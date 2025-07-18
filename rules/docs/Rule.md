# Rule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | Pointer to **int64** | The ID of the organization. | [optional] [readonly] 
**RuleId** | Pointer to **int64** | The ID of the rule. | [optional] [readonly] 
**RuleName** | Pointer to **string** | The name of the rule. The name must be unique across all rules for the organization&#39;s policy. The name can have no more than 256 characters. | [optional] 
**RuleDescription** | Pointer to **string** | The meaningful information about the rule. The description can have no more than 256 characters. | [optional] 
**RuleAction** | Pointer to [**RuleAction**](RuleAction.md) |  | [optional] 
**RulePriority** | Pointer to **int64** | The positive integer that represents the priority of the rule. The priority is unique across all rules on the policy for the organization. | [optional] 
**RuleIsDefault** | Pointer to **bool** | Specifies whether the rule is the default rule. | [optional] 
**RuleIsEnabled** | Pointer to **bool** | Specifies whether the rule is enabled. | [optional] 
**RuleConditions** | Pointer to [**[]RuleConditionsInner**](RuleConditionsInner.md) | The list of conditions that are set on the rule. Updates to \&quot;ReadOnly\&quot; attributes are ignored. | [optional] 
**RuleSettings** | Pointer to [**[]SettingResponseInner**](SettingResponseInner.md) | The properties of the policy settings. | [optional] 
**ModifiedBy** | Pointer to **string** | A string that includes the organization ID and user ID for the account that modified the access rule. | [optional] 
**ModifiedAt** | Pointer to **string** | The date and time that the rule was modified. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | The date and time that the rule was created. | [optional] [readonly] 

## Methods

### NewRule

`func NewRule() *Rule`

NewRule instantiates a new Rule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleWithDefaults

`func NewRuleWithDefaults() *Rule`

NewRuleWithDefaults instantiates a new Rule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *Rule) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Rule) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Rule) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Rule) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetRuleId

`func (o *Rule) GetRuleId() int64`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *Rule) GetRuleIdOk() (*int64, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *Rule) SetRuleId(v int64)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *Rule) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetRuleName

`func (o *Rule) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *Rule) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *Rule) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.

### HasRuleName

`func (o *Rule) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### GetRuleDescription

`func (o *Rule) GetRuleDescription() string`

GetRuleDescription returns the RuleDescription field if non-nil, zero value otherwise.

### GetRuleDescriptionOk

`func (o *Rule) GetRuleDescriptionOk() (*string, bool)`

GetRuleDescriptionOk returns a tuple with the RuleDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleDescription

`func (o *Rule) SetRuleDescription(v string)`

SetRuleDescription sets RuleDescription field to given value.

### HasRuleDescription

`func (o *Rule) HasRuleDescription() bool`

HasRuleDescription returns a boolean if a field has been set.

### GetRuleAction

`func (o *Rule) GetRuleAction() RuleAction`

GetRuleAction returns the RuleAction field if non-nil, zero value otherwise.

### GetRuleActionOk

`func (o *Rule) GetRuleActionOk() (*RuleAction, bool)`

GetRuleActionOk returns a tuple with the RuleAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleAction

`func (o *Rule) SetRuleAction(v RuleAction)`

SetRuleAction sets RuleAction field to given value.

### HasRuleAction

`func (o *Rule) HasRuleAction() bool`

HasRuleAction returns a boolean if a field has been set.

### GetRulePriority

`func (o *Rule) GetRulePriority() int64`

GetRulePriority returns the RulePriority field if non-nil, zero value otherwise.

### GetRulePriorityOk

`func (o *Rule) GetRulePriorityOk() (*int64, bool)`

GetRulePriorityOk returns a tuple with the RulePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulePriority

`func (o *Rule) SetRulePriority(v int64)`

SetRulePriority sets RulePriority field to given value.

### HasRulePriority

`func (o *Rule) HasRulePriority() bool`

HasRulePriority returns a boolean if a field has been set.

### GetRuleIsDefault

`func (o *Rule) GetRuleIsDefault() bool`

GetRuleIsDefault returns the RuleIsDefault field if non-nil, zero value otherwise.

### GetRuleIsDefaultOk

`func (o *Rule) GetRuleIsDefaultOk() (*bool, bool)`

GetRuleIsDefaultOk returns a tuple with the RuleIsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIsDefault

`func (o *Rule) SetRuleIsDefault(v bool)`

SetRuleIsDefault sets RuleIsDefault field to given value.

### HasRuleIsDefault

`func (o *Rule) HasRuleIsDefault() bool`

HasRuleIsDefault returns a boolean if a field has been set.

### GetRuleIsEnabled

`func (o *Rule) GetRuleIsEnabled() bool`

GetRuleIsEnabled returns the RuleIsEnabled field if non-nil, zero value otherwise.

### GetRuleIsEnabledOk

`func (o *Rule) GetRuleIsEnabledOk() (*bool, bool)`

GetRuleIsEnabledOk returns a tuple with the RuleIsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIsEnabled

`func (o *Rule) SetRuleIsEnabled(v bool)`

SetRuleIsEnabled sets RuleIsEnabled field to given value.

### HasRuleIsEnabled

`func (o *Rule) HasRuleIsEnabled() bool`

HasRuleIsEnabled returns a boolean if a field has been set.

### GetRuleConditions

`func (o *Rule) GetRuleConditions() []RuleConditionsInner`

GetRuleConditions returns the RuleConditions field if non-nil, zero value otherwise.

### GetRuleConditionsOk

`func (o *Rule) GetRuleConditionsOk() (*[]RuleConditionsInner, bool)`

GetRuleConditionsOk returns a tuple with the RuleConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleConditions

`func (o *Rule) SetRuleConditions(v []RuleConditionsInner)`

SetRuleConditions sets RuleConditions field to given value.

### HasRuleConditions

`func (o *Rule) HasRuleConditions() bool`

HasRuleConditions returns a boolean if a field has been set.

### GetRuleSettings

`func (o *Rule) GetRuleSettings() []SettingResponseInner`

GetRuleSettings returns the RuleSettings field if non-nil, zero value otherwise.

### GetRuleSettingsOk

`func (o *Rule) GetRuleSettingsOk() (*[]SettingResponseInner, bool)`

GetRuleSettingsOk returns a tuple with the RuleSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleSettings

`func (o *Rule) SetRuleSettings(v []SettingResponseInner)`

SetRuleSettings sets RuleSettings field to given value.

### HasRuleSettings

`func (o *Rule) HasRuleSettings() bool`

HasRuleSettings returns a boolean if a field has been set.

### GetModifiedBy

`func (o *Rule) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Rule) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Rule) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *Rule) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Rule) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Rule) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Rule) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Rule) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Rule) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Rule) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Rule) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Rule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


