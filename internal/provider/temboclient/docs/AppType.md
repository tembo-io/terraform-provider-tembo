# AppType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Restapi** | [**NullableAppConfig**](AppConfig.md) |  | 
**Http** | [**NullableAppConfig**](AppConfig.md) |  | 
**MqApi** | [**NullableAppConfig**](AppConfig.md) |  | 
**Embeddings** | [**NullableAppConfig**](AppConfig.md) |  | 
**Pganalyze** | [**NullableAppConfig**](AppConfig.md) |  | 
**Sqlrunner** | [**NullableAppConfig**](AppConfig.md) |  | 
**Custom** | [**AppService**](AppService.md) |  | 

## Methods

### NewAppType

`func NewAppType(restapi NullableAppConfig, http NullableAppConfig, mqApi NullableAppConfig, embeddings NullableAppConfig, pganalyze NullableAppConfig, sqlrunner NullableAppConfig, custom AppService, ) *AppType`

NewAppType instantiates a new AppType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppTypeWithDefaults

`func NewAppTypeWithDefaults() *AppType`

NewAppTypeWithDefaults instantiates a new AppType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestapi

`func (o *AppType) GetRestapi() AppConfig`

GetRestapi returns the Restapi field if non-nil, zero value otherwise.

### GetRestapiOk

`func (o *AppType) GetRestapiOk() (*AppConfig, bool)`

GetRestapiOk returns a tuple with the Restapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestapi

`func (o *AppType) SetRestapi(v AppConfig)`

SetRestapi sets Restapi field to given value.


### SetRestapiNil

`func (o *AppType) SetRestapiNil(b bool)`

 SetRestapiNil sets the value for Restapi to be an explicit nil

### UnsetRestapi
`func (o *AppType) UnsetRestapi()`

UnsetRestapi ensures that no value is present for Restapi, not even an explicit nil
### GetHttp

`func (o *AppType) GetHttp() AppConfig`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *AppType) GetHttpOk() (*AppConfig, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *AppType) SetHttp(v AppConfig)`

SetHttp sets Http field to given value.


### SetHttpNil

`func (o *AppType) SetHttpNil(b bool)`

 SetHttpNil sets the value for Http to be an explicit nil

### UnsetHttp
`func (o *AppType) UnsetHttp()`

UnsetHttp ensures that no value is present for Http, not even an explicit nil
### GetMqApi

`func (o *AppType) GetMqApi() AppConfig`

GetMqApi returns the MqApi field if non-nil, zero value otherwise.

### GetMqApiOk

`func (o *AppType) GetMqApiOk() (*AppConfig, bool)`

GetMqApiOk returns a tuple with the MqApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqApi

`func (o *AppType) SetMqApi(v AppConfig)`

SetMqApi sets MqApi field to given value.


### SetMqApiNil

`func (o *AppType) SetMqApiNil(b bool)`

 SetMqApiNil sets the value for MqApi to be an explicit nil

### UnsetMqApi
`func (o *AppType) UnsetMqApi()`

UnsetMqApi ensures that no value is present for MqApi, not even an explicit nil
### GetEmbeddings

`func (o *AppType) GetEmbeddings() AppConfig`

GetEmbeddings returns the Embeddings field if non-nil, zero value otherwise.

### GetEmbeddingsOk

`func (o *AppType) GetEmbeddingsOk() (*AppConfig, bool)`

GetEmbeddingsOk returns a tuple with the Embeddings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddings

`func (o *AppType) SetEmbeddings(v AppConfig)`

SetEmbeddings sets Embeddings field to given value.


### SetEmbeddingsNil

`func (o *AppType) SetEmbeddingsNil(b bool)`

 SetEmbeddingsNil sets the value for Embeddings to be an explicit nil

### UnsetEmbeddings
`func (o *AppType) UnsetEmbeddings()`

UnsetEmbeddings ensures that no value is present for Embeddings, not even an explicit nil
### GetPganalyze

`func (o *AppType) GetPganalyze() AppConfig`

GetPganalyze returns the Pganalyze field if non-nil, zero value otherwise.

### GetPganalyzeOk

`func (o *AppType) GetPganalyzeOk() (*AppConfig, bool)`

GetPganalyzeOk returns a tuple with the Pganalyze field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPganalyze

`func (o *AppType) SetPganalyze(v AppConfig)`

SetPganalyze sets Pganalyze field to given value.


### SetPganalyzeNil

`func (o *AppType) SetPganalyzeNil(b bool)`

 SetPganalyzeNil sets the value for Pganalyze to be an explicit nil

### UnsetPganalyze
`func (o *AppType) UnsetPganalyze()`

UnsetPganalyze ensures that no value is present for Pganalyze, not even an explicit nil
### GetSqlrunner

`func (o *AppType) GetSqlrunner() AppConfig`

GetSqlrunner returns the Sqlrunner field if non-nil, zero value otherwise.

### GetSqlrunnerOk

`func (o *AppType) GetSqlrunnerOk() (*AppConfig, bool)`

GetSqlrunnerOk returns a tuple with the Sqlrunner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlrunner

`func (o *AppType) SetSqlrunner(v AppConfig)`

SetSqlrunner sets Sqlrunner field to given value.


### SetSqlrunnerNil

`func (o *AppType) SetSqlrunnerNil(b bool)`

 SetSqlrunnerNil sets the value for Sqlrunner to be an explicit nil

### UnsetSqlrunner
`func (o *AppType) UnsetSqlrunner()`

UnsetSqlrunner ensures that no value is present for Sqlrunner, not even an explicit nil
### GetCustom

`func (o *AppType) GetCustom() AppService`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *AppType) GetCustomOk() (*AppService, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *AppType) SetCustom(v AppService)`

SetCustom sets Custom field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


