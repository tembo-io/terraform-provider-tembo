# AppService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **[]string** |  | [optional] 
**Command** | Pointer to **[]string** |  | [optional] 
**Env** | Pointer to [**[]EnvVar**](EnvVar.md) |  | [optional] 
**Image** | **string** |  | 
**Middlewares** | Pointer to [**[]Middleware**](Middleware.md) |  | [optional] 
**Name** | **string** |  | 
**Probes** | Pointer to [**NullableProbes**](Probes.md) |  | [optional] 
**Resources** | Pointer to [**NullableResourceRequirements**](ResourceRequirements.md) |  | [optional] 
**Routing** | Pointer to [**[]Routing**](Routing.md) |  | [optional] 

## Methods

### NewAppService

`func NewAppService(image string, name string, ) *AppService`

NewAppService instantiates a new AppService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceWithDefaults

`func NewAppServiceWithDefaults() *AppService`

NewAppServiceWithDefaults instantiates a new AppService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *AppService) GetArgs() []string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *AppService) GetArgsOk() (*[]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *AppService) SetArgs(v []string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *AppService) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### SetArgsNil

`func (o *AppService) SetArgsNil(b bool)`

 SetArgsNil sets the value for Args to be an explicit nil

### UnsetArgs
`func (o *AppService) UnsetArgs()`

UnsetArgs ensures that no value is present for Args, not even an explicit nil
### GetCommand

`func (o *AppService) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *AppService) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *AppService) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *AppService) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### SetCommandNil

`func (o *AppService) SetCommandNil(b bool)`

 SetCommandNil sets the value for Command to be an explicit nil

### UnsetCommand
`func (o *AppService) UnsetCommand()`

UnsetCommand ensures that no value is present for Command, not even an explicit nil
### GetEnv

`func (o *AppService) GetEnv() []EnvVar`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *AppService) GetEnvOk() (*[]EnvVar, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *AppService) SetEnv(v []EnvVar)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *AppService) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### SetEnvNil

`func (o *AppService) SetEnvNil(b bool)`

 SetEnvNil sets the value for Env to be an explicit nil

### UnsetEnv
`func (o *AppService) UnsetEnv()`

UnsetEnv ensures that no value is present for Env, not even an explicit nil
### GetImage

`func (o *AppService) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AppService) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AppService) SetImage(v string)`

SetImage sets Image field to given value.


### GetMiddlewares

`func (o *AppService) GetMiddlewares() []Middleware`

GetMiddlewares returns the Middlewares field if non-nil, zero value otherwise.

### GetMiddlewaresOk

`func (o *AppService) GetMiddlewaresOk() (*[]Middleware, bool)`

GetMiddlewaresOk returns a tuple with the Middlewares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddlewares

`func (o *AppService) SetMiddlewares(v []Middleware)`

SetMiddlewares sets Middlewares field to given value.

### HasMiddlewares

`func (o *AppService) HasMiddlewares() bool`

HasMiddlewares returns a boolean if a field has been set.

### SetMiddlewaresNil

`func (o *AppService) SetMiddlewaresNil(b bool)`

 SetMiddlewaresNil sets the value for Middlewares to be an explicit nil

### UnsetMiddlewares
`func (o *AppService) UnsetMiddlewares()`

UnsetMiddlewares ensures that no value is present for Middlewares, not even an explicit nil
### GetName

`func (o *AppService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppService) SetName(v string)`

SetName sets Name field to given value.


### GetProbes

`func (o *AppService) GetProbes() Probes`

GetProbes returns the Probes field if non-nil, zero value otherwise.

### GetProbesOk

`func (o *AppService) GetProbesOk() (*Probes, bool)`

GetProbesOk returns a tuple with the Probes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbes

`func (o *AppService) SetProbes(v Probes)`

SetProbes sets Probes field to given value.

### HasProbes

`func (o *AppService) HasProbes() bool`

HasProbes returns a boolean if a field has been set.

### SetProbesNil

`func (o *AppService) SetProbesNil(b bool)`

 SetProbesNil sets the value for Probes to be an explicit nil

### UnsetProbes
`func (o *AppService) UnsetProbes()`

UnsetProbes ensures that no value is present for Probes, not even an explicit nil
### GetResources

`func (o *AppService) GetResources() ResourceRequirements`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AppService) GetResourcesOk() (*ResourceRequirements, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AppService) SetResources(v ResourceRequirements)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AppService) HasResources() bool`

HasResources returns a boolean if a field has been set.

### SetResourcesNil

`func (o *AppService) SetResourcesNil(b bool)`

 SetResourcesNil sets the value for Resources to be an explicit nil

### UnsetResources
`func (o *AppService) UnsetResources()`

UnsetResources ensures that no value is present for Resources, not even an explicit nil
### GetRouting

`func (o *AppService) GetRouting() []Routing`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *AppService) GetRoutingOk() (*[]Routing, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *AppService) SetRouting(v []Routing)`

SetRouting sets Routing field to given value.

### HasRouting

`func (o *AppService) HasRouting() bool`

HasRouting returns a boolean if a field has been set.

### SetRoutingNil

`func (o *AppService) SetRoutingNil(b bool)`

 SetRoutingNil sets the value for Routing to be an explicit nil

### UnsetRouting
`func (o *AppService) UnsetRouting()`

UnsetRouting ensures that no value is present for Routing, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


