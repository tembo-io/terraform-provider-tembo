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

// checks if the PoolerTemplateSpecContainersResourcesClaims type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PoolerTemplateSpecContainersResourcesClaims{}

// PoolerTemplateSpecContainersResourcesClaims struct for PoolerTemplateSpecContainersResourcesClaims
type PoolerTemplateSpecContainersResourcesClaims struct {
	Name string `json:"name"`
}

// NewPoolerTemplateSpecContainersResourcesClaims instantiates a new PoolerTemplateSpecContainersResourcesClaims object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolerTemplateSpecContainersResourcesClaims(name string) *PoolerTemplateSpecContainersResourcesClaims {
	this := PoolerTemplateSpecContainersResourcesClaims{}
	this.Name = name
	return &this
}

// NewPoolerTemplateSpecContainersResourcesClaimsWithDefaults instantiates a new PoolerTemplateSpecContainersResourcesClaims object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolerTemplateSpecContainersResourcesClaimsWithDefaults() *PoolerTemplateSpecContainersResourcesClaims {
	this := PoolerTemplateSpecContainersResourcesClaims{}
	return &this
}

// GetName returns the Name field value
func (o *PoolerTemplateSpecContainersResourcesClaims) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PoolerTemplateSpecContainersResourcesClaims) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PoolerTemplateSpecContainersResourcesClaims) SetName(v string) {
	o.Name = v
}

func (o PoolerTemplateSpecContainersResourcesClaims) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoolerTemplateSpecContainersResourcesClaims) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullablePoolerTemplateSpecContainersResourcesClaims struct {
	value *PoolerTemplateSpecContainersResourcesClaims
	isSet bool
}

func (v NullablePoolerTemplateSpecContainersResourcesClaims) Get() *PoolerTemplateSpecContainersResourcesClaims {
	return v.value
}

func (v *NullablePoolerTemplateSpecContainersResourcesClaims) Set(val *PoolerTemplateSpecContainersResourcesClaims) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolerTemplateSpecContainersResourcesClaims) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolerTemplateSpecContainersResourcesClaims) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolerTemplateSpecContainersResourcesClaims(val *PoolerTemplateSpecContainersResourcesClaims) *NullablePoolerTemplateSpecContainersResourcesClaims {
	return &NullablePoolerTemplateSpecContainersResourcesClaims{value: val, isSet: true}
}

func (v NullablePoolerTemplateSpecContainersResourcesClaims) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolerTemplateSpecContainersResourcesClaims) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


