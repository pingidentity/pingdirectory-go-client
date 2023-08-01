# HttpServletExtensionListResponseResourcesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsidebandApiHttpServletExtensionSchemaUrn**](EnumsidebandApiHttpServletExtensionSchemaUrn.md) |  | 
**Id** | **string** | Name of the HTTP Servlet Extension | 
**Description** | Pointer to **string** | A description for this HTTP Servlet Extension | [optional] 
**CrossOriginPolicy** | Pointer to **string** | The cross-origin request policy to use for the HTTP Servlet Extension. | [optional] 
**ResponseHeader** | Pointer to **[]string** | Specifies HTTP header fields and values added to response headers for all requests. | [optional] 
**CorrelationIDResponseHeader** | Pointer to **string** | Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \&quot;Correlation-Id\&quot;, \&quot;X-Amzn-Trace-Id\&quot;, and \&quot;X-Request-Id\&quot;. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**BasicAuthEnabled** | Pointer to **bool** | Enables HTTP Basic authentication, using a username and password. The Identity Mapper specified by the identity-mapper property will be used to map the username to a DN. | [optional] 
**IdentityMapper** | Pointer to **string** | Specifies the Identity Mapper that is to be used for associating user entries with basic authentication usernames. | [optional] 
**AccessTokenValidator** | Pointer to **[]string** | If specified, the Access Token Validator(s) that may be used to validate access tokens for requests submitted to this Directory REST API HTTP Servlet Extension. | [optional] 
**AccessTokenScope** | Pointer to **string** | The name of a scope that must be present in an access token accepted by the Directory REST API HTTP Servlet Extension. | [optional] 
**Audience** | Pointer to **string** | A string or URI that identifies the Directory REST API HTTP Servlet Extension in the context of OAuth2 authorization. | [optional] 
**Server** | Pointer to **string** | Specifies the PingFederate server to be configured. | [optional] 
**BaseContextPath** | **string** | The context path to use to access the SCIM 2.0 interface. The value must start with a forward slash and must represent a valid HTTP context path. | 
**AvailableStatusCode** | **int64** | Specifies the HTTP status code that the servlet should return if the server considers itself to be available. | 
**DegradedStatusCode** | **int64** | Specifies the HTTP status code that the servlet should return if the server considers itself to be degraded. | 
**UnavailableStatusCode** | **int64** | Specifies the HTTP status code that the servlet should return if the server considers itself to be unavailable. | 
**OverrideStatusCode** | Pointer to **int64** | Specifies a HTTP status code that the servlet should always return, regardless of the server&#39;s availability. If this value is defined, it will override the availability-based return codes. | [optional] 
**IncludeResponseBody** | Pointer to **bool** | Indicates whether the response should include a body that is a JSON object. | [optional] 
**AdditionalResponseContents** | Pointer to **string** | A JSON-formatted string containing additional fields to be returned in the response body. For example, an additional-response-contents value of &#39;{ \&quot;key\&quot;: \&quot;value\&quot; }&#39; would add the key and value to the root of the JSON response body. | [optional] 
**IncludeInstanceNameLabel** | Pointer to **bool** | Indicates whether generated metrics should include an \&quot;instance\&quot; label whose value is the instance name for this Directory Server instance. | [optional] 
**IncludeProductNameLabel** | Pointer to **bool** | Indicates whether generated metrics should include a \&quot;product\&quot; label whose value is the product name for this Directory Server instance. | [optional] 
**IncludeLocationNameLabel** | Pointer to **bool** | Indicates whether generated metrics should include a \&quot;location\&quot; label whose value is the location name for this Directory Server instance. | [optional] 
**AlwaysIncludeMonitorEntryNameLabel** | Pointer to **bool** | Indicates whether generated metrics should always include a \&quot;monitor_entry\&quot; label whose value is the name of the monitor entry from which the metric was obtained. | [optional] 
**IncludeMonitorObjectClassNameLabel** | Pointer to **bool** | Indicates whether generated metrics should include a \&quot;monitor_object_class\&quot; label whose value is the name of the object class for the monitor entry from which the metric was obtained. | [optional] 
**IncludeMonitorAttributeNameLabel** | Pointer to **bool** | Indicates whether generated metrics should include a \&quot;monitor_attribute\&quot; label whose value is the name of the monitor attribute from which the metric was obtained. | [optional] 
**LabelNameValuePair** | Pointer to **[]string** | A set of name-value pairs for labels that should be included in all metrics exposed by this Directory Server instance. | [optional] 
**StaticContextPath** | Pointer to **string** | The path below the base context path by which static, non-template content such as images, CSS, and Javascript files are accessible. | [optional] 
**StaticContentDirectory** | Pointer to **string** | Specifies the base directory in which static, non-template content such as images, CSS, and Javascript files are stored on the filesystem. | [optional] 
**StaticCustomDirectory** | Pointer to **string** | Specifies the base directory in which custom static, non-template content such as images, CSS, and Javascript files are stored on the filesystem. Files in this directory will override those with the same name in the directory specified by the static-content-directory property. | [optional] 
**TemplateDirectory** | **[]string** | Specifies an ordered list of directories in which to search for the template files. | 
**ExposeRequestAttributes** | Pointer to **bool** | Specifies whether the HTTP request will be exposed to templates. | [optional] 
**ExposeSessionAttributes** | Pointer to **bool** | Specifies whether the HTTP session will be exposed to templates. | [optional] 
**ExposeServerContext** | Pointer to **bool** | Specifies whether a server context will be exposed under context key &#39;ubid_server&#39; for all template contexts. | [optional] 
**AllowContextOverride** | Pointer to **bool** | Indicates whether context providers may override existing context objects with new values. | [optional] 
**MimeTypesFile** | Pointer to **string** | Specifies the path to a file that contains MIME type mappings that will be used to determine the appropriate value to return for the Content-Type header based on the extension of the requested file. | [optional] 
**DefaultMIMEType** | Pointer to **string** | Specifies the default MIME type to use for the Content-Type header when a mapping cannot be found. | [optional] 
**CharacterEncoding** | Pointer to **string** | Specifies the value that will be used for all responses&#39; Content-Type headers&#39; charset parameter that indicates the character encoding of the document. | [optional] 
**StaticResponseHeader** | Pointer to **[]string** | Specifies HTTP header fields and values added to response headers for static content requests such as images and scripts. | [optional] 
**RequireAuthentication** | Pointer to **bool** | Indicates whether the servlet extension should only accept requests from authenticated clients. | [optional] 
**BearerTokenAuthEnabled** | Pointer to **bool** | Enables HTTP bearer token authentication. | [optional] 
**OAuthTokenHandler** | Pointer to **string** | Specifies the OAuth Token Handler implementation that should be used to validate OAuth 2.0 bearer tokens when they are included in a SCIM request. | [optional] 
**ResourceMappingFile** | Pointer to **string** | The path to an XML file defining the resources supported by the SCIM interface and the SCIM-to-LDAP attribute mappings to use. | [optional] 
**IncludeLDAPObjectclass** | Pointer to **[]string** | Specifies the LDAP object classes that should be exposed directly as SCIM resources. | [optional] 
**ExcludeLDAPObjectclass** | Pointer to **[]string** | Specifies the LDAP object classes that should be not be exposed directly as SCIM resources. | [optional] 
**IncludeLDAPBaseDN** | Pointer to **[]string** | Specifies the base DNs for the branches of the DIT that should be exposed via the Identity Access API. | [optional] 
**ExcludeLDAPBaseDN** | Pointer to **[]string** | Specifies the base DNs for the branches of the DIT that should not be exposed via the Identity Access API. | [optional] 
**EntityTagLDAPAttribute** | Pointer to **string** | Specifies the LDAP attribute whose value should be used as the entity tag value to enable SCIM resource versioning support. | [optional] 
**TemporaryDirectory** | **string** | Specifies the location of the directory that is used to create temporary files containing SCIM request data. | 
**TemporaryDirectoryPermissions** | **string** | Specifies the permissions that should be applied to the directory that is used to create temporary files. | 
**MaxResults** | Pointer to **int64** | The maximum number of resources that are returned in a response. | [optional] 
**BulkMaxOperations** | Pointer to **int64** | The maximum number of operations that are permitted in a bulk request. | [optional] 
**BulkMaxPayloadSize** | Pointer to **string** | The maximum payload size in bytes of a bulk request. | [optional] 
**BulkMaxConcurrentRequests** | Pointer to **int64** | The maximum number of bulk requests that may be processed concurrently by the server. Any bulk request that would cause this limit to be exceeded is rejected with HTTP status code 503. | [optional] 
**DebugEnabled** | Pointer to **bool** | Enables debug logging of the SCIM 2.0 SDK. Debug messages will be forwarded to the Directory Server debug logger with the scope of com.unboundid.directory.broker.http.scim2.extension.SCIM2HTTPServletExtension. | [optional] 
**DebugLevel** | [**EnumhttpServletExtensionDebugLevelProp**](EnumhttpServletExtensionDebugLevelProp.md) |  | 
**DebugType** | [**[]EnumhttpServletExtensionDebugTypeProp**](EnumhttpServletExtensionDebugTypeProp.md) |  | 
**IncludeStackTrace** | **bool** | Indicates whether a stack trace of the thread which called the debug method should be included in debug log messages. | 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted HTTP Servlet Extension. | 
**ScriptArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Scripted HTTP Servlet Extension. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**FapiFinancialID** | Pointer to **string** | The unique financial id of the ASPSP associated with this Open Banking service. | [optional] 
**PathPrefix** | Pointer to **string** | An optional ASPSP-specific path prefix to include in the base URI path. If specified with the value \&quot;myPrefix\&quot;, for example, the base path will be /myPrefix/open-banking/v1.1/. | [optional] 
**ConsentServer** | Pointer to **string** | Specifies the PingDirectory instance that is hosting the Consent API for storage and retrieval of Account Requests. | [optional] 
**ConsentDefinitionID** | **string** | The name/id of the consent definition, as defined in the consent-server, that is used for storing Account Requests. | 
**DocumentRootDirectory** | **string** | Specifies the path to the directory on the local filesystem containing the files to be served by this File Server HTTP Servlet Extension. The path must exist, and it must be a directory. | 
**EnableDirectoryIndexing** | Pointer to **bool** | Indicates whether to generate a default HTML page with a listing of available files if the requested path refers to a directory rather than a file, and that directory does not contain an index file. | [optional] 
**IndexFile** | Pointer to **[]string** | Specifies the name of a file whose contents may be returned to the client if the requested path refers to a directory rather than a file. | [optional] 
**AllowedAuthenticationType** | Pointer to [**[]EnumhttpServletExtensionAllowedAuthenticationTypeProp**](EnumhttpServletExtensionAllowedAuthenticationTypeProp.md) |  | [optional] 
**IdTokenValidator** | Pointer to **[]string** | The ID token validators that may be used to verify the authenticity of an of an OpenID Connect ID token. | [optional] 
**RequireFileServletAccessPrivilege** | Pointer to **bool** | Indicates whether the servlet extension should only accept requests from authenticated clients that have the file-servlet-access privilege. | [optional] 
**RequireGroup** | Pointer to **[]string** | The DN of a group whose members will be permitted to access to the associated files. If multiple group DNs are configured, then anyone who is a member of at least one of those groups will be granted access. | [optional] 
**SharedSecretHeaderName** | **string** | The request header used to find the shared secret header for incoming Sideband API HTTP Servlet Extension requests. | 
**SharedSecrets** | Pointer to **[]string** | Shared secrets between the third-party API Gateway and Sideband API HTTP Servlet Extension. | [optional] 
**RequireAPIAuthentication** | Pointer to **bool** | Require authentication when accessing the REST API. | [optional] 
**OmitErrorMessageDetails** | Pointer to **bool** | Specifies that API error messages for invalid queries, unknown resources, service unavailable, and internal server errors are generic in nature. | [optional] 
**ApiAuthenticationTimeout** | Pointer to **string** | Length of time before a REST API authentication session expires. | [optional] 
**MapAccessTokensToLocalUsers** | Pointer to [**EnumhttpServletExtensionMapAccessTokensToLocalUsersProp**](EnumhttpServletExtensionMapAccessTokensToLocalUsersProp.md) |  | [optional] 
**SwaggerEnabled** | Pointer to **bool** | Indicates whether the SCIM2 HTTP Servlet Extension will generate a Swagger specification document. | [optional] 
**MaxPageSize** | Pointer to **int64** | The maximum number of entries to be returned in one page of search results. | [optional] 
**SchemasEndpointObjectclass** | Pointer to **[]string** | The list of object classes which will be returned by the schemas endpoint. | [optional] 
**DefaultOperationalAttribute** | Pointer to **[]string** | A set of operational attributes that will be returned with entries by default. | [optional] 
**RejectExpansionAttribute** | Pointer to **[]string** | A set of attributes which the client is not allowed to provide for the expand query parameters. This should be used for attributes that could either have a large number of values or that reference entries that are very large like groups. | [optional] 
**AlwaysUsePermissiveModify** | Pointer to **bool** | Indicates whether to always use permissive modify behavior for PATCH requests, even if the request did not include the permissive modify request control. | [optional] 
**AllowedControl** | Pointer to [**[]EnumhttpServletExtensionAllowedControlProp**](EnumhttpServletExtensionAllowedControlProp.md) |  | [optional] 
**ExcludedAPIEndpoint** | Pointer to **[]string** | Specifies any Gateway API Endpoints that will not be handled by the Gateway HTTP Servlet Extension. | [optional] 
**RequestLimit** | Pointer to **string** | The maximum number of bytes allowed per request body. | [optional] 
**ResponseLimit** | Pointer to **string** | The maximum number of bytes allowed per response body. | [optional] 
**NumForwardThreads** | Pointer to **int64** | The number of threads used to forward responses to the API backend. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party HTTP Servlet Extension. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party HTTP Servlet Extension. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**RequestContextMethod** | Pointer to [**EnumhttpServletExtensionRequestContextMethodProp**](EnumhttpServletExtensionRequestContextMethodProp.md) |  | [optional] 

## Methods

### NewHttpServletExtensionListResponseResourcesInner

`func NewHttpServletExtensionListResponseResourcesInner(schemas []EnumsidebandApiHttpServletExtensionSchemaUrn, id string, baseContextPath string, availableStatusCode int64, degradedStatusCode int64, unavailableStatusCode int64, templateDirectory []string, temporaryDirectory string, temporaryDirectoryPermissions string, debugLevel EnumhttpServletExtensionDebugLevelProp, debugType []EnumhttpServletExtensionDebugTypeProp, includeStackTrace bool, scriptClass string, consentDefinitionID string, documentRootDirectory string, sharedSecretHeaderName string, extensionClass string, ) *HttpServletExtensionListResponseResourcesInner`

NewHttpServletExtensionListResponseResourcesInner instantiates a new HttpServletExtensionListResponseResourcesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpServletExtensionListResponseResourcesInnerWithDefaults

`func NewHttpServletExtensionListResponseResourcesInnerWithDefaults() *HttpServletExtensionListResponseResourcesInner`

NewHttpServletExtensionListResponseResourcesInnerWithDefaults instantiates a new HttpServletExtensionListResponseResourcesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *HttpServletExtensionListResponseResourcesInner) GetSchemas() []EnumsidebandApiHttpServletExtensionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetSchemasOk() (*[]EnumsidebandApiHttpServletExtensionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *HttpServletExtensionListResponseResourcesInner) SetSchemas(v []EnumsidebandApiHttpServletExtensionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *HttpServletExtensionListResponseResourcesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HttpServletExtensionListResponseResourcesInner) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *HttpServletExtensionListResponseResourcesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HttpServletExtensionListResponseResourcesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HttpServletExtensionListResponseResourcesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCrossOriginPolicy

`func (o *HttpServletExtensionListResponseResourcesInner) GetCrossOriginPolicy() string`

GetCrossOriginPolicy returns the CrossOriginPolicy field if non-nil, zero value otherwise.

### GetCrossOriginPolicyOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetCrossOriginPolicyOk() (*string, bool)`

GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossOriginPolicy

`func (o *HttpServletExtensionListResponseResourcesInner) SetCrossOriginPolicy(v string)`

SetCrossOriginPolicy sets CrossOriginPolicy field to given value.

### HasCrossOriginPolicy

`func (o *HttpServletExtensionListResponseResourcesInner) HasCrossOriginPolicy() bool`

HasCrossOriginPolicy returns a boolean if a field has been set.

### GetResponseHeader

`func (o *HttpServletExtensionListResponseResourcesInner) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *HttpServletExtensionListResponseResourcesInner) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *HttpServletExtensionListResponseResourcesInner) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetCorrelationIDResponseHeader

`func (o *HttpServletExtensionListResponseResourcesInner) GetCorrelationIDResponseHeader() string`

GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field if non-nil, zero value otherwise.

### GetCorrelationIDResponseHeaderOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetCorrelationIDResponseHeaderOk() (*string, bool)`

GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDResponseHeader

`func (o *HttpServletExtensionListResponseResourcesInner) SetCorrelationIDResponseHeader(v string)`

SetCorrelationIDResponseHeader sets CorrelationIDResponseHeader field to given value.

### HasCorrelationIDResponseHeader

`func (o *HttpServletExtensionListResponseResourcesInner) HasCorrelationIDResponseHeader() bool`

HasCorrelationIDResponseHeader returns a boolean if a field has been set.

### GetMeta

`func (o *HttpServletExtensionListResponseResourcesInner) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *HttpServletExtensionListResponseResourcesInner) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *HttpServletExtensionListResponseResourcesInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *HttpServletExtensionListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *HttpServletExtensionListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *HttpServletExtensionListResponseResourcesInner) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *HttpServletExtensionListResponseResourcesInner) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetBasicAuthEnabled

`func (o *HttpServletExtensionListResponseResourcesInner) GetBasicAuthEnabled() bool`

GetBasicAuthEnabled returns the BasicAuthEnabled field if non-nil, zero value otherwise.

### GetBasicAuthEnabledOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetBasicAuthEnabledOk() (*bool, bool)`

GetBasicAuthEnabledOk returns a tuple with the BasicAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthEnabled

`func (o *HttpServletExtensionListResponseResourcesInner) SetBasicAuthEnabled(v bool)`

SetBasicAuthEnabled sets BasicAuthEnabled field to given value.

### HasBasicAuthEnabled

`func (o *HttpServletExtensionListResponseResourcesInner) HasBasicAuthEnabled() bool`

HasBasicAuthEnabled returns a boolean if a field has been set.

### GetIdentityMapper

`func (o *HttpServletExtensionListResponseResourcesInner) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *HttpServletExtensionListResponseResourcesInner) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.

### HasIdentityMapper

`func (o *HttpServletExtensionListResponseResourcesInner) HasIdentityMapper() bool`

HasIdentityMapper returns a boolean if a field has been set.

### GetAccessTokenValidator

`func (o *HttpServletExtensionListResponseResourcesInner) GetAccessTokenValidator() []string`

GetAccessTokenValidator returns the AccessTokenValidator field if non-nil, zero value otherwise.

### GetAccessTokenValidatorOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetAccessTokenValidatorOk() (*[]string, bool)`

GetAccessTokenValidatorOk returns a tuple with the AccessTokenValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValidator

`func (o *HttpServletExtensionListResponseResourcesInner) SetAccessTokenValidator(v []string)`

SetAccessTokenValidator sets AccessTokenValidator field to given value.

### HasAccessTokenValidator

`func (o *HttpServletExtensionListResponseResourcesInner) HasAccessTokenValidator() bool`

HasAccessTokenValidator returns a boolean if a field has been set.

### GetAccessTokenScope

`func (o *HttpServletExtensionListResponseResourcesInner) GetAccessTokenScope() string`

GetAccessTokenScope returns the AccessTokenScope field if non-nil, zero value otherwise.

### GetAccessTokenScopeOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetAccessTokenScopeOk() (*string, bool)`

GetAccessTokenScopeOk returns a tuple with the AccessTokenScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenScope

`func (o *HttpServletExtensionListResponseResourcesInner) SetAccessTokenScope(v string)`

SetAccessTokenScope sets AccessTokenScope field to given value.

### HasAccessTokenScope

`func (o *HttpServletExtensionListResponseResourcesInner) HasAccessTokenScope() bool`

HasAccessTokenScope returns a boolean if a field has been set.

### GetAudience

`func (o *HttpServletExtensionListResponseResourcesInner) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *HttpServletExtensionListResponseResourcesInner) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *HttpServletExtensionListResponseResourcesInner) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetServer

`func (o *HttpServletExtensionListResponseResourcesInner) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *HttpServletExtensionListResponseResourcesInner) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *HttpServletExtensionListResponseResourcesInner) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetBaseContextPath

`func (o *HttpServletExtensionListResponseResourcesInner) GetBaseContextPath() string`

GetBaseContextPath returns the BaseContextPath field if non-nil, zero value otherwise.

### GetBaseContextPathOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetBaseContextPathOk() (*string, bool)`

GetBaseContextPathOk returns a tuple with the BaseContextPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseContextPath

`func (o *HttpServletExtensionListResponseResourcesInner) SetBaseContextPath(v string)`

SetBaseContextPath sets BaseContextPath field to given value.


### GetAvailableStatusCode

`func (o *HttpServletExtensionListResponseResourcesInner) GetAvailableStatusCode() int64`

GetAvailableStatusCode returns the AvailableStatusCode field if non-nil, zero value otherwise.

### GetAvailableStatusCodeOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetAvailableStatusCodeOk() (*int64, bool)`

GetAvailableStatusCodeOk returns a tuple with the AvailableStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableStatusCode

`func (o *HttpServletExtensionListResponseResourcesInner) SetAvailableStatusCode(v int64)`

SetAvailableStatusCode sets AvailableStatusCode field to given value.


### GetDegradedStatusCode

`func (o *HttpServletExtensionListResponseResourcesInner) GetDegradedStatusCode() int64`

GetDegradedStatusCode returns the DegradedStatusCode field if non-nil, zero value otherwise.

### GetDegradedStatusCodeOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetDegradedStatusCodeOk() (*int64, bool)`

GetDegradedStatusCodeOk returns a tuple with the DegradedStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegradedStatusCode

`func (o *HttpServletExtensionListResponseResourcesInner) SetDegradedStatusCode(v int64)`

SetDegradedStatusCode sets DegradedStatusCode field to given value.


### GetUnavailableStatusCode

`func (o *HttpServletExtensionListResponseResourcesInner) GetUnavailableStatusCode() int64`

GetUnavailableStatusCode returns the UnavailableStatusCode field if non-nil, zero value otherwise.

### GetUnavailableStatusCodeOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetUnavailableStatusCodeOk() (*int64, bool)`

GetUnavailableStatusCodeOk returns a tuple with the UnavailableStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailableStatusCode

`func (o *HttpServletExtensionListResponseResourcesInner) SetUnavailableStatusCode(v int64)`

SetUnavailableStatusCode sets UnavailableStatusCode field to given value.


### GetOverrideStatusCode

`func (o *HttpServletExtensionListResponseResourcesInner) GetOverrideStatusCode() int64`

GetOverrideStatusCode returns the OverrideStatusCode field if non-nil, zero value otherwise.

### GetOverrideStatusCodeOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetOverrideStatusCodeOk() (*int64, bool)`

GetOverrideStatusCodeOk returns a tuple with the OverrideStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideStatusCode

`func (o *HttpServletExtensionListResponseResourcesInner) SetOverrideStatusCode(v int64)`

SetOverrideStatusCode sets OverrideStatusCode field to given value.

### HasOverrideStatusCode

`func (o *HttpServletExtensionListResponseResourcesInner) HasOverrideStatusCode() bool`

HasOverrideStatusCode returns a boolean if a field has been set.

### GetIncludeResponseBody

`func (o *HttpServletExtensionListResponseResourcesInner) GetIncludeResponseBody() bool`

GetIncludeResponseBody returns the IncludeResponseBody field if non-nil, zero value otherwise.

### GetIncludeResponseBodyOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetIncludeResponseBodyOk() (*bool, bool)`

GetIncludeResponseBodyOk returns a tuple with the IncludeResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeResponseBody

`func (o *HttpServletExtensionListResponseResourcesInner) SetIncludeResponseBody(v bool)`

SetIncludeResponseBody sets IncludeResponseBody field to given value.

### HasIncludeResponseBody

`func (o *HttpServletExtensionListResponseResourcesInner) HasIncludeResponseBody() bool`

HasIncludeResponseBody returns a boolean if a field has been set.

### GetAdditionalResponseContents

`func (o *HttpServletExtensionListResponseResourcesInner) GetAdditionalResponseContents() string`

GetAdditionalResponseContents returns the AdditionalResponseContents field if non-nil, zero value otherwise.

### GetAdditionalResponseContentsOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetAdditionalResponseContentsOk() (*string, bool)`

GetAdditionalResponseContentsOk returns a tuple with the AdditionalResponseContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalResponseContents

`func (o *HttpServletExtensionListResponseResourcesInner) SetAdditionalResponseContents(v string)`

SetAdditionalResponseContents sets AdditionalResponseContents field to given value.

### HasAdditionalResponseContents

`func (o *HttpServletExtensionListResponseResourcesInner) HasAdditionalResponseContents() bool`

HasAdditionalResponseContents returns a boolean if a field has been set.

### GetIncludeInstanceNameLabel

`func (o *HttpServletExtensionListResponseResourcesInner) GetIncludeInstanceNameLabel() bool`

GetIncludeInstanceNameLabel returns the IncludeInstanceNameLabel field if non-nil, zero value otherwise.

### GetIncludeInstanceNameLabelOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetIncludeInstanceNameLabelOk() (*bool, bool)`

GetIncludeInstanceNameLabelOk returns a tuple with the IncludeInstanceNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInstanceNameLabel

`func (o *HttpServletExtensionListResponseResourcesInner) SetIncludeInstanceNameLabel(v bool)`

SetIncludeInstanceNameLabel sets IncludeInstanceNameLabel field to given value.

### HasIncludeInstanceNameLabel

`func (o *HttpServletExtensionListResponseResourcesInner) HasIncludeInstanceNameLabel() bool`

HasIncludeInstanceNameLabel returns a boolean if a field has been set.

### GetIncludeProductNameLabel

`func (o *HttpServletExtensionListResponseResourcesInner) GetIncludeProductNameLabel() bool`

GetIncludeProductNameLabel returns the IncludeProductNameLabel field if non-nil, zero value otherwise.

### GetIncludeProductNameLabelOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetIncludeProductNameLabelOk() (*bool, bool)`

GetIncludeProductNameLabelOk returns a tuple with the IncludeProductNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeProductNameLabel

`func (o *HttpServletExtensionListResponseResourcesInner) SetIncludeProductNameLabel(v bool)`

SetIncludeProductNameLabel sets IncludeProductNameLabel field to given value.

### HasIncludeProductNameLabel

`func (o *HttpServletExtensionListResponseResourcesInner) HasIncludeProductNameLabel() bool`

HasIncludeProductNameLabel returns a boolean if a field has been set.

### GetIncludeLocationNameLabel

`func (o *HttpServletExtensionListResponseResourcesInner) GetIncludeLocationNameLabel() bool`

GetIncludeLocationNameLabel returns the IncludeLocationNameLabel field if non-nil, zero value otherwise.

### GetIncludeLocationNameLabelOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetIncludeLocationNameLabelOk() (*bool, bool)`

GetIncludeLocationNameLabelOk returns a tuple with the IncludeLocationNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLocationNameLabel

`func (o *HttpServletExtensionListResponseResourcesInner) SetIncludeLocationNameLabel(v bool)`

SetIncludeLocationNameLabel sets IncludeLocationNameLabel field to given value.

### HasIncludeLocationNameLabel

`func (o *HttpServletExtensionListResponseResourcesInner) HasIncludeLocationNameLabel() bool`

HasIncludeLocationNameLabel returns a boolean if a field has been set.

### GetAlwaysIncludeMonitorEntryNameLabel

`func (o *HttpServletExtensionListResponseResourcesInner) GetAlwaysIncludeMonitorEntryNameLabel() bool`

GetAlwaysIncludeMonitorEntryNameLabel returns the AlwaysIncludeMonitorEntryNameLabel field if non-nil, zero value otherwise.

### GetAlwaysIncludeMonitorEntryNameLabelOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetAlwaysIncludeMonitorEntryNameLabelOk() (*bool, bool)`

GetAlwaysIncludeMonitorEntryNameLabelOk returns a tuple with the AlwaysIncludeMonitorEntryNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysIncludeMonitorEntryNameLabel

`func (o *HttpServletExtensionListResponseResourcesInner) SetAlwaysIncludeMonitorEntryNameLabel(v bool)`

SetAlwaysIncludeMonitorEntryNameLabel sets AlwaysIncludeMonitorEntryNameLabel field to given value.

### HasAlwaysIncludeMonitorEntryNameLabel

`func (o *HttpServletExtensionListResponseResourcesInner) HasAlwaysIncludeMonitorEntryNameLabel() bool`

HasAlwaysIncludeMonitorEntryNameLabel returns a boolean if a field has been set.

### GetIncludeMonitorObjectClassNameLabel

`func (o *HttpServletExtensionListResponseResourcesInner) GetIncludeMonitorObjectClassNameLabel() bool`

GetIncludeMonitorObjectClassNameLabel returns the IncludeMonitorObjectClassNameLabel field if non-nil, zero value otherwise.

### GetIncludeMonitorObjectClassNameLabelOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetIncludeMonitorObjectClassNameLabelOk() (*bool, bool)`

GetIncludeMonitorObjectClassNameLabelOk returns a tuple with the IncludeMonitorObjectClassNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMonitorObjectClassNameLabel

`func (o *HttpServletExtensionListResponseResourcesInner) SetIncludeMonitorObjectClassNameLabel(v bool)`

SetIncludeMonitorObjectClassNameLabel sets IncludeMonitorObjectClassNameLabel field to given value.

### HasIncludeMonitorObjectClassNameLabel

`func (o *HttpServletExtensionListResponseResourcesInner) HasIncludeMonitorObjectClassNameLabel() bool`

HasIncludeMonitorObjectClassNameLabel returns a boolean if a field has been set.

### GetIncludeMonitorAttributeNameLabel

`func (o *HttpServletExtensionListResponseResourcesInner) GetIncludeMonitorAttributeNameLabel() bool`

GetIncludeMonitorAttributeNameLabel returns the IncludeMonitorAttributeNameLabel field if non-nil, zero value otherwise.

### GetIncludeMonitorAttributeNameLabelOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetIncludeMonitorAttributeNameLabelOk() (*bool, bool)`

GetIncludeMonitorAttributeNameLabelOk returns a tuple with the IncludeMonitorAttributeNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMonitorAttributeNameLabel

`func (o *HttpServletExtensionListResponseResourcesInner) SetIncludeMonitorAttributeNameLabel(v bool)`

SetIncludeMonitorAttributeNameLabel sets IncludeMonitorAttributeNameLabel field to given value.

### HasIncludeMonitorAttributeNameLabel

`func (o *HttpServletExtensionListResponseResourcesInner) HasIncludeMonitorAttributeNameLabel() bool`

HasIncludeMonitorAttributeNameLabel returns a boolean if a field has been set.

### GetLabelNameValuePair

`func (o *HttpServletExtensionListResponseResourcesInner) GetLabelNameValuePair() []string`

GetLabelNameValuePair returns the LabelNameValuePair field if non-nil, zero value otherwise.

### GetLabelNameValuePairOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetLabelNameValuePairOk() (*[]string, bool)`

GetLabelNameValuePairOk returns a tuple with the LabelNameValuePair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelNameValuePair

`func (o *HttpServletExtensionListResponseResourcesInner) SetLabelNameValuePair(v []string)`

SetLabelNameValuePair sets LabelNameValuePair field to given value.

### HasLabelNameValuePair

`func (o *HttpServletExtensionListResponseResourcesInner) HasLabelNameValuePair() bool`

HasLabelNameValuePair returns a boolean if a field has been set.

### GetStaticContextPath

`func (o *HttpServletExtensionListResponseResourcesInner) GetStaticContextPath() string`

GetStaticContextPath returns the StaticContextPath field if non-nil, zero value otherwise.

### GetStaticContextPathOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetStaticContextPathOk() (*string, bool)`

GetStaticContextPathOk returns a tuple with the StaticContextPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticContextPath

`func (o *HttpServletExtensionListResponseResourcesInner) SetStaticContextPath(v string)`

SetStaticContextPath sets StaticContextPath field to given value.

### HasStaticContextPath

`func (o *HttpServletExtensionListResponseResourcesInner) HasStaticContextPath() bool`

HasStaticContextPath returns a boolean if a field has been set.

### GetStaticContentDirectory

`func (o *HttpServletExtensionListResponseResourcesInner) GetStaticContentDirectory() string`

GetStaticContentDirectory returns the StaticContentDirectory field if non-nil, zero value otherwise.

### GetStaticContentDirectoryOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetStaticContentDirectoryOk() (*string, bool)`

GetStaticContentDirectoryOk returns a tuple with the StaticContentDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticContentDirectory

`func (o *HttpServletExtensionListResponseResourcesInner) SetStaticContentDirectory(v string)`

SetStaticContentDirectory sets StaticContentDirectory field to given value.

### HasStaticContentDirectory

`func (o *HttpServletExtensionListResponseResourcesInner) HasStaticContentDirectory() bool`

HasStaticContentDirectory returns a boolean if a field has been set.

### GetStaticCustomDirectory

`func (o *HttpServletExtensionListResponseResourcesInner) GetStaticCustomDirectory() string`

GetStaticCustomDirectory returns the StaticCustomDirectory field if non-nil, zero value otherwise.

### GetStaticCustomDirectoryOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetStaticCustomDirectoryOk() (*string, bool)`

GetStaticCustomDirectoryOk returns a tuple with the StaticCustomDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticCustomDirectory

`func (o *HttpServletExtensionListResponseResourcesInner) SetStaticCustomDirectory(v string)`

SetStaticCustomDirectory sets StaticCustomDirectory field to given value.

### HasStaticCustomDirectory

`func (o *HttpServletExtensionListResponseResourcesInner) HasStaticCustomDirectory() bool`

HasStaticCustomDirectory returns a boolean if a field has been set.

### GetTemplateDirectory

`func (o *HttpServletExtensionListResponseResourcesInner) GetTemplateDirectory() []string`

GetTemplateDirectory returns the TemplateDirectory field if non-nil, zero value otherwise.

### GetTemplateDirectoryOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetTemplateDirectoryOk() (*[]string, bool)`

GetTemplateDirectoryOk returns a tuple with the TemplateDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateDirectory

`func (o *HttpServletExtensionListResponseResourcesInner) SetTemplateDirectory(v []string)`

SetTemplateDirectory sets TemplateDirectory field to given value.


### GetExposeRequestAttributes

`func (o *HttpServletExtensionListResponseResourcesInner) GetExposeRequestAttributes() bool`

GetExposeRequestAttributes returns the ExposeRequestAttributes field if non-nil, zero value otherwise.

### GetExposeRequestAttributesOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetExposeRequestAttributesOk() (*bool, bool)`

GetExposeRequestAttributesOk returns a tuple with the ExposeRequestAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposeRequestAttributes

`func (o *HttpServletExtensionListResponseResourcesInner) SetExposeRequestAttributes(v bool)`

SetExposeRequestAttributes sets ExposeRequestAttributes field to given value.

### HasExposeRequestAttributes

`func (o *HttpServletExtensionListResponseResourcesInner) HasExposeRequestAttributes() bool`

HasExposeRequestAttributes returns a boolean if a field has been set.

### GetExposeSessionAttributes

`func (o *HttpServletExtensionListResponseResourcesInner) GetExposeSessionAttributes() bool`

GetExposeSessionAttributes returns the ExposeSessionAttributes field if non-nil, zero value otherwise.

### GetExposeSessionAttributesOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetExposeSessionAttributesOk() (*bool, bool)`

GetExposeSessionAttributesOk returns a tuple with the ExposeSessionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposeSessionAttributes

`func (o *HttpServletExtensionListResponseResourcesInner) SetExposeSessionAttributes(v bool)`

SetExposeSessionAttributes sets ExposeSessionAttributes field to given value.

### HasExposeSessionAttributes

`func (o *HttpServletExtensionListResponseResourcesInner) HasExposeSessionAttributes() bool`

HasExposeSessionAttributes returns a boolean if a field has been set.

### GetExposeServerContext

`func (o *HttpServletExtensionListResponseResourcesInner) GetExposeServerContext() bool`

GetExposeServerContext returns the ExposeServerContext field if non-nil, zero value otherwise.

### GetExposeServerContextOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetExposeServerContextOk() (*bool, bool)`

GetExposeServerContextOk returns a tuple with the ExposeServerContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposeServerContext

`func (o *HttpServletExtensionListResponseResourcesInner) SetExposeServerContext(v bool)`

SetExposeServerContext sets ExposeServerContext field to given value.

### HasExposeServerContext

`func (o *HttpServletExtensionListResponseResourcesInner) HasExposeServerContext() bool`

HasExposeServerContext returns a boolean if a field has been set.

### GetAllowContextOverride

`func (o *HttpServletExtensionListResponseResourcesInner) GetAllowContextOverride() bool`

GetAllowContextOverride returns the AllowContextOverride field if non-nil, zero value otherwise.

### GetAllowContextOverrideOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetAllowContextOverrideOk() (*bool, bool)`

GetAllowContextOverrideOk returns a tuple with the AllowContextOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowContextOverride

`func (o *HttpServletExtensionListResponseResourcesInner) SetAllowContextOverride(v bool)`

SetAllowContextOverride sets AllowContextOverride field to given value.

### HasAllowContextOverride

`func (o *HttpServletExtensionListResponseResourcesInner) HasAllowContextOverride() bool`

HasAllowContextOverride returns a boolean if a field has been set.

### GetMimeTypesFile

`func (o *HttpServletExtensionListResponseResourcesInner) GetMimeTypesFile() string`

GetMimeTypesFile returns the MimeTypesFile field if non-nil, zero value otherwise.

### GetMimeTypesFileOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetMimeTypesFileOk() (*string, bool)`

GetMimeTypesFileOk returns a tuple with the MimeTypesFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeTypesFile

`func (o *HttpServletExtensionListResponseResourcesInner) SetMimeTypesFile(v string)`

SetMimeTypesFile sets MimeTypesFile field to given value.

### HasMimeTypesFile

`func (o *HttpServletExtensionListResponseResourcesInner) HasMimeTypesFile() bool`

HasMimeTypesFile returns a boolean if a field has been set.

### GetDefaultMIMEType

`func (o *HttpServletExtensionListResponseResourcesInner) GetDefaultMIMEType() string`

GetDefaultMIMEType returns the DefaultMIMEType field if non-nil, zero value otherwise.

### GetDefaultMIMETypeOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetDefaultMIMETypeOk() (*string, bool)`

GetDefaultMIMETypeOk returns a tuple with the DefaultMIMEType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMIMEType

`func (o *HttpServletExtensionListResponseResourcesInner) SetDefaultMIMEType(v string)`

SetDefaultMIMEType sets DefaultMIMEType field to given value.

### HasDefaultMIMEType

`func (o *HttpServletExtensionListResponseResourcesInner) HasDefaultMIMEType() bool`

HasDefaultMIMEType returns a boolean if a field has been set.

### GetCharacterEncoding

`func (o *HttpServletExtensionListResponseResourcesInner) GetCharacterEncoding() string`

GetCharacterEncoding returns the CharacterEncoding field if non-nil, zero value otherwise.

### GetCharacterEncodingOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetCharacterEncodingOk() (*string, bool)`

GetCharacterEncodingOk returns a tuple with the CharacterEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterEncoding

`func (o *HttpServletExtensionListResponseResourcesInner) SetCharacterEncoding(v string)`

SetCharacterEncoding sets CharacterEncoding field to given value.

### HasCharacterEncoding

`func (o *HttpServletExtensionListResponseResourcesInner) HasCharacterEncoding() bool`

HasCharacterEncoding returns a boolean if a field has been set.

### GetStaticResponseHeader

`func (o *HttpServletExtensionListResponseResourcesInner) GetStaticResponseHeader() []string`

GetStaticResponseHeader returns the StaticResponseHeader field if non-nil, zero value otherwise.

### GetStaticResponseHeaderOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetStaticResponseHeaderOk() (*[]string, bool)`

GetStaticResponseHeaderOk returns a tuple with the StaticResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticResponseHeader

`func (o *HttpServletExtensionListResponseResourcesInner) SetStaticResponseHeader(v []string)`

SetStaticResponseHeader sets StaticResponseHeader field to given value.

### HasStaticResponseHeader

`func (o *HttpServletExtensionListResponseResourcesInner) HasStaticResponseHeader() bool`

HasStaticResponseHeader returns a boolean if a field has been set.

### GetRequireAuthentication

`func (o *HttpServletExtensionListResponseResourcesInner) GetRequireAuthentication() bool`

GetRequireAuthentication returns the RequireAuthentication field if non-nil, zero value otherwise.

### GetRequireAuthenticationOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetRequireAuthenticationOk() (*bool, bool)`

GetRequireAuthenticationOk returns a tuple with the RequireAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAuthentication

`func (o *HttpServletExtensionListResponseResourcesInner) SetRequireAuthentication(v bool)`

SetRequireAuthentication sets RequireAuthentication field to given value.

### HasRequireAuthentication

`func (o *HttpServletExtensionListResponseResourcesInner) HasRequireAuthentication() bool`

HasRequireAuthentication returns a boolean if a field has been set.

### GetBearerTokenAuthEnabled

`func (o *HttpServletExtensionListResponseResourcesInner) GetBearerTokenAuthEnabled() bool`

GetBearerTokenAuthEnabled returns the BearerTokenAuthEnabled field if non-nil, zero value otherwise.

### GetBearerTokenAuthEnabledOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetBearerTokenAuthEnabledOk() (*bool, bool)`

GetBearerTokenAuthEnabledOk returns a tuple with the BearerTokenAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerTokenAuthEnabled

`func (o *HttpServletExtensionListResponseResourcesInner) SetBearerTokenAuthEnabled(v bool)`

SetBearerTokenAuthEnabled sets BearerTokenAuthEnabled field to given value.

### HasBearerTokenAuthEnabled

`func (o *HttpServletExtensionListResponseResourcesInner) HasBearerTokenAuthEnabled() bool`

HasBearerTokenAuthEnabled returns a boolean if a field has been set.

### GetOAuthTokenHandler

`func (o *HttpServletExtensionListResponseResourcesInner) GetOAuthTokenHandler() string`

GetOAuthTokenHandler returns the OAuthTokenHandler field if non-nil, zero value otherwise.

### GetOAuthTokenHandlerOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetOAuthTokenHandlerOk() (*string, bool)`

GetOAuthTokenHandlerOk returns a tuple with the OAuthTokenHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthTokenHandler

`func (o *HttpServletExtensionListResponseResourcesInner) SetOAuthTokenHandler(v string)`

SetOAuthTokenHandler sets OAuthTokenHandler field to given value.

### HasOAuthTokenHandler

`func (o *HttpServletExtensionListResponseResourcesInner) HasOAuthTokenHandler() bool`

HasOAuthTokenHandler returns a boolean if a field has been set.

### GetResourceMappingFile

`func (o *HttpServletExtensionListResponseResourcesInner) GetResourceMappingFile() string`

GetResourceMappingFile returns the ResourceMappingFile field if non-nil, zero value otherwise.

### GetResourceMappingFileOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetResourceMappingFileOk() (*string, bool)`

GetResourceMappingFileOk returns a tuple with the ResourceMappingFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceMappingFile

`func (o *HttpServletExtensionListResponseResourcesInner) SetResourceMappingFile(v string)`

SetResourceMappingFile sets ResourceMappingFile field to given value.

### HasResourceMappingFile

`func (o *HttpServletExtensionListResponseResourcesInner) HasResourceMappingFile() bool`

HasResourceMappingFile returns a boolean if a field has been set.

### GetIncludeLDAPObjectclass

`func (o *HttpServletExtensionListResponseResourcesInner) GetIncludeLDAPObjectclass() []string`

GetIncludeLDAPObjectclass returns the IncludeLDAPObjectclass field if non-nil, zero value otherwise.

### GetIncludeLDAPObjectclassOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetIncludeLDAPObjectclassOk() (*[]string, bool)`

GetIncludeLDAPObjectclassOk returns a tuple with the IncludeLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLDAPObjectclass

`func (o *HttpServletExtensionListResponseResourcesInner) SetIncludeLDAPObjectclass(v []string)`

SetIncludeLDAPObjectclass sets IncludeLDAPObjectclass field to given value.

### HasIncludeLDAPObjectclass

`func (o *HttpServletExtensionListResponseResourcesInner) HasIncludeLDAPObjectclass() bool`

HasIncludeLDAPObjectclass returns a boolean if a field has been set.

### GetExcludeLDAPObjectclass

`func (o *HttpServletExtensionListResponseResourcesInner) GetExcludeLDAPObjectclass() []string`

GetExcludeLDAPObjectclass returns the ExcludeLDAPObjectclass field if non-nil, zero value otherwise.

### GetExcludeLDAPObjectclassOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetExcludeLDAPObjectclassOk() (*[]string, bool)`

GetExcludeLDAPObjectclassOk returns a tuple with the ExcludeLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeLDAPObjectclass

`func (o *HttpServletExtensionListResponseResourcesInner) SetExcludeLDAPObjectclass(v []string)`

SetExcludeLDAPObjectclass sets ExcludeLDAPObjectclass field to given value.

### HasExcludeLDAPObjectclass

`func (o *HttpServletExtensionListResponseResourcesInner) HasExcludeLDAPObjectclass() bool`

HasExcludeLDAPObjectclass returns a boolean if a field has been set.

### GetIncludeLDAPBaseDN

`func (o *HttpServletExtensionListResponseResourcesInner) GetIncludeLDAPBaseDN() []string`

GetIncludeLDAPBaseDN returns the IncludeLDAPBaseDN field if non-nil, zero value otherwise.

### GetIncludeLDAPBaseDNOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetIncludeLDAPBaseDNOk() (*[]string, bool)`

GetIncludeLDAPBaseDNOk returns a tuple with the IncludeLDAPBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLDAPBaseDN

`func (o *HttpServletExtensionListResponseResourcesInner) SetIncludeLDAPBaseDN(v []string)`

SetIncludeLDAPBaseDN sets IncludeLDAPBaseDN field to given value.

### HasIncludeLDAPBaseDN

`func (o *HttpServletExtensionListResponseResourcesInner) HasIncludeLDAPBaseDN() bool`

HasIncludeLDAPBaseDN returns a boolean if a field has been set.

### GetExcludeLDAPBaseDN

`func (o *HttpServletExtensionListResponseResourcesInner) GetExcludeLDAPBaseDN() []string`

GetExcludeLDAPBaseDN returns the ExcludeLDAPBaseDN field if non-nil, zero value otherwise.

### GetExcludeLDAPBaseDNOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetExcludeLDAPBaseDNOk() (*[]string, bool)`

GetExcludeLDAPBaseDNOk returns a tuple with the ExcludeLDAPBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeLDAPBaseDN

`func (o *HttpServletExtensionListResponseResourcesInner) SetExcludeLDAPBaseDN(v []string)`

SetExcludeLDAPBaseDN sets ExcludeLDAPBaseDN field to given value.

### HasExcludeLDAPBaseDN

`func (o *HttpServletExtensionListResponseResourcesInner) HasExcludeLDAPBaseDN() bool`

HasExcludeLDAPBaseDN returns a boolean if a field has been set.

### GetEntityTagLDAPAttribute

`func (o *HttpServletExtensionListResponseResourcesInner) GetEntityTagLDAPAttribute() string`

GetEntityTagLDAPAttribute returns the EntityTagLDAPAttribute field if non-nil, zero value otherwise.

### GetEntityTagLDAPAttributeOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetEntityTagLDAPAttributeOk() (*string, bool)`

GetEntityTagLDAPAttributeOk returns a tuple with the EntityTagLDAPAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTagLDAPAttribute

`func (o *HttpServletExtensionListResponseResourcesInner) SetEntityTagLDAPAttribute(v string)`

SetEntityTagLDAPAttribute sets EntityTagLDAPAttribute field to given value.

### HasEntityTagLDAPAttribute

`func (o *HttpServletExtensionListResponseResourcesInner) HasEntityTagLDAPAttribute() bool`

HasEntityTagLDAPAttribute returns a boolean if a field has been set.

### GetTemporaryDirectory

`func (o *HttpServletExtensionListResponseResourcesInner) GetTemporaryDirectory() string`

GetTemporaryDirectory returns the TemporaryDirectory field if non-nil, zero value otherwise.

### GetTemporaryDirectoryOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetTemporaryDirectoryOk() (*string, bool)`

GetTemporaryDirectoryOk returns a tuple with the TemporaryDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryDirectory

`func (o *HttpServletExtensionListResponseResourcesInner) SetTemporaryDirectory(v string)`

SetTemporaryDirectory sets TemporaryDirectory field to given value.


### GetTemporaryDirectoryPermissions

`func (o *HttpServletExtensionListResponseResourcesInner) GetTemporaryDirectoryPermissions() string`

GetTemporaryDirectoryPermissions returns the TemporaryDirectoryPermissions field if non-nil, zero value otherwise.

### GetTemporaryDirectoryPermissionsOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetTemporaryDirectoryPermissionsOk() (*string, bool)`

GetTemporaryDirectoryPermissionsOk returns a tuple with the TemporaryDirectoryPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryDirectoryPermissions

`func (o *HttpServletExtensionListResponseResourcesInner) SetTemporaryDirectoryPermissions(v string)`

SetTemporaryDirectoryPermissions sets TemporaryDirectoryPermissions field to given value.


### GetMaxResults

`func (o *HttpServletExtensionListResponseResourcesInner) GetMaxResults() int64`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetMaxResultsOk() (*int64, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *HttpServletExtensionListResponseResourcesInner) SetMaxResults(v int64)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *HttpServletExtensionListResponseResourcesInner) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetBulkMaxOperations

`func (o *HttpServletExtensionListResponseResourcesInner) GetBulkMaxOperations() int64`

GetBulkMaxOperations returns the BulkMaxOperations field if non-nil, zero value otherwise.

### GetBulkMaxOperationsOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetBulkMaxOperationsOk() (*int64, bool)`

GetBulkMaxOperationsOk returns a tuple with the BulkMaxOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkMaxOperations

`func (o *HttpServletExtensionListResponseResourcesInner) SetBulkMaxOperations(v int64)`

SetBulkMaxOperations sets BulkMaxOperations field to given value.

### HasBulkMaxOperations

`func (o *HttpServletExtensionListResponseResourcesInner) HasBulkMaxOperations() bool`

HasBulkMaxOperations returns a boolean if a field has been set.

### GetBulkMaxPayloadSize

`func (o *HttpServletExtensionListResponseResourcesInner) GetBulkMaxPayloadSize() string`

GetBulkMaxPayloadSize returns the BulkMaxPayloadSize field if non-nil, zero value otherwise.

### GetBulkMaxPayloadSizeOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetBulkMaxPayloadSizeOk() (*string, bool)`

GetBulkMaxPayloadSizeOk returns a tuple with the BulkMaxPayloadSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkMaxPayloadSize

`func (o *HttpServletExtensionListResponseResourcesInner) SetBulkMaxPayloadSize(v string)`

SetBulkMaxPayloadSize sets BulkMaxPayloadSize field to given value.

### HasBulkMaxPayloadSize

`func (o *HttpServletExtensionListResponseResourcesInner) HasBulkMaxPayloadSize() bool`

HasBulkMaxPayloadSize returns a boolean if a field has been set.

### GetBulkMaxConcurrentRequests

`func (o *HttpServletExtensionListResponseResourcesInner) GetBulkMaxConcurrentRequests() int64`

GetBulkMaxConcurrentRequests returns the BulkMaxConcurrentRequests field if non-nil, zero value otherwise.

### GetBulkMaxConcurrentRequestsOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetBulkMaxConcurrentRequestsOk() (*int64, bool)`

GetBulkMaxConcurrentRequestsOk returns a tuple with the BulkMaxConcurrentRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkMaxConcurrentRequests

`func (o *HttpServletExtensionListResponseResourcesInner) SetBulkMaxConcurrentRequests(v int64)`

SetBulkMaxConcurrentRequests sets BulkMaxConcurrentRequests field to given value.

### HasBulkMaxConcurrentRequests

`func (o *HttpServletExtensionListResponseResourcesInner) HasBulkMaxConcurrentRequests() bool`

HasBulkMaxConcurrentRequests returns a boolean if a field has been set.

### GetDebugEnabled

`func (o *HttpServletExtensionListResponseResourcesInner) GetDebugEnabled() bool`

GetDebugEnabled returns the DebugEnabled field if non-nil, zero value otherwise.

### GetDebugEnabledOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetDebugEnabledOk() (*bool, bool)`

GetDebugEnabledOk returns a tuple with the DebugEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugEnabled

`func (o *HttpServletExtensionListResponseResourcesInner) SetDebugEnabled(v bool)`

SetDebugEnabled sets DebugEnabled field to given value.

### HasDebugEnabled

`func (o *HttpServletExtensionListResponseResourcesInner) HasDebugEnabled() bool`

HasDebugEnabled returns a boolean if a field has been set.

### GetDebugLevel

`func (o *HttpServletExtensionListResponseResourcesInner) GetDebugLevel() EnumhttpServletExtensionDebugLevelProp`

GetDebugLevel returns the DebugLevel field if non-nil, zero value otherwise.

### GetDebugLevelOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetDebugLevelOk() (*EnumhttpServletExtensionDebugLevelProp, bool)`

GetDebugLevelOk returns a tuple with the DebugLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugLevel

`func (o *HttpServletExtensionListResponseResourcesInner) SetDebugLevel(v EnumhttpServletExtensionDebugLevelProp)`

SetDebugLevel sets DebugLevel field to given value.


### GetDebugType

`func (o *HttpServletExtensionListResponseResourcesInner) GetDebugType() []EnumhttpServletExtensionDebugTypeProp`

GetDebugType returns the DebugType field if non-nil, zero value otherwise.

### GetDebugTypeOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetDebugTypeOk() (*[]EnumhttpServletExtensionDebugTypeProp, bool)`

GetDebugTypeOk returns a tuple with the DebugType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugType

`func (o *HttpServletExtensionListResponseResourcesInner) SetDebugType(v []EnumhttpServletExtensionDebugTypeProp)`

SetDebugType sets DebugType field to given value.


### GetIncludeStackTrace

`func (o *HttpServletExtensionListResponseResourcesInner) GetIncludeStackTrace() bool`

GetIncludeStackTrace returns the IncludeStackTrace field if non-nil, zero value otherwise.

### GetIncludeStackTraceOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetIncludeStackTraceOk() (*bool, bool)`

GetIncludeStackTraceOk returns a tuple with the IncludeStackTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeStackTrace

`func (o *HttpServletExtensionListResponseResourcesInner) SetIncludeStackTrace(v bool)`

SetIncludeStackTrace sets IncludeStackTrace field to given value.


### GetScriptClass

`func (o *HttpServletExtensionListResponseResourcesInner) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *HttpServletExtensionListResponseResourcesInner) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *HttpServletExtensionListResponseResourcesInner) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *HttpServletExtensionListResponseResourcesInner) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *HttpServletExtensionListResponseResourcesInner) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetFapiFinancialID

`func (o *HttpServletExtensionListResponseResourcesInner) GetFapiFinancialID() string`

GetFapiFinancialID returns the FapiFinancialID field if non-nil, zero value otherwise.

### GetFapiFinancialIDOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetFapiFinancialIDOk() (*string, bool)`

GetFapiFinancialIDOk returns a tuple with the FapiFinancialID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFapiFinancialID

`func (o *HttpServletExtensionListResponseResourcesInner) SetFapiFinancialID(v string)`

SetFapiFinancialID sets FapiFinancialID field to given value.

### HasFapiFinancialID

`func (o *HttpServletExtensionListResponseResourcesInner) HasFapiFinancialID() bool`

HasFapiFinancialID returns a boolean if a field has been set.

### GetPathPrefix

`func (o *HttpServletExtensionListResponseResourcesInner) GetPathPrefix() string`

GetPathPrefix returns the PathPrefix field if non-nil, zero value otherwise.

### GetPathPrefixOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetPathPrefixOk() (*string, bool)`

GetPathPrefixOk returns a tuple with the PathPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPrefix

`func (o *HttpServletExtensionListResponseResourcesInner) SetPathPrefix(v string)`

SetPathPrefix sets PathPrefix field to given value.

### HasPathPrefix

`func (o *HttpServletExtensionListResponseResourcesInner) HasPathPrefix() bool`

HasPathPrefix returns a boolean if a field has been set.

### GetConsentServer

`func (o *HttpServletExtensionListResponseResourcesInner) GetConsentServer() string`

GetConsentServer returns the ConsentServer field if non-nil, zero value otherwise.

### GetConsentServerOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetConsentServerOk() (*string, bool)`

GetConsentServerOk returns a tuple with the ConsentServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentServer

`func (o *HttpServletExtensionListResponseResourcesInner) SetConsentServer(v string)`

SetConsentServer sets ConsentServer field to given value.

### HasConsentServer

`func (o *HttpServletExtensionListResponseResourcesInner) HasConsentServer() bool`

HasConsentServer returns a boolean if a field has been set.

### GetConsentDefinitionID

`func (o *HttpServletExtensionListResponseResourcesInner) GetConsentDefinitionID() string`

GetConsentDefinitionID returns the ConsentDefinitionID field if non-nil, zero value otherwise.

### GetConsentDefinitionIDOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetConsentDefinitionIDOk() (*string, bool)`

GetConsentDefinitionIDOk returns a tuple with the ConsentDefinitionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentDefinitionID

`func (o *HttpServletExtensionListResponseResourcesInner) SetConsentDefinitionID(v string)`

SetConsentDefinitionID sets ConsentDefinitionID field to given value.


### GetDocumentRootDirectory

`func (o *HttpServletExtensionListResponseResourcesInner) GetDocumentRootDirectory() string`

GetDocumentRootDirectory returns the DocumentRootDirectory field if non-nil, zero value otherwise.

### GetDocumentRootDirectoryOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetDocumentRootDirectoryOk() (*string, bool)`

GetDocumentRootDirectoryOk returns a tuple with the DocumentRootDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentRootDirectory

`func (o *HttpServletExtensionListResponseResourcesInner) SetDocumentRootDirectory(v string)`

SetDocumentRootDirectory sets DocumentRootDirectory field to given value.


### GetEnableDirectoryIndexing

`func (o *HttpServletExtensionListResponseResourcesInner) GetEnableDirectoryIndexing() bool`

GetEnableDirectoryIndexing returns the EnableDirectoryIndexing field if non-nil, zero value otherwise.

### GetEnableDirectoryIndexingOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetEnableDirectoryIndexingOk() (*bool, bool)`

GetEnableDirectoryIndexingOk returns a tuple with the EnableDirectoryIndexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDirectoryIndexing

`func (o *HttpServletExtensionListResponseResourcesInner) SetEnableDirectoryIndexing(v bool)`

SetEnableDirectoryIndexing sets EnableDirectoryIndexing field to given value.

### HasEnableDirectoryIndexing

`func (o *HttpServletExtensionListResponseResourcesInner) HasEnableDirectoryIndexing() bool`

HasEnableDirectoryIndexing returns a boolean if a field has been set.

### GetIndexFile

`func (o *HttpServletExtensionListResponseResourcesInner) GetIndexFile() []string`

GetIndexFile returns the IndexFile field if non-nil, zero value otherwise.

### GetIndexFileOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetIndexFileOk() (*[]string, bool)`

GetIndexFileOk returns a tuple with the IndexFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexFile

`func (o *HttpServletExtensionListResponseResourcesInner) SetIndexFile(v []string)`

SetIndexFile sets IndexFile field to given value.

### HasIndexFile

`func (o *HttpServletExtensionListResponseResourcesInner) HasIndexFile() bool`

HasIndexFile returns a boolean if a field has been set.

### GetAllowedAuthenticationType

`func (o *HttpServletExtensionListResponseResourcesInner) GetAllowedAuthenticationType() []EnumhttpServletExtensionAllowedAuthenticationTypeProp`

GetAllowedAuthenticationType returns the AllowedAuthenticationType field if non-nil, zero value otherwise.

### GetAllowedAuthenticationTypeOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetAllowedAuthenticationTypeOk() (*[]EnumhttpServletExtensionAllowedAuthenticationTypeProp, bool)`

GetAllowedAuthenticationTypeOk returns a tuple with the AllowedAuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAuthenticationType

`func (o *HttpServletExtensionListResponseResourcesInner) SetAllowedAuthenticationType(v []EnumhttpServletExtensionAllowedAuthenticationTypeProp)`

SetAllowedAuthenticationType sets AllowedAuthenticationType field to given value.

### HasAllowedAuthenticationType

`func (o *HttpServletExtensionListResponseResourcesInner) HasAllowedAuthenticationType() bool`

HasAllowedAuthenticationType returns a boolean if a field has been set.

### GetIdTokenValidator

`func (o *HttpServletExtensionListResponseResourcesInner) GetIdTokenValidator() []string`

GetIdTokenValidator returns the IdTokenValidator field if non-nil, zero value otherwise.

### GetIdTokenValidatorOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetIdTokenValidatorOk() (*[]string, bool)`

GetIdTokenValidatorOk returns a tuple with the IdTokenValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenValidator

`func (o *HttpServletExtensionListResponseResourcesInner) SetIdTokenValidator(v []string)`

SetIdTokenValidator sets IdTokenValidator field to given value.

### HasIdTokenValidator

`func (o *HttpServletExtensionListResponseResourcesInner) HasIdTokenValidator() bool`

HasIdTokenValidator returns a boolean if a field has been set.

### GetRequireFileServletAccessPrivilege

`func (o *HttpServletExtensionListResponseResourcesInner) GetRequireFileServletAccessPrivilege() bool`

GetRequireFileServletAccessPrivilege returns the RequireFileServletAccessPrivilege field if non-nil, zero value otherwise.

### GetRequireFileServletAccessPrivilegeOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetRequireFileServletAccessPrivilegeOk() (*bool, bool)`

GetRequireFileServletAccessPrivilegeOk returns a tuple with the RequireFileServletAccessPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireFileServletAccessPrivilege

`func (o *HttpServletExtensionListResponseResourcesInner) SetRequireFileServletAccessPrivilege(v bool)`

SetRequireFileServletAccessPrivilege sets RequireFileServletAccessPrivilege field to given value.

### HasRequireFileServletAccessPrivilege

`func (o *HttpServletExtensionListResponseResourcesInner) HasRequireFileServletAccessPrivilege() bool`

HasRequireFileServletAccessPrivilege returns a boolean if a field has been set.

### GetRequireGroup

`func (o *HttpServletExtensionListResponseResourcesInner) GetRequireGroup() []string`

GetRequireGroup returns the RequireGroup field if non-nil, zero value otherwise.

### GetRequireGroupOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetRequireGroupOk() (*[]string, bool)`

GetRequireGroupOk returns a tuple with the RequireGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireGroup

`func (o *HttpServletExtensionListResponseResourcesInner) SetRequireGroup(v []string)`

SetRequireGroup sets RequireGroup field to given value.

### HasRequireGroup

`func (o *HttpServletExtensionListResponseResourcesInner) HasRequireGroup() bool`

HasRequireGroup returns a boolean if a field has been set.

### GetSharedSecretHeaderName

`func (o *HttpServletExtensionListResponseResourcesInner) GetSharedSecretHeaderName() string`

GetSharedSecretHeaderName returns the SharedSecretHeaderName field if non-nil, zero value otherwise.

### GetSharedSecretHeaderNameOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetSharedSecretHeaderNameOk() (*string, bool)`

GetSharedSecretHeaderNameOk returns a tuple with the SharedSecretHeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecretHeaderName

`func (o *HttpServletExtensionListResponseResourcesInner) SetSharedSecretHeaderName(v string)`

SetSharedSecretHeaderName sets SharedSecretHeaderName field to given value.


### GetSharedSecrets

`func (o *HttpServletExtensionListResponseResourcesInner) GetSharedSecrets() []string`

GetSharedSecrets returns the SharedSecrets field if non-nil, zero value otherwise.

### GetSharedSecretsOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetSharedSecretsOk() (*[]string, bool)`

GetSharedSecretsOk returns a tuple with the SharedSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecrets

`func (o *HttpServletExtensionListResponseResourcesInner) SetSharedSecrets(v []string)`

SetSharedSecrets sets SharedSecrets field to given value.

### HasSharedSecrets

`func (o *HttpServletExtensionListResponseResourcesInner) HasSharedSecrets() bool`

HasSharedSecrets returns a boolean if a field has been set.

### GetRequireAPIAuthentication

`func (o *HttpServletExtensionListResponseResourcesInner) GetRequireAPIAuthentication() bool`

GetRequireAPIAuthentication returns the RequireAPIAuthentication field if non-nil, zero value otherwise.

### GetRequireAPIAuthenticationOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetRequireAPIAuthenticationOk() (*bool, bool)`

GetRequireAPIAuthenticationOk returns a tuple with the RequireAPIAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAPIAuthentication

`func (o *HttpServletExtensionListResponseResourcesInner) SetRequireAPIAuthentication(v bool)`

SetRequireAPIAuthentication sets RequireAPIAuthentication field to given value.

### HasRequireAPIAuthentication

`func (o *HttpServletExtensionListResponseResourcesInner) HasRequireAPIAuthentication() bool`

HasRequireAPIAuthentication returns a boolean if a field has been set.

### GetOmitErrorMessageDetails

`func (o *HttpServletExtensionListResponseResourcesInner) GetOmitErrorMessageDetails() bool`

GetOmitErrorMessageDetails returns the OmitErrorMessageDetails field if non-nil, zero value otherwise.

### GetOmitErrorMessageDetailsOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetOmitErrorMessageDetailsOk() (*bool, bool)`

GetOmitErrorMessageDetailsOk returns a tuple with the OmitErrorMessageDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmitErrorMessageDetails

`func (o *HttpServletExtensionListResponseResourcesInner) SetOmitErrorMessageDetails(v bool)`

SetOmitErrorMessageDetails sets OmitErrorMessageDetails field to given value.

### HasOmitErrorMessageDetails

`func (o *HttpServletExtensionListResponseResourcesInner) HasOmitErrorMessageDetails() bool`

HasOmitErrorMessageDetails returns a boolean if a field has been set.

### GetApiAuthenticationTimeout

`func (o *HttpServletExtensionListResponseResourcesInner) GetApiAuthenticationTimeout() string`

GetApiAuthenticationTimeout returns the ApiAuthenticationTimeout field if non-nil, zero value otherwise.

### GetApiAuthenticationTimeoutOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetApiAuthenticationTimeoutOk() (*string, bool)`

GetApiAuthenticationTimeoutOk returns a tuple with the ApiAuthenticationTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiAuthenticationTimeout

`func (o *HttpServletExtensionListResponseResourcesInner) SetApiAuthenticationTimeout(v string)`

SetApiAuthenticationTimeout sets ApiAuthenticationTimeout field to given value.

### HasApiAuthenticationTimeout

`func (o *HttpServletExtensionListResponseResourcesInner) HasApiAuthenticationTimeout() bool`

HasApiAuthenticationTimeout returns a boolean if a field has been set.

### GetMapAccessTokensToLocalUsers

`func (o *HttpServletExtensionListResponseResourcesInner) GetMapAccessTokensToLocalUsers() EnumhttpServletExtensionMapAccessTokensToLocalUsersProp`

GetMapAccessTokensToLocalUsers returns the MapAccessTokensToLocalUsers field if non-nil, zero value otherwise.

### GetMapAccessTokensToLocalUsersOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetMapAccessTokensToLocalUsersOk() (*EnumhttpServletExtensionMapAccessTokensToLocalUsersProp, bool)`

GetMapAccessTokensToLocalUsersOk returns a tuple with the MapAccessTokensToLocalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapAccessTokensToLocalUsers

`func (o *HttpServletExtensionListResponseResourcesInner) SetMapAccessTokensToLocalUsers(v EnumhttpServletExtensionMapAccessTokensToLocalUsersProp)`

SetMapAccessTokensToLocalUsers sets MapAccessTokensToLocalUsers field to given value.

### HasMapAccessTokensToLocalUsers

`func (o *HttpServletExtensionListResponseResourcesInner) HasMapAccessTokensToLocalUsers() bool`

HasMapAccessTokensToLocalUsers returns a boolean if a field has been set.

### GetSwaggerEnabled

`func (o *HttpServletExtensionListResponseResourcesInner) GetSwaggerEnabled() bool`

GetSwaggerEnabled returns the SwaggerEnabled field if non-nil, zero value otherwise.

### GetSwaggerEnabledOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetSwaggerEnabledOk() (*bool, bool)`

GetSwaggerEnabledOk returns a tuple with the SwaggerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwaggerEnabled

`func (o *HttpServletExtensionListResponseResourcesInner) SetSwaggerEnabled(v bool)`

SetSwaggerEnabled sets SwaggerEnabled field to given value.

### HasSwaggerEnabled

`func (o *HttpServletExtensionListResponseResourcesInner) HasSwaggerEnabled() bool`

HasSwaggerEnabled returns a boolean if a field has been set.

### GetMaxPageSize

`func (o *HttpServletExtensionListResponseResourcesInner) GetMaxPageSize() int64`

GetMaxPageSize returns the MaxPageSize field if non-nil, zero value otherwise.

### GetMaxPageSizeOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetMaxPageSizeOk() (*int64, bool)`

GetMaxPageSizeOk returns a tuple with the MaxPageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPageSize

`func (o *HttpServletExtensionListResponseResourcesInner) SetMaxPageSize(v int64)`

SetMaxPageSize sets MaxPageSize field to given value.

### HasMaxPageSize

`func (o *HttpServletExtensionListResponseResourcesInner) HasMaxPageSize() bool`

HasMaxPageSize returns a boolean if a field has been set.

### GetSchemasEndpointObjectclass

`func (o *HttpServletExtensionListResponseResourcesInner) GetSchemasEndpointObjectclass() []string`

GetSchemasEndpointObjectclass returns the SchemasEndpointObjectclass field if non-nil, zero value otherwise.

### GetSchemasEndpointObjectclassOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetSchemasEndpointObjectclassOk() (*[]string, bool)`

GetSchemasEndpointObjectclassOk returns a tuple with the SchemasEndpointObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemasEndpointObjectclass

`func (o *HttpServletExtensionListResponseResourcesInner) SetSchemasEndpointObjectclass(v []string)`

SetSchemasEndpointObjectclass sets SchemasEndpointObjectclass field to given value.

### HasSchemasEndpointObjectclass

`func (o *HttpServletExtensionListResponseResourcesInner) HasSchemasEndpointObjectclass() bool`

HasSchemasEndpointObjectclass returns a boolean if a field has been set.

### GetDefaultOperationalAttribute

`func (o *HttpServletExtensionListResponseResourcesInner) GetDefaultOperationalAttribute() []string`

GetDefaultOperationalAttribute returns the DefaultOperationalAttribute field if non-nil, zero value otherwise.

### GetDefaultOperationalAttributeOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetDefaultOperationalAttributeOk() (*[]string, bool)`

GetDefaultOperationalAttributeOk returns a tuple with the DefaultOperationalAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOperationalAttribute

`func (o *HttpServletExtensionListResponseResourcesInner) SetDefaultOperationalAttribute(v []string)`

SetDefaultOperationalAttribute sets DefaultOperationalAttribute field to given value.

### HasDefaultOperationalAttribute

`func (o *HttpServletExtensionListResponseResourcesInner) HasDefaultOperationalAttribute() bool`

HasDefaultOperationalAttribute returns a boolean if a field has been set.

### GetRejectExpansionAttribute

`func (o *HttpServletExtensionListResponseResourcesInner) GetRejectExpansionAttribute() []string`

GetRejectExpansionAttribute returns the RejectExpansionAttribute field if non-nil, zero value otherwise.

### GetRejectExpansionAttributeOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetRejectExpansionAttributeOk() (*[]string, bool)`

GetRejectExpansionAttributeOk returns a tuple with the RejectExpansionAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectExpansionAttribute

`func (o *HttpServletExtensionListResponseResourcesInner) SetRejectExpansionAttribute(v []string)`

SetRejectExpansionAttribute sets RejectExpansionAttribute field to given value.

### HasRejectExpansionAttribute

`func (o *HttpServletExtensionListResponseResourcesInner) HasRejectExpansionAttribute() bool`

HasRejectExpansionAttribute returns a boolean if a field has been set.

### GetAlwaysUsePermissiveModify

`func (o *HttpServletExtensionListResponseResourcesInner) GetAlwaysUsePermissiveModify() bool`

GetAlwaysUsePermissiveModify returns the AlwaysUsePermissiveModify field if non-nil, zero value otherwise.

### GetAlwaysUsePermissiveModifyOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetAlwaysUsePermissiveModifyOk() (*bool, bool)`

GetAlwaysUsePermissiveModifyOk returns a tuple with the AlwaysUsePermissiveModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysUsePermissiveModify

`func (o *HttpServletExtensionListResponseResourcesInner) SetAlwaysUsePermissiveModify(v bool)`

SetAlwaysUsePermissiveModify sets AlwaysUsePermissiveModify field to given value.

### HasAlwaysUsePermissiveModify

`func (o *HttpServletExtensionListResponseResourcesInner) HasAlwaysUsePermissiveModify() bool`

HasAlwaysUsePermissiveModify returns a boolean if a field has been set.

### GetAllowedControl

`func (o *HttpServletExtensionListResponseResourcesInner) GetAllowedControl() []EnumhttpServletExtensionAllowedControlProp`

GetAllowedControl returns the AllowedControl field if non-nil, zero value otherwise.

### GetAllowedControlOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetAllowedControlOk() (*[]EnumhttpServletExtensionAllowedControlProp, bool)`

GetAllowedControlOk returns a tuple with the AllowedControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedControl

`func (o *HttpServletExtensionListResponseResourcesInner) SetAllowedControl(v []EnumhttpServletExtensionAllowedControlProp)`

SetAllowedControl sets AllowedControl field to given value.

### HasAllowedControl

`func (o *HttpServletExtensionListResponseResourcesInner) HasAllowedControl() bool`

HasAllowedControl returns a boolean if a field has been set.

### GetExcludedAPIEndpoint

`func (o *HttpServletExtensionListResponseResourcesInner) GetExcludedAPIEndpoint() []string`

GetExcludedAPIEndpoint returns the ExcludedAPIEndpoint field if non-nil, zero value otherwise.

### GetExcludedAPIEndpointOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetExcludedAPIEndpointOk() (*[]string, bool)`

GetExcludedAPIEndpointOk returns a tuple with the ExcludedAPIEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedAPIEndpoint

`func (o *HttpServletExtensionListResponseResourcesInner) SetExcludedAPIEndpoint(v []string)`

SetExcludedAPIEndpoint sets ExcludedAPIEndpoint field to given value.

### HasExcludedAPIEndpoint

`func (o *HttpServletExtensionListResponseResourcesInner) HasExcludedAPIEndpoint() bool`

HasExcludedAPIEndpoint returns a boolean if a field has been set.

### GetRequestLimit

`func (o *HttpServletExtensionListResponseResourcesInner) GetRequestLimit() string`

GetRequestLimit returns the RequestLimit field if non-nil, zero value otherwise.

### GetRequestLimitOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetRequestLimitOk() (*string, bool)`

GetRequestLimitOk returns a tuple with the RequestLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestLimit

`func (o *HttpServletExtensionListResponseResourcesInner) SetRequestLimit(v string)`

SetRequestLimit sets RequestLimit field to given value.

### HasRequestLimit

`func (o *HttpServletExtensionListResponseResourcesInner) HasRequestLimit() bool`

HasRequestLimit returns a boolean if a field has been set.

### GetResponseLimit

`func (o *HttpServletExtensionListResponseResourcesInner) GetResponseLimit() string`

GetResponseLimit returns the ResponseLimit field if non-nil, zero value otherwise.

### GetResponseLimitOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetResponseLimitOk() (*string, bool)`

GetResponseLimitOk returns a tuple with the ResponseLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseLimit

`func (o *HttpServletExtensionListResponseResourcesInner) SetResponseLimit(v string)`

SetResponseLimit sets ResponseLimit field to given value.

### HasResponseLimit

`func (o *HttpServletExtensionListResponseResourcesInner) HasResponseLimit() bool`

HasResponseLimit returns a boolean if a field has been set.

### GetNumForwardThreads

`func (o *HttpServletExtensionListResponseResourcesInner) GetNumForwardThreads() int64`

GetNumForwardThreads returns the NumForwardThreads field if non-nil, zero value otherwise.

### GetNumForwardThreadsOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetNumForwardThreadsOk() (*int64, bool)`

GetNumForwardThreadsOk returns a tuple with the NumForwardThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumForwardThreads

`func (o *HttpServletExtensionListResponseResourcesInner) SetNumForwardThreads(v int64)`

SetNumForwardThreads sets NumForwardThreads field to given value.

### HasNumForwardThreads

`func (o *HttpServletExtensionListResponseResourcesInner) HasNumForwardThreads() bool`

HasNumForwardThreads returns a boolean if a field has been set.

### GetExtensionClass

`func (o *HttpServletExtensionListResponseResourcesInner) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *HttpServletExtensionListResponseResourcesInner) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *HttpServletExtensionListResponseResourcesInner) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *HttpServletExtensionListResponseResourcesInner) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *HttpServletExtensionListResponseResourcesInner) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetRequestContextMethod

`func (o *HttpServletExtensionListResponseResourcesInner) GetRequestContextMethod() EnumhttpServletExtensionRequestContextMethodProp`

GetRequestContextMethod returns the RequestContextMethod field if non-nil, zero value otherwise.

### GetRequestContextMethodOk

`func (o *HttpServletExtensionListResponseResourcesInner) GetRequestContextMethodOk() (*EnumhttpServletExtensionRequestContextMethodProp, bool)`

GetRequestContextMethodOk returns a tuple with the RequestContextMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestContextMethod

`func (o *HttpServletExtensionListResponseResourcesInner) SetRequestContextMethod(v EnumhttpServletExtensionRequestContextMethodProp)`

SetRequestContextMethod sets RequestContextMethod field to given value.

### HasRequestContextMethod

`func (o *HttpServletExtensionListResponseResourcesInner) HasRequestContextMethod() bool`

HasRequestContextMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


