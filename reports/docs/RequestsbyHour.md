# RequestsbyHour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** | The number of requests in the hour. | 
**Counts** | [**RequestAndConnectionCounts**](RequestAndConnectionCounts.md) |  | 
**Timestamp** | **int64** | The timestamp represented in milliseconds for the bucket. | 
**Date** | **string** | The date from the timestamp based on the timezone parameter. | 
**Time** | **string** | The time in 24-hour format based on the timezone parameter. | 

## Methods

### NewRequestsbyHour

`func NewRequestsbyHour(count int64, counts RequestAndConnectionCounts, timestamp int64, date string, time string, ) *RequestsbyHour`

NewRequestsbyHour instantiates a new RequestsbyHour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestsbyHourWithDefaults

`func NewRequestsbyHourWithDefaults() *RequestsbyHour`

NewRequestsbyHourWithDefaults instantiates a new RequestsbyHour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *RequestsbyHour) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RequestsbyHour) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RequestsbyHour) SetCount(v int64)`

SetCount sets Count field to given value.


### GetCounts

`func (o *RequestsbyHour) GetCounts() RequestAndConnectionCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *RequestsbyHour) GetCountsOk() (*RequestAndConnectionCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *RequestsbyHour) SetCounts(v RequestAndConnectionCounts)`

SetCounts sets Counts field to given value.


### GetTimestamp

`func (o *RequestsbyHour) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RequestsbyHour) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RequestsbyHour) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetDate

`func (o *RequestsbyHour) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *RequestsbyHour) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *RequestsbyHour) SetDate(v string)`

SetDate sets Date field to given value.


### GetTime

`func (o *RequestsbyHour) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *RequestsbyHour) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *RequestsbyHour) SetTime(v string)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


