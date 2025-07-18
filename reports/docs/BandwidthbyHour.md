# BandwidthbyHour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inboundbytes** | **int64** | The number of inbound bytes. | 
**Outboundbytes** | **NullableInt64** | The number of outbound bytes. | 
**Timestamp** | **int64** | The timestamp represented in milliseconds for the bucket. | 
**Date** | **string** | The date from the timestamp based on the timezone parameter. | 
**Time** | **string** | The time in 24-hour format based on the timezone parameter. | 

## Methods

### NewBandwidthbyHour

`func NewBandwidthbyHour(inboundbytes int64, outboundbytes NullableInt64, timestamp int64, date string, time string, ) *BandwidthbyHour`

NewBandwidthbyHour instantiates a new BandwidthbyHour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBandwidthbyHourWithDefaults

`func NewBandwidthbyHourWithDefaults() *BandwidthbyHour`

NewBandwidthbyHourWithDefaults instantiates a new BandwidthbyHour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInboundbytes

`func (o *BandwidthbyHour) GetInboundbytes() int64`

GetInboundbytes returns the Inboundbytes field if non-nil, zero value otherwise.

### GetInboundbytesOk

`func (o *BandwidthbyHour) GetInboundbytesOk() (*int64, bool)`

GetInboundbytesOk returns a tuple with the Inboundbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundbytes

`func (o *BandwidthbyHour) SetInboundbytes(v int64)`

SetInboundbytes sets Inboundbytes field to given value.


### GetOutboundbytes

`func (o *BandwidthbyHour) GetOutboundbytes() int64`

GetOutboundbytes returns the Outboundbytes field if non-nil, zero value otherwise.

### GetOutboundbytesOk

`func (o *BandwidthbyHour) GetOutboundbytesOk() (*int64, bool)`

GetOutboundbytesOk returns a tuple with the Outboundbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundbytes

`func (o *BandwidthbyHour) SetOutboundbytes(v int64)`

SetOutboundbytes sets Outboundbytes field to given value.


### SetOutboundbytesNil

`func (o *BandwidthbyHour) SetOutboundbytesNil(b bool)`

 SetOutboundbytesNil sets the value for Outboundbytes to be an explicit nil

### UnsetOutboundbytes
`func (o *BandwidthbyHour) UnsetOutboundbytes()`

UnsetOutboundbytes ensures that no value is present for Outboundbytes, not even an explicit nil
### GetTimestamp

`func (o *BandwidthbyHour) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BandwidthbyHour) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BandwidthbyHour) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetDate

`func (o *BandwidthbyHour) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *BandwidthbyHour) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *BandwidthbyHour) SetDate(v string)`

SetDate sets Date field to given value.


### GetTime

`func (o *BandwidthbyHour) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *BandwidthbyHour) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *BandwidthbyHour) SetTime(v string)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


