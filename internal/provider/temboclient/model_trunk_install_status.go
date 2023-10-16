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

// checks if the TrunkInstallStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrunkInstallStatus{}

// TrunkInstallStatus struct for TrunkInstallStatus
type TrunkInstallStatus struct {
	Error bool `json:"error"`
	ErrorMessage NullableString `json:"error_message,omitempty"`
	InstalledToPods []string `json:"installed_to_pods,omitempty"`
	Name string `json:"name"`
	Version NullableString `json:"version,omitempty"`
}

// NewTrunkInstallStatus instantiates a new TrunkInstallStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrunkInstallStatus(error_ bool, name string) *TrunkInstallStatus {
	this := TrunkInstallStatus{}
	this.Error = error_
	this.Name = name
	return &this
}

// NewTrunkInstallStatusWithDefaults instantiates a new TrunkInstallStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrunkInstallStatusWithDefaults() *TrunkInstallStatus {
	this := TrunkInstallStatus{}
	return &this
}

// GetError returns the Error field value
func (o *TrunkInstallStatus) GetError() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *TrunkInstallStatus) GetErrorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *TrunkInstallStatus) SetError(v bool) {
	o.Error = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrunkInstallStatus) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrunkInstallStatus) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *TrunkInstallStatus) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *TrunkInstallStatus) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}
// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *TrunkInstallStatus) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *TrunkInstallStatus) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

// GetInstalledToPods returns the InstalledToPods field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrunkInstallStatus) GetInstalledToPods() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.InstalledToPods
}

// GetInstalledToPodsOk returns a tuple with the InstalledToPods field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrunkInstallStatus) GetInstalledToPodsOk() ([]string, bool) {
	if o == nil || IsNil(o.InstalledToPods) {
		return nil, false
	}
	return o.InstalledToPods, true
}

// HasInstalledToPods returns a boolean if a field has been set.
func (o *TrunkInstallStatus) HasInstalledToPods() bool {
	if o != nil && IsNil(o.InstalledToPods) {
		return true
	}

	return false
}

// SetInstalledToPods gets a reference to the given []string and assigns it to the InstalledToPods field.
func (o *TrunkInstallStatus) SetInstalledToPods(v []string) {
	o.InstalledToPods = v
}

// GetName returns the Name field value
func (o *TrunkInstallStatus) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TrunkInstallStatus) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TrunkInstallStatus) SetName(v string) {
	o.Name = v
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrunkInstallStatus) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrunkInstallStatus) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *TrunkInstallStatus) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *TrunkInstallStatus) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *TrunkInstallStatus) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *TrunkInstallStatus) UnsetVersion() {
	o.Version.Unset()
}

func (o TrunkInstallStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrunkInstallStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	if o.ErrorMessage.IsSet() {
		toSerialize["error_message"] = o.ErrorMessage.Get()
	}
	if o.InstalledToPods != nil {
		toSerialize["installed_to_pods"] = o.InstalledToPods
	}
	toSerialize["name"] = o.Name
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	return toSerialize, nil
}

type NullableTrunkInstallStatus struct {
	value *TrunkInstallStatus
	isSet bool
}

func (v NullableTrunkInstallStatus) Get() *TrunkInstallStatus {
	return v.value
}

func (v *NullableTrunkInstallStatus) Set(val *TrunkInstallStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTrunkInstallStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTrunkInstallStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrunkInstallStatus(val *TrunkInstallStatus) *NullableTrunkInstallStatus {
	return &NullableTrunkInstallStatus{value: val, isSet: true}
}

func (v NullableTrunkInstallStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrunkInstallStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


