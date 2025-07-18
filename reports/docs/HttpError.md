# HttpError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of the error, either &#x60;CertificateError&#x60; or &#x60;TLSError&#x60;. | [optional] 
**Code** | Pointer to **int64** | The HTTP error code. | [optional] 
**Reason** | Pointer to **string** | The name of the error. | [optional] 
**Attributes** | Pointer to **map[string]interface{}** | The properties of the additional information for the error. | [optional] 

## Methods

### NewHttpError

`func NewHttpError() *HttpError`

NewHttpError instantiates a new HttpError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpErrorWithDefaults

`func NewHttpErrorWithDefaults() *HttpError`

NewHttpErrorWithDefaults instantiates a new HttpError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *HttpError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HttpError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HttpError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HttpError) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCode

`func (o *HttpError) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *HttpError) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *HttpError) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *HttpError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetReason

`func (o *HttpError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *HttpError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *HttpError) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *HttpError) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetAttributes

`func (o *HttpError) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *HttpError) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *HttpError) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *HttpError) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


