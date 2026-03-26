# ListSecureWebGatewayDeviceSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginIds** | **[]int64** | The list of origin IDs. The list can contain 1–100 origin IDs. | 

## Methods

### NewListSecureWebGatewayDeviceSettingsRequest

`func NewListSecureWebGatewayDeviceSettingsRequest(originIds []int64, ) *ListSecureWebGatewayDeviceSettingsRequest`

NewListSecureWebGatewayDeviceSettingsRequest instantiates a new ListSecureWebGatewayDeviceSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSecureWebGatewayDeviceSettingsRequestWithDefaults

`func NewListSecureWebGatewayDeviceSettingsRequestWithDefaults() *ListSecureWebGatewayDeviceSettingsRequest`

NewListSecureWebGatewayDeviceSettingsRequestWithDefaults instantiates a new ListSecureWebGatewayDeviceSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginIds

`func (o *ListSecureWebGatewayDeviceSettingsRequest) GetOriginIds() []int64`

GetOriginIds returns the OriginIds field if non-nil, zero value otherwise.

### GetOriginIdsOk

`func (o *ListSecureWebGatewayDeviceSettingsRequest) GetOriginIdsOk() (*[]int64, bool)`

GetOriginIdsOk returns a tuple with the OriginIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginIds

`func (o *ListSecureWebGatewayDeviceSettingsRequest) SetOriginIds(v []int64)`

SetOriginIds sets OriginIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


