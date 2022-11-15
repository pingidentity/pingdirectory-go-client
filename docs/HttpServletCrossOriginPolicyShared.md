# HttpServletCrossOriginPolicyShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumhttpServletCrossOriginPolicySchemaUrn**](EnumhttpServletCrossOriginPolicySchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this HTTP Servlet Cross Origin Policy | [optional] 
**CorsAllowedMethods** | Pointer to **[]string** |  | [optional] 
**CorsAllowedOrigins** | Pointer to **[]string** |  | [optional] 
**CorsExposedHeaders** | Pointer to **[]string** |  | [optional] 
**CorsAllowedHeaders** | Pointer to **[]string** |  | [optional] 
**CorsPreflightMaxAge** | Pointer to **string** | The maximum amount of time that a preflight request can be cached by a client. | [optional] 
**CorsAllowCredentials** | Pointer to **bool** | Indicates whether the servlet extension allows CORS requests with username/password credentials. | [optional] 

## Methods

### NewHttpServletCrossOriginPolicyShared

`func NewHttpServletCrossOriginPolicyShared() *HttpServletCrossOriginPolicyShared`

NewHttpServletCrossOriginPolicyShared instantiates a new HttpServletCrossOriginPolicyShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpServletCrossOriginPolicySharedWithDefaults

`func NewHttpServletCrossOriginPolicySharedWithDefaults() *HttpServletCrossOriginPolicyShared`

NewHttpServletCrossOriginPolicySharedWithDefaults instantiates a new HttpServletCrossOriginPolicyShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *HttpServletCrossOriginPolicyShared) GetSchemas() []EnumhttpServletCrossOriginPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *HttpServletCrossOriginPolicyShared) GetSchemasOk() (*[]EnumhttpServletCrossOriginPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *HttpServletCrossOriginPolicyShared) SetSchemas(v []EnumhttpServletCrossOriginPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *HttpServletCrossOriginPolicyShared) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *HttpServletCrossOriginPolicyShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HttpServletCrossOriginPolicyShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HttpServletCrossOriginPolicyShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HttpServletCrossOriginPolicyShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCorsAllowedMethods

`func (o *HttpServletCrossOriginPolicyShared) GetCorsAllowedMethods() []string`

GetCorsAllowedMethods returns the CorsAllowedMethods field if non-nil, zero value otherwise.

### GetCorsAllowedMethodsOk

`func (o *HttpServletCrossOriginPolicyShared) GetCorsAllowedMethodsOk() (*[]string, bool)`

GetCorsAllowedMethodsOk returns a tuple with the CorsAllowedMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsAllowedMethods

`func (o *HttpServletCrossOriginPolicyShared) SetCorsAllowedMethods(v []string)`

SetCorsAllowedMethods sets CorsAllowedMethods field to given value.

### HasCorsAllowedMethods

`func (o *HttpServletCrossOriginPolicyShared) HasCorsAllowedMethods() bool`

HasCorsAllowedMethods returns a boolean if a field has been set.

### GetCorsAllowedOrigins

`func (o *HttpServletCrossOriginPolicyShared) GetCorsAllowedOrigins() []string`

GetCorsAllowedOrigins returns the CorsAllowedOrigins field if non-nil, zero value otherwise.

### GetCorsAllowedOriginsOk

`func (o *HttpServletCrossOriginPolicyShared) GetCorsAllowedOriginsOk() (*[]string, bool)`

GetCorsAllowedOriginsOk returns a tuple with the CorsAllowedOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsAllowedOrigins

`func (o *HttpServletCrossOriginPolicyShared) SetCorsAllowedOrigins(v []string)`

SetCorsAllowedOrigins sets CorsAllowedOrigins field to given value.

### HasCorsAllowedOrigins

`func (o *HttpServletCrossOriginPolicyShared) HasCorsAllowedOrigins() bool`

HasCorsAllowedOrigins returns a boolean if a field has been set.

### GetCorsExposedHeaders

`func (o *HttpServletCrossOriginPolicyShared) GetCorsExposedHeaders() []string`

GetCorsExposedHeaders returns the CorsExposedHeaders field if non-nil, zero value otherwise.

### GetCorsExposedHeadersOk

`func (o *HttpServletCrossOriginPolicyShared) GetCorsExposedHeadersOk() (*[]string, bool)`

GetCorsExposedHeadersOk returns a tuple with the CorsExposedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsExposedHeaders

`func (o *HttpServletCrossOriginPolicyShared) SetCorsExposedHeaders(v []string)`

SetCorsExposedHeaders sets CorsExposedHeaders field to given value.

### HasCorsExposedHeaders

`func (o *HttpServletCrossOriginPolicyShared) HasCorsExposedHeaders() bool`

HasCorsExposedHeaders returns a boolean if a field has been set.

### GetCorsAllowedHeaders

`func (o *HttpServletCrossOriginPolicyShared) GetCorsAllowedHeaders() []string`

GetCorsAllowedHeaders returns the CorsAllowedHeaders field if non-nil, zero value otherwise.

### GetCorsAllowedHeadersOk

`func (o *HttpServletCrossOriginPolicyShared) GetCorsAllowedHeadersOk() (*[]string, bool)`

GetCorsAllowedHeadersOk returns a tuple with the CorsAllowedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsAllowedHeaders

`func (o *HttpServletCrossOriginPolicyShared) SetCorsAllowedHeaders(v []string)`

SetCorsAllowedHeaders sets CorsAllowedHeaders field to given value.

### HasCorsAllowedHeaders

`func (o *HttpServletCrossOriginPolicyShared) HasCorsAllowedHeaders() bool`

HasCorsAllowedHeaders returns a boolean if a field has been set.

### GetCorsPreflightMaxAge

`func (o *HttpServletCrossOriginPolicyShared) GetCorsPreflightMaxAge() string`

GetCorsPreflightMaxAge returns the CorsPreflightMaxAge field if non-nil, zero value otherwise.

### GetCorsPreflightMaxAgeOk

`func (o *HttpServletCrossOriginPolicyShared) GetCorsPreflightMaxAgeOk() (*string, bool)`

GetCorsPreflightMaxAgeOk returns a tuple with the CorsPreflightMaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsPreflightMaxAge

`func (o *HttpServletCrossOriginPolicyShared) SetCorsPreflightMaxAge(v string)`

SetCorsPreflightMaxAge sets CorsPreflightMaxAge field to given value.

### HasCorsPreflightMaxAge

`func (o *HttpServletCrossOriginPolicyShared) HasCorsPreflightMaxAge() bool`

HasCorsPreflightMaxAge returns a boolean if a field has been set.

### GetCorsAllowCredentials

`func (o *HttpServletCrossOriginPolicyShared) GetCorsAllowCredentials() bool`

GetCorsAllowCredentials returns the CorsAllowCredentials field if non-nil, zero value otherwise.

### GetCorsAllowCredentialsOk

`func (o *HttpServletCrossOriginPolicyShared) GetCorsAllowCredentialsOk() (*bool, bool)`

GetCorsAllowCredentialsOk returns a tuple with the CorsAllowCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsAllowCredentials

`func (o *HttpServletCrossOriginPolicyShared) SetCorsAllowCredentials(v bool)`

SetCorsAllowCredentials sets CorsAllowCredentials field to given value.

### HasCorsAllowCredentials

`func (o *HttpServletCrossOriginPolicyShared) HasCorsAllowCredentials() bool`

HasCorsAllowCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


