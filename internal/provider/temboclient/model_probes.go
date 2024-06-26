/*
Tembo Cloud

Platform API for Tembo Cloud             </br>             </br>             To find a Tembo Data API, please find it here:             </br>             </br>             [AWS US East 1](https://api.data-1.use1.tembo.io/swagger-ui/)             

API version: v1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package temboclient

import (
	"encoding/json"
	"fmt"
)

// checks if the Probes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Probes{}

// Probes Probes are used to determine the health of a container. You define this in the same manner as you would for all Kubernetes containers. See the [Kubernetes docs](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/).
type Probes struct {
	Liveness Probe `json:"liveness"`
	Readiness Probe `json:"readiness"`
	AdditionalProperties map[string]interface{}
}

type _Probes Probes

// NewProbes instantiates a new Probes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProbes(liveness Probe, readiness Probe) *Probes {
	this := Probes{}
	this.Liveness = liveness
	this.Readiness = readiness
	return &this
}

// NewProbesWithDefaults instantiates a new Probes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProbesWithDefaults() *Probes {
	this := Probes{}
	return &this
}

// GetLiveness returns the Liveness field value
func (o *Probes) GetLiveness() Probe {
	if o == nil {
		var ret Probe
		return ret
	}

	return o.Liveness
}

// GetLivenessOk returns a tuple with the Liveness field value
// and a boolean to check if the value has been set.
func (o *Probes) GetLivenessOk() (*Probe, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Liveness, true
}

// SetLiveness sets field value
func (o *Probes) SetLiveness(v Probe) {
	o.Liveness = v
}

// GetReadiness returns the Readiness field value
func (o *Probes) GetReadiness() Probe {
	if o == nil {
		var ret Probe
		return ret
	}

	return o.Readiness
}

// GetReadinessOk returns a tuple with the Readiness field value
// and a boolean to check if the value has been set.
func (o *Probes) GetReadinessOk() (*Probe, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Readiness, true
}

// SetReadiness sets field value
func (o *Probes) SetReadiness(v Probe) {
	o.Readiness = v
}

func (o Probes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Probes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["liveness"] = o.Liveness
	toSerialize["readiness"] = o.Readiness

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Probes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"liveness",
		"readiness",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varProbes := _Probes{}

	err = json.Unmarshal(data, &varProbes)

	if err != nil {
		return err
	}

	*o = Probes(varProbes)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "liveness")
		delete(additionalProperties, "readiness")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProbes struct {
	value *Probes
	isSet bool
}

func (v NullableProbes) Get() *Probes {
	return v.value
}

func (v *NullableProbes) Set(val *Probes) {
	v.value = val
	v.isSet = true
}

func (v NullableProbes) IsSet() bool {
	return v.isSet
}

func (v *NullableProbes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProbes(val *Probes) *NullableProbes {
	return &NullableProbes{value: val, isSet: true}
}

func (v NullableProbes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProbes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


