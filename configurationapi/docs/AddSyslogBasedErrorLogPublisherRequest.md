# AddSyslogBasedErrorLogPublisherRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsyslogBasedErrorLogPublisherSchemaUrn**](EnumsyslogBasedErrorLogPublisherSchemaUrn.md) |  | 
**Enabled** | **bool** | Indicates whether the Syslog Based Error Log Publisher is enabled for use. | 
**ServerHostName** | Pointer to **string** | Specifies the hostname or IP address of the syslogd host to log to. It is highly recommend to use localhost. | [optional] 
**ServerPort** | Pointer to **int64** | Specifies the port number of the syslogd host to log to. | [optional] 
**SyslogFacility** | Pointer to **int64** | Specifies the syslog facility to use for this Syslog Based Error Log Publisher | [optional] 
**AutoFlush** | Pointer to **bool** | Specifies whether to flush the writer after every log record. | [optional] 
**Asynchronous** | Pointer to **bool** | Indicates whether the Syslog Based Error Log Publisher will publish records asynchronously. | [optional] 
**QueueSize** | Pointer to **int64** | The maximum number of log records that can be stored in the asynchronous queue. | [optional] 
**DefaultSeverity** | Pointer to [**[]EnumlogPublisherDefaultSeverityProp**](EnumlogPublisherDefaultSeverityProp.md) |  | [optional] 
**OverrideSeverity** | Pointer to **[]string** | Specifies the override severity levels for the logger based on the category of the messages. | [optional] 
**Description** | Pointer to **string** | A description for this Log Publisher | [optional] 
**LoggingErrorBehavior** | Pointer to [**EnumlogPublisherLoggingErrorBehaviorProp**](EnumlogPublisherLoggingErrorBehaviorProp.md) |  | [optional] 
**PublisherName** | **string** | Name of the new Log Publisher | 

## Methods

### NewAddSyslogBasedErrorLogPublisherRequest

`func NewAddSyslogBasedErrorLogPublisherRequest(schemas []EnumsyslogBasedErrorLogPublisherSchemaUrn, enabled bool, publisherName string, ) *AddSyslogBasedErrorLogPublisherRequest`

NewAddSyslogBasedErrorLogPublisherRequest instantiates a new AddSyslogBasedErrorLogPublisherRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSyslogBasedErrorLogPublisherRequestWithDefaults

`func NewAddSyslogBasedErrorLogPublisherRequestWithDefaults() *AddSyslogBasedErrorLogPublisherRequest`

NewAddSyslogBasedErrorLogPublisherRequestWithDefaults instantiates a new AddSyslogBasedErrorLogPublisherRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddSyslogBasedErrorLogPublisherRequest) GetSchemas() []EnumsyslogBasedErrorLogPublisherSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddSyslogBasedErrorLogPublisherRequest) GetSchemasOk() (*[]EnumsyslogBasedErrorLogPublisherSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddSyslogBasedErrorLogPublisherRequest) SetSchemas(v []EnumsyslogBasedErrorLogPublisherSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEnabled

`func (o *AddSyslogBasedErrorLogPublisherRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddSyslogBasedErrorLogPublisherRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddSyslogBasedErrorLogPublisherRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetServerHostName

`func (o *AddSyslogBasedErrorLogPublisherRequest) GetServerHostName() string`

GetServerHostName returns the ServerHostName field if non-nil, zero value otherwise.

### GetServerHostNameOk

`func (o *AddSyslogBasedErrorLogPublisherRequest) GetServerHostNameOk() (*string, bool)`

GetServerHostNameOk returns a tuple with the ServerHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostName

`func (o *AddSyslogBasedErrorLogPublisherRequest) SetServerHostName(v string)`

SetServerHostName sets ServerHostName field to given value.

### HasServerHostName

`func (o *AddSyslogBasedErrorLogPublisherRequest) HasServerHostName() bool`

HasServerHostName returns a boolean if a field has been set.

### GetServerPort

`func (o *AddSyslogBasedErrorLogPublisherRequest) GetServerPort() int64`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *AddSyslogBasedErrorLogPublisherRequest) GetServerPortOk() (*int64, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *AddSyslogBasedErrorLogPublisherRequest) SetServerPort(v int64)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *AddSyslogBasedErrorLogPublisherRequest) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### GetSyslogFacility

`func (o *AddSyslogBasedErrorLogPublisherRequest) GetSyslogFacility() int64`

GetSyslogFacility returns the SyslogFacility field if non-nil, zero value otherwise.

### GetSyslogFacilityOk

`func (o *AddSyslogBasedErrorLogPublisherRequest) GetSyslogFacilityOk() (*int64, bool)`

GetSyslogFacilityOk returns a tuple with the SyslogFacility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogFacility

`func (o *AddSyslogBasedErrorLogPublisherRequest) SetSyslogFacility(v int64)`

SetSyslogFacility sets SyslogFacility field to given value.

### HasSyslogFacility

`func (o *AddSyslogBasedErrorLogPublisherRequest) HasSyslogFacility() bool`

HasSyslogFacility returns a boolean if a field has been set.

### GetAutoFlush

`func (o *AddSyslogBasedErrorLogPublisherRequest) GetAutoFlush() bool`

GetAutoFlush returns the AutoFlush field if non-nil, zero value otherwise.

### GetAutoFlushOk

`func (o *AddSyslogBasedErrorLogPublisherRequest) GetAutoFlushOk() (*bool, bool)`

GetAutoFlushOk returns a tuple with the AutoFlush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoFlush

`func (o *AddSyslogBasedErrorLogPublisherRequest) SetAutoFlush(v bool)`

SetAutoFlush sets AutoFlush field to given value.

### HasAutoFlush

`func (o *AddSyslogBasedErrorLogPublisherRequest) HasAutoFlush() bool`

HasAutoFlush returns a boolean if a field has been set.

### GetAsynchronous

`func (o *AddSyslogBasedErrorLogPublisherRequest) GetAsynchronous() bool`

GetAsynchronous returns the Asynchronous field if non-nil, zero value otherwise.

### GetAsynchronousOk

`func (o *AddSyslogBasedErrorLogPublisherRequest) GetAsynchronousOk() (*bool, bool)`

GetAsynchronousOk returns a tuple with the Asynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsynchronous

`func (o *AddSyslogBasedErrorLogPublisherRequest) SetAsynchronous(v bool)`

SetAsynchronous sets Asynchronous field to given value.

### HasAsynchronous

`func (o *AddSyslogBasedErrorLogPublisherRequest) HasAsynchronous() bool`

HasAsynchronous returns a boolean if a field has been set.

### GetQueueSize

`func (o *AddSyslogBasedErrorLogPublisherRequest) GetQueueSize() int64`

GetQueueSize returns the QueueSize field if non-nil, zero value otherwise.

### GetQueueSizeOk

`func (o *AddSyslogBasedErrorLogPublisherRequest) GetQueueSizeOk() (*int64, bool)`

GetQueueSizeOk returns a tuple with the QueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueSize

`func (o *AddSyslogBasedErrorLogPublisherRequest) SetQueueSize(v int64)`

SetQueueSize sets QueueSize field to given value.

### HasQueueSize

`func (o *AddSyslogBasedErrorLogPublisherRequest) HasQueueSize() bool`

HasQueueSize returns a boolean if a field has been set.

### GetDefaultSeverity

`func (o *AddSyslogBasedErrorLogPublisherRequest) GetDefaultSeverity() []EnumlogPublisherDefaultSeverityProp`

GetDefaultSeverity returns the DefaultSeverity field if non-nil, zero value otherwise.

### GetDefaultSeverityOk

`func (o *AddSyslogBasedErrorLogPublisherRequest) GetDefaultSeverityOk() (*[]EnumlogPublisherDefaultSeverityProp, bool)`

GetDefaultSeverityOk returns a tuple with the DefaultSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSeverity

`func (o *AddSyslogBasedErrorLogPublisherRequest) SetDefaultSeverity(v []EnumlogPublisherDefaultSeverityProp)`

SetDefaultSeverity sets DefaultSeverity field to given value.

### HasDefaultSeverity

`func (o *AddSyslogBasedErrorLogPublisherRequest) HasDefaultSeverity() bool`

HasDefaultSeverity returns a boolean if a field has been set.

### GetOverrideSeverity

`func (o *AddSyslogBasedErrorLogPublisherRequest) GetOverrideSeverity() []string`

GetOverrideSeverity returns the OverrideSeverity field if non-nil, zero value otherwise.

### GetOverrideSeverityOk

`func (o *AddSyslogBasedErrorLogPublisherRequest) GetOverrideSeverityOk() (*[]string, bool)`

GetOverrideSeverityOk returns a tuple with the OverrideSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSeverity

`func (o *AddSyslogBasedErrorLogPublisherRequest) SetOverrideSeverity(v []string)`

SetOverrideSeverity sets OverrideSeverity field to given value.

### HasOverrideSeverity

`func (o *AddSyslogBasedErrorLogPublisherRequest) HasOverrideSeverity() bool`

HasOverrideSeverity returns a boolean if a field has been set.

### GetDescription

`func (o *AddSyslogBasedErrorLogPublisherRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSyslogBasedErrorLogPublisherRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSyslogBasedErrorLogPublisherRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSyslogBasedErrorLogPublisherRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLoggingErrorBehavior

`func (o *AddSyslogBasedErrorLogPublisherRequest) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp`

GetLoggingErrorBehavior returns the LoggingErrorBehavior field if non-nil, zero value otherwise.

### GetLoggingErrorBehaviorOk

`func (o *AddSyslogBasedErrorLogPublisherRequest) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool)`

GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingErrorBehavior

`func (o *AddSyslogBasedErrorLogPublisherRequest) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp)`

SetLoggingErrorBehavior sets LoggingErrorBehavior field to given value.

### HasLoggingErrorBehavior

`func (o *AddSyslogBasedErrorLogPublisherRequest) HasLoggingErrorBehavior() bool`

HasLoggingErrorBehavior returns a boolean if a field has been set.

### GetPublisherName

`func (o *AddSyslogBasedErrorLogPublisherRequest) GetPublisherName() string`

GetPublisherName returns the PublisherName field if non-nil, zero value otherwise.

### GetPublisherNameOk

`func (o *AddSyslogBasedErrorLogPublisherRequest) GetPublisherNameOk() (*string, bool)`

GetPublisherNameOk returns a tuple with the PublisherName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherName

`func (o *AddSyslogBasedErrorLogPublisherRequest) SetPublisherName(v string)`

SetPublisherName sets PublisherName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


