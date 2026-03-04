# NetworkObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginId** | **int64** | The origin ID. | 
**Name** | **string** | The name of the network. | 
**IpAddress** | **string** | The IP address of the network. | 
**PrefixLength** | **int64** | The length of the prefix. Set a prefix length that is greater than 28 and less than 33. | 
**IsDynamic** | **bool** | Specifies whether the network has a dynamic IP. | 
**IsVerified** | **bool** | Specifies whether the network is verified. | 
**Status** | **string** | The status of the network. | 
**CreatedAt** | **time.Time** | The date and time (timestamp) when the network was created. | 

## Methods

### NewNetworkObject

`func NewNetworkObject(originId int64, name string, ipAddress string, prefixLength int64, isDynamic bool, isVerified bool, status string, createdAt time.Time, ) *NetworkObject`

NewNetworkObject instantiates a new NetworkObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkObjectWithDefaults

`func NewNetworkObjectWithDefaults() *NetworkObject`

NewNetworkObjectWithDefaults instantiates a new NetworkObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginId

`func (o *NetworkObject) GetOriginId() int64`

GetOriginId returns the OriginId field if non-nil, zero value otherwise.

### GetOriginIdOk

`func (o *NetworkObject) GetOriginIdOk() (*int64, bool)`

GetOriginIdOk returns a tuple with the OriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginId

`func (o *NetworkObject) SetOriginId(v int64)`

SetOriginId sets OriginId field to given value.


### GetName

`func (o *NetworkObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkObject) SetName(v string)`

SetName sets Name field to given value.


### GetIpAddress

`func (o *NetworkObject) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *NetworkObject) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *NetworkObject) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetPrefixLength

`func (o *NetworkObject) GetPrefixLength() int64`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *NetworkObject) GetPrefixLengthOk() (*int64, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *NetworkObject) SetPrefixLength(v int64)`

SetPrefixLength sets PrefixLength field to given value.


### GetIsDynamic

`func (o *NetworkObject) GetIsDynamic() bool`

GetIsDynamic returns the IsDynamic field if non-nil, zero value otherwise.

### GetIsDynamicOk

`func (o *NetworkObject) GetIsDynamicOk() (*bool, bool)`

GetIsDynamicOk returns a tuple with the IsDynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDynamic

`func (o *NetworkObject) SetIsDynamic(v bool)`

SetIsDynamic sets IsDynamic field to given value.


### GetIsVerified

`func (o *NetworkObject) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *NetworkObject) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *NetworkObject) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.


### GetStatus

`func (o *NetworkObject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkObject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkObject) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *NetworkObject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkObject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkObject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


