/*
Tembo Cloud

Platform API for Tembo Cloud             </br>             </br>             To find a Tembo Data API, please find it here:             </br>             </br>             [AWS US East 1](https://api.data-1.use1.tembo.io/swagger-ui/)             

API version: v1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package temboclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AppTypeOneOf5 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppTypeOneOf5{}

// AppTypeOneOf5 struct for AppTypeOneOf5
type AppTypeOneOf5 struct {
	Custom AppService `json:"custom"`
}

type _AppTypeOneOf5 AppTypeOneOf5

// NewAppTypeOneOf5 instantiates a new AppTypeOneOf5 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppTypeOneOf5(custom AppService) *AppTypeOneOf5 {
	this := AppTypeOneOf5{}
	this.Custom = custom
	return &this
}

// NewAppTypeOneOf5WithDefaults instantiates a new AppTypeOneOf5 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppTypeOneOf5WithDefaults() *AppTypeOneOf5 {
	this := AppTypeOneOf5{}
	return &this
}

// GetCustom returns the Custom field value
func (o *AppTypeOneOf5) GetCustom() AppService {
	if o == nil {
		var ret AppService
		return ret
	}

	return o.Custom
}

// GetCustomOk returns a tuple with the Custom field value
// and a boolean to check if the value has been set.
func (o *AppTypeOneOf5) GetCustomOk() (*AppService, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Custom, true
}

// SetCustom sets field value
func (o *AppTypeOneOf5) SetCustom(v AppService) {
	o.Custom = v
}

func (o AppTypeOneOf5) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppTypeOneOf5) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["custom"] = o.Custom
	return toSerialize, nil
}

func (o *AppTypeOneOf5) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"custom",
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

	varAppTypeOneOf5 := _AppTypeOneOf5{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppTypeOneOf5)

	if err != nil {
		return err
	}

	*o = AppTypeOneOf5(varAppTypeOneOf5)

	return err
}

type NullableAppTypeOneOf5 struct {
	value *AppTypeOneOf5
	isSet bool
}

func (v NullableAppTypeOneOf5) Get() *AppTypeOneOf5 {
	return v.value
}

func (v *NullableAppTypeOneOf5) Set(val *AppTypeOneOf5) {
	v.value = val
	v.isSet = true
}

func (v NullableAppTypeOneOf5) IsSet() bool {
	return v.isSet
}

func (v *NullableAppTypeOneOf5) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppTypeOneOf5(val *AppTypeOneOf5) *NullableAppTypeOneOf5 {
	return &NullableAppTypeOneOf5{value: val, isSet: true}
}

func (v NullableAppTypeOneOf5) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppTypeOneOf5) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


