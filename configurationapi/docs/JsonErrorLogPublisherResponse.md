# JsonErrorLogPublisherResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumjsonErrorLogPublisherSchemaUrn**](EnumjsonErrorLogPublisherSchemaUrn.md) |  | 
**LogFile** | **string** | The file name to use for the log files generated by the JSON Error Log Publisher. The path to the file can be specified either as relative to the server root or as an absolute path. | 
**LogFilePermissions** | **string** | The UNIX permissions of the log files created by this JSON Error Log Publisher. | 
**RotationPolicy** | **[]string** | The rotation policy to use for the JSON Error Log Publisher . | 
**RotationListener** | Pointer to **[]string** | A listener that should be notified whenever a log file is rotated out of service. | [optional] 
**RetentionPolicy** | **[]string** | The retention policy to use for the JSON Error Log Publisher . | 
**CompressionMechanism** | Pointer to [**EnumlogPublisherCompressionMechanismProp**](EnumlogPublisherCompressionMechanismProp.md) |  | [optional] 
**SignLog** | Pointer to **bool** | Indicates whether the log should be cryptographically signed so that the log content cannot be altered in an undetectable manner. | [optional] 
**EncryptLog** | Pointer to **bool** | Indicates whether log files should be encrypted so that their content is not available to unauthorized users. | [optional] 
**EncryptionSettingsDefinitionID** | Pointer to **string** | Specifies the ID of the encryption settings definition that should be used to encrypt the data. If this is not provided, the server&#39;s preferred encryption settings definition will be used. The \&quot;encryption-settings list\&quot; command can be used to obtain a list of the encryption settings definitions available in the server. | [optional] 
**Append** | Pointer to **bool** | Specifies whether to append to existing log files. | [optional] 
**Asynchronous** | **bool** | Indicates whether the JSON Error Log Publisher will publish records asynchronously. | 
**AutoFlush** | Pointer to **bool** | Specifies whether to flush the writer after every log record. | [optional] 
**BufferSize** | Pointer to **string** | Specifies the log file buffer size. | [optional] 
**QueueSize** | Pointer to **int64** | The maximum number of log records that can be stored in the asynchronous queue. | [optional] 
**TimeInterval** | Pointer to **string** | Specifies the interval at which to check whether the log files need to be rotated. | [optional] 
**WriteMultiLineMessages** | Pointer to **bool** | Indicates whether the JSON objects should be formatted to span multiple lines with a single element on each line. The multi-line format is potentially more user friendly (if administrators may need to look at the log files), but each message will be larger because of the additional spaces and end-of-line markers. | [optional] 
**IncludeProductName** | Pointer to **bool** | Indicates whether log messages should include the product name for the Directory Server. | [optional] 
**IncludeInstanceName** | Pointer to **bool** | Indicates whether log messages should include the instance name for the Directory Server. | [optional] 
**IncludeStartupID** | Pointer to **bool** | Indicates whether log messages should include the startup ID for the Directory Server, which is a value assigned to the server instance at startup and may be used to identify when the server has been restarted. | [optional] 
**IncludeThreadID** | Pointer to **bool** | Indicates whether log messages should include the thread ID for the Directory Server in each log message. This ID can be used to correlate log messages from the same thread within a single log as well as generated by the same thread across different types of log files. More information about the thread with a specific ID can be obtained using the cn&#x3D;JVM Stack Trace,cn&#x3D;monitor entry. | [optional] 
**GenerifyMessageStringsWhenPossible** | Pointer to **bool** | Indicates whether to use the generified version of the log message string (which may use placeholders like %s for a string or %d for an integer), rather than the version of the message with those placeholders replaced with specific values that would normally be written to the log. | [optional] 
**DefaultSeverity** | Pointer to [**[]EnumlogPublisherDefaultSeverityProp**](EnumlogPublisherDefaultSeverityProp.md) |  | [optional] 
**OverrideSeverity** | Pointer to **[]string** | Specifies the override severity levels for the logger based on the category of the messages. | [optional] 
**Description** | Pointer to **string** | A description for this Log Publisher | [optional] 
**Enabled** | **bool** | Indicates whether the Log Publisher is enabled for use. | 
**LoggingErrorBehavior** | Pointer to [**EnumlogPublisherLoggingErrorBehaviorProp**](EnumlogPublisherLoggingErrorBehaviorProp.md) |  | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Log Publisher | 

## Methods

### NewJsonErrorLogPublisherResponse

`func NewJsonErrorLogPublisherResponse(schemas []EnumjsonErrorLogPublisherSchemaUrn, logFile string, logFilePermissions string, rotationPolicy []string, retentionPolicy []string, asynchronous bool, enabled bool, id string, ) *JsonErrorLogPublisherResponse`

NewJsonErrorLogPublisherResponse instantiates a new JsonErrorLogPublisherResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonErrorLogPublisherResponseWithDefaults

`func NewJsonErrorLogPublisherResponseWithDefaults() *JsonErrorLogPublisherResponse`

NewJsonErrorLogPublisherResponseWithDefaults instantiates a new JsonErrorLogPublisherResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *JsonErrorLogPublisherResponse) GetSchemas() []EnumjsonErrorLogPublisherSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *JsonErrorLogPublisherResponse) GetSchemasOk() (*[]EnumjsonErrorLogPublisherSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *JsonErrorLogPublisherResponse) SetSchemas(v []EnumjsonErrorLogPublisherSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetLogFile

`func (o *JsonErrorLogPublisherResponse) GetLogFile() string`

GetLogFile returns the LogFile field if non-nil, zero value otherwise.

### GetLogFileOk

`func (o *JsonErrorLogPublisherResponse) GetLogFileOk() (*string, bool)`

GetLogFileOk returns a tuple with the LogFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFile

`func (o *JsonErrorLogPublisherResponse) SetLogFile(v string)`

SetLogFile sets LogFile field to given value.


### GetLogFilePermissions

`func (o *JsonErrorLogPublisherResponse) GetLogFilePermissions() string`

GetLogFilePermissions returns the LogFilePermissions field if non-nil, zero value otherwise.

### GetLogFilePermissionsOk

`func (o *JsonErrorLogPublisherResponse) GetLogFilePermissionsOk() (*string, bool)`

GetLogFilePermissionsOk returns a tuple with the LogFilePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFilePermissions

`func (o *JsonErrorLogPublisherResponse) SetLogFilePermissions(v string)`

SetLogFilePermissions sets LogFilePermissions field to given value.


### GetRotationPolicy

`func (o *JsonErrorLogPublisherResponse) GetRotationPolicy() []string`

GetRotationPolicy returns the RotationPolicy field if non-nil, zero value otherwise.

### GetRotationPolicyOk

`func (o *JsonErrorLogPublisherResponse) GetRotationPolicyOk() (*[]string, bool)`

GetRotationPolicyOk returns a tuple with the RotationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationPolicy

`func (o *JsonErrorLogPublisherResponse) SetRotationPolicy(v []string)`

SetRotationPolicy sets RotationPolicy field to given value.


### GetRotationListener

`func (o *JsonErrorLogPublisherResponse) GetRotationListener() []string`

GetRotationListener returns the RotationListener field if non-nil, zero value otherwise.

### GetRotationListenerOk

`func (o *JsonErrorLogPublisherResponse) GetRotationListenerOk() (*[]string, bool)`

GetRotationListenerOk returns a tuple with the RotationListener field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationListener

`func (o *JsonErrorLogPublisherResponse) SetRotationListener(v []string)`

SetRotationListener sets RotationListener field to given value.

### HasRotationListener

`func (o *JsonErrorLogPublisherResponse) HasRotationListener() bool`

HasRotationListener returns a boolean if a field has been set.

### GetRetentionPolicy

`func (o *JsonErrorLogPublisherResponse) GetRetentionPolicy() []string`

GetRetentionPolicy returns the RetentionPolicy field if non-nil, zero value otherwise.

### GetRetentionPolicyOk

`func (o *JsonErrorLogPublisherResponse) GetRetentionPolicyOk() (*[]string, bool)`

GetRetentionPolicyOk returns a tuple with the RetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicy

`func (o *JsonErrorLogPublisherResponse) SetRetentionPolicy(v []string)`

SetRetentionPolicy sets RetentionPolicy field to given value.


### GetCompressionMechanism

`func (o *JsonErrorLogPublisherResponse) GetCompressionMechanism() EnumlogPublisherCompressionMechanismProp`

GetCompressionMechanism returns the CompressionMechanism field if non-nil, zero value otherwise.

### GetCompressionMechanismOk

`func (o *JsonErrorLogPublisherResponse) GetCompressionMechanismOk() (*EnumlogPublisherCompressionMechanismProp, bool)`

GetCompressionMechanismOk returns a tuple with the CompressionMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionMechanism

`func (o *JsonErrorLogPublisherResponse) SetCompressionMechanism(v EnumlogPublisherCompressionMechanismProp)`

SetCompressionMechanism sets CompressionMechanism field to given value.

### HasCompressionMechanism

`func (o *JsonErrorLogPublisherResponse) HasCompressionMechanism() bool`

HasCompressionMechanism returns a boolean if a field has been set.

### GetSignLog

`func (o *JsonErrorLogPublisherResponse) GetSignLog() bool`

GetSignLog returns the SignLog field if non-nil, zero value otherwise.

### GetSignLogOk

`func (o *JsonErrorLogPublisherResponse) GetSignLogOk() (*bool, bool)`

GetSignLogOk returns a tuple with the SignLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignLog

`func (o *JsonErrorLogPublisherResponse) SetSignLog(v bool)`

SetSignLog sets SignLog field to given value.

### HasSignLog

`func (o *JsonErrorLogPublisherResponse) HasSignLog() bool`

HasSignLog returns a boolean if a field has been set.

### GetEncryptLog

`func (o *JsonErrorLogPublisherResponse) GetEncryptLog() bool`

GetEncryptLog returns the EncryptLog field if non-nil, zero value otherwise.

### GetEncryptLogOk

`func (o *JsonErrorLogPublisherResponse) GetEncryptLogOk() (*bool, bool)`

GetEncryptLogOk returns a tuple with the EncryptLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptLog

`func (o *JsonErrorLogPublisherResponse) SetEncryptLog(v bool)`

SetEncryptLog sets EncryptLog field to given value.

### HasEncryptLog

`func (o *JsonErrorLogPublisherResponse) HasEncryptLog() bool`

HasEncryptLog returns a boolean if a field has been set.

### GetEncryptionSettingsDefinitionID

`func (o *JsonErrorLogPublisherResponse) GetEncryptionSettingsDefinitionID() string`

GetEncryptionSettingsDefinitionID returns the EncryptionSettingsDefinitionID field if non-nil, zero value otherwise.

### GetEncryptionSettingsDefinitionIDOk

`func (o *JsonErrorLogPublisherResponse) GetEncryptionSettingsDefinitionIDOk() (*string, bool)`

GetEncryptionSettingsDefinitionIDOk returns a tuple with the EncryptionSettingsDefinitionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionSettingsDefinitionID

`func (o *JsonErrorLogPublisherResponse) SetEncryptionSettingsDefinitionID(v string)`

SetEncryptionSettingsDefinitionID sets EncryptionSettingsDefinitionID field to given value.

### HasEncryptionSettingsDefinitionID

`func (o *JsonErrorLogPublisherResponse) HasEncryptionSettingsDefinitionID() bool`

HasEncryptionSettingsDefinitionID returns a boolean if a field has been set.

### GetAppend

`func (o *JsonErrorLogPublisherResponse) GetAppend() bool`

GetAppend returns the Append field if non-nil, zero value otherwise.

### GetAppendOk

`func (o *JsonErrorLogPublisherResponse) GetAppendOk() (*bool, bool)`

GetAppendOk returns a tuple with the Append field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppend

`func (o *JsonErrorLogPublisherResponse) SetAppend(v bool)`

SetAppend sets Append field to given value.

### HasAppend

`func (o *JsonErrorLogPublisherResponse) HasAppend() bool`

HasAppend returns a boolean if a field has been set.

### GetAsynchronous

`func (o *JsonErrorLogPublisherResponse) GetAsynchronous() bool`

GetAsynchronous returns the Asynchronous field if non-nil, zero value otherwise.

### GetAsynchronousOk

`func (o *JsonErrorLogPublisherResponse) GetAsynchronousOk() (*bool, bool)`

GetAsynchronousOk returns a tuple with the Asynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsynchronous

`func (o *JsonErrorLogPublisherResponse) SetAsynchronous(v bool)`

SetAsynchronous sets Asynchronous field to given value.


### GetAutoFlush

`func (o *JsonErrorLogPublisherResponse) GetAutoFlush() bool`

GetAutoFlush returns the AutoFlush field if non-nil, zero value otherwise.

### GetAutoFlushOk

`func (o *JsonErrorLogPublisherResponse) GetAutoFlushOk() (*bool, bool)`

GetAutoFlushOk returns a tuple with the AutoFlush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoFlush

`func (o *JsonErrorLogPublisherResponse) SetAutoFlush(v bool)`

SetAutoFlush sets AutoFlush field to given value.

### HasAutoFlush

`func (o *JsonErrorLogPublisherResponse) HasAutoFlush() bool`

HasAutoFlush returns a boolean if a field has been set.

### GetBufferSize

`func (o *JsonErrorLogPublisherResponse) GetBufferSize() string`

GetBufferSize returns the BufferSize field if non-nil, zero value otherwise.

### GetBufferSizeOk

`func (o *JsonErrorLogPublisherResponse) GetBufferSizeOk() (*string, bool)`

GetBufferSizeOk returns a tuple with the BufferSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferSize

`func (o *JsonErrorLogPublisherResponse) SetBufferSize(v string)`

SetBufferSize sets BufferSize field to given value.

### HasBufferSize

`func (o *JsonErrorLogPublisherResponse) HasBufferSize() bool`

HasBufferSize returns a boolean if a field has been set.

### GetQueueSize

`func (o *JsonErrorLogPublisherResponse) GetQueueSize() int64`

GetQueueSize returns the QueueSize field if non-nil, zero value otherwise.

### GetQueueSizeOk

`func (o *JsonErrorLogPublisherResponse) GetQueueSizeOk() (*int64, bool)`

GetQueueSizeOk returns a tuple with the QueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueSize

`func (o *JsonErrorLogPublisherResponse) SetQueueSize(v int64)`

SetQueueSize sets QueueSize field to given value.

### HasQueueSize

`func (o *JsonErrorLogPublisherResponse) HasQueueSize() bool`

HasQueueSize returns a boolean if a field has been set.

### GetTimeInterval

`func (o *JsonErrorLogPublisherResponse) GetTimeInterval() string`

GetTimeInterval returns the TimeInterval field if non-nil, zero value otherwise.

### GetTimeIntervalOk

`func (o *JsonErrorLogPublisherResponse) GetTimeIntervalOk() (*string, bool)`

GetTimeIntervalOk returns a tuple with the TimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInterval

`func (o *JsonErrorLogPublisherResponse) SetTimeInterval(v string)`

SetTimeInterval sets TimeInterval field to given value.

### HasTimeInterval

`func (o *JsonErrorLogPublisherResponse) HasTimeInterval() bool`

HasTimeInterval returns a boolean if a field has been set.

### GetWriteMultiLineMessages

`func (o *JsonErrorLogPublisherResponse) GetWriteMultiLineMessages() bool`

GetWriteMultiLineMessages returns the WriteMultiLineMessages field if non-nil, zero value otherwise.

### GetWriteMultiLineMessagesOk

`func (o *JsonErrorLogPublisherResponse) GetWriteMultiLineMessagesOk() (*bool, bool)`

GetWriteMultiLineMessagesOk returns a tuple with the WriteMultiLineMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteMultiLineMessages

`func (o *JsonErrorLogPublisherResponse) SetWriteMultiLineMessages(v bool)`

SetWriteMultiLineMessages sets WriteMultiLineMessages field to given value.

### HasWriteMultiLineMessages

`func (o *JsonErrorLogPublisherResponse) HasWriteMultiLineMessages() bool`

HasWriteMultiLineMessages returns a boolean if a field has been set.

### GetIncludeProductName

`func (o *JsonErrorLogPublisherResponse) GetIncludeProductName() bool`

GetIncludeProductName returns the IncludeProductName field if non-nil, zero value otherwise.

### GetIncludeProductNameOk

`func (o *JsonErrorLogPublisherResponse) GetIncludeProductNameOk() (*bool, bool)`

GetIncludeProductNameOk returns a tuple with the IncludeProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeProductName

`func (o *JsonErrorLogPublisherResponse) SetIncludeProductName(v bool)`

SetIncludeProductName sets IncludeProductName field to given value.

### HasIncludeProductName

`func (o *JsonErrorLogPublisherResponse) HasIncludeProductName() bool`

HasIncludeProductName returns a boolean if a field has been set.

### GetIncludeInstanceName

`func (o *JsonErrorLogPublisherResponse) GetIncludeInstanceName() bool`

GetIncludeInstanceName returns the IncludeInstanceName field if non-nil, zero value otherwise.

### GetIncludeInstanceNameOk

`func (o *JsonErrorLogPublisherResponse) GetIncludeInstanceNameOk() (*bool, bool)`

GetIncludeInstanceNameOk returns a tuple with the IncludeInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInstanceName

`func (o *JsonErrorLogPublisherResponse) SetIncludeInstanceName(v bool)`

SetIncludeInstanceName sets IncludeInstanceName field to given value.

### HasIncludeInstanceName

`func (o *JsonErrorLogPublisherResponse) HasIncludeInstanceName() bool`

HasIncludeInstanceName returns a boolean if a field has been set.

### GetIncludeStartupID

`func (o *JsonErrorLogPublisherResponse) GetIncludeStartupID() bool`

GetIncludeStartupID returns the IncludeStartupID field if non-nil, zero value otherwise.

### GetIncludeStartupIDOk

`func (o *JsonErrorLogPublisherResponse) GetIncludeStartupIDOk() (*bool, bool)`

GetIncludeStartupIDOk returns a tuple with the IncludeStartupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeStartupID

`func (o *JsonErrorLogPublisherResponse) SetIncludeStartupID(v bool)`

SetIncludeStartupID sets IncludeStartupID field to given value.

### HasIncludeStartupID

`func (o *JsonErrorLogPublisherResponse) HasIncludeStartupID() bool`

HasIncludeStartupID returns a boolean if a field has been set.

### GetIncludeThreadID

`func (o *JsonErrorLogPublisherResponse) GetIncludeThreadID() bool`

GetIncludeThreadID returns the IncludeThreadID field if non-nil, zero value otherwise.

### GetIncludeThreadIDOk

`func (o *JsonErrorLogPublisherResponse) GetIncludeThreadIDOk() (*bool, bool)`

GetIncludeThreadIDOk returns a tuple with the IncludeThreadID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeThreadID

`func (o *JsonErrorLogPublisherResponse) SetIncludeThreadID(v bool)`

SetIncludeThreadID sets IncludeThreadID field to given value.

### HasIncludeThreadID

`func (o *JsonErrorLogPublisherResponse) HasIncludeThreadID() bool`

HasIncludeThreadID returns a boolean if a field has been set.

### GetGenerifyMessageStringsWhenPossible

`func (o *JsonErrorLogPublisherResponse) GetGenerifyMessageStringsWhenPossible() bool`

GetGenerifyMessageStringsWhenPossible returns the GenerifyMessageStringsWhenPossible field if non-nil, zero value otherwise.

### GetGenerifyMessageStringsWhenPossibleOk

`func (o *JsonErrorLogPublisherResponse) GetGenerifyMessageStringsWhenPossibleOk() (*bool, bool)`

GetGenerifyMessageStringsWhenPossibleOk returns a tuple with the GenerifyMessageStringsWhenPossible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerifyMessageStringsWhenPossible

`func (o *JsonErrorLogPublisherResponse) SetGenerifyMessageStringsWhenPossible(v bool)`

SetGenerifyMessageStringsWhenPossible sets GenerifyMessageStringsWhenPossible field to given value.

### HasGenerifyMessageStringsWhenPossible

`func (o *JsonErrorLogPublisherResponse) HasGenerifyMessageStringsWhenPossible() bool`

HasGenerifyMessageStringsWhenPossible returns a boolean if a field has been set.

### GetDefaultSeverity

`func (o *JsonErrorLogPublisherResponse) GetDefaultSeverity() []EnumlogPublisherDefaultSeverityProp`

GetDefaultSeverity returns the DefaultSeverity field if non-nil, zero value otherwise.

### GetDefaultSeverityOk

`func (o *JsonErrorLogPublisherResponse) GetDefaultSeverityOk() (*[]EnumlogPublisherDefaultSeverityProp, bool)`

GetDefaultSeverityOk returns a tuple with the DefaultSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSeverity

`func (o *JsonErrorLogPublisherResponse) SetDefaultSeverity(v []EnumlogPublisherDefaultSeverityProp)`

SetDefaultSeverity sets DefaultSeverity field to given value.

### HasDefaultSeverity

`func (o *JsonErrorLogPublisherResponse) HasDefaultSeverity() bool`

HasDefaultSeverity returns a boolean if a field has been set.

### GetOverrideSeverity

`func (o *JsonErrorLogPublisherResponse) GetOverrideSeverity() []string`

GetOverrideSeverity returns the OverrideSeverity field if non-nil, zero value otherwise.

### GetOverrideSeverityOk

`func (o *JsonErrorLogPublisherResponse) GetOverrideSeverityOk() (*[]string, bool)`

GetOverrideSeverityOk returns a tuple with the OverrideSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSeverity

`func (o *JsonErrorLogPublisherResponse) SetOverrideSeverity(v []string)`

SetOverrideSeverity sets OverrideSeverity field to given value.

### HasOverrideSeverity

`func (o *JsonErrorLogPublisherResponse) HasOverrideSeverity() bool`

HasOverrideSeverity returns a boolean if a field has been set.

### GetDescription

`func (o *JsonErrorLogPublisherResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JsonErrorLogPublisherResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JsonErrorLogPublisherResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JsonErrorLogPublisherResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *JsonErrorLogPublisherResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *JsonErrorLogPublisherResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *JsonErrorLogPublisherResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLoggingErrorBehavior

`func (o *JsonErrorLogPublisherResponse) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp`

GetLoggingErrorBehavior returns the LoggingErrorBehavior field if non-nil, zero value otherwise.

### GetLoggingErrorBehaviorOk

`func (o *JsonErrorLogPublisherResponse) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool)`

GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingErrorBehavior

`func (o *JsonErrorLogPublisherResponse) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp)`

SetLoggingErrorBehavior sets LoggingErrorBehavior field to given value.

### HasLoggingErrorBehavior

`func (o *JsonErrorLogPublisherResponse) HasLoggingErrorBehavior() bool`

HasLoggingErrorBehavior returns a boolean if a field has been set.

### GetMeta

`func (o *JsonErrorLogPublisherResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *JsonErrorLogPublisherResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *JsonErrorLogPublisherResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *JsonErrorLogPublisherResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *JsonErrorLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *JsonErrorLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *JsonErrorLogPublisherResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *JsonErrorLogPublisherResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *JsonErrorLogPublisherResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JsonErrorLogPublisherResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JsonErrorLogPublisherResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


