# Error400BulkStates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Error message explaining the reason for failure. | [optional] 
**RequestId** | Pointer to **string** | The ID of the request. | [optional] 
**ValidationErrors** | Pointer to [**Error400BulkStatesValidationErrors**](Error400BulkStatesValidationErrors.md) |  | [optional] 

## Methods

### NewError400BulkStates

`func NewError400BulkStates() *Error400BulkStates`

NewError400BulkStates instantiates a new Error400BulkStates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewError400BulkStatesWithDefaults

`func NewError400BulkStatesWithDefaults() *Error400BulkStates`

NewError400BulkStatesWithDefaults instantiates a new Error400BulkStates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *Error400BulkStates) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Error400BulkStates) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Error400BulkStates) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Error400BulkStates) HasError() bool`

HasError returns a boolean if a field has been set.

### GetRequestId

`func (o *Error400BulkStates) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *Error400BulkStates) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *Error400BulkStates) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *Error400BulkStates) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetValidationErrors

`func (o *Error400BulkStates) GetValidationErrors() Error400BulkStatesValidationErrors`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *Error400BulkStates) GetValidationErrorsOk() (*Error400BulkStatesValidationErrors, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *Error400BulkStates) SetValidationErrors(v Error400BulkStatesValidationErrors)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *Error400BulkStates) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


