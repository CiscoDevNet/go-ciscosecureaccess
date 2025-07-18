# ConnectorGroupPatchReqInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | [**Op**](Op.md) |  | 
**Path** | **string** | The Resource Connector Group property. You can set the following properties on the Connector Group: &#x60;name&#x60;, &#x60;confirmedAgentsEnabled&#x60;, &#x60;provisioningKey&#x60;, and &#x60;resourceIds&#x60;. | 
**Value** | **string** | The value of the property. When the value of the &#x60;path&#x60; field is &#x60;provisioningKey&#x60;, the &#x60;value&#x60; field accepts an empty string or is not required. | 

## Methods

### NewConnectorGroupPatchReqInner

`func NewConnectorGroupPatchReqInner(op Op, path string, value string, ) *ConnectorGroupPatchReqInner`

NewConnectorGroupPatchReqInner instantiates a new ConnectorGroupPatchReqInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorGroupPatchReqInnerWithDefaults

`func NewConnectorGroupPatchReqInnerWithDefaults() *ConnectorGroupPatchReqInner`

NewConnectorGroupPatchReqInnerWithDefaults instantiates a new ConnectorGroupPatchReqInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *ConnectorGroupPatchReqInner) GetOp() Op`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *ConnectorGroupPatchReqInner) GetOpOk() (*Op, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *ConnectorGroupPatchReqInner) SetOp(v Op)`

SetOp sets Op field to given value.


### GetPath

`func (o *ConnectorGroupPatchReqInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ConnectorGroupPatchReqInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ConnectorGroupPatchReqInner) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *ConnectorGroupPatchReqInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConnectorGroupPatchReqInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConnectorGroupPatchReqInner) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


