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

// AddTextAccessLogFieldBehaviorRequest struct for AddTextAccessLogFieldBehaviorRequest
type AddTextAccessLogFieldBehaviorRequest struct {
	// Name of the new Log Field Behavior
	BehaviorName  string                                    `json:"behaviorName"`
	Schemas       []EnumtextAccessLogFieldBehaviorSchemaUrn `json:"schemas"`
	PreserveField []EnumlogFieldBehaviorPreserveFieldProp   `json:"preserveField,omitempty"`
	// The names of any custom fields whose values should be preserved. This should generally only be used for fields that are not available through the preserve-field property (for example, custom log fields defined in Server SDK extensions).
	PreserveFieldName []string                            `json:"preserveFieldName,omitempty"`
	OmitField         []EnumlogFieldBehaviorOmitFieldProp `json:"omitField,omitempty"`
	// The names of any custom fields that should be omitted from log messages. This should generally only be used for fields that are not available through the omit-field property (for example, custom log fields defined in Server SDK extensions).
	OmitFieldName          []string                                         `json:"omitFieldName,omitempty"`
	RedactEntireValueField []EnumlogFieldBehaviorRedactEntireValueFieldProp `json:"redactEntireValueField,omitempty"`
	// The names of any custom fields whose values should be completely redacted. This should generally only be used for fields that are not available through the redact-entire-value-field property (for example, custom log fields defined in Server SDK extensions).
	RedactEntireValueFieldName []string                                             `json:"redactEntireValueFieldName,omitempty"`
	RedactValueComponentsField []EnumlogFieldBehaviorRedactValueComponentsFieldProp `json:"redactValueComponentsField,omitempty"`
	// The names of any custom fields for which to redact components within the value. This should generally only be used for fields that are not available through the redact-value-components-field property (for example, custom log fields defined in Server SDK extensions).
	RedactValueComponentsFieldName []string                                           `json:"redactValueComponentsFieldName,omitempty"`
	TokenizeEntireValueField       []EnumlogFieldBehaviorTokenizeEntireValueFieldProp `json:"tokenizeEntireValueField,omitempty"`
	// The names of any custom fields whose values should be completely tokenized. This should generally only be used for fields that are not available through the tokenize-entire-value-field property (for example, custom log fields defined in Server SDK extensions).
	TokenizeEntireValueFieldName []string                                               `json:"tokenizeEntireValueFieldName,omitempty"`
	TokenizeValueComponentsField []EnumlogFieldBehaviorTokenizeValueComponentsFieldProp `json:"tokenizeValueComponentsField,omitempty"`
	// The names of any custom fields for which to tokenize components within the value. This should generally only be used for fields that are not available through the tokenize-value-components-field property (for example, custom log fields defined in Server SDK extensions).
	TokenizeValueComponentsFieldName []string `json:"tokenizeValueComponentsFieldName,omitempty"`
	// A description for this Log Field Behavior
	Description     *string                                  `json:"description,omitempty"`
	DefaultBehavior *EnumlogFieldBehaviorDefaultBehaviorProp `json:"defaultBehavior,omitempty"`
}

// NewAddTextAccessLogFieldBehaviorRequest instantiates a new AddTextAccessLogFieldBehaviorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddTextAccessLogFieldBehaviorRequest(behaviorName string, schemas []EnumtextAccessLogFieldBehaviorSchemaUrn) *AddTextAccessLogFieldBehaviorRequest {
	this := AddTextAccessLogFieldBehaviorRequest{}
	this.BehaviorName = behaviorName
	this.Schemas = schemas
	return &this
}

// NewAddTextAccessLogFieldBehaviorRequestWithDefaults instantiates a new AddTextAccessLogFieldBehaviorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddTextAccessLogFieldBehaviorRequestWithDefaults() *AddTextAccessLogFieldBehaviorRequest {
	this := AddTextAccessLogFieldBehaviorRequest{}
	return &this
}

// GetBehaviorName returns the BehaviorName field value
func (o *AddTextAccessLogFieldBehaviorRequest) GetBehaviorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BehaviorName
}

// GetBehaviorNameOk returns a tuple with the BehaviorName field value
// and a boolean to check if the value has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) GetBehaviorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BehaviorName, true
}

// SetBehaviorName sets field value
func (o *AddTextAccessLogFieldBehaviorRequest) SetBehaviorName(v string) {
	o.BehaviorName = v
}

// GetSchemas returns the Schemas field value
func (o *AddTextAccessLogFieldBehaviorRequest) GetSchemas() []EnumtextAccessLogFieldBehaviorSchemaUrn {
	if o == nil {
		var ret []EnumtextAccessLogFieldBehaviorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) GetSchemasOk() ([]EnumtextAccessLogFieldBehaviorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddTextAccessLogFieldBehaviorRequest) SetSchemas(v []EnumtextAccessLogFieldBehaviorSchemaUrn) {
	o.Schemas = v
}

// GetPreserveField returns the PreserveField field value if set, zero value otherwise.
func (o *AddTextAccessLogFieldBehaviorRequest) GetPreserveField() []EnumlogFieldBehaviorPreserveFieldProp {
	if o == nil || isNil(o.PreserveField) {
		var ret []EnumlogFieldBehaviorPreserveFieldProp
		return ret
	}
	return o.PreserveField
}

// GetPreserveFieldOk returns a tuple with the PreserveField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) GetPreserveFieldOk() ([]EnumlogFieldBehaviorPreserveFieldProp, bool) {
	if o == nil || isNil(o.PreserveField) {
		return nil, false
	}
	return o.PreserveField, true
}

// HasPreserveField returns a boolean if a field has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) HasPreserveField() bool {
	if o != nil && !isNil(o.PreserveField) {
		return true
	}

	return false
}

// SetPreserveField gets a reference to the given []EnumlogFieldBehaviorPreserveFieldProp and assigns it to the PreserveField field.
func (o *AddTextAccessLogFieldBehaviorRequest) SetPreserveField(v []EnumlogFieldBehaviorPreserveFieldProp) {
	o.PreserveField = v
}

// GetPreserveFieldName returns the PreserveFieldName field value if set, zero value otherwise.
func (o *AddTextAccessLogFieldBehaviorRequest) GetPreserveFieldName() []string {
	if o == nil || isNil(o.PreserveFieldName) {
		var ret []string
		return ret
	}
	return o.PreserveFieldName
}

// GetPreserveFieldNameOk returns a tuple with the PreserveFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) GetPreserveFieldNameOk() ([]string, bool) {
	if o == nil || isNil(o.PreserveFieldName) {
		return nil, false
	}
	return o.PreserveFieldName, true
}

// HasPreserveFieldName returns a boolean if a field has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) HasPreserveFieldName() bool {
	if o != nil && !isNil(o.PreserveFieldName) {
		return true
	}

	return false
}

// SetPreserveFieldName gets a reference to the given []string and assigns it to the PreserveFieldName field.
func (o *AddTextAccessLogFieldBehaviorRequest) SetPreserveFieldName(v []string) {
	o.PreserveFieldName = v
}

// GetOmitField returns the OmitField field value if set, zero value otherwise.
func (o *AddTextAccessLogFieldBehaviorRequest) GetOmitField() []EnumlogFieldBehaviorOmitFieldProp {
	if o == nil || isNil(o.OmitField) {
		var ret []EnumlogFieldBehaviorOmitFieldProp
		return ret
	}
	return o.OmitField
}

// GetOmitFieldOk returns a tuple with the OmitField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) GetOmitFieldOk() ([]EnumlogFieldBehaviorOmitFieldProp, bool) {
	if o == nil || isNil(o.OmitField) {
		return nil, false
	}
	return o.OmitField, true
}

// HasOmitField returns a boolean if a field has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) HasOmitField() bool {
	if o != nil && !isNil(o.OmitField) {
		return true
	}

	return false
}

// SetOmitField gets a reference to the given []EnumlogFieldBehaviorOmitFieldProp and assigns it to the OmitField field.
func (o *AddTextAccessLogFieldBehaviorRequest) SetOmitField(v []EnumlogFieldBehaviorOmitFieldProp) {
	o.OmitField = v
}

// GetOmitFieldName returns the OmitFieldName field value if set, zero value otherwise.
func (o *AddTextAccessLogFieldBehaviorRequest) GetOmitFieldName() []string {
	if o == nil || isNil(o.OmitFieldName) {
		var ret []string
		return ret
	}
	return o.OmitFieldName
}

// GetOmitFieldNameOk returns a tuple with the OmitFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) GetOmitFieldNameOk() ([]string, bool) {
	if o == nil || isNil(o.OmitFieldName) {
		return nil, false
	}
	return o.OmitFieldName, true
}

// HasOmitFieldName returns a boolean if a field has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) HasOmitFieldName() bool {
	if o != nil && !isNil(o.OmitFieldName) {
		return true
	}

	return false
}

// SetOmitFieldName gets a reference to the given []string and assigns it to the OmitFieldName field.
func (o *AddTextAccessLogFieldBehaviorRequest) SetOmitFieldName(v []string) {
	o.OmitFieldName = v
}

// GetRedactEntireValueField returns the RedactEntireValueField field value if set, zero value otherwise.
func (o *AddTextAccessLogFieldBehaviorRequest) GetRedactEntireValueField() []EnumlogFieldBehaviorRedactEntireValueFieldProp {
	if o == nil || isNil(o.RedactEntireValueField) {
		var ret []EnumlogFieldBehaviorRedactEntireValueFieldProp
		return ret
	}
	return o.RedactEntireValueField
}

// GetRedactEntireValueFieldOk returns a tuple with the RedactEntireValueField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) GetRedactEntireValueFieldOk() ([]EnumlogFieldBehaviorRedactEntireValueFieldProp, bool) {
	if o == nil || isNil(o.RedactEntireValueField) {
		return nil, false
	}
	return o.RedactEntireValueField, true
}

// HasRedactEntireValueField returns a boolean if a field has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) HasRedactEntireValueField() bool {
	if o != nil && !isNil(o.RedactEntireValueField) {
		return true
	}

	return false
}

// SetRedactEntireValueField gets a reference to the given []EnumlogFieldBehaviorRedactEntireValueFieldProp and assigns it to the RedactEntireValueField field.
func (o *AddTextAccessLogFieldBehaviorRequest) SetRedactEntireValueField(v []EnumlogFieldBehaviorRedactEntireValueFieldProp) {
	o.RedactEntireValueField = v
}

// GetRedactEntireValueFieldName returns the RedactEntireValueFieldName field value if set, zero value otherwise.
func (o *AddTextAccessLogFieldBehaviorRequest) GetRedactEntireValueFieldName() []string {
	if o == nil || isNil(o.RedactEntireValueFieldName) {
		var ret []string
		return ret
	}
	return o.RedactEntireValueFieldName
}

// GetRedactEntireValueFieldNameOk returns a tuple with the RedactEntireValueFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) GetRedactEntireValueFieldNameOk() ([]string, bool) {
	if o == nil || isNil(o.RedactEntireValueFieldName) {
		return nil, false
	}
	return o.RedactEntireValueFieldName, true
}

// HasRedactEntireValueFieldName returns a boolean if a field has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) HasRedactEntireValueFieldName() bool {
	if o != nil && !isNil(o.RedactEntireValueFieldName) {
		return true
	}

	return false
}

// SetRedactEntireValueFieldName gets a reference to the given []string and assigns it to the RedactEntireValueFieldName field.
func (o *AddTextAccessLogFieldBehaviorRequest) SetRedactEntireValueFieldName(v []string) {
	o.RedactEntireValueFieldName = v
}

// GetRedactValueComponentsField returns the RedactValueComponentsField field value if set, zero value otherwise.
func (o *AddTextAccessLogFieldBehaviorRequest) GetRedactValueComponentsField() []EnumlogFieldBehaviorRedactValueComponentsFieldProp {
	if o == nil || isNil(o.RedactValueComponentsField) {
		var ret []EnumlogFieldBehaviorRedactValueComponentsFieldProp
		return ret
	}
	return o.RedactValueComponentsField
}

// GetRedactValueComponentsFieldOk returns a tuple with the RedactValueComponentsField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) GetRedactValueComponentsFieldOk() ([]EnumlogFieldBehaviorRedactValueComponentsFieldProp, bool) {
	if o == nil || isNil(o.RedactValueComponentsField) {
		return nil, false
	}
	return o.RedactValueComponentsField, true
}

// HasRedactValueComponentsField returns a boolean if a field has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) HasRedactValueComponentsField() bool {
	if o != nil && !isNil(o.RedactValueComponentsField) {
		return true
	}

	return false
}

// SetRedactValueComponentsField gets a reference to the given []EnumlogFieldBehaviorRedactValueComponentsFieldProp and assigns it to the RedactValueComponentsField field.
func (o *AddTextAccessLogFieldBehaviorRequest) SetRedactValueComponentsField(v []EnumlogFieldBehaviorRedactValueComponentsFieldProp) {
	o.RedactValueComponentsField = v
}

// GetRedactValueComponentsFieldName returns the RedactValueComponentsFieldName field value if set, zero value otherwise.
func (o *AddTextAccessLogFieldBehaviorRequest) GetRedactValueComponentsFieldName() []string {
	if o == nil || isNil(o.RedactValueComponentsFieldName) {
		var ret []string
		return ret
	}
	return o.RedactValueComponentsFieldName
}

// GetRedactValueComponentsFieldNameOk returns a tuple with the RedactValueComponentsFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) GetRedactValueComponentsFieldNameOk() ([]string, bool) {
	if o == nil || isNil(o.RedactValueComponentsFieldName) {
		return nil, false
	}
	return o.RedactValueComponentsFieldName, true
}

// HasRedactValueComponentsFieldName returns a boolean if a field has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) HasRedactValueComponentsFieldName() bool {
	if o != nil && !isNil(o.RedactValueComponentsFieldName) {
		return true
	}

	return false
}

// SetRedactValueComponentsFieldName gets a reference to the given []string and assigns it to the RedactValueComponentsFieldName field.
func (o *AddTextAccessLogFieldBehaviorRequest) SetRedactValueComponentsFieldName(v []string) {
	o.RedactValueComponentsFieldName = v
}

// GetTokenizeEntireValueField returns the TokenizeEntireValueField field value if set, zero value otherwise.
func (o *AddTextAccessLogFieldBehaviorRequest) GetTokenizeEntireValueField() []EnumlogFieldBehaviorTokenizeEntireValueFieldProp {
	if o == nil || isNil(o.TokenizeEntireValueField) {
		var ret []EnumlogFieldBehaviorTokenizeEntireValueFieldProp
		return ret
	}
	return o.TokenizeEntireValueField
}

// GetTokenizeEntireValueFieldOk returns a tuple with the TokenizeEntireValueField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) GetTokenizeEntireValueFieldOk() ([]EnumlogFieldBehaviorTokenizeEntireValueFieldProp, bool) {
	if o == nil || isNil(o.TokenizeEntireValueField) {
		return nil, false
	}
	return o.TokenizeEntireValueField, true
}

// HasTokenizeEntireValueField returns a boolean if a field has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) HasTokenizeEntireValueField() bool {
	if o != nil && !isNil(o.TokenizeEntireValueField) {
		return true
	}

	return false
}

// SetTokenizeEntireValueField gets a reference to the given []EnumlogFieldBehaviorTokenizeEntireValueFieldProp and assigns it to the TokenizeEntireValueField field.
func (o *AddTextAccessLogFieldBehaviorRequest) SetTokenizeEntireValueField(v []EnumlogFieldBehaviorTokenizeEntireValueFieldProp) {
	o.TokenizeEntireValueField = v
}

// GetTokenizeEntireValueFieldName returns the TokenizeEntireValueFieldName field value if set, zero value otherwise.
func (o *AddTextAccessLogFieldBehaviorRequest) GetTokenizeEntireValueFieldName() []string {
	if o == nil || isNil(o.TokenizeEntireValueFieldName) {
		var ret []string
		return ret
	}
	return o.TokenizeEntireValueFieldName
}

// GetTokenizeEntireValueFieldNameOk returns a tuple with the TokenizeEntireValueFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) GetTokenizeEntireValueFieldNameOk() ([]string, bool) {
	if o == nil || isNil(o.TokenizeEntireValueFieldName) {
		return nil, false
	}
	return o.TokenizeEntireValueFieldName, true
}

// HasTokenizeEntireValueFieldName returns a boolean if a field has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) HasTokenizeEntireValueFieldName() bool {
	if o != nil && !isNil(o.TokenizeEntireValueFieldName) {
		return true
	}

	return false
}

// SetTokenizeEntireValueFieldName gets a reference to the given []string and assigns it to the TokenizeEntireValueFieldName field.
func (o *AddTextAccessLogFieldBehaviorRequest) SetTokenizeEntireValueFieldName(v []string) {
	o.TokenizeEntireValueFieldName = v
}

// GetTokenizeValueComponentsField returns the TokenizeValueComponentsField field value if set, zero value otherwise.
func (o *AddTextAccessLogFieldBehaviorRequest) GetTokenizeValueComponentsField() []EnumlogFieldBehaviorTokenizeValueComponentsFieldProp {
	if o == nil || isNil(o.TokenizeValueComponentsField) {
		var ret []EnumlogFieldBehaviorTokenizeValueComponentsFieldProp
		return ret
	}
	return o.TokenizeValueComponentsField
}

// GetTokenizeValueComponentsFieldOk returns a tuple with the TokenizeValueComponentsField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) GetTokenizeValueComponentsFieldOk() ([]EnumlogFieldBehaviorTokenizeValueComponentsFieldProp, bool) {
	if o == nil || isNil(o.TokenizeValueComponentsField) {
		return nil, false
	}
	return o.TokenizeValueComponentsField, true
}

// HasTokenizeValueComponentsField returns a boolean if a field has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) HasTokenizeValueComponentsField() bool {
	if o != nil && !isNil(o.TokenizeValueComponentsField) {
		return true
	}

	return false
}

// SetTokenizeValueComponentsField gets a reference to the given []EnumlogFieldBehaviorTokenizeValueComponentsFieldProp and assigns it to the TokenizeValueComponentsField field.
func (o *AddTextAccessLogFieldBehaviorRequest) SetTokenizeValueComponentsField(v []EnumlogFieldBehaviorTokenizeValueComponentsFieldProp) {
	o.TokenizeValueComponentsField = v
}

// GetTokenizeValueComponentsFieldName returns the TokenizeValueComponentsFieldName field value if set, zero value otherwise.
func (o *AddTextAccessLogFieldBehaviorRequest) GetTokenizeValueComponentsFieldName() []string {
	if o == nil || isNil(o.TokenizeValueComponentsFieldName) {
		var ret []string
		return ret
	}
	return o.TokenizeValueComponentsFieldName
}

// GetTokenizeValueComponentsFieldNameOk returns a tuple with the TokenizeValueComponentsFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) GetTokenizeValueComponentsFieldNameOk() ([]string, bool) {
	if o == nil || isNil(o.TokenizeValueComponentsFieldName) {
		return nil, false
	}
	return o.TokenizeValueComponentsFieldName, true
}

// HasTokenizeValueComponentsFieldName returns a boolean if a field has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) HasTokenizeValueComponentsFieldName() bool {
	if o != nil && !isNil(o.TokenizeValueComponentsFieldName) {
		return true
	}

	return false
}

// SetTokenizeValueComponentsFieldName gets a reference to the given []string and assigns it to the TokenizeValueComponentsFieldName field.
func (o *AddTextAccessLogFieldBehaviorRequest) SetTokenizeValueComponentsFieldName(v []string) {
	o.TokenizeValueComponentsFieldName = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddTextAccessLogFieldBehaviorRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddTextAccessLogFieldBehaviorRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDefaultBehavior returns the DefaultBehavior field value if set, zero value otherwise.
func (o *AddTextAccessLogFieldBehaviorRequest) GetDefaultBehavior() EnumlogFieldBehaviorDefaultBehaviorProp {
	if o == nil || isNil(o.DefaultBehavior) {
		var ret EnumlogFieldBehaviorDefaultBehaviorProp
		return ret
	}
	return *o.DefaultBehavior
}

// GetDefaultBehaviorOk returns a tuple with the DefaultBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) GetDefaultBehaviorOk() (*EnumlogFieldBehaviorDefaultBehaviorProp, bool) {
	if o == nil || isNil(o.DefaultBehavior) {
		return nil, false
	}
	return o.DefaultBehavior, true
}

// HasDefaultBehavior returns a boolean if a field has been set.
func (o *AddTextAccessLogFieldBehaviorRequest) HasDefaultBehavior() bool {
	if o != nil && !isNil(o.DefaultBehavior) {
		return true
	}

	return false
}

// SetDefaultBehavior gets a reference to the given EnumlogFieldBehaviorDefaultBehaviorProp and assigns it to the DefaultBehavior field.
func (o *AddTextAccessLogFieldBehaviorRequest) SetDefaultBehavior(v EnumlogFieldBehaviorDefaultBehaviorProp) {
	o.DefaultBehavior = &v
}

func (o AddTextAccessLogFieldBehaviorRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["behaviorName"] = o.BehaviorName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.PreserveField) {
		toSerialize["preserveField"] = o.PreserveField
	}
	if !isNil(o.PreserveFieldName) {
		toSerialize["preserveFieldName"] = o.PreserveFieldName
	}
	if !isNil(o.OmitField) {
		toSerialize["omitField"] = o.OmitField
	}
	if !isNil(o.OmitFieldName) {
		toSerialize["omitFieldName"] = o.OmitFieldName
	}
	if !isNil(o.RedactEntireValueField) {
		toSerialize["redactEntireValueField"] = o.RedactEntireValueField
	}
	if !isNil(o.RedactEntireValueFieldName) {
		toSerialize["redactEntireValueFieldName"] = o.RedactEntireValueFieldName
	}
	if !isNil(o.RedactValueComponentsField) {
		toSerialize["redactValueComponentsField"] = o.RedactValueComponentsField
	}
	if !isNil(o.RedactValueComponentsFieldName) {
		toSerialize["redactValueComponentsFieldName"] = o.RedactValueComponentsFieldName
	}
	if !isNil(o.TokenizeEntireValueField) {
		toSerialize["tokenizeEntireValueField"] = o.TokenizeEntireValueField
	}
	if !isNil(o.TokenizeEntireValueFieldName) {
		toSerialize["tokenizeEntireValueFieldName"] = o.TokenizeEntireValueFieldName
	}
	if !isNil(o.TokenizeValueComponentsField) {
		toSerialize["tokenizeValueComponentsField"] = o.TokenizeValueComponentsField
	}
	if !isNil(o.TokenizeValueComponentsFieldName) {
		toSerialize["tokenizeValueComponentsFieldName"] = o.TokenizeValueComponentsFieldName
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.DefaultBehavior) {
		toSerialize["defaultBehavior"] = o.DefaultBehavior
	}
	return json.Marshal(toSerialize)
}

type NullableAddTextAccessLogFieldBehaviorRequest struct {
	value *AddTextAccessLogFieldBehaviorRequest
	isSet bool
}

func (v NullableAddTextAccessLogFieldBehaviorRequest) Get() *AddTextAccessLogFieldBehaviorRequest {
	return v.value
}

func (v *NullableAddTextAccessLogFieldBehaviorRequest) Set(val *AddTextAccessLogFieldBehaviorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddTextAccessLogFieldBehaviorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddTextAccessLogFieldBehaviorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddTextAccessLogFieldBehaviorRequest(val *AddTextAccessLogFieldBehaviorRequest) *NullableAddTextAccessLogFieldBehaviorRequest {
	return &NullableAddTextAccessLogFieldBehaviorRequest{value: val, isSet: true}
}

func (v NullableAddTextAccessLogFieldBehaviorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddTextAccessLogFieldBehaviorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
