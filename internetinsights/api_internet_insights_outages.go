/*
Internet Insights API

We are happy to announce the release of the Internet Insights API set. This limited release includes endpoints that:  * Make our catalog provider and Internet outage data accessible to API users. * Provide access to advanced filtering, which is part of our next-generation API efforts to allow API users to fine-tune queries across all of our APIs in a consistent manner.  Internet Insights provide visibility into core Internet infrastructure, including ISPs, DNS providers, IaaS, CDNs , and SaaS providers. It tracks the macro-level impact of Internet events on individual users and enterprise networks connecting at the edge of the Internet. These events include Outages, Routing hijacks and leaks, DDoS attacks, And political interference, among others.  Future releases of the Internet Insights API set will further unlock access to core Internet Insights functionality, unlocking potential integrations to enrich customer process flows.  For more information about Internet Insights, see the [Internet Insights](https://docs.thousandeyes.com/product-documentation/internet-insights). 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package internetinsights

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


// InternetInsightsOutagesAPIService InternetInsightsOutagesAPI service
type InternetInsightsOutagesAPIService client.Service

type ApiFilterOutagesRequest struct {

	ApiService *InternetInsightsOutagesAPIService
	apiOutageFilter *ApiOutageFilter
	aid *string
}

func (r ApiFilterOutagesRequest) ApiOutageFilter(apiOutageFilter ApiOutageFilter) ApiFilterOutagesRequest {
	r.apiOutageFilter = &apiOutageFilter
	return r
}

// A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response.
func (r ApiFilterOutagesRequest) Aid(aid string) ApiFilterOutagesRequest {
	r.aid = &aid
	return r
}

func (r ApiFilterOutagesRequest) Execute() (*ApiOutagesResponse, *http.Response, error) {
	return r.ApiService.FilterOutagesExecute(r)
}

/*
FilterOutages List network and application outages

Returns a list of network and application outages using a filter object. Advanced Filter persistance is not currently supported.

<b>Note:</b> Support for pagination will be added in the future.

 ## Samples Queries with Different Filter Permutations

 ### Window
 ```
 curl --location --request POST 'https://api.thousandeyes.com/v7/internet-insights/outages/filter'
--header 'Authorization: Bearer $token
--header 'Accept-Encoding: application/gzip'
--header 'Content-Type: application/json'
--data-raw '{
  "window" : "1d"
  }'
```

### Date Range
```
curl --location --request POST 'https://api.thousandeyes.com/v7/internet-insights/outages/filter'
--header 'Authorization: Bearer $token'
--header 'Content-Type: application/json'
--data-raw '{
    "startDate": "2022-03-01T01:30:00Z",
    "endDate"  : "2022-03-01T23:30:15Z"
  }'
```

### Date Range with Scope
```
curl --location --request POST 'https://api.thousandeyes.com/v7/internet-insights/outages/filter'
--header 'Authorization: Bearer $token'
--header 'Content-Type: application/json'
--data-raw '{
    "startDate": "2022-03-01T01:30:00Z",
    "endDate"  : "2022-03-01T23:30:15Z",
    "outageScope": "with-affected-test"
  }'
```
### Date Range with Network
```
  curl --location --request POST 'https://api.thousandeyes.com/v7/internet-insights/outages/filter'
  --header 'Authorization: Bearer $token'
  --header 'Content-Type: application/json'
  --data-raw '{
      "startDate": "2022-03-01T01:30:00Z",
      "endDate"  : "2022-03-01T23:30:15Z",
      "interfaceNetwork":  ["Telianet"]
    }'
```

### Date Range with Application
```
curl --location --request POST 'https://api.thousandeyes.com/v7/internet-insights/outages/filter'
  --header 'Authorization: Bearer $token'
  --header 'Content-Type: application/json'
  --data-raw '{
      "startDate": "2022-03-01T01:30:00Z",
      "endDate"  : "2022-03-01T23:30:15Z",
      "applicationName": ["Google"]
  }'
```
### Date Range with Provider
```
curl --location --request POST 'https://api.thousandeyes.com/v7/internet-insights/outages/filter'
--header 'Authorization: Bearer $token'
--header 'Content-Type: application/json'
--data-raw '{
      "startDate": "2022-03-01T01:30:00Z",
      "endDate"  : "2022-03-01T23:30:15Z",
      "providerName": ["Century Link", "Microsoft"]
  }'

```
### Date Range with Application and Scope
```
curl --location --request POST 'https://api.thousandeyes.com/v7/internet-insights/outages/filter'
--header 'Authorization: Bearer $token'
--header 'Content-Type: application/json'
--data-raw '{
    "startDate": "2022-03-01T01:30:00Z",
    "endDate"  : "2022-03-01T23:30:15Z",
    "outageScope": "all",
    "applicationName": ["Google"]
}'
```



 @return ApiFilterOutagesRequest
*/
func (a *InternetInsightsOutagesAPIService) FilterOutages() ApiFilterOutagesRequest {
	return ApiFilterOutagesRequest{
		ApiService: a,
	}
}

// Execute executes the request
//  @return ApiOutagesResponse
func (a *InternetInsightsOutagesAPIService) FilterOutagesExecute(r ApiFilterOutagesRequest) (*ApiOutagesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		localVarReturnValue  *ApiOutagesResponse
	)

	localBasePath := a.Client.GetConfig().ServerURL

	localVarPath := localBasePath + "/internet-insights/outages/filter"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiOutageFilter == nil {
		return localVarReturnValue, nil, internalerror.ReportError("apiOutageFilter is required and must be specified")
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
	localVarPostBody = r.apiOutageFilter
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

type ApiGetAppOutageRequest struct {

	ApiService *InternetInsightsOutagesAPIService
	outageId string
	aid *string
}

// A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response.
func (r ApiGetAppOutageRequest) Aid(aid string) ApiGetAppOutageRequest {
	r.aid = &aid
	return r
}

func (r ApiGetAppOutageRequest) Execute() (*ApiApplicationOutageDetails, *http.Response, error) {
	return r.ApiService.GetAppOutageExecute(r)
}

/*
GetAppOutage Retrieve application outage

Returns the details of an application outage.


 @param outageId
 @return ApiGetAppOutageRequest
*/
func (a *InternetInsightsOutagesAPIService) GetAppOutage(outageId string ) ApiGetAppOutageRequest {
	return ApiGetAppOutageRequest{
		ApiService: a,
		outageId: outageId,
	}
}

// Execute executes the request
//  @return ApiApplicationOutageDetails
func (a *InternetInsightsOutagesAPIService) GetAppOutageExecute(r ApiGetAppOutageRequest) (*ApiApplicationOutageDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarReturnValue  *ApiApplicationOutageDetails
	)

	localBasePath := a.Client.GetConfig().ServerURL

	localVarPath := localBasePath + "/internet-insights/outages/app/{outageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"outageId"+"}", url.PathEscape(request.ParameterValueToString(r.outageId, "outageId")), -1)

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

type ApiGetNetworkOutageRequest struct {

	ApiService *InternetInsightsOutagesAPIService
	outageId string
	aid *string
}

// A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response.
func (r ApiGetNetworkOutageRequest) Aid(aid string) ApiGetNetworkOutageRequest {
	r.aid = &aid
	return r
}

func (r ApiGetNetworkOutageRequest) Execute() (*ApiNetworkOutageDetails, *http.Response, error) {
	return r.ApiService.GetNetworkOutageExecute(r)
}

/*
GetNetworkOutage Retrieve network outage

Returns the details of a network outage.


 @param outageId
 @return ApiGetNetworkOutageRequest
*/
func (a *InternetInsightsOutagesAPIService) GetNetworkOutage(outageId string ) ApiGetNetworkOutageRequest {
	return ApiGetNetworkOutageRequest{
		ApiService: a,
		outageId: outageId,
	}
}

// Execute executes the request
//  @return ApiNetworkOutageDetails
func (a *InternetInsightsOutagesAPIService) GetNetworkOutageExecute(r ApiGetNetworkOutageRequest) (*ApiNetworkOutageDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarReturnValue  *ApiNetworkOutageDetails
	)

	localBasePath := a.Client.GetConfig().ServerURL

	localVarPath := localBasePath + "/internet-insights/outages/net/{outageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"outageId"+"}", url.PathEscape(request.ParameterValueToString(r.outageId, "outageId")), -1)

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

func (a *InternetInsightsOutagesAPIService) decodeError(v interface{}, localVarBody []byte, localVarHTTPResponse *http.Response, newErr *internalerror.GenericAPIError) {
    err := a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
    if err != nil {
        newErr.ErrorMessage = err.Error()
        return
    }
    newErr.ErrorMessage = internalerror.FormatErrorMessage(localVarHTTPResponse.Status, &v)
    newErr.Model = v
}
