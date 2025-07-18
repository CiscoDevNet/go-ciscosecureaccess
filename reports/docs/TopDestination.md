# TopDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bandwidth** | **NullableInt64** | The amount of bandwidth | 
**Rank** | **int64** | The rank of the result based on the number of requests. | 
**Domain** | **string** | The domain name. | 
**Count** | **int64** | The total number of requests made for this destination. | 
**Counts** | [**RequestCounts**](RequestCounts.md) |  | 
**Categories** | [**[]Category**](Category.md) | The list of categories. | 
**Policycategories** | [**[]Category**](Category.md) | The policy categories that are associated with the destination. | 

## Methods

### NewTopDestination

`func NewTopDestination(bandwidth NullableInt64, rank int64, domain string, count int64, counts RequestCounts, categories []Category, policycategories []Category, ) *TopDestination`

NewTopDestination instantiates a new TopDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopDestinationWithDefaults

`func NewTopDestinationWithDefaults() *TopDestination`

NewTopDestinationWithDefaults instantiates a new TopDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandwidth

`func (o *TopDestination) GetBandwidth() int64`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *TopDestination) GetBandwidthOk() (*int64, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *TopDestination) SetBandwidth(v int64)`

SetBandwidth sets Bandwidth field to given value.


### SetBandwidthNil

`func (o *TopDestination) SetBandwidthNil(b bool)`

 SetBandwidthNil sets the value for Bandwidth to be an explicit nil

### UnsetBandwidth
`func (o *TopDestination) UnsetBandwidth()`

UnsetBandwidth ensures that no value is present for Bandwidth, not even an explicit nil
### GetRank

`func (o *TopDestination) GetRank() int64`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *TopDestination) GetRankOk() (*int64, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *TopDestination) SetRank(v int64)`

SetRank sets Rank field to given value.


### GetDomain

`func (o *TopDestination) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *TopDestination) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *TopDestination) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetCount

`func (o *TopDestination) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TopDestination) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TopDestination) SetCount(v int64)`

SetCount sets Count field to given value.


### GetCounts

`func (o *TopDestination) GetCounts() RequestCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *TopDestination) GetCountsOk() (*RequestCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *TopDestination) SetCounts(v RequestCounts)`

SetCounts sets Counts field to given value.


### GetCategories

`func (o *TopDestination) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *TopDestination) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *TopDestination) SetCategories(v []Category)`

SetCategories sets Categories field to given value.


### GetPolicycategories

`func (o *TopDestination) GetPolicycategories() []Category`

GetPolicycategories returns the Policycategories field if non-nil, zero value otherwise.

### GetPolicycategoriesOk

`func (o *TopDestination) GetPolicycategoriesOk() (*[]Category, bool)`

GetPolicycategoriesOk returns a tuple with the Policycategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicycategories

`func (o *TopDestination) SetPolicycategories(v []Category)`

SetPolicycategories sets Policycategories field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


