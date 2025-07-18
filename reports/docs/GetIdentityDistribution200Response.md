# GetIdentityDistribution200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]IdentityDistribution**](IdentityDistribution.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetIdentityDistribution200Response

`func NewGetIdentityDistribution200Response(data []IdentityDistribution, meta map[string]interface{}, ) *GetIdentityDistribution200Response`

NewGetIdentityDistribution200Response instantiates a new GetIdentityDistribution200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIdentityDistribution200ResponseWithDefaults

`func NewGetIdentityDistribution200ResponseWithDefaults() *GetIdentityDistribution200Response`

NewGetIdentityDistribution200ResponseWithDefaults instantiates a new GetIdentityDistribution200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetIdentityDistribution200Response) GetData() []IdentityDistribution`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetIdentityDistribution200Response) GetDataOk() (*[]IdentityDistribution, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetIdentityDistribution200Response) SetData(v []IdentityDistribution)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetIdentityDistribution200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetIdentityDistribution200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetIdentityDistribution200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


