# Rules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | The number of rules in the policy. | [optional] 
**Result** | Pointer to [**[]RulesResponseInner**](RulesResponseInner.md) | The list of rules in the policy. | [optional] 

## Methods

### NewRules

`func NewRules() *Rules`

NewRules instantiates a new Rules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulesWithDefaults

`func NewRulesWithDefaults() *Rules`

NewRulesWithDefaults instantiates a new Rules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *Rules) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Rules) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Rules) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *Rules) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResult

`func (o *Rules) GetResult() []RulesResponseInner`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Rules) GetResultOk() (*[]RulesResponseInner, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Rules) SetResult(v []RulesResponseInner)`

SetResult sets Result field to given value.

### HasResult

`func (o *Rules) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


