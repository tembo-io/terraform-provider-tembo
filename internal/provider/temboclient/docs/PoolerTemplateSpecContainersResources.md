# PoolerTemplateSpecContainersResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Claims** | Pointer to [**[]PoolerTemplateSpecContainersResourcesClaims**](PoolerTemplateSpecContainersResourcesClaims.md) |  | [optional] 
**Limits** | Pointer to [**map[string]IntOrString**](IntOrString.md) |  | [optional] 
**Requests** | Pointer to [**map[string]IntOrString**](IntOrString.md) |  | [optional] 

## Methods

### NewPoolerTemplateSpecContainersResources

`func NewPoolerTemplateSpecContainersResources() *PoolerTemplateSpecContainersResources`

NewPoolerTemplateSpecContainersResources instantiates a new PoolerTemplateSpecContainersResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolerTemplateSpecContainersResourcesWithDefaults

`func NewPoolerTemplateSpecContainersResourcesWithDefaults() *PoolerTemplateSpecContainersResources`

NewPoolerTemplateSpecContainersResourcesWithDefaults instantiates a new PoolerTemplateSpecContainersResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaims

`func (o *PoolerTemplateSpecContainersResources) GetClaims() []PoolerTemplateSpecContainersResourcesClaims`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *PoolerTemplateSpecContainersResources) GetClaimsOk() (*[]PoolerTemplateSpecContainersResourcesClaims, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *PoolerTemplateSpecContainersResources) SetClaims(v []PoolerTemplateSpecContainersResourcesClaims)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *PoolerTemplateSpecContainersResources) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### SetClaimsNil

`func (o *PoolerTemplateSpecContainersResources) SetClaimsNil(b bool)`

 SetClaimsNil sets the value for Claims to be an explicit nil

### UnsetClaims
`func (o *PoolerTemplateSpecContainersResources) UnsetClaims()`

UnsetClaims ensures that no value is present for Claims, not even an explicit nil
### GetLimits

`func (o *PoolerTemplateSpecContainersResources) GetLimits() map[string]IntOrString`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *PoolerTemplateSpecContainersResources) GetLimitsOk() (*map[string]IntOrString, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *PoolerTemplateSpecContainersResources) SetLimits(v map[string]IntOrString)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *PoolerTemplateSpecContainersResources) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### SetLimitsNil

`func (o *PoolerTemplateSpecContainersResources) SetLimitsNil(b bool)`

 SetLimitsNil sets the value for Limits to be an explicit nil

### UnsetLimits
`func (o *PoolerTemplateSpecContainersResources) UnsetLimits()`

UnsetLimits ensures that no value is present for Limits, not even an explicit nil
### GetRequests

`func (o *PoolerTemplateSpecContainersResources) GetRequests() map[string]IntOrString`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *PoolerTemplateSpecContainersResources) GetRequestsOk() (*map[string]IntOrString, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *PoolerTemplateSpecContainersResources) SetRequests(v map[string]IntOrString)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *PoolerTemplateSpecContainersResources) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### SetRequestsNil

`func (o *PoolerTemplateSpecContainersResources) SetRequestsNil(b bool)`

 SetRequestsNil sets the value for Requests to be an explicit nil

### UnsetRequests
`func (o *PoolerTemplateSpecContainersResources) UnsetRequests()`

UnsetRequests ensures that no value is present for Requests, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


