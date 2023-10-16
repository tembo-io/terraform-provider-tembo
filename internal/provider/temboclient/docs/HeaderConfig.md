# HeaderConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | **map[string]string** |  | 
**Name** | **string** |  | 

## Methods

### NewHeaderConfig

`func NewHeaderConfig(config map[string]string, name string, ) *HeaderConfig`

NewHeaderConfig instantiates a new HeaderConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeaderConfigWithDefaults

`func NewHeaderConfigWithDefaults() *HeaderConfig`

NewHeaderConfigWithDefaults instantiates a new HeaderConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *HeaderConfig) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *HeaderConfig) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *HeaderConfig) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.


### GetName

`func (o *HeaderConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HeaderConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HeaderConfig) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


