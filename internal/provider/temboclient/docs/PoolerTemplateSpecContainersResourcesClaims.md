# PoolerTemplateSpecContainersResourcesClaims

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container. | 

## Methods

### NewPoolerTemplateSpecContainersResourcesClaims

`func NewPoolerTemplateSpecContainersResourcesClaims(name string, ) *PoolerTemplateSpecContainersResourcesClaims`

NewPoolerTemplateSpecContainersResourcesClaims instantiates a new PoolerTemplateSpecContainersResourcesClaims object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolerTemplateSpecContainersResourcesClaimsWithDefaults

`func NewPoolerTemplateSpecContainersResourcesClaimsWithDefaults() *PoolerTemplateSpecContainersResourcesClaims`

NewPoolerTemplateSpecContainersResourcesClaimsWithDefaults instantiates a new PoolerTemplateSpecContainersResourcesClaims object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PoolerTemplateSpecContainersResourcesClaims) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PoolerTemplateSpecContainersResourcesClaims) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PoolerTemplateSpecContainersResourcesClaims) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


