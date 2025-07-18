# GetActivities200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GetActivities200ResponseDataInner**](GetActivities200ResponseDataInner.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetActivities200Response

`func NewGetActivities200Response(data []GetActivities200ResponseDataInner, meta map[string]interface{}, ) *GetActivities200Response`

NewGetActivities200Response instantiates a new GetActivities200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetActivities200ResponseWithDefaults

`func NewGetActivities200ResponseWithDefaults() *GetActivities200Response`

NewGetActivities200ResponseWithDefaults instantiates a new GetActivities200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetActivities200Response) GetData() []GetActivities200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetActivities200Response) GetDataOk() (*[]GetActivities200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetActivities200Response) SetData(v []GetActivities200ResponseDataInner)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetActivities200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetActivities200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetActivities200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


