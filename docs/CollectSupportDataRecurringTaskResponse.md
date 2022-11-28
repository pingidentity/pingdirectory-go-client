# CollectSupportDataRecurringTaskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Recurring Task | 
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

### NewCollectSupportDataRecurringTaskResponse

`func NewCollectSupportDataRecurringTaskResponse(id string, schemas []EnumcollectSupportDataRecurringTaskSchemaUrn, outputDirectory string, ) *CollectSupportDataRecurringTaskResponse`

NewCollectSupportDataRecurringTaskResponse instantiates a new CollectSupportDataRecurringTaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectSupportDataRecurringTaskResponseWithDefaults

`func NewCollectSupportDataRecurringTaskResponseWithDefaults() *CollectSupportDataRecurringTaskResponse`

NewCollectSupportDataRecurringTaskResponseWithDefaults instantiates a new CollectSupportDataRecurringTaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *CollectSupportDataRecurringTaskResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CollectSupportDataRecurringTaskResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CollectSupportDataRecurringTaskResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CollectSupportDataRecurringTaskResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *CollectSupportDataRecurringTaskResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *CollectSupportDataRecurringTaskResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *CollectSupportDataRecurringTaskResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *CollectSupportDataRecurringTaskResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *CollectSupportDataRecurringTaskResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CollectSupportDataRecurringTaskResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CollectSupportDataRecurringTaskResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *CollectSupportDataRecurringTaskResponse) GetSchemas() []EnumcollectSupportDataRecurringTaskSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *CollectSupportDataRecurringTaskResponse) GetSchemasOk() (*[]EnumcollectSupportDataRecurringTaskSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *CollectSupportDataRecurringTaskResponse) SetSchemas(v []EnumcollectSupportDataRecurringTaskSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetOutputDirectory

`func (o *CollectSupportDataRecurringTaskResponse) GetOutputDirectory() string`

GetOutputDirectory returns the OutputDirectory field if non-nil, zero value otherwise.

### GetOutputDirectoryOk

`func (o *CollectSupportDataRecurringTaskResponse) GetOutputDirectoryOk() (*string, bool)`

GetOutputDirectoryOk returns a tuple with the OutputDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDirectory

`func (o *CollectSupportDataRecurringTaskResponse) SetOutputDirectory(v string)`

SetOutputDirectory sets OutputDirectory field to given value.


### GetEncryptionPassphraseFile

`func (o *CollectSupportDataRecurringTaskResponse) GetEncryptionPassphraseFile() string`

GetEncryptionPassphraseFile returns the EncryptionPassphraseFile field if non-nil, zero value otherwise.

### GetEncryptionPassphraseFileOk

`func (o *CollectSupportDataRecurringTaskResponse) GetEncryptionPassphraseFileOk() (*string, bool)`

GetEncryptionPassphraseFileOk returns a tuple with the EncryptionPassphraseFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionPassphraseFile

`func (o *CollectSupportDataRecurringTaskResponse) SetEncryptionPassphraseFile(v string)`

SetEncryptionPassphraseFile sets EncryptionPassphraseFile field to given value.

### HasEncryptionPassphraseFile

`func (o *CollectSupportDataRecurringTaskResponse) HasEncryptionPassphraseFile() bool`

HasEncryptionPassphraseFile returns a boolean if a field has been set.

### GetIncludeExpensiveData

`func (o *CollectSupportDataRecurringTaskResponse) GetIncludeExpensiveData() bool`

GetIncludeExpensiveData returns the IncludeExpensiveData field if non-nil, zero value otherwise.

### GetIncludeExpensiveDataOk

`func (o *CollectSupportDataRecurringTaskResponse) GetIncludeExpensiveDataOk() (*bool, bool)`

GetIncludeExpensiveDataOk returns a tuple with the IncludeExpensiveData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeExpensiveData

`func (o *CollectSupportDataRecurringTaskResponse) SetIncludeExpensiveData(v bool)`

SetIncludeExpensiveData sets IncludeExpensiveData field to given value.

### HasIncludeExpensiveData

`func (o *CollectSupportDataRecurringTaskResponse) HasIncludeExpensiveData() bool`

HasIncludeExpensiveData returns a boolean if a field has been set.

### GetIncludeReplicationStateDump

`func (o *CollectSupportDataRecurringTaskResponse) GetIncludeReplicationStateDump() bool`

GetIncludeReplicationStateDump returns the IncludeReplicationStateDump field if non-nil, zero value otherwise.

### GetIncludeReplicationStateDumpOk

`func (o *CollectSupportDataRecurringTaskResponse) GetIncludeReplicationStateDumpOk() (*bool, bool)`

GetIncludeReplicationStateDumpOk returns a tuple with the IncludeReplicationStateDump field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeReplicationStateDump

`func (o *CollectSupportDataRecurringTaskResponse) SetIncludeReplicationStateDump(v bool)`

SetIncludeReplicationStateDump sets IncludeReplicationStateDump field to given value.

### HasIncludeReplicationStateDump

`func (o *CollectSupportDataRecurringTaskResponse) HasIncludeReplicationStateDump() bool`

HasIncludeReplicationStateDump returns a boolean if a field has been set.

### GetIncludeBinaryFiles

`func (o *CollectSupportDataRecurringTaskResponse) GetIncludeBinaryFiles() bool`

GetIncludeBinaryFiles returns the IncludeBinaryFiles field if non-nil, zero value otherwise.

### GetIncludeBinaryFilesOk

`func (o *CollectSupportDataRecurringTaskResponse) GetIncludeBinaryFilesOk() (*bool, bool)`

GetIncludeBinaryFilesOk returns a tuple with the IncludeBinaryFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBinaryFiles

`func (o *CollectSupportDataRecurringTaskResponse) SetIncludeBinaryFiles(v bool)`

SetIncludeBinaryFiles sets IncludeBinaryFiles field to given value.

### HasIncludeBinaryFiles

`func (o *CollectSupportDataRecurringTaskResponse) HasIncludeBinaryFiles() bool`

HasIncludeBinaryFiles returns a boolean if a field has been set.

### GetIncludeExtensionSource

`func (o *CollectSupportDataRecurringTaskResponse) GetIncludeExtensionSource() bool`

GetIncludeExtensionSource returns the IncludeExtensionSource field if non-nil, zero value otherwise.

### GetIncludeExtensionSourceOk

`func (o *CollectSupportDataRecurringTaskResponse) GetIncludeExtensionSourceOk() (*bool, bool)`

GetIncludeExtensionSourceOk returns a tuple with the IncludeExtensionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeExtensionSource

`func (o *CollectSupportDataRecurringTaskResponse) SetIncludeExtensionSource(v bool)`

SetIncludeExtensionSource sets IncludeExtensionSource field to given value.

### HasIncludeExtensionSource

`func (o *CollectSupportDataRecurringTaskResponse) HasIncludeExtensionSource() bool`

HasIncludeExtensionSource returns a boolean if a field has been set.

### GetUseSequentialMode

`func (o *CollectSupportDataRecurringTaskResponse) GetUseSequentialMode() bool`

GetUseSequentialMode returns the UseSequentialMode field if non-nil, zero value otherwise.

### GetUseSequentialModeOk

`func (o *CollectSupportDataRecurringTaskResponse) GetUseSequentialModeOk() (*bool, bool)`

GetUseSequentialModeOk returns a tuple with the UseSequentialMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSequentialMode

`func (o *CollectSupportDataRecurringTaskResponse) SetUseSequentialMode(v bool)`

SetUseSequentialMode sets UseSequentialMode field to given value.

### HasUseSequentialMode

`func (o *CollectSupportDataRecurringTaskResponse) HasUseSequentialMode() bool`

HasUseSequentialMode returns a boolean if a field has been set.

### GetSecurityLevel

`func (o *CollectSupportDataRecurringTaskResponse) GetSecurityLevel() EnumrecurringTaskSecurityLevelProp`

GetSecurityLevel returns the SecurityLevel field if non-nil, zero value otherwise.

### GetSecurityLevelOk

`func (o *CollectSupportDataRecurringTaskResponse) GetSecurityLevelOk() (*EnumrecurringTaskSecurityLevelProp, bool)`

GetSecurityLevelOk returns a tuple with the SecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLevel

`func (o *CollectSupportDataRecurringTaskResponse) SetSecurityLevel(v EnumrecurringTaskSecurityLevelProp)`

SetSecurityLevel sets SecurityLevel field to given value.

### HasSecurityLevel

`func (o *CollectSupportDataRecurringTaskResponse) HasSecurityLevel() bool`

HasSecurityLevel returns a boolean if a field has been set.

### GetJstackCount

`func (o *CollectSupportDataRecurringTaskResponse) GetJstackCount() int32`

GetJstackCount returns the JstackCount field if non-nil, zero value otherwise.

### GetJstackCountOk

`func (o *CollectSupportDataRecurringTaskResponse) GetJstackCountOk() (*int32, bool)`

GetJstackCountOk returns a tuple with the JstackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJstackCount

`func (o *CollectSupportDataRecurringTaskResponse) SetJstackCount(v int32)`

SetJstackCount sets JstackCount field to given value.

### HasJstackCount

`func (o *CollectSupportDataRecurringTaskResponse) HasJstackCount() bool`

HasJstackCount returns a boolean if a field has been set.

### GetReportCount

`func (o *CollectSupportDataRecurringTaskResponse) GetReportCount() int32`

GetReportCount returns the ReportCount field if non-nil, zero value otherwise.

### GetReportCountOk

`func (o *CollectSupportDataRecurringTaskResponse) GetReportCountOk() (*int32, bool)`

GetReportCountOk returns a tuple with the ReportCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportCount

`func (o *CollectSupportDataRecurringTaskResponse) SetReportCount(v int32)`

SetReportCount sets ReportCount field to given value.

### HasReportCount

`func (o *CollectSupportDataRecurringTaskResponse) HasReportCount() bool`

HasReportCount returns a boolean if a field has been set.

### GetReportIntervalSeconds

`func (o *CollectSupportDataRecurringTaskResponse) GetReportIntervalSeconds() int32`

GetReportIntervalSeconds returns the ReportIntervalSeconds field if non-nil, zero value otherwise.

### GetReportIntervalSecondsOk

`func (o *CollectSupportDataRecurringTaskResponse) GetReportIntervalSecondsOk() (*int32, bool)`

GetReportIntervalSecondsOk returns a tuple with the ReportIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportIntervalSeconds

`func (o *CollectSupportDataRecurringTaskResponse) SetReportIntervalSeconds(v int32)`

SetReportIntervalSeconds sets ReportIntervalSeconds field to given value.

### HasReportIntervalSeconds

`func (o *CollectSupportDataRecurringTaskResponse) HasReportIntervalSeconds() bool`

HasReportIntervalSeconds returns a boolean if a field has been set.

### GetLogDuration

`func (o *CollectSupportDataRecurringTaskResponse) GetLogDuration() string`

GetLogDuration returns the LogDuration field if non-nil, zero value otherwise.

### GetLogDurationOk

`func (o *CollectSupportDataRecurringTaskResponse) GetLogDurationOk() (*string, bool)`

GetLogDurationOk returns a tuple with the LogDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDuration

`func (o *CollectSupportDataRecurringTaskResponse) SetLogDuration(v string)`

SetLogDuration sets LogDuration field to given value.

### HasLogDuration

`func (o *CollectSupportDataRecurringTaskResponse) HasLogDuration() bool`

HasLogDuration returns a boolean if a field has been set.

### GetLogFileHeadCollectionSize

`func (o *CollectSupportDataRecurringTaskResponse) GetLogFileHeadCollectionSize() string`

GetLogFileHeadCollectionSize returns the LogFileHeadCollectionSize field if non-nil, zero value otherwise.

### GetLogFileHeadCollectionSizeOk

`func (o *CollectSupportDataRecurringTaskResponse) GetLogFileHeadCollectionSizeOk() (*string, bool)`

GetLogFileHeadCollectionSizeOk returns a tuple with the LogFileHeadCollectionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFileHeadCollectionSize

`func (o *CollectSupportDataRecurringTaskResponse) SetLogFileHeadCollectionSize(v string)`

SetLogFileHeadCollectionSize sets LogFileHeadCollectionSize field to given value.

### HasLogFileHeadCollectionSize

`func (o *CollectSupportDataRecurringTaskResponse) HasLogFileHeadCollectionSize() bool`

HasLogFileHeadCollectionSize returns a boolean if a field has been set.

### GetLogFileTailCollectionSize

`func (o *CollectSupportDataRecurringTaskResponse) GetLogFileTailCollectionSize() string`

GetLogFileTailCollectionSize returns the LogFileTailCollectionSize field if non-nil, zero value otherwise.

### GetLogFileTailCollectionSizeOk

`func (o *CollectSupportDataRecurringTaskResponse) GetLogFileTailCollectionSizeOk() (*string, bool)`

GetLogFileTailCollectionSizeOk returns a tuple with the LogFileTailCollectionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFileTailCollectionSize

`func (o *CollectSupportDataRecurringTaskResponse) SetLogFileTailCollectionSize(v string)`

SetLogFileTailCollectionSize sets LogFileTailCollectionSize field to given value.

### HasLogFileTailCollectionSize

`func (o *CollectSupportDataRecurringTaskResponse) HasLogFileTailCollectionSize() bool`

HasLogFileTailCollectionSize returns a boolean if a field has been set.

### GetComment

`func (o *CollectSupportDataRecurringTaskResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CollectSupportDataRecurringTaskResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CollectSupportDataRecurringTaskResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CollectSupportDataRecurringTaskResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetRetainPreviousSupportDataArchiveCount

`func (o *CollectSupportDataRecurringTaskResponse) GetRetainPreviousSupportDataArchiveCount() int32`

GetRetainPreviousSupportDataArchiveCount returns the RetainPreviousSupportDataArchiveCount field if non-nil, zero value otherwise.

### GetRetainPreviousSupportDataArchiveCountOk

`func (o *CollectSupportDataRecurringTaskResponse) GetRetainPreviousSupportDataArchiveCountOk() (*int32, bool)`

GetRetainPreviousSupportDataArchiveCountOk returns a tuple with the RetainPreviousSupportDataArchiveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousSupportDataArchiveCount

`func (o *CollectSupportDataRecurringTaskResponse) SetRetainPreviousSupportDataArchiveCount(v int32)`

SetRetainPreviousSupportDataArchiveCount sets RetainPreviousSupportDataArchiveCount field to given value.

### HasRetainPreviousSupportDataArchiveCount

`func (o *CollectSupportDataRecurringTaskResponse) HasRetainPreviousSupportDataArchiveCount() bool`

HasRetainPreviousSupportDataArchiveCount returns a boolean if a field has been set.

### GetRetainPreviousSupportDataArchiveAge

`func (o *CollectSupportDataRecurringTaskResponse) GetRetainPreviousSupportDataArchiveAge() string`

GetRetainPreviousSupportDataArchiveAge returns the RetainPreviousSupportDataArchiveAge field if non-nil, zero value otherwise.

### GetRetainPreviousSupportDataArchiveAgeOk

`func (o *CollectSupportDataRecurringTaskResponse) GetRetainPreviousSupportDataArchiveAgeOk() (*string, bool)`

GetRetainPreviousSupportDataArchiveAgeOk returns a tuple with the RetainPreviousSupportDataArchiveAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousSupportDataArchiveAge

`func (o *CollectSupportDataRecurringTaskResponse) SetRetainPreviousSupportDataArchiveAge(v string)`

SetRetainPreviousSupportDataArchiveAge sets RetainPreviousSupportDataArchiveAge field to given value.

### HasRetainPreviousSupportDataArchiveAge

`func (o *CollectSupportDataRecurringTaskResponse) HasRetainPreviousSupportDataArchiveAge() bool`

HasRetainPreviousSupportDataArchiveAge returns a boolean if a field has been set.

### GetDescription

`func (o *CollectSupportDataRecurringTaskResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CollectSupportDataRecurringTaskResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CollectSupportDataRecurringTaskResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CollectSupportDataRecurringTaskResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCancelOnTaskDependencyFailure

`func (o *CollectSupportDataRecurringTaskResponse) GetCancelOnTaskDependencyFailure() bool`

GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field if non-nil, zero value otherwise.

### GetCancelOnTaskDependencyFailureOk

`func (o *CollectSupportDataRecurringTaskResponse) GetCancelOnTaskDependencyFailureOk() (*bool, bool)`

GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOnTaskDependencyFailure

`func (o *CollectSupportDataRecurringTaskResponse) SetCancelOnTaskDependencyFailure(v bool)`

SetCancelOnTaskDependencyFailure sets CancelOnTaskDependencyFailure field to given value.

### HasCancelOnTaskDependencyFailure

`func (o *CollectSupportDataRecurringTaskResponse) HasCancelOnTaskDependencyFailure() bool`

HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.

### GetEmailOnStart

`func (o *CollectSupportDataRecurringTaskResponse) GetEmailOnStart() []string`

GetEmailOnStart returns the EmailOnStart field if non-nil, zero value otherwise.

### GetEmailOnStartOk

`func (o *CollectSupportDataRecurringTaskResponse) GetEmailOnStartOk() (*[]string, bool)`

GetEmailOnStartOk returns a tuple with the EmailOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnStart

`func (o *CollectSupportDataRecurringTaskResponse) SetEmailOnStart(v []string)`

SetEmailOnStart sets EmailOnStart field to given value.

### HasEmailOnStart

`func (o *CollectSupportDataRecurringTaskResponse) HasEmailOnStart() bool`

HasEmailOnStart returns a boolean if a field has been set.

### GetEmailOnSuccess

`func (o *CollectSupportDataRecurringTaskResponse) GetEmailOnSuccess() []string`

GetEmailOnSuccess returns the EmailOnSuccess field if non-nil, zero value otherwise.

### GetEmailOnSuccessOk

`func (o *CollectSupportDataRecurringTaskResponse) GetEmailOnSuccessOk() (*[]string, bool)`

GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnSuccess

`func (o *CollectSupportDataRecurringTaskResponse) SetEmailOnSuccess(v []string)`

SetEmailOnSuccess sets EmailOnSuccess field to given value.

### HasEmailOnSuccess

`func (o *CollectSupportDataRecurringTaskResponse) HasEmailOnSuccess() bool`

HasEmailOnSuccess returns a boolean if a field has been set.

### GetEmailOnFailure

`func (o *CollectSupportDataRecurringTaskResponse) GetEmailOnFailure() []string`

GetEmailOnFailure returns the EmailOnFailure field if non-nil, zero value otherwise.

### GetEmailOnFailureOk

`func (o *CollectSupportDataRecurringTaskResponse) GetEmailOnFailureOk() (*[]string, bool)`

GetEmailOnFailureOk returns a tuple with the EmailOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnFailure

`func (o *CollectSupportDataRecurringTaskResponse) SetEmailOnFailure(v []string)`

SetEmailOnFailure sets EmailOnFailure field to given value.

### HasEmailOnFailure

`func (o *CollectSupportDataRecurringTaskResponse) HasEmailOnFailure() bool`

HasEmailOnFailure returns a boolean if a field has been set.

### GetAlertOnStart

`func (o *CollectSupportDataRecurringTaskResponse) GetAlertOnStart() bool`

GetAlertOnStart returns the AlertOnStart field if non-nil, zero value otherwise.

### GetAlertOnStartOk

`func (o *CollectSupportDataRecurringTaskResponse) GetAlertOnStartOk() (*bool, bool)`

GetAlertOnStartOk returns a tuple with the AlertOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnStart

`func (o *CollectSupportDataRecurringTaskResponse) SetAlertOnStart(v bool)`

SetAlertOnStart sets AlertOnStart field to given value.

### HasAlertOnStart

`func (o *CollectSupportDataRecurringTaskResponse) HasAlertOnStart() bool`

HasAlertOnStart returns a boolean if a field has been set.

### GetAlertOnSuccess

`func (o *CollectSupportDataRecurringTaskResponse) GetAlertOnSuccess() bool`

GetAlertOnSuccess returns the AlertOnSuccess field if non-nil, zero value otherwise.

### GetAlertOnSuccessOk

`func (o *CollectSupportDataRecurringTaskResponse) GetAlertOnSuccessOk() (*bool, bool)`

GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnSuccess

`func (o *CollectSupportDataRecurringTaskResponse) SetAlertOnSuccess(v bool)`

SetAlertOnSuccess sets AlertOnSuccess field to given value.

### HasAlertOnSuccess

`func (o *CollectSupportDataRecurringTaskResponse) HasAlertOnSuccess() bool`

HasAlertOnSuccess returns a boolean if a field has been set.

### GetAlertOnFailure

`func (o *CollectSupportDataRecurringTaskResponse) GetAlertOnFailure() bool`

GetAlertOnFailure returns the AlertOnFailure field if non-nil, zero value otherwise.

### GetAlertOnFailureOk

`func (o *CollectSupportDataRecurringTaskResponse) GetAlertOnFailureOk() (*bool, bool)`

GetAlertOnFailureOk returns a tuple with the AlertOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFailure

`func (o *CollectSupportDataRecurringTaskResponse) SetAlertOnFailure(v bool)`

SetAlertOnFailure sets AlertOnFailure field to given value.

### HasAlertOnFailure

`func (o *CollectSupportDataRecurringTaskResponse) HasAlertOnFailure() bool`

HasAlertOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


