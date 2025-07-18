# HubsWithIPInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the Hub. | [optional] [readonly] 
**IsPrimary** | Pointer to **bool** | Specifies whether the Hub is a primary data center. | [optional] [readonly] 
**Datacenter** | Pointer to [**DatacenterWithIP**](DatacenterWithIP.md) |  | [optional] 
**AuthId** | Pointer to **string** | An IP address or email used to authenticate the tunnel. | [optional] [readonly] 
**Status** | Pointer to [**HubState**](HubState.md) |  | [optional] 
**TunnelsCount** | Pointer to **int64** | The number of tunnels in the hub. | [optional] [readonly] 

## Methods

### NewHubsWithIPInner

`func NewHubsWithIPInner() *HubsWithIPInner`

NewHubsWithIPInner instantiates a new HubsWithIPInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHubsWithIPInnerWithDefaults

`func NewHubsWithIPInnerWithDefaults() *HubsWithIPInner`

NewHubsWithIPInnerWithDefaults instantiates a new HubsWithIPInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HubsWithIPInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HubsWithIPInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HubsWithIPInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *HubsWithIPInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsPrimary

`func (o *HubsWithIPInner) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *HubsWithIPInner) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *HubsWithIPInner) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *HubsWithIPInner) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetDatacenter

`func (o *HubsWithIPInner) GetDatacenter() DatacenterWithIP`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *HubsWithIPInner) GetDatacenterOk() (*DatacenterWithIP, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *HubsWithIPInner) SetDatacenter(v DatacenterWithIP)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *HubsWithIPInner) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetAuthId

`func (o *HubsWithIPInner) GetAuthId() string`

GetAuthId returns the AuthId field if non-nil, zero value otherwise.

### GetAuthIdOk

`func (o *HubsWithIPInner) GetAuthIdOk() (*string, bool)`

GetAuthIdOk returns a tuple with the AuthId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthId

`func (o *HubsWithIPInner) SetAuthId(v string)`

SetAuthId sets AuthId field to given value.

### HasAuthId

`func (o *HubsWithIPInner) HasAuthId() bool`

HasAuthId returns a boolean if a field has been set.

### GetStatus

`func (o *HubsWithIPInner) GetStatus() HubState`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HubsWithIPInner) GetStatusOk() (*HubState, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HubsWithIPInner) SetStatus(v HubState)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HubsWithIPInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTunnelsCount

`func (o *HubsWithIPInner) GetTunnelsCount() int64`

GetTunnelsCount returns the TunnelsCount field if non-nil, zero value otherwise.

### GetTunnelsCountOk

`func (o *HubsWithIPInner) GetTunnelsCountOk() (*int64, bool)`

GetTunnelsCountOk returns a tuple with the TunnelsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelsCount

`func (o *HubsWithIPInner) SetTunnelsCount(v int64)`

SetTunnelsCount sets TunnelsCount field to given value.

### HasTunnelsCount

`func (o *HubsWithIPInner) HasTunnelsCount() bool`

HasTunnelsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


