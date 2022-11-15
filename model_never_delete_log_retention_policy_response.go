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

// NeverDeleteLogRetentionPolicyResponse struct for NeverDeleteLogRetentionPolicyResponse
type NeverDeleteLogRetentionPolicyResponse struct {
	// Name of the Log Retention Policy
	Id string `json:"id"`
	Schemas []EnumneverDeleteLogRetentionPolicySchemaUrn `json:"schemas"`
	// A description for this Log Retention Policy
	Description *string `json:"description,omitempty"`
}

// NewNeverDeleteLogRetentionPolicyResponse instantiates a new NeverDeleteLogRetentionPolicyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNeverDeleteLogRetentionPolicyResponse(id string, schemas []EnumneverDeleteLogRetentionPolicySchemaUrn) *NeverDeleteLogRetentionPolicyResponse {
	this := NeverDeleteLogRetentionPolicyResponse{}
	this.Id = id
	this.Schemas = schemas
	return &this
}

// NewNeverDeleteLogRetentionPolicyResponseWithDefaults instantiates a new NeverDeleteLogRetentionPolicyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNeverDeleteLogRetentionPolicyResponseWithDefaults() *NeverDeleteLogRetentionPolicyResponse {
	this := NeverDeleteLogRetentionPolicyResponse{}
	return &this
}

// GetId returns the Id field value
func (o *NeverDeleteLogRetentionPolicyResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NeverDeleteLogRetentionPolicyResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NeverDeleteLogRetentionPolicyResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *NeverDeleteLogRetentionPolicyResponse) GetSchemas() []EnumneverDeleteLogRetentionPolicySchemaUrn {
	if o == nil {
		var ret []EnumneverDeleteLogRetentionPolicySchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *NeverDeleteLogRetentionPolicyResponse) GetSchemasOk() ([]EnumneverDeleteLogRetentionPolicySchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *NeverDeleteLogRetentionPolicyResponse) SetSchemas(v []EnumneverDeleteLogRetentionPolicySchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NeverDeleteLogRetentionPolicyResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NeverDeleteLogRetentionPolicyResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NeverDeleteLogRetentionPolicyResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NeverDeleteLogRetentionPolicyResponse) SetDescription(v string) {
	o.Description = &v
}

func (o NeverDeleteLogRetentionPolicyResponse) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableNeverDeleteLogRetentionPolicyResponse struct {
	value *NeverDeleteLogRetentionPolicyResponse
	isSet bool
}

func (v NullableNeverDeleteLogRetentionPolicyResponse) Get() *NeverDeleteLogRetentionPolicyResponse {
	return v.value
}

func (v *NullableNeverDeleteLogRetentionPolicyResponse) Set(val *NeverDeleteLogRetentionPolicyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNeverDeleteLogRetentionPolicyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNeverDeleteLogRetentionPolicyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNeverDeleteLogRetentionPolicyResponse(val *NeverDeleteLogRetentionPolicyResponse) *NullableNeverDeleteLogRetentionPolicyResponse {
	return &NullableNeverDeleteLogRetentionPolicyResponse{value: val, isSet: true}
}

func (v NullableNeverDeleteLogRetentionPolicyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNeverDeleteLogRetentionPolicyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


