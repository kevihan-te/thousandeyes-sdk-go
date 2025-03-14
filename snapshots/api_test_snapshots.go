/*
Test Snapshots API

Creates a new test snapshot in ThousandEyes.

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package snapshots

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


// TestSnapshotsAPIService TestSnapshotsAPI service
type TestSnapshotsAPIService client.Service

type ApiCreateTestSnapshotRequest struct {

	ApiService *TestSnapshotsAPIService
	testId string
	snapshotRequest *SnapshotRequest
	aid *string
}

func (r ApiCreateTestSnapshotRequest) SnapshotRequest(snapshotRequest SnapshotRequest) ApiCreateTestSnapshotRequest {
	r.snapshotRequest = &snapshotRequest
	return r
}

// A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response.
func (r ApiCreateTestSnapshotRequest) Aid(aid string) ApiCreateTestSnapshotRequest {
	r.aid = &aid
	return r
}

func (r ApiCreateTestSnapshotRequest) Execute() (*SnapshotResponse, *http.Response, error) {
	return r.ApiService.CreateTestSnapshotExecute(r)
}

/*
CreateTestSnapshot Create test snapshot

This operation creates a test snapshot based on the properties provided in the POST data.

* To use this endpoint, you need the `Create snapshot shares` permission.
* You can create a maximum of 5 snapshots per organization within a 5-minute interval.
* Snapshots generated through this operation have a 30-day expiration period.
* The time range specified with the `from` and `to` parameters must adhere to one of the following intervals: 1, 2, 4, 6, 12, 24, or 48 hours.
* The `endDate` field of the snapshot must be set to the present or a past date.
* Certain regions may not have public snapshots enabled for compliance reasons. In that case you will get a 403 Forbidden as a response.

**Note**: This operation does not support the creation of operation Agent snapshots.


 @param testId Test ID
 @return ApiCreateTestSnapshotRequest
*/
func (a *TestSnapshotsAPIService) CreateTestSnapshot(testId string ) ApiCreateTestSnapshotRequest {
	return ApiCreateTestSnapshotRequest{
		ApiService: a,
		testId: testId,
	}
}

// Execute executes the request
//  @return SnapshotResponse
func (a *TestSnapshotsAPIService) CreateTestSnapshotExecute(r ApiCreateTestSnapshotRequest) (*SnapshotResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		localVarReturnValue  *SnapshotResponse
	)

	localBasePath := a.Client.GetConfig().ServerURL

	localVarPath := localBasePath + "/tests/{testId}/snapshot"
	localVarPath = strings.Replace(localVarPath, "{"+"testId"+"}", url.PathEscape(request.ParameterValueToString(r.testId, "testId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.snapshotRequest == nil {
		return localVarReturnValue, nil, internalerror.ReportError("snapshotRequest is required and must be specified")
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
	localVarPostBody = r.snapshotRequest
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

func (a *TestSnapshotsAPIService) decodeError(v interface{}, localVarBody []byte, localVarHTTPResponse *http.Response, newErr *internalerror.GenericAPIError) {
    err := a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
    if err != nil {
        newErr.ErrorMessage = err.Error()
        return
    }
    newErr.ErrorMessage = internalerror.FormatErrorMessage(localVarHTTPResponse.Status, &v)
    newErr.Model = v
}
