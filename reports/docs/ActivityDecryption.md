# ActivityDecryption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | The date from the timestamp based on the timezone parameter. | 
**Time** | **string** | The time in 24-hour format based on the timezone parameter. | 
**Type** | **string** | type of the request. Decryption | 
**DecryptAction** | **string** | Type of decryption action (Decrypt Inbound, Decrypt Outbound, Do Not Decrypt, Decrypt Error). | 
**Timestamp** | **int64** | The timestamp represented in milliseconds. | 
**Identities** | [**[]Identity**](Identity.md) | The list of identities for the entry. | 

## Methods

### NewActivityDecryption

`func NewActivityDecryption(date string, time string, type_ string, decryptAction string, timestamp int64, identities []Identity, ) *ActivityDecryption`

NewActivityDecryption instantiates a new ActivityDecryption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityDecryptionWithDefaults

`func NewActivityDecryptionWithDefaults() *ActivityDecryption`

NewActivityDecryptionWithDefaults instantiates a new ActivityDecryption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *ActivityDecryption) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ActivityDecryption) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ActivityDecryption) SetDate(v string)`

SetDate sets Date field to given value.


### GetTime

`func (o *ActivityDecryption) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ActivityDecryption) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ActivityDecryption) SetTime(v string)`

SetTime sets Time field to given value.


### GetType

`func (o *ActivityDecryption) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActivityDecryption) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActivityDecryption) SetType(v string)`

SetType sets Type field to given value.


### GetDecryptAction

`func (o *ActivityDecryption) GetDecryptAction() string`

GetDecryptAction returns the DecryptAction field if non-nil, zero value otherwise.

### GetDecryptActionOk

`func (o *ActivityDecryption) GetDecryptActionOk() (*string, bool)`

GetDecryptActionOk returns a tuple with the DecryptAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecryptAction

`func (o *ActivityDecryption) SetDecryptAction(v string)`

SetDecryptAction sets DecryptAction field to given value.


### GetTimestamp

`func (o *ActivityDecryption) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ActivityDecryption) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ActivityDecryption) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetIdentities

`func (o *ActivityDecryption) GetIdentities() []Identity`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *ActivityDecryption) GetIdentitiesOk() (*[]Identity, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *ActivityDecryption) SetIdentities(v []Identity)`

SetIdentities sets Identities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


