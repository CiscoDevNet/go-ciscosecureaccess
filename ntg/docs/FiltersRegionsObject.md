# FiltersRegionsObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeerIP** | Pointer to **string** | Provide the public IP address of a peer (tunnel). List the regions sorted in ascending order and based on the regional distances from the geolocation of the provided peer IP. Only a public IP is allowed as the value of &#x60;peerIP&#x60;. &#x60;peerIP&#x60; is ignored if you provide &#x60;latitude&#x60; and &#x60;longitude&#x60;.&lt;/br&gt; **Note:** If you set the &#x60;latitude&#x60; and &#x60;longitude&#x60; filters, the &#x60;peerIP&#x60; filter is ignored. | [optional] 
**Latitude** | Pointer to **string** | List the regions in ascending order based on the distance of the regions from the provided coordinates. **Note:** Set both the &#x60;latitude&#x60; and &#x60;longitude&#x60; filters together. | [optional] 
**Longitude** | Pointer to **string** | List the regions in ascending order based on the distance of the regions from the provided coordinates. | [optional] 
**Status** | Pointer to **string** | Specify either all regions or only the regions that can accept new tunnels. Set the value of the &#x60;status&#x60; filter as &#x60;all&#x60; or &#x60;available&#x60;. The default value is &#x60;available&#x60;. | [optional] 

## Methods

### NewFiltersRegionsObject

`func NewFiltersRegionsObject() *FiltersRegionsObject`

NewFiltersRegionsObject instantiates a new FiltersRegionsObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersRegionsObjectWithDefaults

`func NewFiltersRegionsObjectWithDefaults() *FiltersRegionsObject`

NewFiltersRegionsObjectWithDefaults instantiates a new FiltersRegionsObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeerIP

`func (o *FiltersRegionsObject) GetPeerIP() string`

GetPeerIP returns the PeerIP field if non-nil, zero value otherwise.

### GetPeerIPOk

`func (o *FiltersRegionsObject) GetPeerIPOk() (*string, bool)`

GetPeerIPOk returns a tuple with the PeerIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerIP

`func (o *FiltersRegionsObject) SetPeerIP(v string)`

SetPeerIP sets PeerIP field to given value.

### HasPeerIP

`func (o *FiltersRegionsObject) HasPeerIP() bool`

HasPeerIP returns a boolean if a field has been set.

### GetLatitude

`func (o *FiltersRegionsObject) GetLatitude() string`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *FiltersRegionsObject) GetLatitudeOk() (*string, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *FiltersRegionsObject) SetLatitude(v string)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *FiltersRegionsObject) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *FiltersRegionsObject) GetLongitude() string`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *FiltersRegionsObject) GetLongitudeOk() (*string, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *FiltersRegionsObject) SetLongitude(v string)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *FiltersRegionsObject) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetStatus

`func (o *FiltersRegionsObject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FiltersRegionsObject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FiltersRegionsObject) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FiltersRegionsObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


