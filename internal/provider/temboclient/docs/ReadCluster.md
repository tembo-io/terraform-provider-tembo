# ReadCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionInfo** | Pointer to [**NullableConnectionInfo**](ConnectionInfo.md) |  | [optional] 
**Cpu** | Pointer to [**NullableCpu**](Cpu.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**EntityProperties** | Pointer to **interface{}** |  | [optional] 
**EntityType** | [**EntityType**](EntityType.md) |  | 
**Environment** | [**Environment**](Environment.md) |  | 
**Extensions** | Pointer to [**[]ExtensionStatus**](ExtensionStatus.md) |  | [optional] 
**InstanceId** | Pointer to **NullableString** |  | [optional] 
**InstanceName** | **string** |  | 
**LastUpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Memory** | Pointer to [**NullableMemory**](Memory.md) |  | [optional] 
**OrganizationId** | Pointer to **NullableString** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**State** | Pointer to [**NullableState**](State.md) |  | [optional] 
**Storage** | Pointer to [**NullableStorage**](Storage.md) |  | [optional] 
**TrunkInstalls** | Pointer to [**[]TrunkInstallStatus**](TrunkInstallStatus.md) |  | [optional] 

## Methods

### NewReadCluster

`func NewReadCluster(entityType EntityType, environment Environment, instanceName string, ) *ReadCluster`

NewReadCluster instantiates a new ReadCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadClusterWithDefaults

`func NewReadClusterWithDefaults() *ReadCluster`

NewReadClusterWithDefaults instantiates a new ReadCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionInfo

`func (o *ReadCluster) GetConnectionInfo() ConnectionInfo`

GetConnectionInfo returns the ConnectionInfo field if non-nil, zero value otherwise.

### GetConnectionInfoOk

`func (o *ReadCluster) GetConnectionInfoOk() (*ConnectionInfo, bool)`

GetConnectionInfoOk returns a tuple with the ConnectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionInfo

`func (o *ReadCluster) SetConnectionInfo(v ConnectionInfo)`

SetConnectionInfo sets ConnectionInfo field to given value.

### HasConnectionInfo

`func (o *ReadCluster) HasConnectionInfo() bool`

HasConnectionInfo returns a boolean if a field has been set.

### SetConnectionInfoNil

`func (o *ReadCluster) SetConnectionInfoNil(b bool)`

 SetConnectionInfoNil sets the value for ConnectionInfo to be an explicit nil

### UnsetConnectionInfo
`func (o *ReadCluster) UnsetConnectionInfo()`

UnsetConnectionInfo ensures that no value is present for ConnectionInfo, not even an explicit nil
### GetCpu

`func (o *ReadCluster) GetCpu() Cpu`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ReadCluster) GetCpuOk() (*Cpu, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ReadCluster) SetCpu(v Cpu)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ReadCluster) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *ReadCluster) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *ReadCluster) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetCreatedAt

`func (o *ReadCluster) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReadCluster) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReadCluster) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ReadCluster) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEntityProperties

`func (o *ReadCluster) GetEntityProperties() interface{}`

GetEntityProperties returns the EntityProperties field if non-nil, zero value otherwise.

### GetEntityPropertiesOk

`func (o *ReadCluster) GetEntityPropertiesOk() (*interface{}, bool)`

GetEntityPropertiesOk returns a tuple with the EntityProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityProperties

`func (o *ReadCluster) SetEntityProperties(v interface{})`

SetEntityProperties sets EntityProperties field to given value.

### HasEntityProperties

`func (o *ReadCluster) HasEntityProperties() bool`

HasEntityProperties returns a boolean if a field has been set.

### SetEntityPropertiesNil

`func (o *ReadCluster) SetEntityPropertiesNil(b bool)`

 SetEntityPropertiesNil sets the value for EntityProperties to be an explicit nil

### UnsetEntityProperties
`func (o *ReadCluster) UnsetEntityProperties()`

UnsetEntityProperties ensures that no value is present for EntityProperties, not even an explicit nil
### GetEntityType

`func (o *ReadCluster) GetEntityType() EntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *ReadCluster) GetEntityTypeOk() (*EntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *ReadCluster) SetEntityType(v EntityType)`

SetEntityType sets EntityType field to given value.


### GetEnvironment

`func (o *ReadCluster) GetEnvironment() Environment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ReadCluster) GetEnvironmentOk() (*Environment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ReadCluster) SetEnvironment(v Environment)`

SetEnvironment sets Environment field to given value.


### GetExtensions

`func (o *ReadCluster) GetExtensions() []ExtensionStatus`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *ReadCluster) GetExtensionsOk() (*[]ExtensionStatus, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *ReadCluster) SetExtensions(v []ExtensionStatus)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *ReadCluster) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensionsNil

`func (o *ReadCluster) SetExtensionsNil(b bool)`

 SetExtensionsNil sets the value for Extensions to be an explicit nil

### UnsetExtensions
`func (o *ReadCluster) UnsetExtensions()`

UnsetExtensions ensures that no value is present for Extensions, not even an explicit nil
### GetInstanceId

`func (o *ReadCluster) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ReadCluster) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ReadCluster) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *ReadCluster) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### SetInstanceIdNil

`func (o *ReadCluster) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *ReadCluster) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
### GetInstanceName

`func (o *ReadCluster) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *ReadCluster) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *ReadCluster) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.


### GetLastUpdatedAt

`func (o *ReadCluster) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *ReadCluster) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *ReadCluster) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *ReadCluster) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetMemory

`func (o *ReadCluster) GetMemory() Memory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ReadCluster) GetMemoryOk() (*Memory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ReadCluster) SetMemory(v Memory)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ReadCluster) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *ReadCluster) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *ReadCluster) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetOrganizationId

`func (o *ReadCluster) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ReadCluster) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ReadCluster) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ReadCluster) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *ReadCluster) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *ReadCluster) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetOrganizationName

`func (o *ReadCluster) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ReadCluster) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ReadCluster) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *ReadCluster) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *ReadCluster) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *ReadCluster) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetState

`func (o *ReadCluster) GetState() State`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ReadCluster) GetStateOk() (*State, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ReadCluster) SetState(v State)`

SetState sets State field to given value.

### HasState

`func (o *ReadCluster) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *ReadCluster) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *ReadCluster) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetStorage

`func (o *ReadCluster) GetStorage() Storage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ReadCluster) GetStorageOk() (*Storage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ReadCluster) SetStorage(v Storage)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *ReadCluster) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *ReadCluster) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *ReadCluster) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetTrunkInstalls

`func (o *ReadCluster) GetTrunkInstalls() []TrunkInstallStatus`

GetTrunkInstalls returns the TrunkInstalls field if non-nil, zero value otherwise.

### GetTrunkInstallsOk

`func (o *ReadCluster) GetTrunkInstallsOk() (*[]TrunkInstallStatus, bool)`

GetTrunkInstallsOk returns a tuple with the TrunkInstalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunkInstalls

`func (o *ReadCluster) SetTrunkInstalls(v []TrunkInstallStatus)`

SetTrunkInstalls sets TrunkInstalls field to given value.

### HasTrunkInstalls

`func (o *ReadCluster) HasTrunkInstalls() bool`

HasTrunkInstalls returns a boolean if a field has been set.

### SetTrunkInstallsNil

`func (o *ReadCluster) SetTrunkInstallsNil(b bool)`

 SetTrunkInstallsNil sets the value for TrunkInstalls to be an explicit nil

### UnsetTrunkInstalls
`func (o *ReadCluster) UnsetTrunkInstalls()`

UnsetTrunkInstalls ensures that no value is present for TrunkInstalls, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


