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

// checks if the Utf8PasswordValidatorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Utf8PasswordValidatorResponse{}

// Utf8PasswordValidatorResponse struct for Utf8PasswordValidatorResponse
type Utf8PasswordValidatorResponse struct {
	Schemas []Enumutf8PasswordValidatorSchemaUrn `json:"schemas"`
	// Indicates whether passwords will be allowed to include characters from outside the ASCII character set.
	AllowNonAsciiCharacters *bool `json:"allowNonAsciiCharacters,omitempty"`
	// Indicates whether passwords will be allowed to include characters that are not recognized by the JVM's Unicode support.
	AllowUnknownCharacters *bool                                           `json:"allowUnknownCharacters,omitempty"`
	AllowedCharacterType   []EnumpasswordValidatorAllowedCharacterTypeProp `json:"allowedCharacterType,omitempty"`
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

// NewUtf8PasswordValidatorResponse instantiates a new Utf8PasswordValidatorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUtf8PasswordValidatorResponse(schemas []Enumutf8PasswordValidatorSchemaUrn, enabled bool, id string) *Utf8PasswordValidatorResponse {
	this := Utf8PasswordValidatorResponse{}
	this.Schemas = schemas
	this.Enabled = enabled
	this.Id = id
	return &this
}

// NewUtf8PasswordValidatorResponseWithDefaults instantiates a new Utf8PasswordValidatorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUtf8PasswordValidatorResponseWithDefaults() *Utf8PasswordValidatorResponse {
	this := Utf8PasswordValidatorResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *Utf8PasswordValidatorResponse) GetSchemas() []Enumutf8PasswordValidatorSchemaUrn {
	if o == nil {
		var ret []Enumutf8PasswordValidatorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *Utf8PasswordValidatorResponse) GetSchemasOk() ([]Enumutf8PasswordValidatorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *Utf8PasswordValidatorResponse) SetSchemas(v []Enumutf8PasswordValidatorSchemaUrn) {
	o.Schemas = v
}

// GetAllowNonAsciiCharacters returns the AllowNonAsciiCharacters field value if set, zero value otherwise.
func (o *Utf8PasswordValidatorResponse) GetAllowNonAsciiCharacters() bool {
	if o == nil || IsNil(o.AllowNonAsciiCharacters) {
		var ret bool
		return ret
	}
	return *o.AllowNonAsciiCharacters
}

// GetAllowNonAsciiCharactersOk returns a tuple with the AllowNonAsciiCharacters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Utf8PasswordValidatorResponse) GetAllowNonAsciiCharactersOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowNonAsciiCharacters) {
		return nil, false
	}
	return o.AllowNonAsciiCharacters, true
}

// HasAllowNonAsciiCharacters returns a boolean if a field has been set.
func (o *Utf8PasswordValidatorResponse) HasAllowNonAsciiCharacters() bool {
	if o != nil && !IsNil(o.AllowNonAsciiCharacters) {
		return true
	}

	return false
}

// SetAllowNonAsciiCharacters gets a reference to the given bool and assigns it to the AllowNonAsciiCharacters field.
func (o *Utf8PasswordValidatorResponse) SetAllowNonAsciiCharacters(v bool) {
	o.AllowNonAsciiCharacters = &v
}

// GetAllowUnknownCharacters returns the AllowUnknownCharacters field value if set, zero value otherwise.
func (o *Utf8PasswordValidatorResponse) GetAllowUnknownCharacters() bool {
	if o == nil || IsNil(o.AllowUnknownCharacters) {
		var ret bool
		return ret
	}
	return *o.AllowUnknownCharacters
}

// GetAllowUnknownCharactersOk returns a tuple with the AllowUnknownCharacters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Utf8PasswordValidatorResponse) GetAllowUnknownCharactersOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowUnknownCharacters) {
		return nil, false
	}
	return o.AllowUnknownCharacters, true
}

// HasAllowUnknownCharacters returns a boolean if a field has been set.
func (o *Utf8PasswordValidatorResponse) HasAllowUnknownCharacters() bool {
	if o != nil && !IsNil(o.AllowUnknownCharacters) {
		return true
	}

	return false
}

// SetAllowUnknownCharacters gets a reference to the given bool and assigns it to the AllowUnknownCharacters field.
func (o *Utf8PasswordValidatorResponse) SetAllowUnknownCharacters(v bool) {
	o.AllowUnknownCharacters = &v
}

// GetAllowedCharacterType returns the AllowedCharacterType field value if set, zero value otherwise.
func (o *Utf8PasswordValidatorResponse) GetAllowedCharacterType() []EnumpasswordValidatorAllowedCharacterTypeProp {
	if o == nil || IsNil(o.AllowedCharacterType) {
		var ret []EnumpasswordValidatorAllowedCharacterTypeProp
		return ret
	}
	return o.AllowedCharacterType
}

// GetAllowedCharacterTypeOk returns a tuple with the AllowedCharacterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Utf8PasswordValidatorResponse) GetAllowedCharacterTypeOk() ([]EnumpasswordValidatorAllowedCharacterTypeProp, bool) {
	if o == nil || IsNil(o.AllowedCharacterType) {
		return nil, false
	}
	return o.AllowedCharacterType, true
}

// HasAllowedCharacterType returns a boolean if a field has been set.
func (o *Utf8PasswordValidatorResponse) HasAllowedCharacterType() bool {
	if o != nil && !IsNil(o.AllowedCharacterType) {
		return true
	}

	return false
}

// SetAllowedCharacterType gets a reference to the given []EnumpasswordValidatorAllowedCharacterTypeProp and assigns it to the AllowedCharacterType field.
func (o *Utf8PasswordValidatorResponse) SetAllowedCharacterType(v []EnumpasswordValidatorAllowedCharacterTypeProp) {
	o.AllowedCharacterType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Utf8PasswordValidatorResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Utf8PasswordValidatorResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Utf8PasswordValidatorResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Utf8PasswordValidatorResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *Utf8PasswordValidatorResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Utf8PasswordValidatorResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *Utf8PasswordValidatorResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetValidatorRequirementDescription returns the ValidatorRequirementDescription field value if set, zero value otherwise.
func (o *Utf8PasswordValidatorResponse) GetValidatorRequirementDescription() string {
	if o == nil || IsNil(o.ValidatorRequirementDescription) {
		var ret string
		return ret
	}
	return *o.ValidatorRequirementDescription
}

// GetValidatorRequirementDescriptionOk returns a tuple with the ValidatorRequirementDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Utf8PasswordValidatorResponse) GetValidatorRequirementDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ValidatorRequirementDescription) {
		return nil, false
	}
	return o.ValidatorRequirementDescription, true
}

// HasValidatorRequirementDescription returns a boolean if a field has been set.
func (o *Utf8PasswordValidatorResponse) HasValidatorRequirementDescription() bool {
	if o != nil && !IsNil(o.ValidatorRequirementDescription) {
		return true
	}

	return false
}

// SetValidatorRequirementDescription gets a reference to the given string and assigns it to the ValidatorRequirementDescription field.
func (o *Utf8PasswordValidatorResponse) SetValidatorRequirementDescription(v string) {
	o.ValidatorRequirementDescription = &v
}

// GetValidatorFailureMessage returns the ValidatorFailureMessage field value if set, zero value otherwise.
func (o *Utf8PasswordValidatorResponse) GetValidatorFailureMessage() string {
	if o == nil || IsNil(o.ValidatorFailureMessage) {
		var ret string
		return ret
	}
	return *o.ValidatorFailureMessage
}

// GetValidatorFailureMessageOk returns a tuple with the ValidatorFailureMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Utf8PasswordValidatorResponse) GetValidatorFailureMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ValidatorFailureMessage) {
		return nil, false
	}
	return o.ValidatorFailureMessage, true
}

// HasValidatorFailureMessage returns a boolean if a field has been set.
func (o *Utf8PasswordValidatorResponse) HasValidatorFailureMessage() bool {
	if o != nil && !IsNil(o.ValidatorFailureMessage) {
		return true
	}

	return false
}

// SetValidatorFailureMessage gets a reference to the given string and assigns it to the ValidatorFailureMessage field.
func (o *Utf8PasswordValidatorResponse) SetValidatorFailureMessage(v string) {
	o.ValidatorFailureMessage = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Utf8PasswordValidatorResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Utf8PasswordValidatorResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Utf8PasswordValidatorResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *Utf8PasswordValidatorResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *Utf8PasswordValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Utf8PasswordValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *Utf8PasswordValidatorResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *Utf8PasswordValidatorResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *Utf8PasswordValidatorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Utf8PasswordValidatorResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Utf8PasswordValidatorResponse) SetId(v string) {
	o.Id = v
}

func (o Utf8PasswordValidatorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Utf8PasswordValidatorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.AllowNonAsciiCharacters) {
		toSerialize["allowNonAsciiCharacters"] = o.AllowNonAsciiCharacters
	}
	if !IsNil(o.AllowUnknownCharacters) {
		toSerialize["allowUnknownCharacters"] = o.AllowUnknownCharacters
	}
	if !IsNil(o.AllowedCharacterType) {
		toSerialize["allowedCharacterType"] = o.AllowedCharacterType
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
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableUtf8PasswordValidatorResponse struct {
	value *Utf8PasswordValidatorResponse
	isSet bool
}

func (v NullableUtf8PasswordValidatorResponse) Get() *Utf8PasswordValidatorResponse {
	return v.value
}

func (v *NullableUtf8PasswordValidatorResponse) Set(val *Utf8PasswordValidatorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUtf8PasswordValidatorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUtf8PasswordValidatorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUtf8PasswordValidatorResponse(val *Utf8PasswordValidatorResponse) *NullableUtf8PasswordValidatorResponse {
	return &NullableUtf8PasswordValidatorResponse{value: val, isSet: true}
}

func (v NullableUtf8PasswordValidatorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUtf8PasswordValidatorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
