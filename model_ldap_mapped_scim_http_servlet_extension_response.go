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

// LdapMappedScimHttpServletExtensionResponse struct for LdapMappedScimHttpServletExtensionResponse
type LdapMappedScimHttpServletExtensionResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the HTTP Servlet Extension
	Id      string                                            `json:"id"`
	Schemas []EnumldapMappedScimHttpServletExtensionSchemaUrn `json:"schemas"`
	// Specifies the OAuth Token Handler implementation that should be used to validate OAuth 2.0 bearer tokens when they are included in a SCIM request.
	OAuthTokenHandler *string `json:"OAuthTokenHandler,omitempty"`
	// Enables HTTP Basic authentication, using a username and password.
	BasicAuthEnabled *bool `json:"basicAuthEnabled,omitempty"`
	// Specifies the name of the identity mapper that is to be used to match the username included in the HTTP Basic authentication header to the corresponding user in the directory.
	IdentityMapper *string `json:"identityMapper,omitempty"`
	// The path to an XML file defining the resources supported by the SCIM interface and the SCIM-to-LDAP attribute mappings to use.
	ResourceMappingFile *string `json:"resourceMappingFile,omitempty"`
	// Specifies the LDAP object classes that should be exposed directly as SCIM resources.
	IncludeLDAPObjectclass []string `json:"includeLDAPObjectclass,omitempty"`
	// Specifies the LDAP object classes that should be not be exposed directly as SCIM resources.
	ExcludeLDAPObjectclass []string `json:"excludeLDAPObjectclass,omitempty"`
	// Specifies the base DNs for the branches of the DIT that should be exposed via the Identity Access API.
	IncludeLDAPBaseDN []string `json:"includeLDAPBaseDN,omitempty"`
	// Specifies the base DNs for the branches of the DIT that should not be exposed via the Identity Access API.
	ExcludeLDAPBaseDN []string `json:"excludeLDAPBaseDN,omitempty"`
	// Specifies the LDAP attribute whose value should be used as the entity tag value to enable SCIM resource versioning support.
	EntityTagLDAPAttribute *string `json:"entityTagLDAPAttribute,omitempty"`
	// The context path to use to access the SCIM interface. The value must start with a forward slash and must represent a valid HTTP context path.
	BaseContextPath string `json:"baseContextPath"`
	// Specifies the location of the directory that is used to create temporary files containing SCIM request data.
	TemporaryDirectory string `json:"temporaryDirectory"`
	// Specifies the permissions that should be applied to the directory that is used to create temporary files.
	TemporaryDirectoryPermissions string `json:"temporaryDirectoryPermissions"`
	// The maximum number of resources that are returned in a response.
	MaxResults *int32 `json:"maxResults,omitempty"`
	// The maximum number of operations that are permitted in a bulk request.
	BulkMaxOperations *int32 `json:"bulkMaxOperations,omitempty"`
	// The maximum payload size in bytes of a bulk request.
	BulkMaxPayloadSize *string `json:"bulkMaxPayloadSize,omitempty"`
	// The maximum number of bulk requests that may be processed concurrently by the server. Any bulk request that would cause this limit to be exceeded is rejected with HTTP status code 503.
	BulkMaxConcurrentRequests *int32 `json:"bulkMaxConcurrentRequests,omitempty"`
	// Enables debug logging of the SCIM SDK. Debug messages will be forwarded to the Directory Server debug logger with the scope of com.unboundid.directory.server.extensions.scim.SCIMHTTPServletExtension.
	DebugEnabled *bool                                   `json:"debugEnabled,omitempty"`
	DebugLevel   EnumhttpServletExtensionDebugLevelProp  `json:"debugLevel"`
	DebugType    []EnumhttpServletExtensionDebugTypeProp `json:"debugType"`
	// Indicates whether a stack trace of the thread which called the debug method should be included in debug log messages.
	IncludeStackTrace bool `json:"includeStackTrace"`
	// A description for this HTTP Servlet Extension
	Description *string `json:"description,omitempty"`
	// The cross-origin request policy to use for the HTTP Servlet Extension.
	CrossOriginPolicy *string `json:"crossOriginPolicy,omitempty"`
	// Specifies HTTP header fields and values added to response headers for all requests.
	ResponseHeader []string `json:"responseHeader,omitempty"`
	// Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \"Correlation-Id\", \"X-Amzn-Trace-Id\", and \"X-Request-Id\".
	CorrelationIDResponseHeader *string `json:"correlationIDResponseHeader,omitempty"`
}

// NewLdapMappedScimHttpServletExtensionResponse instantiates a new LdapMappedScimHttpServletExtensionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapMappedScimHttpServletExtensionResponse(id string, schemas []EnumldapMappedScimHttpServletExtensionSchemaUrn, baseContextPath string, temporaryDirectory string, temporaryDirectoryPermissions string, debugLevel EnumhttpServletExtensionDebugLevelProp, debugType []EnumhttpServletExtensionDebugTypeProp, includeStackTrace bool) *LdapMappedScimHttpServletExtensionResponse {
	this := LdapMappedScimHttpServletExtensionResponse{}
	this.Id = id
	this.Schemas = schemas
	this.BaseContextPath = baseContextPath
	this.TemporaryDirectory = temporaryDirectory
	this.TemporaryDirectoryPermissions = temporaryDirectoryPermissions
	this.DebugLevel = debugLevel
	this.DebugType = debugType
	this.IncludeStackTrace = includeStackTrace
	return &this
}

// NewLdapMappedScimHttpServletExtensionResponseWithDefaults instantiates a new LdapMappedScimHttpServletExtensionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapMappedScimHttpServletExtensionResponseWithDefaults() *LdapMappedScimHttpServletExtensionResponse {
	this := LdapMappedScimHttpServletExtensionResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *LdapMappedScimHttpServletExtensionResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *LdapMappedScimHttpServletExtensionResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *LdapMappedScimHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *LdapMappedScimHttpServletExtensionResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *LdapMappedScimHttpServletExtensionResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LdapMappedScimHttpServletExtensionResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *LdapMappedScimHttpServletExtensionResponse) GetSchemas() []EnumldapMappedScimHttpServletExtensionSchemaUrn {
	if o == nil {
		var ret []EnumldapMappedScimHttpServletExtensionSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) GetSchemasOk() ([]EnumldapMappedScimHttpServletExtensionSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *LdapMappedScimHttpServletExtensionResponse) SetSchemas(v []EnumldapMappedScimHttpServletExtensionSchemaUrn) {
	o.Schemas = v
}

// GetOAuthTokenHandler returns the OAuthTokenHandler field value if set, zero value otherwise.
func (o *LdapMappedScimHttpServletExtensionResponse) GetOAuthTokenHandler() string {
	if o == nil || isNil(o.OAuthTokenHandler) {
		var ret string
		return ret
	}
	return *o.OAuthTokenHandler
}

// GetOAuthTokenHandlerOk returns a tuple with the OAuthTokenHandler field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) GetOAuthTokenHandlerOk() (*string, bool) {
	if o == nil || isNil(o.OAuthTokenHandler) {
		return nil, false
	}
	return o.OAuthTokenHandler, true
}

// HasOAuthTokenHandler returns a boolean if a field has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) HasOAuthTokenHandler() bool {
	if o != nil && !isNil(o.OAuthTokenHandler) {
		return true
	}

	return false
}

// SetOAuthTokenHandler gets a reference to the given string and assigns it to the OAuthTokenHandler field.
func (o *LdapMappedScimHttpServletExtensionResponse) SetOAuthTokenHandler(v string) {
	o.OAuthTokenHandler = &v
}

// GetBasicAuthEnabled returns the BasicAuthEnabled field value if set, zero value otherwise.
func (o *LdapMappedScimHttpServletExtensionResponse) GetBasicAuthEnabled() bool {
	if o == nil || isNil(o.BasicAuthEnabled) {
		var ret bool
		return ret
	}
	return *o.BasicAuthEnabled
}

// GetBasicAuthEnabledOk returns a tuple with the BasicAuthEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) GetBasicAuthEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.BasicAuthEnabled) {
		return nil, false
	}
	return o.BasicAuthEnabled, true
}

// HasBasicAuthEnabled returns a boolean if a field has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) HasBasicAuthEnabled() bool {
	if o != nil && !isNil(o.BasicAuthEnabled) {
		return true
	}

	return false
}

// SetBasicAuthEnabled gets a reference to the given bool and assigns it to the BasicAuthEnabled field.
func (o *LdapMappedScimHttpServletExtensionResponse) SetBasicAuthEnabled(v bool) {
	o.BasicAuthEnabled = &v
}

// GetIdentityMapper returns the IdentityMapper field value if set, zero value otherwise.
func (o *LdapMappedScimHttpServletExtensionResponse) GetIdentityMapper() string {
	if o == nil || isNil(o.IdentityMapper) {
		var ret string
		return ret
	}
	return *o.IdentityMapper
}

// GetIdentityMapperOk returns a tuple with the IdentityMapper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) GetIdentityMapperOk() (*string, bool) {
	if o == nil || isNil(o.IdentityMapper) {
		return nil, false
	}
	return o.IdentityMapper, true
}

// HasIdentityMapper returns a boolean if a field has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) HasIdentityMapper() bool {
	if o != nil && !isNil(o.IdentityMapper) {
		return true
	}

	return false
}

// SetIdentityMapper gets a reference to the given string and assigns it to the IdentityMapper field.
func (o *LdapMappedScimHttpServletExtensionResponse) SetIdentityMapper(v string) {
	o.IdentityMapper = &v
}

// GetResourceMappingFile returns the ResourceMappingFile field value if set, zero value otherwise.
func (o *LdapMappedScimHttpServletExtensionResponse) GetResourceMappingFile() string {
	if o == nil || isNil(o.ResourceMappingFile) {
		var ret string
		return ret
	}
	return *o.ResourceMappingFile
}

// GetResourceMappingFileOk returns a tuple with the ResourceMappingFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) GetResourceMappingFileOk() (*string, bool) {
	if o == nil || isNil(o.ResourceMappingFile) {
		return nil, false
	}
	return o.ResourceMappingFile, true
}

// HasResourceMappingFile returns a boolean if a field has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) HasResourceMappingFile() bool {
	if o != nil && !isNil(o.ResourceMappingFile) {
		return true
	}

	return false
}

// SetResourceMappingFile gets a reference to the given string and assigns it to the ResourceMappingFile field.
func (o *LdapMappedScimHttpServletExtensionResponse) SetResourceMappingFile(v string) {
	o.ResourceMappingFile = &v
}

// GetIncludeLDAPObjectclass returns the IncludeLDAPObjectclass field value if set, zero value otherwise.
func (o *LdapMappedScimHttpServletExtensionResponse) GetIncludeLDAPObjectclass() []string {
	if o == nil || isNil(o.IncludeLDAPObjectclass) {
		var ret []string
		return ret
	}
	return o.IncludeLDAPObjectclass
}

// GetIncludeLDAPObjectclassOk returns a tuple with the IncludeLDAPObjectclass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) GetIncludeLDAPObjectclassOk() ([]string, bool) {
	if o == nil || isNil(o.IncludeLDAPObjectclass) {
		return nil, false
	}
	return o.IncludeLDAPObjectclass, true
}

// HasIncludeLDAPObjectclass returns a boolean if a field has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) HasIncludeLDAPObjectclass() bool {
	if o != nil && !isNil(o.IncludeLDAPObjectclass) {
		return true
	}

	return false
}

// SetIncludeLDAPObjectclass gets a reference to the given []string and assigns it to the IncludeLDAPObjectclass field.
func (o *LdapMappedScimHttpServletExtensionResponse) SetIncludeLDAPObjectclass(v []string) {
	o.IncludeLDAPObjectclass = v
}

// GetExcludeLDAPObjectclass returns the ExcludeLDAPObjectclass field value if set, zero value otherwise.
func (o *LdapMappedScimHttpServletExtensionResponse) GetExcludeLDAPObjectclass() []string {
	if o == nil || isNil(o.ExcludeLDAPObjectclass) {
		var ret []string
		return ret
	}
	return o.ExcludeLDAPObjectclass
}

// GetExcludeLDAPObjectclassOk returns a tuple with the ExcludeLDAPObjectclass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) GetExcludeLDAPObjectclassOk() ([]string, bool) {
	if o == nil || isNil(o.ExcludeLDAPObjectclass) {
		return nil, false
	}
	return o.ExcludeLDAPObjectclass, true
}

// HasExcludeLDAPObjectclass returns a boolean if a field has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) HasExcludeLDAPObjectclass() bool {
	if o != nil && !isNil(o.ExcludeLDAPObjectclass) {
		return true
	}

	return false
}

// SetExcludeLDAPObjectclass gets a reference to the given []string and assigns it to the ExcludeLDAPObjectclass field.
func (o *LdapMappedScimHttpServletExtensionResponse) SetExcludeLDAPObjectclass(v []string) {
	o.ExcludeLDAPObjectclass = v
}

// GetIncludeLDAPBaseDN returns the IncludeLDAPBaseDN field value if set, zero value otherwise.
func (o *LdapMappedScimHttpServletExtensionResponse) GetIncludeLDAPBaseDN() []string {
	if o == nil || isNil(o.IncludeLDAPBaseDN) {
		var ret []string
		return ret
	}
	return o.IncludeLDAPBaseDN
}

// GetIncludeLDAPBaseDNOk returns a tuple with the IncludeLDAPBaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) GetIncludeLDAPBaseDNOk() ([]string, bool) {
	if o == nil || isNil(o.IncludeLDAPBaseDN) {
		return nil, false
	}
	return o.IncludeLDAPBaseDN, true
}

// HasIncludeLDAPBaseDN returns a boolean if a field has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) HasIncludeLDAPBaseDN() bool {
	if o != nil && !isNil(o.IncludeLDAPBaseDN) {
		return true
	}

	return false
}

// SetIncludeLDAPBaseDN gets a reference to the given []string and assigns it to the IncludeLDAPBaseDN field.
func (o *LdapMappedScimHttpServletExtensionResponse) SetIncludeLDAPBaseDN(v []string) {
	o.IncludeLDAPBaseDN = v
}

// GetExcludeLDAPBaseDN returns the ExcludeLDAPBaseDN field value if set, zero value otherwise.
func (o *LdapMappedScimHttpServletExtensionResponse) GetExcludeLDAPBaseDN() []string {
	if o == nil || isNil(o.ExcludeLDAPBaseDN) {
		var ret []string
		return ret
	}
	return o.ExcludeLDAPBaseDN
}

// GetExcludeLDAPBaseDNOk returns a tuple with the ExcludeLDAPBaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) GetExcludeLDAPBaseDNOk() ([]string, bool) {
	if o == nil || isNil(o.ExcludeLDAPBaseDN) {
		return nil, false
	}
	return o.ExcludeLDAPBaseDN, true
}

// HasExcludeLDAPBaseDN returns a boolean if a field has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) HasExcludeLDAPBaseDN() bool {
	if o != nil && !isNil(o.ExcludeLDAPBaseDN) {
		return true
	}

	return false
}

// SetExcludeLDAPBaseDN gets a reference to the given []string and assigns it to the ExcludeLDAPBaseDN field.
func (o *LdapMappedScimHttpServletExtensionResponse) SetExcludeLDAPBaseDN(v []string) {
	o.ExcludeLDAPBaseDN = v
}

// GetEntityTagLDAPAttribute returns the EntityTagLDAPAttribute field value if set, zero value otherwise.
func (o *LdapMappedScimHttpServletExtensionResponse) GetEntityTagLDAPAttribute() string {
	if o == nil || isNil(o.EntityTagLDAPAttribute) {
		var ret string
		return ret
	}
	return *o.EntityTagLDAPAttribute
}

// GetEntityTagLDAPAttributeOk returns a tuple with the EntityTagLDAPAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) GetEntityTagLDAPAttributeOk() (*string, bool) {
	if o == nil || isNil(o.EntityTagLDAPAttribute) {
		return nil, false
	}
	return o.EntityTagLDAPAttribute, true
}

// HasEntityTagLDAPAttribute returns a boolean if a field has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) HasEntityTagLDAPAttribute() bool {
	if o != nil && !isNil(o.EntityTagLDAPAttribute) {
		return true
	}

	return false
}

// SetEntityTagLDAPAttribute gets a reference to the given string and assigns it to the EntityTagLDAPAttribute field.
func (o *LdapMappedScimHttpServletExtensionResponse) SetEntityTagLDAPAttribute(v string) {
	o.EntityTagLDAPAttribute = &v
}

// GetBaseContextPath returns the BaseContextPath field value
func (o *LdapMappedScimHttpServletExtensionResponse) GetBaseContextPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseContextPath
}

// GetBaseContextPathOk returns a tuple with the BaseContextPath field value
// and a boolean to check if the value has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) GetBaseContextPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseContextPath, true
}

// SetBaseContextPath sets field value
func (o *LdapMappedScimHttpServletExtensionResponse) SetBaseContextPath(v string) {
	o.BaseContextPath = v
}

// GetTemporaryDirectory returns the TemporaryDirectory field value
func (o *LdapMappedScimHttpServletExtensionResponse) GetTemporaryDirectory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemporaryDirectory
}

// GetTemporaryDirectoryOk returns a tuple with the TemporaryDirectory field value
// and a boolean to check if the value has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) GetTemporaryDirectoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemporaryDirectory, true
}

// SetTemporaryDirectory sets field value
func (o *LdapMappedScimHttpServletExtensionResponse) SetTemporaryDirectory(v string) {
	o.TemporaryDirectory = v
}

// GetTemporaryDirectoryPermissions returns the TemporaryDirectoryPermissions field value
func (o *LdapMappedScimHttpServletExtensionResponse) GetTemporaryDirectoryPermissions() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemporaryDirectoryPermissions
}

// GetTemporaryDirectoryPermissionsOk returns a tuple with the TemporaryDirectoryPermissions field value
// and a boolean to check if the value has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) GetTemporaryDirectoryPermissionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemporaryDirectoryPermissions, true
}

// SetTemporaryDirectoryPermissions sets field value
func (o *LdapMappedScimHttpServletExtensionResponse) SetTemporaryDirectoryPermissions(v string) {
	o.TemporaryDirectoryPermissions = v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *LdapMappedScimHttpServletExtensionResponse) GetMaxResults() int32 {
	if o == nil || isNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) GetMaxResultsOk() (*int32, bool) {
	if o == nil || isNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) HasMaxResults() bool {
	if o != nil && !isNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *LdapMappedScimHttpServletExtensionResponse) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetBulkMaxOperations returns the BulkMaxOperations field value if set, zero value otherwise.
func (o *LdapMappedScimHttpServletExtensionResponse) GetBulkMaxOperations() int32 {
	if o == nil || isNil(o.BulkMaxOperations) {
		var ret int32
		return ret
	}
	return *o.BulkMaxOperations
}

// GetBulkMaxOperationsOk returns a tuple with the BulkMaxOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) GetBulkMaxOperationsOk() (*int32, bool) {
	if o == nil || isNil(o.BulkMaxOperations) {
		return nil, false
	}
	return o.BulkMaxOperations, true
}

// HasBulkMaxOperations returns a boolean if a field has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) HasBulkMaxOperations() bool {
	if o != nil && !isNil(o.BulkMaxOperations) {
		return true
	}

	return false
}

// SetBulkMaxOperations gets a reference to the given int32 and assigns it to the BulkMaxOperations field.
func (o *LdapMappedScimHttpServletExtensionResponse) SetBulkMaxOperations(v int32) {
	o.BulkMaxOperations = &v
}

// GetBulkMaxPayloadSize returns the BulkMaxPayloadSize field value if set, zero value otherwise.
func (o *LdapMappedScimHttpServletExtensionResponse) GetBulkMaxPayloadSize() string {
	if o == nil || isNil(o.BulkMaxPayloadSize) {
		var ret string
		return ret
	}
	return *o.BulkMaxPayloadSize
}

// GetBulkMaxPayloadSizeOk returns a tuple with the BulkMaxPayloadSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) GetBulkMaxPayloadSizeOk() (*string, bool) {
	if o == nil || isNil(o.BulkMaxPayloadSize) {
		return nil, false
	}
	return o.BulkMaxPayloadSize, true
}

// HasBulkMaxPayloadSize returns a boolean if a field has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) HasBulkMaxPayloadSize() bool {
	if o != nil && !isNil(o.BulkMaxPayloadSize) {
		return true
	}

	return false
}

// SetBulkMaxPayloadSize gets a reference to the given string and assigns it to the BulkMaxPayloadSize field.
func (o *LdapMappedScimHttpServletExtensionResponse) SetBulkMaxPayloadSize(v string) {
	o.BulkMaxPayloadSize = &v
}

// GetBulkMaxConcurrentRequests returns the BulkMaxConcurrentRequests field value if set, zero value otherwise.
func (o *LdapMappedScimHttpServletExtensionResponse) GetBulkMaxConcurrentRequests() int32 {
	if o == nil || isNil(o.BulkMaxConcurrentRequests) {
		var ret int32
		return ret
	}
	return *o.BulkMaxConcurrentRequests
}

// GetBulkMaxConcurrentRequestsOk returns a tuple with the BulkMaxConcurrentRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) GetBulkMaxConcurrentRequestsOk() (*int32, bool) {
	if o == nil || isNil(o.BulkMaxConcurrentRequests) {
		return nil, false
	}
	return o.BulkMaxConcurrentRequests, true
}

// HasBulkMaxConcurrentRequests returns a boolean if a field has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) HasBulkMaxConcurrentRequests() bool {
	if o != nil && !isNil(o.BulkMaxConcurrentRequests) {
		return true
	}

	return false
}

// SetBulkMaxConcurrentRequests gets a reference to the given int32 and assigns it to the BulkMaxConcurrentRequests field.
func (o *LdapMappedScimHttpServletExtensionResponse) SetBulkMaxConcurrentRequests(v int32) {
	o.BulkMaxConcurrentRequests = &v
}

// GetDebugEnabled returns the DebugEnabled field value if set, zero value otherwise.
func (o *LdapMappedScimHttpServletExtensionResponse) GetDebugEnabled() bool {
	if o == nil || isNil(o.DebugEnabled) {
		var ret bool
		return ret
	}
	return *o.DebugEnabled
}

// GetDebugEnabledOk returns a tuple with the DebugEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) GetDebugEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.DebugEnabled) {
		return nil, false
	}
	return o.DebugEnabled, true
}

// HasDebugEnabled returns a boolean if a field has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) HasDebugEnabled() bool {
	if o != nil && !isNil(o.DebugEnabled) {
		return true
	}

	return false
}

// SetDebugEnabled gets a reference to the given bool and assigns it to the DebugEnabled field.
func (o *LdapMappedScimHttpServletExtensionResponse) SetDebugEnabled(v bool) {
	o.DebugEnabled = &v
}

// GetDebugLevel returns the DebugLevel field value
func (o *LdapMappedScimHttpServletExtensionResponse) GetDebugLevel() EnumhttpServletExtensionDebugLevelProp {
	if o == nil {
		var ret EnumhttpServletExtensionDebugLevelProp
		return ret
	}

	return o.DebugLevel
}

// GetDebugLevelOk returns a tuple with the DebugLevel field value
// and a boolean to check if the value has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) GetDebugLevelOk() (*EnumhttpServletExtensionDebugLevelProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DebugLevel, true
}

// SetDebugLevel sets field value
func (o *LdapMappedScimHttpServletExtensionResponse) SetDebugLevel(v EnumhttpServletExtensionDebugLevelProp) {
	o.DebugLevel = v
}

// GetDebugType returns the DebugType field value
func (o *LdapMappedScimHttpServletExtensionResponse) GetDebugType() []EnumhttpServletExtensionDebugTypeProp {
	if o == nil {
		var ret []EnumhttpServletExtensionDebugTypeProp
		return ret
	}

	return o.DebugType
}

// GetDebugTypeOk returns a tuple with the DebugType field value
// and a boolean to check if the value has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) GetDebugTypeOk() ([]EnumhttpServletExtensionDebugTypeProp, bool) {
	if o == nil {
		return nil, false
	}
	return o.DebugType, true
}

// SetDebugType sets field value
func (o *LdapMappedScimHttpServletExtensionResponse) SetDebugType(v []EnumhttpServletExtensionDebugTypeProp) {
	o.DebugType = v
}

// GetIncludeStackTrace returns the IncludeStackTrace field value
func (o *LdapMappedScimHttpServletExtensionResponse) GetIncludeStackTrace() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IncludeStackTrace
}

// GetIncludeStackTraceOk returns a tuple with the IncludeStackTrace field value
// and a boolean to check if the value has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) GetIncludeStackTraceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludeStackTrace, true
}

// SetIncludeStackTrace sets field value
func (o *LdapMappedScimHttpServletExtensionResponse) SetIncludeStackTrace(v bool) {
	o.IncludeStackTrace = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LdapMappedScimHttpServletExtensionResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LdapMappedScimHttpServletExtensionResponse) SetDescription(v string) {
	o.Description = &v
}

// GetCrossOriginPolicy returns the CrossOriginPolicy field value if set, zero value otherwise.
func (o *LdapMappedScimHttpServletExtensionResponse) GetCrossOriginPolicy() string {
	if o == nil || isNil(o.CrossOriginPolicy) {
		var ret string
		return ret
	}
	return *o.CrossOriginPolicy
}

// GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) GetCrossOriginPolicyOk() (*string, bool) {
	if o == nil || isNil(o.CrossOriginPolicy) {
		return nil, false
	}
	return o.CrossOriginPolicy, true
}

// HasCrossOriginPolicy returns a boolean if a field has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) HasCrossOriginPolicy() bool {
	if o != nil && !isNil(o.CrossOriginPolicy) {
		return true
	}

	return false
}

// SetCrossOriginPolicy gets a reference to the given string and assigns it to the CrossOriginPolicy field.
func (o *LdapMappedScimHttpServletExtensionResponse) SetCrossOriginPolicy(v string) {
	o.CrossOriginPolicy = &v
}

// GetResponseHeader returns the ResponseHeader field value if set, zero value otherwise.
func (o *LdapMappedScimHttpServletExtensionResponse) GetResponseHeader() []string {
	if o == nil || isNil(o.ResponseHeader) {
		var ret []string
		return ret
	}
	return o.ResponseHeader
}

// GetResponseHeaderOk returns a tuple with the ResponseHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) GetResponseHeaderOk() ([]string, bool) {
	if o == nil || isNil(o.ResponseHeader) {
		return nil, false
	}
	return o.ResponseHeader, true
}

// HasResponseHeader returns a boolean if a field has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) HasResponseHeader() bool {
	if o != nil && !isNil(o.ResponseHeader) {
		return true
	}

	return false
}

// SetResponseHeader gets a reference to the given []string and assigns it to the ResponseHeader field.
func (o *LdapMappedScimHttpServletExtensionResponse) SetResponseHeader(v []string) {
	o.ResponseHeader = v
}

// GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field value if set, zero value otherwise.
func (o *LdapMappedScimHttpServletExtensionResponse) GetCorrelationIDResponseHeader() string {
	if o == nil || isNil(o.CorrelationIDResponseHeader) {
		var ret string
		return ret
	}
	return *o.CorrelationIDResponseHeader
}

// GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) GetCorrelationIDResponseHeaderOk() (*string, bool) {
	if o == nil || isNil(o.CorrelationIDResponseHeader) {
		return nil, false
	}
	return o.CorrelationIDResponseHeader, true
}

// HasCorrelationIDResponseHeader returns a boolean if a field has been set.
func (o *LdapMappedScimHttpServletExtensionResponse) HasCorrelationIDResponseHeader() bool {
	if o != nil && !isNil(o.CorrelationIDResponseHeader) {
		return true
	}

	return false
}

// SetCorrelationIDResponseHeader gets a reference to the given string and assigns it to the CorrelationIDResponseHeader field.
func (o *LdapMappedScimHttpServletExtensionResponse) SetCorrelationIDResponseHeader(v string) {
	o.CorrelationIDResponseHeader = &v
}

func (o LdapMappedScimHttpServletExtensionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.OAuthTokenHandler) {
		toSerialize["OAuthTokenHandler"] = o.OAuthTokenHandler
	}
	if !isNil(o.BasicAuthEnabled) {
		toSerialize["basicAuthEnabled"] = o.BasicAuthEnabled
	}
	if !isNil(o.IdentityMapper) {
		toSerialize["identityMapper"] = o.IdentityMapper
	}
	if !isNil(o.ResourceMappingFile) {
		toSerialize["resourceMappingFile"] = o.ResourceMappingFile
	}
	if !isNil(o.IncludeLDAPObjectclass) {
		toSerialize["includeLDAPObjectclass"] = o.IncludeLDAPObjectclass
	}
	if !isNil(o.ExcludeLDAPObjectclass) {
		toSerialize["excludeLDAPObjectclass"] = o.ExcludeLDAPObjectclass
	}
	if !isNil(o.IncludeLDAPBaseDN) {
		toSerialize["includeLDAPBaseDN"] = o.IncludeLDAPBaseDN
	}
	if !isNil(o.ExcludeLDAPBaseDN) {
		toSerialize["excludeLDAPBaseDN"] = o.ExcludeLDAPBaseDN
	}
	if !isNil(o.EntityTagLDAPAttribute) {
		toSerialize["entityTagLDAPAttribute"] = o.EntityTagLDAPAttribute
	}
	if true {
		toSerialize["baseContextPath"] = o.BaseContextPath
	}
	if true {
		toSerialize["temporaryDirectory"] = o.TemporaryDirectory
	}
	if true {
		toSerialize["temporaryDirectoryPermissions"] = o.TemporaryDirectoryPermissions
	}
	if !isNil(o.MaxResults) {
		toSerialize["maxResults"] = o.MaxResults
	}
	if !isNil(o.BulkMaxOperations) {
		toSerialize["bulkMaxOperations"] = o.BulkMaxOperations
	}
	if !isNil(o.BulkMaxPayloadSize) {
		toSerialize["bulkMaxPayloadSize"] = o.BulkMaxPayloadSize
	}
	if !isNil(o.BulkMaxConcurrentRequests) {
		toSerialize["bulkMaxConcurrentRequests"] = o.BulkMaxConcurrentRequests
	}
	if !isNil(o.DebugEnabled) {
		toSerialize["debugEnabled"] = o.DebugEnabled
	}
	if true {
		toSerialize["debugLevel"] = o.DebugLevel
	}
	if true {
		toSerialize["debugType"] = o.DebugType
	}
	if true {
		toSerialize["includeStackTrace"] = o.IncludeStackTrace
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

type NullableLdapMappedScimHttpServletExtensionResponse struct {
	value *LdapMappedScimHttpServletExtensionResponse
	isSet bool
}

func (v NullableLdapMappedScimHttpServletExtensionResponse) Get() *LdapMappedScimHttpServletExtensionResponse {
	return v.value
}

func (v *NullableLdapMappedScimHttpServletExtensionResponse) Set(val *LdapMappedScimHttpServletExtensionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapMappedScimHttpServletExtensionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapMappedScimHttpServletExtensionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapMappedScimHttpServletExtensionResponse(val *LdapMappedScimHttpServletExtensionResponse) *NullableLdapMappedScimHttpServletExtensionResponse {
	return &NullableLdapMappedScimHttpServletExtensionResponse{value: val, isSet: true}
}

func (v NullableLdapMappedScimHttpServletExtensionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapMappedScimHttpServletExtensionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
