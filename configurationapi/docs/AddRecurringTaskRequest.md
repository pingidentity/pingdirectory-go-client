# AddRecurringTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskName** | **string** | Name of the new Recurring Task | 
**Schemas** | [**[]EnumthirdPartyRecurringTaskSchemaUrn**](EnumthirdPartyRecurringTaskSchemaUrn.md) |  | 
**ProfileDirectory** | **string** | The directory in which the generated server profiles will be placed. The files will be named with the pattern \&quot;server-profile-{timestamp}.zip\&quot;, where \&quot;{timestamp}\&quot; represents the time that the profile was generated. | 
**IncludePath** | Pointer to **[]string** | An optional set of additional paths to files within the instance root that should be included in the generated server profile. All paths must be within the instance root, and relative paths will be relative to the instance root. | [optional] 
**RetainPreviousProfileCount** | Pointer to **int32** | The minimum number of previous server profile zip files that should be preserved after a new profile is generated. | [optional] 
**RetainPreviousProfileAge** | Pointer to **string** | The minimum age of previous server profile zip files that should be preserved after a new profile is generated. | [optional] 
**Description** | Pointer to **string** | A description for this Recurring Task | [optional] 
**CancelOnTaskDependencyFailure** | Pointer to **bool** | Indicates whether an instance of this Recurring Task should be canceled if the task immediately before it in the recurring task chain fails to complete successfully (including if it is canceled by an administrator before it starts or while it is running). | [optional] 
**EmailOnStart** | Pointer to **[]string** | The email addresses to which a message should be sent whenever an instance of this Recurring Task starts running. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**EmailOnSuccess** | Pointer to **[]string** | The email addresses to which a message should be sent whenever an instance of this Recurring Task completes successfully. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**EmailOnFailure** | Pointer to **[]string** | The email addresses to which a message should be sent if an instance of this Recurring Task fails to complete successfully. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**AlertOnStart** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task starts running. | [optional] 
**AlertOnSuccess** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task completes successfully. | [optional] 
**AlertOnFailure** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task fails to complete successfully. | [optional] 
**Reason** | Pointer to **string** | The reason that the server is being placed in lockdown mode. | [optional] 
**BackupDirectory** | Pointer to **string** | The directory in which backup files will be placed. When backing up a single backend, the backup files will be placed directly in this directory. When backing up multiple backends, the backup files for each backend will be placed in a subdirectory whose name is the corresponding backend ID. | [optional] 
**IncludedBackendID** | Pointer to **[]string** | The backend IDs of any backends that should be included in the backup. | [optional] 
**ExcludedBackendID** | Pointer to **[]string** | The backend IDs of any backends that should be excluded from the backup. All backends that support backups and are not listed will be included. | [optional] 
**Compress** | Pointer to **bool** | Indicates whether to compress the LDIF data as it is exported. | [optional] 
**Encrypt** | Pointer to **bool** | Indicates whether to encrypt the LDIF data as it exported. | [optional] 
**EncryptionSettingsDefinitionID** | Pointer to **string** | The ID of an encryption settings definition to use to obtain the LDIF export encryption key. | [optional] 
**Sign** | Pointer to **bool** | Indicates whether to cryptographically sign the exported data, which will make it possible to detect whether the LDIF data has been altered since it was exported. | [optional] 
**RetainPreviousFullBackupCount** | Pointer to **int32** | The minimum number of previous full backups that should be preserved after a new backup completes successfully. | [optional] 
**RetainPreviousFullBackupAge** | Pointer to **string** | The minimum age of previous full backups that should be preserved after a new backup completes successfully. | [optional] 
**MaxMegabytesPerSecond** | Pointer to **int32** | The maximum rate, in megabytes per second, at which LDIF exports should be written. | [optional] 
**SleepDuration** | Pointer to **string** | The length of time to sleep before the task completes. | [optional] 
**DurationToWaitForWorkQueueIdle** | Pointer to **string** | Indicates that task should wait for up to the specified length of time for the work queue to report that all worker threads are idle and there are no pending operations. Note that this primarily monitors operations that use worker threads, which does not include internal operations (for example, those invoked by extensions), and may not include requests from non-LDAP clients (for example, HTTP-based clients). | [optional] 
**LdapURLForSearchExpectedToReturnEntries** | Pointer to **[]string** | An LDAP URL that provides the criteria for a search request that is expected to return at least one entry. The search will be performed internally, and only the base DN, scope, and filter from the URL will be used; any host, port, or requested attributes included in the URL will be ignored. | [optional] 
**SearchInterval** | Pointer to **string** | The length of time the server should sleep between searches performed using the criteria from the ldap-url-for-search-expected-to-return-entries property. | [optional] 
**SearchTimeLimit** | Pointer to **string** | The length of time that the server will wait for a response to each internal search performed using the criteria from the ldap-url-for-search-expected-to-return-entries property. | [optional] 
**DurationToWaitForSearchToReturnEntries** | Pointer to **string** | The maximum length of time that the server will continue to perform internal searches using the criteria from the ldap-url-for-search-expected-to-return-entries property. | [optional] 
**TaskReturnStateIfTimeoutIsEncountered** | Pointer to [**EnumrecurringTaskTaskReturnStateIfTimeoutIsEncounteredProp**](EnumrecurringTaskTaskReturnStateIfTimeoutIsEncounteredProp.md) |  | [optional] 
**TaskJavaClass** | **string** | The fully-qualified name of the Java class that provides the logic for the task to be invoked. | 
**TaskObjectClass** | **[]string** | The names or OIDs of the object classes to include in the tasks that are scheduled from this Statically Defined Recurring Task. All object classes must be defined in the server schema, and the combination of object classes must be valid for a task entry. | 
**TaskAttributeValue** | Pointer to **[]string** | The set of attribute values that should be included in the tasks that are scheduled from this Statically Defined Recurring Task. Each value must be in the form {attribute-type}&#x3D;{value}, where {attribute-type} is the name or OID of an attribute type that is defined in the schema and permitted with the configured set of object classes, and {value} is a value to assign to an attribute with that type. A multivalued attribute can be created by providing multiple name-value pairs with the same name and different values. | [optional] 
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
**LdifDirectory** | Pointer to **string** | The directory in which LDIF export files will be placed. The directory must already exist. | [optional] 
**BackendID** | Pointer to **[]string** | The backend ID for a backend to be exported. | [optional] 
**ExcludeBackendID** | Pointer to **[]string** | The backend ID for a backend to be excluded from the export. | [optional] 
**RetainPreviousLDIFExportCount** | Pointer to **int32** | The minimum number of previous LDIF exports that should be preserved after a new export completes successfully. | [optional] 
**RetainPreviousLDIFExportAge** | Pointer to **string** | The minimum age of previous LDIF exports that should be preserved after a new export completes successfully. | [optional] 
**CommandPath** | **string** | The absolute path to the command to execute. It must be an absolute path, the corresponding file must exist, and it must be listed in the config/exec-command-whitelist.txt file. | 
**CommandArguments** | Pointer to **string** | A string containing the arguments to provide to the command. If the command should be run without arguments, this property should be left undefined. If there should be multiple arguments, then they should be separated with spaces. | [optional] 
**CommandOutputFileBaseName** | Pointer to **string** | The path and base name for a file to which the command output (both standard output and standard error) should be written. This may be left undefined if the command output should not be recorded into a file. | [optional] 
**RetainPreviousOutputFileCount** | Pointer to **int32** | The minimum number of previous command output files that should be preserved after a new instance of the command is invoked. | [optional] 
**RetainPreviousOutputFileAge** | Pointer to **string** | The minimum age of previous command output files that should be preserved after a new instance of the command is invoked. | [optional] 
**LogCommandOutput** | Pointer to **bool** | Indicates whether the command&#39;s output (both standard output and standard error) should be recorded in the server&#39;s error log. | [optional] 
**TaskCompletionStateForNonzeroExitCode** | Pointer to [**EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp**](EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp.md) |  | [optional] 
**WorkingDirectory** | Pointer to **string** | The absolute path to a working directory where the command should be executed. It must be an absolute path and the corresponding directory must exist. | [optional] 
**TargetDirectory** | **string** | The path to the directory containing the files to examine. The directory must exist. | 
**FilenamePattern** | **string** | A pattern that specifies the names of the files to examine. The pattern may contain zero or more asterisks as wildcards, where each wildcard matches zero or more characters. It may also contain at most one occurrence of the special string \&quot;${timestamp}\&quot;, which will match a timestamp with the format specified using the timestamp-format property. All other characters in the pattern will be treated literally. | 
**TimestampFormat** | [**EnumrecurringTaskTimestampFormatProp**](EnumrecurringTaskTimestampFormatProp.md) |  | 
**RetainFileCount** | Pointer to **int32** | The minimum number of files matching the pattern that will be retained. | [optional] 
**RetainFileAge** | Pointer to **string** | The minimum age of files matching the pattern that will be retained. | [optional] 
**RetainAggregateFileSize** | Pointer to **string** | The minimum aggregate size of files that will be retained. The size should be specified as an integer followed by a unit that is one of \&quot;b\&quot; or \&quot;bytes\&quot;, \&quot;kb\&quot; or \&quot;kilobytes\&quot;, \&quot;mb\&quot; or \&quot;megabytes\&quot;, \&quot;gb\&quot; or \&quot;gigabytes\&quot;, or \&quot;tb\&quot; or \&quot;terabytes\&quot;. For example, a value of \&quot;1 gb\&quot; indicates that at least one gigabyte of files should be retained. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Recurring Task. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Recurring Task. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddRecurringTaskRequest

`func NewAddRecurringTaskRequest(taskName string, schemas []EnumthirdPartyRecurringTaskSchemaUrn, profileDirectory string, taskJavaClass string, taskObjectClass []string, outputDirectory string, commandPath string, targetDirectory string, filenamePattern string, timestampFormat EnumrecurringTaskTimestampFormatProp, extensionClass string, ) *AddRecurringTaskRequest`

NewAddRecurringTaskRequest instantiates a new AddRecurringTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddRecurringTaskRequestWithDefaults

`func NewAddRecurringTaskRequestWithDefaults() *AddRecurringTaskRequest`

NewAddRecurringTaskRequestWithDefaults instantiates a new AddRecurringTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskName

`func (o *AddRecurringTaskRequest) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *AddRecurringTaskRequest) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *AddRecurringTaskRequest) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.


### GetSchemas

`func (o *AddRecurringTaskRequest) GetSchemas() []EnumthirdPartyRecurringTaskSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddRecurringTaskRequest) GetSchemasOk() (*[]EnumthirdPartyRecurringTaskSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddRecurringTaskRequest) SetSchemas(v []EnumthirdPartyRecurringTaskSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetProfileDirectory

`func (o *AddRecurringTaskRequest) GetProfileDirectory() string`

GetProfileDirectory returns the ProfileDirectory field if non-nil, zero value otherwise.

### GetProfileDirectoryOk

`func (o *AddRecurringTaskRequest) GetProfileDirectoryOk() (*string, bool)`

GetProfileDirectoryOk returns a tuple with the ProfileDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileDirectory

`func (o *AddRecurringTaskRequest) SetProfileDirectory(v string)`

SetProfileDirectory sets ProfileDirectory field to given value.


### GetIncludePath

`func (o *AddRecurringTaskRequest) GetIncludePath() []string`

GetIncludePath returns the IncludePath field if non-nil, zero value otherwise.

### GetIncludePathOk

`func (o *AddRecurringTaskRequest) GetIncludePathOk() (*[]string, bool)`

GetIncludePathOk returns a tuple with the IncludePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePath

`func (o *AddRecurringTaskRequest) SetIncludePath(v []string)`

SetIncludePath sets IncludePath field to given value.

### HasIncludePath

`func (o *AddRecurringTaskRequest) HasIncludePath() bool`

HasIncludePath returns a boolean if a field has been set.

### GetRetainPreviousProfileCount

`func (o *AddRecurringTaskRequest) GetRetainPreviousProfileCount() int32`

GetRetainPreviousProfileCount returns the RetainPreviousProfileCount field if non-nil, zero value otherwise.

### GetRetainPreviousProfileCountOk

`func (o *AddRecurringTaskRequest) GetRetainPreviousProfileCountOk() (*int32, bool)`

GetRetainPreviousProfileCountOk returns a tuple with the RetainPreviousProfileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousProfileCount

`func (o *AddRecurringTaskRequest) SetRetainPreviousProfileCount(v int32)`

SetRetainPreviousProfileCount sets RetainPreviousProfileCount field to given value.

### HasRetainPreviousProfileCount

`func (o *AddRecurringTaskRequest) HasRetainPreviousProfileCount() bool`

HasRetainPreviousProfileCount returns a boolean if a field has been set.

### GetRetainPreviousProfileAge

`func (o *AddRecurringTaskRequest) GetRetainPreviousProfileAge() string`

GetRetainPreviousProfileAge returns the RetainPreviousProfileAge field if non-nil, zero value otherwise.

### GetRetainPreviousProfileAgeOk

`func (o *AddRecurringTaskRequest) GetRetainPreviousProfileAgeOk() (*string, bool)`

GetRetainPreviousProfileAgeOk returns a tuple with the RetainPreviousProfileAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousProfileAge

`func (o *AddRecurringTaskRequest) SetRetainPreviousProfileAge(v string)`

SetRetainPreviousProfileAge sets RetainPreviousProfileAge field to given value.

### HasRetainPreviousProfileAge

`func (o *AddRecurringTaskRequest) HasRetainPreviousProfileAge() bool`

HasRetainPreviousProfileAge returns a boolean if a field has been set.

### GetDescription

`func (o *AddRecurringTaskRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddRecurringTaskRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddRecurringTaskRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddRecurringTaskRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCancelOnTaskDependencyFailure

`func (o *AddRecurringTaskRequest) GetCancelOnTaskDependencyFailure() bool`

GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field if non-nil, zero value otherwise.

### GetCancelOnTaskDependencyFailureOk

`func (o *AddRecurringTaskRequest) GetCancelOnTaskDependencyFailureOk() (*bool, bool)`

GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOnTaskDependencyFailure

`func (o *AddRecurringTaskRequest) SetCancelOnTaskDependencyFailure(v bool)`

SetCancelOnTaskDependencyFailure sets CancelOnTaskDependencyFailure field to given value.

### HasCancelOnTaskDependencyFailure

`func (o *AddRecurringTaskRequest) HasCancelOnTaskDependencyFailure() bool`

HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.

### GetEmailOnStart

`func (o *AddRecurringTaskRequest) GetEmailOnStart() []string`

GetEmailOnStart returns the EmailOnStart field if non-nil, zero value otherwise.

### GetEmailOnStartOk

`func (o *AddRecurringTaskRequest) GetEmailOnStartOk() (*[]string, bool)`

GetEmailOnStartOk returns a tuple with the EmailOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnStart

`func (o *AddRecurringTaskRequest) SetEmailOnStart(v []string)`

SetEmailOnStart sets EmailOnStart field to given value.

### HasEmailOnStart

`func (o *AddRecurringTaskRequest) HasEmailOnStart() bool`

HasEmailOnStart returns a boolean if a field has been set.

### GetEmailOnSuccess

`func (o *AddRecurringTaskRequest) GetEmailOnSuccess() []string`

GetEmailOnSuccess returns the EmailOnSuccess field if non-nil, zero value otherwise.

### GetEmailOnSuccessOk

`func (o *AddRecurringTaskRequest) GetEmailOnSuccessOk() (*[]string, bool)`

GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnSuccess

`func (o *AddRecurringTaskRequest) SetEmailOnSuccess(v []string)`

SetEmailOnSuccess sets EmailOnSuccess field to given value.

### HasEmailOnSuccess

`func (o *AddRecurringTaskRequest) HasEmailOnSuccess() bool`

HasEmailOnSuccess returns a boolean if a field has been set.

### GetEmailOnFailure

`func (o *AddRecurringTaskRequest) GetEmailOnFailure() []string`

GetEmailOnFailure returns the EmailOnFailure field if non-nil, zero value otherwise.

### GetEmailOnFailureOk

`func (o *AddRecurringTaskRequest) GetEmailOnFailureOk() (*[]string, bool)`

GetEmailOnFailureOk returns a tuple with the EmailOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnFailure

`func (o *AddRecurringTaskRequest) SetEmailOnFailure(v []string)`

SetEmailOnFailure sets EmailOnFailure field to given value.

### HasEmailOnFailure

`func (o *AddRecurringTaskRequest) HasEmailOnFailure() bool`

HasEmailOnFailure returns a boolean if a field has been set.

### GetAlertOnStart

`func (o *AddRecurringTaskRequest) GetAlertOnStart() bool`

GetAlertOnStart returns the AlertOnStart field if non-nil, zero value otherwise.

### GetAlertOnStartOk

`func (o *AddRecurringTaskRequest) GetAlertOnStartOk() (*bool, bool)`

GetAlertOnStartOk returns a tuple with the AlertOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnStart

`func (o *AddRecurringTaskRequest) SetAlertOnStart(v bool)`

SetAlertOnStart sets AlertOnStart field to given value.

### HasAlertOnStart

`func (o *AddRecurringTaskRequest) HasAlertOnStart() bool`

HasAlertOnStart returns a boolean if a field has been set.

### GetAlertOnSuccess

`func (o *AddRecurringTaskRequest) GetAlertOnSuccess() bool`

GetAlertOnSuccess returns the AlertOnSuccess field if non-nil, zero value otherwise.

### GetAlertOnSuccessOk

`func (o *AddRecurringTaskRequest) GetAlertOnSuccessOk() (*bool, bool)`

GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnSuccess

`func (o *AddRecurringTaskRequest) SetAlertOnSuccess(v bool)`

SetAlertOnSuccess sets AlertOnSuccess field to given value.

### HasAlertOnSuccess

`func (o *AddRecurringTaskRequest) HasAlertOnSuccess() bool`

HasAlertOnSuccess returns a boolean if a field has been set.

### GetAlertOnFailure

`func (o *AddRecurringTaskRequest) GetAlertOnFailure() bool`

GetAlertOnFailure returns the AlertOnFailure field if non-nil, zero value otherwise.

### GetAlertOnFailureOk

`func (o *AddRecurringTaskRequest) GetAlertOnFailureOk() (*bool, bool)`

GetAlertOnFailureOk returns a tuple with the AlertOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFailure

`func (o *AddRecurringTaskRequest) SetAlertOnFailure(v bool)`

SetAlertOnFailure sets AlertOnFailure field to given value.

### HasAlertOnFailure

`func (o *AddRecurringTaskRequest) HasAlertOnFailure() bool`

HasAlertOnFailure returns a boolean if a field has been set.

### GetReason

`func (o *AddRecurringTaskRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AddRecurringTaskRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AddRecurringTaskRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AddRecurringTaskRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetBackupDirectory

`func (o *AddRecurringTaskRequest) GetBackupDirectory() string`

GetBackupDirectory returns the BackupDirectory field if non-nil, zero value otherwise.

### GetBackupDirectoryOk

`func (o *AddRecurringTaskRequest) GetBackupDirectoryOk() (*string, bool)`

GetBackupDirectoryOk returns a tuple with the BackupDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDirectory

`func (o *AddRecurringTaskRequest) SetBackupDirectory(v string)`

SetBackupDirectory sets BackupDirectory field to given value.

### HasBackupDirectory

`func (o *AddRecurringTaskRequest) HasBackupDirectory() bool`

HasBackupDirectory returns a boolean if a field has been set.

### GetIncludedBackendID

`func (o *AddRecurringTaskRequest) GetIncludedBackendID() []string`

GetIncludedBackendID returns the IncludedBackendID field if non-nil, zero value otherwise.

### GetIncludedBackendIDOk

`func (o *AddRecurringTaskRequest) GetIncludedBackendIDOk() (*[]string, bool)`

GetIncludedBackendIDOk returns a tuple with the IncludedBackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedBackendID

`func (o *AddRecurringTaskRequest) SetIncludedBackendID(v []string)`

SetIncludedBackendID sets IncludedBackendID field to given value.

### HasIncludedBackendID

`func (o *AddRecurringTaskRequest) HasIncludedBackendID() bool`

HasIncludedBackendID returns a boolean if a field has been set.

### GetExcludedBackendID

`func (o *AddRecurringTaskRequest) GetExcludedBackendID() []string`

GetExcludedBackendID returns the ExcludedBackendID field if non-nil, zero value otherwise.

### GetExcludedBackendIDOk

`func (o *AddRecurringTaskRequest) GetExcludedBackendIDOk() (*[]string, bool)`

GetExcludedBackendIDOk returns a tuple with the ExcludedBackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedBackendID

`func (o *AddRecurringTaskRequest) SetExcludedBackendID(v []string)`

SetExcludedBackendID sets ExcludedBackendID field to given value.

### HasExcludedBackendID

`func (o *AddRecurringTaskRequest) HasExcludedBackendID() bool`

HasExcludedBackendID returns a boolean if a field has been set.

### GetCompress

`func (o *AddRecurringTaskRequest) GetCompress() bool`

GetCompress returns the Compress field if non-nil, zero value otherwise.

### GetCompressOk

`func (o *AddRecurringTaskRequest) GetCompressOk() (*bool, bool)`

GetCompressOk returns a tuple with the Compress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompress

`func (o *AddRecurringTaskRequest) SetCompress(v bool)`

SetCompress sets Compress field to given value.

### HasCompress

`func (o *AddRecurringTaskRequest) HasCompress() bool`

HasCompress returns a boolean if a field has been set.

### GetEncrypt

`func (o *AddRecurringTaskRequest) GetEncrypt() bool`

GetEncrypt returns the Encrypt field if non-nil, zero value otherwise.

### GetEncryptOk

`func (o *AddRecurringTaskRequest) GetEncryptOk() (*bool, bool)`

GetEncryptOk returns a tuple with the Encrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypt

`func (o *AddRecurringTaskRequest) SetEncrypt(v bool)`

SetEncrypt sets Encrypt field to given value.

### HasEncrypt

`func (o *AddRecurringTaskRequest) HasEncrypt() bool`

HasEncrypt returns a boolean if a field has been set.

### GetEncryptionSettingsDefinitionID

`func (o *AddRecurringTaskRequest) GetEncryptionSettingsDefinitionID() string`

GetEncryptionSettingsDefinitionID returns the EncryptionSettingsDefinitionID field if non-nil, zero value otherwise.

### GetEncryptionSettingsDefinitionIDOk

`func (o *AddRecurringTaskRequest) GetEncryptionSettingsDefinitionIDOk() (*string, bool)`

GetEncryptionSettingsDefinitionIDOk returns a tuple with the EncryptionSettingsDefinitionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionSettingsDefinitionID

`func (o *AddRecurringTaskRequest) SetEncryptionSettingsDefinitionID(v string)`

SetEncryptionSettingsDefinitionID sets EncryptionSettingsDefinitionID field to given value.

### HasEncryptionSettingsDefinitionID

`func (o *AddRecurringTaskRequest) HasEncryptionSettingsDefinitionID() bool`

HasEncryptionSettingsDefinitionID returns a boolean if a field has been set.

### GetSign

`func (o *AddRecurringTaskRequest) GetSign() bool`

GetSign returns the Sign field if non-nil, zero value otherwise.

### GetSignOk

`func (o *AddRecurringTaskRequest) GetSignOk() (*bool, bool)`

GetSignOk returns a tuple with the Sign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSign

`func (o *AddRecurringTaskRequest) SetSign(v bool)`

SetSign sets Sign field to given value.

### HasSign

`func (o *AddRecurringTaskRequest) HasSign() bool`

HasSign returns a boolean if a field has been set.

### GetRetainPreviousFullBackupCount

`func (o *AddRecurringTaskRequest) GetRetainPreviousFullBackupCount() int32`

GetRetainPreviousFullBackupCount returns the RetainPreviousFullBackupCount field if non-nil, zero value otherwise.

### GetRetainPreviousFullBackupCountOk

`func (o *AddRecurringTaskRequest) GetRetainPreviousFullBackupCountOk() (*int32, bool)`

GetRetainPreviousFullBackupCountOk returns a tuple with the RetainPreviousFullBackupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousFullBackupCount

`func (o *AddRecurringTaskRequest) SetRetainPreviousFullBackupCount(v int32)`

SetRetainPreviousFullBackupCount sets RetainPreviousFullBackupCount field to given value.

### HasRetainPreviousFullBackupCount

`func (o *AddRecurringTaskRequest) HasRetainPreviousFullBackupCount() bool`

HasRetainPreviousFullBackupCount returns a boolean if a field has been set.

### GetRetainPreviousFullBackupAge

`func (o *AddRecurringTaskRequest) GetRetainPreviousFullBackupAge() string`

GetRetainPreviousFullBackupAge returns the RetainPreviousFullBackupAge field if non-nil, zero value otherwise.

### GetRetainPreviousFullBackupAgeOk

`func (o *AddRecurringTaskRequest) GetRetainPreviousFullBackupAgeOk() (*string, bool)`

GetRetainPreviousFullBackupAgeOk returns a tuple with the RetainPreviousFullBackupAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousFullBackupAge

`func (o *AddRecurringTaskRequest) SetRetainPreviousFullBackupAge(v string)`

SetRetainPreviousFullBackupAge sets RetainPreviousFullBackupAge field to given value.

### HasRetainPreviousFullBackupAge

`func (o *AddRecurringTaskRequest) HasRetainPreviousFullBackupAge() bool`

HasRetainPreviousFullBackupAge returns a boolean if a field has been set.

### GetMaxMegabytesPerSecond

`func (o *AddRecurringTaskRequest) GetMaxMegabytesPerSecond() int32`

GetMaxMegabytesPerSecond returns the MaxMegabytesPerSecond field if non-nil, zero value otherwise.

### GetMaxMegabytesPerSecondOk

`func (o *AddRecurringTaskRequest) GetMaxMegabytesPerSecondOk() (*int32, bool)`

GetMaxMegabytesPerSecondOk returns a tuple with the MaxMegabytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMegabytesPerSecond

`func (o *AddRecurringTaskRequest) SetMaxMegabytesPerSecond(v int32)`

SetMaxMegabytesPerSecond sets MaxMegabytesPerSecond field to given value.

### HasMaxMegabytesPerSecond

`func (o *AddRecurringTaskRequest) HasMaxMegabytesPerSecond() bool`

HasMaxMegabytesPerSecond returns a boolean if a field has been set.

### GetSleepDuration

`func (o *AddRecurringTaskRequest) GetSleepDuration() string`

GetSleepDuration returns the SleepDuration field if non-nil, zero value otherwise.

### GetSleepDurationOk

`func (o *AddRecurringTaskRequest) GetSleepDurationOk() (*string, bool)`

GetSleepDurationOk returns a tuple with the SleepDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepDuration

`func (o *AddRecurringTaskRequest) SetSleepDuration(v string)`

SetSleepDuration sets SleepDuration field to given value.

### HasSleepDuration

`func (o *AddRecurringTaskRequest) HasSleepDuration() bool`

HasSleepDuration returns a boolean if a field has been set.

### GetDurationToWaitForWorkQueueIdle

`func (o *AddRecurringTaskRequest) GetDurationToWaitForWorkQueueIdle() string`

GetDurationToWaitForWorkQueueIdle returns the DurationToWaitForWorkQueueIdle field if non-nil, zero value otherwise.

### GetDurationToWaitForWorkQueueIdleOk

`func (o *AddRecurringTaskRequest) GetDurationToWaitForWorkQueueIdleOk() (*string, bool)`

GetDurationToWaitForWorkQueueIdleOk returns a tuple with the DurationToWaitForWorkQueueIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationToWaitForWorkQueueIdle

`func (o *AddRecurringTaskRequest) SetDurationToWaitForWorkQueueIdle(v string)`

SetDurationToWaitForWorkQueueIdle sets DurationToWaitForWorkQueueIdle field to given value.

### HasDurationToWaitForWorkQueueIdle

`func (o *AddRecurringTaskRequest) HasDurationToWaitForWorkQueueIdle() bool`

HasDurationToWaitForWorkQueueIdle returns a boolean if a field has been set.

### GetLdapURLForSearchExpectedToReturnEntries

`func (o *AddRecurringTaskRequest) GetLdapURLForSearchExpectedToReturnEntries() []string`

GetLdapURLForSearchExpectedToReturnEntries returns the LdapURLForSearchExpectedToReturnEntries field if non-nil, zero value otherwise.

### GetLdapURLForSearchExpectedToReturnEntriesOk

`func (o *AddRecurringTaskRequest) GetLdapURLForSearchExpectedToReturnEntriesOk() (*[]string, bool)`

GetLdapURLForSearchExpectedToReturnEntriesOk returns a tuple with the LdapURLForSearchExpectedToReturnEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapURLForSearchExpectedToReturnEntries

`func (o *AddRecurringTaskRequest) SetLdapURLForSearchExpectedToReturnEntries(v []string)`

SetLdapURLForSearchExpectedToReturnEntries sets LdapURLForSearchExpectedToReturnEntries field to given value.

### HasLdapURLForSearchExpectedToReturnEntries

`func (o *AddRecurringTaskRequest) HasLdapURLForSearchExpectedToReturnEntries() bool`

HasLdapURLForSearchExpectedToReturnEntries returns a boolean if a field has been set.

### GetSearchInterval

`func (o *AddRecurringTaskRequest) GetSearchInterval() string`

GetSearchInterval returns the SearchInterval field if non-nil, zero value otherwise.

### GetSearchIntervalOk

`func (o *AddRecurringTaskRequest) GetSearchIntervalOk() (*string, bool)`

GetSearchIntervalOk returns a tuple with the SearchInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchInterval

`func (o *AddRecurringTaskRequest) SetSearchInterval(v string)`

SetSearchInterval sets SearchInterval field to given value.

### HasSearchInterval

`func (o *AddRecurringTaskRequest) HasSearchInterval() bool`

HasSearchInterval returns a boolean if a field has been set.

### GetSearchTimeLimit

`func (o *AddRecurringTaskRequest) GetSearchTimeLimit() string`

GetSearchTimeLimit returns the SearchTimeLimit field if non-nil, zero value otherwise.

### GetSearchTimeLimitOk

`func (o *AddRecurringTaskRequest) GetSearchTimeLimitOk() (*string, bool)`

GetSearchTimeLimitOk returns a tuple with the SearchTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTimeLimit

`func (o *AddRecurringTaskRequest) SetSearchTimeLimit(v string)`

SetSearchTimeLimit sets SearchTimeLimit field to given value.

### HasSearchTimeLimit

`func (o *AddRecurringTaskRequest) HasSearchTimeLimit() bool`

HasSearchTimeLimit returns a boolean if a field has been set.

### GetDurationToWaitForSearchToReturnEntries

`func (o *AddRecurringTaskRequest) GetDurationToWaitForSearchToReturnEntries() string`

GetDurationToWaitForSearchToReturnEntries returns the DurationToWaitForSearchToReturnEntries field if non-nil, zero value otherwise.

### GetDurationToWaitForSearchToReturnEntriesOk

`func (o *AddRecurringTaskRequest) GetDurationToWaitForSearchToReturnEntriesOk() (*string, bool)`

GetDurationToWaitForSearchToReturnEntriesOk returns a tuple with the DurationToWaitForSearchToReturnEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationToWaitForSearchToReturnEntries

`func (o *AddRecurringTaskRequest) SetDurationToWaitForSearchToReturnEntries(v string)`

SetDurationToWaitForSearchToReturnEntries sets DurationToWaitForSearchToReturnEntries field to given value.

### HasDurationToWaitForSearchToReturnEntries

`func (o *AddRecurringTaskRequest) HasDurationToWaitForSearchToReturnEntries() bool`

HasDurationToWaitForSearchToReturnEntries returns a boolean if a field has been set.

### GetTaskReturnStateIfTimeoutIsEncountered

`func (o *AddRecurringTaskRequest) GetTaskReturnStateIfTimeoutIsEncountered() EnumrecurringTaskTaskReturnStateIfTimeoutIsEncounteredProp`

GetTaskReturnStateIfTimeoutIsEncountered returns the TaskReturnStateIfTimeoutIsEncountered field if non-nil, zero value otherwise.

### GetTaskReturnStateIfTimeoutIsEncounteredOk

`func (o *AddRecurringTaskRequest) GetTaskReturnStateIfTimeoutIsEncounteredOk() (*EnumrecurringTaskTaskReturnStateIfTimeoutIsEncounteredProp, bool)`

GetTaskReturnStateIfTimeoutIsEncounteredOk returns a tuple with the TaskReturnStateIfTimeoutIsEncountered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskReturnStateIfTimeoutIsEncountered

`func (o *AddRecurringTaskRequest) SetTaskReturnStateIfTimeoutIsEncountered(v EnumrecurringTaskTaskReturnStateIfTimeoutIsEncounteredProp)`

SetTaskReturnStateIfTimeoutIsEncountered sets TaskReturnStateIfTimeoutIsEncountered field to given value.

### HasTaskReturnStateIfTimeoutIsEncountered

`func (o *AddRecurringTaskRequest) HasTaskReturnStateIfTimeoutIsEncountered() bool`

HasTaskReturnStateIfTimeoutIsEncountered returns a boolean if a field has been set.

### GetTaskJavaClass

`func (o *AddRecurringTaskRequest) GetTaskJavaClass() string`

GetTaskJavaClass returns the TaskJavaClass field if non-nil, zero value otherwise.

### GetTaskJavaClassOk

`func (o *AddRecurringTaskRequest) GetTaskJavaClassOk() (*string, bool)`

GetTaskJavaClassOk returns a tuple with the TaskJavaClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskJavaClass

`func (o *AddRecurringTaskRequest) SetTaskJavaClass(v string)`

SetTaskJavaClass sets TaskJavaClass field to given value.


### GetTaskObjectClass

`func (o *AddRecurringTaskRequest) GetTaskObjectClass() []string`

GetTaskObjectClass returns the TaskObjectClass field if non-nil, zero value otherwise.

### GetTaskObjectClassOk

`func (o *AddRecurringTaskRequest) GetTaskObjectClassOk() (*[]string, bool)`

GetTaskObjectClassOk returns a tuple with the TaskObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskObjectClass

`func (o *AddRecurringTaskRequest) SetTaskObjectClass(v []string)`

SetTaskObjectClass sets TaskObjectClass field to given value.


### GetTaskAttributeValue

`func (o *AddRecurringTaskRequest) GetTaskAttributeValue() []string`

GetTaskAttributeValue returns the TaskAttributeValue field if non-nil, zero value otherwise.

### GetTaskAttributeValueOk

`func (o *AddRecurringTaskRequest) GetTaskAttributeValueOk() (*[]string, bool)`

GetTaskAttributeValueOk returns a tuple with the TaskAttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskAttributeValue

`func (o *AddRecurringTaskRequest) SetTaskAttributeValue(v []string)`

SetTaskAttributeValue sets TaskAttributeValue field to given value.

### HasTaskAttributeValue

`func (o *AddRecurringTaskRequest) HasTaskAttributeValue() bool`

HasTaskAttributeValue returns a boolean if a field has been set.

### GetOutputDirectory

`func (o *AddRecurringTaskRequest) GetOutputDirectory() string`

GetOutputDirectory returns the OutputDirectory field if non-nil, zero value otherwise.

### GetOutputDirectoryOk

`func (o *AddRecurringTaskRequest) GetOutputDirectoryOk() (*string, bool)`

GetOutputDirectoryOk returns a tuple with the OutputDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDirectory

`func (o *AddRecurringTaskRequest) SetOutputDirectory(v string)`

SetOutputDirectory sets OutputDirectory field to given value.


### GetEncryptionPassphraseFile

`func (o *AddRecurringTaskRequest) GetEncryptionPassphraseFile() string`

GetEncryptionPassphraseFile returns the EncryptionPassphraseFile field if non-nil, zero value otherwise.

### GetEncryptionPassphraseFileOk

`func (o *AddRecurringTaskRequest) GetEncryptionPassphraseFileOk() (*string, bool)`

GetEncryptionPassphraseFileOk returns a tuple with the EncryptionPassphraseFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionPassphraseFile

`func (o *AddRecurringTaskRequest) SetEncryptionPassphraseFile(v string)`

SetEncryptionPassphraseFile sets EncryptionPassphraseFile field to given value.

### HasEncryptionPassphraseFile

`func (o *AddRecurringTaskRequest) HasEncryptionPassphraseFile() bool`

HasEncryptionPassphraseFile returns a boolean if a field has been set.

### GetIncludeExpensiveData

`func (o *AddRecurringTaskRequest) GetIncludeExpensiveData() bool`

GetIncludeExpensiveData returns the IncludeExpensiveData field if non-nil, zero value otherwise.

### GetIncludeExpensiveDataOk

`func (o *AddRecurringTaskRequest) GetIncludeExpensiveDataOk() (*bool, bool)`

GetIncludeExpensiveDataOk returns a tuple with the IncludeExpensiveData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeExpensiveData

`func (o *AddRecurringTaskRequest) SetIncludeExpensiveData(v bool)`

SetIncludeExpensiveData sets IncludeExpensiveData field to given value.

### HasIncludeExpensiveData

`func (o *AddRecurringTaskRequest) HasIncludeExpensiveData() bool`

HasIncludeExpensiveData returns a boolean if a field has been set.

### GetIncludeReplicationStateDump

`func (o *AddRecurringTaskRequest) GetIncludeReplicationStateDump() bool`

GetIncludeReplicationStateDump returns the IncludeReplicationStateDump field if non-nil, zero value otherwise.

### GetIncludeReplicationStateDumpOk

`func (o *AddRecurringTaskRequest) GetIncludeReplicationStateDumpOk() (*bool, bool)`

GetIncludeReplicationStateDumpOk returns a tuple with the IncludeReplicationStateDump field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeReplicationStateDump

`func (o *AddRecurringTaskRequest) SetIncludeReplicationStateDump(v bool)`

SetIncludeReplicationStateDump sets IncludeReplicationStateDump field to given value.

### HasIncludeReplicationStateDump

`func (o *AddRecurringTaskRequest) HasIncludeReplicationStateDump() bool`

HasIncludeReplicationStateDump returns a boolean if a field has been set.

### GetIncludeBinaryFiles

`func (o *AddRecurringTaskRequest) GetIncludeBinaryFiles() bool`

GetIncludeBinaryFiles returns the IncludeBinaryFiles field if non-nil, zero value otherwise.

### GetIncludeBinaryFilesOk

`func (o *AddRecurringTaskRequest) GetIncludeBinaryFilesOk() (*bool, bool)`

GetIncludeBinaryFilesOk returns a tuple with the IncludeBinaryFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBinaryFiles

`func (o *AddRecurringTaskRequest) SetIncludeBinaryFiles(v bool)`

SetIncludeBinaryFiles sets IncludeBinaryFiles field to given value.

### HasIncludeBinaryFiles

`func (o *AddRecurringTaskRequest) HasIncludeBinaryFiles() bool`

HasIncludeBinaryFiles returns a boolean if a field has been set.

### GetIncludeExtensionSource

`func (o *AddRecurringTaskRequest) GetIncludeExtensionSource() bool`

GetIncludeExtensionSource returns the IncludeExtensionSource field if non-nil, zero value otherwise.

### GetIncludeExtensionSourceOk

`func (o *AddRecurringTaskRequest) GetIncludeExtensionSourceOk() (*bool, bool)`

GetIncludeExtensionSourceOk returns a tuple with the IncludeExtensionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeExtensionSource

`func (o *AddRecurringTaskRequest) SetIncludeExtensionSource(v bool)`

SetIncludeExtensionSource sets IncludeExtensionSource field to given value.

### HasIncludeExtensionSource

`func (o *AddRecurringTaskRequest) HasIncludeExtensionSource() bool`

HasIncludeExtensionSource returns a boolean if a field has been set.

### GetUseSequentialMode

`func (o *AddRecurringTaskRequest) GetUseSequentialMode() bool`

GetUseSequentialMode returns the UseSequentialMode field if non-nil, zero value otherwise.

### GetUseSequentialModeOk

`func (o *AddRecurringTaskRequest) GetUseSequentialModeOk() (*bool, bool)`

GetUseSequentialModeOk returns a tuple with the UseSequentialMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSequentialMode

`func (o *AddRecurringTaskRequest) SetUseSequentialMode(v bool)`

SetUseSequentialMode sets UseSequentialMode field to given value.

### HasUseSequentialMode

`func (o *AddRecurringTaskRequest) HasUseSequentialMode() bool`

HasUseSequentialMode returns a boolean if a field has been set.

### GetSecurityLevel

`func (o *AddRecurringTaskRequest) GetSecurityLevel() EnumrecurringTaskSecurityLevelProp`

GetSecurityLevel returns the SecurityLevel field if non-nil, zero value otherwise.

### GetSecurityLevelOk

`func (o *AddRecurringTaskRequest) GetSecurityLevelOk() (*EnumrecurringTaskSecurityLevelProp, bool)`

GetSecurityLevelOk returns a tuple with the SecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLevel

`func (o *AddRecurringTaskRequest) SetSecurityLevel(v EnumrecurringTaskSecurityLevelProp)`

SetSecurityLevel sets SecurityLevel field to given value.

### HasSecurityLevel

`func (o *AddRecurringTaskRequest) HasSecurityLevel() bool`

HasSecurityLevel returns a boolean if a field has been set.

### GetJstackCount

`func (o *AddRecurringTaskRequest) GetJstackCount() int32`

GetJstackCount returns the JstackCount field if non-nil, zero value otherwise.

### GetJstackCountOk

`func (o *AddRecurringTaskRequest) GetJstackCountOk() (*int32, bool)`

GetJstackCountOk returns a tuple with the JstackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJstackCount

`func (o *AddRecurringTaskRequest) SetJstackCount(v int32)`

SetJstackCount sets JstackCount field to given value.

### HasJstackCount

`func (o *AddRecurringTaskRequest) HasJstackCount() bool`

HasJstackCount returns a boolean if a field has been set.

### GetReportCount

`func (o *AddRecurringTaskRequest) GetReportCount() int32`

GetReportCount returns the ReportCount field if non-nil, zero value otherwise.

### GetReportCountOk

`func (o *AddRecurringTaskRequest) GetReportCountOk() (*int32, bool)`

GetReportCountOk returns a tuple with the ReportCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportCount

`func (o *AddRecurringTaskRequest) SetReportCount(v int32)`

SetReportCount sets ReportCount field to given value.

### HasReportCount

`func (o *AddRecurringTaskRequest) HasReportCount() bool`

HasReportCount returns a boolean if a field has been set.

### GetReportIntervalSeconds

`func (o *AddRecurringTaskRequest) GetReportIntervalSeconds() int32`

GetReportIntervalSeconds returns the ReportIntervalSeconds field if non-nil, zero value otherwise.

### GetReportIntervalSecondsOk

`func (o *AddRecurringTaskRequest) GetReportIntervalSecondsOk() (*int32, bool)`

GetReportIntervalSecondsOk returns a tuple with the ReportIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportIntervalSeconds

`func (o *AddRecurringTaskRequest) SetReportIntervalSeconds(v int32)`

SetReportIntervalSeconds sets ReportIntervalSeconds field to given value.

### HasReportIntervalSeconds

`func (o *AddRecurringTaskRequest) HasReportIntervalSeconds() bool`

HasReportIntervalSeconds returns a boolean if a field has been set.

### GetLogDuration

`func (o *AddRecurringTaskRequest) GetLogDuration() string`

GetLogDuration returns the LogDuration field if non-nil, zero value otherwise.

### GetLogDurationOk

`func (o *AddRecurringTaskRequest) GetLogDurationOk() (*string, bool)`

GetLogDurationOk returns a tuple with the LogDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDuration

`func (o *AddRecurringTaskRequest) SetLogDuration(v string)`

SetLogDuration sets LogDuration field to given value.

### HasLogDuration

`func (o *AddRecurringTaskRequest) HasLogDuration() bool`

HasLogDuration returns a boolean if a field has been set.

### GetLogFileHeadCollectionSize

`func (o *AddRecurringTaskRequest) GetLogFileHeadCollectionSize() string`

GetLogFileHeadCollectionSize returns the LogFileHeadCollectionSize field if non-nil, zero value otherwise.

### GetLogFileHeadCollectionSizeOk

`func (o *AddRecurringTaskRequest) GetLogFileHeadCollectionSizeOk() (*string, bool)`

GetLogFileHeadCollectionSizeOk returns a tuple with the LogFileHeadCollectionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFileHeadCollectionSize

`func (o *AddRecurringTaskRequest) SetLogFileHeadCollectionSize(v string)`

SetLogFileHeadCollectionSize sets LogFileHeadCollectionSize field to given value.

### HasLogFileHeadCollectionSize

`func (o *AddRecurringTaskRequest) HasLogFileHeadCollectionSize() bool`

HasLogFileHeadCollectionSize returns a boolean if a field has been set.

### GetLogFileTailCollectionSize

`func (o *AddRecurringTaskRequest) GetLogFileTailCollectionSize() string`

GetLogFileTailCollectionSize returns the LogFileTailCollectionSize field if non-nil, zero value otherwise.

### GetLogFileTailCollectionSizeOk

`func (o *AddRecurringTaskRequest) GetLogFileTailCollectionSizeOk() (*string, bool)`

GetLogFileTailCollectionSizeOk returns a tuple with the LogFileTailCollectionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFileTailCollectionSize

`func (o *AddRecurringTaskRequest) SetLogFileTailCollectionSize(v string)`

SetLogFileTailCollectionSize sets LogFileTailCollectionSize field to given value.

### HasLogFileTailCollectionSize

`func (o *AddRecurringTaskRequest) HasLogFileTailCollectionSize() bool`

HasLogFileTailCollectionSize returns a boolean if a field has been set.

### GetComment

`func (o *AddRecurringTaskRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *AddRecurringTaskRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *AddRecurringTaskRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *AddRecurringTaskRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetRetainPreviousSupportDataArchiveCount

`func (o *AddRecurringTaskRequest) GetRetainPreviousSupportDataArchiveCount() int32`

GetRetainPreviousSupportDataArchiveCount returns the RetainPreviousSupportDataArchiveCount field if non-nil, zero value otherwise.

### GetRetainPreviousSupportDataArchiveCountOk

`func (o *AddRecurringTaskRequest) GetRetainPreviousSupportDataArchiveCountOk() (*int32, bool)`

GetRetainPreviousSupportDataArchiveCountOk returns a tuple with the RetainPreviousSupportDataArchiveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousSupportDataArchiveCount

`func (o *AddRecurringTaskRequest) SetRetainPreviousSupportDataArchiveCount(v int32)`

SetRetainPreviousSupportDataArchiveCount sets RetainPreviousSupportDataArchiveCount field to given value.

### HasRetainPreviousSupportDataArchiveCount

`func (o *AddRecurringTaskRequest) HasRetainPreviousSupportDataArchiveCount() bool`

HasRetainPreviousSupportDataArchiveCount returns a boolean if a field has been set.

### GetRetainPreviousSupportDataArchiveAge

`func (o *AddRecurringTaskRequest) GetRetainPreviousSupportDataArchiveAge() string`

GetRetainPreviousSupportDataArchiveAge returns the RetainPreviousSupportDataArchiveAge field if non-nil, zero value otherwise.

### GetRetainPreviousSupportDataArchiveAgeOk

`func (o *AddRecurringTaskRequest) GetRetainPreviousSupportDataArchiveAgeOk() (*string, bool)`

GetRetainPreviousSupportDataArchiveAgeOk returns a tuple with the RetainPreviousSupportDataArchiveAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousSupportDataArchiveAge

`func (o *AddRecurringTaskRequest) SetRetainPreviousSupportDataArchiveAge(v string)`

SetRetainPreviousSupportDataArchiveAge sets RetainPreviousSupportDataArchiveAge field to given value.

### HasRetainPreviousSupportDataArchiveAge

`func (o *AddRecurringTaskRequest) HasRetainPreviousSupportDataArchiveAge() bool`

HasRetainPreviousSupportDataArchiveAge returns a boolean if a field has been set.

### GetLdifDirectory

`func (o *AddRecurringTaskRequest) GetLdifDirectory() string`

GetLdifDirectory returns the LdifDirectory field if non-nil, zero value otherwise.

### GetLdifDirectoryOk

`func (o *AddRecurringTaskRequest) GetLdifDirectoryOk() (*string, bool)`

GetLdifDirectoryOk returns a tuple with the LdifDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdifDirectory

`func (o *AddRecurringTaskRequest) SetLdifDirectory(v string)`

SetLdifDirectory sets LdifDirectory field to given value.

### HasLdifDirectory

`func (o *AddRecurringTaskRequest) HasLdifDirectory() bool`

HasLdifDirectory returns a boolean if a field has been set.

### GetBackendID

`func (o *AddRecurringTaskRequest) GetBackendID() []string`

GetBackendID returns the BackendID field if non-nil, zero value otherwise.

### GetBackendIDOk

`func (o *AddRecurringTaskRequest) GetBackendIDOk() (*[]string, bool)`

GetBackendIDOk returns a tuple with the BackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendID

`func (o *AddRecurringTaskRequest) SetBackendID(v []string)`

SetBackendID sets BackendID field to given value.

### HasBackendID

`func (o *AddRecurringTaskRequest) HasBackendID() bool`

HasBackendID returns a boolean if a field has been set.

### GetExcludeBackendID

`func (o *AddRecurringTaskRequest) GetExcludeBackendID() []string`

GetExcludeBackendID returns the ExcludeBackendID field if non-nil, zero value otherwise.

### GetExcludeBackendIDOk

`func (o *AddRecurringTaskRequest) GetExcludeBackendIDOk() (*[]string, bool)`

GetExcludeBackendIDOk returns a tuple with the ExcludeBackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeBackendID

`func (o *AddRecurringTaskRequest) SetExcludeBackendID(v []string)`

SetExcludeBackendID sets ExcludeBackendID field to given value.

### HasExcludeBackendID

`func (o *AddRecurringTaskRequest) HasExcludeBackendID() bool`

HasExcludeBackendID returns a boolean if a field has been set.

### GetRetainPreviousLDIFExportCount

`func (o *AddRecurringTaskRequest) GetRetainPreviousLDIFExportCount() int32`

GetRetainPreviousLDIFExportCount returns the RetainPreviousLDIFExportCount field if non-nil, zero value otherwise.

### GetRetainPreviousLDIFExportCountOk

`func (o *AddRecurringTaskRequest) GetRetainPreviousLDIFExportCountOk() (*int32, bool)`

GetRetainPreviousLDIFExportCountOk returns a tuple with the RetainPreviousLDIFExportCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousLDIFExportCount

`func (o *AddRecurringTaskRequest) SetRetainPreviousLDIFExportCount(v int32)`

SetRetainPreviousLDIFExportCount sets RetainPreviousLDIFExportCount field to given value.

### HasRetainPreviousLDIFExportCount

`func (o *AddRecurringTaskRequest) HasRetainPreviousLDIFExportCount() bool`

HasRetainPreviousLDIFExportCount returns a boolean if a field has been set.

### GetRetainPreviousLDIFExportAge

`func (o *AddRecurringTaskRequest) GetRetainPreviousLDIFExportAge() string`

GetRetainPreviousLDIFExportAge returns the RetainPreviousLDIFExportAge field if non-nil, zero value otherwise.

### GetRetainPreviousLDIFExportAgeOk

`func (o *AddRecurringTaskRequest) GetRetainPreviousLDIFExportAgeOk() (*string, bool)`

GetRetainPreviousLDIFExportAgeOk returns a tuple with the RetainPreviousLDIFExportAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousLDIFExportAge

`func (o *AddRecurringTaskRequest) SetRetainPreviousLDIFExportAge(v string)`

SetRetainPreviousLDIFExportAge sets RetainPreviousLDIFExportAge field to given value.

### HasRetainPreviousLDIFExportAge

`func (o *AddRecurringTaskRequest) HasRetainPreviousLDIFExportAge() bool`

HasRetainPreviousLDIFExportAge returns a boolean if a field has been set.

### GetCommandPath

`func (o *AddRecurringTaskRequest) GetCommandPath() string`

GetCommandPath returns the CommandPath field if non-nil, zero value otherwise.

### GetCommandPathOk

`func (o *AddRecurringTaskRequest) GetCommandPathOk() (*string, bool)`

GetCommandPathOk returns a tuple with the CommandPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandPath

`func (o *AddRecurringTaskRequest) SetCommandPath(v string)`

SetCommandPath sets CommandPath field to given value.


### GetCommandArguments

`func (o *AddRecurringTaskRequest) GetCommandArguments() string`

GetCommandArguments returns the CommandArguments field if non-nil, zero value otherwise.

### GetCommandArgumentsOk

`func (o *AddRecurringTaskRequest) GetCommandArgumentsOk() (*string, bool)`

GetCommandArgumentsOk returns a tuple with the CommandArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandArguments

`func (o *AddRecurringTaskRequest) SetCommandArguments(v string)`

SetCommandArguments sets CommandArguments field to given value.

### HasCommandArguments

`func (o *AddRecurringTaskRequest) HasCommandArguments() bool`

HasCommandArguments returns a boolean if a field has been set.

### GetCommandOutputFileBaseName

`func (o *AddRecurringTaskRequest) GetCommandOutputFileBaseName() string`

GetCommandOutputFileBaseName returns the CommandOutputFileBaseName field if non-nil, zero value otherwise.

### GetCommandOutputFileBaseNameOk

`func (o *AddRecurringTaskRequest) GetCommandOutputFileBaseNameOk() (*string, bool)`

GetCommandOutputFileBaseNameOk returns a tuple with the CommandOutputFileBaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandOutputFileBaseName

`func (o *AddRecurringTaskRequest) SetCommandOutputFileBaseName(v string)`

SetCommandOutputFileBaseName sets CommandOutputFileBaseName field to given value.

### HasCommandOutputFileBaseName

`func (o *AddRecurringTaskRequest) HasCommandOutputFileBaseName() bool`

HasCommandOutputFileBaseName returns a boolean if a field has been set.

### GetRetainPreviousOutputFileCount

`func (o *AddRecurringTaskRequest) GetRetainPreviousOutputFileCount() int32`

GetRetainPreviousOutputFileCount returns the RetainPreviousOutputFileCount field if non-nil, zero value otherwise.

### GetRetainPreviousOutputFileCountOk

`func (o *AddRecurringTaskRequest) GetRetainPreviousOutputFileCountOk() (*int32, bool)`

GetRetainPreviousOutputFileCountOk returns a tuple with the RetainPreviousOutputFileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousOutputFileCount

`func (o *AddRecurringTaskRequest) SetRetainPreviousOutputFileCount(v int32)`

SetRetainPreviousOutputFileCount sets RetainPreviousOutputFileCount field to given value.

### HasRetainPreviousOutputFileCount

`func (o *AddRecurringTaskRequest) HasRetainPreviousOutputFileCount() bool`

HasRetainPreviousOutputFileCount returns a boolean if a field has been set.

### GetRetainPreviousOutputFileAge

`func (o *AddRecurringTaskRequest) GetRetainPreviousOutputFileAge() string`

GetRetainPreviousOutputFileAge returns the RetainPreviousOutputFileAge field if non-nil, zero value otherwise.

### GetRetainPreviousOutputFileAgeOk

`func (o *AddRecurringTaskRequest) GetRetainPreviousOutputFileAgeOk() (*string, bool)`

GetRetainPreviousOutputFileAgeOk returns a tuple with the RetainPreviousOutputFileAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousOutputFileAge

`func (o *AddRecurringTaskRequest) SetRetainPreviousOutputFileAge(v string)`

SetRetainPreviousOutputFileAge sets RetainPreviousOutputFileAge field to given value.

### HasRetainPreviousOutputFileAge

`func (o *AddRecurringTaskRequest) HasRetainPreviousOutputFileAge() bool`

HasRetainPreviousOutputFileAge returns a boolean if a field has been set.

### GetLogCommandOutput

`func (o *AddRecurringTaskRequest) GetLogCommandOutput() bool`

GetLogCommandOutput returns the LogCommandOutput field if non-nil, zero value otherwise.

### GetLogCommandOutputOk

`func (o *AddRecurringTaskRequest) GetLogCommandOutputOk() (*bool, bool)`

GetLogCommandOutputOk returns a tuple with the LogCommandOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogCommandOutput

`func (o *AddRecurringTaskRequest) SetLogCommandOutput(v bool)`

SetLogCommandOutput sets LogCommandOutput field to given value.

### HasLogCommandOutput

`func (o *AddRecurringTaskRequest) HasLogCommandOutput() bool`

HasLogCommandOutput returns a boolean if a field has been set.

### GetTaskCompletionStateForNonzeroExitCode

`func (o *AddRecurringTaskRequest) GetTaskCompletionStateForNonzeroExitCode() EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp`

GetTaskCompletionStateForNonzeroExitCode returns the TaskCompletionStateForNonzeroExitCode field if non-nil, zero value otherwise.

### GetTaskCompletionStateForNonzeroExitCodeOk

`func (o *AddRecurringTaskRequest) GetTaskCompletionStateForNonzeroExitCodeOk() (*EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp, bool)`

GetTaskCompletionStateForNonzeroExitCodeOk returns a tuple with the TaskCompletionStateForNonzeroExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCompletionStateForNonzeroExitCode

`func (o *AddRecurringTaskRequest) SetTaskCompletionStateForNonzeroExitCode(v EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp)`

SetTaskCompletionStateForNonzeroExitCode sets TaskCompletionStateForNonzeroExitCode field to given value.

### HasTaskCompletionStateForNonzeroExitCode

`func (o *AddRecurringTaskRequest) HasTaskCompletionStateForNonzeroExitCode() bool`

HasTaskCompletionStateForNonzeroExitCode returns a boolean if a field has been set.

### GetWorkingDirectory

`func (o *AddRecurringTaskRequest) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *AddRecurringTaskRequest) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *AddRecurringTaskRequest) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.

### HasWorkingDirectory

`func (o *AddRecurringTaskRequest) HasWorkingDirectory() bool`

HasWorkingDirectory returns a boolean if a field has been set.

### GetTargetDirectory

`func (o *AddRecurringTaskRequest) GetTargetDirectory() string`

GetTargetDirectory returns the TargetDirectory field if non-nil, zero value otherwise.

### GetTargetDirectoryOk

`func (o *AddRecurringTaskRequest) GetTargetDirectoryOk() (*string, bool)`

GetTargetDirectoryOk returns a tuple with the TargetDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDirectory

`func (o *AddRecurringTaskRequest) SetTargetDirectory(v string)`

SetTargetDirectory sets TargetDirectory field to given value.


### GetFilenamePattern

`func (o *AddRecurringTaskRequest) GetFilenamePattern() string`

GetFilenamePattern returns the FilenamePattern field if non-nil, zero value otherwise.

### GetFilenamePatternOk

`func (o *AddRecurringTaskRequest) GetFilenamePatternOk() (*string, bool)`

GetFilenamePatternOk returns a tuple with the FilenamePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilenamePattern

`func (o *AddRecurringTaskRequest) SetFilenamePattern(v string)`

SetFilenamePattern sets FilenamePattern field to given value.


### GetTimestampFormat

`func (o *AddRecurringTaskRequest) GetTimestampFormat() EnumrecurringTaskTimestampFormatProp`

GetTimestampFormat returns the TimestampFormat field if non-nil, zero value otherwise.

### GetTimestampFormatOk

`func (o *AddRecurringTaskRequest) GetTimestampFormatOk() (*EnumrecurringTaskTimestampFormatProp, bool)`

GetTimestampFormatOk returns a tuple with the TimestampFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampFormat

`func (o *AddRecurringTaskRequest) SetTimestampFormat(v EnumrecurringTaskTimestampFormatProp)`

SetTimestampFormat sets TimestampFormat field to given value.


### GetRetainFileCount

`func (o *AddRecurringTaskRequest) GetRetainFileCount() int32`

GetRetainFileCount returns the RetainFileCount field if non-nil, zero value otherwise.

### GetRetainFileCountOk

`func (o *AddRecurringTaskRequest) GetRetainFileCountOk() (*int32, bool)`

GetRetainFileCountOk returns a tuple with the RetainFileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainFileCount

`func (o *AddRecurringTaskRequest) SetRetainFileCount(v int32)`

SetRetainFileCount sets RetainFileCount field to given value.

### HasRetainFileCount

`func (o *AddRecurringTaskRequest) HasRetainFileCount() bool`

HasRetainFileCount returns a boolean if a field has been set.

### GetRetainFileAge

`func (o *AddRecurringTaskRequest) GetRetainFileAge() string`

GetRetainFileAge returns the RetainFileAge field if non-nil, zero value otherwise.

### GetRetainFileAgeOk

`func (o *AddRecurringTaskRequest) GetRetainFileAgeOk() (*string, bool)`

GetRetainFileAgeOk returns a tuple with the RetainFileAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainFileAge

`func (o *AddRecurringTaskRequest) SetRetainFileAge(v string)`

SetRetainFileAge sets RetainFileAge field to given value.

### HasRetainFileAge

`func (o *AddRecurringTaskRequest) HasRetainFileAge() bool`

HasRetainFileAge returns a boolean if a field has been set.

### GetRetainAggregateFileSize

`func (o *AddRecurringTaskRequest) GetRetainAggregateFileSize() string`

GetRetainAggregateFileSize returns the RetainAggregateFileSize field if non-nil, zero value otherwise.

### GetRetainAggregateFileSizeOk

`func (o *AddRecurringTaskRequest) GetRetainAggregateFileSizeOk() (*string, bool)`

GetRetainAggregateFileSizeOk returns a tuple with the RetainAggregateFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainAggregateFileSize

`func (o *AddRecurringTaskRequest) SetRetainAggregateFileSize(v string)`

SetRetainAggregateFileSize sets RetainAggregateFileSize field to given value.

### HasRetainAggregateFileSize

`func (o *AddRecurringTaskRequest) HasRetainAggregateFileSize() bool`

HasRetainAggregateFileSize returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddRecurringTaskRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddRecurringTaskRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddRecurringTaskRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddRecurringTaskRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddRecurringTaskRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddRecurringTaskRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddRecurringTaskRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


