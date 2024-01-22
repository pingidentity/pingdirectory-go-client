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

// checks if the AddVaultExternalServerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddVaultExternalServerRequest{}

// AddVaultExternalServerRequest struct for AddVaultExternalServerRequest
type AddVaultExternalServerRequest struct {
	Schemas []EnumvaultExternalServerSchemaUrn `json:"schemas"`
	// The base URL needed to access the Vault server. The base URL should consist of the protocol (\"http\" or \"https\"), the server address (resolvable name or IP address), and the port number. For example, \"https://vault.example.com:8200/\".
	VaultServerBaseURI []string `json:"vaultServerBaseURI"`
	// The mechanism used to authenticate to the Vault server.
	VaultAuthenticationMethod string `json:"vaultAuthenticationMethod"`
	// The maximum length of time to wait to obtain an HTTP connection.
	HttpConnectTimeout *string `json:"httpConnectTimeout,omitempty"`
	// The maximum length of time to wait for a response to an HTTP request.
	HttpResponseTimeout *string `json:"httpResponseTimeout,omitempty"`
	// The path to a file containing the information needed to trust the certificate presented by the Vault servers.
	TrustStoreFile *string `json:"trustStoreFile,omitempty"`
	// The passphrase needed to access the contents of the trust store. This is only required if a trust store file is required, and if that trust store requires a PIN to access its contents.
	TrustStorePin *string `json:"trustStorePin,omitempty"`
	// The store type for the specified trust store file. The value should likely be one of \"JKS\", \"PKCS12\", or \"BCFKS\".
	TrustStoreType *string `json:"trustStoreType,omitempty"`
	// A description for this External Server
	Description *string `json:"description,omitempty"`
	// Name of the new External Server
	ServerName string `json:"serverName"`
}

// NewAddVaultExternalServerRequest instantiates a new AddVaultExternalServerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddVaultExternalServerRequest(schemas []EnumvaultExternalServerSchemaUrn, vaultServerBaseURI []string, vaultAuthenticationMethod string, serverName string) *AddVaultExternalServerRequest {
	this := AddVaultExternalServerRequest{}
	this.Schemas = schemas
	this.VaultServerBaseURI = vaultServerBaseURI
	this.VaultAuthenticationMethod = vaultAuthenticationMethod
	this.ServerName = serverName
	return &this
}

// NewAddVaultExternalServerRequestWithDefaults instantiates a new AddVaultExternalServerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddVaultExternalServerRequestWithDefaults() *AddVaultExternalServerRequest {
	this := AddVaultExternalServerRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddVaultExternalServerRequest) GetSchemas() []EnumvaultExternalServerSchemaUrn {
	if o == nil {
		var ret []EnumvaultExternalServerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddVaultExternalServerRequest) GetSchemasOk() ([]EnumvaultExternalServerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddVaultExternalServerRequest) SetSchemas(v []EnumvaultExternalServerSchemaUrn) {
	o.Schemas = v
}

// GetVaultServerBaseURI returns the VaultServerBaseURI field value
func (o *AddVaultExternalServerRequest) GetVaultServerBaseURI() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.VaultServerBaseURI
}

// GetVaultServerBaseURIOk returns a tuple with the VaultServerBaseURI field value
// and a boolean to check if the value has been set.
func (o *AddVaultExternalServerRequest) GetVaultServerBaseURIOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VaultServerBaseURI, true
}

// SetVaultServerBaseURI sets field value
func (o *AddVaultExternalServerRequest) SetVaultServerBaseURI(v []string) {
	o.VaultServerBaseURI = v
}

// GetVaultAuthenticationMethod returns the VaultAuthenticationMethod field value
func (o *AddVaultExternalServerRequest) GetVaultAuthenticationMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultAuthenticationMethod
}

// GetVaultAuthenticationMethodOk returns a tuple with the VaultAuthenticationMethod field value
// and a boolean to check if the value has been set.
func (o *AddVaultExternalServerRequest) GetVaultAuthenticationMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VaultAuthenticationMethod, true
}

// SetVaultAuthenticationMethod sets field value
func (o *AddVaultExternalServerRequest) SetVaultAuthenticationMethod(v string) {
	o.VaultAuthenticationMethod = v
}

// GetHttpConnectTimeout returns the HttpConnectTimeout field value if set, zero value otherwise.
func (o *AddVaultExternalServerRequest) GetHttpConnectTimeout() string {
	if o == nil || IsNil(o.HttpConnectTimeout) {
		var ret string
		return ret
	}
	return *o.HttpConnectTimeout
}

// GetHttpConnectTimeoutOk returns a tuple with the HttpConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVaultExternalServerRequest) GetHttpConnectTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.HttpConnectTimeout) {
		return nil, false
	}
	return o.HttpConnectTimeout, true
}

// HasHttpConnectTimeout returns a boolean if a field has been set.
func (o *AddVaultExternalServerRequest) HasHttpConnectTimeout() bool {
	if o != nil && !IsNil(o.HttpConnectTimeout) {
		return true
	}

	return false
}

// SetHttpConnectTimeout gets a reference to the given string and assigns it to the HttpConnectTimeout field.
func (o *AddVaultExternalServerRequest) SetHttpConnectTimeout(v string) {
	o.HttpConnectTimeout = &v
}

// GetHttpResponseTimeout returns the HttpResponseTimeout field value if set, zero value otherwise.
func (o *AddVaultExternalServerRequest) GetHttpResponseTimeout() string {
	if o == nil || IsNil(o.HttpResponseTimeout) {
		var ret string
		return ret
	}
	return *o.HttpResponseTimeout
}

// GetHttpResponseTimeoutOk returns a tuple with the HttpResponseTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVaultExternalServerRequest) GetHttpResponseTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.HttpResponseTimeout) {
		return nil, false
	}
	return o.HttpResponseTimeout, true
}

// HasHttpResponseTimeout returns a boolean if a field has been set.
func (o *AddVaultExternalServerRequest) HasHttpResponseTimeout() bool {
	if o != nil && !IsNil(o.HttpResponseTimeout) {
		return true
	}

	return false
}

// SetHttpResponseTimeout gets a reference to the given string and assigns it to the HttpResponseTimeout field.
func (o *AddVaultExternalServerRequest) SetHttpResponseTimeout(v string) {
	o.HttpResponseTimeout = &v
}

// GetTrustStoreFile returns the TrustStoreFile field value if set, zero value otherwise.
func (o *AddVaultExternalServerRequest) GetTrustStoreFile() string {
	if o == nil || IsNil(o.TrustStoreFile) {
		var ret string
		return ret
	}
	return *o.TrustStoreFile
}

// GetTrustStoreFileOk returns a tuple with the TrustStoreFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVaultExternalServerRequest) GetTrustStoreFileOk() (*string, bool) {
	if o == nil || IsNil(o.TrustStoreFile) {
		return nil, false
	}
	return o.TrustStoreFile, true
}

// HasTrustStoreFile returns a boolean if a field has been set.
func (o *AddVaultExternalServerRequest) HasTrustStoreFile() bool {
	if o != nil && !IsNil(o.TrustStoreFile) {
		return true
	}

	return false
}

// SetTrustStoreFile gets a reference to the given string and assigns it to the TrustStoreFile field.
func (o *AddVaultExternalServerRequest) SetTrustStoreFile(v string) {
	o.TrustStoreFile = &v
}

// GetTrustStorePin returns the TrustStorePin field value if set, zero value otherwise.
func (o *AddVaultExternalServerRequest) GetTrustStorePin() string {
	if o == nil || IsNil(o.TrustStorePin) {
		var ret string
		return ret
	}
	return *o.TrustStorePin
}

// GetTrustStorePinOk returns a tuple with the TrustStorePin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVaultExternalServerRequest) GetTrustStorePinOk() (*string, bool) {
	if o == nil || IsNil(o.TrustStorePin) {
		return nil, false
	}
	return o.TrustStorePin, true
}

// HasTrustStorePin returns a boolean if a field has been set.
func (o *AddVaultExternalServerRequest) HasTrustStorePin() bool {
	if o != nil && !IsNil(o.TrustStorePin) {
		return true
	}

	return false
}

// SetTrustStorePin gets a reference to the given string and assigns it to the TrustStorePin field.
func (o *AddVaultExternalServerRequest) SetTrustStorePin(v string) {
	o.TrustStorePin = &v
}

// GetTrustStoreType returns the TrustStoreType field value if set, zero value otherwise.
func (o *AddVaultExternalServerRequest) GetTrustStoreType() string {
	if o == nil || IsNil(o.TrustStoreType) {
		var ret string
		return ret
	}
	return *o.TrustStoreType
}

// GetTrustStoreTypeOk returns a tuple with the TrustStoreType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVaultExternalServerRequest) GetTrustStoreTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TrustStoreType) {
		return nil, false
	}
	return o.TrustStoreType, true
}

// HasTrustStoreType returns a boolean if a field has been set.
func (o *AddVaultExternalServerRequest) HasTrustStoreType() bool {
	if o != nil && !IsNil(o.TrustStoreType) {
		return true
	}

	return false
}

// SetTrustStoreType gets a reference to the given string and assigns it to the TrustStoreType field.
func (o *AddVaultExternalServerRequest) SetTrustStoreType(v string) {
	o.TrustStoreType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddVaultExternalServerRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVaultExternalServerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddVaultExternalServerRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddVaultExternalServerRequest) SetDescription(v string) {
	o.Description = &v
}

// GetServerName returns the ServerName field value
func (o *AddVaultExternalServerRequest) GetServerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value
// and a boolean to check if the value has been set.
func (o *AddVaultExternalServerRequest) GetServerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerName, true
}

// SetServerName sets field value
func (o *AddVaultExternalServerRequest) SetServerName(v string) {
	o.ServerName = v
}

func (o AddVaultExternalServerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddVaultExternalServerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["vaultServerBaseURI"] = o.VaultServerBaseURI
	toSerialize["vaultAuthenticationMethod"] = o.VaultAuthenticationMethod
	if !IsNil(o.HttpConnectTimeout) {
		toSerialize["httpConnectTimeout"] = o.HttpConnectTimeout
	}
	if !IsNil(o.HttpResponseTimeout) {
		toSerialize["httpResponseTimeout"] = o.HttpResponseTimeout
	}
	if !IsNil(o.TrustStoreFile) {
		toSerialize["trustStoreFile"] = o.TrustStoreFile
	}
	if !IsNil(o.TrustStorePin) {
		toSerialize["trustStorePin"] = o.TrustStorePin
	}
	if !IsNil(o.TrustStoreType) {
		toSerialize["trustStoreType"] = o.TrustStoreType
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["serverName"] = o.ServerName
	return toSerialize, nil
}

type NullableAddVaultExternalServerRequest struct {
	value *AddVaultExternalServerRequest
	isSet bool
}

func (v NullableAddVaultExternalServerRequest) Get() *AddVaultExternalServerRequest {
	return v.value
}

func (v *NullableAddVaultExternalServerRequest) Set(val *AddVaultExternalServerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddVaultExternalServerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddVaultExternalServerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddVaultExternalServerRequest(val *AddVaultExternalServerRequest) *NullableAddVaultExternalServerRequest {
	return &NullableAddVaultExternalServerRequest{value: val, isSet: true}
}

func (v NullableAddVaultExternalServerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddVaultExternalServerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
