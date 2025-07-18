# BgpDataResponseObj

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsNumber** | **string** | The border gateway protocol (BGP) autonomous system (AS) number for private access network tunnels. Only required for the &#x60;bgp&#x60; routing type. Any other routing types except &#x60;bgp&#x60; are ignored. Specify an integer between 0–65536. | 

## Methods

### NewBgpDataResponseObj

`func NewBgpDataResponseObj(asNumber string, ) *BgpDataResponseObj`

NewBgpDataResponseObj instantiates a new BgpDataResponseObj object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpDataResponseObjWithDefaults

`func NewBgpDataResponseObjWithDefaults() *BgpDataResponseObj`

NewBgpDataResponseObjWithDefaults instantiates a new BgpDataResponseObj object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsNumber

`func (o *BgpDataResponseObj) GetAsNumber() string`

GetAsNumber returns the AsNumber field if non-nil, zero value otherwise.

### GetAsNumberOk

`func (o *BgpDataResponseObj) GetAsNumberOk() (*string, bool)`

GetAsNumberOk returns a tuple with the AsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsNumber

`func (o *BgpDataResponseObj) SetAsNumber(v string)`

SetAsNumber sets AsNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


