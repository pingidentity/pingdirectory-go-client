# GroovyScriptedChangeSubscriptionHandlerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn**](EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn.md) |  | 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Change Subscription Handler. | 
**ScriptArgument** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Change Subscription Handler | [optional] 
**Enabled** | **bool** | Indicates whether this change subscription handler is enabled within the server. | 
**ChangeSubscription** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGroovyScriptedChangeSubscriptionHandlerShared

`func NewGroovyScriptedChangeSubscriptionHandlerShared(schemas []EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn, scriptClass string, enabled bool, ) *GroovyScriptedChangeSubscriptionHandlerShared`

NewGroovyScriptedChangeSubscriptionHandlerShared instantiates a new GroovyScriptedChangeSubscriptionHandlerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroovyScriptedChangeSubscriptionHandlerSharedWithDefaults

`func NewGroovyScriptedChangeSubscriptionHandlerSharedWithDefaults() *GroovyScriptedChangeSubscriptionHandlerShared`

NewGroovyScriptedChangeSubscriptionHandlerSharedWithDefaults instantiates a new GroovyScriptedChangeSubscriptionHandlerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *GroovyScriptedChangeSubscriptionHandlerShared) GetSchemas() []EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GroovyScriptedChangeSubscriptionHandlerShared) GetSchemasOk() (*[]EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GroovyScriptedChangeSubscriptionHandlerShared) SetSchemas(v []EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetScriptClass

`func (o *GroovyScriptedChangeSubscriptionHandlerShared) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *GroovyScriptedChangeSubscriptionHandlerShared) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *GroovyScriptedChangeSubscriptionHandlerShared) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *GroovyScriptedChangeSubscriptionHandlerShared) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *GroovyScriptedChangeSubscriptionHandlerShared) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *GroovyScriptedChangeSubscriptionHandlerShared) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *GroovyScriptedChangeSubscriptionHandlerShared) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetDescription

`func (o *GroovyScriptedChangeSubscriptionHandlerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroovyScriptedChangeSubscriptionHandlerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroovyScriptedChangeSubscriptionHandlerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroovyScriptedChangeSubscriptionHandlerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *GroovyScriptedChangeSubscriptionHandlerShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GroovyScriptedChangeSubscriptionHandlerShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GroovyScriptedChangeSubscriptionHandlerShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetChangeSubscription

`func (o *GroovyScriptedChangeSubscriptionHandlerShared) GetChangeSubscription() []string`

GetChangeSubscription returns the ChangeSubscription field if non-nil, zero value otherwise.

### GetChangeSubscriptionOk

`func (o *GroovyScriptedChangeSubscriptionHandlerShared) GetChangeSubscriptionOk() (*[]string, bool)`

GetChangeSubscriptionOk returns a tuple with the ChangeSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeSubscription

`func (o *GroovyScriptedChangeSubscriptionHandlerShared) SetChangeSubscription(v []string)`

SetChangeSubscription sets ChangeSubscription field to given value.

### HasChangeSubscription

`func (o *GroovyScriptedChangeSubscriptionHandlerShared) HasChangeSubscription() bool`

HasChangeSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


