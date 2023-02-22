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

// AzureKeyVaultPassphraseProviderResponse struct for AzureKeyVaultPassphraseProviderResponse
type AzureKeyVaultPassphraseProviderResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Passphrase Provider
	Id      string                                         `json:"id"`
	Schemas []EnumazureKeyVaultPassphraseProviderSchemaUrn `json:"schemas"`
	// The URI that identifies the Azure Key Vault from which the secret is to be retrieved.
	KeyVaultURI string `json:"keyVaultURI"`
	// The mechanism used to authenticate to the Azure service.
	AzureAuthenticationMethod string `json:"azureAuthenticationMethod"`
	// The name of the secret to retrieve.
	SecretName string `json:"secretName"`
	// The maximum length of time that the passphrase provider may cache the passphrase that has been read from Azure Key Vault. A value of zero seconds indicates that the provider should always attempt to read the passphrase from the Azure service.
	MaxCacheDuration *string `json:"maxCacheDuration,omitempty"`
	// A description for this Passphrase Provider
	Description *string `json:"description,omitempty"`
	// Indicates whether this Passphrase Provider is enabled for use in the server.
	Enabled bool `json:"enabled"`
}

// NewAzureKeyVaultPassphraseProviderResponse instantiates a new AzureKeyVaultPassphraseProviderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureKeyVaultPassphraseProviderResponse(id string, schemas []EnumazureKeyVaultPassphraseProviderSchemaUrn, keyVaultURI string, azureAuthenticationMethod string, secretName string, enabled bool) *AzureKeyVaultPassphraseProviderResponse {
	this := AzureKeyVaultPassphraseProviderResponse{}
	this.Id = id
	this.Schemas = schemas
	this.KeyVaultURI = keyVaultURI
	this.AzureAuthenticationMethod = azureAuthenticationMethod
	this.SecretName = secretName
	this.Enabled = enabled
	return &this
}

// NewAzureKeyVaultPassphraseProviderResponseWithDefaults instantiates a new AzureKeyVaultPassphraseProviderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureKeyVaultPassphraseProviderResponseWithDefaults() *AzureKeyVaultPassphraseProviderResponse {
	this := AzureKeyVaultPassphraseProviderResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AzureKeyVaultPassphraseProviderResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPassphraseProviderResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AzureKeyVaultPassphraseProviderResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *AzureKeyVaultPassphraseProviderResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *AzureKeyVaultPassphraseProviderResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPassphraseProviderResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *AzureKeyVaultPassphraseProviderResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *AzureKeyVaultPassphraseProviderResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *AzureKeyVaultPassphraseProviderResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPassphraseProviderResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AzureKeyVaultPassphraseProviderResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *AzureKeyVaultPassphraseProviderResponse) GetSchemas() []EnumazureKeyVaultPassphraseProviderSchemaUrn {
	if o == nil {
		var ret []EnumazureKeyVaultPassphraseProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPassphraseProviderResponse) GetSchemasOk() ([]EnumazureKeyVaultPassphraseProviderSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AzureKeyVaultPassphraseProviderResponse) SetSchemas(v []EnumazureKeyVaultPassphraseProviderSchemaUrn) {
	o.Schemas = v
}

// GetKeyVaultURI returns the KeyVaultURI field value
func (o *AzureKeyVaultPassphraseProviderResponse) GetKeyVaultURI() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyVaultURI
}

// GetKeyVaultURIOk returns a tuple with the KeyVaultURI field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPassphraseProviderResponse) GetKeyVaultURIOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyVaultURI, true
}

// SetKeyVaultURI sets field value
func (o *AzureKeyVaultPassphraseProviderResponse) SetKeyVaultURI(v string) {
	o.KeyVaultURI = v
}

// GetAzureAuthenticationMethod returns the AzureAuthenticationMethod field value
func (o *AzureKeyVaultPassphraseProviderResponse) GetAzureAuthenticationMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AzureAuthenticationMethod
}

// GetAzureAuthenticationMethodOk returns a tuple with the AzureAuthenticationMethod field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPassphraseProviderResponse) GetAzureAuthenticationMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureAuthenticationMethod, true
}

// SetAzureAuthenticationMethod sets field value
func (o *AzureKeyVaultPassphraseProviderResponse) SetAzureAuthenticationMethod(v string) {
	o.AzureAuthenticationMethod = v
}

// GetSecretName returns the SecretName field value
func (o *AzureKeyVaultPassphraseProviderResponse) GetSecretName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretName
}

// GetSecretNameOk returns a tuple with the SecretName field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPassphraseProviderResponse) GetSecretNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretName, true
}

// SetSecretName sets field value
func (o *AzureKeyVaultPassphraseProviderResponse) SetSecretName(v string) {
	o.SecretName = v
}

// GetMaxCacheDuration returns the MaxCacheDuration field value if set, zero value otherwise.
func (o *AzureKeyVaultPassphraseProviderResponse) GetMaxCacheDuration() string {
	if o == nil || isNil(o.MaxCacheDuration) {
		var ret string
		return ret
	}
	return *o.MaxCacheDuration
}

// GetMaxCacheDurationOk returns a tuple with the MaxCacheDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPassphraseProviderResponse) GetMaxCacheDurationOk() (*string, bool) {
	if o == nil || isNil(o.MaxCacheDuration) {
		return nil, false
	}
	return o.MaxCacheDuration, true
}

// HasMaxCacheDuration returns a boolean if a field has been set.
func (o *AzureKeyVaultPassphraseProviderResponse) HasMaxCacheDuration() bool {
	if o != nil && !isNil(o.MaxCacheDuration) {
		return true
	}

	return false
}

// SetMaxCacheDuration gets a reference to the given string and assigns it to the MaxCacheDuration field.
func (o *AzureKeyVaultPassphraseProviderResponse) SetMaxCacheDuration(v string) {
	o.MaxCacheDuration = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AzureKeyVaultPassphraseProviderResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPassphraseProviderResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AzureKeyVaultPassphraseProviderResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AzureKeyVaultPassphraseProviderResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AzureKeyVaultPassphraseProviderResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPassphraseProviderResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AzureKeyVaultPassphraseProviderResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AzureKeyVaultPassphraseProviderResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
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
	if !isNil(o.MaxCacheDuration) {
		toSerialize["maxCacheDuration"] = o.MaxCacheDuration
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableAzureKeyVaultPassphraseProviderResponse struct {
	value *AzureKeyVaultPassphraseProviderResponse
	isSet bool
}

func (v NullableAzureKeyVaultPassphraseProviderResponse) Get() *AzureKeyVaultPassphraseProviderResponse {
	return v.value
}

func (v *NullableAzureKeyVaultPassphraseProviderResponse) Set(val *AzureKeyVaultPassphraseProviderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureKeyVaultPassphraseProviderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureKeyVaultPassphraseProviderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureKeyVaultPassphraseProviderResponse(val *AzureKeyVaultPassphraseProviderResponse) *NullableAzureKeyVaultPassphraseProviderResponse {
	return &NullableAzureKeyVaultPassphraseProviderResponse{value: val, isSet: true}
}

func (v NullableAzureKeyVaultPassphraseProviderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureKeyVaultPassphraseProviderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}