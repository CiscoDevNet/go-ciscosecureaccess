# RequestAndConnectionCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | Pointer to [**Requests**](Requests.md) |  | [optional] 
**Allowedrequests** | Pointer to **int64** | The number of requests that were allowed. | [optional] 
**Blockedrequests** | Pointer to **int64** | The number of requests that were blocked. | [optional] 
**Connectionevents** | Pointer to **int64** | The number of remote access connection events. | [optional] 
**Disconnectionevents** | Pointer to **int64** | The number of remote access disconnection events. | [optional] 

## Methods

### NewRequestAndConnectionCounts

`func NewRequestAndConnectionCounts() *RequestAndConnectionCounts`

NewRequestAndConnectionCounts instantiates a new RequestAndConnectionCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestAndConnectionCountsWithDefaults

`func NewRequestAndConnectionCountsWithDefaults() *RequestAndConnectionCounts`

NewRequestAndConnectionCountsWithDefaults instantiates a new RequestAndConnectionCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *RequestAndConnectionCounts) GetRequests() Requests`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *RequestAndConnectionCounts) GetRequestsOk() (*Requests, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *RequestAndConnectionCounts) SetRequests(v Requests)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *RequestAndConnectionCounts) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetAllowedrequests

`func (o *RequestAndConnectionCounts) GetAllowedrequests() int64`

GetAllowedrequests returns the Allowedrequests field if non-nil, zero value otherwise.

### GetAllowedrequestsOk

`func (o *RequestAndConnectionCounts) GetAllowedrequestsOk() (*int64, bool)`

GetAllowedrequestsOk returns a tuple with the Allowedrequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedrequests

`func (o *RequestAndConnectionCounts) SetAllowedrequests(v int64)`

SetAllowedrequests sets Allowedrequests field to given value.

### HasAllowedrequests

`func (o *RequestAndConnectionCounts) HasAllowedrequests() bool`

HasAllowedrequests returns a boolean if a field has been set.

### GetBlockedrequests

`func (o *RequestAndConnectionCounts) GetBlockedrequests() int64`

GetBlockedrequests returns the Blockedrequests field if non-nil, zero value otherwise.

### GetBlockedrequestsOk

`func (o *RequestAndConnectionCounts) GetBlockedrequestsOk() (*int64, bool)`

GetBlockedrequestsOk returns a tuple with the Blockedrequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedrequests

`func (o *RequestAndConnectionCounts) SetBlockedrequests(v int64)`

SetBlockedrequests sets Blockedrequests field to given value.

### HasBlockedrequests

`func (o *RequestAndConnectionCounts) HasBlockedrequests() bool`

HasBlockedrequests returns a boolean if a field has been set.

### GetConnectionevents

`func (o *RequestAndConnectionCounts) GetConnectionevents() int64`

GetConnectionevents returns the Connectionevents field if non-nil, zero value otherwise.

### GetConnectioneventsOk

`func (o *RequestAndConnectionCounts) GetConnectioneventsOk() (*int64, bool)`

GetConnectioneventsOk returns a tuple with the Connectionevents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionevents

`func (o *RequestAndConnectionCounts) SetConnectionevents(v int64)`

SetConnectionevents sets Connectionevents field to given value.

### HasConnectionevents

`func (o *RequestAndConnectionCounts) HasConnectionevents() bool`

HasConnectionevents returns a boolean if a field has been set.

### GetDisconnectionevents

`func (o *RequestAndConnectionCounts) GetDisconnectionevents() int64`

GetDisconnectionevents returns the Disconnectionevents field if non-nil, zero value otherwise.

### GetDisconnectioneventsOk

`func (o *RequestAndConnectionCounts) GetDisconnectioneventsOk() (*int64, bool)`

GetDisconnectioneventsOk returns a tuple with the Disconnectionevents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectionevents

`func (o *RequestAndConnectionCounts) SetDisconnectionevents(v int64)`

SetDisconnectionevents sets Disconnectionevents field to given value.

### HasDisconnectionevents

`func (o *RequestAndConnectionCounts) HasDisconnectionevents() bool`

HasDisconnectionevents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


