# ListInternalDomains500Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | Pointer to **int64** | HTTP status code | [optional] 
**Error** | Pointer to **string** | A brief description of the error | [optional] 
**Message** | Pointer to **string** | Detailed error message | [optional] 

## Methods

### NewListInternalDomains500Response

`func NewListInternalDomains500Response() *ListInternalDomains500Response`

NewListInternalDomains500Response instantiates a new ListInternalDomains500Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInternalDomains500ResponseWithDefaults

`func NewListInternalDomains500ResponseWithDefaults() *ListInternalDomains500Response`

NewListInternalDomains500ResponseWithDefaults instantiates a new ListInternalDomains500Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *ListInternalDomains500Response) GetStatusCode() int64`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ListInternalDomains500Response) GetStatusCodeOk() (*int64, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ListInternalDomains500Response) SetStatusCode(v int64)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ListInternalDomains500Response) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetError

`func (o *ListInternalDomains500Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ListInternalDomains500Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ListInternalDomains500Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ListInternalDomains500Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *ListInternalDomains500Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListInternalDomains500Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListInternalDomains500Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ListInternalDomains500Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


