# ActivityDns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Externalip** | **string** | The external IP for the entry. | 
**Internalip** | **string** | The internal IP for the entry. | 
**Policycategories** | [**[]Category**](Category.md) | The list of the policy categories. | 
**Categories** | [**[]Category**](Category.md) | The list of categories. | 
**Verdict** | [**Verdict**](Verdict.md) |  | 
**Domain** | **string** | The domain name for the entry. | 
**Timestamp** | **int64** | The timestamp represented in milliseconds. | 
**Identities** | [**[]Identity**](Identity.md) | The list of identities for the entry. | 
**Allapplications** | [**[]Application**](Application.md) | The list of all applications for the entry. | 
**Threats** | [**[]Threat**](Threat.md) | The list of threats for the entry. | 
**Type** | **string** | The type of the request. A DNS request always has the type dns. | 
**Querytype** | **string** | The type of the DNS request. | 
**Date** | **string** | The date from the timestamp based on the timezone parameter. | 
**Time** | **string** | The time in 24-hour format based on the timezone parameter. | 
**Returncode** | **int64** | The DNS return code for this request. | 
**Allowedapplications** | [**[]Application**](Application.md) | The list of allowed applications. | 
**Blockedapplications** | [**[]Application**](Application.md) | The list of blocked applications. | 

## Methods

### NewActivityDns

`func NewActivityDns(externalip string, internalip string, policycategories []Category, categories []Category, verdict Verdict, domain string, timestamp int64, identities []Identity, allapplications []Application, threats []Threat, type_ string, querytype string, date string, time string, returncode int64, allowedapplications []Application, blockedapplications []Application, ) *ActivityDns`

NewActivityDns instantiates a new ActivityDns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityDnsWithDefaults

`func NewActivityDnsWithDefaults() *ActivityDns`

NewActivityDnsWithDefaults instantiates a new ActivityDns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalip

`func (o *ActivityDns) GetExternalip() string`

GetExternalip returns the Externalip field if non-nil, zero value otherwise.

### GetExternalipOk

`func (o *ActivityDns) GetExternalipOk() (*string, bool)`

GetExternalipOk returns a tuple with the Externalip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalip

`func (o *ActivityDns) SetExternalip(v string)`

SetExternalip sets Externalip field to given value.


### GetInternalip

`func (o *ActivityDns) GetInternalip() string`

GetInternalip returns the Internalip field if non-nil, zero value otherwise.

### GetInternalipOk

`func (o *ActivityDns) GetInternalipOk() (*string, bool)`

GetInternalipOk returns a tuple with the Internalip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalip

`func (o *ActivityDns) SetInternalip(v string)`

SetInternalip sets Internalip field to given value.


### GetPolicycategories

`func (o *ActivityDns) GetPolicycategories() []Category`

GetPolicycategories returns the Policycategories field if non-nil, zero value otherwise.

### GetPolicycategoriesOk

`func (o *ActivityDns) GetPolicycategoriesOk() (*[]Category, bool)`

GetPolicycategoriesOk returns a tuple with the Policycategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicycategories

`func (o *ActivityDns) SetPolicycategories(v []Category)`

SetPolicycategories sets Policycategories field to given value.


### GetCategories

`func (o *ActivityDns) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ActivityDns) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ActivityDns) SetCategories(v []Category)`

SetCategories sets Categories field to given value.


### GetVerdict

`func (o *ActivityDns) GetVerdict() Verdict`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *ActivityDns) GetVerdictOk() (*Verdict, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *ActivityDns) SetVerdict(v Verdict)`

SetVerdict sets Verdict field to given value.


### GetDomain

`func (o *ActivityDns) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ActivityDns) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ActivityDns) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetTimestamp

`func (o *ActivityDns) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ActivityDns) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ActivityDns) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetIdentities

`func (o *ActivityDns) GetIdentities() []Identity`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *ActivityDns) GetIdentitiesOk() (*[]Identity, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *ActivityDns) SetIdentities(v []Identity)`

SetIdentities sets Identities field to given value.


### GetAllapplications

`func (o *ActivityDns) GetAllapplications() []Application`

GetAllapplications returns the Allapplications field if non-nil, zero value otherwise.

### GetAllapplicationsOk

`func (o *ActivityDns) GetAllapplicationsOk() (*[]Application, bool)`

GetAllapplicationsOk returns a tuple with the Allapplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllapplications

`func (o *ActivityDns) SetAllapplications(v []Application)`

SetAllapplications sets Allapplications field to given value.


### GetThreats

`func (o *ActivityDns) GetThreats() []Threat`

GetThreats returns the Threats field if non-nil, zero value otherwise.

### GetThreatsOk

`func (o *ActivityDns) GetThreatsOk() (*[]Threat, bool)`

GetThreatsOk returns a tuple with the Threats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreats

`func (o *ActivityDns) SetThreats(v []Threat)`

SetThreats sets Threats field to given value.


### GetType

`func (o *ActivityDns) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActivityDns) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActivityDns) SetType(v string)`

SetType sets Type field to given value.


### GetQuerytype

`func (o *ActivityDns) GetQuerytype() string`

GetQuerytype returns the Querytype field if non-nil, zero value otherwise.

### GetQuerytypeOk

`func (o *ActivityDns) GetQuerytypeOk() (*string, bool)`

GetQuerytypeOk returns a tuple with the Querytype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerytype

`func (o *ActivityDns) SetQuerytype(v string)`

SetQuerytype sets Querytype field to given value.


### GetDate

`func (o *ActivityDns) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ActivityDns) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ActivityDns) SetDate(v string)`

SetDate sets Date field to given value.


### GetTime

`func (o *ActivityDns) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ActivityDns) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ActivityDns) SetTime(v string)`

SetTime sets Time field to given value.


### GetReturncode

`func (o *ActivityDns) GetReturncode() int64`

GetReturncode returns the Returncode field if non-nil, zero value otherwise.

### GetReturncodeOk

`func (o *ActivityDns) GetReturncodeOk() (*int64, bool)`

GetReturncodeOk returns a tuple with the Returncode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturncode

`func (o *ActivityDns) SetReturncode(v int64)`

SetReturncode sets Returncode field to given value.


### GetAllowedapplications

`func (o *ActivityDns) GetAllowedapplications() []Application`

GetAllowedapplications returns the Allowedapplications field if non-nil, zero value otherwise.

### GetAllowedapplicationsOk

`func (o *ActivityDns) GetAllowedapplicationsOk() (*[]Application, bool)`

GetAllowedapplicationsOk returns a tuple with the Allowedapplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedapplications

`func (o *ActivityDns) SetAllowedapplications(v []Application)`

SetAllowedapplications sets Allowedapplications field to given value.


### GetBlockedapplications

`func (o *ActivityDns) GetBlockedapplications() []Application`

GetBlockedapplications returns the Blockedapplications field if non-nil, zero value otherwise.

### GetBlockedapplicationsOk

`func (o *ActivityDns) GetBlockedapplicationsOk() (*[]Application, bool)`

GetBlockedapplicationsOk returns a tuple with the Blockedapplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedapplications

`func (o *ActivityDns) SetBlockedapplications(v []Application)`

SetBlockedapplications sets Blockedapplications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


