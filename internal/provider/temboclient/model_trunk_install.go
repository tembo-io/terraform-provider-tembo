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

// checks if the TrunkInstall type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrunkInstall{}

// TrunkInstall TrunkInstall allows installation of extensions from the [pgtrunk](https://pgt.dev) registry.  This list should be a list of extension names and versions that you wish to install at runtime using the pgtrunk API.  This example will install the pg_stat_statements extension at version 1.10.0.  ```yaml apiVersion: coredb.io/v1alpha1 kind: CoreDB metadata: name: test-db spec: trunk_installs: - name: pg_stat_statements version: 1.10.0 ```
type TrunkInstall struct {
	// The name of the extension to install. This must be the name of the extension as it appears in the [pgtrunk](https://pgt.dev) registry.
	Name string `json:"name"`
	// The version of the extension to install. If not specified, the latest version will be used. (Optional)
	Version NullableString `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TrunkInstall TrunkInstall

// NewTrunkInstall instantiates a new TrunkInstall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrunkInstall(name string) *TrunkInstall {
	this := TrunkInstall{}
	this.Name = name
	return &this
}

// NewTrunkInstallWithDefaults instantiates a new TrunkInstall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrunkInstallWithDefaults() *TrunkInstall {
	this := TrunkInstall{}
	return &this
}

// GetName returns the Name field value
func (o *TrunkInstall) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TrunkInstall) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TrunkInstall) SetName(v string) {
	o.Name = v
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrunkInstall) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrunkInstall) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *TrunkInstall) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *TrunkInstall) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *TrunkInstall) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *TrunkInstall) UnsetVersion() {
	o.Version.Unset()
}

func (o TrunkInstall) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrunkInstall) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TrunkInstall) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varTrunkInstall := _TrunkInstall{}

	err = json.Unmarshal(data, &varTrunkInstall)

	if err != nil {
		return err
	}

	*o = TrunkInstall(varTrunkInstall)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTrunkInstall struct {
	value *TrunkInstall
	isSet bool
}

func (v NullableTrunkInstall) Get() *TrunkInstall {
	return v.value
}

func (v *NullableTrunkInstall) Set(val *TrunkInstall) {
	v.value = val
	v.isSet = true
}

func (v NullableTrunkInstall) IsSet() bool {
	return v.isSet
}

func (v *NullableTrunkInstall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrunkInstall(val *TrunkInstall) *NullableTrunkInstall {
	return &NullableTrunkInstall{value: val, isSet: true}
}

func (v NullableTrunkInstall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrunkInstall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


