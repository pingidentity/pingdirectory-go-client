# \OtpDeliveryMechanismApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOtpDeliveryMechanism**](OtpDeliveryMechanismApi.md#AddOtpDeliveryMechanism) | **Post** /otp-delivery-mechanisms | Add a new OTP Delivery Mechanism to the config
[**DeleteOtpDeliveryMechanism**](OtpDeliveryMechanismApi.md#DeleteOtpDeliveryMechanism) | **Delete** /otp-delivery-mechanisms/{otp-delivery-mechanism-name} | Delete a OTP Delivery Mechanism
[**GetOtpDeliveryMechanism**](OtpDeliveryMechanismApi.md#GetOtpDeliveryMechanism) | **Get** /otp-delivery-mechanisms/{otp-delivery-mechanism-name} | Returns a single OTP Delivery Mechanism
[**UpdateOtpDeliveryMechanism**](OtpDeliveryMechanismApi.md#UpdateOtpDeliveryMechanism) | **Patch** /otp-delivery-mechanisms/{otp-delivery-mechanism-name} | Update an existing OTP Delivery Mechanism by name



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
    openapiclient "./openapi"
)

func main() {
    addOtpDeliveryMechanismRequest := openapiclient.add_otp_delivery_mechanism_request{AddEmailOtpDeliveryMechanismRequest: openapiclient.NewAddEmailOtpDeliveryMechanismRequest("MechanismName_example", []openapiclient.EnumemailOtpDeliveryMechanismSchemaUrn{openapiclient.Enumemail-otp-delivery-mechanismSchemaUrn("urn:pingidentity:schemas:configuration:2.0:otp-delivery-mechanism:email")}, "EmailAddressAttributeType_example", "SenderAddress_example", "MessageSubject_example", false)} // AddOtpDeliveryMechanismRequest | Create a new OTP Delivery Mechanism in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OtpDeliveryMechanismApi.AddOtpDeliveryMechanism(context.Background()).AddOtpDeliveryMechanismRequest(addOtpDeliveryMechanismRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtpDeliveryMechanismApi.AddOtpDeliveryMechanism``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddOtpDeliveryMechanism`: AddOtpDeliveryMechanism200Response
    fmt.Fprintf(os.Stdout, "Response from `OtpDeliveryMechanismApi.AddOtpDeliveryMechanism`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    otpDeliveryMechanismName := "otpDeliveryMechanismName_example" // string | Name of the OTP Delivery Mechanism to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OtpDeliveryMechanismApi.DeleteOtpDeliveryMechanism(context.Background(), otpDeliveryMechanismName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtpDeliveryMechanismApi.DeleteOtpDeliveryMechanism``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**otpDeliveryMechanismName** | **string** | Name of the OTP Delivery Mechanism to be deleted | 

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
    openapiclient "./openapi"
)

func main() {
    otpDeliveryMechanismName := "otpDeliveryMechanismName_example" // string | Name of the OTP Delivery Mechanism to be read

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OtpDeliveryMechanismApi.GetOtpDeliveryMechanism(context.Background(), otpDeliveryMechanismName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtpDeliveryMechanismApi.GetOtpDeliveryMechanism``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOtpDeliveryMechanism`: AddOtpDeliveryMechanism200Response
    fmt.Fprintf(os.Stdout, "Response from `OtpDeliveryMechanismApi.GetOtpDeliveryMechanism`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**otpDeliveryMechanismName** | **string** | Name of the OTP Delivery Mechanism to be read | 

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
    openapiclient "./openapi"
)

func main() {
    otpDeliveryMechanismName := "otpDeliveryMechanismName_example" // string | Name of the OTP Delivery Mechanism to be updated
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing OTP Delivery Mechanism

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OtpDeliveryMechanismApi.UpdateOtpDeliveryMechanism(context.Background(), otpDeliveryMechanismName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtpDeliveryMechanismApi.UpdateOtpDeliveryMechanism``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOtpDeliveryMechanism`: AddOtpDeliveryMechanism200Response
    fmt.Fprintf(os.Stdout, "Response from `OtpDeliveryMechanismApi.UpdateOtpDeliveryMechanism`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**otpDeliveryMechanismName** | **string** | Name of the OTP Delivery Mechanism to be updated | 

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

