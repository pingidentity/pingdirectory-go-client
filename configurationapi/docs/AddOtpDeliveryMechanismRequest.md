# AddOtpDeliveryMechanismRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MechanismName** | **string** | Name of the new OTP Delivery Mechanism | 
**Schemas** | [**[]EnumthirdPartyOtpDeliveryMechanismSchemaUrn**](EnumthirdPartyOtpDeliveryMechanismSchemaUrn.md) |  | 
**HttpProxyExternalServer** | Pointer to **string** | A reference to an HTTP proxy server that should be used for requests sent to the Twilio service. | [optional] 
**TwilioAccountSID** | **string** | The unique identifier assigned to the Twilio account that will be used. | 
**TwilioAuthToken** | Pointer to **string** | The auth token for the Twilio account that will be used. | [optional] 
**TwilioAuthTokenPassphraseProvider** | Pointer to **string** | The passphrase provider that may be used to obtain the auth token for the Twilio account that will be used. | [optional] 
**PhoneNumberAttributeType** | Pointer to **string** | The name or OID of the attribute in the user&#39;s entry that holds the phone number to which the message should be sent. | [optional] 
**PhoneNumberJSONField** | Pointer to **string** | The name of the JSON field whose value is the phone number to which the message should be sent. The phone number must be contained in a top-level field whose value is a single string. | [optional] 
**PhoneNumberJSONObjectFilter** | Pointer to **string** | A JSON object filter that may be used to identify which phone number value to use when sending the message. | [optional] 
**SenderPhoneNumber** | **[]string** | The outgoing phone number to use for the messages. Values must be phone numbers you have obtained for use with your Twilio account. | 
**MessageTextBeforeOTP** | Pointer to **string** | Any text that should appear in the message before the one-time password value. | [optional] 
**MessageTextAfterOTP** | Pointer to **string** | Any text that should appear in the message after the one-time password value. | [optional] 
**Description** | Pointer to **string** | A description for this OTP Delivery Mechanism | [optional] 
**Enabled** | **bool** | Indicates whether this OTP Delivery Mechanism is enabled for use in the server. | 
**EmailAddressAttributeType** | Pointer to **string** | The name or OID of the attribute that holds the email address to which the message should be sent. | [optional] 
**EmailAddressJSONField** | Pointer to **string** | The name of the JSON field whose value is the email address to which the message should be sent. The email address must be contained in a top-level field whose value is a single string. | [optional] 
**EmailAddressJSONObjectFilter** | Pointer to **string** | A JSON object filter that may be used to identify which email address value to use when sending the message. | [optional] 
**SenderAddress** | **string** | The e-mail address to use as the sender for the one-time password. | 
**MessageSubject** | Pointer to **string** | The subject to use for the e-mail message. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party OTP Delivery Mechanism. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party OTP Delivery Mechanism. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddOtpDeliveryMechanismRequest

`func NewAddOtpDeliveryMechanismRequest(mechanismName string, schemas []EnumthirdPartyOtpDeliveryMechanismSchemaUrn, twilioAccountSID string, senderPhoneNumber []string, enabled bool, senderAddress string, extensionClass string, ) *AddOtpDeliveryMechanismRequest`

NewAddOtpDeliveryMechanismRequest instantiates a new AddOtpDeliveryMechanismRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddOtpDeliveryMechanismRequestWithDefaults

`func NewAddOtpDeliveryMechanismRequestWithDefaults() *AddOtpDeliveryMechanismRequest`

NewAddOtpDeliveryMechanismRequestWithDefaults instantiates a new AddOtpDeliveryMechanismRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMechanismName

`func (o *AddOtpDeliveryMechanismRequest) GetMechanismName() string`

GetMechanismName returns the MechanismName field if non-nil, zero value otherwise.

### GetMechanismNameOk

`func (o *AddOtpDeliveryMechanismRequest) GetMechanismNameOk() (*string, bool)`

GetMechanismNameOk returns a tuple with the MechanismName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMechanismName

`func (o *AddOtpDeliveryMechanismRequest) SetMechanismName(v string)`

SetMechanismName sets MechanismName field to given value.


### GetSchemas

`func (o *AddOtpDeliveryMechanismRequest) GetSchemas() []EnumthirdPartyOtpDeliveryMechanismSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddOtpDeliveryMechanismRequest) GetSchemasOk() (*[]EnumthirdPartyOtpDeliveryMechanismSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddOtpDeliveryMechanismRequest) SetSchemas(v []EnumthirdPartyOtpDeliveryMechanismSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetHttpProxyExternalServer

`func (o *AddOtpDeliveryMechanismRequest) GetHttpProxyExternalServer() string`

GetHttpProxyExternalServer returns the HttpProxyExternalServer field if non-nil, zero value otherwise.

### GetHttpProxyExternalServerOk

`func (o *AddOtpDeliveryMechanismRequest) GetHttpProxyExternalServerOk() (*string, bool)`

GetHttpProxyExternalServerOk returns a tuple with the HttpProxyExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyExternalServer

`func (o *AddOtpDeliveryMechanismRequest) SetHttpProxyExternalServer(v string)`

SetHttpProxyExternalServer sets HttpProxyExternalServer field to given value.

### HasHttpProxyExternalServer

`func (o *AddOtpDeliveryMechanismRequest) HasHttpProxyExternalServer() bool`

HasHttpProxyExternalServer returns a boolean if a field has been set.

### GetTwilioAccountSID

`func (o *AddOtpDeliveryMechanismRequest) GetTwilioAccountSID() string`

GetTwilioAccountSID returns the TwilioAccountSID field if non-nil, zero value otherwise.

### GetTwilioAccountSIDOk

`func (o *AddOtpDeliveryMechanismRequest) GetTwilioAccountSIDOk() (*string, bool)`

GetTwilioAccountSIDOk returns a tuple with the TwilioAccountSID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAccountSID

`func (o *AddOtpDeliveryMechanismRequest) SetTwilioAccountSID(v string)`

SetTwilioAccountSID sets TwilioAccountSID field to given value.


### GetTwilioAuthToken

`func (o *AddOtpDeliveryMechanismRequest) GetTwilioAuthToken() string`

GetTwilioAuthToken returns the TwilioAuthToken field if non-nil, zero value otherwise.

### GetTwilioAuthTokenOk

`func (o *AddOtpDeliveryMechanismRequest) GetTwilioAuthTokenOk() (*string, bool)`

GetTwilioAuthTokenOk returns a tuple with the TwilioAuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAuthToken

`func (o *AddOtpDeliveryMechanismRequest) SetTwilioAuthToken(v string)`

SetTwilioAuthToken sets TwilioAuthToken field to given value.

### HasTwilioAuthToken

`func (o *AddOtpDeliveryMechanismRequest) HasTwilioAuthToken() bool`

HasTwilioAuthToken returns a boolean if a field has been set.

### GetTwilioAuthTokenPassphraseProvider

`func (o *AddOtpDeliveryMechanismRequest) GetTwilioAuthTokenPassphraseProvider() string`

GetTwilioAuthTokenPassphraseProvider returns the TwilioAuthTokenPassphraseProvider field if non-nil, zero value otherwise.

### GetTwilioAuthTokenPassphraseProviderOk

`func (o *AddOtpDeliveryMechanismRequest) GetTwilioAuthTokenPassphraseProviderOk() (*string, bool)`

GetTwilioAuthTokenPassphraseProviderOk returns a tuple with the TwilioAuthTokenPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAuthTokenPassphraseProvider

`func (o *AddOtpDeliveryMechanismRequest) SetTwilioAuthTokenPassphraseProvider(v string)`

SetTwilioAuthTokenPassphraseProvider sets TwilioAuthTokenPassphraseProvider field to given value.

### HasTwilioAuthTokenPassphraseProvider

`func (o *AddOtpDeliveryMechanismRequest) HasTwilioAuthTokenPassphraseProvider() bool`

HasTwilioAuthTokenPassphraseProvider returns a boolean if a field has been set.

### GetPhoneNumberAttributeType

`func (o *AddOtpDeliveryMechanismRequest) GetPhoneNumberAttributeType() string`

GetPhoneNumberAttributeType returns the PhoneNumberAttributeType field if non-nil, zero value otherwise.

### GetPhoneNumberAttributeTypeOk

`func (o *AddOtpDeliveryMechanismRequest) GetPhoneNumberAttributeTypeOk() (*string, bool)`

GetPhoneNumberAttributeTypeOk returns a tuple with the PhoneNumberAttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberAttributeType

`func (o *AddOtpDeliveryMechanismRequest) SetPhoneNumberAttributeType(v string)`

SetPhoneNumberAttributeType sets PhoneNumberAttributeType field to given value.

### HasPhoneNumberAttributeType

`func (o *AddOtpDeliveryMechanismRequest) HasPhoneNumberAttributeType() bool`

HasPhoneNumberAttributeType returns a boolean if a field has been set.

### GetPhoneNumberJSONField

`func (o *AddOtpDeliveryMechanismRequest) GetPhoneNumberJSONField() string`

GetPhoneNumberJSONField returns the PhoneNumberJSONField field if non-nil, zero value otherwise.

### GetPhoneNumberJSONFieldOk

`func (o *AddOtpDeliveryMechanismRequest) GetPhoneNumberJSONFieldOk() (*string, bool)`

GetPhoneNumberJSONFieldOk returns a tuple with the PhoneNumberJSONField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberJSONField

`func (o *AddOtpDeliveryMechanismRequest) SetPhoneNumberJSONField(v string)`

SetPhoneNumberJSONField sets PhoneNumberJSONField field to given value.

### HasPhoneNumberJSONField

`func (o *AddOtpDeliveryMechanismRequest) HasPhoneNumberJSONField() bool`

HasPhoneNumberJSONField returns a boolean if a field has been set.

### GetPhoneNumberJSONObjectFilter

`func (o *AddOtpDeliveryMechanismRequest) GetPhoneNumberJSONObjectFilter() string`

GetPhoneNumberJSONObjectFilter returns the PhoneNumberJSONObjectFilter field if non-nil, zero value otherwise.

### GetPhoneNumberJSONObjectFilterOk

`func (o *AddOtpDeliveryMechanismRequest) GetPhoneNumberJSONObjectFilterOk() (*string, bool)`

GetPhoneNumberJSONObjectFilterOk returns a tuple with the PhoneNumberJSONObjectFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberJSONObjectFilter

`func (o *AddOtpDeliveryMechanismRequest) SetPhoneNumberJSONObjectFilter(v string)`

SetPhoneNumberJSONObjectFilter sets PhoneNumberJSONObjectFilter field to given value.

### HasPhoneNumberJSONObjectFilter

`func (o *AddOtpDeliveryMechanismRequest) HasPhoneNumberJSONObjectFilter() bool`

HasPhoneNumberJSONObjectFilter returns a boolean if a field has been set.

### GetSenderPhoneNumber

`func (o *AddOtpDeliveryMechanismRequest) GetSenderPhoneNumber() []string`

GetSenderPhoneNumber returns the SenderPhoneNumber field if non-nil, zero value otherwise.

### GetSenderPhoneNumberOk

`func (o *AddOtpDeliveryMechanismRequest) GetSenderPhoneNumberOk() (*[]string, bool)`

GetSenderPhoneNumberOk returns a tuple with the SenderPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderPhoneNumber

`func (o *AddOtpDeliveryMechanismRequest) SetSenderPhoneNumber(v []string)`

SetSenderPhoneNumber sets SenderPhoneNumber field to given value.


### GetMessageTextBeforeOTP

`func (o *AddOtpDeliveryMechanismRequest) GetMessageTextBeforeOTP() string`

GetMessageTextBeforeOTP returns the MessageTextBeforeOTP field if non-nil, zero value otherwise.

### GetMessageTextBeforeOTPOk

`func (o *AddOtpDeliveryMechanismRequest) GetMessageTextBeforeOTPOk() (*string, bool)`

GetMessageTextBeforeOTPOk returns a tuple with the MessageTextBeforeOTP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTextBeforeOTP

`func (o *AddOtpDeliveryMechanismRequest) SetMessageTextBeforeOTP(v string)`

SetMessageTextBeforeOTP sets MessageTextBeforeOTP field to given value.

### HasMessageTextBeforeOTP

`func (o *AddOtpDeliveryMechanismRequest) HasMessageTextBeforeOTP() bool`

HasMessageTextBeforeOTP returns a boolean if a field has been set.

### GetMessageTextAfterOTP

`func (o *AddOtpDeliveryMechanismRequest) GetMessageTextAfterOTP() string`

GetMessageTextAfterOTP returns the MessageTextAfterOTP field if non-nil, zero value otherwise.

### GetMessageTextAfterOTPOk

`func (o *AddOtpDeliveryMechanismRequest) GetMessageTextAfterOTPOk() (*string, bool)`

GetMessageTextAfterOTPOk returns a tuple with the MessageTextAfterOTP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTextAfterOTP

`func (o *AddOtpDeliveryMechanismRequest) SetMessageTextAfterOTP(v string)`

SetMessageTextAfterOTP sets MessageTextAfterOTP field to given value.

### HasMessageTextAfterOTP

`func (o *AddOtpDeliveryMechanismRequest) HasMessageTextAfterOTP() bool`

HasMessageTextAfterOTP returns a boolean if a field has been set.

### GetDescription

`func (o *AddOtpDeliveryMechanismRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddOtpDeliveryMechanismRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddOtpDeliveryMechanismRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddOtpDeliveryMechanismRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddOtpDeliveryMechanismRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddOtpDeliveryMechanismRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddOtpDeliveryMechanismRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEmailAddressAttributeType

`func (o *AddOtpDeliveryMechanismRequest) GetEmailAddressAttributeType() string`

GetEmailAddressAttributeType returns the EmailAddressAttributeType field if non-nil, zero value otherwise.

### GetEmailAddressAttributeTypeOk

`func (o *AddOtpDeliveryMechanismRequest) GetEmailAddressAttributeTypeOk() (*string, bool)`

GetEmailAddressAttributeTypeOk returns a tuple with the EmailAddressAttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddressAttributeType

`func (o *AddOtpDeliveryMechanismRequest) SetEmailAddressAttributeType(v string)`

SetEmailAddressAttributeType sets EmailAddressAttributeType field to given value.

### HasEmailAddressAttributeType

`func (o *AddOtpDeliveryMechanismRequest) HasEmailAddressAttributeType() bool`

HasEmailAddressAttributeType returns a boolean if a field has been set.

### GetEmailAddressJSONField

`func (o *AddOtpDeliveryMechanismRequest) GetEmailAddressJSONField() string`

GetEmailAddressJSONField returns the EmailAddressJSONField field if non-nil, zero value otherwise.

### GetEmailAddressJSONFieldOk

`func (o *AddOtpDeliveryMechanismRequest) GetEmailAddressJSONFieldOk() (*string, bool)`

GetEmailAddressJSONFieldOk returns a tuple with the EmailAddressJSONField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddressJSONField

`func (o *AddOtpDeliveryMechanismRequest) SetEmailAddressJSONField(v string)`

SetEmailAddressJSONField sets EmailAddressJSONField field to given value.

### HasEmailAddressJSONField

`func (o *AddOtpDeliveryMechanismRequest) HasEmailAddressJSONField() bool`

HasEmailAddressJSONField returns a boolean if a field has been set.

### GetEmailAddressJSONObjectFilter

`func (o *AddOtpDeliveryMechanismRequest) GetEmailAddressJSONObjectFilter() string`

GetEmailAddressJSONObjectFilter returns the EmailAddressJSONObjectFilter field if non-nil, zero value otherwise.

### GetEmailAddressJSONObjectFilterOk

`func (o *AddOtpDeliveryMechanismRequest) GetEmailAddressJSONObjectFilterOk() (*string, bool)`

GetEmailAddressJSONObjectFilterOk returns a tuple with the EmailAddressJSONObjectFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddressJSONObjectFilter

`func (o *AddOtpDeliveryMechanismRequest) SetEmailAddressJSONObjectFilter(v string)`

SetEmailAddressJSONObjectFilter sets EmailAddressJSONObjectFilter field to given value.

### HasEmailAddressJSONObjectFilter

`func (o *AddOtpDeliveryMechanismRequest) HasEmailAddressJSONObjectFilter() bool`

HasEmailAddressJSONObjectFilter returns a boolean if a field has been set.

### GetSenderAddress

`func (o *AddOtpDeliveryMechanismRequest) GetSenderAddress() string`

GetSenderAddress returns the SenderAddress field if non-nil, zero value otherwise.

### GetSenderAddressOk

`func (o *AddOtpDeliveryMechanismRequest) GetSenderAddressOk() (*string, bool)`

GetSenderAddressOk returns a tuple with the SenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddress

`func (o *AddOtpDeliveryMechanismRequest) SetSenderAddress(v string)`

SetSenderAddress sets SenderAddress field to given value.


### GetMessageSubject

`func (o *AddOtpDeliveryMechanismRequest) GetMessageSubject() string`

GetMessageSubject returns the MessageSubject field if non-nil, zero value otherwise.

### GetMessageSubjectOk

`func (o *AddOtpDeliveryMechanismRequest) GetMessageSubjectOk() (*string, bool)`

GetMessageSubjectOk returns a tuple with the MessageSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSubject

`func (o *AddOtpDeliveryMechanismRequest) SetMessageSubject(v string)`

SetMessageSubject sets MessageSubject field to given value.

### HasMessageSubject

`func (o *AddOtpDeliveryMechanismRequest) HasMessageSubject() bool`

HasMessageSubject returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddOtpDeliveryMechanismRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddOtpDeliveryMechanismRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddOtpDeliveryMechanismRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddOtpDeliveryMechanismRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddOtpDeliveryMechanismRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddOtpDeliveryMechanismRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddOtpDeliveryMechanismRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


