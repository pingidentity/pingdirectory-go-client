# GroovyScriptedPluginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Plugin Root | 
**Schemas** | [**[]EnumgroovyScriptedPluginSchemaUrn**](EnumgroovyScriptedPluginSchemaUrn.md) |  | 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Plugin. | 
**RequestCriteria** | Pointer to **string** | Specifies a set of request criteria that may be used to indicate that this Groovy Scripted Plugin should only be invoked for operations in which the associated request matches this criteria. | [optional] 
**ScriptArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Scripted Plugin. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**PluginType** | [**[]EnumpluginPluginTypeProp**](EnumpluginPluginTypeProp.md) | Specifies the set of plug-in types for the plug-in, which specifies the times at which the plug-in is invoked. | 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewGroovyScriptedPluginResponse

`func NewGroovyScriptedPluginResponse(id string, schemas []EnumgroovyScriptedPluginSchemaUrn, scriptClass string, enabled bool, pluginType []EnumpluginPluginTypeProp, ) *GroovyScriptedPluginResponse`

NewGroovyScriptedPluginResponse instantiates a new GroovyScriptedPluginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroovyScriptedPluginResponseWithDefaults

`func NewGroovyScriptedPluginResponseWithDefaults() *GroovyScriptedPluginResponse`

NewGroovyScriptedPluginResponseWithDefaults instantiates a new GroovyScriptedPluginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroovyScriptedPluginResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroovyScriptedPluginResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroovyScriptedPluginResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *GroovyScriptedPluginResponse) GetSchemas() []EnumgroovyScriptedPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GroovyScriptedPluginResponse) GetSchemasOk() (*[]EnumgroovyScriptedPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GroovyScriptedPluginResponse) SetSchemas(v []EnumgroovyScriptedPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetScriptClass

`func (o *GroovyScriptedPluginResponse) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *GroovyScriptedPluginResponse) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *GroovyScriptedPluginResponse) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetRequestCriteria

`func (o *GroovyScriptedPluginResponse) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *GroovyScriptedPluginResponse) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *GroovyScriptedPluginResponse) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *GroovyScriptedPluginResponse) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetScriptArgument

`func (o *GroovyScriptedPluginResponse) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *GroovyScriptedPluginResponse) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *GroovyScriptedPluginResponse) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *GroovyScriptedPluginResponse) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetDescription

`func (o *GroovyScriptedPluginResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroovyScriptedPluginResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroovyScriptedPluginResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroovyScriptedPluginResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *GroovyScriptedPluginResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GroovyScriptedPluginResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GroovyScriptedPluginResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetPluginType

`func (o *GroovyScriptedPluginResponse) GetPluginType() []EnumpluginPluginTypeProp`

GetPluginType returns the PluginType field if non-nil, zero value otherwise.

### GetPluginTypeOk

`func (o *GroovyScriptedPluginResponse) GetPluginTypeOk() (*[]EnumpluginPluginTypeProp, bool)`

GetPluginTypeOk returns a tuple with the PluginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginType

`func (o *GroovyScriptedPluginResponse) SetPluginType(v []EnumpluginPluginTypeProp)`

SetPluginType sets PluginType field to given value.


### GetInvokeForInternalOperations

`func (o *GroovyScriptedPluginResponse) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *GroovyScriptedPluginResponse) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *GroovyScriptedPluginResponse) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *GroovyScriptedPluginResponse) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.

### GetMeta

`func (o *GroovyScriptedPluginResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GroovyScriptedPluginResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GroovyScriptedPluginResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GroovyScriptedPluginResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *GroovyScriptedPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *GroovyScriptedPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *GroovyScriptedPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *GroovyScriptedPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


