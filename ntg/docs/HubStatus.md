# HubStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The status of the hub. | [optional] [readonly] 
**Time** | Pointer to **time.Time** | The date and time (timestamp) when the hub status was last reported. | [optional] [readonly] 

## Methods

### NewHubStatus

`func NewHubStatus() *HubStatus`

NewHubStatus instantiates a new HubStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHubStatusWithDefaults

`func NewHubStatusWithDefaults() *HubStatus`

NewHubStatusWithDefaults instantiates a new HubStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HubStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HubStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HubStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HubStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTime

`func (o *HubStatus) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *HubStatus) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *HubStatus) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *HubStatus) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


