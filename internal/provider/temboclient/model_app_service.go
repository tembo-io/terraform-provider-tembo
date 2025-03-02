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

// checks if the AppService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppService{}

// AppService AppService significantly extends the functionality of your Tembo Postgres instance by running tools and software built by the Postgres open source community.  **Example**: This will configure and install a Postgrest container along side the Postgres instance, install pg_graphql extension, and configure the ingress routing to expose the Postgrest service.  ```yaml apiVersion: coredb.io/v1alpha1 kind: CoreDB metadata: name: test-db spec: trunk_installs: - name: pg_graphql version: 1.2.0 extensions: - name: pg_graphql locations: - database: postgres enabled: true  appServices: - name: postgrest image: postgrest/postgrest:v10.0.0 routing: # only expose /rest/v1 and /graphql/v1 - port: 3000 ingressPath: /rest/v1 middlewares: - my-headers - port: 3000 ingressPath: /graphql/v1 middlewares: - map-gql - my-headers middlewares: - customRequestHeaders: name: my-headers config: # removes auth header from request Authorization: \"\" Content-Profile: graphql Accept-Profile: graphql - stripPrefix: name: my-strip-prefix config: - /rest/v1 # reroute gql and rest requests - replacePathRegex: name: map-gql config: regex: /graphql/v1/? replacement: /rpc/resolve env: - name: PGRST_DB_URI valueFromPlatform: ReadWriteConnection - name: PGRST_DB_SCHEMA value: \"public, graphql\" - name: PGRST_DB_ANON_ROLE value: postgres - name: PGRST_LOG_LEVEL value: info ```
type AppService struct {
	// Defines the arguments to pass into the container if needed. You define this in the same manner as you would for all Kubernetes containers. See the [Kubernetes docs](https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container).
	Args []string `json:"args,omitempty"`
	// Defines the command into the container if needed. You define this in the same manner as you would for all Kubernetes containers. See the [Kubernetes docs](https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container).
	Command []string `json:"command,omitempty"`
	// Defines the environment variables to pass into the container if needed. You define this in the same manner as you would for all Kubernetes containers. See the [Kubernetes docs](https://kubernetes.io/docs/tasks/inject-data-application/define-environment-variable-container).
	Env []EnvVar `json:"env,omitempty"`
	// Defines the container image to use for the appService.
	Image string `json:"image"`
	Metrics NullableAppMetrics `json:"metrics,omitempty"`
	// Defines the ingress middeware configuration for the appService. This is specifically configured for the ingress controller Traefik.
	Middlewares []Middleware `json:"middlewares,omitempty"`
	// Defines the name of the appService.
	Name string `json:"name"`
	Probes NullableProbes `json:"probes,omitempty"`
	Resources *ResourceRequirements `json:"resources,omitempty"`
	// Defines the routing configuration for the appService.
	Routing []Routing `json:"routing,omitempty"`
	Storage NullableStorageConfig `json:"storage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppService AppService

// NewAppService instantiates a new AppService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppService(image string, name string) *AppService {
	this := AppService{}
	this.Image = image
	this.Name = name
	return &this
}

// NewAppServiceWithDefaults instantiates a new AppService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppServiceWithDefaults() *AppService {
	this := AppService{}
	return &this
}

// GetArgs returns the Args field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppService) GetArgs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppService) GetArgsOk() ([]string, bool) {
	if o == nil || IsNil(o.Args) {
		return nil, false
	}
	return o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *AppService) HasArgs() bool {
	if o != nil && !IsNil(o.Args) {
		return true
	}

	return false
}

// SetArgs gets a reference to the given []string and assigns it to the Args field.
func (o *AppService) SetArgs(v []string) {
	o.Args = v
}

// GetCommand returns the Command field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppService) GetCommand() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppService) GetCommandOk() ([]string, bool) {
	if o == nil || IsNil(o.Command) {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *AppService) HasCommand() bool {
	if o != nil && !IsNil(o.Command) {
		return true
	}

	return false
}

// SetCommand gets a reference to the given []string and assigns it to the Command field.
func (o *AppService) SetCommand(v []string) {
	o.Command = v
}

// GetEnv returns the Env field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppService) GetEnv() []EnvVar {
	if o == nil {
		var ret []EnvVar
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppService) GetEnvOk() ([]EnvVar, bool) {
	if o == nil || IsNil(o.Env) {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *AppService) HasEnv() bool {
	if o != nil && !IsNil(o.Env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given []EnvVar and assigns it to the Env field.
func (o *AppService) SetEnv(v []EnvVar) {
	o.Env = v
}

// GetImage returns the Image field value
func (o *AppService) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *AppService) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *AppService) SetImage(v string) {
	o.Image = v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppService) GetMetrics() AppMetrics {
	if o == nil || IsNil(o.Metrics.Get()) {
		var ret AppMetrics
		return ret
	}
	return *o.Metrics.Get()
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppService) GetMetricsOk() (*AppMetrics, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metrics.Get(), o.Metrics.IsSet()
}

// HasMetrics returns a boolean if a field has been set.
func (o *AppService) HasMetrics() bool {
	if o != nil && o.Metrics.IsSet() {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given NullableAppMetrics and assigns it to the Metrics field.
func (o *AppService) SetMetrics(v AppMetrics) {
	o.Metrics.Set(&v)
}
// SetMetricsNil sets the value for Metrics to be an explicit nil
func (o *AppService) SetMetricsNil() {
	o.Metrics.Set(nil)
}

// UnsetMetrics ensures that no value is present for Metrics, not even an explicit nil
func (o *AppService) UnsetMetrics() {
	o.Metrics.Unset()
}

// GetMiddlewares returns the Middlewares field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppService) GetMiddlewares() []Middleware {
	if o == nil {
		var ret []Middleware
		return ret
	}
	return o.Middlewares
}

// GetMiddlewaresOk returns a tuple with the Middlewares field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppService) GetMiddlewaresOk() ([]Middleware, bool) {
	if o == nil || IsNil(o.Middlewares) {
		return nil, false
	}
	return o.Middlewares, true
}

// HasMiddlewares returns a boolean if a field has been set.
func (o *AppService) HasMiddlewares() bool {
	if o != nil && !IsNil(o.Middlewares) {
		return true
	}

	return false
}

// SetMiddlewares gets a reference to the given []Middleware and assigns it to the Middlewares field.
func (o *AppService) SetMiddlewares(v []Middleware) {
	o.Middlewares = v
}

// GetName returns the Name field value
func (o *AppService) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AppService) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AppService) SetName(v string) {
	o.Name = v
}

// GetProbes returns the Probes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppService) GetProbes() Probes {
	if o == nil || IsNil(o.Probes.Get()) {
		var ret Probes
		return ret
	}
	return *o.Probes.Get()
}

// GetProbesOk returns a tuple with the Probes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppService) GetProbesOk() (*Probes, bool) {
	if o == nil {
		return nil, false
	}
	return o.Probes.Get(), o.Probes.IsSet()
}

// HasProbes returns a boolean if a field has been set.
func (o *AppService) HasProbes() bool {
	if o != nil && o.Probes.IsSet() {
		return true
	}

	return false
}

// SetProbes gets a reference to the given NullableProbes and assigns it to the Probes field.
func (o *AppService) SetProbes(v Probes) {
	o.Probes.Set(&v)
}
// SetProbesNil sets the value for Probes to be an explicit nil
func (o *AppService) SetProbesNil() {
	o.Probes.Set(nil)
}

// UnsetProbes ensures that no value is present for Probes, not even an explicit nil
func (o *AppService) UnsetProbes() {
	o.Probes.Unset()
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *AppService) GetResources() ResourceRequirements {
	if o == nil || IsNil(o.Resources) {
		var ret ResourceRequirements
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppService) GetResourcesOk() (*ResourceRequirements, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *AppService) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given ResourceRequirements and assigns it to the Resources field.
func (o *AppService) SetResources(v ResourceRequirements) {
	o.Resources = &v
}

// GetRouting returns the Routing field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppService) GetRouting() []Routing {
	if o == nil {
		var ret []Routing
		return ret
	}
	return o.Routing
}

// GetRoutingOk returns a tuple with the Routing field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppService) GetRoutingOk() ([]Routing, bool) {
	if o == nil || IsNil(o.Routing) {
		return nil, false
	}
	return o.Routing, true
}

// HasRouting returns a boolean if a field has been set.
func (o *AppService) HasRouting() bool {
	if o != nil && !IsNil(o.Routing) {
		return true
	}

	return false
}

// SetRouting gets a reference to the given []Routing and assigns it to the Routing field.
func (o *AppService) SetRouting(v []Routing) {
	o.Routing = v
}

// GetStorage returns the Storage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppService) GetStorage() StorageConfig {
	if o == nil || IsNil(o.Storage.Get()) {
		var ret StorageConfig
		return ret
	}
	return *o.Storage.Get()
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppService) GetStorageOk() (*StorageConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.Storage.Get(), o.Storage.IsSet()
}

// HasStorage returns a boolean if a field has been set.
func (o *AppService) HasStorage() bool {
	if o != nil && o.Storage.IsSet() {
		return true
	}

	return false
}

// SetStorage gets a reference to the given NullableStorageConfig and assigns it to the Storage field.
func (o *AppService) SetStorage(v StorageConfig) {
	o.Storage.Set(&v)
}
// SetStorageNil sets the value for Storage to be an explicit nil
func (o *AppService) SetStorageNil() {
	o.Storage.Set(nil)
}

// UnsetStorage ensures that no value is present for Storage, not even an explicit nil
func (o *AppService) UnsetStorage() {
	o.Storage.Unset()
}

func (o AppService) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Args != nil {
		toSerialize["args"] = o.Args
	}
	if o.Command != nil {
		toSerialize["command"] = o.Command
	}
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	toSerialize["image"] = o.Image
	if o.Metrics.IsSet() {
		toSerialize["metrics"] = o.Metrics.Get()
	}
	if o.Middlewares != nil {
		toSerialize["middlewares"] = o.Middlewares
	}
	toSerialize["name"] = o.Name
	if o.Probes.IsSet() {
		toSerialize["probes"] = o.Probes.Get()
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if o.Routing != nil {
		toSerialize["routing"] = o.Routing
	}
	if o.Storage.IsSet() {
		toSerialize["storage"] = o.Storage.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppService) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"image",
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

	varAppService := _AppService{}

	err = json.Unmarshal(data, &varAppService)

	if err != nil {
		return err
	}

	*o = AppService(varAppService)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "args")
		delete(additionalProperties, "command")
		delete(additionalProperties, "env")
		delete(additionalProperties, "image")
		delete(additionalProperties, "metrics")
		delete(additionalProperties, "middlewares")
		delete(additionalProperties, "name")
		delete(additionalProperties, "probes")
		delete(additionalProperties, "resources")
		delete(additionalProperties, "routing")
		delete(additionalProperties, "storage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppService struct {
	value *AppService
	isSet bool
}

func (v NullableAppService) Get() *AppService {
	return v.value
}

func (v *NullableAppService) Set(val *AppService) {
	v.value = val
	v.isSet = true
}

func (v NullableAppService) IsSet() bool {
	return v.isSet
}

func (v *NullableAppService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppService(val *AppService) *NullableAppService {
	return &NullableAppService{value: val, isSet: true}
}

func (v NullableAppService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


