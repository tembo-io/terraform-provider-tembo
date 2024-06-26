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

// checks if the AppConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppConfig{}

// AppConfig struct for AppConfig
type AppConfig struct {
	Env []EnvVar `json:"env,omitempty"`
	Resources NullableResourceRequirements `json:"resources,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppConfig AppConfig

// NewAppConfig instantiates a new AppConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppConfig() *AppConfig {
	this := AppConfig{}
	return &this
}

// NewAppConfigWithDefaults instantiates a new AppConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppConfigWithDefaults() *AppConfig {
	this := AppConfig{}
	return &this
}

// GetEnv returns the Env field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppConfig) GetEnv() []EnvVar {
	if o == nil {
		var ret []EnvVar
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppConfig) GetEnvOk() ([]EnvVar, bool) {
	if o == nil || IsNil(o.Env) {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *AppConfig) HasEnv() bool {
	if o != nil && !IsNil(o.Env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given []EnvVar and assigns it to the Env field.
func (o *AppConfig) SetEnv(v []EnvVar) {
	o.Env = v
}

// GetResources returns the Resources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppConfig) GetResources() ResourceRequirements {
	if o == nil || IsNil(o.Resources.Get()) {
		var ret ResourceRequirements
		return ret
	}
	return *o.Resources.Get()
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppConfig) GetResourcesOk() (*ResourceRequirements, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resources.Get(), o.Resources.IsSet()
}

// HasResources returns a boolean if a field has been set.
func (o *AppConfig) HasResources() bool {
	if o != nil && o.Resources.IsSet() {
		return true
	}

	return false
}

// SetResources gets a reference to the given NullableResourceRequirements and assigns it to the Resources field.
func (o *AppConfig) SetResources(v ResourceRequirements) {
	o.Resources.Set(&v)
}
// SetResourcesNil sets the value for Resources to be an explicit nil
func (o *AppConfig) SetResourcesNil() {
	o.Resources.Set(nil)
}

// UnsetResources ensures that no value is present for Resources, not even an explicit nil
func (o *AppConfig) UnsetResources() {
	o.Resources.Unset()
}

func (o AppConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	if o.Resources.IsSet() {
		toSerialize["resources"] = o.Resources.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppConfig) UnmarshalJSON(data []byte) (err error) {
	varAppConfig := _AppConfig{}

	err = json.Unmarshal(data, &varAppConfig)

	if err != nil {
		return err
	}

	*o = AppConfig(varAppConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "env")
		delete(additionalProperties, "resources")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppConfig struct {
	value *AppConfig
	isSet bool
}

func (v NullableAppConfig) Get() *AppConfig {
	return v.value
}

func (v *NullableAppConfig) Set(val *AppConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAppConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAppConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppConfig(val *AppConfig) *NullableAppConfig {
	return &NullableAppConfig{value: val, isSet: true}
}

func (v NullableAppConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


