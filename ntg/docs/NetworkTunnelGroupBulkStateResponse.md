# NetworkTunnelGroupBulkStateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | Pointer to **int64** | An integer that represents the place to start reading in the collection. | [optional] 
**Limit** | Pointer to **int64** | The number of items returned in the response. | [optional] 
**Total** | Pointer to **int64** | The total number of items read from the collection. | [optional] 
**Data** | Pointer to [**[]NetworkTunnelGroupStateResponse**](NetworkTunnelGroupStateResponse.md) | The list of the Network Tunnel Groups in the organization. | [optional] 

## Methods

### NewNetworkTunnelGroupBulkStateResponse

`func NewNetworkTunnelGroupBulkStateResponse() *NetworkTunnelGroupBulkStateResponse`

NewNetworkTunnelGroupBulkStateResponse instantiates a new NetworkTunnelGroupBulkStateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkTunnelGroupBulkStateResponseWithDefaults

`func NewNetworkTunnelGroupBulkStateResponseWithDefaults() *NetworkTunnelGroupBulkStateResponse`

NewNetworkTunnelGroupBulkStateResponseWithDefaults instantiates a new NetworkTunnelGroupBulkStateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *NetworkTunnelGroupBulkStateResponse) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *NetworkTunnelGroupBulkStateResponse) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *NetworkTunnelGroupBulkStateResponse) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *NetworkTunnelGroupBulkStateResponse) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *NetworkTunnelGroupBulkStateResponse) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *NetworkTunnelGroupBulkStateResponse) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *NetworkTunnelGroupBulkStateResponse) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *NetworkTunnelGroupBulkStateResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTotal

`func (o *NetworkTunnelGroupBulkStateResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *NetworkTunnelGroupBulkStateResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *NetworkTunnelGroupBulkStateResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *NetworkTunnelGroupBulkStateResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetData

`func (o *NetworkTunnelGroupBulkStateResponse) GetData() []NetworkTunnelGroupStateResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NetworkTunnelGroupBulkStateResponse) GetDataOk() (*[]NetworkTunnelGroupStateResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NetworkTunnelGroupBulkStateResponse) SetData(v []NetworkTunnelGroupStateResponse)`

SetData sets Data field to given value.

### HasData

`func (o *NetworkTunnelGroupBulkStateResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


