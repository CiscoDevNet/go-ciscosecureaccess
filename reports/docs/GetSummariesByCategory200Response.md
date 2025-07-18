# GetSummariesByCategory200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]SummaryWithCategory**](SummaryWithCategory.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetSummariesByCategory200Response

`func NewGetSummariesByCategory200Response(data []SummaryWithCategory, meta map[string]interface{}, ) *GetSummariesByCategory200Response`

NewGetSummariesByCategory200Response instantiates a new GetSummariesByCategory200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSummariesByCategory200ResponseWithDefaults

`func NewGetSummariesByCategory200ResponseWithDefaults() *GetSummariesByCategory200Response`

NewGetSummariesByCategory200ResponseWithDefaults instantiates a new GetSummariesByCategory200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetSummariesByCategory200Response) GetData() []SummaryWithCategory`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetSummariesByCategory200Response) GetDataOk() (*[]SummaryWithCategory, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetSummariesByCategory200Response) SetData(v []SummaryWithCategory)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetSummariesByCategory200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetSummariesByCategory200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetSummariesByCategory200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


