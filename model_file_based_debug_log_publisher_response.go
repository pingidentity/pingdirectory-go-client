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

// FileBasedDebugLogPublisherResponse struct for FileBasedDebugLogPublisherResponse
type FileBasedDebugLogPublisherResponse struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Log Publisher
	Id string `json:"id"`
	Schemas []EnumfileBasedDebugLogPublisherSchemaUrn `json:"schemas"`
	// The file name to use for the log files generated by the File Based Debug Log Publisher. The path to the file can be specified either as relative to the server root or as an absolute path.
	LogFile string `json:"logFile"`
	// The UNIX permissions of the log files created by this File Based Debug Log Publisher.
	LogFilePermissions string `json:"logFilePermissions"`
	// The rotation policy to use for the File Based Debug Log Publisher .
	RotationPolicy []string `json:"rotationPolicy"`
	// A listener that should be notified whenever a log file is rotated out of service.
	RotationListener []string `json:"rotationListener,omitempty"`
	// The retention policy to use for the File Based Debug Log Publisher .
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
	// Indicates whether the File Based Debug Log Publisher will publish records asynchronously.
	Asynchronous bool `json:"asynchronous"`
	// Specifies whether to flush the writer after every log record.
	AutoFlush *bool `json:"autoFlush,omitempty"`
	// Specifies the log file buffer size.
	BufferSize *string `json:"bufferSize,omitempty"`
	// The maximum number of log records that can be stored in the asynchronous queue.
	QueueSize *int32 `json:"queueSize,omitempty"`
	// Specifies the interval at which to check whether the log files need to be rotated.
	TimeInterval *string `json:"timeInterval,omitempty"`
	TimestampPrecision *EnumlogPublisherTimestampPrecisionProp `json:"timestampPrecision,omitempty"`
	DefaultDebugLevel EnumlogPublisherDefaultDebugLevelProp `json:"defaultDebugLevel"`
	DefaultDebugCategory []EnumlogPublisherDefaultDebugCategoryProp `json:"defaultDebugCategory,omitempty"`
	// Indicates whether to include method arguments in debug messages logged by default.
	DefaultOmitMethodEntryArguments *bool `json:"defaultOmitMethodEntryArguments,omitempty"`
	// Indicates whether to include the return value in debug messages logged by default.
	DefaultOmitMethodReturnValue *bool `json:"defaultOmitMethodReturnValue,omitempty"`
	// Indicates whether to include the cause of exceptions in exception thrown and caught messages logged by default.
	DefaultIncludeThrowableCause *bool `json:"defaultIncludeThrowableCause,omitempty"`
	// Indicates the number of stack frames to include in the stack trace for method entry and exception thrown messages.
	DefaultThrowableStackFrames *int32 `json:"defaultThrowableStackFrames,omitempty"`
	// A description for this Log Publisher
	Description *string `json:"description,omitempty"`
	// Indicates whether the Log Publisher is enabled for use.
	Enabled bool `json:"enabled"`
	LoggingErrorBehavior *EnumlogPublisherLoggingErrorBehaviorProp `json:"loggingErrorBehavior,omitempty"`
}

// NewFileBasedDebugLogPublisherResponse instantiates a new FileBasedDebugLogPublisherResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileBasedDebugLogPublisherResponse(id string, schemas []EnumfileBasedDebugLogPublisherSchemaUrn, logFile string, logFilePermissions string, rotationPolicy []string, retentionPolicy []string, asynchronous bool, defaultDebugLevel EnumlogPublisherDefaultDebugLevelProp, enabled bool) *FileBasedDebugLogPublisherResponse {
	this := FileBasedDebugLogPublisherResponse{}
	this.Id = id
	this.Schemas = schemas
	this.LogFile = logFile
	this.LogFilePermissions = logFilePermissions
	this.RotationPolicy = rotationPolicy
	this.RetentionPolicy = retentionPolicy
	this.Asynchronous = asynchronous
	this.DefaultDebugLevel = defaultDebugLevel
	this.Enabled = enabled
	return &this
}

// NewFileBasedDebugLogPublisherResponseWithDefaults instantiates a new FileBasedDebugLogPublisherResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileBasedDebugLogPublisherResponseWithDefaults() *FileBasedDebugLogPublisherResponse {
	this := FileBasedDebugLogPublisherResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *FileBasedDebugLogPublisherResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedDebugLogPublisherResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *FileBasedDebugLogPublisherResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *FileBasedDebugLogPublisherResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *FileBasedDebugLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedDebugLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
    return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *FileBasedDebugLogPublisherResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *FileBasedDebugLogPublisherResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *FileBasedDebugLogPublisherResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FileBasedDebugLogPublisherResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FileBasedDebugLogPublisherResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *FileBasedDebugLogPublisherResponse) GetSchemas() []EnumfileBasedDebugLogPublisherSchemaUrn {
	if o == nil {
		var ret []EnumfileBasedDebugLogPublisherSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *FileBasedDebugLogPublisherResponse) GetSchemasOk() ([]EnumfileBasedDebugLogPublisherSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *FileBasedDebugLogPublisherResponse) SetSchemas(v []EnumfileBasedDebugLogPublisherSchemaUrn) {
	o.Schemas = v
}

// GetLogFile returns the LogFile field value
func (o *FileBasedDebugLogPublisherResponse) GetLogFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogFile
}

// GetLogFileOk returns a tuple with the LogFile field value
// and a boolean to check if the value has been set.
func (o *FileBasedDebugLogPublisherResponse) GetLogFileOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LogFile, true
}

// SetLogFile sets field value
func (o *FileBasedDebugLogPublisherResponse) SetLogFile(v string) {
	o.LogFile = v
}

// GetLogFilePermissions returns the LogFilePermissions field value
func (o *FileBasedDebugLogPublisherResponse) GetLogFilePermissions() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogFilePermissions
}

// GetLogFilePermissionsOk returns a tuple with the LogFilePermissions field value
// and a boolean to check if the value has been set.
func (o *FileBasedDebugLogPublisherResponse) GetLogFilePermissionsOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LogFilePermissions, true
}

// SetLogFilePermissions sets field value
func (o *FileBasedDebugLogPublisherResponse) SetLogFilePermissions(v string) {
	o.LogFilePermissions = v
}

// GetRotationPolicy returns the RotationPolicy field value
func (o *FileBasedDebugLogPublisherResponse) GetRotationPolicy() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RotationPolicy
}

// GetRotationPolicyOk returns a tuple with the RotationPolicy field value
// and a boolean to check if the value has been set.
func (o *FileBasedDebugLogPublisherResponse) GetRotationPolicyOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RotationPolicy, true
}

// SetRotationPolicy sets field value
func (o *FileBasedDebugLogPublisherResponse) SetRotationPolicy(v []string) {
	o.RotationPolicy = v
}

// GetRotationListener returns the RotationListener field value if set, zero value otherwise.
func (o *FileBasedDebugLogPublisherResponse) GetRotationListener() []string {
	if o == nil || isNil(o.RotationListener) {
		var ret []string
		return ret
	}
	return o.RotationListener
}

// GetRotationListenerOk returns a tuple with the RotationListener field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedDebugLogPublisherResponse) GetRotationListenerOk() ([]string, bool) {
	if o == nil || isNil(o.RotationListener) {
    return nil, false
	}
	return o.RotationListener, true
}

// HasRotationListener returns a boolean if a field has been set.
func (o *FileBasedDebugLogPublisherResponse) HasRotationListener() bool {
	if o != nil && !isNil(o.RotationListener) {
		return true
	}

	return false
}

// SetRotationListener gets a reference to the given []string and assigns it to the RotationListener field.
func (o *FileBasedDebugLogPublisherResponse) SetRotationListener(v []string) {
	o.RotationListener = v
}

// GetRetentionPolicy returns the RetentionPolicy field value
func (o *FileBasedDebugLogPublisherResponse) GetRetentionPolicy() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RetentionPolicy
}

// GetRetentionPolicyOk returns a tuple with the RetentionPolicy field value
// and a boolean to check if the value has been set.
func (o *FileBasedDebugLogPublisherResponse) GetRetentionPolicyOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RetentionPolicy, true
}

// SetRetentionPolicy sets field value
func (o *FileBasedDebugLogPublisherResponse) SetRetentionPolicy(v []string) {
	o.RetentionPolicy = v
}

// GetCompressionMechanism returns the CompressionMechanism field value if set, zero value otherwise.
func (o *FileBasedDebugLogPublisherResponse) GetCompressionMechanism() EnumlogPublisherCompressionMechanismProp {
	if o == nil || isNil(o.CompressionMechanism) {
		var ret EnumlogPublisherCompressionMechanismProp
		return ret
	}
	return *o.CompressionMechanism
}

// GetCompressionMechanismOk returns a tuple with the CompressionMechanism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedDebugLogPublisherResponse) GetCompressionMechanismOk() (*EnumlogPublisherCompressionMechanismProp, bool) {
	if o == nil || isNil(o.CompressionMechanism) {
    return nil, false
	}
	return o.CompressionMechanism, true
}

// HasCompressionMechanism returns a boolean if a field has been set.
func (o *FileBasedDebugLogPublisherResponse) HasCompressionMechanism() bool {
	if o != nil && !isNil(o.CompressionMechanism) {
		return true
	}

	return false
}

// SetCompressionMechanism gets a reference to the given EnumlogPublisherCompressionMechanismProp and assigns it to the CompressionMechanism field.
func (o *FileBasedDebugLogPublisherResponse) SetCompressionMechanism(v EnumlogPublisherCompressionMechanismProp) {
	o.CompressionMechanism = &v
}

// GetSignLog returns the SignLog field value if set, zero value otherwise.
func (o *FileBasedDebugLogPublisherResponse) GetSignLog() bool {
	if o == nil || isNil(o.SignLog) {
		var ret bool
		return ret
	}
	return *o.SignLog
}

// GetSignLogOk returns a tuple with the SignLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedDebugLogPublisherResponse) GetSignLogOk() (*bool, bool) {
	if o == nil || isNil(o.SignLog) {
    return nil, false
	}
	return o.SignLog, true
}

// HasSignLog returns a boolean if a field has been set.
func (o *FileBasedDebugLogPublisherResponse) HasSignLog() bool {
	if o != nil && !isNil(o.SignLog) {
		return true
	}

	return false
}

// SetSignLog gets a reference to the given bool and assigns it to the SignLog field.
func (o *FileBasedDebugLogPublisherResponse) SetSignLog(v bool) {
	o.SignLog = &v
}

// GetEncryptLog returns the EncryptLog field value if set, zero value otherwise.
func (o *FileBasedDebugLogPublisherResponse) GetEncryptLog() bool {
	if o == nil || isNil(o.EncryptLog) {
		var ret bool
		return ret
	}
	return *o.EncryptLog
}

// GetEncryptLogOk returns a tuple with the EncryptLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedDebugLogPublisherResponse) GetEncryptLogOk() (*bool, bool) {
	if o == nil || isNil(o.EncryptLog) {
    return nil, false
	}
	return o.EncryptLog, true
}

// HasEncryptLog returns a boolean if a field has been set.
func (o *FileBasedDebugLogPublisherResponse) HasEncryptLog() bool {
	if o != nil && !isNil(o.EncryptLog) {
		return true
	}

	return false
}

// SetEncryptLog gets a reference to the given bool and assigns it to the EncryptLog field.
func (o *FileBasedDebugLogPublisherResponse) SetEncryptLog(v bool) {
	o.EncryptLog = &v
}

// GetEncryptionSettingsDefinitionID returns the EncryptionSettingsDefinitionID field value if set, zero value otherwise.
func (o *FileBasedDebugLogPublisherResponse) GetEncryptionSettingsDefinitionID() string {
	if o == nil || isNil(o.EncryptionSettingsDefinitionID) {
		var ret string
		return ret
	}
	return *o.EncryptionSettingsDefinitionID
}

// GetEncryptionSettingsDefinitionIDOk returns a tuple with the EncryptionSettingsDefinitionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedDebugLogPublisherResponse) GetEncryptionSettingsDefinitionIDOk() (*string, bool) {
	if o == nil || isNil(o.EncryptionSettingsDefinitionID) {
    return nil, false
	}
	return o.EncryptionSettingsDefinitionID, true
}

// HasEncryptionSettingsDefinitionID returns a boolean if a field has been set.
func (o *FileBasedDebugLogPublisherResponse) HasEncryptionSettingsDefinitionID() bool {
	if o != nil && !isNil(o.EncryptionSettingsDefinitionID) {
		return true
	}

	return false
}

// SetEncryptionSettingsDefinitionID gets a reference to the given string and assigns it to the EncryptionSettingsDefinitionID field.
func (o *FileBasedDebugLogPublisherResponse) SetEncryptionSettingsDefinitionID(v string) {
	o.EncryptionSettingsDefinitionID = &v
}

// GetAppend returns the Append field value if set, zero value otherwise.
func (o *FileBasedDebugLogPublisherResponse) GetAppend() bool {
	if o == nil || isNil(o.Append) {
		var ret bool
		return ret
	}
	return *o.Append
}

// GetAppendOk returns a tuple with the Append field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedDebugLogPublisherResponse) GetAppendOk() (*bool, bool) {
	if o == nil || isNil(o.Append) {
    return nil, false
	}
	return o.Append, true
}

// HasAppend returns a boolean if a field has been set.
func (o *FileBasedDebugLogPublisherResponse) HasAppend() bool {
	if o != nil && !isNil(o.Append) {
		return true
	}

	return false
}

// SetAppend gets a reference to the given bool and assigns it to the Append field.
func (o *FileBasedDebugLogPublisherResponse) SetAppend(v bool) {
	o.Append = &v
}

// GetAsynchronous returns the Asynchronous field value
func (o *FileBasedDebugLogPublisherResponse) GetAsynchronous() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Asynchronous
}

// GetAsynchronousOk returns a tuple with the Asynchronous field value
// and a boolean to check if the value has been set.
func (o *FileBasedDebugLogPublisherResponse) GetAsynchronousOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Asynchronous, true
}

// SetAsynchronous sets field value
func (o *FileBasedDebugLogPublisherResponse) SetAsynchronous(v bool) {
	o.Asynchronous = v
}

// GetAutoFlush returns the AutoFlush field value if set, zero value otherwise.
func (o *FileBasedDebugLogPublisherResponse) GetAutoFlush() bool {
	if o == nil || isNil(o.AutoFlush) {
		var ret bool
		return ret
	}
	return *o.AutoFlush
}

// GetAutoFlushOk returns a tuple with the AutoFlush field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedDebugLogPublisherResponse) GetAutoFlushOk() (*bool, bool) {
	if o == nil || isNil(o.AutoFlush) {
    return nil, false
	}
	return o.AutoFlush, true
}

// HasAutoFlush returns a boolean if a field has been set.
func (o *FileBasedDebugLogPublisherResponse) HasAutoFlush() bool {
	if o != nil && !isNil(o.AutoFlush) {
		return true
	}

	return false
}

// SetAutoFlush gets a reference to the given bool and assigns it to the AutoFlush field.
func (o *FileBasedDebugLogPublisherResponse) SetAutoFlush(v bool) {
	o.AutoFlush = &v
}

// GetBufferSize returns the BufferSize field value if set, zero value otherwise.
func (o *FileBasedDebugLogPublisherResponse) GetBufferSize() string {
	if o == nil || isNil(o.BufferSize) {
		var ret string
		return ret
	}
	return *o.BufferSize
}

// GetBufferSizeOk returns a tuple with the BufferSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedDebugLogPublisherResponse) GetBufferSizeOk() (*string, bool) {
	if o == nil || isNil(o.BufferSize) {
    return nil, false
	}
	return o.BufferSize, true
}

// HasBufferSize returns a boolean if a field has been set.
func (o *FileBasedDebugLogPublisherResponse) HasBufferSize() bool {
	if o != nil && !isNil(o.BufferSize) {
		return true
	}

	return false
}

// SetBufferSize gets a reference to the given string and assigns it to the BufferSize field.
func (o *FileBasedDebugLogPublisherResponse) SetBufferSize(v string) {
	o.BufferSize = &v
}

// GetQueueSize returns the QueueSize field value if set, zero value otherwise.
func (o *FileBasedDebugLogPublisherResponse) GetQueueSize() int32 {
	if o == nil || isNil(o.QueueSize) {
		var ret int32
		return ret
	}
	return *o.QueueSize
}

// GetQueueSizeOk returns a tuple with the QueueSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedDebugLogPublisherResponse) GetQueueSizeOk() (*int32, bool) {
	if o == nil || isNil(o.QueueSize) {
    return nil, false
	}
	return o.QueueSize, true
}

// HasQueueSize returns a boolean if a field has been set.
func (o *FileBasedDebugLogPublisherResponse) HasQueueSize() bool {
	if o != nil && !isNil(o.QueueSize) {
		return true
	}

	return false
}

// SetQueueSize gets a reference to the given int32 and assigns it to the QueueSize field.
func (o *FileBasedDebugLogPublisherResponse) SetQueueSize(v int32) {
	o.QueueSize = &v
}

// GetTimeInterval returns the TimeInterval field value if set, zero value otherwise.
func (o *FileBasedDebugLogPublisherResponse) GetTimeInterval() string {
	if o == nil || isNil(o.TimeInterval) {
		var ret string
		return ret
	}
	return *o.TimeInterval
}

// GetTimeIntervalOk returns a tuple with the TimeInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedDebugLogPublisherResponse) GetTimeIntervalOk() (*string, bool) {
	if o == nil || isNil(o.TimeInterval) {
    return nil, false
	}
	return o.TimeInterval, true
}

// HasTimeInterval returns a boolean if a field has been set.
func (o *FileBasedDebugLogPublisherResponse) HasTimeInterval() bool {
	if o != nil && !isNil(o.TimeInterval) {
		return true
	}

	return false
}

// SetTimeInterval gets a reference to the given string and assigns it to the TimeInterval field.
func (o *FileBasedDebugLogPublisherResponse) SetTimeInterval(v string) {
	o.TimeInterval = &v
}

// GetTimestampPrecision returns the TimestampPrecision field value if set, zero value otherwise.
func (o *FileBasedDebugLogPublisherResponse) GetTimestampPrecision() EnumlogPublisherTimestampPrecisionProp {
	if o == nil || isNil(o.TimestampPrecision) {
		var ret EnumlogPublisherTimestampPrecisionProp
		return ret
	}
	return *o.TimestampPrecision
}

// GetTimestampPrecisionOk returns a tuple with the TimestampPrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedDebugLogPublisherResponse) GetTimestampPrecisionOk() (*EnumlogPublisherTimestampPrecisionProp, bool) {
	if o == nil || isNil(o.TimestampPrecision) {
    return nil, false
	}
	return o.TimestampPrecision, true
}

// HasTimestampPrecision returns a boolean if a field has been set.
func (o *FileBasedDebugLogPublisherResponse) HasTimestampPrecision() bool {
	if o != nil && !isNil(o.TimestampPrecision) {
		return true
	}

	return false
}

// SetTimestampPrecision gets a reference to the given EnumlogPublisherTimestampPrecisionProp and assigns it to the TimestampPrecision field.
func (o *FileBasedDebugLogPublisherResponse) SetTimestampPrecision(v EnumlogPublisherTimestampPrecisionProp) {
	o.TimestampPrecision = &v
}

// GetDefaultDebugLevel returns the DefaultDebugLevel field value
func (o *FileBasedDebugLogPublisherResponse) GetDefaultDebugLevel() EnumlogPublisherDefaultDebugLevelProp {
	if o == nil {
		var ret EnumlogPublisherDefaultDebugLevelProp
		return ret
	}

	return o.DefaultDebugLevel
}

// GetDefaultDebugLevelOk returns a tuple with the DefaultDebugLevel field value
// and a boolean to check if the value has been set.
func (o *FileBasedDebugLogPublisherResponse) GetDefaultDebugLevelOk() (*EnumlogPublisherDefaultDebugLevelProp, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DefaultDebugLevel, true
}

// SetDefaultDebugLevel sets field value
func (o *FileBasedDebugLogPublisherResponse) SetDefaultDebugLevel(v EnumlogPublisherDefaultDebugLevelProp) {
	o.DefaultDebugLevel = v
}

// GetDefaultDebugCategory returns the DefaultDebugCategory field value if set, zero value otherwise.
func (o *FileBasedDebugLogPublisherResponse) GetDefaultDebugCategory() []EnumlogPublisherDefaultDebugCategoryProp {
	if o == nil || isNil(o.DefaultDebugCategory) {
		var ret []EnumlogPublisherDefaultDebugCategoryProp
		return ret
	}
	return o.DefaultDebugCategory
}

// GetDefaultDebugCategoryOk returns a tuple with the DefaultDebugCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedDebugLogPublisherResponse) GetDefaultDebugCategoryOk() ([]EnumlogPublisherDefaultDebugCategoryProp, bool) {
	if o == nil || isNil(o.DefaultDebugCategory) {
    return nil, false
	}
	return o.DefaultDebugCategory, true
}

// HasDefaultDebugCategory returns a boolean if a field has been set.
func (o *FileBasedDebugLogPublisherResponse) HasDefaultDebugCategory() bool {
	if o != nil && !isNil(o.DefaultDebugCategory) {
		return true
	}

	return false
}

// SetDefaultDebugCategory gets a reference to the given []EnumlogPublisherDefaultDebugCategoryProp and assigns it to the DefaultDebugCategory field.
func (o *FileBasedDebugLogPublisherResponse) SetDefaultDebugCategory(v []EnumlogPublisherDefaultDebugCategoryProp) {
	o.DefaultDebugCategory = v
}

// GetDefaultOmitMethodEntryArguments returns the DefaultOmitMethodEntryArguments field value if set, zero value otherwise.
func (o *FileBasedDebugLogPublisherResponse) GetDefaultOmitMethodEntryArguments() bool {
	if o == nil || isNil(o.DefaultOmitMethodEntryArguments) {
		var ret bool
		return ret
	}
	return *o.DefaultOmitMethodEntryArguments
}

// GetDefaultOmitMethodEntryArgumentsOk returns a tuple with the DefaultOmitMethodEntryArguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedDebugLogPublisherResponse) GetDefaultOmitMethodEntryArgumentsOk() (*bool, bool) {
	if o == nil || isNil(o.DefaultOmitMethodEntryArguments) {
    return nil, false
	}
	return o.DefaultOmitMethodEntryArguments, true
}

// HasDefaultOmitMethodEntryArguments returns a boolean if a field has been set.
func (o *FileBasedDebugLogPublisherResponse) HasDefaultOmitMethodEntryArguments() bool {
	if o != nil && !isNil(o.DefaultOmitMethodEntryArguments) {
		return true
	}

	return false
}

// SetDefaultOmitMethodEntryArguments gets a reference to the given bool and assigns it to the DefaultOmitMethodEntryArguments field.
func (o *FileBasedDebugLogPublisherResponse) SetDefaultOmitMethodEntryArguments(v bool) {
	o.DefaultOmitMethodEntryArguments = &v
}

// GetDefaultOmitMethodReturnValue returns the DefaultOmitMethodReturnValue field value if set, zero value otherwise.
func (o *FileBasedDebugLogPublisherResponse) GetDefaultOmitMethodReturnValue() bool {
	if o == nil || isNil(o.DefaultOmitMethodReturnValue) {
		var ret bool
		return ret
	}
	return *o.DefaultOmitMethodReturnValue
}

// GetDefaultOmitMethodReturnValueOk returns a tuple with the DefaultOmitMethodReturnValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedDebugLogPublisherResponse) GetDefaultOmitMethodReturnValueOk() (*bool, bool) {
	if o == nil || isNil(o.DefaultOmitMethodReturnValue) {
    return nil, false
	}
	return o.DefaultOmitMethodReturnValue, true
}

// HasDefaultOmitMethodReturnValue returns a boolean if a field has been set.
func (o *FileBasedDebugLogPublisherResponse) HasDefaultOmitMethodReturnValue() bool {
	if o != nil && !isNil(o.DefaultOmitMethodReturnValue) {
		return true
	}

	return false
}

// SetDefaultOmitMethodReturnValue gets a reference to the given bool and assigns it to the DefaultOmitMethodReturnValue field.
func (o *FileBasedDebugLogPublisherResponse) SetDefaultOmitMethodReturnValue(v bool) {
	o.DefaultOmitMethodReturnValue = &v
}

// GetDefaultIncludeThrowableCause returns the DefaultIncludeThrowableCause field value if set, zero value otherwise.
func (o *FileBasedDebugLogPublisherResponse) GetDefaultIncludeThrowableCause() bool {
	if o == nil || isNil(o.DefaultIncludeThrowableCause) {
		var ret bool
		return ret
	}
	return *o.DefaultIncludeThrowableCause
}

// GetDefaultIncludeThrowableCauseOk returns a tuple with the DefaultIncludeThrowableCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedDebugLogPublisherResponse) GetDefaultIncludeThrowableCauseOk() (*bool, bool) {
	if o == nil || isNil(o.DefaultIncludeThrowableCause) {
    return nil, false
	}
	return o.DefaultIncludeThrowableCause, true
}

// HasDefaultIncludeThrowableCause returns a boolean if a field has been set.
func (o *FileBasedDebugLogPublisherResponse) HasDefaultIncludeThrowableCause() bool {
	if o != nil && !isNil(o.DefaultIncludeThrowableCause) {
		return true
	}

	return false
}

// SetDefaultIncludeThrowableCause gets a reference to the given bool and assigns it to the DefaultIncludeThrowableCause field.
func (o *FileBasedDebugLogPublisherResponse) SetDefaultIncludeThrowableCause(v bool) {
	o.DefaultIncludeThrowableCause = &v
}

// GetDefaultThrowableStackFrames returns the DefaultThrowableStackFrames field value if set, zero value otherwise.
func (o *FileBasedDebugLogPublisherResponse) GetDefaultThrowableStackFrames() int32 {
	if o == nil || isNil(o.DefaultThrowableStackFrames) {
		var ret int32
		return ret
	}
	return *o.DefaultThrowableStackFrames
}

// GetDefaultThrowableStackFramesOk returns a tuple with the DefaultThrowableStackFrames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedDebugLogPublisherResponse) GetDefaultThrowableStackFramesOk() (*int32, bool) {
	if o == nil || isNil(o.DefaultThrowableStackFrames) {
    return nil, false
	}
	return o.DefaultThrowableStackFrames, true
}

// HasDefaultThrowableStackFrames returns a boolean if a field has been set.
func (o *FileBasedDebugLogPublisherResponse) HasDefaultThrowableStackFrames() bool {
	if o != nil && !isNil(o.DefaultThrowableStackFrames) {
		return true
	}

	return false
}

// SetDefaultThrowableStackFrames gets a reference to the given int32 and assigns it to the DefaultThrowableStackFrames field.
func (o *FileBasedDebugLogPublisherResponse) SetDefaultThrowableStackFrames(v int32) {
	o.DefaultThrowableStackFrames = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FileBasedDebugLogPublisherResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedDebugLogPublisherResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FileBasedDebugLogPublisherResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FileBasedDebugLogPublisherResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *FileBasedDebugLogPublisherResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *FileBasedDebugLogPublisherResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *FileBasedDebugLogPublisherResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetLoggingErrorBehavior returns the LoggingErrorBehavior field value if set, zero value otherwise.
func (o *FileBasedDebugLogPublisherResponse) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp {
	if o == nil || isNil(o.LoggingErrorBehavior) {
		var ret EnumlogPublisherLoggingErrorBehaviorProp
		return ret
	}
	return *o.LoggingErrorBehavior
}

// GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedDebugLogPublisherResponse) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool) {
	if o == nil || isNil(o.LoggingErrorBehavior) {
    return nil, false
	}
	return o.LoggingErrorBehavior, true
}

// HasLoggingErrorBehavior returns a boolean if a field has been set.
func (o *FileBasedDebugLogPublisherResponse) HasLoggingErrorBehavior() bool {
	if o != nil && !isNil(o.LoggingErrorBehavior) {
		return true
	}

	return false
}

// SetLoggingErrorBehavior gets a reference to the given EnumlogPublisherLoggingErrorBehaviorProp and assigns it to the LoggingErrorBehavior field.
func (o *FileBasedDebugLogPublisherResponse) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp) {
	o.LoggingErrorBehavior = &v
}

func (o FileBasedDebugLogPublisherResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
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
	if !isNil(o.TimestampPrecision) {
		toSerialize["timestampPrecision"] = o.TimestampPrecision
	}
	if true {
		toSerialize["defaultDebugLevel"] = o.DefaultDebugLevel
	}
	if !isNil(o.DefaultDebugCategory) {
		toSerialize["defaultDebugCategory"] = o.DefaultDebugCategory
	}
	if !isNil(o.DefaultOmitMethodEntryArguments) {
		toSerialize["defaultOmitMethodEntryArguments"] = o.DefaultOmitMethodEntryArguments
	}
	if !isNil(o.DefaultOmitMethodReturnValue) {
		toSerialize["defaultOmitMethodReturnValue"] = o.DefaultOmitMethodReturnValue
	}
	if !isNil(o.DefaultIncludeThrowableCause) {
		toSerialize["defaultIncludeThrowableCause"] = o.DefaultIncludeThrowableCause
	}
	if !isNil(o.DefaultThrowableStackFrames) {
		toSerialize["defaultThrowableStackFrames"] = o.DefaultThrowableStackFrames
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

type NullableFileBasedDebugLogPublisherResponse struct {
	value *FileBasedDebugLogPublisherResponse
	isSet bool
}

func (v NullableFileBasedDebugLogPublisherResponse) Get() *FileBasedDebugLogPublisherResponse {
	return v.value
}

func (v *NullableFileBasedDebugLogPublisherResponse) Set(val *FileBasedDebugLogPublisherResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFileBasedDebugLogPublisherResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFileBasedDebugLogPublisherResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileBasedDebugLogPublisherResponse(val *FileBasedDebugLogPublisherResponse) *NullableFileBasedDebugLogPublisherResponse {
	return &NullableFileBasedDebugLogPublisherResponse{value: val, isSet: true}
}

func (v NullableFileBasedDebugLogPublisherResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileBasedDebugLogPublisherResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


