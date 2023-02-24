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

// AddReplicationAssuranceResultCriteriaRequest struct for AddReplicationAssuranceResultCriteriaRequest
type AddReplicationAssuranceResultCriteriaRequest struct {
	// Name of the new Result Criteria
	CriteriaName             string                                            `json:"criteriaName"`
	Schemas                  []EnumreplicationAssuranceResultCriteriaSchemaUrn `json:"schemas"`
	LocalAssuranceLevel      []EnumresultCriteriaLocalAssuranceLevelProp       `json:"localAssuranceLevel,omitempty"`
	RemoteAssuranceLevel     []EnumresultCriteriaRemoteAssuranceLevelProp      `json:"remoteAssuranceLevel,omitempty"`
	AssuranceTimeoutCriteria *EnumresultCriteriaAssuranceTimeoutCriteriaProp   `json:"assuranceTimeoutCriteria,omitempty"`
	// The value to use for performing matching based on the assurance timeout. This will be ignored if the assurance-timeout-criteria is \"any\".
	AssuranceTimeoutValue             *string                                                  `json:"assuranceTimeoutValue,omitempty"`
	ResponseDelayedByAssurance        *EnumresultCriteriaResponseDelayedByAssuranceProp        `json:"responseDelayedByAssurance,omitempty"`
	AssuranceBehaviorAlteredByControl *EnumresultCriteriaAssuranceBehaviorAlteredByControlProp `json:"assuranceBehaviorAlteredByControl,omitempty"`
	AssuranceSatisfied                *EnumresultCriteriaAssuranceSatisfiedProp                `json:"assuranceSatisfied,omitempty"`
	// A description for this Result Criteria
	Description *string `json:"description,omitempty"`
}

// NewAddReplicationAssuranceResultCriteriaRequest instantiates a new AddReplicationAssuranceResultCriteriaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddReplicationAssuranceResultCriteriaRequest(criteriaName string, schemas []EnumreplicationAssuranceResultCriteriaSchemaUrn) *AddReplicationAssuranceResultCriteriaRequest {
	this := AddReplicationAssuranceResultCriteriaRequest{}
	this.CriteriaName = criteriaName
	this.Schemas = schemas
	return &this
}

// NewAddReplicationAssuranceResultCriteriaRequestWithDefaults instantiates a new AddReplicationAssuranceResultCriteriaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddReplicationAssuranceResultCriteriaRequestWithDefaults() *AddReplicationAssuranceResultCriteriaRequest {
	this := AddReplicationAssuranceResultCriteriaRequest{}
	return &this
}

// GetCriteriaName returns the CriteriaName field value
func (o *AddReplicationAssuranceResultCriteriaRequest) GetCriteriaName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CriteriaName
}

// GetCriteriaNameOk returns a tuple with the CriteriaName field value
// and a boolean to check if the value has been set.
func (o *AddReplicationAssuranceResultCriteriaRequest) GetCriteriaNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CriteriaName, true
}

// SetCriteriaName sets field value
func (o *AddReplicationAssuranceResultCriteriaRequest) SetCriteriaName(v string) {
	o.CriteriaName = v
}

// GetSchemas returns the Schemas field value
func (o *AddReplicationAssuranceResultCriteriaRequest) GetSchemas() []EnumreplicationAssuranceResultCriteriaSchemaUrn {
	if o == nil {
		var ret []EnumreplicationAssuranceResultCriteriaSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddReplicationAssuranceResultCriteriaRequest) GetSchemasOk() ([]EnumreplicationAssuranceResultCriteriaSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddReplicationAssuranceResultCriteriaRequest) SetSchemas(v []EnumreplicationAssuranceResultCriteriaSchemaUrn) {
	o.Schemas = v
}

// GetLocalAssuranceLevel returns the LocalAssuranceLevel field value if set, zero value otherwise.
func (o *AddReplicationAssuranceResultCriteriaRequest) GetLocalAssuranceLevel() []EnumresultCriteriaLocalAssuranceLevelProp {
	if o == nil || isNil(o.LocalAssuranceLevel) {
		var ret []EnumresultCriteriaLocalAssuranceLevelProp
		return ret
	}
	return o.LocalAssuranceLevel
}

// GetLocalAssuranceLevelOk returns a tuple with the LocalAssuranceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddReplicationAssuranceResultCriteriaRequest) GetLocalAssuranceLevelOk() ([]EnumresultCriteriaLocalAssuranceLevelProp, bool) {
	if o == nil || isNil(o.LocalAssuranceLevel) {
		return nil, false
	}
	return o.LocalAssuranceLevel, true
}

// HasLocalAssuranceLevel returns a boolean if a field has been set.
func (o *AddReplicationAssuranceResultCriteriaRequest) HasLocalAssuranceLevel() bool {
	if o != nil && !isNil(o.LocalAssuranceLevel) {
		return true
	}

	return false
}

// SetLocalAssuranceLevel gets a reference to the given []EnumresultCriteriaLocalAssuranceLevelProp and assigns it to the LocalAssuranceLevel field.
func (o *AddReplicationAssuranceResultCriteriaRequest) SetLocalAssuranceLevel(v []EnumresultCriteriaLocalAssuranceLevelProp) {
	o.LocalAssuranceLevel = v
}

// GetRemoteAssuranceLevel returns the RemoteAssuranceLevel field value if set, zero value otherwise.
func (o *AddReplicationAssuranceResultCriteriaRequest) GetRemoteAssuranceLevel() []EnumresultCriteriaRemoteAssuranceLevelProp {
	if o == nil || isNil(o.RemoteAssuranceLevel) {
		var ret []EnumresultCriteriaRemoteAssuranceLevelProp
		return ret
	}
	return o.RemoteAssuranceLevel
}

// GetRemoteAssuranceLevelOk returns a tuple with the RemoteAssuranceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddReplicationAssuranceResultCriteriaRequest) GetRemoteAssuranceLevelOk() ([]EnumresultCriteriaRemoteAssuranceLevelProp, bool) {
	if o == nil || isNil(o.RemoteAssuranceLevel) {
		return nil, false
	}
	return o.RemoteAssuranceLevel, true
}

// HasRemoteAssuranceLevel returns a boolean if a field has been set.
func (o *AddReplicationAssuranceResultCriteriaRequest) HasRemoteAssuranceLevel() bool {
	if o != nil && !isNil(o.RemoteAssuranceLevel) {
		return true
	}

	return false
}

// SetRemoteAssuranceLevel gets a reference to the given []EnumresultCriteriaRemoteAssuranceLevelProp and assigns it to the RemoteAssuranceLevel field.
func (o *AddReplicationAssuranceResultCriteriaRequest) SetRemoteAssuranceLevel(v []EnumresultCriteriaRemoteAssuranceLevelProp) {
	o.RemoteAssuranceLevel = v
}

// GetAssuranceTimeoutCriteria returns the AssuranceTimeoutCriteria field value if set, zero value otherwise.
func (o *AddReplicationAssuranceResultCriteriaRequest) GetAssuranceTimeoutCriteria() EnumresultCriteriaAssuranceTimeoutCriteriaProp {
	if o == nil || isNil(o.AssuranceTimeoutCriteria) {
		var ret EnumresultCriteriaAssuranceTimeoutCriteriaProp
		return ret
	}
	return *o.AssuranceTimeoutCriteria
}

// GetAssuranceTimeoutCriteriaOk returns a tuple with the AssuranceTimeoutCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddReplicationAssuranceResultCriteriaRequest) GetAssuranceTimeoutCriteriaOk() (*EnumresultCriteriaAssuranceTimeoutCriteriaProp, bool) {
	if o == nil || isNil(o.AssuranceTimeoutCriteria) {
		return nil, false
	}
	return o.AssuranceTimeoutCriteria, true
}

// HasAssuranceTimeoutCriteria returns a boolean if a field has been set.
func (o *AddReplicationAssuranceResultCriteriaRequest) HasAssuranceTimeoutCriteria() bool {
	if o != nil && !isNil(o.AssuranceTimeoutCriteria) {
		return true
	}

	return false
}

// SetAssuranceTimeoutCriteria gets a reference to the given EnumresultCriteriaAssuranceTimeoutCriteriaProp and assigns it to the AssuranceTimeoutCriteria field.
func (o *AddReplicationAssuranceResultCriteriaRequest) SetAssuranceTimeoutCriteria(v EnumresultCriteriaAssuranceTimeoutCriteriaProp) {
	o.AssuranceTimeoutCriteria = &v
}

// GetAssuranceTimeoutValue returns the AssuranceTimeoutValue field value if set, zero value otherwise.
func (o *AddReplicationAssuranceResultCriteriaRequest) GetAssuranceTimeoutValue() string {
	if o == nil || isNil(o.AssuranceTimeoutValue) {
		var ret string
		return ret
	}
	return *o.AssuranceTimeoutValue
}

// GetAssuranceTimeoutValueOk returns a tuple with the AssuranceTimeoutValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddReplicationAssuranceResultCriteriaRequest) GetAssuranceTimeoutValueOk() (*string, bool) {
	if o == nil || isNil(o.AssuranceTimeoutValue) {
		return nil, false
	}
	return o.AssuranceTimeoutValue, true
}

// HasAssuranceTimeoutValue returns a boolean if a field has been set.
func (o *AddReplicationAssuranceResultCriteriaRequest) HasAssuranceTimeoutValue() bool {
	if o != nil && !isNil(o.AssuranceTimeoutValue) {
		return true
	}

	return false
}

// SetAssuranceTimeoutValue gets a reference to the given string and assigns it to the AssuranceTimeoutValue field.
func (o *AddReplicationAssuranceResultCriteriaRequest) SetAssuranceTimeoutValue(v string) {
	o.AssuranceTimeoutValue = &v
}

// GetResponseDelayedByAssurance returns the ResponseDelayedByAssurance field value if set, zero value otherwise.
func (o *AddReplicationAssuranceResultCriteriaRequest) GetResponseDelayedByAssurance() EnumresultCriteriaResponseDelayedByAssuranceProp {
	if o == nil || isNil(o.ResponseDelayedByAssurance) {
		var ret EnumresultCriteriaResponseDelayedByAssuranceProp
		return ret
	}
	return *o.ResponseDelayedByAssurance
}

// GetResponseDelayedByAssuranceOk returns a tuple with the ResponseDelayedByAssurance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddReplicationAssuranceResultCriteriaRequest) GetResponseDelayedByAssuranceOk() (*EnumresultCriteriaResponseDelayedByAssuranceProp, bool) {
	if o == nil || isNil(o.ResponseDelayedByAssurance) {
		return nil, false
	}
	return o.ResponseDelayedByAssurance, true
}

// HasResponseDelayedByAssurance returns a boolean if a field has been set.
func (o *AddReplicationAssuranceResultCriteriaRequest) HasResponseDelayedByAssurance() bool {
	if o != nil && !isNil(o.ResponseDelayedByAssurance) {
		return true
	}

	return false
}

// SetResponseDelayedByAssurance gets a reference to the given EnumresultCriteriaResponseDelayedByAssuranceProp and assigns it to the ResponseDelayedByAssurance field.
func (o *AddReplicationAssuranceResultCriteriaRequest) SetResponseDelayedByAssurance(v EnumresultCriteriaResponseDelayedByAssuranceProp) {
	o.ResponseDelayedByAssurance = &v
}

// GetAssuranceBehaviorAlteredByControl returns the AssuranceBehaviorAlteredByControl field value if set, zero value otherwise.
func (o *AddReplicationAssuranceResultCriteriaRequest) GetAssuranceBehaviorAlteredByControl() EnumresultCriteriaAssuranceBehaviorAlteredByControlProp {
	if o == nil || isNil(o.AssuranceBehaviorAlteredByControl) {
		var ret EnumresultCriteriaAssuranceBehaviorAlteredByControlProp
		return ret
	}
	return *o.AssuranceBehaviorAlteredByControl
}

// GetAssuranceBehaviorAlteredByControlOk returns a tuple with the AssuranceBehaviorAlteredByControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddReplicationAssuranceResultCriteriaRequest) GetAssuranceBehaviorAlteredByControlOk() (*EnumresultCriteriaAssuranceBehaviorAlteredByControlProp, bool) {
	if o == nil || isNil(o.AssuranceBehaviorAlteredByControl) {
		return nil, false
	}
	return o.AssuranceBehaviorAlteredByControl, true
}

// HasAssuranceBehaviorAlteredByControl returns a boolean if a field has been set.
func (o *AddReplicationAssuranceResultCriteriaRequest) HasAssuranceBehaviorAlteredByControl() bool {
	if o != nil && !isNil(o.AssuranceBehaviorAlteredByControl) {
		return true
	}

	return false
}

// SetAssuranceBehaviorAlteredByControl gets a reference to the given EnumresultCriteriaAssuranceBehaviorAlteredByControlProp and assigns it to the AssuranceBehaviorAlteredByControl field.
func (o *AddReplicationAssuranceResultCriteriaRequest) SetAssuranceBehaviorAlteredByControl(v EnumresultCriteriaAssuranceBehaviorAlteredByControlProp) {
	o.AssuranceBehaviorAlteredByControl = &v
}

// GetAssuranceSatisfied returns the AssuranceSatisfied field value if set, zero value otherwise.
func (o *AddReplicationAssuranceResultCriteriaRequest) GetAssuranceSatisfied() EnumresultCriteriaAssuranceSatisfiedProp {
	if o == nil || isNil(o.AssuranceSatisfied) {
		var ret EnumresultCriteriaAssuranceSatisfiedProp
		return ret
	}
	return *o.AssuranceSatisfied
}

// GetAssuranceSatisfiedOk returns a tuple with the AssuranceSatisfied field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddReplicationAssuranceResultCriteriaRequest) GetAssuranceSatisfiedOk() (*EnumresultCriteriaAssuranceSatisfiedProp, bool) {
	if o == nil || isNil(o.AssuranceSatisfied) {
		return nil, false
	}
	return o.AssuranceSatisfied, true
}

// HasAssuranceSatisfied returns a boolean if a field has been set.
func (o *AddReplicationAssuranceResultCriteriaRequest) HasAssuranceSatisfied() bool {
	if o != nil && !isNil(o.AssuranceSatisfied) {
		return true
	}

	return false
}

// SetAssuranceSatisfied gets a reference to the given EnumresultCriteriaAssuranceSatisfiedProp and assigns it to the AssuranceSatisfied field.
func (o *AddReplicationAssuranceResultCriteriaRequest) SetAssuranceSatisfied(v EnumresultCriteriaAssuranceSatisfiedProp) {
	o.AssuranceSatisfied = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddReplicationAssuranceResultCriteriaRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddReplicationAssuranceResultCriteriaRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddReplicationAssuranceResultCriteriaRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddReplicationAssuranceResultCriteriaRequest) SetDescription(v string) {
	o.Description = &v
}

func (o AddReplicationAssuranceResultCriteriaRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["criteriaName"] = o.CriteriaName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.LocalAssuranceLevel) {
		toSerialize["localAssuranceLevel"] = o.LocalAssuranceLevel
	}
	if !isNil(o.RemoteAssuranceLevel) {
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

type NullableAddReplicationAssuranceResultCriteriaRequest struct {
	value *AddReplicationAssuranceResultCriteriaRequest
	isSet bool
}

func (v NullableAddReplicationAssuranceResultCriteriaRequest) Get() *AddReplicationAssuranceResultCriteriaRequest {
	return v.value
}

func (v *NullableAddReplicationAssuranceResultCriteriaRequest) Set(val *AddReplicationAssuranceResultCriteriaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddReplicationAssuranceResultCriteriaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddReplicationAssuranceResultCriteriaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddReplicationAssuranceResultCriteriaRequest(val *AddReplicationAssuranceResultCriteriaRequest) *NullableAddReplicationAssuranceResultCriteriaRequest {
	return &NullableAddReplicationAssuranceResultCriteriaRequest{value: val, isSet: true}
}

func (v NullableAddReplicationAssuranceResultCriteriaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddReplicationAssuranceResultCriteriaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
