# DestinationListObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The unique ID of the destination list. | 
**OrganizationId** | **int64** | The organization ID. | 
**Access** | **string** | Configure your access rules to block or allow certain destination lists. No support for the access type on destination lists. Valid values are: &#x60;allow&#x60;, &#x60;block&#x60;, &#x60;url_proxy&#x60;, &#x60;no_decrypt&#x60;, &#x60;warn&#x60;, or &#x60;none&#x60;. | 
**IsGlobal** | **bool** | Specifies whether the destination list is a global destination list. There is only one default &#x60;allow&#x60; destination list and one default &#x60;block&#x60; destination list for an organization. **Note:** No support for global destination lists. When you creat a destination list, set the &#x60;isGlobal&#x60; field to &#x60;false&#x60;. | 
**Name** | **string** | The name of the destination list. | 
**ThirdpartyCategoryId** | **int64** | The third-party category ID of the destination list. | 
**CreatedAt** | **int64** | The date and time when the destination list was created. | 
**ModifiedAt** | **int64** | The date and time when the destination list was modified. | 
**IsMspDefault** | **bool** | Specifies whether MSP is the default. | 
**MarkedForDeletion** | **bool** | Specifies whether the destination list is marked for deletion. | 
**BundleTypeId** | Pointer to [**BundleTypeId**](BundleTypeId.md) |  | [optional] 
**Meta** | Pointer to [**DestinationListObjectMeta**](DestinationListObjectMeta.md) |  | [optional] 

## Methods

### NewDestinationListObject

`func NewDestinationListObject(id int64, organizationId int64, access string, isGlobal bool, name string, thirdpartyCategoryId int64, createdAt int64, modifiedAt int64, isMspDefault bool, markedForDeletion bool, ) *DestinationListObject`

NewDestinationListObject instantiates a new DestinationListObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationListObjectWithDefaults

`func NewDestinationListObjectWithDefaults() *DestinationListObject`

NewDestinationListObjectWithDefaults instantiates a new DestinationListObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DestinationListObject) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DestinationListObject) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DestinationListObject) SetId(v int64)`

SetId sets Id field to given value.


### GetOrganizationId

`func (o *DestinationListObject) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *DestinationListObject) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *DestinationListObject) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.


### GetAccess

`func (o *DestinationListObject) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *DestinationListObject) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *DestinationListObject) SetAccess(v string)`

SetAccess sets Access field to given value.


### GetIsGlobal

`func (o *DestinationListObject) GetIsGlobal() bool`

GetIsGlobal returns the IsGlobal field if non-nil, zero value otherwise.

### GetIsGlobalOk

`func (o *DestinationListObject) GetIsGlobalOk() (*bool, bool)`

GetIsGlobalOk returns a tuple with the IsGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobal

`func (o *DestinationListObject) SetIsGlobal(v bool)`

SetIsGlobal sets IsGlobal field to given value.


### GetName

`func (o *DestinationListObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DestinationListObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DestinationListObject) SetName(v string)`

SetName sets Name field to given value.


### GetThirdpartyCategoryId

`func (o *DestinationListObject) GetThirdpartyCategoryId() int64`

GetThirdpartyCategoryId returns the ThirdpartyCategoryId field if non-nil, zero value otherwise.

### GetThirdpartyCategoryIdOk

`func (o *DestinationListObject) GetThirdpartyCategoryIdOk() (*int64, bool)`

GetThirdpartyCategoryIdOk returns a tuple with the ThirdpartyCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdpartyCategoryId

`func (o *DestinationListObject) SetThirdpartyCategoryId(v int64)`

SetThirdpartyCategoryId sets ThirdpartyCategoryId field to given value.


### GetCreatedAt

`func (o *DestinationListObject) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DestinationListObject) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DestinationListObject) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *DestinationListObject) GetModifiedAt() int64`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *DestinationListObject) GetModifiedAtOk() (*int64, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *DestinationListObject) SetModifiedAt(v int64)`

SetModifiedAt sets ModifiedAt field to given value.


### GetIsMspDefault

`func (o *DestinationListObject) GetIsMspDefault() bool`

GetIsMspDefault returns the IsMspDefault field if non-nil, zero value otherwise.

### GetIsMspDefaultOk

`func (o *DestinationListObject) GetIsMspDefaultOk() (*bool, bool)`

GetIsMspDefaultOk returns a tuple with the IsMspDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMspDefault

`func (o *DestinationListObject) SetIsMspDefault(v bool)`

SetIsMspDefault sets IsMspDefault field to given value.


### GetMarkedForDeletion

`func (o *DestinationListObject) GetMarkedForDeletion() bool`

GetMarkedForDeletion returns the MarkedForDeletion field if non-nil, zero value otherwise.

### GetMarkedForDeletionOk

`func (o *DestinationListObject) GetMarkedForDeletionOk() (*bool, bool)`

GetMarkedForDeletionOk returns a tuple with the MarkedForDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedForDeletion

`func (o *DestinationListObject) SetMarkedForDeletion(v bool)`

SetMarkedForDeletion sets MarkedForDeletion field to given value.


### GetBundleTypeId

`func (o *DestinationListObject) GetBundleTypeId() BundleTypeId`

GetBundleTypeId returns the BundleTypeId field if non-nil, zero value otherwise.

### GetBundleTypeIdOk

`func (o *DestinationListObject) GetBundleTypeIdOk() (*BundleTypeId, bool)`

GetBundleTypeIdOk returns a tuple with the BundleTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleTypeId

`func (o *DestinationListObject) SetBundleTypeId(v BundleTypeId)`

SetBundleTypeId sets BundleTypeId field to given value.

### HasBundleTypeId

`func (o *DestinationListObject) HasBundleTypeId() bool`

HasBundleTypeId returns a boolean if a field has been set.

### GetMeta

`func (o *DestinationListObject) GetMeta() DestinationListObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DestinationListObject) GetMetaOk() (*DestinationListObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DestinationListObject) SetMeta(v DestinationListObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DestinationListObject) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


