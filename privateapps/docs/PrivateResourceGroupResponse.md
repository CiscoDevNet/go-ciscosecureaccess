# PrivateResourceGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceGroupId** | Pointer to **int64** | The ID of the Private Resource Group. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the Private Resource Group. The Private Resource Group name must not have any special characters other than spaces and hyphens. | [optional] 
**Description** | Pointer to **string** | The description of the Private Resource Group. | [optional] 
**ResourceIds** | Pointer to **[]int64** | The list of IDs for the private resources that belong in the group. The list may be empty. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when the Private Resource Group was created. | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** | The date and time when the Private Resource Group was updated. | [optional] [readonly] 
**CreatedBy** | Pointer to **string** | The ID of the user who created the Private Resource Group. | [optional] 
**ModifiedBy** | Pointer to **string** | The ID of the user who last modified the Private Resource Group. | [optional] 

## Methods

### NewPrivateResourceGroupResponse

`func NewPrivateResourceGroupResponse() *PrivateResourceGroupResponse`

NewPrivateResourceGroupResponse instantiates a new PrivateResourceGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateResourceGroupResponseWithDefaults

`func NewPrivateResourceGroupResponseWithDefaults() *PrivateResourceGroupResponse`

NewPrivateResourceGroupResponseWithDefaults instantiates a new PrivateResourceGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceGroupId

`func (o *PrivateResourceGroupResponse) GetResourceGroupId() int64`

GetResourceGroupId returns the ResourceGroupId field if non-nil, zero value otherwise.

### GetResourceGroupIdOk

`func (o *PrivateResourceGroupResponse) GetResourceGroupIdOk() (*int64, bool)`

GetResourceGroupIdOk returns a tuple with the ResourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupId

`func (o *PrivateResourceGroupResponse) SetResourceGroupId(v int64)`

SetResourceGroupId sets ResourceGroupId field to given value.

### HasResourceGroupId

`func (o *PrivateResourceGroupResponse) HasResourceGroupId() bool`

HasResourceGroupId returns a boolean if a field has been set.

### GetName

`func (o *PrivateResourceGroupResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateResourceGroupResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateResourceGroupResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrivateResourceGroupResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PrivateResourceGroupResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrivateResourceGroupResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrivateResourceGroupResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PrivateResourceGroupResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResourceIds

`func (o *PrivateResourceGroupResponse) GetResourceIds() []int64`

GetResourceIds returns the ResourceIds field if non-nil, zero value otherwise.

### GetResourceIdsOk

`func (o *PrivateResourceGroupResponse) GetResourceIdsOk() (*[]int64, bool)`

GetResourceIdsOk returns a tuple with the ResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIds

`func (o *PrivateResourceGroupResponse) SetResourceIds(v []int64)`

SetResourceIds sets ResourceIds field to given value.

### HasResourceIds

`func (o *PrivateResourceGroupResponse) HasResourceIds() bool`

HasResourceIds returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PrivateResourceGroupResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PrivateResourceGroupResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PrivateResourceGroupResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PrivateResourceGroupResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *PrivateResourceGroupResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PrivateResourceGroupResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PrivateResourceGroupResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *PrivateResourceGroupResponse) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *PrivateResourceGroupResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PrivateResourceGroupResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PrivateResourceGroupResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PrivateResourceGroupResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetModifiedBy

`func (o *PrivateResourceGroupResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *PrivateResourceGroupResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *PrivateResourceGroupResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *PrivateResourceGroupResponse) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


