# PrivateResourceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | Pointer to **int64** | The ID of the Private Resource. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the Private Resource that is unique across all resources in the organization. The Resource name must not have any special characters other than spaces and hyphens. | [optional] 
**Description** | Pointer to **string** | The description of the Private Resource. | [optional] 
**DnsServerId** | Pointer to **int64** | The unique identifier of the DNS server that is used to resolve IP addresses. | [optional] 
**CertificateId** | Pointer to **int64** | The ID of the certificate that is used to decrypt traffic to the Private Resource. | [optional] 
**AccessTypes** | Pointer to [**[]AccessTypesInner**](AccessTypesInner.md) | The list of connection properties that describe how end users can access private resources in the organization. Browser-based access requires that all resource addresses use the HTTP or HTTPS protocols. Branch access is always enabled for a private resource. | [optional] 
**ResourceAddresses** | Pointer to [**[]ResourceAddressesInner**](ResourceAddressesInner.md) | The list of resource addresses for the Private Resources. | [optional] 
**ResourceGroupIds** | Pointer to **[]int64** | The list of IDs for the Private Resource Groups that include the Private Resource. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when the private resource was created. | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** | The date and time when the private resource was updated. | [optional] [readonly] 
**CreatedBy** | Pointer to **string** | The ID of the user who created the private resource. | [optional] 
**ModifiedBy** | Pointer to **string** | The ID of the user who last modified private resource. | [optional] 

## Methods

### NewPrivateResourceResponse

`func NewPrivateResourceResponse() *PrivateResourceResponse`

NewPrivateResourceResponse instantiates a new PrivateResourceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateResourceResponseWithDefaults

`func NewPrivateResourceResponseWithDefaults() *PrivateResourceResponse`

NewPrivateResourceResponseWithDefaults instantiates a new PrivateResourceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *PrivateResourceResponse) GetResourceId() int64`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *PrivateResourceResponse) GetResourceIdOk() (*int64, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *PrivateResourceResponse) SetResourceId(v int64)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *PrivateResourceResponse) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetName

`func (o *PrivateResourceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateResourceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateResourceResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrivateResourceResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PrivateResourceResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrivateResourceResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrivateResourceResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PrivateResourceResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDnsServerId

`func (o *PrivateResourceResponse) GetDnsServerId() int64`

GetDnsServerId returns the DnsServerId field if non-nil, zero value otherwise.

### GetDnsServerIdOk

`func (o *PrivateResourceResponse) GetDnsServerIdOk() (*int64, bool)`

GetDnsServerIdOk returns a tuple with the DnsServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServerId

`func (o *PrivateResourceResponse) SetDnsServerId(v int64)`

SetDnsServerId sets DnsServerId field to given value.

### HasDnsServerId

`func (o *PrivateResourceResponse) HasDnsServerId() bool`

HasDnsServerId returns a boolean if a field has been set.

### GetCertificateId

`func (o *PrivateResourceResponse) GetCertificateId() int64`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *PrivateResourceResponse) GetCertificateIdOk() (*int64, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *PrivateResourceResponse) SetCertificateId(v int64)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *PrivateResourceResponse) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### GetAccessTypes

`func (o *PrivateResourceResponse) GetAccessTypes() []AccessTypesInner`

GetAccessTypes returns the AccessTypes field if non-nil, zero value otherwise.

### GetAccessTypesOk

`func (o *PrivateResourceResponse) GetAccessTypesOk() (*[]AccessTypesInner, bool)`

GetAccessTypesOk returns a tuple with the AccessTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTypes

`func (o *PrivateResourceResponse) SetAccessTypes(v []AccessTypesInner)`

SetAccessTypes sets AccessTypes field to given value.

### HasAccessTypes

`func (o *PrivateResourceResponse) HasAccessTypes() bool`

HasAccessTypes returns a boolean if a field has been set.

### GetResourceAddresses

`func (o *PrivateResourceResponse) GetResourceAddresses() []ResourceAddressesInner`

GetResourceAddresses returns the ResourceAddresses field if non-nil, zero value otherwise.

### GetResourceAddressesOk

`func (o *PrivateResourceResponse) GetResourceAddressesOk() (*[]ResourceAddressesInner, bool)`

GetResourceAddressesOk returns a tuple with the ResourceAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAddresses

`func (o *PrivateResourceResponse) SetResourceAddresses(v []ResourceAddressesInner)`

SetResourceAddresses sets ResourceAddresses field to given value.

### HasResourceAddresses

`func (o *PrivateResourceResponse) HasResourceAddresses() bool`

HasResourceAddresses returns a boolean if a field has been set.

### GetResourceGroupIds

`func (o *PrivateResourceResponse) GetResourceGroupIds() []int64`

GetResourceGroupIds returns the ResourceGroupIds field if non-nil, zero value otherwise.

### GetResourceGroupIdsOk

`func (o *PrivateResourceResponse) GetResourceGroupIdsOk() (*[]int64, bool)`

GetResourceGroupIdsOk returns a tuple with the ResourceGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupIds

`func (o *PrivateResourceResponse) SetResourceGroupIds(v []int64)`

SetResourceGroupIds sets ResourceGroupIds field to given value.

### HasResourceGroupIds

`func (o *PrivateResourceResponse) HasResourceGroupIds() bool`

HasResourceGroupIds returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PrivateResourceResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PrivateResourceResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PrivateResourceResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PrivateResourceResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *PrivateResourceResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PrivateResourceResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PrivateResourceResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *PrivateResourceResponse) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *PrivateResourceResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PrivateResourceResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PrivateResourceResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PrivateResourceResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetModifiedBy

`func (o *PrivateResourceResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *PrivateResourceResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *PrivateResourceResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *PrivateResourceResponse) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


