# GetActivityDns200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ActivityDns**](ActivityDns.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetActivityDns200Response

`func NewGetActivityDns200Response(data []ActivityDns, meta map[string]interface{}, ) *GetActivityDns200Response`

NewGetActivityDns200Response instantiates a new GetActivityDns200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetActivityDns200ResponseWithDefaults

`func NewGetActivityDns200ResponseWithDefaults() *GetActivityDns200Response`

NewGetActivityDns200ResponseWithDefaults instantiates a new GetActivityDns200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetActivityDns200Response) GetData() []ActivityDns`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetActivityDns200Response) GetDataOk() (*[]ActivityDns, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetActivityDns200Response) SetData(v []ActivityDns)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetActivityDns200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetActivityDns200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetActivityDns200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


