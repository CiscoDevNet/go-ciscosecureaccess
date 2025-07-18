# NetworkTunnelGroupStateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the Network Tunnel Group. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the Network Tunnel Group. A Network Tunnel Group name is a sequence of 1–50 characters. The &#x60;name&#x60; field cannot have any special characters other than spaces and hyphens. | [optional] 
**OrganizationId** | Pointer to **int64** | The ID of the organization. | [optional] [readonly] 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Hubs** | Pointer to [**[]HubsWithStateInner**](HubsWithStateInner.md) | The list of Hubs for a Network Tunnel Group. Only one Hub is the primary data center. | [optional] 

## Methods

### NewNetworkTunnelGroupStateResponse

`func NewNetworkTunnelGroupStateResponse() *NetworkTunnelGroupStateResponse`

NewNetworkTunnelGroupStateResponse instantiates a new NetworkTunnelGroupStateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkTunnelGroupStateResponseWithDefaults

`func NewNetworkTunnelGroupStateResponseWithDefaults() *NetworkTunnelGroupStateResponse`

NewNetworkTunnelGroupStateResponseWithDefaults instantiates a new NetworkTunnelGroupStateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkTunnelGroupStateResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkTunnelGroupStateResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkTunnelGroupStateResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkTunnelGroupStateResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NetworkTunnelGroupStateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkTunnelGroupStateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkTunnelGroupStateResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkTunnelGroupStateResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *NetworkTunnelGroupStateResponse) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *NetworkTunnelGroupStateResponse) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *NetworkTunnelGroupStateResponse) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *NetworkTunnelGroupStateResponse) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkTunnelGroupStateResponse) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkTunnelGroupStateResponse) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkTunnelGroupStateResponse) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkTunnelGroupStateResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHubs

`func (o *NetworkTunnelGroupStateResponse) GetHubs() []HubsWithStateInner`

GetHubs returns the Hubs field if non-nil, zero value otherwise.

### GetHubsOk

`func (o *NetworkTunnelGroupStateResponse) GetHubsOk() (*[]HubsWithStateInner, bool)`

GetHubsOk returns a tuple with the Hubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubs

`func (o *NetworkTunnelGroupStateResponse) SetHubs(v []HubsWithStateInner)`

SetHubs sets Hubs field to given value.

### HasHubs

`func (o *NetworkTunnelGroupStateResponse) HasHubs() bool`

HasHubs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


