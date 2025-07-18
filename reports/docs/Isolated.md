# Isolated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | The state of the isolated file. | 
**Fileaction** | **string** | The action taken for the file. | 

## Methods

### NewIsolated

`func NewIsolated(state string, fileaction string, ) *Isolated`

NewIsolated instantiates a new Isolated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIsolatedWithDefaults

`func NewIsolatedWithDefaults() *Isolated`

NewIsolatedWithDefaults instantiates a new Isolated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *Isolated) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Isolated) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Isolated) SetState(v string)`

SetState sets State field to given value.


### GetFileaction

`func (o *Isolated) GetFileaction() string`

GetFileaction returns the Fileaction field if non-nil, zero value otherwise.

### GetFileactionOk

`func (o *Isolated) GetFileactionOk() (*string, bool)`

GetFileactionOk returns a tuple with the Fileaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileaction

`func (o *Isolated) SetFileaction(v string)`

SetFileaction sets Fileaction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


