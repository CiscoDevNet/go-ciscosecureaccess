# SignatureListSummarySignaturesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Generatorid** | Pointer to **int64** | The generator ID of the signature. | [optional] 
**Id** | Pointer to **int64** | The ID of the signature. | [optional] 
**Lasteventat** | Pointer to **int64** | The date and time in milliseconds of the last event. | [optional] 
**Counts** | Pointer to [**SignatureListSummarySignaturesInnerCounts**](SignatureListSummarySignaturesInnerCounts.md) |  | [optional] 

## Methods

### NewSignatureListSummarySignaturesInner

`func NewSignatureListSummarySignaturesInner() *SignatureListSummarySignaturesInner`

NewSignatureListSummarySignaturesInner instantiates a new SignatureListSummarySignaturesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureListSummarySignaturesInnerWithDefaults

`func NewSignatureListSummarySignaturesInnerWithDefaults() *SignatureListSummarySignaturesInner`

NewSignatureListSummarySignaturesInnerWithDefaults instantiates a new SignatureListSummarySignaturesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneratorid

`func (o *SignatureListSummarySignaturesInner) GetGeneratorid() int64`

GetGeneratorid returns the Generatorid field if non-nil, zero value otherwise.

### GetGeneratoridOk

`func (o *SignatureListSummarySignaturesInner) GetGeneratoridOk() (*int64, bool)`

GetGeneratoridOk returns a tuple with the Generatorid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratorid

`func (o *SignatureListSummarySignaturesInner) SetGeneratorid(v int64)`

SetGeneratorid sets Generatorid field to given value.

### HasGeneratorid

`func (o *SignatureListSummarySignaturesInner) HasGeneratorid() bool`

HasGeneratorid returns a boolean if a field has been set.

### GetId

`func (o *SignatureListSummarySignaturesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SignatureListSummarySignaturesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SignatureListSummarySignaturesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SignatureListSummarySignaturesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLasteventat

`func (o *SignatureListSummarySignaturesInner) GetLasteventat() int64`

GetLasteventat returns the Lasteventat field if non-nil, zero value otherwise.

### GetLasteventatOk

`func (o *SignatureListSummarySignaturesInner) GetLasteventatOk() (*int64, bool)`

GetLasteventatOk returns a tuple with the Lasteventat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLasteventat

`func (o *SignatureListSummarySignaturesInner) SetLasteventat(v int64)`

SetLasteventat sets Lasteventat field to given value.

### HasLasteventat

`func (o *SignatureListSummarySignaturesInner) HasLasteventat() bool`

HasLasteventat returns a boolean if a field has been set.

### GetCounts

`func (o *SignatureListSummarySignaturesInner) GetCounts() SignatureListSummarySignaturesInnerCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *SignatureListSummarySignaturesInner) GetCountsOk() (*SignatureListSummarySignaturesInnerCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *SignatureListSummarySignaturesInner) SetCounts(v SignatureListSummarySignaturesInnerCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *SignatureListSummarySignaturesInner) HasCounts() bool`

HasCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


