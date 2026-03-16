# \NetworksAPI

All URIs are relative to *https://api.sse.cisco.com/deployments/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetwork**](NetworksAPI.md#CreateNetwork) | **Post** /networks | Create Network
[**DeleteNetwork**](NetworksAPI.md#DeleteNetwork) | **Delete** /networks/{networkId} | Delete Network
[**GetNetwork**](NetworksAPI.md#GetNetwork) | **Get** /networks/{networkId} | Get Network
[**ListNetworks**](NetworksAPI.md#ListNetworks) | **Get** /networks | List Networks
[**UpdateNetwork**](NetworksAPI.md#UpdateNetwork) | **Put** /networks/{networkId} | Update Network



## CreateNetwork

> NetworkObject CreateNetwork(ctx).CreateNetworkRequest(createNetworkRequest).Execute()

Create Network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/networks"
)

func main() {
	createNetworkRequest := *openapiclient.NewCreateNetworkRequest("Name_example", int64(30), false, "Status_example") // CreateNetworkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetwork(context.Background()).CreateNetworkRequest(createNetworkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetwork`: NetworkObject
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetwork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworkRequest** | [**CreateNetworkRequest**](CreateNetworkRequest.md) |  | 

### Return type

[**NetworkObject**](NetworkObject.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetwork

> string DeleteNetwork(ctx, networkId).Execute()

Delete Network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/networks"
)

func main() {
	networkId := int64(13456) // int64 | The ID of the network.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.DeleteNetwork(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetwork`: string
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **int64** | The ID of the network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetwork

> NetworkObject GetNetwork(ctx, networkId).Execute()

Get Network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/networks"
)

func main() {
	networkId := int64(13456) // int64 | The ID of the network.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetwork(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetwork`: NetworkObject
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **int64** | The ID of the network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkObject**](NetworkObject.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworks

> []NetworkObject ListNetworks(ctx).Page(page).Limit(limit).Execute()

List Networks



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/networks"
)

func main() {
	page := int64(56) // int64 | The number of a page in the collection. (optional) (default to 1)
	limit := int64(56) // int64 | The number of records to return from the collection on the page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.ListNetworks(context.Background()).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.ListNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNetworks`: []NetworkObject
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.ListNetworks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | The number of a page in the collection. | [default to 1]
 **limit** | **int64** | The number of records to return from the collection on the page. | [default to 100]

### Return type

[**[]NetworkObject**](NetworkObject.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetwork

> NetworkObject UpdateNetwork(ctx, networkId).UpdateNetworkRequest(updateNetworkRequest).Execute()

Update Network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/networks"
)

func main() {
	networkId := int64(13456) // int64 | The ID of the network.
	updateNetworkRequest := *openapiclient.NewUpdateNetworkRequest("Name_example", false, "Status_example") // UpdateNetworkRequest | Update a network.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetwork(context.Background(), networkId).UpdateNetworkRequest(updateNetworkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetwork`: NetworkObject
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **int64** | The ID of the network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkRequest** | [**UpdateNetworkRequest**](UpdateNetworkRequest.md) | Update a network. | 

### Return type

[**NetworkObject**](NetworkObject.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

