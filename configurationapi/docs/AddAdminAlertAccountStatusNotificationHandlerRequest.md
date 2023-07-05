# AddAdminAlertAccountStatusNotificationHandlerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HandlerName** | **string** | Name of the new Account Status Notification Handler | 
**Schemas** | [**[]EnumadminAlertAccountStatusNotificationHandlerSchemaUrn**](EnumadminAlertAccountStatusNotificationHandlerSchemaUrn.md) |  | 
**AccountStatusNotificationType** | [**[]EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp**](EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp.md) |  | 
**Description** | Pointer to **string** | A description for this Account Status Notification Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Account Status Notification Handler is enabled. Only enabled handlers are invoked whenever a related event occurs in the server. | 
**Asynchronous** | Pointer to **bool** | Indicates whether the server should attempt to invoke this Account Status Notification Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver a message) will not delay processing for the operation that triggered the notification. | [optional] 
**AccountAuthenticationNotificationResultCriteria** | Pointer to **string** | A result criteria object that identifies which successful bind operations should result in account authentication notifications for this handler. | [optional] 
**AccountCreationNotificationRequestCriteria** | Pointer to **string** | A request criteria object that identifies which add requests should result in account creation notifications for this handler. | [optional] 
**AccountDeletionNotificationRequestCriteria** | Pointer to **string** | A request criteria object that identifies which delete requests should result in account deletion notifications for this handler. | [optional] 
**AccountUpdateNotificationRequestCriteria** | Pointer to **string** | A request criteria object that identifies which modify and modify DN requests should result in account update notifications for this handler. | [optional] 

## Methods

### NewAddAdminAlertAccountStatusNotificationHandlerRequest

`func NewAddAdminAlertAccountStatusNotificationHandlerRequest(handlerName string, schemas []EnumadminAlertAccountStatusNotificationHandlerSchemaUrn, accountStatusNotificationType []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp, enabled bool, ) *AddAdminAlertAccountStatusNotificationHandlerRequest`

NewAddAdminAlertAccountStatusNotificationHandlerRequest instantiates a new AddAdminAlertAccountStatusNotificationHandlerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAdminAlertAccountStatusNotificationHandlerRequestWithDefaults

`func NewAddAdminAlertAccountStatusNotificationHandlerRequestWithDefaults() *AddAdminAlertAccountStatusNotificationHandlerRequest`

NewAddAdminAlertAccountStatusNotificationHandlerRequestWithDefaults instantiates a new AddAdminAlertAccountStatusNotificationHandlerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandlerName

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetHandlerName() string`

GetHandlerName returns the HandlerName field if non-nil, zero value otherwise.

### GetHandlerNameOk

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetHandlerNameOk() (*string, bool)`

GetHandlerNameOk returns a tuple with the HandlerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerName

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) SetHandlerName(v string)`

SetHandlerName sets HandlerName field to given value.


### GetSchemas

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetSchemas() []EnumadminAlertAccountStatusNotificationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetSchemasOk() (*[]EnumadminAlertAccountStatusNotificationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) SetSchemas(v []EnumadminAlertAccountStatusNotificationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAccountStatusNotificationType

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetAccountStatusNotificationType() []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp`

GetAccountStatusNotificationType returns the AccountStatusNotificationType field if non-nil, zero value otherwise.

### GetAccountStatusNotificationTypeOk

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetAccountStatusNotificationTypeOk() (*[]EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp, bool)`

GetAccountStatusNotificationTypeOk returns a tuple with the AccountStatusNotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatusNotificationType

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) SetAccountStatusNotificationType(v []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp)`

SetAccountStatusNotificationType sets AccountStatusNotificationType field to given value.


### GetDescription

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAsynchronous

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetAsynchronous() bool`

GetAsynchronous returns the Asynchronous field if non-nil, zero value otherwise.

### GetAsynchronousOk

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetAsynchronousOk() (*bool, bool)`

GetAsynchronousOk returns a tuple with the Asynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsynchronous

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) SetAsynchronous(v bool)`

SetAsynchronous sets Asynchronous field to given value.

### HasAsynchronous

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) HasAsynchronous() bool`

HasAsynchronous returns a boolean if a field has been set.

### GetAccountAuthenticationNotificationResultCriteria

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetAccountAuthenticationNotificationResultCriteria() string`

GetAccountAuthenticationNotificationResultCriteria returns the AccountAuthenticationNotificationResultCriteria field if non-nil, zero value otherwise.

### GetAccountAuthenticationNotificationResultCriteriaOk

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetAccountAuthenticationNotificationResultCriteriaOk() (*string, bool)`

GetAccountAuthenticationNotificationResultCriteriaOk returns a tuple with the AccountAuthenticationNotificationResultCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAuthenticationNotificationResultCriteria

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) SetAccountAuthenticationNotificationResultCriteria(v string)`

SetAccountAuthenticationNotificationResultCriteria sets AccountAuthenticationNotificationResultCriteria field to given value.

### HasAccountAuthenticationNotificationResultCriteria

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) HasAccountAuthenticationNotificationResultCriteria() bool`

HasAccountAuthenticationNotificationResultCriteria returns a boolean if a field has been set.

### GetAccountCreationNotificationRequestCriteria

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetAccountCreationNotificationRequestCriteria() string`

GetAccountCreationNotificationRequestCriteria returns the AccountCreationNotificationRequestCriteria field if non-nil, zero value otherwise.

### GetAccountCreationNotificationRequestCriteriaOk

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetAccountCreationNotificationRequestCriteriaOk() (*string, bool)`

GetAccountCreationNotificationRequestCriteriaOk returns a tuple with the AccountCreationNotificationRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCreationNotificationRequestCriteria

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) SetAccountCreationNotificationRequestCriteria(v string)`

SetAccountCreationNotificationRequestCriteria sets AccountCreationNotificationRequestCriteria field to given value.

### HasAccountCreationNotificationRequestCriteria

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) HasAccountCreationNotificationRequestCriteria() bool`

HasAccountCreationNotificationRequestCriteria returns a boolean if a field has been set.

### GetAccountDeletionNotificationRequestCriteria

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetAccountDeletionNotificationRequestCriteria() string`

GetAccountDeletionNotificationRequestCriteria returns the AccountDeletionNotificationRequestCriteria field if non-nil, zero value otherwise.

### GetAccountDeletionNotificationRequestCriteriaOk

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetAccountDeletionNotificationRequestCriteriaOk() (*string, bool)`

GetAccountDeletionNotificationRequestCriteriaOk returns a tuple with the AccountDeletionNotificationRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDeletionNotificationRequestCriteria

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) SetAccountDeletionNotificationRequestCriteria(v string)`

SetAccountDeletionNotificationRequestCriteria sets AccountDeletionNotificationRequestCriteria field to given value.

### HasAccountDeletionNotificationRequestCriteria

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) HasAccountDeletionNotificationRequestCriteria() bool`

HasAccountDeletionNotificationRequestCriteria returns a boolean if a field has been set.

### GetAccountUpdateNotificationRequestCriteria

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetAccountUpdateNotificationRequestCriteria() string`

GetAccountUpdateNotificationRequestCriteria returns the AccountUpdateNotificationRequestCriteria field if non-nil, zero value otherwise.

### GetAccountUpdateNotificationRequestCriteriaOk

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) GetAccountUpdateNotificationRequestCriteriaOk() (*string, bool)`

GetAccountUpdateNotificationRequestCriteriaOk returns a tuple with the AccountUpdateNotificationRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUpdateNotificationRequestCriteria

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) SetAccountUpdateNotificationRequestCriteria(v string)`

SetAccountUpdateNotificationRequestCriteria sets AccountUpdateNotificationRequestCriteria field to given value.

### HasAccountUpdateNotificationRequestCriteria

`func (o *AddAdminAlertAccountStatusNotificationHandlerRequest) HasAccountUpdateNotificationRequestCriteria() bool`

HasAccountUpdateNotificationRequestCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


