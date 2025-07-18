# ConnectorPatchReqInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | Pointer to [**Op**](Op.md) |  | [optional] 
**Path** | Pointer to **string** | The Connector property. You can set the following properties on the Connector: &#x60;enabled&#x60;, &#x60;confirmed&#x60;, and &#x60;revoked&#x60;. | [optional] 
**Value** | Pointer to **bool** | The value of the property. When the value of the &#x60;path&#x60; field is &#x60;confirmed&#x60; or &#x60;revoked&#x60;, set the &#x60;value&#x60; property to &#x60;true&#x60;. | [optional] 

## Methods

### NewConnectorPatchReqInner

`func NewConnectorPatchReqInner() *ConnectorPatchReqInner`

NewConnectorPatchReqInner instantiates a new ConnectorPatchReqInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorPatchReqInnerWithDefaults

`func NewConnectorPatchReqInnerWithDefaults() *ConnectorPatchReqInner`

NewConnectorPatchReqInnerWithDefaults instantiates a new ConnectorPatchReqInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *ConnectorPatchReqInner) GetOp() Op`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *ConnectorPatchReqInner) GetOpOk() (*Op, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *ConnectorPatchReqInner) SetOp(v Op)`

SetOp sets Op field to given value.

### HasOp

`func (o *ConnectorPatchReqInner) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetPath

`func (o *ConnectorPatchReqInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ConnectorPatchReqInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ConnectorPatchReqInner) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ConnectorPatchReqInner) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetValue

`func (o *ConnectorPatchReqInner) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConnectorPatchReqInner) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConnectorPatchReqInner) SetValue(v bool)`

SetValue sets Value field to given value.

### HasValue

`func (o *ConnectorPatchReqInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


