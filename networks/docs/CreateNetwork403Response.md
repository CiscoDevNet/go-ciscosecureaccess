# CreateNetwork403Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | Pointer to **int64** | HTTP status code | [optional] 
**Error** | Pointer to **string** | Shorthand error message | [optional] 
**Message** | Pointer to **string** | Detailed error message | [optional] 

## Methods

### NewCreateNetwork403Response

`func NewCreateNetwork403Response() *CreateNetwork403Response`

NewCreateNetwork403Response instantiates a new CreateNetwork403Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetwork403ResponseWithDefaults

`func NewCreateNetwork403ResponseWithDefaults() *CreateNetwork403Response`

NewCreateNetwork403ResponseWithDefaults instantiates a new CreateNetwork403Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *CreateNetwork403Response) GetStatusCode() int64`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *CreateNetwork403Response) GetStatusCodeOk() (*int64, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *CreateNetwork403Response) SetStatusCode(v int64)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *CreateNetwork403Response) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetError

`func (o *CreateNetwork403Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CreateNetwork403Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CreateNetwork403Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CreateNetwork403Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *CreateNetwork403Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateNetwork403Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateNetwork403Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateNetwork403Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


