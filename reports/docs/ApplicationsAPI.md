# \ApplicationsAPI

All URIs are relative to *https://api.sse.cisco.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApplication**](ApplicationsAPI.md#GetApplication) | **Get** /reports/v2/appDiscovery/applications/{applicationId} | Get Application
[**GetApplicationAttributes**](ApplicationsAPI.md#GetApplicationAttributes) | **Get** /reports/v2/appDiscovery/applications/{applicationId}/attributes | List Application Attributes
[**GetApplicationIdentities**](ApplicationsAPI.md#GetApplicationIdentities) | **Get** /reports/v2/appDiscovery/applications/{applicationId}/identities | List Application Identities
[**GetApplicationRisk**](ApplicationsAPI.md#GetApplicationRisk) | **Get** /reports/v2/appDiscovery/applications/{applicationId}/risk | Get Application Risk
[**PatchApplications**](ApplicationsAPI.md#PatchApplications) | **Patch** /reports/v2/appDiscovery/applications | Update Applications
[**UpdateApplication**](ApplicationsAPI.md#UpdateApplication) | **Patch** /reports/v2/appDiscovery/applications/{applicationId} | Update Application



## GetApplication

> ApplicationObject GetApplication(ctx, applicationId).Execute()

Get Application



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
	applicationId := "applicationId_example" // string | The ID of the application.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.GetApplication(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplication`: ApplicationObject
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.GetApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationObject**](ApplicationObject.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationAttributes

> ApplicationAttributeCategoryList GetApplicationAttributes(ctx, applicationId).Categories(categories).Execute()

List Application Attributes



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
	applicationId := "applicationId_example" // string | The ID of the application.
	categories := []string{"Inner_example"} // []string | The attributes' categories to filter the collection. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.GetApplicationAttributes(context.Background(), applicationId).Categories(categories).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetApplicationAttributes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationAttributes`: ApplicationAttributeCategoryList
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.GetApplicationAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **categories** | **[]string** | The attributes&#39; categories to filter the collection. | 

### Return type

[**ApplicationAttributeCategoryList**](ApplicationAttributeCategoryList.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationIdentities

> ApplicationIdentityList GetApplicationIdentities(ctx, applicationId).Date(date).Limit(limit).Offset(offset).Sort(sort).Order(order).Execute()

List Application Identities



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
	applicationId := "applicationId_example" // string | The ID of the application.
	date := time.Now() // string | Specify a date to search for data within a twenty-four hour time period. If you do not provide a date, the last 90 days period is used to query the collection. (optional)
	limit := int64(56) // int64 | The maximum number of items to return in the collection. (optional)
	offset := int64(56) // int64 | The number of items to skip before starting to collect the result set. (optional)
	sort := "lastDetected" // string | Specify the name of a field to sort the application identities. (optional)
	order := "asc" // string | Specify the order to sort the collection. Valid values are: `asc` (ascending) or `desc` (descending). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.GetApplicationIdentities(context.Background(), applicationId).Date(date).Limit(limit).Offset(offset).Sort(sort).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetApplicationIdentities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationIdentities`: ApplicationIdentityList
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.GetApplicationIdentities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **date** | **string** | Specify a date to search for data within a twenty-four hour time period. If you do not provide a date, the last 90 days period is used to query the collection. | 
 **limit** | **int64** | The maximum number of items to return in the collection. | 
 **offset** | **int64** | The number of items to skip before starting to collect the result set. | 
 **sort** | **string** | Specify the name of a field to sort the application identities. | 
 **order** | **string** | Specify the order to sort the collection. Valid values are: &#x60;asc&#x60; (ascending) or &#x60;desc&#x60; (descending). | 

### Return type

[**ApplicationIdentityList**](ApplicationIdentityList.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationRisk

> ApplicationRisk GetApplicationRisk(ctx, applicationId).Execute()

Get Application Risk



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
	applicationId := "applicationId_example" // string | The ID of the application.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.GetApplicationRisk(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetApplicationRisk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationRisk`: ApplicationRisk
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.GetApplicationRisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationRiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationRisk**](ApplicationRisk.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchApplications

> BulkLabelApplications PatchApplications(ctx).PatchApplicationsRequest(patchApplicationsRequest).Execute()

Update Applications



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
	patchApplicationsRequest := *openapiclient.NewPatchApplicationsRequest() // PatchApplicationsRequest | A JSON object containing application information for bulk label update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.PatchApplications(context.Background()).PatchApplicationsRequest(patchApplicationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.PatchApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchApplications`: BulkLabelApplications
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.PatchApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchApplicationsRequest** | [**PatchApplicationsRequest**](PatchApplicationsRequest.md) | A JSON object containing application information for bulk label update. | 

### Return type

[**BulkLabelApplications**](BulkLabelApplications.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplication

> ApplicationObject UpdateApplication(ctx, applicationId).UpdateApplicationRequest(updateApplicationRequest).Execute()

Update Application



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
	applicationId := "applicationId_example" // string | The ID of the application.
	updateApplicationRequest := *openapiclient.NewUpdateApplicationRequest(openapiclient.Label("unreviewed")) // UpdateApplicationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.UpdateApplication(context.Background(), applicationId).UpdateApplicationRequest(updateApplicationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.UpdateApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplication`: ApplicationObject
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.UpdateApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateApplicationRequest** | [**UpdateApplicationRequest**](UpdateApplicationRequest.md) |  | 

### Return type

[**ApplicationObject**](ApplicationObject.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

