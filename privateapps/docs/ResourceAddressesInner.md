# ResourceAddressesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationAddr** | **[]string** | The list of IP address, CIDRs, FQDN, or wildcard FQDN destinations. IPv6 is not supported. | 
**ProtocolPorts** | [**[]ResourceAddressesInnerProtocolPortsInner**](ResourceAddressesInnerProtocolPortsInner.md) | The list of protocols and ports for the IP address destinations. The protocols must be unique. | 

## Methods

### NewResourceAddressesInner

`func NewResourceAddressesInner(destinationAddr []string, protocolPorts []ResourceAddressesInnerProtocolPortsInner, ) *ResourceAddressesInner`

NewResourceAddressesInner instantiates a new ResourceAddressesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceAddressesInnerWithDefaults

`func NewResourceAddressesInnerWithDefaults() *ResourceAddressesInner`

NewResourceAddressesInnerWithDefaults instantiates a new ResourceAddressesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationAddr

`func (o *ResourceAddressesInner) GetDestinationAddr() []string`

GetDestinationAddr returns the DestinationAddr field if non-nil, zero value otherwise.

### GetDestinationAddrOk

`func (o *ResourceAddressesInner) GetDestinationAddrOk() (*[]string, bool)`

GetDestinationAddrOk returns a tuple with the DestinationAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddr

`func (o *ResourceAddressesInner) SetDestinationAddr(v []string)`

SetDestinationAddr sets DestinationAddr field to given value.


### GetProtocolPorts

`func (o *ResourceAddressesInner) GetProtocolPorts() []ResourceAddressesInnerProtocolPortsInner`

GetProtocolPorts returns the ProtocolPorts field if non-nil, zero value otherwise.

### GetProtocolPortsOk

`func (o *ResourceAddressesInner) GetProtocolPortsOk() (*[]ResourceAddressesInnerProtocolPortsInner, bool)`

GetProtocolPortsOk returns a tuple with the ProtocolPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolPorts

`func (o *ResourceAddressesInner) SetProtocolPorts(v []ResourceAddressesInnerProtocolPortsInner)`

SetProtocolPorts sets ProtocolPorts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


