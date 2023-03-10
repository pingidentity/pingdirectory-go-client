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

// checks if the AddIdleAccountDataSecurityAuditorRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddIdleAccountDataSecurityAuditorRequest{}

// AddIdleAccountDataSecurityAuditorRequest struct for AddIdleAccountDataSecurityAuditorRequest
type AddIdleAccountDataSecurityAuditorRequest struct {
	// Name of the new Data Security Auditor
	AuditorName string                                        `json:"auditorName"`
	Schemas     []EnumidleAccountDataSecurityAuditorSchemaUrn `json:"schemas"`
	// Specifies the name of the detailed report file.
	ReportFile *string `json:"reportFile,omitempty"`
	// The length of time to use as the warning interval for idle accounts. If the length of time since a user last authenticated is greater than the warning interval but less than the error interval (or if it is greater than the warning interval and no error interval is defined), then a warning will be generated for that account.
	IdleAccountWarningInterval string `json:"idleAccountWarningInterval"`
	// The length of time to use as the error interval for idle accounts. If the length of time since a user last authenticated is greater than the error interval, then an error will be generated for that account. If no error interval is defined, then only the warning interval will be used.
	IdleAccountErrorInterval *string `json:"idleAccountErrorInterval,omitempty"`
	// The length of time to use as the warning interval for accounts that do not appear to have authenticated. If this is not specified, then the idle account warning interval will be used.
	NeverLoggedInAccountWarningInterval *string `json:"neverLoggedInAccountWarningInterval,omitempty"`
	// The length of time to use as the error interval for accounts that do not appear to have authenticated. If this is not specified, then the never-logged-in warning interval will be used. The idle account warning and error intervals will be used if no never-logged-in interval is configured.
	NeverLoggedInAccountErrorInterval *string `json:"neverLoggedInAccountErrorInterval,omitempty"`
	// Indicates whether the Data Security Auditor is enabled for use.
	Enabled *bool `json:"enabled,omitempty"`
	// Specifies the attributes from the audited entries that should be included detailed reports. By default, no attributes are included.
	IncludeAttribute []string `json:"includeAttribute,omitempty"`
	// Specifies which backends the data security auditor may be applied to. By default, the data security auditors will audit entries in all backend types that support data auditing (Local DB, LDIF, and Config File Handler).
	AuditBackend  []string                                  `json:"auditBackend,omitempty"`
	AuditSeverity *EnumdataSecurityAuditorAuditSeverityProp `json:"auditSeverity,omitempty"`
}

// NewAddIdleAccountDataSecurityAuditorRequest instantiates a new AddIdleAccountDataSecurityAuditorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddIdleAccountDataSecurityAuditorRequest(auditorName string, schemas []EnumidleAccountDataSecurityAuditorSchemaUrn, idleAccountWarningInterval string) *AddIdleAccountDataSecurityAuditorRequest {
	this := AddIdleAccountDataSecurityAuditorRequest{}
	this.AuditorName = auditorName
	this.Schemas = schemas
	this.IdleAccountWarningInterval = idleAccountWarningInterval
	return &this
}

// NewAddIdleAccountDataSecurityAuditorRequestWithDefaults instantiates a new AddIdleAccountDataSecurityAuditorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddIdleAccountDataSecurityAuditorRequestWithDefaults() *AddIdleAccountDataSecurityAuditorRequest {
	this := AddIdleAccountDataSecurityAuditorRequest{}
	return &this
}

// GetAuditorName returns the AuditorName field value
func (o *AddIdleAccountDataSecurityAuditorRequest) GetAuditorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuditorName
}

// GetAuditorNameOk returns a tuple with the AuditorName field value
// and a boolean to check if the value has been set.
func (o *AddIdleAccountDataSecurityAuditorRequest) GetAuditorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuditorName, true
}

// SetAuditorName sets field value
func (o *AddIdleAccountDataSecurityAuditorRequest) SetAuditorName(v string) {
	o.AuditorName = v
}

// GetSchemas returns the Schemas field value
func (o *AddIdleAccountDataSecurityAuditorRequest) GetSchemas() []EnumidleAccountDataSecurityAuditorSchemaUrn {
	if o == nil {
		var ret []EnumidleAccountDataSecurityAuditorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddIdleAccountDataSecurityAuditorRequest) GetSchemasOk() ([]EnumidleAccountDataSecurityAuditorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddIdleAccountDataSecurityAuditorRequest) SetSchemas(v []EnumidleAccountDataSecurityAuditorSchemaUrn) {
	o.Schemas = v
}

// GetReportFile returns the ReportFile field value if set, zero value otherwise.
func (o *AddIdleAccountDataSecurityAuditorRequest) GetReportFile() string {
	if o == nil || IsNil(o.ReportFile) {
		var ret string
		return ret
	}
	return *o.ReportFile
}

// GetReportFileOk returns a tuple with the ReportFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIdleAccountDataSecurityAuditorRequest) GetReportFileOk() (*string, bool) {
	if o == nil || IsNil(o.ReportFile) {
		return nil, false
	}
	return o.ReportFile, true
}

// HasReportFile returns a boolean if a field has been set.
func (o *AddIdleAccountDataSecurityAuditorRequest) HasReportFile() bool {
	if o != nil && !IsNil(o.ReportFile) {
		return true
	}

	return false
}

// SetReportFile gets a reference to the given string and assigns it to the ReportFile field.
func (o *AddIdleAccountDataSecurityAuditorRequest) SetReportFile(v string) {
	o.ReportFile = &v
}

// GetIdleAccountWarningInterval returns the IdleAccountWarningInterval field value
func (o *AddIdleAccountDataSecurityAuditorRequest) GetIdleAccountWarningInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdleAccountWarningInterval
}

// GetIdleAccountWarningIntervalOk returns a tuple with the IdleAccountWarningInterval field value
// and a boolean to check if the value has been set.
func (o *AddIdleAccountDataSecurityAuditorRequest) GetIdleAccountWarningIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdleAccountWarningInterval, true
}

// SetIdleAccountWarningInterval sets field value
func (o *AddIdleAccountDataSecurityAuditorRequest) SetIdleAccountWarningInterval(v string) {
	o.IdleAccountWarningInterval = v
}

// GetIdleAccountErrorInterval returns the IdleAccountErrorInterval field value if set, zero value otherwise.
func (o *AddIdleAccountDataSecurityAuditorRequest) GetIdleAccountErrorInterval() string {
	if o == nil || IsNil(o.IdleAccountErrorInterval) {
		var ret string
		return ret
	}
	return *o.IdleAccountErrorInterval
}

// GetIdleAccountErrorIntervalOk returns a tuple with the IdleAccountErrorInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIdleAccountDataSecurityAuditorRequest) GetIdleAccountErrorIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.IdleAccountErrorInterval) {
		return nil, false
	}
	return o.IdleAccountErrorInterval, true
}

// HasIdleAccountErrorInterval returns a boolean if a field has been set.
func (o *AddIdleAccountDataSecurityAuditorRequest) HasIdleAccountErrorInterval() bool {
	if o != nil && !IsNil(o.IdleAccountErrorInterval) {
		return true
	}

	return false
}

// SetIdleAccountErrorInterval gets a reference to the given string and assigns it to the IdleAccountErrorInterval field.
func (o *AddIdleAccountDataSecurityAuditorRequest) SetIdleAccountErrorInterval(v string) {
	o.IdleAccountErrorInterval = &v
}

// GetNeverLoggedInAccountWarningInterval returns the NeverLoggedInAccountWarningInterval field value if set, zero value otherwise.
func (o *AddIdleAccountDataSecurityAuditorRequest) GetNeverLoggedInAccountWarningInterval() string {
	if o == nil || IsNil(o.NeverLoggedInAccountWarningInterval) {
		var ret string
		return ret
	}
	return *o.NeverLoggedInAccountWarningInterval
}

// GetNeverLoggedInAccountWarningIntervalOk returns a tuple with the NeverLoggedInAccountWarningInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIdleAccountDataSecurityAuditorRequest) GetNeverLoggedInAccountWarningIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.NeverLoggedInAccountWarningInterval) {
		return nil, false
	}
	return o.NeverLoggedInAccountWarningInterval, true
}

// HasNeverLoggedInAccountWarningInterval returns a boolean if a field has been set.
func (o *AddIdleAccountDataSecurityAuditorRequest) HasNeverLoggedInAccountWarningInterval() bool {
	if o != nil && !IsNil(o.NeverLoggedInAccountWarningInterval) {
		return true
	}

	return false
}

// SetNeverLoggedInAccountWarningInterval gets a reference to the given string and assigns it to the NeverLoggedInAccountWarningInterval field.
func (o *AddIdleAccountDataSecurityAuditorRequest) SetNeverLoggedInAccountWarningInterval(v string) {
	o.NeverLoggedInAccountWarningInterval = &v
}

// GetNeverLoggedInAccountErrorInterval returns the NeverLoggedInAccountErrorInterval field value if set, zero value otherwise.
func (o *AddIdleAccountDataSecurityAuditorRequest) GetNeverLoggedInAccountErrorInterval() string {
	if o == nil || IsNil(o.NeverLoggedInAccountErrorInterval) {
		var ret string
		return ret
	}
	return *o.NeverLoggedInAccountErrorInterval
}

// GetNeverLoggedInAccountErrorIntervalOk returns a tuple with the NeverLoggedInAccountErrorInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIdleAccountDataSecurityAuditorRequest) GetNeverLoggedInAccountErrorIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.NeverLoggedInAccountErrorInterval) {
		return nil, false
	}
	return o.NeverLoggedInAccountErrorInterval, true
}

// HasNeverLoggedInAccountErrorInterval returns a boolean if a field has been set.
func (o *AddIdleAccountDataSecurityAuditorRequest) HasNeverLoggedInAccountErrorInterval() bool {
	if o != nil && !IsNil(o.NeverLoggedInAccountErrorInterval) {
		return true
	}

	return false
}

// SetNeverLoggedInAccountErrorInterval gets a reference to the given string and assigns it to the NeverLoggedInAccountErrorInterval field.
func (o *AddIdleAccountDataSecurityAuditorRequest) SetNeverLoggedInAccountErrorInterval(v string) {
	o.NeverLoggedInAccountErrorInterval = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AddIdleAccountDataSecurityAuditorRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIdleAccountDataSecurityAuditorRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AddIdleAccountDataSecurityAuditorRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AddIdleAccountDataSecurityAuditorRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIncludeAttribute returns the IncludeAttribute field value if set, zero value otherwise.
func (o *AddIdleAccountDataSecurityAuditorRequest) GetIncludeAttribute() []string {
	if o == nil || IsNil(o.IncludeAttribute) {
		var ret []string
		return ret
	}
	return o.IncludeAttribute
}

// GetIncludeAttributeOk returns a tuple with the IncludeAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIdleAccountDataSecurityAuditorRequest) GetIncludeAttributeOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeAttribute) {
		return nil, false
	}
	return o.IncludeAttribute, true
}

// HasIncludeAttribute returns a boolean if a field has been set.
func (o *AddIdleAccountDataSecurityAuditorRequest) HasIncludeAttribute() bool {
	if o != nil && !IsNil(o.IncludeAttribute) {
		return true
	}

	return false
}

// SetIncludeAttribute gets a reference to the given []string and assigns it to the IncludeAttribute field.
func (o *AddIdleAccountDataSecurityAuditorRequest) SetIncludeAttribute(v []string) {
	o.IncludeAttribute = v
}

// GetAuditBackend returns the AuditBackend field value if set, zero value otherwise.
func (o *AddIdleAccountDataSecurityAuditorRequest) GetAuditBackend() []string {
	if o == nil || IsNil(o.AuditBackend) {
		var ret []string
		return ret
	}
	return o.AuditBackend
}

// GetAuditBackendOk returns a tuple with the AuditBackend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIdleAccountDataSecurityAuditorRequest) GetAuditBackendOk() ([]string, bool) {
	if o == nil || IsNil(o.AuditBackend) {
		return nil, false
	}
	return o.AuditBackend, true
}

// HasAuditBackend returns a boolean if a field has been set.
func (o *AddIdleAccountDataSecurityAuditorRequest) HasAuditBackend() bool {
	if o != nil && !IsNil(o.AuditBackend) {
		return true
	}

	return false
}

// SetAuditBackend gets a reference to the given []string and assigns it to the AuditBackend field.
func (o *AddIdleAccountDataSecurityAuditorRequest) SetAuditBackend(v []string) {
	o.AuditBackend = v
}

// GetAuditSeverity returns the AuditSeverity field value if set, zero value otherwise.
func (o *AddIdleAccountDataSecurityAuditorRequest) GetAuditSeverity() EnumdataSecurityAuditorAuditSeverityProp {
	if o == nil || IsNil(o.AuditSeverity) {
		var ret EnumdataSecurityAuditorAuditSeverityProp
		return ret
	}
	return *o.AuditSeverity
}

// GetAuditSeverityOk returns a tuple with the AuditSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIdleAccountDataSecurityAuditorRequest) GetAuditSeverityOk() (*EnumdataSecurityAuditorAuditSeverityProp, bool) {
	if o == nil || IsNil(o.AuditSeverity) {
		return nil, false
	}
	return o.AuditSeverity, true
}

// HasAuditSeverity returns a boolean if a field has been set.
func (o *AddIdleAccountDataSecurityAuditorRequest) HasAuditSeverity() bool {
	if o != nil && !IsNil(o.AuditSeverity) {
		return true
	}

	return false
}

// SetAuditSeverity gets a reference to the given EnumdataSecurityAuditorAuditSeverityProp and assigns it to the AuditSeverity field.
func (o *AddIdleAccountDataSecurityAuditorRequest) SetAuditSeverity(v EnumdataSecurityAuditorAuditSeverityProp) {
	o.AuditSeverity = &v
}

func (o AddIdleAccountDataSecurityAuditorRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddIdleAccountDataSecurityAuditorRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["auditorName"] = o.AuditorName
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.ReportFile) {
		toSerialize["reportFile"] = o.ReportFile
	}
	toSerialize["idleAccountWarningInterval"] = o.IdleAccountWarningInterval
	if !IsNil(o.IdleAccountErrorInterval) {
		toSerialize["idleAccountErrorInterval"] = o.IdleAccountErrorInterval
	}
	if !IsNil(o.NeverLoggedInAccountWarningInterval) {
		toSerialize["neverLoggedInAccountWarningInterval"] = o.NeverLoggedInAccountWarningInterval
	}
	if !IsNil(o.NeverLoggedInAccountErrorInterval) {
		toSerialize["neverLoggedInAccountErrorInterval"] = o.NeverLoggedInAccountErrorInterval
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.IncludeAttribute) {
		toSerialize["includeAttribute"] = o.IncludeAttribute
	}
	if !IsNil(o.AuditBackend) {
		toSerialize["auditBackend"] = o.AuditBackend
	}
	if !IsNil(o.AuditSeverity) {
		toSerialize["auditSeverity"] = o.AuditSeverity
	}
	return toSerialize, nil
}

type NullableAddIdleAccountDataSecurityAuditorRequest struct {
	value *AddIdleAccountDataSecurityAuditorRequest
	isSet bool
}

func (v NullableAddIdleAccountDataSecurityAuditorRequest) Get() *AddIdleAccountDataSecurityAuditorRequest {
	return v.value
}

func (v *NullableAddIdleAccountDataSecurityAuditorRequest) Set(val *AddIdleAccountDataSecurityAuditorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddIdleAccountDataSecurityAuditorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddIdleAccountDataSecurityAuditorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddIdleAccountDataSecurityAuditorRequest(val *AddIdleAccountDataSecurityAuditorRequest) *NullableAddIdleAccountDataSecurityAuditorRequest {
	return &NullableAddIdleAccountDataSecurityAuditorRequest{value: val, isSet: true}
}

func (v NullableAddIdleAccountDataSecurityAuditorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddIdleAccountDataSecurityAuditorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
