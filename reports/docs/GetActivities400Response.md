# GetActivities400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | The description of the error. | [optional] 
**Code** | Pointer to **int64** | The status code of the error. | [optional] 

## Methods

### NewGetActivities400Response

`func NewGetActivities400Response() *GetActivities400Response`

NewGetActivities400Response instantiates a new GetActivities400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetActivities400ResponseWithDefaults

`func NewGetActivities400ResponseWithDefaults() *GetActivities400Response`

NewGetActivities400ResponseWithDefaults instantiates a new GetActivities400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetActivities400Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetActivities400Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetActivities400Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetActivities400Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCode

`func (o *GetActivities400Response) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetActivities400Response) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetActivities400Response) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetActivities400Response) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


