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

// checks if the AddNeverDeleteLogRetentionPolicyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddNeverDeleteLogRetentionPolicyRequest{}

// AddNeverDeleteLogRetentionPolicyRequest struct for AddNeverDeleteLogRetentionPolicyRequest
type AddNeverDeleteLogRetentionPolicyRequest struct {
	Schemas []EnumneverDeleteLogRetentionPolicySchemaUrn `json:"schemas"`
	// A description for this Log Retention Policy
	Description *string `json:"description,omitempty"`
	// Name of the new Log Retention Policy
	PolicyName string `json:"policyName"`
}

// NewAddNeverDeleteLogRetentionPolicyRequest instantiates a new AddNeverDeleteLogRetentionPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddNeverDeleteLogRetentionPolicyRequest(schemas []EnumneverDeleteLogRetentionPolicySchemaUrn, policyName string) *AddNeverDeleteLogRetentionPolicyRequest {
	this := AddNeverDeleteLogRetentionPolicyRequest{}
	this.Schemas = schemas
	this.PolicyName = policyName
	return &this
}

// NewAddNeverDeleteLogRetentionPolicyRequestWithDefaults instantiates a new AddNeverDeleteLogRetentionPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddNeverDeleteLogRetentionPolicyRequestWithDefaults() *AddNeverDeleteLogRetentionPolicyRequest {
	this := AddNeverDeleteLogRetentionPolicyRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddNeverDeleteLogRetentionPolicyRequest) GetSchemas() []EnumneverDeleteLogRetentionPolicySchemaUrn {
	if o == nil {
		var ret []EnumneverDeleteLogRetentionPolicySchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddNeverDeleteLogRetentionPolicyRequest) GetSchemasOk() ([]EnumneverDeleteLogRetentionPolicySchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddNeverDeleteLogRetentionPolicyRequest) SetSchemas(v []EnumneverDeleteLogRetentionPolicySchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddNeverDeleteLogRetentionPolicyRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNeverDeleteLogRetentionPolicyRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddNeverDeleteLogRetentionPolicyRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddNeverDeleteLogRetentionPolicyRequest) SetDescription(v string) {
	o.Description = &v
}

// GetPolicyName returns the PolicyName field value
func (o *AddNeverDeleteLogRetentionPolicyRequest) GetPolicyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value
// and a boolean to check if the value has been set.
func (o *AddNeverDeleteLogRetentionPolicyRequest) GetPolicyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyName, true
}

// SetPolicyName sets field value
func (o *AddNeverDeleteLogRetentionPolicyRequest) SetPolicyName(v string) {
	o.PolicyName = v
}

func (o AddNeverDeleteLogRetentionPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddNeverDeleteLogRetentionPolicyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["policyName"] = o.PolicyName
	return toSerialize, nil
}

type NullableAddNeverDeleteLogRetentionPolicyRequest struct {
	value *AddNeverDeleteLogRetentionPolicyRequest
	isSet bool
}

func (v NullableAddNeverDeleteLogRetentionPolicyRequest) Get() *AddNeverDeleteLogRetentionPolicyRequest {
	return v.value
}

func (v *NullableAddNeverDeleteLogRetentionPolicyRequest) Set(val *AddNeverDeleteLogRetentionPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddNeverDeleteLogRetentionPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddNeverDeleteLogRetentionPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddNeverDeleteLogRetentionPolicyRequest(val *AddNeverDeleteLogRetentionPolicyRequest) *NullableAddNeverDeleteLogRetentionPolicyRequest {
	return &NullableAddNeverDeleteLogRetentionPolicyRequest{value: val, isSet: true}
}

func (v NullableAddNeverDeleteLogRetentionPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddNeverDeleteLogRetentionPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
