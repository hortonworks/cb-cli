# \DatabasesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDatabaseByCrn**](DatabasesApi.md#DeleteDatabaseByCrn) | **Delete** /v4/databases/{crn} | delete a database config by CRN
[**DeleteDatabaseByName**](DatabasesApi.md#DeleteDatabaseByName) | **Delete** /v4/databases/name/{name} | delete a database config by name
[**DeleteMultipleDatabasesByCrn**](DatabasesApi.md#DeleteMultipleDatabasesByCrn) | **Delete** /v4/databases | delete multiple database configs by CRN
[**GetDatabaseByCrn**](DatabasesApi.md#GetDatabaseByCrn) | **Get** /v4/databases/{crn} | get a database config by CRN
[**GetDatabaseByName**](DatabasesApi.md#GetDatabaseByName) | **Get** /v4/databases/name/{name} | get a database config by name
[**ListDatabases**](DatabasesApi.md#ListDatabases) | **Get** /v4/databases | list database configs
[**RegisterDatabase**](DatabasesApi.md#RegisterDatabase) | **Post** /v4/databases/register | register a database config of existing database
[**TestDatabaseConnection**](DatabasesApi.md#TestDatabaseConnection) | **Post** /v4/databases/test | test database connectivity



## DeleteDatabaseByCrn

> DatabaseV4Response DeleteDatabaseByCrn(ctx, crn).Execute()

delete a database config by CRN



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
    crn := "crn_example" // string | CRN of the database

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesApi.DeleteDatabaseByCrn(context.Background(), crn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesApi.DeleteDatabaseByCrn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDatabaseByCrn`: DatabaseV4Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesApi.DeleteDatabaseByCrn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**crn** | **string** | CRN of the database | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatabaseByCrnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DatabaseV4Response**](DatabaseV4Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDatabaseByName

> DatabaseV4Response DeleteDatabaseByName(ctx, name).EnvironmentCrn(environmentCrn).Execute()

delete a database config by name



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
    environmentCrn := "environmentCrn_example" // string | CRN of the environment of the database(s)
    name := "name_example" // string | Name of the database

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesApi.DeleteDatabaseByName(context.Background(), name).EnvironmentCrn(environmentCrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesApi.DeleteDatabaseByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDatabaseByName`: DatabaseV4Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesApi.DeleteDatabaseByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the database | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatabaseByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentCrn** | **string** | CRN of the environment of the database(s) | 


### Return type

[**DatabaseV4Response**](DatabaseV4Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMultipleDatabasesByCrn

> DatabaseV4Responses DeleteMultipleDatabasesByCrn(ctx).RequestBody(requestBody).Execute()

delete multiple database configs by CRN



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
    requestBody := []string{"Property_example"} // []string | CRNs of the databases (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesApi.DeleteMultipleDatabasesByCrn(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesApi.DeleteMultipleDatabasesByCrn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMultipleDatabasesByCrn`: DatabaseV4Responses
    fmt.Fprintf(os.Stdout, "Response from `DatabasesApi.DeleteMultipleDatabasesByCrn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMultipleDatabasesByCrnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** | CRNs of the databases | 

### Return type

[**DatabaseV4Responses**](DatabaseV4Responses.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseByCrn

> DatabaseV4Response GetDatabaseByCrn(ctx, crn).Execute()

get a database config by CRN



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
    crn := "crn_example" // string | CRN of the database

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesApi.GetDatabaseByCrn(context.Background(), crn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesApi.GetDatabaseByCrn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseByCrn`: DatabaseV4Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesApi.GetDatabaseByCrn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**crn** | **string** | CRN of the database | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseByCrnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DatabaseV4Response**](DatabaseV4Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseByName

> DatabaseV4Response GetDatabaseByName(ctx, name).EnvironmentCrn(environmentCrn).Execute()

get a database config by name



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
    environmentCrn := "environmentCrn_example" // string | CRN of the environment of the database(s)
    name := "name_example" // string | Name of the database

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesApi.GetDatabaseByName(context.Background(), name).EnvironmentCrn(environmentCrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesApi.GetDatabaseByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseByName`: DatabaseV4Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesApi.GetDatabaseByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the database | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentCrn** | **string** | CRN of the environment of the database(s) | 


### Return type

[**DatabaseV4Response**](DatabaseV4Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatabases

> DatabaseV4Responses ListDatabases(ctx).EnvironmentCrn(environmentCrn).Execute()

list database configs



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
    environmentCrn := "environmentCrn_example" // string | CRN of the environment of the database(s)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesApi.ListDatabases(context.Background()).EnvironmentCrn(environmentCrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesApi.ListDatabases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDatabases`: DatabaseV4Responses
    fmt.Fprintf(os.Stdout, "Response from `DatabasesApi.ListDatabases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentCrn** | **string** | CRN of the environment of the database(s) | 

### Return type

[**DatabaseV4Responses**](DatabaseV4Responses.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterDatabase

> DatabaseV4Response RegisterDatabase(ctx).DatabaseV4Request(databaseV4Request).Execute()

register a database config of existing database



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
    databaseV4Request := *openapiclient.NewDatabaseV4Request("Name_example", "ConnectionURL_example", "Type_example", "EnvironmentCrn_example", "ConnectionUserName_example", "ConnectionPassword_example") // DatabaseV4Request | Request containing information about a database to be registered (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesApi.RegisterDatabase(context.Background()).DatabaseV4Request(databaseV4Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesApi.RegisterDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterDatabase`: DatabaseV4Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesApi.RegisterDatabase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **databaseV4Request** | [**DatabaseV4Request**](DatabaseV4Request.md) | Request containing information about a database to be registered | 

### Return type

[**DatabaseV4Response**](DatabaseV4Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestDatabaseConnection

> DatabaseTestV4Response TestDatabaseConnection(ctx).DatabaseTestV4Request(databaseTestV4Request).Execute()

test database connectivity



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
    databaseTestV4Request := *openapiclient.NewDatabaseTestV4Request() // DatabaseTestV4Request | Request for testing connectivity to a database (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesApi.TestDatabaseConnection(context.Background()).DatabaseTestV4Request(databaseTestV4Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesApi.TestDatabaseConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestDatabaseConnection`: DatabaseTestV4Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesApi.TestDatabaseConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestDatabaseConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **databaseTestV4Request** | [**DatabaseTestV4Request**](DatabaseTestV4Request.md) | Request for testing connectivity to a database | 

### Return type

[**DatabaseTestV4Response**](DatabaseTestV4Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

