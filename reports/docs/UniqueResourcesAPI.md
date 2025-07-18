# \UniqueResourcesAPI

All URIs are relative to *https://api.sse.cisco.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUniqueResources**](UniqueResourcesAPI.md#GetUniqueResources) | **Get** /reports/v2/unique-resources | Get Unique Resources (All)



## GetUniqueResources

> GetUniqueResources200Response GetUniqueResources(ctx).From(from).To(to).Exists(exists).Execute()

Get Unique Resources (All)



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
	exists := "destinationlistids,threattypes" // string | Specify an attribute or comma-separated list of attributes to filter the data. Valid values are: `categories`, `policycategories`, `applicationid`, `nbarapplicationid`, `nbarapplicationtypeids`, `privateapplicationid`, `applicationgroupids`, `sha256`, `filename`, `threats`, `threattypes`, `antivirusthreats`, `destinationlistids`, and `httperrors`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniqueResourcesAPI.GetUniqueResources(context.Background()).From(from).To(to).Exists(exists).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniqueResourcesAPI.GetUniqueResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniqueResources`: GetUniqueResources200Response
	fmt.Fprintf(os.Stdout, "Response from `UniqueResourcesAPI.GetUniqueResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUniqueResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | A timestamp or relative time string (for example: &#39;-1days&#39;). Filter for data that appears after this time. | 
 **to** | **string** | A timestamp or relative time string (for example: &#39;now&#39;). Filter for data that appears before this time. | 
 **exists** | **string** | Specify an attribute or comma-separated list of attributes to filter the data. Valid values are: &#x60;categories&#x60;, &#x60;policycategories&#x60;, &#x60;applicationid&#x60;, &#x60;nbarapplicationid&#x60;, &#x60;nbarapplicationtypeids&#x60;, &#x60;privateapplicationid&#x60;, &#x60;applicationgroupids&#x60;, &#x60;sha256&#x60;, &#x60;filename&#x60;, &#x60;threats&#x60;, &#x60;threattypes&#x60;, &#x60;antivirusthreats&#x60;, &#x60;destinationlistids&#x60;, and &#x60;httperrors&#x60;. | 

### Return type

[**GetUniqueResources200Response**](GetUniqueResources200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

