# Bgp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsNumber** | Pointer to **string** | The BGP autonomous system (AS) number. | [optional] [readonly] 
**PeerIPs** | Pointer to **[]string** | The list of BGP peer IP addresses. | [optional] 
**PeerRange** | Pointer to **string** | The range of BGP peer addresses. | [optional] [readonly] 

## Methods

### NewBgp

`func NewBgp() *Bgp`

NewBgp instantiates a new Bgp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpWithDefaults

`func NewBgpWithDefaults() *Bgp`

NewBgpWithDefaults instantiates a new Bgp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsNumber

`func (o *Bgp) GetAsNumber() string`

GetAsNumber returns the AsNumber field if non-nil, zero value otherwise.

### GetAsNumberOk

`func (o *Bgp) GetAsNumberOk() (*string, bool)`

GetAsNumberOk returns a tuple with the AsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsNumber

`func (o *Bgp) SetAsNumber(v string)`

SetAsNumber sets AsNumber field to given value.

### HasAsNumber

`func (o *Bgp) HasAsNumber() bool`

HasAsNumber returns a boolean if a field has been set.

### GetPeerIPs

`func (o *Bgp) GetPeerIPs() []string`

GetPeerIPs returns the PeerIPs field if non-nil, zero value otherwise.

### GetPeerIPsOk

`func (o *Bgp) GetPeerIPsOk() (*[]string, bool)`

GetPeerIPsOk returns a tuple with the PeerIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerIPs

`func (o *Bgp) SetPeerIPs(v []string)`

SetPeerIPs sets PeerIPs field to given value.

### HasPeerIPs

`func (o *Bgp) HasPeerIPs() bool`

HasPeerIPs returns a boolean if a field has been set.

### GetPeerRange

`func (o *Bgp) GetPeerRange() string`

GetPeerRange returns the PeerRange field if non-nil, zero value otherwise.

### GetPeerRangeOk

`func (o *Bgp) GetPeerRangeOk() (*string, bool)`

GetPeerRangeOk returns a tuple with the PeerRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerRange

`func (o *Bgp) SetPeerRange(v string)`

SetPeerRange sets PeerRange field to given value.

### HasPeerRange

`func (o *Bgp) HasPeerRange() bool`

HasPeerRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


