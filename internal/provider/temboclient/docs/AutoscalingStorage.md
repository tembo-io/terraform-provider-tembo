# AutoscalingStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Enable autoscaling for storage | 
**Maximum** | Pointer to [**NullableStorage**](Storage.md) |  | [optional] 

## Methods

### NewAutoscalingStorage

`func NewAutoscalingStorage(enabled bool, ) *AutoscalingStorage`

NewAutoscalingStorage instantiates a new AutoscalingStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoscalingStorageWithDefaults

`func NewAutoscalingStorageWithDefaults() *AutoscalingStorage`

NewAutoscalingStorageWithDefaults instantiates a new AutoscalingStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AutoscalingStorage) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AutoscalingStorage) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AutoscalingStorage) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMaximum

`func (o *AutoscalingStorage) GetMaximum() Storage`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *AutoscalingStorage) GetMaximumOk() (*Storage, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *AutoscalingStorage) SetMaximum(v Storage)`

SetMaximum sets Maximum field to given value.

### HasMaximum

`func (o *AutoscalingStorage) HasMaximum() bool`

HasMaximum returns a boolean if a field has been set.

### SetMaximumNil

`func (o *AutoscalingStorage) SetMaximumNil(b bool)`

 SetMaximumNil sets the value for Maximum to be an explicit nil

### UnsetMaximum
`func (o *AutoscalingStorage) UnsetMaximum()`

UnsetMaximum ensures that no value is present for Maximum, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


