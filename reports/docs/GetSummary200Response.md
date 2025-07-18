# GetSummary200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**Summary**](Summary.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetSummary200Response

`func NewGetSummary200Response(data Summary, meta map[string]interface{}, ) *GetSummary200Response`

NewGetSummary200Response instantiates a new GetSummary200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSummary200ResponseWithDefaults

`func NewGetSummary200ResponseWithDefaults() *GetSummary200Response`

NewGetSummary200ResponseWithDefaults instantiates a new GetSummary200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetSummary200Response) GetData() Summary`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetSummary200Response) GetDataOk() (*Summary, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetSummary200Response) SetData(v Summary)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetSummary200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetSummary200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetSummary200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


