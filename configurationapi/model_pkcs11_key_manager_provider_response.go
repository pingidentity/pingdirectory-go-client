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

// checks if the Pkcs11KeyManagerProviderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Pkcs11KeyManagerProviderResponse{}

// Pkcs11KeyManagerProviderResponse struct for Pkcs11KeyManagerProviderResponse
type Pkcs11KeyManagerProviderResponse struct {
	Schemas []Enumpkcs11KeyManagerProviderSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java security provider class that implements support for interacting with PKCS #11 tokens.
	Pkcs11ProviderClass *string `json:"pkcs11ProviderClass,omitempty"`
	// The path to the file to use to configure the security provider that implements support for interacting with PKCS #11 tokens.
	Pkcs11ProviderConfigurationFile *string `json:"pkcs11ProviderConfigurationFile,omitempty"`
	// The key store type to use when obtaining an instance of a key store for interacting with a PKCS #11 token.
	Pkcs11KeyStoreType *string `json:"pkcs11KeyStoreType,omitempty"`
	// The maximum length of time that data retrieved from PKCS #11 tokens may be cached for reuse. Caching might be necessary if there is noticable latency when accessing the token, for example if the token uses a remote key store. A value of zero milliseconds indicates that no caching should be performed.
	Pkcs11MaxCacheDuration *string `json:"pkcs11MaxCacheDuration,omitempty"`
	// Specifies the PIN needed to access the PKCS11 Key Manager Provider.
	KeyStorePin *string `json:"keyStorePin,omitempty"`
	// Specifies the path to the text file whose only contents should be a single line containing the clear-text PIN needed to access the PKCS11 Key Manager Provider.
	KeyStorePinFile *string `json:"keyStorePinFile,omitempty"`
	// The passphrase provider to use to obtain the clear-text PIN needed to access the PKCS11 Key Manager Provider.
	KeyStorePinPassphraseProvider *string `json:"keyStorePinPassphraseProvider,omitempty"`
	// A description for this Key Manager Provider
	Description *string `json:"description,omitempty"`
	// Indicates whether the Key Manager Provider is enabled for use.
	Enabled                                       bool                                               `json:"enabled"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Key Manager Provider
	Id string `json:"id"`
}

// NewPkcs11KeyManagerProviderResponse instantiates a new Pkcs11KeyManagerProviderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkcs11KeyManagerProviderResponse(schemas []Enumpkcs11KeyManagerProviderSchemaUrn, enabled bool, id string) *Pkcs11KeyManagerProviderResponse {
	this := Pkcs11KeyManagerProviderResponse{}
	this.Schemas = schemas
	this.Enabled = enabled
	this.Id = id
	return &this
}

// NewPkcs11KeyManagerProviderResponseWithDefaults instantiates a new Pkcs11KeyManagerProviderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkcs11KeyManagerProviderResponseWithDefaults() *Pkcs11KeyManagerProviderResponse {
	this := Pkcs11KeyManagerProviderResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *Pkcs11KeyManagerProviderResponse) GetSchemas() []Enumpkcs11KeyManagerProviderSchemaUrn {
	if o == nil {
		var ret []Enumpkcs11KeyManagerProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *Pkcs11KeyManagerProviderResponse) GetSchemasOk() ([]Enumpkcs11KeyManagerProviderSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *Pkcs11KeyManagerProviderResponse) SetSchemas(v []Enumpkcs11KeyManagerProviderSchemaUrn) {
	o.Schemas = v
}

// GetPkcs11ProviderClass returns the Pkcs11ProviderClass field value if set, zero value otherwise.
func (o *Pkcs11KeyManagerProviderResponse) GetPkcs11ProviderClass() string {
	if o == nil || IsNil(o.Pkcs11ProviderClass) {
		var ret string
		return ret
	}
	return *o.Pkcs11ProviderClass
}

// GetPkcs11ProviderClassOk returns a tuple with the Pkcs11ProviderClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pkcs11KeyManagerProviderResponse) GetPkcs11ProviderClassOk() (*string, bool) {
	if o == nil || IsNil(o.Pkcs11ProviderClass) {
		return nil, false
	}
	return o.Pkcs11ProviderClass, true
}

// HasPkcs11ProviderClass returns a boolean if a field has been set.
func (o *Pkcs11KeyManagerProviderResponse) HasPkcs11ProviderClass() bool {
	if o != nil && !IsNil(o.Pkcs11ProviderClass) {
		return true
	}

	return false
}

// SetPkcs11ProviderClass gets a reference to the given string and assigns it to the Pkcs11ProviderClass field.
func (o *Pkcs11KeyManagerProviderResponse) SetPkcs11ProviderClass(v string) {
	o.Pkcs11ProviderClass = &v
}

// GetPkcs11ProviderConfigurationFile returns the Pkcs11ProviderConfigurationFile field value if set, zero value otherwise.
func (o *Pkcs11KeyManagerProviderResponse) GetPkcs11ProviderConfigurationFile() string {
	if o == nil || IsNil(o.Pkcs11ProviderConfigurationFile) {
		var ret string
		return ret
	}
	return *o.Pkcs11ProviderConfigurationFile
}

// GetPkcs11ProviderConfigurationFileOk returns a tuple with the Pkcs11ProviderConfigurationFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pkcs11KeyManagerProviderResponse) GetPkcs11ProviderConfigurationFileOk() (*string, bool) {
	if o == nil || IsNil(o.Pkcs11ProviderConfigurationFile) {
		return nil, false
	}
	return o.Pkcs11ProviderConfigurationFile, true
}

// HasPkcs11ProviderConfigurationFile returns a boolean if a field has been set.
func (o *Pkcs11KeyManagerProviderResponse) HasPkcs11ProviderConfigurationFile() bool {
	if o != nil && !IsNil(o.Pkcs11ProviderConfigurationFile) {
		return true
	}

	return false
}

// SetPkcs11ProviderConfigurationFile gets a reference to the given string and assigns it to the Pkcs11ProviderConfigurationFile field.
func (o *Pkcs11KeyManagerProviderResponse) SetPkcs11ProviderConfigurationFile(v string) {
	o.Pkcs11ProviderConfigurationFile = &v
}

// GetPkcs11KeyStoreType returns the Pkcs11KeyStoreType field value if set, zero value otherwise.
func (o *Pkcs11KeyManagerProviderResponse) GetPkcs11KeyStoreType() string {
	if o == nil || IsNil(o.Pkcs11KeyStoreType) {
		var ret string
		return ret
	}
	return *o.Pkcs11KeyStoreType
}

// GetPkcs11KeyStoreTypeOk returns a tuple with the Pkcs11KeyStoreType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pkcs11KeyManagerProviderResponse) GetPkcs11KeyStoreTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Pkcs11KeyStoreType) {
		return nil, false
	}
	return o.Pkcs11KeyStoreType, true
}

// HasPkcs11KeyStoreType returns a boolean if a field has been set.
func (o *Pkcs11KeyManagerProviderResponse) HasPkcs11KeyStoreType() bool {
	if o != nil && !IsNil(o.Pkcs11KeyStoreType) {
		return true
	}

	return false
}

// SetPkcs11KeyStoreType gets a reference to the given string and assigns it to the Pkcs11KeyStoreType field.
func (o *Pkcs11KeyManagerProviderResponse) SetPkcs11KeyStoreType(v string) {
	o.Pkcs11KeyStoreType = &v
}

// GetPkcs11MaxCacheDuration returns the Pkcs11MaxCacheDuration field value if set, zero value otherwise.
func (o *Pkcs11KeyManagerProviderResponse) GetPkcs11MaxCacheDuration() string {
	if o == nil || IsNil(o.Pkcs11MaxCacheDuration) {
		var ret string
		return ret
	}
	return *o.Pkcs11MaxCacheDuration
}

// GetPkcs11MaxCacheDurationOk returns a tuple with the Pkcs11MaxCacheDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pkcs11KeyManagerProviderResponse) GetPkcs11MaxCacheDurationOk() (*string, bool) {
	if o == nil || IsNil(o.Pkcs11MaxCacheDuration) {
		return nil, false
	}
	return o.Pkcs11MaxCacheDuration, true
}

// HasPkcs11MaxCacheDuration returns a boolean if a field has been set.
func (o *Pkcs11KeyManagerProviderResponse) HasPkcs11MaxCacheDuration() bool {
	if o != nil && !IsNil(o.Pkcs11MaxCacheDuration) {
		return true
	}

	return false
}

// SetPkcs11MaxCacheDuration gets a reference to the given string and assigns it to the Pkcs11MaxCacheDuration field.
func (o *Pkcs11KeyManagerProviderResponse) SetPkcs11MaxCacheDuration(v string) {
	o.Pkcs11MaxCacheDuration = &v
}

// GetKeyStorePin returns the KeyStorePin field value if set, zero value otherwise.
func (o *Pkcs11KeyManagerProviderResponse) GetKeyStorePin() string {
	if o == nil || IsNil(o.KeyStorePin) {
		var ret string
		return ret
	}
	return *o.KeyStorePin
}

// GetKeyStorePinOk returns a tuple with the KeyStorePin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pkcs11KeyManagerProviderResponse) GetKeyStorePinOk() (*string, bool) {
	if o == nil || IsNil(o.KeyStorePin) {
		return nil, false
	}
	return o.KeyStorePin, true
}

// HasKeyStorePin returns a boolean if a field has been set.
func (o *Pkcs11KeyManagerProviderResponse) HasKeyStorePin() bool {
	if o != nil && !IsNil(o.KeyStorePin) {
		return true
	}

	return false
}

// SetKeyStorePin gets a reference to the given string and assigns it to the KeyStorePin field.
func (o *Pkcs11KeyManagerProviderResponse) SetKeyStorePin(v string) {
	o.KeyStorePin = &v
}

// GetKeyStorePinFile returns the KeyStorePinFile field value if set, zero value otherwise.
func (o *Pkcs11KeyManagerProviderResponse) GetKeyStorePinFile() string {
	if o == nil || IsNil(o.KeyStorePinFile) {
		var ret string
		return ret
	}
	return *o.KeyStorePinFile
}

// GetKeyStorePinFileOk returns a tuple with the KeyStorePinFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pkcs11KeyManagerProviderResponse) GetKeyStorePinFileOk() (*string, bool) {
	if o == nil || IsNil(o.KeyStorePinFile) {
		return nil, false
	}
	return o.KeyStorePinFile, true
}

// HasKeyStorePinFile returns a boolean if a field has been set.
func (o *Pkcs11KeyManagerProviderResponse) HasKeyStorePinFile() bool {
	if o != nil && !IsNil(o.KeyStorePinFile) {
		return true
	}

	return false
}

// SetKeyStorePinFile gets a reference to the given string and assigns it to the KeyStorePinFile field.
func (o *Pkcs11KeyManagerProviderResponse) SetKeyStorePinFile(v string) {
	o.KeyStorePinFile = &v
}

// GetKeyStorePinPassphraseProvider returns the KeyStorePinPassphraseProvider field value if set, zero value otherwise.
func (o *Pkcs11KeyManagerProviderResponse) GetKeyStorePinPassphraseProvider() string {
	if o == nil || IsNil(o.KeyStorePinPassphraseProvider) {
		var ret string
		return ret
	}
	return *o.KeyStorePinPassphraseProvider
}

// GetKeyStorePinPassphraseProviderOk returns a tuple with the KeyStorePinPassphraseProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pkcs11KeyManagerProviderResponse) GetKeyStorePinPassphraseProviderOk() (*string, bool) {
	if o == nil || IsNil(o.KeyStorePinPassphraseProvider) {
		return nil, false
	}
	return o.KeyStorePinPassphraseProvider, true
}

// HasKeyStorePinPassphraseProvider returns a boolean if a field has been set.
func (o *Pkcs11KeyManagerProviderResponse) HasKeyStorePinPassphraseProvider() bool {
	if o != nil && !IsNil(o.KeyStorePinPassphraseProvider) {
		return true
	}

	return false
}

// SetKeyStorePinPassphraseProvider gets a reference to the given string and assigns it to the KeyStorePinPassphraseProvider field.
func (o *Pkcs11KeyManagerProviderResponse) SetKeyStorePinPassphraseProvider(v string) {
	o.KeyStorePinPassphraseProvider = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Pkcs11KeyManagerProviderResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pkcs11KeyManagerProviderResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Pkcs11KeyManagerProviderResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Pkcs11KeyManagerProviderResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *Pkcs11KeyManagerProviderResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Pkcs11KeyManagerProviderResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *Pkcs11KeyManagerProviderResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Pkcs11KeyManagerProviderResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pkcs11KeyManagerProviderResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Pkcs11KeyManagerProviderResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *Pkcs11KeyManagerProviderResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *Pkcs11KeyManagerProviderResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pkcs11KeyManagerProviderResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *Pkcs11KeyManagerProviderResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *Pkcs11KeyManagerProviderResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *Pkcs11KeyManagerProviderResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Pkcs11KeyManagerProviderResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Pkcs11KeyManagerProviderResponse) SetId(v string) {
	o.Id = v
}

func (o Pkcs11KeyManagerProviderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Pkcs11KeyManagerProviderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.Pkcs11ProviderClass) {
		toSerialize["pkcs11ProviderClass"] = o.Pkcs11ProviderClass
	}
	if !IsNil(o.Pkcs11ProviderConfigurationFile) {
		toSerialize["pkcs11ProviderConfigurationFile"] = o.Pkcs11ProviderConfigurationFile
	}
	if !IsNil(o.Pkcs11KeyStoreType) {
		toSerialize["pkcs11KeyStoreType"] = o.Pkcs11KeyStoreType
	}
	if !IsNil(o.Pkcs11MaxCacheDuration) {
		toSerialize["pkcs11MaxCacheDuration"] = o.Pkcs11MaxCacheDuration
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

type NullablePkcs11KeyManagerProviderResponse struct {
	value *Pkcs11KeyManagerProviderResponse
	isSet bool
}

func (v NullablePkcs11KeyManagerProviderResponse) Get() *Pkcs11KeyManagerProviderResponse {
	return v.value
}

func (v *NullablePkcs11KeyManagerProviderResponse) Set(val *Pkcs11KeyManagerProviderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePkcs11KeyManagerProviderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePkcs11KeyManagerProviderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkcs11KeyManagerProviderResponse(val *Pkcs11KeyManagerProviderResponse) *NullablePkcs11KeyManagerProviderResponse {
	return &NullablePkcs11KeyManagerProviderResponse{value: val, isSet: true}
}

func (v NullablePkcs11KeyManagerProviderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkcs11KeyManagerProviderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
