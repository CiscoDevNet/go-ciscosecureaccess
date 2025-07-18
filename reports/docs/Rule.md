# Rule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the rule. | 
**Label** | **string** | The name of the rule. | 
**Privateapplicationgroup** | Pointer to [**PrivateApplicationGroup**](PrivateApplicationGroup.md) |  | [optional] 

## Methods

### NewRule

`func NewRule(id int64, label string, ) *Rule`

NewRule instantiates a new Rule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleWithDefaults

`func NewRuleWithDefaults() *Rule`

NewRuleWithDefaults instantiates a new Rule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Rule) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Rule) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Rule) SetId(v int64)`

SetId sets Id field to given value.


### GetLabel

`func (o *Rule) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Rule) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Rule) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetPrivateapplicationgroup

`func (o *Rule) GetPrivateapplicationgroup() PrivateApplicationGroup`

GetPrivateapplicationgroup returns the Privateapplicationgroup field if non-nil, zero value otherwise.

### GetPrivateapplicationgroupOk

`func (o *Rule) GetPrivateapplicationgroupOk() (*PrivateApplicationGroup, bool)`

GetPrivateapplicationgroupOk returns a tuple with the Privateapplicationgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateapplicationgroup

`func (o *Rule) SetPrivateapplicationgroup(v PrivateApplicationGroup)`

SetPrivateapplicationgroup sets Privateapplicationgroup field to given value.

### HasPrivateapplicationgroup

`func (o *Rule) HasPrivateapplicationgroup() bool`

HasPrivateapplicationgroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


