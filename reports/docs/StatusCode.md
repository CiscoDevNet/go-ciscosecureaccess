# StatusCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | **int64** |  | 
**Body** | [**StatusCodeBody**](StatusCodeBody.md) |  | 

## Methods

### NewStatusCode

`func NewStatusCode(statusCode int64, body StatusCodeBody, ) *StatusCode`

NewStatusCode instantiates a new StatusCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusCodeWithDefaults

`func NewStatusCodeWithDefaults() *StatusCode`

NewStatusCodeWithDefaults instantiates a new StatusCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *StatusCode) GetStatusCode() int64`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *StatusCode) GetStatusCodeOk() (*int64, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *StatusCode) SetStatusCode(v int64)`

SetStatusCode sets StatusCode field to given value.


### GetBody

`func (o *StatusCode) GetBody() StatusCodeBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *StatusCode) GetBodyOk() (*StatusCodeBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *StatusCode) SetBody(v StatusCodeBody)`

SetBody sets Body field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


