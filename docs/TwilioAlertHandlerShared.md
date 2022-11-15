# TwilioAlertHandlerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumtwilioAlertHandlerSchemaUrn**](EnumtwilioAlertHandlerSchemaUrn.md) |  | 
**Asynchronous** | Pointer to **bool** | Indicates whether the server should attempt to invoke this Twilio Alert Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver the alert notification) will not delay whatever processing the server was performing when the alert was generated. | [optional] 
**TwilioAccountSID** | **string** | The unique identifier assigned to the Twilio account that will be used. | 
**TwilioAuthToken** | Pointer to **string** | The auth token for the Twilio account that will be used. | [optional] 
**TwilioAuthTokenPassphraseProvider** | Pointer to **string** | The passphrase provider that may be used to obtain the auth token for the Twilio account that will be used. | [optional] 
**SenderPhoneNumber** | **[]string** |  | 
**RecipientPhoneNumber** | **[]string** |  | 
**LongMessageBehavior** | [**EnumalertHandlerLongMessageBehaviorProp**](EnumalertHandlerLongMessageBehaviorProp.md) |  | 
**Description** | Pointer to **string** | A description for this Alert Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Alert Handler is enabled. | 
**EnabledAlertSeverity** | Pointer to [**[]EnumalertHandlerEnabledAlertSeverityProp**](EnumalertHandlerEnabledAlertSeverityProp.md) |  | [optional] 
**EnabledAlertType** | Pointer to [**[]EnumalertHandlerEnabledAlertTypeProp**](EnumalertHandlerEnabledAlertTypeProp.md) |  | [optional] 
**DisabledAlertType** | Pointer to [**[]EnumalertHandlerDisabledAlertTypeProp**](EnumalertHandlerDisabledAlertTypeProp.md) |  | [optional] 

## Methods

### NewTwilioAlertHandlerShared

`func NewTwilioAlertHandlerShared(schemas []EnumtwilioAlertHandlerSchemaUrn, twilioAccountSID string, senderPhoneNumber []string, recipientPhoneNumber []string, longMessageBehavior EnumalertHandlerLongMessageBehaviorProp, enabled bool, ) *TwilioAlertHandlerShared`

NewTwilioAlertHandlerShared instantiates a new TwilioAlertHandlerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTwilioAlertHandlerSharedWithDefaults

`func NewTwilioAlertHandlerSharedWithDefaults() *TwilioAlertHandlerShared`

NewTwilioAlertHandlerSharedWithDefaults instantiates a new TwilioAlertHandlerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *TwilioAlertHandlerShared) GetSchemas() []EnumtwilioAlertHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *TwilioAlertHandlerShared) GetSchemasOk() (*[]EnumtwilioAlertHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *TwilioAlertHandlerShared) SetSchemas(v []EnumtwilioAlertHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAsynchronous

`func (o *TwilioAlertHandlerShared) GetAsynchronous() bool`

GetAsynchronous returns the Asynchronous field if non-nil, zero value otherwise.

### GetAsynchronousOk

`func (o *TwilioAlertHandlerShared) GetAsynchronousOk() (*bool, bool)`

GetAsynchronousOk returns a tuple with the Asynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsynchronous

`func (o *TwilioAlertHandlerShared) SetAsynchronous(v bool)`

SetAsynchronous sets Asynchronous field to given value.

### HasAsynchronous

`func (o *TwilioAlertHandlerShared) HasAsynchronous() bool`

HasAsynchronous returns a boolean if a field has been set.

### GetTwilioAccountSID

`func (o *TwilioAlertHandlerShared) GetTwilioAccountSID() string`

GetTwilioAccountSID returns the TwilioAccountSID field if non-nil, zero value otherwise.

### GetTwilioAccountSIDOk

`func (o *TwilioAlertHandlerShared) GetTwilioAccountSIDOk() (*string, bool)`

GetTwilioAccountSIDOk returns a tuple with the TwilioAccountSID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAccountSID

`func (o *TwilioAlertHandlerShared) SetTwilioAccountSID(v string)`

SetTwilioAccountSID sets TwilioAccountSID field to given value.


### GetTwilioAuthToken

`func (o *TwilioAlertHandlerShared) GetTwilioAuthToken() string`

GetTwilioAuthToken returns the TwilioAuthToken field if non-nil, zero value otherwise.

### GetTwilioAuthTokenOk

`func (o *TwilioAlertHandlerShared) GetTwilioAuthTokenOk() (*string, bool)`

GetTwilioAuthTokenOk returns a tuple with the TwilioAuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAuthToken

`func (o *TwilioAlertHandlerShared) SetTwilioAuthToken(v string)`

SetTwilioAuthToken sets TwilioAuthToken field to given value.

### HasTwilioAuthToken

`func (o *TwilioAlertHandlerShared) HasTwilioAuthToken() bool`

HasTwilioAuthToken returns a boolean if a field has been set.

### GetTwilioAuthTokenPassphraseProvider

`func (o *TwilioAlertHandlerShared) GetTwilioAuthTokenPassphraseProvider() string`

GetTwilioAuthTokenPassphraseProvider returns the TwilioAuthTokenPassphraseProvider field if non-nil, zero value otherwise.

### GetTwilioAuthTokenPassphraseProviderOk

`func (o *TwilioAlertHandlerShared) GetTwilioAuthTokenPassphraseProviderOk() (*string, bool)`

GetTwilioAuthTokenPassphraseProviderOk returns a tuple with the TwilioAuthTokenPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAuthTokenPassphraseProvider

`func (o *TwilioAlertHandlerShared) SetTwilioAuthTokenPassphraseProvider(v string)`

SetTwilioAuthTokenPassphraseProvider sets TwilioAuthTokenPassphraseProvider field to given value.

### HasTwilioAuthTokenPassphraseProvider

`func (o *TwilioAlertHandlerShared) HasTwilioAuthTokenPassphraseProvider() bool`

HasTwilioAuthTokenPassphraseProvider returns a boolean if a field has been set.

### GetSenderPhoneNumber

`func (o *TwilioAlertHandlerShared) GetSenderPhoneNumber() []string`

GetSenderPhoneNumber returns the SenderPhoneNumber field if non-nil, zero value otherwise.

### GetSenderPhoneNumberOk

`func (o *TwilioAlertHandlerShared) GetSenderPhoneNumberOk() (*[]string, bool)`

GetSenderPhoneNumberOk returns a tuple with the SenderPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderPhoneNumber

`func (o *TwilioAlertHandlerShared) SetSenderPhoneNumber(v []string)`

SetSenderPhoneNumber sets SenderPhoneNumber field to given value.


### GetRecipientPhoneNumber

`func (o *TwilioAlertHandlerShared) GetRecipientPhoneNumber() []string`

GetRecipientPhoneNumber returns the RecipientPhoneNumber field if non-nil, zero value otherwise.

### GetRecipientPhoneNumberOk

`func (o *TwilioAlertHandlerShared) GetRecipientPhoneNumberOk() (*[]string, bool)`

GetRecipientPhoneNumberOk returns a tuple with the RecipientPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientPhoneNumber

`func (o *TwilioAlertHandlerShared) SetRecipientPhoneNumber(v []string)`

SetRecipientPhoneNumber sets RecipientPhoneNumber field to given value.


### GetLongMessageBehavior

`func (o *TwilioAlertHandlerShared) GetLongMessageBehavior() EnumalertHandlerLongMessageBehaviorProp`

GetLongMessageBehavior returns the LongMessageBehavior field if non-nil, zero value otherwise.

### GetLongMessageBehaviorOk

`func (o *TwilioAlertHandlerShared) GetLongMessageBehaviorOk() (*EnumalertHandlerLongMessageBehaviorProp, bool)`

GetLongMessageBehaviorOk returns a tuple with the LongMessageBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongMessageBehavior

`func (o *TwilioAlertHandlerShared) SetLongMessageBehavior(v EnumalertHandlerLongMessageBehaviorProp)`

SetLongMessageBehavior sets LongMessageBehavior field to given value.


### GetDescription

`func (o *TwilioAlertHandlerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TwilioAlertHandlerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TwilioAlertHandlerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TwilioAlertHandlerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *TwilioAlertHandlerShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TwilioAlertHandlerShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TwilioAlertHandlerShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEnabledAlertSeverity

`func (o *TwilioAlertHandlerShared) GetEnabledAlertSeverity() []EnumalertHandlerEnabledAlertSeverityProp`

GetEnabledAlertSeverity returns the EnabledAlertSeverity field if non-nil, zero value otherwise.

### GetEnabledAlertSeverityOk

`func (o *TwilioAlertHandlerShared) GetEnabledAlertSeverityOk() (*[]EnumalertHandlerEnabledAlertSeverityProp, bool)`

GetEnabledAlertSeverityOk returns a tuple with the EnabledAlertSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAlertSeverity

`func (o *TwilioAlertHandlerShared) SetEnabledAlertSeverity(v []EnumalertHandlerEnabledAlertSeverityProp)`

SetEnabledAlertSeverity sets EnabledAlertSeverity field to given value.

### HasEnabledAlertSeverity

`func (o *TwilioAlertHandlerShared) HasEnabledAlertSeverity() bool`

HasEnabledAlertSeverity returns a boolean if a field has been set.

### GetEnabledAlertType

`func (o *TwilioAlertHandlerShared) GetEnabledAlertType() []EnumalertHandlerEnabledAlertTypeProp`

GetEnabledAlertType returns the EnabledAlertType field if non-nil, zero value otherwise.

### GetEnabledAlertTypeOk

`func (o *TwilioAlertHandlerShared) GetEnabledAlertTypeOk() (*[]EnumalertHandlerEnabledAlertTypeProp, bool)`

GetEnabledAlertTypeOk returns a tuple with the EnabledAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAlertType

`func (o *TwilioAlertHandlerShared) SetEnabledAlertType(v []EnumalertHandlerEnabledAlertTypeProp)`

SetEnabledAlertType sets EnabledAlertType field to given value.

### HasEnabledAlertType

`func (o *TwilioAlertHandlerShared) HasEnabledAlertType() bool`

HasEnabledAlertType returns a boolean if a field has been set.

### GetDisabledAlertType

`func (o *TwilioAlertHandlerShared) GetDisabledAlertType() []EnumalertHandlerDisabledAlertTypeProp`

GetDisabledAlertType returns the DisabledAlertType field if non-nil, zero value otherwise.

### GetDisabledAlertTypeOk

`func (o *TwilioAlertHandlerShared) GetDisabledAlertTypeOk() (*[]EnumalertHandlerDisabledAlertTypeProp, bool)`

GetDisabledAlertTypeOk returns a tuple with the DisabledAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAlertType

`func (o *TwilioAlertHandlerShared) SetDisabledAlertType(v []EnumalertHandlerDisabledAlertTypeProp)`

SetDisabledAlertType sets DisabledAlertType field to given value.

### HasDisabledAlertType

`func (o *TwilioAlertHandlerShared) HasDisabledAlertType() bool`

HasDisabledAlertType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


