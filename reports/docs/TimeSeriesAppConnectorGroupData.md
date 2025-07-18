# TimeSeriesAppConnectorGroupData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuUtilization** | Pointer to **int64** | The CPU utilization of the Resource Connector. | [optional] 
**Timestamp** | Pointer to **int64** | The timestamp represented in milliseconds for the bucket. | [optional] 

## Methods

### NewTimeSeriesAppConnectorGroupData

`func NewTimeSeriesAppConnectorGroupData() *TimeSeriesAppConnectorGroupData`

NewTimeSeriesAppConnectorGroupData instantiates a new TimeSeriesAppConnectorGroupData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSeriesAppConnectorGroupDataWithDefaults

`func NewTimeSeriesAppConnectorGroupDataWithDefaults() *TimeSeriesAppConnectorGroupData`

NewTimeSeriesAppConnectorGroupDataWithDefaults instantiates a new TimeSeriesAppConnectorGroupData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuUtilization

`func (o *TimeSeriesAppConnectorGroupData) GetCpuUtilization() int64`

GetCpuUtilization returns the CpuUtilization field if non-nil, zero value otherwise.

### GetCpuUtilizationOk

`func (o *TimeSeriesAppConnectorGroupData) GetCpuUtilizationOk() (*int64, bool)`

GetCpuUtilizationOk returns a tuple with the CpuUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtilization

`func (o *TimeSeriesAppConnectorGroupData) SetCpuUtilization(v int64)`

SetCpuUtilization sets CpuUtilization field to given value.

### HasCpuUtilization

`func (o *TimeSeriesAppConnectorGroupData) HasCpuUtilization() bool`

HasCpuUtilization returns a boolean if a field has been set.

### GetTimestamp

`func (o *TimeSeriesAppConnectorGroupData) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TimeSeriesAppConnectorGroupData) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TimeSeriesAppConnectorGroupData) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TimeSeriesAppConnectorGroupData) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


