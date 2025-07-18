# SwgSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The type of traffic associated with the source. | [optional] 
**TotalTraffic** | Pointer to **int64** | The total number of bytes inbound and outbound for this app in your environment. | [optional] 
**BytesIn** | Pointer to **int64** | The number of bytes received (inbound). | [optional] 
**BytesOut** | Pointer to **int64** | The number of bytes sent (outbound). | [optional] 
**BlockedBytesOut** | Pointer to **int64** | The number of bytes of outbound traffic that are blocked. | [optional] 

## Methods

### NewSwgSource

`func NewSwgSource() *SwgSource`

NewSwgSource instantiates a new SwgSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwgSourceWithDefaults

`func NewSwgSourceWithDefaults() *SwgSource`

NewSwgSourceWithDefaults instantiates a new SwgSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SwgSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SwgSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SwgSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SwgSource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTotalTraffic

`func (o *SwgSource) GetTotalTraffic() int64`

GetTotalTraffic returns the TotalTraffic field if non-nil, zero value otherwise.

### GetTotalTrafficOk

`func (o *SwgSource) GetTotalTrafficOk() (*int64, bool)`

GetTotalTrafficOk returns a tuple with the TotalTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTraffic

`func (o *SwgSource) SetTotalTraffic(v int64)`

SetTotalTraffic sets TotalTraffic field to given value.

### HasTotalTraffic

`func (o *SwgSource) HasTotalTraffic() bool`

HasTotalTraffic returns a boolean if a field has been set.

### GetBytesIn

`func (o *SwgSource) GetBytesIn() int64`

GetBytesIn returns the BytesIn field if non-nil, zero value otherwise.

### GetBytesInOk

`func (o *SwgSource) GetBytesInOk() (*int64, bool)`

GetBytesInOk returns a tuple with the BytesIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesIn

`func (o *SwgSource) SetBytesIn(v int64)`

SetBytesIn sets BytesIn field to given value.

### HasBytesIn

`func (o *SwgSource) HasBytesIn() bool`

HasBytesIn returns a boolean if a field has been set.

### GetBytesOut

`func (o *SwgSource) GetBytesOut() int64`

GetBytesOut returns the BytesOut field if non-nil, zero value otherwise.

### GetBytesOutOk

`func (o *SwgSource) GetBytesOutOk() (*int64, bool)`

GetBytesOutOk returns a tuple with the BytesOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesOut

`func (o *SwgSource) SetBytesOut(v int64)`

SetBytesOut sets BytesOut field to given value.

### HasBytesOut

`func (o *SwgSource) HasBytesOut() bool`

HasBytesOut returns a boolean if a field has been set.

### GetBlockedBytesOut

`func (o *SwgSource) GetBlockedBytesOut() int64`

GetBlockedBytesOut returns the BlockedBytesOut field if non-nil, zero value otherwise.

### GetBlockedBytesOutOk

`func (o *SwgSource) GetBlockedBytesOutOk() (*int64, bool)`

GetBlockedBytesOutOk returns a tuple with the BlockedBytesOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedBytesOut

`func (o *SwgSource) SetBlockedBytesOut(v int64)`

SetBlockedBytesOut sets BlockedBytesOut field to given value.

### HasBlockedBytesOut

`func (o *SwgSource) HasBlockedBytesOut() bool`

HasBlockedBytesOut returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


