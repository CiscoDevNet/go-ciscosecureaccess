# ApplicationAttributeValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the attribute value. | [optional] 
**Value** | Pointer to **string** | The value of the attribute value. | [optional] 
**CreatedOn** | Pointer to **time.Time** | The date when the application attribute was created. | [optional] 
**UpdatedOn** | Pointer to **time.Time** | The date when the application attribute was updated. | [optional] 

## Methods

### NewApplicationAttributeValue

`func NewApplicationAttributeValue() *ApplicationAttributeValue`

NewApplicationAttributeValue instantiates a new ApplicationAttributeValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationAttributeValueWithDefaults

`func NewApplicationAttributeValueWithDefaults() *ApplicationAttributeValue`

NewApplicationAttributeValueWithDefaults instantiates a new ApplicationAttributeValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationAttributeValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationAttributeValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationAttributeValue) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationAttributeValue) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *ApplicationAttributeValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ApplicationAttributeValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ApplicationAttributeValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ApplicationAttributeValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCreatedOn

`func (o *ApplicationAttributeValue) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *ApplicationAttributeValue) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *ApplicationAttributeValue) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *ApplicationAttributeValue) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *ApplicationAttributeValue) GetUpdatedOn() time.Time`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *ApplicationAttributeValue) GetUpdatedOnOk() (*time.Time, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *ApplicationAttributeValue) SetUpdatedOn(v time.Time)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *ApplicationAttributeValue) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


