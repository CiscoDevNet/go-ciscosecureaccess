# GetIdentities200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Identity**](Identity.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetIdentities200Response

`func NewGetIdentities200Response(data []Identity, meta map[string]interface{}, ) *GetIdentities200Response`

NewGetIdentities200Response instantiates a new GetIdentities200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIdentities200ResponseWithDefaults

`func NewGetIdentities200ResponseWithDefaults() *GetIdentities200Response`

NewGetIdentities200ResponseWithDefaults instantiates a new GetIdentities200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetIdentities200Response) GetData() []Identity`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetIdentities200Response) GetDataOk() (*[]Identity, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetIdentities200Response) SetData(v []Identity)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetIdentities200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetIdentities200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetIdentities200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


