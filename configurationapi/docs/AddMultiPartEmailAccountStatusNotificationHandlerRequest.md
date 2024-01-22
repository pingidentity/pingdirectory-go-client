# AddMultiPartEmailAccountStatusNotificationHandlerRequest

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
**AccountAuthenticatedMessageTemplate** | Pointer to **string** | The path to a file containing the template to use to generate the email message to send in the event that an account has successfully authenticated in a bind operation that matches the criteria provided in the account-authentication-notification-request-criteria property. | [optional] 
**AccountCreatedMessageTemplate** | Pointer to **string** | The path to a file containing the template to use to generate the email message to send in the event that a new account is created in an add request that matches the criteria provided in the account-creation-notification-request-criteria property. | [optional] 
**AccountDeletedMessageTemplate** | Pointer to **string** | The path to a file containing the template to use to generate the email message to send in the event that an existing accout has been removed in a delete request that matches the criteria provided in the account-deletion-notification-request-criteria property. | [optional] 
**AccountUpdatedMessageTemplate** | Pointer to **string** | The path to a file containing the template to use to generate the email message to send in the event that an existing account is updated with a modify or modify DN operation that matches the criteria provided in the account-update-notification-request-criteria property. | [optional] 
**BindPasswordFailedValidationMessageTemplate** | Pointer to **string** | The path to a file containing the template to use to generate the email message to send in the event that a user authenticated with a password that failed to satisfy the criteria for one or more of the configured password validators. | [optional] 
**MustChangePasswordMessageTemplate** | Pointer to **string** | The path to a file containing the template to use to generate the email message to send in the event that a user successfully authenticates to the server but will be required to choose a new password before they will be allowed to perform any other operations. | [optional] 
**Description** | Pointer to **string** | A description for this Account Status Notification Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Account Status Notification Handler is enabled. Only enabled handlers are invoked whenever a related event occurs in the server. | 
**Asynchronous** | Pointer to **bool** | Indicates whether the server should attempt to invoke this Account Status Notification Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver a message) will not delay processing for the operation that triggered the notification. | [optional] 
**AccountAuthenticationNotificationResultCriteria** | Pointer to **string** | A result criteria object that identifies which successful bind operations should result in account authentication notifications for this handler. | [optional] 
**AccountCreationNotificationRequestCriteria** | Pointer to **string** | A request criteria object that identifies which add requests should result in account creation notifications for this handler. | [optional] 
**AccountDeletionNotificationRequestCriteria** | Pointer to **string** | A request criteria object that identifies which delete requests should result in account deletion notifications for this handler. | [optional] 
**AccountUpdateNotificationRequestCriteria** | Pointer to **string** | A request criteria object that identifies which modify and modify DN requests should result in account update notifications for this handler. | [optional] 
**HandlerName** | **string** | Name of the new Account Status Notification Handler | 

## Methods

### NewAddMultiPartEmailAccountStatusNotificationHandlerRequest

`func NewAddMultiPartEmailAccountStatusNotificationHandlerRequest(schemas []EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn, enabled bool, handlerName string, ) *AddMultiPartEmailAccountStatusNotificationHandlerRequest`

NewAddMultiPartEmailAccountStatusNotificationHandlerRequest instantiates a new AddMultiPartEmailAccountStatusNotificationHandlerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddMultiPartEmailAccountStatusNotificationHandlerRequestWithDefaults

`func NewAddMultiPartEmailAccountStatusNotificationHandlerRequestWithDefaults() *AddMultiPartEmailAccountStatusNotificationHandlerRequest`

NewAddMultiPartEmailAccountStatusNotificationHandlerRequestWithDefaults instantiates a new AddMultiPartEmailAccountStatusNotificationHandlerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetSchemas() []EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetSchemasOk() (*[]EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) SetSchemas(v []EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAccountTemporarilyFailureLockedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountTemporarilyFailureLockedMessageTemplate() string`

GetAccountTemporarilyFailureLockedMessageTemplate returns the AccountTemporarilyFailureLockedMessageTemplate field if non-nil, zero value otherwise.

### GetAccountTemporarilyFailureLockedMessageTemplateOk

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountTemporarilyFailureLockedMessageTemplateOk() (*string, bool)`

GetAccountTemporarilyFailureLockedMessageTemplateOk returns a tuple with the AccountTemporarilyFailureLockedMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTemporarilyFailureLockedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) SetAccountTemporarilyFailureLockedMessageTemplate(v string)`

SetAccountTemporarilyFailureLockedMessageTemplate sets AccountTemporarilyFailureLockedMessageTemplate field to given value.

### HasAccountTemporarilyFailureLockedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) HasAccountTemporarilyFailureLockedMessageTemplate() bool`

HasAccountTemporarilyFailureLockedMessageTemplate returns a boolean if a field has been set.

### GetAccountPermanentlyFailureLockedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountPermanentlyFailureLockedMessageTemplate() string`

GetAccountPermanentlyFailureLockedMessageTemplate returns the AccountPermanentlyFailureLockedMessageTemplate field if non-nil, zero value otherwise.

### GetAccountPermanentlyFailureLockedMessageTemplateOk

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountPermanentlyFailureLockedMessageTemplateOk() (*string, bool)`

GetAccountPermanentlyFailureLockedMessageTemplateOk returns a tuple with the AccountPermanentlyFailureLockedMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPermanentlyFailureLockedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) SetAccountPermanentlyFailureLockedMessageTemplate(v string)`

SetAccountPermanentlyFailureLockedMessageTemplate sets AccountPermanentlyFailureLockedMessageTemplate field to given value.

### HasAccountPermanentlyFailureLockedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) HasAccountPermanentlyFailureLockedMessageTemplate() bool`

HasAccountPermanentlyFailureLockedMessageTemplate returns a boolean if a field has been set.

### GetAccountIdleLockedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountIdleLockedMessageTemplate() string`

GetAccountIdleLockedMessageTemplate returns the AccountIdleLockedMessageTemplate field if non-nil, zero value otherwise.

### GetAccountIdleLockedMessageTemplateOk

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountIdleLockedMessageTemplateOk() (*string, bool)`

GetAccountIdleLockedMessageTemplateOk returns a tuple with the AccountIdleLockedMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdleLockedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) SetAccountIdleLockedMessageTemplate(v string)`

SetAccountIdleLockedMessageTemplate sets AccountIdleLockedMessageTemplate field to given value.

### HasAccountIdleLockedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) HasAccountIdleLockedMessageTemplate() bool`

HasAccountIdleLockedMessageTemplate returns a boolean if a field has been set.

### GetAccountResetLockedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountResetLockedMessageTemplate() string`

GetAccountResetLockedMessageTemplate returns the AccountResetLockedMessageTemplate field if non-nil, zero value otherwise.

### GetAccountResetLockedMessageTemplateOk

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountResetLockedMessageTemplateOk() (*string, bool)`

GetAccountResetLockedMessageTemplateOk returns a tuple with the AccountResetLockedMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountResetLockedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) SetAccountResetLockedMessageTemplate(v string)`

SetAccountResetLockedMessageTemplate sets AccountResetLockedMessageTemplate field to given value.

### HasAccountResetLockedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) HasAccountResetLockedMessageTemplate() bool`

HasAccountResetLockedMessageTemplate returns a boolean if a field has been set.

### GetAccountUnlockedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountUnlockedMessageTemplate() string`

GetAccountUnlockedMessageTemplate returns the AccountUnlockedMessageTemplate field if non-nil, zero value otherwise.

### GetAccountUnlockedMessageTemplateOk

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountUnlockedMessageTemplateOk() (*string, bool)`

GetAccountUnlockedMessageTemplateOk returns a tuple with the AccountUnlockedMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUnlockedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) SetAccountUnlockedMessageTemplate(v string)`

SetAccountUnlockedMessageTemplate sets AccountUnlockedMessageTemplate field to given value.

### HasAccountUnlockedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) HasAccountUnlockedMessageTemplate() bool`

HasAccountUnlockedMessageTemplate returns a boolean if a field has been set.

### GetAccountDisabledMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountDisabledMessageTemplate() string`

GetAccountDisabledMessageTemplate returns the AccountDisabledMessageTemplate field if non-nil, zero value otherwise.

### GetAccountDisabledMessageTemplateOk

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountDisabledMessageTemplateOk() (*string, bool)`

GetAccountDisabledMessageTemplateOk returns a tuple with the AccountDisabledMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDisabledMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) SetAccountDisabledMessageTemplate(v string)`

SetAccountDisabledMessageTemplate sets AccountDisabledMessageTemplate field to given value.

### HasAccountDisabledMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) HasAccountDisabledMessageTemplate() bool`

HasAccountDisabledMessageTemplate returns a boolean if a field has been set.

### GetAccountEnabledMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountEnabledMessageTemplate() string`

GetAccountEnabledMessageTemplate returns the AccountEnabledMessageTemplate field if non-nil, zero value otherwise.

### GetAccountEnabledMessageTemplateOk

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountEnabledMessageTemplateOk() (*string, bool)`

GetAccountEnabledMessageTemplateOk returns a tuple with the AccountEnabledMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountEnabledMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) SetAccountEnabledMessageTemplate(v string)`

SetAccountEnabledMessageTemplate sets AccountEnabledMessageTemplate field to given value.

### HasAccountEnabledMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) HasAccountEnabledMessageTemplate() bool`

HasAccountEnabledMessageTemplate returns a boolean if a field has been set.

### GetAccountNotYetActiveMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountNotYetActiveMessageTemplate() string`

GetAccountNotYetActiveMessageTemplate returns the AccountNotYetActiveMessageTemplate field if non-nil, zero value otherwise.

### GetAccountNotYetActiveMessageTemplateOk

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountNotYetActiveMessageTemplateOk() (*string, bool)`

GetAccountNotYetActiveMessageTemplateOk returns a tuple with the AccountNotYetActiveMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNotYetActiveMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) SetAccountNotYetActiveMessageTemplate(v string)`

SetAccountNotYetActiveMessageTemplate sets AccountNotYetActiveMessageTemplate field to given value.

### HasAccountNotYetActiveMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) HasAccountNotYetActiveMessageTemplate() bool`

HasAccountNotYetActiveMessageTemplate returns a boolean if a field has been set.

### GetAccountExpiredMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountExpiredMessageTemplate() string`

GetAccountExpiredMessageTemplate returns the AccountExpiredMessageTemplate field if non-nil, zero value otherwise.

### GetAccountExpiredMessageTemplateOk

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountExpiredMessageTemplateOk() (*string, bool)`

GetAccountExpiredMessageTemplateOk returns a tuple with the AccountExpiredMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountExpiredMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) SetAccountExpiredMessageTemplate(v string)`

SetAccountExpiredMessageTemplate sets AccountExpiredMessageTemplate field to given value.

### HasAccountExpiredMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) HasAccountExpiredMessageTemplate() bool`

HasAccountExpiredMessageTemplate returns a boolean if a field has been set.

### GetPasswordExpiredMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetPasswordExpiredMessageTemplate() string`

GetPasswordExpiredMessageTemplate returns the PasswordExpiredMessageTemplate field if non-nil, zero value otherwise.

### GetPasswordExpiredMessageTemplateOk

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetPasswordExpiredMessageTemplateOk() (*string, bool)`

GetPasswordExpiredMessageTemplateOk returns a tuple with the PasswordExpiredMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpiredMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) SetPasswordExpiredMessageTemplate(v string)`

SetPasswordExpiredMessageTemplate sets PasswordExpiredMessageTemplate field to given value.

### HasPasswordExpiredMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) HasPasswordExpiredMessageTemplate() bool`

HasPasswordExpiredMessageTemplate returns a boolean if a field has been set.

### GetPasswordExpiringMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetPasswordExpiringMessageTemplate() string`

GetPasswordExpiringMessageTemplate returns the PasswordExpiringMessageTemplate field if non-nil, zero value otherwise.

### GetPasswordExpiringMessageTemplateOk

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetPasswordExpiringMessageTemplateOk() (*string, bool)`

GetPasswordExpiringMessageTemplateOk returns a tuple with the PasswordExpiringMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpiringMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) SetPasswordExpiringMessageTemplate(v string)`

SetPasswordExpiringMessageTemplate sets PasswordExpiringMessageTemplate field to given value.

### HasPasswordExpiringMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) HasPasswordExpiringMessageTemplate() bool`

HasPasswordExpiringMessageTemplate returns a boolean if a field has been set.

### GetPasswordResetMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetPasswordResetMessageTemplate() string`

GetPasswordResetMessageTemplate returns the PasswordResetMessageTemplate field if non-nil, zero value otherwise.

### GetPasswordResetMessageTemplateOk

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetPasswordResetMessageTemplateOk() (*string, bool)`

GetPasswordResetMessageTemplateOk returns a tuple with the PasswordResetMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordResetMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) SetPasswordResetMessageTemplate(v string)`

SetPasswordResetMessageTemplate sets PasswordResetMessageTemplate field to given value.

### HasPasswordResetMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) HasPasswordResetMessageTemplate() bool`

HasPasswordResetMessageTemplate returns a boolean if a field has been set.

### GetPasswordChangedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetPasswordChangedMessageTemplate() string`

GetPasswordChangedMessageTemplate returns the PasswordChangedMessageTemplate field if non-nil, zero value otherwise.

### GetPasswordChangedMessageTemplateOk

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetPasswordChangedMessageTemplateOk() (*string, bool)`

GetPasswordChangedMessageTemplateOk returns a tuple with the PasswordChangedMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordChangedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) SetPasswordChangedMessageTemplate(v string)`

SetPasswordChangedMessageTemplate sets PasswordChangedMessageTemplate field to given value.

### HasPasswordChangedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) HasPasswordChangedMessageTemplate() bool`

HasPasswordChangedMessageTemplate returns a boolean if a field has been set.

### GetAccountAuthenticatedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountAuthenticatedMessageTemplate() string`

GetAccountAuthenticatedMessageTemplate returns the AccountAuthenticatedMessageTemplate field if non-nil, zero value otherwise.

### GetAccountAuthenticatedMessageTemplateOk

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountAuthenticatedMessageTemplateOk() (*string, bool)`

GetAccountAuthenticatedMessageTemplateOk returns a tuple with the AccountAuthenticatedMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAuthenticatedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) SetAccountAuthenticatedMessageTemplate(v string)`

SetAccountAuthenticatedMessageTemplate sets AccountAuthenticatedMessageTemplate field to given value.

### HasAccountAuthenticatedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) HasAccountAuthenticatedMessageTemplate() bool`

HasAccountAuthenticatedMessageTemplate returns a boolean if a field has been set.

### GetAccountCreatedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountCreatedMessageTemplate() string`

GetAccountCreatedMessageTemplate returns the AccountCreatedMessageTemplate field if non-nil, zero value otherwise.

### GetAccountCreatedMessageTemplateOk

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountCreatedMessageTemplateOk() (*string, bool)`

GetAccountCreatedMessageTemplateOk returns a tuple with the AccountCreatedMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCreatedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) SetAccountCreatedMessageTemplate(v string)`

SetAccountCreatedMessageTemplate sets AccountCreatedMessageTemplate field to given value.

### HasAccountCreatedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) HasAccountCreatedMessageTemplate() bool`

HasAccountCreatedMessageTemplate returns a boolean if a field has been set.

### GetAccountDeletedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountDeletedMessageTemplate() string`

GetAccountDeletedMessageTemplate returns the AccountDeletedMessageTemplate field if non-nil, zero value otherwise.

### GetAccountDeletedMessageTemplateOk

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountDeletedMessageTemplateOk() (*string, bool)`

GetAccountDeletedMessageTemplateOk returns a tuple with the AccountDeletedMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDeletedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) SetAccountDeletedMessageTemplate(v string)`

SetAccountDeletedMessageTemplate sets AccountDeletedMessageTemplate field to given value.

### HasAccountDeletedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) HasAccountDeletedMessageTemplate() bool`

HasAccountDeletedMessageTemplate returns a boolean if a field has been set.

### GetAccountUpdatedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountUpdatedMessageTemplate() string`

GetAccountUpdatedMessageTemplate returns the AccountUpdatedMessageTemplate field if non-nil, zero value otherwise.

### GetAccountUpdatedMessageTemplateOk

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountUpdatedMessageTemplateOk() (*string, bool)`

GetAccountUpdatedMessageTemplateOk returns a tuple with the AccountUpdatedMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUpdatedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) SetAccountUpdatedMessageTemplate(v string)`

SetAccountUpdatedMessageTemplate sets AccountUpdatedMessageTemplate field to given value.

### HasAccountUpdatedMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) HasAccountUpdatedMessageTemplate() bool`

HasAccountUpdatedMessageTemplate returns a boolean if a field has been set.

### GetBindPasswordFailedValidationMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetBindPasswordFailedValidationMessageTemplate() string`

GetBindPasswordFailedValidationMessageTemplate returns the BindPasswordFailedValidationMessageTemplate field if non-nil, zero value otherwise.

### GetBindPasswordFailedValidationMessageTemplateOk

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetBindPasswordFailedValidationMessageTemplateOk() (*string, bool)`

GetBindPasswordFailedValidationMessageTemplateOk returns a tuple with the BindPasswordFailedValidationMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPasswordFailedValidationMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) SetBindPasswordFailedValidationMessageTemplate(v string)`

SetBindPasswordFailedValidationMessageTemplate sets BindPasswordFailedValidationMessageTemplate field to given value.

### HasBindPasswordFailedValidationMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) HasBindPasswordFailedValidationMessageTemplate() bool`

HasBindPasswordFailedValidationMessageTemplate returns a boolean if a field has been set.

### GetMustChangePasswordMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetMustChangePasswordMessageTemplate() string`

GetMustChangePasswordMessageTemplate returns the MustChangePasswordMessageTemplate field if non-nil, zero value otherwise.

### GetMustChangePasswordMessageTemplateOk

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetMustChangePasswordMessageTemplateOk() (*string, bool)`

GetMustChangePasswordMessageTemplateOk returns a tuple with the MustChangePasswordMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustChangePasswordMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) SetMustChangePasswordMessageTemplate(v string)`

SetMustChangePasswordMessageTemplate sets MustChangePasswordMessageTemplate field to given value.

### HasMustChangePasswordMessageTemplate

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) HasMustChangePasswordMessageTemplate() bool`

HasMustChangePasswordMessageTemplate returns a boolean if a field has been set.

### GetDescription

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAsynchronous

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAsynchronous() bool`

GetAsynchronous returns the Asynchronous field if non-nil, zero value otherwise.

### GetAsynchronousOk

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAsynchronousOk() (*bool, bool)`

GetAsynchronousOk returns a tuple with the Asynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsynchronous

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) SetAsynchronous(v bool)`

SetAsynchronous sets Asynchronous field to given value.

### HasAsynchronous

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) HasAsynchronous() bool`

HasAsynchronous returns a boolean if a field has been set.

### GetAccountAuthenticationNotificationResultCriteria

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountAuthenticationNotificationResultCriteria() string`

GetAccountAuthenticationNotificationResultCriteria returns the AccountAuthenticationNotificationResultCriteria field if non-nil, zero value otherwise.

### GetAccountAuthenticationNotificationResultCriteriaOk

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountAuthenticationNotificationResultCriteriaOk() (*string, bool)`

GetAccountAuthenticationNotificationResultCriteriaOk returns a tuple with the AccountAuthenticationNotificationResultCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAuthenticationNotificationResultCriteria

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) SetAccountAuthenticationNotificationResultCriteria(v string)`

SetAccountAuthenticationNotificationResultCriteria sets AccountAuthenticationNotificationResultCriteria field to given value.

### HasAccountAuthenticationNotificationResultCriteria

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) HasAccountAuthenticationNotificationResultCriteria() bool`

HasAccountAuthenticationNotificationResultCriteria returns a boolean if a field has been set.

### GetAccountCreationNotificationRequestCriteria

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountCreationNotificationRequestCriteria() string`

GetAccountCreationNotificationRequestCriteria returns the AccountCreationNotificationRequestCriteria field if non-nil, zero value otherwise.

### GetAccountCreationNotificationRequestCriteriaOk

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountCreationNotificationRequestCriteriaOk() (*string, bool)`

GetAccountCreationNotificationRequestCriteriaOk returns a tuple with the AccountCreationNotificationRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCreationNotificationRequestCriteria

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) SetAccountCreationNotificationRequestCriteria(v string)`

SetAccountCreationNotificationRequestCriteria sets AccountCreationNotificationRequestCriteria field to given value.

### HasAccountCreationNotificationRequestCriteria

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) HasAccountCreationNotificationRequestCriteria() bool`

HasAccountCreationNotificationRequestCriteria returns a boolean if a field has been set.

### GetAccountDeletionNotificationRequestCriteria

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountDeletionNotificationRequestCriteria() string`

GetAccountDeletionNotificationRequestCriteria returns the AccountDeletionNotificationRequestCriteria field if non-nil, zero value otherwise.

### GetAccountDeletionNotificationRequestCriteriaOk

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountDeletionNotificationRequestCriteriaOk() (*string, bool)`

GetAccountDeletionNotificationRequestCriteriaOk returns a tuple with the AccountDeletionNotificationRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDeletionNotificationRequestCriteria

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) SetAccountDeletionNotificationRequestCriteria(v string)`

SetAccountDeletionNotificationRequestCriteria sets AccountDeletionNotificationRequestCriteria field to given value.

### HasAccountDeletionNotificationRequestCriteria

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) HasAccountDeletionNotificationRequestCriteria() bool`

HasAccountDeletionNotificationRequestCriteria returns a boolean if a field has been set.

### GetAccountUpdateNotificationRequestCriteria

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountUpdateNotificationRequestCriteria() string`

GetAccountUpdateNotificationRequestCriteria returns the AccountUpdateNotificationRequestCriteria field if non-nil, zero value otherwise.

### GetAccountUpdateNotificationRequestCriteriaOk

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetAccountUpdateNotificationRequestCriteriaOk() (*string, bool)`

GetAccountUpdateNotificationRequestCriteriaOk returns a tuple with the AccountUpdateNotificationRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUpdateNotificationRequestCriteria

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) SetAccountUpdateNotificationRequestCriteria(v string)`

SetAccountUpdateNotificationRequestCriteria sets AccountUpdateNotificationRequestCriteria field to given value.

### HasAccountUpdateNotificationRequestCriteria

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) HasAccountUpdateNotificationRequestCriteria() bool`

HasAccountUpdateNotificationRequestCriteria returns a boolean if a field has been set.

### GetHandlerName

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetHandlerName() string`

GetHandlerName returns the HandlerName field if non-nil, zero value otherwise.

### GetHandlerNameOk

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) GetHandlerNameOk() (*string, bool)`

GetHandlerNameOk returns a tuple with the HandlerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerName

`func (o *AddMultiPartEmailAccountStatusNotificationHandlerRequest) SetHandlerName(v string)`

SetHandlerName sets HandlerName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


