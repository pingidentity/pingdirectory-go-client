# AddTwilioAlertHandlerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumtwilioAlertHandlerSchemaUrn**](EnumtwilioAlertHandlerSchemaUrn.md) |  | 
**Asynchronous** | Pointer to **bool** | Indicates whether the server should attempt to invoke this Twilio Alert Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver the alert notification) will not delay whatever processing the server was performing when the alert was generated. | [optional] 
**HttpProxyExternalServer** | Pointer to **string** | A reference to an HTTP proxy server that should be used for requests sent to the Twilio service. | [optional] 
**TwilioAccountSID** | **string** | The unique identifier assigned to the Twilio account that will be used. | 
**TwilioAuthToken** | Pointer to **string** | The auth token for the Twilio account that will be used. | [optional] 
**TwilioAuthTokenPassphraseProvider** | Pointer to **string** | The passphrase provider that may be used to obtain the auth token for the Twilio account that will be used. | [optional] 
**SenderPhoneNumber** | **[]string** | The outgoing phone number to use for the messages. Values must be phone numbers you have obtained for use with your Twilio account. | 
**RecipientPhoneNumber** | **[]string** | The phone number to which alert notifications should be delivered. | 
**LongMessageBehavior** | Pointer to [**EnumalertHandlerLongMessageBehaviorProp**](EnumalertHandlerLongMessageBehaviorProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Alert Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Alert Handler is enabled. | 
**EnabledAlertSeverity** | Pointer to [**[]EnumalertHandlerEnabledAlertSeverityProp**](EnumalertHandlerEnabledAlertSeverityProp.md) |  | [optional] 
**EnabledAlertType** | Pointer to [**[]EnumalertHandlerEnabledAlertTypeProp**](EnumalertHandlerEnabledAlertTypeProp.md) |  | [optional] 
**DisabledAlertType** | Pointer to [**[]EnumalertHandlerDisabledAlertTypeProp**](EnumalertHandlerDisabledAlertTypeProp.md) |  | [optional] 
**HandlerName** | **string** | Name of the new Alert Handler | 

## Methods

### NewAddTwilioAlertHandlerRequest

`func NewAddTwilioAlertHandlerRequest(schemas []EnumtwilioAlertHandlerSchemaUrn, twilioAccountSID string, senderPhoneNumber []string, recipientPhoneNumber []string, enabled bool, handlerName string, ) *AddTwilioAlertHandlerRequest`

NewAddTwilioAlertHandlerRequest instantiates a new AddTwilioAlertHandlerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTwilioAlertHandlerRequestWithDefaults

`func NewAddTwilioAlertHandlerRequestWithDefaults() *AddTwilioAlertHandlerRequest`

NewAddTwilioAlertHandlerRequestWithDefaults instantiates a new AddTwilioAlertHandlerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddTwilioAlertHandlerRequest) GetSchemas() []EnumtwilioAlertHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddTwilioAlertHandlerRequest) GetSchemasOk() (*[]EnumtwilioAlertHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddTwilioAlertHandlerRequest) SetSchemas(v []EnumtwilioAlertHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAsynchronous

`func (o *AddTwilioAlertHandlerRequest) GetAsynchronous() bool`

GetAsynchronous returns the Asynchronous field if non-nil, zero value otherwise.

### GetAsynchronousOk

`func (o *AddTwilioAlertHandlerRequest) GetAsynchronousOk() (*bool, bool)`

GetAsynchronousOk returns a tuple with the Asynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsynchronous

`func (o *AddTwilioAlertHandlerRequest) SetAsynchronous(v bool)`

SetAsynchronous sets Asynchronous field to given value.

### HasAsynchronous

`func (o *AddTwilioAlertHandlerRequest) HasAsynchronous() bool`

HasAsynchronous returns a boolean if a field has been set.

### GetHttpProxyExternalServer

`func (o *AddTwilioAlertHandlerRequest) GetHttpProxyExternalServer() string`

GetHttpProxyExternalServer returns the HttpProxyExternalServer field if non-nil, zero value otherwise.

### GetHttpProxyExternalServerOk

`func (o *AddTwilioAlertHandlerRequest) GetHttpProxyExternalServerOk() (*string, bool)`

GetHttpProxyExternalServerOk returns a tuple with the HttpProxyExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyExternalServer

`func (o *AddTwilioAlertHandlerRequest) SetHttpProxyExternalServer(v string)`

SetHttpProxyExternalServer sets HttpProxyExternalServer field to given value.

### HasHttpProxyExternalServer

`func (o *AddTwilioAlertHandlerRequest) HasHttpProxyExternalServer() bool`

HasHttpProxyExternalServer returns a boolean if a field has been set.

### GetTwilioAccountSID

`func (o *AddTwilioAlertHandlerRequest) GetTwilioAccountSID() string`

GetTwilioAccountSID returns the TwilioAccountSID field if non-nil, zero value otherwise.

### GetTwilioAccountSIDOk

`func (o *AddTwilioAlertHandlerRequest) GetTwilioAccountSIDOk() (*string, bool)`

GetTwilioAccountSIDOk returns a tuple with the TwilioAccountSID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAccountSID

`func (o *AddTwilioAlertHandlerRequest) SetTwilioAccountSID(v string)`

SetTwilioAccountSID sets TwilioAccountSID field to given value.


### GetTwilioAuthToken

`func (o *AddTwilioAlertHandlerRequest) GetTwilioAuthToken() string`

GetTwilioAuthToken returns the TwilioAuthToken field if non-nil, zero value otherwise.

### GetTwilioAuthTokenOk

`func (o *AddTwilioAlertHandlerRequest) GetTwilioAuthTokenOk() (*string, bool)`

GetTwilioAuthTokenOk returns a tuple with the TwilioAuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAuthToken

`func (o *AddTwilioAlertHandlerRequest) SetTwilioAuthToken(v string)`

SetTwilioAuthToken sets TwilioAuthToken field to given value.

### HasTwilioAuthToken

`func (o *AddTwilioAlertHandlerRequest) HasTwilioAuthToken() bool`

HasTwilioAuthToken returns a boolean if a field has been set.

### GetTwilioAuthTokenPassphraseProvider

`func (o *AddTwilioAlertHandlerRequest) GetTwilioAuthTokenPassphraseProvider() string`

GetTwilioAuthTokenPassphraseProvider returns the TwilioAuthTokenPassphraseProvider field if non-nil, zero value otherwise.

### GetTwilioAuthTokenPassphraseProviderOk

`func (o *AddTwilioAlertHandlerRequest) GetTwilioAuthTokenPassphraseProviderOk() (*string, bool)`

GetTwilioAuthTokenPassphraseProviderOk returns a tuple with the TwilioAuthTokenPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAuthTokenPassphraseProvider

`func (o *AddTwilioAlertHandlerRequest) SetTwilioAuthTokenPassphraseProvider(v string)`

SetTwilioAuthTokenPassphraseProvider sets TwilioAuthTokenPassphraseProvider field to given value.

### HasTwilioAuthTokenPassphraseProvider

`func (o *AddTwilioAlertHandlerRequest) HasTwilioAuthTokenPassphraseProvider() bool`

HasTwilioAuthTokenPassphraseProvider returns a boolean if a field has been set.

### GetSenderPhoneNumber

`func (o *AddTwilioAlertHandlerRequest) GetSenderPhoneNumber() []string`

GetSenderPhoneNumber returns the SenderPhoneNumber field if non-nil, zero value otherwise.

### GetSenderPhoneNumberOk

`func (o *AddTwilioAlertHandlerRequest) GetSenderPhoneNumberOk() (*[]string, bool)`

GetSenderPhoneNumberOk returns a tuple with the SenderPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderPhoneNumber

`func (o *AddTwilioAlertHandlerRequest) SetSenderPhoneNumber(v []string)`

SetSenderPhoneNumber sets SenderPhoneNumber field to given value.


### GetRecipientPhoneNumber

`func (o *AddTwilioAlertHandlerRequest) GetRecipientPhoneNumber() []string`

GetRecipientPhoneNumber returns the RecipientPhoneNumber field if non-nil, zero value otherwise.

### GetRecipientPhoneNumberOk

`func (o *AddTwilioAlertHandlerRequest) GetRecipientPhoneNumberOk() (*[]string, bool)`

GetRecipientPhoneNumberOk returns a tuple with the RecipientPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientPhoneNumber

`func (o *AddTwilioAlertHandlerRequest) SetRecipientPhoneNumber(v []string)`

SetRecipientPhoneNumber sets RecipientPhoneNumber field to given value.


### GetLongMessageBehavior

`func (o *AddTwilioAlertHandlerRequest) GetLongMessageBehavior() EnumalertHandlerLongMessageBehaviorProp`

GetLongMessageBehavior returns the LongMessageBehavior field if non-nil, zero value otherwise.

### GetLongMessageBehaviorOk

`func (o *AddTwilioAlertHandlerRequest) GetLongMessageBehaviorOk() (*EnumalertHandlerLongMessageBehaviorProp, bool)`

GetLongMessageBehaviorOk returns a tuple with the LongMessageBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongMessageBehavior

`func (o *AddTwilioAlertHandlerRequest) SetLongMessageBehavior(v EnumalertHandlerLongMessageBehaviorProp)`

SetLongMessageBehavior sets LongMessageBehavior field to given value.

### HasLongMessageBehavior

`func (o *AddTwilioAlertHandlerRequest) HasLongMessageBehavior() bool`

HasLongMessageBehavior returns a boolean if a field has been set.

### GetDescription

`func (o *AddTwilioAlertHandlerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddTwilioAlertHandlerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddTwilioAlertHandlerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddTwilioAlertHandlerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddTwilioAlertHandlerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddTwilioAlertHandlerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddTwilioAlertHandlerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEnabledAlertSeverity

`func (o *AddTwilioAlertHandlerRequest) GetEnabledAlertSeverity() []EnumalertHandlerEnabledAlertSeverityProp`

GetEnabledAlertSeverity returns the EnabledAlertSeverity field if non-nil, zero value otherwise.

### GetEnabledAlertSeverityOk

`func (o *AddTwilioAlertHandlerRequest) GetEnabledAlertSeverityOk() (*[]EnumalertHandlerEnabledAlertSeverityProp, bool)`

GetEnabledAlertSeverityOk returns a tuple with the EnabledAlertSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAlertSeverity

`func (o *AddTwilioAlertHandlerRequest) SetEnabledAlertSeverity(v []EnumalertHandlerEnabledAlertSeverityProp)`

SetEnabledAlertSeverity sets EnabledAlertSeverity field to given value.

### HasEnabledAlertSeverity

`func (o *AddTwilioAlertHandlerRequest) HasEnabledAlertSeverity() bool`

HasEnabledAlertSeverity returns a boolean if a field has been set.

### GetEnabledAlertType

`func (o *AddTwilioAlertHandlerRequest) GetEnabledAlertType() []EnumalertHandlerEnabledAlertTypeProp`

GetEnabledAlertType returns the EnabledAlertType field if non-nil, zero value otherwise.

### GetEnabledAlertTypeOk

`func (o *AddTwilioAlertHandlerRequest) GetEnabledAlertTypeOk() (*[]EnumalertHandlerEnabledAlertTypeProp, bool)`

GetEnabledAlertTypeOk returns a tuple with the EnabledAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAlertType

`func (o *AddTwilioAlertHandlerRequest) SetEnabledAlertType(v []EnumalertHandlerEnabledAlertTypeProp)`

SetEnabledAlertType sets EnabledAlertType field to given value.

### HasEnabledAlertType

`func (o *AddTwilioAlertHandlerRequest) HasEnabledAlertType() bool`

HasEnabledAlertType returns a boolean if a field has been set.

### GetDisabledAlertType

`func (o *AddTwilioAlertHandlerRequest) GetDisabledAlertType() []EnumalertHandlerDisabledAlertTypeProp`

GetDisabledAlertType returns the DisabledAlertType field if non-nil, zero value otherwise.

### GetDisabledAlertTypeOk

`func (o *AddTwilioAlertHandlerRequest) GetDisabledAlertTypeOk() (*[]EnumalertHandlerDisabledAlertTypeProp, bool)`

GetDisabledAlertTypeOk returns a tuple with the DisabledAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAlertType

`func (o *AddTwilioAlertHandlerRequest) SetDisabledAlertType(v []EnumalertHandlerDisabledAlertTypeProp)`

SetDisabledAlertType sets DisabledAlertType field to given value.

### HasDisabledAlertType

`func (o *AddTwilioAlertHandlerRequest) HasDisabledAlertType() bool`

HasDisabledAlertType returns a boolean if a field has been set.

### GetHandlerName

`func (o *AddTwilioAlertHandlerRequest) GetHandlerName() string`

GetHandlerName returns the HandlerName field if non-nil, zero value otherwise.

### GetHandlerNameOk

`func (o *AddTwilioAlertHandlerRequest) GetHandlerNameOk() (*string, bool)`

GetHandlerNameOk returns a tuple with the HandlerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerName

`func (o *AddTwilioAlertHandlerRequest) SetHandlerName(v string)`

SetHandlerName sets HandlerName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


