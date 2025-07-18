# \NetworkTunnelGroupsPeerStateAPI

All URIs are relative to *https://api.sse.cisco.com/deployments/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPeerState**](NetworkTunnelGroupsPeerStateAPI.md#GetPeerState) | **Get** /networktunnelgroups/{id}/networktunnelhubs/{hub_id}/peers/{peer_id}/state | Get Tunnel State for Network Tunnel Group and Hub
[**GetPeersState**](NetworkTunnelGroupsPeerStateAPI.md#GetPeersState) | **Get** /networktunnelgroups/{id}/peers | Get Peers States for Network Tunnel Group and Hub



## GetPeerState

> TunnelState GetPeerState(ctx, id, hubId, peerId).Execute()

Get Tunnel State for Network Tunnel Group and Hub



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
	hubId := int64(1234567) // int64 | The ID of the Network Tunnel Hub.
	peerId := int64(123455) // int64 | The ID of the peer (tunnel).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkTunnelGroupsPeerStateAPI.GetPeerState(context.Background(), id, hubId, peerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkTunnelGroupsPeerStateAPI.GetPeerState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPeerState`: TunnelState
	fmt.Fprintf(os.Stdout, "Response from `NetworkTunnelGroupsPeerStateAPI.GetPeerState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the Network Tunnel Group. | 
**hubId** | **int64** | The ID of the Network Tunnel Hub. | 
**peerId** | **int64** | The ID of the peer (tunnel). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPeerStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TunnelState**](TunnelState.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPeersState

> []TunnelState GetPeersState(ctx, id).Limit(limit).Offset(offset).SortBy(sortBy).SortOrder(sortOrder).Execute()

Get Peers States for Network Tunnel Group and Hub



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
	limit := int64(25) // int64 | An integer that represents the number of records to return in the response. The default value is 10. (optional) (default to 10)
	offset := int64(10) // int64 | An integer that represents the place to start reading in the collection. When the offset is `0`, the first page is returned from the collection. If the `limit` is 10, the `offset` for the next page is 10. The default value is 0. (optional) (default to 0)
	sortBy := "sortBy_example" // string | Specify the field that will be used to sort the items from the collection in the response. (optional) (default to "modifiedAt")
	sortOrder := "asc" // string | Specify the sort order (ascending or descending) for the items in the response. (optional) (default to "asc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkTunnelGroupsPeerStateAPI.GetPeersState(context.Background(), id).Limit(limit).Offset(offset).SortBy(sortBy).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkTunnelGroupsPeerStateAPI.GetPeersState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPeersState`: []TunnelState
	fmt.Fprintf(os.Stdout, "Response from `NetworkTunnelGroupsPeerStateAPI.GetPeersState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the Network Tunnel Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPeersStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | An integer that represents the number of records to return in the response. The default value is 10. | [default to 10]
 **offset** | **int64** | An integer that represents the place to start reading in the collection. When the offset is &#x60;0&#x60;, the first page is returned from the collection. If the &#x60;limit&#x60; is 10, the &#x60;offset&#x60; for the next page is 10. The default value is 0. | [default to 0]
 **sortBy** | **string** | Specify the field that will be used to sort the items from the collection in the response. | [default to &quot;modifiedAt&quot;]
 **sortOrder** | **string** | Specify the sort order (ascending or descending) for the items in the response. | [default to &quot;asc&quot;]

### Return type

[**[]TunnelState**](TunnelState.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

