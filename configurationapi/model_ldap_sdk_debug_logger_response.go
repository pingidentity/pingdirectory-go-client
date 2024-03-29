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

// checks if the LdapSdkDebugLoggerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LdapSdkDebugLoggerResponse{}

// LdapSdkDebugLoggerResponse struct for LdapSdkDebugLoggerResponse
type LdapSdkDebugLoggerResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumldapSdkDebugLoggerSchemaUrn                  `json:"schemas,omitempty"`
	// A description for this LDAP SDK Debug Logger
	Description *string `json:"description,omitempty"`
	// Indicates whether this LDAP SDK Debug Logger is enabled.
	Enabled bool `json:"enabled"`
	// The path and base name of the file to use for log files generated by this LDAP SDK Debug Logger. The path may be either absolute or relative to the server root.
	LogFile    string                                `json:"logFile"`
	DebugLevel EnumldapSdkDebugLoggerDebugLevelProp  `json:"debugLevel"`
	DebugType  []EnumldapSdkDebugLoggerDebugTypeProp `json:"debugType"`
	// Indicates whether a stack trace of the thread which called the debug method should be included in debug log messages.
	IncludeStackTrace bool `json:"includeStackTrace"`
	// The UNIX permissions of the log files created by this LDAP SDK Debug Logger.
	LogFilePermissions string `json:"logFilePermissions"`
	// Specifies the interval at which to check whether the log files need to be rotated.
	TimeInterval *string `json:"timeInterval,omitempty"`
	// Specifies whether to flush the writer after every log record.
	AutoFlush *bool `json:"autoFlush,omitempty"`
	// Indicates whether the LDAP SDK Debug Logger will publish records asynchronously.
	Asynchronous bool `json:"asynchronous"`
	// The maximum number of log records that can be stored in the asynchronous queue.
	QueueSize *int64 `json:"queueSize,omitempty"`
	// Specifies the log file buffer size.
	BufferSize *string `json:"bufferSize,omitempty"`
	// Specifies whether to append to existing log files.
	Append *bool `json:"append,omitempty"`
	// The rotation policy to use for the LDAP SDK Debug Logger .
	RotationPolicy []string `json:"rotationPolicy"`
	// A listener that should be notified whenever a log file is rotated out of service.
	RotationListener []string `json:"rotationListener,omitempty"`
	// The retention policy to use for the LDAP SDK Debug Logger .
	RetentionPolicy      []string                                        `json:"retentionPolicy"`
	CompressionMechanism *EnumldapSdkDebugLoggerCompressionMechanismProp `json:"compressionMechanism,omitempty"`
	// Indicates whether the log should be cryptographically signed so that the log content cannot be altered in an undetectable manner.
	SignLog *bool `json:"signLog,omitempty"`
	// Indicates whether log files should be encrypted so that their content is not available to unauthorized users.
	EncryptLog *bool `json:"encryptLog,omitempty"`
	// Specifies the ID of the encryption settings definition that should be used to encrypt the data. If this is not provided, the server's preferred encryption settings definition will be used. The \"encryption-settings list\" command can be used to obtain a list of the encryption settings definitions available in the server.
	EncryptionSettingsDefinitionID *string                                         `json:"encryptionSettingsDefinitionID,omitempty"`
	TimestampPrecision             *EnumldapSdkDebugLoggerTimestampPrecisionProp   `json:"timestampPrecision,omitempty"`
	LoggingErrorBehavior           *EnumldapSdkDebugLoggerLoggingErrorBehaviorProp `json:"loggingErrorBehavior,omitempty"`
}

// NewLdapSdkDebugLoggerResponse instantiates a new LdapSdkDebugLoggerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapSdkDebugLoggerResponse(enabled bool, logFile string, debugLevel EnumldapSdkDebugLoggerDebugLevelProp, debugType []EnumldapSdkDebugLoggerDebugTypeProp, includeStackTrace bool, logFilePermissions string, asynchronous bool, rotationPolicy []string, retentionPolicy []string) *LdapSdkDebugLoggerResponse {
	this := LdapSdkDebugLoggerResponse{}
	this.Enabled = enabled
	this.LogFile = logFile
	this.DebugLevel = debugLevel
	this.DebugType = debugType
	this.IncludeStackTrace = includeStackTrace
	this.LogFilePermissions = logFilePermissions
	this.Asynchronous = asynchronous
	this.RotationPolicy = rotationPolicy
	this.RetentionPolicy = retentionPolicy
	return &this
}

// NewLdapSdkDebugLoggerResponseWithDefaults instantiates a new LdapSdkDebugLoggerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapSdkDebugLoggerResponseWithDefaults() *LdapSdkDebugLoggerResponse {
	this := LdapSdkDebugLoggerResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *LdapSdkDebugLoggerResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapSdkDebugLoggerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *LdapSdkDebugLoggerResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *LdapSdkDebugLoggerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *LdapSdkDebugLoggerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapSdkDebugLoggerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *LdapSdkDebugLoggerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *LdapSdkDebugLoggerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *LdapSdkDebugLoggerResponse) GetSchemas() []EnumldapSdkDebugLoggerSchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumldapSdkDebugLoggerSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapSdkDebugLoggerResponse) GetSchemasOk() ([]EnumldapSdkDebugLoggerSchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *LdapSdkDebugLoggerResponse) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumldapSdkDebugLoggerSchemaUrn and assigns it to the Schemas field.
func (o *LdapSdkDebugLoggerResponse) SetSchemas(v []EnumldapSdkDebugLoggerSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LdapSdkDebugLoggerResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapSdkDebugLoggerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LdapSdkDebugLoggerResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LdapSdkDebugLoggerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *LdapSdkDebugLoggerResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *LdapSdkDebugLoggerResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *LdapSdkDebugLoggerResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetLogFile returns the LogFile field value
func (o *LdapSdkDebugLoggerResponse) GetLogFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogFile
}

// GetLogFileOk returns a tuple with the LogFile field value
// and a boolean to check if the value has been set.
func (o *LdapSdkDebugLoggerResponse) GetLogFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogFile, true
}

// SetLogFile sets field value
func (o *LdapSdkDebugLoggerResponse) SetLogFile(v string) {
	o.LogFile = v
}

// GetDebugLevel returns the DebugLevel field value
func (o *LdapSdkDebugLoggerResponse) GetDebugLevel() EnumldapSdkDebugLoggerDebugLevelProp {
	if o == nil {
		var ret EnumldapSdkDebugLoggerDebugLevelProp
		return ret
	}

	return o.DebugLevel
}

// GetDebugLevelOk returns a tuple with the DebugLevel field value
// and a boolean to check if the value has been set.
func (o *LdapSdkDebugLoggerResponse) GetDebugLevelOk() (*EnumldapSdkDebugLoggerDebugLevelProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DebugLevel, true
}

// SetDebugLevel sets field value
func (o *LdapSdkDebugLoggerResponse) SetDebugLevel(v EnumldapSdkDebugLoggerDebugLevelProp) {
	o.DebugLevel = v
}

// GetDebugType returns the DebugType field value
func (o *LdapSdkDebugLoggerResponse) GetDebugType() []EnumldapSdkDebugLoggerDebugTypeProp {
	if o == nil {
		var ret []EnumldapSdkDebugLoggerDebugTypeProp
		return ret
	}

	return o.DebugType
}

// GetDebugTypeOk returns a tuple with the DebugType field value
// and a boolean to check if the value has been set.
func (o *LdapSdkDebugLoggerResponse) GetDebugTypeOk() ([]EnumldapSdkDebugLoggerDebugTypeProp, bool) {
	if o == nil {
		return nil, false
	}
	return o.DebugType, true
}

// SetDebugType sets field value
func (o *LdapSdkDebugLoggerResponse) SetDebugType(v []EnumldapSdkDebugLoggerDebugTypeProp) {
	o.DebugType = v
}

// GetIncludeStackTrace returns the IncludeStackTrace field value
func (o *LdapSdkDebugLoggerResponse) GetIncludeStackTrace() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IncludeStackTrace
}

// GetIncludeStackTraceOk returns a tuple with the IncludeStackTrace field value
// and a boolean to check if the value has been set.
func (o *LdapSdkDebugLoggerResponse) GetIncludeStackTraceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludeStackTrace, true
}

// SetIncludeStackTrace sets field value
func (o *LdapSdkDebugLoggerResponse) SetIncludeStackTrace(v bool) {
	o.IncludeStackTrace = v
}

// GetLogFilePermissions returns the LogFilePermissions field value
func (o *LdapSdkDebugLoggerResponse) GetLogFilePermissions() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogFilePermissions
}

// GetLogFilePermissionsOk returns a tuple with the LogFilePermissions field value
// and a boolean to check if the value has been set.
func (o *LdapSdkDebugLoggerResponse) GetLogFilePermissionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogFilePermissions, true
}

// SetLogFilePermissions sets field value
func (o *LdapSdkDebugLoggerResponse) SetLogFilePermissions(v string) {
	o.LogFilePermissions = v
}

// GetTimeInterval returns the TimeInterval field value if set, zero value otherwise.
func (o *LdapSdkDebugLoggerResponse) GetTimeInterval() string {
	if o == nil || IsNil(o.TimeInterval) {
		var ret string
		return ret
	}
	return *o.TimeInterval
}

// GetTimeIntervalOk returns a tuple with the TimeInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapSdkDebugLoggerResponse) GetTimeIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.TimeInterval) {
		return nil, false
	}
	return o.TimeInterval, true
}

// HasTimeInterval returns a boolean if a field has been set.
func (o *LdapSdkDebugLoggerResponse) HasTimeInterval() bool {
	if o != nil && !IsNil(o.TimeInterval) {
		return true
	}

	return false
}

// SetTimeInterval gets a reference to the given string and assigns it to the TimeInterval field.
func (o *LdapSdkDebugLoggerResponse) SetTimeInterval(v string) {
	o.TimeInterval = &v
}

// GetAutoFlush returns the AutoFlush field value if set, zero value otherwise.
func (o *LdapSdkDebugLoggerResponse) GetAutoFlush() bool {
	if o == nil || IsNil(o.AutoFlush) {
		var ret bool
		return ret
	}
	return *o.AutoFlush
}

// GetAutoFlushOk returns a tuple with the AutoFlush field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapSdkDebugLoggerResponse) GetAutoFlushOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoFlush) {
		return nil, false
	}
	return o.AutoFlush, true
}

// HasAutoFlush returns a boolean if a field has been set.
func (o *LdapSdkDebugLoggerResponse) HasAutoFlush() bool {
	if o != nil && !IsNil(o.AutoFlush) {
		return true
	}

	return false
}

// SetAutoFlush gets a reference to the given bool and assigns it to the AutoFlush field.
func (o *LdapSdkDebugLoggerResponse) SetAutoFlush(v bool) {
	o.AutoFlush = &v
}

// GetAsynchronous returns the Asynchronous field value
func (o *LdapSdkDebugLoggerResponse) GetAsynchronous() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Asynchronous
}

// GetAsynchronousOk returns a tuple with the Asynchronous field value
// and a boolean to check if the value has been set.
func (o *LdapSdkDebugLoggerResponse) GetAsynchronousOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asynchronous, true
}

// SetAsynchronous sets field value
func (o *LdapSdkDebugLoggerResponse) SetAsynchronous(v bool) {
	o.Asynchronous = v
}

// GetQueueSize returns the QueueSize field value if set, zero value otherwise.
func (o *LdapSdkDebugLoggerResponse) GetQueueSize() int64 {
	if o == nil || IsNil(o.QueueSize) {
		var ret int64
		return ret
	}
	return *o.QueueSize
}

// GetQueueSizeOk returns a tuple with the QueueSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapSdkDebugLoggerResponse) GetQueueSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.QueueSize) {
		return nil, false
	}
	return o.QueueSize, true
}

// HasQueueSize returns a boolean if a field has been set.
func (o *LdapSdkDebugLoggerResponse) HasQueueSize() bool {
	if o != nil && !IsNil(o.QueueSize) {
		return true
	}

	return false
}

// SetQueueSize gets a reference to the given int64 and assigns it to the QueueSize field.
func (o *LdapSdkDebugLoggerResponse) SetQueueSize(v int64) {
	o.QueueSize = &v
}

// GetBufferSize returns the BufferSize field value if set, zero value otherwise.
func (o *LdapSdkDebugLoggerResponse) GetBufferSize() string {
	if o == nil || IsNil(o.BufferSize) {
		var ret string
		return ret
	}
	return *o.BufferSize
}

// GetBufferSizeOk returns a tuple with the BufferSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapSdkDebugLoggerResponse) GetBufferSizeOk() (*string, bool) {
	if o == nil || IsNil(o.BufferSize) {
		return nil, false
	}
	return o.BufferSize, true
}

// HasBufferSize returns a boolean if a field has been set.
func (o *LdapSdkDebugLoggerResponse) HasBufferSize() bool {
	if o != nil && !IsNil(o.BufferSize) {
		return true
	}

	return false
}

// SetBufferSize gets a reference to the given string and assigns it to the BufferSize field.
func (o *LdapSdkDebugLoggerResponse) SetBufferSize(v string) {
	o.BufferSize = &v
}

// GetAppend returns the Append field value if set, zero value otherwise.
func (o *LdapSdkDebugLoggerResponse) GetAppend() bool {
	if o == nil || IsNil(o.Append) {
		var ret bool
		return ret
	}
	return *o.Append
}

// GetAppendOk returns a tuple with the Append field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapSdkDebugLoggerResponse) GetAppendOk() (*bool, bool) {
	if o == nil || IsNil(o.Append) {
		return nil, false
	}
	return o.Append, true
}

// HasAppend returns a boolean if a field has been set.
func (o *LdapSdkDebugLoggerResponse) HasAppend() bool {
	if o != nil && !IsNil(o.Append) {
		return true
	}

	return false
}

// SetAppend gets a reference to the given bool and assigns it to the Append field.
func (o *LdapSdkDebugLoggerResponse) SetAppend(v bool) {
	o.Append = &v
}

// GetRotationPolicy returns the RotationPolicy field value
func (o *LdapSdkDebugLoggerResponse) GetRotationPolicy() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RotationPolicy
}

// GetRotationPolicyOk returns a tuple with the RotationPolicy field value
// and a boolean to check if the value has been set.
func (o *LdapSdkDebugLoggerResponse) GetRotationPolicyOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RotationPolicy, true
}

// SetRotationPolicy sets field value
func (o *LdapSdkDebugLoggerResponse) SetRotationPolicy(v []string) {
	o.RotationPolicy = v
}

// GetRotationListener returns the RotationListener field value if set, zero value otherwise.
func (o *LdapSdkDebugLoggerResponse) GetRotationListener() []string {
	if o == nil || IsNil(o.RotationListener) {
		var ret []string
		return ret
	}
	return o.RotationListener
}

// GetRotationListenerOk returns a tuple with the RotationListener field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapSdkDebugLoggerResponse) GetRotationListenerOk() ([]string, bool) {
	if o == nil || IsNil(o.RotationListener) {
		return nil, false
	}
	return o.RotationListener, true
}

// HasRotationListener returns a boolean if a field has been set.
func (o *LdapSdkDebugLoggerResponse) HasRotationListener() bool {
	if o != nil && !IsNil(o.RotationListener) {
		return true
	}

	return false
}

// SetRotationListener gets a reference to the given []string and assigns it to the RotationListener field.
func (o *LdapSdkDebugLoggerResponse) SetRotationListener(v []string) {
	o.RotationListener = v
}

// GetRetentionPolicy returns the RetentionPolicy field value
func (o *LdapSdkDebugLoggerResponse) GetRetentionPolicy() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RetentionPolicy
}

// GetRetentionPolicyOk returns a tuple with the RetentionPolicy field value
// and a boolean to check if the value has been set.
func (o *LdapSdkDebugLoggerResponse) GetRetentionPolicyOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetentionPolicy, true
}

// SetRetentionPolicy sets field value
func (o *LdapSdkDebugLoggerResponse) SetRetentionPolicy(v []string) {
	o.RetentionPolicy = v
}

// GetCompressionMechanism returns the CompressionMechanism field value if set, zero value otherwise.
func (o *LdapSdkDebugLoggerResponse) GetCompressionMechanism() EnumldapSdkDebugLoggerCompressionMechanismProp {
	if o == nil || IsNil(o.CompressionMechanism) {
		var ret EnumldapSdkDebugLoggerCompressionMechanismProp
		return ret
	}
	return *o.CompressionMechanism
}

// GetCompressionMechanismOk returns a tuple with the CompressionMechanism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapSdkDebugLoggerResponse) GetCompressionMechanismOk() (*EnumldapSdkDebugLoggerCompressionMechanismProp, bool) {
	if o == nil || IsNil(o.CompressionMechanism) {
		return nil, false
	}
	return o.CompressionMechanism, true
}

// HasCompressionMechanism returns a boolean if a field has been set.
func (o *LdapSdkDebugLoggerResponse) HasCompressionMechanism() bool {
	if o != nil && !IsNil(o.CompressionMechanism) {
		return true
	}

	return false
}

// SetCompressionMechanism gets a reference to the given EnumldapSdkDebugLoggerCompressionMechanismProp and assigns it to the CompressionMechanism field.
func (o *LdapSdkDebugLoggerResponse) SetCompressionMechanism(v EnumldapSdkDebugLoggerCompressionMechanismProp) {
	o.CompressionMechanism = &v
}

// GetSignLog returns the SignLog field value if set, zero value otherwise.
func (o *LdapSdkDebugLoggerResponse) GetSignLog() bool {
	if o == nil || IsNil(o.SignLog) {
		var ret bool
		return ret
	}
	return *o.SignLog
}

// GetSignLogOk returns a tuple with the SignLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapSdkDebugLoggerResponse) GetSignLogOk() (*bool, bool) {
	if o == nil || IsNil(o.SignLog) {
		return nil, false
	}
	return o.SignLog, true
}

// HasSignLog returns a boolean if a field has been set.
func (o *LdapSdkDebugLoggerResponse) HasSignLog() bool {
	if o != nil && !IsNil(o.SignLog) {
		return true
	}

	return false
}

// SetSignLog gets a reference to the given bool and assigns it to the SignLog field.
func (o *LdapSdkDebugLoggerResponse) SetSignLog(v bool) {
	o.SignLog = &v
}

// GetEncryptLog returns the EncryptLog field value if set, zero value otherwise.
func (o *LdapSdkDebugLoggerResponse) GetEncryptLog() bool {
	if o == nil || IsNil(o.EncryptLog) {
		var ret bool
		return ret
	}
	return *o.EncryptLog
}

// GetEncryptLogOk returns a tuple with the EncryptLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapSdkDebugLoggerResponse) GetEncryptLogOk() (*bool, bool) {
	if o == nil || IsNil(o.EncryptLog) {
		return nil, false
	}
	return o.EncryptLog, true
}

// HasEncryptLog returns a boolean if a field has been set.
func (o *LdapSdkDebugLoggerResponse) HasEncryptLog() bool {
	if o != nil && !IsNil(o.EncryptLog) {
		return true
	}

	return false
}

// SetEncryptLog gets a reference to the given bool and assigns it to the EncryptLog field.
func (o *LdapSdkDebugLoggerResponse) SetEncryptLog(v bool) {
	o.EncryptLog = &v
}

// GetEncryptionSettingsDefinitionID returns the EncryptionSettingsDefinitionID field value if set, zero value otherwise.
func (o *LdapSdkDebugLoggerResponse) GetEncryptionSettingsDefinitionID() string {
	if o == nil || IsNil(o.EncryptionSettingsDefinitionID) {
		var ret string
		return ret
	}
	return *o.EncryptionSettingsDefinitionID
}

// GetEncryptionSettingsDefinitionIDOk returns a tuple with the EncryptionSettingsDefinitionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapSdkDebugLoggerResponse) GetEncryptionSettingsDefinitionIDOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionSettingsDefinitionID) {
		return nil, false
	}
	return o.EncryptionSettingsDefinitionID, true
}

// HasEncryptionSettingsDefinitionID returns a boolean if a field has been set.
func (o *LdapSdkDebugLoggerResponse) HasEncryptionSettingsDefinitionID() bool {
	if o != nil && !IsNil(o.EncryptionSettingsDefinitionID) {
		return true
	}

	return false
}

// SetEncryptionSettingsDefinitionID gets a reference to the given string and assigns it to the EncryptionSettingsDefinitionID field.
func (o *LdapSdkDebugLoggerResponse) SetEncryptionSettingsDefinitionID(v string) {
	o.EncryptionSettingsDefinitionID = &v
}

// GetTimestampPrecision returns the TimestampPrecision field value if set, zero value otherwise.
func (o *LdapSdkDebugLoggerResponse) GetTimestampPrecision() EnumldapSdkDebugLoggerTimestampPrecisionProp {
	if o == nil || IsNil(o.TimestampPrecision) {
		var ret EnumldapSdkDebugLoggerTimestampPrecisionProp
		return ret
	}
	return *o.TimestampPrecision
}

// GetTimestampPrecisionOk returns a tuple with the TimestampPrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapSdkDebugLoggerResponse) GetTimestampPrecisionOk() (*EnumldapSdkDebugLoggerTimestampPrecisionProp, bool) {
	if o == nil || IsNil(o.TimestampPrecision) {
		return nil, false
	}
	return o.TimestampPrecision, true
}

// HasTimestampPrecision returns a boolean if a field has been set.
func (o *LdapSdkDebugLoggerResponse) HasTimestampPrecision() bool {
	if o != nil && !IsNil(o.TimestampPrecision) {
		return true
	}

	return false
}

// SetTimestampPrecision gets a reference to the given EnumldapSdkDebugLoggerTimestampPrecisionProp and assigns it to the TimestampPrecision field.
func (o *LdapSdkDebugLoggerResponse) SetTimestampPrecision(v EnumldapSdkDebugLoggerTimestampPrecisionProp) {
	o.TimestampPrecision = &v
}

// GetLoggingErrorBehavior returns the LoggingErrorBehavior field value if set, zero value otherwise.
func (o *LdapSdkDebugLoggerResponse) GetLoggingErrorBehavior() EnumldapSdkDebugLoggerLoggingErrorBehaviorProp {
	if o == nil || IsNil(o.LoggingErrorBehavior) {
		var ret EnumldapSdkDebugLoggerLoggingErrorBehaviorProp
		return ret
	}
	return *o.LoggingErrorBehavior
}

// GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapSdkDebugLoggerResponse) GetLoggingErrorBehaviorOk() (*EnumldapSdkDebugLoggerLoggingErrorBehaviorProp, bool) {
	if o == nil || IsNil(o.LoggingErrorBehavior) {
		return nil, false
	}
	return o.LoggingErrorBehavior, true
}

// HasLoggingErrorBehavior returns a boolean if a field has been set.
func (o *LdapSdkDebugLoggerResponse) HasLoggingErrorBehavior() bool {
	if o != nil && !IsNil(o.LoggingErrorBehavior) {
		return true
	}

	return false
}

// SetLoggingErrorBehavior gets a reference to the given EnumldapSdkDebugLoggerLoggingErrorBehaviorProp and assigns it to the LoggingErrorBehavior field.
func (o *LdapSdkDebugLoggerResponse) SetLoggingErrorBehavior(v EnumldapSdkDebugLoggerLoggingErrorBehaviorProp) {
	o.LoggingErrorBehavior = &v
}

func (o LdapSdkDebugLoggerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LdapSdkDebugLoggerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["logFile"] = o.LogFile
	toSerialize["debugLevel"] = o.DebugLevel
	toSerialize["debugType"] = o.DebugType
	toSerialize["includeStackTrace"] = o.IncludeStackTrace
	toSerialize["logFilePermissions"] = o.LogFilePermissions
	if !IsNil(o.TimeInterval) {
		toSerialize["timeInterval"] = o.TimeInterval
	}
	if !IsNil(o.AutoFlush) {
		toSerialize["autoFlush"] = o.AutoFlush
	}
	toSerialize["asynchronous"] = o.Asynchronous
	if !IsNil(o.QueueSize) {
		toSerialize["queueSize"] = o.QueueSize
	}
	if !IsNil(o.BufferSize) {
		toSerialize["bufferSize"] = o.BufferSize
	}
	if !IsNil(o.Append) {
		toSerialize["append"] = o.Append
	}
	toSerialize["rotationPolicy"] = o.RotationPolicy
	if !IsNil(o.RotationListener) {
		toSerialize["rotationListener"] = o.RotationListener
	}
	toSerialize["retentionPolicy"] = o.RetentionPolicy
	if !IsNil(o.CompressionMechanism) {
		toSerialize["compressionMechanism"] = o.CompressionMechanism
	}
	if !IsNil(o.SignLog) {
		toSerialize["signLog"] = o.SignLog
	}
	if !IsNil(o.EncryptLog) {
		toSerialize["encryptLog"] = o.EncryptLog
	}
	if !IsNil(o.EncryptionSettingsDefinitionID) {
		toSerialize["encryptionSettingsDefinitionID"] = o.EncryptionSettingsDefinitionID
	}
	if !IsNil(o.TimestampPrecision) {
		toSerialize["timestampPrecision"] = o.TimestampPrecision
	}
	if !IsNil(o.LoggingErrorBehavior) {
		toSerialize["loggingErrorBehavior"] = o.LoggingErrorBehavior
	}
	return toSerialize, nil
}

type NullableLdapSdkDebugLoggerResponse struct {
	value *LdapSdkDebugLoggerResponse
	isSet bool
}

func (v NullableLdapSdkDebugLoggerResponse) Get() *LdapSdkDebugLoggerResponse {
	return v.value
}

func (v *NullableLdapSdkDebugLoggerResponse) Set(val *LdapSdkDebugLoggerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapSdkDebugLoggerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapSdkDebugLoggerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapSdkDebugLoggerResponse(val *LdapSdkDebugLoggerResponse) *NullableLdapSdkDebugLoggerResponse {
	return &NullableLdapSdkDebugLoggerResponse{value: val, isSet: true}
}

func (v NullableLdapSdkDebugLoggerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapSdkDebugLoggerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
