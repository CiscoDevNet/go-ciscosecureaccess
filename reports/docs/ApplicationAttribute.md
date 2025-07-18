# ApplicationAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the attribute. | [optional] 
**Name** | Pointer to **string** | The name of the attribute. | [optional] 
**Description** | Pointer to **string** | The description of the attribute. | [optional] 
**Values** | Pointer to [**[]ApplicationAttributeValue**](ApplicationAttributeValue.md) |  | [optional] 

## Methods

### NewApplicationAttribute

`func NewApplicationAttribute() *ApplicationAttribute`

NewApplicationAttribute instantiates a new ApplicationAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationAttributeWithDefaults

`func NewApplicationAttributeWithDefaults() *ApplicationAttribute`

NewApplicationAttributeWithDefaults instantiates a new ApplicationAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationAttribute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationAttribute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationAttribute) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationAttribute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApplicationAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationAttribute) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationAttribute) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationAttribute) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationAttribute) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationAttribute) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationAttribute) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetValues

`func (o *ApplicationAttribute) GetValues() []ApplicationAttributeValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ApplicationAttribute) GetValuesOk() (*[]ApplicationAttributeValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ApplicationAttribute) SetValues(v []ApplicationAttributeValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *ApplicationAttribute) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


