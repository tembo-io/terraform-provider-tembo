/*
Tembo Cloud

Platform API for Tembo Cloud             </br>             </br>             To find a Tembo Data API, please find it here:             </br>             </br>             [AWS US East 1](https://api.data-1.use1.tembo.io/swagger-ui/)             

API version: v1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package temboclient

import (
	"encoding/json"
)

// checks if the Probe type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Probe{}

// Probe struct for Probe
type Probe struct {
	InitialDelaySeconds int32 `json:"initialDelaySeconds"`
	Path string `json:"path"`
	Port string `json:"port"`
}

// NewProbe instantiates a new Probe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProbe(initialDelaySeconds int32, path string, port string) *Probe {
	this := Probe{}
	this.InitialDelaySeconds = initialDelaySeconds
	this.Path = path
	this.Port = port
	return &this
}

// NewProbeWithDefaults instantiates a new Probe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProbeWithDefaults() *Probe {
	this := Probe{}
	return &this
}

// GetInitialDelaySeconds returns the InitialDelaySeconds field value
func (o *Probe) GetInitialDelaySeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InitialDelaySeconds
}

// GetInitialDelaySecondsOk returns a tuple with the InitialDelaySeconds field value
// and a boolean to check if the value has been set.
func (o *Probe) GetInitialDelaySecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InitialDelaySeconds, true
}

// SetInitialDelaySeconds sets field value
func (o *Probe) SetInitialDelaySeconds(v int32) {
	o.InitialDelaySeconds = v
}

// GetPath returns the Path field value
func (o *Probe) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *Probe) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *Probe) SetPath(v string) {
	o.Path = v
}

// GetPort returns the Port field value
func (o *Probe) GetPort() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *Probe) GetPortOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *Probe) SetPort(v string) {
	o.Port = v
}

func (o Probe) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Probe) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["initialDelaySeconds"] = o.InitialDelaySeconds
	toSerialize["path"] = o.Path
	toSerialize["port"] = o.Port
	return toSerialize, nil
}

type NullableProbe struct {
	value *Probe
	isSet bool
}

func (v NullableProbe) Get() *Probe {
	return v.value
}

func (v *NullableProbe) Set(val *Probe) {
	v.value = val
	v.isSet = true
}

func (v NullableProbe) IsSet() bool {
	return v.isSet
}

func (v *NullableProbe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProbe(val *Probe) *NullableProbe {
	return &NullableProbe{value: val, isSet: true}
}

func (v NullableProbe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProbe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

