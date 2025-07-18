# Requests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **string** | The date and time where to start reading in the collection. | 
**To** | **string** | The date and time where to stop reading in the collection. | 
**Count** | **int64** | The total number of API requests. | 
**Items** | [**[]RequestsInformationInner**](RequestsInformationInner.md) | The information about the API requests. | 

## Methods

### NewRequests

`func NewRequests(from string, to string, count int64, items []RequestsInformationInner, ) *Requests`

NewRequests instantiates a new Requests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestsWithDefaults

`func NewRequestsWithDefaults() *Requests`

NewRequestsWithDefaults instantiates a new Requests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *Requests) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Requests) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Requests) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *Requests) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Requests) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Requests) SetTo(v string)`

SetTo sets To field to given value.


### GetCount

`func (o *Requests) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Requests) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Requests) SetCount(v int64)`

SetCount sets Count field to given value.


### GetItems

`func (o *Requests) GetItems() []RequestsInformationInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Requests) GetItemsOk() (*[]RequestsInformationInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Requests) SetItems(v []RequestsInformationInner)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


