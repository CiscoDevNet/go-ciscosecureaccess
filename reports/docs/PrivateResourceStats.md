# PrivateResourceStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Privateresourceid** | Pointer to **int64** | The private resource ID. | [optional] 
**Idscount** | Pointer to **int64** | The total number of identities that accessed the private resource. | [optional] 
**Hitscount** | Pointer to **int64** | The total number of times that an ednpoint accessed the private resource. | [optional] 
**Lasteventat** | Pointer to **int64** | The last time in milliseconds since the Unix Epoch of the request for the resource. | [optional] 

## Methods

### NewPrivateResourceStats

`func NewPrivateResourceStats() *PrivateResourceStats`

NewPrivateResourceStats instantiates a new PrivateResourceStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateResourceStatsWithDefaults

`func NewPrivateResourceStatsWithDefaults() *PrivateResourceStats`

NewPrivateResourceStatsWithDefaults instantiates a new PrivateResourceStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateresourceid

`func (o *PrivateResourceStats) GetPrivateresourceid() int64`

GetPrivateresourceid returns the Privateresourceid field if non-nil, zero value otherwise.

### GetPrivateresourceidOk

`func (o *PrivateResourceStats) GetPrivateresourceidOk() (*int64, bool)`

GetPrivateresourceidOk returns a tuple with the Privateresourceid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateresourceid

`func (o *PrivateResourceStats) SetPrivateresourceid(v int64)`

SetPrivateresourceid sets Privateresourceid field to given value.

### HasPrivateresourceid

`func (o *PrivateResourceStats) HasPrivateresourceid() bool`

HasPrivateresourceid returns a boolean if a field has been set.

### GetIdscount

`func (o *PrivateResourceStats) GetIdscount() int64`

GetIdscount returns the Idscount field if non-nil, zero value otherwise.

### GetIdscountOk

`func (o *PrivateResourceStats) GetIdscountOk() (*int64, bool)`

GetIdscountOk returns a tuple with the Idscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdscount

`func (o *PrivateResourceStats) SetIdscount(v int64)`

SetIdscount sets Idscount field to given value.

### HasIdscount

`func (o *PrivateResourceStats) HasIdscount() bool`

HasIdscount returns a boolean if a field has been set.

### GetHitscount

`func (o *PrivateResourceStats) GetHitscount() int64`

GetHitscount returns the Hitscount field if non-nil, zero value otherwise.

### GetHitscountOk

`func (o *PrivateResourceStats) GetHitscountOk() (*int64, bool)`

GetHitscountOk returns a tuple with the Hitscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitscount

`func (o *PrivateResourceStats) SetHitscount(v int64)`

SetHitscount sets Hitscount field to given value.

### HasHitscount

`func (o *PrivateResourceStats) HasHitscount() bool`

HasHitscount returns a boolean if a field has been set.

### GetLasteventat

`func (o *PrivateResourceStats) GetLasteventat() int64`

GetLasteventat returns the Lasteventat field if non-nil, zero value otherwise.

### GetLasteventatOk

`func (o *PrivateResourceStats) GetLasteventatOk() (*int64, bool)`

GetLasteventatOk returns a tuple with the Lasteventat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLasteventat

`func (o *PrivateResourceStats) SetLasteventat(v int64)`

SetLasteventat sets Lasteventat field to given value.

### HasLasteventat

`func (o *PrivateResourceStats) HasLasteventat() bool`

HasLasteventat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


