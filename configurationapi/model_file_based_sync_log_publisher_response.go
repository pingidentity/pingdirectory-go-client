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

// checks if the FileBasedSyncLogPublisherResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileBasedSyncLogPublisherResponse{}

// FileBasedSyncLogPublisherResponse struct for FileBasedSyncLogPublisherResponse
type FileBasedSyncLogPublisherResponse struct {
	Schemas []EnumfileBasedSyncLogPublisherSchemaUrn `json:"schemas"`
	// Name of the Log Publisher
	Id string `json:"id"`
	// The file name to use for the log files generated by the File Based Sync Log Publisher. The path to the file can be specified either as relative to the server root or as an absolute path.
	LogFile string `json:"logFile"`
	// The UNIX permissions of the log files created by this File Based Sync Log Publisher.
	LogFilePermissions string `json:"logFilePermissions"`
	// The rotation policy to use for the File Based Sync Log Publisher .
	RotationPolicy []string `json:"rotationPolicy"`
	// A listener that should be notified whenever a log file is rotated out of service.
	RotationListener []string `json:"rotationListener,omitempty"`
	// The retention policy to use for the File Based Sync Log Publisher .
	RetentionPolicy      []string                                  `json:"retentionPolicy"`
	CompressionMechanism *EnumlogPublisherCompressionMechanismProp `json:"compressionMechanism,omitempty"`
	// Indicates whether the log should be cryptographically signed so that the log content cannot be altered in an undetectable manner.
	SignLog *bool `json:"signLog,omitempty"`
	// Indicates whether log files should be encrypted so that their content is not available to unauthorized users.
	EncryptLog *bool `json:"encryptLog,omitempty"`
	// Specifies the ID of the encryption settings definition that should be used to encrypt the data. If this is not provided, the server's preferred encryption settings definition will be used. The \"encryption-settings list\" command can be used to obtain a list of the encryption settings definitions available in the server.
	EncryptionSettingsDefinitionID *string `json:"encryptionSettingsDefinitionID,omitempty"`
	// Specifies whether to append to existing log files.
	Append *bool `json:"append,omitempty"`
	// Indicates whether the File Based Sync Log Publisher will publish records asynchronously.
	Asynchronous bool `json:"asynchronous"`
	// Specifies whether to flush the writer after every log record.
	AutoFlush *bool `json:"autoFlush,omitempty"`
	// Specifies the log file buffer size.
	BufferSize *string `json:"bufferSize,omitempty"`
	// The maximum number of log records that can be stored in the asynchronous queue.
	QueueSize *int64 `json:"queueSize,omitempty"`
	// Specifies the interval at which to check whether the log files need to be rotated.
	TimeInterval       *string                                 `json:"timeInterval,omitempty"`
	TimestampPrecision *EnumlogPublisherTimestampPrecisionProp `json:"timestampPrecision,omitempty"`
	// Specifies which Sync Pipes can log messages to this Sync Log Publisher.
	IncludeSyncPipe   []string                                `json:"includeSyncPipe,omitempty"`
	LoggedMessageType []EnumlogPublisherLoggedMessageTypeProp `json:"loggedMessageType,omitempty"`
	// A description for this Log Publisher
	Description *string `json:"description,omitempty"`
	// Indicates whether the Log Publisher is enabled for use.
	Enabled                                       bool                                               `json:"enabled"`
	LoggingErrorBehavior                          *EnumlogPublisherLoggingErrorBehaviorProp          `json:"loggingErrorBehavior,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewFileBasedSyncLogPublisherResponse instantiates a new FileBasedSyncLogPublisherResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileBasedSyncLogPublisherResponse(schemas []EnumfileBasedSyncLogPublisherSchemaUrn, id string, logFile string, logFilePermissions string, rotationPolicy []string, retentionPolicy []string, asynchronous bool, enabled bool) *FileBasedSyncLogPublisherResponse {
	this := FileBasedSyncLogPublisherResponse{}
	this.Schemas = schemas
	this.Id = id
	this.LogFile = logFile
	this.LogFilePermissions = logFilePermissions
	this.RotationPolicy = rotationPolicy
	this.RetentionPolicy = retentionPolicy
	this.Asynchronous = asynchronous
	this.Enabled = enabled
	return &this
}

// NewFileBasedSyncLogPublisherResponseWithDefaults instantiates a new FileBasedSyncLogPublisherResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileBasedSyncLogPublisherResponseWithDefaults() *FileBasedSyncLogPublisherResponse {
	this := FileBasedSyncLogPublisherResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *FileBasedSyncLogPublisherResponse) GetSchemas() []EnumfileBasedSyncLogPublisherSchemaUrn {
	if o == nil {
		var ret []EnumfileBasedSyncLogPublisherSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *FileBasedSyncLogPublisherResponse) GetSchemasOk() ([]EnumfileBasedSyncLogPublisherSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *FileBasedSyncLogPublisherResponse) SetSchemas(v []EnumfileBasedSyncLogPublisherSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *FileBasedSyncLogPublisherResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FileBasedSyncLogPublisherResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FileBasedSyncLogPublisherResponse) SetId(v string) {
	o.Id = v
}

// GetLogFile returns the LogFile field value
func (o *FileBasedSyncLogPublisherResponse) GetLogFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogFile
}

// GetLogFileOk returns a tuple with the LogFile field value
// and a boolean to check if the value has been set.
func (o *FileBasedSyncLogPublisherResponse) GetLogFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogFile, true
}

// SetLogFile sets field value
func (o *FileBasedSyncLogPublisherResponse) SetLogFile(v string) {
	o.LogFile = v
}

// GetLogFilePermissions returns the LogFilePermissions field value
func (o *FileBasedSyncLogPublisherResponse) GetLogFilePermissions() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogFilePermissions
}

// GetLogFilePermissionsOk returns a tuple with the LogFilePermissions field value
// and a boolean to check if the value has been set.
func (o *FileBasedSyncLogPublisherResponse) GetLogFilePermissionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogFilePermissions, true
}

// SetLogFilePermissions sets field value
func (o *FileBasedSyncLogPublisherResponse) SetLogFilePermissions(v string) {
	o.LogFilePermissions = v
}

// GetRotationPolicy returns the RotationPolicy field value
func (o *FileBasedSyncLogPublisherResponse) GetRotationPolicy() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RotationPolicy
}

// GetRotationPolicyOk returns a tuple with the RotationPolicy field value
// and a boolean to check if the value has been set.
func (o *FileBasedSyncLogPublisherResponse) GetRotationPolicyOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RotationPolicy, true
}

// SetRotationPolicy sets field value
func (o *FileBasedSyncLogPublisherResponse) SetRotationPolicy(v []string) {
	o.RotationPolicy = v
}

// GetRotationListener returns the RotationListener field value if set, zero value otherwise.
func (o *FileBasedSyncLogPublisherResponse) GetRotationListener() []string {
	if o == nil || IsNil(o.RotationListener) {
		var ret []string
		return ret
	}
	return o.RotationListener
}

// GetRotationListenerOk returns a tuple with the RotationListener field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedSyncLogPublisherResponse) GetRotationListenerOk() ([]string, bool) {
	if o == nil || IsNil(o.RotationListener) {
		return nil, false
	}
	return o.RotationListener, true
}

// HasRotationListener returns a boolean if a field has been set.
func (o *FileBasedSyncLogPublisherResponse) HasRotationListener() bool {
	if o != nil && !IsNil(o.RotationListener) {
		return true
	}

	return false
}

// SetRotationListener gets a reference to the given []string and assigns it to the RotationListener field.
func (o *FileBasedSyncLogPublisherResponse) SetRotationListener(v []string) {
	o.RotationListener = v
}

// GetRetentionPolicy returns the RetentionPolicy field value
func (o *FileBasedSyncLogPublisherResponse) GetRetentionPolicy() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RetentionPolicy
}

// GetRetentionPolicyOk returns a tuple with the RetentionPolicy field value
// and a boolean to check if the value has been set.
func (o *FileBasedSyncLogPublisherResponse) GetRetentionPolicyOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetentionPolicy, true
}

// SetRetentionPolicy sets field value
func (o *FileBasedSyncLogPublisherResponse) SetRetentionPolicy(v []string) {
	o.RetentionPolicy = v
}

// GetCompressionMechanism returns the CompressionMechanism field value if set, zero value otherwise.
func (o *FileBasedSyncLogPublisherResponse) GetCompressionMechanism() EnumlogPublisherCompressionMechanismProp {
	if o == nil || IsNil(o.CompressionMechanism) {
		var ret EnumlogPublisherCompressionMechanismProp
		return ret
	}
	return *o.CompressionMechanism
}

// GetCompressionMechanismOk returns a tuple with the CompressionMechanism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedSyncLogPublisherResponse) GetCompressionMechanismOk() (*EnumlogPublisherCompressionMechanismProp, bool) {
	if o == nil || IsNil(o.CompressionMechanism) {
		return nil, false
	}
	return o.CompressionMechanism, true
}

// HasCompressionMechanism returns a boolean if a field has been set.
func (o *FileBasedSyncLogPublisherResponse) HasCompressionMechanism() bool {
	if o != nil && !IsNil(o.CompressionMechanism) {
		return true
	}

	return false
}

// SetCompressionMechanism gets a reference to the given EnumlogPublisherCompressionMechanismProp and assigns it to the CompressionMechanism field.
func (o *FileBasedSyncLogPublisherResponse) SetCompressionMechanism(v EnumlogPublisherCompressionMechanismProp) {
	o.CompressionMechanism = &v
}

// GetSignLog returns the SignLog field value if set, zero value otherwise.
func (o *FileBasedSyncLogPublisherResponse) GetSignLog() bool {
	if o == nil || IsNil(o.SignLog) {
		var ret bool
		return ret
	}
	return *o.SignLog
}

// GetSignLogOk returns a tuple with the SignLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedSyncLogPublisherResponse) GetSignLogOk() (*bool, bool) {
	if o == nil || IsNil(o.SignLog) {
		return nil, false
	}
	return o.SignLog, true
}

// HasSignLog returns a boolean if a field has been set.
func (o *FileBasedSyncLogPublisherResponse) HasSignLog() bool {
	if o != nil && !IsNil(o.SignLog) {
		return true
	}

	return false
}

// SetSignLog gets a reference to the given bool and assigns it to the SignLog field.
func (o *FileBasedSyncLogPublisherResponse) SetSignLog(v bool) {
	o.SignLog = &v
}

// GetEncryptLog returns the EncryptLog field value if set, zero value otherwise.
func (o *FileBasedSyncLogPublisherResponse) GetEncryptLog() bool {
	if o == nil || IsNil(o.EncryptLog) {
		var ret bool
		return ret
	}
	return *o.EncryptLog
}

// GetEncryptLogOk returns a tuple with the EncryptLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedSyncLogPublisherResponse) GetEncryptLogOk() (*bool, bool) {
	if o == nil || IsNil(o.EncryptLog) {
		return nil, false
	}
	return o.EncryptLog, true
}

// HasEncryptLog returns a boolean if a field has been set.
func (o *FileBasedSyncLogPublisherResponse) HasEncryptLog() bool {
	if o != nil && !IsNil(o.EncryptLog) {
		return true
	}

	return false
}

// SetEncryptLog gets a reference to the given bool and assigns it to the EncryptLog field.
func (o *FileBasedSyncLogPublisherResponse) SetEncryptLog(v bool) {
	o.EncryptLog = &v
}

// GetEncryptionSettingsDefinitionID returns the EncryptionSettingsDefinitionID field value if set, zero value otherwise.
func (o *FileBasedSyncLogPublisherResponse) GetEncryptionSettingsDefinitionID() string {
	if o == nil || IsNil(o.EncryptionSettingsDefinitionID) {
		var ret string
		return ret
	}
	return *o.EncryptionSettingsDefinitionID
}

// GetEncryptionSettingsDefinitionIDOk returns a tuple with the EncryptionSettingsDefinitionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedSyncLogPublisherResponse) GetEncryptionSettingsDefinitionIDOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionSettingsDefinitionID) {
		return nil, false
	}
	return o.EncryptionSettingsDefinitionID, true
}

// HasEncryptionSettingsDefinitionID returns a boolean if a field has been set.
func (o *FileBasedSyncLogPublisherResponse) HasEncryptionSettingsDefinitionID() bool {
	if o != nil && !IsNil(o.EncryptionSettingsDefinitionID) {
		return true
	}

	return false
}

// SetEncryptionSettingsDefinitionID gets a reference to the given string and assigns it to the EncryptionSettingsDefinitionID field.
func (o *FileBasedSyncLogPublisherResponse) SetEncryptionSettingsDefinitionID(v string) {
	o.EncryptionSettingsDefinitionID = &v
}

// GetAppend returns the Append field value if set, zero value otherwise.
func (o *FileBasedSyncLogPublisherResponse) GetAppend() bool {
	if o == nil || IsNil(o.Append) {
		var ret bool
		return ret
	}
	return *o.Append
}

// GetAppendOk returns a tuple with the Append field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedSyncLogPublisherResponse) GetAppendOk() (*bool, bool) {
	if o == nil || IsNil(o.Append) {
		return nil, false
	}
	return o.Append, true
}

// HasAppend returns a boolean if a field has been set.
func (o *FileBasedSyncLogPublisherResponse) HasAppend() bool {
	if o != nil && !IsNil(o.Append) {
		return true
	}

	return false
}

// SetAppend gets a reference to the given bool and assigns it to the Append field.
func (o *FileBasedSyncLogPublisherResponse) SetAppend(v bool) {
	o.Append = &v
}

// GetAsynchronous returns the Asynchronous field value
func (o *FileBasedSyncLogPublisherResponse) GetAsynchronous() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Asynchronous
}

// GetAsynchronousOk returns a tuple with the Asynchronous field value
// and a boolean to check if the value has been set.
func (o *FileBasedSyncLogPublisherResponse) GetAsynchronousOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asynchronous, true
}

// SetAsynchronous sets field value
func (o *FileBasedSyncLogPublisherResponse) SetAsynchronous(v bool) {
	o.Asynchronous = v
}

// GetAutoFlush returns the AutoFlush field value if set, zero value otherwise.
func (o *FileBasedSyncLogPublisherResponse) GetAutoFlush() bool {
	if o == nil || IsNil(o.AutoFlush) {
		var ret bool
		return ret
	}
	return *o.AutoFlush
}

// GetAutoFlushOk returns a tuple with the AutoFlush field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedSyncLogPublisherResponse) GetAutoFlushOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoFlush) {
		return nil, false
	}
	return o.AutoFlush, true
}

// HasAutoFlush returns a boolean if a field has been set.
func (o *FileBasedSyncLogPublisherResponse) HasAutoFlush() bool {
	if o != nil && !IsNil(o.AutoFlush) {
		return true
	}

	return false
}

// SetAutoFlush gets a reference to the given bool and assigns it to the AutoFlush field.
func (o *FileBasedSyncLogPublisherResponse) SetAutoFlush(v bool) {
	o.AutoFlush = &v
}

// GetBufferSize returns the BufferSize field value if set, zero value otherwise.
func (o *FileBasedSyncLogPublisherResponse) GetBufferSize() string {
	if o == nil || IsNil(o.BufferSize) {
		var ret string
		return ret
	}
	return *o.BufferSize
}

// GetBufferSizeOk returns a tuple with the BufferSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedSyncLogPublisherResponse) GetBufferSizeOk() (*string, bool) {
	if o == nil || IsNil(o.BufferSize) {
		return nil, false
	}
	return o.BufferSize, true
}

// HasBufferSize returns a boolean if a field has been set.
func (o *FileBasedSyncLogPublisherResponse) HasBufferSize() bool {
	if o != nil && !IsNil(o.BufferSize) {
		return true
	}

	return false
}

// SetBufferSize gets a reference to the given string and assigns it to the BufferSize field.
func (o *FileBasedSyncLogPublisherResponse) SetBufferSize(v string) {
	o.BufferSize = &v
}

// GetQueueSize returns the QueueSize field value if set, zero value otherwise.
func (o *FileBasedSyncLogPublisherResponse) GetQueueSize() int64 {
	if o == nil || IsNil(o.QueueSize) {
		var ret int64
		return ret
	}
	return *o.QueueSize
}

// GetQueueSizeOk returns a tuple with the QueueSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedSyncLogPublisherResponse) GetQueueSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.QueueSize) {
		return nil, false
	}
	return o.QueueSize, true
}

// HasQueueSize returns a boolean if a field has been set.
func (o *FileBasedSyncLogPublisherResponse) HasQueueSize() bool {
	if o != nil && !IsNil(o.QueueSize) {
		return true
	}

	return false
}

// SetQueueSize gets a reference to the given int64 and assigns it to the QueueSize field.
func (o *FileBasedSyncLogPublisherResponse) SetQueueSize(v int64) {
	o.QueueSize = &v
}

// GetTimeInterval returns the TimeInterval field value if set, zero value otherwise.
func (o *FileBasedSyncLogPublisherResponse) GetTimeInterval() string {
	if o == nil || IsNil(o.TimeInterval) {
		var ret string
		return ret
	}
	return *o.TimeInterval
}

// GetTimeIntervalOk returns a tuple with the TimeInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedSyncLogPublisherResponse) GetTimeIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.TimeInterval) {
		return nil, false
	}
	return o.TimeInterval, true
}

// HasTimeInterval returns a boolean if a field has been set.
func (o *FileBasedSyncLogPublisherResponse) HasTimeInterval() bool {
	if o != nil && !IsNil(o.TimeInterval) {
		return true
	}

	return false
}

// SetTimeInterval gets a reference to the given string and assigns it to the TimeInterval field.
func (o *FileBasedSyncLogPublisherResponse) SetTimeInterval(v string) {
	o.TimeInterval = &v
}

// GetTimestampPrecision returns the TimestampPrecision field value if set, zero value otherwise.
func (o *FileBasedSyncLogPublisherResponse) GetTimestampPrecision() EnumlogPublisherTimestampPrecisionProp {
	if o == nil || IsNil(o.TimestampPrecision) {
		var ret EnumlogPublisherTimestampPrecisionProp
		return ret
	}
	return *o.TimestampPrecision
}

// GetTimestampPrecisionOk returns a tuple with the TimestampPrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedSyncLogPublisherResponse) GetTimestampPrecisionOk() (*EnumlogPublisherTimestampPrecisionProp, bool) {
	if o == nil || IsNil(o.TimestampPrecision) {
		return nil, false
	}
	return o.TimestampPrecision, true
}

// HasTimestampPrecision returns a boolean if a field has been set.
func (o *FileBasedSyncLogPublisherResponse) HasTimestampPrecision() bool {
	if o != nil && !IsNil(o.TimestampPrecision) {
		return true
	}

	return false
}

// SetTimestampPrecision gets a reference to the given EnumlogPublisherTimestampPrecisionProp and assigns it to the TimestampPrecision field.
func (o *FileBasedSyncLogPublisherResponse) SetTimestampPrecision(v EnumlogPublisherTimestampPrecisionProp) {
	o.TimestampPrecision = &v
}

// GetIncludeSyncPipe returns the IncludeSyncPipe field value if set, zero value otherwise.
func (o *FileBasedSyncLogPublisherResponse) GetIncludeSyncPipe() []string {
	if o == nil || IsNil(o.IncludeSyncPipe) {
		var ret []string
		return ret
	}
	return o.IncludeSyncPipe
}

// GetIncludeSyncPipeOk returns a tuple with the IncludeSyncPipe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedSyncLogPublisherResponse) GetIncludeSyncPipeOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeSyncPipe) {
		return nil, false
	}
	return o.IncludeSyncPipe, true
}

// HasIncludeSyncPipe returns a boolean if a field has been set.
func (o *FileBasedSyncLogPublisherResponse) HasIncludeSyncPipe() bool {
	if o != nil && !IsNil(o.IncludeSyncPipe) {
		return true
	}

	return false
}

// SetIncludeSyncPipe gets a reference to the given []string and assigns it to the IncludeSyncPipe field.
func (o *FileBasedSyncLogPublisherResponse) SetIncludeSyncPipe(v []string) {
	o.IncludeSyncPipe = v
}

// GetLoggedMessageType returns the LoggedMessageType field value if set, zero value otherwise.
func (o *FileBasedSyncLogPublisherResponse) GetLoggedMessageType() []EnumlogPublisherLoggedMessageTypeProp {
	if o == nil || IsNil(o.LoggedMessageType) {
		var ret []EnumlogPublisherLoggedMessageTypeProp
		return ret
	}
	return o.LoggedMessageType
}

// GetLoggedMessageTypeOk returns a tuple with the LoggedMessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedSyncLogPublisherResponse) GetLoggedMessageTypeOk() ([]EnumlogPublisherLoggedMessageTypeProp, bool) {
	if o == nil || IsNil(o.LoggedMessageType) {
		return nil, false
	}
	return o.LoggedMessageType, true
}

// HasLoggedMessageType returns a boolean if a field has been set.
func (o *FileBasedSyncLogPublisherResponse) HasLoggedMessageType() bool {
	if o != nil && !IsNil(o.LoggedMessageType) {
		return true
	}

	return false
}

// SetLoggedMessageType gets a reference to the given []EnumlogPublisherLoggedMessageTypeProp and assigns it to the LoggedMessageType field.
func (o *FileBasedSyncLogPublisherResponse) SetLoggedMessageType(v []EnumlogPublisherLoggedMessageTypeProp) {
	o.LoggedMessageType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FileBasedSyncLogPublisherResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedSyncLogPublisherResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FileBasedSyncLogPublisherResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FileBasedSyncLogPublisherResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *FileBasedSyncLogPublisherResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *FileBasedSyncLogPublisherResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *FileBasedSyncLogPublisherResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetLoggingErrorBehavior returns the LoggingErrorBehavior field value if set, zero value otherwise.
func (o *FileBasedSyncLogPublisherResponse) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp {
	if o == nil || IsNil(o.LoggingErrorBehavior) {
		var ret EnumlogPublisherLoggingErrorBehaviorProp
		return ret
	}
	return *o.LoggingErrorBehavior
}

// GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedSyncLogPublisherResponse) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool) {
	if o == nil || IsNil(o.LoggingErrorBehavior) {
		return nil, false
	}
	return o.LoggingErrorBehavior, true
}

// HasLoggingErrorBehavior returns a boolean if a field has been set.
func (o *FileBasedSyncLogPublisherResponse) HasLoggingErrorBehavior() bool {
	if o != nil && !IsNil(o.LoggingErrorBehavior) {
		return true
	}

	return false
}

// SetLoggingErrorBehavior gets a reference to the given EnumlogPublisherLoggingErrorBehaviorProp and assigns it to the LoggingErrorBehavior field.
func (o *FileBasedSyncLogPublisherResponse) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp) {
	o.LoggingErrorBehavior = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *FileBasedSyncLogPublisherResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedSyncLogPublisherResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *FileBasedSyncLogPublisherResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *FileBasedSyncLogPublisherResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *FileBasedSyncLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedSyncLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *FileBasedSyncLogPublisherResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *FileBasedSyncLogPublisherResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o FileBasedSyncLogPublisherResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileBasedSyncLogPublisherResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
	toSerialize["logFile"] = o.LogFile
	toSerialize["logFilePermissions"] = o.LogFilePermissions
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
	if !IsNil(o.Append) {
		toSerialize["append"] = o.Append
	}
	toSerialize["asynchronous"] = o.Asynchronous
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
	if !IsNil(o.TimestampPrecision) {
		toSerialize["timestampPrecision"] = o.TimestampPrecision
	}
	if !IsNil(o.IncludeSyncPipe) {
		toSerialize["includeSyncPipe"] = o.IncludeSyncPipe
	}
	if !IsNil(o.LoggedMessageType) {
		toSerialize["loggedMessageType"] = o.LoggedMessageType
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

type NullableFileBasedSyncLogPublisherResponse struct {
	value *FileBasedSyncLogPublisherResponse
	isSet bool
}

func (v NullableFileBasedSyncLogPublisherResponse) Get() *FileBasedSyncLogPublisherResponse {
	return v.value
}

func (v *NullableFileBasedSyncLogPublisherResponse) Set(val *FileBasedSyncLogPublisherResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFileBasedSyncLogPublisherResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFileBasedSyncLogPublisherResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileBasedSyncLogPublisherResponse(val *FileBasedSyncLogPublisherResponse) *NullableFileBasedSyncLogPublisherResponse {
	return &NullableFileBasedSyncLogPublisherResponse{value: val, isSet: true}
}

func (v NullableFileBasedSyncLogPublisherResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileBasedSyncLogPublisherResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
