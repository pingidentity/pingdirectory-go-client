# AddGroovyScriptedOauthTokenHandlerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HandlerName** | **string** | Name of the new OAuth Token Handler | 
**Schemas** | [**[]EnumgroovyScriptedOauthTokenHandlerSchemaUrn**](EnumgroovyScriptedOauthTokenHandlerSchemaUrn.md) |  | 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted OAuth Token Handler. | 
**ScriptArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Scripted OAuth Token Handler. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**Description** | Pointer to **string** | A description for this OAuth Token Handler | [optional] 

## Methods

### NewAddGroovyScriptedOauthTokenHandlerRequest

`func NewAddGroovyScriptedOauthTokenHandlerRequest(handlerName string, schemas []EnumgroovyScriptedOauthTokenHandlerSchemaUrn, scriptClass string, ) *AddGroovyScriptedOauthTokenHandlerRequest`

NewAddGroovyScriptedOauthTokenHandlerRequest instantiates a new AddGroovyScriptedOauthTokenHandlerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddGroovyScriptedOauthTokenHandlerRequestWithDefaults

`func NewAddGroovyScriptedOauthTokenHandlerRequestWithDefaults() *AddGroovyScriptedOauthTokenHandlerRequest`

NewAddGroovyScriptedOauthTokenHandlerRequestWithDefaults instantiates a new AddGroovyScriptedOauthTokenHandlerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandlerName

`func (o *AddGroovyScriptedOauthTokenHandlerRequest) GetHandlerName() string`

GetHandlerName returns the HandlerName field if non-nil, zero value otherwise.

### GetHandlerNameOk

`func (o *AddGroovyScriptedOauthTokenHandlerRequest) GetHandlerNameOk() (*string, bool)`

GetHandlerNameOk returns a tuple with the HandlerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerName

`func (o *AddGroovyScriptedOauthTokenHandlerRequest) SetHandlerName(v string)`

SetHandlerName sets HandlerName field to given value.


### GetSchemas

`func (o *AddGroovyScriptedOauthTokenHandlerRequest) GetSchemas() []EnumgroovyScriptedOauthTokenHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddGroovyScriptedOauthTokenHandlerRequest) GetSchemasOk() (*[]EnumgroovyScriptedOauthTokenHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddGroovyScriptedOauthTokenHandlerRequest) SetSchemas(v []EnumgroovyScriptedOauthTokenHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetScriptClass

`func (o *AddGroovyScriptedOauthTokenHandlerRequest) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *AddGroovyScriptedOauthTokenHandlerRequest) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *AddGroovyScriptedOauthTokenHandlerRequest) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *AddGroovyScriptedOauthTokenHandlerRequest) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *AddGroovyScriptedOauthTokenHandlerRequest) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *AddGroovyScriptedOauthTokenHandlerRequest) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *AddGroovyScriptedOauthTokenHandlerRequest) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetDescription

`func (o *AddGroovyScriptedOauthTokenHandlerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddGroovyScriptedOauthTokenHandlerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddGroovyScriptedOauthTokenHandlerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddGroovyScriptedOauthTokenHandlerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


