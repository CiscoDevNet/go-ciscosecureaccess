# GetSummariesByDestination200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]SummaryWithDestination**](SummaryWithDestination.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetSummariesByDestination200Response

`func NewGetSummariesByDestination200Response(data []SummaryWithDestination, meta map[string]interface{}, ) *GetSummariesByDestination200Response`

NewGetSummariesByDestination200Response instantiates a new GetSummariesByDestination200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSummariesByDestination200ResponseWithDefaults

`func NewGetSummariesByDestination200ResponseWithDefaults() *GetSummariesByDestination200Response`

NewGetSummariesByDestination200ResponseWithDefaults instantiates a new GetSummariesByDestination200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetSummariesByDestination200Response) GetData() []SummaryWithDestination`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetSummariesByDestination200Response) GetDataOk() (*[]SummaryWithDestination, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetSummariesByDestination200Response) SetData(v []SummaryWithDestination)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetSummariesByDestination200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetSummariesByDestination200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetSummariesByDestination200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


