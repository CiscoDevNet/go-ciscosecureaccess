# HubsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the Hub. | [optional] [readonly] 
**IsPrimary** | Pointer to **bool** | Specifies whether the Hub is a primary data center. | [optional] [readonly] 
**Datacenter** | Pointer to [**Datacenter**](Datacenter.md) |  | [optional] 
**AuthId** | Pointer to **string** | An IP address or email used to authenticate the tunnel. | [optional] [readonly] 
**Status** | Pointer to [**HubStatus**](HubStatus.md) |  | [optional] 
**TunnelsCount** | Pointer to **int64** | The number of tunnels in the hub. | [optional] [readonly] 

## Methods

### NewHubsInner

`func NewHubsInner() *HubsInner`

NewHubsInner instantiates a new HubsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHubsInnerWithDefaults

`func NewHubsInnerWithDefaults() *HubsInner`

NewHubsInnerWithDefaults instantiates a new HubsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HubsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HubsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HubsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *HubsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsPrimary

`func (o *HubsInner) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *HubsInner) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *HubsInner) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *HubsInner) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetDatacenter

`func (o *HubsInner) GetDatacenter() Datacenter`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *HubsInner) GetDatacenterOk() (*Datacenter, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *HubsInner) SetDatacenter(v Datacenter)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *HubsInner) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetAuthId

`func (o *HubsInner) GetAuthId() string`

GetAuthId returns the AuthId field if non-nil, zero value otherwise.

### GetAuthIdOk

`func (o *HubsInner) GetAuthIdOk() (*string, bool)`

GetAuthIdOk returns a tuple with the AuthId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthId

`func (o *HubsInner) SetAuthId(v string)`

SetAuthId sets AuthId field to given value.

### HasAuthId

`func (o *HubsInner) HasAuthId() bool`

HasAuthId returns a boolean if a field has been set.

### GetStatus

`func (o *HubsInner) GetStatus() HubStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HubsInner) GetStatusOk() (*HubStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HubsInner) SetStatus(v HubStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HubsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTunnelsCount

`func (o *HubsInner) GetTunnelsCount() int64`

GetTunnelsCount returns the TunnelsCount field if non-nil, zero value otherwise.

### GetTunnelsCountOk

`func (o *HubsInner) GetTunnelsCountOk() (*int64, bool)`

GetTunnelsCountOk returns a tuple with the TunnelsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelsCount

`func (o *HubsInner) SetTunnelsCount(v int64)`

SetTunnelsCount sets TunnelsCount field to given value.

### HasTunnelsCount

`func (o *HubsInner) HasTunnelsCount() bool`

HasTunnelsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


