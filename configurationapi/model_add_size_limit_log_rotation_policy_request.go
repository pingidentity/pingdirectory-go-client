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

// checks if the AddSizeLimitLogRotationPolicyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddSizeLimitLogRotationPolicyRequest{}

// AddSizeLimitLogRotationPolicyRequest struct for AddSizeLimitLogRotationPolicyRequest
type AddSizeLimitLogRotationPolicyRequest struct {
	Schemas []EnumsizeLimitLogRotationPolicySchemaUrn `json:"schemas"`
	// Specifies the maximum size that a log file can reach before it is rotated.
	FileSizeLimit string `json:"fileSizeLimit"`
	// A description for this Log Rotation Policy
	Description *string `json:"description,omitempty"`
	// Name of the new Log Rotation Policy
	PolicyName string `json:"policyName"`
}

// NewAddSizeLimitLogRotationPolicyRequest instantiates a new AddSizeLimitLogRotationPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSizeLimitLogRotationPolicyRequest(schemas []EnumsizeLimitLogRotationPolicySchemaUrn, fileSizeLimit string, policyName string) *AddSizeLimitLogRotationPolicyRequest {
	this := AddSizeLimitLogRotationPolicyRequest{}
	this.Schemas = schemas
	this.FileSizeLimit = fileSizeLimit
	this.PolicyName = policyName
	return &this
}

// NewAddSizeLimitLogRotationPolicyRequestWithDefaults instantiates a new AddSizeLimitLogRotationPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSizeLimitLogRotationPolicyRequestWithDefaults() *AddSizeLimitLogRotationPolicyRequest {
	this := AddSizeLimitLogRotationPolicyRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddSizeLimitLogRotationPolicyRequest) GetSchemas() []EnumsizeLimitLogRotationPolicySchemaUrn {
	if o == nil {
		var ret []EnumsizeLimitLogRotationPolicySchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddSizeLimitLogRotationPolicyRequest) GetSchemasOk() ([]EnumsizeLimitLogRotationPolicySchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddSizeLimitLogRotationPolicyRequest) SetSchemas(v []EnumsizeLimitLogRotationPolicySchemaUrn) {
	o.Schemas = v
}

// GetFileSizeLimit returns the FileSizeLimit field value
func (o *AddSizeLimitLogRotationPolicyRequest) GetFileSizeLimit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileSizeLimit
}

// GetFileSizeLimitOk returns a tuple with the FileSizeLimit field value
// and a boolean to check if the value has been set.
func (o *AddSizeLimitLogRotationPolicyRequest) GetFileSizeLimitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileSizeLimit, true
}

// SetFileSizeLimit sets field value
func (o *AddSizeLimitLogRotationPolicyRequest) SetFileSizeLimit(v string) {
	o.FileSizeLimit = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddSizeLimitLogRotationPolicyRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSizeLimitLogRotationPolicyRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddSizeLimitLogRotationPolicyRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddSizeLimitLogRotationPolicyRequest) SetDescription(v string) {
	o.Description = &v
}

// GetPolicyName returns the PolicyName field value
func (o *AddSizeLimitLogRotationPolicyRequest) GetPolicyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value
// and a boolean to check if the value has been set.
func (o *AddSizeLimitLogRotationPolicyRequest) GetPolicyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyName, true
}

// SetPolicyName sets field value
func (o *AddSizeLimitLogRotationPolicyRequest) SetPolicyName(v string) {
	o.PolicyName = v
}

func (o AddSizeLimitLogRotationPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddSizeLimitLogRotationPolicyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["fileSizeLimit"] = o.FileSizeLimit
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["policyName"] = o.PolicyName
	return toSerialize, nil
}

type NullableAddSizeLimitLogRotationPolicyRequest struct {
	value *AddSizeLimitLogRotationPolicyRequest
	isSet bool
}

func (v NullableAddSizeLimitLogRotationPolicyRequest) Get() *AddSizeLimitLogRotationPolicyRequest {
	return v.value
}

func (v *NullableAddSizeLimitLogRotationPolicyRequest) Set(val *AddSizeLimitLogRotationPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSizeLimitLogRotationPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSizeLimitLogRotationPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSizeLimitLogRotationPolicyRequest(val *AddSizeLimitLogRotationPolicyRequest) *NullableAddSizeLimitLogRotationPolicyRequest {
	return &NullableAddSizeLimitLogRotationPolicyRequest{value: val, isSet: true}
}

func (v NullableAddSizeLimitLogRotationPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSizeLimitLogRotationPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
