# PatchConnector400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Error message explaining the reason for failure. | [optional] 
**RequestId** | Pointer to **string** | The ID of the request. | [optional] 
**ValidationErrors** | Pointer to [**PatchConnector400ResponseValidationErrors**](PatchConnector400ResponseValidationErrors.md) |  | [optional] 

## Methods

### NewPatchConnector400Response

`func NewPatchConnector400Response() *PatchConnector400Response`

NewPatchConnector400Response instantiates a new PatchConnector400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchConnector400ResponseWithDefaults

`func NewPatchConnector400ResponseWithDefaults() *PatchConnector400Response`

NewPatchConnector400ResponseWithDefaults instantiates a new PatchConnector400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *PatchConnector400Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PatchConnector400Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PatchConnector400Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *PatchConnector400Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetRequestId

`func (o *PatchConnector400Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *PatchConnector400Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *PatchConnector400Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *PatchConnector400Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetValidationErrors

`func (o *PatchConnector400Response) GetValidationErrors() PatchConnector400ResponseValidationErrors`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *PatchConnector400Response) GetValidationErrorsOk() (*PatchConnector400ResponseValidationErrors, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *PatchConnector400Response) SetValidationErrors(v PatchConnector400ResponseValidationErrors)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *PatchConnector400Response) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


