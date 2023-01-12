# \FlowApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFlowChainsStatusesByChainIds1**](FlowApi.md#GetFlowChainsStatusesByChainIds1) | **Get** /flow/check/chainIds | Gets flow check responses for parent chains - Input size max 50
[**GetFlowLogsByFlowId1**](FlowApi.md#GetFlowLogsByFlowId1) | **Get** /flow/logs/{flowId} | Get flow logs by flow id
[**GetFlowLogsByFlowIds1**](FlowApi.md#GetFlowLogsByFlowIds1) | **Get** /flow/logs/flowIds | Get flow logs by a list of flow ids - Input size max 50
[**GetFlowLogsByResourceCrn1**](FlowApi.md#GetFlowLogsByResourceCrn1) | **Get** /flow/logs/resource/crn/{resourceCrn} | Get flow logs for resource by resource CRN
[**GetFlowLogsByResourceName1**](FlowApi.md#GetFlowLogsByResourceName1) | **Get** /flow/logs/resource/name/{resourceName} | Get flow logs for resource by resource name
[**GetLastFlowById1**](FlowApi.md#GetLastFlowById1) | **Get** /flow/logs/{flowId}/last | Get last flow log by flow id
[**GetLastFlowByResourceCrn1**](FlowApi.md#GetLastFlowByResourceCrn1) | **Get** /flow/logs/resource/crn/{resourceCrn}/last | Get last flow log for resource by resource CRN
[**GetLastFlowByResourceName1**](FlowApi.md#GetLastFlowByResourceName1) | **Get** /flow/logs/resource/name/{resourceName}/last | Get last flow log for resource by resource name
[**HasFlowRunningByChainId1**](FlowApi.md#HasFlowRunningByChainId1) | **Get** /flow/check/chainId/{chainId} | Check if there is a running flow for chain id
[**HasFlowRunningByFlowId1**](FlowApi.md#HasFlowRunningByFlowId1) | **Get** /flow/check/flowId/{flowId} | Check if there is a running flow for flow id



## GetFlowChainsStatusesByChainIds1

> PageFlowCheckResponse GetFlowChainsStatusesByChainIds1(ctx).ChainIds(chainIds).Size(size).Page(page).Execute()

Gets flow check responses for parent chains - Input size max 50



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
    chainIds := []string{"Inner_example"} // []string | 
    size := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlowApi.GetFlowChainsStatusesByChainIds1(context.Background()).ChainIds(chainIds).Size(size).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetFlowChainsStatusesByChainIds1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlowChainsStatusesByChainIds1`: PageFlowCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetFlowChainsStatusesByChainIds1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowChainsStatusesByChainIds1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chainIds** | **[]string** |  | 
 **size** | **int32** |  | 
 **page** | **int32** |  | 

### Return type

[**PageFlowCheckResponse**](PageFlowCheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlowLogsByFlowId1

> []FlowLogResponse GetFlowLogsByFlowId1(ctx, flowId).Execute()

Get flow logs by flow id



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlowApi.GetFlowLogsByFlowId1(context.Background(), flowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetFlowLogsByFlowId1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlowLogsByFlowId1`: []FlowLogResponse
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetFlowLogsByFlowId1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowLogsByFlowId1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FlowLogResponse**](FlowLogResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlowLogsByFlowIds1

> PageFlowLogResponse GetFlowLogsByFlowIds1(ctx).FlowIds(flowIds).Size(size).Page(page).Execute()

Get flow logs by a list of flow ids - Input size max 50



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
    flowIds := []string{"Inner_example"} // []string | 
    size := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlowApi.GetFlowLogsByFlowIds1(context.Background()).FlowIds(flowIds).Size(size).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetFlowLogsByFlowIds1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlowLogsByFlowIds1`: PageFlowLogResponse
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetFlowLogsByFlowIds1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowLogsByFlowIds1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flowIds** | **[]string** |  | 
 **size** | **int32** |  | 
 **page** | **int32** |  | 

### Return type

[**PageFlowLogResponse**](PageFlowLogResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlowLogsByResourceCrn1

> []FlowLogResponse GetFlowLogsByResourceCrn1(ctx, resourceCrn).Execute()

Get flow logs for resource by resource CRN



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
    resp, r, err := apiClient.FlowApi.GetFlowLogsByResourceCrn1(context.Background(), resourceCrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetFlowLogsByResourceCrn1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlowLogsByResourceCrn1`: []FlowLogResponse
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetFlowLogsByResourceCrn1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceCrn** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowLogsByResourceCrn1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FlowLogResponse**](FlowLogResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlowLogsByResourceName1

> []FlowLogResponse GetFlowLogsByResourceName1(ctx, resourceName).AccountId(accountId).Execute()

Get flow logs for resource by resource name



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
    resourceName := "resourceName_example" // string | 
    accountId := "accountId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlowApi.GetFlowLogsByResourceName1(context.Background(), resourceName).AccountId(accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetFlowLogsByResourceName1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlowLogsByResourceName1`: []FlowLogResponse
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetFlowLogsByResourceName1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowLogsByResourceName1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** |  | 

### Return type

[**[]FlowLogResponse**](FlowLogResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLastFlowById1

> FlowLogResponse GetLastFlowById1(ctx, flowId).Execute()

Get last flow log by flow id



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlowApi.GetLastFlowById1(context.Background(), flowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetLastFlowById1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLastFlowById1`: FlowLogResponse
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetLastFlowById1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLastFlowById1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FlowLogResponse**](FlowLogResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLastFlowByResourceCrn1

> FlowLogResponse GetLastFlowByResourceCrn1(ctx, resourceCrn).Execute()

Get last flow log for resource by resource CRN



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
    resp, r, err := apiClient.FlowApi.GetLastFlowByResourceCrn1(context.Background(), resourceCrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetLastFlowByResourceCrn1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLastFlowByResourceCrn1`: FlowLogResponse
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetLastFlowByResourceCrn1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceCrn** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLastFlowByResourceCrn1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FlowLogResponse**](FlowLogResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLastFlowByResourceName1

> FlowLogResponse GetLastFlowByResourceName1(ctx, resourceName).AccountId(accountId).Execute()

Get last flow log for resource by resource name



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
    resourceName := "resourceName_example" // string | 
    accountId := "accountId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlowApi.GetLastFlowByResourceName1(context.Background(), resourceName).AccountId(accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.GetLastFlowByResourceName1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLastFlowByResourceName1`: FlowLogResponse
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.GetLastFlowByResourceName1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLastFlowByResourceName1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** |  | 

### Return type

[**FlowLogResponse**](FlowLogResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HasFlowRunningByChainId1

> FlowCheckResponse HasFlowRunningByChainId1(ctx, chainId).Execute()

Check if there is a running flow for chain id



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlowApi.HasFlowRunningByChainId1(context.Background(), chainId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.HasFlowRunningByChainId1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HasFlowRunningByChainId1`: FlowCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.HasFlowRunningByChainId1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHasFlowRunningByChainId1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## HasFlowRunningByFlowId1

> FlowCheckResponse HasFlowRunningByFlowId1(ctx, flowId).Execute()

Check if there is a running flow for flow id



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlowApi.HasFlowRunningByFlowId1(context.Background(), flowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowApi.HasFlowRunningByFlowId1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HasFlowRunningByFlowId1`: FlowCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `FlowApi.HasFlowRunningByFlowId1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHasFlowRunningByFlowId1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

