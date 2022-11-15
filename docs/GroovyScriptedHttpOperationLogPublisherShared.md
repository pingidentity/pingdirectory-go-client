# GroovyScriptedHttpOperationLogPublisherShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumgroovyScriptedHttpOperationLogPublisherSchemaUrn**](EnumgroovyScriptedHttpOperationLogPublisherSchemaUrn.md) |  | 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted HTTP Operation Log Publisher. | 
**ScriptArgument** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Log Publisher | [optional] 
**Enabled** | **bool** | Indicates whether the Log Publisher is enabled for use. | 
**LoggingErrorBehavior** | Pointer to [**EnumlogPublisherLoggingErrorBehaviorProp**](EnumlogPublisherLoggingErrorBehaviorProp.md) |  | [optional] 

## Methods

### NewGroovyScriptedHttpOperationLogPublisherShared

`func NewGroovyScriptedHttpOperationLogPublisherShared(schemas []EnumgroovyScriptedHttpOperationLogPublisherSchemaUrn, scriptClass string, enabled bool, ) *GroovyScriptedHttpOperationLogPublisherShared`

NewGroovyScriptedHttpOperationLogPublisherShared instantiates a new GroovyScriptedHttpOperationLogPublisherShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroovyScriptedHttpOperationLogPublisherSharedWithDefaults

`func NewGroovyScriptedHttpOperationLogPublisherSharedWithDefaults() *GroovyScriptedHttpOperationLogPublisherShared`

NewGroovyScriptedHttpOperationLogPublisherSharedWithDefaults instantiates a new GroovyScriptedHttpOperationLogPublisherShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *GroovyScriptedHttpOperationLogPublisherShared) GetSchemas() []EnumgroovyScriptedHttpOperationLogPublisherSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GroovyScriptedHttpOperationLogPublisherShared) GetSchemasOk() (*[]EnumgroovyScriptedHttpOperationLogPublisherSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GroovyScriptedHttpOperationLogPublisherShared) SetSchemas(v []EnumgroovyScriptedHttpOperationLogPublisherSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetScriptClass

`func (o *GroovyScriptedHttpOperationLogPublisherShared) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *GroovyScriptedHttpOperationLogPublisherShared) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *GroovyScriptedHttpOperationLogPublisherShared) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *GroovyScriptedHttpOperationLogPublisherShared) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *GroovyScriptedHttpOperationLogPublisherShared) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *GroovyScriptedHttpOperationLogPublisherShared) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *GroovyScriptedHttpOperationLogPublisherShared) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetDescription

`func (o *GroovyScriptedHttpOperationLogPublisherShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroovyScriptedHttpOperationLogPublisherShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroovyScriptedHttpOperationLogPublisherShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroovyScriptedHttpOperationLogPublisherShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *GroovyScriptedHttpOperationLogPublisherShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GroovyScriptedHttpOperationLogPublisherShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GroovyScriptedHttpOperationLogPublisherShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLoggingErrorBehavior

`func (o *GroovyScriptedHttpOperationLogPublisherShared) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp`

GetLoggingErrorBehavior returns the LoggingErrorBehavior field if non-nil, zero value otherwise.

### GetLoggingErrorBehaviorOk

`func (o *GroovyScriptedHttpOperationLogPublisherShared) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool)`

GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingErrorBehavior

`func (o *GroovyScriptedHttpOperationLogPublisherShared) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp)`

SetLoggingErrorBehavior sets LoggingErrorBehavior field to given value.

### HasLoggingErrorBehavior

`func (o *GroovyScriptedHttpOperationLogPublisherShared) HasLoggingErrorBehavior() bool`

HasLoggingErrorBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


