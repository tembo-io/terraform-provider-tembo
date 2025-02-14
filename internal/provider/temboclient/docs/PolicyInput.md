# PolicyInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | A valid Action ID. Available Action IDs include &#39;CreateInstance&#39; and &#39;ManagePermissions&#39;. Find all available Actions on the Actions API. | 
**Allowed** | **bool** | Whether the Action is allowed or not for the Role | 
**Role** | **string** | A valid Role ID. Available Role IDs include &#39;admin&#39; and &#39;basic_member&#39;. | 

## Methods

### NewPolicyInput

`func NewPolicyInput(action string, allowed bool, role string, ) *PolicyInput`

NewPolicyInput instantiates a new PolicyInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyInputWithDefaults

`func NewPolicyInputWithDefaults() *PolicyInput`

NewPolicyInputWithDefaults instantiates a new PolicyInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PolicyInput) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PolicyInput) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PolicyInput) SetAction(v string)`

SetAction sets Action field to given value.


### GetAllowed

`func (o *PolicyInput) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *PolicyInput) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *PolicyInput) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.


### GetRole

`func (o *PolicyInput) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PolicyInput) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PolicyInput) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


