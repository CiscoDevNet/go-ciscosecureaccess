# Policy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timebasedrule** | **bool** | Specify whether the policy triggered a time-of-day rule. | 
**Destinationlistids** | **[]int64** | The list of destination lists that the rules triggered. | 
**Ruleid** | **NullableInt64** | The ID of the rule in the policy. | 
**Rulesetid** | **NullableInt64** | The ID of the ruleset in the policy. | 

## Methods

### NewPolicy

`func NewPolicy(timebasedrule bool, destinationlistids []int64, ruleid NullableInt64, rulesetid NullableInt64, ) *Policy`

NewPolicy instantiates a new Policy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyWithDefaults

`func NewPolicyWithDefaults() *Policy`

NewPolicyWithDefaults instantiates a new Policy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimebasedrule

`func (o *Policy) GetTimebasedrule() bool`

GetTimebasedrule returns the Timebasedrule field if non-nil, zero value otherwise.

### GetTimebasedruleOk

`func (o *Policy) GetTimebasedruleOk() (*bool, bool)`

GetTimebasedruleOk returns a tuple with the Timebasedrule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimebasedrule

`func (o *Policy) SetTimebasedrule(v bool)`

SetTimebasedrule sets Timebasedrule field to given value.


### GetDestinationlistids

`func (o *Policy) GetDestinationlistids() []int64`

GetDestinationlistids returns the Destinationlistids field if non-nil, zero value otherwise.

### GetDestinationlistidsOk

`func (o *Policy) GetDestinationlistidsOk() (*[]int64, bool)`

GetDestinationlistidsOk returns a tuple with the Destinationlistids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationlistids

`func (o *Policy) SetDestinationlistids(v []int64)`

SetDestinationlistids sets Destinationlistids field to given value.


### GetRuleid

`func (o *Policy) GetRuleid() int64`

GetRuleid returns the Ruleid field if non-nil, zero value otherwise.

### GetRuleidOk

`func (o *Policy) GetRuleidOk() (*int64, bool)`

GetRuleidOk returns a tuple with the Ruleid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleid

`func (o *Policy) SetRuleid(v int64)`

SetRuleid sets Ruleid field to given value.


### SetRuleidNil

`func (o *Policy) SetRuleidNil(b bool)`

 SetRuleidNil sets the value for Ruleid to be an explicit nil

### UnsetRuleid
`func (o *Policy) UnsetRuleid()`

UnsetRuleid ensures that no value is present for Ruleid, not even an explicit nil
### GetRulesetid

`func (o *Policy) GetRulesetid() int64`

GetRulesetid returns the Rulesetid field if non-nil, zero value otherwise.

### GetRulesetidOk

`func (o *Policy) GetRulesetidOk() (*int64, bool)`

GetRulesetidOk returns a tuple with the Rulesetid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesetid

`func (o *Policy) SetRulesetid(v int64)`

SetRulesetid sets Rulesetid field to given value.


### SetRulesetidNil

`func (o *Policy) SetRulesetidNil(b bool)`

 SetRulesetidNil sets the value for Rulesetid to be an explicit nil

### UnsetRulesetid
`func (o *Policy) UnsetRulesetid()`

UnsetRulesetid ensures that no value is present for Rulesetid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


