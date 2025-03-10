# CreateInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppServices** | Pointer to [**[]AppType**](AppType.md) |  | [optional] 
**Autoscaling** | Pointer to [**NullablePatchAutoscaling**](PatchAutoscaling.md) |  | [optional] 
**ConnectionPooler** | Pointer to [**NullableConnectionPooler**](ConnectionPooler.md) |  | [optional] 
**Cpu** | [**Cpu**](Cpu.md) |  | 
**DedicatedNetworking** | Pointer to [**NullableDedicatedNetworking**](DedicatedNetworking.md) |  | [optional] 
**Environment** | [**Environment**](Environment.md) |  | 
**Experimental** | Pointer to [**NullableExperimental**](Experimental.md) |  | [optional] 
**Extensions** | Pointer to [**[]Extension**](Extension.md) |  | [optional] 
**ExtraDomainsRw** | Pointer to **[]string** |  | [optional] 
**InstanceName** | **string** |  | 
**IpAllowList** | Pointer to **[]string** |  | [optional] 
**Memory** | [**Memory**](Memory.md) |  | 
**PgVersion** | Pointer to **int32** |  | [optional] 
**PostgresConfigs** | Pointer to [**[]PgConfig**](PgConfig.md) |  | [optional] 
**ProviderId** | Pointer to **string** |  | [optional] 
**RegionId** | Pointer to **string** |  | [optional] 
**Replicas** | Pointer to **int32** |  | [optional] 
**Spot** | Pointer to **NullableBool** |  | [optional] 
**StackType** | [**StackType**](StackType.md) |  | 
**Storage** | [**Storage**](Storage.md) |  | 
**TrunkInstalls** | Pointer to [**[]TrunkInstall**](TrunkInstall.md) |  | [optional] 

## Methods

### NewCreateInstance

`func NewCreateInstance(cpu Cpu, environment Environment, instanceName string, memory Memory, stackType StackType, storage Storage, ) *CreateInstance`

NewCreateInstance instantiates a new CreateInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInstanceWithDefaults

`func NewCreateInstanceWithDefaults() *CreateInstance`

NewCreateInstanceWithDefaults instantiates a new CreateInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppServices

`func (o *CreateInstance) GetAppServices() []AppType`

GetAppServices returns the AppServices field if non-nil, zero value otherwise.

### GetAppServicesOk

`func (o *CreateInstance) GetAppServicesOk() (*[]AppType, bool)`

GetAppServicesOk returns a tuple with the AppServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServices

`func (o *CreateInstance) SetAppServices(v []AppType)`

SetAppServices sets AppServices field to given value.

### HasAppServices

`func (o *CreateInstance) HasAppServices() bool`

HasAppServices returns a boolean if a field has been set.

### SetAppServicesNil

`func (o *CreateInstance) SetAppServicesNil(b bool)`

 SetAppServicesNil sets the value for AppServices to be an explicit nil

### UnsetAppServices
`func (o *CreateInstance) UnsetAppServices()`

UnsetAppServices ensures that no value is present for AppServices, not even an explicit nil
### GetAutoscaling

`func (o *CreateInstance) GetAutoscaling() PatchAutoscaling`

GetAutoscaling returns the Autoscaling field if non-nil, zero value otherwise.

### GetAutoscalingOk

`func (o *CreateInstance) GetAutoscalingOk() (*PatchAutoscaling, bool)`

GetAutoscalingOk returns a tuple with the Autoscaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaling

`func (o *CreateInstance) SetAutoscaling(v PatchAutoscaling)`

SetAutoscaling sets Autoscaling field to given value.

### HasAutoscaling

`func (o *CreateInstance) HasAutoscaling() bool`

HasAutoscaling returns a boolean if a field has been set.

### SetAutoscalingNil

`func (o *CreateInstance) SetAutoscalingNil(b bool)`

 SetAutoscalingNil sets the value for Autoscaling to be an explicit nil

### UnsetAutoscaling
`func (o *CreateInstance) UnsetAutoscaling()`

UnsetAutoscaling ensures that no value is present for Autoscaling, not even an explicit nil
### GetConnectionPooler

`func (o *CreateInstance) GetConnectionPooler() ConnectionPooler`

GetConnectionPooler returns the ConnectionPooler field if non-nil, zero value otherwise.

### GetConnectionPoolerOk

`func (o *CreateInstance) GetConnectionPoolerOk() (*ConnectionPooler, bool)`

GetConnectionPoolerOk returns a tuple with the ConnectionPooler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPooler

`func (o *CreateInstance) SetConnectionPooler(v ConnectionPooler)`

SetConnectionPooler sets ConnectionPooler field to given value.

### HasConnectionPooler

`func (o *CreateInstance) HasConnectionPooler() bool`

HasConnectionPooler returns a boolean if a field has been set.

### SetConnectionPoolerNil

`func (o *CreateInstance) SetConnectionPoolerNil(b bool)`

 SetConnectionPoolerNil sets the value for ConnectionPooler to be an explicit nil

### UnsetConnectionPooler
`func (o *CreateInstance) UnsetConnectionPooler()`

UnsetConnectionPooler ensures that no value is present for ConnectionPooler, not even an explicit nil
### GetCpu

`func (o *CreateInstance) GetCpu() Cpu`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *CreateInstance) GetCpuOk() (*Cpu, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *CreateInstance) SetCpu(v Cpu)`

SetCpu sets Cpu field to given value.


### GetDedicatedNetworking

`func (o *CreateInstance) GetDedicatedNetworking() DedicatedNetworking`

GetDedicatedNetworking returns the DedicatedNetworking field if non-nil, zero value otherwise.

### GetDedicatedNetworkingOk

`func (o *CreateInstance) GetDedicatedNetworkingOk() (*DedicatedNetworking, bool)`

GetDedicatedNetworkingOk returns a tuple with the DedicatedNetworking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedNetworking

`func (o *CreateInstance) SetDedicatedNetworking(v DedicatedNetworking)`

SetDedicatedNetworking sets DedicatedNetworking field to given value.

### HasDedicatedNetworking

`func (o *CreateInstance) HasDedicatedNetworking() bool`

HasDedicatedNetworking returns a boolean if a field has been set.

### SetDedicatedNetworkingNil

`func (o *CreateInstance) SetDedicatedNetworkingNil(b bool)`

 SetDedicatedNetworkingNil sets the value for DedicatedNetworking to be an explicit nil

### UnsetDedicatedNetworking
`func (o *CreateInstance) UnsetDedicatedNetworking()`

UnsetDedicatedNetworking ensures that no value is present for DedicatedNetworking, not even an explicit nil
### GetEnvironment

`func (o *CreateInstance) GetEnvironment() Environment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CreateInstance) GetEnvironmentOk() (*Environment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CreateInstance) SetEnvironment(v Environment)`

SetEnvironment sets Environment field to given value.


### GetExperimental

`func (o *CreateInstance) GetExperimental() Experimental`

GetExperimental returns the Experimental field if non-nil, zero value otherwise.

### GetExperimentalOk

`func (o *CreateInstance) GetExperimentalOk() (*Experimental, bool)`

GetExperimentalOk returns a tuple with the Experimental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimental

`func (o *CreateInstance) SetExperimental(v Experimental)`

SetExperimental sets Experimental field to given value.

### HasExperimental

`func (o *CreateInstance) HasExperimental() bool`

HasExperimental returns a boolean if a field has been set.

### SetExperimentalNil

`func (o *CreateInstance) SetExperimentalNil(b bool)`

 SetExperimentalNil sets the value for Experimental to be an explicit nil

### UnsetExperimental
`func (o *CreateInstance) UnsetExperimental()`

UnsetExperimental ensures that no value is present for Experimental, not even an explicit nil
### GetExtensions

`func (o *CreateInstance) GetExtensions() []Extension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *CreateInstance) GetExtensionsOk() (*[]Extension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *CreateInstance) SetExtensions(v []Extension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *CreateInstance) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensionsNil

`func (o *CreateInstance) SetExtensionsNil(b bool)`

 SetExtensionsNil sets the value for Extensions to be an explicit nil

### UnsetExtensions
`func (o *CreateInstance) UnsetExtensions()`

UnsetExtensions ensures that no value is present for Extensions, not even an explicit nil
### GetExtraDomainsRw

`func (o *CreateInstance) GetExtraDomainsRw() []string`

GetExtraDomainsRw returns the ExtraDomainsRw field if non-nil, zero value otherwise.

### GetExtraDomainsRwOk

`func (o *CreateInstance) GetExtraDomainsRwOk() (*[]string, bool)`

GetExtraDomainsRwOk returns a tuple with the ExtraDomainsRw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDomainsRw

`func (o *CreateInstance) SetExtraDomainsRw(v []string)`

SetExtraDomainsRw sets ExtraDomainsRw field to given value.

### HasExtraDomainsRw

`func (o *CreateInstance) HasExtraDomainsRw() bool`

HasExtraDomainsRw returns a boolean if a field has been set.

### SetExtraDomainsRwNil

`func (o *CreateInstance) SetExtraDomainsRwNil(b bool)`

 SetExtraDomainsRwNil sets the value for ExtraDomainsRw to be an explicit nil

### UnsetExtraDomainsRw
`func (o *CreateInstance) UnsetExtraDomainsRw()`

UnsetExtraDomainsRw ensures that no value is present for ExtraDomainsRw, not even an explicit nil
### GetInstanceName

`func (o *CreateInstance) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *CreateInstance) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *CreateInstance) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.


### GetIpAllowList

`func (o *CreateInstance) GetIpAllowList() []string`

GetIpAllowList returns the IpAllowList field if non-nil, zero value otherwise.

### GetIpAllowListOk

`func (o *CreateInstance) GetIpAllowListOk() (*[]string, bool)`

GetIpAllowListOk returns a tuple with the IpAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAllowList

`func (o *CreateInstance) SetIpAllowList(v []string)`

SetIpAllowList sets IpAllowList field to given value.

### HasIpAllowList

`func (o *CreateInstance) HasIpAllowList() bool`

HasIpAllowList returns a boolean if a field has been set.

### SetIpAllowListNil

`func (o *CreateInstance) SetIpAllowListNil(b bool)`

 SetIpAllowListNil sets the value for IpAllowList to be an explicit nil

### UnsetIpAllowList
`func (o *CreateInstance) UnsetIpAllowList()`

UnsetIpAllowList ensures that no value is present for IpAllowList, not even an explicit nil
### GetMemory

`func (o *CreateInstance) GetMemory() Memory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *CreateInstance) GetMemoryOk() (*Memory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *CreateInstance) SetMemory(v Memory)`

SetMemory sets Memory field to given value.


### GetPgVersion

`func (o *CreateInstance) GetPgVersion() int32`

GetPgVersion returns the PgVersion field if non-nil, zero value otherwise.

### GetPgVersionOk

`func (o *CreateInstance) GetPgVersionOk() (*int32, bool)`

GetPgVersionOk returns a tuple with the PgVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgVersion

`func (o *CreateInstance) SetPgVersion(v int32)`

SetPgVersion sets PgVersion field to given value.

### HasPgVersion

`func (o *CreateInstance) HasPgVersion() bool`

HasPgVersion returns a boolean if a field has been set.

### GetPostgresConfigs

`func (o *CreateInstance) GetPostgresConfigs() []PgConfig`

GetPostgresConfigs returns the PostgresConfigs field if non-nil, zero value otherwise.

### GetPostgresConfigsOk

`func (o *CreateInstance) GetPostgresConfigsOk() (*[]PgConfig, bool)`

GetPostgresConfigsOk returns a tuple with the PostgresConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresConfigs

`func (o *CreateInstance) SetPostgresConfigs(v []PgConfig)`

SetPostgresConfigs sets PostgresConfigs field to given value.

### HasPostgresConfigs

`func (o *CreateInstance) HasPostgresConfigs() bool`

HasPostgresConfigs returns a boolean if a field has been set.

### SetPostgresConfigsNil

`func (o *CreateInstance) SetPostgresConfigsNil(b bool)`

 SetPostgresConfigsNil sets the value for PostgresConfigs to be an explicit nil

### UnsetPostgresConfigs
`func (o *CreateInstance) UnsetPostgresConfigs()`

UnsetPostgresConfigs ensures that no value is present for PostgresConfigs, not even an explicit nil
### GetProviderId

`func (o *CreateInstance) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *CreateInstance) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *CreateInstance) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *CreateInstance) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetRegionId

`func (o *CreateInstance) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *CreateInstance) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *CreateInstance) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *CreateInstance) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetReplicas

`func (o *CreateInstance) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *CreateInstance) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *CreateInstance) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *CreateInstance) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetSpot

`func (o *CreateInstance) GetSpot() bool`

GetSpot returns the Spot field if non-nil, zero value otherwise.

### GetSpotOk

`func (o *CreateInstance) GetSpotOk() (*bool, bool)`

GetSpotOk returns a tuple with the Spot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpot

`func (o *CreateInstance) SetSpot(v bool)`

SetSpot sets Spot field to given value.

### HasSpot

`func (o *CreateInstance) HasSpot() bool`

HasSpot returns a boolean if a field has been set.

### SetSpotNil

`func (o *CreateInstance) SetSpotNil(b bool)`

 SetSpotNil sets the value for Spot to be an explicit nil

### UnsetSpot
`func (o *CreateInstance) UnsetSpot()`

UnsetSpot ensures that no value is present for Spot, not even an explicit nil
### GetStackType

`func (o *CreateInstance) GetStackType() StackType`

GetStackType returns the StackType field if non-nil, zero value otherwise.

### GetStackTypeOk

`func (o *CreateInstance) GetStackTypeOk() (*StackType, bool)`

GetStackTypeOk returns a tuple with the StackType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackType

`func (o *CreateInstance) SetStackType(v StackType)`

SetStackType sets StackType field to given value.


### GetStorage

`func (o *CreateInstance) GetStorage() Storage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *CreateInstance) GetStorageOk() (*Storage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *CreateInstance) SetStorage(v Storage)`

SetStorage sets Storage field to given value.


### GetTrunkInstalls

`func (o *CreateInstance) GetTrunkInstalls() []TrunkInstall`

GetTrunkInstalls returns the TrunkInstalls field if non-nil, zero value otherwise.

### GetTrunkInstallsOk

`func (o *CreateInstance) GetTrunkInstallsOk() (*[]TrunkInstall, bool)`

GetTrunkInstallsOk returns a tuple with the TrunkInstalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunkInstalls

`func (o *CreateInstance) SetTrunkInstalls(v []TrunkInstall)`

SetTrunkInstalls sets TrunkInstalls field to given value.

### HasTrunkInstalls

`func (o *CreateInstance) HasTrunkInstalls() bool`

HasTrunkInstalls returns a boolean if a field has been set.

### SetTrunkInstallsNil

`func (o *CreateInstance) SetTrunkInstallsNil(b bool)`

 SetTrunkInstallsNil sets the value for TrunkInstalls to be an explicit nil

### UnsetTrunkInstalls
`func (o *CreateInstance) UnsetTrunkInstalls()`

UnsetTrunkInstalls ensures that no value is present for TrunkInstalls, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


