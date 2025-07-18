# \SummariesByDestinationAPI

All URIs are relative to *https://api.sse.cisco.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSummariesByDestination**](SummariesByDestinationAPI.md#GetSummariesByDestination) | **Get** /reports/v2/summaries-by-destination | Get Summaries by Destination (All)
[**GetSummariesByDestinationType**](SummariesByDestinationAPI.md#GetSummariesByDestinationType) | **Get** /reports/v2/summaries-by-destination/{type} | Get Summaries by Destination For Type



## GetSummariesByDestination

> GetSummariesByDestination200Response GetSummariesByDestination(ctx).From(from).To(to).Limit(limit).Offset(offset).Domains(domains).Urls(urls).Categories(categories).Policycategories(policycategories).Ip(ip).Identityids(identityids).Identitytypes(identitytypes).Applicationid(applicationid).Verdict(verdict).Ruleid(ruleid).Filename(filename).Securityoverridden(securityoverridden).Bundleid(bundleid).Threats(threats).Threattypes(threattypes).Ampdisposition(ampdisposition).Antivirusthreats(antivirusthreats).Datalosspreventionstate(datalosspreventionstate).Filternoisydomains(filternoisydomains).Execute()

Get Summaries by Destination (All)



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
	datalosspreventionstate := "blocked" // string | A string that describes the status of a destination. Filter for requests that are blocked by the DLP layer security. (optional)
	filternoisydomains := true // bool | Filter out domains that generate a lot of insignificant traffic (noise). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SummariesByDestinationAPI.GetSummariesByDestination(context.Background()).From(from).To(to).Limit(limit).Offset(offset).Domains(domains).Urls(urls).Categories(categories).Policycategories(policycategories).Ip(ip).Identityids(identityids).Identitytypes(identitytypes).Applicationid(applicationid).Verdict(verdict).Ruleid(ruleid).Filename(filename).Securityoverridden(securityoverridden).Bundleid(bundleid).Threats(threats).Threattypes(threattypes).Ampdisposition(ampdisposition).Antivirusthreats(antivirusthreats).Datalosspreventionstate(datalosspreventionstate).Filternoisydomains(filternoisydomains).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SummariesByDestinationAPI.GetSummariesByDestination``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSummariesByDestination`: GetSummariesByDestination200Response
	fmt.Fprintf(os.Stdout, "Response from `SummariesByDestinationAPI.GetSummariesByDestination`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSummariesByDestinationRequest struct via the builder pattern


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
 **datalosspreventionstate** | **string** | A string that describes the status of a destination. Filter for requests that are blocked by the DLP layer security. | 
 **filternoisydomains** | **bool** | Filter out domains that generate a lot of insignificant traffic (noise). | 

### Return type

[**GetSummariesByDestination200Response**](GetSummariesByDestination200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSummariesByDestinationType

> GetSummariesByDestination200Response GetSummariesByDestinationType(ctx, type_).From(from).To(to).Limit(limit).Offset(offset).Domains(domains).Urls(urls).Categories(categories).Policycategories(policycategories).Ip(ip).Identityids(identityids).Identitytypes(identitytypes).Applicationid(applicationid).Verdict(verdict).Ruleid(ruleid).Filename(filename).Securityoverridden(securityoverridden).Bundleid(bundleid).Threats(threats).Threattypes(threattypes).Ampdisposition(ampdisposition).Antivirusthreats(antivirusthreats).Datalosspreventionstate(datalosspreventionstate).Filternoisydomains(filternoisydomains).Execute()

Get Summaries by Destination For Type



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
	type_ := "dns" // string | Specify the type of traffic.
	from := "1639146300000" // string | A timestamp or relative time string (for example: '-1days'). Filter for data that appears after this time.
	to := "1640010300000" // string | A timestamp or relative time string (for example: 'now'). Filter for data that appears before this time.
	limit := int64(100) // int64 | The maximum number of records to return from the collection. (default to 100)
	offset := int64(0) // int64 | A number that represents an index in the collection. (optional) (default to 0)
	domains := "cisco.com,nasa.gov" // string | A domain name or comma-delimited list of domain name. (optional)
	urls := "https://google.com,facebook.com/help" // string | A URL or comma-delimited list of URL. (optional)
	categories := "148,151,66" // string | A category ID or comma-delimited list of category ID. (optional)
	policycategories := "67,69" // string | A category ID or comma-delimited list of category ID. Filter the request by the categories that trigger a policy. (optional)
	ip := "10.10.10.10" // string | An IP address. (optional)
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
	datalosspreventionstate := "blocked" // string | A string that describes the status of a destination. Filter for requests that are blocked by the DLP layer security. (optional)
	filternoisydomains := true // bool | Filter out domains that generate a lot of insignificant traffic (noise). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SummariesByDestinationAPI.GetSummariesByDestinationType(context.Background(), type_).From(from).To(to).Limit(limit).Offset(offset).Domains(domains).Urls(urls).Categories(categories).Policycategories(policycategories).Ip(ip).Identityids(identityids).Identitytypes(identitytypes).Applicationid(applicationid).Verdict(verdict).Ruleid(ruleid).Filename(filename).Securityoverridden(securityoverridden).Bundleid(bundleid).Threats(threats).Threattypes(threattypes).Ampdisposition(ampdisposition).Antivirusthreats(antivirusthreats).Datalosspreventionstate(datalosspreventionstate).Filternoisydomains(filternoisydomains).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SummariesByDestinationAPI.GetSummariesByDestinationType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSummariesByDestinationType`: GetSummariesByDestination200Response
	fmt.Fprintf(os.Stdout, "Response from `SummariesByDestinationAPI.GetSummariesByDestinationType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | Specify the type of traffic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSummariesByDestinationTypeRequest struct via the builder pattern


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
 **datalosspreventionstate** | **string** | A string that describes the status of a destination. Filter for requests that are blocked by the DLP layer security. | 
 **filternoisydomains** | **bool** | Filter out domains that generate a lot of insignificant traffic (noise). | 

### Return type

[**GetSummariesByDestination200Response**](GetSummariesByDestination200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

