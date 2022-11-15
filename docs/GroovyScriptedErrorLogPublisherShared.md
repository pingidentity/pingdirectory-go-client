# GroovyScriptedErrorLogPublisherShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumgroovyScriptedErrorLogPublisherSchemaUrn**](EnumgroovyScriptedErrorLogPublisherSchemaUrn.md) |  | 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Error Log Publisher. | 
**ScriptArgument** | Pointer to **[]string** |  | [optional] 
**DefaultSeverity** | Pointer to [**[]EnumlogPublisherDefaultSeverityProp**](EnumlogPublisherDefaultSeverityProp.md) |  | [optional] 
**OverrideSeverity** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Log Publisher | [optional] 
**Enabled** | **bool** | Indicates whether the Log Publisher is enabled for use. | 
**LoggingErrorBehavior** | Pointer to [**EnumlogPublisherLoggingErrorBehaviorProp**](EnumlogPublisherLoggingErrorBehaviorProp.md) |  | [optional] 

## Methods

### NewGroovyScriptedErrorLogPublisherShared

`func NewGroovyScriptedErrorLogPublisherShared(schemas []EnumgroovyScriptedErrorLogPublisherSchemaUrn, scriptClass string, enabled bool, ) *GroovyScriptedErrorLogPublisherShared`

NewGroovyScriptedErrorLogPublisherShared instantiates a new GroovyScriptedErrorLogPublisherShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroovyScriptedErrorLogPublisherSharedWithDefaults

`func NewGroovyScriptedErrorLogPublisherSharedWithDefaults() *GroovyScriptedErrorLogPublisherShared`

NewGroovyScriptedErrorLogPublisherSharedWithDefaults instantiates a new GroovyScriptedErrorLogPublisherShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *GroovyScriptedErrorLogPublisherShared) GetSchemas() []EnumgroovyScriptedErrorLogPublisherSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GroovyScriptedErrorLogPublisherShared) GetSchemasOk() (*[]EnumgroovyScriptedErrorLogPublisherSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GroovyScriptedErrorLogPublisherShared) SetSchemas(v []EnumgroovyScriptedErrorLogPublisherSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetScriptClass

`func (o *GroovyScriptedErrorLogPublisherShared) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *GroovyScriptedErrorLogPublisherShared) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *GroovyScriptedErrorLogPublisherShared) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *GroovyScriptedErrorLogPublisherShared) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *GroovyScriptedErrorLogPublisherShared) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *GroovyScriptedErrorLogPublisherShared) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *GroovyScriptedErrorLogPublisherShared) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetDefaultSeverity

`func (o *GroovyScriptedErrorLogPublisherShared) GetDefaultSeverity() []EnumlogPublisherDefaultSeverityProp`

GetDefaultSeverity returns the DefaultSeverity field if non-nil, zero value otherwise.

### GetDefaultSeverityOk

`func (o *GroovyScriptedErrorLogPublisherShared) GetDefaultSeverityOk() (*[]EnumlogPublisherDefaultSeverityProp, bool)`

GetDefaultSeverityOk returns a tuple with the DefaultSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSeverity

`func (o *GroovyScriptedErrorLogPublisherShared) SetDefaultSeverity(v []EnumlogPublisherDefaultSeverityProp)`

SetDefaultSeverity sets DefaultSeverity field to given value.

### HasDefaultSeverity

`func (o *GroovyScriptedErrorLogPublisherShared) HasDefaultSeverity() bool`

HasDefaultSeverity returns a boolean if a field has been set.

### GetOverrideSeverity

`func (o *GroovyScriptedErrorLogPublisherShared) GetOverrideSeverity() []string`

GetOverrideSeverity returns the OverrideSeverity field if non-nil, zero value otherwise.

### GetOverrideSeverityOk

`func (o *GroovyScriptedErrorLogPublisherShared) GetOverrideSeverityOk() (*[]string, bool)`

GetOverrideSeverityOk returns a tuple with the OverrideSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSeverity

`func (o *GroovyScriptedErrorLogPublisherShared) SetOverrideSeverity(v []string)`

SetOverrideSeverity sets OverrideSeverity field to given value.

### HasOverrideSeverity

`func (o *GroovyScriptedErrorLogPublisherShared) HasOverrideSeverity() bool`

HasOverrideSeverity returns a boolean if a field has been set.

### GetDescription

`func (o *GroovyScriptedErrorLogPublisherShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroovyScriptedErrorLogPublisherShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroovyScriptedErrorLogPublisherShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroovyScriptedErrorLogPublisherShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *GroovyScriptedErrorLogPublisherShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GroovyScriptedErrorLogPublisherShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GroovyScriptedErrorLogPublisherShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLoggingErrorBehavior

`func (o *GroovyScriptedErrorLogPublisherShared) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp`

GetLoggingErrorBehavior returns the LoggingErrorBehavior field if non-nil, zero value otherwise.

### GetLoggingErrorBehaviorOk

`func (o *GroovyScriptedErrorLogPublisherShared) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool)`

GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingErrorBehavior

`func (o *GroovyScriptedErrorLogPublisherShared) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp)`

SetLoggingErrorBehavior sets LoggingErrorBehavior field to given value.

### HasLoggingErrorBehavior

`func (o *GroovyScriptedErrorLogPublisherShared) HasLoggingErrorBehavior() bool`

HasLoggingErrorBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


