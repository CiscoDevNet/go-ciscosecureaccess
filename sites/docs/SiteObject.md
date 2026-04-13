# SiteObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginId** | **int64** | The origin ID of the Site. | 
**Name** | **string** | The name of the Site. The name is a sequence of 1–255 characters. | 
**SiteId** | **int64** | The ID of the Site. | 
**IsDefault** | **bool** | Specifies whether the Site is the default Site. | 
**Type** | Pointer to **string** | The type of the Site. | [optional] 
**InternalNetworkCount** | Pointer to **int64** | The number of internal networks that are associated with the Site. | [optional] 
**VaCount** | Pointer to **int64** | The number of virtual appliances that are associated with the Site. | [optional] 
**CreatedAt** | **time.Time** | The date and time (ISO 8601 timestamp) when the Site was created.  | 
**ModifiedAt** | **time.Time** | The date and time (ISO 8601 timestamp) when the Site was modified. | 

## Methods

### NewSiteObject

`func NewSiteObject(originId int64, name string, siteId int64, isDefault bool, createdAt time.Time, modifiedAt time.Time, ) *SiteObject`

NewSiteObject instantiates a new SiteObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteObjectWithDefaults

`func NewSiteObjectWithDefaults() *SiteObject`

NewSiteObjectWithDefaults instantiates a new SiteObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginId

`func (o *SiteObject) GetOriginId() int64`

GetOriginId returns the OriginId field if non-nil, zero value otherwise.

### GetOriginIdOk

`func (o *SiteObject) GetOriginIdOk() (*int64, bool)`

GetOriginIdOk returns a tuple with the OriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginId

`func (o *SiteObject) SetOriginId(v int64)`

SetOriginId sets OriginId field to given value.


### GetName

`func (o *SiteObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteObject) SetName(v string)`

SetName sets Name field to given value.


### GetSiteId

`func (o *SiteObject) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SiteObject) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SiteObject) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.


### GetIsDefault

`func (o *SiteObject) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *SiteObject) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *SiteObject) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetType

`func (o *SiteObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SiteObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SiteObject) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SiteObject) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInternalNetworkCount

`func (o *SiteObject) GetInternalNetworkCount() int64`

GetInternalNetworkCount returns the InternalNetworkCount field if non-nil, zero value otherwise.

### GetInternalNetworkCountOk

`func (o *SiteObject) GetInternalNetworkCountOk() (*int64, bool)`

GetInternalNetworkCountOk returns a tuple with the InternalNetworkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalNetworkCount

`func (o *SiteObject) SetInternalNetworkCount(v int64)`

SetInternalNetworkCount sets InternalNetworkCount field to given value.

### HasInternalNetworkCount

`func (o *SiteObject) HasInternalNetworkCount() bool`

HasInternalNetworkCount returns a boolean if a field has been set.

### GetVaCount

`func (o *SiteObject) GetVaCount() int64`

GetVaCount returns the VaCount field if non-nil, zero value otherwise.

### GetVaCountOk

`func (o *SiteObject) GetVaCountOk() (*int64, bool)`

GetVaCountOk returns a tuple with the VaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaCount

`func (o *SiteObject) SetVaCount(v int64)`

SetVaCount sets VaCount field to given value.

### HasVaCount

`func (o *SiteObject) HasVaCount() bool`

HasVaCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SiteObject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SiteObject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SiteObject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *SiteObject) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SiteObject) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *SiteObject) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


