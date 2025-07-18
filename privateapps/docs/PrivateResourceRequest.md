# PrivateResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Private Resource that is unique across all resources in the organization. The Resource name must not have any special characters other than spaces and hyphens. | 
**Description** | Pointer to **string** | The description of the Private Resource. | [optional] 
**DnsServerId** | Pointer to **int64** | The unique identifier of the DNS server that is used to resolve IP addresses. | [optional] 
**CertificateId** | Pointer to **int64** | The ID of the certificate that is used to decrypt traffic to the Private Resource. | [optional] 
**AccessTypes** | [**[]AccessTypesRequestInner**](AccessTypesRequestInner.md) | The list of connection properties that describe how end users can access private resources in the organization. Browser-based access requires that all resource addresses use the HTTP or HTTPS protocols. Branch access is always enabled for a private resource. | 
**ResourceAddresses** | [**[]ResourceAddressesInner**](ResourceAddressesInner.md) | The list of resource addresses for the Private Resources. | 
**ResourceGroupIds** | Pointer to **[]int64** | The list of IDs for the Private Resource Groups that include the Private Resource. | [optional] 

## Methods

### NewPrivateResourceRequest

`func NewPrivateResourceRequest(name string, accessTypes []AccessTypesRequestInner, resourceAddresses []ResourceAddressesInner, ) *PrivateResourceRequest`

NewPrivateResourceRequest instantiates a new PrivateResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateResourceRequestWithDefaults

`func NewPrivateResourceRequestWithDefaults() *PrivateResourceRequest`

NewPrivateResourceRequestWithDefaults instantiates a new PrivateResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PrivateResourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateResourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateResourceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PrivateResourceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrivateResourceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrivateResourceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PrivateResourceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDnsServerId

`func (o *PrivateResourceRequest) GetDnsServerId() int64`

GetDnsServerId returns the DnsServerId field if non-nil, zero value otherwise.

### GetDnsServerIdOk

`func (o *PrivateResourceRequest) GetDnsServerIdOk() (*int64, bool)`

GetDnsServerIdOk returns a tuple with the DnsServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServerId

`func (o *PrivateResourceRequest) SetDnsServerId(v int64)`

SetDnsServerId sets DnsServerId field to given value.

### HasDnsServerId

`func (o *PrivateResourceRequest) HasDnsServerId() bool`

HasDnsServerId returns a boolean if a field has been set.

### GetCertificateId

`func (o *PrivateResourceRequest) GetCertificateId() int64`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *PrivateResourceRequest) GetCertificateIdOk() (*int64, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *PrivateResourceRequest) SetCertificateId(v int64)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *PrivateResourceRequest) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### GetAccessTypes

`func (o *PrivateResourceRequest) GetAccessTypes() []AccessTypesRequestInner`

GetAccessTypes returns the AccessTypes field if non-nil, zero value otherwise.

### GetAccessTypesOk

`func (o *PrivateResourceRequest) GetAccessTypesOk() (*[]AccessTypesRequestInner, bool)`

GetAccessTypesOk returns a tuple with the AccessTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTypes

`func (o *PrivateResourceRequest) SetAccessTypes(v []AccessTypesRequestInner)`

SetAccessTypes sets AccessTypes field to given value.


### GetResourceAddresses

`func (o *PrivateResourceRequest) GetResourceAddresses() []ResourceAddressesInner`

GetResourceAddresses returns the ResourceAddresses field if non-nil, zero value otherwise.

### GetResourceAddressesOk

`func (o *PrivateResourceRequest) GetResourceAddressesOk() (*[]ResourceAddressesInner, bool)`

GetResourceAddressesOk returns a tuple with the ResourceAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAddresses

`func (o *PrivateResourceRequest) SetResourceAddresses(v []ResourceAddressesInner)`

SetResourceAddresses sets ResourceAddresses field to given value.


### GetResourceGroupIds

`func (o *PrivateResourceRequest) GetResourceGroupIds() []int64`

GetResourceGroupIds returns the ResourceGroupIds field if non-nil, zero value otherwise.

### GetResourceGroupIdsOk

`func (o *PrivateResourceRequest) GetResourceGroupIdsOk() (*[]int64, bool)`

GetResourceGroupIdsOk returns a tuple with the ResourceGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupIds

`func (o *PrivateResourceRequest) SetResourceGroupIds(v []int64)`

SetResourceGroupIds sets ResourceGroupIds field to given value.

### HasResourceGroupIds

`func (o *PrivateResourceRequest) HasResourceGroupIds() bool`

HasResourceGroupIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


