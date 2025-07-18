# \RemoteAccessAPI

All URIs are relative to *https://api.sse.cisco.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRemoteAccessEvents**](RemoteAccessAPI.md#GetRemoteAccessEvents) | **Get** /reports/v2/remote-access-events | Get Remote Access Events



## GetRemoteAccessEvents

> GetRemoteAccessEvents200Response GetRemoteAccessEvents(ctx).From(from).To(to).Limit(limit).Ip(ip).Identityids(identityids).Identitytypes(identitytypes).Connectionevent(connectionevent).Anyconnectversions(anyconnectversions).Osversions(osversions).Timezone(timezone).Execute()

Get Remote Access Events



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
	from := "1639146300000" // string | A timestamp or relative time string (for example: '-1days'). Filter for data that appears after this time.
	to := "1640010300000" // string | A timestamp or relative time string (for example: 'now'). Filter for data that appears before this time.
	limit := int64(100) // int64 | The maximum number of records to return from the collection. (default to 100)
	ip := "10.10.10.10" // string | An IP address. (optional)
	identityids := "1,2,3" // string | An identity ID or comma-delimited list of identity IDs. (optional)
	identitytypes := "network,roaming" // string | An identity type or comma-delimited list of identity types. (optional)
	connectionevent := "connected" // string | Specify the type of connection event. (optional)
	anyconnectversions := "4.10.05095,5.10" // string | Specify a comma-separated list of AnyConnect Roaming Security module versions. (optional)
	osversions := "linux-64-Ubuntu 20.04.5 LTS (Focal Fossa)" // string | Specify a comma-separated list of OS versions. (optional)
	timezone := "ASIA%2fCALCUTTA" // string | Display the timestamp of the traffic events in the specified timezone. For the timezone, provide a continent and city separated by an url-encoded forward slash ('/'), for example: timezone='ASIA%2fCALCUTTA'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteAccessAPI.GetRemoteAccessEvents(context.Background()).From(from).To(to).Limit(limit).Ip(ip).Identityids(identityids).Identitytypes(identitytypes).Connectionevent(connectionevent).Anyconnectversions(anyconnectversions).Osversions(osversions).Timezone(timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteAccessAPI.GetRemoteAccessEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemoteAccessEvents`: GetRemoteAccessEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `RemoteAccessAPI.GetRemoteAccessEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteAccessEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | A timestamp or relative time string (for example: &#39;-1days&#39;). Filter for data that appears after this time. | 
 **to** | **string** | A timestamp or relative time string (for example: &#39;now&#39;). Filter for data that appears before this time. | 
 **limit** | **int64** | The maximum number of records to return from the collection. | [default to 100]
 **ip** | **string** | An IP address. | 
 **identityids** | **string** | An identity ID or comma-delimited list of identity IDs. | 
 **identitytypes** | **string** | An identity type or comma-delimited list of identity types. | 
 **connectionevent** | **string** | Specify the type of connection event. | 
 **anyconnectversions** | **string** | Specify a comma-separated list of AnyConnect Roaming Security module versions. | 
 **osversions** | **string** | Specify a comma-separated list of OS versions. | 
 **timezone** | **string** | Display the timestamp of the traffic events in the specified timezone. For the timezone, provide a continent and city separated by an url-encoded forward slash (&#39;/&#39;), for example: timezone&#x3D;&#39;ASIA%2fCALCUTTA&#39;. | 

### Return type

[**GetRemoteAccessEvents200Response**](GetRemoteAccessEvents200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

