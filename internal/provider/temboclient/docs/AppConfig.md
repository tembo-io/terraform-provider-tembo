# AppConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Env** | Pointer to [**[]EnvVar**](EnvVar.md) |  | [optional] 
**Resources** | Pointer to [**NullableResourceRequirements**](ResourceRequirements.md) |  | [optional] 

## Methods

### NewAppConfig

`func NewAppConfig() *AppConfig`

NewAppConfig instantiates a new AppConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppConfigWithDefaults

`func NewAppConfigWithDefaults() *AppConfig`

NewAppConfigWithDefaults instantiates a new AppConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnv

`func (o *AppConfig) GetEnv() []EnvVar`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *AppConfig) GetEnvOk() (*[]EnvVar, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *AppConfig) SetEnv(v []EnvVar)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *AppConfig) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### SetEnvNil

`func (o *AppConfig) SetEnvNil(b bool)`

 SetEnvNil sets the value for Env to be an explicit nil

### UnsetEnv
`func (o *AppConfig) UnsetEnv()`

UnsetEnv ensures that no value is present for Env, not even an explicit nil
### GetResources

`func (o *AppConfig) GetResources() ResourceRequirements`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AppConfig) GetResourcesOk() (*ResourceRequirements, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AppConfig) SetResources(v ResourceRequirements)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AppConfig) HasResources() bool`

HasResources returns a boolean if a field has been set.

### SetResourcesNil

`func (o *AppConfig) SetResourcesNil(b bool)`

 SetResourcesNil sets the value for Resources to be an explicit nil

### UnsetResources
`func (o *AppConfig) UnsetResources()`

UnsetResources ensures that no value is present for Resources, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


