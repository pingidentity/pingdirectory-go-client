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

// AddSyslogBasedErrorLogPublisherRequest struct for AddSyslogBasedErrorLogPublisherRequest
type AddSyslogBasedErrorLogPublisherRequest struct {
	// Name of the new Log Publisher
	PublisherName string `json:"publisherName"`
	Schemas []EnumsyslogBasedErrorLogPublisherSchemaUrn `json:"schemas"`
	// Indicates whether the Syslog Based Error Log Publisher is enabled for use.
	Enabled bool `json:"enabled"`
	// Specifies the hostname or IP address of the syslogd host to log to. It is highly recommend to use localhost.
	ServerHostName string `json:"serverHostName"`
	// Specifies the port number of the syslogd host to log to.
	ServerPort int32 `json:"serverPort"`
	// Specifies the syslog facility to use for this Syslog Based Error Log Publisher
	SyslogFacility int32 `json:"syslogFacility"`
	// Specifies whether to flush the writer after every log record.
	AutoFlush *bool `json:"autoFlush,omitempty"`
	// Indicates whether the Syslog Based Error Log Publisher will publish records asynchronously.
	Asynchronous bool `json:"asynchronous"`
	// The maximum number of log records that can be stored in the asynchronous queue.
	QueueSize *int32 `json:"queueSize,omitempty"`
	DefaultSeverity []EnumlogPublisherDefaultSeverityProp `json:"defaultSeverity,omitempty"`
	OverrideSeverity []string `json:"overrideSeverity,omitempty"`
	// A description for this Log Publisher
	Description *string `json:"description,omitempty"`
	LoggingErrorBehavior *EnumlogPublisherLoggingErrorBehaviorProp `json:"loggingErrorBehavior,omitempty"`
}

// NewAddSyslogBasedErrorLogPublisherRequest instantiates a new AddSyslogBasedErrorLogPublisherRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSyslogBasedErrorLogPublisherRequest(publisherName string, schemas []EnumsyslogBasedErrorLogPublisherSchemaUrn, enabled bool, serverHostName string, serverPort int32, syslogFacility int32, asynchronous bool) *AddSyslogBasedErrorLogPublisherRequest {
	this := AddSyslogBasedErrorLogPublisherRequest{}
	this.PublisherName = publisherName
	this.Schemas = schemas
	this.Enabled = enabled
	this.ServerHostName = serverHostName
	this.ServerPort = serverPort
	this.SyslogFacility = syslogFacility
	this.Asynchronous = asynchronous
	return &this
}

// NewAddSyslogBasedErrorLogPublisherRequestWithDefaults instantiates a new AddSyslogBasedErrorLogPublisherRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSyslogBasedErrorLogPublisherRequestWithDefaults() *AddSyslogBasedErrorLogPublisherRequest {
	this := AddSyslogBasedErrorLogPublisherRequest{}
	return &this
}

// GetPublisherName returns the PublisherName field value
func (o *AddSyslogBasedErrorLogPublisherRequest) GetPublisherName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublisherName
}

// GetPublisherNameOk returns a tuple with the PublisherName field value
// and a boolean to check if the value has been set.
func (o *AddSyslogBasedErrorLogPublisherRequest) GetPublisherNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PublisherName, true
}

// SetPublisherName sets field value
func (o *AddSyslogBasedErrorLogPublisherRequest) SetPublisherName(v string) {
	o.PublisherName = v
}

// GetSchemas returns the Schemas field value
func (o *AddSyslogBasedErrorLogPublisherRequest) GetSchemas() []EnumsyslogBasedErrorLogPublisherSchemaUrn {
	if o == nil {
		var ret []EnumsyslogBasedErrorLogPublisherSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddSyslogBasedErrorLogPublisherRequest) GetSchemasOk() ([]EnumsyslogBasedErrorLogPublisherSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddSyslogBasedErrorLogPublisherRequest) SetSchemas(v []EnumsyslogBasedErrorLogPublisherSchemaUrn) {
	o.Schemas = v
}

// GetEnabled returns the Enabled field value
func (o *AddSyslogBasedErrorLogPublisherRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddSyslogBasedErrorLogPublisherRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddSyslogBasedErrorLogPublisherRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetServerHostName returns the ServerHostName field value
func (o *AddSyslogBasedErrorLogPublisherRequest) GetServerHostName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerHostName
}

// GetServerHostNameOk returns a tuple with the ServerHostName field value
// and a boolean to check if the value has been set.
func (o *AddSyslogBasedErrorLogPublisherRequest) GetServerHostNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ServerHostName, true
}

// SetServerHostName sets field value
func (o *AddSyslogBasedErrorLogPublisherRequest) SetServerHostName(v string) {
	o.ServerHostName = v
}

// GetServerPort returns the ServerPort field value
func (o *AddSyslogBasedErrorLogPublisherRequest) GetServerPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ServerPort
}

// GetServerPortOk returns a tuple with the ServerPort field value
// and a boolean to check if the value has been set.
func (o *AddSyslogBasedErrorLogPublisherRequest) GetServerPortOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ServerPort, true
}

// SetServerPort sets field value
func (o *AddSyslogBasedErrorLogPublisherRequest) SetServerPort(v int32) {
	o.ServerPort = v
}

// GetSyslogFacility returns the SyslogFacility field value
func (o *AddSyslogBasedErrorLogPublisherRequest) GetSyslogFacility() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SyslogFacility
}

// GetSyslogFacilityOk returns a tuple with the SyslogFacility field value
// and a boolean to check if the value has been set.
func (o *AddSyslogBasedErrorLogPublisherRequest) GetSyslogFacilityOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SyslogFacility, true
}

// SetSyslogFacility sets field value
func (o *AddSyslogBasedErrorLogPublisherRequest) SetSyslogFacility(v int32) {
	o.SyslogFacility = v
}

// GetAutoFlush returns the AutoFlush field value if set, zero value otherwise.
func (o *AddSyslogBasedErrorLogPublisherRequest) GetAutoFlush() bool {
	if o == nil || isNil(o.AutoFlush) {
		var ret bool
		return ret
	}
	return *o.AutoFlush
}

// GetAutoFlushOk returns a tuple with the AutoFlush field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSyslogBasedErrorLogPublisherRequest) GetAutoFlushOk() (*bool, bool) {
	if o == nil || isNil(o.AutoFlush) {
    return nil, false
	}
	return o.AutoFlush, true
}

// HasAutoFlush returns a boolean if a field has been set.
func (o *AddSyslogBasedErrorLogPublisherRequest) HasAutoFlush() bool {
	if o != nil && !isNil(o.AutoFlush) {
		return true
	}

	return false
}

// SetAutoFlush gets a reference to the given bool and assigns it to the AutoFlush field.
func (o *AddSyslogBasedErrorLogPublisherRequest) SetAutoFlush(v bool) {
	o.AutoFlush = &v
}

// GetAsynchronous returns the Asynchronous field value
func (o *AddSyslogBasedErrorLogPublisherRequest) GetAsynchronous() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Asynchronous
}

// GetAsynchronousOk returns a tuple with the Asynchronous field value
// and a boolean to check if the value has been set.
func (o *AddSyslogBasedErrorLogPublisherRequest) GetAsynchronousOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Asynchronous, true
}

// SetAsynchronous sets field value
func (o *AddSyslogBasedErrorLogPublisherRequest) SetAsynchronous(v bool) {
	o.Asynchronous = v
}

// GetQueueSize returns the QueueSize field value if set, zero value otherwise.
func (o *AddSyslogBasedErrorLogPublisherRequest) GetQueueSize() int32 {
	if o == nil || isNil(o.QueueSize) {
		var ret int32
		return ret
	}
	return *o.QueueSize
}

// GetQueueSizeOk returns a tuple with the QueueSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSyslogBasedErrorLogPublisherRequest) GetQueueSizeOk() (*int32, bool) {
	if o == nil || isNil(o.QueueSize) {
    return nil, false
	}
	return o.QueueSize, true
}

// HasQueueSize returns a boolean if a field has been set.
func (o *AddSyslogBasedErrorLogPublisherRequest) HasQueueSize() bool {
	if o != nil && !isNil(o.QueueSize) {
		return true
	}

	return false
}

// SetQueueSize gets a reference to the given int32 and assigns it to the QueueSize field.
func (o *AddSyslogBasedErrorLogPublisherRequest) SetQueueSize(v int32) {
	o.QueueSize = &v
}

// GetDefaultSeverity returns the DefaultSeverity field value if set, zero value otherwise.
func (o *AddSyslogBasedErrorLogPublisherRequest) GetDefaultSeverity() []EnumlogPublisherDefaultSeverityProp {
	if o == nil || isNil(o.DefaultSeverity) {
		var ret []EnumlogPublisherDefaultSeverityProp
		return ret
	}
	return o.DefaultSeverity
}

// GetDefaultSeverityOk returns a tuple with the DefaultSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSyslogBasedErrorLogPublisherRequest) GetDefaultSeverityOk() ([]EnumlogPublisherDefaultSeverityProp, bool) {
	if o == nil || isNil(o.DefaultSeverity) {
    return nil, false
	}
	return o.DefaultSeverity, true
}

// HasDefaultSeverity returns a boolean if a field has been set.
func (o *AddSyslogBasedErrorLogPublisherRequest) HasDefaultSeverity() bool {
	if o != nil && !isNil(o.DefaultSeverity) {
		return true
	}

	return false
}

// SetDefaultSeverity gets a reference to the given []EnumlogPublisherDefaultSeverityProp and assigns it to the DefaultSeverity field.
func (o *AddSyslogBasedErrorLogPublisherRequest) SetDefaultSeverity(v []EnumlogPublisherDefaultSeverityProp) {
	o.DefaultSeverity = v
}

// GetOverrideSeverity returns the OverrideSeverity field value if set, zero value otherwise.
func (o *AddSyslogBasedErrorLogPublisherRequest) GetOverrideSeverity() []string {
	if o == nil || isNil(o.OverrideSeverity) {
		var ret []string
		return ret
	}
	return o.OverrideSeverity
}

// GetOverrideSeverityOk returns a tuple with the OverrideSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSyslogBasedErrorLogPublisherRequest) GetOverrideSeverityOk() ([]string, bool) {
	if o == nil || isNil(o.OverrideSeverity) {
    return nil, false
	}
	return o.OverrideSeverity, true
}

// HasOverrideSeverity returns a boolean if a field has been set.
func (o *AddSyslogBasedErrorLogPublisherRequest) HasOverrideSeverity() bool {
	if o != nil && !isNil(o.OverrideSeverity) {
		return true
	}

	return false
}

// SetOverrideSeverity gets a reference to the given []string and assigns it to the OverrideSeverity field.
func (o *AddSyslogBasedErrorLogPublisherRequest) SetOverrideSeverity(v []string) {
	o.OverrideSeverity = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddSyslogBasedErrorLogPublisherRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSyslogBasedErrorLogPublisherRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddSyslogBasedErrorLogPublisherRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddSyslogBasedErrorLogPublisherRequest) SetDescription(v string) {
	o.Description = &v
}

// GetLoggingErrorBehavior returns the LoggingErrorBehavior field value if set, zero value otherwise.
func (o *AddSyslogBasedErrorLogPublisherRequest) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp {
	if o == nil || isNil(o.LoggingErrorBehavior) {
		var ret EnumlogPublisherLoggingErrorBehaviorProp
		return ret
	}
	return *o.LoggingErrorBehavior
}

// GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSyslogBasedErrorLogPublisherRequest) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool) {
	if o == nil || isNil(o.LoggingErrorBehavior) {
    return nil, false
	}
	return o.LoggingErrorBehavior, true
}

// HasLoggingErrorBehavior returns a boolean if a field has been set.
func (o *AddSyslogBasedErrorLogPublisherRequest) HasLoggingErrorBehavior() bool {
	if o != nil && !isNil(o.LoggingErrorBehavior) {
		return true
	}

	return false
}

// SetLoggingErrorBehavior gets a reference to the given EnumlogPublisherLoggingErrorBehaviorProp and assigns it to the LoggingErrorBehavior field.
func (o *AddSyslogBasedErrorLogPublisherRequest) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp) {
	o.LoggingErrorBehavior = &v
}

func (o AddSyslogBasedErrorLogPublisherRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["publisherName"] = o.PublisherName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["serverHostName"] = o.ServerHostName
	}
	if true {
		toSerialize["serverPort"] = o.ServerPort
	}
	if true {
		toSerialize["syslogFacility"] = o.SyslogFacility
	}
	if !isNil(o.AutoFlush) {
		toSerialize["autoFlush"] = o.AutoFlush
	}
	if true {
		toSerialize["asynchronous"] = o.Asynchronous
	}
	if !isNil(o.QueueSize) {
		toSerialize["queueSize"] = o.QueueSize
	}
	if !isNil(o.DefaultSeverity) {
		toSerialize["defaultSeverity"] = o.DefaultSeverity
	}
	if !isNil(o.OverrideSeverity) {
		toSerialize["overrideSeverity"] = o.OverrideSeverity
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.LoggingErrorBehavior) {
		toSerialize["loggingErrorBehavior"] = o.LoggingErrorBehavior
	}
	return json.Marshal(toSerialize)
}

type NullableAddSyslogBasedErrorLogPublisherRequest struct {
	value *AddSyslogBasedErrorLogPublisherRequest
	isSet bool
}

func (v NullableAddSyslogBasedErrorLogPublisherRequest) Get() *AddSyslogBasedErrorLogPublisherRequest {
	return v.value
}

func (v *NullableAddSyslogBasedErrorLogPublisherRequest) Set(val *AddSyslogBasedErrorLogPublisherRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSyslogBasedErrorLogPublisherRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSyslogBasedErrorLogPublisherRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSyslogBasedErrorLogPublisherRequest(val *AddSyslogBasedErrorLogPublisherRequest) *NullableAddSyslogBasedErrorLogPublisherRequest {
	return &NullableAddSyslogBasedErrorLogPublisherRequest{value: val, isSet: true}
}

func (v NullableAddSyslogBasedErrorLogPublisherRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSyslogBasedErrorLogPublisherRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


