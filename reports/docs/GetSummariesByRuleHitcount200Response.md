# GetSummariesByRuleHitcount200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]RuleHitcountSummary**](RuleHitcountSummary.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetSummariesByRuleHitcount200Response

`func NewGetSummariesByRuleHitcount200Response(data []RuleHitcountSummary, meta map[string]interface{}, ) *GetSummariesByRuleHitcount200Response`

NewGetSummariesByRuleHitcount200Response instantiates a new GetSummariesByRuleHitcount200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSummariesByRuleHitcount200ResponseWithDefaults

`func NewGetSummariesByRuleHitcount200ResponseWithDefaults() *GetSummariesByRuleHitcount200Response`

NewGetSummariesByRuleHitcount200ResponseWithDefaults instantiates a new GetSummariesByRuleHitcount200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetSummariesByRuleHitcount200Response) GetData() []RuleHitcountSummary`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetSummariesByRuleHitcount200Response) GetDataOk() (*[]RuleHitcountSummary, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetSummariesByRuleHitcount200Response) SetData(v []RuleHitcountSummary)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetSummariesByRuleHitcount200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetSummariesByRuleHitcount200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetSummariesByRuleHitcount200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


