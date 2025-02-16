# PatchAutoscaling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autostop** | Pointer to [**NullableAutoStop**](AutoStop.md) |  | [optional] 
**Storage** | Pointer to [**NullableAutoscalingStorage**](AutoscalingStorage.md) |  | [optional] 

## Methods

### NewPatchAutoscaling

`func NewPatchAutoscaling() *PatchAutoscaling`

NewPatchAutoscaling instantiates a new PatchAutoscaling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchAutoscalingWithDefaults

`func NewPatchAutoscalingWithDefaults() *PatchAutoscaling`

NewPatchAutoscalingWithDefaults instantiates a new PatchAutoscaling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutostop

`func (o *PatchAutoscaling) GetAutostop() AutoStop`

GetAutostop returns the Autostop field if non-nil, zero value otherwise.

### GetAutostopOk

`func (o *PatchAutoscaling) GetAutostopOk() (*AutoStop, bool)`

GetAutostopOk returns a tuple with the Autostop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutostop

`func (o *PatchAutoscaling) SetAutostop(v AutoStop)`

SetAutostop sets Autostop field to given value.

### HasAutostop

`func (o *PatchAutoscaling) HasAutostop() bool`

HasAutostop returns a boolean if a field has been set.

### SetAutostopNil

`func (o *PatchAutoscaling) SetAutostopNil(b bool)`

 SetAutostopNil sets the value for Autostop to be an explicit nil

### UnsetAutostop
`func (o *PatchAutoscaling) UnsetAutostop()`

UnsetAutostop ensures that no value is present for Autostop, not even an explicit nil
### GetStorage

`func (o *PatchAutoscaling) GetStorage() AutoscalingStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *PatchAutoscaling) GetStorageOk() (*AutoscalingStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *PatchAutoscaling) SetStorage(v AutoscalingStorage)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *PatchAutoscaling) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *PatchAutoscaling) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *PatchAutoscaling) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


