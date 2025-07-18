# GetActivityIntrusion200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ActivityIntrusion**](ActivityIntrusion.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetActivityIntrusion200Response

`func NewGetActivityIntrusion200Response(data []ActivityIntrusion, meta map[string]interface{}, ) *GetActivityIntrusion200Response`

NewGetActivityIntrusion200Response instantiates a new GetActivityIntrusion200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetActivityIntrusion200ResponseWithDefaults

`func NewGetActivityIntrusion200ResponseWithDefaults() *GetActivityIntrusion200Response`

NewGetActivityIntrusion200ResponseWithDefaults instantiates a new GetActivityIntrusion200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetActivityIntrusion200Response) GetData() []ActivityIntrusion`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetActivityIntrusion200Response) GetDataOk() (*[]ActivityIntrusion, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetActivityIntrusion200Response) SetData(v []ActivityIntrusion)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetActivityIntrusion200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetActivityIntrusion200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetActivityIntrusion200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


