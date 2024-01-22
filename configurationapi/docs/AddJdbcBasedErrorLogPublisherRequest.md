# AddJdbcBasedErrorLogPublisherRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumjdbcBasedErrorLogPublisherSchemaUrn**](EnumjdbcBasedErrorLogPublisherSchemaUrn.md) |  | 
**Server** | **string** | The JDBC-based Database Server to use for a connection. | 
**LogFieldMapping** | **string** | The log field mapping associates loggable fields to database column names. The table name is not part of this mapping. | 
**LogTableName** | Pointer to **string** | The table name to log entries to the database server. | [optional] 
**QueueSize** | Pointer to **int64** | The maximum number of log records that can be stored in the asynchronous queue. | [optional] 
**DefaultSeverity** | Pointer to [**[]EnumlogPublisherDefaultSeverityProp**](EnumlogPublisherDefaultSeverityProp.md) |  | [optional] 
**OverrideSeverity** | Pointer to **[]string** | Specifies the override severity levels for the logger based on the category of the messages. | [optional] 
**Description** | Pointer to **string** | A description for this Log Publisher | [optional] 
**Enabled** | **bool** | Indicates whether the Log Publisher is enabled for use. | 
**LoggingErrorBehavior** | Pointer to [**EnumlogPublisherLoggingErrorBehaviorProp**](EnumlogPublisherLoggingErrorBehaviorProp.md) |  | [optional] 
**PublisherName** | **string** | Name of the new Log Publisher | 

## Methods

### NewAddJdbcBasedErrorLogPublisherRequest

`func NewAddJdbcBasedErrorLogPublisherRequest(schemas []EnumjdbcBasedErrorLogPublisherSchemaUrn, server string, logFieldMapping string, enabled bool, publisherName string, ) *AddJdbcBasedErrorLogPublisherRequest`

NewAddJdbcBasedErrorLogPublisherRequest instantiates a new AddJdbcBasedErrorLogPublisherRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddJdbcBasedErrorLogPublisherRequestWithDefaults

`func NewAddJdbcBasedErrorLogPublisherRequestWithDefaults() *AddJdbcBasedErrorLogPublisherRequest`

NewAddJdbcBasedErrorLogPublisherRequestWithDefaults instantiates a new AddJdbcBasedErrorLogPublisherRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddJdbcBasedErrorLogPublisherRequest) GetSchemas() []EnumjdbcBasedErrorLogPublisherSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddJdbcBasedErrorLogPublisherRequest) GetSchemasOk() (*[]EnumjdbcBasedErrorLogPublisherSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddJdbcBasedErrorLogPublisherRequest) SetSchemas(v []EnumjdbcBasedErrorLogPublisherSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetServer

`func (o *AddJdbcBasedErrorLogPublisherRequest) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *AddJdbcBasedErrorLogPublisherRequest) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *AddJdbcBasedErrorLogPublisherRequest) SetServer(v string)`

SetServer sets Server field to given value.


### GetLogFieldMapping

`func (o *AddJdbcBasedErrorLogPublisherRequest) GetLogFieldMapping() string`

GetLogFieldMapping returns the LogFieldMapping field if non-nil, zero value otherwise.

### GetLogFieldMappingOk

`func (o *AddJdbcBasedErrorLogPublisherRequest) GetLogFieldMappingOk() (*string, bool)`

GetLogFieldMappingOk returns a tuple with the LogFieldMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldMapping

`func (o *AddJdbcBasedErrorLogPublisherRequest) SetLogFieldMapping(v string)`

SetLogFieldMapping sets LogFieldMapping field to given value.


### GetLogTableName

`func (o *AddJdbcBasedErrorLogPublisherRequest) GetLogTableName() string`

GetLogTableName returns the LogTableName field if non-nil, zero value otherwise.

### GetLogTableNameOk

`func (o *AddJdbcBasedErrorLogPublisherRequest) GetLogTableNameOk() (*string, bool)`

GetLogTableNameOk returns a tuple with the LogTableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogTableName

`func (o *AddJdbcBasedErrorLogPublisherRequest) SetLogTableName(v string)`

SetLogTableName sets LogTableName field to given value.

### HasLogTableName

`func (o *AddJdbcBasedErrorLogPublisherRequest) HasLogTableName() bool`

HasLogTableName returns a boolean if a field has been set.

### GetQueueSize

`func (o *AddJdbcBasedErrorLogPublisherRequest) GetQueueSize() int64`

GetQueueSize returns the QueueSize field if non-nil, zero value otherwise.

### GetQueueSizeOk

`func (o *AddJdbcBasedErrorLogPublisherRequest) GetQueueSizeOk() (*int64, bool)`

GetQueueSizeOk returns a tuple with the QueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueSize

`func (o *AddJdbcBasedErrorLogPublisherRequest) SetQueueSize(v int64)`

SetQueueSize sets QueueSize field to given value.

### HasQueueSize

`func (o *AddJdbcBasedErrorLogPublisherRequest) HasQueueSize() bool`

HasQueueSize returns a boolean if a field has been set.

### GetDefaultSeverity

`func (o *AddJdbcBasedErrorLogPublisherRequest) GetDefaultSeverity() []EnumlogPublisherDefaultSeverityProp`

GetDefaultSeverity returns the DefaultSeverity field if non-nil, zero value otherwise.

### GetDefaultSeverityOk

`func (o *AddJdbcBasedErrorLogPublisherRequest) GetDefaultSeverityOk() (*[]EnumlogPublisherDefaultSeverityProp, bool)`

GetDefaultSeverityOk returns a tuple with the DefaultSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSeverity

`func (o *AddJdbcBasedErrorLogPublisherRequest) SetDefaultSeverity(v []EnumlogPublisherDefaultSeverityProp)`

SetDefaultSeverity sets DefaultSeverity field to given value.

### HasDefaultSeverity

`func (o *AddJdbcBasedErrorLogPublisherRequest) HasDefaultSeverity() bool`

HasDefaultSeverity returns a boolean if a field has been set.

### GetOverrideSeverity

`func (o *AddJdbcBasedErrorLogPublisherRequest) GetOverrideSeverity() []string`

GetOverrideSeverity returns the OverrideSeverity field if non-nil, zero value otherwise.

### GetOverrideSeverityOk

`func (o *AddJdbcBasedErrorLogPublisherRequest) GetOverrideSeverityOk() (*[]string, bool)`

GetOverrideSeverityOk returns a tuple with the OverrideSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSeverity

`func (o *AddJdbcBasedErrorLogPublisherRequest) SetOverrideSeverity(v []string)`

SetOverrideSeverity sets OverrideSeverity field to given value.

### HasOverrideSeverity

`func (o *AddJdbcBasedErrorLogPublisherRequest) HasOverrideSeverity() bool`

HasOverrideSeverity returns a boolean if a field has been set.

### GetDescription

`func (o *AddJdbcBasedErrorLogPublisherRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddJdbcBasedErrorLogPublisherRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddJdbcBasedErrorLogPublisherRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddJdbcBasedErrorLogPublisherRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddJdbcBasedErrorLogPublisherRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddJdbcBasedErrorLogPublisherRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddJdbcBasedErrorLogPublisherRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLoggingErrorBehavior

`func (o *AddJdbcBasedErrorLogPublisherRequest) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp`

GetLoggingErrorBehavior returns the LoggingErrorBehavior field if non-nil, zero value otherwise.

### GetLoggingErrorBehaviorOk

`func (o *AddJdbcBasedErrorLogPublisherRequest) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool)`

GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingErrorBehavior

`func (o *AddJdbcBasedErrorLogPublisherRequest) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp)`

SetLoggingErrorBehavior sets LoggingErrorBehavior field to given value.

### HasLoggingErrorBehavior

`func (o *AddJdbcBasedErrorLogPublisherRequest) HasLoggingErrorBehavior() bool`

HasLoggingErrorBehavior returns a boolean if a field has been set.

### GetPublisherName

`func (o *AddJdbcBasedErrorLogPublisherRequest) GetPublisherName() string`

GetPublisherName returns the PublisherName field if non-nil, zero value otherwise.

### GetPublisherNameOk

`func (o *AddJdbcBasedErrorLogPublisherRequest) GetPublisherNameOk() (*string, bool)`

GetPublisherNameOk returns a tuple with the PublisherName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherName

`func (o *AddJdbcBasedErrorLogPublisherRequest) SetPublisherName(v string)`

SetPublisherName sets PublisherName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


