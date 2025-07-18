# GetSummariesByRuleIntrusion200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]SignatureListSummary**](SignatureListSummary.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetSummariesByRuleIntrusion200Response

`func NewGetSummariesByRuleIntrusion200Response(data []SignatureListSummary, meta map[string]interface{}, ) *GetSummariesByRuleIntrusion200Response`

NewGetSummariesByRuleIntrusion200Response instantiates a new GetSummariesByRuleIntrusion200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSummariesByRuleIntrusion200ResponseWithDefaults

`func NewGetSummariesByRuleIntrusion200ResponseWithDefaults() *GetSummariesByRuleIntrusion200Response`

NewGetSummariesByRuleIntrusion200ResponseWithDefaults instantiates a new GetSummariesByRuleIntrusion200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetSummariesByRuleIntrusion200Response) GetData() []SignatureListSummary`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetSummariesByRuleIntrusion200Response) GetDataOk() (*[]SignatureListSummary, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetSummariesByRuleIntrusion200Response) SetData(v []SignatureListSummary)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetSummariesByRuleIntrusion200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetSummariesByRuleIntrusion200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetSummariesByRuleIntrusion200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


