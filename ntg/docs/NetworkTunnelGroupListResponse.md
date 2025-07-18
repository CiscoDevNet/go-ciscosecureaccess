# NetworkTunnelGroupListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the Network Tunnel Group. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the Network Tunnel Group. A Network Tunnel Group name is a sequence of 1–50 characters. The &#x60;name&#x60; field cannot have any special characters other than spaces and hyphens. | [optional] 
**OrganizationId** | Pointer to **int64** | The ID of the organization. | [optional] [readonly] 
**DeviceType** | Pointer to [**DeviceType**](DeviceType.md) |  | [optional] 
**Region** | Pointer to **string** | The name of the region that is used to get the primary and secondary data centers for the Hubs. | [optional] 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Hubs** | Pointer to [**[]HubsInner**](HubsInner.md) | The list of Hubs for a Network Tunnel Group. Only one Hub is the primary data center. | [optional] 
**Routing** | Pointer to [**RoutingResponse**](RoutingResponse.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time (timestamp) when the network tunnel group was created. | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** | The date and time of the last update (timestamp) for the network tunnel group. | [optional] [readonly] 

## Methods

### NewNetworkTunnelGroupListResponse

`func NewNetworkTunnelGroupListResponse() *NetworkTunnelGroupListResponse`

NewNetworkTunnelGroupListResponse instantiates a new NetworkTunnelGroupListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkTunnelGroupListResponseWithDefaults

`func NewNetworkTunnelGroupListResponseWithDefaults() *NetworkTunnelGroupListResponse`

NewNetworkTunnelGroupListResponseWithDefaults instantiates a new NetworkTunnelGroupListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkTunnelGroupListResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkTunnelGroupListResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkTunnelGroupListResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkTunnelGroupListResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NetworkTunnelGroupListResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkTunnelGroupListResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkTunnelGroupListResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkTunnelGroupListResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *NetworkTunnelGroupListResponse) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *NetworkTunnelGroupListResponse) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *NetworkTunnelGroupListResponse) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *NetworkTunnelGroupListResponse) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetDeviceType

`func (o *NetworkTunnelGroupListResponse) GetDeviceType() DeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *NetworkTunnelGroupListResponse) GetDeviceTypeOk() (*DeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *NetworkTunnelGroupListResponse) SetDeviceType(v DeviceType)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *NetworkTunnelGroupListResponse) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetRegion

`func (o *NetworkTunnelGroupListResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NetworkTunnelGroupListResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NetworkTunnelGroupListResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *NetworkTunnelGroupListResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkTunnelGroupListResponse) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkTunnelGroupListResponse) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkTunnelGroupListResponse) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkTunnelGroupListResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHubs

`func (o *NetworkTunnelGroupListResponse) GetHubs() []HubsInner`

GetHubs returns the Hubs field if non-nil, zero value otherwise.

### GetHubsOk

`func (o *NetworkTunnelGroupListResponse) GetHubsOk() (*[]HubsInner, bool)`

GetHubsOk returns a tuple with the Hubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubs

`func (o *NetworkTunnelGroupListResponse) SetHubs(v []HubsInner)`

SetHubs sets Hubs field to given value.

### HasHubs

`func (o *NetworkTunnelGroupListResponse) HasHubs() bool`

HasHubs returns a boolean if a field has been set.

### GetRouting

`func (o *NetworkTunnelGroupListResponse) GetRouting() RoutingResponse`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *NetworkTunnelGroupListResponse) GetRoutingOk() (*RoutingResponse, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *NetworkTunnelGroupListResponse) SetRouting(v RoutingResponse)`

SetRouting sets Routing field to given value.

### HasRouting

`func (o *NetworkTunnelGroupListResponse) HasRouting() bool`

HasRouting returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NetworkTunnelGroupListResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkTunnelGroupListResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkTunnelGroupListResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NetworkTunnelGroupListResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *NetworkTunnelGroupListResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *NetworkTunnelGroupListResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *NetworkTunnelGroupListResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *NetworkTunnelGroupListResponse) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


