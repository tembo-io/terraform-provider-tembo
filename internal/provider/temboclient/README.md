# Go API client for temboclient

Platform API for Tembo Cloud
            </br>
            </br>
            To find a Tembo Data API, please find it here:
            </br>
            </br>
            [AWS US East 1](https://api.data-1.use1.tembo.io/swagger-ui/)
            

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1.0.0
- Package version: 1.0.0
- Generator version: 7.12.0-SNAPSHOT
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import temboclient "github.com/tembo-io/terraform-provider-tembo/temboclient"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `temboclient.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), temboclient.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `temboclient.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), temboclient.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `temboclient.ContextOperationServerIndices` and `temboclient.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), temboclient.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), temboclient.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AiAPI* | [**Ask**](docs/AiAPI.md#ask) | **Get** /api/v1/ask/ | Ask a question to Tembo Docs
*AiAPI* | [**Search**](docs/AiAPI.md#search) | **Get** /api/v1/search/ | Search Tembo Docs Database
*AppAPI* | [**GetAllApps**](docs/AppAPI.md#getallapps) | **Get** /api/v1/apps | Attributes for all apps
*AppAPI* | [**GetApp**](docs/AppAPI.md#getapp) | **Get** /api/v1/apps/{type} | Get the attributes of a single App
*DataplaneAPI* | [**GetAllDataplanes**](docs/DataplaneAPI.md#getalldataplanes) | **Get** /api/v1/orgs/{org_id}/dataplanes | 
*InstanceAPI* | [**CreateInstance**](docs/InstanceAPI.md#createinstance) | **Post** /api/v1/orgs/{org_id}/instances | Create a new Tembo instance
*InstanceAPI* | [**DeleteInstance**](docs/InstanceAPI.md#deleteinstance) | **Delete** /api/v1/orgs/{org_id}/instances/{instance_id} | Delete an existing Tembo instance
*InstanceAPI* | [**GetAll**](docs/InstanceAPI.md#getall) | **Get** /api/v1/orgs/{org_id}/instances | Get all Tembo instances in an organization
*InstanceAPI* | [**GetInstance**](docs/InstanceAPI.md#getinstance) | **Get** /api/v1/orgs/{org_id}/instances/{instance_id} | Get an existing Tembo instance
*InstanceAPI* | [**GetInstanceHistory**](docs/InstanceAPI.md#getinstancehistory) | **Get** /api/v1/orgs/{org_id}/instances/{instance_id}/history/{field} | Get historical information for a given instance
*InstanceAPI* | [**GetInstanceToml**](docs/InstanceAPI.md#getinstancetoml) | **Get** /api/v1/orgs/{org_id}/instances/{instance_id}/toml | Generate text for Tembo.toml
*InstanceAPI* | [**GetSchema**](docs/InstanceAPI.md#getschema) | **Get** /api/v1/orgs/instances/schema | Get the json-schema for an instance
*InstanceAPI* | [**InstanceEvent**](docs/InstanceAPI.md#instanceevent) | **Post** /api/v1/orgs/{org_id}/instances/{instance_id} | Lifecycle events for a Tembo instance
*InstanceAPI* | [**PatchInstance**](docs/InstanceAPI.md#patchinstance) | **Patch** /api/v1/orgs/{org_id}/instances/{instance_id} | Update attributes on an existing Tembo instance
*InstanceAPI* | [**RestoreInstance**](docs/InstanceAPI.md#restoreinstance) | **Post** /api/v1/orgs/{org_id}/restore | Restore a Tembo instance
*RbacAPI* | [**DeletePolicy**](docs/RbacAPI.md#deletepolicy) | **Delete** /api/v1/orgs/{org_id}/policies | Delete a policy
*RbacAPI* | [**GetActions**](docs/RbacAPI.md#getactions) | **Get** /api/v1/orgs/{org_id}/actions | Get all actions
*RbacAPI* | [**GetPolicies**](docs/RbacAPI.md#getpolicies) | **Get** /api/v1/orgs/{org_id}/policies | Get all policies
*RbacAPI* | [**GetRoles**](docs/RbacAPI.md#getroles) | **Get** /api/v1/orgs/{org_id}/roles | Get all roles
*RbacAPI* | [**SetPolicy**](docs/RbacAPI.md#setpolicy) | **Post** /api/v1/orgs/{org_id}/policies | Create or update a policy
*StackAPI* | [**GetAllEntities**](docs/StackAPI.md#getallentities) | **Get** /api/v1/stacks | Attributes for all stacks
*StackAPI* | [**GetEntity**](docs/StackAPI.md#getentity) | **Get** /api/v1/stacks/{type} | Get the attributes of a single stack
*TokensAPI* | [**InvalidateToken**](docs/TokensAPI.md#invalidatetoken) | **Delete** /api/v1/tokens/invalidate | Invalidate the token included in the Authorization header


## Documentation For Models

 - [Action](docs/Action.md)
 - [AppConfig](docs/AppConfig.md)
 - [AppMetrics](docs/AppMetrics.md)
 - [AppService](docs/AppService.md)
 - [AppType](docs/AppType.md)
 - [AppTypeOneOf](docs/AppTypeOneOf.md)
 - [AppTypeOneOf1](docs/AppTypeOneOf1.md)
 - [AppTypeOneOf2](docs/AppTypeOneOf2.md)
 - [AppTypeOneOf3](docs/AppTypeOneOf3.md)
 - [AppTypeOneOf4](docs/AppTypeOneOf4.md)
 - [AppTypeOneOf5](docs/AppTypeOneOf5.md)
 - [AppTypeOneOf6](docs/AppTypeOneOf6.md)
 - [AskParams](docs/AskParams.md)
 - [AskResult](docs/AskResult.md)
 - [AutoStop](docs/AutoStop.md)
 - [Autoscaling](docs/Autoscaling.md)
 - [AutoscalingStorage](docs/AutoscalingStorage.md)
 - [ConnectionInfo](docs/ConnectionInfo.md)
 - [ConnectionPooler](docs/ConnectionPooler.md)
 - [Cpu](docs/Cpu.md)
 - [CreateInstance](docs/CreateInstance.md)
 - [DataPlane](docs/DataPlane.md)
 - [DedicatedNetworking](docs/DedicatedNetworking.md)
 - [EnvVar](docs/EnvVar.md)
 - [EnvVarRef](docs/EnvVarRef.md)
 - [Environment](docs/Environment.md)
 - [ErrorResponseSchema](docs/ErrorResponseSchema.md)
 - [Experimental](docs/Experimental.md)
 - [Extension](docs/Extension.md)
 - [ExtensionInstallLocation](docs/ExtensionInstallLocation.md)
 - [ExtensionInstallLocationStatus](docs/ExtensionInstallLocationStatus.md)
 - [ExtensionStatus](docs/ExtensionStatus.md)
 - [HeaderConfig](docs/HeaderConfig.md)
 - [HistoryEntry](docs/HistoryEntry.md)
 - [HistoryPage](docs/HistoryPage.md)
 - [Infrastructure](docs/Infrastructure.md)
 - [Ingress](docs/Ingress.md)
 - [IngressType](docs/IngressType.md)
 - [Instance](docs/Instance.md)
 - [InstanceEvent](docs/InstanceEvent.md)
 - [InstanceState](docs/InstanceState.md)
 - [IntOrString](docs/IntOrString.md)
 - [IntOrStringOneOf](docs/IntOrStringOneOf.md)
 - [IntOrStringOneOf1](docs/IntOrStringOneOf1.md)
 - [Memory](docs/Memory.md)
 - [Middleware](docs/Middleware.md)
 - [MiddlewareOneOf](docs/MiddlewareOneOf.md)
 - [MiddlewareOneOf1](docs/MiddlewareOneOf1.md)
 - [MiddlewareOneOf2](docs/MiddlewareOneOf2.md)
 - [PaginationInfo](docs/PaginationInfo.md)
 - [PatchAutoscaling](docs/PatchAutoscaling.md)
 - [PatchInstance](docs/PatchInstance.md)
 - [PgBouncer](docs/PgBouncer.md)
 - [PgConfig](docs/PgConfig.md)
 - [PolicyData](docs/PolicyData.md)
 - [PolicyInput](docs/PolicyInput.md)
 - [PoolerPgbouncerPoolMode](docs/PoolerPgbouncerPoolMode.md)
 - [PoolerTemplateSpecContainersResources](docs/PoolerTemplateSpecContainersResources.md)
 - [PoolerTemplateSpecContainersResourcesClaims](docs/PoolerTemplateSpecContainersResourcesClaims.md)
 - [Probe](docs/Probe.md)
 - [Probes](docs/Probes.md)
 - [ReplacePathRegexConfig](docs/ReplacePathRegexConfig.md)
 - [ReplacePathRegexConfigType](docs/ReplacePathRegexConfigType.md)
 - [Resource](docs/Resource.md)
 - [ResourceRequirements](docs/ResourceRequirements.md)
 - [Restore](docs/Restore.md)
 - [RestoreInstance](docs/RestoreInstance.md)
 - [Role](docs/Role.md)
 - [Routing](docs/Routing.md)
 - [SearchParams](docs/SearchParams.md)
 - [SearchResponse](docs/SearchResponse.md)
 - [StackType](docs/StackType.md)
 - [State](docs/State.md)
 - [Storage](docs/Storage.md)
 - [StorageConfig](docs/StorageConfig.md)
 - [StripPrefixConfig](docs/StripPrefixConfig.md)
 - [TrunkInstall](docs/TrunkInstall.md)
 - [TrunkInstallStatus](docs/TrunkInstallStatus.md)
 - [VolumeMount](docs/VolumeMount.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### jwt_token

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), temboclient.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



