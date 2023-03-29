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

// checks if the JdbcBasedErrorLogPublisherResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JdbcBasedErrorLogPublisherResponse{}

// JdbcBasedErrorLogPublisherResponse struct for JdbcBasedErrorLogPublisherResponse
type JdbcBasedErrorLogPublisherResponse struct {
	// Name of the Log Publisher
	Id      string                                    `json:"id"`
	Schemas []EnumjdbcBasedErrorLogPublisherSchemaUrn `json:"schemas"`
	// The JDBC-based Database Server to use for a connection.
	Server string `json:"server"`
	// The log field mapping associates loggable fields to database column names. The table name is not part of this mapping.
	LogFieldMapping string `json:"logFieldMapping"`
	// The table name to log entries to the database server.
	LogTableName string `json:"logTableName"`
	// The maximum number of log records that can be stored in the asynchronous queue.
	QueueSize *int32 `json:"queueSize,omitempty"`
	// Specifies the default severity levels for the logger.
	DefaultSeverity []EnumlogPublisherDefaultSeverityProp `json:"defaultSeverity,omitempty"`
	// Specifies the override severity levels for the logger based on the category of the messages.
	OverrideSeverity []string `json:"overrideSeverity,omitempty"`
	// A description for this Log Publisher
	Description *string `json:"description,omitempty"`
	// Indicates whether the Log Publisher is enabled for use.
	Enabled                                       bool                                               `json:"enabled"`
	LoggingErrorBehavior                          *EnumlogPublisherLoggingErrorBehaviorProp          `json:"loggingErrorBehavior,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewJdbcBasedErrorLogPublisherResponse instantiates a new JdbcBasedErrorLogPublisherResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJdbcBasedErrorLogPublisherResponse(id string, schemas []EnumjdbcBasedErrorLogPublisherSchemaUrn, server string, logFieldMapping string, logTableName string, enabled bool) *JdbcBasedErrorLogPublisherResponse {
	this := JdbcBasedErrorLogPublisherResponse{}
	this.Id = id
	this.Schemas = schemas
	this.Server = server
	this.LogFieldMapping = logFieldMapping
	this.LogTableName = logTableName
	this.Enabled = enabled
	return &this
}

// NewJdbcBasedErrorLogPublisherResponseWithDefaults instantiates a new JdbcBasedErrorLogPublisherResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJdbcBasedErrorLogPublisherResponseWithDefaults() *JdbcBasedErrorLogPublisherResponse {
	this := JdbcBasedErrorLogPublisherResponse{}
	return &this
}

// GetId returns the Id field value
func (o *JdbcBasedErrorLogPublisherResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JdbcBasedErrorLogPublisherResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JdbcBasedErrorLogPublisherResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *JdbcBasedErrorLogPublisherResponse) GetSchemas() []EnumjdbcBasedErrorLogPublisherSchemaUrn {
	if o == nil {
		var ret []EnumjdbcBasedErrorLogPublisherSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *JdbcBasedErrorLogPublisherResponse) GetSchemasOk() ([]EnumjdbcBasedErrorLogPublisherSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *JdbcBasedErrorLogPublisherResponse) SetSchemas(v []EnumjdbcBasedErrorLogPublisherSchemaUrn) {
	o.Schemas = v
}

// GetServer returns the Server field value
func (o *JdbcBasedErrorLogPublisherResponse) GetServer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Server
}

// GetServerOk returns a tuple with the Server field value
// and a boolean to check if the value has been set.
func (o *JdbcBasedErrorLogPublisherResponse) GetServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Server, true
}

// SetServer sets field value
func (o *JdbcBasedErrorLogPublisherResponse) SetServer(v string) {
	o.Server = v
}

// GetLogFieldMapping returns the LogFieldMapping field value
func (o *JdbcBasedErrorLogPublisherResponse) GetLogFieldMapping() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogFieldMapping
}

// GetLogFieldMappingOk returns a tuple with the LogFieldMapping field value
// and a boolean to check if the value has been set.
func (o *JdbcBasedErrorLogPublisherResponse) GetLogFieldMappingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogFieldMapping, true
}

// SetLogFieldMapping sets field value
func (o *JdbcBasedErrorLogPublisherResponse) SetLogFieldMapping(v string) {
	o.LogFieldMapping = v
}

// GetLogTableName returns the LogTableName field value
func (o *JdbcBasedErrorLogPublisherResponse) GetLogTableName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogTableName
}

// GetLogTableNameOk returns a tuple with the LogTableName field value
// and a boolean to check if the value has been set.
func (o *JdbcBasedErrorLogPublisherResponse) GetLogTableNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogTableName, true
}

// SetLogTableName sets field value
func (o *JdbcBasedErrorLogPublisherResponse) SetLogTableName(v string) {
	o.LogTableName = v
}

// GetQueueSize returns the QueueSize field value if set, zero value otherwise.
func (o *JdbcBasedErrorLogPublisherResponse) GetQueueSize() int32 {
	if o == nil || IsNil(o.QueueSize) {
		var ret int32
		return ret
	}
	return *o.QueueSize
}

// GetQueueSizeOk returns a tuple with the QueueSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcBasedErrorLogPublisherResponse) GetQueueSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.QueueSize) {
		return nil, false
	}
	return o.QueueSize, true
}

// HasQueueSize returns a boolean if a field has been set.
func (o *JdbcBasedErrorLogPublisherResponse) HasQueueSize() bool {
	if o != nil && !IsNil(o.QueueSize) {
		return true
	}

	return false
}

// SetQueueSize gets a reference to the given int32 and assigns it to the QueueSize field.
func (o *JdbcBasedErrorLogPublisherResponse) SetQueueSize(v int32) {
	o.QueueSize = &v
}

// GetDefaultSeverity returns the DefaultSeverity field value if set, zero value otherwise.
func (o *JdbcBasedErrorLogPublisherResponse) GetDefaultSeverity() []EnumlogPublisherDefaultSeverityProp {
	if o == nil || IsNil(o.DefaultSeverity) {
		var ret []EnumlogPublisherDefaultSeverityProp
		return ret
	}
	return o.DefaultSeverity
}

// GetDefaultSeverityOk returns a tuple with the DefaultSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcBasedErrorLogPublisherResponse) GetDefaultSeverityOk() ([]EnumlogPublisherDefaultSeverityProp, bool) {
	if o == nil || IsNil(o.DefaultSeverity) {
		return nil, false
	}
	return o.DefaultSeverity, true
}

// HasDefaultSeverity returns a boolean if a field has been set.
func (o *JdbcBasedErrorLogPublisherResponse) HasDefaultSeverity() bool {
	if o != nil && !IsNil(o.DefaultSeverity) {
		return true
	}

	return false
}

// SetDefaultSeverity gets a reference to the given []EnumlogPublisherDefaultSeverityProp and assigns it to the DefaultSeverity field.
func (o *JdbcBasedErrorLogPublisherResponse) SetDefaultSeverity(v []EnumlogPublisherDefaultSeverityProp) {
	o.DefaultSeverity = v
}

// GetOverrideSeverity returns the OverrideSeverity field value if set, zero value otherwise.
func (o *JdbcBasedErrorLogPublisherResponse) GetOverrideSeverity() []string {
	if o == nil || IsNil(o.OverrideSeverity) {
		var ret []string
		return ret
	}
	return o.OverrideSeverity
}

// GetOverrideSeverityOk returns a tuple with the OverrideSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcBasedErrorLogPublisherResponse) GetOverrideSeverityOk() ([]string, bool) {
	if o == nil || IsNil(o.OverrideSeverity) {
		return nil, false
	}
	return o.OverrideSeverity, true
}

// HasOverrideSeverity returns a boolean if a field has been set.
func (o *JdbcBasedErrorLogPublisherResponse) HasOverrideSeverity() bool {
	if o != nil && !IsNil(o.OverrideSeverity) {
		return true
	}

	return false
}

// SetOverrideSeverity gets a reference to the given []string and assigns it to the OverrideSeverity field.
func (o *JdbcBasedErrorLogPublisherResponse) SetOverrideSeverity(v []string) {
	o.OverrideSeverity = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *JdbcBasedErrorLogPublisherResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcBasedErrorLogPublisherResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *JdbcBasedErrorLogPublisherResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *JdbcBasedErrorLogPublisherResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *JdbcBasedErrorLogPublisherResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *JdbcBasedErrorLogPublisherResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *JdbcBasedErrorLogPublisherResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetLoggingErrorBehavior returns the LoggingErrorBehavior field value if set, zero value otherwise.
func (o *JdbcBasedErrorLogPublisherResponse) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp {
	if o == nil || IsNil(o.LoggingErrorBehavior) {
		var ret EnumlogPublisherLoggingErrorBehaviorProp
		return ret
	}
	return *o.LoggingErrorBehavior
}

// GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcBasedErrorLogPublisherResponse) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool) {
	if o == nil || IsNil(o.LoggingErrorBehavior) {
		return nil, false
	}
	return o.LoggingErrorBehavior, true
}

// HasLoggingErrorBehavior returns a boolean if a field has been set.
func (o *JdbcBasedErrorLogPublisherResponse) HasLoggingErrorBehavior() bool {
	if o != nil && !IsNil(o.LoggingErrorBehavior) {
		return true
	}

	return false
}

// SetLoggingErrorBehavior gets a reference to the given EnumlogPublisherLoggingErrorBehaviorProp and assigns it to the LoggingErrorBehavior field.
func (o *JdbcBasedErrorLogPublisherResponse) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp) {
	o.LoggingErrorBehavior = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *JdbcBasedErrorLogPublisherResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcBasedErrorLogPublisherResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *JdbcBasedErrorLogPublisherResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *JdbcBasedErrorLogPublisherResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *JdbcBasedErrorLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcBasedErrorLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *JdbcBasedErrorLogPublisherResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *JdbcBasedErrorLogPublisherResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o JdbcBasedErrorLogPublisherResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JdbcBasedErrorLogPublisherResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["schemas"] = o.Schemas
	toSerialize["server"] = o.Server
	toSerialize["logFieldMapping"] = o.LogFieldMapping
	toSerialize["logTableName"] = o.LogTableName
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
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableJdbcBasedErrorLogPublisherResponse struct {
	value *JdbcBasedErrorLogPublisherResponse
	isSet bool
}

func (v NullableJdbcBasedErrorLogPublisherResponse) Get() *JdbcBasedErrorLogPublisherResponse {
	return v.value
}

func (v *NullableJdbcBasedErrorLogPublisherResponse) Set(val *JdbcBasedErrorLogPublisherResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJdbcBasedErrorLogPublisherResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJdbcBasedErrorLogPublisherResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJdbcBasedErrorLogPublisherResponse(val *JdbcBasedErrorLogPublisherResponse) *NullableJdbcBasedErrorLogPublisherResponse {
	return &NullableJdbcBasedErrorLogPublisherResponse{value: val, isSet: true}
}

func (v NullableJdbcBasedErrorLogPublisherResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJdbcBasedErrorLogPublisherResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
