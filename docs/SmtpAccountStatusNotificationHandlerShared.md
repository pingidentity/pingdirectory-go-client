# SmtpAccountStatusNotificationHandlerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsmtpAccountStatusNotificationHandlerSchemaUrn**](EnumsmtpAccountStatusNotificationHandlerSchemaUrn.md) |  | 
**EmailAddressAttributeType** | Pointer to **[]string** |  | [optional] 
**EmailAddressJSONField** | Pointer to **string** | The name of the JSON field whose value is the email address to which the message should be sent. The email address must be contained in a top-level field whose value is a single string. | [optional] 
**EmailAddressJSONObjectFilter** | Pointer to **string** | A JSON object filter that may be used to identify which email address value to use when sending the message. | [optional] 
**RecipientAddress** | Pointer to **[]string** |  | [optional] 
**SendMessageWithoutEndUserAddress** | **bool** | Indicates whether an email notification message should be generated and sent to the set of notification recipients even if the user entry does not contain any values for any of the email address attributes (that is, in cases when it is not possible to notify the end user). | 
**SenderAddress** | **string** | Specifies the email address from which the message is sent. Note that this does not necessarily have to be a legitimate email address. | 
**MessageSubject** | **[]string** |  | 
**MessageTemplateFile** | **[]string** |  | 
**Description** | Pointer to **string** | A description for this Account Status Notification Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Account Status Notification Handler is enabled. Only enabled handlers are invoked whenever a related event occurs in the server. | 
**Asynchronous** | Pointer to **bool** | Indicates whether the server should attempt to invoke this Account Status Notification Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver a message) will not delay processing for the operation that triggered the notification. | [optional] 
**AccountCreationNotificationRequestCriteria** | Pointer to **string** | A request criteria object that identifies which add requests should result in account creation notifications for this handler. | [optional] 
**AccountUpdateNotificationRequestCriteria** | Pointer to **string** | A request criteria object that identifies which modify and modify DN requests should result in account update notifications for this handler. | [optional] 

## Methods

### NewSmtpAccountStatusNotificationHandlerShared

`func NewSmtpAccountStatusNotificationHandlerShared(schemas []EnumsmtpAccountStatusNotificationHandlerSchemaUrn, sendMessageWithoutEndUserAddress bool, senderAddress string, messageSubject []string, messageTemplateFile []string, enabled bool, ) *SmtpAccountStatusNotificationHandlerShared`

NewSmtpAccountStatusNotificationHandlerShared instantiates a new SmtpAccountStatusNotificationHandlerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmtpAccountStatusNotificationHandlerSharedWithDefaults

`func NewSmtpAccountStatusNotificationHandlerSharedWithDefaults() *SmtpAccountStatusNotificationHandlerShared`

NewSmtpAccountStatusNotificationHandlerSharedWithDefaults instantiates a new SmtpAccountStatusNotificationHandlerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *SmtpAccountStatusNotificationHandlerShared) GetSchemas() []EnumsmtpAccountStatusNotificationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SmtpAccountStatusNotificationHandlerShared) GetSchemasOk() (*[]EnumsmtpAccountStatusNotificationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SmtpAccountStatusNotificationHandlerShared) SetSchemas(v []EnumsmtpAccountStatusNotificationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEmailAddressAttributeType

`func (o *SmtpAccountStatusNotificationHandlerShared) GetEmailAddressAttributeType() []string`

GetEmailAddressAttributeType returns the EmailAddressAttributeType field if non-nil, zero value otherwise.

### GetEmailAddressAttributeTypeOk

`func (o *SmtpAccountStatusNotificationHandlerShared) GetEmailAddressAttributeTypeOk() (*[]string, bool)`

GetEmailAddressAttributeTypeOk returns a tuple with the EmailAddressAttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddressAttributeType

`func (o *SmtpAccountStatusNotificationHandlerShared) SetEmailAddressAttributeType(v []string)`

SetEmailAddressAttributeType sets EmailAddressAttributeType field to given value.

### HasEmailAddressAttributeType

`func (o *SmtpAccountStatusNotificationHandlerShared) HasEmailAddressAttributeType() bool`

HasEmailAddressAttributeType returns a boolean if a field has been set.

### GetEmailAddressJSONField

`func (o *SmtpAccountStatusNotificationHandlerShared) GetEmailAddressJSONField() string`

GetEmailAddressJSONField returns the EmailAddressJSONField field if non-nil, zero value otherwise.

### GetEmailAddressJSONFieldOk

`func (o *SmtpAccountStatusNotificationHandlerShared) GetEmailAddressJSONFieldOk() (*string, bool)`

GetEmailAddressJSONFieldOk returns a tuple with the EmailAddressJSONField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddressJSONField

`func (o *SmtpAccountStatusNotificationHandlerShared) SetEmailAddressJSONField(v string)`

SetEmailAddressJSONField sets EmailAddressJSONField field to given value.

### HasEmailAddressJSONField

`func (o *SmtpAccountStatusNotificationHandlerShared) HasEmailAddressJSONField() bool`

HasEmailAddressJSONField returns a boolean if a field has been set.

### GetEmailAddressJSONObjectFilter

`func (o *SmtpAccountStatusNotificationHandlerShared) GetEmailAddressJSONObjectFilter() string`

GetEmailAddressJSONObjectFilter returns the EmailAddressJSONObjectFilter field if non-nil, zero value otherwise.

### GetEmailAddressJSONObjectFilterOk

`func (o *SmtpAccountStatusNotificationHandlerShared) GetEmailAddressJSONObjectFilterOk() (*string, bool)`

GetEmailAddressJSONObjectFilterOk returns a tuple with the EmailAddressJSONObjectFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddressJSONObjectFilter

`func (o *SmtpAccountStatusNotificationHandlerShared) SetEmailAddressJSONObjectFilter(v string)`

SetEmailAddressJSONObjectFilter sets EmailAddressJSONObjectFilter field to given value.

### HasEmailAddressJSONObjectFilter

`func (o *SmtpAccountStatusNotificationHandlerShared) HasEmailAddressJSONObjectFilter() bool`

HasEmailAddressJSONObjectFilter returns a boolean if a field has been set.

### GetRecipientAddress

`func (o *SmtpAccountStatusNotificationHandlerShared) GetRecipientAddress() []string`

GetRecipientAddress returns the RecipientAddress field if non-nil, zero value otherwise.

### GetRecipientAddressOk

`func (o *SmtpAccountStatusNotificationHandlerShared) GetRecipientAddressOk() (*[]string, bool)`

GetRecipientAddressOk returns a tuple with the RecipientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientAddress

`func (o *SmtpAccountStatusNotificationHandlerShared) SetRecipientAddress(v []string)`

SetRecipientAddress sets RecipientAddress field to given value.

### HasRecipientAddress

`func (o *SmtpAccountStatusNotificationHandlerShared) HasRecipientAddress() bool`

HasRecipientAddress returns a boolean if a field has been set.

### GetSendMessageWithoutEndUserAddress

`func (o *SmtpAccountStatusNotificationHandlerShared) GetSendMessageWithoutEndUserAddress() bool`

GetSendMessageWithoutEndUserAddress returns the SendMessageWithoutEndUserAddress field if non-nil, zero value otherwise.

### GetSendMessageWithoutEndUserAddressOk

`func (o *SmtpAccountStatusNotificationHandlerShared) GetSendMessageWithoutEndUserAddressOk() (*bool, bool)`

GetSendMessageWithoutEndUserAddressOk returns a tuple with the SendMessageWithoutEndUserAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendMessageWithoutEndUserAddress

`func (o *SmtpAccountStatusNotificationHandlerShared) SetSendMessageWithoutEndUserAddress(v bool)`

SetSendMessageWithoutEndUserAddress sets SendMessageWithoutEndUserAddress field to given value.


### GetSenderAddress

`func (o *SmtpAccountStatusNotificationHandlerShared) GetSenderAddress() string`

GetSenderAddress returns the SenderAddress field if non-nil, zero value otherwise.

### GetSenderAddressOk

`func (o *SmtpAccountStatusNotificationHandlerShared) GetSenderAddressOk() (*string, bool)`

GetSenderAddressOk returns a tuple with the SenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddress

`func (o *SmtpAccountStatusNotificationHandlerShared) SetSenderAddress(v string)`

SetSenderAddress sets SenderAddress field to given value.


### GetMessageSubject

`func (o *SmtpAccountStatusNotificationHandlerShared) GetMessageSubject() []string`

GetMessageSubject returns the MessageSubject field if non-nil, zero value otherwise.

### GetMessageSubjectOk

`func (o *SmtpAccountStatusNotificationHandlerShared) GetMessageSubjectOk() (*[]string, bool)`

GetMessageSubjectOk returns a tuple with the MessageSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSubject

`func (o *SmtpAccountStatusNotificationHandlerShared) SetMessageSubject(v []string)`

SetMessageSubject sets MessageSubject field to given value.


### GetMessageTemplateFile

`func (o *SmtpAccountStatusNotificationHandlerShared) GetMessageTemplateFile() []string`

GetMessageTemplateFile returns the MessageTemplateFile field if non-nil, zero value otherwise.

### GetMessageTemplateFileOk

`func (o *SmtpAccountStatusNotificationHandlerShared) GetMessageTemplateFileOk() (*[]string, bool)`

GetMessageTemplateFileOk returns a tuple with the MessageTemplateFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTemplateFile

`func (o *SmtpAccountStatusNotificationHandlerShared) SetMessageTemplateFile(v []string)`

SetMessageTemplateFile sets MessageTemplateFile field to given value.


### GetDescription

`func (o *SmtpAccountStatusNotificationHandlerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SmtpAccountStatusNotificationHandlerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SmtpAccountStatusNotificationHandlerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SmtpAccountStatusNotificationHandlerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SmtpAccountStatusNotificationHandlerShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SmtpAccountStatusNotificationHandlerShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SmtpAccountStatusNotificationHandlerShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAsynchronous

`func (o *SmtpAccountStatusNotificationHandlerShared) GetAsynchronous() bool`

GetAsynchronous returns the Asynchronous field if non-nil, zero value otherwise.

### GetAsynchronousOk

`func (o *SmtpAccountStatusNotificationHandlerShared) GetAsynchronousOk() (*bool, bool)`

GetAsynchronousOk returns a tuple with the Asynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsynchronous

`func (o *SmtpAccountStatusNotificationHandlerShared) SetAsynchronous(v bool)`

SetAsynchronous sets Asynchronous field to given value.

### HasAsynchronous

`func (o *SmtpAccountStatusNotificationHandlerShared) HasAsynchronous() bool`

HasAsynchronous returns a boolean if a field has been set.

### GetAccountCreationNotificationRequestCriteria

`func (o *SmtpAccountStatusNotificationHandlerShared) GetAccountCreationNotificationRequestCriteria() string`

GetAccountCreationNotificationRequestCriteria returns the AccountCreationNotificationRequestCriteria field if non-nil, zero value otherwise.

### GetAccountCreationNotificationRequestCriteriaOk

`func (o *SmtpAccountStatusNotificationHandlerShared) GetAccountCreationNotificationRequestCriteriaOk() (*string, bool)`

GetAccountCreationNotificationRequestCriteriaOk returns a tuple with the AccountCreationNotificationRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCreationNotificationRequestCriteria

`func (o *SmtpAccountStatusNotificationHandlerShared) SetAccountCreationNotificationRequestCriteria(v string)`

SetAccountCreationNotificationRequestCriteria sets AccountCreationNotificationRequestCriteria field to given value.

### HasAccountCreationNotificationRequestCriteria

`func (o *SmtpAccountStatusNotificationHandlerShared) HasAccountCreationNotificationRequestCriteria() bool`

HasAccountCreationNotificationRequestCriteria returns a boolean if a field has been set.

### GetAccountUpdateNotificationRequestCriteria

`func (o *SmtpAccountStatusNotificationHandlerShared) GetAccountUpdateNotificationRequestCriteria() string`

GetAccountUpdateNotificationRequestCriteria returns the AccountUpdateNotificationRequestCriteria field if non-nil, zero value otherwise.

### GetAccountUpdateNotificationRequestCriteriaOk

`func (o *SmtpAccountStatusNotificationHandlerShared) GetAccountUpdateNotificationRequestCriteriaOk() (*string, bool)`

GetAccountUpdateNotificationRequestCriteriaOk returns a tuple with the AccountUpdateNotificationRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUpdateNotificationRequestCriteria

`func (o *SmtpAccountStatusNotificationHandlerShared) SetAccountUpdateNotificationRequestCriteria(v string)`

SetAccountUpdateNotificationRequestCriteria sets AccountUpdateNotificationRequestCriteria field to given value.

### HasAccountUpdateNotificationRequestCriteria

`func (o *SmtpAccountStatusNotificationHandlerShared) HasAccountUpdateNotificationRequestCriteria() bool`

HasAccountUpdateNotificationRequestCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


