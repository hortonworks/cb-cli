# \AuthorizationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizationInfo**](AuthorizationApi.md#AuthorizationInfo) | **Get** /authorization/info | list of required permissions for APIs



## AuthorizationInfo

> []ApiAuthorizationInfo AuthorizationInfo(ctx).Execute()

list of required permissions for APIs

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationApi.AuthorizationInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.AuthorizationInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthorizationInfo`: []ApiAuthorizationInfo
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.AuthorizationInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizationInfoRequest struct via the builder pattern


### Return type

[**[]ApiAuthorizationInfo**](ApiAuthorizationInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

