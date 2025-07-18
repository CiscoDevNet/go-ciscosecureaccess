# TopThreats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Threat** | **string** | The name of the threat. | 
**Threattype** | **string** | The type of the threat. | 
**Count** | **int64** | The number of requests for the threat name. | 

## Methods

### NewTopThreats

`func NewTopThreats(threat string, threattype string, count int64, ) *TopThreats`

NewTopThreats instantiates a new TopThreats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopThreatsWithDefaults

`func NewTopThreatsWithDefaults() *TopThreats`

NewTopThreatsWithDefaults instantiates a new TopThreats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreat

`func (o *TopThreats) GetThreat() string`

GetThreat returns the Threat field if non-nil, zero value otherwise.

### GetThreatOk

`func (o *TopThreats) GetThreatOk() (*string, bool)`

GetThreatOk returns a tuple with the Threat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreat

`func (o *TopThreats) SetThreat(v string)`

SetThreat sets Threat field to given value.


### GetThreattype

`func (o *TopThreats) GetThreattype() string`

GetThreattype returns the Threattype field if non-nil, zero value otherwise.

### GetThreattypeOk

`func (o *TopThreats) GetThreattypeOk() (*string, bool)`

GetThreattypeOk returns a tuple with the Threattype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreattype

`func (o *TopThreats) SetThreattype(v string)`

SetThreattype sets Threattype field to given value.


### GetCount

`func (o *TopThreats) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TopThreats) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TopThreats) SetCount(v int64)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


