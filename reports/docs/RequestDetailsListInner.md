# RequestDetailsListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | The path of the API endpoint. | 
**Verb** | **string** | The name of the API operation. | 
**Count** | **int64** | The number of requests to the API endpoint. | 

## Methods

### NewRequestDetailsListInner

`func NewRequestDetailsListInner(path string, verb string, count int64, ) *RequestDetailsListInner`

NewRequestDetailsListInner instantiates a new RequestDetailsListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestDetailsListInnerWithDefaults

`func NewRequestDetailsListInnerWithDefaults() *RequestDetailsListInner`

NewRequestDetailsListInnerWithDefaults instantiates a new RequestDetailsListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *RequestDetailsListInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RequestDetailsListInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RequestDetailsListInner) SetPath(v string)`

SetPath sets Path field to given value.


### GetVerb

`func (o *RequestDetailsListInner) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *RequestDetailsListInner) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *RequestDetailsListInner) SetVerb(v string)`

SetVerb sets Verb field to given value.


### GetCount

`func (o *RequestDetailsListInner) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RequestDetailsListInner) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RequestDetailsListInner) SetCount(v int64)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


