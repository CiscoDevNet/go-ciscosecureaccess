# \RoamingComputersAPI

All URIs are relative to *https://api.sse.cisco.com/deployments/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRoamingComputer**](RoamingComputersAPI.md#DeleteRoamingComputer) | **Delete** /roamingcomputers/{deviceId} | Delete Roaming Computer
[**GetRoamingComputer**](RoamingComputersAPI.md#GetRoamingComputer) | **Get** /roamingcomputers/{deviceId} | Get Roaming Computer
[**ListRoamingComputers**](RoamingComputersAPI.md#ListRoamingComputers) | **Get** /roamingcomputers | List Roaming Computers
[**UpdateRoamingComputer**](RoamingComputersAPI.md#UpdateRoamingComputer) | **Put** /roamingcomputers/{deviceId} | Update Roaming Computer



## DeleteRoamingComputer

> DeleteRoamingComputer(ctx, deviceId).Execute()

Delete Roaming Computer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/roaming"
)

func main() {
	deviceId := "AB00C7DCEC99D211" // string | The device ID (deviceId) of the roaming computer.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoamingComputersAPI.DeleteRoamingComputer(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoamingComputersAPI.DeleteRoamingComputer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | The device ID (deviceId) of the roaming computer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoamingComputerRequest struct via the builder pattern


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


## GetRoamingComputer

> RoamingComputerObject GetRoamingComputer(ctx, deviceId).Execute()

Get Roaming Computer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/roaming"
)

func main() {
	deviceId := "AB00C7DCEC99D211" // string | The device ID (deviceId) of the roaming computer.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoamingComputersAPI.GetRoamingComputer(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoamingComputersAPI.GetRoamingComputer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoamingComputer`: RoamingComputerObject
	fmt.Fprintf(os.Stdout, "Response from `RoamingComputersAPI.GetRoamingComputer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | The device ID (deviceId) of the roaming computer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoamingComputerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RoamingComputerObject**](RoamingComputerObject.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoamingComputers

> []RoamingComputerObject ListRoamingComputers(ctx).Page(page).Limit(limit).Name(name).Status(status).SwgStatus(swgStatus).LastSyncBefore(lastSyncBefore).LastSyncAfter(lastSyncAfter).Execute()

List Roaming Computers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/roaming"
)

func main() {
	page := int64(3) // int64 | The number of a page in the collection. (optional) (default to 1)
	limit := int64(40) // int64 | The number of records in the collection to return on the page. (optional) (default to 100)
	name := "roaming-device-1" // string | The name of the roaming computer. (optional)
	status := "Network" // string | Filter for the status of the roaming computer with DNS-layer security. (optional)
	swgStatus := "Protected" // string | Filter for the status of the roaming computer with Internet security (Secure Web Gateway). (optional)
	lastSyncBefore := time.Now() // time.Time | The date and time (timestamp) before the last sync. (optional)
	lastSyncAfter := time.Now() // time.Time | The date and time (timestamp) after the last sync. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoamingComputersAPI.ListRoamingComputers(context.Background()).Page(page).Limit(limit).Name(name).Status(status).SwgStatus(swgStatus).LastSyncBefore(lastSyncBefore).LastSyncAfter(lastSyncAfter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoamingComputersAPI.ListRoamingComputers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRoamingComputers`: []RoamingComputerObject
	fmt.Fprintf(os.Stdout, "Response from `RoamingComputersAPI.ListRoamingComputers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRoamingComputersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | The number of a page in the collection. | [default to 1]
 **limit** | **int64** | The number of records in the collection to return on the page. | [default to 100]
 **name** | **string** | The name of the roaming computer. | 
 **status** | **string** | Filter for the status of the roaming computer with DNS-layer security. | 
 **swgStatus** | **string** | Filter for the status of the roaming computer with Internet security (Secure Web Gateway). | 
 **lastSyncBefore** | **time.Time** | The date and time (timestamp) before the last sync. | 
 **lastSyncAfter** | **time.Time** | The date and time (timestamp) after the last sync. | 

### Return type

[**[]RoamingComputerObject**](RoamingComputerObject.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoamingComputer

> RoamingComputerObject UpdateRoamingComputer(ctx, deviceId).UpdateRoamingComputerRequest(updateRoamingComputerRequest).Execute()

Update Roaming Computer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/roaming"
)

func main() {
	deviceId := "AB00C7DCEC99D211" // string | The device ID (deviceId) of the roaming computer.
	updateRoamingComputerRequest := *openapiclient.NewUpdateRoamingComputerRequest("roaming-computer-one") // UpdateRoamingComputerRequest | Update the roaming computer. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoamingComputersAPI.UpdateRoamingComputer(context.Background(), deviceId).UpdateRoamingComputerRequest(updateRoamingComputerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoamingComputersAPI.UpdateRoamingComputer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoamingComputer`: RoamingComputerObject
	fmt.Fprintf(os.Stdout, "Response from `RoamingComputersAPI.UpdateRoamingComputer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | The device ID (deviceId) of the roaming computer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoamingComputerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRoamingComputerRequest** | [**UpdateRoamingComputerRequest**](UpdateRoamingComputerRequest.md) | Update the roaming computer. | 

### Return type

[**RoamingComputerObject**](RoamingComputerObject.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

