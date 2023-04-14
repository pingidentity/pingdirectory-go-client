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

// checks if the AddJmxConnectionHandlerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddJmxConnectionHandlerRequest{}

// AddJmxConnectionHandlerRequest struct for AddJmxConnectionHandlerRequest
type AddJmxConnectionHandlerRequest struct {
	// Name of the new Connection Handler
	HandlerName string                              `json:"handlerName"`
	Schemas     []EnumjmxConnectionHandlerSchemaUrn `json:"schemas"`
	// Specifies the port number on which the JMX Connection Handler will listen for connections from clients.
	ListenPort int64 `json:"listenPort"`
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
}

// NewAddJmxConnectionHandlerRequest instantiates a new AddJmxConnectionHandlerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddJmxConnectionHandlerRequest(handlerName string, schemas []EnumjmxConnectionHandlerSchemaUrn, listenPort int64, enabled bool) *AddJmxConnectionHandlerRequest {
	this := AddJmxConnectionHandlerRequest{}
	this.HandlerName = handlerName
	this.Schemas = schemas
	this.ListenPort = listenPort
	this.Enabled = enabled
	return &this
}

// NewAddJmxConnectionHandlerRequestWithDefaults instantiates a new AddJmxConnectionHandlerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddJmxConnectionHandlerRequestWithDefaults() *AddJmxConnectionHandlerRequest {
	this := AddJmxConnectionHandlerRequest{}
	return &this
}

// GetHandlerName returns the HandlerName field value
func (o *AddJmxConnectionHandlerRequest) GetHandlerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HandlerName
}

// GetHandlerNameOk returns a tuple with the HandlerName field value
// and a boolean to check if the value has been set.
func (o *AddJmxConnectionHandlerRequest) GetHandlerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HandlerName, true
}

// SetHandlerName sets field value
func (o *AddJmxConnectionHandlerRequest) SetHandlerName(v string) {
	o.HandlerName = v
}

// GetSchemas returns the Schemas field value
func (o *AddJmxConnectionHandlerRequest) GetSchemas() []EnumjmxConnectionHandlerSchemaUrn {
	if o == nil {
		var ret []EnumjmxConnectionHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddJmxConnectionHandlerRequest) GetSchemasOk() ([]EnumjmxConnectionHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddJmxConnectionHandlerRequest) SetSchemas(v []EnumjmxConnectionHandlerSchemaUrn) {
	o.Schemas = v
}

// GetListenPort returns the ListenPort field value
func (o *AddJmxConnectionHandlerRequest) GetListenPort() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ListenPort
}

// GetListenPortOk returns a tuple with the ListenPort field value
// and a boolean to check if the value has been set.
func (o *AddJmxConnectionHandlerRequest) GetListenPortOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListenPort, true
}

// SetListenPort sets field value
func (o *AddJmxConnectionHandlerRequest) SetListenPort(v int64) {
	o.ListenPort = v
}

// GetUseSSL returns the UseSSL field value if set, zero value otherwise.
func (o *AddJmxConnectionHandlerRequest) GetUseSSL() bool {
	if o == nil || IsNil(o.UseSSL) {
		var ret bool
		return ret
	}
	return *o.UseSSL
}

// GetUseSSLOk returns a tuple with the UseSSL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJmxConnectionHandlerRequest) GetUseSSLOk() (*bool, bool) {
	if o == nil || IsNil(o.UseSSL) {
		return nil, false
	}
	return o.UseSSL, true
}

// HasUseSSL returns a boolean if a field has been set.
func (o *AddJmxConnectionHandlerRequest) HasUseSSL() bool {
	if o != nil && !IsNil(o.UseSSL) {
		return true
	}

	return false
}

// SetUseSSL gets a reference to the given bool and assigns it to the UseSSL field.
func (o *AddJmxConnectionHandlerRequest) SetUseSSL(v bool) {
	o.UseSSL = &v
}

// GetSslCertNickname returns the SslCertNickname field value if set, zero value otherwise.
func (o *AddJmxConnectionHandlerRequest) GetSslCertNickname() string {
	if o == nil || IsNil(o.SslCertNickname) {
		var ret string
		return ret
	}
	return *o.SslCertNickname
}

// GetSslCertNicknameOk returns a tuple with the SslCertNickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJmxConnectionHandlerRequest) GetSslCertNicknameOk() (*string, bool) {
	if o == nil || IsNil(o.SslCertNickname) {
		return nil, false
	}
	return o.SslCertNickname, true
}

// HasSslCertNickname returns a boolean if a field has been set.
func (o *AddJmxConnectionHandlerRequest) HasSslCertNickname() bool {
	if o != nil && !IsNil(o.SslCertNickname) {
		return true
	}

	return false
}

// SetSslCertNickname gets a reference to the given string and assigns it to the SslCertNickname field.
func (o *AddJmxConnectionHandlerRequest) SetSslCertNickname(v string) {
	o.SslCertNickname = &v
}

// GetKeyManagerProvider returns the KeyManagerProvider field value if set, zero value otherwise.
func (o *AddJmxConnectionHandlerRequest) GetKeyManagerProvider() string {
	if o == nil || IsNil(o.KeyManagerProvider) {
		var ret string
		return ret
	}
	return *o.KeyManagerProvider
}

// GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJmxConnectionHandlerRequest) GetKeyManagerProviderOk() (*string, bool) {
	if o == nil || IsNil(o.KeyManagerProvider) {
		return nil, false
	}
	return o.KeyManagerProvider, true
}

// HasKeyManagerProvider returns a boolean if a field has been set.
func (o *AddJmxConnectionHandlerRequest) HasKeyManagerProvider() bool {
	if o != nil && !IsNil(o.KeyManagerProvider) {
		return true
	}

	return false
}

// SetKeyManagerProvider gets a reference to the given string and assigns it to the KeyManagerProvider field.
func (o *AddJmxConnectionHandlerRequest) SetKeyManagerProvider(v string) {
	o.KeyManagerProvider = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddJmxConnectionHandlerRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJmxConnectionHandlerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddJmxConnectionHandlerRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddJmxConnectionHandlerRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddJmxConnectionHandlerRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddJmxConnectionHandlerRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddJmxConnectionHandlerRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAllowedClient returns the AllowedClient field value if set, zero value otherwise.
func (o *AddJmxConnectionHandlerRequest) GetAllowedClient() []string {
	if o == nil || IsNil(o.AllowedClient) {
		var ret []string
		return ret
	}
	return o.AllowedClient
}

// GetAllowedClientOk returns a tuple with the AllowedClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJmxConnectionHandlerRequest) GetAllowedClientOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedClient) {
		return nil, false
	}
	return o.AllowedClient, true
}

// HasAllowedClient returns a boolean if a field has been set.
func (o *AddJmxConnectionHandlerRequest) HasAllowedClient() bool {
	if o != nil && !IsNil(o.AllowedClient) {
		return true
	}

	return false
}

// SetAllowedClient gets a reference to the given []string and assigns it to the AllowedClient field.
func (o *AddJmxConnectionHandlerRequest) SetAllowedClient(v []string) {
	o.AllowedClient = v
}

// GetDeniedClient returns the DeniedClient field value if set, zero value otherwise.
func (o *AddJmxConnectionHandlerRequest) GetDeniedClient() []string {
	if o == nil || IsNil(o.DeniedClient) {
		var ret []string
		return ret
	}
	return o.DeniedClient
}

// GetDeniedClientOk returns a tuple with the DeniedClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJmxConnectionHandlerRequest) GetDeniedClientOk() ([]string, bool) {
	if o == nil || IsNil(o.DeniedClient) {
		return nil, false
	}
	return o.DeniedClient, true
}

// HasDeniedClient returns a boolean if a field has been set.
func (o *AddJmxConnectionHandlerRequest) HasDeniedClient() bool {
	if o != nil && !IsNil(o.DeniedClient) {
		return true
	}

	return false
}

// SetDeniedClient gets a reference to the given []string and assigns it to the DeniedClient field.
func (o *AddJmxConnectionHandlerRequest) SetDeniedClient(v []string) {
	o.DeniedClient = v
}

func (o AddJmxConnectionHandlerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddJmxConnectionHandlerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["handlerName"] = o.HandlerName
	toSerialize["schemas"] = o.Schemas
	toSerialize["listenPort"] = o.ListenPort
	if !IsNil(o.UseSSL) {
		toSerialize["useSSL"] = o.UseSSL
	}
	if !IsNil(o.SslCertNickname) {
		toSerialize["sslCertNickname"] = o.SslCertNickname
	}
	if !IsNil(o.KeyManagerProvider) {
		toSerialize["keyManagerProvider"] = o.KeyManagerProvider
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.AllowedClient) {
		toSerialize["allowedClient"] = o.AllowedClient
	}
	if !IsNil(o.DeniedClient) {
		toSerialize["deniedClient"] = o.DeniedClient
	}
	return toSerialize, nil
}

type NullableAddJmxConnectionHandlerRequest struct {
	value *AddJmxConnectionHandlerRequest
	isSet bool
}

func (v NullableAddJmxConnectionHandlerRequest) Get() *AddJmxConnectionHandlerRequest {
	return v.value
}

func (v *NullableAddJmxConnectionHandlerRequest) Set(val *AddJmxConnectionHandlerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddJmxConnectionHandlerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddJmxConnectionHandlerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddJmxConnectionHandlerRequest(val *AddJmxConnectionHandlerRequest) *NullableAddJmxConnectionHandlerRequest {
	return &NullableAddJmxConnectionHandlerRequest{value: val, isSet: true}
}

func (v NullableAddJmxConnectionHandlerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddJmxConnectionHandlerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
