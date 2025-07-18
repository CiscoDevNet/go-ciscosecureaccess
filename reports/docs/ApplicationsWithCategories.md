# ApplicationsWithCategories

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to [**[]Application**](Application.md) | The list of applications. | [optional] 
**Categories** | Pointer to [**[]ApplicationCategories**](ApplicationCategories.md) | The list of application categories. | [optional] 

## Methods

### NewApplicationsWithCategories

`func NewApplicationsWithCategories() *ApplicationsWithCategories`

NewApplicationsWithCategories instantiates a new ApplicationsWithCategories object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationsWithCategoriesWithDefaults

`func NewApplicationsWithCategoriesWithDefaults() *ApplicationsWithCategories`

NewApplicationsWithCategoriesWithDefaults instantiates a new ApplicationsWithCategories object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *ApplicationsWithCategories) GetApplications() []Application`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *ApplicationsWithCategories) GetApplicationsOk() (*[]Application, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *ApplicationsWithCategories) SetApplications(v []Application)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *ApplicationsWithCategories) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetCategories

`func (o *ApplicationsWithCategories) GetCategories() []ApplicationCategories`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ApplicationsWithCategories) GetCategoriesOk() (*[]ApplicationCategories, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ApplicationsWithCategories) SetCategories(v []ApplicationCategories)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ApplicationsWithCategories) HasCategories() bool`

HasCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


