# GetIdentityDistributionType200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]IdentityDistribution**](IdentityDistribution.md) | The list of identity distributions. | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetIdentityDistributionType200Response

`func NewGetIdentityDistributionType200Response(data []IdentityDistribution, meta map[string]interface{}, ) *GetIdentityDistributionType200Response`

NewGetIdentityDistributionType200Response instantiates a new GetIdentityDistributionType200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIdentityDistributionType200ResponseWithDefaults

`func NewGetIdentityDistributionType200ResponseWithDefaults() *GetIdentityDistributionType200Response`

NewGetIdentityDistributionType200ResponseWithDefaults instantiates a new GetIdentityDistributionType200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetIdentityDistributionType200Response) GetData() []IdentityDistribution`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetIdentityDistributionType200Response) GetDataOk() (*[]IdentityDistribution, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetIdentityDistributionType200Response) SetData(v []IdentityDistribution)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetIdentityDistributionType200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetIdentityDistributionType200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetIdentityDistributionType200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


