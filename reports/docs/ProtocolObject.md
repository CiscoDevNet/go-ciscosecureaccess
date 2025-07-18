# ProtocolObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the protocol. | [optional] 
**Name** | Pointer to **string** | The name of the protocol. | [optional] 
**Description** | Pointer to **string** | The description of the protocol. | [optional] 
**IdentitiesCount** | Pointer to **int64** | The number of identities. | [optional] 
**Events** | Pointer to **int64** | The number of identity events. | [optional] 
**BlockedEvents** | Pointer to **int64** | The number of blocked events for the identities. | [optional] 
**FirstDetected** | Pointer to **time.Time** | The date when the protocol was first detected for the identities. | [optional] 
**LastDetected** | Pointer to **time.Time** | The date when the protocol was last detected for the identities. | [optional] 

## Methods

### NewProtocolObject

`func NewProtocolObject() *ProtocolObject`

NewProtocolObject instantiates a new ProtocolObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtocolObjectWithDefaults

`func NewProtocolObjectWithDefaults() *ProtocolObject`

NewProtocolObjectWithDefaults instantiates a new ProtocolObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProtocolObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProtocolObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProtocolObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProtocolObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProtocolObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProtocolObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProtocolObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProtocolObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ProtocolObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProtocolObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProtocolObject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProtocolObject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIdentitiesCount

`func (o *ProtocolObject) GetIdentitiesCount() int64`

GetIdentitiesCount returns the IdentitiesCount field if non-nil, zero value otherwise.

### GetIdentitiesCountOk

`func (o *ProtocolObject) GetIdentitiesCountOk() (*int64, bool)`

GetIdentitiesCountOk returns a tuple with the IdentitiesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentitiesCount

`func (o *ProtocolObject) SetIdentitiesCount(v int64)`

SetIdentitiesCount sets IdentitiesCount field to given value.

### HasIdentitiesCount

`func (o *ProtocolObject) HasIdentitiesCount() bool`

HasIdentitiesCount returns a boolean if a field has been set.

### GetEvents

`func (o *ProtocolObject) GetEvents() int64`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ProtocolObject) GetEventsOk() (*int64, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ProtocolObject) SetEvents(v int64)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ProtocolObject) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetBlockedEvents

`func (o *ProtocolObject) GetBlockedEvents() int64`

GetBlockedEvents returns the BlockedEvents field if non-nil, zero value otherwise.

### GetBlockedEventsOk

`func (o *ProtocolObject) GetBlockedEventsOk() (*int64, bool)`

GetBlockedEventsOk returns a tuple with the BlockedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedEvents

`func (o *ProtocolObject) SetBlockedEvents(v int64)`

SetBlockedEvents sets BlockedEvents field to given value.

### HasBlockedEvents

`func (o *ProtocolObject) HasBlockedEvents() bool`

HasBlockedEvents returns a boolean if a field has been set.

### GetFirstDetected

`func (o *ProtocolObject) GetFirstDetected() time.Time`

GetFirstDetected returns the FirstDetected field if non-nil, zero value otherwise.

### GetFirstDetectedOk

`func (o *ProtocolObject) GetFirstDetectedOk() (*time.Time, bool)`

GetFirstDetectedOk returns a tuple with the FirstDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstDetected

`func (o *ProtocolObject) SetFirstDetected(v time.Time)`

SetFirstDetected sets FirstDetected field to given value.

### HasFirstDetected

`func (o *ProtocolObject) HasFirstDetected() bool`

HasFirstDetected returns a boolean if a field has been set.

### GetLastDetected

`func (o *ProtocolObject) GetLastDetected() time.Time`

GetLastDetected returns the LastDetected field if non-nil, zero value otherwise.

### GetLastDetectedOk

`func (o *ProtocolObject) GetLastDetectedOk() (*time.Time, bool)`

GetLastDetectedOk returns a tuple with the LastDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDetected

`func (o *ProtocolObject) SetLastDetected(v time.Time)`

SetLastDetected sets LastDetected field to given value.

### HasLastDetected

`func (o *ProtocolObject) HasLastDetected() bool`

HasLastDetected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


