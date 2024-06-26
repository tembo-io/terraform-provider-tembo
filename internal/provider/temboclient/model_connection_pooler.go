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

// checks if the ConnectionPooler type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionPooler{}

// ConnectionPooler A connection pooler is a tool used to manage database connections, sitting between your application and Postgres instance. Because of the way Postgres handles connections, the server may encounter resource constraint issues when managing a few thousand connections. Using a pooler can alleviate these issues by using actual Postgres connections only when necessary  **Example**: A typical connection pooler configuration  ```yaml apiVersion: coredb.io/v1alpha1 kind: CoreDB metadata: name: test-db spec: connectionPooler: enabled: true pooler: poolMode: transaction # Valid parameter values can be found at https://www.pgbouncer.org/config.html parameters: default_pool_size: \"50\" max_client_conn: \"5000\" resources: limits: cpu: 200m memory: 256Mi requests: cpu: 100m memory: 128Mi ```
type ConnectionPooler struct {
	// Enable the connection pooler  **Default**: false.
	Enabled *bool `json:"enabled,omitempty"`
	Pooler *PgBouncer `json:"pooler,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectionPooler ConnectionPooler

// NewConnectionPooler instantiates a new ConnectionPooler object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionPooler() *ConnectionPooler {
	this := ConnectionPooler{}
	return &this
}

// NewConnectionPoolerWithDefaults instantiates a new ConnectionPooler object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionPoolerWithDefaults() *ConnectionPooler {
	this := ConnectionPooler{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ConnectionPooler) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionPooler) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ConnectionPooler) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ConnectionPooler) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPooler returns the Pooler field value if set, zero value otherwise.
func (o *ConnectionPooler) GetPooler() PgBouncer {
	if o == nil || IsNil(o.Pooler) {
		var ret PgBouncer
		return ret
	}
	return *o.Pooler
}

// GetPoolerOk returns a tuple with the Pooler field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionPooler) GetPoolerOk() (*PgBouncer, bool) {
	if o == nil || IsNil(o.Pooler) {
		return nil, false
	}
	return o.Pooler, true
}

// HasPooler returns a boolean if a field has been set.
func (o *ConnectionPooler) HasPooler() bool {
	if o != nil && !IsNil(o.Pooler) {
		return true
	}

	return false
}

// SetPooler gets a reference to the given PgBouncer and assigns it to the Pooler field.
func (o *ConnectionPooler) SetPooler(v PgBouncer) {
	o.Pooler = &v
}

func (o ConnectionPooler) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionPooler) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Pooler) {
		toSerialize["pooler"] = o.Pooler
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConnectionPooler) UnmarshalJSON(data []byte) (err error) {
	varConnectionPooler := _ConnectionPooler{}

	err = json.Unmarshal(data, &varConnectionPooler)

	if err != nil {
		return err
	}

	*o = ConnectionPooler(varConnectionPooler)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "pooler")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectionPooler struct {
	value *ConnectionPooler
	isSet bool
}

func (v NullableConnectionPooler) Get() *ConnectionPooler {
	return v.value
}

func (v *NullableConnectionPooler) Set(val *ConnectionPooler) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionPooler) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionPooler) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionPooler(val *ConnectionPooler) *NullableConnectionPooler {
	return &NullableConnectionPooler{value: val, isSet: true}
}

func (v NullableConnectionPooler) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionPooler) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


