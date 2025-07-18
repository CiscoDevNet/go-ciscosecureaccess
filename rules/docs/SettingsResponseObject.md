# SettingsResponseObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingName** | [**SettingName**](SettingName.md) |  | 
**SettingValue** | [**SettingValue**](SettingValue.md) |  | 
**CreatedAt** | **string** | The date and time that the rule setting was created. | [readonly] 
**ModifiedAt** | **string** | The date and time that the rule setting was modified. | [readonly] 

## Methods

### NewSettingsResponseObject

`func NewSettingsResponseObject(settingName SettingName, settingValue SettingValue, createdAt string, modifiedAt string, ) *SettingsResponseObject`

NewSettingsResponseObject instantiates a new SettingsResponseObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsResponseObjectWithDefaults

`func NewSettingsResponseObjectWithDefaults() *SettingsResponseObject`

NewSettingsResponseObjectWithDefaults instantiates a new SettingsResponseObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettingName

`func (o *SettingsResponseObject) GetSettingName() SettingName`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *SettingsResponseObject) GetSettingNameOk() (*SettingName, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *SettingsResponseObject) SetSettingName(v SettingName)`

SetSettingName sets SettingName field to given value.


### GetSettingValue

`func (o *SettingsResponseObject) GetSettingValue() SettingValue`

GetSettingValue returns the SettingValue field if non-nil, zero value otherwise.

### GetSettingValueOk

`func (o *SettingsResponseObject) GetSettingValueOk() (*SettingValue, bool)`

GetSettingValueOk returns a tuple with the SettingValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingValue

`func (o *SettingsResponseObject) SetSettingValue(v SettingValue)`

SetSettingValue sets SettingValue field to given value.


### GetCreatedAt

`func (o *SettingsResponseObject) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SettingsResponseObject) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SettingsResponseObject) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *SettingsResponseObject) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SettingsResponseObject) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *SettingsResponseObject) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


