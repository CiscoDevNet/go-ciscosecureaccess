# GetRemoteAccessEvents200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]RemoteAccess**](RemoteAccess.md) | The list of remote access events. | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetRemoteAccessEvents200Response

`func NewGetRemoteAccessEvents200Response(data []RemoteAccess, meta map[string]interface{}, ) *GetRemoteAccessEvents200Response`

NewGetRemoteAccessEvents200Response instantiates a new GetRemoteAccessEvents200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRemoteAccessEvents200ResponseWithDefaults

`func NewGetRemoteAccessEvents200ResponseWithDefaults() *GetRemoteAccessEvents200Response`

NewGetRemoteAccessEvents200ResponseWithDefaults instantiates a new GetRemoteAccessEvents200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetRemoteAccessEvents200Response) GetData() []RemoteAccess`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetRemoteAccessEvents200Response) GetDataOk() (*[]RemoteAccess, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetRemoteAccessEvents200Response) SetData(v []RemoteAccess)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetRemoteAccessEvents200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetRemoteAccessEvents200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetRemoteAccessEvents200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


