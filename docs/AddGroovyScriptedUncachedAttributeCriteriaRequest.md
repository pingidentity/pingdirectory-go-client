# AddGroovyScriptedUncachedAttributeCriteriaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CriteriaName** | **string** | Name of the new Uncached Attribute Criteria | 
**Schemas** | [**[]EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn**](EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn.md) |  | 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Uncached Attribute Criteria. | 
**ScriptArgument** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Uncached Attribute Criteria | [optional] 
**Enabled** | **bool** | Indicates whether this Uncached Attribute Criteria is enabled for use in the server. | 

## Methods

### NewAddGroovyScriptedUncachedAttributeCriteriaRequest

`func NewAddGroovyScriptedUncachedAttributeCriteriaRequest(criteriaName string, schemas []EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn, scriptClass string, enabled bool, ) *AddGroovyScriptedUncachedAttributeCriteriaRequest`

NewAddGroovyScriptedUncachedAttributeCriteriaRequest instantiates a new AddGroovyScriptedUncachedAttributeCriteriaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddGroovyScriptedUncachedAttributeCriteriaRequestWithDefaults

`func NewAddGroovyScriptedUncachedAttributeCriteriaRequestWithDefaults() *AddGroovyScriptedUncachedAttributeCriteriaRequest`

NewAddGroovyScriptedUncachedAttributeCriteriaRequestWithDefaults instantiates a new AddGroovyScriptedUncachedAttributeCriteriaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteriaName

`func (o *AddGroovyScriptedUncachedAttributeCriteriaRequest) GetCriteriaName() string`

GetCriteriaName returns the CriteriaName field if non-nil, zero value otherwise.

### GetCriteriaNameOk

`func (o *AddGroovyScriptedUncachedAttributeCriteriaRequest) GetCriteriaNameOk() (*string, bool)`

GetCriteriaNameOk returns a tuple with the CriteriaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteriaName

`func (o *AddGroovyScriptedUncachedAttributeCriteriaRequest) SetCriteriaName(v string)`

SetCriteriaName sets CriteriaName field to given value.


### GetSchemas

`func (o *AddGroovyScriptedUncachedAttributeCriteriaRequest) GetSchemas() []EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddGroovyScriptedUncachedAttributeCriteriaRequest) GetSchemasOk() (*[]EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddGroovyScriptedUncachedAttributeCriteriaRequest) SetSchemas(v []EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetScriptClass

`func (o *AddGroovyScriptedUncachedAttributeCriteriaRequest) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *AddGroovyScriptedUncachedAttributeCriteriaRequest) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *AddGroovyScriptedUncachedAttributeCriteriaRequest) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *AddGroovyScriptedUncachedAttributeCriteriaRequest) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *AddGroovyScriptedUncachedAttributeCriteriaRequest) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *AddGroovyScriptedUncachedAttributeCriteriaRequest) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *AddGroovyScriptedUncachedAttributeCriteriaRequest) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetDescription

`func (o *AddGroovyScriptedUncachedAttributeCriteriaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddGroovyScriptedUncachedAttributeCriteriaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddGroovyScriptedUncachedAttributeCriteriaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddGroovyScriptedUncachedAttributeCriteriaRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddGroovyScriptedUncachedAttributeCriteriaRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddGroovyScriptedUncachedAttributeCriteriaRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddGroovyScriptedUncachedAttributeCriteriaRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


