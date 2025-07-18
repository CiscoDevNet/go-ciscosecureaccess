# GetTopIps200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TopIP**](TopIP.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetTopIps200Response

`func NewGetTopIps200Response(data []TopIP, meta map[string]interface{}, ) *GetTopIps200Response`

NewGetTopIps200Response instantiates a new GetTopIps200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTopIps200ResponseWithDefaults

`func NewGetTopIps200ResponseWithDefaults() *GetTopIps200Response`

NewGetTopIps200ResponseWithDefaults instantiates a new GetTopIps200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetTopIps200Response) GetData() []TopIP`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTopIps200Response) GetDataOk() (*[]TopIP, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTopIps200Response) SetData(v []TopIP)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetTopIps200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTopIps200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTopIps200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


