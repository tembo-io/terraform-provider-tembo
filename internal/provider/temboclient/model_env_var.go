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

// checks if the EnvVar type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvVar{}

// EnvVar struct for EnvVar
type EnvVar struct {
	Name string `json:"name"`
	Value NullableString `json:"value,omitempty"`
	ValueFromPlatform NullableEnvVarRef `json:"valueFromPlatform,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnvVar EnvVar

// NewEnvVar instantiates a new EnvVar object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvVar(name string) *EnvVar {
	this := EnvVar{}
	this.Name = name
	return &this
}

// NewEnvVarWithDefaults instantiates a new EnvVar object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvVarWithDefaults() *EnvVar {
	this := EnvVar{}
	return &this
}

// GetName returns the Name field value
func (o *EnvVar) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EnvVar) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EnvVar) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvVar) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvVar) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *EnvVar) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *EnvVar) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *EnvVar) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *EnvVar) UnsetValue() {
	o.Value.Unset()
}

// GetValueFromPlatform returns the ValueFromPlatform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvVar) GetValueFromPlatform() EnvVarRef {
	if o == nil || IsNil(o.ValueFromPlatform.Get()) {
		var ret EnvVarRef
		return ret
	}
	return *o.ValueFromPlatform.Get()
}

// GetValueFromPlatformOk returns a tuple with the ValueFromPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvVar) GetValueFromPlatformOk() (*EnvVarRef, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValueFromPlatform.Get(), o.ValueFromPlatform.IsSet()
}

// HasValueFromPlatform returns a boolean if a field has been set.
func (o *EnvVar) HasValueFromPlatform() bool {
	if o != nil && o.ValueFromPlatform.IsSet() {
		return true
	}

	return false
}

// SetValueFromPlatform gets a reference to the given NullableEnvVarRef and assigns it to the ValueFromPlatform field.
func (o *EnvVar) SetValueFromPlatform(v EnvVarRef) {
	o.ValueFromPlatform.Set(&v)
}
// SetValueFromPlatformNil sets the value for ValueFromPlatform to be an explicit nil
func (o *EnvVar) SetValueFromPlatformNil() {
	o.ValueFromPlatform.Set(nil)
}

// UnsetValueFromPlatform ensures that no value is present for ValueFromPlatform, not even an explicit nil
func (o *EnvVar) UnsetValueFromPlatform() {
	o.ValueFromPlatform.Unset()
}

func (o EnvVar) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvVar) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if o.ValueFromPlatform.IsSet() {
		toSerialize["valueFromPlatform"] = o.ValueFromPlatform.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EnvVar) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varEnvVar := _EnvVar{}

	err = json.Unmarshal(data, &varEnvVar)

	if err != nil {
		return err
	}

	*o = EnvVar(varEnvVar)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "value")
		delete(additionalProperties, "valueFromPlatform")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnvVar struct {
	value *EnvVar
	isSet bool
}

func (v NullableEnvVar) Get() *EnvVar {
	return v.value
}

func (v *NullableEnvVar) Set(val *EnvVar) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvVar) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvVar) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvVar(val *EnvVar) *NullableEnvVar {
	return &NullableEnvVar{value: val, isSet: true}
}

func (v NullableEnvVar) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvVar) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


