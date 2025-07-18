# RequestCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | Pointer to **int64** | The total number of requests. | [optional] 
**Allowedrequests** | Pointer to **int64** | The number of requests that were allowed. | [optional] 
**Blockedrequests** | Pointer to **int64** | The number of requests that were blocked. | [optional] 

## Methods

### NewRequestCounts

`func NewRequestCounts() *RequestCounts`

NewRequestCounts instantiates a new RequestCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestCountsWithDefaults

`func NewRequestCountsWithDefaults() *RequestCounts`

NewRequestCountsWithDefaults instantiates a new RequestCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *RequestCounts) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *RequestCounts) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *RequestCounts) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *RequestCounts) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetAllowedrequests

`func (o *RequestCounts) GetAllowedrequests() int64`

GetAllowedrequests returns the Allowedrequests field if non-nil, zero value otherwise.

### GetAllowedrequestsOk

`func (o *RequestCounts) GetAllowedrequestsOk() (*int64, bool)`

GetAllowedrequestsOk returns a tuple with the Allowedrequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedrequests

`func (o *RequestCounts) SetAllowedrequests(v int64)`

SetAllowedrequests sets Allowedrequests field to given value.

### HasAllowedrequests

`func (o *RequestCounts) HasAllowedrequests() bool`

HasAllowedrequests returns a boolean if a field has been set.

### GetBlockedrequests

`func (o *RequestCounts) GetBlockedrequests() int64`

GetBlockedrequests returns the Blockedrequests field if non-nil, zero value otherwise.

### GetBlockedrequestsOk

`func (o *RequestCounts) GetBlockedrequestsOk() (*int64, bool)`

GetBlockedrequestsOk returns a tuple with the Blockedrequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedrequests

`func (o *RequestCounts) SetBlockedrequests(v int64)`

SetBlockedrequests sets Blockedrequests field to given value.

### HasBlockedrequests

`func (o *RequestCounts) HasBlockedrequests() bool`

HasBlockedrequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


