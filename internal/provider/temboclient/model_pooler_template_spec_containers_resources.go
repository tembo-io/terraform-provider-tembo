/*
Tembo Cloud

Platform API for Tembo Cloud             </br>             </br>             To find a Tembo Data API, please find it here:             </br>             </br>             [AWS US East 1](https://api.data-1.use1.tembo.io/swagger-ui/)             

API version: v1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package temboclient

import (
	"encoding/json"
)

// checks if the PoolerTemplateSpecContainersResources type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PoolerTemplateSpecContainersResources{}

// PoolerTemplateSpecContainersResources Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
type PoolerTemplateSpecContainersResources struct {
	// Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.
	Claims []PoolerTemplateSpecContainersResourcesClaims `json:"claims,omitempty"`
	// Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Limits map[string]IntOrString `json:"limits,omitempty"`
	// Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Requests map[string]IntOrString `json:"requests,omitempty"`
}

// NewPoolerTemplateSpecContainersResources instantiates a new PoolerTemplateSpecContainersResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolerTemplateSpecContainersResources() *PoolerTemplateSpecContainersResources {
	this := PoolerTemplateSpecContainersResources{}
	return &this
}

// NewPoolerTemplateSpecContainersResourcesWithDefaults instantiates a new PoolerTemplateSpecContainersResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolerTemplateSpecContainersResourcesWithDefaults() *PoolerTemplateSpecContainersResources {
	this := PoolerTemplateSpecContainersResources{}
	return &this
}

// GetClaims returns the Claims field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolerTemplateSpecContainersResources) GetClaims() []PoolerTemplateSpecContainersResourcesClaims {
	if o == nil {
		var ret []PoolerTemplateSpecContainersResourcesClaims
		return ret
	}
	return o.Claims
}

// GetClaimsOk returns a tuple with the Claims field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolerTemplateSpecContainersResources) GetClaimsOk() ([]PoolerTemplateSpecContainersResourcesClaims, bool) {
	if o == nil || IsNil(o.Claims) {
		return nil, false
	}
	return o.Claims, true
}

// HasClaims returns a boolean if a field has been set.
func (o *PoolerTemplateSpecContainersResources) HasClaims() bool {
	if o != nil && !IsNil(o.Claims) {
		return true
	}

	return false
}

// SetClaims gets a reference to the given []PoolerTemplateSpecContainersResourcesClaims and assigns it to the Claims field.
func (o *PoolerTemplateSpecContainersResources) SetClaims(v []PoolerTemplateSpecContainersResourcesClaims) {
	o.Claims = v
}

// GetLimits returns the Limits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolerTemplateSpecContainersResources) GetLimits() map[string]IntOrString {
	if o == nil {
		var ret map[string]IntOrString
		return ret
	}
	return o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolerTemplateSpecContainersResources) GetLimitsOk() (*map[string]IntOrString, bool) {
	if o == nil || IsNil(o.Limits) {
		return nil, false
	}
	return &o.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *PoolerTemplateSpecContainersResources) HasLimits() bool {
	if o != nil && !IsNil(o.Limits) {
		return true
	}

	return false
}

// SetLimits gets a reference to the given map[string]IntOrString and assigns it to the Limits field.
func (o *PoolerTemplateSpecContainersResources) SetLimits(v map[string]IntOrString) {
	o.Limits = v
}

// GetRequests returns the Requests field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolerTemplateSpecContainersResources) GetRequests() map[string]IntOrString {
	if o == nil {
		var ret map[string]IntOrString
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolerTemplateSpecContainersResources) GetRequestsOk() (*map[string]IntOrString, bool) {
	if o == nil || IsNil(o.Requests) {
		return nil, false
	}
	return &o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *PoolerTemplateSpecContainersResources) HasRequests() bool {
	if o != nil && !IsNil(o.Requests) {
		return true
	}

	return false
}

// SetRequests gets a reference to the given map[string]IntOrString and assigns it to the Requests field.
func (o *PoolerTemplateSpecContainersResources) SetRequests(v map[string]IntOrString) {
	o.Requests = v
}

func (o PoolerTemplateSpecContainersResources) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoolerTemplateSpecContainersResources) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Claims != nil {
		toSerialize["claims"] = o.Claims
	}
	if o.Limits != nil {
		toSerialize["limits"] = o.Limits
	}
	if o.Requests != nil {
		toSerialize["requests"] = o.Requests
	}
	return toSerialize, nil
}

type NullablePoolerTemplateSpecContainersResources struct {
	value *PoolerTemplateSpecContainersResources
	isSet bool
}

func (v NullablePoolerTemplateSpecContainersResources) Get() *PoolerTemplateSpecContainersResources {
	return v.value
}

func (v *NullablePoolerTemplateSpecContainersResources) Set(val *PoolerTemplateSpecContainersResources) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolerTemplateSpecContainersResources) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolerTemplateSpecContainersResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolerTemplateSpecContainersResources(val *PoolerTemplateSpecContainersResources) *NullablePoolerTemplateSpecContainersResources {
	return &NullablePoolerTemplateSpecContainersResources{value: val, isSet: true}
}

func (v NullablePoolerTemplateSpecContainersResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolerTemplateSpecContainersResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


