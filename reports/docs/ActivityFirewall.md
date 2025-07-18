# ActivityFirewall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | The date from the timestamp based on the timezone parameter. | 
**Destinationip** | **string** | The destination IP for the entry. | 
**Sourceip** | **string** | The source IP for the entry. | 
**Sourceport** | **int64** | The source port for the entry. | 
**Destinationport** | **int64** | The destination port for entry. | 
**Categories** | Pointer to [**[]Category**](Category.md) | The list of categories. | [optional] 
**Verdict** | [**Verdict**](Verdict.md) |  | 
**Time** | **string** | The time in 24-hour format based on the timezone parameter. | 
**Timestamp** | **int64** | The timestamp represented in milliseconds. | 
**Identities** | [**[]Identity**](Identity.md) | The list of identities for the entry. | 
**Protocol** | [**Protocol**](Protocol.md) |  | 
**Rule** | [**Rule**](Rule.md) |  | 
**Type** | **string** | The type of the request. A firewall request always has type firewall. | 
**Allapplications** | [**[]Application**](Application.md) | The list of applications for the entry. | 
**Applicationprotocols** | [**[]FirewallApplication**](FirewallApplication.md) | A list of firewall application protocols. | 
**Direction** | **string** | The direction of the packet. It is destined either towards the internet or to the customer&#39;s network. | 
**Packetsize** | **int64** | The size of the packet that was received. | 

## Methods

### NewActivityFirewall

`func NewActivityFirewall(date string, destinationip string, sourceip string, sourceport int64, destinationport int64, verdict Verdict, time string, timestamp int64, identities []Identity, protocol Protocol, rule Rule, type_ string, allapplications []Application, applicationprotocols []FirewallApplication, direction string, packetsize int64, ) *ActivityFirewall`

NewActivityFirewall instantiates a new ActivityFirewall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityFirewallWithDefaults

`func NewActivityFirewallWithDefaults() *ActivityFirewall`

NewActivityFirewallWithDefaults instantiates a new ActivityFirewall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *ActivityFirewall) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ActivityFirewall) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ActivityFirewall) SetDate(v string)`

SetDate sets Date field to given value.


### GetDestinationip

`func (o *ActivityFirewall) GetDestinationip() string`

GetDestinationip returns the Destinationip field if non-nil, zero value otherwise.

### GetDestinationipOk

`func (o *ActivityFirewall) GetDestinationipOk() (*string, bool)`

GetDestinationipOk returns a tuple with the Destinationip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationip

`func (o *ActivityFirewall) SetDestinationip(v string)`

SetDestinationip sets Destinationip field to given value.


### GetSourceip

`func (o *ActivityFirewall) GetSourceip() string`

GetSourceip returns the Sourceip field if non-nil, zero value otherwise.

### GetSourceipOk

`func (o *ActivityFirewall) GetSourceipOk() (*string, bool)`

GetSourceipOk returns a tuple with the Sourceip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceip

`func (o *ActivityFirewall) SetSourceip(v string)`

SetSourceip sets Sourceip field to given value.


### GetSourceport

`func (o *ActivityFirewall) GetSourceport() int64`

GetSourceport returns the Sourceport field if non-nil, zero value otherwise.

### GetSourceportOk

`func (o *ActivityFirewall) GetSourceportOk() (*int64, bool)`

GetSourceportOk returns a tuple with the Sourceport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceport

`func (o *ActivityFirewall) SetSourceport(v int64)`

SetSourceport sets Sourceport field to given value.


### GetDestinationport

`func (o *ActivityFirewall) GetDestinationport() int64`

GetDestinationport returns the Destinationport field if non-nil, zero value otherwise.

### GetDestinationportOk

`func (o *ActivityFirewall) GetDestinationportOk() (*int64, bool)`

GetDestinationportOk returns a tuple with the Destinationport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationport

`func (o *ActivityFirewall) SetDestinationport(v int64)`

SetDestinationport sets Destinationport field to given value.


### GetCategories

`func (o *ActivityFirewall) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ActivityFirewall) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ActivityFirewall) SetCategories(v []Category)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ActivityFirewall) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetVerdict

`func (o *ActivityFirewall) GetVerdict() Verdict`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *ActivityFirewall) GetVerdictOk() (*Verdict, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *ActivityFirewall) SetVerdict(v Verdict)`

SetVerdict sets Verdict field to given value.


### GetTime

`func (o *ActivityFirewall) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ActivityFirewall) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ActivityFirewall) SetTime(v string)`

SetTime sets Time field to given value.


### GetTimestamp

`func (o *ActivityFirewall) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ActivityFirewall) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ActivityFirewall) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetIdentities

`func (o *ActivityFirewall) GetIdentities() []Identity`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *ActivityFirewall) GetIdentitiesOk() (*[]Identity, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *ActivityFirewall) SetIdentities(v []Identity)`

SetIdentities sets Identities field to given value.


### GetProtocol

`func (o *ActivityFirewall) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ActivityFirewall) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ActivityFirewall) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.


### GetRule

`func (o *ActivityFirewall) GetRule() Rule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *ActivityFirewall) GetRuleOk() (*Rule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *ActivityFirewall) SetRule(v Rule)`

SetRule sets Rule field to given value.


### GetType

`func (o *ActivityFirewall) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActivityFirewall) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActivityFirewall) SetType(v string)`

SetType sets Type field to given value.


### GetAllapplications

`func (o *ActivityFirewall) GetAllapplications() []Application`

GetAllapplications returns the Allapplications field if non-nil, zero value otherwise.

### GetAllapplicationsOk

`func (o *ActivityFirewall) GetAllapplicationsOk() (*[]Application, bool)`

GetAllapplicationsOk returns a tuple with the Allapplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllapplications

`func (o *ActivityFirewall) SetAllapplications(v []Application)`

SetAllapplications sets Allapplications field to given value.


### GetApplicationprotocols

`func (o *ActivityFirewall) GetApplicationprotocols() []FirewallApplication`

GetApplicationprotocols returns the Applicationprotocols field if non-nil, zero value otherwise.

### GetApplicationprotocolsOk

`func (o *ActivityFirewall) GetApplicationprotocolsOk() (*[]FirewallApplication, bool)`

GetApplicationprotocolsOk returns a tuple with the Applicationprotocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationprotocols

`func (o *ActivityFirewall) SetApplicationprotocols(v []FirewallApplication)`

SetApplicationprotocols sets Applicationprotocols field to given value.


### GetDirection

`func (o *ActivityFirewall) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ActivityFirewall) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ActivityFirewall) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetPacketsize

`func (o *ActivityFirewall) GetPacketsize() int64`

GetPacketsize returns the Packetsize field if non-nil, zero value otherwise.

### GetPacketsizeOk

`func (o *ActivityFirewall) GetPacketsizeOk() (*int64, bool)`

GetPacketsizeOk returns a tuple with the Packetsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketsize

`func (o *ActivityFirewall) SetPacketsize(v int64)`

SetPacketsize sets Packetsize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


