# AddNetworkTunnelGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Network Tunnel Group. A Network Tunnel Group name is a sequence of 1–50 characters. The &#x60;name&#x60; field cannot have any special characters other than spaces and hyphens. | 
**Region** | **string** | The name of the region that is used to get the primary and secondary data centers for the Hubs. | 
**DeviceType** | Pointer to [**DeviceType**](DeviceType.md) |  | [optional] 
**AuthIdPrefix** | [**AddNetworkTunnelGroupRequestAuthIdPrefix**](AddNetworkTunnelGroupRequestAuthIdPrefix.md) |  | 
**Passphrase** | **string** | The passphrase for the primary and secondary tunnels. Provide a sequence of characters where the length of the passphrase is 16–64 characters. The passphrase must contain at least one upper and one lowercase letter as well as one numeral. The passphrase may not include special characters. | 
**Routing** | Pointer to [**RoutingRequest**](RoutingRequest.md) |  | [optional] 

## Methods

### NewAddNetworkTunnelGroupRequest

`func NewAddNetworkTunnelGroupRequest(name string, region string, authIdPrefix AddNetworkTunnelGroupRequestAuthIdPrefix, passphrase string, ) *AddNetworkTunnelGroupRequest`

NewAddNetworkTunnelGroupRequest instantiates a new AddNetworkTunnelGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddNetworkTunnelGroupRequestWithDefaults

`func NewAddNetworkTunnelGroupRequestWithDefaults() *AddNetworkTunnelGroupRequest`

NewAddNetworkTunnelGroupRequestWithDefaults instantiates a new AddNetworkTunnelGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddNetworkTunnelGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddNetworkTunnelGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddNetworkTunnelGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *AddNetworkTunnelGroupRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AddNetworkTunnelGroupRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AddNetworkTunnelGroupRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetDeviceType

`func (o *AddNetworkTunnelGroupRequest) GetDeviceType() DeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *AddNetworkTunnelGroupRequest) GetDeviceTypeOk() (*DeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *AddNetworkTunnelGroupRequest) SetDeviceType(v DeviceType)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *AddNetworkTunnelGroupRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetAuthIdPrefix

`func (o *AddNetworkTunnelGroupRequest) GetAuthIdPrefix() AddNetworkTunnelGroupRequestAuthIdPrefix`

GetAuthIdPrefix returns the AuthIdPrefix field if non-nil, zero value otherwise.

### GetAuthIdPrefixOk

`func (o *AddNetworkTunnelGroupRequest) GetAuthIdPrefixOk() (*AddNetworkTunnelGroupRequestAuthIdPrefix, bool)`

GetAuthIdPrefixOk returns a tuple with the AuthIdPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthIdPrefix

`func (o *AddNetworkTunnelGroupRequest) SetAuthIdPrefix(v AddNetworkTunnelGroupRequestAuthIdPrefix)`

SetAuthIdPrefix sets AuthIdPrefix field to given value.


### GetPassphrase

`func (o *AddNetworkTunnelGroupRequest) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *AddNetworkTunnelGroupRequest) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *AddNetworkTunnelGroupRequest) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.


### GetRouting

`func (o *AddNetworkTunnelGroupRequest) GetRouting() RoutingRequest`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *AddNetworkTunnelGroupRequest) GetRoutingOk() (*RoutingRequest, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *AddNetworkTunnelGroupRequest) SetRouting(v RoutingRequest)`

SetRouting sets Routing field to given value.

### HasRouting

`func (o *AddNetworkTunnelGroupRequest) HasRouting() bool`

HasRouting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


