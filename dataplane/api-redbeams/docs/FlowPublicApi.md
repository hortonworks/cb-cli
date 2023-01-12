# \FlowPublicApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HasFlowRunningByChainIdAndResourceCrn**](FlowPublicApi.md#HasFlowRunningByChainIdAndResourceCrn) | **Get** /flow-public/check/chainId/{chainId} | Check if there is a running flow for chain id and resourceCrn
[**HasFlowRunningByFlowIdAndResourceCrn**](FlowPublicApi.md#HasFlowRunningByFlowIdAndResourceCrn) | **Get** /flow-public/check/flowId/{flowId} | Check if there is a running flow for flow id and resourceId



## HasFlowRunningByChainIdAndResourceCrn

> FlowCheckResponse HasFlowRunningByChainIdAndResourceCrn(ctx, chainId).ResourceCrn(resourceCrn).Execute()

Check if there is a running flow for chain id and resourceCrn



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
    chainId := "chainId_example" // string | 
    resourceCrn := "resourceCrn_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlowPublicApi.HasFlowRunningByChainIdAndResourceCrn(context.Background(), chainId).ResourceCrn(resourceCrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowPublicApi.HasFlowRunningByChainIdAndResourceCrn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HasFlowRunningByChainIdAndResourceCrn`: FlowCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `FlowPublicApi.HasFlowRunningByChainIdAndResourceCrn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHasFlowRunningByChainIdAndResourceCrnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceCrn** | **string** |  | 

### Return type

[**FlowCheckResponse**](FlowCheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HasFlowRunningByFlowIdAndResourceCrn

> FlowCheckResponse HasFlowRunningByFlowIdAndResourceCrn(ctx, flowId).ResourceCrn(resourceCrn).Execute()

Check if there is a running flow for flow id and resourceId



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
    flowId := "flowId_example" // string | 
    resourceCrn := "resourceCrn_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlowPublicApi.HasFlowRunningByFlowIdAndResourceCrn(context.Background(), flowId).ResourceCrn(resourceCrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowPublicApi.HasFlowRunningByFlowIdAndResourceCrn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HasFlowRunningByFlowIdAndResourceCrn`: FlowCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `FlowPublicApi.HasFlowRunningByFlowIdAndResourceCrn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHasFlowRunningByFlowIdAndResourceCrnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceCrn** | **string** |  | 

### Return type

[**FlowCheckResponse**](FlowCheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

