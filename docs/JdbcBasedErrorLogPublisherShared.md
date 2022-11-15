# JdbcBasedErrorLogPublisherShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumjdbcBasedErrorLogPublisherSchemaUrn**](EnumjdbcBasedErrorLogPublisherSchemaUrn.md) |  | 
**Server** | **string** | The JDBC-based Database Server to use for a connection. | 
**LogFieldMapping** | **string** | The log field mapping associates loggable fields to database column names. The table name is not part of this mapping. | 
**LogTableName** | **string** | The table name to log entries to the database server. | 
**QueueSize** | Pointer to **int32** | The maximum number of log records that can be stored in the asynchronous queue. | [optional] 
**DefaultSeverity** | Pointer to [**[]EnumlogPublisherDefaultSeverityProp**](EnumlogPublisherDefaultSeverityProp.md) |  | [optional] 
**OverrideSeverity** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Log Publisher | [optional] 
**Enabled** | **bool** | Indicates whether the Log Publisher is enabled for use. | 
**LoggingErrorBehavior** | Pointer to [**EnumlogPublisherLoggingErrorBehaviorProp**](EnumlogPublisherLoggingErrorBehaviorProp.md) |  | [optional] 

## Methods

### NewJdbcBasedErrorLogPublisherShared

`func NewJdbcBasedErrorLogPublisherShared(schemas []EnumjdbcBasedErrorLogPublisherSchemaUrn, server string, logFieldMapping string, logTableName string, enabled bool, ) *JdbcBasedErrorLogPublisherShared`

NewJdbcBasedErrorLogPublisherShared instantiates a new JdbcBasedErrorLogPublisherShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJdbcBasedErrorLogPublisherSharedWithDefaults

`func NewJdbcBasedErrorLogPublisherSharedWithDefaults() *JdbcBasedErrorLogPublisherShared`

NewJdbcBasedErrorLogPublisherSharedWithDefaults instantiates a new JdbcBasedErrorLogPublisherShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *JdbcBasedErrorLogPublisherShared) GetSchemas() []EnumjdbcBasedErrorLogPublisherSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *JdbcBasedErrorLogPublisherShared) GetSchemasOk() (*[]EnumjdbcBasedErrorLogPublisherSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *JdbcBasedErrorLogPublisherShared) SetSchemas(v []EnumjdbcBasedErrorLogPublisherSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetServer

`func (o *JdbcBasedErrorLogPublisherShared) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *JdbcBasedErrorLogPublisherShared) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *JdbcBasedErrorLogPublisherShared) SetServer(v string)`

SetServer sets Server field to given value.


### GetLogFieldMapping

`func (o *JdbcBasedErrorLogPublisherShared) GetLogFieldMapping() string`

GetLogFieldMapping returns the LogFieldMapping field if non-nil, zero value otherwise.

### GetLogFieldMappingOk

`func (o *JdbcBasedErrorLogPublisherShared) GetLogFieldMappingOk() (*string, bool)`

GetLogFieldMappingOk returns a tuple with the LogFieldMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldMapping

`func (o *JdbcBasedErrorLogPublisherShared) SetLogFieldMapping(v string)`

SetLogFieldMapping sets LogFieldMapping field to given value.


### GetLogTableName

`func (o *JdbcBasedErrorLogPublisherShared) GetLogTableName() string`

GetLogTableName returns the LogTableName field if non-nil, zero value otherwise.

### GetLogTableNameOk

`func (o *JdbcBasedErrorLogPublisherShared) GetLogTableNameOk() (*string, bool)`

GetLogTableNameOk returns a tuple with the LogTableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogTableName

`func (o *JdbcBasedErrorLogPublisherShared) SetLogTableName(v string)`

SetLogTableName sets LogTableName field to given value.


### GetQueueSize

`func (o *JdbcBasedErrorLogPublisherShared) GetQueueSize() int32`

GetQueueSize returns the QueueSize field if non-nil, zero value otherwise.

### GetQueueSizeOk

`func (o *JdbcBasedErrorLogPublisherShared) GetQueueSizeOk() (*int32, bool)`

GetQueueSizeOk returns a tuple with the QueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueSize

`func (o *JdbcBasedErrorLogPublisherShared) SetQueueSize(v int32)`

SetQueueSize sets QueueSize field to given value.

### HasQueueSize

`func (o *JdbcBasedErrorLogPublisherShared) HasQueueSize() bool`

HasQueueSize returns a boolean if a field has been set.

### GetDefaultSeverity

`func (o *JdbcBasedErrorLogPublisherShared) GetDefaultSeverity() []EnumlogPublisherDefaultSeverityProp`

GetDefaultSeverity returns the DefaultSeverity field if non-nil, zero value otherwise.

### GetDefaultSeverityOk

`func (o *JdbcBasedErrorLogPublisherShared) GetDefaultSeverityOk() (*[]EnumlogPublisherDefaultSeverityProp, bool)`

GetDefaultSeverityOk returns a tuple with the DefaultSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSeverity

`func (o *JdbcBasedErrorLogPublisherShared) SetDefaultSeverity(v []EnumlogPublisherDefaultSeverityProp)`

SetDefaultSeverity sets DefaultSeverity field to given value.

### HasDefaultSeverity

`func (o *JdbcBasedErrorLogPublisherShared) HasDefaultSeverity() bool`

HasDefaultSeverity returns a boolean if a field has been set.

### GetOverrideSeverity

`func (o *JdbcBasedErrorLogPublisherShared) GetOverrideSeverity() []string`

GetOverrideSeverity returns the OverrideSeverity field if non-nil, zero value otherwise.

### GetOverrideSeverityOk

`func (o *JdbcBasedErrorLogPublisherShared) GetOverrideSeverityOk() (*[]string, bool)`

GetOverrideSeverityOk returns a tuple with the OverrideSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSeverity

`func (o *JdbcBasedErrorLogPublisherShared) SetOverrideSeverity(v []string)`

SetOverrideSeverity sets OverrideSeverity field to given value.

### HasOverrideSeverity

`func (o *JdbcBasedErrorLogPublisherShared) HasOverrideSeverity() bool`

HasOverrideSeverity returns a boolean if a field has been set.

### GetDescription

`func (o *JdbcBasedErrorLogPublisherShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JdbcBasedErrorLogPublisherShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JdbcBasedErrorLogPublisherShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JdbcBasedErrorLogPublisherShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *JdbcBasedErrorLogPublisherShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *JdbcBasedErrorLogPublisherShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *JdbcBasedErrorLogPublisherShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLoggingErrorBehavior

`func (o *JdbcBasedErrorLogPublisherShared) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp`

GetLoggingErrorBehavior returns the LoggingErrorBehavior field if non-nil, zero value otherwise.

### GetLoggingErrorBehaviorOk

`func (o *JdbcBasedErrorLogPublisherShared) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool)`

GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingErrorBehavior

`func (o *JdbcBasedErrorLogPublisherShared) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp)`

SetLoggingErrorBehavior sets LoggingErrorBehavior field to given value.

### HasLoggingErrorBehavior

`func (o *JdbcBasedErrorLogPublisherShared) HasLoggingErrorBehavior() bool`

HasLoggingErrorBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


