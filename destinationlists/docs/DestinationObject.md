# DestinationObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the destination. | 
**Destination** | **string** | A domain, URL, or IP. | 
**Type** | [**ModelType**](ModelType.md) |  | 
**Comment** | Pointer to **string** | The comment about the destination. | [optional] 
**CreatedAt** | **string** | The date and time when the destination list was created. | 

## Methods

### NewDestinationObject

`func NewDestinationObject(id string, destination string, type_ ModelType, createdAt string, ) *DestinationObject`

NewDestinationObject instantiates a new DestinationObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationObjectWithDefaults

`func NewDestinationObjectWithDefaults() *DestinationObject`

NewDestinationObjectWithDefaults instantiates a new DestinationObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DestinationObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DestinationObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DestinationObject) SetId(v string)`

SetId sets Id field to given value.


### GetDestination

`func (o *DestinationObject) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *DestinationObject) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *DestinationObject) SetDestination(v string)`

SetDestination sets Destination field to given value.


### GetType

`func (o *DestinationObject) GetType() ModelType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestinationObject) GetTypeOk() (*ModelType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestinationObject) SetType(v ModelType)`

SetType sets Type field to given value.


### GetComment

`func (o *DestinationObject) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DestinationObject) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DestinationObject) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DestinationObject) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DestinationObject) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DestinationObject) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DestinationObject) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


