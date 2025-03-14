# \StackAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllEntities**](StackAPI.md#GetAllEntities) | **Get** /api/v1/stacks | Attributes for all stacks
[**GetEntity**](StackAPI.md#GetEntity) | **Get** /api/v1/stacks/{type} | Get the attributes of a single stack



## GetAllEntities

> []interface{} GetAllEntities(ctx).Execute()

Attributes for all stacks



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
	resp, r, err := apiClient.StackAPI.GetAllEntities(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.GetAllEntities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllEntities`: []interface{}
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.GetAllEntities`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllEntitiesRequest struct via the builder pattern


### Return type

**[]interface{}**

### Authorization

[jwt_token](../README.md#jwt_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntity

> interface{} GetEntity(ctx, type_).Execute()

Get the attributes of a single stack



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
	type_ := openapiclient.StackType("Analytics") // StackType | the type of entity

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.GetEntity(context.Background(), type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.GetEntity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntity`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.GetEntity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | [**StackType**](.md) | the type of entity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[jwt_token](../README.md#jwt_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

