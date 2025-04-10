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

// checks if the RestoreInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestoreInstance{}

// RestoreInstance struct for RestoreInstance
type RestoreInstance struct {
	AppServices []AppType `json:"app_services,omitempty"`
	Autoscaling NullablePatchAutoscaling `json:"autoscaling,omitempty"`
	ConnectionPooler NullableConnectionPooler `json:"connection_pooler,omitempty"`
	Cpu NullableCpu `json:"cpu,omitempty"`
	Environment NullableEnvironment `json:"environment,omitempty"`
	ExtraDomainsRw []string `json:"extra_domains_rw,omitempty"`
	InstanceName string `json:"instance_name"`
	Memory NullableMemory `json:"memory,omitempty"`
	Restore Restore `json:"restore"`
	Storage NullableStorage `json:"storage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RestoreInstance RestoreInstance

// NewRestoreInstance instantiates a new RestoreInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoreInstance(instanceName string, restore Restore) *RestoreInstance {
	this := RestoreInstance{}
	this.InstanceName = instanceName
	this.Restore = restore
	return &this
}

// NewRestoreInstanceWithDefaults instantiates a new RestoreInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoreInstanceWithDefaults() *RestoreInstance {
	this := RestoreInstance{}
	return &this
}

// GetAppServices returns the AppServices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoreInstance) GetAppServices() []AppType {
	if o == nil {
		var ret []AppType
		return ret
	}
	return o.AppServices
}

// GetAppServicesOk returns a tuple with the AppServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoreInstance) GetAppServicesOk() ([]AppType, bool) {
	if o == nil || IsNil(o.AppServices) {
		return nil, false
	}
	return o.AppServices, true
}

// HasAppServices returns a boolean if a field has been set.
func (o *RestoreInstance) HasAppServices() bool {
	if o != nil && !IsNil(o.AppServices) {
		return true
	}

	return false
}

// SetAppServices gets a reference to the given []AppType and assigns it to the AppServices field.
func (o *RestoreInstance) SetAppServices(v []AppType) {
	o.AppServices = v
}

// GetAutoscaling returns the Autoscaling field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoreInstance) GetAutoscaling() PatchAutoscaling {
	if o == nil || IsNil(o.Autoscaling.Get()) {
		var ret PatchAutoscaling
		return ret
	}
	return *o.Autoscaling.Get()
}

// GetAutoscalingOk returns a tuple with the Autoscaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoreInstance) GetAutoscalingOk() (*PatchAutoscaling, bool) {
	if o == nil {
		return nil, false
	}
	return o.Autoscaling.Get(), o.Autoscaling.IsSet()
}

// HasAutoscaling returns a boolean if a field has been set.
func (o *RestoreInstance) HasAutoscaling() bool {
	if o != nil && o.Autoscaling.IsSet() {
		return true
	}

	return false
}

// SetAutoscaling gets a reference to the given NullablePatchAutoscaling and assigns it to the Autoscaling field.
func (o *RestoreInstance) SetAutoscaling(v PatchAutoscaling) {
	o.Autoscaling.Set(&v)
}
// SetAutoscalingNil sets the value for Autoscaling to be an explicit nil
func (o *RestoreInstance) SetAutoscalingNil() {
	o.Autoscaling.Set(nil)
}

// UnsetAutoscaling ensures that no value is present for Autoscaling, not even an explicit nil
func (o *RestoreInstance) UnsetAutoscaling() {
	o.Autoscaling.Unset()
}

// GetConnectionPooler returns the ConnectionPooler field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoreInstance) GetConnectionPooler() ConnectionPooler {
	if o == nil || IsNil(o.ConnectionPooler.Get()) {
		var ret ConnectionPooler
		return ret
	}
	return *o.ConnectionPooler.Get()
}

// GetConnectionPoolerOk returns a tuple with the ConnectionPooler field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoreInstance) GetConnectionPoolerOk() (*ConnectionPooler, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectionPooler.Get(), o.ConnectionPooler.IsSet()
}

// HasConnectionPooler returns a boolean if a field has been set.
func (o *RestoreInstance) HasConnectionPooler() bool {
	if o != nil && o.ConnectionPooler.IsSet() {
		return true
	}

	return false
}

// SetConnectionPooler gets a reference to the given NullableConnectionPooler and assigns it to the ConnectionPooler field.
func (o *RestoreInstance) SetConnectionPooler(v ConnectionPooler) {
	o.ConnectionPooler.Set(&v)
}
// SetConnectionPoolerNil sets the value for ConnectionPooler to be an explicit nil
func (o *RestoreInstance) SetConnectionPoolerNil() {
	o.ConnectionPooler.Set(nil)
}

// UnsetConnectionPooler ensures that no value is present for ConnectionPooler, not even an explicit nil
func (o *RestoreInstance) UnsetConnectionPooler() {
	o.ConnectionPooler.Unset()
}

// GetCpu returns the Cpu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoreInstance) GetCpu() Cpu {
	if o == nil || IsNil(o.Cpu.Get()) {
		var ret Cpu
		return ret
	}
	return *o.Cpu.Get()
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoreInstance) GetCpuOk() (*Cpu, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cpu.Get(), o.Cpu.IsSet()
}

// HasCpu returns a boolean if a field has been set.
func (o *RestoreInstance) HasCpu() bool {
	if o != nil && o.Cpu.IsSet() {
		return true
	}

	return false
}

// SetCpu gets a reference to the given NullableCpu and assigns it to the Cpu field.
func (o *RestoreInstance) SetCpu(v Cpu) {
	o.Cpu.Set(&v)
}
// SetCpuNil sets the value for Cpu to be an explicit nil
func (o *RestoreInstance) SetCpuNil() {
	o.Cpu.Set(nil)
}

// UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
func (o *RestoreInstance) UnsetCpu() {
	o.Cpu.Unset()
}

// GetEnvironment returns the Environment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoreInstance) GetEnvironment() Environment {
	if o == nil || IsNil(o.Environment.Get()) {
		var ret Environment
		return ret
	}
	return *o.Environment.Get()
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoreInstance) GetEnvironmentOk() (*Environment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Environment.Get(), o.Environment.IsSet()
}

// HasEnvironment returns a boolean if a field has been set.
func (o *RestoreInstance) HasEnvironment() bool {
	if o != nil && o.Environment.IsSet() {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given NullableEnvironment and assigns it to the Environment field.
func (o *RestoreInstance) SetEnvironment(v Environment) {
	o.Environment.Set(&v)
}
// SetEnvironmentNil sets the value for Environment to be an explicit nil
func (o *RestoreInstance) SetEnvironmentNil() {
	o.Environment.Set(nil)
}

// UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
func (o *RestoreInstance) UnsetEnvironment() {
	o.Environment.Unset()
}

// GetExtraDomainsRw returns the ExtraDomainsRw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoreInstance) GetExtraDomainsRw() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExtraDomainsRw
}

// GetExtraDomainsRwOk returns a tuple with the ExtraDomainsRw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoreInstance) GetExtraDomainsRwOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtraDomainsRw) {
		return nil, false
	}
	return o.ExtraDomainsRw, true
}

// HasExtraDomainsRw returns a boolean if a field has been set.
func (o *RestoreInstance) HasExtraDomainsRw() bool {
	if o != nil && !IsNil(o.ExtraDomainsRw) {
		return true
	}

	return false
}

// SetExtraDomainsRw gets a reference to the given []string and assigns it to the ExtraDomainsRw field.
func (o *RestoreInstance) SetExtraDomainsRw(v []string) {
	o.ExtraDomainsRw = v
}

// GetInstanceName returns the InstanceName field value
func (o *RestoreInstance) GetInstanceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value
// and a boolean to check if the value has been set.
func (o *RestoreInstance) GetInstanceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceName, true
}

// SetInstanceName sets field value
func (o *RestoreInstance) SetInstanceName(v string) {
	o.InstanceName = v
}

// GetMemory returns the Memory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoreInstance) GetMemory() Memory {
	if o == nil || IsNil(o.Memory.Get()) {
		var ret Memory
		return ret
	}
	return *o.Memory.Get()
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoreInstance) GetMemoryOk() (*Memory, bool) {
	if o == nil {
		return nil, false
	}
	return o.Memory.Get(), o.Memory.IsSet()
}

// HasMemory returns a boolean if a field has been set.
func (o *RestoreInstance) HasMemory() bool {
	if o != nil && o.Memory.IsSet() {
		return true
	}

	return false
}

// SetMemory gets a reference to the given NullableMemory and assigns it to the Memory field.
func (o *RestoreInstance) SetMemory(v Memory) {
	o.Memory.Set(&v)
}
// SetMemoryNil sets the value for Memory to be an explicit nil
func (o *RestoreInstance) SetMemoryNil() {
	o.Memory.Set(nil)
}

// UnsetMemory ensures that no value is present for Memory, not even an explicit nil
func (o *RestoreInstance) UnsetMemory() {
	o.Memory.Unset()
}

// GetRestore returns the Restore field value
func (o *RestoreInstance) GetRestore() Restore {
	if o == nil {
		var ret Restore
		return ret
	}

	return o.Restore
}

// GetRestoreOk returns a tuple with the Restore field value
// and a boolean to check if the value has been set.
func (o *RestoreInstance) GetRestoreOk() (*Restore, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Restore, true
}

// SetRestore sets field value
func (o *RestoreInstance) SetRestore(v Restore) {
	o.Restore = v
}

// GetStorage returns the Storage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoreInstance) GetStorage() Storage {
	if o == nil || IsNil(o.Storage.Get()) {
		var ret Storage
		return ret
	}
	return *o.Storage.Get()
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoreInstance) GetStorageOk() (*Storage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Storage.Get(), o.Storage.IsSet()
}

// HasStorage returns a boolean if a field has been set.
func (o *RestoreInstance) HasStorage() bool {
	if o != nil && o.Storage.IsSet() {
		return true
	}

	return false
}

// SetStorage gets a reference to the given NullableStorage and assigns it to the Storage field.
func (o *RestoreInstance) SetStorage(v Storage) {
	o.Storage.Set(&v)
}
// SetStorageNil sets the value for Storage to be an explicit nil
func (o *RestoreInstance) SetStorageNil() {
	o.Storage.Set(nil)
}

// UnsetStorage ensures that no value is present for Storage, not even an explicit nil
func (o *RestoreInstance) UnsetStorage() {
	o.Storage.Unset()
}

func (o RestoreInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestoreInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AppServices != nil {
		toSerialize["app_services"] = o.AppServices
	}
	if o.Autoscaling.IsSet() {
		toSerialize["autoscaling"] = o.Autoscaling.Get()
	}
	if o.ConnectionPooler.IsSet() {
		toSerialize["connection_pooler"] = o.ConnectionPooler.Get()
	}
	if o.Cpu.IsSet() {
		toSerialize["cpu"] = o.Cpu.Get()
	}
	if o.Environment.IsSet() {
		toSerialize["environment"] = o.Environment.Get()
	}
	if o.ExtraDomainsRw != nil {
		toSerialize["extra_domains_rw"] = o.ExtraDomainsRw
	}
	toSerialize["instance_name"] = o.InstanceName
	if o.Memory.IsSet() {
		toSerialize["memory"] = o.Memory.Get()
	}
	toSerialize["restore"] = o.Restore
	if o.Storage.IsSet() {
		toSerialize["storage"] = o.Storage.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RestoreInstance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"instance_name",
		"restore",
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

	varRestoreInstance := _RestoreInstance{}

	err = json.Unmarshal(data, &varRestoreInstance)

	if err != nil {
		return err
	}

	*o = RestoreInstance(varRestoreInstance)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "app_services")
		delete(additionalProperties, "autoscaling")
		delete(additionalProperties, "connection_pooler")
		delete(additionalProperties, "cpu")
		delete(additionalProperties, "environment")
		delete(additionalProperties, "extra_domains_rw")
		delete(additionalProperties, "instance_name")
		delete(additionalProperties, "memory")
		delete(additionalProperties, "restore")
		delete(additionalProperties, "storage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRestoreInstance struct {
	value *RestoreInstance
	isSet bool
}

func (v NullableRestoreInstance) Get() *RestoreInstance {
	return v.value
}

func (v *NullableRestoreInstance) Set(val *RestoreInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoreInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoreInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoreInstance(val *RestoreInstance) *NullableRestoreInstance {
	return &NullableRestoreInstance{value: val, isSet: true}
}

func (v NullableRestoreInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoreInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


