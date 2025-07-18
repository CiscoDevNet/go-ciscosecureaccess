# TopDnsQueryType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | [**Requests**](Requests.md) |  | 
**Querytype** | **string** | The type of the DNS query. | 

## Methods

### NewTopDnsQueryType

`func NewTopDnsQueryType(requests Requests, querytype string, ) *TopDnsQueryType`

NewTopDnsQueryType instantiates a new TopDnsQueryType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopDnsQueryTypeWithDefaults

`func NewTopDnsQueryTypeWithDefaults() *TopDnsQueryType`

NewTopDnsQueryTypeWithDefaults instantiates a new TopDnsQueryType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *TopDnsQueryType) GetRequests() Requests`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *TopDnsQueryType) GetRequestsOk() (*Requests, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *TopDnsQueryType) SetRequests(v Requests)`

SetRequests sets Requests field to given value.


### GetQuerytype

`func (o *TopDnsQueryType) GetQuerytype() string`

GetQuerytype returns the Querytype field if non-nil, zero value otherwise.

### GetQuerytypeOk

`func (o *TopDnsQueryType) GetQuerytypeOk() (*string, bool)`

GetQuerytypeOk returns a tuple with the Querytype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerytype

`func (o *TopDnsQueryType) SetQuerytype(v string)`

SetQuerytype sets Querytype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


