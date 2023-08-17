# UpdateCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to [**NullableCpu**](Cpu.md) |  | [optional] 
**EntityProperties** | Pointer to **interface{}** |  | [optional] 
**Environment** | Pointer to [**NullableEnvironment**](Environment.md) |  | [optional] 
**Extensions** | Pointer to [**[]Extension**](Extension.md) |  | [optional] 
**Memory** | Pointer to [**NullableMemory**](Memory.md) |  | [optional] 
**Storage** | Pointer to [**NullableStorage**](Storage.md) |  | [optional] 
**TrunkInstalls** | Pointer to [**[]TrunkInstall**](TrunkInstall.md) |  | [optional] 

## Methods

### NewUpdateCluster

`func NewUpdateCluster() *UpdateCluster`

NewUpdateCluster instantiates a new UpdateCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterWithDefaults

`func NewUpdateClusterWithDefaults() *UpdateCluster`

NewUpdateClusterWithDefaults instantiates a new UpdateCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *UpdateCluster) GetCpu() Cpu`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *UpdateCluster) GetCpuOk() (*Cpu, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *UpdateCluster) SetCpu(v Cpu)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *UpdateCluster) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *UpdateCluster) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *UpdateCluster) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetEntityProperties

`func (o *UpdateCluster) GetEntityProperties() interface{}`

GetEntityProperties returns the EntityProperties field if non-nil, zero value otherwise.

### GetEntityPropertiesOk

`func (o *UpdateCluster) GetEntityPropertiesOk() (*interface{}, bool)`

GetEntityPropertiesOk returns a tuple with the EntityProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityProperties

`func (o *UpdateCluster) SetEntityProperties(v interface{})`

SetEntityProperties sets EntityProperties field to given value.

### HasEntityProperties

`func (o *UpdateCluster) HasEntityProperties() bool`

HasEntityProperties returns a boolean if a field has been set.

### SetEntityPropertiesNil

`func (o *UpdateCluster) SetEntityPropertiesNil(b bool)`

 SetEntityPropertiesNil sets the value for EntityProperties to be an explicit nil

### UnsetEntityProperties
`func (o *UpdateCluster) UnsetEntityProperties()`

UnsetEntityProperties ensures that no value is present for EntityProperties, not even an explicit nil
### GetEnvironment

`func (o *UpdateCluster) GetEnvironment() Environment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *UpdateCluster) GetEnvironmentOk() (*Environment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *UpdateCluster) SetEnvironment(v Environment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *UpdateCluster) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *UpdateCluster) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *UpdateCluster) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetExtensions

`func (o *UpdateCluster) GetExtensions() []Extension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *UpdateCluster) GetExtensionsOk() (*[]Extension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *UpdateCluster) SetExtensions(v []Extension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *UpdateCluster) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensionsNil

`func (o *UpdateCluster) SetExtensionsNil(b bool)`

 SetExtensionsNil sets the value for Extensions to be an explicit nil

### UnsetExtensions
`func (o *UpdateCluster) UnsetExtensions()`

UnsetExtensions ensures that no value is present for Extensions, not even an explicit nil
### GetMemory

`func (o *UpdateCluster) GetMemory() Memory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *UpdateCluster) GetMemoryOk() (*Memory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *UpdateCluster) SetMemory(v Memory)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *UpdateCluster) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *UpdateCluster) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *UpdateCluster) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetStorage

`func (o *UpdateCluster) GetStorage() Storage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *UpdateCluster) GetStorageOk() (*Storage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *UpdateCluster) SetStorage(v Storage)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *UpdateCluster) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *UpdateCluster) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *UpdateCluster) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetTrunkInstalls

`func (o *UpdateCluster) GetTrunkInstalls() []TrunkInstall`

GetTrunkInstalls returns the TrunkInstalls field if non-nil, zero value otherwise.

### GetTrunkInstallsOk

`func (o *UpdateCluster) GetTrunkInstallsOk() (*[]TrunkInstall, bool)`

GetTrunkInstallsOk returns a tuple with the TrunkInstalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunkInstalls

`func (o *UpdateCluster) SetTrunkInstalls(v []TrunkInstall)`

SetTrunkInstalls sets TrunkInstalls field to given value.

### HasTrunkInstalls

`func (o *UpdateCluster) HasTrunkInstalls() bool`

HasTrunkInstalls returns a boolean if a field has been set.

### SetTrunkInstallsNil

`func (o *UpdateCluster) SetTrunkInstallsNil(b bool)`

 SetTrunkInstallsNil sets the value for TrunkInstalls to be an explicit nil

### UnsetTrunkInstalls
`func (o *UpdateCluster) UnsetTrunkInstalls()`

UnsetTrunkInstalls ensures that no value is present for TrunkInstalls, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


