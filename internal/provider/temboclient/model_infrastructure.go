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

// checks if the Infrastructure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Infrastructure{}

// Infrastructure struct for Infrastructure
type Infrastructure struct {
	Cpu *string `json:"cpu,omitempty"`
	Memory *string `json:"memory,omitempty"`
	Storage *string `json:"storage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Infrastructure Infrastructure

// NewInfrastructure instantiates a new Infrastructure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfrastructure() *Infrastructure {
	this := Infrastructure{}
	return &this
}

// NewInfrastructureWithDefaults instantiates a new Infrastructure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfrastructureWithDefaults() *Infrastructure {
	this := Infrastructure{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *Infrastructure) GetCpu() string {
	if o == nil || IsNil(o.Cpu) {
		var ret string
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetCpuOk() (*string, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *Infrastructure) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given string and assigns it to the Cpu field.
func (o *Infrastructure) SetCpu(v string) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *Infrastructure) GetMemory() string {
	if o == nil || IsNil(o.Memory) {
		var ret string
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetMemoryOk() (*string, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *Infrastructure) HasMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given string and assigns it to the Memory field.
func (o *Infrastructure) SetMemory(v string) {
	o.Memory = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *Infrastructure) GetStorage() string {
	if o == nil || IsNil(o.Storage) {
		var ret string
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetStorageOk() (*string, bool) {
	if o == nil || IsNil(o.Storage) {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *Infrastructure) HasStorage() bool {
	if o != nil && !IsNil(o.Storage) {
		return true
	}

	return false
}

// SetStorage gets a reference to the given string and assigns it to the Storage field.
func (o *Infrastructure) SetStorage(v string) {
	o.Storage = &v
}

func (o Infrastructure) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Infrastructure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	if !IsNil(o.Storage) {
		toSerialize["storage"] = o.Storage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Infrastructure) UnmarshalJSON(data []byte) (err error) {
	varInfrastructure := _Infrastructure{}

	err = json.Unmarshal(data, &varInfrastructure)

	if err != nil {
		return err
	}

	*o = Infrastructure(varInfrastructure)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cpu")
		delete(additionalProperties, "memory")
		delete(additionalProperties, "storage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInfrastructure struct {
	value *Infrastructure
	isSet bool
}

func (v NullableInfrastructure) Get() *Infrastructure {
	return v.value
}

func (v *NullableInfrastructure) Set(val *Infrastructure) {
	v.value = val
	v.isSet = true
}

func (v NullableInfrastructure) IsSet() bool {
	return v.isSet
}

func (v *NullableInfrastructure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfrastructure(val *Infrastructure) *NullableInfrastructure {
	return &NullableInfrastructure{value: val, isSet: true}
}

func (v NullableInfrastructure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfrastructure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


