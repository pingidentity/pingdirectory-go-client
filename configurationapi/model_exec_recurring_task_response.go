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

// checks if the ExecRecurringTaskResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExecRecurringTaskResponse{}

// ExecRecurringTaskResponse struct for ExecRecurringTaskResponse
type ExecRecurringTaskResponse struct {
	// Name of the Recurring Task
	Id      string                           `json:"id"`
	Schemas []EnumexecRecurringTaskSchemaUrn `json:"schemas"`
	// The absolute path to the command to execute. It must be an absolute path, the corresponding file must exist, and it must be listed in the config/exec-command-whitelist.txt file.
	CommandPath string `json:"commandPath"`
	// A string containing the arguments to provide to the command. If the command should be run without arguments, this property should be left undefined. If there should be multiple arguments, then they should be separated with spaces.
	CommandArguments *string `json:"commandArguments,omitempty"`
	// The path and base name for a file to which the command output (both standard output and standard error) should be written. This may be left undefined if the command output should not be recorded into a file.
	CommandOutputFileBaseName *string `json:"commandOutputFileBaseName,omitempty"`
	// The minimum number of previous command output files that should be preserved after a new instance of the command is invoked.
	RetainPreviousOutputFileCount *int32 `json:"retainPreviousOutputFileCount,omitempty"`
	// The minimum age of previous command output files that should be preserved after a new instance of the command is invoked.
	RetainPreviousOutputFileAge *string `json:"retainPreviousOutputFileAge,omitempty"`
	// Indicates whether the command's output (both standard output and standard error) should be recorded in the server's error log.
	LogCommandOutput                      *bool                                                       `json:"logCommandOutput,omitempty"`
	TaskCompletionStateForNonzeroExitCode *EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp `json:"taskCompletionStateForNonzeroExitCode,omitempty"`
	// The absolute path to a working directory where the command should be executed. It must be an absolute path and the corresponding directory must exist.
	WorkingDirectory *string `json:"workingDirectory,omitempty"`
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
	AlertOnFailure                                *bool                                              `json:"alertOnFailure,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewExecRecurringTaskResponse instantiates a new ExecRecurringTaskResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecRecurringTaskResponse(id string, schemas []EnumexecRecurringTaskSchemaUrn, commandPath string) *ExecRecurringTaskResponse {
	this := ExecRecurringTaskResponse{}
	this.Id = id
	this.Schemas = schemas
	this.CommandPath = commandPath
	return &this
}

// NewExecRecurringTaskResponseWithDefaults instantiates a new ExecRecurringTaskResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecRecurringTaskResponseWithDefaults() *ExecRecurringTaskResponse {
	this := ExecRecurringTaskResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ExecRecurringTaskResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExecRecurringTaskResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExecRecurringTaskResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *ExecRecurringTaskResponse) GetSchemas() []EnumexecRecurringTaskSchemaUrn {
	if o == nil {
		var ret []EnumexecRecurringTaskSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ExecRecurringTaskResponse) GetSchemasOk() ([]EnumexecRecurringTaskSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ExecRecurringTaskResponse) SetSchemas(v []EnumexecRecurringTaskSchemaUrn) {
	o.Schemas = v
}

// GetCommandPath returns the CommandPath field value
func (o *ExecRecurringTaskResponse) GetCommandPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommandPath
}

// GetCommandPathOk returns a tuple with the CommandPath field value
// and a boolean to check if the value has been set.
func (o *ExecRecurringTaskResponse) GetCommandPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommandPath, true
}

// SetCommandPath sets field value
func (o *ExecRecurringTaskResponse) SetCommandPath(v string) {
	o.CommandPath = v
}

// GetCommandArguments returns the CommandArguments field value if set, zero value otherwise.
func (o *ExecRecurringTaskResponse) GetCommandArguments() string {
	if o == nil || IsNil(o.CommandArguments) {
		var ret string
		return ret
	}
	return *o.CommandArguments
}

// GetCommandArgumentsOk returns a tuple with the CommandArguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecRecurringTaskResponse) GetCommandArgumentsOk() (*string, bool) {
	if o == nil || IsNil(o.CommandArguments) {
		return nil, false
	}
	return o.CommandArguments, true
}

// HasCommandArguments returns a boolean if a field has been set.
func (o *ExecRecurringTaskResponse) HasCommandArguments() bool {
	if o != nil && !IsNil(o.CommandArguments) {
		return true
	}

	return false
}

// SetCommandArguments gets a reference to the given string and assigns it to the CommandArguments field.
func (o *ExecRecurringTaskResponse) SetCommandArguments(v string) {
	o.CommandArguments = &v
}

// GetCommandOutputFileBaseName returns the CommandOutputFileBaseName field value if set, zero value otherwise.
func (o *ExecRecurringTaskResponse) GetCommandOutputFileBaseName() string {
	if o == nil || IsNil(o.CommandOutputFileBaseName) {
		var ret string
		return ret
	}
	return *o.CommandOutputFileBaseName
}

// GetCommandOutputFileBaseNameOk returns a tuple with the CommandOutputFileBaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecRecurringTaskResponse) GetCommandOutputFileBaseNameOk() (*string, bool) {
	if o == nil || IsNil(o.CommandOutputFileBaseName) {
		return nil, false
	}
	return o.CommandOutputFileBaseName, true
}

// HasCommandOutputFileBaseName returns a boolean if a field has been set.
func (o *ExecRecurringTaskResponse) HasCommandOutputFileBaseName() bool {
	if o != nil && !IsNil(o.CommandOutputFileBaseName) {
		return true
	}

	return false
}

// SetCommandOutputFileBaseName gets a reference to the given string and assigns it to the CommandOutputFileBaseName field.
func (o *ExecRecurringTaskResponse) SetCommandOutputFileBaseName(v string) {
	o.CommandOutputFileBaseName = &v
}

// GetRetainPreviousOutputFileCount returns the RetainPreviousOutputFileCount field value if set, zero value otherwise.
func (o *ExecRecurringTaskResponse) GetRetainPreviousOutputFileCount() int32 {
	if o == nil || IsNil(o.RetainPreviousOutputFileCount) {
		var ret int32
		return ret
	}
	return *o.RetainPreviousOutputFileCount
}

// GetRetainPreviousOutputFileCountOk returns a tuple with the RetainPreviousOutputFileCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecRecurringTaskResponse) GetRetainPreviousOutputFileCountOk() (*int32, bool) {
	if o == nil || IsNil(o.RetainPreviousOutputFileCount) {
		return nil, false
	}
	return o.RetainPreviousOutputFileCount, true
}

// HasRetainPreviousOutputFileCount returns a boolean if a field has been set.
func (o *ExecRecurringTaskResponse) HasRetainPreviousOutputFileCount() bool {
	if o != nil && !IsNil(o.RetainPreviousOutputFileCount) {
		return true
	}

	return false
}

// SetRetainPreviousOutputFileCount gets a reference to the given int32 and assigns it to the RetainPreviousOutputFileCount field.
func (o *ExecRecurringTaskResponse) SetRetainPreviousOutputFileCount(v int32) {
	o.RetainPreviousOutputFileCount = &v
}

// GetRetainPreviousOutputFileAge returns the RetainPreviousOutputFileAge field value if set, zero value otherwise.
func (o *ExecRecurringTaskResponse) GetRetainPreviousOutputFileAge() string {
	if o == nil || IsNil(o.RetainPreviousOutputFileAge) {
		var ret string
		return ret
	}
	return *o.RetainPreviousOutputFileAge
}

// GetRetainPreviousOutputFileAgeOk returns a tuple with the RetainPreviousOutputFileAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecRecurringTaskResponse) GetRetainPreviousOutputFileAgeOk() (*string, bool) {
	if o == nil || IsNil(o.RetainPreviousOutputFileAge) {
		return nil, false
	}
	return o.RetainPreviousOutputFileAge, true
}

// HasRetainPreviousOutputFileAge returns a boolean if a field has been set.
func (o *ExecRecurringTaskResponse) HasRetainPreviousOutputFileAge() bool {
	if o != nil && !IsNil(o.RetainPreviousOutputFileAge) {
		return true
	}

	return false
}

// SetRetainPreviousOutputFileAge gets a reference to the given string and assigns it to the RetainPreviousOutputFileAge field.
func (o *ExecRecurringTaskResponse) SetRetainPreviousOutputFileAge(v string) {
	o.RetainPreviousOutputFileAge = &v
}

// GetLogCommandOutput returns the LogCommandOutput field value if set, zero value otherwise.
func (o *ExecRecurringTaskResponse) GetLogCommandOutput() bool {
	if o == nil || IsNil(o.LogCommandOutput) {
		var ret bool
		return ret
	}
	return *o.LogCommandOutput
}

// GetLogCommandOutputOk returns a tuple with the LogCommandOutput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecRecurringTaskResponse) GetLogCommandOutputOk() (*bool, bool) {
	if o == nil || IsNil(o.LogCommandOutput) {
		return nil, false
	}
	return o.LogCommandOutput, true
}

// HasLogCommandOutput returns a boolean if a field has been set.
func (o *ExecRecurringTaskResponse) HasLogCommandOutput() bool {
	if o != nil && !IsNil(o.LogCommandOutput) {
		return true
	}

	return false
}

// SetLogCommandOutput gets a reference to the given bool and assigns it to the LogCommandOutput field.
func (o *ExecRecurringTaskResponse) SetLogCommandOutput(v bool) {
	o.LogCommandOutput = &v
}

// GetTaskCompletionStateForNonzeroExitCode returns the TaskCompletionStateForNonzeroExitCode field value if set, zero value otherwise.
func (o *ExecRecurringTaskResponse) GetTaskCompletionStateForNonzeroExitCode() EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp {
	if o == nil || IsNil(o.TaskCompletionStateForNonzeroExitCode) {
		var ret EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp
		return ret
	}
	return *o.TaskCompletionStateForNonzeroExitCode
}

// GetTaskCompletionStateForNonzeroExitCodeOk returns a tuple with the TaskCompletionStateForNonzeroExitCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecRecurringTaskResponse) GetTaskCompletionStateForNonzeroExitCodeOk() (*EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp, bool) {
	if o == nil || IsNil(o.TaskCompletionStateForNonzeroExitCode) {
		return nil, false
	}
	return o.TaskCompletionStateForNonzeroExitCode, true
}

// HasTaskCompletionStateForNonzeroExitCode returns a boolean if a field has been set.
func (o *ExecRecurringTaskResponse) HasTaskCompletionStateForNonzeroExitCode() bool {
	if o != nil && !IsNil(o.TaskCompletionStateForNonzeroExitCode) {
		return true
	}

	return false
}

// SetTaskCompletionStateForNonzeroExitCode gets a reference to the given EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp and assigns it to the TaskCompletionStateForNonzeroExitCode field.
func (o *ExecRecurringTaskResponse) SetTaskCompletionStateForNonzeroExitCode(v EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp) {
	o.TaskCompletionStateForNonzeroExitCode = &v
}

// GetWorkingDirectory returns the WorkingDirectory field value if set, zero value otherwise.
func (o *ExecRecurringTaskResponse) GetWorkingDirectory() string {
	if o == nil || IsNil(o.WorkingDirectory) {
		var ret string
		return ret
	}
	return *o.WorkingDirectory
}

// GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecRecurringTaskResponse) GetWorkingDirectoryOk() (*string, bool) {
	if o == nil || IsNil(o.WorkingDirectory) {
		return nil, false
	}
	return o.WorkingDirectory, true
}

// HasWorkingDirectory returns a boolean if a field has been set.
func (o *ExecRecurringTaskResponse) HasWorkingDirectory() bool {
	if o != nil && !IsNil(o.WorkingDirectory) {
		return true
	}

	return false
}

// SetWorkingDirectory gets a reference to the given string and assigns it to the WorkingDirectory field.
func (o *ExecRecurringTaskResponse) SetWorkingDirectory(v string) {
	o.WorkingDirectory = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExecRecurringTaskResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecRecurringTaskResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExecRecurringTaskResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExecRecurringTaskResponse) SetDescription(v string) {
	o.Description = &v
}

// GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field value if set, zero value otherwise.
func (o *ExecRecurringTaskResponse) GetCancelOnTaskDependencyFailure() bool {
	if o == nil || IsNil(o.CancelOnTaskDependencyFailure) {
		var ret bool
		return ret
	}
	return *o.CancelOnTaskDependencyFailure
}

// GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecRecurringTaskResponse) GetCancelOnTaskDependencyFailureOk() (*bool, bool) {
	if o == nil || IsNil(o.CancelOnTaskDependencyFailure) {
		return nil, false
	}
	return o.CancelOnTaskDependencyFailure, true
}

// HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.
func (o *ExecRecurringTaskResponse) HasCancelOnTaskDependencyFailure() bool {
	if o != nil && !IsNil(o.CancelOnTaskDependencyFailure) {
		return true
	}

	return false
}

// SetCancelOnTaskDependencyFailure gets a reference to the given bool and assigns it to the CancelOnTaskDependencyFailure field.
func (o *ExecRecurringTaskResponse) SetCancelOnTaskDependencyFailure(v bool) {
	o.CancelOnTaskDependencyFailure = &v
}

// GetEmailOnStart returns the EmailOnStart field value if set, zero value otherwise.
func (o *ExecRecurringTaskResponse) GetEmailOnStart() []string {
	if o == nil || IsNil(o.EmailOnStart) {
		var ret []string
		return ret
	}
	return o.EmailOnStart
}

// GetEmailOnStartOk returns a tuple with the EmailOnStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecRecurringTaskResponse) GetEmailOnStartOk() ([]string, bool) {
	if o == nil || IsNil(o.EmailOnStart) {
		return nil, false
	}
	return o.EmailOnStart, true
}

// HasEmailOnStart returns a boolean if a field has been set.
func (o *ExecRecurringTaskResponse) HasEmailOnStart() bool {
	if o != nil && !IsNil(o.EmailOnStart) {
		return true
	}

	return false
}

// SetEmailOnStart gets a reference to the given []string and assigns it to the EmailOnStart field.
func (o *ExecRecurringTaskResponse) SetEmailOnStart(v []string) {
	o.EmailOnStart = v
}

// GetEmailOnSuccess returns the EmailOnSuccess field value if set, zero value otherwise.
func (o *ExecRecurringTaskResponse) GetEmailOnSuccess() []string {
	if o == nil || IsNil(o.EmailOnSuccess) {
		var ret []string
		return ret
	}
	return o.EmailOnSuccess
}

// GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecRecurringTaskResponse) GetEmailOnSuccessOk() ([]string, bool) {
	if o == nil || IsNil(o.EmailOnSuccess) {
		return nil, false
	}
	return o.EmailOnSuccess, true
}

// HasEmailOnSuccess returns a boolean if a field has been set.
func (o *ExecRecurringTaskResponse) HasEmailOnSuccess() bool {
	if o != nil && !IsNil(o.EmailOnSuccess) {
		return true
	}

	return false
}

// SetEmailOnSuccess gets a reference to the given []string and assigns it to the EmailOnSuccess field.
func (o *ExecRecurringTaskResponse) SetEmailOnSuccess(v []string) {
	o.EmailOnSuccess = v
}

// GetEmailOnFailure returns the EmailOnFailure field value if set, zero value otherwise.
func (o *ExecRecurringTaskResponse) GetEmailOnFailure() []string {
	if o == nil || IsNil(o.EmailOnFailure) {
		var ret []string
		return ret
	}
	return o.EmailOnFailure
}

// GetEmailOnFailureOk returns a tuple with the EmailOnFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecRecurringTaskResponse) GetEmailOnFailureOk() ([]string, bool) {
	if o == nil || IsNil(o.EmailOnFailure) {
		return nil, false
	}
	return o.EmailOnFailure, true
}

// HasEmailOnFailure returns a boolean if a field has been set.
func (o *ExecRecurringTaskResponse) HasEmailOnFailure() bool {
	if o != nil && !IsNil(o.EmailOnFailure) {
		return true
	}

	return false
}

// SetEmailOnFailure gets a reference to the given []string and assigns it to the EmailOnFailure field.
func (o *ExecRecurringTaskResponse) SetEmailOnFailure(v []string) {
	o.EmailOnFailure = v
}

// GetAlertOnStart returns the AlertOnStart field value if set, zero value otherwise.
func (o *ExecRecurringTaskResponse) GetAlertOnStart() bool {
	if o == nil || IsNil(o.AlertOnStart) {
		var ret bool
		return ret
	}
	return *o.AlertOnStart
}

// GetAlertOnStartOk returns a tuple with the AlertOnStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecRecurringTaskResponse) GetAlertOnStartOk() (*bool, bool) {
	if o == nil || IsNil(o.AlertOnStart) {
		return nil, false
	}
	return o.AlertOnStart, true
}

// HasAlertOnStart returns a boolean if a field has been set.
func (o *ExecRecurringTaskResponse) HasAlertOnStart() bool {
	if o != nil && !IsNil(o.AlertOnStart) {
		return true
	}

	return false
}

// SetAlertOnStart gets a reference to the given bool and assigns it to the AlertOnStart field.
func (o *ExecRecurringTaskResponse) SetAlertOnStart(v bool) {
	o.AlertOnStart = &v
}

// GetAlertOnSuccess returns the AlertOnSuccess field value if set, zero value otherwise.
func (o *ExecRecurringTaskResponse) GetAlertOnSuccess() bool {
	if o == nil || IsNil(o.AlertOnSuccess) {
		var ret bool
		return ret
	}
	return *o.AlertOnSuccess
}

// GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecRecurringTaskResponse) GetAlertOnSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.AlertOnSuccess) {
		return nil, false
	}
	return o.AlertOnSuccess, true
}

// HasAlertOnSuccess returns a boolean if a field has been set.
func (o *ExecRecurringTaskResponse) HasAlertOnSuccess() bool {
	if o != nil && !IsNil(o.AlertOnSuccess) {
		return true
	}

	return false
}

// SetAlertOnSuccess gets a reference to the given bool and assigns it to the AlertOnSuccess field.
func (o *ExecRecurringTaskResponse) SetAlertOnSuccess(v bool) {
	o.AlertOnSuccess = &v
}

// GetAlertOnFailure returns the AlertOnFailure field value if set, zero value otherwise.
func (o *ExecRecurringTaskResponse) GetAlertOnFailure() bool {
	if o == nil || IsNil(o.AlertOnFailure) {
		var ret bool
		return ret
	}
	return *o.AlertOnFailure
}

// GetAlertOnFailureOk returns a tuple with the AlertOnFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecRecurringTaskResponse) GetAlertOnFailureOk() (*bool, bool) {
	if o == nil || IsNil(o.AlertOnFailure) {
		return nil, false
	}
	return o.AlertOnFailure, true
}

// HasAlertOnFailure returns a boolean if a field has been set.
func (o *ExecRecurringTaskResponse) HasAlertOnFailure() bool {
	if o != nil && !IsNil(o.AlertOnFailure) {
		return true
	}

	return false
}

// SetAlertOnFailure gets a reference to the given bool and assigns it to the AlertOnFailure field.
func (o *ExecRecurringTaskResponse) SetAlertOnFailure(v bool) {
	o.AlertOnFailure = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ExecRecurringTaskResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecRecurringTaskResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ExecRecurringTaskResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ExecRecurringTaskResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ExecRecurringTaskResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecRecurringTaskResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ExecRecurringTaskResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ExecRecurringTaskResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o ExecRecurringTaskResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecRecurringTaskResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["schemas"] = o.Schemas
	toSerialize["commandPath"] = o.CommandPath
	if !IsNil(o.CommandArguments) {
		toSerialize["commandArguments"] = o.CommandArguments
	}
	if !IsNil(o.CommandOutputFileBaseName) {
		toSerialize["commandOutputFileBaseName"] = o.CommandOutputFileBaseName
	}
	if !IsNil(o.RetainPreviousOutputFileCount) {
		toSerialize["retainPreviousOutputFileCount"] = o.RetainPreviousOutputFileCount
	}
	if !IsNil(o.RetainPreviousOutputFileAge) {
		toSerialize["retainPreviousOutputFileAge"] = o.RetainPreviousOutputFileAge
	}
	if !IsNil(o.LogCommandOutput) {
		toSerialize["logCommandOutput"] = o.LogCommandOutput
	}
	if !IsNil(o.TaskCompletionStateForNonzeroExitCode) {
		toSerialize["taskCompletionStateForNonzeroExitCode"] = o.TaskCompletionStateForNonzeroExitCode
	}
	if !IsNil(o.WorkingDirectory) {
		toSerialize["workingDirectory"] = o.WorkingDirectory
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
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableExecRecurringTaskResponse struct {
	value *ExecRecurringTaskResponse
	isSet bool
}

func (v NullableExecRecurringTaskResponse) Get() *ExecRecurringTaskResponse {
	return v.value
}

func (v *NullableExecRecurringTaskResponse) Set(val *ExecRecurringTaskResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExecRecurringTaskResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExecRecurringTaskResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecRecurringTaskResponse(val *ExecRecurringTaskResponse) *NullableExecRecurringTaskResponse {
	return &NullableExecRecurringTaskResponse{value: val, isSet: true}
}

func (v NullableExecRecurringTaskResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecRecurringTaskResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
