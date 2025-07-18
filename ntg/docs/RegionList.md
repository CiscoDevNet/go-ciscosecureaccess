# RegionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regions** | Pointer to [**[]RegionListRegionsInner**](RegionListRegionsInner.md) | The list of regions for the Network Tunnel Groups. | [optional] 
**Bgp** | Pointer to [**Bgp**](Bgp.md) |  | [optional] 

## Methods

### NewRegionList

`func NewRegionList() *RegionList`

NewRegionList instantiates a new RegionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionListWithDefaults

`func NewRegionListWithDefaults() *RegionList`

NewRegionListWithDefaults instantiates a new RegionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegions

`func (o *RegionList) GetRegions() []RegionListRegionsInner`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *RegionList) GetRegionsOk() (*[]RegionListRegionsInner, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *RegionList) SetRegions(v []RegionListRegionsInner)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *RegionList) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetBgp

`func (o *RegionList) GetBgp() Bgp`

GetBgp returns the Bgp field if non-nil, zero value otherwise.

### GetBgpOk

`func (o *RegionList) GetBgpOk() (*Bgp, bool)`

GetBgpOk returns a tuple with the Bgp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgp

`func (o *RegionList) SetBgp(v Bgp)`

SetBgp sets Bgp field to given value.

### HasBgp

`func (o *RegionList) HasBgp() bool`

HasBgp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


