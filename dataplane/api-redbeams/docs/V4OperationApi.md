# \V4OperationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRedbeamsOperationProgressByResourceCrn**](V4OperationApi.md#GetRedbeamsOperationProgressByResourceCrn) | **Get** /v4/operation/resource/crn/{resourceCrn} | Get flow operation progress details for resource by resource crn



## GetRedbeamsOperationProgressByResourceCrn

> OperationView GetRedbeamsOperationProgressByResourceCrn(ctx, resourceCrn).Detailed(detailed).Execute()

Get flow operation progress details for resource by resource crn



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
    resourceCrn := "resourceCrn_example" // string | 
    detailed := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.V4OperationApi.GetRedbeamsOperationProgressByResourceCrn(context.Background(), resourceCrn).Detailed(detailed).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V4OperationApi.GetRedbeamsOperationProgressByResourceCrn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRedbeamsOperationProgressByResourceCrn`: OperationView
    fmt.Fprintf(os.Stdout, "Response from `V4OperationApi.GetRedbeamsOperationProgressByResourceCrn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceCrn** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedbeamsOperationProgressByResourceCrnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **detailed** | **bool** |  | [default to false]

### Return type

[**OperationView**](OperationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

