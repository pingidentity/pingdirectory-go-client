# RecurringTaskListResponseResourcesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Recurring Task | 
**Schemas** | [**[]EnumthirdPartyRecurringTaskSchemaUrn**](EnumthirdPartyRecurringTaskSchemaUrn.md) |  | 
**ProfileDirectory** | **string** | The directory in which the generated server profiles will be placed. The files will be named with the pattern \&quot;server-profile-{timestamp}.zip\&quot;, where \&quot;{timestamp}\&quot; represents the time that the profile was generated. | 
**IncludePath** | Pointer to **[]string** | An optional set of additional paths to files within the instance root that should be included in the generated server profile. All paths must be within the instance root, and relative paths will be relative to the instance root. | [optional] 
**RetainPreviousProfileCount** | Pointer to **int64** | The minimum number of previous server profile zip files that should be preserved after a new profile is generated. | [optional] 
**RetainPreviousProfileAge** | Pointer to **string** | The minimum age of previous server profile zip files that should be preserved after a new profile is generated. | [optional] 
**Description** | Pointer to **string** | A description for this Recurring Task | [optional] 
**CancelOnTaskDependencyFailure** | Pointer to **bool** | Indicates whether an instance of this Recurring Task should be canceled if the task immediately before it in the recurring task chain fails to complete successfully (including if it is canceled by an administrator before it starts or while it is running). | [optional] 
**EmailOnStart** | Pointer to **[]string** | The email addresses to which a message should be sent whenever an instance of this Recurring Task starts running. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**EmailOnSuccess** | Pointer to **[]string** | The email addresses to which a message should be sent whenever an instance of this Recurring Task completes successfully. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**EmailOnFailure** | Pointer to **[]string** | The email addresses to which a message should be sent if an instance of this Recurring Task fails to complete successfully. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**AlertOnStart** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task starts running. | [optional] 
**AlertOnSuccess** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task completes successfully. | [optional] 
**AlertOnFailure** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task fails to complete successfully. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Reason** | Pointer to **string** | The reason that the server is being placed in lockdown mode. | [optional] 
**BackupDirectory** | **string** | The directory in which backup files will be placed. When backing up a single backend, the backup files will be placed directly in this directory. When backing up multiple backends, the backup files for each backend will be placed in a subdirectory whose name is the corresponding backend ID. | 
**IncludedBackendID** | Pointer to **[]string** | The backend IDs of any backends that should be included in the backup. | [optional] 
**ExcludedBackendID** | Pointer to **[]string** | The backend IDs of any backends that should be excluded from the backup. All backends that support backups and are not listed will be included. | [optional] 
**Compress** | Pointer to **bool** | Indicates whether to compress the LDIF data as it is exported. | [optional] 
**Encrypt** | Pointer to **bool** | Indicates whether to encrypt the LDIF data as it exported. | [optional] 
**EncryptionSettingsDefinitionID** | Pointer to **string** | The ID of an encryption settings definition to use to obtain the LDIF export encryption key. | [optional] 
**Sign** | Pointer to **bool** | Indicates whether to cryptographically sign the exported data, which will make it possible to detect whether the LDIF data has been altered since it was exported. | [optional] 
**RetainPreviousFullBackupCount** | Pointer to **int64** | The minimum number of previous full backups that should be preserved after a new backup completes successfully. | [optional] 
**RetainPreviousFullBackupAge** | Pointer to **string** | The minimum age of previous full backups that should be preserved after a new backup completes successfully. | [optional] 
**MaxMegabytesPerSecond** | Pointer to **int64** | The maximum rate, in megabytes per second, at which LDIF exports should be written. | [optional] 
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
**JstackCount** | Pointer to **int64** | The number of times to invoke the jstack utility to obtain a stack trace of all threads running in the JVM. A value of zero indicates that the jstack utility should not be invoked. | [optional] 
**ReportCount** | Pointer to **int64** | The number of intervals of data to collect from tools that use sample-based reporting, like vmstat, iostat, and mpstat. A value of zero indicates that these kinds of tools should not be used to collect any information. | [optional] 
**ReportIntervalSeconds** | Pointer to **int64** | The duration (in seconds) between each interval of data to collect from tools that use sample-based reporting, like vmstat, iostat, and mpstat. | [optional] 
**LogDuration** | Pointer to **string** | The maximum age (leading up to the time the collect-support-data tool was invoked) for log content to include in the support data archive. | [optional] 
**LogFileHeadCollectionSize** | Pointer to **string** | The amount of data to collect from the beginning of each log file included in the support data archive. | [optional] 
**LogFileTailCollectionSize** | Pointer to **string** | The amount of data to collect from the end of each log file included in the support data archive. | [optional] 
**Comment** | Pointer to **string** | An optional comment to include in a README file within the support data archive. | [optional] 
**RetainPreviousSupportDataArchiveCount** | Pointer to **int64** | The minimum number of previous support data archives that should be preserved after a new archive is generated. | [optional] 
**RetainPreviousSupportDataArchiveAge** | Pointer to **string** | The minimum age of previous support data archives that should be preserved after a new archive is generated. | [optional] 
**BaseOutputDirectory** | **string** | The base directory below which generated reports will be written. Each invocation of the audit-data-security task will create a new subdirectory below this base directory whose name is a timestamp indicating when the report was generated. | 
**DataSecurityAuditor** | Pointer to **[]string** | The set of data security auditors that should be invoked. If no auditors are specified, then all auditors defined in the configuration will be used. | [optional] 
**Backend** | Pointer to **[]string** | The set of backends that should be examined. If no backends are specified, then all backends that support this functionality will be included. | [optional] 
**IncludeFilter** | Pointer to **[]string** | A filter that will be used to identify entries that may be included in the generated report. If multiple filters are specified, then any entry that matches at least one of the filters will be included. If no filters are specified, then all entries will be included. | [optional] 
**RetainPreviousReportCount** | Pointer to **int64** | The minimum number of previous reports that should be preserved after a new report is generated. | [optional] 
**RetainPreviousReportAge** | Pointer to **string** | The minimum age of previous reports that should be preserved after a new report completes successfully. | [optional] 
**SleepDuration** | Pointer to **string** | The length of time to sleep before the task completes. | [optional] 
**DurationToWaitForWorkQueueIdle** | Pointer to **string** | Indicates that task should wait for up to the specified length of time for the work queue to report that all worker threads are idle and there are no pending operations. Note that this primarily monitors operations that use worker threads, which does not include internal operations (for example, those invoked by extensions), and may not include requests from non-LDAP clients (for example, HTTP-based clients). | [optional] 
**LdapURLForSearchExpectedToReturnEntries** | Pointer to **[]string** | An LDAP URL that provides the criteria for a search request that is expected to return at least one entry. The search will be performed internally, and only the base DN, scope, and filter from the URL will be used; any host, port, or requested attributes included in the URL will be ignored. | [optional] 
**SearchInterval** | Pointer to **string** | The length of time the server should sleep between searches performed using the criteria from the ldap-url-for-search-expected-to-return-entries property. | [optional] 
**SearchTimeLimit** | Pointer to **string** | The length of time that the server will wait for a response to each internal search performed using the criteria from the ldap-url-for-search-expected-to-return-entries property. | [optional] 
**DurationToWaitForSearchToReturnEntries** | Pointer to **string** | The maximum length of time that the server will continue to perform internal searches using the criteria from the ldap-url-for-search-expected-to-return-entries property. | [optional] 
**TaskReturnStateIfTimeoutIsEncountered** | Pointer to [**EnumrecurringTaskTaskReturnStateIfTimeoutIsEncounteredProp**](EnumrecurringTaskTaskReturnStateIfTimeoutIsEncounteredProp.md) |  | [optional] 
**LdifDirectory** | **string** | The directory in which LDIF export files will be placed. The directory must already exist. | 
**BackendID** | Pointer to **[]string** | The backend ID for a backend to be exported. | [optional] 
**ExcludeBackendID** | Pointer to **[]string** | The backend ID for a backend to be excluded from the export. | [optional] 
**RetainPreviousLDIFExportCount** | Pointer to **int64** | The minimum number of previous LDIF exports that should be preserved after a new export completes successfully. | [optional] 
**RetainPreviousLDIFExportAge** | Pointer to **string** | The minimum age of previous LDIF exports that should be preserved after a new export completes successfully. | [optional] 
**PostLDIFExportTaskProcessor** | Pointer to **[]string** | An optional set of post-LDIF-export task processors that should be invoked for the resulting LDIF export files. | [optional] 
**CommandPath** | **string** | The absolute path to the command to execute. It must be an absolute path, the corresponding file must exist, and it must be listed in the config/exec-command-whitelist.txt file. | 
**CommandArguments** | Pointer to **string** | A string containing the arguments to provide to the command. If the command should be run without arguments, this property should be left undefined. If there should be multiple arguments, then they should be separated with spaces. | [optional] 
**CommandOutputFileBaseName** | Pointer to **string** | The path and base name for a file to which the command output (both standard output and standard error) should be written. This may be left undefined if the command output should not be recorded into a file. | [optional] 
**RetainPreviousOutputFileCount** | Pointer to **int64** | The minimum number of previous command output files that should be preserved after a new instance of the command is invoked. | [optional] 
**RetainPreviousOutputFileAge** | Pointer to **string** | The minimum age of previous command output files that should be preserved after a new instance of the command is invoked. | [optional] 
**LogCommandOutput** | Pointer to **bool** | Indicates whether the command&#39;s output (both standard output and standard error) should be recorded in the server&#39;s error log. | [optional] 
**TaskCompletionStateForNonzeroExitCode** | Pointer to [**EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp**](EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp.md) |  | [optional] 
**WorkingDirectory** | Pointer to **string** | The absolute path to a working directory where the command should be executed. It must be an absolute path and the corresponding directory must exist. | [optional] 
**TargetDirectory** | **string** | The path to the directory containing the files to examine. The directory must exist. | 
**FilenamePattern** | **string** | A pattern that specifies the names of the files to examine. The pattern may contain zero or more asterisks as wildcards, where each wildcard matches zero or more characters. It may also contain at most one occurrence of the special string \&quot;${timestamp}\&quot;, which will match a timestamp with the format specified using the timestamp-format property. All other characters in the pattern will be treated literally. | 
**TimestampFormat** | [**EnumrecurringTaskTimestampFormatProp**](EnumrecurringTaskTimestampFormatProp.md) |  | 
**RetainFileCount** | Pointer to **int64** | The minimum number of files matching the pattern that will be retained. | [optional] 
**RetainFileAge** | Pointer to **string** | The minimum age of files matching the pattern that will be retained. | [optional] 
**RetainAggregateFileSize** | Pointer to **string** | The minimum aggregate size of files that will be retained. The size should be specified as an integer followed by a unit that is one of \&quot;b\&quot; or \&quot;bytes\&quot;, \&quot;kb\&quot; or \&quot;kilobytes\&quot;, \&quot;mb\&quot; or \&quot;megabytes\&quot;, \&quot;gb\&quot; or \&quot;gigabytes\&quot;, or \&quot;tb\&quot; or \&quot;terabytes\&quot;. For example, a value of \&quot;1 gb\&quot; indicates that at least one gigabyte of files should be retained. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Recurring Task. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Recurring Task. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewRecurringTaskListResponseResourcesInner

`func NewRecurringTaskListResponseResourcesInner(id string, schemas []EnumthirdPartyRecurringTaskSchemaUrn, profileDirectory string, backupDirectory string, taskJavaClass string, taskObjectClass []string, outputDirectory string, baseOutputDirectory string, ldifDirectory string, commandPath string, targetDirectory string, filenamePattern string, timestampFormat EnumrecurringTaskTimestampFormatProp, extensionClass string, ) *RecurringTaskListResponseResourcesInner`

NewRecurringTaskListResponseResourcesInner instantiates a new RecurringTaskListResponseResourcesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurringTaskListResponseResourcesInnerWithDefaults

`func NewRecurringTaskListResponseResourcesInnerWithDefaults() *RecurringTaskListResponseResourcesInner`

NewRecurringTaskListResponseResourcesInnerWithDefaults instantiates a new RecurringTaskListResponseResourcesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RecurringTaskListResponseResourcesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RecurringTaskListResponseResourcesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RecurringTaskListResponseResourcesInner) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *RecurringTaskListResponseResourcesInner) GetSchemas() []EnumthirdPartyRecurringTaskSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *RecurringTaskListResponseResourcesInner) GetSchemasOk() (*[]EnumthirdPartyRecurringTaskSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *RecurringTaskListResponseResourcesInner) SetSchemas(v []EnumthirdPartyRecurringTaskSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetProfileDirectory

`func (o *RecurringTaskListResponseResourcesInner) GetProfileDirectory() string`

GetProfileDirectory returns the ProfileDirectory field if non-nil, zero value otherwise.

### GetProfileDirectoryOk

`func (o *RecurringTaskListResponseResourcesInner) GetProfileDirectoryOk() (*string, bool)`

GetProfileDirectoryOk returns a tuple with the ProfileDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileDirectory

`func (o *RecurringTaskListResponseResourcesInner) SetProfileDirectory(v string)`

SetProfileDirectory sets ProfileDirectory field to given value.


### GetIncludePath

`func (o *RecurringTaskListResponseResourcesInner) GetIncludePath() []string`

GetIncludePath returns the IncludePath field if non-nil, zero value otherwise.

### GetIncludePathOk

`func (o *RecurringTaskListResponseResourcesInner) GetIncludePathOk() (*[]string, bool)`

GetIncludePathOk returns a tuple with the IncludePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePath

`func (o *RecurringTaskListResponseResourcesInner) SetIncludePath(v []string)`

SetIncludePath sets IncludePath field to given value.

### HasIncludePath

`func (o *RecurringTaskListResponseResourcesInner) HasIncludePath() bool`

HasIncludePath returns a boolean if a field has been set.

### GetRetainPreviousProfileCount

`func (o *RecurringTaskListResponseResourcesInner) GetRetainPreviousProfileCount() int64`

GetRetainPreviousProfileCount returns the RetainPreviousProfileCount field if non-nil, zero value otherwise.

### GetRetainPreviousProfileCountOk

`func (o *RecurringTaskListResponseResourcesInner) GetRetainPreviousProfileCountOk() (*int64, bool)`

GetRetainPreviousProfileCountOk returns a tuple with the RetainPreviousProfileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousProfileCount

`func (o *RecurringTaskListResponseResourcesInner) SetRetainPreviousProfileCount(v int64)`

SetRetainPreviousProfileCount sets RetainPreviousProfileCount field to given value.

### HasRetainPreviousProfileCount

`func (o *RecurringTaskListResponseResourcesInner) HasRetainPreviousProfileCount() bool`

HasRetainPreviousProfileCount returns a boolean if a field has been set.

### GetRetainPreviousProfileAge

`func (o *RecurringTaskListResponseResourcesInner) GetRetainPreviousProfileAge() string`

GetRetainPreviousProfileAge returns the RetainPreviousProfileAge field if non-nil, zero value otherwise.

### GetRetainPreviousProfileAgeOk

`func (o *RecurringTaskListResponseResourcesInner) GetRetainPreviousProfileAgeOk() (*string, bool)`

GetRetainPreviousProfileAgeOk returns a tuple with the RetainPreviousProfileAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousProfileAge

`func (o *RecurringTaskListResponseResourcesInner) SetRetainPreviousProfileAge(v string)`

SetRetainPreviousProfileAge sets RetainPreviousProfileAge field to given value.

### HasRetainPreviousProfileAge

`func (o *RecurringTaskListResponseResourcesInner) HasRetainPreviousProfileAge() bool`

HasRetainPreviousProfileAge returns a boolean if a field has been set.

### GetDescription

`func (o *RecurringTaskListResponseResourcesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RecurringTaskListResponseResourcesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RecurringTaskListResponseResourcesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RecurringTaskListResponseResourcesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCancelOnTaskDependencyFailure

`func (o *RecurringTaskListResponseResourcesInner) GetCancelOnTaskDependencyFailure() bool`

GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field if non-nil, zero value otherwise.

### GetCancelOnTaskDependencyFailureOk

`func (o *RecurringTaskListResponseResourcesInner) GetCancelOnTaskDependencyFailureOk() (*bool, bool)`

GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOnTaskDependencyFailure

`func (o *RecurringTaskListResponseResourcesInner) SetCancelOnTaskDependencyFailure(v bool)`

SetCancelOnTaskDependencyFailure sets CancelOnTaskDependencyFailure field to given value.

### HasCancelOnTaskDependencyFailure

`func (o *RecurringTaskListResponseResourcesInner) HasCancelOnTaskDependencyFailure() bool`

HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.

### GetEmailOnStart

`func (o *RecurringTaskListResponseResourcesInner) GetEmailOnStart() []string`

GetEmailOnStart returns the EmailOnStart field if non-nil, zero value otherwise.

### GetEmailOnStartOk

`func (o *RecurringTaskListResponseResourcesInner) GetEmailOnStartOk() (*[]string, bool)`

GetEmailOnStartOk returns a tuple with the EmailOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnStart

`func (o *RecurringTaskListResponseResourcesInner) SetEmailOnStart(v []string)`

SetEmailOnStart sets EmailOnStart field to given value.

### HasEmailOnStart

`func (o *RecurringTaskListResponseResourcesInner) HasEmailOnStart() bool`

HasEmailOnStart returns a boolean if a field has been set.

### GetEmailOnSuccess

`func (o *RecurringTaskListResponseResourcesInner) GetEmailOnSuccess() []string`

GetEmailOnSuccess returns the EmailOnSuccess field if non-nil, zero value otherwise.

### GetEmailOnSuccessOk

`func (o *RecurringTaskListResponseResourcesInner) GetEmailOnSuccessOk() (*[]string, bool)`

GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnSuccess

`func (o *RecurringTaskListResponseResourcesInner) SetEmailOnSuccess(v []string)`

SetEmailOnSuccess sets EmailOnSuccess field to given value.

### HasEmailOnSuccess

`func (o *RecurringTaskListResponseResourcesInner) HasEmailOnSuccess() bool`

HasEmailOnSuccess returns a boolean if a field has been set.

### GetEmailOnFailure

`func (o *RecurringTaskListResponseResourcesInner) GetEmailOnFailure() []string`

GetEmailOnFailure returns the EmailOnFailure field if non-nil, zero value otherwise.

### GetEmailOnFailureOk

`func (o *RecurringTaskListResponseResourcesInner) GetEmailOnFailureOk() (*[]string, bool)`

GetEmailOnFailureOk returns a tuple with the EmailOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnFailure

`func (o *RecurringTaskListResponseResourcesInner) SetEmailOnFailure(v []string)`

SetEmailOnFailure sets EmailOnFailure field to given value.

### HasEmailOnFailure

`func (o *RecurringTaskListResponseResourcesInner) HasEmailOnFailure() bool`

HasEmailOnFailure returns a boolean if a field has been set.

### GetAlertOnStart

`func (o *RecurringTaskListResponseResourcesInner) GetAlertOnStart() bool`

GetAlertOnStart returns the AlertOnStart field if non-nil, zero value otherwise.

### GetAlertOnStartOk

`func (o *RecurringTaskListResponseResourcesInner) GetAlertOnStartOk() (*bool, bool)`

GetAlertOnStartOk returns a tuple with the AlertOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnStart

`func (o *RecurringTaskListResponseResourcesInner) SetAlertOnStart(v bool)`

SetAlertOnStart sets AlertOnStart field to given value.

### HasAlertOnStart

`func (o *RecurringTaskListResponseResourcesInner) HasAlertOnStart() bool`

HasAlertOnStart returns a boolean if a field has been set.

### GetAlertOnSuccess

`func (o *RecurringTaskListResponseResourcesInner) GetAlertOnSuccess() bool`

GetAlertOnSuccess returns the AlertOnSuccess field if non-nil, zero value otherwise.

### GetAlertOnSuccessOk

`func (o *RecurringTaskListResponseResourcesInner) GetAlertOnSuccessOk() (*bool, bool)`

GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnSuccess

`func (o *RecurringTaskListResponseResourcesInner) SetAlertOnSuccess(v bool)`

SetAlertOnSuccess sets AlertOnSuccess field to given value.

### HasAlertOnSuccess

`func (o *RecurringTaskListResponseResourcesInner) HasAlertOnSuccess() bool`

HasAlertOnSuccess returns a boolean if a field has been set.

### GetAlertOnFailure

`func (o *RecurringTaskListResponseResourcesInner) GetAlertOnFailure() bool`

GetAlertOnFailure returns the AlertOnFailure field if non-nil, zero value otherwise.

### GetAlertOnFailureOk

`func (o *RecurringTaskListResponseResourcesInner) GetAlertOnFailureOk() (*bool, bool)`

GetAlertOnFailureOk returns a tuple with the AlertOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFailure

`func (o *RecurringTaskListResponseResourcesInner) SetAlertOnFailure(v bool)`

SetAlertOnFailure sets AlertOnFailure field to given value.

### HasAlertOnFailure

`func (o *RecurringTaskListResponseResourcesInner) HasAlertOnFailure() bool`

HasAlertOnFailure returns a boolean if a field has been set.

### GetMeta

`func (o *RecurringTaskListResponseResourcesInner) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RecurringTaskListResponseResourcesInner) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RecurringTaskListResponseResourcesInner) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *RecurringTaskListResponseResourcesInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *RecurringTaskListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *RecurringTaskListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *RecurringTaskListResponseResourcesInner) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *RecurringTaskListResponseResourcesInner) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetReason

`func (o *RecurringTaskListResponseResourcesInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RecurringTaskListResponseResourcesInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RecurringTaskListResponseResourcesInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RecurringTaskListResponseResourcesInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetBackupDirectory

`func (o *RecurringTaskListResponseResourcesInner) GetBackupDirectory() string`

GetBackupDirectory returns the BackupDirectory field if non-nil, zero value otherwise.

### GetBackupDirectoryOk

`func (o *RecurringTaskListResponseResourcesInner) GetBackupDirectoryOk() (*string, bool)`

GetBackupDirectoryOk returns a tuple with the BackupDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDirectory

`func (o *RecurringTaskListResponseResourcesInner) SetBackupDirectory(v string)`

SetBackupDirectory sets BackupDirectory field to given value.


### GetIncludedBackendID

`func (o *RecurringTaskListResponseResourcesInner) GetIncludedBackendID() []string`

GetIncludedBackendID returns the IncludedBackendID field if non-nil, zero value otherwise.

### GetIncludedBackendIDOk

`func (o *RecurringTaskListResponseResourcesInner) GetIncludedBackendIDOk() (*[]string, bool)`

GetIncludedBackendIDOk returns a tuple with the IncludedBackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedBackendID

`func (o *RecurringTaskListResponseResourcesInner) SetIncludedBackendID(v []string)`

SetIncludedBackendID sets IncludedBackendID field to given value.

### HasIncludedBackendID

`func (o *RecurringTaskListResponseResourcesInner) HasIncludedBackendID() bool`

HasIncludedBackendID returns a boolean if a field has been set.

### GetExcludedBackendID

`func (o *RecurringTaskListResponseResourcesInner) GetExcludedBackendID() []string`

GetExcludedBackendID returns the ExcludedBackendID field if non-nil, zero value otherwise.

### GetExcludedBackendIDOk

`func (o *RecurringTaskListResponseResourcesInner) GetExcludedBackendIDOk() (*[]string, bool)`

GetExcludedBackendIDOk returns a tuple with the ExcludedBackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedBackendID

`func (o *RecurringTaskListResponseResourcesInner) SetExcludedBackendID(v []string)`

SetExcludedBackendID sets ExcludedBackendID field to given value.

### HasExcludedBackendID

`func (o *RecurringTaskListResponseResourcesInner) HasExcludedBackendID() bool`

HasExcludedBackendID returns a boolean if a field has been set.

### GetCompress

`func (o *RecurringTaskListResponseResourcesInner) GetCompress() bool`

GetCompress returns the Compress field if non-nil, zero value otherwise.

### GetCompressOk

`func (o *RecurringTaskListResponseResourcesInner) GetCompressOk() (*bool, bool)`

GetCompressOk returns a tuple with the Compress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompress

`func (o *RecurringTaskListResponseResourcesInner) SetCompress(v bool)`

SetCompress sets Compress field to given value.

### HasCompress

`func (o *RecurringTaskListResponseResourcesInner) HasCompress() bool`

HasCompress returns a boolean if a field has been set.

### GetEncrypt

`func (o *RecurringTaskListResponseResourcesInner) GetEncrypt() bool`

GetEncrypt returns the Encrypt field if non-nil, zero value otherwise.

### GetEncryptOk

`func (o *RecurringTaskListResponseResourcesInner) GetEncryptOk() (*bool, bool)`

GetEncryptOk returns a tuple with the Encrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypt

`func (o *RecurringTaskListResponseResourcesInner) SetEncrypt(v bool)`

SetEncrypt sets Encrypt field to given value.

### HasEncrypt

`func (o *RecurringTaskListResponseResourcesInner) HasEncrypt() bool`

HasEncrypt returns a boolean if a field has been set.

### GetEncryptionSettingsDefinitionID

`func (o *RecurringTaskListResponseResourcesInner) GetEncryptionSettingsDefinitionID() string`

GetEncryptionSettingsDefinitionID returns the EncryptionSettingsDefinitionID field if non-nil, zero value otherwise.

### GetEncryptionSettingsDefinitionIDOk

`func (o *RecurringTaskListResponseResourcesInner) GetEncryptionSettingsDefinitionIDOk() (*string, bool)`

GetEncryptionSettingsDefinitionIDOk returns a tuple with the EncryptionSettingsDefinitionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionSettingsDefinitionID

`func (o *RecurringTaskListResponseResourcesInner) SetEncryptionSettingsDefinitionID(v string)`

SetEncryptionSettingsDefinitionID sets EncryptionSettingsDefinitionID field to given value.

### HasEncryptionSettingsDefinitionID

`func (o *RecurringTaskListResponseResourcesInner) HasEncryptionSettingsDefinitionID() bool`

HasEncryptionSettingsDefinitionID returns a boolean if a field has been set.

### GetSign

`func (o *RecurringTaskListResponseResourcesInner) GetSign() bool`

GetSign returns the Sign field if non-nil, zero value otherwise.

### GetSignOk

`func (o *RecurringTaskListResponseResourcesInner) GetSignOk() (*bool, bool)`

GetSignOk returns a tuple with the Sign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSign

`func (o *RecurringTaskListResponseResourcesInner) SetSign(v bool)`

SetSign sets Sign field to given value.

### HasSign

`func (o *RecurringTaskListResponseResourcesInner) HasSign() bool`

HasSign returns a boolean if a field has been set.

### GetRetainPreviousFullBackupCount

`func (o *RecurringTaskListResponseResourcesInner) GetRetainPreviousFullBackupCount() int64`

GetRetainPreviousFullBackupCount returns the RetainPreviousFullBackupCount field if non-nil, zero value otherwise.

### GetRetainPreviousFullBackupCountOk

`func (o *RecurringTaskListResponseResourcesInner) GetRetainPreviousFullBackupCountOk() (*int64, bool)`

GetRetainPreviousFullBackupCountOk returns a tuple with the RetainPreviousFullBackupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousFullBackupCount

`func (o *RecurringTaskListResponseResourcesInner) SetRetainPreviousFullBackupCount(v int64)`

SetRetainPreviousFullBackupCount sets RetainPreviousFullBackupCount field to given value.

### HasRetainPreviousFullBackupCount

`func (o *RecurringTaskListResponseResourcesInner) HasRetainPreviousFullBackupCount() bool`

HasRetainPreviousFullBackupCount returns a boolean if a field has been set.

### GetRetainPreviousFullBackupAge

`func (o *RecurringTaskListResponseResourcesInner) GetRetainPreviousFullBackupAge() string`

GetRetainPreviousFullBackupAge returns the RetainPreviousFullBackupAge field if non-nil, zero value otherwise.

### GetRetainPreviousFullBackupAgeOk

`func (o *RecurringTaskListResponseResourcesInner) GetRetainPreviousFullBackupAgeOk() (*string, bool)`

GetRetainPreviousFullBackupAgeOk returns a tuple with the RetainPreviousFullBackupAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousFullBackupAge

`func (o *RecurringTaskListResponseResourcesInner) SetRetainPreviousFullBackupAge(v string)`

SetRetainPreviousFullBackupAge sets RetainPreviousFullBackupAge field to given value.

### HasRetainPreviousFullBackupAge

`func (o *RecurringTaskListResponseResourcesInner) HasRetainPreviousFullBackupAge() bool`

HasRetainPreviousFullBackupAge returns a boolean if a field has been set.

### GetMaxMegabytesPerSecond

`func (o *RecurringTaskListResponseResourcesInner) GetMaxMegabytesPerSecond() int64`

GetMaxMegabytesPerSecond returns the MaxMegabytesPerSecond field if non-nil, zero value otherwise.

### GetMaxMegabytesPerSecondOk

`func (o *RecurringTaskListResponseResourcesInner) GetMaxMegabytesPerSecondOk() (*int64, bool)`

GetMaxMegabytesPerSecondOk returns a tuple with the MaxMegabytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMegabytesPerSecond

`func (o *RecurringTaskListResponseResourcesInner) SetMaxMegabytesPerSecond(v int64)`

SetMaxMegabytesPerSecond sets MaxMegabytesPerSecond field to given value.

### HasMaxMegabytesPerSecond

`func (o *RecurringTaskListResponseResourcesInner) HasMaxMegabytesPerSecond() bool`

HasMaxMegabytesPerSecond returns a boolean if a field has been set.

### GetTaskJavaClass

`func (o *RecurringTaskListResponseResourcesInner) GetTaskJavaClass() string`

GetTaskJavaClass returns the TaskJavaClass field if non-nil, zero value otherwise.

### GetTaskJavaClassOk

`func (o *RecurringTaskListResponseResourcesInner) GetTaskJavaClassOk() (*string, bool)`

GetTaskJavaClassOk returns a tuple with the TaskJavaClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskJavaClass

`func (o *RecurringTaskListResponseResourcesInner) SetTaskJavaClass(v string)`

SetTaskJavaClass sets TaskJavaClass field to given value.


### GetTaskObjectClass

`func (o *RecurringTaskListResponseResourcesInner) GetTaskObjectClass() []string`

GetTaskObjectClass returns the TaskObjectClass field if non-nil, zero value otherwise.

### GetTaskObjectClassOk

`func (o *RecurringTaskListResponseResourcesInner) GetTaskObjectClassOk() (*[]string, bool)`

GetTaskObjectClassOk returns a tuple with the TaskObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskObjectClass

`func (o *RecurringTaskListResponseResourcesInner) SetTaskObjectClass(v []string)`

SetTaskObjectClass sets TaskObjectClass field to given value.


### GetTaskAttributeValue

`func (o *RecurringTaskListResponseResourcesInner) GetTaskAttributeValue() []string`

GetTaskAttributeValue returns the TaskAttributeValue field if non-nil, zero value otherwise.

### GetTaskAttributeValueOk

`func (o *RecurringTaskListResponseResourcesInner) GetTaskAttributeValueOk() (*[]string, bool)`

GetTaskAttributeValueOk returns a tuple with the TaskAttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskAttributeValue

`func (o *RecurringTaskListResponseResourcesInner) SetTaskAttributeValue(v []string)`

SetTaskAttributeValue sets TaskAttributeValue field to given value.

### HasTaskAttributeValue

`func (o *RecurringTaskListResponseResourcesInner) HasTaskAttributeValue() bool`

HasTaskAttributeValue returns a boolean if a field has been set.

### GetOutputDirectory

`func (o *RecurringTaskListResponseResourcesInner) GetOutputDirectory() string`

GetOutputDirectory returns the OutputDirectory field if non-nil, zero value otherwise.

### GetOutputDirectoryOk

`func (o *RecurringTaskListResponseResourcesInner) GetOutputDirectoryOk() (*string, bool)`

GetOutputDirectoryOk returns a tuple with the OutputDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDirectory

`func (o *RecurringTaskListResponseResourcesInner) SetOutputDirectory(v string)`

SetOutputDirectory sets OutputDirectory field to given value.


### GetEncryptionPassphraseFile

`func (o *RecurringTaskListResponseResourcesInner) GetEncryptionPassphraseFile() string`

GetEncryptionPassphraseFile returns the EncryptionPassphraseFile field if non-nil, zero value otherwise.

### GetEncryptionPassphraseFileOk

`func (o *RecurringTaskListResponseResourcesInner) GetEncryptionPassphraseFileOk() (*string, bool)`

GetEncryptionPassphraseFileOk returns a tuple with the EncryptionPassphraseFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionPassphraseFile

`func (o *RecurringTaskListResponseResourcesInner) SetEncryptionPassphraseFile(v string)`

SetEncryptionPassphraseFile sets EncryptionPassphraseFile field to given value.

### HasEncryptionPassphraseFile

`func (o *RecurringTaskListResponseResourcesInner) HasEncryptionPassphraseFile() bool`

HasEncryptionPassphraseFile returns a boolean if a field has been set.

### GetIncludeExpensiveData

`func (o *RecurringTaskListResponseResourcesInner) GetIncludeExpensiveData() bool`

GetIncludeExpensiveData returns the IncludeExpensiveData field if non-nil, zero value otherwise.

### GetIncludeExpensiveDataOk

`func (o *RecurringTaskListResponseResourcesInner) GetIncludeExpensiveDataOk() (*bool, bool)`

GetIncludeExpensiveDataOk returns a tuple with the IncludeExpensiveData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeExpensiveData

`func (o *RecurringTaskListResponseResourcesInner) SetIncludeExpensiveData(v bool)`

SetIncludeExpensiveData sets IncludeExpensiveData field to given value.

### HasIncludeExpensiveData

`func (o *RecurringTaskListResponseResourcesInner) HasIncludeExpensiveData() bool`

HasIncludeExpensiveData returns a boolean if a field has been set.

### GetIncludeReplicationStateDump

`func (o *RecurringTaskListResponseResourcesInner) GetIncludeReplicationStateDump() bool`

GetIncludeReplicationStateDump returns the IncludeReplicationStateDump field if non-nil, zero value otherwise.

### GetIncludeReplicationStateDumpOk

`func (o *RecurringTaskListResponseResourcesInner) GetIncludeReplicationStateDumpOk() (*bool, bool)`

GetIncludeReplicationStateDumpOk returns a tuple with the IncludeReplicationStateDump field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeReplicationStateDump

`func (o *RecurringTaskListResponseResourcesInner) SetIncludeReplicationStateDump(v bool)`

SetIncludeReplicationStateDump sets IncludeReplicationStateDump field to given value.

### HasIncludeReplicationStateDump

`func (o *RecurringTaskListResponseResourcesInner) HasIncludeReplicationStateDump() bool`

HasIncludeReplicationStateDump returns a boolean if a field has been set.

### GetIncludeBinaryFiles

`func (o *RecurringTaskListResponseResourcesInner) GetIncludeBinaryFiles() bool`

GetIncludeBinaryFiles returns the IncludeBinaryFiles field if non-nil, zero value otherwise.

### GetIncludeBinaryFilesOk

`func (o *RecurringTaskListResponseResourcesInner) GetIncludeBinaryFilesOk() (*bool, bool)`

GetIncludeBinaryFilesOk returns a tuple with the IncludeBinaryFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBinaryFiles

`func (o *RecurringTaskListResponseResourcesInner) SetIncludeBinaryFiles(v bool)`

SetIncludeBinaryFiles sets IncludeBinaryFiles field to given value.

### HasIncludeBinaryFiles

`func (o *RecurringTaskListResponseResourcesInner) HasIncludeBinaryFiles() bool`

HasIncludeBinaryFiles returns a boolean if a field has been set.

### GetIncludeExtensionSource

`func (o *RecurringTaskListResponseResourcesInner) GetIncludeExtensionSource() bool`

GetIncludeExtensionSource returns the IncludeExtensionSource field if non-nil, zero value otherwise.

### GetIncludeExtensionSourceOk

`func (o *RecurringTaskListResponseResourcesInner) GetIncludeExtensionSourceOk() (*bool, bool)`

GetIncludeExtensionSourceOk returns a tuple with the IncludeExtensionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeExtensionSource

`func (o *RecurringTaskListResponseResourcesInner) SetIncludeExtensionSource(v bool)`

SetIncludeExtensionSource sets IncludeExtensionSource field to given value.

### HasIncludeExtensionSource

`func (o *RecurringTaskListResponseResourcesInner) HasIncludeExtensionSource() bool`

HasIncludeExtensionSource returns a boolean if a field has been set.

### GetUseSequentialMode

`func (o *RecurringTaskListResponseResourcesInner) GetUseSequentialMode() bool`

GetUseSequentialMode returns the UseSequentialMode field if non-nil, zero value otherwise.

### GetUseSequentialModeOk

`func (o *RecurringTaskListResponseResourcesInner) GetUseSequentialModeOk() (*bool, bool)`

GetUseSequentialModeOk returns a tuple with the UseSequentialMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSequentialMode

`func (o *RecurringTaskListResponseResourcesInner) SetUseSequentialMode(v bool)`

SetUseSequentialMode sets UseSequentialMode field to given value.

### HasUseSequentialMode

`func (o *RecurringTaskListResponseResourcesInner) HasUseSequentialMode() bool`

HasUseSequentialMode returns a boolean if a field has been set.

### GetSecurityLevel

`func (o *RecurringTaskListResponseResourcesInner) GetSecurityLevel() EnumrecurringTaskSecurityLevelProp`

GetSecurityLevel returns the SecurityLevel field if non-nil, zero value otherwise.

### GetSecurityLevelOk

`func (o *RecurringTaskListResponseResourcesInner) GetSecurityLevelOk() (*EnumrecurringTaskSecurityLevelProp, bool)`

GetSecurityLevelOk returns a tuple with the SecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLevel

`func (o *RecurringTaskListResponseResourcesInner) SetSecurityLevel(v EnumrecurringTaskSecurityLevelProp)`

SetSecurityLevel sets SecurityLevel field to given value.

### HasSecurityLevel

`func (o *RecurringTaskListResponseResourcesInner) HasSecurityLevel() bool`

HasSecurityLevel returns a boolean if a field has been set.

### GetJstackCount

`func (o *RecurringTaskListResponseResourcesInner) GetJstackCount() int64`

GetJstackCount returns the JstackCount field if non-nil, zero value otherwise.

### GetJstackCountOk

`func (o *RecurringTaskListResponseResourcesInner) GetJstackCountOk() (*int64, bool)`

GetJstackCountOk returns a tuple with the JstackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJstackCount

`func (o *RecurringTaskListResponseResourcesInner) SetJstackCount(v int64)`

SetJstackCount sets JstackCount field to given value.

### HasJstackCount

`func (o *RecurringTaskListResponseResourcesInner) HasJstackCount() bool`

HasJstackCount returns a boolean if a field has been set.

### GetReportCount

`func (o *RecurringTaskListResponseResourcesInner) GetReportCount() int64`

GetReportCount returns the ReportCount field if non-nil, zero value otherwise.

### GetReportCountOk

`func (o *RecurringTaskListResponseResourcesInner) GetReportCountOk() (*int64, bool)`

GetReportCountOk returns a tuple with the ReportCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportCount

`func (o *RecurringTaskListResponseResourcesInner) SetReportCount(v int64)`

SetReportCount sets ReportCount field to given value.

### HasReportCount

`func (o *RecurringTaskListResponseResourcesInner) HasReportCount() bool`

HasReportCount returns a boolean if a field has been set.

### GetReportIntervalSeconds

`func (o *RecurringTaskListResponseResourcesInner) GetReportIntervalSeconds() int64`

GetReportIntervalSeconds returns the ReportIntervalSeconds field if non-nil, zero value otherwise.

### GetReportIntervalSecondsOk

`func (o *RecurringTaskListResponseResourcesInner) GetReportIntervalSecondsOk() (*int64, bool)`

GetReportIntervalSecondsOk returns a tuple with the ReportIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportIntervalSeconds

`func (o *RecurringTaskListResponseResourcesInner) SetReportIntervalSeconds(v int64)`

SetReportIntervalSeconds sets ReportIntervalSeconds field to given value.

### HasReportIntervalSeconds

`func (o *RecurringTaskListResponseResourcesInner) HasReportIntervalSeconds() bool`

HasReportIntervalSeconds returns a boolean if a field has been set.

### GetLogDuration

`func (o *RecurringTaskListResponseResourcesInner) GetLogDuration() string`

GetLogDuration returns the LogDuration field if non-nil, zero value otherwise.

### GetLogDurationOk

`func (o *RecurringTaskListResponseResourcesInner) GetLogDurationOk() (*string, bool)`

GetLogDurationOk returns a tuple with the LogDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDuration

`func (o *RecurringTaskListResponseResourcesInner) SetLogDuration(v string)`

SetLogDuration sets LogDuration field to given value.

### HasLogDuration

`func (o *RecurringTaskListResponseResourcesInner) HasLogDuration() bool`

HasLogDuration returns a boolean if a field has been set.

### GetLogFileHeadCollectionSize

`func (o *RecurringTaskListResponseResourcesInner) GetLogFileHeadCollectionSize() string`

GetLogFileHeadCollectionSize returns the LogFileHeadCollectionSize field if non-nil, zero value otherwise.

### GetLogFileHeadCollectionSizeOk

`func (o *RecurringTaskListResponseResourcesInner) GetLogFileHeadCollectionSizeOk() (*string, bool)`

GetLogFileHeadCollectionSizeOk returns a tuple with the LogFileHeadCollectionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFileHeadCollectionSize

`func (o *RecurringTaskListResponseResourcesInner) SetLogFileHeadCollectionSize(v string)`

SetLogFileHeadCollectionSize sets LogFileHeadCollectionSize field to given value.

### HasLogFileHeadCollectionSize

`func (o *RecurringTaskListResponseResourcesInner) HasLogFileHeadCollectionSize() bool`

HasLogFileHeadCollectionSize returns a boolean if a field has been set.

### GetLogFileTailCollectionSize

`func (o *RecurringTaskListResponseResourcesInner) GetLogFileTailCollectionSize() string`

GetLogFileTailCollectionSize returns the LogFileTailCollectionSize field if non-nil, zero value otherwise.

### GetLogFileTailCollectionSizeOk

`func (o *RecurringTaskListResponseResourcesInner) GetLogFileTailCollectionSizeOk() (*string, bool)`

GetLogFileTailCollectionSizeOk returns a tuple with the LogFileTailCollectionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFileTailCollectionSize

`func (o *RecurringTaskListResponseResourcesInner) SetLogFileTailCollectionSize(v string)`

SetLogFileTailCollectionSize sets LogFileTailCollectionSize field to given value.

### HasLogFileTailCollectionSize

`func (o *RecurringTaskListResponseResourcesInner) HasLogFileTailCollectionSize() bool`

HasLogFileTailCollectionSize returns a boolean if a field has been set.

### GetComment

`func (o *RecurringTaskListResponseResourcesInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RecurringTaskListResponseResourcesInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RecurringTaskListResponseResourcesInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RecurringTaskListResponseResourcesInner) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetRetainPreviousSupportDataArchiveCount

`func (o *RecurringTaskListResponseResourcesInner) GetRetainPreviousSupportDataArchiveCount() int64`

GetRetainPreviousSupportDataArchiveCount returns the RetainPreviousSupportDataArchiveCount field if non-nil, zero value otherwise.

### GetRetainPreviousSupportDataArchiveCountOk

`func (o *RecurringTaskListResponseResourcesInner) GetRetainPreviousSupportDataArchiveCountOk() (*int64, bool)`

GetRetainPreviousSupportDataArchiveCountOk returns a tuple with the RetainPreviousSupportDataArchiveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousSupportDataArchiveCount

`func (o *RecurringTaskListResponseResourcesInner) SetRetainPreviousSupportDataArchiveCount(v int64)`

SetRetainPreviousSupportDataArchiveCount sets RetainPreviousSupportDataArchiveCount field to given value.

### HasRetainPreviousSupportDataArchiveCount

`func (o *RecurringTaskListResponseResourcesInner) HasRetainPreviousSupportDataArchiveCount() bool`

HasRetainPreviousSupportDataArchiveCount returns a boolean if a field has been set.

### GetRetainPreviousSupportDataArchiveAge

`func (o *RecurringTaskListResponseResourcesInner) GetRetainPreviousSupportDataArchiveAge() string`

GetRetainPreviousSupportDataArchiveAge returns the RetainPreviousSupportDataArchiveAge field if non-nil, zero value otherwise.

### GetRetainPreviousSupportDataArchiveAgeOk

`func (o *RecurringTaskListResponseResourcesInner) GetRetainPreviousSupportDataArchiveAgeOk() (*string, bool)`

GetRetainPreviousSupportDataArchiveAgeOk returns a tuple with the RetainPreviousSupportDataArchiveAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousSupportDataArchiveAge

`func (o *RecurringTaskListResponseResourcesInner) SetRetainPreviousSupportDataArchiveAge(v string)`

SetRetainPreviousSupportDataArchiveAge sets RetainPreviousSupportDataArchiveAge field to given value.

### HasRetainPreviousSupportDataArchiveAge

`func (o *RecurringTaskListResponseResourcesInner) HasRetainPreviousSupportDataArchiveAge() bool`

HasRetainPreviousSupportDataArchiveAge returns a boolean if a field has been set.

### GetBaseOutputDirectory

`func (o *RecurringTaskListResponseResourcesInner) GetBaseOutputDirectory() string`

GetBaseOutputDirectory returns the BaseOutputDirectory field if non-nil, zero value otherwise.

### GetBaseOutputDirectoryOk

`func (o *RecurringTaskListResponseResourcesInner) GetBaseOutputDirectoryOk() (*string, bool)`

GetBaseOutputDirectoryOk returns a tuple with the BaseOutputDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseOutputDirectory

`func (o *RecurringTaskListResponseResourcesInner) SetBaseOutputDirectory(v string)`

SetBaseOutputDirectory sets BaseOutputDirectory field to given value.


### GetDataSecurityAuditor

`func (o *RecurringTaskListResponseResourcesInner) GetDataSecurityAuditor() []string`

GetDataSecurityAuditor returns the DataSecurityAuditor field if non-nil, zero value otherwise.

### GetDataSecurityAuditorOk

`func (o *RecurringTaskListResponseResourcesInner) GetDataSecurityAuditorOk() (*[]string, bool)`

GetDataSecurityAuditorOk returns a tuple with the DataSecurityAuditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSecurityAuditor

`func (o *RecurringTaskListResponseResourcesInner) SetDataSecurityAuditor(v []string)`

SetDataSecurityAuditor sets DataSecurityAuditor field to given value.

### HasDataSecurityAuditor

`func (o *RecurringTaskListResponseResourcesInner) HasDataSecurityAuditor() bool`

HasDataSecurityAuditor returns a boolean if a field has been set.

### GetBackend

`func (o *RecurringTaskListResponseResourcesInner) GetBackend() []string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *RecurringTaskListResponseResourcesInner) GetBackendOk() (*[]string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *RecurringTaskListResponseResourcesInner) SetBackend(v []string)`

SetBackend sets Backend field to given value.

### HasBackend

`func (o *RecurringTaskListResponseResourcesInner) HasBackend() bool`

HasBackend returns a boolean if a field has been set.

### GetIncludeFilter

`func (o *RecurringTaskListResponseResourcesInner) GetIncludeFilter() []string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *RecurringTaskListResponseResourcesInner) GetIncludeFilterOk() (*[]string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *RecurringTaskListResponseResourcesInner) SetIncludeFilter(v []string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *RecurringTaskListResponseResourcesInner) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetRetainPreviousReportCount

`func (o *RecurringTaskListResponseResourcesInner) GetRetainPreviousReportCount() int64`

GetRetainPreviousReportCount returns the RetainPreviousReportCount field if non-nil, zero value otherwise.

### GetRetainPreviousReportCountOk

`func (o *RecurringTaskListResponseResourcesInner) GetRetainPreviousReportCountOk() (*int64, bool)`

GetRetainPreviousReportCountOk returns a tuple with the RetainPreviousReportCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousReportCount

`func (o *RecurringTaskListResponseResourcesInner) SetRetainPreviousReportCount(v int64)`

SetRetainPreviousReportCount sets RetainPreviousReportCount field to given value.

### HasRetainPreviousReportCount

`func (o *RecurringTaskListResponseResourcesInner) HasRetainPreviousReportCount() bool`

HasRetainPreviousReportCount returns a boolean if a field has been set.

### GetRetainPreviousReportAge

`func (o *RecurringTaskListResponseResourcesInner) GetRetainPreviousReportAge() string`

GetRetainPreviousReportAge returns the RetainPreviousReportAge field if non-nil, zero value otherwise.

### GetRetainPreviousReportAgeOk

`func (o *RecurringTaskListResponseResourcesInner) GetRetainPreviousReportAgeOk() (*string, bool)`

GetRetainPreviousReportAgeOk returns a tuple with the RetainPreviousReportAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousReportAge

`func (o *RecurringTaskListResponseResourcesInner) SetRetainPreviousReportAge(v string)`

SetRetainPreviousReportAge sets RetainPreviousReportAge field to given value.

### HasRetainPreviousReportAge

`func (o *RecurringTaskListResponseResourcesInner) HasRetainPreviousReportAge() bool`

HasRetainPreviousReportAge returns a boolean if a field has been set.

### GetSleepDuration

`func (o *RecurringTaskListResponseResourcesInner) GetSleepDuration() string`

GetSleepDuration returns the SleepDuration field if non-nil, zero value otherwise.

### GetSleepDurationOk

`func (o *RecurringTaskListResponseResourcesInner) GetSleepDurationOk() (*string, bool)`

GetSleepDurationOk returns a tuple with the SleepDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepDuration

`func (o *RecurringTaskListResponseResourcesInner) SetSleepDuration(v string)`

SetSleepDuration sets SleepDuration field to given value.

### HasSleepDuration

`func (o *RecurringTaskListResponseResourcesInner) HasSleepDuration() bool`

HasSleepDuration returns a boolean if a field has been set.

### GetDurationToWaitForWorkQueueIdle

`func (o *RecurringTaskListResponseResourcesInner) GetDurationToWaitForWorkQueueIdle() string`

GetDurationToWaitForWorkQueueIdle returns the DurationToWaitForWorkQueueIdle field if non-nil, zero value otherwise.

### GetDurationToWaitForWorkQueueIdleOk

`func (o *RecurringTaskListResponseResourcesInner) GetDurationToWaitForWorkQueueIdleOk() (*string, bool)`

GetDurationToWaitForWorkQueueIdleOk returns a tuple with the DurationToWaitForWorkQueueIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationToWaitForWorkQueueIdle

`func (o *RecurringTaskListResponseResourcesInner) SetDurationToWaitForWorkQueueIdle(v string)`

SetDurationToWaitForWorkQueueIdle sets DurationToWaitForWorkQueueIdle field to given value.

### HasDurationToWaitForWorkQueueIdle

`func (o *RecurringTaskListResponseResourcesInner) HasDurationToWaitForWorkQueueIdle() bool`

HasDurationToWaitForWorkQueueIdle returns a boolean if a field has been set.

### GetLdapURLForSearchExpectedToReturnEntries

`func (o *RecurringTaskListResponseResourcesInner) GetLdapURLForSearchExpectedToReturnEntries() []string`

GetLdapURLForSearchExpectedToReturnEntries returns the LdapURLForSearchExpectedToReturnEntries field if non-nil, zero value otherwise.

### GetLdapURLForSearchExpectedToReturnEntriesOk

`func (o *RecurringTaskListResponseResourcesInner) GetLdapURLForSearchExpectedToReturnEntriesOk() (*[]string, bool)`

GetLdapURLForSearchExpectedToReturnEntriesOk returns a tuple with the LdapURLForSearchExpectedToReturnEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapURLForSearchExpectedToReturnEntries

`func (o *RecurringTaskListResponseResourcesInner) SetLdapURLForSearchExpectedToReturnEntries(v []string)`

SetLdapURLForSearchExpectedToReturnEntries sets LdapURLForSearchExpectedToReturnEntries field to given value.

### HasLdapURLForSearchExpectedToReturnEntries

`func (o *RecurringTaskListResponseResourcesInner) HasLdapURLForSearchExpectedToReturnEntries() bool`

HasLdapURLForSearchExpectedToReturnEntries returns a boolean if a field has been set.

### GetSearchInterval

`func (o *RecurringTaskListResponseResourcesInner) GetSearchInterval() string`

GetSearchInterval returns the SearchInterval field if non-nil, zero value otherwise.

### GetSearchIntervalOk

`func (o *RecurringTaskListResponseResourcesInner) GetSearchIntervalOk() (*string, bool)`

GetSearchIntervalOk returns a tuple with the SearchInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchInterval

`func (o *RecurringTaskListResponseResourcesInner) SetSearchInterval(v string)`

SetSearchInterval sets SearchInterval field to given value.

### HasSearchInterval

`func (o *RecurringTaskListResponseResourcesInner) HasSearchInterval() bool`

HasSearchInterval returns a boolean if a field has been set.

### GetSearchTimeLimit

`func (o *RecurringTaskListResponseResourcesInner) GetSearchTimeLimit() string`

GetSearchTimeLimit returns the SearchTimeLimit field if non-nil, zero value otherwise.

### GetSearchTimeLimitOk

`func (o *RecurringTaskListResponseResourcesInner) GetSearchTimeLimitOk() (*string, bool)`

GetSearchTimeLimitOk returns a tuple with the SearchTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTimeLimit

`func (o *RecurringTaskListResponseResourcesInner) SetSearchTimeLimit(v string)`

SetSearchTimeLimit sets SearchTimeLimit field to given value.

### HasSearchTimeLimit

`func (o *RecurringTaskListResponseResourcesInner) HasSearchTimeLimit() bool`

HasSearchTimeLimit returns a boolean if a field has been set.

### GetDurationToWaitForSearchToReturnEntries

`func (o *RecurringTaskListResponseResourcesInner) GetDurationToWaitForSearchToReturnEntries() string`

GetDurationToWaitForSearchToReturnEntries returns the DurationToWaitForSearchToReturnEntries field if non-nil, zero value otherwise.

### GetDurationToWaitForSearchToReturnEntriesOk

`func (o *RecurringTaskListResponseResourcesInner) GetDurationToWaitForSearchToReturnEntriesOk() (*string, bool)`

GetDurationToWaitForSearchToReturnEntriesOk returns a tuple with the DurationToWaitForSearchToReturnEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationToWaitForSearchToReturnEntries

`func (o *RecurringTaskListResponseResourcesInner) SetDurationToWaitForSearchToReturnEntries(v string)`

SetDurationToWaitForSearchToReturnEntries sets DurationToWaitForSearchToReturnEntries field to given value.

### HasDurationToWaitForSearchToReturnEntries

`func (o *RecurringTaskListResponseResourcesInner) HasDurationToWaitForSearchToReturnEntries() bool`

HasDurationToWaitForSearchToReturnEntries returns a boolean if a field has been set.

### GetTaskReturnStateIfTimeoutIsEncountered

`func (o *RecurringTaskListResponseResourcesInner) GetTaskReturnStateIfTimeoutIsEncountered() EnumrecurringTaskTaskReturnStateIfTimeoutIsEncounteredProp`

GetTaskReturnStateIfTimeoutIsEncountered returns the TaskReturnStateIfTimeoutIsEncountered field if non-nil, zero value otherwise.

### GetTaskReturnStateIfTimeoutIsEncounteredOk

`func (o *RecurringTaskListResponseResourcesInner) GetTaskReturnStateIfTimeoutIsEncounteredOk() (*EnumrecurringTaskTaskReturnStateIfTimeoutIsEncounteredProp, bool)`

GetTaskReturnStateIfTimeoutIsEncounteredOk returns a tuple with the TaskReturnStateIfTimeoutIsEncountered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskReturnStateIfTimeoutIsEncountered

`func (o *RecurringTaskListResponseResourcesInner) SetTaskReturnStateIfTimeoutIsEncountered(v EnumrecurringTaskTaskReturnStateIfTimeoutIsEncounteredProp)`

SetTaskReturnStateIfTimeoutIsEncountered sets TaskReturnStateIfTimeoutIsEncountered field to given value.

### HasTaskReturnStateIfTimeoutIsEncountered

`func (o *RecurringTaskListResponseResourcesInner) HasTaskReturnStateIfTimeoutIsEncountered() bool`

HasTaskReturnStateIfTimeoutIsEncountered returns a boolean if a field has been set.

### GetLdifDirectory

`func (o *RecurringTaskListResponseResourcesInner) GetLdifDirectory() string`

GetLdifDirectory returns the LdifDirectory field if non-nil, zero value otherwise.

### GetLdifDirectoryOk

`func (o *RecurringTaskListResponseResourcesInner) GetLdifDirectoryOk() (*string, bool)`

GetLdifDirectoryOk returns a tuple with the LdifDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdifDirectory

`func (o *RecurringTaskListResponseResourcesInner) SetLdifDirectory(v string)`

SetLdifDirectory sets LdifDirectory field to given value.


### GetBackendID

`func (o *RecurringTaskListResponseResourcesInner) GetBackendID() []string`

GetBackendID returns the BackendID field if non-nil, zero value otherwise.

### GetBackendIDOk

`func (o *RecurringTaskListResponseResourcesInner) GetBackendIDOk() (*[]string, bool)`

GetBackendIDOk returns a tuple with the BackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendID

`func (o *RecurringTaskListResponseResourcesInner) SetBackendID(v []string)`

SetBackendID sets BackendID field to given value.

### HasBackendID

`func (o *RecurringTaskListResponseResourcesInner) HasBackendID() bool`

HasBackendID returns a boolean if a field has been set.

### GetExcludeBackendID

`func (o *RecurringTaskListResponseResourcesInner) GetExcludeBackendID() []string`

GetExcludeBackendID returns the ExcludeBackendID field if non-nil, zero value otherwise.

### GetExcludeBackendIDOk

`func (o *RecurringTaskListResponseResourcesInner) GetExcludeBackendIDOk() (*[]string, bool)`

GetExcludeBackendIDOk returns a tuple with the ExcludeBackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeBackendID

`func (o *RecurringTaskListResponseResourcesInner) SetExcludeBackendID(v []string)`

SetExcludeBackendID sets ExcludeBackendID field to given value.

### HasExcludeBackendID

`func (o *RecurringTaskListResponseResourcesInner) HasExcludeBackendID() bool`

HasExcludeBackendID returns a boolean if a field has been set.

### GetRetainPreviousLDIFExportCount

`func (o *RecurringTaskListResponseResourcesInner) GetRetainPreviousLDIFExportCount() int64`

GetRetainPreviousLDIFExportCount returns the RetainPreviousLDIFExportCount field if non-nil, zero value otherwise.

### GetRetainPreviousLDIFExportCountOk

`func (o *RecurringTaskListResponseResourcesInner) GetRetainPreviousLDIFExportCountOk() (*int64, bool)`

GetRetainPreviousLDIFExportCountOk returns a tuple with the RetainPreviousLDIFExportCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousLDIFExportCount

`func (o *RecurringTaskListResponseResourcesInner) SetRetainPreviousLDIFExportCount(v int64)`

SetRetainPreviousLDIFExportCount sets RetainPreviousLDIFExportCount field to given value.

### HasRetainPreviousLDIFExportCount

`func (o *RecurringTaskListResponseResourcesInner) HasRetainPreviousLDIFExportCount() bool`

HasRetainPreviousLDIFExportCount returns a boolean if a field has been set.

### GetRetainPreviousLDIFExportAge

`func (o *RecurringTaskListResponseResourcesInner) GetRetainPreviousLDIFExportAge() string`

GetRetainPreviousLDIFExportAge returns the RetainPreviousLDIFExportAge field if non-nil, zero value otherwise.

### GetRetainPreviousLDIFExportAgeOk

`func (o *RecurringTaskListResponseResourcesInner) GetRetainPreviousLDIFExportAgeOk() (*string, bool)`

GetRetainPreviousLDIFExportAgeOk returns a tuple with the RetainPreviousLDIFExportAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousLDIFExportAge

`func (o *RecurringTaskListResponseResourcesInner) SetRetainPreviousLDIFExportAge(v string)`

SetRetainPreviousLDIFExportAge sets RetainPreviousLDIFExportAge field to given value.

### HasRetainPreviousLDIFExportAge

`func (o *RecurringTaskListResponseResourcesInner) HasRetainPreviousLDIFExportAge() bool`

HasRetainPreviousLDIFExportAge returns a boolean if a field has been set.

### GetPostLDIFExportTaskProcessor

`func (o *RecurringTaskListResponseResourcesInner) GetPostLDIFExportTaskProcessor() []string`

GetPostLDIFExportTaskProcessor returns the PostLDIFExportTaskProcessor field if non-nil, zero value otherwise.

### GetPostLDIFExportTaskProcessorOk

`func (o *RecurringTaskListResponseResourcesInner) GetPostLDIFExportTaskProcessorOk() (*[]string, bool)`

GetPostLDIFExportTaskProcessorOk returns a tuple with the PostLDIFExportTaskProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostLDIFExportTaskProcessor

`func (o *RecurringTaskListResponseResourcesInner) SetPostLDIFExportTaskProcessor(v []string)`

SetPostLDIFExportTaskProcessor sets PostLDIFExportTaskProcessor field to given value.

### HasPostLDIFExportTaskProcessor

`func (o *RecurringTaskListResponseResourcesInner) HasPostLDIFExportTaskProcessor() bool`

HasPostLDIFExportTaskProcessor returns a boolean if a field has been set.

### GetCommandPath

`func (o *RecurringTaskListResponseResourcesInner) GetCommandPath() string`

GetCommandPath returns the CommandPath field if non-nil, zero value otherwise.

### GetCommandPathOk

`func (o *RecurringTaskListResponseResourcesInner) GetCommandPathOk() (*string, bool)`

GetCommandPathOk returns a tuple with the CommandPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandPath

`func (o *RecurringTaskListResponseResourcesInner) SetCommandPath(v string)`

SetCommandPath sets CommandPath field to given value.


### GetCommandArguments

`func (o *RecurringTaskListResponseResourcesInner) GetCommandArguments() string`

GetCommandArguments returns the CommandArguments field if non-nil, zero value otherwise.

### GetCommandArgumentsOk

`func (o *RecurringTaskListResponseResourcesInner) GetCommandArgumentsOk() (*string, bool)`

GetCommandArgumentsOk returns a tuple with the CommandArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandArguments

`func (o *RecurringTaskListResponseResourcesInner) SetCommandArguments(v string)`

SetCommandArguments sets CommandArguments field to given value.

### HasCommandArguments

`func (o *RecurringTaskListResponseResourcesInner) HasCommandArguments() bool`

HasCommandArguments returns a boolean if a field has been set.

### GetCommandOutputFileBaseName

`func (o *RecurringTaskListResponseResourcesInner) GetCommandOutputFileBaseName() string`

GetCommandOutputFileBaseName returns the CommandOutputFileBaseName field if non-nil, zero value otherwise.

### GetCommandOutputFileBaseNameOk

`func (o *RecurringTaskListResponseResourcesInner) GetCommandOutputFileBaseNameOk() (*string, bool)`

GetCommandOutputFileBaseNameOk returns a tuple with the CommandOutputFileBaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandOutputFileBaseName

`func (o *RecurringTaskListResponseResourcesInner) SetCommandOutputFileBaseName(v string)`

SetCommandOutputFileBaseName sets CommandOutputFileBaseName field to given value.

### HasCommandOutputFileBaseName

`func (o *RecurringTaskListResponseResourcesInner) HasCommandOutputFileBaseName() bool`

HasCommandOutputFileBaseName returns a boolean if a field has been set.

### GetRetainPreviousOutputFileCount

`func (o *RecurringTaskListResponseResourcesInner) GetRetainPreviousOutputFileCount() int64`

GetRetainPreviousOutputFileCount returns the RetainPreviousOutputFileCount field if non-nil, zero value otherwise.

### GetRetainPreviousOutputFileCountOk

`func (o *RecurringTaskListResponseResourcesInner) GetRetainPreviousOutputFileCountOk() (*int64, bool)`

GetRetainPreviousOutputFileCountOk returns a tuple with the RetainPreviousOutputFileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousOutputFileCount

`func (o *RecurringTaskListResponseResourcesInner) SetRetainPreviousOutputFileCount(v int64)`

SetRetainPreviousOutputFileCount sets RetainPreviousOutputFileCount field to given value.

### HasRetainPreviousOutputFileCount

`func (o *RecurringTaskListResponseResourcesInner) HasRetainPreviousOutputFileCount() bool`

HasRetainPreviousOutputFileCount returns a boolean if a field has been set.

### GetRetainPreviousOutputFileAge

`func (o *RecurringTaskListResponseResourcesInner) GetRetainPreviousOutputFileAge() string`

GetRetainPreviousOutputFileAge returns the RetainPreviousOutputFileAge field if non-nil, zero value otherwise.

### GetRetainPreviousOutputFileAgeOk

`func (o *RecurringTaskListResponseResourcesInner) GetRetainPreviousOutputFileAgeOk() (*string, bool)`

GetRetainPreviousOutputFileAgeOk returns a tuple with the RetainPreviousOutputFileAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousOutputFileAge

`func (o *RecurringTaskListResponseResourcesInner) SetRetainPreviousOutputFileAge(v string)`

SetRetainPreviousOutputFileAge sets RetainPreviousOutputFileAge field to given value.

### HasRetainPreviousOutputFileAge

`func (o *RecurringTaskListResponseResourcesInner) HasRetainPreviousOutputFileAge() bool`

HasRetainPreviousOutputFileAge returns a boolean if a field has been set.

### GetLogCommandOutput

`func (o *RecurringTaskListResponseResourcesInner) GetLogCommandOutput() bool`

GetLogCommandOutput returns the LogCommandOutput field if non-nil, zero value otherwise.

### GetLogCommandOutputOk

`func (o *RecurringTaskListResponseResourcesInner) GetLogCommandOutputOk() (*bool, bool)`

GetLogCommandOutputOk returns a tuple with the LogCommandOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogCommandOutput

`func (o *RecurringTaskListResponseResourcesInner) SetLogCommandOutput(v bool)`

SetLogCommandOutput sets LogCommandOutput field to given value.

### HasLogCommandOutput

`func (o *RecurringTaskListResponseResourcesInner) HasLogCommandOutput() bool`

HasLogCommandOutput returns a boolean if a field has been set.

### GetTaskCompletionStateForNonzeroExitCode

`func (o *RecurringTaskListResponseResourcesInner) GetTaskCompletionStateForNonzeroExitCode() EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp`

GetTaskCompletionStateForNonzeroExitCode returns the TaskCompletionStateForNonzeroExitCode field if non-nil, zero value otherwise.

### GetTaskCompletionStateForNonzeroExitCodeOk

`func (o *RecurringTaskListResponseResourcesInner) GetTaskCompletionStateForNonzeroExitCodeOk() (*EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp, bool)`

GetTaskCompletionStateForNonzeroExitCodeOk returns a tuple with the TaskCompletionStateForNonzeroExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCompletionStateForNonzeroExitCode

`func (o *RecurringTaskListResponseResourcesInner) SetTaskCompletionStateForNonzeroExitCode(v EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp)`

SetTaskCompletionStateForNonzeroExitCode sets TaskCompletionStateForNonzeroExitCode field to given value.

### HasTaskCompletionStateForNonzeroExitCode

`func (o *RecurringTaskListResponseResourcesInner) HasTaskCompletionStateForNonzeroExitCode() bool`

HasTaskCompletionStateForNonzeroExitCode returns a boolean if a field has been set.

### GetWorkingDirectory

`func (o *RecurringTaskListResponseResourcesInner) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *RecurringTaskListResponseResourcesInner) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *RecurringTaskListResponseResourcesInner) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.

### HasWorkingDirectory

`func (o *RecurringTaskListResponseResourcesInner) HasWorkingDirectory() bool`

HasWorkingDirectory returns a boolean if a field has been set.

### GetTargetDirectory

`func (o *RecurringTaskListResponseResourcesInner) GetTargetDirectory() string`

GetTargetDirectory returns the TargetDirectory field if non-nil, zero value otherwise.

### GetTargetDirectoryOk

`func (o *RecurringTaskListResponseResourcesInner) GetTargetDirectoryOk() (*string, bool)`

GetTargetDirectoryOk returns a tuple with the TargetDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDirectory

`func (o *RecurringTaskListResponseResourcesInner) SetTargetDirectory(v string)`

SetTargetDirectory sets TargetDirectory field to given value.


### GetFilenamePattern

`func (o *RecurringTaskListResponseResourcesInner) GetFilenamePattern() string`

GetFilenamePattern returns the FilenamePattern field if non-nil, zero value otherwise.

### GetFilenamePatternOk

`func (o *RecurringTaskListResponseResourcesInner) GetFilenamePatternOk() (*string, bool)`

GetFilenamePatternOk returns a tuple with the FilenamePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilenamePattern

`func (o *RecurringTaskListResponseResourcesInner) SetFilenamePattern(v string)`

SetFilenamePattern sets FilenamePattern field to given value.


### GetTimestampFormat

`func (o *RecurringTaskListResponseResourcesInner) GetTimestampFormat() EnumrecurringTaskTimestampFormatProp`

GetTimestampFormat returns the TimestampFormat field if non-nil, zero value otherwise.

### GetTimestampFormatOk

`func (o *RecurringTaskListResponseResourcesInner) GetTimestampFormatOk() (*EnumrecurringTaskTimestampFormatProp, bool)`

GetTimestampFormatOk returns a tuple with the TimestampFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampFormat

`func (o *RecurringTaskListResponseResourcesInner) SetTimestampFormat(v EnumrecurringTaskTimestampFormatProp)`

SetTimestampFormat sets TimestampFormat field to given value.


### GetRetainFileCount

`func (o *RecurringTaskListResponseResourcesInner) GetRetainFileCount() int64`

GetRetainFileCount returns the RetainFileCount field if non-nil, zero value otherwise.

### GetRetainFileCountOk

`func (o *RecurringTaskListResponseResourcesInner) GetRetainFileCountOk() (*int64, bool)`

GetRetainFileCountOk returns a tuple with the RetainFileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainFileCount

`func (o *RecurringTaskListResponseResourcesInner) SetRetainFileCount(v int64)`

SetRetainFileCount sets RetainFileCount field to given value.

### HasRetainFileCount

`func (o *RecurringTaskListResponseResourcesInner) HasRetainFileCount() bool`

HasRetainFileCount returns a boolean if a field has been set.

### GetRetainFileAge

`func (o *RecurringTaskListResponseResourcesInner) GetRetainFileAge() string`

GetRetainFileAge returns the RetainFileAge field if non-nil, zero value otherwise.

### GetRetainFileAgeOk

`func (o *RecurringTaskListResponseResourcesInner) GetRetainFileAgeOk() (*string, bool)`

GetRetainFileAgeOk returns a tuple with the RetainFileAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainFileAge

`func (o *RecurringTaskListResponseResourcesInner) SetRetainFileAge(v string)`

SetRetainFileAge sets RetainFileAge field to given value.

### HasRetainFileAge

`func (o *RecurringTaskListResponseResourcesInner) HasRetainFileAge() bool`

HasRetainFileAge returns a boolean if a field has been set.

### GetRetainAggregateFileSize

`func (o *RecurringTaskListResponseResourcesInner) GetRetainAggregateFileSize() string`

GetRetainAggregateFileSize returns the RetainAggregateFileSize field if non-nil, zero value otherwise.

### GetRetainAggregateFileSizeOk

`func (o *RecurringTaskListResponseResourcesInner) GetRetainAggregateFileSizeOk() (*string, bool)`

GetRetainAggregateFileSizeOk returns a tuple with the RetainAggregateFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainAggregateFileSize

`func (o *RecurringTaskListResponseResourcesInner) SetRetainAggregateFileSize(v string)`

SetRetainAggregateFileSize sets RetainAggregateFileSize field to given value.

### HasRetainAggregateFileSize

`func (o *RecurringTaskListResponseResourcesInner) HasRetainAggregateFileSize() bool`

HasRetainAggregateFileSize returns a boolean if a field has been set.

### GetExtensionClass

`func (o *RecurringTaskListResponseResourcesInner) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *RecurringTaskListResponseResourcesInner) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *RecurringTaskListResponseResourcesInner) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *RecurringTaskListResponseResourcesInner) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *RecurringTaskListResponseResourcesInner) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *RecurringTaskListResponseResourcesInner) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *RecurringTaskListResponseResourcesInner) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


