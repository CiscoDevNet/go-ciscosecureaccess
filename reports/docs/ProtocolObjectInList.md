# ProtocolObjectInList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the protocol. | 
**Name** | **string** | The name of the protocol. | 
**Description** | **string** | The description of the protocol. | 
**Events** | **int64** |  | 
**BlockedEvents** | **int64** |  | 
**FirstDetected** | **time.Time** |  | 
**LastDetected** | **time.Time** |  | 

## Methods

### NewProtocolObjectInList

`func NewProtocolObjectInList(id string, name string, description string, events int64, blockedEvents int64, firstDetected time.Time, lastDetected time.Time, ) *ProtocolObjectInList`

NewProtocolObjectInList instantiates a new ProtocolObjectInList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtocolObjectInListWithDefaults

`func NewProtocolObjectInListWithDefaults() *ProtocolObjectInList`

NewProtocolObjectInListWithDefaults instantiates a new ProtocolObjectInList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProtocolObjectInList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProtocolObjectInList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProtocolObjectInList) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ProtocolObjectInList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProtocolObjectInList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProtocolObjectInList) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProtocolObjectInList) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProtocolObjectInList) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProtocolObjectInList) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEvents

`func (o *ProtocolObjectInList) GetEvents() int64`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ProtocolObjectInList) GetEventsOk() (*int64, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ProtocolObjectInList) SetEvents(v int64)`

SetEvents sets Events field to given value.


### GetBlockedEvents

`func (o *ProtocolObjectInList) GetBlockedEvents() int64`

GetBlockedEvents returns the BlockedEvents field if non-nil, zero value otherwise.

### GetBlockedEventsOk

`func (o *ProtocolObjectInList) GetBlockedEventsOk() (*int64, bool)`

GetBlockedEventsOk returns a tuple with the BlockedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedEvents

`func (o *ProtocolObjectInList) SetBlockedEvents(v int64)`

SetBlockedEvents sets BlockedEvents field to given value.


### GetFirstDetected

`func (o *ProtocolObjectInList) GetFirstDetected() time.Time`

GetFirstDetected returns the FirstDetected field if non-nil, zero value otherwise.

### GetFirstDetectedOk

`func (o *ProtocolObjectInList) GetFirstDetectedOk() (*time.Time, bool)`

GetFirstDetectedOk returns a tuple with the FirstDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstDetected

`func (o *ProtocolObjectInList) SetFirstDetected(v time.Time)`

SetFirstDetected sets FirstDetected field to given value.


### GetLastDetected

`func (o *ProtocolObjectInList) GetLastDetected() time.Time`

GetLastDetected returns the LastDetected field if non-nil, zero value otherwise.

### GetLastDetectedOk

`func (o *ProtocolObjectInList) GetLastDetectedOk() (*time.Time, bool)`

GetLastDetectedOk returns a tuple with the LastDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDetected

`func (o *ProtocolObjectInList) SetLastDetected(v time.Time)`

SetLastDetected sets LastDetected field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


