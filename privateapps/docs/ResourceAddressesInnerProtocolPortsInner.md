# ResourceAddressesInnerProtocolPortsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to [**ProtocolClientToResource**](ProtocolClientToResource.md) |  | [optional] 
**Ports** | Pointer to **string** | The port number or list of comma-separated port numbers that you can use to connect to the Private Resource. A port number must be within the range of &#x60;1-65535&#x60;. | [optional] 

## Methods

### NewResourceAddressesInnerProtocolPortsInner

`func NewResourceAddressesInnerProtocolPortsInner() *ResourceAddressesInnerProtocolPortsInner`

NewResourceAddressesInnerProtocolPortsInner instantiates a new ResourceAddressesInnerProtocolPortsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceAddressesInnerProtocolPortsInnerWithDefaults

`func NewResourceAddressesInnerProtocolPortsInnerWithDefaults() *ResourceAddressesInnerProtocolPortsInner`

NewResourceAddressesInnerProtocolPortsInnerWithDefaults instantiates a new ResourceAddressesInnerProtocolPortsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *ResourceAddressesInnerProtocolPortsInner) GetProtocol() ProtocolClientToResource`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ResourceAddressesInnerProtocolPortsInner) GetProtocolOk() (*ProtocolClientToResource, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ResourceAddressesInnerProtocolPortsInner) SetProtocol(v ProtocolClientToResource)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ResourceAddressesInnerProtocolPortsInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPorts

`func (o *ResourceAddressesInnerProtocolPortsInner) GetPorts() string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ResourceAddressesInnerProtocolPortsInner) GetPortsOk() (*string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ResourceAddressesInnerProtocolPortsInner) SetPorts(v string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ResourceAddressesInnerProtocolPortsInner) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


