# \MetricsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryRange**](MetricsApi.md#QueryRange) | **Get** /{namespace}/metrics/query_range | 



## QueryRange

> interface{} QueryRange(ctx, namespace).Query(query).Start(start).End(end).Step(step).Execute()



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
    namespace := "org-coredb-inst-control-plane-dev" // string | Instance namespace
    query := "(sum by (namespace) (max_over_time(pg_stat_activity_count{namespace="org-coredb-inst-control-plane-dev"}[1h])))" // string | PromQL range query, must include a 'namespace' label matching the query path
    start := int64(1686780828) // int64 | Range start, unix timestamp
    end := int64(1686862041) // int64 | Range end, unix timestamp. Default is now. (optional)
    step := "60s" // string | Step size duration string, defaults to 60s (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsApi.QueryRange(context.Background(), namespace).Query(query).Start(start).End(end).Step(step).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.QueryRange``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryRange`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `MetricsApi.QueryRange`: %v\n", resp)
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

