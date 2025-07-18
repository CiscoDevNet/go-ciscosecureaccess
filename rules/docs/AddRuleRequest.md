# AddRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleName** | **string** | The name of the rule. The name must be unique across all rules for the organization&#39;s policy. The name can have no more than 256 characters. | 
**RuleDescription** | Pointer to **string** | The meaningful information about the rule. The description can have no more than 256 characters. | [optional] 
**RuleAction** | [**RuleAction**](RuleAction.md) |  | 
**RulePriority** | Pointer to **int64** | The positive integer that represents the priority of the rule. The priority is unique across all rules on the policy for the organization. | [optional] 
**RuleIsEnabled** | Pointer to **bool** | Specifies whether the rule is enabled. | [optional] 
**RuleConditions** | [**[]RuleConditionsInner**](RuleConditionsInner.md) | The list of conditions that are set on the rule. Updates to \&quot;ReadOnly\&quot; attributes are ignored. | 
**RuleSettings** | [**[]RuleSettingsInner**](RuleSettingsInner.md) | The list of settings on a rule. | 

## Methods

### NewAddRuleRequest

`func NewAddRuleRequest(ruleName string, ruleAction RuleAction, ruleConditions []RuleConditionsInner, ruleSettings []RuleSettingsInner, ) *AddRuleRequest`

NewAddRuleRequest instantiates a new AddRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddRuleRequestWithDefaults

`func NewAddRuleRequestWithDefaults() *AddRuleRequest`

NewAddRuleRequestWithDefaults instantiates a new AddRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleName

`func (o *AddRuleRequest) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *AddRuleRequest) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *AddRuleRequest) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.


### GetRuleDescription

`func (o *AddRuleRequest) GetRuleDescription() string`

GetRuleDescription returns the RuleDescription field if non-nil, zero value otherwise.

### GetRuleDescriptionOk

`func (o *AddRuleRequest) GetRuleDescriptionOk() (*string, bool)`

GetRuleDescriptionOk returns a tuple with the RuleDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleDescription

`func (o *AddRuleRequest) SetRuleDescription(v string)`

SetRuleDescription sets RuleDescription field to given value.

### HasRuleDescription

`func (o *AddRuleRequest) HasRuleDescription() bool`

HasRuleDescription returns a boolean if a field has been set.

### GetRuleAction

`func (o *AddRuleRequest) GetRuleAction() RuleAction`

GetRuleAction returns the RuleAction field if non-nil, zero value otherwise.

### GetRuleActionOk

`func (o *AddRuleRequest) GetRuleActionOk() (*RuleAction, bool)`

GetRuleActionOk returns a tuple with the RuleAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleAction

`func (o *AddRuleRequest) SetRuleAction(v RuleAction)`

SetRuleAction sets RuleAction field to given value.


### GetRulePriority

`func (o *AddRuleRequest) GetRulePriority() int64`

GetRulePriority returns the RulePriority field if non-nil, zero value otherwise.

### GetRulePriorityOk

`func (o *AddRuleRequest) GetRulePriorityOk() (*int64, bool)`

GetRulePriorityOk returns a tuple with the RulePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulePriority

`func (o *AddRuleRequest) SetRulePriority(v int64)`

SetRulePriority sets RulePriority field to given value.

### HasRulePriority

`func (o *AddRuleRequest) HasRulePriority() bool`

HasRulePriority returns a boolean if a field has been set.

### GetRuleIsEnabled

`func (o *AddRuleRequest) GetRuleIsEnabled() bool`

GetRuleIsEnabled returns the RuleIsEnabled field if non-nil, zero value otherwise.

### GetRuleIsEnabledOk

`func (o *AddRuleRequest) GetRuleIsEnabledOk() (*bool, bool)`

GetRuleIsEnabledOk returns a tuple with the RuleIsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIsEnabled

`func (o *AddRuleRequest) SetRuleIsEnabled(v bool)`

SetRuleIsEnabled sets RuleIsEnabled field to given value.

### HasRuleIsEnabled

`func (o *AddRuleRequest) HasRuleIsEnabled() bool`

HasRuleIsEnabled returns a boolean if a field has been set.

### GetRuleConditions

`func (o *AddRuleRequest) GetRuleConditions() []RuleConditionsInner`

GetRuleConditions returns the RuleConditions field if non-nil, zero value otherwise.

### GetRuleConditionsOk

`func (o *AddRuleRequest) GetRuleConditionsOk() (*[]RuleConditionsInner, bool)`

GetRuleConditionsOk returns a tuple with the RuleConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleConditions

`func (o *AddRuleRequest) SetRuleConditions(v []RuleConditionsInner)`

SetRuleConditions sets RuleConditions field to given value.


### GetRuleSettings

`func (o *AddRuleRequest) GetRuleSettings() []RuleSettingsInner`

GetRuleSettings returns the RuleSettings field if non-nil, zero value otherwise.

### GetRuleSettingsOk

`func (o *AddRuleRequest) GetRuleSettingsOk() (*[]RuleSettingsInner, bool)`

GetRuleSettingsOk returns a tuple with the RuleSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleSettings

`func (o *AddRuleRequest) SetRuleSettings(v []RuleSettingsInner)`

SetRuleSettings sets RuleSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


