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

// checks if the AddSyslogTextErrorLogPublisherRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddSyslogTextErrorLogPublisherRequest{}

// AddSyslogTextErrorLogPublisherRequest struct for AddSyslogTextErrorLogPublisherRequest
type AddSyslogTextErrorLogPublisherRequest struct {
	// Name of the new Log Publisher
	PublisherName   string                                     `json:"publisherName"`
	Schemas         []EnumsyslogTextErrorLogPublisherSchemaUrn `json:"schemas"`
	DefaultSeverity []EnumlogPublisherDefaultSeverityProp      `json:"defaultSeverity,omitempty"`
	// The syslog server to which messages should be sent.
	SyslogExternalServer []string                            `json:"syslogExternalServer"`
	SyslogFacility       *EnumlogPublisherSyslogFacilityProp `json:"syslogFacility,omitempty"`
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
	Enabled              bool                                      `json:"enabled"`
	LoggingErrorBehavior *EnumlogPublisherLoggingErrorBehaviorProp `json:"loggingErrorBehavior,omitempty"`
}

// NewAddSyslogTextErrorLogPublisherRequest instantiates a new AddSyslogTextErrorLogPublisherRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSyslogTextErrorLogPublisherRequest(publisherName string, schemas []EnumsyslogTextErrorLogPublisherSchemaUrn, syslogExternalServer []string, enabled bool) *AddSyslogTextErrorLogPublisherRequest {
	this := AddSyslogTextErrorLogPublisherRequest{}
	this.PublisherName = publisherName
	this.Schemas = schemas
	this.SyslogExternalServer = syslogExternalServer
	this.Enabled = enabled
	return &this
}

// NewAddSyslogTextErrorLogPublisherRequestWithDefaults instantiates a new AddSyslogTextErrorLogPublisherRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSyslogTextErrorLogPublisherRequestWithDefaults() *AddSyslogTextErrorLogPublisherRequest {
	this := AddSyslogTextErrorLogPublisherRequest{}
	return &this
}

// GetPublisherName returns the PublisherName field value
func (o *AddSyslogTextErrorLogPublisherRequest) GetPublisherName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublisherName
}

// GetPublisherNameOk returns a tuple with the PublisherName field value
// and a boolean to check if the value has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) GetPublisherNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublisherName, true
}

// SetPublisherName sets field value
func (o *AddSyslogTextErrorLogPublisherRequest) SetPublisherName(v string) {
	o.PublisherName = v
}

// GetSchemas returns the Schemas field value
func (o *AddSyslogTextErrorLogPublisherRequest) GetSchemas() []EnumsyslogTextErrorLogPublisherSchemaUrn {
	if o == nil {
		var ret []EnumsyslogTextErrorLogPublisherSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) GetSchemasOk() ([]EnumsyslogTextErrorLogPublisherSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddSyslogTextErrorLogPublisherRequest) SetSchemas(v []EnumsyslogTextErrorLogPublisherSchemaUrn) {
	o.Schemas = v
}

// GetDefaultSeverity returns the DefaultSeverity field value if set, zero value otherwise.
func (o *AddSyslogTextErrorLogPublisherRequest) GetDefaultSeverity() []EnumlogPublisherDefaultSeverityProp {
	if o == nil || IsNil(o.DefaultSeverity) {
		var ret []EnumlogPublisherDefaultSeverityProp
		return ret
	}
	return o.DefaultSeverity
}

// GetDefaultSeverityOk returns a tuple with the DefaultSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) GetDefaultSeverityOk() ([]EnumlogPublisherDefaultSeverityProp, bool) {
	if o == nil || IsNil(o.DefaultSeverity) {
		return nil, false
	}
	return o.DefaultSeverity, true
}

// HasDefaultSeverity returns a boolean if a field has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) HasDefaultSeverity() bool {
	if o != nil && !IsNil(o.DefaultSeverity) {
		return true
	}

	return false
}

// SetDefaultSeverity gets a reference to the given []EnumlogPublisherDefaultSeverityProp and assigns it to the DefaultSeverity field.
func (o *AddSyslogTextErrorLogPublisherRequest) SetDefaultSeverity(v []EnumlogPublisherDefaultSeverityProp) {
	o.DefaultSeverity = v
}

// GetSyslogExternalServer returns the SyslogExternalServer field value
func (o *AddSyslogTextErrorLogPublisherRequest) GetSyslogExternalServer() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SyslogExternalServer
}

// GetSyslogExternalServerOk returns a tuple with the SyslogExternalServer field value
// and a boolean to check if the value has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) GetSyslogExternalServerOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SyslogExternalServer, true
}

// SetSyslogExternalServer sets field value
func (o *AddSyslogTextErrorLogPublisherRequest) SetSyslogExternalServer(v []string) {
	o.SyslogExternalServer = v
}

// GetSyslogFacility returns the SyslogFacility field value if set, zero value otherwise.
func (o *AddSyslogTextErrorLogPublisherRequest) GetSyslogFacility() EnumlogPublisherSyslogFacilityProp {
	if o == nil || IsNil(o.SyslogFacility) {
		var ret EnumlogPublisherSyslogFacilityProp
		return ret
	}
	return *o.SyslogFacility
}

// GetSyslogFacilityOk returns a tuple with the SyslogFacility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) GetSyslogFacilityOk() (*EnumlogPublisherSyslogFacilityProp, bool) {
	if o == nil || IsNil(o.SyslogFacility) {
		return nil, false
	}
	return o.SyslogFacility, true
}

// HasSyslogFacility returns a boolean if a field has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) HasSyslogFacility() bool {
	if o != nil && !IsNil(o.SyslogFacility) {
		return true
	}

	return false
}

// SetSyslogFacility gets a reference to the given EnumlogPublisherSyslogFacilityProp and assigns it to the SyslogFacility field.
func (o *AddSyslogTextErrorLogPublisherRequest) SetSyslogFacility(v EnumlogPublisherSyslogFacilityProp) {
	o.SyslogFacility = &v
}

// GetSyslogSeverity returns the SyslogSeverity field value if set, zero value otherwise.
func (o *AddSyslogTextErrorLogPublisherRequest) GetSyslogSeverity() EnumlogPublisherSyslogSeverityProp {
	if o == nil || IsNil(o.SyslogSeverity) {
		var ret EnumlogPublisherSyslogSeverityProp
		return ret
	}
	return *o.SyslogSeverity
}

// GetSyslogSeverityOk returns a tuple with the SyslogSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) GetSyslogSeverityOk() (*EnumlogPublisherSyslogSeverityProp, bool) {
	if o == nil || IsNil(o.SyslogSeverity) {
		return nil, false
	}
	return o.SyslogSeverity, true
}

// HasSyslogSeverity returns a boolean if a field has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) HasSyslogSeverity() bool {
	if o != nil && !IsNil(o.SyslogSeverity) {
		return true
	}

	return false
}

// SetSyslogSeverity gets a reference to the given EnumlogPublisherSyslogSeverityProp and assigns it to the SyslogSeverity field.
func (o *AddSyslogTextErrorLogPublisherRequest) SetSyslogSeverity(v EnumlogPublisherSyslogSeverityProp) {
	o.SyslogSeverity = &v
}

// GetSyslogMessageHostName returns the SyslogMessageHostName field value if set, zero value otherwise.
func (o *AddSyslogTextErrorLogPublisherRequest) GetSyslogMessageHostName() string {
	if o == nil || IsNil(o.SyslogMessageHostName) {
		var ret string
		return ret
	}
	return *o.SyslogMessageHostName
}

// GetSyslogMessageHostNameOk returns a tuple with the SyslogMessageHostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) GetSyslogMessageHostNameOk() (*string, bool) {
	if o == nil || IsNil(o.SyslogMessageHostName) {
		return nil, false
	}
	return o.SyslogMessageHostName, true
}

// HasSyslogMessageHostName returns a boolean if a field has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) HasSyslogMessageHostName() bool {
	if o != nil && !IsNil(o.SyslogMessageHostName) {
		return true
	}

	return false
}

// SetSyslogMessageHostName gets a reference to the given string and assigns it to the SyslogMessageHostName field.
func (o *AddSyslogTextErrorLogPublisherRequest) SetSyslogMessageHostName(v string) {
	o.SyslogMessageHostName = &v
}

// GetSyslogMessageApplicationName returns the SyslogMessageApplicationName field value if set, zero value otherwise.
func (o *AddSyslogTextErrorLogPublisherRequest) GetSyslogMessageApplicationName() string {
	if o == nil || IsNil(o.SyslogMessageApplicationName) {
		var ret string
		return ret
	}
	return *o.SyslogMessageApplicationName
}

// GetSyslogMessageApplicationNameOk returns a tuple with the SyslogMessageApplicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) GetSyslogMessageApplicationNameOk() (*string, bool) {
	if o == nil || IsNil(o.SyslogMessageApplicationName) {
		return nil, false
	}
	return o.SyslogMessageApplicationName, true
}

// HasSyslogMessageApplicationName returns a boolean if a field has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) HasSyslogMessageApplicationName() bool {
	if o != nil && !IsNil(o.SyslogMessageApplicationName) {
		return true
	}

	return false
}

// SetSyslogMessageApplicationName gets a reference to the given string and assigns it to the SyslogMessageApplicationName field.
func (o *AddSyslogTextErrorLogPublisherRequest) SetSyslogMessageApplicationName(v string) {
	o.SyslogMessageApplicationName = &v
}

// GetIncludeProductName returns the IncludeProductName field value if set, zero value otherwise.
func (o *AddSyslogTextErrorLogPublisherRequest) GetIncludeProductName() bool {
	if o == nil || IsNil(o.IncludeProductName) {
		var ret bool
		return ret
	}
	return *o.IncludeProductName
}

// GetIncludeProductNameOk returns a tuple with the IncludeProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) GetIncludeProductNameOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeProductName) {
		return nil, false
	}
	return o.IncludeProductName, true
}

// HasIncludeProductName returns a boolean if a field has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) HasIncludeProductName() bool {
	if o != nil && !IsNil(o.IncludeProductName) {
		return true
	}

	return false
}

// SetIncludeProductName gets a reference to the given bool and assigns it to the IncludeProductName field.
func (o *AddSyslogTextErrorLogPublisherRequest) SetIncludeProductName(v bool) {
	o.IncludeProductName = &v
}

// GetIncludeInstanceName returns the IncludeInstanceName field value if set, zero value otherwise.
func (o *AddSyslogTextErrorLogPublisherRequest) GetIncludeInstanceName() bool {
	if o == nil || IsNil(o.IncludeInstanceName) {
		var ret bool
		return ret
	}
	return *o.IncludeInstanceName
}

// GetIncludeInstanceNameOk returns a tuple with the IncludeInstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) GetIncludeInstanceNameOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeInstanceName) {
		return nil, false
	}
	return o.IncludeInstanceName, true
}

// HasIncludeInstanceName returns a boolean if a field has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) HasIncludeInstanceName() bool {
	if o != nil && !IsNil(o.IncludeInstanceName) {
		return true
	}

	return false
}

// SetIncludeInstanceName gets a reference to the given bool and assigns it to the IncludeInstanceName field.
func (o *AddSyslogTextErrorLogPublisherRequest) SetIncludeInstanceName(v bool) {
	o.IncludeInstanceName = &v
}

// GetIncludeStartupID returns the IncludeStartupID field value if set, zero value otherwise.
func (o *AddSyslogTextErrorLogPublisherRequest) GetIncludeStartupID() bool {
	if o == nil || IsNil(o.IncludeStartupID) {
		var ret bool
		return ret
	}
	return *o.IncludeStartupID
}

// GetIncludeStartupIDOk returns a tuple with the IncludeStartupID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) GetIncludeStartupIDOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeStartupID) {
		return nil, false
	}
	return o.IncludeStartupID, true
}

// HasIncludeStartupID returns a boolean if a field has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) HasIncludeStartupID() bool {
	if o != nil && !IsNil(o.IncludeStartupID) {
		return true
	}

	return false
}

// SetIncludeStartupID gets a reference to the given bool and assigns it to the IncludeStartupID field.
func (o *AddSyslogTextErrorLogPublisherRequest) SetIncludeStartupID(v bool) {
	o.IncludeStartupID = &v
}

// GetIncludeThreadID returns the IncludeThreadID field value if set, zero value otherwise.
func (o *AddSyslogTextErrorLogPublisherRequest) GetIncludeThreadID() bool {
	if o == nil || IsNil(o.IncludeThreadID) {
		var ret bool
		return ret
	}
	return *o.IncludeThreadID
}

// GetIncludeThreadIDOk returns a tuple with the IncludeThreadID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) GetIncludeThreadIDOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeThreadID) {
		return nil, false
	}
	return o.IncludeThreadID, true
}

// HasIncludeThreadID returns a boolean if a field has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) HasIncludeThreadID() bool {
	if o != nil && !IsNil(o.IncludeThreadID) {
		return true
	}

	return false
}

// SetIncludeThreadID gets a reference to the given bool and assigns it to the IncludeThreadID field.
func (o *AddSyslogTextErrorLogPublisherRequest) SetIncludeThreadID(v bool) {
	o.IncludeThreadID = &v
}

// GetGenerifyMessageStringsWhenPossible returns the GenerifyMessageStringsWhenPossible field value if set, zero value otherwise.
func (o *AddSyslogTextErrorLogPublisherRequest) GetGenerifyMessageStringsWhenPossible() bool {
	if o == nil || IsNil(o.GenerifyMessageStringsWhenPossible) {
		var ret bool
		return ret
	}
	return *o.GenerifyMessageStringsWhenPossible
}

// GetGenerifyMessageStringsWhenPossibleOk returns a tuple with the GenerifyMessageStringsWhenPossible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) GetGenerifyMessageStringsWhenPossibleOk() (*bool, bool) {
	if o == nil || IsNil(o.GenerifyMessageStringsWhenPossible) {
		return nil, false
	}
	return o.GenerifyMessageStringsWhenPossible, true
}

// HasGenerifyMessageStringsWhenPossible returns a boolean if a field has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) HasGenerifyMessageStringsWhenPossible() bool {
	if o != nil && !IsNil(o.GenerifyMessageStringsWhenPossible) {
		return true
	}

	return false
}

// SetGenerifyMessageStringsWhenPossible gets a reference to the given bool and assigns it to the GenerifyMessageStringsWhenPossible field.
func (o *AddSyslogTextErrorLogPublisherRequest) SetGenerifyMessageStringsWhenPossible(v bool) {
	o.GenerifyMessageStringsWhenPossible = &v
}

// GetTimestampPrecision returns the TimestampPrecision field value if set, zero value otherwise.
func (o *AddSyslogTextErrorLogPublisherRequest) GetTimestampPrecision() EnumlogPublisherTimestampPrecisionProp {
	if o == nil || IsNil(o.TimestampPrecision) {
		var ret EnumlogPublisherTimestampPrecisionProp
		return ret
	}
	return *o.TimestampPrecision
}

// GetTimestampPrecisionOk returns a tuple with the TimestampPrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) GetTimestampPrecisionOk() (*EnumlogPublisherTimestampPrecisionProp, bool) {
	if o == nil || IsNil(o.TimestampPrecision) {
		return nil, false
	}
	return o.TimestampPrecision, true
}

// HasTimestampPrecision returns a boolean if a field has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) HasTimestampPrecision() bool {
	if o != nil && !IsNil(o.TimestampPrecision) {
		return true
	}

	return false
}

// SetTimestampPrecision gets a reference to the given EnumlogPublisherTimestampPrecisionProp and assigns it to the TimestampPrecision field.
func (o *AddSyslogTextErrorLogPublisherRequest) SetTimestampPrecision(v EnumlogPublisherTimestampPrecisionProp) {
	o.TimestampPrecision = &v
}

// GetQueueSize returns the QueueSize field value if set, zero value otherwise.
func (o *AddSyslogTextErrorLogPublisherRequest) GetQueueSize() int64 {
	if o == nil || IsNil(o.QueueSize) {
		var ret int64
		return ret
	}
	return *o.QueueSize
}

// GetQueueSizeOk returns a tuple with the QueueSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) GetQueueSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.QueueSize) {
		return nil, false
	}
	return o.QueueSize, true
}

// HasQueueSize returns a boolean if a field has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) HasQueueSize() bool {
	if o != nil && !IsNil(o.QueueSize) {
		return true
	}

	return false
}

// SetQueueSize gets a reference to the given int64 and assigns it to the QueueSize field.
func (o *AddSyslogTextErrorLogPublisherRequest) SetQueueSize(v int64) {
	o.QueueSize = &v
}

// GetOverrideSeverity returns the OverrideSeverity field value if set, zero value otherwise.
func (o *AddSyslogTextErrorLogPublisherRequest) GetOverrideSeverity() []string {
	if o == nil || IsNil(o.OverrideSeverity) {
		var ret []string
		return ret
	}
	return o.OverrideSeverity
}

// GetOverrideSeverityOk returns a tuple with the OverrideSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) GetOverrideSeverityOk() ([]string, bool) {
	if o == nil || IsNil(o.OverrideSeverity) {
		return nil, false
	}
	return o.OverrideSeverity, true
}

// HasOverrideSeverity returns a boolean if a field has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) HasOverrideSeverity() bool {
	if o != nil && !IsNil(o.OverrideSeverity) {
		return true
	}

	return false
}

// SetOverrideSeverity gets a reference to the given []string and assigns it to the OverrideSeverity field.
func (o *AddSyslogTextErrorLogPublisherRequest) SetOverrideSeverity(v []string) {
	o.OverrideSeverity = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddSyslogTextErrorLogPublisherRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddSyslogTextErrorLogPublisherRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddSyslogTextErrorLogPublisherRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddSyslogTextErrorLogPublisherRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetLoggingErrorBehavior returns the LoggingErrorBehavior field value if set, zero value otherwise.
func (o *AddSyslogTextErrorLogPublisherRequest) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp {
	if o == nil || IsNil(o.LoggingErrorBehavior) {
		var ret EnumlogPublisherLoggingErrorBehaviorProp
		return ret
	}
	return *o.LoggingErrorBehavior
}

// GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool) {
	if o == nil || IsNil(o.LoggingErrorBehavior) {
		return nil, false
	}
	return o.LoggingErrorBehavior, true
}

// HasLoggingErrorBehavior returns a boolean if a field has been set.
func (o *AddSyslogTextErrorLogPublisherRequest) HasLoggingErrorBehavior() bool {
	if o != nil && !IsNil(o.LoggingErrorBehavior) {
		return true
	}

	return false
}

// SetLoggingErrorBehavior gets a reference to the given EnumlogPublisherLoggingErrorBehaviorProp and assigns it to the LoggingErrorBehavior field.
func (o *AddSyslogTextErrorLogPublisherRequest) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp) {
	o.LoggingErrorBehavior = &v
}

func (o AddSyslogTextErrorLogPublisherRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddSyslogTextErrorLogPublisherRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["publisherName"] = o.PublisherName
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.DefaultSeverity) {
		toSerialize["defaultSeverity"] = o.DefaultSeverity
	}
	toSerialize["syslogExternalServer"] = o.SyslogExternalServer
	if !IsNil(o.SyslogFacility) {
		toSerialize["syslogFacility"] = o.SyslogFacility
	}
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
	return toSerialize, nil
}

type NullableAddSyslogTextErrorLogPublisherRequest struct {
	value *AddSyslogTextErrorLogPublisherRequest
	isSet bool
}

func (v NullableAddSyslogTextErrorLogPublisherRequest) Get() *AddSyslogTextErrorLogPublisherRequest {
	return v.value
}

func (v *NullableAddSyslogTextErrorLogPublisherRequest) Set(val *AddSyslogTextErrorLogPublisherRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSyslogTextErrorLogPublisherRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSyslogTextErrorLogPublisherRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSyslogTextErrorLogPublisherRequest(val *AddSyslogTextErrorLogPublisherRequest) *NullableAddSyslogTextErrorLogPublisherRequest {
	return &NullableAddSyslogTextErrorLogPublisherRequest{value: val, isSet: true}
}

func (v NullableAddSyslogTextErrorLogPublisherRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSyslogTextErrorLogPublisherRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
