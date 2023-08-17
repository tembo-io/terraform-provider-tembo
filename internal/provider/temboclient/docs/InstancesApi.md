# \InstancesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInstance**](InstancesApi.md#CreateInstance) | **Post** /api/orgs/{org_id}/instances/{entity_type} | Create a new instance of an entity
[**DeleteInstance**](InstancesApi.md#DeleteInstance) | **Delete** /api/orgs/{org_id}/instances/{entity_type}/{instance_id} | Delete an existing instance of any entity_type
[**GetAll**](InstancesApi.md#GetAll) | **Get** /api/orgs/{org_id}/instances | Get all instances of all entities in an organization
[**GetInstance**](InstancesApi.md#GetInstance) | **Get** /api/orgs/{org_id}/instances/{entity_type}/{instance_id} | Gets current attributes of an existing instance
[**UpdateInstance**](InstancesApi.md#UpdateInstance) | **Put** /api/orgs/{org_id}/instances/{entity_type}/{instance_id} | Update or make changes to an existing instance



## CreateInstance

> ReadCluster CreateInstance(ctx, orgId, entityType).CreateCluster(createCluster).Execute()

Create a new instance of an entity



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
    orgId := "orgId_example" // string | organization id that the instance belongs to
    entityType := openapiclient.EntityType("Standard") // EntityType | the type of entity for this instance
    createCluster := *openapiclient.NewCreateCluster(openapiclient.Cpu("1"), openapiclient.Environment("dev"), "InstanceName_example", openapiclient.Memory("1Gi"), openapiclient.Storage("10Gi")) // CreateCluster | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstancesApi.CreateInstance(context.Background(), orgId, entityType).CreateCluster(createCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.CreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInstance`: ReadCluster
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.CreateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | organization id that the instance belongs to | 
**entityType** | [**EntityType**](.md) | the type of entity for this instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createCluster** | [**CreateCluster**](CreateCluster.md) |  | 

### Return type

[**ReadCluster**](ReadCluster.md)

### Authorization

[jwt_token](../README.md#jwt_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInstance

> DeleteInstance(ctx, orgId, entityType, instanceId).Execute()

Delete an existing instance of any entity_type



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
    entityType := "entityType_example" // string | Type of entity for this request
    instanceId := "instanceId_example" // string | Unique identifier of the instance to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.InstancesApi.DeleteInstance(context.Background(), orgId, entityType, instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.DeleteInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization id of the instance to delete | 
**entityType** | **string** | Type of entity for this request | 
**instanceId** | **string** | Unique identifier of the instance to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[jwt_token](../README.md#jwt_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAll

> []ReadCluster GetAll(ctx, orgId).Execute()

Get all instances of all entities in an organization



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
    resp, r, err := apiClient.InstancesApi.GetAll(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.GetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAll`: []ReadCluster
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.GetAll`: %v\n", resp)
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

[**[]ReadCluster**](ReadCluster.md)

### Authorization

[jwt_token](../README.md#jwt_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstance

> ReadCluster GetInstance(ctx, orgId, entityType, instanceId).Execute()

Gets current attributes of an existing instance



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
    orgId := "orgId_example" // string | Organization id of the instance to update
    entityType := "entityType_example" // string | Type of entity for this request
    instanceId := "instanceId_example" // string | Unique identifier of the instance to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstancesApi.GetInstance(context.Background(), orgId, entityType, instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.GetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstance`: ReadCluster
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.GetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization id of the instance to update | 
**entityType** | **string** | Type of entity for this request | 
**instanceId** | **string** | Unique identifier of the instance to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ReadCluster**](ReadCluster.md)

### Authorization

[jwt_token](../README.md#jwt_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInstance

> UpdateInstance(ctx, orgId, entityType, instanceId).EventType(eventType).UpdateCluster(updateCluster).Execute()

Update or make changes to an existing instance



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
    orgId := "orgId_example" // string | Organization id of the instance to update
    entityType := "entityType_example" // string | Type of entity for this request
    instanceId := "instanceId_example" // string | Unique identifier of the instance to update
    eventType := openapiclient.UpdateEvent("extension") // UpdateEvent | 
    updateCluster := *openapiclient.NewUpdateCluster() // UpdateCluster | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.InstancesApi.UpdateInstance(context.Background(), orgId, entityType, instanceId).EventType(eventType).UpdateCluster(updateCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.UpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization id of the instance to update | 
**entityType** | **string** | Type of entity for this request | 
**instanceId** | **string** | Unique identifier of the instance to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **eventType** | [**UpdateEvent**](UpdateEvent.md) |  | 
 **updateCluster** | [**UpdateCluster**](UpdateCluster.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt_token](../README.md#jwt_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

