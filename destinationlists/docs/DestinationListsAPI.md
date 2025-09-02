# \DestinationListsAPI

All URIs are relative to *https://api.sse.cisco.com/policies/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDestinationList**](DestinationListsAPI.md#CreateDestinationList) | **Post** /destinationlists | Create Destination List
[**DeleteDestinationList**](DestinationListsAPI.md#DeleteDestinationList) | **Delete** /destinationlists/{destinationListId} | Delete Destination List
[**GetDestinationList**](DestinationListsAPI.md#GetDestinationList) | **Get** /destinationlists/{destinationListId} | Get Destination List
[**GetDestinationLists**](DestinationListsAPI.md#GetDestinationLists) | **Get** /destinationlists | Get Destination Lists
[**UpdateDestinationLists**](DestinationListsAPI.md#UpdateDestinationLists) | **Patch** /destinationlists/{destinationListId} | Update Destination List



## CreateDestinationList

> DestinationListResponse CreateDestinationList(ctx).DestinationListCreate(destinationListCreate).Execute()

Create Destination List



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cisco-sbg/go-ciscosecureaccess/destinationlists"
)

func main() {
	destinationListCreate := *openapiclient.NewDestinationListCreate(openapiclient.access("allow"), false, "Global Allow list") // DestinationListCreate | Provide destination information and an optional array of destination objects. * In the API request, accepts no more than 500 destination objects in the destination list. * Does not support global destination lists. You must set the `isGlobal` field to `false`. * You cannot use the Destination Lists API to create a destination list with the `access` type of `thirdparty_block`. * You must set the `bundleTypeId` to `2`. Web profiles are applied on your policy rules.  If you make an API request on the POST operation that adds a URL on a high-volume domain to a destination list, the operation may succeed (`HTTP/200 OK`). However, the server returns an error message (`HTTP/400 Bad Request`) that indicates that the destination is on a high-volume domain. **Note:** Secure Access does not add URLs that are on high-volume domains to destination lists. Instead, we recommend that you add the domain only.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DestinationListsAPI.CreateDestinationList(context.Background()).DestinationListCreate(destinationListCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationListsAPI.CreateDestinationList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDestinationList`: DestinationListResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationListsAPI.CreateDestinationList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDestinationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationListCreate** | [**DestinationListCreate**](DestinationListCreate.md) | Provide destination information and an optional array of destination objects. * In the API request, accepts no more than 500 destination objects in the destination list. * Does not support global destination lists. You must set the &#x60;isGlobal&#x60; field to &#x60;false&#x60;. * You cannot use the Destination Lists API to create a destination list with the &#x60;access&#x60; type of &#x60;thirdparty_block&#x60;. * You must set the &#x60;bundleTypeId&#x60; to &#x60;2&#x60;. Web profiles are applied on your policy rules.  If you make an API request on the POST operation that adds a URL on a high-volume domain to a destination list, the operation may succeed (&#x60;HTTP/200 OK&#x60;). However, the server returns an error message (&#x60;HTTP/400 Bad Request&#x60;) that indicates that the destination is on a high-volume domain. **Note:** Secure Access does not add URLs that are on high-volume domains to destination lists. Instead, we recommend that you add the domain only. | 

### Return type

[**DestinationListResponse**](DestinationListResponse.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDestinationList

> DestinationListDelete DeleteDestinationList(ctx, destinationListId).Execute()

Delete Destination List



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cisco-sbg/go-ciscosecureaccess/destinationlists"
)

func main() {
	destinationListId := int64(245) // int64 | The unique ID of the destination list.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DestinationListsAPI.DeleteDestinationList(context.Background(), destinationListId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationListsAPI.DeleteDestinationList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDestinationList`: DestinationListDelete
	fmt.Fprintf(os.Stdout, "Response from `DestinationListsAPI.DeleteDestinationList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destinationListId** | **int64** | The unique ID of the destination list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDestinationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DestinationListDelete**](DestinationListDelete.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDestinationList

> DestinationListResponse GetDestinationList(ctx, destinationListId).Execute()

Get Destination List



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cisco-sbg/go-ciscosecureaccess/destinationlists"
)

func main() {
	destinationListId := int64(245) // int64 | The unique ID of the destination list.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DestinationListsAPI.GetDestinationList(context.Background(), destinationListId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationListsAPI.GetDestinationList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDestinationList`: DestinationListResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationListsAPI.GetDestinationList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destinationListId** | **int64** | The unique ID of the destination list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDestinationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DestinationListResponse**](DestinationListResponse.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDestinationLists

> PaginatedDestinationListsResponse GetDestinationLists(ctx).Page(page).Limit(limit).Execute()

Get Destination Lists



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cisco-sbg/go-ciscosecureaccess/destinationlists"
)

func main() {
	page := int64(4) // int64 | The number of a page in the collection. (optional) (default to 1)
	limit := int64(50) // int64 | The number of records in the collection to return on the page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DestinationListsAPI.GetDestinationLists(context.Background()).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationListsAPI.GetDestinationLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDestinationLists`: PaginatedDestinationListsResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationListsAPI.GetDestinationLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDestinationListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | The number of a page in the collection. | [default to 1]
 **limit** | **int64** | The number of records in the collection to return on the page. | [default to 100]

### Return type

[**PaginatedDestinationListsResponse**](PaginatedDestinationListsResponse.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDestinationLists

> DestinationListResponse UpdateDestinationLists(ctx, destinationListId).DestinationListPatch(destinationListPatch).Execute()

Update Destination List



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cisco-sbg/go-ciscosecureaccess/destinationlists"
)

func main() {
	destinationListId := int64(245) // int64 | The unique ID of the destination list.
	destinationListPatch := *openapiclient.NewDestinationListPatch("Global Allow list") // DestinationListPatch | Update a destination list.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DestinationListsAPI.UpdateDestinationLists(context.Background(), destinationListId).DestinationListPatch(destinationListPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationListsAPI.UpdateDestinationLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDestinationLists`: DestinationListResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationListsAPI.UpdateDestinationLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destinationListId** | **int64** | The unique ID of the destination list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDestinationListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **destinationListPatch** | [**DestinationListPatch**](DestinationListPatch.md) | Update a destination list. | 

### Return type

[**DestinationListResponse**](DestinationListResponse.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

