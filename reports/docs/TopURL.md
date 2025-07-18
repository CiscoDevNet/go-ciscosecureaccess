# TopURL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** | The total number of API requests. | 
**Path** | **string** | The URL path. | 
**Categories** | [**[]Category**](Category.md) | The list of categories. | 
**Rank** | **int64** | The numeric rank of the top URL. | 

## Methods

### NewTopURL

`func NewTopURL(count int64, path string, categories []Category, rank int64, ) *TopURL`

NewTopURL instantiates a new TopURL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopURLWithDefaults

`func NewTopURLWithDefaults() *TopURL`

NewTopURLWithDefaults instantiates a new TopURL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *TopURL) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TopURL) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TopURL) SetCount(v int64)`

SetCount sets Count field to given value.


### GetPath

`func (o *TopURL) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *TopURL) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *TopURL) SetPath(v string)`

SetPath sets Path field to given value.


### GetCategories

`func (o *TopURL) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *TopURL) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *TopURL) SetCategories(v []Category)`

SetCategories sets Categories field to given value.


### GetRank

`func (o *TopURL) GetRank() int64`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *TopURL) GetRankOk() (*int64, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *TopURL) SetRank(v int64)`

SetRank sets Rank field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


