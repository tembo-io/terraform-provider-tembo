# ExtensionInstallLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Database** | Pointer to **string** |  | [optional] 
**Enabled** | **bool** |  | 
**Schema** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewExtensionInstallLocation

`func NewExtensionInstallLocation(enabled bool, ) *ExtensionInstallLocation`

NewExtensionInstallLocation instantiates a new ExtensionInstallLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionInstallLocationWithDefaults

`func NewExtensionInstallLocationWithDefaults() *ExtensionInstallLocation`

NewExtensionInstallLocationWithDefaults instantiates a new ExtensionInstallLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabase

`func (o *ExtensionInstallLocation) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *ExtensionInstallLocation) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *ExtensionInstallLocation) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *ExtensionInstallLocation) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetEnabled

`func (o *ExtensionInstallLocation) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ExtensionInstallLocation) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ExtensionInstallLocation) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSchema

`func (o *ExtensionInstallLocation) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *ExtensionInstallLocation) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *ExtensionInstallLocation) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *ExtensionInstallLocation) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *ExtensionInstallLocation) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *ExtensionInstallLocation) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetVersion

`func (o *ExtensionInstallLocation) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExtensionInstallLocation) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExtensionInstallLocation) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExtensionInstallLocation) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ExtensionInstallLocation) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ExtensionInstallLocation) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


