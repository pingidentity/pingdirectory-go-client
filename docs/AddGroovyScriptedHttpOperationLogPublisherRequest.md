# AddGroovyScriptedHttpOperationLogPublisherRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublisherName** | **string** | Name of the new Log Publisher | 
**Schemas** | [**[]EnumgroovyScriptedHttpOperationLogPublisherSchemaUrn**](EnumgroovyScriptedHttpOperationLogPublisherSchemaUrn.md) |  | 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted HTTP Operation Log Publisher. | 
**ScriptArgument** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Log Publisher | [optional] 
**Enabled** | **bool** | Indicates whether the Log Publisher is enabled for use. | 
**LoggingErrorBehavior** | Pointer to [**EnumlogPublisherLoggingErrorBehaviorProp**](EnumlogPublisherLoggingErrorBehaviorProp.md) |  | [optional] 

## Methods

### NewAddGroovyScriptedHttpOperationLogPublisherRequest

`func NewAddGroovyScriptedHttpOperationLogPublisherRequest(publisherName string, schemas []EnumgroovyScriptedHttpOperationLogPublisherSchemaUrn, scriptClass string, enabled bool, ) *AddGroovyScriptedHttpOperationLogPublisherRequest`

NewAddGroovyScriptedHttpOperationLogPublisherRequest instantiates a new AddGroovyScriptedHttpOperationLogPublisherRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddGroovyScriptedHttpOperationLogPublisherRequestWithDefaults

`func NewAddGroovyScriptedHttpOperationLogPublisherRequestWithDefaults() *AddGroovyScriptedHttpOperationLogPublisherRequest`

NewAddGroovyScriptedHttpOperationLogPublisherRequestWithDefaults instantiates a new AddGroovyScriptedHttpOperationLogPublisherRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublisherName

`func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) GetPublisherName() string`

GetPublisherName returns the PublisherName field if non-nil, zero value otherwise.

### GetPublisherNameOk

`func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) GetPublisherNameOk() (*string, bool)`

GetPublisherNameOk returns a tuple with the PublisherName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherName

`func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) SetPublisherName(v string)`

SetPublisherName sets PublisherName field to given value.


### GetSchemas

`func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) GetSchemas() []EnumgroovyScriptedHttpOperationLogPublisherSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) GetSchemasOk() (*[]EnumgroovyScriptedHttpOperationLogPublisherSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) SetSchemas(v []EnumgroovyScriptedHttpOperationLogPublisherSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetScriptClass

`func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetDescription

`func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLoggingErrorBehavior

`func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp`

GetLoggingErrorBehavior returns the LoggingErrorBehavior field if non-nil, zero value otherwise.

### GetLoggingErrorBehaviorOk

`func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool)`

GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingErrorBehavior

`func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp)`

SetLoggingErrorBehavior sets LoggingErrorBehavior field to given value.

### HasLoggingErrorBehavior

`func (o *AddGroovyScriptedHttpOperationLogPublisherRequest) HasLoggingErrorBehavior() bool`

HasLoggingErrorBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


