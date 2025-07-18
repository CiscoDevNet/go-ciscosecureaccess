# \RuleSettingsAndDefaultsAPI

All URIs are relative to *https://api.sse.cisco.com/policies/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPolicySetting**](RuleSettingsAndDefaultsAPI.md#GetPolicySetting) | **Get** /settings/{settingName} | Get Policy Setting
[**GetPolicySettingTypes**](RuleSettingsAndDefaultsAPI.md#GetPolicySettingTypes) | **Get** /settingTypes | List Policy Setting Types
[**GetPolicySettings**](RuleSettingsAndDefaultsAPI.md#GetPolicySettings) | **Get** /settings | Get Rule Defaults and Global Policy Settings
[**GetSettingTypesByName**](RuleSettingsAndDefaultsAPI.md#GetSettingTypesByName) | **Get** /settingTypes/{settingName} | Get Policy Setting Type
[**PutPolicySetting**](RuleSettingsAndDefaultsAPI.md#PutPolicySetting) | **Put** /settings/{settingName} | Update Policy Setting
[**PutPolicySettings**](RuleSettingsAndDefaultsAPI.md#PutPolicySettings) | **Put** /settings | Update Rule Defaults and Global Policy Settings



## GetPolicySetting

> SettingsResponseObject GetPolicySetting(ctx, settingName).Execute()

Get Policy Setting



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
	settingName := "sse.decryption.logPrivate" // string | The name of the global setting on the policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RuleSettingsAndDefaultsAPI.GetPolicySetting(context.Background(), settingName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RuleSettingsAndDefaultsAPI.GetPolicySetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicySetting`: SettingsResponseObject
	fmt.Fprintf(os.Stdout, "Response from `RuleSettingsAndDefaultsAPI.GetPolicySetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settingName** | **string** | The name of the global setting on the policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicySettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SettingsResponseObject**](SettingsResponseObject.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicySettingTypes

> []SettingTypesResponseInner GetPolicySettingTypes(ctx).Execute()

List Policy Setting Types



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
	resp, r, err := apiClient.RuleSettingsAndDefaultsAPI.GetPolicySettingTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RuleSettingsAndDefaultsAPI.GetPolicySettingTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicySettingTypes`: []SettingTypesResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RuleSettingsAndDefaultsAPI.GetPolicySettingTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicySettingTypesRequest struct via the builder pattern


### Return type

[**[]SettingTypesResponseInner**](SettingTypesResponseInner.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicySettings

> []SettingsResponseInner GetPolicySettings(ctx).Execute()

Get Rule Defaults and Global Policy Settings



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
	resp, r, err := apiClient.RuleSettingsAndDefaultsAPI.GetPolicySettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RuleSettingsAndDefaultsAPI.GetPolicySettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicySettings`: []SettingsResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RuleSettingsAndDefaultsAPI.GetPolicySettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicySettingsRequest struct via the builder pattern


### Return type

[**[]SettingsResponseInner**](SettingsResponseInner.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSettingTypesByName

> []SettingTypesResponseInner GetSettingTypesByName(ctx, settingName).Execute()

Get Policy Setting Type



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
	settingName := "sse.decryption.logPrivate" // string | The name of the global setting on the policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RuleSettingsAndDefaultsAPI.GetSettingTypesByName(context.Background(), settingName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RuleSettingsAndDefaultsAPI.GetSettingTypesByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettingTypesByName`: []SettingTypesResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RuleSettingsAndDefaultsAPI.GetSettingTypesByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settingName** | **string** | The name of the global setting on the policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingTypesByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SettingTypesResponseInner**](SettingTypesResponseInner.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPolicySetting

> SettingsResponseObject PutPolicySetting(ctx, settingName).SettingsRequestObject(settingsRequestObject).Execute()

Update Policy Setting



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
	settingName := "sse.decryption.logPrivate" // string | The name of the global setting on the policy.
	settingsRequestObject := *openapiclient.NewSettingsRequestObject() // SettingsRequestObject | Update the policy setting.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RuleSettingsAndDefaultsAPI.PutPolicySetting(context.Background(), settingName).SettingsRequestObject(settingsRequestObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RuleSettingsAndDefaultsAPI.PutPolicySetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPolicySetting`: SettingsResponseObject
	fmt.Fprintf(os.Stdout, "Response from `RuleSettingsAndDefaultsAPI.PutPolicySetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settingName** | **string** | The name of the global setting on the policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPolicySettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **settingsRequestObject** | [**SettingsRequestObject**](SettingsRequestObject.md) | Update the policy setting. | 

### Return type

[**SettingsResponseObject**](SettingsResponseObject.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPolicySettings

> []SettingResponseInner PutPolicySettings(ctx).SettingsRequestInner(settingsRequestInner).Execute()

Update Rule Defaults and Global Policy Settings



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
	settingsRequestInner := []openapiclient.SettingsRequestInner{*openapiclient.NewSettingsRequestInner()} // []SettingsRequestInner | Update the policy settings.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RuleSettingsAndDefaultsAPI.PutPolicySettings(context.Background()).SettingsRequestInner(settingsRequestInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RuleSettingsAndDefaultsAPI.PutPolicySettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPolicySettings`: []SettingResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RuleSettingsAndDefaultsAPI.PutPolicySettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutPolicySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **settingsRequestInner** | [**[]SettingsRequestInner**](SettingsRequestInner.md) | Update the policy settings. | 

### Return type

[**[]SettingResponseInner**](SettingResponseInner.md)

### Authorization

[oauthFlow](../README.md#oauthFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

