# GetTopCategories200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TopCategory**](TopCategory.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetTopCategories200Response

`func NewGetTopCategories200Response(data []TopCategory, meta map[string]interface{}, ) *GetTopCategories200Response`

NewGetTopCategories200Response instantiates a new GetTopCategories200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTopCategories200ResponseWithDefaults

`func NewGetTopCategories200ResponseWithDefaults() *GetTopCategories200Response`

NewGetTopCategories200ResponseWithDefaults instantiates a new GetTopCategories200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetTopCategories200Response) GetData() []TopCategory`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTopCategories200Response) GetDataOk() (*[]TopCategory, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTopCategories200Response) SetData(v []TopCategory)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetTopCategories200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTopCategories200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTopCategories200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


