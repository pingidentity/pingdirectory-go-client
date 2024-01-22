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

// checks if the UniqueCharactersPasswordValidatorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UniqueCharactersPasswordValidatorResponse{}

// UniqueCharactersPasswordValidatorResponse struct for UniqueCharactersPasswordValidatorResponse
type UniqueCharactersPasswordValidatorResponse struct {
	Schemas []EnumuniqueCharactersPasswordValidatorSchemaUrn `json:"schemas"`
	// Specifies the minimum number of unique characters that a password will be allowed to contain.
	MinUniqueCharacters int64 `json:"minUniqueCharacters"`
	// Indicates whether this password validator should treat password characters in a case-sensitive manner.
	CaseSensitiveValidation bool `json:"caseSensitiveValidation"`
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
	// Name of the Password Validator
	Id string `json:"id"`
}

// NewUniqueCharactersPasswordValidatorResponse instantiates a new UniqueCharactersPasswordValidatorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUniqueCharactersPasswordValidatorResponse(schemas []EnumuniqueCharactersPasswordValidatorSchemaUrn, minUniqueCharacters int64, caseSensitiveValidation bool, enabled bool, id string) *UniqueCharactersPasswordValidatorResponse {
	this := UniqueCharactersPasswordValidatorResponse{}
	this.Schemas = schemas
	this.MinUniqueCharacters = minUniqueCharacters
	this.CaseSensitiveValidation = caseSensitiveValidation
	this.Enabled = enabled
	this.Id = id
	return &this
}

// NewUniqueCharactersPasswordValidatorResponseWithDefaults instantiates a new UniqueCharactersPasswordValidatorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUniqueCharactersPasswordValidatorResponseWithDefaults() *UniqueCharactersPasswordValidatorResponse {
	this := UniqueCharactersPasswordValidatorResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *UniqueCharactersPasswordValidatorResponse) GetSchemas() []EnumuniqueCharactersPasswordValidatorSchemaUrn {
	if o == nil {
		var ret []EnumuniqueCharactersPasswordValidatorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *UniqueCharactersPasswordValidatorResponse) GetSchemasOk() ([]EnumuniqueCharactersPasswordValidatorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *UniqueCharactersPasswordValidatorResponse) SetSchemas(v []EnumuniqueCharactersPasswordValidatorSchemaUrn) {
	o.Schemas = v
}

// GetMinUniqueCharacters returns the MinUniqueCharacters field value
func (o *UniqueCharactersPasswordValidatorResponse) GetMinUniqueCharacters() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MinUniqueCharacters
}

// GetMinUniqueCharactersOk returns a tuple with the MinUniqueCharacters field value
// and a boolean to check if the value has been set.
func (o *UniqueCharactersPasswordValidatorResponse) GetMinUniqueCharactersOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinUniqueCharacters, true
}

// SetMinUniqueCharacters sets field value
func (o *UniqueCharactersPasswordValidatorResponse) SetMinUniqueCharacters(v int64) {
	o.MinUniqueCharacters = v
}

// GetCaseSensitiveValidation returns the CaseSensitiveValidation field value
func (o *UniqueCharactersPasswordValidatorResponse) GetCaseSensitiveValidation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CaseSensitiveValidation
}

// GetCaseSensitiveValidationOk returns a tuple with the CaseSensitiveValidation field value
// and a boolean to check if the value has been set.
func (o *UniqueCharactersPasswordValidatorResponse) GetCaseSensitiveValidationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CaseSensitiveValidation, true
}

// SetCaseSensitiveValidation sets field value
func (o *UniqueCharactersPasswordValidatorResponse) SetCaseSensitiveValidation(v bool) {
	o.CaseSensitiveValidation = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UniqueCharactersPasswordValidatorResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniqueCharactersPasswordValidatorResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UniqueCharactersPasswordValidatorResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UniqueCharactersPasswordValidatorResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *UniqueCharactersPasswordValidatorResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *UniqueCharactersPasswordValidatorResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *UniqueCharactersPasswordValidatorResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetValidatorRequirementDescription returns the ValidatorRequirementDescription field value if set, zero value otherwise.
func (o *UniqueCharactersPasswordValidatorResponse) GetValidatorRequirementDescription() string {
	if o == nil || IsNil(o.ValidatorRequirementDescription) {
		var ret string
		return ret
	}
	return *o.ValidatorRequirementDescription
}

// GetValidatorRequirementDescriptionOk returns a tuple with the ValidatorRequirementDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniqueCharactersPasswordValidatorResponse) GetValidatorRequirementDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ValidatorRequirementDescription) {
		return nil, false
	}
	return o.ValidatorRequirementDescription, true
}

// HasValidatorRequirementDescription returns a boolean if a field has been set.
func (o *UniqueCharactersPasswordValidatorResponse) HasValidatorRequirementDescription() bool {
	if o != nil && !IsNil(o.ValidatorRequirementDescription) {
		return true
	}

	return false
}

// SetValidatorRequirementDescription gets a reference to the given string and assigns it to the ValidatorRequirementDescription field.
func (o *UniqueCharactersPasswordValidatorResponse) SetValidatorRequirementDescription(v string) {
	o.ValidatorRequirementDescription = &v
}

// GetValidatorFailureMessage returns the ValidatorFailureMessage field value if set, zero value otherwise.
func (o *UniqueCharactersPasswordValidatorResponse) GetValidatorFailureMessage() string {
	if o == nil || IsNil(o.ValidatorFailureMessage) {
		var ret string
		return ret
	}
	return *o.ValidatorFailureMessage
}

// GetValidatorFailureMessageOk returns a tuple with the ValidatorFailureMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniqueCharactersPasswordValidatorResponse) GetValidatorFailureMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ValidatorFailureMessage) {
		return nil, false
	}
	return o.ValidatorFailureMessage, true
}

// HasValidatorFailureMessage returns a boolean if a field has been set.
func (o *UniqueCharactersPasswordValidatorResponse) HasValidatorFailureMessage() bool {
	if o != nil && !IsNil(o.ValidatorFailureMessage) {
		return true
	}

	return false
}

// SetValidatorFailureMessage gets a reference to the given string and assigns it to the ValidatorFailureMessage field.
func (o *UniqueCharactersPasswordValidatorResponse) SetValidatorFailureMessage(v string) {
	o.ValidatorFailureMessage = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *UniqueCharactersPasswordValidatorResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniqueCharactersPasswordValidatorResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *UniqueCharactersPasswordValidatorResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *UniqueCharactersPasswordValidatorResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *UniqueCharactersPasswordValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniqueCharactersPasswordValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *UniqueCharactersPasswordValidatorResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *UniqueCharactersPasswordValidatorResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *UniqueCharactersPasswordValidatorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UniqueCharactersPasswordValidatorResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UniqueCharactersPasswordValidatorResponse) SetId(v string) {
	o.Id = v
}

func (o UniqueCharactersPasswordValidatorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UniqueCharactersPasswordValidatorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["minUniqueCharacters"] = o.MinUniqueCharacters
	toSerialize["caseSensitiveValidation"] = o.CaseSensitiveValidation
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
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableUniqueCharactersPasswordValidatorResponse struct {
	value *UniqueCharactersPasswordValidatorResponse
	isSet bool
}

func (v NullableUniqueCharactersPasswordValidatorResponse) Get() *UniqueCharactersPasswordValidatorResponse {
	return v.value
}

func (v *NullableUniqueCharactersPasswordValidatorResponse) Set(val *UniqueCharactersPasswordValidatorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUniqueCharactersPasswordValidatorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUniqueCharactersPasswordValidatorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUniqueCharactersPasswordValidatorResponse(val *UniqueCharactersPasswordValidatorResponse) *NullableUniqueCharactersPasswordValidatorResponse {
	return &NullableUniqueCharactersPasswordValidatorResponse{value: val, isSet: true}
}

func (v NullableUniqueCharactersPasswordValidatorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUniqueCharactersPasswordValidatorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
