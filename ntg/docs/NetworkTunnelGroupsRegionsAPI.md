# \NetworkTunnelGroupsRegionsAPI

All URIs are relative to *https://api.sse.cisco.com/deployments/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListNetworkTunnelGroupRegions**](NetworkTunnelGroupsRegionsAPI.md#ListNetworkTunnelGroupRegions) | **Get** /regions | List Regions for Network Tunnel Groups



## ListNetworkTunnelGroupRegions

> RegionList ListNetworkTunnelGroupRegions(ctx).Filters(filters).Execute()

List Regions for Network Tunnel Groups



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
	filters := *openapiclient.NewFiltersRegionsObject() // FiltersRegionsObject | Filter the regions by one or more properties: `peerIP` or `latitude` and `longitude`.  * **peerIP** - List the regions in ascending order based on the distance of the regions to the location of the given peer IP.   You can only set a public IP for `peerIP`. If `latitude` and `longitude` are provided, `peerIP` is ignored. * **latitude** and **longitude** - List the regions in ascending order based on the distance of the regions from the provided coordinates.   When included with a request, set both `latitude` and `longitude`.  Specify the filters in the JSON format.  Examples:  ``` {   \"peerIP\": \"25.123.22.10\" } ```  ``` {   \"latitude\": \"39.0299604\",   \"longitude\": \"39.0299604\" } ``` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkTunnelGroupsRegionsAPI.ListNetworkTunnelGroupRegions(context.Background()).Filters(filters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkTunnelGroupsRegionsAPI.ListNetworkTunnelGroupRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNetworkTunnelGroupRegions`: RegionList
	fmt.Fprintf(os.Stdout, "Response from `NetworkTunnelGroupsRegionsAPI.ListNetworkTunnelGroupRegions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkTunnelGroupRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | [**FiltersRegionsObject**](FiltersRegionsObject.md) | Filter the regions by one or more properties: &#x60;peerIP&#x60; or &#x60;latitude&#x60; and &#x60;longitude&#x60;.  * **peerIP** - List the regions in ascending order based on the distance of the regions to the location of the given peer IP.   You can only set a public IP for &#x60;peerIP&#x60;. If &#x60;latitude&#x60; and &#x60;longitude&#x60; are provided, &#x60;peerIP&#x60; is ignored. * **latitude** and **longitude** - List the regions in ascending order based on the distance of the regions from the provided coordinates.   When included with a request, set both &#x60;latitude&#x60; and &#x60;longitude&#x60;.  Specify the filters in the JSON format.  Examples:  &#x60;&#x60;&#x60; {   \&quot;peerIP\&quot;: \&quot;25.123.22.10\&quot; } &#x60;&#x60;&#x60;  &#x60;&#x60;&#x60; {   \&quot;latitude\&quot;: \&quot;39.0299604\&quot;,   \&quot;longitude\&quot;: \&quot;39.0299604\&quot; } &#x60;&#x60;&#x60; | 

### Return type

[**RegionList**](RegionList.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

