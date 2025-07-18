# \ApplicationListsAPI

All URIs are relative to *https://api.sse.cisco.com/policies/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationList**](ApplicationListsAPI.md#CreateApplicationList) | **Post** /applicationLists | Create Application List
[**DeleteApplicationList**](ApplicationListsAPI.md#DeleteApplicationList) | **Delete** /applicationLists/{applicationListId} | Delete Application List
[**GetApplicationList**](ApplicationListsAPI.md#GetApplicationList) | **Get** /applicationLists/{applicationListId} | Get Application List
[**GetApplicationLists**](ApplicationListsAPI.md#GetApplicationLists) | **Get** /applicationLists | Get Application Lists
[**PutApplicationList**](ApplicationListsAPI.md#PutApplicationList) | **Put** /applicationLists/{applicationListId} | Update Application List
[**UpdateUsageApplications**](ApplicationListsAPI.md#UpdateUsageApplications) | **Get** /applications/usage | Get Usage of Applications



## CreateApplicationList

> ApplicationList CreateApplicationList(ctx).ApplicationListRequest(applicationListRequest).Execute()

Create Application List



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/rules"
)

func main() {
	applicationListRequest := *openapiclient.NewApplicationListRequest("Application lists on branch_1.", false, []int64{int64(2453454)}) // ApplicationListRequest | Create the application list.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationListsAPI.CreateApplicationList(context.Background()).ApplicationListRequest(applicationListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationListsAPI.CreateApplicationList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplicationList`: ApplicationList
	fmt.Fprintf(os.Stdout, "Response from `ApplicationListsAPI.CreateApplicationList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationListRequest** | [**ApplicationListRequest**](ApplicationListRequest.md) | Create the application list. | 

### Return type

[**ApplicationList**](ApplicationList.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationList

> ApplicationLists DeleteApplicationList(ctx, applicationListId).Execute()

Delete Application List



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/rules"
)

func main() {
	applicationListId := int64(56) // int64 | The ID of the application list.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationListsAPI.DeleteApplicationList(context.Background(), applicationListId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationListsAPI.DeleteApplicationList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApplicationList`: ApplicationLists
	fmt.Fprintf(os.Stdout, "Response from `ApplicationListsAPI.DeleteApplicationList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationListId** | **int64** | The ID of the application list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationLists**](ApplicationLists.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationList

> ApplicationList GetApplicationList(ctx, applicationListId).Execute()

Get Application List



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/rules"
)

func main() {
	applicationListId := int64(56) // int64 | The ID of the application list.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationListsAPI.GetApplicationList(context.Background(), applicationListId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationListsAPI.GetApplicationList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationList`: ApplicationList
	fmt.Fprintf(os.Stdout, "Response from `ApplicationListsAPI.GetApplicationList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationListId** | **int64** | The ID of the application list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationList**](ApplicationList.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationLists

> ApplicationLists GetApplicationLists(ctx).Execute()

Get Application Lists



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/rules"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationListsAPI.GetApplicationLists(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationListsAPI.GetApplicationLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationLists`: ApplicationLists
	fmt.Fprintf(os.Stdout, "Response from `ApplicationListsAPI.GetApplicationLists`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationListsRequest struct via the builder pattern


### Return type

[**ApplicationLists**](ApplicationLists.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutApplicationList

> ApplicationList PutApplicationList(ctx, applicationListId).ApplicationListRequest(applicationListRequest).Execute()

Update Application List



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/rules"
)

func main() {
	applicationListId := int64(56) // int64 | The ID of the application list.
	applicationListRequest := *openapiclient.NewApplicationListRequest("Application lists on branch_1.", false, []int64{int64(2453454)}) // ApplicationListRequest | Update the application list.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationListsAPI.PutApplicationList(context.Background(), applicationListId).ApplicationListRequest(applicationListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationListsAPI.PutApplicationList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutApplicationList`: ApplicationList
	fmt.Fprintf(os.Stdout, "Response from `ApplicationListsAPI.PutApplicationList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationListId** | **int64** | The ID of the application list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutApplicationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationListRequest** | [**ApplicationListRequest**](ApplicationListRequest.md) | Update the application list. | 

### Return type

[**ApplicationList**](ApplicationList.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUsageApplications

> []ApplicationUsageResponseInner UpdateUsageApplications(ctx).AttributeName(attributeName).AttributeValue(attributeValue).Execute()

Get Usage of Applications



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/rules"
)

func main() {
	attributeName := openapiclient.attributeNameInQuery("umbrella.destination.application_ids") // AttributeNameInQuery | Filter on the name of the rule attribute.
	attributeValue := []int64{int64(123)} // []int64 | Filter on the value of the rule attribute.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationListsAPI.UpdateUsageApplications(context.Background()).AttributeName(attributeName).AttributeValue(attributeValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationListsAPI.UpdateUsageApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUsageApplications`: []ApplicationUsageResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplicationListsAPI.UpdateUsageApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUsageApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attributeName** | [**AttributeNameInQuery**](AttributeNameInQuery.md) | Filter on the name of the rule attribute. | 
 **attributeValue** | **[]int64** | Filter on the value of the rule attribute. | 

### Return type

[**[]ApplicationUsageResponseInner**](ApplicationUsageResponseInner.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

