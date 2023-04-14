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

// checks if the PassphrasePasswordGeneratorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PassphrasePasswordGeneratorResponse{}

// PassphrasePasswordGeneratorResponse struct for PassphrasePasswordGeneratorResponse
type PassphrasePasswordGeneratorResponse struct {
	// Name of the Password Generator
	Id      string                                     `json:"id"`
	Schemas []EnumpassphrasePasswordGeneratorSchemaUrn `json:"schemas"`
	// The path to the dictionary file that will be used to obtain the words for use in generated passwords.
	DictionaryFile string `json:"dictionaryFile"`
	// The minimum number of characters that generated passwords will be required to have.
	MinimumPasswordCharacters *int64 `json:"minimumPasswordCharacters,omitempty"`
	// The minimum number of words that must be concatenated in the course of generating a password.
	MinimumPasswordWords *int64 `json:"minimumPasswordWords,omitempty"`
	// Indicates whether to capitalize each word used in the generated password.
	CapitalizeWords *bool `json:"capitalizeWords,omitempty"`
	// A description for this Password Generator
	Description *string `json:"description,omitempty"`
	// Indicates whether the Password Generator is enabled for use.
	Enabled                                       bool                                               `json:"enabled"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewPassphrasePasswordGeneratorResponse instantiates a new PassphrasePasswordGeneratorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPassphrasePasswordGeneratorResponse(id string, schemas []EnumpassphrasePasswordGeneratorSchemaUrn, dictionaryFile string, enabled bool) *PassphrasePasswordGeneratorResponse {
	this := PassphrasePasswordGeneratorResponse{}
	this.Id = id
	this.Schemas = schemas
	this.DictionaryFile = dictionaryFile
	this.Enabled = enabled
	return &this
}

// NewPassphrasePasswordGeneratorResponseWithDefaults instantiates a new PassphrasePasswordGeneratorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPassphrasePasswordGeneratorResponseWithDefaults() *PassphrasePasswordGeneratorResponse {
	this := PassphrasePasswordGeneratorResponse{}
	return &this
}

// GetId returns the Id field value
func (o *PassphrasePasswordGeneratorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PassphrasePasswordGeneratorResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PassphrasePasswordGeneratorResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *PassphrasePasswordGeneratorResponse) GetSchemas() []EnumpassphrasePasswordGeneratorSchemaUrn {
	if o == nil {
		var ret []EnumpassphrasePasswordGeneratorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *PassphrasePasswordGeneratorResponse) GetSchemasOk() ([]EnumpassphrasePasswordGeneratorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *PassphrasePasswordGeneratorResponse) SetSchemas(v []EnumpassphrasePasswordGeneratorSchemaUrn) {
	o.Schemas = v
}

// GetDictionaryFile returns the DictionaryFile field value
func (o *PassphrasePasswordGeneratorResponse) GetDictionaryFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DictionaryFile
}

// GetDictionaryFileOk returns a tuple with the DictionaryFile field value
// and a boolean to check if the value has been set.
func (o *PassphrasePasswordGeneratorResponse) GetDictionaryFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DictionaryFile, true
}

// SetDictionaryFile sets field value
func (o *PassphrasePasswordGeneratorResponse) SetDictionaryFile(v string) {
	o.DictionaryFile = v
}

// GetMinimumPasswordCharacters returns the MinimumPasswordCharacters field value if set, zero value otherwise.
func (o *PassphrasePasswordGeneratorResponse) GetMinimumPasswordCharacters() int64 {
	if o == nil || IsNil(o.MinimumPasswordCharacters) {
		var ret int64
		return ret
	}
	return *o.MinimumPasswordCharacters
}

// GetMinimumPasswordCharactersOk returns a tuple with the MinimumPasswordCharacters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PassphrasePasswordGeneratorResponse) GetMinimumPasswordCharactersOk() (*int64, bool) {
	if o == nil || IsNil(o.MinimumPasswordCharacters) {
		return nil, false
	}
	return o.MinimumPasswordCharacters, true
}

// HasMinimumPasswordCharacters returns a boolean if a field has been set.
func (o *PassphrasePasswordGeneratorResponse) HasMinimumPasswordCharacters() bool {
	if o != nil && !IsNil(o.MinimumPasswordCharacters) {
		return true
	}

	return false
}

// SetMinimumPasswordCharacters gets a reference to the given int64 and assigns it to the MinimumPasswordCharacters field.
func (o *PassphrasePasswordGeneratorResponse) SetMinimumPasswordCharacters(v int64) {
	o.MinimumPasswordCharacters = &v
}

// GetMinimumPasswordWords returns the MinimumPasswordWords field value if set, zero value otherwise.
func (o *PassphrasePasswordGeneratorResponse) GetMinimumPasswordWords() int64 {
	if o == nil || IsNil(o.MinimumPasswordWords) {
		var ret int64
		return ret
	}
	return *o.MinimumPasswordWords
}

// GetMinimumPasswordWordsOk returns a tuple with the MinimumPasswordWords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PassphrasePasswordGeneratorResponse) GetMinimumPasswordWordsOk() (*int64, bool) {
	if o == nil || IsNil(o.MinimumPasswordWords) {
		return nil, false
	}
	return o.MinimumPasswordWords, true
}

// HasMinimumPasswordWords returns a boolean if a field has been set.
func (o *PassphrasePasswordGeneratorResponse) HasMinimumPasswordWords() bool {
	if o != nil && !IsNil(o.MinimumPasswordWords) {
		return true
	}

	return false
}

// SetMinimumPasswordWords gets a reference to the given int64 and assigns it to the MinimumPasswordWords field.
func (o *PassphrasePasswordGeneratorResponse) SetMinimumPasswordWords(v int64) {
	o.MinimumPasswordWords = &v
}

// GetCapitalizeWords returns the CapitalizeWords field value if set, zero value otherwise.
func (o *PassphrasePasswordGeneratorResponse) GetCapitalizeWords() bool {
	if o == nil || IsNil(o.CapitalizeWords) {
		var ret bool
		return ret
	}
	return *o.CapitalizeWords
}

// GetCapitalizeWordsOk returns a tuple with the CapitalizeWords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PassphrasePasswordGeneratorResponse) GetCapitalizeWordsOk() (*bool, bool) {
	if o == nil || IsNil(o.CapitalizeWords) {
		return nil, false
	}
	return o.CapitalizeWords, true
}

// HasCapitalizeWords returns a boolean if a field has been set.
func (o *PassphrasePasswordGeneratorResponse) HasCapitalizeWords() bool {
	if o != nil && !IsNil(o.CapitalizeWords) {
		return true
	}

	return false
}

// SetCapitalizeWords gets a reference to the given bool and assigns it to the CapitalizeWords field.
func (o *PassphrasePasswordGeneratorResponse) SetCapitalizeWords(v bool) {
	o.CapitalizeWords = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PassphrasePasswordGeneratorResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PassphrasePasswordGeneratorResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PassphrasePasswordGeneratorResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PassphrasePasswordGeneratorResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *PassphrasePasswordGeneratorResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *PassphrasePasswordGeneratorResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *PassphrasePasswordGeneratorResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *PassphrasePasswordGeneratorResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PassphrasePasswordGeneratorResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *PassphrasePasswordGeneratorResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *PassphrasePasswordGeneratorResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *PassphrasePasswordGeneratorResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PassphrasePasswordGeneratorResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *PassphrasePasswordGeneratorResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *PassphrasePasswordGeneratorResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o PassphrasePasswordGeneratorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PassphrasePasswordGeneratorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["schemas"] = o.Schemas
	toSerialize["dictionaryFile"] = o.DictionaryFile
	if !IsNil(o.MinimumPasswordCharacters) {
		toSerialize["minimumPasswordCharacters"] = o.MinimumPasswordCharacters
	}
	if !IsNil(o.MinimumPasswordWords) {
		toSerialize["minimumPasswordWords"] = o.MinimumPasswordWords
	}
	if !IsNil(o.CapitalizeWords) {
		toSerialize["capitalizeWords"] = o.CapitalizeWords
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
	return toSerialize, nil
}

type NullablePassphrasePasswordGeneratorResponse struct {
	value *PassphrasePasswordGeneratorResponse
	isSet bool
}

func (v NullablePassphrasePasswordGeneratorResponse) Get() *PassphrasePasswordGeneratorResponse {
	return v.value
}

func (v *NullablePassphrasePasswordGeneratorResponse) Set(val *PassphrasePasswordGeneratorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePassphrasePasswordGeneratorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePassphrasePasswordGeneratorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePassphrasePasswordGeneratorResponse(val *PassphrasePasswordGeneratorResponse) *NullablePassphrasePasswordGeneratorResponse {
	return &NullablePassphrasePasswordGeneratorResponse{value: val, isSet: true}
}

func (v NullablePassphrasePasswordGeneratorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePassphrasePasswordGeneratorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
