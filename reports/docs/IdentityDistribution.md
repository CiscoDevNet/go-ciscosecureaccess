# IdentityDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Counts** | [**RequestCounts**](RequestCounts.md) |  | 
**Requests** | **int64** | The requests made by the identity type. | 
**UniqueIdentityCount** | **int64** | The number unique identities associated with the identity type. | 
**Identitytype** | [**IdentityType**](IdentityType.md) |  | 

## Methods

### NewIdentityDistribution

`func NewIdentityDistribution(counts RequestCounts, requests int64, uniqueIdentityCount int64, identitytype IdentityType, ) *IdentityDistribution`

NewIdentityDistribution instantiates a new IdentityDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityDistributionWithDefaults

`func NewIdentityDistributionWithDefaults() *IdentityDistribution`

NewIdentityDistributionWithDefaults instantiates a new IdentityDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounts

`func (o *IdentityDistribution) GetCounts() RequestCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *IdentityDistribution) GetCountsOk() (*RequestCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *IdentityDistribution) SetCounts(v RequestCounts)`

SetCounts sets Counts field to given value.


### GetRequests

`func (o *IdentityDistribution) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *IdentityDistribution) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *IdentityDistribution) SetRequests(v int64)`

SetRequests sets Requests field to given value.


### GetUniqueIdentityCount

`func (o *IdentityDistribution) GetUniqueIdentityCount() int64`

GetUniqueIdentityCount returns the UniqueIdentityCount field if non-nil, zero value otherwise.

### GetUniqueIdentityCountOk

`func (o *IdentityDistribution) GetUniqueIdentityCountOk() (*int64, bool)`

GetUniqueIdentityCountOk returns a tuple with the UniqueIdentityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdentityCount

`func (o *IdentityDistribution) SetUniqueIdentityCount(v int64)`

SetUniqueIdentityCount sets UniqueIdentityCount field to given value.


### GetIdentitytype

`func (o *IdentityDistribution) GetIdentitytype() IdentityType`

GetIdentitytype returns the Identitytype field if non-nil, zero value otherwise.

### GetIdentitytypeOk

`func (o *IdentityDistribution) GetIdentitytypeOk() (*IdentityType, bool)`

GetIdentitytypeOk returns a tuple with the Identitytype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentitytype

`func (o *IdentityDistribution) SetIdentitytype(v IdentityType)`

SetIdentitytype sets Identitytype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


