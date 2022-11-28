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

// ReplicationAssurancePolicyResponse struct for ReplicationAssurancePolicyResponse
type ReplicationAssurancePolicyResponse struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Replication Assurance Policy
	Id string `json:"id"`
	Schemas []EnumreplicationAssurancePolicySchemaUrn `json:"schemas,omitempty"`
	// Description of the Replication Assurance Policy.
	Description *string `json:"description,omitempty"`
	// Indicates whether this Replication Assurance Policy is enabled for use in the server. If a Replication Assurance Policy is disabled, then no new operations will be associated with it.
	Enabled bool `json:"enabled"`
	// When multiple Replication Assurance Policies are defined, this property determines the evaluation order for finding a Replication Assurance Policy match against an operation. Policies are evaluated based on this index from least to greatest. Values of this property must be unique but not necessarily contiguous.
	EvaluationOrderIndex int32 `json:"evaluationOrderIndex"`
	LocalLevel EnumreplicationAssurancePolicyLocalLevelProp `json:"localLevel"`
	RemoteLevel EnumreplicationAssurancePolicyRemoteLevelProp `json:"remoteLevel"`
	// Specifies the maximum length of time to wait for the replication assurance requirements to be met before timing out and replying to the client.
	Timeout string `json:"timeout"`
	// Specifies a connection criteria used to indicate which operations from clients matching this criteria use this policy. If both a connection criteria and a request criteria are specified for a policy, then both must match an operation for the policy to be assigned.
	ConnectionCriteria *string `json:"connectionCriteria,omitempty"`
	// Specifies a request criteria used to indicate which operations from clients matching this criteria use this policy. If both a connection criteria and a request criteria are specified for a policy, then both must match an operation for the policy to be assigned.
	RequestCriteria *string `json:"requestCriteria,omitempty"`
}

// NewReplicationAssurancePolicyResponse instantiates a new ReplicationAssurancePolicyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplicationAssurancePolicyResponse(id string, enabled bool, evaluationOrderIndex int32, localLevel EnumreplicationAssurancePolicyLocalLevelProp, remoteLevel EnumreplicationAssurancePolicyRemoteLevelProp, timeout string) *ReplicationAssurancePolicyResponse {
	this := ReplicationAssurancePolicyResponse{}
	this.Id = id
	this.Enabled = enabled
	this.EvaluationOrderIndex = evaluationOrderIndex
	this.LocalLevel = localLevel
	this.RemoteLevel = remoteLevel
	this.Timeout = timeout
	return &this
}

// NewReplicationAssurancePolicyResponseWithDefaults instantiates a new ReplicationAssurancePolicyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplicationAssurancePolicyResponseWithDefaults() *ReplicationAssurancePolicyResponse {
	this := ReplicationAssurancePolicyResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ReplicationAssurancePolicyResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationAssurancePolicyResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ReplicationAssurancePolicyResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ReplicationAssurancePolicyResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ReplicationAssurancePolicyResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationAssurancePolicyResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
    return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ReplicationAssurancePolicyResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ReplicationAssurancePolicyResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *ReplicationAssurancePolicyResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReplicationAssurancePolicyResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReplicationAssurancePolicyResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *ReplicationAssurancePolicyResponse) GetSchemas() []EnumreplicationAssurancePolicySchemaUrn {
	if o == nil || isNil(o.Schemas) {
		var ret []EnumreplicationAssurancePolicySchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationAssurancePolicyResponse) GetSchemasOk() ([]EnumreplicationAssurancePolicySchemaUrn, bool) {
	if o == nil || isNil(o.Schemas) {
    return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *ReplicationAssurancePolicyResponse) HasSchemas() bool {
	if o != nil && !isNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumreplicationAssurancePolicySchemaUrn and assigns it to the Schemas field.
func (o *ReplicationAssurancePolicyResponse) SetSchemas(v []EnumreplicationAssurancePolicySchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ReplicationAssurancePolicyResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationAssurancePolicyResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ReplicationAssurancePolicyResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ReplicationAssurancePolicyResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *ReplicationAssurancePolicyResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ReplicationAssurancePolicyResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ReplicationAssurancePolicyResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEvaluationOrderIndex returns the EvaluationOrderIndex field value
func (o *ReplicationAssurancePolicyResponse) GetEvaluationOrderIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EvaluationOrderIndex
}

// GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field value
// and a boolean to check if the value has been set.
func (o *ReplicationAssurancePolicyResponse) GetEvaluationOrderIndexOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EvaluationOrderIndex, true
}

// SetEvaluationOrderIndex sets field value
func (o *ReplicationAssurancePolicyResponse) SetEvaluationOrderIndex(v int32) {
	o.EvaluationOrderIndex = v
}

// GetLocalLevel returns the LocalLevel field value
func (o *ReplicationAssurancePolicyResponse) GetLocalLevel() EnumreplicationAssurancePolicyLocalLevelProp {
	if o == nil {
		var ret EnumreplicationAssurancePolicyLocalLevelProp
		return ret
	}

	return o.LocalLevel
}

// GetLocalLevelOk returns a tuple with the LocalLevel field value
// and a boolean to check if the value has been set.
func (o *ReplicationAssurancePolicyResponse) GetLocalLevelOk() (*EnumreplicationAssurancePolicyLocalLevelProp, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LocalLevel, true
}

// SetLocalLevel sets field value
func (o *ReplicationAssurancePolicyResponse) SetLocalLevel(v EnumreplicationAssurancePolicyLocalLevelProp) {
	o.LocalLevel = v
}

// GetRemoteLevel returns the RemoteLevel field value
func (o *ReplicationAssurancePolicyResponse) GetRemoteLevel() EnumreplicationAssurancePolicyRemoteLevelProp {
	if o == nil {
		var ret EnumreplicationAssurancePolicyRemoteLevelProp
		return ret
	}

	return o.RemoteLevel
}

// GetRemoteLevelOk returns a tuple with the RemoteLevel field value
// and a boolean to check if the value has been set.
func (o *ReplicationAssurancePolicyResponse) GetRemoteLevelOk() (*EnumreplicationAssurancePolicyRemoteLevelProp, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RemoteLevel, true
}

// SetRemoteLevel sets field value
func (o *ReplicationAssurancePolicyResponse) SetRemoteLevel(v EnumreplicationAssurancePolicyRemoteLevelProp) {
	o.RemoteLevel = v
}

// GetTimeout returns the Timeout field value
func (o *ReplicationAssurancePolicyResponse) GetTimeout() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value
// and a boolean to check if the value has been set.
func (o *ReplicationAssurancePolicyResponse) GetTimeoutOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Timeout, true
}

// SetTimeout sets field value
func (o *ReplicationAssurancePolicyResponse) SetTimeout(v string) {
	o.Timeout = v
}

// GetConnectionCriteria returns the ConnectionCriteria field value if set, zero value otherwise.
func (o *ReplicationAssurancePolicyResponse) GetConnectionCriteria() string {
	if o == nil || isNil(o.ConnectionCriteria) {
		var ret string
		return ret
	}
	return *o.ConnectionCriteria
}

// GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationAssurancePolicyResponse) GetConnectionCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.ConnectionCriteria) {
    return nil, false
	}
	return o.ConnectionCriteria, true
}

// HasConnectionCriteria returns a boolean if a field has been set.
func (o *ReplicationAssurancePolicyResponse) HasConnectionCriteria() bool {
	if o != nil && !isNil(o.ConnectionCriteria) {
		return true
	}

	return false
}

// SetConnectionCriteria gets a reference to the given string and assigns it to the ConnectionCriteria field.
func (o *ReplicationAssurancePolicyResponse) SetConnectionCriteria(v string) {
	o.ConnectionCriteria = &v
}

// GetRequestCriteria returns the RequestCriteria field value if set, zero value otherwise.
func (o *ReplicationAssurancePolicyResponse) GetRequestCriteria() string {
	if o == nil || isNil(o.RequestCriteria) {
		var ret string
		return ret
	}
	return *o.RequestCriteria
}

// GetRequestCriteriaOk returns a tuple with the RequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationAssurancePolicyResponse) GetRequestCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.RequestCriteria) {
    return nil, false
	}
	return o.RequestCriteria, true
}

// HasRequestCriteria returns a boolean if a field has been set.
func (o *ReplicationAssurancePolicyResponse) HasRequestCriteria() bool {
	if o != nil && !isNil(o.RequestCriteria) {
		return true
	}

	return false
}

// SetRequestCriteria gets a reference to the given string and assigns it to the RequestCriteria field.
func (o *ReplicationAssurancePolicyResponse) SetRequestCriteria(v string) {
	o.RequestCriteria = &v
}

func (o ReplicationAssurancePolicyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["evaluationOrderIndex"] = o.EvaluationOrderIndex
	}
	if true {
		toSerialize["localLevel"] = o.LocalLevel
	}
	if true {
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

type NullableReplicationAssurancePolicyResponse struct {
	value *ReplicationAssurancePolicyResponse
	isSet bool
}

func (v NullableReplicationAssurancePolicyResponse) Get() *ReplicationAssurancePolicyResponse {
	return v.value
}

func (v *NullableReplicationAssurancePolicyResponse) Set(val *ReplicationAssurancePolicyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReplicationAssurancePolicyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReplicationAssurancePolicyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplicationAssurancePolicyResponse(val *ReplicationAssurancePolicyResponse) *NullableReplicationAssurancePolicyResponse {
	return &NullableReplicationAssurancePolicyResponse{value: val, isSet: true}
}

func (v NullableReplicationAssurancePolicyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplicationAssurancePolicyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


