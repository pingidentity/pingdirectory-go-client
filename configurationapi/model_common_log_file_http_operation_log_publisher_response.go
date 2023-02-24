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

// checks if the CommonLogFileHttpOperationLogPublisherResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonLogFileHttpOperationLogPublisherResponse{}

// CommonLogFileHttpOperationLogPublisherResponse struct for CommonLogFileHttpOperationLogPublisherResponse
type CommonLogFileHttpOperationLogPublisherResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Log Publisher
	Id      string                                                `json:"id"`
	Schemas []EnumcommonLogFileHttpOperationLogPublisherSchemaUrn `json:"schemas"`
	// The file name to use for the log files generated by the Common Log File HTTP Operation Log Publisher. The path to the file can be specified either as relative to the server root or as an absolute path.
	LogFile string `json:"logFile"`
	// The UNIX permissions of the log files created by this Common Log File HTTP Operation Log Publisher.
	LogFilePermissions string `json:"logFilePermissions"`
	// The rotation policy to use for the Common Log File HTTP Operation Log Publisher .
	RotationPolicy []string `json:"rotationPolicy"`
	// A listener that should be notified whenever a log file is rotated out of service.
	RotationListener []string `json:"rotationListener,omitempty"`
	// The retention policy to use for the Common Log File HTTP Operation Log Publisher .
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
	// Indicates whether the Common Log File HTTP Operation Log Publisher will publish records asynchronously.
	Asynchronous bool `json:"asynchronous"`
	// Specifies whether to flush the writer after every log record.
	AutoFlush *bool `json:"autoFlush,omitempty"`
	// Specifies the log file buffer size.
	BufferSize *string `json:"bufferSize,omitempty"`
	// The maximum number of log records that can be stored in the asynchronous queue.
	QueueSize *int32 `json:"queueSize,omitempty"`
	// Specifies the interval at which to check whether the log files need to be rotated.
	TimeInterval *string `json:"timeInterval,omitempty"`
	// A description for this Log Publisher
	Description *string `json:"description,omitempty"`
	// Indicates whether the Log Publisher is enabled for use.
	Enabled              bool                                      `json:"enabled"`
	LoggingErrorBehavior *EnumlogPublisherLoggingErrorBehaviorProp `json:"loggingErrorBehavior,omitempty"`
}

// NewCommonLogFileHttpOperationLogPublisherResponse instantiates a new CommonLogFileHttpOperationLogPublisherResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonLogFileHttpOperationLogPublisherResponse(id string, schemas []EnumcommonLogFileHttpOperationLogPublisherSchemaUrn, logFile string, logFilePermissions string, rotationPolicy []string, retentionPolicy []string, asynchronous bool, enabled bool) *CommonLogFileHttpOperationLogPublisherResponse {
	this := CommonLogFileHttpOperationLogPublisherResponse{}
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

// NewCommonLogFileHttpOperationLogPublisherResponseWithDefaults instantiates a new CommonLogFileHttpOperationLogPublisherResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonLogFileHttpOperationLogPublisherResponseWithDefaults() *CommonLogFileHttpOperationLogPublisherResponse {
	this := CommonLogFileHttpOperationLogPublisherResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *CommonLogFileHttpOperationLogPublisherResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *CommonLogFileHttpOperationLogPublisherResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CommonLogFileHttpOperationLogPublisherResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetSchemas() []EnumcommonLogFileHttpOperationLogPublisherSchemaUrn {
	if o == nil {
		var ret []EnumcommonLogFileHttpOperationLogPublisherSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetSchemasOk() ([]EnumcommonLogFileHttpOperationLogPublisherSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *CommonLogFileHttpOperationLogPublisherResponse) SetSchemas(v []EnumcommonLogFileHttpOperationLogPublisherSchemaUrn) {
	o.Schemas = v
}

// GetLogFile returns the LogFile field value
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetLogFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogFile
}

// GetLogFileOk returns a tuple with the LogFile field value
// and a boolean to check if the value has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetLogFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogFile, true
}

// SetLogFile sets field value
func (o *CommonLogFileHttpOperationLogPublisherResponse) SetLogFile(v string) {
	o.LogFile = v
}

// GetLogFilePermissions returns the LogFilePermissions field value
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetLogFilePermissions() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogFilePermissions
}

// GetLogFilePermissionsOk returns a tuple with the LogFilePermissions field value
// and a boolean to check if the value has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetLogFilePermissionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogFilePermissions, true
}

// SetLogFilePermissions sets field value
func (o *CommonLogFileHttpOperationLogPublisherResponse) SetLogFilePermissions(v string) {
	o.LogFilePermissions = v
}

// GetRotationPolicy returns the RotationPolicy field value
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetRotationPolicy() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RotationPolicy
}

// GetRotationPolicyOk returns a tuple with the RotationPolicy field value
// and a boolean to check if the value has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetRotationPolicyOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RotationPolicy, true
}

// SetRotationPolicy sets field value
func (o *CommonLogFileHttpOperationLogPublisherResponse) SetRotationPolicy(v []string) {
	o.RotationPolicy = v
}

// GetRotationListener returns the RotationListener field value if set, zero value otherwise.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetRotationListener() []string {
	if o == nil || IsNil(o.RotationListener) {
		var ret []string
		return ret
	}
	return o.RotationListener
}

// GetRotationListenerOk returns a tuple with the RotationListener field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetRotationListenerOk() ([]string, bool) {
	if o == nil || IsNil(o.RotationListener) {
		return nil, false
	}
	return o.RotationListener, true
}

// HasRotationListener returns a boolean if a field has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) HasRotationListener() bool {
	if o != nil && !IsNil(o.RotationListener) {
		return true
	}

	return false
}

// SetRotationListener gets a reference to the given []string and assigns it to the RotationListener field.
func (o *CommonLogFileHttpOperationLogPublisherResponse) SetRotationListener(v []string) {
	o.RotationListener = v
}

// GetRetentionPolicy returns the RetentionPolicy field value
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetRetentionPolicy() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RetentionPolicy
}

// GetRetentionPolicyOk returns a tuple with the RetentionPolicy field value
// and a boolean to check if the value has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetRetentionPolicyOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetentionPolicy, true
}

// SetRetentionPolicy sets field value
func (o *CommonLogFileHttpOperationLogPublisherResponse) SetRetentionPolicy(v []string) {
	o.RetentionPolicy = v
}

// GetCompressionMechanism returns the CompressionMechanism field value if set, zero value otherwise.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetCompressionMechanism() EnumlogPublisherCompressionMechanismProp {
	if o == nil || IsNil(o.CompressionMechanism) {
		var ret EnumlogPublisherCompressionMechanismProp
		return ret
	}
	return *o.CompressionMechanism
}

// GetCompressionMechanismOk returns a tuple with the CompressionMechanism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetCompressionMechanismOk() (*EnumlogPublisherCompressionMechanismProp, bool) {
	if o == nil || IsNil(o.CompressionMechanism) {
		return nil, false
	}
	return o.CompressionMechanism, true
}

// HasCompressionMechanism returns a boolean if a field has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) HasCompressionMechanism() bool {
	if o != nil && !IsNil(o.CompressionMechanism) {
		return true
	}

	return false
}

// SetCompressionMechanism gets a reference to the given EnumlogPublisherCompressionMechanismProp and assigns it to the CompressionMechanism field.
func (o *CommonLogFileHttpOperationLogPublisherResponse) SetCompressionMechanism(v EnumlogPublisherCompressionMechanismProp) {
	o.CompressionMechanism = &v
}

// GetSignLog returns the SignLog field value if set, zero value otherwise.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetSignLog() bool {
	if o == nil || IsNil(o.SignLog) {
		var ret bool
		return ret
	}
	return *o.SignLog
}

// GetSignLogOk returns a tuple with the SignLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetSignLogOk() (*bool, bool) {
	if o == nil || IsNil(o.SignLog) {
		return nil, false
	}
	return o.SignLog, true
}

// HasSignLog returns a boolean if a field has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) HasSignLog() bool {
	if o != nil && !IsNil(o.SignLog) {
		return true
	}

	return false
}

// SetSignLog gets a reference to the given bool and assigns it to the SignLog field.
func (o *CommonLogFileHttpOperationLogPublisherResponse) SetSignLog(v bool) {
	o.SignLog = &v
}

// GetEncryptLog returns the EncryptLog field value if set, zero value otherwise.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetEncryptLog() bool {
	if o == nil || IsNil(o.EncryptLog) {
		var ret bool
		return ret
	}
	return *o.EncryptLog
}

// GetEncryptLogOk returns a tuple with the EncryptLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetEncryptLogOk() (*bool, bool) {
	if o == nil || IsNil(o.EncryptLog) {
		return nil, false
	}
	return o.EncryptLog, true
}

// HasEncryptLog returns a boolean if a field has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) HasEncryptLog() bool {
	if o != nil && !IsNil(o.EncryptLog) {
		return true
	}

	return false
}

// SetEncryptLog gets a reference to the given bool and assigns it to the EncryptLog field.
func (o *CommonLogFileHttpOperationLogPublisherResponse) SetEncryptLog(v bool) {
	o.EncryptLog = &v
}

// GetEncryptionSettingsDefinitionID returns the EncryptionSettingsDefinitionID field value if set, zero value otherwise.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetEncryptionSettingsDefinitionID() string {
	if o == nil || IsNil(o.EncryptionSettingsDefinitionID) {
		var ret string
		return ret
	}
	return *o.EncryptionSettingsDefinitionID
}

// GetEncryptionSettingsDefinitionIDOk returns a tuple with the EncryptionSettingsDefinitionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetEncryptionSettingsDefinitionIDOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionSettingsDefinitionID) {
		return nil, false
	}
	return o.EncryptionSettingsDefinitionID, true
}

// HasEncryptionSettingsDefinitionID returns a boolean if a field has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) HasEncryptionSettingsDefinitionID() bool {
	if o != nil && !IsNil(o.EncryptionSettingsDefinitionID) {
		return true
	}

	return false
}

// SetEncryptionSettingsDefinitionID gets a reference to the given string and assigns it to the EncryptionSettingsDefinitionID field.
func (o *CommonLogFileHttpOperationLogPublisherResponse) SetEncryptionSettingsDefinitionID(v string) {
	o.EncryptionSettingsDefinitionID = &v
}

// GetAppend returns the Append field value if set, zero value otherwise.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetAppend() bool {
	if o == nil || IsNil(o.Append) {
		var ret bool
		return ret
	}
	return *o.Append
}

// GetAppendOk returns a tuple with the Append field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetAppendOk() (*bool, bool) {
	if o == nil || IsNil(o.Append) {
		return nil, false
	}
	return o.Append, true
}

// HasAppend returns a boolean if a field has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) HasAppend() bool {
	if o != nil && !IsNil(o.Append) {
		return true
	}

	return false
}

// SetAppend gets a reference to the given bool and assigns it to the Append field.
func (o *CommonLogFileHttpOperationLogPublisherResponse) SetAppend(v bool) {
	o.Append = &v
}

// GetAsynchronous returns the Asynchronous field value
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetAsynchronous() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Asynchronous
}

// GetAsynchronousOk returns a tuple with the Asynchronous field value
// and a boolean to check if the value has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetAsynchronousOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asynchronous, true
}

// SetAsynchronous sets field value
func (o *CommonLogFileHttpOperationLogPublisherResponse) SetAsynchronous(v bool) {
	o.Asynchronous = v
}

// GetAutoFlush returns the AutoFlush field value if set, zero value otherwise.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetAutoFlush() bool {
	if o == nil || IsNil(o.AutoFlush) {
		var ret bool
		return ret
	}
	return *o.AutoFlush
}

// GetAutoFlushOk returns a tuple with the AutoFlush field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetAutoFlushOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoFlush) {
		return nil, false
	}
	return o.AutoFlush, true
}

// HasAutoFlush returns a boolean if a field has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) HasAutoFlush() bool {
	if o != nil && !IsNil(o.AutoFlush) {
		return true
	}

	return false
}

// SetAutoFlush gets a reference to the given bool and assigns it to the AutoFlush field.
func (o *CommonLogFileHttpOperationLogPublisherResponse) SetAutoFlush(v bool) {
	o.AutoFlush = &v
}

// GetBufferSize returns the BufferSize field value if set, zero value otherwise.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetBufferSize() string {
	if o == nil || IsNil(o.BufferSize) {
		var ret string
		return ret
	}
	return *o.BufferSize
}

// GetBufferSizeOk returns a tuple with the BufferSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetBufferSizeOk() (*string, bool) {
	if o == nil || IsNil(o.BufferSize) {
		return nil, false
	}
	return o.BufferSize, true
}

// HasBufferSize returns a boolean if a field has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) HasBufferSize() bool {
	if o != nil && !IsNil(o.BufferSize) {
		return true
	}

	return false
}

// SetBufferSize gets a reference to the given string and assigns it to the BufferSize field.
func (o *CommonLogFileHttpOperationLogPublisherResponse) SetBufferSize(v string) {
	o.BufferSize = &v
}

// GetQueueSize returns the QueueSize field value if set, zero value otherwise.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetQueueSize() int32 {
	if o == nil || IsNil(o.QueueSize) {
		var ret int32
		return ret
	}
	return *o.QueueSize
}

// GetQueueSizeOk returns a tuple with the QueueSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetQueueSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.QueueSize) {
		return nil, false
	}
	return o.QueueSize, true
}

// HasQueueSize returns a boolean if a field has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) HasQueueSize() bool {
	if o != nil && !IsNil(o.QueueSize) {
		return true
	}

	return false
}

// SetQueueSize gets a reference to the given int32 and assigns it to the QueueSize field.
func (o *CommonLogFileHttpOperationLogPublisherResponse) SetQueueSize(v int32) {
	o.QueueSize = &v
}

// GetTimeInterval returns the TimeInterval field value if set, zero value otherwise.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetTimeInterval() string {
	if o == nil || IsNil(o.TimeInterval) {
		var ret string
		return ret
	}
	return *o.TimeInterval
}

// GetTimeIntervalOk returns a tuple with the TimeInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetTimeIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.TimeInterval) {
		return nil, false
	}
	return o.TimeInterval, true
}

// HasTimeInterval returns a boolean if a field has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) HasTimeInterval() bool {
	if o != nil && !IsNil(o.TimeInterval) {
		return true
	}

	return false
}

// SetTimeInterval gets a reference to the given string and assigns it to the TimeInterval field.
func (o *CommonLogFileHttpOperationLogPublisherResponse) SetTimeInterval(v string) {
	o.TimeInterval = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CommonLogFileHttpOperationLogPublisherResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CommonLogFileHttpOperationLogPublisherResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetLoggingErrorBehavior returns the LoggingErrorBehavior field value if set, zero value otherwise.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp {
	if o == nil || IsNil(o.LoggingErrorBehavior) {
		var ret EnumlogPublisherLoggingErrorBehaviorProp
		return ret
	}
	return *o.LoggingErrorBehavior
}

// GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool) {
	if o == nil || IsNil(o.LoggingErrorBehavior) {
		return nil, false
	}
	return o.LoggingErrorBehavior, true
}

// HasLoggingErrorBehavior returns a boolean if a field has been set.
func (o *CommonLogFileHttpOperationLogPublisherResponse) HasLoggingErrorBehavior() bool {
	if o != nil && !IsNil(o.LoggingErrorBehavior) {
		return true
	}

	return false
}

// SetLoggingErrorBehavior gets a reference to the given EnumlogPublisherLoggingErrorBehaviorProp and assigns it to the LoggingErrorBehavior field.
func (o *CommonLogFileHttpOperationLogPublisherResponse) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp) {
	o.LoggingErrorBehavior = &v
}

func (o CommonLogFileHttpOperationLogPublisherResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonLogFileHttpOperationLogPublisherResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	toSerialize["schemas"] = o.Schemas
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
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.LoggingErrorBehavior) {
		toSerialize["loggingErrorBehavior"] = o.LoggingErrorBehavior
	}
	return toSerialize, nil
}

type NullableCommonLogFileHttpOperationLogPublisherResponse struct {
	value *CommonLogFileHttpOperationLogPublisherResponse
	isSet bool
}

func (v NullableCommonLogFileHttpOperationLogPublisherResponse) Get() *CommonLogFileHttpOperationLogPublisherResponse {
	return v.value
}

func (v *NullableCommonLogFileHttpOperationLogPublisherResponse) Set(val *CommonLogFileHttpOperationLogPublisherResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonLogFileHttpOperationLogPublisherResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonLogFileHttpOperationLogPublisherResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonLogFileHttpOperationLogPublisherResponse(val *CommonLogFileHttpOperationLogPublisherResponse) *NullableCommonLogFileHttpOperationLogPublisherResponse {
	return &NullableCommonLogFileHttpOperationLogPublisherResponse{value: val, isSet: true}
}

func (v NullableCommonLogFileHttpOperationLogPublisherResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonLogFileHttpOperationLogPublisherResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
