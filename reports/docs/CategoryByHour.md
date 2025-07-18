# CategoryByHour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **int64** | The timestamp represented in milliseconds for the bucket. | 
**Date** | **string** | The date from the timestamp based on the timezone parameter. | 
**Time** | **string** | The time in 24-hour format based on the timezone parameter. | 
**Counts** | [**[]CategoryCount**](CategoryCount.md) | The list of counts for the category. | 

## Methods

### NewCategoryByHour

`func NewCategoryByHour(timestamp int64, date string, time string, counts []CategoryCount, ) *CategoryByHour`

NewCategoryByHour instantiates a new CategoryByHour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryByHourWithDefaults

`func NewCategoryByHourWithDefaults() *CategoryByHour`

NewCategoryByHourWithDefaults instantiates a new CategoryByHour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *CategoryByHour) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CategoryByHour) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CategoryByHour) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetDate

`func (o *CategoryByHour) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CategoryByHour) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CategoryByHour) SetDate(v string)`

SetDate sets Date field to given value.


### GetTime

`func (o *CategoryByHour) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *CategoryByHour) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *CategoryByHour) SetTime(v string)`

SetTime sets Time field to given value.


### GetCounts

`func (o *CategoryByHour) GetCounts() []CategoryCount`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *CategoryByHour) GetCountsOk() (*[]CategoryCount, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *CategoryByHour) SetCounts(v []CategoryCount)`

SetCounts sets Counts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


