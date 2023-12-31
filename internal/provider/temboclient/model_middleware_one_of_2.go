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

// checks if the MiddlewareOneOf2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MiddlewareOneOf2{}

// MiddlewareOneOf2 struct for MiddlewareOneOf2
type MiddlewareOneOf2 struct {
	ReplacePathRegex ReplacePathRegexConfig `json:"replacePathRegex"`
}

// NewMiddlewareOneOf2 instantiates a new MiddlewareOneOf2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMiddlewareOneOf2(replacePathRegex ReplacePathRegexConfig) *MiddlewareOneOf2 {
	this := MiddlewareOneOf2{}
	this.ReplacePathRegex = replacePathRegex
	return &this
}

// NewMiddlewareOneOf2WithDefaults instantiates a new MiddlewareOneOf2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMiddlewareOneOf2WithDefaults() *MiddlewareOneOf2 {
	this := MiddlewareOneOf2{}
	return &this
}

// GetReplacePathRegex returns the ReplacePathRegex field value
func (o *MiddlewareOneOf2) GetReplacePathRegex() ReplacePathRegexConfig {
	if o == nil {
		var ret ReplacePathRegexConfig
		return ret
	}

	return o.ReplacePathRegex
}

// GetReplacePathRegexOk returns a tuple with the ReplacePathRegex field value
// and a boolean to check if the value has been set.
func (o *MiddlewareOneOf2) GetReplacePathRegexOk() (*ReplacePathRegexConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReplacePathRegex, true
}

// SetReplacePathRegex sets field value
func (o *MiddlewareOneOf2) SetReplacePathRegex(v ReplacePathRegexConfig) {
	o.ReplacePathRegex = v
}

func (o MiddlewareOneOf2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MiddlewareOneOf2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["replacePathRegex"] = o.ReplacePathRegex
	return toSerialize, nil
}

type NullableMiddlewareOneOf2 struct {
	value *MiddlewareOneOf2
	isSet bool
}

func (v NullableMiddlewareOneOf2) Get() *MiddlewareOneOf2 {
	return v.value
}

func (v *NullableMiddlewareOneOf2) Set(val *MiddlewareOneOf2) {
	v.value = val
	v.isSet = true
}

func (v NullableMiddlewareOneOf2) IsSet() bool {
	return v.isSet
}

func (v *NullableMiddlewareOneOf2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMiddlewareOneOf2(val *MiddlewareOneOf2) *NullableMiddlewareOneOf2 {
	return &NullableMiddlewareOneOf2{value: val, isSet: true}
}

func (v NullableMiddlewareOneOf2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMiddlewareOneOf2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


