# EmailOtpDeliveryMechanismShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumemailOtpDeliveryMechanismSchemaUrn**](EnumemailOtpDeliveryMechanismSchemaUrn.md) |  | 
**EmailAddressAttributeType** | **string** | The name or OID of the attribute that holds the email address to which the message should be sent. | 
**EmailAddressJSONField** | Pointer to **string** | The name of the JSON field whose value is the email address to which the message should be sent. The email address must be contained in a top-level field whose value is a single string. | [optional] 
**EmailAddressJSONObjectFilter** | Pointer to **string** | A JSON object filter that may be used to identify which email address value to use when sending the message. | [optional] 
**SenderAddress** | **string** | The e-mail address to use as the sender for the one-time password. | 
**MessageSubject** | **string** | The subject to use for the e-mail message. | 
**MessageTextBeforeOTP** | Pointer to **string** | Any text that should appear in the message before the one-time password value. | [optional] 
**MessageTextAfterOTP** | Pointer to **string** | Any text that should appear in the message after the one-time password value. | [optional] 
**Description** | Pointer to **string** | A description for this OTP Delivery Mechanism | [optional] 
**Enabled** | **bool** | Indicates whether this OTP Delivery Mechanism is enabled for use in the server. | 

## Methods

### NewEmailOtpDeliveryMechanismShared

`func NewEmailOtpDeliveryMechanismShared(schemas []EnumemailOtpDeliveryMechanismSchemaUrn, emailAddressAttributeType string, senderAddress string, messageSubject string, enabled bool, ) *EmailOtpDeliveryMechanismShared`

NewEmailOtpDeliveryMechanismShared instantiates a new EmailOtpDeliveryMechanismShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailOtpDeliveryMechanismSharedWithDefaults

`func NewEmailOtpDeliveryMechanismSharedWithDefaults() *EmailOtpDeliveryMechanismShared`

NewEmailOtpDeliveryMechanismSharedWithDefaults instantiates a new EmailOtpDeliveryMechanismShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *EmailOtpDeliveryMechanismShared) GetSchemas() []EnumemailOtpDeliveryMechanismSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *EmailOtpDeliveryMechanismShared) GetSchemasOk() (*[]EnumemailOtpDeliveryMechanismSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *EmailOtpDeliveryMechanismShared) SetSchemas(v []EnumemailOtpDeliveryMechanismSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEmailAddressAttributeType

`func (o *EmailOtpDeliveryMechanismShared) GetEmailAddressAttributeType() string`

GetEmailAddressAttributeType returns the EmailAddressAttributeType field if non-nil, zero value otherwise.

### GetEmailAddressAttributeTypeOk

`func (o *EmailOtpDeliveryMechanismShared) GetEmailAddressAttributeTypeOk() (*string, bool)`

GetEmailAddressAttributeTypeOk returns a tuple with the EmailAddressAttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddressAttributeType

`func (o *EmailOtpDeliveryMechanismShared) SetEmailAddressAttributeType(v string)`

SetEmailAddressAttributeType sets EmailAddressAttributeType field to given value.


### GetEmailAddressJSONField

`func (o *EmailOtpDeliveryMechanismShared) GetEmailAddressJSONField() string`

GetEmailAddressJSONField returns the EmailAddressJSONField field if non-nil, zero value otherwise.

### GetEmailAddressJSONFieldOk

`func (o *EmailOtpDeliveryMechanismShared) GetEmailAddressJSONFieldOk() (*string, bool)`

GetEmailAddressJSONFieldOk returns a tuple with the EmailAddressJSONField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddressJSONField

`func (o *EmailOtpDeliveryMechanismShared) SetEmailAddressJSONField(v string)`

SetEmailAddressJSONField sets EmailAddressJSONField field to given value.

### HasEmailAddressJSONField

`func (o *EmailOtpDeliveryMechanismShared) HasEmailAddressJSONField() bool`

HasEmailAddressJSONField returns a boolean if a field has been set.

### GetEmailAddressJSONObjectFilter

`func (o *EmailOtpDeliveryMechanismShared) GetEmailAddressJSONObjectFilter() string`

GetEmailAddressJSONObjectFilter returns the EmailAddressJSONObjectFilter field if non-nil, zero value otherwise.

### GetEmailAddressJSONObjectFilterOk

`func (o *EmailOtpDeliveryMechanismShared) GetEmailAddressJSONObjectFilterOk() (*string, bool)`

GetEmailAddressJSONObjectFilterOk returns a tuple with the EmailAddressJSONObjectFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddressJSONObjectFilter

`func (o *EmailOtpDeliveryMechanismShared) SetEmailAddressJSONObjectFilter(v string)`

SetEmailAddressJSONObjectFilter sets EmailAddressJSONObjectFilter field to given value.

### HasEmailAddressJSONObjectFilter

`func (o *EmailOtpDeliveryMechanismShared) HasEmailAddressJSONObjectFilter() bool`

HasEmailAddressJSONObjectFilter returns a boolean if a field has been set.

### GetSenderAddress

`func (o *EmailOtpDeliveryMechanismShared) GetSenderAddress() string`

GetSenderAddress returns the SenderAddress field if non-nil, zero value otherwise.

### GetSenderAddressOk

`func (o *EmailOtpDeliveryMechanismShared) GetSenderAddressOk() (*string, bool)`

GetSenderAddressOk returns a tuple with the SenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddress

`func (o *EmailOtpDeliveryMechanismShared) SetSenderAddress(v string)`

SetSenderAddress sets SenderAddress field to given value.


### GetMessageSubject

`func (o *EmailOtpDeliveryMechanismShared) GetMessageSubject() string`

GetMessageSubject returns the MessageSubject field if non-nil, zero value otherwise.

### GetMessageSubjectOk

`func (o *EmailOtpDeliveryMechanismShared) GetMessageSubjectOk() (*string, bool)`

GetMessageSubjectOk returns a tuple with the MessageSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSubject

`func (o *EmailOtpDeliveryMechanismShared) SetMessageSubject(v string)`

SetMessageSubject sets MessageSubject field to given value.


### GetMessageTextBeforeOTP

`func (o *EmailOtpDeliveryMechanismShared) GetMessageTextBeforeOTP() string`

GetMessageTextBeforeOTP returns the MessageTextBeforeOTP field if non-nil, zero value otherwise.

### GetMessageTextBeforeOTPOk

`func (o *EmailOtpDeliveryMechanismShared) GetMessageTextBeforeOTPOk() (*string, bool)`

GetMessageTextBeforeOTPOk returns a tuple with the MessageTextBeforeOTP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTextBeforeOTP

`func (o *EmailOtpDeliveryMechanismShared) SetMessageTextBeforeOTP(v string)`

SetMessageTextBeforeOTP sets MessageTextBeforeOTP field to given value.

### HasMessageTextBeforeOTP

`func (o *EmailOtpDeliveryMechanismShared) HasMessageTextBeforeOTP() bool`

HasMessageTextBeforeOTP returns a boolean if a field has been set.

### GetMessageTextAfterOTP

`func (o *EmailOtpDeliveryMechanismShared) GetMessageTextAfterOTP() string`

GetMessageTextAfterOTP returns the MessageTextAfterOTP field if non-nil, zero value otherwise.

### GetMessageTextAfterOTPOk

`func (o *EmailOtpDeliveryMechanismShared) GetMessageTextAfterOTPOk() (*string, bool)`

GetMessageTextAfterOTPOk returns a tuple with the MessageTextAfterOTP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTextAfterOTP

`func (o *EmailOtpDeliveryMechanismShared) SetMessageTextAfterOTP(v string)`

SetMessageTextAfterOTP sets MessageTextAfterOTP field to given value.

### HasMessageTextAfterOTP

`func (o *EmailOtpDeliveryMechanismShared) HasMessageTextAfterOTP() bool`

HasMessageTextAfterOTP returns a boolean if a field has been set.

### GetDescription

`func (o *EmailOtpDeliveryMechanismShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EmailOtpDeliveryMechanismShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EmailOtpDeliveryMechanismShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EmailOtpDeliveryMechanismShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *EmailOtpDeliveryMechanismShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EmailOtpDeliveryMechanismShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EmailOtpDeliveryMechanismShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


