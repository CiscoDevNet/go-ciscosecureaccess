# GetActivityFirewall200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ActivityFirewall**](ActivityFirewall.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetActivityFirewall200Response

`func NewGetActivityFirewall200Response(data []ActivityFirewall, meta map[string]interface{}, ) *GetActivityFirewall200Response`

NewGetActivityFirewall200Response instantiates a new GetActivityFirewall200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetActivityFirewall200ResponseWithDefaults

`func NewGetActivityFirewall200ResponseWithDefaults() *GetActivityFirewall200Response`

NewGetActivityFirewall200ResponseWithDefaults instantiates a new GetActivityFirewall200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetActivityFirewall200Response) GetData() []ActivityFirewall`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetActivityFirewall200Response) GetDataOk() (*[]ActivityFirewall, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetActivityFirewall200Response) SetData(v []ActivityFirewall)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetActivityFirewall200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetActivityFirewall200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetActivityFirewall200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


