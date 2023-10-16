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

// State the model 'State'
type State string

// List of State
const (
	SUBMITTED State = "Submitted"
	UP State = "Up"
	CONFIGURING State = "Configuring"
	ERROR State = "Error"
	RESTARTING State = "Restarting"
	STARTING State = "Starting"
	STOPPING State = "Stopping"
	STOPPED State = "Stopped"
	DELETING State = "Deleting"
	DELETED State = "Deleted"
)

// All allowed values of State enum
var AllowedStateEnumValues = []State{
	"Submitted",
	"Up",
	"Configuring",
	"Error",
	"Restarting",
	"Starting",
	"Stopping",
	"Stopped",
	"Deleting",
	"Deleted",
}

func (v *State) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := State(value)
	for _, existing := range AllowedStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid State", value)
}

// NewStateFromValue returns a pointer to a valid State
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStateFromValue(v string) (*State, error) {
	ev := State(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for State: valid values are %v", v, AllowedStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v State) IsValid() bool {
	for _, existing := range AllowedStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to State value
func (v State) Ptr() *State {
	return &v
}

type NullableState struct {
	value *State
	isSet bool
}

func (v NullableState) Get() *State {
	return v.value
}

func (v *NullableState) Set(val *State) {
	v.value = val
	v.isSet = true
}

func (v NullableState) IsSet() bool {
	return v.isSet
}

func (v *NullableState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableState(val *State) *NullableState {
	return &NullableState{value: val, isSet: true}
}

func (v NullableState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

