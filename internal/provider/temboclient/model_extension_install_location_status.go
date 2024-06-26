/*
Tembo Cloud

Platform API for Tembo Cloud             </br>             </br>             To find a Tembo Data API, please find it here:             </br>             </br>             [AWS US East 1](https://api.data-1.use1.tembo.io/swagger-ui/)             

API version: v1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package temboclient

import (
	"encoding/json"
)

// checks if the ExtensionInstallLocationStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtensionInstallLocationStatus{}

// ExtensionInstallLocationStatus struct for ExtensionInstallLocationStatus
type ExtensionInstallLocationStatus struct {
	Database *string `json:"database,omitempty"`
	Enabled NullableBool `json:"enabled,omitempty"`
	Error NullableBool `json:"error,omitempty"`
	ErrorMessage NullableString `json:"error_message,omitempty"`
	Schema NullableString `json:"schema,omitempty"`
	Version NullableString `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExtensionInstallLocationStatus ExtensionInstallLocationStatus

// NewExtensionInstallLocationStatus instantiates a new ExtensionInstallLocationStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionInstallLocationStatus() *ExtensionInstallLocationStatus {
	this := ExtensionInstallLocationStatus{}
	return &this
}

// NewExtensionInstallLocationStatusWithDefaults instantiates a new ExtensionInstallLocationStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionInstallLocationStatusWithDefaults() *ExtensionInstallLocationStatus {
	this := ExtensionInstallLocationStatus{}
	return &this
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *ExtensionInstallLocationStatus) GetDatabase() string {
	if o == nil || IsNil(o.Database) {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionInstallLocationStatus) GetDatabaseOk() (*string, bool) {
	if o == nil || IsNil(o.Database) {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *ExtensionInstallLocationStatus) HasDatabase() bool {
	if o != nil && !IsNil(o.Database) {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *ExtensionInstallLocationStatus) SetDatabase(v string) {
	o.Database = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExtensionInstallLocationStatus) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExtensionInstallLocationStatus) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *ExtensionInstallLocationStatus) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *ExtensionInstallLocationStatus) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}
// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *ExtensionInstallLocationStatus) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *ExtensionInstallLocationStatus) UnsetEnabled() {
	o.Enabled.Unset()
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExtensionInstallLocationStatus) GetError() bool {
	if o == nil || IsNil(o.Error.Get()) {
		var ret bool
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExtensionInstallLocationStatus) GetErrorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *ExtensionInstallLocationStatus) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableBool and assigns it to the Error field.
func (o *ExtensionInstallLocationStatus) SetError(v bool) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *ExtensionInstallLocationStatus) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *ExtensionInstallLocationStatus) UnsetError() {
	o.Error.Unset()
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExtensionInstallLocationStatus) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExtensionInstallLocationStatus) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ExtensionInstallLocationStatus) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *ExtensionInstallLocationStatus) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}
// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *ExtensionInstallLocationStatus) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *ExtensionInstallLocationStatus) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

// GetSchema returns the Schema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExtensionInstallLocationStatus) GetSchema() string {
	if o == nil || IsNil(o.Schema.Get()) {
		var ret string
		return ret
	}
	return *o.Schema.Get()
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExtensionInstallLocationStatus) GetSchemaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schema.Get(), o.Schema.IsSet()
}

// HasSchema returns a boolean if a field has been set.
func (o *ExtensionInstallLocationStatus) HasSchema() bool {
	if o != nil && o.Schema.IsSet() {
		return true
	}

	return false
}

// SetSchema gets a reference to the given NullableString and assigns it to the Schema field.
func (o *ExtensionInstallLocationStatus) SetSchema(v string) {
	o.Schema.Set(&v)
}
// SetSchemaNil sets the value for Schema to be an explicit nil
func (o *ExtensionInstallLocationStatus) SetSchemaNil() {
	o.Schema.Set(nil)
}

// UnsetSchema ensures that no value is present for Schema, not even an explicit nil
func (o *ExtensionInstallLocationStatus) UnsetSchema() {
	o.Schema.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExtensionInstallLocationStatus) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExtensionInstallLocationStatus) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *ExtensionInstallLocationStatus) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *ExtensionInstallLocationStatus) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *ExtensionInstallLocationStatus) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *ExtensionInstallLocationStatus) UnsetVersion() {
	o.Version.Unset()
}

func (o ExtensionInstallLocationStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtensionInstallLocationStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Database) {
		toSerialize["database"] = o.Database
	}
	if o.Enabled.IsSet() {
		toSerialize["enabled"] = o.Enabled.Get()
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["error_message"] = o.ErrorMessage.Get()
	}
	if o.Schema.IsSet() {
		toSerialize["schema"] = o.Schema.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExtensionInstallLocationStatus) UnmarshalJSON(data []byte) (err error) {
	varExtensionInstallLocationStatus := _ExtensionInstallLocationStatus{}

	err = json.Unmarshal(data, &varExtensionInstallLocationStatus)

	if err != nil {
		return err
	}

	*o = ExtensionInstallLocationStatus(varExtensionInstallLocationStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "database")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "error")
		delete(additionalProperties, "error_message")
		delete(additionalProperties, "schema")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExtensionInstallLocationStatus struct {
	value *ExtensionInstallLocationStatus
	isSet bool
}

func (v NullableExtensionInstallLocationStatus) Get() *ExtensionInstallLocationStatus {
	return v.value
}

func (v *NullableExtensionInstallLocationStatus) Set(val *ExtensionInstallLocationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionInstallLocationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionInstallLocationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionInstallLocationStatus(val *ExtensionInstallLocationStatus) *NullableExtensionInstallLocationStatus {
	return &NullableExtensionInstallLocationStatus{value: val, isSet: true}
}

func (v NullableExtensionInstallLocationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionInstallLocationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


