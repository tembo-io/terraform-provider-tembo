# \InstanceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInstance**](InstanceApi.md#CreateInstance) | **Post** /api/v1/orgs/{org_id}/instances | Create a new Tembo instance
[**DeleteInstance**](InstanceApi.md#DeleteInstance) | **Delete** /api/v1/orgs/{org_id}/instances/{instance_id} | Delete an existing Tembo instance
[**GetAll**](InstanceApi.md#GetAll) | **Get** /api/v1/orgs/{org_id}/instances | Get all Tembo instances in an organization
[**GetInstance**](InstanceApi.md#GetInstance) | **Get** /api/v1/orgs/{org_id}/instances/{instance_id} | Get an existing Tembo instance
[**GetSchema**](InstanceApi.md#GetSchema) | **Get** /api/v1/orgs/instances/schema | Get the json-schema for an instance
[**InstanceEvent**](InstanceApi.md#InstanceEvent) | **Post** /api/v1/orgs/{org_id}/instances/{instance_id} | Lifecycle events for a Tembo instance
[**PatchInstance**](InstanceApi.md#PatchInstance) | **Patch** /api/v1/orgs/{org_id}/instances/{instance_id} | Update attributes on an existing Tembo instance
[**PutInstance**](InstanceApi.md#PutInstance) | **Put** /api/v1/orgs/{org_id}/instances/{instance_id} | Replace all attributes of an existing Tembo instance



## CreateInstance

> Instance CreateInstance(ctx, orgId).CreateInstance(createInstance).Execute()

Create a new Tembo instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    orgId := "orgId_example" // string | Organization ID that owns the Tembo instance
    createInstance := *openapiclient.NewCreateInstance(openapiclient.Cpu("1"), openapiclient.Environment("dev"), "InstanceName_example", openapiclient.Memory("1Gi"), openapiclient.StackType("Standard"), openapiclient.Storage("10Gi")) // CreateInstance | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstanceApi.CreateInstance(context.Background(), orgId).CreateInstance(createInstance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.CreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInstance`: Instance
    fmt.Fprintf(os.Stdout, "Response from `InstanceApi.CreateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID that owns the Tembo instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createInstance** | [**CreateInstance**](CreateInstance.md) |  | 

### Return type

[**Instance**](Instance.md)

### Authorization

[jwt_token](../README.md#jwt_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInstance

> Instance DeleteInstance(ctx, orgId, instanceId).Execute()

Delete an existing Tembo instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    orgId := "orgId_example" // string | Organization id of the instance to delete
    instanceId := "instanceId_example" // string | Delete this instance id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstanceApi.DeleteInstance(context.Background(), orgId, instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.DeleteInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteInstance`: Instance
    fmt.Fprintf(os.Stdout, "Response from `InstanceApi.DeleteInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization id of the instance to delete | 
**instanceId** | **string** | Delete this instance id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Instance**](Instance.md)

### Authorization

[jwt_token](../README.md#jwt_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAll

> []Instance GetAll(ctx, orgId).Execute()

Get all Tembo instances in an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    orgId := "orgId_example" // string | organization id for the request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstanceApi.GetAll(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.GetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAll`: []Instance
    fmt.Fprintf(os.Stdout, "Response from `InstanceApi.GetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization id for the request | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Instance**](Instance.md)

### Authorization

[jwt_token](../README.md#jwt_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstance

> Instance GetInstance(ctx, orgId, instanceId).Execute()

Get an existing Tembo instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    orgId := "orgId_example" // string | Organization ID that owns the instance
    instanceId := "instanceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstanceApi.GetInstance(context.Background(), orgId, instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.GetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstance`: Instance
    fmt.Fprintf(os.Stdout, "Response from `InstanceApi.GetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID that owns the instance | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Instance**](Instance.md)

### Authorization

[jwt_token](../README.md#jwt_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchema

> ErrorResponseSchema GetSchema(ctx).Execute()

Get the json-schema for an instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstanceApi.GetSchema(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.GetSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchema`: ErrorResponseSchema
    fmt.Fprintf(os.Stdout, "Response from `InstanceApi.GetSchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaRequest struct via the builder pattern


### Return type

[**ErrorResponseSchema**](ErrorResponseSchema.md)

### Authorization

[jwt_token](../README.md#jwt_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstanceEvent

> Instance InstanceEvent(ctx, orgId, instanceId).EventType(eventType).Execute()

Lifecycle events for a Tembo instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    orgId := "orgId_example" // string | Organization ID that owns the Tembo instance
    eventType := openapiclient.InstanceEvent("restart") // InstanceEvent | 
    instanceId := "instanceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstanceApi.InstanceEvent(context.Background(), orgId, instanceId).EventType(eventType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.InstanceEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstanceEvent`: Instance
    fmt.Fprintf(os.Stdout, "Response from `InstanceApi.InstanceEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID that owns the Tembo instance | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstanceEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventType** | [**InstanceEvent**](InstanceEvent.md) |  | 


### Return type

[**Instance**](Instance.md)

### Authorization

[jwt_token](../README.md#jwt_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchInstance

> Instance PatchInstance(ctx, orgId, instanceId).PatchInstance(patchInstance).Execute()

Update attributes on an existing Tembo instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    orgId := "orgId_example" // string | Organization ID that owns the instance
    instanceId := "instanceId_example" // string | 
    patchInstance := *openapiclient.NewPatchInstance() // PatchInstance | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstanceApi.PatchInstance(context.Background(), orgId, instanceId).PatchInstance(patchInstance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.PatchInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchInstance`: Instance
    fmt.Fprintf(os.Stdout, "Response from `InstanceApi.PatchInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID that owns the instance | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchInstance** | [**PatchInstance**](PatchInstance.md) |  | 

### Return type

[**Instance**](Instance.md)

### Authorization

[jwt_token](../README.md#jwt_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutInstance

> Instance PutInstance(ctx, orgId, instanceId).UpdateInstance(updateInstance).Execute()

Replace all attributes of an existing Tembo instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    orgId := "orgId_example" // string | Organization ID that owns the Tembo instance
    instanceId := "instanceId_example" // string | 
    updateInstance := *openapiclient.NewUpdateInstance(openapiclient.Cpu("1"), openapiclient.Environment("dev"), openapiclient.Memory("1Gi"), int32(123), openapiclient.Storage("10Gi")) // UpdateInstance | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstanceApi.PutInstance(context.Background(), orgId, instanceId).UpdateInstance(updateInstance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.PutInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutInstance`: Instance
    fmt.Fprintf(os.Stdout, "Response from `InstanceApi.PutInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID that owns the Tembo instance | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateInstance** | [**UpdateInstance**](UpdateInstance.md) |  | 

### Return type

[**Instance**](Instance.md)

### Authorization

[jwt_token](../README.md#jwt_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

