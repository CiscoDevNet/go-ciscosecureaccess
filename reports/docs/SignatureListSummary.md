# SignatureListSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signaturelist** | Pointer to [**SignatureList**](SignatureList.md) |  | [optional] 
**Signatures** | Pointer to [**[]SignatureListSummarySignaturesInner**](SignatureListSummarySignaturesInner.md) | The list of information about the signatures. | [optional] 
**Rule** | Pointer to [**RuleParam**](RuleParam.md) |  | [optional] 

## Methods

### NewSignatureListSummary

`func NewSignatureListSummary() *SignatureListSummary`

NewSignatureListSummary instantiates a new SignatureListSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureListSummaryWithDefaults

`func NewSignatureListSummaryWithDefaults() *SignatureListSummary`

NewSignatureListSummaryWithDefaults instantiates a new SignatureListSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignaturelist

`func (o *SignatureListSummary) GetSignaturelist() SignatureList`

GetSignaturelist returns the Signaturelist field if non-nil, zero value otherwise.

### GetSignaturelistOk

`func (o *SignatureListSummary) GetSignaturelistOk() (*SignatureList, bool)`

GetSignaturelistOk returns a tuple with the Signaturelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignaturelist

`func (o *SignatureListSummary) SetSignaturelist(v SignatureList)`

SetSignaturelist sets Signaturelist field to given value.

### HasSignaturelist

`func (o *SignatureListSummary) HasSignaturelist() bool`

HasSignaturelist returns a boolean if a field has been set.

### GetSignatures

`func (o *SignatureListSummary) GetSignatures() []SignatureListSummarySignaturesInner`

GetSignatures returns the Signatures field if non-nil, zero value otherwise.

### GetSignaturesOk

`func (o *SignatureListSummary) GetSignaturesOk() (*[]SignatureListSummarySignaturesInner, bool)`

GetSignaturesOk returns a tuple with the Signatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatures

`func (o *SignatureListSummary) SetSignatures(v []SignatureListSummarySignaturesInner)`

SetSignatures sets Signatures field to given value.

### HasSignatures

`func (o *SignatureListSummary) HasSignatures() bool`

HasSignatures returns a boolean if a field has been set.

### GetRule

`func (o *SignatureListSummary) GetRule() RuleParam`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *SignatureListSummary) GetRuleOk() (*RuleParam, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *SignatureListSummary) SetRule(v RuleParam)`

SetRule sets Rule field to given value.

### HasRule

`func (o *SignatureListSummary) HasRule() bool`

HasRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


