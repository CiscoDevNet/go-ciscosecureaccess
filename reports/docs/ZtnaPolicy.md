# ZtnaPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ruleset** | [**ZtnaPolicyRuleset**](ZtnaPolicyRuleset.md) |  | 
**Privateapplicationgroup** | Pointer to [**ZtnaPolicyRulesetRule**](ZtnaPolicyRulesetRule.md) |  | [optional] 

## Methods

### NewZtnaPolicy

`func NewZtnaPolicy(ruleset ZtnaPolicyRuleset, ) *ZtnaPolicy`

NewZtnaPolicy instantiates a new ZtnaPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZtnaPolicyWithDefaults

`func NewZtnaPolicyWithDefaults() *ZtnaPolicy`

NewZtnaPolicyWithDefaults instantiates a new ZtnaPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleset

`func (o *ZtnaPolicy) GetRuleset() ZtnaPolicyRuleset`

GetRuleset returns the Ruleset field if non-nil, zero value otherwise.

### GetRulesetOk

`func (o *ZtnaPolicy) GetRulesetOk() (*ZtnaPolicyRuleset, bool)`

GetRulesetOk returns a tuple with the Ruleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleset

`func (o *ZtnaPolicy) SetRuleset(v ZtnaPolicyRuleset)`

SetRuleset sets Ruleset field to given value.


### GetPrivateapplicationgroup

`func (o *ZtnaPolicy) GetPrivateapplicationgroup() ZtnaPolicyRulesetRule`

GetPrivateapplicationgroup returns the Privateapplicationgroup field if non-nil, zero value otherwise.

### GetPrivateapplicationgroupOk

`func (o *ZtnaPolicy) GetPrivateapplicationgroupOk() (*ZtnaPolicyRulesetRule, bool)`

GetPrivateapplicationgroupOk returns a tuple with the Privateapplicationgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateapplicationgroup

`func (o *ZtnaPolicy) SetPrivateapplicationgroup(v ZtnaPolicyRulesetRule)`

SetPrivateapplicationgroup sets Privateapplicationgroup field to given value.

### HasPrivateapplicationgroup

`func (o *ZtnaPolicy) HasPrivateapplicationgroup() bool`

HasPrivateapplicationgroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


