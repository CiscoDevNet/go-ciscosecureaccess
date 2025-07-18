# ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to **string** | The information about the error condition. | [optional] 
**RequiredParameters** | Pointer to **[]string** | The list of query parameters that you must include in the API request. | [optional] 

## Methods

### NewResponseData

`func NewResponseData() *ResponseData`

NewResponseData instantiates a new ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDataWithDefaults

`func NewResponseDataWithDefaults() *ResponseData`

NewResponseDataWithDefaults instantiates a new ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *ResponseData) GetErrors() string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ResponseData) GetErrorsOk() (*string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ResponseData) SetErrors(v string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ResponseData) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetRequiredParameters

`func (o *ResponseData) GetRequiredParameters() []string`

GetRequiredParameters returns the RequiredParameters field if non-nil, zero value otherwise.

### GetRequiredParametersOk

`func (o *ResponseData) GetRequiredParametersOk() (*[]string, bool)`

GetRequiredParametersOk returns a tuple with the RequiredParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredParameters

`func (o *ResponseData) SetRequiredParameters(v []string)`

SetRequiredParameters sets RequiredParameters field to given value.

### HasRequiredParameters

`func (o *ResponseData) HasRequiredParameters() bool`

HasRequiredParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


