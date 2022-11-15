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

// JsonFormattedAccessLogFieldBehaviorShared struct for JsonFormattedAccessLogFieldBehaviorShared
type JsonFormattedAccessLogFieldBehaviorShared struct {
	Schemas []EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn `json:"schemas"`
	PreserveField []EnumlogFieldBehaviorPreserveFieldProp `json:"preserveField,omitempty"`
	PreserveFieldName []string `json:"preserveFieldName,omitempty"`
	OmitField []EnumlogFieldBehaviorOmitFieldProp `json:"omitField,omitempty"`
	OmitFieldName []string `json:"omitFieldName,omitempty"`
	RedactEntireValueField []EnumlogFieldBehaviorRedactEntireValueFieldProp `json:"redactEntireValueField,omitempty"`
	RedactEntireValueFieldName []string `json:"redactEntireValueFieldName,omitempty"`
	RedactValueComponentsField []EnumlogFieldBehaviorRedactValueComponentsFieldProp `json:"redactValueComponentsField,omitempty"`
	RedactValueComponentsFieldName []string `json:"redactValueComponentsFieldName,omitempty"`
	TokenizeEntireValueField []EnumlogFieldBehaviorTokenizeEntireValueFieldProp `json:"tokenizeEntireValueField,omitempty"`
	TokenizeEntireValueFieldName []string `json:"tokenizeEntireValueFieldName,omitempty"`
	TokenizeValueComponentsField []EnumlogFieldBehaviorTokenizeValueComponentsFieldProp `json:"tokenizeValueComponentsField,omitempty"`
	TokenizeValueComponentsFieldName []string `json:"tokenizeValueComponentsFieldName,omitempty"`
	// A description for this Log Field Behavior
	Description *string `json:"description,omitempty"`
	DefaultBehavior *EnumlogFieldBehaviorDefaultBehaviorProp `json:"defaultBehavior,omitempty"`
}

// NewJsonFormattedAccessLogFieldBehaviorShared instantiates a new JsonFormattedAccessLogFieldBehaviorShared object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonFormattedAccessLogFieldBehaviorShared(schemas []EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn) *JsonFormattedAccessLogFieldBehaviorShared {
	this := JsonFormattedAccessLogFieldBehaviorShared{}
	this.Schemas = schemas
	return &this
}

// NewJsonFormattedAccessLogFieldBehaviorSharedWithDefaults instantiates a new JsonFormattedAccessLogFieldBehaviorShared object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonFormattedAccessLogFieldBehaviorSharedWithDefaults() *JsonFormattedAccessLogFieldBehaviorShared {
	this := JsonFormattedAccessLogFieldBehaviorShared{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetSchemas() []EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn {
	if o == nil {
		var ret []EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetSchemasOk() ([]EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *JsonFormattedAccessLogFieldBehaviorShared) SetSchemas(v []EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn) {
	o.Schemas = v
}

// GetPreserveField returns the PreserveField field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetPreserveField() []EnumlogFieldBehaviorPreserveFieldProp {
	if o == nil || isNil(o.PreserveField) {
		var ret []EnumlogFieldBehaviorPreserveFieldProp
		return ret
	}
	return o.PreserveField
}

// GetPreserveFieldOk returns a tuple with the PreserveField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetPreserveFieldOk() ([]EnumlogFieldBehaviorPreserveFieldProp, bool) {
	if o == nil || isNil(o.PreserveField) {
    return nil, false
	}
	return o.PreserveField, true
}

// HasPreserveField returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorShared) HasPreserveField() bool {
	if o != nil && !isNil(o.PreserveField) {
		return true
	}

	return false
}

// SetPreserveField gets a reference to the given []EnumlogFieldBehaviorPreserveFieldProp and assigns it to the PreserveField field.
func (o *JsonFormattedAccessLogFieldBehaviorShared) SetPreserveField(v []EnumlogFieldBehaviorPreserveFieldProp) {
	o.PreserveField = v
}

// GetPreserveFieldName returns the PreserveFieldName field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetPreserveFieldName() []string {
	if o == nil || isNil(o.PreserveFieldName) {
		var ret []string
		return ret
	}
	return o.PreserveFieldName
}

// GetPreserveFieldNameOk returns a tuple with the PreserveFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetPreserveFieldNameOk() ([]string, bool) {
	if o == nil || isNil(o.PreserveFieldName) {
    return nil, false
	}
	return o.PreserveFieldName, true
}

// HasPreserveFieldName returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorShared) HasPreserveFieldName() bool {
	if o != nil && !isNil(o.PreserveFieldName) {
		return true
	}

	return false
}

// SetPreserveFieldName gets a reference to the given []string and assigns it to the PreserveFieldName field.
func (o *JsonFormattedAccessLogFieldBehaviorShared) SetPreserveFieldName(v []string) {
	o.PreserveFieldName = v
}

// GetOmitField returns the OmitField field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetOmitField() []EnumlogFieldBehaviorOmitFieldProp {
	if o == nil || isNil(o.OmitField) {
		var ret []EnumlogFieldBehaviorOmitFieldProp
		return ret
	}
	return o.OmitField
}

// GetOmitFieldOk returns a tuple with the OmitField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetOmitFieldOk() ([]EnumlogFieldBehaviorOmitFieldProp, bool) {
	if o == nil || isNil(o.OmitField) {
    return nil, false
	}
	return o.OmitField, true
}

// HasOmitField returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorShared) HasOmitField() bool {
	if o != nil && !isNil(o.OmitField) {
		return true
	}

	return false
}

// SetOmitField gets a reference to the given []EnumlogFieldBehaviorOmitFieldProp and assigns it to the OmitField field.
func (o *JsonFormattedAccessLogFieldBehaviorShared) SetOmitField(v []EnumlogFieldBehaviorOmitFieldProp) {
	o.OmitField = v
}

// GetOmitFieldName returns the OmitFieldName field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetOmitFieldName() []string {
	if o == nil || isNil(o.OmitFieldName) {
		var ret []string
		return ret
	}
	return o.OmitFieldName
}

// GetOmitFieldNameOk returns a tuple with the OmitFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetOmitFieldNameOk() ([]string, bool) {
	if o == nil || isNil(o.OmitFieldName) {
    return nil, false
	}
	return o.OmitFieldName, true
}

// HasOmitFieldName returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorShared) HasOmitFieldName() bool {
	if o != nil && !isNil(o.OmitFieldName) {
		return true
	}

	return false
}

// SetOmitFieldName gets a reference to the given []string and assigns it to the OmitFieldName field.
func (o *JsonFormattedAccessLogFieldBehaviorShared) SetOmitFieldName(v []string) {
	o.OmitFieldName = v
}

// GetRedactEntireValueField returns the RedactEntireValueField field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetRedactEntireValueField() []EnumlogFieldBehaviorRedactEntireValueFieldProp {
	if o == nil || isNil(o.RedactEntireValueField) {
		var ret []EnumlogFieldBehaviorRedactEntireValueFieldProp
		return ret
	}
	return o.RedactEntireValueField
}

// GetRedactEntireValueFieldOk returns a tuple with the RedactEntireValueField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetRedactEntireValueFieldOk() ([]EnumlogFieldBehaviorRedactEntireValueFieldProp, bool) {
	if o == nil || isNil(o.RedactEntireValueField) {
    return nil, false
	}
	return o.RedactEntireValueField, true
}

// HasRedactEntireValueField returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorShared) HasRedactEntireValueField() bool {
	if o != nil && !isNil(o.RedactEntireValueField) {
		return true
	}

	return false
}

// SetRedactEntireValueField gets a reference to the given []EnumlogFieldBehaviorRedactEntireValueFieldProp and assigns it to the RedactEntireValueField field.
func (o *JsonFormattedAccessLogFieldBehaviorShared) SetRedactEntireValueField(v []EnumlogFieldBehaviorRedactEntireValueFieldProp) {
	o.RedactEntireValueField = v
}

// GetRedactEntireValueFieldName returns the RedactEntireValueFieldName field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetRedactEntireValueFieldName() []string {
	if o == nil || isNil(o.RedactEntireValueFieldName) {
		var ret []string
		return ret
	}
	return o.RedactEntireValueFieldName
}

// GetRedactEntireValueFieldNameOk returns a tuple with the RedactEntireValueFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetRedactEntireValueFieldNameOk() ([]string, bool) {
	if o == nil || isNil(o.RedactEntireValueFieldName) {
    return nil, false
	}
	return o.RedactEntireValueFieldName, true
}

// HasRedactEntireValueFieldName returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorShared) HasRedactEntireValueFieldName() bool {
	if o != nil && !isNil(o.RedactEntireValueFieldName) {
		return true
	}

	return false
}

// SetRedactEntireValueFieldName gets a reference to the given []string and assigns it to the RedactEntireValueFieldName field.
func (o *JsonFormattedAccessLogFieldBehaviorShared) SetRedactEntireValueFieldName(v []string) {
	o.RedactEntireValueFieldName = v
}

// GetRedactValueComponentsField returns the RedactValueComponentsField field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetRedactValueComponentsField() []EnumlogFieldBehaviorRedactValueComponentsFieldProp {
	if o == nil || isNil(o.RedactValueComponentsField) {
		var ret []EnumlogFieldBehaviorRedactValueComponentsFieldProp
		return ret
	}
	return o.RedactValueComponentsField
}

// GetRedactValueComponentsFieldOk returns a tuple with the RedactValueComponentsField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetRedactValueComponentsFieldOk() ([]EnumlogFieldBehaviorRedactValueComponentsFieldProp, bool) {
	if o == nil || isNil(o.RedactValueComponentsField) {
    return nil, false
	}
	return o.RedactValueComponentsField, true
}

// HasRedactValueComponentsField returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorShared) HasRedactValueComponentsField() bool {
	if o != nil && !isNil(o.RedactValueComponentsField) {
		return true
	}

	return false
}

// SetRedactValueComponentsField gets a reference to the given []EnumlogFieldBehaviorRedactValueComponentsFieldProp and assigns it to the RedactValueComponentsField field.
func (o *JsonFormattedAccessLogFieldBehaviorShared) SetRedactValueComponentsField(v []EnumlogFieldBehaviorRedactValueComponentsFieldProp) {
	o.RedactValueComponentsField = v
}

// GetRedactValueComponentsFieldName returns the RedactValueComponentsFieldName field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetRedactValueComponentsFieldName() []string {
	if o == nil || isNil(o.RedactValueComponentsFieldName) {
		var ret []string
		return ret
	}
	return o.RedactValueComponentsFieldName
}

// GetRedactValueComponentsFieldNameOk returns a tuple with the RedactValueComponentsFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetRedactValueComponentsFieldNameOk() ([]string, bool) {
	if o == nil || isNil(o.RedactValueComponentsFieldName) {
    return nil, false
	}
	return o.RedactValueComponentsFieldName, true
}

// HasRedactValueComponentsFieldName returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorShared) HasRedactValueComponentsFieldName() bool {
	if o != nil && !isNil(o.RedactValueComponentsFieldName) {
		return true
	}

	return false
}

// SetRedactValueComponentsFieldName gets a reference to the given []string and assigns it to the RedactValueComponentsFieldName field.
func (o *JsonFormattedAccessLogFieldBehaviorShared) SetRedactValueComponentsFieldName(v []string) {
	o.RedactValueComponentsFieldName = v
}

// GetTokenizeEntireValueField returns the TokenizeEntireValueField field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetTokenizeEntireValueField() []EnumlogFieldBehaviorTokenizeEntireValueFieldProp {
	if o == nil || isNil(o.TokenizeEntireValueField) {
		var ret []EnumlogFieldBehaviorTokenizeEntireValueFieldProp
		return ret
	}
	return o.TokenizeEntireValueField
}

// GetTokenizeEntireValueFieldOk returns a tuple with the TokenizeEntireValueField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetTokenizeEntireValueFieldOk() ([]EnumlogFieldBehaviorTokenizeEntireValueFieldProp, bool) {
	if o == nil || isNil(o.TokenizeEntireValueField) {
    return nil, false
	}
	return o.TokenizeEntireValueField, true
}

// HasTokenizeEntireValueField returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorShared) HasTokenizeEntireValueField() bool {
	if o != nil && !isNil(o.TokenizeEntireValueField) {
		return true
	}

	return false
}

// SetTokenizeEntireValueField gets a reference to the given []EnumlogFieldBehaviorTokenizeEntireValueFieldProp and assigns it to the TokenizeEntireValueField field.
func (o *JsonFormattedAccessLogFieldBehaviorShared) SetTokenizeEntireValueField(v []EnumlogFieldBehaviorTokenizeEntireValueFieldProp) {
	o.TokenizeEntireValueField = v
}

// GetTokenizeEntireValueFieldName returns the TokenizeEntireValueFieldName field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetTokenizeEntireValueFieldName() []string {
	if o == nil || isNil(o.TokenizeEntireValueFieldName) {
		var ret []string
		return ret
	}
	return o.TokenizeEntireValueFieldName
}

// GetTokenizeEntireValueFieldNameOk returns a tuple with the TokenizeEntireValueFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetTokenizeEntireValueFieldNameOk() ([]string, bool) {
	if o == nil || isNil(o.TokenizeEntireValueFieldName) {
    return nil, false
	}
	return o.TokenizeEntireValueFieldName, true
}

// HasTokenizeEntireValueFieldName returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorShared) HasTokenizeEntireValueFieldName() bool {
	if o != nil && !isNil(o.TokenizeEntireValueFieldName) {
		return true
	}

	return false
}

// SetTokenizeEntireValueFieldName gets a reference to the given []string and assigns it to the TokenizeEntireValueFieldName field.
func (o *JsonFormattedAccessLogFieldBehaviorShared) SetTokenizeEntireValueFieldName(v []string) {
	o.TokenizeEntireValueFieldName = v
}

// GetTokenizeValueComponentsField returns the TokenizeValueComponentsField field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetTokenizeValueComponentsField() []EnumlogFieldBehaviorTokenizeValueComponentsFieldProp {
	if o == nil || isNil(o.TokenizeValueComponentsField) {
		var ret []EnumlogFieldBehaviorTokenizeValueComponentsFieldProp
		return ret
	}
	return o.TokenizeValueComponentsField
}

// GetTokenizeValueComponentsFieldOk returns a tuple with the TokenizeValueComponentsField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetTokenizeValueComponentsFieldOk() ([]EnumlogFieldBehaviorTokenizeValueComponentsFieldProp, bool) {
	if o == nil || isNil(o.TokenizeValueComponentsField) {
    return nil, false
	}
	return o.TokenizeValueComponentsField, true
}

// HasTokenizeValueComponentsField returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorShared) HasTokenizeValueComponentsField() bool {
	if o != nil && !isNil(o.TokenizeValueComponentsField) {
		return true
	}

	return false
}

// SetTokenizeValueComponentsField gets a reference to the given []EnumlogFieldBehaviorTokenizeValueComponentsFieldProp and assigns it to the TokenizeValueComponentsField field.
func (o *JsonFormattedAccessLogFieldBehaviorShared) SetTokenizeValueComponentsField(v []EnumlogFieldBehaviorTokenizeValueComponentsFieldProp) {
	o.TokenizeValueComponentsField = v
}

// GetTokenizeValueComponentsFieldName returns the TokenizeValueComponentsFieldName field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetTokenizeValueComponentsFieldName() []string {
	if o == nil || isNil(o.TokenizeValueComponentsFieldName) {
		var ret []string
		return ret
	}
	return o.TokenizeValueComponentsFieldName
}

// GetTokenizeValueComponentsFieldNameOk returns a tuple with the TokenizeValueComponentsFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetTokenizeValueComponentsFieldNameOk() ([]string, bool) {
	if o == nil || isNil(o.TokenizeValueComponentsFieldName) {
    return nil, false
	}
	return o.TokenizeValueComponentsFieldName, true
}

// HasTokenizeValueComponentsFieldName returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorShared) HasTokenizeValueComponentsFieldName() bool {
	if o != nil && !isNil(o.TokenizeValueComponentsFieldName) {
		return true
	}

	return false
}

// SetTokenizeValueComponentsFieldName gets a reference to the given []string and assigns it to the TokenizeValueComponentsFieldName field.
func (o *JsonFormattedAccessLogFieldBehaviorShared) SetTokenizeValueComponentsFieldName(v []string) {
	o.TokenizeValueComponentsFieldName = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorShared) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *JsonFormattedAccessLogFieldBehaviorShared) SetDescription(v string) {
	o.Description = &v
}

// GetDefaultBehavior returns the DefaultBehavior field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetDefaultBehavior() EnumlogFieldBehaviorDefaultBehaviorProp {
	if o == nil || isNil(o.DefaultBehavior) {
		var ret EnumlogFieldBehaviorDefaultBehaviorProp
		return ret
	}
	return *o.DefaultBehavior
}

// GetDefaultBehaviorOk returns a tuple with the DefaultBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorShared) GetDefaultBehaviorOk() (*EnumlogFieldBehaviorDefaultBehaviorProp, bool) {
	if o == nil || isNil(o.DefaultBehavior) {
    return nil, false
	}
	return o.DefaultBehavior, true
}

// HasDefaultBehavior returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorShared) HasDefaultBehavior() bool {
	if o != nil && !isNil(o.DefaultBehavior) {
		return true
	}

	return false
}

// SetDefaultBehavior gets a reference to the given EnumlogFieldBehaviorDefaultBehaviorProp and assigns it to the DefaultBehavior field.
func (o *JsonFormattedAccessLogFieldBehaviorShared) SetDefaultBehavior(v EnumlogFieldBehaviorDefaultBehaviorProp) {
	o.DefaultBehavior = &v
}

func (o JsonFormattedAccessLogFieldBehaviorShared) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableJsonFormattedAccessLogFieldBehaviorShared struct {
	value *JsonFormattedAccessLogFieldBehaviorShared
	isSet bool
}

func (v NullableJsonFormattedAccessLogFieldBehaviorShared) Get() *JsonFormattedAccessLogFieldBehaviorShared {
	return v.value
}

func (v *NullableJsonFormattedAccessLogFieldBehaviorShared) Set(val *JsonFormattedAccessLogFieldBehaviorShared) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonFormattedAccessLogFieldBehaviorShared) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonFormattedAccessLogFieldBehaviorShared) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonFormattedAccessLogFieldBehaviorShared(val *JsonFormattedAccessLogFieldBehaviorShared) *NullableJsonFormattedAccessLogFieldBehaviorShared {
	return &NullableJsonFormattedAccessLogFieldBehaviorShared{value: val, isSet: true}
}

func (v NullableJsonFormattedAccessLogFieldBehaviorShared) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonFormattedAccessLogFieldBehaviorShared) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


