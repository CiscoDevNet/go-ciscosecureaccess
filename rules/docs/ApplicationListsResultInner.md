# ApplicationListsResultInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationListId** | Pointer to **int64** | The ID of the application list. | [optional] 
**ApplicationListName** | Pointer to **string** | The descriptive label for the application list. | [optional] 
**IsDefault** | Pointer to **bool** | Specifies whether the application list is the default application list. | [optional] 
**CreatedAt** | Pointer to **string** | The date and time that the application list was created. | [optional] [readonly] 
**ModifiedAt** | Pointer to **string** | The date and time that the application list was modified. | [optional] [readonly] 

## Methods

### NewApplicationListsResultInner

`func NewApplicationListsResultInner() *ApplicationListsResultInner`

NewApplicationListsResultInner instantiates a new ApplicationListsResultInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationListsResultInnerWithDefaults

`func NewApplicationListsResultInnerWithDefaults() *ApplicationListsResultInner`

NewApplicationListsResultInnerWithDefaults instantiates a new ApplicationListsResultInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationListId

`func (o *ApplicationListsResultInner) GetApplicationListId() int64`

GetApplicationListId returns the ApplicationListId field if non-nil, zero value otherwise.

### GetApplicationListIdOk

`func (o *ApplicationListsResultInner) GetApplicationListIdOk() (*int64, bool)`

GetApplicationListIdOk returns a tuple with the ApplicationListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationListId

`func (o *ApplicationListsResultInner) SetApplicationListId(v int64)`

SetApplicationListId sets ApplicationListId field to given value.

### HasApplicationListId

`func (o *ApplicationListsResultInner) HasApplicationListId() bool`

HasApplicationListId returns a boolean if a field has been set.

### GetApplicationListName

`func (o *ApplicationListsResultInner) GetApplicationListName() string`

GetApplicationListName returns the ApplicationListName field if non-nil, zero value otherwise.

### GetApplicationListNameOk

`func (o *ApplicationListsResultInner) GetApplicationListNameOk() (*string, bool)`

GetApplicationListNameOk returns a tuple with the ApplicationListName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationListName

`func (o *ApplicationListsResultInner) SetApplicationListName(v string)`

SetApplicationListName sets ApplicationListName field to given value.

### HasApplicationListName

`func (o *ApplicationListsResultInner) HasApplicationListName() bool`

HasApplicationListName returns a boolean if a field has been set.

### GetIsDefault

`func (o *ApplicationListsResultInner) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ApplicationListsResultInner) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ApplicationListsResultInner) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ApplicationListsResultInner) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApplicationListsResultInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApplicationListsResultInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApplicationListsResultInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApplicationListsResultInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *ApplicationListsResultInner) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ApplicationListsResultInner) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ApplicationListsResultInner) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *ApplicationListsResultInner) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


