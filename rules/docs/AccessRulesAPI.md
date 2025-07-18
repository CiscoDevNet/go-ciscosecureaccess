# \AccessRulesAPI

All URIs are relative to *https://api.sse.cisco.com/policies/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRule**](AccessRulesAPI.md#AddRule) | **Post** /rules | Create Rule
[**DeleteRule**](AccessRulesAPI.md#DeleteRule) | **Delete** /rules/{ruleId} | Delete Rule
[**GetRegionalLocations**](AccessRulesAPI.md#GetRegionalLocations) | **Get** /geolocations | List GeoLocations
[**GetRule**](AccessRulesAPI.md#GetRule) | **Get** /rules/{ruleId} | Get Rule
[**ListRules**](AccessRulesAPI.md#ListRules) | **Get** /rules | List Rules
[**PutRule**](AccessRulesAPI.md#PutRule) | **Put** /rules/{ruleId} | Update Rule
[**PutRules**](AccessRulesAPI.md#PutRules) | **Put** /rules | Enable Rules



## AddRule

> Rule AddRule(ctx).AddRuleRequest(addRuleRequest).Execute()

Create Rule



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
	addRuleRequest := *openapiclient.NewAddRuleRequest("SSE Rule 1", openapiclient.ruleAction("allow"), []openapiclient.RuleConditionsInner{*openapiclient.NewRuleConditionsInner()}, []openapiclient.RuleSettingsInner{*openapiclient.NewRuleSettingsInner()}) // AddRuleRequest | Create the rule with the specific properties.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRulesAPI.AddRule(context.Background()).AddRuleRequest(addRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRulesAPI.AddRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddRule`: Rule
	fmt.Fprintf(os.Stdout, "Response from `AccessRulesAPI.AddRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addRuleRequest** | [**AddRuleRequest**](AddRuleRequest.md) | Create the rule with the specific properties. | 

### Return type

[**Rule**](Rule.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRule

> DeleteRule(ctx, ruleId).Execute()

Delete Rule



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
	ruleId := int64(2456) // int64 | The ID of the rule.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccessRulesAPI.DeleteRule(context.Background(), ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRulesAPI.DeleteRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **int64** | The ID of the rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleRequest struct via the builder pattern


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


## GetRegionalLocations

> Regions GetRegionalLocations(ctx).Execute()

List GeoLocations



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
	resp, r, err := apiClient.AccessRulesAPI.GetRegionalLocations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRulesAPI.GetRegionalLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegionalLocations`: Regions
	fmt.Fprintf(os.Stdout, "Response from `AccessRulesAPI.GetRegionalLocations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegionalLocationsRequest struct via the builder pattern


### Return type

[**Regions**](Regions.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRule

> Rule GetRule(ctx, ruleId).Execute()

Get Rule



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
	ruleId := int64(2456) // int64 | The ID of the rule.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRulesAPI.GetRule(context.Background(), ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRulesAPI.GetRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRule`: Rule
	fmt.Fprintf(os.Stdout, "Response from `AccessRulesAPI.GetRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **int64** | The ID of the rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Rule**](Rule.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRules

> Rules ListRules(ctx).Offset(offset).Limit(limit).RuleName(ruleName).Filters(filters).Execute()

List Rules



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
	offset := int64(10) // int64 | The place to start reading in the collection. The default value is `0`. For example, the offset of the first page in the collection is `0`. If the limit is 10, the offset for next page is `10`. (optional) (default to 0)
	limit := int64(20) // int64 | Set the number of items on a page. The maximum items that are allowed on a page in the response are 1000. (optional) (default to 10)
	ruleName := "Allow all rule" // string | Filter the rules using the name of the rule or a sequence of characters found in the rule name. (optional)
	filters := map[string][]openapiclient.RuleFilters{"key": openapiclient.RuleFilters_value{RuleSourceDestFilters: openapiclient.NewRuleSourceDestFilters()}} // RuleFilters | Filter the rules by one or more properties. Filter on: `ruleName`, `ruleDescription`, `ruleIsEnabled`, `ruleIsDefault`, `ruleAction`, `attributeName` and `attributeValue`, `settingName` and `settingValue`, `rulePriority`, `destinations`, `sources`, `ruleConditions`. Specify the properties in the JSON format. **Note:** Filter on either `ruleConditions` or the `sources` and `destinations` properties, but not both sets of properties.  Example:  ``` {   \"destinations\": [     \"umbrella.destination.destination_list_ids\": \"*\"   ],   \"sources\": [     \"umbrella.source.identity_ids\": \"*\"   ],   \"rulePriority\": \"1,2,3\" } ``` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRulesAPI.ListRules(context.Background()).Offset(offset).Limit(limit).RuleName(ruleName).Filters(filters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRulesAPI.ListRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRules`: Rules
	fmt.Fprintf(os.Stdout, "Response from `AccessRulesAPI.ListRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int64** | The place to start reading in the collection. The default value is &#x60;0&#x60;. For example, the offset of the first page in the collection is &#x60;0&#x60;. If the limit is 10, the offset for next page is &#x60;10&#x60;. | [default to 0]
 **limit** | **int64** | Set the number of items on a page. The maximum items that are allowed on a page in the response are 1000. | [default to 10]
 **ruleName** | **string** | Filter the rules using the name of the rule or a sequence of characters found in the rule name. | 
 **filters** | [**map[string]RuleFiltersValue**](RuleFiltersValue.md) | Filter the rules by one or more properties. Filter on: &#x60;ruleName&#x60;, &#x60;ruleDescription&#x60;, &#x60;ruleIsEnabled&#x60;, &#x60;ruleIsDefault&#x60;, &#x60;ruleAction&#x60;, &#x60;attributeName&#x60; and &#x60;attributeValue&#x60;, &#x60;settingName&#x60; and &#x60;settingValue&#x60;, &#x60;rulePriority&#x60;, &#x60;destinations&#x60;, &#x60;sources&#x60;, &#x60;ruleConditions&#x60;. Specify the properties in the JSON format. **Note:** Filter on either &#x60;ruleConditions&#x60; or the &#x60;sources&#x60; and &#x60;destinations&#x60; properties, but not both sets of properties.  Example:  &#x60;&#x60;&#x60; {   \&quot;destinations\&quot;: [     \&quot;umbrella.destination.destination_list_ids\&quot;: \&quot;*\&quot;   ],   \&quot;sources\&quot;: [     \&quot;umbrella.source.identity_ids\&quot;: \&quot;*\&quot;   ],   \&quot;rulePriority\&quot;: \&quot;1,2,3\&quot; } &#x60;&#x60;&#x60; | 

### Return type

[**Rules**](Rules.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRule

> Rule PutRule(ctx, ruleId).PutRuleRequest(putRuleRequest).Execute()

Update Rule



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
	ruleId := int64(2456) // int64 | The ID of the rule.
	putRuleRequest := *openapiclient.NewPutRuleRequest("SSE Rule 1", openapiclient.ruleAction("allow"), int64(1), []openapiclient.RuleConditionsInner{*openapiclient.NewRuleConditionsInner()}, []openapiclient.RuleSettingsInner{*openapiclient.NewRuleSettingsInner()}) // PutRuleRequest | Create the rule with the specific properties.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRulesAPI.PutRule(context.Background(), ruleId).PutRuleRequest(putRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRulesAPI.PutRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutRule`: Rule
	fmt.Fprintf(os.Stdout, "Response from `AccessRulesAPI.PutRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **int64** | The ID of the rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putRuleRequest** | [**PutRuleRequest**](PutRuleRequest.md) | Create the rule with the specific properties. | 

### Return type

[**Rule**](Rule.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRules

> []RulesResponseInner PutRules(ctx).PutRulesRequest(putRulesRequest).Execute()

Enable Rules



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
	putRulesRequest := *openapiclient.NewPutRulesRequest([]int64{int64(183456)}, []openapiclient.PutRulesRequestPropertiesInner{*openapiclient.NewPutRulesRequestPropertiesInner()}) // PutRulesRequest | Provide the list of rule IDs and set the `ruleIsEnabled` property on these rules. **Note:** You cannot update the `ruleIsEnabled` property on the default rules. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRulesAPI.PutRules(context.Background()).PutRulesRequest(putRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRulesAPI.PutRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutRules`: []RulesResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AccessRulesAPI.PutRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putRulesRequest** | [**PutRulesRequest**](PutRulesRequest.md) | Provide the list of rule IDs and set the &#x60;ruleIsEnabled&#x60; property on these rules. **Note:** You cannot update the &#x60;ruleIsEnabled&#x60; property on the default rules. | 

### Return type

[**[]RulesResponseInner**](RulesResponseInner.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

