# ApplicationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationListName** | Pointer to **string** | The descriptive label for the application list. | [optional] 
**IsDefault** | Pointer to **bool** | Specifies whether the application list is the default application list. | [optional] 
**ApplicationIds** | Pointer to **[]int64** | The list of IDs for the applications. | [optional] 
**ApplicationCategoryIds** | Pointer to **[]int64** | The list of IDs for the application categories. | [optional] 
**CreatedAt** | Pointer to **string** | The date and time that the application list was created. | [optional] [readonly] 
**ModifiedAt** | Pointer to **string** | The date and time that the application list was modified. | [optional] [readonly] 

## Methods

### NewApplicationList

`func NewApplicationList() *ApplicationList`

NewApplicationList instantiates a new ApplicationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationListWithDefaults

`func NewApplicationListWithDefaults() *ApplicationList`

NewApplicationListWithDefaults instantiates a new ApplicationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationListName

`func (o *ApplicationList) GetApplicationListName() string`

GetApplicationListName returns the ApplicationListName field if non-nil, zero value otherwise.

### GetApplicationListNameOk

`func (o *ApplicationList) GetApplicationListNameOk() (*string, bool)`

GetApplicationListNameOk returns a tuple with the ApplicationListName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationListName

`func (o *ApplicationList) SetApplicationListName(v string)`

SetApplicationListName sets ApplicationListName field to given value.

### HasApplicationListName

`func (o *ApplicationList) HasApplicationListName() bool`

HasApplicationListName returns a boolean if a field has been set.

### GetIsDefault

`func (o *ApplicationList) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ApplicationList) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ApplicationList) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ApplicationList) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetApplicationIds

`func (o *ApplicationList) GetApplicationIds() []int64`

GetApplicationIds returns the ApplicationIds field if non-nil, zero value otherwise.

### GetApplicationIdsOk

`func (o *ApplicationList) GetApplicationIdsOk() (*[]int64, bool)`

GetApplicationIdsOk returns a tuple with the ApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIds

`func (o *ApplicationList) SetApplicationIds(v []int64)`

SetApplicationIds sets ApplicationIds field to given value.

### HasApplicationIds

`func (o *ApplicationList) HasApplicationIds() bool`

HasApplicationIds returns a boolean if a field has been set.

### GetApplicationCategoryIds

`func (o *ApplicationList) GetApplicationCategoryIds() []int64`

GetApplicationCategoryIds returns the ApplicationCategoryIds field if non-nil, zero value otherwise.

### GetApplicationCategoryIdsOk

`func (o *ApplicationList) GetApplicationCategoryIdsOk() (*[]int64, bool)`

GetApplicationCategoryIdsOk returns a tuple with the ApplicationCategoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCategoryIds

`func (o *ApplicationList) SetApplicationCategoryIds(v []int64)`

SetApplicationCategoryIds sets ApplicationCategoryIds field to given value.

### HasApplicationCategoryIds

`func (o *ApplicationList) HasApplicationCategoryIds() bool`

HasApplicationCategoryIds returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApplicationList) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApplicationList) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApplicationList) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApplicationList) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *ApplicationList) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ApplicationList) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ApplicationList) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *ApplicationList) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


