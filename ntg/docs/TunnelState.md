# TunnelState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **time.Time** | The date and time (UTC time with milliseconds) when the state event record was generated. | [optional] [readonly] 
**Status** | Pointer to **string** | The high-level status of the tunnel: * UP - The tunnel is active. * DOWN - The tunnel is inactive. * FAILED - The tunnel is in a failed state. * UNKNOWN - The current status is unknown and pending updated information. | [optional] [readonly] 
**Dc** | Pointer to **string** | The domain name of the data center. | [optional] [readonly] 
**DcName** | Pointer to **string** | The name of the data center. | [optional] [readonly] 
**DcDesc** | Pointer to **string** | The city and country, or regional location of the data center. | [optional] [readonly] 
**IkeState** | Pointer to **string** | IKE SA State: * CREATED * CONNECTING * ESTABLISHED * PASSIVE * REKEYING * REKEYED * DELETING * DESTROYING | [optional] [readonly] 
**IpsecState** | Pointer to **string** | IPsec state: * CREATED * ROUTED * INSTALLING * INSTALLED * UPDATING * REKEYING * REKEYED * RETRYING * DELETING * DELETED * DESTROYING | [optional] [readonly] 
**PeerId** | Pointer to **string** | The IKE ID of the remote peer. | [optional] [readonly] 
**PeerIp** | Pointer to **string** | The IP address of the remote peer. | [optional] [readonly] 
**PeerPort** | Pointer to **string** | The port of the remote peer. | [optional] [readonly] 
**LocalIp** | Pointer to **string** | The public IP address assigned to an endpoint device (for example: ISR, Viptela). | [optional] [readonly] 
**LocalPort** | Pointer to **string** | The port on the device. | [optional] [readonly] 
**Ike** | Pointer to [**TunnelIKEState**](TunnelIKEState.md) |  | [optional] 
**Ipsec** | Pointer to [**TunnelIPSecState**](TunnelIPSecState.md) |  | [optional] 
**Data** | Pointer to [**TunnelDataState**](TunnelDataState.md) |  | [optional] 
**RoutingStats** | Pointer to [**TunnelRoutingStats**](TunnelRoutingStats.md) |  | [optional] 

## Methods

### NewTunnelState

`func NewTunnelState() *TunnelState`

NewTunnelState instantiates a new TunnelState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelStateWithDefaults

`func NewTunnelStateWithDefaults() *TunnelState`

NewTunnelStateWithDefaults instantiates a new TunnelState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *TunnelState) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *TunnelState) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *TunnelState) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *TunnelState) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetStatus

`func (o *TunnelState) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TunnelState) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TunnelState) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TunnelState) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDc

`func (o *TunnelState) GetDc() string`

GetDc returns the Dc field if non-nil, zero value otherwise.

### GetDcOk

`func (o *TunnelState) GetDcOk() (*string, bool)`

GetDcOk returns a tuple with the Dc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDc

`func (o *TunnelState) SetDc(v string)`

SetDc sets Dc field to given value.

### HasDc

`func (o *TunnelState) HasDc() bool`

HasDc returns a boolean if a field has been set.

### GetDcName

`func (o *TunnelState) GetDcName() string`

GetDcName returns the DcName field if non-nil, zero value otherwise.

### GetDcNameOk

`func (o *TunnelState) GetDcNameOk() (*string, bool)`

GetDcNameOk returns a tuple with the DcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcName

`func (o *TunnelState) SetDcName(v string)`

SetDcName sets DcName field to given value.

### HasDcName

`func (o *TunnelState) HasDcName() bool`

HasDcName returns a boolean if a field has been set.

### GetDcDesc

`func (o *TunnelState) GetDcDesc() string`

GetDcDesc returns the DcDesc field if non-nil, zero value otherwise.

### GetDcDescOk

`func (o *TunnelState) GetDcDescOk() (*string, bool)`

GetDcDescOk returns a tuple with the DcDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcDesc

`func (o *TunnelState) SetDcDesc(v string)`

SetDcDesc sets DcDesc field to given value.

### HasDcDesc

`func (o *TunnelState) HasDcDesc() bool`

HasDcDesc returns a boolean if a field has been set.

### GetIkeState

`func (o *TunnelState) GetIkeState() string`

GetIkeState returns the IkeState field if non-nil, zero value otherwise.

### GetIkeStateOk

`func (o *TunnelState) GetIkeStateOk() (*string, bool)`

GetIkeStateOk returns a tuple with the IkeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeState

`func (o *TunnelState) SetIkeState(v string)`

SetIkeState sets IkeState field to given value.

### HasIkeState

`func (o *TunnelState) HasIkeState() bool`

HasIkeState returns a boolean if a field has been set.

### GetIpsecState

`func (o *TunnelState) GetIpsecState() string`

GetIpsecState returns the IpsecState field if non-nil, zero value otherwise.

### GetIpsecStateOk

`func (o *TunnelState) GetIpsecStateOk() (*string, bool)`

GetIpsecStateOk returns a tuple with the IpsecState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecState

`func (o *TunnelState) SetIpsecState(v string)`

SetIpsecState sets IpsecState field to given value.

### HasIpsecState

`func (o *TunnelState) HasIpsecState() bool`

HasIpsecState returns a boolean if a field has been set.

### GetPeerId

`func (o *TunnelState) GetPeerId() string`

GetPeerId returns the PeerId field if non-nil, zero value otherwise.

### GetPeerIdOk

`func (o *TunnelState) GetPeerIdOk() (*string, bool)`

GetPeerIdOk returns a tuple with the PeerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerId

`func (o *TunnelState) SetPeerId(v string)`

SetPeerId sets PeerId field to given value.

### HasPeerId

`func (o *TunnelState) HasPeerId() bool`

HasPeerId returns a boolean if a field has been set.

### GetPeerIp

`func (o *TunnelState) GetPeerIp() string`

GetPeerIp returns the PeerIp field if non-nil, zero value otherwise.

### GetPeerIpOk

`func (o *TunnelState) GetPeerIpOk() (*string, bool)`

GetPeerIpOk returns a tuple with the PeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerIp

`func (o *TunnelState) SetPeerIp(v string)`

SetPeerIp sets PeerIp field to given value.

### HasPeerIp

`func (o *TunnelState) HasPeerIp() bool`

HasPeerIp returns a boolean if a field has been set.

### GetPeerPort

`func (o *TunnelState) GetPeerPort() string`

GetPeerPort returns the PeerPort field if non-nil, zero value otherwise.

### GetPeerPortOk

`func (o *TunnelState) GetPeerPortOk() (*string, bool)`

GetPeerPortOk returns a tuple with the PeerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerPort

`func (o *TunnelState) SetPeerPort(v string)`

SetPeerPort sets PeerPort field to given value.

### HasPeerPort

`func (o *TunnelState) HasPeerPort() bool`

HasPeerPort returns a boolean if a field has been set.

### GetLocalIp

`func (o *TunnelState) GetLocalIp() string`

GetLocalIp returns the LocalIp field if non-nil, zero value otherwise.

### GetLocalIpOk

`func (o *TunnelState) GetLocalIpOk() (*string, bool)`

GetLocalIpOk returns a tuple with the LocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIp

`func (o *TunnelState) SetLocalIp(v string)`

SetLocalIp sets LocalIp field to given value.

### HasLocalIp

`func (o *TunnelState) HasLocalIp() bool`

HasLocalIp returns a boolean if a field has been set.

### GetLocalPort

`func (o *TunnelState) GetLocalPort() string`

GetLocalPort returns the LocalPort field if non-nil, zero value otherwise.

### GetLocalPortOk

`func (o *TunnelState) GetLocalPortOk() (*string, bool)`

GetLocalPortOk returns a tuple with the LocalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPort

`func (o *TunnelState) SetLocalPort(v string)`

SetLocalPort sets LocalPort field to given value.

### HasLocalPort

`func (o *TunnelState) HasLocalPort() bool`

HasLocalPort returns a boolean if a field has been set.

### GetIke

`func (o *TunnelState) GetIke() TunnelIKEState`

GetIke returns the Ike field if non-nil, zero value otherwise.

### GetIkeOk

`func (o *TunnelState) GetIkeOk() (*TunnelIKEState, bool)`

GetIkeOk returns a tuple with the Ike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIke

`func (o *TunnelState) SetIke(v TunnelIKEState)`

SetIke sets Ike field to given value.

### HasIke

`func (o *TunnelState) HasIke() bool`

HasIke returns a boolean if a field has been set.

### GetIpsec

`func (o *TunnelState) GetIpsec() TunnelIPSecState`

GetIpsec returns the Ipsec field if non-nil, zero value otherwise.

### GetIpsecOk

`func (o *TunnelState) GetIpsecOk() (*TunnelIPSecState, bool)`

GetIpsecOk returns a tuple with the Ipsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsec

`func (o *TunnelState) SetIpsec(v TunnelIPSecState)`

SetIpsec sets Ipsec field to given value.

### HasIpsec

`func (o *TunnelState) HasIpsec() bool`

HasIpsec returns a boolean if a field has been set.

### GetData

`func (o *TunnelState) GetData() TunnelDataState`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TunnelState) GetDataOk() (*TunnelDataState, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TunnelState) SetData(v TunnelDataState)`

SetData sets Data field to given value.

### HasData

`func (o *TunnelState) HasData() bool`

HasData returns a boolean if a field has been set.

### GetRoutingStats

`func (o *TunnelState) GetRoutingStats() TunnelRoutingStats`

GetRoutingStats returns the RoutingStats field if non-nil, zero value otherwise.

### GetRoutingStatsOk

`func (o *TunnelState) GetRoutingStatsOk() (*TunnelRoutingStats, bool)`

GetRoutingStatsOk returns a tuple with the RoutingStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingStats

`func (o *TunnelState) SetRoutingStats(v TunnelRoutingStats)`

SetRoutingStats sets RoutingStats field to given value.

### HasRoutingStats

`func (o *TunnelState) HasRoutingStats() bool`

HasRoutingStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


