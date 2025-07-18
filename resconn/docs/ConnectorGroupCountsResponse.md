# ConnectorGroupCountsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connected** | Pointer to **int64** | The number of Resource Connector Groups in the organization that are in the connected state. | [optional] 
**Disabled** | Pointer to **int64** | The number of Resource Connector Groups in the organization that are in the disabled state. | [optional] 
**Disconnected** | Pointer to **int64** | The number of the Resource Connector Groups in the organization that are in the disconnected state. | [optional] 
**HasDisconnectedConnector** | Pointer to **int64** | The number of Resource Connector Groups in the organization that have at least one disconnected Connector. | [optional] 
**NoAssignedResources** | Pointer to **int64** | The number of Resource Connector Groups in the organization without any assigned private resources. | [optional] 
**Total** | Pointer to **int64** | The total number of Resource Connector Groups in the organization. | [optional] 
**Waiting** | Pointer to **int64** | The number of Resource Connector Groups in the organziation that are in the waiting state. | [optional] 

## Methods

### NewConnectorGroupCountsResponse

`func NewConnectorGroupCountsResponse() *ConnectorGroupCountsResponse`

NewConnectorGroupCountsResponse instantiates a new ConnectorGroupCountsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorGroupCountsResponseWithDefaults

`func NewConnectorGroupCountsResponseWithDefaults() *ConnectorGroupCountsResponse`

NewConnectorGroupCountsResponseWithDefaults instantiates a new ConnectorGroupCountsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnected

`func (o *ConnectorGroupCountsResponse) GetConnected() int64`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *ConnectorGroupCountsResponse) GetConnectedOk() (*int64, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *ConnectorGroupCountsResponse) SetConnected(v int64)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *ConnectorGroupCountsResponse) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetDisabled

`func (o *ConnectorGroupCountsResponse) GetDisabled() int64`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ConnectorGroupCountsResponse) GetDisabledOk() (*int64, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ConnectorGroupCountsResponse) SetDisabled(v int64)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ConnectorGroupCountsResponse) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDisconnected

`func (o *ConnectorGroupCountsResponse) GetDisconnected() int64`

GetDisconnected returns the Disconnected field if non-nil, zero value otherwise.

### GetDisconnectedOk

`func (o *ConnectorGroupCountsResponse) GetDisconnectedOk() (*int64, bool)`

GetDisconnectedOk returns a tuple with the Disconnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnected

`func (o *ConnectorGroupCountsResponse) SetDisconnected(v int64)`

SetDisconnected sets Disconnected field to given value.

### HasDisconnected

`func (o *ConnectorGroupCountsResponse) HasDisconnected() bool`

HasDisconnected returns a boolean if a field has been set.

### GetHasDisconnectedConnector

`func (o *ConnectorGroupCountsResponse) GetHasDisconnectedConnector() int64`

GetHasDisconnectedConnector returns the HasDisconnectedConnector field if non-nil, zero value otherwise.

### GetHasDisconnectedConnectorOk

`func (o *ConnectorGroupCountsResponse) GetHasDisconnectedConnectorOk() (*int64, bool)`

GetHasDisconnectedConnectorOk returns a tuple with the HasDisconnectedConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDisconnectedConnector

`func (o *ConnectorGroupCountsResponse) SetHasDisconnectedConnector(v int64)`

SetHasDisconnectedConnector sets HasDisconnectedConnector field to given value.

### HasHasDisconnectedConnector

`func (o *ConnectorGroupCountsResponse) HasHasDisconnectedConnector() bool`

HasHasDisconnectedConnector returns a boolean if a field has been set.

### GetNoAssignedResources

`func (o *ConnectorGroupCountsResponse) GetNoAssignedResources() int64`

GetNoAssignedResources returns the NoAssignedResources field if non-nil, zero value otherwise.

### GetNoAssignedResourcesOk

`func (o *ConnectorGroupCountsResponse) GetNoAssignedResourcesOk() (*int64, bool)`

GetNoAssignedResourcesOk returns a tuple with the NoAssignedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAssignedResources

`func (o *ConnectorGroupCountsResponse) SetNoAssignedResources(v int64)`

SetNoAssignedResources sets NoAssignedResources field to given value.

### HasNoAssignedResources

`func (o *ConnectorGroupCountsResponse) HasNoAssignedResources() bool`

HasNoAssignedResources returns a boolean if a field has been set.

### GetTotal

`func (o *ConnectorGroupCountsResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ConnectorGroupCountsResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ConnectorGroupCountsResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ConnectorGroupCountsResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetWaiting

`func (o *ConnectorGroupCountsResponse) GetWaiting() int64`

GetWaiting returns the Waiting field if non-nil, zero value otherwise.

### GetWaitingOk

`func (o *ConnectorGroupCountsResponse) GetWaitingOk() (*int64, bool)`

GetWaitingOk returns a tuple with the Waiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaiting

`func (o *ConnectorGroupCountsResponse) SetWaiting(v int64)`

SetWaiting sets Waiting field to given value.

### HasWaiting

`func (o *ConnectorGroupCountsResponse) HasWaiting() bool`

HasWaiting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


