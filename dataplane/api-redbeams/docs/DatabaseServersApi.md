# \DatabaseServersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDatabaseOnServer**](DatabaseServersApi.md#CreateDatabaseOnServer) | **Post** /v4/databaseservers/createDatabase | create a database on an existing database server
[**CreateDatabaseServer**](DatabaseServersApi.md#CreateDatabaseServer) | **Post** /v4/databaseservers/managed | create and register a database server in a cloud provider
[**CreateDatabaseServerInternal**](DatabaseServersApi.md#CreateDatabaseServerInternal) | **Post** /v4/databaseservers/internal/managed | create and register a database server in a cloud provider with internal actor
[**DeleteDatabaseServerByCrn**](DatabaseServersApi.md#DeleteDatabaseServerByCrn) | **Delete** /v4/databaseservers/{crn} | terminate and/or deregister a database server by CRN
[**DeleteDatabaseServerByName**](DatabaseServersApi.md#DeleteDatabaseServerByName) | **Delete** /v4/databaseservers/name/{name} | terminate and/or deregister a database server by name
[**DeleteMultipleDatabaseServersByCrn**](DatabaseServersApi.md#DeleteMultipleDatabaseServersByCrn) | **Delete** /v4/databaseservers | terminate and/or deregister multiple database servers by CRN
[**GetDatabaseServerByClusterCrn**](DatabaseServersApi.md#GetDatabaseServerByClusterCrn) | **Get** /v4/databaseservers/clusterCrn/{clusterCrn} | get a database server by cluster CRN
[**GetDatabaseServerByCrn**](DatabaseServersApi.md#GetDatabaseServerByCrn) | **Get** /v4/databaseservers/{crn} | get a database server by CRN
[**GetDatabaseServerByName**](DatabaseServersApi.md#GetDatabaseServerByName) | **Get** /v4/databaseservers/name/{name} | get a database server by name
[**GetUsedSubnetsByEnvironment**](DatabaseServersApi.md#GetUsedSubnetsByEnvironment) | **Get** /v4/databaseservers/internal/used_subnets | list the used subnets by the given Environment resource CRN
[**ListDatabaseServers**](DatabaseServersApi.md#ListDatabaseServers) | **Get** /v4/databaseservers | list database servers
[**RegisterDatabaseServer**](DatabaseServersApi.md#RegisterDatabaseServer) | **Post** /v4/databaseservers/register | register a database server
[**ReleaseManagedDatabaseServer**](DatabaseServersApi.md#ReleaseManagedDatabaseServer) | **Put** /v4/databaseservers/{crn}/release | release management of a service-managed database server
[**StartDatabaseServer**](DatabaseServersApi.md#StartDatabaseServer) | **Put** /v4/databaseservers/{crn}/start | start database server
[**StopDatabaseServer**](DatabaseServersApi.md#StopDatabaseServer) | **Put** /v4/databaseservers/{crn}/stop | stop database server
[**TestDatabaseServerConnection**](DatabaseServersApi.md#TestDatabaseServerConnection) | **Post** /v4/databaseservers/test | test database server connectivity
[**UpdateClusterCrn**](DatabaseServersApi.md#UpdateClusterCrn) | **Post** /v4/databaseservers/updateclustercrn | Update the cluster crn associated with the database
[**UpgradeDatabaseServer**](DatabaseServersApi.md#UpgradeDatabaseServer) | **Put** /v4/databaseservers/{crn}/upgrade | upgrade a database server in a cloud provider to a higher major version
[**ValidateUpgradeDatabaseServer**](DatabaseServersApi.md#ValidateUpgradeDatabaseServer) | **Put** /v4/databaseservers/{crn}/validate_upgrade | validate if upgrade is possible on the database server in a cloud provider to a higher major version



## CreateDatabaseOnServer

> CreateDatabaseV4Response CreateDatabaseOnServer(ctx).CreateDatabaseV4Request(createDatabaseV4Request).Execute()

create a database on an existing database server



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
    createDatabaseV4Request := *openapiclient.NewCreateDatabaseV4Request("ExistingDatabaseServerCrn_example", "DatabaseName_example", "Type_example") // CreateDatabaseV4Request | Request for creating a new database on a registered database server (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseServersApi.CreateDatabaseOnServer(context.Background()).CreateDatabaseV4Request(createDatabaseV4Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseServersApi.CreateDatabaseOnServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDatabaseOnServer`: CreateDatabaseV4Response
    fmt.Fprintf(os.Stdout, "Response from `DatabaseServersApi.CreateDatabaseOnServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatabaseOnServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDatabaseV4Request** | [**CreateDatabaseV4Request**](CreateDatabaseV4Request.md) | Request for creating a new database on a registered database server | 

### Return type

[**CreateDatabaseV4Response**](CreateDatabaseV4Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDatabaseServer

> DatabaseServerStatusV4Response CreateDatabaseServer(ctx).AllocateDatabaseServerV4Request(allocateDatabaseServerV4Request).Execute()

create and register a database server in a cloud provider



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
    allocateDatabaseServerV4Request := *openapiclient.NewAllocateDatabaseServerV4Request("EnvironmentCrn_example", "ClusterCrn_example", *openapiclient.NewDatabaseServerV4StackRequest()) // AllocateDatabaseServerV4Request | Request for allocating a new database server in a provider (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseServersApi.CreateDatabaseServer(context.Background()).AllocateDatabaseServerV4Request(allocateDatabaseServerV4Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseServersApi.CreateDatabaseServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDatabaseServer`: DatabaseServerStatusV4Response
    fmt.Fprintf(os.Stdout, "Response from `DatabaseServersApi.CreateDatabaseServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatabaseServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allocateDatabaseServerV4Request** | [**AllocateDatabaseServerV4Request**](AllocateDatabaseServerV4Request.md) | Request for allocating a new database server in a provider | 

### Return type

[**DatabaseServerStatusV4Response**](DatabaseServerStatusV4Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDatabaseServerInternal

> DatabaseServerStatusV4Response CreateDatabaseServerInternal(ctx).InitiatorUserCrn(initiatorUserCrn).AllocateDatabaseServerV4Request(allocateDatabaseServerV4Request).Execute()

create and register a database server in a cloud provider with internal actor



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
    initiatorUserCrn := "initiatorUserCrn_example" // string |  (optional)
    allocateDatabaseServerV4Request := *openapiclient.NewAllocateDatabaseServerV4Request("EnvironmentCrn_example", "ClusterCrn_example", *openapiclient.NewDatabaseServerV4StackRequest()) // AllocateDatabaseServerV4Request | Request for allocating a new database server in a provider (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseServersApi.CreateDatabaseServerInternal(context.Background()).InitiatorUserCrn(initiatorUserCrn).AllocateDatabaseServerV4Request(allocateDatabaseServerV4Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseServersApi.CreateDatabaseServerInternal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDatabaseServerInternal`: DatabaseServerStatusV4Response
    fmt.Fprintf(os.Stdout, "Response from `DatabaseServersApi.CreateDatabaseServerInternal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatabaseServerInternalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **initiatorUserCrn** | **string** |  | 
 **allocateDatabaseServerV4Request** | [**AllocateDatabaseServerV4Request**](AllocateDatabaseServerV4Request.md) | Request for allocating a new database server in a provider | 

### Return type

[**DatabaseServerStatusV4Response**](DatabaseServerStatusV4Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDatabaseServerByCrn

> DatabaseServerV4Response DeleteDatabaseServerByCrn(ctx, crn).Force(force).Execute()

terminate and/or deregister a database server by CRN



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
    crn := "crn_example" // string | CRN of the database server
    force := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseServersApi.DeleteDatabaseServerByCrn(context.Background(), crn).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseServersApi.DeleteDatabaseServerByCrn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDatabaseServerByCrn`: DatabaseServerV4Response
    fmt.Fprintf(os.Stdout, "Response from `DatabaseServersApi.DeleteDatabaseServerByCrn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**crn** | **string** | CRN of the database server | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatabaseServerByCrnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** |  | [default to false]

### Return type

[**DatabaseServerV4Response**](DatabaseServerV4Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDatabaseServerByName

> DatabaseServerV4Response DeleteDatabaseServerByName(ctx, name).EnvironmentCrn(environmentCrn).Force(force).Execute()

terminate and/or deregister a database server by name



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
    environmentCrn := "environmentCrn_example" // string | CRN of the environment of the database server(s)
    name := "name_example" // string | Name of the database server
    force := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseServersApi.DeleteDatabaseServerByName(context.Background(), name).EnvironmentCrn(environmentCrn).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseServersApi.DeleteDatabaseServerByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDatabaseServerByName`: DatabaseServerV4Response
    fmt.Fprintf(os.Stdout, "Response from `DatabaseServersApi.DeleteDatabaseServerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the database server | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatabaseServerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentCrn** | **string** | CRN of the environment of the database server(s) | 

 **force** | **bool** |  | [default to false]

### Return type

[**DatabaseServerV4Response**](DatabaseServerV4Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMultipleDatabaseServersByCrn

> DatabaseServerV4Responses DeleteMultipleDatabaseServersByCrn(ctx).Force(force).RequestBody(requestBody).Execute()

terminate and/or deregister multiple database servers by CRN



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
    force := true // bool |  (optional) (default to false)
    requestBody := []string{"Property_example"} // []string | CRNs of the database servers (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseServersApi.DeleteMultipleDatabaseServersByCrn(context.Background()).Force(force).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseServersApi.DeleteMultipleDatabaseServersByCrn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMultipleDatabaseServersByCrn`: DatabaseServerV4Responses
    fmt.Fprintf(os.Stdout, "Response from `DatabaseServersApi.DeleteMultipleDatabaseServersByCrn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMultipleDatabaseServersByCrnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **force** | **bool** |  | [default to false]
 **requestBody** | **[]string** | CRNs of the database servers | 

### Return type

[**DatabaseServerV4Responses**](DatabaseServerV4Responses.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseServerByClusterCrn

> DatabaseServerV4Response GetDatabaseServerByClusterCrn(ctx, clusterCrn).EnvironmentCrn(environmentCrn).Execute()

get a database server by cluster CRN



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
    environmentCrn := "environmentCrn_example" // string | CRN of the environment of the database server(s)
    clusterCrn := "clusterCrn_example" // string | CRN of cluster of the database server

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseServersApi.GetDatabaseServerByClusterCrn(context.Background(), clusterCrn).EnvironmentCrn(environmentCrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseServersApi.GetDatabaseServerByClusterCrn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseServerByClusterCrn`: DatabaseServerV4Response
    fmt.Fprintf(os.Stdout, "Response from `DatabaseServersApi.GetDatabaseServerByClusterCrn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterCrn** | **string** | CRN of cluster of the database server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseServerByClusterCrnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentCrn** | **string** | CRN of the environment of the database server(s) | 


### Return type

[**DatabaseServerV4Response**](DatabaseServerV4Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseServerByCrn

> DatabaseServerV4Response GetDatabaseServerByCrn(ctx, crn).Execute()

get a database server by CRN



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
    crn := "crn_example" // string | CRN of the database server

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseServersApi.GetDatabaseServerByCrn(context.Background(), crn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseServersApi.GetDatabaseServerByCrn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseServerByCrn`: DatabaseServerV4Response
    fmt.Fprintf(os.Stdout, "Response from `DatabaseServersApi.GetDatabaseServerByCrn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**crn** | **string** | CRN of the database server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseServerByCrnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DatabaseServerV4Response**](DatabaseServerV4Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseServerByName

> DatabaseServerV4Response GetDatabaseServerByName(ctx, name).EnvironmentCrn(environmentCrn).Execute()

get a database server by name



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
    environmentCrn := "environmentCrn_example" // string | CRN of the environment of the database server(s)
    name := "name_example" // string | Name of the database server

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseServersApi.GetDatabaseServerByName(context.Background(), name).EnvironmentCrn(environmentCrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseServersApi.GetDatabaseServerByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseServerByName`: DatabaseServerV4Response
    fmt.Fprintf(os.Stdout, "Response from `DatabaseServersApi.GetDatabaseServerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the database server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseServerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentCrn** | **string** | CRN of the environment of the database server(s) | 


### Return type

[**DatabaseServerV4Response**](DatabaseServerV4Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsedSubnetsByEnvironment

> UsedSubnetsByEnvironmentResponse GetUsedSubnetsByEnvironment(ctx).EnvironmentCrn(environmentCrn).Execute()

list the used subnets by the given Environment resource CRN



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
    environmentCrn := "environmentCrn_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseServersApi.GetUsedSubnetsByEnvironment(context.Background()).EnvironmentCrn(environmentCrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseServersApi.GetUsedSubnetsByEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsedSubnetsByEnvironment`: UsedSubnetsByEnvironmentResponse
    fmt.Fprintf(os.Stdout, "Response from `DatabaseServersApi.GetUsedSubnetsByEnvironment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsedSubnetsByEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentCrn** | **string** |  | 

### Return type

[**UsedSubnetsByEnvironmentResponse**](UsedSubnetsByEnvironmentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatabaseServers

> DatabaseServerV4Responses ListDatabaseServers(ctx).EnvironmentCrn(environmentCrn).Execute()

list database servers



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
    environmentCrn := "environmentCrn_example" // string | CRN of the environment of the database server(s)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseServersApi.ListDatabaseServers(context.Background()).EnvironmentCrn(environmentCrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseServersApi.ListDatabaseServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDatabaseServers`: DatabaseServerV4Responses
    fmt.Fprintf(os.Stdout, "Response from `DatabaseServersApi.ListDatabaseServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDatabaseServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentCrn** | **string** | CRN of the environment of the database server(s) | 

### Return type

[**DatabaseServerV4Responses**](DatabaseServerV4Responses.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterDatabaseServer

> DatabaseServerV4Response RegisterDatabaseServer(ctx).DatabaseServerV4Request(databaseServerV4Request).Execute()

register a database server



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
    databaseServerV4Request := *openapiclient.NewDatabaseServerV4Request("Name_example", "Host_example", int32(123), "DatabaseVendor_example", "EnvironmentCrn_example", "ConnectionUserName_example", "ConnectionPassword_example") // DatabaseServerV4Request | Request containing information about a database server to be registered (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseServersApi.RegisterDatabaseServer(context.Background()).DatabaseServerV4Request(databaseServerV4Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseServersApi.RegisterDatabaseServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterDatabaseServer`: DatabaseServerV4Response
    fmt.Fprintf(os.Stdout, "Response from `DatabaseServersApi.RegisterDatabaseServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterDatabaseServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **databaseServerV4Request** | [**DatabaseServerV4Request**](DatabaseServerV4Request.md) | Request containing information about a database server to be registered | 

### Return type

[**DatabaseServerV4Response**](DatabaseServerV4Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReleaseManagedDatabaseServer

> DatabaseServerV4Response ReleaseManagedDatabaseServer(ctx, crn).Execute()

release management of a service-managed database server



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
    crn := "crn_example" // string | CRN of the database server

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseServersApi.ReleaseManagedDatabaseServer(context.Background(), crn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseServersApi.ReleaseManagedDatabaseServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReleaseManagedDatabaseServer`: DatabaseServerV4Response
    fmt.Fprintf(os.Stdout, "Response from `DatabaseServersApi.ReleaseManagedDatabaseServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**crn** | **string** | CRN of the database server | 

### Other Parameters

Other parameters are passed through a pointer to a apiReleaseManagedDatabaseServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DatabaseServerV4Response**](DatabaseServerV4Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartDatabaseServer

> StartDatabaseServer(ctx, crn).Execute()

start database server



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
    crn := "crn_example" // string | CRN of the database server

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseServersApi.StartDatabaseServer(context.Background(), crn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseServersApi.StartDatabaseServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**crn** | **string** | CRN of the database server | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartDatabaseServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopDatabaseServer

> StopDatabaseServer(ctx, crn).Execute()

stop database server



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
    crn := "crn_example" // string | CRN of the database server

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseServersApi.StopDatabaseServer(context.Background(), crn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseServersApi.StopDatabaseServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**crn** | **string** | CRN of the database server | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopDatabaseServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestDatabaseServerConnection

> DatabaseServerTestV4Response TestDatabaseServerConnection(ctx).DatabaseServerTestV4Request(databaseServerTestV4Request).Execute()

test database server connectivity



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
    databaseServerTestV4Request := *openapiclient.NewDatabaseServerTestV4Request() // DatabaseServerTestV4Request | Request for testing connectivity to a database server (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseServersApi.TestDatabaseServerConnection(context.Background()).DatabaseServerTestV4Request(databaseServerTestV4Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseServersApi.TestDatabaseServerConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestDatabaseServerConnection`: DatabaseServerTestV4Response
    fmt.Fprintf(os.Stdout, "Response from `DatabaseServersApi.TestDatabaseServerConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestDatabaseServerConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **databaseServerTestV4Request** | [**DatabaseServerTestV4Request**](DatabaseServerTestV4Request.md) | Request for testing connectivity to a database server | 

### Return type

[**DatabaseServerTestV4Response**](DatabaseServerTestV4Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClusterCrn

> UpdateClusterCrn(ctx).EnvironmentCrn(environmentCrn).CurrentClusterCrn(currentClusterCrn).NewClusterCrn(newClusterCrn).InitiatorUserCrn(initiatorUserCrn).Execute()

Update the cluster crn associated with the database



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
    environmentCrn := "environmentCrn_example" // string |  (optional)
    currentClusterCrn := "currentClusterCrn_example" // string |  (optional)
    newClusterCrn := "newClusterCrn_example" // string |  (optional)
    initiatorUserCrn := "initiatorUserCrn_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseServersApi.UpdateClusterCrn(context.Background()).EnvironmentCrn(environmentCrn).CurrentClusterCrn(currentClusterCrn).NewClusterCrn(newClusterCrn).InitiatorUserCrn(initiatorUserCrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseServersApi.UpdateClusterCrn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterCrnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentCrn** | **string** |  | 
 **currentClusterCrn** | **string** |  | 
 **newClusterCrn** | **string** |  | 
 **initiatorUserCrn** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeDatabaseServer

> UpgradeDatabaseServerV4Response UpgradeDatabaseServer(ctx, crn).UpgradeDatabaseServerV4Request(upgradeDatabaseServerV4Request).Execute()

upgrade a database server in a cloud provider to a higher major version



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
    crn := "crn_example" // string | CRN of the database server
    upgradeDatabaseServerV4Request := *openapiclient.NewUpgradeDatabaseServerV4Request("UpgradeTargetMajorVersion_example") // UpgradeDatabaseServerV4Request | Request for upgrading a database server in a provider to a higher major version

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseServersApi.UpgradeDatabaseServer(context.Background(), crn).UpgradeDatabaseServerV4Request(upgradeDatabaseServerV4Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseServersApi.UpgradeDatabaseServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpgradeDatabaseServer`: UpgradeDatabaseServerV4Response
    fmt.Fprintf(os.Stdout, "Response from `DatabaseServersApi.UpgradeDatabaseServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**crn** | **string** | CRN of the database server | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeDatabaseServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upgradeDatabaseServerV4Request** | [**UpgradeDatabaseServerV4Request**](UpgradeDatabaseServerV4Request.md) | Request for upgrading a database server in a provider to a higher major version | 

### Return type

[**UpgradeDatabaseServerV4Response**](UpgradeDatabaseServerV4Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateUpgradeDatabaseServer

> UpgradeDatabaseServerV4Response ValidateUpgradeDatabaseServer(ctx, crn).UpgradeDatabaseServerV4Request(upgradeDatabaseServerV4Request).Execute()

validate if upgrade is possible on the database server in a cloud provider to a higher major version



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
    crn := "crn_example" // string | CRN of the database server
    upgradeDatabaseServerV4Request := *openapiclient.NewUpgradeDatabaseServerV4Request("UpgradeTargetMajorVersion_example") // UpgradeDatabaseServerV4Request | Request for upgrading a database server in a provider to a higher major version

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseServersApi.ValidateUpgradeDatabaseServer(context.Background(), crn).UpgradeDatabaseServerV4Request(upgradeDatabaseServerV4Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseServersApi.ValidateUpgradeDatabaseServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateUpgradeDatabaseServer`: UpgradeDatabaseServerV4Response
    fmt.Fprintf(os.Stdout, "Response from `DatabaseServersApi.ValidateUpgradeDatabaseServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**crn** | **string** | CRN of the database server | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpgradeDatabaseServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upgradeDatabaseServerV4Request** | [**UpgradeDatabaseServerV4Request**](UpgradeDatabaseServerV4Request.md) | Request for upgrading a database server in a provider to a higher major version | 

### Return type

[**UpgradeDatabaseServerV4Response**](UpgradeDatabaseServerV4Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

