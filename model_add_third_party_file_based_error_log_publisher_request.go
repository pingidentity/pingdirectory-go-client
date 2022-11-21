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

// AddThirdPartyFileBasedErrorLogPublisherRequest struct for AddThirdPartyFileBasedErrorLogPublisherRequest
type AddThirdPartyFileBasedErrorLogPublisherRequest struct {
	// Name of the new Log Publisher
	PublisherName string `json:"publisherName"`
	Schemas []EnumthirdPartyFileBasedErrorLogPublisherSchemaUrn `json:"schemas"`
	// The file name to use for the log files generated by the Third Party File Based Error Log Publisher. The path to the file can be specified either as relative to the server root or as an absolute path.
	LogFile string `json:"logFile"`
	// The UNIX permissions of the log files created by this Third Party File Based Error Log Publisher.
	LogFilePermissions string `json:"logFilePermissions"`
	// The rotation policy to use for the Third Party File Based Error Log Publisher .
	RotationPolicy []string `json:"rotationPolicy"`
	// A listener that should be notified whenever a log file is rotated out of service.
	RotationListener []string `json:"rotationListener,omitempty"`
	// The retention policy to use for the Third Party File Based Error Log Publisher .
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
	// The fully-qualified name of the Java class providing the logic for the Third Party File Based Error Log Publisher.
	ExtensionClass string `json:"extensionClass"`
	// The set of arguments used to customize the behavior for the Third Party File Based Error Log Publisher. Each configuration property should be given in the form 'name=value'.
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// Indicates whether the Third Party File Based Error Log Publisher will publish records asynchronously.
	Asynchronous bool `json:"asynchronous"`
	// Specifies whether to flush the writer after every log record.
	AutoFlush *bool `json:"autoFlush,omitempty"`
	// Specifies the log file buffer size.
	BufferSize *string `json:"bufferSize,omitempty"`
	// The maximum number of log records that can be stored in the asynchronous queue.
	QueueSize *int32 `json:"queueSize,omitempty"`
	// Specifies the interval at which to check whether the log files need to be rotated.
	TimeInterval *string `json:"timeInterval,omitempty"`
	DefaultSeverity []EnumlogPublisherDefaultSeverityProp `json:"defaultSeverity,omitempty"`
	// Specifies the override severity levels for the logger based on the category of the messages.
	OverrideSeverity []string `json:"overrideSeverity,omitempty"`
	// A description for this Log Publisher
	Description *string `json:"description,omitempty"`
	// Indicates whether the Log Publisher is enabled for use.
	Enabled bool `json:"enabled"`
	LoggingErrorBehavior *EnumlogPublisherLoggingErrorBehaviorProp `json:"loggingErrorBehavior,omitempty"`
}

// NewAddThirdPartyFileBasedErrorLogPublisherRequest instantiates a new AddThirdPartyFileBasedErrorLogPublisherRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddThirdPartyFileBasedErrorLogPublisherRequest(publisherName string, schemas []EnumthirdPartyFileBasedErrorLogPublisherSchemaUrn, logFile string, logFilePermissions string, rotationPolicy []string, retentionPolicy []string, extensionClass string, asynchronous bool, enabled bool) *AddThirdPartyFileBasedErrorLogPublisherRequest {
	this := AddThirdPartyFileBasedErrorLogPublisherRequest{}
	this.PublisherName = publisherName
	this.Schemas = schemas
	this.LogFile = logFile
	this.LogFilePermissions = logFilePermissions
	this.RotationPolicy = rotationPolicy
	this.RetentionPolicy = retentionPolicy
	this.ExtensionClass = extensionClass
	this.Asynchronous = asynchronous
	this.Enabled = enabled
	return &this
}

// NewAddThirdPartyFileBasedErrorLogPublisherRequestWithDefaults instantiates a new AddThirdPartyFileBasedErrorLogPublisherRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddThirdPartyFileBasedErrorLogPublisherRequestWithDefaults() *AddThirdPartyFileBasedErrorLogPublisherRequest {
	this := AddThirdPartyFileBasedErrorLogPublisherRequest{}
	return &this
}

// GetPublisherName returns the PublisherName field value
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetPublisherName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublisherName
}

// GetPublisherNameOk returns a tuple with the PublisherName field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetPublisherNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PublisherName, true
}

// SetPublisherName sets field value
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) SetPublisherName(v string) {
	o.PublisherName = v
}

// GetSchemas returns the Schemas field value
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetSchemas() []EnumthirdPartyFileBasedErrorLogPublisherSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyFileBasedErrorLogPublisherSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetSchemasOk() ([]EnumthirdPartyFileBasedErrorLogPublisherSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) SetSchemas(v []EnumthirdPartyFileBasedErrorLogPublisherSchemaUrn) {
	o.Schemas = v
}

// GetLogFile returns the LogFile field value
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetLogFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogFile
}

// GetLogFileOk returns a tuple with the LogFile field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetLogFileOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LogFile, true
}

// SetLogFile sets field value
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) SetLogFile(v string) {
	o.LogFile = v
}

// GetLogFilePermissions returns the LogFilePermissions field value
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetLogFilePermissions() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogFilePermissions
}

// GetLogFilePermissionsOk returns a tuple with the LogFilePermissions field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetLogFilePermissionsOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LogFilePermissions, true
}

// SetLogFilePermissions sets field value
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) SetLogFilePermissions(v string) {
	o.LogFilePermissions = v
}

// GetRotationPolicy returns the RotationPolicy field value
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetRotationPolicy() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RotationPolicy
}

// GetRotationPolicyOk returns a tuple with the RotationPolicy field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetRotationPolicyOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RotationPolicy, true
}

// SetRotationPolicy sets field value
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) SetRotationPolicy(v []string) {
	o.RotationPolicy = v
}

// GetRotationListener returns the RotationListener field value if set, zero value otherwise.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetRotationListener() []string {
	if o == nil || isNil(o.RotationListener) {
		var ret []string
		return ret
	}
	return o.RotationListener
}

// GetRotationListenerOk returns a tuple with the RotationListener field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetRotationListenerOk() ([]string, bool) {
	if o == nil || isNil(o.RotationListener) {
    return nil, false
	}
	return o.RotationListener, true
}

// HasRotationListener returns a boolean if a field has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) HasRotationListener() bool {
	if o != nil && !isNil(o.RotationListener) {
		return true
	}

	return false
}

// SetRotationListener gets a reference to the given []string and assigns it to the RotationListener field.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) SetRotationListener(v []string) {
	o.RotationListener = v
}

// GetRetentionPolicy returns the RetentionPolicy field value
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetRetentionPolicy() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RetentionPolicy
}

// GetRetentionPolicyOk returns a tuple with the RetentionPolicy field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetRetentionPolicyOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RetentionPolicy, true
}

// SetRetentionPolicy sets field value
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) SetRetentionPolicy(v []string) {
	o.RetentionPolicy = v
}

// GetCompressionMechanism returns the CompressionMechanism field value if set, zero value otherwise.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetCompressionMechanism() EnumlogPublisherCompressionMechanismProp {
	if o == nil || isNil(o.CompressionMechanism) {
		var ret EnumlogPublisherCompressionMechanismProp
		return ret
	}
	return *o.CompressionMechanism
}

// GetCompressionMechanismOk returns a tuple with the CompressionMechanism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetCompressionMechanismOk() (*EnumlogPublisherCompressionMechanismProp, bool) {
	if o == nil || isNil(o.CompressionMechanism) {
    return nil, false
	}
	return o.CompressionMechanism, true
}

// HasCompressionMechanism returns a boolean if a field has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) HasCompressionMechanism() bool {
	if o != nil && !isNil(o.CompressionMechanism) {
		return true
	}

	return false
}

// SetCompressionMechanism gets a reference to the given EnumlogPublisherCompressionMechanismProp and assigns it to the CompressionMechanism field.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) SetCompressionMechanism(v EnumlogPublisherCompressionMechanismProp) {
	o.CompressionMechanism = &v
}

// GetSignLog returns the SignLog field value if set, zero value otherwise.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetSignLog() bool {
	if o == nil || isNil(o.SignLog) {
		var ret bool
		return ret
	}
	return *o.SignLog
}

// GetSignLogOk returns a tuple with the SignLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetSignLogOk() (*bool, bool) {
	if o == nil || isNil(o.SignLog) {
    return nil, false
	}
	return o.SignLog, true
}

// HasSignLog returns a boolean if a field has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) HasSignLog() bool {
	if o != nil && !isNil(o.SignLog) {
		return true
	}

	return false
}

// SetSignLog gets a reference to the given bool and assigns it to the SignLog field.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) SetSignLog(v bool) {
	o.SignLog = &v
}

// GetEncryptLog returns the EncryptLog field value if set, zero value otherwise.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetEncryptLog() bool {
	if o == nil || isNil(o.EncryptLog) {
		var ret bool
		return ret
	}
	return *o.EncryptLog
}

// GetEncryptLogOk returns a tuple with the EncryptLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetEncryptLogOk() (*bool, bool) {
	if o == nil || isNil(o.EncryptLog) {
    return nil, false
	}
	return o.EncryptLog, true
}

// HasEncryptLog returns a boolean if a field has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) HasEncryptLog() bool {
	if o != nil && !isNil(o.EncryptLog) {
		return true
	}

	return false
}

// SetEncryptLog gets a reference to the given bool and assigns it to the EncryptLog field.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) SetEncryptLog(v bool) {
	o.EncryptLog = &v
}

// GetEncryptionSettingsDefinitionID returns the EncryptionSettingsDefinitionID field value if set, zero value otherwise.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetEncryptionSettingsDefinitionID() string {
	if o == nil || isNil(o.EncryptionSettingsDefinitionID) {
		var ret string
		return ret
	}
	return *o.EncryptionSettingsDefinitionID
}

// GetEncryptionSettingsDefinitionIDOk returns a tuple with the EncryptionSettingsDefinitionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetEncryptionSettingsDefinitionIDOk() (*string, bool) {
	if o == nil || isNil(o.EncryptionSettingsDefinitionID) {
    return nil, false
	}
	return o.EncryptionSettingsDefinitionID, true
}

// HasEncryptionSettingsDefinitionID returns a boolean if a field has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) HasEncryptionSettingsDefinitionID() bool {
	if o != nil && !isNil(o.EncryptionSettingsDefinitionID) {
		return true
	}

	return false
}

// SetEncryptionSettingsDefinitionID gets a reference to the given string and assigns it to the EncryptionSettingsDefinitionID field.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) SetEncryptionSettingsDefinitionID(v string) {
	o.EncryptionSettingsDefinitionID = &v
}

// GetAppend returns the Append field value if set, zero value otherwise.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetAppend() bool {
	if o == nil || isNil(o.Append) {
		var ret bool
		return ret
	}
	return *o.Append
}

// GetAppendOk returns a tuple with the Append field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetAppendOk() (*bool, bool) {
	if o == nil || isNil(o.Append) {
    return nil, false
	}
	return o.Append, true
}

// HasAppend returns a boolean if a field has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) HasAppend() bool {
	if o != nil && !isNil(o.Append) {
		return true
	}

	return false
}

// SetAppend gets a reference to the given bool and assigns it to the Append field.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) SetAppend(v bool) {
	o.Append = &v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetExtensionClassOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetExtensionArgument() []string {
	if o == nil || isNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || isNil(o.ExtensionArgument) {
    return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) HasExtensionArgument() bool {
	if o != nil && !isNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetAsynchronous returns the Asynchronous field value
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetAsynchronous() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Asynchronous
}

// GetAsynchronousOk returns a tuple with the Asynchronous field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetAsynchronousOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Asynchronous, true
}

// SetAsynchronous sets field value
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) SetAsynchronous(v bool) {
	o.Asynchronous = v
}

// GetAutoFlush returns the AutoFlush field value if set, zero value otherwise.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetAutoFlush() bool {
	if o == nil || isNil(o.AutoFlush) {
		var ret bool
		return ret
	}
	return *o.AutoFlush
}

// GetAutoFlushOk returns a tuple with the AutoFlush field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetAutoFlushOk() (*bool, bool) {
	if o == nil || isNil(o.AutoFlush) {
    return nil, false
	}
	return o.AutoFlush, true
}

// HasAutoFlush returns a boolean if a field has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) HasAutoFlush() bool {
	if o != nil && !isNil(o.AutoFlush) {
		return true
	}

	return false
}

// SetAutoFlush gets a reference to the given bool and assigns it to the AutoFlush field.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) SetAutoFlush(v bool) {
	o.AutoFlush = &v
}

// GetBufferSize returns the BufferSize field value if set, zero value otherwise.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetBufferSize() string {
	if o == nil || isNil(o.BufferSize) {
		var ret string
		return ret
	}
	return *o.BufferSize
}

// GetBufferSizeOk returns a tuple with the BufferSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetBufferSizeOk() (*string, bool) {
	if o == nil || isNil(o.BufferSize) {
    return nil, false
	}
	return o.BufferSize, true
}

// HasBufferSize returns a boolean if a field has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) HasBufferSize() bool {
	if o != nil && !isNil(o.BufferSize) {
		return true
	}

	return false
}

// SetBufferSize gets a reference to the given string and assigns it to the BufferSize field.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) SetBufferSize(v string) {
	o.BufferSize = &v
}

// GetQueueSize returns the QueueSize field value if set, zero value otherwise.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetQueueSize() int32 {
	if o == nil || isNil(o.QueueSize) {
		var ret int32
		return ret
	}
	return *o.QueueSize
}

// GetQueueSizeOk returns a tuple with the QueueSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetQueueSizeOk() (*int32, bool) {
	if o == nil || isNil(o.QueueSize) {
    return nil, false
	}
	return o.QueueSize, true
}

// HasQueueSize returns a boolean if a field has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) HasQueueSize() bool {
	if o != nil && !isNil(o.QueueSize) {
		return true
	}

	return false
}

// SetQueueSize gets a reference to the given int32 and assigns it to the QueueSize field.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) SetQueueSize(v int32) {
	o.QueueSize = &v
}

// GetTimeInterval returns the TimeInterval field value if set, zero value otherwise.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetTimeInterval() string {
	if o == nil || isNil(o.TimeInterval) {
		var ret string
		return ret
	}
	return *o.TimeInterval
}

// GetTimeIntervalOk returns a tuple with the TimeInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetTimeIntervalOk() (*string, bool) {
	if o == nil || isNil(o.TimeInterval) {
    return nil, false
	}
	return o.TimeInterval, true
}

// HasTimeInterval returns a boolean if a field has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) HasTimeInterval() bool {
	if o != nil && !isNil(o.TimeInterval) {
		return true
	}

	return false
}

// SetTimeInterval gets a reference to the given string and assigns it to the TimeInterval field.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) SetTimeInterval(v string) {
	o.TimeInterval = &v
}

// GetDefaultSeverity returns the DefaultSeverity field value if set, zero value otherwise.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetDefaultSeverity() []EnumlogPublisherDefaultSeverityProp {
	if o == nil || isNil(o.DefaultSeverity) {
		var ret []EnumlogPublisherDefaultSeverityProp
		return ret
	}
	return o.DefaultSeverity
}

// GetDefaultSeverityOk returns a tuple with the DefaultSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetDefaultSeverityOk() ([]EnumlogPublisherDefaultSeverityProp, bool) {
	if o == nil || isNil(o.DefaultSeverity) {
    return nil, false
	}
	return o.DefaultSeverity, true
}

// HasDefaultSeverity returns a boolean if a field has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) HasDefaultSeverity() bool {
	if o != nil && !isNil(o.DefaultSeverity) {
		return true
	}

	return false
}

// SetDefaultSeverity gets a reference to the given []EnumlogPublisherDefaultSeverityProp and assigns it to the DefaultSeverity field.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) SetDefaultSeverity(v []EnumlogPublisherDefaultSeverityProp) {
	o.DefaultSeverity = v
}

// GetOverrideSeverity returns the OverrideSeverity field value if set, zero value otherwise.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetOverrideSeverity() []string {
	if o == nil || isNil(o.OverrideSeverity) {
		var ret []string
		return ret
	}
	return o.OverrideSeverity
}

// GetOverrideSeverityOk returns a tuple with the OverrideSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetOverrideSeverityOk() ([]string, bool) {
	if o == nil || isNil(o.OverrideSeverity) {
    return nil, false
	}
	return o.OverrideSeverity, true
}

// HasOverrideSeverity returns a boolean if a field has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) HasOverrideSeverity() bool {
	if o != nil && !isNil(o.OverrideSeverity) {
		return true
	}

	return false
}

// SetOverrideSeverity gets a reference to the given []string and assigns it to the OverrideSeverity field.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) SetOverrideSeverity(v []string) {
	o.OverrideSeverity = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetLoggingErrorBehavior returns the LoggingErrorBehavior field value if set, zero value otherwise.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp {
	if o == nil || isNil(o.LoggingErrorBehavior) {
		var ret EnumlogPublisherLoggingErrorBehaviorProp
		return ret
	}
	return *o.LoggingErrorBehavior
}

// GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool) {
	if o == nil || isNil(o.LoggingErrorBehavior) {
    return nil, false
	}
	return o.LoggingErrorBehavior, true
}

// HasLoggingErrorBehavior returns a boolean if a field has been set.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) HasLoggingErrorBehavior() bool {
	if o != nil && !isNil(o.LoggingErrorBehavior) {
		return true
	}

	return false
}

// SetLoggingErrorBehavior gets a reference to the given EnumlogPublisherLoggingErrorBehaviorProp and assigns it to the LoggingErrorBehavior field.
func (o *AddThirdPartyFileBasedErrorLogPublisherRequest) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp) {
	o.LoggingErrorBehavior = &v
}

func (o AddThirdPartyFileBasedErrorLogPublisherRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["publisherName"] = o.PublisherName
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
		toSerialize["extensionClass"] = o.ExtensionClass
	}
	if !isNil(o.ExtensionArgument) {
		toSerialize["extensionArgument"] = o.ExtensionArgument
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
	return json.Marshal(toSerialize)
}

type NullableAddThirdPartyFileBasedErrorLogPublisherRequest struct {
	value *AddThirdPartyFileBasedErrorLogPublisherRequest
	isSet bool
}

func (v NullableAddThirdPartyFileBasedErrorLogPublisherRequest) Get() *AddThirdPartyFileBasedErrorLogPublisherRequest {
	return v.value
}

func (v *NullableAddThirdPartyFileBasedErrorLogPublisherRequest) Set(val *AddThirdPartyFileBasedErrorLogPublisherRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddThirdPartyFileBasedErrorLogPublisherRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddThirdPartyFileBasedErrorLogPublisherRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddThirdPartyFileBasedErrorLogPublisherRequest(val *AddThirdPartyFileBasedErrorLogPublisherRequest) *NullableAddThirdPartyFileBasedErrorLogPublisherRequest {
	return &NullableAddThirdPartyFileBasedErrorLogPublisherRequest{value: val, isSet: true}
}

func (v NullableAddThirdPartyFileBasedErrorLogPublisherRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddThirdPartyFileBasedErrorLogPublisherRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


