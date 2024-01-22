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

// checks if the AddLockedAccountDataSecurityAuditorRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddLockedAccountDataSecurityAuditorRequest{}

// AddLockedAccountDataSecurityAuditorRequest struct for AddLockedAccountDataSecurityAuditorRequest
type AddLockedAccountDataSecurityAuditorRequest struct {
	Schemas []EnumlockedAccountDataSecurityAuditorSchemaUrn `json:"schemas"`
	// Specifies the name of the detailed report file.
	ReportFile *string `json:"reportFile,omitempty"`
	// Specifies the attributes from the audited entries that should be included detailed reports. By default, no attributes are included.
	IncludeAttribute []string `json:"includeAttribute,omitempty"`
	// If set, users that have not authenticated for more than the specified time will be reported even if idle account lockout is not configured. Note that users may only be reported if the last login time tracking is enabled.
	MaximumIdleTime *string `json:"maximumIdleTime,omitempty"`
	// Indicates whether the Data Security Auditor is enabled for use.
	Enabled *bool `json:"enabled,omitempty"`
	// Specifies which backends the data security auditor may be applied to. By default, the data security auditors will audit entries in all backend types that support data auditing (Local DB, LDIF, and Config File Handler).
	AuditBackend  []string                                  `json:"auditBackend,omitempty"`
	AuditSeverity *EnumdataSecurityAuditorAuditSeverityProp `json:"auditSeverity,omitempty"`
	// Name of the new Data Security Auditor
	AuditorName string `json:"auditorName"`
}

// NewAddLockedAccountDataSecurityAuditorRequest instantiates a new AddLockedAccountDataSecurityAuditorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddLockedAccountDataSecurityAuditorRequest(schemas []EnumlockedAccountDataSecurityAuditorSchemaUrn, auditorName string) *AddLockedAccountDataSecurityAuditorRequest {
	this := AddLockedAccountDataSecurityAuditorRequest{}
	this.Schemas = schemas
	this.AuditorName = auditorName
	return &this
}

// NewAddLockedAccountDataSecurityAuditorRequestWithDefaults instantiates a new AddLockedAccountDataSecurityAuditorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddLockedAccountDataSecurityAuditorRequestWithDefaults() *AddLockedAccountDataSecurityAuditorRequest {
	this := AddLockedAccountDataSecurityAuditorRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddLockedAccountDataSecurityAuditorRequest) GetSchemas() []EnumlockedAccountDataSecurityAuditorSchemaUrn {
	if o == nil {
		var ret []EnumlockedAccountDataSecurityAuditorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddLockedAccountDataSecurityAuditorRequest) GetSchemasOk() ([]EnumlockedAccountDataSecurityAuditorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddLockedAccountDataSecurityAuditorRequest) SetSchemas(v []EnumlockedAccountDataSecurityAuditorSchemaUrn) {
	o.Schemas = v
}

// GetReportFile returns the ReportFile field value if set, zero value otherwise.
func (o *AddLockedAccountDataSecurityAuditorRequest) GetReportFile() string {
	if o == nil || IsNil(o.ReportFile) {
		var ret string
		return ret
	}
	return *o.ReportFile
}

// GetReportFileOk returns a tuple with the ReportFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLockedAccountDataSecurityAuditorRequest) GetReportFileOk() (*string, bool) {
	if o == nil || IsNil(o.ReportFile) {
		return nil, false
	}
	return o.ReportFile, true
}

// HasReportFile returns a boolean if a field has been set.
func (o *AddLockedAccountDataSecurityAuditorRequest) HasReportFile() bool {
	if o != nil && !IsNil(o.ReportFile) {
		return true
	}

	return false
}

// SetReportFile gets a reference to the given string and assigns it to the ReportFile field.
func (o *AddLockedAccountDataSecurityAuditorRequest) SetReportFile(v string) {
	o.ReportFile = &v
}

// GetIncludeAttribute returns the IncludeAttribute field value if set, zero value otherwise.
func (o *AddLockedAccountDataSecurityAuditorRequest) GetIncludeAttribute() []string {
	if o == nil || IsNil(o.IncludeAttribute) {
		var ret []string
		return ret
	}
	return o.IncludeAttribute
}

// GetIncludeAttributeOk returns a tuple with the IncludeAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLockedAccountDataSecurityAuditorRequest) GetIncludeAttributeOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeAttribute) {
		return nil, false
	}
	return o.IncludeAttribute, true
}

// HasIncludeAttribute returns a boolean if a field has been set.
func (o *AddLockedAccountDataSecurityAuditorRequest) HasIncludeAttribute() bool {
	if o != nil && !IsNil(o.IncludeAttribute) {
		return true
	}

	return false
}

// SetIncludeAttribute gets a reference to the given []string and assigns it to the IncludeAttribute field.
func (o *AddLockedAccountDataSecurityAuditorRequest) SetIncludeAttribute(v []string) {
	o.IncludeAttribute = v
}

// GetMaximumIdleTime returns the MaximumIdleTime field value if set, zero value otherwise.
func (o *AddLockedAccountDataSecurityAuditorRequest) GetMaximumIdleTime() string {
	if o == nil || IsNil(o.MaximumIdleTime) {
		var ret string
		return ret
	}
	return *o.MaximumIdleTime
}

// GetMaximumIdleTimeOk returns a tuple with the MaximumIdleTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLockedAccountDataSecurityAuditorRequest) GetMaximumIdleTimeOk() (*string, bool) {
	if o == nil || IsNil(o.MaximumIdleTime) {
		return nil, false
	}
	return o.MaximumIdleTime, true
}

// HasMaximumIdleTime returns a boolean if a field has been set.
func (o *AddLockedAccountDataSecurityAuditorRequest) HasMaximumIdleTime() bool {
	if o != nil && !IsNil(o.MaximumIdleTime) {
		return true
	}

	return false
}

// SetMaximumIdleTime gets a reference to the given string and assigns it to the MaximumIdleTime field.
func (o *AddLockedAccountDataSecurityAuditorRequest) SetMaximumIdleTime(v string) {
	o.MaximumIdleTime = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AddLockedAccountDataSecurityAuditorRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLockedAccountDataSecurityAuditorRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AddLockedAccountDataSecurityAuditorRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AddLockedAccountDataSecurityAuditorRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAuditBackend returns the AuditBackend field value if set, zero value otherwise.
func (o *AddLockedAccountDataSecurityAuditorRequest) GetAuditBackend() []string {
	if o == nil || IsNil(o.AuditBackend) {
		var ret []string
		return ret
	}
	return o.AuditBackend
}

// GetAuditBackendOk returns a tuple with the AuditBackend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLockedAccountDataSecurityAuditorRequest) GetAuditBackendOk() ([]string, bool) {
	if o == nil || IsNil(o.AuditBackend) {
		return nil, false
	}
	return o.AuditBackend, true
}

// HasAuditBackend returns a boolean if a field has been set.
func (o *AddLockedAccountDataSecurityAuditorRequest) HasAuditBackend() bool {
	if o != nil && !IsNil(o.AuditBackend) {
		return true
	}

	return false
}

// SetAuditBackend gets a reference to the given []string and assigns it to the AuditBackend field.
func (o *AddLockedAccountDataSecurityAuditorRequest) SetAuditBackend(v []string) {
	o.AuditBackend = v
}

// GetAuditSeverity returns the AuditSeverity field value if set, zero value otherwise.
func (o *AddLockedAccountDataSecurityAuditorRequest) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp {
	if o == nil || IsNil(o.AuditSeverity) {
		var ret EnumdataSecurityAuditorAuditSeverityProp
		return ret
	}
	return *o.AuditSeverity
}

// GetAuditSeverityOk returns a tuple with the AuditSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLockedAccountDataSecurityAuditorRequest) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool) {
	if o == nil || IsNil(o.AuditSeverity) {
		return nil, false
	}
	return o.AuditSeverity, true
}

// HasAuditSeverity returns a boolean if a field has been set.
func (o *AddLockedAccountDataSecurityAuditorRequest) HasAuditSeverity() bool {
	if o != nil && !IsNil(o.AuditSeverity) {
		return true
	}

	return false
}

// SetAuditSeverity gets a reference to the given EnumdataSecurityAuditorAuditSeverityProp and assigns it to the AuditSeverity field.
func (o *AddLockedAccountDataSecurityAuditorRequest) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp) {
	o.AuditSeverity = &v
}

// GetAuditorName returns the AuditorName field value
func (o *AddLockedAccountDataSecurityAuditorRequest) GetAuditorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuditorName
}

// GetAuditorNameOk returns a tuple with the AuditorName field value
// and a boolean to check if the value has been set.
func (o *AddLockedAccountDataSecurityAuditorRequest) GetAuditorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuditorName, true
}

// SetAuditorName sets field value
func (o *AddLockedAccountDataSecurityAuditorRequest) SetAuditorName(v string) {
	o.AuditorName = v
}

func (o AddLockedAccountDataSecurityAuditorRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddLockedAccountDataSecurityAuditorRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.ReportFile) {
		toSerialize["reportFile"] = o.ReportFile
	}
	if !IsNil(o.IncludeAttribute) {
		toSerialize["includeAttribute"] = o.IncludeAttribute
	}
	if !IsNil(o.MaximumIdleTime) {
		toSerialize["maximumIdleTime"] = o.MaximumIdleTime
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.AuditBackend) {
		toSerialize["auditBackend"] = o.AuditBackend
	}
	if !IsNil(o.AuditSeverity) {
		toSerialize["auditSeverity"] = o.AuditSeverity
	}
	toSerialize["auditorName"] = o.AuditorName
	return toSerialize, nil
}

type NullableAddLockedAccountDataSecurityAuditorRequest struct {
	value *AddLockedAccountDataSecurityAuditorRequest
	isSet bool
}

func (v NullableAddLockedAccountDataSecurityAuditorRequest) Get() *AddLockedAccountDataSecurityAuditorRequest {
	return v.value
}

func (v *NullableAddLockedAccountDataSecurityAuditorRequest) Set(val *AddLockedAccountDataSecurityAuditorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddLockedAccountDataSecurityAuditorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddLockedAccountDataSecurityAuditorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddLockedAccountDataSecurityAuditorRequest(val *AddLockedAccountDataSecurityAuditorRequest) *NullableAddLockedAccountDataSecurityAuditorRequest {
	return &NullableAddLockedAccountDataSecurityAuditorRequest{value: val, isSet: true}
}

func (v NullableAddLockedAccountDataSecurityAuditorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddLockedAccountDataSecurityAuditorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
