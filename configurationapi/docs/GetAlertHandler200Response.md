# GetAlertHandler200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumthirdPartyAlertHandlerSchemaUrn**](EnumthirdPartyAlertHandlerSchemaUrn.md) |  | 
**Id** | **string** | Name of the Alert Handler | 
**OutputLocation** | Pointer to [**EnumalertHandlerOutputLocationProp**](EnumalertHandlerOutputLocationProp.md) |  | [optional] 
**OutputFormat** | Pointer to [**EnumalertHandlerOutputFormatProp**](EnumalertHandlerOutputFormatProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Alert Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Alert Handler is enabled. | 
**Asynchronous** | Pointer to **bool** | Indicates whether the server should attempt to invoke this Alert Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver the alert notification) will not delay whatever processing the server was performing when the alert was generated. | [optional] 
**EnabledAlertSeverity** | Pointer to [**[]EnumalertHandlerEnabledAlertSeverityProp**](EnumalertHandlerEnabledAlertSeverityProp.md) |  | [optional] 
**EnabledAlertType** | Pointer to [**[]EnumalertHandlerEnabledAlertTypeProp**](EnumalertHandlerEnabledAlertTypeProp.md) |  | [optional] 
**DisabledAlertType** | Pointer to [**[]EnumalertHandlerDisabledAlertTypeProp**](EnumalertHandlerDisabledAlertTypeProp.md) |  | [optional] 
**SenderAddress** | **string** | Specifies the email address to use as the sender for messages generated by this alert handler. | 
**RecipientAddress** | **[]string** | Specifies an email address to which the messages should be sent. | 
**MessageSubject** | **string** | Specifies the subject that should be used for email messages generated by this alert handler. | 
**MessageBody** | **string** | Specifies the body that should be used for email messages generated by this alert handler. | 
**IncludeMonitorDataFilter** | Pointer to **string** | Include monitor entries that match this filter. | [optional] 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Alert Handler. | 
**ScriptArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Scripted Alert Handler. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**ServerHostName** | **string** | Specifies the address of the SNMP agent to which traps will be sent. | 
**ServerPort** | **int32** | Specifies the port number of the SNMP agent to which traps will be sent. | 
**CommunityName** | **string** | Specifies the name of the community to which the traps will be sent. | 
**TwilioAccountSID** | **string** | The unique identifier assigned to the Twilio account that will be used. | 
**TwilioAuthToken** | Pointer to **string** | The auth token for the Twilio account that will be used. | [optional] 
**TwilioAuthTokenPassphraseProvider** | Pointer to **string** | The passphrase provider that may be used to obtain the auth token for the Twilio account that will be used. | [optional] 
**SenderPhoneNumber** | **[]string** | The outgoing phone number to use for the messages. Values must be phone numbers you have obtained for use with your Twilio account. | 
**RecipientPhoneNumber** | **[]string** | The phone number to which alert notifications should be delivered. | 
**LongMessageBehavior** | [**EnumalertHandlerLongMessageBehaviorProp**](EnumalertHandlerLongMessageBehaviorProp.md) |  | 
**Command** | **string** | Specifies the path of the command to execute, without any arguments. It must be an absolute path for reasons of security and reliability. | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Alert Handler. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Alert Handler. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewGetAlertHandler200Response

`func NewGetAlertHandler200Response(schemas []EnumthirdPartyAlertHandlerSchemaUrn, id string, enabled bool, senderAddress string, recipientAddress []string, messageSubject string, messageBody string, scriptClass string, serverHostName string, serverPort int32, communityName string, twilioAccountSID string, senderPhoneNumber []string, recipientPhoneNumber []string, longMessageBehavior EnumalertHandlerLongMessageBehaviorProp, command string, extensionClass string, ) *GetAlertHandler200Response`

NewGetAlertHandler200Response instantiates a new GetAlertHandler200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAlertHandler200ResponseWithDefaults

`func NewGetAlertHandler200ResponseWithDefaults() *GetAlertHandler200Response`

NewGetAlertHandler200ResponseWithDefaults instantiates a new GetAlertHandler200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetAlertHandler200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAlertHandler200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAlertHandler200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetAlertHandler200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *GetAlertHandler200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *GetAlertHandler200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *GetAlertHandler200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *GetAlertHandler200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *GetAlertHandler200Response) GetSchemas() []EnumthirdPartyAlertHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GetAlertHandler200Response) GetSchemasOk() (*[]EnumthirdPartyAlertHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GetAlertHandler200Response) SetSchemas(v []EnumthirdPartyAlertHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *GetAlertHandler200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAlertHandler200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAlertHandler200Response) SetId(v string)`

SetId sets Id field to given value.


### GetOutputLocation

`func (o *GetAlertHandler200Response) GetOutputLocation() EnumalertHandlerOutputLocationProp`

GetOutputLocation returns the OutputLocation field if non-nil, zero value otherwise.

### GetOutputLocationOk

`func (o *GetAlertHandler200Response) GetOutputLocationOk() (*EnumalertHandlerOutputLocationProp, bool)`

GetOutputLocationOk returns a tuple with the OutputLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputLocation

`func (o *GetAlertHandler200Response) SetOutputLocation(v EnumalertHandlerOutputLocationProp)`

SetOutputLocation sets OutputLocation field to given value.

### HasOutputLocation

`func (o *GetAlertHandler200Response) HasOutputLocation() bool`

HasOutputLocation returns a boolean if a field has been set.

### GetOutputFormat

`func (o *GetAlertHandler200Response) GetOutputFormat() EnumalertHandlerOutputFormatProp`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *GetAlertHandler200Response) GetOutputFormatOk() (*EnumalertHandlerOutputFormatProp, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *GetAlertHandler200Response) SetOutputFormat(v EnumalertHandlerOutputFormatProp)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *GetAlertHandler200Response) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

### GetDescription

`func (o *GetAlertHandler200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetAlertHandler200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetAlertHandler200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetAlertHandler200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *GetAlertHandler200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetAlertHandler200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetAlertHandler200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAsynchronous

`func (o *GetAlertHandler200Response) GetAsynchronous() bool`

GetAsynchronous returns the Asynchronous field if non-nil, zero value otherwise.

### GetAsynchronousOk

`func (o *GetAlertHandler200Response) GetAsynchronousOk() (*bool, bool)`

GetAsynchronousOk returns a tuple with the Asynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsynchronous

`func (o *GetAlertHandler200Response) SetAsynchronous(v bool)`

SetAsynchronous sets Asynchronous field to given value.

### HasAsynchronous

`func (o *GetAlertHandler200Response) HasAsynchronous() bool`

HasAsynchronous returns a boolean if a field has been set.

### GetEnabledAlertSeverity

`func (o *GetAlertHandler200Response) GetEnabledAlertSeverity() []EnumalertHandlerEnabledAlertSeverityProp`

GetEnabledAlertSeverity returns the EnabledAlertSeverity field if non-nil, zero value otherwise.

### GetEnabledAlertSeverityOk

`func (o *GetAlertHandler200Response) GetEnabledAlertSeverityOk() (*[]EnumalertHandlerEnabledAlertSeverityProp, bool)`

GetEnabledAlertSeverityOk returns a tuple with the EnabledAlertSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAlertSeverity

`func (o *GetAlertHandler200Response) SetEnabledAlertSeverity(v []EnumalertHandlerEnabledAlertSeverityProp)`

SetEnabledAlertSeverity sets EnabledAlertSeverity field to given value.

### HasEnabledAlertSeverity

`func (o *GetAlertHandler200Response) HasEnabledAlertSeverity() bool`

HasEnabledAlertSeverity returns a boolean if a field has been set.

### GetEnabledAlertType

`func (o *GetAlertHandler200Response) GetEnabledAlertType() []EnumalertHandlerEnabledAlertTypeProp`

GetEnabledAlertType returns the EnabledAlertType field if non-nil, zero value otherwise.

### GetEnabledAlertTypeOk

`func (o *GetAlertHandler200Response) GetEnabledAlertTypeOk() (*[]EnumalertHandlerEnabledAlertTypeProp, bool)`

GetEnabledAlertTypeOk returns a tuple with the EnabledAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAlertType

`func (o *GetAlertHandler200Response) SetEnabledAlertType(v []EnumalertHandlerEnabledAlertTypeProp)`

SetEnabledAlertType sets EnabledAlertType field to given value.

### HasEnabledAlertType

`func (o *GetAlertHandler200Response) HasEnabledAlertType() bool`

HasEnabledAlertType returns a boolean if a field has been set.

### GetDisabledAlertType

`func (o *GetAlertHandler200Response) GetDisabledAlertType() []EnumalertHandlerDisabledAlertTypeProp`

GetDisabledAlertType returns the DisabledAlertType field if non-nil, zero value otherwise.

### GetDisabledAlertTypeOk

`func (o *GetAlertHandler200Response) GetDisabledAlertTypeOk() (*[]EnumalertHandlerDisabledAlertTypeProp, bool)`

GetDisabledAlertTypeOk returns a tuple with the DisabledAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAlertType

`func (o *GetAlertHandler200Response) SetDisabledAlertType(v []EnumalertHandlerDisabledAlertTypeProp)`

SetDisabledAlertType sets DisabledAlertType field to given value.

### HasDisabledAlertType

`func (o *GetAlertHandler200Response) HasDisabledAlertType() bool`

HasDisabledAlertType returns a boolean if a field has been set.

### GetSenderAddress

`func (o *GetAlertHandler200Response) GetSenderAddress() string`

GetSenderAddress returns the SenderAddress field if non-nil, zero value otherwise.

### GetSenderAddressOk

`func (o *GetAlertHandler200Response) GetSenderAddressOk() (*string, bool)`

GetSenderAddressOk returns a tuple with the SenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddress

`func (o *GetAlertHandler200Response) SetSenderAddress(v string)`

SetSenderAddress sets SenderAddress field to given value.


### GetRecipientAddress

`func (o *GetAlertHandler200Response) GetRecipientAddress() []string`

GetRecipientAddress returns the RecipientAddress field if non-nil, zero value otherwise.

### GetRecipientAddressOk

`func (o *GetAlertHandler200Response) GetRecipientAddressOk() (*[]string, bool)`

GetRecipientAddressOk returns a tuple with the RecipientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientAddress

`func (o *GetAlertHandler200Response) SetRecipientAddress(v []string)`

SetRecipientAddress sets RecipientAddress field to given value.


### GetMessageSubject

`func (o *GetAlertHandler200Response) GetMessageSubject() string`

GetMessageSubject returns the MessageSubject field if non-nil, zero value otherwise.

### GetMessageSubjectOk

`func (o *GetAlertHandler200Response) GetMessageSubjectOk() (*string, bool)`

GetMessageSubjectOk returns a tuple with the MessageSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSubject

`func (o *GetAlertHandler200Response) SetMessageSubject(v string)`

SetMessageSubject sets MessageSubject field to given value.


### GetMessageBody

`func (o *GetAlertHandler200Response) GetMessageBody() string`

GetMessageBody returns the MessageBody field if non-nil, zero value otherwise.

### GetMessageBodyOk

`func (o *GetAlertHandler200Response) GetMessageBodyOk() (*string, bool)`

GetMessageBodyOk returns a tuple with the MessageBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageBody

`func (o *GetAlertHandler200Response) SetMessageBody(v string)`

SetMessageBody sets MessageBody field to given value.


### GetIncludeMonitorDataFilter

`func (o *GetAlertHandler200Response) GetIncludeMonitorDataFilter() string`

GetIncludeMonitorDataFilter returns the IncludeMonitorDataFilter field if non-nil, zero value otherwise.

### GetIncludeMonitorDataFilterOk

`func (o *GetAlertHandler200Response) GetIncludeMonitorDataFilterOk() (*string, bool)`

GetIncludeMonitorDataFilterOk returns a tuple with the IncludeMonitorDataFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMonitorDataFilter

`func (o *GetAlertHandler200Response) SetIncludeMonitorDataFilter(v string)`

SetIncludeMonitorDataFilter sets IncludeMonitorDataFilter field to given value.

### HasIncludeMonitorDataFilter

`func (o *GetAlertHandler200Response) HasIncludeMonitorDataFilter() bool`

HasIncludeMonitorDataFilter returns a boolean if a field has been set.

### GetScriptClass

`func (o *GetAlertHandler200Response) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *GetAlertHandler200Response) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *GetAlertHandler200Response) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *GetAlertHandler200Response) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *GetAlertHandler200Response) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *GetAlertHandler200Response) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *GetAlertHandler200Response) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetServerHostName

`func (o *GetAlertHandler200Response) GetServerHostName() string`

GetServerHostName returns the ServerHostName field if non-nil, zero value otherwise.

### GetServerHostNameOk

`func (o *GetAlertHandler200Response) GetServerHostNameOk() (*string, bool)`

GetServerHostNameOk returns a tuple with the ServerHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostName

`func (o *GetAlertHandler200Response) SetServerHostName(v string)`

SetServerHostName sets ServerHostName field to given value.


### GetServerPort

`func (o *GetAlertHandler200Response) GetServerPort() int32`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *GetAlertHandler200Response) GetServerPortOk() (*int32, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *GetAlertHandler200Response) SetServerPort(v int32)`

SetServerPort sets ServerPort field to given value.


### GetCommunityName

`func (o *GetAlertHandler200Response) GetCommunityName() string`

GetCommunityName returns the CommunityName field if non-nil, zero value otherwise.

### GetCommunityNameOk

`func (o *GetAlertHandler200Response) GetCommunityNameOk() (*string, bool)`

GetCommunityNameOk returns a tuple with the CommunityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityName

`func (o *GetAlertHandler200Response) SetCommunityName(v string)`

SetCommunityName sets CommunityName field to given value.


### GetTwilioAccountSID

`func (o *GetAlertHandler200Response) GetTwilioAccountSID() string`

GetTwilioAccountSID returns the TwilioAccountSID field if non-nil, zero value otherwise.

### GetTwilioAccountSIDOk

`func (o *GetAlertHandler200Response) GetTwilioAccountSIDOk() (*string, bool)`

GetTwilioAccountSIDOk returns a tuple with the TwilioAccountSID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAccountSID

`func (o *GetAlertHandler200Response) SetTwilioAccountSID(v string)`

SetTwilioAccountSID sets TwilioAccountSID field to given value.


### GetTwilioAuthToken

`func (o *GetAlertHandler200Response) GetTwilioAuthToken() string`

GetTwilioAuthToken returns the TwilioAuthToken field if non-nil, zero value otherwise.

### GetTwilioAuthTokenOk

`func (o *GetAlertHandler200Response) GetTwilioAuthTokenOk() (*string, bool)`

GetTwilioAuthTokenOk returns a tuple with the TwilioAuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAuthToken

`func (o *GetAlertHandler200Response) SetTwilioAuthToken(v string)`

SetTwilioAuthToken sets TwilioAuthToken field to given value.

### HasTwilioAuthToken

`func (o *GetAlertHandler200Response) HasTwilioAuthToken() bool`

HasTwilioAuthToken returns a boolean if a field has been set.

### GetTwilioAuthTokenPassphraseProvider

`func (o *GetAlertHandler200Response) GetTwilioAuthTokenPassphraseProvider() string`

GetTwilioAuthTokenPassphraseProvider returns the TwilioAuthTokenPassphraseProvider field if non-nil, zero value otherwise.

### GetTwilioAuthTokenPassphraseProviderOk

`func (o *GetAlertHandler200Response) GetTwilioAuthTokenPassphraseProviderOk() (*string, bool)`

GetTwilioAuthTokenPassphraseProviderOk returns a tuple with the TwilioAuthTokenPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAuthTokenPassphraseProvider

`func (o *GetAlertHandler200Response) SetTwilioAuthTokenPassphraseProvider(v string)`

SetTwilioAuthTokenPassphraseProvider sets TwilioAuthTokenPassphraseProvider field to given value.

### HasTwilioAuthTokenPassphraseProvider

`func (o *GetAlertHandler200Response) HasTwilioAuthTokenPassphraseProvider() bool`

HasTwilioAuthTokenPassphraseProvider returns a boolean if a field has been set.

### GetSenderPhoneNumber

`func (o *GetAlertHandler200Response) GetSenderPhoneNumber() []string`

GetSenderPhoneNumber returns the SenderPhoneNumber field if non-nil, zero value otherwise.

### GetSenderPhoneNumberOk

`func (o *GetAlertHandler200Response) GetSenderPhoneNumberOk() (*[]string, bool)`

GetSenderPhoneNumberOk returns a tuple with the SenderPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderPhoneNumber

`func (o *GetAlertHandler200Response) SetSenderPhoneNumber(v []string)`

SetSenderPhoneNumber sets SenderPhoneNumber field to given value.


### GetRecipientPhoneNumber

`func (o *GetAlertHandler200Response) GetRecipientPhoneNumber() []string`

GetRecipientPhoneNumber returns the RecipientPhoneNumber field if non-nil, zero value otherwise.

### GetRecipientPhoneNumberOk

`func (o *GetAlertHandler200Response) GetRecipientPhoneNumberOk() (*[]string, bool)`

GetRecipientPhoneNumberOk returns a tuple with the RecipientPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientPhoneNumber

`func (o *GetAlertHandler200Response) SetRecipientPhoneNumber(v []string)`

SetRecipientPhoneNumber sets RecipientPhoneNumber field to given value.


### GetLongMessageBehavior

`func (o *GetAlertHandler200Response) GetLongMessageBehavior() EnumalertHandlerLongMessageBehaviorProp`

GetLongMessageBehavior returns the LongMessageBehavior field if non-nil, zero value otherwise.

### GetLongMessageBehaviorOk

`func (o *GetAlertHandler200Response) GetLongMessageBehaviorOk() (*EnumalertHandlerLongMessageBehaviorProp, bool)`

GetLongMessageBehaviorOk returns a tuple with the LongMessageBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongMessageBehavior

`func (o *GetAlertHandler200Response) SetLongMessageBehavior(v EnumalertHandlerLongMessageBehaviorProp)`

SetLongMessageBehavior sets LongMessageBehavior field to given value.


### GetCommand

`func (o *GetAlertHandler200Response) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *GetAlertHandler200Response) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *GetAlertHandler200Response) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetExtensionClass

`func (o *GetAlertHandler200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *GetAlertHandler200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *GetAlertHandler200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *GetAlertHandler200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *GetAlertHandler200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *GetAlertHandler200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *GetAlertHandler200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

