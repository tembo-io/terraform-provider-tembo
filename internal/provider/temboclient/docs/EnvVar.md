# EnvVar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Value** | Pointer to **NullableString** |  | [optional] 
**ValueFromPlatform** | Pointer to [**NullableEnvVarRef**](EnvVarRef.md) |  | [optional] 

## Methods

### NewEnvVar

`func NewEnvVar(name string, ) *EnvVar`

NewEnvVar instantiates a new EnvVar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvVarWithDefaults

`func NewEnvVarWithDefaults() *EnvVar`

NewEnvVarWithDefaults instantiates a new EnvVar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvVar) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvVar) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvVar) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *EnvVar) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EnvVar) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EnvVar) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *EnvVar) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *EnvVar) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *EnvVar) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetValueFromPlatform

`func (o *EnvVar) GetValueFromPlatform() EnvVarRef`

GetValueFromPlatform returns the ValueFromPlatform field if non-nil, zero value otherwise.

### GetValueFromPlatformOk

`func (o *EnvVar) GetValueFromPlatformOk() (*EnvVarRef, bool)`

GetValueFromPlatformOk returns a tuple with the ValueFromPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFromPlatform

`func (o *EnvVar) SetValueFromPlatform(v EnvVarRef)`

SetValueFromPlatform sets ValueFromPlatform field to given value.

### HasValueFromPlatform

`func (o *EnvVar) HasValueFromPlatform() bool`

HasValueFromPlatform returns a boolean if a field has been set.

### SetValueFromPlatformNil

`func (o *EnvVar) SetValueFromPlatformNil(b bool)`

 SetValueFromPlatformNil sets the value for ValueFromPlatform to be an explicit nil

### UnsetValueFromPlatform
`func (o *EnvVar) UnsetValueFromPlatform()`

UnsetValueFromPlatform ensures that no value is present for ValueFromPlatform, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


