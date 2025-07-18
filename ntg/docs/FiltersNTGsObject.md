# FiltersNTGsObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of a Network Tunnel Group. The value of &#x60;name&#x60; is a sequence of case-insensitive characters. | [optional] 
**ExactName** | Pointer to **string** | The sequence of case-insensitive characters that exactly match the name of the Network Tunnel Group. When &#x60;exactName&#x60; is included as a filter, &#x60;name&#x60; is ignored. | [optional] 
**Region** | Pointer to **string** | The region for the Network Tunnel Group. The value of &#x60;region&#x60; is a sequence of case-insensitive characters. | [optional] 
**NetworkTunnelGroupIds** | Pointer to **string** | The comma-separated list of Network Tunnel Group IDs. | [optional] 
**ExactAuthIdPrefix** | Pointer to **string** | The case-sensitive value of the auth ID prefix (&#x60;authIdPrefix&#x60;) or IP for the network tunnel Hub. | [optional] 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**DuplicateCIDRs** | Pointer to **string** | List the network tunnel groups that have duplicate CIDRs. Provide the CIDRs and optionally provide the regional scope and region properties. If the regional scope is enabled, only duplicates in the same region are found. You can not use the &#x60;duplicateCIDRs&#x60; filter with any other filter. | [optional] 

## Methods

### NewFiltersNTGsObject

`func NewFiltersNTGsObject() *FiltersNTGsObject`

NewFiltersNTGsObject instantiates a new FiltersNTGsObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersNTGsObjectWithDefaults

`func NewFiltersNTGsObjectWithDefaults() *FiltersNTGsObject`

NewFiltersNTGsObjectWithDefaults instantiates a new FiltersNTGsObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FiltersNTGsObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FiltersNTGsObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FiltersNTGsObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FiltersNTGsObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExactName

`func (o *FiltersNTGsObject) GetExactName() string`

GetExactName returns the ExactName field if non-nil, zero value otherwise.

### GetExactNameOk

`func (o *FiltersNTGsObject) GetExactNameOk() (*string, bool)`

GetExactNameOk returns a tuple with the ExactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactName

`func (o *FiltersNTGsObject) SetExactName(v string)`

SetExactName sets ExactName field to given value.

### HasExactName

`func (o *FiltersNTGsObject) HasExactName() bool`

HasExactName returns a boolean if a field has been set.

### GetRegion

`func (o *FiltersNTGsObject) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *FiltersNTGsObject) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *FiltersNTGsObject) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *FiltersNTGsObject) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetNetworkTunnelGroupIds

`func (o *FiltersNTGsObject) GetNetworkTunnelGroupIds() string`

GetNetworkTunnelGroupIds returns the NetworkTunnelGroupIds field if non-nil, zero value otherwise.

### GetNetworkTunnelGroupIdsOk

`func (o *FiltersNTGsObject) GetNetworkTunnelGroupIdsOk() (*string, bool)`

GetNetworkTunnelGroupIdsOk returns a tuple with the NetworkTunnelGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTunnelGroupIds

`func (o *FiltersNTGsObject) SetNetworkTunnelGroupIds(v string)`

SetNetworkTunnelGroupIds sets NetworkTunnelGroupIds field to given value.

### HasNetworkTunnelGroupIds

`func (o *FiltersNTGsObject) HasNetworkTunnelGroupIds() bool`

HasNetworkTunnelGroupIds returns a boolean if a field has been set.

### GetExactAuthIdPrefix

`func (o *FiltersNTGsObject) GetExactAuthIdPrefix() string`

GetExactAuthIdPrefix returns the ExactAuthIdPrefix field if non-nil, zero value otherwise.

### GetExactAuthIdPrefixOk

`func (o *FiltersNTGsObject) GetExactAuthIdPrefixOk() (*string, bool)`

GetExactAuthIdPrefixOk returns a tuple with the ExactAuthIdPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactAuthIdPrefix

`func (o *FiltersNTGsObject) SetExactAuthIdPrefix(v string)`

SetExactAuthIdPrefix sets ExactAuthIdPrefix field to given value.

### HasExactAuthIdPrefix

`func (o *FiltersNTGsObject) HasExactAuthIdPrefix() bool`

HasExactAuthIdPrefix returns a boolean if a field has been set.

### GetStatus

`func (o *FiltersNTGsObject) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FiltersNTGsObject) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FiltersNTGsObject) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FiltersNTGsObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDuplicateCIDRs

`func (o *FiltersNTGsObject) GetDuplicateCIDRs() string`

GetDuplicateCIDRs returns the DuplicateCIDRs field if non-nil, zero value otherwise.

### GetDuplicateCIDRsOk

`func (o *FiltersNTGsObject) GetDuplicateCIDRsOk() (*string, bool)`

GetDuplicateCIDRsOk returns a tuple with the DuplicateCIDRs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateCIDRs

`func (o *FiltersNTGsObject) SetDuplicateCIDRs(v string)`

SetDuplicateCIDRs sets DuplicateCIDRs field to given value.

### HasDuplicateCIDRs

`func (o *FiltersNTGsObject) HasDuplicateCIDRs() bool`

HasDuplicateCIDRs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


