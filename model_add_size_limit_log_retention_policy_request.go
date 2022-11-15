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

// AddSizeLimitLogRetentionPolicyRequest struct for AddSizeLimitLogRetentionPolicyRequest
type AddSizeLimitLogRetentionPolicyRequest struct {
	// Name of the new Log Retention Policy
	PolicyName string `json:"policyName"`
	Schemas []EnumsizeLimitLogRetentionPolicySchemaUrn `json:"schemas"`
	// Specifies the maximum total disk space used by the log files.
	DiskSpaceUsed string `json:"diskSpaceUsed"`
	// A description for this Log Retention Policy
	Description *string `json:"description,omitempty"`
}

// NewAddSizeLimitLogRetentionPolicyRequest instantiates a new AddSizeLimitLogRetentionPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSizeLimitLogRetentionPolicyRequest(policyName string, schemas []EnumsizeLimitLogRetentionPolicySchemaUrn, diskSpaceUsed string) *AddSizeLimitLogRetentionPolicyRequest {
	this := AddSizeLimitLogRetentionPolicyRequest{}
	this.PolicyName = policyName
	this.Schemas = schemas
	this.DiskSpaceUsed = diskSpaceUsed
	return &this
}

// NewAddSizeLimitLogRetentionPolicyRequestWithDefaults instantiates a new AddSizeLimitLogRetentionPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSizeLimitLogRetentionPolicyRequestWithDefaults() *AddSizeLimitLogRetentionPolicyRequest {
	this := AddSizeLimitLogRetentionPolicyRequest{}
	return &this
}

// GetPolicyName returns the PolicyName field value
func (o *AddSizeLimitLogRetentionPolicyRequest) GetPolicyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value
// and a boolean to check if the value has been set.
func (o *AddSizeLimitLogRetentionPolicyRequest) GetPolicyNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PolicyName, true
}

// SetPolicyName sets field value
func (o *AddSizeLimitLogRetentionPolicyRequest) SetPolicyName(v string) {
	o.PolicyName = v
}

// GetSchemas returns the Schemas field value
func (o *AddSizeLimitLogRetentionPolicyRequest) GetSchemas() []EnumsizeLimitLogRetentionPolicySchemaUrn {
	if o == nil {
		var ret []EnumsizeLimitLogRetentionPolicySchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddSizeLimitLogRetentionPolicyRequest) GetSchemasOk() ([]EnumsizeLimitLogRetentionPolicySchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddSizeLimitLogRetentionPolicyRequest) SetSchemas(v []EnumsizeLimitLogRetentionPolicySchemaUrn) {
	o.Schemas = v
}

// GetDiskSpaceUsed returns the DiskSpaceUsed field value
func (o *AddSizeLimitLogRetentionPolicyRequest) GetDiskSpaceUsed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DiskSpaceUsed
}

// GetDiskSpaceUsedOk returns a tuple with the DiskSpaceUsed field value
// and a boolean to check if the value has been set.
func (o *AddSizeLimitLogRetentionPolicyRequest) GetDiskSpaceUsedOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DiskSpaceUsed, true
}

// SetDiskSpaceUsed sets field value
func (o *AddSizeLimitLogRetentionPolicyRequest) SetDiskSpaceUsed(v string) {
	o.DiskSpaceUsed = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddSizeLimitLogRetentionPolicyRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSizeLimitLogRetentionPolicyRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddSizeLimitLogRetentionPolicyRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddSizeLimitLogRetentionPolicyRequest) SetDescription(v string) {
	o.Description = &v
}

func (o AddSizeLimitLogRetentionPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["policyName"] = o.PolicyName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["diskSpaceUsed"] = o.DiskSpaceUsed
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableAddSizeLimitLogRetentionPolicyRequest struct {
	value *AddSizeLimitLogRetentionPolicyRequest
	isSet bool
}

func (v NullableAddSizeLimitLogRetentionPolicyRequest) Get() *AddSizeLimitLogRetentionPolicyRequest {
	return v.value
}

func (v *NullableAddSizeLimitLogRetentionPolicyRequest) Set(val *AddSizeLimitLogRetentionPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSizeLimitLogRetentionPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSizeLimitLogRetentionPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSizeLimitLogRetentionPolicyRequest(val *AddSizeLimitLogRetentionPolicyRequest) *NullableAddSizeLimitLogRetentionPolicyRequest {
	return &NullableAddSizeLimitLogRetentionPolicyRequest{value: val, isSet: true}
}

func (v NullableAddSizeLimitLogRetentionPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSizeLimitLogRetentionPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


