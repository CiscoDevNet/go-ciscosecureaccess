# AppConnectorAgentSummaryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuUtilization** | Pointer to **int64** | The CPU utilization of the Resource Connector. | [optional] 
**Requests** | Pointer to [**TotalRequest**](TotalRequest.md) |  | [optional] 

## Methods

### NewAppConnectorAgentSummaryData

`func NewAppConnectorAgentSummaryData() *AppConnectorAgentSummaryData`

NewAppConnectorAgentSummaryData instantiates a new AppConnectorAgentSummaryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppConnectorAgentSummaryDataWithDefaults

`func NewAppConnectorAgentSummaryDataWithDefaults() *AppConnectorAgentSummaryData`

NewAppConnectorAgentSummaryDataWithDefaults instantiates a new AppConnectorAgentSummaryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuUtilization

`func (o *AppConnectorAgentSummaryData) GetCpuUtilization() int64`

GetCpuUtilization returns the CpuUtilization field if non-nil, zero value otherwise.

### GetCpuUtilizationOk

`func (o *AppConnectorAgentSummaryData) GetCpuUtilizationOk() (*int64, bool)`

GetCpuUtilizationOk returns a tuple with the CpuUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtilization

`func (o *AppConnectorAgentSummaryData) SetCpuUtilization(v int64)`

SetCpuUtilization sets CpuUtilization field to given value.

### HasCpuUtilization

`func (o *AppConnectorAgentSummaryData) HasCpuUtilization() bool`

HasCpuUtilization returns a boolean if a field has been set.

### GetRequests

`func (o *AppConnectorAgentSummaryData) GetRequests() TotalRequest`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *AppConnectorAgentSummaryData) GetRequestsOk() (*TotalRequest, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *AppConnectorAgentSummaryData) SetRequests(v TotalRequest)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *AppConnectorAgentSummaryData) HasRequests() bool`

HasRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


