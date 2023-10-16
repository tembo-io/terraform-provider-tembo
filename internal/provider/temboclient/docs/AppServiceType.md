# AppServiceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Restapi** | [**NullableAppServiceConfig**](AppServiceConfig.md) |  | 

## Methods

### NewAppServiceType

`func NewAppServiceType(restapi NullableAppServiceConfig, ) *AppServiceType`

NewAppServiceType instantiates a new AppServiceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceTypeWithDefaults

`func NewAppServiceTypeWithDefaults() *AppServiceType`

NewAppServiceTypeWithDefaults instantiates a new AppServiceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestapi

`func (o *AppServiceType) GetRestapi() AppServiceConfig`

GetRestapi returns the Restapi field if non-nil, zero value otherwise.

### GetRestapiOk

`func (o *AppServiceType) GetRestapiOk() (*AppServiceConfig, bool)`

GetRestapiOk returns a tuple with the Restapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestapi

`func (o *AppServiceType) SetRestapi(v AppServiceConfig)`

SetRestapi sets Restapi field to given value.


### SetRestapiNil

`func (o *AppServiceType) SetRestapiNil(b bool)`

 SetRestapiNil sets the value for Restapi to be an explicit nil

### UnsetRestapi
`func (o *AppServiceType) UnsetRestapi()`

UnsetRestapi ensures that no value is present for Restapi, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


