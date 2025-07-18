# GetActivities403Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | The description of the error. | [optional] 
**Code** | Pointer to **int64** | The status code of the error. | [optional] 

## Methods

### NewGetActivities403Response

`func NewGetActivities403Response() *GetActivities403Response`

NewGetActivities403Response instantiates a new GetActivities403Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetActivities403ResponseWithDefaults

`func NewGetActivities403ResponseWithDefaults() *GetActivities403Response`

NewGetActivities403ResponseWithDefaults instantiates a new GetActivities403Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetActivities403Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetActivities403Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetActivities403Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetActivities403Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCode

`func (o *GetActivities403Response) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetActivities403Response) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetActivities403Response) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetActivities403Response) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


