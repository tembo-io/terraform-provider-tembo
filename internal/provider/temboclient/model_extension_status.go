/*
Tembo Cloud

Platform API for Tembo Cloud             </br>             </br>             To find a Tembo Data API, please find it here:             </br>             </br>             [AWS US East 1](https://api.data-1.use1.tembo.io/swagger-ui/)             

API version: v1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package temboclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ExtensionStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtensionStatus{}

// ExtensionStatus struct for ExtensionStatus
type ExtensionStatus struct {
	Description NullableString `json:"description,omitempty"`
	Locations []ExtensionInstallLocationStatus `json:"locations"`
	Name string `json:"name"`
}

type _ExtensionStatus ExtensionStatus

// NewExtensionStatus instantiates a new ExtensionStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionStatus(locations []ExtensionInstallLocationStatus, name string) *ExtensionStatus {
	this := ExtensionStatus{}
	this.Locations = locations
	this.Name = name
	return &this
}

// NewExtensionStatusWithDefaults instantiates a new ExtensionStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionStatusWithDefaults() *ExtensionStatus {
	this := ExtensionStatus{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExtensionStatus) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExtensionStatus) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ExtensionStatus) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ExtensionStatus) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ExtensionStatus) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ExtensionStatus) UnsetDescription() {
	o.Description.Unset()
}

// GetLocations returns the Locations field value
func (o *ExtensionStatus) GetLocations() []ExtensionInstallLocationStatus {
	if o == nil {
		var ret []ExtensionInstallLocationStatus
		return ret
	}

	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value
// and a boolean to check if the value has been set.
func (o *ExtensionStatus) GetLocationsOk() ([]ExtensionInstallLocationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Locations, true
}

// SetLocations sets field value
func (o *ExtensionStatus) SetLocations(v []ExtensionInstallLocationStatus) {
	o.Locations = v
}

// GetName returns the Name field value
func (o *ExtensionStatus) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExtensionStatus) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExtensionStatus) SetName(v string) {
	o.Name = v
}

func (o ExtensionStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtensionStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["locations"] = o.Locations
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *ExtensionStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"locations",
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

	varExtensionStatus := _ExtensionStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExtensionStatus)

	if err != nil {
		return err
	}

	*o = ExtensionStatus(varExtensionStatus)

	return err
}

type NullableExtensionStatus struct {
	value *ExtensionStatus
	isSet bool
}

func (v NullableExtensionStatus) Get() *ExtensionStatus {
	return v.value
}

func (v *NullableExtensionStatus) Set(val *ExtensionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionStatus(val *ExtensionStatus) *NullableExtensionStatus {
	return &NullableExtensionStatus{value: val, isSet: true}
}

func (v NullableExtensionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


