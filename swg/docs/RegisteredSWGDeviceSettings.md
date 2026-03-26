# RegisteredSWGDeviceSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int64** | The total number of devices that requested to update the device setting. | 
**SuccessCount** | **int64** | The number of devices that successfully changed the device setting. | 
**FailCount** | **int64** | The number of devices that failed to change the device setting. | 
**Items** | [**[]RegisteredSWGDeviceSettingsItemsInner**](RegisteredSWGDeviceSettingsItemsInner.md) | The list of device setting status properties. | 
**Value** | [**Value**](Value.md) |  | 

## Methods

### NewRegisteredSWGDeviceSettings

`func NewRegisteredSWGDeviceSettings(totalCount int64, successCount int64, failCount int64, items []RegisteredSWGDeviceSettingsItemsInner, value Value, ) *RegisteredSWGDeviceSettings`

NewRegisteredSWGDeviceSettings instantiates a new RegisteredSWGDeviceSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisteredSWGDeviceSettingsWithDefaults

`func NewRegisteredSWGDeviceSettingsWithDefaults() *RegisteredSWGDeviceSettings`

NewRegisteredSWGDeviceSettingsWithDefaults instantiates a new RegisteredSWGDeviceSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *RegisteredSWGDeviceSettings) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *RegisteredSWGDeviceSettings) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *RegisteredSWGDeviceSettings) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetSuccessCount

`func (o *RegisteredSWGDeviceSettings) GetSuccessCount() int64`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *RegisteredSWGDeviceSettings) GetSuccessCountOk() (*int64, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *RegisteredSWGDeviceSettings) SetSuccessCount(v int64)`

SetSuccessCount sets SuccessCount field to given value.


### GetFailCount

`func (o *RegisteredSWGDeviceSettings) GetFailCount() int64`

GetFailCount returns the FailCount field if non-nil, zero value otherwise.

### GetFailCountOk

`func (o *RegisteredSWGDeviceSettings) GetFailCountOk() (*int64, bool)`

GetFailCountOk returns a tuple with the FailCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailCount

`func (o *RegisteredSWGDeviceSettings) SetFailCount(v int64)`

SetFailCount sets FailCount field to given value.


### GetItems

`func (o *RegisteredSWGDeviceSettings) GetItems() []RegisteredSWGDeviceSettingsItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RegisteredSWGDeviceSettings) GetItemsOk() (*[]RegisteredSWGDeviceSettingsItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RegisteredSWGDeviceSettings) SetItems(v []RegisteredSWGDeviceSettingsItemsInner)`

SetItems sets Items field to given value.


### GetValue

`func (o *RegisteredSWGDeviceSettings) GetValue() Value`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RegisteredSWGDeviceSettings) GetValueOk() (*Value, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RegisteredSWGDeviceSettings) SetValue(v Value)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


