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

// SnmpSubAgentAlertHandlerResponse struct for SnmpSubAgentAlertHandlerResponse
type SnmpSubAgentAlertHandlerResponse struct {
	// Name of the Alert Handler
	Id string `json:"id"`
	Schemas []EnumsnmpSubAgentAlertHandlerSchemaUrn `json:"schemas"`
	// Indicates whether the server should attempt to invoke this SNMP Sub Agent Alert Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver the alert notification) will not delay whatever processing the server was performing when the alert was generated.
	Asynchronous *bool `json:"asynchronous,omitempty"`
	// A description for this Alert Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the Alert Handler is enabled.
	Enabled bool `json:"enabled"`
	EnabledAlertSeverity []EnumalertHandlerEnabledAlertSeverityProp `json:"enabledAlertSeverity,omitempty"`
	EnabledAlertType []EnumalertHandlerEnabledAlertTypeProp `json:"enabledAlertType,omitempty"`
	DisabledAlertType []EnumalertHandlerDisabledAlertTypeProp `json:"disabledAlertType,omitempty"`
	Meta *MetaMeta `json:"meta,omitempty"`
}

// NewSnmpSubAgentAlertHandlerResponse instantiates a new SnmpSubAgentAlertHandlerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnmpSubAgentAlertHandlerResponse(id string, schemas []EnumsnmpSubAgentAlertHandlerSchemaUrn, enabled bool) *SnmpSubAgentAlertHandlerResponse {
	this := SnmpSubAgentAlertHandlerResponse{}
	this.Id = id
	this.Schemas = schemas
	this.Enabled = enabled
	return &this
}

// NewSnmpSubAgentAlertHandlerResponseWithDefaults instantiates a new SnmpSubAgentAlertHandlerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnmpSubAgentAlertHandlerResponseWithDefaults() *SnmpSubAgentAlertHandlerResponse {
	this := SnmpSubAgentAlertHandlerResponse{}
	return &this
}

// GetId returns the Id field value
func (o *SnmpSubAgentAlertHandlerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SnmpSubAgentAlertHandlerResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SnmpSubAgentAlertHandlerResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *SnmpSubAgentAlertHandlerResponse) GetSchemas() []EnumsnmpSubAgentAlertHandlerSchemaUrn {
	if o == nil {
		var ret []EnumsnmpSubAgentAlertHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *SnmpSubAgentAlertHandlerResponse) GetSchemasOk() ([]EnumsnmpSubAgentAlertHandlerSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *SnmpSubAgentAlertHandlerResponse) SetSchemas(v []EnumsnmpSubAgentAlertHandlerSchemaUrn) {
	o.Schemas = v
}

// GetAsynchronous returns the Asynchronous field value if set, zero value otherwise.
func (o *SnmpSubAgentAlertHandlerResponse) GetAsynchronous() bool {
	if o == nil || isNil(o.Asynchronous) {
		var ret bool
		return ret
	}
	return *o.Asynchronous
}

// GetAsynchronousOk returns a tuple with the Asynchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpSubAgentAlertHandlerResponse) GetAsynchronousOk() (*bool, bool) {
	if o == nil || isNil(o.Asynchronous) {
    return nil, false
	}
	return o.Asynchronous, true
}

// HasAsynchronous returns a boolean if a field has been set.
func (o *SnmpSubAgentAlertHandlerResponse) HasAsynchronous() bool {
	if o != nil && !isNil(o.Asynchronous) {
		return true
	}

	return false
}

// SetAsynchronous gets a reference to the given bool and assigns it to the Asynchronous field.
func (o *SnmpSubAgentAlertHandlerResponse) SetAsynchronous(v bool) {
	o.Asynchronous = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SnmpSubAgentAlertHandlerResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpSubAgentAlertHandlerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SnmpSubAgentAlertHandlerResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SnmpSubAgentAlertHandlerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *SnmpSubAgentAlertHandlerResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SnmpSubAgentAlertHandlerResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SnmpSubAgentAlertHandlerResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEnabledAlertSeverity returns the EnabledAlertSeverity field value if set, zero value otherwise.
func (o *SnmpSubAgentAlertHandlerResponse) GetEnabledAlertSeverity() []EnumalertHandlerEnabledAlertSeverityProp {
	if o == nil || isNil(o.EnabledAlertSeverity) {
		var ret []EnumalertHandlerEnabledAlertSeverityProp
		return ret
	}
	return o.EnabledAlertSeverity
}

// GetEnabledAlertSeverityOk returns a tuple with the EnabledAlertSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpSubAgentAlertHandlerResponse) GetEnabledAlertSeverityOk() ([]EnumalertHandlerEnabledAlertSeverityProp, bool) {
	if o == nil || isNil(o.EnabledAlertSeverity) {
    return nil, false
	}
	return o.EnabledAlertSeverity, true
}

// HasEnabledAlertSeverity returns a boolean if a field has been set.
func (o *SnmpSubAgentAlertHandlerResponse) HasEnabledAlertSeverity() bool {
	if o != nil && !isNil(o.EnabledAlertSeverity) {
		return true
	}

	return false
}

// SetEnabledAlertSeverity gets a reference to the given []EnumalertHandlerEnabledAlertSeverityProp and assigns it to the EnabledAlertSeverity field.
func (o *SnmpSubAgentAlertHandlerResponse) SetEnabledAlertSeverity(v []EnumalertHandlerEnabledAlertSeverityProp) {
	o.EnabledAlertSeverity = v
}

// GetEnabledAlertType returns the EnabledAlertType field value if set, zero value otherwise.
func (o *SnmpSubAgentAlertHandlerResponse) GetEnabledAlertType() []EnumalertHandlerEnabledAlertTypeProp {
	if o == nil || isNil(o.EnabledAlertType) {
		var ret []EnumalertHandlerEnabledAlertTypeProp
		return ret
	}
	return o.EnabledAlertType
}

// GetEnabledAlertTypeOk returns a tuple with the EnabledAlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpSubAgentAlertHandlerResponse) GetEnabledAlertTypeOk() ([]EnumalertHandlerEnabledAlertTypeProp, bool) {
	if o == nil || isNil(o.EnabledAlertType) {
    return nil, false
	}
	return o.EnabledAlertType, true
}

// HasEnabledAlertType returns a boolean if a field has been set.
func (o *SnmpSubAgentAlertHandlerResponse) HasEnabledAlertType() bool {
	if o != nil && !isNil(o.EnabledAlertType) {
		return true
	}

	return false
}

// SetEnabledAlertType gets a reference to the given []EnumalertHandlerEnabledAlertTypeProp and assigns it to the EnabledAlertType field.
func (o *SnmpSubAgentAlertHandlerResponse) SetEnabledAlertType(v []EnumalertHandlerEnabledAlertTypeProp) {
	o.EnabledAlertType = v
}

// GetDisabledAlertType returns the DisabledAlertType field value if set, zero value otherwise.
func (o *SnmpSubAgentAlertHandlerResponse) GetDisabledAlertType() []EnumalertHandlerDisabledAlertTypeProp {
	if o == nil || isNil(o.DisabledAlertType) {
		var ret []EnumalertHandlerDisabledAlertTypeProp
		return ret
	}
	return o.DisabledAlertType
}

// GetDisabledAlertTypeOk returns a tuple with the DisabledAlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpSubAgentAlertHandlerResponse) GetDisabledAlertTypeOk() ([]EnumalertHandlerDisabledAlertTypeProp, bool) {
	if o == nil || isNil(o.DisabledAlertType) {
    return nil, false
	}
	return o.DisabledAlertType, true
}

// HasDisabledAlertType returns a boolean if a field has been set.
func (o *SnmpSubAgentAlertHandlerResponse) HasDisabledAlertType() bool {
	if o != nil && !isNil(o.DisabledAlertType) {
		return true
	}

	return false
}

// SetDisabledAlertType gets a reference to the given []EnumalertHandlerDisabledAlertTypeProp and assigns it to the DisabledAlertType field.
func (o *SnmpSubAgentAlertHandlerResponse) SetDisabledAlertType(v []EnumalertHandlerDisabledAlertTypeProp) {
	o.DisabledAlertType = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SnmpSubAgentAlertHandlerResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpSubAgentAlertHandlerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SnmpSubAgentAlertHandlerResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *SnmpSubAgentAlertHandlerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

func (o SnmpSubAgentAlertHandlerResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.Asynchronous) {
		toSerialize["asynchronous"] = o.Asynchronous
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.EnabledAlertSeverity) {
		toSerialize["enabledAlertSeverity"] = o.EnabledAlertSeverity
	}
	if !isNil(o.EnabledAlertType) {
		toSerialize["enabledAlertType"] = o.EnabledAlertType
	}
	if !isNil(o.DisabledAlertType) {
		toSerialize["disabledAlertType"] = o.DisabledAlertType
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableSnmpSubAgentAlertHandlerResponse struct {
	value *SnmpSubAgentAlertHandlerResponse
	isSet bool
}

func (v NullableSnmpSubAgentAlertHandlerResponse) Get() *SnmpSubAgentAlertHandlerResponse {
	return v.value
}

func (v *NullableSnmpSubAgentAlertHandlerResponse) Set(val *SnmpSubAgentAlertHandlerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSnmpSubAgentAlertHandlerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSnmpSubAgentAlertHandlerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnmpSubAgentAlertHandlerResponse(val *SnmpSubAgentAlertHandlerResponse) *NullableSnmpSubAgentAlertHandlerResponse {
	return &NullableSnmpSubAgentAlertHandlerResponse{value: val, isSet: true}
}

func (v NullableSnmpSubAgentAlertHandlerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnmpSubAgentAlertHandlerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


