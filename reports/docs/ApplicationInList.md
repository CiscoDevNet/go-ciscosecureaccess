# ApplicationInList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the app. | 
**Name** | **string** | The name of the app. | 
**Label** | [**Label**](Label.md) |  | 
**WeightedRisk** | [**WeightedRisk**](WeightedRisk.md) |  | 
**Category** | **string** | The category of the app. | 
**AppType** | [**AppType**](AppType.md) |  | 
**Sources** | [**[]SourcesInner**](SourcesInner.md) | The list of app sources where the sources are DNS, Web (Secure Web Gateway), and cloud-delivered firewall (CDFW) traffic events. The list can contain one or more of the source types. | 
**FirstDetected** | **time.Time** | The date when the app was first detected. | 
**LastDetected** | **time.Time** | The date when the app was last detected. | 

## Methods

### NewApplicationInList

`func NewApplicationInList(id string, name string, label Label, weightedRisk WeightedRisk, category string, appType AppType, sources []SourcesInner, firstDetected time.Time, lastDetected time.Time, ) *ApplicationInList`

NewApplicationInList instantiates a new ApplicationInList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInListWithDefaults

`func NewApplicationInListWithDefaults() *ApplicationInList`

NewApplicationInListWithDefaults instantiates a new ApplicationInList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationInList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationInList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationInList) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ApplicationInList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationInList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationInList) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *ApplicationInList) GetLabel() Label`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ApplicationInList) GetLabelOk() (*Label, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ApplicationInList) SetLabel(v Label)`

SetLabel sets Label field to given value.


### GetWeightedRisk

`func (o *ApplicationInList) GetWeightedRisk() WeightedRisk`

GetWeightedRisk returns the WeightedRisk field if non-nil, zero value otherwise.

### GetWeightedRiskOk

`func (o *ApplicationInList) GetWeightedRiskOk() (*WeightedRisk, bool)`

GetWeightedRiskOk returns a tuple with the WeightedRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightedRisk

`func (o *ApplicationInList) SetWeightedRisk(v WeightedRisk)`

SetWeightedRisk sets WeightedRisk field to given value.


### GetCategory

`func (o *ApplicationInList) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ApplicationInList) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ApplicationInList) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetAppType

`func (o *ApplicationInList) GetAppType() AppType`

GetAppType returns the AppType field if non-nil, zero value otherwise.

### GetAppTypeOk

`func (o *ApplicationInList) GetAppTypeOk() (*AppType, bool)`

GetAppTypeOk returns a tuple with the AppType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppType

`func (o *ApplicationInList) SetAppType(v AppType)`

SetAppType sets AppType field to given value.


### GetSources

`func (o *ApplicationInList) GetSources() []SourcesInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *ApplicationInList) GetSourcesOk() (*[]SourcesInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *ApplicationInList) SetSources(v []SourcesInner)`

SetSources sets Sources field to given value.


### GetFirstDetected

`func (o *ApplicationInList) GetFirstDetected() time.Time`

GetFirstDetected returns the FirstDetected field if non-nil, zero value otherwise.

### GetFirstDetectedOk

`func (o *ApplicationInList) GetFirstDetectedOk() (*time.Time, bool)`

GetFirstDetectedOk returns a tuple with the FirstDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstDetected

`func (o *ApplicationInList) SetFirstDetected(v time.Time)`

SetFirstDetected sets FirstDetected field to given value.


### GetLastDetected

`func (o *ApplicationInList) GetLastDetected() time.Time`

GetLastDetected returns the LastDetected field if non-nil, zero value otherwise.

### GetLastDetectedOk

`func (o *ApplicationInList) GetLastDetectedOk() (*time.Time, bool)`

GetLastDetectedOk returns a tuple with the LastDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDetected

`func (o *ApplicationInList) SetLastDetected(v time.Time)`

SetLastDetected sets LastDetected field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


