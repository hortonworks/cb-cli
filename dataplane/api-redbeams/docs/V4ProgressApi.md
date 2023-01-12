# \V4ProgressApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRedbeamsFlowLogsProgressByResourceCrn**](V4ProgressApi.md#GetRedbeamsFlowLogsProgressByResourceCrn) | **Get** /v4/progress/resource/crn/{resourceCrn} | List recent flow operations progress details for resource by resource crn
[**GetRedbeamsLastFlowLogProgressByResourceCrn**](V4ProgressApi.md#GetRedbeamsLastFlowLogProgressByResourceCrn) | **Get** /v4/progress/resource/crn/{resourceCrn}/last | Get last flow operation progress details for resource by resource crn



## GetRedbeamsFlowLogsProgressByResourceCrn

> []FlowProgressResponse GetRedbeamsFlowLogsProgressByResourceCrn(ctx, resourceCrn).Execute()

List recent flow operations progress details for resource by resource crn



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.V4ProgressApi.GetRedbeamsFlowLogsProgressByResourceCrn(context.Background(), resourceCrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V4ProgressApi.GetRedbeamsFlowLogsProgressByResourceCrn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRedbeamsFlowLogsProgressByResourceCrn`: []FlowProgressResponse
    fmt.Fprintf(os.Stdout, "Response from `V4ProgressApi.GetRedbeamsFlowLogsProgressByResourceCrn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceCrn** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedbeamsFlowLogsProgressByResourceCrnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FlowProgressResponse**](FlowProgressResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedbeamsLastFlowLogProgressByResourceCrn

> FlowProgressResponse GetRedbeamsLastFlowLogProgressByResourceCrn(ctx, resourceCrn).Execute()

Get last flow operation progress details for resource by resource crn



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.V4ProgressApi.GetRedbeamsLastFlowLogProgressByResourceCrn(context.Background(), resourceCrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V4ProgressApi.GetRedbeamsLastFlowLogProgressByResourceCrn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRedbeamsLastFlowLogProgressByResourceCrn`: FlowProgressResponse
    fmt.Fprintf(os.Stdout, "Response from `V4ProgressApi.GetRedbeamsLastFlowLogProgressByResourceCrn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceCrn** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedbeamsLastFlowLogProgressByResourceCrnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FlowProgressResponse**](FlowProgressResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

