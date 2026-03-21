# OrgInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | **int64** | The organization ID. | 
**Fingerprint** | **string** | A hash that is used to register the Cisco Secure Client on users devices in the organization. | 
**UserId** | **int64** | The first 32 bits of the API key ID. | 

## Methods

### NewOrgInfo

`func NewOrgInfo(organizationId int64, fingerprint string, userId int64, ) *OrgInfo`

NewOrgInfo instantiates a new OrgInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgInfoWithDefaults

`func NewOrgInfoWithDefaults() *OrgInfo`

NewOrgInfoWithDefaults instantiates a new OrgInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *OrgInfo) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrgInfo) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrgInfo) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.


### GetFingerprint

`func (o *OrgInfo) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *OrgInfo) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *OrgInfo) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.


### GetUserId

`func (o *OrgInfo) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OrgInfo) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OrgInfo) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


