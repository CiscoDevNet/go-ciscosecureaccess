# ApplicationUsageResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **int64** | The ID of the application. | [optional] 
**UsedBy** | Pointer to [**ApplicationUsageResponseInnerUsedBy**](ApplicationUsageResponseInnerUsedBy.md) |  | [optional] 

## Methods

### NewApplicationUsageResponseInner

`func NewApplicationUsageResponseInner() *ApplicationUsageResponseInner`

NewApplicationUsageResponseInner instantiates a new ApplicationUsageResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationUsageResponseInnerWithDefaults

`func NewApplicationUsageResponseInnerWithDefaults() *ApplicationUsageResponseInner`

NewApplicationUsageResponseInnerWithDefaults instantiates a new ApplicationUsageResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *ApplicationUsageResponseInner) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationUsageResponseInner) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ApplicationUsageResponseInner) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *ApplicationUsageResponseInner) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetUsedBy

`func (o *ApplicationUsageResponseInner) GetUsedBy() ApplicationUsageResponseInnerUsedBy`

GetUsedBy returns the UsedBy field if non-nil, zero value otherwise.

### GetUsedByOk

`func (o *ApplicationUsageResponseInner) GetUsedByOk() (*ApplicationUsageResponseInnerUsedBy, bool)`

GetUsedByOk returns a tuple with the UsedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedBy

`func (o *ApplicationUsageResponseInner) SetUsedBy(v ApplicationUsageResponseInnerUsedBy)`

SetUsedBy sets UsedBy field to given value.

### HasUsedBy

`func (o *ApplicationUsageResponseInner) HasUsedBy() bool`

HasUsedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


