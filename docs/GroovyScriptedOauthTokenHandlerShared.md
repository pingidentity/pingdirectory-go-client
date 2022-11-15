# GroovyScriptedOauthTokenHandlerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumgroovyScriptedOauthTokenHandlerSchemaUrn**](EnumgroovyScriptedOauthTokenHandlerSchemaUrn.md) |  | 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted OAuth Token Handler. | 
**ScriptArgument** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this OAuth Token Handler | [optional] 

## Methods

### NewGroovyScriptedOauthTokenHandlerShared

`func NewGroovyScriptedOauthTokenHandlerShared(schemas []EnumgroovyScriptedOauthTokenHandlerSchemaUrn, scriptClass string, ) *GroovyScriptedOauthTokenHandlerShared`

NewGroovyScriptedOauthTokenHandlerShared instantiates a new GroovyScriptedOauthTokenHandlerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroovyScriptedOauthTokenHandlerSharedWithDefaults

`func NewGroovyScriptedOauthTokenHandlerSharedWithDefaults() *GroovyScriptedOauthTokenHandlerShared`

NewGroovyScriptedOauthTokenHandlerSharedWithDefaults instantiates a new GroovyScriptedOauthTokenHandlerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *GroovyScriptedOauthTokenHandlerShared) GetSchemas() []EnumgroovyScriptedOauthTokenHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GroovyScriptedOauthTokenHandlerShared) GetSchemasOk() (*[]EnumgroovyScriptedOauthTokenHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GroovyScriptedOauthTokenHandlerShared) SetSchemas(v []EnumgroovyScriptedOauthTokenHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetScriptClass

`func (o *GroovyScriptedOauthTokenHandlerShared) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *GroovyScriptedOauthTokenHandlerShared) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *GroovyScriptedOauthTokenHandlerShared) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *GroovyScriptedOauthTokenHandlerShared) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *GroovyScriptedOauthTokenHandlerShared) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *GroovyScriptedOauthTokenHandlerShared) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *GroovyScriptedOauthTokenHandlerShared) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetDescription

`func (o *GroovyScriptedOauthTokenHandlerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroovyScriptedOauthTokenHandlerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroovyScriptedOauthTokenHandlerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroovyScriptedOauthTokenHandlerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


