package eventhub

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ConsumerGroupsClient is the azure Event Hubs client
type ConsumerGroupsClient struct {
	BaseClient
}

// NewConsumerGroupsClient creates an instance of the ConsumerGroupsClient client.
func NewConsumerGroupsClient(subscriptionID string) ConsumerGroupsClient {
	return NewConsumerGroupsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewConsumerGroupsClientWithBaseURI creates an instance of the ConsumerGroupsClient client.
func NewConsumerGroupsClientWithBaseURI(baseURI string, subscriptionID string) ConsumerGroupsClient {
	return ConsumerGroupsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates an Event Hubs consumer group as a nested resource within a Namespace.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// namespaceName - the Namespace name
// eventHubName - the Event Hub name
// consumerGroupName - the consumer group name
// parameters - parameters supplied to create or update a consumer group resource.
func (client ConsumerGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string, parameters ConsumerGroup) (result ConsumerGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConsumerGroupsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}},
		{TargetValue: eventHubName,
			Constraints: []validation.Constraint{{Target: "eventHubName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "eventHubName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: consumerGroupName,
			Constraints: []validation.Constraint{{Target: "consumerGroupName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "consumerGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("eventhub.ConsumerGroupsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, namespaceName, eventHubName, consumerGroupName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ConsumerGroupsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string, parameters ConsumerGroup) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"consumerGroupName": autorest.Encode("path", consumerGroupName),
		"eventHubName":      autorest.Encode("path", eventHubName),
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups/{consumerGroupName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ConsumerGroupsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ConsumerGroupsClient) CreateOrUpdateResponder(resp *http.Response) (result ConsumerGroup, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a consumer group from the specified Event Hub and resource group.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// namespaceName - the Namespace name
// eventHubName - the Event Hub name
// consumerGroupName - the consumer group name
func (client ConsumerGroupsClient) Delete(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConsumerGroupsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}},
		{TargetValue: eventHubName,
			Constraints: []validation.Constraint{{Target: "eventHubName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "eventHubName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: consumerGroupName,
			Constraints: []validation.Constraint{{Target: "consumerGroupName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "consumerGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("eventhub.ConsumerGroupsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, namespaceName, eventHubName, consumerGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ConsumerGroupsClient) DeletePreparer(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"consumerGroupName": autorest.Encode("path", consumerGroupName),
		"eventHubName":      autorest.Encode("path", eventHubName),
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups/{consumerGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ConsumerGroupsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ConsumerGroupsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a description for the specified consumer group.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// namespaceName - the Namespace name
// eventHubName - the Event Hub name
// consumerGroupName - the consumer group name
func (client ConsumerGroupsClient) Get(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (result ConsumerGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConsumerGroupsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}},
		{TargetValue: eventHubName,
			Constraints: []validation.Constraint{{Target: "eventHubName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "eventHubName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: consumerGroupName,
			Constraints: []validation.Constraint{{Target: "consumerGroupName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "consumerGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("eventhub.ConsumerGroupsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, namespaceName, eventHubName, consumerGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ConsumerGroupsClient) GetPreparer(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"consumerGroupName": autorest.Encode("path", consumerGroupName),
		"eventHubName":      autorest.Encode("path", eventHubName),
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups/{consumerGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ConsumerGroupsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ConsumerGroupsClient) GetResponder(resp *http.Response) (result ConsumerGroup, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByEventHub gets all the consumer groups in a Namespace. An empty feed is returned if no consumer group exists in
// the Namespace.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// namespaceName - the Namespace name
// eventHubName - the Event Hub name
// skip - skip is only used if a previous operation returned a partial result. If a previous response contains
// a nextLink element, the value of the nextLink element will include a skip parameter that specifies a
// starting point to use for subsequent calls.
// top - may be used to limit the number of results to the most recent N usageDetails.
func (client ConsumerGroupsClient) ListByEventHub(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, skip *int32, top *int32) (result ConsumerGroupListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConsumerGroupsClient.ListByEventHub")
		defer func() {
			sc := -1
			if result.cglr.Response.Response != nil {
				sc = result.cglr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}},
		{TargetValue: eventHubName,
			Constraints: []validation.Constraint{{Target: "eventHubName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "eventHubName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "skip", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil},
				}}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("eventhub.ConsumerGroupsClient", "ListByEventHub", err.Error())
	}

	result.fn = client.listByEventHubNextResults
	req, err := client.ListByEventHubPreparer(ctx, resourceGroupName, namespaceName, eventHubName, skip, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "ListByEventHub", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByEventHubSender(req)
	if err != nil {
		result.cglr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "ListByEventHub", resp, "Failure sending request")
		return
	}

	result.cglr, err = client.ListByEventHubResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "ListByEventHub", resp, "Failure responding to request")
	}

	return
}

// ListByEventHubPreparer prepares the ListByEventHub request.
func (client ConsumerGroupsClient) ListByEventHubPreparer(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, skip *int32, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"eventHubName":      autorest.Encode("path", eventHubName),
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByEventHubSender sends the ListByEventHub request. The method will close the
// http.Response Body if it receives an error.
func (client ConsumerGroupsClient) ListByEventHubSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByEventHubResponder handles the response to the ListByEventHub request. The method always
// closes the http.Response Body.
func (client ConsumerGroupsClient) ListByEventHubResponder(resp *http.Response) (result ConsumerGroupListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByEventHubNextResults retrieves the next set of results, if any.
func (client ConsumerGroupsClient) listByEventHubNextResults(ctx context.Context, lastResults ConsumerGroupListResult) (result ConsumerGroupListResult, err error) {
	req, err := lastResults.consumerGroupListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "listByEventHubNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByEventHubSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "listByEventHubNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByEventHubResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "listByEventHubNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByEventHubComplete enumerates all values, automatically crossing page boundaries as required.
func (client ConsumerGroupsClient) ListByEventHubComplete(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, skip *int32, top *int32) (result ConsumerGroupListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConsumerGroupsClient.ListByEventHub")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByEventHub(ctx, resourceGroupName, namespaceName, eventHubName, skip, top)
	return
}
