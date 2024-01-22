# SyslogJsonSyncLogPublisherResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumsyslogJsonSyncLogPublisherSchemaUrn**](EnumsyslogJsonSyncLogPublisherSchemaUrn.md) |  | 
**Id** | **string** | Name of the Log Publisher | 
**SyslogExternalServer** | **[]string** | The syslog server to which messages should be sent. | 
**SyslogFacility** | [**EnumlogPublisherSyslogFacilityProp**](EnumlogPublisherSyslogFacilityProp.md) |  | 
**SyslogSeverity** | [**EnumlogPublisherSyslogSeverityProp**](EnumlogPublisherSyslogSeverityProp.md) |  | 
**SyslogMessageHostName** | Pointer to **string** | The local host name that will be included in syslog messages that are logged by this Syslog JSON Sync Log Publisher. | [optional] 
**SyslogMessageApplicationName** | Pointer to **string** | The application name that will be included in syslog messages that are logged by this Syslog JSON Sync Log Publisher. | [optional] 
**QueueSize** | Pointer to **int64** | The maximum number of log records that can be stored in the asynchronous queue. | [optional] 
**WriteMultiLineMessages** | Pointer to **bool** | Indicates whether the JSON objects should use a multi-line representation (with each object field and array value on its own line) that may be easier for administrators to read, but each message will be larger (because of additional spaces and end-of-line markers), and it may be more difficult to consume and parse through some text-oriented tools. | [optional] 
**IncludeProductName** | Pointer to **bool** | Indicates whether log messages should include the product name for the Directory Server. | [optional] 
**IncludeInstanceName** | Pointer to **bool** | Indicates whether log messages should include the instance name for the Directory Server. | [optional] 
**IncludeStartupID** | Pointer to **bool** | Indicates whether log messages should include the startup ID for the Directory Server, which is a value assigned to the server instance at startup and may be used to identify when the server has been restarted. | [optional] 
**IncludeThreadID** | Pointer to **bool** | Indicates whether log messages should include the thread ID for the Directory Server in each log message. This ID can be used to correlate log messages from the same thread within a single log as well as generated by the same thread across different types of log files. More information about the thread with a specific ID can be obtained using the cn&#x3D;JVM Stack Trace,cn&#x3D;monitor entry. | [optional] 
**IncludeSyncPipe** | Pointer to **[]string** | Specifies which Sync Pipes can log messages to this Sync Log Publisher. | [optional] 
**LoggedMessageType** | Pointer to [**[]EnumlogPublisherLoggedMessageTypeProp**](EnumlogPublisherLoggedMessageTypeProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Log Publisher | [optional] 
**Enabled** | **bool** | Indicates whether the Log Publisher is enabled for use. | 
**LoggingErrorBehavior** | Pointer to [**EnumlogPublisherLoggingErrorBehaviorProp**](EnumlogPublisherLoggingErrorBehaviorProp.md) |  | [optional] 

## Methods

### NewSyslogJsonSyncLogPublisherResponse

`func NewSyslogJsonSyncLogPublisherResponse(schemas []EnumsyslogJsonSyncLogPublisherSchemaUrn, id string, syslogExternalServer []string, syslogFacility EnumlogPublisherSyslogFacilityProp, syslogSeverity EnumlogPublisherSyslogSeverityProp, enabled bool, ) *SyslogJsonSyncLogPublisherResponse`

NewSyslogJsonSyncLogPublisherResponse instantiates a new SyslogJsonSyncLogPublisherResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogJsonSyncLogPublisherResponseWithDefaults

`func NewSyslogJsonSyncLogPublisherResponseWithDefaults() *SyslogJsonSyncLogPublisherResponse`

NewSyslogJsonSyncLogPublisherResponseWithDefaults instantiates a new SyslogJsonSyncLogPublisherResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *SyslogJsonSyncLogPublisherResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SyslogJsonSyncLogPublisherResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SyslogJsonSyncLogPublisherResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SyslogJsonSyncLogPublisherResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *SyslogJsonSyncLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *SyslogJsonSyncLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *SyslogJsonSyncLogPublisherResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *SyslogJsonSyncLogPublisherResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *SyslogJsonSyncLogPublisherResponse) GetSchemas() []EnumsyslogJsonSyncLogPublisherSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SyslogJsonSyncLogPublisherResponse) GetSchemasOk() (*[]EnumsyslogJsonSyncLogPublisherSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SyslogJsonSyncLogPublisherResponse) SetSchemas(v []EnumsyslogJsonSyncLogPublisherSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *SyslogJsonSyncLogPublisherResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyslogJsonSyncLogPublisherResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyslogJsonSyncLogPublisherResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSyslogExternalServer

`func (o *SyslogJsonSyncLogPublisherResponse) GetSyslogExternalServer() []string`

GetSyslogExternalServer returns the SyslogExternalServer field if non-nil, zero value otherwise.

### GetSyslogExternalServerOk

`func (o *SyslogJsonSyncLogPublisherResponse) GetSyslogExternalServerOk() (*[]string, bool)`

GetSyslogExternalServerOk returns a tuple with the SyslogExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogExternalServer

`func (o *SyslogJsonSyncLogPublisherResponse) SetSyslogExternalServer(v []string)`

SetSyslogExternalServer sets SyslogExternalServer field to given value.


### GetSyslogFacility

`func (o *SyslogJsonSyncLogPublisherResponse) GetSyslogFacility() EnumlogPublisherSyslogFacilityProp`

GetSyslogFacility returns the SyslogFacility field if non-nil, zero value otherwise.

### GetSyslogFacilityOk

`func (o *SyslogJsonSyncLogPublisherResponse) GetSyslogFacilityOk() (*EnumlogPublisherSyslogFacilityProp, bool)`

GetSyslogFacilityOk returns a tuple with the SyslogFacility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogFacility

`func (o *SyslogJsonSyncLogPublisherResponse) SetSyslogFacility(v EnumlogPublisherSyslogFacilityProp)`

SetSyslogFacility sets SyslogFacility field to given value.


### GetSyslogSeverity

`func (o *SyslogJsonSyncLogPublisherResponse) GetSyslogSeverity() EnumlogPublisherSyslogSeverityProp`

GetSyslogSeverity returns the SyslogSeverity field if non-nil, zero value otherwise.

### GetSyslogSeverityOk

`func (o *SyslogJsonSyncLogPublisherResponse) GetSyslogSeverityOk() (*EnumlogPublisherSyslogSeverityProp, bool)`

GetSyslogSeverityOk returns a tuple with the SyslogSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogSeverity

`func (o *SyslogJsonSyncLogPublisherResponse) SetSyslogSeverity(v EnumlogPublisherSyslogSeverityProp)`

SetSyslogSeverity sets SyslogSeverity field to given value.


### GetSyslogMessageHostName

`func (o *SyslogJsonSyncLogPublisherResponse) GetSyslogMessageHostName() string`

GetSyslogMessageHostName returns the SyslogMessageHostName field if non-nil, zero value otherwise.

### GetSyslogMessageHostNameOk

`func (o *SyslogJsonSyncLogPublisherResponse) GetSyslogMessageHostNameOk() (*string, bool)`

GetSyslogMessageHostNameOk returns a tuple with the SyslogMessageHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogMessageHostName

`func (o *SyslogJsonSyncLogPublisherResponse) SetSyslogMessageHostName(v string)`

SetSyslogMessageHostName sets SyslogMessageHostName field to given value.

### HasSyslogMessageHostName

`func (o *SyslogJsonSyncLogPublisherResponse) HasSyslogMessageHostName() bool`

HasSyslogMessageHostName returns a boolean if a field has been set.

### GetSyslogMessageApplicationName

`func (o *SyslogJsonSyncLogPublisherResponse) GetSyslogMessageApplicationName() string`

GetSyslogMessageApplicationName returns the SyslogMessageApplicationName field if non-nil, zero value otherwise.

### GetSyslogMessageApplicationNameOk

`func (o *SyslogJsonSyncLogPublisherResponse) GetSyslogMessageApplicationNameOk() (*string, bool)`

GetSyslogMessageApplicationNameOk returns a tuple with the SyslogMessageApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogMessageApplicationName

`func (o *SyslogJsonSyncLogPublisherResponse) SetSyslogMessageApplicationName(v string)`

SetSyslogMessageApplicationName sets SyslogMessageApplicationName field to given value.

### HasSyslogMessageApplicationName

`func (o *SyslogJsonSyncLogPublisherResponse) HasSyslogMessageApplicationName() bool`

HasSyslogMessageApplicationName returns a boolean if a field has been set.

### GetQueueSize

`func (o *SyslogJsonSyncLogPublisherResponse) GetQueueSize() int64`

GetQueueSize returns the QueueSize field if non-nil, zero value otherwise.

### GetQueueSizeOk

`func (o *SyslogJsonSyncLogPublisherResponse) GetQueueSizeOk() (*int64, bool)`

GetQueueSizeOk returns a tuple with the QueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueSize

`func (o *SyslogJsonSyncLogPublisherResponse) SetQueueSize(v int64)`

SetQueueSize sets QueueSize field to given value.

### HasQueueSize

`func (o *SyslogJsonSyncLogPublisherResponse) HasQueueSize() bool`

HasQueueSize returns a boolean if a field has been set.

### GetWriteMultiLineMessages

`func (o *SyslogJsonSyncLogPublisherResponse) GetWriteMultiLineMessages() bool`

GetWriteMultiLineMessages returns the WriteMultiLineMessages field if non-nil, zero value otherwise.

### GetWriteMultiLineMessagesOk

`func (o *SyslogJsonSyncLogPublisherResponse) GetWriteMultiLineMessagesOk() (*bool, bool)`

GetWriteMultiLineMessagesOk returns a tuple with the WriteMultiLineMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteMultiLineMessages

`func (o *SyslogJsonSyncLogPublisherResponse) SetWriteMultiLineMessages(v bool)`

SetWriteMultiLineMessages sets WriteMultiLineMessages field to given value.

### HasWriteMultiLineMessages

`func (o *SyslogJsonSyncLogPublisherResponse) HasWriteMultiLineMessages() bool`

HasWriteMultiLineMessages returns a boolean if a field has been set.

### GetIncludeProductName

`func (o *SyslogJsonSyncLogPublisherResponse) GetIncludeProductName() bool`

GetIncludeProductName returns the IncludeProductName field if non-nil, zero value otherwise.

### GetIncludeProductNameOk

`func (o *SyslogJsonSyncLogPublisherResponse) GetIncludeProductNameOk() (*bool, bool)`

GetIncludeProductNameOk returns a tuple with the IncludeProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeProductName

`func (o *SyslogJsonSyncLogPublisherResponse) SetIncludeProductName(v bool)`

SetIncludeProductName sets IncludeProductName field to given value.

### HasIncludeProductName

`func (o *SyslogJsonSyncLogPublisherResponse) HasIncludeProductName() bool`

HasIncludeProductName returns a boolean if a field has been set.

### GetIncludeInstanceName

`func (o *SyslogJsonSyncLogPublisherResponse) GetIncludeInstanceName() bool`

GetIncludeInstanceName returns the IncludeInstanceName field if non-nil, zero value otherwise.

### GetIncludeInstanceNameOk

`func (o *SyslogJsonSyncLogPublisherResponse) GetIncludeInstanceNameOk() (*bool, bool)`

GetIncludeInstanceNameOk returns a tuple with the IncludeInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInstanceName

`func (o *SyslogJsonSyncLogPublisherResponse) SetIncludeInstanceName(v bool)`

SetIncludeInstanceName sets IncludeInstanceName field to given value.

### HasIncludeInstanceName

`func (o *SyslogJsonSyncLogPublisherResponse) HasIncludeInstanceName() bool`

HasIncludeInstanceName returns a boolean if a field has been set.

### GetIncludeStartupID

`func (o *SyslogJsonSyncLogPublisherResponse) GetIncludeStartupID() bool`

GetIncludeStartupID returns the IncludeStartupID field if non-nil, zero value otherwise.

### GetIncludeStartupIDOk

`func (o *SyslogJsonSyncLogPublisherResponse) GetIncludeStartupIDOk() (*bool, bool)`

GetIncludeStartupIDOk returns a tuple with the IncludeStartupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeStartupID

`func (o *SyslogJsonSyncLogPublisherResponse) SetIncludeStartupID(v bool)`

SetIncludeStartupID sets IncludeStartupID field to given value.

### HasIncludeStartupID

`func (o *SyslogJsonSyncLogPublisherResponse) HasIncludeStartupID() bool`

HasIncludeStartupID returns a boolean if a field has been set.

### GetIncludeThreadID

`func (o *SyslogJsonSyncLogPublisherResponse) GetIncludeThreadID() bool`

GetIncludeThreadID returns the IncludeThreadID field if non-nil, zero value otherwise.

### GetIncludeThreadIDOk

`func (o *SyslogJsonSyncLogPublisherResponse) GetIncludeThreadIDOk() (*bool, bool)`

GetIncludeThreadIDOk returns a tuple with the IncludeThreadID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeThreadID

`func (o *SyslogJsonSyncLogPublisherResponse) SetIncludeThreadID(v bool)`

SetIncludeThreadID sets IncludeThreadID field to given value.

### HasIncludeThreadID

`func (o *SyslogJsonSyncLogPublisherResponse) HasIncludeThreadID() bool`

HasIncludeThreadID returns a boolean if a field has been set.

### GetIncludeSyncPipe

`func (o *SyslogJsonSyncLogPublisherResponse) GetIncludeSyncPipe() []string`

GetIncludeSyncPipe returns the IncludeSyncPipe field if non-nil, zero value otherwise.

### GetIncludeSyncPipeOk

`func (o *SyslogJsonSyncLogPublisherResponse) GetIncludeSyncPipeOk() (*[]string, bool)`

GetIncludeSyncPipeOk returns a tuple with the IncludeSyncPipe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSyncPipe

`func (o *SyslogJsonSyncLogPublisherResponse) SetIncludeSyncPipe(v []string)`

SetIncludeSyncPipe sets IncludeSyncPipe field to given value.

### HasIncludeSyncPipe

`func (o *SyslogJsonSyncLogPublisherResponse) HasIncludeSyncPipe() bool`

HasIncludeSyncPipe returns a boolean if a field has been set.

### GetLoggedMessageType

`func (o *SyslogJsonSyncLogPublisherResponse) GetLoggedMessageType() []EnumlogPublisherLoggedMessageTypeProp`

GetLoggedMessageType returns the LoggedMessageType field if non-nil, zero value otherwise.

### GetLoggedMessageTypeOk

`func (o *SyslogJsonSyncLogPublisherResponse) GetLoggedMessageTypeOk() (*[]EnumlogPublisherLoggedMessageTypeProp, bool)`

GetLoggedMessageTypeOk returns a tuple with the LoggedMessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedMessageType

`func (o *SyslogJsonSyncLogPublisherResponse) SetLoggedMessageType(v []EnumlogPublisherLoggedMessageTypeProp)`

SetLoggedMessageType sets LoggedMessageType field to given value.

### HasLoggedMessageType

`func (o *SyslogJsonSyncLogPublisherResponse) HasLoggedMessageType() bool`

HasLoggedMessageType returns a boolean if a field has been set.

### GetDescription

`func (o *SyslogJsonSyncLogPublisherResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SyslogJsonSyncLogPublisherResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SyslogJsonSyncLogPublisherResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SyslogJsonSyncLogPublisherResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SyslogJsonSyncLogPublisherResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SyslogJsonSyncLogPublisherResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SyslogJsonSyncLogPublisherResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLoggingErrorBehavior

`func (o *SyslogJsonSyncLogPublisherResponse) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp`

GetLoggingErrorBehavior returns the LoggingErrorBehavior field if non-nil, zero value otherwise.

### GetLoggingErrorBehaviorOk

`func (o *SyslogJsonSyncLogPublisherResponse) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool)`

GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingErrorBehavior

`func (o *SyslogJsonSyncLogPublisherResponse) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp)`

SetLoggingErrorBehavior sets LoggingErrorBehavior field to given value.

### HasLoggingErrorBehavior

`func (o *SyslogJsonSyncLogPublisherResponse) HasLoggingErrorBehavior() bool`

HasLoggingErrorBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

