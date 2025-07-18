# RuleFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleName** | Pointer to **string** | The name of the rule. The name must be unique across all rules for the organization&#39;s policy. The name can have no more than 256 characters. | [optional] 
**RuleDescription** | Pointer to **string** | The meaningful information about the rule. The description can have no more than 256 characters. | [optional] 
**RuleIsEnabled** | Pointer to **bool** | Specifies whether the rule is enabled. | [optional] 
**RuleIsDefault** | Pointer to **bool** | Specifies whether the rule is the default rule. | [optional] 
**RuleAction** | Pointer to [**RuleAction**](RuleAction.md) |  | [optional] 
**AttributeName** | Pointer to [**AttributeName**](AttributeName.md) |  | [optional] 
**AttributeValue** | Pointer to [**AttributeValue**](AttributeValue.md) |  | [optional] 
**SettingName** | Pointer to [**SettingName**](SettingName.md) |  | [optional] 
**SettingValue** | Pointer to [**SettingValue**](SettingValue.md) |  | [optional] 
**RulePriority** | Pointer to **string** | Filter on the priorities of the rules. Provide a comma-separated string of integers that correspond to the priorities of your rules. | [optional] 

## Methods

### NewRuleFilters

`func NewRuleFilters() *RuleFilters`

NewRuleFilters instantiates a new RuleFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleFiltersWithDefaults

`func NewRuleFiltersWithDefaults() *RuleFilters`

NewRuleFiltersWithDefaults instantiates a new RuleFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleName

`func (o *RuleFilters) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *RuleFilters) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *RuleFilters) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.

### HasRuleName

`func (o *RuleFilters) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### GetRuleDescription

`func (o *RuleFilters) GetRuleDescription() string`

GetRuleDescription returns the RuleDescription field if non-nil, zero value otherwise.

### GetRuleDescriptionOk

`func (o *RuleFilters) GetRuleDescriptionOk() (*string, bool)`

GetRuleDescriptionOk returns a tuple with the RuleDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleDescription

`func (o *RuleFilters) SetRuleDescription(v string)`

SetRuleDescription sets RuleDescription field to given value.

### HasRuleDescription

`func (o *RuleFilters) HasRuleDescription() bool`

HasRuleDescription returns a boolean if a field has been set.

### GetRuleIsEnabled

`func (o *RuleFilters) GetRuleIsEnabled() bool`

GetRuleIsEnabled returns the RuleIsEnabled field if non-nil, zero value otherwise.

### GetRuleIsEnabledOk

`func (o *RuleFilters) GetRuleIsEnabledOk() (*bool, bool)`

GetRuleIsEnabledOk returns a tuple with the RuleIsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIsEnabled

`func (o *RuleFilters) SetRuleIsEnabled(v bool)`

SetRuleIsEnabled sets RuleIsEnabled field to given value.

### HasRuleIsEnabled

`func (o *RuleFilters) HasRuleIsEnabled() bool`

HasRuleIsEnabled returns a boolean if a field has been set.

### GetRuleIsDefault

`func (o *RuleFilters) GetRuleIsDefault() bool`

GetRuleIsDefault returns the RuleIsDefault field if non-nil, zero value otherwise.

### GetRuleIsDefaultOk

`func (o *RuleFilters) GetRuleIsDefaultOk() (*bool, bool)`

GetRuleIsDefaultOk returns a tuple with the RuleIsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIsDefault

`func (o *RuleFilters) SetRuleIsDefault(v bool)`

SetRuleIsDefault sets RuleIsDefault field to given value.

### HasRuleIsDefault

`func (o *RuleFilters) HasRuleIsDefault() bool`

HasRuleIsDefault returns a boolean if a field has been set.

### GetRuleAction

`func (o *RuleFilters) GetRuleAction() RuleAction`

GetRuleAction returns the RuleAction field if non-nil, zero value otherwise.

### GetRuleActionOk

`func (o *RuleFilters) GetRuleActionOk() (*RuleAction, bool)`

GetRuleActionOk returns a tuple with the RuleAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleAction

`func (o *RuleFilters) SetRuleAction(v RuleAction)`

SetRuleAction sets RuleAction field to given value.

### HasRuleAction

`func (o *RuleFilters) HasRuleAction() bool`

HasRuleAction returns a boolean if a field has been set.

### GetAttributeName

`func (o *RuleFilters) GetAttributeName() AttributeName`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *RuleFilters) GetAttributeNameOk() (*AttributeName, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *RuleFilters) SetAttributeName(v AttributeName)`

SetAttributeName sets AttributeName field to given value.

### HasAttributeName

`func (o *RuleFilters) HasAttributeName() bool`

HasAttributeName returns a boolean if a field has been set.

### GetAttributeValue

`func (o *RuleFilters) GetAttributeValue() AttributeValue`

GetAttributeValue returns the AttributeValue field if non-nil, zero value otherwise.

### GetAttributeValueOk

`func (o *RuleFilters) GetAttributeValueOk() (*AttributeValue, bool)`

GetAttributeValueOk returns a tuple with the AttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeValue

`func (o *RuleFilters) SetAttributeValue(v AttributeValue)`

SetAttributeValue sets AttributeValue field to given value.

### HasAttributeValue

`func (o *RuleFilters) HasAttributeValue() bool`

HasAttributeValue returns a boolean if a field has been set.

### GetSettingName

`func (o *RuleFilters) GetSettingName() SettingName`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *RuleFilters) GetSettingNameOk() (*SettingName, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *RuleFilters) SetSettingName(v SettingName)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *RuleFilters) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### GetSettingValue

`func (o *RuleFilters) GetSettingValue() SettingValue`

GetSettingValue returns the SettingValue field if non-nil, zero value otherwise.

### GetSettingValueOk

`func (o *RuleFilters) GetSettingValueOk() (*SettingValue, bool)`

GetSettingValueOk returns a tuple with the SettingValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingValue

`func (o *RuleFilters) SetSettingValue(v SettingValue)`

SetSettingValue sets SettingValue field to given value.

### HasSettingValue

`func (o *RuleFilters) HasSettingValue() bool`

HasSettingValue returns a boolean if a field has been set.

### GetRulePriority

`func (o *RuleFilters) GetRulePriority() string`

GetRulePriority returns the RulePriority field if non-nil, zero value otherwise.

### GetRulePriorityOk

`func (o *RuleFilters) GetRulePriorityOk() (*string, bool)`

GetRulePriorityOk returns a tuple with the RulePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulePriority

`func (o *RuleFilters) SetRulePriority(v string)`

SetRulePriority sets RulePriority field to given value.

### HasRulePriority

`func (o *RuleFilters) HasRulePriority() bool`

HasRulePriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


