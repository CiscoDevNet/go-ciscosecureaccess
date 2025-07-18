# BrowserBasedAccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the connection is a browser. | 
**Protocol** | [**ProtocolProxyToResource**](ProtocolProxyToResource.md) |  | 
**Sni** | Pointer to **string** | The Server Name Indication (SNI) domain name. Only applicable for browser-based Zero Trust Access and when you select the HTTPS protocol. The SNI must be a valid domain. | [optional] 
**SslVerificationEnabled** | Pointer to **bool** | Specify whether to enable upstream SSL verification for the internally hosted URL by the customer. Applicable for browser-based Zero Trust Access only and when you select the HTTPS protocol. | [optional] [default to true]
**ExternalFQDN** | Pointer to **string** | The external fully qualified domain name (FQDN). | [optional] 

## Methods

### NewBrowserBasedAccessResponse

`func NewBrowserBasedAccessResponse(type_ string, protocol ProtocolProxyToResource, ) *BrowserBasedAccessResponse`

NewBrowserBasedAccessResponse instantiates a new BrowserBasedAccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrowserBasedAccessResponseWithDefaults

`func NewBrowserBasedAccessResponseWithDefaults() *BrowserBasedAccessResponse`

NewBrowserBasedAccessResponseWithDefaults instantiates a new BrowserBasedAccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BrowserBasedAccessResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BrowserBasedAccessResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BrowserBasedAccessResponse) SetType(v string)`

SetType sets Type field to given value.


### GetProtocol

`func (o *BrowserBasedAccessResponse) GetProtocol() ProtocolProxyToResource`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *BrowserBasedAccessResponse) GetProtocolOk() (*ProtocolProxyToResource, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *BrowserBasedAccessResponse) SetProtocol(v ProtocolProxyToResource)`

SetProtocol sets Protocol field to given value.


### GetSni

`func (o *BrowserBasedAccessResponse) GetSni() string`

GetSni returns the Sni field if non-nil, zero value otherwise.

### GetSniOk

`func (o *BrowserBasedAccessResponse) GetSniOk() (*string, bool)`

GetSniOk returns a tuple with the Sni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSni

`func (o *BrowserBasedAccessResponse) SetSni(v string)`

SetSni sets Sni field to given value.

### HasSni

`func (o *BrowserBasedAccessResponse) HasSni() bool`

HasSni returns a boolean if a field has been set.

### GetSslVerificationEnabled

`func (o *BrowserBasedAccessResponse) GetSslVerificationEnabled() bool`

GetSslVerificationEnabled returns the SslVerificationEnabled field if non-nil, zero value otherwise.

### GetSslVerificationEnabledOk

`func (o *BrowserBasedAccessResponse) GetSslVerificationEnabledOk() (*bool, bool)`

GetSslVerificationEnabledOk returns a tuple with the SslVerificationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVerificationEnabled

`func (o *BrowserBasedAccessResponse) SetSslVerificationEnabled(v bool)`

SetSslVerificationEnabled sets SslVerificationEnabled field to given value.

### HasSslVerificationEnabled

`func (o *BrowserBasedAccessResponse) HasSslVerificationEnabled() bool`

HasSslVerificationEnabled returns a boolean if a field has been set.

### GetExternalFQDN

`func (o *BrowserBasedAccessResponse) GetExternalFQDN() string`

GetExternalFQDN returns the ExternalFQDN field if non-nil, zero value otherwise.

### GetExternalFQDNOk

`func (o *BrowserBasedAccessResponse) GetExternalFQDNOk() (*string, bool)`

GetExternalFQDNOk returns a tuple with the ExternalFQDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFQDN

`func (o *BrowserBasedAccessResponse) SetExternalFQDN(v string)`

SetExternalFQDN sets ExternalFQDN field to given value.

### HasExternalFQDN

`func (o *BrowserBasedAccessResponse) HasExternalFQDN() bool`

HasExternalFQDN returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


