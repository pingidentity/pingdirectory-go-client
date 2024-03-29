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

// checks if the AddVaultCipherStreamProviderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddVaultCipherStreamProviderRequest{}

// AddVaultCipherStreamProviderRequest struct for AddVaultCipherStreamProviderRequest
type AddVaultCipherStreamProviderRequest struct {
	Schemas []EnumvaultCipherStreamProviderSchemaUrn `json:"schemas"`
	// An external server definition with information needed to connect and authenticate to the Vault server.
	VaultExternalServer *string `json:"vaultExternalServer,omitempty"`
	// The base URL needed to access the Vault server. The base URL should consist of the protocol (\"http\" or \"https\"), the server address (resolvable name or IP address), and the port number. For example, \"https://vault.example.com:8200/\".
	VaultServerBaseURI []string `json:"vaultServerBaseURI,omitempty"`
	// The mechanism used to authenticate to the Vault server.
	VaultAuthenticationMethod *string `json:"vaultAuthenticationMethod,omitempty"`
	// The path to the desired secret in the Vault service. This will be appended to the value of the base-url property for the associated Vault external server.
	VaultSecretPath string `json:"vaultSecretPath"`
	// The name of the field in the Vault secret record that contains the passphrase to use to generate the encryption key.
	VaultSecretFieldName string `json:"vaultSecretFieldName"`
	// The path to a file that will hold metadata about the encryption performed by this Vault Cipher Stream Provider.
	VaultEncryptionMetadataFile *string `json:"vaultEncryptionMetadataFile,omitempty"`
	// The path to a file containing the information needed to trust the certificate presented by the Vault servers.
	TrustStoreFile *string `json:"trustStoreFile,omitempty"`
	// The passphrase needed to access the contents of the trust store. This is only required if a trust store file is required, and if that trust store requires a PIN to access its contents.
	TrustStorePin *string `json:"trustStorePin,omitempty"`
	// The store type for the specified trust store file. The value should likely be one of \"JKS\" or \"PKCS12\".
	TrustStoreType *string `json:"trustStoreType,omitempty"`
	// The PBKDF2 iteration count that will be used when deriving the encryption key used to protect the encryption settings database.
	IterationCount *int64 `json:"iterationCount,omitempty"`
	// A description for this Cipher Stream Provider
	Description *string `json:"description,omitempty"`
	// Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server.
	Enabled bool `json:"enabled"`
	// Name of the new Cipher Stream Provider
	ProviderName string `json:"providerName"`
}

// NewAddVaultCipherStreamProviderRequest instantiates a new AddVaultCipherStreamProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddVaultCipherStreamProviderRequest(schemas []EnumvaultCipherStreamProviderSchemaUrn, vaultSecretPath string, vaultSecretFieldName string, enabled bool, providerName string) *AddVaultCipherStreamProviderRequest {
	this := AddVaultCipherStreamProviderRequest{}
	this.Schemas = schemas
	this.VaultSecretPath = vaultSecretPath
	this.VaultSecretFieldName = vaultSecretFieldName
	this.Enabled = enabled
	this.ProviderName = providerName
	return &this
}

// NewAddVaultCipherStreamProviderRequestWithDefaults instantiates a new AddVaultCipherStreamProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddVaultCipherStreamProviderRequestWithDefaults() *AddVaultCipherStreamProviderRequest {
	this := AddVaultCipherStreamProviderRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddVaultCipherStreamProviderRequest) GetSchemas() []EnumvaultCipherStreamProviderSchemaUrn {
	if o == nil {
		var ret []EnumvaultCipherStreamProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddVaultCipherStreamProviderRequest) GetSchemasOk() ([]EnumvaultCipherStreamProviderSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddVaultCipherStreamProviderRequest) SetSchemas(v []EnumvaultCipherStreamProviderSchemaUrn) {
	o.Schemas = v
}

// GetVaultExternalServer returns the VaultExternalServer field value if set, zero value otherwise.
func (o *AddVaultCipherStreamProviderRequest) GetVaultExternalServer() string {
	if o == nil || IsNil(o.VaultExternalServer) {
		var ret string
		return ret
	}
	return *o.VaultExternalServer
}

// GetVaultExternalServerOk returns a tuple with the VaultExternalServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVaultCipherStreamProviderRequest) GetVaultExternalServerOk() (*string, bool) {
	if o == nil || IsNil(o.VaultExternalServer) {
		return nil, false
	}
	return o.VaultExternalServer, true
}

// HasVaultExternalServer returns a boolean if a field has been set.
func (o *AddVaultCipherStreamProviderRequest) HasVaultExternalServer() bool {
	if o != nil && !IsNil(o.VaultExternalServer) {
		return true
	}

	return false
}

// SetVaultExternalServer gets a reference to the given string and assigns it to the VaultExternalServer field.
func (o *AddVaultCipherStreamProviderRequest) SetVaultExternalServer(v string) {
	o.VaultExternalServer = &v
}

// GetVaultServerBaseURI returns the VaultServerBaseURI field value if set, zero value otherwise.
func (o *AddVaultCipherStreamProviderRequest) GetVaultServerBaseURI() []string {
	if o == nil || IsNil(o.VaultServerBaseURI) {
		var ret []string
		return ret
	}
	return o.VaultServerBaseURI
}

// GetVaultServerBaseURIOk returns a tuple with the VaultServerBaseURI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVaultCipherStreamProviderRequest) GetVaultServerBaseURIOk() ([]string, bool) {
	if o == nil || IsNil(o.VaultServerBaseURI) {
		return nil, false
	}
	return o.VaultServerBaseURI, true
}

// HasVaultServerBaseURI returns a boolean if a field has been set.
func (o *AddVaultCipherStreamProviderRequest) HasVaultServerBaseURI() bool {
	if o != nil && !IsNil(o.VaultServerBaseURI) {
		return true
	}

	return false
}

// SetVaultServerBaseURI gets a reference to the given []string and assigns it to the VaultServerBaseURI field.
func (o *AddVaultCipherStreamProviderRequest) SetVaultServerBaseURI(v []string) {
	o.VaultServerBaseURI = v
}

// GetVaultAuthenticationMethod returns the VaultAuthenticationMethod field value if set, zero value otherwise.
func (o *AddVaultCipherStreamProviderRequest) GetVaultAuthenticationMethod() string {
	if o == nil || IsNil(o.VaultAuthenticationMethod) {
		var ret string
		return ret
	}
	return *o.VaultAuthenticationMethod
}

// GetVaultAuthenticationMethodOk returns a tuple with the VaultAuthenticationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVaultCipherStreamProviderRequest) GetVaultAuthenticationMethodOk() (*string, bool) {
	if o == nil || IsNil(o.VaultAuthenticationMethod) {
		return nil, false
	}
	return o.VaultAuthenticationMethod, true
}

// HasVaultAuthenticationMethod returns a boolean if a field has been set.
func (o *AddVaultCipherStreamProviderRequest) HasVaultAuthenticationMethod() bool {
	if o != nil && !IsNil(o.VaultAuthenticationMethod) {
		return true
	}

	return false
}

// SetVaultAuthenticationMethod gets a reference to the given string and assigns it to the VaultAuthenticationMethod field.
func (o *AddVaultCipherStreamProviderRequest) SetVaultAuthenticationMethod(v string) {
	o.VaultAuthenticationMethod = &v
}

// GetVaultSecretPath returns the VaultSecretPath field value
func (o *AddVaultCipherStreamProviderRequest) GetVaultSecretPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultSecretPath
}

// GetVaultSecretPathOk returns a tuple with the VaultSecretPath field value
// and a boolean to check if the value has been set.
func (o *AddVaultCipherStreamProviderRequest) GetVaultSecretPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VaultSecretPath, true
}

// SetVaultSecretPath sets field value
func (o *AddVaultCipherStreamProviderRequest) SetVaultSecretPath(v string) {
	o.VaultSecretPath = v
}

// GetVaultSecretFieldName returns the VaultSecretFieldName field value
func (o *AddVaultCipherStreamProviderRequest) GetVaultSecretFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultSecretFieldName
}

// GetVaultSecretFieldNameOk returns a tuple with the VaultSecretFieldName field value
// and a boolean to check if the value has been set.
func (o *AddVaultCipherStreamProviderRequest) GetVaultSecretFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VaultSecretFieldName, true
}

// SetVaultSecretFieldName sets field value
func (o *AddVaultCipherStreamProviderRequest) SetVaultSecretFieldName(v string) {
	o.VaultSecretFieldName = v
}

// GetVaultEncryptionMetadataFile returns the VaultEncryptionMetadataFile field value if set, zero value otherwise.
func (o *AddVaultCipherStreamProviderRequest) GetVaultEncryptionMetadataFile() string {
	if o == nil || IsNil(o.VaultEncryptionMetadataFile) {
		var ret string
		return ret
	}
	return *o.VaultEncryptionMetadataFile
}

// GetVaultEncryptionMetadataFileOk returns a tuple with the VaultEncryptionMetadataFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVaultCipherStreamProviderRequest) GetVaultEncryptionMetadataFileOk() (*string, bool) {
	if o == nil || IsNil(o.VaultEncryptionMetadataFile) {
		return nil, false
	}
	return o.VaultEncryptionMetadataFile, true
}

// HasVaultEncryptionMetadataFile returns a boolean if a field has been set.
func (o *AddVaultCipherStreamProviderRequest) HasVaultEncryptionMetadataFile() bool {
	if o != nil && !IsNil(o.VaultEncryptionMetadataFile) {
		return true
	}

	return false
}

// SetVaultEncryptionMetadataFile gets a reference to the given string and assigns it to the VaultEncryptionMetadataFile field.
func (o *AddVaultCipherStreamProviderRequest) SetVaultEncryptionMetadataFile(v string) {
	o.VaultEncryptionMetadataFile = &v
}

// GetTrustStoreFile returns the TrustStoreFile field value if set, zero value otherwise.
func (o *AddVaultCipherStreamProviderRequest) GetTrustStoreFile() string {
	if o == nil || IsNil(o.TrustStoreFile) {
		var ret string
		return ret
	}
	return *o.TrustStoreFile
}

// GetTrustStoreFileOk returns a tuple with the TrustStoreFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVaultCipherStreamProviderRequest) GetTrustStoreFileOk() (*string, bool) {
	if o == nil || IsNil(o.TrustStoreFile) {
		return nil, false
	}
	return o.TrustStoreFile, true
}

// HasTrustStoreFile returns a boolean if a field has been set.
func (o *AddVaultCipherStreamProviderRequest) HasTrustStoreFile() bool {
	if o != nil && !IsNil(o.TrustStoreFile) {
		return true
	}

	return false
}

// SetTrustStoreFile gets a reference to the given string and assigns it to the TrustStoreFile field.
func (o *AddVaultCipherStreamProviderRequest) SetTrustStoreFile(v string) {
	o.TrustStoreFile = &v
}

// GetTrustStorePin returns the TrustStorePin field value if set, zero value otherwise.
func (o *AddVaultCipherStreamProviderRequest) GetTrustStorePin() string {
	if o == nil || IsNil(o.TrustStorePin) {
		var ret string
		return ret
	}
	return *o.TrustStorePin
}

// GetTrustStorePinOk returns a tuple with the TrustStorePin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVaultCipherStreamProviderRequest) GetTrustStorePinOk() (*string, bool) {
	if o == nil || IsNil(o.TrustStorePin) {
		return nil, false
	}
	return o.TrustStorePin, true
}

// HasTrustStorePin returns a boolean if a field has been set.
func (o *AddVaultCipherStreamProviderRequest) HasTrustStorePin() bool {
	if o != nil && !IsNil(o.TrustStorePin) {
		return true
	}

	return false
}

// SetTrustStorePin gets a reference to the given string and assigns it to the TrustStorePin field.
func (o *AddVaultCipherStreamProviderRequest) SetTrustStorePin(v string) {
	o.TrustStorePin = &v
}

// GetTrustStoreType returns the TrustStoreType field value if set, zero value otherwise.
func (o *AddVaultCipherStreamProviderRequest) GetTrustStoreType() string {
	if o == nil || IsNil(o.TrustStoreType) {
		var ret string
		return ret
	}
	return *o.TrustStoreType
}

// GetTrustStoreTypeOk returns a tuple with the TrustStoreType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVaultCipherStreamProviderRequest) GetTrustStoreTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TrustStoreType) {
		return nil, false
	}
	return o.TrustStoreType, true
}

// HasTrustStoreType returns a boolean if a field has been set.
func (o *AddVaultCipherStreamProviderRequest) HasTrustStoreType() bool {
	if o != nil && !IsNil(o.TrustStoreType) {
		return true
	}

	return false
}

// SetTrustStoreType gets a reference to the given string and assigns it to the TrustStoreType field.
func (o *AddVaultCipherStreamProviderRequest) SetTrustStoreType(v string) {
	o.TrustStoreType = &v
}

// GetIterationCount returns the IterationCount field value if set, zero value otherwise.
func (o *AddVaultCipherStreamProviderRequest) GetIterationCount() int64 {
	if o == nil || IsNil(o.IterationCount) {
		var ret int64
		return ret
	}
	return *o.IterationCount
}

// GetIterationCountOk returns a tuple with the IterationCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVaultCipherStreamProviderRequest) GetIterationCountOk() (*int64, bool) {
	if o == nil || IsNil(o.IterationCount) {
		return nil, false
	}
	return o.IterationCount, true
}

// HasIterationCount returns a boolean if a field has been set.
func (o *AddVaultCipherStreamProviderRequest) HasIterationCount() bool {
	if o != nil && !IsNil(o.IterationCount) {
		return true
	}

	return false
}

// SetIterationCount gets a reference to the given int64 and assigns it to the IterationCount field.
func (o *AddVaultCipherStreamProviderRequest) SetIterationCount(v int64) {
	o.IterationCount = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddVaultCipherStreamProviderRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVaultCipherStreamProviderRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddVaultCipherStreamProviderRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddVaultCipherStreamProviderRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddVaultCipherStreamProviderRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddVaultCipherStreamProviderRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddVaultCipherStreamProviderRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetProviderName returns the ProviderName field value
func (o *AddVaultCipherStreamProviderRequest) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *AddVaultCipherStreamProviderRequest) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *AddVaultCipherStreamProviderRequest) SetProviderName(v string) {
	o.ProviderName = v
}

func (o AddVaultCipherStreamProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddVaultCipherStreamProviderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.VaultExternalServer) {
		toSerialize["vaultExternalServer"] = o.VaultExternalServer
	}
	if !IsNil(o.VaultServerBaseURI) {
		toSerialize["vaultServerBaseURI"] = o.VaultServerBaseURI
	}
	if !IsNil(o.VaultAuthenticationMethod) {
		toSerialize["vaultAuthenticationMethod"] = o.VaultAuthenticationMethod
	}
	toSerialize["vaultSecretPath"] = o.VaultSecretPath
	toSerialize["vaultSecretFieldName"] = o.VaultSecretFieldName
	if !IsNil(o.VaultEncryptionMetadataFile) {
		toSerialize["vaultEncryptionMetadataFile"] = o.VaultEncryptionMetadataFile
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
	if !IsNil(o.IterationCount) {
		toSerialize["iterationCount"] = o.IterationCount
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["providerName"] = o.ProviderName
	return toSerialize, nil
}

type NullableAddVaultCipherStreamProviderRequest struct {
	value *AddVaultCipherStreamProviderRequest
	isSet bool
}

func (v NullableAddVaultCipherStreamProviderRequest) Get() *AddVaultCipherStreamProviderRequest {
	return v.value
}

func (v *NullableAddVaultCipherStreamProviderRequest) Set(val *AddVaultCipherStreamProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddVaultCipherStreamProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddVaultCipherStreamProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddVaultCipherStreamProviderRequest(val *AddVaultCipherStreamProviderRequest) *NullableAddVaultCipherStreamProviderRequest {
	return &NullableAddVaultCipherStreamProviderRequest{value: val, isSet: true}
}

func (v NullableAddVaultCipherStreamProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddVaultCipherStreamProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
