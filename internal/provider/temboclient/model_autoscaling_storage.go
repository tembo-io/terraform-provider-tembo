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

// checks if the AutoscalingStorage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoscalingStorage{}

// AutoscalingStorage AutoscalingStorage configures the automatic scaling of instance storage.
type AutoscalingStorage struct {
	// Enable autoscaling for storage
	Enabled bool `json:"enabled"`
	Maximum NullableStorage `json:"maximum,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AutoscalingStorage AutoscalingStorage

// NewAutoscalingStorage instantiates a new AutoscalingStorage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoscalingStorage(enabled bool) *AutoscalingStorage {
	this := AutoscalingStorage{}
	this.Enabled = enabled
	return &this
}

// NewAutoscalingStorageWithDefaults instantiates a new AutoscalingStorage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoscalingStorageWithDefaults() *AutoscalingStorage {
	this := AutoscalingStorage{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *AutoscalingStorage) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AutoscalingStorage) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AutoscalingStorage) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMaximum returns the Maximum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoscalingStorage) GetMaximum() Storage {
	if o == nil || IsNil(o.Maximum.Get()) {
		var ret Storage
		return ret
	}
	return *o.Maximum.Get()
}

// GetMaximumOk returns a tuple with the Maximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoscalingStorage) GetMaximumOk() (*Storage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Maximum.Get(), o.Maximum.IsSet()
}

// HasMaximum returns a boolean if a field has been set.
func (o *AutoscalingStorage) HasMaximum() bool {
	if o != nil && o.Maximum.IsSet() {
		return true
	}

	return false
}

// SetMaximum gets a reference to the given NullableStorage and assigns it to the Maximum field.
func (o *AutoscalingStorage) SetMaximum(v Storage) {
	o.Maximum.Set(&v)
}
// SetMaximumNil sets the value for Maximum to be an explicit nil
func (o *AutoscalingStorage) SetMaximumNil() {
	o.Maximum.Set(nil)
}

// UnsetMaximum ensures that no value is present for Maximum, not even an explicit nil
func (o *AutoscalingStorage) UnsetMaximum() {
	o.Maximum.Unset()
}

func (o AutoscalingStorage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoscalingStorage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	if o.Maximum.IsSet() {
		toSerialize["maximum"] = o.Maximum.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AutoscalingStorage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
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

	varAutoscalingStorage := _AutoscalingStorage{}

	err = json.Unmarshal(data, &varAutoscalingStorage)

	if err != nil {
		return err
	}

	*o = AutoscalingStorage(varAutoscalingStorage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "maximum")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAutoscalingStorage struct {
	value *AutoscalingStorage
	isSet bool
}

func (v NullableAutoscalingStorage) Get() *AutoscalingStorage {
	return v.value
}

func (v *NullableAutoscalingStorage) Set(val *AutoscalingStorage) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoscalingStorage) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoscalingStorage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoscalingStorage(val *AutoscalingStorage) *NullableAutoscalingStorage {
	return &NullableAutoscalingStorage{value: val, isSet: true}
}

func (v NullableAutoscalingStorage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoscalingStorage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


