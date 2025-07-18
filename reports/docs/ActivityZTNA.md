# ActivityZTNA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allapplications** | Pointer to [**[]ActivityZTNAAllapplicationsInner**](ActivityZTNAAllapplicationsInner.md) | The list of private applications that are connected through Zero Trust Access. | [optional] 
**Date** | **string** | The date from the timestamp based on the timezone parameter. | 
**Destinationip** | Pointer to **string** | The resolved IP for Zero Trust Network Access (ZTNA) client-based events. | [optional] 
**Clientfirewall** | Pointer to **string** | The type of the firewall that is used, either &#x60;SYS&#x60; or &#x60;NONE&#x60;. | [optional] 
**Time** | **string** | The time in 24-hour format based on the timezone parameter. | 
**Type** | **string** | The type of the request. | 
**Diskencryption** | Pointer to **string** | Type of the disk encryption, either &#x60;SYS&#x60;, &#x60;THIRDPARTY&#x60;, &#x60;NONE&#x60;. | [optional] 
**Antimalwareagents** | Pointer to **[]string** | The list of anti-malware agents. | [optional] 
**Policy** | [**PolicyZTNA**](PolicyZTNA.md) |  | 
**Systempassword** | Pointer to **string** | The system password. | [optional] 
**Verdict** | **string** | The verdict for entry. | 
**Device** | Pointer to [**Device**](Device.md) |  | [optional] 
**Timestamp** | **int64** | The timestamp represented in milliseconds. | 
**Identities** | [**[]Identity**](Identity.md) | The list of identities for the entry. | 

## Methods

### NewActivityZTNA

`func NewActivityZTNA(date string, time string, type_ string, policy PolicyZTNA, verdict string, timestamp int64, identities []Identity, ) *ActivityZTNA`

NewActivityZTNA instantiates a new ActivityZTNA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityZTNAWithDefaults

`func NewActivityZTNAWithDefaults() *ActivityZTNA`

NewActivityZTNAWithDefaults instantiates a new ActivityZTNA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllapplications

`func (o *ActivityZTNA) GetAllapplications() []ActivityZTNAAllapplicationsInner`

GetAllapplications returns the Allapplications field if non-nil, zero value otherwise.

### GetAllapplicationsOk

`func (o *ActivityZTNA) GetAllapplicationsOk() (*[]ActivityZTNAAllapplicationsInner, bool)`

GetAllapplicationsOk returns a tuple with the Allapplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllapplications

`func (o *ActivityZTNA) SetAllapplications(v []ActivityZTNAAllapplicationsInner)`

SetAllapplications sets Allapplications field to given value.

### HasAllapplications

`func (o *ActivityZTNA) HasAllapplications() bool`

HasAllapplications returns a boolean if a field has been set.

### GetDate

`func (o *ActivityZTNA) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ActivityZTNA) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ActivityZTNA) SetDate(v string)`

SetDate sets Date field to given value.


### GetDestinationip

`func (o *ActivityZTNA) GetDestinationip() string`

GetDestinationip returns the Destinationip field if non-nil, zero value otherwise.

### GetDestinationipOk

`func (o *ActivityZTNA) GetDestinationipOk() (*string, bool)`

GetDestinationipOk returns a tuple with the Destinationip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationip

`func (o *ActivityZTNA) SetDestinationip(v string)`

SetDestinationip sets Destinationip field to given value.

### HasDestinationip

`func (o *ActivityZTNA) HasDestinationip() bool`

HasDestinationip returns a boolean if a field has been set.

### GetClientfirewall

`func (o *ActivityZTNA) GetClientfirewall() string`

GetClientfirewall returns the Clientfirewall field if non-nil, zero value otherwise.

### GetClientfirewallOk

`func (o *ActivityZTNA) GetClientfirewallOk() (*string, bool)`

GetClientfirewallOk returns a tuple with the Clientfirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientfirewall

`func (o *ActivityZTNA) SetClientfirewall(v string)`

SetClientfirewall sets Clientfirewall field to given value.

### HasClientfirewall

`func (o *ActivityZTNA) HasClientfirewall() bool`

HasClientfirewall returns a boolean if a field has been set.

### GetTime

`func (o *ActivityZTNA) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ActivityZTNA) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ActivityZTNA) SetTime(v string)`

SetTime sets Time field to given value.


### GetType

`func (o *ActivityZTNA) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActivityZTNA) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActivityZTNA) SetType(v string)`

SetType sets Type field to given value.


### GetDiskencryption

`func (o *ActivityZTNA) GetDiskencryption() string`

GetDiskencryption returns the Diskencryption field if non-nil, zero value otherwise.

### GetDiskencryptionOk

`func (o *ActivityZTNA) GetDiskencryptionOk() (*string, bool)`

GetDiskencryptionOk returns a tuple with the Diskencryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskencryption

`func (o *ActivityZTNA) SetDiskencryption(v string)`

SetDiskencryption sets Diskencryption field to given value.

### HasDiskencryption

`func (o *ActivityZTNA) HasDiskencryption() bool`

HasDiskencryption returns a boolean if a field has been set.

### GetAntimalwareagents

`func (o *ActivityZTNA) GetAntimalwareagents() []string`

GetAntimalwareagents returns the Antimalwareagents field if non-nil, zero value otherwise.

### GetAntimalwareagentsOk

`func (o *ActivityZTNA) GetAntimalwareagentsOk() (*[]string, bool)`

GetAntimalwareagentsOk returns a tuple with the Antimalwareagents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntimalwareagents

`func (o *ActivityZTNA) SetAntimalwareagents(v []string)`

SetAntimalwareagents sets Antimalwareagents field to given value.

### HasAntimalwareagents

`func (o *ActivityZTNA) HasAntimalwareagents() bool`

HasAntimalwareagents returns a boolean if a field has been set.

### GetPolicy

`func (o *ActivityZTNA) GetPolicy() PolicyZTNA`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *ActivityZTNA) GetPolicyOk() (*PolicyZTNA, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *ActivityZTNA) SetPolicy(v PolicyZTNA)`

SetPolicy sets Policy field to given value.


### GetSystempassword

`func (o *ActivityZTNA) GetSystempassword() string`

GetSystempassword returns the Systempassword field if non-nil, zero value otherwise.

### GetSystempasswordOk

`func (o *ActivityZTNA) GetSystempasswordOk() (*string, bool)`

GetSystempasswordOk returns a tuple with the Systempassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystempassword

`func (o *ActivityZTNA) SetSystempassword(v string)`

SetSystempassword sets Systempassword field to given value.

### HasSystempassword

`func (o *ActivityZTNA) HasSystempassword() bool`

HasSystempassword returns a boolean if a field has been set.

### GetVerdict

`func (o *ActivityZTNA) GetVerdict() string`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *ActivityZTNA) GetVerdictOk() (*string, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *ActivityZTNA) SetVerdict(v string)`

SetVerdict sets Verdict field to given value.


### GetDevice

`func (o *ActivityZTNA) GetDevice() Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *ActivityZTNA) GetDeviceOk() (*Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *ActivityZTNA) SetDevice(v Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *ActivityZTNA) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetTimestamp

`func (o *ActivityZTNA) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ActivityZTNA) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ActivityZTNA) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetIdentities

`func (o *ActivityZTNA) GetIdentities() []Identity`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *ActivityZTNA) GetIdentitiesOk() (*[]Identity, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *ActivityZTNA) SetIdentities(v []Identity)`

SetIdentities sets Identities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


