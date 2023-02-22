# AddHttpServletCrossOriginPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyName** | **string** | Name of the new HTTP Servlet Cross Origin Policy | 
**Schemas** | Pointer to [**[]EnumhttpServletCrossOriginPolicySchemaUrn**](EnumhttpServletCrossOriginPolicySchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this HTTP Servlet Cross Origin Policy | [optional] 
**CorsAllowedMethods** | Pointer to **[]string** | A list of HTTP methods allowed for cross-origin access to resources. i.e. one or more of GET, POST, PUT, DELETE, etc. | [optional] 
**CorsAllowedOrigins** | Pointer to **[]string** | A list of origins that are allowed to execute cross-origin requests. | [optional] 
**CorsExposedHeaders** | Pointer to **[]string** | A list of HTTP headers other than the simple response headers that browsers are allowed to access. | [optional] 
**CorsAllowedHeaders** | Pointer to **[]string** | A list of HTTP headers that are supported by the resource and can be specified in a cross-origin request. | [optional] 
**CorsPreflightMaxAge** | Pointer to **string** | The maximum amount of time that a preflight request can be cached by a client. | [optional] 
**CorsAllowCredentials** | Pointer to **bool** | Indicates whether the servlet extension allows CORS requests with username/password credentials. | [optional] 

## Methods

### NewAddHttpServletCrossOriginPolicyRequest

`func NewAddHttpServletCrossOriginPolicyRequest(policyName string, ) *AddHttpServletCrossOriginPolicyRequest`

NewAddHttpServletCrossOriginPolicyRequest instantiates a new AddHttpServletCrossOriginPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddHttpServletCrossOriginPolicyRequestWithDefaults

`func NewAddHttpServletCrossOriginPolicyRequestWithDefaults() *AddHttpServletCrossOriginPolicyRequest`

NewAddHttpServletCrossOriginPolicyRequestWithDefaults instantiates a new AddHttpServletCrossOriginPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyName

`func (o *AddHttpServletCrossOriginPolicyRequest) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *AddHttpServletCrossOriginPolicyRequest) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *AddHttpServletCrossOriginPolicyRequest) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.


### GetSchemas

`func (o *AddHttpServletCrossOriginPolicyRequest) GetSchemas() []EnumhttpServletCrossOriginPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddHttpServletCrossOriginPolicyRequest) GetSchemasOk() (*[]EnumhttpServletCrossOriginPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddHttpServletCrossOriginPolicyRequest) SetSchemas(v []EnumhttpServletCrossOriginPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddHttpServletCrossOriginPolicyRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *AddHttpServletCrossOriginPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddHttpServletCrossOriginPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddHttpServletCrossOriginPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddHttpServletCrossOriginPolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCorsAllowedMethods

`func (o *AddHttpServletCrossOriginPolicyRequest) GetCorsAllowedMethods() []string`

GetCorsAllowedMethods returns the CorsAllowedMethods field if non-nil, zero value otherwise.

### GetCorsAllowedMethodsOk

`func (o *AddHttpServletCrossOriginPolicyRequest) GetCorsAllowedMethodsOk() (*[]string, bool)`

GetCorsAllowedMethodsOk returns a tuple with the CorsAllowedMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsAllowedMethods

`func (o *AddHttpServletCrossOriginPolicyRequest) SetCorsAllowedMethods(v []string)`

SetCorsAllowedMethods sets CorsAllowedMethods field to given value.

### HasCorsAllowedMethods

`func (o *AddHttpServletCrossOriginPolicyRequest) HasCorsAllowedMethods() bool`

HasCorsAllowedMethods returns a boolean if a field has been set.

### GetCorsAllowedOrigins

`func (o *AddHttpServletCrossOriginPolicyRequest) GetCorsAllowedOrigins() []string`

GetCorsAllowedOrigins returns the CorsAllowedOrigins field if non-nil, zero value otherwise.

### GetCorsAllowedOriginsOk

`func (o *AddHttpServletCrossOriginPolicyRequest) GetCorsAllowedOriginsOk() (*[]string, bool)`

GetCorsAllowedOriginsOk returns a tuple with the CorsAllowedOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsAllowedOrigins

`func (o *AddHttpServletCrossOriginPolicyRequest) SetCorsAllowedOrigins(v []string)`

SetCorsAllowedOrigins sets CorsAllowedOrigins field to given value.

### HasCorsAllowedOrigins

`func (o *AddHttpServletCrossOriginPolicyRequest) HasCorsAllowedOrigins() bool`

HasCorsAllowedOrigins returns a boolean if a field has been set.

### GetCorsExposedHeaders

`func (o *AddHttpServletCrossOriginPolicyRequest) GetCorsExposedHeaders() []string`

GetCorsExposedHeaders returns the CorsExposedHeaders field if non-nil, zero value otherwise.

### GetCorsExposedHeadersOk

`func (o *AddHttpServletCrossOriginPolicyRequest) GetCorsExposedHeadersOk() (*[]string, bool)`

GetCorsExposedHeadersOk returns a tuple with the CorsExposedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsExposedHeaders

`func (o *AddHttpServletCrossOriginPolicyRequest) SetCorsExposedHeaders(v []string)`

SetCorsExposedHeaders sets CorsExposedHeaders field to given value.

### HasCorsExposedHeaders

`func (o *AddHttpServletCrossOriginPolicyRequest) HasCorsExposedHeaders() bool`

HasCorsExposedHeaders returns a boolean if a field has been set.

### GetCorsAllowedHeaders

`func (o *AddHttpServletCrossOriginPolicyRequest) GetCorsAllowedHeaders() []string`

GetCorsAllowedHeaders returns the CorsAllowedHeaders field if non-nil, zero value otherwise.

### GetCorsAllowedHeadersOk

`func (o *AddHttpServletCrossOriginPolicyRequest) GetCorsAllowedHeadersOk() (*[]string, bool)`

GetCorsAllowedHeadersOk returns a tuple with the CorsAllowedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsAllowedHeaders

`func (o *AddHttpServletCrossOriginPolicyRequest) SetCorsAllowedHeaders(v []string)`

SetCorsAllowedHeaders sets CorsAllowedHeaders field to given value.

### HasCorsAllowedHeaders

`func (o *AddHttpServletCrossOriginPolicyRequest) HasCorsAllowedHeaders() bool`

HasCorsAllowedHeaders returns a boolean if a field has been set.

### GetCorsPreflightMaxAge

`func (o *AddHttpServletCrossOriginPolicyRequest) GetCorsPreflightMaxAge() string`

GetCorsPreflightMaxAge returns the CorsPreflightMaxAge field if non-nil, zero value otherwise.

### GetCorsPreflightMaxAgeOk

`func (o *AddHttpServletCrossOriginPolicyRequest) GetCorsPreflightMaxAgeOk() (*string, bool)`

GetCorsPreflightMaxAgeOk returns a tuple with the CorsPreflightMaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsPreflightMaxAge

`func (o *AddHttpServletCrossOriginPolicyRequest) SetCorsPreflightMaxAge(v string)`

SetCorsPreflightMaxAge sets CorsPreflightMaxAge field to given value.

### HasCorsPreflightMaxAge

`func (o *AddHttpServletCrossOriginPolicyRequest) HasCorsPreflightMaxAge() bool`

HasCorsPreflightMaxAge returns a boolean if a field has been set.

### GetCorsAllowCredentials

`func (o *AddHttpServletCrossOriginPolicyRequest) GetCorsAllowCredentials() bool`

GetCorsAllowCredentials returns the CorsAllowCredentials field if non-nil, zero value otherwise.

### GetCorsAllowCredentialsOk

`func (o *AddHttpServletCrossOriginPolicyRequest) GetCorsAllowCredentialsOk() (*bool, bool)`

GetCorsAllowCredentialsOk returns a tuple with the CorsAllowCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsAllowCredentials

`func (o *AddHttpServletCrossOriginPolicyRequest) SetCorsAllowCredentials(v bool)`

SetCorsAllowCredentials sets CorsAllowCredentials field to given value.

### HasCorsAllowCredentials

`func (o *AddHttpServletCrossOriginPolicyRequest) HasCorsAllowCredentials() bool`

HasCorsAllowCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


