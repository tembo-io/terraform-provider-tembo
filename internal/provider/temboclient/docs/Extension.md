# Extension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | A description of the extension. (Optional)  **Default**: \&quot;No description provided\&quot; | [optional] 
**Locations** | [**[]ExtensionInstallLocation**](ExtensionInstallLocation.md) | A list of locations (databases) to enabled the extension on. | 
**Name** | **string** | The name of the extension to enable. | 

## Methods

### NewExtension

`func NewExtension(locations []ExtensionInstallLocation, name string, ) *Extension`

NewExtension instantiates a new Extension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionWithDefaults

`func NewExtensionWithDefaults() *Extension`

NewExtensionWithDefaults instantiates a new Extension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Extension) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Extension) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Extension) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Extension) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Extension) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Extension) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLocations

`func (o *Extension) GetLocations() []ExtensionInstallLocation`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *Extension) GetLocationsOk() (*[]ExtensionInstallLocation, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *Extension) SetLocations(v []ExtensionInstallLocation)`

SetLocations sets Locations field to given value.


### GetName

`func (o *Extension) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Extension) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Extension) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


