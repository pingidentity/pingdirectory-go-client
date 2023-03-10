# AddTwilioOtpDeliveryMechanismRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MechanismName** | **string** | Name of the new OTP Delivery Mechanism | 
**Schemas** | [**[]EnumtwilioOtpDeliveryMechanismSchemaUrn**](EnumtwilioOtpDeliveryMechanismSchemaUrn.md) |  | 
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

## Methods

### NewAddTwilioOtpDeliveryMechanismRequest

`func NewAddTwilioOtpDeliveryMechanismRequest(mechanismName string, schemas []EnumtwilioOtpDeliveryMechanismSchemaUrn, twilioAccountSID string, senderPhoneNumber []string, enabled bool, ) *AddTwilioOtpDeliveryMechanismRequest`

NewAddTwilioOtpDeliveryMechanismRequest instantiates a new AddTwilioOtpDeliveryMechanismRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTwilioOtpDeliveryMechanismRequestWithDefaults

`func NewAddTwilioOtpDeliveryMechanismRequestWithDefaults() *AddTwilioOtpDeliveryMechanismRequest`

NewAddTwilioOtpDeliveryMechanismRequestWithDefaults instantiates a new AddTwilioOtpDeliveryMechanismRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMechanismName

`func (o *AddTwilioOtpDeliveryMechanismRequest) GetMechanismName() string`

GetMechanismName returns the MechanismName field if non-nil, zero value otherwise.

### GetMechanismNameOk

`func (o *AddTwilioOtpDeliveryMechanismRequest) GetMechanismNameOk() (*string, bool)`

GetMechanismNameOk returns a tuple with the MechanismName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMechanismName

`func (o *AddTwilioOtpDeliveryMechanismRequest) SetMechanismName(v string)`

SetMechanismName sets MechanismName field to given value.


### GetSchemas

`func (o *AddTwilioOtpDeliveryMechanismRequest) GetSchemas() []EnumtwilioOtpDeliveryMechanismSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddTwilioOtpDeliveryMechanismRequest) GetSchemasOk() (*[]EnumtwilioOtpDeliveryMechanismSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddTwilioOtpDeliveryMechanismRequest) SetSchemas(v []EnumtwilioOtpDeliveryMechanismSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetHttpProxyExternalServer

`func (o *AddTwilioOtpDeliveryMechanismRequest) GetHttpProxyExternalServer() string`

GetHttpProxyExternalServer returns the HttpProxyExternalServer field if non-nil, zero value otherwise.

### GetHttpProxyExternalServerOk

`func (o *AddTwilioOtpDeliveryMechanismRequest) GetHttpProxyExternalServerOk() (*string, bool)`

GetHttpProxyExternalServerOk returns a tuple with the HttpProxyExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyExternalServer

`func (o *AddTwilioOtpDeliveryMechanismRequest) SetHttpProxyExternalServer(v string)`

SetHttpProxyExternalServer sets HttpProxyExternalServer field to given value.

### HasHttpProxyExternalServer

`func (o *AddTwilioOtpDeliveryMechanismRequest) HasHttpProxyExternalServer() bool`

HasHttpProxyExternalServer returns a boolean if a field has been set.

### GetTwilioAccountSID

`func (o *AddTwilioOtpDeliveryMechanismRequest) GetTwilioAccountSID() string`

GetTwilioAccountSID returns the TwilioAccountSID field if non-nil, zero value otherwise.

### GetTwilioAccountSIDOk

`func (o *AddTwilioOtpDeliveryMechanismRequest) GetTwilioAccountSIDOk() (*string, bool)`

GetTwilioAccountSIDOk returns a tuple with the TwilioAccountSID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAccountSID

`func (o *AddTwilioOtpDeliveryMechanismRequest) SetTwilioAccountSID(v string)`

SetTwilioAccountSID sets TwilioAccountSID field to given value.


### GetTwilioAuthToken

`func (o *AddTwilioOtpDeliveryMechanismRequest) GetTwilioAuthToken() string`

GetTwilioAuthToken returns the TwilioAuthToken field if non-nil, zero value otherwise.

### GetTwilioAuthTokenOk

`func (o *AddTwilioOtpDeliveryMechanismRequest) GetTwilioAuthTokenOk() (*string, bool)`

GetTwilioAuthTokenOk returns a tuple with the TwilioAuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAuthToken

`func (o *AddTwilioOtpDeliveryMechanismRequest) SetTwilioAuthToken(v string)`

SetTwilioAuthToken sets TwilioAuthToken field to given value.

### HasTwilioAuthToken

`func (o *AddTwilioOtpDeliveryMechanismRequest) HasTwilioAuthToken() bool`

HasTwilioAuthToken returns a boolean if a field has been set.

### GetTwilioAuthTokenPassphraseProvider

`func (o *AddTwilioOtpDeliveryMechanismRequest) GetTwilioAuthTokenPassphraseProvider() string`

GetTwilioAuthTokenPassphraseProvider returns the TwilioAuthTokenPassphraseProvider field if non-nil, zero value otherwise.

### GetTwilioAuthTokenPassphraseProviderOk

`func (o *AddTwilioOtpDeliveryMechanismRequest) GetTwilioAuthTokenPassphraseProviderOk() (*string, bool)`

GetTwilioAuthTokenPassphraseProviderOk returns a tuple with the TwilioAuthTokenPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAuthTokenPassphraseProvider

`func (o *AddTwilioOtpDeliveryMechanismRequest) SetTwilioAuthTokenPassphraseProvider(v string)`

SetTwilioAuthTokenPassphraseProvider sets TwilioAuthTokenPassphraseProvider field to given value.

### HasTwilioAuthTokenPassphraseProvider

`func (o *AddTwilioOtpDeliveryMechanismRequest) HasTwilioAuthTokenPassphraseProvider() bool`

HasTwilioAuthTokenPassphraseProvider returns a boolean if a field has been set.

### GetPhoneNumberAttributeType

`func (o *AddTwilioOtpDeliveryMechanismRequest) GetPhoneNumberAttributeType() string`

GetPhoneNumberAttributeType returns the PhoneNumberAttributeType field if non-nil, zero value otherwise.

### GetPhoneNumberAttributeTypeOk

`func (o *AddTwilioOtpDeliveryMechanismRequest) GetPhoneNumberAttributeTypeOk() (*string, bool)`

GetPhoneNumberAttributeTypeOk returns a tuple with the PhoneNumberAttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberAttributeType

`func (o *AddTwilioOtpDeliveryMechanismRequest) SetPhoneNumberAttributeType(v string)`

SetPhoneNumberAttributeType sets PhoneNumberAttributeType field to given value.

### HasPhoneNumberAttributeType

`func (o *AddTwilioOtpDeliveryMechanismRequest) HasPhoneNumberAttributeType() bool`

HasPhoneNumberAttributeType returns a boolean if a field has been set.

### GetPhoneNumberJSONField

`func (o *AddTwilioOtpDeliveryMechanismRequest) GetPhoneNumberJSONField() string`

GetPhoneNumberJSONField returns the PhoneNumberJSONField field if non-nil, zero value otherwise.

### GetPhoneNumberJSONFieldOk

`func (o *AddTwilioOtpDeliveryMechanismRequest) GetPhoneNumberJSONFieldOk() (*string, bool)`

GetPhoneNumberJSONFieldOk returns a tuple with the PhoneNumberJSONField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberJSONField

`func (o *AddTwilioOtpDeliveryMechanismRequest) SetPhoneNumberJSONField(v string)`

SetPhoneNumberJSONField sets PhoneNumberJSONField field to given value.

### HasPhoneNumberJSONField

`func (o *AddTwilioOtpDeliveryMechanismRequest) HasPhoneNumberJSONField() bool`

HasPhoneNumberJSONField returns a boolean if a field has been set.

### GetPhoneNumberJSONObjectFilter

`func (o *AddTwilioOtpDeliveryMechanismRequest) GetPhoneNumberJSONObjectFilter() string`

GetPhoneNumberJSONObjectFilter returns the PhoneNumberJSONObjectFilter field if non-nil, zero value otherwise.

### GetPhoneNumberJSONObjectFilterOk

`func (o *AddTwilioOtpDeliveryMechanismRequest) GetPhoneNumberJSONObjectFilterOk() (*string, bool)`

GetPhoneNumberJSONObjectFilterOk returns a tuple with the PhoneNumberJSONObjectFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberJSONObjectFilter

`func (o *AddTwilioOtpDeliveryMechanismRequest) SetPhoneNumberJSONObjectFilter(v string)`

SetPhoneNumberJSONObjectFilter sets PhoneNumberJSONObjectFilter field to given value.

### HasPhoneNumberJSONObjectFilter

`func (o *AddTwilioOtpDeliveryMechanismRequest) HasPhoneNumberJSONObjectFilter() bool`

HasPhoneNumberJSONObjectFilter returns a boolean if a field has been set.

### GetSenderPhoneNumber

`func (o *AddTwilioOtpDeliveryMechanismRequest) GetSenderPhoneNumber() []string`

GetSenderPhoneNumber returns the SenderPhoneNumber field if non-nil, zero value otherwise.

### GetSenderPhoneNumberOk

`func (o *AddTwilioOtpDeliveryMechanismRequest) GetSenderPhoneNumberOk() (*[]string, bool)`

GetSenderPhoneNumberOk returns a tuple with the SenderPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderPhoneNumber

`func (o *AddTwilioOtpDeliveryMechanismRequest) SetSenderPhoneNumber(v []string)`

SetSenderPhoneNumber sets SenderPhoneNumber field to given value.


### GetMessageTextBeforeOTP

`func (o *AddTwilioOtpDeliveryMechanismRequest) GetMessageTextBeforeOTP() string`

GetMessageTextBeforeOTP returns the MessageTextBeforeOTP field if non-nil, zero value otherwise.

### GetMessageTextBeforeOTPOk

`func (o *AddTwilioOtpDeliveryMechanismRequest) GetMessageTextBeforeOTPOk() (*string, bool)`

GetMessageTextBeforeOTPOk returns a tuple with the MessageTextBeforeOTP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTextBeforeOTP

`func (o *AddTwilioOtpDeliveryMechanismRequest) SetMessageTextBeforeOTP(v string)`

SetMessageTextBeforeOTP sets MessageTextBeforeOTP field to given value.

### HasMessageTextBeforeOTP

`func (o *AddTwilioOtpDeliveryMechanismRequest) HasMessageTextBeforeOTP() bool`

HasMessageTextBeforeOTP returns a boolean if a field has been set.

### GetMessageTextAfterOTP

`func (o *AddTwilioOtpDeliveryMechanismRequest) GetMessageTextAfterOTP() string`

GetMessageTextAfterOTP returns the MessageTextAfterOTP field if non-nil, zero value otherwise.

### GetMessageTextAfterOTPOk

`func (o *AddTwilioOtpDeliveryMechanismRequest) GetMessageTextAfterOTPOk() (*string, bool)`

GetMessageTextAfterOTPOk returns a tuple with the MessageTextAfterOTP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTextAfterOTP

`func (o *AddTwilioOtpDeliveryMechanismRequest) SetMessageTextAfterOTP(v string)`

SetMessageTextAfterOTP sets MessageTextAfterOTP field to given value.

### HasMessageTextAfterOTP

`func (o *AddTwilioOtpDeliveryMechanismRequest) HasMessageTextAfterOTP() bool`

HasMessageTextAfterOTP returns a boolean if a field has been set.

### GetDescription

`func (o *AddTwilioOtpDeliveryMechanismRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddTwilioOtpDeliveryMechanismRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddTwilioOtpDeliveryMechanismRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddTwilioOtpDeliveryMechanismRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddTwilioOtpDeliveryMechanismRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddTwilioOtpDeliveryMechanismRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddTwilioOtpDeliveryMechanismRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


