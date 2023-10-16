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

// StackType the model 'StackType'
type StackType string

// List of StackType
const (
	STANDARD StackType = "Standard"
	MESSAGE_QUEUE StackType = "MessageQueue"
	MACHINE_LEARNING StackType = "MachineLearning"
	OLAP StackType = "OLAP"
	OLTP StackType = "OLTP"
)

// All allowed values of StackType enum
var AllowedStackTypeEnumValues = []StackType{
	"Standard",
	"MessageQueue",
	"MachineLearning",
	"OLAP",
	"OLTP",
}

func (v *StackType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StackType(value)
	for _, existing := range AllowedStackTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StackType", value)
}

// NewStackTypeFromValue returns a pointer to a valid StackType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStackTypeFromValue(v string) (*StackType, error) {
	ev := StackType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StackType: valid values are %v", v, AllowedStackTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StackType) IsValid() bool {
	for _, existing := range AllowedStackTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StackType value
func (v StackType) Ptr() *StackType {
	return &v
}

type NullableStackType struct {
	value *StackType
	isSet bool
}

func (v NullableStackType) Get() *StackType {
	return v.value
}

func (v *NullableStackType) Set(val *StackType) {
	v.value = val
	v.isSet = true
}

func (v NullableStackType) IsSet() bool {
	return v.isSet
}

func (v *NullableStackType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStackType(val *StackType) *NullableStackType {
	return &NullableStackType{value: val, isSet: true}
}

func (v NullableStackType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStackType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

