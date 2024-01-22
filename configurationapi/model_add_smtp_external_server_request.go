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

// checks if the AddSmtpExternalServerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddSmtpExternalServerRequest{}

// AddSmtpExternalServerRequest struct for AddSmtpExternalServerRequest
type AddSmtpExternalServerRequest struct {
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
	Description *string `json:"description,omitempty"`
	// Name of the new External Server
	ServerName string `json:"serverName"`
}

// NewAddSmtpExternalServerRequest instantiates a new AddSmtpExternalServerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSmtpExternalServerRequest(schemas []EnumsmtpExternalServerSchemaUrn, serverHostName string, serverName string) *AddSmtpExternalServerRequest {
	this := AddSmtpExternalServerRequest{}
	this.Schemas = schemas
	this.ServerHostName = serverHostName
	this.ServerName = serverName
	return &this
}

// NewAddSmtpExternalServerRequestWithDefaults instantiates a new AddSmtpExternalServerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSmtpExternalServerRequestWithDefaults() *AddSmtpExternalServerRequest {
	this := AddSmtpExternalServerRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddSmtpExternalServerRequest) GetSchemas() []EnumsmtpExternalServerSchemaUrn {
	if o == nil {
		var ret []EnumsmtpExternalServerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddSmtpExternalServerRequest) GetSchemasOk() ([]EnumsmtpExternalServerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddSmtpExternalServerRequest) SetSchemas(v []EnumsmtpExternalServerSchemaUrn) {
	o.Schemas = v
}

// GetServerHostName returns the ServerHostName field value
func (o *AddSmtpExternalServerRequest) GetServerHostName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerHostName
}

// GetServerHostNameOk returns a tuple with the ServerHostName field value
// and a boolean to check if the value has been set.
func (o *AddSmtpExternalServerRequest) GetServerHostNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerHostName, true
}

// SetServerHostName sets field value
func (o *AddSmtpExternalServerRequest) SetServerHostName(v string) {
	o.ServerHostName = v
}

// GetServerPort returns the ServerPort field value if set, zero value otherwise.
func (o *AddSmtpExternalServerRequest) GetServerPort() int64 {
	if o == nil || IsNil(o.ServerPort) {
		var ret int64
		return ret
	}
	return *o.ServerPort
}

// GetServerPortOk returns a tuple with the ServerPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSmtpExternalServerRequest) GetServerPortOk() (*int64, bool) {
	if o == nil || IsNil(o.ServerPort) {
		return nil, false
	}
	return o.ServerPort, true
}

// HasServerPort returns a boolean if a field has been set.
func (o *AddSmtpExternalServerRequest) HasServerPort() bool {
	if o != nil && !IsNil(o.ServerPort) {
		return true
	}

	return false
}

// SetServerPort gets a reference to the given int64 and assigns it to the ServerPort field.
func (o *AddSmtpExternalServerRequest) SetServerPort(v int64) {
	o.ServerPort = &v
}

// GetSmtpSecurity returns the SmtpSecurity field value if set, zero value otherwise.
func (o *AddSmtpExternalServerRequest) GetSmtpSecurity() EnumexternalServerSmtpSecurityProp {
	if o == nil || IsNil(o.SmtpSecurity) {
		var ret EnumexternalServerSmtpSecurityProp
		return ret
	}
	return *o.SmtpSecurity
}

// GetSmtpSecurityOk returns a tuple with the SmtpSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSmtpExternalServerRequest) GetSmtpSecurityOk() (*EnumexternalServerSmtpSecurityProp, bool) {
	if o == nil || IsNil(o.SmtpSecurity) {
		return nil, false
	}
	return o.SmtpSecurity, true
}

// HasSmtpSecurity returns a boolean if a field has been set.
func (o *AddSmtpExternalServerRequest) HasSmtpSecurity() bool {
	if o != nil && !IsNil(o.SmtpSecurity) {
		return true
	}

	return false
}

// SetSmtpSecurity gets a reference to the given EnumexternalServerSmtpSecurityProp and assigns it to the SmtpSecurity field.
func (o *AddSmtpExternalServerRequest) SetSmtpSecurity(v EnumexternalServerSmtpSecurityProp) {
	o.SmtpSecurity = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *AddSmtpExternalServerRequest) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSmtpExternalServerRequest) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *AddSmtpExternalServerRequest) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *AddSmtpExternalServerRequest) SetUserName(v string) {
	o.UserName = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *AddSmtpExternalServerRequest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSmtpExternalServerRequest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *AddSmtpExternalServerRequest) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *AddSmtpExternalServerRequest) SetPassword(v string) {
	o.Password = &v
}

// GetPassphraseProvider returns the PassphraseProvider field value if set, zero value otherwise.
func (o *AddSmtpExternalServerRequest) GetPassphraseProvider() string {
	if o == nil || IsNil(o.PassphraseProvider) {
		var ret string
		return ret
	}
	return *o.PassphraseProvider
}

// GetPassphraseProviderOk returns a tuple with the PassphraseProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSmtpExternalServerRequest) GetPassphraseProviderOk() (*string, bool) {
	if o == nil || IsNil(o.PassphraseProvider) {
		return nil, false
	}
	return o.PassphraseProvider, true
}

// HasPassphraseProvider returns a boolean if a field has been set.
func (o *AddSmtpExternalServerRequest) HasPassphraseProvider() bool {
	if o != nil && !IsNil(o.PassphraseProvider) {
		return true
	}

	return false
}

// SetPassphraseProvider gets a reference to the given string and assigns it to the PassphraseProvider field.
func (o *AddSmtpExternalServerRequest) SetPassphraseProvider(v string) {
	o.PassphraseProvider = &v
}

// GetSmtpTimeout returns the SmtpTimeout field value if set, zero value otherwise.
func (o *AddSmtpExternalServerRequest) GetSmtpTimeout() string {
	if o == nil || IsNil(o.SmtpTimeout) {
		var ret string
		return ret
	}
	return *o.SmtpTimeout
}

// GetSmtpTimeoutOk returns a tuple with the SmtpTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSmtpExternalServerRequest) GetSmtpTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.SmtpTimeout) {
		return nil, false
	}
	return o.SmtpTimeout, true
}

// HasSmtpTimeout returns a boolean if a field has been set.
func (o *AddSmtpExternalServerRequest) HasSmtpTimeout() bool {
	if o != nil && !IsNil(o.SmtpTimeout) {
		return true
	}

	return false
}

// SetSmtpTimeout gets a reference to the given string and assigns it to the SmtpTimeout field.
func (o *AddSmtpExternalServerRequest) SetSmtpTimeout(v string) {
	o.SmtpTimeout = &v
}

// GetSmtpConnectionProperties returns the SmtpConnectionProperties field value if set, zero value otherwise.
func (o *AddSmtpExternalServerRequest) GetSmtpConnectionProperties() []string {
	if o == nil || IsNil(o.SmtpConnectionProperties) {
		var ret []string
		return ret
	}
	return o.SmtpConnectionProperties
}

// GetSmtpConnectionPropertiesOk returns a tuple with the SmtpConnectionProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSmtpExternalServerRequest) GetSmtpConnectionPropertiesOk() ([]string, bool) {
	if o == nil || IsNil(o.SmtpConnectionProperties) {
		return nil, false
	}
	return o.SmtpConnectionProperties, true
}

// HasSmtpConnectionProperties returns a boolean if a field has been set.
func (o *AddSmtpExternalServerRequest) HasSmtpConnectionProperties() bool {
	if o != nil && !IsNil(o.SmtpConnectionProperties) {
		return true
	}

	return false
}

// SetSmtpConnectionProperties gets a reference to the given []string and assigns it to the SmtpConnectionProperties field.
func (o *AddSmtpExternalServerRequest) SetSmtpConnectionProperties(v []string) {
	o.SmtpConnectionProperties = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddSmtpExternalServerRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSmtpExternalServerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddSmtpExternalServerRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddSmtpExternalServerRequest) SetDescription(v string) {
	o.Description = &v
}

// GetServerName returns the ServerName field value
func (o *AddSmtpExternalServerRequest) GetServerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value
// and a boolean to check if the value has been set.
func (o *AddSmtpExternalServerRequest) GetServerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerName, true
}

// SetServerName sets field value
func (o *AddSmtpExternalServerRequest) SetServerName(v string) {
	o.ServerName = v
}

func (o AddSmtpExternalServerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddSmtpExternalServerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	toSerialize["serverName"] = o.ServerName
	return toSerialize, nil
}

type NullableAddSmtpExternalServerRequest struct {
	value *AddSmtpExternalServerRequest
	isSet bool
}

func (v NullableAddSmtpExternalServerRequest) Get() *AddSmtpExternalServerRequest {
	return v.value
}

func (v *NullableAddSmtpExternalServerRequest) Set(val *AddSmtpExternalServerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSmtpExternalServerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSmtpExternalServerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSmtpExternalServerRequest(val *AddSmtpExternalServerRequest) *NullableAddSmtpExternalServerRequest {
	return &NullableAddSmtpExternalServerRequest{value: val, isSet: true}
}

func (v NullableAddSmtpExternalServerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSmtpExternalServerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
