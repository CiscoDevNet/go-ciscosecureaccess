# KeysInformationInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyName** | **string** | The name of the API key. | 
**KeyId** | **string** | The ID of the API key. | 
**Count** | **int64** | The total number of API requests. | 
**Requests** | [**[]RequestDetailsListInner**](RequestDetailsListInner.md) | The list of API request information. | 

## Methods

### NewKeysInformationInner

`func NewKeysInformationInner(keyName string, keyId string, count int64, requests []RequestDetailsListInner, ) *KeysInformationInner`

NewKeysInformationInner instantiates a new KeysInformationInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeysInformationInnerWithDefaults

`func NewKeysInformationInnerWithDefaults() *KeysInformationInner`

NewKeysInformationInnerWithDefaults instantiates a new KeysInformationInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyName

`func (o *KeysInformationInner) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *KeysInformationInner) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *KeysInformationInner) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### GetKeyId

`func (o *KeysInformationInner) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *KeysInformationInner) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *KeysInformationInner) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### GetCount

`func (o *KeysInformationInner) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *KeysInformationInner) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *KeysInformationInner) SetCount(v int64)`

SetCount sets Count field to given value.


### GetRequests

`func (o *KeysInformationInner) GetRequests() []RequestDetailsListInner`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *KeysInformationInner) GetRequestsOk() (*[]RequestDetailsListInner, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *KeysInformationInner) SetRequests(v []RequestDetailsListInner)`

SetRequests sets Requests field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


