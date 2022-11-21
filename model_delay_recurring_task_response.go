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

// DelayRecurringTaskResponse struct for DelayRecurringTaskResponse
type DelayRecurringTaskResponse struct {
	// Name of the Recurring Task
	Id string `json:"id"`
	Schemas []EnumdelayRecurringTaskSchemaUrn `json:"schemas"`
	// The length of time to sleep before the task completes.
	SleepDuration *string `json:"sleepDuration,omitempty"`
	// Indicates that task should wait for up to the specified length of time for the work queue to report that all worker threads are idle and there are no pending operations. Note that this primarily monitors operations that use worker threads, which does not include internal operations (for example, those invoked by extensions), and may not include requests from non-LDAP clients (for example, HTTP-based clients).
	DurationToWaitForWorkQueueIdle *string `json:"durationToWaitForWorkQueueIdle,omitempty"`
	// An LDAP URL that provides the criteria for a search request that is expected to return at least one entry. The search will be performed internally, and only the base DN, scope, and filter from the URL will be used; any host, port, or requested attributes included in the URL will be ignored.
	LdapURLForSearchExpectedToReturnEntries []string `json:"ldapURLForSearchExpectedToReturnEntries,omitempty"`
	// The length of time the server should sleep between searches performed using the criteria from the ldap-url-for-search-expected-to-return-entries property.
	SearchInterval *string `json:"searchInterval,omitempty"`
	// The length of time that the server will wait for a response to each internal search performed using the criteria from the ldap-url-for-search-expected-to-return-entries property.
	SearchTimeLimit *string `json:"searchTimeLimit,omitempty"`
	// The maximum length of time that the server will continue to perform internal searches using the criteria from the ldap-url-for-search-expected-to-return-entries property.
	DurationToWaitForSearchToReturnEntries *string `json:"durationToWaitForSearchToReturnEntries,omitempty"`
	TaskReturnStateIfTimeoutIsEncountered *EnumrecurringTaskTaskReturnStateIfTimeoutIsEncounteredProp `json:"taskReturnStateIfTimeoutIsEncountered,omitempty"`
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
	Meta *MetaMeta `json:"meta,omitempty"`
}

// NewDelayRecurringTaskResponse instantiates a new DelayRecurringTaskResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDelayRecurringTaskResponse(id string, schemas []EnumdelayRecurringTaskSchemaUrn) *DelayRecurringTaskResponse {
	this := DelayRecurringTaskResponse{}
	this.Id = id
	this.Schemas = schemas
	return &this
}

// NewDelayRecurringTaskResponseWithDefaults instantiates a new DelayRecurringTaskResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDelayRecurringTaskResponseWithDefaults() *DelayRecurringTaskResponse {
	this := DelayRecurringTaskResponse{}
	return &this
}

// GetId returns the Id field value
func (o *DelayRecurringTaskResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DelayRecurringTaskResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DelayRecurringTaskResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *DelayRecurringTaskResponse) GetSchemas() []EnumdelayRecurringTaskSchemaUrn {
	if o == nil {
		var ret []EnumdelayRecurringTaskSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *DelayRecurringTaskResponse) GetSchemasOk() ([]EnumdelayRecurringTaskSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *DelayRecurringTaskResponse) SetSchemas(v []EnumdelayRecurringTaskSchemaUrn) {
	o.Schemas = v
}

// GetSleepDuration returns the SleepDuration field value if set, zero value otherwise.
func (o *DelayRecurringTaskResponse) GetSleepDuration() string {
	if o == nil || isNil(o.SleepDuration) {
		var ret string
		return ret
	}
	return *o.SleepDuration
}

// GetSleepDurationOk returns a tuple with the SleepDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayRecurringTaskResponse) GetSleepDurationOk() (*string, bool) {
	if o == nil || isNil(o.SleepDuration) {
    return nil, false
	}
	return o.SleepDuration, true
}

// HasSleepDuration returns a boolean if a field has been set.
func (o *DelayRecurringTaskResponse) HasSleepDuration() bool {
	if o != nil && !isNil(o.SleepDuration) {
		return true
	}

	return false
}

// SetSleepDuration gets a reference to the given string and assigns it to the SleepDuration field.
func (o *DelayRecurringTaskResponse) SetSleepDuration(v string) {
	o.SleepDuration = &v
}

// GetDurationToWaitForWorkQueueIdle returns the DurationToWaitForWorkQueueIdle field value if set, zero value otherwise.
func (o *DelayRecurringTaskResponse) GetDurationToWaitForWorkQueueIdle() string {
	if o == nil || isNil(o.DurationToWaitForWorkQueueIdle) {
		var ret string
		return ret
	}
	return *o.DurationToWaitForWorkQueueIdle
}

// GetDurationToWaitForWorkQueueIdleOk returns a tuple with the DurationToWaitForWorkQueueIdle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayRecurringTaskResponse) GetDurationToWaitForWorkQueueIdleOk() (*string, bool) {
	if o == nil || isNil(o.DurationToWaitForWorkQueueIdle) {
    return nil, false
	}
	return o.DurationToWaitForWorkQueueIdle, true
}

// HasDurationToWaitForWorkQueueIdle returns a boolean if a field has been set.
func (o *DelayRecurringTaskResponse) HasDurationToWaitForWorkQueueIdle() bool {
	if o != nil && !isNil(o.DurationToWaitForWorkQueueIdle) {
		return true
	}

	return false
}

// SetDurationToWaitForWorkQueueIdle gets a reference to the given string and assigns it to the DurationToWaitForWorkQueueIdle field.
func (o *DelayRecurringTaskResponse) SetDurationToWaitForWorkQueueIdle(v string) {
	o.DurationToWaitForWorkQueueIdle = &v
}

// GetLdapURLForSearchExpectedToReturnEntries returns the LdapURLForSearchExpectedToReturnEntries field value if set, zero value otherwise.
func (o *DelayRecurringTaskResponse) GetLdapURLForSearchExpectedToReturnEntries() []string {
	if o == nil || isNil(o.LdapURLForSearchExpectedToReturnEntries) {
		var ret []string
		return ret
	}
	return o.LdapURLForSearchExpectedToReturnEntries
}

// GetLdapURLForSearchExpectedToReturnEntriesOk returns a tuple with the LdapURLForSearchExpectedToReturnEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayRecurringTaskResponse) GetLdapURLForSearchExpectedToReturnEntriesOk() ([]string, bool) {
	if o == nil || isNil(o.LdapURLForSearchExpectedToReturnEntries) {
    return nil, false
	}
	return o.LdapURLForSearchExpectedToReturnEntries, true
}

// HasLdapURLForSearchExpectedToReturnEntries returns a boolean if a field has been set.
func (o *DelayRecurringTaskResponse) HasLdapURLForSearchExpectedToReturnEntries() bool {
	if o != nil && !isNil(o.LdapURLForSearchExpectedToReturnEntries) {
		return true
	}

	return false
}

// SetLdapURLForSearchExpectedToReturnEntries gets a reference to the given []string and assigns it to the LdapURLForSearchExpectedToReturnEntries field.
func (o *DelayRecurringTaskResponse) SetLdapURLForSearchExpectedToReturnEntries(v []string) {
	o.LdapURLForSearchExpectedToReturnEntries = v
}

// GetSearchInterval returns the SearchInterval field value if set, zero value otherwise.
func (o *DelayRecurringTaskResponse) GetSearchInterval() string {
	if o == nil || isNil(o.SearchInterval) {
		var ret string
		return ret
	}
	return *o.SearchInterval
}

// GetSearchIntervalOk returns a tuple with the SearchInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayRecurringTaskResponse) GetSearchIntervalOk() (*string, bool) {
	if o == nil || isNil(o.SearchInterval) {
    return nil, false
	}
	return o.SearchInterval, true
}

// HasSearchInterval returns a boolean if a field has been set.
func (o *DelayRecurringTaskResponse) HasSearchInterval() bool {
	if o != nil && !isNil(o.SearchInterval) {
		return true
	}

	return false
}

// SetSearchInterval gets a reference to the given string and assigns it to the SearchInterval field.
func (o *DelayRecurringTaskResponse) SetSearchInterval(v string) {
	o.SearchInterval = &v
}

// GetSearchTimeLimit returns the SearchTimeLimit field value if set, zero value otherwise.
func (o *DelayRecurringTaskResponse) GetSearchTimeLimit() string {
	if o == nil || isNil(o.SearchTimeLimit) {
		var ret string
		return ret
	}
	return *o.SearchTimeLimit
}

// GetSearchTimeLimitOk returns a tuple with the SearchTimeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayRecurringTaskResponse) GetSearchTimeLimitOk() (*string, bool) {
	if o == nil || isNil(o.SearchTimeLimit) {
    return nil, false
	}
	return o.SearchTimeLimit, true
}

// HasSearchTimeLimit returns a boolean if a field has been set.
func (o *DelayRecurringTaskResponse) HasSearchTimeLimit() bool {
	if o != nil && !isNil(o.SearchTimeLimit) {
		return true
	}

	return false
}

// SetSearchTimeLimit gets a reference to the given string and assigns it to the SearchTimeLimit field.
func (o *DelayRecurringTaskResponse) SetSearchTimeLimit(v string) {
	o.SearchTimeLimit = &v
}

// GetDurationToWaitForSearchToReturnEntries returns the DurationToWaitForSearchToReturnEntries field value if set, zero value otherwise.
func (o *DelayRecurringTaskResponse) GetDurationToWaitForSearchToReturnEntries() string {
	if o == nil || isNil(o.DurationToWaitForSearchToReturnEntries) {
		var ret string
		return ret
	}
	return *o.DurationToWaitForSearchToReturnEntries
}

// GetDurationToWaitForSearchToReturnEntriesOk returns a tuple with the DurationToWaitForSearchToReturnEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayRecurringTaskResponse) GetDurationToWaitForSearchToReturnEntriesOk() (*string, bool) {
	if o == nil || isNil(o.DurationToWaitForSearchToReturnEntries) {
    return nil, false
	}
	return o.DurationToWaitForSearchToReturnEntries, true
}

// HasDurationToWaitForSearchToReturnEntries returns a boolean if a field has been set.
func (o *DelayRecurringTaskResponse) HasDurationToWaitForSearchToReturnEntries() bool {
	if o != nil && !isNil(o.DurationToWaitForSearchToReturnEntries) {
		return true
	}

	return false
}

// SetDurationToWaitForSearchToReturnEntries gets a reference to the given string and assigns it to the DurationToWaitForSearchToReturnEntries field.
func (o *DelayRecurringTaskResponse) SetDurationToWaitForSearchToReturnEntries(v string) {
	o.DurationToWaitForSearchToReturnEntries = &v
}

// GetTaskReturnStateIfTimeoutIsEncountered returns the TaskReturnStateIfTimeoutIsEncountered field value if set, zero value otherwise.
func (o *DelayRecurringTaskResponse) GetTaskReturnStateIfTimeoutIsEncountered() EnumrecurringTaskTaskReturnStateIfTimeoutIsEncounteredProp {
	if o == nil || isNil(o.TaskReturnStateIfTimeoutIsEncountered) {
		var ret EnumrecurringTaskTaskReturnStateIfTimeoutIsEncounteredProp
		return ret
	}
	return *o.TaskReturnStateIfTimeoutIsEncountered
}

// GetTaskReturnStateIfTimeoutIsEncounteredOk returns a tuple with the TaskReturnStateIfTimeoutIsEncountered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayRecurringTaskResponse) GetTaskReturnStateIfTimeoutIsEncounteredOk() (*EnumrecurringTaskTaskReturnStateIfTimeoutIsEncounteredProp, bool) {
	if o == nil || isNil(o.TaskReturnStateIfTimeoutIsEncountered) {
    return nil, false
	}
	return o.TaskReturnStateIfTimeoutIsEncountered, true
}

// HasTaskReturnStateIfTimeoutIsEncountered returns a boolean if a field has been set.
func (o *DelayRecurringTaskResponse) HasTaskReturnStateIfTimeoutIsEncountered() bool {
	if o != nil && !isNil(o.TaskReturnStateIfTimeoutIsEncountered) {
		return true
	}

	return false
}

// SetTaskReturnStateIfTimeoutIsEncountered gets a reference to the given EnumrecurringTaskTaskReturnStateIfTimeoutIsEncounteredProp and assigns it to the TaskReturnStateIfTimeoutIsEncountered field.
func (o *DelayRecurringTaskResponse) SetTaskReturnStateIfTimeoutIsEncountered(v EnumrecurringTaskTaskReturnStateIfTimeoutIsEncounteredProp) {
	o.TaskReturnStateIfTimeoutIsEncountered = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DelayRecurringTaskResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayRecurringTaskResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DelayRecurringTaskResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DelayRecurringTaskResponse) SetDescription(v string) {
	o.Description = &v
}

// GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field value if set, zero value otherwise.
func (o *DelayRecurringTaskResponse) GetCancelOnTaskDependencyFailure() bool {
	if o == nil || isNil(o.CancelOnTaskDependencyFailure) {
		var ret bool
		return ret
	}
	return *o.CancelOnTaskDependencyFailure
}

// GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayRecurringTaskResponse) GetCancelOnTaskDependencyFailureOk() (*bool, bool) {
	if o == nil || isNil(o.CancelOnTaskDependencyFailure) {
    return nil, false
	}
	return o.CancelOnTaskDependencyFailure, true
}

// HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.
func (o *DelayRecurringTaskResponse) HasCancelOnTaskDependencyFailure() bool {
	if o != nil && !isNil(o.CancelOnTaskDependencyFailure) {
		return true
	}

	return false
}

// SetCancelOnTaskDependencyFailure gets a reference to the given bool and assigns it to the CancelOnTaskDependencyFailure field.
func (o *DelayRecurringTaskResponse) SetCancelOnTaskDependencyFailure(v bool) {
	o.CancelOnTaskDependencyFailure = &v
}

// GetEmailOnStart returns the EmailOnStart field value if set, zero value otherwise.
func (o *DelayRecurringTaskResponse) GetEmailOnStart() []string {
	if o == nil || isNil(o.EmailOnStart) {
		var ret []string
		return ret
	}
	return o.EmailOnStart
}

// GetEmailOnStartOk returns a tuple with the EmailOnStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayRecurringTaskResponse) GetEmailOnStartOk() ([]string, bool) {
	if o == nil || isNil(o.EmailOnStart) {
    return nil, false
	}
	return o.EmailOnStart, true
}

// HasEmailOnStart returns a boolean if a field has been set.
func (o *DelayRecurringTaskResponse) HasEmailOnStart() bool {
	if o != nil && !isNil(o.EmailOnStart) {
		return true
	}

	return false
}

// SetEmailOnStart gets a reference to the given []string and assigns it to the EmailOnStart field.
func (o *DelayRecurringTaskResponse) SetEmailOnStart(v []string) {
	o.EmailOnStart = v
}

// GetEmailOnSuccess returns the EmailOnSuccess field value if set, zero value otherwise.
func (o *DelayRecurringTaskResponse) GetEmailOnSuccess() []string {
	if o == nil || isNil(o.EmailOnSuccess) {
		var ret []string
		return ret
	}
	return o.EmailOnSuccess
}

// GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayRecurringTaskResponse) GetEmailOnSuccessOk() ([]string, bool) {
	if o == nil || isNil(o.EmailOnSuccess) {
    return nil, false
	}
	return o.EmailOnSuccess, true
}

// HasEmailOnSuccess returns a boolean if a field has been set.
func (o *DelayRecurringTaskResponse) HasEmailOnSuccess() bool {
	if o != nil && !isNil(o.EmailOnSuccess) {
		return true
	}

	return false
}

// SetEmailOnSuccess gets a reference to the given []string and assigns it to the EmailOnSuccess field.
func (o *DelayRecurringTaskResponse) SetEmailOnSuccess(v []string) {
	o.EmailOnSuccess = v
}

// GetEmailOnFailure returns the EmailOnFailure field value if set, zero value otherwise.
func (o *DelayRecurringTaskResponse) GetEmailOnFailure() []string {
	if o == nil || isNil(o.EmailOnFailure) {
		var ret []string
		return ret
	}
	return o.EmailOnFailure
}

// GetEmailOnFailureOk returns a tuple with the EmailOnFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayRecurringTaskResponse) GetEmailOnFailureOk() ([]string, bool) {
	if o == nil || isNil(o.EmailOnFailure) {
    return nil, false
	}
	return o.EmailOnFailure, true
}

// HasEmailOnFailure returns a boolean if a field has been set.
func (o *DelayRecurringTaskResponse) HasEmailOnFailure() bool {
	if o != nil && !isNil(o.EmailOnFailure) {
		return true
	}

	return false
}

// SetEmailOnFailure gets a reference to the given []string and assigns it to the EmailOnFailure field.
func (o *DelayRecurringTaskResponse) SetEmailOnFailure(v []string) {
	o.EmailOnFailure = v
}

// GetAlertOnStart returns the AlertOnStart field value if set, zero value otherwise.
func (o *DelayRecurringTaskResponse) GetAlertOnStart() bool {
	if o == nil || isNil(o.AlertOnStart) {
		var ret bool
		return ret
	}
	return *o.AlertOnStart
}

// GetAlertOnStartOk returns a tuple with the AlertOnStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayRecurringTaskResponse) GetAlertOnStartOk() (*bool, bool) {
	if o == nil || isNil(o.AlertOnStart) {
    return nil, false
	}
	return o.AlertOnStart, true
}

// HasAlertOnStart returns a boolean if a field has been set.
func (o *DelayRecurringTaskResponse) HasAlertOnStart() bool {
	if o != nil && !isNil(o.AlertOnStart) {
		return true
	}

	return false
}

// SetAlertOnStart gets a reference to the given bool and assigns it to the AlertOnStart field.
func (o *DelayRecurringTaskResponse) SetAlertOnStart(v bool) {
	o.AlertOnStart = &v
}

// GetAlertOnSuccess returns the AlertOnSuccess field value if set, zero value otherwise.
func (o *DelayRecurringTaskResponse) GetAlertOnSuccess() bool {
	if o == nil || isNil(o.AlertOnSuccess) {
		var ret bool
		return ret
	}
	return *o.AlertOnSuccess
}

// GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayRecurringTaskResponse) GetAlertOnSuccessOk() (*bool, bool) {
	if o == nil || isNil(o.AlertOnSuccess) {
    return nil, false
	}
	return o.AlertOnSuccess, true
}

// HasAlertOnSuccess returns a boolean if a field has been set.
func (o *DelayRecurringTaskResponse) HasAlertOnSuccess() bool {
	if o != nil && !isNil(o.AlertOnSuccess) {
		return true
	}

	return false
}

// SetAlertOnSuccess gets a reference to the given bool and assigns it to the AlertOnSuccess field.
func (o *DelayRecurringTaskResponse) SetAlertOnSuccess(v bool) {
	o.AlertOnSuccess = &v
}

// GetAlertOnFailure returns the AlertOnFailure field value if set, zero value otherwise.
func (o *DelayRecurringTaskResponse) GetAlertOnFailure() bool {
	if o == nil || isNil(o.AlertOnFailure) {
		var ret bool
		return ret
	}
	return *o.AlertOnFailure
}

// GetAlertOnFailureOk returns a tuple with the AlertOnFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayRecurringTaskResponse) GetAlertOnFailureOk() (*bool, bool) {
	if o == nil || isNil(o.AlertOnFailure) {
    return nil, false
	}
	return o.AlertOnFailure, true
}

// HasAlertOnFailure returns a boolean if a field has been set.
func (o *DelayRecurringTaskResponse) HasAlertOnFailure() bool {
	if o != nil && !isNil(o.AlertOnFailure) {
		return true
	}

	return false
}

// SetAlertOnFailure gets a reference to the given bool and assigns it to the AlertOnFailure field.
func (o *DelayRecurringTaskResponse) SetAlertOnFailure(v bool) {
	o.AlertOnFailure = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *DelayRecurringTaskResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayRecurringTaskResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *DelayRecurringTaskResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *DelayRecurringTaskResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

func (o DelayRecurringTaskResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.SleepDuration) {
		toSerialize["sleepDuration"] = o.SleepDuration
	}
	if !isNil(o.DurationToWaitForWorkQueueIdle) {
		toSerialize["durationToWaitForWorkQueueIdle"] = o.DurationToWaitForWorkQueueIdle
	}
	if !isNil(o.LdapURLForSearchExpectedToReturnEntries) {
		toSerialize["ldapURLForSearchExpectedToReturnEntries"] = o.LdapURLForSearchExpectedToReturnEntries
	}
	if !isNil(o.SearchInterval) {
		toSerialize["searchInterval"] = o.SearchInterval
	}
	if !isNil(o.SearchTimeLimit) {
		toSerialize["searchTimeLimit"] = o.SearchTimeLimit
	}
	if !isNil(o.DurationToWaitForSearchToReturnEntries) {
		toSerialize["durationToWaitForSearchToReturnEntries"] = o.DurationToWaitForSearchToReturnEntries
	}
	if !isNil(o.TaskReturnStateIfTimeoutIsEncountered) {
		toSerialize["taskReturnStateIfTimeoutIsEncountered"] = o.TaskReturnStateIfTimeoutIsEncountered
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
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableDelayRecurringTaskResponse struct {
	value *DelayRecurringTaskResponse
	isSet bool
}

func (v NullableDelayRecurringTaskResponse) Get() *DelayRecurringTaskResponse {
	return v.value
}

func (v *NullableDelayRecurringTaskResponse) Set(val *DelayRecurringTaskResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDelayRecurringTaskResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDelayRecurringTaskResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDelayRecurringTaskResponse(val *DelayRecurringTaskResponse) *NullableDelayRecurringTaskResponse {
	return &NullableDelayRecurringTaskResponse{value: val, isSet: true}
}

func (v NullableDelayRecurringTaskResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDelayRecurringTaskResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


