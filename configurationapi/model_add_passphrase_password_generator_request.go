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

// checks if the AddPassphrasePasswordGeneratorRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddPassphrasePasswordGeneratorRequest{}

// AddPassphrasePasswordGeneratorRequest struct for AddPassphrasePasswordGeneratorRequest
type AddPassphrasePasswordGeneratorRequest struct {
	// Name of the new Password Generator
	GeneratorName string                                     `json:"generatorName"`
	Schemas       []EnumpassphrasePasswordGeneratorSchemaUrn `json:"schemas"`
	// The path to the dictionary file that will be used to obtain the words for use in generated passwords.
	DictionaryFile string `json:"dictionaryFile"`
	// The minimum number of characters that generated passwords will be required to have.
	MinimumPasswordCharacters *int32 `json:"minimumPasswordCharacters,omitempty"`
	// The minimum number of words that must be concatenated in the course of generating a password.
	MinimumPasswordWords *int32 `json:"minimumPasswordWords,omitempty"`
	// Indicates whether to capitalize each word used in the generated password.
	CapitalizeWords *bool `json:"capitalizeWords,omitempty"`
	// A description for this Password Generator
	Description *string `json:"description,omitempty"`
	// Indicates whether the Password Generator is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewAddPassphrasePasswordGeneratorRequest instantiates a new AddPassphrasePasswordGeneratorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPassphrasePasswordGeneratorRequest(generatorName string, schemas []EnumpassphrasePasswordGeneratorSchemaUrn, dictionaryFile string, enabled bool) *AddPassphrasePasswordGeneratorRequest {
	this := AddPassphrasePasswordGeneratorRequest{}
	this.GeneratorName = generatorName
	this.Schemas = schemas
	this.DictionaryFile = dictionaryFile
	this.Enabled = enabled
	return &this
}

// NewAddPassphrasePasswordGeneratorRequestWithDefaults instantiates a new AddPassphrasePasswordGeneratorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPassphrasePasswordGeneratorRequestWithDefaults() *AddPassphrasePasswordGeneratorRequest {
	this := AddPassphrasePasswordGeneratorRequest{}
	return &this
}

// GetGeneratorName returns the GeneratorName field value
func (o *AddPassphrasePasswordGeneratorRequest) GetGeneratorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GeneratorName
}

// GetGeneratorNameOk returns a tuple with the GeneratorName field value
// and a boolean to check if the value has been set.
func (o *AddPassphrasePasswordGeneratorRequest) GetGeneratorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GeneratorName, true
}

// SetGeneratorName sets field value
func (o *AddPassphrasePasswordGeneratorRequest) SetGeneratorName(v string) {
	o.GeneratorName = v
}

// GetSchemas returns the Schemas field value
func (o *AddPassphrasePasswordGeneratorRequest) GetSchemas() []EnumpassphrasePasswordGeneratorSchemaUrn {
	if o == nil {
		var ret []EnumpassphrasePasswordGeneratorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddPassphrasePasswordGeneratorRequest) GetSchemasOk() ([]EnumpassphrasePasswordGeneratorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddPassphrasePasswordGeneratorRequest) SetSchemas(v []EnumpassphrasePasswordGeneratorSchemaUrn) {
	o.Schemas = v
}

// GetDictionaryFile returns the DictionaryFile field value
func (o *AddPassphrasePasswordGeneratorRequest) GetDictionaryFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DictionaryFile
}

// GetDictionaryFileOk returns a tuple with the DictionaryFile field value
// and a boolean to check if the value has been set.
func (o *AddPassphrasePasswordGeneratorRequest) GetDictionaryFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DictionaryFile, true
}

// SetDictionaryFile sets field value
func (o *AddPassphrasePasswordGeneratorRequest) SetDictionaryFile(v string) {
	o.DictionaryFile = v
}

// GetMinimumPasswordCharacters returns the MinimumPasswordCharacters field value if set, zero value otherwise.
func (o *AddPassphrasePasswordGeneratorRequest) GetMinimumPasswordCharacters() int32 {
	if o == nil || IsNil(o.MinimumPasswordCharacters) {
		var ret int32
		return ret
	}
	return *o.MinimumPasswordCharacters
}

// GetMinimumPasswordCharactersOk returns a tuple with the MinimumPasswordCharacters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPassphrasePasswordGeneratorRequest) GetMinimumPasswordCharactersOk() (*int32, bool) {
	if o == nil || IsNil(o.MinimumPasswordCharacters) {
		return nil, false
	}
	return o.MinimumPasswordCharacters, true
}

// HasMinimumPasswordCharacters returns a boolean if a field has been set.
func (o *AddPassphrasePasswordGeneratorRequest) HasMinimumPasswordCharacters() bool {
	if o != nil && !IsNil(o.MinimumPasswordCharacters) {
		return true
	}

	return false
}

// SetMinimumPasswordCharacters gets a reference to the given int32 and assigns it to the MinimumPasswordCharacters field.
func (o *AddPassphrasePasswordGeneratorRequest) SetMinimumPasswordCharacters(v int32) {
	o.MinimumPasswordCharacters = &v
}

// GetMinimumPasswordWords returns the MinimumPasswordWords field value if set, zero value otherwise.
func (o *AddPassphrasePasswordGeneratorRequest) GetMinimumPasswordWords() int32 {
	if o == nil || IsNil(o.MinimumPasswordWords) {
		var ret int32
		return ret
	}
	return *o.MinimumPasswordWords
}

// GetMinimumPasswordWordsOk returns a tuple with the MinimumPasswordWords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPassphrasePasswordGeneratorRequest) GetMinimumPasswordWordsOk() (*int32, bool) {
	if o == nil || IsNil(o.MinimumPasswordWords) {
		return nil, false
	}
	return o.MinimumPasswordWords, true
}

// HasMinimumPasswordWords returns a boolean if a field has been set.
func (o *AddPassphrasePasswordGeneratorRequest) HasMinimumPasswordWords() bool {
	if o != nil && !IsNil(o.MinimumPasswordWords) {
		return true
	}

	return false
}

// SetMinimumPasswordWords gets a reference to the given int32 and assigns it to the MinimumPasswordWords field.
func (o *AddPassphrasePasswordGeneratorRequest) SetMinimumPasswordWords(v int32) {
	o.MinimumPasswordWords = &v
}

// GetCapitalizeWords returns the CapitalizeWords field value if set, zero value otherwise.
func (o *AddPassphrasePasswordGeneratorRequest) GetCapitalizeWords() bool {
	if o == nil || IsNil(o.CapitalizeWords) {
		var ret bool
		return ret
	}
	return *o.CapitalizeWords
}

// GetCapitalizeWordsOk returns a tuple with the CapitalizeWords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPassphrasePasswordGeneratorRequest) GetCapitalizeWordsOk() (*bool, bool) {
	if o == nil || IsNil(o.CapitalizeWords) {
		return nil, false
	}
	return o.CapitalizeWords, true
}

// HasCapitalizeWords returns a boolean if a field has been set.
func (o *AddPassphrasePasswordGeneratorRequest) HasCapitalizeWords() bool {
	if o != nil && !IsNil(o.CapitalizeWords) {
		return true
	}

	return false
}

// SetCapitalizeWords gets a reference to the given bool and assigns it to the CapitalizeWords field.
func (o *AddPassphrasePasswordGeneratorRequest) SetCapitalizeWords(v bool) {
	o.CapitalizeWords = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddPassphrasePasswordGeneratorRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPassphrasePasswordGeneratorRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddPassphrasePasswordGeneratorRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddPassphrasePasswordGeneratorRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddPassphrasePasswordGeneratorRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddPassphrasePasswordGeneratorRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddPassphrasePasswordGeneratorRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddPassphrasePasswordGeneratorRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddPassphrasePasswordGeneratorRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["generatorName"] = o.GeneratorName
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
	return toSerialize, nil
}

type NullableAddPassphrasePasswordGeneratorRequest struct {
	value *AddPassphrasePasswordGeneratorRequest
	isSet bool
}

func (v NullableAddPassphrasePasswordGeneratorRequest) Get() *AddPassphrasePasswordGeneratorRequest {
	return v.value
}

func (v *NullableAddPassphrasePasswordGeneratorRequest) Set(val *AddPassphrasePasswordGeneratorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPassphrasePasswordGeneratorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPassphrasePasswordGeneratorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPassphrasePasswordGeneratorRequest(val *AddPassphrasePasswordGeneratorRequest) *NullableAddPassphrasePasswordGeneratorRequest {
	return &NullableAddPassphrasePasswordGeneratorRequest{value: val, isSet: true}
}

func (v NullableAddPassphrasePasswordGeneratorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPassphrasePasswordGeneratorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
