# TopFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | **int64** | The number of requests for the file. | 
**Sha256** | **string** | The SHA256 hash for the entry. | 
**Categories** | [**[]Category**](Category.md) | The list of categories. | 
**Identitycount** | **int64** | The count of the identities for the entry. | 
**Filenames** | **[]string** | The list of filenames for the entry. | 
**Filetypes** | **[]string** | The list of filetypes for the entry. | 

## Methods

### NewTopFile

`func NewTopFile(requests int64, sha256 string, categories []Category, identitycount int64, filenames []string, filetypes []string, ) *TopFile`

NewTopFile instantiates a new TopFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopFileWithDefaults

`func NewTopFileWithDefaults() *TopFile`

NewTopFileWithDefaults instantiates a new TopFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *TopFile) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *TopFile) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *TopFile) SetRequests(v int64)`

SetRequests sets Requests field to given value.


### GetSha256

`func (o *TopFile) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *TopFile) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *TopFile) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.


### GetCategories

`func (o *TopFile) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *TopFile) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *TopFile) SetCategories(v []Category)`

SetCategories sets Categories field to given value.


### GetIdentitycount

`func (o *TopFile) GetIdentitycount() int64`

GetIdentitycount returns the Identitycount field if non-nil, zero value otherwise.

### GetIdentitycountOk

`func (o *TopFile) GetIdentitycountOk() (*int64, bool)`

GetIdentitycountOk returns a tuple with the Identitycount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentitycount

`func (o *TopFile) SetIdentitycount(v int64)`

SetIdentitycount sets Identitycount field to given value.


### GetFilenames

`func (o *TopFile) GetFilenames() []string`

GetFilenames returns the Filenames field if non-nil, zero value otherwise.

### GetFilenamesOk

`func (o *TopFile) GetFilenamesOk() (*[]string, bool)`

GetFilenamesOk returns a tuple with the Filenames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilenames

`func (o *TopFile) SetFilenames(v []string)`

SetFilenames sets Filenames field to given value.


### GetFiletypes

`func (o *TopFile) GetFiletypes() []string`

GetFiletypes returns the Filetypes field if non-nil, zero value otherwise.

### GetFiletypesOk

`func (o *TopFile) GetFiletypesOk() (*[]string, bool)`

GetFiletypesOk returns a tuple with the Filetypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiletypes

`func (o *TopFile) SetFiletypes(v []string)`

SetFiletypes sets Filetypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


