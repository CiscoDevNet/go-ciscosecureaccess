# PrivateResourceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]PrivateResourceResponse**](PrivateResourceResponse.md) |  | [optional] 
**Offset** | Pointer to **int64** | The index where the server starts to read the collection. | [optional] 
**Limit** | Pointer to **int64** | The number of items returned on a page in the response. | [optional] 
**Total** | Pointer to **int64** | The total count of the Private Resources read in the collection. | [optional] 

## Methods

### NewPrivateResourceList

`func NewPrivateResourceList() *PrivateResourceList`

NewPrivateResourceList instantiates a new PrivateResourceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateResourceListWithDefaults

`func NewPrivateResourceListWithDefaults() *PrivateResourceList`

NewPrivateResourceListWithDefaults instantiates a new PrivateResourceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *PrivateResourceList) GetItems() []PrivateResourceResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PrivateResourceList) GetItemsOk() (*[]PrivateResourceResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PrivateResourceList) SetItems(v []PrivateResourceResponse)`

SetItems sets Items field to given value.

### HasItems

`func (o *PrivateResourceList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOffset

`func (o *PrivateResourceList) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PrivateResourceList) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PrivateResourceList) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *PrivateResourceList) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *PrivateResourceList) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PrivateResourceList) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PrivateResourceList) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PrivateResourceList) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTotal

`func (o *PrivateResourceList) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PrivateResourceList) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PrivateResourceList) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PrivateResourceList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


