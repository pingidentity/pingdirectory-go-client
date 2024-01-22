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

// checks if the SyslogTextErrorLogPublisherResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyslogTextErrorLogPublisherResponse{}

// SyslogTextErrorLogPublisherResponse struct for SyslogTextErrorLogPublisherResponse
type SyslogTextErrorLogPublisherResponse struct {
	Schemas         []EnumsyslogTextErrorLogPublisherSchemaUrn `json:"schemas"`
	DefaultSeverity []EnumlogPublisherDefaultSeverityProp      `json:"defaultSeverity,omitempty"`
	// The syslog server to which messages should be sent.
	SyslogExternalServer []string                            `json:"syslogExternalServer"`
	SyslogFacility       EnumlogPublisherSyslogFacilityProp  `json:"syslogFacility"`
	SyslogSeverity       *EnumlogPublisherSyslogSeverityProp `json:"syslogSeverity,omitempty"`
	// The local host name that will be included in syslog messages that are logged by this Syslog Text Error Log Publisher.
	SyslogMessageHostName *string `json:"syslogMessageHostName,omitempty"`
	// The application name that will be included in syslog messages that are logged by this Syslog Text Error Log Publisher.
	SyslogMessageApplicationName *string `json:"syslogMessageApplicationName,omitempty"`
	// Indicates whether log messages should include the product name for the Directory Server.
	IncludeProductName *bool `json:"includeProductName,omitempty"`
	// Indicates whether log messages should include the instance name for the Directory Server.
	IncludeInstanceName *bool `json:"includeInstanceName,omitempty"`
	// Indicates whether log messages should include the startup ID for the Directory Server, which is a value assigned to the server instance at startup and may be used to identify when the server has been restarted.
	IncludeStartupID *bool `json:"includeStartupID,omitempty"`
	// Indicates whether log messages should include the thread ID for the Directory Server in each log message. This ID can be used to correlate log messages from the same thread within a single log as well as generated by the same thread across different types of log files. More information about the thread with a specific ID can be obtained using the cn=JVM Stack Trace,cn=monitor entry.
	IncludeThreadID *bool `json:"includeThreadID,omitempty"`
	// Indicates whether to use the generified version of the log message string (which may use placeholders like %s for a string or %d for an integer), rather than the version of the message with those placeholders replaced with specific values that would normally be written to the log.
	GenerifyMessageStringsWhenPossible *bool                                   `json:"generifyMessageStringsWhenPossible,omitempty"`
	TimestampPrecision                 *EnumlogPublisherTimestampPrecisionProp `json:"timestampPrecision,omitempty"`
	// The maximum number of log records that can be stored in the asynchronous queue.
	QueueSize *int64 `json:"queueSize,omitempty"`
	// Specifies the override severity levels for the logger based on the category of the messages.
	OverrideSeverity []string `json:"overrideSeverity,omitempty"`
	// A description for this Log Publisher
	Description *string `json:"description,omitempty"`
	// Indicates whether the Log Publisher is enabled for use.
	Enabled                                       bool                                               `json:"enabled"`
	LoggingErrorBehavior                          *EnumlogPublisherLoggingErrorBehaviorProp          `json:"loggingErrorBehavior,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Log Publisher
	Id string `json:"id"`
}

// NewSyslogTextErrorLogPublisherResponse instantiates a new SyslogTextErrorLogPublisherResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyslogTextErrorLogPublisherResponse(schemas []EnumsyslogTextErrorLogPublisherSchemaUrn, syslogExternalServer []string, syslogFacility EnumlogPublisherSyslogFacilityProp, enabled bool, id string) *SyslogTextErrorLogPublisherResponse {
	this := SyslogTextErrorLogPublisherResponse{}
	this.Schemas = schemas
	this.SyslogExternalServer = syslogExternalServer
	this.SyslogFacility = syslogFacility
	this.Enabled = enabled
	this.Id = id
	return &this
}

// NewSyslogTextErrorLogPublisherResponseWithDefaults instantiates a new SyslogTextErrorLogPublisherResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyslogTextErrorLogPublisherResponseWithDefaults() *SyslogTextErrorLogPublisherResponse {
	this := SyslogTextErrorLogPublisherResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *SyslogTextErrorLogPublisherResponse) GetSchemas() []EnumsyslogTextErrorLogPublisherSchemaUrn {
	if o == nil {
		var ret []EnumsyslogTextErrorLogPublisherSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *SyslogTextErrorLogPublisherResponse) GetSchemasOk() ([]EnumsyslogTextErrorLogPublisherSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *SyslogTextErrorLogPublisherResponse) SetSchemas(v []EnumsyslogTextErrorLogPublisherSchemaUrn) {
	o.Schemas = v
}

// GetDefaultSeverity returns the DefaultSeverity field value if set, zero value otherwise.
func (o *SyslogTextErrorLogPublisherResponse) GetDefaultSeverity() []EnumlogPublisherDefaultSeverityProp {
	if o == nil || IsNil(o.DefaultSeverity) {
		var ret []EnumlogPublisherDefaultSeverityProp
		return ret
	}
	return o.DefaultSeverity
}

// GetDefaultSeverityOk returns a tuple with the DefaultSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogTextErrorLogPublisherResponse) GetDefaultSeverityOk() ([]EnumlogPublisherDefaultSeverityProp, bool) {
	if o == nil || IsNil(o.DefaultSeverity) {
		return nil, false
	}
	return o.DefaultSeverity, true
}

// HasDefaultSeverity returns a boolean if a field has been set.
func (o *SyslogTextErrorLogPublisherResponse) HasDefaultSeverity() bool {
	if o != nil && !IsNil(o.DefaultSeverity) {
		return true
	}

	return false
}

// SetDefaultSeverity gets a reference to the given []EnumlogPublisherDefaultSeverityProp and assigns it to the DefaultSeverity field.
func (o *SyslogTextErrorLogPublisherResponse) SetDefaultSeverity(v []EnumlogPublisherDefaultSeverityProp) {
	o.DefaultSeverity = v
}

// GetSyslogExternalServer returns the SyslogExternalServer field value
func (o *SyslogTextErrorLogPublisherResponse) GetSyslogExternalServer() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SyslogExternalServer
}

// GetSyslogExternalServerOk returns a tuple with the SyslogExternalServer field value
// and a boolean to check if the value has been set.
func (o *SyslogTextErrorLogPublisherResponse) GetSyslogExternalServerOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SyslogExternalServer, true
}

// SetSyslogExternalServer sets field value
func (o *SyslogTextErrorLogPublisherResponse) SetSyslogExternalServer(v []string) {
	o.SyslogExternalServer = v
}

// GetSyslogFacility returns the SyslogFacility field value
func (o *SyslogTextErrorLogPublisherResponse) GetSyslogFacility() EnumlogPublisherSyslogFacilityProp {
	if o == nil {
		var ret EnumlogPublisherSyslogFacilityProp
		return ret
	}

	return o.SyslogFacility
}

// GetSyslogFacilityOk returns a tuple with the SyslogFacility field value
// and a boolean to check if the value has been set.
func (o *SyslogTextErrorLogPublisherResponse) GetSyslogFacilityOk() (*EnumlogPublisherSyslogFacilityProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SyslogFacility, true
}

// SetSyslogFacility sets field value
func (o *SyslogTextErrorLogPublisherResponse) SetSyslogFacility(v EnumlogPublisherSyslogFacilityProp) {
	o.SyslogFacility = v
}

// GetSyslogSeverity returns the SyslogSeverity field value if set, zero value otherwise.
func (o *SyslogTextErrorLogPublisherResponse) GetSyslogSeverity() EnumlogPublisherSyslogSeverityProp {
	if o == nil || IsNil(o.SyslogSeverity) {
		var ret EnumlogPublisherSyslogSeverityProp
		return ret
	}
	return *o.SyslogSeverity
}

// GetSyslogSeverityOk returns a tuple with the SyslogSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogTextErrorLogPublisherResponse) GetSyslogSeverityOk() (*EnumlogPublisherSyslogSeverityProp, bool) {
	if o == nil || IsNil(o.SyslogSeverity) {
		return nil, false
	}
	return o.SyslogSeverity, true
}

// HasSyslogSeverity returns a boolean if a field has been set.
func (o *SyslogTextErrorLogPublisherResponse) HasSyslogSeverity() bool {
	if o != nil && !IsNil(o.SyslogSeverity) {
		return true
	}

	return false
}

// SetSyslogSeverity gets a reference to the given EnumlogPublisherSyslogSeverityProp and assigns it to the SyslogSeverity field.
func (o *SyslogTextErrorLogPublisherResponse) SetSyslogSeverity(v EnumlogPublisherSyslogSeverityProp) {
	o.SyslogSeverity = &v
}

// GetSyslogMessageHostName returns the SyslogMessageHostName field value if set, zero value otherwise.
func (o *SyslogTextErrorLogPublisherResponse) GetSyslogMessageHostName() string {
	if o == nil || IsNil(o.SyslogMessageHostName) {
		var ret string
		return ret
	}
	return *o.SyslogMessageHostName
}

// GetSyslogMessageHostNameOk returns a tuple with the SyslogMessageHostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogTextErrorLogPublisherResponse) GetSyslogMessageHostNameOk() (*string, bool) {
	if o == nil || IsNil(o.SyslogMessageHostName) {
		return nil, false
	}
	return o.SyslogMessageHostName, true
}

// HasSyslogMessageHostName returns a boolean if a field has been set.
func (o *SyslogTextErrorLogPublisherResponse) HasSyslogMessageHostName() bool {
	if o != nil && !IsNil(o.SyslogMessageHostName) {
		return true
	}

	return false
}

// SetSyslogMessageHostName gets a reference to the given string and assigns it to the SyslogMessageHostName field.
func (o *SyslogTextErrorLogPublisherResponse) SetSyslogMessageHostName(v string) {
	o.SyslogMessageHostName = &v
}

// GetSyslogMessageApplicationName returns the SyslogMessageApplicationName field value if set, zero value otherwise.
func (o *SyslogTextErrorLogPublisherResponse) GetSyslogMessageApplicationName() string {
	if o == nil || IsNil(o.SyslogMessageApplicationName) {
		var ret string
		return ret
	}
	return *o.SyslogMessageApplicationName
}

// GetSyslogMessageApplicationNameOk returns a tuple with the SyslogMessageApplicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogTextErrorLogPublisherResponse) GetSyslogMessageApplicationNameOk() (*string, bool) {
	if o == nil || IsNil(o.SyslogMessageApplicationName) {
		return nil, false
	}
	return o.SyslogMessageApplicationName, true
}

// HasSyslogMessageApplicationName returns a boolean if a field has been set.
func (o *SyslogTextErrorLogPublisherResponse) HasSyslogMessageApplicationName() bool {
	if o != nil && !IsNil(o.SyslogMessageApplicationName) {
		return true
	}

	return false
}

// SetSyslogMessageApplicationName gets a reference to the given string and assigns it to the SyslogMessageApplicationName field.
func (o *SyslogTextErrorLogPublisherResponse) SetSyslogMessageApplicationName(v string) {
	o.SyslogMessageApplicationName = &v
}

// GetIncludeProductName returns the IncludeProductName field value if set, zero value otherwise.
func (o *SyslogTextErrorLogPublisherResponse) GetIncludeProductName() bool {
	if o == nil || IsNil(o.IncludeProductName) {
		var ret bool
		return ret
	}
	return *o.IncludeProductName
}

// GetIncludeProductNameOk returns a tuple with the IncludeProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogTextErrorLogPublisherResponse) GetIncludeProductNameOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeProductName) {
		return nil, false
	}
	return o.IncludeProductName, true
}

// HasIncludeProductName returns a boolean if a field has been set.
func (o *SyslogTextErrorLogPublisherResponse) HasIncludeProductName() bool {
	if o != nil && !IsNil(o.IncludeProductName) {
		return true
	}

	return false
}

// SetIncludeProductName gets a reference to the given bool and assigns it to the IncludeProductName field.
func (o *SyslogTextErrorLogPublisherResponse) SetIncludeProductName(v bool) {
	o.IncludeProductName = &v
}

// GetIncludeInstanceName returns the IncludeInstanceName field value if set, zero value otherwise.
func (o *SyslogTextErrorLogPublisherResponse) GetIncludeInstanceName() bool {
	if o == nil || IsNil(o.IncludeInstanceName) {
		var ret bool
		return ret
	}
	return *o.IncludeInstanceName
}

// GetIncludeInstanceNameOk returns a tuple with the IncludeInstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogTextErrorLogPublisherResponse) GetIncludeInstanceNameOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeInstanceName) {
		return nil, false
	}
	return o.IncludeInstanceName, true
}

// HasIncludeInstanceName returns a boolean if a field has been set.
func (o *SyslogTextErrorLogPublisherResponse) HasIncludeInstanceName() bool {
	if o != nil && !IsNil(o.IncludeInstanceName) {
		return true
	}

	return false
}

// SetIncludeInstanceName gets a reference to the given bool and assigns it to the IncludeInstanceName field.
func (o *SyslogTextErrorLogPublisherResponse) SetIncludeInstanceName(v bool) {
	o.IncludeInstanceName = &v
}

// GetIncludeStartupID returns the IncludeStartupID field value if set, zero value otherwise.
func (o *SyslogTextErrorLogPublisherResponse) GetIncludeStartupID() bool {
	if o == nil || IsNil(o.IncludeStartupID) {
		var ret bool
		return ret
	}
	return *o.IncludeStartupID
}

// GetIncludeStartupIDOk returns a tuple with the IncludeStartupID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogTextErrorLogPublisherResponse) GetIncludeStartupIDOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeStartupID) {
		return nil, false
	}
	return o.IncludeStartupID, true
}

// HasIncludeStartupID returns a boolean if a field has been set.
func (o *SyslogTextErrorLogPublisherResponse) HasIncludeStartupID() bool {
	if o != nil && !IsNil(o.IncludeStartupID) {
		return true
	}

	return false
}

// SetIncludeStartupID gets a reference to the given bool and assigns it to the IncludeStartupID field.
func (o *SyslogTextErrorLogPublisherResponse) SetIncludeStartupID(v bool) {
	o.IncludeStartupID = &v
}

// GetIncludeThreadID returns the IncludeThreadID field value if set, zero value otherwise.
func (o *SyslogTextErrorLogPublisherResponse) GetIncludeThreadID() bool {
	if o == nil || IsNil(o.IncludeThreadID) {
		var ret bool
		return ret
	}
	return *o.IncludeThreadID
}

// GetIncludeThreadIDOk returns a tuple with the IncludeThreadID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogTextErrorLogPublisherResponse) GetIncludeThreadIDOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeThreadID) {
		return nil, false
	}
	return o.IncludeThreadID, true
}

// HasIncludeThreadID returns a boolean if a field has been set.
func (o *SyslogTextErrorLogPublisherResponse) HasIncludeThreadID() bool {
	if o != nil && !IsNil(o.IncludeThreadID) {
		return true
	}

	return false
}

// SetIncludeThreadID gets a reference to the given bool and assigns it to the IncludeThreadID field.
func (o *SyslogTextErrorLogPublisherResponse) SetIncludeThreadID(v bool) {
	o.IncludeThreadID = &v
}

// GetGenerifyMessageStringsWhenPossible returns the GenerifyMessageStringsWhenPossible field value if set, zero value otherwise.
func (o *SyslogTextErrorLogPublisherResponse) GetGenerifyMessageStringsWhenPossible() bool {
	if o == nil || IsNil(o.GenerifyMessageStringsWhenPossible) {
		var ret bool
		return ret
	}
	return *o.GenerifyMessageStringsWhenPossible
}

// GetGenerifyMessageStringsWhenPossibleOk returns a tuple with the GenerifyMessageStringsWhenPossible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogTextErrorLogPublisherResponse) GetGenerifyMessageStringsWhenPossibleOk() (*bool, bool) {
	if o == nil || IsNil(o.GenerifyMessageStringsWhenPossible) {
		return nil, false
	}
	return o.GenerifyMessageStringsWhenPossible, true
}

// HasGenerifyMessageStringsWhenPossible returns a boolean if a field has been set.
func (o *SyslogTextErrorLogPublisherResponse) HasGenerifyMessageStringsWhenPossible() bool {
	if o != nil && !IsNil(o.GenerifyMessageStringsWhenPossible) {
		return true
	}

	return false
}

// SetGenerifyMessageStringsWhenPossible gets a reference to the given bool and assigns it to the GenerifyMessageStringsWhenPossible field.
func (o *SyslogTextErrorLogPublisherResponse) SetGenerifyMessageStringsWhenPossible(v bool) {
	o.GenerifyMessageStringsWhenPossible = &v
}

// GetTimestampPrecision returns the TimestampPrecision field value if set, zero value otherwise.
func (o *SyslogTextErrorLogPublisherResponse) GetTimestampPrecision() EnumlogPublisherTimestampPrecisionProp {
	if o == nil || IsNil(o.TimestampPrecision) {
		var ret EnumlogPublisherTimestampPrecisionProp
		return ret
	}
	return *o.TimestampPrecision
}

// GetTimestampPrecisionOk returns a tuple with the TimestampPrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogTextErrorLogPublisherResponse) GetTimestampPrecisionOk() (*EnumlogPublisherTimestampPrecisionProp, bool) {
	if o == nil || IsNil(o.TimestampPrecision) {
		return nil, false
	}
	return o.TimestampPrecision, true
}

// HasTimestampPrecision returns a boolean if a field has been set.
func (o *SyslogTextErrorLogPublisherResponse) HasTimestampPrecision() bool {
	if o != nil && !IsNil(o.TimestampPrecision) {
		return true
	}

	return false
}

// SetTimestampPrecision gets a reference to the given EnumlogPublisherTimestampPrecisionProp and assigns it to the TimestampPrecision field.
func (o *SyslogTextErrorLogPublisherResponse) SetTimestampPrecision(v EnumlogPublisherTimestampPrecisionProp) {
	o.TimestampPrecision = &v
}

// GetQueueSize returns the QueueSize field value if set, zero value otherwise.
func (o *SyslogTextErrorLogPublisherResponse) GetQueueSize() int64 {
	if o == nil || IsNil(o.QueueSize) {
		var ret int64
		return ret
	}
	return *o.QueueSize
}

// GetQueueSizeOk returns a tuple with the QueueSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogTextErrorLogPublisherResponse) GetQueueSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.QueueSize) {
		return nil, false
	}
	return o.QueueSize, true
}

// HasQueueSize returns a boolean if a field has been set.
func (o *SyslogTextErrorLogPublisherResponse) HasQueueSize() bool {
	if o != nil && !IsNil(o.QueueSize) {
		return true
	}

	return false
}

// SetQueueSize gets a reference to the given int64 and assigns it to the QueueSize field.
func (o *SyslogTextErrorLogPublisherResponse) SetQueueSize(v int64) {
	o.QueueSize = &v
}

// GetOverrideSeverity returns the OverrideSeverity field value if set, zero value otherwise.
func (o *SyslogTextErrorLogPublisherResponse) GetOverrideSeverity() []string {
	if o == nil || IsNil(o.OverrideSeverity) {
		var ret []string
		return ret
	}
	return o.OverrideSeverity
}

// GetOverrideSeverityOk returns a tuple with the OverrideSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogTextErrorLogPublisherResponse) GetOverrideSeverityOk() ([]string, bool) {
	if o == nil || IsNil(o.OverrideSeverity) {
		return nil, false
	}
	return o.OverrideSeverity, true
}

// HasOverrideSeverity returns a boolean if a field has been set.
func (o *SyslogTextErrorLogPublisherResponse) HasOverrideSeverity() bool {
	if o != nil && !IsNil(o.OverrideSeverity) {
		return true
	}

	return false
}

// SetOverrideSeverity gets a reference to the given []string and assigns it to the OverrideSeverity field.
func (o *SyslogTextErrorLogPublisherResponse) SetOverrideSeverity(v []string) {
	o.OverrideSeverity = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SyslogTextErrorLogPublisherResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogTextErrorLogPublisherResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SyslogTextErrorLogPublisherResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SyslogTextErrorLogPublisherResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *SyslogTextErrorLogPublisherResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SyslogTextErrorLogPublisherResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SyslogTextErrorLogPublisherResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetLoggingErrorBehavior returns the LoggingErrorBehavior field value if set, zero value otherwise.
func (o *SyslogTextErrorLogPublisherResponse) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp {
	if o == nil || IsNil(o.LoggingErrorBehavior) {
		var ret EnumlogPublisherLoggingErrorBehaviorProp
		return ret
	}
	return *o.LoggingErrorBehavior
}

// GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogTextErrorLogPublisherResponse) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool) {
	if o == nil || IsNil(o.LoggingErrorBehavior) {
		return nil, false
	}
	return o.LoggingErrorBehavior, true
}

// HasLoggingErrorBehavior returns a boolean if a field has been set.
func (o *SyslogTextErrorLogPublisherResponse) HasLoggingErrorBehavior() bool {
	if o != nil && !IsNil(o.LoggingErrorBehavior) {
		return true
	}

	return false
}

// SetLoggingErrorBehavior gets a reference to the given EnumlogPublisherLoggingErrorBehaviorProp and assigns it to the LoggingErrorBehavior field.
func (o *SyslogTextErrorLogPublisherResponse) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp) {
	o.LoggingErrorBehavior = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SyslogTextErrorLogPublisherResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogTextErrorLogPublisherResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SyslogTextErrorLogPublisherResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *SyslogTextErrorLogPublisherResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *SyslogTextErrorLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogTextErrorLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *SyslogTextErrorLogPublisherResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *SyslogTextErrorLogPublisherResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *SyslogTextErrorLogPublisherResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SyslogTextErrorLogPublisherResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SyslogTextErrorLogPublisherResponse) SetId(v string) {
	o.Id = v
}

func (o SyslogTextErrorLogPublisherResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyslogTextErrorLogPublisherResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.DefaultSeverity) {
		toSerialize["defaultSeverity"] = o.DefaultSeverity
	}
	toSerialize["syslogExternalServer"] = o.SyslogExternalServer
	toSerialize["syslogFacility"] = o.SyslogFacility
	if !IsNil(o.SyslogSeverity) {
		toSerialize["syslogSeverity"] = o.SyslogSeverity
	}
	if !IsNil(o.SyslogMessageHostName) {
		toSerialize["syslogMessageHostName"] = o.SyslogMessageHostName
	}
	if !IsNil(o.SyslogMessageApplicationName) {
		toSerialize["syslogMessageApplicationName"] = o.SyslogMessageApplicationName
	}
	if !IsNil(o.IncludeProductName) {
		toSerialize["includeProductName"] = o.IncludeProductName
	}
	if !IsNil(o.IncludeInstanceName) {
		toSerialize["includeInstanceName"] = o.IncludeInstanceName
	}
	if !IsNil(o.IncludeStartupID) {
		toSerialize["includeStartupID"] = o.IncludeStartupID
	}
	if !IsNil(o.IncludeThreadID) {
		toSerialize["includeThreadID"] = o.IncludeThreadID
	}
	if !IsNil(o.GenerifyMessageStringsWhenPossible) {
		toSerialize["generifyMessageStringsWhenPossible"] = o.GenerifyMessageStringsWhenPossible
	}
	if !IsNil(o.TimestampPrecision) {
		toSerialize["timestampPrecision"] = o.TimestampPrecision
	}
	if !IsNil(o.QueueSize) {
		toSerialize["queueSize"] = o.QueueSize
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
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableSyslogTextErrorLogPublisherResponse struct {
	value *SyslogTextErrorLogPublisherResponse
	isSet bool
}

func (v NullableSyslogTextErrorLogPublisherResponse) Get() *SyslogTextErrorLogPublisherResponse {
	return v.value
}

func (v *NullableSyslogTextErrorLogPublisherResponse) Set(val *SyslogTextErrorLogPublisherResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSyslogTextErrorLogPublisherResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSyslogTextErrorLogPublisherResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyslogTextErrorLogPublisherResponse(val *SyslogTextErrorLogPublisherResponse) *NullableSyslogTextErrorLogPublisherResponse {
	return &NullableSyslogTextErrorLogPublisherResponse{value: val, isSet: true}
}

func (v NullableSyslogTextErrorLogPublisherResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyslogTextErrorLogPublisherResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
