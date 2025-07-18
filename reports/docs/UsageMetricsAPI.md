# \UsageMetricsAPI

All URIs are relative to *https://api.sse.cisco.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsageMetrics**](UsageMetricsAPI.md#GetUsageMetrics) | **Get** /reports/v2/usage/metrics | List Usage Metrics



## GetUsageMetrics

> GetUsageMetrics200Response GetUsageMetrics(ctx).OrganizationId(organizationId).ProductId(productId).From(from).To(to).TimeGroupingInterval(timeGroupingInterval).Page(page).Execute()

List Usage Metrics



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
	organizationId := int64(1234567) // int64 | The unique Secure Access organization ID.
	productId := "0" // string | The unique source ID of the usage report, for example: SWG, RAVPN, CHNE. Use `0` to search for all products.
	from := "2023-10-01" // string | The date and time specified in the RFC-3339 format, for example: 2023-10-01. Filter the collection for data that appears after this time.
	to := "2023-10-01" // string | The date and time specified in the RFC-3339 format, for example: 2023-10-01. Filter the collection for data that appears before this time.
	timeGroupingInterval := "DAY" // string | A time period specified as a segment of time. The usage report is collected within the time internal. Valid values are: `HOUR`, `DAY`, `WEEK`, `MONTH`.
	page := "2" // string | The page number that represents where to start reading in the collection. A page in the collection has a maximum of 1000 records. When a response has more than 1000 records, the `next` field is returned in the response. The `next` field is set to the API request URL for the next page of records in the collection. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageMetricsAPI.GetUsageMetrics(context.Background()).OrganizationId(organizationId).ProductId(productId).From(from).To(to).TimeGroupingInterval(timeGroupingInterval).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageMetricsAPI.GetUsageMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsageMetrics`: GetUsageMetrics200Response
	fmt.Fprintf(os.Stdout, "Response from `UsageMetricsAPI.GetUsageMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int64** | The unique Secure Access organization ID. | 
 **productId** | **string** | The unique source ID of the usage report, for example: SWG, RAVPN, CHNE. Use &#x60;0&#x60; to search for all products. | 
 **from** | **string** | The date and time specified in the RFC-3339 format, for example: 2023-10-01. Filter the collection for data that appears after this time. | 
 **to** | **string** | The date and time specified in the RFC-3339 format, for example: 2023-10-01. Filter the collection for data that appears before this time. | 
 **timeGroupingInterval** | **string** | A time period specified as a segment of time. The usage report is collected within the time internal. Valid values are: &#x60;HOUR&#x60;, &#x60;DAY&#x60;, &#x60;WEEK&#x60;, &#x60;MONTH&#x60;. | 
 **page** | **string** | The page number that represents where to start reading in the collection. A page in the collection has a maximum of 1000 records. When a response has more than 1000 records, the &#x60;next&#x60; field is returned in the response. The &#x60;next&#x60; field is set to the API request URL for the next page of records in the collection. | 

### Return type

[**GetUsageMetrics200Response**](GetUsageMetrics200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

