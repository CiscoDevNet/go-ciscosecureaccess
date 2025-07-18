# ZtnaPolicyRuleset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Rule** | Pointer to [**ZtnaPolicyRulesetRule**](ZtnaPolicyRulesetRule.md) |  | [optional] 

## Methods

### NewZtnaPolicyRuleset

`func NewZtnaPolicyRuleset() *ZtnaPolicyRuleset`

NewZtnaPolicyRuleset instantiates a new ZtnaPolicyRuleset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZtnaPolicyRulesetWithDefaults

`func NewZtnaPolicyRulesetWithDefaults() *ZtnaPolicyRuleset`

NewZtnaPolicyRulesetWithDefaults instantiates a new ZtnaPolicyRuleset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ZtnaPolicyRuleset) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZtnaPolicyRuleset) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZtnaPolicyRuleset) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ZtnaPolicyRuleset) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *ZtnaPolicyRuleset) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ZtnaPolicyRuleset) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ZtnaPolicyRuleset) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ZtnaPolicyRuleset) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRule

`func (o *ZtnaPolicyRuleset) GetRule() ZtnaPolicyRulesetRule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *ZtnaPolicyRuleset) GetRuleOk() (*ZtnaPolicyRulesetRule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *ZtnaPolicyRuleset) SetRule(v ZtnaPolicyRulesetRule)`

SetRule sets Rule field to given value.

### HasRule

`func (o *ZtnaPolicyRuleset) HasRule() bool`

HasRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


