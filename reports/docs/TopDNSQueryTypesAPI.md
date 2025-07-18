# \TopDNSQueryTypesAPI

All URIs are relative to *https://api.sse.cisco.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTopDNSQueryTypes**](TopDNSQueryTypesAPI.md#GetTopDNSQueryTypes) | **Get** /reports/v2/top-dns-query-types | Get Top DNS Query Types



## GetTopDNSQueryTypes

> GetTopDNSQueryTypes200Response GetTopDNSQueryTypes(ctx).From(from).To(to).Limit(limit).Offset(offset).Order(order).Domains(domains).Categories(categories).Policycategories(policycategories).Ip(ip).Identityids(identityids).Identitytypes(identitytypes).Applicationid(applicationid).Verdict(verdict).Threats(threats).Threattypes(threattypes).Filternoisydomains(filternoisydomains).Timezone(timezone).Execute()

Get Top DNS Query Types



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
	resp, r, err := apiClient.TopDNSQueryTypesAPI.GetTopDNSQueryTypes(context.Background()).From(from).To(to).Limit(limit).Offset(offset).Order(order).Domains(domains).Categories(categories).Policycategories(policycategories).Ip(ip).Identityids(identityids).Identitytypes(identitytypes).Applicationid(applicationid).Verdict(verdict).Threats(threats).Threattypes(threattypes).Filternoisydomains(filternoisydomains).Timezone(timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopDNSQueryTypesAPI.GetTopDNSQueryTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTopDNSQueryTypes`: GetTopDNSQueryTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `TopDNSQueryTypesAPI.GetTopDNSQueryTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTopDNSQueryTypesRequest struct via the builder pattern


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

[**GetTopDNSQueryTypes200Response**](GetTopDNSQueryTypes200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

