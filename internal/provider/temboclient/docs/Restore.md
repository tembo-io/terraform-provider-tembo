# Restore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | **string** |  | 
**RecoveryTargetTime** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewRestore

`func NewRestore(instanceId string, ) *Restore`

NewRestore instantiates a new Restore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreWithDefaults

`func NewRestoreWithDefaults() *Restore`

NewRestoreWithDefaults instantiates a new Restore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *Restore) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *Restore) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *Restore) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetRecoveryTargetTime

`func (o *Restore) GetRecoveryTargetTime() time.Time`

GetRecoveryTargetTime returns the RecoveryTargetTime field if non-nil, zero value otherwise.

### GetRecoveryTargetTimeOk

`func (o *Restore) GetRecoveryTargetTimeOk() (*time.Time, bool)`

GetRecoveryTargetTimeOk returns a tuple with the RecoveryTargetTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTargetTime

`func (o *Restore) SetRecoveryTargetTime(v time.Time)`

SetRecoveryTargetTime sets RecoveryTargetTime field to given value.

### HasRecoveryTargetTime

`func (o *Restore) HasRecoveryTargetTime() bool`

HasRecoveryTargetTime returns a boolean if a field has been set.

### SetRecoveryTargetTimeNil

`func (o *Restore) SetRecoveryTargetTimeNil(b bool)`

 SetRecoveryTargetTimeNil sets the value for RecoveryTargetTime to be an explicit nil

### UnsetRecoveryTargetTime
`func (o *Restore) UnsetRecoveryTargetTime()`

UnsetRecoveryTargetTime ensures that no value is present for RecoveryTargetTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


