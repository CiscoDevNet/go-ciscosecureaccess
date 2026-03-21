# CreateSecureWebGatewayDeviceSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginIds** | **[]int64** | The list of origin IDs. The list can contain 1–100 origin IDs. | 
**Value** | [**Value**](Value.md) |  | 

## Methods

### NewCreateSecureWebGatewayDeviceSettingsRequest

`func NewCreateSecureWebGatewayDeviceSettingsRequest(originIds []int64, value Value, ) *CreateSecureWebGatewayDeviceSettingsRequest`

NewCreateSecureWebGatewayDeviceSettingsRequest instantiates a new CreateSecureWebGatewayDeviceSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSecureWebGatewayDeviceSettingsRequestWithDefaults

`func NewCreateSecureWebGatewayDeviceSettingsRequestWithDefaults() *CreateSecureWebGatewayDeviceSettingsRequest`

NewCreateSecureWebGatewayDeviceSettingsRequestWithDefaults instantiates a new CreateSecureWebGatewayDeviceSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginIds

`func (o *CreateSecureWebGatewayDeviceSettingsRequest) GetOriginIds() []int64`

GetOriginIds returns the OriginIds field if non-nil, zero value otherwise.

### GetOriginIdsOk

`func (o *CreateSecureWebGatewayDeviceSettingsRequest) GetOriginIdsOk() (*[]int64, bool)`

GetOriginIdsOk returns a tuple with the OriginIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginIds

`func (o *CreateSecureWebGatewayDeviceSettingsRequest) SetOriginIds(v []int64)`

SetOriginIds sets OriginIds field to given value.


### GetValue

`func (o *CreateSecureWebGatewayDeviceSettingsRequest) GetValue() Value`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateSecureWebGatewayDeviceSettingsRequest) GetValueOk() (*Value, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateSecureWebGatewayDeviceSettingsRequest) SetValue(v Value)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


