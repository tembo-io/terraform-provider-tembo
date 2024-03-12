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
	"bytes"
	"fmt"
)

// checks if the Restore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Restore{}

// Restore struct for Restore
type Restore struct {
	InstanceId string `json:"instance_id"`
	RecoveryTargetTime NullableTime `json:"recovery_target_time,omitempty"`
}

type _Restore Restore

// NewRestore instantiates a new Restore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestore(instanceId string) *Restore {
	this := Restore{}
	this.InstanceId = instanceId
	return &this
}

// NewRestoreWithDefaults instantiates a new Restore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoreWithDefaults() *Restore {
	this := Restore{}
	return &this
}

// GetInstanceId returns the InstanceId field value
func (o *Restore) GetInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *Restore) GetInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value
func (o *Restore) SetInstanceId(v string) {
	o.InstanceId = v
}

// GetRecoveryTargetTime returns the RecoveryTargetTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Restore) GetRecoveryTargetTime() time.Time {
	if o == nil || IsNil(o.RecoveryTargetTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.RecoveryTargetTime.Get()
}

// GetRecoveryTargetTimeOk returns a tuple with the RecoveryTargetTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Restore) GetRecoveryTargetTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecoveryTargetTime.Get(), o.RecoveryTargetTime.IsSet()
}

// HasRecoveryTargetTime returns a boolean if a field has been set.
func (o *Restore) HasRecoveryTargetTime() bool {
	if o != nil && o.RecoveryTargetTime.IsSet() {
		return true
	}

	return false
}

// SetRecoveryTargetTime gets a reference to the given NullableTime and assigns it to the RecoveryTargetTime field.
func (o *Restore) SetRecoveryTargetTime(v time.Time) {
	o.RecoveryTargetTime.Set(&v)
}
// SetRecoveryTargetTimeNil sets the value for RecoveryTargetTime to be an explicit nil
func (o *Restore) SetRecoveryTargetTimeNil() {
	o.RecoveryTargetTime.Set(nil)
}

// UnsetRecoveryTargetTime ensures that no value is present for RecoveryTargetTime, not even an explicit nil
func (o *Restore) UnsetRecoveryTargetTime() {
	o.RecoveryTargetTime.Unset()
}

func (o Restore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Restore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["instance_id"] = o.InstanceId
	if o.RecoveryTargetTime.IsSet() {
		toSerialize["recovery_target_time"] = o.RecoveryTargetTime.Get()
	}
	return toSerialize, nil
}

func (o *Restore) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"instance_id",
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

	varRestore := _Restore{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRestore)

	if err != nil {
		return err
	}

	*o = Restore(varRestore)

	return err
}

type NullableRestore struct {
	value *Restore
	isSet bool
}

func (v NullableRestore) Get() *Restore {
	return v.value
}

func (v *NullableRestore) Set(val *Restore) {
	v.value = val
	v.isSet = true
}

func (v NullableRestore) IsSet() bool {
	return v.isSet
}

func (v *NullableRestore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestore(val *Restore) *NullableRestore {
	return &NullableRestore{value: val, isSet: true}
}

func (v NullableRestore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


