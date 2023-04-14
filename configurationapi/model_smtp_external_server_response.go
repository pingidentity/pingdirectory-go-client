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

// checks if the SmtpExternalServerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmtpExternalServerResponse{}

// SmtpExternalServerResponse struct for SmtpExternalServerResponse
type SmtpExternalServerResponse struct {
	// Name of the External Server
	Id      string                            `json:"id"`
	Schemas []EnumsmtpExternalServerSchemaUrn `json:"schemas"`
	// The host name of the smtp server.
	ServerHostName string `json:"serverHostName"`
	// The port number where the smtp server listens for requests.
	ServerPort   *int64                              `json:"serverPort,omitempty"`
	SmtpSecurity *EnumexternalServerSmtpSecurityProp `json:"smtpSecurity,omitempty"`
	// The name of the login account to use when connecting to the smtp server. Both username and password must be supplied if this attribute is set.
	UserName *string `json:"userName,omitempty"`
	// The login password for the specified user name. Both username and password must be supplied if this attribute is set.
	Password *string `json:"password,omitempty"`
	// The passphrase provider to use to obtain the login password for the specified user.
	PassphraseProvider *string `json:"passphraseProvider,omitempty"`
	// Specifies the maximum length of time that a connection or attempted connection to a SMTP server may take.
	SmtpTimeout *string `json:"smtpTimeout,omitempty"`
	// Specifies the connection properties for the smtp server.
	SmtpConnectionProperties []string `json:"smtpConnectionProperties,omitempty"`
	// A description for this External Server
	Description                                   *string                                            `json:"description,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewSmtpExternalServerResponse instantiates a new SmtpExternalServerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmtpExternalServerResponse(id string, schemas []EnumsmtpExternalServerSchemaUrn, serverHostName string) *SmtpExternalServerResponse {
	this := SmtpExternalServerResponse{}
	this.Id = id
	this.Schemas = schemas
	this.ServerHostName = serverHostName
	return &this
}

// NewSmtpExternalServerResponseWithDefaults instantiates a new SmtpExternalServerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmtpExternalServerResponseWithDefaults() *SmtpExternalServerResponse {
	this := SmtpExternalServerResponse{}
	return &this
}

// GetId returns the Id field value
func (o *SmtpExternalServerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SmtpExternalServerResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SmtpExternalServerResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *SmtpExternalServerResponse) GetSchemas() []EnumsmtpExternalServerSchemaUrn {
	if o == nil {
		var ret []EnumsmtpExternalServerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *SmtpExternalServerResponse) GetSchemasOk() ([]EnumsmtpExternalServerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *SmtpExternalServerResponse) SetSchemas(v []EnumsmtpExternalServerSchemaUrn) {
	o.Schemas = v
}

// GetServerHostName returns the ServerHostName field value
func (o *SmtpExternalServerResponse) GetServerHostName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerHostName
}

// GetServerHostNameOk returns a tuple with the ServerHostName field value
// and a boolean to check if the value has been set.
func (o *SmtpExternalServerResponse) GetServerHostNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerHostName, true
}

// SetServerHostName sets field value
func (o *SmtpExternalServerResponse) SetServerHostName(v string) {
	o.ServerHostName = v
}

// GetServerPort returns the ServerPort field value if set, zero value otherwise.
func (o *SmtpExternalServerResponse) GetServerPort() int64 {
	if o == nil || IsNil(o.ServerPort) {
		var ret int64
		return ret
	}
	return *o.ServerPort
}

// GetServerPortOk returns a tuple with the ServerPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpExternalServerResponse) GetServerPortOk() (*int64, bool) {
	if o == nil || IsNil(o.ServerPort) {
		return nil, false
	}
	return o.ServerPort, true
}

// HasServerPort returns a boolean if a field has been set.
func (o *SmtpExternalServerResponse) HasServerPort() bool {
	if o != nil && !IsNil(o.ServerPort) {
		return true
	}

	return false
}

// SetServerPort gets a reference to the given int64 and assigns it to the ServerPort field.
func (o *SmtpExternalServerResponse) SetServerPort(v int64) {
	o.ServerPort = &v
}

// GetSmtpSecurity returns the SmtpSecurity field value if set, zero value otherwise.
func (o *SmtpExternalServerResponse) GetSmtpSecurity() EnumexternalServerSmtpSecurityProp {
	if o == nil || IsNil(o.SmtpSecurity) {
		var ret EnumexternalServerSmtpSecurityProp
		return ret
	}
	return *o.SmtpSecurity
}

// GetSmtpSecurityOk returns a tuple with the SmtpSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpExternalServerResponse) GetSmtpSecurityOk() (*EnumexternalServerSmtpSecurityProp, bool) {
	if o == nil || IsNil(o.SmtpSecurity) {
		return nil, false
	}
	return o.SmtpSecurity, true
}

// HasSmtpSecurity returns a boolean if a field has been set.
func (o *SmtpExternalServerResponse) HasSmtpSecurity() bool {
	if o != nil && !IsNil(o.SmtpSecurity) {
		return true
	}

	return false
}

// SetSmtpSecurity gets a reference to the given EnumexternalServerSmtpSecurityProp and assigns it to the SmtpSecurity field.
func (o *SmtpExternalServerResponse) SetSmtpSecurity(v EnumexternalServerSmtpSecurityProp) {
	o.SmtpSecurity = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *SmtpExternalServerResponse) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpExternalServerResponse) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *SmtpExternalServerResponse) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *SmtpExternalServerResponse) SetUserName(v string) {
	o.UserName = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SmtpExternalServerResponse) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpExternalServerResponse) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *SmtpExternalServerResponse) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *SmtpExternalServerResponse) SetPassword(v string) {
	o.Password = &v
}

// GetPassphraseProvider returns the PassphraseProvider field value if set, zero value otherwise.
func (o *SmtpExternalServerResponse) GetPassphraseProvider() string {
	if o == nil || IsNil(o.PassphraseProvider) {
		var ret string
		return ret
	}
	return *o.PassphraseProvider
}

// GetPassphraseProviderOk returns a tuple with the PassphraseProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpExternalServerResponse) GetPassphraseProviderOk() (*string, bool) {
	if o == nil || IsNil(o.PassphraseProvider) {
		return nil, false
	}
	return o.PassphraseProvider, true
}

// HasPassphraseProvider returns a boolean if a field has been set.
func (o *SmtpExternalServerResponse) HasPassphraseProvider() bool {
	if o != nil && !IsNil(o.PassphraseProvider) {
		return true
	}

	return false
}

// SetPassphraseProvider gets a reference to the given string and assigns it to the PassphraseProvider field.
func (o *SmtpExternalServerResponse) SetPassphraseProvider(v string) {
	o.PassphraseProvider = &v
}

// GetSmtpTimeout returns the SmtpTimeout field value if set, zero value otherwise.
func (o *SmtpExternalServerResponse) GetSmtpTimeout() string {
	if o == nil || IsNil(o.SmtpTimeout) {
		var ret string
		return ret
	}
	return *o.SmtpTimeout
}

// GetSmtpTimeoutOk returns a tuple with the SmtpTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpExternalServerResponse) GetSmtpTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.SmtpTimeout) {
		return nil, false
	}
	return o.SmtpTimeout, true
}

// HasSmtpTimeout returns a boolean if a field has been set.
func (o *SmtpExternalServerResponse) HasSmtpTimeout() bool {
	if o != nil && !IsNil(o.SmtpTimeout) {
		return true
	}

	return false
}

// SetSmtpTimeout gets a reference to the given string and assigns it to the SmtpTimeout field.
func (o *SmtpExternalServerResponse) SetSmtpTimeout(v string) {
	o.SmtpTimeout = &v
}

// GetSmtpConnectionProperties returns the SmtpConnectionProperties field value if set, zero value otherwise.
func (o *SmtpExternalServerResponse) GetSmtpConnectionProperties() []string {
	if o == nil || IsNil(o.SmtpConnectionProperties) {
		var ret []string
		return ret
	}
	return o.SmtpConnectionProperties
}

// GetSmtpConnectionPropertiesOk returns a tuple with the SmtpConnectionProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpExternalServerResponse) GetSmtpConnectionPropertiesOk() ([]string, bool) {
	if o == nil || IsNil(o.SmtpConnectionProperties) {
		return nil, false
	}
	return o.SmtpConnectionProperties, true
}

// HasSmtpConnectionProperties returns a boolean if a field has been set.
func (o *SmtpExternalServerResponse) HasSmtpConnectionProperties() bool {
	if o != nil && !IsNil(o.SmtpConnectionProperties) {
		return true
	}

	return false
}

// SetSmtpConnectionProperties gets a reference to the given []string and assigns it to the SmtpConnectionProperties field.
func (o *SmtpExternalServerResponse) SetSmtpConnectionProperties(v []string) {
	o.SmtpConnectionProperties = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SmtpExternalServerResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpExternalServerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SmtpExternalServerResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SmtpExternalServerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SmtpExternalServerResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpExternalServerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SmtpExternalServerResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *SmtpExternalServerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *SmtpExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *SmtpExternalServerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *SmtpExternalServerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o SmtpExternalServerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmtpExternalServerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["schemas"] = o.Schemas
	toSerialize["serverHostName"] = o.ServerHostName
	if !IsNil(o.ServerPort) {
		toSerialize["serverPort"] = o.ServerPort
	}
	if !IsNil(o.SmtpSecurity) {
		toSerialize["smtpSecurity"] = o.SmtpSecurity
	}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.PassphraseProvider) {
		toSerialize["passphraseProvider"] = o.PassphraseProvider
	}
	if !IsNil(o.SmtpTimeout) {
		toSerialize["smtpTimeout"] = o.SmtpTimeout
	}
	if !IsNil(o.SmtpConnectionProperties) {
		toSerialize["smtpConnectionProperties"] = o.SmtpConnectionProperties
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableSmtpExternalServerResponse struct {
	value *SmtpExternalServerResponse
	isSet bool
}

func (v NullableSmtpExternalServerResponse) Get() *SmtpExternalServerResponse {
	return v.value
}

func (v *NullableSmtpExternalServerResponse) Set(val *SmtpExternalServerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSmtpExternalServerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSmtpExternalServerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmtpExternalServerResponse(val *SmtpExternalServerResponse) *NullableSmtpExternalServerResponse {
	return &NullableSmtpExternalServerResponse{value: val, isSet: true}
}

func (v NullableSmtpExternalServerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmtpExternalServerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
