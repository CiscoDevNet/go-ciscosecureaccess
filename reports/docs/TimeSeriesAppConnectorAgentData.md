# TimeSeriesAppConnectorAgentData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TunnelStatus** | Pointer to [**TunnelStatus**](TunnelStatus.md) |  | [optional] 
**CpuUtilization** | Pointer to **int64** | The CPU utilization of the Resource Connector. | [optional] 
**Requests** | Pointer to [**TotalRequest**](TotalRequest.md) |  | [optional] 
**Timestamp** | Pointer to **int64** | The timestamp represented in milliseconds for the bucket. | [optional] 

## Methods

### NewTimeSeriesAppConnectorAgentData

`func NewTimeSeriesAppConnectorAgentData() *TimeSeriesAppConnectorAgentData`

NewTimeSeriesAppConnectorAgentData instantiates a new TimeSeriesAppConnectorAgentData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSeriesAppConnectorAgentDataWithDefaults

`func NewTimeSeriesAppConnectorAgentDataWithDefaults() *TimeSeriesAppConnectorAgentData`

NewTimeSeriesAppConnectorAgentDataWithDefaults instantiates a new TimeSeriesAppConnectorAgentData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTunnelStatus

`func (o *TimeSeriesAppConnectorAgentData) GetTunnelStatus() TunnelStatus`

GetTunnelStatus returns the TunnelStatus field if non-nil, zero value otherwise.

### GetTunnelStatusOk

`func (o *TimeSeriesAppConnectorAgentData) GetTunnelStatusOk() (*TunnelStatus, bool)`

GetTunnelStatusOk returns a tuple with the TunnelStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelStatus

`func (o *TimeSeriesAppConnectorAgentData) SetTunnelStatus(v TunnelStatus)`

SetTunnelStatus sets TunnelStatus field to given value.

### HasTunnelStatus

`func (o *TimeSeriesAppConnectorAgentData) HasTunnelStatus() bool`

HasTunnelStatus returns a boolean if a field has been set.

### GetCpuUtilization

`func (o *TimeSeriesAppConnectorAgentData) GetCpuUtilization() int64`

GetCpuUtilization returns the CpuUtilization field if non-nil, zero value otherwise.

### GetCpuUtilizationOk

`func (o *TimeSeriesAppConnectorAgentData) GetCpuUtilizationOk() (*int64, bool)`

GetCpuUtilizationOk returns a tuple with the CpuUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtilization

`func (o *TimeSeriesAppConnectorAgentData) SetCpuUtilization(v int64)`

SetCpuUtilization sets CpuUtilization field to given value.

### HasCpuUtilization

`func (o *TimeSeriesAppConnectorAgentData) HasCpuUtilization() bool`

HasCpuUtilization returns a boolean if a field has been set.

### GetRequests

`func (o *TimeSeriesAppConnectorAgentData) GetRequests() TotalRequest`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *TimeSeriesAppConnectorAgentData) GetRequestsOk() (*TotalRequest, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *TimeSeriesAppConnectorAgentData) SetRequests(v TotalRequest)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *TimeSeriesAppConnectorAgentData) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetTimestamp

`func (o *TimeSeriesAppConnectorAgentData) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TimeSeriesAppConnectorAgentData) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TimeSeriesAppConnectorAgentData) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TimeSeriesAppConnectorAgentData) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


