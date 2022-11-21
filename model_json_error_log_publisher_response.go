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

// JsonErrorLogPublisherResponse struct for JsonErrorLogPublisherResponse
type JsonErrorLogPublisherResponse struct {
	// Name of the Log Publisher
	Id string `json:"id"`
	Schemas []EnumjsonErrorLogPublisherSchemaUrn `json:"schemas"`
	// The file name to use for the log files generated by the JSON Error Log Publisher. The path to the file can be specified either as relative to the server root or as an absolute path.
	LogFile string `json:"logFile"`
	// The UNIX permissions of the log files created by this JSON Error Log Publisher.
	LogFilePermissions string `json:"logFilePermissions"`
	// The rotation policy to use for the JSON Error Log Publisher .
	RotationPolicy []string `json:"rotationPolicy"`
	// A listener that should be notified whenever a log file is rotated out of service.
	RotationListener []string `json:"rotationListener,omitempty"`
	// The retention policy to use for the JSON Error Log Publisher .
	RetentionPolicy []string `json:"retentionPolicy"`
	CompressionMechanism *EnumlogPublisherCompressionMechanismProp `json:"compressionMechanism,omitempty"`
	// Indicates whether the log should be cryptographically signed so that the log content cannot be altered in an undetectable manner.
	SignLog *bool `json:"signLog,omitempty"`
	// Indicates whether log files should be encrypted so that their content is not available to unauthorized users.
	EncryptLog *bool `json:"encryptLog,omitempty"`
	// Specifies the ID of the encryption settings definition that should be used to encrypt the data. If this is not provided, the server's preferred encryption settings definition will be used. The \"encryption-settings list\" command can be used to obtain a list of the encryption settings definitions available in the server.
	EncryptionSettingsDefinitionID *string `json:"encryptionSettingsDefinitionID,omitempty"`
	// Specifies whether to append to existing log files.
	Append *bool `json:"append,omitempty"`
	// Indicates whether the JSON Error Log Publisher will publish records asynchronously.
	Asynchronous bool `json:"asynchronous"`
	// Specifies whether to flush the writer after every log record.
	AutoFlush *bool `json:"autoFlush,omitempty"`
	// Specifies the log file buffer size.
	BufferSize *string `json:"bufferSize,omitempty"`
	// The maximum number of log records that can be stored in the asynchronous queue.
	QueueSize *int32 `json:"queueSize,omitempty"`
	// Specifies the interval at which to check whether the log files need to be rotated.
	TimeInterval *string `json:"timeInterval,omitempty"`
	// Indicates whether the JSON objects should be formatted to span multiple lines with a single element on each line. The multi-line format is potentially more user friendly (if administrators may need to look at the log files), but each message will be larger because of the additional spaces and end-of-line markers.
	WriteMultiLineMessages *bool `json:"writeMultiLineMessages,omitempty"`
	// Indicates whether log messages should include the product name for the Directory Server.
	IncludeProductName *bool `json:"includeProductName,omitempty"`
	// Indicates whether log messages should include the instance name for the Directory Server.
	IncludeInstanceName *bool `json:"includeInstanceName,omitempty"`
	// Indicates whether log messages should include the startup ID for the Directory Server, which is a value assigned to the server instance at startup and may be used to identify when the server has been restarted.
	IncludeStartupID *bool `json:"includeStartupID,omitempty"`
	// Indicates whether log messages should include the thread ID for the Directory Server in each log message. This ID can be used to correlate log messages from the same thread within a single log as well as generated by the same thread across different types of log files. More information about the thread with a specific ID can be obtained using the cn=JVM Stack Trace,cn=monitor entry.
	IncludeThreadID *bool `json:"includeThreadID,omitempty"`
	// Indicates whether to use the generified version of the log message string (which may use placeholders like %s for a string or %d for an integer), rather than the version of the message with those placeholders replaced with specific values that would normally be written to the log.
	GenerifyMessageStringsWhenPossible *bool `json:"generifyMessageStringsWhenPossible,omitempty"`
	DefaultSeverity []EnumlogPublisherDefaultSeverityProp `json:"defaultSeverity,omitempty"`
	// Specifies the override severity levels for the logger based on the category of the messages.
	OverrideSeverity []string `json:"overrideSeverity,omitempty"`
	// A description for this Log Publisher
	Description *string `json:"description,omitempty"`
	// Indicates whether the Log Publisher is enabled for use.
	Enabled bool `json:"enabled"`
	LoggingErrorBehavior *EnumlogPublisherLoggingErrorBehaviorProp `json:"loggingErrorBehavior,omitempty"`
	Meta *MetaMeta `json:"meta,omitempty"`
}

// NewJsonErrorLogPublisherResponse instantiates a new JsonErrorLogPublisherResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonErrorLogPublisherResponse(id string, schemas []EnumjsonErrorLogPublisherSchemaUrn, logFile string, logFilePermissions string, rotationPolicy []string, retentionPolicy []string, asynchronous bool, enabled bool) *JsonErrorLogPublisherResponse {
	this := JsonErrorLogPublisherResponse{}
	this.Id = id
	this.Schemas = schemas
	this.LogFile = logFile
	this.LogFilePermissions = logFilePermissions
	this.RotationPolicy = rotationPolicy
	this.RetentionPolicy = retentionPolicy
	this.Asynchronous = asynchronous
	this.Enabled = enabled
	return &this
}

// NewJsonErrorLogPublisherResponseWithDefaults instantiates a new JsonErrorLogPublisherResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonErrorLogPublisherResponseWithDefaults() *JsonErrorLogPublisherResponse {
	this := JsonErrorLogPublisherResponse{}
	return &this
}

// GetId returns the Id field value
func (o *JsonErrorLogPublisherResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JsonErrorLogPublisherResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JsonErrorLogPublisherResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *JsonErrorLogPublisherResponse) GetSchemas() []EnumjsonErrorLogPublisherSchemaUrn {
	if o == nil {
		var ret []EnumjsonErrorLogPublisherSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *JsonErrorLogPublisherResponse) GetSchemasOk() ([]EnumjsonErrorLogPublisherSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *JsonErrorLogPublisherResponse) SetSchemas(v []EnumjsonErrorLogPublisherSchemaUrn) {
	o.Schemas = v
}

// GetLogFile returns the LogFile field value
func (o *JsonErrorLogPublisherResponse) GetLogFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogFile
}

// GetLogFileOk returns a tuple with the LogFile field value
// and a boolean to check if the value has been set.
func (o *JsonErrorLogPublisherResponse) GetLogFileOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LogFile, true
}

// SetLogFile sets field value
func (o *JsonErrorLogPublisherResponse) SetLogFile(v string) {
	o.LogFile = v
}

// GetLogFilePermissions returns the LogFilePermissions field value
func (o *JsonErrorLogPublisherResponse) GetLogFilePermissions() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogFilePermissions
}

// GetLogFilePermissionsOk returns a tuple with the LogFilePermissions field value
// and a boolean to check if the value has been set.
func (o *JsonErrorLogPublisherResponse) GetLogFilePermissionsOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LogFilePermissions, true
}

// SetLogFilePermissions sets field value
func (o *JsonErrorLogPublisherResponse) SetLogFilePermissions(v string) {
	o.LogFilePermissions = v
}

// GetRotationPolicy returns the RotationPolicy field value
func (o *JsonErrorLogPublisherResponse) GetRotationPolicy() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RotationPolicy
}

// GetRotationPolicyOk returns a tuple with the RotationPolicy field value
// and a boolean to check if the value has been set.
func (o *JsonErrorLogPublisherResponse) GetRotationPolicyOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RotationPolicy, true
}

// SetRotationPolicy sets field value
func (o *JsonErrorLogPublisherResponse) SetRotationPolicy(v []string) {
	o.RotationPolicy = v
}

// GetRotationListener returns the RotationListener field value if set, zero value otherwise.
func (o *JsonErrorLogPublisherResponse) GetRotationListener() []string {
	if o == nil || isNil(o.RotationListener) {
		var ret []string
		return ret
	}
	return o.RotationListener
}

// GetRotationListenerOk returns a tuple with the RotationListener field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonErrorLogPublisherResponse) GetRotationListenerOk() ([]string, bool) {
	if o == nil || isNil(o.RotationListener) {
    return nil, false
	}
	return o.RotationListener, true
}

// HasRotationListener returns a boolean if a field has been set.
func (o *JsonErrorLogPublisherResponse) HasRotationListener() bool {
	if o != nil && !isNil(o.RotationListener) {
		return true
	}

	return false
}

// SetRotationListener gets a reference to the given []string and assigns it to the RotationListener field.
func (o *JsonErrorLogPublisherResponse) SetRotationListener(v []string) {
	o.RotationListener = v
}

// GetRetentionPolicy returns the RetentionPolicy field value
func (o *JsonErrorLogPublisherResponse) GetRetentionPolicy() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RetentionPolicy
}

// GetRetentionPolicyOk returns a tuple with the RetentionPolicy field value
// and a boolean to check if the value has been set.
func (o *JsonErrorLogPublisherResponse) GetRetentionPolicyOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RetentionPolicy, true
}

// SetRetentionPolicy sets field value
func (o *JsonErrorLogPublisherResponse) SetRetentionPolicy(v []string) {
	o.RetentionPolicy = v
}

// GetCompressionMechanism returns the CompressionMechanism field value if set, zero value otherwise.
func (o *JsonErrorLogPublisherResponse) GetCompressionMechanism() EnumlogPublisherCompressionMechanismProp {
	if o == nil || isNil(o.CompressionMechanism) {
		var ret EnumlogPublisherCompressionMechanismProp
		return ret
	}
	return *o.CompressionMechanism
}

// GetCompressionMechanismOk returns a tuple with the CompressionMechanism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonErrorLogPublisherResponse) GetCompressionMechanismOk() (*EnumlogPublisherCompressionMechanismProp, bool) {
	if o == nil || isNil(o.CompressionMechanism) {
    return nil, false
	}
	return o.CompressionMechanism, true
}

// HasCompressionMechanism returns a boolean if a field has been set.
func (o *JsonErrorLogPublisherResponse) HasCompressionMechanism() bool {
	if o != nil && !isNil(o.CompressionMechanism) {
		return true
	}

	return false
}

// SetCompressionMechanism gets a reference to the given EnumlogPublisherCompressionMechanismProp and assigns it to the CompressionMechanism field.
func (o *JsonErrorLogPublisherResponse) SetCompressionMechanism(v EnumlogPublisherCompressionMechanismProp) {
	o.CompressionMechanism = &v
}

// GetSignLog returns the SignLog field value if set, zero value otherwise.
func (o *JsonErrorLogPublisherResponse) GetSignLog() bool {
	if o == nil || isNil(o.SignLog) {
		var ret bool
		return ret
	}
	return *o.SignLog
}

// GetSignLogOk returns a tuple with the SignLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonErrorLogPublisherResponse) GetSignLogOk() (*bool, bool) {
	if o == nil || isNil(o.SignLog) {
    return nil, false
	}
	return o.SignLog, true
}

// HasSignLog returns a boolean if a field has been set.
func (o *JsonErrorLogPublisherResponse) HasSignLog() bool {
	if o != nil && !isNil(o.SignLog) {
		return true
	}

	return false
}

// SetSignLog gets a reference to the given bool and assigns it to the SignLog field.
func (o *JsonErrorLogPublisherResponse) SetSignLog(v bool) {
	o.SignLog = &v
}

// GetEncryptLog returns the EncryptLog field value if set, zero value otherwise.
func (o *JsonErrorLogPublisherResponse) GetEncryptLog() bool {
	if o == nil || isNil(o.EncryptLog) {
		var ret bool
		return ret
	}
	return *o.EncryptLog
}

// GetEncryptLogOk returns a tuple with the EncryptLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonErrorLogPublisherResponse) GetEncryptLogOk() (*bool, bool) {
	if o == nil || isNil(o.EncryptLog) {
    return nil, false
	}
	return o.EncryptLog, true
}

// HasEncryptLog returns a boolean if a field has been set.
func (o *JsonErrorLogPublisherResponse) HasEncryptLog() bool {
	if o != nil && !isNil(o.EncryptLog) {
		return true
	}

	return false
}

// SetEncryptLog gets a reference to the given bool and assigns it to the EncryptLog field.
func (o *JsonErrorLogPublisherResponse) SetEncryptLog(v bool) {
	o.EncryptLog = &v
}

// GetEncryptionSettingsDefinitionID returns the EncryptionSettingsDefinitionID field value if set, zero value otherwise.
func (o *JsonErrorLogPublisherResponse) GetEncryptionSettingsDefinitionID() string {
	if o == nil || isNil(o.EncryptionSettingsDefinitionID) {
		var ret string
		return ret
	}
	return *o.EncryptionSettingsDefinitionID
}

// GetEncryptionSettingsDefinitionIDOk returns a tuple with the EncryptionSettingsDefinitionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonErrorLogPublisherResponse) GetEncryptionSettingsDefinitionIDOk() (*string, bool) {
	if o == nil || isNil(o.EncryptionSettingsDefinitionID) {
    return nil, false
	}
	return o.EncryptionSettingsDefinitionID, true
}

// HasEncryptionSettingsDefinitionID returns a boolean if a field has been set.
func (o *JsonErrorLogPublisherResponse) HasEncryptionSettingsDefinitionID() bool {
	if o != nil && !isNil(o.EncryptionSettingsDefinitionID) {
		return true
	}

	return false
}

// SetEncryptionSettingsDefinitionID gets a reference to the given string and assigns it to the EncryptionSettingsDefinitionID field.
func (o *JsonErrorLogPublisherResponse) SetEncryptionSettingsDefinitionID(v string) {
	o.EncryptionSettingsDefinitionID = &v
}

// GetAppend returns the Append field value if set, zero value otherwise.
func (o *JsonErrorLogPublisherResponse) GetAppend() bool {
	if o == nil || isNil(o.Append) {
		var ret bool
		return ret
	}
	return *o.Append
}

// GetAppendOk returns a tuple with the Append field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonErrorLogPublisherResponse) GetAppendOk() (*bool, bool) {
	if o == nil || isNil(o.Append) {
    return nil, false
	}
	return o.Append, true
}

// HasAppend returns a boolean if a field has been set.
func (o *JsonErrorLogPublisherResponse) HasAppend() bool {
	if o != nil && !isNil(o.Append) {
		return true
	}

	return false
}

// SetAppend gets a reference to the given bool and assigns it to the Append field.
func (o *JsonErrorLogPublisherResponse) SetAppend(v bool) {
	o.Append = &v
}

// GetAsynchronous returns the Asynchronous field value
func (o *JsonErrorLogPublisherResponse) GetAsynchronous() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Asynchronous
}

// GetAsynchronousOk returns a tuple with the Asynchronous field value
// and a boolean to check if the value has been set.
func (o *JsonErrorLogPublisherResponse) GetAsynchronousOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Asynchronous, true
}

// SetAsynchronous sets field value
func (o *JsonErrorLogPublisherResponse) SetAsynchronous(v bool) {
	o.Asynchronous = v
}

// GetAutoFlush returns the AutoFlush field value if set, zero value otherwise.
func (o *JsonErrorLogPublisherResponse) GetAutoFlush() bool {
	if o == nil || isNil(o.AutoFlush) {
		var ret bool
		return ret
	}
	return *o.AutoFlush
}

// GetAutoFlushOk returns a tuple with the AutoFlush field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonErrorLogPublisherResponse) GetAutoFlushOk() (*bool, bool) {
	if o == nil || isNil(o.AutoFlush) {
    return nil, false
	}
	return o.AutoFlush, true
}

// HasAutoFlush returns a boolean if a field has been set.
func (o *JsonErrorLogPublisherResponse) HasAutoFlush() bool {
	if o != nil && !isNil(o.AutoFlush) {
		return true
	}

	return false
}

// SetAutoFlush gets a reference to the given bool and assigns it to the AutoFlush field.
func (o *JsonErrorLogPublisherResponse) SetAutoFlush(v bool) {
	o.AutoFlush = &v
}

// GetBufferSize returns the BufferSize field value if set, zero value otherwise.
func (o *JsonErrorLogPublisherResponse) GetBufferSize() string {
	if o == nil || isNil(o.BufferSize) {
		var ret string
		return ret
	}
	return *o.BufferSize
}

// GetBufferSizeOk returns a tuple with the BufferSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonErrorLogPublisherResponse) GetBufferSizeOk() (*string, bool) {
	if o == nil || isNil(o.BufferSize) {
    return nil, false
	}
	return o.BufferSize, true
}

// HasBufferSize returns a boolean if a field has been set.
func (o *JsonErrorLogPublisherResponse) HasBufferSize() bool {
	if o != nil && !isNil(o.BufferSize) {
		return true
	}

	return false
}

// SetBufferSize gets a reference to the given string and assigns it to the BufferSize field.
func (o *JsonErrorLogPublisherResponse) SetBufferSize(v string) {
	o.BufferSize = &v
}

// GetQueueSize returns the QueueSize field value if set, zero value otherwise.
func (o *JsonErrorLogPublisherResponse) GetQueueSize() int32 {
	if o == nil || isNil(o.QueueSize) {
		var ret int32
		return ret
	}
	return *o.QueueSize
}

// GetQueueSizeOk returns a tuple with the QueueSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonErrorLogPublisherResponse) GetQueueSizeOk() (*int32, bool) {
	if o == nil || isNil(o.QueueSize) {
    return nil, false
	}
	return o.QueueSize, true
}

// HasQueueSize returns a boolean if a field has been set.
func (o *JsonErrorLogPublisherResponse) HasQueueSize() bool {
	if o != nil && !isNil(o.QueueSize) {
		return true
	}

	return false
}

// SetQueueSize gets a reference to the given int32 and assigns it to the QueueSize field.
func (o *JsonErrorLogPublisherResponse) SetQueueSize(v int32) {
	o.QueueSize = &v
}

// GetTimeInterval returns the TimeInterval field value if set, zero value otherwise.
func (o *JsonErrorLogPublisherResponse) GetTimeInterval() string {
	if o == nil || isNil(o.TimeInterval) {
		var ret string
		return ret
	}
	return *o.TimeInterval
}

// GetTimeIntervalOk returns a tuple with the TimeInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonErrorLogPublisherResponse) GetTimeIntervalOk() (*string, bool) {
	if o == nil || isNil(o.TimeInterval) {
    return nil, false
	}
	return o.TimeInterval, true
}

// HasTimeInterval returns a boolean if a field has been set.
func (o *JsonErrorLogPublisherResponse) HasTimeInterval() bool {
	if o != nil && !isNil(o.TimeInterval) {
		return true
	}

	return false
}

// SetTimeInterval gets a reference to the given string and assigns it to the TimeInterval field.
func (o *JsonErrorLogPublisherResponse) SetTimeInterval(v string) {
	o.TimeInterval = &v
}

// GetWriteMultiLineMessages returns the WriteMultiLineMessages field value if set, zero value otherwise.
func (o *JsonErrorLogPublisherResponse) GetWriteMultiLineMessages() bool {
	if o == nil || isNil(o.WriteMultiLineMessages) {
		var ret bool
		return ret
	}
	return *o.WriteMultiLineMessages
}

// GetWriteMultiLineMessagesOk returns a tuple with the WriteMultiLineMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonErrorLogPublisherResponse) GetWriteMultiLineMessagesOk() (*bool, bool) {
	if o == nil || isNil(o.WriteMultiLineMessages) {
    return nil, false
	}
	return o.WriteMultiLineMessages, true
}

// HasWriteMultiLineMessages returns a boolean if a field has been set.
func (o *JsonErrorLogPublisherResponse) HasWriteMultiLineMessages() bool {
	if o != nil && !isNil(o.WriteMultiLineMessages) {
		return true
	}

	return false
}

// SetWriteMultiLineMessages gets a reference to the given bool and assigns it to the WriteMultiLineMessages field.
func (o *JsonErrorLogPublisherResponse) SetWriteMultiLineMessages(v bool) {
	o.WriteMultiLineMessages = &v
}

// GetIncludeProductName returns the IncludeProductName field value if set, zero value otherwise.
func (o *JsonErrorLogPublisherResponse) GetIncludeProductName() bool {
	if o == nil || isNil(o.IncludeProductName) {
		var ret bool
		return ret
	}
	return *o.IncludeProductName
}

// GetIncludeProductNameOk returns a tuple with the IncludeProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonErrorLogPublisherResponse) GetIncludeProductNameOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeProductName) {
    return nil, false
	}
	return o.IncludeProductName, true
}

// HasIncludeProductName returns a boolean if a field has been set.
func (o *JsonErrorLogPublisherResponse) HasIncludeProductName() bool {
	if o != nil && !isNil(o.IncludeProductName) {
		return true
	}

	return false
}

// SetIncludeProductName gets a reference to the given bool and assigns it to the IncludeProductName field.
func (o *JsonErrorLogPublisherResponse) SetIncludeProductName(v bool) {
	o.IncludeProductName = &v
}

// GetIncludeInstanceName returns the IncludeInstanceName field value if set, zero value otherwise.
func (o *JsonErrorLogPublisherResponse) GetIncludeInstanceName() bool {
	if o == nil || isNil(o.IncludeInstanceName) {
		var ret bool
		return ret
	}
	return *o.IncludeInstanceName
}

// GetIncludeInstanceNameOk returns a tuple with the IncludeInstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonErrorLogPublisherResponse) GetIncludeInstanceNameOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeInstanceName) {
    return nil, false
	}
	return o.IncludeInstanceName, true
}

// HasIncludeInstanceName returns a boolean if a field has been set.
func (o *JsonErrorLogPublisherResponse) HasIncludeInstanceName() bool {
	if o != nil && !isNil(o.IncludeInstanceName) {
		return true
	}

	return false
}

// SetIncludeInstanceName gets a reference to the given bool and assigns it to the IncludeInstanceName field.
func (o *JsonErrorLogPublisherResponse) SetIncludeInstanceName(v bool) {
	o.IncludeInstanceName = &v
}

// GetIncludeStartupID returns the IncludeStartupID field value if set, zero value otherwise.
func (o *JsonErrorLogPublisherResponse) GetIncludeStartupID() bool {
	if o == nil || isNil(o.IncludeStartupID) {
		var ret bool
		return ret
	}
	return *o.IncludeStartupID
}

// GetIncludeStartupIDOk returns a tuple with the IncludeStartupID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonErrorLogPublisherResponse) GetIncludeStartupIDOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeStartupID) {
    return nil, false
	}
	return o.IncludeStartupID, true
}

// HasIncludeStartupID returns a boolean if a field has been set.
func (o *JsonErrorLogPublisherResponse) HasIncludeStartupID() bool {
	if o != nil && !isNil(o.IncludeStartupID) {
		return true
	}

	return false
}

// SetIncludeStartupID gets a reference to the given bool and assigns it to the IncludeStartupID field.
func (o *JsonErrorLogPublisherResponse) SetIncludeStartupID(v bool) {
	o.IncludeStartupID = &v
}

// GetIncludeThreadID returns the IncludeThreadID field value if set, zero value otherwise.
func (o *JsonErrorLogPublisherResponse) GetIncludeThreadID() bool {
	if o == nil || isNil(o.IncludeThreadID) {
		var ret bool
		return ret
	}
	return *o.IncludeThreadID
}

// GetIncludeThreadIDOk returns a tuple with the IncludeThreadID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonErrorLogPublisherResponse) GetIncludeThreadIDOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeThreadID) {
    return nil, false
	}
	return o.IncludeThreadID, true
}

// HasIncludeThreadID returns a boolean if a field has been set.
func (o *JsonErrorLogPublisherResponse) HasIncludeThreadID() bool {
	if o != nil && !isNil(o.IncludeThreadID) {
		return true
	}

	return false
}

// SetIncludeThreadID gets a reference to the given bool and assigns it to the IncludeThreadID field.
func (o *JsonErrorLogPublisherResponse) SetIncludeThreadID(v bool) {
	o.IncludeThreadID = &v
}

// GetGenerifyMessageStringsWhenPossible returns the GenerifyMessageStringsWhenPossible field value if set, zero value otherwise.
func (o *JsonErrorLogPublisherResponse) GetGenerifyMessageStringsWhenPossible() bool {
	if o == nil || isNil(o.GenerifyMessageStringsWhenPossible) {
		var ret bool
		return ret
	}
	return *o.GenerifyMessageStringsWhenPossible
}

// GetGenerifyMessageStringsWhenPossibleOk returns a tuple with the GenerifyMessageStringsWhenPossible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonErrorLogPublisherResponse) GetGenerifyMessageStringsWhenPossibleOk() (*bool, bool) {
	if o == nil || isNil(o.GenerifyMessageStringsWhenPossible) {
    return nil, false
	}
	return o.GenerifyMessageStringsWhenPossible, true
}

// HasGenerifyMessageStringsWhenPossible returns a boolean if a field has been set.
func (o *JsonErrorLogPublisherResponse) HasGenerifyMessageStringsWhenPossible() bool {
	if o != nil && !isNil(o.GenerifyMessageStringsWhenPossible) {
		return true
	}

	return false
}

// SetGenerifyMessageStringsWhenPossible gets a reference to the given bool and assigns it to the GenerifyMessageStringsWhenPossible field.
func (o *JsonErrorLogPublisherResponse) SetGenerifyMessageStringsWhenPossible(v bool) {
	o.GenerifyMessageStringsWhenPossible = &v
}

// GetDefaultSeverity returns the DefaultSeverity field value if set, zero value otherwise.
func (o *JsonErrorLogPublisherResponse) GetDefaultSeverity() []EnumlogPublisherDefaultSeverityProp {
	if o == nil || isNil(o.DefaultSeverity) {
		var ret []EnumlogPublisherDefaultSeverityProp
		return ret
	}
	return o.DefaultSeverity
}

// GetDefaultSeverityOk returns a tuple with the DefaultSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonErrorLogPublisherResponse) GetDefaultSeverityOk() ([]EnumlogPublisherDefaultSeverityProp, bool) {
	if o == nil || isNil(o.DefaultSeverity) {
    return nil, false
	}
	return o.DefaultSeverity, true
}

// HasDefaultSeverity returns a boolean if a field has been set.
func (o *JsonErrorLogPublisherResponse) HasDefaultSeverity() bool {
	if o != nil && !isNil(o.DefaultSeverity) {
		return true
	}

	return false
}

// SetDefaultSeverity gets a reference to the given []EnumlogPublisherDefaultSeverityProp and assigns it to the DefaultSeverity field.
func (o *JsonErrorLogPublisherResponse) SetDefaultSeverity(v []EnumlogPublisherDefaultSeverityProp) {
	o.DefaultSeverity = v
}

// GetOverrideSeverity returns the OverrideSeverity field value if set, zero value otherwise.
func (o *JsonErrorLogPublisherResponse) GetOverrideSeverity() []string {
	if o == nil || isNil(o.OverrideSeverity) {
		var ret []string
		return ret
	}
	return o.OverrideSeverity
}

// GetOverrideSeverityOk returns a tuple with the OverrideSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonErrorLogPublisherResponse) GetOverrideSeverityOk() ([]string, bool) {
	if o == nil || isNil(o.OverrideSeverity) {
    return nil, false
	}
	return o.OverrideSeverity, true
}

// HasOverrideSeverity returns a boolean if a field has been set.
func (o *JsonErrorLogPublisherResponse) HasOverrideSeverity() bool {
	if o != nil && !isNil(o.OverrideSeverity) {
		return true
	}

	return false
}

// SetOverrideSeverity gets a reference to the given []string and assigns it to the OverrideSeverity field.
func (o *JsonErrorLogPublisherResponse) SetOverrideSeverity(v []string) {
	o.OverrideSeverity = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *JsonErrorLogPublisherResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonErrorLogPublisherResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *JsonErrorLogPublisherResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *JsonErrorLogPublisherResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *JsonErrorLogPublisherResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *JsonErrorLogPublisherResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *JsonErrorLogPublisherResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetLoggingErrorBehavior returns the LoggingErrorBehavior field value if set, zero value otherwise.
func (o *JsonErrorLogPublisherResponse) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp {
	if o == nil || isNil(o.LoggingErrorBehavior) {
		var ret EnumlogPublisherLoggingErrorBehaviorProp
		return ret
	}
	return *o.LoggingErrorBehavior
}

// GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonErrorLogPublisherResponse) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool) {
	if o == nil || isNil(o.LoggingErrorBehavior) {
    return nil, false
	}
	return o.LoggingErrorBehavior, true
}

// HasLoggingErrorBehavior returns a boolean if a field has been set.
func (o *JsonErrorLogPublisherResponse) HasLoggingErrorBehavior() bool {
	if o != nil && !isNil(o.LoggingErrorBehavior) {
		return true
	}

	return false
}

// SetLoggingErrorBehavior gets a reference to the given EnumlogPublisherLoggingErrorBehaviorProp and assigns it to the LoggingErrorBehavior field.
func (o *JsonErrorLogPublisherResponse) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp) {
	o.LoggingErrorBehavior = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *JsonErrorLogPublisherResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonErrorLogPublisherResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *JsonErrorLogPublisherResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *JsonErrorLogPublisherResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

func (o JsonErrorLogPublisherResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["logFile"] = o.LogFile
	}
	if true {
		toSerialize["logFilePermissions"] = o.LogFilePermissions
	}
	if true {
		toSerialize["rotationPolicy"] = o.RotationPolicy
	}
	if !isNil(o.RotationListener) {
		toSerialize["rotationListener"] = o.RotationListener
	}
	if true {
		toSerialize["retentionPolicy"] = o.RetentionPolicy
	}
	if !isNil(o.CompressionMechanism) {
		toSerialize["compressionMechanism"] = o.CompressionMechanism
	}
	if !isNil(o.SignLog) {
		toSerialize["signLog"] = o.SignLog
	}
	if !isNil(o.EncryptLog) {
		toSerialize["encryptLog"] = o.EncryptLog
	}
	if !isNil(o.EncryptionSettingsDefinitionID) {
		toSerialize["encryptionSettingsDefinitionID"] = o.EncryptionSettingsDefinitionID
	}
	if !isNil(o.Append) {
		toSerialize["append"] = o.Append
	}
	if true {
		toSerialize["asynchronous"] = o.Asynchronous
	}
	if !isNil(o.AutoFlush) {
		toSerialize["autoFlush"] = o.AutoFlush
	}
	if !isNil(o.BufferSize) {
		toSerialize["bufferSize"] = o.BufferSize
	}
	if !isNil(o.QueueSize) {
		toSerialize["queueSize"] = o.QueueSize
	}
	if !isNil(o.TimeInterval) {
		toSerialize["timeInterval"] = o.TimeInterval
	}
	if !isNil(o.WriteMultiLineMessages) {
		toSerialize["writeMultiLineMessages"] = o.WriteMultiLineMessages
	}
	if !isNil(o.IncludeProductName) {
		toSerialize["includeProductName"] = o.IncludeProductName
	}
	if !isNil(o.IncludeInstanceName) {
		toSerialize["includeInstanceName"] = o.IncludeInstanceName
	}
	if !isNil(o.IncludeStartupID) {
		toSerialize["includeStartupID"] = o.IncludeStartupID
	}
	if !isNil(o.IncludeThreadID) {
		toSerialize["includeThreadID"] = o.IncludeThreadID
	}
	if !isNil(o.GenerifyMessageStringsWhenPossible) {
		toSerialize["generifyMessageStringsWhenPossible"] = o.GenerifyMessageStringsWhenPossible
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
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.LoggingErrorBehavior) {
		toSerialize["loggingErrorBehavior"] = o.LoggingErrorBehavior
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableJsonErrorLogPublisherResponse struct {
	value *JsonErrorLogPublisherResponse
	isSet bool
}

func (v NullableJsonErrorLogPublisherResponse) Get() *JsonErrorLogPublisherResponse {
	return v.value
}

func (v *NullableJsonErrorLogPublisherResponse) Set(val *JsonErrorLogPublisherResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonErrorLogPublisherResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonErrorLogPublisherResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonErrorLogPublisherResponse(val *JsonErrorLogPublisherResponse) *NullableJsonErrorLogPublisherResponse {
	return &NullableJsonErrorLogPublisherResponse{value: val, isSet: true}
}

func (v NullableJsonErrorLogPublisherResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonErrorLogPublisherResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


