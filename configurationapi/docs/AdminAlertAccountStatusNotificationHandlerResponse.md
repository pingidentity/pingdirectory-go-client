# AdminAlertAccountStatusNotificationHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumadminAlertAccountStatusNotificationHandlerSchemaUrn**](EnumadminAlertAccountStatusNotificationHandlerSchemaUrn.md) |  | 
**AccountStatusNotificationType** | [**[]EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp**](EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp.md) |  | 
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

### NewAdminAlertAccountStatusNotificationHandlerResponse

`func NewAdminAlertAccountStatusNotificationHandlerResponse(schemas []EnumadminAlertAccountStatusNotificationHandlerSchemaUrn, accountStatusNotificationType []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp, enabled bool, id string, ) *AdminAlertAccountStatusNotificationHandlerResponse`

NewAdminAlertAccountStatusNotificationHandlerResponse instantiates a new AdminAlertAccountStatusNotificationHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminAlertAccountStatusNotificationHandlerResponseWithDefaults

`func NewAdminAlertAccountStatusNotificationHandlerResponseWithDefaults() *AdminAlertAccountStatusNotificationHandlerResponse`

NewAdminAlertAccountStatusNotificationHandlerResponseWithDefaults instantiates a new AdminAlertAccountStatusNotificationHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetSchemas() []EnumadminAlertAccountStatusNotificationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetSchemasOk() (*[]EnumadminAlertAccountStatusNotificationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) SetSchemas(v []EnumadminAlertAccountStatusNotificationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAccountStatusNotificationType

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetAccountStatusNotificationType() []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp`

GetAccountStatusNotificationType returns the AccountStatusNotificationType field if non-nil, zero value otherwise.

### GetAccountStatusNotificationTypeOk

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetAccountStatusNotificationTypeOk() (*[]EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp, bool)`

GetAccountStatusNotificationTypeOk returns a tuple with the AccountStatusNotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatusNotificationType

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) SetAccountStatusNotificationType(v []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp)`

SetAccountStatusNotificationType sets AccountStatusNotificationType field to given value.


### GetDescription

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAsynchronous

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetAsynchronous() bool`

GetAsynchronous returns the Asynchronous field if non-nil, zero value otherwise.

### GetAsynchronousOk

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetAsynchronousOk() (*bool, bool)`

GetAsynchronousOk returns a tuple with the Asynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsynchronous

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) SetAsynchronous(v bool)`

SetAsynchronous sets Asynchronous field to given value.

### HasAsynchronous

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) HasAsynchronous() bool`

HasAsynchronous returns a boolean if a field has been set.

### GetAccountAuthenticationNotificationResultCriteria

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetAccountAuthenticationNotificationResultCriteria() string`

GetAccountAuthenticationNotificationResultCriteria returns the AccountAuthenticationNotificationResultCriteria field if non-nil, zero value otherwise.

### GetAccountAuthenticationNotificationResultCriteriaOk

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetAccountAuthenticationNotificationResultCriteriaOk() (*string, bool)`

GetAccountAuthenticationNotificationResultCriteriaOk returns a tuple with the AccountAuthenticationNotificationResultCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAuthenticationNotificationResultCriteria

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) SetAccountAuthenticationNotificationResultCriteria(v string)`

SetAccountAuthenticationNotificationResultCriteria sets AccountAuthenticationNotificationResultCriteria field to given value.

### HasAccountAuthenticationNotificationResultCriteria

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) HasAccountAuthenticationNotificationResultCriteria() bool`

HasAccountAuthenticationNotificationResultCriteria returns a boolean if a field has been set.

### GetAccountCreationNotificationRequestCriteria

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetAccountCreationNotificationRequestCriteria() string`

GetAccountCreationNotificationRequestCriteria returns the AccountCreationNotificationRequestCriteria field if non-nil, zero value otherwise.

### GetAccountCreationNotificationRequestCriteriaOk

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetAccountCreationNotificationRequestCriteriaOk() (*string, bool)`

GetAccountCreationNotificationRequestCriteriaOk returns a tuple with the AccountCreationNotificationRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCreationNotificationRequestCriteria

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) SetAccountCreationNotificationRequestCriteria(v string)`

SetAccountCreationNotificationRequestCriteria sets AccountCreationNotificationRequestCriteria field to given value.

### HasAccountCreationNotificationRequestCriteria

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) HasAccountCreationNotificationRequestCriteria() bool`

HasAccountCreationNotificationRequestCriteria returns a boolean if a field has been set.

### GetAccountDeletionNotificationRequestCriteria

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetAccountDeletionNotificationRequestCriteria() string`

GetAccountDeletionNotificationRequestCriteria returns the AccountDeletionNotificationRequestCriteria field if non-nil, zero value otherwise.

### GetAccountDeletionNotificationRequestCriteriaOk

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetAccountDeletionNotificationRequestCriteriaOk() (*string, bool)`

GetAccountDeletionNotificationRequestCriteriaOk returns a tuple with the AccountDeletionNotificationRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDeletionNotificationRequestCriteria

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) SetAccountDeletionNotificationRequestCriteria(v string)`

SetAccountDeletionNotificationRequestCriteria sets AccountDeletionNotificationRequestCriteria field to given value.

### HasAccountDeletionNotificationRequestCriteria

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) HasAccountDeletionNotificationRequestCriteria() bool`

HasAccountDeletionNotificationRequestCriteria returns a boolean if a field has been set.

### GetAccountUpdateNotificationRequestCriteria

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetAccountUpdateNotificationRequestCriteria() string`

GetAccountUpdateNotificationRequestCriteria returns the AccountUpdateNotificationRequestCriteria field if non-nil, zero value otherwise.

### GetAccountUpdateNotificationRequestCriteriaOk

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetAccountUpdateNotificationRequestCriteriaOk() (*string, bool)`

GetAccountUpdateNotificationRequestCriteriaOk returns a tuple with the AccountUpdateNotificationRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUpdateNotificationRequestCriteria

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) SetAccountUpdateNotificationRequestCriteria(v string)`

SetAccountUpdateNotificationRequestCriteria sets AccountUpdateNotificationRequestCriteria field to given value.

### HasAccountUpdateNotificationRequestCriteria

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) HasAccountUpdateNotificationRequestCriteria() bool`

HasAccountUpdateNotificationRequestCriteria returns a boolean if a field has been set.

### GetMeta

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdminAlertAccountStatusNotificationHandlerResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


