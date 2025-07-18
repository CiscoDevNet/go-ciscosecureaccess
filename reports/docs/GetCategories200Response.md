# GetCategories200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CategoryWithLegacyId**](CategoryWithLegacyId.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetCategories200Response

`func NewGetCategories200Response(data []CategoryWithLegacyId, meta map[string]interface{}, ) *GetCategories200Response`

NewGetCategories200Response instantiates a new GetCategories200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCategories200ResponseWithDefaults

`func NewGetCategories200ResponseWithDefaults() *GetCategories200Response`

NewGetCategories200ResponseWithDefaults instantiates a new GetCategories200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetCategories200Response) GetData() []CategoryWithLegacyId`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCategories200Response) GetDataOk() (*[]CategoryWithLegacyId, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCategories200Response) SetData(v []CategoryWithLegacyId)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetCategories200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetCategories200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetCategories200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


