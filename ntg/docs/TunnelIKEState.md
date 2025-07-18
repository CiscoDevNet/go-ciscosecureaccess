# TunnelIKEState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | Pointer to **string** | Established state age in seconds. | [optional] [readonly] 
**DhGroup** | Pointer to **string** | IKE Diffie-Hellman group. | [optional] [readonly] 
**PrfAlgo** | Pointer to **string** | IKE pseudo random function. | [optional] [readonly] 
**EncAlgo** | Pointer to **string** | IKE encryption algorithm. | [optional] [readonly] 
**InitiatorSpi** | Pointer to **string** | Hex encoded initiator SPI / cookie. | [optional] [readonly] 
**ResponderSpi** | Pointer to **string** | Hex encoded responder SPI / cookie. | [optional] [readonly] 

## Methods

### NewTunnelIKEState

`func NewTunnelIKEState() *TunnelIKEState`

NewTunnelIKEState instantiates a new TunnelIKEState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelIKEStateWithDefaults

`func NewTunnelIKEStateWithDefaults() *TunnelIKEState`

NewTunnelIKEStateWithDefaults instantiates a new TunnelIKEState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *TunnelIKEState) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *TunnelIKEState) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *TunnelIKEState) SetAge(v string)`

SetAge sets Age field to given value.

### HasAge

`func (o *TunnelIKEState) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetDhGroup

`func (o *TunnelIKEState) GetDhGroup() string`

GetDhGroup returns the DhGroup field if non-nil, zero value otherwise.

### GetDhGroupOk

`func (o *TunnelIKEState) GetDhGroupOk() (*string, bool)`

GetDhGroupOk returns a tuple with the DhGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhGroup

`func (o *TunnelIKEState) SetDhGroup(v string)`

SetDhGroup sets DhGroup field to given value.

### HasDhGroup

`func (o *TunnelIKEState) HasDhGroup() bool`

HasDhGroup returns a boolean if a field has been set.

### GetPrfAlgo

`func (o *TunnelIKEState) GetPrfAlgo() string`

GetPrfAlgo returns the PrfAlgo field if non-nil, zero value otherwise.

### GetPrfAlgoOk

`func (o *TunnelIKEState) GetPrfAlgoOk() (*string, bool)`

GetPrfAlgoOk returns a tuple with the PrfAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrfAlgo

`func (o *TunnelIKEState) SetPrfAlgo(v string)`

SetPrfAlgo sets PrfAlgo field to given value.

### HasPrfAlgo

`func (o *TunnelIKEState) HasPrfAlgo() bool`

HasPrfAlgo returns a boolean if a field has been set.

### GetEncAlgo

`func (o *TunnelIKEState) GetEncAlgo() string`

GetEncAlgo returns the EncAlgo field if non-nil, zero value otherwise.

### GetEncAlgoOk

`func (o *TunnelIKEState) GetEncAlgoOk() (*string, bool)`

GetEncAlgoOk returns a tuple with the EncAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncAlgo

`func (o *TunnelIKEState) SetEncAlgo(v string)`

SetEncAlgo sets EncAlgo field to given value.

### HasEncAlgo

`func (o *TunnelIKEState) HasEncAlgo() bool`

HasEncAlgo returns a boolean if a field has been set.

### GetInitiatorSpi

`func (o *TunnelIKEState) GetInitiatorSpi() string`

GetInitiatorSpi returns the InitiatorSpi field if non-nil, zero value otherwise.

### GetInitiatorSpiOk

`func (o *TunnelIKEState) GetInitiatorSpiOk() (*string, bool)`

GetInitiatorSpiOk returns a tuple with the InitiatorSpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorSpi

`func (o *TunnelIKEState) SetInitiatorSpi(v string)`

SetInitiatorSpi sets InitiatorSpi field to given value.

### HasInitiatorSpi

`func (o *TunnelIKEState) HasInitiatorSpi() bool`

HasInitiatorSpi returns a boolean if a field has been set.

### GetResponderSpi

`func (o *TunnelIKEState) GetResponderSpi() string`

GetResponderSpi returns the ResponderSpi field if non-nil, zero value otherwise.

### GetResponderSpiOk

`func (o *TunnelIKEState) GetResponderSpiOk() (*string, bool)`

GetResponderSpiOk returns a tuple with the ResponderSpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponderSpi

`func (o *TunnelIKEState) SetResponderSpi(v string)`

SetResponderSpi sets ResponderSpi field to given value.

### HasResponderSpi

`func (o *TunnelIKEState) HasResponderSpi() bool`

HasResponderSpi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


