# \OrganizationInformationAPI

All URIs are relative to *https://api.sse.cisco.com/deployments/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationInfo**](OrganizationInformationAPI.md#GetOrganizationInfo) | **Get** /roamingcomputers/orgInfo | Get OrgInfo Properties



## GetOrganizationInfo

> OrgInfo GetOrganizationInfo(ctx).Execute()

Get OrgInfo Properties



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/go-ciscosecureaccess/roaming"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationInformationAPI.GetOrganizationInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationInformationAPI.GetOrganizationInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationInfo`: OrgInfo
	fmt.Fprintf(os.Stdout, "Response from `OrganizationInformationAPI.GetOrganizationInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInfoRequest struct via the builder pattern


### Return type

[**OrgInfo**](OrgInfo.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

