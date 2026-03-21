# InternalDomainObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the internal domain. | 
**Domain** | **string** | The domain name of the internal domain. | 
**Description** | **string** | The description of the internal domain. | 
**IncludeAllVAs** | **bool** | Specifies whether to apply the internal domain to all virtual appliances. | 
**IncludeAllMobileDevices** | **bool** | Specifies whether to apply the internal domain to all mobile devices. | 
**CreatedAt** | **time.Time** | The date and time (ISO 8601 timestamp) when the internal domain was created. | 
**ModifiedAt** | **time.Time** | The date and time (ISO 8601 timestamp) when the internal domain was modified. | 
**SiteIds** | **[]int64** | The list of site IDs associated with the domain. | 

## Methods

### NewInternalDomainObject

`func NewInternalDomainObject(id int64, domain string, description string, includeAllVAs bool, includeAllMobileDevices bool, createdAt time.Time, modifiedAt time.Time, siteIds []int64, ) *InternalDomainObject`

NewInternalDomainObject instantiates a new InternalDomainObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalDomainObjectWithDefaults

`func NewInternalDomainObjectWithDefaults() *InternalDomainObject`

NewInternalDomainObjectWithDefaults instantiates a new InternalDomainObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InternalDomainObject) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternalDomainObject) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternalDomainObject) SetId(v int64)`

SetId sets Id field to given value.


### GetDomain

`func (o *InternalDomainObject) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *InternalDomainObject) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *InternalDomainObject) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetDescription

`func (o *InternalDomainObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InternalDomainObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InternalDomainObject) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIncludeAllVAs

`func (o *InternalDomainObject) GetIncludeAllVAs() bool`

GetIncludeAllVAs returns the IncludeAllVAs field if non-nil, zero value otherwise.

### GetIncludeAllVAsOk

`func (o *InternalDomainObject) GetIncludeAllVAsOk() (*bool, bool)`

GetIncludeAllVAsOk returns a tuple with the IncludeAllVAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllVAs

`func (o *InternalDomainObject) SetIncludeAllVAs(v bool)`

SetIncludeAllVAs sets IncludeAllVAs field to given value.


### GetIncludeAllMobileDevices

`func (o *InternalDomainObject) GetIncludeAllMobileDevices() bool`

GetIncludeAllMobileDevices returns the IncludeAllMobileDevices field if non-nil, zero value otherwise.

### GetIncludeAllMobileDevicesOk

`func (o *InternalDomainObject) GetIncludeAllMobileDevicesOk() (*bool, bool)`

GetIncludeAllMobileDevicesOk returns a tuple with the IncludeAllMobileDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllMobileDevices

`func (o *InternalDomainObject) SetIncludeAllMobileDevices(v bool)`

SetIncludeAllMobileDevices sets IncludeAllMobileDevices field to given value.


### GetCreatedAt

`func (o *InternalDomainObject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InternalDomainObject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InternalDomainObject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *InternalDomainObject) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *InternalDomainObject) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *InternalDomainObject) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetSiteIds

`func (o *InternalDomainObject) GetSiteIds() []int64`

GetSiteIds returns the SiteIds field if non-nil, zero value otherwise.

### GetSiteIdsOk

`func (o *InternalDomainObject) GetSiteIdsOk() (*[]int64, bool)`

GetSiteIdsOk returns a tuple with the SiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteIds

`func (o *InternalDomainObject) SetSiteIds(v []int64)`

SetSiteIds sets SiteIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


