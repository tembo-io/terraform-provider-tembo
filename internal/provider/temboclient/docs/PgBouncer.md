# PgBouncer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | Pointer to **map[string]string** |  | [optional] 
**PoolMode** | Pointer to [**PoolerPgbouncerPoolMode**](PoolerPgbouncerPoolMode.md) |  | [optional] 
**Resources** | Pointer to [**NullablePoolerTemplateSpecContainersResources**](PoolerTemplateSpecContainersResources.md) |  | [optional] 

## Methods

### NewPgBouncer

`func NewPgBouncer() *PgBouncer`

NewPgBouncer instantiates a new PgBouncer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPgBouncerWithDefaults

`func NewPgBouncerWithDefaults() *PgBouncer`

NewPgBouncerWithDefaults instantiates a new PgBouncer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameters

`func (o *PgBouncer) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *PgBouncer) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *PgBouncer) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *PgBouncer) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *PgBouncer) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *PgBouncer) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetPoolMode

`func (o *PgBouncer) GetPoolMode() PoolerPgbouncerPoolMode`

GetPoolMode returns the PoolMode field if non-nil, zero value otherwise.

### GetPoolModeOk

`func (o *PgBouncer) GetPoolModeOk() (*PoolerPgbouncerPoolMode, bool)`

GetPoolModeOk returns a tuple with the PoolMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolMode

`func (o *PgBouncer) SetPoolMode(v PoolerPgbouncerPoolMode)`

SetPoolMode sets PoolMode field to given value.

### HasPoolMode

`func (o *PgBouncer) HasPoolMode() bool`

HasPoolMode returns a boolean if a field has been set.

### GetResources

`func (o *PgBouncer) GetResources() PoolerTemplateSpecContainersResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *PgBouncer) GetResourcesOk() (*PoolerTemplateSpecContainersResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *PgBouncer) SetResources(v PoolerTemplateSpecContainersResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *PgBouncer) HasResources() bool`

HasResources returns a boolean if a field has been set.

### SetResourcesNil

`func (o *PgBouncer) SetResourcesNil(b bool)`

 SetResourcesNil sets the value for Resources to be an explicit nil

### UnsetResources
`func (o *PgBouncer) UnsetResources()`

UnsetResources ensures that no value is present for Resources, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


