# GetActivityDecryption200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ActivityDecryption**](ActivityDecryption.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetActivityDecryption200Response

`func NewGetActivityDecryption200Response(data []ActivityDecryption, meta map[string]interface{}, ) *GetActivityDecryption200Response`

NewGetActivityDecryption200Response instantiates a new GetActivityDecryption200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetActivityDecryption200ResponseWithDefaults

`func NewGetActivityDecryption200ResponseWithDefaults() *GetActivityDecryption200Response`

NewGetActivityDecryption200ResponseWithDefaults instantiates a new GetActivityDecryption200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetActivityDecryption200Response) GetData() []ActivityDecryption`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetActivityDecryption200Response) GetDataOk() (*[]ActivityDecryption, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetActivityDecryption200Response) SetData(v []ActivityDecryption)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetActivityDecryption200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetActivityDecryption200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetActivityDecryption200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


