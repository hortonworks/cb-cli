# \V4UtilsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckRightByCrn**](V4UtilsApi.md#CheckRightByCrn) | **Post** /v4/utils/check_right_by_crn | Checking rights from UI by resource CRN
[**CheckRightInAccount**](V4UtilsApi.md#CheckRightInAccount) | **Post** /v4/utils/check_right | Checking rights from UI in account
[**CheckRightOnResources**](V4UtilsApi.md#CheckRightOnResources) | **Post** /v4/utils/check_right_on_resources | Checking right from Uluwatu by resource CRNs



## CheckRightByCrn

> CheckResourceRightsV4Response CheckRightByCrn(ctx).CheckResourceRightsV4Request(checkResourceRightsV4Request).Execute()

Checking rights from UI by resource CRN



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
    checkResourceRightsV4Request := *openapiclient.NewCheckResourceRightsV4Request() // CheckResourceRightsV4Request |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.V4UtilsApi.CheckRightByCrn(context.Background()).CheckResourceRightsV4Request(checkResourceRightsV4Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V4UtilsApi.CheckRightByCrn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckRightByCrn`: CheckResourceRightsV4Response
    fmt.Fprintf(os.Stdout, "Response from `V4UtilsApi.CheckRightByCrn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckRightByCrnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkResourceRightsV4Request** | [**CheckResourceRightsV4Request**](CheckResourceRightsV4Request.md) |  | 

### Return type

[**CheckResourceRightsV4Response**](CheckResourceRightsV4Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckRightInAccount

> CheckRightV4Response CheckRightInAccount(ctx).CheckRightV4Request(checkRightV4Request).Execute()

Checking rights from UI in account



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
    checkRightV4Request := *openapiclient.NewCheckRightV4Request() // CheckRightV4Request |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.V4UtilsApi.CheckRightInAccount(context.Background()).CheckRightV4Request(checkRightV4Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V4UtilsApi.CheckRightInAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckRightInAccount`: CheckRightV4Response
    fmt.Fprintf(os.Stdout, "Response from `V4UtilsApi.CheckRightInAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckRightInAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkRightV4Request** | [**CheckRightV4Request**](CheckRightV4Request.md) |  | 

### Return type

[**CheckRightV4Response**](CheckRightV4Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckRightOnResources

> CheckRightOnResourcesV4Response CheckRightOnResources(ctx).CheckRightOnResourcesV4Request(checkRightOnResourcesV4Request).Execute()

Checking right from Uluwatu by resource CRNs



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
    checkRightOnResourcesV4Request := *openapiclient.NewCheckRightOnResourcesV4Request("Right_example", []string{"ResourceCrns_example"}) // CheckRightOnResourcesV4Request |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.V4UtilsApi.CheckRightOnResources(context.Background()).CheckRightOnResourcesV4Request(checkRightOnResourcesV4Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V4UtilsApi.CheckRightOnResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckRightOnResources`: CheckRightOnResourcesV4Response
    fmt.Fprintf(os.Stdout, "Response from `V4UtilsApi.CheckRightOnResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckRightOnResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkRightOnResourcesV4Request** | [**CheckRightOnResourcesV4Request**](CheckRightOnResourcesV4Request.md) |  | 

### Return type

[**CheckRightOnResourcesV4Response**](CheckRightOnResourcesV4Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

