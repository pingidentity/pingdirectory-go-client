# GetHttpServletExtension200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumthirdPartyHttpServletExtensionSchemaUrn**](EnumthirdPartyHttpServletExtensionSchemaUrn.md) |  | 
**Id** | **string** | Name of the HTTP Servlet Extension | 
**BasicAuthEnabled** | Pointer to **bool** | Enables HTTP Basic authentication, using a username and password. The Identity Mapper specified by the identity-mapper property will be used to map the username to a DN. | [optional] 
**IdentityMapper** | Pointer to **string** | Specifies the Identity Mapper that is to be used for associating user entries with basic authentication usernames. | [optional] 
**AccessTokenValidator** | Pointer to **[]string** | If specified, the Access Token Validator(s) that may be used to validate access tokens for requests submitted to this Directory REST API HTTP Servlet Extension. | [optional] 
**AccessTokenScope** | Pointer to **string** | The name of a scope that must be present in an access token accepted by the Directory REST API HTTP Servlet Extension. | [optional] 
**Audience** | Pointer to **string** | A string or URI that identifies the Directory REST API HTTP Servlet Extension in the context of OAuth2 authorization. | [optional] 
**Description** | Pointer to **string** | A description for this HTTP Servlet Extension | [optional] 
**CrossOriginPolicy** | Pointer to **string** | The cross-origin request policy to use for the HTTP Servlet Extension. | [optional] 
**ResponseHeader** | Pointer to **[]string** | Specifies HTTP header fields and values added to response headers for all requests. | [optional] 
**CorrelationIDResponseHeader** | Pointer to **string** | Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \&quot;Correlation-Id\&quot;, \&quot;X-Amzn-Trace-Id\&quot;, and \&quot;X-Request-Id\&quot;. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
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
**DocumentRootDirectory** | **string** | Specifies the path to the directory on the local filesystem containing the files to be served by this File Server HTTP Servlet Extension. The path must exist, and it must be a directory. | 
**EnableDirectoryIndexing** | Pointer to **bool** | Indicates whether to generate a default HTML page with a listing of available files if the requested path refers to a directory rather than a file, and that directory does not contain an index file. | [optional] 
**IndexFile** | Pointer to **[]string** | Specifies the name of a file whose contents may be returned to the client if the requested path refers to a directory rather than a file. | [optional] 
**AllowedAuthenticationType** | Pointer to [**[]EnumhttpServletExtensionAllowedAuthenticationTypeProp**](EnumhttpServletExtensionAllowedAuthenticationTypeProp.md) |  | [optional] 
**IdTokenValidator** | Pointer to **[]string** | The ID token validators that may be used to verify the authenticity of an of an OpenID Connect ID token. | [optional] 
**RequireFileServletAccessPrivilege** | Pointer to **bool** | Indicates whether the servlet extension should only accept requests from authenticated clients that have the file-servlet-access privilege. | [optional] 
**RequireGroup** | Pointer to **[]string** | The DN of a group whose members will be permitted to access to the associated files. If multiple group DNs are configured, then anyone who is a member of at least one of those groups will be granted access. | [optional] 
**MapAccessTokensToLocalUsers** | Pointer to [**EnumhttpServletExtensionMapAccessTokensToLocalUsersProp**](EnumhttpServletExtensionMapAccessTokensToLocalUsersProp.md) |  | [optional] 
**SwaggerEnabled** | Pointer to **bool** | Indicates whether the SCIM2 HTTP Servlet Extension will generate a Swagger specification document. | [optional] 
**MaxPageSize** | Pointer to **int64** | The maximum number of entries to be returned in one page of search results. | [optional] 
**SchemasEndpointObjectclass** | Pointer to **[]string** | The list of object classes which will be returned by the schemas endpoint. | [optional] 
**DefaultOperationalAttribute** | Pointer to **[]string** | A set of operational attributes that will be returned with entries by default. | [optional] 
**RejectExpansionAttribute** | Pointer to **[]string** | A set of attributes which the client is not allowed to provide for the expand query parameters. This should be used for attributes that could either have a large number of values or that reference entries that are very large like groups. | [optional] 
**AlwaysUsePermissiveModify** | Pointer to **bool** | Indicates whether to always use permissive modify behavior for PATCH requests, even if the request did not include the permissive modify request control. | [optional] 
**AllowedControl** | Pointer to [**[]EnumhttpServletExtensionAllowedControlProp**](EnumhttpServletExtensionAllowedControlProp.md) |  | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party HTTP Servlet Extension. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party HTTP Servlet Extension. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewGetHttpServletExtension200Response

`func NewGetHttpServletExtension200Response(schemas []EnumthirdPartyHttpServletExtensionSchemaUrn, id string, baseContextPath string, availableStatusCode int64, degradedStatusCode int64, unavailableStatusCode int64, templateDirectory []string, temporaryDirectory string, temporaryDirectoryPermissions string, debugLevel EnumhttpServletExtensionDebugLevelProp, debugType []EnumhttpServletExtensionDebugTypeProp, includeStackTrace bool, scriptClass string, documentRootDirectory string, extensionClass string, ) *GetHttpServletExtension200Response`

NewGetHttpServletExtension200Response instantiates a new GetHttpServletExtension200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetHttpServletExtension200ResponseWithDefaults

`func NewGetHttpServletExtension200ResponseWithDefaults() *GetHttpServletExtension200Response`

NewGetHttpServletExtension200ResponseWithDefaults instantiates a new GetHttpServletExtension200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *GetHttpServletExtension200Response) GetSchemas() []EnumthirdPartyHttpServletExtensionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GetHttpServletExtension200Response) GetSchemasOk() (*[]EnumthirdPartyHttpServletExtensionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GetHttpServletExtension200Response) SetSchemas(v []EnumthirdPartyHttpServletExtensionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *GetHttpServletExtension200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetHttpServletExtension200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetHttpServletExtension200Response) SetId(v string)`

SetId sets Id field to given value.


### GetBasicAuthEnabled

`func (o *GetHttpServletExtension200Response) GetBasicAuthEnabled() bool`

GetBasicAuthEnabled returns the BasicAuthEnabled field if non-nil, zero value otherwise.

### GetBasicAuthEnabledOk

`func (o *GetHttpServletExtension200Response) GetBasicAuthEnabledOk() (*bool, bool)`

GetBasicAuthEnabledOk returns a tuple with the BasicAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthEnabled

`func (o *GetHttpServletExtension200Response) SetBasicAuthEnabled(v bool)`

SetBasicAuthEnabled sets BasicAuthEnabled field to given value.

### HasBasicAuthEnabled

`func (o *GetHttpServletExtension200Response) HasBasicAuthEnabled() bool`

HasBasicAuthEnabled returns a boolean if a field has been set.

### GetIdentityMapper

`func (o *GetHttpServletExtension200Response) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *GetHttpServletExtension200Response) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *GetHttpServletExtension200Response) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.

### HasIdentityMapper

`func (o *GetHttpServletExtension200Response) HasIdentityMapper() bool`

HasIdentityMapper returns a boolean if a field has been set.

### GetAccessTokenValidator

`func (o *GetHttpServletExtension200Response) GetAccessTokenValidator() []string`

GetAccessTokenValidator returns the AccessTokenValidator field if non-nil, zero value otherwise.

### GetAccessTokenValidatorOk

`func (o *GetHttpServletExtension200Response) GetAccessTokenValidatorOk() (*[]string, bool)`

GetAccessTokenValidatorOk returns a tuple with the AccessTokenValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValidator

`func (o *GetHttpServletExtension200Response) SetAccessTokenValidator(v []string)`

SetAccessTokenValidator sets AccessTokenValidator field to given value.

### HasAccessTokenValidator

`func (o *GetHttpServletExtension200Response) HasAccessTokenValidator() bool`

HasAccessTokenValidator returns a boolean if a field has been set.

### GetAccessTokenScope

`func (o *GetHttpServletExtension200Response) GetAccessTokenScope() string`

GetAccessTokenScope returns the AccessTokenScope field if non-nil, zero value otherwise.

### GetAccessTokenScopeOk

`func (o *GetHttpServletExtension200Response) GetAccessTokenScopeOk() (*string, bool)`

GetAccessTokenScopeOk returns a tuple with the AccessTokenScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenScope

`func (o *GetHttpServletExtension200Response) SetAccessTokenScope(v string)`

SetAccessTokenScope sets AccessTokenScope field to given value.

### HasAccessTokenScope

`func (o *GetHttpServletExtension200Response) HasAccessTokenScope() bool`

HasAccessTokenScope returns a boolean if a field has been set.

### GetAudience

`func (o *GetHttpServletExtension200Response) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *GetHttpServletExtension200Response) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *GetHttpServletExtension200Response) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *GetHttpServletExtension200Response) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetDescription

`func (o *GetHttpServletExtension200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetHttpServletExtension200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetHttpServletExtension200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetHttpServletExtension200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCrossOriginPolicy

`func (o *GetHttpServletExtension200Response) GetCrossOriginPolicy() string`

GetCrossOriginPolicy returns the CrossOriginPolicy field if non-nil, zero value otherwise.

### GetCrossOriginPolicyOk

`func (o *GetHttpServletExtension200Response) GetCrossOriginPolicyOk() (*string, bool)`

GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossOriginPolicy

`func (o *GetHttpServletExtension200Response) SetCrossOriginPolicy(v string)`

SetCrossOriginPolicy sets CrossOriginPolicy field to given value.

### HasCrossOriginPolicy

`func (o *GetHttpServletExtension200Response) HasCrossOriginPolicy() bool`

HasCrossOriginPolicy returns a boolean if a field has been set.

### GetResponseHeader

`func (o *GetHttpServletExtension200Response) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *GetHttpServletExtension200Response) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *GetHttpServletExtension200Response) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *GetHttpServletExtension200Response) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetCorrelationIDResponseHeader

`func (o *GetHttpServletExtension200Response) GetCorrelationIDResponseHeader() string`

GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field if non-nil, zero value otherwise.

### GetCorrelationIDResponseHeaderOk

`func (o *GetHttpServletExtension200Response) GetCorrelationIDResponseHeaderOk() (*string, bool)`

GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDResponseHeader

`func (o *GetHttpServletExtension200Response) SetCorrelationIDResponseHeader(v string)`

SetCorrelationIDResponseHeader sets CorrelationIDResponseHeader field to given value.

### HasCorrelationIDResponseHeader

`func (o *GetHttpServletExtension200Response) HasCorrelationIDResponseHeader() bool`

HasCorrelationIDResponseHeader returns a boolean if a field has been set.

### GetMeta

`func (o *GetHttpServletExtension200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetHttpServletExtension200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetHttpServletExtension200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetHttpServletExtension200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *GetHttpServletExtension200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *GetHttpServletExtension200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *GetHttpServletExtension200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *GetHttpServletExtension200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetServer

`func (o *GetHttpServletExtension200Response) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *GetHttpServletExtension200Response) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *GetHttpServletExtension200Response) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *GetHttpServletExtension200Response) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetBaseContextPath

`func (o *GetHttpServletExtension200Response) GetBaseContextPath() string`

GetBaseContextPath returns the BaseContextPath field if non-nil, zero value otherwise.

### GetBaseContextPathOk

`func (o *GetHttpServletExtension200Response) GetBaseContextPathOk() (*string, bool)`

GetBaseContextPathOk returns a tuple with the BaseContextPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseContextPath

`func (o *GetHttpServletExtension200Response) SetBaseContextPath(v string)`

SetBaseContextPath sets BaseContextPath field to given value.


### GetAvailableStatusCode

`func (o *GetHttpServletExtension200Response) GetAvailableStatusCode() int64`

GetAvailableStatusCode returns the AvailableStatusCode field if non-nil, zero value otherwise.

### GetAvailableStatusCodeOk

`func (o *GetHttpServletExtension200Response) GetAvailableStatusCodeOk() (*int64, bool)`

GetAvailableStatusCodeOk returns a tuple with the AvailableStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableStatusCode

`func (o *GetHttpServletExtension200Response) SetAvailableStatusCode(v int64)`

SetAvailableStatusCode sets AvailableStatusCode field to given value.


### GetDegradedStatusCode

`func (o *GetHttpServletExtension200Response) GetDegradedStatusCode() int64`

GetDegradedStatusCode returns the DegradedStatusCode field if non-nil, zero value otherwise.

### GetDegradedStatusCodeOk

`func (o *GetHttpServletExtension200Response) GetDegradedStatusCodeOk() (*int64, bool)`

GetDegradedStatusCodeOk returns a tuple with the DegradedStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegradedStatusCode

`func (o *GetHttpServletExtension200Response) SetDegradedStatusCode(v int64)`

SetDegradedStatusCode sets DegradedStatusCode field to given value.


### GetUnavailableStatusCode

`func (o *GetHttpServletExtension200Response) GetUnavailableStatusCode() int64`

GetUnavailableStatusCode returns the UnavailableStatusCode field if non-nil, zero value otherwise.

### GetUnavailableStatusCodeOk

`func (o *GetHttpServletExtension200Response) GetUnavailableStatusCodeOk() (*int64, bool)`

GetUnavailableStatusCodeOk returns a tuple with the UnavailableStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailableStatusCode

`func (o *GetHttpServletExtension200Response) SetUnavailableStatusCode(v int64)`

SetUnavailableStatusCode sets UnavailableStatusCode field to given value.


### GetOverrideStatusCode

`func (o *GetHttpServletExtension200Response) GetOverrideStatusCode() int64`

GetOverrideStatusCode returns the OverrideStatusCode field if non-nil, zero value otherwise.

### GetOverrideStatusCodeOk

`func (o *GetHttpServletExtension200Response) GetOverrideStatusCodeOk() (*int64, bool)`

GetOverrideStatusCodeOk returns a tuple with the OverrideStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideStatusCode

`func (o *GetHttpServletExtension200Response) SetOverrideStatusCode(v int64)`

SetOverrideStatusCode sets OverrideStatusCode field to given value.

### HasOverrideStatusCode

`func (o *GetHttpServletExtension200Response) HasOverrideStatusCode() bool`

HasOverrideStatusCode returns a boolean if a field has been set.

### GetIncludeResponseBody

`func (o *GetHttpServletExtension200Response) GetIncludeResponseBody() bool`

GetIncludeResponseBody returns the IncludeResponseBody field if non-nil, zero value otherwise.

### GetIncludeResponseBodyOk

`func (o *GetHttpServletExtension200Response) GetIncludeResponseBodyOk() (*bool, bool)`

GetIncludeResponseBodyOk returns a tuple with the IncludeResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeResponseBody

`func (o *GetHttpServletExtension200Response) SetIncludeResponseBody(v bool)`

SetIncludeResponseBody sets IncludeResponseBody field to given value.

### HasIncludeResponseBody

`func (o *GetHttpServletExtension200Response) HasIncludeResponseBody() bool`

HasIncludeResponseBody returns a boolean if a field has been set.

### GetAdditionalResponseContents

`func (o *GetHttpServletExtension200Response) GetAdditionalResponseContents() string`

GetAdditionalResponseContents returns the AdditionalResponseContents field if non-nil, zero value otherwise.

### GetAdditionalResponseContentsOk

`func (o *GetHttpServletExtension200Response) GetAdditionalResponseContentsOk() (*string, bool)`

GetAdditionalResponseContentsOk returns a tuple with the AdditionalResponseContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalResponseContents

`func (o *GetHttpServletExtension200Response) SetAdditionalResponseContents(v string)`

SetAdditionalResponseContents sets AdditionalResponseContents field to given value.

### HasAdditionalResponseContents

`func (o *GetHttpServletExtension200Response) HasAdditionalResponseContents() bool`

HasAdditionalResponseContents returns a boolean if a field has been set.

### GetIncludeInstanceNameLabel

`func (o *GetHttpServletExtension200Response) GetIncludeInstanceNameLabel() bool`

GetIncludeInstanceNameLabel returns the IncludeInstanceNameLabel field if non-nil, zero value otherwise.

### GetIncludeInstanceNameLabelOk

`func (o *GetHttpServletExtension200Response) GetIncludeInstanceNameLabelOk() (*bool, bool)`

GetIncludeInstanceNameLabelOk returns a tuple with the IncludeInstanceNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInstanceNameLabel

`func (o *GetHttpServletExtension200Response) SetIncludeInstanceNameLabel(v bool)`

SetIncludeInstanceNameLabel sets IncludeInstanceNameLabel field to given value.

### HasIncludeInstanceNameLabel

`func (o *GetHttpServletExtension200Response) HasIncludeInstanceNameLabel() bool`

HasIncludeInstanceNameLabel returns a boolean if a field has been set.

### GetIncludeProductNameLabel

`func (o *GetHttpServletExtension200Response) GetIncludeProductNameLabel() bool`

GetIncludeProductNameLabel returns the IncludeProductNameLabel field if non-nil, zero value otherwise.

### GetIncludeProductNameLabelOk

`func (o *GetHttpServletExtension200Response) GetIncludeProductNameLabelOk() (*bool, bool)`

GetIncludeProductNameLabelOk returns a tuple with the IncludeProductNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeProductNameLabel

`func (o *GetHttpServletExtension200Response) SetIncludeProductNameLabel(v bool)`

SetIncludeProductNameLabel sets IncludeProductNameLabel field to given value.

### HasIncludeProductNameLabel

`func (o *GetHttpServletExtension200Response) HasIncludeProductNameLabel() bool`

HasIncludeProductNameLabel returns a boolean if a field has been set.

### GetIncludeLocationNameLabel

`func (o *GetHttpServletExtension200Response) GetIncludeLocationNameLabel() bool`

GetIncludeLocationNameLabel returns the IncludeLocationNameLabel field if non-nil, zero value otherwise.

### GetIncludeLocationNameLabelOk

`func (o *GetHttpServletExtension200Response) GetIncludeLocationNameLabelOk() (*bool, bool)`

GetIncludeLocationNameLabelOk returns a tuple with the IncludeLocationNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLocationNameLabel

`func (o *GetHttpServletExtension200Response) SetIncludeLocationNameLabel(v bool)`

SetIncludeLocationNameLabel sets IncludeLocationNameLabel field to given value.

### HasIncludeLocationNameLabel

`func (o *GetHttpServletExtension200Response) HasIncludeLocationNameLabel() bool`

HasIncludeLocationNameLabel returns a boolean if a field has been set.

### GetAlwaysIncludeMonitorEntryNameLabel

`func (o *GetHttpServletExtension200Response) GetAlwaysIncludeMonitorEntryNameLabel() bool`

GetAlwaysIncludeMonitorEntryNameLabel returns the AlwaysIncludeMonitorEntryNameLabel field if non-nil, zero value otherwise.

### GetAlwaysIncludeMonitorEntryNameLabelOk

`func (o *GetHttpServletExtension200Response) GetAlwaysIncludeMonitorEntryNameLabelOk() (*bool, bool)`

GetAlwaysIncludeMonitorEntryNameLabelOk returns a tuple with the AlwaysIncludeMonitorEntryNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysIncludeMonitorEntryNameLabel

`func (o *GetHttpServletExtension200Response) SetAlwaysIncludeMonitorEntryNameLabel(v bool)`

SetAlwaysIncludeMonitorEntryNameLabel sets AlwaysIncludeMonitorEntryNameLabel field to given value.

### HasAlwaysIncludeMonitorEntryNameLabel

`func (o *GetHttpServletExtension200Response) HasAlwaysIncludeMonitorEntryNameLabel() bool`

HasAlwaysIncludeMonitorEntryNameLabel returns a boolean if a field has been set.

### GetIncludeMonitorObjectClassNameLabel

`func (o *GetHttpServletExtension200Response) GetIncludeMonitorObjectClassNameLabel() bool`

GetIncludeMonitorObjectClassNameLabel returns the IncludeMonitorObjectClassNameLabel field if non-nil, zero value otherwise.

### GetIncludeMonitorObjectClassNameLabelOk

`func (o *GetHttpServletExtension200Response) GetIncludeMonitorObjectClassNameLabelOk() (*bool, bool)`

GetIncludeMonitorObjectClassNameLabelOk returns a tuple with the IncludeMonitorObjectClassNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMonitorObjectClassNameLabel

`func (o *GetHttpServletExtension200Response) SetIncludeMonitorObjectClassNameLabel(v bool)`

SetIncludeMonitorObjectClassNameLabel sets IncludeMonitorObjectClassNameLabel field to given value.

### HasIncludeMonitorObjectClassNameLabel

`func (o *GetHttpServletExtension200Response) HasIncludeMonitorObjectClassNameLabel() bool`

HasIncludeMonitorObjectClassNameLabel returns a boolean if a field has been set.

### GetIncludeMonitorAttributeNameLabel

`func (o *GetHttpServletExtension200Response) GetIncludeMonitorAttributeNameLabel() bool`

GetIncludeMonitorAttributeNameLabel returns the IncludeMonitorAttributeNameLabel field if non-nil, zero value otherwise.

### GetIncludeMonitorAttributeNameLabelOk

`func (o *GetHttpServletExtension200Response) GetIncludeMonitorAttributeNameLabelOk() (*bool, bool)`

GetIncludeMonitorAttributeNameLabelOk returns a tuple with the IncludeMonitorAttributeNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMonitorAttributeNameLabel

`func (o *GetHttpServletExtension200Response) SetIncludeMonitorAttributeNameLabel(v bool)`

SetIncludeMonitorAttributeNameLabel sets IncludeMonitorAttributeNameLabel field to given value.

### HasIncludeMonitorAttributeNameLabel

`func (o *GetHttpServletExtension200Response) HasIncludeMonitorAttributeNameLabel() bool`

HasIncludeMonitorAttributeNameLabel returns a boolean if a field has been set.

### GetLabelNameValuePair

`func (o *GetHttpServletExtension200Response) GetLabelNameValuePair() []string`

GetLabelNameValuePair returns the LabelNameValuePair field if non-nil, zero value otherwise.

### GetLabelNameValuePairOk

`func (o *GetHttpServletExtension200Response) GetLabelNameValuePairOk() (*[]string, bool)`

GetLabelNameValuePairOk returns a tuple with the LabelNameValuePair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelNameValuePair

`func (o *GetHttpServletExtension200Response) SetLabelNameValuePair(v []string)`

SetLabelNameValuePair sets LabelNameValuePair field to given value.

### HasLabelNameValuePair

`func (o *GetHttpServletExtension200Response) HasLabelNameValuePair() bool`

HasLabelNameValuePair returns a boolean if a field has been set.

### GetStaticContextPath

`func (o *GetHttpServletExtension200Response) GetStaticContextPath() string`

GetStaticContextPath returns the StaticContextPath field if non-nil, zero value otherwise.

### GetStaticContextPathOk

`func (o *GetHttpServletExtension200Response) GetStaticContextPathOk() (*string, bool)`

GetStaticContextPathOk returns a tuple with the StaticContextPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticContextPath

`func (o *GetHttpServletExtension200Response) SetStaticContextPath(v string)`

SetStaticContextPath sets StaticContextPath field to given value.

### HasStaticContextPath

`func (o *GetHttpServletExtension200Response) HasStaticContextPath() bool`

HasStaticContextPath returns a boolean if a field has been set.

### GetStaticContentDirectory

`func (o *GetHttpServletExtension200Response) GetStaticContentDirectory() string`

GetStaticContentDirectory returns the StaticContentDirectory field if non-nil, zero value otherwise.

### GetStaticContentDirectoryOk

`func (o *GetHttpServletExtension200Response) GetStaticContentDirectoryOk() (*string, bool)`

GetStaticContentDirectoryOk returns a tuple with the StaticContentDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticContentDirectory

`func (o *GetHttpServletExtension200Response) SetStaticContentDirectory(v string)`

SetStaticContentDirectory sets StaticContentDirectory field to given value.

### HasStaticContentDirectory

`func (o *GetHttpServletExtension200Response) HasStaticContentDirectory() bool`

HasStaticContentDirectory returns a boolean if a field has been set.

### GetStaticCustomDirectory

`func (o *GetHttpServletExtension200Response) GetStaticCustomDirectory() string`

GetStaticCustomDirectory returns the StaticCustomDirectory field if non-nil, zero value otherwise.

### GetStaticCustomDirectoryOk

`func (o *GetHttpServletExtension200Response) GetStaticCustomDirectoryOk() (*string, bool)`

GetStaticCustomDirectoryOk returns a tuple with the StaticCustomDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticCustomDirectory

`func (o *GetHttpServletExtension200Response) SetStaticCustomDirectory(v string)`

SetStaticCustomDirectory sets StaticCustomDirectory field to given value.

### HasStaticCustomDirectory

`func (o *GetHttpServletExtension200Response) HasStaticCustomDirectory() bool`

HasStaticCustomDirectory returns a boolean if a field has been set.

### GetTemplateDirectory

`func (o *GetHttpServletExtension200Response) GetTemplateDirectory() []string`

GetTemplateDirectory returns the TemplateDirectory field if non-nil, zero value otherwise.

### GetTemplateDirectoryOk

`func (o *GetHttpServletExtension200Response) GetTemplateDirectoryOk() (*[]string, bool)`

GetTemplateDirectoryOk returns a tuple with the TemplateDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateDirectory

`func (o *GetHttpServletExtension200Response) SetTemplateDirectory(v []string)`

SetTemplateDirectory sets TemplateDirectory field to given value.


### GetExposeRequestAttributes

`func (o *GetHttpServletExtension200Response) GetExposeRequestAttributes() bool`

GetExposeRequestAttributes returns the ExposeRequestAttributes field if non-nil, zero value otherwise.

### GetExposeRequestAttributesOk

`func (o *GetHttpServletExtension200Response) GetExposeRequestAttributesOk() (*bool, bool)`

GetExposeRequestAttributesOk returns a tuple with the ExposeRequestAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposeRequestAttributes

`func (o *GetHttpServletExtension200Response) SetExposeRequestAttributes(v bool)`

SetExposeRequestAttributes sets ExposeRequestAttributes field to given value.

### HasExposeRequestAttributes

`func (o *GetHttpServletExtension200Response) HasExposeRequestAttributes() bool`

HasExposeRequestAttributes returns a boolean if a field has been set.

### GetExposeSessionAttributes

`func (o *GetHttpServletExtension200Response) GetExposeSessionAttributes() bool`

GetExposeSessionAttributes returns the ExposeSessionAttributes field if non-nil, zero value otherwise.

### GetExposeSessionAttributesOk

`func (o *GetHttpServletExtension200Response) GetExposeSessionAttributesOk() (*bool, bool)`

GetExposeSessionAttributesOk returns a tuple with the ExposeSessionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposeSessionAttributes

`func (o *GetHttpServletExtension200Response) SetExposeSessionAttributes(v bool)`

SetExposeSessionAttributes sets ExposeSessionAttributes field to given value.

### HasExposeSessionAttributes

`func (o *GetHttpServletExtension200Response) HasExposeSessionAttributes() bool`

HasExposeSessionAttributes returns a boolean if a field has been set.

### GetExposeServerContext

`func (o *GetHttpServletExtension200Response) GetExposeServerContext() bool`

GetExposeServerContext returns the ExposeServerContext field if non-nil, zero value otherwise.

### GetExposeServerContextOk

`func (o *GetHttpServletExtension200Response) GetExposeServerContextOk() (*bool, bool)`

GetExposeServerContextOk returns a tuple with the ExposeServerContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposeServerContext

`func (o *GetHttpServletExtension200Response) SetExposeServerContext(v bool)`

SetExposeServerContext sets ExposeServerContext field to given value.

### HasExposeServerContext

`func (o *GetHttpServletExtension200Response) HasExposeServerContext() bool`

HasExposeServerContext returns a boolean if a field has been set.

### GetAllowContextOverride

`func (o *GetHttpServletExtension200Response) GetAllowContextOverride() bool`

GetAllowContextOverride returns the AllowContextOverride field if non-nil, zero value otherwise.

### GetAllowContextOverrideOk

`func (o *GetHttpServletExtension200Response) GetAllowContextOverrideOk() (*bool, bool)`

GetAllowContextOverrideOk returns a tuple with the AllowContextOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowContextOverride

`func (o *GetHttpServletExtension200Response) SetAllowContextOverride(v bool)`

SetAllowContextOverride sets AllowContextOverride field to given value.

### HasAllowContextOverride

`func (o *GetHttpServletExtension200Response) HasAllowContextOverride() bool`

HasAllowContextOverride returns a boolean if a field has been set.

### GetMimeTypesFile

`func (o *GetHttpServletExtension200Response) GetMimeTypesFile() string`

GetMimeTypesFile returns the MimeTypesFile field if non-nil, zero value otherwise.

### GetMimeTypesFileOk

`func (o *GetHttpServletExtension200Response) GetMimeTypesFileOk() (*string, bool)`

GetMimeTypesFileOk returns a tuple with the MimeTypesFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeTypesFile

`func (o *GetHttpServletExtension200Response) SetMimeTypesFile(v string)`

SetMimeTypesFile sets MimeTypesFile field to given value.

### HasMimeTypesFile

`func (o *GetHttpServletExtension200Response) HasMimeTypesFile() bool`

HasMimeTypesFile returns a boolean if a field has been set.

### GetDefaultMIMEType

`func (o *GetHttpServletExtension200Response) GetDefaultMIMEType() string`

GetDefaultMIMEType returns the DefaultMIMEType field if non-nil, zero value otherwise.

### GetDefaultMIMETypeOk

`func (o *GetHttpServletExtension200Response) GetDefaultMIMETypeOk() (*string, bool)`

GetDefaultMIMETypeOk returns a tuple with the DefaultMIMEType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMIMEType

`func (o *GetHttpServletExtension200Response) SetDefaultMIMEType(v string)`

SetDefaultMIMEType sets DefaultMIMEType field to given value.

### HasDefaultMIMEType

`func (o *GetHttpServletExtension200Response) HasDefaultMIMEType() bool`

HasDefaultMIMEType returns a boolean if a field has been set.

### GetCharacterEncoding

`func (o *GetHttpServletExtension200Response) GetCharacterEncoding() string`

GetCharacterEncoding returns the CharacterEncoding field if non-nil, zero value otherwise.

### GetCharacterEncodingOk

`func (o *GetHttpServletExtension200Response) GetCharacterEncodingOk() (*string, bool)`

GetCharacterEncodingOk returns a tuple with the CharacterEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterEncoding

`func (o *GetHttpServletExtension200Response) SetCharacterEncoding(v string)`

SetCharacterEncoding sets CharacterEncoding field to given value.

### HasCharacterEncoding

`func (o *GetHttpServletExtension200Response) HasCharacterEncoding() bool`

HasCharacterEncoding returns a boolean if a field has been set.

### GetStaticResponseHeader

`func (o *GetHttpServletExtension200Response) GetStaticResponseHeader() []string`

GetStaticResponseHeader returns the StaticResponseHeader field if non-nil, zero value otherwise.

### GetStaticResponseHeaderOk

`func (o *GetHttpServletExtension200Response) GetStaticResponseHeaderOk() (*[]string, bool)`

GetStaticResponseHeaderOk returns a tuple with the StaticResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticResponseHeader

`func (o *GetHttpServletExtension200Response) SetStaticResponseHeader(v []string)`

SetStaticResponseHeader sets StaticResponseHeader field to given value.

### HasStaticResponseHeader

`func (o *GetHttpServletExtension200Response) HasStaticResponseHeader() bool`

HasStaticResponseHeader returns a boolean if a field has been set.

### GetRequireAuthentication

`func (o *GetHttpServletExtension200Response) GetRequireAuthentication() bool`

GetRequireAuthentication returns the RequireAuthentication field if non-nil, zero value otherwise.

### GetRequireAuthenticationOk

`func (o *GetHttpServletExtension200Response) GetRequireAuthenticationOk() (*bool, bool)`

GetRequireAuthenticationOk returns a tuple with the RequireAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAuthentication

`func (o *GetHttpServletExtension200Response) SetRequireAuthentication(v bool)`

SetRequireAuthentication sets RequireAuthentication field to given value.

### HasRequireAuthentication

`func (o *GetHttpServletExtension200Response) HasRequireAuthentication() bool`

HasRequireAuthentication returns a boolean if a field has been set.

### GetBearerTokenAuthEnabled

`func (o *GetHttpServletExtension200Response) GetBearerTokenAuthEnabled() bool`

GetBearerTokenAuthEnabled returns the BearerTokenAuthEnabled field if non-nil, zero value otherwise.

### GetBearerTokenAuthEnabledOk

`func (o *GetHttpServletExtension200Response) GetBearerTokenAuthEnabledOk() (*bool, bool)`

GetBearerTokenAuthEnabledOk returns a tuple with the BearerTokenAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerTokenAuthEnabled

`func (o *GetHttpServletExtension200Response) SetBearerTokenAuthEnabled(v bool)`

SetBearerTokenAuthEnabled sets BearerTokenAuthEnabled field to given value.

### HasBearerTokenAuthEnabled

`func (o *GetHttpServletExtension200Response) HasBearerTokenAuthEnabled() bool`

HasBearerTokenAuthEnabled returns a boolean if a field has been set.

### GetOAuthTokenHandler

`func (o *GetHttpServletExtension200Response) GetOAuthTokenHandler() string`

GetOAuthTokenHandler returns the OAuthTokenHandler field if non-nil, zero value otherwise.

### GetOAuthTokenHandlerOk

`func (o *GetHttpServletExtension200Response) GetOAuthTokenHandlerOk() (*string, bool)`

GetOAuthTokenHandlerOk returns a tuple with the OAuthTokenHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthTokenHandler

`func (o *GetHttpServletExtension200Response) SetOAuthTokenHandler(v string)`

SetOAuthTokenHandler sets OAuthTokenHandler field to given value.

### HasOAuthTokenHandler

`func (o *GetHttpServletExtension200Response) HasOAuthTokenHandler() bool`

HasOAuthTokenHandler returns a boolean if a field has been set.

### GetResourceMappingFile

`func (o *GetHttpServletExtension200Response) GetResourceMappingFile() string`

GetResourceMappingFile returns the ResourceMappingFile field if non-nil, zero value otherwise.

### GetResourceMappingFileOk

`func (o *GetHttpServletExtension200Response) GetResourceMappingFileOk() (*string, bool)`

GetResourceMappingFileOk returns a tuple with the ResourceMappingFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceMappingFile

`func (o *GetHttpServletExtension200Response) SetResourceMappingFile(v string)`

SetResourceMappingFile sets ResourceMappingFile field to given value.

### HasResourceMappingFile

`func (o *GetHttpServletExtension200Response) HasResourceMappingFile() bool`

HasResourceMappingFile returns a boolean if a field has been set.

### GetIncludeLDAPObjectclass

`func (o *GetHttpServletExtension200Response) GetIncludeLDAPObjectclass() []string`

GetIncludeLDAPObjectclass returns the IncludeLDAPObjectclass field if non-nil, zero value otherwise.

### GetIncludeLDAPObjectclassOk

`func (o *GetHttpServletExtension200Response) GetIncludeLDAPObjectclassOk() (*[]string, bool)`

GetIncludeLDAPObjectclassOk returns a tuple with the IncludeLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLDAPObjectclass

`func (o *GetHttpServletExtension200Response) SetIncludeLDAPObjectclass(v []string)`

SetIncludeLDAPObjectclass sets IncludeLDAPObjectclass field to given value.

### HasIncludeLDAPObjectclass

`func (o *GetHttpServletExtension200Response) HasIncludeLDAPObjectclass() bool`

HasIncludeLDAPObjectclass returns a boolean if a field has been set.

### GetExcludeLDAPObjectclass

`func (o *GetHttpServletExtension200Response) GetExcludeLDAPObjectclass() []string`

GetExcludeLDAPObjectclass returns the ExcludeLDAPObjectclass field if non-nil, zero value otherwise.

### GetExcludeLDAPObjectclassOk

`func (o *GetHttpServletExtension200Response) GetExcludeLDAPObjectclassOk() (*[]string, bool)`

GetExcludeLDAPObjectclassOk returns a tuple with the ExcludeLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeLDAPObjectclass

`func (o *GetHttpServletExtension200Response) SetExcludeLDAPObjectclass(v []string)`

SetExcludeLDAPObjectclass sets ExcludeLDAPObjectclass field to given value.

### HasExcludeLDAPObjectclass

`func (o *GetHttpServletExtension200Response) HasExcludeLDAPObjectclass() bool`

HasExcludeLDAPObjectclass returns a boolean if a field has been set.

### GetIncludeLDAPBaseDN

`func (o *GetHttpServletExtension200Response) GetIncludeLDAPBaseDN() []string`

GetIncludeLDAPBaseDN returns the IncludeLDAPBaseDN field if non-nil, zero value otherwise.

### GetIncludeLDAPBaseDNOk

`func (o *GetHttpServletExtension200Response) GetIncludeLDAPBaseDNOk() (*[]string, bool)`

GetIncludeLDAPBaseDNOk returns a tuple with the IncludeLDAPBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLDAPBaseDN

`func (o *GetHttpServletExtension200Response) SetIncludeLDAPBaseDN(v []string)`

SetIncludeLDAPBaseDN sets IncludeLDAPBaseDN field to given value.

### HasIncludeLDAPBaseDN

`func (o *GetHttpServletExtension200Response) HasIncludeLDAPBaseDN() bool`

HasIncludeLDAPBaseDN returns a boolean if a field has been set.

### GetExcludeLDAPBaseDN

`func (o *GetHttpServletExtension200Response) GetExcludeLDAPBaseDN() []string`

GetExcludeLDAPBaseDN returns the ExcludeLDAPBaseDN field if non-nil, zero value otherwise.

### GetExcludeLDAPBaseDNOk

`func (o *GetHttpServletExtension200Response) GetExcludeLDAPBaseDNOk() (*[]string, bool)`

GetExcludeLDAPBaseDNOk returns a tuple with the ExcludeLDAPBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeLDAPBaseDN

`func (o *GetHttpServletExtension200Response) SetExcludeLDAPBaseDN(v []string)`

SetExcludeLDAPBaseDN sets ExcludeLDAPBaseDN field to given value.

### HasExcludeLDAPBaseDN

`func (o *GetHttpServletExtension200Response) HasExcludeLDAPBaseDN() bool`

HasExcludeLDAPBaseDN returns a boolean if a field has been set.

### GetEntityTagLDAPAttribute

`func (o *GetHttpServletExtension200Response) GetEntityTagLDAPAttribute() string`

GetEntityTagLDAPAttribute returns the EntityTagLDAPAttribute field if non-nil, zero value otherwise.

### GetEntityTagLDAPAttributeOk

`func (o *GetHttpServletExtension200Response) GetEntityTagLDAPAttributeOk() (*string, bool)`

GetEntityTagLDAPAttributeOk returns a tuple with the EntityTagLDAPAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTagLDAPAttribute

`func (o *GetHttpServletExtension200Response) SetEntityTagLDAPAttribute(v string)`

SetEntityTagLDAPAttribute sets EntityTagLDAPAttribute field to given value.

### HasEntityTagLDAPAttribute

`func (o *GetHttpServletExtension200Response) HasEntityTagLDAPAttribute() bool`

HasEntityTagLDAPAttribute returns a boolean if a field has been set.

### GetTemporaryDirectory

`func (o *GetHttpServletExtension200Response) GetTemporaryDirectory() string`

GetTemporaryDirectory returns the TemporaryDirectory field if non-nil, zero value otherwise.

### GetTemporaryDirectoryOk

`func (o *GetHttpServletExtension200Response) GetTemporaryDirectoryOk() (*string, bool)`

GetTemporaryDirectoryOk returns a tuple with the TemporaryDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryDirectory

`func (o *GetHttpServletExtension200Response) SetTemporaryDirectory(v string)`

SetTemporaryDirectory sets TemporaryDirectory field to given value.


### GetTemporaryDirectoryPermissions

`func (o *GetHttpServletExtension200Response) GetTemporaryDirectoryPermissions() string`

GetTemporaryDirectoryPermissions returns the TemporaryDirectoryPermissions field if non-nil, zero value otherwise.

### GetTemporaryDirectoryPermissionsOk

`func (o *GetHttpServletExtension200Response) GetTemporaryDirectoryPermissionsOk() (*string, bool)`

GetTemporaryDirectoryPermissionsOk returns a tuple with the TemporaryDirectoryPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryDirectoryPermissions

`func (o *GetHttpServletExtension200Response) SetTemporaryDirectoryPermissions(v string)`

SetTemporaryDirectoryPermissions sets TemporaryDirectoryPermissions field to given value.


### GetMaxResults

`func (o *GetHttpServletExtension200Response) GetMaxResults() int64`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *GetHttpServletExtension200Response) GetMaxResultsOk() (*int64, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *GetHttpServletExtension200Response) SetMaxResults(v int64)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *GetHttpServletExtension200Response) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetBulkMaxOperations

`func (o *GetHttpServletExtension200Response) GetBulkMaxOperations() int64`

GetBulkMaxOperations returns the BulkMaxOperations field if non-nil, zero value otherwise.

### GetBulkMaxOperationsOk

`func (o *GetHttpServletExtension200Response) GetBulkMaxOperationsOk() (*int64, bool)`

GetBulkMaxOperationsOk returns a tuple with the BulkMaxOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkMaxOperations

`func (o *GetHttpServletExtension200Response) SetBulkMaxOperations(v int64)`

SetBulkMaxOperations sets BulkMaxOperations field to given value.

### HasBulkMaxOperations

`func (o *GetHttpServletExtension200Response) HasBulkMaxOperations() bool`

HasBulkMaxOperations returns a boolean if a field has been set.

### GetBulkMaxPayloadSize

`func (o *GetHttpServletExtension200Response) GetBulkMaxPayloadSize() string`

GetBulkMaxPayloadSize returns the BulkMaxPayloadSize field if non-nil, zero value otherwise.

### GetBulkMaxPayloadSizeOk

`func (o *GetHttpServletExtension200Response) GetBulkMaxPayloadSizeOk() (*string, bool)`

GetBulkMaxPayloadSizeOk returns a tuple with the BulkMaxPayloadSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkMaxPayloadSize

`func (o *GetHttpServletExtension200Response) SetBulkMaxPayloadSize(v string)`

SetBulkMaxPayloadSize sets BulkMaxPayloadSize field to given value.

### HasBulkMaxPayloadSize

`func (o *GetHttpServletExtension200Response) HasBulkMaxPayloadSize() bool`

HasBulkMaxPayloadSize returns a boolean if a field has been set.

### GetBulkMaxConcurrentRequests

`func (o *GetHttpServletExtension200Response) GetBulkMaxConcurrentRequests() int64`

GetBulkMaxConcurrentRequests returns the BulkMaxConcurrentRequests field if non-nil, zero value otherwise.

### GetBulkMaxConcurrentRequestsOk

`func (o *GetHttpServletExtension200Response) GetBulkMaxConcurrentRequestsOk() (*int64, bool)`

GetBulkMaxConcurrentRequestsOk returns a tuple with the BulkMaxConcurrentRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkMaxConcurrentRequests

`func (o *GetHttpServletExtension200Response) SetBulkMaxConcurrentRequests(v int64)`

SetBulkMaxConcurrentRequests sets BulkMaxConcurrentRequests field to given value.

### HasBulkMaxConcurrentRequests

`func (o *GetHttpServletExtension200Response) HasBulkMaxConcurrentRequests() bool`

HasBulkMaxConcurrentRequests returns a boolean if a field has been set.

### GetDebugEnabled

`func (o *GetHttpServletExtension200Response) GetDebugEnabled() bool`

GetDebugEnabled returns the DebugEnabled field if non-nil, zero value otherwise.

### GetDebugEnabledOk

`func (o *GetHttpServletExtension200Response) GetDebugEnabledOk() (*bool, bool)`

GetDebugEnabledOk returns a tuple with the DebugEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugEnabled

`func (o *GetHttpServletExtension200Response) SetDebugEnabled(v bool)`

SetDebugEnabled sets DebugEnabled field to given value.

### HasDebugEnabled

`func (o *GetHttpServletExtension200Response) HasDebugEnabled() bool`

HasDebugEnabled returns a boolean if a field has been set.

### GetDebugLevel

`func (o *GetHttpServletExtension200Response) GetDebugLevel() EnumhttpServletExtensionDebugLevelProp`

GetDebugLevel returns the DebugLevel field if non-nil, zero value otherwise.

### GetDebugLevelOk

`func (o *GetHttpServletExtension200Response) GetDebugLevelOk() (*EnumhttpServletExtensionDebugLevelProp, bool)`

GetDebugLevelOk returns a tuple with the DebugLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugLevel

`func (o *GetHttpServletExtension200Response) SetDebugLevel(v EnumhttpServletExtensionDebugLevelProp)`

SetDebugLevel sets DebugLevel field to given value.


### GetDebugType

`func (o *GetHttpServletExtension200Response) GetDebugType() []EnumhttpServletExtensionDebugTypeProp`

GetDebugType returns the DebugType field if non-nil, zero value otherwise.

### GetDebugTypeOk

`func (o *GetHttpServletExtension200Response) GetDebugTypeOk() (*[]EnumhttpServletExtensionDebugTypeProp, bool)`

GetDebugTypeOk returns a tuple with the DebugType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugType

`func (o *GetHttpServletExtension200Response) SetDebugType(v []EnumhttpServletExtensionDebugTypeProp)`

SetDebugType sets DebugType field to given value.


### GetIncludeStackTrace

`func (o *GetHttpServletExtension200Response) GetIncludeStackTrace() bool`

GetIncludeStackTrace returns the IncludeStackTrace field if non-nil, zero value otherwise.

### GetIncludeStackTraceOk

`func (o *GetHttpServletExtension200Response) GetIncludeStackTraceOk() (*bool, bool)`

GetIncludeStackTraceOk returns a tuple with the IncludeStackTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeStackTrace

`func (o *GetHttpServletExtension200Response) SetIncludeStackTrace(v bool)`

SetIncludeStackTrace sets IncludeStackTrace field to given value.


### GetScriptClass

`func (o *GetHttpServletExtension200Response) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *GetHttpServletExtension200Response) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *GetHttpServletExtension200Response) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *GetHttpServletExtension200Response) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *GetHttpServletExtension200Response) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *GetHttpServletExtension200Response) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *GetHttpServletExtension200Response) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetDocumentRootDirectory

`func (o *GetHttpServletExtension200Response) GetDocumentRootDirectory() string`

GetDocumentRootDirectory returns the DocumentRootDirectory field if non-nil, zero value otherwise.

### GetDocumentRootDirectoryOk

`func (o *GetHttpServletExtension200Response) GetDocumentRootDirectoryOk() (*string, bool)`

GetDocumentRootDirectoryOk returns a tuple with the DocumentRootDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentRootDirectory

`func (o *GetHttpServletExtension200Response) SetDocumentRootDirectory(v string)`

SetDocumentRootDirectory sets DocumentRootDirectory field to given value.


### GetEnableDirectoryIndexing

`func (o *GetHttpServletExtension200Response) GetEnableDirectoryIndexing() bool`

GetEnableDirectoryIndexing returns the EnableDirectoryIndexing field if non-nil, zero value otherwise.

### GetEnableDirectoryIndexingOk

`func (o *GetHttpServletExtension200Response) GetEnableDirectoryIndexingOk() (*bool, bool)`

GetEnableDirectoryIndexingOk returns a tuple with the EnableDirectoryIndexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDirectoryIndexing

`func (o *GetHttpServletExtension200Response) SetEnableDirectoryIndexing(v bool)`

SetEnableDirectoryIndexing sets EnableDirectoryIndexing field to given value.

### HasEnableDirectoryIndexing

`func (o *GetHttpServletExtension200Response) HasEnableDirectoryIndexing() bool`

HasEnableDirectoryIndexing returns a boolean if a field has been set.

### GetIndexFile

`func (o *GetHttpServletExtension200Response) GetIndexFile() []string`

GetIndexFile returns the IndexFile field if non-nil, zero value otherwise.

### GetIndexFileOk

`func (o *GetHttpServletExtension200Response) GetIndexFileOk() (*[]string, bool)`

GetIndexFileOk returns a tuple with the IndexFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexFile

`func (o *GetHttpServletExtension200Response) SetIndexFile(v []string)`

SetIndexFile sets IndexFile field to given value.

### HasIndexFile

`func (o *GetHttpServletExtension200Response) HasIndexFile() bool`

HasIndexFile returns a boolean if a field has been set.

### GetAllowedAuthenticationType

`func (o *GetHttpServletExtension200Response) GetAllowedAuthenticationType() []EnumhttpServletExtensionAllowedAuthenticationTypeProp`

GetAllowedAuthenticationType returns the AllowedAuthenticationType field if non-nil, zero value otherwise.

### GetAllowedAuthenticationTypeOk

`func (o *GetHttpServletExtension200Response) GetAllowedAuthenticationTypeOk() (*[]EnumhttpServletExtensionAllowedAuthenticationTypeProp, bool)`

GetAllowedAuthenticationTypeOk returns a tuple with the AllowedAuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAuthenticationType

`func (o *GetHttpServletExtension200Response) SetAllowedAuthenticationType(v []EnumhttpServletExtensionAllowedAuthenticationTypeProp)`

SetAllowedAuthenticationType sets AllowedAuthenticationType field to given value.

### HasAllowedAuthenticationType

`func (o *GetHttpServletExtension200Response) HasAllowedAuthenticationType() bool`

HasAllowedAuthenticationType returns a boolean if a field has been set.

### GetIdTokenValidator

`func (o *GetHttpServletExtension200Response) GetIdTokenValidator() []string`

GetIdTokenValidator returns the IdTokenValidator field if non-nil, zero value otherwise.

### GetIdTokenValidatorOk

`func (o *GetHttpServletExtension200Response) GetIdTokenValidatorOk() (*[]string, bool)`

GetIdTokenValidatorOk returns a tuple with the IdTokenValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenValidator

`func (o *GetHttpServletExtension200Response) SetIdTokenValidator(v []string)`

SetIdTokenValidator sets IdTokenValidator field to given value.

### HasIdTokenValidator

`func (o *GetHttpServletExtension200Response) HasIdTokenValidator() bool`

HasIdTokenValidator returns a boolean if a field has been set.

### GetRequireFileServletAccessPrivilege

`func (o *GetHttpServletExtension200Response) GetRequireFileServletAccessPrivilege() bool`

GetRequireFileServletAccessPrivilege returns the RequireFileServletAccessPrivilege field if non-nil, zero value otherwise.

### GetRequireFileServletAccessPrivilegeOk

`func (o *GetHttpServletExtension200Response) GetRequireFileServletAccessPrivilegeOk() (*bool, bool)`

GetRequireFileServletAccessPrivilegeOk returns a tuple with the RequireFileServletAccessPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireFileServletAccessPrivilege

`func (o *GetHttpServletExtension200Response) SetRequireFileServletAccessPrivilege(v bool)`

SetRequireFileServletAccessPrivilege sets RequireFileServletAccessPrivilege field to given value.

### HasRequireFileServletAccessPrivilege

`func (o *GetHttpServletExtension200Response) HasRequireFileServletAccessPrivilege() bool`

HasRequireFileServletAccessPrivilege returns a boolean if a field has been set.

### GetRequireGroup

`func (o *GetHttpServletExtension200Response) GetRequireGroup() []string`

GetRequireGroup returns the RequireGroup field if non-nil, zero value otherwise.

### GetRequireGroupOk

`func (o *GetHttpServletExtension200Response) GetRequireGroupOk() (*[]string, bool)`

GetRequireGroupOk returns a tuple with the RequireGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireGroup

`func (o *GetHttpServletExtension200Response) SetRequireGroup(v []string)`

SetRequireGroup sets RequireGroup field to given value.

### HasRequireGroup

`func (o *GetHttpServletExtension200Response) HasRequireGroup() bool`

HasRequireGroup returns a boolean if a field has been set.

### GetMapAccessTokensToLocalUsers

`func (o *GetHttpServletExtension200Response) GetMapAccessTokensToLocalUsers() EnumhttpServletExtensionMapAccessTokensToLocalUsersProp`

GetMapAccessTokensToLocalUsers returns the MapAccessTokensToLocalUsers field if non-nil, zero value otherwise.

### GetMapAccessTokensToLocalUsersOk

`func (o *GetHttpServletExtension200Response) GetMapAccessTokensToLocalUsersOk() (*EnumhttpServletExtensionMapAccessTokensToLocalUsersProp, bool)`

GetMapAccessTokensToLocalUsersOk returns a tuple with the MapAccessTokensToLocalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapAccessTokensToLocalUsers

`func (o *GetHttpServletExtension200Response) SetMapAccessTokensToLocalUsers(v EnumhttpServletExtensionMapAccessTokensToLocalUsersProp)`

SetMapAccessTokensToLocalUsers sets MapAccessTokensToLocalUsers field to given value.

### HasMapAccessTokensToLocalUsers

`func (o *GetHttpServletExtension200Response) HasMapAccessTokensToLocalUsers() bool`

HasMapAccessTokensToLocalUsers returns a boolean if a field has been set.

### GetSwaggerEnabled

`func (o *GetHttpServletExtension200Response) GetSwaggerEnabled() bool`

GetSwaggerEnabled returns the SwaggerEnabled field if non-nil, zero value otherwise.

### GetSwaggerEnabledOk

`func (o *GetHttpServletExtension200Response) GetSwaggerEnabledOk() (*bool, bool)`

GetSwaggerEnabledOk returns a tuple with the SwaggerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwaggerEnabled

`func (o *GetHttpServletExtension200Response) SetSwaggerEnabled(v bool)`

SetSwaggerEnabled sets SwaggerEnabled field to given value.

### HasSwaggerEnabled

`func (o *GetHttpServletExtension200Response) HasSwaggerEnabled() bool`

HasSwaggerEnabled returns a boolean if a field has been set.

### GetMaxPageSize

`func (o *GetHttpServletExtension200Response) GetMaxPageSize() int64`

GetMaxPageSize returns the MaxPageSize field if non-nil, zero value otherwise.

### GetMaxPageSizeOk

`func (o *GetHttpServletExtension200Response) GetMaxPageSizeOk() (*int64, bool)`

GetMaxPageSizeOk returns a tuple with the MaxPageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPageSize

`func (o *GetHttpServletExtension200Response) SetMaxPageSize(v int64)`

SetMaxPageSize sets MaxPageSize field to given value.

### HasMaxPageSize

`func (o *GetHttpServletExtension200Response) HasMaxPageSize() bool`

HasMaxPageSize returns a boolean if a field has been set.

### GetSchemasEndpointObjectclass

`func (o *GetHttpServletExtension200Response) GetSchemasEndpointObjectclass() []string`

GetSchemasEndpointObjectclass returns the SchemasEndpointObjectclass field if non-nil, zero value otherwise.

### GetSchemasEndpointObjectclassOk

`func (o *GetHttpServletExtension200Response) GetSchemasEndpointObjectclassOk() (*[]string, bool)`

GetSchemasEndpointObjectclassOk returns a tuple with the SchemasEndpointObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemasEndpointObjectclass

`func (o *GetHttpServletExtension200Response) SetSchemasEndpointObjectclass(v []string)`

SetSchemasEndpointObjectclass sets SchemasEndpointObjectclass field to given value.

### HasSchemasEndpointObjectclass

`func (o *GetHttpServletExtension200Response) HasSchemasEndpointObjectclass() bool`

HasSchemasEndpointObjectclass returns a boolean if a field has been set.

### GetDefaultOperationalAttribute

`func (o *GetHttpServletExtension200Response) GetDefaultOperationalAttribute() []string`

GetDefaultOperationalAttribute returns the DefaultOperationalAttribute field if non-nil, zero value otherwise.

### GetDefaultOperationalAttributeOk

`func (o *GetHttpServletExtension200Response) GetDefaultOperationalAttributeOk() (*[]string, bool)`

GetDefaultOperationalAttributeOk returns a tuple with the DefaultOperationalAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOperationalAttribute

`func (o *GetHttpServletExtension200Response) SetDefaultOperationalAttribute(v []string)`

SetDefaultOperationalAttribute sets DefaultOperationalAttribute field to given value.

### HasDefaultOperationalAttribute

`func (o *GetHttpServletExtension200Response) HasDefaultOperationalAttribute() bool`

HasDefaultOperationalAttribute returns a boolean if a field has been set.

### GetRejectExpansionAttribute

`func (o *GetHttpServletExtension200Response) GetRejectExpansionAttribute() []string`

GetRejectExpansionAttribute returns the RejectExpansionAttribute field if non-nil, zero value otherwise.

### GetRejectExpansionAttributeOk

`func (o *GetHttpServletExtension200Response) GetRejectExpansionAttributeOk() (*[]string, bool)`

GetRejectExpansionAttributeOk returns a tuple with the RejectExpansionAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectExpansionAttribute

`func (o *GetHttpServletExtension200Response) SetRejectExpansionAttribute(v []string)`

SetRejectExpansionAttribute sets RejectExpansionAttribute field to given value.

### HasRejectExpansionAttribute

`func (o *GetHttpServletExtension200Response) HasRejectExpansionAttribute() bool`

HasRejectExpansionAttribute returns a boolean if a field has been set.

### GetAlwaysUsePermissiveModify

`func (o *GetHttpServletExtension200Response) GetAlwaysUsePermissiveModify() bool`

GetAlwaysUsePermissiveModify returns the AlwaysUsePermissiveModify field if non-nil, zero value otherwise.

### GetAlwaysUsePermissiveModifyOk

`func (o *GetHttpServletExtension200Response) GetAlwaysUsePermissiveModifyOk() (*bool, bool)`

GetAlwaysUsePermissiveModifyOk returns a tuple with the AlwaysUsePermissiveModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysUsePermissiveModify

`func (o *GetHttpServletExtension200Response) SetAlwaysUsePermissiveModify(v bool)`

SetAlwaysUsePermissiveModify sets AlwaysUsePermissiveModify field to given value.

### HasAlwaysUsePermissiveModify

`func (o *GetHttpServletExtension200Response) HasAlwaysUsePermissiveModify() bool`

HasAlwaysUsePermissiveModify returns a boolean if a field has been set.

### GetAllowedControl

`func (o *GetHttpServletExtension200Response) GetAllowedControl() []EnumhttpServletExtensionAllowedControlProp`

GetAllowedControl returns the AllowedControl field if non-nil, zero value otherwise.

### GetAllowedControlOk

`func (o *GetHttpServletExtension200Response) GetAllowedControlOk() (*[]EnumhttpServletExtensionAllowedControlProp, bool)`

GetAllowedControlOk returns a tuple with the AllowedControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedControl

`func (o *GetHttpServletExtension200Response) SetAllowedControl(v []EnumhttpServletExtensionAllowedControlProp)`

SetAllowedControl sets AllowedControl field to given value.

### HasAllowedControl

`func (o *GetHttpServletExtension200Response) HasAllowedControl() bool`

HasAllowedControl returns a boolean if a field has been set.

### GetExtensionClass

`func (o *GetHttpServletExtension200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *GetHttpServletExtension200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *GetHttpServletExtension200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *GetHttpServletExtension200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *GetHttpServletExtension200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *GetHttpServletExtension200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *GetHttpServletExtension200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


