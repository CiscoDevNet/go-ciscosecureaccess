# AntivirusThreats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Puas** | **[]map[string]interface{}** | The list of potentially unwanted applications. | 
**Viruses** | **[]string** | The list of viruses. | 
**Others** | **[]map[string]interface{}** | The list of other antivirus threats. | 

## Methods

### NewAntivirusThreats

`func NewAntivirusThreats(puas []map[string]interface{}, viruses []string, others []map[string]interface{}, ) *AntivirusThreats`

NewAntivirusThreats instantiates a new AntivirusThreats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAntivirusThreatsWithDefaults

`func NewAntivirusThreatsWithDefaults() *AntivirusThreats`

NewAntivirusThreatsWithDefaults instantiates a new AntivirusThreats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPuas

`func (o *AntivirusThreats) GetPuas() []map[string]interface{}`

GetPuas returns the Puas field if non-nil, zero value otherwise.

### GetPuasOk

`func (o *AntivirusThreats) GetPuasOk() (*[]map[string]interface{}, bool)`

GetPuasOk returns a tuple with the Puas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuas

`func (o *AntivirusThreats) SetPuas(v []map[string]interface{})`

SetPuas sets Puas field to given value.


### GetViruses

`func (o *AntivirusThreats) GetViruses() []string`

GetViruses returns the Viruses field if non-nil, zero value otherwise.

### GetVirusesOk

`func (o *AntivirusThreats) GetVirusesOk() (*[]string, bool)`

GetVirusesOk returns a tuple with the Viruses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViruses

`func (o *AntivirusThreats) SetViruses(v []string)`

SetViruses sets Viruses field to given value.


### GetOthers

`func (o *AntivirusThreats) GetOthers() []map[string]interface{}`

GetOthers returns the Others field if non-nil, zero value otherwise.

### GetOthersOk

`func (o *AntivirusThreats) GetOthersOk() (*[]map[string]interface{}, bool)`

GetOthersOk returns a tuple with the Others field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOthers

`func (o *AntivirusThreats) SetOthers(v []map[string]interface{})`

SetOthers sets Others field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


