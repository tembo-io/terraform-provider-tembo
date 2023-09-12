/*
Tembo Cloud

Platform API for Tembo Cloud

API version: v1.0.0-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package temboclient

import (
	"encoding/json"
)

// checks if the ExtensionInstallLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtensionInstallLocation{}

// ExtensionInstallLocation struct for ExtensionInstallLocation
type ExtensionInstallLocation struct {
	Database string `json:"database"`
	Enabled bool `json:"enabled"`
	Schema NullableString `json:"schema,omitempty"`
	Version NullableString `json:"version,omitempty"`
}

// NewExtensionInstallLocation instantiates a new ExtensionInstallLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionInstallLocation(database string, enabled bool) *ExtensionInstallLocation {
	this := ExtensionInstallLocation{}
	this.Database = database
	this.Enabled = enabled
	return &this
}

// NewExtensionInstallLocationWithDefaults instantiates a new ExtensionInstallLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionInstallLocationWithDefaults() *ExtensionInstallLocation {
	this := ExtensionInstallLocation{}
	return &this
}

// GetDatabase returns the Database field value
func (o *ExtensionInstallLocation) GetDatabase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *ExtensionInstallLocation) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Database, true
}

// SetDatabase sets field value
func (o *ExtensionInstallLocation) SetDatabase(v string) {
	o.Database = v
}

// GetEnabled returns the Enabled field value
func (o *ExtensionInstallLocation) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ExtensionInstallLocation) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ExtensionInstallLocation) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSchema returns the Schema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExtensionInstallLocation) GetSchema() string {
	if o == nil || IsNil(o.Schema.Get()) {
		var ret string
		return ret
	}
	return *o.Schema.Get()
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExtensionInstallLocation) GetSchemaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schema.Get(), o.Schema.IsSet()
}

// HasSchema returns a boolean if a field has been set.
func (o *ExtensionInstallLocation) HasSchema() bool {
	if o != nil && o.Schema.IsSet() {
		return true
	}

	return false
}

// SetSchema gets a reference to the given NullableString and assigns it to the Schema field.
func (o *ExtensionInstallLocation) SetSchema(v string) {
	o.Schema.Set(&v)
}
// SetSchemaNil sets the value for Schema to be an explicit nil
func (o *ExtensionInstallLocation) SetSchemaNil() {
	o.Schema.Set(nil)
}

// UnsetSchema ensures that no value is present for Schema, not even an explicit nil
func (o *ExtensionInstallLocation) UnsetSchema() {
	o.Schema.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExtensionInstallLocation) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExtensionInstallLocation) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *ExtensionInstallLocation) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *ExtensionInstallLocation) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *ExtensionInstallLocation) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *ExtensionInstallLocation) UnsetVersion() {
	o.Version.Unset()
}

func (o ExtensionInstallLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtensionInstallLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["database"] = o.Database
	toSerialize["enabled"] = o.Enabled
	if o.Schema.IsSet() {
		toSerialize["schema"] = o.Schema.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	return toSerialize, nil
}

type NullableExtensionInstallLocation struct {
	value *ExtensionInstallLocation
	isSet bool
}

func (v NullableExtensionInstallLocation) Get() *ExtensionInstallLocation {
	return v.value
}

func (v *NullableExtensionInstallLocation) Set(val *ExtensionInstallLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionInstallLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionInstallLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionInstallLocation(val *ExtensionInstallLocation) *NullableExtensionInstallLocation {
	return &NullableExtensionInstallLocation{value: val, isSet: true}
}

func (v NullableExtensionInstallLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionInstallLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


