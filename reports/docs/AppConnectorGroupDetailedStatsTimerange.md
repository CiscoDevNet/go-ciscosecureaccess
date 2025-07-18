# AppConnectorGroupDetailedStatsTimerange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **int64** | The group ID of the Resource Connector. | [optional] 
**SummaryData** | Pointer to [**AppConnectorGroupSummaryData**](AppConnectorGroupSummaryData.md) |  | [optional] 
**TimeSeriesData** | Pointer to [**[]TimeSeriesAppConnectorGroupData**](TimeSeriesAppConnectorGroupData.md) | The list of the time series data. | [optional] 

## Methods

### NewAppConnectorGroupDetailedStatsTimerange

`func NewAppConnectorGroupDetailedStatsTimerange() *AppConnectorGroupDetailedStatsTimerange`

NewAppConnectorGroupDetailedStatsTimerange instantiates a new AppConnectorGroupDetailedStatsTimerange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppConnectorGroupDetailedStatsTimerangeWithDefaults

`func NewAppConnectorGroupDetailedStatsTimerangeWithDefaults() *AppConnectorGroupDetailedStatsTimerange`

NewAppConnectorGroupDetailedStatsTimerangeWithDefaults instantiates a new AppConnectorGroupDetailedStatsTimerange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *AppConnectorGroupDetailedStatsTimerange) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *AppConnectorGroupDetailedStatsTimerange) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *AppConnectorGroupDetailedStatsTimerange) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *AppConnectorGroupDetailedStatsTimerange) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetSummaryData

`func (o *AppConnectorGroupDetailedStatsTimerange) GetSummaryData() AppConnectorGroupSummaryData`

GetSummaryData returns the SummaryData field if non-nil, zero value otherwise.

### GetSummaryDataOk

`func (o *AppConnectorGroupDetailedStatsTimerange) GetSummaryDataOk() (*AppConnectorGroupSummaryData, bool)`

GetSummaryDataOk returns a tuple with the SummaryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryData

`func (o *AppConnectorGroupDetailedStatsTimerange) SetSummaryData(v AppConnectorGroupSummaryData)`

SetSummaryData sets SummaryData field to given value.

### HasSummaryData

`func (o *AppConnectorGroupDetailedStatsTimerange) HasSummaryData() bool`

HasSummaryData returns a boolean if a field has been set.

### GetTimeSeriesData

`func (o *AppConnectorGroupDetailedStatsTimerange) GetTimeSeriesData() []TimeSeriesAppConnectorGroupData`

GetTimeSeriesData returns the TimeSeriesData field if non-nil, zero value otherwise.

### GetTimeSeriesDataOk

`func (o *AppConnectorGroupDetailedStatsTimerange) GetTimeSeriesDataOk() (*[]TimeSeriesAppConnectorGroupData, bool)`

GetTimeSeriesDataOk returns a tuple with the TimeSeriesData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSeriesData

`func (o *AppConnectorGroupDetailedStatsTimerange) SetTimeSeriesData(v []TimeSeriesAppConnectorGroupData)`

SetTimeSeriesData sets TimeSeriesData field to given value.

### HasTimeSeriesData

`func (o *AppConnectorGroupDetailedStatsTimerange) HasTimeSeriesData() bool`

HasTimeSeriesData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


