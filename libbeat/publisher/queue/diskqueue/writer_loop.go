// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package diskqueue

import (
	"encoding/binary"
	"os"
	"time"

	"github.com/elastic/beats/v7/libbeat/logp"
)

// A segmentedFrame is a data frame waiting to be written to disk along with
// the segment it has been assigned to.
type segmentedFrame struct {
	// The frame to be written to disk.
	frame *writeFrame

	// The segment to which this frame should be written.
	segment *queueSegment
}

// A writer loop request contains a list of writeFrames with the
// segment each should be written to.
//
// Input invariant (segment ids are sorted): If a frame f is included in a
// writerLoopRequest, then every subsequent frame in this and future
// requests must have segment id at least f.segment.id.
//
// That is: we must write all frames for segment 0 before we start writing
// to frame 1, etc. This assumption allows all file operations to happen
// safely in the writer loop without any knowledge of the broader queue state.
type writerLoopRequest struct {
	frames []segmentedFrame
}

// A writerLoopResponse reports the number of bytes written to each
// segment in the request. There is guaranteed to be one entry for each
// segment that appeared in the request, in the same order. If there is
// more than one entry, then all but the last segment have been closed.
type writerLoopResponse struct {
	bytesWritten []int64
}

type writerLoop struct {
	// The settings for the queue that created this loop.
	settings Settings

	// The logger for the writer loop, assigned when the queue creates it.
	logger *logp.Logger

	// The writer loop listens on requestChan for frames to write, and
	// writes them to disk immediately (all queue capacity checking etc. is
	// done by the core loop before sending it to the writer).
	// When this channel is closed, any in-progress writes are aborted and
	// the run loop terminates.
	requestChan chan writerLoopRequest

	// The writer loop sends to responseChan when it has finished handling a
	// request, to signal the core loop that it is ready for the next one.
	responseChan chan writerLoopResponse

	// The most recent segment that has been written to, if there is one.
	// This segment
	currentSegment *queueSegment

	// The file handle corresponding to currentSegment. When currentSegment
	// changes, this handle is closed and a new one is created.
	outputFile *os.File
}

func newWriterLoop(logger *logp.Logger, settings Settings) *writerLoop {
	return &writerLoop{
		logger:   logger,
		settings: settings,

		requestChan:  make(chan writerLoopRequest, 1),
		responseChan: make(chan writerLoopResponse),
	}
}

func (wl *writerLoop) run() {
	for {
		block, ok := <-wl.requestChan
		if !ok {
			// The request channel is closed, we are done
			return
		}
		bytesWritten := wl.processRequest(block)
		wl.responseChan <- writerLoopResponse{bytesWritten: bytesWritten}
	}
}

// processRequest writes the frames in the given request to disk and returns
// the number of bytes written to each segment, in the order they were
// encountered.
func (wl *writerLoop) processRequest(request writerLoopRequest) []int64 {
	// retryWriter wraps the file handle with timed retries.
	// retryWriter.Write is guaranteed to return only if the write
	// completely succeeded or the queue is being closed.
	retryWriter := callbackRetryWriter{retry: wl.retryCallback}

	// We keep track of how many frames are written during this request,
	// and send the associated ACKs to the queue / producer listeners
	// in a batch at the end (since each ACK call can involve a round-trip
	// to the registry).
	totalACKCount := 0
	producerACKCounts := make(map[*diskQueueProducer]int)

	var bytesWritten []int64    // Bytes written to all segments.
	curBytesWritten := int64(0) // Bytes written to the current segment.
outerLoop:
	for _, frameRequest := range request.frames {
		// If the new segment doesn't match the last one, we need to open a new
		// file handle and possibly clean up the old one.
		if wl.currentSegment != frameRequest.segment {
			wl.logger.Debugf(
				"Creating new segment file with id %v\n", frameRequest.segment.id)
			if wl.outputFile != nil {
				// Try to sync to disk, then close the file.
				wl.outputFile.Sync()
				wl.outputFile.Close()
				wl.outputFile = nil
				// We are done with this segment, add the byte count to the list and
				// reset the current counter.
				bytesWritten = append(bytesWritten, curBytesWritten)
				curBytesWritten = 0
			}
			wl.currentSegment = frameRequest.segment
			file, err := wl.currentSegment.getWriterWithRetry(
				wl.settings, wl.retryCallback)
			if err != nil {
				// This can only happen if the queue is being closed; abort.
				break
			}
			wl.outputFile = file
		}
		// Make sure our writer points to the current file handle.
		retryWriter.wrapped = wl.outputFile

		// We have the data and a file to write it to. We are now committed
		// to writing this block unless the queue is closed in the meantime.
		frameSize := uint32(frameRequest.frame.sizeOnDisk())

		// The Write calls below all pass through retryWriter, so they can
		// only return an error if the write should be aborted. Thus, all we
		// need to do when we see an error is break out of the request loop.
		err := binary.Write(retryWriter, binary.LittleEndian, frameSize)
		if err != nil {
			break
		}
		_, err = retryWriter.Write(frameRequest.frame.serialized)
		if err != nil {
			break
		}
		// Compute / write the frame's checksum
		checksum := computeChecksum(frameRequest.frame.serialized)
		err = binary.Write(wl.outputFile, binary.LittleEndian, checksum)
		if err != nil {
			break
		}
		// Write the frame footer's (duplicate) length
		err = binary.Write(wl.outputFile, binary.LittleEndian, frameSize)
		if err != nil {
			break
		}
		// Update the byte count as the last step: that way if we abort while
		// a frame is partially written, we only report up to the last
		// complete frame. (This almost never matters, but it allows for
		// more controlled recovery after a bad shutdown.)
		curBytesWritten += int64(frameSize)

		// Update the ACKs that will be sent at the end of the request.
		totalACKCount++
		if frameRequest.frame.producer.config.ACK != nil {
			producerACKCounts[frameRequest.frame.producer]++
		}

		// Explicitly check if we should abort before starting the next frame.
		select {
		case <-wl.requestChan:
			break outerLoop
		default:
		}
	}
	// Try to sync the written data to disk.
	wl.outputFile.Sync()

	// If the queue has an ACK listener, notify it the frames were written.
	if wl.settings.WriteToDiskListener != nil {
		wl.settings.WriteToDiskListener.OnACK(totalACKCount)
	}

	// Notify any producers with ACK listeners that their frames were written.
	for producer, ackCount := range producerACKCounts {
		producer.config.ACK(ackCount)
	}

	// Return the total byte counts, including the final segment.
	return append(bytesWritten, curBytesWritten)
}

// retryCallback is called (by way of retryCallbackWriter) when there is
// an error writing to a segment file. It pauses for a configurable
// interval and returns true if the operation should be retried (which
// it always should, unless the queue is being closed).
func (wl *writerLoop) retryCallback(err error) bool {
	if writeErrorIsRetriable(err) {
		return true
	}
	// If the error is not immediately retriable, log the error
	// and wait for the retry interval before trying again, but
	// abort if the queue is closed (indicated by the request channel
	// becoming unblocked).
	wl.logger.Errorf("Writing to segment %v: %v",
		wl.currentSegment.id, err)
	select {
	case <-time.After(time.Second):
		// TODO: use a configurable interval here
		return true
	case <-wl.requestChan:
		return false
	}
}
