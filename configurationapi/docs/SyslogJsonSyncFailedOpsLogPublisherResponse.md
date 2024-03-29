# SyslogJsonSyncFailedOpsLogPublisherResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumsyslogJsonSyncFailedOpsLogPublisherSchemaUrn**](EnumsyslogJsonSyncFailedOpsLogPublisherSchemaUrn.md) |  | 
**Id** | **string** | Name of the Log Publisher | 
**SyslogExternalServer** | **[]string** | The syslog server to which messages should be sent. | 
**SyslogFacility** | [**EnumlogPublisherSyslogFacilityProp**](EnumlogPublisherSyslogFacilityProp.md) |  | 
**SyslogSeverity** | [**EnumlogPublisherSyslogSeverityProp**](EnumlogPublisherSyslogSeverityProp.md) |  | 
**SyslogMessageHostName** | Pointer to **string** | The local host name that will be included in syslog messages that are logged by this Syslog JSON Sync Failed Ops Log Publisher. | [optional] 
**SyslogMessageApplicationName** | Pointer to **string** | The application name that will be included in syslog messages that are logged by this Syslog JSON Sync Failed Ops Log Publisher. | [optional] 
**QueueSize** | Pointer to **int64** | The maximum number of log records that can be stored in the asynchronous queue. | [optional] 
**WriteMultiLineMessages** | Pointer to **bool** | Indicates whether the JSON objects should use a multi-line representation (with each object field and array value on its own line) that may be easier for administrators to read, but each message will be larger (because of additional spaces and end-of-line markers), and it may be more difficult to consume and parse through some text-oriented tools. | [optional] 
**IncludeProductName** | Pointer to **bool** | Indicates whether log messages should include the product name for the Directory Server. | [optional] 
**IncludeInstanceName** | Pointer to **bool** | Indicates whether log messages should include the instance name for the Directory Server. | [optional] 
**IncludeStartupID** | Pointer to **bool** | Indicates whether log messages should include the startup ID for the Directory Server, which is a value assigned to the server instance at startup and may be used to identify when the server has been restarted. | [optional] 
**IncludeThreadID** | Pointer to **bool** | Indicates whether log messages should include the thread ID for the Directory Server in each log message. This ID can be used to correlate log messages from the same thread within a single log as well as generated by the same thread across different types of log files. More information about the thread with a specific ID can be obtained using the cn&#x3D;JVM Stack Trace,cn&#x3D;monitor entry. | [optional] 
**IncludeSyncPipe** | Pointer to **[]string** | Specifies which Sync Pipes can log messages to this Sync Log Publisher. | [optional] 
**Description** | Pointer to **string** | A description for this Log Publisher | [optional] 
**Enabled** | **bool** | Indicates whether the Log Publisher is enabled for use. | 
**LoggingErrorBehavior** | Pointer to [**EnumlogPublisherLoggingErrorBehaviorProp**](EnumlogPublisherLoggingErrorBehaviorProp.md) |  | [optional] 

## Methods

### NewSyslogJsonSyncFailedOpsLogPublisherResponse

`func NewSyslogJsonSyncFailedOpsLogPublisherResponse(schemas []EnumsyslogJsonSyncFailedOpsLogPublisherSchemaUrn, id string, syslogExternalServer []string, syslogFacility EnumlogPublisherSyslogFacilityProp, syslogSeverity EnumlogPublisherSyslogSeverityProp, enabled bool, ) *SyslogJsonSyncFailedOpsLogPublisherResponse`

NewSyslogJsonSyncFailedOpsLogPublisherResponse instantiates a new SyslogJsonSyncFailedOpsLogPublisherResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogJsonSyncFailedOpsLogPublisherResponseWithDefaults

`func NewSyslogJsonSyncFailedOpsLogPublisherResponseWithDefaults() *SyslogJsonSyncFailedOpsLogPublisherResponse`

NewSyslogJsonSyncFailedOpsLogPublisherResponseWithDefaults instantiates a new SyslogJsonSyncFailedOpsLogPublisherResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetSchemas() []EnumsyslogJsonSyncFailedOpsLogPublisherSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetSchemasOk() (*[]EnumsyslogJsonSyncFailedOpsLogPublisherSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) SetSchemas(v []EnumsyslogJsonSyncFailedOpsLogPublisherSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSyslogExternalServer

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetSyslogExternalServer() []string`

GetSyslogExternalServer returns the SyslogExternalServer field if non-nil, zero value otherwise.

### GetSyslogExternalServerOk

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetSyslogExternalServerOk() (*[]string, bool)`

GetSyslogExternalServerOk returns a tuple with the SyslogExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogExternalServer

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) SetSyslogExternalServer(v []string)`

SetSyslogExternalServer sets SyslogExternalServer field to given value.


### GetSyslogFacility

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetSyslogFacility() EnumlogPublisherSyslogFacilityProp`

GetSyslogFacility returns the SyslogFacility field if non-nil, zero value otherwise.

### GetSyslogFacilityOk

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetSyslogFacilityOk() (*EnumlogPublisherSyslogFacilityProp, bool)`

GetSyslogFacilityOk returns a tuple with the SyslogFacility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogFacility

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) SetSyslogFacility(v EnumlogPublisherSyslogFacilityProp)`

SetSyslogFacility sets SyslogFacility field to given value.


### GetSyslogSeverity

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetSyslogSeverity() EnumlogPublisherSyslogSeverityProp`

GetSyslogSeverity returns the SyslogSeverity field if non-nil, zero value otherwise.

### GetSyslogSeverityOk

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetSyslogSeverityOk() (*EnumlogPublisherSyslogSeverityProp, bool)`

GetSyslogSeverityOk returns a tuple with the SyslogSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogSeverity

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) SetSyslogSeverity(v EnumlogPublisherSyslogSeverityProp)`

SetSyslogSeverity sets SyslogSeverity field to given value.


### GetSyslogMessageHostName

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetSyslogMessageHostName() string`

GetSyslogMessageHostName returns the SyslogMessageHostName field if non-nil, zero value otherwise.

### GetSyslogMessageHostNameOk

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetSyslogMessageHostNameOk() (*string, bool)`

GetSyslogMessageHostNameOk returns a tuple with the SyslogMessageHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogMessageHostName

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) SetSyslogMessageHostName(v string)`

SetSyslogMessageHostName sets SyslogMessageHostName field to given value.

### HasSyslogMessageHostName

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) HasSyslogMessageHostName() bool`

HasSyslogMessageHostName returns a boolean if a field has been set.

### GetSyslogMessageApplicationName

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetSyslogMessageApplicationName() string`

GetSyslogMessageApplicationName returns the SyslogMessageApplicationName field if non-nil, zero value otherwise.

### GetSyslogMessageApplicationNameOk

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetSyslogMessageApplicationNameOk() (*string, bool)`

GetSyslogMessageApplicationNameOk returns a tuple with the SyslogMessageApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogMessageApplicationName

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) SetSyslogMessageApplicationName(v string)`

SetSyslogMessageApplicationName sets SyslogMessageApplicationName field to given value.

### HasSyslogMessageApplicationName

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) HasSyslogMessageApplicationName() bool`

HasSyslogMessageApplicationName returns a boolean if a field has been set.

### GetQueueSize

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetQueueSize() int64`

GetQueueSize returns the QueueSize field if non-nil, zero value otherwise.

### GetQueueSizeOk

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetQueueSizeOk() (*int64, bool)`

GetQueueSizeOk returns a tuple with the QueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueSize

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) SetQueueSize(v int64)`

SetQueueSize sets QueueSize field to given value.

### HasQueueSize

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) HasQueueSize() bool`

HasQueueSize returns a boolean if a field has been set.

### GetWriteMultiLineMessages

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetWriteMultiLineMessages() bool`

GetWriteMultiLineMessages returns the WriteMultiLineMessages field if non-nil, zero value otherwise.

### GetWriteMultiLineMessagesOk

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetWriteMultiLineMessagesOk() (*bool, bool)`

GetWriteMultiLineMessagesOk returns a tuple with the WriteMultiLineMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteMultiLineMessages

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) SetWriteMultiLineMessages(v bool)`

SetWriteMultiLineMessages sets WriteMultiLineMessages field to given value.

### HasWriteMultiLineMessages

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) HasWriteMultiLineMessages() bool`

HasWriteMultiLineMessages returns a boolean if a field has been set.

### GetIncludeProductName

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetIncludeProductName() bool`

GetIncludeProductName returns the IncludeProductName field if non-nil, zero value otherwise.

### GetIncludeProductNameOk

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetIncludeProductNameOk() (*bool, bool)`

GetIncludeProductNameOk returns a tuple with the IncludeProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeProductName

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) SetIncludeProductName(v bool)`

SetIncludeProductName sets IncludeProductName field to given value.

### HasIncludeProductName

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) HasIncludeProductName() bool`

HasIncludeProductName returns a boolean if a field has been set.

### GetIncludeInstanceName

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetIncludeInstanceName() bool`

GetIncludeInstanceName returns the IncludeInstanceName field if non-nil, zero value otherwise.

### GetIncludeInstanceNameOk

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetIncludeInstanceNameOk() (*bool, bool)`

GetIncludeInstanceNameOk returns a tuple with the IncludeInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInstanceName

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) SetIncludeInstanceName(v bool)`

SetIncludeInstanceName sets IncludeInstanceName field to given value.

### HasIncludeInstanceName

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) HasIncludeInstanceName() bool`

HasIncludeInstanceName returns a boolean if a field has been set.

### GetIncludeStartupID

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetIncludeStartupID() bool`

GetIncludeStartupID returns the IncludeStartupID field if non-nil, zero value otherwise.

### GetIncludeStartupIDOk

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetIncludeStartupIDOk() (*bool, bool)`

GetIncludeStartupIDOk returns a tuple with the IncludeStartupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeStartupID

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) SetIncludeStartupID(v bool)`

SetIncludeStartupID sets IncludeStartupID field to given value.

### HasIncludeStartupID

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) HasIncludeStartupID() bool`

HasIncludeStartupID returns a boolean if a field has been set.

### GetIncludeThreadID

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetIncludeThreadID() bool`

GetIncludeThreadID returns the IncludeThreadID field if non-nil, zero value otherwise.

### GetIncludeThreadIDOk

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetIncludeThreadIDOk() (*bool, bool)`

GetIncludeThreadIDOk returns a tuple with the IncludeThreadID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeThreadID

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) SetIncludeThreadID(v bool)`

SetIncludeThreadID sets IncludeThreadID field to given value.

### HasIncludeThreadID

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) HasIncludeThreadID() bool`

HasIncludeThreadID returns a boolean if a field has been set.

### GetIncludeSyncPipe

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetIncludeSyncPipe() []string`

GetIncludeSyncPipe returns the IncludeSyncPipe field if non-nil, zero value otherwise.

### GetIncludeSyncPipeOk

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetIncludeSyncPipeOk() (*[]string, bool)`

GetIncludeSyncPipeOk returns a tuple with the IncludeSyncPipe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSyncPipe

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) SetIncludeSyncPipe(v []string)`

SetIncludeSyncPipe sets IncludeSyncPipe field to given value.

### HasIncludeSyncPipe

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) HasIncludeSyncPipe() bool`

HasIncludeSyncPipe returns a boolean if a field has been set.

### GetDescription

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLoggingErrorBehavior

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp`

GetLoggingErrorBehavior returns the LoggingErrorBehavior field if non-nil, zero value otherwise.

### GetLoggingErrorBehaviorOk

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool)`

GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingErrorBehavior

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp)`

SetLoggingErrorBehavior sets LoggingErrorBehavior field to given value.

### HasLoggingErrorBehavior

`func (o *SyslogJsonSyncFailedOpsLogPublisherResponse) HasLoggingErrorBehavior() bool`

HasLoggingErrorBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


