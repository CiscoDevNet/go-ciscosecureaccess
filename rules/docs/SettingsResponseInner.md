# SettingsResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingName** | [**SettingName**](SettingName.md) |  | 
**SettingValue** | [**SettingValue**](SettingValue.md) |  | 
**CreatedAt** | **string** | The date and time that the rule was created. | [readonly] 
**ModifiedAt** | **string** | The date and time that the rule was modified. | [readonly] 
**SettingId** | **int64** |  | 
**IsGlobal** | **bool** |  | 

## Methods

### NewSettingsResponseInner

`func NewSettingsResponseInner(settingName SettingName, settingValue SettingValue, createdAt string, modifiedAt string, settingId int64, isGlobal bool, ) *SettingsResponseInner`

NewSettingsResponseInner instantiates a new SettingsResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsResponseInnerWithDefaults

`func NewSettingsResponseInnerWithDefaults() *SettingsResponseInner`

NewSettingsResponseInnerWithDefaults instantiates a new SettingsResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettingName

`func (o *SettingsResponseInner) GetSettingName() SettingName`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *SettingsResponseInner) GetSettingNameOk() (*SettingName, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *SettingsResponseInner) SetSettingName(v SettingName)`

SetSettingName sets SettingName field to given value.


### GetSettingValue

`func (o *SettingsResponseInner) GetSettingValue() SettingValue`

GetSettingValue returns the SettingValue field if non-nil, zero value otherwise.

### GetSettingValueOk

`func (o *SettingsResponseInner) GetSettingValueOk() (*SettingValue, bool)`

GetSettingValueOk returns a tuple with the SettingValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingValue

`func (o *SettingsResponseInner) SetSettingValue(v SettingValue)`

SetSettingValue sets SettingValue field to given value.


### GetCreatedAt

`func (o *SettingsResponseInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SettingsResponseInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SettingsResponseInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *SettingsResponseInner) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SettingsResponseInner) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *SettingsResponseInner) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.


### GetSettingId

`func (o *SettingsResponseInner) GetSettingId() int64`

GetSettingId returns the SettingId field if non-nil, zero value otherwise.

### GetSettingIdOk

`func (o *SettingsResponseInner) GetSettingIdOk() (*int64, bool)`

GetSettingIdOk returns a tuple with the SettingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingId

`func (o *SettingsResponseInner) SetSettingId(v int64)`

SetSettingId sets SettingId field to given value.


### GetIsGlobal

`func (o *SettingsResponseInner) GetIsGlobal() bool`

GetIsGlobal returns the IsGlobal field if non-nil, zero value otherwise.

### GetIsGlobalOk

`func (o *SettingsResponseInner) GetIsGlobalOk() (*bool, bool)`

GetIsGlobalOk returns a tuple with the IsGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobal

`func (o *SettingsResponseInner) SetIsGlobal(v bool)`

SetIsGlobal sets IsGlobal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


