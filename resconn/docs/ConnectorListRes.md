# ConnectorListRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ConnectorResponse**](ConnectorResponse.md) | The list of Connectors. | [optional] 
**Offset** | Pointer to **int64** | The place to start reading in the collection. The offset starts at 0. If the &#x60;limit&#x60; is 10, the offset for the next response is 10. The default value is 0. | [optional] 
**Limit** | Pointer to **int64** | The maximum number of items in the response. | [optional] 
**Total** | Pointer to **int64** | The total number of items read from the collection. | [optional] 

## Methods

### NewConnectorListRes

`func NewConnectorListRes() *ConnectorListRes`

NewConnectorListRes instantiates a new ConnectorListRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorListResWithDefaults

`func NewConnectorListResWithDefaults() *ConnectorListRes`

NewConnectorListResWithDefaults instantiates a new ConnectorListRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ConnectorListRes) GetData() []ConnectorResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConnectorListRes) GetDataOk() (*[]ConnectorResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConnectorListRes) SetData(v []ConnectorResponse)`

SetData sets Data field to given value.

### HasData

`func (o *ConnectorListRes) HasData() bool`

HasData returns a boolean if a field has been set.

### GetOffset

`func (o *ConnectorListRes) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ConnectorListRes) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ConnectorListRes) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ConnectorListRes) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *ConnectorListRes) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ConnectorListRes) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ConnectorListRes) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ConnectorListRes) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTotal

`func (o *ConnectorListRes) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ConnectorListRes) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ConnectorListRes) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ConnectorListRes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


