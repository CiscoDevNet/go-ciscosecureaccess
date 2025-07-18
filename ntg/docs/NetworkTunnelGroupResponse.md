# NetworkTunnelGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the Network Tunnel Group. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the Network Tunnel Group. A Network Tunnel Group name is a sequence of 1–50 characters. The &#x60;name&#x60; field cannot have any special characters other than spaces and hyphens. | [optional] 
**OrganizationId** | Pointer to **int64** | The ID of the organization. | [optional] [readonly] 
**DeviceType** | Pointer to [**DeviceType**](DeviceType.md) |  | [optional] 
**Region** | Pointer to **string** | The name of the region that is used to get the primary and secondary data centers for the Hubs. | [optional] 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Hubs** | Pointer to [**[]HubsWithIPInner**](HubsWithIPInner.md) | The list of Hubs for a Network Tunnel Group. Only one Hub is the primary data center. | [optional] 
**Routing** | Pointer to [**RoutingResponse**](RoutingResponse.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time (timestamp) when the network tunnel group was created. | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** | The date and time of the last update (timestamp) for the network tunnel group. | [optional] [readonly] 

## Methods

### NewNetworkTunnelGroupResponse

`func NewNetworkTunnelGroupResponse() *NetworkTunnelGroupResponse`

NewNetworkTunnelGroupResponse instantiates a new NetworkTunnelGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkTunnelGroupResponseWithDefaults

`func NewNetworkTunnelGroupResponseWithDefaults() *NetworkTunnelGroupResponse`

NewNetworkTunnelGroupResponseWithDefaults instantiates a new NetworkTunnelGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkTunnelGroupResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkTunnelGroupResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkTunnelGroupResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkTunnelGroupResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NetworkTunnelGroupResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkTunnelGroupResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkTunnelGroupResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkTunnelGroupResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *NetworkTunnelGroupResponse) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *NetworkTunnelGroupResponse) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *NetworkTunnelGroupResponse) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *NetworkTunnelGroupResponse) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetDeviceType

`func (o *NetworkTunnelGroupResponse) GetDeviceType() DeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *NetworkTunnelGroupResponse) GetDeviceTypeOk() (*DeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *NetworkTunnelGroupResponse) SetDeviceType(v DeviceType)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *NetworkTunnelGroupResponse) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetRegion

`func (o *NetworkTunnelGroupResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NetworkTunnelGroupResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NetworkTunnelGroupResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *NetworkTunnelGroupResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkTunnelGroupResponse) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkTunnelGroupResponse) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkTunnelGroupResponse) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkTunnelGroupResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHubs

`func (o *NetworkTunnelGroupResponse) GetHubs() []HubsWithIPInner`

GetHubs returns the Hubs field if non-nil, zero value otherwise.

### GetHubsOk

`func (o *NetworkTunnelGroupResponse) GetHubsOk() (*[]HubsWithIPInner, bool)`

GetHubsOk returns a tuple with the Hubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubs

`func (o *NetworkTunnelGroupResponse) SetHubs(v []HubsWithIPInner)`

SetHubs sets Hubs field to given value.

### HasHubs

`func (o *NetworkTunnelGroupResponse) HasHubs() bool`

HasHubs returns a boolean if a field has been set.

### GetRouting

`func (o *NetworkTunnelGroupResponse) GetRouting() RoutingResponse`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *NetworkTunnelGroupResponse) GetRoutingOk() (*RoutingResponse, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *NetworkTunnelGroupResponse) SetRouting(v RoutingResponse)`

SetRouting sets Routing field to given value.

### HasRouting

`func (o *NetworkTunnelGroupResponse) HasRouting() bool`

HasRouting returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NetworkTunnelGroupResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkTunnelGroupResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkTunnelGroupResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NetworkTunnelGroupResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *NetworkTunnelGroupResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *NetworkTunnelGroupResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *NetworkTunnelGroupResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *NetworkTunnelGroupResponse) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


