# RemoteAccessIdentitiesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the remote access endpoint. | [optional] 
**Type** | Pointer to [**RemoteAccessIdentitiesInnerType**](RemoteAccessIdentitiesInnerType.md) |  | [optional] 
**Label** | Pointer to **string** | The label for the remote access identity. | [optional] 
**Deleted** | Pointer to **bool** | Specifies whether the identity is available. | [optional] 

## Methods

### NewRemoteAccessIdentitiesInner

`func NewRemoteAccessIdentitiesInner() *RemoteAccessIdentitiesInner`

NewRemoteAccessIdentitiesInner instantiates a new RemoteAccessIdentitiesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteAccessIdentitiesInnerWithDefaults

`func NewRemoteAccessIdentitiesInnerWithDefaults() *RemoteAccessIdentitiesInner`

NewRemoteAccessIdentitiesInnerWithDefaults instantiates a new RemoteAccessIdentitiesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RemoteAccessIdentitiesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemoteAccessIdentitiesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemoteAccessIdentitiesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RemoteAccessIdentitiesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RemoteAccessIdentitiesInner) GetType() RemoteAccessIdentitiesInnerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RemoteAccessIdentitiesInner) GetTypeOk() (*RemoteAccessIdentitiesInnerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RemoteAccessIdentitiesInner) SetType(v RemoteAccessIdentitiesInnerType)`

SetType sets Type field to given value.

### HasType

`func (o *RemoteAccessIdentitiesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLabel

`func (o *RemoteAccessIdentitiesInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RemoteAccessIdentitiesInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RemoteAccessIdentitiesInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *RemoteAccessIdentitiesInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDeleted

`func (o *RemoteAccessIdentitiesInner) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *RemoteAccessIdentitiesInner) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *RemoteAccessIdentitiesInner) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *RemoteAccessIdentitiesInner) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


