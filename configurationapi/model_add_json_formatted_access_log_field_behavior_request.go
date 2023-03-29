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

// checks if the AddJsonFormattedAccessLogFieldBehaviorRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddJsonFormattedAccessLogFieldBehaviorRequest{}

// AddJsonFormattedAccessLogFieldBehaviorRequest struct for AddJsonFormattedAccessLogFieldBehaviorRequest
type AddJsonFormattedAccessLogFieldBehaviorRequest struct {
	// Name of the new Log Field Behavior
	BehaviorName string                                             `json:"behaviorName"`
	Schemas      []EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn `json:"schemas"`
	// The log fields whose values should be logged with the intended value. The values for these fields will be preserved, although they may be sanitized for parsability or safety purposes (for example, to escape special characters in the value), and values that are too long may be truncated.
	PreserveField []EnumlogFieldBehaviorJsonFormattedAccessPreserveFieldProp `json:"preserveField,omitempty"`
	// The names of any custom fields whose values should be preserved. This should generally only be used for fields that are not available through the preserve-field property (for example, custom log fields defined in Server SDK extensions).
	PreserveFieldName []string `json:"preserveFieldName,omitempty"`
	// The log fields that should be omitted entirely from log messages. Neither the field name nor value will be included.
	OmitField []EnumlogFieldBehaviorJsonFormattedAccessOmitFieldProp `json:"omitField,omitempty"`
	// The names of any custom fields that should be omitted from log messages. This should generally only be used for fields that are not available through the omit-field property (for example, custom log fields defined in Server SDK extensions).
	OmitFieldName []string `json:"omitFieldName,omitempty"`
	// The log fields whose values should be completely redacted in log messages. The field name will be included, but with a fixed value that does not reflect the actual value for the field.
	RedactEntireValueField []EnumlogFieldBehaviorJsonFormattedAccessRedactEntireValueFieldProp `json:"redactEntireValueField,omitempty"`
	// The names of any custom fields whose values should be completely redacted. This should generally only be used for fields that are not available through the redact-entire-value-field property (for example, custom log fields defined in Server SDK extensions).
	RedactEntireValueFieldName []string `json:"redactEntireValueFieldName,omitempty"`
	// The log fields whose values will include redacted components.
	RedactValueComponentsField []EnumlogFieldBehaviorJsonFormattedAccessRedactValueComponentsFieldProp `json:"redactValueComponentsField,omitempty"`
	// The names of any custom fields for which to redact components within the value. This should generally only be used for fields that are not available through the redact-value-components-field property (for example, custom log fields defined in Server SDK extensions).
	RedactValueComponentsFieldName []string `json:"redactValueComponentsFieldName,omitempty"`
	// The log fields whose values should be completely tokenized in log messages. The field name will be included, but the value will be replaced with a token that does not reveal the actual value, but that is generated from the value.
	TokenizeEntireValueField []EnumlogFieldBehaviorJsonFormattedAccessTokenizeEntireValueFieldProp `json:"tokenizeEntireValueField,omitempty"`
	// The names of any custom fields whose values should be completely tokenized. This should generally only be used for fields that are not available through the tokenize-entire-value-field property (for example, custom log fields defined in Server SDK extensions).
	TokenizeEntireValueFieldName []string `json:"tokenizeEntireValueFieldName,omitempty"`
	// The log fields whose values will include tokenized components.
	TokenizeValueComponentsField []EnumlogFieldBehaviorJsonFormattedAccessTokenizeValueComponentsFieldProp `json:"tokenizeValueComponentsField,omitempty"`
	// The names of any custom fields for which to tokenize components within the value. This should generally only be used for fields that are not available through the tokenize-value-components-field property (for example, custom log fields defined in Server SDK extensions).
	TokenizeValueComponentsFieldName []string `json:"tokenizeValueComponentsFieldName,omitempty"`
	// A description for this Log Field Behavior
	Description     *string                                  `json:"description,omitempty"`
	DefaultBehavior *EnumlogFieldBehaviorDefaultBehaviorProp `json:"defaultBehavior,omitempty"`
}

// NewAddJsonFormattedAccessLogFieldBehaviorRequest instantiates a new AddJsonFormattedAccessLogFieldBehaviorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddJsonFormattedAccessLogFieldBehaviorRequest(behaviorName string, schemas []EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn) *AddJsonFormattedAccessLogFieldBehaviorRequest {
	this := AddJsonFormattedAccessLogFieldBehaviorRequest{}
	this.BehaviorName = behaviorName
	this.Schemas = schemas
	return &this
}

// NewAddJsonFormattedAccessLogFieldBehaviorRequestWithDefaults instantiates a new AddJsonFormattedAccessLogFieldBehaviorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddJsonFormattedAccessLogFieldBehaviorRequestWithDefaults() *AddJsonFormattedAccessLogFieldBehaviorRequest {
	this := AddJsonFormattedAccessLogFieldBehaviorRequest{}
	return &this
}

// GetBehaviorName returns the BehaviorName field value
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetBehaviorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BehaviorName
}

// GetBehaviorNameOk returns a tuple with the BehaviorName field value
// and a boolean to check if the value has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetBehaviorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BehaviorName, true
}

// SetBehaviorName sets field value
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetBehaviorName(v string) {
	o.BehaviorName = v
}

// GetSchemas returns the Schemas field value
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetSchemas() []EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn {
	if o == nil {
		var ret []EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetSchemasOk() ([]EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetSchemas(v []EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn) {
	o.Schemas = v
}

// GetPreserveField returns the PreserveField field value if set, zero value otherwise.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetPreserveField() []EnumlogFieldBehaviorJsonFormattedAccessPreserveFieldProp {
	if o == nil || IsNil(o.PreserveField) {
		var ret []EnumlogFieldBehaviorJsonFormattedAccessPreserveFieldProp
		return ret
	}
	return o.PreserveField
}

// GetPreserveFieldOk returns a tuple with the PreserveField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetPreserveFieldOk() ([]EnumlogFieldBehaviorJsonFormattedAccessPreserveFieldProp, bool) {
	if o == nil || IsNil(o.PreserveField) {
		return nil, false
	}
	return o.PreserveField, true
}

// HasPreserveField returns a boolean if a field has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) HasPreserveField() bool {
	if o != nil && !IsNil(o.PreserveField) {
		return true
	}

	return false
}

// SetPreserveField gets a reference to the given []EnumlogFieldBehaviorJsonFormattedAccessPreserveFieldProp and assigns it to the PreserveField field.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetPreserveField(v []EnumlogFieldBehaviorJsonFormattedAccessPreserveFieldProp) {
	o.PreserveField = v
}

// GetPreserveFieldName returns the PreserveFieldName field value if set, zero value otherwise.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetPreserveFieldName() []string {
	if o == nil || IsNil(o.PreserveFieldName) {
		var ret []string
		return ret
	}
	return o.PreserveFieldName
}

// GetPreserveFieldNameOk returns a tuple with the PreserveFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetPreserveFieldNameOk() ([]string, bool) {
	if o == nil || IsNil(o.PreserveFieldName) {
		return nil, false
	}
	return o.PreserveFieldName, true
}

// HasPreserveFieldName returns a boolean if a field has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) HasPreserveFieldName() bool {
	if o != nil && !IsNil(o.PreserveFieldName) {
		return true
	}

	return false
}

// SetPreserveFieldName gets a reference to the given []string and assigns it to the PreserveFieldName field.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetPreserveFieldName(v []string) {
	o.PreserveFieldName = v
}

// GetOmitField returns the OmitField field value if set, zero value otherwise.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetOmitField() []EnumlogFieldBehaviorJsonFormattedAccessOmitFieldProp {
	if o == nil || IsNil(o.OmitField) {
		var ret []EnumlogFieldBehaviorJsonFormattedAccessOmitFieldProp
		return ret
	}
	return o.OmitField
}

// GetOmitFieldOk returns a tuple with the OmitField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetOmitFieldOk() ([]EnumlogFieldBehaviorJsonFormattedAccessOmitFieldProp, bool) {
	if o == nil || IsNil(o.OmitField) {
		return nil, false
	}
	return o.OmitField, true
}

// HasOmitField returns a boolean if a field has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) HasOmitField() bool {
	if o != nil && !IsNil(o.OmitField) {
		return true
	}

	return false
}

// SetOmitField gets a reference to the given []EnumlogFieldBehaviorJsonFormattedAccessOmitFieldProp and assigns it to the OmitField field.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetOmitField(v []EnumlogFieldBehaviorJsonFormattedAccessOmitFieldProp) {
	o.OmitField = v
}

// GetOmitFieldName returns the OmitFieldName field value if set, zero value otherwise.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetOmitFieldName() []string {
	if o == nil || IsNil(o.OmitFieldName) {
		var ret []string
		return ret
	}
	return o.OmitFieldName
}

// GetOmitFieldNameOk returns a tuple with the OmitFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetOmitFieldNameOk() ([]string, bool) {
	if o == nil || IsNil(o.OmitFieldName) {
		return nil, false
	}
	return o.OmitFieldName, true
}

// HasOmitFieldName returns a boolean if a field has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) HasOmitFieldName() bool {
	if o != nil && !IsNil(o.OmitFieldName) {
		return true
	}

	return false
}

// SetOmitFieldName gets a reference to the given []string and assigns it to the OmitFieldName field.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetOmitFieldName(v []string) {
	o.OmitFieldName = v
}

// GetRedactEntireValueField returns the RedactEntireValueField field value if set, zero value otherwise.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetRedactEntireValueField() []EnumlogFieldBehaviorJsonFormattedAccessRedactEntireValueFieldProp {
	if o == nil || IsNil(o.RedactEntireValueField) {
		var ret []EnumlogFieldBehaviorJsonFormattedAccessRedactEntireValueFieldProp
		return ret
	}
	return o.RedactEntireValueField
}

// GetRedactEntireValueFieldOk returns a tuple with the RedactEntireValueField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetRedactEntireValueFieldOk() ([]EnumlogFieldBehaviorJsonFormattedAccessRedactEntireValueFieldProp, bool) {
	if o == nil || IsNil(o.RedactEntireValueField) {
		return nil, false
	}
	return o.RedactEntireValueField, true
}

// HasRedactEntireValueField returns a boolean if a field has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) HasRedactEntireValueField() bool {
	if o != nil && !IsNil(o.RedactEntireValueField) {
		return true
	}

	return false
}

// SetRedactEntireValueField gets a reference to the given []EnumlogFieldBehaviorJsonFormattedAccessRedactEntireValueFieldProp and assigns it to the RedactEntireValueField field.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetRedactEntireValueField(v []EnumlogFieldBehaviorJsonFormattedAccessRedactEntireValueFieldProp) {
	o.RedactEntireValueField = v
}

// GetRedactEntireValueFieldName returns the RedactEntireValueFieldName field value if set, zero value otherwise.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetRedactEntireValueFieldName() []string {
	if o == nil || IsNil(o.RedactEntireValueFieldName) {
		var ret []string
		return ret
	}
	return o.RedactEntireValueFieldName
}

// GetRedactEntireValueFieldNameOk returns a tuple with the RedactEntireValueFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetRedactEntireValueFieldNameOk() ([]string, bool) {
	if o == nil || IsNil(o.RedactEntireValueFieldName) {
		return nil, false
	}
	return o.RedactEntireValueFieldName, true
}

// HasRedactEntireValueFieldName returns a boolean if a field has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) HasRedactEntireValueFieldName() bool {
	if o != nil && !IsNil(o.RedactEntireValueFieldName) {
		return true
	}

	return false
}

// SetRedactEntireValueFieldName gets a reference to the given []string and assigns it to the RedactEntireValueFieldName field.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetRedactEntireValueFieldName(v []string) {
	o.RedactEntireValueFieldName = v
}

// GetRedactValueComponentsField returns the RedactValueComponentsField field value if set, zero value otherwise.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetRedactValueComponentsField() []EnumlogFieldBehaviorJsonFormattedAccessRedactValueComponentsFieldProp {
	if o == nil || IsNil(o.RedactValueComponentsField) {
		var ret []EnumlogFieldBehaviorJsonFormattedAccessRedactValueComponentsFieldProp
		return ret
	}
	return o.RedactValueComponentsField
}

// GetRedactValueComponentsFieldOk returns a tuple with the RedactValueComponentsField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetRedactValueComponentsFieldOk() ([]EnumlogFieldBehaviorJsonFormattedAccessRedactValueComponentsFieldProp, bool) {
	if o == nil || IsNil(o.RedactValueComponentsField) {
		return nil, false
	}
	return o.RedactValueComponentsField, true
}

// HasRedactValueComponentsField returns a boolean if a field has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) HasRedactValueComponentsField() bool {
	if o != nil && !IsNil(o.RedactValueComponentsField) {
		return true
	}

	return false
}

// SetRedactValueComponentsField gets a reference to the given []EnumlogFieldBehaviorJsonFormattedAccessRedactValueComponentsFieldProp and assigns it to the RedactValueComponentsField field.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetRedactValueComponentsField(v []EnumlogFieldBehaviorJsonFormattedAccessRedactValueComponentsFieldProp) {
	o.RedactValueComponentsField = v
}

// GetRedactValueComponentsFieldName returns the RedactValueComponentsFieldName field value if set, zero value otherwise.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetRedactValueComponentsFieldName() []string {
	if o == nil || IsNil(o.RedactValueComponentsFieldName) {
		var ret []string
		return ret
	}
	return o.RedactValueComponentsFieldName
}

// GetRedactValueComponentsFieldNameOk returns a tuple with the RedactValueComponentsFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetRedactValueComponentsFieldNameOk() ([]string, bool) {
	if o == nil || IsNil(o.RedactValueComponentsFieldName) {
		return nil, false
	}
	return o.RedactValueComponentsFieldName, true
}

// HasRedactValueComponentsFieldName returns a boolean if a field has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) HasRedactValueComponentsFieldName() bool {
	if o != nil && !IsNil(o.RedactValueComponentsFieldName) {
		return true
	}

	return false
}

// SetRedactValueComponentsFieldName gets a reference to the given []string and assigns it to the RedactValueComponentsFieldName field.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetRedactValueComponentsFieldName(v []string) {
	o.RedactValueComponentsFieldName = v
}

// GetTokenizeEntireValueField returns the TokenizeEntireValueField field value if set, zero value otherwise.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetTokenizeEntireValueField() []EnumlogFieldBehaviorJsonFormattedAccessTokenizeEntireValueFieldProp {
	if o == nil || IsNil(o.TokenizeEntireValueField) {
		var ret []EnumlogFieldBehaviorJsonFormattedAccessTokenizeEntireValueFieldProp
		return ret
	}
	return o.TokenizeEntireValueField
}

// GetTokenizeEntireValueFieldOk returns a tuple with the TokenizeEntireValueField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetTokenizeEntireValueFieldOk() ([]EnumlogFieldBehaviorJsonFormattedAccessTokenizeEntireValueFieldProp, bool) {
	if o == nil || IsNil(o.TokenizeEntireValueField) {
		return nil, false
	}
	return o.TokenizeEntireValueField, true
}

// HasTokenizeEntireValueField returns a boolean if a field has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) HasTokenizeEntireValueField() bool {
	if o != nil && !IsNil(o.TokenizeEntireValueField) {
		return true
	}

	return false
}

// SetTokenizeEntireValueField gets a reference to the given []EnumlogFieldBehaviorJsonFormattedAccessTokenizeEntireValueFieldProp and assigns it to the TokenizeEntireValueField field.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetTokenizeEntireValueField(v []EnumlogFieldBehaviorJsonFormattedAccessTokenizeEntireValueFieldProp) {
	o.TokenizeEntireValueField = v
}

// GetTokenizeEntireValueFieldName returns the TokenizeEntireValueFieldName field value if set, zero value otherwise.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetTokenizeEntireValueFieldName() []string {
	if o == nil || IsNil(o.TokenizeEntireValueFieldName) {
		var ret []string
		return ret
	}
	return o.TokenizeEntireValueFieldName
}

// GetTokenizeEntireValueFieldNameOk returns a tuple with the TokenizeEntireValueFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetTokenizeEntireValueFieldNameOk() ([]string, bool) {
	if o == nil || IsNil(o.TokenizeEntireValueFieldName) {
		return nil, false
	}
	return o.TokenizeEntireValueFieldName, true
}

// HasTokenizeEntireValueFieldName returns a boolean if a field has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) HasTokenizeEntireValueFieldName() bool {
	if o != nil && !IsNil(o.TokenizeEntireValueFieldName) {
		return true
	}

	return false
}

// SetTokenizeEntireValueFieldName gets a reference to the given []string and assigns it to the TokenizeEntireValueFieldName field.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetTokenizeEntireValueFieldName(v []string) {
	o.TokenizeEntireValueFieldName = v
}

// GetTokenizeValueComponentsField returns the TokenizeValueComponentsField field value if set, zero value otherwise.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetTokenizeValueComponentsField() []EnumlogFieldBehaviorJsonFormattedAccessTokenizeValueComponentsFieldProp {
	if o == nil || IsNil(o.TokenizeValueComponentsField) {
		var ret []EnumlogFieldBehaviorJsonFormattedAccessTokenizeValueComponentsFieldProp
		return ret
	}
	return o.TokenizeValueComponentsField
}

// GetTokenizeValueComponentsFieldOk returns a tuple with the TokenizeValueComponentsField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetTokenizeValueComponentsFieldOk() ([]EnumlogFieldBehaviorJsonFormattedAccessTokenizeValueComponentsFieldProp, bool) {
	if o == nil || IsNil(o.TokenizeValueComponentsField) {
		return nil, false
	}
	return o.TokenizeValueComponentsField, true
}

// HasTokenizeValueComponentsField returns a boolean if a field has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) HasTokenizeValueComponentsField() bool {
	if o != nil && !IsNil(o.TokenizeValueComponentsField) {
		return true
	}

	return false
}

// SetTokenizeValueComponentsField gets a reference to the given []EnumlogFieldBehaviorJsonFormattedAccessTokenizeValueComponentsFieldProp and assigns it to the TokenizeValueComponentsField field.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetTokenizeValueComponentsField(v []EnumlogFieldBehaviorJsonFormattedAccessTokenizeValueComponentsFieldProp) {
	o.TokenizeValueComponentsField = v
}

// GetTokenizeValueComponentsFieldName returns the TokenizeValueComponentsFieldName field value if set, zero value otherwise.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetTokenizeValueComponentsFieldName() []string {
	if o == nil || IsNil(o.TokenizeValueComponentsFieldName) {
		var ret []string
		return ret
	}
	return o.TokenizeValueComponentsFieldName
}

// GetTokenizeValueComponentsFieldNameOk returns a tuple with the TokenizeValueComponentsFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetTokenizeValueComponentsFieldNameOk() ([]string, bool) {
	if o == nil || IsNil(o.TokenizeValueComponentsFieldName) {
		return nil, false
	}
	return o.TokenizeValueComponentsFieldName, true
}

// HasTokenizeValueComponentsFieldName returns a boolean if a field has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) HasTokenizeValueComponentsFieldName() bool {
	if o != nil && !IsNil(o.TokenizeValueComponentsFieldName) {
		return true
	}

	return false
}

// SetTokenizeValueComponentsFieldName gets a reference to the given []string and assigns it to the TokenizeValueComponentsFieldName field.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetTokenizeValueComponentsFieldName(v []string) {
	o.TokenizeValueComponentsFieldName = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDefaultBehavior returns the DefaultBehavior field value if set, zero value otherwise.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetDefaultBehavior() EnumlogFieldBehaviorDefaultBehaviorProp {
	if o == nil || IsNil(o.DefaultBehavior) {
		var ret EnumlogFieldBehaviorDefaultBehaviorProp
		return ret
	}
	return *o.DefaultBehavior
}

// GetDefaultBehaviorOk returns a tuple with the DefaultBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetDefaultBehaviorOk() (*EnumlogFieldBehaviorDefaultBehaviorProp, bool) {
	if o == nil || IsNil(o.DefaultBehavior) {
		return nil, false
	}
	return o.DefaultBehavior, true
}

// HasDefaultBehavior returns a boolean if a field has been set.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) HasDefaultBehavior() bool {
	if o != nil && !IsNil(o.DefaultBehavior) {
		return true
	}

	return false
}

// SetDefaultBehavior gets a reference to the given EnumlogFieldBehaviorDefaultBehaviorProp and assigns it to the DefaultBehavior field.
func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetDefaultBehavior(v EnumlogFieldBehaviorDefaultBehaviorProp) {
	o.DefaultBehavior = &v
}

func (o AddJsonFormattedAccessLogFieldBehaviorRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddJsonFormattedAccessLogFieldBehaviorRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["behaviorName"] = o.BehaviorName
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.PreserveField) {
		toSerialize["preserveField"] = o.PreserveField
	}
	if !IsNil(o.PreserveFieldName) {
		toSerialize["preserveFieldName"] = o.PreserveFieldName
	}
	if !IsNil(o.OmitField) {
		toSerialize["omitField"] = o.OmitField
	}
	if !IsNil(o.OmitFieldName) {
		toSerialize["omitFieldName"] = o.OmitFieldName
	}
	if !IsNil(o.RedactEntireValueField) {
		toSerialize["redactEntireValueField"] = o.RedactEntireValueField
	}
	if !IsNil(o.RedactEntireValueFieldName) {
		toSerialize["redactEntireValueFieldName"] = o.RedactEntireValueFieldName
	}
	if !IsNil(o.RedactValueComponentsField) {
		toSerialize["redactValueComponentsField"] = o.RedactValueComponentsField
	}
	if !IsNil(o.RedactValueComponentsFieldName) {
		toSerialize["redactValueComponentsFieldName"] = o.RedactValueComponentsFieldName
	}
	if !IsNil(o.TokenizeEntireValueField) {
		toSerialize["tokenizeEntireValueField"] = o.TokenizeEntireValueField
	}
	if !IsNil(o.TokenizeEntireValueFieldName) {
		toSerialize["tokenizeEntireValueFieldName"] = o.TokenizeEntireValueFieldName
	}
	if !IsNil(o.TokenizeValueComponentsField) {
		toSerialize["tokenizeValueComponentsField"] = o.TokenizeValueComponentsField
	}
	if !IsNil(o.TokenizeValueComponentsFieldName) {
		toSerialize["tokenizeValueComponentsFieldName"] = o.TokenizeValueComponentsFieldName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DefaultBehavior) {
		toSerialize["defaultBehavior"] = o.DefaultBehavior
	}
	return toSerialize, nil
}

type NullableAddJsonFormattedAccessLogFieldBehaviorRequest struct {
	value *AddJsonFormattedAccessLogFieldBehaviorRequest
	isSet bool
}

func (v NullableAddJsonFormattedAccessLogFieldBehaviorRequest) Get() *AddJsonFormattedAccessLogFieldBehaviorRequest {
	return v.value
}

func (v *NullableAddJsonFormattedAccessLogFieldBehaviorRequest) Set(val *AddJsonFormattedAccessLogFieldBehaviorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddJsonFormattedAccessLogFieldBehaviorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddJsonFormattedAccessLogFieldBehaviorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddJsonFormattedAccessLogFieldBehaviorRequest(val *AddJsonFormattedAccessLogFieldBehaviorRequest) *NullableAddJsonFormattedAccessLogFieldBehaviorRequest {
	return &NullableAddJsonFormattedAccessLogFieldBehaviorRequest{value: val, isSet: true}
}

func (v NullableAddJsonFormattedAccessLogFieldBehaviorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddJsonFormattedAccessLogFieldBehaviorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
