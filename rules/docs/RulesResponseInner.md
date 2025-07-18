# RulesResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | Pointer to **int64** | The ID of the organization. | [optional] [readonly] 
**RuleId** | **int64** | The ID of the rule. | [readonly] 
**RuleName** | Pointer to **string** | The name of the rule. The name must be unique across all rules for the organization&#39;s policy. The name can have no more than 256 characters. | [optional] 
**RuleDescription** | Pointer to **string** | The meaningful information about the rule. The description can have no more than 256 characters. | [optional] 
**RuleAction** | Pointer to [**RuleAction**](RuleAction.md) |  | [optional] 
**RulePriority** | Pointer to **int64** | The positive integer that represents the priority of the rule. The priority is unique across all rules on the policy for the organization. | [optional] 
**RuleIsDefault** | Pointer to **bool** | Specifies whether the rule is the default rule. | [optional] 
**RuleIsEnabled** | Pointer to **bool** | Specifies whether the rule is enabled. | [optional] 
**ModifiedBy** | Pointer to **string** | A string that includes the organization ID and user ID for the account that modified the access rule. | [optional] 
**ModifiedAt** | Pointer to **string** | The date and time that the rule was modified. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | The date and time that the rule was created. | [optional] [readonly] 

## Methods

### NewRulesResponseInner

`func NewRulesResponseInner(ruleId int64, ) *RulesResponseInner`

NewRulesResponseInner instantiates a new RulesResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulesResponseInnerWithDefaults

`func NewRulesResponseInnerWithDefaults() *RulesResponseInner`

NewRulesResponseInnerWithDefaults instantiates a new RulesResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *RulesResponseInner) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *RulesResponseInner) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *RulesResponseInner) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *RulesResponseInner) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetRuleId

`func (o *RulesResponseInner) GetRuleId() int64`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *RulesResponseInner) GetRuleIdOk() (*int64, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *RulesResponseInner) SetRuleId(v int64)`

SetRuleId sets RuleId field to given value.


### GetRuleName

`func (o *RulesResponseInner) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *RulesResponseInner) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *RulesResponseInner) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.

### HasRuleName

`func (o *RulesResponseInner) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### GetRuleDescription

`func (o *RulesResponseInner) GetRuleDescription() string`

GetRuleDescription returns the RuleDescription field if non-nil, zero value otherwise.

### GetRuleDescriptionOk

`func (o *RulesResponseInner) GetRuleDescriptionOk() (*string, bool)`

GetRuleDescriptionOk returns a tuple with the RuleDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleDescription

`func (o *RulesResponseInner) SetRuleDescription(v string)`

SetRuleDescription sets RuleDescription field to given value.

### HasRuleDescription

`func (o *RulesResponseInner) HasRuleDescription() bool`

HasRuleDescription returns a boolean if a field has been set.

### GetRuleAction

`func (o *RulesResponseInner) GetRuleAction() RuleAction`

GetRuleAction returns the RuleAction field if non-nil, zero value otherwise.

### GetRuleActionOk

`func (o *RulesResponseInner) GetRuleActionOk() (*RuleAction, bool)`

GetRuleActionOk returns a tuple with the RuleAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleAction

`func (o *RulesResponseInner) SetRuleAction(v RuleAction)`

SetRuleAction sets RuleAction field to given value.

### HasRuleAction

`func (o *RulesResponseInner) HasRuleAction() bool`

HasRuleAction returns a boolean if a field has been set.

### GetRulePriority

`func (o *RulesResponseInner) GetRulePriority() int64`

GetRulePriority returns the RulePriority field if non-nil, zero value otherwise.

### GetRulePriorityOk

`func (o *RulesResponseInner) GetRulePriorityOk() (*int64, bool)`

GetRulePriorityOk returns a tuple with the RulePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulePriority

`func (o *RulesResponseInner) SetRulePriority(v int64)`

SetRulePriority sets RulePriority field to given value.

### HasRulePriority

`func (o *RulesResponseInner) HasRulePriority() bool`

HasRulePriority returns a boolean if a field has been set.

### GetRuleIsDefault

`func (o *RulesResponseInner) GetRuleIsDefault() bool`

GetRuleIsDefault returns the RuleIsDefault field if non-nil, zero value otherwise.

### GetRuleIsDefaultOk

`func (o *RulesResponseInner) GetRuleIsDefaultOk() (*bool, bool)`

GetRuleIsDefaultOk returns a tuple with the RuleIsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIsDefault

`func (o *RulesResponseInner) SetRuleIsDefault(v bool)`

SetRuleIsDefault sets RuleIsDefault field to given value.

### HasRuleIsDefault

`func (o *RulesResponseInner) HasRuleIsDefault() bool`

HasRuleIsDefault returns a boolean if a field has been set.

### GetRuleIsEnabled

`func (o *RulesResponseInner) GetRuleIsEnabled() bool`

GetRuleIsEnabled returns the RuleIsEnabled field if non-nil, zero value otherwise.

### GetRuleIsEnabledOk

`func (o *RulesResponseInner) GetRuleIsEnabledOk() (*bool, bool)`

GetRuleIsEnabledOk returns a tuple with the RuleIsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIsEnabled

`func (o *RulesResponseInner) SetRuleIsEnabled(v bool)`

SetRuleIsEnabled sets RuleIsEnabled field to given value.

### HasRuleIsEnabled

`func (o *RulesResponseInner) HasRuleIsEnabled() bool`

HasRuleIsEnabled returns a boolean if a field has been set.

### GetModifiedBy

`func (o *RulesResponseInner) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *RulesResponseInner) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *RulesResponseInner) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *RulesResponseInner) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetModifiedAt

`func (o *RulesResponseInner) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *RulesResponseInner) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *RulesResponseInner) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *RulesResponseInner) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RulesResponseInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RulesResponseInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RulesResponseInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RulesResponseInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


