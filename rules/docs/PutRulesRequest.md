# PutRulesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleIds** | **[]int64** | The list of rule IDs for the rules. Set the &#x60;rulesIsEnabled&#x60; property on these rules. | 
**Properties** | [**[]PutRulesRequestPropertiesInner**](PutRulesRequestPropertiesInner.md) | The list of the &#x60;ruleIsEnabled&#x60; setting for the rules. | 

## Methods

### NewPutRulesRequest

`func NewPutRulesRequest(ruleIds []int64, properties []PutRulesRequestPropertiesInner, ) *PutRulesRequest`

NewPutRulesRequest instantiates a new PutRulesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutRulesRequestWithDefaults

`func NewPutRulesRequestWithDefaults() *PutRulesRequest`

NewPutRulesRequestWithDefaults instantiates a new PutRulesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleIds

`func (o *PutRulesRequest) GetRuleIds() []int64`

GetRuleIds returns the RuleIds field if non-nil, zero value otherwise.

### GetRuleIdsOk

`func (o *PutRulesRequest) GetRuleIdsOk() (*[]int64, bool)`

GetRuleIdsOk returns a tuple with the RuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIds

`func (o *PutRulesRequest) SetRuleIds(v []int64)`

SetRuleIds sets RuleIds field to given value.


### GetProperties

`func (o *PutRulesRequest) GetProperties() []PutRulesRequestPropertiesInner`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PutRulesRequest) GetPropertiesOk() (*[]PutRulesRequestPropertiesInner, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PutRulesRequest) SetProperties(v []PutRulesRequestPropertiesInner)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


