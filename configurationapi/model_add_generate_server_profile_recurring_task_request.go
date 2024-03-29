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

// checks if the AddGenerateServerProfileRecurringTaskRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddGenerateServerProfileRecurringTaskRequest{}

// AddGenerateServerProfileRecurringTaskRequest struct for AddGenerateServerProfileRecurringTaskRequest
type AddGenerateServerProfileRecurringTaskRequest struct {
	Schemas []EnumgenerateServerProfileRecurringTaskSchemaUrn `json:"schemas"`
	// The directory in which the generated server profiles will be placed. The files will be named with the pattern \"server-profile-{timestamp}.zip\", where \"{timestamp}\" represents the time that the profile was generated.
	ProfileDirectory string `json:"profileDirectory"`
	// An optional set of additional paths to files within the instance root that should be included in the generated server profile. All paths must be within the instance root, and relative paths will be relative to the instance root.
	IncludePath []string `json:"includePath,omitempty"`
	// The minimum number of previous server profile zip files that should be preserved after a new profile is generated.
	RetainPreviousProfileCount *int64 `json:"retainPreviousProfileCount,omitempty"`
	// The minimum age of previous server profile zip files that should be preserved after a new profile is generated.
	RetainPreviousProfileAge *string `json:"retainPreviousProfileAge,omitempty"`
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
	// Name of the new Recurring Task
	TaskName string `json:"taskName"`
}

// NewAddGenerateServerProfileRecurringTaskRequest instantiates a new AddGenerateServerProfileRecurringTaskRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddGenerateServerProfileRecurringTaskRequest(schemas []EnumgenerateServerProfileRecurringTaskSchemaUrn, profileDirectory string, taskName string) *AddGenerateServerProfileRecurringTaskRequest {
	this := AddGenerateServerProfileRecurringTaskRequest{}
	this.Schemas = schemas
	this.ProfileDirectory = profileDirectory
	this.TaskName = taskName
	return &this
}

// NewAddGenerateServerProfileRecurringTaskRequestWithDefaults instantiates a new AddGenerateServerProfileRecurringTaskRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddGenerateServerProfileRecurringTaskRequestWithDefaults() *AddGenerateServerProfileRecurringTaskRequest {
	this := AddGenerateServerProfileRecurringTaskRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddGenerateServerProfileRecurringTaskRequest) GetSchemas() []EnumgenerateServerProfileRecurringTaskSchemaUrn {
	if o == nil {
		var ret []EnumgenerateServerProfileRecurringTaskSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddGenerateServerProfileRecurringTaskRequest) GetSchemasOk() ([]EnumgenerateServerProfileRecurringTaskSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddGenerateServerProfileRecurringTaskRequest) SetSchemas(v []EnumgenerateServerProfileRecurringTaskSchemaUrn) {
	o.Schemas = v
}

// GetProfileDirectory returns the ProfileDirectory field value
func (o *AddGenerateServerProfileRecurringTaskRequest) GetProfileDirectory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfileDirectory
}

// GetProfileDirectoryOk returns a tuple with the ProfileDirectory field value
// and a boolean to check if the value has been set.
func (o *AddGenerateServerProfileRecurringTaskRequest) GetProfileDirectoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfileDirectory, true
}

// SetProfileDirectory sets field value
func (o *AddGenerateServerProfileRecurringTaskRequest) SetProfileDirectory(v string) {
	o.ProfileDirectory = v
}

// GetIncludePath returns the IncludePath field value if set, zero value otherwise.
func (o *AddGenerateServerProfileRecurringTaskRequest) GetIncludePath() []string {
	if o == nil || IsNil(o.IncludePath) {
		var ret []string
		return ret
	}
	return o.IncludePath
}

// GetIncludePathOk returns a tuple with the IncludePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGenerateServerProfileRecurringTaskRequest) GetIncludePathOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludePath) {
		return nil, false
	}
	return o.IncludePath, true
}

// HasIncludePath returns a boolean if a field has been set.
func (o *AddGenerateServerProfileRecurringTaskRequest) HasIncludePath() bool {
	if o != nil && !IsNil(o.IncludePath) {
		return true
	}

	return false
}

// SetIncludePath gets a reference to the given []string and assigns it to the IncludePath field.
func (o *AddGenerateServerProfileRecurringTaskRequest) SetIncludePath(v []string) {
	o.IncludePath = v
}

// GetRetainPreviousProfileCount returns the RetainPreviousProfileCount field value if set, zero value otherwise.
func (o *AddGenerateServerProfileRecurringTaskRequest) GetRetainPreviousProfileCount() int64 {
	if o == nil || IsNil(o.RetainPreviousProfileCount) {
		var ret int64
		return ret
	}
	return *o.RetainPreviousProfileCount
}

// GetRetainPreviousProfileCountOk returns a tuple with the RetainPreviousProfileCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGenerateServerProfileRecurringTaskRequest) GetRetainPreviousProfileCountOk() (*int64, bool) {
	if o == nil || IsNil(o.RetainPreviousProfileCount) {
		return nil, false
	}
	return o.RetainPreviousProfileCount, true
}

// HasRetainPreviousProfileCount returns a boolean if a field has been set.
func (o *AddGenerateServerProfileRecurringTaskRequest) HasRetainPreviousProfileCount() bool {
	if o != nil && !IsNil(o.RetainPreviousProfileCount) {
		return true
	}

	return false
}

// SetRetainPreviousProfileCount gets a reference to the given int64 and assigns it to the RetainPreviousProfileCount field.
func (o *AddGenerateServerProfileRecurringTaskRequest) SetRetainPreviousProfileCount(v int64) {
	o.RetainPreviousProfileCount = &v
}

// GetRetainPreviousProfileAge returns the RetainPreviousProfileAge field value if set, zero value otherwise.
func (o *AddGenerateServerProfileRecurringTaskRequest) GetRetainPreviousProfileAge() string {
	if o == nil || IsNil(o.RetainPreviousProfileAge) {
		var ret string
		return ret
	}
	return *o.RetainPreviousProfileAge
}

// GetRetainPreviousProfileAgeOk returns a tuple with the RetainPreviousProfileAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGenerateServerProfileRecurringTaskRequest) GetRetainPreviousProfileAgeOk() (*string, bool) {
	if o == nil || IsNil(o.RetainPreviousProfileAge) {
		return nil, false
	}
	return o.RetainPreviousProfileAge, true
}

// HasRetainPreviousProfileAge returns a boolean if a field has been set.
func (o *AddGenerateServerProfileRecurringTaskRequest) HasRetainPreviousProfileAge() bool {
	if o != nil && !IsNil(o.RetainPreviousProfileAge) {
		return true
	}

	return false
}

// SetRetainPreviousProfileAge gets a reference to the given string and assigns it to the RetainPreviousProfileAge field.
func (o *AddGenerateServerProfileRecurringTaskRequest) SetRetainPreviousProfileAge(v string) {
	o.RetainPreviousProfileAge = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddGenerateServerProfileRecurringTaskRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGenerateServerProfileRecurringTaskRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddGenerateServerProfileRecurringTaskRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddGenerateServerProfileRecurringTaskRequest) SetDescription(v string) {
	o.Description = &v
}

// GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field value if set, zero value otherwise.
func (o *AddGenerateServerProfileRecurringTaskRequest) GetCancelOnTaskDependencyFailure() bool {
	if o == nil || IsNil(o.CancelOnTaskDependencyFailure) {
		var ret bool
		return ret
	}
	return *o.CancelOnTaskDependencyFailure
}

// GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGenerateServerProfileRecurringTaskRequest) GetCancelOnTaskDependencyFailureOk() (*bool, bool) {
	if o == nil || IsNil(o.CancelOnTaskDependencyFailure) {
		return nil, false
	}
	return o.CancelOnTaskDependencyFailure, true
}

// HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.
func (o *AddGenerateServerProfileRecurringTaskRequest) HasCancelOnTaskDependencyFailure() bool {
	if o != nil && !IsNil(o.CancelOnTaskDependencyFailure) {
		return true
	}

	return false
}

// SetCancelOnTaskDependencyFailure gets a reference to the given bool and assigns it to the CancelOnTaskDependencyFailure field.
func (o *AddGenerateServerProfileRecurringTaskRequest) SetCancelOnTaskDependencyFailure(v bool) {
	o.CancelOnTaskDependencyFailure = &v
}

// GetEmailOnStart returns the EmailOnStart field value if set, zero value otherwise.
func (o *AddGenerateServerProfileRecurringTaskRequest) GetEmailOnStart() []string {
	if o == nil || IsNil(o.EmailOnStart) {
		var ret []string
		return ret
	}
	return o.EmailOnStart
}

// GetEmailOnStartOk returns a tuple with the EmailOnStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGenerateServerProfileRecurringTaskRequest) GetEmailOnStartOk() ([]string, bool) {
	if o == nil || IsNil(o.EmailOnStart) {
		return nil, false
	}
	return o.EmailOnStart, true
}

// HasEmailOnStart returns a boolean if a field has been set.
func (o *AddGenerateServerProfileRecurringTaskRequest) HasEmailOnStart() bool {
	if o != nil && !IsNil(o.EmailOnStart) {
		return true
	}

	return false
}

// SetEmailOnStart gets a reference to the given []string and assigns it to the EmailOnStart field.
func (o *AddGenerateServerProfileRecurringTaskRequest) SetEmailOnStart(v []string) {
	o.EmailOnStart = v
}

// GetEmailOnSuccess returns the EmailOnSuccess field value if set, zero value otherwise.
func (o *AddGenerateServerProfileRecurringTaskRequest) GetEmailOnSuccess() []string {
	if o == nil || IsNil(o.EmailOnSuccess) {
		var ret []string
		return ret
	}
	return o.EmailOnSuccess
}

// GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGenerateServerProfileRecurringTaskRequest) GetEmailOnSuccessOk() ([]string, bool) {
	if o == nil || IsNil(o.EmailOnSuccess) {
		return nil, false
	}
	return o.EmailOnSuccess, true
}

// HasEmailOnSuccess returns a boolean if a field has been set.
func (o *AddGenerateServerProfileRecurringTaskRequest) HasEmailOnSuccess() bool {
	if o != nil && !IsNil(o.EmailOnSuccess) {
		return true
	}

	return false
}

// SetEmailOnSuccess gets a reference to the given []string and assigns it to the EmailOnSuccess field.
func (o *AddGenerateServerProfileRecurringTaskRequest) SetEmailOnSuccess(v []string) {
	o.EmailOnSuccess = v
}

// GetEmailOnFailure returns the EmailOnFailure field value if set, zero value otherwise.
func (o *AddGenerateServerProfileRecurringTaskRequest) GetEmailOnFailure() []string {
	if o == nil || IsNil(o.EmailOnFailure) {
		var ret []string
		return ret
	}
	return o.EmailOnFailure
}

// GetEmailOnFailureOk returns a tuple with the EmailOnFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGenerateServerProfileRecurringTaskRequest) GetEmailOnFailureOk() ([]string, bool) {
	if o == nil || IsNil(o.EmailOnFailure) {
		return nil, false
	}
	return o.EmailOnFailure, true
}

// HasEmailOnFailure returns a boolean if a field has been set.
func (o *AddGenerateServerProfileRecurringTaskRequest) HasEmailOnFailure() bool {
	if o != nil && !IsNil(o.EmailOnFailure) {
		return true
	}

	return false
}

// SetEmailOnFailure gets a reference to the given []string and assigns it to the EmailOnFailure field.
func (o *AddGenerateServerProfileRecurringTaskRequest) SetEmailOnFailure(v []string) {
	o.EmailOnFailure = v
}

// GetAlertOnStart returns the AlertOnStart field value if set, zero value otherwise.
func (o *AddGenerateServerProfileRecurringTaskRequest) GetAlertOnStart() bool {
	if o == nil || IsNil(o.AlertOnStart) {
		var ret bool
		return ret
	}
	return *o.AlertOnStart
}

// GetAlertOnStartOk returns a tuple with the AlertOnStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGenerateServerProfileRecurringTaskRequest) GetAlertOnStartOk() (*bool, bool) {
	if o == nil || IsNil(o.AlertOnStart) {
		return nil, false
	}
	return o.AlertOnStart, true
}

// HasAlertOnStart returns a boolean if a field has been set.
func (o *AddGenerateServerProfileRecurringTaskRequest) HasAlertOnStart() bool {
	if o != nil && !IsNil(o.AlertOnStart) {
		return true
	}

	return false
}

// SetAlertOnStart gets a reference to the given bool and assigns it to the AlertOnStart field.
func (o *AddGenerateServerProfileRecurringTaskRequest) SetAlertOnStart(v bool) {
	o.AlertOnStart = &v
}

// GetAlertOnSuccess returns the AlertOnSuccess field value if set, zero value otherwise.
func (o *AddGenerateServerProfileRecurringTaskRequest) GetAlertOnSuccess() bool {
	if o == nil || IsNil(o.AlertOnSuccess) {
		var ret bool
		return ret
	}
	return *o.AlertOnSuccess
}

// GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGenerateServerProfileRecurringTaskRequest) GetAlertOnSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.AlertOnSuccess) {
		return nil, false
	}
	return o.AlertOnSuccess, true
}

// HasAlertOnSuccess returns a boolean if a field has been set.
func (o *AddGenerateServerProfileRecurringTaskRequest) HasAlertOnSuccess() bool {
	if o != nil && !IsNil(o.AlertOnSuccess) {
		return true
	}

	return false
}

// SetAlertOnSuccess gets a reference to the given bool and assigns it to the AlertOnSuccess field.
func (o *AddGenerateServerProfileRecurringTaskRequest) SetAlertOnSuccess(v bool) {
	o.AlertOnSuccess = &v
}

// GetAlertOnFailure returns the AlertOnFailure field value if set, zero value otherwise.
func (o *AddGenerateServerProfileRecurringTaskRequest) GetAlertOnFailure() bool {
	if o == nil || IsNil(o.AlertOnFailure) {
		var ret bool
		return ret
	}
	return *o.AlertOnFailure
}

// GetAlertOnFailureOk returns a tuple with the AlertOnFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGenerateServerProfileRecurringTaskRequest) GetAlertOnFailureOk() (*bool, bool) {
	if o == nil || IsNil(o.AlertOnFailure) {
		return nil, false
	}
	return o.AlertOnFailure, true
}

// HasAlertOnFailure returns a boolean if a field has been set.
func (o *AddGenerateServerProfileRecurringTaskRequest) HasAlertOnFailure() bool {
	if o != nil && !IsNil(o.AlertOnFailure) {
		return true
	}

	return false
}

// SetAlertOnFailure gets a reference to the given bool and assigns it to the AlertOnFailure field.
func (o *AddGenerateServerProfileRecurringTaskRequest) SetAlertOnFailure(v bool) {
	o.AlertOnFailure = &v
}

// GetTaskName returns the TaskName field value
func (o *AddGenerateServerProfileRecurringTaskRequest) GetTaskName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaskName
}

// GetTaskNameOk returns a tuple with the TaskName field value
// and a boolean to check if the value has been set.
func (o *AddGenerateServerProfileRecurringTaskRequest) GetTaskNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskName, true
}

// SetTaskName sets field value
func (o *AddGenerateServerProfileRecurringTaskRequest) SetTaskName(v string) {
	o.TaskName = v
}

func (o AddGenerateServerProfileRecurringTaskRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddGenerateServerProfileRecurringTaskRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["profileDirectory"] = o.ProfileDirectory
	if !IsNil(o.IncludePath) {
		toSerialize["includePath"] = o.IncludePath
	}
	if !IsNil(o.RetainPreviousProfileCount) {
		toSerialize["retainPreviousProfileCount"] = o.RetainPreviousProfileCount
	}
	if !IsNil(o.RetainPreviousProfileAge) {
		toSerialize["retainPreviousProfileAge"] = o.RetainPreviousProfileAge
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
	toSerialize["taskName"] = o.TaskName
	return toSerialize, nil
}

type NullableAddGenerateServerProfileRecurringTaskRequest struct {
	value *AddGenerateServerProfileRecurringTaskRequest
	isSet bool
}

func (v NullableAddGenerateServerProfileRecurringTaskRequest) Get() *AddGenerateServerProfileRecurringTaskRequest {
	return v.value
}

func (v *NullableAddGenerateServerProfileRecurringTaskRequest) Set(val *AddGenerateServerProfileRecurringTaskRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddGenerateServerProfileRecurringTaskRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddGenerateServerProfileRecurringTaskRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddGenerateServerProfileRecurringTaskRequest(val *AddGenerateServerProfileRecurringTaskRequest) *NullableAddGenerateServerProfileRecurringTaskRequest {
	return &NullableAddGenerateServerProfileRecurringTaskRequest{value: val, isSet: true}
}

func (v NullableAddGenerateServerProfileRecurringTaskRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddGenerateServerProfileRecurringTaskRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
