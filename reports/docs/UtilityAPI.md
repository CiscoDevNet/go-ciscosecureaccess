# \UtilityAPI

All URIs are relative to *https://api.sse.cisco.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCategories**](UtilityAPI.md#GetCategories) | **Get** /reports/v2/categories | Get Categories
[**GetIdentities**](UtilityAPI.md#GetIdentities) | **Get** /reports/v2/identities | Get Identities
[**GetIdentity**](UtilityAPI.md#GetIdentity) | **Get** /reports/v2/identities/{identityid} | Get Identity
[**GetThreatName**](UtilityAPI.md#GetThreatName) | **Get** /reports/v2/threat-names/{threatnameid} | Get Threat Name By Threat ID
[**GetThreatNames**](UtilityAPI.md#GetThreatNames) | **Get** /reports/v2/threat-names | Get Threat Names
[**GetThreatType**](UtilityAPI.md#GetThreatType) | **Get** /reports/v2/threat-types/{threattypeid} | Get Threat Type By Threat ID
[**GetThreatTypes**](UtilityAPI.md#GetThreatTypes) | **Get** /reports/v2/threat-types | Get Threat Types
[**PostIdentities**](UtilityAPI.md#PostIdentities) | **Post** /reports/v2/identities | Get Identities By IDs



## GetCategories

> GetCategories200Response GetCategories(ctx).Execute()

Get Categories



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilityAPI.GetCategories(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilityAPI.GetCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCategories`: GetCategories200Response
	fmt.Fprintf(os.Stdout, "Response from `UtilityAPI.GetCategories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoriesRequest struct via the builder pattern


### Return type

[**GetCategories200Response**](GetCategories200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentities

> GetIdentities200Response GetIdentities(ctx).Limit(limit).Offset(offset).Search(search).Identitytypes(identitytypes).Execute()

Get Identities



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
	limit := int64(100) // int64 | (Identities utility endpoint) The number of records to return from the collection. The default limit is 100. In a single response, the server returns at most 5000 records from the collection. (default to 100)
	offset := int64(0) // int64 | A number that represents an index in the collection. (optional) (default to 0)
	search := "somelabel" // string | A string that represents a search parameter. Filter data for requests where the search string appears in the endpoint data. (optional)
	identitytypes := "network,roaming" // string | An identity type or comma-delimited list of identity types. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilityAPI.GetIdentities(context.Background()).Limit(limit).Offset(offset).Search(search).Identitytypes(identitytypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilityAPI.GetIdentities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentities`: GetIdentities200Response
	fmt.Fprintf(os.Stdout, "Response from `UtilityAPI.GetIdentities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | (Identities utility endpoint) The number of records to return from the collection. The default limit is 100. In a single response, the server returns at most 5000 records from the collection. | [default to 100]
 **offset** | **int64** | A number that represents an index in the collection. | [default to 0]
 **search** | **string** | A string that represents a search parameter. Filter data for requests where the search string appears in the endpoint data. | 
 **identitytypes** | **string** | An identity type or comma-delimited list of identity types. | 

### Return type

[**GetIdentities200Response**](GetIdentities200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentity

> GetIdentity200Response GetIdentity(ctx, identityid).Execute()

Get Identity



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
	identityid := int64(42) // int64 | An identity ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilityAPI.GetIdentity(context.Background(), identityid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilityAPI.GetIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentity`: GetIdentity200Response
	fmt.Fprintf(os.Stdout, "Response from `UtilityAPI.GetIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityid** | **int64** | An identity ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetIdentity200Response**](GetIdentity200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThreatName

> GetThreatName200Response GetThreatName(ctx, threatnameid).Execute()

Get Threat Name By Threat ID



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
	threatnameid := "WannaCry" // string | The name of the threat.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilityAPI.GetThreatName(context.Background(), threatnameid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilityAPI.GetThreatName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetThreatName`: GetThreatName200Response
	fmt.Fprintf(os.Stdout, "Response from `UtilityAPI.GetThreatName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threatnameid** | **string** | The name of the threat. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThreatNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetThreatName200Response**](GetThreatName200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThreatNames

> GetThreatNames200Response GetThreatNames(ctx).Execute()

Get Threat Names



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilityAPI.GetThreatNames(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilityAPI.GetThreatNames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetThreatNames`: GetThreatNames200Response
	fmt.Fprintf(os.Stdout, "Response from `UtilityAPI.GetThreatNames`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetThreatNamesRequest struct via the builder pattern


### Return type

[**GetThreatNames200Response**](GetThreatNames200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThreatType

> GetThreatType200Response GetThreatType(ctx, threattypeid).Execute()

Get Threat Type By Threat ID



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
	threattypeid := "Ransomware" // string | The name of the threat type.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilityAPI.GetThreatType(context.Background(), threattypeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilityAPI.GetThreatType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetThreatType`: GetThreatType200Response
	fmt.Fprintf(os.Stdout, "Response from `UtilityAPI.GetThreatType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threattypeid** | **string** | The name of the threat type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThreatTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetThreatType200Response**](GetThreatType200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThreatTypes

> GetThreatTypes200Response GetThreatTypes(ctx).Execute()

Get Threat Types



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilityAPI.GetThreatTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilityAPI.GetThreatTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetThreatTypes`: GetThreatTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `UtilityAPI.GetThreatTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetThreatTypesRequest struct via the builder pattern


### Return type

[**GetThreatTypes200Response**](GetThreatTypes200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentities

> GetIdentities200Response PostIdentities(ctx).Limit(limit).PostIdentitiesRequest(postIdentitiesRequest).Execute()

Get Identities By IDs



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
	limit := int64(100) // int64 | The maximum number of records to return from the collection. (default to 100)
	postIdentitiesRequest := *openapiclient.NewPostIdentitiesRequest() // PostIdentitiesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilityAPI.PostIdentities(context.Background()).Limit(limit).PostIdentitiesRequest(postIdentitiesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilityAPI.PostIdentities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIdentities`: GetIdentities200Response
	fmt.Fprintf(os.Stdout, "Response from `UtilityAPI.PostIdentities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | The maximum number of records to return from the collection. | [default to 100]
 **postIdentitiesRequest** | [**PostIdentitiesRequest**](PostIdentitiesRequest.md) |  | 

### Return type

[**GetIdentities200Response**](GetIdentities200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

