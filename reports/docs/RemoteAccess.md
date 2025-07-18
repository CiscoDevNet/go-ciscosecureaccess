# RemoteAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Osversion** | **string** | The OS version of the host. | 
**Internalip** | **string** | The internal IP address of the host. | 
**Connecttimestamp** | **int64** | The time and date when the host connected, expressed in seconds. | 
**Reason** | **string** | The reason the event occurred for the remote access endpoint. | 
**Failedreasons** | **[]string** | A list of messages that describe the remote access error. | 
**Connectionevent** | **string** | The type of event associated with the remote connection. | 
**Anyconnectversion** | **string** | The version of the AnyConnect Roaming Security module. | 
**Publicip** | **string** | The public IP address of the host | 
**Vpnprofile** | **string** | The VPN profile | 
**Sessiontype** | **string** | The session type of the remote connection. | 
**Timestamp** | **int64** | The timestamp represented in seconds. | 
**Identities** | [**[]RemoteAccessIdentitiesInner**](RemoteAccessIdentitiesInner.md) |  | 

## Methods

### NewRemoteAccess

`func NewRemoteAccess(osversion string, internalip string, connecttimestamp int64, reason string, failedreasons []string, connectionevent string, anyconnectversion string, publicip string, vpnprofile string, sessiontype string, timestamp int64, identities []RemoteAccessIdentitiesInner, ) *RemoteAccess`

NewRemoteAccess instantiates a new RemoteAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteAccessWithDefaults

`func NewRemoteAccessWithDefaults() *RemoteAccess`

NewRemoteAccessWithDefaults instantiates a new RemoteAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOsversion

`func (o *RemoteAccess) GetOsversion() string`

GetOsversion returns the Osversion field if non-nil, zero value otherwise.

### GetOsversionOk

`func (o *RemoteAccess) GetOsversionOk() (*string, bool)`

GetOsversionOk returns a tuple with the Osversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsversion

`func (o *RemoteAccess) SetOsversion(v string)`

SetOsversion sets Osversion field to given value.


### GetInternalip

`func (o *RemoteAccess) GetInternalip() string`

GetInternalip returns the Internalip field if non-nil, zero value otherwise.

### GetInternalipOk

`func (o *RemoteAccess) GetInternalipOk() (*string, bool)`

GetInternalipOk returns a tuple with the Internalip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalip

`func (o *RemoteAccess) SetInternalip(v string)`

SetInternalip sets Internalip field to given value.


### GetConnecttimestamp

`func (o *RemoteAccess) GetConnecttimestamp() int64`

GetConnecttimestamp returns the Connecttimestamp field if non-nil, zero value otherwise.

### GetConnecttimestampOk

`func (o *RemoteAccess) GetConnecttimestampOk() (*int64, bool)`

GetConnecttimestampOk returns a tuple with the Connecttimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnecttimestamp

`func (o *RemoteAccess) SetConnecttimestamp(v int64)`

SetConnecttimestamp sets Connecttimestamp field to given value.


### GetReason

`func (o *RemoteAccess) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RemoteAccess) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RemoteAccess) SetReason(v string)`

SetReason sets Reason field to given value.


### GetFailedreasons

`func (o *RemoteAccess) GetFailedreasons() []string`

GetFailedreasons returns the Failedreasons field if non-nil, zero value otherwise.

### GetFailedreasonsOk

`func (o *RemoteAccess) GetFailedreasonsOk() (*[]string, bool)`

GetFailedreasonsOk returns a tuple with the Failedreasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedreasons

`func (o *RemoteAccess) SetFailedreasons(v []string)`

SetFailedreasons sets Failedreasons field to given value.


### GetConnectionevent

`func (o *RemoteAccess) GetConnectionevent() string`

GetConnectionevent returns the Connectionevent field if non-nil, zero value otherwise.

### GetConnectioneventOk

`func (o *RemoteAccess) GetConnectioneventOk() (*string, bool)`

GetConnectioneventOk returns a tuple with the Connectionevent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionevent

`func (o *RemoteAccess) SetConnectionevent(v string)`

SetConnectionevent sets Connectionevent field to given value.


### GetAnyconnectversion

`func (o *RemoteAccess) GetAnyconnectversion() string`

GetAnyconnectversion returns the Anyconnectversion field if non-nil, zero value otherwise.

### GetAnyconnectversionOk

`func (o *RemoteAccess) GetAnyconnectversionOk() (*string, bool)`

GetAnyconnectversionOk returns a tuple with the Anyconnectversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyconnectversion

`func (o *RemoteAccess) SetAnyconnectversion(v string)`

SetAnyconnectversion sets Anyconnectversion field to given value.


### GetPublicip

`func (o *RemoteAccess) GetPublicip() string`

GetPublicip returns the Publicip field if non-nil, zero value otherwise.

### GetPublicipOk

`func (o *RemoteAccess) GetPublicipOk() (*string, bool)`

GetPublicipOk returns a tuple with the Publicip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicip

`func (o *RemoteAccess) SetPublicip(v string)`

SetPublicip sets Publicip field to given value.


### GetVpnprofile

`func (o *RemoteAccess) GetVpnprofile() string`

GetVpnprofile returns the Vpnprofile field if non-nil, zero value otherwise.

### GetVpnprofileOk

`func (o *RemoteAccess) GetVpnprofileOk() (*string, bool)`

GetVpnprofileOk returns a tuple with the Vpnprofile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnprofile

`func (o *RemoteAccess) SetVpnprofile(v string)`

SetVpnprofile sets Vpnprofile field to given value.


### GetSessiontype

`func (o *RemoteAccess) GetSessiontype() string`

GetSessiontype returns the Sessiontype field if non-nil, zero value otherwise.

### GetSessiontypeOk

`func (o *RemoteAccess) GetSessiontypeOk() (*string, bool)`

GetSessiontypeOk returns a tuple with the Sessiontype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessiontype

`func (o *RemoteAccess) SetSessiontype(v string)`

SetSessiontype sets Sessiontype field to given value.


### GetTimestamp

`func (o *RemoteAccess) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RemoteAccess) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RemoteAccess) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetIdentities

`func (o *RemoteAccess) GetIdentities() []RemoteAccessIdentitiesInner`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *RemoteAccess) GetIdentitiesOk() (*[]RemoteAccessIdentitiesInner, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *RemoteAccess) SetIdentities(v []RemoteAccessIdentitiesInner)`

SetIdentities sets Identities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


