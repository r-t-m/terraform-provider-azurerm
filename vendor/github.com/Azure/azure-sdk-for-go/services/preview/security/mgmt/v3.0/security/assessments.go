package security

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

// AssessmentsClient is the API spec for Microsoft.Security (Azure Security Center) resource provider
type AssessmentsClient struct {
	BaseClient
}

// NewAssessmentsClient creates an instance of the AssessmentsClient client.
func NewAssessmentsClient(subscriptionID string, ascLocation string) AssessmentsClient {
	return NewAssessmentsClientWithBaseURI(DefaultBaseURI, subscriptionID, ascLocation)
}

// NewAssessmentsClientWithBaseURI creates an instance of the AssessmentsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewAssessmentsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) AssessmentsClient {
	return AssessmentsClient{NewWithBaseURI(baseURI, subscriptionID, ascLocation)}
}

// CreateOrUpdate create a security assessment on your resource. An assessment metadata that describes this assessment
// must be predefined with the same name before inserting the assessment result
// Parameters:
// resourceID - the identifier of the resource.
// assessmentName - the Assessment Key - Unique key for the assessment type
// assessment - calculated assessment on a pre-defined assessment metadata
func (client AssessmentsClient) CreateOrUpdate(ctx context.Context, resourceID string, assessmentName string, assessment Assessment) (result Assessment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssessmentsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: assessment,
			Constraints: []validation.Constraint{{Target: "assessment.AssessmentProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "assessment.AssessmentProperties.Status", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "assessment.AssessmentProperties.Metadata", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "assessment.AssessmentProperties.Metadata.DisplayName", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "assessment.AssessmentProperties.Metadata.PartnerData", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "assessment.AssessmentProperties.Metadata.PartnerData.PartnerName", Name: validation.Null, Rule: true, Chain: nil},
									{Target: "assessment.AssessmentProperties.Metadata.PartnerData.Secret", Name: validation.Null, Rule: true, Chain: nil},
								}},
						}},
					{Target: "assessment.AssessmentProperties.PartnersData", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "assessment.AssessmentProperties.PartnersData.PartnerName", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "assessment.AssessmentProperties.PartnersData.Secret", Name: validation.Null, Rule: true, Chain: nil},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("security.AssessmentsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceID, assessmentName, assessment)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AssessmentsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.AssessmentsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AssessmentsClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client AssessmentsClient) CreateOrUpdatePreparer(ctx context.Context, resourceID string, assessmentName string, assessment Assessment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"assessmentName": autorest.Encode("path", assessmentName),
		"resourceId":     resourceID,
	}

	const APIVersion = "2020-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceId}/providers/Microsoft.Security/assessments/{assessmentName}", pathParameters),
		autorest.WithJSON(assessment),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client AssessmentsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client AssessmentsClient) CreateOrUpdateResponder(resp *http.Response) (result Assessment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a security assessment on your resource. An assessment metadata that describes this assessment must be
// predefined with the same name before inserting the assessment result
// Parameters:
// resourceID - the identifier of the resource.
// assessmentName - the Assessment Key - Unique key for the assessment type
func (client AssessmentsClient) Delete(ctx context.Context, resourceID string, assessmentName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssessmentsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceID, assessmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AssessmentsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "security.AssessmentsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AssessmentsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AssessmentsClient) DeletePreparer(ctx context.Context, resourceID string, assessmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"assessmentName": autorest.Encode("path", assessmentName),
		"resourceId":     resourceID,
	}

	const APIVersion = "2020-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceId}/providers/Microsoft.Security/assessments/{assessmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AssessmentsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AssessmentsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get a security assessment on your scanned resource
// Parameters:
// resourceID - the identifier of the resource.
// assessmentName - the Assessment Key - Unique key for the assessment type
// expand - oData expand. Optional.
func (client AssessmentsClient) Get(ctx context.Context, resourceID string, assessmentName string, expand ExpandEnum) (result Assessment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssessmentsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceID, assessmentName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AssessmentsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.AssessmentsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AssessmentsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client AssessmentsClient) GetPreparer(ctx context.Context, resourceID string, assessmentName string, expand ExpandEnum) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"assessmentName": autorest.Encode("path", assessmentName),
		"resourceId":     resourceID,
	}

	const APIVersion = "2020-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(expand)) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceId}/providers/Microsoft.Security/assessments/{assessmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AssessmentsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AssessmentsClient) GetResponder(resp *http.Response) (result Assessment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get security assessments on all your scanned resources inside a scope
// Parameters:
// scope - scope of the query, can be subscription (/subscriptions/0b06d9ea-afe6-4779-bd59-30e5c2d9d13f) or
// management group (/providers/Microsoft.Management/managementGroups/mgName).
func (client AssessmentsClient) List(ctx context.Context, scope string) (result AssessmentListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssessmentsClient.List")
		defer func() {
			sc := -1
			if result.al.Response.Response != nil {
				sc = result.al.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, scope)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AssessmentsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.al.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.AssessmentsClient", "List", resp, "Failure sending request")
		return
	}

	result.al, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AssessmentsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.al.hasNextLink() && result.al.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client AssessmentsClient) ListPreparer(ctx context.Context, scope string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope": scope,
	}

	const APIVersion = "2020-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Security/assessments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AssessmentsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AssessmentsClient) ListResponder(resp *http.Response) (result AssessmentList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client AssessmentsClient) listNextResults(ctx context.Context, lastResults AssessmentList) (result AssessmentList, err error) {
	req, err := lastResults.assessmentListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "security.AssessmentsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "security.AssessmentsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AssessmentsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client AssessmentsClient) ListComplete(ctx context.Context, scope string) (result AssessmentListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssessmentsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, scope)
	return
}
