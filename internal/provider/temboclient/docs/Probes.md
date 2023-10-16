# Probes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Liveness** | [**Probe**](Probe.md) |  | 
**Readiness** | [**Probe**](Probe.md) |  | 

## Methods

### NewProbes

`func NewProbes(liveness Probe, readiness Probe, ) *Probes`

NewProbes instantiates a new Probes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProbesWithDefaults

`func NewProbesWithDefaults() *Probes`

NewProbesWithDefaults instantiates a new Probes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLiveness

`func (o *Probes) GetLiveness() Probe`

GetLiveness returns the Liveness field if non-nil, zero value otherwise.

### GetLivenessOk

`func (o *Probes) GetLivenessOk() (*Probe, bool)`

GetLivenessOk returns a tuple with the Liveness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveness

`func (o *Probes) SetLiveness(v Probe)`

SetLiveness sets Liveness field to given value.


### GetReadiness

`func (o *Probes) GetReadiness() Probe`

GetReadiness returns the Readiness field if non-nil, zero value otherwise.

### GetReadinessOk

`func (o *Probes) GetReadinessOk() (*Probe, bool)`

GetReadinessOk returns a tuple with the Readiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadiness

`func (o *Probes) SetReadiness(v Probe)`

SetReadiness sets Readiness field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


