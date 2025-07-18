# ActivityIP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | The date from the timestamp based on the timezone parameter. | 
**Destinationip** | **string** | The destination IP for the entry. | 
**Sourceip** | **string** | The source IP for the entry. | 
**Sourceport** | **int64** | The source port for the entry. | 
**Destinationport** | **int64** | The destination port for entry. | 
**Verdict** | [**Verdict**](Verdict.md) |  | 
**Time** | **string** | The time in 24-hour format based on the timezone parameter. | 
**Timestamp** | **int64** | The timestamp represented in milliseconds. | 
**Identities** | [**[]Identity**](Identity.md) | The list of identities for the entry. | 
**Categories** | [**[]Category**](Category.md) | The list of categories. | 
**Type** | **string** | The type of the request. An IP request always has type &#x60;ip&#x60;. | 

## Methods

### NewActivityIP

`func NewActivityIP(date string, destinationip string, sourceip string, sourceport int64, destinationport int64, verdict Verdict, time string, timestamp int64, identities []Identity, categories []Category, type_ string, ) *ActivityIP`

NewActivityIP instantiates a new ActivityIP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityIPWithDefaults

`func NewActivityIPWithDefaults() *ActivityIP`

NewActivityIPWithDefaults instantiates a new ActivityIP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *ActivityIP) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ActivityIP) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ActivityIP) SetDate(v string)`

SetDate sets Date field to given value.


### GetDestinationip

`func (o *ActivityIP) GetDestinationip() string`

GetDestinationip returns the Destinationip field if non-nil, zero value otherwise.

### GetDestinationipOk

`func (o *ActivityIP) GetDestinationipOk() (*string, bool)`

GetDestinationipOk returns a tuple with the Destinationip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationip

`func (o *ActivityIP) SetDestinationip(v string)`

SetDestinationip sets Destinationip field to given value.


### GetSourceip

`func (o *ActivityIP) GetSourceip() string`

GetSourceip returns the Sourceip field if non-nil, zero value otherwise.

### GetSourceipOk

`func (o *ActivityIP) GetSourceipOk() (*string, bool)`

GetSourceipOk returns a tuple with the Sourceip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceip

`func (o *ActivityIP) SetSourceip(v string)`

SetSourceip sets Sourceip field to given value.


### GetSourceport

`func (o *ActivityIP) GetSourceport() int64`

GetSourceport returns the Sourceport field if non-nil, zero value otherwise.

### GetSourceportOk

`func (o *ActivityIP) GetSourceportOk() (*int64, bool)`

GetSourceportOk returns a tuple with the Sourceport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceport

`func (o *ActivityIP) SetSourceport(v int64)`

SetSourceport sets Sourceport field to given value.


### GetDestinationport

`func (o *ActivityIP) GetDestinationport() int64`

GetDestinationport returns the Destinationport field if non-nil, zero value otherwise.

### GetDestinationportOk

`func (o *ActivityIP) GetDestinationportOk() (*int64, bool)`

GetDestinationportOk returns a tuple with the Destinationport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationport

`func (o *ActivityIP) SetDestinationport(v int64)`

SetDestinationport sets Destinationport field to given value.


### GetVerdict

`func (o *ActivityIP) GetVerdict() Verdict`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *ActivityIP) GetVerdictOk() (*Verdict, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *ActivityIP) SetVerdict(v Verdict)`

SetVerdict sets Verdict field to given value.


### GetTime

`func (o *ActivityIP) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ActivityIP) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ActivityIP) SetTime(v string)`

SetTime sets Time field to given value.


### GetTimestamp

`func (o *ActivityIP) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ActivityIP) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ActivityIP) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetIdentities

`func (o *ActivityIP) GetIdentities() []Identity`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *ActivityIP) GetIdentitiesOk() (*[]Identity, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *ActivityIP) SetIdentities(v []Identity)`

SetIdentities sets Identities field to given value.


### GetCategories

`func (o *ActivityIP) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ActivityIP) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ActivityIP) SetCategories(v []Category)`

SetCategories sets Categories field to given value.


### GetType

`func (o *ActivityIP) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActivityIP) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActivityIP) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


