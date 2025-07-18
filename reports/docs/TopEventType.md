# TopEventType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Eventtype** | **string** | The type of the event. | 
**Count** | **int64** | The number of requests that match the event type (&#x60;eventtype&#x60;). | 

## Methods

### NewTopEventType

`func NewTopEventType(eventtype string, count int64, ) *TopEventType`

NewTopEventType instantiates a new TopEventType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopEventTypeWithDefaults

`func NewTopEventTypeWithDefaults() *TopEventType`

NewTopEventTypeWithDefaults instantiates a new TopEventType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventtype

`func (o *TopEventType) GetEventtype() string`

GetEventtype returns the Eventtype field if non-nil, zero value otherwise.

### GetEventtypeOk

`func (o *TopEventType) GetEventtypeOk() (*string, bool)`

GetEventtypeOk returns a tuple with the Eventtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventtype

`func (o *TopEventType) SetEventtype(v string)`

SetEventtype sets Eventtype field to given value.


### GetCount

`func (o *TopEventType) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TopEventType) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TopEventType) SetCount(v int64)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


