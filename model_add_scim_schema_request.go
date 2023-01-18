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

// AddScimSchemaRequest struct for AddScimSchemaRequest
type AddScimSchemaRequest struct {
	// Name of the new SCIM Schema
	SchemaName string `json:"schemaName"`
	Schemas []EnumscimSchemaSchemaUrn `json:"schemas,omitempty"`
	// A description for this SCIM Schema
	Description *string `json:"description,omitempty"`
	// The URN which identifies this SCIM Schema.
	SchemaURN string `json:"schemaURN"`
	// The human readable name for this SCIM Schema.
	DisplayName *string `json:"displayName,omitempty"`
}

// NewAddScimSchemaRequest instantiates a new AddScimSchemaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddScimSchemaRequest(schemaName string, schemaURN string) *AddScimSchemaRequest {
	this := AddScimSchemaRequest{}
	this.SchemaName = schemaName
	this.SchemaURN = schemaURN
	return &this
}

// NewAddScimSchemaRequestWithDefaults instantiates a new AddScimSchemaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddScimSchemaRequestWithDefaults() *AddScimSchemaRequest {
	this := AddScimSchemaRequest{}
	return &this
}

// GetSchemaName returns the SchemaName field value
func (o *AddScimSchemaRequest) GetSchemaName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemaName
}

// GetSchemaNameOk returns a tuple with the SchemaName field value
// and a boolean to check if the value has been set.
func (o *AddScimSchemaRequest) GetSchemaNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SchemaName, true
}

// SetSchemaName sets field value
func (o *AddScimSchemaRequest) SetSchemaName(v string) {
	o.SchemaName = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *AddScimSchemaRequest) GetSchemas() []EnumscimSchemaSchemaUrn {
	if o == nil || isNil(o.Schemas) {
		var ret []EnumscimSchemaSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScimSchemaRequest) GetSchemasOk() ([]EnumscimSchemaSchemaUrn, bool) {
	if o == nil || isNil(o.Schemas) {
    return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *AddScimSchemaRequest) HasSchemas() bool {
	if o != nil && !isNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumscimSchemaSchemaUrn and assigns it to the Schemas field.
func (o *AddScimSchemaRequest) SetSchemas(v []EnumscimSchemaSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddScimSchemaRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScimSchemaRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddScimSchemaRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddScimSchemaRequest) SetDescription(v string) {
	o.Description = &v
}

// GetSchemaURN returns the SchemaURN field value
func (o *AddScimSchemaRequest) GetSchemaURN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemaURN
}

// GetSchemaURNOk returns a tuple with the SchemaURN field value
// and a boolean to check if the value has been set.
func (o *AddScimSchemaRequest) GetSchemaURNOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SchemaURN, true
}

// SetSchemaURN sets field value
func (o *AddScimSchemaRequest) SetSchemaURN(v string) {
	o.SchemaURN = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AddScimSchemaRequest) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScimSchemaRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AddScimSchemaRequest) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AddScimSchemaRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

func (o AddScimSchemaRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemaName"] = o.SchemaName
	}
	if !isNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["schemaURN"] = o.SchemaURN
	}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	return json.Marshal(toSerialize)
}

type NullableAddScimSchemaRequest struct {
	value *AddScimSchemaRequest
	isSet bool
}

func (v NullableAddScimSchemaRequest) Get() *AddScimSchemaRequest {
	return v.value
}

func (v *NullableAddScimSchemaRequest) Set(val *AddScimSchemaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddScimSchemaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddScimSchemaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddScimSchemaRequest(val *AddScimSchemaRequest) *NullableAddScimSchemaRequest {
	return &NullableAddScimSchemaRequest{value: val, isSet: true}
}

func (v NullableAddScimSchemaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddScimSchemaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

