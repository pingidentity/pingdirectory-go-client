# GroovyScriptedAccountStatusNotificationHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumgroovyScriptedAccountStatusNotificationHandlerSchemaUrn**](EnumgroovyScriptedAccountStatusNotificationHandlerSchemaUrn.md) |  | 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Account Status Notification Handler. | 
**ScriptArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Scripted Account Status Notification Handler. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**Description** | Pointer to **string** | A description for this Account Status Notification Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Account Status Notification Handler is enabled. Only enabled handlers are invoked whenever a related event occurs in the server. | 
**Asynchronous** | Pointer to **bool** | Indicates whether the server should attempt to invoke this Account Status Notification Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver a message) will not delay processing for the operation that triggered the notification. | [optional] 
**AccountAuthenticationNotificationResultCriteria** | Pointer to **string** | A result criteria object that identifies which successful bind operations should result in account authentication notifications for this handler. | [optional] 
**AccountCreationNotificationRequestCriteria** | Pointer to **string** | A request criteria object that identifies which add requests should result in account creation notifications for this handler. | [optional] 
**AccountDeletionNotificationRequestCriteria** | Pointer to **string** | A request criteria object that identifies which delete requests should result in account deletion notifications for this handler. | [optional] 
**AccountUpdateNotificationRequestCriteria** | Pointer to **string** | A request criteria object that identifies which modify and modify DN requests should result in account update notifications for this handler. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Account Status Notification Handler | 

## Methods

### NewGroovyScriptedAccountStatusNotificationHandlerResponse

`func NewGroovyScriptedAccountStatusNotificationHandlerResponse(schemas []EnumgroovyScriptedAccountStatusNotificationHandlerSchemaUrn, scriptClass string, enabled bool, id string, ) *GroovyScriptedAccountStatusNotificationHandlerResponse`

NewGroovyScriptedAccountStatusNotificationHandlerResponse instantiates a new GroovyScriptedAccountStatusNotificationHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroovyScriptedAccountStatusNotificationHandlerResponseWithDefaults

`func NewGroovyScriptedAccountStatusNotificationHandlerResponseWithDefaults() *GroovyScriptedAccountStatusNotificationHandlerResponse`

NewGroovyScriptedAccountStatusNotificationHandlerResponseWithDefaults instantiates a new GroovyScriptedAccountStatusNotificationHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) GetSchemas() []EnumgroovyScriptedAccountStatusNotificationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) GetSchemasOk() (*[]EnumgroovyScriptedAccountStatusNotificationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) SetSchemas(v []EnumgroovyScriptedAccountStatusNotificationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetScriptClass

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetDescription

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAsynchronous

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) GetAsynchronous() bool`

GetAsynchronous returns the Asynchronous field if non-nil, zero value otherwise.

### GetAsynchronousOk

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) GetAsynchronousOk() (*bool, bool)`

GetAsynchronousOk returns a tuple with the Asynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsynchronous

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) SetAsynchronous(v bool)`

SetAsynchronous sets Asynchronous field to given value.

### HasAsynchronous

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) HasAsynchronous() bool`

HasAsynchronous returns a boolean if a field has been set.

### GetAccountAuthenticationNotificationResultCriteria

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) GetAccountAuthenticationNotificationResultCriteria() string`

GetAccountAuthenticationNotificationResultCriteria returns the AccountAuthenticationNotificationResultCriteria field if non-nil, zero value otherwise.

### GetAccountAuthenticationNotificationResultCriteriaOk

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) GetAccountAuthenticationNotificationResultCriteriaOk() (*string, bool)`

GetAccountAuthenticationNotificationResultCriteriaOk returns a tuple with the AccountAuthenticationNotificationResultCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAuthenticationNotificationResultCriteria

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) SetAccountAuthenticationNotificationResultCriteria(v string)`

SetAccountAuthenticationNotificationResultCriteria sets AccountAuthenticationNotificationResultCriteria field to given value.

### HasAccountAuthenticationNotificationResultCriteria

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) HasAccountAuthenticationNotificationResultCriteria() bool`

HasAccountAuthenticationNotificationResultCriteria returns a boolean if a field has been set.

### GetAccountCreationNotificationRequestCriteria

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) GetAccountCreationNotificationRequestCriteria() string`

GetAccountCreationNotificationRequestCriteria returns the AccountCreationNotificationRequestCriteria field if non-nil, zero value otherwise.

### GetAccountCreationNotificationRequestCriteriaOk

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) GetAccountCreationNotificationRequestCriteriaOk() (*string, bool)`

GetAccountCreationNotificationRequestCriteriaOk returns a tuple with the AccountCreationNotificationRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCreationNotificationRequestCriteria

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) SetAccountCreationNotificationRequestCriteria(v string)`

SetAccountCreationNotificationRequestCriteria sets AccountCreationNotificationRequestCriteria field to given value.

### HasAccountCreationNotificationRequestCriteria

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) HasAccountCreationNotificationRequestCriteria() bool`

HasAccountCreationNotificationRequestCriteria returns a boolean if a field has been set.

### GetAccountDeletionNotificationRequestCriteria

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) GetAccountDeletionNotificationRequestCriteria() string`

GetAccountDeletionNotificationRequestCriteria returns the AccountDeletionNotificationRequestCriteria field if non-nil, zero value otherwise.

### GetAccountDeletionNotificationRequestCriteriaOk

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) GetAccountDeletionNotificationRequestCriteriaOk() (*string, bool)`

GetAccountDeletionNotificationRequestCriteriaOk returns a tuple with the AccountDeletionNotificationRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDeletionNotificationRequestCriteria

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) SetAccountDeletionNotificationRequestCriteria(v string)`

SetAccountDeletionNotificationRequestCriteria sets AccountDeletionNotificationRequestCriteria field to given value.

### HasAccountDeletionNotificationRequestCriteria

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) HasAccountDeletionNotificationRequestCriteria() bool`

HasAccountDeletionNotificationRequestCriteria returns a boolean if a field has been set.

### GetAccountUpdateNotificationRequestCriteria

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) GetAccountUpdateNotificationRequestCriteria() string`

GetAccountUpdateNotificationRequestCriteria returns the AccountUpdateNotificationRequestCriteria field if non-nil, zero value otherwise.

### GetAccountUpdateNotificationRequestCriteriaOk

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) GetAccountUpdateNotificationRequestCriteriaOk() (*string, bool)`

GetAccountUpdateNotificationRequestCriteriaOk returns a tuple with the AccountUpdateNotificationRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUpdateNotificationRequestCriteria

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) SetAccountUpdateNotificationRequestCriteria(v string)`

SetAccountUpdateNotificationRequestCriteria sets AccountUpdateNotificationRequestCriteria field to given value.

### HasAccountUpdateNotificationRequestCriteria

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) HasAccountUpdateNotificationRequestCriteria() bool`

HasAccountUpdateNotificationRequestCriteria returns a boolean if a field has been set.

### GetMeta

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroovyScriptedAccountStatusNotificationHandlerResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


