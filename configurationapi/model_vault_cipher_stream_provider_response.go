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

// checks if the VaultCipherStreamProviderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VaultCipherStreamProviderResponse{}

// VaultCipherStreamProviderResponse struct for VaultCipherStreamProviderResponse
type VaultCipherStreamProviderResponse struct {
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
	VaultEncryptionMetadataFile string `json:"vaultEncryptionMetadataFile"`
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
	Enabled                                       bool                                               `json:"enabled"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Cipher Stream Provider
	Id string `json:"id"`
}

// NewVaultCipherStreamProviderResponse instantiates a new VaultCipherStreamProviderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVaultCipherStreamProviderResponse(schemas []EnumvaultCipherStreamProviderSchemaUrn, vaultSecretPath string, vaultSecretFieldName string, vaultEncryptionMetadataFile string, enabled bool, id string) *VaultCipherStreamProviderResponse {
	this := VaultCipherStreamProviderResponse{}
	this.Schemas = schemas
	this.VaultSecretPath = vaultSecretPath
	this.VaultSecretFieldName = vaultSecretFieldName
	this.VaultEncryptionMetadataFile = vaultEncryptionMetadataFile
	this.Enabled = enabled
	this.Id = id
	return &this
}

// NewVaultCipherStreamProviderResponseWithDefaults instantiates a new VaultCipherStreamProviderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVaultCipherStreamProviderResponseWithDefaults() *VaultCipherStreamProviderResponse {
	this := VaultCipherStreamProviderResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *VaultCipherStreamProviderResponse) GetSchemas() []EnumvaultCipherStreamProviderSchemaUrn {
	if o == nil {
		var ret []EnumvaultCipherStreamProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *VaultCipherStreamProviderResponse) GetSchemasOk() ([]EnumvaultCipherStreamProviderSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *VaultCipherStreamProviderResponse) SetSchemas(v []EnumvaultCipherStreamProviderSchemaUrn) {
	o.Schemas = v
}

// GetVaultExternalServer returns the VaultExternalServer field value if set, zero value otherwise.
func (o *VaultCipherStreamProviderResponse) GetVaultExternalServer() string {
	if o == nil || IsNil(o.VaultExternalServer) {
		var ret string
		return ret
	}
	return *o.VaultExternalServer
}

// GetVaultExternalServerOk returns a tuple with the VaultExternalServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultCipherStreamProviderResponse) GetVaultExternalServerOk() (*string, bool) {
	if o == nil || IsNil(o.VaultExternalServer) {
		return nil, false
	}
	return o.VaultExternalServer, true
}

// HasVaultExternalServer returns a boolean if a field has been set.
func (o *VaultCipherStreamProviderResponse) HasVaultExternalServer() bool {
	if o != nil && !IsNil(o.VaultExternalServer) {
		return true
	}

	return false
}

// SetVaultExternalServer gets a reference to the given string and assigns it to the VaultExternalServer field.
func (o *VaultCipherStreamProviderResponse) SetVaultExternalServer(v string) {
	o.VaultExternalServer = &v
}

// GetVaultServerBaseURI returns the VaultServerBaseURI field value if set, zero value otherwise.
func (o *VaultCipherStreamProviderResponse) GetVaultServerBaseURI() []string {
	if o == nil || IsNil(o.VaultServerBaseURI) {
		var ret []string
		return ret
	}
	return o.VaultServerBaseURI
}

// GetVaultServerBaseURIOk returns a tuple with the VaultServerBaseURI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultCipherStreamProviderResponse) GetVaultServerBaseURIOk() ([]string, bool) {
	if o == nil || IsNil(o.VaultServerBaseURI) {
		return nil, false
	}
	return o.VaultServerBaseURI, true
}

// HasVaultServerBaseURI returns a boolean if a field has been set.
func (o *VaultCipherStreamProviderResponse) HasVaultServerBaseURI() bool {
	if o != nil && !IsNil(o.VaultServerBaseURI) {
		return true
	}

	return false
}

// SetVaultServerBaseURI gets a reference to the given []string and assigns it to the VaultServerBaseURI field.
func (o *VaultCipherStreamProviderResponse) SetVaultServerBaseURI(v []string) {
	o.VaultServerBaseURI = v
}

// GetVaultAuthenticationMethod returns the VaultAuthenticationMethod field value if set, zero value otherwise.
func (o *VaultCipherStreamProviderResponse) GetVaultAuthenticationMethod() string {
	if o == nil || IsNil(o.VaultAuthenticationMethod) {
		var ret string
		return ret
	}
	return *o.VaultAuthenticationMethod
}

// GetVaultAuthenticationMethodOk returns a tuple with the VaultAuthenticationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultCipherStreamProviderResponse) GetVaultAuthenticationMethodOk() (*string, bool) {
	if o == nil || IsNil(o.VaultAuthenticationMethod) {
		return nil, false
	}
	return o.VaultAuthenticationMethod, true
}

// HasVaultAuthenticationMethod returns a boolean if a field has been set.
func (o *VaultCipherStreamProviderResponse) HasVaultAuthenticationMethod() bool {
	if o != nil && !IsNil(o.VaultAuthenticationMethod) {
		return true
	}

	return false
}

// SetVaultAuthenticationMethod gets a reference to the given string and assigns it to the VaultAuthenticationMethod field.
func (o *VaultCipherStreamProviderResponse) SetVaultAuthenticationMethod(v string) {
	o.VaultAuthenticationMethod = &v
}

// GetVaultSecretPath returns the VaultSecretPath field value
func (o *VaultCipherStreamProviderResponse) GetVaultSecretPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultSecretPath
}

// GetVaultSecretPathOk returns a tuple with the VaultSecretPath field value
// and a boolean to check if the value has been set.
func (o *VaultCipherStreamProviderResponse) GetVaultSecretPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VaultSecretPath, true
}

// SetVaultSecretPath sets field value
func (o *VaultCipherStreamProviderResponse) SetVaultSecretPath(v string) {
	o.VaultSecretPath = v
}

// GetVaultSecretFieldName returns the VaultSecretFieldName field value
func (o *VaultCipherStreamProviderResponse) GetVaultSecretFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultSecretFieldName
}

// GetVaultSecretFieldNameOk returns a tuple with the VaultSecretFieldName field value
// and a boolean to check if the value has been set.
func (o *VaultCipherStreamProviderResponse) GetVaultSecretFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VaultSecretFieldName, true
}

// SetVaultSecretFieldName sets field value
func (o *VaultCipherStreamProviderResponse) SetVaultSecretFieldName(v string) {
	o.VaultSecretFieldName = v
}

// GetVaultEncryptionMetadataFile returns the VaultEncryptionMetadataFile field value
func (o *VaultCipherStreamProviderResponse) GetVaultEncryptionMetadataFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultEncryptionMetadataFile
}

// GetVaultEncryptionMetadataFileOk returns a tuple with the VaultEncryptionMetadataFile field value
// and a boolean to check if the value has been set.
func (o *VaultCipherStreamProviderResponse) GetVaultEncryptionMetadataFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VaultEncryptionMetadataFile, true
}

// SetVaultEncryptionMetadataFile sets field value
func (o *VaultCipherStreamProviderResponse) SetVaultEncryptionMetadataFile(v string) {
	o.VaultEncryptionMetadataFile = v
}

// GetTrustStoreFile returns the TrustStoreFile field value if set, zero value otherwise.
func (o *VaultCipherStreamProviderResponse) GetTrustStoreFile() string {
	if o == nil || IsNil(o.TrustStoreFile) {
		var ret string
		return ret
	}
	return *o.TrustStoreFile
}

// GetTrustStoreFileOk returns a tuple with the TrustStoreFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultCipherStreamProviderResponse) GetTrustStoreFileOk() (*string, bool) {
	if o == nil || IsNil(o.TrustStoreFile) {
		return nil, false
	}
	return o.TrustStoreFile, true
}

// HasTrustStoreFile returns a boolean if a field has been set.
func (o *VaultCipherStreamProviderResponse) HasTrustStoreFile() bool {
	if o != nil && !IsNil(o.TrustStoreFile) {
		return true
	}

	return false
}

// SetTrustStoreFile gets a reference to the given string and assigns it to the TrustStoreFile field.
func (o *VaultCipherStreamProviderResponse) SetTrustStoreFile(v string) {
	o.TrustStoreFile = &v
}

// GetTrustStorePin returns the TrustStorePin field value if set, zero value otherwise.
func (o *VaultCipherStreamProviderResponse) GetTrustStorePin() string {
	if o == nil || IsNil(o.TrustStorePin) {
		var ret string
		return ret
	}
	return *o.TrustStorePin
}

// GetTrustStorePinOk returns a tuple with the TrustStorePin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultCipherStreamProviderResponse) GetTrustStorePinOk() (*string, bool) {
	if o == nil || IsNil(o.TrustStorePin) {
		return nil, false
	}
	return o.TrustStorePin, true
}

// HasTrustStorePin returns a boolean if a field has been set.
func (o *VaultCipherStreamProviderResponse) HasTrustStorePin() bool {
	if o != nil && !IsNil(o.TrustStorePin) {
		return true
	}

	return false
}

// SetTrustStorePin gets a reference to the given string and assigns it to the TrustStorePin field.
func (o *VaultCipherStreamProviderResponse) SetTrustStorePin(v string) {
	o.TrustStorePin = &v
}

// GetTrustStoreType returns the TrustStoreType field value if set, zero value otherwise.
func (o *VaultCipherStreamProviderResponse) GetTrustStoreType() string {
	if o == nil || IsNil(o.TrustStoreType) {
		var ret string
		return ret
	}
	return *o.TrustStoreType
}

// GetTrustStoreTypeOk returns a tuple with the TrustStoreType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultCipherStreamProviderResponse) GetTrustStoreTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TrustStoreType) {
		return nil, false
	}
	return o.TrustStoreType, true
}

// HasTrustStoreType returns a boolean if a field has been set.
func (o *VaultCipherStreamProviderResponse) HasTrustStoreType() bool {
	if o != nil && !IsNil(o.TrustStoreType) {
		return true
	}

	return false
}

// SetTrustStoreType gets a reference to the given string and assigns it to the TrustStoreType field.
func (o *VaultCipherStreamProviderResponse) SetTrustStoreType(v string) {
	o.TrustStoreType = &v
}

// GetIterationCount returns the IterationCount field value if set, zero value otherwise.
func (o *VaultCipherStreamProviderResponse) GetIterationCount() int64 {
	if o == nil || IsNil(o.IterationCount) {
		var ret int64
		return ret
	}
	return *o.IterationCount
}

// GetIterationCountOk returns a tuple with the IterationCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultCipherStreamProviderResponse) GetIterationCountOk() (*int64, bool) {
	if o == nil || IsNil(o.IterationCount) {
		return nil, false
	}
	return o.IterationCount, true
}

// HasIterationCount returns a boolean if a field has been set.
func (o *VaultCipherStreamProviderResponse) HasIterationCount() bool {
	if o != nil && !IsNil(o.IterationCount) {
		return true
	}

	return false
}

// SetIterationCount gets a reference to the given int64 and assigns it to the IterationCount field.
func (o *VaultCipherStreamProviderResponse) SetIterationCount(v int64) {
	o.IterationCount = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VaultCipherStreamProviderResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultCipherStreamProviderResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VaultCipherStreamProviderResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VaultCipherStreamProviderResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *VaultCipherStreamProviderResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *VaultCipherStreamProviderResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *VaultCipherStreamProviderResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *VaultCipherStreamProviderResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultCipherStreamProviderResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *VaultCipherStreamProviderResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *VaultCipherStreamProviderResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *VaultCipherStreamProviderResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultCipherStreamProviderResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *VaultCipherStreamProviderResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *VaultCipherStreamProviderResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *VaultCipherStreamProviderResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VaultCipherStreamProviderResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VaultCipherStreamProviderResponse) SetId(v string) {
	o.Id = v
}

func (o VaultCipherStreamProviderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VaultCipherStreamProviderResponse) ToMap() (map[string]interface{}, error) {
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
	toSerialize["vaultEncryptionMetadataFile"] = o.VaultEncryptionMetadataFile
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
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableVaultCipherStreamProviderResponse struct {
	value *VaultCipherStreamProviderResponse
	isSet bool
}

func (v NullableVaultCipherStreamProviderResponse) Get() *VaultCipherStreamProviderResponse {
	return v.value
}

func (v *NullableVaultCipherStreamProviderResponse) Set(val *VaultCipherStreamProviderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVaultCipherStreamProviderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVaultCipherStreamProviderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVaultCipherStreamProviderResponse(val *VaultCipherStreamProviderResponse) *NullableVaultCipherStreamProviderResponse {
	return &NullableVaultCipherStreamProviderResponse{value: val, isSet: true}
}

func (v NullableVaultCipherStreamProviderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVaultCipherStreamProviderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
