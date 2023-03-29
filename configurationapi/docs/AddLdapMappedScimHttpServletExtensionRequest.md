# AddLdapMappedScimHttpServletExtensionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtensionName** | **string** | Name of the new HTTP Servlet Extension | 
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
**BaseContextPath** | Pointer to **string** | The context path to use to access the SCIM interface. The value must start with a forward slash and must represent a valid HTTP context path. | [optional] 
**TemporaryDirectory** | Pointer to **string** | Specifies the location of the directory that is used to create temporary files containing SCIM request data. | [optional] 
**TemporaryDirectoryPermissions** | Pointer to **string** | Specifies the permissions that should be applied to the directory that is used to create temporary files. | [optional] 
**MaxResults** | Pointer to **int32** | The maximum number of resources that are returned in a response. | [optional] 
**BulkMaxOperations** | Pointer to **int32** | The maximum number of operations that are permitted in a bulk request. | [optional] 
**BulkMaxPayloadSize** | Pointer to **string** | The maximum payload size in bytes of a bulk request. | [optional] 
**BulkMaxConcurrentRequests** | Pointer to **int32** | The maximum number of bulk requests that may be processed concurrently by the server. Any bulk request that would cause this limit to be exceeded is rejected with HTTP status code 503. | [optional] 
**DebugEnabled** | Pointer to **bool** | Enables debug logging of the SCIM SDK. Debug messages will be forwarded to the Directory Server debug logger with the scope of com.unboundid.directory.server.extensions.scim.SCIMHTTPServletExtension. | [optional] 
**DebugLevel** | Pointer to [**EnumhttpServletExtensionDebugLevelProp**](EnumhttpServletExtensionDebugLevelProp.md) |  | [optional] 
**DebugType** | Pointer to [**[]EnumhttpServletExtensionDebugTypeProp**](EnumhttpServletExtensionDebugTypeProp.md) | The types of debug messages that should be logged. | [optional] 
**IncludeStackTrace** | Pointer to **bool** | Indicates whether a stack trace of the thread which called the debug method should be included in debug log messages. | [optional] 
**Description** | Pointer to **string** | A description for this HTTP Servlet Extension | [optional] 
**CrossOriginPolicy** | Pointer to **string** | The cross-origin request policy to use for the HTTP Servlet Extension. | [optional] 
**ResponseHeader** | Pointer to **[]string** | Specifies HTTP header fields and values added to response headers for all requests. | [optional] 
**CorrelationIDResponseHeader** | Pointer to **string** | Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \&quot;Correlation-Id\&quot;, \&quot;X-Amzn-Trace-Id\&quot;, and \&quot;X-Request-Id\&quot;. | [optional] 

## Methods

### NewAddLdapMappedScimHttpServletExtensionRequest

`func NewAddLdapMappedScimHttpServletExtensionRequest(extensionName string, schemas []EnumldapMappedScimHttpServletExtensionSchemaUrn, ) *AddLdapMappedScimHttpServletExtensionRequest`

NewAddLdapMappedScimHttpServletExtensionRequest instantiates a new AddLdapMappedScimHttpServletExtensionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLdapMappedScimHttpServletExtensionRequestWithDefaults

`func NewAddLdapMappedScimHttpServletExtensionRequestWithDefaults() *AddLdapMappedScimHttpServletExtensionRequest`

NewAddLdapMappedScimHttpServletExtensionRequestWithDefaults instantiates a new AddLdapMappedScimHttpServletExtensionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensionName

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetExtensionName() string`

GetExtensionName returns the ExtensionName field if non-nil, zero value otherwise.

### GetExtensionNameOk

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetExtensionNameOk() (*string, bool)`

GetExtensionNameOk returns a tuple with the ExtensionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionName

`func (o *AddLdapMappedScimHttpServletExtensionRequest) SetExtensionName(v string)`

SetExtensionName sets ExtensionName field to given value.


### GetSchemas

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetSchemas() []EnumldapMappedScimHttpServletExtensionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetSchemasOk() (*[]EnumldapMappedScimHttpServletExtensionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddLdapMappedScimHttpServletExtensionRequest) SetSchemas(v []EnumldapMappedScimHttpServletExtensionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetOAuthTokenHandler

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetOAuthTokenHandler() string`

GetOAuthTokenHandler returns the OAuthTokenHandler field if non-nil, zero value otherwise.

### GetOAuthTokenHandlerOk

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetOAuthTokenHandlerOk() (*string, bool)`

GetOAuthTokenHandlerOk returns a tuple with the OAuthTokenHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthTokenHandler

`func (o *AddLdapMappedScimHttpServletExtensionRequest) SetOAuthTokenHandler(v string)`

SetOAuthTokenHandler sets OAuthTokenHandler field to given value.

### HasOAuthTokenHandler

`func (o *AddLdapMappedScimHttpServletExtensionRequest) HasOAuthTokenHandler() bool`

HasOAuthTokenHandler returns a boolean if a field has been set.

### GetBasicAuthEnabled

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetBasicAuthEnabled() bool`

GetBasicAuthEnabled returns the BasicAuthEnabled field if non-nil, zero value otherwise.

### GetBasicAuthEnabledOk

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetBasicAuthEnabledOk() (*bool, bool)`

GetBasicAuthEnabledOk returns a tuple with the BasicAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthEnabled

`func (o *AddLdapMappedScimHttpServletExtensionRequest) SetBasicAuthEnabled(v bool)`

SetBasicAuthEnabled sets BasicAuthEnabled field to given value.

### HasBasicAuthEnabled

`func (o *AddLdapMappedScimHttpServletExtensionRequest) HasBasicAuthEnabled() bool`

HasBasicAuthEnabled returns a boolean if a field has been set.

### GetIdentityMapper

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *AddLdapMappedScimHttpServletExtensionRequest) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.

### HasIdentityMapper

`func (o *AddLdapMappedScimHttpServletExtensionRequest) HasIdentityMapper() bool`

HasIdentityMapper returns a boolean if a field has been set.

### GetResourceMappingFile

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetResourceMappingFile() string`

GetResourceMappingFile returns the ResourceMappingFile field if non-nil, zero value otherwise.

### GetResourceMappingFileOk

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetResourceMappingFileOk() (*string, bool)`

GetResourceMappingFileOk returns a tuple with the ResourceMappingFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceMappingFile

`func (o *AddLdapMappedScimHttpServletExtensionRequest) SetResourceMappingFile(v string)`

SetResourceMappingFile sets ResourceMappingFile field to given value.

### HasResourceMappingFile

`func (o *AddLdapMappedScimHttpServletExtensionRequest) HasResourceMappingFile() bool`

HasResourceMappingFile returns a boolean if a field has been set.

### GetIncludeLDAPObjectclass

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetIncludeLDAPObjectclass() []string`

GetIncludeLDAPObjectclass returns the IncludeLDAPObjectclass field if non-nil, zero value otherwise.

### GetIncludeLDAPObjectclassOk

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetIncludeLDAPObjectclassOk() (*[]string, bool)`

GetIncludeLDAPObjectclassOk returns a tuple with the IncludeLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLDAPObjectclass

`func (o *AddLdapMappedScimHttpServletExtensionRequest) SetIncludeLDAPObjectclass(v []string)`

SetIncludeLDAPObjectclass sets IncludeLDAPObjectclass field to given value.

### HasIncludeLDAPObjectclass

`func (o *AddLdapMappedScimHttpServletExtensionRequest) HasIncludeLDAPObjectclass() bool`

HasIncludeLDAPObjectclass returns a boolean if a field has been set.

### GetExcludeLDAPObjectclass

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetExcludeLDAPObjectclass() []string`

GetExcludeLDAPObjectclass returns the ExcludeLDAPObjectclass field if non-nil, zero value otherwise.

### GetExcludeLDAPObjectclassOk

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetExcludeLDAPObjectclassOk() (*[]string, bool)`

GetExcludeLDAPObjectclassOk returns a tuple with the ExcludeLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeLDAPObjectclass

`func (o *AddLdapMappedScimHttpServletExtensionRequest) SetExcludeLDAPObjectclass(v []string)`

SetExcludeLDAPObjectclass sets ExcludeLDAPObjectclass field to given value.

### HasExcludeLDAPObjectclass

`func (o *AddLdapMappedScimHttpServletExtensionRequest) HasExcludeLDAPObjectclass() bool`

HasExcludeLDAPObjectclass returns a boolean if a field has been set.

### GetIncludeLDAPBaseDN

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetIncludeLDAPBaseDN() []string`

GetIncludeLDAPBaseDN returns the IncludeLDAPBaseDN field if non-nil, zero value otherwise.

### GetIncludeLDAPBaseDNOk

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetIncludeLDAPBaseDNOk() (*[]string, bool)`

GetIncludeLDAPBaseDNOk returns a tuple with the IncludeLDAPBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLDAPBaseDN

`func (o *AddLdapMappedScimHttpServletExtensionRequest) SetIncludeLDAPBaseDN(v []string)`

SetIncludeLDAPBaseDN sets IncludeLDAPBaseDN field to given value.

### HasIncludeLDAPBaseDN

`func (o *AddLdapMappedScimHttpServletExtensionRequest) HasIncludeLDAPBaseDN() bool`

HasIncludeLDAPBaseDN returns a boolean if a field has been set.

### GetExcludeLDAPBaseDN

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetExcludeLDAPBaseDN() []string`

GetExcludeLDAPBaseDN returns the ExcludeLDAPBaseDN field if non-nil, zero value otherwise.

### GetExcludeLDAPBaseDNOk

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetExcludeLDAPBaseDNOk() (*[]string, bool)`

GetExcludeLDAPBaseDNOk returns a tuple with the ExcludeLDAPBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeLDAPBaseDN

`func (o *AddLdapMappedScimHttpServletExtensionRequest) SetExcludeLDAPBaseDN(v []string)`

SetExcludeLDAPBaseDN sets ExcludeLDAPBaseDN field to given value.

### HasExcludeLDAPBaseDN

`func (o *AddLdapMappedScimHttpServletExtensionRequest) HasExcludeLDAPBaseDN() bool`

HasExcludeLDAPBaseDN returns a boolean if a field has been set.

### GetEntityTagLDAPAttribute

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetEntityTagLDAPAttribute() string`

GetEntityTagLDAPAttribute returns the EntityTagLDAPAttribute field if non-nil, zero value otherwise.

### GetEntityTagLDAPAttributeOk

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetEntityTagLDAPAttributeOk() (*string, bool)`

GetEntityTagLDAPAttributeOk returns a tuple with the EntityTagLDAPAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTagLDAPAttribute

`func (o *AddLdapMappedScimHttpServletExtensionRequest) SetEntityTagLDAPAttribute(v string)`

SetEntityTagLDAPAttribute sets EntityTagLDAPAttribute field to given value.

### HasEntityTagLDAPAttribute

`func (o *AddLdapMappedScimHttpServletExtensionRequest) HasEntityTagLDAPAttribute() bool`

HasEntityTagLDAPAttribute returns a boolean if a field has been set.

### GetBaseContextPath

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetBaseContextPath() string`

GetBaseContextPath returns the BaseContextPath field if non-nil, zero value otherwise.

### GetBaseContextPathOk

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetBaseContextPathOk() (*string, bool)`

GetBaseContextPathOk returns a tuple with the BaseContextPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseContextPath

`func (o *AddLdapMappedScimHttpServletExtensionRequest) SetBaseContextPath(v string)`

SetBaseContextPath sets BaseContextPath field to given value.

### HasBaseContextPath

`func (o *AddLdapMappedScimHttpServletExtensionRequest) HasBaseContextPath() bool`

HasBaseContextPath returns a boolean if a field has been set.

### GetTemporaryDirectory

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetTemporaryDirectory() string`

GetTemporaryDirectory returns the TemporaryDirectory field if non-nil, zero value otherwise.

### GetTemporaryDirectoryOk

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetTemporaryDirectoryOk() (*string, bool)`

GetTemporaryDirectoryOk returns a tuple with the TemporaryDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryDirectory

`func (o *AddLdapMappedScimHttpServletExtensionRequest) SetTemporaryDirectory(v string)`

SetTemporaryDirectory sets TemporaryDirectory field to given value.

### HasTemporaryDirectory

`func (o *AddLdapMappedScimHttpServletExtensionRequest) HasTemporaryDirectory() bool`

HasTemporaryDirectory returns a boolean if a field has been set.

### GetTemporaryDirectoryPermissions

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetTemporaryDirectoryPermissions() string`

GetTemporaryDirectoryPermissions returns the TemporaryDirectoryPermissions field if non-nil, zero value otherwise.

### GetTemporaryDirectoryPermissionsOk

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetTemporaryDirectoryPermissionsOk() (*string, bool)`

GetTemporaryDirectoryPermissionsOk returns a tuple with the TemporaryDirectoryPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryDirectoryPermissions

`func (o *AddLdapMappedScimHttpServletExtensionRequest) SetTemporaryDirectoryPermissions(v string)`

SetTemporaryDirectoryPermissions sets TemporaryDirectoryPermissions field to given value.

### HasTemporaryDirectoryPermissions

`func (o *AddLdapMappedScimHttpServletExtensionRequest) HasTemporaryDirectoryPermissions() bool`

HasTemporaryDirectoryPermissions returns a boolean if a field has been set.

### GetMaxResults

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *AddLdapMappedScimHttpServletExtensionRequest) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *AddLdapMappedScimHttpServletExtensionRequest) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetBulkMaxOperations

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetBulkMaxOperations() int32`

GetBulkMaxOperations returns the BulkMaxOperations field if non-nil, zero value otherwise.

### GetBulkMaxOperationsOk

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetBulkMaxOperationsOk() (*int32, bool)`

GetBulkMaxOperationsOk returns a tuple with the BulkMaxOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkMaxOperations

`func (o *AddLdapMappedScimHttpServletExtensionRequest) SetBulkMaxOperations(v int32)`

SetBulkMaxOperations sets BulkMaxOperations field to given value.

### HasBulkMaxOperations

`func (o *AddLdapMappedScimHttpServletExtensionRequest) HasBulkMaxOperations() bool`

HasBulkMaxOperations returns a boolean if a field has been set.

### GetBulkMaxPayloadSize

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetBulkMaxPayloadSize() string`

GetBulkMaxPayloadSize returns the BulkMaxPayloadSize field if non-nil, zero value otherwise.

### GetBulkMaxPayloadSizeOk

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetBulkMaxPayloadSizeOk() (*string, bool)`

GetBulkMaxPayloadSizeOk returns a tuple with the BulkMaxPayloadSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkMaxPayloadSize

`func (o *AddLdapMappedScimHttpServletExtensionRequest) SetBulkMaxPayloadSize(v string)`

SetBulkMaxPayloadSize sets BulkMaxPayloadSize field to given value.

### HasBulkMaxPayloadSize

`func (o *AddLdapMappedScimHttpServletExtensionRequest) HasBulkMaxPayloadSize() bool`

HasBulkMaxPayloadSize returns a boolean if a field has been set.

### GetBulkMaxConcurrentRequests

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetBulkMaxConcurrentRequests() int32`

GetBulkMaxConcurrentRequests returns the BulkMaxConcurrentRequests field if non-nil, zero value otherwise.

### GetBulkMaxConcurrentRequestsOk

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetBulkMaxConcurrentRequestsOk() (*int32, bool)`

GetBulkMaxConcurrentRequestsOk returns a tuple with the BulkMaxConcurrentRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkMaxConcurrentRequests

`func (o *AddLdapMappedScimHttpServletExtensionRequest) SetBulkMaxConcurrentRequests(v int32)`

SetBulkMaxConcurrentRequests sets BulkMaxConcurrentRequests field to given value.

### HasBulkMaxConcurrentRequests

`func (o *AddLdapMappedScimHttpServletExtensionRequest) HasBulkMaxConcurrentRequests() bool`

HasBulkMaxConcurrentRequests returns a boolean if a field has been set.

### GetDebugEnabled

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetDebugEnabled() bool`

GetDebugEnabled returns the DebugEnabled field if non-nil, zero value otherwise.

### GetDebugEnabledOk

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetDebugEnabledOk() (*bool, bool)`

GetDebugEnabledOk returns a tuple with the DebugEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugEnabled

`func (o *AddLdapMappedScimHttpServletExtensionRequest) SetDebugEnabled(v bool)`

SetDebugEnabled sets DebugEnabled field to given value.

### HasDebugEnabled

`func (o *AddLdapMappedScimHttpServletExtensionRequest) HasDebugEnabled() bool`

HasDebugEnabled returns a boolean if a field has been set.

### GetDebugLevel

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetDebugLevel() EnumhttpServletExtensionDebugLevelProp`

GetDebugLevel returns the DebugLevel field if non-nil, zero value otherwise.

### GetDebugLevelOk

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetDebugLevelOk() (*EnumhttpServletExtensionDebugLevelProp, bool)`

GetDebugLevelOk returns a tuple with the DebugLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugLevel

`func (o *AddLdapMappedScimHttpServletExtensionRequest) SetDebugLevel(v EnumhttpServletExtensionDebugLevelProp)`

SetDebugLevel sets DebugLevel field to given value.

### HasDebugLevel

`func (o *AddLdapMappedScimHttpServletExtensionRequest) HasDebugLevel() bool`

HasDebugLevel returns a boolean if a field has been set.

### GetDebugType

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetDebugType() []EnumhttpServletExtensionDebugTypeProp`

GetDebugType returns the DebugType field if non-nil, zero value otherwise.

### GetDebugTypeOk

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetDebugTypeOk() (*[]EnumhttpServletExtensionDebugTypeProp, bool)`

GetDebugTypeOk returns a tuple with the DebugType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugType

`func (o *AddLdapMappedScimHttpServletExtensionRequest) SetDebugType(v []EnumhttpServletExtensionDebugTypeProp)`

SetDebugType sets DebugType field to given value.

### HasDebugType

`func (o *AddLdapMappedScimHttpServletExtensionRequest) HasDebugType() bool`

HasDebugType returns a boolean if a field has been set.

### GetIncludeStackTrace

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetIncludeStackTrace() bool`

GetIncludeStackTrace returns the IncludeStackTrace field if non-nil, zero value otherwise.

### GetIncludeStackTraceOk

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetIncludeStackTraceOk() (*bool, bool)`

GetIncludeStackTraceOk returns a tuple with the IncludeStackTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeStackTrace

`func (o *AddLdapMappedScimHttpServletExtensionRequest) SetIncludeStackTrace(v bool)`

SetIncludeStackTrace sets IncludeStackTrace field to given value.

### HasIncludeStackTrace

`func (o *AddLdapMappedScimHttpServletExtensionRequest) HasIncludeStackTrace() bool`

HasIncludeStackTrace returns a boolean if a field has been set.

### GetDescription

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddLdapMappedScimHttpServletExtensionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddLdapMappedScimHttpServletExtensionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCrossOriginPolicy

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetCrossOriginPolicy() string`

GetCrossOriginPolicy returns the CrossOriginPolicy field if non-nil, zero value otherwise.

### GetCrossOriginPolicyOk

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetCrossOriginPolicyOk() (*string, bool)`

GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossOriginPolicy

`func (o *AddLdapMappedScimHttpServletExtensionRequest) SetCrossOriginPolicy(v string)`

SetCrossOriginPolicy sets CrossOriginPolicy field to given value.

### HasCrossOriginPolicy

`func (o *AddLdapMappedScimHttpServletExtensionRequest) HasCrossOriginPolicy() bool`

HasCrossOriginPolicy returns a boolean if a field has been set.

### GetResponseHeader

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *AddLdapMappedScimHttpServletExtensionRequest) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *AddLdapMappedScimHttpServletExtensionRequest) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetCorrelationIDResponseHeader

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetCorrelationIDResponseHeader() string`

GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field if non-nil, zero value otherwise.

### GetCorrelationIDResponseHeaderOk

`func (o *AddLdapMappedScimHttpServletExtensionRequest) GetCorrelationIDResponseHeaderOk() (*string, bool)`

GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDResponseHeader

`func (o *AddLdapMappedScimHttpServletExtensionRequest) SetCorrelationIDResponseHeader(v string)`

SetCorrelationIDResponseHeader sets CorrelationIDResponseHeader field to given value.

### HasCorrelationIDResponseHeader

`func (o *AddLdapMappedScimHttpServletExtensionRequest) HasCorrelationIDResponseHeader() bool`

HasCorrelationIDResponseHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


