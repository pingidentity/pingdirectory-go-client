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

// checks if the ChangelogPasswordEncryptionPluginResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangelogPasswordEncryptionPluginResponse{}

// ChangelogPasswordEncryptionPluginResponse struct for ChangelogPasswordEncryptionPluginResponse
type ChangelogPasswordEncryptionPluginResponse struct {
	Schemas []EnumchangelogPasswordEncryptionPluginSchemaUrn `json:"schemas"`
	// Name of the Plugin
	Id string `json:"id"`
	// A passphrase that may be used to generate the key for encrypting passwords stored in the changelog. The same passphrase also needs to be set (either through the \"changelog-password-decryption-key\" property or the \"changelog-password-decryption-key-passphrase-provider\" property) in the Global Sync Configuration in the Data Sync Server.
	ChangelogPasswordEncryptionKey *string `json:"changelogPasswordEncryptionKey,omitempty"`
	// A passphrase provider that may be used to obtain the passphrase that will be used to generate the key for encrypting passwords stored in the changelog. The same passphrase also needs to be set (either through the \"changelog-password-decryption-key\" property or the \"changelog-password-decryption-key-passphrase-provider\" property) in the Global Sync Configuration in the Data Sync Server.
	ChangelogPasswordEncryptionKeyPassphraseProvider *string                    `json:"changelogPasswordEncryptionKeyPassphraseProvider,omitempty"`
	PluginType                                       []EnumpluginPluginTypeProp `json:"pluginType"`
	// A description for this Plugin
	Description *string `json:"description,omitempty"`
	// Indicates whether the plug-in is enabled for use.
	Enabled bool `json:"enabled"`
	// Indicates whether the plug-in should be invoked for internal operations.
	InvokeForInternalOperations                   *bool                                              `json:"invokeForInternalOperations,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewChangelogPasswordEncryptionPluginResponse instantiates a new ChangelogPasswordEncryptionPluginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangelogPasswordEncryptionPluginResponse(schemas []EnumchangelogPasswordEncryptionPluginSchemaUrn, id string, pluginType []EnumpluginPluginTypeProp, enabled bool) *ChangelogPasswordEncryptionPluginResponse {
	this := ChangelogPasswordEncryptionPluginResponse{}
	this.Schemas = schemas
	this.Id = id
	this.PluginType = pluginType
	this.Enabled = enabled
	return &this
}

// NewChangelogPasswordEncryptionPluginResponseWithDefaults instantiates a new ChangelogPasswordEncryptionPluginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangelogPasswordEncryptionPluginResponseWithDefaults() *ChangelogPasswordEncryptionPluginResponse {
	this := ChangelogPasswordEncryptionPluginResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *ChangelogPasswordEncryptionPluginResponse) GetSchemas() []EnumchangelogPasswordEncryptionPluginSchemaUrn {
	if o == nil {
		var ret []EnumchangelogPasswordEncryptionPluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ChangelogPasswordEncryptionPluginResponse) GetSchemasOk() ([]EnumchangelogPasswordEncryptionPluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ChangelogPasswordEncryptionPluginResponse) SetSchemas(v []EnumchangelogPasswordEncryptionPluginSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *ChangelogPasswordEncryptionPluginResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ChangelogPasswordEncryptionPluginResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ChangelogPasswordEncryptionPluginResponse) SetId(v string) {
	o.Id = v
}

// GetChangelogPasswordEncryptionKey returns the ChangelogPasswordEncryptionKey field value if set, zero value otherwise.
func (o *ChangelogPasswordEncryptionPluginResponse) GetChangelogPasswordEncryptionKey() string {
	if o == nil || IsNil(o.ChangelogPasswordEncryptionKey) {
		var ret string
		return ret
	}
	return *o.ChangelogPasswordEncryptionKey
}

// GetChangelogPasswordEncryptionKeyOk returns a tuple with the ChangelogPasswordEncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangelogPasswordEncryptionPluginResponse) GetChangelogPasswordEncryptionKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ChangelogPasswordEncryptionKey) {
		return nil, false
	}
	return o.ChangelogPasswordEncryptionKey, true
}

// HasChangelogPasswordEncryptionKey returns a boolean if a field has been set.
func (o *ChangelogPasswordEncryptionPluginResponse) HasChangelogPasswordEncryptionKey() bool {
	if o != nil && !IsNil(o.ChangelogPasswordEncryptionKey) {
		return true
	}

	return false
}

// SetChangelogPasswordEncryptionKey gets a reference to the given string and assigns it to the ChangelogPasswordEncryptionKey field.
func (o *ChangelogPasswordEncryptionPluginResponse) SetChangelogPasswordEncryptionKey(v string) {
	o.ChangelogPasswordEncryptionKey = &v
}

// GetChangelogPasswordEncryptionKeyPassphraseProvider returns the ChangelogPasswordEncryptionKeyPassphraseProvider field value if set, zero value otherwise.
func (o *ChangelogPasswordEncryptionPluginResponse) GetChangelogPasswordEncryptionKeyPassphraseProvider() string {
	if o == nil || IsNil(o.ChangelogPasswordEncryptionKeyPassphraseProvider) {
		var ret string
		return ret
	}
	return *o.ChangelogPasswordEncryptionKeyPassphraseProvider
}

// GetChangelogPasswordEncryptionKeyPassphraseProviderOk returns a tuple with the ChangelogPasswordEncryptionKeyPassphraseProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangelogPasswordEncryptionPluginResponse) GetChangelogPasswordEncryptionKeyPassphraseProviderOk() (*string, bool) {
	if o == nil || IsNil(o.ChangelogPasswordEncryptionKeyPassphraseProvider) {
		return nil, false
	}
	return o.ChangelogPasswordEncryptionKeyPassphraseProvider, true
}

// HasChangelogPasswordEncryptionKeyPassphraseProvider returns a boolean if a field has been set.
func (o *ChangelogPasswordEncryptionPluginResponse) HasChangelogPasswordEncryptionKeyPassphraseProvider() bool {
	if o != nil && !IsNil(o.ChangelogPasswordEncryptionKeyPassphraseProvider) {
		return true
	}

	return false
}

// SetChangelogPasswordEncryptionKeyPassphraseProvider gets a reference to the given string and assigns it to the ChangelogPasswordEncryptionKeyPassphraseProvider field.
func (o *ChangelogPasswordEncryptionPluginResponse) SetChangelogPasswordEncryptionKeyPassphraseProvider(v string) {
	o.ChangelogPasswordEncryptionKeyPassphraseProvider = &v
}

// GetPluginType returns the PluginType field value
func (o *ChangelogPasswordEncryptionPluginResponse) GetPluginType() []EnumpluginPluginTypeProp {
	if o == nil {
		var ret []EnumpluginPluginTypeProp
		return ret
	}

	return o.PluginType
}

// GetPluginTypeOk returns a tuple with the PluginType field value
// and a boolean to check if the value has been set.
func (o *ChangelogPasswordEncryptionPluginResponse) GetPluginTypeOk() ([]EnumpluginPluginTypeProp, bool) {
	if o == nil {
		return nil, false
	}
	return o.PluginType, true
}

// SetPluginType sets field value
func (o *ChangelogPasswordEncryptionPluginResponse) SetPluginType(v []EnumpluginPluginTypeProp) {
	o.PluginType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ChangelogPasswordEncryptionPluginResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangelogPasswordEncryptionPluginResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ChangelogPasswordEncryptionPluginResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ChangelogPasswordEncryptionPluginResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *ChangelogPasswordEncryptionPluginResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ChangelogPasswordEncryptionPluginResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ChangelogPasswordEncryptionPluginResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetInvokeForInternalOperations returns the InvokeForInternalOperations field value if set, zero value otherwise.
func (o *ChangelogPasswordEncryptionPluginResponse) GetInvokeForInternalOperations() bool {
	if o == nil || IsNil(o.InvokeForInternalOperations) {
		var ret bool
		return ret
	}
	return *o.InvokeForInternalOperations
}

// GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangelogPasswordEncryptionPluginResponse) GetInvokeForInternalOperationsOk() (*bool, bool) {
	if o == nil || IsNil(o.InvokeForInternalOperations) {
		return nil, false
	}
	return o.InvokeForInternalOperations, true
}

// HasInvokeForInternalOperations returns a boolean if a field has been set.
func (o *ChangelogPasswordEncryptionPluginResponse) HasInvokeForInternalOperations() bool {
	if o != nil && !IsNil(o.InvokeForInternalOperations) {
		return true
	}

	return false
}

// SetInvokeForInternalOperations gets a reference to the given bool and assigns it to the InvokeForInternalOperations field.
func (o *ChangelogPasswordEncryptionPluginResponse) SetInvokeForInternalOperations(v bool) {
	o.InvokeForInternalOperations = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ChangelogPasswordEncryptionPluginResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangelogPasswordEncryptionPluginResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ChangelogPasswordEncryptionPluginResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ChangelogPasswordEncryptionPluginResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ChangelogPasswordEncryptionPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangelogPasswordEncryptionPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ChangelogPasswordEncryptionPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ChangelogPasswordEncryptionPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o ChangelogPasswordEncryptionPluginResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangelogPasswordEncryptionPluginResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
	if !IsNil(o.ChangelogPasswordEncryptionKey) {
		toSerialize["changelogPasswordEncryptionKey"] = o.ChangelogPasswordEncryptionKey
	}
	if !IsNil(o.ChangelogPasswordEncryptionKeyPassphraseProvider) {
		toSerialize["changelogPasswordEncryptionKeyPassphraseProvider"] = o.ChangelogPasswordEncryptionKeyPassphraseProvider
	}
	toSerialize["pluginType"] = o.PluginType
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.InvokeForInternalOperations) {
		toSerialize["invokeForInternalOperations"] = o.InvokeForInternalOperations
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableChangelogPasswordEncryptionPluginResponse struct {
	value *ChangelogPasswordEncryptionPluginResponse
	isSet bool
}

func (v NullableChangelogPasswordEncryptionPluginResponse) Get() *ChangelogPasswordEncryptionPluginResponse {
	return v.value
}

func (v *NullableChangelogPasswordEncryptionPluginResponse) Set(val *ChangelogPasswordEncryptionPluginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChangelogPasswordEncryptionPluginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChangelogPasswordEncryptionPluginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangelogPasswordEncryptionPluginResponse(val *ChangelogPasswordEncryptionPluginResponse) *NullableChangelogPasswordEncryptionPluginResponse {
	return &NullableChangelogPasswordEncryptionPluginResponse{value: val, isSet: true}
}

func (v NullableChangelogPasswordEncryptionPluginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangelogPasswordEncryptionPluginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
