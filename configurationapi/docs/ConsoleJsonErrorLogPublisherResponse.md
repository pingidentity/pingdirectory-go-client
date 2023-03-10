# ConsoleJsonErrorLogPublisherResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumconsoleJsonErrorLogPublisherSchemaUrn**](EnumconsoleJsonErrorLogPublisherSchemaUrn.md) |  | 
**Id** | **string** | Name of the Log Publisher | 
**Enabled** | **bool** | Indicates whether the Console JSON Error Log Publisher is enabled for use. | 
**DefaultSeverity** | Pointer to [**[]EnumlogPublisherDefaultSeverityProp**](EnumlogPublisherDefaultSeverityProp.md) |  | [optional] 
**WriteMultiLineMessages** | Pointer to **bool** | Indicates whether the JSON objects should be formatted to span multiple lines with a single element on each line. The multi-line format is potentially more user friendly (if administrators may need to look at the log files), but each message will be larger because of the additional spaces and end-of-line markers. | [optional] 
**OutputLocation** | Pointer to [**EnumlogPublisherOutputLocationProp**](EnumlogPublisherOutputLocationProp.md) |  | [optional] 
**IncludeProductName** | Pointer to **bool** | Indicates whether log messages should include the product name for the Directory Server. | [optional] 
**IncludeInstanceName** | Pointer to **bool** | Indicates whether log messages should include the instance name for the Directory Server. | [optional] 
**IncludeStartupID** | Pointer to **bool** | Indicates whether log messages should include the startup ID for the Directory Server, which is a value assigned to the server instance at startup and may be used to identify when the server has been restarted. | [optional] 
**IncludeThreadID** | Pointer to **bool** | Indicates whether log messages should include the thread ID for the Directory Server in each log message. This ID can be used to correlate log messages from the same thread within a single log as well as generated by the same thread across different types of log files. More information about the thread with a specific ID can be obtained using the cn&#x3D;JVM Stack Trace,cn&#x3D;monitor entry. | [optional] 
**GenerifyMessageStringsWhenPossible** | Pointer to **bool** | Indicates whether to use the generified version of the log message string (which may use placeholders like %s for a string or %d for an integer), rather than the version of the message with those placeholders replaced with specific values that would normally be written to the log. | [optional] 
**OverrideSeverity** | Pointer to **[]string** | Specifies the override severity levels for the logger based on the category of the messages. | [optional] 
**Description** | Pointer to **string** | A description for this Log Publisher | [optional] 
**LoggingErrorBehavior** | Pointer to [**EnumlogPublisherLoggingErrorBehaviorProp**](EnumlogPublisherLoggingErrorBehaviorProp.md) |  | [optional] 

## Methods

### NewConsoleJsonErrorLogPublisherResponse

`func NewConsoleJsonErrorLogPublisherResponse(schemas []EnumconsoleJsonErrorLogPublisherSchemaUrn, id string, enabled bool, ) *ConsoleJsonErrorLogPublisherResponse`

NewConsoleJsonErrorLogPublisherResponse instantiates a new ConsoleJsonErrorLogPublisherResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleJsonErrorLogPublisherResponseWithDefaults

`func NewConsoleJsonErrorLogPublisherResponseWithDefaults() *ConsoleJsonErrorLogPublisherResponse`

NewConsoleJsonErrorLogPublisherResponseWithDefaults instantiates a new ConsoleJsonErrorLogPublisherResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ConsoleJsonErrorLogPublisherResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ConsoleJsonErrorLogPublisherResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ConsoleJsonErrorLogPublisherResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ConsoleJsonErrorLogPublisherResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ConsoleJsonErrorLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ConsoleJsonErrorLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ConsoleJsonErrorLogPublisherResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ConsoleJsonErrorLogPublisherResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *ConsoleJsonErrorLogPublisherResponse) GetSchemas() []EnumconsoleJsonErrorLogPublisherSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ConsoleJsonErrorLogPublisherResponse) GetSchemasOk() (*[]EnumconsoleJsonErrorLogPublisherSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ConsoleJsonErrorLogPublisherResponse) SetSchemas(v []EnumconsoleJsonErrorLogPublisherSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *ConsoleJsonErrorLogPublisherResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsoleJsonErrorLogPublisherResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsoleJsonErrorLogPublisherResponse) SetId(v string)`

SetId sets Id field to given value.


### GetEnabled

`func (o *ConsoleJsonErrorLogPublisherResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ConsoleJsonErrorLogPublisherResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ConsoleJsonErrorLogPublisherResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetDefaultSeverity

`func (o *ConsoleJsonErrorLogPublisherResponse) GetDefaultSeverity() []EnumlogPublisherDefaultSeverityProp`

GetDefaultSeverity returns the DefaultSeverity field if non-nil, zero value otherwise.

### GetDefaultSeverityOk

`func (o *ConsoleJsonErrorLogPublisherResponse) GetDefaultSeverityOk() (*[]EnumlogPublisherDefaultSeverityProp, bool)`

GetDefaultSeverityOk returns a tuple with the DefaultSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSeverity

`func (o *ConsoleJsonErrorLogPublisherResponse) SetDefaultSeverity(v []EnumlogPublisherDefaultSeverityProp)`

SetDefaultSeverity sets DefaultSeverity field to given value.

### HasDefaultSeverity

`func (o *ConsoleJsonErrorLogPublisherResponse) HasDefaultSeverity() bool`

HasDefaultSeverity returns a boolean if a field has been set.

### GetWriteMultiLineMessages

`func (o *ConsoleJsonErrorLogPublisherResponse) GetWriteMultiLineMessages() bool`

GetWriteMultiLineMessages returns the WriteMultiLineMessages field if non-nil, zero value otherwise.

### GetWriteMultiLineMessagesOk

`func (o *ConsoleJsonErrorLogPublisherResponse) GetWriteMultiLineMessagesOk() (*bool, bool)`

GetWriteMultiLineMessagesOk returns a tuple with the WriteMultiLineMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteMultiLineMessages

`func (o *ConsoleJsonErrorLogPublisherResponse) SetWriteMultiLineMessages(v bool)`

SetWriteMultiLineMessages sets WriteMultiLineMessages field to given value.

### HasWriteMultiLineMessages

`func (o *ConsoleJsonErrorLogPublisherResponse) HasWriteMultiLineMessages() bool`

HasWriteMultiLineMessages returns a boolean if a field has been set.

### GetOutputLocation

`func (o *ConsoleJsonErrorLogPublisherResponse) GetOutputLocation() EnumlogPublisherOutputLocationProp`

GetOutputLocation returns the OutputLocation field if non-nil, zero value otherwise.

### GetOutputLocationOk

`func (o *ConsoleJsonErrorLogPublisherResponse) GetOutputLocationOk() (*EnumlogPublisherOutputLocationProp, bool)`

GetOutputLocationOk returns a tuple with the OutputLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputLocation

`func (o *ConsoleJsonErrorLogPublisherResponse) SetOutputLocation(v EnumlogPublisherOutputLocationProp)`

SetOutputLocation sets OutputLocation field to given value.

### HasOutputLocation

`func (o *ConsoleJsonErrorLogPublisherResponse) HasOutputLocation() bool`

HasOutputLocation returns a boolean if a field has been set.

### GetIncludeProductName

`func (o *ConsoleJsonErrorLogPublisherResponse) GetIncludeProductName() bool`

GetIncludeProductName returns the IncludeProductName field if non-nil, zero value otherwise.

### GetIncludeProductNameOk

`func (o *ConsoleJsonErrorLogPublisherResponse) GetIncludeProductNameOk() (*bool, bool)`

GetIncludeProductNameOk returns a tuple with the IncludeProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeProductName

`func (o *ConsoleJsonErrorLogPublisherResponse) SetIncludeProductName(v bool)`

SetIncludeProductName sets IncludeProductName field to given value.

### HasIncludeProductName

`func (o *ConsoleJsonErrorLogPublisherResponse) HasIncludeProductName() bool`

HasIncludeProductName returns a boolean if a field has been set.

### GetIncludeInstanceName

`func (o *ConsoleJsonErrorLogPublisherResponse) GetIncludeInstanceName() bool`

GetIncludeInstanceName returns the IncludeInstanceName field if non-nil, zero value otherwise.

### GetIncludeInstanceNameOk

`func (o *ConsoleJsonErrorLogPublisherResponse) GetIncludeInstanceNameOk() (*bool, bool)`

GetIncludeInstanceNameOk returns a tuple with the IncludeInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInstanceName

`func (o *ConsoleJsonErrorLogPublisherResponse) SetIncludeInstanceName(v bool)`

SetIncludeInstanceName sets IncludeInstanceName field to given value.

### HasIncludeInstanceName

`func (o *ConsoleJsonErrorLogPublisherResponse) HasIncludeInstanceName() bool`

HasIncludeInstanceName returns a boolean if a field has been set.

### GetIncludeStartupID

`func (o *ConsoleJsonErrorLogPublisherResponse) GetIncludeStartupID() bool`

GetIncludeStartupID returns the IncludeStartupID field if non-nil, zero value otherwise.

### GetIncludeStartupIDOk

`func (o *ConsoleJsonErrorLogPublisherResponse) GetIncludeStartupIDOk() (*bool, bool)`

GetIncludeStartupIDOk returns a tuple with the IncludeStartupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeStartupID

`func (o *ConsoleJsonErrorLogPublisherResponse) SetIncludeStartupID(v bool)`

SetIncludeStartupID sets IncludeStartupID field to given value.

### HasIncludeStartupID

`func (o *ConsoleJsonErrorLogPublisherResponse) HasIncludeStartupID() bool`

HasIncludeStartupID returns a boolean if a field has been set.

### GetIncludeThreadID

`func (o *ConsoleJsonErrorLogPublisherResponse) GetIncludeThreadID() bool`

GetIncludeThreadID returns the IncludeThreadID field if non-nil, zero value otherwise.

### GetIncludeThreadIDOk

`func (o *ConsoleJsonErrorLogPublisherResponse) GetIncludeThreadIDOk() (*bool, bool)`

GetIncludeThreadIDOk returns a tuple with the IncludeThreadID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeThreadID

`func (o *ConsoleJsonErrorLogPublisherResponse) SetIncludeThreadID(v bool)`

SetIncludeThreadID sets IncludeThreadID field to given value.

### HasIncludeThreadID

`func (o *ConsoleJsonErrorLogPublisherResponse) HasIncludeThreadID() bool`

HasIncludeThreadID returns a boolean if a field has been set.

### GetGenerifyMessageStringsWhenPossible

`func (o *ConsoleJsonErrorLogPublisherResponse) GetGenerifyMessageStringsWhenPossible() bool`

GetGenerifyMessageStringsWhenPossible returns the GenerifyMessageStringsWhenPossible field if non-nil, zero value otherwise.

### GetGenerifyMessageStringsWhenPossibleOk

`func (o *ConsoleJsonErrorLogPublisherResponse) GetGenerifyMessageStringsWhenPossibleOk() (*bool, bool)`

GetGenerifyMessageStringsWhenPossibleOk returns a tuple with the GenerifyMessageStringsWhenPossible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerifyMessageStringsWhenPossible

`func (o *ConsoleJsonErrorLogPublisherResponse) SetGenerifyMessageStringsWhenPossible(v bool)`

SetGenerifyMessageStringsWhenPossible sets GenerifyMessageStringsWhenPossible field to given value.

### HasGenerifyMessageStringsWhenPossible

`func (o *ConsoleJsonErrorLogPublisherResponse) HasGenerifyMessageStringsWhenPossible() bool`

HasGenerifyMessageStringsWhenPossible returns a boolean if a field has been set.

### GetOverrideSeverity

`func (o *ConsoleJsonErrorLogPublisherResponse) GetOverrideSeverity() []string`

GetOverrideSeverity returns the OverrideSeverity field if non-nil, zero value otherwise.

### GetOverrideSeverityOk

`func (o *ConsoleJsonErrorLogPublisherResponse) GetOverrideSeverityOk() (*[]string, bool)`

GetOverrideSeverityOk returns a tuple with the OverrideSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSeverity

`func (o *ConsoleJsonErrorLogPublisherResponse) SetOverrideSeverity(v []string)`

SetOverrideSeverity sets OverrideSeverity field to given value.

### HasOverrideSeverity

`func (o *ConsoleJsonErrorLogPublisherResponse) HasOverrideSeverity() bool`

HasOverrideSeverity returns a boolean if a field has been set.

### GetDescription

`func (o *ConsoleJsonErrorLogPublisherResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConsoleJsonErrorLogPublisherResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConsoleJsonErrorLogPublisherResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConsoleJsonErrorLogPublisherResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLoggingErrorBehavior

`func (o *ConsoleJsonErrorLogPublisherResponse) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp`

GetLoggingErrorBehavior returns the LoggingErrorBehavior field if non-nil, zero value otherwise.

### GetLoggingErrorBehaviorOk

`func (o *ConsoleJsonErrorLogPublisherResponse) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool)`

GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingErrorBehavior

`func (o *ConsoleJsonErrorLogPublisherResponse) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp)`

SetLoggingErrorBehavior sets LoggingErrorBehavior field to given value.

### HasLoggingErrorBehavior

`func (o *ConsoleJsonErrorLogPublisherResponse) HasLoggingErrorBehavior() bool`

HasLoggingErrorBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


