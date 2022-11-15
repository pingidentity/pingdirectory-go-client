# AddGroovyScriptedAccountStatusNotificationHandlerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HandlerName** | **string** | Name of the new Account Status Notification Handler | 
**Schemas** | [**[]EnumgroovyScriptedAccountStatusNotificationHandlerSchemaUrn**](EnumgroovyScriptedAccountStatusNotificationHandlerSchemaUrn.md) |  | 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Account Status Notification Handler. | 
**ScriptArgument** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Account Status Notification Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Account Status Notification Handler is enabled. Only enabled handlers are invoked whenever a related event occurs in the server. | 
**Asynchronous** | Pointer to **bool** | Indicates whether the server should attempt to invoke this Account Status Notification Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver a message) will not delay processing for the operation that triggered the notification. | [optional] 
**AccountCreationNotificationRequestCriteria** | Pointer to **string** | A request criteria object that identifies which add requests should result in account creation notifications for this handler. | [optional] 
**AccountUpdateNotificationRequestCriteria** | Pointer to **string** | A request criteria object that identifies which modify and modify DN requests should result in account update notifications for this handler. | [optional] 

## Methods

### NewAddGroovyScriptedAccountStatusNotificationHandlerRequest

`func NewAddGroovyScriptedAccountStatusNotificationHandlerRequest(handlerName string, schemas []EnumgroovyScriptedAccountStatusNotificationHandlerSchemaUrn, scriptClass string, enabled bool, ) *AddGroovyScriptedAccountStatusNotificationHandlerRequest`

NewAddGroovyScriptedAccountStatusNotificationHandlerRequest instantiates a new AddGroovyScriptedAccountStatusNotificationHandlerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddGroovyScriptedAccountStatusNotificationHandlerRequestWithDefaults

`func NewAddGroovyScriptedAccountStatusNotificationHandlerRequestWithDefaults() *AddGroovyScriptedAccountStatusNotificationHandlerRequest`

NewAddGroovyScriptedAccountStatusNotificationHandlerRequestWithDefaults instantiates a new AddGroovyScriptedAccountStatusNotificationHandlerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandlerName

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) GetHandlerName() string`

GetHandlerName returns the HandlerName field if non-nil, zero value otherwise.

### GetHandlerNameOk

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) GetHandlerNameOk() (*string, bool)`

GetHandlerNameOk returns a tuple with the HandlerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerName

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) SetHandlerName(v string)`

SetHandlerName sets HandlerName field to given value.


### GetSchemas

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) GetSchemas() []EnumgroovyScriptedAccountStatusNotificationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) GetSchemasOk() (*[]EnumgroovyScriptedAccountStatusNotificationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) SetSchemas(v []EnumgroovyScriptedAccountStatusNotificationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetScriptClass

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetDescription

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAsynchronous

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) GetAsynchronous() bool`

GetAsynchronous returns the Asynchronous field if non-nil, zero value otherwise.

### GetAsynchronousOk

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) GetAsynchronousOk() (*bool, bool)`

GetAsynchronousOk returns a tuple with the Asynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsynchronous

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) SetAsynchronous(v bool)`

SetAsynchronous sets Asynchronous field to given value.

### HasAsynchronous

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) HasAsynchronous() bool`

HasAsynchronous returns a boolean if a field has been set.

### GetAccountCreationNotificationRequestCriteria

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) GetAccountCreationNotificationRequestCriteria() string`

GetAccountCreationNotificationRequestCriteria returns the AccountCreationNotificationRequestCriteria field if non-nil, zero value otherwise.

### GetAccountCreationNotificationRequestCriteriaOk

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) GetAccountCreationNotificationRequestCriteriaOk() (*string, bool)`

GetAccountCreationNotificationRequestCriteriaOk returns a tuple with the AccountCreationNotificationRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCreationNotificationRequestCriteria

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) SetAccountCreationNotificationRequestCriteria(v string)`

SetAccountCreationNotificationRequestCriteria sets AccountCreationNotificationRequestCriteria field to given value.

### HasAccountCreationNotificationRequestCriteria

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) HasAccountCreationNotificationRequestCriteria() bool`

HasAccountCreationNotificationRequestCriteria returns a boolean if a field has been set.

### GetAccountUpdateNotificationRequestCriteria

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) GetAccountUpdateNotificationRequestCriteria() string`

GetAccountUpdateNotificationRequestCriteria returns the AccountUpdateNotificationRequestCriteria field if non-nil, zero value otherwise.

### GetAccountUpdateNotificationRequestCriteriaOk

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) GetAccountUpdateNotificationRequestCriteriaOk() (*string, bool)`

GetAccountUpdateNotificationRequestCriteriaOk returns a tuple with the AccountUpdateNotificationRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUpdateNotificationRequestCriteria

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) SetAccountUpdateNotificationRequestCriteria(v string)`

SetAccountUpdateNotificationRequestCriteria sets AccountUpdateNotificationRequestCriteria field to given value.

### HasAccountUpdateNotificationRequestCriteria

`func (o *AddGroovyScriptedAccountStatusNotificationHandlerRequest) HasAccountUpdateNotificationRequestCriteria() bool`

HasAccountUpdateNotificationRequestCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


