# \APIUsageReportAPI

All URIs are relative to *https://api.sse.cisco.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAPIUsageKeys**](APIUsageReportAPI.md#GetAPIUsageKeys) | **Get** /reports/v2/apiUsage/keys | Get Keys
[**GetAPIUsageRequests**](APIUsageReportAPI.md#GetAPIUsageRequests) | **Get** /reports/v2/apiUsage/requests | Get Requests
[**GetAPIUsageResponses**](APIUsageReportAPI.md#GetAPIUsageResponses) | **Get** /reports/v2/apiUsage/responses | Get Responses
[**GetAPIUsageSummary**](APIUsageReportAPI.md#GetAPIUsageSummary) | **Get** /reports/v2/apiUsage/summary | Get Summary



## GetAPIUsageKeys

> Keys GetAPIUsageKeys(ctx).From(from).To(to).ApiKeys(apiKeys).Paths(paths).Verbs(verbs).StatusCodes(statusCodes).UserAgents(userAgents).Execute()

Get Keys



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/reports"
)

func main() {
	from := "2023-10-01" // string | The date and time specified in the RFC-3339 format, for example: 2023-10-01. Filter the collection for data that appears after this time.
	to := "2023-10-01" // string | The date and time specified in the RFC-3339 format, for example: 2023-10-01. Filter the collection for data that appears before this time.
	apiKeys := "123abfs77632" // string | Sort the collection using a list of comma-separated API key IDs. (optional)
	paths := "policies/v2/destinationlists/1234556,policies/v2/destinationlists" // string | Sort the collection using a list of comma-separated API resource paths. (optional)
	verbs := "GET,POST,PUT" // string | Sort the collection using a list of comma-separated HTTP verbs. The HTTP verb strings are case sensitive. (optional)
	statusCodes := "200,204,429" // string | Sort the collection using a list of comma-separated HTTP status codes. (optional)
	userAgents := "python-requests/2.20.0,Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36" // string | Sort the collection using a list of comma-separated labels that describe the HTTP client programs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIUsageReportAPI.GetAPIUsageKeys(context.Background()).From(from).To(to).ApiKeys(apiKeys).Paths(paths).Verbs(verbs).StatusCodes(statusCodes).UserAgents(userAgents).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIUsageReportAPI.GetAPIUsageKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAPIUsageKeys`: Keys
	fmt.Fprintf(os.Stdout, "Response from `APIUsageReportAPI.GetAPIUsageKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAPIUsageKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The date and time specified in the RFC-3339 format, for example: 2023-10-01. Filter the collection for data that appears after this time. | 
 **to** | **string** | The date and time specified in the RFC-3339 format, for example: 2023-10-01. Filter the collection for data that appears before this time. | 
 **apiKeys** | **string** | Sort the collection using a list of comma-separated API key IDs. | 
 **paths** | **string** | Sort the collection using a list of comma-separated API resource paths. | 
 **verbs** | **string** | Sort the collection using a list of comma-separated HTTP verbs. The HTTP verb strings are case sensitive. | 
 **statusCodes** | **string** | Sort the collection using a list of comma-separated HTTP status codes. | 
 **userAgents** | **string** | Sort the collection using a list of comma-separated labels that describe the HTTP client programs. | 

### Return type

[**Keys**](Keys.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAPIUsageRequests

> Requests GetAPIUsageRequests(ctx).From(from).To(to).ApiKeys(apiKeys).Paths(paths).Verbs(verbs).StatusCodes(statusCodes).UserAgents(userAgents).Execute()

Get Requests



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/reports"
)

func main() {
	from := "2023-10-01" // string | The date and time specified in the RFC-3339 format, for example: 2023-10-01. Filter the collection for data that appears after this time.
	to := "2023-10-01" // string | The date and time specified in the RFC-3339 format, for example: 2023-10-01. Filter the collection for data that appears before this time.
	apiKeys := "123abfs77632" // string | Sort the collection using a list of comma-separated API key IDs. (optional)
	paths := "policies/v2/destinationlists/1234556,policies/v2/destinationlists" // string | Sort the collection using a list of comma-separated API resource paths. (optional)
	verbs := "GET,POST,PUT" // string | Sort the collection using a list of comma-separated HTTP verbs. The HTTP verb strings are case sensitive. (optional)
	statusCodes := "200,204,429" // string | Sort the collection using a list of comma-separated HTTP status codes. (optional)
	userAgents := "python-requests/2.20.0,Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36" // string | Sort the collection using a list of comma-separated labels that describe the HTTP client programs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIUsageReportAPI.GetAPIUsageRequests(context.Background()).From(from).To(to).ApiKeys(apiKeys).Paths(paths).Verbs(verbs).StatusCodes(statusCodes).UserAgents(userAgents).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIUsageReportAPI.GetAPIUsageRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAPIUsageRequests`: Requests
	fmt.Fprintf(os.Stdout, "Response from `APIUsageReportAPI.GetAPIUsageRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAPIUsageRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The date and time specified in the RFC-3339 format, for example: 2023-10-01. Filter the collection for data that appears after this time. | 
 **to** | **string** | The date and time specified in the RFC-3339 format, for example: 2023-10-01. Filter the collection for data that appears before this time. | 
 **apiKeys** | **string** | Sort the collection using a list of comma-separated API key IDs. | 
 **paths** | **string** | Sort the collection using a list of comma-separated API resource paths. | 
 **verbs** | **string** | Sort the collection using a list of comma-separated HTTP verbs. The HTTP verb strings are case sensitive. | 
 **statusCodes** | **string** | Sort the collection using a list of comma-separated HTTP status codes. | 
 **userAgents** | **string** | Sort the collection using a list of comma-separated labels that describe the HTTP client programs. | 

### Return type

[**Requests**](Requests.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAPIUsageResponses

> Responses GetAPIUsageResponses(ctx).From(from).To(to).ApiKeys(apiKeys).Paths(paths).Verbs(verbs).StatusCodes(statusCodes).UserAgents(userAgents).Execute()

Get Responses



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/reports"
)

func main() {
	from := "2023-10-01" // string | The date and time specified in the RFC-3339 format, for example: 2023-10-01. Filter the collection for data that appears after this time.
	to := "2023-10-01" // string | The date and time specified in the RFC-3339 format, for example: 2023-10-01. Filter the collection for data that appears before this time.
	apiKeys := "123abfs77632" // string | Sort the collection using a list of comma-separated API key IDs. (optional)
	paths := "policies/v2/destinationlists/1234556,policies/v2/destinationlists" // string | Sort the collection using a list of comma-separated API resource paths. (optional)
	verbs := "GET,POST,PUT" // string | Sort the collection using a list of comma-separated HTTP verbs. The HTTP verb strings are case sensitive. (optional)
	statusCodes := "200,204,429" // string | Sort the collection using a list of comma-separated HTTP status codes. (optional)
	userAgents := "python-requests/2.20.0,Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36" // string | Sort the collection using a list of comma-separated labels that describe the HTTP client programs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIUsageReportAPI.GetAPIUsageResponses(context.Background()).From(from).To(to).ApiKeys(apiKeys).Paths(paths).Verbs(verbs).StatusCodes(statusCodes).UserAgents(userAgents).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIUsageReportAPI.GetAPIUsageResponses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAPIUsageResponses`: Responses
	fmt.Fprintf(os.Stdout, "Response from `APIUsageReportAPI.GetAPIUsageResponses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAPIUsageResponsesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The date and time specified in the RFC-3339 format, for example: 2023-10-01. Filter the collection for data that appears after this time. | 
 **to** | **string** | The date and time specified in the RFC-3339 format, for example: 2023-10-01. Filter the collection for data that appears before this time. | 
 **apiKeys** | **string** | Sort the collection using a list of comma-separated API key IDs. | 
 **paths** | **string** | Sort the collection using a list of comma-separated API resource paths. | 
 **verbs** | **string** | Sort the collection using a list of comma-separated HTTP verbs. The HTTP verb strings are case sensitive. | 
 **statusCodes** | **string** | Sort the collection using a list of comma-separated HTTP status codes. | 
 **userAgents** | **string** | Sort the collection using a list of comma-separated labels that describe the HTTP client programs. | 

### Return type

[**Responses**](Responses.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAPIUsageSummary

> Summary GetAPIUsageSummary(ctx).From(from).To(to).ApiKeys(apiKeys).Paths(paths).Verbs(verbs).StatusCodes(statusCodes).UserAgents(userAgents).Execute()

Get Summary



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/reports"
)

func main() {
	from := "2023-10-01" // string | The date and time specified in the RFC-3339 format, for example: 2023-10-01. Filter the collection for data that appears after this time.
	to := "2023-10-01" // string | The date and time specified in the RFC-3339 format, for example: 2023-10-01. Filter the collection for data that appears before this time.
	apiKeys := "123abfs77632" // string | Sort the collection using a list of comma-separated API key IDs. (optional)
	paths := "policies/v2/destinationlists/1234556,policies/v2/destinationlists" // string | Sort the collection using a list of comma-separated API resource paths. (optional)
	verbs := "GET,POST,PUT" // string | Sort the collection using a list of comma-separated HTTP verbs. The HTTP verb strings are case sensitive. (optional)
	statusCodes := "200,204,429" // string | Sort the collection using a list of comma-separated HTTP status codes. (optional)
	userAgents := "python-requests/2.20.0,Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36" // string | Sort the collection using a list of comma-separated labels that describe the HTTP client programs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIUsageReportAPI.GetAPIUsageSummary(context.Background()).From(from).To(to).ApiKeys(apiKeys).Paths(paths).Verbs(verbs).StatusCodes(statusCodes).UserAgents(userAgents).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIUsageReportAPI.GetAPIUsageSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAPIUsageSummary`: Summary
	fmt.Fprintf(os.Stdout, "Response from `APIUsageReportAPI.GetAPIUsageSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAPIUsageSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The date and time specified in the RFC-3339 format, for example: 2023-10-01. Filter the collection for data that appears after this time. | 
 **to** | **string** | The date and time specified in the RFC-3339 format, for example: 2023-10-01. Filter the collection for data that appears before this time. | 
 **apiKeys** | **string** | Sort the collection using a list of comma-separated API key IDs. | 
 **paths** | **string** | Sort the collection using a list of comma-separated API resource paths. | 
 **verbs** | **string** | Sort the collection using a list of comma-separated HTTP verbs. The HTTP verb strings are case sensitive. | 
 **statusCodes** | **string** | Sort the collection using a list of comma-separated HTTP status codes. | 
 **userAgents** | **string** | Sort the collection using a list of comma-separated labels that describe the HTTP client programs. | 

### Return type

[**Summary**](Summary.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

