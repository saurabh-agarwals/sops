package consumption

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
	"net/http"
)

// AggregatedCostClient is the consumption management client provides access to consumption resources for Azure
// Enterprise Subscriptions.
type AggregatedCostClient struct {
	BaseClient
}

// NewAggregatedCostClient creates an instance of the AggregatedCostClient client.
func NewAggregatedCostClient(subscriptionID string) AggregatedCostClient {
	return NewAggregatedCostClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAggregatedCostClientWithBaseURI creates an instance of the AggregatedCostClient client.
func NewAggregatedCostClientWithBaseURI(baseURI string, subscriptionID string) AggregatedCostClient {
	return AggregatedCostClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetByManagementGroup provides the aggregate cost of a management group and all child management groups by current
// billing period.
// Parameters:
// managementGroupID - azure Management Group ID.
func (client AggregatedCostClient) GetByManagementGroup(ctx context.Context, managementGroupID string) (result ManagementGroupAggregatedCostResult, err error) {
	req, err := client.GetByManagementGroupPreparer(ctx, managementGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.AggregatedCostClient", "GetByManagementGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByManagementGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.AggregatedCostClient", "GetByManagementGroup", resp, "Failure sending request")
		return
	}

	result, err = client.GetByManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.AggregatedCostClient", "GetByManagementGroup", resp, "Failure responding to request")
	}

	return
}

// GetByManagementGroupPreparer prepares the GetByManagementGroup request.
func (client AggregatedCostClient) GetByManagementGroupPreparer(ctx context.Context, managementGroupID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId": autorest.Encode("path", managementGroupID),
	}

	const APIVersion = "2018-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Consumption/aggregatedcost", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByManagementGroupSender sends the GetByManagementGroup request. The method will close the
// http.Response Body if it receives an error.
func (client AggregatedCostClient) GetByManagementGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetByManagementGroupResponder handles the response to the GetByManagementGroup request. The method always
// closes the http.Response Body.
func (client AggregatedCostClient) GetByManagementGroupResponder(resp *http.Response) (result ManagementGroupAggregatedCostResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetForBillingPeriodByManagementGroup provides the aggregate cost of a management group and all child management
// groups by specified billing period
// Parameters:
// managementGroupID - azure Management Group ID.
// billingPeriodName - billing Period Name.
func (client AggregatedCostClient) GetForBillingPeriodByManagementGroup(ctx context.Context, managementGroupID string, billingPeriodName string) (result ManagementGroupAggregatedCostResult, err error) {
	req, err := client.GetForBillingPeriodByManagementGroupPreparer(ctx, managementGroupID, billingPeriodName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.AggregatedCostClient", "GetForBillingPeriodByManagementGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetForBillingPeriodByManagementGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.AggregatedCostClient", "GetForBillingPeriodByManagementGroup", resp, "Failure sending request")
		return
	}

	result, err = client.GetForBillingPeriodByManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.AggregatedCostClient", "GetForBillingPeriodByManagementGroup", resp, "Failure responding to request")
	}

	return
}

// GetForBillingPeriodByManagementGroupPreparer prepares the GetForBillingPeriodByManagementGroup request.
func (client AggregatedCostClient) GetForBillingPeriodByManagementGroupPreparer(ctx context.Context, managementGroupID string, billingPeriodName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingPeriodName": autorest.Encode("path", billingPeriodName),
		"managementGroupId": autorest.Encode("path", managementGroupID),
	}

	const APIVersion = "2018-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}/Microsoft.Consumption/aggregatedcost", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetForBillingPeriodByManagementGroupSender sends the GetForBillingPeriodByManagementGroup request. The method will close the
// http.Response Body if it receives an error.
func (client AggregatedCostClient) GetForBillingPeriodByManagementGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetForBillingPeriodByManagementGroupResponder handles the response to the GetForBillingPeriodByManagementGroup request. The method always
// closes the http.Response Body.
func (client AggregatedCostClient) GetForBillingPeriodByManagementGroupResponder(resp *http.Response) (result ManagementGroupAggregatedCostResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
