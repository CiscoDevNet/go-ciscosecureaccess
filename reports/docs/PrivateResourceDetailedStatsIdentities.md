# PrivateResourceDetailedStatsIdentities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Privateresourceid** | Pointer to **int64** | The private resource ID. | [optional] 
**Identities** | Pointer to [**[]IdentityWithStats**](IdentityWithStats.md) | The list of identities. | [optional] 

## Methods

### NewPrivateResourceDetailedStatsIdentities

`func NewPrivateResourceDetailedStatsIdentities() *PrivateResourceDetailedStatsIdentities`

NewPrivateResourceDetailedStatsIdentities instantiates a new PrivateResourceDetailedStatsIdentities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateResourceDetailedStatsIdentitiesWithDefaults

`func NewPrivateResourceDetailedStatsIdentitiesWithDefaults() *PrivateResourceDetailedStatsIdentities`

NewPrivateResourceDetailedStatsIdentitiesWithDefaults instantiates a new PrivateResourceDetailedStatsIdentities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateresourceid

`func (o *PrivateResourceDetailedStatsIdentities) GetPrivateresourceid() int64`

GetPrivateresourceid returns the Privateresourceid field if non-nil, zero value otherwise.

### GetPrivateresourceidOk

`func (o *PrivateResourceDetailedStatsIdentities) GetPrivateresourceidOk() (*int64, bool)`

GetPrivateresourceidOk returns a tuple with the Privateresourceid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateresourceid

`func (o *PrivateResourceDetailedStatsIdentities) SetPrivateresourceid(v int64)`

SetPrivateresourceid sets Privateresourceid field to given value.

### HasPrivateresourceid

`func (o *PrivateResourceDetailedStatsIdentities) HasPrivateresourceid() bool`

HasPrivateresourceid returns a boolean if a field has been set.

### GetIdentities

`func (o *PrivateResourceDetailedStatsIdentities) GetIdentities() []IdentityWithStats`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *PrivateResourceDetailedStatsIdentities) GetIdentitiesOk() (*[]IdentityWithStats, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *PrivateResourceDetailedStatsIdentities) SetIdentities(v []IdentityWithStats)`

SetIdentities sets Identities field to given value.

### HasIdentities

`func (o *PrivateResourceDetailedStatsIdentities) HasIdentities() bool`

HasIdentities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


