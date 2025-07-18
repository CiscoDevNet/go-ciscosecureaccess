# GetRequestsByHourAndCategory200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CategoryByHour**](CategoryByHour.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetRequestsByHourAndCategory200Response

`func NewGetRequestsByHourAndCategory200Response(data []CategoryByHour, meta map[string]interface{}, ) *GetRequestsByHourAndCategory200Response`

NewGetRequestsByHourAndCategory200Response instantiates a new GetRequestsByHourAndCategory200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRequestsByHourAndCategory200ResponseWithDefaults

`func NewGetRequestsByHourAndCategory200ResponseWithDefaults() *GetRequestsByHourAndCategory200Response`

NewGetRequestsByHourAndCategory200ResponseWithDefaults instantiates a new GetRequestsByHourAndCategory200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetRequestsByHourAndCategory200Response) GetData() []CategoryByHour`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetRequestsByHourAndCategory200Response) GetDataOk() (*[]CategoryByHour, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetRequestsByHourAndCategory200Response) SetData(v []CategoryByHour)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetRequestsByHourAndCategory200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetRequestsByHourAndCategory200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetRequestsByHourAndCategory200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


