# AddRecurringTask200Response

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
**JstackCount** | Pointer to **int64** | The number of times to invoke the jstack utility to obtain a stack trace of all threads running in the JVM. A value of zero indicates that the jstack utility should not be invoked. | [optional] 
**ReportCount** | Pointer to **int64** | The number of intervals of data to collect from tools that use sample-based reporting, like vmstat, iostat, and mpstat. A value of zero indicates that these kinds of tools should not be used to collect any information. | [optional] 
**ReportIntervalSeconds** | Pointer to **int64** | The duration (in seconds) between each interval of data to collect from tools that use sample-based reporting, like vmstat, iostat, and mpstat. | [optional] 
**LogDuration** | Pointer to **string** | The maximum age (leading up to the time the collect-support-data tool was invoked) for log content to include in the support data archive. | [optional] 
**LogFileHeadCollectionSize** | Pointer to **string** | The amount of data to collect from the beginning of each log file included in the support data archive. | [optional] 
**LogFileTailCollectionSize** | Pointer to **string** | The amount of data to collect from the end of each log file included in the support data archive. | [optional] 
**Comment** | Pointer to **string** | An optional comment to include in a README file within the support data archive. | [optional] 
**RetainPreviousSupportDataArchiveCount** | Pointer to **int64** | The minimum number of previous support data archives that should be preserved after a new archive is generated. | [optional] 
**RetainPreviousSupportDataArchiveAge** | Pointer to **string** | The minimum age of previous support data archives that should be preserved after a new archive is generated. | [optional] 
**LdifDirectory** | **string** | The directory in which LDIF export files will be placed. The directory must already exist. | 
**BackendID** | Pointer to **[]string** | The backend ID for a backend to be exported. | [optional] 
**ExcludeBackendID** | Pointer to **[]string** | The backend ID for a backend to be excluded from the export. | [optional] 
**RetainPreviousLDIFExportCount** | Pointer to **int64** | The minimum number of previous LDIF exports that should be preserved after a new export completes successfully. | [optional] 
**RetainPreviousLDIFExportAge** | Pointer to **string** | The minimum age of previous LDIF exports that should be preserved after a new export completes successfully. | [optional] 
**PostLDIFExportTaskProcessor** | Pointer to **[]string** | An optional set of post-LDIF-export task processors that should be invoked for the resulting LDIF export files. | [optional] 
**BaseOutputDirectory** | **string** | The base directory below which generated reports will be written. Each invocation of the audit-data-security task will create a new subdirectory below this base directory whose name is a timestamp indicating when the report was generated. | 
**DataSecurityAuditor** | Pointer to **[]string** | The set of data security auditors that should be invoked. If no auditors are specified, then all auditors defined in the configuration will be used. | [optional] 
**Backend** | Pointer to **[]string** | The set of backends that should be examined. If no backends are specified, then all backends that support this functionality will be included. | [optional] 
**IncludeFilter** | Pointer to **[]string** | A filter that will be used to identify entries that may be included in the generated report. If multiple filters are specified, then any entry that matches at least one of the filters will be included. If no filters are specified, then all entries will be included. | [optional] 
**RetainPreviousReportCount** | Pointer to **int64** | The minimum number of previous reports that should be preserved after a new report is generated. | [optional] 
**RetainPreviousReportAge** | Pointer to **string** | The minimum age of previous reports that should be preserved after a new report completes successfully. | [optional] 
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

### NewAddRecurringTask200Response

`func NewAddRecurringTask200Response(id string, schemas []EnumthirdPartyRecurringTaskSchemaUrn, profileDirectory string, backupDirectory string, taskJavaClass string, taskObjectClass []string, outputDirectory string, ldifDirectory string, baseOutputDirectory string, commandPath string, targetDirectory string, filenamePattern string, timestampFormat EnumrecurringTaskTimestampFormatProp, extensionClass string, ) *AddRecurringTask200Response`

NewAddRecurringTask200Response instantiates a new AddRecurringTask200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddRecurringTask200ResponseWithDefaults

`func NewAddRecurringTask200ResponseWithDefaults() *AddRecurringTask200Response`

NewAddRecurringTask200ResponseWithDefaults instantiates a new AddRecurringTask200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddRecurringTask200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddRecurringTask200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddRecurringTask200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddRecurringTask200Response) GetSchemas() []EnumthirdPartyRecurringTaskSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddRecurringTask200Response) GetSchemasOk() (*[]EnumthirdPartyRecurringTaskSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddRecurringTask200Response) SetSchemas(v []EnumthirdPartyRecurringTaskSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetProfileDirectory

`func (o *AddRecurringTask200Response) GetProfileDirectory() string`

GetProfileDirectory returns the ProfileDirectory field if non-nil, zero value otherwise.

### GetProfileDirectoryOk

`func (o *AddRecurringTask200Response) GetProfileDirectoryOk() (*string, bool)`

GetProfileDirectoryOk returns a tuple with the ProfileDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileDirectory

`func (o *AddRecurringTask200Response) SetProfileDirectory(v string)`

SetProfileDirectory sets ProfileDirectory field to given value.


### GetIncludePath

`func (o *AddRecurringTask200Response) GetIncludePath() []string`

GetIncludePath returns the IncludePath field if non-nil, zero value otherwise.

### GetIncludePathOk

`func (o *AddRecurringTask200Response) GetIncludePathOk() (*[]string, bool)`

GetIncludePathOk returns a tuple with the IncludePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePath

`func (o *AddRecurringTask200Response) SetIncludePath(v []string)`

SetIncludePath sets IncludePath field to given value.

### HasIncludePath

`func (o *AddRecurringTask200Response) HasIncludePath() bool`

HasIncludePath returns a boolean if a field has been set.

### GetRetainPreviousProfileCount

`func (o *AddRecurringTask200Response) GetRetainPreviousProfileCount() int64`

GetRetainPreviousProfileCount returns the RetainPreviousProfileCount field if non-nil, zero value otherwise.

### GetRetainPreviousProfileCountOk

`func (o *AddRecurringTask200Response) GetRetainPreviousProfileCountOk() (*int64, bool)`

GetRetainPreviousProfileCountOk returns a tuple with the RetainPreviousProfileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousProfileCount

`func (o *AddRecurringTask200Response) SetRetainPreviousProfileCount(v int64)`

SetRetainPreviousProfileCount sets RetainPreviousProfileCount field to given value.

### HasRetainPreviousProfileCount

`func (o *AddRecurringTask200Response) HasRetainPreviousProfileCount() bool`

HasRetainPreviousProfileCount returns a boolean if a field has been set.

### GetRetainPreviousProfileAge

`func (o *AddRecurringTask200Response) GetRetainPreviousProfileAge() string`

GetRetainPreviousProfileAge returns the RetainPreviousProfileAge field if non-nil, zero value otherwise.

### GetRetainPreviousProfileAgeOk

`func (o *AddRecurringTask200Response) GetRetainPreviousProfileAgeOk() (*string, bool)`

GetRetainPreviousProfileAgeOk returns a tuple with the RetainPreviousProfileAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousProfileAge

`func (o *AddRecurringTask200Response) SetRetainPreviousProfileAge(v string)`

SetRetainPreviousProfileAge sets RetainPreviousProfileAge field to given value.

### HasRetainPreviousProfileAge

`func (o *AddRecurringTask200Response) HasRetainPreviousProfileAge() bool`

HasRetainPreviousProfileAge returns a boolean if a field has been set.

### GetDescription

`func (o *AddRecurringTask200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddRecurringTask200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddRecurringTask200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddRecurringTask200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCancelOnTaskDependencyFailure

`func (o *AddRecurringTask200Response) GetCancelOnTaskDependencyFailure() bool`

GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field if non-nil, zero value otherwise.

### GetCancelOnTaskDependencyFailureOk

`func (o *AddRecurringTask200Response) GetCancelOnTaskDependencyFailureOk() (*bool, bool)`

GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOnTaskDependencyFailure

`func (o *AddRecurringTask200Response) SetCancelOnTaskDependencyFailure(v bool)`

SetCancelOnTaskDependencyFailure sets CancelOnTaskDependencyFailure field to given value.

### HasCancelOnTaskDependencyFailure

`func (o *AddRecurringTask200Response) HasCancelOnTaskDependencyFailure() bool`

HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.

### GetEmailOnStart

`func (o *AddRecurringTask200Response) GetEmailOnStart() []string`

GetEmailOnStart returns the EmailOnStart field if non-nil, zero value otherwise.

### GetEmailOnStartOk

`func (o *AddRecurringTask200Response) GetEmailOnStartOk() (*[]string, bool)`

GetEmailOnStartOk returns a tuple with the EmailOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnStart

`func (o *AddRecurringTask200Response) SetEmailOnStart(v []string)`

SetEmailOnStart sets EmailOnStart field to given value.

### HasEmailOnStart

`func (o *AddRecurringTask200Response) HasEmailOnStart() bool`

HasEmailOnStart returns a boolean if a field has been set.

### GetEmailOnSuccess

`func (o *AddRecurringTask200Response) GetEmailOnSuccess() []string`

GetEmailOnSuccess returns the EmailOnSuccess field if non-nil, zero value otherwise.

### GetEmailOnSuccessOk

`func (o *AddRecurringTask200Response) GetEmailOnSuccessOk() (*[]string, bool)`

GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnSuccess

`func (o *AddRecurringTask200Response) SetEmailOnSuccess(v []string)`

SetEmailOnSuccess sets EmailOnSuccess field to given value.

### HasEmailOnSuccess

`func (o *AddRecurringTask200Response) HasEmailOnSuccess() bool`

HasEmailOnSuccess returns a boolean if a field has been set.

### GetEmailOnFailure

`func (o *AddRecurringTask200Response) GetEmailOnFailure() []string`

GetEmailOnFailure returns the EmailOnFailure field if non-nil, zero value otherwise.

### GetEmailOnFailureOk

`func (o *AddRecurringTask200Response) GetEmailOnFailureOk() (*[]string, bool)`

GetEmailOnFailureOk returns a tuple with the EmailOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnFailure

`func (o *AddRecurringTask200Response) SetEmailOnFailure(v []string)`

SetEmailOnFailure sets EmailOnFailure field to given value.

### HasEmailOnFailure

`func (o *AddRecurringTask200Response) HasEmailOnFailure() bool`

HasEmailOnFailure returns a boolean if a field has been set.

### GetAlertOnStart

`func (o *AddRecurringTask200Response) GetAlertOnStart() bool`

GetAlertOnStart returns the AlertOnStart field if non-nil, zero value otherwise.

### GetAlertOnStartOk

`func (o *AddRecurringTask200Response) GetAlertOnStartOk() (*bool, bool)`

GetAlertOnStartOk returns a tuple with the AlertOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnStart

`func (o *AddRecurringTask200Response) SetAlertOnStart(v bool)`

SetAlertOnStart sets AlertOnStart field to given value.

### HasAlertOnStart

`func (o *AddRecurringTask200Response) HasAlertOnStart() bool`

HasAlertOnStart returns a boolean if a field has been set.

### GetAlertOnSuccess

`func (o *AddRecurringTask200Response) GetAlertOnSuccess() bool`

GetAlertOnSuccess returns the AlertOnSuccess field if non-nil, zero value otherwise.

### GetAlertOnSuccessOk

`func (o *AddRecurringTask200Response) GetAlertOnSuccessOk() (*bool, bool)`

GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnSuccess

`func (o *AddRecurringTask200Response) SetAlertOnSuccess(v bool)`

SetAlertOnSuccess sets AlertOnSuccess field to given value.

### HasAlertOnSuccess

`func (o *AddRecurringTask200Response) HasAlertOnSuccess() bool`

HasAlertOnSuccess returns a boolean if a field has been set.

### GetAlertOnFailure

`func (o *AddRecurringTask200Response) GetAlertOnFailure() bool`

GetAlertOnFailure returns the AlertOnFailure field if non-nil, zero value otherwise.

### GetAlertOnFailureOk

`func (o *AddRecurringTask200Response) GetAlertOnFailureOk() (*bool, bool)`

GetAlertOnFailureOk returns a tuple with the AlertOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFailure

`func (o *AddRecurringTask200Response) SetAlertOnFailure(v bool)`

SetAlertOnFailure sets AlertOnFailure field to given value.

### HasAlertOnFailure

`func (o *AddRecurringTask200Response) HasAlertOnFailure() bool`

HasAlertOnFailure returns a boolean if a field has been set.

### GetMeta

`func (o *AddRecurringTask200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddRecurringTask200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddRecurringTask200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddRecurringTask200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AddRecurringTask200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AddRecurringTask200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AddRecurringTask200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AddRecurringTask200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetReason

`func (o *AddRecurringTask200Response) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AddRecurringTask200Response) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AddRecurringTask200Response) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AddRecurringTask200Response) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetBackupDirectory

`func (o *AddRecurringTask200Response) GetBackupDirectory() string`

GetBackupDirectory returns the BackupDirectory field if non-nil, zero value otherwise.

### GetBackupDirectoryOk

`func (o *AddRecurringTask200Response) GetBackupDirectoryOk() (*string, bool)`

GetBackupDirectoryOk returns a tuple with the BackupDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDirectory

`func (o *AddRecurringTask200Response) SetBackupDirectory(v string)`

SetBackupDirectory sets BackupDirectory field to given value.


### GetIncludedBackendID

`func (o *AddRecurringTask200Response) GetIncludedBackendID() []string`

GetIncludedBackendID returns the IncludedBackendID field if non-nil, zero value otherwise.

### GetIncludedBackendIDOk

`func (o *AddRecurringTask200Response) GetIncludedBackendIDOk() (*[]string, bool)`

GetIncludedBackendIDOk returns a tuple with the IncludedBackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedBackendID

`func (o *AddRecurringTask200Response) SetIncludedBackendID(v []string)`

SetIncludedBackendID sets IncludedBackendID field to given value.

### HasIncludedBackendID

`func (o *AddRecurringTask200Response) HasIncludedBackendID() bool`

HasIncludedBackendID returns a boolean if a field has been set.

### GetExcludedBackendID

`func (o *AddRecurringTask200Response) GetExcludedBackendID() []string`

GetExcludedBackendID returns the ExcludedBackendID field if non-nil, zero value otherwise.

### GetExcludedBackendIDOk

`func (o *AddRecurringTask200Response) GetExcludedBackendIDOk() (*[]string, bool)`

GetExcludedBackendIDOk returns a tuple with the ExcludedBackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedBackendID

`func (o *AddRecurringTask200Response) SetExcludedBackendID(v []string)`

SetExcludedBackendID sets ExcludedBackendID field to given value.

### HasExcludedBackendID

`func (o *AddRecurringTask200Response) HasExcludedBackendID() bool`

HasExcludedBackendID returns a boolean if a field has been set.

### GetCompress

`func (o *AddRecurringTask200Response) GetCompress() bool`

GetCompress returns the Compress field if non-nil, zero value otherwise.

### GetCompressOk

`func (o *AddRecurringTask200Response) GetCompressOk() (*bool, bool)`

GetCompressOk returns a tuple with the Compress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompress

`func (o *AddRecurringTask200Response) SetCompress(v bool)`

SetCompress sets Compress field to given value.

### HasCompress

`func (o *AddRecurringTask200Response) HasCompress() bool`

HasCompress returns a boolean if a field has been set.

### GetEncrypt

`func (o *AddRecurringTask200Response) GetEncrypt() bool`

GetEncrypt returns the Encrypt field if non-nil, zero value otherwise.

### GetEncryptOk

`func (o *AddRecurringTask200Response) GetEncryptOk() (*bool, bool)`

GetEncryptOk returns a tuple with the Encrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypt

`func (o *AddRecurringTask200Response) SetEncrypt(v bool)`

SetEncrypt sets Encrypt field to given value.

### HasEncrypt

`func (o *AddRecurringTask200Response) HasEncrypt() bool`

HasEncrypt returns a boolean if a field has been set.

### GetEncryptionSettingsDefinitionID

`func (o *AddRecurringTask200Response) GetEncryptionSettingsDefinitionID() string`

GetEncryptionSettingsDefinitionID returns the EncryptionSettingsDefinitionID field if non-nil, zero value otherwise.

### GetEncryptionSettingsDefinitionIDOk

`func (o *AddRecurringTask200Response) GetEncryptionSettingsDefinitionIDOk() (*string, bool)`

GetEncryptionSettingsDefinitionIDOk returns a tuple with the EncryptionSettingsDefinitionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionSettingsDefinitionID

`func (o *AddRecurringTask200Response) SetEncryptionSettingsDefinitionID(v string)`

SetEncryptionSettingsDefinitionID sets EncryptionSettingsDefinitionID field to given value.

### HasEncryptionSettingsDefinitionID

`func (o *AddRecurringTask200Response) HasEncryptionSettingsDefinitionID() bool`

HasEncryptionSettingsDefinitionID returns a boolean if a field has been set.

### GetSign

`func (o *AddRecurringTask200Response) GetSign() bool`

GetSign returns the Sign field if non-nil, zero value otherwise.

### GetSignOk

`func (o *AddRecurringTask200Response) GetSignOk() (*bool, bool)`

GetSignOk returns a tuple with the Sign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSign

`func (o *AddRecurringTask200Response) SetSign(v bool)`

SetSign sets Sign field to given value.

### HasSign

`func (o *AddRecurringTask200Response) HasSign() bool`

HasSign returns a boolean if a field has been set.

### GetRetainPreviousFullBackupCount

`func (o *AddRecurringTask200Response) GetRetainPreviousFullBackupCount() int64`

GetRetainPreviousFullBackupCount returns the RetainPreviousFullBackupCount field if non-nil, zero value otherwise.

### GetRetainPreviousFullBackupCountOk

`func (o *AddRecurringTask200Response) GetRetainPreviousFullBackupCountOk() (*int64, bool)`

GetRetainPreviousFullBackupCountOk returns a tuple with the RetainPreviousFullBackupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousFullBackupCount

`func (o *AddRecurringTask200Response) SetRetainPreviousFullBackupCount(v int64)`

SetRetainPreviousFullBackupCount sets RetainPreviousFullBackupCount field to given value.

### HasRetainPreviousFullBackupCount

`func (o *AddRecurringTask200Response) HasRetainPreviousFullBackupCount() bool`

HasRetainPreviousFullBackupCount returns a boolean if a field has been set.

### GetRetainPreviousFullBackupAge

`func (o *AddRecurringTask200Response) GetRetainPreviousFullBackupAge() string`

GetRetainPreviousFullBackupAge returns the RetainPreviousFullBackupAge field if non-nil, zero value otherwise.

### GetRetainPreviousFullBackupAgeOk

`func (o *AddRecurringTask200Response) GetRetainPreviousFullBackupAgeOk() (*string, bool)`

GetRetainPreviousFullBackupAgeOk returns a tuple with the RetainPreviousFullBackupAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousFullBackupAge

`func (o *AddRecurringTask200Response) SetRetainPreviousFullBackupAge(v string)`

SetRetainPreviousFullBackupAge sets RetainPreviousFullBackupAge field to given value.

### HasRetainPreviousFullBackupAge

`func (o *AddRecurringTask200Response) HasRetainPreviousFullBackupAge() bool`

HasRetainPreviousFullBackupAge returns a boolean if a field has been set.

### GetMaxMegabytesPerSecond

`func (o *AddRecurringTask200Response) GetMaxMegabytesPerSecond() int64`

GetMaxMegabytesPerSecond returns the MaxMegabytesPerSecond field if non-nil, zero value otherwise.

### GetMaxMegabytesPerSecondOk

`func (o *AddRecurringTask200Response) GetMaxMegabytesPerSecondOk() (*int64, bool)`

GetMaxMegabytesPerSecondOk returns a tuple with the MaxMegabytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMegabytesPerSecond

`func (o *AddRecurringTask200Response) SetMaxMegabytesPerSecond(v int64)`

SetMaxMegabytesPerSecond sets MaxMegabytesPerSecond field to given value.

### HasMaxMegabytesPerSecond

`func (o *AddRecurringTask200Response) HasMaxMegabytesPerSecond() bool`

HasMaxMegabytesPerSecond returns a boolean if a field has been set.

### GetSleepDuration

`func (o *AddRecurringTask200Response) GetSleepDuration() string`

GetSleepDuration returns the SleepDuration field if non-nil, zero value otherwise.

### GetSleepDurationOk

`func (o *AddRecurringTask200Response) GetSleepDurationOk() (*string, bool)`

GetSleepDurationOk returns a tuple with the SleepDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepDuration

`func (o *AddRecurringTask200Response) SetSleepDuration(v string)`

SetSleepDuration sets SleepDuration field to given value.

### HasSleepDuration

`func (o *AddRecurringTask200Response) HasSleepDuration() bool`

HasSleepDuration returns a boolean if a field has been set.

### GetDurationToWaitForWorkQueueIdle

`func (o *AddRecurringTask200Response) GetDurationToWaitForWorkQueueIdle() string`

GetDurationToWaitForWorkQueueIdle returns the DurationToWaitForWorkQueueIdle field if non-nil, zero value otherwise.

### GetDurationToWaitForWorkQueueIdleOk

`func (o *AddRecurringTask200Response) GetDurationToWaitForWorkQueueIdleOk() (*string, bool)`

GetDurationToWaitForWorkQueueIdleOk returns a tuple with the DurationToWaitForWorkQueueIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationToWaitForWorkQueueIdle

`func (o *AddRecurringTask200Response) SetDurationToWaitForWorkQueueIdle(v string)`

SetDurationToWaitForWorkQueueIdle sets DurationToWaitForWorkQueueIdle field to given value.

### HasDurationToWaitForWorkQueueIdle

`func (o *AddRecurringTask200Response) HasDurationToWaitForWorkQueueIdle() bool`

HasDurationToWaitForWorkQueueIdle returns a boolean if a field has been set.

### GetLdapURLForSearchExpectedToReturnEntries

`func (o *AddRecurringTask200Response) GetLdapURLForSearchExpectedToReturnEntries() []string`

GetLdapURLForSearchExpectedToReturnEntries returns the LdapURLForSearchExpectedToReturnEntries field if non-nil, zero value otherwise.

### GetLdapURLForSearchExpectedToReturnEntriesOk

`func (o *AddRecurringTask200Response) GetLdapURLForSearchExpectedToReturnEntriesOk() (*[]string, bool)`

GetLdapURLForSearchExpectedToReturnEntriesOk returns a tuple with the LdapURLForSearchExpectedToReturnEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapURLForSearchExpectedToReturnEntries

`func (o *AddRecurringTask200Response) SetLdapURLForSearchExpectedToReturnEntries(v []string)`

SetLdapURLForSearchExpectedToReturnEntries sets LdapURLForSearchExpectedToReturnEntries field to given value.

### HasLdapURLForSearchExpectedToReturnEntries

`func (o *AddRecurringTask200Response) HasLdapURLForSearchExpectedToReturnEntries() bool`

HasLdapURLForSearchExpectedToReturnEntries returns a boolean if a field has been set.

### GetSearchInterval

`func (o *AddRecurringTask200Response) GetSearchInterval() string`

GetSearchInterval returns the SearchInterval field if non-nil, zero value otherwise.

### GetSearchIntervalOk

`func (o *AddRecurringTask200Response) GetSearchIntervalOk() (*string, bool)`

GetSearchIntervalOk returns a tuple with the SearchInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchInterval

`func (o *AddRecurringTask200Response) SetSearchInterval(v string)`

SetSearchInterval sets SearchInterval field to given value.

### HasSearchInterval

`func (o *AddRecurringTask200Response) HasSearchInterval() bool`

HasSearchInterval returns a boolean if a field has been set.

### GetSearchTimeLimit

`func (o *AddRecurringTask200Response) GetSearchTimeLimit() string`

GetSearchTimeLimit returns the SearchTimeLimit field if non-nil, zero value otherwise.

### GetSearchTimeLimitOk

`func (o *AddRecurringTask200Response) GetSearchTimeLimitOk() (*string, bool)`

GetSearchTimeLimitOk returns a tuple with the SearchTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTimeLimit

`func (o *AddRecurringTask200Response) SetSearchTimeLimit(v string)`

SetSearchTimeLimit sets SearchTimeLimit field to given value.

### HasSearchTimeLimit

`func (o *AddRecurringTask200Response) HasSearchTimeLimit() bool`

HasSearchTimeLimit returns a boolean if a field has been set.

### GetDurationToWaitForSearchToReturnEntries

`func (o *AddRecurringTask200Response) GetDurationToWaitForSearchToReturnEntries() string`

GetDurationToWaitForSearchToReturnEntries returns the DurationToWaitForSearchToReturnEntries field if non-nil, zero value otherwise.

### GetDurationToWaitForSearchToReturnEntriesOk

`func (o *AddRecurringTask200Response) GetDurationToWaitForSearchToReturnEntriesOk() (*string, bool)`

GetDurationToWaitForSearchToReturnEntriesOk returns a tuple with the DurationToWaitForSearchToReturnEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationToWaitForSearchToReturnEntries

`func (o *AddRecurringTask200Response) SetDurationToWaitForSearchToReturnEntries(v string)`

SetDurationToWaitForSearchToReturnEntries sets DurationToWaitForSearchToReturnEntries field to given value.

### HasDurationToWaitForSearchToReturnEntries

`func (o *AddRecurringTask200Response) HasDurationToWaitForSearchToReturnEntries() bool`

HasDurationToWaitForSearchToReturnEntries returns a boolean if a field has been set.

### GetTaskReturnStateIfTimeoutIsEncountered

`func (o *AddRecurringTask200Response) GetTaskReturnStateIfTimeoutIsEncountered() EnumrecurringTaskTaskReturnStateIfTimeoutIsEncounteredProp`

GetTaskReturnStateIfTimeoutIsEncountered returns the TaskReturnStateIfTimeoutIsEncountered field if non-nil, zero value otherwise.

### GetTaskReturnStateIfTimeoutIsEncounteredOk

`func (o *AddRecurringTask200Response) GetTaskReturnStateIfTimeoutIsEncounteredOk() (*EnumrecurringTaskTaskReturnStateIfTimeoutIsEncounteredProp, bool)`

GetTaskReturnStateIfTimeoutIsEncounteredOk returns a tuple with the TaskReturnStateIfTimeoutIsEncountered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskReturnStateIfTimeoutIsEncountered

`func (o *AddRecurringTask200Response) SetTaskReturnStateIfTimeoutIsEncountered(v EnumrecurringTaskTaskReturnStateIfTimeoutIsEncounteredProp)`

SetTaskReturnStateIfTimeoutIsEncountered sets TaskReturnStateIfTimeoutIsEncountered field to given value.

### HasTaskReturnStateIfTimeoutIsEncountered

`func (o *AddRecurringTask200Response) HasTaskReturnStateIfTimeoutIsEncountered() bool`

HasTaskReturnStateIfTimeoutIsEncountered returns a boolean if a field has been set.

### GetTaskJavaClass

`func (o *AddRecurringTask200Response) GetTaskJavaClass() string`

GetTaskJavaClass returns the TaskJavaClass field if non-nil, zero value otherwise.

### GetTaskJavaClassOk

`func (o *AddRecurringTask200Response) GetTaskJavaClassOk() (*string, bool)`

GetTaskJavaClassOk returns a tuple with the TaskJavaClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskJavaClass

`func (o *AddRecurringTask200Response) SetTaskJavaClass(v string)`

SetTaskJavaClass sets TaskJavaClass field to given value.


### GetTaskObjectClass

`func (o *AddRecurringTask200Response) GetTaskObjectClass() []string`

GetTaskObjectClass returns the TaskObjectClass field if non-nil, zero value otherwise.

### GetTaskObjectClassOk

`func (o *AddRecurringTask200Response) GetTaskObjectClassOk() (*[]string, bool)`

GetTaskObjectClassOk returns a tuple with the TaskObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskObjectClass

`func (o *AddRecurringTask200Response) SetTaskObjectClass(v []string)`

SetTaskObjectClass sets TaskObjectClass field to given value.


### GetTaskAttributeValue

`func (o *AddRecurringTask200Response) GetTaskAttributeValue() []string`

GetTaskAttributeValue returns the TaskAttributeValue field if non-nil, zero value otherwise.

### GetTaskAttributeValueOk

`func (o *AddRecurringTask200Response) GetTaskAttributeValueOk() (*[]string, bool)`

GetTaskAttributeValueOk returns a tuple with the TaskAttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskAttributeValue

`func (o *AddRecurringTask200Response) SetTaskAttributeValue(v []string)`

SetTaskAttributeValue sets TaskAttributeValue field to given value.

### HasTaskAttributeValue

`func (o *AddRecurringTask200Response) HasTaskAttributeValue() bool`

HasTaskAttributeValue returns a boolean if a field has been set.

### GetOutputDirectory

`func (o *AddRecurringTask200Response) GetOutputDirectory() string`

GetOutputDirectory returns the OutputDirectory field if non-nil, zero value otherwise.

### GetOutputDirectoryOk

`func (o *AddRecurringTask200Response) GetOutputDirectoryOk() (*string, bool)`

GetOutputDirectoryOk returns a tuple with the OutputDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDirectory

`func (o *AddRecurringTask200Response) SetOutputDirectory(v string)`

SetOutputDirectory sets OutputDirectory field to given value.


### GetEncryptionPassphraseFile

`func (o *AddRecurringTask200Response) GetEncryptionPassphraseFile() string`

GetEncryptionPassphraseFile returns the EncryptionPassphraseFile field if non-nil, zero value otherwise.

### GetEncryptionPassphraseFileOk

`func (o *AddRecurringTask200Response) GetEncryptionPassphraseFileOk() (*string, bool)`

GetEncryptionPassphraseFileOk returns a tuple with the EncryptionPassphraseFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionPassphraseFile

`func (o *AddRecurringTask200Response) SetEncryptionPassphraseFile(v string)`

SetEncryptionPassphraseFile sets EncryptionPassphraseFile field to given value.

### HasEncryptionPassphraseFile

`func (o *AddRecurringTask200Response) HasEncryptionPassphraseFile() bool`

HasEncryptionPassphraseFile returns a boolean if a field has been set.

### GetIncludeExpensiveData

`func (o *AddRecurringTask200Response) GetIncludeExpensiveData() bool`

GetIncludeExpensiveData returns the IncludeExpensiveData field if non-nil, zero value otherwise.

### GetIncludeExpensiveDataOk

`func (o *AddRecurringTask200Response) GetIncludeExpensiveDataOk() (*bool, bool)`

GetIncludeExpensiveDataOk returns a tuple with the IncludeExpensiveData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeExpensiveData

`func (o *AddRecurringTask200Response) SetIncludeExpensiveData(v bool)`

SetIncludeExpensiveData sets IncludeExpensiveData field to given value.

### HasIncludeExpensiveData

`func (o *AddRecurringTask200Response) HasIncludeExpensiveData() bool`

HasIncludeExpensiveData returns a boolean if a field has been set.

### GetIncludeReplicationStateDump

`func (o *AddRecurringTask200Response) GetIncludeReplicationStateDump() bool`

GetIncludeReplicationStateDump returns the IncludeReplicationStateDump field if non-nil, zero value otherwise.

### GetIncludeReplicationStateDumpOk

`func (o *AddRecurringTask200Response) GetIncludeReplicationStateDumpOk() (*bool, bool)`

GetIncludeReplicationStateDumpOk returns a tuple with the IncludeReplicationStateDump field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeReplicationStateDump

`func (o *AddRecurringTask200Response) SetIncludeReplicationStateDump(v bool)`

SetIncludeReplicationStateDump sets IncludeReplicationStateDump field to given value.

### HasIncludeReplicationStateDump

`func (o *AddRecurringTask200Response) HasIncludeReplicationStateDump() bool`

HasIncludeReplicationStateDump returns a boolean if a field has been set.

### GetIncludeBinaryFiles

`func (o *AddRecurringTask200Response) GetIncludeBinaryFiles() bool`

GetIncludeBinaryFiles returns the IncludeBinaryFiles field if non-nil, zero value otherwise.

### GetIncludeBinaryFilesOk

`func (o *AddRecurringTask200Response) GetIncludeBinaryFilesOk() (*bool, bool)`

GetIncludeBinaryFilesOk returns a tuple with the IncludeBinaryFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBinaryFiles

`func (o *AddRecurringTask200Response) SetIncludeBinaryFiles(v bool)`

SetIncludeBinaryFiles sets IncludeBinaryFiles field to given value.

### HasIncludeBinaryFiles

`func (o *AddRecurringTask200Response) HasIncludeBinaryFiles() bool`

HasIncludeBinaryFiles returns a boolean if a field has been set.

### GetIncludeExtensionSource

`func (o *AddRecurringTask200Response) GetIncludeExtensionSource() bool`

GetIncludeExtensionSource returns the IncludeExtensionSource field if non-nil, zero value otherwise.

### GetIncludeExtensionSourceOk

`func (o *AddRecurringTask200Response) GetIncludeExtensionSourceOk() (*bool, bool)`

GetIncludeExtensionSourceOk returns a tuple with the IncludeExtensionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeExtensionSource

`func (o *AddRecurringTask200Response) SetIncludeExtensionSource(v bool)`

SetIncludeExtensionSource sets IncludeExtensionSource field to given value.

### HasIncludeExtensionSource

`func (o *AddRecurringTask200Response) HasIncludeExtensionSource() bool`

HasIncludeExtensionSource returns a boolean if a field has been set.

### GetUseSequentialMode

`func (o *AddRecurringTask200Response) GetUseSequentialMode() bool`

GetUseSequentialMode returns the UseSequentialMode field if non-nil, zero value otherwise.

### GetUseSequentialModeOk

`func (o *AddRecurringTask200Response) GetUseSequentialModeOk() (*bool, bool)`

GetUseSequentialModeOk returns a tuple with the UseSequentialMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSequentialMode

`func (o *AddRecurringTask200Response) SetUseSequentialMode(v bool)`

SetUseSequentialMode sets UseSequentialMode field to given value.

### HasUseSequentialMode

`func (o *AddRecurringTask200Response) HasUseSequentialMode() bool`

HasUseSequentialMode returns a boolean if a field has been set.

### GetSecurityLevel

`func (o *AddRecurringTask200Response) GetSecurityLevel() EnumrecurringTaskSecurityLevelProp`

GetSecurityLevel returns the SecurityLevel field if non-nil, zero value otherwise.

### GetSecurityLevelOk

`func (o *AddRecurringTask200Response) GetSecurityLevelOk() (*EnumrecurringTaskSecurityLevelProp, bool)`

GetSecurityLevelOk returns a tuple with the SecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLevel

`func (o *AddRecurringTask200Response) SetSecurityLevel(v EnumrecurringTaskSecurityLevelProp)`

SetSecurityLevel sets SecurityLevel field to given value.

### HasSecurityLevel

`func (o *AddRecurringTask200Response) HasSecurityLevel() bool`

HasSecurityLevel returns a boolean if a field has been set.

### GetJstackCount

`func (o *AddRecurringTask200Response) GetJstackCount() int64`

GetJstackCount returns the JstackCount field if non-nil, zero value otherwise.

### GetJstackCountOk

`func (o *AddRecurringTask200Response) GetJstackCountOk() (*int64, bool)`

GetJstackCountOk returns a tuple with the JstackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJstackCount

`func (o *AddRecurringTask200Response) SetJstackCount(v int64)`

SetJstackCount sets JstackCount field to given value.

### HasJstackCount

`func (o *AddRecurringTask200Response) HasJstackCount() bool`

HasJstackCount returns a boolean if a field has been set.

### GetReportCount

`func (o *AddRecurringTask200Response) GetReportCount() int64`

GetReportCount returns the ReportCount field if non-nil, zero value otherwise.

### GetReportCountOk

`func (o *AddRecurringTask200Response) GetReportCountOk() (*int64, bool)`

GetReportCountOk returns a tuple with the ReportCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportCount

`func (o *AddRecurringTask200Response) SetReportCount(v int64)`

SetReportCount sets ReportCount field to given value.

### HasReportCount

`func (o *AddRecurringTask200Response) HasReportCount() bool`

HasReportCount returns a boolean if a field has been set.

### GetReportIntervalSeconds

`func (o *AddRecurringTask200Response) GetReportIntervalSeconds() int64`

GetReportIntervalSeconds returns the ReportIntervalSeconds field if non-nil, zero value otherwise.

### GetReportIntervalSecondsOk

`func (o *AddRecurringTask200Response) GetReportIntervalSecondsOk() (*int64, bool)`

GetReportIntervalSecondsOk returns a tuple with the ReportIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportIntervalSeconds

`func (o *AddRecurringTask200Response) SetReportIntervalSeconds(v int64)`

SetReportIntervalSeconds sets ReportIntervalSeconds field to given value.

### HasReportIntervalSeconds

`func (o *AddRecurringTask200Response) HasReportIntervalSeconds() bool`

HasReportIntervalSeconds returns a boolean if a field has been set.

### GetLogDuration

`func (o *AddRecurringTask200Response) GetLogDuration() string`

GetLogDuration returns the LogDuration field if non-nil, zero value otherwise.

### GetLogDurationOk

`func (o *AddRecurringTask200Response) GetLogDurationOk() (*string, bool)`

GetLogDurationOk returns a tuple with the LogDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDuration

`func (o *AddRecurringTask200Response) SetLogDuration(v string)`

SetLogDuration sets LogDuration field to given value.

### HasLogDuration

`func (o *AddRecurringTask200Response) HasLogDuration() bool`

HasLogDuration returns a boolean if a field has been set.

### GetLogFileHeadCollectionSize

`func (o *AddRecurringTask200Response) GetLogFileHeadCollectionSize() string`

GetLogFileHeadCollectionSize returns the LogFileHeadCollectionSize field if non-nil, zero value otherwise.

### GetLogFileHeadCollectionSizeOk

`func (o *AddRecurringTask200Response) GetLogFileHeadCollectionSizeOk() (*string, bool)`

GetLogFileHeadCollectionSizeOk returns a tuple with the LogFileHeadCollectionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFileHeadCollectionSize

`func (o *AddRecurringTask200Response) SetLogFileHeadCollectionSize(v string)`

SetLogFileHeadCollectionSize sets LogFileHeadCollectionSize field to given value.

### HasLogFileHeadCollectionSize

`func (o *AddRecurringTask200Response) HasLogFileHeadCollectionSize() bool`

HasLogFileHeadCollectionSize returns a boolean if a field has been set.

### GetLogFileTailCollectionSize

`func (o *AddRecurringTask200Response) GetLogFileTailCollectionSize() string`

GetLogFileTailCollectionSize returns the LogFileTailCollectionSize field if non-nil, zero value otherwise.

### GetLogFileTailCollectionSizeOk

`func (o *AddRecurringTask200Response) GetLogFileTailCollectionSizeOk() (*string, bool)`

GetLogFileTailCollectionSizeOk returns a tuple with the LogFileTailCollectionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFileTailCollectionSize

`func (o *AddRecurringTask200Response) SetLogFileTailCollectionSize(v string)`

SetLogFileTailCollectionSize sets LogFileTailCollectionSize field to given value.

### HasLogFileTailCollectionSize

`func (o *AddRecurringTask200Response) HasLogFileTailCollectionSize() bool`

HasLogFileTailCollectionSize returns a boolean if a field has been set.

### GetComment

`func (o *AddRecurringTask200Response) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *AddRecurringTask200Response) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *AddRecurringTask200Response) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *AddRecurringTask200Response) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetRetainPreviousSupportDataArchiveCount

`func (o *AddRecurringTask200Response) GetRetainPreviousSupportDataArchiveCount() int64`

GetRetainPreviousSupportDataArchiveCount returns the RetainPreviousSupportDataArchiveCount field if non-nil, zero value otherwise.

### GetRetainPreviousSupportDataArchiveCountOk

`func (o *AddRecurringTask200Response) GetRetainPreviousSupportDataArchiveCountOk() (*int64, bool)`

GetRetainPreviousSupportDataArchiveCountOk returns a tuple with the RetainPreviousSupportDataArchiveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousSupportDataArchiveCount

`func (o *AddRecurringTask200Response) SetRetainPreviousSupportDataArchiveCount(v int64)`

SetRetainPreviousSupportDataArchiveCount sets RetainPreviousSupportDataArchiveCount field to given value.

### HasRetainPreviousSupportDataArchiveCount

`func (o *AddRecurringTask200Response) HasRetainPreviousSupportDataArchiveCount() bool`

HasRetainPreviousSupportDataArchiveCount returns a boolean if a field has been set.

### GetRetainPreviousSupportDataArchiveAge

`func (o *AddRecurringTask200Response) GetRetainPreviousSupportDataArchiveAge() string`

GetRetainPreviousSupportDataArchiveAge returns the RetainPreviousSupportDataArchiveAge field if non-nil, zero value otherwise.

### GetRetainPreviousSupportDataArchiveAgeOk

`func (o *AddRecurringTask200Response) GetRetainPreviousSupportDataArchiveAgeOk() (*string, bool)`

GetRetainPreviousSupportDataArchiveAgeOk returns a tuple with the RetainPreviousSupportDataArchiveAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousSupportDataArchiveAge

`func (o *AddRecurringTask200Response) SetRetainPreviousSupportDataArchiveAge(v string)`

SetRetainPreviousSupportDataArchiveAge sets RetainPreviousSupportDataArchiveAge field to given value.

### HasRetainPreviousSupportDataArchiveAge

`func (o *AddRecurringTask200Response) HasRetainPreviousSupportDataArchiveAge() bool`

HasRetainPreviousSupportDataArchiveAge returns a boolean if a field has been set.

### GetLdifDirectory

`func (o *AddRecurringTask200Response) GetLdifDirectory() string`

GetLdifDirectory returns the LdifDirectory field if non-nil, zero value otherwise.

### GetLdifDirectoryOk

`func (o *AddRecurringTask200Response) GetLdifDirectoryOk() (*string, bool)`

GetLdifDirectoryOk returns a tuple with the LdifDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdifDirectory

`func (o *AddRecurringTask200Response) SetLdifDirectory(v string)`

SetLdifDirectory sets LdifDirectory field to given value.


### GetBackendID

`func (o *AddRecurringTask200Response) GetBackendID() []string`

GetBackendID returns the BackendID field if non-nil, zero value otherwise.

### GetBackendIDOk

`func (o *AddRecurringTask200Response) GetBackendIDOk() (*[]string, bool)`

GetBackendIDOk returns a tuple with the BackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendID

`func (o *AddRecurringTask200Response) SetBackendID(v []string)`

SetBackendID sets BackendID field to given value.

### HasBackendID

`func (o *AddRecurringTask200Response) HasBackendID() bool`

HasBackendID returns a boolean if a field has been set.

### GetExcludeBackendID

`func (o *AddRecurringTask200Response) GetExcludeBackendID() []string`

GetExcludeBackendID returns the ExcludeBackendID field if non-nil, zero value otherwise.

### GetExcludeBackendIDOk

`func (o *AddRecurringTask200Response) GetExcludeBackendIDOk() (*[]string, bool)`

GetExcludeBackendIDOk returns a tuple with the ExcludeBackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeBackendID

`func (o *AddRecurringTask200Response) SetExcludeBackendID(v []string)`

SetExcludeBackendID sets ExcludeBackendID field to given value.

### HasExcludeBackendID

`func (o *AddRecurringTask200Response) HasExcludeBackendID() bool`

HasExcludeBackendID returns a boolean if a field has been set.

### GetRetainPreviousLDIFExportCount

`func (o *AddRecurringTask200Response) GetRetainPreviousLDIFExportCount() int64`

GetRetainPreviousLDIFExportCount returns the RetainPreviousLDIFExportCount field if non-nil, zero value otherwise.

### GetRetainPreviousLDIFExportCountOk

`func (o *AddRecurringTask200Response) GetRetainPreviousLDIFExportCountOk() (*int64, bool)`

GetRetainPreviousLDIFExportCountOk returns a tuple with the RetainPreviousLDIFExportCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousLDIFExportCount

`func (o *AddRecurringTask200Response) SetRetainPreviousLDIFExportCount(v int64)`

SetRetainPreviousLDIFExportCount sets RetainPreviousLDIFExportCount field to given value.

### HasRetainPreviousLDIFExportCount

`func (o *AddRecurringTask200Response) HasRetainPreviousLDIFExportCount() bool`

HasRetainPreviousLDIFExportCount returns a boolean if a field has been set.

### GetRetainPreviousLDIFExportAge

`func (o *AddRecurringTask200Response) GetRetainPreviousLDIFExportAge() string`

GetRetainPreviousLDIFExportAge returns the RetainPreviousLDIFExportAge field if non-nil, zero value otherwise.

### GetRetainPreviousLDIFExportAgeOk

`func (o *AddRecurringTask200Response) GetRetainPreviousLDIFExportAgeOk() (*string, bool)`

GetRetainPreviousLDIFExportAgeOk returns a tuple with the RetainPreviousLDIFExportAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousLDIFExportAge

`func (o *AddRecurringTask200Response) SetRetainPreviousLDIFExportAge(v string)`

SetRetainPreviousLDIFExportAge sets RetainPreviousLDIFExportAge field to given value.

### HasRetainPreviousLDIFExportAge

`func (o *AddRecurringTask200Response) HasRetainPreviousLDIFExportAge() bool`

HasRetainPreviousLDIFExportAge returns a boolean if a field has been set.

### GetPostLDIFExportTaskProcessor

`func (o *AddRecurringTask200Response) GetPostLDIFExportTaskProcessor() []string`

GetPostLDIFExportTaskProcessor returns the PostLDIFExportTaskProcessor field if non-nil, zero value otherwise.

### GetPostLDIFExportTaskProcessorOk

`func (o *AddRecurringTask200Response) GetPostLDIFExportTaskProcessorOk() (*[]string, bool)`

GetPostLDIFExportTaskProcessorOk returns a tuple with the PostLDIFExportTaskProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostLDIFExportTaskProcessor

`func (o *AddRecurringTask200Response) SetPostLDIFExportTaskProcessor(v []string)`

SetPostLDIFExportTaskProcessor sets PostLDIFExportTaskProcessor field to given value.

### HasPostLDIFExportTaskProcessor

`func (o *AddRecurringTask200Response) HasPostLDIFExportTaskProcessor() bool`

HasPostLDIFExportTaskProcessor returns a boolean if a field has been set.

### GetBaseOutputDirectory

`func (o *AddRecurringTask200Response) GetBaseOutputDirectory() string`

GetBaseOutputDirectory returns the BaseOutputDirectory field if non-nil, zero value otherwise.

### GetBaseOutputDirectoryOk

`func (o *AddRecurringTask200Response) GetBaseOutputDirectoryOk() (*string, bool)`

GetBaseOutputDirectoryOk returns a tuple with the BaseOutputDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseOutputDirectory

`func (o *AddRecurringTask200Response) SetBaseOutputDirectory(v string)`

SetBaseOutputDirectory sets BaseOutputDirectory field to given value.


### GetDataSecurityAuditor

`func (o *AddRecurringTask200Response) GetDataSecurityAuditor() []string`

GetDataSecurityAuditor returns the DataSecurityAuditor field if non-nil, zero value otherwise.

### GetDataSecurityAuditorOk

`func (o *AddRecurringTask200Response) GetDataSecurityAuditorOk() (*[]string, bool)`

GetDataSecurityAuditorOk returns a tuple with the DataSecurityAuditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSecurityAuditor

`func (o *AddRecurringTask200Response) SetDataSecurityAuditor(v []string)`

SetDataSecurityAuditor sets DataSecurityAuditor field to given value.

### HasDataSecurityAuditor

`func (o *AddRecurringTask200Response) HasDataSecurityAuditor() bool`

HasDataSecurityAuditor returns a boolean if a field has been set.

### GetBackend

`func (o *AddRecurringTask200Response) GetBackend() []string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *AddRecurringTask200Response) GetBackendOk() (*[]string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *AddRecurringTask200Response) SetBackend(v []string)`

SetBackend sets Backend field to given value.

### HasBackend

`func (o *AddRecurringTask200Response) HasBackend() bool`

HasBackend returns a boolean if a field has been set.

### GetIncludeFilter

`func (o *AddRecurringTask200Response) GetIncludeFilter() []string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *AddRecurringTask200Response) GetIncludeFilterOk() (*[]string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *AddRecurringTask200Response) SetIncludeFilter(v []string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *AddRecurringTask200Response) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetRetainPreviousReportCount

`func (o *AddRecurringTask200Response) GetRetainPreviousReportCount() int64`

GetRetainPreviousReportCount returns the RetainPreviousReportCount field if non-nil, zero value otherwise.

### GetRetainPreviousReportCountOk

`func (o *AddRecurringTask200Response) GetRetainPreviousReportCountOk() (*int64, bool)`

GetRetainPreviousReportCountOk returns a tuple with the RetainPreviousReportCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousReportCount

`func (o *AddRecurringTask200Response) SetRetainPreviousReportCount(v int64)`

SetRetainPreviousReportCount sets RetainPreviousReportCount field to given value.

### HasRetainPreviousReportCount

`func (o *AddRecurringTask200Response) HasRetainPreviousReportCount() bool`

HasRetainPreviousReportCount returns a boolean if a field has been set.

### GetRetainPreviousReportAge

`func (o *AddRecurringTask200Response) GetRetainPreviousReportAge() string`

GetRetainPreviousReportAge returns the RetainPreviousReportAge field if non-nil, zero value otherwise.

### GetRetainPreviousReportAgeOk

`func (o *AddRecurringTask200Response) GetRetainPreviousReportAgeOk() (*string, bool)`

GetRetainPreviousReportAgeOk returns a tuple with the RetainPreviousReportAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousReportAge

`func (o *AddRecurringTask200Response) SetRetainPreviousReportAge(v string)`

SetRetainPreviousReportAge sets RetainPreviousReportAge field to given value.

### HasRetainPreviousReportAge

`func (o *AddRecurringTask200Response) HasRetainPreviousReportAge() bool`

HasRetainPreviousReportAge returns a boolean if a field has been set.

### GetCommandPath

`func (o *AddRecurringTask200Response) GetCommandPath() string`

GetCommandPath returns the CommandPath field if non-nil, zero value otherwise.

### GetCommandPathOk

`func (o *AddRecurringTask200Response) GetCommandPathOk() (*string, bool)`

GetCommandPathOk returns a tuple with the CommandPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandPath

`func (o *AddRecurringTask200Response) SetCommandPath(v string)`

SetCommandPath sets CommandPath field to given value.


### GetCommandArguments

`func (o *AddRecurringTask200Response) GetCommandArguments() string`

GetCommandArguments returns the CommandArguments field if non-nil, zero value otherwise.

### GetCommandArgumentsOk

`func (o *AddRecurringTask200Response) GetCommandArgumentsOk() (*string, bool)`

GetCommandArgumentsOk returns a tuple with the CommandArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandArguments

`func (o *AddRecurringTask200Response) SetCommandArguments(v string)`

SetCommandArguments sets CommandArguments field to given value.

### HasCommandArguments

`func (o *AddRecurringTask200Response) HasCommandArguments() bool`

HasCommandArguments returns a boolean if a field has been set.

### GetCommandOutputFileBaseName

`func (o *AddRecurringTask200Response) GetCommandOutputFileBaseName() string`

GetCommandOutputFileBaseName returns the CommandOutputFileBaseName field if non-nil, zero value otherwise.

### GetCommandOutputFileBaseNameOk

`func (o *AddRecurringTask200Response) GetCommandOutputFileBaseNameOk() (*string, bool)`

GetCommandOutputFileBaseNameOk returns a tuple with the CommandOutputFileBaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandOutputFileBaseName

`func (o *AddRecurringTask200Response) SetCommandOutputFileBaseName(v string)`

SetCommandOutputFileBaseName sets CommandOutputFileBaseName field to given value.

### HasCommandOutputFileBaseName

`func (o *AddRecurringTask200Response) HasCommandOutputFileBaseName() bool`

HasCommandOutputFileBaseName returns a boolean if a field has been set.

### GetRetainPreviousOutputFileCount

`func (o *AddRecurringTask200Response) GetRetainPreviousOutputFileCount() int64`

GetRetainPreviousOutputFileCount returns the RetainPreviousOutputFileCount field if non-nil, zero value otherwise.

### GetRetainPreviousOutputFileCountOk

`func (o *AddRecurringTask200Response) GetRetainPreviousOutputFileCountOk() (*int64, bool)`

GetRetainPreviousOutputFileCountOk returns a tuple with the RetainPreviousOutputFileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousOutputFileCount

`func (o *AddRecurringTask200Response) SetRetainPreviousOutputFileCount(v int64)`

SetRetainPreviousOutputFileCount sets RetainPreviousOutputFileCount field to given value.

### HasRetainPreviousOutputFileCount

`func (o *AddRecurringTask200Response) HasRetainPreviousOutputFileCount() bool`

HasRetainPreviousOutputFileCount returns a boolean if a field has been set.

### GetRetainPreviousOutputFileAge

`func (o *AddRecurringTask200Response) GetRetainPreviousOutputFileAge() string`

GetRetainPreviousOutputFileAge returns the RetainPreviousOutputFileAge field if non-nil, zero value otherwise.

### GetRetainPreviousOutputFileAgeOk

`func (o *AddRecurringTask200Response) GetRetainPreviousOutputFileAgeOk() (*string, bool)`

GetRetainPreviousOutputFileAgeOk returns a tuple with the RetainPreviousOutputFileAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousOutputFileAge

`func (o *AddRecurringTask200Response) SetRetainPreviousOutputFileAge(v string)`

SetRetainPreviousOutputFileAge sets RetainPreviousOutputFileAge field to given value.

### HasRetainPreviousOutputFileAge

`func (o *AddRecurringTask200Response) HasRetainPreviousOutputFileAge() bool`

HasRetainPreviousOutputFileAge returns a boolean if a field has been set.

### GetLogCommandOutput

`func (o *AddRecurringTask200Response) GetLogCommandOutput() bool`

GetLogCommandOutput returns the LogCommandOutput field if non-nil, zero value otherwise.

### GetLogCommandOutputOk

`func (o *AddRecurringTask200Response) GetLogCommandOutputOk() (*bool, bool)`

GetLogCommandOutputOk returns a tuple with the LogCommandOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogCommandOutput

`func (o *AddRecurringTask200Response) SetLogCommandOutput(v bool)`

SetLogCommandOutput sets LogCommandOutput field to given value.

### HasLogCommandOutput

`func (o *AddRecurringTask200Response) HasLogCommandOutput() bool`

HasLogCommandOutput returns a boolean if a field has been set.

### GetTaskCompletionStateForNonzeroExitCode

`func (o *AddRecurringTask200Response) GetTaskCompletionStateForNonzeroExitCode() EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp`

GetTaskCompletionStateForNonzeroExitCode returns the TaskCompletionStateForNonzeroExitCode field if non-nil, zero value otherwise.

### GetTaskCompletionStateForNonzeroExitCodeOk

`func (o *AddRecurringTask200Response) GetTaskCompletionStateForNonzeroExitCodeOk() (*EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp, bool)`

GetTaskCompletionStateForNonzeroExitCodeOk returns a tuple with the TaskCompletionStateForNonzeroExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCompletionStateForNonzeroExitCode

`func (o *AddRecurringTask200Response) SetTaskCompletionStateForNonzeroExitCode(v EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp)`

SetTaskCompletionStateForNonzeroExitCode sets TaskCompletionStateForNonzeroExitCode field to given value.

### HasTaskCompletionStateForNonzeroExitCode

`func (o *AddRecurringTask200Response) HasTaskCompletionStateForNonzeroExitCode() bool`

HasTaskCompletionStateForNonzeroExitCode returns a boolean if a field has been set.

### GetWorkingDirectory

`func (o *AddRecurringTask200Response) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *AddRecurringTask200Response) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *AddRecurringTask200Response) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.

### HasWorkingDirectory

`func (o *AddRecurringTask200Response) HasWorkingDirectory() bool`

HasWorkingDirectory returns a boolean if a field has been set.

### GetTargetDirectory

`func (o *AddRecurringTask200Response) GetTargetDirectory() string`

GetTargetDirectory returns the TargetDirectory field if non-nil, zero value otherwise.

### GetTargetDirectoryOk

`func (o *AddRecurringTask200Response) GetTargetDirectoryOk() (*string, bool)`

GetTargetDirectoryOk returns a tuple with the TargetDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDirectory

`func (o *AddRecurringTask200Response) SetTargetDirectory(v string)`

SetTargetDirectory sets TargetDirectory field to given value.


### GetFilenamePattern

`func (o *AddRecurringTask200Response) GetFilenamePattern() string`

GetFilenamePattern returns the FilenamePattern field if non-nil, zero value otherwise.

### GetFilenamePatternOk

`func (o *AddRecurringTask200Response) GetFilenamePatternOk() (*string, bool)`

GetFilenamePatternOk returns a tuple with the FilenamePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilenamePattern

`func (o *AddRecurringTask200Response) SetFilenamePattern(v string)`

SetFilenamePattern sets FilenamePattern field to given value.


### GetTimestampFormat

`func (o *AddRecurringTask200Response) GetTimestampFormat() EnumrecurringTaskTimestampFormatProp`

GetTimestampFormat returns the TimestampFormat field if non-nil, zero value otherwise.

### GetTimestampFormatOk

`func (o *AddRecurringTask200Response) GetTimestampFormatOk() (*EnumrecurringTaskTimestampFormatProp, bool)`

GetTimestampFormatOk returns a tuple with the TimestampFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampFormat

`func (o *AddRecurringTask200Response) SetTimestampFormat(v EnumrecurringTaskTimestampFormatProp)`

SetTimestampFormat sets TimestampFormat field to given value.


### GetRetainFileCount

`func (o *AddRecurringTask200Response) GetRetainFileCount() int64`

GetRetainFileCount returns the RetainFileCount field if non-nil, zero value otherwise.

### GetRetainFileCountOk

`func (o *AddRecurringTask200Response) GetRetainFileCountOk() (*int64, bool)`

GetRetainFileCountOk returns a tuple with the RetainFileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainFileCount

`func (o *AddRecurringTask200Response) SetRetainFileCount(v int64)`

SetRetainFileCount sets RetainFileCount field to given value.

### HasRetainFileCount

`func (o *AddRecurringTask200Response) HasRetainFileCount() bool`

HasRetainFileCount returns a boolean if a field has been set.

### GetRetainFileAge

`func (o *AddRecurringTask200Response) GetRetainFileAge() string`

GetRetainFileAge returns the RetainFileAge field if non-nil, zero value otherwise.

### GetRetainFileAgeOk

`func (o *AddRecurringTask200Response) GetRetainFileAgeOk() (*string, bool)`

GetRetainFileAgeOk returns a tuple with the RetainFileAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainFileAge

`func (o *AddRecurringTask200Response) SetRetainFileAge(v string)`

SetRetainFileAge sets RetainFileAge field to given value.

### HasRetainFileAge

`func (o *AddRecurringTask200Response) HasRetainFileAge() bool`

HasRetainFileAge returns a boolean if a field has been set.

### GetRetainAggregateFileSize

`func (o *AddRecurringTask200Response) GetRetainAggregateFileSize() string`

GetRetainAggregateFileSize returns the RetainAggregateFileSize field if non-nil, zero value otherwise.

### GetRetainAggregateFileSizeOk

`func (o *AddRecurringTask200Response) GetRetainAggregateFileSizeOk() (*string, bool)`

GetRetainAggregateFileSizeOk returns a tuple with the RetainAggregateFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainAggregateFileSize

`func (o *AddRecurringTask200Response) SetRetainAggregateFileSize(v string)`

SetRetainAggregateFileSize sets RetainAggregateFileSize field to given value.

### HasRetainAggregateFileSize

`func (o *AddRecurringTask200Response) HasRetainAggregateFileSize() bool`

HasRetainAggregateFileSize returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddRecurringTask200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddRecurringTask200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddRecurringTask200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddRecurringTask200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddRecurringTask200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddRecurringTask200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddRecurringTask200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


