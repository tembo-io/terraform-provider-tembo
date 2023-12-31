# Go API client for tembodataclient

In the case of large or sensitive data, we avoid collecting it into Tembo Cloud. Instead, there is a Tembo Data API for each region, cloud, or private data plane.
            </br>
            </br>
            To find the Tembo Cloud API, please find it [here](https://api.tembo.io/swagger-ui/).
            

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v0.0.1
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
import tembodataclient "github.com/GIT_USER_ID/GIT_REPO_ID"
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
ctx := context.WithValue(context.Background(), tembodataclient.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), tembodataclient.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), tembodataclient.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), tembodataclient.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*MetricsApi* | [**QueryRange**](docs/MetricsApi.md#queryrange) | **Get** /{namespace}/metrics/query_range | 
*SecretsApi* | [**GetSecret**](docs/SecretsApi.md#getsecret) | **Get** /{namespace}/secrets/{secret_name} | Please use /api/v1/orgs/{org_id}/instances/{instance_id}/secrets/{secret_name}
*SecretsApi* | [**GetSecretNames**](docs/SecretsApi.md#getsecretnames) | **Get** /{namespace}/secrets | Please use /api/v1/orgs/{org_id}/instances/{instance_id}/secrets
*SecretsApi* | [**GetSecretNamesV1**](docs/SecretsApi.md#getsecretnamesv1) | **Get** /api/v1/orgs/{org_id}/instances/{instance_id}/secrets | 
*SecretsApi* | [**GetSecretV1**](docs/SecretsApi.md#getsecretv1) | **Get** /api/v1/orgs/{org_id}/instances/{instance_id}/secrets/{secret_name} | 


## Documentation For Models

 - [AvailableSecret](docs/AvailableSecret.md)


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



