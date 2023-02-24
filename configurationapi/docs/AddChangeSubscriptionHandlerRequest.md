# AddChangeSubscriptionHandlerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HandlerName** | **string** | Name of the new Change Subscription Handler | 
**Schemas** | [**[]EnumthirdPartyChangeSubscriptionHandlerSchemaUrn**](EnumthirdPartyChangeSubscriptionHandlerSchemaUrn.md) |  | 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Change Subscription Handler. | 
**ScriptArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Scripted Change Subscription Handler. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**Description** | Pointer to **string** | A description for this Change Subscription Handler | [optional] 
**Enabled** | **bool** | Indicates whether this change subscription handler is enabled within the server. | 
**ChangeSubscription** | Pointer to **[]string** | The set of change subscriptions for which this change subscription handler should be notified. If no values are provided then it will be notified for all change subscriptions defined in the server. | [optional] 
**LogFile** | Pointer to **string** | Specifies the log file in which the change notification messages will be written. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Change Subscription Handler. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Change Subscription Handler. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddChangeSubscriptionHandlerRequest

`func NewAddChangeSubscriptionHandlerRequest(handlerName string, schemas []EnumthirdPartyChangeSubscriptionHandlerSchemaUrn, scriptClass string, enabled bool, extensionClass string, ) *AddChangeSubscriptionHandlerRequest`

NewAddChangeSubscriptionHandlerRequest instantiates a new AddChangeSubscriptionHandlerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddChangeSubscriptionHandlerRequestWithDefaults

`func NewAddChangeSubscriptionHandlerRequestWithDefaults() *AddChangeSubscriptionHandlerRequest`

NewAddChangeSubscriptionHandlerRequestWithDefaults instantiates a new AddChangeSubscriptionHandlerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandlerName

`func (o *AddChangeSubscriptionHandlerRequest) GetHandlerName() string`

GetHandlerName returns the HandlerName field if non-nil, zero value otherwise.

### GetHandlerNameOk

`func (o *AddChangeSubscriptionHandlerRequest) GetHandlerNameOk() (*string, bool)`

GetHandlerNameOk returns a tuple with the HandlerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerName

`func (o *AddChangeSubscriptionHandlerRequest) SetHandlerName(v string)`

SetHandlerName sets HandlerName field to given value.


### GetSchemas

`func (o *AddChangeSubscriptionHandlerRequest) GetSchemas() []EnumthirdPartyChangeSubscriptionHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddChangeSubscriptionHandlerRequest) GetSchemasOk() (*[]EnumthirdPartyChangeSubscriptionHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddChangeSubscriptionHandlerRequest) SetSchemas(v []EnumthirdPartyChangeSubscriptionHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetScriptClass

`func (o *AddChangeSubscriptionHandlerRequest) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *AddChangeSubscriptionHandlerRequest) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *AddChangeSubscriptionHandlerRequest) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *AddChangeSubscriptionHandlerRequest) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *AddChangeSubscriptionHandlerRequest) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *AddChangeSubscriptionHandlerRequest) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *AddChangeSubscriptionHandlerRequest) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetDescription

`func (o *AddChangeSubscriptionHandlerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddChangeSubscriptionHandlerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddChangeSubscriptionHandlerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddChangeSubscriptionHandlerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddChangeSubscriptionHandlerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddChangeSubscriptionHandlerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddChangeSubscriptionHandlerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetChangeSubscription

`func (o *AddChangeSubscriptionHandlerRequest) GetChangeSubscription() []string`

GetChangeSubscription returns the ChangeSubscription field if non-nil, zero value otherwise.

### GetChangeSubscriptionOk

`func (o *AddChangeSubscriptionHandlerRequest) GetChangeSubscriptionOk() (*[]string, bool)`

GetChangeSubscriptionOk returns a tuple with the ChangeSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeSubscription

`func (o *AddChangeSubscriptionHandlerRequest) SetChangeSubscription(v []string)`

SetChangeSubscription sets ChangeSubscription field to given value.

### HasChangeSubscription

`func (o *AddChangeSubscriptionHandlerRequest) HasChangeSubscription() bool`

HasChangeSubscription returns a boolean if a field has been set.

### GetLogFile

`func (o *AddChangeSubscriptionHandlerRequest) GetLogFile() string`

GetLogFile returns the LogFile field if non-nil, zero value otherwise.

### GetLogFileOk

`func (o *AddChangeSubscriptionHandlerRequest) GetLogFileOk() (*string, bool)`

GetLogFileOk returns a tuple with the LogFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFile

`func (o *AddChangeSubscriptionHandlerRequest) SetLogFile(v string)`

SetLogFile sets LogFile field to given value.

### HasLogFile

`func (o *AddChangeSubscriptionHandlerRequest) HasLogFile() bool`

HasLogFile returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddChangeSubscriptionHandlerRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddChangeSubscriptionHandlerRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddChangeSubscriptionHandlerRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddChangeSubscriptionHandlerRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddChangeSubscriptionHandlerRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddChangeSubscriptionHandlerRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddChangeSubscriptionHandlerRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


