# DestinationListCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | [**Access**](Access.md) |  | 
**IsGlobal** | **bool** | Specifies whether the destination list is a global destination list. There is only one default &#x60;allow&#x60; destination list and one default &#x60;block&#x60; destination list for an organization. **Note:** No support for global destination lists. When you creat a destination list, set the &#x60;isGlobal&#x60; field to &#x60;false&#x60;. | 
**Name** | **string** | The name of the destination list. | 
**BundleTypeId** | Pointer to [**BundleTypeId**](BundleTypeId.md) |  | [optional] 
**Destinations** | Pointer to [**[]DestinationListCreateDestinationsInner**](DestinationListCreateDestinationsInner.md) | The list of destinations. | [optional] 

## Methods

### NewDestinationListCreate

`func NewDestinationListCreate(access Access, isGlobal bool, name string, ) *DestinationListCreate`

NewDestinationListCreate instantiates a new DestinationListCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationListCreateWithDefaults

`func NewDestinationListCreateWithDefaults() *DestinationListCreate`

NewDestinationListCreateWithDefaults instantiates a new DestinationListCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *DestinationListCreate) GetAccess() Access`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *DestinationListCreate) GetAccessOk() (*Access, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *DestinationListCreate) SetAccess(v Access)`

SetAccess sets Access field to given value.


### GetIsGlobal

`func (o *DestinationListCreate) GetIsGlobal() bool`

GetIsGlobal returns the IsGlobal field if non-nil, zero value otherwise.

### GetIsGlobalOk

`func (o *DestinationListCreate) GetIsGlobalOk() (*bool, bool)`

GetIsGlobalOk returns a tuple with the IsGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobal

`func (o *DestinationListCreate) SetIsGlobal(v bool)`

SetIsGlobal sets IsGlobal field to given value.


### GetName

`func (o *DestinationListCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DestinationListCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DestinationListCreate) SetName(v string)`

SetName sets Name field to given value.


### GetBundleTypeId

`func (o *DestinationListCreate) GetBundleTypeId() BundleTypeId`

GetBundleTypeId returns the BundleTypeId field if non-nil, zero value otherwise.

### GetBundleTypeIdOk

`func (o *DestinationListCreate) GetBundleTypeIdOk() (*BundleTypeId, bool)`

GetBundleTypeIdOk returns a tuple with the BundleTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleTypeId

`func (o *DestinationListCreate) SetBundleTypeId(v BundleTypeId)`

SetBundleTypeId sets BundleTypeId field to given value.

### HasBundleTypeId

`func (o *DestinationListCreate) HasBundleTypeId() bool`

HasBundleTypeId returns a boolean if a field has been set.

### GetDestinations

`func (o *DestinationListCreate) GetDestinations() []DestinationListCreateDestinationsInner`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *DestinationListCreate) GetDestinationsOk() (*[]DestinationListCreateDestinationsInner, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *DestinationListCreate) SetDestinations(v []DestinationListCreateDestinationsInner)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *DestinationListCreate) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


