# GetActivityProxy200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ActivityProxy**](ActivityProxy.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetActivityProxy200Response

`func NewGetActivityProxy200Response(data []ActivityProxy, meta map[string]interface{}, ) *GetActivityProxy200Response`

NewGetActivityProxy200Response instantiates a new GetActivityProxy200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetActivityProxy200ResponseWithDefaults

`func NewGetActivityProxy200ResponseWithDefaults() *GetActivityProxy200Response`

NewGetActivityProxy200ResponseWithDefaults instantiates a new GetActivityProxy200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetActivityProxy200Response) GetData() []ActivityProxy`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetActivityProxy200Response) GetDataOk() (*[]ActivityProxy, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetActivityProxy200Response) SetData(v []ActivityProxy)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetActivityProxy200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetActivityProxy200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetActivityProxy200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


