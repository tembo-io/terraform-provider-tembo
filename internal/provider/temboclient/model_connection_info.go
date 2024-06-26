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

// checks if the ConnectionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionInfo{}

// ConnectionInfo struct for ConnectionInfo
type ConnectionInfo struct {
	Host string `json:"host"`
	PoolerHost NullableString `json:"pooler_host,omitempty"`
	Port int32 `json:"port"`
	User string `json:"user"`
	AdditionalProperties map[string]interface{}
}

type _ConnectionInfo ConnectionInfo

// NewConnectionInfo instantiates a new ConnectionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionInfo(host string, port int32, user string) *ConnectionInfo {
	this := ConnectionInfo{}
	this.Host = host
	this.Port = port
	this.User = user
	return &this
}

// NewConnectionInfoWithDefaults instantiates a new ConnectionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionInfoWithDefaults() *ConnectionInfo {
	this := ConnectionInfo{}
	return &this
}

// GetHost returns the Host field value
func (o *ConnectionInfo) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *ConnectionInfo) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *ConnectionInfo) SetHost(v string) {
	o.Host = v
}

// GetPoolerHost returns the PoolerHost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectionInfo) GetPoolerHost() string {
	if o == nil || IsNil(o.PoolerHost.Get()) {
		var ret string
		return ret
	}
	return *o.PoolerHost.Get()
}

// GetPoolerHostOk returns a tuple with the PoolerHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionInfo) GetPoolerHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PoolerHost.Get(), o.PoolerHost.IsSet()
}

// HasPoolerHost returns a boolean if a field has been set.
func (o *ConnectionInfo) HasPoolerHost() bool {
	if o != nil && o.PoolerHost.IsSet() {
		return true
	}

	return false
}

// SetPoolerHost gets a reference to the given NullableString and assigns it to the PoolerHost field.
func (o *ConnectionInfo) SetPoolerHost(v string) {
	o.PoolerHost.Set(&v)
}
// SetPoolerHostNil sets the value for PoolerHost to be an explicit nil
func (o *ConnectionInfo) SetPoolerHostNil() {
	o.PoolerHost.Set(nil)
}

// UnsetPoolerHost ensures that no value is present for PoolerHost, not even an explicit nil
func (o *ConnectionInfo) UnsetPoolerHost() {
	o.PoolerHost.Unset()
}

// GetPort returns the Port field value
func (o *ConnectionInfo) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *ConnectionInfo) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *ConnectionInfo) SetPort(v int32) {
	o.Port = v
}

// GetUser returns the User field value
func (o *ConnectionInfo) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *ConnectionInfo) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *ConnectionInfo) SetUser(v string) {
	o.User = v
}

func (o ConnectionInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["host"] = o.Host
	if o.PoolerHost.IsSet() {
		toSerialize["pooler_host"] = o.PoolerHost.Get()
	}
	toSerialize["port"] = o.Port
	toSerialize["user"] = o.User

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConnectionInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"host",
		"port",
		"user",
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

	varConnectionInfo := _ConnectionInfo{}

	err = json.Unmarshal(data, &varConnectionInfo)

	if err != nil {
		return err
	}

	*o = ConnectionInfo(varConnectionInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "host")
		delete(additionalProperties, "pooler_host")
		delete(additionalProperties, "port")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectionInfo struct {
	value *ConnectionInfo
	isSet bool
}

func (v NullableConnectionInfo) Get() *ConnectionInfo {
	return v.value
}

func (v *NullableConnectionInfo) Set(val *ConnectionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionInfo(val *ConnectionInfo) *NullableConnectionInfo {
	return &NullableConnectionInfo{value: val, isSet: true}
}

func (v NullableConnectionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


