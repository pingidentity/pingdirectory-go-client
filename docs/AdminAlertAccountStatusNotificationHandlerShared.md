# AdminAlertAccountStatusNotificationHandlerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumadminAlertAccountStatusNotificationHandlerSchemaUrn**](EnumadminAlertAccountStatusNotificationHandlerSchemaUrn.md) |  | 
**AccountStatusNotificationType** | [**[]EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp**](EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp.md) |  | 
**Description** | Pointer to **string** | A description for this Account Status Notification Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Account Status Notification Handler is enabled. Only enabled handlers are invoked whenever a related event occurs in the server. | 
**Asynchronous** | Pointer to **bool** | Indicates whether the server should attempt to invoke this Account Status Notification Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver a message) will not delay processing for the operation that triggered the notification. | [optional] 
**AccountCreationNotificationRequestCriteria** | Pointer to **string** | A request criteria object that identifies which add requests should result in account creation notifications for this handler. | [optional] 
**AccountUpdateNotificationRequestCriteria** | Pointer to **string** | A request criteria object that identifies which modify and modify DN requests should result in account update notifications for this handler. | [optional] 

## Methods

### NewAdminAlertAccountStatusNotificationHandlerShared

`func NewAdminAlertAccountStatusNotificationHandlerShared(schemas []EnumadminAlertAccountStatusNotificationHandlerSchemaUrn, accountStatusNotificationType []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp, enabled bool, ) *AdminAlertAccountStatusNotificationHandlerShared`

NewAdminAlertAccountStatusNotificationHandlerShared instantiates a new AdminAlertAccountStatusNotificationHandlerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminAlertAccountStatusNotificationHandlerSharedWithDefaults

`func NewAdminAlertAccountStatusNotificationHandlerSharedWithDefaults() *AdminAlertAccountStatusNotificationHandlerShared`

NewAdminAlertAccountStatusNotificationHandlerSharedWithDefaults instantiates a new AdminAlertAccountStatusNotificationHandlerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AdminAlertAccountStatusNotificationHandlerShared) GetSchemas() []EnumadminAlertAccountStatusNotificationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AdminAlertAccountStatusNotificationHandlerShared) GetSchemasOk() (*[]EnumadminAlertAccountStatusNotificationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AdminAlertAccountStatusNotificationHandlerShared) SetSchemas(v []EnumadminAlertAccountStatusNotificationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAccountStatusNotificationType

`func (o *AdminAlertAccountStatusNotificationHandlerShared) GetAccountStatusNotificationType() []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp`

GetAccountStatusNotificationType returns the AccountStatusNotificationType field if non-nil, zero value otherwise.

### GetAccountStatusNotificationTypeOk

`func (o *AdminAlertAccountStatusNotificationHandlerShared) GetAccountStatusNotificationTypeOk() (*[]EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp, bool)`

GetAccountStatusNotificationTypeOk returns a tuple with the AccountStatusNotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatusNotificationType

`func (o *AdminAlertAccountStatusNotificationHandlerShared) SetAccountStatusNotificationType(v []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp)`

SetAccountStatusNotificationType sets AccountStatusNotificationType field to given value.


### GetDescription

`func (o *AdminAlertAccountStatusNotificationHandlerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdminAlertAccountStatusNotificationHandlerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdminAlertAccountStatusNotificationHandlerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdminAlertAccountStatusNotificationHandlerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AdminAlertAccountStatusNotificationHandlerShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AdminAlertAccountStatusNotificationHandlerShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AdminAlertAccountStatusNotificationHandlerShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAsynchronous

`func (o *AdminAlertAccountStatusNotificationHandlerShared) GetAsynchronous() bool`

GetAsynchronous returns the Asynchronous field if non-nil, zero value otherwise.

### GetAsynchronousOk

`func (o *AdminAlertAccountStatusNotificationHandlerShared) GetAsynchronousOk() (*bool, bool)`

GetAsynchronousOk returns a tuple with the Asynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsynchronous

`func (o *AdminAlertAccountStatusNotificationHandlerShared) SetAsynchronous(v bool)`

SetAsynchronous sets Asynchronous field to given value.

### HasAsynchronous

`func (o *AdminAlertAccountStatusNotificationHandlerShared) HasAsynchronous() bool`

HasAsynchronous returns a boolean if a field has been set.

### GetAccountCreationNotificationRequestCriteria

`func (o *AdminAlertAccountStatusNotificationHandlerShared) GetAccountCreationNotificationRequestCriteria() string`

GetAccountCreationNotificationRequestCriteria returns the AccountCreationNotificationRequestCriteria field if non-nil, zero value otherwise.

### GetAccountCreationNotificationRequestCriteriaOk

`func (o *AdminAlertAccountStatusNotificationHandlerShared) GetAccountCreationNotificationRequestCriteriaOk() (*string, bool)`

GetAccountCreationNotificationRequestCriteriaOk returns a tuple with the AccountCreationNotificationRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCreationNotificationRequestCriteria

`func (o *AdminAlertAccountStatusNotificationHandlerShared) SetAccountCreationNotificationRequestCriteria(v string)`

SetAccountCreationNotificationRequestCriteria sets AccountCreationNotificationRequestCriteria field to given value.

### HasAccountCreationNotificationRequestCriteria

`func (o *AdminAlertAccountStatusNotificationHandlerShared) HasAccountCreationNotificationRequestCriteria() bool`

HasAccountCreationNotificationRequestCriteria returns a boolean if a field has been set.

### GetAccountUpdateNotificationRequestCriteria

`func (o *AdminAlertAccountStatusNotificationHandlerShared) GetAccountUpdateNotificationRequestCriteria() string`

GetAccountUpdateNotificationRequestCriteria returns the AccountUpdateNotificationRequestCriteria field if non-nil, zero value otherwise.

### GetAccountUpdateNotificationRequestCriteriaOk

`func (o *AdminAlertAccountStatusNotificationHandlerShared) GetAccountUpdateNotificationRequestCriteriaOk() (*string, bool)`

GetAccountUpdateNotificationRequestCriteriaOk returns a tuple with the AccountUpdateNotificationRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUpdateNotificationRequestCriteria

`func (o *AdminAlertAccountStatusNotificationHandlerShared) SetAccountUpdateNotificationRequestCriteria(v string)`

SetAccountUpdateNotificationRequestCriteria sets AccountUpdateNotificationRequestCriteria field to given value.

### HasAccountUpdateNotificationRequestCriteria

`func (o *AdminAlertAccountStatusNotificationHandlerShared) HasAccountUpdateNotificationRequestCriteria() bool`

HasAccountUpdateNotificationRequestCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


