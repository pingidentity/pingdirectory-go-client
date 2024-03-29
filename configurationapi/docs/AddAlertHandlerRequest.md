# AddAlertHandlerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HandlerName** | **string** | Name of the new Alert Handler | 
**Schemas** | [**[]EnumthirdPartyAlertHandlerSchemaUrn**](EnumthirdPartyAlertHandlerSchemaUrn.md) |  | 
**Asynchronous** | Pointer to **bool** | Indicates whether the server should attempt to invoke this Alert Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver the alert notification) will not delay whatever processing the server was performing when the alert was generated. | [optional] 
**SenderAddress** | **string** | Specifies the email address to use as the sender for messages generated by this alert handler. | 
**RecipientAddress** | **[]string** | Specifies an email address to which the messages should be sent. | 
**MessageSubject** | Pointer to **string** | Specifies the subject that should be used for email messages generated by this alert handler. | [optional] 
**MessageBody** | Pointer to **string** | Specifies the body that should be used for email messages generated by this alert handler. | [optional] 
**IncludeMonitorDataFilter** | Pointer to **string** | Include monitor entries that match this filter. | [optional] 
**Description** | Pointer to **string** | A description for this Alert Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Alert Handler is enabled. | 
**EnabledAlertSeverity** | Pointer to [**[]EnumalertHandlerEnabledAlertSeverityProp**](EnumalertHandlerEnabledAlertSeverityProp.md) |  | [optional] 
**EnabledAlertType** | Pointer to [**[]EnumalertHandlerEnabledAlertTypeProp**](EnumalertHandlerEnabledAlertTypeProp.md) |  | [optional] 
**DisabledAlertType** | Pointer to [**[]EnumalertHandlerDisabledAlertTypeProp**](EnumalertHandlerDisabledAlertTypeProp.md) |  | [optional] 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Alert Handler. | 
**ScriptArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Scripted Alert Handler. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**ServerHostName** | **string** | Specifies the address of the SNMP agent to which traps will be sent. | 
**ServerPort** | Pointer to **int64** | Specifies the port number of the SNMP agent to which traps will be sent. | [optional] 
**CommunityName** | Pointer to **string** | Specifies the name of the community to which the traps will be sent. | [optional] 
**HttpProxyExternalServer** | Pointer to **string** | A reference to an HTTP proxy server that should be used for requests sent to the Twilio service. | [optional] 
**TwilioAccountSID** | **string** | The unique identifier assigned to the Twilio account that will be used. | 
**TwilioAuthToken** | Pointer to **string** | The auth token for the Twilio account that will be used. | [optional] 
**TwilioAuthTokenPassphraseProvider** | Pointer to **string** | The passphrase provider that may be used to obtain the auth token for the Twilio account that will be used. | [optional] 
**SenderPhoneNumber** | **[]string** | The outgoing phone number to use for the messages. Values must be phone numbers you have obtained for use with your Twilio account. | 
**RecipientPhoneNumber** | **[]string** | The phone number to which alert notifications should be delivered. | 
**LongMessageBehavior** | Pointer to [**EnumalertHandlerLongMessageBehaviorProp**](EnumalertHandlerLongMessageBehaviorProp.md) |  | [optional] 
**Command** | **string** | Specifies the path of the command to execute, without any arguments. It must be an absolute path for reasons of security and reliability. | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Alert Handler. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Alert Handler. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddAlertHandlerRequest

`func NewAddAlertHandlerRequest(handlerName string, schemas []EnumthirdPartyAlertHandlerSchemaUrn, senderAddress string, recipientAddress []string, enabled bool, scriptClass string, serverHostName string, twilioAccountSID string, senderPhoneNumber []string, recipientPhoneNumber []string, command string, extensionClass string, ) *AddAlertHandlerRequest`

NewAddAlertHandlerRequest instantiates a new AddAlertHandlerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAlertHandlerRequestWithDefaults

`func NewAddAlertHandlerRequestWithDefaults() *AddAlertHandlerRequest`

NewAddAlertHandlerRequestWithDefaults instantiates a new AddAlertHandlerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandlerName

`func (o *AddAlertHandlerRequest) GetHandlerName() string`

GetHandlerName returns the HandlerName field if non-nil, zero value otherwise.

### GetHandlerNameOk

`func (o *AddAlertHandlerRequest) GetHandlerNameOk() (*string, bool)`

GetHandlerNameOk returns a tuple with the HandlerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerName

`func (o *AddAlertHandlerRequest) SetHandlerName(v string)`

SetHandlerName sets HandlerName field to given value.


### GetSchemas

`func (o *AddAlertHandlerRequest) GetSchemas() []EnumthirdPartyAlertHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddAlertHandlerRequest) GetSchemasOk() (*[]EnumthirdPartyAlertHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddAlertHandlerRequest) SetSchemas(v []EnumthirdPartyAlertHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAsynchronous

`func (o *AddAlertHandlerRequest) GetAsynchronous() bool`

GetAsynchronous returns the Asynchronous field if non-nil, zero value otherwise.

### GetAsynchronousOk

`func (o *AddAlertHandlerRequest) GetAsynchronousOk() (*bool, bool)`

GetAsynchronousOk returns a tuple with the Asynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsynchronous

`func (o *AddAlertHandlerRequest) SetAsynchronous(v bool)`

SetAsynchronous sets Asynchronous field to given value.

### HasAsynchronous

`func (o *AddAlertHandlerRequest) HasAsynchronous() bool`

HasAsynchronous returns a boolean if a field has been set.

### GetSenderAddress

`func (o *AddAlertHandlerRequest) GetSenderAddress() string`

GetSenderAddress returns the SenderAddress field if non-nil, zero value otherwise.

### GetSenderAddressOk

`func (o *AddAlertHandlerRequest) GetSenderAddressOk() (*string, bool)`

GetSenderAddressOk returns a tuple with the SenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddress

`func (o *AddAlertHandlerRequest) SetSenderAddress(v string)`

SetSenderAddress sets SenderAddress field to given value.


### GetRecipientAddress

`func (o *AddAlertHandlerRequest) GetRecipientAddress() []string`

GetRecipientAddress returns the RecipientAddress field if non-nil, zero value otherwise.

### GetRecipientAddressOk

`func (o *AddAlertHandlerRequest) GetRecipientAddressOk() (*[]string, bool)`

GetRecipientAddressOk returns a tuple with the RecipientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientAddress

`func (o *AddAlertHandlerRequest) SetRecipientAddress(v []string)`

SetRecipientAddress sets RecipientAddress field to given value.


### GetMessageSubject

`func (o *AddAlertHandlerRequest) GetMessageSubject() string`

GetMessageSubject returns the MessageSubject field if non-nil, zero value otherwise.

### GetMessageSubjectOk

`func (o *AddAlertHandlerRequest) GetMessageSubjectOk() (*string, bool)`

GetMessageSubjectOk returns a tuple with the MessageSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSubject

`func (o *AddAlertHandlerRequest) SetMessageSubject(v string)`

SetMessageSubject sets MessageSubject field to given value.

### HasMessageSubject

`func (o *AddAlertHandlerRequest) HasMessageSubject() bool`

HasMessageSubject returns a boolean if a field has been set.

### GetMessageBody

`func (o *AddAlertHandlerRequest) GetMessageBody() string`

GetMessageBody returns the MessageBody field if non-nil, zero value otherwise.

### GetMessageBodyOk

`func (o *AddAlertHandlerRequest) GetMessageBodyOk() (*string, bool)`

GetMessageBodyOk returns a tuple with the MessageBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageBody

`func (o *AddAlertHandlerRequest) SetMessageBody(v string)`

SetMessageBody sets MessageBody field to given value.

### HasMessageBody

`func (o *AddAlertHandlerRequest) HasMessageBody() bool`

HasMessageBody returns a boolean if a field has been set.

### GetIncludeMonitorDataFilter

`func (o *AddAlertHandlerRequest) GetIncludeMonitorDataFilter() string`

GetIncludeMonitorDataFilter returns the IncludeMonitorDataFilter field if non-nil, zero value otherwise.

### GetIncludeMonitorDataFilterOk

`func (o *AddAlertHandlerRequest) GetIncludeMonitorDataFilterOk() (*string, bool)`

GetIncludeMonitorDataFilterOk returns a tuple with the IncludeMonitorDataFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMonitorDataFilter

`func (o *AddAlertHandlerRequest) SetIncludeMonitorDataFilter(v string)`

SetIncludeMonitorDataFilter sets IncludeMonitorDataFilter field to given value.

### HasIncludeMonitorDataFilter

`func (o *AddAlertHandlerRequest) HasIncludeMonitorDataFilter() bool`

HasIncludeMonitorDataFilter returns a boolean if a field has been set.

### GetDescription

`func (o *AddAlertHandlerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddAlertHandlerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddAlertHandlerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddAlertHandlerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddAlertHandlerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddAlertHandlerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddAlertHandlerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEnabledAlertSeverity

`func (o *AddAlertHandlerRequest) GetEnabledAlertSeverity() []EnumalertHandlerEnabledAlertSeverityProp`

GetEnabledAlertSeverity returns the EnabledAlertSeverity field if non-nil, zero value otherwise.

### GetEnabledAlertSeverityOk

`func (o *AddAlertHandlerRequest) GetEnabledAlertSeverityOk() (*[]EnumalertHandlerEnabledAlertSeverityProp, bool)`

GetEnabledAlertSeverityOk returns a tuple with the EnabledAlertSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAlertSeverity

`func (o *AddAlertHandlerRequest) SetEnabledAlertSeverity(v []EnumalertHandlerEnabledAlertSeverityProp)`

SetEnabledAlertSeverity sets EnabledAlertSeverity field to given value.

### HasEnabledAlertSeverity

`func (o *AddAlertHandlerRequest) HasEnabledAlertSeverity() bool`

HasEnabledAlertSeverity returns a boolean if a field has been set.

### GetEnabledAlertType

`func (o *AddAlertHandlerRequest) GetEnabledAlertType() []EnumalertHandlerEnabledAlertTypeProp`

GetEnabledAlertType returns the EnabledAlertType field if non-nil, zero value otherwise.

### GetEnabledAlertTypeOk

`func (o *AddAlertHandlerRequest) GetEnabledAlertTypeOk() (*[]EnumalertHandlerEnabledAlertTypeProp, bool)`

GetEnabledAlertTypeOk returns a tuple with the EnabledAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAlertType

`func (o *AddAlertHandlerRequest) SetEnabledAlertType(v []EnumalertHandlerEnabledAlertTypeProp)`

SetEnabledAlertType sets EnabledAlertType field to given value.

### HasEnabledAlertType

`func (o *AddAlertHandlerRequest) HasEnabledAlertType() bool`

HasEnabledAlertType returns a boolean if a field has been set.

### GetDisabledAlertType

`func (o *AddAlertHandlerRequest) GetDisabledAlertType() []EnumalertHandlerDisabledAlertTypeProp`

GetDisabledAlertType returns the DisabledAlertType field if non-nil, zero value otherwise.

### GetDisabledAlertTypeOk

`func (o *AddAlertHandlerRequest) GetDisabledAlertTypeOk() (*[]EnumalertHandlerDisabledAlertTypeProp, bool)`

GetDisabledAlertTypeOk returns a tuple with the DisabledAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAlertType

`func (o *AddAlertHandlerRequest) SetDisabledAlertType(v []EnumalertHandlerDisabledAlertTypeProp)`

SetDisabledAlertType sets DisabledAlertType field to given value.

### HasDisabledAlertType

`func (o *AddAlertHandlerRequest) HasDisabledAlertType() bool`

HasDisabledAlertType returns a boolean if a field has been set.

### GetScriptClass

`func (o *AddAlertHandlerRequest) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *AddAlertHandlerRequest) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *AddAlertHandlerRequest) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *AddAlertHandlerRequest) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *AddAlertHandlerRequest) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *AddAlertHandlerRequest) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *AddAlertHandlerRequest) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetServerHostName

`func (o *AddAlertHandlerRequest) GetServerHostName() string`

GetServerHostName returns the ServerHostName field if non-nil, zero value otherwise.

### GetServerHostNameOk

`func (o *AddAlertHandlerRequest) GetServerHostNameOk() (*string, bool)`

GetServerHostNameOk returns a tuple with the ServerHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostName

`func (o *AddAlertHandlerRequest) SetServerHostName(v string)`

SetServerHostName sets ServerHostName field to given value.


### GetServerPort

`func (o *AddAlertHandlerRequest) GetServerPort() int64`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *AddAlertHandlerRequest) GetServerPortOk() (*int64, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *AddAlertHandlerRequest) SetServerPort(v int64)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *AddAlertHandlerRequest) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### GetCommunityName

`func (o *AddAlertHandlerRequest) GetCommunityName() string`

GetCommunityName returns the CommunityName field if non-nil, zero value otherwise.

### GetCommunityNameOk

`func (o *AddAlertHandlerRequest) GetCommunityNameOk() (*string, bool)`

GetCommunityNameOk returns a tuple with the CommunityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityName

`func (o *AddAlertHandlerRequest) SetCommunityName(v string)`

SetCommunityName sets CommunityName field to given value.

### HasCommunityName

`func (o *AddAlertHandlerRequest) HasCommunityName() bool`

HasCommunityName returns a boolean if a field has been set.

### GetHttpProxyExternalServer

`func (o *AddAlertHandlerRequest) GetHttpProxyExternalServer() string`

GetHttpProxyExternalServer returns the HttpProxyExternalServer field if non-nil, zero value otherwise.

### GetHttpProxyExternalServerOk

`func (o *AddAlertHandlerRequest) GetHttpProxyExternalServerOk() (*string, bool)`

GetHttpProxyExternalServerOk returns a tuple with the HttpProxyExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyExternalServer

`func (o *AddAlertHandlerRequest) SetHttpProxyExternalServer(v string)`

SetHttpProxyExternalServer sets HttpProxyExternalServer field to given value.

### HasHttpProxyExternalServer

`func (o *AddAlertHandlerRequest) HasHttpProxyExternalServer() bool`

HasHttpProxyExternalServer returns a boolean if a field has been set.

### GetTwilioAccountSID

`func (o *AddAlertHandlerRequest) GetTwilioAccountSID() string`

GetTwilioAccountSID returns the TwilioAccountSID field if non-nil, zero value otherwise.

### GetTwilioAccountSIDOk

`func (o *AddAlertHandlerRequest) GetTwilioAccountSIDOk() (*string, bool)`

GetTwilioAccountSIDOk returns a tuple with the TwilioAccountSID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAccountSID

`func (o *AddAlertHandlerRequest) SetTwilioAccountSID(v string)`

SetTwilioAccountSID sets TwilioAccountSID field to given value.


### GetTwilioAuthToken

`func (o *AddAlertHandlerRequest) GetTwilioAuthToken() string`

GetTwilioAuthToken returns the TwilioAuthToken field if non-nil, zero value otherwise.

### GetTwilioAuthTokenOk

`func (o *AddAlertHandlerRequest) GetTwilioAuthTokenOk() (*string, bool)`

GetTwilioAuthTokenOk returns a tuple with the TwilioAuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAuthToken

`func (o *AddAlertHandlerRequest) SetTwilioAuthToken(v string)`

SetTwilioAuthToken sets TwilioAuthToken field to given value.

### HasTwilioAuthToken

`func (o *AddAlertHandlerRequest) HasTwilioAuthToken() bool`

HasTwilioAuthToken returns a boolean if a field has been set.

### GetTwilioAuthTokenPassphraseProvider

`func (o *AddAlertHandlerRequest) GetTwilioAuthTokenPassphraseProvider() string`

GetTwilioAuthTokenPassphraseProvider returns the TwilioAuthTokenPassphraseProvider field if non-nil, zero value otherwise.

### GetTwilioAuthTokenPassphraseProviderOk

`func (o *AddAlertHandlerRequest) GetTwilioAuthTokenPassphraseProviderOk() (*string, bool)`

GetTwilioAuthTokenPassphraseProviderOk returns a tuple with the TwilioAuthTokenPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAuthTokenPassphraseProvider

`func (o *AddAlertHandlerRequest) SetTwilioAuthTokenPassphraseProvider(v string)`

SetTwilioAuthTokenPassphraseProvider sets TwilioAuthTokenPassphraseProvider field to given value.

### HasTwilioAuthTokenPassphraseProvider

`func (o *AddAlertHandlerRequest) HasTwilioAuthTokenPassphraseProvider() bool`

HasTwilioAuthTokenPassphraseProvider returns a boolean if a field has been set.

### GetSenderPhoneNumber

`func (o *AddAlertHandlerRequest) GetSenderPhoneNumber() []string`

GetSenderPhoneNumber returns the SenderPhoneNumber field if non-nil, zero value otherwise.

### GetSenderPhoneNumberOk

`func (o *AddAlertHandlerRequest) GetSenderPhoneNumberOk() (*[]string, bool)`

GetSenderPhoneNumberOk returns a tuple with the SenderPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderPhoneNumber

`func (o *AddAlertHandlerRequest) SetSenderPhoneNumber(v []string)`

SetSenderPhoneNumber sets SenderPhoneNumber field to given value.


### GetRecipientPhoneNumber

`func (o *AddAlertHandlerRequest) GetRecipientPhoneNumber() []string`

GetRecipientPhoneNumber returns the RecipientPhoneNumber field if non-nil, zero value otherwise.

### GetRecipientPhoneNumberOk

`func (o *AddAlertHandlerRequest) GetRecipientPhoneNumberOk() (*[]string, bool)`

GetRecipientPhoneNumberOk returns a tuple with the RecipientPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientPhoneNumber

`func (o *AddAlertHandlerRequest) SetRecipientPhoneNumber(v []string)`

SetRecipientPhoneNumber sets RecipientPhoneNumber field to given value.


### GetLongMessageBehavior

`func (o *AddAlertHandlerRequest) GetLongMessageBehavior() EnumalertHandlerLongMessageBehaviorProp`

GetLongMessageBehavior returns the LongMessageBehavior field if non-nil, zero value otherwise.

### GetLongMessageBehaviorOk

`func (o *AddAlertHandlerRequest) GetLongMessageBehaviorOk() (*EnumalertHandlerLongMessageBehaviorProp, bool)`

GetLongMessageBehaviorOk returns a tuple with the LongMessageBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongMessageBehavior

`func (o *AddAlertHandlerRequest) SetLongMessageBehavior(v EnumalertHandlerLongMessageBehaviorProp)`

SetLongMessageBehavior sets LongMessageBehavior field to given value.

### HasLongMessageBehavior

`func (o *AddAlertHandlerRequest) HasLongMessageBehavior() bool`

HasLongMessageBehavior returns a boolean if a field has been set.

### GetCommand

`func (o *AddAlertHandlerRequest) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *AddAlertHandlerRequest) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *AddAlertHandlerRequest) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetExtensionClass

`func (o *AddAlertHandlerRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddAlertHandlerRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddAlertHandlerRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddAlertHandlerRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddAlertHandlerRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddAlertHandlerRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddAlertHandlerRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


