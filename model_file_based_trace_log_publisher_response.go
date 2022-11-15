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

// FileBasedTraceLogPublisherResponse struct for FileBasedTraceLogPublisherResponse
type FileBasedTraceLogPublisherResponse struct {
	// Name of the Log Publisher
	Id string `json:"id"`
	Schemas []EnumfileBasedTraceLogPublisherSchemaUrn `json:"schemas"`
	// The file name to use for the log files generated by the File Based Trace Log Publisher. The path to the file can be specified either as relative to the server root or as an absolute path.
	LogFile string `json:"logFile"`
	// The UNIX permissions of the log files created by this File Based Trace Log Publisher.
	LogFilePermissions string `json:"logFilePermissions"`
	RotationPolicy []string `json:"rotationPolicy"`
	RotationListener []string `json:"rotationListener,omitempty"`
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
	// Specifies the log file buffer size.
	BufferSize *string `json:"bufferSize,omitempty"`
	// Specifies the interval at which to check whether the log files need to be rotated.
	TimeInterval *string `json:"timeInterval,omitempty"`
	// Indicates whether the Writer Based Trace Log Publisher will publish records asynchronously.
	Asynchronous bool `json:"asynchronous"`
	// The maximum number of log records that can be stored in the asynchronous queue.
	QueueSize *int32 `json:"queueSize,omitempty"`
	// Specifies the maximum number of characters that may be included in any string in a log message before that string is truncated and replaced with a placeholder indicating the number of characters that were omitted. This can help prevent extremely long log messages from being written.
	MaxStringLength *int32 `json:"maxStringLength,omitempty"`
	DebugMessageType []EnumlogPublisherDebugMessageTypeProp `json:"debugMessageType,omitempty"`
	HttpMessageType []EnumlogPublisherHttpMessageTypeProp `json:"httpMessageType,omitempty"`
	AccessTokenValidatorMessageType []EnumlogPublisherAccessTokenValidatorMessageTypeProp `json:"accessTokenValidatorMessageType,omitempty"`
	IdTokenValidatorMessageType []EnumlogPublisherIdTokenValidatorMessageTypeProp `json:"idTokenValidatorMessageType,omitempty"`
	ScimMessageType []EnumlogPublisherScimMessageTypeProp `json:"scimMessageType,omitempty"`
	ConsentMessageType []EnumlogPublisherConsentMessageTypeProp `json:"consentMessageType,omitempty"`
	DirectoryRESTAPIMessageType []EnumlogPublisherDirectoryRESTAPIMessageTypeProp `json:"directoryRESTAPIMessageType,omitempty"`
	ExtensionMessageType []EnumlogPublisherExtensionMessageTypeProp `json:"extensionMessageType,omitempty"`
	IncludePathPattern []string `json:"includePathPattern,omitempty"`
	ExcludePathPattern []string `json:"excludePathPattern,omitempty"`
	// A description for this Log Publisher
	Description *string `json:"description,omitempty"`
	// Indicates whether the Log Publisher is enabled for use.
	Enabled bool `json:"enabled"`
	LoggingErrorBehavior *EnumlogPublisherLoggingErrorBehaviorProp `json:"loggingErrorBehavior,omitempty"`
}

// NewFileBasedTraceLogPublisherResponse instantiates a new FileBasedTraceLogPublisherResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileBasedTraceLogPublisherResponse(id string, schemas []EnumfileBasedTraceLogPublisherSchemaUrn, logFile string, logFilePermissions string, rotationPolicy []string, retentionPolicy []string, asynchronous bool, enabled bool) *FileBasedTraceLogPublisherResponse {
	this := FileBasedTraceLogPublisherResponse{}
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

// NewFileBasedTraceLogPublisherResponseWithDefaults instantiates a new FileBasedTraceLogPublisherResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileBasedTraceLogPublisherResponseWithDefaults() *FileBasedTraceLogPublisherResponse {
	this := FileBasedTraceLogPublisherResponse{}
	return &this
}

// GetId returns the Id field value
func (o *FileBasedTraceLogPublisherResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FileBasedTraceLogPublisherResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *FileBasedTraceLogPublisherResponse) GetSchemas() []EnumfileBasedTraceLogPublisherSchemaUrn {
	if o == nil {
		var ret []EnumfileBasedTraceLogPublisherSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetSchemasOk() ([]EnumfileBasedTraceLogPublisherSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *FileBasedTraceLogPublisherResponse) SetSchemas(v []EnumfileBasedTraceLogPublisherSchemaUrn) {
	o.Schemas = v
}

// GetLogFile returns the LogFile field value
func (o *FileBasedTraceLogPublisherResponse) GetLogFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogFile
}

// GetLogFileOk returns a tuple with the LogFile field value
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetLogFileOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LogFile, true
}

// SetLogFile sets field value
func (o *FileBasedTraceLogPublisherResponse) SetLogFile(v string) {
	o.LogFile = v
}

// GetLogFilePermissions returns the LogFilePermissions field value
func (o *FileBasedTraceLogPublisherResponse) GetLogFilePermissions() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogFilePermissions
}

// GetLogFilePermissionsOk returns a tuple with the LogFilePermissions field value
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetLogFilePermissionsOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LogFilePermissions, true
}

// SetLogFilePermissions sets field value
func (o *FileBasedTraceLogPublisherResponse) SetLogFilePermissions(v string) {
	o.LogFilePermissions = v
}

// GetRotationPolicy returns the RotationPolicy field value
func (o *FileBasedTraceLogPublisherResponse) GetRotationPolicy() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RotationPolicy
}

// GetRotationPolicyOk returns a tuple with the RotationPolicy field value
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetRotationPolicyOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RotationPolicy, true
}

// SetRotationPolicy sets field value
func (o *FileBasedTraceLogPublisherResponse) SetRotationPolicy(v []string) {
	o.RotationPolicy = v
}

// GetRotationListener returns the RotationListener field value if set, zero value otherwise.
func (o *FileBasedTraceLogPublisherResponse) GetRotationListener() []string {
	if o == nil || isNil(o.RotationListener) {
		var ret []string
		return ret
	}
	return o.RotationListener
}

// GetRotationListenerOk returns a tuple with the RotationListener field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetRotationListenerOk() ([]string, bool) {
	if o == nil || isNil(o.RotationListener) {
    return nil, false
	}
	return o.RotationListener, true
}

// HasRotationListener returns a boolean if a field has been set.
func (o *FileBasedTraceLogPublisherResponse) HasRotationListener() bool {
	if o != nil && !isNil(o.RotationListener) {
		return true
	}

	return false
}

// SetRotationListener gets a reference to the given []string and assigns it to the RotationListener field.
func (o *FileBasedTraceLogPublisherResponse) SetRotationListener(v []string) {
	o.RotationListener = v
}

// GetRetentionPolicy returns the RetentionPolicy field value
func (o *FileBasedTraceLogPublisherResponse) GetRetentionPolicy() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RetentionPolicy
}

// GetRetentionPolicyOk returns a tuple with the RetentionPolicy field value
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetRetentionPolicyOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RetentionPolicy, true
}

// SetRetentionPolicy sets field value
func (o *FileBasedTraceLogPublisherResponse) SetRetentionPolicy(v []string) {
	o.RetentionPolicy = v
}

// GetCompressionMechanism returns the CompressionMechanism field value if set, zero value otherwise.
func (o *FileBasedTraceLogPublisherResponse) GetCompressionMechanism() EnumlogPublisherCompressionMechanismProp {
	if o == nil || isNil(o.CompressionMechanism) {
		var ret EnumlogPublisherCompressionMechanismProp
		return ret
	}
	return *o.CompressionMechanism
}

// GetCompressionMechanismOk returns a tuple with the CompressionMechanism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetCompressionMechanismOk() (*EnumlogPublisherCompressionMechanismProp, bool) {
	if o == nil || isNil(o.CompressionMechanism) {
    return nil, false
	}
	return o.CompressionMechanism, true
}

// HasCompressionMechanism returns a boolean if a field has been set.
func (o *FileBasedTraceLogPublisherResponse) HasCompressionMechanism() bool {
	if o != nil && !isNil(o.CompressionMechanism) {
		return true
	}

	return false
}

// SetCompressionMechanism gets a reference to the given EnumlogPublisherCompressionMechanismProp and assigns it to the CompressionMechanism field.
func (o *FileBasedTraceLogPublisherResponse) SetCompressionMechanism(v EnumlogPublisherCompressionMechanismProp) {
	o.CompressionMechanism = &v
}

// GetSignLog returns the SignLog field value if set, zero value otherwise.
func (o *FileBasedTraceLogPublisherResponse) GetSignLog() bool {
	if o == nil || isNil(o.SignLog) {
		var ret bool
		return ret
	}
	return *o.SignLog
}

// GetSignLogOk returns a tuple with the SignLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetSignLogOk() (*bool, bool) {
	if o == nil || isNil(o.SignLog) {
    return nil, false
	}
	return o.SignLog, true
}

// HasSignLog returns a boolean if a field has been set.
func (o *FileBasedTraceLogPublisherResponse) HasSignLog() bool {
	if o != nil && !isNil(o.SignLog) {
		return true
	}

	return false
}

// SetSignLog gets a reference to the given bool and assigns it to the SignLog field.
func (o *FileBasedTraceLogPublisherResponse) SetSignLog(v bool) {
	o.SignLog = &v
}

// GetEncryptLog returns the EncryptLog field value if set, zero value otherwise.
func (o *FileBasedTraceLogPublisherResponse) GetEncryptLog() bool {
	if o == nil || isNil(o.EncryptLog) {
		var ret bool
		return ret
	}
	return *o.EncryptLog
}

// GetEncryptLogOk returns a tuple with the EncryptLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetEncryptLogOk() (*bool, bool) {
	if o == nil || isNil(o.EncryptLog) {
    return nil, false
	}
	return o.EncryptLog, true
}

// HasEncryptLog returns a boolean if a field has been set.
func (o *FileBasedTraceLogPublisherResponse) HasEncryptLog() bool {
	if o != nil && !isNil(o.EncryptLog) {
		return true
	}

	return false
}

// SetEncryptLog gets a reference to the given bool and assigns it to the EncryptLog field.
func (o *FileBasedTraceLogPublisherResponse) SetEncryptLog(v bool) {
	o.EncryptLog = &v
}

// GetEncryptionSettingsDefinitionID returns the EncryptionSettingsDefinitionID field value if set, zero value otherwise.
func (o *FileBasedTraceLogPublisherResponse) GetEncryptionSettingsDefinitionID() string {
	if o == nil || isNil(o.EncryptionSettingsDefinitionID) {
		var ret string
		return ret
	}
	return *o.EncryptionSettingsDefinitionID
}

// GetEncryptionSettingsDefinitionIDOk returns a tuple with the EncryptionSettingsDefinitionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetEncryptionSettingsDefinitionIDOk() (*string, bool) {
	if o == nil || isNil(o.EncryptionSettingsDefinitionID) {
    return nil, false
	}
	return o.EncryptionSettingsDefinitionID, true
}

// HasEncryptionSettingsDefinitionID returns a boolean if a field has been set.
func (o *FileBasedTraceLogPublisherResponse) HasEncryptionSettingsDefinitionID() bool {
	if o != nil && !isNil(o.EncryptionSettingsDefinitionID) {
		return true
	}

	return false
}

// SetEncryptionSettingsDefinitionID gets a reference to the given string and assigns it to the EncryptionSettingsDefinitionID field.
func (o *FileBasedTraceLogPublisherResponse) SetEncryptionSettingsDefinitionID(v string) {
	o.EncryptionSettingsDefinitionID = &v
}

// GetAppend returns the Append field value if set, zero value otherwise.
func (o *FileBasedTraceLogPublisherResponse) GetAppend() bool {
	if o == nil || isNil(o.Append) {
		var ret bool
		return ret
	}
	return *o.Append
}

// GetAppendOk returns a tuple with the Append field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetAppendOk() (*bool, bool) {
	if o == nil || isNil(o.Append) {
    return nil, false
	}
	return o.Append, true
}

// HasAppend returns a boolean if a field has been set.
func (o *FileBasedTraceLogPublisherResponse) HasAppend() bool {
	if o != nil && !isNil(o.Append) {
		return true
	}

	return false
}

// SetAppend gets a reference to the given bool and assigns it to the Append field.
func (o *FileBasedTraceLogPublisherResponse) SetAppend(v bool) {
	o.Append = &v
}

// GetBufferSize returns the BufferSize field value if set, zero value otherwise.
func (o *FileBasedTraceLogPublisherResponse) GetBufferSize() string {
	if o == nil || isNil(o.BufferSize) {
		var ret string
		return ret
	}
	return *o.BufferSize
}

// GetBufferSizeOk returns a tuple with the BufferSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetBufferSizeOk() (*string, bool) {
	if o == nil || isNil(o.BufferSize) {
    return nil, false
	}
	return o.BufferSize, true
}

// HasBufferSize returns a boolean if a field has been set.
func (o *FileBasedTraceLogPublisherResponse) HasBufferSize() bool {
	if o != nil && !isNil(o.BufferSize) {
		return true
	}

	return false
}

// SetBufferSize gets a reference to the given string and assigns it to the BufferSize field.
func (o *FileBasedTraceLogPublisherResponse) SetBufferSize(v string) {
	o.BufferSize = &v
}

// GetTimeInterval returns the TimeInterval field value if set, zero value otherwise.
func (o *FileBasedTraceLogPublisherResponse) GetTimeInterval() string {
	if o == nil || isNil(o.TimeInterval) {
		var ret string
		return ret
	}
	return *o.TimeInterval
}

// GetTimeIntervalOk returns a tuple with the TimeInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetTimeIntervalOk() (*string, bool) {
	if o == nil || isNil(o.TimeInterval) {
    return nil, false
	}
	return o.TimeInterval, true
}

// HasTimeInterval returns a boolean if a field has been set.
func (o *FileBasedTraceLogPublisherResponse) HasTimeInterval() bool {
	if o != nil && !isNil(o.TimeInterval) {
		return true
	}

	return false
}

// SetTimeInterval gets a reference to the given string and assigns it to the TimeInterval field.
func (o *FileBasedTraceLogPublisherResponse) SetTimeInterval(v string) {
	o.TimeInterval = &v
}

// GetAsynchronous returns the Asynchronous field value
func (o *FileBasedTraceLogPublisherResponse) GetAsynchronous() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Asynchronous
}

// GetAsynchronousOk returns a tuple with the Asynchronous field value
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetAsynchronousOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Asynchronous, true
}

// SetAsynchronous sets field value
func (o *FileBasedTraceLogPublisherResponse) SetAsynchronous(v bool) {
	o.Asynchronous = v
}

// GetQueueSize returns the QueueSize field value if set, zero value otherwise.
func (o *FileBasedTraceLogPublisherResponse) GetQueueSize() int32 {
	if o == nil || isNil(o.QueueSize) {
		var ret int32
		return ret
	}
	return *o.QueueSize
}

// GetQueueSizeOk returns a tuple with the QueueSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetQueueSizeOk() (*int32, bool) {
	if o == nil || isNil(o.QueueSize) {
    return nil, false
	}
	return o.QueueSize, true
}

// HasQueueSize returns a boolean if a field has been set.
func (o *FileBasedTraceLogPublisherResponse) HasQueueSize() bool {
	if o != nil && !isNil(o.QueueSize) {
		return true
	}

	return false
}

// SetQueueSize gets a reference to the given int32 and assigns it to the QueueSize field.
func (o *FileBasedTraceLogPublisherResponse) SetQueueSize(v int32) {
	o.QueueSize = &v
}

// GetMaxStringLength returns the MaxStringLength field value if set, zero value otherwise.
func (o *FileBasedTraceLogPublisherResponse) GetMaxStringLength() int32 {
	if o == nil || isNil(o.MaxStringLength) {
		var ret int32
		return ret
	}
	return *o.MaxStringLength
}

// GetMaxStringLengthOk returns a tuple with the MaxStringLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetMaxStringLengthOk() (*int32, bool) {
	if o == nil || isNil(o.MaxStringLength) {
    return nil, false
	}
	return o.MaxStringLength, true
}

// HasMaxStringLength returns a boolean if a field has been set.
func (o *FileBasedTraceLogPublisherResponse) HasMaxStringLength() bool {
	if o != nil && !isNil(o.MaxStringLength) {
		return true
	}

	return false
}

// SetMaxStringLength gets a reference to the given int32 and assigns it to the MaxStringLength field.
func (o *FileBasedTraceLogPublisherResponse) SetMaxStringLength(v int32) {
	o.MaxStringLength = &v
}

// GetDebugMessageType returns the DebugMessageType field value if set, zero value otherwise.
func (o *FileBasedTraceLogPublisherResponse) GetDebugMessageType() []EnumlogPublisherDebugMessageTypeProp {
	if o == nil || isNil(o.DebugMessageType) {
		var ret []EnumlogPublisherDebugMessageTypeProp
		return ret
	}
	return o.DebugMessageType
}

// GetDebugMessageTypeOk returns a tuple with the DebugMessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetDebugMessageTypeOk() ([]EnumlogPublisherDebugMessageTypeProp, bool) {
	if o == nil || isNil(o.DebugMessageType) {
    return nil, false
	}
	return o.DebugMessageType, true
}

// HasDebugMessageType returns a boolean if a field has been set.
func (o *FileBasedTraceLogPublisherResponse) HasDebugMessageType() bool {
	if o != nil && !isNil(o.DebugMessageType) {
		return true
	}

	return false
}

// SetDebugMessageType gets a reference to the given []EnumlogPublisherDebugMessageTypeProp and assigns it to the DebugMessageType field.
func (o *FileBasedTraceLogPublisherResponse) SetDebugMessageType(v []EnumlogPublisherDebugMessageTypeProp) {
	o.DebugMessageType = v
}

// GetHttpMessageType returns the HttpMessageType field value if set, zero value otherwise.
func (o *FileBasedTraceLogPublisherResponse) GetHttpMessageType() []EnumlogPublisherHttpMessageTypeProp {
	if o == nil || isNil(o.HttpMessageType) {
		var ret []EnumlogPublisherHttpMessageTypeProp
		return ret
	}
	return o.HttpMessageType
}

// GetHttpMessageTypeOk returns a tuple with the HttpMessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetHttpMessageTypeOk() ([]EnumlogPublisherHttpMessageTypeProp, bool) {
	if o == nil || isNil(o.HttpMessageType) {
    return nil, false
	}
	return o.HttpMessageType, true
}

// HasHttpMessageType returns a boolean if a field has been set.
func (o *FileBasedTraceLogPublisherResponse) HasHttpMessageType() bool {
	if o != nil && !isNil(o.HttpMessageType) {
		return true
	}

	return false
}

// SetHttpMessageType gets a reference to the given []EnumlogPublisherHttpMessageTypeProp and assigns it to the HttpMessageType field.
func (o *FileBasedTraceLogPublisherResponse) SetHttpMessageType(v []EnumlogPublisherHttpMessageTypeProp) {
	o.HttpMessageType = v
}

// GetAccessTokenValidatorMessageType returns the AccessTokenValidatorMessageType field value if set, zero value otherwise.
func (o *FileBasedTraceLogPublisherResponse) GetAccessTokenValidatorMessageType() []EnumlogPublisherAccessTokenValidatorMessageTypeProp {
	if o == nil || isNil(o.AccessTokenValidatorMessageType) {
		var ret []EnumlogPublisherAccessTokenValidatorMessageTypeProp
		return ret
	}
	return o.AccessTokenValidatorMessageType
}

// GetAccessTokenValidatorMessageTypeOk returns a tuple with the AccessTokenValidatorMessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetAccessTokenValidatorMessageTypeOk() ([]EnumlogPublisherAccessTokenValidatorMessageTypeProp, bool) {
	if o == nil || isNil(o.AccessTokenValidatorMessageType) {
    return nil, false
	}
	return o.AccessTokenValidatorMessageType, true
}

// HasAccessTokenValidatorMessageType returns a boolean if a field has been set.
func (o *FileBasedTraceLogPublisherResponse) HasAccessTokenValidatorMessageType() bool {
	if o != nil && !isNil(o.AccessTokenValidatorMessageType) {
		return true
	}

	return false
}

// SetAccessTokenValidatorMessageType gets a reference to the given []EnumlogPublisherAccessTokenValidatorMessageTypeProp and assigns it to the AccessTokenValidatorMessageType field.
func (o *FileBasedTraceLogPublisherResponse) SetAccessTokenValidatorMessageType(v []EnumlogPublisherAccessTokenValidatorMessageTypeProp) {
	o.AccessTokenValidatorMessageType = v
}

// GetIdTokenValidatorMessageType returns the IdTokenValidatorMessageType field value if set, zero value otherwise.
func (o *FileBasedTraceLogPublisherResponse) GetIdTokenValidatorMessageType() []EnumlogPublisherIdTokenValidatorMessageTypeProp {
	if o == nil || isNil(o.IdTokenValidatorMessageType) {
		var ret []EnumlogPublisherIdTokenValidatorMessageTypeProp
		return ret
	}
	return o.IdTokenValidatorMessageType
}

// GetIdTokenValidatorMessageTypeOk returns a tuple with the IdTokenValidatorMessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetIdTokenValidatorMessageTypeOk() ([]EnumlogPublisherIdTokenValidatorMessageTypeProp, bool) {
	if o == nil || isNil(o.IdTokenValidatorMessageType) {
    return nil, false
	}
	return o.IdTokenValidatorMessageType, true
}

// HasIdTokenValidatorMessageType returns a boolean if a field has been set.
func (o *FileBasedTraceLogPublisherResponse) HasIdTokenValidatorMessageType() bool {
	if o != nil && !isNil(o.IdTokenValidatorMessageType) {
		return true
	}

	return false
}

// SetIdTokenValidatorMessageType gets a reference to the given []EnumlogPublisherIdTokenValidatorMessageTypeProp and assigns it to the IdTokenValidatorMessageType field.
func (o *FileBasedTraceLogPublisherResponse) SetIdTokenValidatorMessageType(v []EnumlogPublisherIdTokenValidatorMessageTypeProp) {
	o.IdTokenValidatorMessageType = v
}

// GetScimMessageType returns the ScimMessageType field value if set, zero value otherwise.
func (o *FileBasedTraceLogPublisherResponse) GetScimMessageType() []EnumlogPublisherScimMessageTypeProp {
	if o == nil || isNil(o.ScimMessageType) {
		var ret []EnumlogPublisherScimMessageTypeProp
		return ret
	}
	return o.ScimMessageType
}

// GetScimMessageTypeOk returns a tuple with the ScimMessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetScimMessageTypeOk() ([]EnumlogPublisherScimMessageTypeProp, bool) {
	if o == nil || isNil(o.ScimMessageType) {
    return nil, false
	}
	return o.ScimMessageType, true
}

// HasScimMessageType returns a boolean if a field has been set.
func (o *FileBasedTraceLogPublisherResponse) HasScimMessageType() bool {
	if o != nil && !isNil(o.ScimMessageType) {
		return true
	}

	return false
}

// SetScimMessageType gets a reference to the given []EnumlogPublisherScimMessageTypeProp and assigns it to the ScimMessageType field.
func (o *FileBasedTraceLogPublisherResponse) SetScimMessageType(v []EnumlogPublisherScimMessageTypeProp) {
	o.ScimMessageType = v
}

// GetConsentMessageType returns the ConsentMessageType field value if set, zero value otherwise.
func (o *FileBasedTraceLogPublisherResponse) GetConsentMessageType() []EnumlogPublisherConsentMessageTypeProp {
	if o == nil || isNil(o.ConsentMessageType) {
		var ret []EnumlogPublisherConsentMessageTypeProp
		return ret
	}
	return o.ConsentMessageType
}

// GetConsentMessageTypeOk returns a tuple with the ConsentMessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetConsentMessageTypeOk() ([]EnumlogPublisherConsentMessageTypeProp, bool) {
	if o == nil || isNil(o.ConsentMessageType) {
    return nil, false
	}
	return o.ConsentMessageType, true
}

// HasConsentMessageType returns a boolean if a field has been set.
func (o *FileBasedTraceLogPublisherResponse) HasConsentMessageType() bool {
	if o != nil && !isNil(o.ConsentMessageType) {
		return true
	}

	return false
}

// SetConsentMessageType gets a reference to the given []EnumlogPublisherConsentMessageTypeProp and assigns it to the ConsentMessageType field.
func (o *FileBasedTraceLogPublisherResponse) SetConsentMessageType(v []EnumlogPublisherConsentMessageTypeProp) {
	o.ConsentMessageType = v
}

// GetDirectoryRESTAPIMessageType returns the DirectoryRESTAPIMessageType field value if set, zero value otherwise.
func (o *FileBasedTraceLogPublisherResponse) GetDirectoryRESTAPIMessageType() []EnumlogPublisherDirectoryRESTAPIMessageTypeProp {
	if o == nil || isNil(o.DirectoryRESTAPIMessageType) {
		var ret []EnumlogPublisherDirectoryRESTAPIMessageTypeProp
		return ret
	}
	return o.DirectoryRESTAPIMessageType
}

// GetDirectoryRESTAPIMessageTypeOk returns a tuple with the DirectoryRESTAPIMessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetDirectoryRESTAPIMessageTypeOk() ([]EnumlogPublisherDirectoryRESTAPIMessageTypeProp, bool) {
	if o == nil || isNil(o.DirectoryRESTAPIMessageType) {
    return nil, false
	}
	return o.DirectoryRESTAPIMessageType, true
}

// HasDirectoryRESTAPIMessageType returns a boolean if a field has been set.
func (o *FileBasedTraceLogPublisherResponse) HasDirectoryRESTAPIMessageType() bool {
	if o != nil && !isNil(o.DirectoryRESTAPIMessageType) {
		return true
	}

	return false
}

// SetDirectoryRESTAPIMessageType gets a reference to the given []EnumlogPublisherDirectoryRESTAPIMessageTypeProp and assigns it to the DirectoryRESTAPIMessageType field.
func (o *FileBasedTraceLogPublisherResponse) SetDirectoryRESTAPIMessageType(v []EnumlogPublisherDirectoryRESTAPIMessageTypeProp) {
	o.DirectoryRESTAPIMessageType = v
}

// GetExtensionMessageType returns the ExtensionMessageType field value if set, zero value otherwise.
func (o *FileBasedTraceLogPublisherResponse) GetExtensionMessageType() []EnumlogPublisherExtensionMessageTypeProp {
	if o == nil || isNil(o.ExtensionMessageType) {
		var ret []EnumlogPublisherExtensionMessageTypeProp
		return ret
	}
	return o.ExtensionMessageType
}

// GetExtensionMessageTypeOk returns a tuple with the ExtensionMessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetExtensionMessageTypeOk() ([]EnumlogPublisherExtensionMessageTypeProp, bool) {
	if o == nil || isNil(o.ExtensionMessageType) {
    return nil, false
	}
	return o.ExtensionMessageType, true
}

// HasExtensionMessageType returns a boolean if a field has been set.
func (o *FileBasedTraceLogPublisherResponse) HasExtensionMessageType() bool {
	if o != nil && !isNil(o.ExtensionMessageType) {
		return true
	}

	return false
}

// SetExtensionMessageType gets a reference to the given []EnumlogPublisherExtensionMessageTypeProp and assigns it to the ExtensionMessageType field.
func (o *FileBasedTraceLogPublisherResponse) SetExtensionMessageType(v []EnumlogPublisherExtensionMessageTypeProp) {
	o.ExtensionMessageType = v
}

// GetIncludePathPattern returns the IncludePathPattern field value if set, zero value otherwise.
func (o *FileBasedTraceLogPublisherResponse) GetIncludePathPattern() []string {
	if o == nil || isNil(o.IncludePathPattern) {
		var ret []string
		return ret
	}
	return o.IncludePathPattern
}

// GetIncludePathPatternOk returns a tuple with the IncludePathPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetIncludePathPatternOk() ([]string, bool) {
	if o == nil || isNil(o.IncludePathPattern) {
    return nil, false
	}
	return o.IncludePathPattern, true
}

// HasIncludePathPattern returns a boolean if a field has been set.
func (o *FileBasedTraceLogPublisherResponse) HasIncludePathPattern() bool {
	if o != nil && !isNil(o.IncludePathPattern) {
		return true
	}

	return false
}

// SetIncludePathPattern gets a reference to the given []string and assigns it to the IncludePathPattern field.
func (o *FileBasedTraceLogPublisherResponse) SetIncludePathPattern(v []string) {
	o.IncludePathPattern = v
}

// GetExcludePathPattern returns the ExcludePathPattern field value if set, zero value otherwise.
func (o *FileBasedTraceLogPublisherResponse) GetExcludePathPattern() []string {
	if o == nil || isNil(o.ExcludePathPattern) {
		var ret []string
		return ret
	}
	return o.ExcludePathPattern
}

// GetExcludePathPatternOk returns a tuple with the ExcludePathPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetExcludePathPatternOk() ([]string, bool) {
	if o == nil || isNil(o.ExcludePathPattern) {
    return nil, false
	}
	return o.ExcludePathPattern, true
}

// HasExcludePathPattern returns a boolean if a field has been set.
func (o *FileBasedTraceLogPublisherResponse) HasExcludePathPattern() bool {
	if o != nil && !isNil(o.ExcludePathPattern) {
		return true
	}

	return false
}

// SetExcludePathPattern gets a reference to the given []string and assigns it to the ExcludePathPattern field.
func (o *FileBasedTraceLogPublisherResponse) SetExcludePathPattern(v []string) {
	o.ExcludePathPattern = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FileBasedTraceLogPublisherResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FileBasedTraceLogPublisherResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FileBasedTraceLogPublisherResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *FileBasedTraceLogPublisherResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *FileBasedTraceLogPublisherResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetLoggingErrorBehavior returns the LoggingErrorBehavior field value if set, zero value otherwise.
func (o *FileBasedTraceLogPublisherResponse) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp {
	if o == nil || isNil(o.LoggingErrorBehavior) {
		var ret EnumlogPublisherLoggingErrorBehaviorProp
		return ret
	}
	return *o.LoggingErrorBehavior
}

// GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileBasedTraceLogPublisherResponse) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool) {
	if o == nil || isNil(o.LoggingErrorBehavior) {
    return nil, false
	}
	return o.LoggingErrorBehavior, true
}

// HasLoggingErrorBehavior returns a boolean if a field has been set.
func (o *FileBasedTraceLogPublisherResponse) HasLoggingErrorBehavior() bool {
	if o != nil && !isNil(o.LoggingErrorBehavior) {
		return true
	}

	return false
}

// SetLoggingErrorBehavior gets a reference to the given EnumlogPublisherLoggingErrorBehaviorProp and assigns it to the LoggingErrorBehavior field.
func (o *FileBasedTraceLogPublisherResponse) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp) {
	o.LoggingErrorBehavior = &v
}

func (o FileBasedTraceLogPublisherResponse) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.BufferSize) {
		toSerialize["bufferSize"] = o.BufferSize
	}
	if !isNil(o.TimeInterval) {
		toSerialize["timeInterval"] = o.TimeInterval
	}
	if true {
		toSerialize["asynchronous"] = o.Asynchronous
	}
	if !isNil(o.QueueSize) {
		toSerialize["queueSize"] = o.QueueSize
	}
	if !isNil(o.MaxStringLength) {
		toSerialize["maxStringLength"] = o.MaxStringLength
	}
	if !isNil(o.DebugMessageType) {
		toSerialize["debugMessageType"] = o.DebugMessageType
	}
	if !isNil(o.HttpMessageType) {
		toSerialize["httpMessageType"] = o.HttpMessageType
	}
	if !isNil(o.AccessTokenValidatorMessageType) {
		toSerialize["accessTokenValidatorMessageType"] = o.AccessTokenValidatorMessageType
	}
	if !isNil(o.IdTokenValidatorMessageType) {
		toSerialize["idTokenValidatorMessageType"] = o.IdTokenValidatorMessageType
	}
	if !isNil(o.ScimMessageType) {
		toSerialize["scimMessageType"] = o.ScimMessageType
	}
	if !isNil(o.ConsentMessageType) {
		toSerialize["consentMessageType"] = o.ConsentMessageType
	}
	if !isNil(o.DirectoryRESTAPIMessageType) {
		toSerialize["directoryRESTAPIMessageType"] = o.DirectoryRESTAPIMessageType
	}
	if !isNil(o.ExtensionMessageType) {
		toSerialize["extensionMessageType"] = o.ExtensionMessageType
	}
	if !isNil(o.IncludePathPattern) {
		toSerialize["includePathPattern"] = o.IncludePathPattern
	}
	if !isNil(o.ExcludePathPattern) {
		toSerialize["excludePathPattern"] = o.ExcludePathPattern
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

type NullableFileBasedTraceLogPublisherResponse struct {
	value *FileBasedTraceLogPublisherResponse
	isSet bool
}

func (v NullableFileBasedTraceLogPublisherResponse) Get() *FileBasedTraceLogPublisherResponse {
	return v.value
}

func (v *NullableFileBasedTraceLogPublisherResponse) Set(val *FileBasedTraceLogPublisherResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFileBasedTraceLogPublisherResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFileBasedTraceLogPublisherResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileBasedTraceLogPublisherResponse(val *FileBasedTraceLogPublisherResponse) *NullableFileBasedTraceLogPublisherResponse {
	return &NullableFileBasedTraceLogPublisherResponse{value: val, isSet: true}
}

func (v NullableFileBasedTraceLogPublisherResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileBasedTraceLogPublisherResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


