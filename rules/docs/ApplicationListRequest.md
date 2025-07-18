# ApplicationListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationListName** | **string** | The descriptive label for the application list. | 
**IsDefault** | **bool** | Specifies whether the application list is the default application list. | 
**ApplicationIds** | **[]int64** | The list of IDs for the applications. | 
**ApplicationCategoryIds** | Pointer to **[]int64** | The list of IDs for the application categories. | [optional] 

## Methods

### NewApplicationListRequest

`func NewApplicationListRequest(applicationListName string, isDefault bool, applicationIds []int64, ) *ApplicationListRequest`

NewApplicationListRequest instantiates a new ApplicationListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationListRequestWithDefaults

`func NewApplicationListRequestWithDefaults() *ApplicationListRequest`

NewApplicationListRequestWithDefaults instantiates a new ApplicationListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationListName

`func (o *ApplicationListRequest) GetApplicationListName() string`

GetApplicationListName returns the ApplicationListName field if non-nil, zero value otherwise.

### GetApplicationListNameOk

`func (o *ApplicationListRequest) GetApplicationListNameOk() (*string, bool)`

GetApplicationListNameOk returns a tuple with the ApplicationListName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationListName

`func (o *ApplicationListRequest) SetApplicationListName(v string)`

SetApplicationListName sets ApplicationListName field to given value.


### GetIsDefault

`func (o *ApplicationListRequest) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ApplicationListRequest) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ApplicationListRequest) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetApplicationIds

`func (o *ApplicationListRequest) GetApplicationIds() []int64`

GetApplicationIds returns the ApplicationIds field if non-nil, zero value otherwise.

### GetApplicationIdsOk

`func (o *ApplicationListRequest) GetApplicationIdsOk() (*[]int64, bool)`

GetApplicationIdsOk returns a tuple with the ApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIds

`func (o *ApplicationListRequest) SetApplicationIds(v []int64)`

SetApplicationIds sets ApplicationIds field to given value.


### GetApplicationCategoryIds

`func (o *ApplicationListRequest) GetApplicationCategoryIds() []int64`

GetApplicationCategoryIds returns the ApplicationCategoryIds field if non-nil, zero value otherwise.

### GetApplicationCategoryIdsOk

`func (o *ApplicationListRequest) GetApplicationCategoryIdsOk() (*[]int64, bool)`

GetApplicationCategoryIdsOk returns a tuple with the ApplicationCategoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCategoryIds

`func (o *ApplicationListRequest) SetApplicationCategoryIds(v []int64)`

SetApplicationCategoryIds sets ApplicationCategoryIds field to given value.

### HasApplicationCategoryIds

`func (o *ApplicationListRequest) HasApplicationCategoryIds() bool`

HasApplicationCategoryIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


