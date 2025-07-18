# UsageMeterItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BytesInDelta** | **int64** | The number of cumulative bytes received by Secure Access during the time window. | 
**BytesOutDelta** | **int64** | The number of cumulative bytes sent to Secure Access during the time window. | 

## Methods

### NewUsageMeterItem

`func NewUsageMeterItem(bytesInDelta int64, bytesOutDelta int64, ) *UsageMeterItem`

NewUsageMeterItem instantiates a new UsageMeterItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageMeterItemWithDefaults

`func NewUsageMeterItemWithDefaults() *UsageMeterItem`

NewUsageMeterItemWithDefaults instantiates a new UsageMeterItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytesInDelta

`func (o *UsageMeterItem) GetBytesInDelta() int64`

GetBytesInDelta returns the BytesInDelta field if non-nil, zero value otherwise.

### GetBytesInDeltaOk

`func (o *UsageMeterItem) GetBytesInDeltaOk() (*int64, bool)`

GetBytesInDeltaOk returns a tuple with the BytesInDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesInDelta

`func (o *UsageMeterItem) SetBytesInDelta(v int64)`

SetBytesInDelta sets BytesInDelta field to given value.


### GetBytesOutDelta

`func (o *UsageMeterItem) GetBytesOutDelta() int64`

GetBytesOutDelta returns the BytesOutDelta field if non-nil, zero value otherwise.

### GetBytesOutDeltaOk

`func (o *UsageMeterItem) GetBytesOutDeltaOk() (*int64, bool)`

GetBytesOutDeltaOk returns a tuple with the BytesOutDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesOutDelta

`func (o *UsageMeterItem) SetBytesOutDelta(v int64)`

SetBytesOutDelta sets BytesOutDelta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


