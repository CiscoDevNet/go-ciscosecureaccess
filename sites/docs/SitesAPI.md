# \SitesAPI

All URIs are relative to *https://api.sse.cisco.com/deployments/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSite**](SitesAPI.md#CreateSite) | **Post** /sites | Create Site
[**DeleteSite**](SitesAPI.md#DeleteSite) | **Delete** /sites/{siteId} | Delete Site
[**GetSite**](SitesAPI.md#GetSite) | **Get** /sites/{siteId} | Get Site
[**ListSites**](SitesAPI.md#ListSites) | **Get** /sites | List Sites
[**UpdateSite**](SitesAPI.md#UpdateSite) | **Put** /sites/{siteId} | Update Site



## CreateSite

> SiteObject CreateSite(ctx).CreateSiteRequest(createSiteRequest).Execute()

Create Site



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/sites"
)

func main() {
	createSiteRequest := *openapiclient.NewCreateSiteRequest("site-one") // CreateSiteRequest | Register a Site. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesAPI.CreateSite(context.Background()).CreateSiteRequest(createSiteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.CreateSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSite`: SiteObject
	fmt.Fprintf(os.Stdout, "Response from `SitesAPI.CreateSite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSiteRequest** | [**CreateSiteRequest**](CreateSiteRequest.md) | Register a Site. | 

### Return type

[**SiteObject**](SiteObject.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSite

> DeleteSite(ctx, siteId).Execute()

Delete Site



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/sites"
)

func main() {
	siteId := int64(13545) // int64 | The ID of the Site.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesAPI.DeleteSite(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.DeleteSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **int64** | The ID of the Site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteRequest struct via the builder pattern


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


## GetSite

> SiteObject GetSite(ctx, siteId).Execute()

Get Site



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/sites"
)

func main() {
	siteId := int64(13545) // int64 | The ID of the Site.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesAPI.GetSite(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.GetSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSite`: SiteObject
	fmt.Fprintf(os.Stdout, "Response from `SitesAPI.GetSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **int64** | The ID of the Site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SiteObject**](SiteObject.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSites

> []SiteObject ListSites(ctx).Page(page).Limit(limit).Execute()

List Sites



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/sites"
)

func main() {
	page := int64(56) // int64 | The number of a page in the collection. (optional) (default to 1)
	limit := int64(56) // int64 | The number of items in the collection to return on the page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesAPI.ListSites(context.Background()).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.ListSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSites`: []SiteObject
	fmt.Fprintf(os.Stdout, "Response from `SitesAPI.ListSites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | The number of a page in the collection. | [default to 1]
 **limit** | **int64** | The number of items in the collection to return on the page. | [default to 100]

### Return type

[**[]SiteObject**](SiteObject.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSite

> SiteObject UpdateSite(ctx, siteId).CreateSiteRequest(createSiteRequest).Execute()

Update Site



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/sites"
)

func main() {
	siteId := int64(13545) // int64 | The ID of the Site.
	createSiteRequest := *openapiclient.NewCreateSiteRequest("site-one") // CreateSiteRequest | Update the properties of the Site. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesAPI.UpdateSite(context.Background(), siteId).CreateSiteRequest(createSiteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.UpdateSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSite`: SiteObject
	fmt.Fprintf(os.Stdout, "Response from `SitesAPI.UpdateSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **int64** | The ID of the Site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSiteRequest** | [**CreateSiteRequest**](CreateSiteRequest.md) | Update the properties of the Site. | 

### Return type

[**SiteObject**](SiteObject.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

