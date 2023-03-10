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

// checks if the AddAzureKeyVaultCipherStreamProviderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddAzureKeyVaultCipherStreamProviderRequest{}

// AddAzureKeyVaultCipherStreamProviderRequest struct for AddAzureKeyVaultCipherStreamProviderRequest
type AddAzureKeyVaultCipherStreamProviderRequest struct {
	// Name of the new Cipher Stream Provider
	ProviderName string                                           `json:"providerName"`
	Schemas      []EnumazureKeyVaultCipherStreamProviderSchemaUrn `json:"schemas"`
	// The URI that identifies the Azure Key Vault from which the secret is to be retrieved.
	KeyVaultURI string `json:"keyVaultURI"`
	// The mechanism used to authenticate to the Azure service.
	AzureAuthenticationMethod string `json:"azureAuthenticationMethod"`
	// A reference to an HTTP proxy server that should be used for requests sent to the Azure service.
	HttpProxyExternalServer *string `json:"httpProxyExternalServer,omitempty"`
	// The name of the secret to retrieve.
	SecretName string `json:"secretName"`
	// The path to a file that will hold metadata about the encryption performed by this Azure Key Vault Cipher Stream Provider.
	EncryptionMetadataFile *string `json:"encryptionMetadataFile,omitempty"`
	// A description for this Cipher Stream Provider
	Description *string `json:"description,omitempty"`
	// Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server.
	Enabled bool `json:"enabled"`
}

// NewAddAzureKeyVaultCipherStreamProviderRequest instantiates a new AddAzureKeyVaultCipherStreamProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddAzureKeyVaultCipherStreamProviderRequest(providerName string, schemas []EnumazureKeyVaultCipherStreamProviderSchemaUrn, keyVaultURI string, azureAuthenticationMethod string, secretName string, enabled bool) *AddAzureKeyVaultCipherStreamProviderRequest {
	this := AddAzureKeyVaultCipherStreamProviderRequest{}
	this.ProviderName = providerName
	this.Schemas = schemas
	this.KeyVaultURI = keyVaultURI
	this.AzureAuthenticationMethod = azureAuthenticationMethod
	this.SecretName = secretName
	this.Enabled = enabled
	return &this
}

// NewAddAzureKeyVaultCipherStreamProviderRequestWithDefaults instantiates a new AddAzureKeyVaultCipherStreamProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddAzureKeyVaultCipherStreamProviderRequestWithDefaults() *AddAzureKeyVaultCipherStreamProviderRequest {
	this := AddAzureKeyVaultCipherStreamProviderRequest{}
	return &this
}

// GetProviderName returns the ProviderName field value
func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *AddAzureKeyVaultCipherStreamProviderRequest) SetProviderName(v string) {
	o.ProviderName = v
}

// GetSchemas returns the Schemas field value
func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetSchemas() []EnumazureKeyVaultCipherStreamProviderSchemaUrn {
	if o == nil {
		var ret []EnumazureKeyVaultCipherStreamProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetSchemasOk() ([]EnumazureKeyVaultCipherStreamProviderSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddAzureKeyVaultCipherStreamProviderRequest) SetSchemas(v []EnumazureKeyVaultCipherStreamProviderSchemaUrn) {
	o.Schemas = v
}

// GetKeyVaultURI returns the KeyVaultURI field value
func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetKeyVaultURI() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyVaultURI
}

// GetKeyVaultURIOk returns a tuple with the KeyVaultURI field value
// and a boolean to check if the value has been set.
func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetKeyVaultURIOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyVaultURI, true
}

// SetKeyVaultURI sets field value
func (o *AddAzureKeyVaultCipherStreamProviderRequest) SetKeyVaultURI(v string) {
	o.KeyVaultURI = v
}

// GetAzureAuthenticationMethod returns the AzureAuthenticationMethod field value
func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetAzureAuthenticationMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AzureAuthenticationMethod
}

// GetAzureAuthenticationMethodOk returns a tuple with the AzureAuthenticationMethod field value
// and a boolean to check if the value has been set.
func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetAzureAuthenticationMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureAuthenticationMethod, true
}

// SetAzureAuthenticationMethod sets field value
func (o *AddAzureKeyVaultCipherStreamProviderRequest) SetAzureAuthenticationMethod(v string) {
	o.AzureAuthenticationMethod = v
}

// GetHttpProxyExternalServer returns the HttpProxyExternalServer field value if set, zero value otherwise.
func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetHttpProxyExternalServer() string {
	if o == nil || IsNil(o.HttpProxyExternalServer) {
		var ret string
		return ret
	}
	return *o.HttpProxyExternalServer
}

// GetHttpProxyExternalServerOk returns a tuple with the HttpProxyExternalServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetHttpProxyExternalServerOk() (*string, bool) {
	if o == nil || IsNil(o.HttpProxyExternalServer) {
		return nil, false
	}
	return o.HttpProxyExternalServer, true
}

// HasHttpProxyExternalServer returns a boolean if a field has been set.
func (o *AddAzureKeyVaultCipherStreamProviderRequest) HasHttpProxyExternalServer() bool {
	if o != nil && !IsNil(o.HttpProxyExternalServer) {
		return true
	}

	return false
}

// SetHttpProxyExternalServer gets a reference to the given string and assigns it to the HttpProxyExternalServer field.
func (o *AddAzureKeyVaultCipherStreamProviderRequest) SetHttpProxyExternalServer(v string) {
	o.HttpProxyExternalServer = &v
}

// GetSecretName returns the SecretName field value
func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetSecretName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretName
}

// GetSecretNameOk returns a tuple with the SecretName field value
// and a boolean to check if the value has been set.
func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetSecretNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretName, true
}

// SetSecretName sets field value
func (o *AddAzureKeyVaultCipherStreamProviderRequest) SetSecretName(v string) {
	o.SecretName = v
}

// GetEncryptionMetadataFile returns the EncryptionMetadataFile field value if set, zero value otherwise.
func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetEncryptionMetadataFile() string {
	if o == nil || IsNil(o.EncryptionMetadataFile) {
		var ret string
		return ret
	}
	return *o.EncryptionMetadataFile
}

// GetEncryptionMetadataFileOk returns a tuple with the EncryptionMetadataFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetEncryptionMetadataFileOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionMetadataFile) {
		return nil, false
	}
	return o.EncryptionMetadataFile, true
}

// HasEncryptionMetadataFile returns a boolean if a field has been set.
func (o *AddAzureKeyVaultCipherStreamProviderRequest) HasEncryptionMetadataFile() bool {
	if o != nil && !IsNil(o.EncryptionMetadataFile) {
		return true
	}

	return false
}

// SetEncryptionMetadataFile gets a reference to the given string and assigns it to the EncryptionMetadataFile field.
func (o *AddAzureKeyVaultCipherStreamProviderRequest) SetEncryptionMetadataFile(v string) {
	o.EncryptionMetadataFile = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddAzureKeyVaultCipherStreamProviderRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddAzureKeyVaultCipherStreamProviderRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddAzureKeyVaultCipherStreamProviderRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddAzureKeyVaultCipherStreamProviderRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddAzureKeyVaultCipherStreamProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddAzureKeyVaultCipherStreamProviderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["providerName"] = o.ProviderName
	toSerialize["schemas"] = o.Schemas
	toSerialize["keyVaultURI"] = o.KeyVaultURI
	toSerialize["azureAuthenticationMethod"] = o.AzureAuthenticationMethod
	if !IsNil(o.HttpProxyExternalServer) {
		toSerialize["httpProxyExternalServer"] = o.HttpProxyExternalServer
	}
	toSerialize["secretName"] = o.SecretName
	if !IsNil(o.EncryptionMetadataFile) {
		toSerialize["encryptionMetadataFile"] = o.EncryptionMetadataFile
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableAddAzureKeyVaultCipherStreamProviderRequest struct {
	value *AddAzureKeyVaultCipherStreamProviderRequest
	isSet bool
}

func (v NullableAddAzureKeyVaultCipherStreamProviderRequest) Get() *AddAzureKeyVaultCipherStreamProviderRequest {
	return v.value
}

func (v *NullableAddAzureKeyVaultCipherStreamProviderRequest) Set(val *AddAzureKeyVaultCipherStreamProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddAzureKeyVaultCipherStreamProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddAzureKeyVaultCipherStreamProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddAzureKeyVaultCipherStreamProviderRequest(val *AddAzureKeyVaultCipherStreamProviderRequest) *NullableAddAzureKeyVaultCipherStreamProviderRequest {
	return &NullableAddAzureKeyVaultCipherStreamProviderRequest{value: val, isSet: true}
}

func (v NullableAddAzureKeyVaultCipherStreamProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddAzureKeyVaultCipherStreamProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
