# ProtocolIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the identity. | [optional] 
**Name** | Pointer to **string** | The name of the identity. | [optional] 
**Events** | Pointer to **int64** | The number of events for the identity. | [optional] 
**BlockedEvents** | Pointer to **int64** | The number of blocked events for the identity. | [optional] 
**FirstDetected** | Pointer to **time.Time** | The date when the identity was first detected. | [optional] 
**LastDetected** | Pointer to **time.Time** | The date when the identity was last detected. | [optional] 

## Methods

### NewProtocolIdentity

`func NewProtocolIdentity() *ProtocolIdentity`

NewProtocolIdentity instantiates a new ProtocolIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtocolIdentityWithDefaults

`func NewProtocolIdentityWithDefaults() *ProtocolIdentity`

NewProtocolIdentityWithDefaults instantiates a new ProtocolIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProtocolIdentity) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProtocolIdentity) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProtocolIdentity) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ProtocolIdentity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProtocolIdentity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProtocolIdentity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProtocolIdentity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProtocolIdentity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEvents

`func (o *ProtocolIdentity) GetEvents() int64`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ProtocolIdentity) GetEventsOk() (*int64, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ProtocolIdentity) SetEvents(v int64)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ProtocolIdentity) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetBlockedEvents

`func (o *ProtocolIdentity) GetBlockedEvents() int64`

GetBlockedEvents returns the BlockedEvents field if non-nil, zero value otherwise.

### GetBlockedEventsOk

`func (o *ProtocolIdentity) GetBlockedEventsOk() (*int64, bool)`

GetBlockedEventsOk returns a tuple with the BlockedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedEvents

`func (o *ProtocolIdentity) SetBlockedEvents(v int64)`

SetBlockedEvents sets BlockedEvents field to given value.

### HasBlockedEvents

`func (o *ProtocolIdentity) HasBlockedEvents() bool`

HasBlockedEvents returns a boolean if a field has been set.

### GetFirstDetected

`func (o *ProtocolIdentity) GetFirstDetected() time.Time`

GetFirstDetected returns the FirstDetected field if non-nil, zero value otherwise.

### GetFirstDetectedOk

`func (o *ProtocolIdentity) GetFirstDetectedOk() (*time.Time, bool)`

GetFirstDetectedOk returns a tuple with the FirstDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstDetected

`func (o *ProtocolIdentity) SetFirstDetected(v time.Time)`

SetFirstDetected sets FirstDetected field to given value.

### HasFirstDetected

`func (o *ProtocolIdentity) HasFirstDetected() bool`

HasFirstDetected returns a boolean if a field has been set.

### GetLastDetected

`func (o *ProtocolIdentity) GetLastDetected() time.Time`

GetLastDetected returns the LastDetected field if non-nil, zero value otherwise.

### GetLastDetectedOk

`func (o *ProtocolIdentity) GetLastDetectedOk() (*time.Time, bool)`

GetLastDetectedOk returns a tuple with the LastDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDetected

`func (o *ProtocolIdentity) SetLastDetected(v time.Time)`

SetLastDetected sets LastDetected field to given value.

### HasLastDetected

`func (o *ProtocolIdentity) HasLastDetected() bool`

HasLastDetected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


