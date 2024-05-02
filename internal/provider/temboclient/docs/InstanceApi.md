# \InstanceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInstance**](InstanceAPI.md#CreateInstance) | **Post** /api/v1/orgs/{org_id}/instances | Create a new Tembo instance
[**DeleteInstance**](InstanceAPI.md#DeleteInstance) | **Delete** /api/v1/orgs/{org_id}/instances/{instance_id} | Delete an existing Tembo instance
[**GetAll**](InstanceAPI.md#GetAll) | **Get** /api/v1/orgs/{org_id}/instances | Get all Tembo instances in an organization
[**GetInstance**](InstanceAPI.md#GetInstance) | **Get** /api/v1/orgs/{org_id}/instances/{instance_id} | Get an existing Tembo instance
[**GetSchema**](InstanceAPI.md#GetSchema) | **Get** /api/v1/orgs/instances/schema | Get the json-schema for an instance
[**InstanceEvent**](InstanceAPI.md#InstanceEvent) | **Post** /api/v1/orgs/{org_id}/instances/{instance_id} | Lifecycle events for a Tembo instance
[**PatchInstance**](InstanceAPI.md#PatchInstance) | **Patch** /api/v1/orgs/{org_id}/instances/{instance_id} | Update attributes on an existing Tembo instance
[**RestoreInstance**](InstanceAPI.md#RestoreInstance) | **Post** /api/v1/orgs/{org_id}/restore | Restore a Tembo instance



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
	openapiclient "github.com/tembo-io/terraform-provider-tembo/temboclient"
)

func main() {
	orgId := "orgId_example" // string | Organization ID that owns the Tembo instance
	createInstance := *openapiclient.NewCreateInstance(openapiclient.Cpu("0.25"), openapiclient.Environment("dev"), "InstanceName_example", openapiclient.Memory("1Gi"), openapiclient.StackType("Standard"), openapiclient.Storage("10Gi")) // CreateInstance | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.CreateInstance(context.Background(), orgId).CreateInstance(createInstance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.CreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInstance`: Instance
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.CreateInstance`: %v\n", resp)
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
	openapiclient "github.com/tembo-io/terraform-provider-tembo/temboclient"
)

func main() {
	orgId := "orgId_example" // string | Organization id of the instance to delete
	instanceId := "instanceId_example" // string | Delete this instance id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.DeleteInstance(context.Background(), orgId, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.DeleteInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteInstance`: Instance
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.DeleteInstance`: %v\n", resp)
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
	openapiclient "github.com/tembo-io/terraform-provider-tembo/temboclient"
)

func main() {
	orgId := "orgId_example" // string | organization id for the request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.GetAll(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.GetAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAll`: []Instance
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.GetAll`: %v\n", resp)
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
	openapiclient "github.com/tembo-io/terraform-provider-tembo/temboclient"
)

func main() {
	orgId := "orgId_example" // string | Organization ID that owns the instance
	instanceId := "instanceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.GetInstance(context.Background(), orgId, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.GetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstance`: Instance
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.GetInstance`: %v\n", resp)
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
	openapiclient "github.com/tembo-io/terraform-provider-tembo/temboclient"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.GetSchema(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.GetSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSchema`: ErrorResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.GetSchema`: %v\n", resp)
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
	openapiclient "github.com/tembo-io/terraform-provider-tembo/temboclient"
)

func main() {
	orgId := "orgId_example" // string | Organization ID that owns the Tembo instance
	eventType := openapiclient.InstanceEvent("restart") // InstanceEvent | 
	instanceId := "instanceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.InstanceEvent(context.Background(), orgId, instanceId).EventType(eventType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.InstanceEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstanceEvent`: Instance
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.InstanceEvent`: %v\n", resp)
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
	openapiclient "github.com/tembo-io/terraform-provider-tembo/temboclient"
)

func main() {
	orgId := "orgId_example" // string | Organization ID that owns the instance
	instanceId := "instanceId_example" // string | 
	patchInstance := *openapiclient.NewPatchInstance() // PatchInstance | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.PatchInstance(context.Background(), orgId, instanceId).PatchInstance(patchInstance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.PatchInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchInstance`: Instance
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.PatchInstance`: %v\n", resp)
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


## RestoreInstance

> Instance RestoreInstance(ctx, orgId).RestoreInstance(restoreInstance).Execute()

Restore a Tembo instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tembo-io/terraform-provider-tembo/temboclient"
)

func main() {
	orgId := "orgId_example" // string | Organization ID that owns the Tembo instance
	restoreInstance := *openapiclient.NewRestoreInstance("InstanceName_example", *openapiclient.NewRestore("InstanceId_example")) // RestoreInstance | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.RestoreInstance(context.Background(), orgId).RestoreInstance(restoreInstance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.RestoreInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreInstance`: Instance
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.RestoreInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID that owns the Tembo instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restoreInstance** | [**RestoreInstance**](RestoreInstance.md) |  | 

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

