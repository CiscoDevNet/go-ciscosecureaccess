# ApplicationLists

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | The number of application lists in the organization. | [optional] 
**Result** | Pointer to [**[]ApplicationListsResultInner**](ApplicationListsResultInner.md) | The application lists in the organization. | [optional] 

## Methods

### NewApplicationLists

`func NewApplicationLists() *ApplicationLists`

NewApplicationLists instantiates a new ApplicationLists object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationListsWithDefaults

`func NewApplicationListsWithDefaults() *ApplicationLists`

NewApplicationListsWithDefaults instantiates a new ApplicationLists object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ApplicationLists) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ApplicationLists) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ApplicationLists) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ApplicationLists) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResult

`func (o *ApplicationLists) GetResult() []ApplicationListsResultInner`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ApplicationLists) GetResultOk() (*[]ApplicationListsResultInner, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ApplicationLists) SetResult(v []ApplicationListsResultInner)`

SetResult sets Result field to given value.

### HasResult

`func (o *ApplicationLists) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


