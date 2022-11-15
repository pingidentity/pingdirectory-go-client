# TwilioOtpDeliveryMechanismShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumtwilioOtpDeliveryMechanismSchemaUrn**](EnumtwilioOtpDeliveryMechanismSchemaUrn.md) |  | 
**TwilioAccountSID** | **string** | The unique identifier assigned to the Twilio account that will be used. | 
**TwilioAuthToken** | Pointer to **string** | The auth token for the Twilio account that will be used. | [optional] 
**TwilioAuthTokenPassphraseProvider** | Pointer to **string** | The passphrase provider that may be used to obtain the auth token for the Twilio account that will be used. | [optional] 
**PhoneNumberAttributeType** | **string** | The name or OID of the attribute in the user&#39;s entry that holds the phone number to which the message should be sent. | 
**PhoneNumberJSONField** | Pointer to **string** | The name of the JSON field whose value is the phone number to which the message should be sent. The phone number must be contained in a top-level field whose value is a single string. | [optional] 
**PhoneNumberJSONObjectFilter** | Pointer to **string** | A JSON object filter that may be used to identify which phone number value to use when sending the message. | [optional] 
**SenderPhoneNumber** | **[]string** |  | 
**MessageTextBeforeOTP** | Pointer to **string** | Any text that should appear in the message before the one-time password value. | [optional] 
**MessageTextAfterOTP** | Pointer to **string** | Any text that should appear in the message after the one-time password value. | [optional] 
**Description** | Pointer to **string** | A description for this OTP Delivery Mechanism | [optional] 
**Enabled** | **bool** | Indicates whether this OTP Delivery Mechanism is enabled for use in the server. | 

## Methods

### NewTwilioOtpDeliveryMechanismShared

`func NewTwilioOtpDeliveryMechanismShared(schemas []EnumtwilioOtpDeliveryMechanismSchemaUrn, twilioAccountSID string, phoneNumberAttributeType string, senderPhoneNumber []string, enabled bool, ) *TwilioOtpDeliveryMechanismShared`

NewTwilioOtpDeliveryMechanismShared instantiates a new TwilioOtpDeliveryMechanismShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTwilioOtpDeliveryMechanismSharedWithDefaults

`func NewTwilioOtpDeliveryMechanismSharedWithDefaults() *TwilioOtpDeliveryMechanismShared`

NewTwilioOtpDeliveryMechanismSharedWithDefaults instantiates a new TwilioOtpDeliveryMechanismShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *TwilioOtpDeliveryMechanismShared) GetSchemas() []EnumtwilioOtpDeliveryMechanismSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *TwilioOtpDeliveryMechanismShared) GetSchemasOk() (*[]EnumtwilioOtpDeliveryMechanismSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *TwilioOtpDeliveryMechanismShared) SetSchemas(v []EnumtwilioOtpDeliveryMechanismSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetTwilioAccountSID

`func (o *TwilioOtpDeliveryMechanismShared) GetTwilioAccountSID() string`

GetTwilioAccountSID returns the TwilioAccountSID field if non-nil, zero value otherwise.

### GetTwilioAccountSIDOk

`func (o *TwilioOtpDeliveryMechanismShared) GetTwilioAccountSIDOk() (*string, bool)`

GetTwilioAccountSIDOk returns a tuple with the TwilioAccountSID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAccountSID

`func (o *TwilioOtpDeliveryMechanismShared) SetTwilioAccountSID(v string)`

SetTwilioAccountSID sets TwilioAccountSID field to given value.


### GetTwilioAuthToken

`func (o *TwilioOtpDeliveryMechanismShared) GetTwilioAuthToken() string`

GetTwilioAuthToken returns the TwilioAuthToken field if non-nil, zero value otherwise.

### GetTwilioAuthTokenOk

`func (o *TwilioOtpDeliveryMechanismShared) GetTwilioAuthTokenOk() (*string, bool)`

GetTwilioAuthTokenOk returns a tuple with the TwilioAuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAuthToken

`func (o *TwilioOtpDeliveryMechanismShared) SetTwilioAuthToken(v string)`

SetTwilioAuthToken sets TwilioAuthToken field to given value.

### HasTwilioAuthToken

`func (o *TwilioOtpDeliveryMechanismShared) HasTwilioAuthToken() bool`

HasTwilioAuthToken returns a boolean if a field has been set.

### GetTwilioAuthTokenPassphraseProvider

`func (o *TwilioOtpDeliveryMechanismShared) GetTwilioAuthTokenPassphraseProvider() string`

GetTwilioAuthTokenPassphraseProvider returns the TwilioAuthTokenPassphraseProvider field if non-nil, zero value otherwise.

### GetTwilioAuthTokenPassphraseProviderOk

`func (o *TwilioOtpDeliveryMechanismShared) GetTwilioAuthTokenPassphraseProviderOk() (*string, bool)`

GetTwilioAuthTokenPassphraseProviderOk returns a tuple with the TwilioAuthTokenPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAuthTokenPassphraseProvider

`func (o *TwilioOtpDeliveryMechanismShared) SetTwilioAuthTokenPassphraseProvider(v string)`

SetTwilioAuthTokenPassphraseProvider sets TwilioAuthTokenPassphraseProvider field to given value.

### HasTwilioAuthTokenPassphraseProvider

`func (o *TwilioOtpDeliveryMechanismShared) HasTwilioAuthTokenPassphraseProvider() bool`

HasTwilioAuthTokenPassphraseProvider returns a boolean if a field has been set.

### GetPhoneNumberAttributeType

`func (o *TwilioOtpDeliveryMechanismShared) GetPhoneNumberAttributeType() string`

GetPhoneNumberAttributeType returns the PhoneNumberAttributeType field if non-nil, zero value otherwise.

### GetPhoneNumberAttributeTypeOk

`func (o *TwilioOtpDeliveryMechanismShared) GetPhoneNumberAttributeTypeOk() (*string, bool)`

GetPhoneNumberAttributeTypeOk returns a tuple with the PhoneNumberAttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberAttributeType

`func (o *TwilioOtpDeliveryMechanismShared) SetPhoneNumberAttributeType(v string)`

SetPhoneNumberAttributeType sets PhoneNumberAttributeType field to given value.


### GetPhoneNumberJSONField

`func (o *TwilioOtpDeliveryMechanismShared) GetPhoneNumberJSONField() string`

GetPhoneNumberJSONField returns the PhoneNumberJSONField field if non-nil, zero value otherwise.

### GetPhoneNumberJSONFieldOk

`func (o *TwilioOtpDeliveryMechanismShared) GetPhoneNumberJSONFieldOk() (*string, bool)`

GetPhoneNumberJSONFieldOk returns a tuple with the PhoneNumberJSONField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberJSONField

`func (o *TwilioOtpDeliveryMechanismShared) SetPhoneNumberJSONField(v string)`

SetPhoneNumberJSONField sets PhoneNumberJSONField field to given value.

### HasPhoneNumberJSONField

`func (o *TwilioOtpDeliveryMechanismShared) HasPhoneNumberJSONField() bool`

HasPhoneNumberJSONField returns a boolean if a field has been set.

### GetPhoneNumberJSONObjectFilter

`func (o *TwilioOtpDeliveryMechanismShared) GetPhoneNumberJSONObjectFilter() string`

GetPhoneNumberJSONObjectFilter returns the PhoneNumberJSONObjectFilter field if non-nil, zero value otherwise.

### GetPhoneNumberJSONObjectFilterOk

`func (o *TwilioOtpDeliveryMechanismShared) GetPhoneNumberJSONObjectFilterOk() (*string, bool)`

GetPhoneNumberJSONObjectFilterOk returns a tuple with the PhoneNumberJSONObjectFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberJSONObjectFilter

`func (o *TwilioOtpDeliveryMechanismShared) SetPhoneNumberJSONObjectFilter(v string)`

SetPhoneNumberJSONObjectFilter sets PhoneNumberJSONObjectFilter field to given value.

### HasPhoneNumberJSONObjectFilter

`func (o *TwilioOtpDeliveryMechanismShared) HasPhoneNumberJSONObjectFilter() bool`

HasPhoneNumberJSONObjectFilter returns a boolean if a field has been set.

### GetSenderPhoneNumber

`func (o *TwilioOtpDeliveryMechanismShared) GetSenderPhoneNumber() []string`

GetSenderPhoneNumber returns the SenderPhoneNumber field if non-nil, zero value otherwise.

### GetSenderPhoneNumberOk

`func (o *TwilioOtpDeliveryMechanismShared) GetSenderPhoneNumberOk() (*[]string, bool)`

GetSenderPhoneNumberOk returns a tuple with the SenderPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderPhoneNumber

`func (o *TwilioOtpDeliveryMechanismShared) SetSenderPhoneNumber(v []string)`

SetSenderPhoneNumber sets SenderPhoneNumber field to given value.


### GetMessageTextBeforeOTP

`func (o *TwilioOtpDeliveryMechanismShared) GetMessageTextBeforeOTP() string`

GetMessageTextBeforeOTP returns the MessageTextBeforeOTP field if non-nil, zero value otherwise.

### GetMessageTextBeforeOTPOk

`func (o *TwilioOtpDeliveryMechanismShared) GetMessageTextBeforeOTPOk() (*string, bool)`

GetMessageTextBeforeOTPOk returns a tuple with the MessageTextBeforeOTP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTextBeforeOTP

`func (o *TwilioOtpDeliveryMechanismShared) SetMessageTextBeforeOTP(v string)`

SetMessageTextBeforeOTP sets MessageTextBeforeOTP field to given value.

### HasMessageTextBeforeOTP

`func (o *TwilioOtpDeliveryMechanismShared) HasMessageTextBeforeOTP() bool`

HasMessageTextBeforeOTP returns a boolean if a field has been set.

### GetMessageTextAfterOTP

`func (o *TwilioOtpDeliveryMechanismShared) GetMessageTextAfterOTP() string`

GetMessageTextAfterOTP returns the MessageTextAfterOTP field if non-nil, zero value otherwise.

### GetMessageTextAfterOTPOk

`func (o *TwilioOtpDeliveryMechanismShared) GetMessageTextAfterOTPOk() (*string, bool)`

GetMessageTextAfterOTPOk returns a tuple with the MessageTextAfterOTP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTextAfterOTP

`func (o *TwilioOtpDeliveryMechanismShared) SetMessageTextAfterOTP(v string)`

SetMessageTextAfterOTP sets MessageTextAfterOTP field to given value.

### HasMessageTextAfterOTP

`func (o *TwilioOtpDeliveryMechanismShared) HasMessageTextAfterOTP() bool`

HasMessageTextAfterOTP returns a boolean if a field has been set.

### GetDescription

`func (o *TwilioOtpDeliveryMechanismShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TwilioOtpDeliveryMechanismShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TwilioOtpDeliveryMechanismShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TwilioOtpDeliveryMechanismShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *TwilioOtpDeliveryMechanismShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TwilioOtpDeliveryMechanismShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TwilioOtpDeliveryMechanismShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


