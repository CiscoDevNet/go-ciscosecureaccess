# ApplicationAttributeCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the attribute category. | [optional] 
**Name** | Pointer to **string** | The name of the attribute category. | [optional] 
**Attributes** | Pointer to [**[]ApplicationAttribute**](ApplicationAttribute.md) |  | [optional] 

## Methods

### NewApplicationAttributeCategory

`func NewApplicationAttributeCategory() *ApplicationAttributeCategory`

NewApplicationAttributeCategory instantiates a new ApplicationAttributeCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationAttributeCategoryWithDefaults

`func NewApplicationAttributeCategoryWithDefaults() *ApplicationAttributeCategory`

NewApplicationAttributeCategoryWithDefaults instantiates a new ApplicationAttributeCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationAttributeCategory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationAttributeCategory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationAttributeCategory) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationAttributeCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApplicationAttributeCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationAttributeCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationAttributeCategory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationAttributeCategory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAttributes

`func (o *ApplicationAttributeCategory) GetAttributes() []ApplicationAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationAttributeCategory) GetAttributesOk() (*[]ApplicationAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApplicationAttributeCategory) SetAttributes(v []ApplicationAttribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ApplicationAttributeCategory) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


