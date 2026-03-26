# CreateInternalDomainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The internal domain. | 
**Description** | Pointer to **string** | The description of the internal domain. The description is a sequence of characters with a length from 1 through 50. | [optional] 
**IncludeAllVAs** | Pointer to **bool** | Specifies whether to apply the internal domain to all virtual appliances. | [optional] 
**IncludeAllMobileDevices** | Pointer to **bool** | Specifies whether to apply the internal domain to all mobile devices. | [optional] 
**SiteIds** | Pointer to **[]int64** | The list of site IDs associated with the domain. | [optional] 

## Methods

### NewCreateInternalDomainRequest

`func NewCreateInternalDomainRequest(domain string, ) *CreateInternalDomainRequest`

NewCreateInternalDomainRequest instantiates a new CreateInternalDomainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInternalDomainRequestWithDefaults

`func NewCreateInternalDomainRequestWithDefaults() *CreateInternalDomainRequest`

NewCreateInternalDomainRequestWithDefaults instantiates a new CreateInternalDomainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *CreateInternalDomainRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreateInternalDomainRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreateInternalDomainRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetDescription

`func (o *CreateInternalDomainRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateInternalDomainRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateInternalDomainRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateInternalDomainRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIncludeAllVAs

`func (o *CreateInternalDomainRequest) GetIncludeAllVAs() bool`

GetIncludeAllVAs returns the IncludeAllVAs field if non-nil, zero value otherwise.

### GetIncludeAllVAsOk

`func (o *CreateInternalDomainRequest) GetIncludeAllVAsOk() (*bool, bool)`

GetIncludeAllVAsOk returns a tuple with the IncludeAllVAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllVAs

`func (o *CreateInternalDomainRequest) SetIncludeAllVAs(v bool)`

SetIncludeAllVAs sets IncludeAllVAs field to given value.

### HasIncludeAllVAs

`func (o *CreateInternalDomainRequest) HasIncludeAllVAs() bool`

HasIncludeAllVAs returns a boolean if a field has been set.

### GetIncludeAllMobileDevices

`func (o *CreateInternalDomainRequest) GetIncludeAllMobileDevices() bool`

GetIncludeAllMobileDevices returns the IncludeAllMobileDevices field if non-nil, zero value otherwise.

### GetIncludeAllMobileDevicesOk

`func (o *CreateInternalDomainRequest) GetIncludeAllMobileDevicesOk() (*bool, bool)`

GetIncludeAllMobileDevicesOk returns a tuple with the IncludeAllMobileDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllMobileDevices

`func (o *CreateInternalDomainRequest) SetIncludeAllMobileDevices(v bool)`

SetIncludeAllMobileDevices sets IncludeAllMobileDevices field to given value.

### HasIncludeAllMobileDevices

`func (o *CreateInternalDomainRequest) HasIncludeAllMobileDevices() bool`

HasIncludeAllMobileDevices returns a boolean if a field has been set.

### GetSiteIds

`func (o *CreateInternalDomainRequest) GetSiteIds() []int64`

GetSiteIds returns the SiteIds field if non-nil, zero value otherwise.

### GetSiteIdsOk

`func (o *CreateInternalDomainRequest) GetSiteIdsOk() (*[]int64, bool)`

GetSiteIdsOk returns a tuple with the SiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteIds

`func (o *CreateInternalDomainRequest) SetSiteIds(v []int64)`

SetSiteIds sets SiteIds field to given value.

### HasSiteIds

`func (o *CreateInternalDomainRequest) HasSiteIds() bool`

HasSiteIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


