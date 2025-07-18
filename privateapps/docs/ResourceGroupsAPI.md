# \ResourceGroupsAPI

All URIs are relative to *https://api.sse.cisco.com/policies/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPrivateResourceGroup**](ResourceGroupsAPI.md#AddPrivateResourceGroup) | **Post** /privateResourceGroups | Create Resource Group
[**DeletePrivateResourceGroup**](ResourceGroupsAPI.md#DeletePrivateResourceGroup) | **Delete** /privateResourceGroups/{id} | Delete Resource Group
[**GetPrivateResourceGroup**](ResourceGroupsAPI.md#GetPrivateResourceGroup) | **Get** /privateResourceGroups/{id} | Get Resource Group
[**ListResourceGroups**](ResourceGroupsAPI.md#ListResourceGroups) | **Get** /privateResourceGroups | List Resource Groups
[**PutPrivateResourceGroup**](ResourceGroupsAPI.md#PutPrivateResourceGroup) | **Put** /privateResourceGroups/{id} | Update Resource Group



## AddPrivateResourceGroup

> PrivateResourceGroupResponse AddPrivateResourceGroup(ctx).PrivateResourceGroupRequest(privateResourceGroupRequest).Execute()

Create Resource Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/privateapps"
)

func main() {
	privateResourceGroupRequest := *openapiclient.NewPrivateResourceGroupRequest("Westcoast Data Center", []int64{int64(255)}) // PrivateResourceGroupRequest | Create the Resource Group. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceGroupsAPI.AddPrivateResourceGroup(context.Background()).PrivateResourceGroupRequest(privateResourceGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceGroupsAPI.AddPrivateResourceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddPrivateResourceGroup`: PrivateResourceGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourceGroupsAPI.AddPrivateResourceGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPrivateResourceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **privateResourceGroupRequest** | [**PrivateResourceGroupRequest**](PrivateResourceGroupRequest.md) | Create the Resource Group. | 

### Return type

[**PrivateResourceGroupResponse**](PrivateResourceGroupResponse.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePrivateResourceGroup

> string DeletePrivateResourceGroup(ctx, id).Force(force).Execute()

Delete Resource Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/privateapps"
)

func main() {
	id := int64(56) // int64 | The ID of the Private Resource Group.
	force := true // bool | Specify whether to force the removal of the Private Resource Group even if the group is included in one or more private access rules. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceGroupsAPI.DeletePrivateResourceGroup(context.Background(), id).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceGroupsAPI.DeletePrivateResourceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePrivateResourceGroup`: string
	fmt.Fprintf(os.Stdout, "Response from `ResourceGroupsAPI.DeletePrivateResourceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the Private Resource Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrivateResourceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | Specify whether to force the removal of the Private Resource Group even if the group is included in one or more private access rules. | [default to false]

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


## GetPrivateResourceGroup

> PrivateResourceGroupResponse GetPrivateResourceGroup(ctx, id).Execute()

Get Resource Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/privateapps"
)

func main() {
	id := int64(56) // int64 | The ID of the Private Resource Group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceGroupsAPI.GetPrivateResourceGroup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceGroupsAPI.GetPrivateResourceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrivateResourceGroup`: PrivateResourceGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourceGroupsAPI.GetPrivateResourceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the Private Resource Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateResourceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrivateResourceGroupResponse**](PrivateResourceGroupResponse.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResourceGroups

> PrivateResourceGroupList ListResourceGroups(ctx).Filters(filters).Offset(offset).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Execute()

List Resource Groups



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/privateapps"
)

func main() {
	filters := map[string]interface{}{ ... } // map[string]interface{} | Filter the resource groups in the organization by one or more properties: `name`, `description`, `resourceId`, and `accessType`.  * **name**—The name of the resource group. The value is case insensitive. * **description**—The description of the Resource Group. The value is case insensitive. * **resourceId**—The resource group includes the resource, which you can identify by the ID. * **accessType**—The resource group contains at least one resource. The resource has one of the types of access: `branch`, `network`, `client`, or `browser`.  Specify the filters in the JSON format.  Example:  ``` { \"name\": \"West coast\", \"description\": \"West Coast\", \"resourceId\": 1234, \"accessType\": \"network\" } ``` (optional)
	offset := int64(10) // int64 | An integer that represents an index in the collection. If the limit is `10` and the offset is `0`, the offset for next page is `10`. (optional) (default to 0)
	limit := int64(40) // int64 | The number of items from the collection to return on a page. The default value is `10` items on a page. The maximum number of items that is allowed on a page is 1000. (optional) (default to 10)
	sortBy := "name" // string | Specify the field that is used to sort the Private Resource Groups. (optional) (default to "name")
	sortOrder := "desc" // string | Specify the order that is used to sort the Private Resources. The default value is `asc`. (optional) (default to "asc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceGroupsAPI.ListResourceGroups(context.Background()).Filters(filters).Offset(offset).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceGroupsAPI.ListResourceGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListResourceGroups`: PrivateResourceGroupList
	fmt.Fprintf(os.Stdout, "Response from `ResourceGroupsAPI.ListResourceGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListResourceGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | [**map[string]interface{}**](map[string]interface{}.md) | Filter the resource groups in the organization by one or more properties: &#x60;name&#x60;, &#x60;description&#x60;, &#x60;resourceId&#x60;, and &#x60;accessType&#x60;.  * **name**—The name of the resource group. The value is case insensitive. * **description**—The description of the Resource Group. The value is case insensitive. * **resourceId**—The resource group includes the resource, which you can identify by the ID. * **accessType**—The resource group contains at least one resource. The resource has one of the types of access: &#x60;branch&#x60;, &#x60;network&#x60;, &#x60;client&#x60;, or &#x60;browser&#x60;.  Specify the filters in the JSON format.  Example:  &#x60;&#x60;&#x60; { \&quot;name\&quot;: \&quot;West coast\&quot;, \&quot;description\&quot;: \&quot;West Coast\&quot;, \&quot;resourceId\&quot;: 1234, \&quot;accessType\&quot;: \&quot;network\&quot; } &#x60;&#x60;&#x60; | 
 **offset** | **int64** | An integer that represents an index in the collection. If the limit is &#x60;10&#x60; and the offset is &#x60;0&#x60;, the offset for next page is &#x60;10&#x60;. | [default to 0]
 **limit** | **int64** | The number of items from the collection to return on a page. The default value is &#x60;10&#x60; items on a page. The maximum number of items that is allowed on a page is 1000. | [default to 10]
 **sortBy** | **string** | Specify the field that is used to sort the Private Resource Groups. | [default to &quot;name&quot;]
 **sortOrder** | **string** | Specify the order that is used to sort the Private Resources. The default value is &#x60;asc&#x60;. | [default to &quot;asc&quot;]

### Return type

[**PrivateResourceGroupList**](PrivateResourceGroupList.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPrivateResourceGroup

> PrivateResourceGroupResponse PutPrivateResourceGroup(ctx, id).PrivateResourceGroupRequest(privateResourceGroupRequest).Execute()

Update Resource Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/privateapps"
)

func main() {
	id := int64(56) // int64 | The ID of the Private Resource Group.
	privateResourceGroupRequest := *openapiclient.NewPrivateResourceGroupRequest("Westcoast Data Center", []int64{int64(255)}) // PrivateResourceGroupRequest | Update the properties of the Private Resource Group. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceGroupsAPI.PutPrivateResourceGroup(context.Background(), id).PrivateResourceGroupRequest(privateResourceGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceGroupsAPI.PutPrivateResourceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPrivateResourceGroup`: PrivateResourceGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourceGroupsAPI.PutPrivateResourceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the Private Resource Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPrivateResourceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privateResourceGroupRequest** | [**PrivateResourceGroupRequest**](PrivateResourceGroupRequest.md) | Update the properties of the Private Resource Group. | 

### Return type

[**PrivateResourceGroupResponse**](PrivateResourceGroupResponse.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

