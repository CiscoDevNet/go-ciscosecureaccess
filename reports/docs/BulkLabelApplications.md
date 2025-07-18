# BulkLabelApplications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | The time of the update. | [optional] 
**NumberOfApps** | Pointer to **int64** | The number of apps updated | [optional] 

## Methods

### NewBulkLabelApplications

`func NewBulkLabelApplications() *BulkLabelApplications`

NewBulkLabelApplications instantiates a new BulkLabelApplications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkLabelApplicationsWithDefaults

`func NewBulkLabelApplicationsWithDefaults() *BulkLabelApplications`

NewBulkLabelApplicationsWithDefaults instantiates a new BulkLabelApplications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *BulkLabelApplications) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BulkLabelApplications) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BulkLabelApplications) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BulkLabelApplications) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetNumberOfApps

`func (o *BulkLabelApplications) GetNumberOfApps() int64`

GetNumberOfApps returns the NumberOfApps field if non-nil, zero value otherwise.

### GetNumberOfAppsOk

`func (o *BulkLabelApplications) GetNumberOfAppsOk() (*int64, bool)`

GetNumberOfAppsOk returns a tuple with the NumberOfApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfApps

`func (o *BulkLabelApplications) SetNumberOfApps(v int64)`

SetNumberOfApps sets NumberOfApps field to given value.

### HasNumberOfApps

`func (o *BulkLabelApplications) HasNumberOfApps() bool`

HasNumberOfApps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


