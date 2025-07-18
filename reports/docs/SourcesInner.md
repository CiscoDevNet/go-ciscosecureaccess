# SourcesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The type of traffic associated with the source. | [optional] 
**Requests** | Pointer to **int64** | The number of DNS requests for this app. | [optional] 
**BlockedRequests** | Pointer to **int64** | The number of DNS requests blocked by Secure Access, based on policy configurations. | [optional] 
**TotalTraffic** | Pointer to **int64** | The total number of bytes inbound and outbound for this app in your environment. | [optional] 
**BytesIn** | Pointer to **int64** | The number of bytes received (inbound). | [optional] 
**BytesOut** | Pointer to **int64** | The number of bytes sent (outbound). | [optional] 
**BlockedBytesOut** | Pointer to **int64** | The number of bytes of outbound traffic that are blocked. | [optional] 
**Events** | Pointer to **int64** | The number of CDFW L7 events for this app. | [optional] 
**BlockedEvents** | Pointer to **int64** | The number of CDFW L7 events that are blocked. | [optional] 

## Methods

### NewSourcesInner

`func NewSourcesInner() *SourcesInner`

NewSourcesInner instantiates a new SourcesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourcesInnerWithDefaults

`func NewSourcesInnerWithDefaults() *SourcesInner`

NewSourcesInnerWithDefaults instantiates a new SourcesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SourcesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourcesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourcesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SourcesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequests

`func (o *SourcesInner) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *SourcesInner) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *SourcesInner) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *SourcesInner) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetBlockedRequests

`func (o *SourcesInner) GetBlockedRequests() int64`

GetBlockedRequests returns the BlockedRequests field if non-nil, zero value otherwise.

### GetBlockedRequestsOk

`func (o *SourcesInner) GetBlockedRequestsOk() (*int64, bool)`

GetBlockedRequestsOk returns a tuple with the BlockedRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedRequests

`func (o *SourcesInner) SetBlockedRequests(v int64)`

SetBlockedRequests sets BlockedRequests field to given value.

### HasBlockedRequests

`func (o *SourcesInner) HasBlockedRequests() bool`

HasBlockedRequests returns a boolean if a field has been set.

### GetTotalTraffic

`func (o *SourcesInner) GetTotalTraffic() int64`

GetTotalTraffic returns the TotalTraffic field if non-nil, zero value otherwise.

### GetTotalTrafficOk

`func (o *SourcesInner) GetTotalTrafficOk() (*int64, bool)`

GetTotalTrafficOk returns a tuple with the TotalTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTraffic

`func (o *SourcesInner) SetTotalTraffic(v int64)`

SetTotalTraffic sets TotalTraffic field to given value.

### HasTotalTraffic

`func (o *SourcesInner) HasTotalTraffic() bool`

HasTotalTraffic returns a boolean if a field has been set.

### GetBytesIn

`func (o *SourcesInner) GetBytesIn() int64`

GetBytesIn returns the BytesIn field if non-nil, zero value otherwise.

### GetBytesInOk

`func (o *SourcesInner) GetBytesInOk() (*int64, bool)`

GetBytesInOk returns a tuple with the BytesIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesIn

`func (o *SourcesInner) SetBytesIn(v int64)`

SetBytesIn sets BytesIn field to given value.

### HasBytesIn

`func (o *SourcesInner) HasBytesIn() bool`

HasBytesIn returns a boolean if a field has been set.

### GetBytesOut

`func (o *SourcesInner) GetBytesOut() int64`

GetBytesOut returns the BytesOut field if non-nil, zero value otherwise.

### GetBytesOutOk

`func (o *SourcesInner) GetBytesOutOk() (*int64, bool)`

GetBytesOutOk returns a tuple with the BytesOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesOut

`func (o *SourcesInner) SetBytesOut(v int64)`

SetBytesOut sets BytesOut field to given value.

### HasBytesOut

`func (o *SourcesInner) HasBytesOut() bool`

HasBytesOut returns a boolean if a field has been set.

### GetBlockedBytesOut

`func (o *SourcesInner) GetBlockedBytesOut() int64`

GetBlockedBytesOut returns the BlockedBytesOut field if non-nil, zero value otherwise.

### GetBlockedBytesOutOk

`func (o *SourcesInner) GetBlockedBytesOutOk() (*int64, bool)`

GetBlockedBytesOutOk returns a tuple with the BlockedBytesOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedBytesOut

`func (o *SourcesInner) SetBlockedBytesOut(v int64)`

SetBlockedBytesOut sets BlockedBytesOut field to given value.

### HasBlockedBytesOut

`func (o *SourcesInner) HasBlockedBytesOut() bool`

HasBlockedBytesOut returns a boolean if a field has been set.

### GetEvents

`func (o *SourcesInner) GetEvents() int64`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SourcesInner) GetEventsOk() (*int64, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SourcesInner) SetEvents(v int64)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *SourcesInner) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetBlockedEvents

`func (o *SourcesInner) GetBlockedEvents() int64`

GetBlockedEvents returns the BlockedEvents field if non-nil, zero value otherwise.

### GetBlockedEventsOk

`func (o *SourcesInner) GetBlockedEventsOk() (*int64, bool)`

GetBlockedEventsOk returns a tuple with the BlockedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedEvents

`func (o *SourcesInner) SetBlockedEvents(v int64)`

SetBlockedEvents sets BlockedEvents field to given value.

### HasBlockedEvents

`func (o *SourcesInner) HasBlockedEvents() bool`

HasBlockedEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


