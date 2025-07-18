# \ConnectorsAPI

All URIs are relative to *https://api.sse.cisco.com/deployments/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConnector**](ConnectorsAPI.md#DeleteConnector) | **Delete** /connectorAgents/{id} | Delete Connector
[**GetConnector**](ConnectorsAPI.md#GetConnector) | **Get** /connectorAgents/{id} | Get Connector
[**GetConnectorAgentCount**](ConnectorsAPI.md#GetConnectorAgentCount) | **Get** /connectorAgents/counts | Get Counts Connectors States
[**ListConnectors**](ConnectorsAPI.md#ListConnectors) | **Get** /connectorAgents | List Connectors
[**PatchConnector**](ConnectorsAPI.md#PatchConnector) | **Patch** /connectorAgents/{id} | Patch Connector



## DeleteConnector

> string DeleteConnector(ctx, id).Execute()

Delete Connector



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/resconn"
)

func main() {
	id := int64(123455) // int64 | The ID of the Connector.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.DeleteConnector(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.DeleteConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteConnector`: string
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.DeleteConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the Connector. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectorRequest struct via the builder pattern


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


## GetConnector

> ConnectorResponse GetConnector(ctx, id).Execute()

Get Connector



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/resconn"
)

func main() {
	id := int64(123455) // int64 | The ID of the Connector.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.GetConnector(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.GetConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnector`: ConnectorResponse
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.GetConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the Connector. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectorResponse**](ConnectorResponse.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectorAgentCount

> ConnectorCountsResponse GetConnectorAgentCount(ctx).Execute()

Get Counts Connectors States



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/resconn"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.GetConnectorAgentCount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.GetConnectorAgentCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectorAgentCount`: ConnectorCountsResponse
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.GetConnectorAgentCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorAgentCountRequest struct via the builder pattern


### Return type

[**ConnectorCountsResponse**](ConnectorCountsResponse.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: resource/json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectors

> ConnectorListRes ListConnectors(ctx).Filters(filters).Offset(offset).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Execute()

List Connectors



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/resconn"
)

func main() {
	filters := "filters_example" // string | Filter the list of Connectors by one or more properties: `groupId`, `instanceId`, `status`, `confirmed`.  For example:  ``` {   \"groupId\": \"123\",   \"instanceId\": \"i-04568b6cdaeedda25a\",   \"status\": \"connected,reachable\",   \"confirmed\": true } ```  Filters are applied to the following properties:  * **groupId**—A comma-separated list of Connector Group IDs. * **instanceId**—The ID of the Connector instance. The value is not case sensitive. * **status**—A comma-separated list of Connector status labels. * **confirmed**—The confirmed location of the Connector. The value is not case sensitive. (optional)
	offset := int64(25) // int64 | The place to start reading in the collection. The offset starts at 0. If the `limit` is 10, the offset for the next response is 10. The default value is 0. (optional) (default to 0)
	limit := int64(100) // int64 | The maximum number of items to return in the response from the collection. The default value is 10. (optional) (default to 10)
	sortBy := "originIpAddress" // string | Specify a field to filter the collection. (optional) (default to "originIpAddress")
	sortOrder := "desc" // string | Specify a field in the response to order the collection. (optional) (default to "asc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.ListConnectors(context.Background()).Filters(filters).Offset(offset).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.ListConnectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConnectors`: ConnectorListRes
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.ListConnectors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **string** | Filter the list of Connectors by one or more properties: &#x60;groupId&#x60;, &#x60;instanceId&#x60;, &#x60;status&#x60;, &#x60;confirmed&#x60;.  For example:  &#x60;&#x60;&#x60; {   \&quot;groupId\&quot;: \&quot;123\&quot;,   \&quot;instanceId\&quot;: \&quot;i-04568b6cdaeedda25a\&quot;,   \&quot;status\&quot;: \&quot;connected,reachable\&quot;,   \&quot;confirmed\&quot;: true } &#x60;&#x60;&#x60;  Filters are applied to the following properties:  * **groupId**—A comma-separated list of Connector Group IDs. * **instanceId**—The ID of the Connector instance. The value is not case sensitive. * **status**—A comma-separated list of Connector status labels. * **confirmed**—The confirmed location of the Connector. The value is not case sensitive. | 
 **offset** | **int64** | The place to start reading in the collection. The offset starts at 0. If the &#x60;limit&#x60; is 10, the offset for the next response is 10. The default value is 0. | [default to 0]
 **limit** | **int64** | The maximum number of items to return in the response from the collection. The default value is 10. | [default to 10]
 **sortBy** | **string** | Specify a field to filter the collection. | [default to &quot;originIpAddress&quot;]
 **sortOrder** | **string** | Specify a field in the response to order the collection. | [default to &quot;asc&quot;]

### Return type

[**ConnectorListRes**](ConnectorListRes.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchConnector

> ConnectorResponse PatchConnector(ctx, id).ConnectorPatchReqInner(connectorPatchReqInner).Execute()

Patch Connector



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/resconn"
)

func main() {
	id := int64(123455) // int64 | The ID of the Connector.
	connectorPatchReqInner := []openapiclient.ConnectorPatchReqInner{*openapiclient.NewConnectorPatchReqInner()} // []ConnectorPatchReqInner | Update the properties on the Connector. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.PatchConnector(context.Background(), id).ConnectorPatchReqInner(connectorPatchReqInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.PatchConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchConnector`: ConnectorResponse
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.PatchConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the Connector. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectorPatchReqInner** | [**[]ConnectorPatchReqInner**](ConnectorPatchReqInner.md) | Update the properties on the Connector. | 

### Return type

[**ConnectorResponse**](ConnectorResponse.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

