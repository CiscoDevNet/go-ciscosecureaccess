# \ProtocolsAPI

All URIs are relative to *https://api.sse.cisco.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProtocol**](ProtocolsAPI.md#GetProtocol) | **Get** /reports/v2/appDiscovery/protocols/{protocolId} | Get Protocol
[**GetProtocols**](ProtocolsAPI.md#GetProtocols) | **Get** /reports/v2/appDiscovery/protocols | List Protocols



## GetProtocol

> ProtocolObject GetProtocol(ctx, protocolId).Date(date).Execute()

Get Protocol



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/reports"
)

func main() {
	protocolId := "protocolId_example" // string | The ID of the protocol.
	date := time.Now() // string | Specify a date to search for data within a twenty-four hour time period. If you do not provide a date, the last 90 days period is used to query the collection. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProtocolsAPI.GetProtocol(context.Background(), protocolId).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtocolsAPI.GetProtocol``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProtocol`: ProtocolObject
	fmt.Fprintf(os.Stdout, "Response from `ProtocolsAPI.GetProtocol`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**protocolId** | **string** | The ID of the protocol. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProtocolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **date** | **string** | Specify a date to search for data within a twenty-four hour time period. If you do not provide a date, the last 90 days period is used to query the collection. | 

### Return type

[**ProtocolObject**](ProtocolObject.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtocols

> ProtocolList GetProtocols(ctx).Identity(identity).Limit(limit).Offset(offset).Sort(sort).Order(order).Execute()

List Protocols



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/reports"
)

func main() {
	identity := int64(12355) // int64 | Specify the entity ID to filter the collection. (optional)
	limit := int64(56) // int64 | The maximum number of items to return in the collection. (optional)
	offset := int64(56) // int64 | The number of items to skip before starting to collect the result set. (optional)
	sort := "firstDetected" // string | Specify the name of a field to sort the application protocols. (optional)
	order := "asc" // string | Specify the order to sort the collection. Valid values are: `asc` (ascending) or `desc` (descending). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProtocolsAPI.GetProtocols(context.Background()).Identity(identity).Limit(limit).Offset(offset).Sort(sort).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtocolsAPI.GetProtocols``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProtocols`: ProtocolList
	fmt.Fprintf(os.Stdout, "Response from `ProtocolsAPI.GetProtocols`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProtocolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identity** | **int64** | Specify the entity ID to filter the collection. | 
 **limit** | **int64** | The maximum number of items to return in the collection. | 
 **offset** | **int64** | The number of items to skip before starting to collect the result set. | 
 **sort** | **string** | Specify the name of a field to sort the application protocols. | 
 **order** | **string** | Specify the order to sort the collection. Valid values are: &#x60;asc&#x60; (ascending) or &#x60;desc&#x60; (descending). | 

### Return type

[**ProtocolList**](ProtocolList.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

