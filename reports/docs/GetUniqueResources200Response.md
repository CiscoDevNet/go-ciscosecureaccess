# GetUniqueResources200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**UniqueResources**](UniqueResources.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetUniqueResources200Response

`func NewGetUniqueResources200Response(data UniqueResources, meta map[string]interface{}, ) *GetUniqueResources200Response`

NewGetUniqueResources200Response instantiates a new GetUniqueResources200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUniqueResources200ResponseWithDefaults

`func NewGetUniqueResources200ResponseWithDefaults() *GetUniqueResources200Response`

NewGetUniqueResources200ResponseWithDefaults instantiates a new GetUniqueResources200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetUniqueResources200Response) GetData() UniqueResources`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetUniqueResources200Response) GetDataOk() (*UniqueResources, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetUniqueResources200Response) SetData(v UniqueResources)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetUniqueResources200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetUniqueResources200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetUniqueResources200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


