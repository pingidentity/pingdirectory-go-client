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

// checks if the AddGroovyScriptedFileBasedErrorLogPublisherRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddGroovyScriptedFileBasedErrorLogPublisherRequest{}

// AddGroovyScriptedFileBasedErrorLogPublisherRequest struct for AddGroovyScriptedFileBasedErrorLogPublisherRequest
type AddGroovyScriptedFileBasedErrorLogPublisherRequest struct {
	// Name of the new Log Publisher
	PublisherName string                                                  `json:"publisherName"`
	Schemas       []EnumgroovyScriptedFileBasedErrorLogPublisherSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted File Based Error Log Publisher.
	ScriptClass string `json:"scriptClass"`
	// The file name to use for the log files generated by the Scripted File Based Error Log Publisher. The path to the file can be specified either as relative to the server root or as an absolute path.
	LogFile string `json:"logFile"`
	// The UNIX permissions of the log files created by this Scripted File Based Error Log Publisher.
	LogFilePermissions *string `json:"logFilePermissions,omitempty"`
	// The rotation policy to use for the Scripted File Based Error Log Publisher .
	RotationPolicy []string `json:"rotationPolicy,omitempty"`
	// A listener that should be notified whenever a log file is rotated out of service.
	RotationListener []string `json:"rotationListener,omitempty"`
	// The retention policy to use for the Scripted File Based Error Log Publisher .
	RetentionPolicy      []string                                  `json:"retentionPolicy,omitempty"`
	CompressionMechanism *EnumlogPublisherCompressionMechanismProp `json:"compressionMechanism,omitempty"`
	// Indicates whether the log should be cryptographically signed so that the log content cannot be altered in an undetectable manner.
	SignLog *bool `json:"signLog,omitempty"`
	// Indicates whether log files should be encrypted so that their content is not available to unauthorized users.
	EncryptLog *bool `json:"encryptLog,omitempty"`
	// Specifies the ID of the encryption settings definition that should be used to encrypt the data. If this is not provided, the server's preferred encryption settings definition will be used. The \"encryption-settings list\" command can be used to obtain a list of the encryption settings definitions available in the server.
	EncryptionSettingsDefinitionID *string `json:"encryptionSettingsDefinitionID,omitempty"`
	// Specifies whether to append to existing log files.
	Append *bool `json:"append,omitempty"`
	// The set of arguments used to customize the behavior for the Scripted File Based Error Log Publisher. Each configuration property should be given in the form 'name=value'.
	ScriptArgument []string `json:"scriptArgument,omitempty"`
	// Indicates whether the Scripted File Based Error Log Publisher will publish records asynchronously.
	Asynchronous *bool `json:"asynchronous,omitempty"`
	// Specifies whether to flush the writer after every log record.
	AutoFlush *bool `json:"autoFlush,omitempty"`
	// Specifies the log file buffer size.
	BufferSize *string `json:"bufferSize,omitempty"`
	// The maximum number of log records that can be stored in the asynchronous queue.
	QueueSize *int32 `json:"queueSize,omitempty"`
	// Specifies the interval at which to check whether the log files need to be rotated.
	TimeInterval *string `json:"timeInterval,omitempty"`
	// Specifies the default severity levels for the logger.
	DefaultSeverity []EnumlogPublisherDefaultSeverityProp `json:"defaultSeverity,omitempty"`
	// Specifies the override severity levels for the logger based on the category of the messages.
	OverrideSeverity []string `json:"overrideSeverity,omitempty"`
	// A description for this Log Publisher
	Description *string `json:"description,omitempty"`
	// Indicates whether the Log Publisher is enabled for use.
	Enabled              bool                                      `json:"enabled"`
	LoggingErrorBehavior *EnumlogPublisherLoggingErrorBehaviorProp `json:"loggingErrorBehavior,omitempty"`
}

// NewAddGroovyScriptedFileBasedErrorLogPublisherRequest instantiates a new AddGroovyScriptedFileBasedErrorLogPublisherRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddGroovyScriptedFileBasedErrorLogPublisherRequest(publisherName string, schemas []EnumgroovyScriptedFileBasedErrorLogPublisherSchemaUrn, scriptClass string, logFile string, enabled bool) *AddGroovyScriptedFileBasedErrorLogPublisherRequest {
	this := AddGroovyScriptedFileBasedErrorLogPublisherRequest{}
	this.PublisherName = publisherName
	this.Schemas = schemas
	this.ScriptClass = scriptClass
	this.LogFile = logFile
	this.Enabled = enabled
	return &this
}

// NewAddGroovyScriptedFileBasedErrorLogPublisherRequestWithDefaults instantiates a new AddGroovyScriptedFileBasedErrorLogPublisherRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddGroovyScriptedFileBasedErrorLogPublisherRequestWithDefaults() *AddGroovyScriptedFileBasedErrorLogPublisherRequest {
	this := AddGroovyScriptedFileBasedErrorLogPublisherRequest{}
	return &this
}

// GetPublisherName returns the PublisherName field value
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetPublisherName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublisherName
}

// GetPublisherNameOk returns a tuple with the PublisherName field value
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetPublisherNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublisherName, true
}

// SetPublisherName sets field value
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) SetPublisherName(v string) {
	o.PublisherName = v
}

// GetSchemas returns the Schemas field value
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetSchemas() []EnumgroovyScriptedFileBasedErrorLogPublisherSchemaUrn {
	if o == nil {
		var ret []EnumgroovyScriptedFileBasedErrorLogPublisherSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetSchemasOk() ([]EnumgroovyScriptedFileBasedErrorLogPublisherSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) SetSchemas(v []EnumgroovyScriptedFileBasedErrorLogPublisherSchemaUrn) {
	o.Schemas = v
}

// GetScriptClass returns the ScriptClass field value
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetScriptClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScriptClass
}

// GetScriptClassOk returns a tuple with the ScriptClass field value
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetScriptClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptClass, true
}

// SetScriptClass sets field value
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) SetScriptClass(v string) {
	o.ScriptClass = v
}

// GetLogFile returns the LogFile field value
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetLogFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogFile
}

// GetLogFileOk returns a tuple with the LogFile field value
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetLogFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogFile, true
}

// SetLogFile sets field value
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) SetLogFile(v string) {
	o.LogFile = v
}

// GetLogFilePermissions returns the LogFilePermissions field value if set, zero value otherwise.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetLogFilePermissions() string {
	if o == nil || IsNil(o.LogFilePermissions) {
		var ret string
		return ret
	}
	return *o.LogFilePermissions
}

// GetLogFilePermissionsOk returns a tuple with the LogFilePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetLogFilePermissionsOk() (*string, bool) {
	if o == nil || IsNil(o.LogFilePermissions) {
		return nil, false
	}
	return o.LogFilePermissions, true
}

// HasLogFilePermissions returns a boolean if a field has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) HasLogFilePermissions() bool {
	if o != nil && !IsNil(o.LogFilePermissions) {
		return true
	}

	return false
}

// SetLogFilePermissions gets a reference to the given string and assigns it to the LogFilePermissions field.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) SetLogFilePermissions(v string) {
	o.LogFilePermissions = &v
}

// GetRotationPolicy returns the RotationPolicy field value if set, zero value otherwise.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetRotationPolicy() []string {
	if o == nil || IsNil(o.RotationPolicy) {
		var ret []string
		return ret
	}
	return o.RotationPolicy
}

// GetRotationPolicyOk returns a tuple with the RotationPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetRotationPolicyOk() ([]string, bool) {
	if o == nil || IsNil(o.RotationPolicy) {
		return nil, false
	}
	return o.RotationPolicy, true
}

// HasRotationPolicy returns a boolean if a field has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) HasRotationPolicy() bool {
	if o != nil && !IsNil(o.RotationPolicy) {
		return true
	}

	return false
}

// SetRotationPolicy gets a reference to the given []string and assigns it to the RotationPolicy field.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) SetRotationPolicy(v []string) {
	o.RotationPolicy = v
}

// GetRotationListener returns the RotationListener field value if set, zero value otherwise.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetRotationListener() []string {
	if o == nil || IsNil(o.RotationListener) {
		var ret []string
		return ret
	}
	return o.RotationListener
}

// GetRotationListenerOk returns a tuple with the RotationListener field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetRotationListenerOk() ([]string, bool) {
	if o == nil || IsNil(o.RotationListener) {
		return nil, false
	}
	return o.RotationListener, true
}

// HasRotationListener returns a boolean if a field has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) HasRotationListener() bool {
	if o != nil && !IsNil(o.RotationListener) {
		return true
	}

	return false
}

// SetRotationListener gets a reference to the given []string and assigns it to the RotationListener field.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) SetRotationListener(v []string) {
	o.RotationListener = v
}

// GetRetentionPolicy returns the RetentionPolicy field value if set, zero value otherwise.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetRetentionPolicy() []string {
	if o == nil || IsNil(o.RetentionPolicy) {
		var ret []string
		return ret
	}
	return o.RetentionPolicy
}

// GetRetentionPolicyOk returns a tuple with the RetentionPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetRetentionPolicyOk() ([]string, bool) {
	if o == nil || IsNil(o.RetentionPolicy) {
		return nil, false
	}
	return o.RetentionPolicy, true
}

// HasRetentionPolicy returns a boolean if a field has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) HasRetentionPolicy() bool {
	if o != nil && !IsNil(o.RetentionPolicy) {
		return true
	}

	return false
}

// SetRetentionPolicy gets a reference to the given []string and assigns it to the RetentionPolicy field.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) SetRetentionPolicy(v []string) {
	o.RetentionPolicy = v
}

// GetCompressionMechanism returns the CompressionMechanism field value if set, zero value otherwise.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetCompressionMechanism() EnumlogPublisherCompressionMechanismProp {
	if o == nil || IsNil(o.CompressionMechanism) {
		var ret EnumlogPublisherCompressionMechanismProp
		return ret
	}
	return *o.CompressionMechanism
}

// GetCompressionMechanismOk returns a tuple with the CompressionMechanism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetCompressionMechanismOk() (*EnumlogPublisherCompressionMechanismProp, bool) {
	if o == nil || IsNil(o.CompressionMechanism) {
		return nil, false
	}
	return o.CompressionMechanism, true
}

// HasCompressionMechanism returns a boolean if a field has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) HasCompressionMechanism() bool {
	if o != nil && !IsNil(o.CompressionMechanism) {
		return true
	}

	return false
}

// SetCompressionMechanism gets a reference to the given EnumlogPublisherCompressionMechanismProp and assigns it to the CompressionMechanism field.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) SetCompressionMechanism(v EnumlogPublisherCompressionMechanismProp) {
	o.CompressionMechanism = &v
}

// GetSignLog returns the SignLog field value if set, zero value otherwise.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetSignLog() bool {
	if o == nil || IsNil(o.SignLog) {
		var ret bool
		return ret
	}
	return *o.SignLog
}

// GetSignLogOk returns a tuple with the SignLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetSignLogOk() (*bool, bool) {
	if o == nil || IsNil(o.SignLog) {
		return nil, false
	}
	return o.SignLog, true
}

// HasSignLog returns a boolean if a field has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) HasSignLog() bool {
	if o != nil && !IsNil(o.SignLog) {
		return true
	}

	return false
}

// SetSignLog gets a reference to the given bool and assigns it to the SignLog field.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) SetSignLog(v bool) {
	o.SignLog = &v
}

// GetEncryptLog returns the EncryptLog field value if set, zero value otherwise.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetEncryptLog() bool {
	if o == nil || IsNil(o.EncryptLog) {
		var ret bool
		return ret
	}
	return *o.EncryptLog
}

// GetEncryptLogOk returns a tuple with the EncryptLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetEncryptLogOk() (*bool, bool) {
	if o == nil || IsNil(o.EncryptLog) {
		return nil, false
	}
	return o.EncryptLog, true
}

// HasEncryptLog returns a boolean if a field has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) HasEncryptLog() bool {
	if o != nil && !IsNil(o.EncryptLog) {
		return true
	}

	return false
}

// SetEncryptLog gets a reference to the given bool and assigns it to the EncryptLog field.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) SetEncryptLog(v bool) {
	o.EncryptLog = &v
}

// GetEncryptionSettingsDefinitionID returns the EncryptionSettingsDefinitionID field value if set, zero value otherwise.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetEncryptionSettingsDefinitionID() string {
	if o == nil || IsNil(o.EncryptionSettingsDefinitionID) {
		var ret string
		return ret
	}
	return *o.EncryptionSettingsDefinitionID
}

// GetEncryptionSettingsDefinitionIDOk returns a tuple with the EncryptionSettingsDefinitionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetEncryptionSettingsDefinitionIDOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionSettingsDefinitionID) {
		return nil, false
	}
	return o.EncryptionSettingsDefinitionID, true
}

// HasEncryptionSettingsDefinitionID returns a boolean if a field has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) HasEncryptionSettingsDefinitionID() bool {
	if o != nil && !IsNil(o.EncryptionSettingsDefinitionID) {
		return true
	}

	return false
}

// SetEncryptionSettingsDefinitionID gets a reference to the given string and assigns it to the EncryptionSettingsDefinitionID field.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) SetEncryptionSettingsDefinitionID(v string) {
	o.EncryptionSettingsDefinitionID = &v
}

// GetAppend returns the Append field value if set, zero value otherwise.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetAppend() bool {
	if o == nil || IsNil(o.Append) {
		var ret bool
		return ret
	}
	return *o.Append
}

// GetAppendOk returns a tuple with the Append field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetAppendOk() (*bool, bool) {
	if o == nil || IsNil(o.Append) {
		return nil, false
	}
	return o.Append, true
}

// HasAppend returns a boolean if a field has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) HasAppend() bool {
	if o != nil && !IsNil(o.Append) {
		return true
	}

	return false
}

// SetAppend gets a reference to the given bool and assigns it to the Append field.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) SetAppend(v bool) {
	o.Append = &v
}

// GetScriptArgument returns the ScriptArgument field value if set, zero value otherwise.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetScriptArgument() []string {
	if o == nil || IsNil(o.ScriptArgument) {
		var ret []string
		return ret
	}
	return o.ScriptArgument
}

// GetScriptArgumentOk returns a tuple with the ScriptArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetScriptArgumentOk() ([]string, bool) {
	if o == nil || IsNil(o.ScriptArgument) {
		return nil, false
	}
	return o.ScriptArgument, true
}

// HasScriptArgument returns a boolean if a field has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) HasScriptArgument() bool {
	if o != nil && !IsNil(o.ScriptArgument) {
		return true
	}

	return false
}

// SetScriptArgument gets a reference to the given []string and assigns it to the ScriptArgument field.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) SetScriptArgument(v []string) {
	o.ScriptArgument = v
}

// GetAsynchronous returns the Asynchronous field value if set, zero value otherwise.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetAsynchronous() bool {
	if o == nil || IsNil(o.Asynchronous) {
		var ret bool
		return ret
	}
	return *o.Asynchronous
}

// GetAsynchronousOk returns a tuple with the Asynchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetAsynchronousOk() (*bool, bool) {
	if o == nil || IsNil(o.Asynchronous) {
		return nil, false
	}
	return o.Asynchronous, true
}

// HasAsynchronous returns a boolean if a field has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) HasAsynchronous() bool {
	if o != nil && !IsNil(o.Asynchronous) {
		return true
	}

	return false
}

// SetAsynchronous gets a reference to the given bool and assigns it to the Asynchronous field.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) SetAsynchronous(v bool) {
	o.Asynchronous = &v
}

// GetAutoFlush returns the AutoFlush field value if set, zero value otherwise.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetAutoFlush() bool {
	if o == nil || IsNil(o.AutoFlush) {
		var ret bool
		return ret
	}
	return *o.AutoFlush
}

// GetAutoFlushOk returns a tuple with the AutoFlush field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetAutoFlushOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoFlush) {
		return nil, false
	}
	return o.AutoFlush, true
}

// HasAutoFlush returns a boolean if a field has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) HasAutoFlush() bool {
	if o != nil && !IsNil(o.AutoFlush) {
		return true
	}

	return false
}

// SetAutoFlush gets a reference to the given bool and assigns it to the AutoFlush field.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) SetAutoFlush(v bool) {
	o.AutoFlush = &v
}

// GetBufferSize returns the BufferSize field value if set, zero value otherwise.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetBufferSize() string {
	if o == nil || IsNil(o.BufferSize) {
		var ret string
		return ret
	}
	return *o.BufferSize
}

// GetBufferSizeOk returns a tuple with the BufferSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetBufferSizeOk() (*string, bool) {
	if o == nil || IsNil(o.BufferSize) {
		return nil, false
	}
	return o.BufferSize, true
}

// HasBufferSize returns a boolean if a field has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) HasBufferSize() bool {
	if o != nil && !IsNil(o.BufferSize) {
		return true
	}

	return false
}

// SetBufferSize gets a reference to the given string and assigns it to the BufferSize field.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) SetBufferSize(v string) {
	o.BufferSize = &v
}

// GetQueueSize returns the QueueSize field value if set, zero value otherwise.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetQueueSize() int32 {
	if o == nil || IsNil(o.QueueSize) {
		var ret int32
		return ret
	}
	return *o.QueueSize
}

// GetQueueSizeOk returns a tuple with the QueueSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetQueueSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.QueueSize) {
		return nil, false
	}
	return o.QueueSize, true
}

// HasQueueSize returns a boolean if a field has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) HasQueueSize() bool {
	if o != nil && !IsNil(o.QueueSize) {
		return true
	}

	return false
}

// SetQueueSize gets a reference to the given int32 and assigns it to the QueueSize field.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) SetQueueSize(v int32) {
	o.QueueSize = &v
}

// GetTimeInterval returns the TimeInterval field value if set, zero value otherwise.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetTimeInterval() string {
	if o == nil || IsNil(o.TimeInterval) {
		var ret string
		return ret
	}
	return *o.TimeInterval
}

// GetTimeIntervalOk returns a tuple with the TimeInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetTimeIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.TimeInterval) {
		return nil, false
	}
	return o.TimeInterval, true
}

// HasTimeInterval returns a boolean if a field has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) HasTimeInterval() bool {
	if o != nil && !IsNil(o.TimeInterval) {
		return true
	}

	return false
}

// SetTimeInterval gets a reference to the given string and assigns it to the TimeInterval field.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) SetTimeInterval(v string) {
	o.TimeInterval = &v
}

// GetDefaultSeverity returns the DefaultSeverity field value if set, zero value otherwise.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetDefaultSeverity() []EnumlogPublisherDefaultSeverityProp {
	if o == nil || IsNil(o.DefaultSeverity) {
		var ret []EnumlogPublisherDefaultSeverityProp
		return ret
	}
	return o.DefaultSeverity
}

// GetDefaultSeverityOk returns a tuple with the DefaultSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetDefaultSeverityOk() ([]EnumlogPublisherDefaultSeverityProp, bool) {
	if o == nil || IsNil(o.DefaultSeverity) {
		return nil, false
	}
	return o.DefaultSeverity, true
}

// HasDefaultSeverity returns a boolean if a field has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) HasDefaultSeverity() bool {
	if o != nil && !IsNil(o.DefaultSeverity) {
		return true
	}

	return false
}

// SetDefaultSeverity gets a reference to the given []EnumlogPublisherDefaultSeverityProp and assigns it to the DefaultSeverity field.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) SetDefaultSeverity(v []EnumlogPublisherDefaultSeverityProp) {
	o.DefaultSeverity = v
}

// GetOverrideSeverity returns the OverrideSeverity field value if set, zero value otherwise.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetOverrideSeverity() []string {
	if o == nil || IsNil(o.OverrideSeverity) {
		var ret []string
		return ret
	}
	return o.OverrideSeverity
}

// GetOverrideSeverityOk returns a tuple with the OverrideSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetOverrideSeverityOk() ([]string, bool) {
	if o == nil || IsNil(o.OverrideSeverity) {
		return nil, false
	}
	return o.OverrideSeverity, true
}

// HasOverrideSeverity returns a boolean if a field has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) HasOverrideSeverity() bool {
	if o != nil && !IsNil(o.OverrideSeverity) {
		return true
	}

	return false
}

// SetOverrideSeverity gets a reference to the given []string and assigns it to the OverrideSeverity field.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) SetOverrideSeverity(v []string) {
	o.OverrideSeverity = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetLoggingErrorBehavior returns the LoggingErrorBehavior field value if set, zero value otherwise.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp {
	if o == nil || IsNil(o.LoggingErrorBehavior) {
		var ret EnumlogPublisherLoggingErrorBehaviorProp
		return ret
	}
	return *o.LoggingErrorBehavior
}

// GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool) {
	if o == nil || IsNil(o.LoggingErrorBehavior) {
		return nil, false
	}
	return o.LoggingErrorBehavior, true
}

// HasLoggingErrorBehavior returns a boolean if a field has been set.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) HasLoggingErrorBehavior() bool {
	if o != nil && !IsNil(o.LoggingErrorBehavior) {
		return true
	}

	return false
}

// SetLoggingErrorBehavior gets a reference to the given EnumlogPublisherLoggingErrorBehaviorProp and assigns it to the LoggingErrorBehavior field.
func (o *AddGroovyScriptedFileBasedErrorLogPublisherRequest) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp) {
	o.LoggingErrorBehavior = &v
}

func (o AddGroovyScriptedFileBasedErrorLogPublisherRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddGroovyScriptedFileBasedErrorLogPublisherRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["publisherName"] = o.PublisherName
	toSerialize["schemas"] = o.Schemas
	toSerialize["scriptClass"] = o.ScriptClass
	toSerialize["logFile"] = o.LogFile
	if !IsNil(o.LogFilePermissions) {
		toSerialize["logFilePermissions"] = o.LogFilePermissions
	}
	if !IsNil(o.RotationPolicy) {
		toSerialize["rotationPolicy"] = o.RotationPolicy
	}
	if !IsNil(o.RotationListener) {
		toSerialize["rotationListener"] = o.RotationListener
	}
	if !IsNil(o.RetentionPolicy) {
		toSerialize["retentionPolicy"] = o.RetentionPolicy
	}
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
	if !IsNil(o.Append) {
		toSerialize["append"] = o.Append
	}
	if !IsNil(o.ScriptArgument) {
		toSerialize["scriptArgument"] = o.ScriptArgument
	}
	if !IsNil(o.Asynchronous) {
		toSerialize["asynchronous"] = o.Asynchronous
	}
	if !IsNil(o.AutoFlush) {
		toSerialize["autoFlush"] = o.AutoFlush
	}
	if !IsNil(o.BufferSize) {
		toSerialize["bufferSize"] = o.BufferSize
	}
	if !IsNil(o.QueueSize) {
		toSerialize["queueSize"] = o.QueueSize
	}
	if !IsNil(o.TimeInterval) {
		toSerialize["timeInterval"] = o.TimeInterval
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

type NullableAddGroovyScriptedFileBasedErrorLogPublisherRequest struct {
	value *AddGroovyScriptedFileBasedErrorLogPublisherRequest
	isSet bool
}

func (v NullableAddGroovyScriptedFileBasedErrorLogPublisherRequest) Get() *AddGroovyScriptedFileBasedErrorLogPublisherRequest {
	return v.value
}

func (v *NullableAddGroovyScriptedFileBasedErrorLogPublisherRequest) Set(val *AddGroovyScriptedFileBasedErrorLogPublisherRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddGroovyScriptedFileBasedErrorLogPublisherRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddGroovyScriptedFileBasedErrorLogPublisherRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddGroovyScriptedFileBasedErrorLogPublisherRequest(val *AddGroovyScriptedFileBasedErrorLogPublisherRequest) *NullableAddGroovyScriptedFileBasedErrorLogPublisherRequest {
	return &NullableAddGroovyScriptedFileBasedErrorLogPublisherRequest{value: val, isSet: true}
}

func (v NullableAddGroovyScriptedFileBasedErrorLogPublisherRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddGroovyScriptedFileBasedErrorLogPublisherRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
