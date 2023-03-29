# LdapMappedScimHttpServletExtensionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the HTTP Servlet Extension | 
**Schemas** | [**[]EnumldapMappedScimHttpServletExtensionSchemaUrn**](EnumldapMappedScimHttpServletExtensionSchemaUrn.md) |  | 
**OAuthTokenHandler** | Pointer to **string** | Specifies the OAuth Token Handler implementation that should be used to validate OAuth 2.0 bearer tokens when they are included in a SCIM request. | [optional] 
**BasicAuthEnabled** | Pointer to **bool** | Enables HTTP Basic authentication, using a username and password. | [optional] 
**IdentityMapper** | Pointer to **string** | Specifies the name of the identity mapper that is to be used to match the username included in the HTTP Basic authentication header to the corresponding user in the directory. | [optional] 
**ResourceMappingFile** | Pointer to **string** | The path to an XML file defining the resources supported by the SCIM interface and the SCIM-to-LDAP attribute mappings to use. | [optional] 
**IncludeLDAPObjectclass** | Pointer to **[]string** | Specifies the LDAP object classes that should be exposed directly as SCIM resources. | [optional] 
**ExcludeLDAPObjectclass** | Pointer to **[]string** | Specifies the LDAP object classes that should be not be exposed directly as SCIM resources. | [optional] 
**IncludeLDAPBaseDN** | Pointer to **[]string** | Specifies the base DNs for the branches of the DIT that should be exposed via the Identity Access API. | [optional] 
**ExcludeLDAPBaseDN** | Pointer to **[]string** | Specifies the base DNs for the branches of the DIT that should not be exposed via the Identity Access API. | [optional] 
**EntityTagLDAPAttribute** | Pointer to **string** | Specifies the LDAP attribute whose value should be used as the entity tag value to enable SCIM resource versioning support. | [optional] 
**BaseContextPath** | **string** | The context path to use to access the SCIM interface. The value must start with a forward slash and must represent a valid HTTP context path. | 
**TemporaryDirectory** | **string** | Specifies the location of the directory that is used to create temporary files containing SCIM request data. | 
**TemporaryDirectoryPermissions** | **string** | Specifies the permissions that should be applied to the directory that is used to create temporary files. | 
**MaxResults** | Pointer to **int32** | The maximum number of resources that are returned in a response. | [optional] 
**BulkMaxOperations** | Pointer to **int32** | The maximum number of operations that are permitted in a bulk request. | [optional] 
**BulkMaxPayloadSize** | Pointer to **string** | The maximum payload size in bytes of a bulk request. | [optional] 
**BulkMaxConcurrentRequests** | Pointer to **int32** | The maximum number of bulk requests that may be processed concurrently by the server. Any bulk request that would cause this limit to be exceeded is rejected with HTTP status code 503. | [optional] 
**DebugEnabled** | Pointer to **bool** | Enables debug logging of the SCIM SDK. Debug messages will be forwarded to the Directory Server debug logger with the scope of com.unboundid.directory.server.extensions.scim.SCIMHTTPServletExtension. | [optional] 
**DebugLevel** | [**EnumhttpServletExtensionDebugLevelProp**](EnumhttpServletExtensionDebugLevelProp.md) |  | 
**DebugType** | [**[]EnumhttpServletExtensionDebugTypeProp**](EnumhttpServletExtensionDebugTypeProp.md) |  | 
**IncludeStackTrace** | **bool** | Indicates whether a stack trace of the thread which called the debug method should be included in debug log messages. | 
**Description** | Pointer to **string** | A description for this HTTP Servlet Extension | [optional] 
**CrossOriginPolicy** | Pointer to **string** | The cross-origin request policy to use for the HTTP Servlet Extension. | [optional] 
**ResponseHeader** | Pointer to **[]string** | Specifies HTTP header fields and values added to response headers for all requests. | [optional] 
**CorrelationIDResponseHeader** | Pointer to **string** | Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \&quot;Correlation-Id\&quot;, \&quot;X-Amzn-Trace-Id\&quot;, and \&quot;X-Request-Id\&quot;. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewLdapMappedScimHttpServletExtensionResponse

`func NewLdapMappedScimHttpServletExtensionResponse(id string, schemas []EnumldapMappedScimHttpServletExtensionSchemaUrn, baseContextPath string, temporaryDirectory string, temporaryDirectoryPermissions string, debugLevel EnumhttpServletExtensionDebugLevelProp, debugType []EnumhttpServletExtensionDebugTypeProp, includeStackTrace bool, ) *LdapMappedScimHttpServletExtensionResponse`

NewLdapMappedScimHttpServletExtensionResponse instantiates a new LdapMappedScimHttpServletExtensionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapMappedScimHttpServletExtensionResponseWithDefaults

`func NewLdapMappedScimHttpServletExtensionResponseWithDefaults() *LdapMappedScimHttpServletExtensionResponse`

NewLdapMappedScimHttpServletExtensionResponseWithDefaults instantiates a new LdapMappedScimHttpServletExtensionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LdapMappedScimHttpServletExtensionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LdapMappedScimHttpServletExtensionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LdapMappedScimHttpServletExtensionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *LdapMappedScimHttpServletExtensionResponse) GetSchemas() []EnumldapMappedScimHttpServletExtensionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *LdapMappedScimHttpServletExtensionResponse) GetSchemasOk() (*[]EnumldapMappedScimHttpServletExtensionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *LdapMappedScimHttpServletExtensionResponse) SetSchemas(v []EnumldapMappedScimHttpServletExtensionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetOAuthTokenHandler

`func (o *LdapMappedScimHttpServletExtensionResponse) GetOAuthTokenHandler() string`

GetOAuthTokenHandler returns the OAuthTokenHandler field if non-nil, zero value otherwise.

### GetOAuthTokenHandlerOk

`func (o *LdapMappedScimHttpServletExtensionResponse) GetOAuthTokenHandlerOk() (*string, bool)`

GetOAuthTokenHandlerOk returns a tuple with the OAuthTokenHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthTokenHandler

`func (o *LdapMappedScimHttpServletExtensionResponse) SetOAuthTokenHandler(v string)`

SetOAuthTokenHandler sets OAuthTokenHandler field to given value.

### HasOAuthTokenHandler

`func (o *LdapMappedScimHttpServletExtensionResponse) HasOAuthTokenHandler() bool`

HasOAuthTokenHandler returns a boolean if a field has been set.

### GetBasicAuthEnabled

`func (o *LdapMappedScimHttpServletExtensionResponse) GetBasicAuthEnabled() bool`

GetBasicAuthEnabled returns the BasicAuthEnabled field if non-nil, zero value otherwise.

### GetBasicAuthEnabledOk

`func (o *LdapMappedScimHttpServletExtensionResponse) GetBasicAuthEnabledOk() (*bool, bool)`

GetBasicAuthEnabledOk returns a tuple with the BasicAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthEnabled

`func (o *LdapMappedScimHttpServletExtensionResponse) SetBasicAuthEnabled(v bool)`

SetBasicAuthEnabled sets BasicAuthEnabled field to given value.

### HasBasicAuthEnabled

`func (o *LdapMappedScimHttpServletExtensionResponse) HasBasicAuthEnabled() bool`

HasBasicAuthEnabled returns a boolean if a field has been set.

### GetIdentityMapper

`func (o *LdapMappedScimHttpServletExtensionResponse) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *LdapMappedScimHttpServletExtensionResponse) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *LdapMappedScimHttpServletExtensionResponse) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.

### HasIdentityMapper

`func (o *LdapMappedScimHttpServletExtensionResponse) HasIdentityMapper() bool`

HasIdentityMapper returns a boolean if a field has been set.

### GetResourceMappingFile

`func (o *LdapMappedScimHttpServletExtensionResponse) GetResourceMappingFile() string`

GetResourceMappingFile returns the ResourceMappingFile field if non-nil, zero value otherwise.

### GetResourceMappingFileOk

`func (o *LdapMappedScimHttpServletExtensionResponse) GetResourceMappingFileOk() (*string, bool)`

GetResourceMappingFileOk returns a tuple with the ResourceMappingFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceMappingFile

`func (o *LdapMappedScimHttpServletExtensionResponse) SetResourceMappingFile(v string)`

SetResourceMappingFile sets ResourceMappingFile field to given value.

### HasResourceMappingFile

`func (o *LdapMappedScimHttpServletExtensionResponse) HasResourceMappingFile() bool`

HasResourceMappingFile returns a boolean if a field has been set.

### GetIncludeLDAPObjectclass

`func (o *LdapMappedScimHttpServletExtensionResponse) GetIncludeLDAPObjectclass() []string`

GetIncludeLDAPObjectclass returns the IncludeLDAPObjectclass field if non-nil, zero value otherwise.

### GetIncludeLDAPObjectclassOk

`func (o *LdapMappedScimHttpServletExtensionResponse) GetIncludeLDAPObjectclassOk() (*[]string, bool)`

GetIncludeLDAPObjectclassOk returns a tuple with the IncludeLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLDAPObjectclass

`func (o *LdapMappedScimHttpServletExtensionResponse) SetIncludeLDAPObjectclass(v []string)`

SetIncludeLDAPObjectclass sets IncludeLDAPObjectclass field to given value.

### HasIncludeLDAPObjectclass

`func (o *LdapMappedScimHttpServletExtensionResponse) HasIncludeLDAPObjectclass() bool`

HasIncludeLDAPObjectclass returns a boolean if a field has been set.

### GetExcludeLDAPObjectclass

`func (o *LdapMappedScimHttpServletExtensionResponse) GetExcludeLDAPObjectclass() []string`

GetExcludeLDAPObjectclass returns the ExcludeLDAPObjectclass field if non-nil, zero value otherwise.

### GetExcludeLDAPObjectclassOk

`func (o *LdapMappedScimHttpServletExtensionResponse) GetExcludeLDAPObjectclassOk() (*[]string, bool)`

GetExcludeLDAPObjectclassOk returns a tuple with the ExcludeLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeLDAPObjectclass

`func (o *LdapMappedScimHttpServletExtensionResponse) SetExcludeLDAPObjectclass(v []string)`

SetExcludeLDAPObjectclass sets ExcludeLDAPObjectclass field to given value.

### HasExcludeLDAPObjectclass

`func (o *LdapMappedScimHttpServletExtensionResponse) HasExcludeLDAPObjectclass() bool`

HasExcludeLDAPObjectclass returns a boolean if a field has been set.

### GetIncludeLDAPBaseDN

`func (o *LdapMappedScimHttpServletExtensionResponse) GetIncludeLDAPBaseDN() []string`

GetIncludeLDAPBaseDN returns the IncludeLDAPBaseDN field if non-nil, zero value otherwise.

### GetIncludeLDAPBaseDNOk

`func (o *LdapMappedScimHttpServletExtensionResponse) GetIncludeLDAPBaseDNOk() (*[]string, bool)`

GetIncludeLDAPBaseDNOk returns a tuple with the IncludeLDAPBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLDAPBaseDN

`func (o *LdapMappedScimHttpServletExtensionResponse) SetIncludeLDAPBaseDN(v []string)`

SetIncludeLDAPBaseDN sets IncludeLDAPBaseDN field to given value.

### HasIncludeLDAPBaseDN

`func (o *LdapMappedScimHttpServletExtensionResponse) HasIncludeLDAPBaseDN() bool`

HasIncludeLDAPBaseDN returns a boolean if a field has been set.

### GetExcludeLDAPBaseDN

`func (o *LdapMappedScimHttpServletExtensionResponse) GetExcludeLDAPBaseDN() []string`

GetExcludeLDAPBaseDN returns the ExcludeLDAPBaseDN field if non-nil, zero value otherwise.

### GetExcludeLDAPBaseDNOk

`func (o *LdapMappedScimHttpServletExtensionResponse) GetExcludeLDAPBaseDNOk() (*[]string, bool)`

GetExcludeLDAPBaseDNOk returns a tuple with the ExcludeLDAPBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeLDAPBaseDN

`func (o *LdapMappedScimHttpServletExtensionResponse) SetExcludeLDAPBaseDN(v []string)`

SetExcludeLDAPBaseDN sets ExcludeLDAPBaseDN field to given value.

### HasExcludeLDAPBaseDN

`func (o *LdapMappedScimHttpServletExtensionResponse) HasExcludeLDAPBaseDN() bool`

HasExcludeLDAPBaseDN returns a boolean if a field has been set.

### GetEntityTagLDAPAttribute

`func (o *LdapMappedScimHttpServletExtensionResponse) GetEntityTagLDAPAttribute() string`

GetEntityTagLDAPAttribute returns the EntityTagLDAPAttribute field if non-nil, zero value otherwise.

### GetEntityTagLDAPAttributeOk

`func (o *LdapMappedScimHttpServletExtensionResponse) GetEntityTagLDAPAttributeOk() (*string, bool)`

GetEntityTagLDAPAttributeOk returns a tuple with the EntityTagLDAPAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTagLDAPAttribute

`func (o *LdapMappedScimHttpServletExtensionResponse) SetEntityTagLDAPAttribute(v string)`

SetEntityTagLDAPAttribute sets EntityTagLDAPAttribute field to given value.

### HasEntityTagLDAPAttribute

`func (o *LdapMappedScimHttpServletExtensionResponse) HasEntityTagLDAPAttribute() bool`

HasEntityTagLDAPAttribute returns a boolean if a field has been set.

### GetBaseContextPath

`func (o *LdapMappedScimHttpServletExtensionResponse) GetBaseContextPath() string`

GetBaseContextPath returns the BaseContextPath field if non-nil, zero value otherwise.

### GetBaseContextPathOk

`func (o *LdapMappedScimHttpServletExtensionResponse) GetBaseContextPathOk() (*string, bool)`

GetBaseContextPathOk returns a tuple with the BaseContextPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseContextPath

`func (o *LdapMappedScimHttpServletExtensionResponse) SetBaseContextPath(v string)`

SetBaseContextPath sets BaseContextPath field to given value.


### GetTemporaryDirectory

`func (o *LdapMappedScimHttpServletExtensionResponse) GetTemporaryDirectory() string`

GetTemporaryDirectory returns the TemporaryDirectory field if non-nil, zero value otherwise.

### GetTemporaryDirectoryOk

`func (o *LdapMappedScimHttpServletExtensionResponse) GetTemporaryDirectoryOk() (*string, bool)`

GetTemporaryDirectoryOk returns a tuple with the TemporaryDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryDirectory

`func (o *LdapMappedScimHttpServletExtensionResponse) SetTemporaryDirectory(v string)`

SetTemporaryDirectory sets TemporaryDirectory field to given value.


### GetTemporaryDirectoryPermissions

`func (o *LdapMappedScimHttpServletExtensionResponse) GetTemporaryDirectoryPermissions() string`

GetTemporaryDirectoryPermissions returns the TemporaryDirectoryPermissions field if non-nil, zero value otherwise.

### GetTemporaryDirectoryPermissionsOk

`func (o *LdapMappedScimHttpServletExtensionResponse) GetTemporaryDirectoryPermissionsOk() (*string, bool)`

GetTemporaryDirectoryPermissionsOk returns a tuple with the TemporaryDirectoryPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryDirectoryPermissions

`func (o *LdapMappedScimHttpServletExtensionResponse) SetTemporaryDirectoryPermissions(v string)`

SetTemporaryDirectoryPermissions sets TemporaryDirectoryPermissions field to given value.


### GetMaxResults

`func (o *LdapMappedScimHttpServletExtensionResponse) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *LdapMappedScimHttpServletExtensionResponse) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *LdapMappedScimHttpServletExtensionResponse) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *LdapMappedScimHttpServletExtensionResponse) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetBulkMaxOperations

`func (o *LdapMappedScimHttpServletExtensionResponse) GetBulkMaxOperations() int32`

GetBulkMaxOperations returns the BulkMaxOperations field if non-nil, zero value otherwise.

### GetBulkMaxOperationsOk

`func (o *LdapMappedScimHttpServletExtensionResponse) GetBulkMaxOperationsOk() (*int32, bool)`

GetBulkMaxOperationsOk returns a tuple with the BulkMaxOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkMaxOperations

`func (o *LdapMappedScimHttpServletExtensionResponse) SetBulkMaxOperations(v int32)`

SetBulkMaxOperations sets BulkMaxOperations field to given value.

### HasBulkMaxOperations

`func (o *LdapMappedScimHttpServletExtensionResponse) HasBulkMaxOperations() bool`

HasBulkMaxOperations returns a boolean if a field has been set.

### GetBulkMaxPayloadSize

`func (o *LdapMappedScimHttpServletExtensionResponse) GetBulkMaxPayloadSize() string`

GetBulkMaxPayloadSize returns the BulkMaxPayloadSize field if non-nil, zero value otherwise.

### GetBulkMaxPayloadSizeOk

`func (o *LdapMappedScimHttpServletExtensionResponse) GetBulkMaxPayloadSizeOk() (*string, bool)`

GetBulkMaxPayloadSizeOk returns a tuple with the BulkMaxPayloadSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkMaxPayloadSize

`func (o *LdapMappedScimHttpServletExtensionResponse) SetBulkMaxPayloadSize(v string)`

SetBulkMaxPayloadSize sets BulkMaxPayloadSize field to given value.

### HasBulkMaxPayloadSize

`func (o *LdapMappedScimHttpServletExtensionResponse) HasBulkMaxPayloadSize() bool`

HasBulkMaxPayloadSize returns a boolean if a field has been set.

### GetBulkMaxConcurrentRequests

`func (o *LdapMappedScimHttpServletExtensionResponse) GetBulkMaxConcurrentRequests() int32`

GetBulkMaxConcurrentRequests returns the BulkMaxConcurrentRequests field if non-nil, zero value otherwise.

### GetBulkMaxConcurrentRequestsOk

`func (o *LdapMappedScimHttpServletExtensionResponse) GetBulkMaxConcurrentRequestsOk() (*int32, bool)`

GetBulkMaxConcurrentRequestsOk returns a tuple with the BulkMaxConcurrentRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkMaxConcurrentRequests

`func (o *LdapMappedScimHttpServletExtensionResponse) SetBulkMaxConcurrentRequests(v int32)`

SetBulkMaxConcurrentRequests sets BulkMaxConcurrentRequests field to given value.

### HasBulkMaxConcurrentRequests

`func (o *LdapMappedScimHttpServletExtensionResponse) HasBulkMaxConcurrentRequests() bool`

HasBulkMaxConcurrentRequests returns a boolean if a field has been set.

### GetDebugEnabled

`func (o *LdapMappedScimHttpServletExtensionResponse) GetDebugEnabled() bool`

GetDebugEnabled returns the DebugEnabled field if non-nil, zero value otherwise.

### GetDebugEnabledOk

`func (o *LdapMappedScimHttpServletExtensionResponse) GetDebugEnabledOk() (*bool, bool)`

GetDebugEnabledOk returns a tuple with the DebugEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugEnabled

`func (o *LdapMappedScimHttpServletExtensionResponse) SetDebugEnabled(v bool)`

SetDebugEnabled sets DebugEnabled field to given value.

### HasDebugEnabled

`func (o *LdapMappedScimHttpServletExtensionResponse) HasDebugEnabled() bool`

HasDebugEnabled returns a boolean if a field has been set.

### GetDebugLevel

`func (o *LdapMappedScimHttpServletExtensionResponse) GetDebugLevel() EnumhttpServletExtensionDebugLevelProp`

GetDebugLevel returns the DebugLevel field if non-nil, zero value otherwise.

### GetDebugLevelOk

`func (o *LdapMappedScimHttpServletExtensionResponse) GetDebugLevelOk() (*EnumhttpServletExtensionDebugLevelProp, bool)`

GetDebugLevelOk returns a tuple with the DebugLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugLevel

`func (o *LdapMappedScimHttpServletExtensionResponse) SetDebugLevel(v EnumhttpServletExtensionDebugLevelProp)`

SetDebugLevel sets DebugLevel field to given value.


### GetDebugType

`func (o *LdapMappedScimHttpServletExtensionResponse) GetDebugType() []EnumhttpServletExtensionDebugTypeProp`

GetDebugType returns the DebugType field if non-nil, zero value otherwise.

### GetDebugTypeOk

`func (o *LdapMappedScimHttpServletExtensionResponse) GetDebugTypeOk() (*[]EnumhttpServletExtensionDebugTypeProp, bool)`

GetDebugTypeOk returns a tuple with the DebugType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugType

`func (o *LdapMappedScimHttpServletExtensionResponse) SetDebugType(v []EnumhttpServletExtensionDebugTypeProp)`

SetDebugType sets DebugType field to given value.


### GetIncludeStackTrace

`func (o *LdapMappedScimHttpServletExtensionResponse) GetIncludeStackTrace() bool`

GetIncludeStackTrace returns the IncludeStackTrace field if non-nil, zero value otherwise.

### GetIncludeStackTraceOk

`func (o *LdapMappedScimHttpServletExtensionResponse) GetIncludeStackTraceOk() (*bool, bool)`

GetIncludeStackTraceOk returns a tuple with the IncludeStackTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeStackTrace

`func (o *LdapMappedScimHttpServletExtensionResponse) SetIncludeStackTrace(v bool)`

SetIncludeStackTrace sets IncludeStackTrace field to given value.


### GetDescription

`func (o *LdapMappedScimHttpServletExtensionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LdapMappedScimHttpServletExtensionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LdapMappedScimHttpServletExtensionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LdapMappedScimHttpServletExtensionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCrossOriginPolicy

`func (o *LdapMappedScimHttpServletExtensionResponse) GetCrossOriginPolicy() string`

GetCrossOriginPolicy returns the CrossOriginPolicy field if non-nil, zero value otherwise.

### GetCrossOriginPolicyOk

`func (o *LdapMappedScimHttpServletExtensionResponse) GetCrossOriginPolicyOk() (*string, bool)`

GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossOriginPolicy

`func (o *LdapMappedScimHttpServletExtensionResponse) SetCrossOriginPolicy(v string)`

SetCrossOriginPolicy sets CrossOriginPolicy field to given value.

### HasCrossOriginPolicy

`func (o *LdapMappedScimHttpServletExtensionResponse) HasCrossOriginPolicy() bool`

HasCrossOriginPolicy returns a boolean if a field has been set.

### GetResponseHeader

`func (o *LdapMappedScimHttpServletExtensionResponse) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *LdapMappedScimHttpServletExtensionResponse) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *LdapMappedScimHttpServletExtensionResponse) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *LdapMappedScimHttpServletExtensionResponse) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetCorrelationIDResponseHeader

`func (o *LdapMappedScimHttpServletExtensionResponse) GetCorrelationIDResponseHeader() string`

GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field if non-nil, zero value otherwise.

### GetCorrelationIDResponseHeaderOk

`func (o *LdapMappedScimHttpServletExtensionResponse) GetCorrelationIDResponseHeaderOk() (*string, bool)`

GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDResponseHeader

`func (o *LdapMappedScimHttpServletExtensionResponse) SetCorrelationIDResponseHeader(v string)`

SetCorrelationIDResponseHeader sets CorrelationIDResponseHeader field to given value.

### HasCorrelationIDResponseHeader

`func (o *LdapMappedScimHttpServletExtensionResponse) HasCorrelationIDResponseHeader() bool`

HasCorrelationIDResponseHeader returns a boolean if a field has been set.

### GetMeta

`func (o *LdapMappedScimHttpServletExtensionResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *LdapMappedScimHttpServletExtensionResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *LdapMappedScimHttpServletExtensionResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *LdapMappedScimHttpServletExtensionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *LdapMappedScimHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *LdapMappedScimHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *LdapMappedScimHttpServletExtensionResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *LdapMappedScimHttpServletExtensionResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


