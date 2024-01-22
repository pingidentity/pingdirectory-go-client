# Scim2HttpServletExtensionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]Enumscim2HttpServletExtensionSchemaUrn**](Enumscim2HttpServletExtensionSchemaUrn.md) |  | 
**Id** | **string** | Name of the HTTP Servlet Extension | 
**BaseContextPath** | **string** | The context path to use to access the SCIM 2.0 interface. The value must start with a forward slash and must represent a valid HTTP context path. | 
**AccessTokenValidator** | Pointer to **[]string** | If specified, the Access Token Validator(s) that may be used to validate access tokens for requests submitted to this SCIM2 HTTP Servlet Extension. | [optional] 
**MapAccessTokensToLocalUsers** | Pointer to [**EnumhttpServletExtensionMapAccessTokensToLocalUsersProp**](EnumhttpServletExtensionMapAccessTokensToLocalUsersProp.md) |  | [optional] 
**DebugEnabled** | Pointer to **bool** | Enables debug logging of the SCIM 2.0 SDK. Debug messages will be forwarded to the Directory Server debug logger with the scope of com.unboundid.directory.broker.http.scim2.extension.SCIM2HTTPServletExtension. | [optional] 
**DebugLevel** | [**EnumhttpServletExtensionDebugLevelProp**](EnumhttpServletExtensionDebugLevelProp.md) |  | 
**DebugType** | [**[]EnumhttpServletExtensionDebugTypeProp**](EnumhttpServletExtensionDebugTypeProp.md) |  | 
**IncludeStackTrace** | **bool** | Indicates whether a stack trace of the thread which called the debug method should be included in debug log messages. | 
**SwaggerEnabled** | Pointer to **bool** | Indicates whether the SCIM2 HTTP Servlet Extension will generate a Swagger specification document. | [optional] 
**Description** | Pointer to **string** | A description for this HTTP Servlet Extension | [optional] 
**CrossOriginPolicy** | Pointer to **string** | The cross-origin request policy to use for the HTTP Servlet Extension. | [optional] 
**ResponseHeader** | Pointer to **[]string** | Specifies HTTP header fields and values added to response headers for all requests. | [optional] 
**CorrelationIDResponseHeader** | Pointer to **string** | Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \&quot;Correlation-Id\&quot;, \&quot;X-Amzn-Trace-Id\&quot;, and \&quot;X-Request-Id\&quot;. | [optional] 

## Methods

### NewScim2HttpServletExtensionResponse

`func NewScim2HttpServletExtensionResponse(schemas []Enumscim2HttpServletExtensionSchemaUrn, id string, baseContextPath string, debugLevel EnumhttpServletExtensionDebugLevelProp, debugType []EnumhttpServletExtensionDebugTypeProp, includeStackTrace bool, ) *Scim2HttpServletExtensionResponse`

NewScim2HttpServletExtensionResponse instantiates a new Scim2HttpServletExtensionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScim2HttpServletExtensionResponseWithDefaults

`func NewScim2HttpServletExtensionResponseWithDefaults() *Scim2HttpServletExtensionResponse`

NewScim2HttpServletExtensionResponseWithDefaults instantiates a new Scim2HttpServletExtensionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *Scim2HttpServletExtensionResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Scim2HttpServletExtensionResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Scim2HttpServletExtensionResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Scim2HttpServletExtensionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *Scim2HttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *Scim2HttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *Scim2HttpServletExtensionResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *Scim2HttpServletExtensionResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *Scim2HttpServletExtensionResponse) GetSchemas() []Enumscim2HttpServletExtensionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *Scim2HttpServletExtensionResponse) GetSchemasOk() (*[]Enumscim2HttpServletExtensionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *Scim2HttpServletExtensionResponse) SetSchemas(v []Enumscim2HttpServletExtensionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *Scim2HttpServletExtensionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Scim2HttpServletExtensionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Scim2HttpServletExtensionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetBaseContextPath

`func (o *Scim2HttpServletExtensionResponse) GetBaseContextPath() string`

GetBaseContextPath returns the BaseContextPath field if non-nil, zero value otherwise.

### GetBaseContextPathOk

`func (o *Scim2HttpServletExtensionResponse) GetBaseContextPathOk() (*string, bool)`

GetBaseContextPathOk returns a tuple with the BaseContextPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseContextPath

`func (o *Scim2HttpServletExtensionResponse) SetBaseContextPath(v string)`

SetBaseContextPath sets BaseContextPath field to given value.


### GetAccessTokenValidator

`func (o *Scim2HttpServletExtensionResponse) GetAccessTokenValidator() []string`

GetAccessTokenValidator returns the AccessTokenValidator field if non-nil, zero value otherwise.

### GetAccessTokenValidatorOk

`func (o *Scim2HttpServletExtensionResponse) GetAccessTokenValidatorOk() (*[]string, bool)`

GetAccessTokenValidatorOk returns a tuple with the AccessTokenValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValidator

`func (o *Scim2HttpServletExtensionResponse) SetAccessTokenValidator(v []string)`

SetAccessTokenValidator sets AccessTokenValidator field to given value.

### HasAccessTokenValidator

`func (o *Scim2HttpServletExtensionResponse) HasAccessTokenValidator() bool`

HasAccessTokenValidator returns a boolean if a field has been set.

### GetMapAccessTokensToLocalUsers

`func (o *Scim2HttpServletExtensionResponse) GetMapAccessTokensToLocalUsers() EnumhttpServletExtensionMapAccessTokensToLocalUsersProp`

GetMapAccessTokensToLocalUsers returns the MapAccessTokensToLocalUsers field if non-nil, zero value otherwise.

### GetMapAccessTokensToLocalUsersOk

`func (o *Scim2HttpServletExtensionResponse) GetMapAccessTokensToLocalUsersOk() (*EnumhttpServletExtensionMapAccessTokensToLocalUsersProp, bool)`

GetMapAccessTokensToLocalUsersOk returns a tuple with the MapAccessTokensToLocalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapAccessTokensToLocalUsers

`func (o *Scim2HttpServletExtensionResponse) SetMapAccessTokensToLocalUsers(v EnumhttpServletExtensionMapAccessTokensToLocalUsersProp)`

SetMapAccessTokensToLocalUsers sets MapAccessTokensToLocalUsers field to given value.

### HasMapAccessTokensToLocalUsers

`func (o *Scim2HttpServletExtensionResponse) HasMapAccessTokensToLocalUsers() bool`

HasMapAccessTokensToLocalUsers returns a boolean if a field has been set.

### GetDebugEnabled

`func (o *Scim2HttpServletExtensionResponse) GetDebugEnabled() bool`

GetDebugEnabled returns the DebugEnabled field if non-nil, zero value otherwise.

### GetDebugEnabledOk

`func (o *Scim2HttpServletExtensionResponse) GetDebugEnabledOk() (*bool, bool)`

GetDebugEnabledOk returns a tuple with the DebugEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugEnabled

`func (o *Scim2HttpServletExtensionResponse) SetDebugEnabled(v bool)`

SetDebugEnabled sets DebugEnabled field to given value.

### HasDebugEnabled

`func (o *Scim2HttpServletExtensionResponse) HasDebugEnabled() bool`

HasDebugEnabled returns a boolean if a field has been set.

### GetDebugLevel

`func (o *Scim2HttpServletExtensionResponse) GetDebugLevel() EnumhttpServletExtensionDebugLevelProp`

GetDebugLevel returns the DebugLevel field if non-nil, zero value otherwise.

### GetDebugLevelOk

`func (o *Scim2HttpServletExtensionResponse) GetDebugLevelOk() (*EnumhttpServletExtensionDebugLevelProp, bool)`

GetDebugLevelOk returns a tuple with the DebugLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugLevel

`func (o *Scim2HttpServletExtensionResponse) SetDebugLevel(v EnumhttpServletExtensionDebugLevelProp)`

SetDebugLevel sets DebugLevel field to given value.


### GetDebugType

`func (o *Scim2HttpServletExtensionResponse) GetDebugType() []EnumhttpServletExtensionDebugTypeProp`

GetDebugType returns the DebugType field if non-nil, zero value otherwise.

### GetDebugTypeOk

`func (o *Scim2HttpServletExtensionResponse) GetDebugTypeOk() (*[]EnumhttpServletExtensionDebugTypeProp, bool)`

GetDebugTypeOk returns a tuple with the DebugType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugType

`func (o *Scim2HttpServletExtensionResponse) SetDebugType(v []EnumhttpServletExtensionDebugTypeProp)`

SetDebugType sets DebugType field to given value.


### GetIncludeStackTrace

`func (o *Scim2HttpServletExtensionResponse) GetIncludeStackTrace() bool`

GetIncludeStackTrace returns the IncludeStackTrace field if non-nil, zero value otherwise.

### GetIncludeStackTraceOk

`func (o *Scim2HttpServletExtensionResponse) GetIncludeStackTraceOk() (*bool, bool)`

GetIncludeStackTraceOk returns a tuple with the IncludeStackTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeStackTrace

`func (o *Scim2HttpServletExtensionResponse) SetIncludeStackTrace(v bool)`

SetIncludeStackTrace sets IncludeStackTrace field to given value.


### GetSwaggerEnabled

`func (o *Scim2HttpServletExtensionResponse) GetSwaggerEnabled() bool`

GetSwaggerEnabled returns the SwaggerEnabled field if non-nil, zero value otherwise.

### GetSwaggerEnabledOk

`func (o *Scim2HttpServletExtensionResponse) GetSwaggerEnabledOk() (*bool, bool)`

GetSwaggerEnabledOk returns a tuple with the SwaggerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwaggerEnabled

`func (o *Scim2HttpServletExtensionResponse) SetSwaggerEnabled(v bool)`

SetSwaggerEnabled sets SwaggerEnabled field to given value.

### HasSwaggerEnabled

`func (o *Scim2HttpServletExtensionResponse) HasSwaggerEnabled() bool`

HasSwaggerEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *Scim2HttpServletExtensionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Scim2HttpServletExtensionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Scim2HttpServletExtensionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Scim2HttpServletExtensionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCrossOriginPolicy

`func (o *Scim2HttpServletExtensionResponse) GetCrossOriginPolicy() string`

GetCrossOriginPolicy returns the CrossOriginPolicy field if non-nil, zero value otherwise.

### GetCrossOriginPolicyOk

`func (o *Scim2HttpServletExtensionResponse) GetCrossOriginPolicyOk() (*string, bool)`

GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossOriginPolicy

`func (o *Scim2HttpServletExtensionResponse) SetCrossOriginPolicy(v string)`

SetCrossOriginPolicy sets CrossOriginPolicy field to given value.

### HasCrossOriginPolicy

`func (o *Scim2HttpServletExtensionResponse) HasCrossOriginPolicy() bool`

HasCrossOriginPolicy returns a boolean if a field has been set.

### GetResponseHeader

`func (o *Scim2HttpServletExtensionResponse) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *Scim2HttpServletExtensionResponse) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *Scim2HttpServletExtensionResponse) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *Scim2HttpServletExtensionResponse) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetCorrelationIDResponseHeader

`func (o *Scim2HttpServletExtensionResponse) GetCorrelationIDResponseHeader() string`

GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field if non-nil, zero value otherwise.

### GetCorrelationIDResponseHeaderOk

`func (o *Scim2HttpServletExtensionResponse) GetCorrelationIDResponseHeaderOk() (*string, bool)`

GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDResponseHeader

`func (o *Scim2HttpServletExtensionResponse) SetCorrelationIDResponseHeader(v string)`

SetCorrelationIDResponseHeader sets CorrelationIDResponseHeader field to given value.

### HasCorrelationIDResponseHeader

`func (o *Scim2HttpServletExtensionResponse) HasCorrelationIDResponseHeader() bool`

HasCorrelationIDResponseHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


