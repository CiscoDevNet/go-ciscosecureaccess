# TopThreatTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Threattype** | **string** | The type of the threat. | 
**Threatscount** | **int64** | The count of the threats for the threat type. | 
**Count** | **int64** | The number of requests for the threat type. | 

## Methods

### NewTopThreatTypes

`func NewTopThreatTypes(threattype string, threatscount int64, count int64, ) *TopThreatTypes`

NewTopThreatTypes instantiates a new TopThreatTypes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopThreatTypesWithDefaults

`func NewTopThreatTypesWithDefaults() *TopThreatTypes`

NewTopThreatTypesWithDefaults instantiates a new TopThreatTypes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreattype

`func (o *TopThreatTypes) GetThreattype() string`

GetThreattype returns the Threattype field if non-nil, zero value otherwise.

### GetThreattypeOk

`func (o *TopThreatTypes) GetThreattypeOk() (*string, bool)`

GetThreattypeOk returns a tuple with the Threattype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreattype

`func (o *TopThreatTypes) SetThreattype(v string)`

SetThreattype sets Threattype field to given value.


### GetThreatscount

`func (o *TopThreatTypes) GetThreatscount() int64`

GetThreatscount returns the Threatscount field if non-nil, zero value otherwise.

### GetThreatscountOk

`func (o *TopThreatTypes) GetThreatscountOk() (*int64, bool)`

GetThreatscountOk returns a tuple with the Threatscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatscount

`func (o *TopThreatTypes) SetThreatscount(v int64)`

SetThreatscount sets Threatscount field to given value.


### GetCount

`func (o *TopThreatTypes) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TopThreatTypes) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TopThreatTypes) SetCount(v int64)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


