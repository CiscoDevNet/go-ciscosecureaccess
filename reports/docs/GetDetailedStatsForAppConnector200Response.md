# GetDetailedStatsForAppConnector200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AppConnectorAgentDetailedStatsTimerange**](AppConnectorAgentDetailedStatsTimerange.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetDetailedStatsForAppConnector200Response

`func NewGetDetailedStatsForAppConnector200Response(data []AppConnectorAgentDetailedStatsTimerange, meta map[string]interface{}, ) *GetDetailedStatsForAppConnector200Response`

NewGetDetailedStatsForAppConnector200Response instantiates a new GetDetailedStatsForAppConnector200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDetailedStatsForAppConnector200ResponseWithDefaults

`func NewGetDetailedStatsForAppConnector200ResponseWithDefaults() *GetDetailedStatsForAppConnector200Response`

NewGetDetailedStatsForAppConnector200ResponseWithDefaults instantiates a new GetDetailedStatsForAppConnector200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetDetailedStatsForAppConnector200Response) GetData() []AppConnectorAgentDetailedStatsTimerange`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetDetailedStatsForAppConnector200Response) GetDataOk() (*[]AppConnectorAgentDetailedStatsTimerange, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetDetailedStatsForAppConnector200Response) SetData(v []AppConnectorAgentDetailedStatsTimerange)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetDetailedStatsForAppConnector200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDetailedStatsForAppConnector200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDetailedStatsForAppConnector200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


