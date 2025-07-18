# CdfwSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The type of traffic associated with the source. | [optional] 
**Events** | Pointer to **int64** | The number of CDFW L7 events for this app. | [optional] 
**BlockedEvents** | Pointer to **int64** | The number of CDFW L7 events that are blocked. | [optional] 

## Methods

### NewCdfwSource

`func NewCdfwSource() *CdfwSource`

NewCdfwSource instantiates a new CdfwSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdfwSourceWithDefaults

`func NewCdfwSourceWithDefaults() *CdfwSource`

NewCdfwSourceWithDefaults instantiates a new CdfwSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CdfwSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CdfwSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CdfwSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CdfwSource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEvents

`func (o *CdfwSource) GetEvents() int64`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CdfwSource) GetEventsOk() (*int64, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CdfwSource) SetEvents(v int64)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *CdfwSource) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetBlockedEvents

`func (o *CdfwSource) GetBlockedEvents() int64`

GetBlockedEvents returns the BlockedEvents field if non-nil, zero value otherwise.

### GetBlockedEventsOk

`func (o *CdfwSource) GetBlockedEventsOk() (*int64, bool)`

GetBlockedEventsOk returns a tuple with the BlockedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedEvents

`func (o *CdfwSource) SetBlockedEvents(v int64)`

SetBlockedEvents sets BlockedEvents field to given value.

### HasBlockedEvents

`func (o *CdfwSource) HasBlockedEvents() bool`

HasBlockedEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


