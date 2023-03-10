# AddOtpDeliveryMechanism200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the OTP Delivery Mechanism | 
**Schemas** | [**[]EnumthirdPartyOtpDeliveryMechanismSchemaUrn**](EnumthirdPartyOtpDeliveryMechanismSchemaUrn.md) |  | 
**HttpProxyExternalServer** | Pointer to **string** | A reference to an HTTP proxy server that should be used for requests sent to the Twilio service. | [optional] 
**TwilioAccountSID** | **string** | The unique identifier assigned to the Twilio account that will be used. | 
**TwilioAuthToken** | Pointer to **string** | The auth token for the Twilio account that will be used. | [optional] 
**TwilioAuthTokenPassphraseProvider** | Pointer to **string** | The passphrase provider that may be used to obtain the auth token for the Twilio account that will be used. | [optional] 
**PhoneNumberAttributeType** | **string** | The name or OID of the attribute in the user&#39;s entry that holds the phone number to which the message should be sent. | 
**PhoneNumberJSONField** | Pointer to **string** | The name of the JSON field whose value is the phone number to which the message should be sent. The phone number must be contained in a top-level field whose value is a single string. | [optional] 
**PhoneNumberJSONObjectFilter** | Pointer to **string** | A JSON object filter that may be used to identify which phone number value to use when sending the message. | [optional] 
**SenderPhoneNumber** | **[]string** | The outgoing phone number to use for the messages. Values must be phone numbers you have obtained for use with your Twilio account. | 
**MessageTextBeforeOTP** | Pointer to **string** | Any text that should appear in the message before the one-time password value. | [optional] 
**MessageTextAfterOTP** | Pointer to **string** | Any text that should appear in the message after the one-time password value. | [optional] 
**Description** | Pointer to **string** | A description for this OTP Delivery Mechanism | [optional] 
**Enabled** | **bool** | Indicates whether this OTP Delivery Mechanism is enabled for use in the server. | 
**EmailAddressAttributeType** | **string** | The name or OID of the attribute that holds the email address to which the message should be sent. | 
**EmailAddressJSONField** | Pointer to **string** | The name of the JSON field whose value is the email address to which the message should be sent. The email address must be contained in a top-level field whose value is a single string. | [optional] 
**EmailAddressJSONObjectFilter** | Pointer to **string** | A JSON object filter that may be used to identify which email address value to use when sending the message. | [optional] 
**SenderAddress** | **string** | The e-mail address to use as the sender for the one-time password. | 
**MessageSubject** | **string** | The subject to use for the e-mail message. | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party OTP Delivery Mechanism. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party OTP Delivery Mechanism. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddOtpDeliveryMechanism200Response

`func NewAddOtpDeliveryMechanism200Response(id string, schemas []EnumthirdPartyOtpDeliveryMechanismSchemaUrn, twilioAccountSID string, phoneNumberAttributeType string, senderPhoneNumber []string, enabled bool, emailAddressAttributeType string, senderAddress string, messageSubject string, extensionClass string, ) *AddOtpDeliveryMechanism200Response`

NewAddOtpDeliveryMechanism200Response instantiates a new AddOtpDeliveryMechanism200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddOtpDeliveryMechanism200ResponseWithDefaults

`func NewAddOtpDeliveryMechanism200ResponseWithDefaults() *AddOtpDeliveryMechanism200Response`

NewAddOtpDeliveryMechanism200ResponseWithDefaults instantiates a new AddOtpDeliveryMechanism200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *AddOtpDeliveryMechanism200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddOtpDeliveryMechanism200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddOtpDeliveryMechanism200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddOtpDeliveryMechanism200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AddOtpDeliveryMechanism200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AddOtpDeliveryMechanism200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AddOtpDeliveryMechanism200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AddOtpDeliveryMechanism200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *AddOtpDeliveryMechanism200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddOtpDeliveryMechanism200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddOtpDeliveryMechanism200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddOtpDeliveryMechanism200Response) GetSchemas() []EnumthirdPartyOtpDeliveryMechanismSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddOtpDeliveryMechanism200Response) GetSchemasOk() (*[]EnumthirdPartyOtpDeliveryMechanismSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddOtpDeliveryMechanism200Response) SetSchemas(v []EnumthirdPartyOtpDeliveryMechanismSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetHttpProxyExternalServer

`func (o *AddOtpDeliveryMechanism200Response) GetHttpProxyExternalServer() string`

GetHttpProxyExternalServer returns the HttpProxyExternalServer field if non-nil, zero value otherwise.

### GetHttpProxyExternalServerOk

`func (o *AddOtpDeliveryMechanism200Response) GetHttpProxyExternalServerOk() (*string, bool)`

GetHttpProxyExternalServerOk returns a tuple with the HttpProxyExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyExternalServer

`func (o *AddOtpDeliveryMechanism200Response) SetHttpProxyExternalServer(v string)`

SetHttpProxyExternalServer sets HttpProxyExternalServer field to given value.

### HasHttpProxyExternalServer

`func (o *AddOtpDeliveryMechanism200Response) HasHttpProxyExternalServer() bool`

HasHttpProxyExternalServer returns a boolean if a field has been set.

### GetTwilioAccountSID

`func (o *AddOtpDeliveryMechanism200Response) GetTwilioAccountSID() string`

GetTwilioAccountSID returns the TwilioAccountSID field if non-nil, zero value otherwise.

### GetTwilioAccountSIDOk

`func (o *AddOtpDeliveryMechanism200Response) GetTwilioAccountSIDOk() (*string, bool)`

GetTwilioAccountSIDOk returns a tuple with the TwilioAccountSID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAccountSID

`func (o *AddOtpDeliveryMechanism200Response) SetTwilioAccountSID(v string)`

SetTwilioAccountSID sets TwilioAccountSID field to given value.


### GetTwilioAuthToken

`func (o *AddOtpDeliveryMechanism200Response) GetTwilioAuthToken() string`

GetTwilioAuthToken returns the TwilioAuthToken field if non-nil, zero value otherwise.

### GetTwilioAuthTokenOk

`func (o *AddOtpDeliveryMechanism200Response) GetTwilioAuthTokenOk() (*string, bool)`

GetTwilioAuthTokenOk returns a tuple with the TwilioAuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAuthToken

`func (o *AddOtpDeliveryMechanism200Response) SetTwilioAuthToken(v string)`

SetTwilioAuthToken sets TwilioAuthToken field to given value.

### HasTwilioAuthToken

`func (o *AddOtpDeliveryMechanism200Response) HasTwilioAuthToken() bool`

HasTwilioAuthToken returns a boolean if a field has been set.

### GetTwilioAuthTokenPassphraseProvider

`func (o *AddOtpDeliveryMechanism200Response) GetTwilioAuthTokenPassphraseProvider() string`

GetTwilioAuthTokenPassphraseProvider returns the TwilioAuthTokenPassphraseProvider field if non-nil, zero value otherwise.

### GetTwilioAuthTokenPassphraseProviderOk

`func (o *AddOtpDeliveryMechanism200Response) GetTwilioAuthTokenPassphraseProviderOk() (*string, bool)`

GetTwilioAuthTokenPassphraseProviderOk returns a tuple with the TwilioAuthTokenPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAuthTokenPassphraseProvider

`func (o *AddOtpDeliveryMechanism200Response) SetTwilioAuthTokenPassphraseProvider(v string)`

SetTwilioAuthTokenPassphraseProvider sets TwilioAuthTokenPassphraseProvider field to given value.

### HasTwilioAuthTokenPassphraseProvider

`func (o *AddOtpDeliveryMechanism200Response) HasTwilioAuthTokenPassphraseProvider() bool`

HasTwilioAuthTokenPassphraseProvider returns a boolean if a field has been set.

### GetPhoneNumberAttributeType

`func (o *AddOtpDeliveryMechanism200Response) GetPhoneNumberAttributeType() string`

GetPhoneNumberAttributeType returns the PhoneNumberAttributeType field if non-nil, zero value otherwise.

### GetPhoneNumberAttributeTypeOk

`func (o *AddOtpDeliveryMechanism200Response) GetPhoneNumberAttributeTypeOk() (*string, bool)`

GetPhoneNumberAttributeTypeOk returns a tuple with the PhoneNumberAttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberAttributeType

`func (o *AddOtpDeliveryMechanism200Response) SetPhoneNumberAttributeType(v string)`

SetPhoneNumberAttributeType sets PhoneNumberAttributeType field to given value.


### GetPhoneNumberJSONField

`func (o *AddOtpDeliveryMechanism200Response) GetPhoneNumberJSONField() string`

GetPhoneNumberJSONField returns the PhoneNumberJSONField field if non-nil, zero value otherwise.

### GetPhoneNumberJSONFieldOk

`func (o *AddOtpDeliveryMechanism200Response) GetPhoneNumberJSONFieldOk() (*string, bool)`

GetPhoneNumberJSONFieldOk returns a tuple with the PhoneNumberJSONField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberJSONField

`func (o *AddOtpDeliveryMechanism200Response) SetPhoneNumberJSONField(v string)`

SetPhoneNumberJSONField sets PhoneNumberJSONField field to given value.

### HasPhoneNumberJSONField

`func (o *AddOtpDeliveryMechanism200Response) HasPhoneNumberJSONField() bool`

HasPhoneNumberJSONField returns a boolean if a field has been set.

### GetPhoneNumberJSONObjectFilter

`func (o *AddOtpDeliveryMechanism200Response) GetPhoneNumberJSONObjectFilter() string`

GetPhoneNumberJSONObjectFilter returns the PhoneNumberJSONObjectFilter field if non-nil, zero value otherwise.

### GetPhoneNumberJSONObjectFilterOk

`func (o *AddOtpDeliveryMechanism200Response) GetPhoneNumberJSONObjectFilterOk() (*string, bool)`

GetPhoneNumberJSONObjectFilterOk returns a tuple with the PhoneNumberJSONObjectFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberJSONObjectFilter

`func (o *AddOtpDeliveryMechanism200Response) SetPhoneNumberJSONObjectFilter(v string)`

SetPhoneNumberJSONObjectFilter sets PhoneNumberJSONObjectFilter field to given value.

### HasPhoneNumberJSONObjectFilter

`func (o *AddOtpDeliveryMechanism200Response) HasPhoneNumberJSONObjectFilter() bool`

HasPhoneNumberJSONObjectFilter returns a boolean if a field has been set.

### GetSenderPhoneNumber

`func (o *AddOtpDeliveryMechanism200Response) GetSenderPhoneNumber() []string`

GetSenderPhoneNumber returns the SenderPhoneNumber field if non-nil, zero value otherwise.

### GetSenderPhoneNumberOk

`func (o *AddOtpDeliveryMechanism200Response) GetSenderPhoneNumberOk() (*[]string, bool)`

GetSenderPhoneNumberOk returns a tuple with the SenderPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderPhoneNumber

`func (o *AddOtpDeliveryMechanism200Response) SetSenderPhoneNumber(v []string)`

SetSenderPhoneNumber sets SenderPhoneNumber field to given value.


### GetMessageTextBeforeOTP

`func (o *AddOtpDeliveryMechanism200Response) GetMessageTextBeforeOTP() string`

GetMessageTextBeforeOTP returns the MessageTextBeforeOTP field if non-nil, zero value otherwise.

### GetMessageTextBeforeOTPOk

`func (o *AddOtpDeliveryMechanism200Response) GetMessageTextBeforeOTPOk() (*string, bool)`

GetMessageTextBeforeOTPOk returns a tuple with the MessageTextBeforeOTP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTextBeforeOTP

`func (o *AddOtpDeliveryMechanism200Response) SetMessageTextBeforeOTP(v string)`

SetMessageTextBeforeOTP sets MessageTextBeforeOTP field to given value.

### HasMessageTextBeforeOTP

`func (o *AddOtpDeliveryMechanism200Response) HasMessageTextBeforeOTP() bool`

HasMessageTextBeforeOTP returns a boolean if a field has been set.

### GetMessageTextAfterOTP

`func (o *AddOtpDeliveryMechanism200Response) GetMessageTextAfterOTP() string`

GetMessageTextAfterOTP returns the MessageTextAfterOTP field if non-nil, zero value otherwise.

### GetMessageTextAfterOTPOk

`func (o *AddOtpDeliveryMechanism200Response) GetMessageTextAfterOTPOk() (*string, bool)`

GetMessageTextAfterOTPOk returns a tuple with the MessageTextAfterOTP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTextAfterOTP

`func (o *AddOtpDeliveryMechanism200Response) SetMessageTextAfterOTP(v string)`

SetMessageTextAfterOTP sets MessageTextAfterOTP field to given value.

### HasMessageTextAfterOTP

`func (o *AddOtpDeliveryMechanism200Response) HasMessageTextAfterOTP() bool`

HasMessageTextAfterOTP returns a boolean if a field has been set.

### GetDescription

`func (o *AddOtpDeliveryMechanism200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddOtpDeliveryMechanism200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddOtpDeliveryMechanism200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddOtpDeliveryMechanism200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddOtpDeliveryMechanism200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddOtpDeliveryMechanism200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddOtpDeliveryMechanism200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEmailAddressAttributeType

`func (o *AddOtpDeliveryMechanism200Response) GetEmailAddressAttributeType() string`

GetEmailAddressAttributeType returns the EmailAddressAttributeType field if non-nil, zero value otherwise.

### GetEmailAddressAttributeTypeOk

`func (o *AddOtpDeliveryMechanism200Response) GetEmailAddressAttributeTypeOk() (*string, bool)`

GetEmailAddressAttributeTypeOk returns a tuple with the EmailAddressAttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddressAttributeType

`func (o *AddOtpDeliveryMechanism200Response) SetEmailAddressAttributeType(v string)`

SetEmailAddressAttributeType sets EmailAddressAttributeType field to given value.


### GetEmailAddressJSONField

`func (o *AddOtpDeliveryMechanism200Response) GetEmailAddressJSONField() string`

GetEmailAddressJSONField returns the EmailAddressJSONField field if non-nil, zero value otherwise.

### GetEmailAddressJSONFieldOk

`func (o *AddOtpDeliveryMechanism200Response) GetEmailAddressJSONFieldOk() (*string, bool)`

GetEmailAddressJSONFieldOk returns a tuple with the EmailAddressJSONField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddressJSONField

`func (o *AddOtpDeliveryMechanism200Response) SetEmailAddressJSONField(v string)`

SetEmailAddressJSONField sets EmailAddressJSONField field to given value.

### HasEmailAddressJSONField

`func (o *AddOtpDeliveryMechanism200Response) HasEmailAddressJSONField() bool`

HasEmailAddressJSONField returns a boolean if a field has been set.

### GetEmailAddressJSONObjectFilter

`func (o *AddOtpDeliveryMechanism200Response) GetEmailAddressJSONObjectFilter() string`

GetEmailAddressJSONObjectFilter returns the EmailAddressJSONObjectFilter field if non-nil, zero value otherwise.

### GetEmailAddressJSONObjectFilterOk

`func (o *AddOtpDeliveryMechanism200Response) GetEmailAddressJSONObjectFilterOk() (*string, bool)`

GetEmailAddressJSONObjectFilterOk returns a tuple with the EmailAddressJSONObjectFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddressJSONObjectFilter

`func (o *AddOtpDeliveryMechanism200Response) SetEmailAddressJSONObjectFilter(v string)`

SetEmailAddressJSONObjectFilter sets EmailAddressJSONObjectFilter field to given value.

### HasEmailAddressJSONObjectFilter

`func (o *AddOtpDeliveryMechanism200Response) HasEmailAddressJSONObjectFilter() bool`

HasEmailAddressJSONObjectFilter returns a boolean if a field has been set.

### GetSenderAddress

`func (o *AddOtpDeliveryMechanism200Response) GetSenderAddress() string`

GetSenderAddress returns the SenderAddress field if non-nil, zero value otherwise.

### GetSenderAddressOk

`func (o *AddOtpDeliveryMechanism200Response) GetSenderAddressOk() (*string, bool)`

GetSenderAddressOk returns a tuple with the SenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddress

`func (o *AddOtpDeliveryMechanism200Response) SetSenderAddress(v string)`

SetSenderAddress sets SenderAddress field to given value.


### GetMessageSubject

`func (o *AddOtpDeliveryMechanism200Response) GetMessageSubject() string`

GetMessageSubject returns the MessageSubject field if non-nil, zero value otherwise.

### GetMessageSubjectOk

`func (o *AddOtpDeliveryMechanism200Response) GetMessageSubjectOk() (*string, bool)`

GetMessageSubjectOk returns a tuple with the MessageSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSubject

`func (o *AddOtpDeliveryMechanism200Response) SetMessageSubject(v string)`

SetMessageSubject sets MessageSubject field to given value.


### GetExtensionClass

`func (o *AddOtpDeliveryMechanism200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddOtpDeliveryMechanism200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddOtpDeliveryMechanism200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddOtpDeliveryMechanism200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddOtpDeliveryMechanism200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddOtpDeliveryMechanism200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddOtpDeliveryMechanism200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


