# BrowserBasedAccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the connection is a browser. | 
**Protocol** | [**ProtocolProxyToResource**](ProtocolProxyToResource.md) |  | 
**Sni** | Pointer to **string** | The Server Name Indication (SNI) domain name. Only applicable for browser-based Zero Trust Access and when you select the HTTPS protocol. The SNI must be a valid domain. | [optional] 
**SslVerificationEnabled** | Pointer to **bool** | Specify whether to enable upstream SSL verification for the internally hosted URL by the customer. Applicable for browser-based Zero Trust Access only and when you select the HTTPS protocol. | [optional] [default to true]
**ExternalFQDNPrefix** | Pointer to **string** | The external fully qualified domain name (FQDN) prefix. | [optional] 

## Methods

### NewBrowserBasedAccessRequest

`func NewBrowserBasedAccessRequest(type_ string, protocol ProtocolProxyToResource, ) *BrowserBasedAccessRequest`

NewBrowserBasedAccessRequest instantiates a new BrowserBasedAccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrowserBasedAccessRequestWithDefaults

`func NewBrowserBasedAccessRequestWithDefaults() *BrowserBasedAccessRequest`

NewBrowserBasedAccessRequestWithDefaults instantiates a new BrowserBasedAccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BrowserBasedAccessRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BrowserBasedAccessRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BrowserBasedAccessRequest) SetType(v string)`

SetType sets Type field to given value.


### GetProtocol

`func (o *BrowserBasedAccessRequest) GetProtocol() ProtocolProxyToResource`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *BrowserBasedAccessRequest) GetProtocolOk() (*ProtocolProxyToResource, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *BrowserBasedAccessRequest) SetProtocol(v ProtocolProxyToResource)`

SetProtocol sets Protocol field to given value.


### GetSni

`func (o *BrowserBasedAccessRequest) GetSni() string`

GetSni returns the Sni field if non-nil, zero value otherwise.

### GetSniOk

`func (o *BrowserBasedAccessRequest) GetSniOk() (*string, bool)`

GetSniOk returns a tuple with the Sni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSni

`func (o *BrowserBasedAccessRequest) SetSni(v string)`

SetSni sets Sni field to given value.

### HasSni

`func (o *BrowserBasedAccessRequest) HasSni() bool`

HasSni returns a boolean if a field has been set.

### GetSslVerificationEnabled

`func (o *BrowserBasedAccessRequest) GetSslVerificationEnabled() bool`

GetSslVerificationEnabled returns the SslVerificationEnabled field if non-nil, zero value otherwise.

### GetSslVerificationEnabledOk

`func (o *BrowserBasedAccessRequest) GetSslVerificationEnabledOk() (*bool, bool)`

GetSslVerificationEnabledOk returns a tuple with the SslVerificationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVerificationEnabled

`func (o *BrowserBasedAccessRequest) SetSslVerificationEnabled(v bool)`

SetSslVerificationEnabled sets SslVerificationEnabled field to given value.

### HasSslVerificationEnabled

`func (o *BrowserBasedAccessRequest) HasSslVerificationEnabled() bool`

HasSslVerificationEnabled returns a boolean if a field has been set.

### GetExternalFQDNPrefix

`func (o *BrowserBasedAccessRequest) GetExternalFQDNPrefix() string`

GetExternalFQDNPrefix returns the ExternalFQDNPrefix field if non-nil, zero value otherwise.

### GetExternalFQDNPrefixOk

`func (o *BrowserBasedAccessRequest) GetExternalFQDNPrefixOk() (*string, bool)`

GetExternalFQDNPrefixOk returns a tuple with the ExternalFQDNPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFQDNPrefix

`func (o *BrowserBasedAccessRequest) SetExternalFQDNPrefix(v string)`

SetExternalFQDNPrefix sets ExternalFQDNPrefix field to given value.

### HasExternalFQDNPrefix

`func (o *BrowserBasedAccessRequest) HasExternalFQDNPrefix() bool`

HasExternalFQDNPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


