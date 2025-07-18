# RuleSourceDestFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sources** | Pointer to [**[]map[string]RuleFilterSourcesExampleValue**](map[string]RuleFilterSourcesExampleValue.md) | The list of sources to filter your collection of rules. | [optional] 
**Destinations** | Pointer to [**[]map[string]RuleFilterDestinationsExampleValue**](map[string]RuleFilterDestinationsExampleValue.md) | The list of destinations to filter your collection of rules. | [optional] 

## Methods

### NewRuleSourceDestFilters

`func NewRuleSourceDestFilters() *RuleSourceDestFilters`

NewRuleSourceDestFilters instantiates a new RuleSourceDestFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleSourceDestFiltersWithDefaults

`func NewRuleSourceDestFiltersWithDefaults() *RuleSourceDestFilters`

NewRuleSourceDestFiltersWithDefaults instantiates a new RuleSourceDestFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSources

`func (o *RuleSourceDestFilters) GetSources() []map[string]RuleFilterSourcesExampleValue`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *RuleSourceDestFilters) GetSourcesOk() (*[]map[string]RuleFilterSourcesExampleValue, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *RuleSourceDestFilters) SetSources(v []map[string]RuleFilterSourcesExampleValue)`

SetSources sets Sources field to given value.

### HasSources

`func (o *RuleSourceDestFilters) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetDestinations

`func (o *RuleSourceDestFilters) GetDestinations() []map[string]RuleFilterDestinationsExampleValue`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *RuleSourceDestFilters) GetDestinationsOk() (*[]map[string]RuleFilterDestinationsExampleValue, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *RuleSourceDestFilters) SetDestinations(v []map[string]RuleFilterDestinationsExampleValue)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *RuleSourceDestFilters) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


