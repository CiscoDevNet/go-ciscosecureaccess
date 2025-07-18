# ClientBasedAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the connection is Secure Client. | 
**ReachableAddresses** | **[]string** | The list of IP address, CIDRs, FQDN, or wildcard FQDN destinations. IPv6 is not supported. Only applies to resources that are configured for client-based Zero Trust Access. | 

## Methods

### NewClientBasedAccess

`func NewClientBasedAccess(type_ string, reachableAddresses []string, ) *ClientBasedAccess`

NewClientBasedAccess instantiates a new ClientBasedAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientBasedAccessWithDefaults

`func NewClientBasedAccessWithDefaults() *ClientBasedAccess`

NewClientBasedAccessWithDefaults instantiates a new ClientBasedAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ClientBasedAccess) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientBasedAccess) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientBasedAccess) SetType(v string)`

SetType sets Type field to given value.


### GetReachableAddresses

`func (o *ClientBasedAccess) GetReachableAddresses() []string`

GetReachableAddresses returns the ReachableAddresses field if non-nil, zero value otherwise.

### GetReachableAddressesOk

`func (o *ClientBasedAccess) GetReachableAddressesOk() (*[]string, bool)`

GetReachableAddressesOk returns a tuple with the ReachableAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachableAddresses

`func (o *ClientBasedAccess) SetReachableAddresses(v []string)`

SetReachableAddresses sets ReachableAddresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


