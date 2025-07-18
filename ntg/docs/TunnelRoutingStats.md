# TunnelRoutingStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientRouteStats** | Pointer to [**TunnelRoutingStatsClientRouteStats**](TunnelRoutingStatsClientRouteStats.md) |  | [optional] 
**CloudRouteStats** | Pointer to [**TunnelRoutingStatsCloudRouteStats**](TunnelRoutingStatsCloudRouteStats.md) |  | [optional] 

## Methods

### NewTunnelRoutingStats

`func NewTunnelRoutingStats() *TunnelRoutingStats`

NewTunnelRoutingStats instantiates a new TunnelRoutingStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelRoutingStatsWithDefaults

`func NewTunnelRoutingStatsWithDefaults() *TunnelRoutingStats`

NewTunnelRoutingStatsWithDefaults instantiates a new TunnelRoutingStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientRouteStats

`func (o *TunnelRoutingStats) GetClientRouteStats() TunnelRoutingStatsClientRouteStats`

GetClientRouteStats returns the ClientRouteStats field if non-nil, zero value otherwise.

### GetClientRouteStatsOk

`func (o *TunnelRoutingStats) GetClientRouteStatsOk() (*TunnelRoutingStatsClientRouteStats, bool)`

GetClientRouteStatsOk returns a tuple with the ClientRouteStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRouteStats

`func (o *TunnelRoutingStats) SetClientRouteStats(v TunnelRoutingStatsClientRouteStats)`

SetClientRouteStats sets ClientRouteStats field to given value.

### HasClientRouteStats

`func (o *TunnelRoutingStats) HasClientRouteStats() bool`

HasClientRouteStats returns a boolean if a field has been set.

### GetCloudRouteStats

`func (o *TunnelRoutingStats) GetCloudRouteStats() TunnelRoutingStatsCloudRouteStats`

GetCloudRouteStats returns the CloudRouteStats field if non-nil, zero value otherwise.

### GetCloudRouteStatsOk

`func (o *TunnelRoutingStats) GetCloudRouteStatsOk() (*TunnelRoutingStatsCloudRouteStats, bool)`

GetCloudRouteStatsOk returns a tuple with the CloudRouteStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudRouteStats

`func (o *TunnelRoutingStats) SetCloudRouteStats(v TunnelRoutingStatsCloudRouteStats)`

SetCloudRouteStats sets CloudRouteStats field to given value.

### HasCloudRouteStats

`func (o *TunnelRoutingStats) HasCloudRouteStats() bool`

HasCloudRouteStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


