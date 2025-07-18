# PatchApplicationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to [**Label**](Label.md) |  | [optional] 
**ApplicationsList** | Pointer to **[]string** | Provide a list of application ID to update. | [optional] 

## Methods

### NewPatchApplicationsRequest

`func NewPatchApplicationsRequest() *PatchApplicationsRequest`

NewPatchApplicationsRequest instantiates a new PatchApplicationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchApplicationsRequestWithDefaults

`func NewPatchApplicationsRequestWithDefaults() *PatchApplicationsRequest`

NewPatchApplicationsRequestWithDefaults instantiates a new PatchApplicationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *PatchApplicationsRequest) GetLabel() Label`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchApplicationsRequest) GetLabelOk() (*Label, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchApplicationsRequest) SetLabel(v Label)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchApplicationsRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetApplicationsList

`func (o *PatchApplicationsRequest) GetApplicationsList() []string`

GetApplicationsList returns the ApplicationsList field if non-nil, zero value otherwise.

### GetApplicationsListOk

`func (o *PatchApplicationsRequest) GetApplicationsListOk() (*[]string, bool)`

GetApplicationsListOk returns a tuple with the ApplicationsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationsList

`func (o *PatchApplicationsRequest) SetApplicationsList(v []string)`

SetApplicationsList sets ApplicationsList field to given value.

### HasApplicationsList

`func (o *PatchApplicationsRequest) HasApplicationsList() bool`

HasApplicationsList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


