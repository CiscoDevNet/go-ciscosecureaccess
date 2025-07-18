# CiscoAMP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | **int64** | The AMP score. | 
**Disposition** | **string** | The AMP disposition. | 
**Malware** | **string** | The AMP malware. | 

## Methods

### NewCiscoAMP

`func NewCiscoAMP(score int64, disposition string, malware string, ) *CiscoAMP`

NewCiscoAMP instantiates a new CiscoAMP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCiscoAMPWithDefaults

`func NewCiscoAMPWithDefaults() *CiscoAMP`

NewCiscoAMPWithDefaults instantiates a new CiscoAMP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *CiscoAMP) GetScore() int64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *CiscoAMP) GetScoreOk() (*int64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *CiscoAMP) SetScore(v int64)`

SetScore sets Score field to given value.


### GetDisposition

`func (o *CiscoAMP) GetDisposition() string`

GetDisposition returns the Disposition field if non-nil, zero value otherwise.

### GetDispositionOk

`func (o *CiscoAMP) GetDispositionOk() (*string, bool)`

GetDispositionOk returns a tuple with the Disposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisposition

`func (o *CiscoAMP) SetDisposition(v string)`

SetDisposition sets Disposition field to given value.


### GetMalware

`func (o *CiscoAMP) GetMalware() string`

GetMalware returns the Malware field if non-nil, zero value otherwise.

### GetMalwareOk

`func (o *CiscoAMP) GetMalwareOk() (*string, bool)`

GetMalwareOk returns a tuple with the Malware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalware

`func (o *CiscoAMP) SetMalware(v string)`

SetMalware sets Malware field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


