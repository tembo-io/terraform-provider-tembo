# Autoscaling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autostop** | Pointer to [**AutoStop**](AutoStop.md) |  | [optional] 
**Storage** | Pointer to [**AutoscalingStorage**](AutoscalingStorage.md) |  | [optional] 

## Methods

### NewAutoscaling

`func NewAutoscaling() *Autoscaling`

NewAutoscaling instantiates a new Autoscaling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoscalingWithDefaults

`func NewAutoscalingWithDefaults() *Autoscaling`

NewAutoscalingWithDefaults instantiates a new Autoscaling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutostop

`func (o *Autoscaling) GetAutostop() AutoStop`

GetAutostop returns the Autostop field if non-nil, zero value otherwise.

### GetAutostopOk

`func (o *Autoscaling) GetAutostopOk() (*AutoStop, bool)`

GetAutostopOk returns a tuple with the Autostop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutostop

`func (o *Autoscaling) SetAutostop(v AutoStop)`

SetAutostop sets Autostop field to given value.

### HasAutostop

`func (o *Autoscaling) HasAutostop() bool`

HasAutostop returns a boolean if a field has been set.

### GetStorage

`func (o *Autoscaling) GetStorage() AutoscalingStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *Autoscaling) GetStorageOk() (*AutoscalingStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *Autoscaling) SetStorage(v AutoscalingStorage)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *Autoscaling) HasStorage() bool`

HasStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


