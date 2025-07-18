# SettingResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingValue** | Pointer to [**SettingValue**](SettingValue.md) |  | [optional] 
**SettingName** | Pointer to [**SettingName**](SettingName.md) |  | [optional] 
**CreatedAt** | Pointer to **string** | The date and time that the rule setting was created. | [optional] [readonly] 
**ModifiedAt** | Pointer to **string** | The date and time that the rule setting was modified. | [optional] [readonly] 

## Methods

### NewSettingResponseInner

`func NewSettingResponseInner() *SettingResponseInner`

NewSettingResponseInner instantiates a new SettingResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingResponseInnerWithDefaults

`func NewSettingResponseInnerWithDefaults() *SettingResponseInner`

NewSettingResponseInnerWithDefaults instantiates a new SettingResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettingValue

`func (o *SettingResponseInner) GetSettingValue() SettingValue`

GetSettingValue returns the SettingValue field if non-nil, zero value otherwise.

### GetSettingValueOk

`func (o *SettingResponseInner) GetSettingValueOk() (*SettingValue, bool)`

GetSettingValueOk returns a tuple with the SettingValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingValue

`func (o *SettingResponseInner) SetSettingValue(v SettingValue)`

SetSettingValue sets SettingValue field to given value.

### HasSettingValue

`func (o *SettingResponseInner) HasSettingValue() bool`

HasSettingValue returns a boolean if a field has been set.

### GetSettingName

`func (o *SettingResponseInner) GetSettingName() SettingName`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *SettingResponseInner) GetSettingNameOk() (*SettingName, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *SettingResponseInner) SetSettingName(v SettingName)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *SettingResponseInner) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SettingResponseInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SettingResponseInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SettingResponseInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SettingResponseInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *SettingResponseInner) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SettingResponseInner) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *SettingResponseInner) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *SettingResponseInner) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


