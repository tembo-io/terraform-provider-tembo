/*
Tembo Cloud

Platform API for Tembo Cloud

API version: v1.0.0-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package temboclient

import (
	"encoding/json"
	"fmt"
)

// InstanceTypes the model 'InstanceTypes'
type InstanceTypes string

// List of InstanceTypes
const (
	GENERAL_PURPOSE InstanceTypes = "GeneralPurpose"
)

// All allowed values of InstanceTypes enum
var AllowedInstanceTypesEnumValues = []InstanceTypes{
	"GeneralPurpose",
}

func (v *InstanceTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InstanceTypes(value)
	for _, existing := range AllowedInstanceTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InstanceTypes", value)
}

// NewInstanceTypesFromValue returns a pointer to a valid InstanceTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInstanceTypesFromValue(v string) (*InstanceTypes, error) {
	ev := InstanceTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InstanceTypes: valid values are %v", v, AllowedInstanceTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InstanceTypes) IsValid() bool {
	for _, existing := range AllowedInstanceTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InstanceTypes value
func (v InstanceTypes) Ptr() *InstanceTypes {
	return &v
}

type NullableInstanceTypes struct {
	value *InstanceTypes
	isSet bool
}

func (v NullableInstanceTypes) Get() *InstanceTypes {
	return v.value
}

func (v *NullableInstanceTypes) Set(val *InstanceTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceTypes(val *InstanceTypes) *NullableInstanceTypes {
	return &NullableInstanceTypes{value: val, isSet: true}
}

func (v NullableInstanceTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

