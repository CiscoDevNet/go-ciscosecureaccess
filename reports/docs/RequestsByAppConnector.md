# RequestsByAppConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the Resource Connector. | 
**Requests** | [**RequestSummaryAppConnector**](RequestSummaryAppConnector.md) |  | 

## Methods

### NewRequestsByAppConnector

`func NewRequestsByAppConnector(id int64, requests RequestSummaryAppConnector, ) *RequestsByAppConnector`

NewRequestsByAppConnector instantiates a new RequestsByAppConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestsByAppConnectorWithDefaults

`func NewRequestsByAppConnectorWithDefaults() *RequestsByAppConnector`

NewRequestsByAppConnectorWithDefaults instantiates a new RequestsByAppConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequestsByAppConnector) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestsByAppConnector) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestsByAppConnector) SetId(v int64)`

SetId sets Id field to given value.


### GetRequests

`func (o *RequestsByAppConnector) GetRequests() RequestSummaryAppConnector`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *RequestsByAppConnector) GetRequestsOk() (*RequestSummaryAppConnector, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *RequestsByAppConnector) SetRequests(v RequestSummaryAppConnector)`

SetRequests sets Requests field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


