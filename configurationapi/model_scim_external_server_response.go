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

// checks if the ScimExternalServerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScimExternalServerResponse{}

// ScimExternalServerResponse struct for ScimExternalServerResponse
type ScimExternalServerResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumscimExternalServerSchemaUrn                  `json:"schemas"`
	// Name of the External Server
	Id string `json:"id"`
	// The complete URL which will be used to access the SCIM service provider.
	ScimServiceURL string `json:"scimServiceURL"`
	// The name of the login account to use when connecting to the SCIM service provider. This is used in conjunction with the chosen authentication-method.
	UserName *string `json:"userName,omitempty"`
	// The login password for the specified user name. This is used in conjunction with the chosen authentication-method.
	Password *string `json:"password,omitempty"`
	// The passphrase provider to use to obtain the login password for the specified user.
	PassphraseProvider *string `json:"passphraseProvider,omitempty"`
	// Specifies the location for the SCIM External Server.
	Location                   *string                                               `json:"location,omitempty"`
	ConnectionSecurity         EnumexternalServerScimConnectionSecurityProp          `json:"connectionSecurity"`
	AuthenticationMethod       EnumexternalServerScimAuthenticationMethodProp        `json:"authenticationMethod"`
	HostnameVerificationMethod *EnumexternalServerScimHostnameVerificationMethodProp `json:"hostnameVerificationMethod,omitempty"`
	// A reference to an HTTP proxy server that should be used for requests sent to the SCIM service.
	HttpProxyExternalServer *string `json:"httpProxyExternalServer,omitempty"`
	// The key manager provider to use if SSL is to be used for connection-level security. When specifying a value for this property (except when using the Null key manager provider) you must ensure that the external server trusts this server's public certificate by adding this server's public certificate to the external server's trust store.
	KeyManagerProvider *string `json:"keyManagerProvider,omitempty"`
	// The trust manager provider to use if SSL is to be used for connection-level security.
	TrustManagerProvider *string `json:"trustManagerProvider,omitempty"`
	// Specifies the amount of time to wait for a response from the service provider when establishing a connection. If the timeout is exceeded, the Directory Server will attempt to fail over to a different server. A value of zero indicates no timeout.
	ConnectTimeout *string `json:"connectTimeout,omitempty"`
	// Specifies the maximum length of time that an operation should be allowed to block while waiting for a response from the SCIM service provider. A value of zero indicates that there should be no client-side timeout.
	ResponseTimeout *string                               `json:"responseTimeout,omitempty"`
	OAuthTokenType  *EnumexternalServerOAuthTokenTypeProp `json:"OAuthTokenType,omitempty"`
	// The token to use in conjunction with the OAuth authentication-method and the chosen oauth-token-type.
	OAuthToken *string `json:"OAuthToken,omitempty"`
	// A description for this External Server
	Description *string `json:"description,omitempty"`
}

// NewScimExternalServerResponse instantiates a new ScimExternalServerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScimExternalServerResponse(schemas []EnumscimExternalServerSchemaUrn, id string, scimServiceURL string, connectionSecurity EnumexternalServerScimConnectionSecurityProp, authenticationMethod EnumexternalServerScimAuthenticationMethodProp) *ScimExternalServerResponse {
	this := ScimExternalServerResponse{}
	this.Schemas = schemas
	this.Id = id
	this.ScimServiceURL = scimServiceURL
	this.ConnectionSecurity = connectionSecurity
	this.AuthenticationMethod = authenticationMethod
	return &this
}

// NewScimExternalServerResponseWithDefaults instantiates a new ScimExternalServerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScimExternalServerResponseWithDefaults() *ScimExternalServerResponse {
	this := ScimExternalServerResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ScimExternalServerResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimExternalServerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ScimExternalServerResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ScimExternalServerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ScimExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ScimExternalServerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ScimExternalServerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *ScimExternalServerResponse) GetSchemas() []EnumscimExternalServerSchemaUrn {
	if o == nil {
		var ret []EnumscimExternalServerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ScimExternalServerResponse) GetSchemasOk() ([]EnumscimExternalServerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ScimExternalServerResponse) SetSchemas(v []EnumscimExternalServerSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *ScimExternalServerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ScimExternalServerResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ScimExternalServerResponse) SetId(v string) {
	o.Id = v
}

// GetScimServiceURL returns the ScimServiceURL field value
func (o *ScimExternalServerResponse) GetScimServiceURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScimServiceURL
}

// GetScimServiceURLOk returns a tuple with the ScimServiceURL field value
// and a boolean to check if the value has been set.
func (o *ScimExternalServerResponse) GetScimServiceURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScimServiceURL, true
}

// SetScimServiceURL sets field value
func (o *ScimExternalServerResponse) SetScimServiceURL(v string) {
	o.ScimServiceURL = v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *ScimExternalServerResponse) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimExternalServerResponse) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *ScimExternalServerResponse) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *ScimExternalServerResponse) SetUserName(v string) {
	o.UserName = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ScimExternalServerResponse) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimExternalServerResponse) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ScimExternalServerResponse) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ScimExternalServerResponse) SetPassword(v string) {
	o.Password = &v
}

// GetPassphraseProvider returns the PassphraseProvider field value if set, zero value otherwise.
func (o *ScimExternalServerResponse) GetPassphraseProvider() string {
	if o == nil || IsNil(o.PassphraseProvider) {
		var ret string
		return ret
	}
	return *o.PassphraseProvider
}

// GetPassphraseProviderOk returns a tuple with the PassphraseProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimExternalServerResponse) GetPassphraseProviderOk() (*string, bool) {
	if o == nil || IsNil(o.PassphraseProvider) {
		return nil, false
	}
	return o.PassphraseProvider, true
}

// HasPassphraseProvider returns a boolean if a field has been set.
func (o *ScimExternalServerResponse) HasPassphraseProvider() bool {
	if o != nil && !IsNil(o.PassphraseProvider) {
		return true
	}

	return false
}

// SetPassphraseProvider gets a reference to the given string and assigns it to the PassphraseProvider field.
func (o *ScimExternalServerResponse) SetPassphraseProvider(v string) {
	o.PassphraseProvider = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ScimExternalServerResponse) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimExternalServerResponse) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ScimExternalServerResponse) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *ScimExternalServerResponse) SetLocation(v string) {
	o.Location = &v
}

// GetConnectionSecurity returns the ConnectionSecurity field value
func (o *ScimExternalServerResponse) GetConnectionSecurity() EnumexternalServerScimConnectionSecurityProp {
	if o == nil {
		var ret EnumexternalServerScimConnectionSecurityProp
		return ret
	}

	return o.ConnectionSecurity
}

// GetConnectionSecurityOk returns a tuple with the ConnectionSecurity field value
// and a boolean to check if the value has been set.
func (o *ScimExternalServerResponse) GetConnectionSecurityOk() (*EnumexternalServerScimConnectionSecurityProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionSecurity, true
}

// SetConnectionSecurity sets field value
func (o *ScimExternalServerResponse) SetConnectionSecurity(v EnumexternalServerScimConnectionSecurityProp) {
	o.ConnectionSecurity = v
}

// GetAuthenticationMethod returns the AuthenticationMethod field value
func (o *ScimExternalServerResponse) GetAuthenticationMethod() EnumexternalServerScimAuthenticationMethodProp {
	if o == nil {
		var ret EnumexternalServerScimAuthenticationMethodProp
		return ret
	}

	return o.AuthenticationMethod
}

// GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field value
// and a boolean to check if the value has been set.
func (o *ScimExternalServerResponse) GetAuthenticationMethodOk() (*EnumexternalServerScimAuthenticationMethodProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationMethod, true
}

// SetAuthenticationMethod sets field value
func (o *ScimExternalServerResponse) SetAuthenticationMethod(v EnumexternalServerScimAuthenticationMethodProp) {
	o.AuthenticationMethod = v
}

// GetHostnameVerificationMethod returns the HostnameVerificationMethod field value if set, zero value otherwise.
func (o *ScimExternalServerResponse) GetHostnameVerificationMethod() EnumexternalServerScimHostnameVerificationMethodProp {
	if o == nil || IsNil(o.HostnameVerificationMethod) {
		var ret EnumexternalServerScimHostnameVerificationMethodProp
		return ret
	}
	return *o.HostnameVerificationMethod
}

// GetHostnameVerificationMethodOk returns a tuple with the HostnameVerificationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimExternalServerResponse) GetHostnameVerificationMethodOk() (*EnumexternalServerScimHostnameVerificationMethodProp, bool) {
	if o == nil || IsNil(o.HostnameVerificationMethod) {
		return nil, false
	}
	return o.HostnameVerificationMethod, true
}

// HasHostnameVerificationMethod returns a boolean if a field has been set.
func (o *ScimExternalServerResponse) HasHostnameVerificationMethod() bool {
	if o != nil && !IsNil(o.HostnameVerificationMethod) {
		return true
	}

	return false
}

// SetHostnameVerificationMethod gets a reference to the given EnumexternalServerScimHostnameVerificationMethodProp and assigns it to the HostnameVerificationMethod field.
func (o *ScimExternalServerResponse) SetHostnameVerificationMethod(v EnumexternalServerScimHostnameVerificationMethodProp) {
	o.HostnameVerificationMethod = &v
}

// GetHttpProxyExternalServer returns the HttpProxyExternalServer field value if set, zero value otherwise.
func (o *ScimExternalServerResponse) GetHttpProxyExternalServer() string {
	if o == nil || IsNil(o.HttpProxyExternalServer) {
		var ret string
		return ret
	}
	return *o.HttpProxyExternalServer
}

// GetHttpProxyExternalServerOk returns a tuple with the HttpProxyExternalServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimExternalServerResponse) GetHttpProxyExternalServerOk() (*string, bool) {
	if o == nil || IsNil(o.HttpProxyExternalServer) {
		return nil, false
	}
	return o.HttpProxyExternalServer, true
}

// HasHttpProxyExternalServer returns a boolean if a field has been set.
func (o *ScimExternalServerResponse) HasHttpProxyExternalServer() bool {
	if o != nil && !IsNil(o.HttpProxyExternalServer) {
		return true
	}

	return false
}

// SetHttpProxyExternalServer gets a reference to the given string and assigns it to the HttpProxyExternalServer field.
func (o *ScimExternalServerResponse) SetHttpProxyExternalServer(v string) {
	o.HttpProxyExternalServer = &v
}

// GetKeyManagerProvider returns the KeyManagerProvider field value if set, zero value otherwise.
func (o *ScimExternalServerResponse) GetKeyManagerProvider() string {
	if o == nil || IsNil(o.KeyManagerProvider) {
		var ret string
		return ret
	}
	return *o.KeyManagerProvider
}

// GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimExternalServerResponse) GetKeyManagerProviderOk() (*string, bool) {
	if o == nil || IsNil(o.KeyManagerProvider) {
		return nil, false
	}
	return o.KeyManagerProvider, true
}

// HasKeyManagerProvider returns a boolean if a field has been set.
func (o *ScimExternalServerResponse) HasKeyManagerProvider() bool {
	if o != nil && !IsNil(o.KeyManagerProvider) {
		return true
	}

	return false
}

// SetKeyManagerProvider gets a reference to the given string and assigns it to the KeyManagerProvider field.
func (o *ScimExternalServerResponse) SetKeyManagerProvider(v string) {
	o.KeyManagerProvider = &v
}

// GetTrustManagerProvider returns the TrustManagerProvider field value if set, zero value otherwise.
func (o *ScimExternalServerResponse) GetTrustManagerProvider() string {
	if o == nil || IsNil(o.TrustManagerProvider) {
		var ret string
		return ret
	}
	return *o.TrustManagerProvider
}

// GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimExternalServerResponse) GetTrustManagerProviderOk() (*string, bool) {
	if o == nil || IsNil(o.TrustManagerProvider) {
		return nil, false
	}
	return o.TrustManagerProvider, true
}

// HasTrustManagerProvider returns a boolean if a field has been set.
func (o *ScimExternalServerResponse) HasTrustManagerProvider() bool {
	if o != nil && !IsNil(o.TrustManagerProvider) {
		return true
	}

	return false
}

// SetTrustManagerProvider gets a reference to the given string and assigns it to the TrustManagerProvider field.
func (o *ScimExternalServerResponse) SetTrustManagerProvider(v string) {
	o.TrustManagerProvider = &v
}

// GetConnectTimeout returns the ConnectTimeout field value if set, zero value otherwise.
func (o *ScimExternalServerResponse) GetConnectTimeout() string {
	if o == nil || IsNil(o.ConnectTimeout) {
		var ret string
		return ret
	}
	return *o.ConnectTimeout
}

// GetConnectTimeoutOk returns a tuple with the ConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimExternalServerResponse) GetConnectTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectTimeout) {
		return nil, false
	}
	return o.ConnectTimeout, true
}

// HasConnectTimeout returns a boolean if a field has been set.
func (o *ScimExternalServerResponse) HasConnectTimeout() bool {
	if o != nil && !IsNil(o.ConnectTimeout) {
		return true
	}

	return false
}

// SetConnectTimeout gets a reference to the given string and assigns it to the ConnectTimeout field.
func (o *ScimExternalServerResponse) SetConnectTimeout(v string) {
	o.ConnectTimeout = &v
}

// GetResponseTimeout returns the ResponseTimeout field value if set, zero value otherwise.
func (o *ScimExternalServerResponse) GetResponseTimeout() string {
	if o == nil || IsNil(o.ResponseTimeout) {
		var ret string
		return ret
	}
	return *o.ResponseTimeout
}

// GetResponseTimeoutOk returns a tuple with the ResponseTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimExternalServerResponse) GetResponseTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseTimeout) {
		return nil, false
	}
	return o.ResponseTimeout, true
}

// HasResponseTimeout returns a boolean if a field has been set.
func (o *ScimExternalServerResponse) HasResponseTimeout() bool {
	if o != nil && !IsNil(o.ResponseTimeout) {
		return true
	}

	return false
}

// SetResponseTimeout gets a reference to the given string and assigns it to the ResponseTimeout field.
func (o *ScimExternalServerResponse) SetResponseTimeout(v string) {
	o.ResponseTimeout = &v
}

// GetOAuthTokenType returns the OAuthTokenType field value if set, zero value otherwise.
func (o *ScimExternalServerResponse) GetOAuthTokenType() EnumexternalServerOAuthTokenTypeProp {
	if o == nil || IsNil(o.OAuthTokenType) {
		var ret EnumexternalServerOAuthTokenTypeProp
		return ret
	}
	return *o.OAuthTokenType
}

// GetOAuthTokenTypeOk returns a tuple with the OAuthTokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimExternalServerResponse) GetOAuthTokenTypeOk() (*EnumexternalServerOAuthTokenTypeProp, bool) {
	if o == nil || IsNil(o.OAuthTokenType) {
		return nil, false
	}
	return o.OAuthTokenType, true
}

// HasOAuthTokenType returns a boolean if a field has been set.
func (o *ScimExternalServerResponse) HasOAuthTokenType() bool {
	if o != nil && !IsNil(o.OAuthTokenType) {
		return true
	}

	return false
}

// SetOAuthTokenType gets a reference to the given EnumexternalServerOAuthTokenTypeProp and assigns it to the OAuthTokenType field.
func (o *ScimExternalServerResponse) SetOAuthTokenType(v EnumexternalServerOAuthTokenTypeProp) {
	o.OAuthTokenType = &v
}

// GetOAuthToken returns the OAuthToken field value if set, zero value otherwise.
func (o *ScimExternalServerResponse) GetOAuthToken() string {
	if o == nil || IsNil(o.OAuthToken) {
		var ret string
		return ret
	}
	return *o.OAuthToken
}

// GetOAuthTokenOk returns a tuple with the OAuthToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimExternalServerResponse) GetOAuthTokenOk() (*string, bool) {
	if o == nil || IsNil(o.OAuthToken) {
		return nil, false
	}
	return o.OAuthToken, true
}

// HasOAuthToken returns a boolean if a field has been set.
func (o *ScimExternalServerResponse) HasOAuthToken() bool {
	if o != nil && !IsNil(o.OAuthToken) {
		return true
	}

	return false
}

// SetOAuthToken gets a reference to the given string and assigns it to the OAuthToken field.
func (o *ScimExternalServerResponse) SetOAuthToken(v string) {
	o.OAuthToken = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ScimExternalServerResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimExternalServerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ScimExternalServerResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ScimExternalServerResponse) SetDescription(v string) {
	o.Description = &v
}

func (o ScimExternalServerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScimExternalServerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
	toSerialize["scimServiceURL"] = o.ScimServiceURL
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.PassphraseProvider) {
		toSerialize["passphraseProvider"] = o.PassphraseProvider
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	toSerialize["connectionSecurity"] = o.ConnectionSecurity
	toSerialize["authenticationMethod"] = o.AuthenticationMethod
	if !IsNil(o.HostnameVerificationMethod) {
		toSerialize["hostnameVerificationMethod"] = o.HostnameVerificationMethod
	}
	if !IsNil(o.HttpProxyExternalServer) {
		toSerialize["httpProxyExternalServer"] = o.HttpProxyExternalServer
	}
	if !IsNil(o.KeyManagerProvider) {
		toSerialize["keyManagerProvider"] = o.KeyManagerProvider
	}
	if !IsNil(o.TrustManagerProvider) {
		toSerialize["trustManagerProvider"] = o.TrustManagerProvider
	}
	if !IsNil(o.ConnectTimeout) {
		toSerialize["connectTimeout"] = o.ConnectTimeout
	}
	if !IsNil(o.ResponseTimeout) {
		toSerialize["responseTimeout"] = o.ResponseTimeout
	}
	if !IsNil(o.OAuthTokenType) {
		toSerialize["OAuthTokenType"] = o.OAuthTokenType
	}
	if !IsNil(o.OAuthToken) {
		toSerialize["OAuthToken"] = o.OAuthToken
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableScimExternalServerResponse struct {
	value *ScimExternalServerResponse
	isSet bool
}

func (v NullableScimExternalServerResponse) Get() *ScimExternalServerResponse {
	return v.value
}

func (v *NullableScimExternalServerResponse) Set(val *ScimExternalServerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableScimExternalServerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableScimExternalServerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScimExternalServerResponse(val *ScimExternalServerResponse) *NullableScimExternalServerResponse {
	return &NullableScimExternalServerResponse{value: val, isSet: true}
}

func (v NullableScimExternalServerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScimExternalServerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
