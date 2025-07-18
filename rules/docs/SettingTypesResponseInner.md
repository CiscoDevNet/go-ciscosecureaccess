# SettingTypesResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceScopes** | Pointer to **[]string** | The list of services that are related to the policy setting type. | [optional] 
**SettingForRulesets** | Pointer to **bool** | Specify whether the setting is for the rulesets. | [optional] 
**SettingId** | Pointer to **int64** | The ID of the setting. | [optional] 
**SettingName** | Pointer to [**SettingName**](SettingName.md) |  | [optional] 
**TypeDescription** | Pointer to **string** | The type of the setting description. | [optional] 
**TypeName** | Pointer to **string** | The type of the setting name. | [optional] 
**TypeDefaultValue** | Pointer to **bool** | Specifiy whether the setting has the default value. | [optional] 
**SettingDescription** | Pointer to **string** | The description of the access rule setting. | [optional] 
**TypeValidationRegex** | Pointer to **string** | The regex pattern for validating the type. | [optional] 
**SettingForRules** | Pointer to **bool** | Specify whether the setting is for the rules. | [optional] 
**SettingForOrganizations** | Pointer to **bool** | Specify whether the setting is for the organization. | [optional] 

## Methods

### NewSettingTypesResponseInner

`func NewSettingTypesResponseInner() *SettingTypesResponseInner`

NewSettingTypesResponseInner instantiates a new SettingTypesResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingTypesResponseInnerWithDefaults

`func NewSettingTypesResponseInnerWithDefaults() *SettingTypesResponseInner`

NewSettingTypesResponseInnerWithDefaults instantiates a new SettingTypesResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceScopes

`func (o *SettingTypesResponseInner) GetServiceScopes() []string`

GetServiceScopes returns the ServiceScopes field if non-nil, zero value otherwise.

### GetServiceScopesOk

`func (o *SettingTypesResponseInner) GetServiceScopesOk() (*[]string, bool)`

GetServiceScopesOk returns a tuple with the ServiceScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceScopes

`func (o *SettingTypesResponseInner) SetServiceScopes(v []string)`

SetServiceScopes sets ServiceScopes field to given value.

### HasServiceScopes

`func (o *SettingTypesResponseInner) HasServiceScopes() bool`

HasServiceScopes returns a boolean if a field has been set.

### GetSettingForRulesets

`func (o *SettingTypesResponseInner) GetSettingForRulesets() bool`

GetSettingForRulesets returns the SettingForRulesets field if non-nil, zero value otherwise.

### GetSettingForRulesetsOk

`func (o *SettingTypesResponseInner) GetSettingForRulesetsOk() (*bool, bool)`

GetSettingForRulesetsOk returns a tuple with the SettingForRulesets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingForRulesets

`func (o *SettingTypesResponseInner) SetSettingForRulesets(v bool)`

SetSettingForRulesets sets SettingForRulesets field to given value.

### HasSettingForRulesets

`func (o *SettingTypesResponseInner) HasSettingForRulesets() bool`

HasSettingForRulesets returns a boolean if a field has been set.

### GetSettingId

`func (o *SettingTypesResponseInner) GetSettingId() int64`

GetSettingId returns the SettingId field if non-nil, zero value otherwise.

### GetSettingIdOk

`func (o *SettingTypesResponseInner) GetSettingIdOk() (*int64, bool)`

GetSettingIdOk returns a tuple with the SettingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingId

`func (o *SettingTypesResponseInner) SetSettingId(v int64)`

SetSettingId sets SettingId field to given value.

### HasSettingId

`func (o *SettingTypesResponseInner) HasSettingId() bool`

HasSettingId returns a boolean if a field has been set.

### GetSettingName

`func (o *SettingTypesResponseInner) GetSettingName() SettingName`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *SettingTypesResponseInner) GetSettingNameOk() (*SettingName, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *SettingTypesResponseInner) SetSettingName(v SettingName)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *SettingTypesResponseInner) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### GetTypeDescription

`func (o *SettingTypesResponseInner) GetTypeDescription() string`

GetTypeDescription returns the TypeDescription field if non-nil, zero value otherwise.

### GetTypeDescriptionOk

`func (o *SettingTypesResponseInner) GetTypeDescriptionOk() (*string, bool)`

GetTypeDescriptionOk returns a tuple with the TypeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDescription

`func (o *SettingTypesResponseInner) SetTypeDescription(v string)`

SetTypeDescription sets TypeDescription field to given value.

### HasTypeDescription

`func (o *SettingTypesResponseInner) HasTypeDescription() bool`

HasTypeDescription returns a boolean if a field has been set.

### GetTypeName

`func (o *SettingTypesResponseInner) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *SettingTypesResponseInner) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *SettingTypesResponseInner) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *SettingTypesResponseInner) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### GetTypeDefaultValue

`func (o *SettingTypesResponseInner) GetTypeDefaultValue() bool`

GetTypeDefaultValue returns the TypeDefaultValue field if non-nil, zero value otherwise.

### GetTypeDefaultValueOk

`func (o *SettingTypesResponseInner) GetTypeDefaultValueOk() (*bool, bool)`

GetTypeDefaultValueOk returns a tuple with the TypeDefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDefaultValue

`func (o *SettingTypesResponseInner) SetTypeDefaultValue(v bool)`

SetTypeDefaultValue sets TypeDefaultValue field to given value.

### HasTypeDefaultValue

`func (o *SettingTypesResponseInner) HasTypeDefaultValue() bool`

HasTypeDefaultValue returns a boolean if a field has been set.

### GetSettingDescription

`func (o *SettingTypesResponseInner) GetSettingDescription() string`

GetSettingDescription returns the SettingDescription field if non-nil, zero value otherwise.

### GetSettingDescriptionOk

`func (o *SettingTypesResponseInner) GetSettingDescriptionOk() (*string, bool)`

GetSettingDescriptionOk returns a tuple with the SettingDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingDescription

`func (o *SettingTypesResponseInner) SetSettingDescription(v string)`

SetSettingDescription sets SettingDescription field to given value.

### HasSettingDescription

`func (o *SettingTypesResponseInner) HasSettingDescription() bool`

HasSettingDescription returns a boolean if a field has been set.

### GetTypeValidationRegex

`func (o *SettingTypesResponseInner) GetTypeValidationRegex() string`

GetTypeValidationRegex returns the TypeValidationRegex field if non-nil, zero value otherwise.

### GetTypeValidationRegexOk

`func (o *SettingTypesResponseInner) GetTypeValidationRegexOk() (*string, bool)`

GetTypeValidationRegexOk returns a tuple with the TypeValidationRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeValidationRegex

`func (o *SettingTypesResponseInner) SetTypeValidationRegex(v string)`

SetTypeValidationRegex sets TypeValidationRegex field to given value.

### HasTypeValidationRegex

`func (o *SettingTypesResponseInner) HasTypeValidationRegex() bool`

HasTypeValidationRegex returns a boolean if a field has been set.

### GetSettingForRules

`func (o *SettingTypesResponseInner) GetSettingForRules() bool`

GetSettingForRules returns the SettingForRules field if non-nil, zero value otherwise.

### GetSettingForRulesOk

`func (o *SettingTypesResponseInner) GetSettingForRulesOk() (*bool, bool)`

GetSettingForRulesOk returns a tuple with the SettingForRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingForRules

`func (o *SettingTypesResponseInner) SetSettingForRules(v bool)`

SetSettingForRules sets SettingForRules field to given value.

### HasSettingForRules

`func (o *SettingTypesResponseInner) HasSettingForRules() bool`

HasSettingForRules returns a boolean if a field has been set.

### GetSettingForOrganizations

`func (o *SettingTypesResponseInner) GetSettingForOrganizations() bool`

GetSettingForOrganizations returns the SettingForOrganizations field if non-nil, zero value otherwise.

### GetSettingForOrganizationsOk

`func (o *SettingTypesResponseInner) GetSettingForOrganizationsOk() (*bool, bool)`

GetSettingForOrganizationsOk returns a tuple with the SettingForOrganizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingForOrganizations

`func (o *SettingTypesResponseInner) SetSettingForOrganizations(v bool)`

SetSettingForOrganizations sets SettingForOrganizations field to given value.

### HasSettingForOrganizations

`func (o *SettingTypesResponseInner) HasSettingForOrganizations() bool`

HasSettingForOrganizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


