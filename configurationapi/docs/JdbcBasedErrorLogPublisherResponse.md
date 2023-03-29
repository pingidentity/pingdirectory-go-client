# JdbcBasedErrorLogPublisherResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Log Publisher | 
**Schemas** | [**[]EnumjdbcBasedErrorLogPublisherSchemaUrn**](EnumjdbcBasedErrorLogPublisherSchemaUrn.md) |  | 
**Server** | **string** | The JDBC-based Database Server to use for a connection. | 
**LogFieldMapping** | **string** | The log field mapping associates loggable fields to database column names. The table name is not part of this mapping. | 
**LogTableName** | **string** | The table name to log entries to the database server. | 
**QueueSize** | Pointer to **int32** | The maximum number of log records that can be stored in the asynchronous queue. | [optional] 
**DefaultSeverity** | Pointer to [**[]EnumlogPublisherDefaultSeverityProp**](EnumlogPublisherDefaultSeverityProp.md) |  | [optional] 
**OverrideSeverity** | Pointer to **[]string** | Specifies the override severity levels for the logger based on the category of the messages. | [optional] 
**Description** | Pointer to **string** | A description for this Log Publisher | [optional] 
**Enabled** | **bool** | Indicates whether the Log Publisher is enabled for use. | 
**LoggingErrorBehavior** | Pointer to [**EnumlogPublisherLoggingErrorBehaviorProp**](EnumlogPublisherLoggingErrorBehaviorProp.md) |  | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewJdbcBasedErrorLogPublisherResponse

`func NewJdbcBasedErrorLogPublisherResponse(id string, schemas []EnumjdbcBasedErrorLogPublisherSchemaUrn, server string, logFieldMapping string, logTableName string, enabled bool, ) *JdbcBasedErrorLogPublisherResponse`

NewJdbcBasedErrorLogPublisherResponse instantiates a new JdbcBasedErrorLogPublisherResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJdbcBasedErrorLogPublisherResponseWithDefaults

`func NewJdbcBasedErrorLogPublisherResponseWithDefaults() *JdbcBasedErrorLogPublisherResponse`

NewJdbcBasedErrorLogPublisherResponseWithDefaults instantiates a new JdbcBasedErrorLogPublisherResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JdbcBasedErrorLogPublisherResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JdbcBasedErrorLogPublisherResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JdbcBasedErrorLogPublisherResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *JdbcBasedErrorLogPublisherResponse) GetSchemas() []EnumjdbcBasedErrorLogPublisherSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *JdbcBasedErrorLogPublisherResponse) GetSchemasOk() (*[]EnumjdbcBasedErrorLogPublisherSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *JdbcBasedErrorLogPublisherResponse) SetSchemas(v []EnumjdbcBasedErrorLogPublisherSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetServer

`func (o *JdbcBasedErrorLogPublisherResponse) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *JdbcBasedErrorLogPublisherResponse) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *JdbcBasedErrorLogPublisherResponse) SetServer(v string)`

SetServer sets Server field to given value.


### GetLogFieldMapping

`func (o *JdbcBasedErrorLogPublisherResponse) GetLogFieldMapping() string`

GetLogFieldMapping returns the LogFieldMapping field if non-nil, zero value otherwise.

### GetLogFieldMappingOk

`func (o *JdbcBasedErrorLogPublisherResponse) GetLogFieldMappingOk() (*string, bool)`

GetLogFieldMappingOk returns a tuple with the LogFieldMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldMapping

`func (o *JdbcBasedErrorLogPublisherResponse) SetLogFieldMapping(v string)`

SetLogFieldMapping sets LogFieldMapping field to given value.


### GetLogTableName

`func (o *JdbcBasedErrorLogPublisherResponse) GetLogTableName() string`

GetLogTableName returns the LogTableName field if non-nil, zero value otherwise.

### GetLogTableNameOk

`func (o *JdbcBasedErrorLogPublisherResponse) GetLogTableNameOk() (*string, bool)`

GetLogTableNameOk returns a tuple with the LogTableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogTableName

`func (o *JdbcBasedErrorLogPublisherResponse) SetLogTableName(v string)`

SetLogTableName sets LogTableName field to given value.


### GetQueueSize

`func (o *JdbcBasedErrorLogPublisherResponse) GetQueueSize() int32`

GetQueueSize returns the QueueSize field if non-nil, zero value otherwise.

### GetQueueSizeOk

`func (o *JdbcBasedErrorLogPublisherResponse) GetQueueSizeOk() (*int32, bool)`

GetQueueSizeOk returns a tuple with the QueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueSize

`func (o *JdbcBasedErrorLogPublisherResponse) SetQueueSize(v int32)`

SetQueueSize sets QueueSize field to given value.

### HasQueueSize

`func (o *JdbcBasedErrorLogPublisherResponse) HasQueueSize() bool`

HasQueueSize returns a boolean if a field has been set.

### GetDefaultSeverity

`func (o *JdbcBasedErrorLogPublisherResponse) GetDefaultSeverity() []EnumlogPublisherDefaultSeverityProp`

GetDefaultSeverity returns the DefaultSeverity field if non-nil, zero value otherwise.

### GetDefaultSeverityOk

`func (o *JdbcBasedErrorLogPublisherResponse) GetDefaultSeverityOk() (*[]EnumlogPublisherDefaultSeverityProp, bool)`

GetDefaultSeverityOk returns a tuple with the DefaultSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSeverity

`func (o *JdbcBasedErrorLogPublisherResponse) SetDefaultSeverity(v []EnumlogPublisherDefaultSeverityProp)`

SetDefaultSeverity sets DefaultSeverity field to given value.

### HasDefaultSeverity

`func (o *JdbcBasedErrorLogPublisherResponse) HasDefaultSeverity() bool`

HasDefaultSeverity returns a boolean if a field has been set.

### GetOverrideSeverity

`func (o *JdbcBasedErrorLogPublisherResponse) GetOverrideSeverity() []string`

GetOverrideSeverity returns the OverrideSeverity field if non-nil, zero value otherwise.

### GetOverrideSeverityOk

`func (o *JdbcBasedErrorLogPublisherResponse) GetOverrideSeverityOk() (*[]string, bool)`

GetOverrideSeverityOk returns a tuple with the OverrideSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSeverity

`func (o *JdbcBasedErrorLogPublisherResponse) SetOverrideSeverity(v []string)`

SetOverrideSeverity sets OverrideSeverity field to given value.

### HasOverrideSeverity

`func (o *JdbcBasedErrorLogPublisherResponse) HasOverrideSeverity() bool`

HasOverrideSeverity returns a boolean if a field has been set.

### GetDescription

`func (o *JdbcBasedErrorLogPublisherResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JdbcBasedErrorLogPublisherResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JdbcBasedErrorLogPublisherResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JdbcBasedErrorLogPublisherResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *JdbcBasedErrorLogPublisherResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *JdbcBasedErrorLogPublisherResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *JdbcBasedErrorLogPublisherResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLoggingErrorBehavior

`func (o *JdbcBasedErrorLogPublisherResponse) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp`

GetLoggingErrorBehavior returns the LoggingErrorBehavior field if non-nil, zero value otherwise.

### GetLoggingErrorBehaviorOk

`func (o *JdbcBasedErrorLogPublisherResponse) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool)`

GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingErrorBehavior

`func (o *JdbcBasedErrorLogPublisherResponse) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp)`

SetLoggingErrorBehavior sets LoggingErrorBehavior field to given value.

### HasLoggingErrorBehavior

`func (o *JdbcBasedErrorLogPublisherResponse) HasLoggingErrorBehavior() bool`

HasLoggingErrorBehavior returns a boolean if a field has been set.

### GetMeta

`func (o *JdbcBasedErrorLogPublisherResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *JdbcBasedErrorLogPublisherResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *JdbcBasedErrorLogPublisherResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *JdbcBasedErrorLogPublisherResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *JdbcBasedErrorLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *JdbcBasedErrorLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *JdbcBasedErrorLogPublisherResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *JdbcBasedErrorLogPublisherResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


