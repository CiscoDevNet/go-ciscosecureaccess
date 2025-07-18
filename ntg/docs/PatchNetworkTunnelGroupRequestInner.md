# PatchNetworkTunnelGroupRequestInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | **string** | The operation that needs to be done. The only available operation is &#x60;replace&#x60;. | 
**Path** | **string** | The path of the property that needs to be updated. Available paths are &#x60;/name&#x60;, &#x60;/authIdPrefix&#x60;, &#x60;/passphrase&#x60;, &#x60;/region&#x60;, and &#x60;/routing&#x60;. | 
**Value** | [**PatchNetworkTunnelGroupRequestInnerValue**](PatchNetworkTunnelGroupRequestInnerValue.md) |  | 

## Methods

### NewPatchNetworkTunnelGroupRequestInner

`func NewPatchNetworkTunnelGroupRequestInner(op string, path string, value PatchNetworkTunnelGroupRequestInnerValue, ) *PatchNetworkTunnelGroupRequestInner`

NewPatchNetworkTunnelGroupRequestInner instantiates a new PatchNetworkTunnelGroupRequestInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchNetworkTunnelGroupRequestInnerWithDefaults

`func NewPatchNetworkTunnelGroupRequestInnerWithDefaults() *PatchNetworkTunnelGroupRequestInner`

NewPatchNetworkTunnelGroupRequestInnerWithDefaults instantiates a new PatchNetworkTunnelGroupRequestInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *PatchNetworkTunnelGroupRequestInner) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *PatchNetworkTunnelGroupRequestInner) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *PatchNetworkTunnelGroupRequestInner) SetOp(v string)`

SetOp sets Op field to given value.


### GetPath

`func (o *PatchNetworkTunnelGroupRequestInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PatchNetworkTunnelGroupRequestInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PatchNetworkTunnelGroupRequestInner) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *PatchNetworkTunnelGroupRequestInner) GetValue() PatchNetworkTunnelGroupRequestInnerValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PatchNetworkTunnelGroupRequestInner) GetValueOk() (*PatchNetworkTunnelGroupRequestInnerValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PatchNetworkTunnelGroupRequestInner) SetValue(v PatchNetworkTunnelGroupRequestInnerValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


