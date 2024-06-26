/*
Tembo Cloud

Platform API for Tembo Cloud             </br>             </br>             To find a Tembo Data API, please find it here:             </br>             </br>             [AWS US East 1](https://api.data-1.use1.tembo.io/swagger-ui/)             

API version: v1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package temboclient

import (
	"encoding/json"
	"fmt"
)

// checks if the MiddlewareOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MiddlewareOneOf{}

// MiddlewareOneOf struct for MiddlewareOneOf
type MiddlewareOneOf struct {
	CustomRequestHeaders HeaderConfig `json:"customRequestHeaders"`
	AdditionalProperties map[string]interface{}
}

type _MiddlewareOneOf MiddlewareOneOf

// NewMiddlewareOneOf instantiates a new MiddlewareOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMiddlewareOneOf(customRequestHeaders HeaderConfig) *MiddlewareOneOf {
	this := MiddlewareOneOf{}
	this.CustomRequestHeaders = customRequestHeaders
	return &this
}

// NewMiddlewareOneOfWithDefaults instantiates a new MiddlewareOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMiddlewareOneOfWithDefaults() *MiddlewareOneOf {
	this := MiddlewareOneOf{}
	return &this
}

// GetCustomRequestHeaders returns the CustomRequestHeaders field value
func (o *MiddlewareOneOf) GetCustomRequestHeaders() HeaderConfig {
	if o == nil {
		var ret HeaderConfig
		return ret
	}

	return o.CustomRequestHeaders
}

// GetCustomRequestHeadersOk returns a tuple with the CustomRequestHeaders field value
// and a boolean to check if the value has been set.
func (o *MiddlewareOneOf) GetCustomRequestHeadersOk() (*HeaderConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomRequestHeaders, true
}

// SetCustomRequestHeaders sets field value
func (o *MiddlewareOneOf) SetCustomRequestHeaders(v HeaderConfig) {
	o.CustomRequestHeaders = v
}

func (o MiddlewareOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MiddlewareOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customRequestHeaders"] = o.CustomRequestHeaders

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MiddlewareOneOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"customRequestHeaders",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMiddlewareOneOf := _MiddlewareOneOf{}

	err = json.Unmarshal(data, &varMiddlewareOneOf)

	if err != nil {
		return err
	}

	*o = MiddlewareOneOf(varMiddlewareOneOf)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "customRequestHeaders")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMiddlewareOneOf struct {
	value *MiddlewareOneOf
	isSet bool
}

func (v NullableMiddlewareOneOf) Get() *MiddlewareOneOf {
	return v.value
}

func (v *NullableMiddlewareOneOf) Set(val *MiddlewareOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMiddlewareOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMiddlewareOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMiddlewareOneOf(val *MiddlewareOneOf) *NullableMiddlewareOneOf {
	return &NullableMiddlewareOneOf{value: val, isSet: true}
}

func (v NullableMiddlewareOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMiddlewareOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


