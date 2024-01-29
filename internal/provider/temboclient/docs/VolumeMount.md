# VolumeMount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MountPath** | **string** |  | 
**MountPropagation** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**ReadOnly** | Pointer to **NullableBool** |  | [optional] 
**SubPath** | Pointer to **NullableString** |  | [optional] 
**SubPathExpr** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewVolumeMount

`func NewVolumeMount(mountPath string, name string, ) *VolumeMount`

NewVolumeMount instantiates a new VolumeMount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeMountWithDefaults

`func NewVolumeMountWithDefaults() *VolumeMount`

NewVolumeMountWithDefaults instantiates a new VolumeMount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMountPath

`func (o *VolumeMount) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *VolumeMount) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *VolumeMount) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.


### GetMountPropagation

`func (o *VolumeMount) GetMountPropagation() string`

GetMountPropagation returns the MountPropagation field if non-nil, zero value otherwise.

### GetMountPropagationOk

`func (o *VolumeMount) GetMountPropagationOk() (*string, bool)`

GetMountPropagationOk returns a tuple with the MountPropagation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPropagation

`func (o *VolumeMount) SetMountPropagation(v string)`

SetMountPropagation sets MountPropagation field to given value.

### HasMountPropagation

`func (o *VolumeMount) HasMountPropagation() bool`

HasMountPropagation returns a boolean if a field has been set.

### SetMountPropagationNil

`func (o *VolumeMount) SetMountPropagationNil(b bool)`

 SetMountPropagationNil sets the value for MountPropagation to be an explicit nil

### UnsetMountPropagation
`func (o *VolumeMount) UnsetMountPropagation()`

UnsetMountPropagation ensures that no value is present for MountPropagation, not even an explicit nil
### GetName

`func (o *VolumeMount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeMount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeMount) SetName(v string)`

SetName sets Name field to given value.


### GetReadOnly

`func (o *VolumeMount) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *VolumeMount) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *VolumeMount) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *VolumeMount) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### SetReadOnlyNil

`func (o *VolumeMount) SetReadOnlyNil(b bool)`

 SetReadOnlyNil sets the value for ReadOnly to be an explicit nil

### UnsetReadOnly
`func (o *VolumeMount) UnsetReadOnly()`

UnsetReadOnly ensures that no value is present for ReadOnly, not even an explicit nil
### GetSubPath

`func (o *VolumeMount) GetSubPath() string`

GetSubPath returns the SubPath field if non-nil, zero value otherwise.

### GetSubPathOk

`func (o *VolumeMount) GetSubPathOk() (*string, bool)`

GetSubPathOk returns a tuple with the SubPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubPath

`func (o *VolumeMount) SetSubPath(v string)`

SetSubPath sets SubPath field to given value.

### HasSubPath

`func (o *VolumeMount) HasSubPath() bool`

HasSubPath returns a boolean if a field has been set.

### SetSubPathNil

`func (o *VolumeMount) SetSubPathNil(b bool)`

 SetSubPathNil sets the value for SubPath to be an explicit nil

### UnsetSubPath
`func (o *VolumeMount) UnsetSubPath()`

UnsetSubPath ensures that no value is present for SubPath, not even an explicit nil
### GetSubPathExpr

`func (o *VolumeMount) GetSubPathExpr() string`

GetSubPathExpr returns the SubPathExpr field if non-nil, zero value otherwise.

### GetSubPathExprOk

`func (o *VolumeMount) GetSubPathExprOk() (*string, bool)`

GetSubPathExprOk returns a tuple with the SubPathExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubPathExpr

`func (o *VolumeMount) SetSubPathExpr(v string)`

SetSubPathExpr sets SubPathExpr field to given value.

### HasSubPathExpr

`func (o *VolumeMount) HasSubPathExpr() bool`

HasSubPathExpr returns a boolean if a field has been set.

### SetSubPathExprNil

`func (o *VolumeMount) SetSubPathExprNil(b bool)`

 SetSubPathExprNil sets the value for SubPathExpr to be an explicit nil

### UnsetSubPathExpr
`func (o *VolumeMount) UnsetSubPathExpr()`

UnsetSubPathExpr ensures that no value is present for SubPathExpr, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


