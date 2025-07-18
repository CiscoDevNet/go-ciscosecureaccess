# GetTopThreats200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TopThreats**](TopThreats.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetTopThreats200Response

`func NewGetTopThreats200Response(data []TopThreats, meta map[string]interface{}, ) *GetTopThreats200Response`

NewGetTopThreats200Response instantiates a new GetTopThreats200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTopThreats200ResponseWithDefaults

`func NewGetTopThreats200ResponseWithDefaults() *GetTopThreats200Response`

NewGetTopThreats200ResponseWithDefaults instantiates a new GetTopThreats200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetTopThreats200Response) GetData() []TopThreats`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTopThreats200Response) GetDataOk() (*[]TopThreats, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTopThreats200Response) SetData(v []TopThreats)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetTopThreats200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTopThreats200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTopThreats200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


