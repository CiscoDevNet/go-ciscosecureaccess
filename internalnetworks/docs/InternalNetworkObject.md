# InternalNetworkObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginId** | **int64** | The origin ID of the internal network. | 
**Name** | **string** | The name of the internal network. | 
**IpAddress** | **string** | The IP (IPv4 or IPv6) address of the internal network. | 
**PrefixLength** | **int64** | The length of the internal network&#39;s prefix. The prefix length is between 8 and 32. | 
**SiteName** | Pointer to **string** | The name of the site that is associated with the internal network. | [optional] 
**SiteId** | Pointer to **int64** | The ID of the site, which is associated with the internal network. The internal network is associated with either a &#x60;siteId&#x60;, &#x60;networkId&#x60;, or &#x60;tunnelId&#x60;. | [optional] 
**NetworkName** | Pointer to **string** | The name of the network that is associated with the internal network. | [optional] 
**NetworkId** | Pointer to **int64** | The ID of the network, which is associated with the internal network. The internal network is associated with either a &#x60;siteId&#x60;, &#x60;networkId&#x60;, or &#x60;tunnelId&#x60;. | [optional] 
**TunnelName** | Pointer to **string** | The name of the network tunnel group that is associated with the internal network. | [optional] 
**TunnelId** | Pointer to **int64** | The ID of the network tunnel group ID, which is associated with the internal network. The internal network is associated with either a &#x60;siteId&#x60;, &#x60;networkId&#x60;, or &#x60;tunnelId&#x60;. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time (ISO8601 timestamp) when the internal network was created. | [optional] 
**ModifiedAt** | Pointer to **time.Time** | The date and time (ISO8601 timestamp) when the internal network was modified. | [optional] 

## Methods

### NewInternalNetworkObject

`func NewInternalNetworkObject(originId int64, name string, ipAddress string, prefixLength int64, ) *InternalNetworkObject`

NewInternalNetworkObject instantiates a new InternalNetworkObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalNetworkObjectWithDefaults

`func NewInternalNetworkObjectWithDefaults() *InternalNetworkObject`

NewInternalNetworkObjectWithDefaults instantiates a new InternalNetworkObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginId

`func (o *InternalNetworkObject) GetOriginId() int64`

GetOriginId returns the OriginId field if non-nil, zero value otherwise.

### GetOriginIdOk

`func (o *InternalNetworkObject) GetOriginIdOk() (*int64, bool)`

GetOriginIdOk returns a tuple with the OriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginId

`func (o *InternalNetworkObject) SetOriginId(v int64)`

SetOriginId sets OriginId field to given value.


### GetName

`func (o *InternalNetworkObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InternalNetworkObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InternalNetworkObject) SetName(v string)`

SetName sets Name field to given value.


### GetIpAddress

`func (o *InternalNetworkObject) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *InternalNetworkObject) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *InternalNetworkObject) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetPrefixLength

`func (o *InternalNetworkObject) GetPrefixLength() int64`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *InternalNetworkObject) GetPrefixLengthOk() (*int64, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *InternalNetworkObject) SetPrefixLength(v int64)`

SetPrefixLength sets PrefixLength field to given value.


### GetSiteName

`func (o *InternalNetworkObject) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *InternalNetworkObject) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *InternalNetworkObject) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *InternalNetworkObject) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSiteId

`func (o *InternalNetworkObject) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *InternalNetworkObject) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *InternalNetworkObject) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *InternalNetworkObject) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetNetworkName

`func (o *InternalNetworkObject) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *InternalNetworkObject) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *InternalNetworkObject) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *InternalNetworkObject) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetNetworkId

`func (o *InternalNetworkObject) GetNetworkId() int64`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InternalNetworkObject) GetNetworkIdOk() (*int64, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InternalNetworkObject) SetNetworkId(v int64)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InternalNetworkObject) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetTunnelName

`func (o *InternalNetworkObject) GetTunnelName() string`

GetTunnelName returns the TunnelName field if non-nil, zero value otherwise.

### GetTunnelNameOk

`func (o *InternalNetworkObject) GetTunnelNameOk() (*string, bool)`

GetTunnelNameOk returns a tuple with the TunnelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelName

`func (o *InternalNetworkObject) SetTunnelName(v string)`

SetTunnelName sets TunnelName field to given value.

### HasTunnelName

`func (o *InternalNetworkObject) HasTunnelName() bool`

HasTunnelName returns a boolean if a field has been set.

### GetTunnelId

`func (o *InternalNetworkObject) GetTunnelId() int64`

GetTunnelId returns the TunnelId field if non-nil, zero value otherwise.

### GetTunnelIdOk

`func (o *InternalNetworkObject) GetTunnelIdOk() (*int64, bool)`

GetTunnelIdOk returns a tuple with the TunnelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelId

`func (o *InternalNetworkObject) SetTunnelId(v int64)`

SetTunnelId sets TunnelId field to given value.

### HasTunnelId

`func (o *InternalNetworkObject) HasTunnelId() bool`

HasTunnelId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InternalNetworkObject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InternalNetworkObject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InternalNetworkObject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InternalNetworkObject) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *InternalNetworkObject) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *InternalNetworkObject) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *InternalNetworkObject) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *InternalNetworkObject) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


