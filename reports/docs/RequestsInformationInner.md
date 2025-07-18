# RequestsInformationInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserAgent** | **string** | The name of the client program. | 
**Count** | **int64** | The total number of API requests. | 
**Requests** | [**[]RequestDetailsListInner**](RequestDetailsListInner.md) | The list of API request information. | 

## Methods

### NewRequestsInformationInner

`func NewRequestsInformationInner(userAgent string, count int64, requests []RequestDetailsListInner, ) *RequestsInformationInner`

NewRequestsInformationInner instantiates a new RequestsInformationInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestsInformationInnerWithDefaults

`func NewRequestsInformationInnerWithDefaults() *RequestsInformationInner`

NewRequestsInformationInnerWithDefaults instantiates a new RequestsInformationInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserAgent

`func (o *RequestsInformationInner) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *RequestsInformationInner) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *RequestsInformationInner) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.


### GetCount

`func (o *RequestsInformationInner) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RequestsInformationInner) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RequestsInformationInner) SetCount(v int64)`

SetCount sets Count field to given value.


### GetRequests

`func (o *RequestsInformationInner) GetRequests() []RequestDetailsListInner`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *RequestsInformationInner) GetRequestsOk() (*[]RequestDetailsListInner, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *RequestsInformationInner) SetRequests(v []RequestDetailsListInner)`

SetRequests sets Requests field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


