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

// AddReplicationAssurancePolicyRequest struct for AddReplicationAssurancePolicyRequest
type AddReplicationAssurancePolicyRequest struct {
	// Name of the new Replication Assurance Policy
	PolicyName string                                    `json:"policyName"`
	Schemas    []EnumreplicationAssurancePolicySchemaUrn `json:"schemas,omitempty"`
	// Description of the Replication Assurance Policy.
	Description *string `json:"description,omitempty"`
	// Indicates whether this Replication Assurance Policy is enabled for use in the server. If a Replication Assurance Policy is disabled, then no new operations will be associated with it.
	Enabled *bool `json:"enabled,omitempty"`
	// When multiple Replication Assurance Policies are defined, this property determines the evaluation order for finding a Replication Assurance Policy match against an operation. Policies are evaluated based on this index from least to greatest. Values of this property must be unique but not necessarily contiguous.
	EvaluationOrderIndex int32                                          `json:"evaluationOrderIndex"`
	LocalLevel           *EnumreplicationAssurancePolicyLocalLevelProp  `json:"localLevel,omitempty"`
	RemoteLevel          *EnumreplicationAssurancePolicyRemoteLevelProp `json:"remoteLevel,omitempty"`
	// Specifies the maximum length of time to wait for the replication assurance requirements to be met before timing out and replying to the client.
	Timeout string `json:"timeout"`
	// Specifies a connection criteria used to indicate which operations from clients matching this criteria use this policy. If both a connection criteria and a request criteria are specified for a policy, then both must match an operation for the policy to be assigned.
	ConnectionCriteria *string `json:"connectionCriteria,omitempty"`
	// Specifies a request criteria used to indicate which operations from clients matching this criteria use this policy. If both a connection criteria and a request criteria are specified for a policy, then both must match an operation for the policy to be assigned.
	RequestCriteria *string `json:"requestCriteria,omitempty"`
}

// NewAddReplicationAssurancePolicyRequest instantiates a new AddReplicationAssurancePolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddReplicationAssurancePolicyRequest(policyName string, evaluationOrderIndex int32, timeout string) *AddReplicationAssurancePolicyRequest {
	this := AddReplicationAssurancePolicyRequest{}
	this.PolicyName = policyName
	this.EvaluationOrderIndex = evaluationOrderIndex
	this.Timeout = timeout
	return &this
}

// NewAddReplicationAssurancePolicyRequestWithDefaults instantiates a new AddReplicationAssurancePolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddReplicationAssurancePolicyRequestWithDefaults() *AddReplicationAssurancePolicyRequest {
	this := AddReplicationAssurancePolicyRequest{}
	return &this
}

// GetPolicyName returns the PolicyName field value
func (o *AddReplicationAssurancePolicyRequest) GetPolicyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value
// and a boolean to check if the value has been set.
func (o *AddReplicationAssurancePolicyRequest) GetPolicyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyName, true
}

// SetPolicyName sets field value
func (o *AddReplicationAssurancePolicyRequest) SetPolicyName(v string) {
	o.PolicyName = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *AddReplicationAssurancePolicyRequest) GetSchemas() []EnumreplicationAssurancePolicySchemaUrn {
	if o == nil || isNil(o.Schemas) {
		var ret []EnumreplicationAssurancePolicySchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddReplicationAssurancePolicyRequest) GetSchemasOk() ([]EnumreplicationAssurancePolicySchemaUrn, bool) {
	if o == nil || isNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *AddReplicationAssurancePolicyRequest) HasSchemas() bool {
	if o != nil && !isNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumreplicationAssurancePolicySchemaUrn and assigns it to the Schemas field.
func (o *AddReplicationAssurancePolicyRequest) SetSchemas(v []EnumreplicationAssurancePolicySchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddReplicationAssurancePolicyRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddReplicationAssurancePolicyRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddReplicationAssurancePolicyRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddReplicationAssurancePolicyRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AddReplicationAssurancePolicyRequest) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddReplicationAssurancePolicyRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AddReplicationAssurancePolicyRequest) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AddReplicationAssurancePolicyRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEvaluationOrderIndex returns the EvaluationOrderIndex field value
func (o *AddReplicationAssurancePolicyRequest) GetEvaluationOrderIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EvaluationOrderIndex
}

// GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field value
// and a boolean to check if the value has been set.
func (o *AddReplicationAssurancePolicyRequest) GetEvaluationOrderIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvaluationOrderIndex, true
}

// SetEvaluationOrderIndex sets field value
func (o *AddReplicationAssurancePolicyRequest) SetEvaluationOrderIndex(v int32) {
	o.EvaluationOrderIndex = v
}

// GetLocalLevel returns the LocalLevel field value if set, zero value otherwise.
func (o *AddReplicationAssurancePolicyRequest) GetLocalLevel() EnumreplicationAssurancePolicyLocalLevelProp {
	if o == nil || isNil(o.LocalLevel) {
		var ret EnumreplicationAssurancePolicyLocalLevelProp
		return ret
	}
	return *o.LocalLevel
}

// GetLocalLevelOk returns a tuple with the LocalLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddReplicationAssurancePolicyRequest) GetLocalLevelOk() (*EnumreplicationAssurancePolicyLocalLevelProp, bool) {
	if o == nil || isNil(o.LocalLevel) {
		return nil, false
	}
	return o.LocalLevel, true
}

// HasLocalLevel returns a boolean if a field has been set.
func (o *AddReplicationAssurancePolicyRequest) HasLocalLevel() bool {
	if o != nil && !isNil(o.LocalLevel) {
		return true
	}

	return false
}

// SetLocalLevel gets a reference to the given EnumreplicationAssurancePolicyLocalLevelProp and assigns it to the LocalLevel field.
func (o *AddReplicationAssurancePolicyRequest) SetLocalLevel(v EnumreplicationAssurancePolicyLocalLevelProp) {
	o.LocalLevel = &v
}

// GetRemoteLevel returns the RemoteLevel field value if set, zero value otherwise.
func (o *AddReplicationAssurancePolicyRequest) GetRemoteLevel() EnumreplicationAssurancePolicyRemoteLevelProp {
	if o == nil || isNil(o.RemoteLevel) {
		var ret EnumreplicationAssurancePolicyRemoteLevelProp
		return ret
	}
	return *o.RemoteLevel
}

// GetRemoteLevelOk returns a tuple with the RemoteLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddReplicationAssurancePolicyRequest) GetRemoteLevelOk() (*EnumreplicationAssurancePolicyRemoteLevelProp, bool) {
	if o == nil || isNil(o.RemoteLevel) {
		return nil, false
	}
	return o.RemoteLevel, true
}

// HasRemoteLevel returns a boolean if a field has been set.
func (o *AddReplicationAssurancePolicyRequest) HasRemoteLevel() bool {
	if o != nil && !isNil(o.RemoteLevel) {
		return true
	}

	return false
}

// SetRemoteLevel gets a reference to the given EnumreplicationAssurancePolicyRemoteLevelProp and assigns it to the RemoteLevel field.
func (o *AddReplicationAssurancePolicyRequest) SetRemoteLevel(v EnumreplicationAssurancePolicyRemoteLevelProp) {
	o.RemoteLevel = &v
}

// GetTimeout returns the Timeout field value
func (o *AddReplicationAssurancePolicyRequest) GetTimeout() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value
// and a boolean to check if the value has been set.
func (o *AddReplicationAssurancePolicyRequest) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeout, true
}

// SetTimeout sets field value
func (o *AddReplicationAssurancePolicyRequest) SetTimeout(v string) {
	o.Timeout = v
}

// GetConnectionCriteria returns the ConnectionCriteria field value if set, zero value otherwise.
func (o *AddReplicationAssurancePolicyRequest) GetConnectionCriteria() string {
	if o == nil || isNil(o.ConnectionCriteria) {
		var ret string
		return ret
	}
	return *o.ConnectionCriteria
}

// GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddReplicationAssurancePolicyRequest) GetConnectionCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.ConnectionCriteria) {
		return nil, false
	}
	return o.ConnectionCriteria, true
}

// HasConnectionCriteria returns a boolean if a field has been set.
func (o *AddReplicationAssurancePolicyRequest) HasConnectionCriteria() bool {
	if o != nil && !isNil(o.ConnectionCriteria) {
		return true
	}

	return false
}

// SetConnectionCriteria gets a reference to the given string and assigns it to the ConnectionCriteria field.
func (o *AddReplicationAssurancePolicyRequest) SetConnectionCriteria(v string) {
	o.ConnectionCriteria = &v
}

// GetRequestCriteria returns the RequestCriteria field value if set, zero value otherwise.
func (o *AddReplicationAssurancePolicyRequest) GetRequestCriteria() string {
	if o == nil || isNil(o.RequestCriteria) {
		var ret string
		return ret
	}
	return *o.RequestCriteria
}

// GetRequestCriteriaOk returns a tuple with the RequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddReplicationAssurancePolicyRequest) GetRequestCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.RequestCriteria) {
		return nil, false
	}
	return o.RequestCriteria, true
}

// HasRequestCriteria returns a boolean if a field has been set.
func (o *AddReplicationAssurancePolicyRequest) HasRequestCriteria() bool {
	if o != nil && !isNil(o.RequestCriteria) {
		return true
	}

	return false
}

// SetRequestCriteria gets a reference to the given string and assigns it to the RequestCriteria field.
func (o *AddReplicationAssurancePolicyRequest) SetRequestCriteria(v string) {
	o.RequestCriteria = &v
}

func (o AddReplicationAssurancePolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["policyName"] = o.PolicyName
	}
	if !isNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["evaluationOrderIndex"] = o.EvaluationOrderIndex
	}
	if !isNil(o.LocalLevel) {
		toSerialize["localLevel"] = o.LocalLevel
	}
	if !isNil(o.RemoteLevel) {
		toSerialize["remoteLevel"] = o.RemoteLevel
	}
	if true {
		toSerialize["timeout"] = o.Timeout
	}
	if !isNil(o.ConnectionCriteria) {
		toSerialize["connectionCriteria"] = o.ConnectionCriteria
	}
	if !isNil(o.RequestCriteria) {
		toSerialize["requestCriteria"] = o.RequestCriteria
	}
	return json.Marshal(toSerialize)
}

type NullableAddReplicationAssurancePolicyRequest struct {
	value *AddReplicationAssurancePolicyRequest
	isSet bool
}

func (v NullableAddReplicationAssurancePolicyRequest) Get() *AddReplicationAssurancePolicyRequest {
	return v.value
}

func (v *NullableAddReplicationAssurancePolicyRequest) Set(val *AddReplicationAssurancePolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddReplicationAssurancePolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddReplicationAssurancePolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddReplicationAssurancePolicyRequest(val *AddReplicationAssurancePolicyRequest) *NullableAddReplicationAssurancePolicyRequest {
	return &NullableAddReplicationAssurancePolicyRequest{value: val, isSet: true}
}

func (v NullableAddReplicationAssurancePolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddReplicationAssurancePolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
