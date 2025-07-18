# CategoryCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | [**Category**](Category.md) |  | 
**Requests** | **int64** | The total number of requests for the category. | 

## Methods

### NewCategoryCount

`func NewCategoryCount(category Category, requests int64, ) *CategoryCount`

NewCategoryCount instantiates a new CategoryCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryCountWithDefaults

`func NewCategoryCountWithDefaults() *CategoryCount`

NewCategoryCountWithDefaults instantiates a new CategoryCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *CategoryCount) GetCategory() Category`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CategoryCount) GetCategoryOk() (*Category, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CategoryCount) SetCategory(v Category)`

SetCategory sets Category field to given value.


### GetRequests

`func (o *CategoryCount) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *CategoryCount) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *CategoryCount) SetRequests(v int64)`

SetRequests sets Requests field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


