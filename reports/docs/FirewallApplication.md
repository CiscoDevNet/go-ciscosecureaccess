# FirewallApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the application or protocol. | [optional] 
**Label** | Pointer to **string** | The descriptive label for the application or protocol. | [optional] 
**App** | Pointer to **string** | The information about the app type. | [optional] 

## Methods

### NewFirewallApplication

`func NewFirewallApplication() *FirewallApplication`

NewFirewallApplication instantiates a new FirewallApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallApplicationWithDefaults

`func NewFirewallApplicationWithDefaults() *FirewallApplication`

NewFirewallApplicationWithDefaults instantiates a new FirewallApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FirewallApplication) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirewallApplication) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirewallApplication) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FirewallApplication) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *FirewallApplication) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FirewallApplication) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FirewallApplication) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FirewallApplication) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetApp

`func (o *FirewallApplication) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *FirewallApplication) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *FirewallApplication) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *FirewallApplication) HasApp() bool`

HasApp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


