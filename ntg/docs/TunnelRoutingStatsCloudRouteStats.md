# TunnelRoutingStatsCloudRouteStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsClipped** | Pointer to **bool** | Indicates whether cloud route stats array has been truncated | [optional] 
**Stats** | Pointer to [**[]TunnelRoutingStatsCloudRouteStatsStatsInner**](TunnelRoutingStatsCloudRouteStatsStatsInner.md) | The list of the cloud routing statistics. | [optional] 

## Methods

### NewTunnelRoutingStatsCloudRouteStats

`func NewTunnelRoutingStatsCloudRouteStats() *TunnelRoutingStatsCloudRouteStats`

NewTunnelRoutingStatsCloudRouteStats instantiates a new TunnelRoutingStatsCloudRouteStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelRoutingStatsCloudRouteStatsWithDefaults

`func NewTunnelRoutingStatsCloudRouteStatsWithDefaults() *TunnelRoutingStatsCloudRouteStats`

NewTunnelRoutingStatsCloudRouteStatsWithDefaults instantiates a new TunnelRoutingStatsCloudRouteStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsClipped

`func (o *TunnelRoutingStatsCloudRouteStats) GetIsClipped() bool`

GetIsClipped returns the IsClipped field if non-nil, zero value otherwise.

### GetIsClippedOk

`func (o *TunnelRoutingStatsCloudRouteStats) GetIsClippedOk() (*bool, bool)`

GetIsClippedOk returns a tuple with the IsClipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClipped

`func (o *TunnelRoutingStatsCloudRouteStats) SetIsClipped(v bool)`

SetIsClipped sets IsClipped field to given value.

### HasIsClipped

`func (o *TunnelRoutingStatsCloudRouteStats) HasIsClipped() bool`

HasIsClipped returns a boolean if a field has been set.

### GetStats

`func (o *TunnelRoutingStatsCloudRouteStats) GetStats() []TunnelRoutingStatsCloudRouteStatsStatsInner`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *TunnelRoutingStatsCloudRouteStats) GetStatsOk() (*[]TunnelRoutingStatsCloudRouteStatsStatsInner, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *TunnelRoutingStatsCloudRouteStats) SetStats(v []TunnelRoutingStatsCloudRouteStatsStatsInner)`

SetStats sets Stats field to given value.

### HasStats

`func (o *TunnelRoutingStatsCloudRouteStats) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


