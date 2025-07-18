# Error400ListError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Error message explaining the reason for failure. | [optional] 
**RequestId** | Pointer to **string** | The ID of the request. | [optional] 
**ValidationErrors** | Pointer to [**Error400ListOneOfValidation**](Error400ListOneOfValidation.md) |  | [optional] 

## Methods

### NewError400ListError

`func NewError400ListError() *Error400ListError`

NewError400ListError instantiates a new Error400ListError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewError400ListErrorWithDefaults

`func NewError400ListErrorWithDefaults() *Error400ListError`

NewError400ListErrorWithDefaults instantiates a new Error400ListError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *Error400ListError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Error400ListError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Error400ListError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Error400ListError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetRequestId

`func (o *Error400ListError) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *Error400ListError) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *Error400ListError) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *Error400ListError) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetValidationErrors

`func (o *Error400ListError) GetValidationErrors() Error400ListOneOfValidation`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *Error400ListError) GetValidationErrorsOk() (*Error400ListOneOfValidation, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *Error400ListError) SetValidationErrors(v Error400ListOneOfValidation)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *Error400ListError) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


