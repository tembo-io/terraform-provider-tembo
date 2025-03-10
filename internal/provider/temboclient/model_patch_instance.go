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

// checks if the PatchInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchInstance{}

// PatchInstance struct for PatchInstance
type PatchInstance struct {
	AppServices []AppType `json:"app_services,omitempty"`
	Autoscaling NullablePatchAutoscaling `json:"autoscaling,omitempty"`
	ConnectionPooler NullableConnectionPooler `json:"connection_pooler,omitempty"`
	Cpu NullableCpu `json:"cpu,omitempty"`
	DedicatedNetworking NullableDedicatedNetworking `json:"dedicated_networking,omitempty"`
	Environment NullableEnvironment `json:"environment,omitempty"`
	Experimental NullableExperimental `json:"experimental,omitempty"`
	Extensions []Extension `json:"extensions,omitempty"`
	ExtraDomainsRw []string `json:"extra_domains_rw,omitempty"`
	InstanceName NullableString `json:"instance_name,omitempty"`
	IpAllowList []string `json:"ip_allow_list,omitempty"`
	Memory NullableMemory `json:"memory,omitempty"`
	PostgresConfigs []PgConfig `json:"postgres_configs,omitempty"`
	Replicas NullableInt32 `json:"replicas,omitempty"`
	Spot NullableBool `json:"spot,omitempty"`
	Storage NullableStorage `json:"storage,omitempty"`
	TrunkInstalls []TrunkInstall `json:"trunk_installs,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchInstance PatchInstance

// NewPatchInstance instantiates a new PatchInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchInstance() *PatchInstance {
	this := PatchInstance{}
	return &this
}

// NewPatchInstanceWithDefaults instantiates a new PatchInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchInstanceWithDefaults() *PatchInstance {
	this := PatchInstance{}
	return &this
}

// GetAppServices returns the AppServices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchInstance) GetAppServices() []AppType {
	if o == nil {
		var ret []AppType
		return ret
	}
	return o.AppServices
}

// GetAppServicesOk returns a tuple with the AppServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchInstance) GetAppServicesOk() ([]AppType, bool) {
	if o == nil || IsNil(o.AppServices) {
		return nil, false
	}
	return o.AppServices, true
}

// HasAppServices returns a boolean if a field has been set.
func (o *PatchInstance) HasAppServices() bool {
	if o != nil && !IsNil(o.AppServices) {
		return true
	}

	return false
}

// SetAppServices gets a reference to the given []AppType and assigns it to the AppServices field.
func (o *PatchInstance) SetAppServices(v []AppType) {
	o.AppServices = v
}

// GetAutoscaling returns the Autoscaling field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchInstance) GetAutoscaling() PatchAutoscaling {
	if o == nil || IsNil(o.Autoscaling.Get()) {
		var ret PatchAutoscaling
		return ret
	}
	return *o.Autoscaling.Get()
}

// GetAutoscalingOk returns a tuple with the Autoscaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchInstance) GetAutoscalingOk() (*PatchAutoscaling, bool) {
	if o == nil {
		return nil, false
	}
	return o.Autoscaling.Get(), o.Autoscaling.IsSet()
}

// HasAutoscaling returns a boolean if a field has been set.
func (o *PatchInstance) HasAutoscaling() bool {
	if o != nil && o.Autoscaling.IsSet() {
		return true
	}

	return false
}

// SetAutoscaling gets a reference to the given NullablePatchAutoscaling and assigns it to the Autoscaling field.
func (o *PatchInstance) SetAutoscaling(v PatchAutoscaling) {
	o.Autoscaling.Set(&v)
}
// SetAutoscalingNil sets the value for Autoscaling to be an explicit nil
func (o *PatchInstance) SetAutoscalingNil() {
	o.Autoscaling.Set(nil)
}

// UnsetAutoscaling ensures that no value is present for Autoscaling, not even an explicit nil
func (o *PatchInstance) UnsetAutoscaling() {
	o.Autoscaling.Unset()
}

// GetConnectionPooler returns the ConnectionPooler field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchInstance) GetConnectionPooler() ConnectionPooler {
	if o == nil || IsNil(o.ConnectionPooler.Get()) {
		var ret ConnectionPooler
		return ret
	}
	return *o.ConnectionPooler.Get()
}

// GetConnectionPoolerOk returns a tuple with the ConnectionPooler field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchInstance) GetConnectionPoolerOk() (*ConnectionPooler, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectionPooler.Get(), o.ConnectionPooler.IsSet()
}

// HasConnectionPooler returns a boolean if a field has been set.
func (o *PatchInstance) HasConnectionPooler() bool {
	if o != nil && o.ConnectionPooler.IsSet() {
		return true
	}

	return false
}

// SetConnectionPooler gets a reference to the given NullableConnectionPooler and assigns it to the ConnectionPooler field.
func (o *PatchInstance) SetConnectionPooler(v ConnectionPooler) {
	o.ConnectionPooler.Set(&v)
}
// SetConnectionPoolerNil sets the value for ConnectionPooler to be an explicit nil
func (o *PatchInstance) SetConnectionPoolerNil() {
	o.ConnectionPooler.Set(nil)
}

// UnsetConnectionPooler ensures that no value is present for ConnectionPooler, not even an explicit nil
func (o *PatchInstance) UnsetConnectionPooler() {
	o.ConnectionPooler.Unset()
}

// GetCpu returns the Cpu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchInstance) GetCpu() Cpu {
	if o == nil || IsNil(o.Cpu.Get()) {
		var ret Cpu
		return ret
	}
	return *o.Cpu.Get()
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchInstance) GetCpuOk() (*Cpu, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cpu.Get(), o.Cpu.IsSet()
}

// HasCpu returns a boolean if a field has been set.
func (o *PatchInstance) HasCpu() bool {
	if o != nil && o.Cpu.IsSet() {
		return true
	}

	return false
}

// SetCpu gets a reference to the given NullableCpu and assigns it to the Cpu field.
func (o *PatchInstance) SetCpu(v Cpu) {
	o.Cpu.Set(&v)
}
// SetCpuNil sets the value for Cpu to be an explicit nil
func (o *PatchInstance) SetCpuNil() {
	o.Cpu.Set(nil)
}

// UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
func (o *PatchInstance) UnsetCpu() {
	o.Cpu.Unset()
}

// GetDedicatedNetworking returns the DedicatedNetworking field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchInstance) GetDedicatedNetworking() DedicatedNetworking {
	if o == nil || IsNil(o.DedicatedNetworking.Get()) {
		var ret DedicatedNetworking
		return ret
	}
	return *o.DedicatedNetworking.Get()
}

// GetDedicatedNetworkingOk returns a tuple with the DedicatedNetworking field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchInstance) GetDedicatedNetworkingOk() (*DedicatedNetworking, bool) {
	if o == nil {
		return nil, false
	}
	return o.DedicatedNetworking.Get(), o.DedicatedNetworking.IsSet()
}

// HasDedicatedNetworking returns a boolean if a field has been set.
func (o *PatchInstance) HasDedicatedNetworking() bool {
	if o != nil && o.DedicatedNetworking.IsSet() {
		return true
	}

	return false
}

// SetDedicatedNetworking gets a reference to the given NullableDedicatedNetworking and assigns it to the DedicatedNetworking field.
func (o *PatchInstance) SetDedicatedNetworking(v DedicatedNetworking) {
	o.DedicatedNetworking.Set(&v)
}
// SetDedicatedNetworkingNil sets the value for DedicatedNetworking to be an explicit nil
func (o *PatchInstance) SetDedicatedNetworkingNil() {
	o.DedicatedNetworking.Set(nil)
}

// UnsetDedicatedNetworking ensures that no value is present for DedicatedNetworking, not even an explicit nil
func (o *PatchInstance) UnsetDedicatedNetworking() {
	o.DedicatedNetworking.Unset()
}

// GetEnvironment returns the Environment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchInstance) GetEnvironment() Environment {
	if o == nil || IsNil(o.Environment.Get()) {
		var ret Environment
		return ret
	}
	return *o.Environment.Get()
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchInstance) GetEnvironmentOk() (*Environment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Environment.Get(), o.Environment.IsSet()
}

// HasEnvironment returns a boolean if a field has been set.
func (o *PatchInstance) HasEnvironment() bool {
	if o != nil && o.Environment.IsSet() {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given NullableEnvironment and assigns it to the Environment field.
func (o *PatchInstance) SetEnvironment(v Environment) {
	o.Environment.Set(&v)
}
// SetEnvironmentNil sets the value for Environment to be an explicit nil
func (o *PatchInstance) SetEnvironmentNil() {
	o.Environment.Set(nil)
}

// UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
func (o *PatchInstance) UnsetEnvironment() {
	o.Environment.Unset()
}

// GetExperimental returns the Experimental field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchInstance) GetExperimental() Experimental {
	if o == nil || IsNil(o.Experimental.Get()) {
		var ret Experimental
		return ret
	}
	return *o.Experimental.Get()
}

// GetExperimentalOk returns a tuple with the Experimental field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchInstance) GetExperimentalOk() (*Experimental, bool) {
	if o == nil {
		return nil, false
	}
	return o.Experimental.Get(), o.Experimental.IsSet()
}

// HasExperimental returns a boolean if a field has been set.
func (o *PatchInstance) HasExperimental() bool {
	if o != nil && o.Experimental.IsSet() {
		return true
	}

	return false
}

// SetExperimental gets a reference to the given NullableExperimental and assigns it to the Experimental field.
func (o *PatchInstance) SetExperimental(v Experimental) {
	o.Experimental.Set(&v)
}
// SetExperimentalNil sets the value for Experimental to be an explicit nil
func (o *PatchInstance) SetExperimentalNil() {
	o.Experimental.Set(nil)
}

// UnsetExperimental ensures that no value is present for Experimental, not even an explicit nil
func (o *PatchInstance) UnsetExperimental() {
	o.Experimental.Unset()
}

// GetExtensions returns the Extensions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchInstance) GetExtensions() []Extension {
	if o == nil {
		var ret []Extension
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchInstance) GetExtensionsOk() ([]Extension, bool) {
	if o == nil || IsNil(o.Extensions) {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *PatchInstance) HasExtensions() bool {
	if o != nil && !IsNil(o.Extensions) {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given []Extension and assigns it to the Extensions field.
func (o *PatchInstance) SetExtensions(v []Extension) {
	o.Extensions = v
}

// GetExtraDomainsRw returns the ExtraDomainsRw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchInstance) GetExtraDomainsRw() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExtraDomainsRw
}

// GetExtraDomainsRwOk returns a tuple with the ExtraDomainsRw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchInstance) GetExtraDomainsRwOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtraDomainsRw) {
		return nil, false
	}
	return o.ExtraDomainsRw, true
}

// HasExtraDomainsRw returns a boolean if a field has been set.
func (o *PatchInstance) HasExtraDomainsRw() bool {
	if o != nil && !IsNil(o.ExtraDomainsRw) {
		return true
	}

	return false
}

// SetExtraDomainsRw gets a reference to the given []string and assigns it to the ExtraDomainsRw field.
func (o *PatchInstance) SetExtraDomainsRw(v []string) {
	o.ExtraDomainsRw = v
}

// GetInstanceName returns the InstanceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchInstance) GetInstanceName() string {
	if o == nil || IsNil(o.InstanceName.Get()) {
		var ret string
		return ret
	}
	return *o.InstanceName.Get()
}

// GetInstanceNameOk returns a tuple with the InstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchInstance) GetInstanceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstanceName.Get(), o.InstanceName.IsSet()
}

// HasInstanceName returns a boolean if a field has been set.
func (o *PatchInstance) HasInstanceName() bool {
	if o != nil && o.InstanceName.IsSet() {
		return true
	}

	return false
}

// SetInstanceName gets a reference to the given NullableString and assigns it to the InstanceName field.
func (o *PatchInstance) SetInstanceName(v string) {
	o.InstanceName.Set(&v)
}
// SetInstanceNameNil sets the value for InstanceName to be an explicit nil
func (o *PatchInstance) SetInstanceNameNil() {
	o.InstanceName.Set(nil)
}

// UnsetInstanceName ensures that no value is present for InstanceName, not even an explicit nil
func (o *PatchInstance) UnsetInstanceName() {
	o.InstanceName.Unset()
}

// GetIpAllowList returns the IpAllowList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchInstance) GetIpAllowList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IpAllowList
}

// GetIpAllowListOk returns a tuple with the IpAllowList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchInstance) GetIpAllowListOk() ([]string, bool) {
	if o == nil || IsNil(o.IpAllowList) {
		return nil, false
	}
	return o.IpAllowList, true
}

// HasIpAllowList returns a boolean if a field has been set.
func (o *PatchInstance) HasIpAllowList() bool {
	if o != nil && !IsNil(o.IpAllowList) {
		return true
	}

	return false
}

// SetIpAllowList gets a reference to the given []string and assigns it to the IpAllowList field.
func (o *PatchInstance) SetIpAllowList(v []string) {
	o.IpAllowList = v
}

// GetMemory returns the Memory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchInstance) GetMemory() Memory {
	if o == nil || IsNil(o.Memory.Get()) {
		var ret Memory
		return ret
	}
	return *o.Memory.Get()
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchInstance) GetMemoryOk() (*Memory, bool) {
	if o == nil {
		return nil, false
	}
	return o.Memory.Get(), o.Memory.IsSet()
}

// HasMemory returns a boolean if a field has been set.
func (o *PatchInstance) HasMemory() bool {
	if o != nil && o.Memory.IsSet() {
		return true
	}

	return false
}

// SetMemory gets a reference to the given NullableMemory and assigns it to the Memory field.
func (o *PatchInstance) SetMemory(v Memory) {
	o.Memory.Set(&v)
}
// SetMemoryNil sets the value for Memory to be an explicit nil
func (o *PatchInstance) SetMemoryNil() {
	o.Memory.Set(nil)
}

// UnsetMemory ensures that no value is present for Memory, not even an explicit nil
func (o *PatchInstance) UnsetMemory() {
	o.Memory.Unset()
}

// GetPostgresConfigs returns the PostgresConfigs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchInstance) GetPostgresConfigs() []PgConfig {
	if o == nil {
		var ret []PgConfig
		return ret
	}
	return o.PostgresConfigs
}

// GetPostgresConfigsOk returns a tuple with the PostgresConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchInstance) GetPostgresConfigsOk() ([]PgConfig, bool) {
	if o == nil || IsNil(o.PostgresConfigs) {
		return nil, false
	}
	return o.PostgresConfigs, true
}

// HasPostgresConfigs returns a boolean if a field has been set.
func (o *PatchInstance) HasPostgresConfigs() bool {
	if o != nil && !IsNil(o.PostgresConfigs) {
		return true
	}

	return false
}

// SetPostgresConfigs gets a reference to the given []PgConfig and assigns it to the PostgresConfigs field.
func (o *PatchInstance) SetPostgresConfigs(v []PgConfig) {
	o.PostgresConfigs = v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchInstance) GetReplicas() int32 {
	if o == nil || IsNil(o.Replicas.Get()) {
		var ret int32
		return ret
	}
	return *o.Replicas.Get()
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchInstance) GetReplicasOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Replicas.Get(), o.Replicas.IsSet()
}

// HasReplicas returns a boolean if a field has been set.
func (o *PatchInstance) HasReplicas() bool {
	if o != nil && o.Replicas.IsSet() {
		return true
	}

	return false
}

// SetReplicas gets a reference to the given NullableInt32 and assigns it to the Replicas field.
func (o *PatchInstance) SetReplicas(v int32) {
	o.Replicas.Set(&v)
}
// SetReplicasNil sets the value for Replicas to be an explicit nil
func (o *PatchInstance) SetReplicasNil() {
	o.Replicas.Set(nil)
}

// UnsetReplicas ensures that no value is present for Replicas, not even an explicit nil
func (o *PatchInstance) UnsetReplicas() {
	o.Replicas.Unset()
}

// GetSpot returns the Spot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchInstance) GetSpot() bool {
	if o == nil || IsNil(o.Spot.Get()) {
		var ret bool
		return ret
	}
	return *o.Spot.Get()
}

// GetSpotOk returns a tuple with the Spot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchInstance) GetSpotOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Spot.Get(), o.Spot.IsSet()
}

// HasSpot returns a boolean if a field has been set.
func (o *PatchInstance) HasSpot() bool {
	if o != nil && o.Spot.IsSet() {
		return true
	}

	return false
}

// SetSpot gets a reference to the given NullableBool and assigns it to the Spot field.
func (o *PatchInstance) SetSpot(v bool) {
	o.Spot.Set(&v)
}
// SetSpotNil sets the value for Spot to be an explicit nil
func (o *PatchInstance) SetSpotNil() {
	o.Spot.Set(nil)
}

// UnsetSpot ensures that no value is present for Spot, not even an explicit nil
func (o *PatchInstance) UnsetSpot() {
	o.Spot.Unset()
}

// GetStorage returns the Storage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchInstance) GetStorage() Storage {
	if o == nil || IsNil(o.Storage.Get()) {
		var ret Storage
		return ret
	}
	return *o.Storage.Get()
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchInstance) GetStorageOk() (*Storage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Storage.Get(), o.Storage.IsSet()
}

// HasStorage returns a boolean if a field has been set.
func (o *PatchInstance) HasStorage() bool {
	if o != nil && o.Storage.IsSet() {
		return true
	}

	return false
}

// SetStorage gets a reference to the given NullableStorage and assigns it to the Storage field.
func (o *PatchInstance) SetStorage(v Storage) {
	o.Storage.Set(&v)
}
// SetStorageNil sets the value for Storage to be an explicit nil
func (o *PatchInstance) SetStorageNil() {
	o.Storage.Set(nil)
}

// UnsetStorage ensures that no value is present for Storage, not even an explicit nil
func (o *PatchInstance) UnsetStorage() {
	o.Storage.Unset()
}

// GetTrunkInstalls returns the TrunkInstalls field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchInstance) GetTrunkInstalls() []TrunkInstall {
	if o == nil {
		var ret []TrunkInstall
		return ret
	}
	return o.TrunkInstalls
}

// GetTrunkInstallsOk returns a tuple with the TrunkInstalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchInstance) GetTrunkInstallsOk() ([]TrunkInstall, bool) {
	if o == nil || IsNil(o.TrunkInstalls) {
		return nil, false
	}
	return o.TrunkInstalls, true
}

// HasTrunkInstalls returns a boolean if a field has been set.
func (o *PatchInstance) HasTrunkInstalls() bool {
	if o != nil && !IsNil(o.TrunkInstalls) {
		return true
	}

	return false
}

// SetTrunkInstalls gets a reference to the given []TrunkInstall and assigns it to the TrunkInstalls field.
func (o *PatchInstance) SetTrunkInstalls(v []TrunkInstall) {
	o.TrunkInstalls = v
}

func (o PatchInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchInstance) ToMap() (map[string]interface{}, error) {
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
	if o.DedicatedNetworking.IsSet() {
		toSerialize["dedicated_networking"] = o.DedicatedNetworking.Get()
	}
	if o.Environment.IsSet() {
		toSerialize["environment"] = o.Environment.Get()
	}
	if o.Experimental.IsSet() {
		toSerialize["experimental"] = o.Experimental.Get()
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.ExtraDomainsRw != nil {
		toSerialize["extra_domains_rw"] = o.ExtraDomainsRw
	}
	if o.InstanceName.IsSet() {
		toSerialize["instance_name"] = o.InstanceName.Get()
	}
	if o.IpAllowList != nil {
		toSerialize["ip_allow_list"] = o.IpAllowList
	}
	if o.Memory.IsSet() {
		toSerialize["memory"] = o.Memory.Get()
	}
	if o.PostgresConfigs != nil {
		toSerialize["postgres_configs"] = o.PostgresConfigs
	}
	if o.Replicas.IsSet() {
		toSerialize["replicas"] = o.Replicas.Get()
	}
	if o.Spot.IsSet() {
		toSerialize["spot"] = o.Spot.Get()
	}
	if o.Storage.IsSet() {
		toSerialize["storage"] = o.Storage.Get()
	}
	if o.TrunkInstalls != nil {
		toSerialize["trunk_installs"] = o.TrunkInstalls
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchInstance) UnmarshalJSON(data []byte) (err error) {
	varPatchInstance := _PatchInstance{}

	err = json.Unmarshal(data, &varPatchInstance)

	if err != nil {
		return err
	}

	*o = PatchInstance(varPatchInstance)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "app_services")
		delete(additionalProperties, "autoscaling")
		delete(additionalProperties, "connection_pooler")
		delete(additionalProperties, "cpu")
		delete(additionalProperties, "dedicated_networking")
		delete(additionalProperties, "environment")
		delete(additionalProperties, "experimental")
		delete(additionalProperties, "extensions")
		delete(additionalProperties, "extra_domains_rw")
		delete(additionalProperties, "instance_name")
		delete(additionalProperties, "ip_allow_list")
		delete(additionalProperties, "memory")
		delete(additionalProperties, "postgres_configs")
		delete(additionalProperties, "replicas")
		delete(additionalProperties, "spot")
		delete(additionalProperties, "storage")
		delete(additionalProperties, "trunk_installs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchInstance struct {
	value *PatchInstance
	isSet bool
}

func (v NullablePatchInstance) Get() *PatchInstance {
	return v.value
}

func (v *NullablePatchInstance) Set(val *PatchInstance) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchInstance) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchInstance(val *PatchInstance) *NullablePatchInstance {
	return &NullablePatchInstance{value: val, isSet: true}
}

func (v NullablePatchInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


