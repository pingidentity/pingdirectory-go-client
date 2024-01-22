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

// checks if the AddSmtpAlertHandlerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddSmtpAlertHandlerRequest{}

// AddSmtpAlertHandlerRequest struct for AddSmtpAlertHandlerRequest
type AddSmtpAlertHandlerRequest struct {
	Schemas []EnumsmtpAlertHandlerSchemaUrn `json:"schemas"`
	// Indicates whether the server should attempt to invoke this SMTP Alert Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver the alert notification) will not delay whatever processing the server was performing when the alert was generated.
	Asynchronous *bool `json:"asynchronous,omitempty"`
	// Specifies the email address to use as the sender for messages generated by this alert handler.
	SenderAddress string `json:"senderAddress"`
	// Specifies an email address to which the messages should be sent.
	RecipientAddress []string `json:"recipientAddress"`
	// Specifies the subject that should be used for email messages generated by this alert handler.
	MessageSubject *string `json:"messageSubject,omitempty"`
	// Specifies the body that should be used for email messages generated by this alert handler.
	MessageBody *string `json:"messageBody,omitempty"`
	// Include monitor entries that match this filter.
	IncludeMonitorDataFilter *string `json:"includeMonitorDataFilter,omitempty"`
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

// NewAddSmtpAlertHandlerRequest instantiates a new AddSmtpAlertHandlerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSmtpAlertHandlerRequest(schemas []EnumsmtpAlertHandlerSchemaUrn, senderAddress string, recipientAddress []string, enabled bool, handlerName string) *AddSmtpAlertHandlerRequest {
	this := AddSmtpAlertHandlerRequest{}
	this.Schemas = schemas
	this.SenderAddress = senderAddress
	this.RecipientAddress = recipientAddress
	this.Enabled = enabled
	this.HandlerName = handlerName
	return &this
}

// NewAddSmtpAlertHandlerRequestWithDefaults instantiates a new AddSmtpAlertHandlerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSmtpAlertHandlerRequestWithDefaults() *AddSmtpAlertHandlerRequest {
	this := AddSmtpAlertHandlerRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddSmtpAlertHandlerRequest) GetSchemas() []EnumsmtpAlertHandlerSchemaUrn {
	if o == nil {
		var ret []EnumsmtpAlertHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddSmtpAlertHandlerRequest) GetSchemasOk() ([]EnumsmtpAlertHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddSmtpAlertHandlerRequest) SetSchemas(v []EnumsmtpAlertHandlerSchemaUrn) {
	o.Schemas = v
}

// GetAsynchronous returns the Asynchronous field value if set, zero value otherwise.
func (o *AddSmtpAlertHandlerRequest) GetAsynchronous() bool {
	if o == nil || IsNil(o.Asynchronous) {
		var ret bool
		return ret
	}
	return *o.Asynchronous
}

// GetAsynchronousOk returns a tuple with the Asynchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSmtpAlertHandlerRequest) GetAsynchronousOk() (*bool, bool) {
	if o == nil || IsNil(o.Asynchronous) {
		return nil, false
	}
	return o.Asynchronous, true
}

// HasAsynchronous returns a boolean if a field has been set.
func (o *AddSmtpAlertHandlerRequest) HasAsynchronous() bool {
	if o != nil && !IsNil(o.Asynchronous) {
		return true
	}

	return false
}

// SetAsynchronous gets a reference to the given bool and assigns it to the Asynchronous field.
func (o *AddSmtpAlertHandlerRequest) SetAsynchronous(v bool) {
	o.Asynchronous = &v
}

// GetSenderAddress returns the SenderAddress field value
func (o *AddSmtpAlertHandlerRequest) GetSenderAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SenderAddress
}

// GetSenderAddressOk returns a tuple with the SenderAddress field value
// and a boolean to check if the value has been set.
func (o *AddSmtpAlertHandlerRequest) GetSenderAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderAddress, true
}

// SetSenderAddress sets field value
func (o *AddSmtpAlertHandlerRequest) SetSenderAddress(v string) {
	o.SenderAddress = v
}

// GetRecipientAddress returns the RecipientAddress field value
func (o *AddSmtpAlertHandlerRequest) GetRecipientAddress() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RecipientAddress
}

// GetRecipientAddressOk returns a tuple with the RecipientAddress field value
// and a boolean to check if the value has been set.
func (o *AddSmtpAlertHandlerRequest) GetRecipientAddressOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecipientAddress, true
}

// SetRecipientAddress sets field value
func (o *AddSmtpAlertHandlerRequest) SetRecipientAddress(v []string) {
	o.RecipientAddress = v
}

// GetMessageSubject returns the MessageSubject field value if set, zero value otherwise.
func (o *AddSmtpAlertHandlerRequest) GetMessageSubject() string {
	if o == nil || IsNil(o.MessageSubject) {
		var ret string
		return ret
	}
	return *o.MessageSubject
}

// GetMessageSubjectOk returns a tuple with the MessageSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSmtpAlertHandlerRequest) GetMessageSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.MessageSubject) {
		return nil, false
	}
	return o.MessageSubject, true
}

// HasMessageSubject returns a boolean if a field has been set.
func (o *AddSmtpAlertHandlerRequest) HasMessageSubject() bool {
	if o != nil && !IsNil(o.MessageSubject) {
		return true
	}

	return false
}

// SetMessageSubject gets a reference to the given string and assigns it to the MessageSubject field.
func (o *AddSmtpAlertHandlerRequest) SetMessageSubject(v string) {
	o.MessageSubject = &v
}

// GetMessageBody returns the MessageBody field value if set, zero value otherwise.
func (o *AddSmtpAlertHandlerRequest) GetMessageBody() string {
	if o == nil || IsNil(o.MessageBody) {
		var ret string
		return ret
	}
	return *o.MessageBody
}

// GetMessageBodyOk returns a tuple with the MessageBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSmtpAlertHandlerRequest) GetMessageBodyOk() (*string, bool) {
	if o == nil || IsNil(o.MessageBody) {
		return nil, false
	}
	return o.MessageBody, true
}

// HasMessageBody returns a boolean if a field has been set.
func (o *AddSmtpAlertHandlerRequest) HasMessageBody() bool {
	if o != nil && !IsNil(o.MessageBody) {
		return true
	}

	return false
}

// SetMessageBody gets a reference to the given string and assigns it to the MessageBody field.
func (o *AddSmtpAlertHandlerRequest) SetMessageBody(v string) {
	o.MessageBody = &v
}

// GetIncludeMonitorDataFilter returns the IncludeMonitorDataFilter field value if set, zero value otherwise.
func (o *AddSmtpAlertHandlerRequest) GetIncludeMonitorDataFilter() string {
	if o == nil || IsNil(o.IncludeMonitorDataFilter) {
		var ret string
		return ret
	}
	return *o.IncludeMonitorDataFilter
}

// GetIncludeMonitorDataFilterOk returns a tuple with the IncludeMonitorDataFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSmtpAlertHandlerRequest) GetIncludeMonitorDataFilterOk() (*string, bool) {
	if o == nil || IsNil(o.IncludeMonitorDataFilter) {
		return nil, false
	}
	return o.IncludeMonitorDataFilter, true
}

// HasIncludeMonitorDataFilter returns a boolean if a field has been set.
func (o *AddSmtpAlertHandlerRequest) HasIncludeMonitorDataFilter() bool {
	if o != nil && !IsNil(o.IncludeMonitorDataFilter) {
		return true
	}

	return false
}

// SetIncludeMonitorDataFilter gets a reference to the given string and assigns it to the IncludeMonitorDataFilter field.
func (o *AddSmtpAlertHandlerRequest) SetIncludeMonitorDataFilter(v string) {
	o.IncludeMonitorDataFilter = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddSmtpAlertHandlerRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSmtpAlertHandlerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddSmtpAlertHandlerRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddSmtpAlertHandlerRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddSmtpAlertHandlerRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddSmtpAlertHandlerRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddSmtpAlertHandlerRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEnabledAlertSeverity returns the EnabledAlertSeverity field value if set, zero value otherwise.
func (o *AddSmtpAlertHandlerRequest) GetEnabledAlertSeverity() []EnumalertHandlerEnabledAlertSeverityProp {
	if o == nil || IsNil(o.EnabledAlertSeverity) {
		var ret []EnumalertHandlerEnabledAlertSeverityProp
		return ret
	}
	return o.EnabledAlertSeverity
}

// GetEnabledAlertSeverityOk returns a tuple with the EnabledAlertSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSmtpAlertHandlerRequest) GetEnabledAlertSeverityOk() ([]EnumalertHandlerEnabledAlertSeverityProp, bool) {
	if o == nil || IsNil(o.EnabledAlertSeverity) {
		return nil, false
	}
	return o.EnabledAlertSeverity, true
}

// HasEnabledAlertSeverity returns a boolean if a field has been set.
func (o *AddSmtpAlertHandlerRequest) HasEnabledAlertSeverity() bool {
	if o != nil && !IsNil(o.EnabledAlertSeverity) {
		return true
	}

	return false
}

// SetEnabledAlertSeverity gets a reference to the given []EnumalertHandlerEnabledAlertSeverityProp and assigns it to the EnabledAlertSeverity field.
func (o *AddSmtpAlertHandlerRequest) SetEnabledAlertSeverity(v []EnumalertHandlerEnabledAlertSeverityProp) {
	o.EnabledAlertSeverity = v
}

// GetEnabledAlertType returns the EnabledAlertType field value if set, zero value otherwise.
func (o *AddSmtpAlertHandlerRequest) GetEnabledAlertType() []EnumalertHandlerEnabledAlertTypeProp {
	if o == nil || IsNil(o.EnabledAlertType) {
		var ret []EnumalertHandlerEnabledAlertTypeProp
		return ret
	}
	return o.EnabledAlertType
}

// GetEnabledAlertTypeOk returns a tuple with the EnabledAlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSmtpAlertHandlerRequest) GetEnabledAlertTypeOk() ([]EnumalertHandlerEnabledAlertTypeProp, bool) {
	if o == nil || IsNil(o.EnabledAlertType) {
		return nil, false
	}
	return o.EnabledAlertType, true
}

// HasEnabledAlertType returns a boolean if a field has been set.
func (o *AddSmtpAlertHandlerRequest) HasEnabledAlertType() bool {
	if o != nil && !IsNil(o.EnabledAlertType) {
		return true
	}

	return false
}

// SetEnabledAlertType gets a reference to the given []EnumalertHandlerEnabledAlertTypeProp and assigns it to the EnabledAlertType field.
func (o *AddSmtpAlertHandlerRequest) SetEnabledAlertType(v []EnumalertHandlerEnabledAlertTypeProp) {
	o.EnabledAlertType = v
}

// GetDisabledAlertType returns the DisabledAlertType field value if set, zero value otherwise.
func (o *AddSmtpAlertHandlerRequest) GetDisabledAlertType() []EnumalertHandlerDisabledAlertTypeProp {
	if o == nil || IsNil(o.DisabledAlertType) {
		var ret []EnumalertHandlerDisabledAlertTypeProp
		return ret
	}
	return o.DisabledAlertType
}

// GetDisabledAlertTypeOk returns a tuple with the DisabledAlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSmtpAlertHandlerRequest) GetDisabledAlertTypeOk() ([]EnumalertHandlerDisabledAlertTypeProp, bool) {
	if o == nil || IsNil(o.DisabledAlertType) {
		return nil, false
	}
	return o.DisabledAlertType, true
}

// HasDisabledAlertType returns a boolean if a field has been set.
func (o *AddSmtpAlertHandlerRequest) HasDisabledAlertType() bool {
	if o != nil && !IsNil(o.DisabledAlertType) {
		return true
	}

	return false
}

// SetDisabledAlertType gets a reference to the given []EnumalertHandlerDisabledAlertTypeProp and assigns it to the DisabledAlertType field.
func (o *AddSmtpAlertHandlerRequest) SetDisabledAlertType(v []EnumalertHandlerDisabledAlertTypeProp) {
	o.DisabledAlertType = v
}

// GetHandlerName returns the HandlerName field value
func (o *AddSmtpAlertHandlerRequest) GetHandlerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HandlerName
}

// GetHandlerNameOk returns a tuple with the HandlerName field value
// and a boolean to check if the value has been set.
func (o *AddSmtpAlertHandlerRequest) GetHandlerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HandlerName, true
}

// SetHandlerName sets field value
func (o *AddSmtpAlertHandlerRequest) SetHandlerName(v string) {
	o.HandlerName = v
}

func (o AddSmtpAlertHandlerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddSmtpAlertHandlerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.Asynchronous) {
		toSerialize["asynchronous"] = o.Asynchronous
	}
	toSerialize["senderAddress"] = o.SenderAddress
	toSerialize["recipientAddress"] = o.RecipientAddress
	if !IsNil(o.MessageSubject) {
		toSerialize["messageSubject"] = o.MessageSubject
	}
	if !IsNil(o.MessageBody) {
		toSerialize["messageBody"] = o.MessageBody
	}
	if !IsNil(o.IncludeMonitorDataFilter) {
		toSerialize["includeMonitorDataFilter"] = o.IncludeMonitorDataFilter
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

type NullableAddSmtpAlertHandlerRequest struct {
	value *AddSmtpAlertHandlerRequest
	isSet bool
}

func (v NullableAddSmtpAlertHandlerRequest) Get() *AddSmtpAlertHandlerRequest {
	return v.value
}

func (v *NullableAddSmtpAlertHandlerRequest) Set(val *AddSmtpAlertHandlerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSmtpAlertHandlerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSmtpAlertHandlerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSmtpAlertHandlerRequest(val *AddSmtpAlertHandlerRequest) *NullableAddSmtpAlertHandlerRequest {
	return &NullableAddSmtpAlertHandlerRequest{value: val, isSet: true}
}

func (v NullableAddSmtpAlertHandlerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSmtpAlertHandlerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
