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

// ExportReversiblePasswordsExtendedOperationHandlerResponse struct for ExportReversiblePasswordsExtendedOperationHandlerResponse
type ExportReversiblePasswordsExtendedOperationHandlerResponse struct {
	// Name of the Extended Operation Handler
	Id string `json:"id"`
	Schemas []EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn `json:"schemas"`
	// A description for this Extended Operation Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server).
	Enabled bool `json:"enabled"`
}

// NewExportReversiblePasswordsExtendedOperationHandlerResponse instantiates a new ExportReversiblePasswordsExtendedOperationHandlerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportReversiblePasswordsExtendedOperationHandlerResponse(id string, schemas []EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn, enabled bool) *ExportReversiblePasswordsExtendedOperationHandlerResponse {
	this := ExportReversiblePasswordsExtendedOperationHandlerResponse{}
	this.Id = id
	this.Schemas = schemas
	this.Enabled = enabled
	return &this
}

// NewExportReversiblePasswordsExtendedOperationHandlerResponseWithDefaults instantiates a new ExportReversiblePasswordsExtendedOperationHandlerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportReversiblePasswordsExtendedOperationHandlerResponseWithDefaults() *ExportReversiblePasswordsExtendedOperationHandlerResponse {
	this := ExportReversiblePasswordsExtendedOperationHandlerResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) GetSchemas() []EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn {
	if o == nil {
		var ret []EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) GetSchemasOk() ([]EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) SetSchemas(v []EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o ExportReversiblePasswordsExtendedOperationHandlerResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableExportReversiblePasswordsExtendedOperationHandlerResponse struct {
	value *ExportReversiblePasswordsExtendedOperationHandlerResponse
	isSet bool
}

func (v NullableExportReversiblePasswordsExtendedOperationHandlerResponse) Get() *ExportReversiblePasswordsExtendedOperationHandlerResponse {
	return v.value
}

func (v *NullableExportReversiblePasswordsExtendedOperationHandlerResponse) Set(val *ExportReversiblePasswordsExtendedOperationHandlerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExportReversiblePasswordsExtendedOperationHandlerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExportReversiblePasswordsExtendedOperationHandlerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportReversiblePasswordsExtendedOperationHandlerResponse(val *ExportReversiblePasswordsExtendedOperationHandlerResponse) *NullableExportReversiblePasswordsExtendedOperationHandlerResponse {
	return &NullableExportReversiblePasswordsExtendedOperationHandlerResponse{value: val, isSet: true}
}

func (v NullableExportReversiblePasswordsExtendedOperationHandlerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportReversiblePasswordsExtendedOperationHandlerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


