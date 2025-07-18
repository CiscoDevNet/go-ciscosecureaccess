# DeploymentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**IdentityType**](IdentityType.md) |  | 
**Activecount** | **int64** | The count of the active identity type. | 
**Count** | **int64** | The total count of the identity type. | 

## Methods

### NewDeploymentStatus

`func NewDeploymentStatus(type_ IdentityType, activecount int64, count int64, ) *DeploymentStatus`

NewDeploymentStatus instantiates a new DeploymentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentStatusWithDefaults

`func NewDeploymentStatusWithDefaults() *DeploymentStatus`

NewDeploymentStatusWithDefaults instantiates a new DeploymentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DeploymentStatus) GetType() IdentityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeploymentStatus) GetTypeOk() (*IdentityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeploymentStatus) SetType(v IdentityType)`

SetType sets Type field to given value.


### GetActivecount

`func (o *DeploymentStatus) GetActivecount() int64`

GetActivecount returns the Activecount field if non-nil, zero value otherwise.

### GetActivecountOk

`func (o *DeploymentStatus) GetActivecountOk() (*int64, bool)`

GetActivecountOk returns a tuple with the Activecount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivecount

`func (o *DeploymentStatus) SetActivecount(v int64)`

SetActivecount sets Activecount field to given value.


### GetCount

`func (o *DeploymentStatus) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DeploymentStatus) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DeploymentStatus) SetCount(v int64)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


