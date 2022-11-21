# AddCollectSupportDataRecurringTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskName** | **string** | Name of the new Recurring Task | 
**Schemas** | [**[]EnumcollectSupportDataRecurringTaskSchemaUrn**](EnumcollectSupportDataRecurringTaskSchemaUrn.md) |  | 
**OutputDirectory** | **string** | The directory in which the support data archive files will be placed. The path must be a directory, and that directory must already exist. Relative paths will be interpreted as relative to the server root. | 
**EncryptionPassphraseFile** | Pointer to **string** | The path to a file that contains the passphrase to encrypt the contents of the support data archive. | [optional] 
**IncludeExpensiveData** | Pointer to **bool** | Indicates whether the support data archive should include information that may be expensive to obtain, and that may temporarily affect the server&#39;s performance or responsiveness. | [optional] 
**IncludeReplicationStateDump** | Pointer to **bool** | Indicates whether the support data archive should include a replication state dump, which may be several megabytes in size. | [optional] 
**IncludeBinaryFiles** | Pointer to **bool** | Indicates whether the support data archive should include binary files that may not have otherwise been included. Note that it may not be possible to obscure or redact sensitive information in binary files. | [optional] 
**IncludeExtensionSource** | Pointer to **bool** | Indicates whether the support data archive should include the source code (if available) for any third-party extensions that may be installed in the server. | [optional] 
**UseSequentialMode** | Pointer to **bool** | Indicates whether to capture support data information sequentially rather than in parallel. Capturing data in sequential mode may reduce the amount of memory that the tool requires to operate, at the cost of taking longer to run. | [optional] 
**SecurityLevel** | Pointer to [**EnumrecurringTaskSecurityLevelProp**](EnumrecurringTaskSecurityLevelProp.md) |  | [optional] 
**JstackCount** | Pointer to **int32** | The number of times to invoke the jstack utility to obtain a stack trace of all threads running in the JVM. A value of zero indicates that the jstack utility should not be invoked. | [optional] 
**ReportCount** | Pointer to **int32** | The number of intervals of data to collect from tools that use sample-based reporting, like vmstat, iostat, and mpstat. A value of zero indicates that these kinds of tools should not be used to collect any information. | [optional] 
**ReportIntervalSeconds** | Pointer to **int32** | The duration (in seconds) between each interval of data to collect from tools that use sample-based reporting, like vmstat, iostat, and mpstat. | [optional] 
**LogDuration** | Pointer to **string** | The maximum age (leading up to the time the collect-support-data tool was invoked) for log content to include in the support data archive. | [optional] 
**LogFileHeadCollectionSize** | Pointer to **string** | The amount of data to collect from the beginning of each log file included in the support data archive. | [optional] 
**LogFileTailCollectionSize** | Pointer to **string** | The amount of data to collect from the end of each log file included in the support data archive. | [optional] 
**Comment** | Pointer to **string** | An optional comment to include in a README file within the support data archive. | [optional] 
**RetainPreviousSupportDataArchiveCount** | Pointer to **int32** | The minimum number of previous support data archives that should be preserved after a new archive is generated. | [optional] 
**RetainPreviousSupportDataArchiveAge** | Pointer to **string** | The minimum age of previous support data archives that should be preserved after a new archive is generated. | [optional] 
**Description** | Pointer to **string** | A description for this Recurring Task | [optional] 
**CancelOnTaskDependencyFailure** | Pointer to **bool** | Indicates whether an instance of this Recurring Task should be canceled if the task immediately before it in the recurring task chain fails to complete successfully (including if it is canceled by an administrator before it starts or while it is running). | [optional] 
**EmailOnStart** | Pointer to **[]string** | The email addresses to which a message should be sent whenever an instance of this Recurring Task starts running. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**EmailOnSuccess** | Pointer to **[]string** | The email addresses to which a message should be sent whenever an instance of this Recurring Task completes successfully. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**EmailOnFailure** | Pointer to **[]string** | The email addresses to which a message should be sent if an instance of this Recurring Task fails to complete successfully. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**AlertOnStart** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task starts running. | [optional] 
**AlertOnSuccess** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task completes successfully. | [optional] 
**AlertOnFailure** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task fails to complete successfully. | [optional] 

## Methods

### NewAddCollectSupportDataRecurringTaskRequest

`func NewAddCollectSupportDataRecurringTaskRequest(taskName string, schemas []EnumcollectSupportDataRecurringTaskSchemaUrn, outputDirectory string, ) *AddCollectSupportDataRecurringTaskRequest`

NewAddCollectSupportDataRecurringTaskRequest instantiates a new AddCollectSupportDataRecurringTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCollectSupportDataRecurringTaskRequestWithDefaults

`func NewAddCollectSupportDataRecurringTaskRequestWithDefaults() *AddCollectSupportDataRecurringTaskRequest`

NewAddCollectSupportDataRecurringTaskRequestWithDefaults instantiates a new AddCollectSupportDataRecurringTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskName

`func (o *AddCollectSupportDataRecurringTaskRequest) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *AddCollectSupportDataRecurringTaskRequest) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *AddCollectSupportDataRecurringTaskRequest) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.


### GetSchemas

`func (o *AddCollectSupportDataRecurringTaskRequest) GetSchemas() []EnumcollectSupportDataRecurringTaskSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddCollectSupportDataRecurringTaskRequest) GetSchemasOk() (*[]EnumcollectSupportDataRecurringTaskSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddCollectSupportDataRecurringTaskRequest) SetSchemas(v []EnumcollectSupportDataRecurringTaskSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetOutputDirectory

`func (o *AddCollectSupportDataRecurringTaskRequest) GetOutputDirectory() string`

GetOutputDirectory returns the OutputDirectory field if non-nil, zero value otherwise.

### GetOutputDirectoryOk

`func (o *AddCollectSupportDataRecurringTaskRequest) GetOutputDirectoryOk() (*string, bool)`

GetOutputDirectoryOk returns a tuple with the OutputDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDirectory

`func (o *AddCollectSupportDataRecurringTaskRequest) SetOutputDirectory(v string)`

SetOutputDirectory sets OutputDirectory field to given value.


### GetEncryptionPassphraseFile

`func (o *AddCollectSupportDataRecurringTaskRequest) GetEncryptionPassphraseFile() string`

GetEncryptionPassphraseFile returns the EncryptionPassphraseFile field if non-nil, zero value otherwise.

### GetEncryptionPassphraseFileOk

`func (o *AddCollectSupportDataRecurringTaskRequest) GetEncryptionPassphraseFileOk() (*string, bool)`

GetEncryptionPassphraseFileOk returns a tuple with the EncryptionPassphraseFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionPassphraseFile

`func (o *AddCollectSupportDataRecurringTaskRequest) SetEncryptionPassphraseFile(v string)`

SetEncryptionPassphraseFile sets EncryptionPassphraseFile field to given value.

### HasEncryptionPassphraseFile

`func (o *AddCollectSupportDataRecurringTaskRequest) HasEncryptionPassphraseFile() bool`

HasEncryptionPassphraseFile returns a boolean if a field has been set.

### GetIncludeExpensiveData

`func (o *AddCollectSupportDataRecurringTaskRequest) GetIncludeExpensiveData() bool`

GetIncludeExpensiveData returns the IncludeExpensiveData field if non-nil, zero value otherwise.

### GetIncludeExpensiveDataOk

`func (o *AddCollectSupportDataRecurringTaskRequest) GetIncludeExpensiveDataOk() (*bool, bool)`

GetIncludeExpensiveDataOk returns a tuple with the IncludeExpensiveData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeExpensiveData

`func (o *AddCollectSupportDataRecurringTaskRequest) SetIncludeExpensiveData(v bool)`

SetIncludeExpensiveData sets IncludeExpensiveData field to given value.

### HasIncludeExpensiveData

`func (o *AddCollectSupportDataRecurringTaskRequest) HasIncludeExpensiveData() bool`

HasIncludeExpensiveData returns a boolean if a field has been set.

### GetIncludeReplicationStateDump

`func (o *AddCollectSupportDataRecurringTaskRequest) GetIncludeReplicationStateDump() bool`

GetIncludeReplicationStateDump returns the IncludeReplicationStateDump field if non-nil, zero value otherwise.

### GetIncludeReplicationStateDumpOk

`func (o *AddCollectSupportDataRecurringTaskRequest) GetIncludeReplicationStateDumpOk() (*bool, bool)`

GetIncludeReplicationStateDumpOk returns a tuple with the IncludeReplicationStateDump field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeReplicationStateDump

`func (o *AddCollectSupportDataRecurringTaskRequest) SetIncludeReplicationStateDump(v bool)`

SetIncludeReplicationStateDump sets IncludeReplicationStateDump field to given value.

### HasIncludeReplicationStateDump

`func (o *AddCollectSupportDataRecurringTaskRequest) HasIncludeReplicationStateDump() bool`

HasIncludeReplicationStateDump returns a boolean if a field has been set.

### GetIncludeBinaryFiles

`func (o *AddCollectSupportDataRecurringTaskRequest) GetIncludeBinaryFiles() bool`

GetIncludeBinaryFiles returns the IncludeBinaryFiles field if non-nil, zero value otherwise.

### GetIncludeBinaryFilesOk

`func (o *AddCollectSupportDataRecurringTaskRequest) GetIncludeBinaryFilesOk() (*bool, bool)`

GetIncludeBinaryFilesOk returns a tuple with the IncludeBinaryFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBinaryFiles

`func (o *AddCollectSupportDataRecurringTaskRequest) SetIncludeBinaryFiles(v bool)`

SetIncludeBinaryFiles sets IncludeBinaryFiles field to given value.

### HasIncludeBinaryFiles

`func (o *AddCollectSupportDataRecurringTaskRequest) HasIncludeBinaryFiles() bool`

HasIncludeBinaryFiles returns a boolean if a field has been set.

### GetIncludeExtensionSource

`func (o *AddCollectSupportDataRecurringTaskRequest) GetIncludeExtensionSource() bool`

GetIncludeExtensionSource returns the IncludeExtensionSource field if non-nil, zero value otherwise.

### GetIncludeExtensionSourceOk

`func (o *AddCollectSupportDataRecurringTaskRequest) GetIncludeExtensionSourceOk() (*bool, bool)`

GetIncludeExtensionSourceOk returns a tuple with the IncludeExtensionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeExtensionSource

`func (o *AddCollectSupportDataRecurringTaskRequest) SetIncludeExtensionSource(v bool)`

SetIncludeExtensionSource sets IncludeExtensionSource field to given value.

### HasIncludeExtensionSource

`func (o *AddCollectSupportDataRecurringTaskRequest) HasIncludeExtensionSource() bool`

HasIncludeExtensionSource returns a boolean if a field has been set.

### GetUseSequentialMode

`func (o *AddCollectSupportDataRecurringTaskRequest) GetUseSequentialMode() bool`

GetUseSequentialMode returns the UseSequentialMode field if non-nil, zero value otherwise.

### GetUseSequentialModeOk

`func (o *AddCollectSupportDataRecurringTaskRequest) GetUseSequentialModeOk() (*bool, bool)`

GetUseSequentialModeOk returns a tuple with the UseSequentialMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSequentialMode

`func (o *AddCollectSupportDataRecurringTaskRequest) SetUseSequentialMode(v bool)`

SetUseSequentialMode sets UseSequentialMode field to given value.

### HasUseSequentialMode

`func (o *AddCollectSupportDataRecurringTaskRequest) HasUseSequentialMode() bool`

HasUseSequentialMode returns a boolean if a field has been set.

### GetSecurityLevel

`func (o *AddCollectSupportDataRecurringTaskRequest) GetSecurityLevel() EnumrecurringTaskSecurityLevelProp`

GetSecurityLevel returns the SecurityLevel field if non-nil, zero value otherwise.

### GetSecurityLevelOk

`func (o *AddCollectSupportDataRecurringTaskRequest) GetSecurityLevelOk() (*EnumrecurringTaskSecurityLevelProp, bool)`

GetSecurityLevelOk returns a tuple with the SecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLevel

`func (o *AddCollectSupportDataRecurringTaskRequest) SetSecurityLevel(v EnumrecurringTaskSecurityLevelProp)`

SetSecurityLevel sets SecurityLevel field to given value.

### HasSecurityLevel

`func (o *AddCollectSupportDataRecurringTaskRequest) HasSecurityLevel() bool`

HasSecurityLevel returns a boolean if a field has been set.

### GetJstackCount

`func (o *AddCollectSupportDataRecurringTaskRequest) GetJstackCount() int32`

GetJstackCount returns the JstackCount field if non-nil, zero value otherwise.

### GetJstackCountOk

`func (o *AddCollectSupportDataRecurringTaskRequest) GetJstackCountOk() (*int32, bool)`

GetJstackCountOk returns a tuple with the JstackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJstackCount

`func (o *AddCollectSupportDataRecurringTaskRequest) SetJstackCount(v int32)`

SetJstackCount sets JstackCount field to given value.

### HasJstackCount

`func (o *AddCollectSupportDataRecurringTaskRequest) HasJstackCount() bool`

HasJstackCount returns a boolean if a field has been set.

### GetReportCount

`func (o *AddCollectSupportDataRecurringTaskRequest) GetReportCount() int32`

GetReportCount returns the ReportCount field if non-nil, zero value otherwise.

### GetReportCountOk

`func (o *AddCollectSupportDataRecurringTaskRequest) GetReportCountOk() (*int32, bool)`

GetReportCountOk returns a tuple with the ReportCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportCount

`func (o *AddCollectSupportDataRecurringTaskRequest) SetReportCount(v int32)`

SetReportCount sets ReportCount field to given value.

### HasReportCount

`func (o *AddCollectSupportDataRecurringTaskRequest) HasReportCount() bool`

HasReportCount returns a boolean if a field has been set.

### GetReportIntervalSeconds

`func (o *AddCollectSupportDataRecurringTaskRequest) GetReportIntervalSeconds() int32`

GetReportIntervalSeconds returns the ReportIntervalSeconds field if non-nil, zero value otherwise.

### GetReportIntervalSecondsOk

`func (o *AddCollectSupportDataRecurringTaskRequest) GetReportIntervalSecondsOk() (*int32, bool)`

GetReportIntervalSecondsOk returns a tuple with the ReportIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportIntervalSeconds

`func (o *AddCollectSupportDataRecurringTaskRequest) SetReportIntervalSeconds(v int32)`

SetReportIntervalSeconds sets ReportIntervalSeconds field to given value.

### HasReportIntervalSeconds

`func (o *AddCollectSupportDataRecurringTaskRequest) HasReportIntervalSeconds() bool`

HasReportIntervalSeconds returns a boolean if a field has been set.

### GetLogDuration

`func (o *AddCollectSupportDataRecurringTaskRequest) GetLogDuration() string`

GetLogDuration returns the LogDuration field if non-nil, zero value otherwise.

### GetLogDurationOk

`func (o *AddCollectSupportDataRecurringTaskRequest) GetLogDurationOk() (*string, bool)`

GetLogDurationOk returns a tuple with the LogDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDuration

`func (o *AddCollectSupportDataRecurringTaskRequest) SetLogDuration(v string)`

SetLogDuration sets LogDuration field to given value.

### HasLogDuration

`func (o *AddCollectSupportDataRecurringTaskRequest) HasLogDuration() bool`

HasLogDuration returns a boolean if a field has been set.

### GetLogFileHeadCollectionSize

`func (o *AddCollectSupportDataRecurringTaskRequest) GetLogFileHeadCollectionSize() string`

GetLogFileHeadCollectionSize returns the LogFileHeadCollectionSize field if non-nil, zero value otherwise.

### GetLogFileHeadCollectionSizeOk

`func (o *AddCollectSupportDataRecurringTaskRequest) GetLogFileHeadCollectionSizeOk() (*string, bool)`

GetLogFileHeadCollectionSizeOk returns a tuple with the LogFileHeadCollectionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFileHeadCollectionSize

`func (o *AddCollectSupportDataRecurringTaskRequest) SetLogFileHeadCollectionSize(v string)`

SetLogFileHeadCollectionSize sets LogFileHeadCollectionSize field to given value.

### HasLogFileHeadCollectionSize

`func (o *AddCollectSupportDataRecurringTaskRequest) HasLogFileHeadCollectionSize() bool`

HasLogFileHeadCollectionSize returns a boolean if a field has been set.

### GetLogFileTailCollectionSize

`func (o *AddCollectSupportDataRecurringTaskRequest) GetLogFileTailCollectionSize() string`

GetLogFileTailCollectionSize returns the LogFileTailCollectionSize field if non-nil, zero value otherwise.

### GetLogFileTailCollectionSizeOk

`func (o *AddCollectSupportDataRecurringTaskRequest) GetLogFileTailCollectionSizeOk() (*string, bool)`

GetLogFileTailCollectionSizeOk returns a tuple with the LogFileTailCollectionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFileTailCollectionSize

`func (o *AddCollectSupportDataRecurringTaskRequest) SetLogFileTailCollectionSize(v string)`

SetLogFileTailCollectionSize sets LogFileTailCollectionSize field to given value.

### HasLogFileTailCollectionSize

`func (o *AddCollectSupportDataRecurringTaskRequest) HasLogFileTailCollectionSize() bool`

HasLogFileTailCollectionSize returns a boolean if a field has been set.

### GetComment

`func (o *AddCollectSupportDataRecurringTaskRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *AddCollectSupportDataRecurringTaskRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *AddCollectSupportDataRecurringTaskRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *AddCollectSupportDataRecurringTaskRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetRetainPreviousSupportDataArchiveCount

`func (o *AddCollectSupportDataRecurringTaskRequest) GetRetainPreviousSupportDataArchiveCount() int32`

GetRetainPreviousSupportDataArchiveCount returns the RetainPreviousSupportDataArchiveCount field if non-nil, zero value otherwise.

### GetRetainPreviousSupportDataArchiveCountOk

`func (o *AddCollectSupportDataRecurringTaskRequest) GetRetainPreviousSupportDataArchiveCountOk() (*int32, bool)`

GetRetainPreviousSupportDataArchiveCountOk returns a tuple with the RetainPreviousSupportDataArchiveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousSupportDataArchiveCount

`func (o *AddCollectSupportDataRecurringTaskRequest) SetRetainPreviousSupportDataArchiveCount(v int32)`

SetRetainPreviousSupportDataArchiveCount sets RetainPreviousSupportDataArchiveCount field to given value.

### HasRetainPreviousSupportDataArchiveCount

`func (o *AddCollectSupportDataRecurringTaskRequest) HasRetainPreviousSupportDataArchiveCount() bool`

HasRetainPreviousSupportDataArchiveCount returns a boolean if a field has been set.

### GetRetainPreviousSupportDataArchiveAge

`func (o *AddCollectSupportDataRecurringTaskRequest) GetRetainPreviousSupportDataArchiveAge() string`

GetRetainPreviousSupportDataArchiveAge returns the RetainPreviousSupportDataArchiveAge field if non-nil, zero value otherwise.

### GetRetainPreviousSupportDataArchiveAgeOk

`func (o *AddCollectSupportDataRecurringTaskRequest) GetRetainPreviousSupportDataArchiveAgeOk() (*string, bool)`

GetRetainPreviousSupportDataArchiveAgeOk returns a tuple with the RetainPreviousSupportDataArchiveAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousSupportDataArchiveAge

`func (o *AddCollectSupportDataRecurringTaskRequest) SetRetainPreviousSupportDataArchiveAge(v string)`

SetRetainPreviousSupportDataArchiveAge sets RetainPreviousSupportDataArchiveAge field to given value.

### HasRetainPreviousSupportDataArchiveAge

`func (o *AddCollectSupportDataRecurringTaskRequest) HasRetainPreviousSupportDataArchiveAge() bool`

HasRetainPreviousSupportDataArchiveAge returns a boolean if a field has been set.

### GetDescription

`func (o *AddCollectSupportDataRecurringTaskRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddCollectSupportDataRecurringTaskRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddCollectSupportDataRecurringTaskRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddCollectSupportDataRecurringTaskRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCancelOnTaskDependencyFailure

`func (o *AddCollectSupportDataRecurringTaskRequest) GetCancelOnTaskDependencyFailure() bool`

GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field if non-nil, zero value otherwise.

### GetCancelOnTaskDependencyFailureOk

`func (o *AddCollectSupportDataRecurringTaskRequest) GetCancelOnTaskDependencyFailureOk() (*bool, bool)`

GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOnTaskDependencyFailure

`func (o *AddCollectSupportDataRecurringTaskRequest) SetCancelOnTaskDependencyFailure(v bool)`

SetCancelOnTaskDependencyFailure sets CancelOnTaskDependencyFailure field to given value.

### HasCancelOnTaskDependencyFailure

`func (o *AddCollectSupportDataRecurringTaskRequest) HasCancelOnTaskDependencyFailure() bool`

HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.

### GetEmailOnStart

`func (o *AddCollectSupportDataRecurringTaskRequest) GetEmailOnStart() []string`

GetEmailOnStart returns the EmailOnStart field if non-nil, zero value otherwise.

### GetEmailOnStartOk

`func (o *AddCollectSupportDataRecurringTaskRequest) GetEmailOnStartOk() (*[]string, bool)`

GetEmailOnStartOk returns a tuple with the EmailOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnStart

`func (o *AddCollectSupportDataRecurringTaskRequest) SetEmailOnStart(v []string)`

SetEmailOnStart sets EmailOnStart field to given value.

### HasEmailOnStart

`func (o *AddCollectSupportDataRecurringTaskRequest) HasEmailOnStart() bool`

HasEmailOnStart returns a boolean if a field has been set.

### GetEmailOnSuccess

`func (o *AddCollectSupportDataRecurringTaskRequest) GetEmailOnSuccess() []string`

GetEmailOnSuccess returns the EmailOnSuccess field if non-nil, zero value otherwise.

### GetEmailOnSuccessOk

`func (o *AddCollectSupportDataRecurringTaskRequest) GetEmailOnSuccessOk() (*[]string, bool)`

GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnSuccess

`func (o *AddCollectSupportDataRecurringTaskRequest) SetEmailOnSuccess(v []string)`

SetEmailOnSuccess sets EmailOnSuccess field to given value.

### HasEmailOnSuccess

`func (o *AddCollectSupportDataRecurringTaskRequest) HasEmailOnSuccess() bool`

HasEmailOnSuccess returns a boolean if a field has been set.

### GetEmailOnFailure

`func (o *AddCollectSupportDataRecurringTaskRequest) GetEmailOnFailure() []string`

GetEmailOnFailure returns the EmailOnFailure field if non-nil, zero value otherwise.

### GetEmailOnFailureOk

`func (o *AddCollectSupportDataRecurringTaskRequest) GetEmailOnFailureOk() (*[]string, bool)`

GetEmailOnFailureOk returns a tuple with the EmailOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnFailure

`func (o *AddCollectSupportDataRecurringTaskRequest) SetEmailOnFailure(v []string)`

SetEmailOnFailure sets EmailOnFailure field to given value.

### HasEmailOnFailure

`func (o *AddCollectSupportDataRecurringTaskRequest) HasEmailOnFailure() bool`

HasEmailOnFailure returns a boolean if a field has been set.

### GetAlertOnStart

`func (o *AddCollectSupportDataRecurringTaskRequest) GetAlertOnStart() bool`

GetAlertOnStart returns the AlertOnStart field if non-nil, zero value otherwise.

### GetAlertOnStartOk

`func (o *AddCollectSupportDataRecurringTaskRequest) GetAlertOnStartOk() (*bool, bool)`

GetAlertOnStartOk returns a tuple with the AlertOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnStart

`func (o *AddCollectSupportDataRecurringTaskRequest) SetAlertOnStart(v bool)`

SetAlertOnStart sets AlertOnStart field to given value.

### HasAlertOnStart

`func (o *AddCollectSupportDataRecurringTaskRequest) HasAlertOnStart() bool`

HasAlertOnStart returns a boolean if a field has been set.

### GetAlertOnSuccess

`func (o *AddCollectSupportDataRecurringTaskRequest) GetAlertOnSuccess() bool`

GetAlertOnSuccess returns the AlertOnSuccess field if non-nil, zero value otherwise.

### GetAlertOnSuccessOk

`func (o *AddCollectSupportDataRecurringTaskRequest) GetAlertOnSuccessOk() (*bool, bool)`

GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnSuccess

`func (o *AddCollectSupportDataRecurringTaskRequest) SetAlertOnSuccess(v bool)`

SetAlertOnSuccess sets AlertOnSuccess field to given value.

### HasAlertOnSuccess

`func (o *AddCollectSupportDataRecurringTaskRequest) HasAlertOnSuccess() bool`

HasAlertOnSuccess returns a boolean if a field has been set.

### GetAlertOnFailure

`func (o *AddCollectSupportDataRecurringTaskRequest) GetAlertOnFailure() bool`

GetAlertOnFailure returns the AlertOnFailure field if non-nil, zero value otherwise.

### GetAlertOnFailureOk

`func (o *AddCollectSupportDataRecurringTaskRequest) GetAlertOnFailureOk() (*bool, bool)`

GetAlertOnFailureOk returns a tuple with the AlertOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFailure

`func (o *AddCollectSupportDataRecurringTaskRequest) SetAlertOnFailure(v bool)`

SetAlertOnFailure sets AlertOnFailure field to given value.

### HasAlertOnFailure

`func (o *AddCollectSupportDataRecurringTaskRequest) HasAlertOnFailure() bool`

HasAlertOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


