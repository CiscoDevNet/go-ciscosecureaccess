# \NetworkTunnelGroupsAPI

All URIs are relative to *https://api.sse.cisco.com/deployments/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNetworkTunnelGroup**](NetworkTunnelGroupsAPI.md#AddNetworkTunnelGroup) | **Post** /networktunnelgroups | Create Network Tunnel Group
[**DeleteNetworkTunnelGroup**](NetworkTunnelGroupsAPI.md#DeleteNetworkTunnelGroup) | **Delete** /networktunnelgroups/{id} | Delete Network Tunnel Group
[**GetNetworkTunnelGroup**](NetworkTunnelGroupsAPI.md#GetNetworkTunnelGroup) | **Get** /networktunnelgroups/{id} | Get Network Tunnel Group
[**ListNetworkTunnelGroups**](NetworkTunnelGroupsAPI.md#ListNetworkTunnelGroups) | **Get** /networktunnelgroups | List Network Tunnel Groups
[**PatchNetworkTunnelGroup**](NetworkTunnelGroupsAPI.md#PatchNetworkTunnelGroup) | **Patch** /networktunnelgroups/{id} | Update Network Tunnel Group



## AddNetworkTunnelGroup

> NetworkTunnelGroupResponse AddNetworkTunnelGroup(ctx).AddNetworkTunnelGroupRequest(addNetworkTunnelGroupRequest).Execute()

Create Network Tunnel Group



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
	addNetworkTunnelGroupRequest := *openapiclient.NewAddNetworkTunnelGroupRequest("New York Branch Tunnels", "us-east-1", openapiclient.addNetworkTunnelGroup_request_authIdPrefix{ArrayOfString: new([]string)}, "t3stingTunn3lNow") // AddNetworkTunnelGroupRequest | Create the Network Tunnel Group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkTunnelGroupsAPI.AddNetworkTunnelGroup(context.Background()).AddNetworkTunnelGroupRequest(addNetworkTunnelGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkTunnelGroupsAPI.AddNetworkTunnelGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddNetworkTunnelGroup`: NetworkTunnelGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkTunnelGroupsAPI.AddNetworkTunnelGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddNetworkTunnelGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addNetworkTunnelGroupRequest** | [**AddNetworkTunnelGroupRequest**](AddNetworkTunnelGroupRequest.md) | Create the Network Tunnel Group. | 

### Return type

[**NetworkTunnelGroupResponse**](NetworkTunnelGroupResponse.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkTunnelGroup

> DeleteNetworkTunnelGroup(ctx, id).Execute()

Delete Network Tunnel Group



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
	r, err := apiClient.NetworkTunnelGroupsAPI.DeleteNetworkTunnelGroup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkTunnelGroupsAPI.DeleteNetworkTunnelGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the Network Tunnel Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkTunnelGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkTunnelGroup

> NetworkTunnelGroupResponse GetNetworkTunnelGroup(ctx, id).Execute()

Get Network Tunnel Group



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
	resp, r, err := apiClient.NetworkTunnelGroupsAPI.GetNetworkTunnelGroup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkTunnelGroupsAPI.GetNetworkTunnelGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkTunnelGroup`: NetworkTunnelGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkTunnelGroupsAPI.GetNetworkTunnelGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the Network Tunnel Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkTunnelGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkTunnelGroupResponse**](NetworkTunnelGroupResponse.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkTunnelGroups

> NetworkTunnelGroupsList ListNetworkTunnelGroups(ctx).Filters(filters).Offset(offset).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).IncludeStatuses(includeStatuses).Execute()

List Network Tunnel Groups



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
	filters := *openapiclient.NewFiltersNTGsObject() // FiltersNTGsObject | Filter the network tunnel groups by one or more properties:  * **name** - The name of a network tunnel group. The value of `name` is a sequence of case-insensitive characters. * **exactName** - The sequence of case-insensitive characters that exactly match the name of the network tunnel group.   When `exactName` is included as a filter, the `name` filter is ignored. * **networkTunnelGroupIds** - The comma-separated list of network tunnel group IDs. * **exactAuthIdPrefix** - The case-sensitive value of the network tunnel hub auth ID prefix or the IP. * **region** - The region for the network tunnel group. The value of `region` is a sequence of case-insensitive characters. * **status** - The status of the network tunnel group. Valid values are \"connected\", \"disconnected\", and \"warning\". * **duplicateCIDRs** - List the network tunnel groups that have duplicate CIDRs.   Provide the CIDRs and optionally provide the regional scope and region properties.   If the regional scope is enabled, only duplicates in the same region are found.   You can not use the `duplicateCIDRs` filter with any other filter.  Specify the filters in the JSON format.  Example:  ``` {     \"name\": \"Branch 1 Network Tunnel Group\",     \"region\": \"us-east-1\" } ``` or  Example:  ``` {     \"duplicateCIDRs\":     {         \"cidrs\": \"10.0.0.0/8,10.01.0.0/16\",         \"regionalScope\": true,         \"region\": \"us-east\"     } } (optional)
	offset := int64(10) // int64 | An integer that represents the place to start reading in the collection. When the offset is `0`, the first page is returned from the collection. If the `limit` is 10, the `offset` for the next page is 10. The default value is 0. (optional) (default to 0)
	limit := int64(25) // int64 | An integer that represents the number of records to return in the response. The default value is 10. (optional) (default to 10)
	sortBy := "sortBy_example" // string | Specify the field that will be used to sort the items from the collection in the response. (optional) (default to "name")
	sortOrder := "asc" // string | Specify the sort order (ascending or descending) for the items in the response. (optional) (default to "asc")
	includeStatuses := true // bool | Specify whether to include the IPsec tunnel status field (`tunnelsStatus`) for each hub. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkTunnelGroupsAPI.ListNetworkTunnelGroups(context.Background()).Filters(filters).Offset(offset).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).IncludeStatuses(includeStatuses).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkTunnelGroupsAPI.ListNetworkTunnelGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNetworkTunnelGroups`: NetworkTunnelGroupsList
	fmt.Fprintf(os.Stdout, "Response from `NetworkTunnelGroupsAPI.ListNetworkTunnelGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkTunnelGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | [**FiltersNTGsObject**](FiltersNTGsObject.md) | Filter the network tunnel groups by one or more properties:  * **name** - The name of a network tunnel group. The value of &#x60;name&#x60; is a sequence of case-insensitive characters. * **exactName** - The sequence of case-insensitive characters that exactly match the name of the network tunnel group.   When &#x60;exactName&#x60; is included as a filter, the &#x60;name&#x60; filter is ignored. * **networkTunnelGroupIds** - The comma-separated list of network tunnel group IDs. * **exactAuthIdPrefix** - The case-sensitive value of the network tunnel hub auth ID prefix or the IP. * **region** - The region for the network tunnel group. The value of &#x60;region&#x60; is a sequence of case-insensitive characters. * **status** - The status of the network tunnel group. Valid values are \&quot;connected\&quot;, \&quot;disconnected\&quot;, and \&quot;warning\&quot;. * **duplicateCIDRs** - List the network tunnel groups that have duplicate CIDRs.   Provide the CIDRs and optionally provide the regional scope and region properties.   If the regional scope is enabled, only duplicates in the same region are found.   You can not use the &#x60;duplicateCIDRs&#x60; filter with any other filter.  Specify the filters in the JSON format.  Example:  &#x60;&#x60;&#x60; {     \&quot;name\&quot;: \&quot;Branch 1 Network Tunnel Group\&quot;,     \&quot;region\&quot;: \&quot;us-east-1\&quot; } &#x60;&#x60;&#x60; or  Example:  &#x60;&#x60;&#x60; {     \&quot;duplicateCIDRs\&quot;:     {         \&quot;cidrs\&quot;: \&quot;10.0.0.0/8,10.01.0.0/16\&quot;,         \&quot;regionalScope\&quot;: true,         \&quot;region\&quot;: \&quot;us-east\&quot;     } } | 
 **offset** | **int64** | An integer that represents the place to start reading in the collection. When the offset is &#x60;0&#x60;, the first page is returned from the collection. If the &#x60;limit&#x60; is 10, the &#x60;offset&#x60; for the next page is 10. The default value is 0. | [default to 0]
 **limit** | **int64** | An integer that represents the number of records to return in the response. The default value is 10. | [default to 10]
 **sortBy** | **string** | Specify the field that will be used to sort the items from the collection in the response. | [default to &quot;name&quot;]
 **sortOrder** | **string** | Specify the sort order (ascending or descending) for the items in the response. | [default to &quot;asc&quot;]
 **includeStatuses** | **bool** | Specify whether to include the IPsec tunnel status field (&#x60;tunnelsStatus&#x60;) for each hub. | [default to false]

### Return type

[**NetworkTunnelGroupsList**](NetworkTunnelGroupsList.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchNetworkTunnelGroup

> NetworkTunnelGroupResponse PatchNetworkTunnelGroup(ctx, id).PatchNetworkTunnelGroupRequestInner(patchNetworkTunnelGroupRequestInner).Execute()

Update Network Tunnel Group



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
	patchNetworkTunnelGroupRequestInner := []openapiclient.PatchNetworkTunnelGroupRequestInner{*openapiclient.NewPatchNetworkTunnelGroupRequestInner("replace", "/name", openapiclient.patchNetworkTunnelGroup_request_inner_value{RoutingRequest: openapiclient.NewRoutingRequest("Type_example", openapiclient.routingRequest_data{BgpDataRequestObj: openapiclient.NewBgpDataRequestObj("5432")})})} // []PatchNetworkTunnelGroupRequestInner | Update the properties of the Network Tunnel Group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkTunnelGroupsAPI.PatchNetworkTunnelGroup(context.Background(), id).PatchNetworkTunnelGroupRequestInner(patchNetworkTunnelGroupRequestInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkTunnelGroupsAPI.PatchNetworkTunnelGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchNetworkTunnelGroup`: NetworkTunnelGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkTunnelGroupsAPI.PatchNetworkTunnelGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the Network Tunnel Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchNetworkTunnelGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchNetworkTunnelGroupRequestInner** | [**[]PatchNetworkTunnelGroupRequestInner**](PatchNetworkTunnelGroupRequestInner.md) | Update the properties of the Network Tunnel Group. | 

### Return type

[**NetworkTunnelGroupResponse**](NetworkTunnelGroupResponse.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

