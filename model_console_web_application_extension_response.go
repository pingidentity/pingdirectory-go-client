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

// ConsoleWebApplicationExtensionResponse struct for ConsoleWebApplicationExtensionResponse
type ConsoleWebApplicationExtensionResponse struct {
	Schemas []EnumconsoleWebApplicationExtensionSchemaUrn `json:"schemas,omitempty"`
	// Indicates that SSO login into the Administrative Console is enabled.
	SsoEnabled *bool `json:"ssoEnabled,omitempty"`
	// The client ID to use when authenticating to the OpenID Connect provider.
	OidcClientID *string `json:"oidcClientID,omitempty"`
	// The client secret to use when authenticating to the OpenID Connect provider.
	OidcClientSecret *string `json:"oidcClientSecret,omitempty"`
	// A passphrase provider that may be used to obtain the client secret to use when authenticating to the OpenID Connect provider.
	OidcClientSecretPassphraseProvider *string `json:"oidcClientSecretPassphraseProvider,omitempty"`
	// The issuer URL of the OpenID Connect provider.
	OidcIssuerURL *string `json:"oidcIssuerURL,omitempty"`
	// Specifies the path to the truststore file used by this application to evaluate OIDC provider certificates. If this field is left blank, the default JVM trust store will be used.
	OidcTrustStoreFile *string `json:"oidcTrustStoreFile,omitempty"`
	// Specifies the format for the data in the OIDC trust store file.
	OidcTrustStoreType *string `json:"oidcTrustStoreType,omitempty"`
	// The passphrase provider that may be used to obtain the PIN for the trust store used with OIDC providers. This is only required if a trust store file is required, and if that trust store requires a PIN to access its contents.
	OidcTrustStorePinPassphraseProvider *string `json:"oidcTrustStorePinPassphraseProvider,omitempty"`
	// Controls whether or not hostname verification is performed, which checks if the hostname of the OIDC provider matches the name(s) stored inside the certificate it provides. This property should only be set to false for testing purposes.
	OidcStrictHostnameVerification *bool `json:"oidcStrictHostnameVerification,omitempty"`
	// Controls whether or not this application will always trust any certificate that is presented to it, regardless of its contents. This property should only be set to true for testing purposes.
	OidcTrustAll *bool `json:"oidcTrustAll,omitempty"`
	// The LDAP URL used to connect to the managed server.
	LdapServer *string `json:"ldapServer,omitempty"`
	// Specifies the path to the truststore file, which is used by this application to establish trust of managed servers.
	TrustStoreFile *string `json:"trustStoreFile,omitempty"`
	// Specifies the format for the data in the trust store file.
	TrustStoreType *string `json:"trustStoreType,omitempty"`
	// The passphrase provider that may be used to obtain the PIN for the trust store used with managed LDAP servers. This is only required if a trust store file is required, and if that trust store requires a PIN to access its contents.
	TrustStorePinPassphraseProvider *string `json:"trustStorePinPassphraseProvider,omitempty"`
	// The path to the log file for the web application.
	LogFile *string `json:"logFile,omitempty"`
	Complexity *EnumwebApplicationExtensionComplexityProp `json:"complexity,omitempty"`
	// A description for this Web Application Extension
	Description *string `json:"description,omitempty"`
	// Specifies the base context path that should be used by HTTP clients to reference content. The value must start with a forward slash and at least one additional character and must represent a valid HTTP context path.
	BaseContextPath string `json:"baseContextPath"`
	// Specifies the path to a standard web application archive (WAR) file.
	WarFile *string `json:"warFile,omitempty"`
	// Specifies the path to the directory on the local filesystem containing the files to be served by this Web Application Extension. The path must exist, and it must be a directory.
	DocumentRootDirectory *string `json:"documentRootDirectory,omitempty"`
	// Specifies the path to the deployment descriptor file when used with document-root-directory.
	DeploymentDescriptorFile *string `json:"deploymentDescriptorFile,omitempty"`
	// Specifies the path to the directory that may be used to store temporary files such as extracted WAR files and compiled JSP files.
	TemporaryDirectory *string `json:"temporaryDirectory,omitempty"`
	InitParameter []string `json:"initParameter,omitempty"`
}

// NewConsoleWebApplicationExtensionResponse instantiates a new ConsoleWebApplicationExtensionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsoleWebApplicationExtensionResponse(baseContextPath string) *ConsoleWebApplicationExtensionResponse {
	this := ConsoleWebApplicationExtensionResponse{}
	this.BaseContextPath = baseContextPath
	return &this
}

// NewConsoleWebApplicationExtensionResponseWithDefaults instantiates a new ConsoleWebApplicationExtensionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsoleWebApplicationExtensionResponseWithDefaults() *ConsoleWebApplicationExtensionResponse {
	this := ConsoleWebApplicationExtensionResponse{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *ConsoleWebApplicationExtensionResponse) GetSchemas() []EnumconsoleWebApplicationExtensionSchemaUrn {
	if o == nil || isNil(o.Schemas) {
		var ret []EnumconsoleWebApplicationExtensionSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleWebApplicationExtensionResponse) GetSchemasOk() ([]EnumconsoleWebApplicationExtensionSchemaUrn, bool) {
	if o == nil || isNil(o.Schemas) {
    return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *ConsoleWebApplicationExtensionResponse) HasSchemas() bool {
	if o != nil && !isNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumconsoleWebApplicationExtensionSchemaUrn and assigns it to the Schemas field.
func (o *ConsoleWebApplicationExtensionResponse) SetSchemas(v []EnumconsoleWebApplicationExtensionSchemaUrn) {
	o.Schemas = v
}

// GetSsoEnabled returns the SsoEnabled field value if set, zero value otherwise.
func (o *ConsoleWebApplicationExtensionResponse) GetSsoEnabled() bool {
	if o == nil || isNil(o.SsoEnabled) {
		var ret bool
		return ret
	}
	return *o.SsoEnabled
}

// GetSsoEnabledOk returns a tuple with the SsoEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleWebApplicationExtensionResponse) GetSsoEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.SsoEnabled) {
    return nil, false
	}
	return o.SsoEnabled, true
}

// HasSsoEnabled returns a boolean if a field has been set.
func (o *ConsoleWebApplicationExtensionResponse) HasSsoEnabled() bool {
	if o != nil && !isNil(o.SsoEnabled) {
		return true
	}

	return false
}

// SetSsoEnabled gets a reference to the given bool and assigns it to the SsoEnabled field.
func (o *ConsoleWebApplicationExtensionResponse) SetSsoEnabled(v bool) {
	o.SsoEnabled = &v
}

// GetOidcClientID returns the OidcClientID field value if set, zero value otherwise.
func (o *ConsoleWebApplicationExtensionResponse) GetOidcClientID() string {
	if o == nil || isNil(o.OidcClientID) {
		var ret string
		return ret
	}
	return *o.OidcClientID
}

// GetOidcClientIDOk returns a tuple with the OidcClientID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleWebApplicationExtensionResponse) GetOidcClientIDOk() (*string, bool) {
	if o == nil || isNil(o.OidcClientID) {
    return nil, false
	}
	return o.OidcClientID, true
}

// HasOidcClientID returns a boolean if a field has been set.
func (o *ConsoleWebApplicationExtensionResponse) HasOidcClientID() bool {
	if o != nil && !isNil(o.OidcClientID) {
		return true
	}

	return false
}

// SetOidcClientID gets a reference to the given string and assigns it to the OidcClientID field.
func (o *ConsoleWebApplicationExtensionResponse) SetOidcClientID(v string) {
	o.OidcClientID = &v
}

// GetOidcClientSecret returns the OidcClientSecret field value if set, zero value otherwise.
func (o *ConsoleWebApplicationExtensionResponse) GetOidcClientSecret() string {
	if o == nil || isNil(o.OidcClientSecret) {
		var ret string
		return ret
	}
	return *o.OidcClientSecret
}

// GetOidcClientSecretOk returns a tuple with the OidcClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleWebApplicationExtensionResponse) GetOidcClientSecretOk() (*string, bool) {
	if o == nil || isNil(o.OidcClientSecret) {
    return nil, false
	}
	return o.OidcClientSecret, true
}

// HasOidcClientSecret returns a boolean if a field has been set.
func (o *ConsoleWebApplicationExtensionResponse) HasOidcClientSecret() bool {
	if o != nil && !isNil(o.OidcClientSecret) {
		return true
	}

	return false
}

// SetOidcClientSecret gets a reference to the given string and assigns it to the OidcClientSecret field.
func (o *ConsoleWebApplicationExtensionResponse) SetOidcClientSecret(v string) {
	o.OidcClientSecret = &v
}

// GetOidcClientSecretPassphraseProvider returns the OidcClientSecretPassphraseProvider field value if set, zero value otherwise.
func (o *ConsoleWebApplicationExtensionResponse) GetOidcClientSecretPassphraseProvider() string {
	if o == nil || isNil(o.OidcClientSecretPassphraseProvider) {
		var ret string
		return ret
	}
	return *o.OidcClientSecretPassphraseProvider
}

// GetOidcClientSecretPassphraseProviderOk returns a tuple with the OidcClientSecretPassphraseProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleWebApplicationExtensionResponse) GetOidcClientSecretPassphraseProviderOk() (*string, bool) {
	if o == nil || isNil(o.OidcClientSecretPassphraseProvider) {
    return nil, false
	}
	return o.OidcClientSecretPassphraseProvider, true
}

// HasOidcClientSecretPassphraseProvider returns a boolean if a field has been set.
func (o *ConsoleWebApplicationExtensionResponse) HasOidcClientSecretPassphraseProvider() bool {
	if o != nil && !isNil(o.OidcClientSecretPassphraseProvider) {
		return true
	}

	return false
}

// SetOidcClientSecretPassphraseProvider gets a reference to the given string and assigns it to the OidcClientSecretPassphraseProvider field.
func (o *ConsoleWebApplicationExtensionResponse) SetOidcClientSecretPassphraseProvider(v string) {
	o.OidcClientSecretPassphraseProvider = &v
}

// GetOidcIssuerURL returns the OidcIssuerURL field value if set, zero value otherwise.
func (o *ConsoleWebApplicationExtensionResponse) GetOidcIssuerURL() string {
	if o == nil || isNil(o.OidcIssuerURL) {
		var ret string
		return ret
	}
	return *o.OidcIssuerURL
}

// GetOidcIssuerURLOk returns a tuple with the OidcIssuerURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleWebApplicationExtensionResponse) GetOidcIssuerURLOk() (*string, bool) {
	if o == nil || isNil(o.OidcIssuerURL) {
    return nil, false
	}
	return o.OidcIssuerURL, true
}

// HasOidcIssuerURL returns a boolean if a field has been set.
func (o *ConsoleWebApplicationExtensionResponse) HasOidcIssuerURL() bool {
	if o != nil && !isNil(o.OidcIssuerURL) {
		return true
	}

	return false
}

// SetOidcIssuerURL gets a reference to the given string and assigns it to the OidcIssuerURL field.
func (o *ConsoleWebApplicationExtensionResponse) SetOidcIssuerURL(v string) {
	o.OidcIssuerURL = &v
}

// GetOidcTrustStoreFile returns the OidcTrustStoreFile field value if set, zero value otherwise.
func (o *ConsoleWebApplicationExtensionResponse) GetOidcTrustStoreFile() string {
	if o == nil || isNil(o.OidcTrustStoreFile) {
		var ret string
		return ret
	}
	return *o.OidcTrustStoreFile
}

// GetOidcTrustStoreFileOk returns a tuple with the OidcTrustStoreFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleWebApplicationExtensionResponse) GetOidcTrustStoreFileOk() (*string, bool) {
	if o == nil || isNil(o.OidcTrustStoreFile) {
    return nil, false
	}
	return o.OidcTrustStoreFile, true
}

// HasOidcTrustStoreFile returns a boolean if a field has been set.
func (o *ConsoleWebApplicationExtensionResponse) HasOidcTrustStoreFile() bool {
	if o != nil && !isNil(o.OidcTrustStoreFile) {
		return true
	}

	return false
}

// SetOidcTrustStoreFile gets a reference to the given string and assigns it to the OidcTrustStoreFile field.
func (o *ConsoleWebApplicationExtensionResponse) SetOidcTrustStoreFile(v string) {
	o.OidcTrustStoreFile = &v
}

// GetOidcTrustStoreType returns the OidcTrustStoreType field value if set, zero value otherwise.
func (o *ConsoleWebApplicationExtensionResponse) GetOidcTrustStoreType() string {
	if o == nil || isNil(o.OidcTrustStoreType) {
		var ret string
		return ret
	}
	return *o.OidcTrustStoreType
}

// GetOidcTrustStoreTypeOk returns a tuple with the OidcTrustStoreType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleWebApplicationExtensionResponse) GetOidcTrustStoreTypeOk() (*string, bool) {
	if o == nil || isNil(o.OidcTrustStoreType) {
    return nil, false
	}
	return o.OidcTrustStoreType, true
}

// HasOidcTrustStoreType returns a boolean if a field has been set.
func (o *ConsoleWebApplicationExtensionResponse) HasOidcTrustStoreType() bool {
	if o != nil && !isNil(o.OidcTrustStoreType) {
		return true
	}

	return false
}

// SetOidcTrustStoreType gets a reference to the given string and assigns it to the OidcTrustStoreType field.
func (o *ConsoleWebApplicationExtensionResponse) SetOidcTrustStoreType(v string) {
	o.OidcTrustStoreType = &v
}

// GetOidcTrustStorePinPassphraseProvider returns the OidcTrustStorePinPassphraseProvider field value if set, zero value otherwise.
func (o *ConsoleWebApplicationExtensionResponse) GetOidcTrustStorePinPassphraseProvider() string {
	if o == nil || isNil(o.OidcTrustStorePinPassphraseProvider) {
		var ret string
		return ret
	}
	return *o.OidcTrustStorePinPassphraseProvider
}

// GetOidcTrustStorePinPassphraseProviderOk returns a tuple with the OidcTrustStorePinPassphraseProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleWebApplicationExtensionResponse) GetOidcTrustStorePinPassphraseProviderOk() (*string, bool) {
	if o == nil || isNil(o.OidcTrustStorePinPassphraseProvider) {
    return nil, false
	}
	return o.OidcTrustStorePinPassphraseProvider, true
}

// HasOidcTrustStorePinPassphraseProvider returns a boolean if a field has been set.
func (o *ConsoleWebApplicationExtensionResponse) HasOidcTrustStorePinPassphraseProvider() bool {
	if o != nil && !isNil(o.OidcTrustStorePinPassphraseProvider) {
		return true
	}

	return false
}

// SetOidcTrustStorePinPassphraseProvider gets a reference to the given string and assigns it to the OidcTrustStorePinPassphraseProvider field.
func (o *ConsoleWebApplicationExtensionResponse) SetOidcTrustStorePinPassphraseProvider(v string) {
	o.OidcTrustStorePinPassphraseProvider = &v
}

// GetOidcStrictHostnameVerification returns the OidcStrictHostnameVerification field value if set, zero value otherwise.
func (o *ConsoleWebApplicationExtensionResponse) GetOidcStrictHostnameVerification() bool {
	if o == nil || isNil(o.OidcStrictHostnameVerification) {
		var ret bool
		return ret
	}
	return *o.OidcStrictHostnameVerification
}

// GetOidcStrictHostnameVerificationOk returns a tuple with the OidcStrictHostnameVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleWebApplicationExtensionResponse) GetOidcStrictHostnameVerificationOk() (*bool, bool) {
	if o == nil || isNil(o.OidcStrictHostnameVerification) {
    return nil, false
	}
	return o.OidcStrictHostnameVerification, true
}

// HasOidcStrictHostnameVerification returns a boolean if a field has been set.
func (o *ConsoleWebApplicationExtensionResponse) HasOidcStrictHostnameVerification() bool {
	if o != nil && !isNil(o.OidcStrictHostnameVerification) {
		return true
	}

	return false
}

// SetOidcStrictHostnameVerification gets a reference to the given bool and assigns it to the OidcStrictHostnameVerification field.
func (o *ConsoleWebApplicationExtensionResponse) SetOidcStrictHostnameVerification(v bool) {
	o.OidcStrictHostnameVerification = &v
}

// GetOidcTrustAll returns the OidcTrustAll field value if set, zero value otherwise.
func (o *ConsoleWebApplicationExtensionResponse) GetOidcTrustAll() bool {
	if o == nil || isNil(o.OidcTrustAll) {
		var ret bool
		return ret
	}
	return *o.OidcTrustAll
}

// GetOidcTrustAllOk returns a tuple with the OidcTrustAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleWebApplicationExtensionResponse) GetOidcTrustAllOk() (*bool, bool) {
	if o == nil || isNil(o.OidcTrustAll) {
    return nil, false
	}
	return o.OidcTrustAll, true
}

// HasOidcTrustAll returns a boolean if a field has been set.
func (o *ConsoleWebApplicationExtensionResponse) HasOidcTrustAll() bool {
	if o != nil && !isNil(o.OidcTrustAll) {
		return true
	}

	return false
}

// SetOidcTrustAll gets a reference to the given bool and assigns it to the OidcTrustAll field.
func (o *ConsoleWebApplicationExtensionResponse) SetOidcTrustAll(v bool) {
	o.OidcTrustAll = &v
}

// GetLdapServer returns the LdapServer field value if set, zero value otherwise.
func (o *ConsoleWebApplicationExtensionResponse) GetLdapServer() string {
	if o == nil || isNil(o.LdapServer) {
		var ret string
		return ret
	}
	return *o.LdapServer
}

// GetLdapServerOk returns a tuple with the LdapServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleWebApplicationExtensionResponse) GetLdapServerOk() (*string, bool) {
	if o == nil || isNil(o.LdapServer) {
    return nil, false
	}
	return o.LdapServer, true
}

// HasLdapServer returns a boolean if a field has been set.
func (o *ConsoleWebApplicationExtensionResponse) HasLdapServer() bool {
	if o != nil && !isNil(o.LdapServer) {
		return true
	}

	return false
}

// SetLdapServer gets a reference to the given string and assigns it to the LdapServer field.
func (o *ConsoleWebApplicationExtensionResponse) SetLdapServer(v string) {
	o.LdapServer = &v
}

// GetTrustStoreFile returns the TrustStoreFile field value if set, zero value otherwise.
func (o *ConsoleWebApplicationExtensionResponse) GetTrustStoreFile() string {
	if o == nil || isNil(o.TrustStoreFile) {
		var ret string
		return ret
	}
	return *o.TrustStoreFile
}

// GetTrustStoreFileOk returns a tuple with the TrustStoreFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleWebApplicationExtensionResponse) GetTrustStoreFileOk() (*string, bool) {
	if o == nil || isNil(o.TrustStoreFile) {
    return nil, false
	}
	return o.TrustStoreFile, true
}

// HasTrustStoreFile returns a boolean if a field has been set.
func (o *ConsoleWebApplicationExtensionResponse) HasTrustStoreFile() bool {
	if o != nil && !isNil(o.TrustStoreFile) {
		return true
	}

	return false
}

// SetTrustStoreFile gets a reference to the given string and assigns it to the TrustStoreFile field.
func (o *ConsoleWebApplicationExtensionResponse) SetTrustStoreFile(v string) {
	o.TrustStoreFile = &v
}

// GetTrustStoreType returns the TrustStoreType field value if set, zero value otherwise.
func (o *ConsoleWebApplicationExtensionResponse) GetTrustStoreType() string {
	if o == nil || isNil(o.TrustStoreType) {
		var ret string
		return ret
	}
	return *o.TrustStoreType
}

// GetTrustStoreTypeOk returns a tuple with the TrustStoreType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleWebApplicationExtensionResponse) GetTrustStoreTypeOk() (*string, bool) {
	if o == nil || isNil(o.TrustStoreType) {
    return nil, false
	}
	return o.TrustStoreType, true
}

// HasTrustStoreType returns a boolean if a field has been set.
func (o *ConsoleWebApplicationExtensionResponse) HasTrustStoreType() bool {
	if o != nil && !isNil(o.TrustStoreType) {
		return true
	}

	return false
}

// SetTrustStoreType gets a reference to the given string and assigns it to the TrustStoreType field.
func (o *ConsoleWebApplicationExtensionResponse) SetTrustStoreType(v string) {
	o.TrustStoreType = &v
}

// GetTrustStorePinPassphraseProvider returns the TrustStorePinPassphraseProvider field value if set, zero value otherwise.
func (o *ConsoleWebApplicationExtensionResponse) GetTrustStorePinPassphraseProvider() string {
	if o == nil || isNil(o.TrustStorePinPassphraseProvider) {
		var ret string
		return ret
	}
	return *o.TrustStorePinPassphraseProvider
}

// GetTrustStorePinPassphraseProviderOk returns a tuple with the TrustStorePinPassphraseProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleWebApplicationExtensionResponse) GetTrustStorePinPassphraseProviderOk() (*string, bool) {
	if o == nil || isNil(o.TrustStorePinPassphraseProvider) {
    return nil, false
	}
	return o.TrustStorePinPassphraseProvider, true
}

// HasTrustStorePinPassphraseProvider returns a boolean if a field has been set.
func (o *ConsoleWebApplicationExtensionResponse) HasTrustStorePinPassphraseProvider() bool {
	if o != nil && !isNil(o.TrustStorePinPassphraseProvider) {
		return true
	}

	return false
}

// SetTrustStorePinPassphraseProvider gets a reference to the given string and assigns it to the TrustStorePinPassphraseProvider field.
func (o *ConsoleWebApplicationExtensionResponse) SetTrustStorePinPassphraseProvider(v string) {
	o.TrustStorePinPassphraseProvider = &v
}

// GetLogFile returns the LogFile field value if set, zero value otherwise.
func (o *ConsoleWebApplicationExtensionResponse) GetLogFile() string {
	if o == nil || isNil(o.LogFile) {
		var ret string
		return ret
	}
	return *o.LogFile
}

// GetLogFileOk returns a tuple with the LogFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleWebApplicationExtensionResponse) GetLogFileOk() (*string, bool) {
	if o == nil || isNil(o.LogFile) {
    return nil, false
	}
	return o.LogFile, true
}

// HasLogFile returns a boolean if a field has been set.
func (o *ConsoleWebApplicationExtensionResponse) HasLogFile() bool {
	if o != nil && !isNil(o.LogFile) {
		return true
	}

	return false
}

// SetLogFile gets a reference to the given string and assigns it to the LogFile field.
func (o *ConsoleWebApplicationExtensionResponse) SetLogFile(v string) {
	o.LogFile = &v
}

// GetComplexity returns the Complexity field value if set, zero value otherwise.
func (o *ConsoleWebApplicationExtensionResponse) GetComplexity() EnumwebApplicationExtensionComplexityProp {
	if o == nil || isNil(o.Complexity) {
		var ret EnumwebApplicationExtensionComplexityProp
		return ret
	}
	return *o.Complexity
}

// GetComplexityOk returns a tuple with the Complexity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleWebApplicationExtensionResponse) GetComplexityOk() (*EnumwebApplicationExtensionComplexityProp, bool) {
	if o == nil || isNil(o.Complexity) {
    return nil, false
	}
	return o.Complexity, true
}

// HasComplexity returns a boolean if a field has been set.
func (o *ConsoleWebApplicationExtensionResponse) HasComplexity() bool {
	if o != nil && !isNil(o.Complexity) {
		return true
	}

	return false
}

// SetComplexity gets a reference to the given EnumwebApplicationExtensionComplexityProp and assigns it to the Complexity field.
func (o *ConsoleWebApplicationExtensionResponse) SetComplexity(v EnumwebApplicationExtensionComplexityProp) {
	o.Complexity = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConsoleWebApplicationExtensionResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleWebApplicationExtensionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConsoleWebApplicationExtensionResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConsoleWebApplicationExtensionResponse) SetDescription(v string) {
	o.Description = &v
}

// GetBaseContextPath returns the BaseContextPath field value
func (o *ConsoleWebApplicationExtensionResponse) GetBaseContextPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseContextPath
}

// GetBaseContextPathOk returns a tuple with the BaseContextPath field value
// and a boolean to check if the value has been set.
func (o *ConsoleWebApplicationExtensionResponse) GetBaseContextPathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.BaseContextPath, true
}

// SetBaseContextPath sets field value
func (o *ConsoleWebApplicationExtensionResponse) SetBaseContextPath(v string) {
	o.BaseContextPath = v
}

// GetWarFile returns the WarFile field value if set, zero value otherwise.
func (o *ConsoleWebApplicationExtensionResponse) GetWarFile() string {
	if o == nil || isNil(o.WarFile) {
		var ret string
		return ret
	}
	return *o.WarFile
}

// GetWarFileOk returns a tuple with the WarFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleWebApplicationExtensionResponse) GetWarFileOk() (*string, bool) {
	if o == nil || isNil(o.WarFile) {
    return nil, false
	}
	return o.WarFile, true
}

// HasWarFile returns a boolean if a field has been set.
func (o *ConsoleWebApplicationExtensionResponse) HasWarFile() bool {
	if o != nil && !isNil(o.WarFile) {
		return true
	}

	return false
}

// SetWarFile gets a reference to the given string and assigns it to the WarFile field.
func (o *ConsoleWebApplicationExtensionResponse) SetWarFile(v string) {
	o.WarFile = &v
}

// GetDocumentRootDirectory returns the DocumentRootDirectory field value if set, zero value otherwise.
func (o *ConsoleWebApplicationExtensionResponse) GetDocumentRootDirectory() string {
	if o == nil || isNil(o.DocumentRootDirectory) {
		var ret string
		return ret
	}
	return *o.DocumentRootDirectory
}

// GetDocumentRootDirectoryOk returns a tuple with the DocumentRootDirectory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleWebApplicationExtensionResponse) GetDocumentRootDirectoryOk() (*string, bool) {
	if o == nil || isNil(o.DocumentRootDirectory) {
    return nil, false
	}
	return o.DocumentRootDirectory, true
}

// HasDocumentRootDirectory returns a boolean if a field has been set.
func (o *ConsoleWebApplicationExtensionResponse) HasDocumentRootDirectory() bool {
	if o != nil && !isNil(o.DocumentRootDirectory) {
		return true
	}

	return false
}

// SetDocumentRootDirectory gets a reference to the given string and assigns it to the DocumentRootDirectory field.
func (o *ConsoleWebApplicationExtensionResponse) SetDocumentRootDirectory(v string) {
	o.DocumentRootDirectory = &v
}

// GetDeploymentDescriptorFile returns the DeploymentDescriptorFile field value if set, zero value otherwise.
func (o *ConsoleWebApplicationExtensionResponse) GetDeploymentDescriptorFile() string {
	if o == nil || isNil(o.DeploymentDescriptorFile) {
		var ret string
		return ret
	}
	return *o.DeploymentDescriptorFile
}

// GetDeploymentDescriptorFileOk returns a tuple with the DeploymentDescriptorFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleWebApplicationExtensionResponse) GetDeploymentDescriptorFileOk() (*string, bool) {
	if o == nil || isNil(o.DeploymentDescriptorFile) {
    return nil, false
	}
	return o.DeploymentDescriptorFile, true
}

// HasDeploymentDescriptorFile returns a boolean if a field has been set.
func (o *ConsoleWebApplicationExtensionResponse) HasDeploymentDescriptorFile() bool {
	if o != nil && !isNil(o.DeploymentDescriptorFile) {
		return true
	}

	return false
}

// SetDeploymentDescriptorFile gets a reference to the given string and assigns it to the DeploymentDescriptorFile field.
func (o *ConsoleWebApplicationExtensionResponse) SetDeploymentDescriptorFile(v string) {
	o.DeploymentDescriptorFile = &v
}

// GetTemporaryDirectory returns the TemporaryDirectory field value if set, zero value otherwise.
func (o *ConsoleWebApplicationExtensionResponse) GetTemporaryDirectory() string {
	if o == nil || isNil(o.TemporaryDirectory) {
		var ret string
		return ret
	}
	return *o.TemporaryDirectory
}

// GetTemporaryDirectoryOk returns a tuple with the TemporaryDirectory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleWebApplicationExtensionResponse) GetTemporaryDirectoryOk() (*string, bool) {
	if o == nil || isNil(o.TemporaryDirectory) {
    return nil, false
	}
	return o.TemporaryDirectory, true
}

// HasTemporaryDirectory returns a boolean if a field has been set.
func (o *ConsoleWebApplicationExtensionResponse) HasTemporaryDirectory() bool {
	if o != nil && !isNil(o.TemporaryDirectory) {
		return true
	}

	return false
}

// SetTemporaryDirectory gets a reference to the given string and assigns it to the TemporaryDirectory field.
func (o *ConsoleWebApplicationExtensionResponse) SetTemporaryDirectory(v string) {
	o.TemporaryDirectory = &v
}

// GetInitParameter returns the InitParameter field value if set, zero value otherwise.
func (o *ConsoleWebApplicationExtensionResponse) GetInitParameter() []string {
	if o == nil || isNil(o.InitParameter) {
		var ret []string
		return ret
	}
	return o.InitParameter
}

// GetInitParameterOk returns a tuple with the InitParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleWebApplicationExtensionResponse) GetInitParameterOk() ([]string, bool) {
	if o == nil || isNil(o.InitParameter) {
    return nil, false
	}
	return o.InitParameter, true
}

// HasInitParameter returns a boolean if a field has been set.
func (o *ConsoleWebApplicationExtensionResponse) HasInitParameter() bool {
	if o != nil && !isNil(o.InitParameter) {
		return true
	}

	return false
}

// SetInitParameter gets a reference to the given []string and assigns it to the InitParameter field.
func (o *ConsoleWebApplicationExtensionResponse) SetInitParameter(v []string) {
	o.InitParameter = v
}

func (o ConsoleWebApplicationExtensionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.SsoEnabled) {
		toSerialize["ssoEnabled"] = o.SsoEnabled
	}
	if !isNil(o.OidcClientID) {
		toSerialize["oidcClientID"] = o.OidcClientID
	}
	if !isNil(o.OidcClientSecret) {
		toSerialize["oidcClientSecret"] = o.OidcClientSecret
	}
	if !isNil(o.OidcClientSecretPassphraseProvider) {
		toSerialize["oidcClientSecretPassphraseProvider"] = o.OidcClientSecretPassphraseProvider
	}
	if !isNil(o.OidcIssuerURL) {
		toSerialize["oidcIssuerURL"] = o.OidcIssuerURL
	}
	if !isNil(o.OidcTrustStoreFile) {
		toSerialize["oidcTrustStoreFile"] = o.OidcTrustStoreFile
	}
	if !isNil(o.OidcTrustStoreType) {
		toSerialize["oidcTrustStoreType"] = o.OidcTrustStoreType
	}
	if !isNil(o.OidcTrustStorePinPassphraseProvider) {
		toSerialize["oidcTrustStorePinPassphraseProvider"] = o.OidcTrustStorePinPassphraseProvider
	}
	if !isNil(o.OidcStrictHostnameVerification) {
		toSerialize["oidcStrictHostnameVerification"] = o.OidcStrictHostnameVerification
	}
	if !isNil(o.OidcTrustAll) {
		toSerialize["oidcTrustAll"] = o.OidcTrustAll
	}
	if !isNil(o.LdapServer) {
		toSerialize["ldapServer"] = o.LdapServer
	}
	if !isNil(o.TrustStoreFile) {
		toSerialize["trustStoreFile"] = o.TrustStoreFile
	}
	if !isNil(o.TrustStoreType) {
		toSerialize["trustStoreType"] = o.TrustStoreType
	}
	if !isNil(o.TrustStorePinPassphraseProvider) {
		toSerialize["trustStorePinPassphraseProvider"] = o.TrustStorePinPassphraseProvider
	}
	if !isNil(o.LogFile) {
		toSerialize["logFile"] = o.LogFile
	}
	if !isNil(o.Complexity) {
		toSerialize["complexity"] = o.Complexity
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["baseContextPath"] = o.BaseContextPath
	}
	if !isNil(o.WarFile) {
		toSerialize["warFile"] = o.WarFile
	}
	if !isNil(o.DocumentRootDirectory) {
		toSerialize["documentRootDirectory"] = o.DocumentRootDirectory
	}
	if !isNil(o.DeploymentDescriptorFile) {
		toSerialize["deploymentDescriptorFile"] = o.DeploymentDescriptorFile
	}
	if !isNil(o.TemporaryDirectory) {
		toSerialize["temporaryDirectory"] = o.TemporaryDirectory
	}
	if !isNil(o.InitParameter) {
		toSerialize["initParameter"] = o.InitParameter
	}
	return json.Marshal(toSerialize)
}

type NullableConsoleWebApplicationExtensionResponse struct {
	value *ConsoleWebApplicationExtensionResponse
	isSet bool
}

func (v NullableConsoleWebApplicationExtensionResponse) Get() *ConsoleWebApplicationExtensionResponse {
	return v.value
}

func (v *NullableConsoleWebApplicationExtensionResponse) Set(val *ConsoleWebApplicationExtensionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConsoleWebApplicationExtensionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConsoleWebApplicationExtensionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsoleWebApplicationExtensionResponse(val *ConsoleWebApplicationExtensionResponse) *NullableConsoleWebApplicationExtensionResponse {
	return &NullableConsoleWebApplicationExtensionResponse{value: val, isSet: true}
}

func (v NullableConsoleWebApplicationExtensionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsoleWebApplicationExtensionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


