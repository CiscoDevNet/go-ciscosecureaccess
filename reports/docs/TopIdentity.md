# TopIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | **int64** | The total number of requests made by this identity. | 
**Bandwidth** | **NullableInt64** | The amount of bandwidth | 
**Identity** | [**Identity**](Identity.md) |  | 
**Counts** | [**RequestAndConnectionCounts**](RequestAndConnectionCounts.md) |  | 
**Rank** | **int64** | The rank of the result based on the number of requests. | 

## Methods

### NewTopIdentity

`func NewTopIdentity(requests int64, bandwidth NullableInt64, identity Identity, counts RequestAndConnectionCounts, rank int64, ) *TopIdentity`

NewTopIdentity instantiates a new TopIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopIdentityWithDefaults

`func NewTopIdentityWithDefaults() *TopIdentity`

NewTopIdentityWithDefaults instantiates a new TopIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *TopIdentity) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *TopIdentity) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *TopIdentity) SetRequests(v int64)`

SetRequests sets Requests field to given value.


### GetBandwidth

`func (o *TopIdentity) GetBandwidth() int64`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *TopIdentity) GetBandwidthOk() (*int64, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *TopIdentity) SetBandwidth(v int64)`

SetBandwidth sets Bandwidth field to given value.


### SetBandwidthNil

`func (o *TopIdentity) SetBandwidthNil(b bool)`

 SetBandwidthNil sets the value for Bandwidth to be an explicit nil

### UnsetBandwidth
`func (o *TopIdentity) UnsetBandwidth()`

UnsetBandwidth ensures that no value is present for Bandwidth, not even an explicit nil
### GetIdentity

`func (o *TopIdentity) GetIdentity() Identity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *TopIdentity) GetIdentityOk() (*Identity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *TopIdentity) SetIdentity(v Identity)`

SetIdentity sets Identity field to given value.


### GetCounts

`func (o *TopIdentity) GetCounts() RequestAndConnectionCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *TopIdentity) GetCountsOk() (*RequestAndConnectionCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *TopIdentity) SetCounts(v RequestAndConnectionCounts)`

SetCounts sets Counts field to given value.


### GetRank

`func (o *TopIdentity) GetRank() int64`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *TopIdentity) GetRankOk() (*int64, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *TopIdentity) SetRank(v int64)`

SetRank sets Rank field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


