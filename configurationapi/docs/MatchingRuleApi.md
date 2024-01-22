# \MatchingRuleAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMatchingRule**](MatchingRuleAPI.md#GetMatchingRule) | **Get** /matching-rules/{matching-rule-name} | Returns a single Matching Rule
[**ListMatchingRules**](MatchingRuleAPI.md#ListMatchingRules) | **Get** /matching-rules | Returns a list of all Matching Rule objects
[**UpdateMatchingRule**](MatchingRuleAPI.md#UpdateMatchingRule) | **Patch** /matching-rules/{matching-rule-name} | Update an existing Matching Rule by name



## GetMatchingRule

> GetMatchingRule200Response GetMatchingRule(ctx, matchingRuleName).Execute()

Returns a single Matching Rule

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
    matchingRuleName := "matchingRuleName_example" // string | Name of the Matching Rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MatchingRuleAPI.GetMatchingRule(context.Background(), matchingRuleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MatchingRuleAPI.GetMatchingRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMatchingRule`: GetMatchingRule200Response
    fmt.Fprintf(os.Stdout, "Response from `MatchingRuleAPI.GetMatchingRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**matchingRuleName** | **string** | Name of the Matching Rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMatchingRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMatchingRule200Response**](GetMatchingRule200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMatchingRules

> MatchingRuleListResponse ListMatchingRules(ctx).Filter(filter).Execute()

Returns a list of all Matching Rule objects

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
    resp, r, err := apiClient.MatchingRuleAPI.ListMatchingRules(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MatchingRuleAPI.ListMatchingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMatchingRules`: MatchingRuleListResponse
    fmt.Fprintf(os.Stdout, "Response from `MatchingRuleAPI.ListMatchingRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMatchingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**MatchingRuleListResponse**](MatchingRuleListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMatchingRule

> GetMatchingRule200Response UpdateMatchingRule(ctx, matchingRuleName).UpdateRequest(updateRequest).Execute()

Update an existing Matching Rule by name

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
    matchingRuleName := "matchingRuleName_example" // string | Name of the Matching Rule
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Matching Rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MatchingRuleAPI.UpdateMatchingRule(context.Background(), matchingRuleName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MatchingRuleAPI.UpdateMatchingRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMatchingRule`: GetMatchingRule200Response
    fmt.Fprintf(os.Stdout, "Response from `MatchingRuleAPI.UpdateMatchingRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**matchingRuleName** | **string** | Name of the Matching Rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMatchingRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Matching Rule | 

### Return type

[**GetMatchingRule200Response**](GetMatchingRule200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

