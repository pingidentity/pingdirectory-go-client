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

// checks if the EnterLockdownModeRecurringTaskResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnterLockdownModeRecurringTaskResponse{}

// EnterLockdownModeRecurringTaskResponse struct for EnterLockdownModeRecurringTaskResponse
type EnterLockdownModeRecurringTaskResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Recurring Task
	Id      string                                        `json:"id"`
	Schemas []EnumenterLockdownModeRecurringTaskSchemaUrn `json:"schemas"`
	// The reason that the server is being placed in lockdown mode.
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

// NewEnterLockdownModeRecurringTaskResponse instantiates a new EnterLockdownModeRecurringTaskResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnterLockdownModeRecurringTaskResponse(id string, schemas []EnumenterLockdownModeRecurringTaskSchemaUrn) *EnterLockdownModeRecurringTaskResponse {
	this := EnterLockdownModeRecurringTaskResponse{}
	this.Id = id
	this.Schemas = schemas
	return &this
}

// NewEnterLockdownModeRecurringTaskResponseWithDefaults instantiates a new EnterLockdownModeRecurringTaskResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnterLockdownModeRecurringTaskResponseWithDefaults() *EnterLockdownModeRecurringTaskResponse {
	this := EnterLockdownModeRecurringTaskResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *EnterLockdownModeRecurringTaskResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterLockdownModeRecurringTaskResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *EnterLockdownModeRecurringTaskResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *EnterLockdownModeRecurringTaskResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *EnterLockdownModeRecurringTaskResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterLockdownModeRecurringTaskResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *EnterLockdownModeRecurringTaskResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *EnterLockdownModeRecurringTaskResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *EnterLockdownModeRecurringTaskResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EnterLockdownModeRecurringTaskResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EnterLockdownModeRecurringTaskResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *EnterLockdownModeRecurringTaskResponse) GetSchemas() []EnumenterLockdownModeRecurringTaskSchemaUrn {
	if o == nil {
		var ret []EnumenterLockdownModeRecurringTaskSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *EnterLockdownModeRecurringTaskResponse) GetSchemasOk() ([]EnumenterLockdownModeRecurringTaskSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *EnterLockdownModeRecurringTaskResponse) SetSchemas(v []EnumenterLockdownModeRecurringTaskSchemaUrn) {
	o.Schemas = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *EnterLockdownModeRecurringTaskResponse) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterLockdownModeRecurringTaskResponse) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *EnterLockdownModeRecurringTaskResponse) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *EnterLockdownModeRecurringTaskResponse) SetReason(v string) {
	o.Reason = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EnterLockdownModeRecurringTaskResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterLockdownModeRecurringTaskResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EnterLockdownModeRecurringTaskResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EnterLockdownModeRecurringTaskResponse) SetDescription(v string) {
	o.Description = &v
}

// GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field value if set, zero value otherwise.
func (o *EnterLockdownModeRecurringTaskResponse) GetCancelOnTaskDependencyFailure() bool {
	if o == nil || IsNil(o.CancelOnTaskDependencyFailure) {
		var ret bool
		return ret
	}
	return *o.CancelOnTaskDependencyFailure
}

// GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterLockdownModeRecurringTaskResponse) GetCancelOnTaskDependencyFailureOk() (*bool, bool) {
	if o == nil || IsNil(o.CancelOnTaskDependencyFailure) {
		return nil, false
	}
	return o.CancelOnTaskDependencyFailure, true
}

// HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.
func (o *EnterLockdownModeRecurringTaskResponse) HasCancelOnTaskDependencyFailure() bool {
	if o != nil && !IsNil(o.CancelOnTaskDependencyFailure) {
		return true
	}

	return false
}

// SetCancelOnTaskDependencyFailure gets a reference to the given bool and assigns it to the CancelOnTaskDependencyFailure field.
func (o *EnterLockdownModeRecurringTaskResponse) SetCancelOnTaskDependencyFailure(v bool) {
	o.CancelOnTaskDependencyFailure = &v
}

// GetEmailOnStart returns the EmailOnStart field value if set, zero value otherwise.
func (o *EnterLockdownModeRecurringTaskResponse) GetEmailOnStart() []string {
	if o == nil || IsNil(o.EmailOnStart) {
		var ret []string
		return ret
	}
	return o.EmailOnStart
}

// GetEmailOnStartOk returns a tuple with the EmailOnStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterLockdownModeRecurringTaskResponse) GetEmailOnStartOk() ([]string, bool) {
	if o == nil || IsNil(o.EmailOnStart) {
		return nil, false
	}
	return o.EmailOnStart, true
}

// HasEmailOnStart returns a boolean if a field has been set.
func (o *EnterLockdownModeRecurringTaskResponse) HasEmailOnStart() bool {
	if o != nil && !IsNil(o.EmailOnStart) {
		return true
	}

	return false
}

// SetEmailOnStart gets a reference to the given []string and assigns it to the EmailOnStart field.
func (o *EnterLockdownModeRecurringTaskResponse) SetEmailOnStart(v []string) {
	o.EmailOnStart = v
}

// GetEmailOnSuccess returns the EmailOnSuccess field value if set, zero value otherwise.
func (o *EnterLockdownModeRecurringTaskResponse) GetEmailOnSuccess() []string {
	if o == nil || IsNil(o.EmailOnSuccess) {
		var ret []string
		return ret
	}
	return o.EmailOnSuccess
}

// GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterLockdownModeRecurringTaskResponse) GetEmailOnSuccessOk() ([]string, bool) {
	if o == nil || IsNil(o.EmailOnSuccess) {
		return nil, false
	}
	return o.EmailOnSuccess, true
}

// HasEmailOnSuccess returns a boolean if a field has been set.
func (o *EnterLockdownModeRecurringTaskResponse) HasEmailOnSuccess() bool {
	if o != nil && !IsNil(o.EmailOnSuccess) {
		return true
	}

	return false
}

// SetEmailOnSuccess gets a reference to the given []string and assigns it to the EmailOnSuccess field.
func (o *EnterLockdownModeRecurringTaskResponse) SetEmailOnSuccess(v []string) {
	o.EmailOnSuccess = v
}

// GetEmailOnFailure returns the EmailOnFailure field value if set, zero value otherwise.
func (o *EnterLockdownModeRecurringTaskResponse) GetEmailOnFailure() []string {
	if o == nil || IsNil(o.EmailOnFailure) {
		var ret []string
		return ret
	}
	return o.EmailOnFailure
}

// GetEmailOnFailureOk returns a tuple with the EmailOnFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterLockdownModeRecurringTaskResponse) GetEmailOnFailureOk() ([]string, bool) {
	if o == nil || IsNil(o.EmailOnFailure) {
		return nil, false
	}
	return o.EmailOnFailure, true
}

// HasEmailOnFailure returns a boolean if a field has been set.
func (o *EnterLockdownModeRecurringTaskResponse) HasEmailOnFailure() bool {
	if o != nil && !IsNil(o.EmailOnFailure) {
		return true
	}

	return false
}

// SetEmailOnFailure gets a reference to the given []string and assigns it to the EmailOnFailure field.
func (o *EnterLockdownModeRecurringTaskResponse) SetEmailOnFailure(v []string) {
	o.EmailOnFailure = v
}

// GetAlertOnStart returns the AlertOnStart field value if set, zero value otherwise.
func (o *EnterLockdownModeRecurringTaskResponse) GetAlertOnStart() bool {
	if o == nil || IsNil(o.AlertOnStart) {
		var ret bool
		return ret
	}
	return *o.AlertOnStart
}

// GetAlertOnStartOk returns a tuple with the AlertOnStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterLockdownModeRecurringTaskResponse) GetAlertOnStartOk() (*bool, bool) {
	if o == nil || IsNil(o.AlertOnStart) {
		return nil, false
	}
	return o.AlertOnStart, true
}

// HasAlertOnStart returns a boolean if a field has been set.
func (o *EnterLockdownModeRecurringTaskResponse) HasAlertOnStart() bool {
	if o != nil && !IsNil(o.AlertOnStart) {
		return true
	}

	return false
}

// SetAlertOnStart gets a reference to the given bool and assigns it to the AlertOnStart field.
func (o *EnterLockdownModeRecurringTaskResponse) SetAlertOnStart(v bool) {
	o.AlertOnStart = &v
}

// GetAlertOnSuccess returns the AlertOnSuccess field value if set, zero value otherwise.
func (o *EnterLockdownModeRecurringTaskResponse) GetAlertOnSuccess() bool {
	if o == nil || IsNil(o.AlertOnSuccess) {
		var ret bool
		return ret
	}
	return *o.AlertOnSuccess
}

// GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterLockdownModeRecurringTaskResponse) GetAlertOnSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.AlertOnSuccess) {
		return nil, false
	}
	return o.AlertOnSuccess, true
}

// HasAlertOnSuccess returns a boolean if a field has been set.
func (o *EnterLockdownModeRecurringTaskResponse) HasAlertOnSuccess() bool {
	if o != nil && !IsNil(o.AlertOnSuccess) {
		return true
	}

	return false
}

// SetAlertOnSuccess gets a reference to the given bool and assigns it to the AlertOnSuccess field.
func (o *EnterLockdownModeRecurringTaskResponse) SetAlertOnSuccess(v bool) {
	o.AlertOnSuccess = &v
}

// GetAlertOnFailure returns the AlertOnFailure field value if set, zero value otherwise.
func (o *EnterLockdownModeRecurringTaskResponse) GetAlertOnFailure() bool {
	if o == nil || IsNil(o.AlertOnFailure) {
		var ret bool
		return ret
	}
	return *o.AlertOnFailure
}

// GetAlertOnFailureOk returns a tuple with the AlertOnFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterLockdownModeRecurringTaskResponse) GetAlertOnFailureOk() (*bool, bool) {
	if o == nil || IsNil(o.AlertOnFailure) {
		return nil, false
	}
	return o.AlertOnFailure, true
}

// HasAlertOnFailure returns a boolean if a field has been set.
func (o *EnterLockdownModeRecurringTaskResponse) HasAlertOnFailure() bool {
	if o != nil && !IsNil(o.AlertOnFailure) {
		return true
	}

	return false
}

// SetAlertOnFailure gets a reference to the given bool and assigns it to the AlertOnFailure field.
func (o *EnterLockdownModeRecurringTaskResponse) SetAlertOnFailure(v bool) {
	o.AlertOnFailure = &v
}

func (o EnterLockdownModeRecurringTaskResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnterLockdownModeRecurringTaskResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CancelOnTaskDependencyFailure) {
		toSerialize["cancelOnTaskDependencyFailure"] = o.CancelOnTaskDependencyFailure
	}
	if !IsNil(o.EmailOnStart) {
		toSerialize["emailOnStart"] = o.EmailOnStart
	}
	if !IsNil(o.EmailOnSuccess) {
		toSerialize["emailOnSuccess"] = o.EmailOnSuccess
	}
	if !IsNil(o.EmailOnFailure) {
		toSerialize["emailOnFailure"] = o.EmailOnFailure
	}
	if !IsNil(o.AlertOnStart) {
		toSerialize["alertOnStart"] = o.AlertOnStart
	}
	if !IsNil(o.AlertOnSuccess) {
		toSerialize["alertOnSuccess"] = o.AlertOnSuccess
	}
	if !IsNil(o.AlertOnFailure) {
		toSerialize["alertOnFailure"] = o.AlertOnFailure
	}
	return toSerialize, nil
}

type NullableEnterLockdownModeRecurringTaskResponse struct {
	value *EnterLockdownModeRecurringTaskResponse
	isSet bool
}

func (v NullableEnterLockdownModeRecurringTaskResponse) Get() *EnterLockdownModeRecurringTaskResponse {
	return v.value
}

func (v *NullableEnterLockdownModeRecurringTaskResponse) Set(val *EnterLockdownModeRecurringTaskResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEnterLockdownModeRecurringTaskResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEnterLockdownModeRecurringTaskResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnterLockdownModeRecurringTaskResponse(val *EnterLockdownModeRecurringTaskResponse) *NullableEnterLockdownModeRecurringTaskResponse {
	return &NullableEnterLockdownModeRecurringTaskResponse{value: val, isSet: true}
}

func (v NullableEnterLockdownModeRecurringTaskResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnterLockdownModeRecurringTaskResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
