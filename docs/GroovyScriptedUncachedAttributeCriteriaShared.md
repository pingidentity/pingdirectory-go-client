# GroovyScriptedUncachedAttributeCriteriaShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn**](EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn.md) |  | 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Uncached Attribute Criteria. | 
**ScriptArgument** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Uncached Attribute Criteria | [optional] 
**Enabled** | **bool** | Indicates whether this Uncached Attribute Criteria is enabled for use in the server. | 

## Methods

### NewGroovyScriptedUncachedAttributeCriteriaShared

`func NewGroovyScriptedUncachedAttributeCriteriaShared(schemas []EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn, scriptClass string, enabled bool, ) *GroovyScriptedUncachedAttributeCriteriaShared`

NewGroovyScriptedUncachedAttributeCriteriaShared instantiates a new GroovyScriptedUncachedAttributeCriteriaShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroovyScriptedUncachedAttributeCriteriaSharedWithDefaults

`func NewGroovyScriptedUncachedAttributeCriteriaSharedWithDefaults() *GroovyScriptedUncachedAttributeCriteriaShared`

NewGroovyScriptedUncachedAttributeCriteriaSharedWithDefaults instantiates a new GroovyScriptedUncachedAttributeCriteriaShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *GroovyScriptedUncachedAttributeCriteriaShared) GetSchemas() []EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GroovyScriptedUncachedAttributeCriteriaShared) GetSchemasOk() (*[]EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GroovyScriptedUncachedAttributeCriteriaShared) SetSchemas(v []EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetScriptClass

`func (o *GroovyScriptedUncachedAttributeCriteriaShared) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *GroovyScriptedUncachedAttributeCriteriaShared) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *GroovyScriptedUncachedAttributeCriteriaShared) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *GroovyScriptedUncachedAttributeCriteriaShared) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *GroovyScriptedUncachedAttributeCriteriaShared) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *GroovyScriptedUncachedAttributeCriteriaShared) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *GroovyScriptedUncachedAttributeCriteriaShared) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetDescription

`func (o *GroovyScriptedUncachedAttributeCriteriaShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroovyScriptedUncachedAttributeCriteriaShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroovyScriptedUncachedAttributeCriteriaShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroovyScriptedUncachedAttributeCriteriaShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *GroovyScriptedUncachedAttributeCriteriaShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GroovyScriptedUncachedAttributeCriteriaShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GroovyScriptedUncachedAttributeCriteriaShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


