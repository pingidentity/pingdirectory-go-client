# AddHttpServletExtensionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtensionName** | **string** | Name of the new HTTP Servlet Extension | 
**Schemas** | [**[]EnumthirdPartyHttpServletExtensionSchemaUrn**](EnumthirdPartyHttpServletExtensionSchemaUrn.md) |  | 
**Server** | Pointer to **string** | Specifies the PingFederate server to be configured. | [optional] 
**Description** | Pointer to **string** | A description for this HTTP Servlet Extension | [optional] 
**CrossOriginPolicy** | Pointer to **string** | The cross-origin request policy to use for the HTTP Servlet Extension. | [optional] 
**ResponseHeader** | Pointer to **[]string** | Specifies HTTP header fields and values added to response headers for all requests. | [optional] 
**CorrelationIDResponseHeader** | Pointer to **string** | Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \&quot;Correlation-Id\&quot;, \&quot;X-Amzn-Trace-Id\&quot;, and \&quot;X-Request-Id\&quot;. | [optional] 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted HTTP Servlet Extension. | 
**ScriptArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Scripted HTTP Servlet Extension. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**BaseContextPath** | **string** | The context path to use to access the SCIM interface. The value must start with a forward slash and must represent a valid HTTP context path. | 
**AvailableStatusCode** | **int32** | Specifies the HTTP status code that the servlet should return if the server considers itself to be available. | 
**DegradedStatusCode** | **int32** | Specifies the HTTP status code that the servlet should return if the server considers itself to be degraded. | 
**UnavailableStatusCode** | **int32** | Specifies the HTTP status code that the servlet should return if the server considers itself to be unavailable. | 
**OverrideStatusCode** | Pointer to **int32** | Specifies a HTTP status code that the servlet should always return, regardless of the server&#39;s availability. If this value is defined, it will override the availability-based return codes. | [optional] 
**IncludeResponseBody** | Pointer to **bool** | Indicates whether the response should include a body that is a JSON object. | [optional] 
**AdditionalResponseContents** | Pointer to **string** | A JSON-formatted string containing additional fields to be returned in the response body. For example, an additional-response-contents value of &#39;{ \&quot;key\&quot;: \&quot;value\&quot; }&#39; would add the key and value to the root of the JSON response body. | [optional] 
**DocumentRootDirectory** | **string** | Specifies the path to the directory on the local filesystem containing the files to be served by this File Server HTTP Servlet Extension. The path must exist, and it must be a directory. | 
**EnableDirectoryIndexing** | Pointer to **bool** | Indicates whether to generate a default HTML page with a listing of available files if the requested path refers to a directory rather than a file, and that directory does not contain an index file. | [optional] 
**IndexFile** | Pointer to **[]string** | Specifies the name of a file whose contents may be returned to the client if the requested path refers to a directory rather than a file. | [optional] 
**MimeTypesFile** | Pointer to **string** | Specifies the path to a file that contains MIME type mappings that will be used to determine the appropriate value to return for the Content-Type header based on the extension of the requested file. | [optional] 
**DefaultMIMEType** | Pointer to **string** | Specifies the default MIME type to use for the Content-Type header when a mapping cannot be found. | [optional] 
**RequireAuthentication** | Pointer to **bool** | Indicates whether the servlet extension should only accept requests from authenticated clients. | [optional] 
**AllowedAuthenticationType** | Pointer to [**[]EnumhttpServletExtensionAllowedAuthenticationTypeProp**](EnumhttpServletExtensionAllowedAuthenticationTypeProp.md) |  | [optional] 
**AccessTokenValidator** | Pointer to **[]string** | The access token validators that may be used to verify the authenticity of an OAuth 2.0 bearer token. | [optional] 
**IdTokenValidator** | Pointer to **[]string** | The ID token validators that may be used to verify the authenticity of an of an OpenID Connect ID token. | [optional] 
**RequireFileServletAccessPrivilege** | Pointer to **bool** | Indicates whether the servlet extension should only accept requests from authenticated clients that have the file-servlet-access privilege. | [optional] 
**RequireGroup** | Pointer to **[]string** | The DN of a group whose members will be permitted to access to the associated files. If multiple group DNs are configured, then anyone who is a member of at least one of those groups will be granted access. | [optional] 
**IdentityMapper** | Pointer to **string** | Specifies the name of the identity mapper that is to be used to match the username included in the HTTP Basic authentication header to the corresponding user in the directory. | [optional] 
**OAuthTokenHandler** | Pointer to **string** | Specifies the OAuth Token Handler implementation that should be used to validate OAuth 2.0 bearer tokens when they are included in a SCIM request. | [optional] 
**BasicAuthEnabled** | Pointer to **bool** | Enables HTTP Basic authentication, using a username and password. | [optional] 
**ResourceMappingFile** | Pointer to **string** | The path to an XML file defining the resources supported by the SCIM interface and the SCIM-to-LDAP attribute mappings to use. | [optional] 
**IncludeLDAPObjectclass** | Pointer to **[]string** | Specifies the LDAP object classes that should be exposed directly as SCIM resources. | [optional] 
**ExcludeLDAPObjectclass** | Pointer to **[]string** | Specifies the LDAP object classes that should be not be exposed directly as SCIM resources. | [optional] 
**IncludeLDAPBaseDN** | Pointer to **[]string** | Specifies the base DNs for the branches of the DIT that should be exposed via the Identity Access API. | [optional] 
**ExcludeLDAPBaseDN** | Pointer to **[]string** | Specifies the base DNs for the branches of the DIT that should not be exposed via the Identity Access API. | [optional] 
**EntityTagLDAPAttribute** | Pointer to **string** | Specifies the LDAP attribute whose value should be used as the entity tag value to enable SCIM resource versioning support. | [optional] 
**TemporaryDirectory** | Pointer to **string** | Specifies the location of the directory that is used to create temporary files containing SCIM request data. | [optional] 
**TemporaryDirectoryPermissions** | Pointer to **string** | Specifies the permissions that should be applied to the directory that is used to create temporary files. | [optional] 
**MaxResults** | Pointer to **int32** | The maximum number of resources that are returned in a response. | [optional] 
**BulkMaxOperations** | Pointer to **int32** | The maximum number of operations that are permitted in a bulk request. | [optional] 
**BulkMaxPayloadSize** | Pointer to **string** | The maximum payload size in bytes of a bulk request. | [optional] 
**BulkMaxConcurrentRequests** | Pointer to **int32** | The maximum number of bulk requests that may be processed concurrently by the server. Any bulk request that would cause this limit to be exceeded is rejected with HTTP status code 503. | [optional] 
**DebugEnabled** | Pointer to **bool** | Enables debug logging of the SCIM SDK. Debug messages will be forwarded to the Directory Server debug logger with the scope of com.unboundid.directory.server.extensions.scim.SCIMHTTPServletExtension. | [optional] 
**DebugLevel** | Pointer to [**EnumhttpServletExtensionDebugLevelProp**](EnumhttpServletExtensionDebugLevelProp.md) |  | [optional] 
**DebugType** | Pointer to [**[]EnumhttpServletExtensionDebugTypeProp**](EnumhttpServletExtensionDebugTypeProp.md) |  | [optional] 
**IncludeStackTrace** | Pointer to **bool** | Indicates whether a stack trace of the thread which called the debug method should be included in debug log messages. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party HTTP Servlet Extension. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party HTTP Servlet Extension. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddHttpServletExtensionRequest

`func NewAddHttpServletExtensionRequest(extensionName string, schemas []EnumthirdPartyHttpServletExtensionSchemaUrn, scriptClass string, baseContextPath string, availableStatusCode int32, degradedStatusCode int32, unavailableStatusCode int32, documentRootDirectory string, extensionClass string, ) *AddHttpServletExtensionRequest`

NewAddHttpServletExtensionRequest instantiates a new AddHttpServletExtensionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddHttpServletExtensionRequestWithDefaults

`func NewAddHttpServletExtensionRequestWithDefaults() *AddHttpServletExtensionRequest`

NewAddHttpServletExtensionRequestWithDefaults instantiates a new AddHttpServletExtensionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensionName

`func (o *AddHttpServletExtensionRequest) GetExtensionName() string`

GetExtensionName returns the ExtensionName field if non-nil, zero value otherwise.

### GetExtensionNameOk

`func (o *AddHttpServletExtensionRequest) GetExtensionNameOk() (*string, bool)`

GetExtensionNameOk returns a tuple with the ExtensionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionName

`func (o *AddHttpServletExtensionRequest) SetExtensionName(v string)`

SetExtensionName sets ExtensionName field to given value.


### GetSchemas

`func (o *AddHttpServletExtensionRequest) GetSchemas() []EnumthirdPartyHttpServletExtensionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddHttpServletExtensionRequest) GetSchemasOk() (*[]EnumthirdPartyHttpServletExtensionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddHttpServletExtensionRequest) SetSchemas(v []EnumthirdPartyHttpServletExtensionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetServer

`func (o *AddHttpServletExtensionRequest) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *AddHttpServletExtensionRequest) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *AddHttpServletExtensionRequest) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *AddHttpServletExtensionRequest) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetDescription

`func (o *AddHttpServletExtensionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddHttpServletExtensionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddHttpServletExtensionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddHttpServletExtensionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCrossOriginPolicy

`func (o *AddHttpServletExtensionRequest) GetCrossOriginPolicy() string`

GetCrossOriginPolicy returns the CrossOriginPolicy field if non-nil, zero value otherwise.

### GetCrossOriginPolicyOk

`func (o *AddHttpServletExtensionRequest) GetCrossOriginPolicyOk() (*string, bool)`

GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossOriginPolicy

`func (o *AddHttpServletExtensionRequest) SetCrossOriginPolicy(v string)`

SetCrossOriginPolicy sets CrossOriginPolicy field to given value.

### HasCrossOriginPolicy

`func (o *AddHttpServletExtensionRequest) HasCrossOriginPolicy() bool`

HasCrossOriginPolicy returns a boolean if a field has been set.

### GetResponseHeader

`func (o *AddHttpServletExtensionRequest) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *AddHttpServletExtensionRequest) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *AddHttpServletExtensionRequest) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *AddHttpServletExtensionRequest) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetCorrelationIDResponseHeader

`func (o *AddHttpServletExtensionRequest) GetCorrelationIDResponseHeader() string`

GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field if non-nil, zero value otherwise.

### GetCorrelationIDResponseHeaderOk

`func (o *AddHttpServletExtensionRequest) GetCorrelationIDResponseHeaderOk() (*string, bool)`

GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDResponseHeader

`func (o *AddHttpServletExtensionRequest) SetCorrelationIDResponseHeader(v string)`

SetCorrelationIDResponseHeader sets CorrelationIDResponseHeader field to given value.

### HasCorrelationIDResponseHeader

`func (o *AddHttpServletExtensionRequest) HasCorrelationIDResponseHeader() bool`

HasCorrelationIDResponseHeader returns a boolean if a field has been set.

### GetScriptClass

`func (o *AddHttpServletExtensionRequest) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *AddHttpServletExtensionRequest) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *AddHttpServletExtensionRequest) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *AddHttpServletExtensionRequest) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *AddHttpServletExtensionRequest) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *AddHttpServletExtensionRequest) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *AddHttpServletExtensionRequest) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetBaseContextPath

`func (o *AddHttpServletExtensionRequest) GetBaseContextPath() string`

GetBaseContextPath returns the BaseContextPath field if non-nil, zero value otherwise.

### GetBaseContextPathOk

`func (o *AddHttpServletExtensionRequest) GetBaseContextPathOk() (*string, bool)`

GetBaseContextPathOk returns a tuple with the BaseContextPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseContextPath

`func (o *AddHttpServletExtensionRequest) SetBaseContextPath(v string)`

SetBaseContextPath sets BaseContextPath field to given value.


### GetAvailableStatusCode

`func (o *AddHttpServletExtensionRequest) GetAvailableStatusCode() int32`

GetAvailableStatusCode returns the AvailableStatusCode field if non-nil, zero value otherwise.

### GetAvailableStatusCodeOk

`func (o *AddHttpServletExtensionRequest) GetAvailableStatusCodeOk() (*int32, bool)`

GetAvailableStatusCodeOk returns a tuple with the AvailableStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableStatusCode

`func (o *AddHttpServletExtensionRequest) SetAvailableStatusCode(v int32)`

SetAvailableStatusCode sets AvailableStatusCode field to given value.


### GetDegradedStatusCode

`func (o *AddHttpServletExtensionRequest) GetDegradedStatusCode() int32`

GetDegradedStatusCode returns the DegradedStatusCode field if non-nil, zero value otherwise.

### GetDegradedStatusCodeOk

`func (o *AddHttpServletExtensionRequest) GetDegradedStatusCodeOk() (*int32, bool)`

GetDegradedStatusCodeOk returns a tuple with the DegradedStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegradedStatusCode

`func (o *AddHttpServletExtensionRequest) SetDegradedStatusCode(v int32)`

SetDegradedStatusCode sets DegradedStatusCode field to given value.


### GetUnavailableStatusCode

`func (o *AddHttpServletExtensionRequest) GetUnavailableStatusCode() int32`

GetUnavailableStatusCode returns the UnavailableStatusCode field if non-nil, zero value otherwise.

### GetUnavailableStatusCodeOk

`func (o *AddHttpServletExtensionRequest) GetUnavailableStatusCodeOk() (*int32, bool)`

GetUnavailableStatusCodeOk returns a tuple with the UnavailableStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailableStatusCode

`func (o *AddHttpServletExtensionRequest) SetUnavailableStatusCode(v int32)`

SetUnavailableStatusCode sets UnavailableStatusCode field to given value.


### GetOverrideStatusCode

`func (o *AddHttpServletExtensionRequest) GetOverrideStatusCode() int32`

GetOverrideStatusCode returns the OverrideStatusCode field if non-nil, zero value otherwise.

### GetOverrideStatusCodeOk

`func (o *AddHttpServletExtensionRequest) GetOverrideStatusCodeOk() (*int32, bool)`

GetOverrideStatusCodeOk returns a tuple with the OverrideStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideStatusCode

`func (o *AddHttpServletExtensionRequest) SetOverrideStatusCode(v int32)`

SetOverrideStatusCode sets OverrideStatusCode field to given value.

### HasOverrideStatusCode

`func (o *AddHttpServletExtensionRequest) HasOverrideStatusCode() bool`

HasOverrideStatusCode returns a boolean if a field has been set.

### GetIncludeResponseBody

`func (o *AddHttpServletExtensionRequest) GetIncludeResponseBody() bool`

GetIncludeResponseBody returns the IncludeResponseBody field if non-nil, zero value otherwise.

### GetIncludeResponseBodyOk

`func (o *AddHttpServletExtensionRequest) GetIncludeResponseBodyOk() (*bool, bool)`

GetIncludeResponseBodyOk returns a tuple with the IncludeResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeResponseBody

`func (o *AddHttpServletExtensionRequest) SetIncludeResponseBody(v bool)`

SetIncludeResponseBody sets IncludeResponseBody field to given value.

### HasIncludeResponseBody

`func (o *AddHttpServletExtensionRequest) HasIncludeResponseBody() bool`

HasIncludeResponseBody returns a boolean if a field has been set.

### GetAdditionalResponseContents

`func (o *AddHttpServletExtensionRequest) GetAdditionalResponseContents() string`

GetAdditionalResponseContents returns the AdditionalResponseContents field if non-nil, zero value otherwise.

### GetAdditionalResponseContentsOk

`func (o *AddHttpServletExtensionRequest) GetAdditionalResponseContentsOk() (*string, bool)`

GetAdditionalResponseContentsOk returns a tuple with the AdditionalResponseContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalResponseContents

`func (o *AddHttpServletExtensionRequest) SetAdditionalResponseContents(v string)`

SetAdditionalResponseContents sets AdditionalResponseContents field to given value.

### HasAdditionalResponseContents

`func (o *AddHttpServletExtensionRequest) HasAdditionalResponseContents() bool`

HasAdditionalResponseContents returns a boolean if a field has been set.

### GetDocumentRootDirectory

`func (o *AddHttpServletExtensionRequest) GetDocumentRootDirectory() string`

GetDocumentRootDirectory returns the DocumentRootDirectory field if non-nil, zero value otherwise.

### GetDocumentRootDirectoryOk

`func (o *AddHttpServletExtensionRequest) GetDocumentRootDirectoryOk() (*string, bool)`

GetDocumentRootDirectoryOk returns a tuple with the DocumentRootDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentRootDirectory

`func (o *AddHttpServletExtensionRequest) SetDocumentRootDirectory(v string)`

SetDocumentRootDirectory sets DocumentRootDirectory field to given value.


### GetEnableDirectoryIndexing

`func (o *AddHttpServletExtensionRequest) GetEnableDirectoryIndexing() bool`

GetEnableDirectoryIndexing returns the EnableDirectoryIndexing field if non-nil, zero value otherwise.

### GetEnableDirectoryIndexingOk

`func (o *AddHttpServletExtensionRequest) GetEnableDirectoryIndexingOk() (*bool, bool)`

GetEnableDirectoryIndexingOk returns a tuple with the EnableDirectoryIndexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDirectoryIndexing

`func (o *AddHttpServletExtensionRequest) SetEnableDirectoryIndexing(v bool)`

SetEnableDirectoryIndexing sets EnableDirectoryIndexing field to given value.

### HasEnableDirectoryIndexing

`func (o *AddHttpServletExtensionRequest) HasEnableDirectoryIndexing() bool`

HasEnableDirectoryIndexing returns a boolean if a field has been set.

### GetIndexFile

`func (o *AddHttpServletExtensionRequest) GetIndexFile() []string`

GetIndexFile returns the IndexFile field if non-nil, zero value otherwise.

### GetIndexFileOk

`func (o *AddHttpServletExtensionRequest) GetIndexFileOk() (*[]string, bool)`

GetIndexFileOk returns a tuple with the IndexFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexFile

`func (o *AddHttpServletExtensionRequest) SetIndexFile(v []string)`

SetIndexFile sets IndexFile field to given value.

### HasIndexFile

`func (o *AddHttpServletExtensionRequest) HasIndexFile() bool`

HasIndexFile returns a boolean if a field has been set.

### GetMimeTypesFile

`func (o *AddHttpServletExtensionRequest) GetMimeTypesFile() string`

GetMimeTypesFile returns the MimeTypesFile field if non-nil, zero value otherwise.

### GetMimeTypesFileOk

`func (o *AddHttpServletExtensionRequest) GetMimeTypesFileOk() (*string, bool)`

GetMimeTypesFileOk returns a tuple with the MimeTypesFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeTypesFile

`func (o *AddHttpServletExtensionRequest) SetMimeTypesFile(v string)`

SetMimeTypesFile sets MimeTypesFile field to given value.

### HasMimeTypesFile

`func (o *AddHttpServletExtensionRequest) HasMimeTypesFile() bool`

HasMimeTypesFile returns a boolean if a field has been set.

### GetDefaultMIMEType

`func (o *AddHttpServletExtensionRequest) GetDefaultMIMEType() string`

GetDefaultMIMEType returns the DefaultMIMEType field if non-nil, zero value otherwise.

### GetDefaultMIMETypeOk

`func (o *AddHttpServletExtensionRequest) GetDefaultMIMETypeOk() (*string, bool)`

GetDefaultMIMETypeOk returns a tuple with the DefaultMIMEType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMIMEType

`func (o *AddHttpServletExtensionRequest) SetDefaultMIMEType(v string)`

SetDefaultMIMEType sets DefaultMIMEType field to given value.

### HasDefaultMIMEType

`func (o *AddHttpServletExtensionRequest) HasDefaultMIMEType() bool`

HasDefaultMIMEType returns a boolean if a field has been set.

### GetRequireAuthentication

`func (o *AddHttpServletExtensionRequest) GetRequireAuthentication() bool`

GetRequireAuthentication returns the RequireAuthentication field if non-nil, zero value otherwise.

### GetRequireAuthenticationOk

`func (o *AddHttpServletExtensionRequest) GetRequireAuthenticationOk() (*bool, bool)`

GetRequireAuthenticationOk returns a tuple with the RequireAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAuthentication

`func (o *AddHttpServletExtensionRequest) SetRequireAuthentication(v bool)`

SetRequireAuthentication sets RequireAuthentication field to given value.

### HasRequireAuthentication

`func (o *AddHttpServletExtensionRequest) HasRequireAuthentication() bool`

HasRequireAuthentication returns a boolean if a field has been set.

### GetAllowedAuthenticationType

`func (o *AddHttpServletExtensionRequest) GetAllowedAuthenticationType() []EnumhttpServletExtensionAllowedAuthenticationTypeProp`

GetAllowedAuthenticationType returns the AllowedAuthenticationType field if non-nil, zero value otherwise.

### GetAllowedAuthenticationTypeOk

`func (o *AddHttpServletExtensionRequest) GetAllowedAuthenticationTypeOk() (*[]EnumhttpServletExtensionAllowedAuthenticationTypeProp, bool)`

GetAllowedAuthenticationTypeOk returns a tuple with the AllowedAuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAuthenticationType

`func (o *AddHttpServletExtensionRequest) SetAllowedAuthenticationType(v []EnumhttpServletExtensionAllowedAuthenticationTypeProp)`

SetAllowedAuthenticationType sets AllowedAuthenticationType field to given value.

### HasAllowedAuthenticationType

`func (o *AddHttpServletExtensionRequest) HasAllowedAuthenticationType() bool`

HasAllowedAuthenticationType returns a boolean if a field has been set.

### GetAccessTokenValidator

`func (o *AddHttpServletExtensionRequest) GetAccessTokenValidator() []string`

GetAccessTokenValidator returns the AccessTokenValidator field if non-nil, zero value otherwise.

### GetAccessTokenValidatorOk

`func (o *AddHttpServletExtensionRequest) GetAccessTokenValidatorOk() (*[]string, bool)`

GetAccessTokenValidatorOk returns a tuple with the AccessTokenValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValidator

`func (o *AddHttpServletExtensionRequest) SetAccessTokenValidator(v []string)`

SetAccessTokenValidator sets AccessTokenValidator field to given value.

### HasAccessTokenValidator

`func (o *AddHttpServletExtensionRequest) HasAccessTokenValidator() bool`

HasAccessTokenValidator returns a boolean if a field has been set.

### GetIdTokenValidator

`func (o *AddHttpServletExtensionRequest) GetIdTokenValidator() []string`

GetIdTokenValidator returns the IdTokenValidator field if non-nil, zero value otherwise.

### GetIdTokenValidatorOk

`func (o *AddHttpServletExtensionRequest) GetIdTokenValidatorOk() (*[]string, bool)`

GetIdTokenValidatorOk returns a tuple with the IdTokenValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenValidator

`func (o *AddHttpServletExtensionRequest) SetIdTokenValidator(v []string)`

SetIdTokenValidator sets IdTokenValidator field to given value.

### HasIdTokenValidator

`func (o *AddHttpServletExtensionRequest) HasIdTokenValidator() bool`

HasIdTokenValidator returns a boolean if a field has been set.

### GetRequireFileServletAccessPrivilege

`func (o *AddHttpServletExtensionRequest) GetRequireFileServletAccessPrivilege() bool`

GetRequireFileServletAccessPrivilege returns the RequireFileServletAccessPrivilege field if non-nil, zero value otherwise.

### GetRequireFileServletAccessPrivilegeOk

`func (o *AddHttpServletExtensionRequest) GetRequireFileServletAccessPrivilegeOk() (*bool, bool)`

GetRequireFileServletAccessPrivilegeOk returns a tuple with the RequireFileServletAccessPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireFileServletAccessPrivilege

`func (o *AddHttpServletExtensionRequest) SetRequireFileServletAccessPrivilege(v bool)`

SetRequireFileServletAccessPrivilege sets RequireFileServletAccessPrivilege field to given value.

### HasRequireFileServletAccessPrivilege

`func (o *AddHttpServletExtensionRequest) HasRequireFileServletAccessPrivilege() bool`

HasRequireFileServletAccessPrivilege returns a boolean if a field has been set.

### GetRequireGroup

`func (o *AddHttpServletExtensionRequest) GetRequireGroup() []string`

GetRequireGroup returns the RequireGroup field if non-nil, zero value otherwise.

### GetRequireGroupOk

`func (o *AddHttpServletExtensionRequest) GetRequireGroupOk() (*[]string, bool)`

GetRequireGroupOk returns a tuple with the RequireGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireGroup

`func (o *AddHttpServletExtensionRequest) SetRequireGroup(v []string)`

SetRequireGroup sets RequireGroup field to given value.

### HasRequireGroup

`func (o *AddHttpServletExtensionRequest) HasRequireGroup() bool`

HasRequireGroup returns a boolean if a field has been set.

### GetIdentityMapper

`func (o *AddHttpServletExtensionRequest) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *AddHttpServletExtensionRequest) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *AddHttpServletExtensionRequest) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.

### HasIdentityMapper

`func (o *AddHttpServletExtensionRequest) HasIdentityMapper() bool`

HasIdentityMapper returns a boolean if a field has been set.

### GetOAuthTokenHandler

`func (o *AddHttpServletExtensionRequest) GetOAuthTokenHandler() string`

GetOAuthTokenHandler returns the OAuthTokenHandler field if non-nil, zero value otherwise.

### GetOAuthTokenHandlerOk

`func (o *AddHttpServletExtensionRequest) GetOAuthTokenHandlerOk() (*string, bool)`

GetOAuthTokenHandlerOk returns a tuple with the OAuthTokenHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthTokenHandler

`func (o *AddHttpServletExtensionRequest) SetOAuthTokenHandler(v string)`

SetOAuthTokenHandler sets OAuthTokenHandler field to given value.

### HasOAuthTokenHandler

`func (o *AddHttpServletExtensionRequest) HasOAuthTokenHandler() bool`

HasOAuthTokenHandler returns a boolean if a field has been set.

### GetBasicAuthEnabled

`func (o *AddHttpServletExtensionRequest) GetBasicAuthEnabled() bool`

GetBasicAuthEnabled returns the BasicAuthEnabled field if non-nil, zero value otherwise.

### GetBasicAuthEnabledOk

`func (o *AddHttpServletExtensionRequest) GetBasicAuthEnabledOk() (*bool, bool)`

GetBasicAuthEnabledOk returns a tuple with the BasicAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthEnabled

`func (o *AddHttpServletExtensionRequest) SetBasicAuthEnabled(v bool)`

SetBasicAuthEnabled sets BasicAuthEnabled field to given value.

### HasBasicAuthEnabled

`func (o *AddHttpServletExtensionRequest) HasBasicAuthEnabled() bool`

HasBasicAuthEnabled returns a boolean if a field has been set.

### GetResourceMappingFile

`func (o *AddHttpServletExtensionRequest) GetResourceMappingFile() string`

GetResourceMappingFile returns the ResourceMappingFile field if non-nil, zero value otherwise.

### GetResourceMappingFileOk

`func (o *AddHttpServletExtensionRequest) GetResourceMappingFileOk() (*string, bool)`

GetResourceMappingFileOk returns a tuple with the ResourceMappingFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceMappingFile

`func (o *AddHttpServletExtensionRequest) SetResourceMappingFile(v string)`

SetResourceMappingFile sets ResourceMappingFile field to given value.

### HasResourceMappingFile

`func (o *AddHttpServletExtensionRequest) HasResourceMappingFile() bool`

HasResourceMappingFile returns a boolean if a field has been set.

### GetIncludeLDAPObjectclass

`func (o *AddHttpServletExtensionRequest) GetIncludeLDAPObjectclass() []string`

GetIncludeLDAPObjectclass returns the IncludeLDAPObjectclass field if non-nil, zero value otherwise.

### GetIncludeLDAPObjectclassOk

`func (o *AddHttpServletExtensionRequest) GetIncludeLDAPObjectclassOk() (*[]string, bool)`

GetIncludeLDAPObjectclassOk returns a tuple with the IncludeLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLDAPObjectclass

`func (o *AddHttpServletExtensionRequest) SetIncludeLDAPObjectclass(v []string)`

SetIncludeLDAPObjectclass sets IncludeLDAPObjectclass field to given value.

### HasIncludeLDAPObjectclass

`func (o *AddHttpServletExtensionRequest) HasIncludeLDAPObjectclass() bool`

HasIncludeLDAPObjectclass returns a boolean if a field has been set.

### GetExcludeLDAPObjectclass

`func (o *AddHttpServletExtensionRequest) GetExcludeLDAPObjectclass() []string`

GetExcludeLDAPObjectclass returns the ExcludeLDAPObjectclass field if non-nil, zero value otherwise.

### GetExcludeLDAPObjectclassOk

`func (o *AddHttpServletExtensionRequest) GetExcludeLDAPObjectclassOk() (*[]string, bool)`

GetExcludeLDAPObjectclassOk returns a tuple with the ExcludeLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeLDAPObjectclass

`func (o *AddHttpServletExtensionRequest) SetExcludeLDAPObjectclass(v []string)`

SetExcludeLDAPObjectclass sets ExcludeLDAPObjectclass field to given value.

### HasExcludeLDAPObjectclass

`func (o *AddHttpServletExtensionRequest) HasExcludeLDAPObjectclass() bool`

HasExcludeLDAPObjectclass returns a boolean if a field has been set.

### GetIncludeLDAPBaseDN

`func (o *AddHttpServletExtensionRequest) GetIncludeLDAPBaseDN() []string`

GetIncludeLDAPBaseDN returns the IncludeLDAPBaseDN field if non-nil, zero value otherwise.

### GetIncludeLDAPBaseDNOk

`func (o *AddHttpServletExtensionRequest) GetIncludeLDAPBaseDNOk() (*[]string, bool)`

GetIncludeLDAPBaseDNOk returns a tuple with the IncludeLDAPBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLDAPBaseDN

`func (o *AddHttpServletExtensionRequest) SetIncludeLDAPBaseDN(v []string)`

SetIncludeLDAPBaseDN sets IncludeLDAPBaseDN field to given value.

### HasIncludeLDAPBaseDN

`func (o *AddHttpServletExtensionRequest) HasIncludeLDAPBaseDN() bool`

HasIncludeLDAPBaseDN returns a boolean if a field has been set.

### GetExcludeLDAPBaseDN

`func (o *AddHttpServletExtensionRequest) GetExcludeLDAPBaseDN() []string`

GetExcludeLDAPBaseDN returns the ExcludeLDAPBaseDN field if non-nil, zero value otherwise.

### GetExcludeLDAPBaseDNOk

`func (o *AddHttpServletExtensionRequest) GetExcludeLDAPBaseDNOk() (*[]string, bool)`

GetExcludeLDAPBaseDNOk returns a tuple with the ExcludeLDAPBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeLDAPBaseDN

`func (o *AddHttpServletExtensionRequest) SetExcludeLDAPBaseDN(v []string)`

SetExcludeLDAPBaseDN sets ExcludeLDAPBaseDN field to given value.

### HasExcludeLDAPBaseDN

`func (o *AddHttpServletExtensionRequest) HasExcludeLDAPBaseDN() bool`

HasExcludeLDAPBaseDN returns a boolean if a field has been set.

### GetEntityTagLDAPAttribute

`func (o *AddHttpServletExtensionRequest) GetEntityTagLDAPAttribute() string`

GetEntityTagLDAPAttribute returns the EntityTagLDAPAttribute field if non-nil, zero value otherwise.

### GetEntityTagLDAPAttributeOk

`func (o *AddHttpServletExtensionRequest) GetEntityTagLDAPAttributeOk() (*string, bool)`

GetEntityTagLDAPAttributeOk returns a tuple with the EntityTagLDAPAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTagLDAPAttribute

`func (o *AddHttpServletExtensionRequest) SetEntityTagLDAPAttribute(v string)`

SetEntityTagLDAPAttribute sets EntityTagLDAPAttribute field to given value.

### HasEntityTagLDAPAttribute

`func (o *AddHttpServletExtensionRequest) HasEntityTagLDAPAttribute() bool`

HasEntityTagLDAPAttribute returns a boolean if a field has been set.

### GetTemporaryDirectory

`func (o *AddHttpServletExtensionRequest) GetTemporaryDirectory() string`

GetTemporaryDirectory returns the TemporaryDirectory field if non-nil, zero value otherwise.

### GetTemporaryDirectoryOk

`func (o *AddHttpServletExtensionRequest) GetTemporaryDirectoryOk() (*string, bool)`

GetTemporaryDirectoryOk returns a tuple with the TemporaryDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryDirectory

`func (o *AddHttpServletExtensionRequest) SetTemporaryDirectory(v string)`

SetTemporaryDirectory sets TemporaryDirectory field to given value.

### HasTemporaryDirectory

`func (o *AddHttpServletExtensionRequest) HasTemporaryDirectory() bool`

HasTemporaryDirectory returns a boolean if a field has been set.

### GetTemporaryDirectoryPermissions

`func (o *AddHttpServletExtensionRequest) GetTemporaryDirectoryPermissions() string`

GetTemporaryDirectoryPermissions returns the TemporaryDirectoryPermissions field if non-nil, zero value otherwise.

### GetTemporaryDirectoryPermissionsOk

`func (o *AddHttpServletExtensionRequest) GetTemporaryDirectoryPermissionsOk() (*string, bool)`

GetTemporaryDirectoryPermissionsOk returns a tuple with the TemporaryDirectoryPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryDirectoryPermissions

`func (o *AddHttpServletExtensionRequest) SetTemporaryDirectoryPermissions(v string)`

SetTemporaryDirectoryPermissions sets TemporaryDirectoryPermissions field to given value.

### HasTemporaryDirectoryPermissions

`func (o *AddHttpServletExtensionRequest) HasTemporaryDirectoryPermissions() bool`

HasTemporaryDirectoryPermissions returns a boolean if a field has been set.

### GetMaxResults

`func (o *AddHttpServletExtensionRequest) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *AddHttpServletExtensionRequest) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *AddHttpServletExtensionRequest) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *AddHttpServletExtensionRequest) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetBulkMaxOperations

`func (o *AddHttpServletExtensionRequest) GetBulkMaxOperations() int32`

GetBulkMaxOperations returns the BulkMaxOperations field if non-nil, zero value otherwise.

### GetBulkMaxOperationsOk

`func (o *AddHttpServletExtensionRequest) GetBulkMaxOperationsOk() (*int32, bool)`

GetBulkMaxOperationsOk returns a tuple with the BulkMaxOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkMaxOperations

`func (o *AddHttpServletExtensionRequest) SetBulkMaxOperations(v int32)`

SetBulkMaxOperations sets BulkMaxOperations field to given value.

### HasBulkMaxOperations

`func (o *AddHttpServletExtensionRequest) HasBulkMaxOperations() bool`

HasBulkMaxOperations returns a boolean if a field has been set.

### GetBulkMaxPayloadSize

`func (o *AddHttpServletExtensionRequest) GetBulkMaxPayloadSize() string`

GetBulkMaxPayloadSize returns the BulkMaxPayloadSize field if non-nil, zero value otherwise.

### GetBulkMaxPayloadSizeOk

`func (o *AddHttpServletExtensionRequest) GetBulkMaxPayloadSizeOk() (*string, bool)`

GetBulkMaxPayloadSizeOk returns a tuple with the BulkMaxPayloadSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkMaxPayloadSize

`func (o *AddHttpServletExtensionRequest) SetBulkMaxPayloadSize(v string)`

SetBulkMaxPayloadSize sets BulkMaxPayloadSize field to given value.

### HasBulkMaxPayloadSize

`func (o *AddHttpServletExtensionRequest) HasBulkMaxPayloadSize() bool`

HasBulkMaxPayloadSize returns a boolean if a field has been set.

### GetBulkMaxConcurrentRequests

`func (o *AddHttpServletExtensionRequest) GetBulkMaxConcurrentRequests() int32`

GetBulkMaxConcurrentRequests returns the BulkMaxConcurrentRequests field if non-nil, zero value otherwise.

### GetBulkMaxConcurrentRequestsOk

`func (o *AddHttpServletExtensionRequest) GetBulkMaxConcurrentRequestsOk() (*int32, bool)`

GetBulkMaxConcurrentRequestsOk returns a tuple with the BulkMaxConcurrentRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkMaxConcurrentRequests

`func (o *AddHttpServletExtensionRequest) SetBulkMaxConcurrentRequests(v int32)`

SetBulkMaxConcurrentRequests sets BulkMaxConcurrentRequests field to given value.

### HasBulkMaxConcurrentRequests

`func (o *AddHttpServletExtensionRequest) HasBulkMaxConcurrentRequests() bool`

HasBulkMaxConcurrentRequests returns a boolean if a field has been set.

### GetDebugEnabled

`func (o *AddHttpServletExtensionRequest) GetDebugEnabled() bool`

GetDebugEnabled returns the DebugEnabled field if non-nil, zero value otherwise.

### GetDebugEnabledOk

`func (o *AddHttpServletExtensionRequest) GetDebugEnabledOk() (*bool, bool)`

GetDebugEnabledOk returns a tuple with the DebugEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugEnabled

`func (o *AddHttpServletExtensionRequest) SetDebugEnabled(v bool)`

SetDebugEnabled sets DebugEnabled field to given value.

### HasDebugEnabled

`func (o *AddHttpServletExtensionRequest) HasDebugEnabled() bool`

HasDebugEnabled returns a boolean if a field has been set.

### GetDebugLevel

`func (o *AddHttpServletExtensionRequest) GetDebugLevel() EnumhttpServletExtensionDebugLevelProp`

GetDebugLevel returns the DebugLevel field if non-nil, zero value otherwise.

### GetDebugLevelOk

`func (o *AddHttpServletExtensionRequest) GetDebugLevelOk() (*EnumhttpServletExtensionDebugLevelProp, bool)`

GetDebugLevelOk returns a tuple with the DebugLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugLevel

`func (o *AddHttpServletExtensionRequest) SetDebugLevel(v EnumhttpServletExtensionDebugLevelProp)`

SetDebugLevel sets DebugLevel field to given value.

### HasDebugLevel

`func (o *AddHttpServletExtensionRequest) HasDebugLevel() bool`

HasDebugLevel returns a boolean if a field has been set.

### GetDebugType

`func (o *AddHttpServletExtensionRequest) GetDebugType() []EnumhttpServletExtensionDebugTypeProp`

GetDebugType returns the DebugType field if non-nil, zero value otherwise.

### GetDebugTypeOk

`func (o *AddHttpServletExtensionRequest) GetDebugTypeOk() (*[]EnumhttpServletExtensionDebugTypeProp, bool)`

GetDebugTypeOk returns a tuple with the DebugType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugType

`func (o *AddHttpServletExtensionRequest) SetDebugType(v []EnumhttpServletExtensionDebugTypeProp)`

SetDebugType sets DebugType field to given value.

### HasDebugType

`func (o *AddHttpServletExtensionRequest) HasDebugType() bool`

HasDebugType returns a boolean if a field has been set.

### GetIncludeStackTrace

`func (o *AddHttpServletExtensionRequest) GetIncludeStackTrace() bool`

GetIncludeStackTrace returns the IncludeStackTrace field if non-nil, zero value otherwise.

### GetIncludeStackTraceOk

`func (o *AddHttpServletExtensionRequest) GetIncludeStackTraceOk() (*bool, bool)`

GetIncludeStackTraceOk returns a tuple with the IncludeStackTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeStackTrace

`func (o *AddHttpServletExtensionRequest) SetIncludeStackTrace(v bool)`

SetIncludeStackTrace sets IncludeStackTrace field to given value.

### HasIncludeStackTrace

`func (o *AddHttpServletExtensionRequest) HasIncludeStackTrace() bool`

HasIncludeStackTrace returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddHttpServletExtensionRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddHttpServletExtensionRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddHttpServletExtensionRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddHttpServletExtensionRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddHttpServletExtensionRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddHttpServletExtensionRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddHttpServletExtensionRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


