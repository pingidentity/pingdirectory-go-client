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

// checks if the SyslogBasedErrorLogPublisherResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyslogBasedErrorLogPublisherResponse{}

// SyslogBasedErrorLogPublisherResponse struct for SyslogBasedErrorLogPublisherResponse
type SyslogBasedErrorLogPublisherResponse struct {
	// Name of the Log Publisher
	Id      string                                      `json:"id"`
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
	// Specifies the default severity levels for the logger.
	DefaultSeverity []EnumlogPublisherDefaultSeverityProp `json:"defaultSeverity,omitempty"`
	// Specifies the override severity levels for the logger based on the category of the messages.
	OverrideSeverity []string `json:"overrideSeverity,omitempty"`
	// A description for this Log Publisher
	Description                                   *string                                            `json:"description,omitempty"`
	LoggingErrorBehavior                          *EnumlogPublisherLoggingErrorBehaviorProp          `json:"loggingErrorBehavior,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewSyslogBasedErrorLogPublisherResponse instantiates a new SyslogBasedErrorLogPublisherResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyslogBasedErrorLogPublisherResponse(id string, schemas []EnumsyslogBasedErrorLogPublisherSchemaUrn, enabled bool, serverHostName string, serverPort int32, syslogFacility int32, asynchronous bool) *SyslogBasedErrorLogPublisherResponse {
	this := SyslogBasedErrorLogPublisherResponse{}
	this.Id = id
	this.Schemas = schemas
	this.Enabled = enabled
	this.ServerHostName = serverHostName
	this.ServerPort = serverPort
	this.SyslogFacility = syslogFacility
	this.Asynchronous = asynchronous
	return &this
}

// NewSyslogBasedErrorLogPublisherResponseWithDefaults instantiates a new SyslogBasedErrorLogPublisherResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyslogBasedErrorLogPublisherResponseWithDefaults() *SyslogBasedErrorLogPublisherResponse {
	this := SyslogBasedErrorLogPublisherResponse{}
	return &this
}

// GetId returns the Id field value
func (o *SyslogBasedErrorLogPublisherResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SyslogBasedErrorLogPublisherResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SyslogBasedErrorLogPublisherResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *SyslogBasedErrorLogPublisherResponse) GetSchemas() []EnumsyslogBasedErrorLogPublisherSchemaUrn {
	if o == nil {
		var ret []EnumsyslogBasedErrorLogPublisherSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *SyslogBasedErrorLogPublisherResponse) GetSchemasOk() ([]EnumsyslogBasedErrorLogPublisherSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *SyslogBasedErrorLogPublisherResponse) SetSchemas(v []EnumsyslogBasedErrorLogPublisherSchemaUrn) {
	o.Schemas = v
}

// GetEnabled returns the Enabled field value
func (o *SyslogBasedErrorLogPublisherResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SyslogBasedErrorLogPublisherResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SyslogBasedErrorLogPublisherResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetServerHostName returns the ServerHostName field value
func (o *SyslogBasedErrorLogPublisherResponse) GetServerHostName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerHostName
}

// GetServerHostNameOk returns a tuple with the ServerHostName field value
// and a boolean to check if the value has been set.
func (o *SyslogBasedErrorLogPublisherResponse) GetServerHostNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerHostName, true
}

// SetServerHostName sets field value
func (o *SyslogBasedErrorLogPublisherResponse) SetServerHostName(v string) {
	o.ServerHostName = v
}

// GetServerPort returns the ServerPort field value
func (o *SyslogBasedErrorLogPublisherResponse) GetServerPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ServerPort
}

// GetServerPortOk returns a tuple with the ServerPort field value
// and a boolean to check if the value has been set.
func (o *SyslogBasedErrorLogPublisherResponse) GetServerPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerPort, true
}

// SetServerPort sets field value
func (o *SyslogBasedErrorLogPublisherResponse) SetServerPort(v int32) {
	o.ServerPort = v
}

// GetSyslogFacility returns the SyslogFacility field value
func (o *SyslogBasedErrorLogPublisherResponse) GetSyslogFacility() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SyslogFacility
}

// GetSyslogFacilityOk returns a tuple with the SyslogFacility field value
// and a boolean to check if the value has been set.
func (o *SyslogBasedErrorLogPublisherResponse) GetSyslogFacilityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SyslogFacility, true
}

// SetSyslogFacility sets field value
func (o *SyslogBasedErrorLogPublisherResponse) SetSyslogFacility(v int32) {
	o.SyslogFacility = v
}

// GetAutoFlush returns the AutoFlush field value if set, zero value otherwise.
func (o *SyslogBasedErrorLogPublisherResponse) GetAutoFlush() bool {
	if o == nil || IsNil(o.AutoFlush) {
		var ret bool
		return ret
	}
	return *o.AutoFlush
}

// GetAutoFlushOk returns a tuple with the AutoFlush field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogBasedErrorLogPublisherResponse) GetAutoFlushOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoFlush) {
		return nil, false
	}
	return o.AutoFlush, true
}

// HasAutoFlush returns a boolean if a field has been set.
func (o *SyslogBasedErrorLogPublisherResponse) HasAutoFlush() bool {
	if o != nil && !IsNil(o.AutoFlush) {
		return true
	}

	return false
}

// SetAutoFlush gets a reference to the given bool and assigns it to the AutoFlush field.
func (o *SyslogBasedErrorLogPublisherResponse) SetAutoFlush(v bool) {
	o.AutoFlush = &v
}

// GetAsynchronous returns the Asynchronous field value
func (o *SyslogBasedErrorLogPublisherResponse) GetAsynchronous() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Asynchronous
}

// GetAsynchronousOk returns a tuple with the Asynchronous field value
// and a boolean to check if the value has been set.
func (o *SyslogBasedErrorLogPublisherResponse) GetAsynchronousOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asynchronous, true
}

// SetAsynchronous sets field value
func (o *SyslogBasedErrorLogPublisherResponse) SetAsynchronous(v bool) {
	o.Asynchronous = v
}

// GetQueueSize returns the QueueSize field value if set, zero value otherwise.
func (o *SyslogBasedErrorLogPublisherResponse) GetQueueSize() int32 {
	if o == nil || IsNil(o.QueueSize) {
		var ret int32
		return ret
	}
	return *o.QueueSize
}

// GetQueueSizeOk returns a tuple with the QueueSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogBasedErrorLogPublisherResponse) GetQueueSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.QueueSize) {
		return nil, false
	}
	return o.QueueSize, true
}

// HasQueueSize returns a boolean if a field has been set.
func (o *SyslogBasedErrorLogPublisherResponse) HasQueueSize() bool {
	if o != nil && !IsNil(o.QueueSize) {
		return true
	}

	return false
}

// SetQueueSize gets a reference to the given int32 and assigns it to the QueueSize field.
func (o *SyslogBasedErrorLogPublisherResponse) SetQueueSize(v int32) {
	o.QueueSize = &v
}

// GetDefaultSeverity returns the DefaultSeverity field value if set, zero value otherwise.
func (o *SyslogBasedErrorLogPublisherResponse) GetDefaultSeverity() []EnumlogPublisherDefaultSeverityProp {
	if o == nil || IsNil(o.DefaultSeverity) {
		var ret []EnumlogPublisherDefaultSeverityProp
		return ret
	}
	return o.DefaultSeverity
}

// GetDefaultSeverityOk returns a tuple with the DefaultSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogBasedErrorLogPublisherResponse) GetDefaultSeverityOk() ([]EnumlogPublisherDefaultSeverityProp, bool) {
	if o == nil || IsNil(o.DefaultSeverity) {
		return nil, false
	}
	return o.DefaultSeverity, true
}

// HasDefaultSeverity returns a boolean if a field has been set.
func (o *SyslogBasedErrorLogPublisherResponse) HasDefaultSeverity() bool {
	if o != nil && !IsNil(o.DefaultSeverity) {
		return true
	}

	return false
}

// SetDefaultSeverity gets a reference to the given []EnumlogPublisherDefaultSeverityProp and assigns it to the DefaultSeverity field.
func (o *SyslogBasedErrorLogPublisherResponse) SetDefaultSeverity(v []EnumlogPublisherDefaultSeverityProp) {
	o.DefaultSeverity = v
}

// GetOverrideSeverity returns the OverrideSeverity field value if set, zero value otherwise.
func (o *SyslogBasedErrorLogPublisherResponse) GetOverrideSeverity() []string {
	if o == nil || IsNil(o.OverrideSeverity) {
		var ret []string
		return ret
	}
	return o.OverrideSeverity
}

// GetOverrideSeverityOk returns a tuple with the OverrideSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogBasedErrorLogPublisherResponse) GetOverrideSeverityOk() ([]string, bool) {
	if o == nil || IsNil(o.OverrideSeverity) {
		return nil, false
	}
	return o.OverrideSeverity, true
}

// HasOverrideSeverity returns a boolean if a field has been set.
func (o *SyslogBasedErrorLogPublisherResponse) HasOverrideSeverity() bool {
	if o != nil && !IsNil(o.OverrideSeverity) {
		return true
	}

	return false
}

// SetOverrideSeverity gets a reference to the given []string and assigns it to the OverrideSeverity field.
func (o *SyslogBasedErrorLogPublisherResponse) SetOverrideSeverity(v []string) {
	o.OverrideSeverity = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SyslogBasedErrorLogPublisherResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogBasedErrorLogPublisherResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SyslogBasedErrorLogPublisherResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SyslogBasedErrorLogPublisherResponse) SetDescription(v string) {
	o.Description = &v
}

// GetLoggingErrorBehavior returns the LoggingErrorBehavior field value if set, zero value otherwise.
func (o *SyslogBasedErrorLogPublisherResponse) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp {
	if o == nil || IsNil(o.LoggingErrorBehavior) {
		var ret EnumlogPublisherLoggingErrorBehaviorProp
		return ret
	}
	return *o.LoggingErrorBehavior
}

// GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogBasedErrorLogPublisherResponse) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool) {
	if o == nil || IsNil(o.LoggingErrorBehavior) {
		return nil, false
	}
	return o.LoggingErrorBehavior, true
}

// HasLoggingErrorBehavior returns a boolean if a field has been set.
func (o *SyslogBasedErrorLogPublisherResponse) HasLoggingErrorBehavior() bool {
	if o != nil && !IsNil(o.LoggingErrorBehavior) {
		return true
	}

	return false
}

// SetLoggingErrorBehavior gets a reference to the given EnumlogPublisherLoggingErrorBehaviorProp and assigns it to the LoggingErrorBehavior field.
func (o *SyslogBasedErrorLogPublisherResponse) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp) {
	o.LoggingErrorBehavior = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SyslogBasedErrorLogPublisherResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogBasedErrorLogPublisherResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SyslogBasedErrorLogPublisherResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *SyslogBasedErrorLogPublisherResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *SyslogBasedErrorLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogBasedErrorLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *SyslogBasedErrorLogPublisherResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *SyslogBasedErrorLogPublisherResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o SyslogBasedErrorLogPublisherResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyslogBasedErrorLogPublisherResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["schemas"] = o.Schemas
	toSerialize["enabled"] = o.Enabled
	toSerialize["serverHostName"] = o.ServerHostName
	toSerialize["serverPort"] = o.ServerPort
	toSerialize["syslogFacility"] = o.SyslogFacility
	if !IsNil(o.AutoFlush) {
		toSerialize["autoFlush"] = o.AutoFlush
	}
	toSerialize["asynchronous"] = o.Asynchronous
	if !IsNil(o.QueueSize) {
		toSerialize["queueSize"] = o.QueueSize
	}
	if !IsNil(o.DefaultSeverity) {
		toSerialize["defaultSeverity"] = o.DefaultSeverity
	}
	if !IsNil(o.OverrideSeverity) {
		toSerialize["overrideSeverity"] = o.OverrideSeverity
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.LoggingErrorBehavior) {
		toSerialize["loggingErrorBehavior"] = o.LoggingErrorBehavior
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableSyslogBasedErrorLogPublisherResponse struct {
	value *SyslogBasedErrorLogPublisherResponse
	isSet bool
}

func (v NullableSyslogBasedErrorLogPublisherResponse) Get() *SyslogBasedErrorLogPublisherResponse {
	return v.value
}

func (v *NullableSyslogBasedErrorLogPublisherResponse) Set(val *SyslogBasedErrorLogPublisherResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSyslogBasedErrorLogPublisherResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSyslogBasedErrorLogPublisherResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyslogBasedErrorLogPublisherResponse(val *SyslogBasedErrorLogPublisherResponse) *NullableSyslogBasedErrorLogPublisherResponse {
	return &NullableSyslogBasedErrorLogPublisherResponse{value: val, isSet: true}
}

func (v NullableSyslogBasedErrorLogPublisherResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyslogBasedErrorLogPublisherResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
