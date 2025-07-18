# ConnectorCountsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Announced** | Pointer to **int64** | Number of announced connectors in the organization. | [optional] 
**Connected** | Pointer to **int64** | The number of Connectors in the organization that are in the connected state. | [optional] 
**Disabled** | Pointer to **int64** | The number of Connectors in the organization that are in the disabled state. | [optional] 
**Disconnected** | Pointer to **int64** | The number of the Connectors in the organization that are in the disconnected state. | [optional] 
**Expired** | Pointer to **int64** | The number of Connectors in the organziation that are in the expired state. | [optional] 
**Reachable** | Pointer to **int64** | The number of Connectors in the organziation that are in the reachable state. | [optional] 
**Total** | Pointer to **int64** | The total number of Connectors in the organization. | [optional] 
**Upgrading** | Pointer to **int64** | The number of Resource Connector Groups in the organziation that are in the upgrading state. | [optional] 

## Methods

### NewConnectorCountsResponse

`func NewConnectorCountsResponse() *ConnectorCountsResponse`

NewConnectorCountsResponse instantiates a new ConnectorCountsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorCountsResponseWithDefaults

`func NewConnectorCountsResponseWithDefaults() *ConnectorCountsResponse`

NewConnectorCountsResponseWithDefaults instantiates a new ConnectorCountsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnounced

`func (o *ConnectorCountsResponse) GetAnnounced() int64`

GetAnnounced returns the Announced field if non-nil, zero value otherwise.

### GetAnnouncedOk

`func (o *ConnectorCountsResponse) GetAnnouncedOk() (*int64, bool)`

GetAnnouncedOk returns a tuple with the Announced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnounced

`func (o *ConnectorCountsResponse) SetAnnounced(v int64)`

SetAnnounced sets Announced field to given value.

### HasAnnounced

`func (o *ConnectorCountsResponse) HasAnnounced() bool`

HasAnnounced returns a boolean if a field has been set.

### GetConnected

`func (o *ConnectorCountsResponse) GetConnected() int64`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *ConnectorCountsResponse) GetConnectedOk() (*int64, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *ConnectorCountsResponse) SetConnected(v int64)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *ConnectorCountsResponse) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetDisabled

`func (o *ConnectorCountsResponse) GetDisabled() int64`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ConnectorCountsResponse) GetDisabledOk() (*int64, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ConnectorCountsResponse) SetDisabled(v int64)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ConnectorCountsResponse) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDisconnected

`func (o *ConnectorCountsResponse) GetDisconnected() int64`

GetDisconnected returns the Disconnected field if non-nil, zero value otherwise.

### GetDisconnectedOk

`func (o *ConnectorCountsResponse) GetDisconnectedOk() (*int64, bool)`

GetDisconnectedOk returns a tuple with the Disconnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnected

`func (o *ConnectorCountsResponse) SetDisconnected(v int64)`

SetDisconnected sets Disconnected field to given value.

### HasDisconnected

`func (o *ConnectorCountsResponse) HasDisconnected() bool`

HasDisconnected returns a boolean if a field has been set.

### GetExpired

`func (o *ConnectorCountsResponse) GetExpired() int64`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *ConnectorCountsResponse) GetExpiredOk() (*int64, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *ConnectorCountsResponse) SetExpired(v int64)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *ConnectorCountsResponse) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetReachable

`func (o *ConnectorCountsResponse) GetReachable() int64`

GetReachable returns the Reachable field if non-nil, zero value otherwise.

### GetReachableOk

`func (o *ConnectorCountsResponse) GetReachableOk() (*int64, bool)`

GetReachableOk returns a tuple with the Reachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachable

`func (o *ConnectorCountsResponse) SetReachable(v int64)`

SetReachable sets Reachable field to given value.

### HasReachable

`func (o *ConnectorCountsResponse) HasReachable() bool`

HasReachable returns a boolean if a field has been set.

### GetTotal

`func (o *ConnectorCountsResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ConnectorCountsResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ConnectorCountsResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ConnectorCountsResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUpgrading

`func (o *ConnectorCountsResponse) GetUpgrading() int64`

GetUpgrading returns the Upgrading field if non-nil, zero value otherwise.

### GetUpgradingOk

`func (o *ConnectorCountsResponse) GetUpgradingOk() (*int64, bool)`

GetUpgradingOk returns a tuple with the Upgrading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrading

`func (o *ConnectorCountsResponse) SetUpgrading(v int64)`

SetUpgrading sets Upgrading field to given value.

### HasUpgrading

`func (o *ConnectorCountsResponse) HasUpgrading() bool`

HasUpgrading returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


