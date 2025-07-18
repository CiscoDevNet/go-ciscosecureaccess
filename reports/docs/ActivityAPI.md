# \ActivityAPI

All URIs are relative to *https://api.sse.cisco.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActivities**](ActivityAPI.md#GetActivities) | **Get** /reports/v2/activity | Get Activities (all)
[**GetActivityAmpRetrospective**](ActivityAPI.md#GetActivityAmpRetrospective) | **Get** /reports/v2/activity/amp-retrospective | Get Activity AMP Retrospective
[**GetActivityDecryption**](ActivityAPI.md#GetActivityDecryption) | **Get** /reports/v2/activity/decryption | Activity Decryption
[**GetActivityDns**](ActivityAPI.md#GetActivityDns) | **Get** /reports/v2/activity/dns | Get Activity DNS
[**GetActivityFirewall**](ActivityAPI.md#GetActivityFirewall) | **Get** /reports/v2/activity/firewall | Get Activity Firewall
[**GetActivityIP**](ActivityAPI.md#GetActivityIP) | **Get** /reports/v2/activity/ip | Get Activity IP
[**GetActivityIntrusion**](ActivityAPI.md#GetActivityIntrusion) | **Get** /reports/v2/activity/intrusion | Get Activity Intrusion
[**GetActivityProxy**](ActivityAPI.md#GetActivityProxy) | **Get** /reports/v2/activity/proxy | Get Activity Proxy
[**GetActivityZTNA**](ActivityAPI.md#GetActivityZTNA) | **Get** /reports/v2/activity/ztna | Activity ZTNA



## GetActivities

> GetActivities200Response GetActivities(ctx).From(from).To(to).Limit(limit).Offset(offset).Domains(domains).Urls(urls).Categories(categories).Policycategories(policycategories).Ip(ip).Ports(ports).Identityids(identityids).Identitytypes(identitytypes).Applicationid(applicationid).Verdict(verdict).Ruleid(ruleid).Filename(filename).Securityoverridden(securityoverridden).Bundleid(bundleid).Threats(threats).Threattypes(threattypes).Ampdisposition(ampdisposition).Antivirusthreats(antivirusthreats).XTrafficType(xTrafficType).Isolatedstate(isolatedstate).IsolatedFileAction(isolatedFileAction).Datalosspreventionstate(datalosspreventionstate).Filternoisydomains(filternoisydomains).Httperrors(httperrors).Exists(exists).Timezone(timezone).Execute()

Get Activities (all)



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
	ruleid := int64(1) // int64 | The firewall policy rule ID. (optional)
	filename := "myfilename_*" // string | A string that identifies a filename. Filter the request by the filename. Supports globbing or use of the wildcard character ('*'). The asterisk (*) matches zero or more occurrences of any character. (optional)
	securityoverridden := true // bool | Specify whether to filter on requests that override security. (optional)
	bundleid := int64(1) // int64 | A proxy bundle ID. (optional)
	threats := "threats_example" // string | A threat name or comma-delimited list of threat names. (optional)
	threattypes := "threattypes_example" // string | A threat type or comma-delimited list of threat types. (optional)
	ampdisposition := "clean,malicious,unknown" // string | An AMP disposition string or a comma-delimited list of AMP disposition strings. (optional)
	antivirusthreats := "Trojan.Linux.Generic.144075" // string | A threat name or comma-delimited list of threat names. (optional)
	xTrafficType := "dns,proxy,firewall,ip" // string | A string or comma-delimited list of strings that describes the type of traffic. If the header is not set, the default value is `all`. Valid values are: `dns`, `proxy`, `firewall`, and `ip`. (optional)
	isolatedstate := "isolated" // string | A string that describes the remote browser isolation (RBI) isolation type. (optional)
	isolatedFileAction := "downloaded-safe-pdf" // string | A string that describes the remote browser isolation (RBI) file action type. (optional)
	datalosspreventionstate := "blocked" // string | A string that describes the status of a destination. Filter for requests that are blocked by the DLP layer security. (optional)
	filternoisydomains := true // bool | Filter out domains that generate a lot of insignificant traffic (noise). (optional)
	httperrors := "certificateerror" // string | Filter data for requests that resulted in a TLS error or a certificate error. (optional)
	exists := "destinationlistids,threattypes" // string | Specify an attribute or comma-separated list of attributes to filter the data. Valid values are: `categories`, `policycategories`, `applicationid`, `nbarapplicationid`, `nbarapplicationtypeids`, `privateapplicationid`, `applicationgroupids`, `sha256`, `filename`, `threats`, `threattypes`, `antivirusthreats`, `destinationlistids`, and `httperrors`. (optional)
	timezone := "ASIA%2fCALCUTTA" // string | Display the timestamp of the traffic events in the specified timezone. For the timezone, provide a continent and city separated by an url-encoded forward slash ('/'), for example: timezone='ASIA%2fCALCUTTA'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityAPI.GetActivities(context.Background()).From(from).To(to).Limit(limit).Offset(offset).Domains(domains).Urls(urls).Categories(categories).Policycategories(policycategories).Ip(ip).Ports(ports).Identityids(identityids).Identitytypes(identitytypes).Applicationid(applicationid).Verdict(verdict).Ruleid(ruleid).Filename(filename).Securityoverridden(securityoverridden).Bundleid(bundleid).Threats(threats).Threattypes(threattypes).Ampdisposition(ampdisposition).Antivirusthreats(antivirusthreats).XTrafficType(xTrafficType).Isolatedstate(isolatedstate).IsolatedFileAction(isolatedFileAction).Datalosspreventionstate(datalosspreventionstate).Filternoisydomains(filternoisydomains).Httperrors(httperrors).Exists(exists).Timezone(timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityAPI.GetActivities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivities`: GetActivities200Response
	fmt.Fprintf(os.Stdout, "Response from `ActivityAPI.GetActivities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActivitiesRequest struct via the builder pattern


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
 **ruleid** | **int64** | The firewall policy rule ID. | 
 **filename** | **string** | A string that identifies a filename. Filter the request by the filename. Supports globbing or use of the wildcard character (&#39;*&#39;). The asterisk (*) matches zero or more occurrences of any character. | 
 **securityoverridden** | **bool** | Specify whether to filter on requests that override security. | 
 **bundleid** | **int64** | A proxy bundle ID. | 
 **threats** | **string** | A threat name or comma-delimited list of threat names. | 
 **threattypes** | **string** | A threat type or comma-delimited list of threat types. | 
 **ampdisposition** | **string** | An AMP disposition string or a comma-delimited list of AMP disposition strings. | 
 **antivirusthreats** | **string** | A threat name or comma-delimited list of threat names. | 
 **xTrafficType** | **string** | A string or comma-delimited list of strings that describes the type of traffic. If the header is not set, the default value is &#x60;all&#x60;. Valid values are: &#x60;dns&#x60;, &#x60;proxy&#x60;, &#x60;firewall&#x60;, and &#x60;ip&#x60;. | 
 **isolatedstate** | **string** | A string that describes the remote browser isolation (RBI) isolation type. | 
 **isolatedFileAction** | **string** | A string that describes the remote browser isolation (RBI) file action type. | 
 **datalosspreventionstate** | **string** | A string that describes the status of a destination. Filter for requests that are blocked by the DLP layer security. | 
 **filternoisydomains** | **bool** | Filter out domains that generate a lot of insignificant traffic (noise). | 
 **httperrors** | **string** | Filter data for requests that resulted in a TLS error or a certificate error. | 
 **exists** | **string** | Specify an attribute or comma-separated list of attributes to filter the data. Valid values are: &#x60;categories&#x60;, &#x60;policycategories&#x60;, &#x60;applicationid&#x60;, &#x60;nbarapplicationid&#x60;, &#x60;nbarapplicationtypeids&#x60;, &#x60;privateapplicationid&#x60;, &#x60;applicationgroupids&#x60;, &#x60;sha256&#x60;, &#x60;filename&#x60;, &#x60;threats&#x60;, &#x60;threattypes&#x60;, &#x60;antivirusthreats&#x60;, &#x60;destinationlistids&#x60;, and &#x60;httperrors&#x60;. | 
 **timezone** | **string** | Display the timestamp of the traffic events in the specified timezone. For the timezone, provide a continent and city separated by an url-encoded forward slash (&#39;/&#39;), for example: timezone&#x3D;&#39;ASIA%2fCALCUTTA&#39;. | 

### Return type

[**GetActivities200Response**](GetActivities200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActivityAmpRetrospective

> GetActivityAmpRetrospective200Response GetActivityAmpRetrospective(ctx).From(from).To(to).Limit(limit).Offset(offset).Ampdisposition(ampdisposition).Sha256(sha256).Timezone(timezone).Execute()

Get Activity AMP Retrospective



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
	ampdisposition := "clean,malicious,unknown" // string | An AMP disposition string or a comma-delimited list of AMP disposition strings. (optional)
	sha256 := "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad" // string | A SHA-256 hash. (optional)
	timezone := "ASIA%2fCALCUTTA" // string | Display the timestamp of the traffic events in the specified timezone. For the timezone, provide a continent and city separated by an url-encoded forward slash ('/'), for example: timezone='ASIA%2fCALCUTTA'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityAPI.GetActivityAmpRetrospective(context.Background()).From(from).To(to).Limit(limit).Offset(offset).Ampdisposition(ampdisposition).Sha256(sha256).Timezone(timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityAPI.GetActivityAmpRetrospective``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivityAmpRetrospective`: GetActivityAmpRetrospective200Response
	fmt.Fprintf(os.Stdout, "Response from `ActivityAPI.GetActivityAmpRetrospective`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActivityAmpRetrospectiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | A timestamp or relative time string (for example: &#39;-1days&#39;). Filter for data that appears after this time. | 
 **to** | **string** | A timestamp or relative time string (for example: &#39;now&#39;). Filter for data that appears before this time. | 
 **limit** | **int64** | The maximum number of records to return from the collection. | [default to 100]
 **offset** | **int64** | A number that represents an index in the collection. | [default to 0]
 **ampdisposition** | **string** | An AMP disposition string or a comma-delimited list of AMP disposition strings. | 
 **sha256** | **string** | A SHA-256 hash. | 
 **timezone** | **string** | Display the timestamp of the traffic events in the specified timezone. For the timezone, provide a continent and city separated by an url-encoded forward slash (&#39;/&#39;), for example: timezone&#x3D;&#39;ASIA%2fCALCUTTA&#39;. | 

### Return type

[**GetActivityAmpRetrospective200Response**](GetActivityAmpRetrospective200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActivityDecryption

> GetActivityDecryption200Response GetActivityDecryption(ctx).From(from).To(to).Limit(limit).Offset(offset).Decryptaction(decryptaction).Timezone(timezone).Execute()

Activity Decryption



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
	decryptaction := "decryptinbound,donodecrypt" // string | The list of comma-separated decryption actions. Valid values are: `decryptinbound`, `decryptoutbound`, `donotdecrypt`, and `decrypterror`. (optional)
	timezone := "ASIA%2fCALCUTTA" // string | Display the timestamp of the traffic events in the specified timezone. For the timezone, provide a continent and city separated by an url-encoded forward slash ('/'), for example: timezone='ASIA%2fCALCUTTA'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityAPI.GetActivityDecryption(context.Background()).From(from).To(to).Limit(limit).Offset(offset).Decryptaction(decryptaction).Timezone(timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityAPI.GetActivityDecryption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivityDecryption`: GetActivityDecryption200Response
	fmt.Fprintf(os.Stdout, "Response from `ActivityAPI.GetActivityDecryption`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActivityDecryptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | A timestamp or relative time string (for example: &#39;-1days&#39;). Filter for data that appears after this time. | 
 **to** | **string** | A timestamp or relative time string (for example: &#39;now&#39;). Filter for data that appears before this time. | 
 **limit** | **int64** | The maximum number of records to return from the collection. | [default to 100]
 **offset** | **int64** | A number that represents an index in the collection. | [default to 0]
 **decryptaction** | **string** | The list of comma-separated decryption actions. Valid values are: &#x60;decryptinbound&#x60;, &#x60;decryptoutbound&#x60;, &#x60;donotdecrypt&#x60;, and &#x60;decrypterror&#x60;. | 
 **timezone** | **string** | Display the timestamp of the traffic events in the specified timezone. For the timezone, provide a continent and city separated by an url-encoded forward slash (&#39;/&#39;), for example: timezone&#x3D;&#39;ASIA%2fCALCUTTA&#39;. | 

### Return type

[**GetActivityDecryption200Response**](GetActivityDecryption200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActivityDns

> GetActivityDns200Response GetActivityDns(ctx).From(from).To(to).Limit(limit).Offset(offset).Order(order).Domains(domains).Categories(categories).Policycategories(policycategories).Ip(ip).Identityids(identityids).Identitytypes(identitytypes).Applicationid(applicationid).Verdict(verdict).Threats(threats).Threattypes(threattypes).Filternoisydomains(filternoisydomains).Timezone(timezone).Execute()

Get Activity DNS



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
	order := "desc" // string | A string that describes how to order the results: ascending (`asc`) or descending (`desc`). (optional)
	domains := "cisco.com,nasa.gov" // string | A domain name or comma-delimited list of domain name. (optional)
	categories := "148,151,66" // string | A category ID or comma-delimited list of category ID. (optional)
	policycategories := "67,69" // string | A category ID or comma-delimited list of category ID. Filter the request by the categories that trigger a policy. (optional)
	ip := "10.10.10.10" // string | An IP address. (optional)
	identityids := "1,2,3" // string | An identity ID or comma-delimited list of identity IDs. (optional)
	identitytypes := "network,roaming" // string | An identity type or comma-delimited list of identity types. (optional)
	applicationid := "1" // string | The ID of the application. (optional)
	verdict := "allowed,blocked" // string | A string or comma-delimited string that describes whether the traffic can reach the destination. (optional)
	threats := "threats_example" // string | A threat name or comma-delimited list of threat names. (optional)
	threattypes := "threattypes_example" // string | A threat type or comma-delimited list of threat types. (optional)
	filternoisydomains := true // bool | Filter out domains that generate a lot of insignificant traffic (noise). (optional)
	timezone := "ASIA%2fCALCUTTA" // string | Display the timestamp of the traffic events in the specified timezone. For the timezone, provide a continent and city separated by an url-encoded forward slash ('/'), for example: timezone='ASIA%2fCALCUTTA'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityAPI.GetActivityDns(context.Background()).From(from).To(to).Limit(limit).Offset(offset).Order(order).Domains(domains).Categories(categories).Policycategories(policycategories).Ip(ip).Identityids(identityids).Identitytypes(identitytypes).Applicationid(applicationid).Verdict(verdict).Threats(threats).Threattypes(threattypes).Filternoisydomains(filternoisydomains).Timezone(timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityAPI.GetActivityDns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivityDns`: GetActivityDns200Response
	fmt.Fprintf(os.Stdout, "Response from `ActivityAPI.GetActivityDns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActivityDnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | A timestamp or relative time string (for example: &#39;-1days&#39;). Filter for data that appears after this time. | 
 **to** | **string** | A timestamp or relative time string (for example: &#39;now&#39;). Filter for data that appears before this time. | 
 **limit** | **int64** | The maximum number of records to return from the collection. | [default to 100]
 **offset** | **int64** | A number that represents an index in the collection. | [default to 0]
 **order** | **string** | A string that describes how to order the results: ascending (&#x60;asc&#x60;) or descending (&#x60;desc&#x60;). | 
 **domains** | **string** | A domain name or comma-delimited list of domain name. | 
 **categories** | **string** | A category ID or comma-delimited list of category ID. | 
 **policycategories** | **string** | A category ID or comma-delimited list of category ID. Filter the request by the categories that trigger a policy. | 
 **ip** | **string** | An IP address. | 
 **identityids** | **string** | An identity ID or comma-delimited list of identity IDs. | 
 **identitytypes** | **string** | An identity type or comma-delimited list of identity types. | 
 **applicationid** | **string** | The ID of the application. | 
 **verdict** | **string** | A string or comma-delimited string that describes whether the traffic can reach the destination. | 
 **threats** | **string** | A threat name or comma-delimited list of threat names. | 
 **threattypes** | **string** | A threat type or comma-delimited list of threat types. | 
 **filternoisydomains** | **bool** | Filter out domains that generate a lot of insignificant traffic (noise). | 
 **timezone** | **string** | Display the timestamp of the traffic events in the specified timezone. For the timezone, provide a continent and city separated by an url-encoded forward slash (&#39;/&#39;), for example: timezone&#x3D;&#39;ASIA%2fCALCUTTA&#39;. | 

### Return type

[**GetActivityDns200Response**](GetActivityDns200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActivityFirewall

> GetActivityFirewall200Response GetActivityFirewall(ctx).From(from).To(to).Limit(limit).Offset(offset).Identityids(identityids).Ruleid(ruleid).Verdict(verdict).Ip(ip).Ports(ports).Timezone(timezone).Categories(categories).Execute()

Get Activity Firewall



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
	identityids := "1,2,3" // string | An identity ID or comma-delimited list of identity IDs. (optional)
	ruleid := int64(1) // int64 | The firewall policy rule ID. (optional)
	verdict := "allowed,blocked" // string | A string or comma-delimited string that describes whether the traffic can reach the destination. (optional)
	ip := "10.10.10.10" // string | An IP address. (optional)
	ports := "7351,80" // string | A port number or comma-delimited list of port numbers. (optional)
	timezone := "ASIA%2fCALCUTTA" // string | Display the timestamp of the traffic events in the specified timezone. For the timezone, provide a continent and city separated by an url-encoded forward slash ('/'), for example: timezone='ASIA%2fCALCUTTA'. (optional)
	categories := "148,151,66" // string | A category ID or comma-delimited list of category ID. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityAPI.GetActivityFirewall(context.Background()).From(from).To(to).Limit(limit).Offset(offset).Identityids(identityids).Ruleid(ruleid).Verdict(verdict).Ip(ip).Ports(ports).Timezone(timezone).Categories(categories).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityAPI.GetActivityFirewall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivityFirewall`: GetActivityFirewall200Response
	fmt.Fprintf(os.Stdout, "Response from `ActivityAPI.GetActivityFirewall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActivityFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | A timestamp or relative time string (for example: &#39;-1days&#39;). Filter for data that appears after this time. | 
 **to** | **string** | A timestamp or relative time string (for example: &#39;now&#39;). Filter for data that appears before this time. | 
 **limit** | **int64** | The maximum number of records to return from the collection. | [default to 100]
 **offset** | **int64** | A number that represents an index in the collection. | [default to 0]
 **identityids** | **string** | An identity ID or comma-delimited list of identity IDs. | 
 **ruleid** | **int64** | The firewall policy rule ID. | 
 **verdict** | **string** | A string or comma-delimited string that describes whether the traffic can reach the destination. | 
 **ip** | **string** | An IP address. | 
 **ports** | **string** | A port number or comma-delimited list of port numbers. | 
 **timezone** | **string** | Display the timestamp of the traffic events in the specified timezone. For the timezone, provide a continent and city separated by an url-encoded forward slash (&#39;/&#39;), for example: timezone&#x3D;&#39;ASIA%2fCALCUTTA&#39;. | 
 **categories** | **string** | A category ID or comma-delimited list of category ID. | 

### Return type

[**GetActivityFirewall200Response**](GetActivityFirewall200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActivityIP

> GetActivityIP200Response GetActivityIP(ctx).From(from).To(to).Limit(limit).Offset(offset).Identityids(identityids).Identitytypes(identitytypes).Categories(categories).Verdict(verdict).Ip(ip).Ports(ports).Execute()

Get Activity IP



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
	identityids := "1,2,3" // string | An identity ID or comma-delimited list of identity IDs. (optional)
	identitytypes := "network,roaming" // string | An identity type or comma-delimited list of identity types. (optional)
	categories := "148,151,66" // string | A category ID or comma-delimited list of category ID. (optional)
	verdict := "allowed,blocked" // string | A string or comma-delimited string that describes whether the traffic can reach the destination. (optional)
	ip := "10.10.10.10" // string | An IP address. (optional)
	ports := "7351,80" // string | A port number or comma-delimited list of port numbers. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityAPI.GetActivityIP(context.Background()).From(from).To(to).Limit(limit).Offset(offset).Identityids(identityids).Identitytypes(identitytypes).Categories(categories).Verdict(verdict).Ip(ip).Ports(ports).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityAPI.GetActivityIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivityIP`: GetActivityIP200Response
	fmt.Fprintf(os.Stdout, "Response from `ActivityAPI.GetActivityIP`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActivityIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | A timestamp or relative time string (for example: &#39;-1days&#39;). Filter for data that appears after this time. | 
 **to** | **string** | A timestamp or relative time string (for example: &#39;now&#39;). Filter for data that appears before this time. | 
 **limit** | **int64** | The maximum number of records to return from the collection. | [default to 100]
 **offset** | **int64** | A number that represents an index in the collection. | [default to 0]
 **identityids** | **string** | An identity ID or comma-delimited list of identity IDs. | 
 **identitytypes** | **string** | An identity type or comma-delimited list of identity types. | 
 **categories** | **string** | A category ID or comma-delimited list of category ID. | 
 **verdict** | **string** | A string or comma-delimited string that describes whether the traffic can reach the destination. | 
 **ip** | **string** | An IP address. | 
 **ports** | **string** | A port number or comma-delimited list of port numbers. | 

### Return type

[**GetActivityIP200Response**](GetActivityIP200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActivityIntrusion

> GetActivityIntrusion200Response GetActivityIntrusion(ctx).From(from).To(to).Limit(limit).Offset(offset).Identityids(identityids).Signatures(signatures).Signaturelistids(signaturelistids).Intrusionaction(intrusionaction).Ip(ip).Ports(ports).Filternoisydomains(filternoisydomains).Timezone(timezone).Execute()

Get Activity Intrusion



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
	identityids := "1,2,3" // string | An identity ID or comma-delimited list of identity IDs. (optional)
	signatures := "1-2,1-4" // string | The signature or comma-separated list of <signatureid>-<generatorid> signatures. (optional)
	signaturelistids := "1,2" // string | The signature ID or comma-separated list of signature list IDs. (optional)
	intrusionaction := "detected,would_block" // string | An action or list of comma-separated intrusion actions. Valid values are: `would_block`, `blocked`, and `detected`. (optional)
	ip := "10.10.10.10" // string | An IP address. (optional)
	ports := "7351,80" // string | A port number or comma-delimited list of port numbers. (optional)
	filternoisydomains := true // bool | Filter out domains that generate a lot of insignificant traffic (noise). (optional)
	timezone := "ASIA%2fCALCUTTA" // string | Display the timestamp of the traffic events in the specified timezone. For the timezone, provide a continent and city separated by an url-encoded forward slash ('/'), for example: timezone='ASIA%2fCALCUTTA'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityAPI.GetActivityIntrusion(context.Background()).From(from).To(to).Limit(limit).Offset(offset).Identityids(identityids).Signatures(signatures).Signaturelistids(signaturelistids).Intrusionaction(intrusionaction).Ip(ip).Ports(ports).Filternoisydomains(filternoisydomains).Timezone(timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityAPI.GetActivityIntrusion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivityIntrusion`: GetActivityIntrusion200Response
	fmt.Fprintf(os.Stdout, "Response from `ActivityAPI.GetActivityIntrusion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActivityIntrusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | A timestamp or relative time string (for example: &#39;-1days&#39;). Filter for data that appears after this time. | 
 **to** | **string** | A timestamp or relative time string (for example: &#39;now&#39;). Filter for data that appears before this time. | 
 **limit** | **int64** | The maximum number of records to return from the collection. | [default to 100]
 **offset** | **int64** | A number that represents an index in the collection. | [default to 0]
 **identityids** | **string** | An identity ID or comma-delimited list of identity IDs. | 
 **signatures** | **string** | The signature or comma-separated list of &lt;signatureid&gt;-&lt;generatorid&gt; signatures. | 
 **signaturelistids** | **string** | The signature ID or comma-separated list of signature list IDs. | 
 **intrusionaction** | **string** | An action or list of comma-separated intrusion actions. Valid values are: &#x60;would_block&#x60;, &#x60;blocked&#x60;, and &#x60;detected&#x60;. | 
 **ip** | **string** | An IP address. | 
 **ports** | **string** | A port number or comma-delimited list of port numbers. | 
 **filternoisydomains** | **bool** | Filter out domains that generate a lot of insignificant traffic (noise). | 
 **timezone** | **string** | Display the timestamp of the traffic events in the specified timezone. For the timezone, provide a continent and city separated by an url-encoded forward slash (&#39;/&#39;), for example: timezone&#x3D;&#39;ASIA%2fCALCUTTA&#39;. | 

### Return type

[**GetActivityIntrusion200Response**](GetActivityIntrusion200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActivityProxy

> GetActivityProxy200Response GetActivityProxy(ctx).From(from).To(to).Limit(limit).Order(order).Offset(offset).Domains(domains).Urls(urls).Categories(categories).Policycategories(policycategories).Ip(ip).Ports(ports).Identityids(identityids).Identitytypes(identitytypes).Applicationid(applicationid).Verdict(verdict).Ruleid(ruleid).Filename(filename).Securityoverridden(securityoverridden).Bundleid(bundleid).Threats(threats).Threattypes(threattypes).Ampdisposition(ampdisposition).Antivirusthreats(antivirusthreats).Tenantcontrols(tenantcontrols).Isolatedstate(isolatedstate).IsolatedFileAction(isolatedFileAction).Datalosspreventionstate(datalosspreventionstate).Httperrors(httperrors).Timezone(timezone).Execute()

Get Activity Proxy



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
	order := "desc" // string | A string that describes how to order the results: ascending (`asc`) or descending (`desc`). (optional)
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
	ruleid := int64(1) // int64 | The firewall policy rule ID. (optional)
	filename := "myfilename_*" // string | A string that identifies a filename. Filter the request by the filename. Supports globbing or use of the wildcard character ('*'). The asterisk (*) matches zero or more occurrences of any character. (optional)
	securityoverridden := true // bool | Specify whether to filter on requests that override security. (optional)
	bundleid := int64(1) // int64 | A proxy bundle ID. (optional)
	threats := "threats_example" // string | A threat name or comma-delimited list of threat names. (optional)
	threattypes := "threattypes_example" // string | A threat type or comma-delimited list of threat types. (optional)
	ampdisposition := "clean,malicious,unknown" // string | An AMP disposition string or a comma-delimited list of AMP disposition strings. (optional)
	antivirusthreats := "Trojan.Linux.Generic.144075" // string | A threat name or comma-delimited list of threat names. (optional)
	tenantcontrols := true // bool | If set to `true`, filter data for requests that are part of a tenant control policy. (optional)
	isolatedstate := "isolated" // string | A string that describes the remote browser isolation (RBI) isolation type. (optional)
	isolatedFileAction := "downloaded-safe-pdf" // string | A string that describes the remote browser isolation (RBI) file action type. (optional)
	datalosspreventionstate := "blocked" // string | A string that describes the status of a destination. Filter for requests that are blocked by the DLP layer security. (optional)
	httperrors := "certificateerror" // string | Filter data for requests that resulted in a TLS error or a certificate error. (optional)
	timezone := "ASIA%2fCALCUTTA" // string | Display the timestamp of the traffic events in the specified timezone. For the timezone, provide a continent and city separated by an url-encoded forward slash ('/'), for example: timezone='ASIA%2fCALCUTTA'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityAPI.GetActivityProxy(context.Background()).From(from).To(to).Limit(limit).Order(order).Offset(offset).Domains(domains).Urls(urls).Categories(categories).Policycategories(policycategories).Ip(ip).Ports(ports).Identityids(identityids).Identitytypes(identitytypes).Applicationid(applicationid).Verdict(verdict).Ruleid(ruleid).Filename(filename).Securityoverridden(securityoverridden).Bundleid(bundleid).Threats(threats).Threattypes(threattypes).Ampdisposition(ampdisposition).Antivirusthreats(antivirusthreats).Tenantcontrols(tenantcontrols).Isolatedstate(isolatedstate).IsolatedFileAction(isolatedFileAction).Datalosspreventionstate(datalosspreventionstate).Httperrors(httperrors).Timezone(timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityAPI.GetActivityProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivityProxy`: GetActivityProxy200Response
	fmt.Fprintf(os.Stdout, "Response from `ActivityAPI.GetActivityProxy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActivityProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | A timestamp or relative time string (for example: &#39;-1days&#39;). Filter for data that appears after this time. | 
 **to** | **string** | A timestamp or relative time string (for example: &#39;now&#39;). Filter for data that appears before this time. | 
 **limit** | **int64** | The maximum number of records to return from the collection. | [default to 100]
 **order** | **string** | A string that describes how to order the results: ascending (&#x60;asc&#x60;) or descending (&#x60;desc&#x60;). | 
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
 **ruleid** | **int64** | The firewall policy rule ID. | 
 **filename** | **string** | A string that identifies a filename. Filter the request by the filename. Supports globbing or use of the wildcard character (&#39;*&#39;). The asterisk (*) matches zero or more occurrences of any character. | 
 **securityoverridden** | **bool** | Specify whether to filter on requests that override security. | 
 **bundleid** | **int64** | A proxy bundle ID. | 
 **threats** | **string** | A threat name or comma-delimited list of threat names. | 
 **threattypes** | **string** | A threat type or comma-delimited list of threat types. | 
 **ampdisposition** | **string** | An AMP disposition string or a comma-delimited list of AMP disposition strings. | 
 **antivirusthreats** | **string** | A threat name or comma-delimited list of threat names. | 
 **tenantcontrols** | **bool** | If set to &#x60;true&#x60;, filter data for requests that are part of a tenant control policy. | 
 **isolatedstate** | **string** | A string that describes the remote browser isolation (RBI) isolation type. | 
 **isolatedFileAction** | **string** | A string that describes the remote browser isolation (RBI) file action type. | 
 **datalosspreventionstate** | **string** | A string that describes the status of a destination. Filter for requests that are blocked by the DLP layer security. | 
 **httperrors** | **string** | Filter data for requests that resulted in a TLS error or a certificate error. | 
 **timezone** | **string** | Display the timestamp of the traffic events in the specified timezone. For the timezone, provide a continent and city separated by an url-encoded forward slash (&#39;/&#39;), for example: timezone&#x3D;&#39;ASIA%2fCALCUTTA&#39;. | 

### Return type

[**GetActivityProxy200Response**](GetActivityProxy200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActivityZTNA

> GetActivityZTNA200Response GetActivityZTNA(ctx).From(from).To(to).Limit(limit).Offset(offset).Ztnatype(ztnatype).Timezone(timezone).Execute()

Activity ZTNA



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
	ztnatype := "clientless" // string | Specify the Zero Trust Network Access (ZTNA) session type. (optional)
	timezone := "ASIA%2fCALCUTTA" // string | Display the timestamp of the traffic events in the specified timezone. For the timezone, provide a continent and city separated by an url-encoded forward slash ('/'), for example: timezone='ASIA%2fCALCUTTA'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityAPI.GetActivityZTNA(context.Background()).From(from).To(to).Limit(limit).Offset(offset).Ztnatype(ztnatype).Timezone(timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityAPI.GetActivityZTNA``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivityZTNA`: GetActivityZTNA200Response
	fmt.Fprintf(os.Stdout, "Response from `ActivityAPI.GetActivityZTNA`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActivityZTNARequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | A timestamp or relative time string (for example: &#39;-1days&#39;). Filter for data that appears after this time. | 
 **to** | **string** | A timestamp or relative time string (for example: &#39;now&#39;). Filter for data that appears before this time. | 
 **limit** | **int64** | The maximum number of records to return from the collection. | [default to 100]
 **offset** | **int64** | A number that represents an index in the collection. | [default to 0]
 **ztnatype** | **string** | Specify the Zero Trust Network Access (ZTNA) session type. | 
 **timezone** | **string** | Display the timestamp of the traffic events in the specified timezone. For the timezone, provide a continent and city separated by an url-encoded forward slash (&#39;/&#39;), for example: timezone&#x3D;&#39;ASIA%2fCALCUTTA&#39;. | 

### Return type

[**GetActivityZTNA200Response**](GetActivityZTNA200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

