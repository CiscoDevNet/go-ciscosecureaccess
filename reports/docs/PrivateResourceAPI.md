# \PrivateResourceAPI

All URIs are relative to *https://api.sse.cisco.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPrivateResourceDetailedStatsIdentities**](PrivateResourceAPI.md#GetPrivateResourceDetailedStatsIdentities) | **Get** /reports/v2/private-resources/detailed-stats-identities | Get Private Resource Access Statistics in Identity Report
[**GetPrivateResourceDetailedStatsTimerange**](PrivateResourceAPI.md#GetPrivateResourceDetailedStatsTimerange) | **Get** /reports/v2/private-resources/detailed-stats-timerange | Get Private Resource Access Statistics in Details Report
[**GetPrivateResourceStats**](PrivateResourceAPI.md#GetPrivateResourceStats) | **Get** /reports/v2/private-resources/summary-stats | Get Private Resource Access Statistics in Summary Report



## GetPrivateResourceDetailedStatsIdentities

> GetPrivateResourceDetailedStatsIdentities200Response GetPrivateResourceDetailedStatsIdentities(ctx).From(from).To(to).Limit(limit).Privateresourceid(privateresourceid).Offset(offset).Timezone(timezone).Execute()

Get Private Resource Access Statistics in Identity Report



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
	from := "1639146300000" // string | A timestamp or relative time string (for example: '-1days'). Filter for data that appears after this time.
	to := "1640010300000" // string | A timestamp or relative time string (for example: 'now'). Filter for data that appears before this time.
	limit := int64(100) // int64 | The maximum number of records to return from the collection. (default to 100)
	privateresourceid := int64(47) // int64 | A private resource ID.
	offset := int64(0) // int64 | A number that represents an index in the collection. (optional) (default to 0)
	timezone := "ASIA%2fCALCUTTA" // string | Display the timestamp of the traffic events in the specified timezone. For the timezone, provide a continent and city separated by an url-encoded forward slash ('/'), for example: timezone='ASIA%2fCALCUTTA'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivateResourceAPI.GetPrivateResourceDetailedStatsIdentities(context.Background()).From(from).To(to).Limit(limit).Privateresourceid(privateresourceid).Offset(offset).Timezone(timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivateResourceAPI.GetPrivateResourceDetailedStatsIdentities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrivateResourceDetailedStatsIdentities`: GetPrivateResourceDetailedStatsIdentities200Response
	fmt.Fprintf(os.Stdout, "Response from `PrivateResourceAPI.GetPrivateResourceDetailedStatsIdentities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateResourceDetailedStatsIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | A timestamp or relative time string (for example: &#39;-1days&#39;). Filter for data that appears after this time. | 
 **to** | **string** | A timestamp or relative time string (for example: &#39;now&#39;). Filter for data that appears before this time. | 
 **limit** | **int64** | The maximum number of records to return from the collection. | [default to 100]
 **privateresourceid** | **int64** | A private resource ID. | 
 **offset** | **int64** | A number that represents an index in the collection. | [default to 0]
 **timezone** | **string** | Display the timestamp of the traffic events in the specified timezone. For the timezone, provide a continent and city separated by an url-encoded forward slash (&#39;/&#39;), for example: timezone&#x3D;&#39;ASIA%2fCALCUTTA&#39;. | 

### Return type

[**GetPrivateResourceDetailedStatsIdentities200Response**](GetPrivateResourceDetailedStatsIdentities200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateResourceDetailedStatsTimerange

> GetPrivateResourceDetailedStatsTimerange200Response GetPrivateResourceDetailedStatsTimerange(ctx).From(from).To(to).Privateresourceid(privateresourceid).Timerange(timerange).Timezone(timezone).Execute()

Get Private Resource Access Statistics in Details Report



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
	from := "1639146300000" // string | A timestamp or relative time string (for example: '-1days'). Filter for data that appears after this time.
	to := "1640010300000" // string | A timestamp or relative time string (for example: 'now'). Filter for data that appears before this time.
	privateresourceid := int64(47) // int64 | A private resource ID.
	timerange := "minute" // string | A string that represents a range of time. If the header is not set, the header value defaults to `hour`. (optional)
	timezone := "ASIA%2fCALCUTTA" // string | Display the timestamp of the traffic events in the specified timezone. For the timezone, provide a continent and city separated by an url-encoded forward slash ('/'), for example: timezone='ASIA%2fCALCUTTA'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivateResourceAPI.GetPrivateResourceDetailedStatsTimerange(context.Background()).From(from).To(to).Privateresourceid(privateresourceid).Timerange(timerange).Timezone(timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivateResourceAPI.GetPrivateResourceDetailedStatsTimerange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrivateResourceDetailedStatsTimerange`: GetPrivateResourceDetailedStatsTimerange200Response
	fmt.Fprintf(os.Stdout, "Response from `PrivateResourceAPI.GetPrivateResourceDetailedStatsTimerange`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateResourceDetailedStatsTimerangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | A timestamp or relative time string (for example: &#39;-1days&#39;). Filter for data that appears after this time. | 
 **to** | **string** | A timestamp or relative time string (for example: &#39;now&#39;). Filter for data that appears before this time. | 
 **privateresourceid** | **int64** | A private resource ID. | 
 **timerange** | **string** | A string that represents a range of time. If the header is not set, the header value defaults to &#x60;hour&#x60;. | 
 **timezone** | **string** | Display the timestamp of the traffic events in the specified timezone. For the timezone, provide a continent and city separated by an url-encoded forward slash (&#39;/&#39;), for example: timezone&#x3D;&#39;ASIA%2fCALCUTTA&#39;. | 

### Return type

[**GetPrivateResourceDetailedStatsTimerange200Response**](GetPrivateResourceDetailedStatsTimerange200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateResourceStats

> GetPrivateResourceStats200Response GetPrivateResourceStats(ctx).From(from).To(to).Limit(limit).Privateresourceids(privateresourceids).Offset(offset).Execute()

Get Private Resource Access Statistics in Summary Report



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
	from := "1639146300000" // string | A timestamp or relative time string (for example: '-1days'). Filter for data that appears after this time.
	to := "1640010300000" // string | A timestamp or relative time string (for example: 'now'). Filter for data that appears before this time.
	limit := int64(100) // int64 | The maximum number of records to return from the collection. (default to 100)
	privateresourceids := "29,31" // string | A comma-delimited list of private resource IDs.
	offset := int64(0) // int64 | A number that represents an index in the collection. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivateResourceAPI.GetPrivateResourceStats(context.Background()).From(from).To(to).Limit(limit).Privateresourceids(privateresourceids).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivateResourceAPI.GetPrivateResourceStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrivateResourceStats`: GetPrivateResourceStats200Response
	fmt.Fprintf(os.Stdout, "Response from `PrivateResourceAPI.GetPrivateResourceStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateResourceStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | A timestamp or relative time string (for example: &#39;-1days&#39;). Filter for data that appears after this time. | 
 **to** | **string** | A timestamp or relative time string (for example: &#39;now&#39;). Filter for data that appears before this time. | 
 **limit** | **int64** | The maximum number of records to return from the collection. | [default to 100]
 **privateresourceids** | **string** | A comma-delimited list of private resource IDs. | 
 **offset** | **int64** | A number that represents an index in the collection. | [default to 0]

### Return type

[**GetPrivateResourceStats200Response**](GetPrivateResourceStats200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

