# PrivateResourceDetailedStatsTimerangeTotalHitsCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** | The total hit counts for the private resource. | [optional] 
**Success** | Pointer to **int64** | The number of hit counts for the private resource where the clients was allowed to access. | [optional] 
**Blocked** | Pointer to **int64** | The counts of the blocked access to the private resource. | [optional] 

## Methods

### NewPrivateResourceDetailedStatsTimerangeTotalHitsCount

`func NewPrivateResourceDetailedStatsTimerangeTotalHitsCount() *PrivateResourceDetailedStatsTimerangeTotalHitsCount`

NewPrivateResourceDetailedStatsTimerangeTotalHitsCount instantiates a new PrivateResourceDetailedStatsTimerangeTotalHitsCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateResourceDetailedStatsTimerangeTotalHitsCountWithDefaults

`func NewPrivateResourceDetailedStatsTimerangeTotalHitsCountWithDefaults() *PrivateResourceDetailedStatsTimerangeTotalHitsCount`

NewPrivateResourceDetailedStatsTimerangeTotalHitsCountWithDefaults instantiates a new PrivateResourceDetailedStatsTimerangeTotalHitsCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *PrivateResourceDetailedStatsTimerangeTotalHitsCount) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PrivateResourceDetailedStatsTimerangeTotalHitsCount) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PrivateResourceDetailedStatsTimerangeTotalHitsCount) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PrivateResourceDetailedStatsTimerangeTotalHitsCount) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetSuccess

`func (o *PrivateResourceDetailedStatsTimerangeTotalHitsCount) GetSuccess() int64`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *PrivateResourceDetailedStatsTimerangeTotalHitsCount) GetSuccessOk() (*int64, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *PrivateResourceDetailedStatsTimerangeTotalHitsCount) SetSuccess(v int64)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *PrivateResourceDetailedStatsTimerangeTotalHitsCount) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetBlocked

`func (o *PrivateResourceDetailedStatsTimerangeTotalHitsCount) GetBlocked() int64`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *PrivateResourceDetailedStatsTimerangeTotalHitsCount) GetBlockedOk() (*int64, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *PrivateResourceDetailedStatsTimerangeTotalHitsCount) SetBlocked(v int64)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *PrivateResourceDetailedStatsTimerangeTotalHitsCount) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


