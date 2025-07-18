# ZtnaApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | id of the application | [optional] 
**Label** | Pointer to **string** | label of the application | [optional] 
**Category** | Pointer to [**ZtnaApplicationCategory**](ZtnaApplicationCategory.md) |  | [optional] 

## Methods

### NewZtnaApplication

`func NewZtnaApplication() *ZtnaApplication`

NewZtnaApplication instantiates a new ZtnaApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZtnaApplicationWithDefaults

`func NewZtnaApplicationWithDefaults() *ZtnaApplication`

NewZtnaApplicationWithDefaults instantiates a new ZtnaApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ZtnaApplication) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZtnaApplication) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZtnaApplication) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ZtnaApplication) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *ZtnaApplication) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ZtnaApplication) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ZtnaApplication) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ZtnaApplication) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetCategory

`func (o *ZtnaApplication) GetCategory() ZtnaApplicationCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ZtnaApplication) GetCategoryOk() (*ZtnaApplicationCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ZtnaApplication) SetCategory(v ZtnaApplicationCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ZtnaApplication) HasCategory() bool`

HasCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


