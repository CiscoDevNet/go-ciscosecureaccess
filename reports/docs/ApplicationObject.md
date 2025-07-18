# ApplicationObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the app. | 
**Name** | **string** | The name of the app. | 
**Description** | **string** | The description of the app. | 
**Label** | [**Label**](Label.md) |  | 
**WeightedRisk** | [**WeightedRisk**](WeightedRisk.md) |  | 
**Category** | **string** | The category applied to the app. | 
**AppType** | [**AppType**](AppType.md) |  | 
**Url** | **string** | The URL of the app. | 
**Vendor** | **string** | The vendor that owns the app. | 
**IdentitiesCount** | **int64** | The number of identities. | 
**Sources** | [**[]SourcesInner**](SourcesInner.md) | The list of app sources where the sources are DNS, Web (Secure Web Gateway), and cloud-delivered firewall (CDFW) traffic events. The list can contain one or more of the source types. | 
**FirstDetected** | **time.Time** | The date when the app was first detected. | 
**LastDetected** | **time.Time** | The date when the app was last detected. | 

## Methods

### NewApplicationObject

`func NewApplicationObject(id string, name string, description string, label Label, weightedRisk WeightedRisk, category string, appType AppType, url string, vendor string, identitiesCount int64, sources []SourcesInner, firstDetected time.Time, lastDetected time.Time, ) *ApplicationObject`

NewApplicationObject instantiates a new ApplicationObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationObjectWithDefaults

`func NewApplicationObjectWithDefaults() *ApplicationObject`

NewApplicationObjectWithDefaults instantiates a new ApplicationObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationObject) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ApplicationObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationObject) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ApplicationObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationObject) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLabel

`func (o *ApplicationObject) GetLabel() Label`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ApplicationObject) GetLabelOk() (*Label, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ApplicationObject) SetLabel(v Label)`

SetLabel sets Label field to given value.


### GetWeightedRisk

`func (o *ApplicationObject) GetWeightedRisk() WeightedRisk`

GetWeightedRisk returns the WeightedRisk field if non-nil, zero value otherwise.

### GetWeightedRiskOk

`func (o *ApplicationObject) GetWeightedRiskOk() (*WeightedRisk, bool)`

GetWeightedRiskOk returns a tuple with the WeightedRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightedRisk

`func (o *ApplicationObject) SetWeightedRisk(v WeightedRisk)`

SetWeightedRisk sets WeightedRisk field to given value.


### GetCategory

`func (o *ApplicationObject) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ApplicationObject) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ApplicationObject) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetAppType

`func (o *ApplicationObject) GetAppType() AppType`

GetAppType returns the AppType field if non-nil, zero value otherwise.

### GetAppTypeOk

`func (o *ApplicationObject) GetAppTypeOk() (*AppType, bool)`

GetAppTypeOk returns a tuple with the AppType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppType

`func (o *ApplicationObject) SetAppType(v AppType)`

SetAppType sets AppType field to given value.


### GetUrl

`func (o *ApplicationObject) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApplicationObject) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApplicationObject) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetVendor

`func (o *ApplicationObject) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ApplicationObject) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ApplicationObject) SetVendor(v string)`

SetVendor sets Vendor field to given value.


### GetIdentitiesCount

`func (o *ApplicationObject) GetIdentitiesCount() int64`

GetIdentitiesCount returns the IdentitiesCount field if non-nil, zero value otherwise.

### GetIdentitiesCountOk

`func (o *ApplicationObject) GetIdentitiesCountOk() (*int64, bool)`

GetIdentitiesCountOk returns a tuple with the IdentitiesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentitiesCount

`func (o *ApplicationObject) SetIdentitiesCount(v int64)`

SetIdentitiesCount sets IdentitiesCount field to given value.


### GetSources

`func (o *ApplicationObject) GetSources() []SourcesInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *ApplicationObject) GetSourcesOk() (*[]SourcesInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *ApplicationObject) SetSources(v []SourcesInner)`

SetSources sets Sources field to given value.


### GetFirstDetected

`func (o *ApplicationObject) GetFirstDetected() time.Time`

GetFirstDetected returns the FirstDetected field if non-nil, zero value otherwise.

### GetFirstDetectedOk

`func (o *ApplicationObject) GetFirstDetectedOk() (*time.Time, bool)`

GetFirstDetectedOk returns a tuple with the FirstDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstDetected

`func (o *ApplicationObject) SetFirstDetected(v time.Time)`

SetFirstDetected sets FirstDetected field to given value.


### GetLastDetected

`func (o *ApplicationObject) GetLastDetected() time.Time`

GetLastDetected returns the LastDetected field if non-nil, zero value otherwise.

### GetLastDetectedOk

`func (o *ApplicationObject) GetLastDetectedOk() (*time.Time, bool)`

GetLastDetectedOk returns a tuple with the LastDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDetected

`func (o *ApplicationObject) SetLastDetected(v time.Time)`

SetLastDetected sets LastDetected field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


