# Ruleset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the ruleset. | [optional] 
**Label** | Pointer to **string** | The descriptive label for the ruleset. | [optional] 
**Rule** | Pointer to [**Rule**](Rule.md) |  | [optional] 

## Methods

### NewRuleset

`func NewRuleset() *Ruleset`

NewRuleset instantiates a new Ruleset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulesetWithDefaults

`func NewRulesetWithDefaults() *Ruleset`

NewRulesetWithDefaults instantiates a new Ruleset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Ruleset) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ruleset) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Ruleset) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Ruleset) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *Ruleset) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Ruleset) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Ruleset) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Ruleset) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRule

`func (o *Ruleset) GetRule() Rule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *Ruleset) GetRuleOk() (*Rule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *Ruleset) SetRule(v Rule)`

SetRule sets Rule field to given value.

### HasRule

`func (o *Ruleset) HasRule() bool`

HasRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


