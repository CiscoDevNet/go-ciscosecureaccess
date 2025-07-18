# GetUsageMetrics200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | **int64** |  | 
**Body** | [**GetUsageMetrics200ResponseBody**](GetUsageMetrics200ResponseBody.md) |  | 
**Next** | Pointer to **string** | The URL of the next set of usage metric records in the collection. \&quot;https://api.sse.cisco.com/reports/v2/usage/metrics?from&#x3D;2022-07-15T20:00:00.000&amp;to&#x3D;2022-07-16T21:00:00.000\&quot; | [optional] 

## Methods

### NewGetUsageMetrics200Response

`func NewGetUsageMetrics200Response(statusCode int64, body GetUsageMetrics200ResponseBody, ) *GetUsageMetrics200Response`

NewGetUsageMetrics200Response instantiates a new GetUsageMetrics200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUsageMetrics200ResponseWithDefaults

`func NewGetUsageMetrics200ResponseWithDefaults() *GetUsageMetrics200Response`

NewGetUsageMetrics200ResponseWithDefaults instantiates a new GetUsageMetrics200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *GetUsageMetrics200Response) GetStatusCode() int64`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *GetUsageMetrics200Response) GetStatusCodeOk() (*int64, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *GetUsageMetrics200Response) SetStatusCode(v int64)`

SetStatusCode sets StatusCode field to given value.


### GetBody

`func (o *GetUsageMetrics200Response) GetBody() GetUsageMetrics200ResponseBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *GetUsageMetrics200Response) GetBodyOk() (*GetUsageMetrics200ResponseBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *GetUsageMetrics200Response) SetBody(v GetUsageMetrics200ResponseBody)`

SetBody sets Body field to given value.


### GetNext

`func (o *GetUsageMetrics200Response) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *GetUsageMetrics200Response) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *GetUsageMetrics200Response) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *GetUsageMetrics200Response) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


