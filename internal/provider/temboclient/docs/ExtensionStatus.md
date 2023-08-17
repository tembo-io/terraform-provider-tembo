# ExtensionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Locations** | [**[]ExtensionInstallLocationStatus**](ExtensionInstallLocationStatus.md) |  | 
**Name** | **string** |  | 

## Methods

### NewExtensionStatus

`func NewExtensionStatus(locations []ExtensionInstallLocationStatus, name string, ) *ExtensionStatus`

NewExtensionStatus instantiates a new ExtensionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionStatusWithDefaults

`func NewExtensionStatusWithDefaults() *ExtensionStatus`

NewExtensionStatusWithDefaults instantiates a new ExtensionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ExtensionStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExtensionStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExtensionStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExtensionStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ExtensionStatus) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ExtensionStatus) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLocations

`func (o *ExtensionStatus) GetLocations() []ExtensionInstallLocationStatus`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *ExtensionStatus) GetLocationsOk() (*[]ExtensionInstallLocationStatus, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *ExtensionStatus) SetLocations(v []ExtensionInstallLocationStatus)`

SetLocations sets Locations field to given value.


### GetName

`func (o *ExtensionStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtensionStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtensionStatus) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


