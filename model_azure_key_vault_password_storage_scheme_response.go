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

// AzureKeyVaultPasswordStorageSchemeResponse struct for AzureKeyVaultPasswordStorageSchemeResponse
type AzureKeyVaultPasswordStorageSchemeResponse struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Password Storage Scheme
	Id string `json:"id"`
	Schemas []EnumazureKeyVaultPasswordStorageSchemeSchemaUrn `json:"schemas"`
	// The URI that identifies the Azure Key Vault from which the secret is to be retrieved.
	KeyVaultURI string `json:"keyVaultURI"`
	// The mechanism used to authenticate to the Azure service.
	AzureAuthenticationMethod string `json:"azureAuthenticationMethod"`
	// A description for this Password Storage Scheme
	Description *string `json:"description,omitempty"`
	// Indicates whether the Password Storage Scheme is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewAzureKeyVaultPasswordStorageSchemeResponse instantiates a new AzureKeyVaultPasswordStorageSchemeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureKeyVaultPasswordStorageSchemeResponse(id string, schemas []EnumazureKeyVaultPasswordStorageSchemeSchemaUrn, keyVaultURI string, azureAuthenticationMethod string, enabled bool) *AzureKeyVaultPasswordStorageSchemeResponse {
	this := AzureKeyVaultPasswordStorageSchemeResponse{}
	this.Id = id
	this.Schemas = schemas
	this.KeyVaultURI = keyVaultURI
	this.AzureAuthenticationMethod = azureAuthenticationMethod
	this.Enabled = enabled
	return &this
}

// NewAzureKeyVaultPasswordStorageSchemeResponseWithDefaults instantiates a new AzureKeyVaultPasswordStorageSchemeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureKeyVaultPasswordStorageSchemeResponseWithDefaults() *AzureKeyVaultPasswordStorageSchemeResponse {
	this := AzureKeyVaultPasswordStorageSchemeResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AzureKeyVaultPasswordStorageSchemeResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPasswordStorageSchemeResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AzureKeyVaultPasswordStorageSchemeResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *AzureKeyVaultPasswordStorageSchemeResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *AzureKeyVaultPasswordStorageSchemeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPasswordStorageSchemeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
    return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *AzureKeyVaultPasswordStorageSchemeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *AzureKeyVaultPasswordStorageSchemeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *AzureKeyVaultPasswordStorageSchemeResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPasswordStorageSchemeResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AzureKeyVaultPasswordStorageSchemeResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *AzureKeyVaultPasswordStorageSchemeResponse) GetSchemas() []EnumazureKeyVaultPasswordStorageSchemeSchemaUrn {
	if o == nil {
		var ret []EnumazureKeyVaultPasswordStorageSchemeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPasswordStorageSchemeResponse) GetSchemasOk() ([]EnumazureKeyVaultPasswordStorageSchemeSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AzureKeyVaultPasswordStorageSchemeResponse) SetSchemas(v []EnumazureKeyVaultPasswordStorageSchemeSchemaUrn) {
	o.Schemas = v
}

// GetKeyVaultURI returns the KeyVaultURI field value
func (o *AzureKeyVaultPasswordStorageSchemeResponse) GetKeyVaultURI() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyVaultURI
}

// GetKeyVaultURIOk returns a tuple with the KeyVaultURI field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPasswordStorageSchemeResponse) GetKeyVaultURIOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.KeyVaultURI, true
}

// SetKeyVaultURI sets field value
func (o *AzureKeyVaultPasswordStorageSchemeResponse) SetKeyVaultURI(v string) {
	o.KeyVaultURI = v
}

// GetAzureAuthenticationMethod returns the AzureAuthenticationMethod field value
func (o *AzureKeyVaultPasswordStorageSchemeResponse) GetAzureAuthenticationMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AzureAuthenticationMethod
}

// GetAzureAuthenticationMethodOk returns a tuple with the AzureAuthenticationMethod field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPasswordStorageSchemeResponse) GetAzureAuthenticationMethodOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AzureAuthenticationMethod, true
}

// SetAzureAuthenticationMethod sets field value
func (o *AzureKeyVaultPasswordStorageSchemeResponse) SetAzureAuthenticationMethod(v string) {
	o.AzureAuthenticationMethod = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AzureKeyVaultPasswordStorageSchemeResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPasswordStorageSchemeResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AzureKeyVaultPasswordStorageSchemeResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AzureKeyVaultPasswordStorageSchemeResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AzureKeyVaultPasswordStorageSchemeResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPasswordStorageSchemeResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AzureKeyVaultPasswordStorageSchemeResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AzureKeyVaultPasswordStorageSchemeResponse) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableAzureKeyVaultPasswordStorageSchemeResponse struct {
	value *AzureKeyVaultPasswordStorageSchemeResponse
	isSet bool
}

func (v NullableAzureKeyVaultPasswordStorageSchemeResponse) Get() *AzureKeyVaultPasswordStorageSchemeResponse {
	return v.value
}

func (v *NullableAzureKeyVaultPasswordStorageSchemeResponse) Set(val *AzureKeyVaultPasswordStorageSchemeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureKeyVaultPasswordStorageSchemeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureKeyVaultPasswordStorageSchemeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureKeyVaultPasswordStorageSchemeResponse(val *AzureKeyVaultPasswordStorageSchemeResponse) *NullableAzureKeyVaultPasswordStorageSchemeResponse {
	return &NullableAzureKeyVaultPasswordStorageSchemeResponse{value: val, isSet: true}
}

func (v NullableAzureKeyVaultPasswordStorageSchemeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureKeyVaultPasswordStorageSchemeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


