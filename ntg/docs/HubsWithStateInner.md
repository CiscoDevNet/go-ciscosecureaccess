# HubsWithStateInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the Hub. | [optional] [readonly] 
**IsPrimary** | Pointer to **bool** | Specifies whether the Hub is a primary data center. | [optional] [readonly] 
**Datacenter** | Pointer to [**HubsWithStateInnerDatacenter**](HubsWithStateInnerDatacenter.md) |  | [optional] 
**Status** | Pointer to [**HubState**](HubState.md) |  | [optional] 
**TunnelsStatus** | Pointer to [**[]TunnelState**](TunnelState.md) | The list of the states for the Network Tunnels. The maximum number of items in the list of tunnel states is 10. | [optional] 

## Methods

### NewHubsWithStateInner

`func NewHubsWithStateInner() *HubsWithStateInner`

NewHubsWithStateInner instantiates a new HubsWithStateInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHubsWithStateInnerWithDefaults

`func NewHubsWithStateInnerWithDefaults() *HubsWithStateInner`

NewHubsWithStateInnerWithDefaults instantiates a new HubsWithStateInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HubsWithStateInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HubsWithStateInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HubsWithStateInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *HubsWithStateInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsPrimary

`func (o *HubsWithStateInner) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *HubsWithStateInner) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *HubsWithStateInner) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *HubsWithStateInner) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetDatacenter

`func (o *HubsWithStateInner) GetDatacenter() HubsWithStateInnerDatacenter`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *HubsWithStateInner) GetDatacenterOk() (*HubsWithStateInnerDatacenter, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *HubsWithStateInner) SetDatacenter(v HubsWithStateInnerDatacenter)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *HubsWithStateInner) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetStatus

`func (o *HubsWithStateInner) GetStatus() HubState`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HubsWithStateInner) GetStatusOk() (*HubState, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HubsWithStateInner) SetStatus(v HubState)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HubsWithStateInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTunnelsStatus

`func (o *HubsWithStateInner) GetTunnelsStatus() []TunnelState`

GetTunnelsStatus returns the TunnelsStatus field if non-nil, zero value otherwise.

### GetTunnelsStatusOk

`func (o *HubsWithStateInner) GetTunnelsStatusOk() (*[]TunnelState, bool)`

GetTunnelsStatusOk returns a tuple with the TunnelsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelsStatus

`func (o *HubsWithStateInner) SetTunnelsStatus(v []TunnelState)`

SetTunnelsStatus sets TunnelsStatus field to given value.

### HasTunnelsStatus

`func (o *HubsWithStateInner) HasTunnelsStatus() bool`

HasTunnelsStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


