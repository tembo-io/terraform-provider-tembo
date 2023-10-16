# TrunkInstallStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **bool** |  | 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**InstalledToPods** | Pointer to **[]string** |  | [optional] 
**Name** | **string** |  | 
**Version** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTrunkInstallStatus

`func NewTrunkInstallStatus(error_ bool, name string, ) *TrunkInstallStatus`

NewTrunkInstallStatus instantiates a new TrunkInstallStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrunkInstallStatusWithDefaults

`func NewTrunkInstallStatusWithDefaults() *TrunkInstallStatus`

NewTrunkInstallStatusWithDefaults instantiates a new TrunkInstallStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *TrunkInstallStatus) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TrunkInstallStatus) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TrunkInstallStatus) SetError(v bool)`

SetError sets Error field to given value.


### GetErrorMessage

`func (o *TrunkInstallStatus) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TrunkInstallStatus) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TrunkInstallStatus) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *TrunkInstallStatus) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *TrunkInstallStatus) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *TrunkInstallStatus) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetInstalledToPods

`func (o *TrunkInstallStatus) GetInstalledToPods() []string`

GetInstalledToPods returns the InstalledToPods field if non-nil, zero value otherwise.

### GetInstalledToPodsOk

`func (o *TrunkInstallStatus) GetInstalledToPodsOk() (*[]string, bool)`

GetInstalledToPodsOk returns a tuple with the InstalledToPods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledToPods

`func (o *TrunkInstallStatus) SetInstalledToPods(v []string)`

SetInstalledToPods sets InstalledToPods field to given value.

### HasInstalledToPods

`func (o *TrunkInstallStatus) HasInstalledToPods() bool`

HasInstalledToPods returns a boolean if a field has been set.

### SetInstalledToPodsNil

`func (o *TrunkInstallStatus) SetInstalledToPodsNil(b bool)`

 SetInstalledToPodsNil sets the value for InstalledToPods to be an explicit nil

### UnsetInstalledToPods
`func (o *TrunkInstallStatus) UnsetInstalledToPods()`

UnsetInstalledToPods ensures that no value is present for InstalledToPods, not even an explicit nil
### GetName

`func (o *TrunkInstallStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrunkInstallStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrunkInstallStatus) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *TrunkInstallStatus) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TrunkInstallStatus) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TrunkInstallStatus) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TrunkInstallStatus) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *TrunkInstallStatus) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *TrunkInstallStatus) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


