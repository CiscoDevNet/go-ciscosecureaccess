# GetDeploymentStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DeploymentStatus**](DeploymentStatus.md) |  | 
**Meta** | **map[string]interface{}** | The properties of the metadata. | 

## Methods

### NewGetDeploymentStatus200Response

`func NewGetDeploymentStatus200Response(data []DeploymentStatus, meta map[string]interface{}, ) *GetDeploymentStatus200Response`

NewGetDeploymentStatus200Response instantiates a new GetDeploymentStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeploymentStatus200ResponseWithDefaults

`func NewGetDeploymentStatus200ResponseWithDefaults() *GetDeploymentStatus200Response`

NewGetDeploymentStatus200ResponseWithDefaults instantiates a new GetDeploymentStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetDeploymentStatus200Response) GetData() []DeploymentStatus`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetDeploymentStatus200Response) GetDataOk() (*[]DeploymentStatus, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetDeploymentStatus200Response) SetData(v []DeploymentStatus)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetDeploymentStatus200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDeploymentStatus200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDeploymentStatus200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


