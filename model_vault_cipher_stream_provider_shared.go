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

// VaultCipherStreamProviderShared struct for VaultCipherStreamProviderShared
type VaultCipherStreamProviderShared struct {
	Schemas []EnumvaultCipherStreamProviderSchemaUrn `json:"schemas"`
	// An external server definition with information needed to connect and authenticate to the Vault server.
	VaultExternalServer *string `json:"vaultExternalServer,omitempty"`
	VaultServerBaseURI []string `json:"vaultServerBaseURI,omitempty"`
	// The mechanism used to authenticate to the Vault server.
	VaultAuthenticationMethod *string `json:"vaultAuthenticationMethod,omitempty"`
	// The path to the desired secret in the Vault service. This will be appended to the value of the base-url property for the associated Vault external server.
	VaultSecretPath string `json:"vaultSecretPath"`
	// The name of the field in the Vault secret record that contains the passphrase to use to generate the encryption key.
	VaultSecretFieldName string `json:"vaultSecretFieldName"`
	// The path to a file that will hold metadata about the encryption performed by this Vault Cipher Stream Provider.
	VaultEncryptionMetadataFile string `json:"vaultEncryptionMetadataFile"`
	// The path to a file containing the information needed to trust the certificate presented by the Vault servers.
	TrustStoreFile *string `json:"trustStoreFile,omitempty"`
	// The passphrase needed to access the contents of the trust store. This is only required if a trust store file is required, and if that trust store requires a PIN to access its contents.
	TrustStorePin *string `json:"trustStorePin,omitempty"`
	// The store type for the specified trust store file. The value should likely be one of \"JKS\" or \"PKCS12\".
	TrustStoreType *string `json:"trustStoreType,omitempty"`
	// A description for this Cipher Stream Provider
	Description *string `json:"description,omitempty"`
	// Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server.
	Enabled bool `json:"enabled"`
}

// NewVaultCipherStreamProviderShared instantiates a new VaultCipherStreamProviderShared object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVaultCipherStreamProviderShared(schemas []EnumvaultCipherStreamProviderSchemaUrn, vaultSecretPath string, vaultSecretFieldName string, vaultEncryptionMetadataFile string, enabled bool) *VaultCipherStreamProviderShared {
	this := VaultCipherStreamProviderShared{}
	this.Schemas = schemas
	this.VaultSecretPath = vaultSecretPath
	this.VaultSecretFieldName = vaultSecretFieldName
	this.VaultEncryptionMetadataFile = vaultEncryptionMetadataFile
	this.Enabled = enabled
	return &this
}

// NewVaultCipherStreamProviderSharedWithDefaults instantiates a new VaultCipherStreamProviderShared object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVaultCipherStreamProviderSharedWithDefaults() *VaultCipherStreamProviderShared {
	this := VaultCipherStreamProviderShared{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *VaultCipherStreamProviderShared) GetSchemas() []EnumvaultCipherStreamProviderSchemaUrn {
	if o == nil {
		var ret []EnumvaultCipherStreamProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *VaultCipherStreamProviderShared) GetSchemasOk() ([]EnumvaultCipherStreamProviderSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *VaultCipherStreamProviderShared) SetSchemas(v []EnumvaultCipherStreamProviderSchemaUrn) {
	o.Schemas = v
}

// GetVaultExternalServer returns the VaultExternalServer field value if set, zero value otherwise.
func (o *VaultCipherStreamProviderShared) GetVaultExternalServer() string {
	if o == nil || isNil(o.VaultExternalServer) {
		var ret string
		return ret
	}
	return *o.VaultExternalServer
}

// GetVaultExternalServerOk returns a tuple with the VaultExternalServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultCipherStreamProviderShared) GetVaultExternalServerOk() (*string, bool) {
	if o == nil || isNil(o.VaultExternalServer) {
    return nil, false
	}
	return o.VaultExternalServer, true
}

// HasVaultExternalServer returns a boolean if a field has been set.
func (o *VaultCipherStreamProviderShared) HasVaultExternalServer() bool {
	if o != nil && !isNil(o.VaultExternalServer) {
		return true
	}

	return false
}

// SetVaultExternalServer gets a reference to the given string and assigns it to the VaultExternalServer field.
func (o *VaultCipherStreamProviderShared) SetVaultExternalServer(v string) {
	o.VaultExternalServer = &v
}

// GetVaultServerBaseURI returns the VaultServerBaseURI field value if set, zero value otherwise.
func (o *VaultCipherStreamProviderShared) GetVaultServerBaseURI() []string {
	if o == nil || isNil(o.VaultServerBaseURI) {
		var ret []string
		return ret
	}
	return o.VaultServerBaseURI
}

// GetVaultServerBaseURIOk returns a tuple with the VaultServerBaseURI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultCipherStreamProviderShared) GetVaultServerBaseURIOk() ([]string, bool) {
	if o == nil || isNil(o.VaultServerBaseURI) {
    return nil, false
	}
	return o.VaultServerBaseURI, true
}

// HasVaultServerBaseURI returns a boolean if a field has been set.
func (o *VaultCipherStreamProviderShared) HasVaultServerBaseURI() bool {
	if o != nil && !isNil(o.VaultServerBaseURI) {
		return true
	}

	return false
}

// SetVaultServerBaseURI gets a reference to the given []string and assigns it to the VaultServerBaseURI field.
func (o *VaultCipherStreamProviderShared) SetVaultServerBaseURI(v []string) {
	o.VaultServerBaseURI = v
}

// GetVaultAuthenticationMethod returns the VaultAuthenticationMethod field value if set, zero value otherwise.
func (o *VaultCipherStreamProviderShared) GetVaultAuthenticationMethod() string {
	if o == nil || isNil(o.VaultAuthenticationMethod) {
		var ret string
		return ret
	}
	return *o.VaultAuthenticationMethod
}

// GetVaultAuthenticationMethodOk returns a tuple with the VaultAuthenticationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultCipherStreamProviderShared) GetVaultAuthenticationMethodOk() (*string, bool) {
	if o == nil || isNil(o.VaultAuthenticationMethod) {
    return nil, false
	}
	return o.VaultAuthenticationMethod, true
}

// HasVaultAuthenticationMethod returns a boolean if a field has been set.
func (o *VaultCipherStreamProviderShared) HasVaultAuthenticationMethod() bool {
	if o != nil && !isNil(o.VaultAuthenticationMethod) {
		return true
	}

	return false
}

// SetVaultAuthenticationMethod gets a reference to the given string and assigns it to the VaultAuthenticationMethod field.
func (o *VaultCipherStreamProviderShared) SetVaultAuthenticationMethod(v string) {
	o.VaultAuthenticationMethod = &v
}

// GetVaultSecretPath returns the VaultSecretPath field value
func (o *VaultCipherStreamProviderShared) GetVaultSecretPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultSecretPath
}

// GetVaultSecretPathOk returns a tuple with the VaultSecretPath field value
// and a boolean to check if the value has been set.
func (o *VaultCipherStreamProviderShared) GetVaultSecretPathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.VaultSecretPath, true
}

// SetVaultSecretPath sets field value
func (o *VaultCipherStreamProviderShared) SetVaultSecretPath(v string) {
	o.VaultSecretPath = v
}

// GetVaultSecretFieldName returns the VaultSecretFieldName field value
func (o *VaultCipherStreamProviderShared) GetVaultSecretFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultSecretFieldName
}

// GetVaultSecretFieldNameOk returns a tuple with the VaultSecretFieldName field value
// and a boolean to check if the value has been set.
func (o *VaultCipherStreamProviderShared) GetVaultSecretFieldNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.VaultSecretFieldName, true
}

// SetVaultSecretFieldName sets field value
func (o *VaultCipherStreamProviderShared) SetVaultSecretFieldName(v string) {
	o.VaultSecretFieldName = v
}

// GetVaultEncryptionMetadataFile returns the VaultEncryptionMetadataFile field value
func (o *VaultCipherStreamProviderShared) GetVaultEncryptionMetadataFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultEncryptionMetadataFile
}

// GetVaultEncryptionMetadataFileOk returns a tuple with the VaultEncryptionMetadataFile field value
// and a boolean to check if the value has been set.
func (o *VaultCipherStreamProviderShared) GetVaultEncryptionMetadataFileOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.VaultEncryptionMetadataFile, true
}

// SetVaultEncryptionMetadataFile sets field value
func (o *VaultCipherStreamProviderShared) SetVaultEncryptionMetadataFile(v string) {
	o.VaultEncryptionMetadataFile = v
}

// GetTrustStoreFile returns the TrustStoreFile field value if set, zero value otherwise.
func (o *VaultCipherStreamProviderShared) GetTrustStoreFile() string {
	if o == nil || isNil(o.TrustStoreFile) {
		var ret string
		return ret
	}
	return *o.TrustStoreFile
}

// GetTrustStoreFileOk returns a tuple with the TrustStoreFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultCipherStreamProviderShared) GetTrustStoreFileOk() (*string, bool) {
	if o == nil || isNil(o.TrustStoreFile) {
    return nil, false
	}
	return o.TrustStoreFile, true
}

// HasTrustStoreFile returns a boolean if a field has been set.
func (o *VaultCipherStreamProviderShared) HasTrustStoreFile() bool {
	if o != nil && !isNil(o.TrustStoreFile) {
		return true
	}

	return false
}

// SetTrustStoreFile gets a reference to the given string and assigns it to the TrustStoreFile field.
func (o *VaultCipherStreamProviderShared) SetTrustStoreFile(v string) {
	o.TrustStoreFile = &v
}

// GetTrustStorePin returns the TrustStorePin field value if set, zero value otherwise.
func (o *VaultCipherStreamProviderShared) GetTrustStorePin() string {
	if o == nil || isNil(o.TrustStorePin) {
		var ret string
		return ret
	}
	return *o.TrustStorePin
}

// GetTrustStorePinOk returns a tuple with the TrustStorePin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultCipherStreamProviderShared) GetTrustStorePinOk() (*string, bool) {
	if o == nil || isNil(o.TrustStorePin) {
    return nil, false
	}
	return o.TrustStorePin, true
}

// HasTrustStorePin returns a boolean if a field has been set.
func (o *VaultCipherStreamProviderShared) HasTrustStorePin() bool {
	if o != nil && !isNil(o.TrustStorePin) {
		return true
	}

	return false
}

// SetTrustStorePin gets a reference to the given string and assigns it to the TrustStorePin field.
func (o *VaultCipherStreamProviderShared) SetTrustStorePin(v string) {
	o.TrustStorePin = &v
}

// GetTrustStoreType returns the TrustStoreType field value if set, zero value otherwise.
func (o *VaultCipherStreamProviderShared) GetTrustStoreType() string {
	if o == nil || isNil(o.TrustStoreType) {
		var ret string
		return ret
	}
	return *o.TrustStoreType
}

// GetTrustStoreTypeOk returns a tuple with the TrustStoreType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultCipherStreamProviderShared) GetTrustStoreTypeOk() (*string, bool) {
	if o == nil || isNil(o.TrustStoreType) {
    return nil, false
	}
	return o.TrustStoreType, true
}

// HasTrustStoreType returns a boolean if a field has been set.
func (o *VaultCipherStreamProviderShared) HasTrustStoreType() bool {
	if o != nil && !isNil(o.TrustStoreType) {
		return true
	}

	return false
}

// SetTrustStoreType gets a reference to the given string and assigns it to the TrustStoreType field.
func (o *VaultCipherStreamProviderShared) SetTrustStoreType(v string) {
	o.TrustStoreType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VaultCipherStreamProviderShared) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultCipherStreamProviderShared) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VaultCipherStreamProviderShared) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VaultCipherStreamProviderShared) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *VaultCipherStreamProviderShared) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *VaultCipherStreamProviderShared) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *VaultCipherStreamProviderShared) SetEnabled(v bool) {
	o.Enabled = v
}

func (o VaultCipherStreamProviderShared) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.VaultExternalServer) {
		toSerialize["vaultExternalServer"] = o.VaultExternalServer
	}
	if !isNil(o.VaultServerBaseURI) {
		toSerialize["vaultServerBaseURI"] = o.VaultServerBaseURI
	}
	if !isNil(o.VaultAuthenticationMethod) {
		toSerialize["vaultAuthenticationMethod"] = o.VaultAuthenticationMethod
	}
	if true {
		toSerialize["vaultSecretPath"] = o.VaultSecretPath
	}
	if true {
		toSerialize["vaultSecretFieldName"] = o.VaultSecretFieldName
	}
	if true {
		toSerialize["vaultEncryptionMetadataFile"] = o.VaultEncryptionMetadataFile
	}
	if !isNil(o.TrustStoreFile) {
		toSerialize["trustStoreFile"] = o.TrustStoreFile
	}
	if !isNil(o.TrustStorePin) {
		toSerialize["trustStorePin"] = o.TrustStorePin
	}
	if !isNil(o.TrustStoreType) {
		toSerialize["trustStoreType"] = o.TrustStoreType
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableVaultCipherStreamProviderShared struct {
	value *VaultCipherStreamProviderShared
	isSet bool
}

func (v NullableVaultCipherStreamProviderShared) Get() *VaultCipherStreamProviderShared {
	return v.value
}

func (v *NullableVaultCipherStreamProviderShared) Set(val *VaultCipherStreamProviderShared) {
	v.value = val
	v.isSet = true
}

func (v NullableVaultCipherStreamProviderShared) IsSet() bool {
	return v.isSet
}

func (v *NullableVaultCipherStreamProviderShared) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVaultCipherStreamProviderShared(val *VaultCipherStreamProviderShared) *NullableVaultCipherStreamProviderShared {
	return &NullableVaultCipherStreamProviderShared{value: val, isSet: true}
}

func (v NullableVaultCipherStreamProviderShared) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVaultCipherStreamProviderShared) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


