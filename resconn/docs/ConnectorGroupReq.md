# ConnectorGroupReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Resource Connector Group. The Resource Connector Group name may not include any special characters other than spaces and hyphens. | 
**Location** | **string** | The region where the Resource Connector Group is available. | 
**Environment** | Pointer to [**Environment**](Environment.md) |  | [optional] [default to AWS]

## Methods

### NewConnectorGroupReq

`func NewConnectorGroupReq(name string, location string, ) *ConnectorGroupReq`

NewConnectorGroupReq instantiates a new ConnectorGroupReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorGroupReqWithDefaults

`func NewConnectorGroupReqWithDefaults() *ConnectorGroupReq`

NewConnectorGroupReqWithDefaults instantiates a new ConnectorGroupReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConnectorGroupReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorGroupReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorGroupReq) SetName(v string)`

SetName sets Name field to given value.


### GetLocation

`func (o *ConnectorGroupReq) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ConnectorGroupReq) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ConnectorGroupReq) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetEnvironment

`func (o *ConnectorGroupReq) GetEnvironment() Environment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ConnectorGroupReq) GetEnvironmentOk() (*Environment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ConnectorGroupReq) SetEnvironment(v Environment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ConnectorGroupReq) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


