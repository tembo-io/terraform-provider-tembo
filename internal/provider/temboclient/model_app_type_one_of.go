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

// checks if the AppTypeOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppTypeOneOf{}

// AppTypeOneOf struct for AppTypeOneOf
type AppTypeOneOf struct {
	Restapi NullableAppConfig `json:"restapi"`
	AdditionalProperties map[string]interface{}
}

type _AppTypeOneOf AppTypeOneOf

// NewAppTypeOneOf instantiates a new AppTypeOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppTypeOneOf(restapi NullableAppConfig) *AppTypeOneOf {
	this := AppTypeOneOf{}
	this.Restapi = restapi
	return &this
}

// NewAppTypeOneOfWithDefaults instantiates a new AppTypeOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppTypeOneOfWithDefaults() *AppTypeOneOf {
	this := AppTypeOneOf{}
	return &this
}

// GetRestapi returns the Restapi field value
// If the value is explicit nil, the zero value for AppConfig will be returned
func (o *AppTypeOneOf) GetRestapi() AppConfig {
	if o == nil || o.Restapi.Get() == nil {
		var ret AppConfig
		return ret
	}

	return *o.Restapi.Get()
}

// GetRestapiOk returns a tuple with the Restapi field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppTypeOneOf) GetRestapiOk() (*AppConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.Restapi.Get(), o.Restapi.IsSet()
}

// SetRestapi sets field value
func (o *AppTypeOneOf) SetRestapi(v AppConfig) {
	o.Restapi.Set(&v)
}

func (o AppTypeOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppTypeOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["restapi"] = o.Restapi.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppTypeOneOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"restapi",
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

	varAppTypeOneOf := _AppTypeOneOf{}

	err = json.Unmarshal(data, &varAppTypeOneOf)

	if err != nil {
		return err
	}

	*o = AppTypeOneOf(varAppTypeOneOf)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "restapi")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppTypeOneOf struct {
	value *AppTypeOneOf
	isSet bool
}

func (v NullableAppTypeOneOf) Get() *AppTypeOneOf {
	return v.value
}

func (v *NullableAppTypeOneOf) Set(val *AppTypeOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAppTypeOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAppTypeOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppTypeOneOf(val *AppTypeOneOf) *NullableAppTypeOneOf {
	return &NullableAppTypeOneOf{value: val, isSet: true}
}

func (v NullableAppTypeOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppTypeOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


