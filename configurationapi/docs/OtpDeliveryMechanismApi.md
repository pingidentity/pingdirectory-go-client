# \OtpDeliveryMechanismAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOtpDeliveryMechanism**](OtpDeliveryMechanismAPI.md#AddOtpDeliveryMechanism) | **Post** /otp-delivery-mechanisms | Add a new OTP Delivery Mechanism to the config
[**DeleteOtpDeliveryMechanism**](OtpDeliveryMechanismAPI.md#DeleteOtpDeliveryMechanism) | **Delete** /otp-delivery-mechanisms/{otp-delivery-mechanism-name} | Delete a OTP Delivery Mechanism
[**GetOtpDeliveryMechanism**](OtpDeliveryMechanismAPI.md#GetOtpDeliveryMechanism) | **Get** /otp-delivery-mechanisms/{otp-delivery-mechanism-name} | Returns a single OTP Delivery Mechanism
[**ListOtpDeliveryMechanisms**](OtpDeliveryMechanismAPI.md#ListOtpDeliveryMechanisms) | **Get** /otp-delivery-mechanisms | Returns a list of all OTP Delivery Mechanism objects
[**UpdateOtpDeliveryMechanism**](OtpDeliveryMechanismAPI.md#UpdateOtpDeliveryMechanism) | **Patch** /otp-delivery-mechanisms/{otp-delivery-mechanism-name} | Update an existing OTP Delivery Mechanism by name



## AddOtpDeliveryMechanism

> AddOtpDeliveryMechanism200Response AddOtpDeliveryMechanism(ctx).AddOtpDeliveryMechanismRequest(addOtpDeliveryMechanismRequest).Execute()

Add a new OTP Delivery Mechanism to the config

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
    addOtpDeliveryMechanismRequest := openapiclient.add_otp_delivery_mechanism_request{AddEmailOtpDeliveryMechanismRequest: openapiclient.NewAddEmailOtpDeliveryMechanismRequest([]openapiclient.EnumemailOtpDeliveryMechanismSchemaUrn{openapiclient.Enumemail-otp-delivery-mechanismSchemaUrn("urn:pingidentity:schemas:configuration:2.0:otp-delivery-mechanism:email")}, "SenderAddress_example", false, "MechanismName_example")} // AddOtpDeliveryMechanismRequest | Create a new OTP Delivery Mechanism in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OtpDeliveryMechanismAPI.AddOtpDeliveryMechanism(context.Background()).AddOtpDeliveryMechanismRequest(addOtpDeliveryMechanismRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtpDeliveryMechanismAPI.AddOtpDeliveryMechanism``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddOtpDeliveryMechanism`: AddOtpDeliveryMechanism200Response
    fmt.Fprintf(os.Stdout, "Response from `OtpDeliveryMechanismAPI.AddOtpDeliveryMechanism`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddOtpDeliveryMechanismRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addOtpDeliveryMechanismRequest** | [**AddOtpDeliveryMechanismRequest**](AddOtpDeliveryMechanismRequest.md) | Create a new OTP Delivery Mechanism in the config | 

### Return type

[**AddOtpDeliveryMechanism200Response**](AddOtpDeliveryMechanism200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOtpDeliveryMechanism

> DeleteOtpDeliveryMechanism(ctx, otpDeliveryMechanismName).Execute()

Delete a OTP Delivery Mechanism

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
    otpDeliveryMechanismName := "otpDeliveryMechanismName_example" // string | Name of the OTP Delivery Mechanism

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OtpDeliveryMechanismAPI.DeleteOtpDeliveryMechanism(context.Background(), otpDeliveryMechanismName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtpDeliveryMechanismAPI.DeleteOtpDeliveryMechanism``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**otpDeliveryMechanismName** | **string** | Name of the OTP Delivery Mechanism | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOtpDeliveryMechanismRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOtpDeliveryMechanism

> AddOtpDeliveryMechanism200Response GetOtpDeliveryMechanism(ctx, otpDeliveryMechanismName).Execute()

Returns a single OTP Delivery Mechanism

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
    otpDeliveryMechanismName := "otpDeliveryMechanismName_example" // string | Name of the OTP Delivery Mechanism

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OtpDeliveryMechanismAPI.GetOtpDeliveryMechanism(context.Background(), otpDeliveryMechanismName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtpDeliveryMechanismAPI.GetOtpDeliveryMechanism``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOtpDeliveryMechanism`: AddOtpDeliveryMechanism200Response
    fmt.Fprintf(os.Stdout, "Response from `OtpDeliveryMechanismAPI.GetOtpDeliveryMechanism`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**otpDeliveryMechanismName** | **string** | Name of the OTP Delivery Mechanism | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOtpDeliveryMechanismRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddOtpDeliveryMechanism200Response**](AddOtpDeliveryMechanism200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOtpDeliveryMechanisms

> OtpDeliveryMechanismListResponse ListOtpDeliveryMechanisms(ctx).Filter(filter).Execute()

Returns a list of all OTP Delivery Mechanism objects

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
    filter := "filter_example" // string | SCIM filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OtpDeliveryMechanismAPI.ListOtpDeliveryMechanisms(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtpDeliveryMechanismAPI.ListOtpDeliveryMechanisms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOtpDeliveryMechanisms`: OtpDeliveryMechanismListResponse
    fmt.Fprintf(os.Stdout, "Response from `OtpDeliveryMechanismAPI.ListOtpDeliveryMechanisms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOtpDeliveryMechanismsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**OtpDeliveryMechanismListResponse**](OtpDeliveryMechanismListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOtpDeliveryMechanism

> AddOtpDeliveryMechanism200Response UpdateOtpDeliveryMechanism(ctx, otpDeliveryMechanismName).UpdateRequest(updateRequest).Execute()

Update an existing OTP Delivery Mechanism by name

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
    otpDeliveryMechanismName := "otpDeliveryMechanismName_example" // string | Name of the OTP Delivery Mechanism
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing OTP Delivery Mechanism

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OtpDeliveryMechanismAPI.UpdateOtpDeliveryMechanism(context.Background(), otpDeliveryMechanismName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtpDeliveryMechanismAPI.UpdateOtpDeliveryMechanism``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOtpDeliveryMechanism`: AddOtpDeliveryMechanism200Response
    fmt.Fprintf(os.Stdout, "Response from `OtpDeliveryMechanismAPI.UpdateOtpDeliveryMechanism`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**otpDeliveryMechanismName** | **string** | Name of the OTP Delivery Mechanism | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOtpDeliveryMechanismRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing OTP Delivery Mechanism | 

### Return type

[**AddOtpDeliveryMechanism200Response**](AddOtpDeliveryMechanism200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

