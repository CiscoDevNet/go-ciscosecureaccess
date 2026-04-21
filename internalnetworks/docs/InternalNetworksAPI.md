# \InternalNetworksAPI

All URIs are relative to *https://api.sse.cisco.com/deployments/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInternalNetwork**](InternalNetworksAPI.md#CreateInternalNetwork) | **Post** /internalnetworks | Create Internal Network
[**DeleteInternalNetwork**](InternalNetworksAPI.md#DeleteInternalNetwork) | **Delete** /internalnetworks/{internalNetworkId} | Delete Internal Network
[**GetInternalNetwork**](InternalNetworksAPI.md#GetInternalNetwork) | **Get** /internalnetworks/{internalNetworkId} | Get Internal Network
[**ListInternalNetworks**](InternalNetworksAPI.md#ListInternalNetworks) | **Get** /internalnetworks | List Internal Networks
[**UpdateInternalNetwork**](InternalNetworksAPI.md#UpdateInternalNetwork) | **Put** /internalnetworks/{internalNetworkId} | Update Internal Network



## CreateInternalNetwork

> InternalNetworkObject CreateInternalNetwork(ctx).CreateInternalNetworkRequest(createInternalNetworkRequest).Execute()

Create Internal Network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/internalnetworks"
)

func main() {
	createInternalNetworkRequest := *openapiclient.NewCreateInternalNetworkRequest("site one internal network", "198.2.2.4", int64(31)) // CreateInternalNetworkRequest | A JSON object that defines the internal network.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternalNetworksAPI.CreateInternalNetwork(context.Background()).CreateInternalNetworkRequest(createInternalNetworkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalNetworksAPI.CreateInternalNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInternalNetwork`: InternalNetworkObject
	fmt.Fprintf(os.Stdout, "Response from `InternalNetworksAPI.CreateInternalNetwork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInternalNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInternalNetworkRequest** | [**CreateInternalNetworkRequest**](CreateInternalNetworkRequest.md) | A JSON object that defines the internal network. | 

### Return type

[**InternalNetworkObject**](InternalNetworkObject.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInternalNetwork

> DeleteInternalNetwork(ctx, internalNetworkId).Execute()

Delete Internal Network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/internalnetworks"
)

func main() {
	internalNetworkId := int64(56) // int64 | The origin ID (originId) of the internal network.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InternalNetworksAPI.DeleteInternalNetwork(context.Background(), internalNetworkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalNetworksAPI.DeleteInternalNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**internalNetworkId** | **int64** | The origin ID (originId) of the internal network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInternalNetworkRequest struct via the builder pattern


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


## GetInternalNetwork

> InternalNetworkObject GetInternalNetwork(ctx, internalNetworkId).Execute()

Get Internal Network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/internalnetworks"
)

func main() {
	internalNetworkId := int64(56) // int64 | The origin ID (originId) of the internal network

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternalNetworksAPI.GetInternalNetwork(context.Background(), internalNetworkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalNetworksAPI.GetInternalNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInternalNetwork`: InternalNetworkObject
	fmt.Fprintf(os.Stdout, "Response from `InternalNetworksAPI.GetInternalNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**internalNetworkId** | **int64** | The origin ID (originId) of the internal network | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInternalNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InternalNetworkObject**](InternalNetworkObject.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInternalNetworks

> []InternalNetworkObject ListInternalNetworks(ctx).Page(page).Limit(limit).Name(name).Execute()

List Internal Networks



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/internalnetworks"
)

func main() {
	page := int64(56) // int64 | The number of a page in the collection. (optional) (default to 1)
	limit := int64(56) // int64 | The number of records to return from the collection on the page. (optional) (default to 100)
	name := "name_example" // string | internal network label (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternalNetworksAPI.ListInternalNetworks(context.Background()).Page(page).Limit(limit).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalNetworksAPI.ListInternalNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInternalNetworks`: []InternalNetworkObject
	fmt.Fprintf(os.Stdout, "Response from `InternalNetworksAPI.ListInternalNetworks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInternalNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | The number of a page in the collection. | [default to 1]
 **limit** | **int64** | The number of records to return from the collection on the page. | [default to 100]
 **name** | **string** | internal network label | 

### Return type

[**[]InternalNetworkObject**](InternalNetworkObject.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInternalNetwork

> InternalNetworkObject UpdateInternalNetwork(ctx, internalNetworkId).CreateInternalNetworkRequest(createInternalNetworkRequest).Execute()

Update Internal Network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/internalnetworks"
)

func main() {
	internalNetworkId := int64(56) // int64 | The origin ID (originId) of the internal network.
	createInternalNetworkRequest := *openapiclient.NewCreateInternalNetworkRequest("site one internal network", "198.2.2.4", int64(31)) // CreateInternalNetworkRequest | A JSON object that defines the internal network.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternalNetworksAPI.UpdateInternalNetwork(context.Background(), internalNetworkId).CreateInternalNetworkRequest(createInternalNetworkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalNetworksAPI.UpdateInternalNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInternalNetwork`: InternalNetworkObject
	fmt.Fprintf(os.Stdout, "Response from `InternalNetworksAPI.UpdateInternalNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**internalNetworkId** | **int64** | The origin ID (originId) of the internal network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInternalNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createInternalNetworkRequest** | [**CreateInternalNetworkRequest**](CreateInternalNetworkRequest.md) | A JSON object that defines the internal network. | 

### Return type

[**InternalNetworkObject**](InternalNetworkObject.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

