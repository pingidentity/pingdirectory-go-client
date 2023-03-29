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

// checks if the AddJdbcBasedErrorLogPublisherRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddJdbcBasedErrorLogPublisherRequest{}

// AddJdbcBasedErrorLogPublisherRequest struct for AddJdbcBasedErrorLogPublisherRequest
type AddJdbcBasedErrorLogPublisherRequest struct {
	// Name of the new Log Publisher
	PublisherName string                                    `json:"publisherName"`
	Schemas       []EnumjdbcBasedErrorLogPublisherSchemaUrn `json:"schemas"`
	// The JDBC-based Database Server to use for a connection.
	Server string `json:"server"`
	// The log field mapping associates loggable fields to database column names. The table name is not part of this mapping.
	LogFieldMapping string `json:"logFieldMapping"`
	// The table name to log entries to the database server.
	LogTableName *string `json:"logTableName,omitempty"`
	// The maximum number of log records that can be stored in the asynchronous queue.
	QueueSize       *int32                                `json:"queueSize,omitempty"`
	DefaultSeverity []EnumlogPublisherDefaultSeverityProp `json:"defaultSeverity,omitempty"`
	// Specifies the override severity levels for the logger based on the category of the messages.
	OverrideSeverity []string `json:"overrideSeverity,omitempty"`
	// A description for this Log Publisher
	Description *string `json:"description,omitempty"`
	// Indicates whether the Log Publisher is enabled for use.
	Enabled              bool                                      `json:"enabled"`
	LoggingErrorBehavior *EnumlogPublisherLoggingErrorBehaviorProp `json:"loggingErrorBehavior,omitempty"`
}

// NewAddJdbcBasedErrorLogPublisherRequest instantiates a new AddJdbcBasedErrorLogPublisherRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddJdbcBasedErrorLogPublisherRequest(publisherName string, schemas []EnumjdbcBasedErrorLogPublisherSchemaUrn, server string, logFieldMapping string, enabled bool) *AddJdbcBasedErrorLogPublisherRequest {
	this := AddJdbcBasedErrorLogPublisherRequest{}
	this.PublisherName = publisherName
	this.Schemas = schemas
	this.Server = server
	this.LogFieldMapping = logFieldMapping
	this.Enabled = enabled
	return &this
}

// NewAddJdbcBasedErrorLogPublisherRequestWithDefaults instantiates a new AddJdbcBasedErrorLogPublisherRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddJdbcBasedErrorLogPublisherRequestWithDefaults() *AddJdbcBasedErrorLogPublisherRequest {
	this := AddJdbcBasedErrorLogPublisherRequest{}
	return &this
}

// GetPublisherName returns the PublisherName field value
func (o *AddJdbcBasedErrorLogPublisherRequest) GetPublisherName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublisherName
}

// GetPublisherNameOk returns a tuple with the PublisherName field value
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedErrorLogPublisherRequest) GetPublisherNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublisherName, true
}

// SetPublisherName sets field value
func (o *AddJdbcBasedErrorLogPublisherRequest) SetPublisherName(v string) {
	o.PublisherName = v
}

// GetSchemas returns the Schemas field value
func (o *AddJdbcBasedErrorLogPublisherRequest) GetSchemas() []EnumjdbcBasedErrorLogPublisherSchemaUrn {
	if o == nil {
		var ret []EnumjdbcBasedErrorLogPublisherSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedErrorLogPublisherRequest) GetSchemasOk() ([]EnumjdbcBasedErrorLogPublisherSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddJdbcBasedErrorLogPublisherRequest) SetSchemas(v []EnumjdbcBasedErrorLogPublisherSchemaUrn) {
	o.Schemas = v
}

// GetServer returns the Server field value
func (o *AddJdbcBasedErrorLogPublisherRequest) GetServer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Server
}

// GetServerOk returns a tuple with the Server field value
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedErrorLogPublisherRequest) GetServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Server, true
}

// SetServer sets field value
func (o *AddJdbcBasedErrorLogPublisherRequest) SetServer(v string) {
	o.Server = v
}

// GetLogFieldMapping returns the LogFieldMapping field value
func (o *AddJdbcBasedErrorLogPublisherRequest) GetLogFieldMapping() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogFieldMapping
}

// GetLogFieldMappingOk returns a tuple with the LogFieldMapping field value
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedErrorLogPublisherRequest) GetLogFieldMappingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogFieldMapping, true
}

// SetLogFieldMapping sets field value
func (o *AddJdbcBasedErrorLogPublisherRequest) SetLogFieldMapping(v string) {
	o.LogFieldMapping = v
}

// GetLogTableName returns the LogTableName field value if set, zero value otherwise.
func (o *AddJdbcBasedErrorLogPublisherRequest) GetLogTableName() string {
	if o == nil || IsNil(o.LogTableName) {
		var ret string
		return ret
	}
	return *o.LogTableName
}

// GetLogTableNameOk returns a tuple with the LogTableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedErrorLogPublisherRequest) GetLogTableNameOk() (*string, bool) {
	if o == nil || IsNil(o.LogTableName) {
		return nil, false
	}
	return o.LogTableName, true
}

// HasLogTableName returns a boolean if a field has been set.
func (o *AddJdbcBasedErrorLogPublisherRequest) HasLogTableName() bool {
	if o != nil && !IsNil(o.LogTableName) {
		return true
	}

	return false
}

// SetLogTableName gets a reference to the given string and assigns it to the LogTableName field.
func (o *AddJdbcBasedErrorLogPublisherRequest) SetLogTableName(v string) {
	o.LogTableName = &v
}

// GetQueueSize returns the QueueSize field value if set, zero value otherwise.
func (o *AddJdbcBasedErrorLogPublisherRequest) GetQueueSize() int32 {
	if o == nil || IsNil(o.QueueSize) {
		var ret int32
		return ret
	}
	return *o.QueueSize
}

// GetQueueSizeOk returns a tuple with the QueueSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedErrorLogPublisherRequest) GetQueueSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.QueueSize) {
		return nil, false
	}
	return o.QueueSize, true
}

// HasQueueSize returns a boolean if a field has been set.
func (o *AddJdbcBasedErrorLogPublisherRequest) HasQueueSize() bool {
	if o != nil && !IsNil(o.QueueSize) {
		return true
	}

	return false
}

// SetQueueSize gets a reference to the given int32 and assigns it to the QueueSize field.
func (o *AddJdbcBasedErrorLogPublisherRequest) SetQueueSize(v int32) {
	o.QueueSize = &v
}

// GetDefaultSeverity returns the DefaultSeverity field value if set, zero value otherwise.
func (o *AddJdbcBasedErrorLogPublisherRequest) GetDefaultSeverity() []EnumlogPublisherDefaultSeverityProp {
	if o == nil || IsNil(o.DefaultSeverity) {
		var ret []EnumlogPublisherDefaultSeverityProp
		return ret
	}
	return o.DefaultSeverity
}

// GetDefaultSeverityOk returns a tuple with the DefaultSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedErrorLogPublisherRequest) GetDefaultSeverityOk() ([]EnumlogPublisherDefaultSeverityProp, bool) {
	if o == nil || IsNil(o.DefaultSeverity) {
		return nil, false
	}
	return o.DefaultSeverity, true
}

// HasDefaultSeverity returns a boolean if a field has been set.
func (o *AddJdbcBasedErrorLogPublisherRequest) HasDefaultSeverity() bool {
	if o != nil && !IsNil(o.DefaultSeverity) {
		return true
	}

	return false
}

// SetDefaultSeverity gets a reference to the given []EnumlogPublisherDefaultSeverityProp and assigns it to the DefaultSeverity field.
func (o *AddJdbcBasedErrorLogPublisherRequest) SetDefaultSeverity(v []EnumlogPublisherDefaultSeverityProp) {
	o.DefaultSeverity = v
}

// GetOverrideSeverity returns the OverrideSeverity field value if set, zero value otherwise.
func (o *AddJdbcBasedErrorLogPublisherRequest) GetOverrideSeverity() []string {
	if o == nil || IsNil(o.OverrideSeverity) {
		var ret []string
		return ret
	}
	return o.OverrideSeverity
}

// GetOverrideSeverityOk returns a tuple with the OverrideSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedErrorLogPublisherRequest) GetOverrideSeverityOk() ([]string, bool) {
	if o == nil || IsNil(o.OverrideSeverity) {
		return nil, false
	}
	return o.OverrideSeverity, true
}

// HasOverrideSeverity returns a boolean if a field has been set.
func (o *AddJdbcBasedErrorLogPublisherRequest) HasOverrideSeverity() bool {
	if o != nil && !IsNil(o.OverrideSeverity) {
		return true
	}

	return false
}

// SetOverrideSeverity gets a reference to the given []string and assigns it to the OverrideSeverity field.
func (o *AddJdbcBasedErrorLogPublisherRequest) SetOverrideSeverity(v []string) {
	o.OverrideSeverity = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddJdbcBasedErrorLogPublisherRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedErrorLogPublisherRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddJdbcBasedErrorLogPublisherRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddJdbcBasedErrorLogPublisherRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddJdbcBasedErrorLogPublisherRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedErrorLogPublisherRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddJdbcBasedErrorLogPublisherRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetLoggingErrorBehavior returns the LoggingErrorBehavior field value if set, zero value otherwise.
func (o *AddJdbcBasedErrorLogPublisherRequest) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp {
	if o == nil || IsNil(o.LoggingErrorBehavior) {
		var ret EnumlogPublisherLoggingErrorBehaviorProp
		return ret
	}
	return *o.LoggingErrorBehavior
}

// GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedErrorLogPublisherRequest) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool) {
	if o == nil || IsNil(o.LoggingErrorBehavior) {
		return nil, false
	}
	return o.LoggingErrorBehavior, true
}

// HasLoggingErrorBehavior returns a boolean if a field has been set.
func (o *AddJdbcBasedErrorLogPublisherRequest) HasLoggingErrorBehavior() bool {
	if o != nil && !IsNil(o.LoggingErrorBehavior) {
		return true
	}

	return false
}

// SetLoggingErrorBehavior gets a reference to the given EnumlogPublisherLoggingErrorBehaviorProp and assigns it to the LoggingErrorBehavior field.
func (o *AddJdbcBasedErrorLogPublisherRequest) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp) {
	o.LoggingErrorBehavior = &v
}

func (o AddJdbcBasedErrorLogPublisherRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddJdbcBasedErrorLogPublisherRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["publisherName"] = o.PublisherName
	toSerialize["schemas"] = o.Schemas
	toSerialize["server"] = o.Server
	toSerialize["logFieldMapping"] = o.LogFieldMapping
	if !IsNil(o.LogTableName) {
		toSerialize["logTableName"] = o.LogTableName
	}
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
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.LoggingErrorBehavior) {
		toSerialize["loggingErrorBehavior"] = o.LoggingErrorBehavior
	}
	return toSerialize, nil
}

type NullableAddJdbcBasedErrorLogPublisherRequest struct {
	value *AddJdbcBasedErrorLogPublisherRequest
	isSet bool
}

func (v NullableAddJdbcBasedErrorLogPublisherRequest) Get() *AddJdbcBasedErrorLogPublisherRequest {
	return v.value
}

func (v *NullableAddJdbcBasedErrorLogPublisherRequest) Set(val *AddJdbcBasedErrorLogPublisherRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddJdbcBasedErrorLogPublisherRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddJdbcBasedErrorLogPublisherRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddJdbcBasedErrorLogPublisherRequest(val *AddJdbcBasedErrorLogPublisherRequest) *NullableAddJdbcBasedErrorLogPublisherRequest {
	return &NullableAddJdbcBasedErrorLogPublisherRequest{value: val, isSet: true}
}

func (v NullableAddJdbcBasedErrorLogPublisherRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddJdbcBasedErrorLogPublisherRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
