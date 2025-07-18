# \NetworkTunnelGroupsStateAPI

All URIs are relative to *https://api.sse.cisco.com/deployments/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkTunnelGroupBulkState**](NetworkTunnelGroupsStateAPI.md#GetNetworkTunnelGroupBulkState) | **Get** /networktunnelgroupsstate | List State of Network Tunnel Groups
[**GetNetworkTunnelGroupState**](NetworkTunnelGroupsStateAPI.md#GetNetworkTunnelGroupState) | **Get** /networktunnelgroups/{id}/state | Get State of Network Tunnel Group



## GetNetworkTunnelGroupBulkState

> NetworkTunnelGroupBulkStateResponse GetNetworkTunnelGroupBulkState(ctx).Offset(offset).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Execute()

List State of Network Tunnel Groups



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/ntg"
)

func main() {
	offset := int64(10) // int64 | An integer that represents the place to start reading in the collection. When the offset is `0`, the first page is returned from the collection. If the `limit` is 10, the `offset` for the next page is 10. The default value is 0. (optional) (default to 0)
	limit := int64(56) // int64 | An integer that represents the number of records to return in the response. (optional) (default to 10)
	sortBy := "sortBy_example" // string | Specify the field that will be used to sort the items from the collection in the response. (optional) (default to "name")
	sortOrder := "asc" // string | Specify the sort order (ascending or descending) for the items in the response. (optional) (default to "asc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkTunnelGroupsStateAPI.GetNetworkTunnelGroupBulkState(context.Background()).Offset(offset).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkTunnelGroupsStateAPI.GetNetworkTunnelGroupBulkState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkTunnelGroupBulkState`: NetworkTunnelGroupBulkStateResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkTunnelGroupsStateAPI.GetNetworkTunnelGroupBulkState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkTunnelGroupBulkStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int64** | An integer that represents the place to start reading in the collection. When the offset is &#x60;0&#x60;, the first page is returned from the collection. If the &#x60;limit&#x60; is 10, the &#x60;offset&#x60; for the next page is 10. The default value is 0. | [default to 0]
 **limit** | **int64** | An integer that represents the number of records to return in the response. | [default to 10]
 **sortBy** | **string** | Specify the field that will be used to sort the items from the collection in the response. | [default to &quot;name&quot;]
 **sortOrder** | **string** | Specify the sort order (ascending or descending) for the items in the response. | [default to &quot;asc&quot;]

### Return type

[**NetworkTunnelGroupBulkStateResponse**](NetworkTunnelGroupBulkStateResponse.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkTunnelGroupState

> NetworkTunnelGroupStateResponse GetNetworkTunnelGroupState(ctx, id).Execute()

Get State of Network Tunnel Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/ntg"
)

func main() {
	id := int64(123455) // int64 | The ID of the Network Tunnel Group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkTunnelGroupsStateAPI.GetNetworkTunnelGroupState(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkTunnelGroupsStateAPI.GetNetworkTunnelGroupState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkTunnelGroupState`: NetworkTunnelGroupStateResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkTunnelGroupsStateAPI.GetNetworkTunnelGroupState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the Network Tunnel Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkTunnelGroupStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkTunnelGroupStateResponse**](NetworkTunnelGroupStateResponse.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

