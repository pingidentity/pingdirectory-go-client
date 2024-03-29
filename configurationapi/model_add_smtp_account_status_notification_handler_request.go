/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the AddSmtpAccountStatusNotificationHandlerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddSmtpAccountStatusNotificationHandlerRequest{}

// AddSmtpAccountStatusNotificationHandlerRequest struct for AddSmtpAccountStatusNotificationHandlerRequest
type AddSmtpAccountStatusNotificationHandlerRequest struct {
	Schemas []EnumsmtpAccountStatusNotificationHandlerSchemaUrn `json:"schemas"`
	// Specifies which attribute in the user's entries may be used to obtain the email address when notifying the end user.
	EmailAddressAttributeType []string `json:"emailAddressAttributeType,omitempty"`
	// The name of the JSON field whose value is the email address to which the message should be sent. The email address must be contained in a top-level field whose value is a single string.
	EmailAddressJSONField *string `json:"emailAddressJSONField,omitempty"`
	// A JSON object filter that may be used to identify which email address value to use when sending the message.
	EmailAddressJSONObjectFilter *string `json:"emailAddressJSONObjectFilter,omitempty"`
	// Specifies an email address to which notification messages are sent, either instead of or in addition to the end user for whom the notification has been generated.
	RecipientAddress []string `json:"recipientAddress,omitempty"`
	// Indicates whether an email notification message should be generated and sent to the set of notification recipients even if the user entry does not contain any values for any of the email address attributes (that is, in cases when it is not possible to notify the end user).
	SendMessageWithoutEndUserAddress *bool `json:"sendMessageWithoutEndUserAddress,omitempty"`
	// Specifies the email address from which the message is sent. Note that this does not necessarily have to be a legitimate email address.
	SenderAddress string `json:"senderAddress"`
	// Specifies the subject that should be used for email messages generated by this account status notification handler.
	MessageSubject []string `json:"messageSubject"`
	// Specifies the path to the file containing the message template to generate the email notification messages.
	MessageTemplateFile []string `json:"messageTemplateFile"`
	// A description for this Account Status Notification Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the Account Status Notification Handler is enabled. Only enabled handlers are invoked whenever a related event occurs in the server.
	Enabled bool `json:"enabled"`
	// Indicates whether the server should attempt to invoke this Account Status Notification Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver a message) will not delay processing for the operation that triggered the notification.
	Asynchronous *bool `json:"asynchronous,omitempty"`
	// A result criteria object that identifies which successful bind operations should result in account authentication notifications for this handler.
	AccountAuthenticationNotificationResultCriteria *string `json:"accountAuthenticationNotificationResultCriteria,omitempty"`
	// A request criteria object that identifies which add requests should result in account creation notifications for this handler.
	AccountCreationNotificationRequestCriteria *string `json:"accountCreationNotificationRequestCriteria,omitempty"`
	// A request criteria object that identifies which delete requests should result in account deletion notifications for this handler.
	AccountDeletionNotificationRequestCriteria *string `json:"accountDeletionNotificationRequestCriteria,omitempty"`
	// A request criteria object that identifies which modify and modify DN requests should result in account update notifications for this handler.
	AccountUpdateNotificationRequestCriteria *string `json:"accountUpdateNotificationRequestCriteria,omitempty"`
	// Name of the new Account Status Notification Handler
	HandlerName string `json:"handlerName"`
}

// NewAddSmtpAccountStatusNotificationHandlerRequest instantiates a new AddSmtpAccountStatusNotificationHandlerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSmtpAccountStatusNotificationHandlerRequest(schemas []EnumsmtpAccountStatusNotificationHandlerSchemaUrn, senderAddress string, messageSubject []string, messageTemplateFile []string, enabled bool, handlerName string) *AddSmtpAccountStatusNotificationHandlerRequest {
	this := AddSmtpAccountStatusNotificationHandlerRequest{}
	this.Schemas = schemas
	this.SenderAddress = senderAddress
	this.MessageSubject = messageSubject
	this.MessageTemplateFile = messageTemplateFile
	this.Enabled = enabled
	this.HandlerName = handlerName
	return &this
}

// NewAddSmtpAccountStatusNotificationHandlerRequestWithDefaults instantiates a new AddSmtpAccountStatusNotificationHandlerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSmtpAccountStatusNotificationHandlerRequestWithDefaults() *AddSmtpAccountStatusNotificationHandlerRequest {
	this := AddSmtpAccountStatusNotificationHandlerRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetSchemas() []EnumsmtpAccountStatusNotificationHandlerSchemaUrn {
	if o == nil {
		var ret []EnumsmtpAccountStatusNotificationHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetSchemasOk() ([]EnumsmtpAccountStatusNotificationHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddSmtpAccountStatusNotificationHandlerRequest) SetSchemas(v []EnumsmtpAccountStatusNotificationHandlerSchemaUrn) {
	o.Schemas = v
}

// GetEmailAddressAttributeType returns the EmailAddressAttributeType field value if set, zero value otherwise.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetEmailAddressAttributeType() []string {
	if o == nil || IsNil(o.EmailAddressAttributeType) {
		var ret []string
		return ret
	}
	return o.EmailAddressAttributeType
}

// GetEmailAddressAttributeTypeOk returns a tuple with the EmailAddressAttributeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetEmailAddressAttributeTypeOk() ([]string, bool) {
	if o == nil || IsNil(o.EmailAddressAttributeType) {
		return nil, false
	}
	return o.EmailAddressAttributeType, true
}

// HasEmailAddressAttributeType returns a boolean if a field has been set.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) HasEmailAddressAttributeType() bool {
	if o != nil && !IsNil(o.EmailAddressAttributeType) {
		return true
	}

	return false
}

// SetEmailAddressAttributeType gets a reference to the given []string and assigns it to the EmailAddressAttributeType field.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) SetEmailAddressAttributeType(v []string) {
	o.EmailAddressAttributeType = v
}

// GetEmailAddressJSONField returns the EmailAddressJSONField field value if set, zero value otherwise.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetEmailAddressJSONField() string {
	if o == nil || IsNil(o.EmailAddressJSONField) {
		var ret string
		return ret
	}
	return *o.EmailAddressJSONField
}

// GetEmailAddressJSONFieldOk returns a tuple with the EmailAddressJSONField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetEmailAddressJSONFieldOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAddressJSONField) {
		return nil, false
	}
	return o.EmailAddressJSONField, true
}

// HasEmailAddressJSONField returns a boolean if a field has been set.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) HasEmailAddressJSONField() bool {
	if o != nil && !IsNil(o.EmailAddressJSONField) {
		return true
	}

	return false
}

// SetEmailAddressJSONField gets a reference to the given string and assigns it to the EmailAddressJSONField field.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) SetEmailAddressJSONField(v string) {
	o.EmailAddressJSONField = &v
}

// GetEmailAddressJSONObjectFilter returns the EmailAddressJSONObjectFilter field value if set, zero value otherwise.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetEmailAddressJSONObjectFilter() string {
	if o == nil || IsNil(o.EmailAddressJSONObjectFilter) {
		var ret string
		return ret
	}
	return *o.EmailAddressJSONObjectFilter
}

// GetEmailAddressJSONObjectFilterOk returns a tuple with the EmailAddressJSONObjectFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetEmailAddressJSONObjectFilterOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAddressJSONObjectFilter) {
		return nil, false
	}
	return o.EmailAddressJSONObjectFilter, true
}

// HasEmailAddressJSONObjectFilter returns a boolean if a field has been set.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) HasEmailAddressJSONObjectFilter() bool {
	if o != nil && !IsNil(o.EmailAddressJSONObjectFilter) {
		return true
	}

	return false
}

// SetEmailAddressJSONObjectFilter gets a reference to the given string and assigns it to the EmailAddressJSONObjectFilter field.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) SetEmailAddressJSONObjectFilter(v string) {
	o.EmailAddressJSONObjectFilter = &v
}

// GetRecipientAddress returns the RecipientAddress field value if set, zero value otherwise.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetRecipientAddress() []string {
	if o == nil || IsNil(o.RecipientAddress) {
		var ret []string
		return ret
	}
	return o.RecipientAddress
}

// GetRecipientAddressOk returns a tuple with the RecipientAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetRecipientAddressOk() ([]string, bool) {
	if o == nil || IsNil(o.RecipientAddress) {
		return nil, false
	}
	return o.RecipientAddress, true
}

// HasRecipientAddress returns a boolean if a field has been set.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) HasRecipientAddress() bool {
	if o != nil && !IsNil(o.RecipientAddress) {
		return true
	}

	return false
}

// SetRecipientAddress gets a reference to the given []string and assigns it to the RecipientAddress field.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) SetRecipientAddress(v []string) {
	o.RecipientAddress = v
}

// GetSendMessageWithoutEndUserAddress returns the SendMessageWithoutEndUserAddress field value if set, zero value otherwise.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetSendMessageWithoutEndUserAddress() bool {
	if o == nil || IsNil(o.SendMessageWithoutEndUserAddress) {
		var ret bool
		return ret
	}
	return *o.SendMessageWithoutEndUserAddress
}

// GetSendMessageWithoutEndUserAddressOk returns a tuple with the SendMessageWithoutEndUserAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetSendMessageWithoutEndUserAddressOk() (*bool, bool) {
	if o == nil || IsNil(o.SendMessageWithoutEndUserAddress) {
		return nil, false
	}
	return o.SendMessageWithoutEndUserAddress, true
}

// HasSendMessageWithoutEndUserAddress returns a boolean if a field has been set.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) HasSendMessageWithoutEndUserAddress() bool {
	if o != nil && !IsNil(o.SendMessageWithoutEndUserAddress) {
		return true
	}

	return false
}

// SetSendMessageWithoutEndUserAddress gets a reference to the given bool and assigns it to the SendMessageWithoutEndUserAddress field.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) SetSendMessageWithoutEndUserAddress(v bool) {
	o.SendMessageWithoutEndUserAddress = &v
}

// GetSenderAddress returns the SenderAddress field value
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetSenderAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SenderAddress
}

// GetSenderAddressOk returns a tuple with the SenderAddress field value
// and a boolean to check if the value has been set.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetSenderAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderAddress, true
}

// SetSenderAddress sets field value
func (o *AddSmtpAccountStatusNotificationHandlerRequest) SetSenderAddress(v string) {
	o.SenderAddress = v
}

// GetMessageSubject returns the MessageSubject field value
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetMessageSubject() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MessageSubject
}

// GetMessageSubjectOk returns a tuple with the MessageSubject field value
// and a boolean to check if the value has been set.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetMessageSubjectOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MessageSubject, true
}

// SetMessageSubject sets field value
func (o *AddSmtpAccountStatusNotificationHandlerRequest) SetMessageSubject(v []string) {
	o.MessageSubject = v
}

// GetMessageTemplateFile returns the MessageTemplateFile field value
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetMessageTemplateFile() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MessageTemplateFile
}

// GetMessageTemplateFileOk returns a tuple with the MessageTemplateFile field value
// and a boolean to check if the value has been set.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetMessageTemplateFileOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MessageTemplateFile, true
}

// SetMessageTemplateFile sets field value
func (o *AddSmtpAccountStatusNotificationHandlerRequest) SetMessageTemplateFile(v []string) {
	o.MessageTemplateFile = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddSmtpAccountStatusNotificationHandlerRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAsynchronous returns the Asynchronous field value if set, zero value otherwise.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetAsynchronous() bool {
	if o == nil || IsNil(o.Asynchronous) {
		var ret bool
		return ret
	}
	return *o.Asynchronous
}

// GetAsynchronousOk returns a tuple with the Asynchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetAsynchronousOk() (*bool, bool) {
	if o == nil || IsNil(o.Asynchronous) {
		return nil, false
	}
	return o.Asynchronous, true
}

// HasAsynchronous returns a boolean if a field has been set.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) HasAsynchronous() bool {
	if o != nil && !IsNil(o.Asynchronous) {
		return true
	}

	return false
}

// SetAsynchronous gets a reference to the given bool and assigns it to the Asynchronous field.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) SetAsynchronous(v bool) {
	o.Asynchronous = &v
}

// GetAccountAuthenticationNotificationResultCriteria returns the AccountAuthenticationNotificationResultCriteria field value if set, zero value otherwise.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetAccountAuthenticationNotificationResultCriteria() string {
	if o == nil || IsNil(o.AccountAuthenticationNotificationResultCriteria) {
		var ret string
		return ret
	}
	return *o.AccountAuthenticationNotificationResultCriteria
}

// GetAccountAuthenticationNotificationResultCriteriaOk returns a tuple with the AccountAuthenticationNotificationResultCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetAccountAuthenticationNotificationResultCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.AccountAuthenticationNotificationResultCriteria) {
		return nil, false
	}
	return o.AccountAuthenticationNotificationResultCriteria, true
}

// HasAccountAuthenticationNotificationResultCriteria returns a boolean if a field has been set.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) HasAccountAuthenticationNotificationResultCriteria() bool {
	if o != nil && !IsNil(o.AccountAuthenticationNotificationResultCriteria) {
		return true
	}

	return false
}

// SetAccountAuthenticationNotificationResultCriteria gets a reference to the given string and assigns it to the AccountAuthenticationNotificationResultCriteria field.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) SetAccountAuthenticationNotificationResultCriteria(v string) {
	o.AccountAuthenticationNotificationResultCriteria = &v
}

// GetAccountCreationNotificationRequestCriteria returns the AccountCreationNotificationRequestCriteria field value if set, zero value otherwise.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetAccountCreationNotificationRequestCriteria() string {
	if o == nil || IsNil(o.AccountCreationNotificationRequestCriteria) {
		var ret string
		return ret
	}
	return *o.AccountCreationNotificationRequestCriteria
}

// GetAccountCreationNotificationRequestCriteriaOk returns a tuple with the AccountCreationNotificationRequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetAccountCreationNotificationRequestCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.AccountCreationNotificationRequestCriteria) {
		return nil, false
	}
	return o.AccountCreationNotificationRequestCriteria, true
}

// HasAccountCreationNotificationRequestCriteria returns a boolean if a field has been set.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) HasAccountCreationNotificationRequestCriteria() bool {
	if o != nil && !IsNil(o.AccountCreationNotificationRequestCriteria) {
		return true
	}

	return false
}

// SetAccountCreationNotificationRequestCriteria gets a reference to the given string and assigns it to the AccountCreationNotificationRequestCriteria field.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) SetAccountCreationNotificationRequestCriteria(v string) {
	o.AccountCreationNotificationRequestCriteria = &v
}

// GetAccountDeletionNotificationRequestCriteria returns the AccountDeletionNotificationRequestCriteria field value if set, zero value otherwise.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetAccountDeletionNotificationRequestCriteria() string {
	if o == nil || IsNil(o.AccountDeletionNotificationRequestCriteria) {
		var ret string
		return ret
	}
	return *o.AccountDeletionNotificationRequestCriteria
}

// GetAccountDeletionNotificationRequestCriteriaOk returns a tuple with the AccountDeletionNotificationRequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetAccountDeletionNotificationRequestCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.AccountDeletionNotificationRequestCriteria) {
		return nil, false
	}
	return o.AccountDeletionNotificationRequestCriteria, true
}

// HasAccountDeletionNotificationRequestCriteria returns a boolean if a field has been set.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) HasAccountDeletionNotificationRequestCriteria() bool {
	if o != nil && !IsNil(o.AccountDeletionNotificationRequestCriteria) {
		return true
	}

	return false
}

// SetAccountDeletionNotificationRequestCriteria gets a reference to the given string and assigns it to the AccountDeletionNotificationRequestCriteria field.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) SetAccountDeletionNotificationRequestCriteria(v string) {
	o.AccountDeletionNotificationRequestCriteria = &v
}

// GetAccountUpdateNotificationRequestCriteria returns the AccountUpdateNotificationRequestCriteria field value if set, zero value otherwise.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetAccountUpdateNotificationRequestCriteria() string {
	if o == nil || IsNil(o.AccountUpdateNotificationRequestCriteria) {
		var ret string
		return ret
	}
	return *o.AccountUpdateNotificationRequestCriteria
}

// GetAccountUpdateNotificationRequestCriteriaOk returns a tuple with the AccountUpdateNotificationRequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetAccountUpdateNotificationRequestCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.AccountUpdateNotificationRequestCriteria) {
		return nil, false
	}
	return o.AccountUpdateNotificationRequestCriteria, true
}

// HasAccountUpdateNotificationRequestCriteria returns a boolean if a field has been set.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) HasAccountUpdateNotificationRequestCriteria() bool {
	if o != nil && !IsNil(o.AccountUpdateNotificationRequestCriteria) {
		return true
	}

	return false
}

// SetAccountUpdateNotificationRequestCriteria gets a reference to the given string and assigns it to the AccountUpdateNotificationRequestCriteria field.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) SetAccountUpdateNotificationRequestCriteria(v string) {
	o.AccountUpdateNotificationRequestCriteria = &v
}

// GetHandlerName returns the HandlerName field value
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetHandlerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HandlerName
}

// GetHandlerNameOk returns a tuple with the HandlerName field value
// and a boolean to check if the value has been set.
func (o *AddSmtpAccountStatusNotificationHandlerRequest) GetHandlerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HandlerName, true
}

// SetHandlerName sets field value
func (o *AddSmtpAccountStatusNotificationHandlerRequest) SetHandlerName(v string) {
	o.HandlerName = v
}

func (o AddSmtpAccountStatusNotificationHandlerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddSmtpAccountStatusNotificationHandlerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.EmailAddressAttributeType) {
		toSerialize["emailAddressAttributeType"] = o.EmailAddressAttributeType
	}
	if !IsNil(o.EmailAddressJSONField) {
		toSerialize["emailAddressJSONField"] = o.EmailAddressJSONField
	}
	if !IsNil(o.EmailAddressJSONObjectFilter) {
		toSerialize["emailAddressJSONObjectFilter"] = o.EmailAddressJSONObjectFilter
	}
	if !IsNil(o.RecipientAddress) {
		toSerialize["recipientAddress"] = o.RecipientAddress
	}
	if !IsNil(o.SendMessageWithoutEndUserAddress) {
		toSerialize["sendMessageWithoutEndUserAddress"] = o.SendMessageWithoutEndUserAddress
	}
	toSerialize["senderAddress"] = o.SenderAddress
	toSerialize["messageSubject"] = o.MessageSubject
	toSerialize["messageTemplateFile"] = o.MessageTemplateFile
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Asynchronous) {
		toSerialize["asynchronous"] = o.Asynchronous
	}
	if !IsNil(o.AccountAuthenticationNotificationResultCriteria) {
		toSerialize["accountAuthenticationNotificationResultCriteria"] = o.AccountAuthenticationNotificationResultCriteria
	}
	if !IsNil(o.AccountCreationNotificationRequestCriteria) {
		toSerialize["accountCreationNotificationRequestCriteria"] = o.AccountCreationNotificationRequestCriteria
	}
	if !IsNil(o.AccountDeletionNotificationRequestCriteria) {
		toSerialize["accountDeletionNotificationRequestCriteria"] = o.AccountDeletionNotificationRequestCriteria
	}
	if !IsNil(o.AccountUpdateNotificationRequestCriteria) {
		toSerialize["accountUpdateNotificationRequestCriteria"] = o.AccountUpdateNotificationRequestCriteria
	}
	toSerialize["handlerName"] = o.HandlerName
	return toSerialize, nil
}

type NullableAddSmtpAccountStatusNotificationHandlerRequest struct {
	value *AddSmtpAccountStatusNotificationHandlerRequest
	isSet bool
}

func (v NullableAddSmtpAccountStatusNotificationHandlerRequest) Get() *AddSmtpAccountStatusNotificationHandlerRequest {
	return v.value
}

func (v *NullableAddSmtpAccountStatusNotificationHandlerRequest) Set(val *AddSmtpAccountStatusNotificationHandlerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSmtpAccountStatusNotificationHandlerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSmtpAccountStatusNotificationHandlerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSmtpAccountStatusNotificationHandlerRequest(val *AddSmtpAccountStatusNotificationHandlerRequest) *NullableAddSmtpAccountStatusNotificationHandlerRequest {
	return &NullableAddSmtpAccountStatusNotificationHandlerRequest{value: val, isSet: true}
}

func (v NullableAddSmtpAccountStatusNotificationHandlerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSmtpAccountStatusNotificationHandlerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
