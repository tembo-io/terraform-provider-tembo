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

// PoolerPgbouncerPoolMode The PgBouncer configuration
type PoolerPgbouncerPoolMode string

// List of PoolerPgbouncerPoolMode
const (
	SESSION PoolerPgbouncerPoolMode = "session"
	TRANSACTION PoolerPgbouncerPoolMode = "transaction"
)

// All allowed values of PoolerPgbouncerPoolMode enum
var AllowedPoolerPgbouncerPoolModeEnumValues = []PoolerPgbouncerPoolMode{
	"session",
	"transaction",
}

func (v *PoolerPgbouncerPoolMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PoolerPgbouncerPoolMode(value)
	for _, existing := range AllowedPoolerPgbouncerPoolModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PoolerPgbouncerPoolMode", value)
}

// NewPoolerPgbouncerPoolModeFromValue returns a pointer to a valid PoolerPgbouncerPoolMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPoolerPgbouncerPoolModeFromValue(v string) (*PoolerPgbouncerPoolMode, error) {
	ev := PoolerPgbouncerPoolMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PoolerPgbouncerPoolMode: valid values are %v", v, AllowedPoolerPgbouncerPoolModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PoolerPgbouncerPoolMode) IsValid() bool {
	for _, existing := range AllowedPoolerPgbouncerPoolModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PoolerPgbouncerPoolMode value
func (v PoolerPgbouncerPoolMode) Ptr() *PoolerPgbouncerPoolMode {
	return &v
}

type NullablePoolerPgbouncerPoolMode struct {
	value *PoolerPgbouncerPoolMode
	isSet bool
}

func (v NullablePoolerPgbouncerPoolMode) Get() *PoolerPgbouncerPoolMode {
	return v.value
}

func (v *NullablePoolerPgbouncerPoolMode) Set(val *PoolerPgbouncerPoolMode) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolerPgbouncerPoolMode) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolerPgbouncerPoolMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolerPgbouncerPoolMode(val *PoolerPgbouncerPoolMode) *NullablePoolerPgbouncerPoolMode {
	return &NullablePoolerPgbouncerPoolMode{value: val, isSet: true}
}

func (v NullablePoolerPgbouncerPoolMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolerPgbouncerPoolMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

