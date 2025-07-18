# GetThreatName200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**ThreatName**](ThreatName.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetThreatName200Response

`func NewGetThreatName200Response(data ThreatName, meta map[string]interface{}, ) *GetThreatName200Response`

NewGetThreatName200Response instantiates a new GetThreatName200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetThreatName200ResponseWithDefaults

`func NewGetThreatName200ResponseWithDefaults() *GetThreatName200Response`

NewGetThreatName200ResponseWithDefaults instantiates a new GetThreatName200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetThreatName200Response) GetData() ThreatName`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetThreatName200Response) GetDataOk() (*ThreatName, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetThreatName200Response) SetData(v ThreatName)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetThreatName200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetThreatName200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetThreatName200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


