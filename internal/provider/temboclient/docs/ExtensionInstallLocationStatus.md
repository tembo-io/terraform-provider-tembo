# ExtensionInstallLocationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Database** | **string** |  | 
**Enabled** | Pointer to **NullableBool** |  | [optional] 
**Error** | Pointer to **NullableBool** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**Schema** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewExtensionInstallLocationStatus

`func NewExtensionInstallLocationStatus(database string, ) *ExtensionInstallLocationStatus`

NewExtensionInstallLocationStatus instantiates a new ExtensionInstallLocationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionInstallLocationStatusWithDefaults

`func NewExtensionInstallLocationStatusWithDefaults() *ExtensionInstallLocationStatus`

NewExtensionInstallLocationStatusWithDefaults instantiates a new ExtensionInstallLocationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabase

`func (o *ExtensionInstallLocationStatus) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *ExtensionInstallLocationStatus) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *ExtensionInstallLocationStatus) SetDatabase(v string)`

SetDatabase sets Database field to given value.


### GetEnabled

`func (o *ExtensionInstallLocationStatus) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ExtensionInstallLocationStatus) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ExtensionInstallLocationStatus) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ExtensionInstallLocationStatus) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *ExtensionInstallLocationStatus) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *ExtensionInstallLocationStatus) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetError

`func (o *ExtensionInstallLocationStatus) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ExtensionInstallLocationStatus) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ExtensionInstallLocationStatus) SetError(v bool)`

SetError sets Error field to given value.

### HasError

`func (o *ExtensionInstallLocationStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ExtensionInstallLocationStatus) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ExtensionInstallLocationStatus) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetErrorMessage

`func (o *ExtensionInstallLocationStatus) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ExtensionInstallLocationStatus) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ExtensionInstallLocationStatus) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ExtensionInstallLocationStatus) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ExtensionInstallLocationStatus) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ExtensionInstallLocationStatus) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetSchema

`func (o *ExtensionInstallLocationStatus) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *ExtensionInstallLocationStatus) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *ExtensionInstallLocationStatus) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *ExtensionInstallLocationStatus) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *ExtensionInstallLocationStatus) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *ExtensionInstallLocationStatus) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetVersion

`func (o *ExtensionInstallLocationStatus) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExtensionInstallLocationStatus) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExtensionInstallLocationStatus) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExtensionInstallLocationStatus) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ExtensionInstallLocationStatus) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ExtensionInstallLocationStatus) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


