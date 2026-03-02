# \PrivateResourcesAPI

All URIs are relative to *https://api.sse.cisco.com/policies/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPrivateResource**](PrivateResourcesAPI.md#AddPrivateResource) | **Post** /privateResources | Create Private Resource
[**DeletePrivateResource**](PrivateResourcesAPI.md#DeletePrivateResource) | **Delete** /privateResources/{id} | Delete Private Resource
[**GetPrivateResource**](PrivateResourcesAPI.md#GetPrivateResource) | **Get** /privateResources/{id} | Get Private Resource
[**ListPrivateResources**](PrivateResourcesAPI.md#ListPrivateResources) | **Get** /privateResources | List Private Resources
[**PutPrivateResource**](PrivateResourcesAPI.md#PutPrivateResource) | **Put** /privateResources/{id} | Update Private Resource



## AddPrivateResource

> PrivateResourceResponse AddPrivateResource(ctx).PrivateResourceRequest(privateResourceRequest).Execute()

Create Private Resource



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cisco-sbg/go-ciscosecureaccess/privateapps"
)

func main() {
	privateResourceRequest := *openapiclient.NewPrivateResourceRequest("Jira", []openapiclient.AccessTypesRequestInner{*openapiclient.NewAccessTypesRequestInner("client", []string{"172.6.0.0/32"}, openapiclient.protocolProxyToResource("HTTPS"))}, []openapiclient.ResourceAddressesInner{*openapiclient.NewResourceAddressesInner([]string{"DestinationAddr_example"}, []openapiclient.ResourceAddressesInnerProtocolPortsInner{*openapiclient.NewResourceAddressesInnerProtocolPortsInner()})}) // PrivateResourceRequest | Create the Private Resource. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivateResourcesAPI.AddPrivateResource(context.Background()).PrivateResourceRequest(privateResourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivateResourcesAPI.AddPrivateResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddPrivateResource`: PrivateResourceResponse
	fmt.Fprintf(os.Stdout, "Response from `PrivateResourcesAPI.AddPrivateResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPrivateResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **privateResourceRequest** | [**PrivateResourceRequest**](PrivateResourceRequest.md) | Create the Private Resource. | 

### Return type

[**PrivateResourceResponse**](PrivateResourceResponse.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePrivateResource

> string DeletePrivateResource(ctx, id).Force(force).Execute()

Delete Private Resource



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cisco-sbg/go-ciscosecureaccess/privateapps"
)

func main() {
	id := int64(56) // int64 | The ID of the Private Resource.
	force := true // bool | Specify whether to force the removal of the Private Resource even if the resource is included on one or more private access rules. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivateResourcesAPI.DeletePrivateResource(context.Background(), id).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivateResourcesAPI.DeletePrivateResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePrivateResource`: string
	fmt.Fprintf(os.Stdout, "Response from `PrivateResourcesAPI.DeletePrivateResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the Private Resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrivateResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | Specify whether to force the removal of the Private Resource even if the resource is included on one or more private access rules. | [default to false]

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


## GetPrivateResource

> PrivateResourceResponse GetPrivateResource(ctx, id).Execute()

Get Private Resource



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cisco-sbg/go-ciscosecureaccess/privateapps"
)

func main() {
	id := int64(56) // int64 | The ID of the Private Resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivateResourcesAPI.GetPrivateResource(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivateResourcesAPI.GetPrivateResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrivateResource`: PrivateResourceResponse
	fmt.Fprintf(os.Stdout, "Response from `PrivateResourcesAPI.GetPrivateResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the Private Resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrivateResourceResponse**](PrivateResourceResponse.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPrivateResources

> PrivateResourceList ListPrivateResources(ctx).Filters(filters).Offset(offset).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Execute()

List Private Resources



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cisco-sbg/go-ciscosecureaccess/privateapps"
)

func main() {
	filters := map[string]interface{}{ ... } // map[string]interface{} | Filter the private resources in the organization by one or more properties: `name`, `description`, `resourceGroupId`, `accessType`, and `externalFQDN`.  * **name**-The name of the resource. The value is case insensitive. * **description**-The description of the resource. The value is case insensitive. * **resourceGroupId**-The ID of the resource group that includes the resource. * **accessType**-The type of access configured for the resource. Valid values are: `browser`, `network`, or `client`, or a combination of these values. * **externalFQDN**-The value of the resource's external fully qualified domain name (FQDN).  Specify the filters in the JSON format.  Example:  ``` { \"name\": \"West coast\", \"description\": \"West Coast\", \"resourceGroupId\": 1234, \"accessType\": \"network\", \"externalFQDN\": \"https://jira-2390150.ztna.ciscoplus.com\" } ``` (optional)
	offset := int64(10) // int64 | An integer that represents an index in the collection. If the limit is `10` and the offset is `0`, the offset for next page is `10`. (optional) (default to 0)
	limit := int64(40) // int64 | The number of items from the collection to return on a page. The default value is `10` items on a page. The maximum number of items that is allowed on a page is 1000. (optional) (default to 10)
	sortBy := "resourceId" // string | Specify the field that is used to sort the Private Resources. (optional) (default to "name")
	sortOrder := "desc" // string | Specify the order that is used to sort the Private Resources. The default value is `asc`. (optional) (default to "asc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivateResourcesAPI.ListPrivateResources(context.Background()).Filters(filters).Offset(offset).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivateResourcesAPI.ListPrivateResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPrivateResources`: PrivateResourceList
	fmt.Fprintf(os.Stdout, "Response from `PrivateResourcesAPI.ListPrivateResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPrivateResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | [**map[string]interface{}**](map[string]interface{}.md) | Filter the private resources in the organization by one or more properties: &#x60;name&#x60;, &#x60;description&#x60;, &#x60;resourceGroupId&#x60;, &#x60;accessType&#x60;, and &#x60;externalFQDN&#x60;.  * **name**-The name of the resource. The value is case insensitive. * **description**-The description of the resource. The value is case insensitive. * **resourceGroupId**-The ID of the resource group that includes the resource. * **accessType**-The type of access configured for the resource. Valid values are: &#x60;browser&#x60;, &#x60;network&#x60;, or &#x60;client&#x60;, or a combination of these values. * **externalFQDN**-The value of the resource&#39;s external fully qualified domain name (FQDN).  Specify the filters in the JSON format.  Example:  &#x60;&#x60;&#x60; { \&quot;name\&quot;: \&quot;West coast\&quot;, \&quot;description\&quot;: \&quot;West Coast\&quot;, \&quot;resourceGroupId\&quot;: 1234, \&quot;accessType\&quot;: \&quot;network\&quot;, \&quot;externalFQDN\&quot;: \&quot;https://jira-2390150.ztna.ciscoplus.com\&quot; } &#x60;&#x60;&#x60; | 
 **offset** | **int64** | An integer that represents an index in the collection. If the limit is &#x60;10&#x60; and the offset is &#x60;0&#x60;, the offset for next page is &#x60;10&#x60;. | [default to 0]
 **limit** | **int64** | The number of items from the collection to return on a page. The default value is &#x60;10&#x60; items on a page. The maximum number of items that is allowed on a page is 1000. | [default to 10]
 **sortBy** | **string** | Specify the field that is used to sort the Private Resources. | [default to &quot;name&quot;]
 **sortOrder** | **string** | Specify the order that is used to sort the Private Resources. The default value is &#x60;asc&#x60;. | [default to &quot;asc&quot;]

### Return type

[**PrivateResourceList**](PrivateResourceList.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPrivateResource

> PrivateResourceResponse PutPrivateResource(ctx, id).PrivateResourceRequest(privateResourceRequest).Execute()

Update Private Resource



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cisco-sbg/go-ciscosecureaccess/privateapps"
)

func main() {
	id := int64(56) // int64 | The ID of the Private Resource.
	privateResourceRequest := *openapiclient.NewPrivateResourceRequest("Jira", []openapiclient.AccessTypesRequestInner{*openapiclient.NewAccessTypesRequestInner("client", []string{"172.6.0.0/32"}, openapiclient.protocolProxyToResource("HTTPS"))}, []openapiclient.ResourceAddressesInner{*openapiclient.NewResourceAddressesInner([]string{"DestinationAddr_example"}, []openapiclient.ResourceAddressesInnerProtocolPortsInner{*openapiclient.NewResourceAddressesInnerProtocolPortsInner()})}) // PrivateResourceRequest | Update a Private Resource. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivateResourcesAPI.PutPrivateResource(context.Background(), id).PrivateResourceRequest(privateResourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivateResourcesAPI.PutPrivateResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPrivateResource`: PrivateResourceResponse
	fmt.Fprintf(os.Stdout, "Response from `PrivateResourcesAPI.PutPrivateResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the Private Resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPrivateResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privateResourceRequest** | [**PrivateResourceRequest**](PrivateResourceRequest.md) | Update a Private Resource. | 

### Return type

[**PrivateResourceResponse**](PrivateResourceResponse.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

