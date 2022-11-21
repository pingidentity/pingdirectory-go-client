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

// JmxConnectionHandlerResponse struct for JmxConnectionHandlerResponse
type JmxConnectionHandlerResponse struct {
	// Name of the Connection Handler
	Id string `json:"id"`
	Schemas []EnumjmxConnectionHandlerSchemaUrn `json:"schemas"`
	// Specifies the port number on which the JMX Connection Handler will listen for connections from clients.
	ListenPort int32 `json:"listenPort"`
	// Indicates whether the JMX Connection Handler should use SSL.
	UseSSL *bool `json:"useSSL,omitempty"`
	// Specifies the nickname (also called the alias) of the certificate that the JMX Connection Handler should use when performing SSL communication.
	SslCertNickname *string `json:"sslCertNickname,omitempty"`
	// Specifies the name of the key manager that should be used with this JMX Connection Handler .
	KeyManagerProvider *string `json:"keyManagerProvider,omitempty"`
	// A description for this Connection Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the Connection Handler is enabled.
	Enabled bool `json:"enabled"`
	// Specifies a set of address masks that determines the addresses of the clients that are allowed to establish connections to this connection handler.
	AllowedClient []string `json:"allowedClient,omitempty"`
	// Specifies a set of address masks that determines the addresses of the clients that are not allowed to establish connections to this connection handler.
	DeniedClient []string `json:"deniedClient,omitempty"`
	Meta *MetaMeta `json:"meta,omitempty"`
}

// NewJmxConnectionHandlerResponse instantiates a new JmxConnectionHandlerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJmxConnectionHandlerResponse(id string, schemas []EnumjmxConnectionHandlerSchemaUrn, listenPort int32, enabled bool) *JmxConnectionHandlerResponse {
	this := JmxConnectionHandlerResponse{}
	this.Id = id
	this.Schemas = schemas
	this.ListenPort = listenPort
	this.Enabled = enabled
	return &this
}

// NewJmxConnectionHandlerResponseWithDefaults instantiates a new JmxConnectionHandlerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJmxConnectionHandlerResponseWithDefaults() *JmxConnectionHandlerResponse {
	this := JmxConnectionHandlerResponse{}
	return &this
}

// GetId returns the Id field value
func (o *JmxConnectionHandlerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JmxConnectionHandlerResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JmxConnectionHandlerResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *JmxConnectionHandlerResponse) GetSchemas() []EnumjmxConnectionHandlerSchemaUrn {
	if o == nil {
		var ret []EnumjmxConnectionHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *JmxConnectionHandlerResponse) GetSchemasOk() ([]EnumjmxConnectionHandlerSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *JmxConnectionHandlerResponse) SetSchemas(v []EnumjmxConnectionHandlerSchemaUrn) {
	o.Schemas = v
}

// GetListenPort returns the ListenPort field value
func (o *JmxConnectionHandlerResponse) GetListenPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ListenPort
}

// GetListenPortOk returns a tuple with the ListenPort field value
// and a boolean to check if the value has been set.
func (o *JmxConnectionHandlerResponse) GetListenPortOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ListenPort, true
}

// SetListenPort sets field value
func (o *JmxConnectionHandlerResponse) SetListenPort(v int32) {
	o.ListenPort = v
}

// GetUseSSL returns the UseSSL field value if set, zero value otherwise.
func (o *JmxConnectionHandlerResponse) GetUseSSL() bool {
	if o == nil || isNil(o.UseSSL) {
		var ret bool
		return ret
	}
	return *o.UseSSL
}

// GetUseSSLOk returns a tuple with the UseSSL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JmxConnectionHandlerResponse) GetUseSSLOk() (*bool, bool) {
	if o == nil || isNil(o.UseSSL) {
    return nil, false
	}
	return o.UseSSL, true
}

// HasUseSSL returns a boolean if a field has been set.
func (o *JmxConnectionHandlerResponse) HasUseSSL() bool {
	if o != nil && !isNil(o.UseSSL) {
		return true
	}

	return false
}

// SetUseSSL gets a reference to the given bool and assigns it to the UseSSL field.
func (o *JmxConnectionHandlerResponse) SetUseSSL(v bool) {
	o.UseSSL = &v
}

// GetSslCertNickname returns the SslCertNickname field value if set, zero value otherwise.
func (o *JmxConnectionHandlerResponse) GetSslCertNickname() string {
	if o == nil || isNil(o.SslCertNickname) {
		var ret string
		return ret
	}
	return *o.SslCertNickname
}

// GetSslCertNicknameOk returns a tuple with the SslCertNickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JmxConnectionHandlerResponse) GetSslCertNicknameOk() (*string, bool) {
	if o == nil || isNil(o.SslCertNickname) {
    return nil, false
	}
	return o.SslCertNickname, true
}

// HasSslCertNickname returns a boolean if a field has been set.
func (o *JmxConnectionHandlerResponse) HasSslCertNickname() bool {
	if o != nil && !isNil(o.SslCertNickname) {
		return true
	}

	return false
}

// SetSslCertNickname gets a reference to the given string and assigns it to the SslCertNickname field.
func (o *JmxConnectionHandlerResponse) SetSslCertNickname(v string) {
	o.SslCertNickname = &v
}

// GetKeyManagerProvider returns the KeyManagerProvider field value if set, zero value otherwise.
func (o *JmxConnectionHandlerResponse) GetKeyManagerProvider() string {
	if o == nil || isNil(o.KeyManagerProvider) {
		var ret string
		return ret
	}
	return *o.KeyManagerProvider
}

// GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JmxConnectionHandlerResponse) GetKeyManagerProviderOk() (*string, bool) {
	if o == nil || isNil(o.KeyManagerProvider) {
    return nil, false
	}
	return o.KeyManagerProvider, true
}

// HasKeyManagerProvider returns a boolean if a field has been set.
func (o *JmxConnectionHandlerResponse) HasKeyManagerProvider() bool {
	if o != nil && !isNil(o.KeyManagerProvider) {
		return true
	}

	return false
}

// SetKeyManagerProvider gets a reference to the given string and assigns it to the KeyManagerProvider field.
func (o *JmxConnectionHandlerResponse) SetKeyManagerProvider(v string) {
	o.KeyManagerProvider = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *JmxConnectionHandlerResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JmxConnectionHandlerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *JmxConnectionHandlerResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *JmxConnectionHandlerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *JmxConnectionHandlerResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *JmxConnectionHandlerResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *JmxConnectionHandlerResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAllowedClient returns the AllowedClient field value if set, zero value otherwise.
func (o *JmxConnectionHandlerResponse) GetAllowedClient() []string {
	if o == nil || isNil(o.AllowedClient) {
		var ret []string
		return ret
	}
	return o.AllowedClient
}

// GetAllowedClientOk returns a tuple with the AllowedClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JmxConnectionHandlerResponse) GetAllowedClientOk() ([]string, bool) {
	if o == nil || isNil(o.AllowedClient) {
    return nil, false
	}
	return o.AllowedClient, true
}

// HasAllowedClient returns a boolean if a field has been set.
func (o *JmxConnectionHandlerResponse) HasAllowedClient() bool {
	if o != nil && !isNil(o.AllowedClient) {
		return true
	}

	return false
}

// SetAllowedClient gets a reference to the given []string and assigns it to the AllowedClient field.
func (o *JmxConnectionHandlerResponse) SetAllowedClient(v []string) {
	o.AllowedClient = v
}

// GetDeniedClient returns the DeniedClient field value if set, zero value otherwise.
func (o *JmxConnectionHandlerResponse) GetDeniedClient() []string {
	if o == nil || isNil(o.DeniedClient) {
		var ret []string
		return ret
	}
	return o.DeniedClient
}

// GetDeniedClientOk returns a tuple with the DeniedClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JmxConnectionHandlerResponse) GetDeniedClientOk() ([]string, bool) {
	if o == nil || isNil(o.DeniedClient) {
    return nil, false
	}
	return o.DeniedClient, true
}

// HasDeniedClient returns a boolean if a field has been set.
func (o *JmxConnectionHandlerResponse) HasDeniedClient() bool {
	if o != nil && !isNil(o.DeniedClient) {
		return true
	}

	return false
}

// SetDeniedClient gets a reference to the given []string and assigns it to the DeniedClient field.
func (o *JmxConnectionHandlerResponse) SetDeniedClient(v []string) {
	o.DeniedClient = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *JmxConnectionHandlerResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JmxConnectionHandlerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *JmxConnectionHandlerResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *JmxConnectionHandlerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

func (o JmxConnectionHandlerResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["listenPort"] = o.ListenPort
	}
	if !isNil(o.UseSSL) {
		toSerialize["useSSL"] = o.UseSSL
	}
	if !isNil(o.SslCertNickname) {
		toSerialize["sslCertNickname"] = o.SslCertNickname
	}
	if !isNil(o.KeyManagerProvider) {
		toSerialize["keyManagerProvider"] = o.KeyManagerProvider
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.AllowedClient) {
		toSerialize["allowedClient"] = o.AllowedClient
	}
	if !isNil(o.DeniedClient) {
		toSerialize["deniedClient"] = o.DeniedClient
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableJmxConnectionHandlerResponse struct {
	value *JmxConnectionHandlerResponse
	isSet bool
}

func (v NullableJmxConnectionHandlerResponse) Get() *JmxConnectionHandlerResponse {
	return v.value
}

func (v *NullableJmxConnectionHandlerResponse) Set(val *JmxConnectionHandlerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJmxConnectionHandlerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJmxConnectionHandlerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJmxConnectionHandlerResponse(val *JmxConnectionHandlerResponse) *NullableJmxConnectionHandlerResponse {
	return &NullableJmxConnectionHandlerResponse{value: val, isSet: true}
}

func (v NullableJmxConnectionHandlerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJmxConnectionHandlerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


