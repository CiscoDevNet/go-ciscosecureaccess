# CreateInternalNetworkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the internal network. | 
**IpAddress** | **string** | The IP (IPv4 or IPv6) address of the internal network. | 
**PrefixLength** | **int64** | The length of the internal network&#39;s prefix. The prefix length is between 8 and 32. | 
**SiteId** | Pointer to **int64** | The ID of the site, which is associated with the internal network. The internal network is associated with either a &#x60;siteId&#x60;, &#x60;networkId&#x60;, or &#x60;tunnelId&#x60;. | [optional] 
**NetworkId** | Pointer to **int64** | The ID of the network, which is associated with the internal network. The internal network is associated with either a &#x60;siteId&#x60;, &#x60;networkId&#x60;, or &#x60;tunnelId&#x60;. | [optional] 
**TunnelId** | Pointer to **int64** | The ID of the network tunnel group ID, which is associated with the internal network. The internal network is associated with either a &#x60;siteId&#x60;, &#x60;networkId&#x60;, or &#x60;tunnelId&#x60;. | [optional] 

## Methods

### NewCreateInternalNetworkRequest

`func NewCreateInternalNetworkRequest(name string, ipAddress string, prefixLength int64, ) *CreateInternalNetworkRequest`

NewCreateInternalNetworkRequest instantiates a new CreateInternalNetworkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInternalNetworkRequestWithDefaults

`func NewCreateInternalNetworkRequestWithDefaults() *CreateInternalNetworkRequest`

NewCreateInternalNetworkRequestWithDefaults instantiates a new CreateInternalNetworkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateInternalNetworkRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateInternalNetworkRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateInternalNetworkRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIpAddress

`func (o *CreateInternalNetworkRequest) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *CreateInternalNetworkRequest) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *CreateInternalNetworkRequest) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetPrefixLength

`func (o *CreateInternalNetworkRequest) GetPrefixLength() int64`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *CreateInternalNetworkRequest) GetPrefixLengthOk() (*int64, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *CreateInternalNetworkRequest) SetPrefixLength(v int64)`

SetPrefixLength sets PrefixLength field to given value.


### GetSiteId

`func (o *CreateInternalNetworkRequest) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *CreateInternalNetworkRequest) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *CreateInternalNetworkRequest) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *CreateInternalNetworkRequest) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetNetworkId

`func (o *CreateInternalNetworkRequest) GetNetworkId() int64`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *CreateInternalNetworkRequest) GetNetworkIdOk() (*int64, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *CreateInternalNetworkRequest) SetNetworkId(v int64)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *CreateInternalNetworkRequest) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetTunnelId

`func (o *CreateInternalNetworkRequest) GetTunnelId() int64`

GetTunnelId returns the TunnelId field if non-nil, zero value otherwise.

### GetTunnelIdOk

`func (o *CreateInternalNetworkRequest) GetTunnelIdOk() (*int64, bool)`

GetTunnelIdOk returns a tuple with the TunnelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelId

`func (o *CreateInternalNetworkRequest) SetTunnelId(v int64)`

SetTunnelId sets TunnelId field to given value.

### HasTunnelId

`func (o *CreateInternalNetworkRequest) HasTunnelId() bool`

HasTunnelId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


