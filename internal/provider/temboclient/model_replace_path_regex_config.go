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

// checks if the ReplacePathRegexConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplacePathRegexConfig{}

// ReplacePathRegexConfig struct for ReplacePathRegexConfig
type ReplacePathRegexConfig struct {
	Config ReplacePathRegexConfigType `json:"config"`
	Name string `json:"name"`
}

// NewReplacePathRegexConfig instantiates a new ReplacePathRegexConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplacePathRegexConfig(config ReplacePathRegexConfigType, name string) *ReplacePathRegexConfig {
	this := ReplacePathRegexConfig{}
	this.Config = config
	this.Name = name
	return &this
}

// NewReplacePathRegexConfigWithDefaults instantiates a new ReplacePathRegexConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplacePathRegexConfigWithDefaults() *ReplacePathRegexConfig {
	this := ReplacePathRegexConfig{}
	return &this
}

// GetConfig returns the Config field value
func (o *ReplacePathRegexConfig) GetConfig() ReplacePathRegexConfigType {
	if o == nil {
		var ret ReplacePathRegexConfigType
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *ReplacePathRegexConfig) GetConfigOk() (*ReplacePathRegexConfigType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *ReplacePathRegexConfig) SetConfig(v ReplacePathRegexConfigType) {
	o.Config = v
}

// GetName returns the Name field value
func (o *ReplacePathRegexConfig) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ReplacePathRegexConfig) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ReplacePathRegexConfig) SetName(v string) {
	o.Name = v
}

func (o ReplacePathRegexConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplacePathRegexConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["config"] = o.Config
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableReplacePathRegexConfig struct {
	value *ReplacePathRegexConfig
	isSet bool
}

func (v NullableReplacePathRegexConfig) Get() *ReplacePathRegexConfig {
	return v.value
}

func (v *NullableReplacePathRegexConfig) Set(val *ReplacePathRegexConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableReplacePathRegexConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableReplacePathRegexConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplacePathRegexConfig(val *ReplacePathRegexConfig) *NullableReplacePathRegexConfig {
	return &NullableReplacePathRegexConfig{value: val, isSet: true}
}

func (v NullableReplacePathRegexConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplacePathRegexConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


