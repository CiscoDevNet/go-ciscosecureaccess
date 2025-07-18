# GetTopIdentities200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TopIdentity**](TopIdentity.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetTopIdentities200Response

`func NewGetTopIdentities200Response(data []TopIdentity, meta map[string]interface{}, ) *GetTopIdentities200Response`

NewGetTopIdentities200Response instantiates a new GetTopIdentities200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTopIdentities200ResponseWithDefaults

`func NewGetTopIdentities200ResponseWithDefaults() *GetTopIdentities200Response`

NewGetTopIdentities200ResponseWithDefaults instantiates a new GetTopIdentities200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetTopIdentities200Response) GetData() []TopIdentity`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTopIdentities200Response) GetDataOk() (*[]TopIdentity, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTopIdentities200Response) SetData(v []TopIdentity)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetTopIdentities200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTopIdentities200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTopIdentities200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


