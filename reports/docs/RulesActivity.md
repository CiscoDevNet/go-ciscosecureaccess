# RulesActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rule** | [**Rule**](Rule.md) |  | 
**Timestamp** | **int64** | The date and time that the event happened in milliseconds. | 
**Verdict** | [**Verdict**](Verdict.md) |  | 
**Resource** | Pointer to [**PrivateApplicationGroup**](PrivateApplicationGroup.md) |  | [optional] 

## Methods

### NewRulesActivity

`func NewRulesActivity(rule Rule, timestamp int64, verdict Verdict, ) *RulesActivity`

NewRulesActivity instantiates a new RulesActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulesActivityWithDefaults

`func NewRulesActivityWithDefaults() *RulesActivity`

NewRulesActivityWithDefaults instantiates a new RulesActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRule

`func (o *RulesActivity) GetRule() Rule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *RulesActivity) GetRuleOk() (*Rule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *RulesActivity) SetRule(v Rule)`

SetRule sets Rule field to given value.


### GetTimestamp

`func (o *RulesActivity) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RulesActivity) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RulesActivity) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetVerdict

`func (o *RulesActivity) GetVerdict() Verdict`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *RulesActivity) GetVerdictOk() (*Verdict, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *RulesActivity) SetVerdict(v Verdict)`

SetVerdict sets Verdict field to given value.


### GetResource

`func (o *RulesActivity) GetResource() PrivateApplicationGroup`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *RulesActivity) GetResourceOk() (*PrivateApplicationGroup, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *RulesActivity) SetResource(v PrivateApplicationGroup)`

SetResource sets Resource field to given value.

### HasResource

`func (o *RulesActivity) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


