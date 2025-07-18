# TopResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to [**TopResourceApplication**](TopResourceApplication.md) |  | [optional] 
**Count** | Pointer to **int64** | The number of times a source connected to the resource. | [optional] 
**Rank** | Pointer to **int64** | The score assigned to the app. | [optional] 

## Methods

### NewTopResource

`func NewTopResource() *TopResource`

NewTopResource instantiates a new TopResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopResourceWithDefaults

`func NewTopResourceWithDefaults() *TopResource`

NewTopResourceWithDefaults instantiates a new TopResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *TopResource) GetApplication() TopResourceApplication`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *TopResource) GetApplicationOk() (*TopResourceApplication, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *TopResource) SetApplication(v TopResourceApplication)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *TopResource) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetCount

`func (o *TopResource) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TopResource) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TopResource) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *TopResource) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetRank

`func (o *TopResource) GetRank() int64`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *TopResource) GetRankOk() (*int64, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *TopResource) SetRank(v int64)`

SetRank sets Rank field to given value.

### HasRank

`func (o *TopResource) HasRank() bool`

HasRank returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


