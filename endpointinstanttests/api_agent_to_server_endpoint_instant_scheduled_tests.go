/*
Endpoint Instant Scheduled Tests API

 You can create and execute a new endpoint instant scheduled test within ThousandEyes using this API. The test parameters are specified in the `POST` data.  The following applies to the Endpoint Instant Scheduled Tests API:  * To initiate the creation and execution of an instant scheduled test, the user must possess the `Edit endpoint tests` permission.  * Upon successful creation of an instant scheduled test, the API responds with an HTTP/201 CREATED status code and return the test definition. * It's important to note that the response does not include the results of the instant scheduled test. To retrieve test results, users can utilize the Endpoint Test Data endpoints. The URLs for these API test data endpoints are provided within the test definition output when an instant scheduled test is created. 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpointinstanttests

import (
	"bytes"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	internalerror "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/error"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/request"
	"io"
	"net/http"
	"net/url"
)


// AgentToServerEndpointInstantScheduledTestsAPIService AgentToServerEndpointInstantScheduledTestsAPI service
type AgentToServerEndpointInstantScheduledTestsAPIService client.Service

type ApiCreateAgentToServerScheduledInstantTestRequest struct {

	ApiService *AgentToServerEndpointInstantScheduledTestsAPIService
	endpointAgentToServerInstantTest *EndpointAgentToServerInstantTest
	aid *string
}

func (r ApiCreateAgentToServerScheduledInstantTestRequest) EndpointAgentToServerInstantTest(endpointAgentToServerInstantTest EndpointAgentToServerInstantTest) ApiCreateAgentToServerScheduledInstantTestRequest {
	r.endpointAgentToServerInstantTest = &endpointAgentToServerInstantTest
	return r
}

// A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response.
func (r ApiCreateAgentToServerScheduledInstantTestRequest) Aid(aid string) ApiCreateAgentToServerScheduledInstantTestRequest {
	r.aid = &aid
	return r
}

func (r ApiCreateAgentToServerScheduledInstantTestRequest) Execute() (*EndpointAgentToServerTest, *http.Response, error) {
	return r.ApiService.CreateAgentToServerScheduledInstantTestExecute(r)
}

/*
CreateAgentToServerScheduledInstantTest Run agent to server instant scheduled test

Creates and runs a new endpoint agent to server instant scheduled test in ThousandEyes.


 @return ApiCreateAgentToServerScheduledInstantTestRequest
*/
func (a *AgentToServerEndpointInstantScheduledTestsAPIService) CreateAgentToServerScheduledInstantTest() ApiCreateAgentToServerScheduledInstantTestRequest {
	return ApiCreateAgentToServerScheduledInstantTestRequest{
		ApiService: a,
	}
}

// Execute executes the request
//  @return EndpointAgentToServerTest
func (a *AgentToServerEndpointInstantScheduledTestsAPIService) CreateAgentToServerScheduledInstantTestExecute(r ApiCreateAgentToServerScheduledInstantTestRequest) (*EndpointAgentToServerTest, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		localVarReturnValue  *EndpointAgentToServerTest
	)

	localBasePath := a.Client.GetConfig().ServerURL

	localVarPath := localBasePath + "/endpoint/tests/scheduled-tests/agent-to-server/instant"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.endpointAgentToServerInstantTest == nil {
		return localVarReturnValue, nil, internalerror.ReportError("endpointAgentToServerInstantTest is required and must be specified")
	}

	if r.aid != nil {
		request.ParameterAddToHeaderOrQuery(localVarQueryParams, "aid", r.aid, "")
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
	localVarPostBody = r.endpointAgentToServerInstantTest
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

func (a *AgentToServerEndpointInstantScheduledTestsAPIService) decodeError(v interface{}, localVarBody []byte, localVarHTTPResponse *http.Response, newErr *internalerror.GenericAPIError) {
    err := a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
    if err != nil {
        newErr.ErrorMessage = err.Error()
        return
    }
    newErr.ErrorMessage = internalerror.FormatErrorMessage(localVarHTTPResponse.Status, &v)
    newErr.Model = v
}
