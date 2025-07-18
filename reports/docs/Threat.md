# Threat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | The descriptive label for the threat name. | [optional] 
**Type** | Pointer to **string** | The type of the threat. | [optional] 

## Methods

### NewThreat

`func NewThreat() *Threat`

NewThreat instantiates a new Threat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatWithDefaults

`func NewThreatWithDefaults() *Threat`

NewThreatWithDefaults instantiates a new Threat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *Threat) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Threat) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Threat) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Threat) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *Threat) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Threat) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Threat) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Threat) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


