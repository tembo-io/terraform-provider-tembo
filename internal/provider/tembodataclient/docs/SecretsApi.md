# \SecretsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSecret**](SecretsAPI.md#GetSecret) | **Get** /{namespace}/secrets/{secret_name} | Please use /api/v1/orgs/{org_id}/instances/{instance_id}/secrets/{secret_name}
[**GetSecretNames**](SecretsAPI.md#GetSecretNames) | **Get** /{namespace}/secrets | Please use /api/v1/orgs/{org_id}/instances/{instance_id}/secrets
[**GetSecretNamesV1**](SecretsAPI.md#GetSecretNamesV1) | **Get** /api/v1/orgs/{org_id}/instances/{instance_id}/secrets | 
[**GetSecretV1**](SecretsAPI.md#GetSecretV1) | **Get** /api/v1/orgs/{org_id}/instances/{instance_id}/secrets/{secret_name} | 



## GetSecret

> map[string]string GetSecret(ctx, namespace, secretName).Execute()

Please use /api/v1/orgs/{org_id}/instances/{instance_id}/secrets/{secret_name}



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tembo-io/terraform-provider-tembo/tembodataclient"
)

func main() {
	namespace := "org-myco-inst-prod" // string | Instance namespace
	secretName := "readonly-role" // string | Secret name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.GetSecret(context.Background(), namespace, secretName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.GetSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecret`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.GetSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Instance namespace | 
**secretName** | **string** | Secret name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]string**

### Authorization

[jwt_token](../README.md#jwt_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecretNames

> []AvailableSecret GetSecretNames(ctx, namespace).Execute()

Please use /api/v1/orgs/{org_id}/instances/{instance_id}/secrets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tembo-io/terraform-provider-tembo/tembodataclient"
)

func main() {
	namespace := "org-myco-inst-prod" // string | Instance namespace

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.GetSecretNames(context.Background(), namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.GetSecretNames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecretNames`: []AvailableSecret
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.GetSecretNames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Instance namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AvailableSecret**](AvailableSecret.md)

### Authorization

[jwt_token](../README.md#jwt_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecretNamesV1

> []AvailableSecret GetSecretNamesV1(ctx, orgId, instanceId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tembo-io/terraform-provider-tembo/tembodataclient"
)

func main() {
	orgId := "org_2T7FJA0DpaNBnELVLU1IS4XzZG0" // string | Tembo Cloud Organization ID
	instanceId := "inst_1696253936968_TblNOY_6" // string | Tembo Cloud Instance ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.GetSecretNamesV1(context.Background(), orgId, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.GetSecretNamesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecretNamesV1`: []AvailableSecret
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.GetSecretNamesV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Tembo Cloud Organization ID | 
**instanceId** | **string** | Tembo Cloud Instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretNamesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]AvailableSecret**](AvailableSecret.md)

### Authorization

[jwt_token](../README.md#jwt_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecretV1

> map[string]string GetSecretV1(ctx, orgId, instanceId, secretName).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tembo-io/terraform-provider-tembo/tembodataclient"
)

func main() {
	orgId := "org_2T7FJA0DpaNBnELVLU1IS4XzZG0" // string | Tembo Cloud Organization ID
	instanceId := "inst_1696253936968_TblNOY_6" // string | Tembo Cloud Instance ID
	secretName := "readonly-role" // string | Secret name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.GetSecretV1(context.Background(), orgId, instanceId, secretName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.GetSecretV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecretV1`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.GetSecretV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Tembo Cloud Organization ID | 
**instanceId** | **string** | Tembo Cloud Instance ID | 
**secretName** | **string** | Secret name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**map[string]string**

### Authorization

[jwt_token](../README.md#jwt_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

