# ResponsesInformationInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | **string** | The HTTP response status code. | 
**Count** | **int64** | The total number of API requests. | 
**Requests** | [**[]RequestDetailsListInner**](RequestDetailsListInner.md) | The list of API request information. | 

## Methods

### NewResponsesInformationInner

`func NewResponsesInformationInner(statusCode string, count int64, requests []RequestDetailsListInner, ) *ResponsesInformationInner`

NewResponsesInformationInner instantiates a new ResponsesInformationInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesInformationInnerWithDefaults

`func NewResponsesInformationInnerWithDefaults() *ResponsesInformationInner`

NewResponsesInformationInnerWithDefaults instantiates a new ResponsesInformationInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *ResponsesInformationInner) GetStatusCode() string`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ResponsesInformationInner) GetStatusCodeOk() (*string, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ResponsesInformationInner) SetStatusCode(v string)`

SetStatusCode sets StatusCode field to given value.


### GetCount

`func (o *ResponsesInformationInner) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ResponsesInformationInner) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ResponsesInformationInner) SetCount(v int64)`

SetCount sets Count field to given value.


### GetRequests

`func (o *ResponsesInformationInner) GetRequests() []RequestDetailsListInner`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *ResponsesInformationInner) GetRequestsOk() (*[]RequestDetailsListInner, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *ResponsesInformationInner) SetRequests(v []RequestDetailsListInner)`

SetRequests sets Requests field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


