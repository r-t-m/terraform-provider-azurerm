package dataprotection

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ExportJobsOperationResultClient is the open API 2.0 Specs for Azure Data Protection service
type ExportJobsOperationResultClient struct {
	BaseClient
}

// NewExportJobsOperationResultClient creates an instance of the ExportJobsOperationResultClient client.
func NewExportJobsOperationResultClient(subscriptionID string) ExportJobsOperationResultClient {
	return NewExportJobsOperationResultClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewExportJobsOperationResultClientWithBaseURI creates an instance of the ExportJobsOperationResultClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewExportJobsOperationResultClientWithBaseURI(baseURI string, subscriptionID string) ExportJobsOperationResultClient {
	return ExportJobsOperationResultClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets the operation result of operation triggered by Export Jobs API. If the operation is successful, then it
// also contains URL of a Blob and a SAS key to access the same. The blob contains exported jobs in JSON serialized
// format.
// Parameters:
// resourceGroupName - the name of the resource group where the backup vault is present.
// vaultName - the name of the backup vault.
// operationID - operationID which represents the export job.
func (client ExportJobsOperationResultClient) Get(ctx context.Context, resourceGroupName string, vaultName string, operationID string) (result ExportJobsResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExportJobsOperationResultClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, vaultName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.ExportJobsOperationResultClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dataprotection.ExportJobsOperationResultClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.ExportJobsOperationResultClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ExportJobsOperationResultClient) GetPreparer(ctx context.Context, resourceGroupName string, vaultName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationId":       autorest.Encode("path", operationID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2021-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupJobs/operations/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ExportJobsOperationResultClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ExportJobsOperationResultClient) GetResponder(resp *http.Response) (result ExportJobsResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
