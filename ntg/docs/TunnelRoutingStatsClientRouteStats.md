# TunnelRoutingStatsClientRouteStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsClipped** | Pointer to **bool** | Indicates whether the routing statistics for the client route are truncated. | [optional] 
**Stats** | Pointer to [**[]TunnelRoutingStatsClientRouteStatsStatsInner**](TunnelRoutingStatsClientRouteStatsStatsInner.md) | The list of the client routing statistics. | [optional] 

## Methods

### NewTunnelRoutingStatsClientRouteStats

`func NewTunnelRoutingStatsClientRouteStats() *TunnelRoutingStatsClientRouteStats`

NewTunnelRoutingStatsClientRouteStats instantiates a new TunnelRoutingStatsClientRouteStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelRoutingStatsClientRouteStatsWithDefaults

`func NewTunnelRoutingStatsClientRouteStatsWithDefaults() *TunnelRoutingStatsClientRouteStats`

NewTunnelRoutingStatsClientRouteStatsWithDefaults instantiates a new TunnelRoutingStatsClientRouteStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsClipped

`func (o *TunnelRoutingStatsClientRouteStats) GetIsClipped() bool`

GetIsClipped returns the IsClipped field if non-nil, zero value otherwise.

### GetIsClippedOk

`func (o *TunnelRoutingStatsClientRouteStats) GetIsClippedOk() (*bool, bool)`

GetIsClippedOk returns a tuple with the IsClipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClipped

`func (o *TunnelRoutingStatsClientRouteStats) SetIsClipped(v bool)`

SetIsClipped sets IsClipped field to given value.

### HasIsClipped

`func (o *TunnelRoutingStatsClientRouteStats) HasIsClipped() bool`

HasIsClipped returns a boolean if a field has been set.

### GetStats

`func (o *TunnelRoutingStatsClientRouteStats) GetStats() []TunnelRoutingStatsClientRouteStatsStatsInner`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *TunnelRoutingStatsClientRouteStats) GetStatsOk() (*[]TunnelRoutingStatsClientRouteStatsStatsInner, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *TunnelRoutingStatsClientRouteStats) SetStats(v []TunnelRoutingStatsClientRouteStatsStatsInner)`

SetStats sets Stats field to given value.

### HasStats

`func (o *TunnelRoutingStatsClientRouteStats) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


