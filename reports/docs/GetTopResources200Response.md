# GetTopResources200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TopResource**](TopResource.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetTopResources200Response

`func NewGetTopResources200Response(data []TopResource, meta map[string]interface{}, ) *GetTopResources200Response`

NewGetTopResources200Response instantiates a new GetTopResources200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTopResources200ResponseWithDefaults

`func NewGetTopResources200ResponseWithDefaults() *GetTopResources200Response`

NewGetTopResources200ResponseWithDefaults instantiates a new GetTopResources200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetTopResources200Response) GetData() []TopResource`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTopResources200Response) GetDataOk() (*[]TopResource, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTopResources200Response) SetData(v []TopResource)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetTopResources200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTopResources200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTopResources200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


