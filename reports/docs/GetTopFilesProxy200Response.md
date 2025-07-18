# GetTopFilesProxy200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TopFile**](TopFile.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetTopFilesProxy200Response

`func NewGetTopFilesProxy200Response(data []TopFile, meta map[string]interface{}, ) *GetTopFilesProxy200Response`

NewGetTopFilesProxy200Response instantiates a new GetTopFilesProxy200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTopFilesProxy200ResponseWithDefaults

`func NewGetTopFilesProxy200ResponseWithDefaults() *GetTopFilesProxy200Response`

NewGetTopFilesProxy200ResponseWithDefaults instantiates a new GetTopFilesProxy200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetTopFilesProxy200Response) GetData() []TopFile`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTopFilesProxy200Response) GetDataOk() (*[]TopFile, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTopFilesProxy200Response) SetData(v []TopFile)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetTopFilesProxy200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTopFilesProxy200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTopFilesProxy200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


