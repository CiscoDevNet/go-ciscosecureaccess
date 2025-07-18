# ConnectorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the Connector. | [optional] [readonly] 
**GroupId** | Pointer to **int64** | The ID of the Resource Connector Group. | [optional] [readonly] 
**InstanceId** | Pointer to **string** | The globally unique ID of the Connector instance. | [optional] 
**Confirmed** | Pointer to **bool** | Indicates whether the Connector exists. | [optional] [default to false]
**Enabled** | Pointer to **bool** | Indicates whether the Connector can receive traffic. | [optional] [default to false]
**Version** | Pointer to **string** | The runtime version of the Connector image. | [optional] [readonly] 
**Sha1** | Pointer to **string** | The unique ID of the Connector formatted as a hash string. | [optional] 
**Hostname** | Pointer to **string** | The unique hostname of the device that manages the runtime of the Connector. | [optional] 
**OriginIpAddress** | Pointer to **string** | The IP address of the Connector. | [optional] [readonly] 
**BaseVersion** | Pointer to **string** | The Base OS version of the Connector agent image. Updating the base OS version requires that you download and redeploy the latest Connector image. | [optional] [readonly] 
**IsLatestBaseVersion** | Pointer to **bool** | Indicates whether the Connector is using the latest available image. If this value is false, a new image is available to download. We recommend that you redeploy this Connector with the latest image. | [optional] [readonly] 
**UpgradeStatus** | Pointer to **string** | Indicates the status of the latest over-the-air update for this Connector. Over-the-air updates modify the Connector agent software, which are made without downloading a new image and redeploying the Connector. | [optional] [readonly] 
**Status** | Pointer to **string** | The label that describes the status of the Connector. | [optional] [readonly] [default to "announced"]
**StatusUpdatedAt** | Pointer to **time.Time** | The date and time of the Connector&#39;s status update, specified in the ISO 8601 format. | [optional] [readonly] 
**RevokedAt** | Pointer to **time.Time** | The date and time of the removal of the Connector, specified in the ISO 8601 format. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | The date and time of the addition of the new Connector, specified in the ISO 8601 format. | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** | The date and time of the Connector&#39;s update, specified in the ISO 8601 format. | [optional] [readonly] 

## Methods

### NewConnectorResponse

`func NewConnectorResponse() *ConnectorResponse`

NewConnectorResponse instantiates a new ConnectorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorResponseWithDefaults

`func NewConnectorResponseWithDefaults() *ConnectorResponse`

NewConnectorResponseWithDefaults instantiates a new ConnectorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectorResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectorResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectorResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectorResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroupId

`func (o *ConnectorResponse) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ConnectorResponse) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ConnectorResponse) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ConnectorResponse) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetInstanceId

`func (o *ConnectorResponse) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ConnectorResponse) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ConnectorResponse) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *ConnectorResponse) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetConfirmed

`func (o *ConnectorResponse) GetConfirmed() bool`

GetConfirmed returns the Confirmed field if non-nil, zero value otherwise.

### GetConfirmedOk

`func (o *ConnectorResponse) GetConfirmedOk() (*bool, bool)`

GetConfirmedOk returns a tuple with the Confirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmed

`func (o *ConnectorResponse) SetConfirmed(v bool)`

SetConfirmed sets Confirmed field to given value.

### HasConfirmed

`func (o *ConnectorResponse) HasConfirmed() bool`

HasConfirmed returns a boolean if a field has been set.

### GetEnabled

`func (o *ConnectorResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ConnectorResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ConnectorResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ConnectorResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVersion

`func (o *ConnectorResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConnectorResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConnectorResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ConnectorResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSha1

`func (o *ConnectorResponse) GetSha1() string`

GetSha1 returns the Sha1 field if non-nil, zero value otherwise.

### GetSha1Ok

`func (o *ConnectorResponse) GetSha1Ok() (*string, bool)`

GetSha1Ok returns a tuple with the Sha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1

`func (o *ConnectorResponse) SetSha1(v string)`

SetSha1 sets Sha1 field to given value.

### HasSha1

`func (o *ConnectorResponse) HasSha1() bool`

HasSha1 returns a boolean if a field has been set.

### GetHostname

`func (o *ConnectorResponse) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ConnectorResponse) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ConnectorResponse) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ConnectorResponse) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetOriginIpAddress

`func (o *ConnectorResponse) GetOriginIpAddress() string`

GetOriginIpAddress returns the OriginIpAddress field if non-nil, zero value otherwise.

### GetOriginIpAddressOk

`func (o *ConnectorResponse) GetOriginIpAddressOk() (*string, bool)`

GetOriginIpAddressOk returns a tuple with the OriginIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginIpAddress

`func (o *ConnectorResponse) SetOriginIpAddress(v string)`

SetOriginIpAddress sets OriginIpAddress field to given value.

### HasOriginIpAddress

`func (o *ConnectorResponse) HasOriginIpAddress() bool`

HasOriginIpAddress returns a boolean if a field has been set.

### GetBaseVersion

`func (o *ConnectorResponse) GetBaseVersion() string`

GetBaseVersion returns the BaseVersion field if non-nil, zero value otherwise.

### GetBaseVersionOk

`func (o *ConnectorResponse) GetBaseVersionOk() (*string, bool)`

GetBaseVersionOk returns a tuple with the BaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseVersion

`func (o *ConnectorResponse) SetBaseVersion(v string)`

SetBaseVersion sets BaseVersion field to given value.

### HasBaseVersion

`func (o *ConnectorResponse) HasBaseVersion() bool`

HasBaseVersion returns a boolean if a field has been set.

### GetIsLatestBaseVersion

`func (o *ConnectorResponse) GetIsLatestBaseVersion() bool`

GetIsLatestBaseVersion returns the IsLatestBaseVersion field if non-nil, zero value otherwise.

### GetIsLatestBaseVersionOk

`func (o *ConnectorResponse) GetIsLatestBaseVersionOk() (*bool, bool)`

GetIsLatestBaseVersionOk returns a tuple with the IsLatestBaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLatestBaseVersion

`func (o *ConnectorResponse) SetIsLatestBaseVersion(v bool)`

SetIsLatestBaseVersion sets IsLatestBaseVersion field to given value.

### HasIsLatestBaseVersion

`func (o *ConnectorResponse) HasIsLatestBaseVersion() bool`

HasIsLatestBaseVersion returns a boolean if a field has been set.

### GetUpgradeStatus

`func (o *ConnectorResponse) GetUpgradeStatus() string`

GetUpgradeStatus returns the UpgradeStatus field if non-nil, zero value otherwise.

### GetUpgradeStatusOk

`func (o *ConnectorResponse) GetUpgradeStatusOk() (*string, bool)`

GetUpgradeStatusOk returns a tuple with the UpgradeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStatus

`func (o *ConnectorResponse) SetUpgradeStatus(v string)`

SetUpgradeStatus sets UpgradeStatus field to given value.

### HasUpgradeStatus

`func (o *ConnectorResponse) HasUpgradeStatus() bool`

HasUpgradeStatus returns a boolean if a field has been set.

### GetStatus

`func (o *ConnectorResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectorResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectorResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConnectorResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusUpdatedAt

`func (o *ConnectorResponse) GetStatusUpdatedAt() time.Time`

GetStatusUpdatedAt returns the StatusUpdatedAt field if non-nil, zero value otherwise.

### GetStatusUpdatedAtOk

`func (o *ConnectorResponse) GetStatusUpdatedAtOk() (*time.Time, bool)`

GetStatusUpdatedAtOk returns a tuple with the StatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUpdatedAt

`func (o *ConnectorResponse) SetStatusUpdatedAt(v time.Time)`

SetStatusUpdatedAt sets StatusUpdatedAt field to given value.

### HasStatusUpdatedAt

`func (o *ConnectorResponse) HasStatusUpdatedAt() bool`

HasStatusUpdatedAt returns a boolean if a field has been set.

### GetRevokedAt

`func (o *ConnectorResponse) GetRevokedAt() time.Time`

GetRevokedAt returns the RevokedAt field if non-nil, zero value otherwise.

### GetRevokedAtOk

`func (o *ConnectorResponse) GetRevokedAtOk() (*time.Time, bool)`

GetRevokedAtOk returns a tuple with the RevokedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedAt

`func (o *ConnectorResponse) SetRevokedAt(v time.Time)`

SetRevokedAt sets RevokedAt field to given value.

### HasRevokedAt

`func (o *ConnectorResponse) HasRevokedAt() bool`

HasRevokedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ConnectorResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConnectorResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConnectorResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ConnectorResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *ConnectorResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ConnectorResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ConnectorResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *ConnectorResponse) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


