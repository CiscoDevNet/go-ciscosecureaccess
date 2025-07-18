# TunnelDataState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PacketsIn** | Pointer to **string** | The number of processed input packets (tunnel ingress). | [optional] [readonly] 
**BytesIn** | Pointer to **string** | The number of processed input bytes (tunnel ingress). | [optional] [readonly] 
**IdleTimeIn** | Pointer to **string** | The seconds since the last inbound packet (the time that the tunnel is idle). | [optional] [readonly] 
**PacketsOut** | Pointer to **string** | The number of processed output packets (tunnel egress). | [optional] [readonly] 
**BytesOut** | Pointer to **string** | The number of processed output bytes (tunnel egress). | [optional] [readonly] 
**IdleTimeOut** | Pointer to **string** | The seconds since the last outbound packet (the time that the tunnel is idle). | [optional] [readonly] 
**Initialized** | Pointer to **string** | The time when the packet and byte counters were initialized to 0. | [optional] [readonly] 

## Methods

### NewTunnelDataState

`func NewTunnelDataState() *TunnelDataState`

NewTunnelDataState instantiates a new TunnelDataState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelDataStateWithDefaults

`func NewTunnelDataStateWithDefaults() *TunnelDataState`

NewTunnelDataStateWithDefaults instantiates a new TunnelDataState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPacketsIn

`func (o *TunnelDataState) GetPacketsIn() string`

GetPacketsIn returns the PacketsIn field if non-nil, zero value otherwise.

### GetPacketsInOk

`func (o *TunnelDataState) GetPacketsInOk() (*string, bool)`

GetPacketsInOk returns a tuple with the PacketsIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketsIn

`func (o *TunnelDataState) SetPacketsIn(v string)`

SetPacketsIn sets PacketsIn field to given value.

### HasPacketsIn

`func (o *TunnelDataState) HasPacketsIn() bool`

HasPacketsIn returns a boolean if a field has been set.

### GetBytesIn

`func (o *TunnelDataState) GetBytesIn() string`

GetBytesIn returns the BytesIn field if non-nil, zero value otherwise.

### GetBytesInOk

`func (o *TunnelDataState) GetBytesInOk() (*string, bool)`

GetBytesInOk returns a tuple with the BytesIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesIn

`func (o *TunnelDataState) SetBytesIn(v string)`

SetBytesIn sets BytesIn field to given value.

### HasBytesIn

`func (o *TunnelDataState) HasBytesIn() bool`

HasBytesIn returns a boolean if a field has been set.

### GetIdleTimeIn

`func (o *TunnelDataState) GetIdleTimeIn() string`

GetIdleTimeIn returns the IdleTimeIn field if non-nil, zero value otherwise.

### GetIdleTimeInOk

`func (o *TunnelDataState) GetIdleTimeInOk() (*string, bool)`

GetIdleTimeInOk returns a tuple with the IdleTimeIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeIn

`func (o *TunnelDataState) SetIdleTimeIn(v string)`

SetIdleTimeIn sets IdleTimeIn field to given value.

### HasIdleTimeIn

`func (o *TunnelDataState) HasIdleTimeIn() bool`

HasIdleTimeIn returns a boolean if a field has been set.

### GetPacketsOut

`func (o *TunnelDataState) GetPacketsOut() string`

GetPacketsOut returns the PacketsOut field if non-nil, zero value otherwise.

### GetPacketsOutOk

`func (o *TunnelDataState) GetPacketsOutOk() (*string, bool)`

GetPacketsOutOk returns a tuple with the PacketsOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketsOut

`func (o *TunnelDataState) SetPacketsOut(v string)`

SetPacketsOut sets PacketsOut field to given value.

### HasPacketsOut

`func (o *TunnelDataState) HasPacketsOut() bool`

HasPacketsOut returns a boolean if a field has been set.

### GetBytesOut

`func (o *TunnelDataState) GetBytesOut() string`

GetBytesOut returns the BytesOut field if non-nil, zero value otherwise.

### GetBytesOutOk

`func (o *TunnelDataState) GetBytesOutOk() (*string, bool)`

GetBytesOutOk returns a tuple with the BytesOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesOut

`func (o *TunnelDataState) SetBytesOut(v string)`

SetBytesOut sets BytesOut field to given value.

### HasBytesOut

`func (o *TunnelDataState) HasBytesOut() bool`

HasBytesOut returns a boolean if a field has been set.

### GetIdleTimeOut

`func (o *TunnelDataState) GetIdleTimeOut() string`

GetIdleTimeOut returns the IdleTimeOut field if non-nil, zero value otherwise.

### GetIdleTimeOutOk

`func (o *TunnelDataState) GetIdleTimeOutOk() (*string, bool)`

GetIdleTimeOutOk returns a tuple with the IdleTimeOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeOut

`func (o *TunnelDataState) SetIdleTimeOut(v string)`

SetIdleTimeOut sets IdleTimeOut field to given value.

### HasIdleTimeOut

`func (o *TunnelDataState) HasIdleTimeOut() bool`

HasIdleTimeOut returns a boolean if a field has been set.

### GetInitialized

`func (o *TunnelDataState) GetInitialized() string`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *TunnelDataState) GetInitializedOk() (*string, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *TunnelDataState) SetInitialized(v string)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *TunnelDataState) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


