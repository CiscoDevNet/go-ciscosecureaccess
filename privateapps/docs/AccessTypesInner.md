# AccessTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the connection is branch. | 
**Protocol** | [**ProtocolProxyToResource**](ProtocolProxyToResource.md) |  | 
**Sni** | Pointer to **string** | The Server Name Indication (SNI) domain name. Only applicable for browser-based Zero Trust Access and when you select the HTTPS protocol. The SNI must be a valid domain. | [optional] 
**SslVerificationEnabled** | Pointer to **bool** | Specify whether to enable upstream SSL verification for the internally hosted URL by the customer. Applicable for browser-based Zero Trust Access only and when you select the HTTPS protocol. | [optional] [default to true]
**ExternalFQDN** | Pointer to **string** | The external fully qualified domain name (FQDN). | [optional] 
**ReachableAddresses** | **[]string** | The list of IP address, CIDRs, FQDN, or wildcard FQDN destinations. IPv6 is not supported. Only applies to resources that are configured for client-based Zero Trust Access. | 

## Methods

### NewAccessTypesInner

`func NewAccessTypesInner(type_ string, protocol ProtocolProxyToResource, reachableAddresses []string, ) *AccessTypesInner`

NewAccessTypesInner instantiates a new AccessTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTypesInnerWithDefaults

`func NewAccessTypesInnerWithDefaults() *AccessTypesInner`

NewAccessTypesInnerWithDefaults instantiates a new AccessTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AccessTypesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccessTypesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccessTypesInner) SetType(v string)`

SetType sets Type field to given value.


### GetProtocol

`func (o *AccessTypesInner) GetProtocol() ProtocolProxyToResource`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *AccessTypesInner) GetProtocolOk() (*ProtocolProxyToResource, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *AccessTypesInner) SetProtocol(v ProtocolProxyToResource)`

SetProtocol sets Protocol field to given value.


### GetSni

`func (o *AccessTypesInner) GetSni() string`

GetSni returns the Sni field if non-nil, zero value otherwise.

### GetSniOk

`func (o *AccessTypesInner) GetSniOk() (*string, bool)`

GetSniOk returns a tuple with the Sni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSni

`func (o *AccessTypesInner) SetSni(v string)`

SetSni sets Sni field to given value.

### HasSni

`func (o *AccessTypesInner) HasSni() bool`

HasSni returns a boolean if a field has been set.

### GetSslVerificationEnabled

`func (o *AccessTypesInner) GetSslVerificationEnabled() bool`

GetSslVerificationEnabled returns the SslVerificationEnabled field if non-nil, zero value otherwise.

### GetSslVerificationEnabledOk

`func (o *AccessTypesInner) GetSslVerificationEnabledOk() (*bool, bool)`

GetSslVerificationEnabledOk returns a tuple with the SslVerificationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVerificationEnabled

`func (o *AccessTypesInner) SetSslVerificationEnabled(v bool)`

SetSslVerificationEnabled sets SslVerificationEnabled field to given value.

### HasSslVerificationEnabled

`func (o *AccessTypesInner) HasSslVerificationEnabled() bool`

HasSslVerificationEnabled returns a boolean if a field has been set.

### GetExternalFQDN

`func (o *AccessTypesInner) GetExternalFQDN() string`

GetExternalFQDN returns the ExternalFQDN field if non-nil, zero value otherwise.

### GetExternalFQDNOk

`func (o *AccessTypesInner) GetExternalFQDNOk() (*string, bool)`

GetExternalFQDNOk returns a tuple with the ExternalFQDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFQDN

`func (o *AccessTypesInner) SetExternalFQDN(v string)`

SetExternalFQDN sets ExternalFQDN field to given value.

### HasExternalFQDN

`func (o *AccessTypesInner) HasExternalFQDN() bool`

HasExternalFQDN returns a boolean if a field has been set.

### GetReachableAddresses

`func (o *AccessTypesInner) GetReachableAddresses() []string`

GetReachableAddresses returns the ReachableAddresses field if non-nil, zero value otherwise.

### GetReachableAddressesOk

`func (o *AccessTypesInner) GetReachableAddressesOk() (*[]string, bool)`

GetReachableAddressesOk returns a tuple with the ReachableAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachableAddresses

`func (o *AccessTypesInner) SetReachableAddresses(v []string)`

SetReachableAddresses sets ReachableAddresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


