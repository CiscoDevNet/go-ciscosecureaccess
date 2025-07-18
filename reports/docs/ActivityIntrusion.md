# ActivityIntrusion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classification** | **string** | The category of attack detected by a rule that is part of a more general type of attack class, such as trojan-activity, attempted-user, and unknown. | 
**Date** | **string** | The date from the timestamp based on the timezone parameter. | 
**Destinationip** | **string** | The destination IP for the entry. | 
**Destinationport** | **int64** | The destination port for entry. | 
**Identities** | [**[]Identity**](Identity.md) | The list of identities for the entry. | 
**Protocol** | [**Protocol**](Protocol.md) |  | 
**Sessionid** | **int64** | The unique identifier of a session, which is used to group the correlated events between various services. | 
**Severity** | [**Severity**](Severity.md) |  | 
**Signature** | [**Signature**](Signature.md) |  | 
**Signaturelist** | [**SignatureList**](SignatureList.md) |  | 
**Sourceip** | **string** | The source IP for the entry. | 
**Sourceport** | **int64** | The source port for the entry. | 
**Time** | **string** | The time in 24-hour format based on the timezone parameter. | 
**Timestamp** | **int64** | The timestamp represented in milliseconds. | 
**Type** | **string** | The type of the request. An intrusion request always has type intrusion. | 
**Verdict** | [**VerdictDetected**](VerdictDetected.md) |  | 

## Methods

### NewActivityIntrusion

`func NewActivityIntrusion(classification string, date string, destinationip string, destinationport int64, identities []Identity, protocol Protocol, sessionid int64, severity Severity, signature Signature, signaturelist SignatureList, sourceip string, sourceport int64, time string, timestamp int64, type_ string, verdict VerdictDetected, ) *ActivityIntrusion`

NewActivityIntrusion instantiates a new ActivityIntrusion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityIntrusionWithDefaults

`func NewActivityIntrusionWithDefaults() *ActivityIntrusion`

NewActivityIntrusionWithDefaults instantiates a new ActivityIntrusion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassification

`func (o *ActivityIntrusion) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *ActivityIntrusion) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *ActivityIntrusion) SetClassification(v string)`

SetClassification sets Classification field to given value.


### GetDate

`func (o *ActivityIntrusion) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ActivityIntrusion) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ActivityIntrusion) SetDate(v string)`

SetDate sets Date field to given value.


### GetDestinationip

`func (o *ActivityIntrusion) GetDestinationip() string`

GetDestinationip returns the Destinationip field if non-nil, zero value otherwise.

### GetDestinationipOk

`func (o *ActivityIntrusion) GetDestinationipOk() (*string, bool)`

GetDestinationipOk returns a tuple with the Destinationip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationip

`func (o *ActivityIntrusion) SetDestinationip(v string)`

SetDestinationip sets Destinationip field to given value.


### GetDestinationport

`func (o *ActivityIntrusion) GetDestinationport() int64`

GetDestinationport returns the Destinationport field if non-nil, zero value otherwise.

### GetDestinationportOk

`func (o *ActivityIntrusion) GetDestinationportOk() (*int64, bool)`

GetDestinationportOk returns a tuple with the Destinationport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationport

`func (o *ActivityIntrusion) SetDestinationport(v int64)`

SetDestinationport sets Destinationport field to given value.


### GetIdentities

`func (o *ActivityIntrusion) GetIdentities() []Identity`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *ActivityIntrusion) GetIdentitiesOk() (*[]Identity, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *ActivityIntrusion) SetIdentities(v []Identity)`

SetIdentities sets Identities field to given value.


### GetProtocol

`func (o *ActivityIntrusion) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ActivityIntrusion) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ActivityIntrusion) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.


### GetSessionid

`func (o *ActivityIntrusion) GetSessionid() int64`

GetSessionid returns the Sessionid field if non-nil, zero value otherwise.

### GetSessionidOk

`func (o *ActivityIntrusion) GetSessionidOk() (*int64, bool)`

GetSessionidOk returns a tuple with the Sessionid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionid

`func (o *ActivityIntrusion) SetSessionid(v int64)`

SetSessionid sets Sessionid field to given value.


### GetSeverity

`func (o *ActivityIntrusion) GetSeverity() Severity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ActivityIntrusion) GetSeverityOk() (*Severity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ActivityIntrusion) SetSeverity(v Severity)`

SetSeverity sets Severity field to given value.


### GetSignature

`func (o *ActivityIntrusion) GetSignature() Signature`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *ActivityIntrusion) GetSignatureOk() (*Signature, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *ActivityIntrusion) SetSignature(v Signature)`

SetSignature sets Signature field to given value.


### GetSignaturelist

`func (o *ActivityIntrusion) GetSignaturelist() SignatureList`

GetSignaturelist returns the Signaturelist field if non-nil, zero value otherwise.

### GetSignaturelistOk

`func (o *ActivityIntrusion) GetSignaturelistOk() (*SignatureList, bool)`

GetSignaturelistOk returns a tuple with the Signaturelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignaturelist

`func (o *ActivityIntrusion) SetSignaturelist(v SignatureList)`

SetSignaturelist sets Signaturelist field to given value.


### GetSourceip

`func (o *ActivityIntrusion) GetSourceip() string`

GetSourceip returns the Sourceip field if non-nil, zero value otherwise.

### GetSourceipOk

`func (o *ActivityIntrusion) GetSourceipOk() (*string, bool)`

GetSourceipOk returns a tuple with the Sourceip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceip

`func (o *ActivityIntrusion) SetSourceip(v string)`

SetSourceip sets Sourceip field to given value.


### GetSourceport

`func (o *ActivityIntrusion) GetSourceport() int64`

GetSourceport returns the Sourceport field if non-nil, zero value otherwise.

### GetSourceportOk

`func (o *ActivityIntrusion) GetSourceportOk() (*int64, bool)`

GetSourceportOk returns a tuple with the Sourceport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceport

`func (o *ActivityIntrusion) SetSourceport(v int64)`

SetSourceport sets Sourceport field to given value.


### GetTime

`func (o *ActivityIntrusion) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ActivityIntrusion) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ActivityIntrusion) SetTime(v string)`

SetTime sets Time field to given value.


### GetTimestamp

`func (o *ActivityIntrusion) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ActivityIntrusion) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ActivityIntrusion) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetType

`func (o *ActivityIntrusion) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActivityIntrusion) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActivityIntrusion) SetType(v string)`

SetType sets Type field to given value.


### GetVerdict

`func (o *ActivityIntrusion) GetVerdict() VerdictDetected`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *ActivityIntrusion) GetVerdictOk() (*VerdictDetected, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *ActivityIntrusion) SetVerdict(v VerdictDetected)`

SetVerdict sets Verdict field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


