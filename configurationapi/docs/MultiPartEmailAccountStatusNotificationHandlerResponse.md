# MultiPartEmailAccountStatusNotificationHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Account Status Notification Handler | 
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
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewMultiPartEmailAccountStatusNotificationHandlerResponse

`func NewMultiPartEmailAccountStatusNotificationHandlerResponse(id string, schemas []EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn, enabled bool, ) *MultiPartEmailAccountStatusNotificationHandlerResponse`

NewMultiPartEmailAccountStatusNotificationHandlerResponse instantiates a new MultiPartEmailAccountStatusNotificationHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiPartEmailAccountStatusNotificationHandlerResponseWithDefaults

`func NewMultiPartEmailAccountStatusNotificationHandlerResponseWithDefaults() *MultiPartEmailAccountStatusNotificationHandlerResponse`

NewMultiPartEmailAccountStatusNotificationHandlerResponseWithDefaults instantiates a new MultiPartEmailAccountStatusNotificationHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetSchemas() []EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetSchemasOk() (*[]EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) SetSchemas(v []EnummultiPartEmailAccountStatusNotificationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAccountTemporarilyFailureLockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetAccountTemporarilyFailureLockedMessageTemplate() string`

GetAccountTemporarilyFailureLockedMessageTemplate returns the AccountTemporarilyFailureLockedMessageTemplate field if non-nil, zero value otherwise.

### GetAccountTemporarilyFailureLockedMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetAccountTemporarilyFailureLockedMessageTemplateOk() (*string, bool)`

GetAccountTemporarilyFailureLockedMessageTemplateOk returns a tuple with the AccountTemporarilyFailureLockedMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTemporarilyFailureLockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) SetAccountTemporarilyFailureLockedMessageTemplate(v string)`

SetAccountTemporarilyFailureLockedMessageTemplate sets AccountTemporarilyFailureLockedMessageTemplate field to given value.

### HasAccountTemporarilyFailureLockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) HasAccountTemporarilyFailureLockedMessageTemplate() bool`

HasAccountTemporarilyFailureLockedMessageTemplate returns a boolean if a field has been set.

### GetAccountPermanentlyFailureLockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetAccountPermanentlyFailureLockedMessageTemplate() string`

GetAccountPermanentlyFailureLockedMessageTemplate returns the AccountPermanentlyFailureLockedMessageTemplate field if non-nil, zero value otherwise.

### GetAccountPermanentlyFailureLockedMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetAccountPermanentlyFailureLockedMessageTemplateOk() (*string, bool)`

GetAccountPermanentlyFailureLockedMessageTemplateOk returns a tuple with the AccountPermanentlyFailureLockedMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPermanentlyFailureLockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) SetAccountPermanentlyFailureLockedMessageTemplate(v string)`

SetAccountPermanentlyFailureLockedMessageTemplate sets AccountPermanentlyFailureLockedMessageTemplate field to given value.

### HasAccountPermanentlyFailureLockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) HasAccountPermanentlyFailureLockedMessageTemplate() bool`

HasAccountPermanentlyFailureLockedMessageTemplate returns a boolean if a field has been set.

### GetAccountIdleLockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetAccountIdleLockedMessageTemplate() string`

GetAccountIdleLockedMessageTemplate returns the AccountIdleLockedMessageTemplate field if non-nil, zero value otherwise.

### GetAccountIdleLockedMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetAccountIdleLockedMessageTemplateOk() (*string, bool)`

GetAccountIdleLockedMessageTemplateOk returns a tuple with the AccountIdleLockedMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdleLockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) SetAccountIdleLockedMessageTemplate(v string)`

SetAccountIdleLockedMessageTemplate sets AccountIdleLockedMessageTemplate field to given value.

### HasAccountIdleLockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) HasAccountIdleLockedMessageTemplate() bool`

HasAccountIdleLockedMessageTemplate returns a boolean if a field has been set.

### GetAccountResetLockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetAccountResetLockedMessageTemplate() string`

GetAccountResetLockedMessageTemplate returns the AccountResetLockedMessageTemplate field if non-nil, zero value otherwise.

### GetAccountResetLockedMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetAccountResetLockedMessageTemplateOk() (*string, bool)`

GetAccountResetLockedMessageTemplateOk returns a tuple with the AccountResetLockedMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountResetLockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) SetAccountResetLockedMessageTemplate(v string)`

SetAccountResetLockedMessageTemplate sets AccountResetLockedMessageTemplate field to given value.

### HasAccountResetLockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) HasAccountResetLockedMessageTemplate() bool`

HasAccountResetLockedMessageTemplate returns a boolean if a field has been set.

### GetAccountUnlockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetAccountUnlockedMessageTemplate() string`

GetAccountUnlockedMessageTemplate returns the AccountUnlockedMessageTemplate field if non-nil, zero value otherwise.

### GetAccountUnlockedMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetAccountUnlockedMessageTemplateOk() (*string, bool)`

GetAccountUnlockedMessageTemplateOk returns a tuple with the AccountUnlockedMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUnlockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) SetAccountUnlockedMessageTemplate(v string)`

SetAccountUnlockedMessageTemplate sets AccountUnlockedMessageTemplate field to given value.

### HasAccountUnlockedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) HasAccountUnlockedMessageTemplate() bool`

HasAccountUnlockedMessageTemplate returns a boolean if a field has been set.

### GetAccountDisabledMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetAccountDisabledMessageTemplate() string`

GetAccountDisabledMessageTemplate returns the AccountDisabledMessageTemplate field if non-nil, zero value otherwise.

### GetAccountDisabledMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetAccountDisabledMessageTemplateOk() (*string, bool)`

GetAccountDisabledMessageTemplateOk returns a tuple with the AccountDisabledMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDisabledMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) SetAccountDisabledMessageTemplate(v string)`

SetAccountDisabledMessageTemplate sets AccountDisabledMessageTemplate field to given value.

### HasAccountDisabledMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) HasAccountDisabledMessageTemplate() bool`

HasAccountDisabledMessageTemplate returns a boolean if a field has been set.

### GetAccountEnabledMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetAccountEnabledMessageTemplate() string`

GetAccountEnabledMessageTemplate returns the AccountEnabledMessageTemplate field if non-nil, zero value otherwise.

### GetAccountEnabledMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetAccountEnabledMessageTemplateOk() (*string, bool)`

GetAccountEnabledMessageTemplateOk returns a tuple with the AccountEnabledMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountEnabledMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) SetAccountEnabledMessageTemplate(v string)`

SetAccountEnabledMessageTemplate sets AccountEnabledMessageTemplate field to given value.

### HasAccountEnabledMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) HasAccountEnabledMessageTemplate() bool`

HasAccountEnabledMessageTemplate returns a boolean if a field has been set.

### GetAccountNotYetActiveMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetAccountNotYetActiveMessageTemplate() string`

GetAccountNotYetActiveMessageTemplate returns the AccountNotYetActiveMessageTemplate field if non-nil, zero value otherwise.

### GetAccountNotYetActiveMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetAccountNotYetActiveMessageTemplateOk() (*string, bool)`

GetAccountNotYetActiveMessageTemplateOk returns a tuple with the AccountNotYetActiveMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNotYetActiveMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) SetAccountNotYetActiveMessageTemplate(v string)`

SetAccountNotYetActiveMessageTemplate sets AccountNotYetActiveMessageTemplate field to given value.

### HasAccountNotYetActiveMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) HasAccountNotYetActiveMessageTemplate() bool`

HasAccountNotYetActiveMessageTemplate returns a boolean if a field has been set.

### GetAccountExpiredMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetAccountExpiredMessageTemplate() string`

GetAccountExpiredMessageTemplate returns the AccountExpiredMessageTemplate field if non-nil, zero value otherwise.

### GetAccountExpiredMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetAccountExpiredMessageTemplateOk() (*string, bool)`

GetAccountExpiredMessageTemplateOk returns a tuple with the AccountExpiredMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountExpiredMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) SetAccountExpiredMessageTemplate(v string)`

SetAccountExpiredMessageTemplate sets AccountExpiredMessageTemplate field to given value.

### HasAccountExpiredMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) HasAccountExpiredMessageTemplate() bool`

HasAccountExpiredMessageTemplate returns a boolean if a field has been set.

### GetPasswordExpiredMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetPasswordExpiredMessageTemplate() string`

GetPasswordExpiredMessageTemplate returns the PasswordExpiredMessageTemplate field if non-nil, zero value otherwise.

### GetPasswordExpiredMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetPasswordExpiredMessageTemplateOk() (*string, bool)`

GetPasswordExpiredMessageTemplateOk returns a tuple with the PasswordExpiredMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpiredMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) SetPasswordExpiredMessageTemplate(v string)`

SetPasswordExpiredMessageTemplate sets PasswordExpiredMessageTemplate field to given value.

### HasPasswordExpiredMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) HasPasswordExpiredMessageTemplate() bool`

HasPasswordExpiredMessageTemplate returns a boolean if a field has been set.

### GetPasswordExpiringMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetPasswordExpiringMessageTemplate() string`

GetPasswordExpiringMessageTemplate returns the PasswordExpiringMessageTemplate field if non-nil, zero value otherwise.

### GetPasswordExpiringMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetPasswordExpiringMessageTemplateOk() (*string, bool)`

GetPasswordExpiringMessageTemplateOk returns a tuple with the PasswordExpiringMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpiringMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) SetPasswordExpiringMessageTemplate(v string)`

SetPasswordExpiringMessageTemplate sets PasswordExpiringMessageTemplate field to given value.

### HasPasswordExpiringMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) HasPasswordExpiringMessageTemplate() bool`

HasPasswordExpiringMessageTemplate returns a boolean if a field has been set.

### GetPasswordResetMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetPasswordResetMessageTemplate() string`

GetPasswordResetMessageTemplate returns the PasswordResetMessageTemplate field if non-nil, zero value otherwise.

### GetPasswordResetMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetPasswordResetMessageTemplateOk() (*string, bool)`

GetPasswordResetMessageTemplateOk returns a tuple with the PasswordResetMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordResetMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) SetPasswordResetMessageTemplate(v string)`

SetPasswordResetMessageTemplate sets PasswordResetMessageTemplate field to given value.

### HasPasswordResetMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) HasPasswordResetMessageTemplate() bool`

HasPasswordResetMessageTemplate returns a boolean if a field has been set.

### GetPasswordChangedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetPasswordChangedMessageTemplate() string`

GetPasswordChangedMessageTemplate returns the PasswordChangedMessageTemplate field if non-nil, zero value otherwise.

### GetPasswordChangedMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetPasswordChangedMessageTemplateOk() (*string, bool)`

GetPasswordChangedMessageTemplateOk returns a tuple with the PasswordChangedMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordChangedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) SetPasswordChangedMessageTemplate(v string)`

SetPasswordChangedMessageTemplate sets PasswordChangedMessageTemplate field to given value.

### HasPasswordChangedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) HasPasswordChangedMessageTemplate() bool`

HasPasswordChangedMessageTemplate returns a boolean if a field has been set.

### GetAccountCreatedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetAccountCreatedMessageTemplate() string`

GetAccountCreatedMessageTemplate returns the AccountCreatedMessageTemplate field if non-nil, zero value otherwise.

### GetAccountCreatedMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetAccountCreatedMessageTemplateOk() (*string, bool)`

GetAccountCreatedMessageTemplateOk returns a tuple with the AccountCreatedMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCreatedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) SetAccountCreatedMessageTemplate(v string)`

SetAccountCreatedMessageTemplate sets AccountCreatedMessageTemplate field to given value.

### HasAccountCreatedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) HasAccountCreatedMessageTemplate() bool`

HasAccountCreatedMessageTemplate returns a boolean if a field has been set.

### GetAccountUpdatedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetAccountUpdatedMessageTemplate() string`

GetAccountUpdatedMessageTemplate returns the AccountUpdatedMessageTemplate field if non-nil, zero value otherwise.

### GetAccountUpdatedMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetAccountUpdatedMessageTemplateOk() (*string, bool)`

GetAccountUpdatedMessageTemplateOk returns a tuple with the AccountUpdatedMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUpdatedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) SetAccountUpdatedMessageTemplate(v string)`

SetAccountUpdatedMessageTemplate sets AccountUpdatedMessageTemplate field to given value.

### HasAccountUpdatedMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) HasAccountUpdatedMessageTemplate() bool`

HasAccountUpdatedMessageTemplate returns a boolean if a field has been set.

### GetBindPasswordFailedValidationMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetBindPasswordFailedValidationMessageTemplate() string`

GetBindPasswordFailedValidationMessageTemplate returns the BindPasswordFailedValidationMessageTemplate field if non-nil, zero value otherwise.

### GetBindPasswordFailedValidationMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetBindPasswordFailedValidationMessageTemplateOk() (*string, bool)`

GetBindPasswordFailedValidationMessageTemplateOk returns a tuple with the BindPasswordFailedValidationMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPasswordFailedValidationMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) SetBindPasswordFailedValidationMessageTemplate(v string)`

SetBindPasswordFailedValidationMessageTemplate sets BindPasswordFailedValidationMessageTemplate field to given value.

### HasBindPasswordFailedValidationMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) HasBindPasswordFailedValidationMessageTemplate() bool`

HasBindPasswordFailedValidationMessageTemplate returns a boolean if a field has been set.

### GetMustChangePasswordMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetMustChangePasswordMessageTemplate() string`

GetMustChangePasswordMessageTemplate returns the MustChangePasswordMessageTemplate field if non-nil, zero value otherwise.

### GetMustChangePasswordMessageTemplateOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetMustChangePasswordMessageTemplateOk() (*string, bool)`

GetMustChangePasswordMessageTemplateOk returns a tuple with the MustChangePasswordMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustChangePasswordMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) SetMustChangePasswordMessageTemplate(v string)`

SetMustChangePasswordMessageTemplate sets MustChangePasswordMessageTemplate field to given value.

### HasMustChangePasswordMessageTemplate

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) HasMustChangePasswordMessageTemplate() bool`

HasMustChangePasswordMessageTemplate returns a boolean if a field has been set.

### GetDescription

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAsynchronous

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetAsynchronous() bool`

GetAsynchronous returns the Asynchronous field if non-nil, zero value otherwise.

### GetAsynchronousOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetAsynchronousOk() (*bool, bool)`

GetAsynchronousOk returns a tuple with the Asynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsynchronous

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) SetAsynchronous(v bool)`

SetAsynchronous sets Asynchronous field to given value.

### HasAsynchronous

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) HasAsynchronous() bool`

HasAsynchronous returns a boolean if a field has been set.

### GetAccountCreationNotificationRequestCriteria

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetAccountCreationNotificationRequestCriteria() string`

GetAccountCreationNotificationRequestCriteria returns the AccountCreationNotificationRequestCriteria field if non-nil, zero value otherwise.

### GetAccountCreationNotificationRequestCriteriaOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetAccountCreationNotificationRequestCriteriaOk() (*string, bool)`

GetAccountCreationNotificationRequestCriteriaOk returns a tuple with the AccountCreationNotificationRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCreationNotificationRequestCriteria

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) SetAccountCreationNotificationRequestCriteria(v string)`

SetAccountCreationNotificationRequestCriteria sets AccountCreationNotificationRequestCriteria field to given value.

### HasAccountCreationNotificationRequestCriteria

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) HasAccountCreationNotificationRequestCriteria() bool`

HasAccountCreationNotificationRequestCriteria returns a boolean if a field has been set.

### GetAccountUpdateNotificationRequestCriteria

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetAccountUpdateNotificationRequestCriteria() string`

GetAccountUpdateNotificationRequestCriteria returns the AccountUpdateNotificationRequestCriteria field if non-nil, zero value otherwise.

### GetAccountUpdateNotificationRequestCriteriaOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetAccountUpdateNotificationRequestCriteriaOk() (*string, bool)`

GetAccountUpdateNotificationRequestCriteriaOk returns a tuple with the AccountUpdateNotificationRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUpdateNotificationRequestCriteria

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) SetAccountUpdateNotificationRequestCriteria(v string)`

SetAccountUpdateNotificationRequestCriteria sets AccountUpdateNotificationRequestCriteria field to given value.

### HasAccountUpdateNotificationRequestCriteria

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) HasAccountUpdateNotificationRequestCriteria() bool`

HasAccountUpdateNotificationRequestCriteria returns a boolean if a field has been set.

### GetMeta

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *MultiPartEmailAccountStatusNotificationHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


