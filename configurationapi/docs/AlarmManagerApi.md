# \AlarmManagerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAlarmManager**](AlarmManagerAPI.md#GetAlarmManager) | **Get** /alarm-manager | Returns a single Alarm Manager
[**UpdateAlarmManager**](AlarmManagerAPI.md#UpdateAlarmManager) | **Patch** /alarm-manager | Update an existing Alarm Manager by name



## GetAlarmManager

> AlarmManagerResponse GetAlarmManager(ctx).Execute()

Returns a single Alarm Manager

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlarmManagerAPI.GetAlarmManager(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlarmManagerAPI.GetAlarmManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlarmManager`: AlarmManagerResponse
    fmt.Fprintf(os.Stdout, "Response from `AlarmManagerAPI.GetAlarmManager`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlarmManagerRequest struct via the builder pattern


### Return type

[**AlarmManagerResponse**](AlarmManagerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlarmManager

> AlarmManagerResponse UpdateAlarmManager(ctx).UpdateRequest(updateRequest).Execute()

Update an existing Alarm Manager by name

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Alarm Manager

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlarmManagerAPI.UpdateAlarmManager(context.Background()).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlarmManagerAPI.UpdateAlarmManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAlarmManager`: AlarmManagerResponse
    fmt.Fprintf(os.Stdout, "Response from `AlarmManagerAPI.UpdateAlarmManager`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlarmManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Alarm Manager | 

### Return type

[**AlarmManagerResponse**](AlarmManagerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

