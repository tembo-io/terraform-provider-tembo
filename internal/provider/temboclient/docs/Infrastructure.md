# Infrastructure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | [**Cpu**](Cpu.md) |  | 
**InstanceType** | [**InstanceTypes**](InstanceTypes.md) |  | 
**Memory** | [**Memory**](Memory.md) |  | 
**Provider** | Pointer to [**CloudProvider**](CloudProvider.md) |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**StorageClass** | [**StorageClass**](StorageClass.md) |  | 
**StorageSize** | [**Storage**](Storage.md) |  | 

## Methods

### NewInfrastructure

`func NewInfrastructure(cpu Cpu, instanceType InstanceTypes, memory Memory, storageClass StorageClass, storageSize Storage, ) *Infrastructure`

NewInfrastructure instantiates a new Infrastructure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureWithDefaults

`func NewInfrastructureWithDefaults() *Infrastructure`

NewInfrastructureWithDefaults instantiates a new Infrastructure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *Infrastructure) GetCpu() Cpu`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *Infrastructure) GetCpuOk() (*Cpu, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *Infrastructure) SetCpu(v Cpu)`

SetCpu sets Cpu field to given value.


### GetInstanceType

`func (o *Infrastructure) GetInstanceType() InstanceTypes`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *Infrastructure) GetInstanceTypeOk() (*InstanceTypes, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *Infrastructure) SetInstanceType(v InstanceTypes)`

SetInstanceType sets InstanceType field to given value.


### GetMemory

`func (o *Infrastructure) GetMemory() Memory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *Infrastructure) GetMemoryOk() (*Memory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *Infrastructure) SetMemory(v Memory)`

SetMemory sets Memory field to given value.


### GetProvider

`func (o *Infrastructure) GetProvider() CloudProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Infrastructure) GetProviderOk() (*CloudProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Infrastructure) SetProvider(v CloudProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *Infrastructure) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRegion

`func (o *Infrastructure) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Infrastructure) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Infrastructure) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Infrastructure) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStorageClass

`func (o *Infrastructure) GetStorageClass() StorageClass`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *Infrastructure) GetStorageClassOk() (*StorageClass, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *Infrastructure) SetStorageClass(v StorageClass)`

SetStorageClass sets StorageClass field to given value.


### GetStorageSize

`func (o *Infrastructure) GetStorageSize() Storage`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *Infrastructure) GetStorageSizeOk() (*Storage, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *Infrastructure) SetStorageSize(v Storage)`

SetStorageSize sets StorageSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


