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

// checks if the SmtpAlertHandlerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmtpAlertHandlerResponse{}

// SmtpAlertHandlerResponse struct for SmtpAlertHandlerResponse
type SmtpAlertHandlerResponse struct {
	Schemas []EnumsmtpAlertHandlerSchemaUrn `json:"schemas"`
	// Indicates whether the server should attempt to invoke this SMTP Alert Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver the alert notification) will not delay whatever processing the server was performing when the alert was generated.
	Asynchronous *bool `json:"asynchronous,omitempty"`
	// Specifies the email address to use as the sender for messages generated by this alert handler.
	SenderAddress string `json:"senderAddress"`
	// Specifies an email address to which the messages should be sent.
	RecipientAddress []string `json:"recipientAddress"`
	// Specifies the subject that should be used for email messages generated by this alert handler.
	MessageSubject string `json:"messageSubject"`
	// Specifies the body that should be used for email messages generated by this alert handler.
	MessageBody string `json:"messageBody"`
	// Include monitor entries that match this filter.
	IncludeMonitorDataFilter *string `json:"includeMonitorDataFilter,omitempty"`
	// A description for this Alert Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the Alert Handler is enabled.
	Enabled                                       bool                                               `json:"enabled"`
	EnabledAlertSeverity                          []EnumalertHandlerEnabledAlertSeverityProp         `json:"enabledAlertSeverity,omitempty"`
	EnabledAlertType                              []EnumalertHandlerEnabledAlertTypeProp             `json:"enabledAlertType,omitempty"`
	DisabledAlertType                             []EnumalertHandlerDisabledAlertTypeProp            `json:"disabledAlertType,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Alert Handler
	Id string `json:"id"`
}

// NewSmtpAlertHandlerResponse instantiates a new SmtpAlertHandlerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmtpAlertHandlerResponse(schemas []EnumsmtpAlertHandlerSchemaUrn, senderAddress string, recipientAddress []string, messageSubject string, messageBody string, enabled bool, id string) *SmtpAlertHandlerResponse {
	this := SmtpAlertHandlerResponse{}
	this.Schemas = schemas
	this.SenderAddress = senderAddress
	this.RecipientAddress = recipientAddress
	this.MessageSubject = messageSubject
	this.MessageBody = messageBody
	this.Enabled = enabled
	this.Id = id
	return &this
}

// NewSmtpAlertHandlerResponseWithDefaults instantiates a new SmtpAlertHandlerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmtpAlertHandlerResponseWithDefaults() *SmtpAlertHandlerResponse {
	this := SmtpAlertHandlerResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *SmtpAlertHandlerResponse) GetSchemas() []EnumsmtpAlertHandlerSchemaUrn {
	if o == nil {
		var ret []EnumsmtpAlertHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *SmtpAlertHandlerResponse) GetSchemasOk() ([]EnumsmtpAlertHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *SmtpAlertHandlerResponse) SetSchemas(v []EnumsmtpAlertHandlerSchemaUrn) {
	o.Schemas = v
}

// GetAsynchronous returns the Asynchronous field value if set, zero value otherwise.
func (o *SmtpAlertHandlerResponse) GetAsynchronous() bool {
	if o == nil || IsNil(o.Asynchronous) {
		var ret bool
		return ret
	}
	return *o.Asynchronous
}

// GetAsynchronousOk returns a tuple with the Asynchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpAlertHandlerResponse) GetAsynchronousOk() (*bool, bool) {
	if o == nil || IsNil(o.Asynchronous) {
		return nil, false
	}
	return o.Asynchronous, true
}

// HasAsynchronous returns a boolean if a field has been set.
func (o *SmtpAlertHandlerResponse) HasAsynchronous() bool {
	if o != nil && !IsNil(o.Asynchronous) {
		return true
	}

	return false
}

// SetAsynchronous gets a reference to the given bool and assigns it to the Asynchronous field.
func (o *SmtpAlertHandlerResponse) SetAsynchronous(v bool) {
	o.Asynchronous = &v
}

// GetSenderAddress returns the SenderAddress field value
func (o *SmtpAlertHandlerResponse) GetSenderAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SenderAddress
}

// GetSenderAddressOk returns a tuple with the SenderAddress field value
// and a boolean to check if the value has been set.
func (o *SmtpAlertHandlerResponse) GetSenderAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderAddress, true
}

// SetSenderAddress sets field value
func (o *SmtpAlertHandlerResponse) SetSenderAddress(v string) {
	o.SenderAddress = v
}

// GetRecipientAddress returns the RecipientAddress field value
func (o *SmtpAlertHandlerResponse) GetRecipientAddress() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RecipientAddress
}

// GetRecipientAddressOk returns a tuple with the RecipientAddress field value
// and a boolean to check if the value has been set.
func (o *SmtpAlertHandlerResponse) GetRecipientAddressOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecipientAddress, true
}

// SetRecipientAddress sets field value
func (o *SmtpAlertHandlerResponse) SetRecipientAddress(v []string) {
	o.RecipientAddress = v
}

// GetMessageSubject returns the MessageSubject field value
func (o *SmtpAlertHandlerResponse) GetMessageSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageSubject
}

// GetMessageSubjectOk returns a tuple with the MessageSubject field value
// and a boolean to check if the value has been set.
func (o *SmtpAlertHandlerResponse) GetMessageSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageSubject, true
}

// SetMessageSubject sets field value
func (o *SmtpAlertHandlerResponse) SetMessageSubject(v string) {
	o.MessageSubject = v
}

// GetMessageBody returns the MessageBody field value
func (o *SmtpAlertHandlerResponse) GetMessageBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageBody
}

// GetMessageBodyOk returns a tuple with the MessageBody field value
// and a boolean to check if the value has been set.
func (o *SmtpAlertHandlerResponse) GetMessageBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageBody, true
}

// SetMessageBody sets field value
func (o *SmtpAlertHandlerResponse) SetMessageBody(v string) {
	o.MessageBody = v
}

// GetIncludeMonitorDataFilter returns the IncludeMonitorDataFilter field value if set, zero value otherwise.
func (o *SmtpAlertHandlerResponse) GetIncludeMonitorDataFilter() string {
	if o == nil || IsNil(o.IncludeMonitorDataFilter) {
		var ret string
		return ret
	}
	return *o.IncludeMonitorDataFilter
}

// GetIncludeMonitorDataFilterOk returns a tuple with the IncludeMonitorDataFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpAlertHandlerResponse) GetIncludeMonitorDataFilterOk() (*string, bool) {
	if o == nil || IsNil(o.IncludeMonitorDataFilter) {
		return nil, false
	}
	return o.IncludeMonitorDataFilter, true
}

// HasIncludeMonitorDataFilter returns a boolean if a field has been set.
func (o *SmtpAlertHandlerResponse) HasIncludeMonitorDataFilter() bool {
	if o != nil && !IsNil(o.IncludeMonitorDataFilter) {
		return true
	}

	return false
}

// SetIncludeMonitorDataFilter gets a reference to the given string and assigns it to the IncludeMonitorDataFilter field.
func (o *SmtpAlertHandlerResponse) SetIncludeMonitorDataFilter(v string) {
	o.IncludeMonitorDataFilter = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SmtpAlertHandlerResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpAlertHandlerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SmtpAlertHandlerResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SmtpAlertHandlerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *SmtpAlertHandlerResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SmtpAlertHandlerResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SmtpAlertHandlerResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEnabledAlertSeverity returns the EnabledAlertSeverity field value if set, zero value otherwise.
func (o *SmtpAlertHandlerResponse) GetEnabledAlertSeverity() []EnumalertHandlerEnabledAlertSeverityProp {
	if o == nil || IsNil(o.EnabledAlertSeverity) {
		var ret []EnumalertHandlerEnabledAlertSeverityProp
		return ret
	}
	return o.EnabledAlertSeverity
}

// GetEnabledAlertSeverityOk returns a tuple with the EnabledAlertSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpAlertHandlerResponse) GetEnabledAlertSeverityOk() ([]EnumalertHandlerEnabledAlertSeverityProp, bool) {
	if o == nil || IsNil(o.EnabledAlertSeverity) {
		return nil, false
	}
	return o.EnabledAlertSeverity, true
}

// HasEnabledAlertSeverity returns a boolean if a field has been set.
func (o *SmtpAlertHandlerResponse) HasEnabledAlertSeverity() bool {
	if o != nil && !IsNil(o.EnabledAlertSeverity) {
		return true
	}

	return false
}

// SetEnabledAlertSeverity gets a reference to the given []EnumalertHandlerEnabledAlertSeverityProp and assigns it to the EnabledAlertSeverity field.
func (o *SmtpAlertHandlerResponse) SetEnabledAlertSeverity(v []EnumalertHandlerEnabledAlertSeverityProp) {
	o.EnabledAlertSeverity = v
}

// GetEnabledAlertType returns the EnabledAlertType field value if set, zero value otherwise.
func (o *SmtpAlertHandlerResponse) GetEnabledAlertType() []EnumalertHandlerEnabledAlertTypeProp {
	if o == nil || IsNil(o.EnabledAlertType) {
		var ret []EnumalertHandlerEnabledAlertTypeProp
		return ret
	}
	return o.EnabledAlertType
}

// GetEnabledAlertTypeOk returns a tuple with the EnabledAlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpAlertHandlerResponse) GetEnabledAlertTypeOk() ([]EnumalertHandlerEnabledAlertTypeProp, bool) {
	if o == nil || IsNil(o.EnabledAlertType) {
		return nil, false
	}
	return o.EnabledAlertType, true
}

// HasEnabledAlertType returns a boolean if a field has been set.
func (o *SmtpAlertHandlerResponse) HasEnabledAlertType() bool {
	if o != nil && !IsNil(o.EnabledAlertType) {
		return true
	}

	return false
}

// SetEnabledAlertType gets a reference to the given []EnumalertHandlerEnabledAlertTypeProp and assigns it to the EnabledAlertType field.
func (o *SmtpAlertHandlerResponse) SetEnabledAlertType(v []EnumalertHandlerEnabledAlertTypeProp) {
	o.EnabledAlertType = v
}

// GetDisabledAlertType returns the DisabledAlertType field value if set, zero value otherwise.
func (o *SmtpAlertHandlerResponse) GetDisabledAlertType() []EnumalertHandlerDisabledAlertTypeProp {
	if o == nil || IsNil(o.DisabledAlertType) {
		var ret []EnumalertHandlerDisabledAlertTypeProp
		return ret
	}
	return o.DisabledAlertType
}

// GetDisabledAlertTypeOk returns a tuple with the DisabledAlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpAlertHandlerResponse) GetDisabledAlertTypeOk() ([]EnumalertHandlerDisabledAlertTypeProp, bool) {
	if o == nil || IsNil(o.DisabledAlertType) {
		return nil, false
	}
	return o.DisabledAlertType, true
}

// HasDisabledAlertType returns a boolean if a field has been set.
func (o *SmtpAlertHandlerResponse) HasDisabledAlertType() bool {
	if o != nil && !IsNil(o.DisabledAlertType) {
		return true
	}

	return false
}

// SetDisabledAlertType gets a reference to the given []EnumalertHandlerDisabledAlertTypeProp and assigns it to the DisabledAlertType field.
func (o *SmtpAlertHandlerResponse) SetDisabledAlertType(v []EnumalertHandlerDisabledAlertTypeProp) {
	o.DisabledAlertType = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SmtpAlertHandlerResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpAlertHandlerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SmtpAlertHandlerResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *SmtpAlertHandlerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *SmtpAlertHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpAlertHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *SmtpAlertHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *SmtpAlertHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *SmtpAlertHandlerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SmtpAlertHandlerResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SmtpAlertHandlerResponse) SetId(v string) {
	o.Id = v
}

func (o SmtpAlertHandlerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmtpAlertHandlerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.Asynchronous) {
		toSerialize["asynchronous"] = o.Asynchronous
	}
	toSerialize["senderAddress"] = o.SenderAddress
	toSerialize["recipientAddress"] = o.RecipientAddress
	toSerialize["messageSubject"] = o.MessageSubject
	toSerialize["messageBody"] = o.MessageBody
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
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableSmtpAlertHandlerResponse struct {
	value *SmtpAlertHandlerResponse
	isSet bool
}

func (v NullableSmtpAlertHandlerResponse) Get() *SmtpAlertHandlerResponse {
	return v.value
}

func (v *NullableSmtpAlertHandlerResponse) Set(val *SmtpAlertHandlerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSmtpAlertHandlerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSmtpAlertHandlerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmtpAlertHandlerResponse(val *SmtpAlertHandlerResponse) *NullableSmtpAlertHandlerResponse {
	return &NullableSmtpAlertHandlerResponse{value: val, isSet: true}
}

func (v NullableSmtpAlertHandlerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmtpAlertHandlerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
