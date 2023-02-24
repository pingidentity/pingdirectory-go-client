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

// checks if the LdifExportRecurringTaskResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LdifExportRecurringTaskResponse{}

// LdifExportRecurringTaskResponse struct for LdifExportRecurringTaskResponse
type LdifExportRecurringTaskResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Recurring Task
	Id      string                                 `json:"id"`
	Schemas []EnumldifExportRecurringTaskSchemaUrn `json:"schemas"`
	// The directory in which LDIF export files will be placed. The directory must already exist.
	LdifDirectory string `json:"ldifDirectory"`
	// The backend ID for a backend to be exported.
	BackendID []string `json:"backendID,omitempty"`
	// The backend ID for a backend to be excluded from the export.
	ExcludeBackendID []string `json:"excludeBackendID,omitempty"`
	// Indicates whether to compress the LDIF data as it is exported.
	Compress *bool `json:"compress,omitempty"`
	// Indicates whether to encrypt the LDIF data as it exported.
	Encrypt *bool `json:"encrypt,omitempty"`
	// The ID of an encryption settings definition to use to obtain the LDIF export encryption key.
	EncryptionSettingsDefinitionID *string `json:"encryptionSettingsDefinitionID,omitempty"`
	// Indicates whether to cryptographically sign the exported data, which will make it possible to detect whether the LDIF data has been altered since it was exported.
	Sign *bool `json:"sign,omitempty"`
	// The minimum number of previous LDIF exports that should be preserved after a new export completes successfully.
	RetainPreviousLDIFExportCount *int32 `json:"retainPreviousLDIFExportCount,omitempty"`
	// The minimum age of previous LDIF exports that should be preserved after a new export completes successfully.
	RetainPreviousLDIFExportAge *string `json:"retainPreviousLDIFExportAge,omitempty"`
	// The maximum rate, in megabytes per second, at which LDIF exports should be written.
	MaxMegabytesPerSecond *int32 `json:"maxMegabytesPerSecond,omitempty"`
	// A description for this Recurring Task
	Description *string `json:"description,omitempty"`
	// Indicates whether an instance of this Recurring Task should be canceled if the task immediately before it in the recurring task chain fails to complete successfully (including if it is canceled by an administrator before it starts or while it is running).
	CancelOnTaskDependencyFailure *bool `json:"cancelOnTaskDependencyFailure,omitempty"`
	// The email addresses to which a message should be sent whenever an instance of this Recurring Task starts running. If this option is used, then at least one smtp-server must be configured in the global configuration.
	EmailOnStart []string `json:"emailOnStart,omitempty"`
	// The email addresses to which a message should be sent whenever an instance of this Recurring Task completes successfully. If this option is used, then at least one smtp-server must be configured in the global configuration.
	EmailOnSuccess []string `json:"emailOnSuccess,omitempty"`
	// The email addresses to which a message should be sent if an instance of this Recurring Task fails to complete successfully. If this option is used, then at least one smtp-server must be configured in the global configuration.
	EmailOnFailure []string `json:"emailOnFailure,omitempty"`
	// Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task starts running.
	AlertOnStart *bool `json:"alertOnStart,omitempty"`
	// Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task completes successfully.
	AlertOnSuccess *bool `json:"alertOnSuccess,omitempty"`
	// Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task fails to complete successfully.
	AlertOnFailure *bool `json:"alertOnFailure,omitempty"`
}

// NewLdifExportRecurringTaskResponse instantiates a new LdifExportRecurringTaskResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdifExportRecurringTaskResponse(id string, schemas []EnumldifExportRecurringTaskSchemaUrn, ldifDirectory string) *LdifExportRecurringTaskResponse {
	this := LdifExportRecurringTaskResponse{}
	this.Id = id
	this.Schemas = schemas
	this.LdifDirectory = ldifDirectory
	return &this
}

// NewLdifExportRecurringTaskResponseWithDefaults instantiates a new LdifExportRecurringTaskResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdifExportRecurringTaskResponseWithDefaults() *LdifExportRecurringTaskResponse {
	this := LdifExportRecurringTaskResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *LdifExportRecurringTaskResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdifExportRecurringTaskResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *LdifExportRecurringTaskResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *LdifExportRecurringTaskResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *LdifExportRecurringTaskResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdifExportRecurringTaskResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *LdifExportRecurringTaskResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *LdifExportRecurringTaskResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *LdifExportRecurringTaskResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LdifExportRecurringTaskResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LdifExportRecurringTaskResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *LdifExportRecurringTaskResponse) GetSchemas() []EnumldifExportRecurringTaskSchemaUrn {
	if o == nil {
		var ret []EnumldifExportRecurringTaskSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *LdifExportRecurringTaskResponse) GetSchemasOk() ([]EnumldifExportRecurringTaskSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *LdifExportRecurringTaskResponse) SetSchemas(v []EnumldifExportRecurringTaskSchemaUrn) {
	o.Schemas = v
}

// GetLdifDirectory returns the LdifDirectory field value
func (o *LdifExportRecurringTaskResponse) GetLdifDirectory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LdifDirectory
}

// GetLdifDirectoryOk returns a tuple with the LdifDirectory field value
// and a boolean to check if the value has been set.
func (o *LdifExportRecurringTaskResponse) GetLdifDirectoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LdifDirectory, true
}

// SetLdifDirectory sets field value
func (o *LdifExportRecurringTaskResponse) SetLdifDirectory(v string) {
	o.LdifDirectory = v
}

// GetBackendID returns the BackendID field value if set, zero value otherwise.
func (o *LdifExportRecurringTaskResponse) GetBackendID() []string {
	if o == nil || IsNil(o.BackendID) {
		var ret []string
		return ret
	}
	return o.BackendID
}

// GetBackendIDOk returns a tuple with the BackendID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdifExportRecurringTaskResponse) GetBackendIDOk() ([]string, bool) {
	if o == nil || IsNil(o.BackendID) {
		return nil, false
	}
	return o.BackendID, true
}

// HasBackendID returns a boolean if a field has been set.
func (o *LdifExportRecurringTaskResponse) HasBackendID() bool {
	if o != nil && !IsNil(o.BackendID) {
		return true
	}

	return false
}

// SetBackendID gets a reference to the given []string and assigns it to the BackendID field.
func (o *LdifExportRecurringTaskResponse) SetBackendID(v []string) {
	o.BackendID = v
}

// GetExcludeBackendID returns the ExcludeBackendID field value if set, zero value otherwise.
func (o *LdifExportRecurringTaskResponse) GetExcludeBackendID() []string {
	if o == nil || IsNil(o.ExcludeBackendID) {
		var ret []string
		return ret
	}
	return o.ExcludeBackendID
}

// GetExcludeBackendIDOk returns a tuple with the ExcludeBackendID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdifExportRecurringTaskResponse) GetExcludeBackendIDOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeBackendID) {
		return nil, false
	}
	return o.ExcludeBackendID, true
}

// HasExcludeBackendID returns a boolean if a field has been set.
func (o *LdifExportRecurringTaskResponse) HasExcludeBackendID() bool {
	if o != nil && !IsNil(o.ExcludeBackendID) {
		return true
	}

	return false
}

// SetExcludeBackendID gets a reference to the given []string and assigns it to the ExcludeBackendID field.
func (o *LdifExportRecurringTaskResponse) SetExcludeBackendID(v []string) {
	o.ExcludeBackendID = v
}

// GetCompress returns the Compress field value if set, zero value otherwise.
func (o *LdifExportRecurringTaskResponse) GetCompress() bool {
	if o == nil || IsNil(o.Compress) {
		var ret bool
		return ret
	}
	return *o.Compress
}

// GetCompressOk returns a tuple with the Compress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdifExportRecurringTaskResponse) GetCompressOk() (*bool, bool) {
	if o == nil || IsNil(o.Compress) {
		return nil, false
	}
	return o.Compress, true
}

// HasCompress returns a boolean if a field has been set.
func (o *LdifExportRecurringTaskResponse) HasCompress() bool {
	if o != nil && !IsNil(o.Compress) {
		return true
	}

	return false
}

// SetCompress gets a reference to the given bool and assigns it to the Compress field.
func (o *LdifExportRecurringTaskResponse) SetCompress(v bool) {
	o.Compress = &v
}

// GetEncrypt returns the Encrypt field value if set, zero value otherwise.
func (o *LdifExportRecurringTaskResponse) GetEncrypt() bool {
	if o == nil || IsNil(o.Encrypt) {
		var ret bool
		return ret
	}
	return *o.Encrypt
}

// GetEncryptOk returns a tuple with the Encrypt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdifExportRecurringTaskResponse) GetEncryptOk() (*bool, bool) {
	if o == nil || IsNil(o.Encrypt) {
		return nil, false
	}
	return o.Encrypt, true
}

// HasEncrypt returns a boolean if a field has been set.
func (o *LdifExportRecurringTaskResponse) HasEncrypt() bool {
	if o != nil && !IsNil(o.Encrypt) {
		return true
	}

	return false
}

// SetEncrypt gets a reference to the given bool and assigns it to the Encrypt field.
func (o *LdifExportRecurringTaskResponse) SetEncrypt(v bool) {
	o.Encrypt = &v
}

// GetEncryptionSettingsDefinitionID returns the EncryptionSettingsDefinitionID field value if set, zero value otherwise.
func (o *LdifExportRecurringTaskResponse) GetEncryptionSettingsDefinitionID() string {
	if o == nil || IsNil(o.EncryptionSettingsDefinitionID) {
		var ret string
		return ret
	}
	return *o.EncryptionSettingsDefinitionID
}

// GetEncryptionSettingsDefinitionIDOk returns a tuple with the EncryptionSettingsDefinitionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdifExportRecurringTaskResponse) GetEncryptionSettingsDefinitionIDOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionSettingsDefinitionID) {
		return nil, false
	}
	return o.EncryptionSettingsDefinitionID, true
}

// HasEncryptionSettingsDefinitionID returns a boolean if a field has been set.
func (o *LdifExportRecurringTaskResponse) HasEncryptionSettingsDefinitionID() bool {
	if o != nil && !IsNil(o.EncryptionSettingsDefinitionID) {
		return true
	}

	return false
}

// SetEncryptionSettingsDefinitionID gets a reference to the given string and assigns it to the EncryptionSettingsDefinitionID field.
func (o *LdifExportRecurringTaskResponse) SetEncryptionSettingsDefinitionID(v string) {
	o.EncryptionSettingsDefinitionID = &v
}

// GetSign returns the Sign field value if set, zero value otherwise.
func (o *LdifExportRecurringTaskResponse) GetSign() bool {
	if o == nil || IsNil(o.Sign) {
		var ret bool
		return ret
	}
	return *o.Sign
}

// GetSignOk returns a tuple with the Sign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdifExportRecurringTaskResponse) GetSignOk() (*bool, bool) {
	if o == nil || IsNil(o.Sign) {
		return nil, false
	}
	return o.Sign, true
}

// HasSign returns a boolean if a field has been set.
func (o *LdifExportRecurringTaskResponse) HasSign() bool {
	if o != nil && !IsNil(o.Sign) {
		return true
	}

	return false
}

// SetSign gets a reference to the given bool and assigns it to the Sign field.
func (o *LdifExportRecurringTaskResponse) SetSign(v bool) {
	o.Sign = &v
}

// GetRetainPreviousLDIFExportCount returns the RetainPreviousLDIFExportCount field value if set, zero value otherwise.
func (o *LdifExportRecurringTaskResponse) GetRetainPreviousLDIFExportCount() int32 {
	if o == nil || IsNil(o.RetainPreviousLDIFExportCount) {
		var ret int32
		return ret
	}
	return *o.RetainPreviousLDIFExportCount
}

// GetRetainPreviousLDIFExportCountOk returns a tuple with the RetainPreviousLDIFExportCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdifExportRecurringTaskResponse) GetRetainPreviousLDIFExportCountOk() (*int32, bool) {
	if o == nil || IsNil(o.RetainPreviousLDIFExportCount) {
		return nil, false
	}
	return o.RetainPreviousLDIFExportCount, true
}

// HasRetainPreviousLDIFExportCount returns a boolean if a field has been set.
func (o *LdifExportRecurringTaskResponse) HasRetainPreviousLDIFExportCount() bool {
	if o != nil && !IsNil(o.RetainPreviousLDIFExportCount) {
		return true
	}

	return false
}

// SetRetainPreviousLDIFExportCount gets a reference to the given int32 and assigns it to the RetainPreviousLDIFExportCount field.
func (o *LdifExportRecurringTaskResponse) SetRetainPreviousLDIFExportCount(v int32) {
	o.RetainPreviousLDIFExportCount = &v
}

// GetRetainPreviousLDIFExportAge returns the RetainPreviousLDIFExportAge field value if set, zero value otherwise.
func (o *LdifExportRecurringTaskResponse) GetRetainPreviousLDIFExportAge() string {
	if o == nil || IsNil(o.RetainPreviousLDIFExportAge) {
		var ret string
		return ret
	}
	return *o.RetainPreviousLDIFExportAge
}

// GetRetainPreviousLDIFExportAgeOk returns a tuple with the RetainPreviousLDIFExportAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdifExportRecurringTaskResponse) GetRetainPreviousLDIFExportAgeOk() (*string, bool) {
	if o == nil || IsNil(o.RetainPreviousLDIFExportAge) {
		return nil, false
	}
	return o.RetainPreviousLDIFExportAge, true
}

// HasRetainPreviousLDIFExportAge returns a boolean if a field has been set.
func (o *LdifExportRecurringTaskResponse) HasRetainPreviousLDIFExportAge() bool {
	if o != nil && !IsNil(o.RetainPreviousLDIFExportAge) {
		return true
	}

	return false
}

// SetRetainPreviousLDIFExportAge gets a reference to the given string and assigns it to the RetainPreviousLDIFExportAge field.
func (o *LdifExportRecurringTaskResponse) SetRetainPreviousLDIFExportAge(v string) {
	o.RetainPreviousLDIFExportAge = &v
}

// GetMaxMegabytesPerSecond returns the MaxMegabytesPerSecond field value if set, zero value otherwise.
func (o *LdifExportRecurringTaskResponse) GetMaxMegabytesPerSecond() int32 {
	if o == nil || IsNil(o.MaxMegabytesPerSecond) {
		var ret int32
		return ret
	}
	return *o.MaxMegabytesPerSecond
}

// GetMaxMegabytesPerSecondOk returns a tuple with the MaxMegabytesPerSecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdifExportRecurringTaskResponse) GetMaxMegabytesPerSecondOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxMegabytesPerSecond) {
		return nil, false
	}
	return o.MaxMegabytesPerSecond, true
}

// HasMaxMegabytesPerSecond returns a boolean if a field has been set.
func (o *LdifExportRecurringTaskResponse) HasMaxMegabytesPerSecond() bool {
	if o != nil && !IsNil(o.MaxMegabytesPerSecond) {
		return true
	}

	return false
}

// SetMaxMegabytesPerSecond gets a reference to the given int32 and assigns it to the MaxMegabytesPerSecond field.
func (o *LdifExportRecurringTaskResponse) SetMaxMegabytesPerSecond(v int32) {
	o.MaxMegabytesPerSecond = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LdifExportRecurringTaskResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdifExportRecurringTaskResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LdifExportRecurringTaskResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LdifExportRecurringTaskResponse) SetDescription(v string) {
	o.Description = &v
}

// GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field value if set, zero value otherwise.
func (o *LdifExportRecurringTaskResponse) GetCancelOnTaskDependencyFailure() bool {
	if o == nil || IsNil(o.CancelOnTaskDependencyFailure) {
		var ret bool
		return ret
	}
	return *o.CancelOnTaskDependencyFailure
}

// GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdifExportRecurringTaskResponse) GetCancelOnTaskDependencyFailureOk() (*bool, bool) {
	if o == nil || IsNil(o.CancelOnTaskDependencyFailure) {
		return nil, false
	}
	return o.CancelOnTaskDependencyFailure, true
}

// HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.
func (o *LdifExportRecurringTaskResponse) HasCancelOnTaskDependencyFailure() bool {
	if o != nil && !IsNil(o.CancelOnTaskDependencyFailure) {
		return true
	}

	return false
}

// SetCancelOnTaskDependencyFailure gets a reference to the given bool and assigns it to the CancelOnTaskDependencyFailure field.
func (o *LdifExportRecurringTaskResponse) SetCancelOnTaskDependencyFailure(v bool) {
	o.CancelOnTaskDependencyFailure = &v
}

// GetEmailOnStart returns the EmailOnStart field value if set, zero value otherwise.
func (o *LdifExportRecurringTaskResponse) GetEmailOnStart() []string {
	if o == nil || IsNil(o.EmailOnStart) {
		var ret []string
		return ret
	}
	return o.EmailOnStart
}

// GetEmailOnStartOk returns a tuple with the EmailOnStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdifExportRecurringTaskResponse) GetEmailOnStartOk() ([]string, bool) {
	if o == nil || IsNil(o.EmailOnStart) {
		return nil, false
	}
	return o.EmailOnStart, true
}

// HasEmailOnStart returns a boolean if a field has been set.
func (o *LdifExportRecurringTaskResponse) HasEmailOnStart() bool {
	if o != nil && !IsNil(o.EmailOnStart) {
		return true
	}

	return false
}

// SetEmailOnStart gets a reference to the given []string and assigns it to the EmailOnStart field.
func (o *LdifExportRecurringTaskResponse) SetEmailOnStart(v []string) {
	o.EmailOnStart = v
}

// GetEmailOnSuccess returns the EmailOnSuccess field value if set, zero value otherwise.
func (o *LdifExportRecurringTaskResponse) GetEmailOnSuccess() []string {
	if o == nil || IsNil(o.EmailOnSuccess) {
		var ret []string
		return ret
	}
	return o.EmailOnSuccess
}

// GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdifExportRecurringTaskResponse) GetEmailOnSuccessOk() ([]string, bool) {
	if o == nil || IsNil(o.EmailOnSuccess) {
		return nil, false
	}
	return o.EmailOnSuccess, true
}

// HasEmailOnSuccess returns a boolean if a field has been set.
func (o *LdifExportRecurringTaskResponse) HasEmailOnSuccess() bool {
	if o != nil && !IsNil(o.EmailOnSuccess) {
		return true
	}

	return false
}

// SetEmailOnSuccess gets a reference to the given []string and assigns it to the EmailOnSuccess field.
func (o *LdifExportRecurringTaskResponse) SetEmailOnSuccess(v []string) {
	o.EmailOnSuccess = v
}

// GetEmailOnFailure returns the EmailOnFailure field value if set, zero value otherwise.
func (o *LdifExportRecurringTaskResponse) GetEmailOnFailure() []string {
	if o == nil || IsNil(o.EmailOnFailure) {
		var ret []string
		return ret
	}
	return o.EmailOnFailure
}

// GetEmailOnFailureOk returns a tuple with the EmailOnFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdifExportRecurringTaskResponse) GetEmailOnFailureOk() ([]string, bool) {
	if o == nil || IsNil(o.EmailOnFailure) {
		return nil, false
	}
	return o.EmailOnFailure, true
}

// HasEmailOnFailure returns a boolean if a field has been set.
func (o *LdifExportRecurringTaskResponse) HasEmailOnFailure() bool {
	if o != nil && !IsNil(o.EmailOnFailure) {
		return true
	}

	return false
}

// SetEmailOnFailure gets a reference to the given []string and assigns it to the EmailOnFailure field.
func (o *LdifExportRecurringTaskResponse) SetEmailOnFailure(v []string) {
	o.EmailOnFailure = v
}

// GetAlertOnStart returns the AlertOnStart field value if set, zero value otherwise.
func (o *LdifExportRecurringTaskResponse) GetAlertOnStart() bool {
	if o == nil || IsNil(o.AlertOnStart) {
		var ret bool
		return ret
	}
	return *o.AlertOnStart
}

// GetAlertOnStartOk returns a tuple with the AlertOnStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdifExportRecurringTaskResponse) GetAlertOnStartOk() (*bool, bool) {
	if o == nil || IsNil(o.AlertOnStart) {
		return nil, false
	}
	return o.AlertOnStart, true
}

// HasAlertOnStart returns a boolean if a field has been set.
func (o *LdifExportRecurringTaskResponse) HasAlertOnStart() bool {
	if o != nil && !IsNil(o.AlertOnStart) {
		return true
	}

	return false
}

// SetAlertOnStart gets a reference to the given bool and assigns it to the AlertOnStart field.
func (o *LdifExportRecurringTaskResponse) SetAlertOnStart(v bool) {
	o.AlertOnStart = &v
}

// GetAlertOnSuccess returns the AlertOnSuccess field value if set, zero value otherwise.
func (o *LdifExportRecurringTaskResponse) GetAlertOnSuccess() bool {
	if o == nil || IsNil(o.AlertOnSuccess) {
		var ret bool
		return ret
	}
	return *o.AlertOnSuccess
}

// GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdifExportRecurringTaskResponse) GetAlertOnSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.AlertOnSuccess) {
		return nil, false
	}
	return o.AlertOnSuccess, true
}

// HasAlertOnSuccess returns a boolean if a field has been set.
func (o *LdifExportRecurringTaskResponse) HasAlertOnSuccess() bool {
	if o != nil && !IsNil(o.AlertOnSuccess) {
		return true
	}

	return false
}

// SetAlertOnSuccess gets a reference to the given bool and assigns it to the AlertOnSuccess field.
func (o *LdifExportRecurringTaskResponse) SetAlertOnSuccess(v bool) {
	o.AlertOnSuccess = &v
}

// GetAlertOnFailure returns the AlertOnFailure field value if set, zero value otherwise.
func (o *LdifExportRecurringTaskResponse) GetAlertOnFailure() bool {
	if o == nil || IsNil(o.AlertOnFailure) {
		var ret bool
		return ret
	}
	return *o.AlertOnFailure
}

// GetAlertOnFailureOk returns a tuple with the AlertOnFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdifExportRecurringTaskResponse) GetAlertOnFailureOk() (*bool, bool) {
	if o == nil || IsNil(o.AlertOnFailure) {
		return nil, false
	}
	return o.AlertOnFailure, true
}

// HasAlertOnFailure returns a boolean if a field has been set.
func (o *LdifExportRecurringTaskResponse) HasAlertOnFailure() bool {
	if o != nil && !IsNil(o.AlertOnFailure) {
		return true
	}

	return false
}

// SetAlertOnFailure gets a reference to the given bool and assigns it to the AlertOnFailure field.
func (o *LdifExportRecurringTaskResponse) SetAlertOnFailure(v bool) {
	o.AlertOnFailure = &v
}

func (o LdifExportRecurringTaskResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LdifExportRecurringTaskResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	toSerialize["schemas"] = o.Schemas
	toSerialize["ldifDirectory"] = o.LdifDirectory
	if !IsNil(o.BackendID) {
		toSerialize["backendID"] = o.BackendID
	}
	if !IsNil(o.ExcludeBackendID) {
		toSerialize["excludeBackendID"] = o.ExcludeBackendID
	}
	if !IsNil(o.Compress) {
		toSerialize["compress"] = o.Compress
	}
	if !IsNil(o.Encrypt) {
		toSerialize["encrypt"] = o.Encrypt
	}
	if !IsNil(o.EncryptionSettingsDefinitionID) {
		toSerialize["encryptionSettingsDefinitionID"] = o.EncryptionSettingsDefinitionID
	}
	if !IsNil(o.Sign) {
		toSerialize["sign"] = o.Sign
	}
	if !IsNil(o.RetainPreviousLDIFExportCount) {
		toSerialize["retainPreviousLDIFExportCount"] = o.RetainPreviousLDIFExportCount
	}
	if !IsNil(o.RetainPreviousLDIFExportAge) {
		toSerialize["retainPreviousLDIFExportAge"] = o.RetainPreviousLDIFExportAge
	}
	if !IsNil(o.MaxMegabytesPerSecond) {
		toSerialize["maxMegabytesPerSecond"] = o.MaxMegabytesPerSecond
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CancelOnTaskDependencyFailure) {
		toSerialize["cancelOnTaskDependencyFailure"] = o.CancelOnTaskDependencyFailure
	}
	if !IsNil(o.EmailOnStart) {
		toSerialize["emailOnStart"] = o.EmailOnStart
	}
	if !IsNil(o.EmailOnSuccess) {
		toSerialize["emailOnSuccess"] = o.EmailOnSuccess
	}
	if !IsNil(o.EmailOnFailure) {
		toSerialize["emailOnFailure"] = o.EmailOnFailure
	}
	if !IsNil(o.AlertOnStart) {
		toSerialize["alertOnStart"] = o.AlertOnStart
	}
	if !IsNil(o.AlertOnSuccess) {
		toSerialize["alertOnSuccess"] = o.AlertOnSuccess
	}
	if !IsNil(o.AlertOnFailure) {
		toSerialize["alertOnFailure"] = o.AlertOnFailure
	}
	return toSerialize, nil
}

type NullableLdifExportRecurringTaskResponse struct {
	value *LdifExportRecurringTaskResponse
	isSet bool
}

func (v NullableLdifExportRecurringTaskResponse) Get() *LdifExportRecurringTaskResponse {
	return v.value
}

func (v *NullableLdifExportRecurringTaskResponse) Set(val *LdifExportRecurringTaskResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLdifExportRecurringTaskResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLdifExportRecurringTaskResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdifExportRecurringTaskResponse(val *LdifExportRecurringTaskResponse) *NullableLdifExportRecurringTaskResponse {
	return &NullableLdifExportRecurringTaskResponse{value: val, isSet: true}
}

func (v NullableLdifExportRecurringTaskResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdifExportRecurringTaskResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
