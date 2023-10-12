# AvailableSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of an available secret | 
**PossibleKeys** | **[]string** | For this secret, available keys | 

## Methods

### NewAvailableSecret

`func NewAvailableSecret(name string, possibleKeys []string, ) *AvailableSecret`

NewAvailableSecret instantiates a new AvailableSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableSecretWithDefaults

`func NewAvailableSecretWithDefaults() *AvailableSecret`

NewAvailableSecretWithDefaults instantiates a new AvailableSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AvailableSecret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AvailableSecret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AvailableSecret) SetName(v string)`

SetName sets Name field to given value.


### GetPossibleKeys

`func (o *AvailableSecret) GetPossibleKeys() []string`

GetPossibleKeys returns the PossibleKeys field if non-nil, zero value otherwise.

### GetPossibleKeysOk

`func (o *AvailableSecret) GetPossibleKeysOk() (*[]string, bool)`

GetPossibleKeysOk returns a tuple with the PossibleKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleKeys

`func (o *AvailableSecret) SetPossibleKeys(v []string)`

SetPossibleKeys sets PossibleKeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


