/*
Tests API

This API supports listing, creating, editing, and deleting Cloud and Enterprise Agent (CEA) based tests. 

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tests

import (
	"bytes"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	internalerror "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/error"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/request"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// HTTPServerTestsAPIService HTTPServerTestsAPI service
type HTTPServerTestsAPIService client.Service

type ApiCreateHttpServerTestRequest struct {

	ApiService *HTTPServerTestsAPIService
	httpServerTestRequest *HttpServerTestRequest
	aid *string
	expand *[]ExpandTestOptions
}

func (r ApiCreateHttpServerTestRequest) HttpServerTestRequest(httpServerTestRequest HttpServerTestRequest) ApiCreateHttpServerTestRequest {
	r.httpServerTestRequest = &httpServerTestRequest
	return r
}

// A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response.
func (r ApiCreateHttpServerTestRequest) Aid(aid string) ApiCreateHttpServerTestRequest {
	r.aid = &aid
	return r
}

// Optional parameter on whether or not to expand the test sub-resources. By default no expansion is going to take place if the query parameter is not present. If the user wishes to expand the &#x60;agents&#x60; sub-resource, they need to pass the &#x60;?expand&#x3D;agent&#x60; query.
func (r ApiCreateHttpServerTestRequest) Expand(expand []ExpandTestOptions) ApiCreateHttpServerTestRequest {
	r.expand = &expand
	return r
}

func (r ApiCreateHttpServerTestRequest) Execute() (*HttpServerTestResponse, *http.Response, error) {
	return r.ApiService.CreateHttpServerTestExecute(r)
}

/*
CreateHttpServerTest Create HTTP Server test

Creates a new HTTP Server test. This method requires Account Admin permissions. 


 @return ApiCreateHttpServerTestRequest
*/
func (a *HTTPServerTestsAPIService) CreateHttpServerTest() ApiCreateHttpServerTestRequest {
	return ApiCreateHttpServerTestRequest{
		ApiService: a,
	}
}

// Execute executes the request
//  @return HttpServerTestResponse
func (a *HTTPServerTestsAPIService) CreateHttpServerTestExecute(r ApiCreateHttpServerTestRequest) (*HttpServerTestResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		localVarReturnValue  *HttpServerTestResponse
	)

	localBasePath := a.Client.GetConfig().ServerURL

	localVarPath := localBasePath + "/tests/http-server"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.httpServerTestRequest == nil {
		return localVarReturnValue, nil, internalerror.ReportError("httpServerTestRequest is required and must be specified")
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
	localVarPostBody = r.httpServerTestRequest
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

type ApiDeleteHttpServerTestRequest struct {

	ApiService *HTTPServerTestsAPIService
	testId string
	aid *string
}

// A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response.
func (r ApiDeleteHttpServerTestRequest) Aid(aid string) ApiDeleteHttpServerTestRequest {
	r.aid = &aid
	return r
}

func (r ApiDeleteHttpServerTestRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteHttpServerTestExecute(r)
}

/*
DeleteHttpServerTest Delete HTTP Server test

Deletes the specified HTTP Server test. This method requires Account Admin permissions.

 @param testId Test ID
 @return ApiDeleteHttpServerTestRequest
*/
func (a *HTTPServerTestsAPIService) DeleteHttpServerTest(testId string ) ApiDeleteHttpServerTestRequest {
	return ApiDeleteHttpServerTestRequest{
		ApiService: a,
		testId: testId,
	}
}

// Execute executes the request
func (a *HTTPServerTestsAPIService) DeleteHttpServerTestExecute(r ApiDeleteHttpServerTestRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
	)

	localBasePath := a.Client.GetConfig().ServerURL

	localVarPath := localBasePath + "/tests/http-server/{testId}"
	localVarPath = strings.Replace(localVarPath, "{"+"testId"+"}", url.PathEscape(request.ParameterValueToString(r.testId, "testId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.aid != nil {
		request.ParameterAddToHeaderOrQuery(localVarQueryParams, "aid", r.aid, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := request.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := request.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &internalerror.GenericAPIError{
			Body:  localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v UnauthorizedError
            a.decodeError(&v, localVarBody, localVarHTTPResponse, newErr)
            return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
            a.decodeError(&v, localVarBody, localVarHTTPResponse, newErr)
            return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
            a.decodeError(&v, localVarBody, localVarHTTPResponse, newErr)
            return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
            a.decodeError(&v, localVarBody, localVarHTTPResponse, newErr)
            return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
            a.decodeError(&v, localVarBody, localVarHTTPResponse, newErr)
            return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 502 {
			var v Error
            a.decodeError(&v, localVarBody, localVarHTTPResponse, newErr)
            return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetHttpServerTestRequest struct {

	ApiService *HTTPServerTestsAPIService
	testId string
	aid *string
	expand *[]ExpandTestOptions
}

// A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response.
func (r ApiGetHttpServerTestRequest) Aid(aid string) ApiGetHttpServerTestRequest {
	r.aid = &aid
	return r
}

// Optional parameter on whether or not to expand the test sub-resources. By default no expansion is going to take place if the query parameter is not present. If the user wishes to expand the &#x60;agents&#x60; sub-resource, they need to pass the &#x60;?expand&#x3D;agent&#x60; query.
func (r ApiGetHttpServerTestRequest) Expand(expand []ExpandTestOptions) ApiGetHttpServerTestRequest {
	r.expand = &expand
	return r
}

func (r ApiGetHttpServerTestRequest) Execute() (*HttpServerTestResponse, *http.Response, error) {
	return r.ApiService.GetHttpServerTestExecute(r)
}

/*
GetHttpServerTest Get HTTP Server test

Returns details for a HTTP Server test, including name, intervals, targets, alert rules and agents.

 @param testId Test ID
 @return ApiGetHttpServerTestRequest
*/
func (a *HTTPServerTestsAPIService) GetHttpServerTest(testId string ) ApiGetHttpServerTestRequest {
	return ApiGetHttpServerTestRequest{
		ApiService: a,
		testId: testId,
	}
}

// Execute executes the request
//  @return HttpServerTestResponse
func (a *HTTPServerTestsAPIService) GetHttpServerTestExecute(r ApiGetHttpServerTestRequest) (*HttpServerTestResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarReturnValue  *HttpServerTestResponse
	)

	localBasePath := a.Client.GetConfig().ServerURL

	localVarPath := localBasePath + "/tests/http-server/{testId}"
	localVarPath = strings.Replace(localVarPath, "{"+"testId"+"}", url.PathEscape(request.ParameterValueToString(r.testId, "testId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.aid != nil {
		request.ParameterAddToHeaderOrQuery(localVarQueryParams, "aid", r.aid, "")
	}
	if r.expand != nil {
		request.ParameterAddToHeaderOrQuery(localVarQueryParams, "expand", r.expand, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiGetHttpServerTestsRequest struct {

	ApiService *HTTPServerTestsAPIService
	aid *string
}

// A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response.
func (r ApiGetHttpServerTestsRequest) Aid(aid string) ApiGetHttpServerTestsRequest {
	r.aid = &aid
	return r
}

func (r ApiGetHttpServerTestsRequest) Execute() (*HttpServerTests, *http.Response, error) {
	return r.ApiService.GetHttpServerTestsExecute(r)
}

/*
GetHttpServerTests List HTTP Server tests

Returns a list of all HTTP Server tests and saved events.


 @return ApiGetHttpServerTestsRequest
*/
func (a *HTTPServerTestsAPIService) GetHttpServerTests() ApiGetHttpServerTestsRequest {
	return ApiGetHttpServerTestsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//  @return HttpServerTests
func (a *HTTPServerTestsAPIService) GetHttpServerTestsExecute(r ApiGetHttpServerTestsRequest) (*HttpServerTests, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarReturnValue  *HttpServerTests
	)

	localBasePath := a.Client.GetConfig().ServerURL

	localVarPath := localBasePath + "/tests/http-server"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.aid != nil {
		request.ParameterAddToHeaderOrQuery(localVarQueryParams, "aid", r.aid, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiUpdateHttpServerTestRequest struct {

	ApiService *HTTPServerTestsAPIService
	testId string
	httpServerTestRequest *HttpServerTestRequest
	aid *string
	expand *[]ExpandTestOptions
}

func (r ApiUpdateHttpServerTestRequest) HttpServerTestRequest(httpServerTestRequest HttpServerTestRequest) ApiUpdateHttpServerTestRequest {
	r.httpServerTestRequest = &httpServerTestRequest
	return r
}

// A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response.
func (r ApiUpdateHttpServerTestRequest) Aid(aid string) ApiUpdateHttpServerTestRequest {
	r.aid = &aid
	return r
}

// Optional parameter on whether or not to expand the test sub-resources. By default no expansion is going to take place if the query parameter is not present. If the user wishes to expand the &#x60;agents&#x60; sub-resource, they need to pass the &#x60;?expand&#x3D;agent&#x60; query.
func (r ApiUpdateHttpServerTestRequest) Expand(expand []ExpandTestOptions) ApiUpdateHttpServerTestRequest {
	r.expand = &expand
	return r
}

func (r ApiUpdateHttpServerTestRequest) Execute() (*HttpServerTestResponse, *http.Response, error) {
	return r.ApiService.UpdateHttpServerTestExecute(r)
}

/*
UpdateHttpServerTest Update HTTP Server test

Updates a HTTP Server test. The target test cannot be a live share or saved event. This method requires Account Admin permissions.

 @param testId Test ID
 @return ApiUpdateHttpServerTestRequest
*/
func (a *HTTPServerTestsAPIService) UpdateHttpServerTest(testId string ) ApiUpdateHttpServerTestRequest {
	return ApiUpdateHttpServerTestRequest{
		ApiService: a,
		testId: testId,
	}
}

// Execute executes the request
//  @return HttpServerTestResponse
func (a *HTTPServerTestsAPIService) UpdateHttpServerTestExecute(r ApiUpdateHttpServerTestRequest) (*HttpServerTestResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		localVarReturnValue  *HttpServerTestResponse
	)

	localBasePath := a.Client.GetConfig().ServerURL

	localVarPath := localBasePath + "/tests/http-server/{testId}"
	localVarPath = strings.Replace(localVarPath, "{"+"testId"+"}", url.PathEscape(request.ParameterValueToString(r.testId, "testId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.httpServerTestRequest == nil {
		return localVarReturnValue, nil, internalerror.ReportError("httpServerTestRequest is required and must be specified")
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
	localVarPostBody = r.httpServerTestRequest
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

func (a *HTTPServerTestsAPIService) decodeError(v interface{}, localVarBody []byte, localVarHTTPResponse *http.Response, newErr *internalerror.GenericAPIError) {
    err := a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
    if err != nil {
        newErr.ErrorMessage = err.Error()
        return
    }
    newErr.ErrorMessage = internalerror.FormatErrorMessage(localVarHTTPResponse.Status, &v)
    newErr.Model = v
}
