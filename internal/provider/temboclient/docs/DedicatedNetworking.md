# DedicatedNetworking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **NullableBool** | Enables or disables dedicated networking. Default is false. | [optional] 
**IncludeStandby** | Pointer to **NullableBool** | If true, includes the standby instance in the dedicated networking setup. Default is false. | [optional] 

## Methods

### NewDedicatedNetworking

`func NewDedicatedNetworking() *DedicatedNetworking`

NewDedicatedNetworking instantiates a new DedicatedNetworking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDedicatedNetworkingWithDefaults

`func NewDedicatedNetworkingWithDefaults() *DedicatedNetworking`

NewDedicatedNetworkingWithDefaults instantiates a new DedicatedNetworking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DedicatedNetworking) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DedicatedNetworking) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DedicatedNetworking) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DedicatedNetworking) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *DedicatedNetworking) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *DedicatedNetworking) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetIncludeStandby

`func (o *DedicatedNetworking) GetIncludeStandby() bool`

GetIncludeStandby returns the IncludeStandby field if non-nil, zero value otherwise.

### GetIncludeStandbyOk

`func (o *DedicatedNetworking) GetIncludeStandbyOk() (*bool, bool)`

GetIncludeStandbyOk returns a tuple with the IncludeStandby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeStandby

`func (o *DedicatedNetworking) SetIncludeStandby(v bool)`

SetIncludeStandby sets IncludeStandby field to given value.

### HasIncludeStandby

`func (o *DedicatedNetworking) HasIncludeStandby() bool`

HasIncludeStandby returns a boolean if a field has been set.

### SetIncludeStandbyNil

`func (o *DedicatedNetworking) SetIncludeStandbyNil(b bool)`

 SetIncludeStandbyNil sets the value for IncludeStandby to be an explicit nil

### UnsetIncludeStandby
`func (o *DedicatedNetworking) UnsetIncludeStandby()`

UnsetIncludeStandby ensures that no value is present for IncludeStandby, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


