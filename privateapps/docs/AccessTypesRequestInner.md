# AccessTypesRequestInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the connection is Secure Client. | 
**ReachableAddresses** | **[]string** | The list of IP address, CIDRs, FQDN, or wildcard FQDN destinations. IPv6 is not supported. Only applies to resources that are configured for client-based Zero Trust Access. | 
**Protocol** | [**ProtocolProxyToResource**](ProtocolProxyToResource.md) |  | 
**Sni** | Pointer to **string** | The Server Name Indication (SNI) domain name. Only applicable for browser-based Zero Trust Access and when you select the HTTPS protocol. The SNI must be a valid domain. | [optional] 
**SslVerificationEnabled** | Pointer to **bool** | Specify whether to enable upstream SSL verification for the internally hosted URL by the customer. Applicable for browser-based Zero Trust Access only and when you select the HTTPS protocol. | [optional] [default to true]
**ExternalFQDNPrefix** | Pointer to **string** | The external fully qualified domain name (FQDN) prefix. | [optional] 

## Methods

### NewAccessTypesRequestInner

`func NewAccessTypesRequestInner(type_ string, reachableAddresses []string, protocol ProtocolProxyToResource, ) *AccessTypesRequestInner`

NewAccessTypesRequestInner instantiates a new AccessTypesRequestInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTypesRequestInnerWithDefaults

`func NewAccessTypesRequestInnerWithDefaults() *AccessTypesRequestInner`

NewAccessTypesRequestInnerWithDefaults instantiates a new AccessTypesRequestInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AccessTypesRequestInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccessTypesRequestInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccessTypesRequestInner) SetType(v string)`

SetType sets Type field to given value.


### GetReachableAddresses

`func (o *AccessTypesRequestInner) GetReachableAddresses() []string`

GetReachableAddresses returns the ReachableAddresses field if non-nil, zero value otherwise.

### GetReachableAddressesOk

`func (o *AccessTypesRequestInner) GetReachableAddressesOk() (*[]string, bool)`

GetReachableAddressesOk returns a tuple with the ReachableAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachableAddresses

`func (o *AccessTypesRequestInner) SetReachableAddresses(v []string)`

SetReachableAddresses sets ReachableAddresses field to given value.


### GetProtocol

`func (o *AccessTypesRequestInner) GetProtocol() ProtocolProxyToResource`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *AccessTypesRequestInner) GetProtocolOk() (*ProtocolProxyToResource, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *AccessTypesRequestInner) SetProtocol(v ProtocolProxyToResource)`

SetProtocol sets Protocol field to given value.


### GetSni

`func (o *AccessTypesRequestInner) GetSni() string`

GetSni returns the Sni field if non-nil, zero value otherwise.

### GetSniOk

`func (o *AccessTypesRequestInner) GetSniOk() (*string, bool)`

GetSniOk returns a tuple with the Sni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSni

`func (o *AccessTypesRequestInner) SetSni(v string)`

SetSni sets Sni field to given value.

### HasSni

`func (o *AccessTypesRequestInner) HasSni() bool`

HasSni returns a boolean if a field has been set.

### GetSslVerificationEnabled

`func (o *AccessTypesRequestInner) GetSslVerificationEnabled() bool`

GetSslVerificationEnabled returns the SslVerificationEnabled field if non-nil, zero value otherwise.

### GetSslVerificationEnabledOk

`func (o *AccessTypesRequestInner) GetSslVerificationEnabledOk() (*bool, bool)`

GetSslVerificationEnabledOk returns a tuple with the SslVerificationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVerificationEnabled

`func (o *AccessTypesRequestInner) SetSslVerificationEnabled(v bool)`

SetSslVerificationEnabled sets SslVerificationEnabled field to given value.

### HasSslVerificationEnabled

`func (o *AccessTypesRequestInner) HasSslVerificationEnabled() bool`

HasSslVerificationEnabled returns a boolean if a field has been set.

### GetExternalFQDNPrefix

`func (o *AccessTypesRequestInner) GetExternalFQDNPrefix() string`

GetExternalFQDNPrefix returns the ExternalFQDNPrefix field if non-nil, zero value otherwise.

### GetExternalFQDNPrefixOk

`func (o *AccessTypesRequestInner) GetExternalFQDNPrefixOk() (*string, bool)`

GetExternalFQDNPrefixOk returns a tuple with the ExternalFQDNPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFQDNPrefix

`func (o *AccessTypesRequestInner) SetExternalFQDNPrefix(v string)`

SetExternalFQDNPrefix sets ExternalFQDNPrefix field to given value.

### HasExternalFQDNPrefix

`func (o *AccessTypesRequestInner) HasExternalFQDNPrefix() bool`

HasExternalFQDNPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


