# TwilioAlertHandlerResponse

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
**LongMessageBehavior** | [**EnumalertHandlerLongMessageBehaviorProp**](EnumalertHandlerLongMessageBehaviorProp.md) |  | 
**Description** | Pointer to **string** | A description for this Alert Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Alert Handler is enabled. | 
**EnabledAlertSeverity** | Pointer to [**[]EnumalertHandlerEnabledAlertSeverityProp**](EnumalertHandlerEnabledAlertSeverityProp.md) |  | [optional] 
**EnabledAlertType** | Pointer to [**[]EnumalertHandlerEnabledAlertTypeProp**](EnumalertHandlerEnabledAlertTypeProp.md) |  | [optional] 
**DisabledAlertType** | Pointer to [**[]EnumalertHandlerDisabledAlertTypeProp**](EnumalertHandlerDisabledAlertTypeProp.md) |  | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Alert Handler | 

## Methods

### NewTwilioAlertHandlerResponse

`func NewTwilioAlertHandlerResponse(schemas []EnumtwilioAlertHandlerSchemaUrn, twilioAccountSID string, senderPhoneNumber []string, recipientPhoneNumber []string, longMessageBehavior EnumalertHandlerLongMessageBehaviorProp, enabled bool, id string, ) *TwilioAlertHandlerResponse`

NewTwilioAlertHandlerResponse instantiates a new TwilioAlertHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTwilioAlertHandlerResponseWithDefaults

`func NewTwilioAlertHandlerResponseWithDefaults() *TwilioAlertHandlerResponse`

NewTwilioAlertHandlerResponseWithDefaults instantiates a new TwilioAlertHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *TwilioAlertHandlerResponse) GetSchemas() []EnumtwilioAlertHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *TwilioAlertHandlerResponse) GetSchemasOk() (*[]EnumtwilioAlertHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *TwilioAlertHandlerResponse) SetSchemas(v []EnumtwilioAlertHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAsynchronous

`func (o *TwilioAlertHandlerResponse) GetAsynchronous() bool`

GetAsynchronous returns the Asynchronous field if non-nil, zero value otherwise.

### GetAsynchronousOk

`func (o *TwilioAlertHandlerResponse) GetAsynchronousOk() (*bool, bool)`

GetAsynchronousOk returns a tuple with the Asynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsynchronous

`func (o *TwilioAlertHandlerResponse) SetAsynchronous(v bool)`

SetAsynchronous sets Asynchronous field to given value.

### HasAsynchronous

`func (o *TwilioAlertHandlerResponse) HasAsynchronous() bool`

HasAsynchronous returns a boolean if a field has been set.

### GetHttpProxyExternalServer

`func (o *TwilioAlertHandlerResponse) GetHttpProxyExternalServer() string`

GetHttpProxyExternalServer returns the HttpProxyExternalServer field if non-nil, zero value otherwise.

### GetHttpProxyExternalServerOk

`func (o *TwilioAlertHandlerResponse) GetHttpProxyExternalServerOk() (*string, bool)`

GetHttpProxyExternalServerOk returns a tuple with the HttpProxyExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyExternalServer

`func (o *TwilioAlertHandlerResponse) SetHttpProxyExternalServer(v string)`

SetHttpProxyExternalServer sets HttpProxyExternalServer field to given value.

### HasHttpProxyExternalServer

`func (o *TwilioAlertHandlerResponse) HasHttpProxyExternalServer() bool`

HasHttpProxyExternalServer returns a boolean if a field has been set.

### GetTwilioAccountSID

`func (o *TwilioAlertHandlerResponse) GetTwilioAccountSID() string`

GetTwilioAccountSID returns the TwilioAccountSID field if non-nil, zero value otherwise.

### GetTwilioAccountSIDOk

`func (o *TwilioAlertHandlerResponse) GetTwilioAccountSIDOk() (*string, bool)`

GetTwilioAccountSIDOk returns a tuple with the TwilioAccountSID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAccountSID

`func (o *TwilioAlertHandlerResponse) SetTwilioAccountSID(v string)`

SetTwilioAccountSID sets TwilioAccountSID field to given value.


### GetTwilioAuthToken

`func (o *TwilioAlertHandlerResponse) GetTwilioAuthToken() string`

GetTwilioAuthToken returns the TwilioAuthToken field if non-nil, zero value otherwise.

### GetTwilioAuthTokenOk

`func (o *TwilioAlertHandlerResponse) GetTwilioAuthTokenOk() (*string, bool)`

GetTwilioAuthTokenOk returns a tuple with the TwilioAuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAuthToken

`func (o *TwilioAlertHandlerResponse) SetTwilioAuthToken(v string)`

SetTwilioAuthToken sets TwilioAuthToken field to given value.

### HasTwilioAuthToken

`func (o *TwilioAlertHandlerResponse) HasTwilioAuthToken() bool`

HasTwilioAuthToken returns a boolean if a field has been set.

### GetTwilioAuthTokenPassphraseProvider

`func (o *TwilioAlertHandlerResponse) GetTwilioAuthTokenPassphraseProvider() string`

GetTwilioAuthTokenPassphraseProvider returns the TwilioAuthTokenPassphraseProvider field if non-nil, zero value otherwise.

### GetTwilioAuthTokenPassphraseProviderOk

`func (o *TwilioAlertHandlerResponse) GetTwilioAuthTokenPassphraseProviderOk() (*string, bool)`

GetTwilioAuthTokenPassphraseProviderOk returns a tuple with the TwilioAuthTokenPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAuthTokenPassphraseProvider

`func (o *TwilioAlertHandlerResponse) SetTwilioAuthTokenPassphraseProvider(v string)`

SetTwilioAuthTokenPassphraseProvider sets TwilioAuthTokenPassphraseProvider field to given value.

### HasTwilioAuthTokenPassphraseProvider

`func (o *TwilioAlertHandlerResponse) HasTwilioAuthTokenPassphraseProvider() bool`

HasTwilioAuthTokenPassphraseProvider returns a boolean if a field has been set.

### GetSenderPhoneNumber

`func (o *TwilioAlertHandlerResponse) GetSenderPhoneNumber() []string`

GetSenderPhoneNumber returns the SenderPhoneNumber field if non-nil, zero value otherwise.

### GetSenderPhoneNumberOk

`func (o *TwilioAlertHandlerResponse) GetSenderPhoneNumberOk() (*[]string, bool)`

GetSenderPhoneNumberOk returns a tuple with the SenderPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderPhoneNumber

`func (o *TwilioAlertHandlerResponse) SetSenderPhoneNumber(v []string)`

SetSenderPhoneNumber sets SenderPhoneNumber field to given value.


### GetRecipientPhoneNumber

`func (o *TwilioAlertHandlerResponse) GetRecipientPhoneNumber() []string`

GetRecipientPhoneNumber returns the RecipientPhoneNumber field if non-nil, zero value otherwise.

### GetRecipientPhoneNumberOk

`func (o *TwilioAlertHandlerResponse) GetRecipientPhoneNumberOk() (*[]string, bool)`

GetRecipientPhoneNumberOk returns a tuple with the RecipientPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientPhoneNumber

`func (o *TwilioAlertHandlerResponse) SetRecipientPhoneNumber(v []string)`

SetRecipientPhoneNumber sets RecipientPhoneNumber field to given value.


### GetLongMessageBehavior

`func (o *TwilioAlertHandlerResponse) GetLongMessageBehavior() EnumalertHandlerLongMessageBehaviorProp`

GetLongMessageBehavior returns the LongMessageBehavior field if non-nil, zero value otherwise.

### GetLongMessageBehaviorOk

`func (o *TwilioAlertHandlerResponse) GetLongMessageBehaviorOk() (*EnumalertHandlerLongMessageBehaviorProp, bool)`

GetLongMessageBehaviorOk returns a tuple with the LongMessageBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongMessageBehavior

`func (o *TwilioAlertHandlerResponse) SetLongMessageBehavior(v EnumalertHandlerLongMessageBehaviorProp)`

SetLongMessageBehavior sets LongMessageBehavior field to given value.


### GetDescription

`func (o *TwilioAlertHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TwilioAlertHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TwilioAlertHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TwilioAlertHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *TwilioAlertHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TwilioAlertHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TwilioAlertHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEnabledAlertSeverity

`func (o *TwilioAlertHandlerResponse) GetEnabledAlertSeverity() []EnumalertHandlerEnabledAlertSeverityProp`

GetEnabledAlertSeverity returns the EnabledAlertSeverity field if non-nil, zero value otherwise.

### GetEnabledAlertSeverityOk

`func (o *TwilioAlertHandlerResponse) GetEnabledAlertSeverityOk() (*[]EnumalertHandlerEnabledAlertSeverityProp, bool)`

GetEnabledAlertSeverityOk returns a tuple with the EnabledAlertSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAlertSeverity

`func (o *TwilioAlertHandlerResponse) SetEnabledAlertSeverity(v []EnumalertHandlerEnabledAlertSeverityProp)`

SetEnabledAlertSeverity sets EnabledAlertSeverity field to given value.

### HasEnabledAlertSeverity

`func (o *TwilioAlertHandlerResponse) HasEnabledAlertSeverity() bool`

HasEnabledAlertSeverity returns a boolean if a field has been set.

### GetEnabledAlertType

`func (o *TwilioAlertHandlerResponse) GetEnabledAlertType() []EnumalertHandlerEnabledAlertTypeProp`

GetEnabledAlertType returns the EnabledAlertType field if non-nil, zero value otherwise.

### GetEnabledAlertTypeOk

`func (o *TwilioAlertHandlerResponse) GetEnabledAlertTypeOk() (*[]EnumalertHandlerEnabledAlertTypeProp, bool)`

GetEnabledAlertTypeOk returns a tuple with the EnabledAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAlertType

`func (o *TwilioAlertHandlerResponse) SetEnabledAlertType(v []EnumalertHandlerEnabledAlertTypeProp)`

SetEnabledAlertType sets EnabledAlertType field to given value.

### HasEnabledAlertType

`func (o *TwilioAlertHandlerResponse) HasEnabledAlertType() bool`

HasEnabledAlertType returns a boolean if a field has been set.

### GetDisabledAlertType

`func (o *TwilioAlertHandlerResponse) GetDisabledAlertType() []EnumalertHandlerDisabledAlertTypeProp`

GetDisabledAlertType returns the DisabledAlertType field if non-nil, zero value otherwise.

### GetDisabledAlertTypeOk

`func (o *TwilioAlertHandlerResponse) GetDisabledAlertTypeOk() (*[]EnumalertHandlerDisabledAlertTypeProp, bool)`

GetDisabledAlertTypeOk returns a tuple with the DisabledAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAlertType

`func (o *TwilioAlertHandlerResponse) SetDisabledAlertType(v []EnumalertHandlerDisabledAlertTypeProp)`

SetDisabledAlertType sets DisabledAlertType field to given value.

### HasDisabledAlertType

`func (o *TwilioAlertHandlerResponse) HasDisabledAlertType() bool`

HasDisabledAlertType returns a boolean if a field has been set.

### GetMeta

`func (o *TwilioAlertHandlerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TwilioAlertHandlerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TwilioAlertHandlerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TwilioAlertHandlerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *TwilioAlertHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *TwilioAlertHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *TwilioAlertHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *TwilioAlertHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *TwilioAlertHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TwilioAlertHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TwilioAlertHandlerResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


