# Signature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Generatorid** | **int64** | The unique ID that is assigned to the part of the IPS, which generated the event. | 
**Id** | **int64** | The ID that is used to uniquely identify signatures. | 
**Label** | **string** | A descriptive label for the the signature. | 
**Cves** | **[]string** | The list of common vulnerabilites and exposures (CVEs). | 

## Methods

### NewSignature

`func NewSignature(generatorid int64, id int64, label string, cves []string, ) *Signature`

NewSignature instantiates a new Signature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureWithDefaults

`func NewSignatureWithDefaults() *Signature`

NewSignatureWithDefaults instantiates a new Signature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneratorid

`func (o *Signature) GetGeneratorid() int64`

GetGeneratorid returns the Generatorid field if non-nil, zero value otherwise.

### GetGeneratoridOk

`func (o *Signature) GetGeneratoridOk() (*int64, bool)`

GetGeneratoridOk returns a tuple with the Generatorid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratorid

`func (o *Signature) SetGeneratorid(v int64)`

SetGeneratorid sets Generatorid field to given value.


### GetId

`func (o *Signature) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Signature) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Signature) SetId(v int64)`

SetId sets Id field to given value.


### GetLabel

`func (o *Signature) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Signature) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Signature) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetCves

`func (o *Signature) GetCves() []string`

GetCves returns the Cves field if non-nil, zero value otherwise.

### GetCvesOk

`func (o *Signature) GetCvesOk() (*[]string, bool)`

GetCvesOk returns a tuple with the Cves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCves

`func (o *Signature) SetCves(v []string)`

SetCves sets Cves field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


