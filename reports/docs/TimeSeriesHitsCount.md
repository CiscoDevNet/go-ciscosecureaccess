# TimeSeriesHitsCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hitscount** | Pointer to **int64** | The total number of times that the user or device accessed the private resource in a specific time period. | [optional] 
**Success** | Pointer to **int64** | The number of times that the user or device was allowed to access the private resource in a specific time period. | [optional] 
**Blocked** | Pointer to **int64** | The number of times that the user or device was blocked to access the private resource in a specific time period. | [optional] 
**Timestamp** | Pointer to **int64** | The timestamp represented in milliseconds for the bucket. | [optional] 
**Date** | Pointer to **string** | The date from the timestamp based on the timezone parameter. | [optional] 
**Time** | Pointer to **string** | The time in 24-hour format based on the timezone parameter. | [optional] 

## Methods

### NewTimeSeriesHitsCount

`func NewTimeSeriesHitsCount() *TimeSeriesHitsCount`

NewTimeSeriesHitsCount instantiates a new TimeSeriesHitsCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSeriesHitsCountWithDefaults

`func NewTimeSeriesHitsCountWithDefaults() *TimeSeriesHitsCount`

NewTimeSeriesHitsCountWithDefaults instantiates a new TimeSeriesHitsCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHitscount

`func (o *TimeSeriesHitsCount) GetHitscount() int64`

GetHitscount returns the Hitscount field if non-nil, zero value otherwise.

### GetHitscountOk

`func (o *TimeSeriesHitsCount) GetHitscountOk() (*int64, bool)`

GetHitscountOk returns a tuple with the Hitscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitscount

`func (o *TimeSeriesHitsCount) SetHitscount(v int64)`

SetHitscount sets Hitscount field to given value.

### HasHitscount

`func (o *TimeSeriesHitsCount) HasHitscount() bool`

HasHitscount returns a boolean if a field has been set.

### GetSuccess

`func (o *TimeSeriesHitsCount) GetSuccess() int64`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *TimeSeriesHitsCount) GetSuccessOk() (*int64, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *TimeSeriesHitsCount) SetSuccess(v int64)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *TimeSeriesHitsCount) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetBlocked

`func (o *TimeSeriesHitsCount) GetBlocked() int64`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *TimeSeriesHitsCount) GetBlockedOk() (*int64, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *TimeSeriesHitsCount) SetBlocked(v int64)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *TimeSeriesHitsCount) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetTimestamp

`func (o *TimeSeriesHitsCount) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TimeSeriesHitsCount) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TimeSeriesHitsCount) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TimeSeriesHitsCount) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetDate

`func (o *TimeSeriesHitsCount) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TimeSeriesHitsCount) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TimeSeriesHitsCount) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *TimeSeriesHitsCount) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetTime

`func (o *TimeSeriesHitsCount) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *TimeSeriesHitsCount) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *TimeSeriesHitsCount) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *TimeSeriesHitsCount) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


