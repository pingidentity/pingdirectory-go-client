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

// checks if the AddFreeDiskSpaceLogRetentionPolicyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddFreeDiskSpaceLogRetentionPolicyRequest{}

// AddFreeDiskSpaceLogRetentionPolicyRequest struct for AddFreeDiskSpaceLogRetentionPolicyRequest
type AddFreeDiskSpaceLogRetentionPolicyRequest struct {
	Schemas []EnumfreeDiskSpaceLogRetentionPolicySchemaUrn `json:"schemas"`
	// Specifies the minimum amount of free disk space that should be available on the file system on which the archived log files are stored.
	FreeDiskSpace string `json:"freeDiskSpace"`
	// A description for this Log Retention Policy
	Description *string `json:"description,omitempty"`
	// Name of the new Log Retention Policy
	PolicyName string `json:"policyName"`
}

// NewAddFreeDiskSpaceLogRetentionPolicyRequest instantiates a new AddFreeDiskSpaceLogRetentionPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddFreeDiskSpaceLogRetentionPolicyRequest(schemas []EnumfreeDiskSpaceLogRetentionPolicySchemaUrn, freeDiskSpace string, policyName string) *AddFreeDiskSpaceLogRetentionPolicyRequest {
	this := AddFreeDiskSpaceLogRetentionPolicyRequest{}
	this.Schemas = schemas
	this.FreeDiskSpace = freeDiskSpace
	this.PolicyName = policyName
	return &this
}

// NewAddFreeDiskSpaceLogRetentionPolicyRequestWithDefaults instantiates a new AddFreeDiskSpaceLogRetentionPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddFreeDiskSpaceLogRetentionPolicyRequestWithDefaults() *AddFreeDiskSpaceLogRetentionPolicyRequest {
	this := AddFreeDiskSpaceLogRetentionPolicyRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddFreeDiskSpaceLogRetentionPolicyRequest) GetSchemas() []EnumfreeDiskSpaceLogRetentionPolicySchemaUrn {
	if o == nil {
		var ret []EnumfreeDiskSpaceLogRetentionPolicySchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddFreeDiskSpaceLogRetentionPolicyRequest) GetSchemasOk() ([]EnumfreeDiskSpaceLogRetentionPolicySchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddFreeDiskSpaceLogRetentionPolicyRequest) SetSchemas(v []EnumfreeDiskSpaceLogRetentionPolicySchemaUrn) {
	o.Schemas = v
}

// GetFreeDiskSpace returns the FreeDiskSpace field value
func (o *AddFreeDiskSpaceLogRetentionPolicyRequest) GetFreeDiskSpace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FreeDiskSpace
}

// GetFreeDiskSpaceOk returns a tuple with the FreeDiskSpace field value
// and a boolean to check if the value has been set.
func (o *AddFreeDiskSpaceLogRetentionPolicyRequest) GetFreeDiskSpaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FreeDiskSpace, true
}

// SetFreeDiskSpace sets field value
func (o *AddFreeDiskSpaceLogRetentionPolicyRequest) SetFreeDiskSpace(v string) {
	o.FreeDiskSpace = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddFreeDiskSpaceLogRetentionPolicyRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFreeDiskSpaceLogRetentionPolicyRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddFreeDiskSpaceLogRetentionPolicyRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddFreeDiskSpaceLogRetentionPolicyRequest) SetDescription(v string) {
	o.Description = &v
}

// GetPolicyName returns the PolicyName field value
func (o *AddFreeDiskSpaceLogRetentionPolicyRequest) GetPolicyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value
// and a boolean to check if the value has been set.
func (o *AddFreeDiskSpaceLogRetentionPolicyRequest) GetPolicyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyName, true
}

// SetPolicyName sets field value
func (o *AddFreeDiskSpaceLogRetentionPolicyRequest) SetPolicyName(v string) {
	o.PolicyName = v
}

func (o AddFreeDiskSpaceLogRetentionPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddFreeDiskSpaceLogRetentionPolicyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["freeDiskSpace"] = o.FreeDiskSpace
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["policyName"] = o.PolicyName
	return toSerialize, nil
}

type NullableAddFreeDiskSpaceLogRetentionPolicyRequest struct {
	value *AddFreeDiskSpaceLogRetentionPolicyRequest
	isSet bool
}

func (v NullableAddFreeDiskSpaceLogRetentionPolicyRequest) Get() *AddFreeDiskSpaceLogRetentionPolicyRequest {
	return v.value
}

func (v *NullableAddFreeDiskSpaceLogRetentionPolicyRequest) Set(val *AddFreeDiskSpaceLogRetentionPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddFreeDiskSpaceLogRetentionPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddFreeDiskSpaceLogRetentionPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddFreeDiskSpaceLogRetentionPolicyRequest(val *AddFreeDiskSpaceLogRetentionPolicyRequest) *NullableAddFreeDiskSpaceLogRetentionPolicyRequest {
	return &NullableAddFreeDiskSpaceLogRetentionPolicyRequest{value: val, isSet: true}
}

func (v NullableAddFreeDiskSpaceLogRetentionPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddFreeDiskSpaceLogRetentionPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
