# PrivateResourceGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Private Resource Group. The Private Resource Group name must not have any special characters other than spaces and hyphens. | 
**Description** | Pointer to **string** | The description of the Private Resource Group. | [optional] 
**ResourceIds** | **[]int64** | The list of IDs for the private resources that belong in the group. The list may be empty. | 

## Methods

### NewPrivateResourceGroupRequest

`func NewPrivateResourceGroupRequest(name string, resourceIds []int64, ) *PrivateResourceGroupRequest`

NewPrivateResourceGroupRequest instantiates a new PrivateResourceGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateResourceGroupRequestWithDefaults

`func NewPrivateResourceGroupRequestWithDefaults() *PrivateResourceGroupRequest`

NewPrivateResourceGroupRequestWithDefaults instantiates a new PrivateResourceGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PrivateResourceGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateResourceGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateResourceGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PrivateResourceGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrivateResourceGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrivateResourceGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PrivateResourceGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResourceIds

`func (o *PrivateResourceGroupRequest) GetResourceIds() []int64`

GetResourceIds returns the ResourceIds field if non-nil, zero value otherwise.

### GetResourceIdsOk

`func (o *PrivateResourceGroupRequest) GetResourceIdsOk() (*[]int64, bool)`

GetResourceIdsOk returns a tuple with the ResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIds

`func (o *PrivateResourceGroupRequest) SetResourceIds(v []int64)`

SetResourceIds sets ResourceIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


