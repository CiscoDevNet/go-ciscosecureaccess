# Usage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | **int64** | The unique Secure Access organization ID. | 
**Product** | **string** | The name of the Secure Access data source service. | 
**RegisteredAt** | **time.Time** | The ISO 8601 timestamp when the usage was recorded in the origin service. The timestamp uses the UTC format following the ISO 8601 date time standard. | 
**UsageMeterItems** | [**UsageMeterItem**](UsageMeterItem.md) |  | 

## Methods

### NewUsage

`func NewUsage(organizationId int64, product string, registeredAt time.Time, usageMeterItems UsageMeterItem, ) *Usage`

NewUsage instantiates a new Usage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageWithDefaults

`func NewUsageWithDefaults() *Usage`

NewUsageWithDefaults instantiates a new Usage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *Usage) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Usage) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Usage) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.


### GetProduct

`func (o *Usage) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *Usage) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *Usage) SetProduct(v string)`

SetProduct sets Product field to given value.


### GetRegisteredAt

`func (o *Usage) GetRegisteredAt() time.Time`

GetRegisteredAt returns the RegisteredAt field if non-nil, zero value otherwise.

### GetRegisteredAtOk

`func (o *Usage) GetRegisteredAtOk() (*time.Time, bool)`

GetRegisteredAtOk returns a tuple with the RegisteredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredAt

`func (o *Usage) SetRegisteredAt(v time.Time)`

SetRegisteredAt sets RegisteredAt field to given value.


### GetUsageMeterItems

`func (o *Usage) GetUsageMeterItems() UsageMeterItem`

GetUsageMeterItems returns the UsageMeterItems field if non-nil, zero value otherwise.

### GetUsageMeterItemsOk

`func (o *Usage) GetUsageMeterItemsOk() (*UsageMeterItem, bool)`

GetUsageMeterItemsOk returns a tuple with the UsageMeterItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageMeterItems

`func (o *Usage) SetUsageMeterItems(v UsageMeterItem)`

SetUsageMeterItems sets UsageMeterItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


