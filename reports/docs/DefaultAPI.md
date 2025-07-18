# \DefaultAPI

All URIs are relative to *https://api.sse.cisco.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApplications**](DefaultAPI.md#GetApplications) | **Get** /reports/v2/applications | Get Applications
[**GetApplicationsAppDiscovery**](DefaultAPI.md#GetApplicationsAppDiscovery) | **Get** /reports/v2/appDiscovery/applications | List Applications



## GetApplications

> GetApplications200Response GetApplications(ctx).Application(application).Execute()

Get Applications



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
	application := "Games" // string | Filter on the name of the application. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetApplications(context.Background()).Application(application).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplications`: GetApplications200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | **string** | Filter on the name of the application. | 

### Return type

[**GetApplications200Response**](GetApplications200Response.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationsAppDiscovery

> ApplicationList GetApplicationsAppDiscovery(ctx).Sources(sources).Identity(identity).Labels(labels).Controllable(controllable).WeightedRisk(weightedRisk).Categories(categories).AppTypes(appTypes).Date(date).Limit(limit).Offset(offset).Sort(sort).Order(order).LabelTimestamp(labelTimestamp).Execute()

List Applications



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
	sources := []string{"Sources_example"} // []string | Specify the types of log source to filter the collection. Valid values are: `dns`, `swg`, `cdfw`. (optional)
	identity := int64(12355) // int64 | Specify the entity ID to filter the collection. (optional)
	labels := []openapiclient.Label{openapiclient.Label("unreviewed")} // []Label | Specify the types of application classification to filter the collection. Valid values are: `unreviewed`, `approved`, `notApproved`, `underAudit`. (optional)
	controllable := "advanced" // string | Specify the type of controllable applications to filter the collection. Valid values are: `all`, `advanced`. (optional)
	weightedRisk := []openapiclient.WeightedRisk{openapiclient.WeightedRisk("veryLow")} // []WeightedRisk | Specify the list of application weighted risk to filter the collection. Valid values are: `veryLow`, `low`, `medium`, `high`, `veryHigh`. (optional)
	categories := []string{"Inner_example"} // []string | Specify the list of application category to filter the collection. (optional)
	appTypes := []openapiclient.AppType{openapiclient.AppType("saas")} // []AppType | Specify the types of application to filter the collection. Valid values are: `saas`, `paas`, and `iaas`. (optional)
	date := time.Now() // string | Specify a date to search for data within a twenty-four hour time period. If you do not provide a date, the last 90 days period is used to query the collection. (optional)
	limit := int64(56) // int64 | The maximum number of items to return in the collection. (optional)
	offset := int64(56) // int64 | The number of items to skip before starting to collect the result set. (optional)
	sort := "firstDetected" // string | Specify the name of a field to sort the applications. (optional)
	order := "asc" // string | Specify the order to sort the collection. Valid values are: `asc` (ascending) or `desc` (descending). (optional)
	labelTimestamp := time.Now() // time.Time | Filter the result set on the time when the label was updated. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetApplicationsAppDiscovery(context.Background()).Sources(sources).Identity(identity).Labels(labels).Controllable(controllable).WeightedRisk(weightedRisk).Categories(categories).AppTypes(appTypes).Date(date).Limit(limit).Offset(offset).Sort(sort).Order(order).LabelTimestamp(labelTimestamp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetApplicationsAppDiscovery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationsAppDiscovery`: ApplicationList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetApplicationsAppDiscovery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationsAppDiscoveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sources** | **[]string** | Specify the types of log source to filter the collection. Valid values are: &#x60;dns&#x60;, &#x60;swg&#x60;, &#x60;cdfw&#x60;. | 
 **identity** | **int64** | Specify the entity ID to filter the collection. | 
 **labels** | [**[]Label**](Label.md) | Specify the types of application classification to filter the collection. Valid values are: &#x60;unreviewed&#x60;, &#x60;approved&#x60;, &#x60;notApproved&#x60;, &#x60;underAudit&#x60;. | 
 **controllable** | **string** | Specify the type of controllable applications to filter the collection. Valid values are: &#x60;all&#x60;, &#x60;advanced&#x60;. | 
 **weightedRisk** | [**[]WeightedRisk**](WeightedRisk.md) | Specify the list of application weighted risk to filter the collection. Valid values are: &#x60;veryLow&#x60;, &#x60;low&#x60;, &#x60;medium&#x60;, &#x60;high&#x60;, &#x60;veryHigh&#x60;. | 
 **categories** | **[]string** | Specify the list of application category to filter the collection. | 
 **appTypes** | [**[]AppType**](AppType.md) | Specify the types of application to filter the collection. Valid values are: &#x60;saas&#x60;, &#x60;paas&#x60;, and &#x60;iaas&#x60;. | 
 **date** | **string** | Specify a date to search for data within a twenty-four hour time period. If you do not provide a date, the last 90 days period is used to query the collection. | 
 **limit** | **int64** | The maximum number of items to return in the collection. | 
 **offset** | **int64** | The number of items to skip before starting to collect the result set. | 
 **sort** | **string** | Specify the name of a field to sort the applications. | 
 **order** | **string** | Specify the order to sort the collection. Valid values are: &#x60;asc&#x60; (ascending) or &#x60;desc&#x60; (descending). | 
 **labelTimestamp** | **time.Time** | Filter the result set on the time when the label was updated. | 

### Return type

[**ApplicationList**](ApplicationList.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

