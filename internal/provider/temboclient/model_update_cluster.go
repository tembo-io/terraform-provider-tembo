/*
cp-webserver

CoreDB backend API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package temboclient

import (
	"encoding/json"
)

// checks if the UpdateCluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCluster{}

// UpdateCluster struct for UpdateCluster
type UpdateCluster struct {
	Cpu NullableCpu `json:"cpu,omitempty"`
	EntityProperties interface{} `json:"entity_properties,omitempty"`
	Environment NullableEnvironment `json:"environment,omitempty"`
	Extensions []Extension `json:"extensions,omitempty"`
	Memory NullableMemory `json:"memory,omitempty"`
	Storage NullableStorage `json:"storage,omitempty"`
	TrunkInstalls []TrunkInstall `json:"trunk_installs,omitempty"`
}

// NewUpdateCluster instantiates a new UpdateCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCluster() *UpdateCluster {
	this := UpdateCluster{}
	return &this
}

// NewUpdateClusterWithDefaults instantiates a new UpdateCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateClusterWithDefaults() *UpdateCluster {
	this := UpdateCluster{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateCluster) GetCpu() Cpu {
	if o == nil || IsNil(o.Cpu.Get()) {
		var ret Cpu
		return ret
	}
	return *o.Cpu.Get()
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCluster) GetCpuOk() (*Cpu, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cpu.Get(), o.Cpu.IsSet()
}

// HasCpu returns a boolean if a field has been set.
func (o *UpdateCluster) HasCpu() bool {
	if o != nil && o.Cpu.IsSet() {
		return true
	}

	return false
}

// SetCpu gets a reference to the given NullableCpu and assigns it to the Cpu field.
func (o *UpdateCluster) SetCpu(v Cpu) {
	o.Cpu.Set(&v)
}
// SetCpuNil sets the value for Cpu to be an explicit nil
func (o *UpdateCluster) SetCpuNil() {
	o.Cpu.Set(nil)
}

// UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
func (o *UpdateCluster) UnsetCpu() {
	o.Cpu.Unset()
}

// GetEntityProperties returns the EntityProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateCluster) GetEntityProperties() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EntityProperties
}

// GetEntityPropertiesOk returns a tuple with the EntityProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCluster) GetEntityPropertiesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EntityProperties) {
		return nil, false
	}
	return &o.EntityProperties, true
}

// HasEntityProperties returns a boolean if a field has been set.
func (o *UpdateCluster) HasEntityProperties() bool {
	if o != nil && IsNil(o.EntityProperties) {
		return true
	}

	return false
}

// SetEntityProperties gets a reference to the given interface{} and assigns it to the EntityProperties field.
func (o *UpdateCluster) SetEntityProperties(v interface{}) {
	o.EntityProperties = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateCluster) GetEnvironment() Environment {
	if o == nil || IsNil(o.Environment.Get()) {
		var ret Environment
		return ret
	}
	return *o.Environment.Get()
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCluster) GetEnvironmentOk() (*Environment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Environment.Get(), o.Environment.IsSet()
}

// HasEnvironment returns a boolean if a field has been set.
func (o *UpdateCluster) HasEnvironment() bool {
	if o != nil && o.Environment.IsSet() {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given NullableEnvironment and assigns it to the Environment field.
func (o *UpdateCluster) SetEnvironment(v Environment) {
	o.Environment.Set(&v)
}
// SetEnvironmentNil sets the value for Environment to be an explicit nil
func (o *UpdateCluster) SetEnvironmentNil() {
	o.Environment.Set(nil)
}

// UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
func (o *UpdateCluster) UnsetEnvironment() {
	o.Environment.Unset()
}

// GetExtensions returns the Extensions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateCluster) GetExtensions() []Extension {
	if o == nil {
		var ret []Extension
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCluster) GetExtensionsOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extensions) {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *UpdateCluster) HasExtensions() bool {
	if o != nil && IsNil(o.Extensions) {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given []Extension and assigns it to the Extensions field.
func (o *UpdateCluster) SetExtensions(v []Extension) {
	o.Extensions = v
}

// GetMemory returns the Memory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateCluster) GetMemory() Memory {
	if o == nil || IsNil(o.Memory.Get()) {
		var ret Memory
		return ret
	}
	return *o.Memory.Get()
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCluster) GetMemoryOk() (*Memory, bool) {
	if o == nil {
		return nil, false
	}
	return o.Memory.Get(), o.Memory.IsSet()
}

// HasMemory returns a boolean if a field has been set.
func (o *UpdateCluster) HasMemory() bool {
	if o != nil && o.Memory.IsSet() {
		return true
	}

	return false
}

// SetMemory gets a reference to the given NullableMemory and assigns it to the Memory field.
func (o *UpdateCluster) SetMemory(v Memory) {
	o.Memory.Set(&v)
}
// SetMemoryNil sets the value for Memory to be an explicit nil
func (o *UpdateCluster) SetMemoryNil() {
	o.Memory.Set(nil)
}

// UnsetMemory ensures that no value is present for Memory, not even an explicit nil
func (o *UpdateCluster) UnsetMemory() {
	o.Memory.Unset()
}

// GetStorage returns the Storage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateCluster) GetStorage() Storage {
	if o == nil || IsNil(o.Storage.Get()) {
		var ret Storage
		return ret
	}
	return *o.Storage.Get()
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCluster) GetStorageOk() (*Storage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Storage.Get(), o.Storage.IsSet()
}

// HasStorage returns a boolean if a field has been set.
func (o *UpdateCluster) HasStorage() bool {
	if o != nil && o.Storage.IsSet() {
		return true
	}

	return false
}

// SetStorage gets a reference to the given NullableStorage and assigns it to the Storage field.
func (o *UpdateCluster) SetStorage(v Storage) {
	o.Storage.Set(&v)
}
// SetStorageNil sets the value for Storage to be an explicit nil
func (o *UpdateCluster) SetStorageNil() {
	o.Storage.Set(nil)
}

// UnsetStorage ensures that no value is present for Storage, not even an explicit nil
func (o *UpdateCluster) UnsetStorage() {
	o.Storage.Unset()
}

// GetTrunkInstalls returns the TrunkInstalls field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateCluster) GetTrunkInstalls() []TrunkInstall {
	if o == nil {
		var ret []TrunkInstall
		return ret
	}
	return o.TrunkInstalls
}

// GetTrunkInstallsOk returns a tuple with the TrunkInstalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCluster) GetTrunkInstallsOk() ([]TrunkInstall, bool) {
	if o == nil || IsNil(o.TrunkInstalls) {
		return nil, false
	}
	return o.TrunkInstalls, true
}

// HasTrunkInstalls returns a boolean if a field has been set.
func (o *UpdateCluster) HasTrunkInstalls() bool {
	if o != nil && IsNil(o.TrunkInstalls) {
		return true
	}

	return false
}

// SetTrunkInstalls gets a reference to the given []TrunkInstall and assigns it to the TrunkInstalls field.
func (o *UpdateCluster) SetTrunkInstalls(v []TrunkInstall) {
	o.TrunkInstalls = v
}

func (o UpdateCluster) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Cpu.IsSet() {
		toSerialize["cpu"] = o.Cpu.Get()
	}
	if o.EntityProperties != nil {
		toSerialize["entity_properties"] = o.EntityProperties
	}
	if o.Environment.IsSet() {
		toSerialize["environment"] = o.Environment.Get()
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Memory.IsSet() {
		toSerialize["memory"] = o.Memory.Get()
	}
	if o.Storage.IsSet() {
		toSerialize["storage"] = o.Storage.Get()
	}
	if o.TrunkInstalls != nil {
		toSerialize["trunk_installs"] = o.TrunkInstalls
	}
	return toSerialize, nil
}

type NullableUpdateCluster struct {
	value *UpdateCluster
	isSet bool
}

func (v NullableUpdateCluster) Get() *UpdateCluster {
	return v.value
}

func (v *NullableUpdateCluster) Set(val *UpdateCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCluster(val *UpdateCluster) *NullableUpdateCluster {
	return &NullableUpdateCluster{value: val, isSet: true}
}

func (v NullableUpdateCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

