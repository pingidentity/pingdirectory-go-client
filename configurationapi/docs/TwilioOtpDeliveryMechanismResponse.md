# TwilioOtpDeliveryMechanismResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumtwilioOtpDeliveryMechanismSchemaUrn**](EnumtwilioOtpDeliveryMechanismSchemaUrn.md) |  | 
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
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the OTP Delivery Mechanism | 

## Methods

### NewTwilioOtpDeliveryMechanismResponse

`func NewTwilioOtpDeliveryMechanismResponse(schemas []EnumtwilioOtpDeliveryMechanismSchemaUrn, twilioAccountSID string, phoneNumberAttributeType string, senderPhoneNumber []string, enabled bool, id string, ) *TwilioOtpDeliveryMechanismResponse`

NewTwilioOtpDeliveryMechanismResponse instantiates a new TwilioOtpDeliveryMechanismResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTwilioOtpDeliveryMechanismResponseWithDefaults

`func NewTwilioOtpDeliveryMechanismResponseWithDefaults() *TwilioOtpDeliveryMechanismResponse`

NewTwilioOtpDeliveryMechanismResponseWithDefaults instantiates a new TwilioOtpDeliveryMechanismResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *TwilioOtpDeliveryMechanismResponse) GetSchemas() []EnumtwilioOtpDeliveryMechanismSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *TwilioOtpDeliveryMechanismResponse) GetSchemasOk() (*[]EnumtwilioOtpDeliveryMechanismSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *TwilioOtpDeliveryMechanismResponse) SetSchemas(v []EnumtwilioOtpDeliveryMechanismSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetHttpProxyExternalServer

`func (o *TwilioOtpDeliveryMechanismResponse) GetHttpProxyExternalServer() string`

GetHttpProxyExternalServer returns the HttpProxyExternalServer field if non-nil, zero value otherwise.

### GetHttpProxyExternalServerOk

`func (o *TwilioOtpDeliveryMechanismResponse) GetHttpProxyExternalServerOk() (*string, bool)`

GetHttpProxyExternalServerOk returns a tuple with the HttpProxyExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyExternalServer

`func (o *TwilioOtpDeliveryMechanismResponse) SetHttpProxyExternalServer(v string)`

SetHttpProxyExternalServer sets HttpProxyExternalServer field to given value.

### HasHttpProxyExternalServer

`func (o *TwilioOtpDeliveryMechanismResponse) HasHttpProxyExternalServer() bool`

HasHttpProxyExternalServer returns a boolean if a field has been set.

### GetTwilioAccountSID

`func (o *TwilioOtpDeliveryMechanismResponse) GetTwilioAccountSID() string`

GetTwilioAccountSID returns the TwilioAccountSID field if non-nil, zero value otherwise.

### GetTwilioAccountSIDOk

`func (o *TwilioOtpDeliveryMechanismResponse) GetTwilioAccountSIDOk() (*string, bool)`

GetTwilioAccountSIDOk returns a tuple with the TwilioAccountSID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAccountSID

`func (o *TwilioOtpDeliveryMechanismResponse) SetTwilioAccountSID(v string)`

SetTwilioAccountSID sets TwilioAccountSID field to given value.


### GetTwilioAuthToken

`func (o *TwilioOtpDeliveryMechanismResponse) GetTwilioAuthToken() string`

GetTwilioAuthToken returns the TwilioAuthToken field if non-nil, zero value otherwise.

### GetTwilioAuthTokenOk

`func (o *TwilioOtpDeliveryMechanismResponse) GetTwilioAuthTokenOk() (*string, bool)`

GetTwilioAuthTokenOk returns a tuple with the TwilioAuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAuthToken

`func (o *TwilioOtpDeliveryMechanismResponse) SetTwilioAuthToken(v string)`

SetTwilioAuthToken sets TwilioAuthToken field to given value.

### HasTwilioAuthToken

`func (o *TwilioOtpDeliveryMechanismResponse) HasTwilioAuthToken() bool`

HasTwilioAuthToken returns a boolean if a field has been set.

### GetTwilioAuthTokenPassphraseProvider

`func (o *TwilioOtpDeliveryMechanismResponse) GetTwilioAuthTokenPassphraseProvider() string`

GetTwilioAuthTokenPassphraseProvider returns the TwilioAuthTokenPassphraseProvider field if non-nil, zero value otherwise.

### GetTwilioAuthTokenPassphraseProviderOk

`func (o *TwilioOtpDeliveryMechanismResponse) GetTwilioAuthTokenPassphraseProviderOk() (*string, bool)`

GetTwilioAuthTokenPassphraseProviderOk returns a tuple with the TwilioAuthTokenPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAuthTokenPassphraseProvider

`func (o *TwilioOtpDeliveryMechanismResponse) SetTwilioAuthTokenPassphraseProvider(v string)`

SetTwilioAuthTokenPassphraseProvider sets TwilioAuthTokenPassphraseProvider field to given value.

### HasTwilioAuthTokenPassphraseProvider

`func (o *TwilioOtpDeliveryMechanismResponse) HasTwilioAuthTokenPassphraseProvider() bool`

HasTwilioAuthTokenPassphraseProvider returns a boolean if a field has been set.

### GetPhoneNumberAttributeType

`func (o *TwilioOtpDeliveryMechanismResponse) GetPhoneNumberAttributeType() string`

GetPhoneNumberAttributeType returns the PhoneNumberAttributeType field if non-nil, zero value otherwise.

### GetPhoneNumberAttributeTypeOk

`func (o *TwilioOtpDeliveryMechanismResponse) GetPhoneNumberAttributeTypeOk() (*string, bool)`

GetPhoneNumberAttributeTypeOk returns a tuple with the PhoneNumberAttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberAttributeType

`func (o *TwilioOtpDeliveryMechanismResponse) SetPhoneNumberAttributeType(v string)`

SetPhoneNumberAttributeType sets PhoneNumberAttributeType field to given value.


### GetPhoneNumberJSONField

`func (o *TwilioOtpDeliveryMechanismResponse) GetPhoneNumberJSONField() string`

GetPhoneNumberJSONField returns the PhoneNumberJSONField field if non-nil, zero value otherwise.

### GetPhoneNumberJSONFieldOk

`func (o *TwilioOtpDeliveryMechanismResponse) GetPhoneNumberJSONFieldOk() (*string, bool)`

GetPhoneNumberJSONFieldOk returns a tuple with the PhoneNumberJSONField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberJSONField

`func (o *TwilioOtpDeliveryMechanismResponse) SetPhoneNumberJSONField(v string)`

SetPhoneNumberJSONField sets PhoneNumberJSONField field to given value.

### HasPhoneNumberJSONField

`func (o *TwilioOtpDeliveryMechanismResponse) HasPhoneNumberJSONField() bool`

HasPhoneNumberJSONField returns a boolean if a field has been set.

### GetPhoneNumberJSONObjectFilter

`func (o *TwilioOtpDeliveryMechanismResponse) GetPhoneNumberJSONObjectFilter() string`

GetPhoneNumberJSONObjectFilter returns the PhoneNumberJSONObjectFilter field if non-nil, zero value otherwise.

### GetPhoneNumberJSONObjectFilterOk

`func (o *TwilioOtpDeliveryMechanismResponse) GetPhoneNumberJSONObjectFilterOk() (*string, bool)`

GetPhoneNumberJSONObjectFilterOk returns a tuple with the PhoneNumberJSONObjectFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberJSONObjectFilter

`func (o *TwilioOtpDeliveryMechanismResponse) SetPhoneNumberJSONObjectFilter(v string)`

SetPhoneNumberJSONObjectFilter sets PhoneNumberJSONObjectFilter field to given value.

### HasPhoneNumberJSONObjectFilter

`func (o *TwilioOtpDeliveryMechanismResponse) HasPhoneNumberJSONObjectFilter() bool`

HasPhoneNumberJSONObjectFilter returns a boolean if a field has been set.

### GetSenderPhoneNumber

`func (o *TwilioOtpDeliveryMechanismResponse) GetSenderPhoneNumber() []string`

GetSenderPhoneNumber returns the SenderPhoneNumber field if non-nil, zero value otherwise.

### GetSenderPhoneNumberOk

`func (o *TwilioOtpDeliveryMechanismResponse) GetSenderPhoneNumberOk() (*[]string, bool)`

GetSenderPhoneNumberOk returns a tuple with the SenderPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderPhoneNumber

`func (o *TwilioOtpDeliveryMechanismResponse) SetSenderPhoneNumber(v []string)`

SetSenderPhoneNumber sets SenderPhoneNumber field to given value.


### GetMessageTextBeforeOTP

`func (o *TwilioOtpDeliveryMechanismResponse) GetMessageTextBeforeOTP() string`

GetMessageTextBeforeOTP returns the MessageTextBeforeOTP field if non-nil, zero value otherwise.

### GetMessageTextBeforeOTPOk

`func (o *TwilioOtpDeliveryMechanismResponse) GetMessageTextBeforeOTPOk() (*string, bool)`

GetMessageTextBeforeOTPOk returns a tuple with the MessageTextBeforeOTP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTextBeforeOTP

`func (o *TwilioOtpDeliveryMechanismResponse) SetMessageTextBeforeOTP(v string)`

SetMessageTextBeforeOTP sets MessageTextBeforeOTP field to given value.

### HasMessageTextBeforeOTP

`func (o *TwilioOtpDeliveryMechanismResponse) HasMessageTextBeforeOTP() bool`

HasMessageTextBeforeOTP returns a boolean if a field has been set.

### GetMessageTextAfterOTP

`func (o *TwilioOtpDeliveryMechanismResponse) GetMessageTextAfterOTP() string`

GetMessageTextAfterOTP returns the MessageTextAfterOTP field if non-nil, zero value otherwise.

### GetMessageTextAfterOTPOk

`func (o *TwilioOtpDeliveryMechanismResponse) GetMessageTextAfterOTPOk() (*string, bool)`

GetMessageTextAfterOTPOk returns a tuple with the MessageTextAfterOTP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTextAfterOTP

`func (o *TwilioOtpDeliveryMechanismResponse) SetMessageTextAfterOTP(v string)`

SetMessageTextAfterOTP sets MessageTextAfterOTP field to given value.

### HasMessageTextAfterOTP

`func (o *TwilioOtpDeliveryMechanismResponse) HasMessageTextAfterOTP() bool`

HasMessageTextAfterOTP returns a boolean if a field has been set.

### GetDescription

`func (o *TwilioOtpDeliveryMechanismResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TwilioOtpDeliveryMechanismResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TwilioOtpDeliveryMechanismResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TwilioOtpDeliveryMechanismResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *TwilioOtpDeliveryMechanismResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TwilioOtpDeliveryMechanismResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TwilioOtpDeliveryMechanismResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *TwilioOtpDeliveryMechanismResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TwilioOtpDeliveryMechanismResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TwilioOtpDeliveryMechanismResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TwilioOtpDeliveryMechanismResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *TwilioOtpDeliveryMechanismResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *TwilioOtpDeliveryMechanismResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *TwilioOtpDeliveryMechanismResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *TwilioOtpDeliveryMechanismResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *TwilioOtpDeliveryMechanismResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TwilioOtpDeliveryMechanismResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TwilioOtpDeliveryMechanismResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


