# CreateCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | [**Cpu**](Cpu.md) |  | 
**EntityProperties** | Pointer to **interface{}** |  | [optional] 
**EntityType** | Pointer to [**EntityType**](EntityType.md) |  | [optional] 
**Environment** | [**Environment**](Environment.md) |  | 
**InstanceName** | **string** |  | 
**Memory** | [**Memory**](Memory.md) |  | 
**OrganizationId** | Pointer to **NullableString** |  | [optional] 
**Storage** | [**Storage**](Storage.md) |  | 

## Methods

### NewCreateCluster

`func NewCreateCluster(cpu Cpu, environment Environment, instanceName string, memory Memory, storage Storage, ) *CreateCluster`

NewCreateCluster instantiates a new CreateCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterWithDefaults

`func NewCreateClusterWithDefaults() *CreateCluster`

NewCreateClusterWithDefaults instantiates a new CreateCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *CreateCluster) GetCpu() Cpu`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *CreateCluster) GetCpuOk() (*Cpu, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *CreateCluster) SetCpu(v Cpu)`

SetCpu sets Cpu field to given value.


### GetEntityProperties

`func (o *CreateCluster) GetEntityProperties() interface{}`

GetEntityProperties returns the EntityProperties field if non-nil, zero value otherwise.

### GetEntityPropertiesOk

`func (o *CreateCluster) GetEntityPropertiesOk() (*interface{}, bool)`

GetEntityPropertiesOk returns a tuple with the EntityProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityProperties

`func (o *CreateCluster) SetEntityProperties(v interface{})`

SetEntityProperties sets EntityProperties field to given value.

### HasEntityProperties

`func (o *CreateCluster) HasEntityProperties() bool`

HasEntityProperties returns a boolean if a field has been set.

### SetEntityPropertiesNil

`func (o *CreateCluster) SetEntityPropertiesNil(b bool)`

 SetEntityPropertiesNil sets the value for EntityProperties to be an explicit nil

### UnsetEntityProperties
`func (o *CreateCluster) UnsetEntityProperties()`

UnsetEntityProperties ensures that no value is present for EntityProperties, not even an explicit nil
### GetEntityType

`func (o *CreateCluster) GetEntityType() EntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *CreateCluster) GetEntityTypeOk() (*EntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *CreateCluster) SetEntityType(v EntityType)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *CreateCluster) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetEnvironment

`func (o *CreateCluster) GetEnvironment() Environment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CreateCluster) GetEnvironmentOk() (*Environment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CreateCluster) SetEnvironment(v Environment)`

SetEnvironment sets Environment field to given value.


### GetInstanceName

`func (o *CreateCluster) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *CreateCluster) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *CreateCluster) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.


### GetMemory

`func (o *CreateCluster) GetMemory() Memory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *CreateCluster) GetMemoryOk() (*Memory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *CreateCluster) SetMemory(v Memory)`

SetMemory sets Memory field to given value.


### GetOrganizationId

`func (o *CreateCluster) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateCluster) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateCluster) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CreateCluster) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *CreateCluster) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *CreateCluster) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetStorage

`func (o *CreateCluster) GetStorage() Storage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *CreateCluster) GetStorageOk() (*Storage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *CreateCluster) SetStorage(v Storage)`

SetStorage sets Storage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


