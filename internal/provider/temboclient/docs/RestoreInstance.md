# RestoreInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppServices** | Pointer to [**[]AppType**](AppType.md) |  | [optional] 
**ConnectionPooler** | Pointer to [**NullableConnectionPooler**](ConnectionPooler.md) |  | [optional] 
**Cpu** | Pointer to [**NullableCpu**](Cpu.md) |  | [optional] 
**Environment** | Pointer to [**NullableEnvironment**](Environment.md) |  | [optional] 
**ExtraDomainsRw** | Pointer to **[]string** |  | [optional] 
**InstanceName** | **string** |  | 
**Memory** | Pointer to [**NullableMemory**](Memory.md) |  | [optional] 
**Restore** | [**Restore**](Restore.md) |  | 
**Storage** | Pointer to [**NullableStorage**](Storage.md) |  | [optional] 

## Methods

### NewRestoreInstance

`func NewRestoreInstance(instanceName string, restore Restore, ) *RestoreInstance`

NewRestoreInstance instantiates a new RestoreInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreInstanceWithDefaults

`func NewRestoreInstanceWithDefaults() *RestoreInstance`

NewRestoreInstanceWithDefaults instantiates a new RestoreInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppServices

`func (o *RestoreInstance) GetAppServices() []AppType`

GetAppServices returns the AppServices field if non-nil, zero value otherwise.

### GetAppServicesOk

`func (o *RestoreInstance) GetAppServicesOk() (*[]AppType, bool)`

GetAppServicesOk returns a tuple with the AppServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServices

`func (o *RestoreInstance) SetAppServices(v []AppType)`

SetAppServices sets AppServices field to given value.

### HasAppServices

`func (o *RestoreInstance) HasAppServices() bool`

HasAppServices returns a boolean if a field has been set.

### SetAppServicesNil

`func (o *RestoreInstance) SetAppServicesNil(b bool)`

 SetAppServicesNil sets the value for AppServices to be an explicit nil

### UnsetAppServices
`func (o *RestoreInstance) UnsetAppServices()`

UnsetAppServices ensures that no value is present for AppServices, not even an explicit nil
### GetConnectionPooler

`func (o *RestoreInstance) GetConnectionPooler() ConnectionPooler`

GetConnectionPooler returns the ConnectionPooler field if non-nil, zero value otherwise.

### GetConnectionPoolerOk

`func (o *RestoreInstance) GetConnectionPoolerOk() (*ConnectionPooler, bool)`

GetConnectionPoolerOk returns a tuple with the ConnectionPooler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPooler

`func (o *RestoreInstance) SetConnectionPooler(v ConnectionPooler)`

SetConnectionPooler sets ConnectionPooler field to given value.

### HasConnectionPooler

`func (o *RestoreInstance) HasConnectionPooler() bool`

HasConnectionPooler returns a boolean if a field has been set.

### SetConnectionPoolerNil

`func (o *RestoreInstance) SetConnectionPoolerNil(b bool)`

 SetConnectionPoolerNil sets the value for ConnectionPooler to be an explicit nil

### UnsetConnectionPooler
`func (o *RestoreInstance) UnsetConnectionPooler()`

UnsetConnectionPooler ensures that no value is present for ConnectionPooler, not even an explicit nil
### GetCpu

`func (o *RestoreInstance) GetCpu() Cpu`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *RestoreInstance) GetCpuOk() (*Cpu, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *RestoreInstance) SetCpu(v Cpu)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *RestoreInstance) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *RestoreInstance) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *RestoreInstance) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetEnvironment

`func (o *RestoreInstance) GetEnvironment() Environment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *RestoreInstance) GetEnvironmentOk() (*Environment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *RestoreInstance) SetEnvironment(v Environment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *RestoreInstance) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *RestoreInstance) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *RestoreInstance) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetExtraDomainsRw

`func (o *RestoreInstance) GetExtraDomainsRw() []string`

GetExtraDomainsRw returns the ExtraDomainsRw field if non-nil, zero value otherwise.

### GetExtraDomainsRwOk

`func (o *RestoreInstance) GetExtraDomainsRwOk() (*[]string, bool)`

GetExtraDomainsRwOk returns a tuple with the ExtraDomainsRw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDomainsRw

`func (o *RestoreInstance) SetExtraDomainsRw(v []string)`

SetExtraDomainsRw sets ExtraDomainsRw field to given value.

### HasExtraDomainsRw

`func (o *RestoreInstance) HasExtraDomainsRw() bool`

HasExtraDomainsRw returns a boolean if a field has been set.

### SetExtraDomainsRwNil

`func (o *RestoreInstance) SetExtraDomainsRwNil(b bool)`

 SetExtraDomainsRwNil sets the value for ExtraDomainsRw to be an explicit nil

### UnsetExtraDomainsRw
`func (o *RestoreInstance) UnsetExtraDomainsRw()`

UnsetExtraDomainsRw ensures that no value is present for ExtraDomainsRw, not even an explicit nil
### GetInstanceName

`func (o *RestoreInstance) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *RestoreInstance) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *RestoreInstance) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.


### GetMemory

`func (o *RestoreInstance) GetMemory() Memory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *RestoreInstance) GetMemoryOk() (*Memory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *RestoreInstance) SetMemory(v Memory)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *RestoreInstance) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *RestoreInstance) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *RestoreInstance) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetRestore

`func (o *RestoreInstance) GetRestore() Restore`

GetRestore returns the Restore field if non-nil, zero value otherwise.

### GetRestoreOk

`func (o *RestoreInstance) GetRestoreOk() (*Restore, bool)`

GetRestoreOk returns a tuple with the Restore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestore

`func (o *RestoreInstance) SetRestore(v Restore)`

SetRestore sets Restore field to given value.


### GetStorage

`func (o *RestoreInstance) GetStorage() Storage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *RestoreInstance) GetStorageOk() (*Storage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *RestoreInstance) SetStorage(v Storage)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *RestoreInstance) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *RestoreInstance) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *RestoreInstance) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


