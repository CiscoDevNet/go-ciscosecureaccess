# UpdateNetworkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the network. | 
**IpAddress** | Pointer to **string** | The IP address of the network. | [optional] 
**PrefixLength** | Pointer to **int64** | The length of the prefix. Set a prefix length that is greater than 28 and less than 33. | [optional] 
**IsDynamic** | **bool** | Specifies whether the IP address is dynamic. | 
**Status** | **string** | The status of the network. | 

## Methods

### NewUpdateNetworkRequest

`func NewUpdateNetworkRequest(name string, isDynamic bool, status string, ) *UpdateNetworkRequest`

NewUpdateNetworkRequest instantiates a new UpdateNetworkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkRequestWithDefaults

`func NewUpdateNetworkRequestWithDefaults() *UpdateNetworkRequest`

NewUpdateNetworkRequestWithDefaults instantiates a new UpdateNetworkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNetworkRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIpAddress

`func (o *UpdateNetworkRequest) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *UpdateNetworkRequest) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *UpdateNetworkRequest) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *UpdateNetworkRequest) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetPrefixLength

`func (o *UpdateNetworkRequest) GetPrefixLength() int64`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *UpdateNetworkRequest) GetPrefixLengthOk() (*int64, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *UpdateNetworkRequest) SetPrefixLength(v int64)`

SetPrefixLength sets PrefixLength field to given value.

### HasPrefixLength

`func (o *UpdateNetworkRequest) HasPrefixLength() bool`

HasPrefixLength returns a boolean if a field has been set.

### GetIsDynamic

`func (o *UpdateNetworkRequest) GetIsDynamic() bool`

GetIsDynamic returns the IsDynamic field if non-nil, zero value otherwise.

### GetIsDynamicOk

`func (o *UpdateNetworkRequest) GetIsDynamicOk() (*bool, bool)`

GetIsDynamicOk returns a tuple with the IsDynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDynamic

`func (o *UpdateNetworkRequest) SetIsDynamic(v bool)`

SetIsDynamic sets IsDynamic field to given value.


### GetStatus

`func (o *UpdateNetworkRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateNetworkRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateNetworkRequest) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


