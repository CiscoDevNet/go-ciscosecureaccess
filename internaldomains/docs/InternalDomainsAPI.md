# \InternalDomainsAPI

All URIs are relative to *https://api.sse.cisco.com/deployments/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInternalDomain**](InternalDomainsAPI.md#CreateInternalDomain) | **Post** /internaldomains | Create Internal Domain
[**DeleteInternalDomain**](InternalDomainsAPI.md#DeleteInternalDomain) | **Delete** /internaldomains/{internalDomainId} | Delete Internal Domain
[**GetInternalDomain**](InternalDomainsAPI.md#GetInternalDomain) | **Get** /internaldomains/{internalDomainId} | Get Internal Domain
[**ListInternalDomains**](InternalDomainsAPI.md#ListInternalDomains) | **Get** /internaldomains | List Internal Domains
[**UpdateInternalDomain**](InternalDomainsAPI.md#UpdateInternalDomain) | **Put** /internaldomains/{internalDomainId} | Update Internal Domain



## CreateInternalDomain

> InternalDomainObject CreateInternalDomain(ctx).CreateInternalDomainRequest(createInternalDomainRequest).Execute()

Create Internal Domain



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/internaldomains"
)

func main() {
	createInternalDomainRequest := *openapiclient.NewCreateInternalDomainRequest("cisco-internal.com") // CreateInternalDomainRequest | Create the internal domain.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternalDomainsAPI.CreateInternalDomain(context.Background()).CreateInternalDomainRequest(createInternalDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalDomainsAPI.CreateInternalDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInternalDomain`: InternalDomainObject
	fmt.Fprintf(os.Stdout, "Response from `InternalDomainsAPI.CreateInternalDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInternalDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInternalDomainRequest** | [**CreateInternalDomainRequest**](CreateInternalDomainRequest.md) | Create the internal domain. | 

### Return type

[**InternalDomainObject**](InternalDomainObject.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInternalDomain

> string DeleteInternalDomain(ctx, internalDomainId).Execute()

Delete Internal Domain



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/internaldomains"
)

func main() {
	internalDomainId := int64(12456) // int64 | The ID of the internal domain.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternalDomainsAPI.DeleteInternalDomain(context.Background(), internalDomainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalDomainsAPI.DeleteInternalDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteInternalDomain`: string
	fmt.Fprintf(os.Stdout, "Response from `InternalDomainsAPI.DeleteInternalDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**internalDomainId** | **int64** | The ID of the internal domain. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInternalDomainRequest struct via the builder pattern


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


## GetInternalDomain

> InternalDomainObject GetInternalDomain(ctx, internalDomainId).Execute()

Get Internal Domain



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/internaldomains"
)

func main() {
	internalDomainId := int64(12456) // int64 | The ID of the internal domain.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternalDomainsAPI.GetInternalDomain(context.Background(), internalDomainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalDomainsAPI.GetInternalDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInternalDomain`: InternalDomainObject
	fmt.Fprintf(os.Stdout, "Response from `InternalDomainsAPI.GetInternalDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**internalDomainId** | **int64** | The ID of the internal domain. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInternalDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InternalDomainObject**](InternalDomainObject.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInternalDomains

> []InternalDomainObject ListInternalDomains(ctx).Page(page).Limit(limit).Execute()

List Internal Domains



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/internaldomains"
)

func main() {
	page := int64(2) // int64 | The number of a page in the collection. (optional) (default to 1)
	limit := int64(50) // int64 | The number of records in the collection to return on the page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternalDomainsAPI.ListInternalDomains(context.Background()).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalDomainsAPI.ListInternalDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInternalDomains`: []InternalDomainObject
	fmt.Fprintf(os.Stdout, "Response from `InternalDomainsAPI.ListInternalDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInternalDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | The number of a page in the collection. | [default to 1]
 **limit** | **int64** | The number of records in the collection to return on the page. | [default to 100]

### Return type

[**[]InternalDomainObject**](InternalDomainObject.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInternalDomain

> InternalDomainObject UpdateInternalDomain(ctx, internalDomainId).CreateInternalDomainRequest(createInternalDomainRequest).Execute()

Update Internal Domain



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/internaldomains"
)

func main() {
	internalDomainId := int64(12456) // int64 | The ID of the internal domain.
	createInternalDomainRequest := *openapiclient.NewCreateInternalDomainRequest("cisco-internal.com") // CreateInternalDomainRequest | Create the internal domain.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternalDomainsAPI.UpdateInternalDomain(context.Background(), internalDomainId).CreateInternalDomainRequest(createInternalDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalDomainsAPI.UpdateInternalDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInternalDomain`: InternalDomainObject
	fmt.Fprintf(os.Stdout, "Response from `InternalDomainsAPI.UpdateInternalDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**internalDomainId** | **int64** | The ID of the internal domain. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInternalDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createInternalDomainRequest** | [**CreateInternalDomainRequest**](CreateInternalDomainRequest.md) | Create the internal domain. | 

### Return type

[**InternalDomainObject**](InternalDomainObject.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

