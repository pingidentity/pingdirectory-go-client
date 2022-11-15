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

// EmailOtpDeliveryMechanismResponse struct for EmailOtpDeliveryMechanismResponse
type EmailOtpDeliveryMechanismResponse struct {
	// Name of the OTP Delivery Mechanism
	Id string `json:"id"`
	Schemas []EnumemailOtpDeliveryMechanismSchemaUrn `json:"schemas"`
	// The name or OID of the attribute that holds the email address to which the message should be sent.
	EmailAddressAttributeType string `json:"emailAddressAttributeType"`
	// The name of the JSON field whose value is the email address to which the message should be sent. The email address must be contained in a top-level field whose value is a single string.
	EmailAddressJSONField *string `json:"emailAddressJSONField,omitempty"`
	// A JSON object filter that may be used to identify which email address value to use when sending the message.
	EmailAddressJSONObjectFilter *string `json:"emailAddressJSONObjectFilter,omitempty"`
	// The e-mail address to use as the sender for the one-time password.
	SenderAddress string `json:"senderAddress"`
	// The subject to use for the e-mail message.
	MessageSubject string `json:"messageSubject"`
	// Any text that should appear in the message before the one-time password value.
	MessageTextBeforeOTP *string `json:"messageTextBeforeOTP,omitempty"`
	// Any text that should appear in the message after the one-time password value.
	MessageTextAfterOTP *string `json:"messageTextAfterOTP,omitempty"`
	// A description for this OTP Delivery Mechanism
	Description *string `json:"description,omitempty"`
	// Indicates whether this OTP Delivery Mechanism is enabled for use in the server.
	Enabled bool `json:"enabled"`
}

// NewEmailOtpDeliveryMechanismResponse instantiates a new EmailOtpDeliveryMechanismResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailOtpDeliveryMechanismResponse(id string, schemas []EnumemailOtpDeliveryMechanismSchemaUrn, emailAddressAttributeType string, senderAddress string, messageSubject string, enabled bool) *EmailOtpDeliveryMechanismResponse {
	this := EmailOtpDeliveryMechanismResponse{}
	this.Id = id
	this.Schemas = schemas
	this.EmailAddressAttributeType = emailAddressAttributeType
	this.SenderAddress = senderAddress
	this.MessageSubject = messageSubject
	this.Enabled = enabled
	return &this
}

// NewEmailOtpDeliveryMechanismResponseWithDefaults instantiates a new EmailOtpDeliveryMechanismResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailOtpDeliveryMechanismResponseWithDefaults() *EmailOtpDeliveryMechanismResponse {
	this := EmailOtpDeliveryMechanismResponse{}
	return &this
}

// GetId returns the Id field value
func (o *EmailOtpDeliveryMechanismResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EmailOtpDeliveryMechanismResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EmailOtpDeliveryMechanismResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *EmailOtpDeliveryMechanismResponse) GetSchemas() []EnumemailOtpDeliveryMechanismSchemaUrn {
	if o == nil {
		var ret []EnumemailOtpDeliveryMechanismSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *EmailOtpDeliveryMechanismResponse) GetSchemasOk() ([]EnumemailOtpDeliveryMechanismSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *EmailOtpDeliveryMechanismResponse) SetSchemas(v []EnumemailOtpDeliveryMechanismSchemaUrn) {
	o.Schemas = v
}

// GetEmailAddressAttributeType returns the EmailAddressAttributeType field value
func (o *EmailOtpDeliveryMechanismResponse) GetEmailAddressAttributeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddressAttributeType
}

// GetEmailAddressAttributeTypeOk returns a tuple with the EmailAddressAttributeType field value
// and a boolean to check if the value has been set.
func (o *EmailOtpDeliveryMechanismResponse) GetEmailAddressAttributeTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EmailAddressAttributeType, true
}

// SetEmailAddressAttributeType sets field value
func (o *EmailOtpDeliveryMechanismResponse) SetEmailAddressAttributeType(v string) {
	o.EmailAddressAttributeType = v
}

// GetEmailAddressJSONField returns the EmailAddressJSONField field value if set, zero value otherwise.
func (o *EmailOtpDeliveryMechanismResponse) GetEmailAddressJSONField() string {
	if o == nil || isNil(o.EmailAddressJSONField) {
		var ret string
		return ret
	}
	return *o.EmailAddressJSONField
}

// GetEmailAddressJSONFieldOk returns a tuple with the EmailAddressJSONField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailOtpDeliveryMechanismResponse) GetEmailAddressJSONFieldOk() (*string, bool) {
	if o == nil || isNil(o.EmailAddressJSONField) {
    return nil, false
	}
	return o.EmailAddressJSONField, true
}

// HasEmailAddressJSONField returns a boolean if a field has been set.
func (o *EmailOtpDeliveryMechanismResponse) HasEmailAddressJSONField() bool {
	if o != nil && !isNil(o.EmailAddressJSONField) {
		return true
	}

	return false
}

// SetEmailAddressJSONField gets a reference to the given string and assigns it to the EmailAddressJSONField field.
func (o *EmailOtpDeliveryMechanismResponse) SetEmailAddressJSONField(v string) {
	o.EmailAddressJSONField = &v
}

// GetEmailAddressJSONObjectFilter returns the EmailAddressJSONObjectFilter field value if set, zero value otherwise.
func (o *EmailOtpDeliveryMechanismResponse) GetEmailAddressJSONObjectFilter() string {
	if o == nil || isNil(o.EmailAddressJSONObjectFilter) {
		var ret string
		return ret
	}
	return *o.EmailAddressJSONObjectFilter
}

// GetEmailAddressJSONObjectFilterOk returns a tuple with the EmailAddressJSONObjectFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailOtpDeliveryMechanismResponse) GetEmailAddressJSONObjectFilterOk() (*string, bool) {
	if o == nil || isNil(o.EmailAddressJSONObjectFilter) {
    return nil, false
	}
	return o.EmailAddressJSONObjectFilter, true
}

// HasEmailAddressJSONObjectFilter returns a boolean if a field has been set.
func (o *EmailOtpDeliveryMechanismResponse) HasEmailAddressJSONObjectFilter() bool {
	if o != nil && !isNil(o.EmailAddressJSONObjectFilter) {
		return true
	}

	return false
}

// SetEmailAddressJSONObjectFilter gets a reference to the given string and assigns it to the EmailAddressJSONObjectFilter field.
func (o *EmailOtpDeliveryMechanismResponse) SetEmailAddressJSONObjectFilter(v string) {
	o.EmailAddressJSONObjectFilter = &v
}

// GetSenderAddress returns the SenderAddress field value
func (o *EmailOtpDeliveryMechanismResponse) GetSenderAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SenderAddress
}

// GetSenderAddressOk returns a tuple with the SenderAddress field value
// and a boolean to check if the value has been set.
func (o *EmailOtpDeliveryMechanismResponse) GetSenderAddressOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SenderAddress, true
}

// SetSenderAddress sets field value
func (o *EmailOtpDeliveryMechanismResponse) SetSenderAddress(v string) {
	o.SenderAddress = v
}

// GetMessageSubject returns the MessageSubject field value
func (o *EmailOtpDeliveryMechanismResponse) GetMessageSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageSubject
}

// GetMessageSubjectOk returns a tuple with the MessageSubject field value
// and a boolean to check if the value has been set.
func (o *EmailOtpDeliveryMechanismResponse) GetMessageSubjectOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MessageSubject, true
}

// SetMessageSubject sets field value
func (o *EmailOtpDeliveryMechanismResponse) SetMessageSubject(v string) {
	o.MessageSubject = v
}

// GetMessageTextBeforeOTP returns the MessageTextBeforeOTP field value if set, zero value otherwise.
func (o *EmailOtpDeliveryMechanismResponse) GetMessageTextBeforeOTP() string {
	if o == nil || isNil(o.MessageTextBeforeOTP) {
		var ret string
		return ret
	}
	return *o.MessageTextBeforeOTP
}

// GetMessageTextBeforeOTPOk returns a tuple with the MessageTextBeforeOTP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailOtpDeliveryMechanismResponse) GetMessageTextBeforeOTPOk() (*string, bool) {
	if o == nil || isNil(o.MessageTextBeforeOTP) {
    return nil, false
	}
	return o.MessageTextBeforeOTP, true
}

// HasMessageTextBeforeOTP returns a boolean if a field has been set.
func (o *EmailOtpDeliveryMechanismResponse) HasMessageTextBeforeOTP() bool {
	if o != nil && !isNil(o.MessageTextBeforeOTP) {
		return true
	}

	return false
}

// SetMessageTextBeforeOTP gets a reference to the given string and assigns it to the MessageTextBeforeOTP field.
func (o *EmailOtpDeliveryMechanismResponse) SetMessageTextBeforeOTP(v string) {
	o.MessageTextBeforeOTP = &v
}

// GetMessageTextAfterOTP returns the MessageTextAfterOTP field value if set, zero value otherwise.
func (o *EmailOtpDeliveryMechanismResponse) GetMessageTextAfterOTP() string {
	if o == nil || isNil(o.MessageTextAfterOTP) {
		var ret string
		return ret
	}
	return *o.MessageTextAfterOTP
}

// GetMessageTextAfterOTPOk returns a tuple with the MessageTextAfterOTP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailOtpDeliveryMechanismResponse) GetMessageTextAfterOTPOk() (*string, bool) {
	if o == nil || isNil(o.MessageTextAfterOTP) {
    return nil, false
	}
	return o.MessageTextAfterOTP, true
}

// HasMessageTextAfterOTP returns a boolean if a field has been set.
func (o *EmailOtpDeliveryMechanismResponse) HasMessageTextAfterOTP() bool {
	if o != nil && !isNil(o.MessageTextAfterOTP) {
		return true
	}

	return false
}

// SetMessageTextAfterOTP gets a reference to the given string and assigns it to the MessageTextAfterOTP field.
func (o *EmailOtpDeliveryMechanismResponse) SetMessageTextAfterOTP(v string) {
	o.MessageTextAfterOTP = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EmailOtpDeliveryMechanismResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailOtpDeliveryMechanismResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EmailOtpDeliveryMechanismResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EmailOtpDeliveryMechanismResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *EmailOtpDeliveryMechanismResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *EmailOtpDeliveryMechanismResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *EmailOtpDeliveryMechanismResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o EmailOtpDeliveryMechanismResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["emailAddressAttributeType"] = o.EmailAddressAttributeType
	}
	if !isNil(o.EmailAddressJSONField) {
		toSerialize["emailAddressJSONField"] = o.EmailAddressJSONField
	}
	if !isNil(o.EmailAddressJSONObjectFilter) {
		toSerialize["emailAddressJSONObjectFilter"] = o.EmailAddressJSONObjectFilter
	}
	if true {
		toSerialize["senderAddress"] = o.SenderAddress
	}
	if true {
		toSerialize["messageSubject"] = o.MessageSubject
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

type NullableEmailOtpDeliveryMechanismResponse struct {
	value *EmailOtpDeliveryMechanismResponse
	isSet bool
}

func (v NullableEmailOtpDeliveryMechanismResponse) Get() *EmailOtpDeliveryMechanismResponse {
	return v.value
}

func (v *NullableEmailOtpDeliveryMechanismResponse) Set(val *EmailOtpDeliveryMechanismResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailOtpDeliveryMechanismResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailOtpDeliveryMechanismResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailOtpDeliveryMechanismResponse(val *EmailOtpDeliveryMechanismResponse) *NullableEmailOtpDeliveryMechanismResponse {
	return &NullableEmailOtpDeliveryMechanismResponse{value: val, isSet: true}
}

func (v NullableEmailOtpDeliveryMechanismResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailOtpDeliveryMechanismResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


