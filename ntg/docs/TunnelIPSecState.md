# TunnelIPSecState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | Pointer to **string** | The installed state age in seconds. | [optional] [readonly] 
**IntegrityAlgo** | Pointer to **string** | The ESP or AH integrity algorithm name. | [optional] [readonly] 
**EncAlgo** | Pointer to **string** | The ESP encryption algorithm name. | [optional] [readonly] 
**EncKeySize** | Pointer to **string** | The ESP encryption key size (optional field, not available with NULL encryption). | [optional] [readonly] 
**SpiIn** | Pointer to **string** | The Hex encoded inbound SPI. | [optional] [readonly] 
**SpiOut** | Pointer to **string** | The Hex encoded outbound SPI. | [optional] [readonly] 
**PeerSelectors** | Pointer to **[]string** | The peer traffic selectors. | [optional] [readonly] 

## Methods

### NewTunnelIPSecState

`func NewTunnelIPSecState() *TunnelIPSecState`

NewTunnelIPSecState instantiates a new TunnelIPSecState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelIPSecStateWithDefaults

`func NewTunnelIPSecStateWithDefaults() *TunnelIPSecState`

NewTunnelIPSecStateWithDefaults instantiates a new TunnelIPSecState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *TunnelIPSecState) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *TunnelIPSecState) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *TunnelIPSecState) SetAge(v string)`

SetAge sets Age field to given value.

### HasAge

`func (o *TunnelIPSecState) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetIntegrityAlgo

`func (o *TunnelIPSecState) GetIntegrityAlgo() string`

GetIntegrityAlgo returns the IntegrityAlgo field if non-nil, zero value otherwise.

### GetIntegrityAlgoOk

`func (o *TunnelIPSecState) GetIntegrityAlgoOk() (*string, bool)`

GetIntegrityAlgoOk returns a tuple with the IntegrityAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityAlgo

`func (o *TunnelIPSecState) SetIntegrityAlgo(v string)`

SetIntegrityAlgo sets IntegrityAlgo field to given value.

### HasIntegrityAlgo

`func (o *TunnelIPSecState) HasIntegrityAlgo() bool`

HasIntegrityAlgo returns a boolean if a field has been set.

### GetEncAlgo

`func (o *TunnelIPSecState) GetEncAlgo() string`

GetEncAlgo returns the EncAlgo field if non-nil, zero value otherwise.

### GetEncAlgoOk

`func (o *TunnelIPSecState) GetEncAlgoOk() (*string, bool)`

GetEncAlgoOk returns a tuple with the EncAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncAlgo

`func (o *TunnelIPSecState) SetEncAlgo(v string)`

SetEncAlgo sets EncAlgo field to given value.

### HasEncAlgo

`func (o *TunnelIPSecState) HasEncAlgo() bool`

HasEncAlgo returns a boolean if a field has been set.

### GetEncKeySize

`func (o *TunnelIPSecState) GetEncKeySize() string`

GetEncKeySize returns the EncKeySize field if non-nil, zero value otherwise.

### GetEncKeySizeOk

`func (o *TunnelIPSecState) GetEncKeySizeOk() (*string, bool)`

GetEncKeySizeOk returns a tuple with the EncKeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncKeySize

`func (o *TunnelIPSecState) SetEncKeySize(v string)`

SetEncKeySize sets EncKeySize field to given value.

### HasEncKeySize

`func (o *TunnelIPSecState) HasEncKeySize() bool`

HasEncKeySize returns a boolean if a field has been set.

### GetSpiIn

`func (o *TunnelIPSecState) GetSpiIn() string`

GetSpiIn returns the SpiIn field if non-nil, zero value otherwise.

### GetSpiInOk

`func (o *TunnelIPSecState) GetSpiInOk() (*string, bool)`

GetSpiInOk returns a tuple with the SpiIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpiIn

`func (o *TunnelIPSecState) SetSpiIn(v string)`

SetSpiIn sets SpiIn field to given value.

### HasSpiIn

`func (o *TunnelIPSecState) HasSpiIn() bool`

HasSpiIn returns a boolean if a field has been set.

### GetSpiOut

`func (o *TunnelIPSecState) GetSpiOut() string`

GetSpiOut returns the SpiOut field if non-nil, zero value otherwise.

### GetSpiOutOk

`func (o *TunnelIPSecState) GetSpiOutOk() (*string, bool)`

GetSpiOutOk returns a tuple with the SpiOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpiOut

`func (o *TunnelIPSecState) SetSpiOut(v string)`

SetSpiOut sets SpiOut field to given value.

### HasSpiOut

`func (o *TunnelIPSecState) HasSpiOut() bool`

HasSpiOut returns a boolean if a field has been set.

### GetPeerSelectors

`func (o *TunnelIPSecState) GetPeerSelectors() []string`

GetPeerSelectors returns the PeerSelectors field if non-nil, zero value otherwise.

### GetPeerSelectorsOk

`func (o *TunnelIPSecState) GetPeerSelectorsOk() (*[]string, bool)`

GetPeerSelectorsOk returns a tuple with the PeerSelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerSelectors

`func (o *TunnelIPSecState) SetPeerSelectors(v []string)`

SetPeerSelectors sets PeerSelectors field to given value.

### HasPeerSelectors

`func (o *TunnelIPSecState) HasPeerSelectors() bool`

HasPeerSelectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


