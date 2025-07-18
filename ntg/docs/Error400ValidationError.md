# Error400ValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Error message explaining the reason for failure. | [optional] 
**RequestId** | Pointer to **string** | The ID of the request. | [optional] 

## Methods

### NewError400ValidationError

`func NewError400ValidationError() *Error400ValidationError`

NewError400ValidationError instantiates a new Error400ValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewError400ValidationErrorWithDefaults

`func NewError400ValidationErrorWithDefaults() *Error400ValidationError`

NewError400ValidationErrorWithDefaults instantiates a new Error400ValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *Error400ValidationError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Error400ValidationError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Error400ValidationError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Error400ValidationError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetRequestId

`func (o *Error400ValidationError) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *Error400ValidationError) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *Error400ValidationError) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *Error400ValidationError) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


