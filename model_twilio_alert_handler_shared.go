/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// TwilioAlertHandlerShared struct for TwilioAlertHandlerShared
type TwilioAlertHandlerShared struct {
	Schemas []EnumtwilioAlertHandlerSchemaUrn `json:"schemas"`
	// Indicates whether the server should attempt to invoke this Twilio Alert Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver the alert notification) will not delay whatever processing the server was performing when the alert was generated.
	Asynchronous *bool `json:"asynchronous,omitempty"`
	// The unique identifier assigned to the Twilio account that will be used.
	TwilioAccountSID string `json:"twilioAccountSID"`
	// The auth token for the Twilio account that will be used.
	TwilioAuthToken *string `json:"twilioAuthToken,omitempty"`
	// The passphrase provider that may be used to obtain the auth token for the Twilio account that will be used.
	TwilioAuthTokenPassphraseProvider *string `json:"twilioAuthTokenPassphraseProvider,omitempty"`
	SenderPhoneNumber []string `json:"senderPhoneNumber"`
	RecipientPhoneNumber []string `json:"recipientPhoneNumber"`
	LongMessageBehavior EnumalertHandlerLongMessageBehaviorProp `json:"longMessageBehavior"`
	// A description for this Alert Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the Alert Handler is enabled.
	Enabled bool `json:"enabled"`
	EnabledAlertSeverity []EnumalertHandlerEnabledAlertSeverityProp `json:"enabledAlertSeverity,omitempty"`
	EnabledAlertType []EnumalertHandlerEnabledAlertTypeProp `json:"enabledAlertType,omitempty"`
	DisabledAlertType []EnumalertHandlerDisabledAlertTypeProp `json:"disabledAlertType,omitempty"`
}

// NewTwilioAlertHandlerShared instantiates a new TwilioAlertHandlerShared object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTwilioAlertHandlerShared(schemas []EnumtwilioAlertHandlerSchemaUrn, twilioAccountSID string, senderPhoneNumber []string, recipientPhoneNumber []string, longMessageBehavior EnumalertHandlerLongMessageBehaviorProp, enabled bool) *TwilioAlertHandlerShared {
	this := TwilioAlertHandlerShared{}
	this.Schemas = schemas
	this.TwilioAccountSID = twilioAccountSID
	this.SenderPhoneNumber = senderPhoneNumber
	this.RecipientPhoneNumber = recipientPhoneNumber
	this.LongMessageBehavior = longMessageBehavior
	this.Enabled = enabled
	return &this
}

// NewTwilioAlertHandlerSharedWithDefaults instantiates a new TwilioAlertHandlerShared object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTwilioAlertHandlerSharedWithDefaults() *TwilioAlertHandlerShared {
	this := TwilioAlertHandlerShared{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *TwilioAlertHandlerShared) GetSchemas() []EnumtwilioAlertHandlerSchemaUrn {
	if o == nil {
		var ret []EnumtwilioAlertHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *TwilioAlertHandlerShared) GetSchemasOk() ([]EnumtwilioAlertHandlerSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *TwilioAlertHandlerShared) SetSchemas(v []EnumtwilioAlertHandlerSchemaUrn) {
	o.Schemas = v
}

// GetAsynchronous returns the Asynchronous field value if set, zero value otherwise.
func (o *TwilioAlertHandlerShared) GetAsynchronous() bool {
	if o == nil || isNil(o.Asynchronous) {
		var ret bool
		return ret
	}
	return *o.Asynchronous
}

// GetAsynchronousOk returns a tuple with the Asynchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioAlertHandlerShared) GetAsynchronousOk() (*bool, bool) {
	if o == nil || isNil(o.Asynchronous) {
    return nil, false
	}
	return o.Asynchronous, true
}

// HasAsynchronous returns a boolean if a field has been set.
func (o *TwilioAlertHandlerShared) HasAsynchronous() bool {
	if o != nil && !isNil(o.Asynchronous) {
		return true
	}

	return false
}

// SetAsynchronous gets a reference to the given bool and assigns it to the Asynchronous field.
func (o *TwilioAlertHandlerShared) SetAsynchronous(v bool) {
	o.Asynchronous = &v
}

// GetTwilioAccountSID returns the TwilioAccountSID field value
func (o *TwilioAlertHandlerShared) GetTwilioAccountSID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TwilioAccountSID
}

// GetTwilioAccountSIDOk returns a tuple with the TwilioAccountSID field value
// and a boolean to check if the value has been set.
func (o *TwilioAlertHandlerShared) GetTwilioAccountSIDOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TwilioAccountSID, true
}

// SetTwilioAccountSID sets field value
func (o *TwilioAlertHandlerShared) SetTwilioAccountSID(v string) {
	o.TwilioAccountSID = v
}

// GetTwilioAuthToken returns the TwilioAuthToken field value if set, zero value otherwise.
func (o *TwilioAlertHandlerShared) GetTwilioAuthToken() string {
	if o == nil || isNil(o.TwilioAuthToken) {
		var ret string
		return ret
	}
	return *o.TwilioAuthToken
}

// GetTwilioAuthTokenOk returns a tuple with the TwilioAuthToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioAlertHandlerShared) GetTwilioAuthTokenOk() (*string, bool) {
	if o == nil || isNil(o.TwilioAuthToken) {
    return nil, false
	}
	return o.TwilioAuthToken, true
}

// HasTwilioAuthToken returns a boolean if a field has been set.
func (o *TwilioAlertHandlerShared) HasTwilioAuthToken() bool {
	if o != nil && !isNil(o.TwilioAuthToken) {
		return true
	}

	return false
}

// SetTwilioAuthToken gets a reference to the given string and assigns it to the TwilioAuthToken field.
func (o *TwilioAlertHandlerShared) SetTwilioAuthToken(v string) {
	o.TwilioAuthToken = &v
}

// GetTwilioAuthTokenPassphraseProvider returns the TwilioAuthTokenPassphraseProvider field value if set, zero value otherwise.
func (o *TwilioAlertHandlerShared) GetTwilioAuthTokenPassphraseProvider() string {
	if o == nil || isNil(o.TwilioAuthTokenPassphraseProvider) {
		var ret string
		return ret
	}
	return *o.TwilioAuthTokenPassphraseProvider
}

// GetTwilioAuthTokenPassphraseProviderOk returns a tuple with the TwilioAuthTokenPassphraseProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioAlertHandlerShared) GetTwilioAuthTokenPassphraseProviderOk() (*string, bool) {
	if o == nil || isNil(o.TwilioAuthTokenPassphraseProvider) {
    return nil, false
	}
	return o.TwilioAuthTokenPassphraseProvider, true
}

// HasTwilioAuthTokenPassphraseProvider returns a boolean if a field has been set.
func (o *TwilioAlertHandlerShared) HasTwilioAuthTokenPassphraseProvider() bool {
	if o != nil && !isNil(o.TwilioAuthTokenPassphraseProvider) {
		return true
	}

	return false
}

// SetTwilioAuthTokenPassphraseProvider gets a reference to the given string and assigns it to the TwilioAuthTokenPassphraseProvider field.
func (o *TwilioAlertHandlerShared) SetTwilioAuthTokenPassphraseProvider(v string) {
	o.TwilioAuthTokenPassphraseProvider = &v
}

// GetSenderPhoneNumber returns the SenderPhoneNumber field value
func (o *TwilioAlertHandlerShared) GetSenderPhoneNumber() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SenderPhoneNumber
}

// GetSenderPhoneNumberOk returns a tuple with the SenderPhoneNumber field value
// and a boolean to check if the value has been set.
func (o *TwilioAlertHandlerShared) GetSenderPhoneNumberOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.SenderPhoneNumber, true
}

// SetSenderPhoneNumber sets field value
func (o *TwilioAlertHandlerShared) SetSenderPhoneNumber(v []string) {
	o.SenderPhoneNumber = v
}

// GetRecipientPhoneNumber returns the RecipientPhoneNumber field value
func (o *TwilioAlertHandlerShared) GetRecipientPhoneNumber() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RecipientPhoneNumber
}

// GetRecipientPhoneNumberOk returns a tuple with the RecipientPhoneNumber field value
// and a boolean to check if the value has been set.
func (o *TwilioAlertHandlerShared) GetRecipientPhoneNumberOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RecipientPhoneNumber, true
}

// SetRecipientPhoneNumber sets field value
func (o *TwilioAlertHandlerShared) SetRecipientPhoneNumber(v []string) {
	o.RecipientPhoneNumber = v
}

// GetLongMessageBehavior returns the LongMessageBehavior field value
func (o *TwilioAlertHandlerShared) GetLongMessageBehavior() EnumalertHandlerLongMessageBehaviorProp {
	if o == nil {
		var ret EnumalertHandlerLongMessageBehaviorProp
		return ret
	}

	return o.LongMessageBehavior
}

// GetLongMessageBehaviorOk returns a tuple with the LongMessageBehavior field value
// and a boolean to check if the value has been set.
func (o *TwilioAlertHandlerShared) GetLongMessageBehaviorOk() (*EnumalertHandlerLongMessageBehaviorProp, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LongMessageBehavior, true
}

// SetLongMessageBehavior sets field value
func (o *TwilioAlertHandlerShared) SetLongMessageBehavior(v EnumalertHandlerLongMessageBehaviorProp) {
	o.LongMessageBehavior = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TwilioAlertHandlerShared) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioAlertHandlerShared) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TwilioAlertHandlerShared) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TwilioAlertHandlerShared) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *TwilioAlertHandlerShared) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *TwilioAlertHandlerShared) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *TwilioAlertHandlerShared) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEnabledAlertSeverity returns the EnabledAlertSeverity field value if set, zero value otherwise.
func (o *TwilioAlertHandlerShared) GetEnabledAlertSeverity() []EnumalertHandlerEnabledAlertSeverityProp {
	if o == nil || isNil(o.EnabledAlertSeverity) {
		var ret []EnumalertHandlerEnabledAlertSeverityProp
		return ret
	}
	return o.EnabledAlertSeverity
}

// GetEnabledAlertSeverityOk returns a tuple with the EnabledAlertSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioAlertHandlerShared) GetEnabledAlertSeverityOk() ([]EnumalertHandlerEnabledAlertSeverityProp, bool) {
	if o == nil || isNil(o.EnabledAlertSeverity) {
    return nil, false
	}
	return o.EnabledAlertSeverity, true
}

// HasEnabledAlertSeverity returns a boolean if a field has been set.
func (o *TwilioAlertHandlerShared) HasEnabledAlertSeverity() bool {
	if o != nil && !isNil(o.EnabledAlertSeverity) {
		return true
	}

	return false
}

// SetEnabledAlertSeverity gets a reference to the given []EnumalertHandlerEnabledAlertSeverityProp and assigns it to the EnabledAlertSeverity field.
func (o *TwilioAlertHandlerShared) SetEnabledAlertSeverity(v []EnumalertHandlerEnabledAlertSeverityProp) {
	o.EnabledAlertSeverity = v
}

// GetEnabledAlertType returns the EnabledAlertType field value if set, zero value otherwise.
func (o *TwilioAlertHandlerShared) GetEnabledAlertType() []EnumalertHandlerEnabledAlertTypeProp {
	if o == nil || isNil(o.EnabledAlertType) {
		var ret []EnumalertHandlerEnabledAlertTypeProp
		return ret
	}
	return o.EnabledAlertType
}

// GetEnabledAlertTypeOk returns a tuple with the EnabledAlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioAlertHandlerShared) GetEnabledAlertTypeOk() ([]EnumalertHandlerEnabledAlertTypeProp, bool) {
	if o == nil || isNil(o.EnabledAlertType) {
    return nil, false
	}
	return o.EnabledAlertType, true
}

// HasEnabledAlertType returns a boolean if a field has been set.
func (o *TwilioAlertHandlerShared) HasEnabledAlertType() bool {
	if o != nil && !isNil(o.EnabledAlertType) {
		return true
	}

	return false
}

// SetEnabledAlertType gets a reference to the given []EnumalertHandlerEnabledAlertTypeProp and assigns it to the EnabledAlertType field.
func (o *TwilioAlertHandlerShared) SetEnabledAlertType(v []EnumalertHandlerEnabledAlertTypeProp) {
	o.EnabledAlertType = v
}

// GetDisabledAlertType returns the DisabledAlertType field value if set, zero value otherwise.
func (o *TwilioAlertHandlerShared) GetDisabledAlertType() []EnumalertHandlerDisabledAlertTypeProp {
	if o == nil || isNil(o.DisabledAlertType) {
		var ret []EnumalertHandlerDisabledAlertTypeProp
		return ret
	}
	return o.DisabledAlertType
}

// GetDisabledAlertTypeOk returns a tuple with the DisabledAlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioAlertHandlerShared) GetDisabledAlertTypeOk() ([]EnumalertHandlerDisabledAlertTypeProp, bool) {
	if o == nil || isNil(o.DisabledAlertType) {
    return nil, false
	}
	return o.DisabledAlertType, true
}

// HasDisabledAlertType returns a boolean if a field has been set.
func (o *TwilioAlertHandlerShared) HasDisabledAlertType() bool {
	if o != nil && !isNil(o.DisabledAlertType) {
		return true
	}

	return false
}

// SetDisabledAlertType gets a reference to the given []EnumalertHandlerDisabledAlertTypeProp and assigns it to the DisabledAlertType field.
func (o *TwilioAlertHandlerShared) SetDisabledAlertType(v []EnumalertHandlerDisabledAlertTypeProp) {
	o.DisabledAlertType = v
}

func (o TwilioAlertHandlerShared) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.Asynchronous) {
		toSerialize["asynchronous"] = o.Asynchronous
	}
	if true {
		toSerialize["twilioAccountSID"] = o.TwilioAccountSID
	}
	if !isNil(o.TwilioAuthToken) {
		toSerialize["twilioAuthToken"] = o.TwilioAuthToken
	}
	if !isNil(o.TwilioAuthTokenPassphraseProvider) {
		toSerialize["twilioAuthTokenPassphraseProvider"] = o.TwilioAuthTokenPassphraseProvider
	}
	if true {
		toSerialize["senderPhoneNumber"] = o.SenderPhoneNumber
	}
	if true {
		toSerialize["recipientPhoneNumber"] = o.RecipientPhoneNumber
	}
	if true {
		toSerialize["longMessageBehavior"] = o.LongMessageBehavior
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.EnabledAlertSeverity) {
		toSerialize["enabledAlertSeverity"] = o.EnabledAlertSeverity
	}
	if !isNil(o.EnabledAlertType) {
		toSerialize["enabledAlertType"] = o.EnabledAlertType
	}
	if !isNil(o.DisabledAlertType) {
		toSerialize["disabledAlertType"] = o.DisabledAlertType
	}
	return json.Marshal(toSerialize)
}

type NullableTwilioAlertHandlerShared struct {
	value *TwilioAlertHandlerShared
	isSet bool
}

func (v NullableTwilioAlertHandlerShared) Get() *TwilioAlertHandlerShared {
	return v.value
}

func (v *NullableTwilioAlertHandlerShared) Set(val *TwilioAlertHandlerShared) {
	v.value = val
	v.isSet = true
}

func (v NullableTwilioAlertHandlerShared) IsSet() bool {
	return v.isSet
}

func (v *NullableTwilioAlertHandlerShared) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTwilioAlertHandlerShared(val *TwilioAlertHandlerShared) *NullableTwilioAlertHandlerShared {
	return &NullableTwilioAlertHandlerShared{value: val, isSet: true}
}

func (v NullableTwilioAlertHandlerShared) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTwilioAlertHandlerShared) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


