# PrivateResourceDetailedStatsTimerange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Privateresourceid** | Pointer to **int64** | The private resource ID. | [optional] 
**TotalHitsCount** | Pointer to [**PrivateResourceDetailedStatsTimerangeTotalHitsCount**](PrivateResourceDetailedStatsTimerangeTotalHitsCount.md) |  | [optional] 
**TimeSeriesHitsCount** | Pointer to [**[]TimeSeriesHitsCount**](TimeSeriesHitsCount.md) | The list of the counts and timestamps for the hit counts. | [optional] 

## Methods

### NewPrivateResourceDetailedStatsTimerange

`func NewPrivateResourceDetailedStatsTimerange() *PrivateResourceDetailedStatsTimerange`

NewPrivateResourceDetailedStatsTimerange instantiates a new PrivateResourceDetailedStatsTimerange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateResourceDetailedStatsTimerangeWithDefaults

`func NewPrivateResourceDetailedStatsTimerangeWithDefaults() *PrivateResourceDetailedStatsTimerange`

NewPrivateResourceDetailedStatsTimerangeWithDefaults instantiates a new PrivateResourceDetailedStatsTimerange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateresourceid

`func (o *PrivateResourceDetailedStatsTimerange) GetPrivateresourceid() int64`

GetPrivateresourceid returns the Privateresourceid field if non-nil, zero value otherwise.

### GetPrivateresourceidOk

`func (o *PrivateResourceDetailedStatsTimerange) GetPrivateresourceidOk() (*int64, bool)`

GetPrivateresourceidOk returns a tuple with the Privateresourceid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateresourceid

`func (o *PrivateResourceDetailedStatsTimerange) SetPrivateresourceid(v int64)`

SetPrivateresourceid sets Privateresourceid field to given value.

### HasPrivateresourceid

`func (o *PrivateResourceDetailedStatsTimerange) HasPrivateresourceid() bool`

HasPrivateresourceid returns a boolean if a field has been set.

### GetTotalHitsCount

`func (o *PrivateResourceDetailedStatsTimerange) GetTotalHitsCount() PrivateResourceDetailedStatsTimerangeTotalHitsCount`

GetTotalHitsCount returns the TotalHitsCount field if non-nil, zero value otherwise.

### GetTotalHitsCountOk

`func (o *PrivateResourceDetailedStatsTimerange) GetTotalHitsCountOk() (*PrivateResourceDetailedStatsTimerangeTotalHitsCount, bool)`

GetTotalHitsCountOk returns a tuple with the TotalHitsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHitsCount

`func (o *PrivateResourceDetailedStatsTimerange) SetTotalHitsCount(v PrivateResourceDetailedStatsTimerangeTotalHitsCount)`

SetTotalHitsCount sets TotalHitsCount field to given value.

### HasTotalHitsCount

`func (o *PrivateResourceDetailedStatsTimerange) HasTotalHitsCount() bool`

HasTotalHitsCount returns a boolean if a field has been set.

### GetTimeSeriesHitsCount

`func (o *PrivateResourceDetailedStatsTimerange) GetTimeSeriesHitsCount() []TimeSeriesHitsCount`

GetTimeSeriesHitsCount returns the TimeSeriesHitsCount field if non-nil, zero value otherwise.

### GetTimeSeriesHitsCountOk

`func (o *PrivateResourceDetailedStatsTimerange) GetTimeSeriesHitsCountOk() (*[]TimeSeriesHitsCount, bool)`

GetTimeSeriesHitsCountOk returns a tuple with the TimeSeriesHitsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSeriesHitsCount

`func (o *PrivateResourceDetailedStatsTimerange) SetTimeSeriesHitsCount(v []TimeSeriesHitsCount)`

SetTimeSeriesHitsCount sets TimeSeriesHitsCount field to given value.

### HasTimeSeriesHitsCount

`func (o *PrivateResourceDetailedStatsTimerange) HasTimeSeriesHitsCount() bool`

HasTimeSeriesHitsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


