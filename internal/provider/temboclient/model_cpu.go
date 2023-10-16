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

// Cpu the model 'Cpu'
type Cpu string

// List of Cpu
const (
	_1 Cpu = "1"
	_2 Cpu = "2"
	_4 Cpu = "4"
	_8 Cpu = "8"
	_16 Cpu = "16"
	_32 Cpu = "32"
)

// All allowed values of Cpu enum
var AllowedCpuEnumValues = []Cpu{
	"1",
	"2",
	"4",
	"8",
	"16",
	"32",
}

func (v *Cpu) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Cpu(value)
	for _, existing := range AllowedCpuEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Cpu", value)
}

// NewCpuFromValue returns a pointer to a valid Cpu
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCpuFromValue(v string) (*Cpu, error) {
	ev := Cpu(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Cpu: valid values are %v", v, AllowedCpuEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Cpu) IsValid() bool {
	for _, existing := range AllowedCpuEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Cpu value
func (v Cpu) Ptr() *Cpu {
	return &v
}

type NullableCpu struct {
	value *Cpu
	isSet bool
}

func (v NullableCpu) Get() *Cpu {
	return v.value
}

func (v *NullableCpu) Set(val *Cpu) {
	v.value = val
	v.isSet = true
}

func (v NullableCpu) IsSet() bool {
	return v.isSet
}

func (v *NullableCpu) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCpu(val *Cpu) *NullableCpu {
	return &NullableCpu{value: val, isSet: true}
}

func (v NullableCpu) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCpu) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

