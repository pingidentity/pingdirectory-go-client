# SyslogBasedErrorLogPublisherShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsyslogBasedErrorLogPublisherSchemaUrn**](EnumsyslogBasedErrorLogPublisherSchemaUrn.md) |  | 
**Enabled** | **bool** | Indicates whether the Syslog Based Error Log Publisher is enabled for use. | 
**ServerHostName** | **string** | Specifies the hostname or IP address of the syslogd host to log to. It is highly recommend to use localhost. | 
**ServerPort** | **int32** | Specifies the port number of the syslogd host to log to. | 
**SyslogFacility** | **int32** | Specifies the syslog facility to use for this Syslog Based Error Log Publisher | 
**AutoFlush** | Pointer to **bool** | Specifies whether to flush the writer after every log record. | [optional] 
**Asynchronous** | **bool** | Indicates whether the Syslog Based Error Log Publisher will publish records asynchronously. | 
**QueueSize** | Pointer to **int32** | The maximum number of log records that can be stored in the asynchronous queue. | [optional] 
**DefaultSeverity** | Pointer to [**[]EnumlogPublisherDefaultSeverityProp**](EnumlogPublisherDefaultSeverityProp.md) |  | [optional] 
**OverrideSeverity** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Log Publisher | [optional] 
**LoggingErrorBehavior** | Pointer to [**EnumlogPublisherLoggingErrorBehaviorProp**](EnumlogPublisherLoggingErrorBehaviorProp.md) |  | [optional] 

## Methods

### NewSyslogBasedErrorLogPublisherShared

`func NewSyslogBasedErrorLogPublisherShared(schemas []EnumsyslogBasedErrorLogPublisherSchemaUrn, enabled bool, serverHostName string, serverPort int32, syslogFacility int32, asynchronous bool, ) *SyslogBasedErrorLogPublisherShared`

NewSyslogBasedErrorLogPublisherShared instantiates a new SyslogBasedErrorLogPublisherShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogBasedErrorLogPublisherSharedWithDefaults

`func NewSyslogBasedErrorLogPublisherSharedWithDefaults() *SyslogBasedErrorLogPublisherShared`

NewSyslogBasedErrorLogPublisherSharedWithDefaults instantiates a new SyslogBasedErrorLogPublisherShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *SyslogBasedErrorLogPublisherShared) GetSchemas() []EnumsyslogBasedErrorLogPublisherSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SyslogBasedErrorLogPublisherShared) GetSchemasOk() (*[]EnumsyslogBasedErrorLogPublisherSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SyslogBasedErrorLogPublisherShared) SetSchemas(v []EnumsyslogBasedErrorLogPublisherSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEnabled

`func (o *SyslogBasedErrorLogPublisherShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SyslogBasedErrorLogPublisherShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SyslogBasedErrorLogPublisherShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetServerHostName

`func (o *SyslogBasedErrorLogPublisherShared) GetServerHostName() string`

GetServerHostName returns the ServerHostName field if non-nil, zero value otherwise.

### GetServerHostNameOk

`func (o *SyslogBasedErrorLogPublisherShared) GetServerHostNameOk() (*string, bool)`

GetServerHostNameOk returns a tuple with the ServerHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostName

`func (o *SyslogBasedErrorLogPublisherShared) SetServerHostName(v string)`

SetServerHostName sets ServerHostName field to given value.


### GetServerPort

`func (o *SyslogBasedErrorLogPublisherShared) GetServerPort() int32`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *SyslogBasedErrorLogPublisherShared) GetServerPortOk() (*int32, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *SyslogBasedErrorLogPublisherShared) SetServerPort(v int32)`

SetServerPort sets ServerPort field to given value.


### GetSyslogFacility

`func (o *SyslogBasedErrorLogPublisherShared) GetSyslogFacility() int32`

GetSyslogFacility returns the SyslogFacility field if non-nil, zero value otherwise.

### GetSyslogFacilityOk

`func (o *SyslogBasedErrorLogPublisherShared) GetSyslogFacilityOk() (*int32, bool)`

GetSyslogFacilityOk returns a tuple with the SyslogFacility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogFacility

`func (o *SyslogBasedErrorLogPublisherShared) SetSyslogFacility(v int32)`

SetSyslogFacility sets SyslogFacility field to given value.


### GetAutoFlush

`func (o *SyslogBasedErrorLogPublisherShared) GetAutoFlush() bool`

GetAutoFlush returns the AutoFlush field if non-nil, zero value otherwise.

### GetAutoFlushOk

`func (o *SyslogBasedErrorLogPublisherShared) GetAutoFlushOk() (*bool, bool)`

GetAutoFlushOk returns a tuple with the AutoFlush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoFlush

`func (o *SyslogBasedErrorLogPublisherShared) SetAutoFlush(v bool)`

SetAutoFlush sets AutoFlush field to given value.

### HasAutoFlush

`func (o *SyslogBasedErrorLogPublisherShared) HasAutoFlush() bool`

HasAutoFlush returns a boolean if a field has been set.

### GetAsynchronous

`func (o *SyslogBasedErrorLogPublisherShared) GetAsynchronous() bool`

GetAsynchronous returns the Asynchronous field if non-nil, zero value otherwise.

### GetAsynchronousOk

`func (o *SyslogBasedErrorLogPublisherShared) GetAsynchronousOk() (*bool, bool)`

GetAsynchronousOk returns a tuple with the Asynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsynchronous

`func (o *SyslogBasedErrorLogPublisherShared) SetAsynchronous(v bool)`

SetAsynchronous sets Asynchronous field to given value.


### GetQueueSize

`func (o *SyslogBasedErrorLogPublisherShared) GetQueueSize() int32`

GetQueueSize returns the QueueSize field if non-nil, zero value otherwise.

### GetQueueSizeOk

`func (o *SyslogBasedErrorLogPublisherShared) GetQueueSizeOk() (*int32, bool)`

GetQueueSizeOk returns a tuple with the QueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueSize

`func (o *SyslogBasedErrorLogPublisherShared) SetQueueSize(v int32)`

SetQueueSize sets QueueSize field to given value.

### HasQueueSize

`func (o *SyslogBasedErrorLogPublisherShared) HasQueueSize() bool`

HasQueueSize returns a boolean if a field has been set.

### GetDefaultSeverity

`func (o *SyslogBasedErrorLogPublisherShared) GetDefaultSeverity() []EnumlogPublisherDefaultSeverityProp`

GetDefaultSeverity returns the DefaultSeverity field if non-nil, zero value otherwise.

### GetDefaultSeverityOk

`func (o *SyslogBasedErrorLogPublisherShared) GetDefaultSeverityOk() (*[]EnumlogPublisherDefaultSeverityProp, bool)`

GetDefaultSeverityOk returns a tuple with the DefaultSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSeverity

`func (o *SyslogBasedErrorLogPublisherShared) SetDefaultSeverity(v []EnumlogPublisherDefaultSeverityProp)`

SetDefaultSeverity sets DefaultSeverity field to given value.

### HasDefaultSeverity

`func (o *SyslogBasedErrorLogPublisherShared) HasDefaultSeverity() bool`

HasDefaultSeverity returns a boolean if a field has been set.

### GetOverrideSeverity

`func (o *SyslogBasedErrorLogPublisherShared) GetOverrideSeverity() []string`

GetOverrideSeverity returns the OverrideSeverity field if non-nil, zero value otherwise.

### GetOverrideSeverityOk

`func (o *SyslogBasedErrorLogPublisherShared) GetOverrideSeverityOk() (*[]string, bool)`

GetOverrideSeverityOk returns a tuple with the OverrideSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSeverity

`func (o *SyslogBasedErrorLogPublisherShared) SetOverrideSeverity(v []string)`

SetOverrideSeverity sets OverrideSeverity field to given value.

### HasOverrideSeverity

`func (o *SyslogBasedErrorLogPublisherShared) HasOverrideSeverity() bool`

HasOverrideSeverity returns a boolean if a field has been set.

### GetDescription

`func (o *SyslogBasedErrorLogPublisherShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SyslogBasedErrorLogPublisherShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SyslogBasedErrorLogPublisherShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SyslogBasedErrorLogPublisherShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLoggingErrorBehavior

`func (o *SyslogBasedErrorLogPublisherShared) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp`

GetLoggingErrorBehavior returns the LoggingErrorBehavior field if non-nil, zero value otherwise.

### GetLoggingErrorBehaviorOk

`func (o *SyslogBasedErrorLogPublisherShared) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool)`

GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingErrorBehavior

`func (o *SyslogBasedErrorLogPublisherShared) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp)`

SetLoggingErrorBehavior sets LoggingErrorBehavior field to given value.

### HasLoggingErrorBehavior

`func (o *SyslogBasedErrorLogPublisherShared) HasLoggingErrorBehavior() bool`

HasLoggingErrorBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


