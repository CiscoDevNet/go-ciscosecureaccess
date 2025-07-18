# \ApplicationCategoriesAPI

All URIs are relative to *https://api.sse.cisco.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApplicationCategories**](ApplicationCategoriesAPI.md#GetApplicationCategories) | **Get** /reports/v2/appDiscovery/applicationCategories | List Application Categories



## GetApplicationCategories

> ApplicationCategoryList GetApplicationCategories(ctx).Limit(limit).Offset(offset).Execute()

List Application Categories



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
	limit := int64(56) // int64 | The maximum number of items to return in the collection. (optional)
	offset := int64(56) // int64 | The number of items to skip before starting to collect the result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationCategoriesAPI.GetApplicationCategories(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCategoriesAPI.GetApplicationCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationCategories`: ApplicationCategoryList
	fmt.Fprintf(os.Stdout, "Response from `ApplicationCategoriesAPI.GetApplicationCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | The maximum number of items to return in the collection. | 
 **offset** | **int64** | The number of items to skip before starting to collect the result set. | 

### Return type

[**ApplicationCategoryList**](ApplicationCategoryList.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

