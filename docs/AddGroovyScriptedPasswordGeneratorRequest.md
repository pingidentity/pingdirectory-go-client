# AddGroovyScriptedPasswordGeneratorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneratorName** | **string** | Name of the new Password Generator | 
**Schemas** | [**[]EnumgroovyScriptedPasswordGeneratorSchemaUrn**](EnumgroovyScriptedPasswordGeneratorSchemaUrn.md) |  | 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Password Generator. | 
**ScriptArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Scripted Password Generator. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**Description** | Pointer to **string** | A description for this Password Generator | [optional] 
**Enabled** | **bool** | Indicates whether the Password Generator is enabled for use. | 

## Methods

### NewAddGroovyScriptedPasswordGeneratorRequest

`func NewAddGroovyScriptedPasswordGeneratorRequest(generatorName string, schemas []EnumgroovyScriptedPasswordGeneratorSchemaUrn, scriptClass string, enabled bool, ) *AddGroovyScriptedPasswordGeneratorRequest`

NewAddGroovyScriptedPasswordGeneratorRequest instantiates a new AddGroovyScriptedPasswordGeneratorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddGroovyScriptedPasswordGeneratorRequestWithDefaults

`func NewAddGroovyScriptedPasswordGeneratorRequestWithDefaults() *AddGroovyScriptedPasswordGeneratorRequest`

NewAddGroovyScriptedPasswordGeneratorRequestWithDefaults instantiates a new AddGroovyScriptedPasswordGeneratorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneratorName

`func (o *AddGroovyScriptedPasswordGeneratorRequest) GetGeneratorName() string`

GetGeneratorName returns the GeneratorName field if non-nil, zero value otherwise.

### GetGeneratorNameOk

`func (o *AddGroovyScriptedPasswordGeneratorRequest) GetGeneratorNameOk() (*string, bool)`

GetGeneratorNameOk returns a tuple with the GeneratorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratorName

`func (o *AddGroovyScriptedPasswordGeneratorRequest) SetGeneratorName(v string)`

SetGeneratorName sets GeneratorName field to given value.


### GetSchemas

`func (o *AddGroovyScriptedPasswordGeneratorRequest) GetSchemas() []EnumgroovyScriptedPasswordGeneratorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddGroovyScriptedPasswordGeneratorRequest) GetSchemasOk() (*[]EnumgroovyScriptedPasswordGeneratorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddGroovyScriptedPasswordGeneratorRequest) SetSchemas(v []EnumgroovyScriptedPasswordGeneratorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetScriptClass

`func (o *AddGroovyScriptedPasswordGeneratorRequest) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *AddGroovyScriptedPasswordGeneratorRequest) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *AddGroovyScriptedPasswordGeneratorRequest) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *AddGroovyScriptedPasswordGeneratorRequest) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *AddGroovyScriptedPasswordGeneratorRequest) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *AddGroovyScriptedPasswordGeneratorRequest) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *AddGroovyScriptedPasswordGeneratorRequest) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetDescription

`func (o *AddGroovyScriptedPasswordGeneratorRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddGroovyScriptedPasswordGeneratorRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddGroovyScriptedPasswordGeneratorRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddGroovyScriptedPasswordGeneratorRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddGroovyScriptedPasswordGeneratorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddGroovyScriptedPasswordGeneratorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddGroovyScriptedPasswordGeneratorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


