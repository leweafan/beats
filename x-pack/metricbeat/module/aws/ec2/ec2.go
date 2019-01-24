// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package ec2

import (
	"fmt"
	"strconv"
	"time"

	awssdk "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/defaults"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/cloudwatchiface"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/ec2iface"
	"github.com/pkg/errors"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/common/cfgwarn"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/metricbeat/mb"
	"github.com/elastic/beats/x-pack/metricbeat/module/aws"
)

var metricsetName = "ec2"

// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host defined in the module's configuration. After the
// MetricSet has been created then Fetch will begin to be called periodically.
func init() {
	mb.Registry.MustAddMetricSet(aws.ModuleName, metricsetName, New,
		mb.DefaultMetricSet(),
	)
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	*aws.MetricSet
	moduleConfig   *aws.Config
	awsConfig      *awssdk.Config
	regionsList    []string
	durationString string
	periodInSec    int
	logger         *logp.Logger
}

// metricIDNameMap is a translating map between createMetricDataQuery id
// and aws ec2 module metric name, cloudwatch ec2 metric name.
var metricIDNameMap = map[string][]string{
	"cpu1":     {"cpu.total.pct", "CPUUtilization"},
	"cpu2":     {"cpu.credit_usage", "CPUCreditUsage"},
	"cpu3":     {"cpu.credit_balance", "CPUCreditBalance"},
	"cpu4":     {"cpu.surplus_credit_balance", "CPUSurplusCreditBalance"},
	"cpu5":     {"cpu.surplus_credits_charged", "CPUSurplusCreditsCharged"},
	"network1": {"network.in.packets", "NetworkPacketsIn"},
	"network2": {"network.out.packets", "NetworkPacketsOut"},
	"network3": {"network.in.bytes", "NetworkIn"},
	"network4": {"network.out.bytes", "NetworkOut"},
	"disk1":    {"diskio.read.bytes", "DiskReadBytes"},
	"disk2":    {"diskio.write.bytes", "DiskWriteBytes"},
	"disk3":    {"diskio.read.ops", "DiskReadOps"},
	"disk4":    {"diskio.write.ops", "DiskWriteOps"},
	"status1":  {"status.check_failed", "StatusCheckFailed"},
	"status2":  {"status.check_failed_system", "StatusCheckFailed_System"},
	"status3":  {"status.check_failed_instance", "StatusCheckFailed_Instance"},
}

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Beta("The aws ec2 metricset is beta.")
	ec2Logger := logp.NewLogger(aws.ModuleName)

	moduleConfig := aws.Config{}
	if err := base.Module().UnpackConfig(&moduleConfig); err != nil {
		return nil, err
	}

	if moduleConfig.Period == "" {
		err := errors.New("period is not set in AWS module config")
		ec2Logger.Error(err)
	}

	metricSet, err := aws.NewMetricSet(base)
	if err != nil {
		return nil, errors.Wrap(err, "error creating aws metricset")
	}

	// Get a list of regions
	awsConfig := defaults.Config()
	awsCreds := awssdk.Credentials{
		AccessKeyID:     moduleConfig.AccessKeyID,
		SecretAccessKey: moduleConfig.SecretAccessKey,
	}
	if moduleConfig.SessionToken != "" {
		awsCreds.SessionToken = moduleConfig.SessionToken
	}

	awsConfig.Credentials = awssdk.StaticCredentialsProvider{
		Value: awsCreds,
	}

	awsConfig.Region = moduleConfig.DefaultRegion
	svcEC2 := ec2.New(awsConfig)
	regionsList, err := getRegions(svcEC2)
	if err != nil {
		err = errors.Wrap(err, "getRegions failed")
		ec2Logger.Error(err.Error())
	}

	// Calculate duration based on period
	durationString, periodSec, err := convertPeriodToDuration(moduleConfig.Period)
	if err != nil {
		ec2Logger.Error(err.Error())
		return nil, err
	}

	// Check if period is set to be multiple of 60s or 300s
	remainder300 := periodSec % 300
	remainder60 := periodSec % 60
	if remainder300 != 0 || remainder60 != 0 {
		err := errors.New("period needs to be set to 60s (or a multiple of 60s) if detailed monitoring is " +
			"enabled for EC2 instances or set to 300s (or a multiple of 300s) if EC2 instances has basic monitoring. " +
			"To avoid data missing or extra costs, please make sure period is set correctly in config.yml")
		ec2Logger.Info(err)
	}

	return &MetricSet{
		MetricSet:      metricSet,
		moduleConfig:   &moduleConfig,
		awsConfig:      &awsConfig,
		regionsList:    regionsList,
		durationString: durationString,
		periodInSec:    periodSec,
		logger:         ec2Logger,
	}, nil
}

// Fetch methods implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(report mb.ReporterV2) {
	for _, regionName := range m.regionsList {
		m.awsConfig.Region = regionName
		svcEC2 := ec2.New(*m.awsConfig)
		instanceIDs, instancesOutputs, err := getInstancesPerRegion(svcEC2)
		if err != nil {
			err = errors.Wrap(err, "getInstancesPerRegion failed, skipping region "+regionName)
			m.logger.Errorf(err.Error())
			report.Error(err)
			continue
		}

		svcCloudwatch := cloudwatch.New(*m.awsConfig)
		for _, instanceID := range instanceIDs {
			init := true
			getMetricDataOutput := &cloudwatch.GetMetricDataOutput{NextToken: nil}
			for init || getMetricDataOutput.NextToken != nil {
				init = false
				output, err := getMetricDataPerRegion(m.durationString, m.periodInSec, instanceID, getMetricDataOutput.NextToken, svcCloudwatch)
				if err != nil {
					err = errors.Wrap(err, "getMetricDataPerRegion failed, skipping region "+regionName+" for instance "+instanceID)
					m.logger.Error(err.Error())
					report.Error(err)
					continue
				}
				getMetricDataOutput.MetricDataResults = append(getMetricDataOutput.MetricDataResults, output.MetricDataResults...)
			}

			event, info, err := createCloudWatchEvents(getMetricDataOutput, instanceID, instancesOutputs[instanceID], regionName)
			if info != "" {
				m.logger.Info(info)
			}

			if err != nil {
				m.logger.Error(err.Error())
				event.Error = err
				report.Event(event)
				continue
			}
			report.Event(event)
		}
	}
}

func getRegions(svc ec2iface.EC2API) (regionsList []string, err error) {
	input := &ec2.DescribeRegionsInput{}
	req := svc.DescribeRegionsRequest(input)
	output, err := req.Send()
	if err != nil {
		err = errors.Wrap(err, "Failed DescribeRegions")
		return
	}
	for _, region := range output.Regions {
		regionsList = append(regionsList, *region.RegionName)
	}
	return
}

func createCloudWatchEvents(getMetricDataOutput *cloudwatch.GetMetricDataOutput, instanceID string, instanceOutput ec2.Instance, regionName string) (event mb.Event, info string, err error) {
	event.Service = metricsetName
	event.RootFields = common.MapStr{}
	mapOfRootFieldsResults := make(map[string]interface{})
	mapOfRootFieldsResults["service.name"] = metricsetName
	mapOfRootFieldsResults["cloud.provider"] = metricsetName
	mapOfRootFieldsResults["cloud.instance.id"] = instanceID

	machineType, err := instanceOutput.InstanceType.MarshalValue()
	if err != nil {
		err = errors.Wrap(err, "instance.InstanceType.MarshalValue failed")
		return
	}

	mapOfRootFieldsResults["cloud.machine.type"] = machineType
	mapOfRootFieldsResults["cloud.availability_zone"] = *instanceOutput.Placement.AvailabilityZone
	mapOfRootFieldsResults["cloud.image.id"] = *instanceOutput.ImageId
	mapOfRootFieldsResults["cloud.region"] = regionName

	resultRootFields, err := eventMapping(mapOfRootFieldsResults, schemaRootFields)
	if err != nil {
		err = errors.Wrap(err, "Error trying to apply schema schemaRootFields in AWS EC2 metricbeat module.")
		return
	}
	event.RootFields = resultRootFields

	mapOfMetricSetFieldResults := make(map[string]interface{})
	for _, output := range getMetricDataOutput.MetricDataResults {
		if len(output.Values) == 0 {
			continue
		}
		metricKey := metricIDNameMap[*output.Id]
		mapOfMetricSetFieldResults[metricKey[0]] = fmt.Sprint(output.Values[0])
	}

	if len(mapOfMetricSetFieldResults) <= 3 {
		info = "Missing Cloudwatch data for instance " + instanceID + ". This is expected for a new instance during the " +
			"first data collection. If this shows up multiple times, please recheck the period setting in config."
		return
	}

	resultMetricSetFields, err := eventMapping(mapOfMetricSetFieldResults, schemaMetricSetFields)
	if err != nil {
		err = errors.Wrap(err, "Error trying to apply schema schemaMetricSetFields in AWS EC2 metricbeat module.")
		return
	}
	event.MetricSetFields = resultMetricSetFields
	return
}

func getInstancesPerRegion(svc ec2iface.EC2API) (instanceIDs []string, instancesOutputs map[string]ec2.Instance, err error) {
	instancesOutputs = map[string]ec2.Instance{}
	output := ec2.DescribeInstancesOutput{NextToken: nil}
	init := true
	for init || output.NextToken != nil {
		init = false
		describeInstanceInput := &ec2.DescribeInstancesInput{
			Filters: []ec2.Filter{
				{
					Name:   awssdk.String("instance-state-name"),
					Values: []string{"running"},
				},
			},
		}

		req := svc.DescribeInstancesRequest(describeInstanceInput)
		output, err := req.Send()
		if err != nil {
			err = errors.Wrap(err, "Error DescribeInstances")
			return nil, nil, err
		}

		for _, reservation := range output.Reservations {
			for _, instance := range reservation.Instances {
				instanceIDs = append(instanceIDs, *instance.InstanceId)
				instancesOutputs[*instance.InstanceId] = instance
			}
		}
	}
	return
}

func convertPeriodToDuration(period string) (string, int, error) {
	// Amazon EC2 sends metrics to Amazon CloudWatch with 5-minute default frequency.
	// If detailed monitoring is enabled, then data will be available in 1-minute period.
	// Set starttime double the default frequency earlier than the endtime in order to make sure
	// GetMetricDataRequest gets the latest data point for each metric.
	numberPeriod, err := strconv.Atoi(period[0 : len(period)-1])
	if err != nil {
		return "", 0, err
	}

	unitPeriod := period[len(period)-1:]
	switch unitPeriod {
	case "s":
		duration := "-" + strconv.Itoa(numberPeriod*2) + unitPeriod
		return duration, numberPeriod, nil
	case "m":
		duration := "-" + strconv.Itoa(numberPeriod*2) + unitPeriod
		periodInSec := numberPeriod * 60
		return duration, periodInSec, nil
	default:
		err = errors.New("invalid period in config. Please reset period in config")
		duration := "-" + strconv.Itoa(numberPeriod*2) + "s"
		return duration, numberPeriod, err
	}
}

func getMetricDataPerRegion(durationString string, periodInSec int, instanceID string, nextToken *string, svc cloudwatchiface.CloudWatchAPI) (*cloudwatch.GetMetricDataOutput, error) {
	endTime := time.Now()
	duration, err := time.ParseDuration(durationString)
	if err != nil {
		logp.Error(errors.Wrap(err, "Error ParseDuration"))
		return nil, err
	}

	startTime := endTime.Add(duration)

	dimName := "InstanceId"
	dim := cloudwatch.Dimension{
		Name:  &dimName,
		Value: &instanceID,
	}

	metricDataQueries := []cloudwatch.MetricDataQuery{}
	for metricID, metricName := range metricIDNameMap {
		metricDataQuery := createMetricDataQuery(metricID, metricName[1], periodInSec, []cloudwatch.Dimension{dim})
		metricDataQueries = append(metricDataQueries, metricDataQuery)
	}

	getMetricDataInput := &cloudwatch.GetMetricDataInput{
		NextToken:         nextToken,
		StartTime:         &startTime,
		EndTime:           &endTime,
		MetricDataQueries: metricDataQueries,
	}

	req := svc.GetMetricDataRequest(getMetricDataInput)
	getMetricDataOutput, err := req.Send()
	if err != nil {
		logp.Error(errors.Wrap(err, "Error GetMetricDataInput"))
		return nil, err
	}
	return getMetricDataOutput, nil
}

func createMetricDataQuery(id string, metricName string, periodInSec int, dimensions []cloudwatch.Dimension) (metricDataQuery cloudwatch.MetricDataQuery) {
	namespace := "AWS/EC2"
	statistic := "Average"
	period := int64(periodInSec)

	metric := cloudwatch.Metric{
		Namespace:  &namespace,
		MetricName: &metricName,
		Dimensions: dimensions,
	}

	metricDataQuery = cloudwatch.MetricDataQuery{
		Id: &id,
		MetricStat: &cloudwatch.MetricStat{
			Period: &period,
			Stat:   &statistic,
			Metric: &metric,
		},
	}
	return
}
