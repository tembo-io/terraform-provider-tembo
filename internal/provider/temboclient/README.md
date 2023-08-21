# Go API client for temboclient

CoreDB backend API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import temboclient "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), temboclient.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), temboclient.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
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
*EntitiesApi* | [**GetAllEntities**](docs/EntitiesApi.md#getallentities) | **Get** /api/entities/all | Get the raw Stack definitions for all entities
*EntitiesApi* | [**GetEntity**](docs/EntitiesApi.md#getentity) | **Get** /api/entities/{entity_type} | Get the json-schema for an entity
*InstancesApi* | [**CreateInstance**](docs/InstancesApi.md#createinstance) | **Post** /api/orgs/{org_id}/instances/{entity_type} | Create a new instance of an entity
*InstancesApi* | [**DeleteInstance**](docs/InstancesApi.md#deleteinstance) | **Delete** /api/orgs/{org_id}/instances/{entity_type}/{instance_id} | Delete an existing instance of any entity_type
*InstancesApi* | [**GetAll**](docs/InstancesApi.md#getall) | **Get** /api/orgs/{org_id}/instances | Get all instances of all entities in an organization
*InstancesApi* | [**GetInstance**](docs/InstancesApi.md#getinstance) | **Get** /api/orgs/{org_id}/instances/{entity_type}/{instance_id} | Gets current attributes of an existing instance
*InstancesApi* | [**UpdateInstance**](docs/InstancesApi.md#updateinstance) | **Put** /api/orgs/{org_id}/instances/{entity_type}/{instance_id} | Update or make changes to an existing instance


## Documentation For Models

 - [ConnectionInfo](docs/ConnectionInfo.md)
 - [Cpu](docs/Cpu.md)
 - [CreateCluster](docs/CreateCluster.md)
 - [EntityType](docs/EntityType.md)
 - [Environment](docs/Environment.md)
 - [Extension](docs/Extension.md)
 - [ExtensionInstallLocation](docs/ExtensionInstallLocation.md)
 - [ExtensionInstallLocationStatus](docs/ExtensionInstallLocationStatus.md)
 - [ExtensionStatus](docs/ExtensionStatus.md)
 - [Memory](docs/Memory.md)
 - [ReadCluster](docs/ReadCluster.md)
 - [State](docs/State.md)
 - [Storage](docs/Storage.md)
 - [TrunkInstall](docs/TrunkInstall.md)
 - [TrunkInstallStatus](docs/TrunkInstallStatus.md)
 - [UpdateCluster](docs/UpdateCluster.md)
 - [UpdateEvent](docs/UpdateEvent.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### jwt_token

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
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


