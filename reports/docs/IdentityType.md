# IdentityType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the origin type for the identity. | [optional] 
**Label** | Pointer to **string** | The label of the origin type for the identity. | [optional] 
**Type** | Pointer to **string** | The name of the origin type for the identity. | [optional] 

## Methods

### NewIdentityType

`func NewIdentityType() *IdentityType`

NewIdentityType instantiates a new IdentityType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityTypeWithDefaults

`func NewIdentityTypeWithDefaults() *IdentityType`

NewIdentityTypeWithDefaults instantiates a new IdentityType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdentityType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *IdentityType) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *IdentityType) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *IdentityType) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *IdentityType) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *IdentityType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdentityType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdentityType) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IdentityType) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


