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

// checks if the AddCommonLogFileHttpOperationLogPublisherRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddCommonLogFileHttpOperationLogPublisherRequest{}

// AddCommonLogFileHttpOperationLogPublisherRequest struct for AddCommonLogFileHttpOperationLogPublisherRequest
type AddCommonLogFileHttpOperationLogPublisherRequest struct {
	Schemas []EnumcommonLogFileHttpOperationLogPublisherSchemaUrn `json:"schemas"`
	// The file name to use for the log files generated by the Common Log File HTTP Operation Log Publisher. The path to the file can be specified either as relative to the server root or as an absolute path.
	LogFile string `json:"logFile"`
	// The UNIX permissions of the log files created by this Common Log File HTTP Operation Log Publisher.
	LogFilePermissions *string `json:"logFilePermissions,omitempty"`
	// The rotation policy to use for the Common Log File HTTP Operation Log Publisher .
	RotationPolicy []string `json:"rotationPolicy,omitempty"`
	// A listener that should be notified whenever a log file is rotated out of service.
	RotationListener []string `json:"rotationListener,omitempty"`
	// The retention policy to use for the Common Log File HTTP Operation Log Publisher .
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
	// Indicates whether the Common Log File HTTP Operation Log Publisher will publish records asynchronously.
	Asynchronous *bool `json:"asynchronous,omitempty"`
	// Specifies whether to flush the writer after every log record.
	AutoFlush *bool `json:"autoFlush,omitempty"`
	// Specifies the log file buffer size.
	BufferSize *string `json:"bufferSize,omitempty"`
	// The maximum number of log records that can be stored in the asynchronous queue.
	QueueSize *int64 `json:"queueSize,omitempty"`
	// Specifies the interval at which to check whether the log files need to be rotated.
	TimeInterval *string `json:"timeInterval,omitempty"`
	// A description for this Log Publisher
	Description *string `json:"description,omitempty"`
	// Indicates whether the Log Publisher is enabled for use.
	Enabled              bool                                      `json:"enabled"`
	LoggingErrorBehavior *EnumlogPublisherLoggingErrorBehaviorProp `json:"loggingErrorBehavior,omitempty"`
	// Name of the new Log Publisher
	PublisherName string `json:"publisherName"`
}

// NewAddCommonLogFileHttpOperationLogPublisherRequest instantiates a new AddCommonLogFileHttpOperationLogPublisherRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddCommonLogFileHttpOperationLogPublisherRequest(schemas []EnumcommonLogFileHttpOperationLogPublisherSchemaUrn, logFile string, enabled bool, publisherName string) *AddCommonLogFileHttpOperationLogPublisherRequest {
	this := AddCommonLogFileHttpOperationLogPublisherRequest{}
	this.Schemas = schemas
	this.LogFile = logFile
	this.Enabled = enabled
	this.PublisherName = publisherName
	return &this
}

// NewAddCommonLogFileHttpOperationLogPublisherRequestWithDefaults instantiates a new AddCommonLogFileHttpOperationLogPublisherRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCommonLogFileHttpOperationLogPublisherRequestWithDefaults() *AddCommonLogFileHttpOperationLogPublisherRequest {
	this := AddCommonLogFileHttpOperationLogPublisherRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetSchemas() []EnumcommonLogFileHttpOperationLogPublisherSchemaUrn {
	if o == nil {
		var ret []EnumcommonLogFileHttpOperationLogPublisherSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetSchemasOk() ([]EnumcommonLogFileHttpOperationLogPublisherSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) SetSchemas(v []EnumcommonLogFileHttpOperationLogPublisherSchemaUrn) {
	o.Schemas = v
}

// GetLogFile returns the LogFile field value
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetLogFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogFile
}

// GetLogFileOk returns a tuple with the LogFile field value
// and a boolean to check if the value has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetLogFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogFile, true
}

// SetLogFile sets field value
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) SetLogFile(v string) {
	o.LogFile = v
}

// GetLogFilePermissions returns the LogFilePermissions field value if set, zero value otherwise.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetLogFilePermissions() string {
	if o == nil || IsNil(o.LogFilePermissions) {
		var ret string
		return ret
	}
	return *o.LogFilePermissions
}

// GetLogFilePermissionsOk returns a tuple with the LogFilePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetLogFilePermissionsOk() (*string, bool) {
	if o == nil || IsNil(o.LogFilePermissions) {
		return nil, false
	}
	return o.LogFilePermissions, true
}

// HasLogFilePermissions returns a boolean if a field has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) HasLogFilePermissions() bool {
	if o != nil && !IsNil(o.LogFilePermissions) {
		return true
	}

	return false
}

// SetLogFilePermissions gets a reference to the given string and assigns it to the LogFilePermissions field.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) SetLogFilePermissions(v string) {
	o.LogFilePermissions = &v
}

// GetRotationPolicy returns the RotationPolicy field value if set, zero value otherwise.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetRotationPolicy() []string {
	if o == nil || IsNil(o.RotationPolicy) {
		var ret []string
		return ret
	}
	return o.RotationPolicy
}

// GetRotationPolicyOk returns a tuple with the RotationPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetRotationPolicyOk() ([]string, bool) {
	if o == nil || IsNil(o.RotationPolicy) {
		return nil, false
	}
	return o.RotationPolicy, true
}

// HasRotationPolicy returns a boolean if a field has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) HasRotationPolicy() bool {
	if o != nil && !IsNil(o.RotationPolicy) {
		return true
	}

	return false
}

// SetRotationPolicy gets a reference to the given []string and assigns it to the RotationPolicy field.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) SetRotationPolicy(v []string) {
	o.RotationPolicy = v
}

// GetRotationListener returns the RotationListener field value if set, zero value otherwise.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetRotationListener() []string {
	if o == nil || IsNil(o.RotationListener) {
		var ret []string
		return ret
	}
	return o.RotationListener
}

// GetRotationListenerOk returns a tuple with the RotationListener field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetRotationListenerOk() ([]string, bool) {
	if o == nil || IsNil(o.RotationListener) {
		return nil, false
	}
	return o.RotationListener, true
}

// HasRotationListener returns a boolean if a field has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) HasRotationListener() bool {
	if o != nil && !IsNil(o.RotationListener) {
		return true
	}

	return false
}

// SetRotationListener gets a reference to the given []string and assigns it to the RotationListener field.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) SetRotationListener(v []string) {
	o.RotationListener = v
}

// GetRetentionPolicy returns the RetentionPolicy field value if set, zero value otherwise.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetRetentionPolicy() []string {
	if o == nil || IsNil(o.RetentionPolicy) {
		var ret []string
		return ret
	}
	return o.RetentionPolicy
}

// GetRetentionPolicyOk returns a tuple with the RetentionPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetRetentionPolicyOk() ([]string, bool) {
	if o == nil || IsNil(o.RetentionPolicy) {
		return nil, false
	}
	return o.RetentionPolicy, true
}

// HasRetentionPolicy returns a boolean if a field has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) HasRetentionPolicy() bool {
	if o != nil && !IsNil(o.RetentionPolicy) {
		return true
	}

	return false
}

// SetRetentionPolicy gets a reference to the given []string and assigns it to the RetentionPolicy field.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) SetRetentionPolicy(v []string) {
	o.RetentionPolicy = v
}

// GetCompressionMechanism returns the CompressionMechanism field value if set, zero value otherwise.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetCompressionMechanism() EnumlogPublisherCompressionMechanismProp {
	if o == nil || IsNil(o.CompressionMechanism) {
		var ret EnumlogPublisherCompressionMechanismProp
		return ret
	}
	return *o.CompressionMechanism
}

// GetCompressionMechanismOk returns a tuple with the CompressionMechanism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetCompressionMechanismOk() (*EnumlogPublisherCompressionMechanismProp, bool) {
	if o == nil || IsNil(o.CompressionMechanism) {
		return nil, false
	}
	return o.CompressionMechanism, true
}

// HasCompressionMechanism returns a boolean if a field has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) HasCompressionMechanism() bool {
	if o != nil && !IsNil(o.CompressionMechanism) {
		return true
	}

	return false
}

// SetCompressionMechanism gets a reference to the given EnumlogPublisherCompressionMechanismProp and assigns it to the CompressionMechanism field.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) SetCompressionMechanism(v EnumlogPublisherCompressionMechanismProp) {
	o.CompressionMechanism = &v
}

// GetSignLog returns the SignLog field value if set, zero value otherwise.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetSignLog() bool {
	if o == nil || IsNil(o.SignLog) {
		var ret bool
		return ret
	}
	return *o.SignLog
}

// GetSignLogOk returns a tuple with the SignLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetSignLogOk() (*bool, bool) {
	if o == nil || IsNil(o.SignLog) {
		return nil, false
	}
	return o.SignLog, true
}

// HasSignLog returns a boolean if a field has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) HasSignLog() bool {
	if o != nil && !IsNil(o.SignLog) {
		return true
	}

	return false
}

// SetSignLog gets a reference to the given bool and assigns it to the SignLog field.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) SetSignLog(v bool) {
	o.SignLog = &v
}

// GetEncryptLog returns the EncryptLog field value if set, zero value otherwise.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetEncryptLog() bool {
	if o == nil || IsNil(o.EncryptLog) {
		var ret bool
		return ret
	}
	return *o.EncryptLog
}

// GetEncryptLogOk returns a tuple with the EncryptLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetEncryptLogOk() (*bool, bool) {
	if o == nil || IsNil(o.EncryptLog) {
		return nil, false
	}
	return o.EncryptLog, true
}

// HasEncryptLog returns a boolean if a field has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) HasEncryptLog() bool {
	if o != nil && !IsNil(o.EncryptLog) {
		return true
	}

	return false
}

// SetEncryptLog gets a reference to the given bool and assigns it to the EncryptLog field.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) SetEncryptLog(v bool) {
	o.EncryptLog = &v
}

// GetEncryptionSettingsDefinitionID returns the EncryptionSettingsDefinitionID field value if set, zero value otherwise.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetEncryptionSettingsDefinitionID() string {
	if o == nil || IsNil(o.EncryptionSettingsDefinitionID) {
		var ret string
		return ret
	}
	return *o.EncryptionSettingsDefinitionID
}

// GetEncryptionSettingsDefinitionIDOk returns a tuple with the EncryptionSettingsDefinitionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetEncryptionSettingsDefinitionIDOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionSettingsDefinitionID) {
		return nil, false
	}
	return o.EncryptionSettingsDefinitionID, true
}

// HasEncryptionSettingsDefinitionID returns a boolean if a field has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) HasEncryptionSettingsDefinitionID() bool {
	if o != nil && !IsNil(o.EncryptionSettingsDefinitionID) {
		return true
	}

	return false
}

// SetEncryptionSettingsDefinitionID gets a reference to the given string and assigns it to the EncryptionSettingsDefinitionID field.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) SetEncryptionSettingsDefinitionID(v string) {
	o.EncryptionSettingsDefinitionID = &v
}

// GetAppend returns the Append field value if set, zero value otherwise.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetAppend() bool {
	if o == nil || IsNil(o.Append) {
		var ret bool
		return ret
	}
	return *o.Append
}

// GetAppendOk returns a tuple with the Append field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetAppendOk() (*bool, bool) {
	if o == nil || IsNil(o.Append) {
		return nil, false
	}
	return o.Append, true
}

// HasAppend returns a boolean if a field has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) HasAppend() bool {
	if o != nil && !IsNil(o.Append) {
		return true
	}

	return false
}

// SetAppend gets a reference to the given bool and assigns it to the Append field.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) SetAppend(v bool) {
	o.Append = &v
}

// GetAsynchronous returns the Asynchronous field value if set, zero value otherwise.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetAsynchronous() bool {
	if o == nil || IsNil(o.Asynchronous) {
		var ret bool
		return ret
	}
	return *o.Asynchronous
}

// GetAsynchronousOk returns a tuple with the Asynchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetAsynchronousOk() (*bool, bool) {
	if o == nil || IsNil(o.Asynchronous) {
		return nil, false
	}
	return o.Asynchronous, true
}

// HasAsynchronous returns a boolean if a field has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) HasAsynchronous() bool {
	if o != nil && !IsNil(o.Asynchronous) {
		return true
	}

	return false
}

// SetAsynchronous gets a reference to the given bool and assigns it to the Asynchronous field.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) SetAsynchronous(v bool) {
	o.Asynchronous = &v
}

// GetAutoFlush returns the AutoFlush field value if set, zero value otherwise.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetAutoFlush() bool {
	if o == nil || IsNil(o.AutoFlush) {
		var ret bool
		return ret
	}
	return *o.AutoFlush
}

// GetAutoFlushOk returns a tuple with the AutoFlush field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetAutoFlushOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoFlush) {
		return nil, false
	}
	return o.AutoFlush, true
}

// HasAutoFlush returns a boolean if a field has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) HasAutoFlush() bool {
	if o != nil && !IsNil(o.AutoFlush) {
		return true
	}

	return false
}

// SetAutoFlush gets a reference to the given bool and assigns it to the AutoFlush field.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) SetAutoFlush(v bool) {
	o.AutoFlush = &v
}

// GetBufferSize returns the BufferSize field value if set, zero value otherwise.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetBufferSize() string {
	if o == nil || IsNil(o.BufferSize) {
		var ret string
		return ret
	}
	return *o.BufferSize
}

// GetBufferSizeOk returns a tuple with the BufferSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetBufferSizeOk() (*string, bool) {
	if o == nil || IsNil(o.BufferSize) {
		return nil, false
	}
	return o.BufferSize, true
}

// HasBufferSize returns a boolean if a field has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) HasBufferSize() bool {
	if o != nil && !IsNil(o.BufferSize) {
		return true
	}

	return false
}

// SetBufferSize gets a reference to the given string and assigns it to the BufferSize field.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) SetBufferSize(v string) {
	o.BufferSize = &v
}

// GetQueueSize returns the QueueSize field value if set, zero value otherwise.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetQueueSize() int64 {
	if o == nil || IsNil(o.QueueSize) {
		var ret int64
		return ret
	}
	return *o.QueueSize
}

// GetQueueSizeOk returns a tuple with the QueueSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetQueueSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.QueueSize) {
		return nil, false
	}
	return o.QueueSize, true
}

// HasQueueSize returns a boolean if a field has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) HasQueueSize() bool {
	if o != nil && !IsNil(o.QueueSize) {
		return true
	}

	return false
}

// SetQueueSize gets a reference to the given int64 and assigns it to the QueueSize field.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) SetQueueSize(v int64) {
	o.QueueSize = &v
}

// GetTimeInterval returns the TimeInterval field value if set, zero value otherwise.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetTimeInterval() string {
	if o == nil || IsNil(o.TimeInterval) {
		var ret string
		return ret
	}
	return *o.TimeInterval
}

// GetTimeIntervalOk returns a tuple with the TimeInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetTimeIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.TimeInterval) {
		return nil, false
	}
	return o.TimeInterval, true
}

// HasTimeInterval returns a boolean if a field has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) HasTimeInterval() bool {
	if o != nil && !IsNil(o.TimeInterval) {
		return true
	}

	return false
}

// SetTimeInterval gets a reference to the given string and assigns it to the TimeInterval field.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) SetTimeInterval(v string) {
	o.TimeInterval = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetLoggingErrorBehavior returns the LoggingErrorBehavior field value if set, zero value otherwise.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp {
	if o == nil || IsNil(o.LoggingErrorBehavior) {
		var ret EnumlogPublisherLoggingErrorBehaviorProp
		return ret
	}
	return *o.LoggingErrorBehavior
}

// GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool) {
	if o == nil || IsNil(o.LoggingErrorBehavior) {
		return nil, false
	}
	return o.LoggingErrorBehavior, true
}

// HasLoggingErrorBehavior returns a boolean if a field has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) HasLoggingErrorBehavior() bool {
	if o != nil && !IsNil(o.LoggingErrorBehavior) {
		return true
	}

	return false
}

// SetLoggingErrorBehavior gets a reference to the given EnumlogPublisherLoggingErrorBehaviorProp and assigns it to the LoggingErrorBehavior field.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp) {
	o.LoggingErrorBehavior = &v
}

// GetPublisherName returns the PublisherName field value
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetPublisherName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublisherName
}

// GetPublisherNameOk returns a tuple with the PublisherName field value
// and a boolean to check if the value has been set.
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) GetPublisherNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublisherName, true
}

// SetPublisherName sets field value
func (o *AddCommonLogFileHttpOperationLogPublisherRequest) SetPublisherName(v string) {
	o.PublisherName = v
}

func (o AddCommonLogFileHttpOperationLogPublisherRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddCommonLogFileHttpOperationLogPublisherRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
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
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.LoggingErrorBehavior) {
		toSerialize["loggingErrorBehavior"] = o.LoggingErrorBehavior
	}
	toSerialize["publisherName"] = o.PublisherName
	return toSerialize, nil
}

type NullableAddCommonLogFileHttpOperationLogPublisherRequest struct {
	value *AddCommonLogFileHttpOperationLogPublisherRequest
	isSet bool
}

func (v NullableAddCommonLogFileHttpOperationLogPublisherRequest) Get() *AddCommonLogFileHttpOperationLogPublisherRequest {
	return v.value
}

func (v *NullableAddCommonLogFileHttpOperationLogPublisherRequest) Set(val *AddCommonLogFileHttpOperationLogPublisherRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddCommonLogFileHttpOperationLogPublisherRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddCommonLogFileHttpOperationLogPublisherRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddCommonLogFileHttpOperationLogPublisherRequest(val *AddCommonLogFileHttpOperationLogPublisherRequest) *NullableAddCommonLogFileHttpOperationLogPublisherRequest {
	return &NullableAddCommonLogFileHttpOperationLogPublisherRequest{value: val, isSet: true}
}

func (v NullableAddCommonLogFileHttpOperationLogPublisherRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddCommonLogFileHttpOperationLogPublisherRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
