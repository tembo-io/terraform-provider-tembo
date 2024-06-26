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

// checks if the HeaderConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HeaderConfig{}

// HeaderConfig struct for HeaderConfig
type HeaderConfig struct {
	Config map[string]string `json:"config"`
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _HeaderConfig HeaderConfig

// NewHeaderConfig instantiates a new HeaderConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeaderConfig(config map[string]string, name string) *HeaderConfig {
	this := HeaderConfig{}
	this.Config = config
	this.Name = name
	return &this
}

// NewHeaderConfigWithDefaults instantiates a new HeaderConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeaderConfigWithDefaults() *HeaderConfig {
	this := HeaderConfig{}
	return &this
}

// GetConfig returns the Config field value
func (o *HeaderConfig) GetConfig() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *HeaderConfig) GetConfigOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *HeaderConfig) SetConfig(v map[string]string) {
	o.Config = v
}

// GetName returns the Name field value
func (o *HeaderConfig) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HeaderConfig) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *HeaderConfig) SetName(v string) {
	o.Name = v
}

func (o HeaderConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HeaderConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["config"] = o.Config
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HeaderConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"config",
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

	varHeaderConfig := _HeaderConfig{}

	err = json.Unmarshal(data, &varHeaderConfig)

	if err != nil {
		return err
	}

	*o = HeaderConfig(varHeaderConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "config")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHeaderConfig struct {
	value *HeaderConfig
	isSet bool
}

func (v NullableHeaderConfig) Get() *HeaderConfig {
	return v.value
}

func (v *NullableHeaderConfig) Set(val *HeaderConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableHeaderConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableHeaderConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeaderConfig(val *HeaderConfig) *NullableHeaderConfig {
	return &NullableHeaderConfig{value: val, isSet: true}
}

func (v NullableHeaderConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeaderConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


