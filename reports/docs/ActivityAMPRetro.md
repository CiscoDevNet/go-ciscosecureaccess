# ActivityAMPRetro

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **int64** | The timestamp represented in seconds. | 
**Firstseenat** | **int64** | The date and time (a timestamp expressed in seconds) when the malware event was first recorded. | 
**Disposition** | **string** | The disposition for the entry. | 
**Score** | **int64** | The score for the entry. | 
**Hostname** | **string** | The hostname for the entry. | 
**Malwarename** | **string** | The name of the malware for the entry. | 
**Sha256** | **string** | The SHA256 hash for the entry. | 

## Methods

### NewActivityAMPRetro

`func NewActivityAMPRetro(timestamp int64, firstseenat int64, disposition string, score int64, hostname string, malwarename string, sha256 string, ) *ActivityAMPRetro`

NewActivityAMPRetro instantiates a new ActivityAMPRetro object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityAMPRetroWithDefaults

`func NewActivityAMPRetroWithDefaults() *ActivityAMPRetro`

NewActivityAMPRetroWithDefaults instantiates a new ActivityAMPRetro object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ActivityAMPRetro) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ActivityAMPRetro) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ActivityAMPRetro) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetFirstseenat

`func (o *ActivityAMPRetro) GetFirstseenat() int64`

GetFirstseenat returns the Firstseenat field if non-nil, zero value otherwise.

### GetFirstseenatOk

`func (o *ActivityAMPRetro) GetFirstseenatOk() (*int64, bool)`

GetFirstseenatOk returns a tuple with the Firstseenat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstseenat

`func (o *ActivityAMPRetro) SetFirstseenat(v int64)`

SetFirstseenat sets Firstseenat field to given value.


### GetDisposition

`func (o *ActivityAMPRetro) GetDisposition() string`

GetDisposition returns the Disposition field if non-nil, zero value otherwise.

### GetDispositionOk

`func (o *ActivityAMPRetro) GetDispositionOk() (*string, bool)`

GetDispositionOk returns a tuple with the Disposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisposition

`func (o *ActivityAMPRetro) SetDisposition(v string)`

SetDisposition sets Disposition field to given value.


### GetScore

`func (o *ActivityAMPRetro) GetScore() int64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ActivityAMPRetro) GetScoreOk() (*int64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ActivityAMPRetro) SetScore(v int64)`

SetScore sets Score field to given value.


### GetHostname

`func (o *ActivityAMPRetro) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ActivityAMPRetro) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ActivityAMPRetro) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetMalwarename

`func (o *ActivityAMPRetro) GetMalwarename() string`

GetMalwarename returns the Malwarename field if non-nil, zero value otherwise.

### GetMalwarenameOk

`func (o *ActivityAMPRetro) GetMalwarenameOk() (*string, bool)`

GetMalwarenameOk returns a tuple with the Malwarename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalwarename

`func (o *ActivityAMPRetro) SetMalwarename(v string)`

SetMalwarename sets Malwarename field to given value.


### GetSha256

`func (o *ActivityAMPRetro) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *ActivityAMPRetro) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *ActivityAMPRetro) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


