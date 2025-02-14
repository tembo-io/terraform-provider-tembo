# AppMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | path to scrape metrics | 
**Port** | **int32** | port must be also exposed in one of AppService.routing[] | 

## Methods

### NewAppMetrics

`func NewAppMetrics(path string, port int32, ) *AppMetrics`

NewAppMetrics instantiates a new AppMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppMetricsWithDefaults

`func NewAppMetricsWithDefaults() *AppMetrics`

NewAppMetricsWithDefaults instantiates a new AppMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *AppMetrics) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AppMetrics) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AppMetrics) SetPath(v string)`

SetPath sets Path field to given value.


### GetPort

`func (o *AppMetrics) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AppMetrics) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AppMetrics) SetPort(v int32)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


