# RuleHitcountSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ruleid** | Pointer to **int64** | The policy rule ID. | [optional] 
**Hitcount** | Pointer to **int64** | The total number of times that the policy activated the rule ID. | [optional] 
**Lasteventat** | Pointer to **int64** | The last time in milliseconds since the Unix Epoch where the policy activated the rule ID. | [optional] 

## Methods

### NewRuleHitcountSummary

`func NewRuleHitcountSummary() *RuleHitcountSummary`

NewRuleHitcountSummary instantiates a new RuleHitcountSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleHitcountSummaryWithDefaults

`func NewRuleHitcountSummaryWithDefaults() *RuleHitcountSummary`

NewRuleHitcountSummaryWithDefaults instantiates a new RuleHitcountSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleid

`func (o *RuleHitcountSummary) GetRuleid() int64`

GetRuleid returns the Ruleid field if non-nil, zero value otherwise.

### GetRuleidOk

`func (o *RuleHitcountSummary) GetRuleidOk() (*int64, bool)`

GetRuleidOk returns a tuple with the Ruleid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleid

`func (o *RuleHitcountSummary) SetRuleid(v int64)`

SetRuleid sets Ruleid field to given value.

### HasRuleid

`func (o *RuleHitcountSummary) HasRuleid() bool`

HasRuleid returns a boolean if a field has been set.

### GetHitcount

`func (o *RuleHitcountSummary) GetHitcount() int64`

GetHitcount returns the Hitcount field if non-nil, zero value otherwise.

### GetHitcountOk

`func (o *RuleHitcountSummary) GetHitcountOk() (*int64, bool)`

GetHitcountOk returns a tuple with the Hitcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitcount

`func (o *RuleHitcountSummary) SetHitcount(v int64)`

SetHitcount sets Hitcount field to given value.

### HasHitcount

`func (o *RuleHitcountSummary) HasHitcount() bool`

HasHitcount returns a boolean if a field has been set.

### GetLasteventat

`func (o *RuleHitcountSummary) GetLasteventat() int64`

GetLasteventat returns the Lasteventat field if non-nil, zero value otherwise.

### GetLasteventatOk

`func (o *RuleHitcountSummary) GetLasteventatOk() (*int64, bool)`

GetLasteventatOk returns a tuple with the Lasteventat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLasteventat

`func (o *RuleHitcountSummary) SetLasteventat(v int64)`

SetLasteventat sets Lasteventat field to given value.

### HasLasteventat

`func (o *RuleHitcountSummary) HasLasteventat() bool`

HasLasteventat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


