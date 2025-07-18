# StaticDataRequestObj

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkCIDRs** | **[]string** | The public and private address ranges that are used internally by your organization. | 

## Methods

### NewStaticDataRequestObj

`func NewStaticDataRequestObj(networkCIDRs []string, ) *StaticDataRequestObj`

NewStaticDataRequestObj instantiates a new StaticDataRequestObj object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticDataRequestObjWithDefaults

`func NewStaticDataRequestObjWithDefaults() *StaticDataRequestObj`

NewStaticDataRequestObjWithDefaults instantiates a new StaticDataRequestObj object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkCIDRs

`func (o *StaticDataRequestObj) GetNetworkCIDRs() []string`

GetNetworkCIDRs returns the NetworkCIDRs field if non-nil, zero value otherwise.

### GetNetworkCIDRsOk

`func (o *StaticDataRequestObj) GetNetworkCIDRsOk() (*[]string, bool)`

GetNetworkCIDRsOk returns a tuple with the NetworkCIDRs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCIDRs

`func (o *StaticDataRequestObj) SetNetworkCIDRs(v []string)`

SetNetworkCIDRs sets NetworkCIDRs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


