# PatchInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppServices** | Pointer to [**[]AppType**](AppType.md) |  | [optional] 
**ConnectionPooler** | Pointer to [**NullableConnectionPooler**](ConnectionPooler.md) |  | [optional] 
**Cpu** | Pointer to [**NullableCpu**](Cpu.md) |  | [optional] 
**Environment** | Pointer to [**NullableEnvironment**](Environment.md) |  | [optional] 
**Experimental** | Pointer to [**NullableExperimental**](Experimental.md) |  | [optional] 
**Extensions** | Pointer to [**[]Extension**](Extension.md) |  | [optional] 
**ExtraDomainsRw** | Pointer to **[]string** |  | [optional] 
**IpAllowList** | Pointer to **[]string** |  | [optional] 
**Memory** | Pointer to [**NullableMemory**](Memory.md) |  | [optional] 
**PostgresConfigs** | Pointer to [**[]PgConfig**](PgConfig.md) |  | [optional] 
**Replicas** | Pointer to **NullableInt32** |  | [optional] 
**Spot** | Pointer to **NullableBool** |  | [optional] 
**Storage** | Pointer to [**NullableStorage**](Storage.md) |  | [optional] 
**TrunkInstalls** | Pointer to [**[]TrunkInstall**](TrunkInstall.md) |  | [optional] 

## Methods

### NewPatchInstance

`func NewPatchInstance() *PatchInstance`

NewPatchInstance instantiates a new PatchInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchInstanceWithDefaults

`func NewPatchInstanceWithDefaults() *PatchInstance`

NewPatchInstanceWithDefaults instantiates a new PatchInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppServices

`func (o *PatchInstance) GetAppServices() []AppType`

GetAppServices returns the AppServices field if non-nil, zero value otherwise.

### GetAppServicesOk

`func (o *PatchInstance) GetAppServicesOk() (*[]AppType, bool)`

GetAppServicesOk returns a tuple with the AppServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServices

`func (o *PatchInstance) SetAppServices(v []AppType)`

SetAppServices sets AppServices field to given value.

### HasAppServices

`func (o *PatchInstance) HasAppServices() bool`

HasAppServices returns a boolean if a field has been set.

### SetAppServicesNil

`func (o *PatchInstance) SetAppServicesNil(b bool)`

 SetAppServicesNil sets the value for AppServices to be an explicit nil

### UnsetAppServices
`func (o *PatchInstance) UnsetAppServices()`

UnsetAppServices ensures that no value is present for AppServices, not even an explicit nil
### GetConnectionPooler

`func (o *PatchInstance) GetConnectionPooler() ConnectionPooler`

GetConnectionPooler returns the ConnectionPooler field if non-nil, zero value otherwise.

### GetConnectionPoolerOk

`func (o *PatchInstance) GetConnectionPoolerOk() (*ConnectionPooler, bool)`

GetConnectionPoolerOk returns a tuple with the ConnectionPooler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPooler

`func (o *PatchInstance) SetConnectionPooler(v ConnectionPooler)`

SetConnectionPooler sets ConnectionPooler field to given value.

### HasConnectionPooler

`func (o *PatchInstance) HasConnectionPooler() bool`

HasConnectionPooler returns a boolean if a field has been set.

### SetConnectionPoolerNil

`func (o *PatchInstance) SetConnectionPoolerNil(b bool)`

 SetConnectionPoolerNil sets the value for ConnectionPooler to be an explicit nil

### UnsetConnectionPooler
`func (o *PatchInstance) UnsetConnectionPooler()`

UnsetConnectionPooler ensures that no value is present for ConnectionPooler, not even an explicit nil
### GetCpu

`func (o *PatchInstance) GetCpu() Cpu`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *PatchInstance) GetCpuOk() (*Cpu, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *PatchInstance) SetCpu(v Cpu)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *PatchInstance) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *PatchInstance) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *PatchInstance) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetEnvironment

`func (o *PatchInstance) GetEnvironment() Environment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PatchInstance) GetEnvironmentOk() (*Environment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PatchInstance) SetEnvironment(v Environment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *PatchInstance) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *PatchInstance) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *PatchInstance) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetExperimental

`func (o *PatchInstance) GetExperimental() Experimental`

GetExperimental returns the Experimental field if non-nil, zero value otherwise.

### GetExperimentalOk

`func (o *PatchInstance) GetExperimentalOk() (*Experimental, bool)`

GetExperimentalOk returns a tuple with the Experimental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimental

`func (o *PatchInstance) SetExperimental(v Experimental)`

SetExperimental sets Experimental field to given value.

### HasExperimental

`func (o *PatchInstance) HasExperimental() bool`

HasExperimental returns a boolean if a field has been set.

### SetExperimentalNil

`func (o *PatchInstance) SetExperimentalNil(b bool)`

 SetExperimentalNil sets the value for Experimental to be an explicit nil

### UnsetExperimental
`func (o *PatchInstance) UnsetExperimental()`

UnsetExperimental ensures that no value is present for Experimental, not even an explicit nil
### GetExtensions

`func (o *PatchInstance) GetExtensions() []Extension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *PatchInstance) GetExtensionsOk() (*[]Extension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *PatchInstance) SetExtensions(v []Extension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *PatchInstance) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensionsNil

`func (o *PatchInstance) SetExtensionsNil(b bool)`

 SetExtensionsNil sets the value for Extensions to be an explicit nil

### UnsetExtensions
`func (o *PatchInstance) UnsetExtensions()`

UnsetExtensions ensures that no value is present for Extensions, not even an explicit nil
### GetExtraDomainsRw

`func (o *PatchInstance) GetExtraDomainsRw() []string`

GetExtraDomainsRw returns the ExtraDomainsRw field if non-nil, zero value otherwise.

### GetExtraDomainsRwOk

`func (o *PatchInstance) GetExtraDomainsRwOk() (*[]string, bool)`

GetExtraDomainsRwOk returns a tuple with the ExtraDomainsRw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDomainsRw

`func (o *PatchInstance) SetExtraDomainsRw(v []string)`

SetExtraDomainsRw sets ExtraDomainsRw field to given value.

### HasExtraDomainsRw

`func (o *PatchInstance) HasExtraDomainsRw() bool`

HasExtraDomainsRw returns a boolean if a field has been set.

### SetExtraDomainsRwNil

`func (o *PatchInstance) SetExtraDomainsRwNil(b bool)`

 SetExtraDomainsRwNil sets the value for ExtraDomainsRw to be an explicit nil

### UnsetExtraDomainsRw
`func (o *PatchInstance) UnsetExtraDomainsRw()`

UnsetExtraDomainsRw ensures that no value is present for ExtraDomainsRw, not even an explicit nil
### GetIpAllowList

`func (o *PatchInstance) GetIpAllowList() []string`

GetIpAllowList returns the IpAllowList field if non-nil, zero value otherwise.

### GetIpAllowListOk

`func (o *PatchInstance) GetIpAllowListOk() (*[]string, bool)`

GetIpAllowListOk returns a tuple with the IpAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAllowList

`func (o *PatchInstance) SetIpAllowList(v []string)`

SetIpAllowList sets IpAllowList field to given value.

### HasIpAllowList

`func (o *PatchInstance) HasIpAllowList() bool`

HasIpAllowList returns a boolean if a field has been set.

### SetIpAllowListNil

`func (o *PatchInstance) SetIpAllowListNil(b bool)`

 SetIpAllowListNil sets the value for IpAllowList to be an explicit nil

### UnsetIpAllowList
`func (o *PatchInstance) UnsetIpAllowList()`

UnsetIpAllowList ensures that no value is present for IpAllowList, not even an explicit nil
### GetMemory

`func (o *PatchInstance) GetMemory() Memory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *PatchInstance) GetMemoryOk() (*Memory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *PatchInstance) SetMemory(v Memory)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *PatchInstance) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *PatchInstance) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *PatchInstance) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetPostgresConfigs

`func (o *PatchInstance) GetPostgresConfigs() []PgConfig`

GetPostgresConfigs returns the PostgresConfigs field if non-nil, zero value otherwise.

### GetPostgresConfigsOk

`func (o *PatchInstance) GetPostgresConfigsOk() (*[]PgConfig, bool)`

GetPostgresConfigsOk returns a tuple with the PostgresConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresConfigs

`func (o *PatchInstance) SetPostgresConfigs(v []PgConfig)`

SetPostgresConfigs sets PostgresConfigs field to given value.

### HasPostgresConfigs

`func (o *PatchInstance) HasPostgresConfigs() bool`

HasPostgresConfigs returns a boolean if a field has been set.

### SetPostgresConfigsNil

`func (o *PatchInstance) SetPostgresConfigsNil(b bool)`

 SetPostgresConfigsNil sets the value for PostgresConfigs to be an explicit nil

### UnsetPostgresConfigs
`func (o *PatchInstance) UnsetPostgresConfigs()`

UnsetPostgresConfigs ensures that no value is present for PostgresConfigs, not even an explicit nil
### GetReplicas

`func (o *PatchInstance) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *PatchInstance) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *PatchInstance) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *PatchInstance) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### SetReplicasNil

`func (o *PatchInstance) SetReplicasNil(b bool)`

 SetReplicasNil sets the value for Replicas to be an explicit nil

### UnsetReplicas
`func (o *PatchInstance) UnsetReplicas()`

UnsetReplicas ensures that no value is present for Replicas, not even an explicit nil
### GetSpot

`func (o *PatchInstance) GetSpot() bool`

GetSpot returns the Spot field if non-nil, zero value otherwise.

### GetSpotOk

`func (o *PatchInstance) GetSpotOk() (*bool, bool)`

GetSpotOk returns a tuple with the Spot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpot

`func (o *PatchInstance) SetSpot(v bool)`

SetSpot sets Spot field to given value.

### HasSpot

`func (o *PatchInstance) HasSpot() bool`

HasSpot returns a boolean if a field has been set.

### SetSpotNil

`func (o *PatchInstance) SetSpotNil(b bool)`

 SetSpotNil sets the value for Spot to be an explicit nil

### UnsetSpot
`func (o *PatchInstance) UnsetSpot()`

UnsetSpot ensures that no value is present for Spot, not even an explicit nil
### GetStorage

`func (o *PatchInstance) GetStorage() Storage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *PatchInstance) GetStorageOk() (*Storage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *PatchInstance) SetStorage(v Storage)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *PatchInstance) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *PatchInstance) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *PatchInstance) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetTrunkInstalls

`func (o *PatchInstance) GetTrunkInstalls() []TrunkInstall`

GetTrunkInstalls returns the TrunkInstalls field if non-nil, zero value otherwise.

### GetTrunkInstallsOk

`func (o *PatchInstance) GetTrunkInstallsOk() (*[]TrunkInstall, bool)`

GetTrunkInstallsOk returns a tuple with the TrunkInstalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunkInstalls

`func (o *PatchInstance) SetTrunkInstalls(v []TrunkInstall)`

SetTrunkInstalls sets TrunkInstalls field to given value.

### HasTrunkInstalls

`func (o *PatchInstance) HasTrunkInstalls() bool`

HasTrunkInstalls returns a boolean if a field has been set.

### SetTrunkInstallsNil

`func (o *PatchInstance) SetTrunkInstallsNil(b bool)`

 SetTrunkInstallsNil sets the value for TrunkInstalls to be an explicit nil

### UnsetTrunkInstalls
`func (o *PatchInstance) UnsetTrunkInstalls()`

UnsetTrunkInstalls ensures that no value is present for TrunkInstalls, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


