# DestinationListCreateDestinationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to **string** | A domain, URL, or IP. | [optional] 
**Type** | Pointer to [**ModelType**](ModelType.md) |  | [optional] 
**Comment** | Pointer to **string** | The comment about the destination. | [optional] 

## Methods

### NewDestinationListCreateDestinationsInner

`func NewDestinationListCreateDestinationsInner() *DestinationListCreateDestinationsInner`

NewDestinationListCreateDestinationsInner instantiates a new DestinationListCreateDestinationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationListCreateDestinationsInnerWithDefaults

`func NewDestinationListCreateDestinationsInnerWithDefaults() *DestinationListCreateDestinationsInner`

NewDestinationListCreateDestinationsInnerWithDefaults instantiates a new DestinationListCreateDestinationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *DestinationListCreateDestinationsInner) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *DestinationListCreateDestinationsInner) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *DestinationListCreateDestinationsInner) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *DestinationListCreateDestinationsInner) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetType

`func (o *DestinationListCreateDestinationsInner) GetType() ModelType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestinationListCreateDestinationsInner) GetTypeOk() (*ModelType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestinationListCreateDestinationsInner) SetType(v ModelType)`

SetType sets Type field to given value.

### HasType

`func (o *DestinationListCreateDestinationsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetComment

`func (o *DestinationListCreateDestinationsInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DestinationListCreateDestinationsInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DestinationListCreateDestinationsInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DestinationListCreateDestinationsInner) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


