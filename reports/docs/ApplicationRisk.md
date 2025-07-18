# ApplicationRisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of app. | [optional] 
**Name** | Pointer to **string** | The name of the app. | [optional] 
**WeightedRisk** | Pointer to [**WeightedRisk**](WeightedRisk.md) |  | [optional] 
**BusinessRisk** | Pointer to [**BusinessRisk**](BusinessRisk.md) |  | [optional] 
**UsageType** | Pointer to [**UsageType**](UsageType.md) |  | [optional] 
**WebReputation** | Pointer to **int64** | Provides accurate conclusions about a given host by tracking a broad set of attributes and using sophisticated security modeling. Powered by Talos Security Intelligence. | [optional] 
**FinancialViability** | Pointer to [**FinancialViability**](FinancialViability.md) |  | [optional] 
**DataStorage** | Pointer to [**DataStorage**](DataStorage.md) |  | [optional] 
**VendorCompliance** | Pointer to [**[]ApplicationRiskVendorComplianceInner**](ApplicationRiskVendorComplianceInner.md) | The vendor&#39;s compliance information. | [optional] 

## Methods

### NewApplicationRisk

`func NewApplicationRisk() *ApplicationRisk`

NewApplicationRisk instantiates a new ApplicationRisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRiskWithDefaults

`func NewApplicationRiskWithDefaults() *ApplicationRisk`

NewApplicationRiskWithDefaults instantiates a new ApplicationRisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationRisk) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationRisk) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationRisk) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationRisk) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApplicationRisk) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationRisk) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationRisk) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationRisk) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWeightedRisk

`func (o *ApplicationRisk) GetWeightedRisk() WeightedRisk`

GetWeightedRisk returns the WeightedRisk field if non-nil, zero value otherwise.

### GetWeightedRiskOk

`func (o *ApplicationRisk) GetWeightedRiskOk() (*WeightedRisk, bool)`

GetWeightedRiskOk returns a tuple with the WeightedRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightedRisk

`func (o *ApplicationRisk) SetWeightedRisk(v WeightedRisk)`

SetWeightedRisk sets WeightedRisk field to given value.

### HasWeightedRisk

`func (o *ApplicationRisk) HasWeightedRisk() bool`

HasWeightedRisk returns a boolean if a field has been set.

### GetBusinessRisk

`func (o *ApplicationRisk) GetBusinessRisk() BusinessRisk`

GetBusinessRisk returns the BusinessRisk field if non-nil, zero value otherwise.

### GetBusinessRiskOk

`func (o *ApplicationRisk) GetBusinessRiskOk() (*BusinessRisk, bool)`

GetBusinessRiskOk returns a tuple with the BusinessRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessRisk

`func (o *ApplicationRisk) SetBusinessRisk(v BusinessRisk)`

SetBusinessRisk sets BusinessRisk field to given value.

### HasBusinessRisk

`func (o *ApplicationRisk) HasBusinessRisk() bool`

HasBusinessRisk returns a boolean if a field has been set.

### GetUsageType

`func (o *ApplicationRisk) GetUsageType() UsageType`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *ApplicationRisk) GetUsageTypeOk() (*UsageType, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *ApplicationRisk) SetUsageType(v UsageType)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *ApplicationRisk) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### GetWebReputation

`func (o *ApplicationRisk) GetWebReputation() int64`

GetWebReputation returns the WebReputation field if non-nil, zero value otherwise.

### GetWebReputationOk

`func (o *ApplicationRisk) GetWebReputationOk() (*int64, bool)`

GetWebReputationOk returns a tuple with the WebReputation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebReputation

`func (o *ApplicationRisk) SetWebReputation(v int64)`

SetWebReputation sets WebReputation field to given value.

### HasWebReputation

`func (o *ApplicationRisk) HasWebReputation() bool`

HasWebReputation returns a boolean if a field has been set.

### GetFinancialViability

`func (o *ApplicationRisk) GetFinancialViability() FinancialViability`

GetFinancialViability returns the FinancialViability field if non-nil, zero value otherwise.

### GetFinancialViabilityOk

`func (o *ApplicationRisk) GetFinancialViabilityOk() (*FinancialViability, bool)`

GetFinancialViabilityOk returns a tuple with the FinancialViability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialViability

`func (o *ApplicationRisk) SetFinancialViability(v FinancialViability)`

SetFinancialViability sets FinancialViability field to given value.

### HasFinancialViability

`func (o *ApplicationRisk) HasFinancialViability() bool`

HasFinancialViability returns a boolean if a field has been set.

### GetDataStorage

`func (o *ApplicationRisk) GetDataStorage() DataStorage`

GetDataStorage returns the DataStorage field if non-nil, zero value otherwise.

### GetDataStorageOk

`func (o *ApplicationRisk) GetDataStorageOk() (*DataStorage, bool)`

GetDataStorageOk returns a tuple with the DataStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStorage

`func (o *ApplicationRisk) SetDataStorage(v DataStorage)`

SetDataStorage sets DataStorage field to given value.

### HasDataStorage

`func (o *ApplicationRisk) HasDataStorage() bool`

HasDataStorage returns a boolean if a field has been set.

### GetVendorCompliance

`func (o *ApplicationRisk) GetVendorCompliance() []ApplicationRiskVendorComplianceInner`

GetVendorCompliance returns the VendorCompliance field if non-nil, zero value otherwise.

### GetVendorComplianceOk

`func (o *ApplicationRisk) GetVendorComplianceOk() (*[]ApplicationRiskVendorComplianceInner, bool)`

GetVendorComplianceOk returns a tuple with the VendorCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCompliance

`func (o *ApplicationRisk) SetVendorCompliance(v []ApplicationRiskVendorComplianceInner)`

SetVendorCompliance sets VendorCompliance field to given value.

### HasVendorCompliance

`func (o *ApplicationRisk) HasVendorCompliance() bool`

HasVendorCompliance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


