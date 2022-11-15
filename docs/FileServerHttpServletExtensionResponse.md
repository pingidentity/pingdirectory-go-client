# FileServerHttpServletExtensionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the HTTP Servlet Extension | 
**Schemas** | [**[]EnumfileServerHttpServletExtensionSchemaUrn**](EnumfileServerHttpServletExtensionSchemaUrn.md) |  | 
**BaseContextPath** | **string** | Specifies the base context path that should be used by HTTP clients to reference content. The value must start with a forward slash and must represent a valid HTTP context path. | 
**DocumentRootDirectory** | **string** | Specifies the path to the directory on the local filesystem containing the files to be served by this File Server HTTP Servlet Extension. The path must exist, and it must be a directory. | 
**EnableDirectoryIndexing** | Pointer to **bool** | Indicates whether to generate a default HTML page with a listing of available files if the requested path refers to a directory rather than a file, and that directory does not contain an index file. | [optional] 
**IndexFile** | Pointer to **[]string** |  | [optional] 
**MimeTypesFile** | Pointer to **string** | Specifies the path to a file that contains MIME type mappings that will be used to determine the appropriate value to return for the Content-Type header based on the extension of the requested file. | [optional] 
**DefaultMIMEType** | Pointer to **string** | Specifies the default MIME type to use for the Content-Type header when a mapping cannot be found. | [optional] 
**RequireAuthentication** | Pointer to **bool** | Indicates whether the servlet extension should only accept requests from authenticated clients. | [optional] 
**AllowedAuthenticationType** | Pointer to [**[]EnumhttpServletExtensionAllowedAuthenticationTypeProp**](EnumhttpServletExtensionAllowedAuthenticationTypeProp.md) |  | [optional] 
**AccessTokenValidator** | Pointer to **[]string** |  | [optional] 
**IdTokenValidator** | Pointer to **[]string** |  | [optional] 
**RequireFileServletAccessPrivilege** | Pointer to **bool** | Indicates whether the servlet extension should only accept requests from authenticated clients that have the file-servlet-access privilege. | [optional] 
**RequireGroup** | Pointer to **[]string** |  | [optional] 
**IdentityMapper** | Pointer to **string** | The identity mapper that will be used to identify the entry with which a username is associated. | [optional] 
**Description** | Pointer to **string** | A description for this HTTP Servlet Extension | [optional] 
**CrossOriginPolicy** | Pointer to **string** | The cross-origin request policy to use for the HTTP Servlet Extension. | [optional] 
**ResponseHeader** | Pointer to **[]string** |  | [optional] 
**CorrelationIDResponseHeader** | Pointer to **string** | Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \&quot;Correlation-Id\&quot;, \&quot;X-Amzn-Trace-Id\&quot;, and \&quot;X-Request-Id\&quot;. | [optional] 

## Methods

### NewFileServerHttpServletExtensionResponse

`func NewFileServerHttpServletExtensionResponse(id string, schemas []EnumfileServerHttpServletExtensionSchemaUrn, baseContextPath string, documentRootDirectory string, ) *FileServerHttpServletExtensionResponse`

NewFileServerHttpServletExtensionResponse instantiates a new FileServerHttpServletExtensionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileServerHttpServletExtensionResponseWithDefaults

`func NewFileServerHttpServletExtensionResponseWithDefaults() *FileServerHttpServletExtensionResponse`

NewFileServerHttpServletExtensionResponseWithDefaults instantiates a new FileServerHttpServletExtensionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FileServerHttpServletExtensionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileServerHttpServletExtensionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileServerHttpServletExtensionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *FileServerHttpServletExtensionResponse) GetSchemas() []EnumfileServerHttpServletExtensionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *FileServerHttpServletExtensionResponse) GetSchemasOk() (*[]EnumfileServerHttpServletExtensionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *FileServerHttpServletExtensionResponse) SetSchemas(v []EnumfileServerHttpServletExtensionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetBaseContextPath

`func (o *FileServerHttpServletExtensionResponse) GetBaseContextPath() string`

GetBaseContextPath returns the BaseContextPath field if non-nil, zero value otherwise.

### GetBaseContextPathOk

`func (o *FileServerHttpServletExtensionResponse) GetBaseContextPathOk() (*string, bool)`

GetBaseContextPathOk returns a tuple with the BaseContextPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseContextPath

`func (o *FileServerHttpServletExtensionResponse) SetBaseContextPath(v string)`

SetBaseContextPath sets BaseContextPath field to given value.


### GetDocumentRootDirectory

`func (o *FileServerHttpServletExtensionResponse) GetDocumentRootDirectory() string`

GetDocumentRootDirectory returns the DocumentRootDirectory field if non-nil, zero value otherwise.

### GetDocumentRootDirectoryOk

`func (o *FileServerHttpServletExtensionResponse) GetDocumentRootDirectoryOk() (*string, bool)`

GetDocumentRootDirectoryOk returns a tuple with the DocumentRootDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentRootDirectory

`func (o *FileServerHttpServletExtensionResponse) SetDocumentRootDirectory(v string)`

SetDocumentRootDirectory sets DocumentRootDirectory field to given value.


### GetEnableDirectoryIndexing

`func (o *FileServerHttpServletExtensionResponse) GetEnableDirectoryIndexing() bool`

GetEnableDirectoryIndexing returns the EnableDirectoryIndexing field if non-nil, zero value otherwise.

### GetEnableDirectoryIndexingOk

`func (o *FileServerHttpServletExtensionResponse) GetEnableDirectoryIndexingOk() (*bool, bool)`

GetEnableDirectoryIndexingOk returns a tuple with the EnableDirectoryIndexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDirectoryIndexing

`func (o *FileServerHttpServletExtensionResponse) SetEnableDirectoryIndexing(v bool)`

SetEnableDirectoryIndexing sets EnableDirectoryIndexing field to given value.

### HasEnableDirectoryIndexing

`func (o *FileServerHttpServletExtensionResponse) HasEnableDirectoryIndexing() bool`

HasEnableDirectoryIndexing returns a boolean if a field has been set.

### GetIndexFile

`func (o *FileServerHttpServletExtensionResponse) GetIndexFile() []string`

GetIndexFile returns the IndexFile field if non-nil, zero value otherwise.

### GetIndexFileOk

`func (o *FileServerHttpServletExtensionResponse) GetIndexFileOk() (*[]string, bool)`

GetIndexFileOk returns a tuple with the IndexFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexFile

`func (o *FileServerHttpServletExtensionResponse) SetIndexFile(v []string)`

SetIndexFile sets IndexFile field to given value.

### HasIndexFile

`func (o *FileServerHttpServletExtensionResponse) HasIndexFile() bool`

HasIndexFile returns a boolean if a field has been set.

### GetMimeTypesFile

`func (o *FileServerHttpServletExtensionResponse) GetMimeTypesFile() string`

GetMimeTypesFile returns the MimeTypesFile field if non-nil, zero value otherwise.

### GetMimeTypesFileOk

`func (o *FileServerHttpServletExtensionResponse) GetMimeTypesFileOk() (*string, bool)`

GetMimeTypesFileOk returns a tuple with the MimeTypesFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeTypesFile

`func (o *FileServerHttpServletExtensionResponse) SetMimeTypesFile(v string)`

SetMimeTypesFile sets MimeTypesFile field to given value.

### HasMimeTypesFile

`func (o *FileServerHttpServletExtensionResponse) HasMimeTypesFile() bool`

HasMimeTypesFile returns a boolean if a field has been set.

### GetDefaultMIMEType

`func (o *FileServerHttpServletExtensionResponse) GetDefaultMIMEType() string`

GetDefaultMIMEType returns the DefaultMIMEType field if non-nil, zero value otherwise.

### GetDefaultMIMETypeOk

`func (o *FileServerHttpServletExtensionResponse) GetDefaultMIMETypeOk() (*string, bool)`

GetDefaultMIMETypeOk returns a tuple with the DefaultMIMEType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMIMEType

`func (o *FileServerHttpServletExtensionResponse) SetDefaultMIMEType(v string)`

SetDefaultMIMEType sets DefaultMIMEType field to given value.

### HasDefaultMIMEType

`func (o *FileServerHttpServletExtensionResponse) HasDefaultMIMEType() bool`

HasDefaultMIMEType returns a boolean if a field has been set.

### GetRequireAuthentication

`func (o *FileServerHttpServletExtensionResponse) GetRequireAuthentication() bool`

GetRequireAuthentication returns the RequireAuthentication field if non-nil, zero value otherwise.

### GetRequireAuthenticationOk

`func (o *FileServerHttpServletExtensionResponse) GetRequireAuthenticationOk() (*bool, bool)`

GetRequireAuthenticationOk returns a tuple with the RequireAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAuthentication

`func (o *FileServerHttpServletExtensionResponse) SetRequireAuthentication(v bool)`

SetRequireAuthentication sets RequireAuthentication field to given value.

### HasRequireAuthentication

`func (o *FileServerHttpServletExtensionResponse) HasRequireAuthentication() bool`

HasRequireAuthentication returns a boolean if a field has been set.

### GetAllowedAuthenticationType

`func (o *FileServerHttpServletExtensionResponse) GetAllowedAuthenticationType() []EnumhttpServletExtensionAllowedAuthenticationTypeProp`

GetAllowedAuthenticationType returns the AllowedAuthenticationType field if non-nil, zero value otherwise.

### GetAllowedAuthenticationTypeOk

`func (o *FileServerHttpServletExtensionResponse) GetAllowedAuthenticationTypeOk() (*[]EnumhttpServletExtensionAllowedAuthenticationTypeProp, bool)`

GetAllowedAuthenticationTypeOk returns a tuple with the AllowedAuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAuthenticationType

`func (o *FileServerHttpServletExtensionResponse) SetAllowedAuthenticationType(v []EnumhttpServletExtensionAllowedAuthenticationTypeProp)`

SetAllowedAuthenticationType sets AllowedAuthenticationType field to given value.

### HasAllowedAuthenticationType

`func (o *FileServerHttpServletExtensionResponse) HasAllowedAuthenticationType() bool`

HasAllowedAuthenticationType returns a boolean if a field has been set.

### GetAccessTokenValidator

`func (o *FileServerHttpServletExtensionResponse) GetAccessTokenValidator() []string`

GetAccessTokenValidator returns the AccessTokenValidator field if non-nil, zero value otherwise.

### GetAccessTokenValidatorOk

`func (o *FileServerHttpServletExtensionResponse) GetAccessTokenValidatorOk() (*[]string, bool)`

GetAccessTokenValidatorOk returns a tuple with the AccessTokenValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValidator

`func (o *FileServerHttpServletExtensionResponse) SetAccessTokenValidator(v []string)`

SetAccessTokenValidator sets AccessTokenValidator field to given value.

### HasAccessTokenValidator

`func (o *FileServerHttpServletExtensionResponse) HasAccessTokenValidator() bool`

HasAccessTokenValidator returns a boolean if a field has been set.

### GetIdTokenValidator

`func (o *FileServerHttpServletExtensionResponse) GetIdTokenValidator() []string`

GetIdTokenValidator returns the IdTokenValidator field if non-nil, zero value otherwise.

### GetIdTokenValidatorOk

`func (o *FileServerHttpServletExtensionResponse) GetIdTokenValidatorOk() (*[]string, bool)`

GetIdTokenValidatorOk returns a tuple with the IdTokenValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenValidator

`func (o *FileServerHttpServletExtensionResponse) SetIdTokenValidator(v []string)`

SetIdTokenValidator sets IdTokenValidator field to given value.

### HasIdTokenValidator

`func (o *FileServerHttpServletExtensionResponse) HasIdTokenValidator() bool`

HasIdTokenValidator returns a boolean if a field has been set.

### GetRequireFileServletAccessPrivilege

`func (o *FileServerHttpServletExtensionResponse) GetRequireFileServletAccessPrivilege() bool`

GetRequireFileServletAccessPrivilege returns the RequireFileServletAccessPrivilege field if non-nil, zero value otherwise.

### GetRequireFileServletAccessPrivilegeOk

`func (o *FileServerHttpServletExtensionResponse) GetRequireFileServletAccessPrivilegeOk() (*bool, bool)`

GetRequireFileServletAccessPrivilegeOk returns a tuple with the RequireFileServletAccessPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireFileServletAccessPrivilege

`func (o *FileServerHttpServletExtensionResponse) SetRequireFileServletAccessPrivilege(v bool)`

SetRequireFileServletAccessPrivilege sets RequireFileServletAccessPrivilege field to given value.

### HasRequireFileServletAccessPrivilege

`func (o *FileServerHttpServletExtensionResponse) HasRequireFileServletAccessPrivilege() bool`

HasRequireFileServletAccessPrivilege returns a boolean if a field has been set.

### GetRequireGroup

`func (o *FileServerHttpServletExtensionResponse) GetRequireGroup() []string`

GetRequireGroup returns the RequireGroup field if non-nil, zero value otherwise.

### GetRequireGroupOk

`func (o *FileServerHttpServletExtensionResponse) GetRequireGroupOk() (*[]string, bool)`

GetRequireGroupOk returns a tuple with the RequireGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireGroup

`func (o *FileServerHttpServletExtensionResponse) SetRequireGroup(v []string)`

SetRequireGroup sets RequireGroup field to given value.

### HasRequireGroup

`func (o *FileServerHttpServletExtensionResponse) HasRequireGroup() bool`

HasRequireGroup returns a boolean if a field has been set.

### GetIdentityMapper

`func (o *FileServerHttpServletExtensionResponse) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *FileServerHttpServletExtensionResponse) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *FileServerHttpServletExtensionResponse) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.

### HasIdentityMapper

`func (o *FileServerHttpServletExtensionResponse) HasIdentityMapper() bool`

HasIdentityMapper returns a boolean if a field has been set.

### GetDescription

`func (o *FileServerHttpServletExtensionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FileServerHttpServletExtensionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FileServerHttpServletExtensionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FileServerHttpServletExtensionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCrossOriginPolicy

`func (o *FileServerHttpServletExtensionResponse) GetCrossOriginPolicy() string`

GetCrossOriginPolicy returns the CrossOriginPolicy field if non-nil, zero value otherwise.

### GetCrossOriginPolicyOk

`func (o *FileServerHttpServletExtensionResponse) GetCrossOriginPolicyOk() (*string, bool)`

GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossOriginPolicy

`func (o *FileServerHttpServletExtensionResponse) SetCrossOriginPolicy(v string)`

SetCrossOriginPolicy sets CrossOriginPolicy field to given value.

### HasCrossOriginPolicy

`func (o *FileServerHttpServletExtensionResponse) HasCrossOriginPolicy() bool`

HasCrossOriginPolicy returns a boolean if a field has been set.

### GetResponseHeader

`func (o *FileServerHttpServletExtensionResponse) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *FileServerHttpServletExtensionResponse) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *FileServerHttpServletExtensionResponse) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *FileServerHttpServletExtensionResponse) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetCorrelationIDResponseHeader

`func (o *FileServerHttpServletExtensionResponse) GetCorrelationIDResponseHeader() string`

GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field if non-nil, zero value otherwise.

### GetCorrelationIDResponseHeaderOk

`func (o *FileServerHttpServletExtensionResponse) GetCorrelationIDResponseHeaderOk() (*string, bool)`

GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDResponseHeader

`func (o *FileServerHttpServletExtensionResponse) SetCorrelationIDResponseHeader(v string)`

SetCorrelationIDResponseHeader sets CorrelationIDResponseHeader field to given value.

### HasCorrelationIDResponseHeader

`func (o *FileServerHttpServletExtensionResponse) HasCorrelationIDResponseHeader() bool`

HasCorrelationIDResponseHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


