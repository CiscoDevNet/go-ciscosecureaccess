# ConnectorGroupList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ConnectorGroupResponse**](ConnectorGroupResponse.md) | The list of the Connector Groups. | [optional] 
**Offset** | Pointer to **int64** | The place to start reading in the collection. The offset starts at 0. If the &#x60;limit&#x60; is 10, the offset for the next response is 10. The default value is 0. | [optional] 
**Limit** | Pointer to **int64** | The maximum number of items in the response. | [optional] 
**Total** | Pointer to **int64** | The total number of items read from the collection. | [optional] 

## Methods

### NewConnectorGroupList

`func NewConnectorGroupList() *ConnectorGroupList`

NewConnectorGroupList instantiates a new ConnectorGroupList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorGroupListWithDefaults

`func NewConnectorGroupListWithDefaults() *ConnectorGroupList`

NewConnectorGroupListWithDefaults instantiates a new ConnectorGroupList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ConnectorGroupList) GetData() []ConnectorGroupResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConnectorGroupList) GetDataOk() (*[]ConnectorGroupResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConnectorGroupList) SetData(v []ConnectorGroupResponse)`

SetData sets Data field to given value.

### HasData

`func (o *ConnectorGroupList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetOffset

`func (o *ConnectorGroupList) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ConnectorGroupList) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ConnectorGroupList) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ConnectorGroupList) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *ConnectorGroupList) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ConnectorGroupList) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ConnectorGroupList) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ConnectorGroupList) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTotal

`func (o *ConnectorGroupList) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ConnectorGroupList) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ConnectorGroupList) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ConnectorGroupList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


