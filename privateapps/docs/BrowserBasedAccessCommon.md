# BrowserBasedAccessCommon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the connection is a browser. | 
**Protocol** | [**ProtocolProxyToResource**](ProtocolProxyToResource.md) |  | 
**Sni** | Pointer to **string** | The Server Name Indication (SNI) domain name. Only applicable for browser-based Zero Trust Access and when you select the HTTPS protocol. The SNI must be a valid domain. | [optional] 
**SslVerificationEnabled** | Pointer to **bool** | Specify whether to enable upstream SSL verification for the internally hosted URL by the customer. Applicable for browser-based Zero Trust Access only and when you select the HTTPS protocol. | [optional] [default to true]

## Methods

### NewBrowserBasedAccessCommon

`func NewBrowserBasedAccessCommon(type_ string, protocol ProtocolProxyToResource, ) *BrowserBasedAccessCommon`

NewBrowserBasedAccessCommon instantiates a new BrowserBasedAccessCommon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrowserBasedAccessCommonWithDefaults

`func NewBrowserBasedAccessCommonWithDefaults() *BrowserBasedAccessCommon`

NewBrowserBasedAccessCommonWithDefaults instantiates a new BrowserBasedAccessCommon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BrowserBasedAccessCommon) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BrowserBasedAccessCommon) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BrowserBasedAccessCommon) SetType(v string)`

SetType sets Type field to given value.


### GetProtocol

`func (o *BrowserBasedAccessCommon) GetProtocol() ProtocolProxyToResource`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *BrowserBasedAccessCommon) GetProtocolOk() (*ProtocolProxyToResource, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *BrowserBasedAccessCommon) SetProtocol(v ProtocolProxyToResource)`

SetProtocol sets Protocol field to given value.


### GetSni

`func (o *BrowserBasedAccessCommon) GetSni() string`

GetSni returns the Sni field if non-nil, zero value otherwise.

### GetSniOk

`func (o *BrowserBasedAccessCommon) GetSniOk() (*string, bool)`

GetSniOk returns a tuple with the Sni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSni

`func (o *BrowserBasedAccessCommon) SetSni(v string)`

SetSni sets Sni field to given value.

### HasSni

`func (o *BrowserBasedAccessCommon) HasSni() bool`

HasSni returns a boolean if a field has been set.

### GetSslVerificationEnabled

`func (o *BrowserBasedAccessCommon) GetSslVerificationEnabled() bool`

GetSslVerificationEnabled returns the SslVerificationEnabled field if non-nil, zero value otherwise.

### GetSslVerificationEnabledOk

`func (o *BrowserBasedAccessCommon) GetSslVerificationEnabledOk() (*bool, bool)`

GetSslVerificationEnabledOk returns a tuple with the SslVerificationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVerificationEnabled

`func (o *BrowserBasedAccessCommon) SetSslVerificationEnabled(v bool)`

SetSslVerificationEnabled sets SslVerificationEnabled field to given value.

### HasSslVerificationEnabled

`func (o *BrowserBasedAccessCommon) HasSslVerificationEnabled() bool`

HasSslVerificationEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


