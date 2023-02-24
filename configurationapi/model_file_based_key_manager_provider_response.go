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

// checks if the FileBasedKeyManagerProviderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileBasedKeyManagerProviderResponse{}

// FileBasedKeyManagerProviderResponse struct for FileBasedKeyManagerProviderResponse
type FileBasedKeyManagerProviderResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Key Manager Provider
	Id      string                                     `json:"id"`
	Schemas []EnumfileBasedKeyManagerProviderSchemaUrn `json:"schemas"`
	// Specifies the path to the file that contains the private key information. This may be an absolute path, or a path that is relative to the Directory Server instance root.
	KeyStoreFile string `json:"keyStoreFile"`
	// Specifies the format for the data in the key store file.
	KeyStoreType *string `json:"keyStoreType,omitempty"`
	// Specifies the PIN needed to access the File Based Key Manager Provider.
	KeyStorePin *string `json:"keyStorePin,omitempty"`
	// Specifies the path to the text file whose only contents should be a single line containing the clear-text PIN needed to access the File Based Key Manager Provider.
	KeyStorePinFile *string `json:"keyStorePinFile,omitempty"`
	// The passphrase provider to use to obtain the clear-text PIN needed to access the File Based Key Manager Provider.
	KeyStorePinPassphraseProvider *string `json:"keyStorePinPassphraseProvider,omitempty"`
	// Specifies the clear-text PIN needed to access the File Based Key Manager Provider private key. If no private key PIN is specified the PIN defaults to the key store PIN.
	PrivateKeyPin *string `json:"privateKeyPin,omitempty"`
	// Specifies the path to the text file whose only contents should be a single line containing the clear-text PIN needed to access the File Based Key Manager Provider private key. If no private key PIN is specified the PIN defaults to the key store PIN.
	PrivateKeyPinFile *string `json:"privateKeyPinFile,omitempty"`
	// The passphrase provider to use to obtain the clear-text PIN needed to access the File Based Key Manager Provider private key. If no private key PIN is specified the PIN defaults to the key store PIN.
	PrivateKeyPinPassphraseProvider *string `json:"privateKeyPinPassphraseProvider,omitempty"`
	// A description for this Key Manager Provider
	Description *string `json:"description,omitempty"`
	// Indicates whether the Key Manager Provider is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewFileBasedKeyManagerProviderResponse instantiates a new FileBasedKeyManagerProviderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileBasedKeyManagerProviderResponse(id string, schemas []EnumfileBasedKeyManagerProviderSchemaUrn, keyStoreFile string, enabled bool) *FileBasedKeyManagerProviderResponse {
	this := FileBasedKeyManagerProviderResponse{}
	this.Id = id
	this.Schemas = schemas
	this.KeyStoreFile = keyStoreFile
	this.Enabled = enabled
	return &this
}

// NewFileBasedKeyManagerProviderResponseWithDefaults instantiates a new FileBasedKeyManagerProviderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileBasedKeyManagerProviderResponseWithDefaults() *FileBasedKeyManagerProviderResponse {
	this := FileBasedKeyManagerProviderResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *FileBasedKeyManagerProviderResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedKeyManagerProviderResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *FileBasedKeyManagerProviderResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *FileBasedKeyManagerProviderResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *FileBasedKeyManagerProviderResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedKeyManagerProviderResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *FileBasedKeyManagerProviderResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *FileBasedKeyManagerProviderResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *FileBasedKeyManagerProviderResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FileBasedKeyManagerProviderResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FileBasedKeyManagerProviderResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *FileBasedKeyManagerProviderResponse) GetSchemas() []EnumfileBasedKeyManagerProviderSchemaUrn {
	if o == nil {
		var ret []EnumfileBasedKeyManagerProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *FileBasedKeyManagerProviderResponse) GetSchemasOk() ([]EnumfileBasedKeyManagerProviderSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *FileBasedKeyManagerProviderResponse) SetSchemas(v []EnumfileBasedKeyManagerProviderSchemaUrn) {
	o.Schemas = v
}

// GetKeyStoreFile returns the KeyStoreFile field value
func (o *FileBasedKeyManagerProviderResponse) GetKeyStoreFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyStoreFile
}

// GetKeyStoreFileOk returns a tuple with the KeyStoreFile field value
// and a boolean to check if the value has been set.
func (o *FileBasedKeyManagerProviderResponse) GetKeyStoreFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyStoreFile, true
}

// SetKeyStoreFile sets field value
func (o *FileBasedKeyManagerProviderResponse) SetKeyStoreFile(v string) {
	o.KeyStoreFile = v
}

// GetKeyStoreType returns the KeyStoreType field value if set, zero value otherwise.
func (o *FileBasedKeyManagerProviderResponse) GetKeyStoreType() string {
	if o == nil || IsNil(o.KeyStoreType) {
		var ret string
		return ret
	}
	return *o.KeyStoreType
}

// GetKeyStoreTypeOk returns a tuple with the KeyStoreType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedKeyManagerProviderResponse) GetKeyStoreTypeOk() (*string, bool) {
	if o == nil || IsNil(o.KeyStoreType) {
		return nil, false
	}
	return o.KeyStoreType, true
}

// HasKeyStoreType returns a boolean if a field has been set.
func (o *FileBasedKeyManagerProviderResponse) HasKeyStoreType() bool {
	if o != nil && !IsNil(o.KeyStoreType) {
		return true
	}

	return false
}

// SetKeyStoreType gets a reference to the given string and assigns it to the KeyStoreType field.
func (o *FileBasedKeyManagerProviderResponse) SetKeyStoreType(v string) {
	o.KeyStoreType = &v
}

// GetKeyStorePin returns the KeyStorePin field value if set, zero value otherwise.
func (o *FileBasedKeyManagerProviderResponse) GetKeyStorePin() string {
	if o == nil || IsNil(o.KeyStorePin) {
		var ret string
		return ret
	}
	return *o.KeyStorePin
}

// GetKeyStorePinOk returns a tuple with the KeyStorePin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedKeyManagerProviderResponse) GetKeyStorePinOk() (*string, bool) {
	if o == nil || IsNil(o.KeyStorePin) {
		return nil, false
	}
	return o.KeyStorePin, true
}

// HasKeyStorePin returns a boolean if a field has been set.
func (o *FileBasedKeyManagerProviderResponse) HasKeyStorePin() bool {
	if o != nil && !IsNil(o.KeyStorePin) {
		return true
	}

	return false
}

// SetKeyStorePin gets a reference to the given string and assigns it to the KeyStorePin field.
func (o *FileBasedKeyManagerProviderResponse) SetKeyStorePin(v string) {
	o.KeyStorePin = &v
}

// GetKeyStorePinFile returns the KeyStorePinFile field value if set, zero value otherwise.
func (o *FileBasedKeyManagerProviderResponse) GetKeyStorePinFile() string {
	if o == nil || IsNil(o.KeyStorePinFile) {
		var ret string
		return ret
	}
	return *o.KeyStorePinFile
}

// GetKeyStorePinFileOk returns a tuple with the KeyStorePinFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedKeyManagerProviderResponse) GetKeyStorePinFileOk() (*string, bool) {
	if o == nil || IsNil(o.KeyStorePinFile) {
		return nil, false
	}
	return o.KeyStorePinFile, true
}

// HasKeyStorePinFile returns a boolean if a field has been set.
func (o *FileBasedKeyManagerProviderResponse) HasKeyStorePinFile() bool {
	if o != nil && !IsNil(o.KeyStorePinFile) {
		return true
	}

	return false
}

// SetKeyStorePinFile gets a reference to the given string and assigns it to the KeyStorePinFile field.
func (o *FileBasedKeyManagerProviderResponse) SetKeyStorePinFile(v string) {
	o.KeyStorePinFile = &v
}

// GetKeyStorePinPassphraseProvider returns the KeyStorePinPassphraseProvider field value if set, zero value otherwise.
func (o *FileBasedKeyManagerProviderResponse) GetKeyStorePinPassphraseProvider() string {
	if o == nil || IsNil(o.KeyStorePinPassphraseProvider) {
		var ret string
		return ret
	}
	return *o.KeyStorePinPassphraseProvider
}

// GetKeyStorePinPassphraseProviderOk returns a tuple with the KeyStorePinPassphraseProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedKeyManagerProviderResponse) GetKeyStorePinPassphraseProviderOk() (*string, bool) {
	if o == nil || IsNil(o.KeyStorePinPassphraseProvider) {
		return nil, false
	}
	return o.KeyStorePinPassphraseProvider, true
}

// HasKeyStorePinPassphraseProvider returns a boolean if a field has been set.
func (o *FileBasedKeyManagerProviderResponse) HasKeyStorePinPassphraseProvider() bool {
	if o != nil && !IsNil(o.KeyStorePinPassphraseProvider) {
		return true
	}

	return false
}

// SetKeyStorePinPassphraseProvider gets a reference to the given string and assigns it to the KeyStorePinPassphraseProvider field.
func (o *FileBasedKeyManagerProviderResponse) SetKeyStorePinPassphraseProvider(v string) {
	o.KeyStorePinPassphraseProvider = &v
}

// GetPrivateKeyPin returns the PrivateKeyPin field value if set, zero value otherwise.
func (o *FileBasedKeyManagerProviderResponse) GetPrivateKeyPin() string {
	if o == nil || IsNil(o.PrivateKeyPin) {
		var ret string
		return ret
	}
	return *o.PrivateKeyPin
}

// GetPrivateKeyPinOk returns a tuple with the PrivateKeyPin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedKeyManagerProviderResponse) GetPrivateKeyPinOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateKeyPin) {
		return nil, false
	}
	return o.PrivateKeyPin, true
}

// HasPrivateKeyPin returns a boolean if a field has been set.
func (o *FileBasedKeyManagerProviderResponse) HasPrivateKeyPin() bool {
	if o != nil && !IsNil(o.PrivateKeyPin) {
		return true
	}

	return false
}

// SetPrivateKeyPin gets a reference to the given string and assigns it to the PrivateKeyPin field.
func (o *FileBasedKeyManagerProviderResponse) SetPrivateKeyPin(v string) {
	o.PrivateKeyPin = &v
}

// GetPrivateKeyPinFile returns the PrivateKeyPinFile field value if set, zero value otherwise.
func (o *FileBasedKeyManagerProviderResponse) GetPrivateKeyPinFile() string {
	if o == nil || IsNil(o.PrivateKeyPinFile) {
		var ret string
		return ret
	}
	return *o.PrivateKeyPinFile
}

// GetPrivateKeyPinFileOk returns a tuple with the PrivateKeyPinFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedKeyManagerProviderResponse) GetPrivateKeyPinFileOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateKeyPinFile) {
		return nil, false
	}
	return o.PrivateKeyPinFile, true
}

// HasPrivateKeyPinFile returns a boolean if a field has been set.
func (o *FileBasedKeyManagerProviderResponse) HasPrivateKeyPinFile() bool {
	if o != nil && !IsNil(o.PrivateKeyPinFile) {
		return true
	}

	return false
}

// SetPrivateKeyPinFile gets a reference to the given string and assigns it to the PrivateKeyPinFile field.
func (o *FileBasedKeyManagerProviderResponse) SetPrivateKeyPinFile(v string) {
	o.PrivateKeyPinFile = &v
}

// GetPrivateKeyPinPassphraseProvider returns the PrivateKeyPinPassphraseProvider field value if set, zero value otherwise.
func (o *FileBasedKeyManagerProviderResponse) GetPrivateKeyPinPassphraseProvider() string {
	if o == nil || IsNil(o.PrivateKeyPinPassphraseProvider) {
		var ret string
		return ret
	}
	return *o.PrivateKeyPinPassphraseProvider
}

// GetPrivateKeyPinPassphraseProviderOk returns a tuple with the PrivateKeyPinPassphraseProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedKeyManagerProviderResponse) GetPrivateKeyPinPassphraseProviderOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateKeyPinPassphraseProvider) {
		return nil, false
	}
	return o.PrivateKeyPinPassphraseProvider, true
}

// HasPrivateKeyPinPassphraseProvider returns a boolean if a field has been set.
func (o *FileBasedKeyManagerProviderResponse) HasPrivateKeyPinPassphraseProvider() bool {
	if o != nil && !IsNil(o.PrivateKeyPinPassphraseProvider) {
		return true
	}

	return false
}

// SetPrivateKeyPinPassphraseProvider gets a reference to the given string and assigns it to the PrivateKeyPinPassphraseProvider field.
func (o *FileBasedKeyManagerProviderResponse) SetPrivateKeyPinPassphraseProvider(v string) {
	o.PrivateKeyPinPassphraseProvider = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FileBasedKeyManagerProviderResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedKeyManagerProviderResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FileBasedKeyManagerProviderResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FileBasedKeyManagerProviderResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *FileBasedKeyManagerProviderResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *FileBasedKeyManagerProviderResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *FileBasedKeyManagerProviderResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o FileBasedKeyManagerProviderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileBasedKeyManagerProviderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	toSerialize["schemas"] = o.Schemas
	toSerialize["keyStoreFile"] = o.KeyStoreFile
	if !IsNil(o.KeyStoreType) {
		toSerialize["keyStoreType"] = o.KeyStoreType
	}
	if !IsNil(o.KeyStorePin) {
		toSerialize["keyStorePin"] = o.KeyStorePin
	}
	if !IsNil(o.KeyStorePinFile) {
		toSerialize["keyStorePinFile"] = o.KeyStorePinFile
	}
	if !IsNil(o.KeyStorePinPassphraseProvider) {
		toSerialize["keyStorePinPassphraseProvider"] = o.KeyStorePinPassphraseProvider
	}
	if !IsNil(o.PrivateKeyPin) {
		toSerialize["privateKeyPin"] = o.PrivateKeyPin
	}
	if !IsNil(o.PrivateKeyPinFile) {
		toSerialize["privateKeyPinFile"] = o.PrivateKeyPinFile
	}
	if !IsNil(o.PrivateKeyPinPassphraseProvider) {
		toSerialize["privateKeyPinPassphraseProvider"] = o.PrivateKeyPinPassphraseProvider
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableFileBasedKeyManagerProviderResponse struct {
	value *FileBasedKeyManagerProviderResponse
	isSet bool
}

func (v NullableFileBasedKeyManagerProviderResponse) Get() *FileBasedKeyManagerProviderResponse {
	return v.value
}

func (v *NullableFileBasedKeyManagerProviderResponse) Set(val *FileBasedKeyManagerProviderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFileBasedKeyManagerProviderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFileBasedKeyManagerProviderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileBasedKeyManagerProviderResponse(val *FileBasedKeyManagerProviderResponse) *NullableFileBasedKeyManagerProviderResponse {
	return &NullableFileBasedKeyManagerProviderResponse{value: val, isSet: true}
}

func (v NullableFileBasedKeyManagerProviderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileBasedKeyManagerProviderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
