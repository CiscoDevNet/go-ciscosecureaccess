# GetApplications200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**ApplicationsWithCategories**](ApplicationsWithCategories.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetApplications200Response

`func NewGetApplications200Response(data ApplicationsWithCategories, meta map[string]interface{}, ) *GetApplications200Response`

NewGetApplications200Response instantiates a new GetApplications200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApplications200ResponseWithDefaults

`func NewGetApplications200ResponseWithDefaults() *GetApplications200Response`

NewGetApplications200ResponseWithDefaults instantiates a new GetApplications200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetApplications200Response) GetData() ApplicationsWithCategories`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetApplications200Response) GetDataOk() (*ApplicationsWithCategories, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetApplications200Response) SetData(v ApplicationsWithCategories)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetApplications200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetApplications200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetApplications200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


