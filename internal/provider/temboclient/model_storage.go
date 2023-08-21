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

// Storage the model 'Storage'
type Storage string

// List of Storage
const (
	_10_GI Storage = "10Gi"
	_50_GI Storage = "50Gi"
	_100_GI Storage = "100Gi"
	_200_GI Storage = "200Gi"
	_300_GI Storage = "300Gi"
	_400_GI Storage = "400Gi"
	_500_GI Storage = "500Gi"
)

// All allowed values of Storage enum
var AllowedStorageEnumValues = []Storage{
	"10Gi",
	"50Gi",
	"100Gi",
	"200Gi",
	"300Gi",
	"400Gi",
	"500Gi",
}

func (v *Storage) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Storage(value)
	for _, existing := range AllowedStorageEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Storage", value)
}

// NewStorageFromValue returns a pointer to a valid Storage
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStorageFromValue(v string) (*Storage, error) {
	ev := Storage(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Storage: valid values are %v", v, AllowedStorageEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Storage) IsValid() bool {
	for _, existing := range AllowedStorageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Storage value
func (v Storage) Ptr() *Storage {
	return &v
}

type NullableStorage struct {
	value *Storage
	isSet bool
}

func (v NullableStorage) Get() *Storage {
	return v.value
}

func (v *NullableStorage) Set(val *Storage) {
	v.value = val
	v.isSet = true
}

func (v NullableStorage) IsSet() bool {
	return v.isSet
}

func (v *NullableStorage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorage(val *Storage) *NullableStorage {
	return &NullableStorage{value: val, isSet: true}
}

func (v NullableStorage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
