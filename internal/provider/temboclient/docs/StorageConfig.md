# StorageConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolumeMounts** | Pointer to [**[]VolumeMount**](VolumeMount.md) |  | [optional] 

## Methods

### NewStorageConfig

`func NewStorageConfig() *StorageConfig`

NewStorageConfig instantiates a new StorageConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageConfigWithDefaults

`func NewStorageConfigWithDefaults() *StorageConfig`

NewStorageConfigWithDefaults instantiates a new StorageConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumeMounts

`func (o *StorageConfig) GetVolumeMounts() []VolumeMount`

GetVolumeMounts returns the VolumeMounts field if non-nil, zero value otherwise.

### GetVolumeMountsOk

`func (o *StorageConfig) GetVolumeMountsOk() (*[]VolumeMount, bool)`

GetVolumeMountsOk returns a tuple with the VolumeMounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeMounts

`func (o *StorageConfig) SetVolumeMounts(v []VolumeMount)`

SetVolumeMounts sets VolumeMounts field to given value.

### HasVolumeMounts

`func (o *StorageConfig) HasVolumeMounts() bool`

HasVolumeMounts returns a boolean if a field has been set.

### SetVolumeMountsNil

`func (o *StorageConfig) SetVolumeMountsNil(b bool)`

 SetVolumeMountsNil sets the value for VolumeMounts to be an explicit nil

### UnsetVolumeMounts
`func (o *StorageConfig) UnsetVolumeMounts()`

UnsetVolumeMounts ensures that no value is present for VolumeMounts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


