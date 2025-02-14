# \RbacAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePolicy**](RbacAPI.md#DeletePolicy) | **Delete** /api/v1/orgs/{org_id}/policies | Delete a policy
[**GetActions**](RbacAPI.md#GetActions) | **Get** /api/v1/orgs/{org_id}/actions | Get all actions
[**GetPolicies**](RbacAPI.md#GetPolicies) | **Get** /api/v1/orgs/{org_id}/policies | Get all policies
[**GetRoles**](RbacAPI.md#GetRoles) | **Get** /api/v1/orgs/{org_id}/roles | Get all roles
[**SetPolicy**](RbacAPI.md#SetPolicy) | **Post** /api/v1/orgs/{org_id}/policies | Create or update a policy



## DeletePolicy

> string DeletePolicy(ctx, orgId).PolicyInput(policyInput).Execute()

Delete a policy



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
	orgId := "orgId_example" // string | Organization ID for the request
	policyInput := *openapiclient.NewPolicyInput("CreateInstance", true, "admin") // PolicyInput | Policy to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbacAPI.DeletePolicy(context.Background(), orgId).PolicyInput(policyInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.DeletePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePolicy`: string
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.DeletePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID for the request | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyInput** | [**PolicyInput**](PolicyInput.md) | Policy to delete | 

### Return type

**string**

### Authorization

[jwt_token](../README.md#jwt_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActions

> []Action GetActions(ctx, orgId).Execute()

Get all actions



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
	orgId := "orgId_example" // string | Organization ID for the request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbacAPI.GetActions(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.GetActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActions`: []Action
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.GetActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID for the request | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Action**](Action.md)

### Authorization

[jwt_token](../README.md#jwt_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicies

> []PolicyData GetPolicies(ctx, orgId).Execute()

Get all policies



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
	orgId := "orgId_example" // string | Organization ID for the request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbacAPI.GetPolicies(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.GetPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicies`: []PolicyData
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.GetPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID for the request | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PolicyData**](PolicyData.md)

### Authorization

[jwt_token](../README.md#jwt_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoles

> []Role GetRoles(ctx, orgId).Execute()

Get all roles



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
	orgId := "orgId_example" // string | Organization ID for the request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbacAPI.GetRoles(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.GetRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoles`: []Role
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.GetRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID for the request | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Role**](Role.md)

### Authorization

[jwt_token](../README.md#jwt_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPolicy

> string SetPolicy(ctx, orgId).PolicyInput(policyInput).Execute()

Create or update a policy



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
	orgId := "orgId_example" // string | Organization ID for the request
	policyInput := *openapiclient.NewPolicyInput("CreateInstance", true, "admin") // PolicyInput | Policy to set

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbacAPI.SetPolicy(context.Background(), orgId).PolicyInput(policyInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.SetPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPolicy`: string
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.SetPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID for the request | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyInput** | [**PolicyInput**](PolicyInput.md) | Policy to set | 

### Return type

**string**

### Authorization

[jwt_token](../README.md#jwt_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

