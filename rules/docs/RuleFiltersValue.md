# RuleFiltersValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sources** | Pointer to [**[]map[string]RuleFilterSourcesExampleValue**](map[string]RuleFilterSourcesExampleValue.md) | The list of sources to filter your collection of rules. | [optional] 
**Destinations** | Pointer to [**[]map[string]RuleFilterDestinationsExampleValue**](map[string]RuleFilterDestinationsExampleValue.md) | The list of destinations to filter your collection of rules. | [optional] 

## Methods

### NewRuleFiltersValue

`func NewRuleFiltersValue() *RuleFiltersValue`

NewRuleFiltersValue instantiates a new RuleFiltersValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleFiltersValueWithDefaults

`func NewRuleFiltersValueWithDefaults() *RuleFiltersValue`

NewRuleFiltersValueWithDefaults instantiates a new RuleFiltersValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSources

`func (o *RuleFiltersValue) GetSources() []map[string]RuleFilterSourcesExampleValue`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *RuleFiltersValue) GetSourcesOk() (*[]map[string]RuleFilterSourcesExampleValue, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *RuleFiltersValue) SetSources(v []map[string]RuleFilterSourcesExampleValue)`

SetSources sets Sources field to given value.

### HasSources

`func (o *RuleFiltersValue) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetDestinations

`func (o *RuleFiltersValue) GetDestinations() []map[string]RuleFilterDestinationsExampleValue`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *RuleFiltersValue) GetDestinationsOk() (*[]map[string]RuleFilterDestinationsExampleValue, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *RuleFiltersValue) SetDestinations(v []map[string]RuleFilterDestinationsExampleValue)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *RuleFiltersValue) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


