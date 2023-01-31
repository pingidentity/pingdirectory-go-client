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

// AddPkcs11CipherStreamProviderRequest struct for AddPkcs11CipherStreamProviderRequest
type AddPkcs11CipherStreamProviderRequest struct {
	// Name of the new Cipher Stream Provider
	ProviderName string                                    `json:"providerName"`
	Schemas      []Enumpkcs11CipherStreamProviderSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java security provider class that implements support for interacting with PKCS #11 tokens.
	Pkcs11ProviderClass *string `json:"pkcs11ProviderClass,omitempty"`
	// The path to the file to use to configure the security provider that implements support for interacting with PKCS #11 tokens.
	Pkcs11ProviderConfigurationFile *string `json:"pkcs11ProviderConfigurationFile,omitempty"`
	// The clear-text user PIN needed to interact with the PKCS #11 token.
	KeyStorePin *string `json:"keyStorePin,omitempty"`
	// The path to a file containing the user PIN needed to interact with the PKCS #11 token. The file must exist and must contain exactly one line with a clear-text representation of the PIN.
	KeyStorePinFile *string `json:"keyStorePinFile,omitempty"`
	// The name of an environment variable whose value is the user PIN needed to interact with the PKCS #11 token. The environment variable must be defined and must contain a clear-text representation of the PIN.
	KeyStorePinEnvironmentVariable *string `json:"keyStorePinEnvironmentVariable,omitempty"`
	// The key store type to use when obtaining an instance of a key store for interacting with a PKCS #11 token.
	Pkcs11KeyStoreType *string `json:"pkcs11KeyStoreType,omitempty"`
	// The alias for the certificate in the PKCS #11 token that will be used to wrap the encryption key. The target certificate must exist in the PKCS #11 token, and it must have an RSA key pair because the JVM does not currently provide adequate key wrapping support for elliptic curve key pairs.  If you have also configured the server to use a PKCS #11 token for accessing listener certificates, we strongly recommend that you use a different certificate to protect the contents of the encryption settings database than you use for negotiating TLS sessions with clients. It is imperative that the certificate used by this PKCS11 Cipher Stream Provider remain constant for the life of the provider because if the certificate were to be replaced, then the contents of the encryption settings database could become inaccessible. Unlike with listener certificates used for TLS negotiation that need to be replaced on a regular basis, this PKCS11 Cipher Stream Provider does not consider the validity period for the associated certificate, and it will continue to function even after the certificate has expired.  If you need to rotate the certificate used to protect the server's encryption settings database, you should first install the desired new certificate in the PKCS #11 token under a different alias. Then, you should create a new instance of this PKCS11 Cipher Stream Provider that is configured to use that certificate, and that also uses a different value for the encryption-metadata-file because the information in that file is tied to the certificate used to generate it. Finally, you will need to update the global configuration so that the encryption-settings-cipher-stream-provider property references the new cipher stream provider rather than this one. The update to the global configuration must be done with the server online so that it can properly re-encrypt the contents of the encryption settings database with the correct key tied to the new certificate.
	SslCertNickname string `json:"sslCertNickname"`
	// The path to a file that will hold metadata about the encryption performed by this PKCS11 Cipher Stream Provider.
	EncryptionMetadataFile string `json:"encryptionMetadataFile"`
	// A description for this Cipher Stream Provider
	Description *string `json:"description,omitempty"`
	// Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server.
	Enabled bool `json:"enabled"`
}

// NewAddPkcs11CipherStreamProviderRequest instantiates a new AddPkcs11CipherStreamProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPkcs11CipherStreamProviderRequest(providerName string, schemas []Enumpkcs11CipherStreamProviderSchemaUrn, sslCertNickname string, encryptionMetadataFile string, enabled bool) *AddPkcs11CipherStreamProviderRequest {
	this := AddPkcs11CipherStreamProviderRequest{}
	this.ProviderName = providerName
	this.Schemas = schemas
	this.SslCertNickname = sslCertNickname
	this.EncryptionMetadataFile = encryptionMetadataFile
	this.Enabled = enabled
	return &this
}

// NewAddPkcs11CipherStreamProviderRequestWithDefaults instantiates a new AddPkcs11CipherStreamProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPkcs11CipherStreamProviderRequestWithDefaults() *AddPkcs11CipherStreamProviderRequest {
	this := AddPkcs11CipherStreamProviderRequest{}
	return &this
}

// GetProviderName returns the ProviderName field value
func (o *AddPkcs11CipherStreamProviderRequest) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *AddPkcs11CipherStreamProviderRequest) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *AddPkcs11CipherStreamProviderRequest) SetProviderName(v string) {
	o.ProviderName = v
}

// GetSchemas returns the Schemas field value
func (o *AddPkcs11CipherStreamProviderRequest) GetSchemas() []Enumpkcs11CipherStreamProviderSchemaUrn {
	if o == nil {
		var ret []Enumpkcs11CipherStreamProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddPkcs11CipherStreamProviderRequest) GetSchemasOk() ([]Enumpkcs11CipherStreamProviderSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddPkcs11CipherStreamProviderRequest) SetSchemas(v []Enumpkcs11CipherStreamProviderSchemaUrn) {
	o.Schemas = v
}

// GetPkcs11ProviderClass returns the Pkcs11ProviderClass field value if set, zero value otherwise.
func (o *AddPkcs11CipherStreamProviderRequest) GetPkcs11ProviderClass() string {
	if o == nil || isNil(o.Pkcs11ProviderClass) {
		var ret string
		return ret
	}
	return *o.Pkcs11ProviderClass
}

// GetPkcs11ProviderClassOk returns a tuple with the Pkcs11ProviderClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPkcs11CipherStreamProviderRequest) GetPkcs11ProviderClassOk() (*string, bool) {
	if o == nil || isNil(o.Pkcs11ProviderClass) {
		return nil, false
	}
	return o.Pkcs11ProviderClass, true
}

// HasPkcs11ProviderClass returns a boolean if a field has been set.
func (o *AddPkcs11CipherStreamProviderRequest) HasPkcs11ProviderClass() bool {
	if o != nil && !isNil(o.Pkcs11ProviderClass) {
		return true
	}

	return false
}

// SetPkcs11ProviderClass gets a reference to the given string and assigns it to the Pkcs11ProviderClass field.
func (o *AddPkcs11CipherStreamProviderRequest) SetPkcs11ProviderClass(v string) {
	o.Pkcs11ProviderClass = &v
}

// GetPkcs11ProviderConfigurationFile returns the Pkcs11ProviderConfigurationFile field value if set, zero value otherwise.
func (o *AddPkcs11CipherStreamProviderRequest) GetPkcs11ProviderConfigurationFile() string {
	if o == nil || isNil(o.Pkcs11ProviderConfigurationFile) {
		var ret string
		return ret
	}
	return *o.Pkcs11ProviderConfigurationFile
}

// GetPkcs11ProviderConfigurationFileOk returns a tuple with the Pkcs11ProviderConfigurationFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPkcs11CipherStreamProviderRequest) GetPkcs11ProviderConfigurationFileOk() (*string, bool) {
	if o == nil || isNil(o.Pkcs11ProviderConfigurationFile) {
		return nil, false
	}
	return o.Pkcs11ProviderConfigurationFile, true
}

// HasPkcs11ProviderConfigurationFile returns a boolean if a field has been set.
func (o *AddPkcs11CipherStreamProviderRequest) HasPkcs11ProviderConfigurationFile() bool {
	if o != nil && !isNil(o.Pkcs11ProviderConfigurationFile) {
		return true
	}

	return false
}

// SetPkcs11ProviderConfigurationFile gets a reference to the given string and assigns it to the Pkcs11ProviderConfigurationFile field.
func (o *AddPkcs11CipherStreamProviderRequest) SetPkcs11ProviderConfigurationFile(v string) {
	o.Pkcs11ProviderConfigurationFile = &v
}

// GetKeyStorePin returns the KeyStorePin field value if set, zero value otherwise.
func (o *AddPkcs11CipherStreamProviderRequest) GetKeyStorePin() string {
	if o == nil || isNil(o.KeyStorePin) {
		var ret string
		return ret
	}
	return *o.KeyStorePin
}

// GetKeyStorePinOk returns a tuple with the KeyStorePin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPkcs11CipherStreamProviderRequest) GetKeyStorePinOk() (*string, bool) {
	if o == nil || isNil(o.KeyStorePin) {
		return nil, false
	}
	return o.KeyStorePin, true
}

// HasKeyStorePin returns a boolean if a field has been set.
func (o *AddPkcs11CipherStreamProviderRequest) HasKeyStorePin() bool {
	if o != nil && !isNil(o.KeyStorePin) {
		return true
	}

	return false
}

// SetKeyStorePin gets a reference to the given string and assigns it to the KeyStorePin field.
func (o *AddPkcs11CipherStreamProviderRequest) SetKeyStorePin(v string) {
	o.KeyStorePin = &v
}

// GetKeyStorePinFile returns the KeyStorePinFile field value if set, zero value otherwise.
func (o *AddPkcs11CipherStreamProviderRequest) GetKeyStorePinFile() string {
	if o == nil || isNil(o.KeyStorePinFile) {
		var ret string
		return ret
	}
	return *o.KeyStorePinFile
}

// GetKeyStorePinFileOk returns a tuple with the KeyStorePinFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPkcs11CipherStreamProviderRequest) GetKeyStorePinFileOk() (*string, bool) {
	if o == nil || isNil(o.KeyStorePinFile) {
		return nil, false
	}
	return o.KeyStorePinFile, true
}

// HasKeyStorePinFile returns a boolean if a field has been set.
func (o *AddPkcs11CipherStreamProviderRequest) HasKeyStorePinFile() bool {
	if o != nil && !isNil(o.KeyStorePinFile) {
		return true
	}

	return false
}

// SetKeyStorePinFile gets a reference to the given string and assigns it to the KeyStorePinFile field.
func (o *AddPkcs11CipherStreamProviderRequest) SetKeyStorePinFile(v string) {
	o.KeyStorePinFile = &v
}

// GetKeyStorePinEnvironmentVariable returns the KeyStorePinEnvironmentVariable field value if set, zero value otherwise.
func (o *AddPkcs11CipherStreamProviderRequest) GetKeyStorePinEnvironmentVariable() string {
	if o == nil || isNil(o.KeyStorePinEnvironmentVariable) {
		var ret string
		return ret
	}
	return *o.KeyStorePinEnvironmentVariable
}

// GetKeyStorePinEnvironmentVariableOk returns a tuple with the KeyStorePinEnvironmentVariable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPkcs11CipherStreamProviderRequest) GetKeyStorePinEnvironmentVariableOk() (*string, bool) {
	if o == nil || isNil(o.KeyStorePinEnvironmentVariable) {
		return nil, false
	}
	return o.KeyStorePinEnvironmentVariable, true
}

// HasKeyStorePinEnvironmentVariable returns a boolean if a field has been set.
func (o *AddPkcs11CipherStreamProviderRequest) HasKeyStorePinEnvironmentVariable() bool {
	if o != nil && !isNil(o.KeyStorePinEnvironmentVariable) {
		return true
	}

	return false
}

// SetKeyStorePinEnvironmentVariable gets a reference to the given string and assigns it to the KeyStorePinEnvironmentVariable field.
func (o *AddPkcs11CipherStreamProviderRequest) SetKeyStorePinEnvironmentVariable(v string) {
	o.KeyStorePinEnvironmentVariable = &v
}

// GetPkcs11KeyStoreType returns the Pkcs11KeyStoreType field value if set, zero value otherwise.
func (o *AddPkcs11CipherStreamProviderRequest) GetPkcs11KeyStoreType() string {
	if o == nil || isNil(o.Pkcs11KeyStoreType) {
		var ret string
		return ret
	}
	return *o.Pkcs11KeyStoreType
}

// GetPkcs11KeyStoreTypeOk returns a tuple with the Pkcs11KeyStoreType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPkcs11CipherStreamProviderRequest) GetPkcs11KeyStoreTypeOk() (*string, bool) {
	if o == nil || isNil(o.Pkcs11KeyStoreType) {
		return nil, false
	}
	return o.Pkcs11KeyStoreType, true
}

// HasPkcs11KeyStoreType returns a boolean if a field has been set.
func (o *AddPkcs11CipherStreamProviderRequest) HasPkcs11KeyStoreType() bool {
	if o != nil && !isNil(o.Pkcs11KeyStoreType) {
		return true
	}

	return false
}

// SetPkcs11KeyStoreType gets a reference to the given string and assigns it to the Pkcs11KeyStoreType field.
func (o *AddPkcs11CipherStreamProviderRequest) SetPkcs11KeyStoreType(v string) {
	o.Pkcs11KeyStoreType = &v
}

// GetSslCertNickname returns the SslCertNickname field value
func (o *AddPkcs11CipherStreamProviderRequest) GetSslCertNickname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SslCertNickname
}

// GetSslCertNicknameOk returns a tuple with the SslCertNickname field value
// and a boolean to check if the value has been set.
func (o *AddPkcs11CipherStreamProviderRequest) GetSslCertNicknameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SslCertNickname, true
}

// SetSslCertNickname sets field value
func (o *AddPkcs11CipherStreamProviderRequest) SetSslCertNickname(v string) {
	o.SslCertNickname = v
}

// GetEncryptionMetadataFile returns the EncryptionMetadataFile field value
func (o *AddPkcs11CipherStreamProviderRequest) GetEncryptionMetadataFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EncryptionMetadataFile
}

// GetEncryptionMetadataFileOk returns a tuple with the EncryptionMetadataFile field value
// and a boolean to check if the value has been set.
func (o *AddPkcs11CipherStreamProviderRequest) GetEncryptionMetadataFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EncryptionMetadataFile, true
}

// SetEncryptionMetadataFile sets field value
func (o *AddPkcs11CipherStreamProviderRequest) SetEncryptionMetadataFile(v string) {
	o.EncryptionMetadataFile = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddPkcs11CipherStreamProviderRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPkcs11CipherStreamProviderRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddPkcs11CipherStreamProviderRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddPkcs11CipherStreamProviderRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddPkcs11CipherStreamProviderRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddPkcs11CipherStreamProviderRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddPkcs11CipherStreamProviderRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddPkcs11CipherStreamProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["providerName"] = o.ProviderName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.Pkcs11ProviderClass) {
		toSerialize["pkcs11ProviderClass"] = o.Pkcs11ProviderClass
	}
	if !isNil(o.Pkcs11ProviderConfigurationFile) {
		toSerialize["pkcs11ProviderConfigurationFile"] = o.Pkcs11ProviderConfigurationFile
	}
	if !isNil(o.KeyStorePin) {
		toSerialize["keyStorePin"] = o.KeyStorePin
	}
	if !isNil(o.KeyStorePinFile) {
		toSerialize["keyStorePinFile"] = o.KeyStorePinFile
	}
	if !isNil(o.KeyStorePinEnvironmentVariable) {
		toSerialize["keyStorePinEnvironmentVariable"] = o.KeyStorePinEnvironmentVariable
	}
	if !isNil(o.Pkcs11KeyStoreType) {
		toSerialize["pkcs11KeyStoreType"] = o.Pkcs11KeyStoreType
	}
	if true {
		toSerialize["sslCertNickname"] = o.SslCertNickname
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

type NullableAddPkcs11CipherStreamProviderRequest struct {
	value *AddPkcs11CipherStreamProviderRequest
	isSet bool
}

func (v NullableAddPkcs11CipherStreamProviderRequest) Get() *AddPkcs11CipherStreamProviderRequest {
	return v.value
}

func (v *NullableAddPkcs11CipherStreamProviderRequest) Set(val *AddPkcs11CipherStreamProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPkcs11CipherStreamProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPkcs11CipherStreamProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPkcs11CipherStreamProviderRequest(val *AddPkcs11CipherStreamProviderRequest) *NullableAddPkcs11CipherStreamProviderRequest {
	return &NullableAddPkcs11CipherStreamProviderRequest{value: val, isSet: true}
}

func (v NullableAddPkcs11CipherStreamProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPkcs11CipherStreamProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
