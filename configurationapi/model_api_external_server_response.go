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

// checks if the ApiExternalServerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiExternalServerResponse{}

// ApiExternalServerResponse struct for ApiExternalServerResponse
type ApiExternalServerResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumapiExternalServerSchemaUrn                   `json:"schemas"`
	// Name of the External Server
	Id string `json:"id"`
	// The base URL of the external server, optionally including port number, for example \"https://externalService:9031\".
	BaseURL                    string                                               `json:"baseURL"`
	HostnameVerificationMethod *EnumexternalServerApiHostnameVerificationMethodProp `json:"hostnameVerificationMethod,omitempty"`
	// The key manager provider to use if SSL (HTTPS) is to be used for connection-level security. When specifying a value for this property (except when using the Null key manager provider) you must ensure that the external server trusts this server's public certificate by adding this server's public certificate to the external server's trust store.
	KeyManagerProvider *string `json:"keyManagerProvider,omitempty"`
	// The trust manager provider to use if SSL (HTTPS) is to be used for connection-level security.
	TrustManagerProvider *string `json:"trustManagerProvider,omitempty"`
	// A list of HTTP headers that will be forwarded by the PingAuthorize Server to the downstream API server.
	AllowedHeader []string `json:"allowedHeader,omitempty"`
	// The name of the login account to use for HTTP requests to the downstream API server using basic authentication. This property is ignored unless an associated Gateway API Endpoint's http-auth-evaluation-behavior property is set to \"evaluate-and-replace\".
	UserName *string `json:"userName,omitempty"`
	// The name of the login password to use for HTTP requests to the downstream API server using basic authentication. This property is ignored unless an associated Gateway API Endpoint's http-auth-evaluation-behavior property is set to \"evaluate-and-replace\".
	Password *string `json:"password,omitempty"`
	// The certificate alias within the keystore to use if SSL (HTTPS) is to be used for connection-level security. When specifying a value for this property you must ensure that the external server trusts this server's public certificate by adding this server's public certificate to the external server's trust store.
	SslCertNickname *string `json:"sslCertNickname,omitempty"`
	// Specifies the maximum length of time to wait for a connection to be established before aborting a request to the server.
	ConnectTimeout *string `json:"connectTimeout,omitempty"`
	// Specifies the maximum length of time to wait for response data to be read from an established connection before aborting a request to the server.
	ResponseTimeout *string `json:"responseTimeout,omitempty"`
	// A description for this External Server
	Description *string `json:"description,omitempty"`
}

// NewApiExternalServerResponse instantiates a new ApiExternalServerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiExternalServerResponse(schemas []EnumapiExternalServerSchemaUrn, id string, baseURL string) *ApiExternalServerResponse {
	this := ApiExternalServerResponse{}
	this.Schemas = schemas
	this.Id = id
	this.BaseURL = baseURL
	return &this
}

// NewApiExternalServerResponseWithDefaults instantiates a new ApiExternalServerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiExternalServerResponseWithDefaults() *ApiExternalServerResponse {
	this := ApiExternalServerResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ApiExternalServerResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiExternalServerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ApiExternalServerResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ApiExternalServerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ApiExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ApiExternalServerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ApiExternalServerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *ApiExternalServerResponse) GetSchemas() []EnumapiExternalServerSchemaUrn {
	if o == nil {
		var ret []EnumapiExternalServerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ApiExternalServerResponse) GetSchemasOk() ([]EnumapiExternalServerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ApiExternalServerResponse) SetSchemas(v []EnumapiExternalServerSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *ApiExternalServerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApiExternalServerResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApiExternalServerResponse) SetId(v string) {
	o.Id = v
}

// GetBaseURL returns the BaseURL field value
func (o *ApiExternalServerResponse) GetBaseURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseURL
}

// GetBaseURLOk returns a tuple with the BaseURL field value
// and a boolean to check if the value has been set.
func (o *ApiExternalServerResponse) GetBaseURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseURL, true
}

// SetBaseURL sets field value
func (o *ApiExternalServerResponse) SetBaseURL(v string) {
	o.BaseURL = v
}

// GetHostnameVerificationMethod returns the HostnameVerificationMethod field value if set, zero value otherwise.
func (o *ApiExternalServerResponse) GetHostnameVerificationMethod() EnumexternalServerApiHostnameVerificationMethodProp {
	if o == nil || IsNil(o.HostnameVerificationMethod) {
		var ret EnumexternalServerApiHostnameVerificationMethodProp
		return ret
	}
	return *o.HostnameVerificationMethod
}

// GetHostnameVerificationMethodOk returns a tuple with the HostnameVerificationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiExternalServerResponse) GetHostnameVerificationMethodOk() (*EnumexternalServerApiHostnameVerificationMethodProp, bool) {
	if o == nil || IsNil(o.HostnameVerificationMethod) {
		return nil, false
	}
	return o.HostnameVerificationMethod, true
}

// HasHostnameVerificationMethod returns a boolean if a field has been set.
func (o *ApiExternalServerResponse) HasHostnameVerificationMethod() bool {
	if o != nil && !IsNil(o.HostnameVerificationMethod) {
		return true
	}

	return false
}

// SetHostnameVerificationMethod gets a reference to the given EnumexternalServerApiHostnameVerificationMethodProp and assigns it to the HostnameVerificationMethod field.
func (o *ApiExternalServerResponse) SetHostnameVerificationMethod(v EnumexternalServerApiHostnameVerificationMethodProp) {
	o.HostnameVerificationMethod = &v
}

// GetKeyManagerProvider returns the KeyManagerProvider field value if set, zero value otherwise.
func (o *ApiExternalServerResponse) GetKeyManagerProvider() string {
	if o == nil || IsNil(o.KeyManagerProvider) {
		var ret string
		return ret
	}
	return *o.KeyManagerProvider
}

// GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiExternalServerResponse) GetKeyManagerProviderOk() (*string, bool) {
	if o == nil || IsNil(o.KeyManagerProvider) {
		return nil, false
	}
	return o.KeyManagerProvider, true
}

// HasKeyManagerProvider returns a boolean if a field has been set.
func (o *ApiExternalServerResponse) HasKeyManagerProvider() bool {
	if o != nil && !IsNil(o.KeyManagerProvider) {
		return true
	}

	return false
}

// SetKeyManagerProvider gets a reference to the given string and assigns it to the KeyManagerProvider field.
func (o *ApiExternalServerResponse) SetKeyManagerProvider(v string) {
	o.KeyManagerProvider = &v
}

// GetTrustManagerProvider returns the TrustManagerProvider field value if set, zero value otherwise.
func (o *ApiExternalServerResponse) GetTrustManagerProvider() string {
	if o == nil || IsNil(o.TrustManagerProvider) {
		var ret string
		return ret
	}
	return *o.TrustManagerProvider
}

// GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiExternalServerResponse) GetTrustManagerProviderOk() (*string, bool) {
	if o == nil || IsNil(o.TrustManagerProvider) {
		return nil, false
	}
	return o.TrustManagerProvider, true
}

// HasTrustManagerProvider returns a boolean if a field has been set.
func (o *ApiExternalServerResponse) HasTrustManagerProvider() bool {
	if o != nil && !IsNil(o.TrustManagerProvider) {
		return true
	}

	return false
}

// SetTrustManagerProvider gets a reference to the given string and assigns it to the TrustManagerProvider field.
func (o *ApiExternalServerResponse) SetTrustManagerProvider(v string) {
	o.TrustManagerProvider = &v
}

// GetAllowedHeader returns the AllowedHeader field value if set, zero value otherwise.
func (o *ApiExternalServerResponse) GetAllowedHeader() []string {
	if o == nil || IsNil(o.AllowedHeader) {
		var ret []string
		return ret
	}
	return o.AllowedHeader
}

// GetAllowedHeaderOk returns a tuple with the AllowedHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiExternalServerResponse) GetAllowedHeaderOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedHeader) {
		return nil, false
	}
	return o.AllowedHeader, true
}

// HasAllowedHeader returns a boolean if a field has been set.
func (o *ApiExternalServerResponse) HasAllowedHeader() bool {
	if o != nil && !IsNil(o.AllowedHeader) {
		return true
	}

	return false
}

// SetAllowedHeader gets a reference to the given []string and assigns it to the AllowedHeader field.
func (o *ApiExternalServerResponse) SetAllowedHeader(v []string) {
	o.AllowedHeader = v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *ApiExternalServerResponse) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiExternalServerResponse) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *ApiExternalServerResponse) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *ApiExternalServerResponse) SetUserName(v string) {
	o.UserName = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ApiExternalServerResponse) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiExternalServerResponse) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ApiExternalServerResponse) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ApiExternalServerResponse) SetPassword(v string) {
	o.Password = &v
}

// GetSslCertNickname returns the SslCertNickname field value if set, zero value otherwise.
func (o *ApiExternalServerResponse) GetSslCertNickname() string {
	if o == nil || IsNil(o.SslCertNickname) {
		var ret string
		return ret
	}
	return *o.SslCertNickname
}

// GetSslCertNicknameOk returns a tuple with the SslCertNickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiExternalServerResponse) GetSslCertNicknameOk() (*string, bool) {
	if o == nil || IsNil(o.SslCertNickname) {
		return nil, false
	}
	return o.SslCertNickname, true
}

// HasSslCertNickname returns a boolean if a field has been set.
func (o *ApiExternalServerResponse) HasSslCertNickname() bool {
	if o != nil && !IsNil(o.SslCertNickname) {
		return true
	}

	return false
}

// SetSslCertNickname gets a reference to the given string and assigns it to the SslCertNickname field.
func (o *ApiExternalServerResponse) SetSslCertNickname(v string) {
	o.SslCertNickname = &v
}

// GetConnectTimeout returns the ConnectTimeout field value if set, zero value otherwise.
func (o *ApiExternalServerResponse) GetConnectTimeout() string {
	if o == nil || IsNil(o.ConnectTimeout) {
		var ret string
		return ret
	}
	return *o.ConnectTimeout
}

// GetConnectTimeoutOk returns a tuple with the ConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiExternalServerResponse) GetConnectTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectTimeout) {
		return nil, false
	}
	return o.ConnectTimeout, true
}

// HasConnectTimeout returns a boolean if a field has been set.
func (o *ApiExternalServerResponse) HasConnectTimeout() bool {
	if o != nil && !IsNil(o.ConnectTimeout) {
		return true
	}

	return false
}

// SetConnectTimeout gets a reference to the given string and assigns it to the ConnectTimeout field.
func (o *ApiExternalServerResponse) SetConnectTimeout(v string) {
	o.ConnectTimeout = &v
}

// GetResponseTimeout returns the ResponseTimeout field value if set, zero value otherwise.
func (o *ApiExternalServerResponse) GetResponseTimeout() string {
	if o == nil || IsNil(o.ResponseTimeout) {
		var ret string
		return ret
	}
	return *o.ResponseTimeout
}

// GetResponseTimeoutOk returns a tuple with the ResponseTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiExternalServerResponse) GetResponseTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseTimeout) {
		return nil, false
	}
	return o.ResponseTimeout, true
}

// HasResponseTimeout returns a boolean if a field has been set.
func (o *ApiExternalServerResponse) HasResponseTimeout() bool {
	if o != nil && !IsNil(o.ResponseTimeout) {
		return true
	}

	return false
}

// SetResponseTimeout gets a reference to the given string and assigns it to the ResponseTimeout field.
func (o *ApiExternalServerResponse) SetResponseTimeout(v string) {
	o.ResponseTimeout = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiExternalServerResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiExternalServerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiExternalServerResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiExternalServerResponse) SetDescription(v string) {
	o.Description = &v
}

func (o ApiExternalServerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiExternalServerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
	toSerialize["baseURL"] = o.BaseURL
	if !IsNil(o.HostnameVerificationMethod) {
		toSerialize["hostnameVerificationMethod"] = o.HostnameVerificationMethod
	}
	if !IsNil(o.KeyManagerProvider) {
		toSerialize["keyManagerProvider"] = o.KeyManagerProvider
	}
	if !IsNil(o.TrustManagerProvider) {
		toSerialize["trustManagerProvider"] = o.TrustManagerProvider
	}
	if !IsNil(o.AllowedHeader) {
		toSerialize["allowedHeader"] = o.AllowedHeader
	}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.SslCertNickname) {
		toSerialize["sslCertNickname"] = o.SslCertNickname
	}
	if !IsNil(o.ConnectTimeout) {
		toSerialize["connectTimeout"] = o.ConnectTimeout
	}
	if !IsNil(o.ResponseTimeout) {
		toSerialize["responseTimeout"] = o.ResponseTimeout
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableApiExternalServerResponse struct {
	value *ApiExternalServerResponse
	isSet bool
}

func (v NullableApiExternalServerResponse) Get() *ApiExternalServerResponse {
	return v.value
}

func (v *NullableApiExternalServerResponse) Set(val *ApiExternalServerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiExternalServerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiExternalServerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiExternalServerResponse(val *ApiExternalServerResponse) *NullableApiExternalServerResponse {
	return &NullableApiExternalServerResponse{value: val, isSet: true}
}

func (v NullableApiExternalServerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiExternalServerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
