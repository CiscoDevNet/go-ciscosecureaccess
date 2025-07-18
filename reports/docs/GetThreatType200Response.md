# GetThreatType200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**ThreatType**](ThreatType.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetThreatType200Response

`func NewGetThreatType200Response(data ThreatType, meta map[string]interface{}, ) *GetThreatType200Response`

NewGetThreatType200Response instantiates a new GetThreatType200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetThreatType200ResponseWithDefaults

`func NewGetThreatType200ResponseWithDefaults() *GetThreatType200Response`

NewGetThreatType200ResponseWithDefaults instantiates a new GetThreatType200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetThreatType200Response) GetData() ThreatType`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetThreatType200Response) GetDataOk() (*ThreatType, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetThreatType200Response) SetData(v ThreatType)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetThreatType200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetThreatType200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetThreatType200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


