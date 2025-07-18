# RoutingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the route. | 
**Data** | [**RoutingRequestData**](RoutingRequestData.md) |  | 

## Methods

### NewRoutingRequest

`func NewRoutingRequest(type_ string, data RoutingRequestData, ) *RoutingRequest`

NewRoutingRequest instantiates a new RoutingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingRequestWithDefaults

`func NewRoutingRequestWithDefaults() *RoutingRequest`

NewRoutingRequestWithDefaults instantiates a new RoutingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RoutingRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoutingRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoutingRequest) SetType(v string)`

SetType sets Type field to given value.


### GetData

`func (o *RoutingRequest) GetData() RoutingRequestData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RoutingRequest) GetDataOk() (*RoutingRequestData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RoutingRequest) SetData(v RoutingRequestData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


