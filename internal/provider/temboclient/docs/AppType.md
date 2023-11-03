# AppType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Restapi** | [**NullableAppConfig**](AppConfig.md) |  | 
**Http** | [**NullableAppConfig**](AppConfig.md) |  | 
**Custom** | [**AppService**](AppService.md) |  | 

## Methods

### NewAppType

`func NewAppType(restapi NullableAppConfig, http NullableAppConfig, custom AppService, ) *AppType`

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


