# PutRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleName** | **string** | The name of the rule. The name must be unique across all rules for the organization&#39;s policy. The name can have no more than 256 characters. | 
**RuleDescription** | Pointer to **string** | The meaningful information about the rule. The description can have no more than 256 characters. | [optional] 
**RuleAction** | [**RuleAction**](RuleAction.md) |  | 
**RulePriority** | **int64** | The positive integer that represents the priority of the rule. The priority is unique across all rules on the policy for the organization. | 
**RuleIsEnabled** | Pointer to **bool** | Specifies whether the rule is enabled. | [optional] 
**RuleConditions** | [**[]RuleConditionsInner**](RuleConditionsInner.md) | The list of conditions that are set on the rule. Updates to \&quot;ReadOnly\&quot; attributes are ignored. | 
**RuleSettings** | [**[]RuleSettingsInner**](RuleSettingsInner.md) | The list of settings on a rule. | 

## Methods

### NewPutRuleRequest

`func NewPutRuleRequest(ruleName string, ruleAction RuleAction, rulePriority int64, ruleConditions []RuleConditionsInner, ruleSettings []RuleSettingsInner, ) *PutRuleRequest`

NewPutRuleRequest instantiates a new PutRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutRuleRequestWithDefaults

`func NewPutRuleRequestWithDefaults() *PutRuleRequest`

NewPutRuleRequestWithDefaults instantiates a new PutRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleName

`func (o *PutRuleRequest) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *PutRuleRequest) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *PutRuleRequest) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.


### GetRuleDescription

`func (o *PutRuleRequest) GetRuleDescription() string`

GetRuleDescription returns the RuleDescription field if non-nil, zero value otherwise.

### GetRuleDescriptionOk

`func (o *PutRuleRequest) GetRuleDescriptionOk() (*string, bool)`

GetRuleDescriptionOk returns a tuple with the RuleDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleDescription

`func (o *PutRuleRequest) SetRuleDescription(v string)`

SetRuleDescription sets RuleDescription field to given value.

### HasRuleDescription

`func (o *PutRuleRequest) HasRuleDescription() bool`

HasRuleDescription returns a boolean if a field has been set.

### GetRuleAction

`func (o *PutRuleRequest) GetRuleAction() RuleAction`

GetRuleAction returns the RuleAction field if non-nil, zero value otherwise.

### GetRuleActionOk

`func (o *PutRuleRequest) GetRuleActionOk() (*RuleAction, bool)`

GetRuleActionOk returns a tuple with the RuleAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleAction

`func (o *PutRuleRequest) SetRuleAction(v RuleAction)`

SetRuleAction sets RuleAction field to given value.


### GetRulePriority

`func (o *PutRuleRequest) GetRulePriority() int64`

GetRulePriority returns the RulePriority field if non-nil, zero value otherwise.

### GetRulePriorityOk

`func (o *PutRuleRequest) GetRulePriorityOk() (*int64, bool)`

GetRulePriorityOk returns a tuple with the RulePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulePriority

`func (o *PutRuleRequest) SetRulePriority(v int64)`

SetRulePriority sets RulePriority field to given value.


### GetRuleIsEnabled

`func (o *PutRuleRequest) GetRuleIsEnabled() bool`

GetRuleIsEnabled returns the RuleIsEnabled field if non-nil, zero value otherwise.

### GetRuleIsEnabledOk

`func (o *PutRuleRequest) GetRuleIsEnabledOk() (*bool, bool)`

GetRuleIsEnabledOk returns a tuple with the RuleIsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIsEnabled

`func (o *PutRuleRequest) SetRuleIsEnabled(v bool)`

SetRuleIsEnabled sets RuleIsEnabled field to given value.

### HasRuleIsEnabled

`func (o *PutRuleRequest) HasRuleIsEnabled() bool`

HasRuleIsEnabled returns a boolean if a field has been set.

### GetRuleConditions

`func (o *PutRuleRequest) GetRuleConditions() []RuleConditionsInner`

GetRuleConditions returns the RuleConditions field if non-nil, zero value otherwise.

### GetRuleConditionsOk

`func (o *PutRuleRequest) GetRuleConditionsOk() (*[]RuleConditionsInner, bool)`

GetRuleConditionsOk returns a tuple with the RuleConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleConditions

`func (o *PutRuleRequest) SetRuleConditions(v []RuleConditionsInner)`

SetRuleConditions sets RuleConditions field to given value.


### GetRuleSettings

`func (o *PutRuleRequest) GetRuleSettings() []RuleSettingsInner`

GetRuleSettings returns the RuleSettings field if non-nil, zero value otherwise.

### GetRuleSettingsOk

`func (o *PutRuleRequest) GetRuleSettingsOk() (*[]RuleSettingsInner, bool)`

GetRuleSettingsOk returns a tuple with the RuleSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleSettings

`func (o *PutRuleRequest) SetRuleSettings(v []RuleSettingsInner)`

SetRuleSettings sets RuleSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


