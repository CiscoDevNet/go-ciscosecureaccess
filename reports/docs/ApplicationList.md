# ApplicationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]ApplicationInList**](ApplicationInList.md) | The list of application properties. | [optional] 

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

### GetItems

`func (o *ApplicationList) GetItems() []ApplicationInList`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ApplicationList) GetItemsOk() (*[]ApplicationInList, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ApplicationList) SetItems(v []ApplicationInList)`

SetItems sets Items field to given value.

### HasItems

`func (o *ApplicationList) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


