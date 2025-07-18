# \ConnectorGroupsAPI

All URIs are relative to *https://api.sse.cisco.com/deployments/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnectorGroup**](ConnectorGroupsAPI.md#CreateConnectorGroup) | **Post** /connectorGroups | Create Resource Connector Group
[**DeleteConnectorGroup**](ConnectorGroupsAPI.md#DeleteConnectorGroup) | **Delete** /connectorGroups/{id} | Delete Resource Connector Group
[**GetConnectorGroup**](ConnectorGroupsAPI.md#GetConnectorGroup) | **Get** /connectorGroups/{id} | Get Resource Connector Group
[**GetConnectorGroupCount**](ConnectorGroupsAPI.md#GetConnectorGroupCount) | **Get** /connectorGroups/counts | Get Counts Resource Connector Groups States
[**ListConnectorGroups**](ConnectorGroupsAPI.md#ListConnectorGroups) | **Get** /connectorGroups | List Resource Connector Groups
[**PatchConnectorGroup**](ConnectorGroupsAPI.md#PatchConnectorGroup) | **Patch** /connectorGroups/{id} | Patch Resource Connector Group
[**PutConnectorGroup**](ConnectorGroupsAPI.md#PutConnectorGroup) | **Put** /connectorGroups/{id} | Update Resource Connector Group



## CreateConnectorGroup

> ConnectorGroupResponse CreateConnectorGroup(ctx).ConnectorGroupReq(connectorGroupReq).Execute()

Create Resource Connector Group



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
	connectorGroupReq := *openapiclient.NewConnectorGroupReq("NYC Connector Group", "us-west-2") // ConnectorGroupReq | Create a Resource Connector Group. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorGroupsAPI.CreateConnectorGroup(context.Background()).ConnectorGroupReq(connectorGroupReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorGroupsAPI.CreateConnectorGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConnectorGroup`: ConnectorGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `ConnectorGroupsAPI.CreateConnectorGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectorGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectorGroupReq** | [**ConnectorGroupReq**](ConnectorGroupReq.md) | Create a Resource Connector Group. | 

### Return type

[**ConnectorGroupResponse**](ConnectorGroupResponse.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, resource/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnectorGroup

> string DeleteConnectorGroup(ctx, id).Execute()

Delete Resource Connector Group



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
	id := int64(123455) // int64 | The ID of the Connector Group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorGroupsAPI.DeleteConnectorGroup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorGroupsAPI.DeleteConnectorGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteConnectorGroup`: string
	fmt.Fprintf(os.Stdout, "Response from `ConnectorGroupsAPI.DeleteConnectorGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the Connector Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectorGroupRequest struct via the builder pattern


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


## GetConnectorGroup

> ConnectorGroupResponse GetConnectorGroup(ctx, id).IncludeProvisioningKey(includeProvisioningKey).Execute()

Get Resource Connector Group



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
	id := int64(123455) // int64 | The ID of the Connector Group.
	includeProvisioningKey := true // bool | Specify whether to include the Connector Group's provisioning key in the response. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorGroupsAPI.GetConnectorGroup(context.Background(), id).IncludeProvisioningKey(includeProvisioningKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorGroupsAPI.GetConnectorGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectorGroup`: ConnectorGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `ConnectorGroupsAPI.GetConnectorGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the Connector Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeProvisioningKey** | **bool** | Specify whether to include the Connector Group&#39;s provisioning key in the response. | [default to false]

### Return type

[**ConnectorGroupResponse**](ConnectorGroupResponse.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectorGroupCount

> ConnectorGroupCountsResponse GetConnectorGroupCount(ctx).Execute()

Get Counts Resource Connector Groups States



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
	resp, r, err := apiClient.ConnectorGroupsAPI.GetConnectorGroupCount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorGroupsAPI.GetConnectorGroupCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectorGroupCount`: ConnectorGroupCountsResponse
	fmt.Fprintf(os.Stdout, "Response from `ConnectorGroupsAPI.GetConnectorGroupCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorGroupCountRequest struct via the builder pattern


### Return type

[**ConnectorGroupCountsResponse**](ConnectorGroupCountsResponse.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: resource/json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectorGroups

> ConnectorGroupList ListConnectorGroups(ctx).IncludeProvisioningKey(includeProvisioningKey).Filters(filters).Offset(offset).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Execute()

List Resource Connector Groups



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
	includeProvisioningKey := true // bool | Specify whether to include the Connector Group's provisioning key in the response. (optional) (default to false)
	filters := "filters_example" // string | Filter the list of Connector Groups by one or more properties: `groupId`, `name`, `location`, `status`, `hasResources`.  For example:  ``` {   \"groupId\": \"123,987\",   \"name\": \"NYC Connector Group\",   \"status\": \"connected,waiting\",   \"location\": \"us-east-1\",   \"hasResources\": false } ```  Filters are applied to the following properties:  * **groupId**—A comma-separated list of Connector Group IDs. * **name**—The name of the Connector Groups. The value is not case sensitive. * **location**—The location of Connector Groups. * **status**—A comma-separated list of Connector Group status labels. * **hasResources**—When `false`, Secure Access returns Connector Groups with no resource. * **hasAgentWithStatus**—List the Connector Groups that have at least one Connector with the specified status. * **environment**—List the Connector Groups that have the specified environment, either `aws`, `azure`, or `esx`. * **resourceIds**—Connector groups that have at least one of the private resources that are defined in the comma separated list that is mapped to the Resource Connector Group. When `true`, Secure Access returns Connector Groups with at least one resource. When the value is not set, Secure Access returns Connector Groups regardless of the number of the resources. (optional)
	offset := int64(25) // int64 | The place to start reading in the collection. The offset starts at 0. If the `limit` is 10, the offset for the next response is 10. The default value is 0. (optional) (default to 0)
	limit := int64(100) // int64 | The maximum number of items to return in the response from the collection. The default value is 10. (optional) (default to 10)
	sortBy := "resourceIds" // string | Specify a field to filter the collection. (optional) (default to "name")
	sortOrder := "desc" // string | Specify a field in the response to order the collection. (optional) (default to "asc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorGroupsAPI.ListConnectorGroups(context.Background()).IncludeProvisioningKey(includeProvisioningKey).Filters(filters).Offset(offset).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorGroupsAPI.ListConnectorGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConnectorGroups`: ConnectorGroupList
	fmt.Fprintf(os.Stdout, "Response from `ConnectorGroupsAPI.ListConnectorGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConnectorGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeProvisioningKey** | **bool** | Specify whether to include the Connector Group&#39;s provisioning key in the response. | [default to false]
 **filters** | **string** | Filter the list of Connector Groups by one or more properties: &#x60;groupId&#x60;, &#x60;name&#x60;, &#x60;location&#x60;, &#x60;status&#x60;, &#x60;hasResources&#x60;.  For example:  &#x60;&#x60;&#x60; {   \&quot;groupId\&quot;: \&quot;123,987\&quot;,   \&quot;name\&quot;: \&quot;NYC Connector Group\&quot;,   \&quot;status\&quot;: \&quot;connected,waiting\&quot;,   \&quot;location\&quot;: \&quot;us-east-1\&quot;,   \&quot;hasResources\&quot;: false } &#x60;&#x60;&#x60;  Filters are applied to the following properties:  * **groupId**—A comma-separated list of Connector Group IDs. * **name**—The name of the Connector Groups. The value is not case sensitive. * **location**—The location of Connector Groups. * **status**—A comma-separated list of Connector Group status labels. * **hasResources**—When &#x60;false&#x60;, Secure Access returns Connector Groups with no resource. * **hasAgentWithStatus**—List the Connector Groups that have at least one Connector with the specified status. * **environment**—List the Connector Groups that have the specified environment, either &#x60;aws&#x60;, &#x60;azure&#x60;, or &#x60;esx&#x60;. * **resourceIds**—Connector groups that have at least one of the private resources that are defined in the comma separated list that is mapped to the Resource Connector Group. When &#x60;true&#x60;, Secure Access returns Connector Groups with at least one resource. When the value is not set, Secure Access returns Connector Groups regardless of the number of the resources. | 
 **offset** | **int64** | The place to start reading in the collection. The offset starts at 0. If the &#x60;limit&#x60; is 10, the offset for the next response is 10. The default value is 0. | [default to 0]
 **limit** | **int64** | The maximum number of items to return in the response from the collection. The default value is 10. | [default to 10]
 **sortBy** | **string** | Specify a field to filter the collection. | [default to &quot;name&quot;]
 **sortOrder** | **string** | Specify a field in the response to order the collection. | [default to &quot;asc&quot;]

### Return type

[**ConnectorGroupList**](ConnectorGroupList.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchConnectorGroup

> ConnectorGroupResponse PatchConnectorGroup(ctx, id).ConnectorGroupPatchReqInner(connectorGroupPatchReqInner).Execute()

Patch Resource Connector Group



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
	id := int64(123455) // int64 | The ID of the Connector Group.
	connectorGroupPatchReqInner := []openapiclient.ConnectorGroupPatchReqInner{*openapiclient.NewConnectorGroupPatchReqInner(openapiclient.op("replace"), "/name", "us-east-1")} // []ConnectorGroupPatchReqInner | Set the properties on the Connector Group. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorGroupsAPI.PatchConnectorGroup(context.Background(), id).ConnectorGroupPatchReqInner(connectorGroupPatchReqInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorGroupsAPI.PatchConnectorGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchConnectorGroup`: ConnectorGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `ConnectorGroupsAPI.PatchConnectorGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the Connector Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchConnectorGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectorGroupPatchReqInner** | [**[]ConnectorGroupPatchReqInner**](ConnectorGroupPatchReqInner.md) | Set the properties on the Connector Group. | 

### Return type

[**ConnectorGroupResponse**](ConnectorGroupResponse.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutConnectorGroup

> ConnectorGroupResponse PutConnectorGroup(ctx, id).ConnectorGroupReq(connectorGroupReq).Execute()

Update Resource Connector Group



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
	id := int64(123455) // int64 | The ID of the Connector Group.
	connectorGroupReq := *openapiclient.NewConnectorGroupReq("NYC Connector Group", "us-west-2") // ConnectorGroupReq | Set the properties on the Resource Connector Group. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorGroupsAPI.PutConnectorGroup(context.Background(), id).ConnectorGroupReq(connectorGroupReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorGroupsAPI.PutConnectorGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutConnectorGroup`: ConnectorGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `ConnectorGroupsAPI.PutConnectorGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the Connector Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutConnectorGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectorGroupReq** | [**ConnectorGroupReq**](ConnectorGroupReq.md) | Set the properties on the Resource Connector Group. | 

### Return type

[**ConnectorGroupResponse**](ConnectorGroupResponse.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

