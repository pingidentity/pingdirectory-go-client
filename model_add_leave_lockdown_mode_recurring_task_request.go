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

// AddLeaveLockdownModeRecurringTaskRequest struct for AddLeaveLockdownModeRecurringTaskRequest
type AddLeaveLockdownModeRecurringTaskRequest struct {
	// Name of the new Recurring Task
	TaskName string `json:"taskName"`
	Schemas []EnumleaveLockdownModeRecurringTaskSchemaUrn `json:"schemas"`
	// The reason that the server is being taken out of in lockdown mode.
	Reason *string `json:"reason,omitempty"`
	// A description for this Recurring Task
	Description *string `json:"description,omitempty"`
	// Indicates whether an instance of this Recurring Task should be canceled if the task immediately before it in the recurring task chain fails to complete successfully (including if it is canceled by an administrator before it starts or while it is running).
	CancelOnTaskDependencyFailure *bool `json:"cancelOnTaskDependencyFailure,omitempty"`
	// The email addresses to which a message should be sent whenever an instance of this Recurring Task starts running. If this option is used, then at least one smtp-server must be configured in the global configuration.
	EmailOnStart []string `json:"emailOnStart,omitempty"`
	// The email addresses to which a message should be sent whenever an instance of this Recurring Task completes successfully. If this option is used, then at least one smtp-server must be configured in the global configuration.
	EmailOnSuccess []string `json:"emailOnSuccess,omitempty"`
	// The email addresses to which a message should be sent if an instance of this Recurring Task fails to complete successfully. If this option is used, then at least one smtp-server must be configured in the global configuration.
	EmailOnFailure []string `json:"emailOnFailure,omitempty"`
	// Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task starts running.
	AlertOnStart *bool `json:"alertOnStart,omitempty"`
	// Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task completes successfully.
	AlertOnSuccess *bool `json:"alertOnSuccess,omitempty"`
	// Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task fails to complete successfully.
	AlertOnFailure *bool `json:"alertOnFailure,omitempty"`
}

// NewAddLeaveLockdownModeRecurringTaskRequest instantiates a new AddLeaveLockdownModeRecurringTaskRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddLeaveLockdownModeRecurringTaskRequest(taskName string, schemas []EnumleaveLockdownModeRecurringTaskSchemaUrn) *AddLeaveLockdownModeRecurringTaskRequest {
	this := AddLeaveLockdownModeRecurringTaskRequest{}
	this.TaskName = taskName
	this.Schemas = schemas
	return &this
}

// NewAddLeaveLockdownModeRecurringTaskRequestWithDefaults instantiates a new AddLeaveLockdownModeRecurringTaskRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddLeaveLockdownModeRecurringTaskRequestWithDefaults() *AddLeaveLockdownModeRecurringTaskRequest {
	this := AddLeaveLockdownModeRecurringTaskRequest{}
	return &this
}

// GetTaskName returns the TaskName field value
func (o *AddLeaveLockdownModeRecurringTaskRequest) GetTaskName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaskName
}

// GetTaskNameOk returns a tuple with the TaskName field value
// and a boolean to check if the value has been set.
func (o *AddLeaveLockdownModeRecurringTaskRequest) GetTaskNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TaskName, true
}

// SetTaskName sets field value
func (o *AddLeaveLockdownModeRecurringTaskRequest) SetTaskName(v string) {
	o.TaskName = v
}

// GetSchemas returns the Schemas field value
func (o *AddLeaveLockdownModeRecurringTaskRequest) GetSchemas() []EnumleaveLockdownModeRecurringTaskSchemaUrn {
	if o == nil {
		var ret []EnumleaveLockdownModeRecurringTaskSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddLeaveLockdownModeRecurringTaskRequest) GetSchemasOk() ([]EnumleaveLockdownModeRecurringTaskSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddLeaveLockdownModeRecurringTaskRequest) SetSchemas(v []EnumleaveLockdownModeRecurringTaskSchemaUrn) {
	o.Schemas = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *AddLeaveLockdownModeRecurringTaskRequest) GetReason() string {
	if o == nil || isNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLeaveLockdownModeRecurringTaskRequest) GetReasonOk() (*string, bool) {
	if o == nil || isNil(o.Reason) {
    return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *AddLeaveLockdownModeRecurringTaskRequest) HasReason() bool {
	if o != nil && !isNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *AddLeaveLockdownModeRecurringTaskRequest) SetReason(v string) {
	o.Reason = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddLeaveLockdownModeRecurringTaskRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLeaveLockdownModeRecurringTaskRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddLeaveLockdownModeRecurringTaskRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddLeaveLockdownModeRecurringTaskRequest) SetDescription(v string) {
	o.Description = &v
}

// GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field value if set, zero value otherwise.
func (o *AddLeaveLockdownModeRecurringTaskRequest) GetCancelOnTaskDependencyFailure() bool {
	if o == nil || isNil(o.CancelOnTaskDependencyFailure) {
		var ret bool
		return ret
	}
	return *o.CancelOnTaskDependencyFailure
}

// GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLeaveLockdownModeRecurringTaskRequest) GetCancelOnTaskDependencyFailureOk() (*bool, bool) {
	if o == nil || isNil(o.CancelOnTaskDependencyFailure) {
    return nil, false
	}
	return o.CancelOnTaskDependencyFailure, true
}

// HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.
func (o *AddLeaveLockdownModeRecurringTaskRequest) HasCancelOnTaskDependencyFailure() bool {
	if o != nil && !isNil(o.CancelOnTaskDependencyFailure) {
		return true
	}

	return false
}

// SetCancelOnTaskDependencyFailure gets a reference to the given bool and assigns it to the CancelOnTaskDependencyFailure field.
func (o *AddLeaveLockdownModeRecurringTaskRequest) SetCancelOnTaskDependencyFailure(v bool) {
	o.CancelOnTaskDependencyFailure = &v
}

// GetEmailOnStart returns the EmailOnStart field value if set, zero value otherwise.
func (o *AddLeaveLockdownModeRecurringTaskRequest) GetEmailOnStart() []string {
	if o == nil || isNil(o.EmailOnStart) {
		var ret []string
		return ret
	}
	return o.EmailOnStart
}

// GetEmailOnStartOk returns a tuple with the EmailOnStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLeaveLockdownModeRecurringTaskRequest) GetEmailOnStartOk() ([]string, bool) {
	if o == nil || isNil(o.EmailOnStart) {
    return nil, false
	}
	return o.EmailOnStart, true
}

// HasEmailOnStart returns a boolean if a field has been set.
func (o *AddLeaveLockdownModeRecurringTaskRequest) HasEmailOnStart() bool {
	if o != nil && !isNil(o.EmailOnStart) {
		return true
	}

	return false
}

// SetEmailOnStart gets a reference to the given []string and assigns it to the EmailOnStart field.
func (o *AddLeaveLockdownModeRecurringTaskRequest) SetEmailOnStart(v []string) {
	o.EmailOnStart = v
}

// GetEmailOnSuccess returns the EmailOnSuccess field value if set, zero value otherwise.
func (o *AddLeaveLockdownModeRecurringTaskRequest) GetEmailOnSuccess() []string {
	if o == nil || isNil(o.EmailOnSuccess) {
		var ret []string
		return ret
	}
	return o.EmailOnSuccess
}

// GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLeaveLockdownModeRecurringTaskRequest) GetEmailOnSuccessOk() ([]string, bool) {
	if o == nil || isNil(o.EmailOnSuccess) {
    return nil, false
	}
	return o.EmailOnSuccess, true
}

// HasEmailOnSuccess returns a boolean if a field has been set.
func (o *AddLeaveLockdownModeRecurringTaskRequest) HasEmailOnSuccess() bool {
	if o != nil && !isNil(o.EmailOnSuccess) {
		return true
	}

	return false
}

// SetEmailOnSuccess gets a reference to the given []string and assigns it to the EmailOnSuccess field.
func (o *AddLeaveLockdownModeRecurringTaskRequest) SetEmailOnSuccess(v []string) {
	o.EmailOnSuccess = v
}

// GetEmailOnFailure returns the EmailOnFailure field value if set, zero value otherwise.
func (o *AddLeaveLockdownModeRecurringTaskRequest) GetEmailOnFailure() []string {
	if o == nil || isNil(o.EmailOnFailure) {
		var ret []string
		return ret
	}
	return o.EmailOnFailure
}

// GetEmailOnFailureOk returns a tuple with the EmailOnFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLeaveLockdownModeRecurringTaskRequest) GetEmailOnFailureOk() ([]string, bool) {
	if o == nil || isNil(o.EmailOnFailure) {
    return nil, false
	}
	return o.EmailOnFailure, true
}

// HasEmailOnFailure returns a boolean if a field has been set.
func (o *AddLeaveLockdownModeRecurringTaskRequest) HasEmailOnFailure() bool {
	if o != nil && !isNil(o.EmailOnFailure) {
		return true
	}

	return false
}

// SetEmailOnFailure gets a reference to the given []string and assigns it to the EmailOnFailure field.
func (o *AddLeaveLockdownModeRecurringTaskRequest) SetEmailOnFailure(v []string) {
	o.EmailOnFailure = v
}

// GetAlertOnStart returns the AlertOnStart field value if set, zero value otherwise.
func (o *AddLeaveLockdownModeRecurringTaskRequest) GetAlertOnStart() bool {
	if o == nil || isNil(o.AlertOnStart) {
		var ret bool
		return ret
	}
	return *o.AlertOnStart
}

// GetAlertOnStartOk returns a tuple with the AlertOnStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLeaveLockdownModeRecurringTaskRequest) GetAlertOnStartOk() (*bool, bool) {
	if o == nil || isNil(o.AlertOnStart) {
    return nil, false
	}
	return o.AlertOnStart, true
}

// HasAlertOnStart returns a boolean if a field has been set.
func (o *AddLeaveLockdownModeRecurringTaskRequest) HasAlertOnStart() bool {
	if o != nil && !isNil(o.AlertOnStart) {
		return true
	}

	return false
}

// SetAlertOnStart gets a reference to the given bool and assigns it to the AlertOnStart field.
func (o *AddLeaveLockdownModeRecurringTaskRequest) SetAlertOnStart(v bool) {
	o.AlertOnStart = &v
}

// GetAlertOnSuccess returns the AlertOnSuccess field value if set, zero value otherwise.
func (o *AddLeaveLockdownModeRecurringTaskRequest) GetAlertOnSuccess() bool {
	if o == nil || isNil(o.AlertOnSuccess) {
		var ret bool
		return ret
	}
	return *o.AlertOnSuccess
}

// GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLeaveLockdownModeRecurringTaskRequest) GetAlertOnSuccessOk() (*bool, bool) {
	if o == nil || isNil(o.AlertOnSuccess) {
    return nil, false
	}
	return o.AlertOnSuccess, true
}

// HasAlertOnSuccess returns a boolean if a field has been set.
func (o *AddLeaveLockdownModeRecurringTaskRequest) HasAlertOnSuccess() bool {
	if o != nil && !isNil(o.AlertOnSuccess) {
		return true
	}

	return false
}

// SetAlertOnSuccess gets a reference to the given bool and assigns it to the AlertOnSuccess field.
func (o *AddLeaveLockdownModeRecurringTaskRequest) SetAlertOnSuccess(v bool) {
	o.AlertOnSuccess = &v
}

// GetAlertOnFailure returns the AlertOnFailure field value if set, zero value otherwise.
func (o *AddLeaveLockdownModeRecurringTaskRequest) GetAlertOnFailure() bool {
	if o == nil || isNil(o.AlertOnFailure) {
		var ret bool
		return ret
	}
	return *o.AlertOnFailure
}

// GetAlertOnFailureOk returns a tuple with the AlertOnFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLeaveLockdownModeRecurringTaskRequest) GetAlertOnFailureOk() (*bool, bool) {
	if o == nil || isNil(o.AlertOnFailure) {
    return nil, false
	}
	return o.AlertOnFailure, true
}

// HasAlertOnFailure returns a boolean if a field has been set.
func (o *AddLeaveLockdownModeRecurringTaskRequest) HasAlertOnFailure() bool {
	if o != nil && !isNil(o.AlertOnFailure) {
		return true
	}

	return false
}

// SetAlertOnFailure gets a reference to the given bool and assigns it to the AlertOnFailure field.
func (o *AddLeaveLockdownModeRecurringTaskRequest) SetAlertOnFailure(v bool) {
	o.AlertOnFailure = &v
}

func (o AddLeaveLockdownModeRecurringTaskRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["taskName"] = o.TaskName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.CancelOnTaskDependencyFailure) {
		toSerialize["cancelOnTaskDependencyFailure"] = o.CancelOnTaskDependencyFailure
	}
	if !isNil(o.EmailOnStart) {
		toSerialize["emailOnStart"] = o.EmailOnStart
	}
	if !isNil(o.EmailOnSuccess) {
		toSerialize["emailOnSuccess"] = o.EmailOnSuccess
	}
	if !isNil(o.EmailOnFailure) {
		toSerialize["emailOnFailure"] = o.EmailOnFailure
	}
	if !isNil(o.AlertOnStart) {
		toSerialize["alertOnStart"] = o.AlertOnStart
	}
	if !isNil(o.AlertOnSuccess) {
		toSerialize["alertOnSuccess"] = o.AlertOnSuccess
	}
	if !isNil(o.AlertOnFailure) {
		toSerialize["alertOnFailure"] = o.AlertOnFailure
	}
	return json.Marshal(toSerialize)
}

type NullableAddLeaveLockdownModeRecurringTaskRequest struct {
	value *AddLeaveLockdownModeRecurringTaskRequest
	isSet bool
}

func (v NullableAddLeaveLockdownModeRecurringTaskRequest) Get() *AddLeaveLockdownModeRecurringTaskRequest {
	return v.value
}

func (v *NullableAddLeaveLockdownModeRecurringTaskRequest) Set(val *AddLeaveLockdownModeRecurringTaskRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddLeaveLockdownModeRecurringTaskRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddLeaveLockdownModeRecurringTaskRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddLeaveLockdownModeRecurringTaskRequest(val *AddLeaveLockdownModeRecurringTaskRequest) *NullableAddLeaveLockdownModeRecurringTaskRequest {
	return &NullableAddLeaveLockdownModeRecurringTaskRequest{value: val, isSet: true}
}

func (v NullableAddLeaveLockdownModeRecurringTaskRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddLeaveLockdownModeRecurringTaskRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


