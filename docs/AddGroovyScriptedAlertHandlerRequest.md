# AddGroovyScriptedAlertHandlerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HandlerName** | **string** | Name of the new Alert Handler | 
**Schemas** | [**[]EnumgroovyScriptedAlertHandlerSchemaUrn**](EnumgroovyScriptedAlertHandlerSchemaUrn.md) |  | 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Alert Handler. | 
**ScriptArgument** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Alert Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Alert Handler is enabled. | 
**Asynchronous** | Pointer to **bool** | Indicates whether the server should attempt to invoke this Alert Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver the alert notification) will not delay whatever processing the server was performing when the alert was generated. | [optional] 
**EnabledAlertSeverity** | Pointer to [**[]EnumalertHandlerEnabledAlertSeverityProp**](EnumalertHandlerEnabledAlertSeverityProp.md) |  | [optional] 
**EnabledAlertType** | Pointer to [**[]EnumalertHandlerEnabledAlertTypeProp**](EnumalertHandlerEnabledAlertTypeProp.md) |  | [optional] 
**DisabledAlertType** | Pointer to [**[]EnumalertHandlerDisabledAlertTypeProp**](EnumalertHandlerDisabledAlertTypeProp.md) |  | [optional] 

## Methods

### NewAddGroovyScriptedAlertHandlerRequest

`func NewAddGroovyScriptedAlertHandlerRequest(handlerName string, schemas []EnumgroovyScriptedAlertHandlerSchemaUrn, scriptClass string, enabled bool, ) *AddGroovyScriptedAlertHandlerRequest`

NewAddGroovyScriptedAlertHandlerRequest instantiates a new AddGroovyScriptedAlertHandlerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddGroovyScriptedAlertHandlerRequestWithDefaults

`func NewAddGroovyScriptedAlertHandlerRequestWithDefaults() *AddGroovyScriptedAlertHandlerRequest`

NewAddGroovyScriptedAlertHandlerRequestWithDefaults instantiates a new AddGroovyScriptedAlertHandlerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandlerName

`func (o *AddGroovyScriptedAlertHandlerRequest) GetHandlerName() string`

GetHandlerName returns the HandlerName field if non-nil, zero value otherwise.

### GetHandlerNameOk

`func (o *AddGroovyScriptedAlertHandlerRequest) GetHandlerNameOk() (*string, bool)`

GetHandlerNameOk returns a tuple with the HandlerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerName

`func (o *AddGroovyScriptedAlertHandlerRequest) SetHandlerName(v string)`

SetHandlerName sets HandlerName field to given value.


### GetSchemas

`func (o *AddGroovyScriptedAlertHandlerRequest) GetSchemas() []EnumgroovyScriptedAlertHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddGroovyScriptedAlertHandlerRequest) GetSchemasOk() (*[]EnumgroovyScriptedAlertHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddGroovyScriptedAlertHandlerRequest) SetSchemas(v []EnumgroovyScriptedAlertHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetScriptClass

`func (o *AddGroovyScriptedAlertHandlerRequest) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *AddGroovyScriptedAlertHandlerRequest) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *AddGroovyScriptedAlertHandlerRequest) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *AddGroovyScriptedAlertHandlerRequest) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *AddGroovyScriptedAlertHandlerRequest) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *AddGroovyScriptedAlertHandlerRequest) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *AddGroovyScriptedAlertHandlerRequest) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetDescription

`func (o *AddGroovyScriptedAlertHandlerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddGroovyScriptedAlertHandlerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddGroovyScriptedAlertHandlerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddGroovyScriptedAlertHandlerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddGroovyScriptedAlertHandlerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddGroovyScriptedAlertHandlerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddGroovyScriptedAlertHandlerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAsynchronous

`func (o *AddGroovyScriptedAlertHandlerRequest) GetAsynchronous() bool`

GetAsynchronous returns the Asynchronous field if non-nil, zero value otherwise.

### GetAsynchronousOk

`func (o *AddGroovyScriptedAlertHandlerRequest) GetAsynchronousOk() (*bool, bool)`

GetAsynchronousOk returns a tuple with the Asynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsynchronous

`func (o *AddGroovyScriptedAlertHandlerRequest) SetAsynchronous(v bool)`

SetAsynchronous sets Asynchronous field to given value.

### HasAsynchronous

`func (o *AddGroovyScriptedAlertHandlerRequest) HasAsynchronous() bool`

HasAsynchronous returns a boolean if a field has been set.

### GetEnabledAlertSeverity

`func (o *AddGroovyScriptedAlertHandlerRequest) GetEnabledAlertSeverity() []EnumalertHandlerEnabledAlertSeverityProp`

GetEnabledAlertSeverity returns the EnabledAlertSeverity field if non-nil, zero value otherwise.

### GetEnabledAlertSeverityOk

`func (o *AddGroovyScriptedAlertHandlerRequest) GetEnabledAlertSeverityOk() (*[]EnumalertHandlerEnabledAlertSeverityProp, bool)`

GetEnabledAlertSeverityOk returns a tuple with the EnabledAlertSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAlertSeverity

`func (o *AddGroovyScriptedAlertHandlerRequest) SetEnabledAlertSeverity(v []EnumalertHandlerEnabledAlertSeverityProp)`

SetEnabledAlertSeverity sets EnabledAlertSeverity field to given value.

### HasEnabledAlertSeverity

`func (o *AddGroovyScriptedAlertHandlerRequest) HasEnabledAlertSeverity() bool`

HasEnabledAlertSeverity returns a boolean if a field has been set.

### GetEnabledAlertType

`func (o *AddGroovyScriptedAlertHandlerRequest) GetEnabledAlertType() []EnumalertHandlerEnabledAlertTypeProp`

GetEnabledAlertType returns the EnabledAlertType field if non-nil, zero value otherwise.

### GetEnabledAlertTypeOk

`func (o *AddGroovyScriptedAlertHandlerRequest) GetEnabledAlertTypeOk() (*[]EnumalertHandlerEnabledAlertTypeProp, bool)`

GetEnabledAlertTypeOk returns a tuple with the EnabledAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAlertType

`func (o *AddGroovyScriptedAlertHandlerRequest) SetEnabledAlertType(v []EnumalertHandlerEnabledAlertTypeProp)`

SetEnabledAlertType sets EnabledAlertType field to given value.

### HasEnabledAlertType

`func (o *AddGroovyScriptedAlertHandlerRequest) HasEnabledAlertType() bool`

HasEnabledAlertType returns a boolean if a field has been set.

### GetDisabledAlertType

`func (o *AddGroovyScriptedAlertHandlerRequest) GetDisabledAlertType() []EnumalertHandlerDisabledAlertTypeProp`

GetDisabledAlertType returns the DisabledAlertType field if non-nil, zero value otherwise.

### GetDisabledAlertTypeOk

`func (o *AddGroovyScriptedAlertHandlerRequest) GetDisabledAlertTypeOk() (*[]EnumalertHandlerDisabledAlertTypeProp, bool)`

GetDisabledAlertTypeOk returns a tuple with the DisabledAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAlertType

`func (o *AddGroovyScriptedAlertHandlerRequest) SetDisabledAlertType(v []EnumalertHandlerDisabledAlertTypeProp)`

SetDisabledAlertType sets DisabledAlertType field to given value.

### HasDisabledAlertType

`func (o *AddGroovyScriptedAlertHandlerRequest) HasDisabledAlertType() bool`

HasDisabledAlertType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


