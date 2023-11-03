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

// checks if the AppTypeOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppTypeOneOf1{}

// AppTypeOneOf1 struct for AppTypeOneOf1
type AppTypeOneOf1 struct {
	Http NullableAppConfig `json:"http"`
}

// NewAppTypeOneOf1 instantiates a new AppTypeOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppTypeOneOf1(http NullableAppConfig) *AppTypeOneOf1 {
	this := AppTypeOneOf1{}
	this.Http = http
	return &this
}

// NewAppTypeOneOf1WithDefaults instantiates a new AppTypeOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppTypeOneOf1WithDefaults() *AppTypeOneOf1 {
	this := AppTypeOneOf1{}
	return &this
}

// GetHttp returns the Http field value
// If the value is explicit nil, the zero value for AppConfig will be returned
func (o *AppTypeOneOf1) GetHttp() AppConfig {
	if o == nil || o.Http.Get() == nil {
		var ret AppConfig
		return ret
	}

	return *o.Http.Get()
}

// GetHttpOk returns a tuple with the Http field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppTypeOneOf1) GetHttpOk() (*AppConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.Http.Get(), o.Http.IsSet()
}

// SetHttp sets field value
func (o *AppTypeOneOf1) SetHttp(v AppConfig) {
	o.Http.Set(&v)
}

func (o AppTypeOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppTypeOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["http"] = o.Http.Get()
	return toSerialize, nil
}

type NullableAppTypeOneOf1 struct {
	value *AppTypeOneOf1
	isSet bool
}

func (v NullableAppTypeOneOf1) Get() *AppTypeOneOf1 {
	return v.value
}

func (v *NullableAppTypeOneOf1) Set(val *AppTypeOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableAppTypeOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableAppTypeOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppTypeOneOf1(val *AppTypeOneOf1) *NullableAppTypeOneOf1 {
	return &NullableAppTypeOneOf1{value: val, isSet: true}
}

func (v NullableAppTypeOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppTypeOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


