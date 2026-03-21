# ListSWGDeviceSettingsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginId** | **int64** | The origin ID of the device. | 
**Name** | **string** | The name of the device setting. | 
**Value** | [**Value**](Value.md) |  | 
**ModifiedAt** | **string** | The date and time when the settings on the device were modified. The timestamp is in the ISO 8601 date format. | 

## Methods

### NewListSWGDeviceSettingsInner

`func NewListSWGDeviceSettingsInner(originId int64, name string, value Value, modifiedAt string, ) *ListSWGDeviceSettingsInner`

NewListSWGDeviceSettingsInner instantiates a new ListSWGDeviceSettingsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSWGDeviceSettingsInnerWithDefaults

`func NewListSWGDeviceSettingsInnerWithDefaults() *ListSWGDeviceSettingsInner`

NewListSWGDeviceSettingsInnerWithDefaults instantiates a new ListSWGDeviceSettingsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginId

`func (o *ListSWGDeviceSettingsInner) GetOriginId() int64`

GetOriginId returns the OriginId field if non-nil, zero value otherwise.

### GetOriginIdOk

`func (o *ListSWGDeviceSettingsInner) GetOriginIdOk() (*int64, bool)`

GetOriginIdOk returns a tuple with the OriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginId

`func (o *ListSWGDeviceSettingsInner) SetOriginId(v int64)`

SetOriginId sets OriginId field to given value.


### GetName

`func (o *ListSWGDeviceSettingsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListSWGDeviceSettingsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListSWGDeviceSettingsInner) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *ListSWGDeviceSettingsInner) GetValue() Value`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListSWGDeviceSettingsInner) GetValueOk() (*Value, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListSWGDeviceSettingsInner) SetValue(v Value)`

SetValue sets Value field to given value.


### GetModifiedAt

`func (o *ListSWGDeviceSettingsInner) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ListSWGDeviceSettingsInner) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ListSWGDeviceSettingsInner) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


