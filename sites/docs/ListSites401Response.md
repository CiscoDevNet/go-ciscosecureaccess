# ListSites401Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | Pointer to **int64** | HTTP status code | [optional] 
**Error** | Pointer to **string** | A brief description of the error | [optional] 
**Message** | Pointer to **string** | Detailed error message | [optional] 

## Methods

### NewListSites401Response

`func NewListSites401Response() *ListSites401Response`

NewListSites401Response instantiates a new ListSites401Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSites401ResponseWithDefaults

`func NewListSites401ResponseWithDefaults() *ListSites401Response`

NewListSites401ResponseWithDefaults instantiates a new ListSites401Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *ListSites401Response) GetStatusCode() int64`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ListSites401Response) GetStatusCodeOk() (*int64, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ListSites401Response) SetStatusCode(v int64)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ListSites401Response) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetError

`func (o *ListSites401Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ListSites401Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ListSites401Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ListSites401Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *ListSites401Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListSites401Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListSites401Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ListSites401Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


