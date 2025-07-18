# HubState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **time.Time** | The date and time (UTC time, with milliseconds) when the state event record was generated. | [readonly] 
**Status** | **string** | The high-level status of the Hub: * UP - The hub is active. * DOWN - The hub is inactive. * UNKNOWN - The current status is unknown and pending updated information.  | [readonly] 

## Methods

### NewHubState

`func NewHubState(time time.Time, status string, ) *HubState`

NewHubState instantiates a new HubState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHubStateWithDefaults

`func NewHubStateWithDefaults() *HubState`

NewHubStateWithDefaults instantiates a new HubState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *HubState) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *HubState) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *HubState) SetTime(v time.Time)`

SetTime sets Time field to given value.


### GetStatus

`func (o *HubState) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HubState) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HubState) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


