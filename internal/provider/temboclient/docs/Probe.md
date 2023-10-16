# Probe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InitialDelaySeconds** | **int32** |  | 
**Path** | **string** |  | 
**Port** | **string** |  | 

## Methods

### NewProbe

`func NewProbe(initialDelaySeconds int32, path string, port string, ) *Probe`

NewProbe instantiates a new Probe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProbeWithDefaults

`func NewProbeWithDefaults() *Probe`

NewProbeWithDefaults instantiates a new Probe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitialDelaySeconds

`func (o *Probe) GetInitialDelaySeconds() int32`

GetInitialDelaySeconds returns the InitialDelaySeconds field if non-nil, zero value otherwise.

### GetInitialDelaySecondsOk

`func (o *Probe) GetInitialDelaySecondsOk() (*int32, bool)`

GetInitialDelaySecondsOk returns a tuple with the InitialDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialDelaySeconds

`func (o *Probe) SetInitialDelaySeconds(v int32)`

SetInitialDelaySeconds sets InitialDelaySeconds field to given value.


### GetPath

`func (o *Probe) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Probe) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Probe) SetPath(v string)`

SetPath sets Path field to given value.


### GetPort

`func (o *Probe) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Probe) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Probe) SetPort(v string)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


