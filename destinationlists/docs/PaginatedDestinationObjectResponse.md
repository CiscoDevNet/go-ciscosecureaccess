# PaginatedDestinationObjectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**Status**](Status.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 
**Data** | [**[]DestinationObjectWithStringId**](DestinationObjectWithStringId.md) | The list of destinations in the destination list. | 

## Methods

### NewPaginatedDestinationObjectResponse

`func NewPaginatedDestinationObjectResponse(status Status, meta Meta, data []DestinationObjectWithStringId, ) *PaginatedDestinationObjectResponse`

NewPaginatedDestinationObjectResponse instantiates a new PaginatedDestinationObjectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedDestinationObjectResponseWithDefaults

`func NewPaginatedDestinationObjectResponseWithDefaults() *PaginatedDestinationObjectResponse`

NewPaginatedDestinationObjectResponseWithDefaults instantiates a new PaginatedDestinationObjectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PaginatedDestinationObjectResponse) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaginatedDestinationObjectResponse) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaginatedDestinationObjectResponse) SetStatus(v Status)`

SetStatus sets Status field to given value.


### GetMeta

`func (o *PaginatedDestinationObjectResponse) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PaginatedDestinationObjectResponse) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PaginatedDestinationObjectResponse) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetData

`func (o *PaginatedDestinationObjectResponse) GetData() []DestinationObjectWithStringId`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaginatedDestinationObjectResponse) GetDataOk() (*[]DestinationObjectWithStringId, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaginatedDestinationObjectResponse) SetData(v []DestinationObjectWithStringId)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


