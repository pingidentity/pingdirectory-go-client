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

// AzureKeyVaultCipherStreamProviderResponse struct for AzureKeyVaultCipherStreamProviderResponse
type AzureKeyVaultCipherStreamProviderResponse struct {
	// Name of the Cipher Stream Provider
	Id string `json:"id"`
	Schemas []EnumazureKeyVaultCipherStreamProviderSchemaUrn `json:"schemas"`
	// The URI that identifies the Azure Key Vault from which the secret is to be retrieved.
	KeyVaultURI string `json:"keyVaultURI"`
	// The mechanism used to authenticate to the Azure service.
	AzureAuthenticationMethod string `json:"azureAuthenticationMethod"`
	// The name of the secret to retrieve.
	SecretName string `json:"secretName"`
	// The path to a file that will hold metadata about the encryption performed by this Azure Key Vault Cipher Stream Provider.
	EncryptionMetadataFile string `json:"encryptionMetadataFile"`
	// A description for this Cipher Stream Provider
	Description *string `json:"description,omitempty"`
	// Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server.
	Enabled bool `json:"enabled"`
}

// NewAzureKeyVaultCipherStreamProviderResponse instantiates a new AzureKeyVaultCipherStreamProviderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureKeyVaultCipherStreamProviderResponse(id string, schemas []EnumazureKeyVaultCipherStreamProviderSchemaUrn, keyVaultURI string, azureAuthenticationMethod string, secretName string, encryptionMetadataFile string, enabled bool) *AzureKeyVaultCipherStreamProviderResponse {
	this := AzureKeyVaultCipherStreamProviderResponse{}
	this.Id = id
	this.Schemas = schemas
	this.KeyVaultURI = keyVaultURI
	this.AzureAuthenticationMethod = azureAuthenticationMethod
	this.SecretName = secretName
	this.EncryptionMetadataFile = encryptionMetadataFile
	this.Enabled = enabled
	return &this
}

// NewAzureKeyVaultCipherStreamProviderResponseWithDefaults instantiates a new AzureKeyVaultCipherStreamProviderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureKeyVaultCipherStreamProviderResponseWithDefaults() *AzureKeyVaultCipherStreamProviderResponse {
	this := AzureKeyVaultCipherStreamProviderResponse{}
	return &this
}

// GetId returns the Id field value
func (o *AzureKeyVaultCipherStreamProviderResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultCipherStreamProviderResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AzureKeyVaultCipherStreamProviderResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *AzureKeyVaultCipherStreamProviderResponse) GetSchemas() []EnumazureKeyVaultCipherStreamProviderSchemaUrn {
	if o == nil {
		var ret []EnumazureKeyVaultCipherStreamProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultCipherStreamProviderResponse) GetSchemasOk() ([]EnumazureKeyVaultCipherStreamProviderSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AzureKeyVaultCipherStreamProviderResponse) SetSchemas(v []EnumazureKeyVaultCipherStreamProviderSchemaUrn) {
	o.Schemas = v
}

// GetKeyVaultURI returns the KeyVaultURI field value
func (o *AzureKeyVaultCipherStreamProviderResponse) GetKeyVaultURI() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyVaultURI
}

// GetKeyVaultURIOk returns a tuple with the KeyVaultURI field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultCipherStreamProviderResponse) GetKeyVaultURIOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.KeyVaultURI, true
}

// SetKeyVaultURI sets field value
func (o *AzureKeyVaultCipherStreamProviderResponse) SetKeyVaultURI(v string) {
	o.KeyVaultURI = v
}

// GetAzureAuthenticationMethod returns the AzureAuthenticationMethod field value
func (o *AzureKeyVaultCipherStreamProviderResponse) GetAzureAuthenticationMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AzureAuthenticationMethod
}

// GetAzureAuthenticationMethodOk returns a tuple with the AzureAuthenticationMethod field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultCipherStreamProviderResponse) GetAzureAuthenticationMethodOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AzureAuthenticationMethod, true
}

// SetAzureAuthenticationMethod sets field value
func (o *AzureKeyVaultCipherStreamProviderResponse) SetAzureAuthenticationMethod(v string) {
	o.AzureAuthenticationMethod = v
}

// GetSecretName returns the SecretName field value
func (o *AzureKeyVaultCipherStreamProviderResponse) GetSecretName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretName
}

// GetSecretNameOk returns a tuple with the SecretName field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultCipherStreamProviderResponse) GetSecretNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SecretName, true
}

// SetSecretName sets field value
func (o *AzureKeyVaultCipherStreamProviderResponse) SetSecretName(v string) {
	o.SecretName = v
}

// GetEncryptionMetadataFile returns the EncryptionMetadataFile field value
func (o *AzureKeyVaultCipherStreamProviderResponse) GetEncryptionMetadataFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EncryptionMetadataFile
}

// GetEncryptionMetadataFileOk returns a tuple with the EncryptionMetadataFile field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultCipherStreamProviderResponse) GetEncryptionMetadataFileOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EncryptionMetadataFile, true
}

// SetEncryptionMetadataFile sets field value
func (o *AzureKeyVaultCipherStreamProviderResponse) SetEncryptionMetadataFile(v string) {
	o.EncryptionMetadataFile = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AzureKeyVaultCipherStreamProviderResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultCipherStreamProviderResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AzureKeyVaultCipherStreamProviderResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AzureKeyVaultCipherStreamProviderResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AzureKeyVaultCipherStreamProviderResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultCipherStreamProviderResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AzureKeyVaultCipherStreamProviderResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AzureKeyVaultCipherStreamProviderResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["keyVaultURI"] = o.KeyVaultURI
	}
	if true {
		toSerialize["azureAuthenticationMethod"] = o.AzureAuthenticationMethod
	}
	if true {
		toSerialize["secretName"] = o.SecretName
	}
	if true {
		toSerialize["encryptionMetadataFile"] = o.EncryptionMetadataFile
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableAzureKeyVaultCipherStreamProviderResponse struct {
	value *AzureKeyVaultCipherStreamProviderResponse
	isSet bool
}

func (v NullableAzureKeyVaultCipherStreamProviderResponse) Get() *AzureKeyVaultCipherStreamProviderResponse {
	return v.value
}

func (v *NullableAzureKeyVaultCipherStreamProviderResponse) Set(val *AzureKeyVaultCipherStreamProviderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureKeyVaultCipherStreamProviderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureKeyVaultCipherStreamProviderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureKeyVaultCipherStreamProviderResponse(val *AzureKeyVaultCipherStreamProviderResponse) *NullableAzureKeyVaultCipherStreamProviderResponse {
	return &NullableAzureKeyVaultCipherStreamProviderResponse{value: val, isSet: true}
}

func (v NullableAzureKeyVaultCipherStreamProviderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureKeyVaultCipherStreamProviderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


