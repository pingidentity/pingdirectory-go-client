# AddHttpServletExtension200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the HTTP Servlet Extension | 
**Schemas** | [**[]EnumthirdPartyHttpServletExtensionSchemaUrn**](EnumthirdPartyHttpServletExtensionSchemaUrn.md) |  | 
**Server** | Pointer to **string** | Specifies the PingFederate server to be configured. | [optional] 
**Description** | Pointer to **string** | A description for this HTTP Servlet Extension | [optional] 
**CrossOriginPolicy** | Pointer to **string** | The cross-origin request policy to use for the HTTP Servlet Extension. | [optional] 
**ResponseHeader** | Pointer to **[]string** | Specifies HTTP header fields and values added to response headers for all requests. | [optional] 
**CorrelationIDResponseHeader** | Pointer to **string** | Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \&quot;Correlation-Id\&quot;, \&quot;X-Amzn-Trace-Id\&quot;, and \&quot;X-Request-Id\&quot;. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**BaseContextPath** | **string** | Specifies the base context path that should be used by HTTP clients to reference content. The value must start with a forward slash and must represent a valid HTTP context path. | 
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
**OAuthTokenHandler** | Pointer to **string** | Specifies the OAuth Token Handler implementation that should be used to validate OAuth 2.0 bearer tokens when they are included in a SCIM request. | [optional] 
**BasicAuthEnabled** | Pointer to **bool** | Enables HTTP Basic authentication, using a username and password. | [optional] 
**IdentityMapper** | Pointer to **string** | The identity mapper that will be used to identify the entry with which a username is associated. | [optional] 
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
**DebugEnabled** | Pointer to **bool** | Enables debug logging of the SCIM SDK. Debug messages will be forwarded to the Directory Server debug logger with the scope of com.unboundid.directory.server.extensions.scim.SCIMHTTPServletExtension. | [optional] 
**DebugLevel** | [**EnumhttpServletExtensionDebugLevelProp**](EnumhttpServletExtensionDebugLevelProp.md) |  | 
**DebugType** | [**[]EnumhttpServletExtensionDebugTypeProp**](EnumhttpServletExtensionDebugTypeProp.md) |  | 
**IncludeStackTrace** | **bool** | Indicates whether a stack trace of the thread which called the debug method should be included in debug log messages. | 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted HTTP Servlet Extension. | 
**ScriptArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Scripted HTTP Servlet Extension. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
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
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party HTTP Servlet Extension. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party HTTP Servlet Extension. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddHttpServletExtension200Response

`func NewAddHttpServletExtension200Response(id string, schemas []EnumthirdPartyHttpServletExtensionSchemaUrn, baseContextPath string, availableStatusCode int64, degradedStatusCode int64, unavailableStatusCode int64, temporaryDirectory string, temporaryDirectoryPermissions string, debugLevel EnumhttpServletExtensionDebugLevelProp, debugType []EnumhttpServletExtensionDebugTypeProp, includeStackTrace bool, scriptClass string, documentRootDirectory string, extensionClass string, ) *AddHttpServletExtension200Response`

NewAddHttpServletExtension200Response instantiates a new AddHttpServletExtension200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddHttpServletExtension200ResponseWithDefaults

`func NewAddHttpServletExtension200ResponseWithDefaults() *AddHttpServletExtension200Response`

NewAddHttpServletExtension200ResponseWithDefaults instantiates a new AddHttpServletExtension200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddHttpServletExtension200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddHttpServletExtension200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddHttpServletExtension200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddHttpServletExtension200Response) GetSchemas() []EnumthirdPartyHttpServletExtensionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddHttpServletExtension200Response) GetSchemasOk() (*[]EnumthirdPartyHttpServletExtensionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddHttpServletExtension200Response) SetSchemas(v []EnumthirdPartyHttpServletExtensionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetServer

`func (o *AddHttpServletExtension200Response) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *AddHttpServletExtension200Response) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *AddHttpServletExtension200Response) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *AddHttpServletExtension200Response) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetDescription

`func (o *AddHttpServletExtension200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddHttpServletExtension200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddHttpServletExtension200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddHttpServletExtension200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCrossOriginPolicy

`func (o *AddHttpServletExtension200Response) GetCrossOriginPolicy() string`

GetCrossOriginPolicy returns the CrossOriginPolicy field if non-nil, zero value otherwise.

### GetCrossOriginPolicyOk

`func (o *AddHttpServletExtension200Response) GetCrossOriginPolicyOk() (*string, bool)`

GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossOriginPolicy

`func (o *AddHttpServletExtension200Response) SetCrossOriginPolicy(v string)`

SetCrossOriginPolicy sets CrossOriginPolicy field to given value.

### HasCrossOriginPolicy

`func (o *AddHttpServletExtension200Response) HasCrossOriginPolicy() bool`

HasCrossOriginPolicy returns a boolean if a field has been set.

### GetResponseHeader

`func (o *AddHttpServletExtension200Response) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *AddHttpServletExtension200Response) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *AddHttpServletExtension200Response) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *AddHttpServletExtension200Response) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetCorrelationIDResponseHeader

`func (o *AddHttpServletExtension200Response) GetCorrelationIDResponseHeader() string`

GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field if non-nil, zero value otherwise.

### GetCorrelationIDResponseHeaderOk

`func (o *AddHttpServletExtension200Response) GetCorrelationIDResponseHeaderOk() (*string, bool)`

GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDResponseHeader

`func (o *AddHttpServletExtension200Response) SetCorrelationIDResponseHeader(v string)`

SetCorrelationIDResponseHeader sets CorrelationIDResponseHeader field to given value.

### HasCorrelationIDResponseHeader

`func (o *AddHttpServletExtension200Response) HasCorrelationIDResponseHeader() bool`

HasCorrelationIDResponseHeader returns a boolean if a field has been set.

### GetMeta

`func (o *AddHttpServletExtension200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddHttpServletExtension200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddHttpServletExtension200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddHttpServletExtension200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AddHttpServletExtension200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AddHttpServletExtension200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AddHttpServletExtension200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AddHttpServletExtension200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetBaseContextPath

`func (o *AddHttpServletExtension200Response) GetBaseContextPath() string`

GetBaseContextPath returns the BaseContextPath field if non-nil, zero value otherwise.

### GetBaseContextPathOk

`func (o *AddHttpServletExtension200Response) GetBaseContextPathOk() (*string, bool)`

GetBaseContextPathOk returns a tuple with the BaseContextPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseContextPath

`func (o *AddHttpServletExtension200Response) SetBaseContextPath(v string)`

SetBaseContextPath sets BaseContextPath field to given value.


### GetAvailableStatusCode

`func (o *AddHttpServletExtension200Response) GetAvailableStatusCode() int64`

GetAvailableStatusCode returns the AvailableStatusCode field if non-nil, zero value otherwise.

### GetAvailableStatusCodeOk

`func (o *AddHttpServletExtension200Response) GetAvailableStatusCodeOk() (*int64, bool)`

GetAvailableStatusCodeOk returns a tuple with the AvailableStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableStatusCode

`func (o *AddHttpServletExtension200Response) SetAvailableStatusCode(v int64)`

SetAvailableStatusCode sets AvailableStatusCode field to given value.


### GetDegradedStatusCode

`func (o *AddHttpServletExtension200Response) GetDegradedStatusCode() int64`

GetDegradedStatusCode returns the DegradedStatusCode field if non-nil, zero value otherwise.

### GetDegradedStatusCodeOk

`func (o *AddHttpServletExtension200Response) GetDegradedStatusCodeOk() (*int64, bool)`

GetDegradedStatusCodeOk returns a tuple with the DegradedStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegradedStatusCode

`func (o *AddHttpServletExtension200Response) SetDegradedStatusCode(v int64)`

SetDegradedStatusCode sets DegradedStatusCode field to given value.


### GetUnavailableStatusCode

`func (o *AddHttpServletExtension200Response) GetUnavailableStatusCode() int64`

GetUnavailableStatusCode returns the UnavailableStatusCode field if non-nil, zero value otherwise.

### GetUnavailableStatusCodeOk

`func (o *AddHttpServletExtension200Response) GetUnavailableStatusCodeOk() (*int64, bool)`

GetUnavailableStatusCodeOk returns a tuple with the UnavailableStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailableStatusCode

`func (o *AddHttpServletExtension200Response) SetUnavailableStatusCode(v int64)`

SetUnavailableStatusCode sets UnavailableStatusCode field to given value.


### GetOverrideStatusCode

`func (o *AddHttpServletExtension200Response) GetOverrideStatusCode() int64`

GetOverrideStatusCode returns the OverrideStatusCode field if non-nil, zero value otherwise.

### GetOverrideStatusCodeOk

`func (o *AddHttpServletExtension200Response) GetOverrideStatusCodeOk() (*int64, bool)`

GetOverrideStatusCodeOk returns a tuple with the OverrideStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideStatusCode

`func (o *AddHttpServletExtension200Response) SetOverrideStatusCode(v int64)`

SetOverrideStatusCode sets OverrideStatusCode field to given value.

### HasOverrideStatusCode

`func (o *AddHttpServletExtension200Response) HasOverrideStatusCode() bool`

HasOverrideStatusCode returns a boolean if a field has been set.

### GetIncludeResponseBody

`func (o *AddHttpServletExtension200Response) GetIncludeResponseBody() bool`

GetIncludeResponseBody returns the IncludeResponseBody field if non-nil, zero value otherwise.

### GetIncludeResponseBodyOk

`func (o *AddHttpServletExtension200Response) GetIncludeResponseBodyOk() (*bool, bool)`

GetIncludeResponseBodyOk returns a tuple with the IncludeResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeResponseBody

`func (o *AddHttpServletExtension200Response) SetIncludeResponseBody(v bool)`

SetIncludeResponseBody sets IncludeResponseBody field to given value.

### HasIncludeResponseBody

`func (o *AddHttpServletExtension200Response) HasIncludeResponseBody() bool`

HasIncludeResponseBody returns a boolean if a field has been set.

### GetAdditionalResponseContents

`func (o *AddHttpServletExtension200Response) GetAdditionalResponseContents() string`

GetAdditionalResponseContents returns the AdditionalResponseContents field if non-nil, zero value otherwise.

### GetAdditionalResponseContentsOk

`func (o *AddHttpServletExtension200Response) GetAdditionalResponseContentsOk() (*string, bool)`

GetAdditionalResponseContentsOk returns a tuple with the AdditionalResponseContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalResponseContents

`func (o *AddHttpServletExtension200Response) SetAdditionalResponseContents(v string)`

SetAdditionalResponseContents sets AdditionalResponseContents field to given value.

### HasAdditionalResponseContents

`func (o *AddHttpServletExtension200Response) HasAdditionalResponseContents() bool`

HasAdditionalResponseContents returns a boolean if a field has been set.

### GetIncludeInstanceNameLabel

`func (o *AddHttpServletExtension200Response) GetIncludeInstanceNameLabel() bool`

GetIncludeInstanceNameLabel returns the IncludeInstanceNameLabel field if non-nil, zero value otherwise.

### GetIncludeInstanceNameLabelOk

`func (o *AddHttpServletExtension200Response) GetIncludeInstanceNameLabelOk() (*bool, bool)`

GetIncludeInstanceNameLabelOk returns a tuple with the IncludeInstanceNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInstanceNameLabel

`func (o *AddHttpServletExtension200Response) SetIncludeInstanceNameLabel(v bool)`

SetIncludeInstanceNameLabel sets IncludeInstanceNameLabel field to given value.

### HasIncludeInstanceNameLabel

`func (o *AddHttpServletExtension200Response) HasIncludeInstanceNameLabel() bool`

HasIncludeInstanceNameLabel returns a boolean if a field has been set.

### GetIncludeProductNameLabel

`func (o *AddHttpServletExtension200Response) GetIncludeProductNameLabel() bool`

GetIncludeProductNameLabel returns the IncludeProductNameLabel field if non-nil, zero value otherwise.

### GetIncludeProductNameLabelOk

`func (o *AddHttpServletExtension200Response) GetIncludeProductNameLabelOk() (*bool, bool)`

GetIncludeProductNameLabelOk returns a tuple with the IncludeProductNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeProductNameLabel

`func (o *AddHttpServletExtension200Response) SetIncludeProductNameLabel(v bool)`

SetIncludeProductNameLabel sets IncludeProductNameLabel field to given value.

### HasIncludeProductNameLabel

`func (o *AddHttpServletExtension200Response) HasIncludeProductNameLabel() bool`

HasIncludeProductNameLabel returns a boolean if a field has been set.

### GetIncludeLocationNameLabel

`func (o *AddHttpServletExtension200Response) GetIncludeLocationNameLabel() bool`

GetIncludeLocationNameLabel returns the IncludeLocationNameLabel field if non-nil, zero value otherwise.

### GetIncludeLocationNameLabelOk

`func (o *AddHttpServletExtension200Response) GetIncludeLocationNameLabelOk() (*bool, bool)`

GetIncludeLocationNameLabelOk returns a tuple with the IncludeLocationNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLocationNameLabel

`func (o *AddHttpServletExtension200Response) SetIncludeLocationNameLabel(v bool)`

SetIncludeLocationNameLabel sets IncludeLocationNameLabel field to given value.

### HasIncludeLocationNameLabel

`func (o *AddHttpServletExtension200Response) HasIncludeLocationNameLabel() bool`

HasIncludeLocationNameLabel returns a boolean if a field has been set.

### GetAlwaysIncludeMonitorEntryNameLabel

`func (o *AddHttpServletExtension200Response) GetAlwaysIncludeMonitorEntryNameLabel() bool`

GetAlwaysIncludeMonitorEntryNameLabel returns the AlwaysIncludeMonitorEntryNameLabel field if non-nil, zero value otherwise.

### GetAlwaysIncludeMonitorEntryNameLabelOk

`func (o *AddHttpServletExtension200Response) GetAlwaysIncludeMonitorEntryNameLabelOk() (*bool, bool)`

GetAlwaysIncludeMonitorEntryNameLabelOk returns a tuple with the AlwaysIncludeMonitorEntryNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysIncludeMonitorEntryNameLabel

`func (o *AddHttpServletExtension200Response) SetAlwaysIncludeMonitorEntryNameLabel(v bool)`

SetAlwaysIncludeMonitorEntryNameLabel sets AlwaysIncludeMonitorEntryNameLabel field to given value.

### HasAlwaysIncludeMonitorEntryNameLabel

`func (o *AddHttpServletExtension200Response) HasAlwaysIncludeMonitorEntryNameLabel() bool`

HasAlwaysIncludeMonitorEntryNameLabel returns a boolean if a field has been set.

### GetIncludeMonitorObjectClassNameLabel

`func (o *AddHttpServletExtension200Response) GetIncludeMonitorObjectClassNameLabel() bool`

GetIncludeMonitorObjectClassNameLabel returns the IncludeMonitorObjectClassNameLabel field if non-nil, zero value otherwise.

### GetIncludeMonitorObjectClassNameLabelOk

`func (o *AddHttpServletExtension200Response) GetIncludeMonitorObjectClassNameLabelOk() (*bool, bool)`

GetIncludeMonitorObjectClassNameLabelOk returns a tuple with the IncludeMonitorObjectClassNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMonitorObjectClassNameLabel

`func (o *AddHttpServletExtension200Response) SetIncludeMonitorObjectClassNameLabel(v bool)`

SetIncludeMonitorObjectClassNameLabel sets IncludeMonitorObjectClassNameLabel field to given value.

### HasIncludeMonitorObjectClassNameLabel

`func (o *AddHttpServletExtension200Response) HasIncludeMonitorObjectClassNameLabel() bool`

HasIncludeMonitorObjectClassNameLabel returns a boolean if a field has been set.

### GetIncludeMonitorAttributeNameLabel

`func (o *AddHttpServletExtension200Response) GetIncludeMonitorAttributeNameLabel() bool`

GetIncludeMonitorAttributeNameLabel returns the IncludeMonitorAttributeNameLabel field if non-nil, zero value otherwise.

### GetIncludeMonitorAttributeNameLabelOk

`func (o *AddHttpServletExtension200Response) GetIncludeMonitorAttributeNameLabelOk() (*bool, bool)`

GetIncludeMonitorAttributeNameLabelOk returns a tuple with the IncludeMonitorAttributeNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMonitorAttributeNameLabel

`func (o *AddHttpServletExtension200Response) SetIncludeMonitorAttributeNameLabel(v bool)`

SetIncludeMonitorAttributeNameLabel sets IncludeMonitorAttributeNameLabel field to given value.

### HasIncludeMonitorAttributeNameLabel

`func (o *AddHttpServletExtension200Response) HasIncludeMonitorAttributeNameLabel() bool`

HasIncludeMonitorAttributeNameLabel returns a boolean if a field has been set.

### GetLabelNameValuePair

`func (o *AddHttpServletExtension200Response) GetLabelNameValuePair() []string`

GetLabelNameValuePair returns the LabelNameValuePair field if non-nil, zero value otherwise.

### GetLabelNameValuePairOk

`func (o *AddHttpServletExtension200Response) GetLabelNameValuePairOk() (*[]string, bool)`

GetLabelNameValuePairOk returns a tuple with the LabelNameValuePair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelNameValuePair

`func (o *AddHttpServletExtension200Response) SetLabelNameValuePair(v []string)`

SetLabelNameValuePair sets LabelNameValuePair field to given value.

### HasLabelNameValuePair

`func (o *AddHttpServletExtension200Response) HasLabelNameValuePair() bool`

HasLabelNameValuePair returns a boolean if a field has been set.

### GetOAuthTokenHandler

`func (o *AddHttpServletExtension200Response) GetOAuthTokenHandler() string`

GetOAuthTokenHandler returns the OAuthTokenHandler field if non-nil, zero value otherwise.

### GetOAuthTokenHandlerOk

`func (o *AddHttpServletExtension200Response) GetOAuthTokenHandlerOk() (*string, bool)`

GetOAuthTokenHandlerOk returns a tuple with the OAuthTokenHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthTokenHandler

`func (o *AddHttpServletExtension200Response) SetOAuthTokenHandler(v string)`

SetOAuthTokenHandler sets OAuthTokenHandler field to given value.

### HasOAuthTokenHandler

`func (o *AddHttpServletExtension200Response) HasOAuthTokenHandler() bool`

HasOAuthTokenHandler returns a boolean if a field has been set.

### GetBasicAuthEnabled

`func (o *AddHttpServletExtension200Response) GetBasicAuthEnabled() bool`

GetBasicAuthEnabled returns the BasicAuthEnabled field if non-nil, zero value otherwise.

### GetBasicAuthEnabledOk

`func (o *AddHttpServletExtension200Response) GetBasicAuthEnabledOk() (*bool, bool)`

GetBasicAuthEnabledOk returns a tuple with the BasicAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthEnabled

`func (o *AddHttpServletExtension200Response) SetBasicAuthEnabled(v bool)`

SetBasicAuthEnabled sets BasicAuthEnabled field to given value.

### HasBasicAuthEnabled

`func (o *AddHttpServletExtension200Response) HasBasicAuthEnabled() bool`

HasBasicAuthEnabled returns a boolean if a field has been set.

### GetIdentityMapper

`func (o *AddHttpServletExtension200Response) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *AddHttpServletExtension200Response) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *AddHttpServletExtension200Response) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.

### HasIdentityMapper

`func (o *AddHttpServletExtension200Response) HasIdentityMapper() bool`

HasIdentityMapper returns a boolean if a field has been set.

### GetResourceMappingFile

`func (o *AddHttpServletExtension200Response) GetResourceMappingFile() string`

GetResourceMappingFile returns the ResourceMappingFile field if non-nil, zero value otherwise.

### GetResourceMappingFileOk

`func (o *AddHttpServletExtension200Response) GetResourceMappingFileOk() (*string, bool)`

GetResourceMappingFileOk returns a tuple with the ResourceMappingFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceMappingFile

`func (o *AddHttpServletExtension200Response) SetResourceMappingFile(v string)`

SetResourceMappingFile sets ResourceMappingFile field to given value.

### HasResourceMappingFile

`func (o *AddHttpServletExtension200Response) HasResourceMappingFile() bool`

HasResourceMappingFile returns a boolean if a field has been set.

### GetIncludeLDAPObjectclass

`func (o *AddHttpServletExtension200Response) GetIncludeLDAPObjectclass() []string`

GetIncludeLDAPObjectclass returns the IncludeLDAPObjectclass field if non-nil, zero value otherwise.

### GetIncludeLDAPObjectclassOk

`func (o *AddHttpServletExtension200Response) GetIncludeLDAPObjectclassOk() (*[]string, bool)`

GetIncludeLDAPObjectclassOk returns a tuple with the IncludeLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLDAPObjectclass

`func (o *AddHttpServletExtension200Response) SetIncludeLDAPObjectclass(v []string)`

SetIncludeLDAPObjectclass sets IncludeLDAPObjectclass field to given value.

### HasIncludeLDAPObjectclass

`func (o *AddHttpServletExtension200Response) HasIncludeLDAPObjectclass() bool`

HasIncludeLDAPObjectclass returns a boolean if a field has been set.

### GetExcludeLDAPObjectclass

`func (o *AddHttpServletExtension200Response) GetExcludeLDAPObjectclass() []string`

GetExcludeLDAPObjectclass returns the ExcludeLDAPObjectclass field if non-nil, zero value otherwise.

### GetExcludeLDAPObjectclassOk

`func (o *AddHttpServletExtension200Response) GetExcludeLDAPObjectclassOk() (*[]string, bool)`

GetExcludeLDAPObjectclassOk returns a tuple with the ExcludeLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeLDAPObjectclass

`func (o *AddHttpServletExtension200Response) SetExcludeLDAPObjectclass(v []string)`

SetExcludeLDAPObjectclass sets ExcludeLDAPObjectclass field to given value.

### HasExcludeLDAPObjectclass

`func (o *AddHttpServletExtension200Response) HasExcludeLDAPObjectclass() bool`

HasExcludeLDAPObjectclass returns a boolean if a field has been set.

### GetIncludeLDAPBaseDN

`func (o *AddHttpServletExtension200Response) GetIncludeLDAPBaseDN() []string`

GetIncludeLDAPBaseDN returns the IncludeLDAPBaseDN field if non-nil, zero value otherwise.

### GetIncludeLDAPBaseDNOk

`func (o *AddHttpServletExtension200Response) GetIncludeLDAPBaseDNOk() (*[]string, bool)`

GetIncludeLDAPBaseDNOk returns a tuple with the IncludeLDAPBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLDAPBaseDN

`func (o *AddHttpServletExtension200Response) SetIncludeLDAPBaseDN(v []string)`

SetIncludeLDAPBaseDN sets IncludeLDAPBaseDN field to given value.

### HasIncludeLDAPBaseDN

`func (o *AddHttpServletExtension200Response) HasIncludeLDAPBaseDN() bool`

HasIncludeLDAPBaseDN returns a boolean if a field has been set.

### GetExcludeLDAPBaseDN

`func (o *AddHttpServletExtension200Response) GetExcludeLDAPBaseDN() []string`

GetExcludeLDAPBaseDN returns the ExcludeLDAPBaseDN field if non-nil, zero value otherwise.

### GetExcludeLDAPBaseDNOk

`func (o *AddHttpServletExtension200Response) GetExcludeLDAPBaseDNOk() (*[]string, bool)`

GetExcludeLDAPBaseDNOk returns a tuple with the ExcludeLDAPBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeLDAPBaseDN

`func (o *AddHttpServletExtension200Response) SetExcludeLDAPBaseDN(v []string)`

SetExcludeLDAPBaseDN sets ExcludeLDAPBaseDN field to given value.

### HasExcludeLDAPBaseDN

`func (o *AddHttpServletExtension200Response) HasExcludeLDAPBaseDN() bool`

HasExcludeLDAPBaseDN returns a boolean if a field has been set.

### GetEntityTagLDAPAttribute

`func (o *AddHttpServletExtension200Response) GetEntityTagLDAPAttribute() string`

GetEntityTagLDAPAttribute returns the EntityTagLDAPAttribute field if non-nil, zero value otherwise.

### GetEntityTagLDAPAttributeOk

`func (o *AddHttpServletExtension200Response) GetEntityTagLDAPAttributeOk() (*string, bool)`

GetEntityTagLDAPAttributeOk returns a tuple with the EntityTagLDAPAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTagLDAPAttribute

`func (o *AddHttpServletExtension200Response) SetEntityTagLDAPAttribute(v string)`

SetEntityTagLDAPAttribute sets EntityTagLDAPAttribute field to given value.

### HasEntityTagLDAPAttribute

`func (o *AddHttpServletExtension200Response) HasEntityTagLDAPAttribute() bool`

HasEntityTagLDAPAttribute returns a boolean if a field has been set.

### GetTemporaryDirectory

`func (o *AddHttpServletExtension200Response) GetTemporaryDirectory() string`

GetTemporaryDirectory returns the TemporaryDirectory field if non-nil, zero value otherwise.

### GetTemporaryDirectoryOk

`func (o *AddHttpServletExtension200Response) GetTemporaryDirectoryOk() (*string, bool)`

GetTemporaryDirectoryOk returns a tuple with the TemporaryDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryDirectory

`func (o *AddHttpServletExtension200Response) SetTemporaryDirectory(v string)`

SetTemporaryDirectory sets TemporaryDirectory field to given value.


### GetTemporaryDirectoryPermissions

`func (o *AddHttpServletExtension200Response) GetTemporaryDirectoryPermissions() string`

GetTemporaryDirectoryPermissions returns the TemporaryDirectoryPermissions field if non-nil, zero value otherwise.

### GetTemporaryDirectoryPermissionsOk

`func (o *AddHttpServletExtension200Response) GetTemporaryDirectoryPermissionsOk() (*string, bool)`

GetTemporaryDirectoryPermissionsOk returns a tuple with the TemporaryDirectoryPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryDirectoryPermissions

`func (o *AddHttpServletExtension200Response) SetTemporaryDirectoryPermissions(v string)`

SetTemporaryDirectoryPermissions sets TemporaryDirectoryPermissions field to given value.


### GetMaxResults

`func (o *AddHttpServletExtension200Response) GetMaxResults() int64`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *AddHttpServletExtension200Response) GetMaxResultsOk() (*int64, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *AddHttpServletExtension200Response) SetMaxResults(v int64)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *AddHttpServletExtension200Response) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetBulkMaxOperations

`func (o *AddHttpServletExtension200Response) GetBulkMaxOperations() int64`

GetBulkMaxOperations returns the BulkMaxOperations field if non-nil, zero value otherwise.

### GetBulkMaxOperationsOk

`func (o *AddHttpServletExtension200Response) GetBulkMaxOperationsOk() (*int64, bool)`

GetBulkMaxOperationsOk returns a tuple with the BulkMaxOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkMaxOperations

`func (o *AddHttpServletExtension200Response) SetBulkMaxOperations(v int64)`

SetBulkMaxOperations sets BulkMaxOperations field to given value.

### HasBulkMaxOperations

`func (o *AddHttpServletExtension200Response) HasBulkMaxOperations() bool`

HasBulkMaxOperations returns a boolean if a field has been set.

### GetBulkMaxPayloadSize

`func (o *AddHttpServletExtension200Response) GetBulkMaxPayloadSize() string`

GetBulkMaxPayloadSize returns the BulkMaxPayloadSize field if non-nil, zero value otherwise.

### GetBulkMaxPayloadSizeOk

`func (o *AddHttpServletExtension200Response) GetBulkMaxPayloadSizeOk() (*string, bool)`

GetBulkMaxPayloadSizeOk returns a tuple with the BulkMaxPayloadSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkMaxPayloadSize

`func (o *AddHttpServletExtension200Response) SetBulkMaxPayloadSize(v string)`

SetBulkMaxPayloadSize sets BulkMaxPayloadSize field to given value.

### HasBulkMaxPayloadSize

`func (o *AddHttpServletExtension200Response) HasBulkMaxPayloadSize() bool`

HasBulkMaxPayloadSize returns a boolean if a field has been set.

### GetBulkMaxConcurrentRequests

`func (o *AddHttpServletExtension200Response) GetBulkMaxConcurrentRequests() int64`

GetBulkMaxConcurrentRequests returns the BulkMaxConcurrentRequests field if non-nil, zero value otherwise.

### GetBulkMaxConcurrentRequestsOk

`func (o *AddHttpServletExtension200Response) GetBulkMaxConcurrentRequestsOk() (*int64, bool)`

GetBulkMaxConcurrentRequestsOk returns a tuple with the BulkMaxConcurrentRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkMaxConcurrentRequests

`func (o *AddHttpServletExtension200Response) SetBulkMaxConcurrentRequests(v int64)`

SetBulkMaxConcurrentRequests sets BulkMaxConcurrentRequests field to given value.

### HasBulkMaxConcurrentRequests

`func (o *AddHttpServletExtension200Response) HasBulkMaxConcurrentRequests() bool`

HasBulkMaxConcurrentRequests returns a boolean if a field has been set.

### GetDebugEnabled

`func (o *AddHttpServletExtension200Response) GetDebugEnabled() bool`

GetDebugEnabled returns the DebugEnabled field if non-nil, zero value otherwise.

### GetDebugEnabledOk

`func (o *AddHttpServletExtension200Response) GetDebugEnabledOk() (*bool, bool)`

GetDebugEnabledOk returns a tuple with the DebugEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugEnabled

`func (o *AddHttpServletExtension200Response) SetDebugEnabled(v bool)`

SetDebugEnabled sets DebugEnabled field to given value.

### HasDebugEnabled

`func (o *AddHttpServletExtension200Response) HasDebugEnabled() bool`

HasDebugEnabled returns a boolean if a field has been set.

### GetDebugLevel

`func (o *AddHttpServletExtension200Response) GetDebugLevel() EnumhttpServletExtensionDebugLevelProp`

GetDebugLevel returns the DebugLevel field if non-nil, zero value otherwise.

### GetDebugLevelOk

`func (o *AddHttpServletExtension200Response) GetDebugLevelOk() (*EnumhttpServletExtensionDebugLevelProp, bool)`

GetDebugLevelOk returns a tuple with the DebugLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugLevel

`func (o *AddHttpServletExtension200Response) SetDebugLevel(v EnumhttpServletExtensionDebugLevelProp)`

SetDebugLevel sets DebugLevel field to given value.


### GetDebugType

`func (o *AddHttpServletExtension200Response) GetDebugType() []EnumhttpServletExtensionDebugTypeProp`

GetDebugType returns the DebugType field if non-nil, zero value otherwise.

### GetDebugTypeOk

`func (o *AddHttpServletExtension200Response) GetDebugTypeOk() (*[]EnumhttpServletExtensionDebugTypeProp, bool)`

GetDebugTypeOk returns a tuple with the DebugType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugType

`func (o *AddHttpServletExtension200Response) SetDebugType(v []EnumhttpServletExtensionDebugTypeProp)`

SetDebugType sets DebugType field to given value.


### GetIncludeStackTrace

`func (o *AddHttpServletExtension200Response) GetIncludeStackTrace() bool`

GetIncludeStackTrace returns the IncludeStackTrace field if non-nil, zero value otherwise.

### GetIncludeStackTraceOk

`func (o *AddHttpServletExtension200Response) GetIncludeStackTraceOk() (*bool, bool)`

GetIncludeStackTraceOk returns a tuple with the IncludeStackTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeStackTrace

`func (o *AddHttpServletExtension200Response) SetIncludeStackTrace(v bool)`

SetIncludeStackTrace sets IncludeStackTrace field to given value.


### GetScriptClass

`func (o *AddHttpServletExtension200Response) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *AddHttpServletExtension200Response) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *AddHttpServletExtension200Response) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *AddHttpServletExtension200Response) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *AddHttpServletExtension200Response) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *AddHttpServletExtension200Response) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *AddHttpServletExtension200Response) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetDocumentRootDirectory

`func (o *AddHttpServletExtension200Response) GetDocumentRootDirectory() string`

GetDocumentRootDirectory returns the DocumentRootDirectory field if non-nil, zero value otherwise.

### GetDocumentRootDirectoryOk

`func (o *AddHttpServletExtension200Response) GetDocumentRootDirectoryOk() (*string, bool)`

GetDocumentRootDirectoryOk returns a tuple with the DocumentRootDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentRootDirectory

`func (o *AddHttpServletExtension200Response) SetDocumentRootDirectory(v string)`

SetDocumentRootDirectory sets DocumentRootDirectory field to given value.


### GetEnableDirectoryIndexing

`func (o *AddHttpServletExtension200Response) GetEnableDirectoryIndexing() bool`

GetEnableDirectoryIndexing returns the EnableDirectoryIndexing field if non-nil, zero value otherwise.

### GetEnableDirectoryIndexingOk

`func (o *AddHttpServletExtension200Response) GetEnableDirectoryIndexingOk() (*bool, bool)`

GetEnableDirectoryIndexingOk returns a tuple with the EnableDirectoryIndexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDirectoryIndexing

`func (o *AddHttpServletExtension200Response) SetEnableDirectoryIndexing(v bool)`

SetEnableDirectoryIndexing sets EnableDirectoryIndexing field to given value.

### HasEnableDirectoryIndexing

`func (o *AddHttpServletExtension200Response) HasEnableDirectoryIndexing() bool`

HasEnableDirectoryIndexing returns a boolean if a field has been set.

### GetIndexFile

`func (o *AddHttpServletExtension200Response) GetIndexFile() []string`

GetIndexFile returns the IndexFile field if non-nil, zero value otherwise.

### GetIndexFileOk

`func (o *AddHttpServletExtension200Response) GetIndexFileOk() (*[]string, bool)`

GetIndexFileOk returns a tuple with the IndexFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexFile

`func (o *AddHttpServletExtension200Response) SetIndexFile(v []string)`

SetIndexFile sets IndexFile field to given value.

### HasIndexFile

`func (o *AddHttpServletExtension200Response) HasIndexFile() bool`

HasIndexFile returns a boolean if a field has been set.

### GetMimeTypesFile

`func (o *AddHttpServletExtension200Response) GetMimeTypesFile() string`

GetMimeTypesFile returns the MimeTypesFile field if non-nil, zero value otherwise.

### GetMimeTypesFileOk

`func (o *AddHttpServletExtension200Response) GetMimeTypesFileOk() (*string, bool)`

GetMimeTypesFileOk returns a tuple with the MimeTypesFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeTypesFile

`func (o *AddHttpServletExtension200Response) SetMimeTypesFile(v string)`

SetMimeTypesFile sets MimeTypesFile field to given value.

### HasMimeTypesFile

`func (o *AddHttpServletExtension200Response) HasMimeTypesFile() bool`

HasMimeTypesFile returns a boolean if a field has been set.

### GetDefaultMIMEType

`func (o *AddHttpServletExtension200Response) GetDefaultMIMEType() string`

GetDefaultMIMEType returns the DefaultMIMEType field if non-nil, zero value otherwise.

### GetDefaultMIMETypeOk

`func (o *AddHttpServletExtension200Response) GetDefaultMIMETypeOk() (*string, bool)`

GetDefaultMIMETypeOk returns a tuple with the DefaultMIMEType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMIMEType

`func (o *AddHttpServletExtension200Response) SetDefaultMIMEType(v string)`

SetDefaultMIMEType sets DefaultMIMEType field to given value.

### HasDefaultMIMEType

`func (o *AddHttpServletExtension200Response) HasDefaultMIMEType() bool`

HasDefaultMIMEType returns a boolean if a field has been set.

### GetRequireAuthentication

`func (o *AddHttpServletExtension200Response) GetRequireAuthentication() bool`

GetRequireAuthentication returns the RequireAuthentication field if non-nil, zero value otherwise.

### GetRequireAuthenticationOk

`func (o *AddHttpServletExtension200Response) GetRequireAuthenticationOk() (*bool, bool)`

GetRequireAuthenticationOk returns a tuple with the RequireAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAuthentication

`func (o *AddHttpServletExtension200Response) SetRequireAuthentication(v bool)`

SetRequireAuthentication sets RequireAuthentication field to given value.

### HasRequireAuthentication

`func (o *AddHttpServletExtension200Response) HasRequireAuthentication() bool`

HasRequireAuthentication returns a boolean if a field has been set.

### GetAllowedAuthenticationType

`func (o *AddHttpServletExtension200Response) GetAllowedAuthenticationType() []EnumhttpServletExtensionAllowedAuthenticationTypeProp`

GetAllowedAuthenticationType returns the AllowedAuthenticationType field if non-nil, zero value otherwise.

### GetAllowedAuthenticationTypeOk

`func (o *AddHttpServletExtension200Response) GetAllowedAuthenticationTypeOk() (*[]EnumhttpServletExtensionAllowedAuthenticationTypeProp, bool)`

GetAllowedAuthenticationTypeOk returns a tuple with the AllowedAuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAuthenticationType

`func (o *AddHttpServletExtension200Response) SetAllowedAuthenticationType(v []EnumhttpServletExtensionAllowedAuthenticationTypeProp)`

SetAllowedAuthenticationType sets AllowedAuthenticationType field to given value.

### HasAllowedAuthenticationType

`func (o *AddHttpServletExtension200Response) HasAllowedAuthenticationType() bool`

HasAllowedAuthenticationType returns a boolean if a field has been set.

### GetAccessTokenValidator

`func (o *AddHttpServletExtension200Response) GetAccessTokenValidator() []string`

GetAccessTokenValidator returns the AccessTokenValidator field if non-nil, zero value otherwise.

### GetAccessTokenValidatorOk

`func (o *AddHttpServletExtension200Response) GetAccessTokenValidatorOk() (*[]string, bool)`

GetAccessTokenValidatorOk returns a tuple with the AccessTokenValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValidator

`func (o *AddHttpServletExtension200Response) SetAccessTokenValidator(v []string)`

SetAccessTokenValidator sets AccessTokenValidator field to given value.

### HasAccessTokenValidator

`func (o *AddHttpServletExtension200Response) HasAccessTokenValidator() bool`

HasAccessTokenValidator returns a boolean if a field has been set.

### GetIdTokenValidator

`func (o *AddHttpServletExtension200Response) GetIdTokenValidator() []string`

GetIdTokenValidator returns the IdTokenValidator field if non-nil, zero value otherwise.

### GetIdTokenValidatorOk

`func (o *AddHttpServletExtension200Response) GetIdTokenValidatorOk() (*[]string, bool)`

GetIdTokenValidatorOk returns a tuple with the IdTokenValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenValidator

`func (o *AddHttpServletExtension200Response) SetIdTokenValidator(v []string)`

SetIdTokenValidator sets IdTokenValidator field to given value.

### HasIdTokenValidator

`func (o *AddHttpServletExtension200Response) HasIdTokenValidator() bool`

HasIdTokenValidator returns a boolean if a field has been set.

### GetRequireFileServletAccessPrivilege

`func (o *AddHttpServletExtension200Response) GetRequireFileServletAccessPrivilege() bool`

GetRequireFileServletAccessPrivilege returns the RequireFileServletAccessPrivilege field if non-nil, zero value otherwise.

### GetRequireFileServletAccessPrivilegeOk

`func (o *AddHttpServletExtension200Response) GetRequireFileServletAccessPrivilegeOk() (*bool, bool)`

GetRequireFileServletAccessPrivilegeOk returns a tuple with the RequireFileServletAccessPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireFileServletAccessPrivilege

`func (o *AddHttpServletExtension200Response) SetRequireFileServletAccessPrivilege(v bool)`

SetRequireFileServletAccessPrivilege sets RequireFileServletAccessPrivilege field to given value.

### HasRequireFileServletAccessPrivilege

`func (o *AddHttpServletExtension200Response) HasRequireFileServletAccessPrivilege() bool`

HasRequireFileServletAccessPrivilege returns a boolean if a field has been set.

### GetRequireGroup

`func (o *AddHttpServletExtension200Response) GetRequireGroup() []string`

GetRequireGroup returns the RequireGroup field if non-nil, zero value otherwise.

### GetRequireGroupOk

`func (o *AddHttpServletExtension200Response) GetRequireGroupOk() (*[]string, bool)`

GetRequireGroupOk returns a tuple with the RequireGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireGroup

`func (o *AddHttpServletExtension200Response) SetRequireGroup(v []string)`

SetRequireGroup sets RequireGroup field to given value.

### HasRequireGroup

`func (o *AddHttpServletExtension200Response) HasRequireGroup() bool`

HasRequireGroup returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddHttpServletExtension200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddHttpServletExtension200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddHttpServletExtension200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddHttpServletExtension200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddHttpServletExtension200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddHttpServletExtension200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddHttpServletExtension200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


