# MultiPartEmailAccountStatusNotificationHandlerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn**](EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn.md) |  | 
**AccountTemporarilyFailureLockedMessageTemplate** | Pointer to **string** | The path to a file containing the template to use to generate the email message to send in the event that an account becomes temporarily locked as a result of too many authentication failures. | [optional] 
**AccountPermanentlyFailureLockedMessageTemplate** | Pointer to **string** | The path to a file containing the template to use to generate the email message to send in the event that an account becomes permanently locked as a result of too many authentication failures. | [optional] 
**AccountIdleLockedMessageTemplate** | Pointer to **string** | The path to a file containing the template to use to generate the email message to send in the event that authentication attempt fails because it has been too long since the user last successfully authenticated. | [optional] 
**AccountResetLockedMessageTemplate** | Pointer to **string** | The path to a file containing the template to use to generate the email message to send in the event that authentication attempt fails because the user failed to choose a new password in a timely manner after an administrative reset. | [optional] 
**AccountUnlockedMessageTemplate** | Pointer to **string** | The path to a file containing the template to use to generate the email message to send in the event that a user&#39;s account has been unlocked (e.g., by an administrative password reset). | [optional] 
**AccountDisabledMessageTemplate** | Pointer to **string** | The path to a file containing the template to use to generate the email message to send in the event that a user&#39;s account is disabled by an administrator. | [optional] 
**AccountEnabledMessageTemplate** | Pointer to **string** | The path to a file containing the template to use to generate the email message to send in the event that a user&#39;s account is enabled by an administrator. | [optional] 
**AccountNotYetActiveMessageTemplate** | Pointer to **string** | The path to a file containing the template to use to generate the email message to send in the event that an authentication attempt fails because the account has an activation time that is in the future. | [optional] 
**AccountExpiredMessageTemplate** | Pointer to **string** | The path to a file containing the template to use to generate the email message to send in the event that an authentication attempt fails because the account has an expiration time that is in the past. | [optional] 
**PasswordExpiredMessageTemplate** | Pointer to **string** | The path to a file containing the template to use to generate the email message to send in the event that an authentication attempt fails because the account has an expired password. | [optional] 
**PasswordExpiringMessageTemplate** | Pointer to **string** | The path to a file containing the template to use to generate the email message to send in the event that an authentication attempt succeeds, but the user&#39;s password is about to expire. This notification will only be generated the first time the user authenticates within the window of time that the server should warn about an upcoming password expiration. | [optional] 
**PasswordResetMessageTemplate** | Pointer to **string** | The path to a file containing the template to use to generate the email message to send in the event that a user&#39;s password has been reset by an administrator. | [optional] 
**PasswordChangedMessageTemplate** | Pointer to **string** | The path to a file containing the template to use to generate the email message to send in the event that a user changes their own password. | [optional] 
**AccountCreatedMessageTemplate** | Pointer to **string** | The path to a file containing the template to use to generate the email message to send in the event that a new account is created in an add request that matches the criteria provided in the account-creation-notification-request-criteria property. | [optional] 
**AccountUpdatedMessageTemplate** | Pointer to **string** | The path to a file containing the template to use to generate the email message to send in the event that an existing account is updated with a modify or modify DN operation that matches the criteria provided in the account-update-notification-request-criteria property. | [optional] 
**BindPasswordFailedValidationMessageTemplate** | Pointer to **string** | The path to a file containing the template to use to generate the email message to send in the event that a user authenticated with a password that failed to satisfy the criteria for one or more of the configured password validators. | [optional] 
**MustChangePasswordMessageTemplate** | Pointer to **string** | The path to a file containing the template to use to generate the email message to send in the event that a user successfully authenticates to the server but will be required to choose a new password before they will be allowed to perform any other operations. | [optional] 
**Description** | Pointer to **string** | A description for this Account Status Notification Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Account Status Notification Handler is enabled. Only enabled handlers are invoked whenever a related event occurs in the server. | 
**Asynchronous** | Pointer to **bool** | Indicates whether the server should attempt to invoke this Account Status Notification Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver a message) will not delay processing for the operation that triggered the notification. | [optional] 
**AccountCreationNotificationRequestCriteria** | Pointer to **string** | A request criteria object that identifies which add requests should result in account creation notifications for this handler. | [optional] 
**AccountUpdateNotificationRequestCriteria** | Pointer to **string** | A request criteria object that identifies which modify and modify DN requests should result in account update notifications for this handler. | [optional] 

## Methods

### NewMultiPartEmailAccountStatusNotificationHandlerShared

`func NewMultiPartEmailAccountStatusNotificationHandlerShared(schemas []EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn, enabled bool, ) *MultiPartEmailAccountStatusNotificationHandlerShared`

NewMultiPartEmailAccountStatusNotificationHandlerShared instantiates a new MultiPartEmailAccountStatusNotificationHandlerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiPartEmailAccountStatusNotificationHandlerSharedWithDefaults

`func NewMultiPartEmailAccountStatusNotificationHandlerSharedWithDefaults() *MultiPartEmailAccountStatusNotificationHandlerShared`

NewMultiPartEmailAccountStatusNotificationHandlerSharedWithDefaults instantiates a new MultiPartEmailAccountStatusNotificationHandlerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetSchemas() []EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetSchemasOk() (*[]EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) SetSchemas(v []EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAccountTemporarilyFailureLockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetAccountTemporarilyFailureLockedMessageTemplate() string`

GetAccountTemporarilyFailureLockedMessageTemplate returns the AccountTemporarilyFailureLockedMessageTemplate field if non-nil, zero value otherwise.

### GetAccountTemporarilyFailureLockedMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetAccountTemporarilyFailureLockedMessageTemplateOk() (*string, bool)`

GetAccountTemporarilyFailureLockedMessageTemplateOk returns a tuple with the AccountTemporarilyFailureLockedMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTemporarilyFailureLockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) SetAccountTemporarilyFailureLockedMessageTemplate(v string)`

SetAccountTemporarilyFailureLockedMessageTemplate sets AccountTemporarilyFailureLockedMessageTemplate field to given value.

### HasAccountTemporarilyFailureLockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) HasAccountTemporarilyFailureLockedMessageTemplate() bool`

HasAccountTemporarilyFailureLockedMessageTemplate returns a boolean if a field has been set.

### GetAccountPermanentlyFailureLockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetAccountPermanentlyFailureLockedMessageTemplate() string`

GetAccountPermanentlyFailureLockedMessageTemplate returns the AccountPermanentlyFailureLockedMessageTemplate field if non-nil, zero value otherwise.

### GetAccountPermanentlyFailureLockedMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetAccountPermanentlyFailureLockedMessageTemplateOk() (*string, bool)`

GetAccountPermanentlyFailureLockedMessageTemplateOk returns a tuple with the AccountPermanentlyFailureLockedMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPermanentlyFailureLockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) SetAccountPermanentlyFailureLockedMessageTemplate(v string)`

SetAccountPermanentlyFailureLockedMessageTemplate sets AccountPermanentlyFailureLockedMessageTemplate field to given value.

### HasAccountPermanentlyFailureLockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) HasAccountPermanentlyFailureLockedMessageTemplate() bool`

HasAccountPermanentlyFailureLockedMessageTemplate returns a boolean if a field has been set.

### GetAccountIdleLockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetAccountIdleLockedMessageTemplate() string`

GetAccountIdleLockedMessageTemplate returns the AccountIdleLockedMessageTemplate field if non-nil, zero value otherwise.

### GetAccountIdleLockedMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetAccountIdleLockedMessageTemplateOk() (*string, bool)`

GetAccountIdleLockedMessageTemplateOk returns a tuple with the AccountIdleLockedMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdleLockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) SetAccountIdleLockedMessageTemplate(v string)`

SetAccountIdleLockedMessageTemplate sets AccountIdleLockedMessageTemplate field to given value.

### HasAccountIdleLockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) HasAccountIdleLockedMessageTemplate() bool`

HasAccountIdleLockedMessageTemplate returns a boolean if a field has been set.

### GetAccountResetLockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetAccountResetLockedMessageTemplate() string`

GetAccountResetLockedMessageTemplate returns the AccountResetLockedMessageTemplate field if non-nil, zero value otherwise.

### GetAccountResetLockedMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetAccountResetLockedMessageTemplateOk() (*string, bool)`

GetAccountResetLockedMessageTemplateOk returns a tuple with the AccountResetLockedMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountResetLockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) SetAccountResetLockedMessageTemplate(v string)`

SetAccountResetLockedMessageTemplate sets AccountResetLockedMessageTemplate field to given value.

### HasAccountResetLockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) HasAccountResetLockedMessageTemplate() bool`

HasAccountResetLockedMessageTemplate returns a boolean if a field has been set.

### GetAccountUnlockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetAccountUnlockedMessageTemplate() string`

GetAccountUnlockedMessageTemplate returns the AccountUnlockedMessageTemplate field if non-nil, zero value otherwise.

### GetAccountUnlockedMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetAccountUnlockedMessageTemplateOk() (*string, bool)`

GetAccountUnlockedMessageTemplateOk returns a tuple with the AccountUnlockedMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUnlockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) SetAccountUnlockedMessageTemplate(v string)`

SetAccountUnlockedMessageTemplate sets AccountUnlockedMessageTemplate field to given value.

### HasAccountUnlockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) HasAccountUnlockedMessageTemplate() bool`

HasAccountUnlockedMessageTemplate returns a boolean if a field has been set.

### GetAccountDisabledMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetAccountDisabledMessageTemplate() string`

GetAccountDisabledMessageTemplate returns the AccountDisabledMessageTemplate field if non-nil, zero value otherwise.

### GetAccountDisabledMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetAccountDisabledMessageTemplateOk() (*string, bool)`

GetAccountDisabledMessageTemplateOk returns a tuple with the AccountDisabledMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDisabledMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) SetAccountDisabledMessageTemplate(v string)`

SetAccountDisabledMessageTemplate sets AccountDisabledMessageTemplate field to given value.

### HasAccountDisabledMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) HasAccountDisabledMessageTemplate() bool`

HasAccountDisabledMessageTemplate returns a boolean if a field has been set.

### GetAccountEnabledMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetAccountEnabledMessageTemplate() string`

GetAccountEnabledMessageTemplate returns the AccountEnabledMessageTemplate field if non-nil, zero value otherwise.

### GetAccountEnabledMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetAccountEnabledMessageTemplateOk() (*string, bool)`

GetAccountEnabledMessageTemplateOk returns a tuple with the AccountEnabledMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountEnabledMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) SetAccountEnabledMessageTemplate(v string)`

SetAccountEnabledMessageTemplate sets AccountEnabledMessageTemplate field to given value.

### HasAccountEnabledMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) HasAccountEnabledMessageTemplate() bool`

HasAccountEnabledMessageTemplate returns a boolean if a field has been set.

### GetAccountNotYetActiveMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetAccountNotYetActiveMessageTemplate() string`

GetAccountNotYetActiveMessageTemplate returns the AccountNotYetActiveMessageTemplate field if non-nil, zero value otherwise.

### GetAccountNotYetActiveMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetAccountNotYetActiveMessageTemplateOk() (*string, bool)`

GetAccountNotYetActiveMessageTemplateOk returns a tuple with the AccountNotYetActiveMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNotYetActiveMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) SetAccountNotYetActiveMessageTemplate(v string)`

SetAccountNotYetActiveMessageTemplate sets AccountNotYetActiveMessageTemplate field to given value.

### HasAccountNotYetActiveMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) HasAccountNotYetActiveMessageTemplate() bool`

HasAccountNotYetActiveMessageTemplate returns a boolean if a field has been set.

### GetAccountExpiredMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetAccountExpiredMessageTemplate() string`

GetAccountExpiredMessageTemplate returns the AccountExpiredMessageTemplate field if non-nil, zero value otherwise.

### GetAccountExpiredMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetAccountExpiredMessageTemplateOk() (*string, bool)`

GetAccountExpiredMessageTemplateOk returns a tuple with the AccountExpiredMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountExpiredMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) SetAccountExpiredMessageTemplate(v string)`

SetAccountExpiredMessageTemplate sets AccountExpiredMessageTemplate field to given value.

### HasAccountExpiredMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) HasAccountExpiredMessageTemplate() bool`

HasAccountExpiredMessageTemplate returns a boolean if a field has been set.

### GetPasswordExpiredMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetPasswordExpiredMessageTemplate() string`

GetPasswordExpiredMessageTemplate returns the PasswordExpiredMessageTemplate field if non-nil, zero value otherwise.

### GetPasswordExpiredMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetPasswordExpiredMessageTemplateOk() (*string, bool)`

GetPasswordExpiredMessageTemplateOk returns a tuple with the PasswordExpiredMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpiredMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) SetPasswordExpiredMessageTemplate(v string)`

SetPasswordExpiredMessageTemplate sets PasswordExpiredMessageTemplate field to given value.

### HasPasswordExpiredMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) HasPasswordExpiredMessageTemplate() bool`

HasPasswordExpiredMessageTemplate returns a boolean if a field has been set.

### GetPasswordExpiringMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetPasswordExpiringMessageTemplate() string`

GetPasswordExpiringMessageTemplate returns the PasswordExpiringMessageTemplate field if non-nil, zero value otherwise.

### GetPasswordExpiringMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetPasswordExpiringMessageTemplateOk() (*string, bool)`

GetPasswordExpiringMessageTemplateOk returns a tuple with the PasswordExpiringMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpiringMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) SetPasswordExpiringMessageTemplate(v string)`

SetPasswordExpiringMessageTemplate sets PasswordExpiringMessageTemplate field to given value.

### HasPasswordExpiringMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) HasPasswordExpiringMessageTemplate() bool`

HasPasswordExpiringMessageTemplate returns a boolean if a field has been set.

### GetPasswordResetMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetPasswordResetMessageTemplate() string`

GetPasswordResetMessageTemplate returns the PasswordResetMessageTemplate field if non-nil, zero value otherwise.

### GetPasswordResetMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetPasswordResetMessageTemplateOk() (*string, bool)`

GetPasswordResetMessageTemplateOk returns a tuple with the PasswordResetMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordResetMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) SetPasswordResetMessageTemplate(v string)`

SetPasswordResetMessageTemplate sets PasswordResetMessageTemplate field to given value.

### HasPasswordResetMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) HasPasswordResetMessageTemplate() bool`

HasPasswordResetMessageTemplate returns a boolean if a field has been set.

### GetPasswordChangedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetPasswordChangedMessageTemplate() string`

GetPasswordChangedMessageTemplate returns the PasswordChangedMessageTemplate field if non-nil, zero value otherwise.

### GetPasswordChangedMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetPasswordChangedMessageTemplateOk() (*string, bool)`

GetPasswordChangedMessageTemplateOk returns a tuple with the PasswordChangedMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordChangedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) SetPasswordChangedMessageTemplate(v string)`

SetPasswordChangedMessageTemplate sets PasswordChangedMessageTemplate field to given value.

### HasPasswordChangedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) HasPasswordChangedMessageTemplate() bool`

HasPasswordChangedMessageTemplate returns a boolean if a field has been set.

### GetAccountCreatedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetAccountCreatedMessageTemplate() string`

GetAccountCreatedMessageTemplate returns the AccountCreatedMessageTemplate field if non-nil, zero value otherwise.

### GetAccountCreatedMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetAccountCreatedMessageTemplateOk() (*string, bool)`

GetAccountCreatedMessageTemplateOk returns a tuple with the AccountCreatedMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCreatedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) SetAccountCreatedMessageTemplate(v string)`

SetAccountCreatedMessageTemplate sets AccountCreatedMessageTemplate field to given value.

### HasAccountCreatedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) HasAccountCreatedMessageTemplate() bool`

HasAccountCreatedMessageTemplate returns a boolean if a field has been set.

### GetAccountUpdatedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetAccountUpdatedMessageTemplate() string`

GetAccountUpdatedMessageTemplate returns the AccountUpdatedMessageTemplate field if non-nil, zero value otherwise.

### GetAccountUpdatedMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetAccountUpdatedMessageTemplateOk() (*string, bool)`

GetAccountUpdatedMessageTemplateOk returns a tuple with the AccountUpdatedMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUpdatedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) SetAccountUpdatedMessageTemplate(v string)`

SetAccountUpdatedMessageTemplate sets AccountUpdatedMessageTemplate field to given value.

### HasAccountUpdatedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) HasAccountUpdatedMessageTemplate() bool`

HasAccountUpdatedMessageTemplate returns a boolean if a field has been set.

### GetBindPasswordFailedValidationMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetBindPasswordFailedValidationMessageTemplate() string`

GetBindPasswordFailedValidationMessageTemplate returns the BindPasswordFailedValidationMessageTemplate field if non-nil, zero value otherwise.

### GetBindPasswordFailedValidationMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetBindPasswordFailedValidationMessageTemplateOk() (*string, bool)`

GetBindPasswordFailedValidationMessageTemplateOk returns a tuple with the BindPasswordFailedValidationMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPasswordFailedValidationMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) SetBindPasswordFailedValidationMessageTemplate(v string)`

SetBindPasswordFailedValidationMessageTemplate sets BindPasswordFailedValidationMessageTemplate field to given value.

### HasBindPasswordFailedValidationMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) HasBindPasswordFailedValidationMessageTemplate() bool`

HasBindPasswordFailedValidationMessageTemplate returns a boolean if a field has been set.

### GetMustChangePasswordMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetMustChangePasswordMessageTemplate() string`

GetMustChangePasswordMessageTemplate returns the MustChangePasswordMessageTemplate field if non-nil, zero value otherwise.

### GetMustChangePasswordMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetMustChangePasswordMessageTemplateOk() (*string, bool)`

GetMustChangePasswordMessageTemplateOk returns a tuple with the MustChangePasswordMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustChangePasswordMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) SetMustChangePasswordMessageTemplate(v string)`

SetMustChangePasswordMessageTemplate sets MustChangePasswordMessageTemplate field to given value.

### HasMustChangePasswordMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) HasMustChangePasswordMessageTemplate() bool`

HasMustChangePasswordMessageTemplate returns a boolean if a field has been set.

### GetDescription

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAsynchronous

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetAsynchronous() bool`

GetAsynchronous returns the Asynchronous field if non-nil, zero value otherwise.

### GetAsynchronousOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetAsynchronousOk() (*bool, bool)`

GetAsynchronousOk returns a tuple with the Asynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsynchronous

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) SetAsynchronous(v bool)`

SetAsynchronous sets Asynchronous field to given value.

### HasAsynchronous

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) HasAsynchronous() bool`

HasAsynchronous returns a boolean if a field has been set.

### GetAccountCreationNotificationRequestCriteria

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetAccountCreationNotificationRequestCriteria() string`

GetAccountCreationNotificationRequestCriteria returns the AccountCreationNotificationRequestCriteria field if non-nil, zero value otherwise.

### GetAccountCreationNotificationRequestCriteriaOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetAccountCreationNotificationRequestCriteriaOk() (*string, bool)`

GetAccountCreationNotificationRequestCriteriaOk returns a tuple with the AccountCreationNotificationRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCreationNotificationRequestCriteria

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) SetAccountCreationNotificationRequestCriteria(v string)`

SetAccountCreationNotificationRequestCriteria sets AccountCreationNotificationRequestCriteria field to given value.

### HasAccountCreationNotificationRequestCriteria

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) HasAccountCreationNotificationRequestCriteria() bool`

HasAccountCreationNotificationRequestCriteria returns a boolean if a field has been set.

### GetAccountUpdateNotificationRequestCriteria

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetAccountUpdateNotificationRequestCriteria() string`

GetAccountUpdateNotificationRequestCriteria returns the AccountUpdateNotificationRequestCriteria field if non-nil, zero value otherwise.

### GetAccountUpdateNotificationRequestCriteriaOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) GetAccountUpdateNotificationRequestCriteriaOk() (*string, bool)`

GetAccountUpdateNotificationRequestCriteriaOk returns a tuple with the AccountUpdateNotificationRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUpdateNotificationRequestCriteria

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) SetAccountUpdateNotificationRequestCriteria(v string)`

SetAccountUpdateNotificationRequestCriteria sets AccountUpdateNotificationRequestCriteria field to given value.

### HasAccountUpdateNotificationRequestCriteria

`func (o *MultiPartEmailAccountStatusNotificationHandlerShared) HasAccountUpdateNotificationRequestCriteria() bool`

HasAccountUpdateNotificationRequestCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


