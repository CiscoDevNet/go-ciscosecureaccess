# NetworkTunnelGroupsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]NetworkTunnelGroupListResponse**](NetworkTunnelGroupListResponse.md) | The list of Network Tunnel Groups. | [optional] 
**Offset** | Pointer to **int64** | An integer that represents the place to start reading in the collection. | [optional] 
**Limit** | Pointer to **int64** | The number of records returned on a page. | [optional] 
**Total** | Pointer to **int64** | The total number of records returned. | [optional] 

## Methods

### NewNetworkTunnelGroupsList

`func NewNetworkTunnelGroupsList() *NetworkTunnelGroupsList`

NewNetworkTunnelGroupsList instantiates a new NetworkTunnelGroupsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkTunnelGroupsListWithDefaults

`func NewNetworkTunnelGroupsListWithDefaults() *NetworkTunnelGroupsList`

NewNetworkTunnelGroupsListWithDefaults instantiates a new NetworkTunnelGroupsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *NetworkTunnelGroupsList) GetData() []NetworkTunnelGroupListResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NetworkTunnelGroupsList) GetDataOk() (*[]NetworkTunnelGroupListResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NetworkTunnelGroupsList) SetData(v []NetworkTunnelGroupListResponse)`

SetData sets Data field to given value.

### HasData

`func (o *NetworkTunnelGroupsList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetOffset

`func (o *NetworkTunnelGroupsList) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *NetworkTunnelGroupsList) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *NetworkTunnelGroupsList) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *NetworkTunnelGroupsList) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *NetworkTunnelGroupsList) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *NetworkTunnelGroupsList) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *NetworkTunnelGroupsList) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *NetworkTunnelGroupsList) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTotal

`func (o *NetworkTunnelGroupsList) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *NetworkTunnelGroupsList) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *NetworkTunnelGroupsList) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *NetworkTunnelGroupsList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


