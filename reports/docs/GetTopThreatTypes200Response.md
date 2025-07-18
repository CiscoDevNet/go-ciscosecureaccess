# GetTopThreatTypes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TopThreatTypes**](TopThreatTypes.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetTopThreatTypes200Response

`func NewGetTopThreatTypes200Response(data []TopThreatTypes, meta map[string]interface{}, ) *GetTopThreatTypes200Response`

NewGetTopThreatTypes200Response instantiates a new GetTopThreatTypes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTopThreatTypes200ResponseWithDefaults

`func NewGetTopThreatTypes200ResponseWithDefaults() *GetTopThreatTypes200Response`

NewGetTopThreatTypes200ResponseWithDefaults instantiates a new GetTopThreatTypes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetTopThreatTypes200Response) GetData() []TopThreatTypes`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTopThreatTypes200Response) GetDataOk() (*[]TopThreatTypes, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTopThreatTypes200Response) SetData(v []TopThreatTypes)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetTopThreatTypes200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTopThreatTypes200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTopThreatTypes200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


