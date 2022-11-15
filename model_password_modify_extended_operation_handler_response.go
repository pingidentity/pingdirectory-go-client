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

// PasswordModifyExtendedOperationHandlerResponse struct for PasswordModifyExtendedOperationHandlerResponse
type PasswordModifyExtendedOperationHandlerResponse struct {
	Schemas []EnumpasswordModifyExtendedOperationHandlerSchemaUrn `json:"schemas"`
	// Specifies the name of the identity mapper that should be used in conjunction with the password modify extended operation.
	IdentityMapper string `json:"identityMapper"`
	// A description for this Extended Operation Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server).
	Enabled bool `json:"enabled"`
}

// NewPasswordModifyExtendedOperationHandlerResponse instantiates a new PasswordModifyExtendedOperationHandlerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordModifyExtendedOperationHandlerResponse(schemas []EnumpasswordModifyExtendedOperationHandlerSchemaUrn, identityMapper string, enabled bool) *PasswordModifyExtendedOperationHandlerResponse {
	this := PasswordModifyExtendedOperationHandlerResponse{}
	this.Schemas = schemas
	this.IdentityMapper = identityMapper
	this.Enabled = enabled
	return &this
}

// NewPasswordModifyExtendedOperationHandlerResponseWithDefaults instantiates a new PasswordModifyExtendedOperationHandlerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordModifyExtendedOperationHandlerResponseWithDefaults() *PasswordModifyExtendedOperationHandlerResponse {
	this := PasswordModifyExtendedOperationHandlerResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *PasswordModifyExtendedOperationHandlerResponse) GetSchemas() []EnumpasswordModifyExtendedOperationHandlerSchemaUrn {
	if o == nil {
		var ret []EnumpasswordModifyExtendedOperationHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *PasswordModifyExtendedOperationHandlerResponse) GetSchemasOk() ([]EnumpasswordModifyExtendedOperationHandlerSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *PasswordModifyExtendedOperationHandlerResponse) SetSchemas(v []EnumpasswordModifyExtendedOperationHandlerSchemaUrn) {
	o.Schemas = v
}

// GetIdentityMapper returns the IdentityMapper field value
func (o *PasswordModifyExtendedOperationHandlerResponse) GetIdentityMapper() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityMapper
}

// GetIdentityMapperOk returns a tuple with the IdentityMapper field value
// and a boolean to check if the value has been set.
func (o *PasswordModifyExtendedOperationHandlerResponse) GetIdentityMapperOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IdentityMapper, true
}

// SetIdentityMapper sets field value
func (o *PasswordModifyExtendedOperationHandlerResponse) SetIdentityMapper(v string) {
	o.IdentityMapper = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PasswordModifyExtendedOperationHandlerResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordModifyExtendedOperationHandlerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PasswordModifyExtendedOperationHandlerResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PasswordModifyExtendedOperationHandlerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *PasswordModifyExtendedOperationHandlerResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *PasswordModifyExtendedOperationHandlerResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *PasswordModifyExtendedOperationHandlerResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o PasswordModifyExtendedOperationHandlerResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["identityMapper"] = o.IdentityMapper
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullablePasswordModifyExtendedOperationHandlerResponse struct {
	value *PasswordModifyExtendedOperationHandlerResponse
	isSet bool
}

func (v NullablePasswordModifyExtendedOperationHandlerResponse) Get() *PasswordModifyExtendedOperationHandlerResponse {
	return v.value
}

func (v *NullablePasswordModifyExtendedOperationHandlerResponse) Set(val *PasswordModifyExtendedOperationHandlerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordModifyExtendedOperationHandlerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordModifyExtendedOperationHandlerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordModifyExtendedOperationHandlerResponse(val *PasswordModifyExtendedOperationHandlerResponse) *NullablePasswordModifyExtendedOperationHandlerResponse {
	return &NullablePasswordModifyExtendedOperationHandlerResponse{value: val, isSet: true}
}

func (v NullablePasswordModifyExtendedOperationHandlerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordModifyExtendedOperationHandlerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


