# AddGroovyScriptedErrorLogPublisherRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublisherName** | **string** | Name of the new Log Publisher | 
**Schemas** | [**[]EnumgroovyScriptedErrorLogPublisherSchemaUrn**](EnumgroovyScriptedErrorLogPublisherSchemaUrn.md) |  | 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Error Log Publisher. | 
**ScriptArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Scripted Error Log Publisher. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**DefaultSeverity** | Pointer to [**[]EnumlogPublisherDefaultSeverityProp**](EnumlogPublisherDefaultSeverityProp.md) | Specifies the default severity levels for the logger. | [optional] 
**OverrideSeverity** | Pointer to **[]string** | Specifies the override severity levels for the logger based on the category of the messages. | [optional] 
**Description** | Pointer to **string** | A description for this Log Publisher | [optional] 
**Enabled** | **bool** | Indicates whether the Log Publisher is enabled for use. | 
**LoggingErrorBehavior** | Pointer to [**EnumlogPublisherLoggingErrorBehaviorProp**](EnumlogPublisherLoggingErrorBehaviorProp.md) |  | [optional] 

## Methods

### NewAddGroovyScriptedErrorLogPublisherRequest

`func NewAddGroovyScriptedErrorLogPublisherRequest(publisherName string, schemas []EnumgroovyScriptedErrorLogPublisherSchemaUrn, scriptClass string, enabled bool, ) *AddGroovyScriptedErrorLogPublisherRequest`

NewAddGroovyScriptedErrorLogPublisherRequest instantiates a new AddGroovyScriptedErrorLogPublisherRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddGroovyScriptedErrorLogPublisherRequestWithDefaults

`func NewAddGroovyScriptedErrorLogPublisherRequestWithDefaults() *AddGroovyScriptedErrorLogPublisherRequest`

NewAddGroovyScriptedErrorLogPublisherRequestWithDefaults instantiates a new AddGroovyScriptedErrorLogPublisherRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublisherName

`func (o *AddGroovyScriptedErrorLogPublisherRequest) GetPublisherName() string`

GetPublisherName returns the PublisherName field if non-nil, zero value otherwise.

### GetPublisherNameOk

`func (o *AddGroovyScriptedErrorLogPublisherRequest) GetPublisherNameOk() (*string, bool)`

GetPublisherNameOk returns a tuple with the PublisherName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherName

`func (o *AddGroovyScriptedErrorLogPublisherRequest) SetPublisherName(v string)`

SetPublisherName sets PublisherName field to given value.


### GetSchemas

`func (o *AddGroovyScriptedErrorLogPublisherRequest) GetSchemas() []EnumgroovyScriptedErrorLogPublisherSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddGroovyScriptedErrorLogPublisherRequest) GetSchemasOk() (*[]EnumgroovyScriptedErrorLogPublisherSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddGroovyScriptedErrorLogPublisherRequest) SetSchemas(v []EnumgroovyScriptedErrorLogPublisherSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetScriptClass

`func (o *AddGroovyScriptedErrorLogPublisherRequest) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *AddGroovyScriptedErrorLogPublisherRequest) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *AddGroovyScriptedErrorLogPublisherRequest) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *AddGroovyScriptedErrorLogPublisherRequest) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *AddGroovyScriptedErrorLogPublisherRequest) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *AddGroovyScriptedErrorLogPublisherRequest) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *AddGroovyScriptedErrorLogPublisherRequest) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetDefaultSeverity

`func (o *AddGroovyScriptedErrorLogPublisherRequest) GetDefaultSeverity() []EnumlogPublisherDefaultSeverityProp`

GetDefaultSeverity returns the DefaultSeverity field if non-nil, zero value otherwise.

### GetDefaultSeverityOk

`func (o *AddGroovyScriptedErrorLogPublisherRequest) GetDefaultSeverityOk() (*[]EnumlogPublisherDefaultSeverityProp, bool)`

GetDefaultSeverityOk returns a tuple with the DefaultSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSeverity

`func (o *AddGroovyScriptedErrorLogPublisherRequest) SetDefaultSeverity(v []EnumlogPublisherDefaultSeverityProp)`

SetDefaultSeverity sets DefaultSeverity field to given value.

### HasDefaultSeverity

`func (o *AddGroovyScriptedErrorLogPublisherRequest) HasDefaultSeverity() bool`

HasDefaultSeverity returns a boolean if a field has been set.

### GetOverrideSeverity

`func (o *AddGroovyScriptedErrorLogPublisherRequest) GetOverrideSeverity() []string`

GetOverrideSeverity returns the OverrideSeverity field if non-nil, zero value otherwise.

### GetOverrideSeverityOk

`func (o *AddGroovyScriptedErrorLogPublisherRequest) GetOverrideSeverityOk() (*[]string, bool)`

GetOverrideSeverityOk returns a tuple with the OverrideSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSeverity

`func (o *AddGroovyScriptedErrorLogPublisherRequest) SetOverrideSeverity(v []string)`

SetOverrideSeverity sets OverrideSeverity field to given value.

### HasOverrideSeverity

`func (o *AddGroovyScriptedErrorLogPublisherRequest) HasOverrideSeverity() bool`

HasOverrideSeverity returns a boolean if a field has been set.

### GetDescription

`func (o *AddGroovyScriptedErrorLogPublisherRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddGroovyScriptedErrorLogPublisherRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddGroovyScriptedErrorLogPublisherRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddGroovyScriptedErrorLogPublisherRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddGroovyScriptedErrorLogPublisherRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddGroovyScriptedErrorLogPublisherRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddGroovyScriptedErrorLogPublisherRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLoggingErrorBehavior

`func (o *AddGroovyScriptedErrorLogPublisherRequest) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp`

GetLoggingErrorBehavior returns the LoggingErrorBehavior field if non-nil, zero value otherwise.

### GetLoggingErrorBehaviorOk

`func (o *AddGroovyScriptedErrorLogPublisherRequest) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool)`

GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingErrorBehavior

`func (o *AddGroovyScriptedErrorLogPublisherRequest) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp)`

SetLoggingErrorBehavior sets LoggingErrorBehavior field to given value.

### HasLoggingErrorBehavior

`func (o *AddGroovyScriptedErrorLogPublisherRequest) HasLoggingErrorBehavior() bool`

HasLoggingErrorBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


