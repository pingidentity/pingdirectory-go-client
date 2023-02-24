/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the FileServerHttpServletExtensionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileServerHttpServletExtensionResponse{}

// FileServerHttpServletExtensionResponse struct for FileServerHttpServletExtensionResponse
type FileServerHttpServletExtensionResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the HTTP Servlet Extension
	Id      string                                        `json:"id"`
	Schemas []EnumfileServerHttpServletExtensionSchemaUrn `json:"schemas"`
	// Specifies the base context path that should be used by HTTP clients to reference content. The value must start with a forward slash and must represent a valid HTTP context path.
	BaseContextPath string `json:"baseContextPath"`
	// Specifies the path to the directory on the local filesystem containing the files to be served by this File Server HTTP Servlet Extension. The path must exist, and it must be a directory.
	DocumentRootDirectory string `json:"documentRootDirectory"`
	// Indicates whether to generate a default HTML page with a listing of available files if the requested path refers to a directory rather than a file, and that directory does not contain an index file.
	EnableDirectoryIndexing *bool `json:"enableDirectoryIndexing,omitempty"`
	// Specifies the name of a file whose contents may be returned to the client if the requested path refers to a directory rather than a file.
	IndexFile []string `json:"indexFile,omitempty"`
	// Specifies the path to a file that contains MIME type mappings that will be used to determine the appropriate value to return for the Content-Type header based on the extension of the requested file.
	MimeTypesFile *string `json:"mimeTypesFile,omitempty"`
	// Specifies the default MIME type to use for the Content-Type header when a mapping cannot be found.
	DefaultMIMEType *string `json:"defaultMIMEType,omitempty"`
	// Indicates whether the servlet extension should only accept requests from authenticated clients.
	RequireAuthentication     *bool                                                   `json:"requireAuthentication,omitempty"`
	AllowedAuthenticationType []EnumhttpServletExtensionAllowedAuthenticationTypeProp `json:"allowedAuthenticationType,omitempty"`
	// The access token validators that may be used to verify the authenticity of an OAuth 2.0 bearer token.
	AccessTokenValidator []string `json:"accessTokenValidator,omitempty"`
	// The ID token validators that may be used to verify the authenticity of an of an OpenID Connect ID token.
	IdTokenValidator []string `json:"idTokenValidator,omitempty"`
	// Indicates whether the servlet extension should only accept requests from authenticated clients that have the file-servlet-access privilege.
	RequireFileServletAccessPrivilege *bool `json:"requireFileServletAccessPrivilege,omitempty"`
	// The DN of a group whose members will be permitted to access to the associated files. If multiple group DNs are configured, then anyone who is a member of at least one of those groups will be granted access.
	RequireGroup []string `json:"requireGroup,omitempty"`
	// The identity mapper that will be used to identify the entry with which a username is associated.
	IdentityMapper *string `json:"identityMapper,omitempty"`
	// A description for this HTTP Servlet Extension
	Description *string `json:"description,omitempty"`
	// The cross-origin request policy to use for the HTTP Servlet Extension.
	CrossOriginPolicy *string `json:"crossOriginPolicy,omitempty"`
	// Specifies HTTP header fields and values added to response headers for all requests.
	ResponseHeader []string `json:"responseHeader,omitempty"`
	// Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \"Correlation-Id\", \"X-Amzn-Trace-Id\", and \"X-Request-Id\".
	CorrelationIDResponseHeader *string `json:"correlationIDResponseHeader,omitempty"`
}

// NewFileServerHttpServletExtensionResponse instantiates a new FileServerHttpServletExtensionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileServerHttpServletExtensionResponse(id string, schemas []EnumfileServerHttpServletExtensionSchemaUrn, baseContextPath string, documentRootDirectory string) *FileServerHttpServletExtensionResponse {
	this := FileServerHttpServletExtensionResponse{}
	this.Id = id
	this.Schemas = schemas
	this.BaseContextPath = baseContextPath
	this.DocumentRootDirectory = documentRootDirectory
	return &this
}

// NewFileServerHttpServletExtensionResponseWithDefaults instantiates a new FileServerHttpServletExtensionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileServerHttpServletExtensionResponseWithDefaults() *FileServerHttpServletExtensionResponse {
	this := FileServerHttpServletExtensionResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *FileServerHttpServletExtensionResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileServerHttpServletExtensionResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *FileServerHttpServletExtensionResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *FileServerHttpServletExtensionResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *FileServerHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileServerHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *FileServerHttpServletExtensionResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *FileServerHttpServletExtensionResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *FileServerHttpServletExtensionResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FileServerHttpServletExtensionResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FileServerHttpServletExtensionResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *FileServerHttpServletExtensionResponse) GetSchemas() []EnumfileServerHttpServletExtensionSchemaUrn {
	if o == nil {
		var ret []EnumfileServerHttpServletExtensionSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *FileServerHttpServletExtensionResponse) GetSchemasOk() ([]EnumfileServerHttpServletExtensionSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *FileServerHttpServletExtensionResponse) SetSchemas(v []EnumfileServerHttpServletExtensionSchemaUrn) {
	o.Schemas = v
}

// GetBaseContextPath returns the BaseContextPath field value
func (o *FileServerHttpServletExtensionResponse) GetBaseContextPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseContextPath
}

// GetBaseContextPathOk returns a tuple with the BaseContextPath field value
// and a boolean to check if the value has been set.
func (o *FileServerHttpServletExtensionResponse) GetBaseContextPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseContextPath, true
}

// SetBaseContextPath sets field value
func (o *FileServerHttpServletExtensionResponse) SetBaseContextPath(v string) {
	o.BaseContextPath = v
}

// GetDocumentRootDirectory returns the DocumentRootDirectory field value
func (o *FileServerHttpServletExtensionResponse) GetDocumentRootDirectory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentRootDirectory
}

// GetDocumentRootDirectoryOk returns a tuple with the DocumentRootDirectory field value
// and a boolean to check if the value has been set.
func (o *FileServerHttpServletExtensionResponse) GetDocumentRootDirectoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentRootDirectory, true
}

// SetDocumentRootDirectory sets field value
func (o *FileServerHttpServletExtensionResponse) SetDocumentRootDirectory(v string) {
	o.DocumentRootDirectory = v
}

// GetEnableDirectoryIndexing returns the EnableDirectoryIndexing field value if set, zero value otherwise.
func (o *FileServerHttpServletExtensionResponse) GetEnableDirectoryIndexing() bool {
	if o == nil || IsNil(o.EnableDirectoryIndexing) {
		var ret bool
		return ret
	}
	return *o.EnableDirectoryIndexing
}

// GetEnableDirectoryIndexingOk returns a tuple with the EnableDirectoryIndexing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileServerHttpServletExtensionResponse) GetEnableDirectoryIndexingOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableDirectoryIndexing) {
		return nil, false
	}
	return o.EnableDirectoryIndexing, true
}

// HasEnableDirectoryIndexing returns a boolean if a field has been set.
func (o *FileServerHttpServletExtensionResponse) HasEnableDirectoryIndexing() bool {
	if o != nil && !IsNil(o.EnableDirectoryIndexing) {
		return true
	}

	return false
}

// SetEnableDirectoryIndexing gets a reference to the given bool and assigns it to the EnableDirectoryIndexing field.
func (o *FileServerHttpServletExtensionResponse) SetEnableDirectoryIndexing(v bool) {
	o.EnableDirectoryIndexing = &v
}

// GetIndexFile returns the IndexFile field value if set, zero value otherwise.
func (o *FileServerHttpServletExtensionResponse) GetIndexFile() []string {
	if o == nil || IsNil(o.IndexFile) {
		var ret []string
		return ret
	}
	return o.IndexFile
}

// GetIndexFileOk returns a tuple with the IndexFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileServerHttpServletExtensionResponse) GetIndexFileOk() ([]string, bool) {
	if o == nil || IsNil(o.IndexFile) {
		return nil, false
	}
	return o.IndexFile, true
}

// HasIndexFile returns a boolean if a field has been set.
func (o *FileServerHttpServletExtensionResponse) HasIndexFile() bool {
	if o != nil && !IsNil(o.IndexFile) {
		return true
	}

	return false
}

// SetIndexFile gets a reference to the given []string and assigns it to the IndexFile field.
func (o *FileServerHttpServletExtensionResponse) SetIndexFile(v []string) {
	o.IndexFile = v
}

// GetMimeTypesFile returns the MimeTypesFile field value if set, zero value otherwise.
func (o *FileServerHttpServletExtensionResponse) GetMimeTypesFile() string {
	if o == nil || IsNil(o.MimeTypesFile) {
		var ret string
		return ret
	}
	return *o.MimeTypesFile
}

// GetMimeTypesFileOk returns a tuple with the MimeTypesFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileServerHttpServletExtensionResponse) GetMimeTypesFileOk() (*string, bool) {
	if o == nil || IsNil(o.MimeTypesFile) {
		return nil, false
	}
	return o.MimeTypesFile, true
}

// HasMimeTypesFile returns a boolean if a field has been set.
func (o *FileServerHttpServletExtensionResponse) HasMimeTypesFile() bool {
	if o != nil && !IsNil(o.MimeTypesFile) {
		return true
	}

	return false
}

// SetMimeTypesFile gets a reference to the given string and assigns it to the MimeTypesFile field.
func (o *FileServerHttpServletExtensionResponse) SetMimeTypesFile(v string) {
	o.MimeTypesFile = &v
}

// GetDefaultMIMEType returns the DefaultMIMEType field value if set, zero value otherwise.
func (o *FileServerHttpServletExtensionResponse) GetDefaultMIMEType() string {
	if o == nil || IsNil(o.DefaultMIMEType) {
		var ret string
		return ret
	}
	return *o.DefaultMIMEType
}

// GetDefaultMIMETypeOk returns a tuple with the DefaultMIMEType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileServerHttpServletExtensionResponse) GetDefaultMIMETypeOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultMIMEType) {
		return nil, false
	}
	return o.DefaultMIMEType, true
}

// HasDefaultMIMEType returns a boolean if a field has been set.
func (o *FileServerHttpServletExtensionResponse) HasDefaultMIMEType() bool {
	if o != nil && !IsNil(o.DefaultMIMEType) {
		return true
	}

	return false
}

// SetDefaultMIMEType gets a reference to the given string and assigns it to the DefaultMIMEType field.
func (o *FileServerHttpServletExtensionResponse) SetDefaultMIMEType(v string) {
	o.DefaultMIMEType = &v
}

// GetRequireAuthentication returns the RequireAuthentication field value if set, zero value otherwise.
func (o *FileServerHttpServletExtensionResponse) GetRequireAuthentication() bool {
	if o == nil || IsNil(o.RequireAuthentication) {
		var ret bool
		return ret
	}
	return *o.RequireAuthentication
}

// GetRequireAuthenticationOk returns a tuple with the RequireAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileServerHttpServletExtensionResponse) GetRequireAuthenticationOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireAuthentication) {
		return nil, false
	}
	return o.RequireAuthentication, true
}

// HasRequireAuthentication returns a boolean if a field has been set.
func (o *FileServerHttpServletExtensionResponse) HasRequireAuthentication() bool {
	if o != nil && !IsNil(o.RequireAuthentication) {
		return true
	}

	return false
}

// SetRequireAuthentication gets a reference to the given bool and assigns it to the RequireAuthentication field.
func (o *FileServerHttpServletExtensionResponse) SetRequireAuthentication(v bool) {
	o.RequireAuthentication = &v
}

// GetAllowedAuthenticationType returns the AllowedAuthenticationType field value if set, zero value otherwise.
func (o *FileServerHttpServletExtensionResponse) GetAllowedAuthenticationType() []EnumhttpServletExtensionAllowedAuthenticationTypeProp {
	if o == nil || IsNil(o.AllowedAuthenticationType) {
		var ret []EnumhttpServletExtensionAllowedAuthenticationTypeProp
		return ret
	}
	return o.AllowedAuthenticationType
}

// GetAllowedAuthenticationTypeOk returns a tuple with the AllowedAuthenticationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileServerHttpServletExtensionResponse) GetAllowedAuthenticationTypeOk() ([]EnumhttpServletExtensionAllowedAuthenticationTypeProp, bool) {
	if o == nil || IsNil(o.AllowedAuthenticationType) {
		return nil, false
	}
	return o.AllowedAuthenticationType, true
}

// HasAllowedAuthenticationType returns a boolean if a field has been set.
func (o *FileServerHttpServletExtensionResponse) HasAllowedAuthenticationType() bool {
	if o != nil && !IsNil(o.AllowedAuthenticationType) {
		return true
	}

	return false
}

// SetAllowedAuthenticationType gets a reference to the given []EnumhttpServletExtensionAllowedAuthenticationTypeProp and assigns it to the AllowedAuthenticationType field.
func (o *FileServerHttpServletExtensionResponse) SetAllowedAuthenticationType(v []EnumhttpServletExtensionAllowedAuthenticationTypeProp) {
	o.AllowedAuthenticationType = v
}

// GetAccessTokenValidator returns the AccessTokenValidator field value if set, zero value otherwise.
func (o *FileServerHttpServletExtensionResponse) GetAccessTokenValidator() []string {
	if o == nil || IsNil(o.AccessTokenValidator) {
		var ret []string
		return ret
	}
	return o.AccessTokenValidator
}

// GetAccessTokenValidatorOk returns a tuple with the AccessTokenValidator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileServerHttpServletExtensionResponse) GetAccessTokenValidatorOk() ([]string, bool) {
	if o == nil || IsNil(o.AccessTokenValidator) {
		return nil, false
	}
	return o.AccessTokenValidator, true
}

// HasAccessTokenValidator returns a boolean if a field has been set.
func (o *FileServerHttpServletExtensionResponse) HasAccessTokenValidator() bool {
	if o != nil && !IsNil(o.AccessTokenValidator) {
		return true
	}

	return false
}

// SetAccessTokenValidator gets a reference to the given []string and assigns it to the AccessTokenValidator field.
func (o *FileServerHttpServletExtensionResponse) SetAccessTokenValidator(v []string) {
	o.AccessTokenValidator = v
}

// GetIdTokenValidator returns the IdTokenValidator field value if set, zero value otherwise.
func (o *FileServerHttpServletExtensionResponse) GetIdTokenValidator() []string {
	if o == nil || IsNil(o.IdTokenValidator) {
		var ret []string
		return ret
	}
	return o.IdTokenValidator
}

// GetIdTokenValidatorOk returns a tuple with the IdTokenValidator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileServerHttpServletExtensionResponse) GetIdTokenValidatorOk() ([]string, bool) {
	if o == nil || IsNil(o.IdTokenValidator) {
		return nil, false
	}
	return o.IdTokenValidator, true
}

// HasIdTokenValidator returns a boolean if a field has been set.
func (o *FileServerHttpServletExtensionResponse) HasIdTokenValidator() bool {
	if o != nil && !IsNil(o.IdTokenValidator) {
		return true
	}

	return false
}

// SetIdTokenValidator gets a reference to the given []string and assigns it to the IdTokenValidator field.
func (o *FileServerHttpServletExtensionResponse) SetIdTokenValidator(v []string) {
	o.IdTokenValidator = v
}

// GetRequireFileServletAccessPrivilege returns the RequireFileServletAccessPrivilege field value if set, zero value otherwise.
func (o *FileServerHttpServletExtensionResponse) GetRequireFileServletAccessPrivilege() bool {
	if o == nil || IsNil(o.RequireFileServletAccessPrivilege) {
		var ret bool
		return ret
	}
	return *o.RequireFileServletAccessPrivilege
}

// GetRequireFileServletAccessPrivilegeOk returns a tuple with the RequireFileServletAccessPrivilege field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileServerHttpServletExtensionResponse) GetRequireFileServletAccessPrivilegeOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireFileServletAccessPrivilege) {
		return nil, false
	}
	return o.RequireFileServletAccessPrivilege, true
}

// HasRequireFileServletAccessPrivilege returns a boolean if a field has been set.
func (o *FileServerHttpServletExtensionResponse) HasRequireFileServletAccessPrivilege() bool {
	if o != nil && !IsNil(o.RequireFileServletAccessPrivilege) {
		return true
	}

	return false
}

// SetRequireFileServletAccessPrivilege gets a reference to the given bool and assigns it to the RequireFileServletAccessPrivilege field.
func (o *FileServerHttpServletExtensionResponse) SetRequireFileServletAccessPrivilege(v bool) {
	o.RequireFileServletAccessPrivilege = &v
}

// GetRequireGroup returns the RequireGroup field value if set, zero value otherwise.
func (o *FileServerHttpServletExtensionResponse) GetRequireGroup() []string {
	if o == nil || IsNil(o.RequireGroup) {
		var ret []string
		return ret
	}
	return o.RequireGroup
}

// GetRequireGroupOk returns a tuple with the RequireGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileServerHttpServletExtensionResponse) GetRequireGroupOk() ([]string, bool) {
	if o == nil || IsNil(o.RequireGroup) {
		return nil, false
	}
	return o.RequireGroup, true
}

// HasRequireGroup returns a boolean if a field has been set.
func (o *FileServerHttpServletExtensionResponse) HasRequireGroup() bool {
	if o != nil && !IsNil(o.RequireGroup) {
		return true
	}

	return false
}

// SetRequireGroup gets a reference to the given []string and assigns it to the RequireGroup field.
func (o *FileServerHttpServletExtensionResponse) SetRequireGroup(v []string) {
	o.RequireGroup = v
}

// GetIdentityMapper returns the IdentityMapper field value if set, zero value otherwise.
func (o *FileServerHttpServletExtensionResponse) GetIdentityMapper() string {
	if o == nil || IsNil(o.IdentityMapper) {
		var ret string
		return ret
	}
	return *o.IdentityMapper
}

// GetIdentityMapperOk returns a tuple with the IdentityMapper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileServerHttpServletExtensionResponse) GetIdentityMapperOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityMapper) {
		return nil, false
	}
	return o.IdentityMapper, true
}

// HasIdentityMapper returns a boolean if a field has been set.
func (o *FileServerHttpServletExtensionResponse) HasIdentityMapper() bool {
	if o != nil && !IsNil(o.IdentityMapper) {
		return true
	}

	return false
}

// SetIdentityMapper gets a reference to the given string and assigns it to the IdentityMapper field.
func (o *FileServerHttpServletExtensionResponse) SetIdentityMapper(v string) {
	o.IdentityMapper = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FileServerHttpServletExtensionResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileServerHttpServletExtensionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FileServerHttpServletExtensionResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FileServerHttpServletExtensionResponse) SetDescription(v string) {
	o.Description = &v
}

// GetCrossOriginPolicy returns the CrossOriginPolicy field value if set, zero value otherwise.
func (o *FileServerHttpServletExtensionResponse) GetCrossOriginPolicy() string {
	if o == nil || IsNil(o.CrossOriginPolicy) {
		var ret string
		return ret
	}
	return *o.CrossOriginPolicy
}

// GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileServerHttpServletExtensionResponse) GetCrossOriginPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.CrossOriginPolicy) {
		return nil, false
	}
	return o.CrossOriginPolicy, true
}

// HasCrossOriginPolicy returns a boolean if a field has been set.
func (o *FileServerHttpServletExtensionResponse) HasCrossOriginPolicy() bool {
	if o != nil && !IsNil(o.CrossOriginPolicy) {
		return true
	}

	return false
}

// SetCrossOriginPolicy gets a reference to the given string and assigns it to the CrossOriginPolicy field.
func (o *FileServerHttpServletExtensionResponse) SetCrossOriginPolicy(v string) {
	o.CrossOriginPolicy = &v
}

// GetResponseHeader returns the ResponseHeader field value if set, zero value otherwise.
func (o *FileServerHttpServletExtensionResponse) GetResponseHeader() []string {
	if o == nil || IsNil(o.ResponseHeader) {
		var ret []string
		return ret
	}
	return o.ResponseHeader
}

// GetResponseHeaderOk returns a tuple with the ResponseHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileServerHttpServletExtensionResponse) GetResponseHeaderOk() ([]string, bool) {
	if o == nil || IsNil(o.ResponseHeader) {
		return nil, false
	}
	return o.ResponseHeader, true
}

// HasResponseHeader returns a boolean if a field has been set.
func (o *FileServerHttpServletExtensionResponse) HasResponseHeader() bool {
	if o != nil && !IsNil(o.ResponseHeader) {
		return true
	}

	return false
}

// SetResponseHeader gets a reference to the given []string and assigns it to the ResponseHeader field.
func (o *FileServerHttpServletExtensionResponse) SetResponseHeader(v []string) {
	o.ResponseHeader = v
}

// GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field value if set, zero value otherwise.
func (o *FileServerHttpServletExtensionResponse) GetCorrelationIDResponseHeader() string {
	if o == nil || IsNil(o.CorrelationIDResponseHeader) {
		var ret string
		return ret
	}
	return *o.CorrelationIDResponseHeader
}

// GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileServerHttpServletExtensionResponse) GetCorrelationIDResponseHeaderOk() (*string, bool) {
	if o == nil || IsNil(o.CorrelationIDResponseHeader) {
		return nil, false
	}
	return o.CorrelationIDResponseHeader, true
}

// HasCorrelationIDResponseHeader returns a boolean if a field has been set.
func (o *FileServerHttpServletExtensionResponse) HasCorrelationIDResponseHeader() bool {
	if o != nil && !IsNil(o.CorrelationIDResponseHeader) {
		return true
	}

	return false
}

// SetCorrelationIDResponseHeader gets a reference to the given string and assigns it to the CorrelationIDResponseHeader field.
func (o *FileServerHttpServletExtensionResponse) SetCorrelationIDResponseHeader(v string) {
	o.CorrelationIDResponseHeader = &v
}

func (o FileServerHttpServletExtensionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileServerHttpServletExtensionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	toSerialize["schemas"] = o.Schemas
	toSerialize["baseContextPath"] = o.BaseContextPath
	toSerialize["documentRootDirectory"] = o.DocumentRootDirectory
	if !IsNil(o.EnableDirectoryIndexing) {
		toSerialize["enableDirectoryIndexing"] = o.EnableDirectoryIndexing
	}
	if !IsNil(o.IndexFile) {
		toSerialize["indexFile"] = o.IndexFile
	}
	if !IsNil(o.MimeTypesFile) {
		toSerialize["mimeTypesFile"] = o.MimeTypesFile
	}
	if !IsNil(o.DefaultMIMEType) {
		toSerialize["defaultMIMEType"] = o.DefaultMIMEType
	}
	if !IsNil(o.RequireAuthentication) {
		toSerialize["requireAuthentication"] = o.RequireAuthentication
	}
	if !IsNil(o.AllowedAuthenticationType) {
		toSerialize["allowedAuthenticationType"] = o.AllowedAuthenticationType
	}
	if !IsNil(o.AccessTokenValidator) {
		toSerialize["accessTokenValidator"] = o.AccessTokenValidator
	}
	if !IsNil(o.IdTokenValidator) {
		toSerialize["idTokenValidator"] = o.IdTokenValidator
	}
	if !IsNil(o.RequireFileServletAccessPrivilege) {
		toSerialize["requireFileServletAccessPrivilege"] = o.RequireFileServletAccessPrivilege
	}
	if !IsNil(o.RequireGroup) {
		toSerialize["requireGroup"] = o.RequireGroup
	}
	if !IsNil(o.IdentityMapper) {
		toSerialize["identityMapper"] = o.IdentityMapper
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CrossOriginPolicy) {
		toSerialize["crossOriginPolicy"] = o.CrossOriginPolicy
	}
	if !IsNil(o.ResponseHeader) {
		toSerialize["responseHeader"] = o.ResponseHeader
	}
	if !IsNil(o.CorrelationIDResponseHeader) {
		toSerialize["correlationIDResponseHeader"] = o.CorrelationIDResponseHeader
	}
	return toSerialize, nil
}

type NullableFileServerHttpServletExtensionResponse struct {
	value *FileServerHttpServletExtensionResponse
	isSet bool
}

func (v NullableFileServerHttpServletExtensionResponse) Get() *FileServerHttpServletExtensionResponse {
	return v.value
}

func (v *NullableFileServerHttpServletExtensionResponse) Set(val *FileServerHttpServletExtensionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFileServerHttpServletExtensionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFileServerHttpServletExtensionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileServerHttpServletExtensionResponse(val *FileServerHttpServletExtensionResponse) *NullableFileServerHttpServletExtensionResponse {
	return &NullableFileServerHttpServletExtensionResponse{value: val, isSet: true}
}

func (v NullableFileServerHttpServletExtensionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileServerHttpServletExtensionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
