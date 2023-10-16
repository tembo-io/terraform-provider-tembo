# Infrastructure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **string** |  | [optional] 
**Memory** | Pointer to **string** |  | [optional] 
**Storage** | Pointer to **string** |  | [optional] 

## Methods

### NewInfrastructure

`func NewInfrastructure() *Infrastructure`

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

`func (o *Infrastructure) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *Infrastructure) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *Infrastructure) SetCpu(v string)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *Infrastructure) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *Infrastructure) GetMemory() string`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *Infrastructure) GetMemoryOk() (*string, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *Infrastructure) SetMemory(v string)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *Infrastructure) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetStorage

`func (o *Infrastructure) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *Infrastructure) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *Infrastructure) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *Infrastructure) HasStorage() bool`

HasStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


