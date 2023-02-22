# AddChangeSubscriptionHandler200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Change Subscription Handler | 
**Schemas** | [**[]EnumthirdPartyChangeSubscriptionHandlerSchemaUrn**](EnumthirdPartyChangeSubscriptionHandlerSchemaUrn.md) |  | 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Change Subscription Handler. | 
**ScriptArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Scripted Change Subscription Handler. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**Description** | Pointer to **string** | A description for this Change Subscription Handler | [optional] 
**Enabled** | **bool** | Indicates whether this change subscription handler is enabled within the server. | 
**ChangeSubscription** | Pointer to **[]string** | The set of change subscriptions for which this change subscription handler should be notified. If no values are provided then it will be notified for all change subscriptions defined in the server. | [optional] 
**LogFile** | **string** | Specifies the log file in which the change notification messages will be written. | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Change Subscription Handler. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Change Subscription Handler. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddChangeSubscriptionHandler200Response

`func NewAddChangeSubscriptionHandler200Response(id string, schemas []EnumthirdPartyChangeSubscriptionHandlerSchemaUrn, scriptClass string, enabled bool, logFile string, extensionClass string, ) *AddChangeSubscriptionHandler200Response`

NewAddChangeSubscriptionHandler200Response instantiates a new AddChangeSubscriptionHandler200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddChangeSubscriptionHandler200ResponseWithDefaults

`func NewAddChangeSubscriptionHandler200ResponseWithDefaults() *AddChangeSubscriptionHandler200Response`

NewAddChangeSubscriptionHandler200ResponseWithDefaults instantiates a new AddChangeSubscriptionHandler200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *AddChangeSubscriptionHandler200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddChangeSubscriptionHandler200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddChangeSubscriptionHandler200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddChangeSubscriptionHandler200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AddChangeSubscriptionHandler200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AddChangeSubscriptionHandler200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AddChangeSubscriptionHandler200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AddChangeSubscriptionHandler200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *AddChangeSubscriptionHandler200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddChangeSubscriptionHandler200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddChangeSubscriptionHandler200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddChangeSubscriptionHandler200Response) GetSchemas() []EnumthirdPartyChangeSubscriptionHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddChangeSubscriptionHandler200Response) GetSchemasOk() (*[]EnumthirdPartyChangeSubscriptionHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddChangeSubscriptionHandler200Response) SetSchemas(v []EnumthirdPartyChangeSubscriptionHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetScriptClass

`func (o *AddChangeSubscriptionHandler200Response) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *AddChangeSubscriptionHandler200Response) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *AddChangeSubscriptionHandler200Response) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *AddChangeSubscriptionHandler200Response) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *AddChangeSubscriptionHandler200Response) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *AddChangeSubscriptionHandler200Response) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *AddChangeSubscriptionHandler200Response) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetDescription

`func (o *AddChangeSubscriptionHandler200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddChangeSubscriptionHandler200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddChangeSubscriptionHandler200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddChangeSubscriptionHandler200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddChangeSubscriptionHandler200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddChangeSubscriptionHandler200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddChangeSubscriptionHandler200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetChangeSubscription

`func (o *AddChangeSubscriptionHandler200Response) GetChangeSubscription() []string`

GetChangeSubscription returns the ChangeSubscription field if non-nil, zero value otherwise.

### GetChangeSubscriptionOk

`func (o *AddChangeSubscriptionHandler200Response) GetChangeSubscriptionOk() (*[]string, bool)`

GetChangeSubscriptionOk returns a tuple with the ChangeSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeSubscription

`func (o *AddChangeSubscriptionHandler200Response) SetChangeSubscription(v []string)`

SetChangeSubscription sets ChangeSubscription field to given value.

### HasChangeSubscription

`func (o *AddChangeSubscriptionHandler200Response) HasChangeSubscription() bool`

HasChangeSubscription returns a boolean if a field has been set.

### GetLogFile

`func (o *AddChangeSubscriptionHandler200Response) GetLogFile() string`

GetLogFile returns the LogFile field if non-nil, zero value otherwise.

### GetLogFileOk

`func (o *AddChangeSubscriptionHandler200Response) GetLogFileOk() (*string, bool)`

GetLogFileOk returns a tuple with the LogFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFile

`func (o *AddChangeSubscriptionHandler200Response) SetLogFile(v string)`

SetLogFile sets LogFile field to given value.


### GetExtensionClass

`func (o *AddChangeSubscriptionHandler200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddChangeSubscriptionHandler200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddChangeSubscriptionHandler200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddChangeSubscriptionHandler200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddChangeSubscriptionHandler200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddChangeSubscriptionHandler200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddChangeSubscriptionHandler200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


