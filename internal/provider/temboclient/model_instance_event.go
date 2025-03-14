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

// InstanceEvent the model 'InstanceEvent'
type InstanceEvent string

// List of InstanceEvent
const (
	RESTART InstanceEvent = "restart"
	STOP InstanceEvent = "stop"
	START InstanceEvent = "start"
)

// All allowed values of InstanceEvent enum
var AllowedInstanceEventEnumValues = []InstanceEvent{
	"restart",
	"stop",
	"start",
}

func (v *InstanceEvent) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InstanceEvent(value)
	for _, existing := range AllowedInstanceEventEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InstanceEvent", value)
}

// NewInstanceEventFromValue returns a pointer to a valid InstanceEvent
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInstanceEventFromValue(v string) (*InstanceEvent, error) {
	ev := InstanceEvent(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InstanceEvent: valid values are %v", v, AllowedInstanceEventEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InstanceEvent) IsValid() bool {
	for _, existing := range AllowedInstanceEventEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InstanceEvent value
func (v InstanceEvent) Ptr() *InstanceEvent {
	return &v
}

type NullableInstanceEvent struct {
	value *InstanceEvent
	isSet bool
}

func (v NullableInstanceEvent) Get() *InstanceEvent {
	return v.value
}

func (v *NullableInstanceEvent) Set(val *InstanceEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceEvent(val *InstanceEvent) *NullableInstanceEvent {
	return &NullableInstanceEvent{value: val, isSet: true}
}

func (v NullableInstanceEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

