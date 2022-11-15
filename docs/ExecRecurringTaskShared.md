# ExecRecurringTaskShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumexecRecurringTaskSchemaUrn**](EnumexecRecurringTaskSchemaUrn.md) |  | 
**CommandPath** | **string** | The absolute path to the command to execute. It must be an absolute path, the corresponding file must exist, and it must be listed in the config/exec-command-whitelist.txt file. | 
**CommandArguments** | Pointer to **string** | A string containing the arguments to provide to the command. If the command should be run without arguments, this property should be left undefined. If there should be multiple arguments, then they should be separated with spaces. | [optional] 
**CommandOutputFileBaseName** | Pointer to **string** | The path and base name for a file to which the command output (both standard output and standard error) should be written. This may be left undefined if the command output should not be recorded into a file. | [optional] 
**RetainPreviousOutputFileCount** | Pointer to **int32** | The minimum number of previous command output files that should be preserved after a new instance of the command is invoked. | [optional] 
**RetainPreviousOutputFileAge** | Pointer to **string** | The minimum age of previous command output files that should be preserved after a new instance of the command is invoked. | [optional] 
**LogCommandOutput** | Pointer to **bool** | Indicates whether the command&#39;s output (both standard output and standard error) should be recorded in the server&#39;s error log. | [optional] 
**TaskCompletionStateForNonzeroExitCode** | Pointer to [**EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp**](EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp.md) |  | [optional] 
**WorkingDirectory** | Pointer to **string** | The absolute path to a working directory where the command should be executed. It must be an absolute path and the corresponding directory must exist. | [optional] 
**Description** | Pointer to **string** | A description for this Recurring Task | [optional] 
**CancelOnTaskDependencyFailure** | Pointer to **bool** | Indicates whether an instance of this Recurring Task should be canceled if the task immediately before it in the recurring task chain fails to complete successfully (including if it is canceled by an administrator before it starts or while it is running). | [optional] 
**EmailOnStart** | Pointer to **[]string** |  | [optional] 
**EmailOnSuccess** | Pointer to **[]string** |  | [optional] 
**EmailOnFailure** | Pointer to **[]string** |  | [optional] 
**AlertOnStart** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task starts running. | [optional] 
**AlertOnSuccess** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task completes successfully. | [optional] 
**AlertOnFailure** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task fails to complete successfully. | [optional] 

## Methods

### NewExecRecurringTaskShared

`func NewExecRecurringTaskShared(schemas []EnumexecRecurringTaskSchemaUrn, commandPath string, ) *ExecRecurringTaskShared`

NewExecRecurringTaskShared instantiates a new ExecRecurringTaskShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecRecurringTaskSharedWithDefaults

`func NewExecRecurringTaskSharedWithDefaults() *ExecRecurringTaskShared`

NewExecRecurringTaskSharedWithDefaults instantiates a new ExecRecurringTaskShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ExecRecurringTaskShared) GetSchemas() []EnumexecRecurringTaskSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ExecRecurringTaskShared) GetSchemasOk() (*[]EnumexecRecurringTaskSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ExecRecurringTaskShared) SetSchemas(v []EnumexecRecurringTaskSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetCommandPath

`func (o *ExecRecurringTaskShared) GetCommandPath() string`

GetCommandPath returns the CommandPath field if non-nil, zero value otherwise.

### GetCommandPathOk

`func (o *ExecRecurringTaskShared) GetCommandPathOk() (*string, bool)`

GetCommandPathOk returns a tuple with the CommandPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandPath

`func (o *ExecRecurringTaskShared) SetCommandPath(v string)`

SetCommandPath sets CommandPath field to given value.


### GetCommandArguments

`func (o *ExecRecurringTaskShared) GetCommandArguments() string`

GetCommandArguments returns the CommandArguments field if non-nil, zero value otherwise.

### GetCommandArgumentsOk

`func (o *ExecRecurringTaskShared) GetCommandArgumentsOk() (*string, bool)`

GetCommandArgumentsOk returns a tuple with the CommandArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandArguments

`func (o *ExecRecurringTaskShared) SetCommandArguments(v string)`

SetCommandArguments sets CommandArguments field to given value.

### HasCommandArguments

`func (o *ExecRecurringTaskShared) HasCommandArguments() bool`

HasCommandArguments returns a boolean if a field has been set.

### GetCommandOutputFileBaseName

`func (o *ExecRecurringTaskShared) GetCommandOutputFileBaseName() string`

GetCommandOutputFileBaseName returns the CommandOutputFileBaseName field if non-nil, zero value otherwise.

### GetCommandOutputFileBaseNameOk

`func (o *ExecRecurringTaskShared) GetCommandOutputFileBaseNameOk() (*string, bool)`

GetCommandOutputFileBaseNameOk returns a tuple with the CommandOutputFileBaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandOutputFileBaseName

`func (o *ExecRecurringTaskShared) SetCommandOutputFileBaseName(v string)`

SetCommandOutputFileBaseName sets CommandOutputFileBaseName field to given value.

### HasCommandOutputFileBaseName

`func (o *ExecRecurringTaskShared) HasCommandOutputFileBaseName() bool`

HasCommandOutputFileBaseName returns a boolean if a field has been set.

### GetRetainPreviousOutputFileCount

`func (o *ExecRecurringTaskShared) GetRetainPreviousOutputFileCount() int32`

GetRetainPreviousOutputFileCount returns the RetainPreviousOutputFileCount field if non-nil, zero value otherwise.

### GetRetainPreviousOutputFileCountOk

`func (o *ExecRecurringTaskShared) GetRetainPreviousOutputFileCountOk() (*int32, bool)`

GetRetainPreviousOutputFileCountOk returns a tuple with the RetainPreviousOutputFileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousOutputFileCount

`func (o *ExecRecurringTaskShared) SetRetainPreviousOutputFileCount(v int32)`

SetRetainPreviousOutputFileCount sets RetainPreviousOutputFileCount field to given value.

### HasRetainPreviousOutputFileCount

`func (o *ExecRecurringTaskShared) HasRetainPreviousOutputFileCount() bool`

HasRetainPreviousOutputFileCount returns a boolean if a field has been set.

### GetRetainPreviousOutputFileAge

`func (o *ExecRecurringTaskShared) GetRetainPreviousOutputFileAge() string`

GetRetainPreviousOutputFileAge returns the RetainPreviousOutputFileAge field if non-nil, zero value otherwise.

### GetRetainPreviousOutputFileAgeOk

`func (o *ExecRecurringTaskShared) GetRetainPreviousOutputFileAgeOk() (*string, bool)`

GetRetainPreviousOutputFileAgeOk returns a tuple with the RetainPreviousOutputFileAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousOutputFileAge

`func (o *ExecRecurringTaskShared) SetRetainPreviousOutputFileAge(v string)`

SetRetainPreviousOutputFileAge sets RetainPreviousOutputFileAge field to given value.

### HasRetainPreviousOutputFileAge

`func (o *ExecRecurringTaskShared) HasRetainPreviousOutputFileAge() bool`

HasRetainPreviousOutputFileAge returns a boolean if a field has been set.

### GetLogCommandOutput

`func (o *ExecRecurringTaskShared) GetLogCommandOutput() bool`

GetLogCommandOutput returns the LogCommandOutput field if non-nil, zero value otherwise.

### GetLogCommandOutputOk

`func (o *ExecRecurringTaskShared) GetLogCommandOutputOk() (*bool, bool)`

GetLogCommandOutputOk returns a tuple with the LogCommandOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogCommandOutput

`func (o *ExecRecurringTaskShared) SetLogCommandOutput(v bool)`

SetLogCommandOutput sets LogCommandOutput field to given value.

### HasLogCommandOutput

`func (o *ExecRecurringTaskShared) HasLogCommandOutput() bool`

HasLogCommandOutput returns a boolean if a field has been set.

### GetTaskCompletionStateForNonzeroExitCode

`func (o *ExecRecurringTaskShared) GetTaskCompletionStateForNonzeroExitCode() EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp`

GetTaskCompletionStateForNonzeroExitCode returns the TaskCompletionStateForNonzeroExitCode field if non-nil, zero value otherwise.

### GetTaskCompletionStateForNonzeroExitCodeOk

`func (o *ExecRecurringTaskShared) GetTaskCompletionStateForNonzeroExitCodeOk() (*EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp, bool)`

GetTaskCompletionStateForNonzeroExitCodeOk returns a tuple with the TaskCompletionStateForNonzeroExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCompletionStateForNonzeroExitCode

`func (o *ExecRecurringTaskShared) SetTaskCompletionStateForNonzeroExitCode(v EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp)`

SetTaskCompletionStateForNonzeroExitCode sets TaskCompletionStateForNonzeroExitCode field to given value.

### HasTaskCompletionStateForNonzeroExitCode

`func (o *ExecRecurringTaskShared) HasTaskCompletionStateForNonzeroExitCode() bool`

HasTaskCompletionStateForNonzeroExitCode returns a boolean if a field has been set.

### GetWorkingDirectory

`func (o *ExecRecurringTaskShared) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *ExecRecurringTaskShared) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *ExecRecurringTaskShared) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.

### HasWorkingDirectory

`func (o *ExecRecurringTaskShared) HasWorkingDirectory() bool`

HasWorkingDirectory returns a boolean if a field has been set.

### GetDescription

`func (o *ExecRecurringTaskShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExecRecurringTaskShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExecRecurringTaskShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExecRecurringTaskShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCancelOnTaskDependencyFailure

`func (o *ExecRecurringTaskShared) GetCancelOnTaskDependencyFailure() bool`

GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field if non-nil, zero value otherwise.

### GetCancelOnTaskDependencyFailureOk

`func (o *ExecRecurringTaskShared) GetCancelOnTaskDependencyFailureOk() (*bool, bool)`

GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOnTaskDependencyFailure

`func (o *ExecRecurringTaskShared) SetCancelOnTaskDependencyFailure(v bool)`

SetCancelOnTaskDependencyFailure sets CancelOnTaskDependencyFailure field to given value.

### HasCancelOnTaskDependencyFailure

`func (o *ExecRecurringTaskShared) HasCancelOnTaskDependencyFailure() bool`

HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.

### GetEmailOnStart

`func (o *ExecRecurringTaskShared) GetEmailOnStart() []string`

GetEmailOnStart returns the EmailOnStart field if non-nil, zero value otherwise.

### GetEmailOnStartOk

`func (o *ExecRecurringTaskShared) GetEmailOnStartOk() (*[]string, bool)`

GetEmailOnStartOk returns a tuple with the EmailOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnStart

`func (o *ExecRecurringTaskShared) SetEmailOnStart(v []string)`

SetEmailOnStart sets EmailOnStart field to given value.

### HasEmailOnStart

`func (o *ExecRecurringTaskShared) HasEmailOnStart() bool`

HasEmailOnStart returns a boolean if a field has been set.

### GetEmailOnSuccess

`func (o *ExecRecurringTaskShared) GetEmailOnSuccess() []string`

GetEmailOnSuccess returns the EmailOnSuccess field if non-nil, zero value otherwise.

### GetEmailOnSuccessOk

`func (o *ExecRecurringTaskShared) GetEmailOnSuccessOk() (*[]string, bool)`

GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnSuccess

`func (o *ExecRecurringTaskShared) SetEmailOnSuccess(v []string)`

SetEmailOnSuccess sets EmailOnSuccess field to given value.

### HasEmailOnSuccess

`func (o *ExecRecurringTaskShared) HasEmailOnSuccess() bool`

HasEmailOnSuccess returns a boolean if a field has been set.

### GetEmailOnFailure

`func (o *ExecRecurringTaskShared) GetEmailOnFailure() []string`

GetEmailOnFailure returns the EmailOnFailure field if non-nil, zero value otherwise.

### GetEmailOnFailureOk

`func (o *ExecRecurringTaskShared) GetEmailOnFailureOk() (*[]string, bool)`

GetEmailOnFailureOk returns a tuple with the EmailOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnFailure

`func (o *ExecRecurringTaskShared) SetEmailOnFailure(v []string)`

SetEmailOnFailure sets EmailOnFailure field to given value.

### HasEmailOnFailure

`func (o *ExecRecurringTaskShared) HasEmailOnFailure() bool`

HasEmailOnFailure returns a boolean if a field has been set.

### GetAlertOnStart

`func (o *ExecRecurringTaskShared) GetAlertOnStart() bool`

GetAlertOnStart returns the AlertOnStart field if non-nil, zero value otherwise.

### GetAlertOnStartOk

`func (o *ExecRecurringTaskShared) GetAlertOnStartOk() (*bool, bool)`

GetAlertOnStartOk returns a tuple with the AlertOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnStart

`func (o *ExecRecurringTaskShared) SetAlertOnStart(v bool)`

SetAlertOnStart sets AlertOnStart field to given value.

### HasAlertOnStart

`func (o *ExecRecurringTaskShared) HasAlertOnStart() bool`

HasAlertOnStart returns a boolean if a field has been set.

### GetAlertOnSuccess

`func (o *ExecRecurringTaskShared) GetAlertOnSuccess() bool`

GetAlertOnSuccess returns the AlertOnSuccess field if non-nil, zero value otherwise.

### GetAlertOnSuccessOk

`func (o *ExecRecurringTaskShared) GetAlertOnSuccessOk() (*bool, bool)`

GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnSuccess

`func (o *ExecRecurringTaskShared) SetAlertOnSuccess(v bool)`

SetAlertOnSuccess sets AlertOnSuccess field to given value.

### HasAlertOnSuccess

`func (o *ExecRecurringTaskShared) HasAlertOnSuccess() bool`

HasAlertOnSuccess returns a boolean if a field has been set.

### GetAlertOnFailure

`func (o *ExecRecurringTaskShared) GetAlertOnFailure() bool`

GetAlertOnFailure returns the AlertOnFailure field if non-nil, zero value otherwise.

### GetAlertOnFailureOk

`func (o *ExecRecurringTaskShared) GetAlertOnFailureOk() (*bool, bool)`

GetAlertOnFailureOk returns a tuple with the AlertOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFailure

`func (o *ExecRecurringTaskShared) SetAlertOnFailure(v bool)`

SetAlertOnFailure sets AlertOnFailure field to given value.

### HasAlertOnFailure

`func (o *ExecRecurringTaskShared) HasAlertOnFailure() bool`

HasAlertOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


