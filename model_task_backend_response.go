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

// TaskBackendResponse struct for TaskBackendResponse
type TaskBackendResponse struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas []EnumtaskBackendSchemaUrn `json:"schemas"`
	// Name of the Backend
	Id string `json:"id"`
	// Specifies a name to identify the associated backend.
	BackendID string `json:"backendID"`
	// Specifies the base DN(s) for the data that the backend handles.
	BaseDN []string `json:"baseDN"`
	WritabilityMode EnumbackendWritabilityModeProp `json:"writabilityMode"`
	// Specifies the path to the backing file for storing information about the tasks configured in the server.
	TaskBackingFile string `json:"taskBackingFile"`
	// The maximum number of log messages to retain in each task entry from the beginning of the processing for that task. If too many messages are logged during task processing, then retaining only a limited number of messages from the beginning and/or end of task processing can reduce the amount of memory that the server consumes by caching information about currently-active and recently-completed tasks.
	MaximumInitialTaskLogMessagesToRetain *int32 `json:"maximumInitialTaskLogMessagesToRetain,omitempty"`
	// The maximum number of log messages to retain in each task entry from the end of the processing for that task. If too many messages are logged during task processing, then retaining only a limited number of messages from the beginning and/or end of task processing can reduce the amount of memory that the server consumes by caching information about currently-active and recently-completed tasks.
	MaximumFinalTaskLogMessagesToRetain *int32 `json:"maximumFinalTaskLogMessagesToRetain,omitempty"`
	// Specifies the length of time that task entries should be retained after processing on the associated task has been completed.
	TaskRetentionTime *string `json:"taskRetentionTime,omitempty"`
	// Specifies the email address to use as the sender address (that is, the \"From:\" address) for notification mail messages generated when a task completes execution.
	NotificationSenderAddress *string `json:"notificationSenderAddress,omitempty"`
	// A description for this Backend
	Description *string `json:"description,omitempty"`
	// Indicates whether the backend is enabled in the server.
	Enabled bool `json:"enabled"`
	// Determines whether the Directory Server enters a DEGRADED state (and sends a corresponding alert) when this Backend is disabled.
	SetDegradedAlertWhenDisabled *bool `json:"setDegradedAlertWhenDisabled,omitempty"`
	// Determines whether any LDAP operation that would use this Backend is to return UNAVAILABLE when this Backend is disabled.
	ReturnUnavailableWhenDisabled *bool `json:"returnUnavailableWhenDisabled,omitempty"`
	// Specifies the permissions that should be applied to files and directories created by a backup of the backend.
	BackupFilePermissions *string `json:"backupFilePermissions,omitempty"`
	// Specifies a notification manager for changes resulting from operations processed through this Backend
	NotificationManager *string `json:"notificationManager,omitempty"`
}

// NewTaskBackendResponse instantiates a new TaskBackendResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskBackendResponse(schemas []EnumtaskBackendSchemaUrn, id string, backendID string, baseDN []string, writabilityMode EnumbackendWritabilityModeProp, taskBackingFile string, enabled bool) *TaskBackendResponse {
	this := TaskBackendResponse{}
	this.Schemas = schemas
	this.Id = id
	this.BackendID = backendID
	this.BaseDN = baseDN
	this.WritabilityMode = writabilityMode
	this.TaskBackingFile = taskBackingFile
	this.Enabled = enabled
	return &this
}

// NewTaskBackendResponseWithDefaults instantiates a new TaskBackendResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskBackendResponseWithDefaults() *TaskBackendResponse {
	this := TaskBackendResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *TaskBackendResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskBackendResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *TaskBackendResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *TaskBackendResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *TaskBackendResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskBackendResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
    return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *TaskBackendResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *TaskBackendResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *TaskBackendResponse) GetSchemas() []EnumtaskBackendSchemaUrn {
	if o == nil {
		var ret []EnumtaskBackendSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *TaskBackendResponse) GetSchemasOk() ([]EnumtaskBackendSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *TaskBackendResponse) SetSchemas(v []EnumtaskBackendSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *TaskBackendResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TaskBackendResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TaskBackendResponse) SetId(v string) {
	o.Id = v
}

// GetBackendID returns the BackendID field value
func (o *TaskBackendResponse) GetBackendID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BackendID
}

// GetBackendIDOk returns a tuple with the BackendID field value
// and a boolean to check if the value has been set.
func (o *TaskBackendResponse) GetBackendIDOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.BackendID, true
}

// SetBackendID sets field value
func (o *TaskBackendResponse) SetBackendID(v string) {
	o.BackendID = v
}

// GetBaseDN returns the BaseDN field value
func (o *TaskBackendResponse) GetBaseDN() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BaseDN
}

// GetBaseDNOk returns a tuple with the BaseDN field value
// and a boolean to check if the value has been set.
func (o *TaskBackendResponse) GetBaseDNOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.BaseDN, true
}

// SetBaseDN sets field value
func (o *TaskBackendResponse) SetBaseDN(v []string) {
	o.BaseDN = v
}

// GetWritabilityMode returns the WritabilityMode field value
func (o *TaskBackendResponse) GetWritabilityMode() EnumbackendWritabilityModeProp {
	if o == nil {
		var ret EnumbackendWritabilityModeProp
		return ret
	}

	return o.WritabilityMode
}

// GetWritabilityModeOk returns a tuple with the WritabilityMode field value
// and a boolean to check if the value has been set.
func (o *TaskBackendResponse) GetWritabilityModeOk() (*EnumbackendWritabilityModeProp, bool) {
	if o == nil {
    return nil, false
	}
	return &o.WritabilityMode, true
}

// SetWritabilityMode sets field value
func (o *TaskBackendResponse) SetWritabilityMode(v EnumbackendWritabilityModeProp) {
	o.WritabilityMode = v
}

// GetTaskBackingFile returns the TaskBackingFile field value
func (o *TaskBackendResponse) GetTaskBackingFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaskBackingFile
}

// GetTaskBackingFileOk returns a tuple with the TaskBackingFile field value
// and a boolean to check if the value has been set.
func (o *TaskBackendResponse) GetTaskBackingFileOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TaskBackingFile, true
}

// SetTaskBackingFile sets field value
func (o *TaskBackendResponse) SetTaskBackingFile(v string) {
	o.TaskBackingFile = v
}

// GetMaximumInitialTaskLogMessagesToRetain returns the MaximumInitialTaskLogMessagesToRetain field value if set, zero value otherwise.
func (o *TaskBackendResponse) GetMaximumInitialTaskLogMessagesToRetain() int32 {
	if o == nil || isNil(o.MaximumInitialTaskLogMessagesToRetain) {
		var ret int32
		return ret
	}
	return *o.MaximumInitialTaskLogMessagesToRetain
}

// GetMaximumInitialTaskLogMessagesToRetainOk returns a tuple with the MaximumInitialTaskLogMessagesToRetain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskBackendResponse) GetMaximumInitialTaskLogMessagesToRetainOk() (*int32, bool) {
	if o == nil || isNil(o.MaximumInitialTaskLogMessagesToRetain) {
    return nil, false
	}
	return o.MaximumInitialTaskLogMessagesToRetain, true
}

// HasMaximumInitialTaskLogMessagesToRetain returns a boolean if a field has been set.
func (o *TaskBackendResponse) HasMaximumInitialTaskLogMessagesToRetain() bool {
	if o != nil && !isNil(o.MaximumInitialTaskLogMessagesToRetain) {
		return true
	}

	return false
}

// SetMaximumInitialTaskLogMessagesToRetain gets a reference to the given int32 and assigns it to the MaximumInitialTaskLogMessagesToRetain field.
func (o *TaskBackendResponse) SetMaximumInitialTaskLogMessagesToRetain(v int32) {
	o.MaximumInitialTaskLogMessagesToRetain = &v
}

// GetMaximumFinalTaskLogMessagesToRetain returns the MaximumFinalTaskLogMessagesToRetain field value if set, zero value otherwise.
func (o *TaskBackendResponse) GetMaximumFinalTaskLogMessagesToRetain() int32 {
	if o == nil || isNil(o.MaximumFinalTaskLogMessagesToRetain) {
		var ret int32
		return ret
	}
	return *o.MaximumFinalTaskLogMessagesToRetain
}

// GetMaximumFinalTaskLogMessagesToRetainOk returns a tuple with the MaximumFinalTaskLogMessagesToRetain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskBackendResponse) GetMaximumFinalTaskLogMessagesToRetainOk() (*int32, bool) {
	if o == nil || isNil(o.MaximumFinalTaskLogMessagesToRetain) {
    return nil, false
	}
	return o.MaximumFinalTaskLogMessagesToRetain, true
}

// HasMaximumFinalTaskLogMessagesToRetain returns a boolean if a field has been set.
func (o *TaskBackendResponse) HasMaximumFinalTaskLogMessagesToRetain() bool {
	if o != nil && !isNil(o.MaximumFinalTaskLogMessagesToRetain) {
		return true
	}

	return false
}

// SetMaximumFinalTaskLogMessagesToRetain gets a reference to the given int32 and assigns it to the MaximumFinalTaskLogMessagesToRetain field.
func (o *TaskBackendResponse) SetMaximumFinalTaskLogMessagesToRetain(v int32) {
	o.MaximumFinalTaskLogMessagesToRetain = &v
}

// GetTaskRetentionTime returns the TaskRetentionTime field value if set, zero value otherwise.
func (o *TaskBackendResponse) GetTaskRetentionTime() string {
	if o == nil || isNil(o.TaskRetentionTime) {
		var ret string
		return ret
	}
	return *o.TaskRetentionTime
}

// GetTaskRetentionTimeOk returns a tuple with the TaskRetentionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskBackendResponse) GetTaskRetentionTimeOk() (*string, bool) {
	if o == nil || isNil(o.TaskRetentionTime) {
    return nil, false
	}
	return o.TaskRetentionTime, true
}

// HasTaskRetentionTime returns a boolean if a field has been set.
func (o *TaskBackendResponse) HasTaskRetentionTime() bool {
	if o != nil && !isNil(o.TaskRetentionTime) {
		return true
	}

	return false
}

// SetTaskRetentionTime gets a reference to the given string and assigns it to the TaskRetentionTime field.
func (o *TaskBackendResponse) SetTaskRetentionTime(v string) {
	o.TaskRetentionTime = &v
}

// GetNotificationSenderAddress returns the NotificationSenderAddress field value if set, zero value otherwise.
func (o *TaskBackendResponse) GetNotificationSenderAddress() string {
	if o == nil || isNil(o.NotificationSenderAddress) {
		var ret string
		return ret
	}
	return *o.NotificationSenderAddress
}

// GetNotificationSenderAddressOk returns a tuple with the NotificationSenderAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskBackendResponse) GetNotificationSenderAddressOk() (*string, bool) {
	if o == nil || isNil(o.NotificationSenderAddress) {
    return nil, false
	}
	return o.NotificationSenderAddress, true
}

// HasNotificationSenderAddress returns a boolean if a field has been set.
func (o *TaskBackendResponse) HasNotificationSenderAddress() bool {
	if o != nil && !isNil(o.NotificationSenderAddress) {
		return true
	}

	return false
}

// SetNotificationSenderAddress gets a reference to the given string and assigns it to the NotificationSenderAddress field.
func (o *TaskBackendResponse) SetNotificationSenderAddress(v string) {
	o.NotificationSenderAddress = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TaskBackendResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskBackendResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TaskBackendResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TaskBackendResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *TaskBackendResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *TaskBackendResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *TaskBackendResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSetDegradedAlertWhenDisabled returns the SetDegradedAlertWhenDisabled field value if set, zero value otherwise.
func (o *TaskBackendResponse) GetSetDegradedAlertWhenDisabled() bool {
	if o == nil || isNil(o.SetDegradedAlertWhenDisabled) {
		var ret bool
		return ret
	}
	return *o.SetDegradedAlertWhenDisabled
}

// GetSetDegradedAlertWhenDisabledOk returns a tuple with the SetDegradedAlertWhenDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskBackendResponse) GetSetDegradedAlertWhenDisabledOk() (*bool, bool) {
	if o == nil || isNil(o.SetDegradedAlertWhenDisabled) {
    return nil, false
	}
	return o.SetDegradedAlertWhenDisabled, true
}

// HasSetDegradedAlertWhenDisabled returns a boolean if a field has been set.
func (o *TaskBackendResponse) HasSetDegradedAlertWhenDisabled() bool {
	if o != nil && !isNil(o.SetDegradedAlertWhenDisabled) {
		return true
	}

	return false
}

// SetSetDegradedAlertWhenDisabled gets a reference to the given bool and assigns it to the SetDegradedAlertWhenDisabled field.
func (o *TaskBackendResponse) SetSetDegradedAlertWhenDisabled(v bool) {
	o.SetDegradedAlertWhenDisabled = &v
}

// GetReturnUnavailableWhenDisabled returns the ReturnUnavailableWhenDisabled field value if set, zero value otherwise.
func (o *TaskBackendResponse) GetReturnUnavailableWhenDisabled() bool {
	if o == nil || isNil(o.ReturnUnavailableWhenDisabled) {
		var ret bool
		return ret
	}
	return *o.ReturnUnavailableWhenDisabled
}

// GetReturnUnavailableWhenDisabledOk returns a tuple with the ReturnUnavailableWhenDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskBackendResponse) GetReturnUnavailableWhenDisabledOk() (*bool, bool) {
	if o == nil || isNil(o.ReturnUnavailableWhenDisabled) {
    return nil, false
	}
	return o.ReturnUnavailableWhenDisabled, true
}

// HasReturnUnavailableWhenDisabled returns a boolean if a field has been set.
func (o *TaskBackendResponse) HasReturnUnavailableWhenDisabled() bool {
	if o != nil && !isNil(o.ReturnUnavailableWhenDisabled) {
		return true
	}

	return false
}

// SetReturnUnavailableWhenDisabled gets a reference to the given bool and assigns it to the ReturnUnavailableWhenDisabled field.
func (o *TaskBackendResponse) SetReturnUnavailableWhenDisabled(v bool) {
	o.ReturnUnavailableWhenDisabled = &v
}

// GetBackupFilePermissions returns the BackupFilePermissions field value if set, zero value otherwise.
func (o *TaskBackendResponse) GetBackupFilePermissions() string {
	if o == nil || isNil(o.BackupFilePermissions) {
		var ret string
		return ret
	}
	return *o.BackupFilePermissions
}

// GetBackupFilePermissionsOk returns a tuple with the BackupFilePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskBackendResponse) GetBackupFilePermissionsOk() (*string, bool) {
	if o == nil || isNil(o.BackupFilePermissions) {
    return nil, false
	}
	return o.BackupFilePermissions, true
}

// HasBackupFilePermissions returns a boolean if a field has been set.
func (o *TaskBackendResponse) HasBackupFilePermissions() bool {
	if o != nil && !isNil(o.BackupFilePermissions) {
		return true
	}

	return false
}

// SetBackupFilePermissions gets a reference to the given string and assigns it to the BackupFilePermissions field.
func (o *TaskBackendResponse) SetBackupFilePermissions(v string) {
	o.BackupFilePermissions = &v
}

// GetNotificationManager returns the NotificationManager field value if set, zero value otherwise.
func (o *TaskBackendResponse) GetNotificationManager() string {
	if o == nil || isNil(o.NotificationManager) {
		var ret string
		return ret
	}
	return *o.NotificationManager
}

// GetNotificationManagerOk returns a tuple with the NotificationManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskBackendResponse) GetNotificationManagerOk() (*string, bool) {
	if o == nil || isNil(o.NotificationManager) {
    return nil, false
	}
	return o.NotificationManager, true
}

// HasNotificationManager returns a boolean if a field has been set.
func (o *TaskBackendResponse) HasNotificationManager() bool {
	if o != nil && !isNil(o.NotificationManager) {
		return true
	}

	return false
}

// SetNotificationManager gets a reference to the given string and assigns it to the NotificationManager field.
func (o *TaskBackendResponse) SetNotificationManager(v string) {
	o.NotificationManager = &v
}

func (o TaskBackendResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["backendID"] = o.BackendID
	}
	if true {
		toSerialize["baseDN"] = o.BaseDN
	}
	if true {
		toSerialize["writabilityMode"] = o.WritabilityMode
	}
	if true {
		toSerialize["taskBackingFile"] = o.TaskBackingFile
	}
	if !isNil(o.MaximumInitialTaskLogMessagesToRetain) {
		toSerialize["maximumInitialTaskLogMessagesToRetain"] = o.MaximumInitialTaskLogMessagesToRetain
	}
	if !isNil(o.MaximumFinalTaskLogMessagesToRetain) {
		toSerialize["maximumFinalTaskLogMessagesToRetain"] = o.MaximumFinalTaskLogMessagesToRetain
	}
	if !isNil(o.TaskRetentionTime) {
		toSerialize["taskRetentionTime"] = o.TaskRetentionTime
	}
	if !isNil(o.NotificationSenderAddress) {
		toSerialize["notificationSenderAddress"] = o.NotificationSenderAddress
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.SetDegradedAlertWhenDisabled) {
		toSerialize["setDegradedAlertWhenDisabled"] = o.SetDegradedAlertWhenDisabled
	}
	if !isNil(o.ReturnUnavailableWhenDisabled) {
		toSerialize["returnUnavailableWhenDisabled"] = o.ReturnUnavailableWhenDisabled
	}
	if !isNil(o.BackupFilePermissions) {
		toSerialize["backupFilePermissions"] = o.BackupFilePermissions
	}
	if !isNil(o.NotificationManager) {
		toSerialize["notificationManager"] = o.NotificationManager
	}
	return json.Marshal(toSerialize)
}

type NullableTaskBackendResponse struct {
	value *TaskBackendResponse
	isSet bool
}

func (v NullableTaskBackendResponse) Get() *TaskBackendResponse {
	return v.value
}

func (v *NullableTaskBackendResponse) Set(val *TaskBackendResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskBackendResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskBackendResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskBackendResponse(val *TaskBackendResponse) *NullableTaskBackendResponse {
	return &NullableTaskBackendResponse{value: val, isSet: true}
}

func (v NullableTaskBackendResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskBackendResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


