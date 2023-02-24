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

// checks if the AddPingOnePassThroughAuthenticationPluginRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddPingOnePassThroughAuthenticationPluginRequest{}

// AddPingOnePassThroughAuthenticationPluginRequest struct for AddPingOnePassThroughAuthenticationPluginRequest
type AddPingOnePassThroughAuthenticationPluginRequest struct {
	// Name of the new Plugin
	PluginName string                                                `json:"pluginName"`
	Schemas    []EnumpingOnePassThroughAuthenticationPluginSchemaUrn `json:"schemas"`
	// Specifies the API endpoint for the PingOne web service.
	ApiURL string `json:"apiURL"`
	// Specifies the API endpoint for the PingOne authentication service.
	AuthURL string `json:"authURL"`
	// Specifies the OAuth Client ID used to authenticate connections to the PingOne API.
	OAuthClientID string `json:"OAuthClientID"`
	// Specifies the OAuth Client Secret used to authenticate connections to the PingOne API.
	OAuthClientSecret *string `json:"OAuthClientSecret,omitempty"`
	// Specifies a passphrase provider that can be used to obtain the OAuth Client Secret used to authenticate connections to the PingOne API.
	OAuthClientSecretPassphraseProvider *string `json:"OAuthClientSecretPassphraseProvider,omitempty"`
	// Specifies the PingOne Environment that will be associated with this PingOne Pass Through Authentication Plugin.
	EnvironmentID string `json:"environmentID"`
	// The base DNs for the local users whose authentication attempts may be passed through to the PingOne service.
	IncludedLocalEntryBaseDN []string `json:"includedLocalEntryBaseDN,omitempty"`
	// A reference to connection criteria that will be used to indicate which bind requests should be passed through to the PingOne service.
	ConnectionCriteria *string `json:"connectionCriteria,omitempty"`
	// A reference to request criteria that will be used to indicate which bind requests should be passed through to the PingOne service.
	RequestCriteria *string `json:"requestCriteria,omitempty"`
	// Indicates whether to attempt the bind in the local server first, or to only send it to the PingOne service.
	TryLocalBind *bool `json:"tryLocalBind,omitempty"`
	// Indicates whether to attempt the authentication in the PingOne service if the local user entry includes a password. This property will only be used if try-local-bind is true.
	OverrideLocalPassword *bool `json:"overrideLocalPassword,omitempty"`
	// Indicates whether to overwrite the user's local password if the local bind fails but the authentication attempt succeeds when attempted in the PingOne service.
	UpdateLocalPassword *bool `json:"updateLocalPassword,omitempty"`
	// This is the DN of the user that will be used to overwrite the user's local password if update-local-password is set. The DN put here should be added to 'ignore-changes-by-dn' in the appropriate Sync Source.
	UpdateLocalPasswordDN *string `json:"updateLocalPasswordDN,omitempty"`
	// Indicates whether to overwrite the user's local password even if the password used to authenticate to the PingOne service would have failed validation if the user attempted to set it directly.
	AllowLaxPassThroughAuthenticationPasswords *bool                                                    `json:"allowLaxPassThroughAuthenticationPasswords,omitempty"`
	IgnoredPasswordPolicyStateErrorCondition   []EnumpluginIgnoredPasswordPolicyStateErrorConditionProp `json:"ignoredPasswordPolicyStateErrorCondition,omitempty"`
	// The names of the attributes in the local user entry whose values must match the values of the corresponding fields in the PingOne service.
	UserMappingLocalAttribute []string `json:"userMappingLocalAttribute"`
	// The names of the fields in the PingOne service whose values must match the values of the corresponding attributes in the local user entry, as specified in the user-mapping-local-attribute property.
	UserMappingRemoteJSONField []string `json:"userMappingRemoteJSONField"`
	// An optional SCIM filter that will be ANDed with the filter created to identify the account in the PingOne service that corresponds to the local entry. Only the \"eq\", \"sw\", \"and\", and \"or\" filter types may be used.
	AdditionalUserMappingSCIMFilter *string `json:"additionalUserMappingSCIMFilter,omitempty"`
	// A description for this Plugin
	Description *string `json:"description,omitempty"`
	// Indicates whether the plug-in is enabled for use.
	Enabled bool `json:"enabled"`
	// Indicates whether the plug-in should be invoked for internal operations.
	InvokeForInternalOperations *bool `json:"invokeForInternalOperations,omitempty"`
}

// NewAddPingOnePassThroughAuthenticationPluginRequest instantiates a new AddPingOnePassThroughAuthenticationPluginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPingOnePassThroughAuthenticationPluginRequest(pluginName string, schemas []EnumpingOnePassThroughAuthenticationPluginSchemaUrn, apiURL string, authURL string, oAuthClientID string, environmentID string, userMappingLocalAttribute []string, userMappingRemoteJSONField []string, enabled bool) *AddPingOnePassThroughAuthenticationPluginRequest {
	this := AddPingOnePassThroughAuthenticationPluginRequest{}
	this.PluginName = pluginName
	this.Schemas = schemas
	this.ApiURL = apiURL
	this.AuthURL = authURL
	this.OAuthClientID = oAuthClientID
	this.EnvironmentID = environmentID
	this.UserMappingLocalAttribute = userMappingLocalAttribute
	this.UserMappingRemoteJSONField = userMappingRemoteJSONField
	this.Enabled = enabled
	return &this
}

// NewAddPingOnePassThroughAuthenticationPluginRequestWithDefaults instantiates a new AddPingOnePassThroughAuthenticationPluginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPingOnePassThroughAuthenticationPluginRequestWithDefaults() *AddPingOnePassThroughAuthenticationPluginRequest {
	this := AddPingOnePassThroughAuthenticationPluginRequest{}
	return &this
}

// GetPluginName returns the PluginName field value
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetPluginName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PluginName
}

// GetPluginNameOk returns a tuple with the PluginName field value
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetPluginNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PluginName, true
}

// SetPluginName sets field value
func (o *AddPingOnePassThroughAuthenticationPluginRequest) SetPluginName(v string) {
	o.PluginName = v
}

// GetSchemas returns the Schemas field value
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetSchemas() []EnumpingOnePassThroughAuthenticationPluginSchemaUrn {
	if o == nil {
		var ret []EnumpingOnePassThroughAuthenticationPluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetSchemasOk() ([]EnumpingOnePassThroughAuthenticationPluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddPingOnePassThroughAuthenticationPluginRequest) SetSchemas(v []EnumpingOnePassThroughAuthenticationPluginSchemaUrn) {
	o.Schemas = v
}

// GetApiURL returns the ApiURL field value
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetApiURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiURL
}

// GetApiURLOk returns a tuple with the ApiURL field value
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetApiURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiURL, true
}

// SetApiURL sets field value
func (o *AddPingOnePassThroughAuthenticationPluginRequest) SetApiURL(v string) {
	o.ApiURL = v
}

// GetAuthURL returns the AuthURL field value
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetAuthURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthURL
}

// GetAuthURLOk returns a tuple with the AuthURL field value
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetAuthURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthURL, true
}

// SetAuthURL sets field value
func (o *AddPingOnePassThroughAuthenticationPluginRequest) SetAuthURL(v string) {
	o.AuthURL = v
}

// GetOAuthClientID returns the OAuthClientID field value
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetOAuthClientID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OAuthClientID
}

// GetOAuthClientIDOk returns a tuple with the OAuthClientID field value
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetOAuthClientIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OAuthClientID, true
}

// SetOAuthClientID sets field value
func (o *AddPingOnePassThroughAuthenticationPluginRequest) SetOAuthClientID(v string) {
	o.OAuthClientID = v
}

// GetOAuthClientSecret returns the OAuthClientSecret field value if set, zero value otherwise.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetOAuthClientSecret() string {
	if o == nil || IsNil(o.OAuthClientSecret) {
		var ret string
		return ret
	}
	return *o.OAuthClientSecret
}

// GetOAuthClientSecretOk returns a tuple with the OAuthClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetOAuthClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.OAuthClientSecret) {
		return nil, false
	}
	return o.OAuthClientSecret, true
}

// HasOAuthClientSecret returns a boolean if a field has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) HasOAuthClientSecret() bool {
	if o != nil && !IsNil(o.OAuthClientSecret) {
		return true
	}

	return false
}

// SetOAuthClientSecret gets a reference to the given string and assigns it to the OAuthClientSecret field.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) SetOAuthClientSecret(v string) {
	o.OAuthClientSecret = &v
}

// GetOAuthClientSecretPassphraseProvider returns the OAuthClientSecretPassphraseProvider field value if set, zero value otherwise.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetOAuthClientSecretPassphraseProvider() string {
	if o == nil || IsNil(o.OAuthClientSecretPassphraseProvider) {
		var ret string
		return ret
	}
	return *o.OAuthClientSecretPassphraseProvider
}

// GetOAuthClientSecretPassphraseProviderOk returns a tuple with the OAuthClientSecretPassphraseProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetOAuthClientSecretPassphraseProviderOk() (*string, bool) {
	if o == nil || IsNil(o.OAuthClientSecretPassphraseProvider) {
		return nil, false
	}
	return o.OAuthClientSecretPassphraseProvider, true
}

// HasOAuthClientSecretPassphraseProvider returns a boolean if a field has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) HasOAuthClientSecretPassphraseProvider() bool {
	if o != nil && !IsNil(o.OAuthClientSecretPassphraseProvider) {
		return true
	}

	return false
}

// SetOAuthClientSecretPassphraseProvider gets a reference to the given string and assigns it to the OAuthClientSecretPassphraseProvider field.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) SetOAuthClientSecretPassphraseProvider(v string) {
	o.OAuthClientSecretPassphraseProvider = &v
}

// GetEnvironmentID returns the EnvironmentID field value
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetEnvironmentID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentID
}

// GetEnvironmentIDOk returns a tuple with the EnvironmentID field value
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetEnvironmentIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentID, true
}

// SetEnvironmentID sets field value
func (o *AddPingOnePassThroughAuthenticationPluginRequest) SetEnvironmentID(v string) {
	o.EnvironmentID = v
}

// GetIncludedLocalEntryBaseDN returns the IncludedLocalEntryBaseDN field value if set, zero value otherwise.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetIncludedLocalEntryBaseDN() []string {
	if o == nil || IsNil(o.IncludedLocalEntryBaseDN) {
		var ret []string
		return ret
	}
	return o.IncludedLocalEntryBaseDN
}

// GetIncludedLocalEntryBaseDNOk returns a tuple with the IncludedLocalEntryBaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetIncludedLocalEntryBaseDNOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludedLocalEntryBaseDN) {
		return nil, false
	}
	return o.IncludedLocalEntryBaseDN, true
}

// HasIncludedLocalEntryBaseDN returns a boolean if a field has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) HasIncludedLocalEntryBaseDN() bool {
	if o != nil && !IsNil(o.IncludedLocalEntryBaseDN) {
		return true
	}

	return false
}

// SetIncludedLocalEntryBaseDN gets a reference to the given []string and assigns it to the IncludedLocalEntryBaseDN field.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) SetIncludedLocalEntryBaseDN(v []string) {
	o.IncludedLocalEntryBaseDN = v
}

// GetConnectionCriteria returns the ConnectionCriteria field value if set, zero value otherwise.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetConnectionCriteria() string {
	if o == nil || IsNil(o.ConnectionCriteria) {
		var ret string
		return ret
	}
	return *o.ConnectionCriteria
}

// GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetConnectionCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionCriteria) {
		return nil, false
	}
	return o.ConnectionCriteria, true
}

// HasConnectionCriteria returns a boolean if a field has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) HasConnectionCriteria() bool {
	if o != nil && !IsNil(o.ConnectionCriteria) {
		return true
	}

	return false
}

// SetConnectionCriteria gets a reference to the given string and assigns it to the ConnectionCriteria field.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) SetConnectionCriteria(v string) {
	o.ConnectionCriteria = &v
}

// GetRequestCriteria returns the RequestCriteria field value if set, zero value otherwise.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetRequestCriteria() string {
	if o == nil || IsNil(o.RequestCriteria) {
		var ret string
		return ret
	}
	return *o.RequestCriteria
}

// GetRequestCriteriaOk returns a tuple with the RequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetRequestCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.RequestCriteria) {
		return nil, false
	}
	return o.RequestCriteria, true
}

// HasRequestCriteria returns a boolean if a field has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) HasRequestCriteria() bool {
	if o != nil && !IsNil(o.RequestCriteria) {
		return true
	}

	return false
}

// SetRequestCriteria gets a reference to the given string and assigns it to the RequestCriteria field.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) SetRequestCriteria(v string) {
	o.RequestCriteria = &v
}

// GetTryLocalBind returns the TryLocalBind field value if set, zero value otherwise.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetTryLocalBind() bool {
	if o == nil || IsNil(o.TryLocalBind) {
		var ret bool
		return ret
	}
	return *o.TryLocalBind
}

// GetTryLocalBindOk returns a tuple with the TryLocalBind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetTryLocalBindOk() (*bool, bool) {
	if o == nil || IsNil(o.TryLocalBind) {
		return nil, false
	}
	return o.TryLocalBind, true
}

// HasTryLocalBind returns a boolean if a field has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) HasTryLocalBind() bool {
	if o != nil && !IsNil(o.TryLocalBind) {
		return true
	}

	return false
}

// SetTryLocalBind gets a reference to the given bool and assigns it to the TryLocalBind field.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) SetTryLocalBind(v bool) {
	o.TryLocalBind = &v
}

// GetOverrideLocalPassword returns the OverrideLocalPassword field value if set, zero value otherwise.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetOverrideLocalPassword() bool {
	if o == nil || IsNil(o.OverrideLocalPassword) {
		var ret bool
		return ret
	}
	return *o.OverrideLocalPassword
}

// GetOverrideLocalPasswordOk returns a tuple with the OverrideLocalPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetOverrideLocalPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.OverrideLocalPassword) {
		return nil, false
	}
	return o.OverrideLocalPassword, true
}

// HasOverrideLocalPassword returns a boolean if a field has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) HasOverrideLocalPassword() bool {
	if o != nil && !IsNil(o.OverrideLocalPassword) {
		return true
	}

	return false
}

// SetOverrideLocalPassword gets a reference to the given bool and assigns it to the OverrideLocalPassword field.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) SetOverrideLocalPassword(v bool) {
	o.OverrideLocalPassword = &v
}

// GetUpdateLocalPassword returns the UpdateLocalPassword field value if set, zero value otherwise.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetUpdateLocalPassword() bool {
	if o == nil || IsNil(o.UpdateLocalPassword) {
		var ret bool
		return ret
	}
	return *o.UpdateLocalPassword
}

// GetUpdateLocalPasswordOk returns a tuple with the UpdateLocalPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetUpdateLocalPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateLocalPassword) {
		return nil, false
	}
	return o.UpdateLocalPassword, true
}

// HasUpdateLocalPassword returns a boolean if a field has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) HasUpdateLocalPassword() bool {
	if o != nil && !IsNil(o.UpdateLocalPassword) {
		return true
	}

	return false
}

// SetUpdateLocalPassword gets a reference to the given bool and assigns it to the UpdateLocalPassword field.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) SetUpdateLocalPassword(v bool) {
	o.UpdateLocalPassword = &v
}

// GetUpdateLocalPasswordDN returns the UpdateLocalPasswordDN field value if set, zero value otherwise.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetUpdateLocalPasswordDN() string {
	if o == nil || IsNil(o.UpdateLocalPasswordDN) {
		var ret string
		return ret
	}
	return *o.UpdateLocalPasswordDN
}

// GetUpdateLocalPasswordDNOk returns a tuple with the UpdateLocalPasswordDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetUpdateLocalPasswordDNOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateLocalPasswordDN) {
		return nil, false
	}
	return o.UpdateLocalPasswordDN, true
}

// HasUpdateLocalPasswordDN returns a boolean if a field has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) HasUpdateLocalPasswordDN() bool {
	if o != nil && !IsNil(o.UpdateLocalPasswordDN) {
		return true
	}

	return false
}

// SetUpdateLocalPasswordDN gets a reference to the given string and assigns it to the UpdateLocalPasswordDN field.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) SetUpdateLocalPasswordDN(v string) {
	o.UpdateLocalPasswordDN = &v
}

// GetAllowLaxPassThroughAuthenticationPasswords returns the AllowLaxPassThroughAuthenticationPasswords field value if set, zero value otherwise.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetAllowLaxPassThroughAuthenticationPasswords() bool {
	if o == nil || IsNil(o.AllowLaxPassThroughAuthenticationPasswords) {
		var ret bool
		return ret
	}
	return *o.AllowLaxPassThroughAuthenticationPasswords
}

// GetAllowLaxPassThroughAuthenticationPasswordsOk returns a tuple with the AllowLaxPassThroughAuthenticationPasswords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetAllowLaxPassThroughAuthenticationPasswordsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowLaxPassThroughAuthenticationPasswords) {
		return nil, false
	}
	return o.AllowLaxPassThroughAuthenticationPasswords, true
}

// HasAllowLaxPassThroughAuthenticationPasswords returns a boolean if a field has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) HasAllowLaxPassThroughAuthenticationPasswords() bool {
	if o != nil && !IsNil(o.AllowLaxPassThroughAuthenticationPasswords) {
		return true
	}

	return false
}

// SetAllowLaxPassThroughAuthenticationPasswords gets a reference to the given bool and assigns it to the AllowLaxPassThroughAuthenticationPasswords field.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) SetAllowLaxPassThroughAuthenticationPasswords(v bool) {
	o.AllowLaxPassThroughAuthenticationPasswords = &v
}

// GetIgnoredPasswordPolicyStateErrorCondition returns the IgnoredPasswordPolicyStateErrorCondition field value if set, zero value otherwise.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetIgnoredPasswordPolicyStateErrorCondition() []EnumpluginIgnoredPasswordPolicyStateErrorConditionProp {
	if o == nil || IsNil(o.IgnoredPasswordPolicyStateErrorCondition) {
		var ret []EnumpluginIgnoredPasswordPolicyStateErrorConditionProp
		return ret
	}
	return o.IgnoredPasswordPolicyStateErrorCondition
}

// GetIgnoredPasswordPolicyStateErrorConditionOk returns a tuple with the IgnoredPasswordPolicyStateErrorCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetIgnoredPasswordPolicyStateErrorConditionOk() ([]EnumpluginIgnoredPasswordPolicyStateErrorConditionProp, bool) {
	if o == nil || IsNil(o.IgnoredPasswordPolicyStateErrorCondition) {
		return nil, false
	}
	return o.IgnoredPasswordPolicyStateErrorCondition, true
}

// HasIgnoredPasswordPolicyStateErrorCondition returns a boolean if a field has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) HasIgnoredPasswordPolicyStateErrorCondition() bool {
	if o != nil && !IsNil(o.IgnoredPasswordPolicyStateErrorCondition) {
		return true
	}

	return false
}

// SetIgnoredPasswordPolicyStateErrorCondition gets a reference to the given []EnumpluginIgnoredPasswordPolicyStateErrorConditionProp and assigns it to the IgnoredPasswordPolicyStateErrorCondition field.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) SetIgnoredPasswordPolicyStateErrorCondition(v []EnumpluginIgnoredPasswordPolicyStateErrorConditionProp) {
	o.IgnoredPasswordPolicyStateErrorCondition = v
}

// GetUserMappingLocalAttribute returns the UserMappingLocalAttribute field value
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetUserMappingLocalAttribute() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.UserMappingLocalAttribute
}

// GetUserMappingLocalAttributeOk returns a tuple with the UserMappingLocalAttribute field value
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetUserMappingLocalAttributeOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserMappingLocalAttribute, true
}

// SetUserMappingLocalAttribute sets field value
func (o *AddPingOnePassThroughAuthenticationPluginRequest) SetUserMappingLocalAttribute(v []string) {
	o.UserMappingLocalAttribute = v
}

// GetUserMappingRemoteJSONField returns the UserMappingRemoteJSONField field value
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetUserMappingRemoteJSONField() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.UserMappingRemoteJSONField
}

// GetUserMappingRemoteJSONFieldOk returns a tuple with the UserMappingRemoteJSONField field value
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetUserMappingRemoteJSONFieldOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserMappingRemoteJSONField, true
}

// SetUserMappingRemoteJSONField sets field value
func (o *AddPingOnePassThroughAuthenticationPluginRequest) SetUserMappingRemoteJSONField(v []string) {
	o.UserMappingRemoteJSONField = v
}

// GetAdditionalUserMappingSCIMFilter returns the AdditionalUserMappingSCIMFilter field value if set, zero value otherwise.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetAdditionalUserMappingSCIMFilter() string {
	if o == nil || IsNil(o.AdditionalUserMappingSCIMFilter) {
		var ret string
		return ret
	}
	return *o.AdditionalUserMappingSCIMFilter
}

// GetAdditionalUserMappingSCIMFilterOk returns a tuple with the AdditionalUserMappingSCIMFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetAdditionalUserMappingSCIMFilterOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalUserMappingSCIMFilter) {
		return nil, false
	}
	return o.AdditionalUserMappingSCIMFilter, true
}

// HasAdditionalUserMappingSCIMFilter returns a boolean if a field has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) HasAdditionalUserMappingSCIMFilter() bool {
	if o != nil && !IsNil(o.AdditionalUserMappingSCIMFilter) {
		return true
	}

	return false
}

// SetAdditionalUserMappingSCIMFilter gets a reference to the given string and assigns it to the AdditionalUserMappingSCIMFilter field.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) SetAdditionalUserMappingSCIMFilter(v string) {
	o.AdditionalUserMappingSCIMFilter = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddPingOnePassThroughAuthenticationPluginRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetInvokeForInternalOperations returns the InvokeForInternalOperations field value if set, zero value otherwise.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetInvokeForInternalOperations() bool {
	if o == nil || IsNil(o.InvokeForInternalOperations) {
		var ret bool
		return ret
	}
	return *o.InvokeForInternalOperations
}

// GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) GetInvokeForInternalOperationsOk() (*bool, bool) {
	if o == nil || IsNil(o.InvokeForInternalOperations) {
		return nil, false
	}
	return o.InvokeForInternalOperations, true
}

// HasInvokeForInternalOperations returns a boolean if a field has been set.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) HasInvokeForInternalOperations() bool {
	if o != nil && !IsNil(o.InvokeForInternalOperations) {
		return true
	}

	return false
}

// SetInvokeForInternalOperations gets a reference to the given bool and assigns it to the InvokeForInternalOperations field.
func (o *AddPingOnePassThroughAuthenticationPluginRequest) SetInvokeForInternalOperations(v bool) {
	o.InvokeForInternalOperations = &v
}

func (o AddPingOnePassThroughAuthenticationPluginRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddPingOnePassThroughAuthenticationPluginRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pluginName"] = o.PluginName
	toSerialize["schemas"] = o.Schemas
	toSerialize["apiURL"] = o.ApiURL
	toSerialize["authURL"] = o.AuthURL
	toSerialize["OAuthClientID"] = o.OAuthClientID
	if !IsNil(o.OAuthClientSecret) {
		toSerialize["OAuthClientSecret"] = o.OAuthClientSecret
	}
	if !IsNil(o.OAuthClientSecretPassphraseProvider) {
		toSerialize["OAuthClientSecretPassphraseProvider"] = o.OAuthClientSecretPassphraseProvider
	}
	toSerialize["environmentID"] = o.EnvironmentID
	if !IsNil(o.IncludedLocalEntryBaseDN) {
		toSerialize["includedLocalEntryBaseDN"] = o.IncludedLocalEntryBaseDN
	}
	if !IsNil(o.ConnectionCriteria) {
		toSerialize["connectionCriteria"] = o.ConnectionCriteria
	}
	if !IsNil(o.RequestCriteria) {
		toSerialize["requestCriteria"] = o.RequestCriteria
	}
	if !IsNil(o.TryLocalBind) {
		toSerialize["tryLocalBind"] = o.TryLocalBind
	}
	if !IsNil(o.OverrideLocalPassword) {
		toSerialize["overrideLocalPassword"] = o.OverrideLocalPassword
	}
	if !IsNil(o.UpdateLocalPassword) {
		toSerialize["updateLocalPassword"] = o.UpdateLocalPassword
	}
	if !IsNil(o.UpdateLocalPasswordDN) {
		toSerialize["updateLocalPasswordDN"] = o.UpdateLocalPasswordDN
	}
	if !IsNil(o.AllowLaxPassThroughAuthenticationPasswords) {
		toSerialize["allowLaxPassThroughAuthenticationPasswords"] = o.AllowLaxPassThroughAuthenticationPasswords
	}
	if !IsNil(o.IgnoredPasswordPolicyStateErrorCondition) {
		toSerialize["ignoredPasswordPolicyStateErrorCondition"] = o.IgnoredPasswordPolicyStateErrorCondition
	}
	toSerialize["userMappingLocalAttribute"] = o.UserMappingLocalAttribute
	toSerialize["userMappingRemoteJSONField"] = o.UserMappingRemoteJSONField
	if !IsNil(o.AdditionalUserMappingSCIMFilter) {
		toSerialize["additionalUserMappingSCIMFilter"] = o.AdditionalUserMappingSCIMFilter
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.InvokeForInternalOperations) {
		toSerialize["invokeForInternalOperations"] = o.InvokeForInternalOperations
	}
	return toSerialize, nil
}

type NullableAddPingOnePassThroughAuthenticationPluginRequest struct {
	value *AddPingOnePassThroughAuthenticationPluginRequest
	isSet bool
}

func (v NullableAddPingOnePassThroughAuthenticationPluginRequest) Get() *AddPingOnePassThroughAuthenticationPluginRequest {
	return v.value
}

func (v *NullableAddPingOnePassThroughAuthenticationPluginRequest) Set(val *AddPingOnePassThroughAuthenticationPluginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPingOnePassThroughAuthenticationPluginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPingOnePassThroughAuthenticationPluginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPingOnePassThroughAuthenticationPluginRequest(val *AddPingOnePassThroughAuthenticationPluginRequest) *NullableAddPingOnePassThroughAuthenticationPluginRequest {
	return &NullableAddPingOnePassThroughAuthenticationPluginRequest{value: val, isSet: true}
}

func (v NullableAddPingOnePassThroughAuthenticationPluginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPingOnePassThroughAuthenticationPluginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
