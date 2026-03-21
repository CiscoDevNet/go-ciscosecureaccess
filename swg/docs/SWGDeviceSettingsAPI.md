# \SWGDeviceSettingsAPI

All URIs are relative to *https://api.sse.cisco.com/deployments/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecureWebGatewayDeviceSettings**](SWGDeviceSettingsAPI.md#CreateSecureWebGatewayDeviceSettings) | **Post** /deviceSettings/SWGEnabled/set | Set SWG Override Device Settings
[**DeleteSecureWebGatewayDeviceSettings**](SWGDeviceSettingsAPI.md#DeleteSecureWebGatewayDeviceSettings) | **Post** /deviceSettings/SWGEnabled/remove | Delete SWG Override Device Settings
[**ListSecureWebGatewayDeviceSettings**](SWGDeviceSettingsAPI.md#ListSecureWebGatewayDeviceSettings) | **Post** /deviceSettings/SWGEnabled/list | List SWG Override Device Settings



## CreateSecureWebGatewayDeviceSettings

> RegisteredSWGDeviceSettings CreateSecureWebGatewayDeviceSettings(ctx).CreateSecureWebGatewayDeviceSettingsRequest(createSecureWebGatewayDeviceSettingsRequest).Execute()

Set SWG Override Device Settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/swg"
)

func main() {
	createSecureWebGatewayDeviceSettingsRequest := *openapiclient.NewCreateSecureWebGatewayDeviceSettingsRequest([]int64{int64(12321231)}, openapiclient.value("0")) // CreateSecureWebGatewayDeviceSettingsRequest | * Provide a list of origin ID for the devices in the organization. The list can contain 1–100 origin IDs. * Provide the Secure Web gateway (SWG) device setting to apply to the devices. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SWGDeviceSettingsAPI.CreateSecureWebGatewayDeviceSettings(context.Background()).CreateSecureWebGatewayDeviceSettingsRequest(createSecureWebGatewayDeviceSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SWGDeviceSettingsAPI.CreateSecureWebGatewayDeviceSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSecureWebGatewayDeviceSettings`: RegisteredSWGDeviceSettings
	fmt.Fprintf(os.Stdout, "Response from `SWGDeviceSettingsAPI.CreateSecureWebGatewayDeviceSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecureWebGatewayDeviceSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSecureWebGatewayDeviceSettingsRequest** | [**CreateSecureWebGatewayDeviceSettingsRequest**](CreateSecureWebGatewayDeviceSettingsRequest.md) | * Provide a list of origin ID for the devices in the organization. The list can contain 1–100 origin IDs. * Provide the Secure Web gateway (SWG) device setting to apply to the devices. | 

### Return type

[**RegisteredSWGDeviceSettings**](RegisteredSWGDeviceSettings.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecureWebGatewayDeviceSettings

> DeleteSecureWebGatewayDeviceSettings200Response DeleteSecureWebGatewayDeviceSettings(ctx).ListSecureWebGatewayDeviceSettingsRequest(listSecureWebGatewayDeviceSettingsRequest).Execute()

Delete SWG Override Device Settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/swg"
)

func main() {
	listSecureWebGatewayDeviceSettingsRequest := *openapiclient.NewListSecureWebGatewayDeviceSettingsRequest([]int64{int64(12321231)}) // ListSecureWebGatewayDeviceSettingsRequest | Provide a list of origin ID for the devices in the organization. The list can contain 1–100 origin IDs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SWGDeviceSettingsAPI.DeleteSecureWebGatewayDeviceSettings(context.Background()).ListSecureWebGatewayDeviceSettingsRequest(listSecureWebGatewayDeviceSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SWGDeviceSettingsAPI.DeleteSecureWebGatewayDeviceSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSecureWebGatewayDeviceSettings`: DeleteSecureWebGatewayDeviceSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `SWGDeviceSettingsAPI.DeleteSecureWebGatewayDeviceSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecureWebGatewayDeviceSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listSecureWebGatewayDeviceSettingsRequest** | [**ListSecureWebGatewayDeviceSettingsRequest**](ListSecureWebGatewayDeviceSettingsRequest.md) | Provide a list of origin ID for the devices in the organization. The list can contain 1–100 origin IDs. | 

### Return type

[**DeleteSecureWebGatewayDeviceSettings200Response**](DeleteSecureWebGatewayDeviceSettings200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSecureWebGatewayDeviceSettings

> []ListSWGDeviceSettingsInner ListSecureWebGatewayDeviceSettings(ctx).ListSecureWebGatewayDeviceSettingsRequest(listSecureWebGatewayDeviceSettingsRequest).Execute()

List SWG Override Device Settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/swg"
)

func main() {
	listSecureWebGatewayDeviceSettingsRequest := *openapiclient.NewListSecureWebGatewayDeviceSettingsRequest([]int64{int64(12321231)}) // ListSecureWebGatewayDeviceSettingsRequest | Provide a list of origin ID for the devices in the organization. The list can contain 1–100 origin IDs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SWGDeviceSettingsAPI.ListSecureWebGatewayDeviceSettings(context.Background()).ListSecureWebGatewayDeviceSettingsRequest(listSecureWebGatewayDeviceSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SWGDeviceSettingsAPI.ListSecureWebGatewayDeviceSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSecureWebGatewayDeviceSettings`: []ListSWGDeviceSettingsInner
	fmt.Fprintf(os.Stdout, "Response from `SWGDeviceSettingsAPI.ListSecureWebGatewayDeviceSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSecureWebGatewayDeviceSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listSecureWebGatewayDeviceSettingsRequest** | [**ListSecureWebGatewayDeviceSettingsRequest**](ListSecureWebGatewayDeviceSettingsRequest.md) | Provide a list of origin ID for the devices in the organization. The list can contain 1–100 origin IDs. | 

### Return type

[**[]ListSWGDeviceSettingsInner**](ListSWGDeviceSettingsInner.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

