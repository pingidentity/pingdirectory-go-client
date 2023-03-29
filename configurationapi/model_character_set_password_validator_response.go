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

// checks if the CharacterSetPasswordValidatorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CharacterSetPasswordValidatorResponse{}

// CharacterSetPasswordValidatorResponse struct for CharacterSetPasswordValidatorResponse
type CharacterSetPasswordValidatorResponse struct {
	// Name of the Password Validator
	Id      string                                       `json:"id"`
	Schemas []EnumcharacterSetPasswordValidatorSchemaUrn `json:"schemas"`
	// Specifies a character set containing characters that a password may contain and a value indicating the minimum number of characters required from that set.
	CharacterSet []string `json:"characterSet"`
	// Indicates whether this password validator allows passwords to contain characters outside of any of the user-defined character sets.
	AllowUnclassifiedCharacters bool `json:"allowUnclassifiedCharacters"`
	// Specifies the minimum number of character sets that must be represented in a proposed password.
	MinimumRequiredCharacterSets *int32 `json:"minimumRequiredCharacterSets,omitempty"`
	// A description for this Password Validator
	Description *string `json:"description,omitempty"`
	// Indicates whether the password validator is enabled for use.
	Enabled bool `json:"enabled"`
	// Specifies a message that can be used to describe the requirements imposed by this password validator to end users. If a value is provided for this property, then it will override any description that may have otherwise been generated by the validator.
	ValidatorRequirementDescription *string `json:"validatorRequirementDescription,omitempty"`
	// Specifies a message that may be provided to the end user in the event that a proposed password is rejected by this validator. If a value is provided for this property, then it will override any failure message that may have otherwise been generated by the validator.
	ValidatorFailureMessage                       *string                                            `json:"validatorFailureMessage,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewCharacterSetPasswordValidatorResponse instantiates a new CharacterSetPasswordValidatorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCharacterSetPasswordValidatorResponse(id string, schemas []EnumcharacterSetPasswordValidatorSchemaUrn, characterSet []string, allowUnclassifiedCharacters bool, enabled bool) *CharacterSetPasswordValidatorResponse {
	this := CharacterSetPasswordValidatorResponse{}
	this.Id = id
	this.Schemas = schemas
	this.CharacterSet = characterSet
	this.AllowUnclassifiedCharacters = allowUnclassifiedCharacters
	this.Enabled = enabled
	return &this
}

// NewCharacterSetPasswordValidatorResponseWithDefaults instantiates a new CharacterSetPasswordValidatorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCharacterSetPasswordValidatorResponseWithDefaults() *CharacterSetPasswordValidatorResponse {
	this := CharacterSetPasswordValidatorResponse{}
	return &this
}

// GetId returns the Id field value
func (o *CharacterSetPasswordValidatorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CharacterSetPasswordValidatorResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CharacterSetPasswordValidatorResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *CharacterSetPasswordValidatorResponse) GetSchemas() []EnumcharacterSetPasswordValidatorSchemaUrn {
	if o == nil {
		var ret []EnumcharacterSetPasswordValidatorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *CharacterSetPasswordValidatorResponse) GetSchemasOk() ([]EnumcharacterSetPasswordValidatorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *CharacterSetPasswordValidatorResponse) SetSchemas(v []EnumcharacterSetPasswordValidatorSchemaUrn) {
	o.Schemas = v
}

// GetCharacterSet returns the CharacterSet field value
func (o *CharacterSetPasswordValidatorResponse) GetCharacterSet() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CharacterSet
}

// GetCharacterSetOk returns a tuple with the CharacterSet field value
// and a boolean to check if the value has been set.
func (o *CharacterSetPasswordValidatorResponse) GetCharacterSetOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CharacterSet, true
}

// SetCharacterSet sets field value
func (o *CharacterSetPasswordValidatorResponse) SetCharacterSet(v []string) {
	o.CharacterSet = v
}

// GetAllowUnclassifiedCharacters returns the AllowUnclassifiedCharacters field value
func (o *CharacterSetPasswordValidatorResponse) GetAllowUnclassifiedCharacters() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowUnclassifiedCharacters
}

// GetAllowUnclassifiedCharactersOk returns a tuple with the AllowUnclassifiedCharacters field value
// and a boolean to check if the value has been set.
func (o *CharacterSetPasswordValidatorResponse) GetAllowUnclassifiedCharactersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowUnclassifiedCharacters, true
}

// SetAllowUnclassifiedCharacters sets field value
func (o *CharacterSetPasswordValidatorResponse) SetAllowUnclassifiedCharacters(v bool) {
	o.AllowUnclassifiedCharacters = v
}

// GetMinimumRequiredCharacterSets returns the MinimumRequiredCharacterSets field value if set, zero value otherwise.
func (o *CharacterSetPasswordValidatorResponse) GetMinimumRequiredCharacterSets() int32 {
	if o == nil || IsNil(o.MinimumRequiredCharacterSets) {
		var ret int32
		return ret
	}
	return *o.MinimumRequiredCharacterSets
}

// GetMinimumRequiredCharacterSetsOk returns a tuple with the MinimumRequiredCharacterSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CharacterSetPasswordValidatorResponse) GetMinimumRequiredCharacterSetsOk() (*int32, bool) {
	if o == nil || IsNil(o.MinimumRequiredCharacterSets) {
		return nil, false
	}
	return o.MinimumRequiredCharacterSets, true
}

// HasMinimumRequiredCharacterSets returns a boolean if a field has been set.
func (o *CharacterSetPasswordValidatorResponse) HasMinimumRequiredCharacterSets() bool {
	if o != nil && !IsNil(o.MinimumRequiredCharacterSets) {
		return true
	}

	return false
}

// SetMinimumRequiredCharacterSets gets a reference to the given int32 and assigns it to the MinimumRequiredCharacterSets field.
func (o *CharacterSetPasswordValidatorResponse) SetMinimumRequiredCharacterSets(v int32) {
	o.MinimumRequiredCharacterSets = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CharacterSetPasswordValidatorResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CharacterSetPasswordValidatorResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CharacterSetPasswordValidatorResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CharacterSetPasswordValidatorResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *CharacterSetPasswordValidatorResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CharacterSetPasswordValidatorResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CharacterSetPasswordValidatorResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetValidatorRequirementDescription returns the ValidatorRequirementDescription field value if set, zero value otherwise.
func (o *CharacterSetPasswordValidatorResponse) GetValidatorRequirementDescription() string {
	if o == nil || IsNil(o.ValidatorRequirementDescription) {
		var ret string
		return ret
	}
	return *o.ValidatorRequirementDescription
}

// GetValidatorRequirementDescriptionOk returns a tuple with the ValidatorRequirementDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CharacterSetPasswordValidatorResponse) GetValidatorRequirementDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ValidatorRequirementDescription) {
		return nil, false
	}
	return o.ValidatorRequirementDescription, true
}

// HasValidatorRequirementDescription returns a boolean if a field has been set.
func (o *CharacterSetPasswordValidatorResponse) HasValidatorRequirementDescription() bool {
	if o != nil && !IsNil(o.ValidatorRequirementDescription) {
		return true
	}

	return false
}

// SetValidatorRequirementDescription gets a reference to the given string and assigns it to the ValidatorRequirementDescription field.
func (o *CharacterSetPasswordValidatorResponse) SetValidatorRequirementDescription(v string) {
	o.ValidatorRequirementDescription = &v
}

// GetValidatorFailureMessage returns the ValidatorFailureMessage field value if set, zero value otherwise.
func (o *CharacterSetPasswordValidatorResponse) GetValidatorFailureMessage() string {
	if o == nil || IsNil(o.ValidatorFailureMessage) {
		var ret string
		return ret
	}
	return *o.ValidatorFailureMessage
}

// GetValidatorFailureMessageOk returns a tuple with the ValidatorFailureMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CharacterSetPasswordValidatorResponse) GetValidatorFailureMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ValidatorFailureMessage) {
		return nil, false
	}
	return o.ValidatorFailureMessage, true
}

// HasValidatorFailureMessage returns a boolean if a field has been set.
func (o *CharacterSetPasswordValidatorResponse) HasValidatorFailureMessage() bool {
	if o != nil && !IsNil(o.ValidatorFailureMessage) {
		return true
	}

	return false
}

// SetValidatorFailureMessage gets a reference to the given string and assigns it to the ValidatorFailureMessage field.
func (o *CharacterSetPasswordValidatorResponse) SetValidatorFailureMessage(v string) {
	o.ValidatorFailureMessage = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CharacterSetPasswordValidatorResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CharacterSetPasswordValidatorResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CharacterSetPasswordValidatorResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *CharacterSetPasswordValidatorResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *CharacterSetPasswordValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CharacterSetPasswordValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *CharacterSetPasswordValidatorResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *CharacterSetPasswordValidatorResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o CharacterSetPasswordValidatorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CharacterSetPasswordValidatorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["schemas"] = o.Schemas
	toSerialize["characterSet"] = o.CharacterSet
	toSerialize["allowUnclassifiedCharacters"] = o.AllowUnclassifiedCharacters
	if !IsNil(o.MinimumRequiredCharacterSets) {
		toSerialize["minimumRequiredCharacterSets"] = o.MinimumRequiredCharacterSets
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.ValidatorRequirementDescription) {
		toSerialize["validatorRequirementDescription"] = o.ValidatorRequirementDescription
	}
	if !IsNil(o.ValidatorFailureMessage) {
		toSerialize["validatorFailureMessage"] = o.ValidatorFailureMessage
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableCharacterSetPasswordValidatorResponse struct {
	value *CharacterSetPasswordValidatorResponse
	isSet bool
}

func (v NullableCharacterSetPasswordValidatorResponse) Get() *CharacterSetPasswordValidatorResponse {
	return v.value
}

func (v *NullableCharacterSetPasswordValidatorResponse) Set(val *CharacterSetPasswordValidatorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCharacterSetPasswordValidatorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCharacterSetPasswordValidatorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCharacterSetPasswordValidatorResponse(val *CharacterSetPasswordValidatorResponse) *NullableCharacterSetPasswordValidatorResponse {
	return &NullableCharacterSetPasswordValidatorResponse{value: val, isSet: true}
}

func (v NullableCharacterSetPasswordValidatorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCharacterSetPasswordValidatorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
