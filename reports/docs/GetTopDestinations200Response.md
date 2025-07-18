# GetTopDestinations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TopDestination**](TopDestination.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetTopDestinations200Response

`func NewGetTopDestinations200Response(data []TopDestination, meta map[string]interface{}, ) *GetTopDestinations200Response`

NewGetTopDestinations200Response instantiates a new GetTopDestinations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTopDestinations200ResponseWithDefaults

`func NewGetTopDestinations200ResponseWithDefaults() *GetTopDestinations200Response`

NewGetTopDestinations200ResponseWithDefaults instantiates a new GetTopDestinations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetTopDestinations200Response) GetData() []TopDestination`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTopDestinations200Response) GetDataOk() (*[]TopDestination, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTopDestinations200Response) SetData(v []TopDestination)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetTopDestinations200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTopDestinations200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTopDestinations200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


