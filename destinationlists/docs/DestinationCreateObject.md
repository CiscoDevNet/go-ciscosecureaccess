# DestinationCreateObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | **string** | A domain, URL, or IP. | 
**Type** | Pointer to [**ModelType**](ModelType.md) |  | [optional] 
**Comment** | Pointer to **string** | The comment about the destination. | [optional] 

## Methods

### NewDestinationCreateObject

`func NewDestinationCreateObject(destination string, ) *DestinationCreateObject`

NewDestinationCreateObject instantiates a new DestinationCreateObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationCreateObjectWithDefaults

`func NewDestinationCreateObjectWithDefaults() *DestinationCreateObject`

NewDestinationCreateObjectWithDefaults instantiates a new DestinationCreateObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *DestinationCreateObject) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *DestinationCreateObject) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *DestinationCreateObject) SetDestination(v string)`

SetDestination sets Destination field to given value.


### GetType

`func (o *DestinationCreateObject) GetType() ModelType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestinationCreateObject) GetTypeOk() (*ModelType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestinationCreateObject) SetType(v ModelType)`

SetType sets Type field to given value.

### HasType

`func (o *DestinationCreateObject) HasType() bool`

HasType returns a boolean if a field has been set.

### GetComment

`func (o *DestinationCreateObject) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DestinationCreateObject) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DestinationCreateObject) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DestinationCreateObject) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


