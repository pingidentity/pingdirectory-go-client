/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the AddTwilioAlertHandlerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddTwilioAlertHandlerRequest{}

// AddTwilioAlertHandlerRequest struct for AddTwilioAlertHandlerRequest
type AddTwilioAlertHandlerRequest struct {
	Schemas []EnumtwilioAlertHandlerSchemaUrn `json:"schemas"`
	// Indicates whether the server should attempt to invoke this Twilio Alert Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver the alert notification) will not delay whatever processing the server was performing when the alert was generated.
	Asynchronous *bool `json:"asynchronous,omitempty"`
	// A reference to an HTTP proxy server that should be used for requests sent to the Twilio service.
	HttpProxyExternalServer *string `json:"httpProxyExternalServer,omitempty"`
	// The unique identifier assigned to the Twilio account that will be used.
	TwilioAccountSID string `json:"twilioAccountSID"`
	// The auth token for the Twilio account that will be used.
	TwilioAuthToken *string `json:"twilioAuthToken,omitempty"`
	// The passphrase provider that may be used to obtain the auth token for the Twilio account that will be used.
	TwilioAuthTokenPassphraseProvider *string `json:"twilioAuthTokenPassphraseProvider,omitempty"`
	// The outgoing phone number to use for the messages. Values must be phone numbers you have obtained for use with your Twilio account.
	SenderPhoneNumber []string `json:"senderPhoneNumber"`
	// The phone number to which alert notifications should be delivered.
	RecipientPhoneNumber []string                                 `json:"recipientPhoneNumber"`
	LongMessageBehavior  *EnumalertHandlerLongMessageBehaviorProp `json:"longMessageBehavior,omitempty"`
	// A description for this Alert Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the Alert Handler is enabled.
	Enabled              bool                                       `json:"enabled"`
	EnabledAlertSeverity []EnumalertHandlerEnabledAlertSeverityProp `json:"enabledAlertSeverity,omitempty"`
	EnabledAlertType     []EnumalertHandlerEnabledAlertTypeProp     `json:"enabledAlertType,omitempty"`
	DisabledAlertType    []EnumalertHandlerDisabledAlertTypeProp    `json:"disabledAlertType,omitempty"`
	// Name of the new Alert Handler
	HandlerName string `json:"handlerName"`
}

// NewAddTwilioAlertHandlerRequest instantiates a new AddTwilioAlertHandlerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddTwilioAlertHandlerRequest(schemas []EnumtwilioAlertHandlerSchemaUrn, twilioAccountSID string, senderPhoneNumber []string, recipientPhoneNumber []string, enabled bool, handlerName string) *AddTwilioAlertHandlerRequest {
	this := AddTwilioAlertHandlerRequest{}
	this.Schemas = schemas
	this.TwilioAccountSID = twilioAccountSID
	this.SenderPhoneNumber = senderPhoneNumber
	this.RecipientPhoneNumber = recipientPhoneNumber
	this.Enabled = enabled
	this.HandlerName = handlerName
	return &this
}

// NewAddTwilioAlertHandlerRequestWithDefaults instantiates a new AddTwilioAlertHandlerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddTwilioAlertHandlerRequestWithDefaults() *AddTwilioAlertHandlerRequest {
	this := AddTwilioAlertHandlerRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddTwilioAlertHandlerRequest) GetSchemas() []EnumtwilioAlertHandlerSchemaUrn {
	if o == nil {
		var ret []EnumtwilioAlertHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddTwilioAlertHandlerRequest) GetSchemasOk() ([]EnumtwilioAlertHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddTwilioAlertHandlerRequest) SetSchemas(v []EnumtwilioAlertHandlerSchemaUrn) {
	o.Schemas = v
}

// GetAsynchronous returns the Asynchronous field value if set, zero value otherwise.
func (o *AddTwilioAlertHandlerRequest) GetAsynchronous() bool {
	if o == nil || IsNil(o.Asynchronous) {
		var ret bool
		return ret
	}
	return *o.Asynchronous
}

// GetAsynchronousOk returns a tuple with the Asynchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTwilioAlertHandlerRequest) GetAsynchronousOk() (*bool, bool) {
	if o == nil || IsNil(o.Asynchronous) {
		return nil, false
	}
	return o.Asynchronous, true
}

// HasAsynchronous returns a boolean if a field has been set.
func (o *AddTwilioAlertHandlerRequest) HasAsynchronous() bool {
	if o != nil && !IsNil(o.Asynchronous) {
		return true
	}

	return false
}

// SetAsynchronous gets a reference to the given bool and assigns it to the Asynchronous field.
func (o *AddTwilioAlertHandlerRequest) SetAsynchronous(v bool) {
	o.Asynchronous = &v
}

// GetHttpProxyExternalServer returns the HttpProxyExternalServer field value if set, zero value otherwise.
func (o *AddTwilioAlertHandlerRequest) GetHttpProxyExternalServer() string {
	if o == nil || IsNil(o.HttpProxyExternalServer) {
		var ret string
		return ret
	}
	return *o.HttpProxyExternalServer
}

// GetHttpProxyExternalServerOk returns a tuple with the HttpProxyExternalServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTwilioAlertHandlerRequest) GetHttpProxyExternalServerOk() (*string, bool) {
	if o == nil || IsNil(o.HttpProxyExternalServer) {
		return nil, false
	}
	return o.HttpProxyExternalServer, true
}

// HasHttpProxyExternalServer returns a boolean if a field has been set.
func (o *AddTwilioAlertHandlerRequest) HasHttpProxyExternalServer() bool {
	if o != nil && !IsNil(o.HttpProxyExternalServer) {
		return true
	}

	return false
}

// SetHttpProxyExternalServer gets a reference to the given string and assigns it to the HttpProxyExternalServer field.
func (o *AddTwilioAlertHandlerRequest) SetHttpProxyExternalServer(v string) {
	o.HttpProxyExternalServer = &v
}

// GetTwilioAccountSID returns the TwilioAccountSID field value
func (o *AddTwilioAlertHandlerRequest) GetTwilioAccountSID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TwilioAccountSID
}

// GetTwilioAccountSIDOk returns a tuple with the TwilioAccountSID field value
// and a boolean to check if the value has been set.
func (o *AddTwilioAlertHandlerRequest) GetTwilioAccountSIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TwilioAccountSID, true
}

// SetTwilioAccountSID sets field value
func (o *AddTwilioAlertHandlerRequest) SetTwilioAccountSID(v string) {
	o.TwilioAccountSID = v
}

// GetTwilioAuthToken returns the TwilioAuthToken field value if set, zero value otherwise.
func (o *AddTwilioAlertHandlerRequest) GetTwilioAuthToken() string {
	if o == nil || IsNil(o.TwilioAuthToken) {
		var ret string
		return ret
	}
	return *o.TwilioAuthToken
}

// GetTwilioAuthTokenOk returns a tuple with the TwilioAuthToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTwilioAlertHandlerRequest) GetTwilioAuthTokenOk() (*string, bool) {
	if o == nil || IsNil(o.TwilioAuthToken) {
		return nil, false
	}
	return o.TwilioAuthToken, true
}

// HasTwilioAuthToken returns a boolean if a field has been set.
func (o *AddTwilioAlertHandlerRequest) HasTwilioAuthToken() bool {
	if o != nil && !IsNil(o.TwilioAuthToken) {
		return true
	}

	return false
}

// SetTwilioAuthToken gets a reference to the given string and assigns it to the TwilioAuthToken field.
func (o *AddTwilioAlertHandlerRequest) SetTwilioAuthToken(v string) {
	o.TwilioAuthToken = &v
}

// GetTwilioAuthTokenPassphraseProvider returns the TwilioAuthTokenPassphraseProvider field value if set, zero value otherwise.
func (o *AddTwilioAlertHandlerRequest) GetTwilioAuthTokenPassphraseProvider() string {
	if o == nil || IsNil(o.TwilioAuthTokenPassphraseProvider) {
		var ret string
		return ret
	}
	return *o.TwilioAuthTokenPassphraseProvider
}

// GetTwilioAuthTokenPassphraseProviderOk returns a tuple with the TwilioAuthTokenPassphraseProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTwilioAlertHandlerRequest) GetTwilioAuthTokenPassphraseProviderOk() (*string, bool) {
	if o == nil || IsNil(o.TwilioAuthTokenPassphraseProvider) {
		return nil, false
	}
	return o.TwilioAuthTokenPassphraseProvider, true
}

// HasTwilioAuthTokenPassphraseProvider returns a boolean if a field has been set.
func (o *AddTwilioAlertHandlerRequest) HasTwilioAuthTokenPassphraseProvider() bool {
	if o != nil && !IsNil(o.TwilioAuthTokenPassphraseProvider) {
		return true
	}

	return false
}

// SetTwilioAuthTokenPassphraseProvider gets a reference to the given string and assigns it to the TwilioAuthTokenPassphraseProvider field.
func (o *AddTwilioAlertHandlerRequest) SetTwilioAuthTokenPassphraseProvider(v string) {
	o.TwilioAuthTokenPassphraseProvider = &v
}

// GetSenderPhoneNumber returns the SenderPhoneNumber field value
func (o *AddTwilioAlertHandlerRequest) GetSenderPhoneNumber() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SenderPhoneNumber
}

// GetSenderPhoneNumberOk returns a tuple with the SenderPhoneNumber field value
// and a boolean to check if the value has been set.
func (o *AddTwilioAlertHandlerRequest) GetSenderPhoneNumberOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SenderPhoneNumber, true
}

// SetSenderPhoneNumber sets field value
func (o *AddTwilioAlertHandlerRequest) SetSenderPhoneNumber(v []string) {
	o.SenderPhoneNumber = v
}

// GetRecipientPhoneNumber returns the RecipientPhoneNumber field value
func (o *AddTwilioAlertHandlerRequest) GetRecipientPhoneNumber() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RecipientPhoneNumber
}

// GetRecipientPhoneNumberOk returns a tuple with the RecipientPhoneNumber field value
// and a boolean to check if the value has been set.
func (o *AddTwilioAlertHandlerRequest) GetRecipientPhoneNumberOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecipientPhoneNumber, true
}

// SetRecipientPhoneNumber sets field value
func (o *AddTwilioAlertHandlerRequest) SetRecipientPhoneNumber(v []string) {
	o.RecipientPhoneNumber = v
}

// GetLongMessageBehavior returns the LongMessageBehavior field value if set, zero value otherwise.
func (o *AddTwilioAlertHandlerRequest) GetLongMessageBehavior() EnumalertHandlerLongMessageBehaviorProp {
	if o == nil || IsNil(o.LongMessageBehavior) {
		var ret EnumalertHandlerLongMessageBehaviorProp
		return ret
	}
	return *o.LongMessageBehavior
}

// GetLongMessageBehaviorOk returns a tuple with the LongMessageBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTwilioAlertHandlerRequest) GetLongMessageBehaviorOk() (*EnumalertHandlerLongMessageBehaviorProp, bool) {
	if o == nil || IsNil(o.LongMessageBehavior) {
		return nil, false
	}
	return o.LongMessageBehavior, true
}

// HasLongMessageBehavior returns a boolean if a field has been set.
func (o *AddTwilioAlertHandlerRequest) HasLongMessageBehavior() bool {
	if o != nil && !IsNil(o.LongMessageBehavior) {
		return true
	}

	return false
}

// SetLongMessageBehavior gets a reference to the given EnumalertHandlerLongMessageBehaviorProp and assigns it to the LongMessageBehavior field.
func (o *AddTwilioAlertHandlerRequest) SetLongMessageBehavior(v EnumalertHandlerLongMessageBehaviorProp) {
	o.LongMessageBehavior = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddTwilioAlertHandlerRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTwilioAlertHandlerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddTwilioAlertHandlerRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddTwilioAlertHandlerRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddTwilioAlertHandlerRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddTwilioAlertHandlerRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddTwilioAlertHandlerRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEnabledAlertSeverity returns the EnabledAlertSeverity field value if set, zero value otherwise.
func (o *AddTwilioAlertHandlerRequest) GetEnabledAlertSeverity() []EnumalertHandlerEnabledAlertSeverityProp {
	if o == nil || IsNil(o.EnabledAlertSeverity) {
		var ret []EnumalertHandlerEnabledAlertSeverityProp
		return ret
	}
	return o.EnabledAlertSeverity
}

// GetEnabledAlertSeverityOk returns a tuple with the EnabledAlertSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTwilioAlertHandlerRequest) GetEnabledAlertSeverityOk() ([]EnumalertHandlerEnabledAlertSeverityProp, bool) {
	if o == nil || IsNil(o.EnabledAlertSeverity) {
		return nil, false
	}
	return o.EnabledAlertSeverity, true
}

// HasEnabledAlertSeverity returns a boolean if a field has been set.
func (o *AddTwilioAlertHandlerRequest) HasEnabledAlertSeverity() bool {
	if o != nil && !IsNil(o.EnabledAlertSeverity) {
		return true
	}

	return false
}

// SetEnabledAlertSeverity gets a reference to the given []EnumalertHandlerEnabledAlertSeverityProp and assigns it to the EnabledAlertSeverity field.
func (o *AddTwilioAlertHandlerRequest) SetEnabledAlertSeverity(v []EnumalertHandlerEnabledAlertSeverityProp) {
	o.EnabledAlertSeverity = v
}

// GetEnabledAlertType returns the EnabledAlertType field value if set, zero value otherwise.
func (o *AddTwilioAlertHandlerRequest) GetEnabledAlertType() []EnumalertHandlerEnabledAlertTypeProp {
	if o == nil || IsNil(o.EnabledAlertType) {
		var ret []EnumalertHandlerEnabledAlertTypeProp
		return ret
	}
	return o.EnabledAlertType
}

// GetEnabledAlertTypeOk returns a tuple with the EnabledAlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTwilioAlertHandlerRequest) GetEnabledAlertTypeOk() ([]EnumalertHandlerEnabledAlertTypeProp, bool) {
	if o == nil || IsNil(o.EnabledAlertType) {
		return nil, false
	}
	return o.EnabledAlertType, true
}

// HasEnabledAlertType returns a boolean if a field has been set.
func (o *AddTwilioAlertHandlerRequest) HasEnabledAlertType() bool {
	if o != nil && !IsNil(o.EnabledAlertType) {
		return true
	}

	return false
}

// SetEnabledAlertType gets a reference to the given []EnumalertHandlerEnabledAlertTypeProp and assigns it to the EnabledAlertType field.
func (o *AddTwilioAlertHandlerRequest) SetEnabledAlertType(v []EnumalertHandlerEnabledAlertTypeProp) {
	o.EnabledAlertType = v
}

// GetDisabledAlertType returns the DisabledAlertType field value if set, zero value otherwise.
func (o *AddTwilioAlertHandlerRequest) GetDisabledAlertType() []EnumalertHandlerDisabledAlertTypeProp {
	if o == nil || IsNil(o.DisabledAlertType) {
		var ret []EnumalertHandlerDisabledAlertTypeProp
		return ret
	}
	return o.DisabledAlertType
}

// GetDisabledAlertTypeOk returns a tuple with the DisabledAlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTwilioAlertHandlerRequest) GetDisabledAlertTypeOk() ([]EnumalertHandlerDisabledAlertTypeProp, bool) {
	if o == nil || IsNil(o.DisabledAlertType) {
		return nil, false
	}
	return o.DisabledAlertType, true
}

// HasDisabledAlertType returns a boolean if a field has been set.
func (o *AddTwilioAlertHandlerRequest) HasDisabledAlertType() bool {
	if o != nil && !IsNil(o.DisabledAlertType) {
		return true
	}

	return false
}

// SetDisabledAlertType gets a reference to the given []EnumalertHandlerDisabledAlertTypeProp and assigns it to the DisabledAlertType field.
func (o *AddTwilioAlertHandlerRequest) SetDisabledAlertType(v []EnumalertHandlerDisabledAlertTypeProp) {
	o.DisabledAlertType = v
}

// GetHandlerName returns the HandlerName field value
func (o *AddTwilioAlertHandlerRequest) GetHandlerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HandlerName
}

// GetHandlerNameOk returns a tuple with the HandlerName field value
// and a boolean to check if the value has been set.
func (o *AddTwilioAlertHandlerRequest) GetHandlerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HandlerName, true
}

// SetHandlerName sets field value
func (o *AddTwilioAlertHandlerRequest) SetHandlerName(v string) {
	o.HandlerName = v
}

func (o AddTwilioAlertHandlerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddTwilioAlertHandlerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.Asynchronous) {
		toSerialize["asynchronous"] = o.Asynchronous
	}
	if !IsNil(o.HttpProxyExternalServer) {
		toSerialize["httpProxyExternalServer"] = o.HttpProxyExternalServer
	}
	toSerialize["twilioAccountSID"] = o.TwilioAccountSID
	if !IsNil(o.TwilioAuthToken) {
		toSerialize["twilioAuthToken"] = o.TwilioAuthToken
	}
	if !IsNil(o.TwilioAuthTokenPassphraseProvider) {
		toSerialize["twilioAuthTokenPassphraseProvider"] = o.TwilioAuthTokenPassphraseProvider
	}
	toSerialize["senderPhoneNumber"] = o.SenderPhoneNumber
	toSerialize["recipientPhoneNumber"] = o.RecipientPhoneNumber
	if !IsNil(o.LongMessageBehavior) {
		toSerialize["longMessageBehavior"] = o.LongMessageBehavior
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.EnabledAlertSeverity) {
		toSerialize["enabledAlertSeverity"] = o.EnabledAlertSeverity
	}
	if !IsNil(o.EnabledAlertType) {
		toSerialize["enabledAlertType"] = o.EnabledAlertType
	}
	if !IsNil(o.DisabledAlertType) {
		toSerialize["disabledAlertType"] = o.DisabledAlertType
	}
	toSerialize["handlerName"] = o.HandlerName
	return toSerialize, nil
}

type NullableAddTwilioAlertHandlerRequest struct {
	value *AddTwilioAlertHandlerRequest
	isSet bool
}

func (v NullableAddTwilioAlertHandlerRequest) Get() *AddTwilioAlertHandlerRequest {
	return v.value
}

func (v *NullableAddTwilioAlertHandlerRequest) Set(val *AddTwilioAlertHandlerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddTwilioAlertHandlerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddTwilioAlertHandlerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddTwilioAlertHandlerRequest(val *AddTwilioAlertHandlerRequest) *NullableAddTwilioAlertHandlerRequest {
	return &NullableAddTwilioAlertHandlerRequest{value: val, isSet: true}
}

func (v NullableAddTwilioAlertHandlerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddTwilioAlertHandlerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
