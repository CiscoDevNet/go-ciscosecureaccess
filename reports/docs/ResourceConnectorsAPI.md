# \ResourceConnectorsAPI

All URIs are relative to *https://api.sse.cisco.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDetailedStatsForAppConnector**](ResourceConnectorsAPI.md#GetDetailedStatsForAppConnector) | **Get** /reports/v2/app-connectors/detailed-stats-timerange | Get Detailed Resource Connector Statistics
[**GetDetailedStatsForAppConnectorGroups**](ResourceConnectorsAPI.md#GetDetailedStatsForAppConnectorGroups) | **Get** /reports/v2/app-connectors/groups/detailed-stats-timerange | Get Detailed Resource Connector Group Statistics
[**GetOverloadedGroupsCount**](ResourceConnectorsAPI.md#GetOverloadedGroupsCount) | **Get** /reports/v2/app-connectors/groups/overloaded-count | Get Count of Overloaded Groups



## GetDetailedStatsForAppConnector

> GetDetailedStatsForAppConnector200Response GetDetailedStatsForAppConnector(ctx).From(from).To(to).Agentids(agentids).Timerange(timerange).Execute()

Get Detailed Resource Connector Statistics



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
	agentids := "31,47" // string | A comma-delimited list of resource connector IDs.
	timerange := "minute" // string | A string that represents a range of time. If the header is not set, the header value defaults to `hour`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceConnectorsAPI.GetDetailedStatsForAppConnector(context.Background()).From(from).To(to).Agentids(agentids).Timerange(timerange).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceConnectorsAPI.GetDetailedStatsForAppConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDetailedStatsForAppConnector`: GetDetailedStatsForAppConnector200Response
	fmt.Fprintf(os.Stdout, "Response from `ResourceConnectorsAPI.GetDetailedStatsForAppConnector`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDetailedStatsForAppConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | A timestamp or relative time string (for example: &#39;-1days&#39;). Filter for data that appears after this time. | 
 **to** | **string** | A timestamp or relative time string (for example: &#39;now&#39;). Filter for data that appears before this time. | 
 **agentids** | **string** | A comma-delimited list of resource connector IDs. | 
 **timerange** | **string** | A string that represents a range of time. If the header is not set, the header value defaults to &#x60;hour&#x60;. | 

### Return type

[**GetDetailedStatsForAppConnector200Response**](GetDetailedStatsForAppConnector200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDetailedStatsForAppConnectorGroups

> GetDetailedStatsForAppConnectorGroups200Response GetDetailedStatsForAppConnectorGroups(ctx).From(from).To(to).Groupids(groupids).Cputhreshold(cputhreshold).Timerange(timerange).Execute()

Get Detailed Resource Connector Group Statistics



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
	groupids := "31,47" // string | A comma-delimited list of resource connector group IDs.
	cputhreshold := int64(16) // int64 | The CPU threshold (`cputhreshold`) percentage that is used to filter overloaded groups.
	timerange := "minute" // string | A string that represents a range of time. If the header is not set, the header value defaults to `hour`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceConnectorsAPI.GetDetailedStatsForAppConnectorGroups(context.Background()).From(from).To(to).Groupids(groupids).Cputhreshold(cputhreshold).Timerange(timerange).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceConnectorsAPI.GetDetailedStatsForAppConnectorGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDetailedStatsForAppConnectorGroups`: GetDetailedStatsForAppConnectorGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `ResourceConnectorsAPI.GetDetailedStatsForAppConnectorGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDetailedStatsForAppConnectorGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | A timestamp or relative time string (for example: &#39;-1days&#39;). Filter for data that appears after this time. | 
 **to** | **string** | A timestamp or relative time string (for example: &#39;now&#39;). Filter for data that appears before this time. | 
 **groupids** | **string** | A comma-delimited list of resource connector group IDs. | 
 **cputhreshold** | **int64** | The CPU threshold (&#x60;cputhreshold&#x60;) percentage that is used to filter overloaded groups. | 
 **timerange** | **string** | A string that represents a range of time. If the header is not set, the header value defaults to &#x60;hour&#x60;. | 

### Return type

[**GetDetailedStatsForAppConnectorGroups200Response**](GetDetailedStatsForAppConnectorGroups200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOverloadedGroupsCount

> GetTotalRequests200Response GetOverloadedGroupsCount(ctx).From(from).To(to).Cputhreshold(cputhreshold).Execute()

Get Count of Overloaded Groups



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
	cputhreshold := int64(16) // int64 | The CPU threshold (`cputhreshold`) percentage that is used to filter overloaded groups.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceConnectorsAPI.GetOverloadedGroupsCount(context.Background()).From(from).To(to).Cputhreshold(cputhreshold).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceConnectorsAPI.GetOverloadedGroupsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOverloadedGroupsCount`: GetTotalRequests200Response
	fmt.Fprintf(os.Stdout, "Response from `ResourceConnectorsAPI.GetOverloadedGroupsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOverloadedGroupsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | A timestamp or relative time string (for example: &#39;-1days&#39;). Filter for data that appears after this time. | 
 **to** | **string** | A timestamp or relative time string (for example: &#39;now&#39;). Filter for data that appears before this time. | 
 **cputhreshold** | **int64** | The CPU threshold (&#x60;cputhreshold&#x60;) percentage that is used to filter overloaded groups. | 

### Return type

[**GetTotalRequests200Response**](GetTotalRequests200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

