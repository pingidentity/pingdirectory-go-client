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

// GenerateServerProfileRecurringTaskResponse struct for GenerateServerProfileRecurringTaskResponse
type GenerateServerProfileRecurringTaskResponse struct {
	// Name of the Recurring Task
	Id string `json:"id"`
	Schemas []EnumgenerateServerProfileRecurringTaskSchemaUrn `json:"schemas"`
	// The directory in which the generated server profiles will be placed. The files will be named with the pattern \"server-profile-{timestamp}.zip\", where \"{timestamp}\" represents the time that the profile was generated.
	ProfileDirectory string `json:"profileDirectory"`
	IncludePath []string `json:"includePath,omitempty"`
	// The minimum number of previous server profile zip files that should be preserved after a new profile is generated.
	RetainPreviousProfileCount *int32 `json:"retainPreviousProfileCount,omitempty"`
	// The minimum age of previous server profile zip files that should be preserved after a new profile is generated.
	RetainPreviousProfileAge *string `json:"retainPreviousProfileAge,omitempty"`
	// A description for this Recurring Task
	Description *string `json:"description,omitempty"`
	// Indicates whether an instance of this Recurring Task should be canceled if the task immediately before it in the recurring task chain fails to complete successfully (including if it is canceled by an administrator before it starts or while it is running).
	CancelOnTaskDependencyFailure *bool `json:"cancelOnTaskDependencyFailure,omitempty"`
	EmailOnStart []string `json:"emailOnStart,omitempty"`
	EmailOnSuccess []string `json:"emailOnSuccess,omitempty"`
	EmailOnFailure []string `json:"emailOnFailure,omitempty"`
	// Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task starts running.
	AlertOnStart *bool `json:"alertOnStart,omitempty"`
	// Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task completes successfully.
	AlertOnSuccess *bool `json:"alertOnSuccess,omitempty"`
	// Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task fails to complete successfully.
	AlertOnFailure *bool `json:"alertOnFailure,omitempty"`
}

// NewGenerateServerProfileRecurringTaskResponse instantiates a new GenerateServerProfileRecurringTaskResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateServerProfileRecurringTaskResponse(id string, schemas []EnumgenerateServerProfileRecurringTaskSchemaUrn, profileDirectory string) *GenerateServerProfileRecurringTaskResponse {
	this := GenerateServerProfileRecurringTaskResponse{}
	this.Id = id
	this.Schemas = schemas
	this.ProfileDirectory = profileDirectory
	return &this
}

// NewGenerateServerProfileRecurringTaskResponseWithDefaults instantiates a new GenerateServerProfileRecurringTaskResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateServerProfileRecurringTaskResponseWithDefaults() *GenerateServerProfileRecurringTaskResponse {
	this := GenerateServerProfileRecurringTaskResponse{}
	return &this
}

// GetId returns the Id field value
func (o *GenerateServerProfileRecurringTaskResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GenerateServerProfileRecurringTaskResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GenerateServerProfileRecurringTaskResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *GenerateServerProfileRecurringTaskResponse) GetSchemas() []EnumgenerateServerProfileRecurringTaskSchemaUrn {
	if o == nil {
		var ret []EnumgenerateServerProfileRecurringTaskSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *GenerateServerProfileRecurringTaskResponse) GetSchemasOk() ([]EnumgenerateServerProfileRecurringTaskSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *GenerateServerProfileRecurringTaskResponse) SetSchemas(v []EnumgenerateServerProfileRecurringTaskSchemaUrn) {
	o.Schemas = v
}

// GetProfileDirectory returns the ProfileDirectory field value
func (o *GenerateServerProfileRecurringTaskResponse) GetProfileDirectory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfileDirectory
}

// GetProfileDirectoryOk returns a tuple with the ProfileDirectory field value
// and a boolean to check if the value has been set.
func (o *GenerateServerProfileRecurringTaskResponse) GetProfileDirectoryOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ProfileDirectory, true
}

// SetProfileDirectory sets field value
func (o *GenerateServerProfileRecurringTaskResponse) SetProfileDirectory(v string) {
	o.ProfileDirectory = v
}

// GetIncludePath returns the IncludePath field value if set, zero value otherwise.
func (o *GenerateServerProfileRecurringTaskResponse) GetIncludePath() []string {
	if o == nil || isNil(o.IncludePath) {
		var ret []string
		return ret
	}
	return o.IncludePath
}

// GetIncludePathOk returns a tuple with the IncludePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateServerProfileRecurringTaskResponse) GetIncludePathOk() ([]string, bool) {
	if o == nil || isNil(o.IncludePath) {
    return nil, false
	}
	return o.IncludePath, true
}

// HasIncludePath returns a boolean if a field has been set.
func (o *GenerateServerProfileRecurringTaskResponse) HasIncludePath() bool {
	if o != nil && !isNil(o.IncludePath) {
		return true
	}

	return false
}

// SetIncludePath gets a reference to the given []string and assigns it to the IncludePath field.
func (o *GenerateServerProfileRecurringTaskResponse) SetIncludePath(v []string) {
	o.IncludePath = v
}

// GetRetainPreviousProfileCount returns the RetainPreviousProfileCount field value if set, zero value otherwise.
func (o *GenerateServerProfileRecurringTaskResponse) GetRetainPreviousProfileCount() int32 {
	if o == nil || isNil(o.RetainPreviousProfileCount) {
		var ret int32
		return ret
	}
	return *o.RetainPreviousProfileCount
}

// GetRetainPreviousProfileCountOk returns a tuple with the RetainPreviousProfileCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateServerProfileRecurringTaskResponse) GetRetainPreviousProfileCountOk() (*int32, bool) {
	if o == nil || isNil(o.RetainPreviousProfileCount) {
    return nil, false
	}
	return o.RetainPreviousProfileCount, true
}

// HasRetainPreviousProfileCount returns a boolean if a field has been set.
func (o *GenerateServerProfileRecurringTaskResponse) HasRetainPreviousProfileCount() bool {
	if o != nil && !isNil(o.RetainPreviousProfileCount) {
		return true
	}

	return false
}

// SetRetainPreviousProfileCount gets a reference to the given int32 and assigns it to the RetainPreviousProfileCount field.
func (o *GenerateServerProfileRecurringTaskResponse) SetRetainPreviousProfileCount(v int32) {
	o.RetainPreviousProfileCount = &v
}

// GetRetainPreviousProfileAge returns the RetainPreviousProfileAge field value if set, zero value otherwise.
func (o *GenerateServerProfileRecurringTaskResponse) GetRetainPreviousProfileAge() string {
	if o == nil || isNil(o.RetainPreviousProfileAge) {
		var ret string
		return ret
	}
	return *o.RetainPreviousProfileAge
}

// GetRetainPreviousProfileAgeOk returns a tuple with the RetainPreviousProfileAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateServerProfileRecurringTaskResponse) GetRetainPreviousProfileAgeOk() (*string, bool) {
	if o == nil || isNil(o.RetainPreviousProfileAge) {
    return nil, false
	}
	return o.RetainPreviousProfileAge, true
}

// HasRetainPreviousProfileAge returns a boolean if a field has been set.
func (o *GenerateServerProfileRecurringTaskResponse) HasRetainPreviousProfileAge() bool {
	if o != nil && !isNil(o.RetainPreviousProfileAge) {
		return true
	}

	return false
}

// SetRetainPreviousProfileAge gets a reference to the given string and assigns it to the RetainPreviousProfileAge field.
func (o *GenerateServerProfileRecurringTaskResponse) SetRetainPreviousProfileAge(v string) {
	o.RetainPreviousProfileAge = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GenerateServerProfileRecurringTaskResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateServerProfileRecurringTaskResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GenerateServerProfileRecurringTaskResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GenerateServerProfileRecurringTaskResponse) SetDescription(v string) {
	o.Description = &v
}

// GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field value if set, zero value otherwise.
func (o *GenerateServerProfileRecurringTaskResponse) GetCancelOnTaskDependencyFailure() bool {
	if o == nil || isNil(o.CancelOnTaskDependencyFailure) {
		var ret bool
		return ret
	}
	return *o.CancelOnTaskDependencyFailure
}

// GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateServerProfileRecurringTaskResponse) GetCancelOnTaskDependencyFailureOk() (*bool, bool) {
	if o == nil || isNil(o.CancelOnTaskDependencyFailure) {
    return nil, false
	}
	return o.CancelOnTaskDependencyFailure, true
}

// HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.
func (o *GenerateServerProfileRecurringTaskResponse) HasCancelOnTaskDependencyFailure() bool {
	if o != nil && !isNil(o.CancelOnTaskDependencyFailure) {
		return true
	}

	return false
}

// SetCancelOnTaskDependencyFailure gets a reference to the given bool and assigns it to the CancelOnTaskDependencyFailure field.
func (o *GenerateServerProfileRecurringTaskResponse) SetCancelOnTaskDependencyFailure(v bool) {
	o.CancelOnTaskDependencyFailure = &v
}

// GetEmailOnStart returns the EmailOnStart field value if set, zero value otherwise.
func (o *GenerateServerProfileRecurringTaskResponse) GetEmailOnStart() []string {
	if o == nil || isNil(o.EmailOnStart) {
		var ret []string
		return ret
	}
	return o.EmailOnStart
}

// GetEmailOnStartOk returns a tuple with the EmailOnStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateServerProfileRecurringTaskResponse) GetEmailOnStartOk() ([]string, bool) {
	if o == nil || isNil(o.EmailOnStart) {
    return nil, false
	}
	return o.EmailOnStart, true
}

// HasEmailOnStart returns a boolean if a field has been set.
func (o *GenerateServerProfileRecurringTaskResponse) HasEmailOnStart() bool {
	if o != nil && !isNil(o.EmailOnStart) {
		return true
	}

	return false
}

// SetEmailOnStart gets a reference to the given []string and assigns it to the EmailOnStart field.
func (o *GenerateServerProfileRecurringTaskResponse) SetEmailOnStart(v []string) {
	o.EmailOnStart = v
}

// GetEmailOnSuccess returns the EmailOnSuccess field value if set, zero value otherwise.
func (o *GenerateServerProfileRecurringTaskResponse) GetEmailOnSuccess() []string {
	if o == nil || isNil(o.EmailOnSuccess) {
		var ret []string
		return ret
	}
	return o.EmailOnSuccess
}

// GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateServerProfileRecurringTaskResponse) GetEmailOnSuccessOk() ([]string, bool) {
	if o == nil || isNil(o.EmailOnSuccess) {
    return nil, false
	}
	return o.EmailOnSuccess, true
}

// HasEmailOnSuccess returns a boolean if a field has been set.
func (o *GenerateServerProfileRecurringTaskResponse) HasEmailOnSuccess() bool {
	if o != nil && !isNil(o.EmailOnSuccess) {
		return true
	}

	return false
}

// SetEmailOnSuccess gets a reference to the given []string and assigns it to the EmailOnSuccess field.
func (o *GenerateServerProfileRecurringTaskResponse) SetEmailOnSuccess(v []string) {
	o.EmailOnSuccess = v
}

// GetEmailOnFailure returns the EmailOnFailure field value if set, zero value otherwise.
func (o *GenerateServerProfileRecurringTaskResponse) GetEmailOnFailure() []string {
	if o == nil || isNil(o.EmailOnFailure) {
		var ret []string
		return ret
	}
	return o.EmailOnFailure
}

// GetEmailOnFailureOk returns a tuple with the EmailOnFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateServerProfileRecurringTaskResponse) GetEmailOnFailureOk() ([]string, bool) {
	if o == nil || isNil(o.EmailOnFailure) {
    return nil, false
	}
	return o.EmailOnFailure, true
}

// HasEmailOnFailure returns a boolean if a field has been set.
func (o *GenerateServerProfileRecurringTaskResponse) HasEmailOnFailure() bool {
	if o != nil && !isNil(o.EmailOnFailure) {
		return true
	}

	return false
}

// SetEmailOnFailure gets a reference to the given []string and assigns it to the EmailOnFailure field.
func (o *GenerateServerProfileRecurringTaskResponse) SetEmailOnFailure(v []string) {
	o.EmailOnFailure = v
}

// GetAlertOnStart returns the AlertOnStart field value if set, zero value otherwise.
func (o *GenerateServerProfileRecurringTaskResponse) GetAlertOnStart() bool {
	if o == nil || isNil(o.AlertOnStart) {
		var ret bool
		return ret
	}
	return *o.AlertOnStart
}

// GetAlertOnStartOk returns a tuple with the AlertOnStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateServerProfileRecurringTaskResponse) GetAlertOnStartOk() (*bool, bool) {
	if o == nil || isNil(o.AlertOnStart) {
    return nil, false
	}
	return o.AlertOnStart, true
}

// HasAlertOnStart returns a boolean if a field has been set.
func (o *GenerateServerProfileRecurringTaskResponse) HasAlertOnStart() bool {
	if o != nil && !isNil(o.AlertOnStart) {
		return true
	}

	return false
}

// SetAlertOnStart gets a reference to the given bool and assigns it to the AlertOnStart field.
func (o *GenerateServerProfileRecurringTaskResponse) SetAlertOnStart(v bool) {
	o.AlertOnStart = &v
}

// GetAlertOnSuccess returns the AlertOnSuccess field value if set, zero value otherwise.
func (o *GenerateServerProfileRecurringTaskResponse) GetAlertOnSuccess() bool {
	if o == nil || isNil(o.AlertOnSuccess) {
		var ret bool
		return ret
	}
	return *o.AlertOnSuccess
}

// GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateServerProfileRecurringTaskResponse) GetAlertOnSuccessOk() (*bool, bool) {
	if o == nil || isNil(o.AlertOnSuccess) {
    return nil, false
	}
	return o.AlertOnSuccess, true
}

// HasAlertOnSuccess returns a boolean if a field has been set.
func (o *GenerateServerProfileRecurringTaskResponse) HasAlertOnSuccess() bool {
	if o != nil && !isNil(o.AlertOnSuccess) {
		return true
	}

	return false
}

// SetAlertOnSuccess gets a reference to the given bool and assigns it to the AlertOnSuccess field.
func (o *GenerateServerProfileRecurringTaskResponse) SetAlertOnSuccess(v bool) {
	o.AlertOnSuccess = &v
}

// GetAlertOnFailure returns the AlertOnFailure field value if set, zero value otherwise.
func (o *GenerateServerProfileRecurringTaskResponse) GetAlertOnFailure() bool {
	if o == nil || isNil(o.AlertOnFailure) {
		var ret bool
		return ret
	}
	return *o.AlertOnFailure
}

// GetAlertOnFailureOk returns a tuple with the AlertOnFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateServerProfileRecurringTaskResponse) GetAlertOnFailureOk() (*bool, bool) {
	if o == nil || isNil(o.AlertOnFailure) {
    return nil, false
	}
	return o.AlertOnFailure, true
}

// HasAlertOnFailure returns a boolean if a field has been set.
func (o *GenerateServerProfileRecurringTaskResponse) HasAlertOnFailure() bool {
	if o != nil && !isNil(o.AlertOnFailure) {
		return true
	}

	return false
}

// SetAlertOnFailure gets a reference to the given bool and assigns it to the AlertOnFailure field.
func (o *GenerateServerProfileRecurringTaskResponse) SetAlertOnFailure(v bool) {
	o.AlertOnFailure = &v
}

func (o GenerateServerProfileRecurringTaskResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["profileDirectory"] = o.ProfileDirectory
	}
	if !isNil(o.IncludePath) {
		toSerialize["includePath"] = o.IncludePath
	}
	if !isNil(o.RetainPreviousProfileCount) {
		toSerialize["retainPreviousProfileCount"] = o.RetainPreviousProfileCount
	}
	if !isNil(o.RetainPreviousProfileAge) {
		toSerialize["retainPreviousProfileAge"] = o.RetainPreviousProfileAge
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

type NullableGenerateServerProfileRecurringTaskResponse struct {
	value *GenerateServerProfileRecurringTaskResponse
	isSet bool
}

func (v NullableGenerateServerProfileRecurringTaskResponse) Get() *GenerateServerProfileRecurringTaskResponse {
	return v.value
}

func (v *NullableGenerateServerProfileRecurringTaskResponse) Set(val *GenerateServerProfileRecurringTaskResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateServerProfileRecurringTaskResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateServerProfileRecurringTaskResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateServerProfileRecurringTaskResponse(val *GenerateServerProfileRecurringTaskResponse) *NullableGenerateServerProfileRecurringTaskResponse {
	return &NullableGenerateServerProfileRecurringTaskResponse{value: val, isSet: true}
}

func (v NullableGenerateServerProfileRecurringTaskResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateServerProfileRecurringTaskResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


