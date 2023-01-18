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

// HaystackPasswordValidatorResponse struct for HaystackPasswordValidatorResponse
type HaystackPasswordValidatorResponse struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Password Validator
	Id string `json:"id"`
	Schemas []EnumhaystackPasswordValidatorSchemaUrn `json:"schemas"`
	// The number of password guesses per second that a potential attacker may be expected to make.
	AssumedPasswordGuessesPerSecond string `json:"assumedPasswordGuessesPerSecond"`
	// The minimum length of time (using the configured number of password guesses per second) required to exhaust the entire search space for a proposed password in order for that password to be considered acceptable.
	MinimumAcceptableTimeToExhaustSearchSpace string `json:"minimumAcceptableTimeToExhaustSearchSpace"`
	// A description for this Password Validator
	Description *string `json:"description,omitempty"`
	// Indicates whether the password validator is enabled for use.
	Enabled bool `json:"enabled"`
	// Specifies a message that can be used to describe the requirements imposed by this password validator to end users. If a value is provided for this property, then it will override any description that may have otherwise been generated by the validator.
	ValidatorRequirementDescription *string `json:"validatorRequirementDescription,omitempty"`
	// Specifies a message that may be provided to the end user in the event that a proposed password is rejected by this validator. If a value is provided for this property, then it will override any failure message that may have otherwise been generated by the validator.
	ValidatorFailureMessage *string `json:"validatorFailureMessage,omitempty"`
}

// NewHaystackPasswordValidatorResponse instantiates a new HaystackPasswordValidatorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHaystackPasswordValidatorResponse(id string, schemas []EnumhaystackPasswordValidatorSchemaUrn, assumedPasswordGuessesPerSecond string, minimumAcceptableTimeToExhaustSearchSpace string, enabled bool) *HaystackPasswordValidatorResponse {
	this := HaystackPasswordValidatorResponse{}
	this.Id = id
	this.Schemas = schemas
	this.AssumedPasswordGuessesPerSecond = assumedPasswordGuessesPerSecond
	this.MinimumAcceptableTimeToExhaustSearchSpace = minimumAcceptableTimeToExhaustSearchSpace
	this.Enabled = enabled
	return &this
}

// NewHaystackPasswordValidatorResponseWithDefaults instantiates a new HaystackPasswordValidatorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHaystackPasswordValidatorResponseWithDefaults() *HaystackPasswordValidatorResponse {
	this := HaystackPasswordValidatorResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *HaystackPasswordValidatorResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HaystackPasswordValidatorResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *HaystackPasswordValidatorResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *HaystackPasswordValidatorResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *HaystackPasswordValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HaystackPasswordValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
    return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *HaystackPasswordValidatorResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *HaystackPasswordValidatorResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *HaystackPasswordValidatorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *HaystackPasswordValidatorResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *HaystackPasswordValidatorResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *HaystackPasswordValidatorResponse) GetSchemas() []EnumhaystackPasswordValidatorSchemaUrn {
	if o == nil {
		var ret []EnumhaystackPasswordValidatorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *HaystackPasswordValidatorResponse) GetSchemasOk() ([]EnumhaystackPasswordValidatorSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *HaystackPasswordValidatorResponse) SetSchemas(v []EnumhaystackPasswordValidatorSchemaUrn) {
	o.Schemas = v
}

// GetAssumedPasswordGuessesPerSecond returns the AssumedPasswordGuessesPerSecond field value
func (o *HaystackPasswordValidatorResponse) GetAssumedPasswordGuessesPerSecond() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssumedPasswordGuessesPerSecond
}

// GetAssumedPasswordGuessesPerSecondOk returns a tuple with the AssumedPasswordGuessesPerSecond field value
// and a boolean to check if the value has been set.
func (o *HaystackPasswordValidatorResponse) GetAssumedPasswordGuessesPerSecondOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AssumedPasswordGuessesPerSecond, true
}

// SetAssumedPasswordGuessesPerSecond sets field value
func (o *HaystackPasswordValidatorResponse) SetAssumedPasswordGuessesPerSecond(v string) {
	o.AssumedPasswordGuessesPerSecond = v
}

// GetMinimumAcceptableTimeToExhaustSearchSpace returns the MinimumAcceptableTimeToExhaustSearchSpace field value
func (o *HaystackPasswordValidatorResponse) GetMinimumAcceptableTimeToExhaustSearchSpace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinimumAcceptableTimeToExhaustSearchSpace
}

// GetMinimumAcceptableTimeToExhaustSearchSpaceOk returns a tuple with the MinimumAcceptableTimeToExhaustSearchSpace field value
// and a boolean to check if the value has been set.
func (o *HaystackPasswordValidatorResponse) GetMinimumAcceptableTimeToExhaustSearchSpaceOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MinimumAcceptableTimeToExhaustSearchSpace, true
}

// SetMinimumAcceptableTimeToExhaustSearchSpace sets field value
func (o *HaystackPasswordValidatorResponse) SetMinimumAcceptableTimeToExhaustSearchSpace(v string) {
	o.MinimumAcceptableTimeToExhaustSearchSpace = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *HaystackPasswordValidatorResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HaystackPasswordValidatorResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *HaystackPasswordValidatorResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *HaystackPasswordValidatorResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *HaystackPasswordValidatorResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *HaystackPasswordValidatorResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *HaystackPasswordValidatorResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetValidatorRequirementDescription returns the ValidatorRequirementDescription field value if set, zero value otherwise.
func (o *HaystackPasswordValidatorResponse) GetValidatorRequirementDescription() string {
	if o == nil || isNil(o.ValidatorRequirementDescription) {
		var ret string
		return ret
	}
	return *o.ValidatorRequirementDescription
}

// GetValidatorRequirementDescriptionOk returns a tuple with the ValidatorRequirementDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HaystackPasswordValidatorResponse) GetValidatorRequirementDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.ValidatorRequirementDescription) {
    return nil, false
	}
	return o.ValidatorRequirementDescription, true
}

// HasValidatorRequirementDescription returns a boolean if a field has been set.
func (o *HaystackPasswordValidatorResponse) HasValidatorRequirementDescription() bool {
	if o != nil && !isNil(o.ValidatorRequirementDescription) {
		return true
	}

	return false
}

// SetValidatorRequirementDescription gets a reference to the given string and assigns it to the ValidatorRequirementDescription field.
func (o *HaystackPasswordValidatorResponse) SetValidatorRequirementDescription(v string) {
	o.ValidatorRequirementDescription = &v
}

// GetValidatorFailureMessage returns the ValidatorFailureMessage field value if set, zero value otherwise.
func (o *HaystackPasswordValidatorResponse) GetValidatorFailureMessage() string {
	if o == nil || isNil(o.ValidatorFailureMessage) {
		var ret string
		return ret
	}
	return *o.ValidatorFailureMessage
}

// GetValidatorFailureMessageOk returns a tuple with the ValidatorFailureMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HaystackPasswordValidatorResponse) GetValidatorFailureMessageOk() (*string, bool) {
	if o == nil || isNil(o.ValidatorFailureMessage) {
    return nil, false
	}
	return o.ValidatorFailureMessage, true
}

// HasValidatorFailureMessage returns a boolean if a field has been set.
func (o *HaystackPasswordValidatorResponse) HasValidatorFailureMessage() bool {
	if o != nil && !isNil(o.ValidatorFailureMessage) {
		return true
	}

	return false
}

// SetValidatorFailureMessage gets a reference to the given string and assigns it to the ValidatorFailureMessage field.
func (o *HaystackPasswordValidatorResponse) SetValidatorFailureMessage(v string) {
	o.ValidatorFailureMessage = &v
}

func (o HaystackPasswordValidatorResponse) MarshalJSON() ([]byte, error) {
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
		toSerialize["assumedPasswordGuessesPerSecond"] = o.AssumedPasswordGuessesPerSecond
	}
	if true {
		toSerialize["minimumAcceptableTimeToExhaustSearchSpace"] = o.MinimumAcceptableTimeToExhaustSearchSpace
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.ValidatorRequirementDescription) {
		toSerialize["validatorRequirementDescription"] = o.ValidatorRequirementDescription
	}
	if !isNil(o.ValidatorFailureMessage) {
		toSerialize["validatorFailureMessage"] = o.ValidatorFailureMessage
	}
	return json.Marshal(toSerialize)
}

type NullableHaystackPasswordValidatorResponse struct {
	value *HaystackPasswordValidatorResponse
	isSet bool
}

func (v NullableHaystackPasswordValidatorResponse) Get() *HaystackPasswordValidatorResponse {
	return v.value
}

func (v *NullableHaystackPasswordValidatorResponse) Set(val *HaystackPasswordValidatorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHaystackPasswordValidatorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHaystackPasswordValidatorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHaystackPasswordValidatorResponse(val *HaystackPasswordValidatorResponse) *NullableHaystackPasswordValidatorResponse {
	return &NullableHaystackPasswordValidatorResponse{value: val, isSet: true}
}

func (v NullableHaystackPasswordValidatorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHaystackPasswordValidatorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

