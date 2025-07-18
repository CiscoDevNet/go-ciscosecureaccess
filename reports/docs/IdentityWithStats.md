# IdentityWithStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the identity concatenated with the organization ID and the organization ID of the identity. | 
**Type** | [**IdentityType**](IdentityType.md) |  | 
**Label** | **string** | The descriptive label for the identity. | 
**Deleted** | **bool** | Indicates whether the identity was deleted. | 
**Hitscount** | **int64** | Indicates the number of times that an identity accessed a private resource. | 
**Success** | **int64** | Indicates the number of times that an identity was allowed to access a private resource. | 
**Blocked** | **int64** | Indicates the number of times that an identity was blocked access to a private resource. | 

## Methods

### NewIdentityWithStats

`func NewIdentityWithStats(id int64, type_ IdentityType, label string, deleted bool, hitscount int64, success int64, blocked int64, ) *IdentityWithStats`

NewIdentityWithStats instantiates a new IdentityWithStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityWithStatsWithDefaults

`func NewIdentityWithStatsWithDefaults() *IdentityWithStats`

NewIdentityWithStatsWithDefaults instantiates a new IdentityWithStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdentityWithStats) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityWithStats) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityWithStats) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *IdentityWithStats) GetType() IdentityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdentityWithStats) GetTypeOk() (*IdentityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdentityWithStats) SetType(v IdentityType)`

SetType sets Type field to given value.


### GetLabel

`func (o *IdentityWithStats) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *IdentityWithStats) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *IdentityWithStats) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDeleted

`func (o *IdentityWithStats) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *IdentityWithStats) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *IdentityWithStats) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.


### GetHitscount

`func (o *IdentityWithStats) GetHitscount() int64`

GetHitscount returns the Hitscount field if non-nil, zero value otherwise.

### GetHitscountOk

`func (o *IdentityWithStats) GetHitscountOk() (*int64, bool)`

GetHitscountOk returns a tuple with the Hitscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitscount

`func (o *IdentityWithStats) SetHitscount(v int64)`

SetHitscount sets Hitscount field to given value.


### GetSuccess

`func (o *IdentityWithStats) GetSuccess() int64`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *IdentityWithStats) GetSuccessOk() (*int64, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *IdentityWithStats) SetSuccess(v int64)`

SetSuccess sets Success field to given value.


### GetBlocked

`func (o *IdentityWithStats) GetBlocked() int64`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *IdentityWithStats) GetBlockedOk() (*int64, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *IdentityWithStats) SetBlocked(v int64)`

SetBlocked sets Blocked field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


