/*
Tembo Cloud

Platform API for Tembo Cloud             </br>             </br>             To find a Tembo Data API, please find it here:             </br>             </br>             [AWS US East 1](https://api.data-1.use1.tembo.io/swagger-ui/)             

API version: v1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package temboclient

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the HistoryEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoryEntry{}

// HistoryEntry struct for HistoryEntry
type HistoryEntry struct {
	Instance InstanceState `json:"instance"`
	// The timestamp of the entry
	Timestamp time.Time `json:"timestamp"`
	AdditionalProperties map[string]interface{}
}

type _HistoryEntry HistoryEntry

// NewHistoryEntry instantiates a new HistoryEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryEntry(instance InstanceState, timestamp time.Time) *HistoryEntry {
	this := HistoryEntry{}
	this.Instance = instance
	this.Timestamp = timestamp
	return &this
}

// NewHistoryEntryWithDefaults instantiates a new HistoryEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryEntryWithDefaults() *HistoryEntry {
	this := HistoryEntry{}
	return &this
}

// GetInstance returns the Instance field value
func (o *HistoryEntry) GetInstance() InstanceState {
	if o == nil {
		var ret InstanceState
		return ret
	}

	return o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value
// and a boolean to check if the value has been set.
func (o *HistoryEntry) GetInstanceOk() (*InstanceState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Instance, true
}

// SetInstance sets field value
func (o *HistoryEntry) SetInstance(v InstanceState) {
	o.Instance = v
}

// GetTimestamp returns the Timestamp field value
func (o *HistoryEntry) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *HistoryEntry) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *HistoryEntry) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

func (o HistoryEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoryEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["instance"] = o.Instance
	toSerialize["timestamp"] = o.Timestamp

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HistoryEntry) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"instance",
		"timestamp",
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

	varHistoryEntry := _HistoryEntry{}

	err = json.Unmarshal(data, &varHistoryEntry)

	if err != nil {
		return err
	}

	*o = HistoryEntry(varHistoryEntry)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "instance")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHistoryEntry struct {
	value *HistoryEntry
	isSet bool
}

func (v NullableHistoryEntry) Get() *HistoryEntry {
	return v.value
}

func (v *NullableHistoryEntry) Set(val *HistoryEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryEntry(val *HistoryEntry) *NullableHistoryEntry {
	return &NullableHistoryEntry{value: val, isSet: true}
}

func (v NullableHistoryEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


