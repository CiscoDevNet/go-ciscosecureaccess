# CategoryWithLegacyId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the category. | 
**Legacyid** | **int64** | The legacy category ID. | 
**Label** | **string** | The label of the category. | 
**Type** | **string** | The type of the category. | 
**Integration** | **bool** | Specifies whether the category is an integration. | 
**Deprecated** | **bool** | Specifies whether the legacy category is deprecated. | 

## Methods

### NewCategoryWithLegacyId

`func NewCategoryWithLegacyId(id int64, legacyid int64, label string, type_ string, integration bool, deprecated bool, ) *CategoryWithLegacyId`

NewCategoryWithLegacyId instantiates a new CategoryWithLegacyId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryWithLegacyIdWithDefaults

`func NewCategoryWithLegacyIdWithDefaults() *CategoryWithLegacyId`

NewCategoryWithLegacyIdWithDefaults instantiates a new CategoryWithLegacyId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CategoryWithLegacyId) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CategoryWithLegacyId) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CategoryWithLegacyId) SetId(v int64)`

SetId sets Id field to given value.


### GetLegacyid

`func (o *CategoryWithLegacyId) GetLegacyid() int64`

GetLegacyid returns the Legacyid field if non-nil, zero value otherwise.

### GetLegacyidOk

`func (o *CategoryWithLegacyId) GetLegacyidOk() (*int64, bool)`

GetLegacyidOk returns a tuple with the Legacyid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyid

`func (o *CategoryWithLegacyId) SetLegacyid(v int64)`

SetLegacyid sets Legacyid field to given value.


### GetLabel

`func (o *CategoryWithLegacyId) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CategoryWithLegacyId) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CategoryWithLegacyId) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetType

`func (o *CategoryWithLegacyId) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CategoryWithLegacyId) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CategoryWithLegacyId) SetType(v string)`

SetType sets Type field to given value.


### GetIntegration

`func (o *CategoryWithLegacyId) GetIntegration() bool`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *CategoryWithLegacyId) GetIntegrationOk() (*bool, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *CategoryWithLegacyId) SetIntegration(v bool)`

SetIntegration sets Integration field to given value.


### GetDeprecated

`func (o *CategoryWithLegacyId) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *CategoryWithLegacyId) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *CategoryWithLegacyId) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


