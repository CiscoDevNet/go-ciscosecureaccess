# PaginatedDestinationListsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**Status**](Status.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 
**Data** | [**[]DestinationListObject**](DestinationListObject.md) | The list of destination lists. | 

## Methods

### NewPaginatedDestinationListsResponse

`func NewPaginatedDestinationListsResponse(status Status, meta Meta, data []DestinationListObject, ) *PaginatedDestinationListsResponse`

NewPaginatedDestinationListsResponse instantiates a new PaginatedDestinationListsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedDestinationListsResponseWithDefaults

`func NewPaginatedDestinationListsResponseWithDefaults() *PaginatedDestinationListsResponse`

NewPaginatedDestinationListsResponseWithDefaults instantiates a new PaginatedDestinationListsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PaginatedDestinationListsResponse) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaginatedDestinationListsResponse) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaginatedDestinationListsResponse) SetStatus(v Status)`

SetStatus sets Status field to given value.


### GetMeta

`func (o *PaginatedDestinationListsResponse) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PaginatedDestinationListsResponse) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PaginatedDestinationListsResponse) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetData

`func (o *PaginatedDestinationListsResponse) GetData() []DestinationListObject`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaginatedDestinationListsResponse) GetDataOk() (*[]DestinationListObject, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaginatedDestinationListsResponse) SetData(v []DestinationListObject)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


