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

// ReplicationAssuranceResultCriteriaResponse struct for ReplicationAssuranceResultCriteriaResponse
type ReplicationAssuranceResultCriteriaResponse struct {
	// Name of the Result Criteria
	Id string `json:"id"`
	Schemas []EnumreplicationAssuranceResultCriteriaSchemaUrn `json:"schemas"`
	LocalAssuranceLevel []EnumresultCriteriaLocalAssuranceLevelProp `json:"localAssuranceLevel"`
	RemoteAssuranceLevel []EnumresultCriteriaRemoteAssuranceLevelProp `json:"remoteAssuranceLevel"`
	AssuranceTimeoutCriteria *EnumresultCriteriaAssuranceTimeoutCriteriaProp `json:"assuranceTimeoutCriteria,omitempty"`
	// The value to use for performing matching based on the assurance timeout. This will be ignored if the assurance-timeout-criteria is \"any\".
	AssuranceTimeoutValue *string `json:"assuranceTimeoutValue,omitempty"`
	ResponseDelayedByAssurance *EnumresultCriteriaResponseDelayedByAssuranceProp `json:"responseDelayedByAssurance,omitempty"`
	AssuranceBehaviorAlteredByControl *EnumresultCriteriaAssuranceBehaviorAlteredByControlProp `json:"assuranceBehaviorAlteredByControl,omitempty"`
	AssuranceSatisfied *EnumresultCriteriaAssuranceSatisfiedProp `json:"assuranceSatisfied,omitempty"`
	// A description for this Result Criteria
	Description *string `json:"description,omitempty"`
}

// NewReplicationAssuranceResultCriteriaResponse instantiates a new ReplicationAssuranceResultCriteriaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplicationAssuranceResultCriteriaResponse(id string, schemas []EnumreplicationAssuranceResultCriteriaSchemaUrn, localAssuranceLevel []EnumresultCriteriaLocalAssuranceLevelProp, remoteAssuranceLevel []EnumresultCriteriaRemoteAssuranceLevelProp) *ReplicationAssuranceResultCriteriaResponse {
	this := ReplicationAssuranceResultCriteriaResponse{}
	this.Id = id
	this.Schemas = schemas
	this.LocalAssuranceLevel = localAssuranceLevel
	this.RemoteAssuranceLevel = remoteAssuranceLevel
	return &this
}

// NewReplicationAssuranceResultCriteriaResponseWithDefaults instantiates a new ReplicationAssuranceResultCriteriaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplicationAssuranceResultCriteriaResponseWithDefaults() *ReplicationAssuranceResultCriteriaResponse {
	this := ReplicationAssuranceResultCriteriaResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ReplicationAssuranceResultCriteriaResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReplicationAssuranceResultCriteriaResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReplicationAssuranceResultCriteriaResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *ReplicationAssuranceResultCriteriaResponse) GetSchemas() []EnumreplicationAssuranceResultCriteriaSchemaUrn {
	if o == nil {
		var ret []EnumreplicationAssuranceResultCriteriaSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ReplicationAssuranceResultCriteriaResponse) GetSchemasOk() ([]EnumreplicationAssuranceResultCriteriaSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ReplicationAssuranceResultCriteriaResponse) SetSchemas(v []EnumreplicationAssuranceResultCriteriaSchemaUrn) {
	o.Schemas = v
}

// GetLocalAssuranceLevel returns the LocalAssuranceLevel field value
func (o *ReplicationAssuranceResultCriteriaResponse) GetLocalAssuranceLevel() []EnumresultCriteriaLocalAssuranceLevelProp {
	if o == nil {
		var ret []EnumresultCriteriaLocalAssuranceLevelProp
		return ret
	}

	return o.LocalAssuranceLevel
}

// GetLocalAssuranceLevelOk returns a tuple with the LocalAssuranceLevel field value
// and a boolean to check if the value has been set.
func (o *ReplicationAssuranceResultCriteriaResponse) GetLocalAssuranceLevelOk() ([]EnumresultCriteriaLocalAssuranceLevelProp, bool) {
	if o == nil {
    return nil, false
	}
	return o.LocalAssuranceLevel, true
}

// SetLocalAssuranceLevel sets field value
func (o *ReplicationAssuranceResultCriteriaResponse) SetLocalAssuranceLevel(v []EnumresultCriteriaLocalAssuranceLevelProp) {
	o.LocalAssuranceLevel = v
}

// GetRemoteAssuranceLevel returns the RemoteAssuranceLevel field value
func (o *ReplicationAssuranceResultCriteriaResponse) GetRemoteAssuranceLevel() []EnumresultCriteriaRemoteAssuranceLevelProp {
	if o == nil {
		var ret []EnumresultCriteriaRemoteAssuranceLevelProp
		return ret
	}

	return o.RemoteAssuranceLevel
}

// GetRemoteAssuranceLevelOk returns a tuple with the RemoteAssuranceLevel field value
// and a boolean to check if the value has been set.
func (o *ReplicationAssuranceResultCriteriaResponse) GetRemoteAssuranceLevelOk() ([]EnumresultCriteriaRemoteAssuranceLevelProp, bool) {
	if o == nil {
    return nil, false
	}
	return o.RemoteAssuranceLevel, true
}

// SetRemoteAssuranceLevel sets field value
func (o *ReplicationAssuranceResultCriteriaResponse) SetRemoteAssuranceLevel(v []EnumresultCriteriaRemoteAssuranceLevelProp) {
	o.RemoteAssuranceLevel = v
}

// GetAssuranceTimeoutCriteria returns the AssuranceTimeoutCriteria field value if set, zero value otherwise.
func (o *ReplicationAssuranceResultCriteriaResponse) GetAssuranceTimeoutCriteria() EnumresultCriteriaAssuranceTimeoutCriteriaProp {
	if o == nil || isNil(o.AssuranceTimeoutCriteria) {
		var ret EnumresultCriteriaAssuranceTimeoutCriteriaProp
		return ret
	}
	return *o.AssuranceTimeoutCriteria
}

// GetAssuranceTimeoutCriteriaOk returns a tuple with the AssuranceTimeoutCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationAssuranceResultCriteriaResponse) GetAssuranceTimeoutCriteriaOk() (*EnumresultCriteriaAssuranceTimeoutCriteriaProp, bool) {
	if o == nil || isNil(o.AssuranceTimeoutCriteria) {
    return nil, false
	}
	return o.AssuranceTimeoutCriteria, true
}

// HasAssuranceTimeoutCriteria returns a boolean if a field has been set.
func (o *ReplicationAssuranceResultCriteriaResponse) HasAssuranceTimeoutCriteria() bool {
	if o != nil && !isNil(o.AssuranceTimeoutCriteria) {
		return true
	}

	return false
}

// SetAssuranceTimeoutCriteria gets a reference to the given EnumresultCriteriaAssuranceTimeoutCriteriaProp and assigns it to the AssuranceTimeoutCriteria field.
func (o *ReplicationAssuranceResultCriteriaResponse) SetAssuranceTimeoutCriteria(v EnumresultCriteriaAssuranceTimeoutCriteriaProp) {
	o.AssuranceTimeoutCriteria = &v
}

// GetAssuranceTimeoutValue returns the AssuranceTimeoutValue field value if set, zero value otherwise.
func (o *ReplicationAssuranceResultCriteriaResponse) GetAssuranceTimeoutValue() string {
	if o == nil || isNil(o.AssuranceTimeoutValue) {
		var ret string
		return ret
	}
	return *o.AssuranceTimeoutValue
}

// GetAssuranceTimeoutValueOk returns a tuple with the AssuranceTimeoutValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationAssuranceResultCriteriaResponse) GetAssuranceTimeoutValueOk() (*string, bool) {
	if o == nil || isNil(o.AssuranceTimeoutValue) {
    return nil, false
	}
	return o.AssuranceTimeoutValue, true
}

// HasAssuranceTimeoutValue returns a boolean if a field has been set.
func (o *ReplicationAssuranceResultCriteriaResponse) HasAssuranceTimeoutValue() bool {
	if o != nil && !isNil(o.AssuranceTimeoutValue) {
		return true
	}

	return false
}

// SetAssuranceTimeoutValue gets a reference to the given string and assigns it to the AssuranceTimeoutValue field.
func (o *ReplicationAssuranceResultCriteriaResponse) SetAssuranceTimeoutValue(v string) {
	o.AssuranceTimeoutValue = &v
}

// GetResponseDelayedByAssurance returns the ResponseDelayedByAssurance field value if set, zero value otherwise.
func (o *ReplicationAssuranceResultCriteriaResponse) GetResponseDelayedByAssurance() EnumresultCriteriaResponseDelayedByAssuranceProp {
	if o == nil || isNil(o.ResponseDelayedByAssurance) {
		var ret EnumresultCriteriaResponseDelayedByAssuranceProp
		return ret
	}
	return *o.ResponseDelayedByAssurance
}

// GetResponseDelayedByAssuranceOk returns a tuple with the ResponseDelayedByAssurance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationAssuranceResultCriteriaResponse) GetResponseDelayedByAssuranceOk() (*EnumresultCriteriaResponseDelayedByAssuranceProp, bool) {
	if o == nil || isNil(o.ResponseDelayedByAssurance) {
    return nil, false
	}
	return o.ResponseDelayedByAssurance, true
}

// HasResponseDelayedByAssurance returns a boolean if a field has been set.
func (o *ReplicationAssuranceResultCriteriaResponse) HasResponseDelayedByAssurance() bool {
	if o != nil && !isNil(o.ResponseDelayedByAssurance) {
		return true
	}

	return false
}

// SetResponseDelayedByAssurance gets a reference to the given EnumresultCriteriaResponseDelayedByAssuranceProp and assigns it to the ResponseDelayedByAssurance field.
func (o *ReplicationAssuranceResultCriteriaResponse) SetResponseDelayedByAssurance(v EnumresultCriteriaResponseDelayedByAssuranceProp) {
	o.ResponseDelayedByAssurance = &v
}

// GetAssuranceBehaviorAlteredByControl returns the AssuranceBehaviorAlteredByControl field value if set, zero value otherwise.
func (o *ReplicationAssuranceResultCriteriaResponse) GetAssuranceBehaviorAlteredByControl() EnumresultCriteriaAssuranceBehaviorAlteredByControlProp {
	if o == nil || isNil(o.AssuranceBehaviorAlteredByControl) {
		var ret EnumresultCriteriaAssuranceBehaviorAlteredByControlProp
		return ret
	}
	return *o.AssuranceBehaviorAlteredByControl
}

// GetAssuranceBehaviorAlteredByControlOk returns a tuple with the AssuranceBehaviorAlteredByControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationAssuranceResultCriteriaResponse) GetAssuranceBehaviorAlteredByControlOk() (*EnumresultCriteriaAssuranceBehaviorAlteredByControlProp, bool) {
	if o == nil || isNil(o.AssuranceBehaviorAlteredByControl) {
    return nil, false
	}
	return o.AssuranceBehaviorAlteredByControl, true
}

// HasAssuranceBehaviorAlteredByControl returns a boolean if a field has been set.
func (o *ReplicationAssuranceResultCriteriaResponse) HasAssuranceBehaviorAlteredByControl() bool {
	if o != nil && !isNil(o.AssuranceBehaviorAlteredByControl) {
		return true
	}

	return false
}

// SetAssuranceBehaviorAlteredByControl gets a reference to the given EnumresultCriteriaAssuranceBehaviorAlteredByControlProp and assigns it to the AssuranceBehaviorAlteredByControl field.
func (o *ReplicationAssuranceResultCriteriaResponse) SetAssuranceBehaviorAlteredByControl(v EnumresultCriteriaAssuranceBehaviorAlteredByControlProp) {
	o.AssuranceBehaviorAlteredByControl = &v
}

// GetAssuranceSatisfied returns the AssuranceSatisfied field value if set, zero value otherwise.
func (o *ReplicationAssuranceResultCriteriaResponse) GetAssuranceSatisfied() EnumresultCriteriaAssuranceSatisfiedProp {
	if o == nil || isNil(o.AssuranceSatisfied) {
		var ret EnumresultCriteriaAssuranceSatisfiedProp
		return ret
	}
	return *o.AssuranceSatisfied
}

// GetAssuranceSatisfiedOk returns a tuple with the AssuranceSatisfied field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationAssuranceResultCriteriaResponse) GetAssuranceSatisfiedOk() (*EnumresultCriteriaAssuranceSatisfiedProp, bool) {
	if o == nil || isNil(o.AssuranceSatisfied) {
    return nil, false
	}
	return o.AssuranceSatisfied, true
}

// HasAssuranceSatisfied returns a boolean if a field has been set.
func (o *ReplicationAssuranceResultCriteriaResponse) HasAssuranceSatisfied() bool {
	if o != nil && !isNil(o.AssuranceSatisfied) {
		return true
	}

	return false
}

// SetAssuranceSatisfied gets a reference to the given EnumresultCriteriaAssuranceSatisfiedProp and assigns it to the AssuranceSatisfied field.
func (o *ReplicationAssuranceResultCriteriaResponse) SetAssuranceSatisfied(v EnumresultCriteriaAssuranceSatisfiedProp) {
	o.AssuranceSatisfied = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ReplicationAssuranceResultCriteriaResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationAssuranceResultCriteriaResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ReplicationAssuranceResultCriteriaResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ReplicationAssuranceResultCriteriaResponse) SetDescription(v string) {
	o.Description = &v
}

func (o ReplicationAssuranceResultCriteriaResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["localAssuranceLevel"] = o.LocalAssuranceLevel
	}
	if true {
		toSerialize["remoteAssuranceLevel"] = o.RemoteAssuranceLevel
	}
	if !isNil(o.AssuranceTimeoutCriteria) {
		toSerialize["assuranceTimeoutCriteria"] = o.AssuranceTimeoutCriteria
	}
	if !isNil(o.AssuranceTimeoutValue) {
		toSerialize["assuranceTimeoutValue"] = o.AssuranceTimeoutValue
	}
	if !isNil(o.ResponseDelayedByAssurance) {
		toSerialize["responseDelayedByAssurance"] = o.ResponseDelayedByAssurance
	}
	if !isNil(o.AssuranceBehaviorAlteredByControl) {
		toSerialize["assuranceBehaviorAlteredByControl"] = o.AssuranceBehaviorAlteredByControl
	}
	if !isNil(o.AssuranceSatisfied) {
		toSerialize["assuranceSatisfied"] = o.AssuranceSatisfied
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableReplicationAssuranceResultCriteriaResponse struct {
	value *ReplicationAssuranceResultCriteriaResponse
	isSet bool
}

func (v NullableReplicationAssuranceResultCriteriaResponse) Get() *ReplicationAssuranceResultCriteriaResponse {
	return v.value
}

func (v *NullableReplicationAssuranceResultCriteriaResponse) Set(val *ReplicationAssuranceResultCriteriaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReplicationAssuranceResultCriteriaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReplicationAssuranceResultCriteriaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplicationAssuranceResultCriteriaResponse(val *ReplicationAssuranceResultCriteriaResponse) *NullableReplicationAssuranceResultCriteriaResponse {
	return &NullableReplicationAssuranceResultCriteriaResponse{value: val, isSet: true}
}

func (v NullableReplicationAssuranceResultCriteriaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplicationAssuranceResultCriteriaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


