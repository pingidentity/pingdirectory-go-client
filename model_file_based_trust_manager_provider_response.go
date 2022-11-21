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

// FileBasedTrustManagerProviderResponse struct for FileBasedTrustManagerProviderResponse
type FileBasedTrustManagerProviderResponse struct {
	// Name of the Trust Manager Provider
	Id string `json:"id"`
	Schemas []EnumfileBasedTrustManagerProviderSchemaUrn `json:"schemas"`
	// Specifies the path to the file containing the trust information. It can be an absolute path or a path that is relative to the Directory Server instance root.
	TrustStoreFile string `json:"trustStoreFile"`
	// Specifies the format for the data in the trust store file.
	TrustStoreType *string `json:"trustStoreType,omitempty"`
	// Specifies the clear-text PIN needed to access the File Based Trust Manager Provider.
	TrustStorePin *string `json:"trustStorePin,omitempty"`
	// Specifies the path to the text file whose only contents should be a single line containing the clear-text PIN needed to access the File Based Trust Manager Provider.
	TrustStorePinFile *string `json:"trustStorePinFile,omitempty"`
	// The passphrase provider to use to obtain the clear-text PIN needed to access the File Based Trust Manager Provider.
	TrustStorePinPassphraseProvider *string `json:"trustStorePinPassphraseProvider,omitempty"`
	// Indicate whether the Trust Manager Provider is enabled for use.
	Enabled bool `json:"enabled"`
	// Indicates whether certificates issued by an authority included in the JVM's set of default issuers should be automatically trusted, even if they would not otherwise be trusted by this provider.
	IncludeJVMDefaultIssuers *bool `json:"includeJVMDefaultIssuers,omitempty"`
	Meta *MetaMeta `json:"meta,omitempty"`
}

// NewFileBasedTrustManagerProviderResponse instantiates a new FileBasedTrustManagerProviderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileBasedTrustManagerProviderResponse(id string, schemas []EnumfileBasedTrustManagerProviderSchemaUrn, trustStoreFile string, enabled bool) *FileBasedTrustManagerProviderResponse {
	this := FileBasedTrustManagerProviderResponse{}
	this.Id = id
	this.Schemas = schemas
	this.TrustStoreFile = trustStoreFile
	this.Enabled = enabled
	return &this
}

// NewFileBasedTrustManagerProviderResponseWithDefaults instantiates a new FileBasedTrustManagerProviderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileBasedTrustManagerProviderResponseWithDefaults() *FileBasedTrustManagerProviderResponse {
	this := FileBasedTrustManagerProviderResponse{}
	return &this
}

// GetId returns the Id field value
func (o *FileBasedTrustManagerProviderResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FileBasedTrustManagerProviderResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FileBasedTrustManagerProviderResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *FileBasedTrustManagerProviderResponse) GetSchemas() []EnumfileBasedTrustManagerProviderSchemaUrn {
	if o == nil {
		var ret []EnumfileBasedTrustManagerProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *FileBasedTrustManagerProviderResponse) GetSchemasOk() ([]EnumfileBasedTrustManagerProviderSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *FileBasedTrustManagerProviderResponse) SetSchemas(v []EnumfileBasedTrustManagerProviderSchemaUrn) {
	o.Schemas = v
}

// GetTrustStoreFile returns the TrustStoreFile field value
func (o *FileBasedTrustManagerProviderResponse) GetTrustStoreFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TrustStoreFile
}

// GetTrustStoreFileOk returns a tuple with the TrustStoreFile field value
// and a boolean to check if the value has been set.
func (o *FileBasedTrustManagerProviderResponse) GetTrustStoreFileOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TrustStoreFile, true
}

// SetTrustStoreFile sets field value
func (o *FileBasedTrustManagerProviderResponse) SetTrustStoreFile(v string) {
	o.TrustStoreFile = v
}

// GetTrustStoreType returns the TrustStoreType field value if set, zero value otherwise.
func (o *FileBasedTrustManagerProviderResponse) GetTrustStoreType() string {
	if o == nil || isNil(o.TrustStoreType) {
		var ret string
		return ret
	}
	return *o.TrustStoreType
}

// GetTrustStoreTypeOk returns a tuple with the TrustStoreType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedTrustManagerProviderResponse) GetTrustStoreTypeOk() (*string, bool) {
	if o == nil || isNil(o.TrustStoreType) {
    return nil, false
	}
	return o.TrustStoreType, true
}

// HasTrustStoreType returns a boolean if a field has been set.
func (o *FileBasedTrustManagerProviderResponse) HasTrustStoreType() bool {
	if o != nil && !isNil(o.TrustStoreType) {
		return true
	}

	return false
}

// SetTrustStoreType gets a reference to the given string and assigns it to the TrustStoreType field.
func (o *FileBasedTrustManagerProviderResponse) SetTrustStoreType(v string) {
	o.TrustStoreType = &v
}

// GetTrustStorePin returns the TrustStorePin field value if set, zero value otherwise.
func (o *FileBasedTrustManagerProviderResponse) GetTrustStorePin() string {
	if o == nil || isNil(o.TrustStorePin) {
		var ret string
		return ret
	}
	return *o.TrustStorePin
}

// GetTrustStorePinOk returns a tuple with the TrustStorePin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedTrustManagerProviderResponse) GetTrustStorePinOk() (*string, bool) {
	if o == nil || isNil(o.TrustStorePin) {
    return nil, false
	}
	return o.TrustStorePin, true
}

// HasTrustStorePin returns a boolean if a field has been set.
func (o *FileBasedTrustManagerProviderResponse) HasTrustStorePin() bool {
	if o != nil && !isNil(o.TrustStorePin) {
		return true
	}

	return false
}

// SetTrustStorePin gets a reference to the given string and assigns it to the TrustStorePin field.
func (o *FileBasedTrustManagerProviderResponse) SetTrustStorePin(v string) {
	o.TrustStorePin = &v
}

// GetTrustStorePinFile returns the TrustStorePinFile field value if set, zero value otherwise.
func (o *FileBasedTrustManagerProviderResponse) GetTrustStorePinFile() string {
	if o == nil || isNil(o.TrustStorePinFile) {
		var ret string
		return ret
	}
	return *o.TrustStorePinFile
}

// GetTrustStorePinFileOk returns a tuple with the TrustStorePinFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedTrustManagerProviderResponse) GetTrustStorePinFileOk() (*string, bool) {
	if o == nil || isNil(o.TrustStorePinFile) {
    return nil, false
	}
	return o.TrustStorePinFile, true
}

// HasTrustStorePinFile returns a boolean if a field has been set.
func (o *FileBasedTrustManagerProviderResponse) HasTrustStorePinFile() bool {
	if o != nil && !isNil(o.TrustStorePinFile) {
		return true
	}

	return false
}

// SetTrustStorePinFile gets a reference to the given string and assigns it to the TrustStorePinFile field.
func (o *FileBasedTrustManagerProviderResponse) SetTrustStorePinFile(v string) {
	o.TrustStorePinFile = &v
}

// GetTrustStorePinPassphraseProvider returns the TrustStorePinPassphraseProvider field value if set, zero value otherwise.
func (o *FileBasedTrustManagerProviderResponse) GetTrustStorePinPassphraseProvider() string {
	if o == nil || isNil(o.TrustStorePinPassphraseProvider) {
		var ret string
		return ret
	}
	return *o.TrustStorePinPassphraseProvider
}

// GetTrustStorePinPassphraseProviderOk returns a tuple with the TrustStorePinPassphraseProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedTrustManagerProviderResponse) GetTrustStorePinPassphraseProviderOk() (*string, bool) {
	if o == nil || isNil(o.TrustStorePinPassphraseProvider) {
    return nil, false
	}
	return o.TrustStorePinPassphraseProvider, true
}

// HasTrustStorePinPassphraseProvider returns a boolean if a field has been set.
func (o *FileBasedTrustManagerProviderResponse) HasTrustStorePinPassphraseProvider() bool {
	if o != nil && !isNil(o.TrustStorePinPassphraseProvider) {
		return true
	}

	return false
}

// SetTrustStorePinPassphraseProvider gets a reference to the given string and assigns it to the TrustStorePinPassphraseProvider field.
func (o *FileBasedTrustManagerProviderResponse) SetTrustStorePinPassphraseProvider(v string) {
	o.TrustStorePinPassphraseProvider = &v
}

// GetEnabled returns the Enabled field value
func (o *FileBasedTrustManagerProviderResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *FileBasedTrustManagerProviderResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *FileBasedTrustManagerProviderResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetIncludeJVMDefaultIssuers returns the IncludeJVMDefaultIssuers field value if set, zero value otherwise.
func (o *FileBasedTrustManagerProviderResponse) GetIncludeJVMDefaultIssuers() bool {
	if o == nil || isNil(o.IncludeJVMDefaultIssuers) {
		var ret bool
		return ret
	}
	return *o.IncludeJVMDefaultIssuers
}

// GetIncludeJVMDefaultIssuersOk returns a tuple with the IncludeJVMDefaultIssuers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedTrustManagerProviderResponse) GetIncludeJVMDefaultIssuersOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeJVMDefaultIssuers) {
    return nil, false
	}
	return o.IncludeJVMDefaultIssuers, true
}

// HasIncludeJVMDefaultIssuers returns a boolean if a field has been set.
func (o *FileBasedTrustManagerProviderResponse) HasIncludeJVMDefaultIssuers() bool {
	if o != nil && !isNil(o.IncludeJVMDefaultIssuers) {
		return true
	}

	return false
}

// SetIncludeJVMDefaultIssuers gets a reference to the given bool and assigns it to the IncludeJVMDefaultIssuers field.
func (o *FileBasedTrustManagerProviderResponse) SetIncludeJVMDefaultIssuers(v bool) {
	o.IncludeJVMDefaultIssuers = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *FileBasedTrustManagerProviderResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedTrustManagerProviderResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *FileBasedTrustManagerProviderResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *FileBasedTrustManagerProviderResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

func (o FileBasedTrustManagerProviderResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["trustStoreFile"] = o.TrustStoreFile
	}
	if !isNil(o.TrustStoreType) {
		toSerialize["trustStoreType"] = o.TrustStoreType
	}
	if !isNil(o.TrustStorePin) {
		toSerialize["trustStorePin"] = o.TrustStorePin
	}
	if !isNil(o.TrustStorePinFile) {
		toSerialize["trustStorePinFile"] = o.TrustStorePinFile
	}
	if !isNil(o.TrustStorePinPassphraseProvider) {
		toSerialize["trustStorePinPassphraseProvider"] = o.TrustStorePinPassphraseProvider
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.IncludeJVMDefaultIssuers) {
		toSerialize["includeJVMDefaultIssuers"] = o.IncludeJVMDefaultIssuers
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableFileBasedTrustManagerProviderResponse struct {
	value *FileBasedTrustManagerProviderResponse
	isSet bool
}

func (v NullableFileBasedTrustManagerProviderResponse) Get() *FileBasedTrustManagerProviderResponse {
	return v.value
}

func (v *NullableFileBasedTrustManagerProviderResponse) Set(val *FileBasedTrustManagerProviderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFileBasedTrustManagerProviderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFileBasedTrustManagerProviderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileBasedTrustManagerProviderResponse(val *FileBasedTrustManagerProviderResponse) *NullableFileBasedTrustManagerProviderResponse {
	return &NullableFileBasedTrustManagerProviderResponse{value: val, isSet: true}
}

func (v NullableFileBasedTrustManagerProviderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileBasedTrustManagerProviderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


