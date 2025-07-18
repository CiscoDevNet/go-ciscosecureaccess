# GetTopUrls200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TopURL**](TopURL.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetTopUrls200Response

`func NewGetTopUrls200Response(data []TopURL, meta map[string]interface{}, ) *GetTopUrls200Response`

NewGetTopUrls200Response instantiates a new GetTopUrls200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTopUrls200ResponseWithDefaults

`func NewGetTopUrls200ResponseWithDefaults() *GetTopUrls200Response`

NewGetTopUrls200ResponseWithDefaults instantiates a new GetTopUrls200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetTopUrls200Response) GetData() []TopURL`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTopUrls200Response) GetDataOk() (*[]TopURL, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTopUrls200Response) SetData(v []TopURL)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetTopUrls200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTopUrls200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTopUrls200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


