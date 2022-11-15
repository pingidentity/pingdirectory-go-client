# GroovyScriptedPasswordGeneratorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Password Generator | 
**Schemas** | [**[]EnumgroovyScriptedPasswordGeneratorSchemaUrn**](EnumgroovyScriptedPasswordGeneratorSchemaUrn.md) |  | 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Password Generator. | 
**ScriptArgument** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Password Generator | [optional] 
**Enabled** | **bool** | Indicates whether the Password Generator is enabled for use. | 

## Methods

### NewGroovyScriptedPasswordGeneratorResponse

`func NewGroovyScriptedPasswordGeneratorResponse(id string, schemas []EnumgroovyScriptedPasswordGeneratorSchemaUrn, scriptClass string, enabled bool, ) *GroovyScriptedPasswordGeneratorResponse`

NewGroovyScriptedPasswordGeneratorResponse instantiates a new GroovyScriptedPasswordGeneratorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroovyScriptedPasswordGeneratorResponseWithDefaults

`func NewGroovyScriptedPasswordGeneratorResponseWithDefaults() *GroovyScriptedPasswordGeneratorResponse`

NewGroovyScriptedPasswordGeneratorResponseWithDefaults instantiates a new GroovyScriptedPasswordGeneratorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroovyScriptedPasswordGeneratorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroovyScriptedPasswordGeneratorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroovyScriptedPasswordGeneratorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *GroovyScriptedPasswordGeneratorResponse) GetSchemas() []EnumgroovyScriptedPasswordGeneratorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GroovyScriptedPasswordGeneratorResponse) GetSchemasOk() (*[]EnumgroovyScriptedPasswordGeneratorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GroovyScriptedPasswordGeneratorResponse) SetSchemas(v []EnumgroovyScriptedPasswordGeneratorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetScriptClass

`func (o *GroovyScriptedPasswordGeneratorResponse) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *GroovyScriptedPasswordGeneratorResponse) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *GroovyScriptedPasswordGeneratorResponse) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *GroovyScriptedPasswordGeneratorResponse) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *GroovyScriptedPasswordGeneratorResponse) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *GroovyScriptedPasswordGeneratorResponse) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *GroovyScriptedPasswordGeneratorResponse) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetDescription

`func (o *GroovyScriptedPasswordGeneratorResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroovyScriptedPasswordGeneratorResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroovyScriptedPasswordGeneratorResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroovyScriptedPasswordGeneratorResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *GroovyScriptedPasswordGeneratorResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GroovyScriptedPasswordGeneratorResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GroovyScriptedPasswordGeneratorResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


