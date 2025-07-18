# SummaryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuccessfulRequests** | **int64** | The number of successful API requests. | 
**FailedRequests** | **int64** | The number of failed API requests. | 
**Total** | **int64** | The total number of API requests. | 

## Methods

### NewSummaryData

`func NewSummaryData(successfulRequests int64, failedRequests int64, total int64, ) *SummaryData`

NewSummaryData instantiates a new SummaryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryDataWithDefaults

`func NewSummaryDataWithDefaults() *SummaryData`

NewSummaryDataWithDefaults instantiates a new SummaryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccessfulRequests

`func (o *SummaryData) GetSuccessfulRequests() int64`

GetSuccessfulRequests returns the SuccessfulRequests field if non-nil, zero value otherwise.

### GetSuccessfulRequestsOk

`func (o *SummaryData) GetSuccessfulRequestsOk() (*int64, bool)`

GetSuccessfulRequestsOk returns a tuple with the SuccessfulRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulRequests

`func (o *SummaryData) SetSuccessfulRequests(v int64)`

SetSuccessfulRequests sets SuccessfulRequests field to given value.


### GetFailedRequests

`func (o *SummaryData) GetFailedRequests() int64`

GetFailedRequests returns the FailedRequests field if non-nil, zero value otherwise.

### GetFailedRequestsOk

`func (o *SummaryData) GetFailedRequestsOk() (*int64, bool)`

GetFailedRequestsOk returns a tuple with the FailedRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRequests

`func (o *SummaryData) SetFailedRequests(v int64)`

SetFailedRequests sets FailedRequests field to given value.


### GetTotal

`func (o *SummaryData) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SummaryData) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SummaryData) SetTotal(v int64)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


