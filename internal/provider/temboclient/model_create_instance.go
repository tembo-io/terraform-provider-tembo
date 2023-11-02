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

// checks if the CreateInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateInstance{}

// CreateInstance struct for CreateInstance
type CreateInstance struct {
	AppServices []AppType `json:"app_services,omitempty"`
	ConnectionPooler NullableConnectionPooler `json:"connection_pooler,omitempty"`
	Cpu Cpu `json:"cpu"`
	Environment Environment `json:"environment"`
	Extensions []Extension `json:"extensions,omitempty"`
	ExtraDomainsRw []string `json:"extra_domains_rw,omitempty"`
	InstanceName string `json:"instance_name"`
	IpAllowList []string `json:"ip_allow_list,omitempty"`
	Memory Memory `json:"memory"`
	PostgresConfigs []PgConfig `json:"postgres_configs,omitempty"`
	Replicas *int32 `json:"replicas,omitempty"`
	StackType StackType `json:"stack_type"`
	Storage Storage `json:"storage"`
	TrunkInstalls []TrunkInstall `json:"trunk_installs,omitempty"`
}

// NewCreateInstance instantiates a new CreateInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInstance(cpu Cpu, environment Environment, instanceName string, memory Memory, stackType StackType, storage Storage) *CreateInstance {
	this := CreateInstance{}
	this.Cpu = cpu
	this.Environment = environment
	this.InstanceName = instanceName
	this.Memory = memory
	this.StackType = stackType
	this.Storage = storage
	return &this
}

// NewCreateInstanceWithDefaults instantiates a new CreateInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInstanceWithDefaults() *CreateInstance {
	this := CreateInstance{}
	return &this
}

// GetAppServices returns the AppServices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateInstance) GetAppServices() []AppType {
	if o == nil {
		var ret []AppType
		return ret
	}
	return o.AppServices
}

// GetAppServicesOk returns a tuple with the AppServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateInstance) GetAppServicesOk() ([]AppType, bool) {
	if o == nil || IsNil(o.AppServices) {
		return nil, false
	}
	return o.AppServices, true
}

// HasAppServices returns a boolean if a field has been set.
func (o *CreateInstance) HasAppServices() bool {
	if o != nil && IsNil(o.AppServices) {
		return true
	}

	return false
}

// SetAppServices gets a reference to the given []AppType and assigns it to the AppServices field.
func (o *CreateInstance) SetAppServices(v []AppType) {
	o.AppServices = v
}

// GetConnectionPooler returns the ConnectionPooler field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateInstance) GetConnectionPooler() ConnectionPooler {
	if o == nil || IsNil(o.ConnectionPooler.Get()) {
		var ret ConnectionPooler
		return ret
	}
	return *o.ConnectionPooler.Get()
}

// GetConnectionPoolerOk returns a tuple with the ConnectionPooler field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateInstance) GetConnectionPoolerOk() (*ConnectionPooler, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectionPooler.Get(), o.ConnectionPooler.IsSet()
}

// HasConnectionPooler returns a boolean if a field has been set.
func (o *CreateInstance) HasConnectionPooler() bool {
	if o != nil && o.ConnectionPooler.IsSet() {
		return true
	}

	return false
}

// SetConnectionPooler gets a reference to the given NullableConnectionPooler and assigns it to the ConnectionPooler field.
func (o *CreateInstance) SetConnectionPooler(v ConnectionPooler) {
	o.ConnectionPooler.Set(&v)
}
// SetConnectionPoolerNil sets the value for ConnectionPooler to be an explicit nil
func (o *CreateInstance) SetConnectionPoolerNil() {
	o.ConnectionPooler.Set(nil)
}

// UnsetConnectionPooler ensures that no value is present for ConnectionPooler, not even an explicit nil
func (o *CreateInstance) UnsetConnectionPooler() {
	o.ConnectionPooler.Unset()
}

// GetCpu returns the Cpu field value
func (o *CreateInstance) GetCpu() Cpu {
	if o == nil {
		var ret Cpu
		return ret
	}

	return o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value
// and a boolean to check if the value has been set.
func (o *CreateInstance) GetCpuOk() (*Cpu, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cpu, true
}

// SetCpu sets field value
func (o *CreateInstance) SetCpu(v Cpu) {
	o.Cpu = v
}

// GetEnvironment returns the Environment field value
func (o *CreateInstance) GetEnvironment() Environment {
	if o == nil {
		var ret Environment
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *CreateInstance) GetEnvironmentOk() (*Environment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *CreateInstance) SetEnvironment(v Environment) {
	o.Environment = v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateInstance) GetExtensions() []Extension {
	if o == nil {
		var ret []Extension
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateInstance) GetExtensionsOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extensions) {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *CreateInstance) HasExtensions() bool {
	if o != nil && IsNil(o.Extensions) {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given []Extension and assigns it to the Extensions field.
func (o *CreateInstance) SetExtensions(v []Extension) {
	o.Extensions = v
}

// GetExtraDomainsRw returns the ExtraDomainsRw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateInstance) GetExtraDomainsRw() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExtraDomainsRw
}

// GetExtraDomainsRwOk returns a tuple with the ExtraDomainsRw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateInstance) GetExtraDomainsRwOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtraDomainsRw) {
		return nil, false
	}
	return o.ExtraDomainsRw, true
}

// HasExtraDomainsRw returns a boolean if a field has been set.
func (o *CreateInstance) HasExtraDomainsRw() bool {
	if o != nil && IsNil(o.ExtraDomainsRw) {
		return true
	}

	return false
}

// SetExtraDomainsRw gets a reference to the given []string and assigns it to the ExtraDomainsRw field.
func (o *CreateInstance) SetExtraDomainsRw(v []string) {
	o.ExtraDomainsRw = v
}

// GetInstanceName returns the InstanceName field value
func (o *CreateInstance) GetInstanceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value
// and a boolean to check if the value has been set.
func (o *CreateInstance) GetInstanceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceName, true
}

// SetInstanceName sets field value
func (o *CreateInstance) SetInstanceName(v string) {
	o.InstanceName = v
}

// GetIpAllowList returns the IpAllowList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateInstance) GetIpAllowList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IpAllowList
}

// GetIpAllowListOk returns a tuple with the IpAllowList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateInstance) GetIpAllowListOk() ([]string, bool) {
	if o == nil || IsNil(o.IpAllowList) {
		return nil, false
	}
	return o.IpAllowList, true
}

// HasIpAllowList returns a boolean if a field has been set.
func (o *CreateInstance) HasIpAllowList() bool {
	if o != nil && IsNil(o.IpAllowList) {
		return true
	}

	return false
}

// SetIpAllowList gets a reference to the given []string and assigns it to the IpAllowList field.
func (o *CreateInstance) SetIpAllowList(v []string) {
	o.IpAllowList = v
}

// GetMemory returns the Memory field value
func (o *CreateInstance) GetMemory() Memory {
	if o == nil {
		var ret Memory
		return ret
	}

	return o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value
// and a boolean to check if the value has been set.
func (o *CreateInstance) GetMemoryOk() (*Memory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memory, true
}

// SetMemory sets field value
func (o *CreateInstance) SetMemory(v Memory) {
	o.Memory = v
}

// GetPostgresConfigs returns the PostgresConfigs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateInstance) GetPostgresConfigs() []PgConfig {
	if o == nil {
		var ret []PgConfig
		return ret
	}
	return o.PostgresConfigs
}

// GetPostgresConfigsOk returns a tuple with the PostgresConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateInstance) GetPostgresConfigsOk() ([]PgConfig, bool) {
	if o == nil || IsNil(o.PostgresConfigs) {
		return nil, false
	}
	return o.PostgresConfigs, true
}

// HasPostgresConfigs returns a boolean if a field has been set.
func (o *CreateInstance) HasPostgresConfigs() bool {
	if o != nil && IsNil(o.PostgresConfigs) {
		return true
	}

	return false
}

// SetPostgresConfigs gets a reference to the given []PgConfig and assigns it to the PostgresConfigs field.
func (o *CreateInstance) SetPostgresConfigs(v []PgConfig) {
	o.PostgresConfigs = v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *CreateInstance) GetReplicas() int32 {
	if o == nil || IsNil(o.Replicas) {
		var ret int32
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstance) GetReplicasOk() (*int32, bool) {
	if o == nil || IsNil(o.Replicas) {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *CreateInstance) HasReplicas() bool {
	if o != nil && !IsNil(o.Replicas) {
		return true
	}

	return false
}

// SetReplicas gets a reference to the given int32 and assigns it to the Replicas field.
func (o *CreateInstance) SetReplicas(v int32) {
	o.Replicas = &v
}

// GetStackType returns the StackType field value
func (o *CreateInstance) GetStackType() StackType {
	if o == nil {
		var ret StackType
		return ret
	}

	return o.StackType
}

// GetStackTypeOk returns a tuple with the StackType field value
// and a boolean to check if the value has been set.
func (o *CreateInstance) GetStackTypeOk() (*StackType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StackType, true
}

// SetStackType sets field value
func (o *CreateInstance) SetStackType(v StackType) {
	o.StackType = v
}

// GetStorage returns the Storage field value
func (o *CreateInstance) GetStorage() Storage {
	if o == nil {
		var ret Storage
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *CreateInstance) GetStorageOk() (*Storage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *CreateInstance) SetStorage(v Storage) {
	o.Storage = v
}

// GetTrunkInstalls returns the TrunkInstalls field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateInstance) GetTrunkInstalls() []TrunkInstall {
	if o == nil {
		var ret []TrunkInstall
		return ret
	}
	return o.TrunkInstalls
}

// GetTrunkInstallsOk returns a tuple with the TrunkInstalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateInstance) GetTrunkInstallsOk() ([]TrunkInstall, bool) {
	if o == nil || IsNil(o.TrunkInstalls) {
		return nil, false
	}
	return o.TrunkInstalls, true
}

// HasTrunkInstalls returns a boolean if a field has been set.
func (o *CreateInstance) HasTrunkInstalls() bool {
	if o != nil && IsNil(o.TrunkInstalls) {
		return true
	}

	return false
}

// SetTrunkInstalls gets a reference to the given []TrunkInstall and assigns it to the TrunkInstalls field.
func (o *CreateInstance) SetTrunkInstalls(v []TrunkInstall) {
	o.TrunkInstalls = v
}

func (o CreateInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AppServices != nil {
		toSerialize["app_services"] = o.AppServices
	}
	if o.ConnectionPooler.IsSet() {
		toSerialize["connection_pooler"] = o.ConnectionPooler.Get()
	}
	toSerialize["cpu"] = o.Cpu
	toSerialize["environment"] = o.Environment
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.ExtraDomainsRw != nil {
		toSerialize["extra_domains_rw"] = o.ExtraDomainsRw
	}
	toSerialize["instance_name"] = o.InstanceName
	if o.IpAllowList != nil {
		toSerialize["ip_allow_list"] = o.IpAllowList
	}
	toSerialize["memory"] = o.Memory
	if o.PostgresConfigs != nil {
		toSerialize["postgres_configs"] = o.PostgresConfigs
	}
	if !IsNil(o.Replicas) {
		toSerialize["replicas"] = o.Replicas
	}
	toSerialize["stack_type"] = o.StackType
	toSerialize["storage"] = o.Storage
	if o.TrunkInstalls != nil {
		toSerialize["trunk_installs"] = o.TrunkInstalls
	}
	return toSerialize, nil
}

type NullableCreateInstance struct {
	value *CreateInstance
	isSet bool
}

func (v NullableCreateInstance) Get() *CreateInstance {
	return v.value
}

func (v *NullableCreateInstance) Set(val *CreateInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInstance(val *CreateInstance) *NullableCreateInstance {
	return &NullableCreateInstance{value: val, isSet: true}
}

func (v NullableCreateInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


