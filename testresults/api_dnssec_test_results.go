/*
Test Results API

Get test result metrics for Cloud and Enterprise Agent tests.

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package testresults

import (
	"bytes"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	internalerror "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/error"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/request"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)


// DNSSECTestResultsAPIService DNSSECTestResultsAPI service
type DNSSECTestResultsAPIService client.Service

type ApiGetTestDnsSecResultsRequest struct {

	ApiService *DNSSECTestResultsAPIService
	testId string
	aid *string
	window *string
	startDate *time.Time
	endDate *time.Time
	cursor *string
}

// A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response.
func (r ApiGetTestDnsSecResultsRequest) Aid(aid string) ApiGetTestDnsSecResultsRequest {
	r.aid = &aid
	return r
}

// A dynamic time interval up to the current time of the request. Specify the interval as a number followed by an optional type: &#x60;s&#x60; for seconds (default if no type is specified), &#x60;m&#x60; for minutes, &#x60;h&#x60; for hours, &#x60;d&#x60; for days, and &#x60;w&#x60; for weeks. For a precise date range, use &#x60;startDate&#x60; and &#x60;endDate&#x60;.
func (r ApiGetTestDnsSecResultsRequest) Window(window string) ApiGetTestDnsSecResultsRequest {
	r.window = &window
	return r
}

// Use with the &#x60;endDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;.
func (r ApiGetTestDnsSecResultsRequest) StartDate(startDate time.Time) ApiGetTestDnsSecResultsRequest {
	r.startDate = &startDate
	return r
}

// Defaults to current time the request is made. Use with the &#x60;startDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;.
func (r ApiGetTestDnsSecResultsRequest) EndDate(endDate time.Time) ApiGetTestDnsSecResultsRequest {
	r.endDate = &endDate
	return r
}

// (Optional) Opaque cursor used for pagination. Clients should use &#x60;next&#x60; value from &#x60;_links&#x60; instead of this parameter.
func (r ApiGetTestDnsSecResultsRequest) Cursor(cursor string) ApiGetTestDnsSecResultsRequest {
	r.cursor = &cursor
	return r
}

func (r ApiGetTestDnsSecResultsRequest) Execute() (*DnssecTestResults, *http.Response, error) {
	return r.ApiService.GetTestDnsSecResultsExecute(r)
}

/*
GetTestDnsSecResults Get DNSSEC test results

Returns the keychain validity for a record on a domain secured using DNSSEC extensions.


 @param testId Test ID
 @return ApiGetTestDnsSecResultsRequest
*/
func (a *DNSSECTestResultsAPIService) GetTestDnsSecResults(testId string ) ApiGetTestDnsSecResultsRequest {
	return ApiGetTestDnsSecResultsRequest{
		ApiService: a,
		testId: testId,
	}
}

// Execute executes the request
//  @return DnssecTestResults
func (a *DNSSECTestResultsAPIService) GetTestDnsSecResultsExecute(r ApiGetTestDnsSecResultsRequest) (*DnssecTestResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarReturnValue  *DnssecTestResults
	)

	localBasePath := a.Client.GetConfig().ServerURL

	localVarPath := localBasePath + "/test-results/{testId}/dnssec"
	localVarPath = strings.Replace(localVarPath, "{"+"testId"+"}", url.PathEscape(request.ParameterValueToString(r.testId, "testId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.aid != nil {
		request.ParameterAddToHeaderOrQuery(localVarQueryParams, "aid", r.aid, "")
	}
	if r.window != nil {
		request.ParameterAddToHeaderOrQuery(localVarQueryParams, "window", r.window, "")
	}
	if r.startDate != nil {
		request.ParameterAddToHeaderOrQuery(localVarQueryParams, "startDate", r.startDate, "")
	}
	if r.endDate != nil {
		request.ParameterAddToHeaderOrQuery(localVarQueryParams, "endDate", r.endDate, "")
	}
	if r.cursor != nil {
		request.ParameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "")
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

func (a *DNSSECTestResultsAPIService) decodeError(v interface{}, localVarBody []byte, localVarHTTPResponse *http.Response, newErr *internalerror.GenericAPIError) {
    err := a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
    if err != nil {
        newErr.ErrorMessage = err.Error()
        return
    }
    newErr.ErrorMessage = internalerror.FormatErrorMessage(localVarHTTPResponse.Status, &v)
    newErr.Model = v
}
