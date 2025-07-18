# ListConnectorGroups400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Error message explaining the reason for failure. | [optional] 
**RequestId** | Pointer to **string** | The ID of the request. | [optional] 
**ValidationErrors** | Pointer to [**ListConnectorGroups400ResponseValidationErrors**](ListConnectorGroups400ResponseValidationErrors.md) |  | [optional] 

## Methods

### NewListConnectorGroups400Response

`func NewListConnectorGroups400Response() *ListConnectorGroups400Response`

NewListConnectorGroups400Response instantiates a new ListConnectorGroups400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConnectorGroups400ResponseWithDefaults

`func NewListConnectorGroups400ResponseWithDefaults() *ListConnectorGroups400Response`

NewListConnectorGroups400ResponseWithDefaults instantiates a new ListConnectorGroups400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ListConnectorGroups400Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ListConnectorGroups400Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ListConnectorGroups400Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ListConnectorGroups400Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetRequestId

`func (o *ListConnectorGroups400Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ListConnectorGroups400Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ListConnectorGroups400Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ListConnectorGroups400Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetValidationErrors

`func (o *ListConnectorGroups400Response) GetValidationErrors() ListConnectorGroups400ResponseValidationErrors`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *ListConnectorGroups400Response) GetValidationErrorsOk() (*ListConnectorGroups400ResponseValidationErrors, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *ListConnectorGroups400Response) SetValidationErrors(v ListConnectorGroups400ResponseValidationErrors)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *ListConnectorGroups400Response) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


