/*
Instant Tests API

The Instant Tests API operations lets you create and run new instant tests. You will need to be a regular user or have the following permissions:   * `API Access`   * `View tests`  The response does not include the immediate test results. Use the Test Results endpoints to get test results after creating and executing an instant test. You can find the URLs for these endpoints in the _links section of the test definition that is returned when you create the instant test. 

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package instanttests

import (
	"bytes"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	internalerror "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/error"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/request"
	"io"
	"net/http"
	"net/url"
)


// AgentToAgentInstantTestsAPIService AgentToAgentInstantTestsAPI service
type AgentToAgentInstantTestsAPIService client.Service

type ApiCreateAgentToAgentInstantTestRequest struct {

	ApiService *AgentToAgentInstantTestsAPIService
	agentToAgentInstantTestRequest *AgentToAgentInstantTestRequest
	aid *string
	expand *[]ExpandInstantTestOptions
}

func (r ApiCreateAgentToAgentInstantTestRequest) AgentToAgentInstantTestRequest(agentToAgentInstantTestRequest AgentToAgentInstantTestRequest) ApiCreateAgentToAgentInstantTestRequest {
	r.agentToAgentInstantTestRequest = &agentToAgentInstantTestRequest
	return r
}

// A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response.
func (r ApiCreateAgentToAgentInstantTestRequest) Aid(aid string) ApiCreateAgentToAgentInstantTestRequest {
	r.aid = &aid
	return r
}

// (Optional) Indicates if the test sub-resources should be expanded. Defaults to no expansion. To expand the &#x60;agents&#x60; sub-resource, use the query &#x60;?expand&#x3D;agent&#x60;.
func (r ApiCreateAgentToAgentInstantTestRequest) Expand(expand []ExpandInstantTestOptions) ApiCreateAgentToAgentInstantTestRequest {
	r.expand = &expand
	return r
}

func (r ApiCreateAgentToAgentInstantTestRequest) Execute() (*AgentToAgentInstantTestResponse, *http.Response, error) {
	return r.ApiService.CreateAgentToAgentInstantTestExecute(r)
}

/*
CreateAgentToAgentInstantTest Create agent-to-agent instant test

Creates and runs a new agent-to-agent instant test.


 @return ApiCreateAgentToAgentInstantTestRequest
*/
func (a *AgentToAgentInstantTestsAPIService) CreateAgentToAgentInstantTest() ApiCreateAgentToAgentInstantTestRequest {
	return ApiCreateAgentToAgentInstantTestRequest{
		ApiService: a,
	}
}

// Execute executes the request
//  @return AgentToAgentInstantTestResponse
func (a *AgentToAgentInstantTestsAPIService) CreateAgentToAgentInstantTestExecute(r ApiCreateAgentToAgentInstantTestRequest) (*AgentToAgentInstantTestResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		localVarReturnValue  *AgentToAgentInstantTestResponse
	)

	localBasePath := a.Client.GetConfig().ServerURL

	localVarPath := localBasePath + "/tests/agent-to-agent/instant"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.agentToAgentInstantTestRequest == nil {
		return localVarReturnValue, nil, internalerror.ReportError("agentToAgentInstantTestRequest is required and must be specified")
	}

	if r.aid != nil {
		request.ParameterAddToHeaderOrQuery(localVarQueryParams, "aid", r.aid, "")
	}
	if r.expand != nil {
		request.ParameterAddToHeaderOrQuery(localVarQueryParams, "expand", r.expand, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := request.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/hal+json", "application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := request.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.agentToAgentInstantTestRequest
	req, err := a.Client.PrepareRequest(localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &internalerror.GenericAPIError{
			Body:  localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ValidationError
            a.decodeError(&v, localVarBody, localVarHTTPResponse, newErr)
            return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v UnauthorizedError
            a.decodeError(&v, localVarBody, localVarHTTPResponse, newErr)
            return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
            a.decodeError(&v, localVarBody, localVarHTTPResponse, newErr)
            return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
            a.decodeError(&v, localVarBody, localVarHTTPResponse, newErr)
            return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
            a.decodeError(&v, localVarBody, localVarHTTPResponse, newErr)
            return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
            a.decodeError(&v, localVarBody, localVarHTTPResponse, newErr)
            return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 502 {
			var v Error
            a.decodeError(&v, localVarBody, localVarHTTPResponse, newErr)
            return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &internalerror.GenericAPIError{
			Body:  localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

func (a *AgentToAgentInstantTestsAPIService) decodeError(v interface{}, localVarBody []byte, localVarHTTPResponse *http.Response, newErr *internalerror.GenericAPIError) {
    err := a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
    if err != nil {
        newErr.ErrorMessage = err.Error()
        return
    }
    newErr.ErrorMessage = internalerror.FormatErrorMessage(localVarHTTPResponse.Status, &v)
    newErr.Model = v
}
