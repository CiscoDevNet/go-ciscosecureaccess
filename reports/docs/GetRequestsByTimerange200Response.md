# GetRequestsByTimerange200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]RequestsbyHour**](RequestsbyHour.md) | The information about the provider&#39;s requests within the timerange. | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetRequestsByTimerange200Response

`func NewGetRequestsByTimerange200Response(data []RequestsbyHour, meta map[string]interface{}, ) *GetRequestsByTimerange200Response`

NewGetRequestsByTimerange200Response instantiates a new GetRequestsByTimerange200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRequestsByTimerange200ResponseWithDefaults

`func NewGetRequestsByTimerange200ResponseWithDefaults() *GetRequestsByTimerange200Response`

NewGetRequestsByTimerange200ResponseWithDefaults instantiates a new GetRequestsByTimerange200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetRequestsByTimerange200Response) GetData() []RequestsbyHour`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetRequestsByTimerange200Response) GetDataOk() (*[]RequestsbyHour, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetRequestsByTimerange200Response) SetData(v []RequestsbyHour)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetRequestsByTimerange200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetRequestsByTimerange200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetRequestsByTimerange200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


