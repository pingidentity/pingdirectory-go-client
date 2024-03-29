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

// checks if the RegularExpressionPasswordValidatorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegularExpressionPasswordValidatorResponse{}

// RegularExpressionPasswordValidatorResponse struct for RegularExpressionPasswordValidatorResponse
type RegularExpressionPasswordValidatorResponse struct {
	Schemas []EnumregularExpressionPasswordValidatorSchemaUrn `json:"schemas"`
	// The regular expression to use for this password validator.
	MatchPattern  string                                 `json:"matchPattern"`
	MatchBehavior EnumpasswordValidatorMatchBehaviorProp `json:"matchBehavior"`
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

// NewRegularExpressionPasswordValidatorResponse instantiates a new RegularExpressionPasswordValidatorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegularExpressionPasswordValidatorResponse(schemas []EnumregularExpressionPasswordValidatorSchemaUrn, matchPattern string, matchBehavior EnumpasswordValidatorMatchBehaviorProp, enabled bool, id string) *RegularExpressionPasswordValidatorResponse {
	this := RegularExpressionPasswordValidatorResponse{}
	this.Schemas = schemas
	this.MatchPattern = matchPattern
	this.MatchBehavior = matchBehavior
	this.Enabled = enabled
	this.Id = id
	return &this
}

// NewRegularExpressionPasswordValidatorResponseWithDefaults instantiates a new RegularExpressionPasswordValidatorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegularExpressionPasswordValidatorResponseWithDefaults() *RegularExpressionPasswordValidatorResponse {
	this := RegularExpressionPasswordValidatorResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *RegularExpressionPasswordValidatorResponse) GetSchemas() []EnumregularExpressionPasswordValidatorSchemaUrn {
	if o == nil {
		var ret []EnumregularExpressionPasswordValidatorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *RegularExpressionPasswordValidatorResponse) GetSchemasOk() ([]EnumregularExpressionPasswordValidatorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *RegularExpressionPasswordValidatorResponse) SetSchemas(v []EnumregularExpressionPasswordValidatorSchemaUrn) {
	o.Schemas = v
}

// GetMatchPattern returns the MatchPattern field value
func (o *RegularExpressionPasswordValidatorResponse) GetMatchPattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MatchPattern
}

// GetMatchPatternOk returns a tuple with the MatchPattern field value
// and a boolean to check if the value has been set.
func (o *RegularExpressionPasswordValidatorResponse) GetMatchPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchPattern, true
}

// SetMatchPattern sets field value
func (o *RegularExpressionPasswordValidatorResponse) SetMatchPattern(v string) {
	o.MatchPattern = v
}

// GetMatchBehavior returns the MatchBehavior field value
func (o *RegularExpressionPasswordValidatorResponse) GetMatchBehavior() EnumpasswordValidatorMatchBehaviorProp {
	if o == nil {
		var ret EnumpasswordValidatorMatchBehaviorProp
		return ret
	}

	return o.MatchBehavior
}

// GetMatchBehaviorOk returns a tuple with the MatchBehavior field value
// and a boolean to check if the value has been set.
func (o *RegularExpressionPasswordValidatorResponse) GetMatchBehaviorOk() (*EnumpasswordValidatorMatchBehaviorProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchBehavior, true
}

// SetMatchBehavior sets field value
func (o *RegularExpressionPasswordValidatorResponse) SetMatchBehavior(v EnumpasswordValidatorMatchBehaviorProp) {
	o.MatchBehavior = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RegularExpressionPasswordValidatorResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegularExpressionPasswordValidatorResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RegularExpressionPasswordValidatorResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RegularExpressionPasswordValidatorResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *RegularExpressionPasswordValidatorResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *RegularExpressionPasswordValidatorResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *RegularExpressionPasswordValidatorResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetValidatorRequirementDescription returns the ValidatorRequirementDescription field value if set, zero value otherwise.
func (o *RegularExpressionPasswordValidatorResponse) GetValidatorRequirementDescription() string {
	if o == nil || IsNil(o.ValidatorRequirementDescription) {
		var ret string
		return ret
	}
	return *o.ValidatorRequirementDescription
}

// GetValidatorRequirementDescriptionOk returns a tuple with the ValidatorRequirementDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegularExpressionPasswordValidatorResponse) GetValidatorRequirementDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ValidatorRequirementDescription) {
		return nil, false
	}
	return o.ValidatorRequirementDescription, true
}

// HasValidatorRequirementDescription returns a boolean if a field has been set.
func (o *RegularExpressionPasswordValidatorResponse) HasValidatorRequirementDescription() bool {
	if o != nil && !IsNil(o.ValidatorRequirementDescription) {
		return true
	}

	return false
}

// SetValidatorRequirementDescription gets a reference to the given string and assigns it to the ValidatorRequirementDescription field.
func (o *RegularExpressionPasswordValidatorResponse) SetValidatorRequirementDescription(v string) {
	o.ValidatorRequirementDescription = &v
}

// GetValidatorFailureMessage returns the ValidatorFailureMessage field value if set, zero value otherwise.
func (o *RegularExpressionPasswordValidatorResponse) GetValidatorFailureMessage() string {
	if o == nil || IsNil(o.ValidatorFailureMessage) {
		var ret string
		return ret
	}
	return *o.ValidatorFailureMessage
}

// GetValidatorFailureMessageOk returns a tuple with the ValidatorFailureMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegularExpressionPasswordValidatorResponse) GetValidatorFailureMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ValidatorFailureMessage) {
		return nil, false
	}
	return o.ValidatorFailureMessage, true
}

// HasValidatorFailureMessage returns a boolean if a field has been set.
func (o *RegularExpressionPasswordValidatorResponse) HasValidatorFailureMessage() bool {
	if o != nil && !IsNil(o.ValidatorFailureMessage) {
		return true
	}

	return false
}

// SetValidatorFailureMessage gets a reference to the given string and assigns it to the ValidatorFailureMessage field.
func (o *RegularExpressionPasswordValidatorResponse) SetValidatorFailureMessage(v string) {
	o.ValidatorFailureMessage = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *RegularExpressionPasswordValidatorResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegularExpressionPasswordValidatorResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *RegularExpressionPasswordValidatorResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *RegularExpressionPasswordValidatorResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *RegularExpressionPasswordValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegularExpressionPasswordValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *RegularExpressionPasswordValidatorResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *RegularExpressionPasswordValidatorResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *RegularExpressionPasswordValidatorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RegularExpressionPasswordValidatorResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RegularExpressionPasswordValidatorResponse) SetId(v string) {
	o.Id = v
}

func (o RegularExpressionPasswordValidatorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegularExpressionPasswordValidatorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["matchPattern"] = o.MatchPattern
	toSerialize["matchBehavior"] = o.MatchBehavior
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

type NullableRegularExpressionPasswordValidatorResponse struct {
	value *RegularExpressionPasswordValidatorResponse
	isSet bool
}

func (v NullableRegularExpressionPasswordValidatorResponse) Get() *RegularExpressionPasswordValidatorResponse {
	return v.value
}

func (v *NullableRegularExpressionPasswordValidatorResponse) Set(val *RegularExpressionPasswordValidatorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRegularExpressionPasswordValidatorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRegularExpressionPasswordValidatorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegularExpressionPasswordValidatorResponse(val *RegularExpressionPasswordValidatorResponse) *NullableRegularExpressionPasswordValidatorResponse {
	return &NullableRegularExpressionPasswordValidatorResponse{value: val, isSet: true}
}

func (v NullableRegularExpressionPasswordValidatorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegularExpressionPasswordValidatorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
