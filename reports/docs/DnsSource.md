# DnsSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The type of traffic associated with the source. | [optional] 
**Requests** | Pointer to **int64** | The number of DNS requests for this app. | [optional] 
**BlockedRequests** | Pointer to **int64** | The number of DNS requests blocked by Secure Access, based on policy configurations. | [optional] 

## Methods

### NewDnsSource

`func NewDnsSource() *DnsSource`

NewDnsSource instantiates a new DnsSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsSourceWithDefaults

`func NewDnsSourceWithDefaults() *DnsSource`

NewDnsSourceWithDefaults instantiates a new DnsSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DnsSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnsSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnsSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DnsSource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequests

`func (o *DnsSource) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *DnsSource) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *DnsSource) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *DnsSource) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetBlockedRequests

`func (o *DnsSource) GetBlockedRequests() int64`

GetBlockedRequests returns the BlockedRequests field if non-nil, zero value otherwise.

### GetBlockedRequestsOk

`func (o *DnsSource) GetBlockedRequestsOk() (*int64, bool)`

GetBlockedRequestsOk returns a tuple with the BlockedRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedRequests

`func (o *DnsSource) SetBlockedRequests(v int64)`

SetBlockedRequests sets BlockedRequests field to given value.

### HasBlockedRequests

`func (o *DnsSource) HasBlockedRequests() bool`

HasBlockedRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


