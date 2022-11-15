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

// TwilioOtpDeliveryMechanismResponse struct for TwilioOtpDeliveryMechanismResponse
type TwilioOtpDeliveryMechanismResponse struct {
	// Name of the OTP Delivery Mechanism
	Id string `json:"id"`
	Schemas []EnumtwilioOtpDeliveryMechanismSchemaUrn `json:"schemas"`
	// The unique identifier assigned to the Twilio account that will be used.
	TwilioAccountSID string `json:"twilioAccountSID"`
	// The auth token for the Twilio account that will be used.
	TwilioAuthToken *string `json:"twilioAuthToken,omitempty"`
	// The passphrase provider that may be used to obtain the auth token for the Twilio account that will be used.
	TwilioAuthTokenPassphraseProvider *string `json:"twilioAuthTokenPassphraseProvider,omitempty"`
	// The name or OID of the attribute in the user's entry that holds the phone number to which the message should be sent.
	PhoneNumberAttributeType string `json:"phoneNumberAttributeType"`
	// The name of the JSON field whose value is the phone number to which the message should be sent. The phone number must be contained in a top-level field whose value is a single string.
	PhoneNumberJSONField *string `json:"phoneNumberJSONField,omitempty"`
	// A JSON object filter that may be used to identify which phone number value to use when sending the message.
	PhoneNumberJSONObjectFilter *string `json:"phoneNumberJSONObjectFilter,omitempty"`
	SenderPhoneNumber []string `json:"senderPhoneNumber"`
	// Any text that should appear in the message before the one-time password value.
	MessageTextBeforeOTP *string `json:"messageTextBeforeOTP,omitempty"`
	// Any text that should appear in the message after the one-time password value.
	MessageTextAfterOTP *string `json:"messageTextAfterOTP,omitempty"`
	// A description for this OTP Delivery Mechanism
	Description *string `json:"description,omitempty"`
	// Indicates whether this OTP Delivery Mechanism is enabled for use in the server.
	Enabled bool `json:"enabled"`
}

// NewTwilioOtpDeliveryMechanismResponse instantiates a new TwilioOtpDeliveryMechanismResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTwilioOtpDeliveryMechanismResponse(id string, schemas []EnumtwilioOtpDeliveryMechanismSchemaUrn, twilioAccountSID string, phoneNumberAttributeType string, senderPhoneNumber []string, enabled bool) *TwilioOtpDeliveryMechanismResponse {
	this := TwilioOtpDeliveryMechanismResponse{}
	this.Id = id
	this.Schemas = schemas
	this.TwilioAccountSID = twilioAccountSID
	this.PhoneNumberAttributeType = phoneNumberAttributeType
	this.SenderPhoneNumber = senderPhoneNumber
	this.Enabled = enabled
	return &this
}

// NewTwilioOtpDeliveryMechanismResponseWithDefaults instantiates a new TwilioOtpDeliveryMechanismResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTwilioOtpDeliveryMechanismResponseWithDefaults() *TwilioOtpDeliveryMechanismResponse {
	this := TwilioOtpDeliveryMechanismResponse{}
	return &this
}

// GetId returns the Id field value
func (o *TwilioOtpDeliveryMechanismResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TwilioOtpDeliveryMechanismResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TwilioOtpDeliveryMechanismResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *TwilioOtpDeliveryMechanismResponse) GetSchemas() []EnumtwilioOtpDeliveryMechanismSchemaUrn {
	if o == nil {
		var ret []EnumtwilioOtpDeliveryMechanismSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *TwilioOtpDeliveryMechanismResponse) GetSchemasOk() ([]EnumtwilioOtpDeliveryMechanismSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *TwilioOtpDeliveryMechanismResponse) SetSchemas(v []EnumtwilioOtpDeliveryMechanismSchemaUrn) {
	o.Schemas = v
}

// GetTwilioAccountSID returns the TwilioAccountSID field value
func (o *TwilioOtpDeliveryMechanismResponse) GetTwilioAccountSID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TwilioAccountSID
}

// GetTwilioAccountSIDOk returns a tuple with the TwilioAccountSID field value
// and a boolean to check if the value has been set.
func (o *TwilioOtpDeliveryMechanismResponse) GetTwilioAccountSIDOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TwilioAccountSID, true
}

// SetTwilioAccountSID sets field value
func (o *TwilioOtpDeliveryMechanismResponse) SetTwilioAccountSID(v string) {
	o.TwilioAccountSID = v
}

// GetTwilioAuthToken returns the TwilioAuthToken field value if set, zero value otherwise.
func (o *TwilioOtpDeliveryMechanismResponse) GetTwilioAuthToken() string {
	if o == nil || isNil(o.TwilioAuthToken) {
		var ret string
		return ret
	}
	return *o.TwilioAuthToken
}

// GetTwilioAuthTokenOk returns a tuple with the TwilioAuthToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioOtpDeliveryMechanismResponse) GetTwilioAuthTokenOk() (*string, bool) {
	if o == nil || isNil(o.TwilioAuthToken) {
    return nil, false
	}
	return o.TwilioAuthToken, true
}

// HasTwilioAuthToken returns a boolean if a field has been set.
func (o *TwilioOtpDeliveryMechanismResponse) HasTwilioAuthToken() bool {
	if o != nil && !isNil(o.TwilioAuthToken) {
		return true
	}

	return false
}

// SetTwilioAuthToken gets a reference to the given string and assigns it to the TwilioAuthToken field.
func (o *TwilioOtpDeliveryMechanismResponse) SetTwilioAuthToken(v string) {
	o.TwilioAuthToken = &v
}

// GetTwilioAuthTokenPassphraseProvider returns the TwilioAuthTokenPassphraseProvider field value if set, zero value otherwise.
func (o *TwilioOtpDeliveryMechanismResponse) GetTwilioAuthTokenPassphraseProvider() string {
	if o == nil || isNil(o.TwilioAuthTokenPassphraseProvider) {
		var ret string
		return ret
	}
	return *o.TwilioAuthTokenPassphraseProvider
}

// GetTwilioAuthTokenPassphraseProviderOk returns a tuple with the TwilioAuthTokenPassphraseProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioOtpDeliveryMechanismResponse) GetTwilioAuthTokenPassphraseProviderOk() (*string, bool) {
	if o == nil || isNil(o.TwilioAuthTokenPassphraseProvider) {
    return nil, false
	}
	return o.TwilioAuthTokenPassphraseProvider, true
}

// HasTwilioAuthTokenPassphraseProvider returns a boolean if a field has been set.
func (o *TwilioOtpDeliveryMechanismResponse) HasTwilioAuthTokenPassphraseProvider() bool {
	if o != nil && !isNil(o.TwilioAuthTokenPassphraseProvider) {
		return true
	}

	return false
}

// SetTwilioAuthTokenPassphraseProvider gets a reference to the given string and assigns it to the TwilioAuthTokenPassphraseProvider field.
func (o *TwilioOtpDeliveryMechanismResponse) SetTwilioAuthTokenPassphraseProvider(v string) {
	o.TwilioAuthTokenPassphraseProvider = &v
}

// GetPhoneNumberAttributeType returns the PhoneNumberAttributeType field value
func (o *TwilioOtpDeliveryMechanismResponse) GetPhoneNumberAttributeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PhoneNumberAttributeType
}

// GetPhoneNumberAttributeTypeOk returns a tuple with the PhoneNumberAttributeType field value
// and a boolean to check if the value has been set.
func (o *TwilioOtpDeliveryMechanismResponse) GetPhoneNumberAttributeTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PhoneNumberAttributeType, true
}

// SetPhoneNumberAttributeType sets field value
func (o *TwilioOtpDeliveryMechanismResponse) SetPhoneNumberAttributeType(v string) {
	o.PhoneNumberAttributeType = v
}

// GetPhoneNumberJSONField returns the PhoneNumberJSONField field value if set, zero value otherwise.
func (o *TwilioOtpDeliveryMechanismResponse) GetPhoneNumberJSONField() string {
	if o == nil || isNil(o.PhoneNumberJSONField) {
		var ret string
		return ret
	}
	return *o.PhoneNumberJSONField
}

// GetPhoneNumberJSONFieldOk returns a tuple with the PhoneNumberJSONField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioOtpDeliveryMechanismResponse) GetPhoneNumberJSONFieldOk() (*string, bool) {
	if o == nil || isNil(o.PhoneNumberJSONField) {
    return nil, false
	}
	return o.PhoneNumberJSONField, true
}

// HasPhoneNumberJSONField returns a boolean if a field has been set.
func (o *TwilioOtpDeliveryMechanismResponse) HasPhoneNumberJSONField() bool {
	if o != nil && !isNil(o.PhoneNumberJSONField) {
		return true
	}

	return false
}

// SetPhoneNumberJSONField gets a reference to the given string and assigns it to the PhoneNumberJSONField field.
func (o *TwilioOtpDeliveryMechanismResponse) SetPhoneNumberJSONField(v string) {
	o.PhoneNumberJSONField = &v
}

// GetPhoneNumberJSONObjectFilter returns the PhoneNumberJSONObjectFilter field value if set, zero value otherwise.
func (o *TwilioOtpDeliveryMechanismResponse) GetPhoneNumberJSONObjectFilter() string {
	if o == nil || isNil(o.PhoneNumberJSONObjectFilter) {
		var ret string
		return ret
	}
	return *o.PhoneNumberJSONObjectFilter
}

// GetPhoneNumberJSONObjectFilterOk returns a tuple with the PhoneNumberJSONObjectFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioOtpDeliveryMechanismResponse) GetPhoneNumberJSONObjectFilterOk() (*string, bool) {
	if o == nil || isNil(o.PhoneNumberJSONObjectFilter) {
    return nil, false
	}
	return o.PhoneNumberJSONObjectFilter, true
}

// HasPhoneNumberJSONObjectFilter returns a boolean if a field has been set.
func (o *TwilioOtpDeliveryMechanismResponse) HasPhoneNumberJSONObjectFilter() bool {
	if o != nil && !isNil(o.PhoneNumberJSONObjectFilter) {
		return true
	}

	return false
}

// SetPhoneNumberJSONObjectFilter gets a reference to the given string and assigns it to the PhoneNumberJSONObjectFilter field.
func (o *TwilioOtpDeliveryMechanismResponse) SetPhoneNumberJSONObjectFilter(v string) {
	o.PhoneNumberJSONObjectFilter = &v
}

// GetSenderPhoneNumber returns the SenderPhoneNumber field value
func (o *TwilioOtpDeliveryMechanismResponse) GetSenderPhoneNumber() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SenderPhoneNumber
}

// GetSenderPhoneNumberOk returns a tuple with the SenderPhoneNumber field value
// and a boolean to check if the value has been set.
func (o *TwilioOtpDeliveryMechanismResponse) GetSenderPhoneNumberOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.SenderPhoneNumber, true
}

// SetSenderPhoneNumber sets field value
func (o *TwilioOtpDeliveryMechanismResponse) SetSenderPhoneNumber(v []string) {
	o.SenderPhoneNumber = v
}

// GetMessageTextBeforeOTP returns the MessageTextBeforeOTP field value if set, zero value otherwise.
func (o *TwilioOtpDeliveryMechanismResponse) GetMessageTextBeforeOTP() string {
	if o == nil || isNil(o.MessageTextBeforeOTP) {
		var ret string
		return ret
	}
	return *o.MessageTextBeforeOTP
}

// GetMessageTextBeforeOTPOk returns a tuple with the MessageTextBeforeOTP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioOtpDeliveryMechanismResponse) GetMessageTextBeforeOTPOk() (*string, bool) {
	if o == nil || isNil(o.MessageTextBeforeOTP) {
    return nil, false
	}
	return o.MessageTextBeforeOTP, true
}

// HasMessageTextBeforeOTP returns a boolean if a field has been set.
func (o *TwilioOtpDeliveryMechanismResponse) HasMessageTextBeforeOTP() bool {
	if o != nil && !isNil(o.MessageTextBeforeOTP) {
		return true
	}

	return false
}

// SetMessageTextBeforeOTP gets a reference to the given string and assigns it to the MessageTextBeforeOTP field.
func (o *TwilioOtpDeliveryMechanismResponse) SetMessageTextBeforeOTP(v string) {
	o.MessageTextBeforeOTP = &v
}

// GetMessageTextAfterOTP returns the MessageTextAfterOTP field value if set, zero value otherwise.
func (o *TwilioOtpDeliveryMechanismResponse) GetMessageTextAfterOTP() string {
	if o == nil || isNil(o.MessageTextAfterOTP) {
		var ret string
		return ret
	}
	return *o.MessageTextAfterOTP
}

// GetMessageTextAfterOTPOk returns a tuple with the MessageTextAfterOTP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioOtpDeliveryMechanismResponse) GetMessageTextAfterOTPOk() (*string, bool) {
	if o == nil || isNil(o.MessageTextAfterOTP) {
    return nil, false
	}
	return o.MessageTextAfterOTP, true
}

// HasMessageTextAfterOTP returns a boolean if a field has been set.
func (o *TwilioOtpDeliveryMechanismResponse) HasMessageTextAfterOTP() bool {
	if o != nil && !isNil(o.MessageTextAfterOTP) {
		return true
	}

	return false
}

// SetMessageTextAfterOTP gets a reference to the given string and assigns it to the MessageTextAfterOTP field.
func (o *TwilioOtpDeliveryMechanismResponse) SetMessageTextAfterOTP(v string) {
	o.MessageTextAfterOTP = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TwilioOtpDeliveryMechanismResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioOtpDeliveryMechanismResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TwilioOtpDeliveryMechanismResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TwilioOtpDeliveryMechanismResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *TwilioOtpDeliveryMechanismResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *TwilioOtpDeliveryMechanismResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *TwilioOtpDeliveryMechanismResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o TwilioOtpDeliveryMechanismResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
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
		toSerialize["phoneNumberAttributeType"] = o.PhoneNumberAttributeType
	}
	if !isNil(o.PhoneNumberJSONField) {
		toSerialize["phoneNumberJSONField"] = o.PhoneNumberJSONField
	}
	if !isNil(o.PhoneNumberJSONObjectFilter) {
		toSerialize["phoneNumberJSONObjectFilter"] = o.PhoneNumberJSONObjectFilter
	}
	if true {
		toSerialize["senderPhoneNumber"] = o.SenderPhoneNumber
	}
	if !isNil(o.MessageTextBeforeOTP) {
		toSerialize["messageTextBeforeOTP"] = o.MessageTextBeforeOTP
	}
	if !isNil(o.MessageTextAfterOTP) {
		toSerialize["messageTextAfterOTP"] = o.MessageTextAfterOTP
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableTwilioOtpDeliveryMechanismResponse struct {
	value *TwilioOtpDeliveryMechanismResponse
	isSet bool
}

func (v NullableTwilioOtpDeliveryMechanismResponse) Get() *TwilioOtpDeliveryMechanismResponse {
	return v.value
}

func (v *NullableTwilioOtpDeliveryMechanismResponse) Set(val *TwilioOtpDeliveryMechanismResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTwilioOtpDeliveryMechanismResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTwilioOtpDeliveryMechanismResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTwilioOtpDeliveryMechanismResponse(val *TwilioOtpDeliveryMechanismResponse) *NullableTwilioOtpDeliveryMechanismResponse {
	return &NullableTwilioOtpDeliveryMechanismResponse{value: val, isSet: true}
}

func (v NullableTwilioOtpDeliveryMechanismResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTwilioOtpDeliveryMechanismResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


