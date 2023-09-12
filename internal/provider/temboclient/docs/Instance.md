# Instance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionInfo** | Pointer to [**NullableConnectionInfo**](ConnectionInfo.md) |  | [optional] 
**Cpu** | [**Cpu**](Cpu.md) |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Environment** | [**Environment**](Environment.md) |  | 
**Extensions** | Pointer to [**[]ExtensionStatus**](ExtensionStatus.md) |  | [optional] 
**ExtraDomainsRw** | Pointer to **[]string** |  | [optional] 
**InstanceId** | **string** |  | 
**InstanceName** | **string** |  | 
**LastUpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Memory** | [**Memory**](Memory.md) |  | 
**OrganizationId** | **string** |  | 
**OrganizationName** | **string** |  | 
**PostgresConfigs** | Pointer to [**[]PgConfig**](PgConfig.md) |  | [optional] 
**Replicas** | **int32** |  | 
**StackType** | [**EntityType**](EntityType.md) |  | 
**State** | [**State**](State.md) |  | 
**Storage** | [**Storage**](Storage.md) |  | 
**TrunkInstalls** | Pointer to [**[]TrunkInstallStatus**](TrunkInstallStatus.md) |  | [optional] 

## Methods

### NewInstance

`func NewInstance(cpu Cpu, environment Environment, instanceId string, instanceName string, memory Memory, organizationId string, organizationName string, replicas int32, stackType EntityType, state State, storage Storage, ) *Instance`

NewInstance instantiates a new Instance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceWithDefaults

`func NewInstanceWithDefaults() *Instance`

NewInstanceWithDefaults instantiates a new Instance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionInfo

`func (o *Instance) GetConnectionInfo() ConnectionInfo`

GetConnectionInfo returns the ConnectionInfo field if non-nil, zero value otherwise.

### GetConnectionInfoOk

`func (o *Instance) GetConnectionInfoOk() (*ConnectionInfo, bool)`

GetConnectionInfoOk returns a tuple with the ConnectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionInfo

`func (o *Instance) SetConnectionInfo(v ConnectionInfo)`

SetConnectionInfo sets ConnectionInfo field to given value.

### HasConnectionInfo

`func (o *Instance) HasConnectionInfo() bool`

HasConnectionInfo returns a boolean if a field has been set.

### SetConnectionInfoNil

`func (o *Instance) SetConnectionInfoNil(b bool)`

 SetConnectionInfoNil sets the value for ConnectionInfo to be an explicit nil

### UnsetConnectionInfo
`func (o *Instance) UnsetConnectionInfo()`

UnsetConnectionInfo ensures that no value is present for ConnectionInfo, not even an explicit nil
### GetCpu

`func (o *Instance) GetCpu() Cpu`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *Instance) GetCpuOk() (*Cpu, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *Instance) SetCpu(v Cpu)`

SetCpu sets Cpu field to given value.


### GetCreatedAt

`func (o *Instance) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Instance) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Instance) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Instance) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEnvironment

`func (o *Instance) GetEnvironment() Environment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Instance) GetEnvironmentOk() (*Environment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Instance) SetEnvironment(v Environment)`

SetEnvironment sets Environment field to given value.


### GetExtensions

`func (o *Instance) GetExtensions() []ExtensionStatus`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Instance) GetExtensionsOk() (*[]ExtensionStatus, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Instance) SetExtensions(v []ExtensionStatus)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Instance) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensionsNil

`func (o *Instance) SetExtensionsNil(b bool)`

 SetExtensionsNil sets the value for Extensions to be an explicit nil

### UnsetExtensions
`func (o *Instance) UnsetExtensions()`

UnsetExtensions ensures that no value is present for Extensions, not even an explicit nil
### GetExtraDomainsRw

`func (o *Instance) GetExtraDomainsRw() []string`

GetExtraDomainsRw returns the ExtraDomainsRw field if non-nil, zero value otherwise.

### GetExtraDomainsRwOk

`func (o *Instance) GetExtraDomainsRwOk() (*[]string, bool)`

GetExtraDomainsRwOk returns a tuple with the ExtraDomainsRw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDomainsRw

`func (o *Instance) SetExtraDomainsRw(v []string)`

SetExtraDomainsRw sets ExtraDomainsRw field to given value.

### HasExtraDomainsRw

`func (o *Instance) HasExtraDomainsRw() bool`

HasExtraDomainsRw returns a boolean if a field has been set.

### SetExtraDomainsRwNil

`func (o *Instance) SetExtraDomainsRwNil(b bool)`

 SetExtraDomainsRwNil sets the value for ExtraDomainsRw to be an explicit nil

### UnsetExtraDomainsRw
`func (o *Instance) UnsetExtraDomainsRw()`

UnsetExtraDomainsRw ensures that no value is present for ExtraDomainsRw, not even an explicit nil
### GetInstanceId

`func (o *Instance) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *Instance) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *Instance) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetInstanceName

`func (o *Instance) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *Instance) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *Instance) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.


### GetLastUpdatedAt

`func (o *Instance) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *Instance) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *Instance) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *Instance) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetMemory

`func (o *Instance) GetMemory() Memory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *Instance) GetMemoryOk() (*Memory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *Instance) SetMemory(v Memory)`

SetMemory sets Memory field to given value.


### GetOrganizationId

`func (o *Instance) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Instance) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Instance) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetOrganizationName

`func (o *Instance) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *Instance) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *Instance) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetPostgresConfigs

`func (o *Instance) GetPostgresConfigs() []PgConfig`

GetPostgresConfigs returns the PostgresConfigs field if non-nil, zero value otherwise.

### GetPostgresConfigsOk

`func (o *Instance) GetPostgresConfigsOk() (*[]PgConfig, bool)`

GetPostgresConfigsOk returns a tuple with the PostgresConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresConfigs

`func (o *Instance) SetPostgresConfigs(v []PgConfig)`

SetPostgresConfigs sets PostgresConfigs field to given value.

### HasPostgresConfigs

`func (o *Instance) HasPostgresConfigs() bool`

HasPostgresConfigs returns a boolean if a field has been set.

### SetPostgresConfigsNil

`func (o *Instance) SetPostgresConfigsNil(b bool)`

 SetPostgresConfigsNil sets the value for PostgresConfigs to be an explicit nil

### UnsetPostgresConfigs
`func (o *Instance) UnsetPostgresConfigs()`

UnsetPostgresConfigs ensures that no value is present for PostgresConfigs, not even an explicit nil
### GetReplicas

`func (o *Instance) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *Instance) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *Instance) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.


### GetStackType

`func (o *Instance) GetStackType() EntityType`

GetStackType returns the StackType field if non-nil, zero value otherwise.

### GetStackTypeOk

`func (o *Instance) GetStackTypeOk() (*EntityType, bool)`

GetStackTypeOk returns a tuple with the StackType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackType

`func (o *Instance) SetStackType(v EntityType)`

SetStackType sets StackType field to given value.


### GetState

`func (o *Instance) GetState() State`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Instance) GetStateOk() (*State, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Instance) SetState(v State)`

SetState sets State field to given value.


### GetStorage

`func (o *Instance) GetStorage() Storage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *Instance) GetStorageOk() (*Storage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *Instance) SetStorage(v Storage)`

SetStorage sets Storage field to given value.


### GetTrunkInstalls

`func (o *Instance) GetTrunkInstalls() []TrunkInstallStatus`

GetTrunkInstalls returns the TrunkInstalls field if non-nil, zero value otherwise.

### GetTrunkInstallsOk

`func (o *Instance) GetTrunkInstallsOk() (*[]TrunkInstallStatus, bool)`

GetTrunkInstallsOk returns a tuple with the TrunkInstalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunkInstalls

`func (o *Instance) SetTrunkInstalls(v []TrunkInstallStatus)`

SetTrunkInstalls sets TrunkInstalls field to given value.

### HasTrunkInstalls

`func (o *Instance) HasTrunkInstalls() bool`

HasTrunkInstalls returns a boolean if a field has been set.

### SetTrunkInstallsNil

`func (o *Instance) SetTrunkInstallsNil(b bool)`

 SetTrunkInstallsNil sets the value for TrunkInstalls to be an explicit nil

### UnsetTrunkInstalls
`func (o *Instance) UnsetTrunkInstalls()`

UnsetTrunkInstalls ensures that no value is present for TrunkInstalls, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


