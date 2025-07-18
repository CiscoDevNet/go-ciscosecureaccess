# RegionsResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinentName** | Pointer to **string** | The name of the continent. | [optional] 
**Countries** | Pointer to [**[]RegionsResultsInnerCountriesInner**](RegionsResultsInnerCountriesInner.md) | The list of countries in the continent. | [optional] 

## Methods

### NewRegionsResultsInner

`func NewRegionsResultsInner() *RegionsResultsInner`

NewRegionsResultsInner instantiates a new RegionsResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionsResultsInnerWithDefaults

`func NewRegionsResultsInnerWithDefaults() *RegionsResultsInner`

NewRegionsResultsInnerWithDefaults instantiates a new RegionsResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinentName

`func (o *RegionsResultsInner) GetContinentName() string`

GetContinentName returns the ContinentName field if non-nil, zero value otherwise.

### GetContinentNameOk

`func (o *RegionsResultsInner) GetContinentNameOk() (*string, bool)`

GetContinentNameOk returns a tuple with the ContinentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinentName

`func (o *RegionsResultsInner) SetContinentName(v string)`

SetContinentName sets ContinentName field to given value.

### HasContinentName

`func (o *RegionsResultsInner) HasContinentName() bool`

HasContinentName returns a boolean if a field has been set.

### GetCountries

`func (o *RegionsResultsInner) GetCountries() []RegionsResultsInnerCountriesInner`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *RegionsResultsInner) GetCountriesOk() (*[]RegionsResultsInnerCountriesInner, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *RegionsResultsInner) SetCountries(v []RegionsResultsInnerCountriesInner)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *RegionsResultsInner) HasCountries() bool`

HasCountries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


