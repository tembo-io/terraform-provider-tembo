# Ingress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**Path** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIngress

`func NewIngress(enabled bool, ) *Ingress`

NewIngress instantiates a new Ingress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngressWithDefaults

`func NewIngressWithDefaults() *Ingress`

NewIngressWithDefaults instantiates a new Ingress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *Ingress) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Ingress) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Ingress) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetPath

`func (o *Ingress) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Ingress) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Ingress) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Ingress) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *Ingress) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *Ingress) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


