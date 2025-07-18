# GetBandwidthByHour200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BandwidthbyHour**](BandwidthbyHour.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetBandwidthByHour200Response

`func NewGetBandwidthByHour200Response(data []BandwidthbyHour, meta map[string]interface{}, ) *GetBandwidthByHour200Response`

NewGetBandwidthByHour200Response instantiates a new GetBandwidthByHour200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBandwidthByHour200ResponseWithDefaults

`func NewGetBandwidthByHour200ResponseWithDefaults() *GetBandwidthByHour200Response`

NewGetBandwidthByHour200ResponseWithDefaults instantiates a new GetBandwidthByHour200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetBandwidthByHour200Response) GetData() []BandwidthbyHour`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetBandwidthByHour200Response) GetDataOk() (*[]BandwidthbyHour, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetBandwidthByHour200Response) SetData(v []BandwidthbyHour)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetBandwidthByHour200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetBandwidthByHour200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetBandwidthByHour200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


