/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AddFileServerHttpServletExtensionRequest struct for AddFileServerHttpServletExtensionRequest
type AddFileServerHttpServletExtensionRequest struct {
	// Name of the new HTTP Servlet Extension
	ExtensionName string `json:"extensionName"`
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
	RequireAuthentication *bool `json:"requireAuthentication,omitempty"`
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

// NewAddFileServerHttpServletExtensionRequest instantiates a new AddFileServerHttpServletExtensionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddFileServerHttpServletExtensionRequest(extensionName string, schemas []EnumfileServerHttpServletExtensionSchemaUrn, baseContextPath string, documentRootDirectory string) *AddFileServerHttpServletExtensionRequest {
	this := AddFileServerHttpServletExtensionRequest{}
	this.ExtensionName = extensionName
	this.Schemas = schemas
	this.BaseContextPath = baseContextPath
	this.DocumentRootDirectory = documentRootDirectory
	return &this
}

// NewAddFileServerHttpServletExtensionRequestWithDefaults instantiates a new AddFileServerHttpServletExtensionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddFileServerHttpServletExtensionRequestWithDefaults() *AddFileServerHttpServletExtensionRequest {
	this := AddFileServerHttpServletExtensionRequest{}
	return &this
}

// GetExtensionName returns the ExtensionName field value
func (o *AddFileServerHttpServletExtensionRequest) GetExtensionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionName
}

// GetExtensionNameOk returns a tuple with the ExtensionName field value
// and a boolean to check if the value has been set.
func (o *AddFileServerHttpServletExtensionRequest) GetExtensionNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExtensionName, true
}

// SetExtensionName sets field value
func (o *AddFileServerHttpServletExtensionRequest) SetExtensionName(v string) {
	o.ExtensionName = v
}

// GetSchemas returns the Schemas field value
func (o *AddFileServerHttpServletExtensionRequest) GetSchemas() []EnumfileServerHttpServletExtensionSchemaUrn {
	if o == nil {
		var ret []EnumfileServerHttpServletExtensionSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddFileServerHttpServletExtensionRequest) GetSchemasOk() ([]EnumfileServerHttpServletExtensionSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddFileServerHttpServletExtensionRequest) SetSchemas(v []EnumfileServerHttpServletExtensionSchemaUrn) {
	o.Schemas = v
}

// GetBaseContextPath returns the BaseContextPath field value
func (o *AddFileServerHttpServletExtensionRequest) GetBaseContextPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseContextPath
}

// GetBaseContextPathOk returns a tuple with the BaseContextPath field value
// and a boolean to check if the value has been set.
func (o *AddFileServerHttpServletExtensionRequest) GetBaseContextPathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.BaseContextPath, true
}

// SetBaseContextPath sets field value
func (o *AddFileServerHttpServletExtensionRequest) SetBaseContextPath(v string) {
	o.BaseContextPath = v
}

// GetDocumentRootDirectory returns the DocumentRootDirectory field value
func (o *AddFileServerHttpServletExtensionRequest) GetDocumentRootDirectory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentRootDirectory
}

// GetDocumentRootDirectoryOk returns a tuple with the DocumentRootDirectory field value
// and a boolean to check if the value has been set.
func (o *AddFileServerHttpServletExtensionRequest) GetDocumentRootDirectoryOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DocumentRootDirectory, true
}

// SetDocumentRootDirectory sets field value
func (o *AddFileServerHttpServletExtensionRequest) SetDocumentRootDirectory(v string) {
	o.DocumentRootDirectory = v
}

// GetEnableDirectoryIndexing returns the EnableDirectoryIndexing field value if set, zero value otherwise.
func (o *AddFileServerHttpServletExtensionRequest) GetEnableDirectoryIndexing() bool {
	if o == nil || isNil(o.EnableDirectoryIndexing) {
		var ret bool
		return ret
	}
	return *o.EnableDirectoryIndexing
}

// GetEnableDirectoryIndexingOk returns a tuple with the EnableDirectoryIndexing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFileServerHttpServletExtensionRequest) GetEnableDirectoryIndexingOk() (*bool, bool) {
	if o == nil || isNil(o.EnableDirectoryIndexing) {
    return nil, false
	}
	return o.EnableDirectoryIndexing, true
}

// HasEnableDirectoryIndexing returns a boolean if a field has been set.
func (o *AddFileServerHttpServletExtensionRequest) HasEnableDirectoryIndexing() bool {
	if o != nil && !isNil(o.EnableDirectoryIndexing) {
		return true
	}

	return false
}

// SetEnableDirectoryIndexing gets a reference to the given bool and assigns it to the EnableDirectoryIndexing field.
func (o *AddFileServerHttpServletExtensionRequest) SetEnableDirectoryIndexing(v bool) {
	o.EnableDirectoryIndexing = &v
}

// GetIndexFile returns the IndexFile field value if set, zero value otherwise.
func (o *AddFileServerHttpServletExtensionRequest) GetIndexFile() []string {
	if o == nil || isNil(o.IndexFile) {
		var ret []string
		return ret
	}
	return o.IndexFile
}

// GetIndexFileOk returns a tuple with the IndexFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFileServerHttpServletExtensionRequest) GetIndexFileOk() ([]string, bool) {
	if o == nil || isNil(o.IndexFile) {
    return nil, false
	}
	return o.IndexFile, true
}

// HasIndexFile returns a boolean if a field has been set.
func (o *AddFileServerHttpServletExtensionRequest) HasIndexFile() bool {
	if o != nil && !isNil(o.IndexFile) {
		return true
	}

	return false
}

// SetIndexFile gets a reference to the given []string and assigns it to the IndexFile field.
func (o *AddFileServerHttpServletExtensionRequest) SetIndexFile(v []string) {
	o.IndexFile = v
}

// GetMimeTypesFile returns the MimeTypesFile field value if set, zero value otherwise.
func (o *AddFileServerHttpServletExtensionRequest) GetMimeTypesFile() string {
	if o == nil || isNil(o.MimeTypesFile) {
		var ret string
		return ret
	}
	return *o.MimeTypesFile
}

// GetMimeTypesFileOk returns a tuple with the MimeTypesFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFileServerHttpServletExtensionRequest) GetMimeTypesFileOk() (*string, bool) {
	if o == nil || isNil(o.MimeTypesFile) {
    return nil, false
	}
	return o.MimeTypesFile, true
}

// HasMimeTypesFile returns a boolean if a field has been set.
func (o *AddFileServerHttpServletExtensionRequest) HasMimeTypesFile() bool {
	if o != nil && !isNil(o.MimeTypesFile) {
		return true
	}

	return false
}

// SetMimeTypesFile gets a reference to the given string and assigns it to the MimeTypesFile field.
func (o *AddFileServerHttpServletExtensionRequest) SetMimeTypesFile(v string) {
	o.MimeTypesFile = &v
}

// GetDefaultMIMEType returns the DefaultMIMEType field value if set, zero value otherwise.
func (o *AddFileServerHttpServletExtensionRequest) GetDefaultMIMEType() string {
	if o == nil || isNil(o.DefaultMIMEType) {
		var ret string
		return ret
	}
	return *o.DefaultMIMEType
}

// GetDefaultMIMETypeOk returns a tuple with the DefaultMIMEType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFileServerHttpServletExtensionRequest) GetDefaultMIMETypeOk() (*string, bool) {
	if o == nil || isNil(o.DefaultMIMEType) {
    return nil, false
	}
	return o.DefaultMIMEType, true
}

// HasDefaultMIMEType returns a boolean if a field has been set.
func (o *AddFileServerHttpServletExtensionRequest) HasDefaultMIMEType() bool {
	if o != nil && !isNil(o.DefaultMIMEType) {
		return true
	}

	return false
}

// SetDefaultMIMEType gets a reference to the given string and assigns it to the DefaultMIMEType field.
func (o *AddFileServerHttpServletExtensionRequest) SetDefaultMIMEType(v string) {
	o.DefaultMIMEType = &v
}

// GetRequireAuthentication returns the RequireAuthentication field value if set, zero value otherwise.
func (o *AddFileServerHttpServletExtensionRequest) GetRequireAuthentication() bool {
	if o == nil || isNil(o.RequireAuthentication) {
		var ret bool
		return ret
	}
	return *o.RequireAuthentication
}

// GetRequireAuthenticationOk returns a tuple with the RequireAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFileServerHttpServletExtensionRequest) GetRequireAuthenticationOk() (*bool, bool) {
	if o == nil || isNil(o.RequireAuthentication) {
    return nil, false
	}
	return o.RequireAuthentication, true
}

// HasRequireAuthentication returns a boolean if a field has been set.
func (o *AddFileServerHttpServletExtensionRequest) HasRequireAuthentication() bool {
	if o != nil && !isNil(o.RequireAuthentication) {
		return true
	}

	return false
}

// SetRequireAuthentication gets a reference to the given bool and assigns it to the RequireAuthentication field.
func (o *AddFileServerHttpServletExtensionRequest) SetRequireAuthentication(v bool) {
	o.RequireAuthentication = &v
}

// GetAllowedAuthenticationType returns the AllowedAuthenticationType field value if set, zero value otherwise.
func (o *AddFileServerHttpServletExtensionRequest) GetAllowedAuthenticationType() []EnumhttpServletExtensionAllowedAuthenticationTypeProp {
	if o == nil || isNil(o.AllowedAuthenticationType) {
		var ret []EnumhttpServletExtensionAllowedAuthenticationTypeProp
		return ret
	}
	return o.AllowedAuthenticationType
}

// GetAllowedAuthenticationTypeOk returns a tuple with the AllowedAuthenticationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFileServerHttpServletExtensionRequest) GetAllowedAuthenticationTypeOk() ([]EnumhttpServletExtensionAllowedAuthenticationTypeProp, bool) {
	if o == nil || isNil(o.AllowedAuthenticationType) {
    return nil, false
	}
	return o.AllowedAuthenticationType, true
}

// HasAllowedAuthenticationType returns a boolean if a field has been set.
func (o *AddFileServerHttpServletExtensionRequest) HasAllowedAuthenticationType() bool {
	if o != nil && !isNil(o.AllowedAuthenticationType) {
		return true
	}

	return false
}

// SetAllowedAuthenticationType gets a reference to the given []EnumhttpServletExtensionAllowedAuthenticationTypeProp and assigns it to the AllowedAuthenticationType field.
func (o *AddFileServerHttpServletExtensionRequest) SetAllowedAuthenticationType(v []EnumhttpServletExtensionAllowedAuthenticationTypeProp) {
	o.AllowedAuthenticationType = v
}

// GetAccessTokenValidator returns the AccessTokenValidator field value if set, zero value otherwise.
func (o *AddFileServerHttpServletExtensionRequest) GetAccessTokenValidator() []string {
	if o == nil || isNil(o.AccessTokenValidator) {
		var ret []string
		return ret
	}
	return o.AccessTokenValidator
}

// GetAccessTokenValidatorOk returns a tuple with the AccessTokenValidator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFileServerHttpServletExtensionRequest) GetAccessTokenValidatorOk() ([]string, bool) {
	if o == nil || isNil(o.AccessTokenValidator) {
    return nil, false
	}
	return o.AccessTokenValidator, true
}

// HasAccessTokenValidator returns a boolean if a field has been set.
func (o *AddFileServerHttpServletExtensionRequest) HasAccessTokenValidator() bool {
	if o != nil && !isNil(o.AccessTokenValidator) {
		return true
	}

	return false
}

// SetAccessTokenValidator gets a reference to the given []string and assigns it to the AccessTokenValidator field.
func (o *AddFileServerHttpServletExtensionRequest) SetAccessTokenValidator(v []string) {
	o.AccessTokenValidator = v
}

// GetIdTokenValidator returns the IdTokenValidator field value if set, zero value otherwise.
func (o *AddFileServerHttpServletExtensionRequest) GetIdTokenValidator() []string {
	if o == nil || isNil(o.IdTokenValidator) {
		var ret []string
		return ret
	}
	return o.IdTokenValidator
}

// GetIdTokenValidatorOk returns a tuple with the IdTokenValidator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFileServerHttpServletExtensionRequest) GetIdTokenValidatorOk() ([]string, bool) {
	if o == nil || isNil(o.IdTokenValidator) {
    return nil, false
	}
	return o.IdTokenValidator, true
}

// HasIdTokenValidator returns a boolean if a field has been set.
func (o *AddFileServerHttpServletExtensionRequest) HasIdTokenValidator() bool {
	if o != nil && !isNil(o.IdTokenValidator) {
		return true
	}

	return false
}

// SetIdTokenValidator gets a reference to the given []string and assigns it to the IdTokenValidator field.
func (o *AddFileServerHttpServletExtensionRequest) SetIdTokenValidator(v []string) {
	o.IdTokenValidator = v
}

// GetRequireFileServletAccessPrivilege returns the RequireFileServletAccessPrivilege field value if set, zero value otherwise.
func (o *AddFileServerHttpServletExtensionRequest) GetRequireFileServletAccessPrivilege() bool {
	if o == nil || isNil(o.RequireFileServletAccessPrivilege) {
		var ret bool
		return ret
	}
	return *o.RequireFileServletAccessPrivilege
}

// GetRequireFileServletAccessPrivilegeOk returns a tuple with the RequireFileServletAccessPrivilege field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFileServerHttpServletExtensionRequest) GetRequireFileServletAccessPrivilegeOk() (*bool, bool) {
	if o == nil || isNil(o.RequireFileServletAccessPrivilege) {
    return nil, false
	}
	return o.RequireFileServletAccessPrivilege, true
}

// HasRequireFileServletAccessPrivilege returns a boolean if a field has been set.
func (o *AddFileServerHttpServletExtensionRequest) HasRequireFileServletAccessPrivilege() bool {
	if o != nil && !isNil(o.RequireFileServletAccessPrivilege) {
		return true
	}

	return false
}

// SetRequireFileServletAccessPrivilege gets a reference to the given bool and assigns it to the RequireFileServletAccessPrivilege field.
func (o *AddFileServerHttpServletExtensionRequest) SetRequireFileServletAccessPrivilege(v bool) {
	o.RequireFileServletAccessPrivilege = &v
}

// GetRequireGroup returns the RequireGroup field value if set, zero value otherwise.
func (o *AddFileServerHttpServletExtensionRequest) GetRequireGroup() []string {
	if o == nil || isNil(o.RequireGroup) {
		var ret []string
		return ret
	}
	return o.RequireGroup
}

// GetRequireGroupOk returns a tuple with the RequireGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFileServerHttpServletExtensionRequest) GetRequireGroupOk() ([]string, bool) {
	if o == nil || isNil(o.RequireGroup) {
    return nil, false
	}
	return o.RequireGroup, true
}

// HasRequireGroup returns a boolean if a field has been set.
func (o *AddFileServerHttpServletExtensionRequest) HasRequireGroup() bool {
	if o != nil && !isNil(o.RequireGroup) {
		return true
	}

	return false
}

// SetRequireGroup gets a reference to the given []string and assigns it to the RequireGroup field.
func (o *AddFileServerHttpServletExtensionRequest) SetRequireGroup(v []string) {
	o.RequireGroup = v
}

// GetIdentityMapper returns the IdentityMapper field value if set, zero value otherwise.
func (o *AddFileServerHttpServletExtensionRequest) GetIdentityMapper() string {
	if o == nil || isNil(o.IdentityMapper) {
		var ret string
		return ret
	}
	return *o.IdentityMapper
}

// GetIdentityMapperOk returns a tuple with the IdentityMapper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFileServerHttpServletExtensionRequest) GetIdentityMapperOk() (*string, bool) {
	if o == nil || isNil(o.IdentityMapper) {
    return nil, false
	}
	return o.IdentityMapper, true
}

// HasIdentityMapper returns a boolean if a field has been set.
func (o *AddFileServerHttpServletExtensionRequest) HasIdentityMapper() bool {
	if o != nil && !isNil(o.IdentityMapper) {
		return true
	}

	return false
}

// SetIdentityMapper gets a reference to the given string and assigns it to the IdentityMapper field.
func (o *AddFileServerHttpServletExtensionRequest) SetIdentityMapper(v string) {
	o.IdentityMapper = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddFileServerHttpServletExtensionRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFileServerHttpServletExtensionRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddFileServerHttpServletExtensionRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddFileServerHttpServletExtensionRequest) SetDescription(v string) {
	o.Description = &v
}

// GetCrossOriginPolicy returns the CrossOriginPolicy field value if set, zero value otherwise.
func (o *AddFileServerHttpServletExtensionRequest) GetCrossOriginPolicy() string {
	if o == nil || isNil(o.CrossOriginPolicy) {
		var ret string
		return ret
	}
	return *o.CrossOriginPolicy
}

// GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFileServerHttpServletExtensionRequest) GetCrossOriginPolicyOk() (*string, bool) {
	if o == nil || isNil(o.CrossOriginPolicy) {
    return nil, false
	}
	return o.CrossOriginPolicy, true
}

// HasCrossOriginPolicy returns a boolean if a field has been set.
func (o *AddFileServerHttpServletExtensionRequest) HasCrossOriginPolicy() bool {
	if o != nil && !isNil(o.CrossOriginPolicy) {
		return true
	}

	return false
}

// SetCrossOriginPolicy gets a reference to the given string and assigns it to the CrossOriginPolicy field.
func (o *AddFileServerHttpServletExtensionRequest) SetCrossOriginPolicy(v string) {
	o.CrossOriginPolicy = &v
}

// GetResponseHeader returns the ResponseHeader field value if set, zero value otherwise.
func (o *AddFileServerHttpServletExtensionRequest) GetResponseHeader() []string {
	if o == nil || isNil(o.ResponseHeader) {
		var ret []string
		return ret
	}
	return o.ResponseHeader
}

// GetResponseHeaderOk returns a tuple with the ResponseHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFileServerHttpServletExtensionRequest) GetResponseHeaderOk() ([]string, bool) {
	if o == nil || isNil(o.ResponseHeader) {
    return nil, false
	}
	return o.ResponseHeader, true
}

// HasResponseHeader returns a boolean if a field has been set.
func (o *AddFileServerHttpServletExtensionRequest) HasResponseHeader() bool {
	if o != nil && !isNil(o.ResponseHeader) {
		return true
	}

	return false
}

// SetResponseHeader gets a reference to the given []string and assigns it to the ResponseHeader field.
func (o *AddFileServerHttpServletExtensionRequest) SetResponseHeader(v []string) {
	o.ResponseHeader = v
}

// GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field value if set, zero value otherwise.
func (o *AddFileServerHttpServletExtensionRequest) GetCorrelationIDResponseHeader() string {
	if o == nil || isNil(o.CorrelationIDResponseHeader) {
		var ret string
		return ret
	}
	return *o.CorrelationIDResponseHeader
}

// GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFileServerHttpServletExtensionRequest) GetCorrelationIDResponseHeaderOk() (*string, bool) {
	if o == nil || isNil(o.CorrelationIDResponseHeader) {
    return nil, false
	}
	return o.CorrelationIDResponseHeader, true
}

// HasCorrelationIDResponseHeader returns a boolean if a field has been set.
func (o *AddFileServerHttpServletExtensionRequest) HasCorrelationIDResponseHeader() bool {
	if o != nil && !isNil(o.CorrelationIDResponseHeader) {
		return true
	}

	return false
}

// SetCorrelationIDResponseHeader gets a reference to the given string and assigns it to the CorrelationIDResponseHeader field.
func (o *AddFileServerHttpServletExtensionRequest) SetCorrelationIDResponseHeader(v string) {
	o.CorrelationIDResponseHeader = &v
}

func (o AddFileServerHttpServletExtensionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["extensionName"] = o.ExtensionName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["baseContextPath"] = o.BaseContextPath
	}
	if true {
		toSerialize["documentRootDirectory"] = o.DocumentRootDirectory
	}
	if !isNil(o.EnableDirectoryIndexing) {
		toSerialize["enableDirectoryIndexing"] = o.EnableDirectoryIndexing
	}
	if !isNil(o.IndexFile) {
		toSerialize["indexFile"] = o.IndexFile
	}
	if !isNil(o.MimeTypesFile) {
		toSerialize["mimeTypesFile"] = o.MimeTypesFile
	}
	if !isNil(o.DefaultMIMEType) {
		toSerialize["defaultMIMEType"] = o.DefaultMIMEType
	}
	if !isNil(o.RequireAuthentication) {
		toSerialize["requireAuthentication"] = o.RequireAuthentication
	}
	if !isNil(o.AllowedAuthenticationType) {
		toSerialize["allowedAuthenticationType"] = o.AllowedAuthenticationType
	}
	if !isNil(o.AccessTokenValidator) {
		toSerialize["accessTokenValidator"] = o.AccessTokenValidator
	}
	if !isNil(o.IdTokenValidator) {
		toSerialize["idTokenValidator"] = o.IdTokenValidator
	}
	if !isNil(o.RequireFileServletAccessPrivilege) {
		toSerialize["requireFileServletAccessPrivilege"] = o.RequireFileServletAccessPrivilege
	}
	if !isNil(o.RequireGroup) {
		toSerialize["requireGroup"] = o.RequireGroup
	}
	if !isNil(o.IdentityMapper) {
		toSerialize["identityMapper"] = o.IdentityMapper
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.CrossOriginPolicy) {
		toSerialize["crossOriginPolicy"] = o.CrossOriginPolicy
	}
	if !isNil(o.ResponseHeader) {
		toSerialize["responseHeader"] = o.ResponseHeader
	}
	if !isNil(o.CorrelationIDResponseHeader) {
		toSerialize["correlationIDResponseHeader"] = o.CorrelationIDResponseHeader
	}
	return json.Marshal(toSerialize)
}

type NullableAddFileServerHttpServletExtensionRequest struct {
	value *AddFileServerHttpServletExtensionRequest
	isSet bool
}

func (v NullableAddFileServerHttpServletExtensionRequest) Get() *AddFileServerHttpServletExtensionRequest {
	return v.value
}

func (v *NullableAddFileServerHttpServletExtensionRequest) Set(val *AddFileServerHttpServletExtensionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddFileServerHttpServletExtensionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddFileServerHttpServletExtensionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddFileServerHttpServletExtensionRequest(val *AddFileServerHttpServletExtensionRequest) *NullableAddFileServerHttpServletExtensionRequest {
	return &NullableAddFileServerHttpServletExtensionRequest{value: val, isSet: true}
}

func (v NullableAddFileServerHttpServletExtensionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddFileServerHttpServletExtensionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


