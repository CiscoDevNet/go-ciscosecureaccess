# ApplicationIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the identity. | 
**Name** | **string** | The name of the identity. | 
**Sources** | [**[]SourcesInner**](SourcesInner.md) | The list of app sources where the sources are DNS, Web (Secure Web Gateway), and cloud-delivered firewall (CDFW) traffic events. The list can contain one or more of the source types. | 
**FirstDetected** | **time.Time** | The date when the identity was first detected for the application. | 
**LastDetected** | **time.Time** | The date when the identity was last detected for the application. | 

## Methods

### NewApplicationIdentity

`func NewApplicationIdentity(id int64, name string, sources []SourcesInner, firstDetected time.Time, lastDetected time.Time, ) *ApplicationIdentity`

NewApplicationIdentity instantiates a new ApplicationIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationIdentityWithDefaults

`func NewApplicationIdentityWithDefaults() *ApplicationIdentity`

NewApplicationIdentityWithDefaults instantiates a new ApplicationIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationIdentity) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationIdentity) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationIdentity) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ApplicationIdentity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationIdentity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationIdentity) SetName(v string)`

SetName sets Name field to given value.


### GetSources

`func (o *ApplicationIdentity) GetSources() []SourcesInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *ApplicationIdentity) GetSourcesOk() (*[]SourcesInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *ApplicationIdentity) SetSources(v []SourcesInner)`

SetSources sets Sources field to given value.


### GetFirstDetected

`func (o *ApplicationIdentity) GetFirstDetected() time.Time`

GetFirstDetected returns the FirstDetected field if non-nil, zero value otherwise.

### GetFirstDetectedOk

`func (o *ApplicationIdentity) GetFirstDetectedOk() (*time.Time, bool)`

GetFirstDetectedOk returns a tuple with the FirstDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstDetected

`func (o *ApplicationIdentity) SetFirstDetected(v time.Time)`

SetFirstDetected sets FirstDetected field to given value.


### GetLastDetected

`func (o *ApplicationIdentity) GetLastDetected() time.Time`

GetLastDetected returns the LastDetected field if non-nil, zero value otherwise.

### GetLastDetectedOk

`func (o *ApplicationIdentity) GetLastDetectedOk() (*time.Time, bool)`

GetLastDetectedOk returns a tuple with the LastDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDetected

`func (o *ApplicationIdentity) SetLastDetected(v time.Time)`

SetLastDetected sets LastDetected field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


