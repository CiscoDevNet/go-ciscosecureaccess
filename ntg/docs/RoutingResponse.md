# RoutingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the route. | 
**Data** | [**RoutingResponseData**](RoutingResponseData.md) |  | 

## Methods

### NewRoutingResponse

`func NewRoutingResponse(type_ string, data RoutingResponseData, ) *RoutingResponse`

NewRoutingResponse instantiates a new RoutingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingResponseWithDefaults

`func NewRoutingResponseWithDefaults() *RoutingResponse`

NewRoutingResponseWithDefaults instantiates a new RoutingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RoutingResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoutingResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoutingResponse) SetType(v string)`

SetType sets Type field to given value.


### GetData

`func (o *RoutingResponse) GetData() RoutingResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RoutingResponse) GetDataOk() (*RoutingResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RoutingResponse) SetData(v RoutingResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


