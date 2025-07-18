# AppConnectorAgentDetailedStatsTimerange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **int64** | The group ID of the Resource Connector. | [optional] 
**AgentId** | Pointer to **int64** | The ID of the Resource Connector. | [optional] 
**SummaryData** | Pointer to [**AppConnectorAgentSummaryData**](AppConnectorAgentSummaryData.md) |  | [optional] 
**TimeSeriesData** | Pointer to [**[]TimeSeriesAppConnectorAgentData**](TimeSeriesAppConnectorAgentData.md) | The list of time series Resource Connector data. | [optional] 

## Methods

### NewAppConnectorAgentDetailedStatsTimerange

`func NewAppConnectorAgentDetailedStatsTimerange() *AppConnectorAgentDetailedStatsTimerange`

NewAppConnectorAgentDetailedStatsTimerange instantiates a new AppConnectorAgentDetailedStatsTimerange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppConnectorAgentDetailedStatsTimerangeWithDefaults

`func NewAppConnectorAgentDetailedStatsTimerangeWithDefaults() *AppConnectorAgentDetailedStatsTimerange`

NewAppConnectorAgentDetailedStatsTimerangeWithDefaults instantiates a new AppConnectorAgentDetailedStatsTimerange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *AppConnectorAgentDetailedStatsTimerange) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *AppConnectorAgentDetailedStatsTimerange) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *AppConnectorAgentDetailedStatsTimerange) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *AppConnectorAgentDetailedStatsTimerange) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetAgentId

`func (o *AppConnectorAgentDetailedStatsTimerange) GetAgentId() int64`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *AppConnectorAgentDetailedStatsTimerange) GetAgentIdOk() (*int64, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *AppConnectorAgentDetailedStatsTimerange) SetAgentId(v int64)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *AppConnectorAgentDetailedStatsTimerange) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetSummaryData

`func (o *AppConnectorAgentDetailedStatsTimerange) GetSummaryData() AppConnectorAgentSummaryData`

GetSummaryData returns the SummaryData field if non-nil, zero value otherwise.

### GetSummaryDataOk

`func (o *AppConnectorAgentDetailedStatsTimerange) GetSummaryDataOk() (*AppConnectorAgentSummaryData, bool)`

GetSummaryDataOk returns a tuple with the SummaryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryData

`func (o *AppConnectorAgentDetailedStatsTimerange) SetSummaryData(v AppConnectorAgentSummaryData)`

SetSummaryData sets SummaryData field to given value.

### HasSummaryData

`func (o *AppConnectorAgentDetailedStatsTimerange) HasSummaryData() bool`

HasSummaryData returns a boolean if a field has been set.

### GetTimeSeriesData

`func (o *AppConnectorAgentDetailedStatsTimerange) GetTimeSeriesData() []TimeSeriesAppConnectorAgentData`

GetTimeSeriesData returns the TimeSeriesData field if non-nil, zero value otherwise.

### GetTimeSeriesDataOk

`func (o *AppConnectorAgentDetailedStatsTimerange) GetTimeSeriesDataOk() (*[]TimeSeriesAppConnectorAgentData, bool)`

GetTimeSeriesDataOk returns a tuple with the TimeSeriesData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSeriesData

`func (o *AppConnectorAgentDetailedStatsTimerange) SetTimeSeriesData(v []TimeSeriesAppConnectorAgentData)`

SetTimeSeriesData sets TimeSeriesData field to given value.

### HasTimeSeriesData

`func (o *AppConnectorAgentDetailedStatsTimerange) HasTimeSeriesData() bool`

HasTimeSeriesData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


