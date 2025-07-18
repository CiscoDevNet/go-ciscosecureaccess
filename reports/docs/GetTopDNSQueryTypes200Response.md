# GetTopDNSQueryTypes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TopDnsQueryType**](TopDnsQueryType.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetTopDNSQueryTypes200Response

`func NewGetTopDNSQueryTypes200Response(data []TopDnsQueryType, meta map[string]interface{}, ) *GetTopDNSQueryTypes200Response`

NewGetTopDNSQueryTypes200Response instantiates a new GetTopDNSQueryTypes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTopDNSQueryTypes200ResponseWithDefaults

`func NewGetTopDNSQueryTypes200ResponseWithDefaults() *GetTopDNSQueryTypes200Response`

NewGetTopDNSQueryTypes200ResponseWithDefaults instantiates a new GetTopDNSQueryTypes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetTopDNSQueryTypes200Response) GetData() []TopDnsQueryType`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTopDNSQueryTypes200Response) GetDataOk() (*[]TopDnsQueryType, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTopDNSQueryTypes200Response) SetData(v []TopDnsQueryType)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetTopDNSQueryTypes200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTopDNSQueryTypes200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTopDNSQueryTypes200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


