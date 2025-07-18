# \SummariesByRuleAPI

All URIs are relative to *https://api.sse.cisco.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSummariesByRuleFirewallHitcount**](SummariesByRuleAPI.md#GetSummariesByRuleFirewallHitcount) | **Get** /reports/v2/summaries-by-rule/firewall-hitcount | Get Summaries by Rule for Policy Rule Firewall Hit Count
[**GetSummariesByRuleHitcount**](SummariesByRuleAPI.md#GetSummariesByRuleHitcount) | **Get** /reports/v2/summaries-by-rule/hitcount | Get Summaries by Rule for Policy Rule Hit Count
[**GetSummariesByRuleIntrusion**](SummariesByRuleAPI.md#GetSummariesByRuleIntrusion) | **Get** /reports/v2/summaries-by-rule/intrusion | Get Summaries by Rule



## GetSummariesByRuleFirewallHitcount

> GetSummariesByRuleFirewallHitcount200Response GetSummariesByRuleFirewallHitcount(ctx).From(from).To(to).Ruleids(ruleids).Execute()

Get Summaries by Rule for Policy Rule Firewall Hit Count



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
	ruleids := "31,47" // string | A comma-delimited list of firewall policy rule IDs.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SummariesByRuleAPI.GetSummariesByRuleFirewallHitcount(context.Background()).From(from).To(to).Ruleids(ruleids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SummariesByRuleAPI.GetSummariesByRuleFirewallHitcount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSummariesByRuleFirewallHitcount`: GetSummariesByRuleFirewallHitcount200Response
	fmt.Fprintf(os.Stdout, "Response from `SummariesByRuleAPI.GetSummariesByRuleFirewallHitcount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSummariesByRuleFirewallHitcountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | A timestamp or relative time string (for example: &#39;-1days&#39;). Filter for data that appears after this time. | 
 **to** | **string** | A timestamp or relative time string (for example: &#39;now&#39;). Filter for data that appears before this time. | 
 **ruleids** | **string** | A comma-delimited list of firewall policy rule IDs. | 

### Return type

[**GetSummariesByRuleFirewallHitcount200Response**](GetSummariesByRuleFirewallHitcount200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSummariesByRuleHitcount

> GetSummariesByRuleHitcount200Response GetSummariesByRuleHitcount(ctx).From(from).To(to).Ruleids(ruleids).Execute()

Get Summaries by Rule for Policy Rule Hit Count



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
	ruleids := "31,47" // string | A comma-delimited list of firewall policy rule IDs.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SummariesByRuleAPI.GetSummariesByRuleHitcount(context.Background()).From(from).To(to).Ruleids(ruleids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SummariesByRuleAPI.GetSummariesByRuleHitcount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSummariesByRuleHitcount`: GetSummariesByRuleHitcount200Response
	fmt.Fprintf(os.Stdout, "Response from `SummariesByRuleAPI.GetSummariesByRuleHitcount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSummariesByRuleHitcountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | A timestamp or relative time string (for example: &#39;-1days&#39;). Filter for data that appears after this time. | 
 **to** | **string** | A timestamp or relative time string (for example: &#39;now&#39;). Filter for data that appears before this time. | 
 **ruleids** | **string** | A comma-delimited list of firewall policy rule IDs. | 

### Return type

[**GetSummariesByRuleHitcount200Response**](GetSummariesByRuleHitcount200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSummariesByRuleIntrusion

> GetSummariesByRuleIntrusion200Response GetSummariesByRuleIntrusion(ctx).From(from).To(to).Limit(limit).Offset(offset).Signatures(signatures).Signaturelistids(signaturelistids).Ipsprofile(ipsprofile).Ip(ip).Identityids(identityids).Identitytypes(identitytypes).Intrusionaction(intrusionaction).Ports(ports).Filternoisydomains(filternoisydomains).Execute()

Get Summaries by Rule



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
	signatures := "1-2,1-4" // string | The signature or comma-separated list of <signatureid>-<generatorid> signatures. (optional)
	signaturelistids := "1,2" // string | The signature ID or comma-separated list of signature list IDs. (optional)
	ipsprofile := "config,profile" // string | An IPS profile string or comma-delimited list of IPS profile string. (optional)
	ip := "10.10.10.10" // string | An IP address. (optional)
	identityids := "1,2,3" // string | An identity ID or comma-delimited list of identity IDs. (optional)
	identitytypes := "network,roaming" // string | An identity type or comma-delimited list of identity types. (optional)
	intrusionaction := "detected,would_block" // string | An action or list of comma-separated intrusion actions. Valid values are: `would_block`, `blocked`, and `detected`. (optional)
	ports := "7351,80" // string | A port number or comma-delimited list of port numbers. (optional)
	filternoisydomains := true // bool | Filter out domains that generate a lot of insignificant traffic (noise). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SummariesByRuleAPI.GetSummariesByRuleIntrusion(context.Background()).From(from).To(to).Limit(limit).Offset(offset).Signatures(signatures).Signaturelistids(signaturelistids).Ipsprofile(ipsprofile).Ip(ip).Identityids(identityids).Identitytypes(identitytypes).Intrusionaction(intrusionaction).Ports(ports).Filternoisydomains(filternoisydomains).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SummariesByRuleAPI.GetSummariesByRuleIntrusion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSummariesByRuleIntrusion`: GetSummariesByRuleIntrusion200Response
	fmt.Fprintf(os.Stdout, "Response from `SummariesByRuleAPI.GetSummariesByRuleIntrusion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSummariesByRuleIntrusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | A timestamp or relative time string (for example: &#39;-1days&#39;). Filter for data that appears after this time. | 
 **to** | **string** | A timestamp or relative time string (for example: &#39;now&#39;). Filter for data that appears before this time. | 
 **limit** | **int64** | The maximum number of records to return from the collection. | [default to 100]
 **offset** | **int64** | A number that represents an index in the collection. | [default to 0]
 **signatures** | **string** | The signature or comma-separated list of &lt;signatureid&gt;-&lt;generatorid&gt; signatures. | 
 **signaturelistids** | **string** | The signature ID or comma-separated list of signature list IDs. | 
 **ipsprofile** | **string** | An IPS profile string or comma-delimited list of IPS profile string. | 
 **ip** | **string** | An IP address. | 
 **identityids** | **string** | An identity ID or comma-delimited list of identity IDs. | 
 **identitytypes** | **string** | An identity type or comma-delimited list of identity types. | 
 **intrusionaction** | **string** | An action or list of comma-separated intrusion actions. Valid values are: &#x60;would_block&#x60;, &#x60;blocked&#x60;, and &#x60;detected&#x60;. | 
 **ports** | **string** | A port number or comma-delimited list of port numbers. | 
 **filternoisydomains** | **bool** | Filter out domains that generate a lot of insignificant traffic (noise). | 

### Return type

[**GetSummariesByRuleIntrusion200Response**](GetSummariesByRuleIntrusion200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

