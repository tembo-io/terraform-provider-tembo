# \TokensAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvalidateToken**](TokensAPI.md#InvalidateToken) | **Delete** /api/v1/tokens/invalidate | Invalidate the token included in the Authorization header



## InvalidateToken

> string InvalidateToken(ctx).Execute()

Invalidate the token included in the Authorization header



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
	resp, r, err := apiClient.TokensAPI.InvalidateToken(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.InvalidateToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvalidateToken`: string
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.InvalidateToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInvalidateTokenRequest struct via the builder pattern


### Return type

**string**

### Authorization

[jwt_token](../README.md#jwt_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

