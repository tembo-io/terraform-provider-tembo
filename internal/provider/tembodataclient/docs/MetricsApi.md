# \MetricsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Query**](MetricsAPI.md#Query) | **Get** /{namespace}/metrics/query | 
[**QueryRange**](MetricsAPI.md#QueryRange) | **Get** /{namespace}/metrics/query_range | 



## Query

> interface{} Query(ctx, namespace).Query(query).Execute()



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
	namespace := "org-coredb-inst-control-plane-dev" // string | Instance namespace
	query := "(sum by (namespace) (max_over_time(pg_stat_activity_count{namespace="org-coredb-inst-control-plane-dev"}[1h])))" // string | PromQL range query, must include a 'namespace' label matching the query path

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.Query(context.Background(), namespace).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.Query``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Query`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.Query`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Instance namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **string** | PromQL range query, must include a &#39;namespace&#39; label matching the query path | 

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


## QueryRange

> interface{} QueryRange(ctx, namespace).Query(query).Start(start).End(end).Step(step).Execute()



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
	namespace := "org-coredb-inst-control-plane-dev" // string | Instance namespace
	query := "(sum by (namespace) (max_over_time(pg_stat_activity_count{namespace="org-coredb-inst-control-plane-dev"}[1h])))" // string | PromQL range query, must include a 'namespace' label matching the query path
	start := int64(1686780828) // int64 | Range start, unix timestamp
	end := int64(1686862041) // int64 | Range end, unix timestamp. Default is now. (optional)
	step := "60s" // string | Step size duration string, defaults to 60s (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.QueryRange(context.Background(), namespace).Query(query).Start(start).End(end).Step(step).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.QueryRange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryRange`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.QueryRange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Instance namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryRangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **string** | PromQL range query, must include a &#39;namespace&#39; label matching the query path | 
 **start** | **int64** | Range start, unix timestamp | 
 **end** | **int64** | Range end, unix timestamp. Default is now. | 
 **step** | **string** | Step size duration string, defaults to 60s | 

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

