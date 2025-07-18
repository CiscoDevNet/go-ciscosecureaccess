# \TopDestinationsAPI

All URIs are relative to *https://api.sse.cisco.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTopDestinations**](TopDestinationsAPI.md#GetTopDestinations) | **Get** /reports/v2/top-destinations | Get Top Destinations
[**GetTopDestinationsType**](TopDestinationsAPI.md#GetTopDestinationsType) | **Get** /reports/v2/top-destinations/{type} | Get Top Destinations By Type
[**GetTopUrls**](TopDestinationsAPI.md#GetTopUrls) | **Get** /reports/v2/top-urls | Get Top URLs



## GetTopDestinations

> GetTopDestinations200Response GetTopDestinations(ctx).From(from).To(to).Limit(limit).Offset(offset).Domains(domains).Urls(urls).Categories(categories).Policycategories(policycategories).Ip(ip).Ports(ports).Identityids(identityids).Identitytypes(identitytypes).Applicationid(applicationid).Verdict(verdict).Sha256(sha256).Securityoverridden(securityoverridden).Bundleid(bundleid).Threats(threats).Threattypes(threattypes).Ampdisposition(ampdisposition).Antivirusthreats(antivirusthreats).Datalosspreventionstate(datalosspreventionstate).Filternoisydomains(filternoisydomains).Timezone(timezone).Execute()

Get Top Destinations



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
	limit := int64(100) // int64 | The maximum number of records to return from the collection. (default to 100)
	offset := int64(0) // int64 | A number that represents an index in the collection. (optional) (default to 0)
	domains := "cisco.com,nasa.gov" // string | A domain name or comma-delimited list of domain name. (optional)
	urls := "https://google.com,facebook.com/help" // string | A URL or comma-delimited list of URL. (optional)
	categories := "148,151,66" // string | A category ID or comma-delimited list of category ID. (optional)
	policycategories := "67,69" // string | A category ID or comma-delimited list of category ID. Filter the request by the categories that trigger a policy. (optional)
	ip := "10.10.10.10" // string | An IP address. (optional)
	ports := "7351,80" // string | A port number or comma-delimited list of port numbers. (optional)
	identityids := "1,2,3" // string | An identity ID or comma-delimited list of identity IDs. (optional)
	identitytypes := "network,roaming" // string | An identity type or comma-delimited list of identity types. (optional)
	applicationid := "1" // string | The ID of the application. (optional)
	verdict := "allowed,blocked" // string | A string or comma-delimited string that describes whether the traffic can reach the destination. (optional)
	sha256 := "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad" // string | A SHA-256 hash. (optional)
	securityoverridden := true // bool | Specify whether to filter on requests that override security. (optional)
	bundleid := int64(1) // int64 | A proxy bundle ID. (optional)
	threats := "threats_example" // string | A threat name or comma-delimited list of threat names. (optional)
	threattypes := "threattypes_example" // string | A threat type or comma-delimited list of threat types. (optional)
	ampdisposition := "clean,malicious,unknown" // string | An AMP disposition string or a comma-delimited list of AMP disposition strings. (optional)
	antivirusthreats := "Trojan.Linux.Generic.144075" // string | A threat name or comma-delimited list of threat names. (optional)
	datalosspreventionstate := "blocked" // string | A string that describes the status of a destination. Filter for requests that are blocked by the DLP layer security. (optional)
	filternoisydomains := true // bool | Filter out domains that generate a lot of insignificant traffic (noise). (optional)
	timezone := "ASIA%2fCALCUTTA" // string | Display the timestamp of the traffic events in the specified timezone. For the timezone, provide a continent and city separated by an url-encoded forward slash ('/'), for example: timezone='ASIA%2fCALCUTTA'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TopDestinationsAPI.GetTopDestinations(context.Background()).From(from).To(to).Limit(limit).Offset(offset).Domains(domains).Urls(urls).Categories(categories).Policycategories(policycategories).Ip(ip).Ports(ports).Identityids(identityids).Identitytypes(identitytypes).Applicationid(applicationid).Verdict(verdict).Sha256(sha256).Securityoverridden(securityoverridden).Bundleid(bundleid).Threats(threats).Threattypes(threattypes).Ampdisposition(ampdisposition).Antivirusthreats(antivirusthreats).Datalosspreventionstate(datalosspreventionstate).Filternoisydomains(filternoisydomains).Timezone(timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopDestinationsAPI.GetTopDestinations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTopDestinations`: GetTopDestinations200Response
	fmt.Fprintf(os.Stdout, "Response from `TopDestinationsAPI.GetTopDestinations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTopDestinationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | A timestamp or relative time string (for example: &#39;-1days&#39;). Filter for data that appears after this time. | 
 **to** | **string** | A timestamp or relative time string (for example: &#39;now&#39;). Filter for data that appears before this time. | 
 **limit** | **int64** | The maximum number of records to return from the collection. | [default to 100]
 **offset** | **int64** | A number that represents an index in the collection. | [default to 0]
 **domains** | **string** | A domain name or comma-delimited list of domain name. | 
 **urls** | **string** | A URL or comma-delimited list of URL. | 
 **categories** | **string** | A category ID or comma-delimited list of category ID. | 
 **policycategories** | **string** | A category ID or comma-delimited list of category ID. Filter the request by the categories that trigger a policy. | 
 **ip** | **string** | An IP address. | 
 **ports** | **string** | A port number or comma-delimited list of port numbers. | 
 **identityids** | **string** | An identity ID or comma-delimited list of identity IDs. | 
 **identitytypes** | **string** | An identity type or comma-delimited list of identity types. | 
 **applicationid** | **string** | The ID of the application. | 
 **verdict** | **string** | A string or comma-delimited string that describes whether the traffic can reach the destination. | 
 **sha256** | **string** | A SHA-256 hash. | 
 **securityoverridden** | **bool** | Specify whether to filter on requests that override security. | 
 **bundleid** | **int64** | A proxy bundle ID. | 
 **threats** | **string** | A threat name or comma-delimited list of threat names. | 
 **threattypes** | **string** | A threat type or comma-delimited list of threat types. | 
 **ampdisposition** | **string** | An AMP disposition string or a comma-delimited list of AMP disposition strings. | 
 **antivirusthreats** | **string** | A threat name or comma-delimited list of threat names. | 
 **datalosspreventionstate** | **string** | A string that describes the status of a destination. Filter for requests that are blocked by the DLP layer security. | 
 **filternoisydomains** | **bool** | Filter out domains that generate a lot of insignificant traffic (noise). | 
 **timezone** | **string** | Display the timestamp of the traffic events in the specified timezone. For the timezone, provide a continent and city separated by an url-encoded forward slash (&#39;/&#39;), for example: timezone&#x3D;&#39;ASIA%2fCALCUTTA&#39;. | 

### Return type

[**GetTopDestinations200Response**](GetTopDestinations200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopDestinationsType

> GetTopDestinations200Response GetTopDestinationsType(ctx, type_).From(from).To(to).Limit(limit).Offset(offset).Domains(domains).Urls(urls).Categories(categories).Policycategories(policycategories).Ip(ip).Ports(ports).Identityids(identityids).Identitytypes(identitytypes).Applicationid(applicationid).Verdict(verdict).Sha256(sha256).Securityoverridden(securityoverridden).Bundleid(bundleid).Threats(threats).Threattypes(threattypes).Ampdisposition(ampdisposition).Antivirusthreats(antivirusthreats).Datalosspreventionstate(datalosspreventionstate).Filternoisydomains(filternoisydomains).Timezone(timezone).Execute()

Get Top Destinations By Type



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
	type_ := "firewall" // string | Specify the type of traffic.
	from := "1639146300000" // string | A timestamp or relative time string (for example: '-1days'). Filter for data that appears after this time.
	to := "1640010300000" // string | A timestamp or relative time string (for example: 'now'). Filter for data that appears before this time.
	limit := int64(100) // int64 | The maximum number of records to return from the collection. (default to 100)
	offset := int64(0) // int64 | A number that represents an index in the collection. (optional) (default to 0)
	domains := "cisco.com,nasa.gov" // string | A domain name or comma-delimited list of domain name. (optional)
	urls := "https://google.com,facebook.com/help" // string | A URL or comma-delimited list of URL. (optional)
	categories := "148,151,66" // string | A category ID or comma-delimited list of category ID. (optional)
	policycategories := "67,69" // string | A category ID or comma-delimited list of category ID. Filter the request by the categories that trigger a policy. (optional)
	ip := "10.10.10.10" // string | An IP address. (optional)
	ports := "7351,80" // string | A port number or comma-delimited list of port numbers. (optional)
	identityids := "1,2,3" // string | An identity ID or comma-delimited list of identity IDs. (optional)
	identitytypes := "network,roaming" // string | An identity type or comma-delimited list of identity types. (optional)
	applicationid := "1" // string | The ID of the application. (optional)
	verdict := "allowed,blocked" // string | A string or comma-delimited string that describes whether the traffic can reach the destination. (optional)
	sha256 := "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad" // string | A SHA-256 hash. (optional)
	securityoverridden := true // bool | Specify whether to filter on requests that override security. (optional)
	bundleid := int64(1) // int64 | A proxy bundle ID. (optional)
	threats := "threats_example" // string | A threat name or comma-delimited list of threat names. (optional)
	threattypes := "threattypes_example" // string | A threat type or comma-delimited list of threat types. (optional)
	ampdisposition := "clean,malicious,unknown" // string | An AMP disposition string or a comma-delimited list of AMP disposition strings. (optional)
	antivirusthreats := "Trojan.Linux.Generic.144075" // string | A threat name or comma-delimited list of threat names. (optional)
	datalosspreventionstate := "blocked" // string | A string that describes the status of a destination. Filter for requests that are blocked by the DLP layer security. (optional)
	filternoisydomains := true // bool | Filter out domains that generate a lot of insignificant traffic (noise). (optional)
	timezone := "ASIA%2fCALCUTTA" // string | Display the timestamp of the traffic events in the specified timezone. For the timezone, provide a continent and city separated by an url-encoded forward slash ('/'), for example: timezone='ASIA%2fCALCUTTA'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TopDestinationsAPI.GetTopDestinationsType(context.Background(), type_).From(from).To(to).Limit(limit).Offset(offset).Domains(domains).Urls(urls).Categories(categories).Policycategories(policycategories).Ip(ip).Ports(ports).Identityids(identityids).Identitytypes(identitytypes).Applicationid(applicationid).Verdict(verdict).Sha256(sha256).Securityoverridden(securityoverridden).Bundleid(bundleid).Threats(threats).Threattypes(threattypes).Ampdisposition(ampdisposition).Antivirusthreats(antivirusthreats).Datalosspreventionstate(datalosspreventionstate).Filternoisydomains(filternoisydomains).Timezone(timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopDestinationsAPI.GetTopDestinationsType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTopDestinationsType`: GetTopDestinations200Response
	fmt.Fprintf(os.Stdout, "Response from `TopDestinationsAPI.GetTopDestinationsType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | Specify the type of traffic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopDestinationsTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | A timestamp or relative time string (for example: &#39;-1days&#39;). Filter for data that appears after this time. | 
 **to** | **string** | A timestamp or relative time string (for example: &#39;now&#39;). Filter for data that appears before this time. | 
 **limit** | **int64** | The maximum number of records to return from the collection. | [default to 100]
 **offset** | **int64** | A number that represents an index in the collection. | [default to 0]
 **domains** | **string** | A domain name or comma-delimited list of domain name. | 
 **urls** | **string** | A URL or comma-delimited list of URL. | 
 **categories** | **string** | A category ID or comma-delimited list of category ID. | 
 **policycategories** | **string** | A category ID or comma-delimited list of category ID. Filter the request by the categories that trigger a policy. | 
 **ip** | **string** | An IP address. | 
 **ports** | **string** | A port number or comma-delimited list of port numbers. | 
 **identityids** | **string** | An identity ID or comma-delimited list of identity IDs. | 
 **identitytypes** | **string** | An identity type or comma-delimited list of identity types. | 
 **applicationid** | **string** | The ID of the application. | 
 **verdict** | **string** | A string or comma-delimited string that describes whether the traffic can reach the destination. | 
 **sha256** | **string** | A SHA-256 hash. | 
 **securityoverridden** | **bool** | Specify whether to filter on requests that override security. | 
 **bundleid** | **int64** | A proxy bundle ID. | 
 **threats** | **string** | A threat name or comma-delimited list of threat names. | 
 **threattypes** | **string** | A threat type or comma-delimited list of threat types. | 
 **ampdisposition** | **string** | An AMP disposition string or a comma-delimited list of AMP disposition strings. | 
 **antivirusthreats** | **string** | A threat name or comma-delimited list of threat names. | 
 **datalosspreventionstate** | **string** | A string that describes the status of a destination. Filter for requests that are blocked by the DLP layer security. | 
 **filternoisydomains** | **bool** | Filter out domains that generate a lot of insignificant traffic (noise). | 
 **timezone** | **string** | Display the timestamp of the traffic events in the specified timezone. For the timezone, provide a continent and city separated by an url-encoded forward slash (&#39;/&#39;), for example: timezone&#x3D;&#39;ASIA%2fCALCUTTA&#39;. | 

### Return type

[**GetTopDestinations200Response**](GetTopDestinations200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopUrls

> GetTopUrls200Response GetTopUrls(ctx).From(from).To(to).Limit(limit).Domains(domains).Offset(offset).Sha256(sha256).Securityoverridden(securityoverridden).Bundleid(bundleid).Ampdisposition(ampdisposition).Antivirusthreats(antivirusthreats).Datalosspreventionstate(datalosspreventionstate).Filternoisydomains(filternoisydomains).Timezone(timezone).Execute()

Get Top URLs



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
	limit := int64(100) // int64 | The maximum number of records to return from the collection. (default to 100)
	domains := "cisco.com,nasa.gov" // string | A domain name or comma-delimited list of domain name. (optional)
	offset := int64(0) // int64 | A number that represents an index in the collection. (optional) (default to 0)
	sha256 := "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad" // string | A SHA-256 hash. (optional)
	securityoverridden := true // bool | Specify whether to filter on requests that override security. (optional)
	bundleid := int64(1) // int64 | A proxy bundle ID. (optional)
	ampdisposition := "clean,malicious,unknown" // string | An AMP disposition string or a comma-delimited list of AMP disposition strings. (optional)
	antivirusthreats := "Trojan.Linux.Generic.144075" // string | A threat name or comma-delimited list of threat names. (optional)
	datalosspreventionstate := "blocked" // string | A string that describes the status of a destination. Filter for requests that are blocked by the DLP layer security. (optional)
	filternoisydomains := true // bool | Filter out domains that generate a lot of insignificant traffic (noise). (optional)
	timezone := "ASIA%2fCALCUTTA" // string | Display the timestamp of the traffic events in the specified timezone. For the timezone, provide a continent and city separated by an url-encoded forward slash ('/'), for example: timezone='ASIA%2fCALCUTTA'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TopDestinationsAPI.GetTopUrls(context.Background()).From(from).To(to).Limit(limit).Domains(domains).Offset(offset).Sha256(sha256).Securityoverridden(securityoverridden).Bundleid(bundleid).Ampdisposition(ampdisposition).Antivirusthreats(antivirusthreats).Datalosspreventionstate(datalosspreventionstate).Filternoisydomains(filternoisydomains).Timezone(timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopDestinationsAPI.GetTopUrls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTopUrls`: GetTopUrls200Response
	fmt.Fprintf(os.Stdout, "Response from `TopDestinationsAPI.GetTopUrls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTopUrlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | A timestamp or relative time string (for example: &#39;-1days&#39;). Filter for data that appears after this time. | 
 **to** | **string** | A timestamp or relative time string (for example: &#39;now&#39;). Filter for data that appears before this time. | 
 **limit** | **int64** | The maximum number of records to return from the collection. | [default to 100]
 **domains** | **string** | A domain name or comma-delimited list of domain name. | 
 **offset** | **int64** | A number that represents an index in the collection. | [default to 0]
 **sha256** | **string** | A SHA-256 hash. | 
 **securityoverridden** | **bool** | Specify whether to filter on requests that override security. | 
 **bundleid** | **int64** | A proxy bundle ID. | 
 **ampdisposition** | **string** | An AMP disposition string or a comma-delimited list of AMP disposition strings. | 
 **antivirusthreats** | **string** | A threat name or comma-delimited list of threat names. | 
 **datalosspreventionstate** | **string** | A string that describes the status of a destination. Filter for requests that are blocked by the DLP layer security. | 
 **filternoisydomains** | **bool** | Filter out domains that generate a lot of insignificant traffic (noise). | 
 **timezone** | **string** | Display the timestamp of the traffic events in the specified timezone. For the timezone, provide a continent and city separated by an url-encoded forward slash (&#39;/&#39;), for example: timezone&#x3D;&#39;ASIA%2fCALCUTTA&#39;. | 

### Return type

[**GetTopUrls200Response**](GetTopUrls200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

