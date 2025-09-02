# \DestinationsAPI

All URIs are relative to *https://api.sse.cisco.com/policies/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDestinations**](DestinationsAPI.md#CreateDestinations) | **Post** /destinationlists/{destinationListId}/destinations | Add Destinations to Destination List
[**DeleteDestinations**](DestinationsAPI.md#DeleteDestinations) | **Delete** /destinationlists/{destinationListId}/destinations/remove | Delete Destinations from Destination List
[**GetDestinations**](DestinationsAPI.md#GetDestinations) | **Get** /destinationlists/{destinationListId}/destinations | Get Destinations in Destination List



## CreateDestinations

> DestinationListResponse CreateDestinations(ctx, destinationListId).DestinationCreateObject(destinationCreateObject).Execute()

Add Destinations to Destination List



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
	destinationCreateObject := []openapiclient.DestinationCreateObject{*openapiclient.NewDestinationCreateObject("cisco.com")} // []DestinationCreateObject | Add destinations to a destination list. Accepts no more than 500 destination objects in the body of the request unless the destination list is of type `thirdparty_block`. If the type of the destination list is `thirdparty_block`, then the system accepts 100 destination objects.  If you make an API request on the POST operation that adds a URL on a high-volume domain to a destination list, the operation may succeed (`HTTP/200 OK`). However, the server returns an error message (`HTTP/400 Bad Request`) that indicates that the destination is on a high-volume domain. **Note:** Secure Access does not add URLs that are on high-volume domains to destination lists. Instead, we recommend that you add the domain only.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DestinationsAPI.CreateDestinations(context.Background(), destinationListId).DestinationCreateObject(destinationCreateObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.CreateDestinations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDestinations`: DestinationListResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.CreateDestinations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destinationListId** | **int64** | The unique ID of the destination list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDestinationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **destinationCreateObject** | [**[]DestinationCreateObject**](DestinationCreateObject.md) | Add destinations to a destination list. Accepts no more than 500 destination objects in the body of the request unless the destination list is of type &#x60;thirdparty_block&#x60;. If the type of the destination list is &#x60;thirdparty_block&#x60;, then the system accepts 100 destination objects.  If you make an API request on the POST operation that adds a URL on a high-volume domain to a destination list, the operation may succeed (&#x60;HTTP/200 OK&#x60;). However, the server returns an error message (&#x60;HTTP/400 Bad Request&#x60;) that indicates that the destination is on a high-volume domain. **Note:** Secure Access does not add URLs that are on high-volume domains to destination lists. Instead, we recommend that you add the domain only. | 

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


## DeleteDestinations

> DestinationListResponse DeleteDestinations(ctx, destinationListId).RequestBody(requestBody).Execute()

Delete Destinations from Destination List



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
	requestBody := []int64{int64(1234)} // []int64 | Add a list of destination ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DestinationsAPI.DeleteDestinations(context.Background(), destinationListId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.DeleteDestinations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDestinations`: DestinationListResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.DeleteDestinations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destinationListId** | **int64** | The unique ID of the destination list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDestinationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]int64** | Add a list of destination ID. | 

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


## GetDestinations

> PaginatedDestinationObjectResponse GetDestinations(ctx, destinationListId).Page(page).Limit(limit).Execute()

Get Destinations in Destination List



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
	page := int64(4) // int64 | The number of a page in the collection. (optional) (default to 1)
	limit := int64(50) // int64 | The number of records in the collection to return on the page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DestinationsAPI.GetDestinations(context.Background(), destinationListId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.GetDestinations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDestinations`: PaginatedDestinationObjectResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.GetDestinations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destinationListId** | **int64** | The unique ID of the destination list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDestinationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | The number of a page in the collection. | [default to 1]
 **limit** | **int64** | The number of records in the collection to return on the page. | [default to 100]

### Return type

[**PaginatedDestinationObjectResponse**](PaginatedDestinationObjectResponse.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

