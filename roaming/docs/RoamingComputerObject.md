# RoamingComputerObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginId** | **int64** | The origin ID for the roaming computer. | 
**Name** | **string** | The name of the roaming computer. &#x60;name&#x60; is a sequence of 1–50 characters. | 
**DeviceId** | **string** | The hex ID of the roaming computer. | 
**Type** | **string** | The type of the roaming computer. | 
**Status** | **string** | The status of the roaming computer with DNS-layer security. | 
**SwgStatus** | **string** | The status of the roaming computer with Internet security (Secure Web Gateway). | 
**LastSyncStatus** | **string** | The status of the last sync on the roaming computer with DNS-layer security. | 
**LastSyncSwgStatus** | **string** | The status of the last sync on the roaming computer with Internet security (Secure Web Gateway). | 
**LastSync** | **time.Time** | The date and time (timestamp) of the last sync. | 
**AppliedBundle** | **int64** | The policy ID. | 
**HasIpBlocking** | **bool** | Specifies whether the roaming computer has IP blocking. | 
**Version** | **string** | The version of the Cisco Secure Client with the Internet Security module deployed on the roaming computer. | 
**OsVersion** | **string** | The OS version of the roaming computer. | 
**OsVersionName** | **string** | The OS version name of the roaming computer. | 
**AnyconnectDeviceId** | Pointer to **string** | The ID of the device that has the Cisco Secure Client deployed with the Internet Security module. | [optional] 

## Methods

### NewRoamingComputerObject

`func NewRoamingComputerObject(originId int64, name string, deviceId string, type_ string, status string, swgStatus string, lastSyncStatus string, lastSyncSwgStatus string, lastSync time.Time, appliedBundle int64, hasIpBlocking bool, version string, osVersion string, osVersionName string, ) *RoamingComputerObject`

NewRoamingComputerObject instantiates a new RoamingComputerObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoamingComputerObjectWithDefaults

`func NewRoamingComputerObjectWithDefaults() *RoamingComputerObject`

NewRoamingComputerObjectWithDefaults instantiates a new RoamingComputerObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginId

`func (o *RoamingComputerObject) GetOriginId() int64`

GetOriginId returns the OriginId field if non-nil, zero value otherwise.

### GetOriginIdOk

`func (o *RoamingComputerObject) GetOriginIdOk() (*int64, bool)`

GetOriginIdOk returns a tuple with the OriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginId

`func (o *RoamingComputerObject) SetOriginId(v int64)`

SetOriginId sets OriginId field to given value.


### GetName

`func (o *RoamingComputerObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoamingComputerObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoamingComputerObject) SetName(v string)`

SetName sets Name field to given value.


### GetDeviceId

`func (o *RoamingComputerObject) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *RoamingComputerObject) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *RoamingComputerObject) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetType

`func (o *RoamingComputerObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoamingComputerObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoamingComputerObject) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *RoamingComputerObject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RoamingComputerObject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RoamingComputerObject) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSwgStatus

`func (o *RoamingComputerObject) GetSwgStatus() string`

GetSwgStatus returns the SwgStatus field if non-nil, zero value otherwise.

### GetSwgStatusOk

`func (o *RoamingComputerObject) GetSwgStatusOk() (*string, bool)`

GetSwgStatusOk returns a tuple with the SwgStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwgStatus

`func (o *RoamingComputerObject) SetSwgStatus(v string)`

SetSwgStatus sets SwgStatus field to given value.


### GetLastSyncStatus

`func (o *RoamingComputerObject) GetLastSyncStatus() string`

GetLastSyncStatus returns the LastSyncStatus field if non-nil, zero value otherwise.

### GetLastSyncStatusOk

`func (o *RoamingComputerObject) GetLastSyncStatusOk() (*string, bool)`

GetLastSyncStatusOk returns a tuple with the LastSyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncStatus

`func (o *RoamingComputerObject) SetLastSyncStatus(v string)`

SetLastSyncStatus sets LastSyncStatus field to given value.


### GetLastSyncSwgStatus

`func (o *RoamingComputerObject) GetLastSyncSwgStatus() string`

GetLastSyncSwgStatus returns the LastSyncSwgStatus field if non-nil, zero value otherwise.

### GetLastSyncSwgStatusOk

`func (o *RoamingComputerObject) GetLastSyncSwgStatusOk() (*string, bool)`

GetLastSyncSwgStatusOk returns a tuple with the LastSyncSwgStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncSwgStatus

`func (o *RoamingComputerObject) SetLastSyncSwgStatus(v string)`

SetLastSyncSwgStatus sets LastSyncSwgStatus field to given value.


### GetLastSync

`func (o *RoamingComputerObject) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *RoamingComputerObject) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *RoamingComputerObject) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.


### GetAppliedBundle

`func (o *RoamingComputerObject) GetAppliedBundle() int64`

GetAppliedBundle returns the AppliedBundle field if non-nil, zero value otherwise.

### GetAppliedBundleOk

`func (o *RoamingComputerObject) GetAppliedBundleOk() (*int64, bool)`

GetAppliedBundleOk returns a tuple with the AppliedBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedBundle

`func (o *RoamingComputerObject) SetAppliedBundle(v int64)`

SetAppliedBundle sets AppliedBundle field to given value.


### GetHasIpBlocking

`func (o *RoamingComputerObject) GetHasIpBlocking() bool`

GetHasIpBlocking returns the HasIpBlocking field if non-nil, zero value otherwise.

### GetHasIpBlockingOk

`func (o *RoamingComputerObject) GetHasIpBlockingOk() (*bool, bool)`

GetHasIpBlockingOk returns a tuple with the HasIpBlocking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIpBlocking

`func (o *RoamingComputerObject) SetHasIpBlocking(v bool)`

SetHasIpBlocking sets HasIpBlocking field to given value.


### GetVersion

`func (o *RoamingComputerObject) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RoamingComputerObject) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RoamingComputerObject) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetOsVersion

`func (o *RoamingComputerObject) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *RoamingComputerObject) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *RoamingComputerObject) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.


### GetOsVersionName

`func (o *RoamingComputerObject) GetOsVersionName() string`

GetOsVersionName returns the OsVersionName field if non-nil, zero value otherwise.

### GetOsVersionNameOk

`func (o *RoamingComputerObject) GetOsVersionNameOk() (*string, bool)`

GetOsVersionNameOk returns a tuple with the OsVersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersionName

`func (o *RoamingComputerObject) SetOsVersionName(v string)`

SetOsVersionName sets OsVersionName field to given value.


### GetAnyconnectDeviceId

`func (o *RoamingComputerObject) GetAnyconnectDeviceId() string`

GetAnyconnectDeviceId returns the AnyconnectDeviceId field if non-nil, zero value otherwise.

### GetAnyconnectDeviceIdOk

`func (o *RoamingComputerObject) GetAnyconnectDeviceIdOk() (*string, bool)`

GetAnyconnectDeviceIdOk returns a tuple with the AnyconnectDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyconnectDeviceId

`func (o *RoamingComputerObject) SetAnyconnectDeviceId(v string)`

SetAnyconnectDeviceId sets AnyconnectDeviceId field to given value.

### HasAnyconnectDeviceId

`func (o *RoamingComputerObject) HasAnyconnectDeviceId() bool`

HasAnyconnectDeviceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


