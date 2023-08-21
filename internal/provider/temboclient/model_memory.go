/*
cp-webserver

CoreDB backend API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package temboclient

import (
	"encoding/json"
	"fmt"
)

// Memory the model 'Memory'
type Memory string

// List of Memory
const (
	_1_GI Memory = "1Gi"
	_2_GI Memory = "2Gi"
	_4_GI Memory = "4Gi"
	_8_GI Memory = "8Gi"
	_16_GI Memory = "16Gi"
	_32_GI Memory = "32Gi"
)

// All allowed values of Memory enum
var AllowedMemoryEnumValues = []Memory{
	"1Gi",
	"2Gi",
	"4Gi",
	"8Gi",
	"16Gi",
	"32Gi",
}

func (v *Memory) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Memory(value)
	for _, existing := range AllowedMemoryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Memory", value)
}

// NewMemoryFromValue returns a pointer to a valid Memory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMemoryFromValue(v string) (*Memory, error) {
	ev := Memory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Memory: valid values are %v", v, AllowedMemoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Memory) IsValid() bool {
	for _, existing := range AllowedMemoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Memory value
func (v Memory) Ptr() *Memory {
	return &v
}

type NullableMemory struct {
	value *Memory
	isSet bool
}

func (v NullableMemory) Get() *Memory {
	return v.value
}

func (v *NullableMemory) Set(val *Memory) {
	v.value = val
	v.isSet = true
}

func (v NullableMemory) IsSet() bool {
	return v.isSet
}

func (v *NullableMemory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemory(val *Memory) *NullableMemory {
	return &NullableMemory{value: val, isSet: true}
}

func (v NullableMemory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
