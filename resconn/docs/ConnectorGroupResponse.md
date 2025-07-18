# ConnectorGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the Resource Connector Group. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the Resource Connector Group. The Resource Connector Group name may not include any special characters other than spaces and hyphens. | [optional] 
**Location** | Pointer to **string** | The region where the Resource Connector Group is available. | [optional] 
**Environment** | Pointer to [**Environment**](Environment.md) |  | [optional] [default to AWS]
**ProvisioningKey** | Pointer to **string** | The Resource Connector Group&#39;s provisioning key. | [optional] 
**ProvisioningKeyExpiresAt** | Pointer to **time.Time** | The date and time of the expiration of the Resource Connector Group&#39;s provisioning key, specified in the ISO 8601 format. | [optional] [readonly] 
**BaseImageDownloadUrl** | Pointer to **string** | The URL of the Connector Group&#39;s base image. | [optional] [readonly] 
**Status** | Pointer to **string** | The label that describes the status of the Resource Connector Group. | [optional] 
**StatusUpdatedAt** | Pointer to **time.Time** | The date and time of the udpate of the Resource Connector Group&#39;s status, specified in the ISO 8601 format. | [optional] [readonly] 
**ConnectorsCount** | Pointer to **int64** | The number of Connectors in the Resource Connector Group. | [optional] 
**ResourceIds** | Pointer to **[]int64** | The list of resource IDs. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time of the addition of the new Resource Connector Group, specified in the ISO 8601 format. | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** | The date and time of the update to the Resource Connector Group, specified in the ISO 8601 format. | [optional] [readonly] 
**DisconnectedConnectorsCount** | Pointer to **int64** | The number of disconnected Connectors in the Resource Connector Group. | [optional] 

## Methods

### NewConnectorGroupResponse

`func NewConnectorGroupResponse() *ConnectorGroupResponse`

NewConnectorGroupResponse instantiates a new ConnectorGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorGroupResponseWithDefaults

`func NewConnectorGroupResponseWithDefaults() *ConnectorGroupResponse`

NewConnectorGroupResponseWithDefaults instantiates a new ConnectorGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectorGroupResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectorGroupResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectorGroupResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectorGroupResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ConnectorGroupResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorGroupResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorGroupResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectorGroupResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocation

`func (o *ConnectorGroupResponse) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ConnectorGroupResponse) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ConnectorGroupResponse) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ConnectorGroupResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetEnvironment

`func (o *ConnectorGroupResponse) GetEnvironment() Environment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ConnectorGroupResponse) GetEnvironmentOk() (*Environment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ConnectorGroupResponse) SetEnvironment(v Environment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ConnectorGroupResponse) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetProvisioningKey

`func (o *ConnectorGroupResponse) GetProvisioningKey() string`

GetProvisioningKey returns the ProvisioningKey field if non-nil, zero value otherwise.

### GetProvisioningKeyOk

`func (o *ConnectorGroupResponse) GetProvisioningKeyOk() (*string, bool)`

GetProvisioningKeyOk returns a tuple with the ProvisioningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningKey

`func (o *ConnectorGroupResponse) SetProvisioningKey(v string)`

SetProvisioningKey sets ProvisioningKey field to given value.

### HasProvisioningKey

`func (o *ConnectorGroupResponse) HasProvisioningKey() bool`

HasProvisioningKey returns a boolean if a field has been set.

### GetProvisioningKeyExpiresAt

`func (o *ConnectorGroupResponse) GetProvisioningKeyExpiresAt() time.Time`

GetProvisioningKeyExpiresAt returns the ProvisioningKeyExpiresAt field if non-nil, zero value otherwise.

### GetProvisioningKeyExpiresAtOk

`func (o *ConnectorGroupResponse) GetProvisioningKeyExpiresAtOk() (*time.Time, bool)`

GetProvisioningKeyExpiresAtOk returns a tuple with the ProvisioningKeyExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningKeyExpiresAt

`func (o *ConnectorGroupResponse) SetProvisioningKeyExpiresAt(v time.Time)`

SetProvisioningKeyExpiresAt sets ProvisioningKeyExpiresAt field to given value.

### HasProvisioningKeyExpiresAt

`func (o *ConnectorGroupResponse) HasProvisioningKeyExpiresAt() bool`

HasProvisioningKeyExpiresAt returns a boolean if a field has been set.

### GetBaseImageDownloadUrl

`func (o *ConnectorGroupResponse) GetBaseImageDownloadUrl() string`

GetBaseImageDownloadUrl returns the BaseImageDownloadUrl field if non-nil, zero value otherwise.

### GetBaseImageDownloadUrlOk

`func (o *ConnectorGroupResponse) GetBaseImageDownloadUrlOk() (*string, bool)`

GetBaseImageDownloadUrlOk returns a tuple with the BaseImageDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseImageDownloadUrl

`func (o *ConnectorGroupResponse) SetBaseImageDownloadUrl(v string)`

SetBaseImageDownloadUrl sets BaseImageDownloadUrl field to given value.

### HasBaseImageDownloadUrl

`func (o *ConnectorGroupResponse) HasBaseImageDownloadUrl() bool`

HasBaseImageDownloadUrl returns a boolean if a field has been set.

### GetStatus

`func (o *ConnectorGroupResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectorGroupResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectorGroupResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConnectorGroupResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusUpdatedAt

`func (o *ConnectorGroupResponse) GetStatusUpdatedAt() time.Time`

GetStatusUpdatedAt returns the StatusUpdatedAt field if non-nil, zero value otherwise.

### GetStatusUpdatedAtOk

`func (o *ConnectorGroupResponse) GetStatusUpdatedAtOk() (*time.Time, bool)`

GetStatusUpdatedAtOk returns a tuple with the StatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUpdatedAt

`func (o *ConnectorGroupResponse) SetStatusUpdatedAt(v time.Time)`

SetStatusUpdatedAt sets StatusUpdatedAt field to given value.

### HasStatusUpdatedAt

`func (o *ConnectorGroupResponse) HasStatusUpdatedAt() bool`

HasStatusUpdatedAt returns a boolean if a field has been set.

### GetConnectorsCount

`func (o *ConnectorGroupResponse) GetConnectorsCount() int64`

GetConnectorsCount returns the ConnectorsCount field if non-nil, zero value otherwise.

### GetConnectorsCountOk

`func (o *ConnectorGroupResponse) GetConnectorsCountOk() (*int64, bool)`

GetConnectorsCountOk returns a tuple with the ConnectorsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorsCount

`func (o *ConnectorGroupResponse) SetConnectorsCount(v int64)`

SetConnectorsCount sets ConnectorsCount field to given value.

### HasConnectorsCount

`func (o *ConnectorGroupResponse) HasConnectorsCount() bool`

HasConnectorsCount returns a boolean if a field has been set.

### GetResourceIds

`func (o *ConnectorGroupResponse) GetResourceIds() []int64`

GetResourceIds returns the ResourceIds field if non-nil, zero value otherwise.

### GetResourceIdsOk

`func (o *ConnectorGroupResponse) GetResourceIdsOk() (*[]int64, bool)`

GetResourceIdsOk returns a tuple with the ResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIds

`func (o *ConnectorGroupResponse) SetResourceIds(v []int64)`

SetResourceIds sets ResourceIds field to given value.

### HasResourceIds

`func (o *ConnectorGroupResponse) HasResourceIds() bool`

HasResourceIds returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ConnectorGroupResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConnectorGroupResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConnectorGroupResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ConnectorGroupResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *ConnectorGroupResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ConnectorGroupResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ConnectorGroupResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *ConnectorGroupResponse) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetDisconnectedConnectorsCount

`func (o *ConnectorGroupResponse) GetDisconnectedConnectorsCount() int64`

GetDisconnectedConnectorsCount returns the DisconnectedConnectorsCount field if non-nil, zero value otherwise.

### GetDisconnectedConnectorsCountOk

`func (o *ConnectorGroupResponse) GetDisconnectedConnectorsCountOk() (*int64, bool)`

GetDisconnectedConnectorsCountOk returns a tuple with the DisconnectedConnectorsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedConnectorsCount

`func (o *ConnectorGroupResponse) SetDisconnectedConnectorsCount(v int64)`

SetDisconnectedConnectorsCount sets DisconnectedConnectorsCount field to given value.

### HasDisconnectedConnectorsCount

`func (o *ConnectorGroupResponse) HasDisconnectedConnectorsCount() bool`

HasDisconnectedConnectorsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


