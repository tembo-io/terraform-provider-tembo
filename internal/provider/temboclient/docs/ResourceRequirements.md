# ResourceRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limits** | Pointer to [**NullableResource**](Resource.md) |  | [optional] 
**Requests** | Pointer to [**NullableResource**](Resource.md) |  | [optional] 

## Methods

### NewResourceRequirements

`func NewResourceRequirements() *ResourceRequirements`

NewResourceRequirements instantiates a new ResourceRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceRequirementsWithDefaults

`func NewResourceRequirementsWithDefaults() *ResourceRequirements`

NewResourceRequirementsWithDefaults instantiates a new ResourceRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimits

`func (o *ResourceRequirements) GetLimits() Resource`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *ResourceRequirements) GetLimitsOk() (*Resource, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *ResourceRequirements) SetLimits(v Resource)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *ResourceRequirements) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### SetLimitsNil

`func (o *ResourceRequirements) SetLimitsNil(b bool)`

 SetLimitsNil sets the value for Limits to be an explicit nil

### UnsetLimits
`func (o *ResourceRequirements) UnsetLimits()`

UnsetLimits ensures that no value is present for Limits, not even an explicit nil
### GetRequests

`func (o *ResourceRequirements) GetRequests() Resource`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *ResourceRequirements) GetRequestsOk() (*Resource, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *ResourceRequirements) SetRequests(v Resource)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *ResourceRequirements) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### SetRequestsNil

`func (o *ResourceRequirements) SetRequestsNil(b bool)`

 SetRequestsNil sets the value for Requests to be an explicit nil

### UnsetRequests
`func (o *ResourceRequirements) UnsetRequests()`

UnsetRequests ensures that no value is present for Requests, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


