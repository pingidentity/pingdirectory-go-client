# CollectSupportDataRecurringTaskShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**EmailOnStart** | Pointer to **[]string** |  | [optional] 
**EmailOnSuccess** | Pointer to **[]string** |  | [optional] 
**EmailOnFailure** | Pointer to **[]string** |  | [optional] 
**AlertOnStart** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task starts running. | [optional] 
**AlertOnSuccess** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task completes successfully. | [optional] 
**AlertOnFailure** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task fails to complete successfully. | [optional] 

## Methods

### NewCollectSupportDataRecurringTaskShared

`func NewCollectSupportDataRecurringTaskShared(schemas []EnumcollectSupportDataRecurringTaskSchemaUrn, outputDirectory string, ) *CollectSupportDataRecurringTaskShared`

NewCollectSupportDataRecurringTaskShared instantiates a new CollectSupportDataRecurringTaskShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectSupportDataRecurringTaskSharedWithDefaults

`func NewCollectSupportDataRecurringTaskSharedWithDefaults() *CollectSupportDataRecurringTaskShared`

NewCollectSupportDataRecurringTaskSharedWithDefaults instantiates a new CollectSupportDataRecurringTaskShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *CollectSupportDataRecurringTaskShared) GetSchemas() []EnumcollectSupportDataRecurringTaskSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *CollectSupportDataRecurringTaskShared) GetSchemasOk() (*[]EnumcollectSupportDataRecurringTaskSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *CollectSupportDataRecurringTaskShared) SetSchemas(v []EnumcollectSupportDataRecurringTaskSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetOutputDirectory

`func (o *CollectSupportDataRecurringTaskShared) GetOutputDirectory() string`

GetOutputDirectory returns the OutputDirectory field if non-nil, zero value otherwise.

### GetOutputDirectoryOk

`func (o *CollectSupportDataRecurringTaskShared) GetOutputDirectoryOk() (*string, bool)`

GetOutputDirectoryOk returns a tuple with the OutputDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDirectory

`func (o *CollectSupportDataRecurringTaskShared) SetOutputDirectory(v string)`

SetOutputDirectory sets OutputDirectory field to given value.


### GetEncryptionPassphraseFile

`func (o *CollectSupportDataRecurringTaskShared) GetEncryptionPassphraseFile() string`

GetEncryptionPassphraseFile returns the EncryptionPassphraseFile field if non-nil, zero value otherwise.

### GetEncryptionPassphraseFileOk

`func (o *CollectSupportDataRecurringTaskShared) GetEncryptionPassphraseFileOk() (*string, bool)`

GetEncryptionPassphraseFileOk returns a tuple with the EncryptionPassphraseFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionPassphraseFile

`func (o *CollectSupportDataRecurringTaskShared) SetEncryptionPassphraseFile(v string)`

SetEncryptionPassphraseFile sets EncryptionPassphraseFile field to given value.

### HasEncryptionPassphraseFile

`func (o *CollectSupportDataRecurringTaskShared) HasEncryptionPassphraseFile() bool`

HasEncryptionPassphraseFile returns a boolean if a field has been set.

### GetIncludeExpensiveData

`func (o *CollectSupportDataRecurringTaskShared) GetIncludeExpensiveData() bool`

GetIncludeExpensiveData returns the IncludeExpensiveData field if non-nil, zero value otherwise.

### GetIncludeExpensiveDataOk

`func (o *CollectSupportDataRecurringTaskShared) GetIncludeExpensiveDataOk() (*bool, bool)`

GetIncludeExpensiveDataOk returns a tuple with the IncludeExpensiveData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeExpensiveData

`func (o *CollectSupportDataRecurringTaskShared) SetIncludeExpensiveData(v bool)`

SetIncludeExpensiveData sets IncludeExpensiveData field to given value.

### HasIncludeExpensiveData

`func (o *CollectSupportDataRecurringTaskShared) HasIncludeExpensiveData() bool`

HasIncludeExpensiveData returns a boolean if a field has been set.

### GetIncludeReplicationStateDump

`func (o *CollectSupportDataRecurringTaskShared) GetIncludeReplicationStateDump() bool`

GetIncludeReplicationStateDump returns the IncludeReplicationStateDump field if non-nil, zero value otherwise.

### GetIncludeReplicationStateDumpOk

`func (o *CollectSupportDataRecurringTaskShared) GetIncludeReplicationStateDumpOk() (*bool, bool)`

GetIncludeReplicationStateDumpOk returns a tuple with the IncludeReplicationStateDump field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeReplicationStateDump

`func (o *CollectSupportDataRecurringTaskShared) SetIncludeReplicationStateDump(v bool)`

SetIncludeReplicationStateDump sets IncludeReplicationStateDump field to given value.

### HasIncludeReplicationStateDump

`func (o *CollectSupportDataRecurringTaskShared) HasIncludeReplicationStateDump() bool`

HasIncludeReplicationStateDump returns a boolean if a field has been set.

### GetIncludeBinaryFiles

`func (o *CollectSupportDataRecurringTaskShared) GetIncludeBinaryFiles() bool`

GetIncludeBinaryFiles returns the IncludeBinaryFiles field if non-nil, zero value otherwise.

### GetIncludeBinaryFilesOk

`func (o *CollectSupportDataRecurringTaskShared) GetIncludeBinaryFilesOk() (*bool, bool)`

GetIncludeBinaryFilesOk returns a tuple with the IncludeBinaryFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBinaryFiles

`func (o *CollectSupportDataRecurringTaskShared) SetIncludeBinaryFiles(v bool)`

SetIncludeBinaryFiles sets IncludeBinaryFiles field to given value.

### HasIncludeBinaryFiles

`func (o *CollectSupportDataRecurringTaskShared) HasIncludeBinaryFiles() bool`

HasIncludeBinaryFiles returns a boolean if a field has been set.

### GetIncludeExtensionSource

`func (o *CollectSupportDataRecurringTaskShared) GetIncludeExtensionSource() bool`

GetIncludeExtensionSource returns the IncludeExtensionSource field if non-nil, zero value otherwise.

### GetIncludeExtensionSourceOk

`func (o *CollectSupportDataRecurringTaskShared) GetIncludeExtensionSourceOk() (*bool, bool)`

GetIncludeExtensionSourceOk returns a tuple with the IncludeExtensionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeExtensionSource

`func (o *CollectSupportDataRecurringTaskShared) SetIncludeExtensionSource(v bool)`

SetIncludeExtensionSource sets IncludeExtensionSource field to given value.

### HasIncludeExtensionSource

`func (o *CollectSupportDataRecurringTaskShared) HasIncludeExtensionSource() bool`

HasIncludeExtensionSource returns a boolean if a field has been set.

### GetUseSequentialMode

`func (o *CollectSupportDataRecurringTaskShared) GetUseSequentialMode() bool`

GetUseSequentialMode returns the UseSequentialMode field if non-nil, zero value otherwise.

### GetUseSequentialModeOk

`func (o *CollectSupportDataRecurringTaskShared) GetUseSequentialModeOk() (*bool, bool)`

GetUseSequentialModeOk returns a tuple with the UseSequentialMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSequentialMode

`func (o *CollectSupportDataRecurringTaskShared) SetUseSequentialMode(v bool)`

SetUseSequentialMode sets UseSequentialMode field to given value.

### HasUseSequentialMode

`func (o *CollectSupportDataRecurringTaskShared) HasUseSequentialMode() bool`

HasUseSequentialMode returns a boolean if a field has been set.

### GetSecurityLevel

`func (o *CollectSupportDataRecurringTaskShared) GetSecurityLevel() EnumrecurringTaskSecurityLevelProp`

GetSecurityLevel returns the SecurityLevel field if non-nil, zero value otherwise.

### GetSecurityLevelOk

`func (o *CollectSupportDataRecurringTaskShared) GetSecurityLevelOk() (*EnumrecurringTaskSecurityLevelProp, bool)`

GetSecurityLevelOk returns a tuple with the SecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLevel

`func (o *CollectSupportDataRecurringTaskShared) SetSecurityLevel(v EnumrecurringTaskSecurityLevelProp)`

SetSecurityLevel sets SecurityLevel field to given value.

### HasSecurityLevel

`func (o *CollectSupportDataRecurringTaskShared) HasSecurityLevel() bool`

HasSecurityLevel returns a boolean if a field has been set.

### GetJstackCount

`func (o *CollectSupportDataRecurringTaskShared) GetJstackCount() int32`

GetJstackCount returns the JstackCount field if non-nil, zero value otherwise.

### GetJstackCountOk

`func (o *CollectSupportDataRecurringTaskShared) GetJstackCountOk() (*int32, bool)`

GetJstackCountOk returns a tuple with the JstackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJstackCount

`func (o *CollectSupportDataRecurringTaskShared) SetJstackCount(v int32)`

SetJstackCount sets JstackCount field to given value.

### HasJstackCount

`func (o *CollectSupportDataRecurringTaskShared) HasJstackCount() bool`

HasJstackCount returns a boolean if a field has been set.

### GetReportCount

`func (o *CollectSupportDataRecurringTaskShared) GetReportCount() int32`

GetReportCount returns the ReportCount field if non-nil, zero value otherwise.

### GetReportCountOk

`func (o *CollectSupportDataRecurringTaskShared) GetReportCountOk() (*int32, bool)`

GetReportCountOk returns a tuple with the ReportCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportCount

`func (o *CollectSupportDataRecurringTaskShared) SetReportCount(v int32)`

SetReportCount sets ReportCount field to given value.

### HasReportCount

`func (o *CollectSupportDataRecurringTaskShared) HasReportCount() bool`

HasReportCount returns a boolean if a field has been set.

### GetReportIntervalSeconds

`func (o *CollectSupportDataRecurringTaskShared) GetReportIntervalSeconds() int32`

GetReportIntervalSeconds returns the ReportIntervalSeconds field if non-nil, zero value otherwise.

### GetReportIntervalSecondsOk

`func (o *CollectSupportDataRecurringTaskShared) GetReportIntervalSecondsOk() (*int32, bool)`

GetReportIntervalSecondsOk returns a tuple with the ReportIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportIntervalSeconds

`func (o *CollectSupportDataRecurringTaskShared) SetReportIntervalSeconds(v int32)`

SetReportIntervalSeconds sets ReportIntervalSeconds field to given value.

### HasReportIntervalSeconds

`func (o *CollectSupportDataRecurringTaskShared) HasReportIntervalSeconds() bool`

HasReportIntervalSeconds returns a boolean if a field has been set.

### GetLogDuration

`func (o *CollectSupportDataRecurringTaskShared) GetLogDuration() string`

GetLogDuration returns the LogDuration field if non-nil, zero value otherwise.

### GetLogDurationOk

`func (o *CollectSupportDataRecurringTaskShared) GetLogDurationOk() (*string, bool)`

GetLogDurationOk returns a tuple with the LogDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDuration

`func (o *CollectSupportDataRecurringTaskShared) SetLogDuration(v string)`

SetLogDuration sets LogDuration field to given value.

### HasLogDuration

`func (o *CollectSupportDataRecurringTaskShared) HasLogDuration() bool`

HasLogDuration returns a boolean if a field has been set.

### GetLogFileHeadCollectionSize

`func (o *CollectSupportDataRecurringTaskShared) GetLogFileHeadCollectionSize() string`

GetLogFileHeadCollectionSize returns the LogFileHeadCollectionSize field if non-nil, zero value otherwise.

### GetLogFileHeadCollectionSizeOk

`func (o *CollectSupportDataRecurringTaskShared) GetLogFileHeadCollectionSizeOk() (*string, bool)`

GetLogFileHeadCollectionSizeOk returns a tuple with the LogFileHeadCollectionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFileHeadCollectionSize

`func (o *CollectSupportDataRecurringTaskShared) SetLogFileHeadCollectionSize(v string)`

SetLogFileHeadCollectionSize sets LogFileHeadCollectionSize field to given value.

### HasLogFileHeadCollectionSize

`func (o *CollectSupportDataRecurringTaskShared) HasLogFileHeadCollectionSize() bool`

HasLogFileHeadCollectionSize returns a boolean if a field has been set.

### GetLogFileTailCollectionSize

`func (o *CollectSupportDataRecurringTaskShared) GetLogFileTailCollectionSize() string`

GetLogFileTailCollectionSize returns the LogFileTailCollectionSize field if non-nil, zero value otherwise.

### GetLogFileTailCollectionSizeOk

`func (o *CollectSupportDataRecurringTaskShared) GetLogFileTailCollectionSizeOk() (*string, bool)`

GetLogFileTailCollectionSizeOk returns a tuple with the LogFileTailCollectionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFileTailCollectionSize

`func (o *CollectSupportDataRecurringTaskShared) SetLogFileTailCollectionSize(v string)`

SetLogFileTailCollectionSize sets LogFileTailCollectionSize field to given value.

### HasLogFileTailCollectionSize

`func (o *CollectSupportDataRecurringTaskShared) HasLogFileTailCollectionSize() bool`

HasLogFileTailCollectionSize returns a boolean if a field has been set.

### GetComment

`func (o *CollectSupportDataRecurringTaskShared) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CollectSupportDataRecurringTaskShared) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CollectSupportDataRecurringTaskShared) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CollectSupportDataRecurringTaskShared) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetRetainPreviousSupportDataArchiveCount

`func (o *CollectSupportDataRecurringTaskShared) GetRetainPreviousSupportDataArchiveCount() int32`

GetRetainPreviousSupportDataArchiveCount returns the RetainPreviousSupportDataArchiveCount field if non-nil, zero value otherwise.

### GetRetainPreviousSupportDataArchiveCountOk

`func (o *CollectSupportDataRecurringTaskShared) GetRetainPreviousSupportDataArchiveCountOk() (*int32, bool)`

GetRetainPreviousSupportDataArchiveCountOk returns a tuple with the RetainPreviousSupportDataArchiveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousSupportDataArchiveCount

`func (o *CollectSupportDataRecurringTaskShared) SetRetainPreviousSupportDataArchiveCount(v int32)`

SetRetainPreviousSupportDataArchiveCount sets RetainPreviousSupportDataArchiveCount field to given value.

### HasRetainPreviousSupportDataArchiveCount

`func (o *CollectSupportDataRecurringTaskShared) HasRetainPreviousSupportDataArchiveCount() bool`

HasRetainPreviousSupportDataArchiveCount returns a boolean if a field has been set.

### GetRetainPreviousSupportDataArchiveAge

`func (o *CollectSupportDataRecurringTaskShared) GetRetainPreviousSupportDataArchiveAge() string`

GetRetainPreviousSupportDataArchiveAge returns the RetainPreviousSupportDataArchiveAge field if non-nil, zero value otherwise.

### GetRetainPreviousSupportDataArchiveAgeOk

`func (o *CollectSupportDataRecurringTaskShared) GetRetainPreviousSupportDataArchiveAgeOk() (*string, bool)`

GetRetainPreviousSupportDataArchiveAgeOk returns a tuple with the RetainPreviousSupportDataArchiveAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousSupportDataArchiveAge

`func (o *CollectSupportDataRecurringTaskShared) SetRetainPreviousSupportDataArchiveAge(v string)`

SetRetainPreviousSupportDataArchiveAge sets RetainPreviousSupportDataArchiveAge field to given value.

### HasRetainPreviousSupportDataArchiveAge

`func (o *CollectSupportDataRecurringTaskShared) HasRetainPreviousSupportDataArchiveAge() bool`

HasRetainPreviousSupportDataArchiveAge returns a boolean if a field has been set.

### GetDescription

`func (o *CollectSupportDataRecurringTaskShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CollectSupportDataRecurringTaskShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CollectSupportDataRecurringTaskShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CollectSupportDataRecurringTaskShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCancelOnTaskDependencyFailure

`func (o *CollectSupportDataRecurringTaskShared) GetCancelOnTaskDependencyFailure() bool`

GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field if non-nil, zero value otherwise.

### GetCancelOnTaskDependencyFailureOk

`func (o *CollectSupportDataRecurringTaskShared) GetCancelOnTaskDependencyFailureOk() (*bool, bool)`

GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOnTaskDependencyFailure

`func (o *CollectSupportDataRecurringTaskShared) SetCancelOnTaskDependencyFailure(v bool)`

SetCancelOnTaskDependencyFailure sets CancelOnTaskDependencyFailure field to given value.

### HasCancelOnTaskDependencyFailure

`func (o *CollectSupportDataRecurringTaskShared) HasCancelOnTaskDependencyFailure() bool`

HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.

### GetEmailOnStart

`func (o *CollectSupportDataRecurringTaskShared) GetEmailOnStart() []string`

GetEmailOnStart returns the EmailOnStart field if non-nil, zero value otherwise.

### GetEmailOnStartOk

`func (o *CollectSupportDataRecurringTaskShared) GetEmailOnStartOk() (*[]string, bool)`

GetEmailOnStartOk returns a tuple with the EmailOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnStart

`func (o *CollectSupportDataRecurringTaskShared) SetEmailOnStart(v []string)`

SetEmailOnStart sets EmailOnStart field to given value.

### HasEmailOnStart

`func (o *CollectSupportDataRecurringTaskShared) HasEmailOnStart() bool`

HasEmailOnStart returns a boolean if a field has been set.

### GetEmailOnSuccess

`func (o *CollectSupportDataRecurringTaskShared) GetEmailOnSuccess() []string`

GetEmailOnSuccess returns the EmailOnSuccess field if non-nil, zero value otherwise.

### GetEmailOnSuccessOk

`func (o *CollectSupportDataRecurringTaskShared) GetEmailOnSuccessOk() (*[]string, bool)`

GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnSuccess

`func (o *CollectSupportDataRecurringTaskShared) SetEmailOnSuccess(v []string)`

SetEmailOnSuccess sets EmailOnSuccess field to given value.

### HasEmailOnSuccess

`func (o *CollectSupportDataRecurringTaskShared) HasEmailOnSuccess() bool`

HasEmailOnSuccess returns a boolean if a field has been set.

### GetEmailOnFailure

`func (o *CollectSupportDataRecurringTaskShared) GetEmailOnFailure() []string`

GetEmailOnFailure returns the EmailOnFailure field if non-nil, zero value otherwise.

### GetEmailOnFailureOk

`func (o *CollectSupportDataRecurringTaskShared) GetEmailOnFailureOk() (*[]string, bool)`

GetEmailOnFailureOk returns a tuple with the EmailOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnFailure

`func (o *CollectSupportDataRecurringTaskShared) SetEmailOnFailure(v []string)`

SetEmailOnFailure sets EmailOnFailure field to given value.

### HasEmailOnFailure

`func (o *CollectSupportDataRecurringTaskShared) HasEmailOnFailure() bool`

HasEmailOnFailure returns a boolean if a field has been set.

### GetAlertOnStart

`func (o *CollectSupportDataRecurringTaskShared) GetAlertOnStart() bool`

GetAlertOnStart returns the AlertOnStart field if non-nil, zero value otherwise.

### GetAlertOnStartOk

`func (o *CollectSupportDataRecurringTaskShared) GetAlertOnStartOk() (*bool, bool)`

GetAlertOnStartOk returns a tuple with the AlertOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnStart

`func (o *CollectSupportDataRecurringTaskShared) SetAlertOnStart(v bool)`

SetAlertOnStart sets AlertOnStart field to given value.

### HasAlertOnStart

`func (o *CollectSupportDataRecurringTaskShared) HasAlertOnStart() bool`

HasAlertOnStart returns a boolean if a field has been set.

### GetAlertOnSuccess

`func (o *CollectSupportDataRecurringTaskShared) GetAlertOnSuccess() bool`

GetAlertOnSuccess returns the AlertOnSuccess field if non-nil, zero value otherwise.

### GetAlertOnSuccessOk

`func (o *CollectSupportDataRecurringTaskShared) GetAlertOnSuccessOk() (*bool, bool)`

GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnSuccess

`func (o *CollectSupportDataRecurringTaskShared) SetAlertOnSuccess(v bool)`

SetAlertOnSuccess sets AlertOnSuccess field to given value.

### HasAlertOnSuccess

`func (o *CollectSupportDataRecurringTaskShared) HasAlertOnSuccess() bool`

HasAlertOnSuccess returns a boolean if a field has been set.

### GetAlertOnFailure

`func (o *CollectSupportDataRecurringTaskShared) GetAlertOnFailure() bool`

GetAlertOnFailure returns the AlertOnFailure field if non-nil, zero value otherwise.

### GetAlertOnFailureOk

`func (o *CollectSupportDataRecurringTaskShared) GetAlertOnFailureOk() (*bool, bool)`

GetAlertOnFailureOk returns a tuple with the AlertOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFailure

`func (o *CollectSupportDataRecurringTaskShared) SetAlertOnFailure(v bool)`

SetAlertOnFailure sets AlertOnFailure field to given value.

### HasAlertOnFailure

`func (o *CollectSupportDataRecurringTaskShared) HasAlertOnFailure() bool`

HasAlertOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


