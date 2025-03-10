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

// checks if the InstanceState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceState{}

// InstanceState struct for InstanceState
type InstanceState struct {
	State State `json:"state"`
	AdditionalProperties map[string]interface{}
}

type _InstanceState InstanceState

// NewInstanceState instantiates a new InstanceState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceState(state State) *InstanceState {
	this := InstanceState{}
	this.State = state
	return &this
}

// NewInstanceStateWithDefaults instantiates a new InstanceState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceStateWithDefaults() *InstanceState {
	this := InstanceState{}
	return &this
}

// GetState returns the State field value
func (o *InstanceState) GetState() State {
	if o == nil {
		var ret State
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *InstanceState) GetStateOk() (*State, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *InstanceState) SetState(v State) {
	o.State = v
}

func (o InstanceState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["state"] = o.State

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InstanceState) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"state",
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

	varInstanceState := _InstanceState{}

	err = json.Unmarshal(data, &varInstanceState)

	if err != nil {
		return err
	}

	*o = InstanceState(varInstanceState)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "state")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInstanceState struct {
	value *InstanceState
	isSet bool
}

func (v NullableInstanceState) Get() *InstanceState {
	return v.value
}

func (v *NullableInstanceState) Set(val *InstanceState) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceState) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceState(val *InstanceState) *NullableInstanceState {
	return &NullableInstanceState{value: val, isSet: true}
}

func (v NullableInstanceState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


