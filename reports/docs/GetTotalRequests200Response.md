# GetTotalRequests200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**TotalRequest**](TotalRequest.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetTotalRequests200Response

`func NewGetTotalRequests200Response(data TotalRequest, meta map[string]interface{}, ) *GetTotalRequests200Response`

NewGetTotalRequests200Response instantiates a new GetTotalRequests200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTotalRequests200ResponseWithDefaults

`func NewGetTotalRequests200ResponseWithDefaults() *GetTotalRequests200Response`

NewGetTotalRequests200ResponseWithDefaults instantiates a new GetTotalRequests200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetTotalRequests200Response) GetData() TotalRequest`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTotalRequests200Response) GetDataOk() (*TotalRequest, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTotalRequests200Response) SetData(v TotalRequest)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetTotalRequests200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTotalRequests200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTotalRequests200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


