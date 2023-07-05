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

// checks if the AddPingOnePassThroughAuthenticationHandlerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddPingOnePassThroughAuthenticationHandlerRequest{}

// AddPingOnePassThroughAuthenticationHandlerRequest struct for AddPingOnePassThroughAuthenticationHandlerRequest
type AddPingOnePassThroughAuthenticationHandlerRequest struct {
	// Name of the new Pass Through Authentication Handler
	HandlerName string                                                 `json:"handlerName"`
	Schemas     []EnumpingOnePassThroughAuthenticationHandlerSchemaUrn `json:"schemas"`
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
	// Specifies the PingOne Environment that will be associated with this PingOne Pass Through Authentication Handler.
	EnvironmentID string `json:"environmentID"`
	// A reference to an HTTP proxy server that should be used for requests sent to the PingOne service.
	HttpProxyExternalServer *string `json:"httpProxyExternalServer,omitempty"`
	// The names of the attributes in the local user entry whose values must match the values of the corresponding fields in the PingOne service.
	UserMappingLocalAttribute []string `json:"userMappingLocalAttribute"`
	// The names of the fields in the PingOne service whose values must match the values of the corresponding attributes in the local user entry, as specified in the user-mapping-local-attribute property.
	UserMappingRemoteJSONField []string `json:"userMappingRemoteJSONField"`
	// An optional SCIM filter that will be ANDed with the filter created to identify the account in the PingOne service that corresponds to the local entry. Only the \"eq\", \"sw\", \"and\", and \"or\" filter types may be used.
	AdditionalUserMappingSCIMFilter *string `json:"additionalUserMappingSCIMFilter,omitempty"`
	// A description for this Pass Through Authentication Handler
	Description *string `json:"description,omitempty"`
	// The base DNs for the local users whose authentication attempts may be passed through to the external authentication service.
	IncludedLocalEntryBaseDN []string `json:"includedLocalEntryBaseDN,omitempty"`
	// A reference to connection criteria that will be used to indicate which bind requests should be passed through to the external authentication service.
	ConnectionCriteria *string `json:"connectionCriteria,omitempty"`
	// A reference to request criteria that will be used to indicate which bind requests should be passed through to the external authentication service.
	RequestCriteria *string `json:"requestCriteria,omitempty"`
}

// NewAddPingOnePassThroughAuthenticationHandlerRequest instantiates a new AddPingOnePassThroughAuthenticationHandlerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPingOnePassThroughAuthenticationHandlerRequest(handlerName string, schemas []EnumpingOnePassThroughAuthenticationHandlerSchemaUrn, apiURL string, authURL string, oAuthClientID string, environmentID string, userMappingLocalAttribute []string, userMappingRemoteJSONField []string) *AddPingOnePassThroughAuthenticationHandlerRequest {
	this := AddPingOnePassThroughAuthenticationHandlerRequest{}
	this.HandlerName = handlerName
	this.Schemas = schemas
	this.ApiURL = apiURL
	this.AuthURL = authURL
	this.OAuthClientID = oAuthClientID
	this.EnvironmentID = environmentID
	this.UserMappingLocalAttribute = userMappingLocalAttribute
	this.UserMappingRemoteJSONField = userMappingRemoteJSONField
	return &this
}

// NewAddPingOnePassThroughAuthenticationHandlerRequestWithDefaults instantiates a new AddPingOnePassThroughAuthenticationHandlerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPingOnePassThroughAuthenticationHandlerRequestWithDefaults() *AddPingOnePassThroughAuthenticationHandlerRequest {
	this := AddPingOnePassThroughAuthenticationHandlerRequest{}
	return &this
}

// GetHandlerName returns the HandlerName field value
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetHandlerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HandlerName
}

// GetHandlerNameOk returns a tuple with the HandlerName field value
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetHandlerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HandlerName, true
}

// SetHandlerName sets field value
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) SetHandlerName(v string) {
	o.HandlerName = v
}

// GetSchemas returns the Schemas field value
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetSchemas() []EnumpingOnePassThroughAuthenticationHandlerSchemaUrn {
	if o == nil {
		var ret []EnumpingOnePassThroughAuthenticationHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetSchemasOk() ([]EnumpingOnePassThroughAuthenticationHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) SetSchemas(v []EnumpingOnePassThroughAuthenticationHandlerSchemaUrn) {
	o.Schemas = v
}

// GetApiURL returns the ApiURL field value
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetApiURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiURL
}

// GetApiURLOk returns a tuple with the ApiURL field value
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetApiURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiURL, true
}

// SetApiURL sets field value
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) SetApiURL(v string) {
	o.ApiURL = v
}

// GetAuthURL returns the AuthURL field value
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetAuthURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthURL
}

// GetAuthURLOk returns a tuple with the AuthURL field value
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetAuthURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthURL, true
}

// SetAuthURL sets field value
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) SetAuthURL(v string) {
	o.AuthURL = v
}

// GetOAuthClientID returns the OAuthClientID field value
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetOAuthClientID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OAuthClientID
}

// GetOAuthClientIDOk returns a tuple with the OAuthClientID field value
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetOAuthClientIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OAuthClientID, true
}

// SetOAuthClientID sets field value
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) SetOAuthClientID(v string) {
	o.OAuthClientID = v
}

// GetOAuthClientSecret returns the OAuthClientSecret field value if set, zero value otherwise.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetOAuthClientSecret() string {
	if o == nil || IsNil(o.OAuthClientSecret) {
		var ret string
		return ret
	}
	return *o.OAuthClientSecret
}

// GetOAuthClientSecretOk returns a tuple with the OAuthClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetOAuthClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.OAuthClientSecret) {
		return nil, false
	}
	return o.OAuthClientSecret, true
}

// HasOAuthClientSecret returns a boolean if a field has been set.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) HasOAuthClientSecret() bool {
	if o != nil && !IsNil(o.OAuthClientSecret) {
		return true
	}

	return false
}

// SetOAuthClientSecret gets a reference to the given string and assigns it to the OAuthClientSecret field.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) SetOAuthClientSecret(v string) {
	o.OAuthClientSecret = &v
}

// GetOAuthClientSecretPassphraseProvider returns the OAuthClientSecretPassphraseProvider field value if set, zero value otherwise.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetOAuthClientSecretPassphraseProvider() string {
	if o == nil || IsNil(o.OAuthClientSecretPassphraseProvider) {
		var ret string
		return ret
	}
	return *o.OAuthClientSecretPassphraseProvider
}

// GetOAuthClientSecretPassphraseProviderOk returns a tuple with the OAuthClientSecretPassphraseProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetOAuthClientSecretPassphraseProviderOk() (*string, bool) {
	if o == nil || IsNil(o.OAuthClientSecretPassphraseProvider) {
		return nil, false
	}
	return o.OAuthClientSecretPassphraseProvider, true
}

// HasOAuthClientSecretPassphraseProvider returns a boolean if a field has been set.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) HasOAuthClientSecretPassphraseProvider() bool {
	if o != nil && !IsNil(o.OAuthClientSecretPassphraseProvider) {
		return true
	}

	return false
}

// SetOAuthClientSecretPassphraseProvider gets a reference to the given string and assigns it to the OAuthClientSecretPassphraseProvider field.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) SetOAuthClientSecretPassphraseProvider(v string) {
	o.OAuthClientSecretPassphraseProvider = &v
}

// GetEnvironmentID returns the EnvironmentID field value
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetEnvironmentID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentID
}

// GetEnvironmentIDOk returns a tuple with the EnvironmentID field value
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetEnvironmentIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentID, true
}

// SetEnvironmentID sets field value
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) SetEnvironmentID(v string) {
	o.EnvironmentID = v
}

// GetHttpProxyExternalServer returns the HttpProxyExternalServer field value if set, zero value otherwise.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetHttpProxyExternalServer() string {
	if o == nil || IsNil(o.HttpProxyExternalServer) {
		var ret string
		return ret
	}
	return *o.HttpProxyExternalServer
}

// GetHttpProxyExternalServerOk returns a tuple with the HttpProxyExternalServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetHttpProxyExternalServerOk() (*string, bool) {
	if o == nil || IsNil(o.HttpProxyExternalServer) {
		return nil, false
	}
	return o.HttpProxyExternalServer, true
}

// HasHttpProxyExternalServer returns a boolean if a field has been set.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) HasHttpProxyExternalServer() bool {
	if o != nil && !IsNil(o.HttpProxyExternalServer) {
		return true
	}

	return false
}

// SetHttpProxyExternalServer gets a reference to the given string and assigns it to the HttpProxyExternalServer field.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) SetHttpProxyExternalServer(v string) {
	o.HttpProxyExternalServer = &v
}

// GetUserMappingLocalAttribute returns the UserMappingLocalAttribute field value
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetUserMappingLocalAttribute() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.UserMappingLocalAttribute
}

// GetUserMappingLocalAttributeOk returns a tuple with the UserMappingLocalAttribute field value
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetUserMappingLocalAttributeOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserMappingLocalAttribute, true
}

// SetUserMappingLocalAttribute sets field value
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) SetUserMappingLocalAttribute(v []string) {
	o.UserMappingLocalAttribute = v
}

// GetUserMappingRemoteJSONField returns the UserMappingRemoteJSONField field value
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetUserMappingRemoteJSONField() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.UserMappingRemoteJSONField
}

// GetUserMappingRemoteJSONFieldOk returns a tuple with the UserMappingRemoteJSONField field value
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetUserMappingRemoteJSONFieldOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserMappingRemoteJSONField, true
}

// SetUserMappingRemoteJSONField sets field value
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) SetUserMappingRemoteJSONField(v []string) {
	o.UserMappingRemoteJSONField = v
}

// GetAdditionalUserMappingSCIMFilter returns the AdditionalUserMappingSCIMFilter field value if set, zero value otherwise.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetAdditionalUserMappingSCIMFilter() string {
	if o == nil || IsNil(o.AdditionalUserMappingSCIMFilter) {
		var ret string
		return ret
	}
	return *o.AdditionalUserMappingSCIMFilter
}

// GetAdditionalUserMappingSCIMFilterOk returns a tuple with the AdditionalUserMappingSCIMFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetAdditionalUserMappingSCIMFilterOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalUserMappingSCIMFilter) {
		return nil, false
	}
	return o.AdditionalUserMappingSCIMFilter, true
}

// HasAdditionalUserMappingSCIMFilter returns a boolean if a field has been set.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) HasAdditionalUserMappingSCIMFilter() bool {
	if o != nil && !IsNil(o.AdditionalUserMappingSCIMFilter) {
		return true
	}

	return false
}

// SetAdditionalUserMappingSCIMFilter gets a reference to the given string and assigns it to the AdditionalUserMappingSCIMFilter field.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) SetAdditionalUserMappingSCIMFilter(v string) {
	o.AdditionalUserMappingSCIMFilter = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) SetDescription(v string) {
	o.Description = &v
}

// GetIncludedLocalEntryBaseDN returns the IncludedLocalEntryBaseDN field value if set, zero value otherwise.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetIncludedLocalEntryBaseDN() []string {
	if o == nil || IsNil(o.IncludedLocalEntryBaseDN) {
		var ret []string
		return ret
	}
	return o.IncludedLocalEntryBaseDN
}

// GetIncludedLocalEntryBaseDNOk returns a tuple with the IncludedLocalEntryBaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetIncludedLocalEntryBaseDNOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludedLocalEntryBaseDN) {
		return nil, false
	}
	return o.IncludedLocalEntryBaseDN, true
}

// HasIncludedLocalEntryBaseDN returns a boolean if a field has been set.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) HasIncludedLocalEntryBaseDN() bool {
	if o != nil && !IsNil(o.IncludedLocalEntryBaseDN) {
		return true
	}

	return false
}

// SetIncludedLocalEntryBaseDN gets a reference to the given []string and assigns it to the IncludedLocalEntryBaseDN field.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) SetIncludedLocalEntryBaseDN(v []string) {
	o.IncludedLocalEntryBaseDN = v
}

// GetConnectionCriteria returns the ConnectionCriteria field value if set, zero value otherwise.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetConnectionCriteria() string {
	if o == nil || IsNil(o.ConnectionCriteria) {
		var ret string
		return ret
	}
	return *o.ConnectionCriteria
}

// GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetConnectionCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionCriteria) {
		return nil, false
	}
	return o.ConnectionCriteria, true
}

// HasConnectionCriteria returns a boolean if a field has been set.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) HasConnectionCriteria() bool {
	if o != nil && !IsNil(o.ConnectionCriteria) {
		return true
	}

	return false
}

// SetConnectionCriteria gets a reference to the given string and assigns it to the ConnectionCriteria field.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) SetConnectionCriteria(v string) {
	o.ConnectionCriteria = &v
}

// GetRequestCriteria returns the RequestCriteria field value if set, zero value otherwise.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetRequestCriteria() string {
	if o == nil || IsNil(o.RequestCriteria) {
		var ret string
		return ret
	}
	return *o.RequestCriteria
}

// GetRequestCriteriaOk returns a tuple with the RequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) GetRequestCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.RequestCriteria) {
		return nil, false
	}
	return o.RequestCriteria, true
}

// HasRequestCriteria returns a boolean if a field has been set.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) HasRequestCriteria() bool {
	if o != nil && !IsNil(o.RequestCriteria) {
		return true
	}

	return false
}

// SetRequestCriteria gets a reference to the given string and assigns it to the RequestCriteria field.
func (o *AddPingOnePassThroughAuthenticationHandlerRequest) SetRequestCriteria(v string) {
	o.RequestCriteria = &v
}

func (o AddPingOnePassThroughAuthenticationHandlerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddPingOnePassThroughAuthenticationHandlerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["handlerName"] = o.HandlerName
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
	if !IsNil(o.HttpProxyExternalServer) {
		toSerialize["httpProxyExternalServer"] = o.HttpProxyExternalServer
	}
	toSerialize["userMappingLocalAttribute"] = o.UserMappingLocalAttribute
	toSerialize["userMappingRemoteJSONField"] = o.UserMappingRemoteJSONField
	if !IsNil(o.AdditionalUserMappingSCIMFilter) {
		toSerialize["additionalUserMappingSCIMFilter"] = o.AdditionalUserMappingSCIMFilter
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IncludedLocalEntryBaseDN) {
		toSerialize["includedLocalEntryBaseDN"] = o.IncludedLocalEntryBaseDN
	}
	if !IsNil(o.ConnectionCriteria) {
		toSerialize["connectionCriteria"] = o.ConnectionCriteria
	}
	if !IsNil(o.RequestCriteria) {
		toSerialize["requestCriteria"] = o.RequestCriteria
	}
	return toSerialize, nil
}

type NullableAddPingOnePassThroughAuthenticationHandlerRequest struct {
	value *AddPingOnePassThroughAuthenticationHandlerRequest
	isSet bool
}

func (v NullableAddPingOnePassThroughAuthenticationHandlerRequest) Get() *AddPingOnePassThroughAuthenticationHandlerRequest {
	return v.value
}

func (v *NullableAddPingOnePassThroughAuthenticationHandlerRequest) Set(val *AddPingOnePassThroughAuthenticationHandlerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPingOnePassThroughAuthenticationHandlerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPingOnePassThroughAuthenticationHandlerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPingOnePassThroughAuthenticationHandlerRequest(val *AddPingOnePassThroughAuthenticationHandlerRequest) *NullableAddPingOnePassThroughAuthenticationHandlerRequest {
	return &NullableAddPingOnePassThroughAuthenticationHandlerRequest{value: val, isSet: true}
}

func (v NullableAddPingOnePassThroughAuthenticationHandlerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPingOnePassThroughAuthenticationHandlerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
