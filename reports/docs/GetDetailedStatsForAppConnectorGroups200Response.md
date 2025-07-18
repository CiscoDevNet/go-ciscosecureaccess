# GetDetailedStatsForAppConnectorGroups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AppConnectorGroupDetailedStatsTimerange**](AppConnectorGroupDetailedStatsTimerange.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetDetailedStatsForAppConnectorGroups200Response

`func NewGetDetailedStatsForAppConnectorGroups200Response(data []AppConnectorGroupDetailedStatsTimerange, meta map[string]interface{}, ) *GetDetailedStatsForAppConnectorGroups200Response`

NewGetDetailedStatsForAppConnectorGroups200Response instantiates a new GetDetailedStatsForAppConnectorGroups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDetailedStatsForAppConnectorGroups200ResponseWithDefaults

`func NewGetDetailedStatsForAppConnectorGroups200ResponseWithDefaults() *GetDetailedStatsForAppConnectorGroups200Response`

NewGetDetailedStatsForAppConnectorGroups200ResponseWithDefaults instantiates a new GetDetailedStatsForAppConnectorGroups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetDetailedStatsForAppConnectorGroups200Response) GetData() []AppConnectorGroupDetailedStatsTimerange`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetDetailedStatsForAppConnectorGroups200Response) GetDataOk() (*[]AppConnectorGroupDetailedStatsTimerange, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetDetailedStatsForAppConnectorGroups200Response) SetData(v []AppConnectorGroupDetailedStatsTimerange)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetDetailedStatsForAppConnectorGroups200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDetailedStatsForAppConnectorGroups200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDetailedStatsForAppConnectorGroups200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


