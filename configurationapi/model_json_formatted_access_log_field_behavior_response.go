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

// checks if the JsonFormattedAccessLogFieldBehaviorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JsonFormattedAccessLogFieldBehaviorResponse{}

// JsonFormattedAccessLogFieldBehaviorResponse struct for JsonFormattedAccessLogFieldBehaviorResponse
type JsonFormattedAccessLogFieldBehaviorResponse struct {
	// Name of the Log Field Behavior
	Id            string                                                     `json:"id"`
	Schemas       []EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn         `json:"schemas"`
	PreserveField []EnumlogFieldBehaviorJsonFormattedAccessPreserveFieldProp `json:"preserveField,omitempty"`
	// The names of any custom fields whose values should be preserved. This should generally only be used for fields that are not available through the preserve-field property (for example, custom log fields defined in Server SDK extensions).
	PreserveFieldName []string                                               `json:"preserveFieldName,omitempty"`
	OmitField         []EnumlogFieldBehaviorJsonFormattedAccessOmitFieldProp `json:"omitField,omitempty"`
	// The names of any custom fields that should be omitted from log messages. This should generally only be used for fields that are not available through the omit-field property (for example, custom log fields defined in Server SDK extensions).
	OmitFieldName          []string                                                            `json:"omitFieldName,omitempty"`
	RedactEntireValueField []EnumlogFieldBehaviorJsonFormattedAccessRedactEntireValueFieldProp `json:"redactEntireValueField,omitempty"`
	// The names of any custom fields whose values should be completely redacted. This should generally only be used for fields that are not available through the redact-entire-value-field property (for example, custom log fields defined in Server SDK extensions).
	RedactEntireValueFieldName []string                                                                `json:"redactEntireValueFieldName,omitempty"`
	RedactValueComponentsField []EnumlogFieldBehaviorJsonFormattedAccessRedactValueComponentsFieldProp `json:"redactValueComponentsField,omitempty"`
	// The names of any custom fields for which to redact components within the value. This should generally only be used for fields that are not available through the redact-value-components-field property (for example, custom log fields defined in Server SDK extensions).
	RedactValueComponentsFieldName []string                                                              `json:"redactValueComponentsFieldName,omitempty"`
	TokenizeEntireValueField       []EnumlogFieldBehaviorJsonFormattedAccessTokenizeEntireValueFieldProp `json:"tokenizeEntireValueField,omitempty"`
	// The names of any custom fields whose values should be completely tokenized. This should generally only be used for fields that are not available through the tokenize-entire-value-field property (for example, custom log fields defined in Server SDK extensions).
	TokenizeEntireValueFieldName []string                                                                  `json:"tokenizeEntireValueFieldName,omitempty"`
	TokenizeValueComponentsField []EnumlogFieldBehaviorJsonFormattedAccessTokenizeValueComponentsFieldProp `json:"tokenizeValueComponentsField,omitempty"`
	// The names of any custom fields for which to tokenize components within the value. This should generally only be used for fields that are not available through the tokenize-value-components-field property (for example, custom log fields defined in Server SDK extensions).
	TokenizeValueComponentsFieldName []string `json:"tokenizeValueComponentsFieldName,omitempty"`
	// A description for this Log Field Behavior
	Description                                   *string                                            `json:"description,omitempty"`
	DefaultBehavior                               *EnumlogFieldBehaviorDefaultBehaviorProp           `json:"defaultBehavior,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewJsonFormattedAccessLogFieldBehaviorResponse instantiates a new JsonFormattedAccessLogFieldBehaviorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonFormattedAccessLogFieldBehaviorResponse(id string, schemas []EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn) *JsonFormattedAccessLogFieldBehaviorResponse {
	this := JsonFormattedAccessLogFieldBehaviorResponse{}
	this.Id = id
	this.Schemas = schemas
	return &this
}

// NewJsonFormattedAccessLogFieldBehaviorResponseWithDefaults instantiates a new JsonFormattedAccessLogFieldBehaviorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonFormattedAccessLogFieldBehaviorResponseWithDefaults() *JsonFormattedAccessLogFieldBehaviorResponse {
	this := JsonFormattedAccessLogFieldBehaviorResponse{}
	return &this
}

// GetId returns the Id field value
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JsonFormattedAccessLogFieldBehaviorResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetSchemas() []EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn {
	if o == nil {
		var ret []EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetSchemasOk() ([]EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *JsonFormattedAccessLogFieldBehaviorResponse) SetSchemas(v []EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn) {
	o.Schemas = v
}

// GetPreserveField returns the PreserveField field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetPreserveField() []EnumlogFieldBehaviorJsonFormattedAccessPreserveFieldProp {
	if o == nil || IsNil(o.PreserveField) {
		var ret []EnumlogFieldBehaviorJsonFormattedAccessPreserveFieldProp
		return ret
	}
	return o.PreserveField
}

// GetPreserveFieldOk returns a tuple with the PreserveField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetPreserveFieldOk() ([]EnumlogFieldBehaviorJsonFormattedAccessPreserveFieldProp, bool) {
	if o == nil || IsNil(o.PreserveField) {
		return nil, false
	}
	return o.PreserveField, true
}

// HasPreserveField returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) HasPreserveField() bool {
	if o != nil && !IsNil(o.PreserveField) {
		return true
	}

	return false
}

// SetPreserveField gets a reference to the given []EnumlogFieldBehaviorJsonFormattedAccessPreserveFieldProp and assigns it to the PreserveField field.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) SetPreserveField(v []EnumlogFieldBehaviorJsonFormattedAccessPreserveFieldProp) {
	o.PreserveField = v
}

// GetPreserveFieldName returns the PreserveFieldName field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetPreserveFieldName() []string {
	if o == nil || IsNil(o.PreserveFieldName) {
		var ret []string
		return ret
	}
	return o.PreserveFieldName
}

// GetPreserveFieldNameOk returns a tuple with the PreserveFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetPreserveFieldNameOk() ([]string, bool) {
	if o == nil || IsNil(o.PreserveFieldName) {
		return nil, false
	}
	return o.PreserveFieldName, true
}

// HasPreserveFieldName returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) HasPreserveFieldName() bool {
	if o != nil && !IsNil(o.PreserveFieldName) {
		return true
	}

	return false
}

// SetPreserveFieldName gets a reference to the given []string and assigns it to the PreserveFieldName field.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) SetPreserveFieldName(v []string) {
	o.PreserveFieldName = v
}

// GetOmitField returns the OmitField field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetOmitField() []EnumlogFieldBehaviorJsonFormattedAccessOmitFieldProp {
	if o == nil || IsNil(o.OmitField) {
		var ret []EnumlogFieldBehaviorJsonFormattedAccessOmitFieldProp
		return ret
	}
	return o.OmitField
}

// GetOmitFieldOk returns a tuple with the OmitField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetOmitFieldOk() ([]EnumlogFieldBehaviorJsonFormattedAccessOmitFieldProp, bool) {
	if o == nil || IsNil(o.OmitField) {
		return nil, false
	}
	return o.OmitField, true
}

// HasOmitField returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) HasOmitField() bool {
	if o != nil && !IsNil(o.OmitField) {
		return true
	}

	return false
}

// SetOmitField gets a reference to the given []EnumlogFieldBehaviorJsonFormattedAccessOmitFieldProp and assigns it to the OmitField field.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) SetOmitField(v []EnumlogFieldBehaviorJsonFormattedAccessOmitFieldProp) {
	o.OmitField = v
}

// GetOmitFieldName returns the OmitFieldName field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetOmitFieldName() []string {
	if o == nil || IsNil(o.OmitFieldName) {
		var ret []string
		return ret
	}
	return o.OmitFieldName
}

// GetOmitFieldNameOk returns a tuple with the OmitFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetOmitFieldNameOk() ([]string, bool) {
	if o == nil || IsNil(o.OmitFieldName) {
		return nil, false
	}
	return o.OmitFieldName, true
}

// HasOmitFieldName returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) HasOmitFieldName() bool {
	if o != nil && !IsNil(o.OmitFieldName) {
		return true
	}

	return false
}

// SetOmitFieldName gets a reference to the given []string and assigns it to the OmitFieldName field.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) SetOmitFieldName(v []string) {
	o.OmitFieldName = v
}

// GetRedactEntireValueField returns the RedactEntireValueField field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetRedactEntireValueField() []EnumlogFieldBehaviorJsonFormattedAccessRedactEntireValueFieldProp {
	if o == nil || IsNil(o.RedactEntireValueField) {
		var ret []EnumlogFieldBehaviorJsonFormattedAccessRedactEntireValueFieldProp
		return ret
	}
	return o.RedactEntireValueField
}

// GetRedactEntireValueFieldOk returns a tuple with the RedactEntireValueField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetRedactEntireValueFieldOk() ([]EnumlogFieldBehaviorJsonFormattedAccessRedactEntireValueFieldProp, bool) {
	if o == nil || IsNil(o.RedactEntireValueField) {
		return nil, false
	}
	return o.RedactEntireValueField, true
}

// HasRedactEntireValueField returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) HasRedactEntireValueField() bool {
	if o != nil && !IsNil(o.RedactEntireValueField) {
		return true
	}

	return false
}

// SetRedactEntireValueField gets a reference to the given []EnumlogFieldBehaviorJsonFormattedAccessRedactEntireValueFieldProp and assigns it to the RedactEntireValueField field.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) SetRedactEntireValueField(v []EnumlogFieldBehaviorJsonFormattedAccessRedactEntireValueFieldProp) {
	o.RedactEntireValueField = v
}

// GetRedactEntireValueFieldName returns the RedactEntireValueFieldName field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetRedactEntireValueFieldName() []string {
	if o == nil || IsNil(o.RedactEntireValueFieldName) {
		var ret []string
		return ret
	}
	return o.RedactEntireValueFieldName
}

// GetRedactEntireValueFieldNameOk returns a tuple with the RedactEntireValueFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetRedactEntireValueFieldNameOk() ([]string, bool) {
	if o == nil || IsNil(o.RedactEntireValueFieldName) {
		return nil, false
	}
	return o.RedactEntireValueFieldName, true
}

// HasRedactEntireValueFieldName returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) HasRedactEntireValueFieldName() bool {
	if o != nil && !IsNil(o.RedactEntireValueFieldName) {
		return true
	}

	return false
}

// SetRedactEntireValueFieldName gets a reference to the given []string and assigns it to the RedactEntireValueFieldName field.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) SetRedactEntireValueFieldName(v []string) {
	o.RedactEntireValueFieldName = v
}

// GetRedactValueComponentsField returns the RedactValueComponentsField field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetRedactValueComponentsField() []EnumlogFieldBehaviorJsonFormattedAccessRedactValueComponentsFieldProp {
	if o == nil || IsNil(o.RedactValueComponentsField) {
		var ret []EnumlogFieldBehaviorJsonFormattedAccessRedactValueComponentsFieldProp
		return ret
	}
	return o.RedactValueComponentsField
}

// GetRedactValueComponentsFieldOk returns a tuple with the RedactValueComponentsField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetRedactValueComponentsFieldOk() ([]EnumlogFieldBehaviorJsonFormattedAccessRedactValueComponentsFieldProp, bool) {
	if o == nil || IsNil(o.RedactValueComponentsField) {
		return nil, false
	}
	return o.RedactValueComponentsField, true
}

// HasRedactValueComponentsField returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) HasRedactValueComponentsField() bool {
	if o != nil && !IsNil(o.RedactValueComponentsField) {
		return true
	}

	return false
}

// SetRedactValueComponentsField gets a reference to the given []EnumlogFieldBehaviorJsonFormattedAccessRedactValueComponentsFieldProp and assigns it to the RedactValueComponentsField field.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) SetRedactValueComponentsField(v []EnumlogFieldBehaviorJsonFormattedAccessRedactValueComponentsFieldProp) {
	o.RedactValueComponentsField = v
}

// GetRedactValueComponentsFieldName returns the RedactValueComponentsFieldName field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetRedactValueComponentsFieldName() []string {
	if o == nil || IsNil(o.RedactValueComponentsFieldName) {
		var ret []string
		return ret
	}
	return o.RedactValueComponentsFieldName
}

// GetRedactValueComponentsFieldNameOk returns a tuple with the RedactValueComponentsFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetRedactValueComponentsFieldNameOk() ([]string, bool) {
	if o == nil || IsNil(o.RedactValueComponentsFieldName) {
		return nil, false
	}
	return o.RedactValueComponentsFieldName, true
}

// HasRedactValueComponentsFieldName returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) HasRedactValueComponentsFieldName() bool {
	if o != nil && !IsNil(o.RedactValueComponentsFieldName) {
		return true
	}

	return false
}

// SetRedactValueComponentsFieldName gets a reference to the given []string and assigns it to the RedactValueComponentsFieldName field.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) SetRedactValueComponentsFieldName(v []string) {
	o.RedactValueComponentsFieldName = v
}

// GetTokenizeEntireValueField returns the TokenizeEntireValueField field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetTokenizeEntireValueField() []EnumlogFieldBehaviorJsonFormattedAccessTokenizeEntireValueFieldProp {
	if o == nil || IsNil(o.TokenizeEntireValueField) {
		var ret []EnumlogFieldBehaviorJsonFormattedAccessTokenizeEntireValueFieldProp
		return ret
	}
	return o.TokenizeEntireValueField
}

// GetTokenizeEntireValueFieldOk returns a tuple with the TokenizeEntireValueField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetTokenizeEntireValueFieldOk() ([]EnumlogFieldBehaviorJsonFormattedAccessTokenizeEntireValueFieldProp, bool) {
	if o == nil || IsNil(o.TokenizeEntireValueField) {
		return nil, false
	}
	return o.TokenizeEntireValueField, true
}

// HasTokenizeEntireValueField returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) HasTokenizeEntireValueField() bool {
	if o != nil && !IsNil(o.TokenizeEntireValueField) {
		return true
	}

	return false
}

// SetTokenizeEntireValueField gets a reference to the given []EnumlogFieldBehaviorJsonFormattedAccessTokenizeEntireValueFieldProp and assigns it to the TokenizeEntireValueField field.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) SetTokenizeEntireValueField(v []EnumlogFieldBehaviorJsonFormattedAccessTokenizeEntireValueFieldProp) {
	o.TokenizeEntireValueField = v
}

// GetTokenizeEntireValueFieldName returns the TokenizeEntireValueFieldName field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetTokenizeEntireValueFieldName() []string {
	if o == nil || IsNil(o.TokenizeEntireValueFieldName) {
		var ret []string
		return ret
	}
	return o.TokenizeEntireValueFieldName
}

// GetTokenizeEntireValueFieldNameOk returns a tuple with the TokenizeEntireValueFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetTokenizeEntireValueFieldNameOk() ([]string, bool) {
	if o == nil || IsNil(o.TokenizeEntireValueFieldName) {
		return nil, false
	}
	return o.TokenizeEntireValueFieldName, true
}

// HasTokenizeEntireValueFieldName returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) HasTokenizeEntireValueFieldName() bool {
	if o != nil && !IsNil(o.TokenizeEntireValueFieldName) {
		return true
	}

	return false
}

// SetTokenizeEntireValueFieldName gets a reference to the given []string and assigns it to the TokenizeEntireValueFieldName field.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) SetTokenizeEntireValueFieldName(v []string) {
	o.TokenizeEntireValueFieldName = v
}

// GetTokenizeValueComponentsField returns the TokenizeValueComponentsField field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetTokenizeValueComponentsField() []EnumlogFieldBehaviorJsonFormattedAccessTokenizeValueComponentsFieldProp {
	if o == nil || IsNil(o.TokenizeValueComponentsField) {
		var ret []EnumlogFieldBehaviorJsonFormattedAccessTokenizeValueComponentsFieldProp
		return ret
	}
	return o.TokenizeValueComponentsField
}

// GetTokenizeValueComponentsFieldOk returns a tuple with the TokenizeValueComponentsField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetTokenizeValueComponentsFieldOk() ([]EnumlogFieldBehaviorJsonFormattedAccessTokenizeValueComponentsFieldProp, bool) {
	if o == nil || IsNil(o.TokenizeValueComponentsField) {
		return nil, false
	}
	return o.TokenizeValueComponentsField, true
}

// HasTokenizeValueComponentsField returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) HasTokenizeValueComponentsField() bool {
	if o != nil && !IsNil(o.TokenizeValueComponentsField) {
		return true
	}

	return false
}

// SetTokenizeValueComponentsField gets a reference to the given []EnumlogFieldBehaviorJsonFormattedAccessTokenizeValueComponentsFieldProp and assigns it to the TokenizeValueComponentsField field.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) SetTokenizeValueComponentsField(v []EnumlogFieldBehaviorJsonFormattedAccessTokenizeValueComponentsFieldProp) {
	o.TokenizeValueComponentsField = v
}

// GetTokenizeValueComponentsFieldName returns the TokenizeValueComponentsFieldName field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetTokenizeValueComponentsFieldName() []string {
	if o == nil || IsNil(o.TokenizeValueComponentsFieldName) {
		var ret []string
		return ret
	}
	return o.TokenizeValueComponentsFieldName
}

// GetTokenizeValueComponentsFieldNameOk returns a tuple with the TokenizeValueComponentsFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetTokenizeValueComponentsFieldNameOk() ([]string, bool) {
	if o == nil || IsNil(o.TokenizeValueComponentsFieldName) {
		return nil, false
	}
	return o.TokenizeValueComponentsFieldName, true
}

// HasTokenizeValueComponentsFieldName returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) HasTokenizeValueComponentsFieldName() bool {
	if o != nil && !IsNil(o.TokenizeValueComponentsFieldName) {
		return true
	}

	return false
}

// SetTokenizeValueComponentsFieldName gets a reference to the given []string and assigns it to the TokenizeValueComponentsFieldName field.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) SetTokenizeValueComponentsFieldName(v []string) {
	o.TokenizeValueComponentsFieldName = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) SetDescription(v string) {
	o.Description = &v
}

// GetDefaultBehavior returns the DefaultBehavior field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetDefaultBehavior() EnumlogFieldBehaviorDefaultBehaviorProp {
	if o == nil || IsNil(o.DefaultBehavior) {
		var ret EnumlogFieldBehaviorDefaultBehaviorProp
		return ret
	}
	return *o.DefaultBehavior
}

// GetDefaultBehaviorOk returns a tuple with the DefaultBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetDefaultBehaviorOk() (*EnumlogFieldBehaviorDefaultBehaviorProp, bool) {
	if o == nil || IsNil(o.DefaultBehavior) {
		return nil, false
	}
	return o.DefaultBehavior, true
}

// HasDefaultBehavior returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) HasDefaultBehavior() bool {
	if o != nil && !IsNil(o.DefaultBehavior) {
		return true
	}

	return false
}

// SetDefaultBehavior gets a reference to the given EnumlogFieldBehaviorDefaultBehaviorProp and assigns it to the DefaultBehavior field.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) SetDefaultBehavior(v EnumlogFieldBehaviorDefaultBehaviorProp) {
	o.DefaultBehavior = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *JsonFormattedAccessLogFieldBehaviorResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o JsonFormattedAccessLogFieldBehaviorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JsonFormattedAccessLogFieldBehaviorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
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
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableJsonFormattedAccessLogFieldBehaviorResponse struct {
	value *JsonFormattedAccessLogFieldBehaviorResponse
	isSet bool
}

func (v NullableJsonFormattedAccessLogFieldBehaviorResponse) Get() *JsonFormattedAccessLogFieldBehaviorResponse {
	return v.value
}

func (v *NullableJsonFormattedAccessLogFieldBehaviorResponse) Set(val *JsonFormattedAccessLogFieldBehaviorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonFormattedAccessLogFieldBehaviorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonFormattedAccessLogFieldBehaviorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonFormattedAccessLogFieldBehaviorResponse(val *JsonFormattedAccessLogFieldBehaviorResponse) *NullableJsonFormattedAccessLogFieldBehaviorResponse {
	return &NullableJsonFormattedAccessLogFieldBehaviorResponse{value: val, isSet: true}
}

func (v NullableJsonFormattedAccessLogFieldBehaviorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonFormattedAccessLogFieldBehaviorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
