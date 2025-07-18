# GetRequestsByHour200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]RequestsbyHour**](RequestsbyHour.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetRequestsByHour200Response

`func NewGetRequestsByHour200Response(data []RequestsbyHour, meta map[string]interface{}, ) *GetRequestsByHour200Response`

NewGetRequestsByHour200Response instantiates a new GetRequestsByHour200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRequestsByHour200ResponseWithDefaults

`func NewGetRequestsByHour200ResponseWithDefaults() *GetRequestsByHour200Response`

NewGetRequestsByHour200ResponseWithDefaults instantiates a new GetRequestsByHour200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetRequestsByHour200Response) GetData() []RequestsbyHour`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetRequestsByHour200Response) GetDataOk() (*[]RequestsbyHour, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetRequestsByHour200Response) SetData(v []RequestsbyHour)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetRequestsByHour200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetRequestsByHour200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetRequestsByHour200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


